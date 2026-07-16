// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/simulation/service_test.go
package simulation

import (
	"context"
	"testing"

	"ms-billing-engine-tax-rates/internal/phase"
	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

// mockEngine implementa TaxEngine para testes.
type mockEngine struct {
	output     models.DocumentoFiscalSaida
	err        error
	callCount  int
}

func (m *mockEngine) ProcessWithPhase(ctx context.Context, input models.DocumentoFiscalEntrada, filter phase.CalculatorFilter) (models.DocumentoFiscalSaida, error) {
	m.callCount++
	if m.err != nil {
		return models.DocumentoFiscalSaida{}, m.err
	}
	return m.output, nil
}

func dec(v string) decimal.Decimal { return decimal.RequireFromString(v) }

// ─── TST-003.01: Simulação com 1 destino — margem líquida positiva ────────

func TestSimulation_MargemPositiva(t *testing.T) {
	engine := &mockEngine{
		output: models.DocumentoFiscalSaida{
			TotalImpostos: dec("246.50"),
			Itens: []models.ItemDocumentoFiscalSaida{
				{ValorLiquido: dec("753.50")},
			},
		},
	}

	svc := NewSimulationService(engine)
	resp, err := svc.Simulate(context.Background(), SimulationRequest{
		Itens: []SimulationItem{
			{SKU: "PROD-001", NCM: "84713012", Quantidade: dec("2"), ValorUnitario: dec("500"), Custo: dec("300")},
		},
		UFOrigem: "SP",
		Destinos: []Destino{{UF: "RJ"}},
	})
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if len(resp.ImpactoPorUF) != 1 {
		t.Fatalf("esperado 1 impacto, obtido %d", len(resp.ImpactoPorUF))
	}

	imp := resp.ImpactoPorUF[0]
	if imp.UF != "RJ" {
		t.Errorf("UF = %q, esperado RJ", imp.UF)
	}
	// Margem = (753.50 - 300) / 753.50 * 100 ≈ 60.19%
	if imp.MargemLiquida.LessThan(dec("60")) || imp.MargemLiquida.GreaterThan(dec("61")) {
		t.Errorf("MargemLiquida = %s, esperado ~60.19", imp.MargemLiquida)
	}
	if imp.AlertaMargem {
		t.Error("AlertaMargem não deveria estar ativo para margem positiva")
	}
}

// ─── TST-003.02: Simulação com 3 destinos — array com impacto por UF ────────

func TestSimulation_MultiplosDestinos(t *testing.T) {
	engine := &mockEngine{
		output: models.DocumentoFiscalSaida{
			TotalImpostos: dec("180.00"),
			Itens: []models.ItemDocumentoFiscalSaida{
				{ValorLiquido: dec("820.00")},
			},
		},
	}

	svc := NewSimulationService(engine)
	resp, err := svc.Simulate(context.Background(), SimulationRequest{
		Itens: []SimulationItem{
			{SKU: "PROD-001", NCM: "84713012", Quantidade: dec("1"), ValorUnitario: dec("1000"), Custo: dec("600")},
		},
		UFOrigem: "SP",
		Destinos: []Destino{
			{UF: "SP"},
			{UF: "RJ"},
			{UF: "MG"},
		},
	})
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(resp.ImpactoPorUF) != 3 {
		t.Fatalf("esperado 3 impactos, obtido %d", len(resp.ImpactoPorUF))
	}
	if engine.callCount != 3 {
		t.Errorf("engine chamado %d vezes, esperado 3", engine.callCount)
	}

	ufs := make([]string, len(resp.ImpactoPorUF))
	for i, imp := range resp.ImpactoPorUF {
		ufs[i] = imp.UF
	}
	t.Logf("Destinos simulados: %v", ufs)
}

// ─── TST-003.03: Margem negativa — retorna alerta_margem ────────────────────

func TestSimulation_MargemNegativa(t *testing.T) {
	engine := &mockEngine{
		output: models.DocumentoFiscalSaida{
			TotalImpostos: dec("400.00"),
			Itens: []models.ItemDocumentoFiscalSaida{
				{ValorLiquido: dec("600.00")}, // impostos > lucro bruto
			},
		},
	}

	svc := NewSimulationService(engine)
	resp, err := svc.Simulate(context.Background(), SimulationRequest{
		Itens: []SimulationItem{
			{SKU: "PROD-001", NCM: "84713012", Quantidade: dec("1"), ValorUnitario: dec("1000"), Custo: dec("700")},
		},
		UFOrigem: "SP",
		Destinos: []Destino{{UF: "RJ"}},
	})
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	imp := resp.ImpactoPorUF[0]
	// Margem = (600 - 700) / 600 * 100 ≈ -16.67%
	if !imp.MargemLiquida.LessThan(decimal.Zero) {
		t.Errorf("MargemLiquida = %s, esperado negativo", imp.MargemLiquida)
	}
	if !imp.AlertaMargem {
		t.Error("AlertaMargem deveria estar ativo para margem negativa")
	}
}

// ─── TST-003.04: Custo zero — margem = 100% ────────────────────────────────

func TestSimulation_CustoZero(t *testing.T) {
	engine := &mockEngine{
		output: models.DocumentoFiscalSaida{
			TotalImpostos: dec("100.00"),
			Itens: []models.ItemDocumentoFiscalSaida{
				{ValorLiquido: dec("900.00")},
			},
		},
	}

	svc := NewSimulationService(engine)
	resp, err := svc.Simulate(context.Background(), SimulationRequest{
		Itens: []SimulationItem{
			{SKU: "PROD-001", NCM: "84713012", Quantidade: dec("1"), ValorUnitario: dec("1000"), Custo: dec("0")},
		},
		UFOrigem: "SP",
		Destinos: []Destino{{UF: "RJ"}},
	})
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	imp := resp.ImpactoPorUF[0]
	// Margem = (900 - 0) / 900 * 100 = 100%
	if !imp.MargemLiquida.Equal(dec("100")) {
		t.Errorf("MargemLiquida = %s, esperado 100", imp.MargemLiquida)
	}
}

// ─── TST-003.05: Destino sem IBS — warning no impacto ──────────────────────

func TestSimulation_SemIBSNoDestino(t *testing.T) {
	engine := &mockEngine{
		output: models.DocumentoFiscalSaida{
			TotalImpostos: dec("0"),
			Itens: []models.ItemDocumentoFiscalSaida{
				{ValorLiquido: dec("1000.00")},
			},
		},
	}

	svc := NewSimulationService(engine)
	resp, err := svc.Simulate(context.Background(), SimulationRequest{
		Itens: []SimulationItem{
			{SKU: "PROD-001", NCM: "84713012", Quantidade: dec("1"), ValorUnitario: dec("1000"), Custo: dec("500")},
		},
		UFOrigem: "SP",
		Destinos: []Destino{{UF: "XX", MunicipioIBGE: "9999999"}},
	})
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	imp := resp.ImpactoPorUF[0]
	if imp.Warning != "aliquota_ibs_nao_encontrada" {
		t.Errorf("Warning = %q, esperado 'aliquota_ibs_nao_encontrada'", imp.Warning)
	}
}

// ─── TST-003.06: Produto sujeito a IS — IS incluso no impacto ──────────────

func TestSimulation_ComImpostoSeletivo(t *testing.T) {
	engine := &mockEngine{
		output: models.DocumentoFiscalSaida{
			TotalImpostos: dec("350.00"), // inclui IS
			Itens: []models.ItemDocumentoFiscalSaida{
				{
					ValorLiquido: dec("650.00"),
					Tributos: []models.TributosItemDocumentoFiscalSaida{
						{Tributo: "IS", Valor: dec("50")},
						{Tributo: "CBS", Valor: dec("90")},
						{Tributo: "IBS", Valor: dec("75")},
						{Tributo: "ICMS", Valor: dec("120")},
						{Tributo: "PIS", Valor: dec("10.725")},
						{Tributo: "COFINS", Valor: dec("4.275")},
					},
				},
			},
		},
	}

	svc := NewSimulationService(engine)
	resp, err := svc.Simulate(context.Background(), SimulationRequest{
		Itens: []SimulationItem{
			{SKU: "BEBIDA", NCM: "22030000", Quantidade: dec("1"), ValorUnitario: dec("1000"), Custo: dec("400")},
		},
		UFOrigem: "SP",
		Destinos: []Destino{{UF: "RJ"}},
	})
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	imp := resp.ImpactoPorUF[0]
	if imp.TotalImpostos.String() != "350" {
		t.Errorf("TotalImpostos = %s, esperado 350", imp.TotalImpostos)
	}
	if imp.MargemLiquida.LessThan(decimal.Zero) {
		t.Error("margem negativa inesperada")
	}
}

// ─── Testes de validação ────────────────────────────────────────────────────

func TestSimulation_ValidacaoSemItens(t *testing.T) {
	svc := NewSimulationService(&mockEngine{})
	_, err := svc.Simulate(context.Background(), SimulationRequest{
		Destinos: []Destino{{UF: "SP"}},
	})
	if err == nil {
		t.Fatal("esperado erro para simulação sem itens")
	}
}

func TestSimulation_ValidacaoSemDestinos(t *testing.T) {
	svc := NewSimulationService(&mockEngine{})
	_, err := svc.Simulate(context.Background(), SimulationRequest{
		Itens: []SimulationItem{
			{SKU: "P", NCM: "84713012", Quantidade: dec("1"), ValorUnitario: dec("100")},
		},
	})
	if err == nil {
		t.Fatal("esperado erro para simulação sem destinos")
	}
}
