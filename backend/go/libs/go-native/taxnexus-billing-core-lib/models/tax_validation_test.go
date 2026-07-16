package models

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

// validDoc returns a minimal valid DocumentoFiscalEntrada for use as a test baseline.
func validDoc() DocumentoFiscalEntrada {
	return DocumentoFiscalEntrada{
		CorrelacaoID:       "corr-001",
		DocumentoID:        "doc-001",
		DataOperacao:       time.Date(2026, 4, 2, 10, 30, 0, 0, time.UTC),
		TipoOperacaoFiscal: "SAIDA",
		CRTEmitente:        "LUCRO_REAL",
		LocalizacaoOrigem:  LocalizacaoFiscal{UF: "SP"},
		LocalizacaoDestino: LocalizacaoFiscal{UF: "RJ"},
		Itens:              []ItemDocumentoFiscalEntrada{validItem()},
	}
}

func validItem() ItemDocumentoFiscalEntrada {
	return ItemDocumentoFiscalEntrada{
		SKU:           "SKU-001",
		NCM:           "22030000",
		Quantidade:    decimal.NewFromInt(1),
		ValorUnitario: decimal.NewFromFloat(10.50),
	}
}

func TestValidate_DocumentoValido(t *testing.T) {
	doc := validDoc()
	if errs := doc.Validate(); len(errs) != 0 {
		t.Errorf("expected no errors, got: %s", errs.Error())
	}
}

func TestValidate_CamposObrigatoriosAusentes(t *testing.T) {
	doc := DocumentoFiscalEntrada{}
	errs := doc.Validate()

	required := map[string]bool{
		"correlacao_id":          false,
		"documento_id":           false,
		"data_operacao":          false,
		"tipo_operacao_fiscal":   false,
		"crt_emitente":           false,
		"localizacao_origem.uf":  false,
		"localizacao_destino.uf": false,
		"itens":                  false,
	}
	for _, e := range errs {
		if _, ok := required[e.Field]; ok {
			required[e.Field] = true
		}
	}
	for field, found := range required {
		if !found {
			t.Errorf("expected REQUIRED error for field %q, but it was not reported", field)
		}
	}
}

func TestValidate_TipoOperacaoInvalido(t *testing.T) {
	doc := validDoc()
	doc.TipoOperacaoFiscal = "INVALIDO"
	errs := doc.Validate()
	if !hasFieldCode(errs, "tipo_operacao_fiscal", "INVALID_VALUE") {
		t.Error("expected INVALID_VALUE for tipo_operacao_fiscal")
	}
}

func TestValidate_TipoOperacaoCaseInsensitive(t *testing.T) {
	doc := validDoc()
	doc.TipoOperacaoFiscal = "saida" // lowercase — must be accepted after normalization
	if errs := doc.Validate(); len(errs) != 0 {
		t.Errorf("expected no errors for lowercase 'saida', got: %s", errs.Error())
	}
}

func TestValidate_CRTEmitenteInvalido(t *testing.T) {
	doc := validDoc()
	doc.CRTEmitente = "DESCONHECIDO"
	errs := doc.Validate()
	if !hasFieldCode(errs, "crt_emitente", "INVALID_VALUE") {
		t.Error("expected INVALID_VALUE for crt_emitente")
	}
}

func TestValidate_UFInvalida(t *testing.T) {
	doc := validDoc()
	doc.LocalizacaoOrigem.UF = "SPA" // 3 chars
	errs := doc.Validate()
	if !hasFieldCode(errs, "localizacao_origem.uf", "INVALID_FORMAT") {
		t.Error("expected INVALID_FORMAT for localizacao_origem.uf with 3-char UF")
	}
}

func TestValidate_IndicadorPresencaValido(t *testing.T) {
	doc := validDoc()
	for _, v := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"} {
		doc.IndicadorPresenca = v
		if errs := doc.Validate(); len(errs) != 0 {
			t.Errorf("expected no errors for indicador_presenca=%q, got: %s", v, errs.Error())
		}
	}
}

func TestValidate_IndicadorPresencaInvalido(t *testing.T) {
	doc := validDoc()
	doc.IndicadorPresenca = "X"
	errs := doc.Validate()
	if !hasFieldCode(errs, "indicador_presenca", "INVALID_VALUE") {
		t.Error("expected INVALID_VALUE for indicador_presenca='X'")
	}
}

func TestValidate_NaturezaOperacaoInvalida(t *testing.T) {
	doc := validDoc()
	doc.NaturezaOperacao = "EXPORTACAO" // not in the valid set
	errs := doc.Validate()
	if !hasFieldCode(errs, "natureza_operacao", "INVALID_VALUE") {
		t.Error("expected INVALID_VALUE for natureza_operacao='EXPORTACAO'")
	}
}

func TestValidate_ZonaEspecialValida(t *testing.T) {
	for _, v := range []string{"ZFM", "ALC", "zfm", "SIM", "TRUE", ""} {
		doc := validDoc()
		doc.DetalhesDocumentoFiscal = []Detalhe{{Key: string(KeyDocumentoInfosZonaEspecial), Value: v}}
		if errs := doc.Validate(); len(errs) != 0 {
			t.Errorf("expected no errors for ZONA_ESPECIAL=%q, got: %s", v, errs.Error())
		}
	}
}

func TestValidate_ZonaEspecialInvalida(t *testing.T) {
	doc := validDoc()
	doc.DetalhesDocumentoFiscal = []Detalhe{{Key: string(KeyDocumentoInfosZonaEspecial), Value: "INVALIDA"}}
	errs := doc.Validate()
	if !hasFieldCode(errs, "detalhes.ZONA_ESPECIAL", "INVALID_VALUE") {
		t.Error("expected INVALID_VALUE for ZONA_ESPECIAL='INVALIDA'")
	}
}

func TestValidate_ItemSemSKU(t *testing.T) {
	doc := validDoc()
	doc.Itens[0].SKU = ""
	errs := doc.Validate()
	if !hasFieldCode(errs, "itens[0].sku", "REQUIRED") {
		t.Error("expected REQUIRED for itens[0].sku")
	}
}

func TestValidate_ItemQuantidadeZero(t *testing.T) {
	doc := validDoc()
	doc.Itens[0].Quantidade = decimal.Zero
	errs := doc.Validate()
	if !hasFieldCode(errs, "itens[0].quantidade", "INVALID_VALUE") {
		t.Error("expected INVALID_VALUE for itens[0].quantidade=0")
	}
}

func TestValidate_ItemCSTPISInvalido(t *testing.T) {
	doc := validDoc()
	doc.Itens[0].DetalhesItemDocumentoFiscal = []Detalhe{
		{Key: string(KeyDocumentoInfosItemPISCOFINSCSTPIS), Value: "99X"},
	}
	errs := doc.Validate()
	expectedField := "itens[0].detalhes." + string(KeyDocumentoInfosItemPISCOFINSCSTPIS)
	if !hasFieldCode(errs, expectedField, "INVALID_VALUE") {
		t.Errorf("expected INVALID_VALUE for %s", expectedField)
	}
}

func TestValidate_ItemCSTPISValido(t *testing.T) {
	doc := validDoc()
	doc.Itens[0].DetalhesItemDocumentoFiscal = []Detalhe{
		{Key: string(KeyDocumentoInfosItemPISCOFINSCSTPIS), Value: "01"},
		{Key: string(KeyDocumentoInfosItemPISCOFINSCSTCOFINS), Value: "02"},
		{Key: string(KeyDocumentoInfosItemSubstituirCSTICMS), Value: "010"},
	}
	if errs := doc.Validate(); len(errs) != 0 {
		t.Errorf("expected no errors for valid CSTs, got: %s", errs.Error())
	}
}

func TestValidate_ItemCSTICMSInvalido(t *testing.T) {
	doc := validDoc()
	doc.Itens[0].DetalhesItemDocumentoFiscal = []Detalhe{
		{Key: string(KeyDocumentoInfosItemSubstituirCSTICMS), Value: "999"},
	}
	errs := doc.Validate()
	expectedField := "itens[0].detalhes." + string(KeyDocumentoInfosItemSubstituirCSTICMS)
	if !hasFieldCode(errs, expectedField, "INVALID_VALUE") {
		t.Errorf("expected INVALID_VALUE for %s", expectedField)
	}
}

func TestValidationErrors_Error(t *testing.T) {
	ve := ValidationErrors{
		{Field: "foo", Code: "REQUIRED", Message: "campo obrigatório"},
		{Field: "bar", Code: "INVALID_VALUE", Message: "valor inválido: \"X\""},
	}
	got := ve.Error()
	if got != `foo: campo obrigatório; bar: valor inválido: "X"` {
		t.Errorf("unexpected Error() output: %q", got)
	}
}

// hasFieldCode is a test helper that returns true if errs contains an entry
// matching the given field and code.
func hasFieldCode(errs ValidationErrors, field, code string) bool {
	for _, e := range errs {
		if e.Field == field && e.Code == code {
			return true
		}
	}
	return false
}
