package legacy

import (
	"context"
	"testing"

	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

// TestFUNTTEL_SCM_Incide verifica o cálculo de FUNTTEL para SCM.
func TestFUNTTEL_SCM_Incide(t *testing.T) {
	calc := NewFUNTTELCalculator()
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
		t.Fatalf("esperado 1 tributo FUNTTEL, obtido %d", len(result[0].Tributos))
	}

	trib := result[0].Tributos[0]
	if trib.Tributo != "FUNTTEL" {
		t.Errorf("tributo = %q, esperado FUNTTEL", trib.Tributo)
	}

	// Exemplo da política:
	// Base = 70,00 − 17,50 − 0,45 − 2,10 = R$ 49,95
	// FUNTTEL = 49,95 × 0,5% = R$ 0,25
	baseEsperada := decimal.NewFromFloat(49.95)
	valorEsperado := decimal.NewFromFloat(0.25)

	if !trib.BaseCalculo.Equal(baseEsperada) {
		t.Errorf("base FUNTTEL = %s, esperado %s", trib.BaseCalculo, baseEsperada)
	}
	if !trib.Valor.Equal(valorEsperado) {
		t.Errorf("valor FUNTTEL = %s, esperado %s", trib.Valor, valorEsperado)
	}
}

// TestFUNTTEL_MesmaBaseFUST verifica que FUNTTEL usa a mesma base que FUST.
func TestFUNTTEL_MesmaBaseFUST(t *testing.T) {
	fustCalc := NewFUSTCalculator()
	funttelCalc := NewFUNTTELCalculator()

	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemServicoTelecomComTributos("SCM-001", 100.00, 1, "SCM", 18.00, 1.65, 7.60),
		},
	}

	fustResult, _ := fustCalc.Calculate(context.Background(), input)
	funttelResult, _ := funttelCalc.Calculate(context.Background(), input)

	fustBase := fustResult[0].Tributos[0].BaseCalculo
	funttelBase := funttelResult[0].Tributos[0].BaseCalculo

	if !fustBase.Equal(funttelBase) {
		t.Errorf("bases diferentes: FUST=%s, FUNTTEL=%s", fustBase, funttelBase)
	}

	// FUST = 72,75 × 1% = 0,73
	// FUNTTEL = 72,75 × 0,5% = 0,36
	fustEsperado := decimal.NewFromFloat(0.73)
	funttelEsperado := decimal.NewFromFloat(0.36)

	if !fustResult[0].Tributos[0].Valor.Equal(fustEsperado) {
		t.Errorf("FUST = %s, esperado %s", fustResult[0].Tributos[0].Valor, fustEsperado)
	}
	if !funttelResult[0].Tributos[0].Valor.Equal(funttelEsperado) {
		t.Errorf("FUNTTEL = %s, esperado %s", funttelResult[0].Tributos[0].Valor, funttelEsperado)
	}
}

// TestFUNTTEL_SVA_NaoIncide verifica que SVA não paga FUNTTEL.
func TestFUNTTEL_SVA_NaoIncide(t *testing.T) {
	calc := NewFUNTTELCalculator()
	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemServicoTelecomComTributos("SVA-001", 50.00, 1, "SVA", 0, 0, 0),
		},
	}

	result, _ := calc.Calculate(context.Background(), input)
	if len(result[0].Tributos) != 0 {
		t.Error("SVA não deve ter FUNTTEL calculado")
	}
}

// TestFUNTTEL_BaseLiquidaNegativa verifica que base negativa resulta em FUNTTEL = 0.
func TestFUNTTEL_BaseLiquidaNegativa(t *testing.T) {
	calc := NewFUNTTELCalculator()
	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemServicoTelecomComTributos("SCM-NEG", 10.00, 1, "SCM", 8.00, 3.00, 5.00),
		},
	}

	result, _ := calc.Calculate(context.Background(), input)
	trib := result[0].Tributos[0]

	if !trib.Valor.IsZero() {
		t.Errorf("FUNTTEL com base negativa deveria ser zero, obtido %s", trib.Valor)
	}
}
