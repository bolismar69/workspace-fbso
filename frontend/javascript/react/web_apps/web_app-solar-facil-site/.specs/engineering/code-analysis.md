# Code Analysis — Solar Fácil Site

> Análise técnica de fluxo de código: componentes, hooks, dados e renderização.
> Gerado por `documentation-writer` em 2026-07-08.

---

## 1. Fluxo de Renderização

### 1.1 Página Inicial (/)

```
Requisição GET /
  └── layout.tsx (Server)
       ├── <html lang="pt-BR">
       ├── Metadata (title, description, openGraph)
       ├── Inter font (next/font — build time)
       └── <body>
            ├── <AnalyticsProvider> (Client — carrega GA4 se configurado)
            ├── <Header> (Server — contém <MobileMenu> Client)
            ├── page.tsx (Server)
            │    ├── <HeroSection> (Server — estático)
            │    ├── <CalculatorSection> (Server — wrapper)
            │    │    ├── <ConsumerCalculator> (Client)
            │    │    │    └── useCalculator('consumer')
            │    │    │         ├── useState: input, result, error, hasCalculated
            │    │    │         └── calculate() → calculator.ts → ConsumerResult
            │    │    └── <ProviderCalculator> (Client)
            │    │         └── useCalculator('provider')
            │    ├── <ProofSection> (Server — METRICS constant)
            │    ├── <HowItWorksSection> (Server — STEPS constant)
            │    ├── <PlansSection> (Client)
            │    │    └── usePlans() → servicePlans.fetchPlans() → PLANS[]
            │    ├── <DifferentiatorsSection> (Server — DIFFERENTIATORS)
            │    └── <FinalCtaSection> (Server — estático)
            └── <Footer> (Server — estático)
```

### 1.2 Página de Planos (/planos)

```
Requisição GET /planos
  └── layout.tsx (Server — mesmo layout)
       └── planos/page.tsx (Server)
            ├── <PlansComparisonTable> (Server — PLANS constant)
            ├── <ProviderHighlight> (Server — PROVIDER_METRICS)
            └── <FaqAccordion> (Client)
                 └── useFaqAccordion() → useState(index)
```

### 1.3 Página de Contato (/contato)

```
Requisição GET /contato
  └── layout.tsx (Server — mesmo layout)
       └── contato/page.tsx (Client)
            ├── <JourneySummary> (Client — searchParams)
            ├── <ContactForm> (Client)
            │    └── useContactForm({initialProfile, initialMessage})
            │         ├── useState: values, errors, isSubmitting, isSuccess
            │         ├── setField() → validação inline
            │         └── handleSubmit()
            │              ├── validateForm() → FormErrors
            │              ├── Anti-spam: elapsed < 3s
            │              └── fetch(FORM_ENDPOINT, POST FormData)
            ├── <DirectChannels> (Client)
            └── <SuccessScreen> (Client — condicional: isSuccess)
```

---

## 2. Fluxo de Dados

```
┌─────────────────────────────────────────────────────┐
│                   DATA SOURCES                       │
│                                                      │
│  constants.ts          mocks/*.json                  │
│  (PLANS, FAQ_ITEMS,    (mockPlans, mockFAQs,        │
│   METRICS, STEPS,       mockConcessionarias,         │
│   DIFFERENTIATORS)      mockConsumoMedio)            │
│                                                      │
└────────────┬────────────────┬───────────────────────┘
             │                │
             ▼                ▼
┌─────────────────┐  ┌─────────────────┐
│  Server          │  │  Services       │
│  Components      │  │  (async, 500ms) │
│  (import direto) │  │                 │
└────────┬────────┘  └────────┬────────┘
         │                    │
         ▼                    ▼
┌──────────────────────────────────────┐
│              HOOKS                    │
│  usePlans, useFAQs, useCalculator... │
└────────────────┬─────────────────────┘
                 │
                 ▼
┌──────────────────────────────────────┐
│           COMPONENTS                  │
│  PlansSection, FaqAccordion, etc.    │
└──────────────────────────────────────┘
```

---

## 3. Decisões de Código Notáveis

### 3.1 Funções Puras em lib/

```typescript
// lib/calculator.ts — testável, sem side effects
export function calculateConsumerEconomy(monthlyBill: number): ConsumerResult

// lib/validation.ts — testável, sem side effects  
export function validateEmail(email: string): string | null
```

### 3.2 Serviços com Delay Controlado

O delay de 500ms nos serviços é intencional — prepara a UX para quando houver API real. Se removido, a transição loading→loaded seria instantânea e o Skeleton nunca apareceria.

### 3.3 Anti-Spam no Formulário

Duas camadas de proteção:
1. **Honeypot**: campo hidden `website` (bots preenchem)
2. **Tempo mínimo**: 3s entre page load e submit (bots submetem instantaneamente)

Se detectado como bot → fake success (não alerta o atacante).

### 3.4 Analytics Condicional

```typescript
// Só dispara se script GA4 foi carregado
export function trackEvent(event, params?) {
  if (typeof window !== 'undefined' && window.gtag) {
    window.gtag('event', event, params);
  }
}
```

---

## 4. Pontos de Extensão

| Local | Propósito | Como extender |
|---|---|---|
| `services/` | Trocar mock por API real | Substituir `setTimeout` por `fetch` |
| `lib/calculator.ts` | Novas regras de cálculo | Adicionar funções puras |
| `lib/validation.ts` | Novas validações | Adicionar funções de validação |
| `hooks/` | Novo estado compartilhado | Criar hook com `useState`/`useCallback` |
| `constants.ts` | Novos dados estáticos | Adicionar arrays/objects |

---

Última atualização: 2026-07-08
