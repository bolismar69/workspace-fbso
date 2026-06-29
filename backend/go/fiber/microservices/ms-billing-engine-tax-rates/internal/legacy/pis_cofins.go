package legacy

import (
	"context"
	"log/slog"
	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"

	"github.com/shopspring/decimal"
)

const (
	defaultAliquotaPIS    = 1.65
	defaultAliquotaCOFINS = 7.6
)

type PISCofinsCalculator struct {
	repo repository.TaxRepository
}

func NewPISCofinsCalculator(r repository.TaxRepository) *PISCofinsCalculator {
	return &PISCofinsCalculator{repo: r}
}

func (c *PISCofinsCalculator) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.TributosItemDocumentoFiscalSaida, error) {
	var tributos []models.TributosItemDocumentoFiscalSaida

	regime := string(models.NormalizeCRTEmitente(input.CRTEmitente))

	for _, item := range input.Itens {
		attrs := item.ToMap()
		slog.Debug("Calculando PIS/COFINS",
			"SKU", item.SKU,
			"Valor", item.ValorUnitario.Mul(item.Quantidade),
			"regime", regime,
		)

		cstPis := models.GetString(attrs, models.KeyDocumentoInfosItemPISCOFINSCSTPIS)
		cstCofins := models.GetString(attrs, models.KeyDocumentoInfosItemPISCOFINSCSTCOFINS)

		valorIcmsExcluir := models.GetDecimal(attrs, models.KeyDocumentoInfosValorExclusaoICMS)

		// 1. Busca regra federal no banco
		federalRule, err := c.repo.GetFederalTaxRule(ctx, regime, cstPis, cstCofins)
		if err != nil {
			slog.Warn("Regra federal nao encontrada no banco, usando defaults",
				"regime", regime,
				"cst_pis", cstPis,
				"cst_cofins", cstCofins,
				"erro", err,
			)
		}

		// 2. Determina alíquotas: banco > default hardcoded
		aliqPis := decimal.NewFromFloat(defaultAliquotaPIS)
		aliqCofins := decimal.NewFromFloat(defaultAliquotaCOFINS)
		excluiICMSBase := true

		if federalRule != nil {
			if !federalRule.AliquotaPIS.IsZero() {
				aliqPis = federalRule.AliquotaPIS
			}
			if !federalRule.AliquotaCOFINS.IsZero() {
				aliqCofins = federalRule.AliquotaCOFINS
			}
			excluiICMSBase = federalRule.ExcluiICMSBase
		}

		slog.Debug("Aliquotas PIS/COFINS resolvidas",
			"SKU", item.SKU,
			"aliq_pis", aliqPis,
			"aliq_cofins", aliqCofins,
			"exclui_icms_base", excluiICMSBase,
			"fonte", map[bool]string{true: "banco", false: "default"}[federalRule != nil],
		)

		// 3. PIS
		if cstPis != "" {
			valorIcmsParaExcluir := decimal.Zero
			if excluiICMSBase {
				valorIcmsParaExcluir = valorIcmsExcluir
			}

			basePis := item.ValorUnitario.Mul(item.Quantidade).Sub(valorIcmsParaExcluir)

			strategy := GetPISStrategy(cstPis)
			info := PISInfo{
				CST:       cstPis,
				Regime:    input.CRTEmitente,
				Aliquota:  aliqPis,
				ValorIcms: valorIcmsParaExcluir,
			}

			valorCalculado := strategy.Calculate(item, info)

			tributos = append(tributos, models.TributosItemDocumentoFiscalSaida{
				Tributo:     "PIS",
				CST:         cstPis,
				BaseCalculo: basePis,
				Aliquota:    aliqPis,
				Valor:       valorCalculado,
				MoreNumericDetails: []models.Detalhe{
					{Key: "valor_item", Value: item.ValorUnitario.Mul(item.Quantidade)},
					{Key: "icms_excluido_base", Value: valorIcmsParaExcluir},
					{Key: "base_calculo_final", Value: basePis},
					{Key: "aliquota_percentual", Value: aliqPis},
					{Key: "valor_tributo", Value: valorCalculado},
				},
				MoreTextDetails: []models.Detalhe{
					{Key: "metodo", Value: map[bool]string{true: "base_com_exclusao_icms", false: "base_sem_exclusao_icms"}[excluiICMSBase]},
					{Key: "regime", Value: input.CRTEmitente},
					{Key: "sku", Value: item.SKU},
					{Key: "fonte_aliquota", Value: map[bool]string{true: "banco_federal_tax_rules", false: "default_hardcoded"}[federalRule != nil]},
				},
			})
		}

		// 4. COFINS
		if cstCofins != "" {
			valorIcmsParaExcluir := decimal.Zero
			if excluiICMSBase {
				valorIcmsParaExcluir = valorIcmsExcluir
			}

			baseCofins := item.ValorUnitario.Mul(item.Quantidade).Sub(valorIcmsParaExcluir)

			strategy := GetCOFINSStrategy(cstCofins)
			info := COFINSInfo{
				CST:       cstCofins,
				Aliquota:  aliqCofins,
				ValorIcms: valorIcmsParaExcluir,
			}

			valorCalculado := strategy.Calculate(item, info)

			tributos = append(tributos, models.TributosItemDocumentoFiscalSaida{
				Tributo:     "COFINS",
				CST:         cstCofins,
				BaseCalculo: baseCofins,
				Aliquota:    aliqCofins,
				Valor:       valorCalculado,
				MoreNumericDetails: []models.Detalhe{
					{Key: "valor_item", Value: item.ValorUnitario.Mul(item.Quantidade)},
					{Key: "icms_excluido_base", Value: valorIcmsParaExcluir},
					{Key: "base_calculo_final", Value: baseCofins},
					{Key: "aliquota_percentual", Value: aliqCofins},
					{Key: "valor_tributo", Value: valorCalculado},
				},
				MoreTextDetails: []models.Detalhe{
					{Key: "metodo", Value: map[bool]string{true: "base_com_exclusao_icms", false: "base_sem_exclusao_icms"}[excluiICMSBase]},
					{Key: "regime", Value: input.CRTEmitente},
					{Key: "sku", Value: item.SKU},
					{Key: "fonte_aliquota", Value: map[bool]string{true: "banco_federal_tax_rules", false: "default_hardcoded"}[federalRule != nil]},
				},
			})
		}
	}

	return tributos, nil
}

func GetPISStrategy(cst string) PISStrategy {
	switch models.NormalizeCSTPISCOFINS(cst) {
	case models.CSTPISCOFINS01, models.CSTPISCOFINS02:
		return &PIS01_02{}
	case models.CSTPISCOFINS03:
		return &PIS03{}
	case models.CSTPISCOFINS04:
		return &PIS04{}
	case models.CSTPISCOFINS05:
		return &PIS05{}
	case models.CSTPISCOFINS06:
		return &PIS06{}
	case models.CSTPISCOFINS49:
		return &PIS49{}
	case models.CSTPISCOFINS99:
		return &PIS99{}
	default:
		return &PIS50To99{}
	}
}

func GetCOFINSStrategy(cst string) COFINSStrategy {
	switch models.NormalizeCSTPISCOFINS(cst) {
	case models.CSTPISCOFINS01, models.CSTPISCOFINS02:
		return &COFINS01_02{}
	case models.CSTPISCOFINS03:
		return &COFINS03{}
	case models.CSTPISCOFINS04:
		return &COFINS04{}
	case models.CSTPISCOFINS05:
		return &COFINS05{}
	case models.CSTPISCOFINS06:
		return &COFINS06{}
	case models.CSTPISCOFINS49:
		return &COFINS49{}
	default:
		return &COFINS50To99{}
	}
}
