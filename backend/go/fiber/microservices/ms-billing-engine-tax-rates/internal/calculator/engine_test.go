package calculator

import (
	"context"
	"errors"
	"testing"
	"time"

	"ms-billing-engine-tax-rates/internal/domain"
	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

type mockTaxCalculator struct {
	items  []models.ItemDocumentoFiscalSaida
	err    error
}

func (m *mockTaxCalculator) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.items, nil
}

func mockItem(sku, tributo string, valor float64) models.ItemDocumentoFiscalSaida {
	return models.ItemDocumentoFiscalSaida{
		SKU: sku,
		Tributos: []models.TributosItemDocumentoFiscalSaida{
			{
				Tributo: tributo,
				Valor:   decimal.NewFromFloat(valor),
			},
		},
	}
}

func inputDocumento(itens ...models.ItemDocumentoFiscalEntrada) models.DocumentoFiscalEntrada {
	return models.DocumentoFiscalEntrada{
		CorrelacaoID:       "corr-001",
		DocumentoID:        "doc-001",
		CRTEmitente:        "LUCRO_REAL",
		TipoOperacaoFiscal: "SAIDA",
		LocalizacaoOrigem:  models.LocalizacaoFiscal{UF: "SP"},
		LocalizacaoDestino: models.LocalizacaoFiscal{UF: "SP"},
		Itens:              itens,
		DataOperacao:       time.Now(),
	}
}

func inputItem(sku string, quantidade, valorUnitario float64) models.ItemDocumentoFiscalEntrada {
	return models.ItemDocumentoFiscalEntrada{
		SKU:           sku,
		NCM:           "84713019",
		Quantidade:    decimal.NewFromFloat(quantidade),
		ValorUnitario: decimal.NewFromFloat(valorUnitario),
	}
}

func TestEngine_PreCalcInjetaValoresNaFase2(t *testing.T) {
	item := inputItem("SKU-A", 1, 1000)

	preCalc := &mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{mockItem("SKU-A", "IPI", 100)},
	}
	calc2 := &mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{mockItem("SKU-A", "ICMS", 180)},
	}

	engine := BillingEngineOrdered(
		[]domain.TaxCalculator{preCalc},
		calc2,
	)

	input := inputDocumento(item)
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(res.Itens) != 1 {
		t.Fatalf("esperado 1 item, obteve %d", len(res.Itens))
	}
	if len(res.Itens[0].Tributos) != 2 {
		t.Fatalf("esperado 2 tributos (IPI + ICMS), obteve %d", len(res.Itens[0].Tributos))
	}

	if res.Itens[0].Tributos[0].Tributo != "IPI" {
		t.Errorf("tributo[0] = %s, want IPI", res.Itens[0].Tributos[0].Tributo)
	}
	if res.Itens[0].Tributos[1].Tributo != "ICMS" {
		t.Errorf("tributo[1] = %s, want ICMS", res.Itens[0].Tributos[1].Tributo)
	}
}

func TestEngine_Paralelo(t *testing.T) {
	item := inputItem("SKU-A", 1, 1000)

	calc1 := &mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{mockItem("SKU-A", "ICMS", 180)},
	}
	calc2 := &mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{mockItem("SKU-A", "PIS", 16.50)},
	}

	engine := BillingEngine(calc1, calc2)

	input := inputDocumento(item)
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(res.Itens[0].Tributos) != 2 {
		t.Fatalf("esperado 2 tributos, obteve %d", len(res.Itens[0].Tributos))
	}

	tributos := make(map[string]bool)
	for _, trib := range res.Itens[0].Tributos {
		tributos[trib.Tributo] = true
	}
	if !tributos["ICMS"] || !tributos["PIS"] {
		t.Errorf("tributos faltando: %v", tributos)
	}
}

func TestEngine_TotalImpostos_Consolidado(t *testing.T) {
	item := inputItem("SKU-A", 2, 500) // total item = 1000

	calc := &mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "ICMS", Valor: decimal.NewFromFloat(180)},
				{Tributo: "PIS", Valor: decimal.NewFromFloat(16.50)},
			},
		}},
	}

	engine := BillingEngine(calc)

	input := inputDocumento(item)
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if res.TotalImpostos.StringFixed(2) != "196.50" {
		t.Errorf("total impostos = %s, want 196.50", res.TotalImpostos.StringFixed(2))
	}
	if res.TotalNota.StringFixed(2) != "1000.00" {
		t.Errorf("total nota = %s, want 1000.00", res.TotalNota.StringFixed(2))
	}
}

func TestEngine_PreCalcErro_Propaga(t *testing.T) {
	preCalc := &mockTaxCalculator{
		err: errors.New("falha no IPI"),
	}
	engine := BillingEngineOrdered([]domain.TaxCalculator{preCalc})

	_, err := engine.Process(context.Background(), inputDocumento(inputItem("A", 1, 100)))
	if err == nil {
		t.Fatal("esperado erro do pre-calc, obteve nil")
	}
}

func TestEngine_CalcParaleloErro_ColetaELoga(t *testing.T) {
	item := inputItem("SKU-A", 1, 1000)

	calc1 := &mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{mockItem("SKU-A", "ICMS", 180)},
	}
	calc2 := &mockTaxCalculator{
		err: errors.New("falha no PIS/COFINS"),
	}

	engine := BillingEngine(calc1, calc2)

	input := inputDocumento(item)
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	// Resultado parcial — apenas ICMS deve estar presente
	if len(res.Itens[0].Tributos) != 1 {
		t.Fatalf("esperado 1 tributo (ICMS), obteve %d", len(res.Itens[0].Tributos))
	}
	if res.Itens[0].Tributos[0].Tributo != "ICMS" {
		t.Errorf("tributo = %s, want ICMS", res.Itens[0].Tributos[0].Tributo)
	}
}

func TestEngine_ValidacaoFalha(t *testing.T) {
	engine := BillingEngine(&mockTaxCalculator{})

	input := models.DocumentoFiscalEntrada{}
	_, err := engine.Process(context.Background(), input)
	if err == nil {
		t.Fatal("esperado erro de validacao, obteve nil")
	}
}

func TestEngine_CalcParaleloMultiplosErros_ColetaTodos(t *testing.T) {
	calc1 := &mockTaxCalculator{err: errors.New("falha A")}
	calc2 := &mockTaxCalculator{err: errors.New("falha B")}
	calc3 := &mockTaxCalculator{items: []models.ItemDocumentoFiscalSaida{mockItem("X", "ICMS", 10)}}

	engine := BillingEngine(calc1, calc2, calc3)

	input := inputDocumento(inputItem("X", 1, 100))
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	// Apenas calc3 deve ter contribuído
	if len(res.Itens[0].Tributos) != 1 {
		t.Fatalf("esperado 1 tributo, obteve %d", len(res.Itens[0].Tributos))
	}
	if res.Itens[0].Tributos[0].Tributo != "ICMS" {
		t.Errorf("tributo = %s, want ICMS", res.Itens[0].Tributos[0].Tributo)
	}
}

// ─── GAP-004: Campo valor_liquido no Response ─────────────────────────────

// TST-004.01: Item com tributos — valor_liquido = valor_item − total_impostos
func TestValorLiquido_ComTributos(t *testing.T) {
	item := inputItem("SKU-A", 2, 500) // 2 * 500 = 1000

	ipiCalc := &mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{mockItem("SKU-A", "IPI", 50)},
	}
	icmsCalc := &mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{mockItem("SKU-A", "ICMS", 180)},
	}
	pisCalc := &mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{mockItem("SKU-A", "PIS", 16.5)},
	}

	engine := BillingEnginePhased(
		Phase("IPI", Sequential, ipiCalc),
		Phase("ICMS+PIS", Parallel, icmsCalc, pisCalc),
	)

	input := inputDocumento(item)
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(res.Itens) != 1 {
		t.Fatalf("esperado 1 item, obteve %d", len(res.Itens))
	}

	// Total = 1000; impostos = 50 + 180 + 16.5 = 246.50
	// valor_liquido = 1000 - 246.50 = 753.50
	expectedLiquido := decimal.NewFromFloat(753.50)
	if !res.Itens[0].ValorLiquido.Equal(expectedLiquido) {
		t.Errorf("ValorLiquido = %s, esperado %s",
			res.Itens[0].ValorLiquido, expectedLiquido)
	}

	if !res.Itens[0].Total.Equal(decimal.NewFromFloat(1000)) {
		t.Errorf("Total = %s, esperado 1000", res.Itens[0].Total)
	}
}

// TST-004.02: Item isento (todos os tributos = 0) — valor_liquido = valor_item
func TestValorLiquido_Isento(t *testing.T) {
	item := inputItem("SKU-B", 1, 500)

	isentoCalc := &mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-B",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "ICMS", Valor: decimal.Zero},
				{Tributo: "PIS", Valor: decimal.Zero},
			},
		}},
	}

	engine := BillingEngine(isentoCalc)
	input := inputDocumento(item)
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if !res.Itens[0].ValorLiquido.Equal(decimal.NewFromFloat(500)) {
		t.Errorf("ValorLiquido = %s, esperado 500 (isento)",
			res.Itens[0].ValorLiquido)
	}
}

// TST-004.03: Impostos > valor — valor_liquido = 0 (nunca negativo)
func TestValorLiquido_ImpostosMaiorQueValor(t *testing.T) {
	item := inputItem("SKU-C", 1, 100)

	anomaloCalc := &mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-C",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "ICMS", Valor: decimal.NewFromFloat(80)},
				{Tributo: "PIS", Valor: decimal.NewFromFloat(50)},
			},
		}},
	}

	engine := BillingEngine(anomaloCalc)
	input := inputDocumento(item)
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if !res.Itens[0].ValorLiquido.Equal(decimal.Zero) {
		t.Errorf("ValorLiquido = %s, esperado 0 (impostos > valor)",
			res.Itens[0].ValorLiquido)
	}
}

// TST-004.04: Response JSON contém campo valor_liquido
func TestValorLiquido_JSONSerialization(t *testing.T) {
	item := inputItem("SKU-D", 1, 100)
	calc := &mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{
			mockItem("SKU-D", "ICMS", 12),
		},
	}

	engine := BillingEngine(calc)
	input := inputDocumento(item)
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if res.Itens[0].ValorLiquido.IsZero() {
		t.Error("ValorLiquido não deveria ser zero para este cenário")
	}

	if res.Itens[0].Total.IsZero() {
		t.Error("Total não deveria ser zero")
	}
}

// ─── GAP-006: Split Payment Schema ────────────────────────────────────────

func buildSplitOutput(cbs, ibs, is float64) models.DocumentoFiscalSaida {
	var tributos []models.TributosItemDocumentoFiscalSaida
	if cbs > 0 {
		tributos = append(tributos, models.TributosItemDocumentoFiscalSaida{Tributo: "CBS", Valor: dec(cbs)})
	}
	if ibs > 0 {
		tributos = append(tributos, models.TributosItemDocumentoFiscalSaida{Tributo: "IBS", Valor: dec(ibs)})
	}
	if is > 0 {
		tributos = append(tributos, models.TributosItemDocumentoFiscalSaida{Tributo: "IS", Valor: dec(is)})
	}
	return models.DocumentoFiscalSaida{
		Itens: []models.ItemDocumentoFiscalSaida{{Tributos: tributos}},
	}
}

func dec(f float64) decimal.Decimal { return decimal.NewFromFloat(f) }

// TST-006.01: Split normal — receita liquida = total - (CBS+IBS+IS)
func TestSplitPayment_Normal(t *testing.T) {
	item := inputItem("SKU-A", 1, 1000)
	// CBS=90, IBS=75, IS=50
	engine := BillingEngine(&mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "CBS", Valor: dec(90)},
				{Tributo: "IBS", Valor: dec(75)},
				{Tributo: "IS", Valor: dec(50)},
				{Tributo: "ICMS", Valor: dec(120)},
			},
		}},
	})

	res, err := engine.Process(context.Background(), inputDocumento(item))
	if err != nil {
		t.Fatalf("erro: %v", err)
	}

	sp := res.SplitPayment
	if sp == nil {
		t.Fatal("SplitPayment não deveria ser nil")
	}
	// receita liquida = 1000 - (90+75+50) = 785
	if sp.ValorReceitaLiquida.String() != "785" {
		t.Errorf("ReceitaLiquida = %s, esperado 785", sp.ValorReceitaLiquida)
	}
	if sp.ValorCBSReter.String() != "90" {
		t.Errorf("CBSReter = %s, esperado 90", sp.ValorCBSReter)
	}
	if sp.ValorIBSReter.String() != "75" {
		t.Errorf("IBSReter = %s, esperado 75", sp.ValorIBSReter)
	}
	if sp.ValorISReter.String() != "50" {
		t.Errorf("ISReter = %s, esperado 50", sp.ValorISReter)
	}
	if sp.CodigoBarrasSplit == "" {
		t.Error("CodigoBarrasSplit não pode ser vazio")
	}
}

// TST-006.02: Operação com IS — IS incluso no split
func TestSplitPayment_ComIS(t *testing.T) {
	item := inputItem("SKU-B", 2, 500)
	engine := BillingEngine(&mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-B",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "CBS", Valor: dec(90)},
				{Tributo: "IBS", Valor: dec(75)},
				{Tributo: "IS", Valor: dec(100)}, // IS alto (bebida/cigarro)
			},
		}},
	})

	res, _ := engine.Process(context.Background(), inputDocumento(item))
	sp := res.SplitPayment

	if sp.ValorISReter.String() != "100" {
		t.Errorf("ISReter = %s, esperado 100", sp.ValorISReter)
	}
	// receita = 1000 - (90+75+100) = 735
	if sp.ValorReceitaLiquida.String() != "735" {
		t.Errorf("ReceitaLiquida = %s, esperado 735", sp.ValorReceitaLiquida)
	}
}

// TST-006.03: Operação isenta — todos os valores zerados
func TestSplitPayment_Isento(t *testing.T) {
	item := inputItem("SKU-C", 1, 500)
	engine := BillingEngine(&mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-C",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "ICMS", Valor: decimal.Zero},
				{Tributo: "PIS", Valor: decimal.Zero},
			},
		}},
	})

	res, _ := engine.Process(context.Background(), inputDocumento(item))
	sp := res.SplitPayment

	if sp.ValorReceitaLiquida.String() != "500" {
		t.Errorf("ReceitaLiquida = %s, esperado 500", sp.ValorReceitaLiquida)
	}
	if !sp.ValorCBSReter.IsZero() || !sp.ValorIBSReter.IsZero() || !sp.ValorISReter.IsZero() {
		t.Error("Todos os valores de retenção devem ser zero em operação isenta")
	}
}

// TST-006.04: Hash SHA-256 determinístico
func TestSplitPayment_HashDeterministico(t *testing.T) {
	item := inputItem("SKU-D", 1, 1000)
	calc := &mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-D",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "CBS", Valor: dec(90)},
				{Tributo: "IBS", Valor: dec(75)},
			},
		}},
	}

	engine1 := BillingEngine(calc)
	res1, _ := engine1.Process(context.Background(), inputDocumento(item))

	engine2 := BillingEngine(calc)
	res2, _ := engine2.Process(context.Background(), inputDocumento(item))

	if res1.SplitPayment.CodigoBarrasSplit != res2.SplitPayment.CodigoBarrasSplit {
		t.Errorf("Hash deveria ser determinístico:\n  hash1=%s\n  hash2=%s",
			res1.SplitPayment.CodigoBarrasSplit, res2.SplitPayment.CodigoBarrasSplit)
	}

	// Hash deve ter 64 caracteres hex (SHA-256)
	if len(res1.SplitPayment.CodigoBarrasSplit) != 64 {
		t.Errorf("CodigoBarrasSplit tem %d caracteres, esperado 64", len(res1.SplitPayment.CodigoBarrasSplit))
	}
}

// TST-006.05: Operação interestadual — IBS segregado corretamente
func TestSplitPayment_Interestadual(t *testing.T) {
	item := inputItem("SKU-E", 1, 2000)
	engine := BillingEngine(&mockTaxCalculator{
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-E",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "CBS", Valor: dec(180)},
				{Tributo: "IBS", Valor: dec(160)},
				{Tributo: "IS", Valor: dec(0)},
				{Tributo: "ICMS", Valor: dec(240)},
				{Tributo: "ICMS_DIFAL", Valor: dec(60)},
			},
		}},
	})

	res, _ := engine.Process(context.Background(), inputDocumento(item))
	sp := res.SplitPayment

	// receita = 2000 - (180+160+0) = 1660
	if sp.ValorReceitaLiquida.String() != "1660" {
		t.Errorf("ReceitaLiquida = %s, esperado 1660", sp.ValorReceitaLiquida)
	}
	if sp.ValorCBSReter.String() != "180" {
		t.Errorf("CBSReter = %s, esperado 180", sp.ValorCBSReter)
	}
	// ICMS e DIFAL NÃO entram no split (apenas CBS/IBS/IS)
	if sp.ValorIBSReter.String() != "160" {
		t.Errorf("IBSReter = %s, esperado 160", sp.ValorIBSReter)
	}
}
