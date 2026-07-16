'use client';
// usePlans.ts — ≡ App (Padrão: { data, loading, error })
import { useState, useEffect } from 'react';
import type { Plan } from '@/lib/types';
import { fetchPlans } from '@/services/servicePlans';

export function usePlans() {
  const [data, setData] = useState<Plan[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchPlans()
      .then((result) => {
        setData(result);
        setLoading(false);
      })
      .catch((err) => {
        setError(err instanceof Error ? err.message : 'Erro ao carregar planos');
        setLoading(false);
      });
  }, []);

  return { data, loading, error };
}
