// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/credit/engine_test.go
package credit

import (
	"context"
	"testing"
	"time"

	"ms-billing-engine-tax-rates/internal/phase"
	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"

	"github.com/shopspring/decimal"
)

type mockEngine struct {
	output models.DocumentoFiscalSaida
}

func (m *mockEngine) ProcessWithPhase(ctx context.Context, input models.DocumentoFiscalEntrada, filter phase.CalculatorFilter) (models.DocumentoFiscalSaida, error) {
	return m.output, nil
}

// mockRepo retorna nil (sem restrição — permite crédito amplo)
type mockRepo struct{}

func (r *mockRepo) GetIvaDualRule(ctx context.Context, ncm, uf, ibge string) (*repository.IvaDualRule, error) {
	return &repository.IvaDualRule{PercentualReducao: decimal.Zero}, nil
}
func (r *mockRepo) GetRateByNCM(ctx context.Context, ncm string, uf string) (float64, error)    { return 0, nil }
func (r *mockRepo) GetIBSRate(ctx context.Context, municipioIBGE string) (float64, error)        { return 0, nil }
func (r *mockRepo) GetFederalTaxRule(ctx context.Context, rt, cstP, cstC string) (*repository.FederalTaxRule, error) { return nil, nil }
func (r *mockRepo) GetICMSRule(ctx context.Context, o, d string) (*repository.ICMSRule, error)   { return nil, nil }
func (r *mockRepo) GetEquivalence(ctx context.Context, csosn, top string) (*repository.TaxEquivalence, error) { return nil, nil }
func (r *mockRepo) GetSimplesFaixa(ctx context.Context, a string, rbt12 decimal.Decimal) (*repository.SimplesFaixa, error) { return nil, nil }
func (r *mockRepo) GetProductException(ctx context.Context, nf, ng, ud, rtd string) (*repository.ProductException, error) { return nil, nil }
func (r *mockRepo) GetIPIRegra(ctx context.Context, n, e, c, t, p, u string, z bool, d string) (*repository.IPIRegra, error) { return nil, nil }
func (r *mockRepo) GetNCMSeletivo(ctx context.Context, ncm string) (*repository.NCMSeletivoRule, error) { return nil, nil }

func input(ncm string, valor float64) models.DocumentoFiscalEntrada {
	return models.DocumentoFiscalEntrada{
		DocumentoID: "NFE-001", TipoOperacaoFiscal: "ENTRADA",
		CRTEmitente: "3", DataOperacao: time.Now(),
		LocalizacaoOrigem:  models.LocalizacaoFiscal{UF: "SP"},
		LocalizacaoDestino: models.LocalizacaoFiscal{UF: "SP"},
		Itens: []models.ItemDocumentoFiscalEntrada{
			{SKU: "P1", NCM: ncm, Quantidade: decimal.NewFromInt(1), ValorUnitario: decimal.NewFromFloat(valor)},
		},
	}
}

// ─── TST-005.01: Fornecedor qualificado → crédito integral ─────────────────

func TestCredit_FornecedorQualificado(t *testing.T) {
	engine := &mockEngine{output: models.DocumentoFiscalSaida{
		Itens: []models.ItemDocumentoFiscalSaida{{
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "CBS", Valor: decimal.NewFromFloat(90)},
				{Tributo: "IBS", Valor: decimal.NewFromFloat(75)},
				{Tributo: "IS", Valor: decimal.NewFromFloat(50)},
			},
		}},
	}}
	ce := NewCreditEngine(engine, &AlwaysQualifiedSupplierChecker{}, &mockRepo{})

	resp, err := ce.Calculate(context.Background(), CreditCalculationRequest{
		DocumentoFiscalEntrada: input("84713012", 1000),
		CNPJFornecedor:         "12345678000199",
	})
	if err != nil {
		t.Fatalf("erro: %v", err)
	}

	if resp.CreditoCBS.String() != "90" {
		t.Errorf("CreditoCBS = %s, esperado 90", resp.CreditoCBS)
	}
	if resp.CreditoIBS.String() != "75" {
		t.Errorf("CreditoIBS = %s, esperado 75", resp.CreditoIBS)
	}
	if resp.CreditoTotal.String() != "165" {
		t.Errorf("CreditoTotal = %s, esperado 165", resp.CreditoTotal)
	}
	// IS não gera crédito (TST-005.05)
	if resp.MotivoBloqueio != "" {
		t.Errorf("MotivoBloqueio = %q, esperado vazio", resp.MotivoBloqueio)
	}
}

// ─── TST-005.02: Fornecedor não qualificado → crédito = 0 ─────────────────

func TestCredit_FornecedorBloqueado(t *testing.T) {
	ce := NewCreditEngine(&mockEngine{}, &BlockedSupplierChecker{}, &mockRepo{})

	resp, err := ce.Calculate(context.Background(), CreditCalculationRequest{
		DocumentoFiscalEntrada: input("84713012", 1000),
		CNPJFornecedor:         "00000000000000",
	})
	if err != nil {
		t.Fatalf("erro: %v", err)
	}

	if !resp.CreditoTotal.IsZero() {
		t.Errorf("CreditoTotal = %s, esperado 0", resp.CreditoTotal)
	}
	if resp.MotivoBloqueio == "" {
		t.Error("MotivoBloqueio deve estar preenchido")
	}
}

// ─── TST-005.03: Fornecedor Simples Nacional → crédito restrito ────────────

func TestCredit_SimplesNacional(t *testing.T) {
	simplesChecker := &BlockedSupplierChecker{} // Simples Nacional = bloqueado
	ce := NewCreditEngine(&mockEngine{}, simplesChecker, &mockRepo{})

	resp, _ := ce.Calculate(context.Background(), CreditCalculationRequest{
		DocumentoFiscalEntrada: input("84713012", 1000),
		CNPJFornecedor:         "11111111000191",
	})

	if !resp.CreditoTotal.IsZero() {
		t.Errorf("Simples Nacional não deve gerar crédito: %s", resp.CreditoTotal)
	}
}

// ─── TST-005.05: IS não gera crédito ───────────────────────────────────────

func TestCredit_ISNaoGeraCredito(t *testing.T) {
	engine := &mockEngine{output: models.DocumentoFiscalSaida{
		Itens: []models.ItemDocumentoFiscalSaida{{
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "CBS", Valor: decimal.NewFromFloat(90)},
				{Tributo: "IBS", Valor: decimal.NewFromFloat(75)},
				{Tributo: "IS", Valor: decimal.NewFromFloat(200)}, // IS alto (bebida)
			},
		}},
	}}
	ce := NewCreditEngine(engine, &AlwaysQualifiedSupplierChecker{}, &mockRepo{})

	resp, _ := ce.Calculate(context.Background(), CreditCalculationRequest{
		DocumentoFiscalEntrada: input("22030000", 1000), CNPJFornecedor: "12345678000199",
	})

	// IS não entra no crédito total
	if resp.CreditoTotal.String() != "165" {
		t.Errorf("CreditoTotal = %s, esperado 165 (sem IS)", resp.CreditoTotal)
	}
}

// ─── TST-005.06: Valor zero → crédito = 0 ──────────────────────────────────

func TestCredit_ValorZero(t *testing.T) {
	engine := &mockEngine{output: models.DocumentoFiscalSaida{
		Itens: []models.ItemDocumentoFiscalSaida{{Tributos: nil}},
	}}
	ce := NewCreditEngine(engine, &AlwaysQualifiedSupplierChecker{}, &mockRepo{})

	resp, _ := ce.Calculate(context.Background(), CreditCalculationRequest{
		DocumentoFiscalEntrada: input("84713012", 0), CNPJFornecedor: "12345678000199",
	})

	if !resp.CreditoTotal.IsZero() {
		t.Errorf("CreditoTotal = %s, esperado 0", resp.CreditoTotal)
	}
}

// ─── TST-005.07: GET /v1/credit/summary ────────────────────────────────────

func TestCredit_Summary(t *testing.T) {
	ce := NewCreditEngine(&mockEngine{}, &AlwaysQualifiedSupplierChecker{}, &mockRepo{})

	resp, err := ce.Summary(context.Background(), "2026-06")
	if err != nil {
		t.Fatalf("erro: %v", err)
	}
	if resp.Periodo != "2026-06" {
		t.Errorf("Periodo = %q, esperado 2026-06", resp.Periodo)
	}
}

// ─── TST-005.08: Período sem dados → totais zerados ─────────────────────────

func TestCredit_SummaryVazio(t *testing.T) {
	ce := NewCreditEngine(&mockEngine{}, &AlwaysQualifiedSupplierChecker{}, &mockRepo{})

	resp, _ := ce.Summary(context.Background(), "2025-01")

	if !resp.CreditoDisponivel.IsZero() {
		t.Errorf("CreditoDisponivel = %s, esperado 0", resp.CreditoDisponivel)
	}
}

// ─── TST-005.10: Múltiplos itens → crédito totalizado ──────────────────────

func TestCredit_MultiplosItens(t *testing.T) {
	engine := &mockEngine{output: models.DocumentoFiscalSaida{
		Itens: []models.ItemDocumentoFiscalSaida{
			{SKU: "A", Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "CBS", Valor: decimal.NewFromFloat(45)},
				{Tributo: "IBS", Valor: decimal.NewFromFloat(30)},
			}},
			{SKU: "B", Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "CBS", Valor: decimal.NewFromFloat(90)},
				{Tributo: "IBS", Valor: decimal.NewFromFloat(75)},
			}},
		},
	}}
	ce := NewCreditEngine(engine, &AlwaysQualifiedSupplierChecker{}, &mockRepo{})

	input2 := models.DocumentoFiscalEntrada{
		DocumentoID: "NFE-002", TipoOperacaoFiscal: "ENTRADA",
		CRTEmitente: "3", DataOperacao: time.Now(),
		LocalizacaoOrigem:  models.LocalizacaoFiscal{UF: "SP"},
		LocalizacaoDestino: models.LocalizacaoFiscal{UF: "SP"},
		Itens: []models.ItemDocumentoFiscalEntrada{
			{SKU: "A", NCM: "84713012", Quantidade: decimal.NewFromInt(1), ValorUnitario: decimal.NewFromFloat(500)},
			{SKU: "B", NCM: "84713012", Quantidade: decimal.NewFromInt(1), ValorUnitario: decimal.NewFromFloat(1000)},
		},
	}

	resp, _ := ce.Calculate(context.Background(), CreditCalculationRequest{
		DocumentoFiscalEntrada: input2, CNPJFornecedor: "12345678000199",
	})

	// Total CBS = 45 + 90 = 135, IBS = 30 + 75 = 105, Total = 240
	if resp.CreditoTotal.String() != "240" {
		t.Errorf("CreditoTotal = %s, esperado 240", resp.CreditoTotal)
	}
	if len(resp.Itens) != 2 {
		t.Errorf("esperado 2 itens, obtido %d", len(resp.Itens))
	}
}
