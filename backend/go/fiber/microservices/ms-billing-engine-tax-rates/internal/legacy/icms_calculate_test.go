package legacy

import (
	"context"
	"errors"
	"testing"
	"time"

	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"

	"github.com/shopspring/decimal"
)

func defaultICMSRule() *repository.ICMSRule {
	return &repository.ICMSRule{
		Id:                    1,
		UFOrigem:              "SP",
		UFDestino:             "SP",
		AliquotaInterna:       decimal.NewFromInt(18),
		AliquotaInterestadual: decimal.NewFromInt(12),
		ReducaoBase:           decimal.Zero,
		PercentualFCP:         decimal.Zero,
		InicioValidade:        time.Now().AddDate(-1, 0, 0),
	}
}

func defaultSimplesFaixa() *repository.SimplesFaixa {
	return &repository.SimplesFaixa{
		Id:             1,
		Anexo:          "I",
		AliqNominal:    decimal.NewFromFloat(4.0),
		ValorDeduzir:   decimal.Zero,
		PercIcmsAnexo:  decimal.NewFromFloat(33.5),
		InicioValidade: time.Now().AddDate(-1, 0, 0),
	}
}

func defaultEquivalence() *repository.TaxEquivalence {
	return &repository.TaxEquivalence{
		Id:             1,
		CSOSN:          "101",
		CSTEquivalente: "000",
		PermiteCredito: true,
		InicioValidade: time.Now().AddDate(-1, 0, 0),
	}
}

func icmsItem(ncm string, quantidade, valorUnitario float64, detalhes ...models.Detalhe) models.ItemDocumentoFiscalEntrada {
	return models.ItemDocumentoFiscalEntrada{
		SKU:                         "SKU-" + ncm,
		NCM:                         ncm,
		Quantidade:                  decimal.NewFromFloat(quantidade),
		ValorUnitario:               decimal.NewFromFloat(valorUnitario),
		DetalhesItemDocumentoFiscal: detalhes,
	}
}

func icmsDocumento(origemUF, destinoUF, crt string, isDestinoFinal bool, itens ...models.ItemDocumentoFiscalEntrada) models.DocumentoFiscalEntrada {
	return models.DocumentoFiscalEntrada{
		CRTEmitente:       crt,
		TipoOperacaoFiscal: "VENDA",
		LocalizacaoOrigem:  models.LocalizacaoFiscal{UF: origemUF},
		LocalizacaoDestino: models.LocalizacaoFiscal{UF: destinoUF},
		IsDestinoFinal:     isDestinoFinal,
		Itens:              itens,
		DataOperacao:       time.Now(),
	}
}

func TestICMS_ProprioInterno(t *testing.T) {
	mock := &mockTaxRepository{
		icmsRule: defaultICMSRule(),
	}
	calc := NewICMSCalculator(mock)

	doc := icmsDocumento("SP", "SP", "3", false, icmsItem("84713019", 1, 1000))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(res.Itens) != 1 || len(res.Itens[0].Tributos) != 1 {
		t.Fatalf("esperado 1 tributo, obteve %d", len(res.Itens[0].Tributos))
	}

	trib := res.Itens[0].Tributos[0]
	if trib.Tributo != "ICMS" {
		t.Errorf("tributo = %s, want ICMS", trib.Tributo)
	}
	// Valor: 1000 * 18% = 180.00
	if trib.Valor.StringFixed(2) != "180.00" {
		t.Errorf("valor = %s, want 180.00", trib.Valor.StringFixed(2))
	}
}

func TestICMS_ProprioInterno_ComReducaoBase(t *testing.T) {
	rule := defaultICMSRule()
	rule.ReducaoBase = decimal.NewFromInt(40) // 40% redução
	mock := &mockTaxRepository{icmsRule: rule}
	calc := NewICMSCalculator(mock)

	doc := icmsDocumento("SP", "SP", "3", false, icmsItem("84713019", 1, 1000))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	trib := res.Itens[0].Tributos[0]
	// Base: 1000 * (1 - 40%) = 600
	if trib.BaseCalculo.StringFixed(2) != "600.00" {
		t.Errorf("base = %s, want 600.00", trib.BaseCalculo.StringFixed(2))
	}
	// Valor: 600 * 18% = 108.00
	if trib.Valor.StringFixed(2) != "108.00" {
		t.Errorf("valor = %s, want 108.00", trib.Valor.StringFixed(2))
	}
}

func TestICMS_ProprioInterno_ComFCP(t *testing.T) {
	rule := defaultICMSRule()
	rule.PercentualFCP = decimal.NewFromInt(2) // 2% FCP
	mock := &mockTaxRepository{icmsRule: rule}
	calc := NewICMSCalculator(mock)

	doc := icmsDocumento("SP", "SP", "3", false, icmsItem("84713019", 1, 1000))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	trib := res.Itens[0].Tributos[0]
	// ICMS: 1000 * 18% = 180, FCP: 1000 * 2% = 20, Total = 200
	if trib.Valor.StringFixed(2) != "200.00" {
		t.Errorf("valor = %s, want 200.00 (180 ICMS + 20 FCP)", trib.Valor.StringFixed(2))
	}
}

func TestICMS_ST_Interno_CST010(t *testing.T) {
	rule := defaultICMSRule()
	mock := &mockTaxRepository{icmsRule: rule}
	calc := NewICMSCalculator(mock)

	doc := icmsDocumento("SP", "SP", "3", false, icmsItem("84713019", 1, 1000,
		models.Detalhe{Key: string(models.KeyDocumentoInfosItemSubstituirCSTICMS), Value: "010"},
		models.Detalhe{Key: string(models.KeyDocumentoInfosItemSubstituirMVAPercentual), Value: decimal.NewFromInt(40)},
	))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	trib := res.Itens[0].Tributos[0]
	if trib.Tributo != "ICMS_ST" {
		t.Errorf("tributo = %s, want ICMS_ST", trib.Tributo)
	}
	// Base ST: 1000 * (1 + 40/100) = 1400
	if trib.BaseCalculo.StringFixed(2) != "1400.00" {
		t.Errorf("base ST = %s, want 1400.00", trib.BaseCalculo.StringFixed(2))
	}
	// Valor ST: 1400 * 18% = 252.00
	if trib.Valor.StringFixed(2) != "252.00" {
		t.Errorf("valor = %s, want 252.00", trib.Valor.StringFixed(2))
	}
}

func TestICMS_ST_Interno_Protocolo(t *testing.T) {
	rule := defaultICMSRule()
	mock := &mockTaxRepository{
		icmsRule: rule,
		productException: &repository.ProductException{
			Id:                 99,
			NCM:                "84713019",
			PossuiProtocoloST:  true,
			MVAST:              decimal.NewFromInt(50),
			AliquotaInternaDestino: decimal.NewFromInt(18),
			InicioValidade:     time.Now().AddDate(-1, 0, 0),
		},
	}
	calc := NewICMSCalculator(mock)

	doc := icmsDocumento("SP", "SP", "3", false, icmsItem("84713019", 1, 1000))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	trib := res.Itens[0].Tributos[0]
	if trib.Tributo != "ICMS_ST" {
		t.Errorf("tributo = %s, want ICMS_ST", trib.Tributo)
	}
	// Base ST: 1000 * (1 + 50/100) = 1500
	if trib.BaseCalculo.StringFixed(2) != "1500.00" {
		t.Errorf("base ST = %s, want 1500.00", trib.BaseCalculo.StringFixed(2))
	}
	if trib.Valor.StringFixed(2) != "270.00" {
		t.Errorf("valor = %s, want 270.00 (1500 * 18%%)", trib.Valor.StringFixed(2))
	}
}

func TestICMS_Interestadual(t *testing.T) {
	rule := defaultICMSRule()
	rule.UFOrigem = "SP"
	rule.UFDestino = "RJ"
	mock := &mockTaxRepository{icmsRule: rule}
	calc := NewICMSCalculator(mock)

	doc := icmsDocumento("SP", "RJ", "3", false, icmsItem("84713019", 2, 500))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	trib := res.Itens[0].Tributos[0]
	if trib.Tributo != "ICMS" {
		t.Errorf("tributo = %s, want ICMS", trib.Tributo)
	}
	// Valor: 1000 * 12% = 120.00
	if trib.Valor.StringFixed(2) != "120.00" {
		t.Errorf("valor = %s, want 120.00", trib.Valor.StringFixed(2))
	}
}

func TestICMS_DIFAL(t *testing.T) {
	rule := defaultICMSRule()
	rule.UFOrigem = "SP"
	rule.UFDestino = "RJ"
	rule.AliquotaInterestadual = decimal.NewFromInt(12)
	rule.AliquotaInterna = decimal.NewFromInt(18)
	mock := &mockTaxRepository{icmsRule: rule}
	calc := NewICMSCalculator(mock)

	doc := icmsDocumento("SP", "RJ", "3", true, icmsItem("84713019", 1, 1000))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(res.Itens[0].Tributos) < 2 {
		t.Fatalf("esperado 2 tributos (ICMS + DIFAL), obteve %d", len(res.Itens[0].Tributos))
	}

	// ICMS Interestadual
	icms := res.Itens[0].Tributos[0]
	if icms.Tributo != "ICMS" {
		t.Errorf("tributo[0] = %s, want ICMS", icms.Tributo)
	}

	// DIFAL
	difal := res.Itens[0].Tributos[1]
	if difal.Tributo != "ICMS_DIFAL" {
		t.Errorf("tributo[1] = %s, want ICMS_DIFAL", difal.Tributo)
	}
	// DIFAL: 1000 * (18-12)% = 60.00
	if difal.Valor.StringFixed(2) != "60.00" {
		t.Errorf("DIFAL valor = %s, want 60.00", difal.Valor.StringFixed(2))
	}
}

func TestICMS_ST_Interestadual(t *testing.T) {
	rule := defaultICMSRule()
	rule.UFOrigem = "SP"
	rule.UFDestino = "RJ"
	rule.AliquotaInterestadual = decimal.NewFromInt(12)
	rule.AliquotaInterna = decimal.NewFromInt(18)
	mock := &mockTaxRepository{
		icmsRule: rule,
		productException: &repository.ProductException{
			Id:                 99,
			NCM:                "84713019",
			PossuiProtocoloST:  true,
			MVAST:              decimal.NewFromInt(35),
			AliquotaInternaDestino: decimal.NewFromInt(18),
			InicioValidade:     time.Now().AddDate(-1, 0, 0),
		},
	}
	calc := NewICMSCalculator(mock)

	doc := icmsDocumento("SP", "RJ", "3", false, icmsItem("84713019", 1, 1000))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(res.Itens[0].Tributos) < 2 {
		t.Fatalf("esperado 2 tributos (ICMS + ST), obteve %d", len(res.Itens[0].Tributos))
	}

	st := res.Itens[0].Tributos[1]
	if st.Tributo != "ICMS_ST" {
		t.Errorf("ST tributo = %s, want ICMS_ST", st.Tributo)
	}
	// Base ST: 1000 * (1 + 35/100) = 1350
	if st.BaseCalculo.StringFixed(2) != "1350.00" {
		t.Errorf("base ST = %s, want 1350.00", st.BaseCalculo.StringFixed(2))
	}
	// Valor ST: 1350 * 18% = 243.00
	if st.Valor.StringFixed(2) != "243.00" {
		t.Errorf("ST valor = %s, want 243.00", st.Valor.StringFixed(2))
	}
}

func TestICMS_SimplesNacional(t *testing.T) {
	mock := &mockTaxRepository{
		icmsRule:      defaultICMSRule(),
		equivalence:   defaultEquivalence(),
		simplesFaixa:  defaultSimplesFaixa(),
	}
	calc := NewICMSCalculator(mock)

	doc := icmsDocumento("SP", "SP", "SIMPLES", false, icmsItem("84713019", 1, 1000,
		models.Detalhe{Key: string(models.KeyDocumentoInfosItemCSOSN), Value: "101"},
		models.Detalhe{Key: string(models.KeyDocumentoInfosAnexoSimples), Value: "I"},
		models.Detalhe{Key: string(models.KeyDocumentoInfosRBT12), Value: decimal.NewFromFloat(180000)},
	))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(res.Itens[0].Tributos) != 1 {
		t.Fatalf("esperado 1 tributo Simples, obteve %d", len(res.Itens[0].Tributos))
	}

	trib := res.Itens[0].Tributos[0]
	if trib.Tributo != "ICMS_SIMPLES" {
		t.Errorf("tributo = %s, want ICMS_SIMPLES", trib.Tributo)
	}
}

func TestICMS_ConfigIndisponivel_PulaItem(t *testing.T) {
	mock := &mockTaxRepository{
		icmsRuleErr: errors.New("regra indisponivel"),
	}
	calc := NewICMSCalculator(mock)

	doc := icmsDocumento("XX", "YY", "3", false, icmsItem("84713019", 1, 500))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(res.Itens[0].Tributos) != 0 {
		t.Errorf("esperado 0 tributos (item pulado), obteve %d", len(res.Itens[0].Tributos))
	}
}

func TestICMS_ProductException_Merge(t *testing.T) {
	rule := defaultICMSRule()
	mock := &mockTaxRepository{
		icmsRule: rule,
		productException: &repository.ProductException{
			Id:                     99,
			NCM:                    "84713019",
			AliquotaInternaDestino: decimal.NewFromInt(7),
			CSTICMS:                "060",
			PossuiProtocoloST:      false,
			InicioValidade:         time.Now().AddDate(-1, 0, 0),
		},
	}
	calc := NewICMSCalculator(mock)

	doc := icmsDocumento("SP", "SP", "3", false, icmsItem("84713019", 1, 1000))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	trib := res.Itens[0].Tributos[0]
	// Aliquota da exceção (7%) prevalece sobre a regra (18%)
	if trib.Aliquota.StringFixed(2) != "7.00" {
		t.Errorf("aliquota = %s, want 7.00 (da excecao)", trib.Aliquota.StringFixed(2))
	}
	if trib.CST != "060" {
		t.Errorf("CST = %s, want 060 (da excecao)", trib.CST)
	}
	// Valor: 1000 * 7% = 70.00
	if trib.Valor.StringFixed(2) != "70.00" {
		t.Errorf("valor = %s, want 70.00", trib.Valor.StringFixed(2))
	}
}

func TestICMS_DIFAL_AliquotaInternaMenor_NaoAplica(t *testing.T) {
	rule := defaultICMSRule()
	rule.AliquotaInterna = decimal.NewFromInt(7)   // menor que interestadual
	rule.AliquotaInterestadual = decimal.NewFromInt(12)
	rule.UFOrigem = "SP"
	rule.UFDestino = "MG"
	mock := &mockTaxRepository{icmsRule: rule}
	calc := NewICMSCalculator(mock)

	doc := icmsDocumento("SP", "MG", "3", true, icmsItem("84713019", 1, 1000))
	res, err := calc.Calculate(context.Background(), doc)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	for _, trib := range res.Itens[0].Tributos {
		if trib.Tributo == "ICMS_DIFAL" {
			t.Error("DIFAL nao deveria ter sido aplicado (aliq interna < interestadual)")
		}
	}
}
