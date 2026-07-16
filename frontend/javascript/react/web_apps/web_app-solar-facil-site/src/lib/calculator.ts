// Solar Fácil — Calculator business logic (pure functions)
// Refs: BR-DER-001..008, DT-001

import type { ConsumerResult, PlanName, ProviderResult } from './types';
import {
  CONSUMER_DISCOUNT_RATE,
  CONSUMER_OUTLIER_MAX,
  CONSUMER_OUTLIER_MIN,
  PLANS,
  PROVIDER_OUTLIER_MAX,
  PROVIDER_OUTLIER_MIN,
  PROVIDER_RATE,
} from './constants';

/**
 * Format a number as Brazilian Real currency string.
 * @example formatBRL(42) → "R$ 42,00"
 */
export function formatBRL(value: number): string {
  return value.toLocaleString('pt-BR', {
    style: 'currency',
    currency: 'BRL',
  });
}

/**
 * Suggest a plan based on monthly electricity bill range.
 * Decision Table DT-001:
 *   ≤ 200 → Basic
 *   200–350 → Special
 *   350–600 → Premium
 */
export function suggestPlan(monthlyBill: number): PlanName | null {
  for (const plan of PLANS) {
    if (monthlyBill >= plan.capacityKwh.min && monthlyBill <= plan.capacityKwh.max) {
      return plan.name;
    }
  }
  return null;
}

/**
 * Calculate estimated monthly economy for a consumer.
 * BR-DER-001: economy = monthlyBill × 0.12
 * BR-DER-005/006: outliers < 50 or > 5000
 */
export function calculateConsumerEconomy(monthlyBill: number): ConsumerResult {
  if (monthlyBill < CONSUMER_OUTLIER_MIN) {
    return {
      economy: 0,
      suggestedPlan: null,
      isOutlier: true,
      message:
        'Valor informado está abaixo da faixa típica. Para uma simulação personalizada, fale conosco pelo WhatsApp.',
    };
  }

  if (monthlyBill > CONSUMER_OUTLIER_MAX) {
    return {
      economy: 0,
      suggestedPlan: null,
      isOutlier: true,
      message:
        'Para valores acima de R$ 5.000, entre em contato para uma proposta personalizada.',
    };
  }

  const economy = Math.round(monthlyBill * CONSUMER_DISCOUNT_RATE * 100) / 100;
  const plan = suggestPlan(monthlyBill);

  return {
    economy,
    suggestedPlan: plan,
    isOutlier: false,
    message: `Você pode economizar ~${formatBRL(economy)}/mês`,
  };
}

/**
 * Calculate estimated monthly gain for a solar energy provider.
 * BR-DER-003: gain = surplusKwh × R$ 0.40
 * BR-DER-007/008: outliers < 50 kWh or > 10000 kWh
 */
export function calculateProviderGain(monthlySurplusKwh: number): ProviderResult {
  if (monthlySurplusKwh < PROVIDER_OUTLIER_MIN) {
    return {
      gain: 0,
      rate: PROVIDER_RATE,
      isOutlier: true,
      message:
        'Valor mínimo para compartilhamento: 50 kWh/mês.',
    };
  }

  if (monthlySurplusKwh > PROVIDER_OUTLIER_MAX) {
    return {
      gain: 0,
      rate: PROVIDER_RATE,
      isOutlier: true,
      message:
        'Para grandes volumes, solicite contato comercial.',
    };
  }

  const gain = Math.round(monthlySurplusKwh * PROVIDER_RATE * 100) / 100;

  return {
    gain,
    rate: PROVIDER_RATE,
    isOutlier: false,
    message: `Você pode ganhar ~${formatBRL(gain)}/mês compartilhando seu excedente`,
  };
}
