import { useQuery } from "@tanstack/react-query";
import { useAssociadosCopilot } from "@/services/database/useAssociadosCopilot";
import { DBResponse } from "@/types/DBResponse";
import { AssociadoType } from "@/types/AssociadoType";

export function useQueryAssociadosSearchByCpfCnpjSenha(
  cpf_cnpj: string,
  senha: string,
  enabled: boolean = true
) {
  const { searchByCpfCnpjSenha } = useAssociadosCopilot();

  return useQuery<DBResponse<AssociadoType[]>, Error>({
    queryKey: ["associados", "byCpfCnpjSenha", cpf_cnpj],
    queryFn: async () => {
      const result = await searchByCpfCnpjSenha(cpf_cnpj, senha);
      if (!result.success) {
        throw new Error(result.error || "Erro ao buscar associado por CPF/CNPJ e Senha");
      }
      return result;
    },
    enabled: enabled && !!cpf_cnpj && !!senha,
    staleTime: 1000 * 60 * 5,
    refetchOnWindowFocus: false,
  });
}
