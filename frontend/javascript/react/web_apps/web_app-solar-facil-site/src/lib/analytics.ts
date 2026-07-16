// Solar Fácil — Google Analytics 4 event helpers
// Refs: BR-ACT-004 to BR-ACT-007

declare global {
  interface Window {
    gtag?: (...args: unknown[]) => void;
    dataLayer?: unknown[];
  }
}

export const GA_MEASUREMENT_ID = process.env.NEXT_PUBLIC_GA_ID || 'G-XXXXXXXXXX';

export type AnalyticsEvent =
  | 'cta_click'
  | 'calculator_use'
  | 'faq_open'
  | 'lead_capture';

interface CtaClickParams {
  cta_type: 'consumidor' | 'fornecedor';
  location: 'hero' | 'final_cta';
}

interface CalculatorUseParams {
  persona: 'consumidor' | 'fornecedor';
  input_value: number;
  result: number;
  plan_suggested?: string;
}

interface FaqOpenParams {
  question_index: number;
}

interface LeadCaptureParams {
  persona: string;
  has_plan: boolean;
}

type EventParams =
  | CtaClickParams
  | CalculatorUseParams
  | FaqOpenParams
  | LeadCaptureParams;

export function trackEvent(event: AnalyticsEvent, params?: EventParams) {
  if (typeof window !== 'undefined' && window.gtag) {
    window.gtag('event', event, params);
  }
}

// Convenience wrappers for each event type (BR-ACT-004..007)
export function trackCtaClick(params: CtaClickParams) {
  trackEvent('cta_click', params);
}

export function trackCalculatorUse(params: CalculatorUseParams) {
  trackEvent('calculator_use', params);
}

export function trackFaqOpen(params: FaqOpenParams) {
  trackEvent('faq_open', params);
}

export function trackLeadCapture(params: LeadCaptureParams) {
  trackEvent('lead_capture', params);
}
