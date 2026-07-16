// src/hooks/mutations/useSeedMovimentacoesMutation.ts
import { useMutation } from "@tanstack/react-query";
import { seedMovimentacoes } from "@/services/database/seedMovimentacoes";
import { DBResponse } from "@/types/DBResponse";
import * as SQLite from "expo-sqlite";

type SeedMovimentacoesInput = {
  associadoId: number;
  db?: SQLite.SQLiteDatabase;
};

export function useMutationSeedMovimentacoes() {
  return useMutation<DBResponse<{ count: number }>, Error, SeedMovimentacoesInput>({
    mutationFn: async ({ associadoId, db }) => {
      try {
        if (!db) {
          throw new Error("useSeedMovimentacoesMutation - Conexão com o banco de dados não fornecida.");
        }
        const result = await seedMovimentacoes(associadoId, db);
        return result;
      } catch (error) {
        console.error("useSeedMovimentacoesMutation - Erro:", error);
        return { success: false, error: "useSeedMovimentacoesMutation - Erro ao realizar seed de movimentações" };
      }
    },
  });
}
