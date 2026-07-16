export type MovimentacaoMensalType = {
  id: number; // Chave primária
  dataCadastro: string; // Data de cadastro no formato ISO
  dataAtualizacao: string; // Data de atualização no formato ISO
  associadoId: number; // ID do associado relacionado
  mes: number; // Mês da movimentação (1-12)
  ano: number; // Ano da movimentação
  valorTotal: number; // Valor total da movimentação
  dataVencimento: string; // Data de vencimento no formato ISO
  dataPagamento: string | null; // Data de pagamento no formato ISO ou null
  statusPagamento: string; // Status do pagamento ("Pago" ou "Pendente")
  observacoes: string; // Observações sobre a movimentação
  energiaRecebidaKwh: number; // Energia recebida em kWh
  valorEnergiaRecebida: number; // Valor da energia recebida
  tarifaUnitariaKwh: number; // Tarifa unitária por kWh
  valorCobrado: number; // Valor cobrado na movimentação
  valorEconomizado: number; // Valor economizado na movimentação
  percentualEconomizado: number; // Percentual economizado na movimentação
};
