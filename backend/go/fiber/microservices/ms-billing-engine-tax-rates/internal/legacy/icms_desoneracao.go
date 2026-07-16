package legacy

import (
	"log/slog"

	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

// MotivoDesoneracaoICMS representa os codigos oficiais SEFAZ para motivo
// de desoneracao de ICMS (campos motDesICMS / vICMSDeson da NF-e).
//
// Ref: Ajuste SINIEF 07/2005, Nota Tecnica 2019.001 v1.51
type MotivoDesoneracaoICMS int

const (
	MotDesoneracaoTaxi                   MotivoDesoneracaoICMS = 1  // Taxi
	MotDesoneracaoDeficienteFisico       MotivoDesoneracaoICMS = 2  // Deficiente fisico
	MotDesoneracaoProdutorAgropecuario   MotivoDesoneracaoICMS = 3  // Produtor Agropecuario
	MotDesoneracaoFrotistaLocadora       MotivoDesoneracaoICMS = 4  // Frotista/Locadora
	MotDesoneracaoDiplomaticoConsular    MotivoDesoneracaoICMS = 5  // Diplomatico/Consular
	MotDesoneracaoUtilitariosAmazonia    MotivoDesoneracaoICMS = 6  // Utilitarios e Motocicletas Amazonia
	MotDesoneracaoSUFRAMA                MotivoDesoneracaoICMS = 7  // SUFRAMA
	MotDesoneracaoVendaOrgaosPublicos    MotivoDesoneracaoICMS = 8  // Venda a Orgaos Publicos
	MotDesoneracaoOutros                 MotivoDesoneracaoICMS = 9  // Outros (default)
	MotDesoneracaoCATGuiado              MotivoDesoneracaoICMS = 10 // Deficiente Condutor (CAT Guiado)
	MotDesoneracaoCATNaoGuiado           MotivoDesoneracaoICMS = 11 // Deficiente Nao Condutor (CAT Nao Guiado)
	MotDesoneracaoOrgaoFomento           MotivoDesoneracaoICMS = 12 // Orgao de Fomento e Desenvolvimento
	MotDesoneracaoOlimpiadas             MotivoDesoneracaoICMS = 90 // Olimpiadas (Rio 2016) — historico
)

// String retorna o codigo SEFAZ como string.
func (m MotivoDesoneracaoICMS) String() string {
	return decimal.NewFromInt(int64(m)).String()
}

// cstsPermitemDesoneracao define os CSTs de ICMS que permitem desoneracao.
// CST 00 (tributacao integral) NAO permite desoneracao (BR-TAX-CONS-013).
//
// CSTs que permitem:
//
//	20 — Com reducao de base
//	30 — Isenta / nao tributada e com cobranca do ICMS por substituicao tributaria
//	40 — Isenta
//	41 — Nao tributada
//	50 — Suspensao
//	70 — Com reducao de base de calculo e cobranca do ICMS por substituicao tributaria
//	90 — Outras
var cstsPermitemDesoneracao = map[models.CSTICMS]bool{
	models.CSTICMS("020"): true,
	models.CSTICMS("030"): true,
	models.CSTICMS("040"): true,
	models.CSTICMS("041"): true,
	models.CSTICMS("050"): true,
	models.CSTICMS("070"): true,
	models.CSTICMS("090"): true,
}

// ModoDesoneracao define o tipo de calculo da desoneracao.
type ModoDesoneracao string

const (
	// ModoReducaoBase: Base_Reduzida = Valor × (1 − PctReducao/100)
	// ICMS = Base_Reduzida × Aliquota/100
	// vICMSDeson = (Valor × Aliquota) − ICMS
	ModoReducaoBase ModoDesoneracao = "REDUCAO_BASE"

	// ModoLimitacaoAliquota: Indice = 1 − (AliqAlvo/AliqNominal)
	// Base_Reduzida = Valor × (AliqAlvo/AliqNominal)
	// ICMS = Base_Reduzida × Aliquota/100
	ModoLimitacaoAliquota ModoDesoneracao = "LIMITACAO_ALIQUOTA"
)

// DesoneracaoParams contem os parametros de entrada para o calculo de desoneracao.
type DesoneracaoParams struct {
	ValorItem          decimal.Decimal // valor total do item (quantidade × unitario)
	AliquotaNominal    decimal.Decimal // aliquota nominal de ICMS (ex: 18%)
	PercentualReducao  decimal.Decimal // percentual de reducao da base (ex: 50% → 50)
	AliquotaAlvo       decimal.Decimal // aliquota efetiva alvo para limitacao (ex: 12%)
	MotivoDesoneracao  MotivoDesoneracaoICMS
	CodBeneficioFiscal string // codigo do beneficio fiscal (ex: "RJ123456")
}

// DesoneracaoResult contem o resultado do calculo de desoneracao.
type DesoneracaoResult struct {
	BaseReduzida     decimal.Decimal // base de calculo apos reducao
	ICMS             decimal.Decimal // valor do ICMS calculado
	ValorICMSDeson   decimal.Decimal // valor desonerado (abatido do total)
	Modo             ModoDesoneracao
	MotDesICMS       string // codigo SEFAZ do motivo (1-12, 90)
}

// ICMSDesoneracao implementa as regras de ICMS Desonerado conforme
// SOP-017, BR-TAX-CALC-021 (Reducao de Base) e BR-TAX-CALC-022
// (Limitacao de Aliquota).
type ICMSDesoneracao struct{}

// NewICMSDesoneracao cria uma nova instancia da calculadora de desoneracao.
func NewICMSDesoneracao() *ICMSDesoneracao {
	return &ICMSDesoneracao{}
}

// CSTPermiteDesoneracao verifica se um CST de ICMS permite desoneracao.
// CST 00 (tributacao integral) NAO permite. (BR-TAX-CONS-013)
func (d *ICMSDesoneracao) CSTPermiteDesoneracao(cst models.CSTICMS) bool {
	return cstsPermitemDesoneracao[cst]
}

// CalcularReducaoBase implementa o Modo Reducao de Base (BR-TAX-CALC-021).
//
// Formula:
//
//	Base_Reduzida = Valor × (1 − PctReducao/100)
//	ICMS = Base_Reduzida × Aliquota/100
//	vICMSDeson = (Valor × Aliquota) − ICMS
func (d *ICMSDesoneracao) CalcularReducaoBase(params DesoneracaoParams) DesoneracaoResult {
	fatorReducao := decimal.NewFromInt(1).Sub(
		params.PercentualReducao.Div(decimal.NewFromInt(100)),
	)
	baseReduzida := params.ValorItem.Mul(fatorReducao)

	aliqPct := params.AliquotaNominal.Div(decimal.NewFromInt(100))
	icms := baseReduzida.Mul(aliqPct).Round(2)

	icmsSemReducao := params.ValorItem.Mul(aliqPct)
	vICMSDeson := icmsSemReducao.Sub(icms).Round(2)

	motDesICMS := params.MotivoDesoneracao
	if motDesICMS == 0 {
		motDesICMS = MotDesoneracaoOutros // default = 9 (Outros)
	}

	slog.Debug("ICMS Desonerado — Reducao de Base",
		"valor_item", params.ValorItem,
		"pct_reducao", params.PercentualReducao,
		"base_reduzida", baseReduzida,
		"icms", icms,
		"vICMSDeson", vICMSDeson,
		"motivo", motDesICMS,
	)

	return DesoneracaoResult{
		BaseReduzida:   baseReduzida,
		ICMS:           icms,
		ValorICMSDeson: vICMSDeson,
		Modo:           ModoReducaoBase,
		MotDesICMS:     motDesICMS.String(),
	}
}

// CalcularLimitacaoAliquota implementa o Modo Limitacao de Aliquota
// (BR-TAX-CALC-022).
//
// Formula:
//
//	Indice = 1 − (AliqAlvo/AliqNominal)
//	Base_Reduzida = Valor × (AliqAlvo/AliqNominal)
//	ICMS = Base_Reduzida × AliquotaNominal/100
//	vICMSDeson = (Valor × AliquotaNominal) − ICMS
func (d *ICMSDesoneracao) CalcularLimitacaoAliquota(params DesoneracaoParams) DesoneracaoResult {
	if params.AliquotaNominal.IsZero() {
		slog.Warn("Aliquota nominal zero na limitacao de aliquota — retornando zero")
		return DesoneracaoResult{
			MotDesICMS: MotDesoneracaoOutros.String(),
			Modo:       ModoLimitacaoAliquota,
		}
	}

	razao := params.AliquotaAlvo.Div(params.AliquotaNominal)
	baseReduzida := params.ValorItem.Mul(razao)

	aliqPct := params.AliquotaNominal.Div(decimal.NewFromInt(100))
	icms := baseReduzida.Mul(aliqPct).Round(2)

	icmsSemLimitacao := params.ValorItem.Mul(aliqPct)
	vICMSDeson := icmsSemLimitacao.Sub(icms).Round(2)

	motDesICMS := params.MotivoDesoneracao
	if motDesICMS == 0 {
		motDesICMS = MotDesoneracaoOutros // default = 9 (Outros)
	}

	slog.Debug("ICMS Desonerado — Limitacao de Aliquota",
		"valor_item", params.ValorItem,
		"aliq_nominal", params.AliquotaNominal,
		"aliq_alvo", params.AliquotaAlvo,
		"indice", decimal.NewFromInt(1).Sub(razao),
		"base_reduzida", baseReduzida,
		"icms", icms,
		"vICMSDeson", vICMSDeson,
		"motivo", motDesICMS,
	)

	return DesoneracaoResult{
		BaseReduzida:   baseReduzida,
		ICMS:           icms,
		ValorICMSDeson: vICMSDeson,
		Modo:           ModoLimitacaoAliquota,
		MotDesICMS:     motDesICMS.String(),
	}
}

// DeterminarModo identifica qual modo de desoneracao aplicar com base nos
// parametros do item.
//
// Prioridade:
//  1. Se AliquotaAlvo > 0 → Modo Limitacao de Aliquota
//  2. Se PercentualReducao > 0 → Modo Reducao de Base
//  3. Senao → sem desoneracao (retorna nil)
func (d *ICMSDesoneracao) DeterminarModo(percentualReducao, aliquotaAlvo decimal.Decimal) *ModoDesoneracao {
	if aliquotaAlvo.GreaterThan(decimal.Zero) {
		modo := ModoLimitacaoAliquota
		return &modo
	}
	if percentualReducao.GreaterThan(decimal.Zero) {
		modo := ModoReducaoBase
		return &modo
	}
	return nil
}

// BuildDesoneracaoDetails constroi os detalhes de auditoria para tributos
// com desoneracao (SOP-017 Step 6).
func BuildDesoneracaoDetails(result DesoneracaoResult, params DesoneracaoParams) (numericDetails []models.Detalhe, textDetails []models.Detalhe) {
	numericDetails = []models.Detalhe{
		{Key: "valor_item", Value: params.ValorItem},
		{Key: "base_reduzida", Value: result.BaseReduzida},
		{Key: "icms_calculado", Value: result.ICMS},
		{Key: "vicms_deson", Value: result.ValorICMSDeson},
		{Key: "aliquota_nominal", Value: params.AliquotaNominal},
	}

	textDetails = []models.Detalhe{
		{Key: "modo_desoneracao", Value: string(result.Modo)},
		{Key: "motivo_desoneracao_icms", Value: result.MotDesICMS},
	}

	if result.Modo == ModoReducaoBase {
		numericDetails = append(numericDetails,
			models.Detalhe{Key: "percentual_reducao_base", Value: params.PercentualReducao},
		)
	} else if result.Modo == ModoLimitacaoAliquota {
		numericDetails = append(numericDetails,
			models.Detalhe{Key: "aliquota_efetiva_alvo", Value: params.AliquotaAlvo},
		)
	}

	if params.CodBeneficioFiscal != "" {
		textDetails = append(textDetails,
			models.Detalhe{Key: "cod_beneficio_fiscal", Value: params.CodBeneficioFiscal},
		)
	}

	return numericDetails, textDetails
}
