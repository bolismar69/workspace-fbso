import { useMutation } from "@tanstack/react-query";
import { useAssociadosCopilot } from "@/services/database/useAssociadosCopilot";
import { AssociadoType } from "@/types/AssociadoType";
import { DBResponse } from "@/types/DBResponse";

export function useMutationAssociadoInsertRecord() {
  const { insertRecord } = useAssociadosCopilot();

  return useMutation<DBResponse<{ rowID: number; numberChanges: number }>, Error, AssociadoType>({
    mutationFn: async (novoAssociado: AssociadoType) => {
      try {
        // insertRecord should return { rowID: number; numberChanges: number }
        const result = await insertRecord(novoAssociado);
        if (result.success && result.data !== undefined) {
          return {
            success: true,
            data: result.data, // unwrap the data property
          };
        } else {
          return {
            success: false,
            error: "Erro ao inserir associado",
          };
        }
      } catch (error) {
        console.error("useInsertAssociadoMutation - Erro:", error);
        return {
          success: false,
          error: "useInsertAssociadoMutation - Erro ao inserir associado",
        };
      }
    },
  });
}
