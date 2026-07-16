// src/services/beneficiadoService.ts
import { BeneficiadoType } from "@/types/BeneficiadoType";

export const salvarBeneficiadoMock = async (
  data: BeneficiadoType
): Promise<{ sucesso: boolean; mensagem: string }> => {
  console.log("📤 Enviando dados do beneficiado (mock):", data);

  // Simula atraso de rede
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve({
        sucesso: true,
        mensagem: "Dados do beneficiado salvos com sucesso (mock).",
      });
    }, 1200);
  });
};
