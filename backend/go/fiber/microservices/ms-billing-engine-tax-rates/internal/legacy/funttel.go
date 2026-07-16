// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/legacy/funttel.go
package legacy

import (
	"context"
	"log/slog"

	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

const (
	// funttelAliquota é a alíquota do FUNTTEL (0,5%) conforme Lei 10.052/2000.
	funttelAliquota = 0.5

	// funttelTributoName é o nome do tributo na resposta.
	funttelTributoName = "FUNTTEL"
)

// FUNTTELCalculator calcula o FUNTTEL (Fundo para o Desenvolvimento
// Tecnológico das Telecomunicações), contribuição setorial de 0,5%
// sobre a mesma base de cálculo do FUST.
//
// Regras de negócio vinculadas:
//   - BR-TAX-CALC-020: FUNTTEL = Base_FUST × 0,5%
//   - BR-TAX-DEF-011: Definição de FUNTTEL (Lei 10.052/2000)
//
// O FUNTTEL compartilha a mesma base de cálculo e a mesma lógica
// de classificação (SCM/STFC vs SVA) com o FUST. Por isso, deve ser
// executado imediatamente APÓS o FUST no pipeline.
//
// Fonte: PROCEDURE-FIN-00001 SOP-016
type FUNTTELCalculator struct {
	classifier *TelecomClassifier
}

// NewFUNTTELCalculator cria uma nova calculadora de FUNTTEL.
func NewFUNTTELCalculator() *FUNTTELCalculator {
	return &FUNTTELCalculator{
		classifier: NewTelecomClassifier(),
	}
}

// Calculate implementa domain.TaxCalculator.
//
// Utiliza a mesma base de cálculo do FUST:
//
//	Base = Valor_Serviço − ICMS − PIS − COFINS
//	FUNTTEL = Base × 0,005 (0,5%)
//
// Assim como o FUST, depende dos valores de ICMS, PIS e COFINS
// já calculados e injetados nos detalhes do item pela engine.
func (c *FUNTTELCalculator) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
	resultado := make([]models.ItemDocumentoFiscalSaida, len(input.Itens))

	for i, item := range input.Itens {
		resultado[i].SKU = item.SKU

		// 1. Mesma classificação do FUST — apenas SCM/STFC
		if !c.classifier.MustCalculateFUST(item) {
			continue
		}

		itemMap := item.ToMap()
		valorServico := item.ValorUnitario.Mul(item.Quantidade)

		// 2. Extrair valores já calculados (injetados pela engine)
		icmsValor := models.GetDecimal(itemMap, models.KeyDocumentoInfos("ITEM_ICMS_VALOR"))
		pisValor := models.GetDecimal(itemMap, models.KeyDocumentoInfos("ITEM_PIS_VALOR"))
		cofinsValor := models.GetDecimal(itemMap, models.KeyDocumentoInfos("ITEM_COFINS_VALOR"))

		// 3. Mesma base líquida do FUST — BR-TAX-CALC-020
		baseLiquida := valorServico.Sub(icmsValor).Sub(pisValor).Sub(cofinsValor)

		if baseLiquida.IsNegative() {
			slog.Warn("FUNTTEL: base líquida negativa, FUNTTEL = 0",
				"SKU", item.SKU,
				"base_liquida", baseLiquida,
			)
			baseLiquida = decimal.Zero
		}

		// 4. Calcular FUNTTEL — BR-TAX-CALC-020
		aliquotaPct := decimal.NewFromFloat(funttelAliquota).Div(decimal.NewFromInt(100))
		valorFUNTTEL := baseLiquida.Mul(aliquotaPct).Round(2)

		natureza, _ := c.classifier.Classify(item)

		tributo := models.TributosItemDocumentoFiscalSaida{
			Tributo:     funttelTributoName,
			CST:         "", // FUNTTEL não utiliza CST
			BaseCalculo: baseLiquida.Round(2),
			Aliquota:    decimal.NewFromFloat(funttelAliquota),
			Valor:       valorFUNTTEL,
			MoreNumericDetails: []models.Detalhe{
				{Key: "valor_servico", Value: valorServico},
				{Key: "icms_deduzido", Value: icmsValor},
				{Key: "pis_deduzido", Value: pisValor},
				{Key: "cofins_deduzido", Value: cofinsValor},
				{Key: "base_liquida", Value: baseLiquida.Round(2)},
				{Key: "aliquota_funttel", Value: decimal.NewFromFloat(funttelAliquota)},
				{Key: "valor_funttel", Value: valorFUNTTEL},
			},
			MoreTextDetails: []models.Detalhe{
				{Key: "natureza_servico", Value: natureza},
				{Key: "sku", Value: item.SKU},
				{Key: "base_legal", Value: "Lei 10.052/2000"},
				{Key: "metodo", Value: "base_liquida"},
			},
		}

		resultado[i].Tributos = append(resultado[i].Tributos, tributo)

		slog.Debug("FUNTTEL calculado",
			"SKU", item.SKU,
			"natureza", natureza,
			"base_liquida", baseLiquida,
			"valor", valorFUNTTEL,
		)
	}

	return resultado, nil
}
