// path: backend/go/fiber/microservices/ms-tax-individual-social-security/services/inss_service.go
package services

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"time"
	"bytes"
	"encoding/json"
	"net/http"

	"taxnexus-core-lib/models"
	"taxnexus-core-lib/repository"

	"github.com/shopspring/decimal"
)

type INSSClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

type INSSService struct {
	repo   *repository.TaxRepository
	logger *slog.Logger
}

func NewINSSClient(baseURL string) *INSSClient {
	return &INSSClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 3 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				IdleConnTimeout:     90 * time.Second,
				MaxIdleConnsPerHost: 20, // Crucial para performance entre microserviços
			},
		},
	}
}

func NewINSSService(repo *repository.TaxRepository) *INSSService {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	return &INSSService{
		repo:   repo,
		logger: slog.New(handler),
	}
}

func (s *INSSService) getLogger(ctx context.Context) *slog.Logger {
	if tid, ok := ctx.Value("requestid").(string); ok {
		return s.logger.With("trace_id", tid)
	}
	return s.logger
}

func (s *INSSService) Calculate(ctx context.Context, req models.UniversalTaxRequest) (models.TaxResponse, error) {
	l := s.getLogger(ctx)
	
	if req.ReferenceDate.IsZero() {
		req.ReferenceDate = time.Now()
	}

	// 1. Busca configurações (ex: teto do INSS)
	configs, err := s.repo.GetTableConfigs(ctx, "INSS", req.ReferenceDate)
	if err != nil {
		l.Error("falha ao carregar configs do INSS", "error", err)
		return models.TaxResponse{}, err
	}

	// 2. Busca todas as fatias/regras do INSS para a data
	// Nota: Precisamos de todas as faixas, não apenas a que o salário se encaixa
	rules, err := s.repo.GetTaxRulesForPeriod(ctx, "INSS", req.ReferenceDate)
	if err != nil {
		l.Error("falha ao carregar fatias do INSS", "error", err)
		return models.TaxResponse{}, err
	}

	// Garante que as faixas estão ordenadas pelo valor mínimo
	sort.Slice(rules, func(i, j int) bool {
		return rules[i].RangeMin.LessThan(rules[j].RangeMin)
	})

	l.Info("calculo INSS iniciado", "gross_income", req.GrossIncome)

	// 3. Lógica de Teto
	ceiling := configs["inss_contribution_ceiling"]
	calculationBase := req.GrossIncome
	if !ceiling.IsZero() && calculationBase.GreaterThan(ceiling) {
		calculationBase = ceiling
		l.Debug("teto do INSS aplicado", "ceiling", ceiling)
	}

	totalTax := decimal.Zero
	var details []models.DeductionDetail
	lastRangeMax := decimal.Zero

	// 4. Cálculo Progressivo por Fatias
	for _, rule := range rules {
		if calculationBase.LessThanOrEqual(lastRangeMax) {
			break
		}

		// Determina o topo da fatia atual
		currentUpperLimit := calculationBase
		if rule.RangeMax != nil && calculationBase.GreaterThan(*rule.RangeMax) {
			currentUpperLimit = *rule.RangeMax
		}

		// Calcula a diferença tributável nesta faixa
		taxableInThisRange := currentUpperLimit.Sub(lastRangeMax)
		aliqFactor := rule.AliqPercent.Div(decimal.NewFromInt(100))
		taxInRange := taxableInThisRange.Mul(aliqFactor)

		totalTax = totalTax.Add(taxInRange)
		
		details = append(details, models.DeductionDetail{
			Type:   fmt.Sprintf("inss_range_%s_percent", rule.AliqPercent.String()),
			Amount: taxInRange.Round(2),
		})

		if rule.RangeMax == nil {
			break
		}
		lastRangeMax = *rule.RangeMax
	}

	effRate := decimal.Zero
	if req.GrossIncome.GreaterThan(decimal.Zero) {
		effRate = totalTax.Div(req.GrossIncome).Mul(decimal.NewFromInt(100))
	}

	return models.TaxResponse{
		ReferenceDate:    req.ReferenceDate.Format("2006-01-02"),
		GrossIncome:      req.GrossIncome.Round(2),
		BaseValue:        calculationBase.Round(2),
		TaxAmount:        totalTax.Round(2),
		EffectiveRate:    effRate.Round(2),
		DeductionDetails: details,
		UsedConfigs:      mapDecimalToString(configs),
	}, nil
}

func mapDecimalToString(m map[string]decimal.Decimal) map[string]string {
	res := make(map[string]string)
	for k, v := range m {
		res[k] = v.String()
	}
	return res
}

func (c *INSSClient) FetchINSS(ctx context.Context, req models.UniversalTaxRequest) (models.TaxResponse, error) {
	body, _ := json.Marshal(req)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.BaseURL+"/api/v1/calculate/inss", bytes.NewBuffer(body))
	if err != nil {
		return models.TaxResponse{}, err
	}
	
	httpReq.Header.Set("Content-Type", "application/json")
	if tid, ok := ctx.Value("requestid").(string); ok {
		httpReq.Header.Set("X-Request-ID", tid)
	}

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return models.TaxResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.TaxResponse{}, fmt.Errorf("inss service status: %d", resp.StatusCode)
	}

	var taxRes models.TaxResponse
	err = json.NewDecoder(resp.Body).Decode(&taxRes)
	return taxRes, err
}