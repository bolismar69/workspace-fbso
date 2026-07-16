// /src/services/database/useMovimentacoesCopilot.ts
import { useDatabase } from "@/context/DatabaseContext";
import { MovimentacaoMensalType } from "@/types/MovimentacaoMensalType";
import { DBResponse } from "@/types/DBResponse";

export function useMovimentacoesCopilotCopy() {
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
    if (!db) throw new Error("useAssociadosCopilot - ensureDbConnected - Banco de dados nao conectado");
    return db;
  }

  async function insertMovimentacao(mov: MovimentacaoMensalType): Promise<number> {
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
    return result?.lastInsertRowId ?? -1;
  }

  async function updateMovimentacao(mov: MovimentacaoMensalType): Promise<number> {
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
    return result?.changes ?? -1;
  }

  async function deleteMovimentacao(id: number): Promise<number> {
    const db = await ensureDbConnected();

    const result = await db.runAsync(`DELETE FROM movimentacoes WHERE id = ?;`, [id]);
    return result?.changes ?? -1;
  }

  async function searchByAssociado(associadoId: number): Promise<MovimentacaoMensalType[]> {
    const db = await ensureDbConnected();

    const result = await db.getAllAsync(
      `SELECT * FROM movimentacoes WHERE associadoId = ? ORDER BY ano DESC, mes DESC;`,
      [associadoId]
    );
    return result as MovimentacaoMensalType[];
  }
  async function searchByAssociadoId(associadoId: number): Promise<DBResponse<MovimentacaoMensalType[]>> {
    try {
      const db = await ensureDbConnected();
      const result = await db.getAllAsync(
        `SELECT * FROM movimentacoes WHERE associadoId = ? ORDER BY ano DESC, mes DESC;`,
        [associadoId]
      );

      if (!result || result.length === 0) {
        return { success: true, data: [] };
      }

      const movimentacoes = result as MovimentacaoMensalType[];
      return { success: true, data: movimentacoes };
    } catch (error) {
      console.error("useMovimentacoesCopilot - Erro ao buscar movimentações:", error);
      return { success: false, error: "Erro ao buscar movimentações" };
    }
  }

  async function searchAll(): Promise<MovimentacaoMensalType[]> {
    const db = await ensureDbConnected();

    const result = await db.getAllAsync(`SELECT * FROM movimentacoes ORDER BY ano DESC, mes DESC;`);
    return result as MovimentacaoMensalType[];
  }

  return {
    insertMovimentacao,
    updateMovimentacao,
    deleteMovimentacao,
    searchByAssociado,
    searchByAssociadoId,
    searchAll,
  };
}
