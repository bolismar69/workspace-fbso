import { useQuery } from "@tanstack/react-query";
import { useAssociadosCopilot } from "@/services/database/useAssociadosCopilot";
import { DBResponse } from "@/types/DBResponse";
import { AssociadoType } from "@/types/AssociadoType";

export function useQueryAssociadosSearchById(id: number, enabled: boolean = true) {
  const { searchById } = useAssociadosCopilot();

  return useQuery<DBResponse<AssociadoType[]>, Error>({
    queryKey: ["associados", "byId", id],
    queryFn: async () => {
      const result = await searchById(id);
      if (!result.success) {
        throw new Error(result.error || "Erro ao buscar associado por ID");
      }
      return result;
    },
    enabled: enabled && !!id,
    staleTime: 1000 * 60 * 5,
    refetchOnWindowFocus: false,
  });
}
