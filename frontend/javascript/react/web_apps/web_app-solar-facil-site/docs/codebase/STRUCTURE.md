# STRUCTURE — Solar Fácil Site

> Layout de diretórios, entry points e arquivos chave.
> Gerado por `acquire-codebase-knowledge` em 2026-07-08.
> Fontes: scan de diretórios, `package.json`, `next.config.ts`.

---

## 1. Estrutura de Diretórios

```
solar-facil-site/
├── src/
│   ├── app/                        ← App Router (Next.js 16)
│   │   ├── layout.tsx              ← Root layout (metadata, fontes, providers)
│   │   ├── page.tsx                ← Homepage (Hero→Calculadora→Prova→Planos→CtaFinal)
│   │   ├── globals.css             ← Estilos globais + design tokens (--color-solar-*)
│   │   ├── robots.ts               ← Geração dinâmica de robots.txt
│   │   ├── sitemap.ts              ← Geração dinâmica de sitemap.xml
│   │   ├── contato/
│   │   │   └── page.tsx            ← Página de contato (formulário + canais diretos)
│   │   └── planos/
│   │       └── page.tsx            ← Página de planos (tabela comparativa + FAQ)
│   │
│   ├── components/
│   │   ├── shared/                 ← Componentes reutilizáveis (6)
│   │   │   ├── AnalyticsProvider.tsx ← Provider Google Analytics 4 (condicional)
│   │   │   ├── Breadcrumb.tsx        ← Navegação breadcrumb
│   │   │   ├── Button.tsx            ← Botão reutilizável (variantes: primary, secondary, outline)
│   │   │   ├── JsonLd.tsx            ← Structured data (Schema.org JSON-LD)
│   │   │   ├── MetricCard.tsx        ← Cartão de métrica (valor + label + ícone)
│   │   │   ├── SectionWrapper.tsx    ← Wrapper de seção (padding, id, fundo)
│   │   │   └── Skeleton.tsx          ← Loading skeleton placeholder
│   │   │
│   │   ├── layout/                 ← Componentes de layout (3)
│   │   │   ├── Header.tsx            ← Header com logo, nav desktop, mobile toggle
│   │   │   ├── Footer.tsx            ← Footer com links, social, copyright
│   │   │   └── MobileMenu.tsx        ← Menu mobile (drawer/overlay)
│   │   │
│   │   ├── home/                   ← Componentes da homepage (9)
│   │   │   ├── HeroSection.tsx       ← Hero com headline, subtítulo, CTAs
│   │   │   ├── CalculatorSection.tsx ← Seção wrapper das calculadoras
│   │   │   ├── ConsumerCalculator.tsx ← Calculadora de economia (consumidor)
│   │   │   ├── ProviderCalculator.tsx ← Calculadora de ganho (fornecedor)
│   │   │   ├── ProofSection.tsx      ← Prova social (métricas + depoimentos)
│   │   │   ├── HowItWorksSection.tsx ← Seção "Como Funciona" (3 passos)
│   │   │   ├── PlansSection.tsx      ← Seção de planos na homepage
│   │   │   ├── DifferentiatorsSection.tsx ← Diferenciais competitivos
│   │   │   └── FinalCtaSection.tsx   ← CTA final da homepage
│   │   │
│   │   ├── plans/                  ← Componentes da página de planos (4)
│   │   │   ├── PlanCard.tsx          ← Card de plano individual
│   │   │   ├── PlansComparisonTable.tsx ← Tabela comparativa de planos
│   │   │   ├── ProviderHighlight.tsx ← Highlight para fornecedores
│   │   │   └── FaqAccordion.tsx      ← Accordion FAQ
│   │   │
│   │   └── contact/                ← Componentes da página de contato (5)
│   │       ├── ContactForm.tsx       ← Formulário de contato principal
│   │       ├── FormField.tsx         ← Campo de formulário reutilizável
│   │       ├── JourneySummary.tsx    ← Resumo da jornada (contexto da calculadora)
│   │       ├── DirectChannels.tsx    ← Canais diretos (WhatsApp, email, redes)
│   │       └── SuccessScreen.tsx     ← Tela de sucesso pós-envio
│   │
│   ├── hooks/                      ← Custom hooks (7)
│   │   ├── useCalculator.ts         ← Lógica da calculadora (consumer/provider)
│   │   ├── useConcessionarias.ts    ← Fetch de concessionárias
│   │   ├── useConsumoMedio.ts       ← Fetch de consumo médio
│   │   ├── useContactForm.ts        ← Lógica do formulário de contato
│   │   ├── useFaqAccordion.ts       ← Estado de accordion FAQ
│   │   ├── useFAQs.ts              ← Fetch de FAQs
│   │   └── usePlans.ts             ← Fetch de planos
│   │
│   ├── lib/                        ← Lógica de negócio (5)
│   │   ├── types.ts                 ← Tipos core (Plan, ConsumerResult, LeadForm, etc.)
│   │   ├── constants.ts             ← Constantes de negócio (PLANS, RATES, FAQ_ITEMS, etc.)
│   │   ├── calculator.ts            ← Funções puras de cálculo (BRL format, economia, ganho)
│   │   ├── validation.ts            ← Validação de formulário (nome, email, telefone, etc.)
│   │   └── analytics.ts             ← Google Analytics 4 (eventos, tracking)
│   │
│   ├── services/                   ← Camada de dados assíncrona (4)
│   │   ├── servicePlans.ts          ← fetchPlans() — 500ms delay simulado
│   │   ├── serviceFAQs.ts          ← fetchFAQs() — 500ms delay simulado
│   │   ├── serviceConcessionarias.ts ← fetchConcessionarias() — 500ms delay simulado
│   │   └── serviceConsumoMedio.ts   ← fetchConsumoMedio() — 500ms delay simulado
│   │
│   ├── types/                      ← Tipos de dados externos (3)
│   │   ├── concessionaria.ts        ← Tipo Concessionaria
│   │   ├── consumo-medio.ts         ← Tipo ConsumoMedio
│   │   └── faq.ts                   ← Tipo FAQ
│   │
│   └── mocks/                      ← Dados mock JSON (4)
│       ├── mockPlans.json           ← 3 planos (Basic, Special, Premium)
│       ├── mockFAQs.json            ← 6 perguntas frequentes
│       ├── mockConcessionarias.json ← Lista de concessionárias
│       └── mockConsumoMedio.json    ← Dados de consumo médio por região
│
├── public/                         ← Assets estáticos
├── .specs/                         ← Documentação de especificação
├── docs/codebase/                  ← Esta documentação
├── .impeccable/                    ← Config do impeccable (auditoria de design)
├── .github/workflows/deploy.yml    ← CI/CD pipeline
├── next.config.ts                  ← Configuração Next.js
├── tsconfig.json                   ← Configuração TypeScript
├── postcss.config.mjs              ← Configuração PostCSS/Tailwind
├── eslint.config.mjs               ← Configuração ESLint
├── package.json                    ← Dependências e scripts
└── README.md                       ← Documentação inicial
```

## 2. Entry Points

| Entry Point | Arquivo | Propósito |
|---|---|---|
| **Root Layout** | `src/app/layout.tsx` | Metadata global, fontes (Inter via next/font), providers (Analytics) |
| **Homepage** | `src/app/page.tsx` | Landing page completa com 8 seções |
| **Planos** | `src/app/planos/page.tsx` | Página de planos standalone |
| **Contato** | `src/app/contato/page.tsx` | Página de contato standalone |
| **Robots** | `src/app/robots.ts` | Geração dinâmica de robots.txt |
| **Sitemap** | `src/app/sitemap.ts` | Geração dinâmica de sitemap.xml |

## 3. Rotas da Aplicação

| Rota | Página | Tipo | Descrição |
|---|---|---|---|
| `/` | `page.tsx` | Server Component | Homepage com Hero, Calculadora, Prova, Planos, CTA |
| `/planos` | `planos/page.tsx` | Server Component | Página de planos com tabela comparativa |
| `/contato` | `contato/page.tsx` | Client Component | Página de contato com formulário |

## 4. Arquivos de Configuração

| Arquivo | Propósito |
|---|---|
| `next.config.ts` | Config Next.js (output standalone, headers, imagens) |
| `tsconfig.json` | Config TypeScript (strict, ES2017, paths `@/*` → `src/*`) |
| `postcss.config.mjs` | PostCSS plugin Tailwind v4 |
| `eslint.config.mjs` | ESLint flat config (Next.js + TypeScript) |
| `.env.example` | Template de variáveis de ambiente |

## 5. Padrão de Organização

- **App Router**: Roteamento baseado em diretórios (`src/app/`), sem `pages/`
- **Componentes por domínio**: `shared/`, `layout/`, `home/`, `plans/`, `contact/`
- **Hooks isolados**: Um hook por responsabilidade, todos `'use client'`
- **Serviços com delay**: `services/` simulam API real com 500ms de delay → preparados para substituição por `fetch` real
- **Funções puras em `lib/`**: Lógica de cálculo e validação sem efeitos colaterais — testáveis isoladamente

## 6. Contagem de Arquivos

| Diretório | Arquivos | Tipo |
|---|---|---|
| `src/app/` | 6 | Páginas + layout + globals + robots + sitemap |
| `src/components/shared/` | 7 | Componentes reutilizáveis |
| `src/components/layout/` | 3 | Header, Footer, MobileMenu |
| `src/components/home/` | 9 | Seções da homepage |
| `src/components/plans/` | 4 | Componentes de planos |
| `src/components/contact/` | 5 | Componentes de contato |
| `src/hooks/` | 7 | Custom hooks |
| `src/lib/` | 5 | Lógica de negócio |
| `src/services/` | 4 | Camada de dados |
| `src/types/` | 3 | Tipos de dados |
| `src/mocks/` | 4 | Dados mock JSON |
| **Total** | **57** | Arquivos fonte TypeScript/TSX |

---

Última atualização: 2026-07-08
