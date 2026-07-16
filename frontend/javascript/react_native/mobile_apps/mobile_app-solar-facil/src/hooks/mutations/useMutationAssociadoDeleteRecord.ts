import { useMutation } from "@tanstack/react-query";
import { useAssociadosCopilot } from "@/services/database/useAssociadosCopilot";
import { DBResponse } from "@/types/DBResponse";

export function useMutationAssociadoDeleteRecord() {
  const { deleteRecord } = useAssociadosCopilot();

  return useMutation<DBResponse<{ rowID: number; numberChanges: number }>, Error, number>({
    mutationFn: async (id: number) => {
      try {
        const result = await deleteRecord(id);
        return {
          success: true,
          data: result,
        };
      } catch (error) {
        console.error("useDeleteAssociadoMutation - Erro:", error);
        return {
          success: false,
          error: "Erro ao deletar associado",
        };
      }
    },
  });
}
