package legacy

import (
	"context"
	"fmt"
	"log/slog"

	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"
	"taxnexus-billing-core-lib/utils"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ICMSCalculator struct {
	repo repository.TaxRepository
}

func NewICMSCalculator(r repository.TaxRepository) *ICMSCalculator {
	return &ICMSCalculator{repo: r}
}

func (c *ICMSCalculator) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) (models.DocumentoFiscalSaida, error) {

	response := models.DocumentoFiscalSaida{
		IDTransaction: uuid.NewString(),
	}

	isInterestadual := input.LocalizacaoOrigem.UF != input.LocalizacaoDestino.UF
	isSimples := models.NormalizeCRTEmitente(input.CRTEmitente) == models.CRTEmitenteSimples

	for _, item := range input.Itens {

		attrs := item.ToMap()
		slog.Debug("Calculando ICMS", "SKU", item.SKU, "Valor", item.ValorUnitario.Mul(item.Quantidade), "Detalhes", attrs)

		itemRes := models.ItemDocumentoFiscalSaida{
			SKU: item.SKU,
		}

		baseItem := item.ValorUnitario.Mul(item.Quantidade)

		// 1. Obtém a configuração efetiva (merge regra geral + exceção de produto)
		config, err := c.getEffectiveTaxConfig(ctx, item, input)
		if err != nil {
			slog.Warn("Configuracao ICMS indisponivel, pulando item", "SKU", item.SKU, "erro", err)
			itemRes.Total = baseItem
			response.Itens = append(response.Itens, itemRes)
			continue
		}

		// 2. Simples Nacional — lógica própria com equivalência CSOSN
		//    (Simples Nacional NAO aplica desoneracao classica — SOP-017 edge case)
		if isSimples {
			c.calcularICMSSimples(ctx, &itemRes, item, input, attrs, baseItem)
		} else {
			// 3. Regime Normal
			if !isInterestadual {
				c.calcularICMSOperacaoInterna(&itemRes, item, input, config, baseItem, attrs)
			} else {
				c.calcularICMSOperacaoInterestadual(&itemRes, item, input, config, baseItem)
			}
		}

		itemRes.Total = baseItem
		response.Itens = append(response.Itens, itemRes)
		response.TotalImpostos = response.TotalImpostos.Add(itemRes.Total)
	}

	return response, nil
}

func (c *ICMSCalculator) calcularICMSSimples(
	ctx context.Context,
	itemRes *models.ItemDocumentoFiscalSaida,
	item models.ItemDocumentoFiscalEntrada,
	input models.DocumentoFiscalEntrada,
	attrs map[string]interface{},
	baseItem decimal.Decimal,
) {
	CSOSN := models.GetString(attrs, models.KeyDocumentoInfosItemCSOSN)
	if CSOSN == "" {
		return
	}

	equiv, err := c.repo.GetEquivalence(ctx, CSOSN, string(models.NormalizeTipoOperacao(input.TipoOperacaoFiscal)))
	if err != nil || !equiv.PermiteCredito {
		return
	}

	faixa, err := c.repo.GetSimplesFaixa(ctx,
		models.GetString(attrs, models.KeyDocumentoInfosAnexoSimples),
		models.GetDecimal(attrs, models.KeyDocumentoInfosRBT12))
	if err != nil {
		slog.Warn("Faixa Simples nao encontrada", "SKU", item.SKU, "erro", err)
		return
	}

	aliq := utils.CalcularAliquotaEfetivaSimples(
		models.GetDecimal(attrs, models.KeyDocumentoInfosRBT12),
		faixa.AliqNominal,
		faixa.ValorDeduzir,
		faixa.PercIcmsAnexo,
	)

	aliqPct := aliq.Div(decimal.NewFromInt(100))
	vlorImposto := baseItem.Mul(aliqPct)

	itemRes.Tributos = append(itemRes.Tributos, models.TributosItemDocumentoFiscalSaida{
		Tributo:     "ICMS_SIMPLES",
		BaseCalculo: baseItem,
		Aliquota:    aliq,
		Valor:       vlorImposto.Round(2),
		MoreNumericDetails: []models.Detalhe{
			{Key: "rbt12", Value: models.GetDecimal(attrs, models.KeyDocumentoInfosRBT12)},
			{Key: "aliq_nominal", Value: faixa.AliqNominal},
			{Key: "valor_deduzir", Value: faixa.ValorDeduzir},
			{Key: "perc_icms_anexo", Value: faixa.PercIcmsAnexo},
			{Key: "aliq_efetiva", Value: aliq.Round(4)},
		},
		MoreTextDetails: []models.Detalhe{
			{Key: "regime", Value: input.CRTEmitente},
			{Key: "anexo_simples", Value: models.GetString(attrs, models.KeyDocumentoInfosAnexoSimples)},
			{Key: "csosn", Value: CSOSN},
			{Key: "cst_equivalente", Value: equiv.CSTEquivalente},
		},
	})
}

func (c *ICMSCalculator) calcularICMSOperacaoInterna(
	itemRes *models.ItemDocumentoFiscalSaida,
	item models.ItemDocumentoFiscalEntrada,
	input models.DocumentoFiscalEntrada,
	config repository.ICMSConfig,
	baseItem decimal.Decimal,
	attrs map[string]interface{},
) {
	cst := models.NormalizeCSTICMS(models.GetString(attrs, models.KeyDocumentoInfosItemSubstituirCSTICMS))
	if cst == "" {
		cst = models.CSTICMS(models.GetString(attrs, models.KeyDocumentoInfos("CST_ICMS")))
	}

	// ICMS-ST (interno): CST 010 ou config com protocolo ST + MVA
	if cst == models.CSTICMS010 || (config.PossuiProtocoloST && !config.MVAPadrao.IsZero()) {
		mva := config.MVAPadrao
		if cst == models.CSTICMS010 {
			mvaOverride := models.GetDecimal(attrs, models.KeyDocumentoInfosItemSubstituirMVAPercentual)
			if !mvaOverride.IsZero() {
				mva = mvaOverride
			}
		}

		aliquotaFator := mva.Div(decimal.NewFromInt(100)).Add(decimal.NewFromInt(1))
		baseST := baseItem.Mul(aliquotaFator)
		aliqIntPct := config.AliquotaInterna.Div(decimal.NewFromInt(100))
		valorST := baseST.Mul(aliqIntPct).Round(2)

		itemRes.Tributos = append(itemRes.Tributos, models.TributosItemDocumentoFiscalSaida{
			Tributo:     "ICMS_ST",
			BaseCalculo: baseST,
			Aliquota:    config.AliquotaInterna,
			Valor:       valorST,
			CST:         string(cst),
			MoreNumericDetails: []models.Detalhe{
				{Key: "valor_item", Value: baseItem},
				{Key: "mva_percentual", Value: mva},
				{Key: "base_st_calculada", Value: baseST},
				{Key: "valor_st_calculado", Value: valorST},
				{Key: "aliquota_interna_destino", Value: config.AliquotaInterna},
			},
			MoreTextDetails: []models.Detalhe{
				{Key: "fonte_regra", Value: "config_efetiva_st_interno"},
				{Key: "uf_origem", Value: input.LocalizacaoOrigem.UF},
				{Key: "uf_destino", Value: input.LocalizacaoDestino.UF},
			},
		})
		return
	}

	// ==================================================================
	// F-004: ICMS Desonerado — Reducao de Base e Limitacao de Aliquota
	// Ref: SOP-017, BR-TAX-CALC-021/022, BR-TAX-CONS-013, BR-TAX-ACT-007
	//
	// Aplica apenas em Regime Normal (nao Simples).
	// CST deve permitir desoneracao: {20, 30, 40, 41, 50, 70, 90}.
	// CST 00 (tributacao integral) NAO permite.
	// ==================================================================
	desoneracaoAplicar := models.GetDecimal(attrs, models.KeyDocumentoInfosItemDesoneracaoAplicar)
	desoneracaoAplicarStr := models.GetString(attrs, models.KeyDocumentoInfosItemDesoneracaoAplicar)
	if desoneracaoAplicar.GreaterThan(decimal.Zero) || desoneracaoAplicarStr == "true" {
		desoneracaoCalc := NewICMSDesoneracao()

		if desoneracaoCalc.CSTPermiteDesoneracao(cst) {
			pctReducao := models.GetDecimal(attrs, models.KeyDocumentoInfosItemDesoneracaoPercentual)
			aliqAlvo := models.GetDecimal(attrs, models.KeyDocumentoInfosItemDesoneracaoCargaEfetivaAlvo)
			motDesICMSCod := models.GetDecimal(attrs, models.KeyDocumentoInfosItemDesoneracaoMotivoICMS)

			params := DesoneracaoParams{
				ValorItem:          baseItem,
				AliquotaNominal:    config.AliquotaInterna,
				PercentualReducao:  pctReducao,
				AliquotaAlvo:       aliqAlvo,
				MotivoDesoneracao:  MotivoDesoneracaoICMS(motDesICMSCod.IntPart()),
				CodBeneficioFiscal: models.GetString(attrs, models.KeyDocumentoInfosItemDetalheCodBeneficioFiscal),
			}

			modo := desoneracaoCalc.DeterminarModo(pctReducao, aliqAlvo)
			if modo != nil {
				var result DesoneracaoResult
				switch *modo {
				case ModoLimitacaoAliquota:
					result = desoneracaoCalc.CalcularLimitacaoAliquota(params)
				default:
					result = desoneracaoCalc.CalcularReducaoBase(params)
				}

				numDetails, textDetails := BuildDesoneracaoDetails(result, params)

				tributoDeson := models.TributosItemDocumentoFiscalSaida{
					Tributo:            "ICMS",
					BaseCalculo:        result.BaseReduzida,
					Aliquota:           config.AliquotaInterna,
					Valor:              result.ICMS,
					CST:                string(cst),
					MoreNumericDetails: numDetails,
					MoreTextDetails:    textDetails,
				}

				// FCP sobre base reduzida (BR-TAX-ACT-007)
				if !config.PercentualFCP.IsZero() {
					fcpPct := config.PercentualFCP.Div(decimal.NewFromInt(100))
					valorFCP := result.BaseReduzida.Mul(fcpPct).Round(2)
					tributoDeson.MoreNumericDetails = append(tributoDeson.MoreNumericDetails,
						models.Detalhe{Key: "fcp_percentual", Value: config.PercentualFCP},
						models.Detalhe{Key: "valor_fcp", Value: valorFCP},
					)
					tributoDeson.Valor = tributoDeson.Valor.Add(valorFCP)
				}

				itemRes.Tributos = append(itemRes.Tributos, tributoDeson)

				slog.Debug("ICMS Desonerado Calculado",
					"SKU", item.SKU,
					"modo", result.Modo,
					"base_reduzida", result.BaseReduzida,
					"icms", result.ICMS,
					"vicms_deson", result.ValorICMSDeson,
					"motivo", result.MotDesICMS,
				)
				return
			}
		} else {
			slog.Warn("Desoneracao solicitada mas CST nao permite",
				"SKU", item.SKU,
				"cst", cst,
			)
		}
	}

	// ICMS Próprio (interno) — sem desoneracao
	base := baseItem
	if !config.ReducaoBase.IsZero() {
		fatorReducao := decimal.NewFromInt(1).Sub(config.ReducaoBase.Div(decimal.NewFromInt(100)))
		base = base.Mul(fatorReducao)
	}

	aliqPct := config.AliquotaInterna.Div(decimal.NewFromInt(100))
	valorICMS := base.Mul(aliqPct).Round(2)

	tributo := models.TributosItemDocumentoFiscalSaida{
		Tributo:     "ICMS",
		BaseCalculo: base,
		Aliquota:    config.AliquotaInterna,
		Valor:       valorICMS,
		CST:         config.CSTPadrao,
		MoreNumericDetails: []models.Detalhe{
			{Key: "valor_item", Value: baseItem},
			{Key: "reducao_base_percentual", Value: config.ReducaoBase},
			{Key: "base_calculo_efetiva", Value: base},
			{Key: "aliquota_interna", Value: config.AliquotaInterna},
			{Key: "valor_icms", Value: valorICMS},
		},
		MoreTextDetails: []models.Detalhe{
			{Key: "fonte_regra", Value: "config_efetiva_icms_proprio_interno"},
			{Key: "uf_origem", Value: input.LocalizacaoOrigem.UF},
			{Key: "uf_destino", Value: input.LocalizacaoDestino.UF},
		},
	}

	// FCP
	if !config.PercentualFCP.IsZero() {
		fcpPct := config.PercentualFCP.Div(decimal.NewFromInt(100))
		valorFCP := base.Mul(fcpPct).Round(2)
		tributo.MoreNumericDetails = append(tributo.MoreNumericDetails,
			models.Detalhe{Key: "fcp_percentual", Value: config.PercentualFCP},
			models.Detalhe{Key: "valor_fcp", Value: valorFCP},
		)
		tributo.Valor = tributo.Valor.Add(valorFCP)
	}

	itemRes.Tributos = append(itemRes.Tributos, tributo)

	slog.Debug("ICMS Proprio Interno Calculado",
		"SKU", item.SKU,
		"base", base,
		"aliquota", config.AliquotaInterna,
		"valor", valorICMS,
	)
}

func (c *ICMSCalculator) calcularICMSOperacaoInterestadual(
	itemRes *models.ItemDocumentoFiscalSaida,
	item models.ItemDocumentoFiscalEntrada,
	input models.DocumentoFiscalEntrada,
	config repository.ICMSConfig,
	baseItem decimal.Decimal,
) {
	aliqInterPct := config.AliquotaInterestadual.Div(decimal.NewFromInt(100))
	base := baseItem
	if !config.ReducaoBase.IsZero() {
		fatorReducao := decimal.NewFromInt(1).Sub(config.ReducaoBase.Div(decimal.NewFromInt(100)))
		base = base.Mul(fatorReducao)
	}

	// ICMS Interestadual (sempre calculado, exceto para Simples que é tratado em separado)
	valorICMSInter := base.Mul(aliqInterPct).Round(2)

	itemRes.Tributos = append(itemRes.Tributos, models.TributosItemDocumentoFiscalSaida{
		Tributo:     "ICMS",
		BaseCalculo: base,
		Aliquota:    config.AliquotaInterestadual,
		Valor:       valorICMSInter,
		CST:         config.CSTPadrao,
		MoreNumericDetails: []models.Detalhe{
			{Key: "valor_item", Value: baseItem},
			{Key: "reducao_base_percentual", Value: config.ReducaoBase},
			{Key: "base_calculo_efetiva", Value: base},
			{Key: "aliquota_interestadual", Value: config.AliquotaInterestadual},
			{Key: "valor_icms_interestadual", Value: valorICMSInter},
		},
		MoreTextDetails: []models.Detalhe{
			{Key: "fonte_regra", Value: "config_efetiva_icms_interestadual"},
			{Key: "uf_origem", Value: input.LocalizacaoOrigem.UF},
			{Key: "uf_destino", Value: input.LocalizacaoDestino.UF},
		},
	})

	slog.Debug("ICMS Interestadual Calculado",
		"SKU", item.SKU,
		"origem", input.LocalizacaoOrigem.UF,
		"destino", input.LocalizacaoDestino.UF,
		"aliquota", config.AliquotaInterestadual,
		"valor", valorICMSInter,
	)

	// DIFAL: Consumidor final + não Simples
	if input.IsDestinoFinal {
		difalRes, ok := c.CalcularDIFAL(item, config)
		if ok {
			itemRes.Tributos = append(itemRes.Tributos, difalRes)
			slog.Debug("DIFAL Aplicado",
				"SKU", item.SKU,
				"locationOrigem", input.LocalizacaoOrigem.UF,
				"locationDestino", input.LocalizacaoDestino.UF,
				"Valor", difalRes.Valor,
			)
		}
		return
	}

	// ICMS-ST interestadual: Protocolo ST + MVA configurado
	if config.PossuiProtocoloST && !config.MVAPadrao.IsZero() {
		mva := config.MVAPadrao
		aliquotaFator := mva.Div(decimal.NewFromInt(100)).Add(decimal.NewFromInt(1))
		baseST := baseItem.Mul(aliquotaFator)
		aliqIntPct := config.AliquotaInterna.Div(decimal.NewFromInt(100))
		valorST := baseST.Mul(aliqIntPct).Round(2)

		itemRes.Tributos = append(itemRes.Tributos, models.TributosItemDocumentoFiscalSaida{
			Tributo:     "ICMS_ST",
			BaseCalculo: baseST,
			Aliquota:    config.AliquotaInterna,
			Valor:       valorST,
			MoreNumericDetails: []models.Detalhe{
				{Key: "valor_item", Value: baseItem},
				{Key: "mva_percentual", Value: mva},
				{Key: "base_st_calculada", Value: baseST},
				{Key: "aliquota_interna_destino", Value: config.AliquotaInterna},
				{Key: "valor_st_calculado", Value: valorST},
			},
			MoreTextDetails: []models.Detalhe{
				{Key: "fonte_regra", Value: "config_efetiva_st_interestadual"},
				{Key: "uf_origem", Value: input.LocalizacaoOrigem.UF},
				{Key: "uf_destino", Value: input.LocalizacaoDestino.UF},
			},
		})

		slog.Debug("ICMS-ST Interestadual Calculado",
			"SKU", item.SKU,
			"mva", mva,
			"baseST", baseST,
			"valorST", valorST,
		)
	}
}

func (c *ICMSCalculator) CalcularDIFAL(item models.ItemDocumentoFiscalEntrada, config repository.ICMSConfig) (models.TributosItemDocumentoFiscalSaida, bool) {
	if config.AliquotaInterna.LessThanOrEqual(config.AliquotaInterestadual) {
		return models.TributosItemDocumentoFiscalSaida{}, false
	}

	aliquotaDifal := config.AliquotaInterna.Sub(config.AliquotaInterestadual)

	aliqPct := aliquotaDifal.Div(decimal.NewFromInt(100))
	valorDifal := item.ValorUnitario.Mul(item.Quantidade).Mul(aliqPct).Round(2)

	var valorF64, aliqIntF64, aliqInterstF64, aliqDifF64 float64
	valorF64, _ = item.ValorUnitario.Mul(item.Quantidade).Float64()
	aliqIntF64, _ = config.AliquotaInterna.Float64()
	aliqInterstF64, _ = config.AliquotaInterestadual.Float64()
	aliqDifF64, _ = aliquotaDifal.Float64()

	memoria := fmt.Sprintf(
		"Base: %.2f | Aliq. Interna Destino: %.2f%% | Aliq. Interestadual: %.2f%% | Diferencial: %.2f%%",
		valorF64, aliqIntF64, aliqInterstF64, aliqDifF64,
	)

	return models.TributosItemDocumentoFiscalSaida{
		Tributo:     "ICMS_DIFAL",
		BaseCalculo: item.ValorUnitario.Mul(item.Quantidade),
		Aliquota:    aliquotaDifal,
		Valor:       valorDifal,
		MoreTextDetails: []models.Detalhe{
			{Key: "MEMORIA_CALCULO", Value: memoria},
			{Key: "REGRA_APLICADA", Value: "EC 87/2015 - Consumidor Final"},
		},
	}, true
}

func (c *ICMSCalculator) getEffectiveTaxConfig(ctx context.Context, item models.ItemDocumentoFiscalEntrada, input models.DocumentoFiscalEntrada) (repository.ICMSConfig, error) {
	rule, err := c.repo.GetICMSRule(ctx, input.LocalizacaoOrigem.UF, input.LocalizacaoDestino.UF)
	if err != nil {
		slog.Error("Regra geral de ICMS nao encontrada", "origem", input.LocalizacaoOrigem.UF, "destino", input.LocalizacaoDestino.UF)
		return repository.ICMSConfig{}, err
	}

	ncmGroup := ""
	if len(item.NCM) >= 4 {
		ncmGroup = item.NCM[:4]
	}

	exc, err := c.repo.GetProductException(ctx, item.NCM, ncmGroup, input.LocalizacaoDestino.UF, string(models.NormalizeCRTEmitente(input.CRTEmitente)))

	config := repository.ICMSConfig{
		Id:                    rule.Id,
		AliquotaInterna:       rule.AliquotaInterna,
		AliquotaInterestadual: rule.AliquotaInterestadual,
		PercentualFCP:         rule.PercentualFCP,
		ReducaoBase:           rule.ReducaoBase,
		CSTPadrao:             string(models.CSTICMS000),
		PossuiProtocoloST:     false,
	}

	if err == nil && exc != nil {
		if !exc.AliquotaInternaDestino.IsZero() {
			config.AliquotaInterna = exc.AliquotaInternaDestino
		}
		if !exc.AliquotaInterestadual.IsZero() {
			config.AliquotaInterestadual = exc.AliquotaInterestadual
		}
		if !exc.PercentualFCP.IsZero() {
			config.PercentualFCP = exc.PercentualFCP
		}
		if !exc.ReducaoBase.IsZero() {
			config.ReducaoBase = exc.ReducaoBase
		}

		config.MVAPadrao = exc.MVAST
		config.PossuiProtocoloST = exc.PossuiProtocoloST
		config.CSTPadrao = exc.CSTICMS
	}

	return config, nil
}
