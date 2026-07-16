// src/services/serviceConsumoMedio.ts — ≡ App (Padrão: async/Promise/delay 500ms/try-catch/fallback [])
import type { ConsumoMedioType } from '@/types/consumo-medio';
import consumoMedio from '@/mocks/mockConsumoMedio.json';

export async function fetchConsumoMedio(): Promise<ConsumoMedioType[]> {
  return new Promise((resolve) => {
    try {
      setTimeout(() => {
        resolve(consumoMedio as ConsumoMedioType[]);
      }, 500);
    } catch (error) {
      console.error("Erro ao buscar dados de consumo médio => ", error);
      resolve([]);
    }
  });
}
