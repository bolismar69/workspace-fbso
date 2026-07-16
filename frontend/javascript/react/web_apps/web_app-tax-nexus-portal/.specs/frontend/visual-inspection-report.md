---
title: "Inspeção Visual — TaxNexus Portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["visual-inspection", "layout", "responsive", "contrast"]
inspection_method: "static-code-analysis"
browser_available: false
screenshots: false
---

# Inspeção Visual — TaxNexus Portal

**⚠️ Método degradado:** Análise estática de classes CSS e estrutura de componentes. Screenshots e verificação de viewport indisponíveis (browser não pôde ser instalado neste ambiente). Issues reportados como prováveis baseados nas classes Tailwind e CSS custom properties encontradas.

---

## Estrutura de Layout

A aplicação tem 2 views principais:

### View 1: Login (CNPJ)

```
┌──────────────────────────────────────┐
│  .min-h-screen .bg-gray-100 .py-10   │
│  ┌────────────────────────────────┐  │
│  │  .max-w-md .bg-white .p-8      │  │
│  │  .rounded-lg .shadow-md        │  │
│  │                                │  │
│  │  TaxNexus TaaS (h1)            │  │
│  │  CNPJ do Contribuinte (label)  │  │
│  │  [00000000000000] (input)      │  │
│  │  [ ACESSAR PORTAL ] (button)   │  │
│  └────────────────────────────────┘  │
└──────────────────────────────────────┘
```

**Classes Tailwind analisadas:**
- Container externo: `min-h-screen bg-gray-100 py-10` — fundo cinza claro, padding vertical generoso
- Card: `max-w-md mx-auto bg-white p-8 rounded-lg shadow-md` — 448px max, centralizado, sombra + border-radius
- Título: `text-2xl font-bold mb-6 text-center text-blue-900` — 24px bold, azul escuro
- Input: `mt-1 block w-full p-3 border rounded-md shadow-sm` — full-width, padding 12px, borda + shadow-sm (⚠️ ghost-card)
- Botão: `w-full bg-blue-900 text-white py-3 rounded-md font-semibold hover:bg-blue-800` — altura 48px (bom para touch)

### View 2: Simulador + Resultados

```
┌──────────────────────────────────────────────┐
│  .container .mx-auto                          │
│  ┌──────────────────────────────────────────┐ │
│  │ Header: flex justify-between              │ │
│  │ TaxNexus Simulator v1.0  |  Sair/CNPJ    │ │
│  └──────────────────────────────────────────┘ │
│  ┌──────────────────────────────────────────┐ │
│  │ .max-w-4xl .bg-white .rounded-lg         │ │
│  │ .shadow-lg .p-6 .space-y-6               │ │
│  │                                          │ │
│  │ CNPJ: XX.XXX.XXX/XXXX-XX (info bar)      │ │
│  │                                          │ │
│  │ [Estado ▽] [Cidade ▽] [NCM: 62011100]    │ │
│  │ ┌──────────────────────────────────────┐ │ │
│  │ │ Saldo Remanescente (R$) [________]   │ │ │
│  │ └──────────────────────────────────────┘ │ │
│  │ [ SIMULAR REFORMA TRIBUTÁRIA ]          │ │
│  │                                          │ │
│  │ ┌──────────┐ ┌──────────┐               │ │
│  │ │ Legado   │ │ Reforma  │               │ │
│  │ │ PIS: R$…│ │ CBS: R$… │               │ │
│  │ │ COFINS… │ │ IBS: R$… │               │ │
│  │ └──────────┘ └──────────┘               │ │
│  │ ┌──────────────────────────────────────┐ │ │
│  │ │     BarChart: CBS ▓▓ + IBS ▓▓        │ │ │
│  │ └──────────────────────────────────────┘ │ │
│  └──────────────────────────────────────────┘ │
└──────────────────────────────────────────────┘
```

---

## Análise de Contraste (Cores → Background)

| Elemento | Cor Texto | Cor Fundo | Contraste Estimado | WCAG AA |
|---|---|---|---|---|
| Body text (`--text`) | #6b6375 | #ffffff (`--bg`) | ~4.6:1 | ✅ AA |
| Heading (`--text-h`) | #08060d | #ffffff | ~18.1:1 | ✅ AAA |
| Botão primary | #ffffff | #1e3a8a (blue-900) | ~12.5:1 | ✅ AAA |
| Botão primary hover | #ffffff | #1e40af (blue-800) | ~10.8:1 | ✅ AAA |
| Botão disabled | #ffffff? | #9ca3af (gray-400) | ~2.9:1 | ❌ FALHA |
| Label CNPJ (gray-700) | #374151 | #ffffff | ~5.9:1 | ✅ AA |
| Card legado título | #9a3412 (orange-800) | #fff7ed (orange-50) | ~5.8:1 | ✅ AA |
| Card reforma título | #1e40af (blue-800) | #eff6ff (blue-50) | ~7.5:1 | ✅ AAA |
| Saldo label (yellow-800) | #854d0e | #fefce8 (yellow-50) | ~6.3:1 | ✅ AA |
| Link "Sair" (red-600) | #dc2626 | #ffffff (bg header) | ~4.0:1 | ⚠️ AA large text only |
| Texto escuro (`--text`) | #9ca3af | #16171d (`--bg-dark`) | ~8.5:1 | ✅ AAA |

**⚠️ Issue de contraste:** O botão em estado `disabled` (`bg-gray-400` com texto branco) falha WCAG AA (2.9:1). Recomendação: usar `text-gray-500` no estado disabled para melhorar contraste.

---

## Análise de Responsividade

### Breakpoints Detectados

| Breakpoint | Tailwind Class | Comportamento |
|---|---|---|
| Default (mobile) | — | Layout single-column, padding reduzido via `index.css` |
| 768px | `md:` | Grid de selects 3 colunas; cards de resultado 2 colunas |
| 1024px | `@media (max-width: 1024px)` em `index.css` | Font-size body reduz de 18px → 16px; headings reduzem |

### Issues de Responsividade Prováveis

| ID | Severidade | Issue | Localização | Viewports Afetados |
|---|---|---|---|---|
| V1 | **P1** | Input CNPJ full-width (`block w-full`) ok, mas sem `max-width` — em desktop wide (1920px) pode ficar exageradamente largo | `App.tsx:26` | ≥1280px |
| V2 | **P1** | Grid de resultados usa `md:grid-cols-2` — em mobile (<768px) cards ficam empilhados, mas sem padding lateral adequado | `TaxSimulator.tsx:105` | <768px |
| V3 | **P2** | Gráfico Recharts com `h-80` (320px) fixo — em mobile muito alto, ocupa a tela inteira | `TaxSimulator.tsx:131` | <640px |
| V4 | **P2** | Selects em grid `md:grid-cols-3` mas sem `sm:grid-cols-1` explícito — confia no comportamento padrão de grid | `TaxSimulator.tsx:65` | <768px |
| V5 | **P3** | Container `max-w-4xl` (896px) — em telas 4K o conteúdo fica estreito com muito espaço vazio nas laterais | `TaxSimulator.tsx:58` | ≥1920px |
| V6 | **P3** | `#root` com `width: 1126px` fixo no `index.css` contradiz os containers responsivos Tailwind dos componentes | `index.css:54` | Todos |

**⚠️ Issue crítico de layout (V6):** O `#root` tem `width: 1126px` e `max-width: 100%`. Isso é um resquício do template Vite que limita a largura máxima. Os componentes usam classes Tailwind responsivas, mas o container root impõe um teto de 1126px que pode conflitar com o comportamento esperado em telas ultra-wide.

---

## Consistência Visual

### Fontes
- ❌ **Inconsistência de font-family:** O `index.css` define `system-ui, 'Segoe UI', Roboto, sans-serif`. Os componentes React usam classes Tailwind que referenciariam a font-family padrão do Tailwind (Inter), mas como Tailwind não está instalado, a fonte cai para o default do navegador. Isso pode resultar em fontes diferentes entre o texto base (system-ui via `index.css`) e os componentes (browser default)
- ✅ Pesos consistentes: `font-bold` (700) para headings, `font-semibold` (600) para botões, `font-medium` (500) para labels

### Cores
- ✅ Sistema semântico consistente: laranja para legado, azul para reforma, amarelo para alertas
- ❌ Cor de acento roxa (`#aa3bff`) definida no `index.css` nunca é usada nos componentes — token órfão
- ❌ Azul da marca inconsistente: `blue-900` (#1e3a8a) no título de login vs. `blue-600` (#2563eb) no botão de simular — duas hierarquias diferentes de azul sem relação clara

### Espaçamento
- ✅ Escala consistente: `space-y-4` (16px), `space-y-6` (24px), `p-6` (24px), `p-8` (32px), `gap-4` (16px), `gap-6` (24px)
- ❌ Mix de `px-6` com `container mx-auto` — padding lateral inconsistente entre header e conteúdo

---

## Estados da Aplicação

| Estado | Cobertura | Issue |
|---|---|---|
| **Default (login)** | ✅ | Formulário renderiza corretamente |
| **Default (simulador)** | ✅ | Todos os campos e botão visíveis |
| **Loading** | ✅ | Botão mostra "Processando..." + fica disabled |
| **Resultado** | ✅ | Cards + gráfico renderizam condicionalmente |
| **Empty (sem estado)** | ✅ | Select de cidade começa vazio e disabled |
| **Error (API)** | ❌ | Sem tratamento visível — botão volta ao normal silenciosamente |
| **CNPJ inválido** | ⚠️ | `alert()` — funcional mas pobre |
| **Sem cidade selecionada** | ✅ | Botão disabled até cidade escolhida |

---

## Sumário

| Severidade | Quantidade | Issues |
|---|---|---|
| P1 | 2 | V1 (input largo em desktop), V2 (cards empilhados sem padding mobile) |
| P2 | 3 | V3 (gráfico altura fixa), V4 (grid sem fallback), Contraste disabled button |
| P3 | 3 | V5 (max-w estreito em 4K), V6 (width fixo no root), Token roxo órfão |

---

🤖 *Inspeção visual gerada por mineração de especificações frontend (PROMPT-MINING-FRONTEND-SPECIFICATION). Screenshots indisponíveis — browser não pôde ser instalado neste ambiente.*
