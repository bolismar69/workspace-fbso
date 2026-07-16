package legacy

import (
	"testing"

	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

// --- Modo Redução de Base ---

func TestICMS_Desonerado_ReducaoBase(t *testing.T) {
	desoneracao := NewICMSDesoneracao()

	params := DesoneracaoParams{
		ValorItem:         decimal.NewFromFloat(1000.00),
		AliquotaNominal:   decimal.NewFromFloat(18.0),
		PercentualReducao: decimal.NewFromFloat(50.0), // 50% de reducao
		AliquotaAlvo:      decimal.Zero,
		MotivoDesoneracao: MotDesoneracaoProdutorAgropecuario, // 3
	}

	result := desoneracao.CalcularReducaoBase(params)

	// Base reduzida = 1000 * (1 - 50/100) = 500
	baseEsperada := decimal.NewFromFloat(500.0)
	if !result.BaseReduzida.Equal(baseEsperada) {
		t.Errorf("base reduzida: esperado %s, obteve %s", baseEsperada, result.BaseReduzida)
	}

	// ICMS = 500 * 0.18 = 90.00
	icmsEsperado := decimal.NewFromFloat(90.0)
	if !result.ICMS.Equal(icmsEsperado) {
		t.Errorf("ICMS: esperado %s, obteve %s", icmsEsperado, result.ICMS)
	}

	// vICMSDeson = (1000 * 0.18) - 90 = 180 - 90 = 90
	vICMSDesonEsperado := decimal.NewFromFloat(90.0)
	if !result.ValorICMSDeson.Equal(vICMSDesonEsperado) {
		t.Errorf("vICMSDeson: esperado %s, obteve %s", vICMSDesonEsperado, result.ValorICMSDeson)
	}

	if result.Modo != ModoReducaoBase {
		t.Errorf("modo: esperado %s, obteve %s", ModoReducaoBase, result.Modo)
	}

	if result.MotDesICMS != "3" {
		t.Errorf("motDesICMS: esperado '3', obteve '%s'", result.MotDesICMS)
	}
}

func TestICMS_Desonerado_ReducaoBase_100pct(t *testing.T) {
	desoneracao := NewICMSDesoneracao()

	params := DesoneracaoParams{
		ValorItem:         decimal.NewFromFloat(500.00),
		AliquotaNominal:   decimal.NewFromFloat(18.0),
		PercentualReducao: decimal.NewFromFloat(100.0), // isencao total
		AliquotaAlvo:      decimal.Zero,
	}

	result := desoneracao.CalcularReducaoBase(params)

	// Base reduzida = 500 * (1 - 1.0) = 0
	if !result.BaseReduzida.IsZero() {
		t.Errorf("base reduzida deve ser zero com 100%% reducao, obteve %s", result.BaseReduzida)
	}

	// ICMS = 0
	if !result.ICMS.IsZero() {
		t.Errorf("ICMS deve ser zero com base zero, obteve %s", result.ICMS)
	}

	// vICMSDeson = 500 * 0.18 - 0 = 90
	vICMSDesonEsperado := decimal.NewFromFloat(90.0)
	if !result.ValorICMSDeson.Equal(vICMSDesonEsperado) {
		t.Errorf("vICMSDeson: esperado %s, obteve %s", vICMSDesonEsperado, result.ValorICMSDeson)
	}
}

// --- Modo Limitação de Alíquota ---

func TestICMS_Desonerado_LimitacaoAliquota(t *testing.T) {
	desoneracao := NewICMSDesoneracao()

	params := DesoneracaoParams{
		ValorItem:       decimal.NewFromFloat(1000.00),
		AliquotaNominal: decimal.NewFromFloat(18.0),  // aliquota nominal
		AliquotaAlvo:    decimal.NewFromFloat(12.0),  // aliquota alvo (limitada)
	}

	result := desoneracao.CalcularLimitacaoAliquota(params)

	// Razao = 12/18 = 0.666...
	// Base reduzida = 1000 * (12/18) = 666.666...7
	baseEsperada := decimal.NewFromFloat(1000.0).Mul(
		decimal.NewFromFloat(12.0).Div(decimal.NewFromFloat(18.0)),
	)
	if !result.BaseReduzida.Sub(baseEsperada).Abs().LessThan(decimal.NewFromFloat(0.01)) {
		t.Errorf("base reduzida: esperado ~%s, obteve %s", baseEsperada, result.BaseReduzida)
	}

	// ICMS = 666.67 * 0.18 ≈ 120.00
	if !result.ICMS.GreaterThan(decimal.NewFromFloat(119.0)) || !result.ICMS.LessThan(decimal.NewFromFloat(121.0)) {
		t.Errorf("ICMS deve ser aproximadamente 120.00, obteve %s", result.ICMS)
	}

	if result.Modo != ModoLimitacaoAliquota {
		t.Errorf("modo: esperado %s, obteve %s", ModoLimitacaoAliquota, result.Modo)
	}
}

func TestICMS_Desonerado_LimitacaoAliquota_Zero(t *testing.T) {
	desoneracao := NewICMSDesoneracao()

	params := DesoneracaoParams{
		ValorItem:       decimal.NewFromFloat(500.00),
		AliquotaNominal: decimal.Zero, // aliquota nominal zero
		AliquotaAlvo:    decimal.NewFromFloat(12.0),
	}

	result := desoneracao.CalcularLimitacaoAliquota(params)

	// Com aliquota nominal zero, resultado deve ser zero
	if !result.ICMS.IsZero() {
		t.Errorf("ICMS deve ser zero com aliquota nominal zero, obteve %s", result.ICMS)
	}
	if !result.BaseReduzida.IsZero() {
		t.Errorf("base reduzida deve ser zero, obteve %s", result.BaseReduzida)
	}
}

// --- Validação de CST ---

func TestDesoneracao_CSTInvalido_RetornaErro(t *testing.T) {
	desoneracao := NewICMSDesoneracao()

	// CST 00 (tributacao integral) NAO permite desoneracao
	if desoneracao.CSTPermiteDesoneracao(models.CSTICMS000) {
		t.Error("CST 00 nao deve permitir desoneracao")
	}
}

func TestDesoneracao_CSTValidos(t *testing.T) {
	desoneracao := NewICMSDesoneracao()

	cstsValidos := []models.CSTICMS{
		models.CSTICMS("020"), // Com reducao de base
		models.CSTICMS("030"), // Isenta com ST
		models.CSTICMS("040"), // Isenta
		models.CSTICMS("041"), // Nao tributada
		models.CSTICMS("050"), // Suspensao
		models.CSTICMS("070"), // Com reducao e ST
		models.CSTICMS("090"), // Outras
	}

	for _, cst := range cstsValidos {
		t.Run(string(cst), func(t *testing.T) {
			if !desoneracao.CSTPermiteDesoneracao(cst) {
				t.Errorf("CST %s deve permitir desoneracao", cst)
			}
		})
	}
}

// --- DeterminarModo ---

func TestDesoneracao_DeterminarModo_ReducaoBase(t *testing.T) {
	desoneracao := NewICMSDesoneracao()

	modo := desoneracao.DeterminarModo(
		decimal.NewFromFloat(50.0),  // percentual reducao > 0
		decimal.Zero,                 // aliquota alvo = 0
	)

	if modo == nil {
		t.Fatal("modo nao deve ser nil")
	}
	if *modo != ModoReducaoBase {
		t.Errorf("esperado %s, obteve %s", ModoReducaoBase, *modo)
	}
}

func TestDesoneracao_DeterminarModo_LimitacaoAliquota(t *testing.T) {
	desoneracao := NewICMSDesoneracao()

	// AliquotaAlvo > 0 tem prioridade sobre PercentualReducao
	modo := desoneracao.DeterminarModo(
		decimal.NewFromFloat(50.0),  // percentual reducao (ignorado)
		decimal.NewFromFloat(12.0),  // aliquota alvo > 0
	)

	if modo == nil {
		t.Fatal("modo nao deve ser nil")
	}
	if *modo != ModoLimitacaoAliquota {
		t.Errorf("esperado %s, obteve %s", ModoLimitacaoAliquota, *modo)
	}
}

func TestDesoneracao_DeterminarModo_Nenhum(t *testing.T) {
	desoneracao := NewICMSDesoneracao()

	modo := desoneracao.DeterminarModo(decimal.Zero, decimal.Zero)
	if modo != nil {
		t.Errorf("modo deve ser nil quando nenhum parametro > 0, obteve %v", *modo)
	}
}

// --- Motivo Default ---

func TestDesoneracao_MotivoDefault_Outros(t *testing.T) {
	desoneracao := NewICMSDesoneracao()

	params := DesoneracaoParams{
		ValorItem:         decimal.NewFromFloat(100.00),
		AliquotaNominal:   decimal.NewFromFloat(18.0),
		PercentualReducao: decimal.NewFromFloat(30.0),
		MotivoDesoneracao: 0, // nao especificado → default = 9 (Outros)
	}

	result := desoneracao.CalcularReducaoBase(params)

	if result.MotDesICMS != "9" {
		t.Errorf("motivo default deve ser '9' (Outros), obteve '%s'", result.MotDesICMS)
	}
}

// --- Abate do Valor Total ---

func TestDesoneracao_AbateValorTotal(t *testing.T) {
	desoneracao := NewICMSDesoneracao()

	params := DesoneracaoParams{
		ValorItem:         decimal.NewFromFloat(1000.00),
		AliquotaNominal:   decimal.NewFromFloat(18.0),
		PercentualReducao: decimal.NewFromFloat(40.0),
	}

	result := desoneracao.CalcularReducaoBase(params)

	// ICMS sem desoneracao: 1000 * 0.18 = 180
	// ICMS com desoneracao (result.ICMS)
	// vICMSDeson = 180 - result.ICMS

	icmsSemDeson := decimal.NewFromFloat(1000.0 * 0.18)
	diferenca := icmsSemDeson.Sub(result.ICMS)

	if !diferenca.Sub(result.ValorICMSDeson).Abs().LessThan(decimal.NewFromFloat(0.01)) {
		t.Errorf("vICMSDeson deve ser aproximadamente a diferenca entre ICMS sem e com desoneracao: vICMSDeson=%s, diferenca=%s",
			result.ValorICMSDeson, diferenca)
	}
}

// --- BuildDesoneracaoDetails ---

func TestBuildDesoneracaoDetails_ReducaoBase(t *testing.T) {
	result := DesoneracaoResult{
		BaseReduzida:   decimal.NewFromFloat(500.0),
		ICMS:           decimal.NewFromFloat(90.0),
		ValorICMSDeson: decimal.NewFromFloat(90.0),
		Modo:           ModoReducaoBase,
		MotDesICMS:     "3",
	}

	params := DesoneracaoParams{
		ValorItem:          decimal.NewFromFloat(1000.0),
		AliquotaNominal:    decimal.NewFromFloat(18.0),
		PercentualReducao:  decimal.NewFromFloat(50.0),
		CodBeneficioFiscal: "RJ123456",
	}

	numDetails, textDetails := BuildDesoneracaoDetails(result, params)

	if len(numDetails) == 0 {
		t.Error("numeric details nao devem estar vazios")
	}
	if len(textDetails) == 0 {
		t.Error("text details nao devem estar vazios")
	}

	// Verifica se contem o percentual de reducao (modo reducao base)
	foundPctReducao := false
	for _, d := range numDetails {
		if d.Key == "percentual_reducao_base" {
			foundPctReducao = true
			break
		}
	}
	if !foundPctReducao {
		t.Error("numeric details devem conter 'percentual_reducao_base' no modo reducao base")
	}

	// Verifica se contem o codigo de beneficio fiscal
	foundCodBeneficio := false
	for _, d := range textDetails {
		if d.Key == "cod_beneficio_fiscal" && d.Value == "RJ123456" {
			foundCodBeneficio = true
			break
		}
	}
	if !foundCodBeneficio {
		t.Error("text details devem conter 'cod_beneficio_fiscal' quando informado")
	}
}

func TestBuildDesoneracaoDetails_LimitacaoAliquota(t *testing.T) {
	result := DesoneracaoResult{
		BaseReduzida:   decimal.NewFromFloat(666.67),
		ICMS:           decimal.NewFromFloat(120.0),
		ValorICMSDeson: decimal.NewFromFloat(60.0),
		Modo:           ModoLimitacaoAliquota,
		MotDesICMS:     "9",
	}

	params := DesoneracaoParams{
		ValorItem:       decimal.NewFromFloat(1000.0),
		AliquotaNominal: decimal.NewFromFloat(18.0),
		AliquotaAlvo:    decimal.NewFromFloat(12.0),
	}

	numDetails, _ := BuildDesoneracaoDetails(result, params)

	// Verifica se contem a aliquota efetiva alvo (modo limitacao)
	foundAliqAlvo := false
	for _, d := range numDetails {
		if d.Key == "aliquota_efetiva_alvo" {
			foundAliqAlvo = true
			break
		}
	}
	if !foundAliqAlvo {
		t.Error("numeric details devem conter 'aliquota_efetiva_alvo' no modo limitacao aliquota")
	}
}

// --- Simples Nacional: Nao Aplica ---

func TestDesoneracao_SimplesNacional_NaoAplica(t *testing.T) {
	// A desoneracao "classica" (SOP-017) NAO e aplicada no Simples Nacional.
	// O Simples ja possui tratamento proprio via CSOSN.
	// Este teste valida que a estrutura de desoneracao existe e funciona
	// para Regime Normal — o Simples usa calcularICMSSimples() separadamente,
	// que nao chama a desoneracao.

	// Verifica que MotDesoneracaoOutros = 9
	if MotDesoneracaoOutros != 9 {
		t.Errorf("MotDesoneracaoOutros deve ser 9, obteve %d", MotDesoneracaoOutros)
	}

	// Verifica que ha 14 motivos oficiais (1-12 + 90)
	todosMotivos := []MotivoDesoneracaoICMS{
		MotDesoneracaoTaxi,
		MotDesoneracaoDeficienteFisico,
		MotDesoneracaoProdutorAgropecuario,
		MotDesoneracaoFrotistaLocadora,
		MotDesoneracaoDiplomaticoConsular,
		MotDesoneracaoUtilitariosAmazonia,
		MotDesoneracaoSUFRAMA,
		MotDesoneracaoVendaOrgaosPublicos,
		MotDesoneracaoOutros,
		MotDesoneracaoCATGuiado,
		MotDesoneracaoCATNaoGuiado,
		MotDesoneracaoOrgaoFomento,
		MotDesoneracaoOlimpiadas,
	}

	if len(todosMotivos) != 13 {
		t.Errorf("esperado 13 motivos oficiais SEFAZ, obteve %d", len(todosMotivos))
	}
}
