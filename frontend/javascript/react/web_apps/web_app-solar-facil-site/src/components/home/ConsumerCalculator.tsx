'use client';

import { useCalculator } from '@/hooks/useCalculator';
import { Home } from 'lucide-react';
import Button from '@/components/shared/Button';
import type { ConsumerResult } from '@/lib/types';

export default function ConsumerCalculator() {
  const { input, setInput, result, error, hasCalculated, calculate, reset } =
    useCalculator('consumer');

  const consumerResult = result as ConsumerResult | null;
  const contactHref = consumerResult?.suggestedPlan
    ? `/contato?perfil=consumidor&plano=${consumerResult.suggestedPlan}&economia=${consumerResult.economy}`
    : '/contato?perfil=consumidor';

  return (
    <div className="rounded-lg border border-solar-border bg-solar-bg p-4 shadow-sm md:p-4">
      <div className="mb-4 flex items-center gap-2">
        <Home size={24} className="text-solar-primary" />
        <h3 className="text-xl font-bold text-solar-text">Consumidor</h3>
      </div>

      <p className="mb-4 text-sm text-solar-text-muted">
        Quanto você pode <span className="font-semibold text-solar-primary">economizar</span>?
      </p>

      {/* Input */}
      <div className="mb-4">
        <label htmlFor="consumer-input" className="mb-1 block text-base font-bold text-solar-text">
          Quanto você gasta de luz por mês?
        </label>
        <div className="relative">
          <span className="absolute left-3 top-1/2 -translate-y-1/2 text-solar-text-muted">R$</span>
          <input
            id="consumer-input"
            type="number"
            min="0"
            step="1"
            inputMode="numeric"
            placeholder="350"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            onKeyDown={(e) => e.key === 'Enter' && calculate()}
            className={`w-full rounded-lg border px-3 py-2.5 pl-10 text-base transition-colors
              focus:outline-none focus:ring-2 focus:ring-solar-primary
              ${error ? 'border-solar-error-border bg-solar-error-surface' : 'border-solar-border'}`}
          />
        </div>
        {error && <p className="mt-1 text-sm text-solar-error-text">{error}</p>}
      </div>

      {/* Action buttons */}
      <div className="flex gap-2">
        <Button variant="primary" size="md" onClick={calculate} className="flex-1">
          Calcular
        </Button>
        {hasCalculated && (
          <Button variant="outline" size="md" onClick={reset}>
            Limpar
          </Button>
        )}
      </div>

      {/* Result */}
      {consumerResult && (
        <div role="status" className="mt-6 motion-safe:animate-fade-in rounded-lg bg-solar-primary-light p-4">
          {consumerResult.isOutlier ? (
            <p className="text-sm text-solar-text">{consumerResult.message}</p>
          ) : (
            <>
              <p className="text-lg font-bold text-solar-primary-dark">
                {consumerResult.message}
              </p>
              {consumerResult.suggestedPlan && (
                <p className="mt-2 text-sm text-solar-text">
                  Plano sugerido:{' '}
                  <span className="font-semibold text-solar-primary">
                    {consumerResult.suggestedPlan}
                  </span>
                </p>
              )}
              <div className="mt-4">
                <Button variant="primary" size="sm" href={contactHref}>
                  Quero este plano
                </Button>
              </div>
            </>
          )}
        </div>
      )}
    </div>
  );
}
