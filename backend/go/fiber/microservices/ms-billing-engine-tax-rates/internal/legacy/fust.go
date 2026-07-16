// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/legacy/fust.go
package legacy

import (
	"context"
	"log/slog"

	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

const (
	// fustAliquota é a alíquota do FUST (1%) conforme Lei 9.998/2000.
	fustAliquota = 1.0

	// fustTributoName é o nome do tributo na resposta.
	fustTributoName = "FUST"
)

// FUSTCalculator calcula o FUST (Fundo de Universalização dos Serviços
// de Telecomunicações), contribuição setorial de 1% sobre a Receita
// Operacional Líquida de serviços de telecom (SCM e STFC).
//
// Regras de negócio vinculadas:
//   - BR-TAX-CALC-019: FUST = Base_Líquida × 1%
//   - BR-TAX-DEF-010: Definição de FUST (Lei 9.998/2000)
//   - BR-TAX-CONS-012: Exclusivo para SCM/STFC — não incide sobre SVA
//   - BR-TAX-INF-007: Classificação de serviço (SCM/STFC vs SVA)
//
// A base de cálculo é a Receita Operacional Líquida:
//
//	Base_FUST = Valor_Serviço − ICMS − PIS − COFINS
//
// O FUST depende dos valores de ICMS, PIS e COFINS já calculados.
// Por isso, deve ser executado APÓS estes tributos no pipeline (Fase 3).
//
// Fonte: PROCEDURE-FIN-00001 SOP-016
type FUSTCalculator struct {
	classifier *TelecomClassifier
}

// NewFUSTCalculator cria uma nova calculadora de FUST.
func NewFUSTCalculator() *FUSTCalculator {
	return &FUSTCalculator{
		classifier: NewTelecomClassifier(),
	}
}

// Calculate implementa domain.TaxCalculator.
//
// Para cada item do documento fiscal:
//  1. Classifica o serviço (SCM/STFC vs SVA) — apenas SCM/STFC pagam FUST
//  2. Extrai os valores de ICMS, PIS e COFINS já calculados dos detalhes do item
//  3. Calcula base líquida: Base = Valor_Serviço − ICMS − PIS − COFINS
//  4. Se base líquida negativa → FUST = 0
//  5. Calcula: FUST = Base × 0,01 (1%)
//
// Os valores de ICMS, PIS e COFINS devem ter sido injetados nos detalhes
// do item pela engine ANTES da chamada a este calculator (Fase 3).
func (c *FUSTCalculator) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
	resultado := make([]models.ItemDocumentoFiscalSaida, len(input.Itens))

	for i, item := range input.Itens {
		resultado[i].SKU = item.SKU

		// 1. Classificar o serviço — apenas SCM/STFC pagam FUST
		if !c.classifier.MustCalculateFUST(item) {
			continue
		}

		itemMap := item.ToMap()
		valorServico := item.ValorUnitario.Mul(item.Quantidade)

		// 2. Extrair valores de ICMS, PIS e COFINS já calculados
		// Estes valores são injetados pela engine na Fase 2 (pós-paralela)
		// usando as chaves ITEM_ICMS_VALOR, ITEM_PIS_VALOR, ITEM_COFINS_VALOR
		icmsValor := models.GetDecimal(itemMap, models.KeyDocumentoInfos("ITEM_ICMS_VALOR"))
		pisValor := models.GetDecimal(itemMap, models.KeyDocumentoInfos("ITEM_PIS_VALOR"))
		cofinsValor := models.GetDecimal(itemMap, models.KeyDocumentoInfos("ITEM_COFINS_VALOR"))

		// 3. Calcular base líquida — BR-TAX-CALC-019
		baseLiquida := valorServico.Sub(icmsValor).Sub(pisValor).Sub(cofinsValor)

		// 4. Base líquida negativa → FUST = 0
		if baseLiquida.IsNegative() {
			slog.Warn("FUST: base líquida negativa, FUST = 0",
				"SKU", item.SKU,
				"valor_servico", valorServico,
				"icms", icmsValor,
				"pis", pisValor,
				"cofins", cofinsValor,
				"base_liquida", baseLiquida,
			)
			baseLiquida = decimal.Zero
		}

		// 5. Calcular FUST — BR-TAX-CALC-019
		aliquotaPct := decimal.NewFromFloat(fustAliquota).Div(decimal.NewFromInt(100))
		valorFUST := baseLiquida.Mul(aliquotaPct).Round(2)

		natureza, _ := c.classifier.Classify(item)

		tributo := models.TributosItemDocumentoFiscalSaida{
			Tributo:     fustTributoName,
			CST:         "", // FUST não utiliza CST
			BaseCalculo: baseLiquida.Round(2),
			Aliquota:    decimal.NewFromFloat(fustAliquota),
			Valor:       valorFUST,
			MoreNumericDetails: []models.Detalhe{
				{Key: "valor_servico", Value: valorServico},
				{Key: "icms_deduzido", Value: icmsValor},
				{Key: "pis_deduzido", Value: pisValor},
				{Key: "cofins_deduzido", Value: cofinsValor},
				{Key: "base_liquida", Value: baseLiquida.Round(2)},
				{Key: "aliquota_fust", Value: decimal.NewFromFloat(fustAliquota)},
				{Key: "valor_fust", Value: valorFUST},
			},
			MoreTextDetails: []models.Detalhe{
				{Key: "natureza_servico", Value: natureza},
				{Key: "sku", Value: item.SKU},
				{Key: "base_legal", Value: "Lei 9.998/2000"},
				{Key: "metodo", Value: "base_liquida"},
			},
		}

		resultado[i].Tributos = append(resultado[i].Tributos, tributo)

		slog.Debug("FUST calculado",
			"SKU", item.SKU,
			"natureza", natureza,
			"valor_servico", valorServico,
			"base_liquida", baseLiquida,
			"valor", valorFUST,
		)
	}

	return resultado, nil
}
