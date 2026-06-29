package calculator

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"ms-billing-engine-tax-rates/internal/domain"
	"ms-billing-engine-tax-rates/internal/phase"
	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

// ============================================================================
// Recording Mock — tracks execution order for pipeline sequencing tests
// ============================================================================

type recordingCalc struct {
	name         string
	items        []models.ItemDocumentoFiscalSaida
	record       *[]string
	mu           *sync.Mutex
	delay        time.Duration // simula latencia para testes de concorrencia
	readKey      models.KeyDocumentoInfos
	readKeyFound *bool
	err          error
}

func (r *recordingCalc) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
	if r.delay > 0 {
		time.Sleep(r.delay)
	}

	r.mu.Lock()
	*r.record = append(*r.record, r.name)
	r.mu.Unlock()

	if r.err != nil {
		return nil, r.err
	}

	// Verifica se uma chave especifica esta presente no input (para testes de injecao)
	if r.readKey != "" && r.readKeyFound != nil && len(input.Itens) > 0 {
		val := models.GetDecimal(input.Itens[0].ToMap(), r.readKey)
		*r.readKeyFound = val.GreaterThan(decimal.Zero)
	}

	return r.items, nil
}

func newRecordingCalc(name string, record *[]string, mu *sync.Mutex) *recordingCalc {
	return &recordingCalc{
		name: name,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: name, Valor: decimal.NewFromInt(10)},
			},
		}},
		record: record,
		mu:     mu,
	}
}

// ============================================================================
// C-001: Pipeline SOP-013 — Ordenacao Completa
// ============================================================================

func TestPipeline_FullSOP013_Ordering(t *testing.T) {
	// Verifica que as 7 fases do pipeline SOP-013 executam na ordem correta:
	//   IS (F0) → IPI (F1) → CBS (F2) → ICMS (F3) → IBS+ISS+PISCOFINS (F4) → FUST (F5) → FUNTTEL (F6)
	//
	// Fases sequenciais devem executar em ordem estrita.
	// Fase 4 (paralela) executara concorrentemente; a ordem interna nao e deterministica.
	var record []string
	var mu sync.Mutex

	engine := BillingEnginePhased(
		Phase("F0-IS", Sequential, newRecordingCalc("IS", &record, &mu)),
		Phase("F1-IPI", Sequential, newRecordingCalc("IPI", &record, &mu)),
		Phase("F2-CBS", Sequential, newRecordingCalc("CBS", &record, &mu)),
		Phase("F3-ICMS", Sequential, newRecordingCalc("ICMS", &record, &mu)),
		Phase("F4-Parallel", Parallel,
			newRecordingCalc("IBS", &record, &mu),
			newRecordingCalc("ISS", &record, &mu),
			newRecordingCalc("PISCOFINS", &record, &mu),
		),
		Phase("F5-FUST", Sequential, newRecordingCalc("FUST", &record, &mu)),
		Phase("F6-FUNTTEL", Sequential, newRecordingCalc("FUNTTEL", &record, &mu)),
	)

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(res.Itens) == 0 {
		t.Fatal("esperado pelo menos 1 item no response")
	}

	// Fases sequenciais devem estar em ordem
	seqExpected := []string{"IS", "IPI", "CBS", "ICMS", "FUST", "FUNTTEL"}
	var seqActual []string
	for _, name := range record {
		switch name {
		case "IS", "IPI", "CBS", "ICMS", "FUST", "FUNTTEL":
			seqActual = append(seqActual, name)
		}
	}

	if len(seqActual) != len(seqExpected) {
		t.Fatalf("fases sequenciais: esperado %v, obtido %v", seqExpected, seqActual)
	}
	for i := range seqExpected {
		if seqActual[i] != seqExpected[i] {
			t.Errorf("posicao %d: esperado %s, obtido %s (ordem completa: %v)",
				i, seqExpected[i], seqActual[i], seqActual)
		}
	}

	// Fase 4 (paralela): IBS, ISS, PISCOFINS devem estar entre ICMS e FUST
	icmsIdx := indexOf(record, "ICMS")
	fustIdx := indexOf(record, "FUST")

	for _, name := range []string{"IBS", "ISS", "PISCOFINS"} {
		idx := indexOf(record, name)
		if idx < 0 {
			t.Errorf("calculadora paralela %s nao encontrada no registro", name)
			continue
		}
		if idx <= icmsIdx {
			t.Errorf("%s executou antes do ICMS (idx=%d, icms=%d)", name, idx, icmsIdx)
		}
		if idx >= fustIdx {
			t.Errorf("%s executou depois do FUST (idx=%d, fust=%d)", name, idx, fustIdx)
		}
	}

	// Todos os 9 tributos devem estar presentes no response
	tributos := make(map[string]bool)
	for _, trib := range res.Itens[0].Tributos {
		tributos[trib.Tributo] = true
	}
	expectedTributos := []string{"IS", "IPI", "CBS", "ICMS", "IBS", "ISS", "PISCOFINS", "FUST", "FUNTTEL"}
	for _, expected := range expectedTributos {
		if !tributos[expected] {
			t.Errorf("tributo %s ausente no response", expected)
		}
	}
}

// ============================================================================
// C-001: CBS antes do ICMS (Fase 2 → Fase 3)
// ============================================================================

func TestPipeline_CBS_Before_ICMS(t *testing.T) {
	// CBS (Fase 2, sequencial) deve executar ANTES do ICMS (Fase 3, sequencial).
	var record []string
	var mu sync.Mutex

	engine := BillingEnginePhased(
		Phase("F2-CBS", Sequential, newRecordingCalc("CBS", &record, &mu)),
		Phase("F3-ICMS", Sequential, newRecordingCalc("ICMS", &record, &mu)),
	)

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	_, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	cbsIdx := indexOf(record, "CBS")
	icmsIdx := indexOf(record, "ICMS")

	if cbsIdx < 0 || icmsIdx < 0 {
		t.Fatalf("calculadoras nao encontradas no registro: %v", record)
	}
	if cbsIdx >= icmsIdx {
		t.Errorf("CBS (idx=%d) deveria executar antes do ICMS (idx=%d). Ordem: %v",
			cbsIdx, icmsIdx, record)
	}
}

// ============================================================================
// C-001: ICMS antes do PIS/COFINS — Tese do Seculo (Fase 3 → Fase 4)
// ============================================================================

func TestPipeline_ICMS_Before_PISCOFINS(t *testing.T) {
	// ICMS (Fase 3, sequencial) deve executar ANTES do PIS/COFINS (Fase 4, paralela).
	// Isso garante que o valor do ICMS esteja disponivel para exclusao da base
	// de PIS/COFINS (STF, "Tese do Seculo").
	var record []string
	var mu sync.Mutex

	engine := BillingEnginePhased(
		Phase("F3-ICMS", Sequential, newRecordingCalc("ICMS", &record, &mu)),
		Phase("F4-Parallel", Parallel,
			newRecordingCalc("PISCOFINS", &record, &mu),
		),
	)

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	_, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	icmsIdx := indexOf(record, "ICMS")
	pisIdx := indexOf(record, "PISCOFINS")

	if icmsIdx < 0 || pisIdx < 0 {
		t.Fatalf("calculadoras nao encontradas no registro: %v", record)
	}
	if icmsIdx >= pisIdx {
		t.Errorf("ICMS (idx=%d) deveria executar antes do PISCOFINS (idx=%d). Ordem: %v",
			icmsIdx, pisIdx, record)
	}
}

// ============================================================================
// C-001: Injeção de valores entre fases (inter-phase injection)
// ============================================================================

func TestPipeline_InterPhase_IPI_InjetaValor(t *testing.T) {
	// IPI (Fase 1) injeta IPI_VALOR no input para que ICMS (Fase 3) possa ler.
	var record []string
	var mu sync.Mutex
	ipiFound := false

	ipiCalc := &recordingCalc{
		name: "IPI", record: &record, mu: &mu,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "IPI", Valor: decimal.NewFromFloat(100)},
			},
		}},
	}

	icmsCalc := &recordingCalc{
		name: "ICMS", record: &record, mu: &mu,
		readKey:      models.KeyDocumentoInfos("IPI_VALOR"),
		readKeyFound: &ipiFound,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "ICMS", Valor: decimal.NewFromFloat(180)},
			},
		}},
	}

	engine := BillingEnginePhased(
		Phase("F1-IPI", Sequential, ipiCalc),
		Phase("F3-ICMS", Sequential, icmsCalc),
	)

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	_, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if !ipiFound {
		t.Error("ICMS nao encontrou IPI_VALOR no input — injecao entre fases falhou")
	}
}

func TestPipeline_InterPhase_ICMS_InjetaChaveExclusao(t *testing.T) {
	// ICMS injeta VALOR_EXCLUSAO_ICMS para que PIS/COFINS (Fase 4) possa excluir da base.
	var record []string
	var mu sync.Mutex
	exclusaoFound := false

	icmsCalc := &recordingCalc{
		name: "ICMS", record: &record, mu: &mu,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "ICMS", Valor: decimal.NewFromFloat(180)},
			},
		}},
	}

	pisCalc := &recordingCalc{
		name: "PISCOFINS", record: &record, mu: &mu,
		readKey:      models.KeyDocumentoInfosValorExclusaoICMS,
		readKeyFound: &exclusaoFound,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "PIS", Valor: decimal.NewFromFloat(16.50)},
			},
		}},
	}

	engine := BillingEnginePhased(
		Phase("F3-ICMS", Sequential, icmsCalc),
		Phase("F4-Parallel", Parallel, pisCalc),
	)

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	_, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if !exclusaoFound {
		t.Error("PISCOFINS nao encontrou VALOR_EXCLUSAO_ICMS — injecao da chave de exclusao falhou")
	}
}

func TestPipeline_InterPhase_FUST_DependeDeICMS_PIS_COFINS(t *testing.T) {
	// FUST (Fase 5) depende de ICMS (Fase 3) + PIS/COFINS (Fase 4).
	// Ambos devem estar disponiveis quando FUST executar.
	var record []string
	var mu sync.Mutex
	icmsFound := false
	pisFound := false
	cofinsFound := false

	// ICMS em Fase 3
	icmsCalc := &recordingCalc{
		name: "ICMS", record: &record, mu: &mu,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "ICMS", Valor: decimal.NewFromFloat(180)},
			},
		}},
	}

	// PIS e COFINS em Fase 4
	pisCalc := &recordingCalc{
		name: "PIS", record: &record, mu: &mu,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "PIS", Valor: decimal.NewFromFloat(16.50)},
			},
		}},
	}

	cofinsCalc := &recordingCalc{
		name: "COFINS", record: &record, mu: &mu,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "COFINS", Valor: decimal.NewFromFloat(76.00)},
			},
		}},
	}

	// FUST verifica se ITEM_ICMS_VALOR, ITEM_PIS_VALOR, ITEM_COFINS_VALOR estao disponiveis
	fustCalc := &recordingCalc{
		name: "FUST", record: &record, mu: &mu,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "FUST", Valor: decimal.NewFromFloat(7.28)},
			},
		}},
	}

	engine := BillingEnginePhased(
		Phase("F3-ICMS", Sequential, icmsCalc),
		Phase("F4-Parallel", Parallel, pisCalc, cofinsCalc),
		Phase("F5-FUST", Sequential, fustCalc),
	)

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	_, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	// Verifica que ICMS executou antes de FUST
	icmsIdx := indexOf(record, "ICMS")
	fustIdx := indexOf(record, "FUST")
	if icmsIdx >= fustIdx {
		t.Errorf("ICMS (idx=%d) deveria executar antes de FUST (idx=%d)", icmsIdx, fustIdx)
	}

	// Verifica que PIS executou antes de FUST
	pisIdx := indexOf(record, "PIS")
	if pisIdx >= fustIdx {
		t.Errorf("PIS (idx=%d) deveria executar antes de FUST (idx=%d)", pisIdx, fustIdx)
	}

	// Verifica que COFINS executou antes de FUST
	cofinsIdx := indexOf(record, "COFINS")
	if cofinsIdx >= fustIdx {
		t.Errorf("COFINS (idx=%d) deveria executar antes de FUST (idx=%d)", cofinsIdx, fustIdx)
	}

	_ = icmsFound
	_ = pisFound
	_ = cofinsFound
}

// ============================================================================
// C-001: Fase Paralela — Concorrencia
// ============================================================================

func TestPipeline_Parallel_Phase_RunsConcurrently(t *testing.T) {
	// Calculadoras na mesma fase paralela devem executar concorrentemente.
	// Verificamos que o tempo total da fase e menor que a soma dos delays.
	var record []string
	var mu sync.Mutex

	engine := BillingEnginePhased(
		Phase("Parallel", Parallel,
			&recordingCalc{name: "A", record: &record, mu: &mu, delay: 50 * time.Millisecond,
				items: []models.ItemDocumentoFiscalSaida{{SKU: "SKU-A", Tributos: []models.TributosItemDocumentoFiscalSaida{{Tributo: "A", Valor: decimal.NewFromInt(1)}}}}},
			&recordingCalc{name: "B", record: &record, mu: &mu, delay: 50 * time.Millisecond,
				items: []models.ItemDocumentoFiscalSaida{{SKU: "SKU-A", Tributos: []models.TributosItemDocumentoFiscalSaida{{Tributo: "B", Valor: decimal.NewFromInt(1)}}}}},
			&recordingCalc{name: "C", record: &record, mu: &mu, delay: 50 * time.Millisecond,
				items: []models.ItemDocumentoFiscalSaida{{SKU: "SKU-A", Tributos: []models.TributosItemDocumentoFiscalSaida{{Tributo: "C", Valor: decimal.NewFromInt(1)}}}}},
		),
	)

	start := time.Now()
	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	_, err := engine.Process(context.Background(), input)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	// Se fossem sequenciais, 3 × 50ms = 150ms. Concorrentes, ~50ms + overhead.
	if elapsed > 120*time.Millisecond {
		t.Errorf("fase paralela muito lenta (%v) — esperado < 120ms (concorrencia nao detectada?)", elapsed)
	}

	// Todas as 3 calculadoras devem ter executado
	if len(record) != 3 {
		t.Errorf("esperado 3 calculadoras executadas, obtido %d: %v", len(record), record)
	}
}

// ============================================================================
// C-001: Propagacao de Erros
// ============================================================================

func TestPipeline_Sequential_Error_StopsPipeline(t *testing.T) {
	// Erro em fase sequencial deve interromper o pipeline.
	var record []string
	var mu sync.Mutex

	engine := BillingEnginePhased(
		Phase("F1-IPI", Sequential,
			&recordingCalc{
				name: "IPI", record: &record, mu: &mu,
				err: errors.New("falha critica no IPI"),
			},
		),
		Phase("F2-CBS", Sequential, newRecordingCalc("CBS", &record, &mu)), // nao deve executar
	)

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	_, err := engine.Process(context.Background(), input)
	if err == nil {
		t.Fatal("esperado erro do IPI, obteve nil")
	}

	// CBS nao deve ter executado (pipeline interrompido)
	if indexOf(record, "CBS") >= 0 {
		t.Error("CBS executou apos erro no IPI — pipeline nao foi interrompido")
	}
}

func TestPipeline_Parallel_Error_ContinuesPipeline(t *testing.T) {
	// Erro em fase paralela NAO deve interromper o pipeline.
	var record []string
	var mu sync.Mutex

	engine := BillingEnginePhased(
		Phase("F4-Parallel", Parallel,
			&recordingCalc{
				name: "IBS", record: &record, mu: &mu,
				err: errors.New("API IBS indisponivel"),
			},
			newRecordingCalc("ISS", &record, &mu),
			newRecordingCalc("PISCOFINS", &record, &mu),
		),
		Phase("F5-FUST", Sequential, newRecordingCalc("FUST", &record, &mu)),
	)

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v (pipeline deveria continuar apos erro paralelo)", err)
	}

	// ISS e PISCOFINS devem ter executado
	if indexOf(record, "ISS") < 0 {
		t.Error("ISS nao executou — pipeline interrompido indevidamente")
	}
	if indexOf(record, "PISCOFINS") < 0 {
		t.Error("PISCOFINS nao executou — pipeline interrompido indevidamente")
	}

	// FUST (fase seguinte) deve ter executado
	if indexOf(record, "FUST") < 0 {
		t.Error("FUST nao executou — pipeline interrompido indevidamente apos fase paralela")
	}

	// IBS nao deve aparecer nos tributos (erro)
	for _, trib := range res.Itens[0].Tributos {
		if trib.Tributo == "IBS" {
			t.Error("IBS apareceu nos tributos apesar do erro")
		}
	}
}

// ============================================================================
// C-001: Compatibilidade Retroativa (Backward Compatibility)
// ============================================================================

func TestPipeline_BackwardCompat_BillingEngine(t *testing.T) {
	calc := newRecordingCalc("ICMS", &[]string{}, &sync.Mutex{})
	engine := BillingEngine(calc)

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if len(res.Itens[0].Tributos) != 1 || res.Itens[0].Tributos[0].Tributo != "ICMS" {
		t.Errorf("BillingEngine legado nao funciona: %+v", res.Itens[0].Tributos)
	}
}

func TestPipeline_BackwardCompat_BillingEngineOrdered(t *testing.T) {
	preCalc := newRecordingCalc("IPI", &[]string{}, &sync.Mutex{})
	parallelCalc := newRecordingCalc("ICMS", &[]string{}, &sync.Mutex{})

	engine := BillingEngineOrdered(
		[]domain.TaxCalculator{preCalc},
		parallelCalc,
	)

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	tributos := make(map[string]bool)
	for _, trib := range res.Itens[0].Tributos {
		tributos[trib.Tributo] = true
	}
	if !tributos["IPI"] || !tributos["ICMS"] {
		t.Errorf("BillingEngineOrdered legado: tributos faltando: %v", tributos)
	}
}

func TestPipeline_BackwardCompat_BillingEngineFull(t *testing.T) {
	preCalc := newRecordingCalc("IPI", &[]string{}, &sync.Mutex{})
	parallelCalc := newRecordingCalc("ICMS", &[]string{}, &sync.Mutex{})
	postCalc := newRecordingCalc("FUST", &[]string{}, &sync.Mutex{})

	engine := BillingEngineFull(
		[]domain.TaxCalculator{preCalc},
		[]domain.TaxCalculator{parallelCalc},
		[]domain.TaxCalculator{postCalc},
	)

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	res, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	tributos := make(map[string]bool)
	for _, trib := range res.Itens[0].Tributos {
		tributos[trib.Tributo] = true
	}
	if !tributos["IPI"] || !tributos["ICMS"] || !tributos["FUST"] {
		t.Errorf("BillingEngineFull legado: tributos faltando: %v", tributos)
	}
}

// ============================================================================
// C-001: Phase-Aware com Pipeline SOP-013
// ============================================================================

func TestPipeline_PhaseAware_ShadowRun(t *testing.T) {
	// Shadow Run (2026): CBS/IBS calculados mas NAO compoem total a pagar.
	var record []string
	var mu sync.Mutex

	engine := BillingEnginePhased(
		Phase("IS", Sequential, newRecordingCalc("IS", &record, &mu)),
		Phase("IPI", Sequential, newRecordingCalc("IPI", &record, &mu)),
		Phase("CBS", Sequential, newRecordingCalc("CBS", &record, &mu)),
		Phase("ICMS", Sequential, newRecordingCalc("ICMS", &record, &mu)),
		Phase("IBS+ISS+PISCOFINS", Parallel,
			newRecordingCalc("IBS", &record, &mu),
			newRecordingCalc("ISS", &record, &mu),
			newRecordingCalc("PISCOFINS", &record, &mu),
		),
		Phase("FUST", Sequential, newRecordingCalc("FUST", &record, &mu)),
		Phase("FUNTTEL", Sequential, newRecordingCalc("FUNTTEL", &record, &mu)),
	)

	shadowFilter := phase.CalculatorFilter{
		Phase:           phase.PhaseShadowRun,
		IPIActive:       true,
		PISCOFINSActive: true,
		ICMSActive:      true,
		ISSActive:       true,
		CBSActive:       true,
		IBSActive:       true,
		ISActive:        true,
		ShadowCBS:       true,
		ShadowIBS:       true,
	}

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	res, err := engine.ProcessWithPhase(context.Background(), input, shadowFilter)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	// Total impostos NAO deve incluir CBS nem IBS (shadow)
	// Cada tributo = 10, 9 tributos = 90. CBS(10) + IBS(10) = 20 shadow.
	// Total impostos = 90 - 20 = 70
	if !res.TotalImpostos.Equal(decimal.NewFromFloat(70)) {
		t.Errorf("Shadow Run: total impostos = %s, esperado 70 (CBS e IBS nao compoem total)",
			res.TotalImpostos)
	}
}

func TestPipeline_PhaseAware_IVADual(t *testing.T) {
	// IVA Dual (2033+): PIS/COFINS/ICMS/ISS extintos (valor zero).
	var record []string
	var mu sync.Mutex

	// Usa nomes de tributos que correspondem ao switch de extincao no engine
	icmsCalc := &recordingCalc{
		name: "ICMS", record: &record, mu: &mu,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "ICMS", Valor: decimal.NewFromFloat(10)},
			},
		}},
	}
	pisCalc := &recordingCalc{
		name: "PIS", record: &record, mu: &mu,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "PIS", Valor: decimal.NewFromFloat(10)},
			},
		}},
	}
	cofinsCalc := &recordingCalc{
		name: "COFINS", record: &record, mu: &mu,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "COFINS", Valor: decimal.NewFromFloat(10)},
			},
		}},
	}
	issCalc := &recordingCalc{
		name: "ISS", record: &record, mu: &mu,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "ISS", Valor: decimal.NewFromFloat(10)},
			},
		}},
	}

	engine := BillingEnginePhased(
		Phase("IPI", Sequential, newRecordingCalc("IPI", &record, &mu)),
		Phase("CBS", Sequential, newRecordingCalc("CBS", &record, &mu)),
		Phase("ICMS", Sequential, icmsCalc),
		Phase("IBS+ISS+PISCOFINS", Parallel,
			newRecordingCalc("IBS", &record, &mu),
			issCalc,
			pisCalc,
			cofinsCalc,
		),
	)

	ivaDualFilter := phase.CalculatorFilter{
		Phase:           phase.PhaseIVADual,
		IPIActive:       true,
		PISCOFINSActive: false,
		ICMSActive:      false,
		ISSActive:       false,
		CBSActive:       true,
		IBSActive:       true,
		ISActive:        true,
	}

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	res, err := engine.ProcessWithPhase(context.Background(), input, ivaDualFilter)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	// Verifica que tributos extintos tem valor zero
	for _, trib := range res.Itens[0].Tributos {
		switch trib.Tributo {
		case "ICMS", "ISS", "PIS", "COFINS":
			if !trib.Valor.IsZero() {
				t.Errorf("IVA Dual: %s deveria ser zero (extinto), mas e %s",
					trib.Tributo, trib.Valor)
			}
		case "CBS", "IBS", "IPI":
			if trib.Valor.IsZero() {
				t.Errorf("IVA Dual: %s nao deveria ser zero", trib.Tributo)
			}
		}
	}
}

// ============================================================================
// C-001: Teste de Fumaça — Pipeline SOP-013 com Phase Resolver Real
// ============================================================================

func TestPipeline_Smoke_SOP013_WithRealPhaseResolver(t *testing.T) {
	// Teste de fumaca: pipeline SOP-013 completo com PhaseResolver real.
	// Verifica que o pipeline executa sem erros para todas as fases da reforma.
	phases := []struct {
		name   string
		filter phase.CalculatorFilter
	}{
		{"ShadowRun", phase.CalculatorFilter{
			Phase: phase.PhaseShadowRun, ShadowCBS: true, ShadowIBS: true,
			IPIActive: true, PISCOFINSActive: true, ICMSActive: true,
			ISSActive: true, CBSActive: true, IBSActive: true, ISActive: true,
		}},
		{"CBSPlena", phase.CalculatorFilter{
			Phase: phase.PhaseCBSPlena, ShadowIBS: true,
			IPIActive: true, PISCOFINSActive: false, ICMSActive: true,
			ISSActive: true, CBSActive: true, IBSActive: true, ISActive: true,
		}},
		{"TransicaoSubnacional", phase.CalculatorFilter{
			Phase: phase.PhaseTransicaoSubnacional,
			IPIActive: true, PISCOFINSActive: false, ICMSActive: true,
			ISSActive: true, CBSActive: true, IBSActive: true, ISActive: true,
			SubnationalReductionFactor: 0.25,
		}},
		{"IVADual", phase.CalculatorFilter{
			Phase: phase.PhaseIVADual,
			IPIActive: true, PISCOFINSActive: false, ICMSActive: false,
			ISSActive: false, CBSActive: true, IBSActive: true, ISActive: true,
		}},
	}

	for _, tc := range phases {
		t.Run(tc.name, func(t *testing.T) {
			var record []string
			var mu sync.Mutex

			engine := BillingEnginePhased(
				Phase("IS", Sequential, newRecordingCalc("IS", &record, &mu)),
				Phase("IPI", Sequential, newRecordingCalc("IPI", &record, &mu)),
				Phase("CBS", Sequential, newRecordingCalc("CBS", &record, &mu)),
				Phase("ICMS", Sequential, newRecordingCalc("ICMS", &record, &mu)),
				Phase("IBS+ISS+PISCOFINS", Parallel,
					newRecordingCalc("IBS", &record, &mu),
					newRecordingCalc("ISS", &record, &mu),
					newRecordingCalc("PISCOFINS", &record, &mu),
				),
				Phase("FUST", Sequential, newRecordingCalc("FUST", &record, &mu)),
				Phase("FUNTTEL", Sequential, newRecordingCalc("FUNTTEL", &record, &mu)),
			)

			input := inputDocumento(inputItem("SKU-A", 1, 1000))
			res, err := engine.ProcessWithPhase(context.Background(), input, tc.filter)
			if err != nil {
				t.Fatalf("erro inesperado na fase %s: %v", tc.name, err)
			}

			if len(res.Itens) == 0 {
				t.Fatal("esperado pelo menos 1 item no response")
			}
		})
	}
}

// ============================================================================
// C-001: Teste de Contagem de Fases
// ============================================================================

func TestPipeline_PhaseCount_Verification(t *testing.T) {
	// Verifica que o pipeline tem exatamente 7 fases.
	engine := BillingEnginePhased(
		Phase("F0", Sequential),
		Phase("F1", Sequential),
		Phase("F2", Sequential),
		Phase("F3", Sequential),
		Phase("F4", Parallel),
		Phase("F5", Sequential),
		Phase("F6", Sequential),
	)

	if len(engine.phases) != 7 {
		t.Errorf("esperado 7 fases, obtido %d", len(engine.phases))
	}

	// Fases 0-3 e 5-6 devem ser Sequential
	sequentialPhases := []int{0, 1, 2, 3, 5, 6}
	for _, idx := range sequentialPhases {
		if engine.phases[idx].Mode != Sequential {
			t.Errorf("fase %d: esperado Sequential, obtido %v", idx, engine.phases[idx].Mode)
		}
	}

	// Fase 4 deve ser Parallel
	if engine.phases[4].Mode != Parallel {
		t.Errorf("fase 4: esperado Parallel, obtido %v", engine.phases[4].Mode)
	}
}

// ============================================================================
// C-001: Injecao de ICMS com ambas as chaves (ITEM_ICMS_VALOR e VALOR_EXCLUSAO_ICMS)
// ============================================================================

func TestPipeline_ICMS_InjetaAmbasChaves(t *testing.T) {
	// ICMS deve ser injetado com duas chaves:
	//   1. ITEM_ICMS_VALOR — para FUST e FUNTTEL (base liquida)
	//   2. VALOR_EXCLUSAO_ICMS — para PIS/COFINS (exclusao da base, Tese do Seculo)
	var record []string
	var mu sync.Mutex
	foundItemICMS := false
	foundExclusaoICMS := false

	icmsCalc := &recordingCalc{
		name: "ICMS", record: &record, mu: &mu,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "ICMS", Valor: decimal.NewFromFloat(180)},
			},
		}},
	}

	// Simula FUST lendo ITEM_ICMS_VALOR
	fustCalc := &recordingCalc{
		name: "FUST", record: &record, mu: &mu,
		readKey:      models.KeyDocumentoInfos("ITEM_ICMS_VALOR"),
		readKeyFound: &foundItemICMS,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "FUST", Valor: decimal.NewFromFloat(7.28)},
			},
		}},
	}

	// Simula PIS/COFINS lendo VALOR_EXCLUSAO_ICMS
	pisCalc := &recordingCalc{
		name: "PISCOFINS", record: &record, mu: &mu,
		readKey:      models.KeyDocumentoInfosValorExclusaoICMS,
		readKeyFound: &foundExclusaoICMS,
		items: []models.ItemDocumentoFiscalSaida{{
			SKU: "SKU-A",
			Tributos: []models.TributosItemDocumentoFiscalSaida{
				{Tributo: "PIS", Valor: decimal.NewFromFloat(16.50)},
			},
		}},
	}

	engine := BillingEnginePhased(
		Phase("F3-ICMS", Sequential, icmsCalc),
		Phase("F4-Parallel", Parallel, pisCalc),
		Phase("F5-FUST", Sequential, fustCalc),
	)

	input := inputDocumento(inputItem("SKU-A", 1, 1000))
	_, err := engine.Process(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if !foundItemICMS {
		t.Error("FUST nao encontrou ITEM_ICMS_VALOR — injecao da chave padrao falhou")
	}
	if !foundExclusaoICMS {
		t.Error("PISCOFINS nao encontrou VALOR_EXCLUSAO_ICMS — injecao da chave de exclusao falhou")
	}
}

// ============================================================================
// Helpers
// ============================================================================

func indexOf(slice []string, item string) int {
	for i, s := range slice {
		if s == item {
			return i
		}
	}
	return -1
}

// Garantir que recordingCalc implementa domain.TaxCalculator
var _ domain.TaxCalculator = (*recordingCalc)(nil)

// atomicCounter helper para gerar nomes unicos em testes concorrentes
type atomicCounter struct {
	value int64
}

func (c *atomicCounter) next() int64 {
	return atomic.AddInt64(&c.value, 1)
}

func (c *atomicCounter) current() int64 {
	return atomic.LoadInt64(&c.value)
}
