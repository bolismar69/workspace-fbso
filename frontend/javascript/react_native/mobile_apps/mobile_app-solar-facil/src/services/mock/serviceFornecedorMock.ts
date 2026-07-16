// src/services/serviceFornecedor.ts
import { FornecedorType } from "@/types/FornecedorType";

export const salvarFornecedorMock= async (
  data: FornecedorType
): Promise<{ sucesso: boolean; mensagem: string }> => {
  console.log("📤 Enviando dados do fornecedor:", data);

  // Simula atraso de rede
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve({
        sucesso: true,
        mensagem: "Dados do fornecedor salvos com sucesso.",
      });
    }, 1200);
  });
};
