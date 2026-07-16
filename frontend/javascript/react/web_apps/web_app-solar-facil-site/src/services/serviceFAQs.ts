// src/services/serviceFAQs.ts — ≡ App (Padrão: async/Promise/delay 500ms/try-catch/fallback [])
// Dados: FAQ_ITEMS de @/lib/constants (Site) — informações de negócio equivalentes ao App
import { FAQ_ITEMS } from '@/lib/constants';

export interface FAQItem {
  question: string;
  answer: string;
}

export async function fetchFAQs(): Promise<FAQItem[]> {
  return new Promise((resolve) => {
    try {
      setTimeout(() => {
        resolve(FAQ_ITEMS);
      }, 500);
    } catch (error) {
      console.error("Erro ao buscar dados de FAQs => ", error);
      resolve([]);
    }
  });
}
