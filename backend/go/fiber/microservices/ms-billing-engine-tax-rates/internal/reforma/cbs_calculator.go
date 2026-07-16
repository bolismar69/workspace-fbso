package reforma

import (
	"context"
	"log/slog"

	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"
)

// CBSCalculator calcula apenas a CBS (Contribuicao sobre Bens e Servicos),
// tributo federal da Reforma Tributaria (EC 132/2023).
//
// No pipeline SOP-013 (C-001), a CBS executa na Fase 2 (sequencial, "por fora"),
// ANTES do ICMS. CBS e "por fora" porque seu valor nao compoe a base de
// calculo de outros tributos.
//
// A CBS substitui PIS e COFINS a partir de 2027 (fase CBS Plena).
type CBSCalculator struct {
	repo repository.TaxRepository
}

// NewCBSCalculator cria uma calculadora exclusiva de CBS.
func NewCBSCalculator(r repository.TaxRepository) *CBSCalculator {
	return &CBSCalculator{repo: r}
}

// Calculate computa apenas a CBS para todos os itens do documento.
// IBS e IS nao sao calculados aqui — use IBSCalculator e ISFilter respectivamente.
func (c *CBSCalculator) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
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
			slog.Debug("Item isento de CBS (reducao 100%)",
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
	}

	return res, nil
}
