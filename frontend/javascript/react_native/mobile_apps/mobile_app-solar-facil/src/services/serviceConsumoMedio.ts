import mockConsumoMedio from "@/mocks/mockConsumoMedio.json";
import { ConsumoMedioType } from "@/types/ConsumoMedioType";

const SIMULATION_DELAY = 500; // Configuração do tempo de simulação

export async function fetchConsumoMedio(): Promise<ConsumoMedioType[]> {
  return new Promise((resolve, reject) => {
    try {
      setTimeout(() => {
      resolve(mockConsumoMedio as ConsumoMedioType[]);
      }, SIMULATION_DELAY);
    } catch (error) {
      console.error("Erro ao buscar dados de consumo médio => ", error);
      resolve([]);
    }
  });
}

export async function fetchConsumoMedioOptions(): Promise<{ label: string; value: number }[]> {
  try {
    const consumoMedio = await fetchConsumoMedio();
    const activeConsumoMedio = consumoMedio.filter((item) => item.status === "active");
    return activeConsumoMedio.map((item) => ({
      label: item.name,
      value: item.id,
    }));
  } catch (error) {
    throw new Error("Erro ao transformar dados de consumo médio em opções => " + error);
  }
}
