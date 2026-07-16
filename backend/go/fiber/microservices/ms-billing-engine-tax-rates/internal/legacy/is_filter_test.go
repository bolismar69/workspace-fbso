// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/legacy/is_filter_test.go
package legacy

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"

	"github.com/shopspring/decimal"
)

// mockNCMSeletivoRepo é um mock especializado para testes do ISFilter
// que suporta múltiplos NCMs via um mapa.
type mockNCMSeletivoRepo struct {
	mockTaxRepository
	rules map[string]*repository.NCMSeletivoRule
	err   error
}

func newMockNCMSeletivoRepo(rule *repository.NCMSeletivoRule, err error) *mockNCMSeletivoRepo {
	m := &mockNCMSeletivoRepo{
		rules: make(map[string]*repository.NCMSeletivoRule),
		err:   err,
	}
	if rule != nil {
		m.rules[rule.NCM] = rule
	}
	return m
}

func newMockNCMSeletivoRepoMap(rules map[string]*repository.NCMSeletivoRule) *mockNCMSeletivoRepo {
	return &mockNCMSeletivoRepo{
		rules: rules,
	}
}

func (m *mockNCMSeletivoRepo) GetNCMSeletivo(ctx context.Context, ncm string) (*repository.NCMSeletivoRule, error) {
	if m.err != nil {
		return nil, m.err
	}
	rule, ok := m.rules[ncm]
	if !ok {
		return nil, nil // NCM não encontrado → IS não incide
	}
	return rule, nil
}

// itemComNCM cria um item com NCM e valor padrão para testes.
func itemComNCM(sku, ncm string, valorUnitario, quantidade float64) models.ItemDocumentoFiscalEntrada {
	return models.ItemDocumentoFiscalEntrada{
		SKU:           sku,
		NCM:           ncm,
		ValorUnitario: decimal.NewFromFloat(valorUnitario),
		Quantidade:    decimal.NewFromFloat(quantidade),
	}
}

// ncmSeletivoRuleCerveja retorna uma regra para bebidas alcoólicas.
func ncmSeletivoRuleCerveja() *repository.NCMSeletivoRule {
	return &repository.NCMSeletivoRule{
		NCM:            "22030000",
		Categoria:      "BEBIDAS_ALCOOLICAS",
		AliquotaIS:     decimal.NewFromFloat(50.0),
		Descricao:      sql.NullString{String: "Cervejas de malte", Valid: true},
		InicioValidade: time.Now().AddDate(-1, 0, 0),
	}
}

// TestIS_BebidaAlcoolica_Incide verifica que cerveja (NCM 22030000) paga IS 50%.
func TestIS_BebidaAlcoolica_Incide(t *testing.T) {
	repo := newMockNCMSeletivoRepo(ncmSeletivoRuleCerveja(), nil)
	filter := NewISFilter(repo)

	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemComNCM("CERV-001", "22030000", 10.00, 1),
		},
	}

	result, err := filter.Calculate(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("esperado 1 item, obtido %d", len(result))
	}

	tribs := result[0].Tributos
	if len(tribs) != 1 {
		t.Fatalf("esperado 1 tributo, obtido %d", len(tribs))
	}

	trib := tribs[0]
	if trib.Tributo != "IS" {
		t.Errorf("tributo = %q, esperado IS", trib.Tributo)
	}

	// 10.00 × 50% = 5.00
	valorEsperado := decimal.NewFromFloat(5.00)
	if !trib.Valor.Equal(valorEsperado) {
		t.Errorf("valor IS = %s, esperado %s", trib.Valor, valorEsperado)
	}

	if !trib.Aliquota.Equal(decimal.NewFromFloat(50.0)) {
		t.Errorf("aliquota IS = %s, esperado 50.0", trib.Aliquota)
	}

	// Verifica fonte nos detalhes
	fonteOk := false
	for _, d := range trib.MoreTextDetails {
		if d.Key == "fonte" && d.Value == "ncm_seletivo" {
			fonteOk = true
		}
		if d.Key == "categoria_is" && d.Value == "BEBIDAS_ALCOOLICAS" {
			// categoria correta
		}
	}
	if !fonteOk {
		t.Error("fonte 'ncm_seletivo' não encontrada nos detalhes")
	}
}

// TestIS_TelecomPuro_NaoIncide verifica que equipamento de rede (NCM 8517.62.59)
// NÃO paga IS (não está na tabela ncm_seletivo).
func TestIS_TelecomPuro_NaoIncide(t *testing.T) {
	repo := newMockNCMSeletivoRepo(nil, nil) // nil → NCM não encontrado
	filter := NewISFilter(repo)

	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemComNCM("TELECOM-001", "85176259", 5000.00, 1),
		},
	}

	result, err := filter.Calculate(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	trib := result[0].Tributos[0]
	if trib.Tributo != "IS" {
		t.Errorf("tributo = %q, esperado IS (mesmo com valor zero)", trib.Tributo)
	}
	if !trib.Valor.IsZero() {
		t.Errorf("telecom puro não deve pagar IS, mas valor = %s", trib.Valor)
	}

	// Verifica que está marcado como isento
	isExempt := false
	for _, d := range trib.MoreTextDetails {
		if d.Key == "is_exempt" && d.Value == "true" {
			isExempt = true
		}
	}
	if !isExempt {
		t.Error("NCM fora da tabela deveria ter is_exempt=true")
	}
}

// TestIS_IsentoFlag_Documento verifica que a flag isento_is no documento
// sobrescreve o IS para todos os itens.
func TestIS_IsentoFlag_Documento(t *testing.T) {
	repo := newMockNCMSeletivoRepo(ncmSeletivoRuleCerveja(), nil)
	filter := NewISFilter(repo)

	input := models.DocumentoFiscalEntrada{
		DetalhesDocumentoFiscal: []models.Detalhe{
			{Key: "ISENTO_IS", Value: "true"},
		},
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemComNCM("CERV-001", "22030000", 10.00, 1),
			itemComNCM("CERV-002", "22030000", 20.00, 2),
		},
	}

	result, err := filter.Calculate(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	for i, item := range result {
		for _, trib := range item.Tributos {
			if !trib.Valor.IsZero() {
				t.Errorf("item %d: isento_is no doc deveria zerar IS, mas valor = %s", i, trib.Valor)
			}
			isExempt := false
			for _, d := range trib.MoreTextDetails {
				if d.Key == "is_exempt" && d.Value == "true" {
					isExempt = true
				}
			}
			if !isExempt {
				t.Errorf("item %d: deveria ter is_exempt=true", i)
			}
		}
	}
}

// TestIS_IsentoFlag_Item verifica que a flag isento_is no item sobrescreve
// apenas aquele item específico.
func TestIS_IsentoFlag_Item(t *testing.T) {
	repo := newMockNCMSeletivoRepo(ncmSeletivoRuleCerveja(), nil)
	filter := NewISFilter(repo)

	// Item 1: normal, deve pagar IS
	item1 := itemComNCM("CERV-001", "22030000", 10.00, 1)
	// Item 2: isento manualmente
	item2 := itemComNCM("CERV-002", "22030000", 10.00, 1)
	item2.AddDetalhe(models.KeyDocumentoInfos("ISENTO_IS"), "true")

	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{item1, item2},
	}

	result, err := filter.Calculate(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("esperado 2 itens, obtido %d", len(result))
	}

	// Item 1: deve pagar IS
	trib1 := result[0].Tributos[0]
	if trib1.Valor.IsZero() {
		t.Error("item 1 deveria pagar IS (sem flag isento)")
	}

	// Item 2: isento
	trib2 := result[1].Tributos[0]
	if !trib2.Valor.IsZero() {
		t.Error("item 2 não deveria pagar IS (flag isento_is)")
	}
}

// TestIS_Cigarros_Aliquota100 verifica que cigarros pagam IS 100%.
func TestIS_Cigarros_Aliquota100(t *testing.T) {
	rule := &repository.NCMSeletivoRule{
		NCM:            "24022000",
		Categoria:      "CIGARROS",
		AliquotaIS:     decimal.NewFromFloat(100.0),
		Descricao:      sql.NullString{String: "Cigarros com filtro", Valid: true},
		InicioValidade: time.Now().AddDate(-1, 0, 0),
	}
	repo := newMockNCMSeletivoRepo(rule, nil)
	filter := NewISFilter(repo)

	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemComNCM("CIG-001", "24022000", 15.00, 1),
		},
	}

	result, _ := filter.Calculate(context.Background(), input)

	// 15.00 × 100% = 15.00
	valorEsperado := decimal.NewFromFloat(15.00)
	if !result[0].Tributos[0].Valor.Equal(valorEsperado) {
		t.Errorf("cigarros IS 100%%: esperado %s, obtido %s",
			valorEsperado, result[0].Tributos[0].Valor)
	}
}

// TestIS_Refrigerantes_Aliquota25 verifica que refrigerantes pagam IS 25%.
func TestIS_Refrigerantes_Aliquota25(t *testing.T) {
	rule := &repository.NCMSeletivoRule{
		NCM:            "22021000",
		Categoria:      "REFRIGERANTES",
		AliquotaIS:     decimal.NewFromFloat(25.0),
		InicioValidade: time.Now().AddDate(-1, 0, 0),
	}
	repo := newMockNCMSeletivoRepo(rule, nil)
	filter := NewISFilter(repo)

	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemComNCM("REFRI-001", "22021000", 8.00, 1),
		},
	}

	result, _ := filter.Calculate(context.Background(), input)

	// 8.00 × 25% = 2.00
	valorEsperado := decimal.NewFromFloat(2.00)
	if !result[0].Tributos[0].Valor.Equal(valorEsperado) {
		t.Errorf("refrigerantes IS 25%%: esperado %s, obtido %s",
			valorEsperado, result[0].Tributos[0].Valor)
	}
}

// TestIS_MultiplosItensMistos verifica IS com itens que pagam e não pagam.
func TestIS_MultiplosItensMistos(t *testing.T) {
	repo := newMockNCMSeletivoRepo(ncmSeletivoRuleCerveja(), nil)
	filter := NewISFilter(repo)

	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemComNCM("CERV-001", "22030000", 10.00, 1),  // paga IS
			itemComNCM("TELECOM-001", "85176259", 5000.00, 1), // não paga
			itemComNCM("CERV-002", "22030000", 20.00, 2),  // paga IS
		},
	}

	result, err := filter.Calculate(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(result) != 3 {
		t.Fatalf("esperado 3 itens, obtido %d", len(result))
	}

	// Item 0: cerveja — IS = 10 × 50% = 5.00
	if result[0].Tributos[0].Valor.IsZero() {
		t.Error("item 0 (cerveja) deveria pagar IS")
	}

	// Item 1: telecom — IS = 0
	if !result[1].Tributos[0].Valor.IsZero() {
		t.Error("item 1 (telecom) não deveria pagar IS")
	}

	// Item 2: cerveja — IS = 40 × 50% = 20.00
	valorEsperado := decimal.NewFromFloat(20.00)
	if !result[2].Tributos[0].Valor.Equal(valorEsperado) {
		t.Errorf("item 2 IS = %s, esperado %s", result[2].Tributos[0].Valor, valorEsperado)
	}
}

// TestIS_ErroRepositorio_FallbackZero verifica que em caso de erro no repositório,
// o IS assume valor zero (fail-safe) e registra auditoria.
func TestIS_ErroRepositorio_FallbackZero(t *testing.T) {
	repo := newMockNCMSeletivoRepo(nil, context.DeadlineExceeded)
	filter := NewISFilter(repo)

	input := models.DocumentoFiscalEntrada{
		Itens: []models.ItemDocumentoFiscalEntrada{
			itemComNCM("CERV-001", "22030000", 10.00, 1),
		},
	}

	result, err := filter.Calculate(context.Background(), input)
	if err != nil {
		t.Fatalf("ISFilter não deve propagar erro do repo, mas propagou: %v", err)
	}

	if len(result) != 1 || len(result[0].Tributos) != 1 {
		t.Fatal("deveria retornar resultado mesmo com erro no repositório")
	}

	if !result[0].Tributos[0].Valor.IsZero() {
		t.Error("erro no repositório deveria resultar em IS=0 (fail-safe)")
	}
}
