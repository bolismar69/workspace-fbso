// Package phase implementa o sistema de resolucao de fase tributaria da Reforma
// Tributaria Brasileira (EC 132/2023, LC 214/2025).
//
// O PhaseResolver mapeia a DataOperacao do documento fiscal para a fase
// correspondente do cronograma de transicao:
//
//	2026              → SHADOW_RUN          (CBS/IBS calculados, mas nao compoe total a pagar)
//	2027–2028         → CBS_PLENA           (CBS substitui PIS/COFINS; IBS ainda shadow)
//	2029–2032         → TRANSICAO_SUBNACIONAL (ICMS/ISS gradual → IBS subnacional)
//	2033+             → IVA_DUAL            (CBS + IBS plenos; PIS/COFINS/ICMS/ISS extintos)
//
// O TaxSelector aplica a matriz DT-001 para determinar quais calculadoras
// estao ativas em cada fase, considerando a natureza da operacao.
//
// Referencias:
//   - POLICE-FIN-00001 §5.1, §5.2, §5.3
//   - PROCEDURE-FIN-00001 SOP-013
//   - RULES-CATALOG-FIN-00001 BR-TAX-INF-001 a 004, DT-001, DT-005,
//     BR-TAX-ACT-005, BR-TAX-ACT-006
package phase

import (
	"time"
)

// Phase representa uma fase do cronograma de transicao da Reforma Tributaria.
type Phase string

const (
	// PhaseShadowRun (2026): CBS e IBS sao calculados em regime de "sombra"
	// (shadow tax result), mas NAO compoem o total_a_pagar. Util para testes,
	// homologacao e adaptacao de sistemas antes da vigencia plena.
	PhaseShadowRun Phase = "SHADOW_RUN"

	// PhaseCBSPlena (2027–2028): CBS substitui definitivamente PIS e COFINS.
	// IBS ainda opera em shadow run (nao compoe total). PIS/COFINS = 0.
	PhaseCBSPlena Phase = "CBS_PLENA"

	// PhaseTransicaoSubnacional (2029–2032): ICMS e ISS comecam a ser
	// gradualmente substituidos pelo IBS subnacional. Aplica fator de
	// reducao progressivo conforme cronograma oficial.
	PhaseTransicaoSubnacional Phase = "TRANSICAO_SUBNACIONAL"

	// PhaseIVADual (2033+): Sistema IVA Dual pleno. CBS (federal) + IBS
	// (subnacional) sao os unicos tributos sobre consumo. PIS, COFINS,
	// ICMS e ISS sao extintos (valor zero).
	PhaseIVADual Phase = "IVA_DUAL"
)

// PhaseInfo contem metadados sobre uma fase tributaria.
type PhaseInfo struct {
	Phase           Phase
	Description     string
	CBSActive       bool // CBS compoe total a pagar?
	IBSActive       bool // IBS compoe total a pagar?
	PISCOFINSActive bool // PIS/COFINS ainda vigentes?
	ICMSActive      bool // ICMS ainda vigente?
	ISSActive       bool // ISS ainda vigente?
	IPIActive       bool // IPI sempre ativo (tributo extrafiscal, nao substituido)
}

// phaseRegistry mapeia cada fase as suas caracteristicas conforme DT-001.
var phaseRegistry = map[Phase]PhaseInfo{
	PhaseShadowRun: {
		Phase:           PhaseShadowRun,
		Description:     "Shadow Run — CBS/IBS calculados, mas nao compoem total a pagar (2026)",
		CBSActive:       false, // shadow apenas
		IBSActive:       false, // shadow apenas
		PISCOFINSActive: true,
		ICMSActive:      true,
		ISSActive:       true,
		IPIActive:       true,
	},
	PhaseCBSPlena: {
		Phase:           PhaseCBSPlena,
		Description:     "CBS Plena — CBS substitui PIS/COFINS; IBS em shadow (2027–2028)",
		CBSActive:       true,
		IBSActive:       false, // shadow apenas
		PISCOFINSActive: false,
		ICMSActive:      true,
		ISSActive:       true,
		IPIActive:       true,
	},
	PhaseTransicaoSubnacional: {
		Phase:           PhaseTransicaoSubnacional,
		Description:     "Transicao Subnacional — ICMS/ISS gradativamente substituidos pelo IBS (2029–2032)",
		CBSActive:       true,
		IBSActive:       true,
		PISCOFINSActive: false,
		ICMSActive:      true, // com fator de reducao
		ISSActive:       true, // com fator de reducao
		IPIActive:       true,
	},
	PhaseIVADual: {
		Phase:           PhaseIVADual,
		Description:     "IVA Dual — CBS + IBS plenos; PIS/COFINS/ICMS/ISS extintos (2033+)",
		CBSActive:       true,
		IBSActive:       true,
		PISCOFINSActive: false,
		ICMSActive:      false,
		ISSActive:       false,
		IPIActive:       true,
	},
}

// PhaseResolver resolve a fase tributaria com base na data da operacao.
// Implementa BR-TAX-INF-001 a 004.
type PhaseResolver struct{}

// NewPhaseResolver cria um novo resolver de fase tributaria.
func NewPhaseResolver() *PhaseResolver {
	return &PhaseResolver{}
}

// Resolve determina a Phase correspondente a uma data de operacao.
//
// Regras (BR-TAX-INF-001 a 004):
//   - 2026: SHADOW_RUN
//   - 2027–2028: CBS_PLENA
//   - 2029–2032: TRANSICAO_SUBNACIONAL
//   - 2033+: IVA_DUAL
func (r *PhaseResolver) Resolve(dataOperacao time.Time) Phase {
	year := dataOperacao.Year()

	switch {
	case year >= 2033:
		return PhaseIVADual
	case year >= 2029 && year <= 2032:
		return PhaseTransicaoSubnacional
	case year >= 2027 && year <= 2028:
		return PhaseCBSPlena
	case year == 2026:
		return PhaseShadowRun
	default:
		// Datas anteriores a 2026: comportamento pre-reforma
		// (apenas tributos legados, sem CBS/IBS)
		return PhaseShadowRun
	}
}

// GetInfo retorna os metadados da fase (quais tributos estao ativos, etc.).
func (r *PhaseResolver) GetInfo(dataOperacao time.Time) PhaseInfo {
	phase := r.Resolve(dataOperacao)
	info, ok := phaseRegistry[phase]
	if !ok {
		// Fallback seguro: apenas tributos legados ativos
		return PhaseInfo{
			Phase:           phase,
			Description:     "Fase desconhecida — fallback para tributos legados",
			CBSActive:       false,
			IBSActive:       false,
			PISCOFINSActive: true,
			ICMSActive:      true,
			ISSActive:       true,
			IPIActive:       true,
		}
	}
	return info
}

// IsShadowRun verifica se a data esta na fase Shadow Run (2026).
// Util para decidir se CBS/IBS compoem o total a pagar.
func (r *PhaseResolver) IsShadowRun(dataOperacao time.Time) bool {
	return r.Resolve(dataOperacao) == PhaseShadowRun
}

// IsLegacyTaxExtinct verifica se tributos legados (PIS/COFINS/ICMS/ISS)
// estao extintos para a data da operacao (fase IVA Dual, 2033+).
func (r *PhaseResolver) IsLegacyTaxExtinct(dataOperacao time.Time) bool {
	return r.Resolve(dataOperacao) == PhaseIVADual
}

// GetReductionFactor retorna o fator de reducao para ICMS/ISS durante a
// transicao subnacional. Fora dessa fase, retorna 0 (sem reducao).
//
// Durante a transicao (2029–2032), o fator cresce linearmente:
//   - 2029: 25% de reducao (fator = 0.25)
//   - 2030: 50% de reducao (fator = 0.50)
//   - 2031: 75% de reducao (fator = 0.75)
//   - 2032: 100% de reducao (fator = 1.00 → tributo extinto)
func (r *PhaseResolver) GetReductionFactor(dataOperacao time.Time) float64 {
	year := dataOperacao.Year()

	switch year {
	case 2029:
		return 0.25
	case 2030:
		return 0.50
	case 2031:
		return 0.75
	case 2032:
		return 1.00
	default:
		return 0.0
	}
}
