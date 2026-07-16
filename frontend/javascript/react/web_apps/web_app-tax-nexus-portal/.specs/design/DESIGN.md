---
name: TaxNexus Portal (TaaS)
description: Portal de simulação da Reforma Tributária 2026 — comparativo fiscal com visualização de dados
colors:
  text-body: "#6b6375"
  text-heading: "#08060d"
  bg-primary: "#ffffff"
  border-default: "#e5e4e7"
  code-bg: "#f4f3ec"
  accent-primary: "#aa3bff"
  accent-bg: "rgba(170, 59, 255, 0.1)"
  accent-border: "rgba(170, 59, 255, 0.5)"
  social-bg: "rgba(244, 243, 236, 0.5)"
  text-body-dark: "#9ca3af"
  text-heading-dark: "#f3f4f6"
  bg-primary-dark: "#16171d"
  border-dark: "#2e303a"
  code-bg-dark: "#1f2028"
  accent-dark: "#c084fc"
  accent-bg-dark: "rgba(192, 132, 252, 0.15)"
  blue-50: "#eff6ff"
  blue-100: "#dbeafe"
  blue-600: "#2563eb"
  blue-800: "#1e40af"
  blue-900: "#1e3a8a"
  orange-50: "#fff7ed"
  orange-200: "#fed7aa"
  orange-800: "#9a3412"
  gray-50: "#f9fafb"
  gray-100: "#f3f4f6"
  gray-400: "#9ca3af"
  gray-700: "#374151"
  yellow-50: "#fefce8"
  yellow-100: "#fef9c3"
  yellow-300: "#fde047"
  yellow-800: "#854d0e"
  red-600: "#dc2626"
  purple-700: "#7e22ce"
  green-600: "#059669"
typography:
  body:
    fontFamily: "system-ui, 'Segoe UI', Roboto, sans-serif"
    fontSize: "18px"
    fontWeight: 400
    lineHeight: 1.45
    letterSpacing: "0.18px"
  heading:
    fontFamily: "system-ui, 'Segoe UI', Roboto, sans-serif"
    fontWeight: 500
  h1:
    fontSize: "56px"
    fontWeight: 500
    letterSpacing: "-1.68px"
  h2:
    fontSize: "24px"
    fontWeight: 500
    lineHeight: 1.18
    letterSpacing: "-0.24px"
  mono:
    fontFamily: "ui-monospace, Consolas, monospace"
    fontSize: "15px"
    lineHeight: 1.35
rounded:
  sm: "4px"
  md: "5px"
  lg: "6px"
  xl: "8px"
spacing:
  xs: "4px"
  sm: "8px"
  md: "16px"
  lg: "24px"
  xl: "32px"
components:
  button-primary:
    backgroundColor: "#2563eb"
    textColor: "#ffffff"
    rounded: "{rounded.xl}"
    padding: "12px 0"
  button-primary-hover:
    backgroundColor: "#1d4ed8"
  button-primary-disabled:
    backgroundColor: "#9ca3af"
  input-text:
    backgroundColor: "{colors.bg-primary}"
    textColor: "{colors.text-heading}"
    rounded: "{rounded.md}"
    padding: "12px"
  card-result:
    backgroundColor: "{colors.blue-50}"
    rounded: "12px"
    padding: "20px"
---

# Design System: TaxNexus Portal

## 1. Overview

**Creative North Star: "A Planilha Fiscal do Futuro"**

O TaxNexus Portal é uma ferramenta de trabalho para profissionais fiscais. Sua linguagem visual empresta da precisão de uma planilha financeira bem diagramada e da clareza de um dashboard executivo. A interface é frugal em decoração e generosa em dados — números formatados, comparações lado a lado, hierarquia visual que guia da entrada à conclusão sem fricção.

A personalidade é **técnica, confiável, direta**. Nada de animações decorativas, nada de glassmorphism, nada de "delight" gratuito. A confiança vem da clareza numérica e da previsibilidade dos componentes. O roxo (#aa3bff) como cor de acento injeta modernidade sem sacrificar a sobriedade fiscal. O azul (#2563eb) governa ações primárias e a marca do produto.

**Key Characteristics:**
- Paleta Restrained com acentos Committed em pontos de ação
- Tipografia de sistema — sem fontes externas, performance máxima
- Componentes com estados completos (default, hover, focus, disabled, loading)
- Dois temas: claro (fiscal/diurno) e escuro (análise/noturno)
- Densidade de dados generosa: tabelas, gráficos e cards de métricas

## 2. Colors

A paleta segue o padrão Tailwind com custom properties complementares. O azul profundo (`blue-900`) ancora a identidade da marca; o roxo (`#aa3bff`) atua como acento de destaque em estados interativos.

### Primary
- **Azul Profundo** (#1e3a8a / blue-900): Cor da marca. Usado em headers, títulos de seção e texto de destaque institucional
- **Azul Ação** (#2563eb / blue-600): Cor de ação primária. Exclusivo para botões de submit, links e indicador de etapa ativa

### Secondary
- **Roxo Acento** (#aa3bff): Cor de ênfase secundária. Estados hover em elementos interativos, badges, bordas de foco. Suavizado para `rgba(170,59,255,0.1)` como fundo de destaque sutil

### Neutral
- **Texto Corpo** (#6b6375): Texto de parágrafo e labels. Contraste 4.62:1 contra fundo branco — atende WCAG AA
- **Texto Heading** (#08060d): Títulos e dados numéricos de alta importância. Contraste 18.1:1 — confortável para leitura prolongada
- **Fundo** (#ffffff): Fundo principal. Neutro puro, sem tint
- **Borda** (#e5e4e7): Bordas de container, divisores, separadores de seção
- **Cinza Superfície** (#f3f4f6 / gray-100): Fundo de página. Leve distinção do conteúdo sem competir

### Semantic
- **Laranja Legado** (#fff7ed / orange-50): Cards do sistema tributário atual — tom quente para "modelo antigo"
- **Azul Reforma** (#eff6ff / blue-50): Cards do novo sistema — tom frio para "modelo novo"
- **Amarelo Alerta** (#fefce8 / yellow-50): Seção de saldo remanescente — atenção sem alarme
- **Vermelho Ação** (#dc2626): Links de saída, ações destrutivas

### Named Rules
**The One Accent Rule.** O roxo (#aa3bff) aparece em ≤5% de qualquer tela. Sua raridade é o que o torna efetivo como sinal de interação.

## 3. Typography

**Display/Heading Font:** System UI (system-ui, Segoe UI, Roboto, sans-serif)
**Body Font:** System UI (system-ui, Segoe UI, Roboto, sans-serif)
**Mono Font:** UI Monospace (ui-monospace, Consolas, monospace)

**Character:** Funcional e invisível. A tipografia não busca chamar atenção para si — fontes de sistema eliminam latency de carregamento e respeitam as preferências do OS do usuário.

### Hierarchy
- **H1** (500, 56px, letter-spacing: -1.68px): Título da página/aplicação. Uso único por view
- **H2** (500, 24px, line-height: 1.18, letter-spacing: -0.24px): Títulos de seção dentro de componentes
- **Body** (400, 18px, line-height: 1.45): Texto de parágrafo. Em telas ≤1024px reduz para 16px
- **Label** (500, 14px / 16px): Labels de formulário, cabeçalhos de card, métricas
- **Mono** (400, 15px, line-height: 1.35): Códigos NCM, CNPJ, valores monetários, dados tabulares

### Named Rules
**The No Display Font Rule.** Fontes display são proibidas em UI de produto. Toda a hierarquia usa a mesma família sans — variação vem de peso, tamanho e cor, não de família.

## 4. Elevation

O sistema é majoritariamente plano. A única sombra definida é `0 10px 15px -3px rgba(0,0,0,0.1), 0 4px 6px -2px rgba(0,0,0,0.05)` — usada exclusivamente em cards de resultado e no formulário de login. Não há escala de elevação (sm/md/lg/xl); a distinção entre camadas é feita por bordas e cor de fundo.

### Shadow Vocabulary
- **Card Elevation** (`box-shadow: 0 10px 15px -3px rgba(0,0,0,0.1), 0 4px 6px -2px rgba(0,0,0,0.05)`): Cards de conteúdo (login, resultados). No tema escuro, shadows são mais intensas: `rgba(0,0,0,0.4)` e `rgba(0,0,0,0.25)`

### Named Rules
**The Flat-By-Default Rule.** Superfícies são planas em repouso. Sombras aparecem apenas em containers elevados (cards, modals). Backgrounds coloridos (blue-50, orange-50) substituem sombras para distinção entre seções.

## 5. Components

### Buttons
- **Shape:** Cantos arredondados 8px (rounded-lg); padding vertical 12px; full-width em formulários
- **Primary:** Fundo `blue-600` (#2563eb), texto branco, font-weight 600. Altura mínima 48px para touch targets
- **Hover:** Escurece para `blue-700` (#1d4ed8)
- **Disabled:** Fundo `gray-400` (#9ca3af), sem hover. Usado durante loading da API
- **Text/Secondary:** Links de texto (ex: "Sair / Trocar CNPJ") usam `red-600` com underline no hover

### Inputs / Fields
- **Style:** Borda 1px solid, fundo branco, border-radius 6px, padding 12px (p-3)
- **Focus:** Sem estilo de foco customizado — usa o default do navegador
- **Select:** Idêntico ao input, com disabled state quando dependente de seleção anterior (ex: cidade depende de estado)
- **Number:** Campo de saldo com borda `yellow-300` sobre fundo `yellow-50` para destaque contextual

### Cards / Containers
- **Login Card:** Fundo branco, padding 32px (p-8), border-radius 8px (rounded-lg), sombra md
- **Result Cards:** Fundo `orange-50` (legado) ou `blue-50` (reforma), padding 20px (p-5), border-radius 12px (rounded-xl), borda `orange-200` ou `blue-200`
- **Simulator Container:** Fundo branco, padding 24px (p-6), max-width 56rem (max-w-4xl), border-radius 8px, sombra lg
- **Saldo Remanescente:** Fundo `yellow-50`, borda `yellow-100`, padding 16px, border-radius 6px — destaque contextual sem ser erro

### Navigation
- **Header:** Flex row com justify-between. Título "TaxNexus Simulator v1.0" à esquerda, botão "Sair" à direita. Texto `blue-900` para brand, `red-600` para ação de saída
- **Sem sidebar ou breadcrumbs** — aplicação single-page com fluxo linear

### Data Visualization
- **BarChart (Recharts):** Container com fundo branco, borda, sombra interna (shadow-inner), altura 320px (h-80). Barras: CBS em `#2563eb` (azul), IBS em `#059669` (verde). Grid tracejada horizontal, sem grid vertical. Tooltip customizada com formatação BRL (pt-BR)

## 6. Do's and Don'ts

### Do:
- **Do** usar `bg-white rounded-lg shadow-md` para cards de conteúdo principal
- **Do** usar `orange-50` para sistema atual (legado) e `blue-50` para pós-reforma — a distinção cromática ensina o modelo mental
- **Do** formatar TODOS os valores monetários com `toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })`
- **Do** usar `disabled:bg-gray-400` em botões durante carregamento assíncrono
- **Do** mostrar o CNPJ do usuário logado no header como indicador de contexto
- **Do** usar fontes de sistema (`system-ui`) — sem Google Fonts, sem latência de rede

### Don't:
- **Don't** usar glassmorphism, backdrop-filter, ou blur decorativo — é ferramenta fiscal, não portfolio de design
- **Don't** usar gradientes em texto (`background-clip: text`) — planilhas não têm texto degradê
- **Don't** usar animações de entrada coreografadas — o usuário quer ver números, não um espetáculo
- **Don't** usar modo escuro com roxo/neon — manter tons neutros no tema escuro
- **Don't** usar "SaaS landing-page clichés" (hero metrics, card grids idênticos, eyebrow labels, glassmorphism)
- **Don't** usar `border-left` > 2px como stripe decorativa em cards
- **Don't** usar `border-radius` > 16px em cards — 12px (rounded-xl) é o teto
- **Don't** usar sombra + borda decorative no mesmo elemento — escolha um
