// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/calculator/legacy_adapter.go
package calculator

import (
	"context"
	"fmt"

	"ms-billing-engine-tax-rates/internal/domain"
	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

type legacyICMSCalculator interface {
	Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) (models.DocumentoFiscalSaida, error)
}

type legacyPISCofinsCalculator interface {
	Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.TributosItemDocumentoFiscalSaida, error)
}

// icmsTaxAdapter adapta ICMSCalculator para a interface TaxCalculator
type icmsTaxAdapter struct {
	calculator legacyICMSCalculator
}

func (a icmsTaxAdapter) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
	response, err := a.calculator.Calculate(ctx, input)
	if err != nil {
		return nil, err
	}
	return response.Itens, nil
}

// legacyTaxAdapter adapta PISCofinsCalculator para a interface TaxCalculator,
// com injeção opcional de valores de ICMS para exclusão da base de cálculo.
type legacyTaxAdapter struct {
	calculator legacyPISCofinsCalculator
	icmsSource legacyICMSCalculator
}

func (a legacyTaxAdapter) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
	// 1. Mapeia valores de ICMS por SKU para a exclusão da base
	icmsValues := make(map[string]decimal.Decimal)
	if a.icmsSource != nil {
		icmsRes, _ := a.icmsSource.Calculate(ctx, input)
		for _, item := range icmsRes.Itens {
			for _, trib := range item.Tributos {
				if trib.Tributo == "ICMS" || trib.Tributo == "ICMS_PROPRIO" {
					icmsValues[item.SKU] = trib.Valor
				}
			}
		}
	}

	// 2. Injeta o valor no input para o PIS/COFINS enxergar
	for i := range input.Itens {
		if val, ok := icmsValues[input.Itens[i].SKU]; ok {
			input.Itens[i].AddDetalhe(models.KeyDocumentoInfosValorExclusaoICMS, val)
		}
	}

	tributos, err := a.calculator.Calculate(ctx, input)
	if err != nil {
		return nil, err
	}

	// 3. Formata o retorno
	res := make([]models.ItemDocumentoFiscalSaida, len(input.Itens))
	for i, item := range input.Itens {
		res[i].SKU = item.SKU
		for _, t := range tributos {
			res[i].Tributos = append(res[i].Tributos, t)
		}
	}

	return res, nil
}

// NewLegacyAdapter cria adaptadores para calculadoras legadas.
// Uso para ICMS:              LegacyAdapter(icmsCalc)
// Uso para PIS/COFINS:        LegacyAdapter(pisCofinsCalc)
// Uso com exclusão de ICMS:   LegacyAdapter(pisCofinsCalc, icmsCalc)
func LegacyAdapter(calc interface{}, icmsProvider ...interface{}) domain.TaxCalculator {
	if typed, ok := calc.(legacyICMSCalculator); ok {
		return icmsTaxAdapter{calculator: typed}
	}

	if typed, ok := calc.(legacyPISCofinsCalculator); ok {
		adapter := legacyTaxAdapter{calculator: typed}
		if len(icmsProvider) > 0 {
			if provider, providerOK := icmsProvider[0].(legacyICMSCalculator); providerOK {
				adapter.icmsSource = provider
			}
		}
		return adapter
	}

	panic(fmt.Sprintf("LegacyAdapter: tipo não suportado: %T", calc))
}
