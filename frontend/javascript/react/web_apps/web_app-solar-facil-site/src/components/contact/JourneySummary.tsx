'use client';

import { useSearchParams } from 'next/navigation';
import { CheckCircle } from 'lucide-react';

export default function JourneySummary() {
  const searchParams = useSearchParams();
  const plano = searchParams.get('plano');
  const economia = searchParams.get('economia');
  const excedente = searchParams.get('excedente');
  const ganho = searchParams.get('ganho');

  const hasConsumerContext = plano && economia;
  const hasProviderContext = excedente && ganho;

  if (!hasConsumerContext && !hasProviderContext) {
    return null;
  }

  return (
    <div className="mb-6 rounded-lg border border-solar-primary/20 bg-solar-primary-light/50 p-4">
      <div className="flex items-start gap-3">
        <CheckCircle size={20} className="mt-0.5 flex-shrink-0 text-solar-primary" />
        <div>
          <p className="text-sm font-medium text-solar-text">
            {hasConsumerContext ? 'Simulação concluída' : 'Cálculo realizado'}
          </p>
          {hasConsumerContext && (
            <p className="mt-1 text-sm text-solar-text-muted">
              Plano sugerido:{' '}
              <span className="font-semibold text-solar-primary">{plano}</span>
              {' — '}
              Economia estimada:{' '}
              <span className="font-semibold text-solar-primary">
                R$ {economia}/mês
              </span>
            </p>
          )}
          {hasProviderContext && (
            <p className="mt-1 text-sm text-solar-text-muted">
              Excedente informado:{' '}
              <span className="font-semibold text-solar-secondary-dark">
                {excedente} kWh/mês
              </span>
              {' — '}
              Ganho estimado:{' '}
              <span className="font-semibold text-solar-secondary-dark">
                R$ {ganho}/mês
              </span>
            </p>
          )}
        </div>
      </div>
    </div>
  );
}
