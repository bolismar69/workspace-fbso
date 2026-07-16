import  mockConcessionaria  from "@/mocks/mockConcessionarias.json";
import { ConcessionariaType } from "@/types/ConcessionariaType";

const SIMULATION_DELAY = 500; // Configuração do tempo de simulação

export async function fetchConcessionarias(): Promise<ConcessionariaType[]> {
  return new Promise((resolve, reject) => {
    try{
      setTimeout(() => {
        resolve(mockConcessionaria as ConcessionariaType[]);
      }, SIMULATION_DELAY); // Simula um atraso de 500ms
    } catch (error) {
        reject(new Error("Erro ao buscar dados de concessionarias => " + error));
    }
  });
}

export async function fetchConcessionariasOptions(): Promise<{ label: string; value: number }[]> {
  const plans = await fetchConcessionarias();
  const activeConcessionarias = plans.filter((plan) => plan.status === "active");
  return activeConcessionarias.map((concessionaria) => ({
    label: concessionaria.name,
    value: concessionaria.id,
  }));
}
