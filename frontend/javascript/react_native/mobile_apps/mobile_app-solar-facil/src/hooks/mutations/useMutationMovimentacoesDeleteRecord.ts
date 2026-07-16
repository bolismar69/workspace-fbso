import { useMutation } from "@tanstack/react-query";
import { useMovimentacoesCopilot } from "@/services/database/useMovimentacoesCopilot";
import { DBResponse } from "@/types/DBResponse";

export function useDeleteMovimentacaoMutation() {
  const { deleteRecord } = useMovimentacoesCopilot();

  return useMutation<DBResponse<{ changes: number }>, Error, number>({
    mutationFn: async (id: number) => {
      try {
        const result = await deleteRecord(id);
        return {
          success: true,
          data: { changes: result.data ?? 0 },
        };
      } catch (error) {
        console.error("useDeleteMovimentacaoMutation - Erro:", error);
        return {
          success: false,
          error: "useDeleteMovimentacaoMutation - Erro ao deletar movimentação",
        };
      }
    },
  });
}
