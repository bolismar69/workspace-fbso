// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/credit/engine.go
package credit

import (
	"context"
	"fmt"
	"log/slog"

	"ms-billing-engine-tax-rates/internal/phase"
	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TaxEngine é a interface de cálculo usada pelo CreditEngine.
type TaxEngine interface {
	ProcessWithPhase(ctx context.Context, input models.DocumentoFiscalEntrada, filter phase.CalculatorFilter) (models.DocumentoFiscalSaida, error)
}

// CreditEngine calcula créditos de CBS/IBS de documentos fiscais de entrada
// (compras de fornecedores) conforme regras do Lucro Real (BR-08).
type CreditEngine struct {
	engine         TaxEngine
	supplierCheck  SupplierChecker
	repo           repository.TaxRepository
}

// NewCreditEngine cria um CreditEngine.
func NewCreditEngine(engine TaxEngine, supplier SupplierChecker, repo repository.TaxRepository) *CreditEngine {
	return &CreditEngine{
		engine:        engine,
		supplierCheck: supplier,
		repo:          repo,
	}
}

// Calculate calcula os créditos de CBS/IBS apropriáveis de uma NF-e de entrada.
//
// Fluxo:
//  1. Verifica qualificação do fornecedor (via SupplierChecker)
//  2. Se não qualificado → crédito = 0 com motivo_bloqueio
//  3. Calcula CBS/IBS via motor (ProcessWithPhase)
//  4. Para cada item, verifica permite_credito_amplo na iva_dual_rules
//  5. IS não gera crédito (não é recuperável)
//  6. Crédito não pode exceder o valor do item
func (ce *CreditEngine) Calculate(ctx context.Context, req CreditCalculationRequest) (*CreditCalculationResponse, error) {
	// 1. Verifica fornecedor
	supplier, err := ce.supplierCheck.Check(ctx, req.CNPJFornecedor)
	if err != nil {
		return nil, fmt.Errorf("falha ao verificar fornecedor: %w", err)
	}

	if !supplier.PermiteCredito {
		slog.Warn("Fornecedor não qualificado para crédito",
			"cnpj", req.CNPJFornecedor,
			"status", supplier.Status,
			"regime", supplier.RegimeTributario,
		)
		return &CreditCalculationResponse{
			IDTransaction:  uuid.NewString(),
			CNPJFornecedor: req.CNPJFornecedor,
			MotivoBloqueio: fmt.Sprintf("fornecedor_nao_qualificado: %s (%s)", supplier.Status, supplier.RegimeTributario),
			Itens:          make([]CreditItem, len(req.Itens)),
		}, nil
	}

	// 2. Calcula tributos de entrada via motor
	defaultFilter := phase.CalculatorFilter{
		Phase:           "PRE_REFORMA",
		CBSActive:       true,
		IBSActive:       true,
		ISActive:        true,
		IPIActive:       false,
		PISCOFINSActive: false,
		ICMSActive:      false,
		ISSActive:       false,
	}

	output, err := ce.engine.ProcessWithPhase(ctx, req.DocumentoFiscalEntrada, defaultFilter)
	if err != nil {
		return nil, fmt.Errorf("falha ao calcular tributos de entrada: %w", err)
	}

	// 3. Calcula créditos por item
	resp := &CreditCalculationResponse{
		IDTransaction:  uuid.NewString(),
		CNPJFornecedor: req.CNPJFornecedor,
		Itens:          make([]CreditItem, len(req.Itens)),
	}

	creditoCBS := decimal.Zero
	creditoIBS := decimal.Zero

	for i, item := range output.Itens {
		ci := CreditItem{SKU: item.SKU, Creditavel: true}
		valorItem := req.Itens[i].ValorUnitario.Mul(req.Itens[i].Quantidade)

		for _, t := range item.Tributos {
			switch t.Tributo {
			case "CBS":
				ci.CBSDestacado = t.Valor
				// Verifica permite_credito_amplo
				if !ce.permiteCreditoAmplo(ctx, req.Itens[i].NCM, req.LocalizacaoDestino.UF) {
					ci.Creditavel = false
					resp.MotivoBloqueio = "credito_amplo_nao_permitido"
				} else {
					creditoCBS = creditoCBS.Add(t.Valor)
				}
			case "IBS":
				ci.IBSDestacado = t.Valor
				if ci.Creditavel {
					creditoIBS = creditoIBS.Add(t.Valor)
				}
			}
		}

		// Piso: crédito não pode exceder valor do item
		if creditoCBS.Add(creditoIBS).GreaterThan(valorItem) {
			resp.MotivoBloqueio = "credito_excede_valor_item"
			ci.Creditavel = false
		}

		resp.Itens[i] = ci
	}

	resp.CreditoCBS = creditoCBS.Round(2)
	resp.CreditoIBS = creditoIBS.Round(2)
	resp.CreditoTotal = creditoCBS.Add(creditoIBS).Round(2)

	slog.Info("Créditos calculados",
		"cnpj", req.CNPJFornecedor,
		"credito_cbs", resp.CreditoCBS,
		"credito_ibs", resp.CreditoIBS,
		"credito_total", resp.CreditoTotal,
	)

	return resp, nil
}

// Summary retorna o resumo de créditos disponíveis vs apropriados no período.
// Na versão atual, retorna valores zerados como stub (GAP-005, T-005.6).
func (ce *CreditEngine) Summary(ctx context.Context, periodo string) (*CreditSummaryResponse, error) {
	slog.Info("Credit summary consultado (stub — persistência pendente GAP-007)",
		"periodo", periodo,
	)
	return &CreditSummaryResponse{
		Periodo: periodo,
	}, nil
}

// permiteCreditoAmplo verifica na iva_dual_rules se o NCM+UF permite crédito amplo.
func (ce *CreditEngine) permiteCreditoAmplo(ctx context.Context, ncm, uf string) bool {
	rule, err := ce.repo.GetIvaDualRule(ctx, ncm, uf, "")
	if err != nil || rule == nil {
		return true // default: permite se não encontrou restrição
	}
	return rule.PercentualReducao.LessThan(decimal.NewFromInt(100))
}
