import { useQuery } from "@tanstack/react-query";
import { useMovimentacoesCopilot } from "@/services/database/useMovimentacoesCopilot";
import { DBResponse } from "@/types/DBResponse";
import { MovimentacaoMensalType } from "@/types/MovimentacaoMensalType";

export function useQueryMovimentacoesSearchAll(enabled: boolean = true) {
  const { searchAll } = useMovimentacoesCopilot();

  return useQuery<DBResponse<MovimentacaoMensalType[]>, Error>({
    queryKey: ["movimentacoes", "all"],
    queryFn: async () => {
      const result = await searchAll();
      if (!result.success) {
        throw new Error(result.error || "Erro ao buscar todas as movimentações");
      }
      return result;
    },
    enabled,
    staleTime: 1000 * 60 * 5,
    refetchOnWindowFocus: false,
  });
}
