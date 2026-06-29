// Package reforma implementa as calculadoras da Reforma Tributaria Brasileira
// (EC 132/2023, LC 214/2025).
//
// O pacote contem tres calculadoras:
//   - ReformaCalculator: calcula CBS e IBS juntos (legado, mantido para compatibilidade)
//   - CBSCalculator: calcula apenas CBS (Fase 2 do pipeline SOP-013 — sequencial, "por fora")
//   - IBSCalculator: calcula apenas IBS (Fase 4 do pipeline SOP-013 — paralela)
//
// A separacao CBS/IBS e necessaria porque o pipeline SOP-013 (C-001)
// requer que CBS execute sequencialmente ANTES do ICMS (Fase 2), enquanto
// IBS executa em paralelo com ISS e PIS/COFINS (Fase 4).
//
// Ambos consultam a mesma tabela iva_dual_rules; o CachedTaxRepository
// garante que a segunda consulta atinja o cache Redis.
package reforma

import (
	"context"
	"log/slog"

	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"

	"github.com/shopspring/decimal"
)

const (
	tributoCBS = "CBS"
	tributoIBS = "IBS"
	cstPadrao  = "01"
	cstIsento  = "04"
)

// ivaDualResult contem os valores computados a partir de uma regra IVA Dual.
type ivaDualResult struct {
	BaseCalculo            decimal.Decimal
	AliquotaEfetivaCBS     decimal.Decimal
	AliquotaEfetivaIBS     decimal.Decimal
	AliquotaNominalCBS     decimal.Decimal
	AliquotaNominalIBSEst  decimal.Decimal
	AliquotaNominalIBSMun  decimal.Decimal
	ValorCBS               decimal.Decimal
	ValorIBSEstadual       decimal.Decimal
	ValorIBSMunicipal      decimal.Decimal
	ValorIBSTotal          decimal.Decimal
	FatorReducao           decimal.Decimal
	PercentualReducao      decimal.Decimal
	EfetivamenteIsento     bool
	CSTEfetivo             string
	UFDestino              string
	NCM                    string
	SKU                    string
}

// computeIvaDual consulta a tabela iva_dual_rules e computa os valores de CBS e IBS.
// Retorna nil se o item nao estiver sujeito a CBS/IBS (sem regra, NCM fora da tabela, etc.).
func computeIvaDual(ctx context.Context, repo repository.TaxRepository, item models.ItemDocumentoFiscalEntrada, ufDestino, municipioIBGE string) *ivaDualResult {
	ncm := item.NCM

	rule, err := repo.GetIvaDualRule(ctx, ncm, ufDestino, municipioIBGE)
	if err != nil {
		slog.Warn("Erro ao buscar regra IVA Dual, item sera ignorado",
			"sku", item.SKU,
			"ncm", ncm,
			"uf_destino", ufDestino,
			"municipio", municipioIBGE,
			"error", err,
		)
		return nil
	}

	if rule == nil {
		slog.Warn("Nenhuma regra IVA Dual encontrada, item ignorado",
			"sku", item.SKU,
			"ncm", ncm,
			"uf_destino", ufDestino,
			"municipio", municipioIBGE,
		)
		return nil
	}

	baseCalculo := item.ValorUnitario.Mul(item.Quantidade)

	fatorReducao := decimal.NewFromInt(1).Sub(
		rule.PercentualReducao.Div(decimal.NewFromInt(100)),
	)

	efetivamenteIsento := rule.PercentualReducao.GreaterThanOrEqual(decimal.NewFromInt(100))

	result := &ivaDualResult{
		BaseCalculo:       baseCalculo,
		FatorReducao:      fatorReducao,
		PercentualReducao: rule.PercentualReducao,
		EfetivamenteIsento: efetivamenteIsento,
		AliquotaNominalCBS: rule.AliquotaCBS,
		AliquotaNominalIBSEst: rule.AliquotaIBSEstadual,
		AliquotaNominalIBSMun: rule.AliquotaIBSMunicipal,
		UFDestino:         ufDestino,
		NCM:               ncm,
		SKU:               item.SKU,
	}

	if efetivamenteIsento {
		result.CSTEfetivo = cstIsento
		return result
	}

	result.CSTEfetivo = cstPadrao
	if rule.PercentualReducao.GreaterThan(decimal.Zero) {
		result.CSTEfetivo = cstIsento
	}

	result.AliquotaEfetivaCBS = rule.AliquotaCBS.Mul(fatorReducao)
	result.AliquotaEfetivaIBS = rule.AliquotaIBSEstadual.Add(rule.AliquotaIBSMunicipal).Mul(fatorReducao)

	aliquotaEfetivaIBSEstadual := rule.AliquotaIBSEstadual.Mul(fatorReducao)
	aliquotaEfetivaIBSMunicipal := rule.AliquotaIBSMunicipal.Mul(fatorReducao)

	result.ValorCBS = baseCalculo.Mul(result.AliquotaEfetivaCBS.Div(decimal.NewFromInt(100)))
	result.ValorIBSEstadual = baseCalculo.Mul(aliquotaEfetivaIBSEstadual.Div(decimal.NewFromInt(100)))
	result.ValorIBSMunicipal = baseCalculo.Mul(aliquotaEfetivaIBSMunicipal.Div(decimal.NewFromInt(100)))
	result.ValorIBSTotal = result.ValorIBSEstadual.Add(result.ValorIBSMunicipal)

	return result
}

// buildCBSDetails constroi os detalhes (numeric e text) para o tributo CBS.
func buildCBSDetails(r *ivaDualResult) ([]models.Detalhe, []models.Detalhe) {
	return []models.Detalhe{
			{Key: "valor_item", Value: r.BaseCalculo},
			{Key: "aliquota_nominal_cbs", Value: r.AliquotaNominalCBS},
			{Key: "aliquota_efetiva_cbs", Value: r.AliquotaEfetivaCBS},
			{Key: "fator_reducao", Value: r.FatorReducao},
			{Key: "percentual_reducao", Value: r.PercentualReducao},
		}, []models.Detalhe{
			{Key: "uf_destino", Value: r.UFDestino},
			{Key: "ncm", Value: r.NCM},
			{Key: "sku", Value: r.SKU},
			{Key: "fonte", Value: "iva_dual_rules"},
		}
}

// buildIBSDetails constroi os detalhes (numeric e text) para o tributo IBS.
func buildIBSDetails(r *ivaDualResult) ([]models.Detalhe, []models.Detalhe) {
	return []models.Detalhe{
			{Key: "valor_item", Value: r.BaseCalculo},
			{Key: "ibs_estadual", Value: r.ValorIBSEstadual},
			{Key: "ibs_municipal", Value: r.ValorIBSMunicipal},
			{Key: "aliquota_ibs_estadual_efetiva", Value: decimal.Zero}, // computed below if needed
			{Key: "aliquota_ibs_municipal_efetiva", Value: decimal.Zero},
			{Key: "fator_reducao", Value: r.FatorReducao},
			{Key: "percentual_reducao", Value: r.PercentualReducao},
		}, []models.Detalhe{
			{Key: "uf_destino", Value: r.UFDestino},
			{Key: "ncm", Value: r.NCM},
			{Key: "sku", Value: r.SKU},
			{Key: "fonte", Value: "iva_dual_rules"},
		}
}

// NOTA: O cálculo do IS (Imposto Seletivo) foi extraído para ISFilter
// (internal/legacy/is_filter.go) como pré-filtro independente (Fase 0).
// O ReformaCalculator agora calcula APENAS CBS e IBS.
// Ref: BR-TAX-CONS-010, F-006

// ReformaCalculator calcula CBS e IBS juntos (comportamento legado).
// Mantido para compatibilidade com codigo existente.
// Para o pipeline SOP-013 (C-001), prefira CBSCalculator e IBSCalculator separadamente.
type ReformaCalculator struct {
	repo repository.TaxRepository
}

// NewReformaCalculator cria uma calculadora legada de CBS + IBS.
func NewReformaCalculator(r repository.TaxRepository) *ReformaCalculator {
	return &ReformaCalculator{repo: r}
}

// Calculate computa CBS e IBS para todos os itens do documento.
func (c *ReformaCalculator) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
	res := make([]models.ItemDocumentoFiscalSaida, len(input.Itens))

	for i, item := range input.Itens {
		res[i].SKU = item.SKU

		ufDestino := input.LocalizacaoDestino.UF
		municipioIBGE := input.LocalizacaoDestino.Municipio

		r := computeIvaDual(ctx, c.repo, item, ufDestino, municipioIBGE)
		if r == nil {
			continue
		}

		if r.EfetivamenteIsento {
			slog.Debug("Item isento de CBS/IBS (reducao 100%)",
				"sku", item.SKU,
				"ncm", item.NCM,
			)
			continue
		}

		numDetails, textDetails := buildCBSDetails(r)
		res[i].Tributos = append(res[i].Tributos, models.TributosItemDocumentoFiscalSaida{
			Tributo:           tributoCBS,
			CST:               r.CSTEfetivo,
			BaseCalculo:       r.BaseCalculo,
			Aliquota:          r.AliquotaEfetivaCBS,
			Valor:             r.ValorCBS,
			MoreNumericDetails: numDetails,
			MoreTextDetails:   textDetails,
		})

		numDetailsIBS, textDetailsIBS := buildIBSDetails(r)
		res[i].Tributos = append(res[i].Tributos, models.TributosItemDocumentoFiscalSaida{
			Tributo:           tributoIBS,
			CST:               r.CSTEfetivo,
			BaseCalculo:       r.BaseCalculo,
			Aliquota:          r.AliquotaEfetivaIBS,
			Valor:             r.ValorIBSTotal,
			MoreNumericDetails: numDetailsIBS,
			MoreTextDetails:   textDetailsIBS,
		})
	}

	return res, nil
}
