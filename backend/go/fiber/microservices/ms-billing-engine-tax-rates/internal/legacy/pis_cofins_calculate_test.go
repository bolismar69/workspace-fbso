package legacy

import (
	"context"
	"errors"
	"testing"

	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"

	"github.com/shopspring/decimal"
)

func makeDocumentoFiscalEntrada(itens ...models.ItemDocumentoFiscalEntrada) models.DocumentoFiscalEntrada {
	return models.DocumentoFiscalEntrada{
		CRTEmitente: "3",
		Itens:       itens,
	}
}

func makeItem(cstPis, cstCofins string, quantidade, valorUnitario float64, valorIcmsExcluir float64) models.ItemDocumentoFiscalEntrada {
	detalhes := []models.Detalhe{
		{Key: string(models.KeyDocumentoInfosItemPISCOFINSCSTPIS), Value: cstPis},
		{Key: string(models.KeyDocumentoInfosItemPISCOFINSCSTCOFINS), Value: cstCofins},
		{Key: string(models.KeyDocumentoInfosValorExclusaoICMS), Value: decimal.NewFromFloat(valorIcmsExcluir)},
	}
	return models.ItemDocumentoFiscalEntrada{
		SKU:                         "SKU-TEST-" + cstPis,
		Quantidade:                  decimal.NewFromFloat(quantidade),
		ValorUnitario:               decimal.NewFromFloat(valorUnitario),
		DetalhesItemDocumentoFiscal: detalhes,
	}
}

func TestPISCofinsCalculator_ExcluiICMSBase_True(t *testing.T) {
	t.Run("ExcluiICMSBase=true: ICMS excluido da base de calculo", func(t *testing.T) {
		mock := newMockWithFederalRule(defaultFederalRule(1.65, 7.6, true))
		calc := NewPISCofinsCalculator(mock)

		item := makeItem("01", "01", 1, 1000, 120)
		input := makeDocumentoFiscalEntrada(item)

		tributos, err := calc.Calculate(context.Background(), input)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		if len(tributos) != 2 {
			t.Fatalf("esperado 2 tributos, obteve %d", len(tributos))
		}

		// PIS
		pis := tributos[0]
		if pis.Tributo != "PIS" {
			t.Errorf("tributo = %s, want PIS", pis.Tributo)
		}
		// Base = 1000 - 120 = 880
		if pis.BaseCalculo.StringFixed(2) != "880.00" {
			t.Errorf("PIS BaseCalculo = %s, want 880.00", pis.BaseCalculo.StringFixed(2))
		}
		// Valor = 880 * 1.65 / 100 = 14.52
		if pis.Valor.StringFixed(2) != "14.52" {
			t.Errorf("PIS Valor = %s, want 14.52", pis.Valor.StringFixed(2))
		}
		// Verifica detalhes
		assertDetail(t, pis.MoreNumericDetails, "icms_excluido_base", "120")
		assertDetail(t, pis.MoreTextDetails, "metodo", "base_com_exclusao_icms")

		// COFINS
		cofins := tributos[1]
		if cofins.Tributo != "COFINS" {
			t.Errorf("tributo = %s, want COFINS", cofins.Tributo)
		}
		if cofins.BaseCalculo.StringFixed(2) != "880.00" {
			t.Errorf("COFINS BaseCalculo = %s, want 880.00", cofins.BaseCalculo.StringFixed(2))
		}
		// Valor = 880 * 7.6 / 100 = 66.88
		if cofins.Valor.StringFixed(2) != "66.88" {
			t.Errorf("COFINS Valor = %s, want 66.88", cofins.Valor.StringFixed(2))
		}
		assertDetail(t, cofins.MoreTextDetails, "metodo", "base_com_exclusao_icms")
	})
}

func TestPISCofinsCalculator_ExcluiICMSBase_False(t *testing.T) {
	t.Run("ExcluiICMSBase=false: ICMS NAO excluido da base de calculo", func(t *testing.T) {
		mock := newMockWithFederalRule(defaultFederalRule(1.65, 7.6, false))
		calc := NewPISCofinsCalculator(mock)

		item := makeItem("01", "01", 1, 1000, 120)
		input := makeDocumentoFiscalEntrada(item)

		tributos, err := calc.Calculate(context.Background(), input)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		// PIS
		pis := tributos[0]
		// Base = 1000 (sem exclusao do ICMS)
		if pis.BaseCalculo.StringFixed(2) != "1000.00" {
			t.Errorf("PIS BaseCalculo = %s, want 1000.00", pis.BaseCalculo.StringFixed(2))
		}
		// Valor = 1000 * 1.65 / 100 = 16.5
		if pis.Valor.StringFixed(2) != "16.50" {
			t.Errorf("PIS Valor = %s, want 16.50", pis.Valor.StringFixed(2))
		}
		assertDetail(t, pis.MoreNumericDetails, "icms_excluido_base", "0")
		assertDetail(t, pis.MoreTextDetails, "metodo", "base_sem_exclusao_icms")

		// COFINS
		cofins := tributos[1]
		if cofins.BaseCalculo.StringFixed(2) != "1000.00" {
			t.Errorf("COFINS BaseCalculo = %s, want 1000.00", cofins.BaseCalculo.StringFixed(2))
		}
		// Valor = 1000 * 7.6 / 100 = 76
		if cofins.Valor.StringFixed(2) != "76.00" {
			t.Errorf("COFINS Valor = %s, want 76.00", cofins.Valor.StringFixed(2))
		}
		assertDetail(t, cofins.MoreTextDetails, "metodo", "base_sem_exclusao_icms")
	})
}

func TestPISCofinsCalculator_FallbackDefaults(t *testing.T) {
	t.Run("Quando repo retorna nil (regra nao encontrada), usa defaults", func(t *testing.T) {
		mock := newMockWithFederalRule(nil)
		calc := NewPISCofinsCalculator(mock)

		item := makeItem("01", "01", 1, 1000, 120)
		input := makeDocumentoFiscalEntrada(item)

		tributos, err := calc.Calculate(context.Background(), input)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		pis := tributos[0]
		// Default: ExcluiICMSBase = true, aliq PIS = 1.65
		// Base = 1000 - 120 = 880
		if pis.BaseCalculo.StringFixed(2) != "880.00" {
			t.Errorf("PIS BaseCalculo = %s, want 880.00", pis.BaseCalculo.StringFixed(2))
		}
		if pis.Valor.StringFixed(2) != "14.52" {
			t.Errorf("PIS Valor = %s, want 14.52", pis.Valor.StringFixed(2))
		}
		assertDetail(t, pis.MoreTextDetails, "fonte_aliquota", "default_hardcoded")
		assertDetail(t, pis.MoreTextDetails, "metodo", "base_com_exclusao_icms")
	})
}

func TestPISCofinsCalculator_FallbackOnError(t *testing.T) {
	t.Run("Quando repo retorna erro, usa defaults com fallback", func(t *testing.T) {
		mock := newMockWithFederalRuleError(errors.New("conexao com banco falhou"))
		calc := NewPISCofinsCalculator(mock)

		item := makeItem("01", "01", 2, 500, 180)
		input := makeDocumentoFiscalEntrada(item)

		tributos, err := calc.Calculate(context.Background(), input)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		pis := tributos[0]
		// Default: ExcluiICMSBase = true, aliq PIS = 1.65
		// Base = 1000 - 180 = 820
		if pis.BaseCalculo.StringFixed(2) != "820.00" {
			t.Errorf("PIS BaseCalculo = %s, want 820.00", pis.BaseCalculo.StringFixed(2))
		}
		assertDetail(t, pis.MoreTextDetails, "fonte_aliquota", "default_hardcoded")
	})
}

func TestPISCofinsCalculator_AliquotasCustomizadas(t *testing.T) {
	t.Run("Aliquotas customizadas do banco sobrescrevem defaults", func(t *testing.T) {
		mock := newMockWithFederalRule(defaultFederalRule(3.0, 9.25, true))
		calc := NewPISCofinsCalculator(mock)

		item := makeItem("02", "02", 1, 1000, 0)
		input := makeDocumentoFiscalEntrada(item)

		tributos, err := calc.Calculate(context.Background(), input)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		pis := tributos[0]
		if pis.Aliquota.StringFixed(2) != "3.00" {
			t.Errorf("PIS Aliquota = %s, want 3.00", pis.Aliquota.StringFixed(2))
		}
		if pis.Valor.StringFixed(2) != "30.00" {
			t.Errorf("PIS Valor = %s, want 30.00 (1000 * 3.0%%)", pis.Valor.StringFixed(2))
		}
		assertDetail(t, pis.MoreTextDetails, "fonte_aliquota", "banco_federal_tax_rules")

		cofins := tributos[1]
		if cofins.Aliquota.StringFixed(2) != "9.25" {
			t.Errorf("COFINS Aliquota = %s, want 9.25", cofins.Aliquota.StringFixed(2))
		}
		if cofins.Valor.StringFixed(2) != "92.50" {
			t.Errorf("COFINS Valor = %s, want 92.50 (1000 * 9.25%%)", cofins.Valor.StringFixed(2))
		}
	})
}

func TestPISCofinsCalculator_CST_03_PorUnidade(t *testing.T) {
	t.Run("CST 03: calculo por unidade (PIS e COFINS)", func(t *testing.T) {
		rule := defaultFederalRule(50, 200, true)
		rule.CSTPIS = "03"
		rule.CSTCOFINS = "03"
		mock := newMockWithFederalRule(rule)
		calc := NewPISCofinsCalculator(mock)

		item := makeItem("03", "03", 100, 0, 0)
		input := makeDocumentoFiscalEntrada(item)

		tributos, err := calc.Calculate(context.Background(), input)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		// PIS: CalcTax(100, 50) = 100 * 50 / 100 = 50.00
		pis := tributos[0]
		if pis.CST != "03" {
			t.Errorf("PIS CST = %s, want 03", pis.CST)
		}
		if pis.Valor.StringFixed(2) != "50.00" {
			t.Errorf("PIS Valor = %s, want 50.00", pis.Valor.StringFixed(2))
		}

		// COFINS: CalcTax(100, 200) = 100 * 200 / 100 = 200.00
		cofins := tributos[1]
		if cofins.CST != "03" {
			t.Errorf("COFINS CST = %s, want 03", cofins.CST)
		}
		if cofins.Valor.StringFixed(2) != "200.00" {
			t.Errorf("COFINS Valor = %s, want 200.00", cofins.Valor.StringFixed(2))
		}
	})
}

func TestPISCofinsCalculator_CST_ZeroValue(t *testing.T) {
	tests := []string{"04", "05", "06", "49", "50", "70", "98", "99"}

	for _, cst := range tests {
		t.Run("CST "+cst+": valor zero", func(t *testing.T) {
			rule := defaultFederalRule(1.65, 7.6, true)
			rule.CSTPIS = cst
			rule.CSTCOFINS = cst
			mock := newMockWithFederalRule(rule)
			calc := NewPISCofinsCalculator(mock)

			item := makeItem(cst, cst, 1, 1000, 0)
			input := makeDocumentoFiscalEntrada(item)

			tributos, err := calc.Calculate(context.Background(), input)
			if err != nil {
				t.Fatalf("erro inesperado: %v", err)
			}

			if len(tributos) != 2 {
				t.Fatalf("esperado 2 tributos para CST %s, obteve %d", cst, len(tributos))
			}

			pis := tributos[0]
			if !pis.Valor.Equal(decimal.Zero) {
				t.Errorf("CST %s PIS Valor = %s, want 0", cst, pis.Valor.String())
			}
			if pis.CST != cst {
				t.Errorf("CST %s PIS CST = %s, want %s", cst, pis.CST, cst)
			}

			cofins := tributos[1]
			if !cofins.Valor.Equal(decimal.Zero) {
				t.Errorf("CST %s COFINS Valor = %s, want 0", cst, cofins.Valor.String())
			}
		})
	}
}

func TestPISCofinsCalculator_MultiplosItens(t *testing.T) {
	t.Run("Multiplos itens com CSTs diferentes", func(t *testing.T) {
		mock := newMockWithFederalRule(defaultFederalRule(1.65, 7.6, true))
		calc := NewPISCofinsCalculator(mock)

		item1 := makeItem("01", "01", 1, 1000, 120)
		item2 := makeItem("04", "04", 2, 500, 0)
		input := makeDocumentoFiscalEntrada(item1, item2)

		tributos, err := calc.Calculate(context.Background(), input)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		if len(tributos) != 4 {
			t.Fatalf("esperado 4 tributos (2 itens x 2 tributos), obteve %d", len(tributos))
		}

		// Item 1: PIS CST 01 = 14.52
		if tributos[0].Valor.StringFixed(2) != "14.52" {
			t.Errorf("Item1 PIS = %s, want 14.52", tributos[0].Valor.StringFixed(2))
		}
		// Item 1: COFINS CST 01 = 66.88
		if tributos[1].Valor.StringFixed(2) != "66.88" {
			t.Errorf("Item1 COFINS = %s, want 66.88", tributos[1].Valor.StringFixed(2))
		}
		// Item 2: PIS CST 04 = 0
		if !tributos[2].Valor.Equal(decimal.Zero) {
			t.Errorf("Item2 PIS = %s, want 0", tributos[2].Valor.String())
		}
		// Item 2: COFINS CST 04 = 0
		if !tributos[3].Valor.Equal(decimal.Zero) {
			t.Errorf("Item2 COFINS = %s, want 0", tributos[3].Valor.String())
		}
	})
}

func TestPISCofinsCalculator_ICMSZeroNaBase(t *testing.T) {
	t.Run("Base negativa apos exclusao ICMS deve ser zero", func(t *testing.T) {
		mock := newMockWithFederalRule(defaultFederalRule(1.65, 7.6, true))
		calc := NewPISCofinsCalculator(mock)

		item := makeItem("01", "01", 1, 100, 200)
		input := makeDocumentoFiscalEntrada(item)

		tributos, err := calc.Calculate(context.Background(), input)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		// Base = 100 - 200 = negativo → zero
		pis := tributos[0]
		if pis.BaseCalculo.StringFixed(2) != "-100.00" {
			t.Errorf("PIS BaseCalculo = %s, want -100.00 (raw base)", pis.BaseCalculo.StringFixed(2))
		}
		// Valor = 0 * 1.65 / 100 = 0 (strategy clamps to zero)
		if !pis.Valor.Equal(decimal.Zero) {
			t.Errorf("PIS Valor = %s, want 0", pis.Valor.String())
		}
	})
}

func TestPISCofinsCalculator_SemCST_NaoCalcula(t *testing.T) {
	t.Run("Item sem CST PIS nao gera tributo PIS", func(t *testing.T) {
		mock := newMockWithFederalRule(defaultFederalRule(1.65, 7.6, true))
		calc := NewPISCofinsCalculator(mock)

		item := makeItem("", "01", 1, 1000, 0)
		input := makeDocumentoFiscalEntrada(item)

		tributos, err := calc.Calculate(context.Background(), input)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		if len(tributos) != 1 {
			t.Fatalf("esperado 1 tributo (apenas COFINS), obteve %d", len(tributos))
		}
		if tributos[0].Tributo != "COFINS" {
			t.Errorf("tributo = %s, want COFINS", tributos[0].Tributo)
		}
	})
}

func TestPISCofinsCalculator_AliquotaZeroNoBancoUsaDefault(t *testing.T) {
	t.Run("Aliquota zero no banco nao sobrescreve default", func(t *testing.T) {
		rule := &repository.FederalTaxRule{
			Id:               1,
			RegimeTributario: "3",
			CSTPIS:           "01",
			CSTCOFINS:        "01",
			AliquotaPIS:      decimal.Zero,
			AliquotaCOFINS:   decimal.Zero,
			ExcluiICMSBase:   false,
		}
		mock := newMockWithFederalRule(rule)
		calc := NewPISCofinsCalculator(mock)

		item := makeItem("01", "01", 1, 1000, 0)
		input := makeDocumentoFiscalEntrada(item)

		tributos, err := calc.Calculate(context.Background(), input)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		pis := tributos[0]
		// Usa default 1.65 (banco tem zero)
		if pis.Aliquota.StringFixed(2) != "1.65" {
			t.Errorf("PIS Aliquota = %s, want 1.65", pis.Aliquota.StringFixed(2))
		}
		if pis.Valor.StringFixed(2) != "16.50" {
			t.Errorf("PIS Valor = %s, want 16.50", pis.Valor.StringFixed(2))
		}

		cofins := tributos[1]
		if cofins.Aliquota.StringFixed(2) != "7.60" {
			t.Errorf("COFINS Aliquota = %s, want 7.60", cofins.Aliquota.StringFixed(2))
		}
	})
}

func TestPISCofinsCalculator_PISOnly(t *testing.T) {
	t.Run("Item com apenas CST PIS (sem COFINS)", func(t *testing.T) {
		mock := newMockWithFederalRule(defaultFederalRule(1.65, 7.6, true))
		calc := NewPISCofinsCalculator(mock)

		item := makeItem("01", "", 1, 1000, 0)
		input := makeDocumentoFiscalEntrada(item)

		tributos, err := calc.Calculate(context.Background(), input)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		if len(tributos) != 1 {
			t.Fatalf("esperado 1 tributo (apenas PIS), obteve %d", len(tributos))
		}
		if tributos[0].Tributo != "PIS" {
			t.Errorf("tributo = %s, want PIS", tributos[0].Tributo)
		}
	})
}

func TestPISCofinsCalculator_ExcluiICMSBaseZeroICMS(t *testing.T) {
	t.Run("ExcluiICMSBase=true mas ICMS zero: base mantida", func(t *testing.T) {
		mock := newMockWithFederalRule(defaultFederalRule(1.65, 7.6, true))
		calc := NewPISCofinsCalculator(mock)

		item := makeItem("01", "01", 3, 500, 0)
		input := makeDocumentoFiscalEntrada(item)

		tributos, err := calc.Calculate(context.Background(), input)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		pis := tributos[0]
		// Base = 1500 - 0 = 1500
		if pis.BaseCalculo.StringFixed(2) != "1500.00" {
			t.Errorf("PIS BaseCalculo = %s, want 1500.00", pis.BaseCalculo.StringFixed(2))
		}
		// Valor = 1500 * 1.65 / 100 = 24.75
		if pis.Valor.StringFixed(2) != "24.75" {
			t.Errorf("PIS Valor = %s, want 24.75", pis.Valor.StringFixed(2))
		}
		assertDetail(t, pis.MoreNumericDetails, "icms_excluido_base", "0")
	})
}

func assertDetail(t *testing.T, details []models.Detalhe, key, want string) {
	t.Helper()
	for _, d := range details {
		if d.Key == key {
			got := ""
			switch v := d.Value.(type) {
			case string:
				got = v
			case decimal.Decimal:
				got = v.String()
			}
			if got != want {
				t.Errorf("detalhe %q = %s, want %s", key, got, want)
			}
			return
		}
	}
	t.Errorf("detalhe %q nao encontrado", key)
}
