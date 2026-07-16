package legacy

import (
	"context"
	"testing"

	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

// itemServicoTelecomComTributos cria um item de serviço telecom com valores
// de ICMS, PIS e COFINS já injetados (simulando o que a engine faz na Fase 2).
func itemServicoTelecomComTributos(sku string, valorUnitario, quantidade float64, natureza string, icms, pis, cofins float64) models.ItemDocumentoFiscalEntrada {
	item := models.ItemDocumentoFiscalEntrada{
		SKU:           sku,
		NCM:           "00000000",
		ValorUnitario: decimal.NewFromFloat(valorUnitario),
		Quantidade:    decimal.NewFromFloat(quantidade),
	}
	item.AddDetalhe(models.KeyDocumentoInfos("NATUREZA_SERVICO"), natureza)
	if icms > 0 {
		item.AddDetalhe(models.KeyDocumentoInfos("ITEM_ICMS_VALOR"), decimal.NewFromFloat(icms))
	}
	if pis > 0 {
		item.AddDetalhe(models.KeyDocumentoInfos("ITEM_PIS_VALOR"), decimal.NewFromFloat(pis))
	}
	if cofins > 0 {
		item.AddDetalhe(models.KeyDocumentoInfos("ITEM_COFINS_VALOR"), decimal.NewFromFloat(cofins))
	}
	return item
}

// TestFUST_SCM_Incide verifica o cálculo de FUST para serviço SCM.
func TestFUST_SCM_Incide(t *testing.T) {
	calc := NewFUSTCalculator()
	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemServicoTelecomComTributos("SCM-001", 70.00, 1, "SCM", 17.50, 0.45, 2.10),
		},
	}

	result, err := calc.Calculate(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(result) != 1 || len(result[0].Tributos) != 1 {
		t.Fatalf("esperado 1 tributo FUST, obtido %d", len(result[0].Tributos))
	}

	trib := result[0].Tributos[0]
	if trib.Tributo != "FUST" {
		t.Errorf("tributo = %q, esperado FUST", trib.Tributo)
	}

	// Exemplo da política:
	// SCM = R$ 70,00; ICMS = R$ 17,50; PIS = R$ 0,45; COFINS = R$ 2,10
	// Base FUST = 70,00 − 17,50 − 0,45 − 2,10 = R$ 49,95
	// FUST = 49,95 × 1% = R$ 0,50
	baseEsperada := decimal.NewFromFloat(49.95)
	valorEsperado := decimal.NewFromFloat(0.50)

	if !trib.BaseCalculo.Equal(baseEsperada) {
		t.Errorf("base FUST = %s, esperado %s", trib.BaseCalculo, baseEsperada)
	}
	if !trib.Valor.Equal(valorEsperado) {
		t.Errorf("valor FUST = %s, esperado %s", trib.Valor, valorEsperado)
	}
}

// TestFUST_STFC_Incide verifica FUST para STFC (telefonia fixa).
func TestFUST_STFC_Incide(t *testing.T) {
	calc := NewFUSTCalculator()
	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemServicoTelecomComTributos("STFC-001", 100.00, 1, "STFC", 25.00, 0.65, 3.00),
		},
	}

	result, _ := calc.Calculate(context.Background(), input)

	// Base: 100 − 25 − 0,65 − 3,00 = 71,35
	// FUST: 71,35 × 1% = 0,71
	baseEsperada := decimal.NewFromFloat(71.35)
	valorEsperado := decimal.NewFromFloat(0.71)

	if !result[0].Tributos[0].BaseCalculo.Equal(baseEsperada) {
		t.Errorf("base FUST = %s, esperado %s", result[0].Tributos[0].BaseCalculo, baseEsperada)
	}
	if !result[0].Tributos[0].Valor.Equal(valorEsperado) {
		t.Errorf("valor FUST = %s, esperado %s", result[0].Tributos[0].Valor, valorEsperado)
	}
}

// TestFUST_SVA_NaoIncide verifica que SVA não paga FUST.
func TestFUST_SVA_NaoIncide(t *testing.T) {
	calc := NewFUSTCalculator()
	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemServicoTelecomComTributos("SVA-001", 50.00, 1, "SVA", 0, 0, 0),
		},
	}

	result, _ := calc.Calculate(context.Background(), input)
	if len(result[0].Tributos) != 0 {
		t.Error("SVA não deve ter FUST calculado")
	}
}

// TestFUST_SemNatureza_NaoIncide verifica que item sem natureza definida não paga FUST.
func TestFUST_SemNatureza_NaoIncide(t *testing.T) {
	calc := NewFUSTCalculator()
	item := models.ItemDocumentoFiscalEntrada{
		SKU:           "SEM-NATUREZA",
		NCM:           "00000000",
		ValorUnitario: decimal.NewFromFloat(100),
		Quantidade:    decimal.NewFromFloat(1),
	}
	// Sem NATUREZA_SERVICO

	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{item},
	}

	result, _ := calc.Calculate(context.Background(), input)
	if len(result[0].Tributos) != 0 {
		t.Error("item sem natureza de serviço não deve ter FUST")
	}
}

// TestFUST_BaseLiquidaNegativa verifica que base negativa resulta em FUST = 0.
func TestFUST_BaseLiquidaNegativa(t *testing.T) {
	calc := NewFUSTCalculator()
	// ICMS + PIS + COFINS > Valor do serviço → base negativa
	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemServicoTelecomComTributos("SCM-NEG", 10.00, 1, "SCM", 8.00, 3.00, 5.00),
		},
	}

	result, _ := calc.Calculate(context.Background(), input)
	trib := result[0].Tributos[0]

	if !trib.BaseCalculo.IsZero() {
		t.Errorf("base negativa deveria ser zero, obtido %s", trib.BaseCalculo)
	}
	if !trib.Valor.IsZero() {
		t.Errorf("FUST com base negativa deveria ser zero, obtido %s", trib.Valor)
	}
}

// TestFUST_MultiplosItensMistos verifica cálculo com itens SCM, SVA e sem natureza.
func TestFUST_MultiplosItensMistos(t *testing.T) {
	calc := NewFUSTCalculator()
	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemServicoTelecomComTributos("SCM-001", 100.00, 1, "SCM", 18.00, 0.65, 3.00),
			itemServicoTelecomComTributos("SVA-001", 50.00, 1, "SVA", 0, 0, 0),
			itemServicoTelecomComTributos("SCM-002", 200.00, 2, "SCM", 72.00, 2.60, 12.00),
		},
	}

	result, _ := calc.Calculate(context.Background(), input)

	// Item 0 (SCM): Base = 100 − 18 − 0,65 − 3 = 78,35; FUST = 0,78
	if len(result[0].Tributos) == 0 {
		t.Error("item 0 (SCM) deveria ter FUST")
	} else {
		esperado := decimal.NewFromFloat(0.78)
		if !result[0].Tributos[0].Valor.Equal(esperado) {
			t.Errorf("item 0 FUST = %s, esperado %s", result[0].Tributos[0].Valor, esperado)
		}
	}

	// Item 1 (SVA): sem FUST
	if len(result[1].Tributos) != 0 {
		t.Error("item 1 (SVA) não deveria ter FUST")
	}

	// Item 2 (SCM): Base = 400 − 72 − 2,60 − 12 = 313,40; FUST = 3,13
	if len(result[2].Tributos) == 0 {
		t.Error("item 2 (SCM) deveria ter FUST")
	} else {
		esperado := decimal.NewFromFloat(3.13)
		if !result[2].Tributos[0].Valor.Equal(esperado) {
			t.Errorf("item 2 FUST = %s, esperado %s", result[2].Tributos[0].Valor, esperado)
		}
	}
}
