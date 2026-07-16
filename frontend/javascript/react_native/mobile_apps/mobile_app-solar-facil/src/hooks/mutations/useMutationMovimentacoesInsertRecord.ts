import { useMutation } from "@tanstack/react-query";
import { useMovimentacoesCopilot } from "@/services/database/useMovimentacoesCopilot";
import { MovimentacaoMensalType } from "@/types/MovimentacaoMensalType";
import { DBResponse } from "@/types/DBResponse";

export function useMutationMovimentacoesInsertRecord() {
  const { insertRecord } = useMovimentacoesCopilot();

  return useMutation<DBResponse<{ rowID: number }>, Error, MovimentacaoMensalType>({
    mutationFn: async (movimentacao: MovimentacaoMensalType) => {
      try {
        const response = await insertRecord(movimentacao);
        if (response.success && response.data !== undefined) {
          return {
            success: true,
            data: { rowID: response.data },
          };
        } else {
          return {
            success: false,
            error: "useInsertMovimentacaoMutation - Erro ao inserir movimentação" + response.error,
          };
        }
      } catch (error) {
        console.error("useInsertMovimentacaoMutation - Erro:", error);
        return {
          success: false,
          error: "useInsertMovimentacaoMutation - Erro ao inserir movimentação",
        };
      }
    },
  });
}
