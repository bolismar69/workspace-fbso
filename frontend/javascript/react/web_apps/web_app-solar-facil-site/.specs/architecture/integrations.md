# Integrações — Solar Fácil Site

> Mapeamento completo de dependências externas, APIs e bibliotecas críticas.
> Gerado por `architecture-designer` em 2026-07-08.

---

## 1. Visão Geral

A aplicação é predominantemente **estática** — não possui backend próprio. As integrações externas são mínimas e majoritariamente client-side.

## 2. Integrações Externas

### 2.1 Formspree (Formulário de Contato)

| Propriedade | Valor |
|---|---|
| **Tipo** | API REST (POST) |
| **Endpoint** | `https://formspree.io/f/{form-id}` |
| **Status** | ⚠️ Placeholder (`formspree.io/f/placeholder`) |
| **Método** | POST |
| **Content-Type** | `multipart/form-data` |
| **Autenticação** | Nenhuma (form-id público) |
| **Timeout** | Padrão do navegador (sem configuração explícita) |
| **Retry** | Não implementado |
| **Fallback** | Mensagem de erro + sugestão de contato via WhatsApp |

### 2.2 Google Analytics 4

| Propriedade | Valor |
|---|---|
| **Tipo** | Script client-side + API de eventos |
| **Measurement ID** | `NEXT_PUBLIC_GA_ID` (env var, default `G-XXXXXXXXXX`) |
| **Status** | ✅ Condicional — carrega apenas se ID configurado |
| **Método** | `window.gtag('event', ...)` |
| **Eventos** | `cta_click`, `calculator_use`, `faq_open`, `lead_capture` |
| **Condição** | `AnalyticsProvider` verifica `NEXT_PUBLIC_GA_ID` antes de injetar script |

### 2.3 Google Fonts (next/font)

| Propriedade | Valor |
|---|---|
| **Tipo** | Download build-time |
| **Fonte** | Inter (variável, sans-serif) |
| **Subsets** | `latin` |
| **Método** | `next/font/google` — download no build, self-hosted em produção |
| **Impacto de rede** | Zero em produção (fontes incluídas no bundle) |

## 3. Dependências npm Críticas

### 3.1 Runtime

| Pacote | Versão | Propósito | Criticalidade |
|---|---|---|---|
| `next` | 16.2.10 | Framework SSR/SSG | 🔴 Estrutural |
| `react` | 19.2.4 | Biblioteca UI | 🔴 Estrutural |
| `react-dom` | 19.2.4 | Renderizador DOM | 🔴 Estrutural |
| `lucide-react` | ^1.23.0 | Ícones SVG | 🟡 Substituível |

### 3.2 Desenvolvimento

| Pacote | Propósito | Criticalidade |
|---|---|---|
| `tailwindcss` + `@tailwindcss/postcss` | Framework CSS | 🔴 Estrutural |
| `typescript` | Type-checking | 🔴 Estrutural |
| `eslint` + `eslint-config-next` | Linting | 🟡 Qualidade |
| `prettier` + `prettier-plugin-tailwindcss` | Formatação | 🟢 Conveniência |

## 4. Serviços que Simulam API (Camada de Dados)

Todos os serviços em `src/services/` seguem o mesmo padrão:

```typescript
export async function fetchX(): Promise<X[]> {
  return new Promise((resolve) => {
    try {
      setTimeout(() => resolve(DATA), 500); // Simula latência de API
    } catch (error) {
      console.error("Erro => ", error);
      resolve([]); // Fallback seguro
    }
  });
}
```

| Serviço | Fonte de Dados | Delay | Quando migrar para API real |
|---|---|---|---|
| `servicePlans` | `PLANS` (constants.ts) | 500ms | Quando backend de planos existir |
| `serviceFAQs` | `FAQ_ITEMS` (constants.ts) | 500ms | Quando backend de FAQs existir |
| `serviceConcessionarias` | `mockConcessionarias.json` | 500ms | Quando backend de concessionárias existir |
| `serviceConsumoMedio` | `mockConsumoMedio.json` | 500ms | Quando backend de consumo médio existir |

## 5. Diagrama de Dependências

```
┌──────────────────────────────────────────┐
│              Solar Fácil Site             │
│                                           │
│  Dependências de Runtime:                 │
│  ├── next (16.2.10)                      │
│  ├── react (19.2.4)                      │
│  ├── react-dom (19.2.4)                  │
│  └── lucide-react (^1.23.0)              │
│                                           │
│  Dependências Externas:                   │
│  ├── Formspree (formulário, placeholder)  │
│  ├── Google Analytics 4 (condicional)     │
│  └── Google Fonts (build-time)            │
│                                           │
│  Dados (estáticos):                       │
│  ├── constants.ts (PLANS, FAQ_ITEMS, ...) │
│  └── mocks/*.json (planos, FAQs, ...)     │
└──────────────────────────────────────────┘
```

## 6. Variáveis de Ambiente

| Variável | Default | Uso |
|---|---|---|
| `NEXT_PUBLIC_GA_ID` | `G-XXXXXXXXXX` | Google Analytics 4 Measurement ID |
| `NEXT_PUBLIC_FORM_ENDPOINT` | `https://formspree.io/f/placeholder` | Endpoint do formulário de contato |

---

Última atualização: 2026-07-08
