# INTEGRATIONS — Solar Fácil Site

> APIs externas, serviços de terceiros, autenticação e monitoramento.
> Gerado por `acquire-codebase-knowledge` em 2026-07-08.
> Fontes: análise do código-fonte, `next.config.ts`, `.env.example`.

---

## 1. Visão Geral

A aplicação é majoritariamente **estática** — não consome APIs backend reais para dados de negócio. As únicas integrações externas são:

| Integração | Tipo | Status | Propósito |
|---|---|---|---|
| **Formspree** | API REST (POST) | ⚠️ Placeholder | Envio de formulário de contato |
| **Google Analytics 4** | Script client-side | ✅ Condicional | Analytics e tracking de eventos |
| **Google Fonts** | CDN (next/font) | ✅ Ativo | Fonte Inter (variável) |
| **Lucide React** | Biblioteca npm | ✅ Ativo | Ícones SVG |

## 2. Formspree (Formulário de Contato)

### 2.1 Configuração

```typescript
// src/hooks/useContactForm.ts:7
const FORM_ENDPOINT = process.env.NEXT_PUBLIC_FORM_ENDPOINT || 'https://formspree.io/f/placeholder';
```

### 2.2 Detalhes da Integração

| Propriedade | Valor |
|---|---|
| **Método** | POST |
| **Content-Type** | `multipart/form-data` (FormData) |
| **Endpoint** | Placeholder (`https://formspree.io/f/placeholder`) |
| **Timeout** | Sem timeout explícito |
| **Retry** | Não implementado |
| **Fallback** | Mensagem de erro genérica + sugestão de WhatsApp |

### 2.3 Campos Enviados

| Campo | Tipo | Descrição |
|---|---|---|
| `name` | string | Nome do contato |
| `email` | string | E-mail do contato |
| `phone` | string | Telefone (opcional) |
| `profile` | string | Perfil (consumidor/fornecedor/cooperativa) |
| `message` | string | Mensagem (opcional, máx 1000 chars) |
| `website` | string (hidden) | Honeypot anti-spam (deve estar vazio) |

### 2.4 Anti-Spam

- **Honeypot**: Campo hidden `website` — bots preenchem, humanos não
- **Tempo mínimo**: 3 segundos entre carregamento da página e submissão (bots submetem instantaneamente)
- **Comportamento**: Se detectado como bot, mostra fake success (não alerta o atacante)

## 3. Google Analytics 4

### 3.1 Configuração

```typescript
// src/lib/analytics.ts:11
export const GA_MEASUREMENT_ID = process.env.NEXT_PUBLIC_GA_ID || 'G-XXXXXXXXXX';
```

### 3.2 Eventos Definidos

| Evento | GA4 Event Name | Parâmetros | Status |
|---|---|---|---|
| CTA Click | `cta_click` | `cta_type`, `location` | ✅ Conectado |
| Uso da Calculadora | `calculator_use` | `persona`, `input_value`, `result`, `plan_suggested` | ✅ Conectado |
| FAQ Open | `faq_open` | `question_index` | ❌ Não conectado |
| Lead Capture | `lead_capture` | `persona`, `has_plan` | ❌ Não conectado |

### 3.3 Implementação

- **Provider**: `<AnalyticsProvider />` em `src/components/shared/AnalyticsProvider.tsx` — carrega script GA4 condicionalmente
- **Tracking**: Funções wrapper em `src/lib/analytics.ts` (tipadas, com verificação de `window.gtag`)
- **Condição**: Só dispara se `window.gtag` existe (script carregado)

## 4. Google Fonts (next/font)

### 4.1 Configuração

```typescript
// src/app/layout.tsx
import { Inter } from 'next/font/google';
const inter = Inter({ subsets: ['latin'] });
```

### 4.2 Detalhes

| Propriedade | Valor |
|---|---|
| **Fonte** | Inter (variável, sans-serif) |
| **Subsets** | `latin` |
| **Método** | `next/font/google` (otimização automática) |
| **Self-hosting** | Sim — Next.js faz download e serve localmente em produção |

## 5. Integrações NÃO Existentes (Planejadas/Necessárias)

| Integração | Motivo | Prioridade |
|---|---|---|
| **API Backend** | Dados de planos, concessionárias, FAQs são estáticos | Alta |
| **WhatsApp Business API** | Número placeholder `5511999999999`, link estático | Alta |
| **App Store / Google Play** | URLs placeholder para apps mobile | Média |
| **Sistema de Autenticação** | Sem login/cadastro atualmente | Baixa |
| **Email Marketing** | Sem integração com Mailchimp, RD Station, etc. | Baixa |
| **CSP Headers** | Sem Content-Security-Policy definido | Alta (segurança) |
| **HSTS** | Sem Strict-Transport-Security header | Alta (segurança) |
| **Sentry/LogRocket** | Sem monitoramento de erros client-side | Média |

## 6. Pacotes npm com Propósito Específico

| Pacote | Propósito | Criticalidade |
|---|---|---|
| `lucide-react` | Único pacote de runtime além de React/Next | Média (substituível) |
| `tailwindcss` + `@tailwindcss/postcss` | Estilização | Alta (estrutural) |
| `prettier` + `prettier-plugin-tailwindcss` | Formatação | Baixa (dev only) |
| `eslint` + `eslint-config-next` | Linting | Média (qualidade) |

## 7. Diagrama de Integrações

```
┌──────────────────────────────────────────┐
│           Solar Fácil Site               │
│                                           │
│  ┌─────────────┐  ┌──────────────────┐   │
│  │ Calculadora  │  │ Form. Contato    │   │
│  │ (client)     │  │ (client)         │   │
│  └──────┬───────┘  └────────┬─────────┘   │
│         │                   │              │
│         │ (funções puras)   │ (fetch POST) │
│         ▼                   ▼              │
│  ┌─────────────┐  ┌──────────────────┐   │
│  │ lib/        │  │ Formspree        │   │
│  │ calculator  │  │ (placeholder)    │   │
│  └─────────────┘  └──────────────────┘   │
│                                           │
│  ┌──────────────────────────────────┐    │
│  │ Google Analytics 4 (condicional) │    │
│  └──────────────────────────────────┘    │
│                                           │
│  ┌──────────────────────────────────┐    │
│  │ Google Fonts — Inter (next/font) │    │
│  └──────────────────────────────────┘    │
└──────────────────────────────────────────┘
```

---

Última atualização: 2026-07-08
