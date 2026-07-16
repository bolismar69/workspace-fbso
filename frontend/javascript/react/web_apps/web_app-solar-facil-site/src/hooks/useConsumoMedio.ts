'use client';
// useConsumoMedio.ts — ≡ App (Padrão: { data, loading, error })
import { useState, useEffect } from 'react';
import type { ConsumoMedioType } from '@/types/consumo-medio';
import { fetchConsumoMedio } from '@/services/serviceConsumoMedio';

export function useConsumoMedio() {
  const [data, setData] = useState<ConsumoMedioType[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchConsumoMedio()
      .then((result) => {
        setData(result);
        setLoading(false);
      })
      .catch((err) => {
        setError(err instanceof Error ? err.message : 'Erro ao carregar dados de consumo médio');
        setLoading(false);
      });
  }, []);

  return { data, loading, error };
}
