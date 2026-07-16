'use client';
// useConcessionarias.ts — ≡ App (Padrão: { data, loading, error })
import { useState, useEffect } from 'react';
import type { ConcessionariaType } from '@/types/concessionaria';
import { fetchConcessionarias } from '@/services/serviceConcessionarias';

export function useConcessionarias() {
  const [data, setData] = useState<ConcessionariaType[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchConcessionarias()
      .then((result) => {
        setData(result);
        setLoading(false);
      })
      .catch((err) => {
        setError(err instanceof Error ? err.message : 'Erro ao carregar concessionárias');
        setLoading(false);
      });
  }, []);

  return { data, loading, error };
}
