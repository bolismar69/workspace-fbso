package reforma

import (
	"context"
	"log/slog"

	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"
)

// IBSCalculator calcula apenas o IBS (Imposto sobre Bens e Servicos),
// tributo subnacional (estadual + municipal) da Reforma Tributaria (EC 132/2023).
//
// No pipeline SOP-013 (C-001), o IBS executa na Fase 4 (paralela), junto com
// ISS e PIS/COFINS. O IBS e independente destes tributos e pode ser calculado
// concorrentemente.
//
// O IBS substitui ICMS (estadual) e ISS (municipal) a partir de 2033
// (fase IVA Dual), com transicao gradual entre 2029–2032.
type IBSCalculator struct {
	repo repository.TaxRepository
}

// NewIBSCalculator cria uma calculadora exclusiva de IBS.
func NewIBSCalculator(r repository.TaxRepository) *IBSCalculator {
	return &IBSCalculator{repo: r}
}

// Calculate computa apenas o IBS para todos os itens do documento.
// CBS e IS nao sao calculados aqui — use CBSCalculator e ISFilter respectivamente.
func (c *IBSCalculator) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
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
			slog.Debug("Item isento de IBS (reducao 100%)",
				"sku", item.SKU,
				"ncm", item.NCM,
			)
			continue
		}

		numDetails, textDetails := buildIBSDetails(r)
		res[i].Tributos = append(res[i].Tributos, models.TributosItemDocumentoFiscalSaida{
			Tributo:           tributoIBS,
			CST:               r.CSTEfetivo,
			BaseCalculo:       r.BaseCalculo,
			Aliquota:          r.AliquotaEfetivaIBS,
			Valor:             r.ValorIBSTotal,
			MoreNumericDetails: numDetails,
			MoreTextDetails:   textDetails,
		})
	}

	return res, nil
}
