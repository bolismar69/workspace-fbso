'use client';

import { useCalculator } from '@/hooks/useCalculator';
import { Zap } from 'lucide-react';
import Button from '@/components/shared/Button';
import type { ProviderResult } from '@/lib/types';
import { PROVIDER_RATE } from '@/lib/constants';

export default function ProviderCalculator() {
  const { input, setInput, result, error, hasCalculated, calculate, reset } =
    useCalculator('provider');

  const providerResult = result as ProviderResult | null;
  const contactHref = providerResult && !providerResult.isOutlier
    ? `/contato?perfil=fornecedor&excedente=${input}&ganho=${providerResult.gain}`
    : '/contato?perfil=fornecedor';

  return (
    <div className="rounded-lg border border-solar-border bg-solar-bg p-4 shadow-sm md:p-4">
      <div className="mb-4 flex items-center gap-2">
        <Zap size={24} className="text-solar-secondary" />
        <h3 className="text-xl font-bold text-solar-text">Fornecedor</h3>
      </div>

      <p className="mb-4 text-sm text-solar-text-muted">
        Quanto seu excedente pode{' '}
        <span className="font-semibold text-solar-secondary">render</span>?
      </p>

      {/* Input */}
      <div className="mb-4">
        <label htmlFor="provider-input" className="mb-1 block text-base font-bold text-solar-text">
          Quantos kWh você gera de excedente por mês?
        </label>
        <div className="relative">
          <input
            id="provider-input"
            type="number"
            min="0"
            step="1"
            inputMode="numeric"
            placeholder="300"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            onKeyDown={(e) => e.key === 'Enter' && calculate()}
            className={`w-full rounded-lg border px-3 py-2.5 pr-12 text-base transition-colors
              focus:outline-none focus:ring-2 focus:ring-solar-secondary
              ${error ? 'border-solar-error-border bg-solar-error-surface' : 'border-solar-border'}`}
          />
          <span className="absolute right-3 top-1/2 -translate-y-1/2 text-sm text-solar-text-muted">
            kWh
          </span>
        </div>
        {error && <p className="mt-1 text-sm text-solar-error-text">{error}</p>}
      </div>

      {/* Reference rate */}
      <p className="mb-3 text-xs text-solar-text-muted">
        Referência:{' '}
        <span className="font-semibold text-solar-secondary">
          {new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(PROVIDER_RATE)}/kWh
        </span>
      </p>

      {/* Action buttons */}
      <div className="flex gap-2">
        <Button variant="secondary" size="md" onClick={calculate} className="flex-1">
          Simular
        </Button>
        {hasCalculated && (
          <Button variant="outline" size="md" onClick={reset}>
            Limpar
          </Button>
        )}
      </div>

      {/* Result */}
      {providerResult && (
        <div role="status" className="mt-6 motion-safe:animate-fade-in rounded-lg bg-solar-secondary-light p-4">
          {providerResult.isOutlier ? (
            <p className="text-sm text-solar-text">{providerResult.message}</p>
          ) : (
            <>
              <p className="text-lg font-bold text-solar-secondary-dark">
                {providerResult.message}
              </p>
              <div className="mt-4">
                <Button variant="secondary" size="sm" href={contactHref}>
                  Cadastrar meus painéis
                </Button>
              </div>
            </>
          )}
        </div>
      )}
    </div>
  );
}
