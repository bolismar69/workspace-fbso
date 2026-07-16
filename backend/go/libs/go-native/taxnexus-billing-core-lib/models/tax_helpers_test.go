package models

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestDetalhesNormalizadosConverteENormalizaChaves(t *testing.T) {
	detalhes := []Detalhe{
		{Key: "cst_pis", Value: "01"},
		{Key: "mva_percentual", Value: 17.5},
		{Key: "rbt12", Value: "1234.56"},
		{Key: "observacao", Value: true},
	}

	normalized := NewDetalhesNormalizados(detalhes)

	if got := GetString(normalized.Todos, KeyDocumentoInfosItemPISCOFINSCSTPIS); got != "01" {
		t.Fatalf("expected CST_PIS to be normalized, got %q", got)
	}

	if got := GetDecimal(normalized.Todos, "MVA_PERCENTUAL"); !got.Equal(decimal.RequireFromString("17.5")) {
		t.Fatalf("expected MVA decimal conversion, got %s", got)
	}

	if _, ok := normalized.Decimais["OBSERVACAO"]; ok {
		t.Fatalf("expected non-numeric detail to be absent from decimal map")
	}

	if got := normalized.Decimais["RBT12"]; !got.Equal(decimal.RequireFromString("1234.56")) {
		t.Fatalf("expected string number conversion to decimal, got %s", got)
	}
}

func TestPrepararMapasDetalhesMergeDocumentoEItem(t *testing.T) {
	documento := DocumentoFiscalEntrada{
		DetalhesDocumentoFiscal: []Detalhe{
			{Key: string(KeyDocumentoInfosRBT12), Value: "1000.00"},
			{Key: string(KeyDocumentoInfosIndicadorPresenca), Value: "2"},
		},
		Itens: []ItemDocumentoFiscalEntrada{
			{
				DetalhesItemDocumentoFiscal: []Detalhe{
					{Key: string(KeyDocumentoInfosRBT12), Value: "2000.00"},
					{Key: string(KeyDocumentoInfosItemSubstituirAliquotaPIS), Value: 1.65},
				},
			},
		},
	}

	documento.PrepararMapasDetalhes()

	if got := GetDecimal(documento.ToMap(), KeyDocumentoInfosRBT12); !got.Equal(decimal.RequireFromString("1000")) {
		t.Fatalf("expected document map to preserve document RBT12, got %s", got)
	}

	itemMap := documento.Itens[0].ToMap()
	if got := GetDecimal(itemMap, KeyDocumentoInfosRBT12); !got.Equal(decimal.RequireFromString("2000")) {
		t.Fatalf("expected item map to override document RBT12, got %s", got)
	}

	if got := GetString(itemMap, KeyDocumentoInfosIndicadorPresenca); got != "2" {
		t.Fatalf("expected merged document detail in item map, got %q", got)
	}

	itemDecimalMap := documento.Itens[0].ToDecimalMap()
	if got := itemDecimalMap[string(KeyDocumentoInfosItemSubstituirAliquotaPIS)]; !got.Equal(decimal.RequireFromString("1.65")) {
		t.Fatalf("expected decimal map to contain item aliquota pis, got %s", got)
	}
	if got := itemDecimalMap[string(KeyDocumentoInfosRBT12)]; !got.Equal(decimal.RequireFromString("2000")) {
		t.Fatalf("expected decimal map override for item RBT12, got %s", got)
	}
}

func TestAddDetalheAtualizaMapas(t *testing.T) {
	item := &ItemDocumentoFiscalEntrada{}
	item.AddDetalhe(KeyDocumentoInfosValorExclusaoICMS, "15.40")

	if got := GetDecimal(item.ToMap(), KeyDocumentoInfosValorExclusaoICMS); !got.Equal(decimal.RequireFromString("15.40")) {
		t.Fatalf("expected added detail to be available in generic map, got %s", got)
	}

	if got := item.ToDecimalMap()[string(KeyDocumentoInfosValorExclusaoICMS)]; !got.Equal(decimal.RequireFromString("15.40")) {
		t.Fatalf("expected added detail to be available in decimal map, got %s", got)
	}

	if len(item.DetalhesItemDocumentoFiscal) != 1 {
		t.Fatalf("expected raw details slice to be kept in sync, got %d entries", len(item.DetalhesItemDocumentoFiscal))
	}
}
