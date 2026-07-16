package models

// constants_v2.go no formato simplificado solicitado:
// um unico type (KeyDocumentoInfos) e um unico bloco const com todas as keys.
type KeyDocumentoInfos string

// Enums de domínio para campos recorrentes de entrada/cálculo.
type CRTEmitente string

const (
	CRTEmitenteSimples        CRTEmitente = "SIMPLES"
	CRTEmitenteLucroReal      CRTEmitente = "LUCRO_REAL"
	CRTEmitenteLucroPresumido CRTEmitente = "LUCRO_PRESUMIDO"
	CRTEmitenteMEI            CRTEmitente = "MEI"
)

type TipoOperacaoFiscal string

const (
	TipoOperacaoFiscalSaida         TipoOperacaoFiscal = "SAIDA"
	TipoOperacaoFiscalEntrada       TipoOperacaoFiscal = "ENTRADA"
	TipoOperacaoFiscalTransferencia TipoOperacaoFiscal = "TRANSFERENCIA"
)

type IndicadorPresenca string

const (
	IndicadorPresencaCodigo0 IndicadorPresenca = "0"
	IndicadorPresencaCodigo1 IndicadorPresenca = "1"
	IndicadorPresencaCodigo2 IndicadorPresenca = "2"
	IndicadorPresencaCodigo3 IndicadorPresenca = "3"
	IndicadorPresencaCodigo4 IndicadorPresenca = "4"
	IndicadorPresencaCodigo5 IndicadorPresenca = "5"
	IndicadorPresencaCodigo6 IndicadorPresenca = "6"
	IndicadorPresencaCodigo7 IndicadorPresenca = "7"
	IndicadorPresencaCodigo8 IndicadorPresenca = "8"
	IndicadorPresencaCodigo9 IndicadorPresenca = "9"
)

type ZonaEspecial string

const (
	ZonaEspecialNenhuma ZonaEspecial = ""
	ZonaEspecialZFM     ZonaEspecial = "ZFM"
	ZonaEspecialALC     ZonaEspecial = "ALC"
)

type MetodoCalculoTributo string

const (
	MetodoCalculoAdValorem MetodoCalculoTributo = "AD_VALOREM"
	MetodoCalculoAdPauta   MetodoCalculoTributo = "AD_PAUTA"
)

type NaturezaOperacao string

const (
	NaturezaOperacaoRevenda          NaturezaOperacao = "REVENDA"
	NaturezaOperacaoConsumo          NaturezaOperacao = "CONSUMO"
	NaturezaOperacaoAtivoImobiliario NaturezaOperacao = "ATIVO_IMOBILIARIO"
	NaturezaOperacaoServico          NaturezaOperacao = "SERVICO"
	NaturezaOperacaoOutros           NaturezaOperacao = "OUTROS"
)

type CSTICMS string

const (
	CSTICMS010 CSTICMS = "010"
	CSTICMS000 CSTICMS = "000"
)

type CSTPISCOFINS string

const (
	CSTPISCOFINS01 CSTPISCOFINS = "01"
	CSTPISCOFINS02 CSTPISCOFINS = "02"
	CSTPISCOFINS03 CSTPISCOFINS = "03"
	CSTPISCOFINS04 CSTPISCOFINS = "04"
	CSTPISCOFINS05 CSTPISCOFINS = "05"
	CSTPISCOFINS06 CSTPISCOFINS = "06"
	CSTPISCOFINS49 CSTPISCOFINS = "49"
	CSTPISCOFINS99 CSTPISCOFINS = "99"
)

const (
	// Documento: componentes de calculo
	KeyDocumentoInfosValorFrete          KeyDocumentoInfos = "VALOR_FRETE"
	KeyDocumentoInfosValorSeguro         KeyDocumentoInfos = "VALOR_SEGURO"
	KeyDocumentoInfosValorOutrasDespesas KeyDocumentoInfos = "VALOR_OUTRAS_DESPESAS"
	KeyDocumentoInfosValorDesconto       KeyDocumentoInfos = "VALOR_DESCONTO"
	KeyDocumentoInfosValorExclusaoICMS   KeyDocumentoInfos = "VALOR_EXCLUSAO_ICMS"
	KeyDocumentoInfosValorExclusaoISS    KeyDocumentoInfos = "VALOR_EXCLUSAO_ISS"

	// Documento: contexto/substituicoes
	KeyDocumentoInfosRBT12                        KeyDocumentoInfos = "RBT12"
	KeyDocumentoInfosAnexoSimples                 KeyDocumentoInfos = "ANEXO_SIMPLES"
	KeyDocumentoInfosIndInscrEstadualDestinataria KeyDocumentoInfos = "IND_INSCR_ESTADUAL_DESTINATARIA"
	KeyDocumentoInfosContribuinteICMS             KeyDocumentoInfos = "CONTRIBUINTE_ICMS"
	KeyDocumentoInfosTipoOperacaoFiscal           KeyDocumentoInfos = "TIPO_OPERACAO_FISCAL"
	KeyDocumentoInfosFinalidadeOperacao           KeyDocumentoInfos = "FINALIDADE_OPERACAO"
	KeyDocumentoInfosDestinoFinal                 KeyDocumentoInfos = "DESTINO_FINAL"
	KeyDocumentoInfosIndicadorPresenca            KeyDocumentoInfos = "INDICADOR_PRESENCA"
	KeyDocumentoInfosProtocoloICMS                KeyDocumentoInfos = "PROTOCOLO_ICMS"
	KeyDocumentoInfosModalidadeFrete              KeyDocumentoInfos = "MODALIDADE_FRETE"
	KeyDocumentoInfosZonaEspecial                 KeyDocumentoInfos = "ZONA_ESPECIAL"

	// Documento: contexto do destinatario
	KeyDocumentoInfosDestinatarioIndIE            KeyDocumentoInfos = "DESTINATARIO_IND_IE"
	KeyDocumentoInfosDestinatarioContribuinteICMS KeyDocumentoInfos = "DESTINATARIO_CONTRIBUINTE_ICMS"
	KeyDocumentoInfosDestinatarioCRT              KeyDocumentoInfos = "DESTINATARIO_CRT"
	KeyDocumentoInfosDestinatarioRegimeTributario KeyDocumentoInfos = "REGIME_TRIBUTARIO_DESTINATARIO"
	KeyDocumentoInfosDestinatarioPerfilComprador  KeyDocumentoInfos = "DESTINATARIO_PERFIL_COMPRADOR" // Perfil do Comprador (PCD, Isento, Contribuinte)

	// Informações do produtor/fabricante
	KeyDocumentoInfosProdutorCNPJ KeyDocumentoInfos = "PRODUTOR_CNPJ"
	KeyDocumentoInfosProdutorUF   KeyDocumentoInfos = "PRODUTOR_UF"

	// Item: ICMS — parâmetros equivalentes aos configuráveis em tabela
	// (alíquotas, CST/CSOSN, ST/FCP, desoneração, regime/tipo de operação e metadados fiscais)
	KeyDocumentoInfosItemSubstituirAliquotaICMS                KeyDocumentoInfos = "ITEM_SUBSTITUIR_ALIQUOTA_ICMS"
	KeyDocumentoInfosItemCSOSN                                 KeyDocumentoInfos = "CSOSN"
	KeyDocumentoInfosItemSubstituirCSTICMS                     KeyDocumentoInfos = "ITEM_SUBSTITUIR_CST_ICMS"
	KeyDocumentoInfosItemSubstituirCSTEquivalente              KeyDocumentoInfos = "ITEM_SUBSTITUIR_CST_EQUIVALENTE"
	KeyDocumentoInfosItemSubstituirPermiteCredito              KeyDocumentoInfos = "ITEM_SUBSTITUIR_PERMITE_CREDITO"
	KeyDocumentoInfosItemSubstituirCEST                        KeyDocumentoInfos = "ITEM_SUBSTITUIR_CEST"
	KeyDocumentoInfosItemSubstituirCFOP                        KeyDocumentoInfos = "ITEM_SUBSTITUIR_CFOP"
	KeyDocumentoInfosItemSubstituirAliquotaInternaDestino      KeyDocumentoInfos = "ITEM_SUBSTITUIR_ALIQUOTA_INTERNA_DESTINO"
	KeyDocumentoInfosItemSubstituirAliquotaInterestadual       KeyDocumentoInfos = "ITEM_SUBSTITUIR_ALIQUOTA_INTERESTADUAL"
	KeyDocumentoInfosItemSubstituirReducaoBase                 KeyDocumentoInfos = "ITEM_SUBSTITUIR_REDUCAO_BASE"
	KeyDocumentoInfosItemSubstituirMVAPercentual               KeyDocumentoInfos = "ITEM_SUBSTITUIR_MVA_PERCENTUAL"
	KeyDocumentoInfosItemSubstituirMVAAjustado                 KeyDocumentoInfos = "ITEM_SUBSTITUIR_MVA_AJUSTADO"
	KeyDocumentoInfosItemSubstituirModalidadeBCST              KeyDocumentoInfos = "ITEM_SUBSTITUIR_MODALIDADE_BC_ST"
	KeyDocumentoInfosItemSubstituirPossuiProtocoloST           KeyDocumentoInfos = "ITEM_SUBSTITUIR_POSSUI_PROTOCOLO_ST"
	KeyDocumentoInfosItemSubstituirPercentualFCP               KeyDocumentoInfos = "ITEM_SUBSTITUIR_PERCENTUAL_FCP"
	KeyDocumentoInfosItemSubstituirPercentualFCPST             KeyDocumentoInfos = "ITEM_SUBSTITUIR_PERCENTUAL_FCP_ST"
	KeyDocumentoInfosItemSubstituirDiferimento                 KeyDocumentoInfos = "ITEM_SUBSTITUIR_DIFERIMENTO"
	KeyDocumentoInfosItemSubstituirRegimeTributarioDestino     KeyDocumentoInfos = "ITEM_SUBSTITUIR_REGIME_TRIBUTARIO_DESTINO"
	KeyDocumentoInfosItemSubstituirTipoOperacaoFiscal          KeyDocumentoInfos = "ITEM_SUBSTITUIR_TIPO_OPERACAO_FISCAL"
	KeyDocumentoInfosItemSubstituirDesoneracaoNaoAplicar       KeyDocumentoInfos = "ITEM_SUBSTITUIR_DESONERACAO_NAO_APLICAR"
	KeyDocumentoInfosItemSubstituirDesoneracaoSimAplicar       KeyDocumentoInfos = "ITEM_SUBSTITUIR_DESONERACAO_SIM_APLICAR"
	KeyDocumentoInfosItemSubstituirDesoneracaoAliquota         KeyDocumentoInfos = "ITEM_SUBSTITUIR_DESONERACAO_ALIQUOTA"
	KeyDocumentoInfosItemSubstituirDesoneracaoLimitarAliquota  KeyDocumentoInfos = "ITEM_SUBSTITUIR_DESONERACAO_LIMITAR_ALIQUOTA"
	KeyDocumentoInfosItemSubstituirDesoneracaoValorJaCalculado KeyDocumentoInfos = "ITEM_SUBSTITUIR_DESONERACAO_VALOR_JA_CALCULADO"
	KeyDocumentoInfosItemDesoneracaoAplicar                    KeyDocumentoInfos = "ITEM_DESONERACAO_APLICAR"
	KeyDocumentoInfosItemDesoneracaoMotivoICMS                 KeyDocumentoInfos = "ITEM_DESONERACAO_MOTIVO_ICMS"
	KeyDocumentoInfosItemDesoneracaoPercentual                 KeyDocumentoInfos = "ITEM_DESONERACAO_PERCENTUAL"
	KeyDocumentoInfosItemDesoneracaoCargaEfetivaAlvo           KeyDocumentoInfos = "ITEM_DESONERACAO_CARGA_EFETIVA_ALVO"
	KeyDocumentoInfosItemDesoneracaoValorICMS                  KeyDocumentoInfos = "ITEM_DESONERACAO_VALOR_ICMS"
	KeyDocumentoInfosItemSTModalidadeBC                        KeyDocumentoInfos = "ITEM_ST_MODALIDADE_BC"
	KeyDocumentoInfosItemSTPossuiProtocolo                     KeyDocumentoInfos = "ITEM_ST_POSSUI_PROTOCOLO"
	KeyDocumentoInfosItemSTMVAAjustado                         KeyDocumentoInfos = "ITEM_ST_MVA_AJUSTADO"
	KeyDocumentoInfosItemSTPercentualFCP                       KeyDocumentoInfos = "ITEM_ST_PERCENTUAL_FCP"
	KeyDocumentoInfosItemSTPercentualFCPST                     KeyDocumentoInfos = "ITEM_ST_PERCENTUAL_FCP_ST"
	KeyDocumentoInfosItemDetalheCodBeneficioFiscal             KeyDocumentoInfos = "ITEM_DETALHE_COD_BENEFICIO_FISCAL"
	KeyDocumentoInfosItemDetalheNumeroProcessoJudicial         KeyDocumentoInfos = "ITEM_DETALHE_NUMERO_PROCESSO_JUDICIAL"
	KeyDocumentoInfosItemDetalheMotivoDesoneracaoICMS          KeyDocumentoInfos = "ITEM_DETALHE_MOTIVO_DESONERACAO_ICMS"
	KeyDocumentoInfosItemDetalheValorICMSDesonerado            KeyDocumentoInfos = "ITEM_DETALHE_VALOR_ICMS_DESONERADO"
	KeyDocumentoInfosItemDetalheRegraAplicadaOrigem            KeyDocumentoInfos = "ITEM_DETALHE_REGRA_APLICADA_ORIGEM"
	KeyDocumentoInfosItemDetalheRegraAplicadaDestino           KeyDocumentoInfos = "ITEM_DETALHE_REGRA_APLICADA_DESTINO"
	KeyDocumentoInfosItemDetalheRegraFonte                     KeyDocumentoInfos = "ITEM_DETALHE_REGRA_FONTE"

	// Item: PIS — parâmetros e overrides
	KeyDocumentoInfosItemSubstituirCSTPIS                   KeyDocumentoInfos = "ITEM_SUBSTITUIR_CST_PIS"
	KeyDocumentoInfosItemSubstituirAliquotaPIS              KeyDocumentoInfos = "ITEM_SUBSTITUIR_ALIQUOTA_PIS"
	KeyDocumentoInfosItemPISCOFINSCSTPIS                    KeyDocumentoInfos = "ITEM_PISCOFINS_CST_PIS"
	KeyDocumentoInfosItemPISCOFINSAliquotaPIS               KeyDocumentoInfos = "ITEM_PISCOFINS_ALIQUOTA_PIS"
	KeyDocumentoInfosItemSubstituirExcluirICMSBasePISCOFINS KeyDocumentoInfos = "ITEM_SUBSTITUIR_EXCLUIR_ICMS_BASE_PIS_COFINS"
	KeyDocumentoInfosItemPISCOFINSExcluirICMSBase           KeyDocumentoInfos = "ITEM_PISCOFINS_EXCLUIR_ICMS_BASE"

	// Item: COFINS — parâmetros e overrides
	KeyDocumentoInfosItemSubstituirCSTCOFINS      KeyDocumentoInfos = "ITEM_SUBSTITUIR_CST_COFINS"
	KeyDocumentoInfosItemSubstituirAliquotaCOFINS KeyDocumentoInfos = "ITEM_SUBSTITUIR_ALIQUOTA_COFINS"
	KeyDocumentoInfosItemPISCOFINSCSTCOFINS       KeyDocumentoInfos = "ITEM_PISCOFINS_CST_COFINS"
	KeyDocumentoInfosItemPISCOFINSAliquotaCOFINS  KeyDocumentoInfos = "ITEM_PISCOFINS_ALIQUOTA_COFINS"

	// Item: overrides de outros tributos/subsistemas
	KeyDocumentoInfosItemSubstituirEX_IPI KeyDocumentoInfos = "ITEM_SUBSTITUIR_EX_IPI" /// O EX-IPI (Exceção do IPI) é um código numérico (geralmente 01, 02...) que indica uma exceção à alíquota padrão de um NCM específico na Tabela de Incidência do IPI (TIPI)

	// Item: IPI — overrides/substituicoes
	KeyDocumentoInfosItemIPIAliquota   KeyDocumentoInfos = "ITEM_IPI_ALIQUOTA"    // override da alíquota ad valorem (%)
	KeyDocumentoInfosItemIPIValorPauta KeyDocumentoInfos = "ITEM_IPI_VALOR_PAUTA" // override do valor de pauta por unidade (ad pauta)
	KeyDocumentoInfosItemIPICST        KeyDocumentoInfos = "ITEM_IPI_CST"         // override do CST-IPI (ex: "50", "53", "99")
	KeyDocumentoInfosItemIPICEnq       KeyDocumentoInfos = "ITEM_IPI_C_ENQ"       // override do código de enquadramento legal
)
