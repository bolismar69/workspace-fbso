// /src/sservices/storage/serviceMovimentacaoMensal.ts

import { adicionarItem, carregarLista, salvarListaV2 } from "@/services/storage/storageUtils";
import { MovimentacaoMensalType } from "@/types/MovimentacaoMensalType";

const FILENAME = "movimentacaomensal.json";

export async function salvarMovimentacaoMensal(data: MovimentacaoMensalType) {
  await adicionarItem(FILENAME, data);
}
export async function listarMovimentacoesMensais(): Promise<MovimentacaoMensalType[]> {
  return carregarLista<MovimentacaoMensalType>(FILENAME);
}
export async function atualizarMovimentacaoMensal(data: MovimentacaoMensalType): Promise<string> {
  const movimentacoes = await carregarLista<MovimentacaoMensalType>(FILENAME);
  const index = movimentacoes.findIndex((item) => item.id === data.id);
  if (index !== -1) {
    movimentacoes[index] = data;
    await salvarListaV2<MovimentacaoMensalType>(FILENAME, movimentacoes);
    return "Movimentação mensal atualizada com sucesso!";
  } else {
    return "Movimentação mensal nao encontrada!";
  }
}
  