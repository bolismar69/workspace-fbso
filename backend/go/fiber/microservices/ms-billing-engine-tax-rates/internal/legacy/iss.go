// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/legacy/iss.go
package legacy

import (
	"context"
	"log/slog"

	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

const (
	// issAliquotaMin e issAliquotaMax definem o range constitucional
	// de alíquotas do ISS por município, conforme LC 116/2003 art. 8º-A.
	issAliquotaMin = 2.0
	issAliquotaMax = 5.0

	// issItemListaServicoTelecom é o código da Lista de Serviços da LC 116/2003
	// para serviços de telecomunicação (item 1.05).
	issItemListaServicoTelecom = "1.05"

	// issTributoName é o nome do tributo na resposta.
	issTributoName = "ISS"
)

// ISSCalculator calcula o ISS (Imposto Sobre Serviços de Qualquer Natureza)
// sobre serviços de telecomunicação e outros serviços.
//
// Regras de negócio vinculadas:
//   - BR-TAX-CALC-016: ISS = Preço_Serviço × Alíquota_Municipal
//   - BR-TAX-CONS-007: 2% ≤ Alíquota_Municipal ≤ 5%
//   - BR-TAX-DEF-007: Serviços de telecom = item 1.05 da LC 116/2003
//
// Fonte: PROCEDURE-FIN-00001 SOP-010
type ISSCalculator struct{}

// NewISSCalculator cria uma nova calculadora de ISS.
// O ISS não requer consulta ao banco de dados — a alíquota municipal
// é fornecida via detalhes do documento ou do item.
func NewISSCalculator() *ISSCalculator {
	return &ISSCalculator{}
}

// Calculate implementa domain.TaxCalculator.
//
// Para cada item do documento fiscal:
//  1. Verifica se o item é um serviço (via detalhe ITEM_LISTA_SERVICO)
//  2. Se não for serviço → skip (ISS não incide sobre mercadorias)
//  3. Obtém a alíquota municipal (documento > item > default)
//  4. Valida que a alíquota está no intervalo [2%, 5%] — BR-TAX-CONS-007
//  5. Calcula ISS = Preço_Serviço × Alíquota — BR-TAX-CALC-016
//  6. Verifica retenção na fonte (ISS retido pelo tomador)
func (c *ISSCalculator) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
	resultado := make([]models.ItemDocumentoFiscalSaida, len(input.Itens))

	// Alíquota municipal no nível do documento (pode ser sobrescrita por item)
	docMap := input.ToMap()
	aliquotaDoc := models.GetDecimal(docMap, models.KeyDocumentoInfos("ISS_ALIQUOTA"))

	for i, item := range input.Itens {
		resultado[i].SKU = item.SKU

		itemMap := item.ToMap()

		// 1. Verifica se o item é um serviço
		listaServico := models.GetString(itemMap, models.KeyDocumentoInfos("ITEM_LISTA_SERVICO"))
		if listaServico == "" {
			// Item sem classificação de serviço — ISS não incide
			slog.Debug("ISS: item sem ItemListaServico, pulando",
				"SKU", item.SKU,
			)
			continue
		}

		// 2. Obtém alíquota municipal: item > documento > default do município
		aliquota := models.GetDecimal(itemMap, models.KeyDocumentoInfos("ISS_ALIQUOTA"))
		if !aliquota.IsPositive() {
			aliquota = aliquotaDoc
		}
		if !aliquota.IsPositive() {
			// Sem alíquota configurada — não é possível calcular
			slog.Warn("ISS: alíquota municipal não configurada, pulando item",
				"SKU", item.SKU,
				"lista_servico", listaServico,
			)
			continue
		}

		// 3. Valida range constitucional [2%, 5%] — BR-TAX-CONS-007
		if aliquota.LessThan(decimal.NewFromFloat(issAliquotaMin)) || aliquota.GreaterThan(decimal.NewFromFloat(issAliquotaMax)) {
			slog.Warn("ISS: alíquota fora do intervalo legal [2%, 5%]",
				"SKU", item.SKU,
				"aliquota", aliquota,
			)
			// Continua o cálculo mesmo com warning — a responsabilidade
			// pela correção da alíquota é do Tax Compliance Officer
		}

		// 4. Calcula o valor do ISS — BR-TAX-CALC-016
		precoServico := item.ValorUnitario.Mul(item.Quantidade)
		aliquotaPct := aliquota.Div(decimal.NewFromInt(100))
		valorISS := precoServico.Mul(aliquotaPct).Round(2)

		// 5. Verifica retenção na fonte
		issRetido := models.GetString(itemMap, models.KeyDocumentoInfos("ISS_RETIDO"))
		retencaoFonte := issRetido == "true" || issRetido == "1"

		fonteAliquota := "documento"
		if models.GetDecimal(itemMap, models.KeyDocumentoInfos("ISS_ALIQUOTA")).IsPositive() {
			fonteAliquota = "item"
		}

		tributo := models.TributosItemDocumentoFiscalSaida{
			Tributo:     issTributoName,
			CST:         "", // ISS não utiliza CST no mesmo formato do ICMS/IPI
			BaseCalculo: precoServico,
			Aliquota:    aliquota,
			Valor:       valorISS,
			MoreNumericDetails: []models.Detalhe{
				{Key: "preco_servico", Value: precoServico},
				{Key: "aliquota_municipal", Value: aliquota},
				{Key: "valor_iss", Value: valorISS},
			},
			MoreTextDetails: []models.Detalhe{
				{Key: "lista_servico", Value: listaServico},
				{Key: "fonte_aliquota", Value: fonteAliquota},
				{Key: "retencao_fonte", Value: map[bool]string{true: "sim", false: "nao"}[retencaoFonte]},
				{Key: "uf_prestador", Value: input.LocalizacaoOrigem.UF},
				{Key: "sku", Value: item.SKU},
				{Key: "base_legal", Value: "LC 116/2003"},
			},
		}

		resultado[i].Tributos = append(resultado[i].Tributos, tributo)

		slog.Debug("ISS calculado",
			"SKU", item.SKU,
			"lista_servico", listaServico,
			"preco_servico", precoServico,
			"aliquota", aliquota,
			"valor", valorISS,
			"retencao_fonte", retencaoFonte,
		)
	}

	return resultado, nil
}
