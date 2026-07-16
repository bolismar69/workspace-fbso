import { useQuery } from "@tanstack/react-query";
import { useAssociadosCopilot } from "@/services/database/useAssociadosCopilot";
import { DBResponse } from "@/types/DBResponse";
import { AssociadoType } from "@/types/AssociadoType";

export function useQueryAssociadosSearchAll(enabled: boolean = true) {
  const { searchAll } = useAssociadosCopilot();

  return useQuery<DBResponse<AssociadoType[]>, Error>({
    queryKey: ["associados", "all"],
    queryFn: async () => {
      const result = await searchAll();
      if (!result.success) {
        throw new Error(result.error || "Erro ao buscar associados");
      }
      return result;
    },
    enabled,
    staleTime: 1000 * 60 * 5,
    refetchOnWindowFocus: false,
  });
}
