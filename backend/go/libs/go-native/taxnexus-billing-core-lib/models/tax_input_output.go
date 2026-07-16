// path: backend/go/libs/go-native/taxnexus-billing-core-lib/models/tax_input_output.go
package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// Estrutura genérica para pares chave-valor
type Detalhe struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type DetalhesNormalizados struct {
	Todos    map[string]interface{}     `json:"-"`
	Decimais map[string]decimal.Decimal `json:"-"`
}

type DocumentoFiscalEntrada struct {
	CorrelacaoID       string    `json:"correlacao_id" validate:"required"`                             // ID para correlação de logs e rastreamento
	DocumentoID        string    `json:"documento_id" validate:"required"`                              // ID único para rastreamento
	DataOperacao       time.Time `json:"data_operacao" validate:"required"`                             // Data da operação (pode ser diferente da data de emissão)
	TipoOperacaoFiscal string    `json:"tipo_operacao_fiscal" validate:"required,tipo_operacao_fiscal"` // SAIDA, ENTRADA, TRANSFERENCIA, etc.
	// antigo: FinalidadeNotaFiscal
	NaturezaOperacao        string                       `json:"natureza_operacao" validate:"omitempty,natureza_operacao"`   // REVENDA, CONSUMO, ATIVO_IMOBILIARIO, SERVICO, OUTROS
	CFOP                    string                       `json:"cfop"`                                                       // CFOP significa Código Fiscal de Operações e Prestações, usado para identificar a natureza da circulação da mercadoria ou a prestação de serviço de transporte.Ele diz para o fisco o que está acontecendo na operação
	CRTEmitente             string                       `json:"crt_emitente" validate:"required,crt_emitente"`              // CRT oficial: 1=Simples Nacional, 2=Simples excesso sublimite, 3=Regime Normal, 4=MEI
	LocalizacaoOrigem       LocalizacaoFiscal            `json:"localizacao_origem" validate:"required"`                     // UF + Municipio
	LocalizacaoDestino      LocalizacaoFiscal            `json:"localizacao_destino" validate:"required"`                    // UF + Municipio
	IsDestinoFinal          bool                         `json:"is_destino_final"`                                           // Consumidor Final ou Não
	IndicadorPresenca       string                       `json:"indicador_presenca" validate:"omitempty,indicador_presenca"` // 0-9 conforme tabela do SPED
	DetalhesDocumentoFiscal []Detalhe                    `json:"detalhes_documento_fiscal"`                                  // Lista de detalhes específicos do item, como CFOP, CEST, etc., para flexibilidade máxima
	Itens                   []ItemDocumentoFiscalEntrada `json:"itens" validate:"required,min=1,dive"`                       // Lista de itens da nota fiscal
	MapaDetalhes            DetalhesNormalizados         `json:"-"`
}

type LocalizacaoFiscal struct {
	UF        string `json:"uf" validate:"required,len=2"` // Unidade Federativa (Estado)
	Municipio string `json:"municipio_codigo_ibge"`        // Código IBGE do município
}

type ItemDocumentoFiscalEntrada struct {
	SKU                         string               `json:"sku" validate:"required"`                                          // SKU do produto, usado para buscar regras específicas
	NCM                         string               `json:"ncm" validate:"required"`                                          // NCM completo, usado para buscar regras específicas
	Quantidade                  decimal.Decimal      `json:"quantidade" validate:"dec_gt_zero"`                                // Quantidade do item, usada para cálculos de impostos que dependem da quantidade
	ValorUnitario               decimal.Decimal      `json:"valor_unitario" validate:"dec_gt_zero"`                            // Valor unitário do item, usado para cálculos de impostos ad valorem
	NaturezaItemNotaFiscal      string               `json:"natureza_item_nota_fiscal" validate:"omitempty,natureza_operacao"` // REVENDA, CONSUMO, ATIVO_IMOBILIARIO, SERVICO, OUTROS
	IsDestinoFinal              bool                 `json:"is_destino_final"`                                                 // Consumidor Final ou Não, pode ser redundante com o campo do documento, mas é enviado para facilitar
	DetalhesItemDocumentoFiscal []Detalhe            `json:"detalhes_item_documento_fiscal" validate:"dive"`                   // Lista de detalhes específicos do item, como CFOP, CEST, etc., para flexibilidade máxima
	MapaDetalhes                DetalhesNormalizados `json:"-"`
}

// --- Response Structs ---

type DocumentoFiscalSaida struct {
	IDTransaction string                     `json:"id_transaction"`
	CorrelacaoID  string                     `json:"correlacao_id"`
	DocumentoID   string                     `json:"documento_id"`
	TotalNota     decimal.Decimal            `json:"total_nota"`
	TotalImpostos decimal.Decimal            `json:"total_impostos"`
	SplitPayment  *SplitPayment              `json:"split_payment,omitempty"`
	Itens         []ItemDocumentoFiscalSaida `json:"itens"`
}

// SplitPayment representa a partição financeira esperada no momento da liquidação
// bancária (BR-09 — Split Payment da Reforma Tributária).
//
// A rede bancária usa estes valores para segregar o pagamento:
//   - ValorReceitaLiquida → conta da empresa (receita livre)
//   - ValorCBSReter        → conta vinculada CBS (União)
//   - ValorIBSReter        → conta vinculada IBS (Comitê Gestor)
//   - ValorISReter         → conta vinculada IS (União)
//
// O CodigoBarrasSplit (SHA-256) permite conciliação bancária determinística.
type SplitPayment struct {
	ValorReceitaLiquida decimal.Decimal `json:"valor_receita_liquida"`
	ValorCBSReter       decimal.Decimal `json:"valor_cbs_reter"`
	ValorIBSReter       decimal.Decimal `json:"valor_ibs_reter"`
	ValorISReter        decimal.Decimal `json:"valor_is_reter"`
	CodigoBarrasSplit   string          `json:"codigo_barras_split"`
}

type ItemDocumentoFiscalSaida struct {
	SKU          string                             `json:"sku"`
	Total        decimal.Decimal                    `json:"total"`
	ValorLiquido decimal.Decimal                    `json:"valor_liquido"`
	Tributos     []TributosItemDocumentoFiscalSaida `json:"tributos"`
}

type TributosItemDocumentoFiscalSaida struct {
	Tributo            string          `json:"tributo"`
	CST                string          `json:"cst,omitempty"`
	CSOSN              string          `json:"csosn,omitempty"`
	CEST               string          `json:"cest,omitempty"`
	BaseCalculo        decimal.Decimal `json:"base_calculo"`
	Aliquota           decimal.Decimal `json:"aliquota"`
	Valor              decimal.Decimal `json:"valor"`
	MoreNumericDetails []Detalhe       `json:"more_numeric_details,omitempty"`
	MoreTextDetails    []Detalhe       `json:"more_text_details,omitempty"`
}

/*
Exemplo de payload de entrada (DocumentoFiscalEntrada):

{
	"correlacao_id": "corr-20260402-0001",
	"documento_id": "nf-123456",
	"data_operacao": "2026-04-02T10:30:00Z",
	"tipo_operacao_fiscal": "SAIDA",
	"finalidade_nota_fiscal": "REVENDA",
	"crt_emitente": "1",
	"localizacao_origem": {
		"uf": "SP",
		"municipio_codigo_ibge": "3550308"
	},
	"localizacao_destino": {
		"uf": "RJ",
		"municipio_codigo_ibge": "3304557"
	},
	"is_destino_final": true,
	"indicador_presenca": "2",
	"detalhes_documento_fiscal": [
		{ "key": "CRT", "value": "1" },
		{ "key": "RBT12", "value": "1250000.55" },
		{ "key": "ANEXO_SIMPLES", "value": "I" },
		{ "key": "DESTINATARIO_IND_IE", "value": "9" },
		{ "key": "DESTINATARIO_CONTRIBUINTE_ICMS", "value": false },
		{ "key": "MODALIDADE_FRETE", "value": "CIF" },
		{ "key": "FRETE", "value": "45.90" },
		{ "key": "SEGURO", "value": "0" },
		{ "key": "OUTRAS_DESPESAS", "value": "12.00" },
		{ "key": "DESCONTO", "value": "5.00" },
		{ "key": "PROTOCOLO_ICMS", "value": "ICMS-45/2026" }
	],
	"itens": [
		{
			"sku": "SKU-001",
			"ncm": "22030000",
			"quantidade": "10",
			"valor_unitario": "12.50",
			"natureza_item_nota_fiscal": "REVENDA",
			"is_destino_final": true,
			"detalhes_item_documento_fiscal": [
				{ "key": "CSOSN", "value": "101" },
				{ "key": "ITEM_SUBSTITUIR_CEST", "value": "0300100" },
				{ "key": "ITEM_SUBSTITUIR_CFOP", "value": "6101" },
				{ "key": "ITEM_SUBSTITUIR_EX_IPI", "value": "01" },
				{ "key": "ITEM_SUBSTITUIR_ALIQUOTA_ICMS", "value": "18.00" },
				{ "key": "ITEM_SUBSTITUIR_ALIQUOTA_INTERESTADUAL", "value": "12.00" },
				{ "key": "ITEM_SUBSTITUIR_ALIQUOTA_INTERNA_DESTINO", "value": "20.00" },
				{ "key": "ITEM_SUBSTITUIR_MVA_PERCENTUAL", "value": "35.00" },
				{ "key": "ITEM_PISCOFINS_CST_PIS", "value": "01" },
				{ "key": "ITEM_PISCOFINS_CST_COFINS", "value": "01" },
				{ "key": "ITEM_PISCOFINS_ALIQUOTA_PIS", "value": "1.65" },
				{ "key": "ITEM_PISCOFINS_ALIQUOTA_COFINS", "value": "7.60" },
				{ "key": "ITEM_PISCOFINS_EXCLUIR_ICMS_BASE", "value": true }
			]
		},
		{
			"sku": "SKU-002",
			"ncm": "30049099",
			"quantidade": "2",
			"valor_unitario": "87.30",
			"natureza_item_nota_fiscal": "CONSUMO",
			"is_destino_final": true,
			"detalhes_item_documento_fiscal": [
				{ "key": "CST_ICMS", "value": "040" },
				{ "key": "ITEM_DESONERACAO_APLICAR", "value": true },
				{ "key": "ITEM_DESONERACAO_MOTIVO_ICMS", "value": "3" },
				{ "key": "ITEM_DETALHE_COD_BENEFICIO_FISCAL", "value": "RJ123456" },
				{ "key": "ITEM_PISCOFINS_CST_PIS", "value": "06" },
				{ "key": "ITEM_PISCOFINS_CST_COFINS", "value": "06" }
			]
		}
	]
}

Exemplo de payload de saída (DocumentoFiscalSaida):

{
	"id_transaction": "tx-20260402-0001",
	"correlacao_id": "corr-20260402-0001",
	"documento_id": "nf-123456",
	"total_nota": "299.60",
	"total_impostos": "74.85",
	"itens": [
		{
			"sku": "SKU-001",
			"total": "125.00",
			"tributos": [
				{
					"tributo": "ICMS",
					"csosn": "101",
					"cest": "0300100",
					"base_calculo": "125.00",
					"aliquota": "18.00",
					"valor": "22.50",
					"more_numeric_details": [
						{ "key": "aliquota_interestadual", "value": "12.00" },
						{ "key": "aliquota_interna_destino", "value": "20.00" },
						{ "key": "mva_percentual", "value": "35.00" }
					],
					"more_text_details": [
						{ "key": "cfop", "value": "6101" },
						{ "key": "regra_fonte", "value": "item_override" }
					]
				},
				{
					"tributo": "PIS",
					"cst": "01",
					"base_calculo": "102.50",
					"aliquota": "1.65",
					"valor": "1.69",
					"more_numeric_details": [
						{ "key": "icms_excluido_base", "value": "22.50" }
					],
					"more_text_details": [
						{ "key": "metodo", "value": "base_com_exclusao_icms" }
					]
				},
				{
					"tributo": "COFINS",
					"cst": "01",
					"base_calculo": "102.50",
					"aliquota": "7.60",
					"valor": "7.79"
				}
			]
		},
		{
			"sku": "SKU-002",
			"total": "174.60",
			"tributos": [
				{
					"tributo": "ICMS",
					"cst": "040",
					"base_calculo": "174.60",
					"aliquota": "0",
					"valor": "0",
					"more_text_details": [
						{ "key": "motivo_desoneracao", "value": "3" },
						{ "key": "cod_beneficio_fiscal", "value": "RJ123456" }
					]
				},
				{
					"tributo": "PIS",
					"cst": "06",
					"base_calculo": "0",
					"aliquota": "0",
					"valor": "0"
				},
				{
					"tributo": "COFINS",
					"cst": "06",
					"base_calculo": "0",
					"aliquota": "0",
					"valor": "0"
				}
			]
		}
	]
}
*/
