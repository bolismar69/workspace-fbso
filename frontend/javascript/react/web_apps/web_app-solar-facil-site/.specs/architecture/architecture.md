# Arquitetura — Solar Fácil Site

> Visão geral consolidada da arquitetura: stack, padrões, decisões e cross-cutting concerns.
> Gerado por `architecture-designer` em 2026-07-08.
> Integra conteúdo de C4 Context, C4 Containers, C4 Components, ADRs e Integrations.

---

## 1. Resumo Executivo

O **Solar Fácil Site** é um portal web construído com **Next.js 16 App Router + React 19 + Tailwind CSS v4 + TypeScript**. Adota o padrão **Server Components por padrão** (páginas renderizadas no servidor) com **Client Components sob demanda** (calculadoras, formulários). Opera sem backend — todos os dados são estáticos (constantes TypeScript + JSON mocks). O output `standalone` permite deploy em qualquer servidor Node.js 20+.

## 2. Padrão Arquitetural

```
Server Components (default)  ←→  Client Components ('use client')
        │                               │
        │ Renderizados no servidor       │ Hidratados no browser
        │ SEO + performance             │ Interatividade + estado
        │                               │
        ▼                               ▼
   Páginas estáticas              Calculadoras, Formulários
   (homepage, planos)             Analytics, Menus
```

## 3. Stack Tecnológica

| Camada | Tecnologia | Versão |
|---|---|---|
| Framework | Next.js (App Router) | 16.2.10 |
| UI Library | React | 19.2.4 |
| Estilização | Tailwind CSS v4 + CSS Custom Properties | ^4 |
| Linguagem | TypeScript (strict) | ^5 |
| Ícones | Lucide React | ^1.23.0 |
| Lint | ESLint 9 + eslint-config-next | ^9 |
| Format | Prettier + prettier-plugin-tailwindcss | ^3.9.4 |
| Deploy | Standalone (DigitalOcean) | — |

## 4. Estrutura de Camadas

```
src/
├── app/            ← App Router: páginas, layout, globals
├── components/     ← UI Layer: shared, layout, home, plans, contact
│   ├── shared/     ← 7 componentes reutilizáveis
│   ├── layout/     ← 3 componentes de layout
│   ├── home/       ← 9 seções da homepage
│   ├── plans/      ← 4 componentes de planos
│   └── contact/    ← 5 componentes de contato
├── hooks/          ← State Layer: 7 custom hooks
├── services/       ← Data Layer: 4 serviços assíncronos (mock)
├── lib/            ← Business Logic: 5 módulos de funções puras
├── types/          ← External Types: 3 definições de tipo
└── mocks/          ← Static Data: 4 arquivos JSON mock
```

## 5. Árvore de Componentes Principal

```
RootLayout (Server)
├── AnalyticsProvider (Client)
├── Header (Server)
│   └── MobileMenu (Client)
├── [Page Content]
│   ├── HomePage (Server)
│   │   ├── HeroSection (Server)
│   │   ├── CalculatorSection (Server)
│   │   │   ├── ConsumerCalculator (Client) → useCalculator
│   │   │   └── ProviderCalculator (Client) → useCalculator
│   │   ├── ProofSection (Server)
│   │   ├── HowItWorksSection (Server)
│   │   ├── PlansSection (Client) → usePlans
│   │   ├── DifferentiatorsSection (Server)
│   │   └── FinalCtaSection (Server)
│   ├── PlanosPage (Server)
│   │   ├── PlansComparisonTable (Server)
│   │   ├── ProviderHighlight (Server)
│   │   └── FaqAccordion (Client) → useFaqAccordion
│   └── ContatoPage (Client)
│       ├── JourneySummary (Client)
│       ├── ContactForm (Client) → useContactForm
│       ├── DirectChannels (Client)
│       └── SuccessScreen (Client)
└── Footer (Server)
```

## 6. Fluxo de Dados

```
Dados Estáticos (build time)
  constants.ts, mock*.json
    ↓
Serviços (async, 500ms delay)
  servicePlans, serviceFAQs, ...
    ↓
Hooks (state management)
  usePlans, useFAQs, ...
    ↓
Componentes (UI)
  PlansSection, FaqAccordion, ...

Dados Dinâmicos (client-side only)
  Calculadora: input → lib/calculator.ts (função pura) → resultado
  Formulário: input → lib/validation.ts → fetch(FORM_ENDPOINT) → resposta
  Analytics: eventos → analytics.ts → window.gtag()
```

## 7. Cross-Cutting Concerns

| Concern | Implementação | Localização |
|---|---|---|
| **Autenticação** | ❌ Não implementada | — |
| **Error Handling** | try-catch com mensagens pt-BR | hooks, services |
| **Loading States** | Skeleton + flags booleanas | Skeleton.tsx, hooks |
| **SEO** | Metadata, robots, sitemap, JSON-LD | layout.tsx, robots.ts, sitemap.ts |
| **Segurança** | Headers HTTP, honeypot anti-spam | next.config.ts, useContactForm.ts |
| **Analytics** | GA4 condicional, 4 eventos | analytics.ts, AnalyticsProvider.tsx |
| **Acessibilidade** | WCAG AA planejado | PRODUCT.md |

## 8. Diagramas C4

- [C4 Context](c4-context.md) — Sistema no ecossistema (atores + sistemas externos)
- [C4 Containers](c4-containers.md) — Containers de deploy (Nginx, Node.js, Browser SPA)
- [C4 Components](c4-components.md) — Árvore de componentes React (25 componentes, 7 hooks)

## 9. Decisões Arquiteturais (ADRs)

- [ADR-001](adrs/adr-001.md) — Next.js 16 App Router + React 19
- [ADR-002](adrs/adr-002.md) — Dados estáticos sem backend
- [ADR-003](adrs/adr-003.md) — Tailwind CSS v4 + CSS Custom Properties
- [ADR-004](adrs/adr-004.md) — Output standalone para hospedagem não-Vercel

## 10. Rotas

| Rota | Tipo | Descrição |
|---|---|---|
| `/` | Server Component | Homepage (8 seções) |
| `/planos` | Server Component | Página de planos + FAQ |
| `/contato` | Client Component | Página de contato + formulário |

---

Última atualização: 2026-07-08
