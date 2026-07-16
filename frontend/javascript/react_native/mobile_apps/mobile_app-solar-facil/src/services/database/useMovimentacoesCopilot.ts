// /src/services/database/useMovimentacoesCopilot.ts
import { useDatabase } from "@/context/DatabaseContext";
import { MovimentacaoMensalType } from "@/types/MovimentacaoMensalType";
import { DBResponse } from "@/types/DBResponse";

export function useMovimentacoesCopilot() {
  const {
    getDatabaseConnection,
    isDatabaseConnected,
    initializeDatabaseConnection,
    dbInstance,
  } = useDatabase();

  console.log("=== INICIO ========================================================================");
  console.log("useMovimentacoesCopilot - Iniciando o hook useMovimentacoesCopilot...");
  console.info("useMovimentacoesCopilot - useDatabase() - Conexão com o banco de dados:", {
    isDatabaseConnected,
    dbInstance,
  });

  async function ensureDbConnected() {
    console.log("useAssociadosCopilot - ensureDbConnected - Verificando conexão com o banco...");
    if (!isDatabaseConnected) {
      await initializeDatabaseConnection();
    }
    const db = getDatabaseConnection();
    if (!db) throw new Error("useMovimentacoesCopilot - Banco de dados não conectado");
    return db;
  }

  async function insertRecord(mov: MovimentacaoMensalType): Promise<DBResponse<number>> {
    try {
      const db = await ensureDbConnected();
      const result = await db.runAsync(
        `INSERT INTO movimentacoes (
          associadoId, mes, ano, valorTotal, dataCadastro, dataAtualizacao,
          dataVencimento, dataPagamento, statusPagamento, observacoes,
          energiaRecebidaKwh, valorEnergiaRecebida, tarifaUnitariaKwh,
          valorCobrado, valorEconomizado, percentualEconomizado
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
        [
          mov.associadoId,
          mov.mes,
          mov.ano,
          mov.valorTotal,
          mov.dataCadastro,
          mov.dataAtualizacao,
          mov.dataVencimento,
          mov.dataPagamento,
          mov.statusPagamento,
          mov.observacoes,
          mov.energiaRecebidaKwh,
          mov.valorEnergiaRecebida,
          mov.tarifaUnitariaKwh,
          mov.valorCobrado,
          mov.valorEconomizado,
          mov.percentualEconomizado,
        ]
      );
      return { success: true, data: result?.lastInsertRowId ?? -1 };
    } catch (error) {
      console.error("useMovimentacoesCopilot - Erro ao inserir movimentação:", error);
      return { success: false, error: "Erro ao inserir movimentação" };
    }
  }

  async function updateRecord(mov: MovimentacaoMensalType): Promise<DBResponse<number>> {
    try {
      const db = await ensureDbConnected();
      const result = await db.runAsync(
        `UPDATE movimentacoes SET
          mes = ?, ano = ?, valorTotal = ?, dataCadastro = ?, dataAtualizacao = ?,
          dataVencimento = ?, dataPagamento = ?, statusPagamento = ?, observacoes = ?,
          energiaRecebidaKwh = ?, valorEnergiaRecebida = ?, tarifaUnitariaKwh = ?,
          valorCobrado = ?, valorEconomizado = ?, percentualEconomizado = ?
        WHERE id = ?;`,
        [
          mov.mes,
          mov.ano,
          mov.valorTotal,
          mov.dataCadastro,
          mov.dataAtualizacao,
          mov.dataVencimento,
          mov.dataPagamento,
          mov.statusPagamento,
          mov.observacoes,
          mov.energiaRecebidaKwh,
          mov.valorEnergiaRecebida,
          mov.tarifaUnitariaKwh,
          mov.valorCobrado,
          mov.valorEconomizado,
          mov.percentualEconomizado,
          mov.id,
        ]
      );
      return { success: true, data: result?.changes ?? 0 };
    } catch (error) {
      console.error("useMovimentacoesCopilot - Erro ao atualizar movimentação:", error);
      return { success: false, error: "Erro ao atualizar movimentação" };
    }
  }

  async function deleteRecord(id: number): Promise<DBResponse<number>> {
    try {
      const db = await ensureDbConnected();
      const result = await db.runAsync(`DELETE FROM movimentacoes WHERE id = ?;`, [id]);
      return { success: true, data: result?.changes ?? 0 };
    } catch (error) {
      console.error("useMovimentacoesCopilot - Erro ao deletar movimentação:", error);
      return { success: false, error: "Erro ao deletar movimentação" };
    }
  }

  async function searchByAssociadoId(associadoId: number): Promise<DBResponse<MovimentacaoMensalType[]>> {
    try {
      const db = await ensureDbConnected();
      const result = await db.getAllAsync(
        `SELECT * FROM movimentacoes WHERE associadoId = ? ORDER BY ano DESC, mes DESC;`,
        [associadoId]
      );
      const movimentacoes = (result ?? []) as MovimentacaoMensalType[];
      return { success: true, data: movimentacoes };
    } catch (error) {
      console.error("useMovimentacoesCopilot - Erro ao buscar movimentações por associado:", error);
      return { success: false, error: "Erro ao buscar movimentações por associado" };
    }
  }

  async function searchAll(): Promise<DBResponse<MovimentacaoMensalType[]>> {
    try {
      const db = await ensureDbConnected();
      const result = await db.getAllAsync(
        `SELECT * FROM movimentacoes ORDER BY ano DESC, mes DESC;`
      );
      const movimentacoes = (result ?? []) as MovimentacaoMensalType[];
      return { success: true, data: movimentacoes };
    } catch (error) {
      console.error("useMovimentacoesCopilot - Erro ao buscar todas movimentações:", error);
      return { success: false, error: "Erro ao buscar todas as movimentações" };
    }
  }

  return {
    insertRecord,
    updateRecord,
    deleteRecord,
    searchByAssociadoId,
    searchAll,
  };
}
