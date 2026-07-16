// src/services/servicePlans.ts — ≡ App (Padrão: async/Promise/delay 500ms/try-catch/fallback [])
// Dados: PLANS de @/lib/constants (Site) — informações de negócio equivalentes ao App
import type { Plan } from '@/lib/types';
import { PLANS } from '@/lib/constants';

export async function fetchPlans(): Promise<Plan[]> {
  return new Promise((resolve) => {
    try {
      setTimeout(() => {
        resolve(PLANS);
      }, 500);
    } catch (error) {
      console.error("Erro ao buscar dados de planos => ", error);
      resolve([]);
    }
  });
}

export async function fetchPlanOptions(): Promise<{ label: string; value: number }[]> {
  const plans = await fetchPlans();
  return plans.map((plan, index) => ({
    label: plan.name,
    value: index,
  }));
}
