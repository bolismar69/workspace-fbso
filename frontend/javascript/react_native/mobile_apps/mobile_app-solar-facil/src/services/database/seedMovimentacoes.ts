// /src/services/database/seedMovimentacoes.ts
import { DBResponse } from "@/types/DBResponse";
import * as SQLite from "expo-sqlite";

export async function seedMovimentacoes(
  associadoId: number,
  db: SQLite.SQLiteDatabase
): Promise<DBResponse<{ count: number }>> {
  console.log("=== INICIO ========================================================================");
  console.log("seedMovimentacoes - Iniciando a inserção de registros de movimentação...", associadoId);

  if (!db) {
    console.error("seedMovimentacoes - Conexão com o banco de dados não fornecida.");
    return { success: false, error: "Conexão com o banco de dados não fornecida." };
  }

  const currentDate = new Date();
  let registrosInseridos = 0;

  try {
    for (let i = 0; i < 12; i++) {
      const movimentacaoDate = new Date(currentDate.getFullYear(), currentDate.getMonth() - i, 1);
      const mes = movimentacaoDate.getMonth() + 1;
      const ano = movimentacaoDate.getFullYear();
      const valorTotal = parseFloat((Math.random() * 1000).toFixed(2));
      const dataCadastro = movimentacaoDate.toISOString();
      const dataAtualizacao = dataCadastro;

      const dataVencimento = new Date(
        movimentacaoDate.getFullYear(),
        movimentacaoDate.getMonth(),
        Math.floor(Math.random() * 28) + 1
      ).toISOString();

      const dataPagamento = Math.random() > 0.5 ? dataVencimento : null;
      const statusPagamento = dataPagamento ? "Pago" : "Pendente";
      const observacoes = `Detalhe ${i + 1} da movimentação ${associadoId}`;

      const energiaRecebidaKwh = parseFloat((Math.random() * 100).toFixed(2));
      const valorEnergiaRecebida = parseFloat((energiaRecebidaKwh * 0.5).toFixed(2));
      const tarifaUnitariaKwh = 0.5;
      const valorCobrado = parseFloat((valorEnergiaRecebida * 0.85).toFixed(2));
      const valorEconomizado = parseFloat((valorEnergiaRecebida - valorCobrado).toFixed(2));
      const percentualEconomizado = parseFloat(((1 - valorCobrado / valorEnergiaRecebida) * 100).toFixed(2));

      const movimentacaoId = associadoId * 100 + i + 1;

      const result = await db.runAsync(
        `INSERT INTO movimentacoes (
          id, dataCadastro, dataAtualizacao, associadoId, mes,
          ano, valorTotal, dataVencimento, dataPagamento, statusPagamento, observacoes,
          energiaRecebidaKwh, valorEnergiaRecebida, tarifaUnitariaKwh, valorCobrado,
          valorEconomizado, percentualEconomizado
        )
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
        [
          movimentacaoId, dataCadastro, dataAtualizacao, associadoId, mes,
          ano, valorTotal, dataVencimento, dataPagamento, statusPagamento, observacoes,
          energiaRecebidaKwh, valorEnergiaRecebida, tarifaUnitariaKwh, valorCobrado,
          valorEconomizado, percentualEconomizado
        ]
      );

      if (result?.changes === 1) {
        registrosInseridos++;
      }
    }

    console.log("seedMovimentacoes - Registros de movimentação inseridos:", registrosInseridos);
    return { success: true, data: { count: registrosInseridos } };

  } catch (error) {
    console.error("seedMovimentacoes - Erro durante a inserção de movimentações:", error);
    return { success: false, error: "Erro ao inserir movimentações para o associado." };
  } finally {
    console.log("=== TERMINO ========================================================================");
  }
}
