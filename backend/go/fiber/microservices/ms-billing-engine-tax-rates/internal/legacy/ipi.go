// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/legacy/ipi.go
package legacy

import (
	"context"
	"log/slog"
	"strings"

	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"

	"github.com/shopspring/decimal"
)

// IPICalculator calcula o IPI (Imposto sobre Produtos Industrializados).
// Deve ser executado ANTES dos demais impostos, pois seu valor compõe a base
// de cálculo do ICMS em operações com consumidor final / ativo imobilizado.
//
// Prioridade de parâmetros (da maior para a menor):
//  1. Detalhes do item (DetalhesItemDocumentoFiscal) — permitem sobrescrever
//     qualquer campo da regra sem necessidade de cadastro no banco.
//  2. Regra encontrada no repositório via GetIPIRegra.
//  3. Se o repositório não retornar regra mas o item fornecer alíquota ou
//     valor de pauta nos detalhes, o IPI é calculado inteiramente a partir
//     dos detalhes (modo "inline").
type IPICalculator struct {
	repo repository.TaxRepository
}

func NewIPICalculator(r repository.TaxRepository) *IPICalculator {
	return &IPICalculator{repo: r}
}

// Calculate implementa TaxCalculator diretamente (sem necessidade de LegacyAdapter).
// Retorna um []ItemDocumentoFiscalSaida indexado na mesma ordem de input.Itens.
func (c *IPICalculator) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
	// Garante que os mapas de detalhes estejam inicializados
	docMap := input.ToMap()

	// Valores de rateio no nível do documento (frete, seguro, etc.)
	frete := models.GetDecimal(docMap, models.KeyDocumentoInfosValorFrete)
	seguro := models.GetDecimal(docMap, models.KeyDocumentoInfosValorSeguro)
	outrasDespesas := models.GetDecimal(docMap, models.KeyDocumentoInfosValorOutrasDespesas)
	desconto := models.GetDecimal(docMap, models.KeyDocumentoInfosValorDesconto)

	// Step 1 — Total de produtos (denominador do fator de rateio)
	totalProdutos := decimal.Zero
	for _, item := range input.Itens {
		totalProdutos = totalProdutos.Add(item.ValorUnitario.Mul(item.Quantidade))
	}

	resultado := make([]models.ItemDocumentoFiscalSaida, len(input.Itens))

	for i, item := range input.Itens {
		resultado[i].SKU = item.SKU

		itemMap := item.ToMap()

		// --- Parâmetros de busca ---
		// TipoOperacaoFiscal: item pode sobrescrever o valor do documento
		tipoOperacaoFiscal := string(models.NormalizeTipoOperacao(input.TipoOperacaoFiscal))
		if override := string(models.NormalizeTipoOperacao(models.GetString(itemMap, models.KeyDocumentoInfosItemSubstituirTipoOperacaoFiscal))); override != "" {
			tipoOperacaoFiscal = override
		}

		ncm := strings.ToUpper(strings.TrimSpace(item.NCM))
		exIPI := strings.ToUpper(strings.TrimSpace(models.GetString(itemMap, models.KeyDocumentoInfosItemSubstituirEX_IPI)))
		crtEmitente := string(models.NormalizeCRTEmitente(input.CRTEmitente))
		perfilComprador := strings.ToUpper(strings.TrimSpace(models.GetString(docMap, models.KeyDocumentoInfosDestinatarioPerfilComprador)))
		ufDestino := strings.ToUpper(strings.TrimSpace(input.LocalizacaoDestino.UF))
		zonaEspecialStr := models.GetString(docMap, models.KeyDocumentoInfosZonaEspecial)
		zonaEspecial := models.IsZonaEspecial(zonaEspecialStr)
		dataOperacao := input.DataOperacao.Format("2006-01-02")

		// --- Overrides de regra provenientes dos Detalhes do item ---
		overrideAliquota := models.GetDecimal(itemMap, models.KeyDocumentoInfosItemIPIAliquota)
		overrideValorPauta := models.GetDecimal(itemMap, models.KeyDocumentoInfosItemIPIValorPauta)
		overrideCST := strings.TrimSpace(models.GetString(itemMap, models.KeyDocumentoInfosItemIPICST))
		overrideCEnq := strings.TrimSpace(models.GetString(itemMap, models.KeyDocumentoInfosItemIPICEnq))

		hasDetalheOverride := overrideAliquota.IsPositive() || overrideValorPauta.IsPositive()
		hasCompleteDetalheConfig := hasDetalheOverride && overrideCST != "" && overrideCEnq != ""

		slog.Debug("Calculando IPI",
			"SKU", item.SKU,
			"NCM", ncm,
			"ExIPI", exIPI,
			"CRT", crtEmitente,
			"TipoOp", tipoOperacaoFiscal,
			"Perfil", perfilComprador,
			"UF", ufDestino,
			"ZonaEspecial", zonaEspecial,
			"DataOp", dataOperacao,
			"override_aliquota", overrideAliquota,
			"override_valor_pauta", overrideValorPauta,
			"detalhe_config_completa", hasCompleteDetalheConfig,
		)

		regra := &repository.IPIRegra{}
		consultouRepositorio := false
		var err error

		if hasCompleteDetalheConfig {
			// Todas as informações necessárias vieram dos parâmetros ITEM_IPI_*.
			// Nesse cenário, evita a ida ao banco para reduzir latência.
			slog.Debug("IPI: pulando consulta no banco, usando somente parâmetros ITEM_IPI_*", "SKU", item.SKU)
		} else {
			// --- Busca a regra no repositório ---
			consultouRepositorio = true
			regra, err = c.repo.GetIPIRegra(
				ctx, ncm, exIPI, crtEmitente, tipoOperacaoFiscal,
				perfilComprador, ufDestino, zonaEspecial, dataOperacao,
			)
			if err != nil {
				if !hasDetalheOverride {
					// Sem regra no banco e sem override nos detalhes: item sem IPI
					slog.Warn("IPI: regra não encontrada e sem override, item sem IPI", "SKU", item.SKU, "NCM", ncm, "err", err)
					continue
				}
				// Sem regra no banco, mas o item forneceu alíquota/pauta inline
				slog.Debug("IPI: regra não encontrada no banco, usando override dos detalhes", "SKU", item.SKU)
				regra = &repository.IPIRegra{}
			}
		}

		// --- Aplica overrides dos Detalhes do item sobre a regra ---
		// Os Detalhes têm precedência sobre o que veio do banco.
		if overrideAliquota.IsPositive() {
			regra.AliquotaIPI = overrideAliquota
		}
		if overrideValorPauta.IsPositive() {
			regra.ValorPautaIPI = overrideValorPauta
		}
		if overrideCST != "" {
			regra.CSTIPI = overrideCST
		}
		if overrideCEnq != "" {
			regra.CEnq = overrideCEnq
		}

		// --- Fator de rateio ---
		valorItem := item.ValorUnitario.Mul(item.Quantidade)
		var fatorRateio decimal.Decimal
		if totalProdutos.IsZero() {
			fatorRateio = decimal.Zero
		} else {
			fatorRateio = valorItem.Div(totalProdutos)
		}

		// --- Base de Cálculo com rateio das despesas acessórias ---
		baseCalculo := valorItem.
			Add(frete.Mul(fatorRateio)).
			Add(seguro.Mul(fatorRateio)).
			Add(outrasDespesas.Mul(fatorRateio)).
			Sub(desconto.Mul(fatorRateio))

		// --- Valor do IPI ---
		var valorIPI decimal.Decimal
		var metodoCalculo string

		if regra.ValorPautaIPI.IsPositive() {
			// Ad Pauta (valor fixo por unidade, ex: bebidas, cigarros)
			valorIPI = item.Quantidade.Mul(regra.ValorPautaIPI)
			metodoCalculo = string(models.MetodoCalculoAdPauta)
		} else {
			// Ad Valorem (percentual sobre base de cálculo)
			valorIPI = baseCalculo.Mul(regra.AliquotaIPI.Div(decimal.NewFromInt(100)))
			metodoCalculo = string(models.MetodoCalculoAdValorem)
		}

		valorIPI = valorIPI.Round(2)

		fonte := "repositorio"
		if hasCompleteDetalheConfig {
			fonte = "detalhe_item"
		} else if hasDetalheOverride && err != nil {
			fonte = "detalhe_item"
		} else if consultouRepositorio && hasDetalheOverride {
			fonte = "repositorio+override"
		}

		tributo := models.TributosItemDocumentoFiscalSaida{
			Tributo:     "IPI",
			CST:         regra.CSTIPI,
			BaseCalculo: baseCalculo.Round(2),
			Aliquota:    regra.AliquotaIPI,
			Valor:       valorIPI,
			MoreNumericDetails: []models.Detalhe{
				{Key: "valor_item", Value: valorItem},
				{Key: "fator_rateio", Value: fatorRateio},
				{Key: "frete_rateado", Value: frete.Mul(fatorRateio).Round(2)},
				{Key: "seguro_rateado", Value: seguro.Mul(fatorRateio).Round(2)},
				{Key: "outras_despesas_rateadas", Value: outrasDespesas.Mul(fatorRateio).Round(2)},
				{Key: "desconto_rateado", Value: desconto.Mul(fatorRateio).Round(2)},
				{Key: "base_calculo", Value: baseCalculo.Round(2)},
				{Key: "valor_ipi", Value: valorIPI},
			},
			MoreTextDetails: []models.Detalhe{
				{Key: "regra_id", Value: regra.ID},
				{Key: "c_enq", Value: regra.CEnq},
				{Key: "metodo_calculo", Value: metodoCalculo},
				{Key: "fonte_regra", Value: fonte},
				{Key: "ncm", Value: ncm},
				{Key: "ex_ipi", Value: exIPI},
				{Key: "crt_emitente", Value: crtEmitente},
				{Key: "tipo_operacao_fiscal", Value: tipoOperacaoFiscal},
				{Key: "perfil_comprador", Value: perfilComprador},
				{Key: "uf_destino", Value: ufDestino},
				{Key: "zona_especial", Value: zonaEspecial},
			},
		}

		resultado[i].Tributos = append(resultado[i].Tributos, tributo)
	}

	return resultado, nil
}
