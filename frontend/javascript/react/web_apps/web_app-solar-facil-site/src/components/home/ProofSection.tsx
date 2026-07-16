'use client';

import { useSearchParams } from 'next/navigation';
import MetricCard from '@/components/shared/MetricCard';
import { METRICS, PROVIDER_METRICS } from '@/lib/constants';
import { ShieldCheck } from 'lucide-react';

export default function ProofSection() {
  const searchParams = useSearchParams();
  const persona = searchParams.get('perfil');
  const showProviderMetrics = persona === 'fornecedor';
  const highlightConsumer = persona === 'consumidor';

  return (
    <section
      id="provas"
      className="bg-solar-primary px-4 py-20 md:py-28"
    >
      <div className="mx-auto max-w-5xl">
        <h2 className="text-center text-3xl font-bold text-white md:text-4xl [text-wrap:balance]">
          Quem já confia na Solar Fácil
        </h2>

        <div className="mt-12 grid gap-6 sm:grid-cols-3">
          {METRICS.map((metric) => {
            const isHighlighted =
              (highlightConsumer && metric.label.includes('desconto')) ||
              (showProviderMetrics && metric.label.includes('cooperativa'));

            return (
              <MetricCard
                key={metric.label}
                value={metric.value}
                label={metric.label}
                highlighted={isHighlighted}
                inverted
              />
            );
          })}
        </div>

        {showProviderMetrics && (
          <div className="mt-6 grid gap-6 sm:grid-cols-3">
            {PROVIDER_METRICS.map((metric) => (
              <MetricCard
                key={metric.label}
                value={metric.value}
                label={metric.label}
                highlighted
                inverted
              />
            ))}
          </div>
        )}

        {/* ANEEL Seal — inverted */}
        <div className="mx-auto mt-12 flex max-w-md items-center gap-3 rounded-lg border border-white/25 bg-white/10 p-4 text-center backdrop-blur-sm">
          <ShieldCheck className="h-8 w-8 flex-shrink-0 text-white" />
          <p className="text-sm text-white/80">
            <span className="font-semibold text-white">
              Plataforma regulamentada
            </span>{' '}
            pela ANEEL (RN 687/2015). Operação dentro das normas do setor
            elétrico brasileiro.
          </p>
        </div>
      </div>
    </section>
  );
}
