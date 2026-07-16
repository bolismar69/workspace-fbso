'use client';

import { useState, useCallback } from 'react';
import type { ConsumerResult, ProviderResult } from '@/lib/types';
import { calculateConsumerEconomy, calculateProviderGain } from '@/lib/calculator';
import { trackCalculatorUse } from '@/lib/analytics';

type CalculatorMode = 'consumer' | 'provider';

interface CalculatorState {
  input: string;
  result: ConsumerResult | ProviderResult | null;
  error: string | null;
  hasCalculated: boolean;
}

export function useCalculator(mode: CalculatorMode) {
  const [state, setState] = useState<CalculatorState>({
    input: '',
    result: null,
    error: null,
    hasCalculated: false,
  });

  const setInput = useCallback((value: string) => {
    setState((prev) => ({ ...prev, input: value, error: null }));
  }, []);

  const calculate = useCallback(() => {
    const numericValue = parseFloat(state.input);

    if (!state.input || isNaN(numericValue) || numericValue <= 0) {
      setState((prev) => ({
        ...prev,
        error: 'Informe um valor para calcular',
        result: null,
      }));
      return;
    }

    setState((prev) => ({ ...prev, error: null }));

    if (mode === 'consumer') {
      const result = calculateConsumerEconomy(numericValue);
      setState((prev) => ({ ...prev, result, hasCalculated: true }));
      trackCalculatorUse({
        persona: 'consumidor',
        input_value: numericValue,
        result: result.economy,
        plan_suggested: result.suggestedPlan ?? undefined,
      });
    } else {
      const result = calculateProviderGain(numericValue);
      setState((prev) => ({ ...prev, result, hasCalculated: true }));
      trackCalculatorUse({
        persona: 'fornecedor',
        input_value: numericValue,
        result: result.gain,
      });
    }
  }, [state.input, mode]);

  const reset = useCallback(() => {
    setState({ input: '', result: null, error: null, hasCalculated: false });
  }, []);

  return {
    input: state.input,
    setInput,
    result: state.result,
    error: state.error,
    hasCalculated: state.hasCalculated,
    calculate,
    reset,
  };
}
