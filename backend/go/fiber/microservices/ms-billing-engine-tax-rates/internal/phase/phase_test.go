package phase

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

// --- PhaseResolver Tests ---

func TestPhaseResolver_2026_ShadowRun(t *testing.T) {
	resolver := NewPhaseResolver()

	tests := []struct {
		name string
		date time.Time
	}{
		{"inicio 2026", time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"meio 2026", time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC)},
		{"fim 2026", time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			phase := resolver.Resolve(tt.date)
			if phase != PhaseShadowRun {
				t.Errorf("esperado %s, obteve %s para data %s", PhaseShadowRun, phase, tt.date.Format("2006-01-02"))
			}
		})
	}
}

func TestPhaseResolver_2027_CBSPlena(t *testing.T) {
	resolver := NewPhaseResolver()

	tests := []struct {
		name string
		date time.Time
	}{
		{"inicio 2027", time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"meio 2027", time.Date(2027, 7, 1, 0, 0, 0, 0, time.UTC)},
		{"fim 2028", time.Date(2028, 12, 31, 23, 59, 59, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			phase := resolver.Resolve(tt.date)
			if phase != PhaseCBSPlena {
				t.Errorf("esperado %s, obteve %s para data %s", PhaseCBSPlena, phase, tt.date.Format("2006-01-02"))
			}
		})
	}
}

func TestPhaseResolver_2030_Transicao(t *testing.T) {
	resolver := NewPhaseResolver()

	tests := []struct {
		name string
		date time.Time
	}{
		{"inicio transicao 2029", time.Date(2029, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"meio 2030", time.Date(2030, 6, 15, 0, 0, 0, 0, time.UTC)},
		{"meio 2031", time.Date(2031, 3, 1, 0, 0, 0, 0, time.UTC)},
		{"fim transicao 2032", time.Date(2032, 12, 31, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			phase := resolver.Resolve(tt.date)
			if phase != PhaseTransicaoSubnacional {
				t.Errorf("esperado %s, obteve %s para data %s", PhaseTransicaoSubnacional, phase, tt.date.Format("2006-01-02"))
			}
		})
	}
}

func TestPhaseResolver_2033_IVADual(t *testing.T) {
	resolver := NewPhaseResolver()

	tests := []struct {
		name string
		date time.Time
	}{
		{"inicio IVA Dual", time.Date(2033, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"futuro 2040", time.Date(2040, 5, 10, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			phase := resolver.Resolve(tt.date)
			if phase != PhaseIVADual {
				t.Errorf("esperado %s, obteve %s para data %s", PhaseIVADual, phase, tt.date.Format("2006-01-02"))
			}
		})
	}
}

func TestPhaseResolver_PreReforma(t *testing.T) {
	resolver := NewPhaseResolver()

	// Datas anteriores a 2026: comportamento pre-reforma (apenas tributos legados)
	date := time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC)
	phase := resolver.Resolve(date)
	if phase != PhaseShadowRun {
		t.Errorf("datas pre-2026 devem usar PhaseShadowRun como fallback, obteve %s", phase)
	}
}

// --- PhaseInfo Tests ---

func TestGetInfo_ShadowRun_Flags(t *testing.T) {
	resolver := NewPhaseResolver()
	date := time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC)

	info := resolver.GetInfo(date)

	if info.CBSActive {
		t.Error("CBS nao deve estar ativo (compoe total) no Shadow Run")
	}
	if info.IBSActive {
		t.Error("IBS nao deve estar ativo (compoe total) no Shadow Run")
	}
	if !info.PISCOFINSActive {
		t.Error("PIS/COFINS devem estar ativos no Shadow Run")
	}
	if !info.ICMSActive {
		t.Error("ICMS deve estar ativo no Shadow Run")
	}
	if !info.ISSActive {
		t.Error("ISS deve estar ativo no Shadow Run")
	}
	if !info.IPIActive {
		t.Error("IPI deve estar ativo em todas as fases")
	}
}

func TestGetInfo_IVADual_Flags(t *testing.T) {
	resolver := NewPhaseResolver()
	date := time.Date(2033, 1, 1, 0, 0, 0, 0, time.UTC)

	info := resolver.GetInfo(date)

	if !info.CBSActive {
		t.Error("CBS deve estar ativo no IVA Dual")
	}
	if !info.IBSActive {
		t.Error("IBS deve estar ativo no IVA Dual")
	}
	if info.PISCOFINSActive {
		t.Error("PIS/COFINS devem estar extintos no IVA Dual")
	}
	if info.ICMSActive {
		t.Error("ICMS deve estar extinto no IVA Dual")
	}
	if info.ISSActive {
		t.Error("ISS deve estar extinto no IVA Dual")
	}
	if !info.IPIActive {
		t.Error("IPI deve estar ativo em todas as fases")
	}
}

// --- Reduction Factor Tests ---

func TestReductionFactor_TransicaoSubnacional(t *testing.T) {
	resolver := NewPhaseResolver()

	tests := []struct {
		year   int
		factor float64
	}{
		{2029, 0.25},
		{2030, 0.50},
		{2031, 0.75},
		{2032, 1.00},
	}

	for _, tt := range tests {
		t.Run("ano_"+decimal.NewFromInt(int64(tt.year)).String(), func(t *testing.T) {
			date := time.Date(tt.year, 6, 15, 0, 0, 0, 0, time.UTC)
			factor := resolver.GetReductionFactor(date)
			if factor != tt.factor {
				t.Errorf("ano %d: esperado fator %.2f, obteve %.2f", tt.year, tt.factor, factor)
			}
		})
	}
}

func TestReductionFactor_OutsideTransicao(t *testing.T) {
	resolver := NewPhaseResolver()

	tests := []time.Time{
		time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC),
		time.Date(2027, 6, 15, 0, 0, 0, 0, time.UTC),
		time.Date(2033, 6, 15, 0, 0, 0, 0, time.UTC),
	}

	for _, date := range tests {
		t.Run(date.Format("2006"), func(t *testing.T) {
			factor := resolver.GetReductionFactor(date)
			if factor != 0.0 {
				t.Errorf("fora da transicao, fator deve ser 0, obteve %.2f", factor)
			}
		})
	}
}

func TestIsShadowRun(t *testing.T) {
	resolver := NewPhaseResolver()

	if !resolver.IsShadowRun(time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC)) {
		t.Error("2026 deve ser Shadow Run")
	}
	if resolver.IsShadowRun(time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC)) {
		t.Error("2027 nao deve ser Shadow Run")
	}
}

func TestIsLegacyTaxExtinct(t *testing.T) {
	resolver := NewPhaseResolver()

	if resolver.IsLegacyTaxExtinct(time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC)) {
		t.Error("2026: tributos legados nao devem estar extintos")
	}
	if resolver.IsLegacyTaxExtinct(time.Date(2030, 6, 15, 0, 0, 0, 0, time.UTC)) {
		t.Error("2030: tributos legados ainda vigentes (transicao)")
	}
	if !resolver.IsLegacyTaxExtinct(time.Date(2033, 1, 1, 0, 0, 0, 0, time.UTC)) {
		t.Error("2033+: tributos legados devem estar extintos (IVA Dual)")
	}
}

// --- TaxSelector Tests ---

func TestTaxSelector_Filter_ShadowRun(t *testing.T) {
	resolver := NewPhaseResolver()
	selector := NewTaxSelector(resolver)

	filter := selector.Filter(time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC))

	if filter.Phase != PhaseShadowRun {
		t.Errorf("esperado %s, obteve %s", PhaseShadowRun, filter.Phase)
	}
	if !filter.ShadowCBS {
		t.Error("ShadowCBS deve ser true no Shadow Run")
	}
	if !filter.ShadowIBS {
		t.Error("ShadowIBS deve ser true no Shadow Run")
	}
	if !filter.ICMSActive {
		t.Error("ICMS deve estar ativo no Shadow Run")
	}
	if !filter.PISCOFINSActive {
		t.Error("PIS/COFINS devem estar ativos no Shadow Run")
	}
}

func TestTaxSelector_Filter_IVADual(t *testing.T) {
	resolver := NewPhaseResolver()
	selector := NewTaxSelector(resolver)

	filter := selector.Filter(time.Date(2033, 1, 1, 0, 0, 0, 0, time.UTC))

	if filter.Phase != PhaseIVADual {
		t.Errorf("esperado %s, obteve %s", PhaseIVADual, filter.Phase)
	}
	if filter.ShadowCBS {
		t.Error("ShadowCBS deve ser false no IVA Dual")
	}
	if filter.ShadowIBS {
		t.Error("ShadowIBS deve ser false no IVA Dual")
	}
	if !filter.CBSActive {
		t.Error("CBS deve estar ativo no IVA Dual")
	}
	if !filter.IBSActive {
		t.Error("IBS deve estar ativo no IVA Dual")
	}
}

func TestTaxSelector_ShouldIncludeInTotal(t *testing.T) {
	resolver := NewPhaseResolver()
	selector := NewTaxSelector(resolver)

	// Shadow Run (2026): CBS e IBS NAO compoem total
	date2026 := time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC)
	if selector.ShouldIncludeInTotal(date2026, "CBS") {
		t.Error("CBS nao deve compor total no Shadow Run (BR-TAX-ACT-005)")
	}
	if selector.ShouldIncludeInTotal(date2026, "IBS") {
		t.Error("IBS nao deve compor total no Shadow Run (BR-TAX-ACT-005)")
	}
	if !selector.ShouldIncludeInTotal(date2026, "ICMS") {
		t.Error("ICMS deve compor total no Shadow Run")
	}
	if !selector.ShouldIncludeInTotal(date2026, "PIS") {
		t.Error("PIS deve compor total no Shadow Run")
	}

	// CBS Plena (2027): CBS ativo, IBS shadow
	date2027 := time.Date(2027, 6, 15, 0, 0, 0, 0, time.UTC)
	if !selector.ShouldIncludeInTotal(date2027, "CBS") {
		t.Error("CBS deve compor total na CBS Plena")
	}
	if selector.ShouldIncludeInTotal(date2027, "IBS") {
		t.Error("IBS nao deve compor total na CBS Plena")
	}

	// IVA Dual (2033): apenas CBS/IBS/IS/IPI compoem
	date2033 := time.Date(2033, 1, 1, 0, 0, 0, 0, time.UTC)
	if selector.ShouldIncludeInTotal(date2033, "PIS") {
		t.Error("PIS nao deve compor total no IVA Dual (BR-TAX-ACT-006)")
	}
	if selector.ShouldIncludeInTotal(date2033, "ICMS") {
		t.Error("ICMS nao deve compor total no IVA Dual (BR-TAX-ACT-006)")
	}
	if selector.ShouldIncludeInTotal(date2033, "ISS") {
		t.Error("ISS nao deve compor total no IVA Dual (BR-TAX-ACT-006)")
	}
	if !selector.ShouldIncludeInTotal(date2033, "CBS") {
		t.Error("CBS deve compor total no IVA Dual")
	}
	if !selector.ShouldIncludeInTotal(date2033, "IPI") {
		t.Error("IPI deve compor total em qualquer fase")
	}
}

func TestTaxSelector_SubnationalReduction(t *testing.T) {
	resolver := NewPhaseResolver()
	selector := NewTaxSelector(resolver)

	// 2030: 50% de reducao
	filter := selector.Filter(time.Date(2030, 6, 15, 0, 0, 0, 0, time.UTC))
	if filter.SubnationalReductionFactor != 0.50 {
		t.Errorf("esperado fator 0.50 para 2030, obteve %.2f", filter.SubnationalReductionFactor)
	}

	// 2026: sem reducao
	filter2026 := selector.Filter(time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC))
	if filter2026.SubnationalReductionFactor != 0.0 {
		t.Errorf("esperado fator 0.0 para 2026, obteve %.2f", filter2026.SubnationalReductionFactor)
	}
}
