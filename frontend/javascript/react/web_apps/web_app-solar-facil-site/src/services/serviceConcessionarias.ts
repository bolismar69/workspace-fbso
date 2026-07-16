// src/services/serviceConcessionarias.ts — ≡ App (Padrão: async/Promise/delay 500ms/try-catch/fallback [])
import type { ConcessionariaType } from '@/types/concessionaria';
import concessionarias from '@/mocks/mockConcessionarias.json';

export async function fetchConcessionarias(): Promise<ConcessionariaType[]> {
  return new Promise((resolve) => {
    try {
      setTimeout(() => {
        resolve(concessionarias as ConcessionariaType[]);
      }, 500);
    } catch (error) {
      console.error("Erro ao buscar dados de concessionárias => ", error);
      resolve([]);
    }
  });
}
