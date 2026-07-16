# Catálogo de Componentes: Solar Fácil Site

> **Especificação de engenharia reversa** — todos os componentes, suas props, estados e comportamentos conforme observado no código-fonte.
> Gerado por `/spec-miner` em 2026-07-05. Revisado para pt-BR em 2026-07-06.

---

## 1. Componentes Compartilhados (Shared)

### Button
**Arquivo**: `src/components/shared/Button.tsx`
**Tipo**: Server component (pode ser usado tanto em server quanto client)

| Prop | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `variant` | `'primary' \| 'secondary' \| 'outline'` | `'primary'` | Estilo visual |
| `size` | `'sm' \| 'md' \| 'lg'` | `'lg'` | Preset de tamanho |
| `children` | `ReactNode` | — | Conteúdo do botão |
| `href` | `string?` | — | Se fornecido, renderiza como `<a>` em vez de `<button>` |
| `className` | `string?` | `''` | Classes adicionais |
| `...props` | `ButtonHTMLAttributes` | — | Todos atributos nativos de button |

| Variante | Fundo | Cor do Texto | Hover |
|----------|-------|-------------|-------|
| `primary` | `bg-solar-primary` | `white` | `bg-solar-primary-dark` |
| `secondary` | `bg-solar-secondary` | `white` | `bg-solar-secondary-dark` |
| `outline` | `transparent` + `border-2 border-solar-primary` | `text-solar-primary` | `bg-solar-primary-light` |

| Tamanho | Padding | Tamanho da Fonte | Border Radius |
|---------|---------|-----------------|---------------|
| `sm` | `px-4 py-2` | `text-sm` | `rounded-lg` |
| `md` | `px-6 py-2.5` | `text-base` | `rounded-lg` |
| `lg` | `px-6 py-2.5` | `text-base` | `rounded-lg` |

**Estados**: padrão, hover (transição de cor 200ms), focus-visible (anel 2px + offset 2px), disabled (nativo)

---

### SectionWrapper
**Arquivo**: `src/components/shared/SectionWrapper.tsx`
**Tipo**: Server component

| Prop | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `id` | `string?` | — | ID da seção para alvos de scroll |
| `children` | `ReactNode` | — | Conteúdo da seção |
| `className` | `string?` | `''` | Classes adicionais |
| `bg` | `'white' \| 'alt'` | `'white'` | Cor de fundo |

Renderiza `<section>` com `px-4 py-16 md:py-24` e container interno `max-w-7xl`.

---

### MetricCard
**Arquivo**: `src/components/shared/MetricCard.tsx`
**Tipo**: Server component

| Prop | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `value` | `string` | — | Valor exibido |
| `label` | `string` | — | Texto descritivo |
| `highlighted` | `boolean` | `false` | Ênfase visual |
| `highlightColor` | `'green' \| 'amber'` | `'green'` | Família de cor do destaque |
| `inverted` | `boolean` | `false` | Modo branco-sobre-escuro |

**Dois modos:**
- **Padrão**: card branco com borda, texto secundário em muted, opção de destaque verde/amber
- **Invertido**: card transparente sobre fundo escuro, texto branco, textos semi-transparentes

**Padrão destacado**: `bg-solar-primary-light` + `ring-2 ring-solar-primary/20` + `scale-105`
**Invertido destacado**: `bg-white/15` + `ring-2 ring-white/30` + `scale-105`

---

### Skeleton
**Arquivo**: `src/components/shared/Skeleton.tsx`
**Tipo**: Server component

| Prop | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `variant` | `'text' \| 'card' \| 'circle' \| 'input'` | `'text'` | Formato |
| `className` | `string?` | `''` | Classes adicionais (ex.: altura personalizada) |

Todas as variantes: `bg-solar-text-muted/10` + `motion-safe:animate-pulse` + `aria-hidden="true"`

---

### Breadcrumb
**Arquivo**: `src/components/shared/Breadcrumb.tsx`
**Tipo**: Server component

| Prop | Tipo | Descrição |
|------|------|-----------|
| `items` | `{ label: string; href?: string }[]` | Trilha de breadcrumb |

Renderiza `<nav aria-label="Breadcrumb">` com `<ol>`, separadores de chevron, e:
- Links para itens não-últimos com `href`
- Texto escuro em negrito para o último item
- Texto muted para itens sem href

---

### JsonLd
**Arquivo**: `src/components/shared/JsonLd.tsx`
**Tipo**: Server component (sem props)

Injeta dados estruturados Schema.org `Organization` via `<script type="application/ld+json">` com:
- Nome, URL, descrição, logo
- sameAs: Instagram, LinkedIn
- contactPoint: serviço de atendimento (email, areaServed: BR, idioma: Portuguese)

---

### AnalyticsProvider
**Arquivo**: `src/components/shared/AnalyticsProvider.tsx`
**Tipo**: Server component (sem props)

Injeta condicionalmente o script Google Analytics 4:
- Omite se `GA_MEASUREMENT_ID` não estiver definido ou for o placeholder `'G-XXXXXXXXXX'`
- Carrega gtag.js de forma assíncrona
- Inicializa com `send_page_view: true`

---

## 2. Componentes de Layout

### Header
**Arquivo**: `src/components/layout/Header.tsx`
**Tipo**: Client component (`'use client'`)

**Estrutura:**
```
<header> (fixo, z-50, backdrop-blur, border-b)
  <div> (max-w-7xl, flex, h-16)
    Logo (Image 120×24, priority)
    Navegação desktop (hidden md:flex)
      Link Planos
      Link Contato
      Botão CTA "Baixar App"
    Hambúrguer mobile (md:hidden)
  Espaçador div (h-16)
  <MobileMenu>
```

**Link ativo**: Detectado via `usePathname()` — link ativo recebe `text-solar-primary font-semibold`

---

### MobileMenu
**Arquivo**: `src/components/layout/MobileMenu.tsx`
**Tipo**: Client component

| Prop | Tipo | Descrição |
|------|------|-----------|
| `open` | `boolean` | Visibilidade |
| `onClose` | `() => void` | Callback de fechamento |

**Comportamentos:**
- Fecha na tecla Escape (listener adicionado/removido ao abrir)
- Bloqueio de scroll do body quando aberto (`overflow: hidden`)
- Backdrop com blur (`bg-black/50 backdrop-blur-sm`)
- Painel deslizante da direita (`motion-safe:animate-slide-in`)
- Links: Home, Planos, Contato, CTA "Baixar App"

**Limpeza**: Funções de retorno do `useEffect` removem o listener de evento e restauram o scroll do body.

---

### Footer
**Arquivo**: `src/components/layout/Footer.tsx`
**Tipo**: Server component

**Estrutura (grid de 4 colunas → empilhado no mobile):**
1. **Marca**: Logo + tagline
2. **Navegação**: Links Home, Planos, Contato
3. **Baixe o App**: Links App Store + Google Play
4. **Contato**: Links Email, Instagram, LinkedIn

**Barra inferior**: Selo ANEEL (ícone ShieldCheck + "Plataforma regulamentada pela ANEEL (RN 687/2015)") + Termos/Privacidade/Copyright

---

## 3. Componentes da Homepage

### HeroSection
**Arquivo**: `src/components/home/HeroSection.tsx`
**Tipo**: Client component

**Estrutura:**
```
<section> (min-h-screen, flex, centralizado)
  Gradiente de fundo (solar-primary-light/50 → solar-bg)
  Conteúdo (max-w-3xl)
    h1: "Energia Limpa\nCompartilhada" (solar-primary na segunda linha)
    p: Proposta de valor
    CTAs duplos:
      "Quero Economizar" (primary) → scroll para calculadora-consumidor
      "Quero Compartilhar Energia" (secondary) → scroll para calculadora-fornecedor
  Indicador de scroll (chevron quicando, aria-hidden)
```

**Analytics**: `trackCtaClick` disparado no clique do CTA com location `'hero'`.

---

### CalculatorSection
**Arquivo**: `src/components/home/CalculatorSection.tsx`
**Tipo**: Server component

Agrupa duas calculadoras em um grid de 2 colunas (`lg:grid-cols-2`) sobre fundo amarelo claro. Cada calculadora tem seu próprio ID de alvo de scroll para navegação por CTA.

---

### ConsumerCalculator
**Arquivo**: `src/components/home/ConsumerCalculator.tsx`
**Tipo**: Client component

**Estados:**
- **Inicial**: Input + botão "Calcular"
- **Erro**: Borda vermelha + texto de erro abaixo do input
- **Resultado**: Painel de resultado animado (`motion-safe:animate-fade-in`) com economia, plano sugerido, CTA "Quero este plano"
- **Limpar**: Botão "Limpar" aparece após o primeiro resultado

**Painel de resultado**: Fundo `bg-solar-primary-light`, `role="status"` para leitores de tela.
**Link CTA**: `/contato?perfil=consumidor&plano={plano}&economia={valor}`

---

### ProviderCalculator
**Arquivo**: `src/components/home/ProviderCalculator.tsx`
**Tipo**: Client component

Mesma estrutura do ConsumerCalculator mas com:
- Esquema de cores teal/secundário (`bg-solar-secondary-light`)
- Sufixo kWh no input
- Exibição da taxa de referência (`{PROVIDER_RATE}/kWh`)
- Link CTA: `/contato?perfil=fornecedor&excedente={valor}&ganho={valor}`

---

### ProofSection
**Arquivo**: `src/components/home/ProofSection.tsx`
**Tipo**: Client component (`useSearchParams`)

**Estrutura:**
```
<section> (bg-solar-primary, full-bleed verde)
  h2: "Quem já confia na Solar Fácil"
  Grid 3 colunas de MetricCard (modo invertido)
  Grid condicional 3 colunas de MetricCard de Fornecedor (se perfil=fornecedor)
  Selo ANEEL (vidro fosco: border-white/25, bg-white/10, backdrop-blur)
```

**Comportamento dinâmico**: Destaca métricas específicas baseado no parâmetro de URL `perfil`.

---

### HowItWorksSection
**Arquivo**: `src/components/home/HowItWorksSection.tsx`
**Tipo**: Server component

**Desktop (≥768px)**: Timeline horizontal com linha de conexão, 3 passos com ícones circulados.
**Mobile (<768px)**: Timeline vertical com segmentos de linha entre passos.

Cada passo: sobrancelha numerada "Passo N" → título → descrição. Ícones mapeados do Lucide via `iconMap`.

---

### PlansSection
**Arquivo**: `src/components/home/PlansSection.tsx`
**Tipo**: Server component

Exibe 3 componentes PlanCard em grid de 3 colunas + card CTA ProviderHighlight + link "Comparar todos os planos" para `/planos`.

---

### DifferentiatorsSection
**Arquivo**: `src/components/home/DifferentiatorsSection.tsx`
**Tipo**: Server component

**Layout assimétrico:**
1. **Diferencial hero**: Card full-width — "R$ 0" em fonte 6xl, mensagem "Zero Capex"
2. **Trio de apoio**: Grid 3 colunas com cards ícone + estatística + descrição
   - Pareamento Automático (teal)
   - Plataforma Legal ANEEL (verde/branco)
   - API + ANEEL (escuro)

---

### FinalCtaSection
**Arquivo**: `src/components/home/FinalCtaSection.tsx`
**Tipo**: Client component

Funil CTA do final da página: headline → subtexto → CTAs duplos (mesmos do hero).
Faz scroll de volta para as calculadoras ao clicar.

---

## 4. Componentes da Página de Planos

### PlanCard
**Arquivo**: `src/components/plans/PlanCard.tsx`
**Tipo**: Server component

| Prop | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `plan` | `Plan` | — | Dados do plano |
| `variant` | `'compact' \| 'full'` | `'compact'` | Quantidade de recursos (3 vs todos) |

**Card destacado**: `border-solar-primary ring-2 ring-solar-primary/30 scale-[1.02]` + selo pill "Mais Popular" (ícone Star, posicionado -top-3).

---

### PlansComparisonTable
**Arquivo**: `src/components/plans/PlansComparisonTable.tsx`
**Tipo**: Server component

**Desktop (≥1024px)**: `<table>` HTML com cabeçalhos de plano, linhas de recursos, linha CTA. Coluna destacada recebe `bg-solar-primary-light` com cantos arredondados.
**Mobile (<1024px)**: Pilha vertical de cards com lista completa de recursos.

---

### ProviderHighlight
**Arquivo**: `src/components/plans/ProviderHighlight.tsx`
**Tipo**: Server component

Card com borda teal, ícone Zap, headline "É um produtor de energia solar?", destaque R$0,40/kWh e CTA para `/contato?perfil=fornecedor`.

---

### FaqAccordion
**Arquivo**: `src/components/plans/FaqAccordion.tsx`
**Tipo**: Client component

**Acessibilidade**:
- `aria-expanded` no botão trigger
- `aria-controls` vinculando trigger ao painel
- `role="region"` no painel
- `aria-labelledby` no painel referenciando o trigger

**Animação**: `transition-[max-height] duration-300` com `max-h-0` (fechado) / `max-h-96` (aberto). Chevron gira 180°.

**Dica inteligente**: Após 3+ FAQs únicos abertos, exibe sugestão de contato.

---

## 5. Componentes da Página de Contato

### ContactForm
**Arquivo**: `src/components/contact/ContactForm.tsx`
**Tipo**: Client component

**Estados:**
- **Editando**: Campos do formulário visíveis, botão "Enviar mensagem"
- **Enviando**: Botão mostra "Enviando..." e fica disabled
- **Sucesso**: Substituído completamente por `<SuccessScreen />`
- **Erro**: Painel de erro de envio com link de fallback WhatsApp

**Ordem dos campos**: Honeypot oculto → Nome* → Email* → Telefone → Perfil* (chips de rádio) → Mensagem → [Painel de erro] → Botão enviar

**Chips de rádio de perfil**: 3 opções com ícones:
- 🏠 Consumidor (quero economizar)
- ⚡ Fornecedor (quero compartilhar)
- 🤝 Cooperativa (quero parceria)

---

### FormField
**Arquivo**: `src/components/contact/FormField.tsx`
**Tipo**: Client component

| Prop | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `label` | `string` | — | Rótulo do campo |
| `name` | `string` | — | Nome/id do input |
| `type` | `'text' \| 'email' \| 'tel' \| 'textarea' \| 'number'` | `'text'` | Tipo de input |
| `value` | `string` | — | Valor controlado |
| `onChange` | `(v: string) => void` | — | Handler de alteração |
| `error` | `string?` | — | Mensagem de erro |
| `required` | `boolean` | `false` | Exibe asterisco vermelho |
| `placeholder` | `string?` | — | Texto placeholder |
| `hidden` | `boolean` | `false` | Renderiza input fora da tela (honeypot) |

**Estado de erro**: Borda vermelha + texto de erro vermelho abaixo do campo.

---

### SuccessScreen
**Arquivo**: `src/components/contact/SuccessScreen.tsx`
**Tipo**: Server component

Layout centralizado: ícone CheckCircle grande (verde) → headline "Obrigado!" → mensagem "resposta em até 24h" → CTA "Baixar App".

---

### JourneySummary
**Arquivo**: `src/components/contact/JourneySummary.tsx`
**Tipo**: Client component

Lê parâmetros de URL (`plano`, `economia`, `excedente`, `ganho`) e exibe um banner de resumo contextual mostrando os resultados da calculadora do usuário. Retorna `null` se nenhum parâmetro de contexto estiver presente.

---

### DirectChannels
**Arquivo**: `src/components/contact/DirectChannels.tsx`
**Tipo**: Server component

Grid de 3 colunas de cards de contato: WhatsApp, Email, Instagram — cada um com ícone, rótulo, informação e link de ação. Rodapé mostra "São Paulo — SP" com ícone de pin de mapa.
