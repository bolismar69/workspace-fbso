// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/calculator/engine.go
package calculator

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log/slog"
	"sync"

	"ms-billing-engine-tax-rates/internal/domain"
	"ms-billing-engine-tax-rates/internal/phase"
	"taxnexus-billing-core-lib/models"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ExecutionMode define como as calculadoras de uma fase sao executadas.
type ExecutionMode int

const (
	// Sequential executa cada calculadora na ordem em que foram registradas,
	// com injeção de valores entre elas (resultados da calc N disponiveis para calc N+1).
	Sequential ExecutionMode = iota

	// Parallel executa todas as calculadoras concorrentemente via goroutines.
	// Erros sao coletados e logados, mas nao interrompem as demais calculadoras.
	Parallel
)

// CalculationPhase representa uma fase do pipeline de calculo com seu modo
// de execucao e calculadoras associadas.
//
// Exemplos:
//
//	// Fase sequencial com uma unica calculadora
//	CalculationPhase{Name: "IPI", Mode: Sequential, Calculators: []domain.TaxCalculator{ipiCalc}}
//
//	// Fase paralela com multiplas calculadoras
//	CalculationPhase{Name: "IBS+ISS+PISCOFINS", Mode: Parallel, Calculators: []domain.TaxCalculator{ibsCalc, issCalc, pisCofinsCalc}}
type CalculationPhase struct {
	Name        string
	Mode        ExecutionMode
	Calculators []domain.TaxCalculator
}

// Phase e um helper para construir fases comuns de forma concisa.
func Phase(name string, mode ExecutionMode, calcs ...domain.TaxCalculator) CalculationPhase {
	return CalculationPhase{Name: name, Mode: mode, Calculators: calcs}
}

// BillingEngineStruct orquestra o calculo de todos os impostos seguindo
// um pipeline de fases configurável.
//
// Cada fase pode ser Sequential (calculadoras executadas em ordem, com
// injeção de valores entre elas) ou Parallel (calculadoras executadas
// concorrentemente via goroutines).
//
// Entre fases, os valores computados sao injetados nos itens do documento
// de entrada (como detalhes com a chave "<TRIBUTO>_VALOR"), tornando-os
// disponíveis para fases subsequentes.
//
// Pipeline SOP-013 completo (C-001):
//
//	Fase 0 (Sequencial): IS (pre-filtro — NCM seletivo)
//	Fase 1 (Sequencial): IPI (compoe base do ICMS para consumidor final)
//	Fase 2 (Sequencial): CBS ("por fora", nao compoe base de outros)
//	Fase 3 (Sequencial): ICMS (excluido da base PIS/COFINS — Tese do Seculo)
//	Fase 4 (Paralela):   IBS + ISS + PIS/COFINS
//	Fase 5 (Sequencial): FUST (depende de ICMS + PIS + COFINS)
//	Fase 6 (Sequencial): FUNTTEL (depende da base do FUST)
type BillingEngineStruct struct {
	phases []CalculationPhase
}

// BillingEngine cria um engine com calculadoras paralelas (comportamento legado).
// Mantido para compatibilidade com codigo existente e testes.
func BillingEngine(calcs ...domain.TaxCalculator) *BillingEngineStruct {
	return &BillingEngineStruct{
		phases: []CalculationPhase{
			{Name: "Parallel", Mode: Parallel, Calculators: calcs},
		},
	}
}

// BillingEngineOrdered cria um engine onde preCalcs sao executadas primeiro,
// sequencialmente, com injecao de valores entre as fases, seguidas por calcs
// em paralelo. Mantido para compatibilidade.
func BillingEngineOrdered(preCalcs []domain.TaxCalculator, calcs ...domain.TaxCalculator) *BillingEngineStruct {
	var phases []CalculationPhase
	if len(preCalcs) > 0 {
		phases = append(phases, CalculationPhase{
			Name: "Pre-Sequential", Mode: Sequential, Calculators: preCalcs,
		})
	}
	if len(calcs) > 0 {
		phases = append(phases, CalculationPhase{
			Name: "Parallel", Mode: Parallel, Calculators: calcs,
		})
	}
	return &BillingEngineStruct{phases: phases}
}

// BillingEngineFull cria um engine com tres fases (pre/parallel/post).
// Mantido para compatibilidade com codigo existente.
func BillingEngineFull(preCalcs []domain.TaxCalculator, calcs []domain.TaxCalculator, postCalcs []domain.TaxCalculator) *BillingEngineStruct {
	var phases []CalculationPhase
	if len(preCalcs) > 0 {
		phases = append(phases, CalculationPhase{
			Name: "Pre-Sequential", Mode: Sequential, Calculators: preCalcs,
		})
	}
	if len(calcs) > 0 {
		phases = append(phases, CalculationPhase{
			Name: "Parallel", Mode: Parallel, Calculators: calcs,
		})
	}
	if len(postCalcs) > 0 {
		phases = append(phases, CalculationPhase{
			Name: "Post-Sequential", Mode: Sequential, Calculators: postCalcs,
		})
	}
	return &BillingEngineStruct{phases: phases}
}

// BillingEnginePhased cria um engine com fases arbitrarias.
// Este e o construtor principal para o pipeline SOP-013 (C-001).
//
// Exemplo de uso com o pipeline completo:
//
//	engine := BillingEnginePhased(
//	    Phase("IS", Sequential, isFilter),
//	    Phase("IPI", Sequential, ipiCalc),
//	    Phase("CBS", Sequential, cbsCalc),
//	    Phase("ICMS", Sequential, icmsAdapter),
//	    Phase("IBS+ISS+PISCOFINS", Parallel, ibsCalc, issCalc, pisCofinsAdapter),
//	    Phase("FUST", Sequential, fustCalc),
//	    Phase("FUNTTEL", Sequential, funttelCalc),
//	)
func BillingEnginePhased(phases ...CalculationPhase) *BillingEngineStruct {
	return &BillingEngineStruct{phases: phases}
}

// Process executa o pipeline de calculo com comportamento pre-reforma
// (todos os tributos ativos, sem shadow tax). Para comportamento
// phase-aware, use ProcessWithPhase.
func (e *BillingEngineStruct) Process(ctx context.Context, input models.DocumentoFiscalEntrada) (models.DocumentoFiscalSaida, error) {
	defaultFilter := phase.CalculatorFilter{
		Phase:           "PRE_REFORMA",
		IPIActive:       true,
		PISCOFINSActive: true,
		ICMSActive:      true,
		ISSActive:       true,
		CBSActive:       true,
		IBSActive:       true,
		ISActive:        true,
	}
	return e.ProcessWithPhase(ctx, input, defaultFilter)
}

// ProcessWithPhase executa o pipeline de calculo respeitando a fase
// da Reforma Tributaria (Shadow Run, CBS Plena, Transicao Subnacional,
// IVA Dual).
//
// O metodo itera sobre as fases configuradas no engine, executando
// cada fase no modo declarado (Sequential ou Parallel). Entre fases,
// os valores computados sao injetados nos itens do documento de entrada
// para que fases subsequentes possam utiliza-los.
//
// Comportamento por fase (BR-TAX-ACT-005/006):
//   - Shadow Run (2026): CBS/IBS calculados mas NAO compoem total a pagar
//   - CBS Plena (2027): IBS em shadow; PIS/COFINS = 0
//   - Transicao Subnacional (2029–2032): ICMS/ISS com fator de reducao
//   - IVA Dual (2033+): PIS/COFINS/ICMS/ISS extintos
func (e *BillingEngineStruct) ProcessWithPhase(ctx context.Context, input models.DocumentoFiscalEntrada, filter phase.CalculatorFilter) (models.DocumentoFiscalSaida, error) {
	if errs := input.Validate(); len(errs) > 0 {
		return models.DocumentoFiscalSaida{}, errs
	}

	response := models.DocumentoFiscalSaida{
		IDTransaction: uuid.NewString(),
		Itens:         make([]models.ItemDocumentoFiscalSaida, len(input.Itens)),
	}

	for i, item := range input.Itens {
		response.Itens[i].SKU = item.SKU
	}

	slog.Info("Motor de calculo iniciado com fase tributaria",
		"phase", filter.Phase,
		"shadow_cbs", filter.ShadowCBS,
		"shadow_ibs", filter.ShadowIBS,
		"subnational_reduction", filter.SubnationalReductionFactor,
		"icms_active", filter.ICMSActive,
		"pis_cofins_active", filter.PISCOFINSActive,
		"iss_active", filter.ISSActive,
		"pipeline_phases", len(e.phases),
	)

	// =====================================================================
	// Pipeline multi-fase (C-001)
	//
	// Cada fase e executada no modo declarado (Sequential ou Parallel).
	// Apos cada fase, os tributos calculados sao injetados nos detalhes
	// do input para que fases subsequentes possam acessa-los.
	// =====================================================================
	for phaseIdx, p := range e.phases {
		slog.Debug("Executando fase do pipeline",
			"phase_idx", phaseIdx,
			"phase_name", p.Name,
			"mode", map[ExecutionMode]string{Sequential: "sequential", Parallel: "parallel"}[p.Mode],
			"calculators", len(p.Calculators),
		)

		switch p.Mode {
		case Sequential:
			if err := e.executeSequentialPhase(ctx, input, &response, p); err != nil {
				return models.DocumentoFiscalSaida{}, err
			}
		case Parallel:
			e.executeParallelPhase(ctx, input, &response, p)
		}
	}

	// =====================================================================
	// Pos-processamento de fase: shadow tax, reducao subnacional, extincao
	// =====================================================================

	// Aplica fator de reducao na transicao subnacional (ICMS/ISS)
	if filter.SubnationalReductionFactor > 0 {
		for i := range response.Itens {
			for j := range response.Itens[i].Tributos {
				t := &response.Itens[i].Tributos[j]
				if t.Tributo == "ICMS" || t.Tributo == "ICMS_PROPRIO" ||
					t.Tributo == "ICMS_ST" || t.Tributo == "ICMS_DIFAL" ||
					t.Tributo == "ISS" {
					reduction := t.Valor.Mul(decimal.NewFromFloat(filter.SubnationalReductionFactor))
					t.Valor = t.Valor.Sub(reduction)
					t.MoreNumericDetails = append(t.MoreNumericDetails,
						models.Detalhe{Key: "reducao_subnacional_fator", Value: decimal.NewFromFloat(filter.SubnationalReductionFactor)},
						models.Detalhe{Key: "reducao_subnacional_valor", Value: reduction},
					)
					t.MoreTextDetails = append(t.MoreTextDetails,
						models.Detalhe{Key: "fase_tributaria", Value: string(filter.Phase)},
					)
				}
			}
		}
	}

	// Na fase IVA Dual (2033+), PIS/COFINS/ICMS/ISS sao extintos (BR-TAX-ACT-006)
	if filter.Phase == phase.PhaseIVADual {
		for i := range response.Itens {
			for j := range response.Itens[i].Tributos {
				t := &response.Itens[i].Tributos[j]
				switch t.Tributo {
				case "PIS", "COFINS", "ICMS", "ICMS_PROPRIO", "ICMS_ST",
					"ICMS_DIFAL", "ICMS_SIMPLES", "ISS":
					t.Valor = decimal.Zero
					t.BaseCalculo = decimal.Zero
					t.Aliquota = decimal.Zero
					t.MoreTextDetails = append(t.MoreTextDetails,
						models.Detalhe{Key: "fase_tributaria", Value: "IVA_DUAL_EXTINTO"},
					)
				}
			}
		}
	}

	// =====================================================================
	// Consolidação de Totais com distinção shadow/não-shadow
	//
	// Shadow taxes (CBS/IBS em Shadow Run) sao calculados e registrados,
	// mas NAO compoem total_impostos. Isso permite que sistemas consumidores
	// visualizem o impacto futuro sem afetar o valor a pagar (BR-TAX-ACT-005).
	// =====================================================================
	totalImpostos := decimal.Zero
	totalShadow := decimal.Zero
	totalNota := decimal.Zero

	for i := range response.Itens {
		valorItem := input.Itens[i].ValorUnitario.Mul(input.Itens[i].Quantidade)
		somaImpostosItem := decimal.Zero
		somaShadowItem := decimal.Zero

		for _, t := range response.Itens[i].Tributos {
			if (filter.ShadowCBS && t.Tributo == "CBS") ||
				(filter.ShadowIBS && t.Tributo == "IBS") {
				somaShadowItem = somaShadowItem.Add(t.Valor)
			} else {
				somaImpostosItem = somaImpostosItem.Add(t.Valor)
			}
		}

		response.Itens[i].Total = valorItem.Round(2)
		valorLiquido := valorItem.Sub(somaImpostosItem).Sub(somaShadowItem)
		if valorLiquido.LessThan(decimal.Zero) {
			valorLiquido = decimal.Zero
		}
		response.Itens[i].ValorLiquido = valorLiquido.Round(2)

		totalImpostos = totalImpostos.Add(somaImpostosItem)
		totalShadow = totalShadow.Add(somaShadowItem)
		totalNota = totalNota.Add(valorItem)
	}

	response.TotalImpostos = totalImpostos.Round(2)
	response.TotalNota = totalNota.Round(2)

	// Split Payment (GAP-006 — BR-09)
	cbsReter, ibsReter, isReter := e.sumSplitTaxes(response)
	receitaLiquida := totalNota.Sub(cbsReter).Sub(ibsReter).Sub(isReter)
	if receitaLiquida.LessThan(decimal.Zero) {
		receitaLiquida = decimal.Zero
	}
	splitPayload := fmt.Sprintf("%s|%s|%s|%s",
		receitaLiquida.Round(2), cbsReter.Round(2),
		ibsReter.Round(2), isReter.Round(2))
	hash := sha256.Sum256([]byte(splitPayload))
	response.SplitPayment = &models.SplitPayment{
		ValorReceitaLiquida: receitaLiquida.Round(2),
		ValorCBSReter:       cbsReter.Round(2),
		ValorIBSReter:       ibsReter.Round(2),
		ValorISReter:        isReter.Round(2),
		CodigoBarrasSplit:   fmt.Sprintf("%x", hash),
	}

	if totalShadow.GreaterThan(decimal.Zero) {
		slog.Info("Shadow tax calculado (nao compoe total a pagar)",
			"phase", filter.Phase,
			"shadow_total", totalShadow.Round(2),
			"total_a_pagar", response.TotalImpostos,
		)
	}

	return response, nil
}

// executeSequentialPhase executa as calculadoras de uma fase em ordem,
// injetando os valores calculados no input para a proxima calculadora.
//
// Erros em fases sequenciais sao propagados (interrompem o pipeline).
// Isso e intencional: se uma calculadora critica como IPI ou ICMS falhar,
// os calculos subsequentes que dependem de seus valores estariam incorretos.
func (e *BillingEngineStruct) executeSequentialPhase(
	ctx context.Context,
	input models.DocumentoFiscalEntrada,
	response *models.DocumentoFiscalSaida,
	p CalculationPhase,
) error {
	for _, calc := range p.Calculators {
		resItens, err := calc.Calculate(ctx, input)
		if err != nil {
			return err
		}

		// Merge tributos no response
		for i := range resItens {
			if i >= len(response.Itens) {
				break
			}
			response.Itens[i].Tributos = append(response.Itens[i].Tributos, resItens[i].Tributos...)
		}

		// Injeta valores calculados no input para fases/calculadoras subsequentes
		e.injectTributoValues(input, resItens)
	}

	return nil
}

// executeParallelPhase executa as calculadoras de uma fase concorrentemente
// via goroutines. Erros sao coletados e logados, mas nao interrompem as
// demais calculadoras.
func (e *BillingEngineStruct) executeParallelPhase(
	ctx context.Context,
	input models.DocumentoFiscalEntrada,
	response *models.DocumentoFiscalSaida,
	p CalculationPhase,
) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	errChan := make(chan error, len(p.Calculators))

	// Coleciona todos os resultados para injecao pos-fase
	type phaseResult struct {
		items []models.ItemDocumentoFiscalSaida
	}
	results := make([]phaseResult, len(p.Calculators))

	for calcIdx, calc := range p.Calculators {
		wg.Add(1)
		go func(idx int, c domain.TaxCalculator) {
			defer wg.Done()
			resItens, err := c.Calculate(ctx, input)
			if err != nil {
				errChan <- err
				return
			}

			mu.Lock()
			results[idx] = phaseResult{items: resItens}
			for i := range resItens {
				if i < len(response.Itens) {
					response.Itens[i].Tributos = append(response.Itens[i].Tributos, resItens[i].Tributos...)
				}
			}
			mu.Unlock()
		}(calcIdx, calc)
	}
	wg.Wait()
	close(errChan)

	for err := range errChan {
		slog.Warn("Erro em calculadora paralela — calculo parcial",
			"phase", p.Name,
			"error", err,
		)
	}

	// Injeta valores de todas as calculadoras paralelas no input
	// para fases subsequentes (ex: FUST/FUNTTEL precisam de ICMS+PIS+COFINS)
	for _, r := range results {
		if r.items != nil {
			e.injectTributoValues(input, r.items)
		}
	}
}

// injectTributoValues extrai os valores de tributos calculados e os injeta
// nos detalhes dos itens do documento de entrada. Isso permite que fases
// e calculadoras subsequentes acessem valores ja computados.
//
// Para cada tributo encontrado nos resultados, uma entrada e adicionada
// aos detalhes do item de entrada correspondente com a chave
// "<TRIBUTO>_VALOR" (ex: "IPI_VALOR", "ICMS_VALOR", "PIS_VALOR").
//
// Casos especiais:
//   - ICMS: injetado com ambas as chaves ITEM_ICMS_VALOR (FUST/FUNTTEL)
//     e VALOR_EXCLUSAO_ICMS (PIS/COFINS — exclusao da base, Tese do Seculo)
func (e *BillingEngineStruct) injectTributoValues(
	input models.DocumentoFiscalEntrada,
	resItens []models.ItemDocumentoFiscalSaida,
) {
	for i := range resItens {
		if i >= len(input.Itens) {
			break
		}
		for _, trib := range resItens[i].Tributos {
			// Chave generica: <TRIBUTO>_VALOR
			genericKey := models.KeyDocumentoInfos(trib.Tributo + "_VALOR")
			input.Itens[i].AddDetalhe(genericKey, trib.Valor)

			// Chaves especificas por tributo
			switch trib.Tributo {
			case "ICMS", "ICMS_PROPRIO", "ICMS_ST", "ICMS_DIFAL", "ICMS_SIMPLES":
				// ITEM_ICMS_VALOR: usado por FUST e FUNTTEL
				existingICMS := models.GetDecimal(input.Itens[i].ToMap(), models.KeyDocumentoInfos("ITEM_ICMS_VALOR"))
				input.Itens[i].AddDetalhe(models.KeyDocumentoInfos("ITEM_ICMS_VALOR"), existingICMS.Add(trib.Valor))

				// VALOR_EXCLUSAO_ICMS: usado por PIS/COFINS (Tese do Seculo)
				existingExclusao := models.GetDecimal(input.Itens[i].ToMap(), models.KeyDocumentoInfosValorExclusaoICMS)
				input.Itens[i].AddDetalhe(models.KeyDocumentoInfosValorExclusaoICMS, existingExclusao.Add(trib.Valor))

			case "PIS":
				existingPIS := models.GetDecimal(input.Itens[i].ToMap(), models.KeyDocumentoInfos("ITEM_PIS_VALOR"))
				input.Itens[i].AddDetalhe(models.KeyDocumentoInfos("ITEM_PIS_VALOR"), existingPIS.Add(trib.Valor))

			case "COFINS":
				existingCOFINS := models.GetDecimal(input.Itens[i].ToMap(), models.KeyDocumentoInfos("ITEM_COFINS_VALOR"))
				input.Itens[i].AddDetalhe(models.KeyDocumentoInfos("ITEM_COFINS_VALOR"), existingCOFINS.Add(trib.Valor))

			case "IPI":
				existingIPI := models.GetDecimal(input.Itens[i].ToMap(), models.KeyDocumentoInfos("ITEM_IPI_VALOR"))
				input.Itens[i].AddDetalhe(models.KeyDocumentoInfos("ITEM_IPI_VALOR"), existingIPI.Add(trib.Valor))
			}
		}
	}
}

// sumSplitTaxes soma CBS, IBS e IS de todos os itens para o split payment (GAP-006).
func (e *BillingEngineStruct) sumSplitTaxes(response models.DocumentoFiscalSaida) (cbs, ibs, is decimal.Decimal) {
	for _, item := range response.Itens {
		for _, t := range item.Tributos {
			switch t.Tributo {
			case "CBS":
				cbs = cbs.Add(t.Valor)
			case "IBS":
				ibs = ibs.Add(t.Valor)
			case "IS":
				is = is.Add(t.Valor)
			}
		}
	}
	return
}
