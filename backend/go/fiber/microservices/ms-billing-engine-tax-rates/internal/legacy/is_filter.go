// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/legacy/is_filter.go
package legacy

import (
	"context"
	"log/slog"

	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"

	"github.com/shopspring/decimal"
)

const (
	isTributoName = "IS"
)

// ISFilter é uma pré-calculadora (Fase 0) que verifica se o NCM do item
// está sujeito ao Imposto Seletivo (IS) ANTES do cálculo da CBS/IBS.
//
// Diferente do cálculo de IS embutido no ReformaCalculator (que consulta
// iva_dual_rules), o ISFilter consulta a tabela independente ncm_seletivo
// conforme determinado pela política organizacional.
//
// Regras de negócio vinculadas:
//   - BR-TAX-INF-005: Consultar tabela ncm_seletivo para verificar NCM sujeito ao IS
//   - BR-TAX-CONS-010: IS é pré-filtro obrigatório ANTES da CBS
//
// Comportamento:
//  1. Se flag isento_is = true (documento ou item) → IS = 0 (override manual)
//  2. Se NCM na tabela ncm_seletivo → IS = Valor × Aliquota_IS_Categoria
//  3. Se NCM não na tabela → IS = 0 (não incide)
//  4. Registra em auditoria mesmo se IS = 0 (com flag is_exempt)
//
// Fonte: PROCEDURE-FIN-00001 SOP-003
type ISFilter struct {
	repo repository.TaxRepository
}

// NewISFilter cria um novo pré-filtro de Imposto Seletivo.
// O repositório é usado para consultar a tabela ncm_seletivo.
func NewISFilter(repo repository.TaxRepository) *ISFilter {
	return &ISFilter{repo: repo}
}

// Calculate implementa domain.TaxCalculator.
//
// É executado ANTES de todos os outros tributos (Fase 0 do pipeline).
// Os valores calculados são injetados nos detalhes do item para que
// calculadoras subsequentes possam acessar o valor do IS já apurado.
func (f *ISFilter) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
	resultado := make([]models.ItemDocumentoFiscalSaida, len(input.Itens))

	// Flag de isenção no nível do documento
	docMap := input.ToMap()
	isentoDoc := models.GetString(docMap, models.KeyDocumentoInfos("ISENTO_IS"))
	isIsentoGlobal := isentoDoc == "true" || isentoDoc == "1"

	for i, item := range input.Itens {
		resultado[i].SKU = item.SKU

		itemMap := item.ToMap()

		// 1. Verifica flag de isenção no nível do item
		isentoItem := models.GetString(itemMap, models.KeyDocumentoInfos("ISENTO_IS"))
		isIsento := isIsentoGlobal || isentoItem == "true" || isentoItem == "1"

		if isIsento {
			// IS isento — registra em auditoria com flag is_exempt
			slog.Debug("IS Filter: item isento de IS (flag isento_is)",
				"SKU", item.SKU,
				"NCM", item.NCM,
			)
			// Registra IS=0 com marcação de isenção para auditoria
			resultado[i].Tributos = append(resultado[i].Tributos, f.buildTributoIsento(item))
			continue
		}

		// 2. Consulta tabela ncm_seletivo
		rule, err := f.repo.GetNCMSeletivo(ctx, item.NCM)
		if err != nil {
			slog.Warn("IS Filter: erro ao consultar ncm_seletivo, assumindo IS=0",
				"SKU", item.SKU,
				"NCM", item.NCM,
				"error", err,
			)
			resultado[i].Tributos = append(resultado[i].Tributos, f.buildTributoIsento(item))
			continue
		}

		// 3. NCM não está na tabela → IS não incide
		if rule == nil {
			slog.Debug("IS Filter: NCM não sujeito ao IS",
				"SKU", item.SKU,
				"NCM", item.NCM,
			)
			resultado[i].Tributos = append(resultado[i].Tributos, f.buildTributoNaoIncide(item))
			continue
		}

		// 4. NCM na tabela → IS incide
		valorItem := item.ValorUnitario.Mul(item.Quantidade)
		aliquotaPct := rule.AliquotaIS.Div(decimal.NewFromInt(100))
		valorIS := valorItem.Mul(aliquotaPct).Round(2)

		tributo := models.TributosItemDocumentoFiscalSaida{
			Tributo:     isTributoName,
			CST:         "", // IS não utiliza CST no formato tradicional
			BaseCalculo: valorItem,
			Aliquota:    rule.AliquotaIS,
			Valor:       valorIS,
			MoreNumericDetails: []models.Detalhe{
				{Key: "valor_item", Value: valorItem},
				{Key: "aliquota_is_categoria", Value: rule.AliquotaIS},
				{Key: "valor_is", Value: valorIS},
			},
			MoreTextDetails: []models.Detalhe{
				{Key: "ncm", Value: item.NCM},
				{Key: "sku", Value: item.SKU},
				{Key: "categoria_is", Value: rule.Categoria},
				{Key: "fonte", Value: "ncm_seletivo"},
				{Key: "is_exempt", Value: "false"},
				{Key: "base_legal", Value: "EC 132/2023"},
			},
		}

		resultado[i].Tributos = append(resultado[i].Tributos, tributo)

		slog.Debug("IS Filter: IS calculado",
			"SKU", item.SKU,
			"NCM", item.NCM,
			"categoria", rule.Categoria,
			"aliquota", rule.AliquotaIS,
			"valor", valorIS,
		)
	}

	return resultado, nil
}

// buildTributoIsento cria um registro de IS=0 com flag de isenção.
// Utilizado quando isento_is=true (override manual de isenção).
func (f *ISFilter) buildTributoIsento(item models.ItemDocumentoFiscalEntrada) models.TributosItemDocumentoFiscalSaida {
	return models.TributosItemDocumentoFiscalSaida{
		Tributo:     isTributoName,
		CST:         "",
		BaseCalculo: decimal.Zero,
		Aliquota:    decimal.Zero,
		Valor:       decimal.Zero,
		MoreNumericDetails: []models.Detalhe{
			{Key: "valor_is", Value: decimal.Zero},
		},
		MoreTextDetails: []models.Detalhe{
			{Key: "ncm", Value: item.NCM},
			{Key: "sku", Value: item.SKU},
			{Key: "fonte", Value: "isento_override"},
			{Key: "is_exempt", Value: "true"},
			{Key: "motivo", Value: "Flag isento_is ativa — isencao manual"},
		},
	}
}

// buildTributoNaoIncide cria um registro de IS=0 quando o NCM não está
// na tabela ncm_seletivo (não é produto sujeito ao IS).
func (f *ISFilter) buildTributoNaoIncide(item models.ItemDocumentoFiscalEntrada) models.TributosItemDocumentoFiscalSaida {
	return models.TributosItemDocumentoFiscalSaida{
		Tributo:     isTributoName,
		CST:         "",
		BaseCalculo: decimal.Zero,
		Aliquota:    decimal.Zero,
		Valor:       decimal.Zero,
		MoreNumericDetails: []models.Detalhe{
			{Key: "valor_is", Value: decimal.Zero},
		},
		MoreTextDetails: []models.Detalhe{
			{Key: "ncm", Value: item.NCM},
			{Key: "sku", Value: item.SKU},
			{Key: "fonte", Value: "ncm_seletivo"},
			{Key: "is_exempt", Value: "true"},
			{Key: "motivo", Value: "NCM nao listado na tabela ncm_seletivo"},
		},
	}
}
