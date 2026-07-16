import { useMutation } from "@tanstack/react-query";
import { useMovimentacoesCopilot } from "@/services/database/useMovimentacoesCopilot";
import { MovimentacaoMensalType } from "@/types/MovimentacaoMensalType";
import { DBResponse } from "@/types/DBResponse";

export function useUpdateMovimentacaoMutation() {
  const { updateRecord } = useMovimentacoesCopilot();

  return useMutation<DBResponse<{ changes: number }>, Error, MovimentacaoMensalType>({
    mutationFn: async (movimentacao: MovimentacaoMensalType) => {
      try {
        const response = await updateRecord(movimentacao);
        if (response.success && typeof response.data === "number") {
          return {
            success: true,
            data: { changes: response.data },
          };
        } else {
          return {
            success: false,
            error: "useUpdateMovimentacaoMutation - Erro ao atualizar movimentação" + response.error,
          };
        }
      } catch (error) {
        console.error("useUpdateMovimentacaoMutation - Erro:", error);
        return {
          success: false,
          error: "useUpdateMovimentacaoMutation - Erro ao atualizar movimentação",
        };
      }
    },
  });
}
