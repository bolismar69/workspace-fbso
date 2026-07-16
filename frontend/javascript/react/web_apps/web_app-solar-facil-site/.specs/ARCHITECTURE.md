# Arquitetura: Solar Fácil Site

> **Especificação de engenharia reversa** — todas as observações fundamentadas em análise do código-fonte.
> Gerado por `/spec-miner` em 2026-07-05. Revisado para pt-BR em 2026-07-06.
> Conteúdo integrado de `PRODUCT.md` (personas, propósito, personalidade da marca).

---

## 1. Contexto do Produto

> Fonte: `.specs/PRODUCT.md`

### Propósito

A **Solar Fácil** conecta produtores de energia solar excedente com consumidores que querem desconto na conta de luz, via cooperativas regulamentadas pela ANEEL. O site é a porta de entrada: educa sobre energia compartilhada, calcula a economia potencial, apresenta os planos e converte visitantes em leads qualificados.

**Critério de sucesso:** o visitante entender em 30 segundos que pode economizar sem investir nada, usar a calculadora, e iniciar contato.

### Personas

| Persona | Perfil | Necessidade |
|---------|--------|-------------|
| **Consumidores** | Pessoas físicas e pequenas empresas que querem reduzir a conta de luz sem instalar painéis solares. Não são especialistas em energia. Acessam principalmente via mobile, durante o horário comercial. | Economia sem complexidade |
| **Fornecedores/Produtores** | Quem já tem painéis solares e gera excedente. Quer rentabilizar esse excedente de forma simples e legalizada. Perfil mais técnico que o consumidor. | Rentabilizar excedente |
| **Cooperativas** | Parceiros B2B que operam a distribuição de energia. Precisam de credibilidade e informações sobre integração com a plataforma. | Integração e credibilidade |

### Personalidade da Marca

**Confiança · Sustentabilidade · Inovação · Amigável · Prático**

Uma marca que fala a língua das pessoas. Não é corporativa nem burocrática — é próxima, direta e transparente. Sustentabilidade não é bandeira verde; é consequência natural de um modelo de negócio inteligente.

**Tom de voz:** "A gente resolve." — Informativo sem ser professoral, entusiasmado sem ser vendedor, técnico sem ser incompreensível.

### Anti-referências (explicitamente rejeitadas)
- ❌ Startup tech fria (dark mode, neon, glassmorphism, estética "SaaS dashboard")
- ❌ Site governamental/estatal (bandeiras, brasões, azul imperial)
- ❌ Site genérico de energia (clichês verdes com folhas, gradientes ecológicos, ícones de gota e lâmpada)

### Princípios de Design

1. **"Energia que aproxima"** — Toda pessoa, independente de conhecimento técnico, entende o que fazemos em segundos.
2. **"Transparência radical"** — Números claros e verificáveis. A calculadora mostra exatamente a economia. Nada de letras miúdas.
3. **"Brasil real, sem clichês"** — Identidade brasileira sem apelar para verde-amarelo, bandeiras ou tropicalismo forçado.
4. **"Confiança ancorada em fatos"** — Prova social com números reais, chancela ANEEL visível, depoimentos verificáveis.
5. **"Simplicidade que respeita"** — A complexidade fica nos bastidores; a interface entrega clareza.

---

## 2. Stack Tecnológico

| Camada | Tecnologia | Versão | Evidência |
|--------|-----------|--------|-----------|
| **Framework** | Next.js (App Router) | 16.2.10 | `package.json:13` |
| **Biblioteca UI** | React | 19.2.4 | `package.json:14` |
| **Estilização** | Tailwind CSS v4 | ^4 | `package.json:26`, `postcss.config.mjs` |
| **Linguagem** | TypeScript | ^5 | `package.json:27`, `tsconfig.json` |
| **Ícones** | Lucide React | ^1.23.0 | `package.json:12` |
| **Linting** | ESLint 9 + `eslint-config-next` | ^9 / 16.2.10 | `eslint.config.mjs` |
| **Formatação** | Prettier + `prettier-plugin-tailwindcss` | ^3.9.4 / ^0.8.0 | `package.json:24-25` |
| **Target** | ES2017 | — | `tsconfig.json:3` |
| **Output** | Standalone (hospedagem não-Vercel) | — | `next.config.ts:5` |

### Design Tokens

Propriedades CSS customizadas com prefixo `--color-solar-*` definidas via Tailwind v4 `@theme inline` em `src/app/globals.css:8-36`.

### Decisões Arquiteturais Fundamentais

- **Família tipográfica única**: Inter atende todos os papéis tipográficos (display, corpo, label) — hierarquia via peso (400→500→700→800), não por troca de fontes
- **Propriedades CSS customizadas**: Todas as cores usam tokens `var(--color-solar-*)`, permitindo tematização futura
- **Output standalone**: `output: 'standalone'` em `next.config.ts:5` mira hospedagem não-Vercel (DigitalOcean, etc.)
- **Sem dark mode** (ainda): Comentário em `globals.css:38-40` indica planejamento para milestone futuro via `next-themes`
- **Sem dependências externas de runtime** além de React, Next e Lucide — zero bibliotecas de cliente API, sem biblioteca de gerenciamento de estado

---

## 3. Estrutura de Diretórios

```
solar-facil-site/
├── src/
│   ├── app/                          # Páginas Next.js App Router
│   │   ├── layout.tsx                # Layout raiz (metadata, Header, Footer, GA, JSON-LD)
│   │   ├── page.tsx                  # Homepage (7 seções compostas)
│   │   ├── globals.css               # Design tokens + estilos base + animações
│   │   ├── robots.ts                 # Geração de robots.txt
│   │   ├── sitemap.ts                # Geração de sitemap XML
│   │   ├── planos/
│   │   │   └── page.tsx             # /planos — tabela comparativa + FAQ
│   │   └── contato/
│   │       └── page.tsx             # /contato — formulário + canais diretos
│   │
│   ├── components/
│   │   ├── shared/                   # Primitivos reutilizáveis
│   │   │   ├── Button.tsx            # 3 variantes × 3 tamanhos + modo link
│   │   │   ├── SectionWrapper.tsx    # Container de seção consistente (fundo white/alt)
│   │   │   ├── MetricCard.tsx        # Exibição de estatística (modos padrão + invertido)
│   │   │   ├── Skeleton.tsx          # Skeleton de carregamento (4 variantes)
│   │   │   ├── Breadcrumb.tsx        # Navegação breadcrumb
│   │   │   ├── JsonLd.tsx            # Dados estruturados Schema.org Organization
│   │   │   └── AnalyticsProvider.tsx  # Injeção de script Google Analytics 4
│   │   │
│   │   ├── layout/                   # Componentes de casca
│   │   │   ├── Header.tsx            # Cabeçalho fixo (logo + nav + toggle mobile)
│   │   │   ├── Footer.tsx            # Rodapé 4 colunas (marca, nav, app, contato)
│   │   │   └── MobileMenu.tsx        # Navegação mobile full-screen (painel slide-in)
│   │   │
│   │   ├── home/                     # Seções da Homepage (7 seções)
│   │   │   ├── HeroSection.tsx       # Hero (headline + CTAs duplos + scroll indicator)
│   │   │   ├── CalculatorSection.tsx # Wrapper das calculadoras (2 abas: consumidor + fornecedor)
│   │   │   ├── ConsumerCalculator.tsx # Calculadora do consumidor — formulário + resultado
│   │   │   ├── ProviderCalculator.tsx # Calculadora do fornecedor — formulário + resultado
│   │   │   ├── ProofSection.tsx      # Métricas de prova social (boundary Suspense)
│   │   │   ├── HowItWorksSection.tsx # Timeline 3 passos (desktop horizontal, mobile vertical)
│   │   │   ├── PlansSection.tsx      # Cards de planos + destaque fornecedor
│   │   │   ├── DifferentiatorsSection.tsx # Layout assimétrico de diferenciais
│   │   │   └── FinalCtaSection.tsx   # Funil CTA inferior
│   │   │
│   │   ├── plans/                    # Componentes da página /planos
│   │   │   ├── PlanCard.tsx          # Card de plano (variantes compact + full)
│   │   │   ├── PlansComparisonTable.tsx # Tabela comparativa (desktop) + cards (mobile)
│   │   │   ├── ProviderHighlight.tsx # Card CTA para fornecedor
│   │   │   └── FaqAccordion.tsx      # Accordion de FAQ acessível
│   │   │
│   │   └── contact/                  # Componentes da página /contato
│   │       ├── ContactForm.tsx       # Formulário de lead (radio profile + campos + honeypot)
│   │       ├── FormField.tsx         # Campo genérico (input/textarea, modo oculto)
│   │       ├── SuccessScreen.tsx     # Tela de sucesso pós-envio
│   │       ├── JourneySummary.tsx    # Resumo contextual da calculadora (URL params)
│   │       └── DirectChannels.tsx    # Links WhatsApp + Email + Redes Sociais
│   │
│   ├── hooks/                        # Hooks React customizados (todos 'use client')
│   │   ├── useCalculator.ts          # Máquina de estado da calculadora + analytics
│   │   ├── useContactForm.ts         # Estado do form + validação + envio
│   │   ├── usePlans.ts              # Fetcher de planos { data, loading, error }
│   │   ├── useFAQs.ts              # Fetcher de FAQs { data, loading, error }
│   │   ├── useConsumoMedio.ts      # Fetcher de Consumo Médio
│   │   ├── useConcessionarias.ts    # Fetcher de Concessionárias
│   │   └── useFaqAccordion.ts       # Estado do accordion + limiar de dica de contato
│   │
│   ├── lib/                          # Lógica de negócio pura (sem React)
│   │   ├── types.ts                 # Tipos core do domínio (Plan, LeadForm, resultados da calculadora)
│   │   ├── constants.ts             # Regras de negócio, planos, métricas, FAQs, URLs
│   │   ├── calculator.ts            # Funções puras de cálculo (economia consumidor, ganho fornecedor)
│   │   ├── validation.ts            # Validação de formulário (nome, email, telefone, perfil, mensagem)
│   │   └── analytics.ts             # Helpers de eventos GA4 (eventos tipados)
│   │
│   ├── services/                     # Busca de dados assíncrona (simulada — delay 500ms)
│   │   ├── servicePlans.ts          # → constante PLANS
│   │   ├── serviceFAQs.ts          # → constante FAQ_ITEMS
│   │   ├── serviceConsumoMedio.ts   # → mockConsumoMedio.json
│   │   └── serviceConcessionarias.ts # → mockConcessionarias.json
│   │
│   ├── types/                        # Tipos de entidade de domínio (mapeados do App)
│   │   ├── concessionaria.ts        # ConcessionariaType
│   │   ├── faq.ts                   # FAQType, FAQCategoryType
│   │   └── consumo-medio.ts         # ConsumoMedioType
│   │
│   └── mocks/                        # Conjuntos de dados JSON estáticos
│       ├── mockPlans.json            # 3 planos (dados de id/nome/preço/indicadores comerciais)
│       ├── mockFAQs.json             # 3 categorias × 3 FAQs cada
│       ├── mockConcessionarias.json  # 50 concessionárias (7 ativas, 43 inativas)
│       └── mockConsumoMedio.json     # 6 faixas de consumo (0–500+ kWh)
│
├── public/
│   ├── logo.svg                      # Logo da marca
│   ├── globe.svg, window.svg, file.svg, next.svg, vercel.svg  # Assets padrão Next
│
├── .env.example                      # GA4 ID, endpoint do formulário, URL do site
├── next.config.ts                    # Output standalone, headers de segurança, formatos de imagem
├── tsconfig.json                     # ES2017, strict, resolução bundler, path aliases
├── eslint.config.mjs                 # ESLint 9 flat config (core-web-vitals + typescript)
├── postcss.config.mjs                # Plugin PostCSS do Tailwind CSS v4
└── package.json                      # Scripts: dev, build, start, lint
```

---

## 4. Fluxo de Dados

### Fluxo de Renderização da Página

```
Requisição do Navegador
    ↓
Next.js App Router → layout.tsx
    ├── Metadata (SEO: OG, robots, JSON-LD)
    ├── <html lang="pt-BR"> com variável de fonte Inter
    ├── Link "Pular para o conteúdo" (acessibilidade)
    ├── JsonLd (Schema.org Organization)
    ├── AnalyticsProvider (injeção de script GA4)
    ├── Header (fixo, backdrop-blur)
    └── <main id="main-content">
        └── page.tsx → HomePage (7 seções em ordem)
            ├── HeroSection (client)
            ├── CalculatorSection (server) → ConsumerCalculator + ProviderCalculator (ambos client)
            ├── ProofSection (client, boundary Suspense)
            ├── HowItWorksSection (server)
            ├── PlansSection (server)
            ├── DifferentiatorsSection (server)
            └── FinalCtaSection (client)
    └── Footer (server)
```

### Fluxo da Calculadora (Exemplo Consumidor)

```
Usuário digita valor da conta mensal (R$)
    ↓
useCalculator('consumer')
    ├── setInput(valor) → atualiza estado
    └── calculate()
        ├── Converte input para float
        ├── validateConsumerInput(conta)
        │   ├── conta < 50  → outlier (mensagem + sem plano)
        │   ├── conta > 5000 → outlier (mensagem + sem plano)
        │   └── Faixa válida → calculateConsumerEconomy(conta)
        │       ├── economia = Math.round(conta × 0.12 × 100) / 100
        │       └── suggestPlan(conta)
        │           ├── 100–200 kWh → Basic
        │           ├── 200–350 kWh → Special
        │           └── 350–600 kWh → Premium
        └── trackCalculatorUse({ persona, input, result, plan })
            → window.gtag('event', 'calculator_use', ...)
```

### Fluxo do Formulário de Contato

```
Usuário preenche formulário → envia
    ↓
useContactForm.handleSubmit(e)
    ├── preventDefault()
    ├── validateForm(values)
    │   ├── validateName(nome)       → mín 2 caracteres
    │   ├── validateEmail(email)     → regex /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    │   ├── validatePhone(telefone)  → opcional, 10-11 dígitos
    │   ├── validateProfile()        → obrigatório
    │   └── validateMessage()        → opcional, máx 1000 caracteres
    ├── hasErrors()? → retorna (exibe erros nos campos)
    ├── Anti-spam: elapsed < 3000ms? → falso sucesso (bloqueio silencioso)
    ├── POST para NEXT_PUBLIC_FORM_ENDPOINT (FormData)
    │   ├── Campo honeypot deve estar vazio
    │   ├── response.ok → setIsSuccess(true)
    │   └── !response.ok → setSubmitError(mensagem + fallback WhatsApp)
```

### Padrão da Camada de Serviços

Todos os serviços seguem o mesmo padrão assíncrono — simulando chamadas de backend com delay de 500ms:

```
fetch[Recurso]()
    └── new Promise((resolve) => {
            setTimeout(() => resolve(dadosEstaticos), 500)
        })
```

Isso significa: todos os dados são estáticos e empacotados em tempo de build. NÃO há chamadas de API reais, NÃO há banco de dados, NÃO há autenticação.

---

## 5. Mapa de Rotas

| Rota | Página | Título Metadata | Componentes Principais |
|------|--------|-----------------|----------------------|
| `/` | `src/app/page.tsx` | "Solar Fácil — Energia Limpa Compartilhada" | Hero, Calculadora (×2), Prova, ComoFunciona, Planos, Diferenciais, CtaFinal |
| `/planos` | `src/app/planos/page.tsx` | "Planos — Solar Fácil" | Breadcrumb, TabelaComparativa, DestaqueFornecedor, FaqAccordion |
| `/contato` | `src/app/contato/page.tsx` | "Contato — Solar Fácil" | Breadcrumb, ResumoJornada, FormulárioContato, CanaisDiretos |
| `/robots.txt` | `src/app/robots.ts` | — | Gerado: allow all, URL do sitemap |
| `/sitemap.xml` | `src/app/sitemap.ts` | — | Gerado: 3 URLs (/, /planos, /contato) |

### Separação Componentes Cliente/Servidor

| Tipo | Componentes |
|------|-------------|
| **Server** (RSC) | `layout.tsx`, `page.tsx`, `SectionWrapper`, `PlansSection`, `HowItWorksSection`, `DifferentiatorsSection`, `PlanCard`, `PlansComparisonTable`, `ProviderHighlight`, `Breadcrumb`, `Footer`, `JsonLd`, `AnalyticsProvider`, `Skeleton` |
| **Client** (`'use client'`) | `Header`, `MobileMenu`, `HeroSection`, `ConsumerCalculator`, `ProviderCalculator`, `ProofSection`, `FinalCtaSection`, `FaqAccordion`, `ContactForm`, `FormField`, `JourneySummary`, `SuccessScreen`, `DirectChannels` |

**Justificativa**: Componentes que precisam de APIs do navegador (`useState`, `useSearchParams`, `usePathname`, scroll, window) são client components. Componentes puramente apresentacionais permanecem server components para máxima performance.

---

## 6. Configuração de Build

### Next.js Config (`next.config.ts`)

```typescript
{
  output: 'standalone',          // Hospedagem não-Vercel (DigitalOcean, etc.)
  images: {
    formats: ['image/avif', 'image/webp'],  // Otimização moderna de imagens
  },
  headers: [
    { source: '/(.*)',
      headers: [
        'X-Frame-Options: DENY',                           // Anti-clickjacking
        'X-Content-Type-Options: nosniff',                 // Proteção MIME sniffing
        'Referrer-Policy: strict-origin-when-cross-origin', // Privacidade
      ]
    }
  ]
}
```

### TypeScript Config (`tsconfig.json`)

- Target: ES2017
- Modo strict: `true`
- Módulo: ESNext com resolução bundler
- Path alias: `@/*` → `./src/*`
- Compilação incremental habilitada

---

## 7. Integrações Externas

| Integração | Propósito | Status |
|------------|-----------|--------|
| **Google Analytics 4** | Page views + 4 eventos customizados | Condicional (apenas se `NEXT_PUBLIC_GA_ID` configurado, `AnalyticsProvider.tsx:4`) |
| **Formspree** (ou similar) | Envio do formulário de contato | Configurável via `NEXT_PUBLIC_FORM_ENDPOINT` |
| **WhatsApp** | Canal de contato alternativo | Número fixo em constantes (`WHATSAPP_NUMBER`, `constants.ts:130`) |
| **Google Fonts** | Tipo de fonte Inter | Via `next/font/google` (`layout.tsx:9`) |

> **Nota**: GA4 é a única integração ativa. O envio de formulário e WhatsApp estão conectados mas usam valores placeholder. Não existe backend de API real — todos os dados são estáticos.

---

## 8. Padrões e Convenções

### Convenções de Nomenclatura
- **Componentes**: PascalCase, declarações de função com `export default`
- **Hooks**: prefixo `use`, um hook por arquivo
- **Serviços**: prefixo `service`, funções async com delay simulado de 500ms
- **Tipos**: Interfaces PascalCase, tipos de entidade de domínio com sufixo `Type`
- **Classes CSS**: Classes utilitárias do Tailwind, sem CSS modules

### Organização do Código
- Lógica de negócio pura em `lib/` (sem imports do React)
- Acesso a dados assíncrono em `services/` (backend simulado)
- Gerenciamento de estado em `hooks/` (apenas estado React, sem store global)
- Tipos de entidade de domínio em `types/` (mapeados do modelo de domínio do App)
- Dados de teste estáticos em `mocks/` (JSON)

### Tratamento de Erros
- Serviços: try/catch com fallback para array vazio + console.error
- Formulários: erros de validação por campo retornados como objeto `FormErrors`
- Calculadora: faixas outlier capturadas antes do cálculo, retorna `isOutlier: true` + mensagem
- Rede: bloco catch do formulário de contato exibe fallback do WhatsApp

### Padrões de Acessibilidade
- Link "Pular para o conteúdo" em `layout.tsx:46-49`
- `aria-expanded` + `aria-controls` no accordion de FAQ (`FaqAccordion.tsx:27-28`)
- `aria-label` no toggle do menu mobile (`Header.tsx:56`)
- `aria-label="Breadcrumb"` na navegação breadcrumb (`Breadcrumb.tsx:10`)
- `role="status"` nos resultados da calculadora (`ConsumerCalculator.tsx:67`)
- `aria-hidden="true"` em elementos decorativos (skeletons, scroll indicator)
- `prefers-reduced-motion` respeitado globalmente (`globals.css:73-82`)
- `text-wrap: balance` em todas as headlines para balanceamento de linhas
- `text-wrap: pretty` em parágrafos longos para prevenção de viúvas/órfãs

### Requisitos de Acessibilidade e Inclusão
> Fonte: `.specs/PRODUCT.md`

- **WCAG AA** como padrão mínimo em todas as páginas
- Contraste ≥ 4.5:1 para texto corrido, ≥ 3:1 para texto grande
- Navegação completa por teclado (focus rings visíveis, tab order lógica)
- Suporte a leitores de tela (ARIA labels, landmarks semânticas, alt text descritivo)
- Testes com os 3 tipos mais comuns de daltonismo (deuteranopia, protanopia, tritanopia)
- **Idiomas:** pt-BR (primário) com arquitetura preparada para espanhol (futuro — rotas com prefixo `/es/`)
