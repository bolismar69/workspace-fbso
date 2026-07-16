# ARCHITECTURE — Solar Fácil Site

> Camadas, padrões arquiteturais e fluxo de dados.
> Gerado por `acquire-codebase-knowledge` em 2026-07-08.
> Fontes: análise do código-fonte em `src/`, `next.config.ts`, `.specs/ARCHITECTURE.md`.

---

## 1. Padrão Arquitetural

**Next.js App Router + Server Components por padrão + Client Components sob demanda.**

- **Páginas** (`page.tsx`): Server Components por padrão (renderizadas no servidor)
- **Componentes interativos**: Marcados com `'use client'` (calculadoras, formulários, menus)
- **Separação**: `lib/` (lógica pura) → `services/` (dados assíncronos) → `hooks/` (estado React) → `components/` (UI)

## 2. Diagrama de Camadas

```
┌─────────────────────────────────────────────┐
│                 components/                  │  ← UI Layer
│  shared/  layout/  home/  plans/  contact/  │
├─────────────────────────────────────────────┤
│                  hooks/                      │  ← State Management
│  useCalculator useContactForm usePlans ...  │
├─────────────────────────────────────────────┤
│                services/                     │  ← Data Layer (async)
│  servicePlans serviceFAQs ...               │
├──────────────────────┬──────────────────────┤
│       lib/           │      types/          │  ← Business Logic + Types
│  calculator.ts       │  concessionaria.ts   │
│  constants.ts        │  consumo-medio.ts    │
│  validation.ts       │  faq.ts              │
│  analytics.ts        │                      │
│  types.ts            │                      │
├──────────────────────┴──────────────────────┤
│                mocks/                        │  ← Static Data (JSON)
│  mockPlans.json mockFAQs.json ...           │
└─────────────────────────────────────────────┘
```

## 3. Fluxo de Dados

### 3.1 Páginas Server Components (Homepage, Planos)

```
Server (build/request time)
  └→ page.tsx (Server Component)
       └→ constants.ts (import direto — sem fetch)
            └→ Renderiza HTML estático
                 └→ Enviado ao cliente
```

### 3.2 Componentes Client (Calculadora)

```
Client (browser)
  └→ useCalculator hook
       └→ calculator.ts (funções puras)
            ├→ calculateConsumerEconomy(monthlyBill) → ConsumerResult
            └→ calculateProviderGain(surplusKwh) → ProviderResult
       └→ analytics.ts → trackCalculatorUse() → window.gtag()
```

### 3.3 Componentes Client (Formulário de Contato)

```
Client (browser)
  └→ useContactForm hook
       ├→ validation.ts → validateForm(values) → FormErrors
       └→ fetch(FORM_ENDPOINT) → POST FormData
            └→ SuccessScreen | submitError
```

### 3.4 Serviços com Dados Mock (Planos, FAQs, etc.)

```
Client (browser)
  └→ usePlans hook
       └→ servicePlans.fetchPlans()
            └→ setTimeout(() => resolve(PLANS), 500)  ← Simula latência de API
                 └→ Retorna Plan[]
```

## 4. Decisões Arquiteturais

### 4.1 Dados Estáticos vs. API Real

**Decisão**: Todos os dados são estáticos (constantes TypeScript + JSON mocks).
**Motivação**: MVP sem backend. A camada `services/` foi projetada com delay de 500ms para simular latência real e facilitar migração futura para API.
**Consequência**: Sem chamadas de rede reais para dados de negócio. Apenas o formulário de contato faz `fetch` para endpoint externo (Formspree placeholder).

### 4.2 Sem Gerenciador de Estado Global

**Decisão**: Usar apenas `useState`/`useCallback` locais por hook.
**Motivação**: App pequena (3 páginas, 25 componentes). Estado não precisa ser compartilhado entre páginas distantes.
**Consequência**: Código mais simples, zero boilerplate de Redux/Zustand/Context. Se a app crescer, migrar para Context ou Zustand.

### 4.3 Server Components por Padrão

**Decisão**: Páginas são Server Components. Apenas componentes interativos são `'use client'`.
**Motivação**: Melhor performance (menos JS enviado ao cliente), SEO nativo (metadata server-side).
**Consequência**: Limite claro entre o que roda no servidor (imports de constantes, metadata) e no cliente (calculadoras, formulários).

### 4.4 Output Standalone

**Decisão**: `output: 'standalone'` em vez do default Vercel.
**Motivação**: Hospedagem em DigitalOcean ou similar (não-Vercel). Gera bundle Node.js autossuficiente.
**Consequência**: Sem edge functions, sem ISR via Vercel, sem analytics de plataforma.

### 4.5 Tailwind CSS v4 (CSS-first)

**Decisão**: Tailwind v4 com `@theme inline` em vez de `tailwind.config.ts`.
**Motivação**: Tailwind v4 é CSS-first — tokens de design são definidos diretamente no CSS, não em JS.
**Consequência**: `globals.css` é o source of truth para design tokens (`--color-solar-*`).

## 5. Cross-Cutting Concerns

| Concern | Implementação | Localização |
|---|---|---|
| **Autenticação** | ❌ Não implementada | — |
| **Autorização** | ❌ Não implementada | — |
| **Error Handling** | try-catch nos hooks, mensagens em pt-BR | `useContactForm.ts`, `usePlans.ts`, etc. |
| **Loading States** | Skeleton component, `isSubmitting`, flags booleanas | `Skeleton.tsx`, hooks |
| **Logging** | `console.error` nos catch blocks | Serviços e hooks |
| **Analytics** | Google Analytics 4 (condicional via env var) | `analytics.ts`, `AnalyticsProvider.tsx` |
| **SEO** | Metadata API, robots.txt, sitemap.xml, JSON-LD | `layout.tsx`, `robots.ts`, `sitemap.ts`, `JsonLd.tsx` |
| **Segurança** | X-Frame-Options, X-Content-Type-Options, Referrer-Policy, honeypot anti-spam | `next.config.ts`, `useContactForm.ts` |
| **Formatação** | Prettier + prettier-plugin-tailwindcss | `package.json` |
| **Linting** | ESLint 9 flat config + eslint-config-next | `eslint.config.mjs` |

## 6. Convenções de Código

| Convenção | Padrão |
|---|---|
| **Naming de arquivos** | PascalCase para componentes, camelCase para hooks/lib/services |
| **Naming de funções** | camelCase (ex: `calculateConsumerEconomy`) |
| **Naming de tipos** | PascalCase (ex: `ConsumerResult`, `LeadForm`) |
| **Exports** | Named exports (sem `export default`) |
| **Imports** | Path alias `@/` → `src/` |
| **Ordenação** | React → Next → libs externas → internos (`@/`) |
| **Componentes** | Funções com tipo implícito (sem `React.FC`) |
| **Estados** | `useState` + `useCallback` (sem Redux/Zustand) |

---

Última atualização: 2026-07-08
