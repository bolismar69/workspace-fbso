import { useQuery } from "@tanstack/react-query";
import { useMovimentacoesCopilot } from "@/services/database/useMovimentacoesCopilot";
import { DBResponse } from "@/types/DBResponse";
import { MovimentacaoMensalType } from "@/types/MovimentacaoMensalType";

export function useQueryMovimentacoesSearchByAssociadoId(associadoId: number, enabled: boolean = true) {
  const { searchByAssociadoId } = useMovimentacoesCopilot();

  return useQuery<DBResponse<MovimentacaoMensalType[]>, Error>({
    queryKey: ["movimentacoes", associadoId],
    queryFn: async () => {
      const result = await searchByAssociadoId(associadoId);
      if (!result.success) {
        throw new Error(result.error || "Erro ao buscar movimentações");
      }
      return result;
    },
    enabled: !!associadoId && enabled,
    staleTime: 1000 * 60 * 2,
    refetchOnWindowFocus: true,
  });
}
