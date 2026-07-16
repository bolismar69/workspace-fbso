---
title: "Design Tokens — TaxNexus Portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
source: "src/index.css (custom properties) + Tailwind utility classes in components"
---

# Design Tokens — TaxNexus Portal

Tokens extraídos do código-fonte. O sistema usa uma combinação de CSS custom properties (herdadas do template Vite) e classes utilitárias Tailwind-like nos componentes React. **Tailwind não está listado como dependência no `package.json`** — as classes funcionam como convenção visual, não como motor de estilização ativo.

---

## Cores

### Tema Claro (default)

| Token | Valor | Uso |
|---|---|---|
| `--text` | `#6b6375` | Texto de corpo — roxo-acinzentado suave |
| `--text-h` | `#08060d` | Títulos e headings — quase preto |
| `--bg` | `#ffffff` | Fundo principal — branco puro |
| `--border` | `#e5e4e7` | Bordas e divisores — cinza claro |
| `--code-bg` | `#f4f3ec` | Fundo de elementos `<code>` — creme quente |
| `--accent` | `#aa3bff` | Cor de destaque (hover, focus) — roxo vibrante |
| `--accent-bg` | `rgba(170, 59, 255, 0.1)` | Fundo de destaque sutil |
| `--accent-border` | `rgba(170, 59, 255, 0.5)` | Borda de destaque |
| `--social-bg` | `rgba(244, 243, 236, 0.5)` | Fundo de elementos sociais |
| `--shadow` | `0 10px 15px -3px rgba(0,0,0,0.1), 0 4px 6px -2px rgba(0,0,0,0.05)` | Sombra de cards |

### Tema Escuro (`prefers-color-scheme: dark`)

| Token | Valor |
|---|---|
| `--text` | `#9ca3af` |
| `--text-h` | `#f3f4f6` |
| `--bg` | `#16171d` |
| `--border` | `#2e303a` |
| `--code-bg` | `#1f2028` |
| `--accent` | `#c084fc` |
| `--accent-bg` | `rgba(192, 132, 252, 0.15)` |
| `--shadow` | `0 10px 15px -3px rgba(0,0,0,0.4), 0 4px 6px -2px rgba(0,0,0,0.25)` |

### Cores Semânticas (Tailwind — usadas nos componentes)

| Cor | Tailwind Class | Uso |
|---|---|---|
| Azul Marca | `blue-900` (#1e3a8a) | Título "TaxNexus TaaS", header |
| Azul Ação | `blue-600` (#2563eb) | Botão "SIMULAR", hover para `blue-700` |
| Laranja Legado | `orange-50` (#fff7ed) | Fundo de card do sistema atual |
| Laranja Legado Borda | `orange-200` (#fed7aa) | Borda de card do sistema atual |
| Azul Reforma | `blue-50` (#eff6ff) | Fundo de card do novo sistema |
| Azul Reforma Borda | `blue-200` (#dbeafe) | Borda de card do novo sistema |
| Amarelo Alerta | `yellow-50` (#fefce8) | Fundo da seção de saldo |
| Amarelo Alerta Borda | `yellow-300` (#fde047) | Borda do input de saldo |
| Cinza Fundo Página | `gray-100` (#f3f4f6) | Background da página |
| Cinza Texto Label | `gray-700` (#374151) | Labels de formulário |
| Cinza Disabled | `gray-400` (#9ca3af) | Botão disabled |
| Vermelho Ação | `red-600` (#dc2626) | Link "Sair / Trocar CNPJ" |
| Roxo Destaque | `purple-700` (#7e22ce) | Texto "Novos IPVA/ITCMD" |
| Verde IBS | `green-600` | Cor da barra IBS no gráfico |

---

## Tipografia

### Font Families

| Token | Valor | Aplicação |
|---|---|---|
| `--sans` | `system-ui, 'Segoe UI', Roboto, sans-serif` | Corpo e UI |
| `--heading` | `system-ui, 'Segoe UI', Roboto, sans-serif` | Títulos (mesma stack) |
| `--mono` | `ui-monospace, Consolas, monospace` | Código, CNPJ, NCM |

### Escala Tipográfica

| Nível | Font Size | Weight | Line Height | Letter Spacing | Uso |
|---|---|---|---|---|---|
| H1 | 56px (36px ≤1024px) | 500 | normal | -1.68px | Título da página |
| H2 | 24px (20px ≤1024px) | 500 | 118% | -0.24px | Títulos de seção |
| Body | 18px (16px ≤1024px) | 400 | 145% | 0.18px | Texto corrido |
| Code | 15px | 400 | 135% | normal | Código inline |
| `.counter` | 16px | 400 | normal | normal | Contador/badge |

### Tailwind Font Sizes (usadas nos componentes)

| Classe | Tamanho | Uso |
|---|---|---|
| `text-2xl` | 24px | Título "TaxNexus TaaS" |
| `text-xl` | 20px | Título "TaxNexus Simulator v1.0" |
| `text-sm` | 14px | Labels, texto secundário, CNPJ no header |
| `text-xs` | 12px | Não usado atualmente |

---

## Espaçamento

### Escala Implícita (valores extraídos das classes)

| Nome | Valor | Classes Representativas |
|---|---|---|
| xs | 4px | `p-1`, `gap-1`, `space-y-1` |
| sm | 8px | `p-2`, `gap-2`, `mt-2`, `mb-2` |
| md | 12px | `p-3`, `py-3` (botões) |
| lg | 16px | `p-4`, `gap-4`, `mb-4`, `space-y-4` |
| xl | 20px | `p-5` (cards de resultado) |
| 2xl | 24px | `p-6`, `mb-6`, `space-y-6` |
| 3xl | 32px | `p-8` (card de login) |
| 4xl | 40px | `py-10` (container da página) |

---

## Border Radius

| Nome | Valor | Uso |
|---|---|---|
| `sm` | 4px | Elementos `<code>`, badges |
| `md` | 6px | Inputs, links sociais |
| `lg` | 8px | Cards, botões (`rounded-lg`) |
| `xl` | 12px | Cards de resultado (`rounded-xl`) |
| `full` | 9999px | Não usado atualmente |

---

## Sombras

| Nome | Valor | Uso |
|---|---|---|
| `md` | `0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)` | Card de login |
| `lg` | `0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)` | Container do simulador |
| `--shadow` | `0 10px 15px -3px rgba(0,0,0,0.1), 0 4px 6px -2px rgba(0,0,0,0.05)` | Custom property (sobrescreve Tailwind) |

---

## Breakpoints

| Nome | Largura | CSS/Classe |
|---|---|---|
| Tablet | ≤1024px | `@media (max-width: 1024px)` em `index.css` |
| Mobile | <768px | Classes `md:` do Tailwind |
| Desktop | ≥1280px | Comportamento default |

---

## Transições & Motion

| Elemento | Transição |
|---|---|
| `.counter` hover | `border-color 0.3s` |
| Links sociais hover | `box-shadow 0.3s` |
| Botão primary hover | `background-color` (instantâneo, sem transição explícita) |

O sistema é majoritariamente sem animações. A única animação declarada é `animate-in slide-in-from-top-4 duration-700` nos cards de resultado — uma classe customizada (provavelmente definida em outro arquivo ou dependência não identificada).

---

🤖 *Tokens extraídos por mineração de especificações frontend (PROMPT-MINING-FRONTEND-SPECIFICATION).*
