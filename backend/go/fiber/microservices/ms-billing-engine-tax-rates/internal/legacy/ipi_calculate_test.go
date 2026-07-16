package legacy

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"

	"github.com/shopspring/decimal"
)

func defaultIPIRegra(aliquota, valorPauta float64, cst, cEnq string) *repository.IPIRegra {
	return &repository.IPIRegra{
		ID:                 1,
		NCM:                "84713019",
		CrtEmitente:        "3",
		TipoOperacaoFiscal: "VENDA",
		AliquotaIPI:        decimal.NewFromFloat(aliquota),
		ValorPautaIPI:      decimal.NewFromFloat(valorPauta),
		CSTIPI:             cst,
		CEnq:               cEnq,
		InicioValidade:     time.Now().AddDate(-1, 0, 0),
	}
}

func ipiItem(ncm string, quantidade, valorUnitario float64, detalhes ...models.Detalhe) models.ItemDocumentoFiscalEntrada {
	return models.ItemDocumentoFiscalEntrada{
		SKU:                         "SKU-" + ncm,
		NCM:                         ncm,
		Quantidade:                  decimal.NewFromFloat(quantidade),
		ValorUnitario:               decimal.NewFromFloat(valorUnitario),
		DetalhesItemDocumentoFiscal: detalhes,
	}
}

func ipiDocumento(itens ...models.ItemDocumentoFiscalEntrada) models.DocumentoFiscalEntrada {
	return models.DocumentoFiscalEntrada{
		CRTEmitente:       "3",
		TipoOperacaoFiscal: "VENDA",
		LocalizacaoOrigem:  models.LocalizacaoFiscal{UF: "SP"},
		LocalizacaoDestino: models.LocalizacaoFiscal{UF: "SP"},
		Itens:              itens,
		DataOperacao:       time.Now(),
	}
}

func TestIPI_AdValorem_Basico(t *testing.T) {
	mock := &mockTaxRepository{
		ipiRegra: defaultIPIRegra(10, 0, "50", "999"),
	}
	calc := NewIPICalculator(mock)

	doc := ipiDocumento(ipiItem("84713019", 2, 500))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(res) != 1 || len(res[0].Tributos) != 1 {
		t.Fatalf("esperado 1 item com 1 tributo, obteve %d itens", len(res))
	}

	trib := res[0].Tributos[0]
	if trib.Tributo != "IPI" {
		t.Errorf("tributo = %s, want IPI", trib.Tributo)
	}
	if trib.CST != "50" {
		t.Errorf("CST = %s, want 50", trib.CST)
	}
	// Base: 2*500 = 1000
	if trib.BaseCalculo.StringFixed(2) != "1000.00" {
		t.Errorf("base = %s, want 1000.00", trib.BaseCalculo.StringFixed(2))
	}
	// Valor: 1000 * 10% = 100.00
	if trib.Valor.StringFixed(2) != "100.00" {
		t.Errorf("valor = %s, want 100.00", trib.Valor.StringFixed(2))
	}
	assertDetailText(t, trib.MoreTextDetails, "metodo_calculo", "AD_VALOREM")
	assertDetailText(t, trib.MoreTextDetails, "fonte_regra", "repositorio")
}

func TestIPI_ComRateioDespesas(t *testing.T) {
	mock := &mockTaxRepository{
		ipiRegra: defaultIPIRegra(10, 0, "50", "999"),
	}
	calc := NewIPICalculator(mock)

	doc := ipiDocumento(ipiItem("84713019", 1, 1000))
	doc.DetalhesDocumentoFiscal = []models.Detalhe{
		{Key: string(models.KeyDocumentoInfosValorFrete), Value: decimal.NewFromInt(100)},
		{Key: string(models.KeyDocumentoInfosValorSeguro), Value: decimal.NewFromInt(50)},
		{Key: string(models.KeyDocumentoInfosValorOutrasDespesas), Value: decimal.NewFromInt(30)},
		{Key: string(models.KeyDocumentoInfosValorDesconto), Value: decimal.NewFromInt(20)},
	}

	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	trib := res[0].Tributos[0]
	// Base: 1000 + 100 + 50 + 30 - 20 = 1160
	if trib.BaseCalculo.StringFixed(2) != "1160.00" {
		t.Errorf("base = %s, want 1160.00", trib.BaseCalculo.StringFixed(2))
	}
	// Valor: 1160 * 10% = 116.00
	if trib.Valor.StringFixed(2) != "116.00" {
		t.Errorf("valor = %s, want 116.00", trib.Valor.StringFixed(2))
	}
}

func TestIPI_AdPauta(t *testing.T) {
	mock := &mockTaxRepository{
		ipiRegra: defaultIPIRegra(0, 2.50, "50", "999"),
	}
	calc := NewIPICalculator(mock)

	doc := ipiDocumento(ipiItem("22030000", 24, 5))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	trib := res[0].Tributos[0]
	// Valor: 24 * R$ 2.50 = 60.00
	if trib.Valor.StringFixed(2) != "60.00" {
		t.Errorf("valor = %s, want 60.00 (24 * 2.50)", trib.Valor.StringFixed(2))
	}
	assertDetailText(t, trib.MoreTextDetails, "metodo_calculo", "AD_PAUTA")
}

func TestIPI_CompleteDetalheOverride(t *testing.T) {
	mock := &mockTaxRepository{} // sem regra no repo
	calc := NewIPICalculator(mock)

	doc := ipiDocumento(ipiItem("84713019", 1, 2000,
		models.Detalhe{Key: string(models.KeyDocumentoInfosItemIPIAliquota), Value: decimal.NewFromInt(15)},
		models.Detalhe{Key: string(models.KeyDocumentoInfosItemIPICST), Value: "50"},
		models.Detalhe{Key: string(models.KeyDocumentoInfosItemIPICEnq), Value: "999"},
	))

	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	trib := res[0].Tributos[0]
	if trib.Valor.StringFixed(2) != "300.00" {
		t.Errorf("valor = %s, want 300.00 (2000*15%%)", trib.Valor.StringFixed(2))
	}
	assertDetailText(t, trib.MoreTextDetails, "fonte_regra", "detalhe_item")
}

func TestIPI_RepoErrorWithOverrideInline(t *testing.T) {
	mock := &mockTaxRepository{
		ipiRegraErr: errors.New("db offline"),
	}
	calc := NewIPICalculator(mock)

	doc := ipiDocumento(ipiItem("84713019", 1, 500,
		models.Detalhe{Key: string(models.KeyDocumentoInfosItemIPIAliquota), Value: decimal.NewFromInt(5)},
	))

	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	trib := res[0].Tributos[0]
	if trib.Valor.StringFixed(2) != "25.00" {
		t.Errorf("valor = %s, want 25.00 (500*5%%)", trib.Valor.StringFixed(2))
	}
}

func TestIPI_RepoErrorNoOverride_Skip(t *testing.T) {
	mock := &mockTaxRepository{
		ipiRegraErr: errors.New("db offline"),
	}
	calc := NewIPICalculator(mock)

	doc := ipiDocumento(ipiItem("84713019", 1, 500))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(res[0].Tributos) != 0 {
		t.Errorf("esperado 0 tributos (item pulado), obteve %d", len(res[0].Tributos))
	}
}

func TestIPI_MultiplosItensComRateio(t *testing.T) {
	mock := &mockTaxRepository{
		ipiRegra: defaultIPIRegra(5, 0, "50", "999"),
	}
	calc := NewIPICalculator(mock)

	doc := ipiDocumento(
		ipiItem("84713019", 1, 600),
		ipiItem("84713020", 1, 400),
	)
	doc.DetalhesDocumentoFiscal = []models.Detalhe{
		{Key: string(models.KeyDocumentoInfosValorFrete), Value: decimal.NewFromInt(200)},
	}

	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	// Item 1: 60% do rateio → frete rateado = 120
	trib1 := res[0].Tributos[0]
	// Base: 600 + 120 = 720
	if trib1.BaseCalculo.StringFixed(2) != "720.00" {
		t.Errorf("item1 base = %s, want 720.00", trib1.BaseCalculo.StringFixed(2))
	}
	// Valor: 720 * 5% = 36.00
	if trib1.Valor.StringFixed(2) != "36.00" {
		t.Errorf("item1 valor = %s, want 36.00", trib1.Valor.StringFixed(2))
	}

	// Item 2: 40% do rateio → frete rateado = 80
	trib2 := res[1].Tributos[0]
	// Base: 400 + 80 = 480
	if trib2.BaseCalculo.StringFixed(2) != "480.00" {
		t.Errorf("item2 base = %s, want 480.00", trib2.BaseCalculo.StringFixed(2))
	}
	// Valor: 480 * 5% = 24.00
	if trib2.Valor.StringFixed(2) != "24.00" {
		t.Errorf("item2 valor = %s, want 24.00", trib2.Valor.StringFixed(2))
	}
}

func assertDetailText(t *testing.T, details []models.Detalhe, key, want string) {
	t.Helper()
	for _, d := range details {
		if d.Key == key {
			if s, ok := d.Value.(string); ok && s == want {
				return
			}
			if s, ok := d.Value.(fmt.Stringer); ok && s.String() == want {
				return
			}
			t.Errorf("detalhe %q = %v, want %s", key, d.Value, want)
			return
		}
	}
	t.Errorf("detalhe %q nao encontrado", key)
}
