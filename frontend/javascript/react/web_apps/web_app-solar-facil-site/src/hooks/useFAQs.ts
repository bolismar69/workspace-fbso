'use client';
// useFAQs.ts — ≡ App (Padrão: { data, loading, error })
import { useState, useEffect } from 'react';
import { fetchFAQs, type FAQItem } from '@/services/serviceFAQs';

export function useFAQs() {
  const [data, setData] = useState<FAQItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchFAQs()
      .then((result) => {
        setData(result);
        setLoading(false);
      })
      .catch((err) => {
        setError(err instanceof Error ? err.message : 'Erro ao carregar FAQs');
        setLoading(false);
      });
  }, []);

  return { data, loading, error };
}
