package phase

import (
	"time"

	"ms-billing-engine-tax-rates/internal/domain"
)

// TaxSelector aplica a matriz DT-001 para selecionar quais calculadoras
// estao ativas em cada fase tributaria.
//
// A matriz DT-001 define, para cada combinacao (fase, natureza da operacao),
// quais tributos devem ser calculados e se compoem ou nao o total a pagar.
//
// Referencias:
//   - RULES-CATALOG-FIN-00001 DT-001
//   - PROCEDURE-FIN-00001 SOP-013
type TaxSelector struct {
	resolver *PhaseResolver
}

// NewTaxSelector cria um novo seletor de calculadoras.
func NewTaxSelector(resolver *PhaseResolver) *TaxSelector {
	return &TaxSelector{resolver: resolver}
}

// CalculatorFilter contem as decisoes de quais calculadoras estao ativas
// para uma determinada operacao fiscal.
type CalculatorFilter struct {
	Phase Phase

	// Calculadoras ativas (devem ser executadas)
	IPIActive       bool
	PISCOFINSActive bool
	ICMSActive      bool
	ISSActive       bool
	CBSActive       bool
	IBSActive       bool
	ISActive        bool

	// Shadow run: tributos calculados mas NAO compoem total a pagar
	// (usado para CBS e IBS em 2026)
	ShadowCBS bool
	ShadowIBS bool

	// Fator de reducao para ICMS/ISS na transicao subnacional (0.0 a 1.0)
	SubnationalReductionFactor float64
}

// Filter determina quais calculadoras estao ativas para uma data de operacao.
// Aplica as regras DT-001 e BR-TAX-ACT-005/006.
func (s *TaxSelector) Filter(dataOperacao time.Time) CalculatorFilter {
	info := s.resolver.GetInfo(dataOperacao)
	reductionFactor := s.resolver.GetReductionFactor(dataOperacao)
	phase := s.resolver.Resolve(dataOperacao)

	filter := CalculatorFilter{
		Phase:                      phase,
		IPIActive:                  info.IPIActive,
		PISCOFINSActive:            info.PISCOFINSActive,
		ICMSActive:                 info.ICMSActive,
		ISSActive:                  info.ISSActive,
		CBSActive:                  true, // CBS sempre calculada (pode ser shadow)
		IBSActive:                  true, // IBS sempre calculada (pode ser shadow)
		ISActive:                   true, // IS sempre calculado (pre-filtro)
		SubnationalReductionFactor: reductionFactor,
	}

	// Shadow run: CBS e IBS nao compoem total a pagar (BR-TAX-ACT-005)
	if phase == PhaseShadowRun {
		filter.ShadowCBS = true
		filter.ShadowIBS = true
	}

	// CBS Plena: CBS ativa, IBS em shadow
	if phase == PhaseCBSPlena {
		filter.ShadowIBS = true
	}

	return filter
}

// ShouldIncludeInTotal indica se um tributo especifico deve compor o
// total a pagar do documento fiscal.
//
// Regras:
//   - Shadow Run (2026): CBS e IBS NAO compoem total (BR-TAX-ACT-005)
//   - CBS Plena (2027): IBS NAO compoe total
//   - IVA Dual (2033+): apenas CBS e IBS compoem total (BR-TAX-ACT-006)
func (s *TaxSelector) ShouldIncludeInTotal(dataOperacao time.Time, tributo string) bool {
	phase := s.resolver.Resolve(dataOperacao)

	switch tributo {
	case "CBS":
		return phase != PhaseShadowRun
	case "IBS":
		return phase != PhaseShadowRun && phase != PhaseCBSPlena
	case "IS":
		return true // IS sempre compoe total
	case "IPI":
		return true // IPI sempre compoe total
	case "PIS", "COFINS":
		return phase != PhaseCBSPlena && phase != PhaseTransicaoSubnacional && phase != PhaseIVADual
	case "ICMS", "ICMS_PROPRIO", "ICMS_ST", "ICMS_DIFAL", "ICMS_SIMPLES":
		return phase != PhaseIVADual
	case "ISS":
		return phase != PhaseIVADual
	default:
		return true
	}
}

// LegacyAdapterSelector cria uma lista de adaptadores para calculadoras
// baseada no filter. Retorna as calculadoras legacy que devem ser executadas.
//
// Deprecated: Use CalculatorFilter diretamente no engine para maior controle.
// Mantido para compatibilidade com codigo existente.
func (s *TaxSelector) LegacyAdapterSelector(
	filter CalculatorFilter,
	ipCalc domain.TaxCalculator,
	pisCofinsCalc domain.TaxCalculator,
	icmsCalc domain.TaxCalculator,
	issCalc domain.TaxCalculator,
	reformaCalc domain.TaxCalculator,
) []domain.TaxCalculator {
	var calcs []domain.TaxCalculator

	if filter.IPIActive && ipCalc != nil {
		calcs = append(calcs, ipCalc)
	}
	if filter.PISCOFINSActive && pisCofinsCalc != nil {
		calcs = append(calcs, pisCofinsCalc)
	}
	if filter.ICMSActive && icmsCalc != nil {
		calcs = append(calcs, icmsCalc)
	}
	if filter.ISSActive && issCalc != nil {
		calcs = append(calcs, issCalc)
	}
	if reformaCalc != nil {
		calcs = append(calcs, reformaCalc)
	}

	return calcs
}
