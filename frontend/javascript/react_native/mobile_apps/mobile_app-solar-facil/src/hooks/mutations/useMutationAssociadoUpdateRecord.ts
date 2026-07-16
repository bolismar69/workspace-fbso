import { useMutation } from "@tanstack/react-query";
import { useAssociadosCopilot } from "@/services/database/useAssociadosCopilot";
import { AssociadoType } from "@/types/AssociadoType";
import { DBResponse } from "@/types/DBResponse";

export function useMutationAssociadoUpdateRecord() {
  const { updateRecord } = useAssociadosCopilot();

  return useMutation<DBResponse<{ rowID: number; numberChanges: number }>, Error, AssociadoType>({
    mutationFn: async (associado: AssociadoType) => {
      try {
        const result = await updateRecord(associado);
        if (result.success && result.data !== undefined) {
          return {
            success: true,
            data: result.data, // unwrap the data property
          }
        }
        else {
          return {
            success: false,
            error: "Erro ao atualizar associado",
          };
        }
      } catch (error) {
        console.error("useUpdateAssociadoMutation - Erro:", error);
        return {
          success: false,
          error: "Erro ao atualizar associado",
        };
      }
    },
  });
}
