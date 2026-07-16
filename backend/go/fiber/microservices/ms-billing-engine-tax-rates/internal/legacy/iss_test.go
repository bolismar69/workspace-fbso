package legacy

import (
	"context"
	"testing"

	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

// itemServicoTelecom cria um item de serviço de telecom com alíquota ISS.
func itemServicoTelecom(sku string, valorUnitario, quantidade, aliquotaISS float64) models.ItemDocumentoFiscalEntrada {
	item := models.ItemDocumentoFiscalEntrada{
		SKU:           sku,
		NCM:           "00000000", // Serviços puros podem não ter NCM
		ValorUnitario: decimal.NewFromFloat(valorUnitario),
		Quantidade:    decimal.NewFromFloat(quantidade),
	}
	item.AddDetalhe(models.KeyDocumentoInfos("ITEM_LISTA_SERVICO"), "1.05")
	item.AddDetalhe(models.KeyDocumentoInfos("ISS_ALIQUOTA"), decimal.NewFromFloat(aliquotaISS))
	return item
}

// TestISS_Valor_AliquotaMunicipal verifica o cálculo básico do ISS.
func TestISS_Valor_AliquotaMunicipal(t *testing.T) {
	calc := NewISSCalculator()
	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemServicoTelecom("SKU-001", 1000.00, 1, 5.00),
		},
	}

	result, err := calc.Calculate(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("esperado 1 item, obtido %d", len(result))
	}

	tributos := result[0].Tributos
	if len(tributos) != 1 {
		t.Fatalf("esperado 1 tributo, obtido %d", len(tributos))
	}

	trib := tributos[0]
	if trib.Tributo != "ISS" {
		t.Errorf("tributo = %q, esperado ISS", trib.Tributo)
	}

	valorEsperado := decimal.NewFromFloat(50.00) // 1000 × 5% = 50
	if !trib.Valor.Equal(valorEsperado) {
		t.Errorf("valor ISS = %s, esperado %s", trib.Valor, valorEsperado)
	}

	if !trib.BaseCalculo.Equal(decimal.NewFromFloat(1000)) {
		t.Errorf("base cálculo = %s, esperado 1000", trib.BaseCalculo)
	}
}

// TestISS_AliquotaForaDoIntervalo_ContinuaComWarning verifica que alíquota
// fora do range [2%, 5%] ainda calcula (com warning em log).
func TestISS_AliquotaForaDoIntervalo_ContinuaComWarning(t *testing.T) {
	calc := NewISSCalculator()
	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemServicoTelecom("SKU-002", 500.00, 2, 6.00), // Alíquota 6% — acima do máximo
		},
	}

	result, err := calc.Calculate(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(result) != 1 || len(result[0].Tributos) != 1 {
		t.Fatal("esperado que o ISS seja calculado mesmo com alíquota fora do range")
	}

	// 1000 × 6% = 60
	valorEsperado := decimal.NewFromFloat(60.00)
	if !result[0].Tributos[0].Valor.Equal(valorEsperado) {
		t.Errorf("valor ISS = %s, esperado %s", result[0].Tributos[0].Valor, valorEsperado)
	}
}

// TestISS_AliquotaMinimaPermitida verifica alíquota no limite inferior (2%).
func TestISS_AliquotaMinimaPermitida(t *testing.T) {
	calc := NewISSCalculator()
	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemServicoTelecom("SKU-003", 2000.00, 1, 2.00),
		},
	}

	result, _ := calc.Calculate(context.Background(), input)
	valorEsperado := decimal.NewFromFloat(40.00) // 2000 × 2%
	if !result[0].Tributos[0].Valor.Equal(valorEsperado) {
		t.Errorf("valor ISS = %s, esperado %s", result[0].Tributos[0].Valor, valorEsperado)
	}
}

// TestISS_ItemSemServico_NaoIncide verifica que mercadorias não pagam ISS.
func TestISS_ItemSemServico_NaoIncide(t *testing.T) {
	calc := NewISSCalculator()
	// Item sem ITEM_LISTA_SERVICO → é mercadoria, ISS não incide
	item := models.ItemDocumentoFiscalEntrada{
		SKU:           "MERCADORIA-001",
		NCM:           "84713012",
		ValorUnitario: decimal.NewFromFloat(5000),
		Quantidade:    decimal.NewFromFloat(1),
	}

	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{item},
	}

	result, err := calc.Calculate(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("esperado 1 item, obtido %d", len(result))
	}

	if len(result[0].Tributos) != 0 {
		t.Errorf("mercadoria não deve ter ISS, mas tem %d tributos", len(result[0].Tributos))
	}
}

// TestISS_MultiplosItensMistos verifica o cálculo com itens de serviço e mercadoria.
func TestISS_MultiplosItensMistos(t *testing.T) {
	calc := NewISSCalculator()
	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemServicoTelecom("SERVICO-001", 100.00, 1, 3.00),
			{SKU: "MERCADORIA-001", NCM: "84713012", ValorUnitario: decimal.NewFromFloat(500), Quantidade: decimal.NewFromFloat(1)},
			itemServicoTelecom("SERVICO-002", 200.00, 2, 5.00),
		},
	}

	result, err := calc.Calculate(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(result) != 3 {
		t.Fatalf("esperado 3 itens, obtido %d", len(result))
	}

	// Item 0: serviço, deve ter ISS
	if len(result[0].Tributos) == 0 {
		t.Error("item 0 (serviço) deveria ter ISS")
	} else {
		// 100 × 3% = 3.00
		if !result[0].Tributos[0].Valor.Equal(decimal.NewFromFloat(3.00)) {
			t.Errorf("item 0 ISS = %s, esperado 3.00", result[0].Tributos[0].Valor)
		}
	}

	// Item 1: mercadoria, não deve ter ISS
	if len(result[1].Tributos) != 0 {
		t.Error("item 1 (mercadoria) não deveria ter ISS")
	}

	// Item 2: serviço, deve ter ISS: 400 × 5% = 20.00
	if len(result[2].Tributos) == 0 {
		t.Error("item 2 (serviço) deveria ter ISS")
	} else {
		if !result[2].Tributos[0].Valor.Equal(decimal.NewFromFloat(20.00)) {
			t.Errorf("item 2 ISS = %s, esperado 20.00", result[2].Tributos[0].Valor)
		}
	}
}

// TestISS_RetencaoFonte verifica que a flag de retenção é registrada nos detalhes.
func TestISS_RetencaoFonte(t *testing.T) {
	calc := NewISSCalculator()
	item := itemServicoTelecom("SKU-RET", 1000.00, 1, 4.00)
	item.AddDetalhe(models.KeyDocumentoInfos("ISS_RETIDO"), "true")

	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{item},
	}

	result, _ := calc.Calculate(context.Background(), input)
	trib := result[0].Tributos[0]

	// Verifica que a flag de retenção aparece nos detalhes textuais
	found := false
	for _, d := range trib.MoreTextDetails {
		if d.Key == "retencao_fonte" && d.Value == "sim" {
			found = true
			break
		}
	}
	if !found {
		t.Error("flag retencao_fonte='sim' não encontrada nos MoreTextDetails")
	}
}

// TestISS_SemAliquota_NaoCalcula verifica que item sem alíquota configurada é pulado.
func TestISS_SemAliquota_NaoCalcula(t *testing.T) {
	calc := NewISSCalculator()
	item := models.ItemDocumentoFiscalEntrada{
		SKU:           "SKU-SEM-ALIQ",
		NCM:           "00000000",
		ValorUnitario: decimal.NewFromFloat(100),
		Quantidade:    decimal.NewFromFloat(1),
	}
	item.AddDetalhe(models.KeyDocumentoInfos("ITEM_LISTA_SERVICO"), "1.05")
	// Sem ISS_ALIQUOTA

	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{item},
	}

	result, _ := calc.Calculate(context.Background(), input)
	if len(result[0].Tributos) != 0 {
		t.Error("item sem alíquota não deveria ter ISS calculado")
	}
}
