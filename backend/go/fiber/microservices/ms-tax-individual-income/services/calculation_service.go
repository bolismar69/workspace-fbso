// path: backend/go/fiber/microservices/ms-tax-individual-income/services/calculation_service.go
package services

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"taxnexus-individual-core-lib/models"
	"taxnexus-individual-core-lib/repository"

	"github.com/shopspring/decimal"
)

type CalculationResult struct {
	Key string
	Res models.TaxResponse
	Err error
}

type CalculationService struct {
	repo       *repository.TaxRepository
	inssClient *INSSClient // Novo campo para comunicação com MS de INSS
	logger     *slog.Logger
}

func NewCalculationService(repo *repository.TaxRepository, inssURL string) *CalculationService {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	return &CalculationService{
		repo:       repo,
		inssClient: NewINSSClient(inssURL), // Inicializa o cliente
		logger:     slog.New(handler),
	}
}

func (s *CalculationService) getLogger(ctx context.Context) *slog.Logger {
	if tid, ok := ctx.Value("requestid").(string); ok {
		return s.logger.With("trace_id", tid)
	}
	return s.logger
}

// func (s *CalculationService) Calculate(ctx context.Context, req models.UniversalTaxRequest) (map[string]models.TaxResponse, error) {
// 	l := s.getLogger(ctx)

// 	if req.ReferenceDate.IsZero() {
// 		req.ReferenceDate = time.Now()
// 	}

// 	configs, err := s.repo.GetTableConfigs(ctx, req.TaxCode, req.ReferenceDate)
// 	if err != nil {
// 		l.Error("erro ao carregar tax_configs", "error", err, "tax_code", req.TaxCode)
// 		return nil, fmt.Errorf("failed to load tax configurations: %w", err)
// 	}

// 	inputMap := make(map[string]decimal.Decimal)
// 	for _, in := range req.Inputs {
// 		inputMap[in.Key] = in.Value
// 	}

// 	results := make(map[string]models.TaxResponse)

// 	resCompleta, err := s.calculateCompleta(ctx, req, inputMap, configs)
// 	if err == nil {
// 		results["completa"] = resCompleta
// 	}

// 	resSimplificada, err := s.calculateSimplificada(ctx, req, configs)
// 	if err == nil {
// 		results["simplificada"] = resSimplificada
// 	}

// 	if len(results) == 2 {
// 		if results["simplificada"].TaxAmount.LessThan(results["completa"].TaxAmount) {
// 			sImp := results["simplificada"]; sImp.IsRecommended = true; results["simplificada"] = sImp
// 		} else {
// 			cImp := results["completa"]; cImp.IsRecommended = true; results["completa"] = cImp
// 		}
// 	}

// 	return results, nil
// }

func (s *CalculationService) Calculate(ctx context.Context, req models.UniversalTaxRequest) (map[string]models.TaxResponse, error) {
	l := s.getLogger(ctx)
	start := time.Now()

	if req.ReferenceDate.IsZero() {
		req.ReferenceDate = time.Now()
	}

	// Carrega configs (Uma única vez para ambos os cálculos)
	configs, err := s.repo.GetTableConfigs(ctx, req.TaxCode, req.ReferenceDate)
	if err != nil {
		return nil, err
	}

	inputMap := make(map[string]decimal.Decimal)
	for _, in := range req.Inputs {
		inputMap[in.Key] = in.Value
	}

	resChan := make(chan CalculationResult, 2)

	// Dispara cálculos em paralelo
	go func() {
		res, err := s.calculateCompleta(ctx, req, inputMap, configs)
		resChan <- CalculationResult{"completa", res, err}
	}()

	go func() {
		res, err := s.calculateSimplificada(ctx, req, configs)
		resChan <- CalculationResult{"simplificada", res, err}
	}()

	results := make(map[string]models.TaxResponse)
	for i := 0; i < 2; i++ {
		result := <-resChan
		if result.Err == nil {
			results[result.Key] = result.Res
		} else {
			l.Error("erro no calculo paralelo", "type", result.Key, "error", result.Err)
		}
	}

	// Lógica de recomendação
	if c, okC := results["completa"]; okC {
		if s, okS := results["simplificada"]; okS {
			if s.TaxAmount.LessThan(c.TaxAmount) {
				s.IsRecommended = true
				results["simplificada"] = s
			} else {
				c.IsRecommended = true
				results["completa"] = c
			}
		}
	}

	l.Info("calculo IRPF finalizado", "duration_ms", time.Since(start).Milliseconds())
	return results, nil
}

func (s *CalculationService) calculateCompleta(ctx context.Context, req models.UniversalTaxRequest, inputs, configs map[string]decimal.Decimal) (models.TaxResponse, error) {
	l := s.getLogger(ctx)
	var details []models.DeductionDetail
	totalDeduction := decimal.Zero

	// --- INTEGRAÇÃO INSS ---
	l.Debug("solicitando calculo de INSS externo")
	inssRes, err := s.inssClient.FetchINSS(ctx, req)
	if err != nil {
		l.Warn("falha ao obter INSS, procedendo sem dedução previdenciária", "error", err)
	} else {
		// Adiciona o INSS calculado como dedução
		details = append(details, models.DeductionDetail{
			Type:   "deduction_social_security",
			Amount: inssRes.TaxAmount.Round(2),
		})
	}
	// -----------------------

	pension := inputs["pension_amount"]
	if pct, ok := inputs["pension_percentage"]; ok && pct.GreaterThan(decimal.Zero) {
		pension = req.GrossIncome.Mul(pct.Div(decimal.NewFromInt(100)))
	}
	if pension.GreaterThan(decimal.Zero) {
		details = append(details, models.DeductionDetail{Type: "pension_official", Amount: pension.Round(2)})
	}

	depQty := inputs["dependents_qty"]
	depKey := "dependent_deduction_monthly"
	if req.CalculationType == "annual" {
		depKey = "dependent_deduction_annual"
	}
	depRate := configs[depKey]
	if depRate.IsZero() {
		depRate = decimal.NewFromFloat(189.59)
	}

	depTotal := depQty.Mul(depRate)
	details = append(details, models.DeductionDetail{Type: "deduction_for_dependents", Amount: depTotal.Round(2)})

	eduSpent := inputs["education_expenses"]
	if eduSpent.GreaterThan(decimal.Zero) {
		eduLimitKey := "education_limit_monthly"
		if req.CalculationType == "annual" {
			eduLimitKey = "education_limit_annual"
		}
		eduLimit := configs[eduLimitKey]
		totalEduLimit := eduLimit.Mul(decimal.NewFromInt(int64(1 + int(depQty.IntPart()))))
		eduDeduction := decimal.Min(eduSpent, totalEduLimit)
		details = append(details, models.DeductionDetail{Type: "deduction_for_education", Amount: eduDeduction.Round(2)})
	}

	healthSpent := inputs["health_expenses"]
	if healthSpent.GreaterThan(decimal.Zero) {
		details = append(details, models.DeductionDetail{Type: "deduction_for_health", Amount: healthSpent.Round(2)})
	}

	pgblSpent := inputs["pgbl_contribution"]
	if pgblSpent.GreaterThan(decimal.Zero) {
		if unit, ok := s.getUnit(req.Inputs, "pgbl_contribution"); ok && unit == "percentage" {
			pgblSpent = req.GrossIncome.Mul(pgblSpent.Div(decimal.NewFromInt(100)))
		}
		pgblLimitPct := configs["pgbl_limit_percentage"]
		if pgblLimitPct.IsZero() {
			pgblLimitPct = decimal.NewFromInt(12)
		}
		pgblLimit := req.GrossIncome.Mul(pgblLimitPct.Div(decimal.NewFromInt(100)))
		pgblDeduction := decimal.Min(pgblSpent, pgblLimit)
		details = append(details, models.DeductionDetail{Type: "deduction_for_pgbl", Amount: pgblDeduction.Round(2)})
	}

	for _, d := range details {
		totalDeduction = totalDeduction.Add(d.Amount)
	}
	baseValue := req.GrossIncome.Sub(totalDeduction)

	return s.runTaxMath(ctx, req, baseValue, totalDeduction, details, configs)
}

func (s *CalculationService) calculateSimplificada(ctx context.Context, req models.UniversalTaxRequest, configs map[string]decimal.Decimal) (models.TaxResponse, error) {
	discount := req.GrossIncome.Mul(decimal.NewFromFloat(0.20))
	limitKey := "simplified_discount_monthly_limit"
	if req.CalculationType == "annual" {
		limitKey = "simplified_discount_annual_limit"
	}
	limit := configs[limitKey]
	if limit.IsZero() {
		limit = decimal.NewFromFloat(564.80)
	}

	if discount.GreaterThan(limit) {
		discount = limit
	}

	details := []models.DeductionDetail{{Type: "simplified_discount", Amount: discount.Round(2)}}
	return s.runTaxMath(ctx, req, req.GrossIncome.Sub(discount), discount, details, configs)
}

func (s *CalculationService) runTaxMath(ctx context.Context, req models.UniversalTaxRequest, baseValue, totalDeduction decimal.Decimal, details []models.DeductionDetail, configs map[string]decimal.Decimal) (models.TaxResponse, error) {
	l := s.getLogger(ctx)
	if baseValue.IsNegative() {
		baseValue = decimal.Zero
	}

	rule, err := s.repo.GetApplicableRule(ctx, req.TaxCode, baseValue, req.ReferenceDate)
	if err != nil {
		l.Error("regra de imposto nao encontrada", "base_value", baseValue)
		return models.TaxResponse{}, err
	}

	// 1. Prepara Metadata para o Response
	appliedRuleStr := fmt.Sprintf("Faixa: %s a %v | Alíquota: %s%% | Parcela a Deduzir: %s",
		rule.RangeMin, rule.RangeMax, rule.AliqPercent, rule.DeductionVal)

	usedConfigsMap := make(map[string]string)
	for k, v := range configs {
		usedConfigsMap[k] = v.String()
	}

	// Logs internos (slog)
	l.Debug("regra aplicada", "aliq", rule.AliqPercent, "deduction", rule.DeductionVal)

	// Cálculo do imposto
	aliqFactor := rule.AliqPercent.Div(decimal.NewFromInt(100))
	taxAmount := baseValue.Mul(aliqFactor).Sub(rule.DeductionVal)
	if taxAmount.IsNegative() {
		taxAmount = decimal.Zero
	}

	// Lógica de transição 2026
	date2026 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	if req.TaxCode == "IRPF" && !req.ReferenceDate.Before(date2026) {
		floor := configs["transition_2026_floor"]
		ceiling := configs["transition_2026_ceiling"]

		if baseValue.LessThanOrEqual(floor) && taxAmount.GreaterThan(decimal.Zero) {
			reduction := taxAmount
			taxAmount = decimal.Zero
			details = append(details, models.DeductionDetail{Type: "reforma_2026_isencao_total", Amount: reduction.Round(2).Neg()})
		} else if baseValue.GreaterThan(floor) && baseValue.LessThanOrEqual(ceiling) {
			fA := configs["transition_2026_factor_a"]
			fB := configs["transition_2026_factor_b"]
			reduction := fA.Sub(fB.Mul(baseValue))
			if reduction.GreaterThan(decimal.Zero) {
				if reduction.GreaterThan(taxAmount) {
					reduction = taxAmount
				}
				taxAmount = taxAmount.Sub(reduction)
				details = append(details, models.DeductionDetail{Type: "reforma_2026_reducao_adicional", Amount: reduction.Round(2).Neg()})
			}
		}
	}

	effRate := decimal.Zero
	if req.GrossIncome.GreaterThan(decimal.Zero) {
		effRate = taxAmount.Div(req.GrossIncome).Mul(decimal.NewFromInt(100))
	}

	return models.TaxResponse{
		ReferenceDate:        req.ReferenceDate.Format("2006-01-02"),
		GrossIncome:          req.GrossIncome.Round(2),
		TotalDeductionAmount: totalDeduction.Round(2),
		BaseValue:            baseValue.Round(2),
		TaxAmount:            taxAmount.Round(2),
		EffectiveRate:        effRate.Round(2),
		DeductionDetails:     details,
		AppliedRule:          appliedRuleStr,
		UsedConfigs:          usedConfigsMap,
	}, nil
}

func (s *CalculationService) getUnit(inputs []models.DocumentoFiscalRequest, key string) (string, bool) {
	for _, in := range inputs {
		if in.Key == key {
			return in.Unit, true
		}
	}
	return "", false
}
