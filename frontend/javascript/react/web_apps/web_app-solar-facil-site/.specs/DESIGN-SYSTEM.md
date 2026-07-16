# Sistema de Design: Solar Fácil Site

> **Especificação conjunta** — sintetizando análise do skill `frontend-design` com observações de engenharia reversa do código.
> Baseado em: `DESIGN.md` (brief de design original), `globals.css` (tokens implementados), todos os componentes.
> Gerado por `/spec-miner` + `frontend-design` em 2026-07-05. Revisado para pt-BR e integrado com DESIGN.md completo em 2026-07-06.

---

## 1. Norte Criativo: "A Casa Conectada"

> *"A Solar Fácil é o encontro entre o conforto do lar brasileiro e a inteligência de uma rede de energia moderna."* — DESIGN.md

A linguagem visual do site trata cada seção como um "cômodo" com sua própria atmosfera — alguns amplos e iluminados, outros íntimos e focados. A energia flui entre os espaços naturalmente. A estética é **acolhedora mas precisa**: warm but precise, acolhedora mas confiável.

A paleta verde-terra e teal transmite sustentabilidade sem clichês ecológicos; o amarelo-luz pontua momentos de destaque como a luz solar entrando por uma janela. Os componentes são **receptivos e vivos**: respondem ao toque com elevação sutil e transições imediatas, como interruptores bem projetados.

### Anti-referências (explicitamente rejeitadas)
- ❌ Frieza de startup tech (dark mode, neon, glassmorphism)
- ❌ Visual de burocracia governamental (bandeiras, brasões, azul imperial)
- ❌ Clichês genéricos de "energia verde" (ícones de folha, gota d'água, gradientes ecológicos)

### Características-Chave
- Paleta verde-terra ancorada em confiança, não em "ecologia de marketing"
- Tipografia única (Inter) com variação expressiva de pesos para hierarquia sem múltiplas famílias
- Componentes que respondem ao toque — hover com elevação, foco visível, transições rápidas
- Ritmo visual variado entre seções — cada "cômodo" tem sua própria atmosfera
- Transparência como linguagem visual: números grandes, claros, verificáveis

---

## 2. Sistema de Cores

### Tokens Implementados

Todos os tokens definidos em `src/app/globals.css:8-36` via Tailwind v4 `@theme inline`.

#### Primárias — *"O Chão da Marca"*

| Token CSS | Nome de Design | Hex | OKLCH | Uso |
|-----------|---------------|-----|-------|-----|
| `--color-solar-primary` | **Verde Terra** | `#1E5631` | `oklch(36% 0.08 155)` | Botões primários, headlines, links ativos, anéis de foco |
| `--color-solar-primary-dark` | **Verde Terra Profundo** | `#153d23` | — | Estados hover/ativos |
| `--color-solar-primary-light` | **Verde Orvalho** | `#dcfce7` | — | Fundos de seleção, painéis de resultado, cards destacados |

**Verde Terra** é a cor da confiança. Usada em botões primários, headlines de seção, elementos de destaque. É o "chão" da marca — presente, estável, sem gritar.
**Verde Orvalho** é a versão mais clara do verde, como luz filtrada por folhas.

#### Secundárias — *"O Contraponto Arejado"*

| Token CSS | Nome de Design | Hex | Uso |
|-----------|---------------|-----|-----|
| `--color-solar-secondary` | **Teal Serena** | `#A5C9CA` | Botões secundários, calculadora fornecedor, links |
| `--color-solar-secondary-dark` | **Teal Profundo** | `#7ab5b5` | Estados hover |
| `--color-solar-secondary-light` | **Névoa Teal** | `#e6f4f4` | Painéis de resultado secundários, seções de fundo |

**Teal Serena** é o contraponto arejado ao verde-terra. Traz leveza sem perder seriedade.
**Névoa Teal** é o fundo alternativo frio para seções que pedem calma — FAQ, informações regulatórias.

#### Base — *"Fundação Neutra"*

| Token CSS | Nome de Design | Hex | Uso |
|-----------|---------------|-----|-----|
| `--color-solar-bg` | **Branco Solar** | `#ffffff` | Fundo principal |
| `--color-solar-bg-alt` | **Luz do Sol** | `#ffffbf` | Fundo alternado de seções (amarelo claro) |
| `--color-solar-text` | **Terra Escura** | `#111827` | Texto principal (quase preto, leve tom azulado) |
| `--color-solar-text-muted` | **Sombra Suave** | `#444444` | Texto secundário/muted (contraste ≥4.5:1 verificado) |
| `--color-solar-border` | **Linha do Horizonte** | `#cccccc` | Bordas, divisores |

**Luz do Sol** é um amarelo bem claro e quente, como luz solar indireta. Usado com moderação.
**Terra Escura** é a cor de todo número exibido na interface.

#### Status

| Token CSS | Nome de Design | Hex | Uso |
|-----------|---------------|-----|-----|
| `--color-solar-success` | **Verde Confirma** | `#e0f7e9` | Fundos de sucesso |
| `--color-solar-warning` | **Amarelo Alerta** | `#fffde7` | Fundos de aviso |
| `--color-solar-error` | **Vermelho Atenção** | `#fbe9e7` | Fundos de erro |
| `--color-solar-error-border` | — | `#ef4444` | Bordas de campo com erro |
| `--color-solar-error-text` | — | `#dc2626` | Mensagens de erro |
| `--color-solar-error-surface` | — | `#fef2f2` | Fundos de painel de erro |

### Regras de Cor Nomeadas (do DESIGN.md, verificadas no código)

**Regra da Uma Voz**: O Verde Terra (`--color-solar-primary`) ocupa ≤15% de qualquer tela. Sua raridade é o que lhe dá peso. Está presente em:
- Botões primários (`Button.tsx`)
- Link ativo da navegação (`Header.tsx:17`)
- Texto de destaque na headline (`HeroSection.tsx:21`)
- Anéis de foco (global)
- *Verificado: Sim, esses usos são esparsos o suficiente para cumprir.*

**Regra da Luz Natural**: Luz do Sol (`--color-solar-bg-alt`, `#ffffbf`) aparece em no máximo 30% das seções. Uso atual:
- CalculatorSection (`page.tsx:32`: `bg="alt"`)
- PlansSection (`PlansSection.tsx:8`: `bg="alt"`)
- HowItWorksSection (`HowItWorksSection.tsx:13`: `bg="alt"`)
- Footer (`Footer.tsx:13`: `bg-solar-bg-alt`)
- *São 4 de ~10 seções = 40%. Excede ligeiramente a regra.*

**Regra da Confiança Visível**: Todo número exibido na interface (economia, desconto, kWh) usa Terra Escura (`#111827`) com peso bold ou extrabold. Verificado em:
- MetricCard (`font-extrabold text-solar-text`)
- Preço do PlanCard (`font-extrabold text-solar-text`)
- Resultados da calculadora (valores de economia/ganho)
- *Verificado: Nenhum número escondido em texto muted.*

---

## 3. Sistema Tipográfico

**Fonte**: Inter — família única, carregada via `next/font/google` em `layout.tsx:9-12`.
**Filosofia**: Hierarquia através de peso (400→500→700→800) e tamanho, não através de troca de fontes.
**Caráter**: Inter é uma sans-serif geométrica-humanista que funciona em densidades extremas — do Thin (100) ao ExtraBold (800). A variação de peso constrói hierarquia onde outras marcas usariam múltiplas famílias.

### Escala Implementada

| Papel | Peso | Tamanho | Altura de Linha | Letter Spacing | CSS |
|-------|------|---------|-----------------|----------------|-----|
| **Display** | 800 (extrabold) | `clamp(2.25rem, 5vw, 3.75rem)` | 1.1 | `-0.02em` | Headline do hero |
| **Headline** | 700 (bold) | `clamp(1.875rem, 3vw, 2.25rem)` | 1.2 | `-0.01em` | Títulos de seção |
| **Title** | 700 (bold) | `1.25rem` | 1.3 | — | Títulos de cards, headings de formulário |
| **Body** | 400 (normal) | `1rem` | 1.6 | — | Texto corrido |
| **Label** | 500 (medium) | `0.875rem` | 1.4 | — | Navegação, labels de formulário, badges |

### Observações no Código

```css
/* Hero — implementado como utilitários Tailwind */
text-4xl font-extrabold tracking-tight    /* mobile */
sm:text-5xl md:text-6xl                   /* escala responsiva */
[text-wrap:balance]                       /* linhas balanceadas */

/* Títulos de seção */
text-3xl font-bold md:text-4xl
[text-wrap:balance]

/* Números em resultados */
text-3xl md:text-4xl font-extrabold

/* Texto corrido */
text-base text-solar-text-muted sm:text-lg md:text-xl
[text-wrap:pretty]
```

**Verificação da Regra do Peso Único**: Apenas Inter é usada. Nenhuma fonte secundária existe no código. *Verificado.*

**Verificação da Regra do Piso do Display**: Tracking é `-0.02em` (via `tracking-tight`), bem acima do mínimo de -0.04em. *Verificado.*

---

## 4. Espaçamento e Layout

### Ritmo de Seção

```
SectionWrapper padrão:
  padding-y: py-16 (mobile) / md:py-24 (desktop)
  padding-x: px-4
  max-width: max-w-7xl (1280px)
  scroll-margin-top: 80px (compensa cabeçalho fixo)
```

### Alternância de Fundo de Seção (Homepage)

| Ordem | Seção | Fundo |
|-------|-------|-------|
| 1 | HeroSection | branco + gradiente overlay |
| 2 | CalculatorSection | `bg="alt"` (Luz do Sol) |
| 3 | ProofSection | Full-bleed `bg-solar-primary` (Verde Terra) |
| 4 | HowItWorksSection | `bg="alt"` (Luz do Sol) |
| 5 | PlansSection | `bg="alt"` (Luz do Sol) |
| 6 | DifferentiatorsSection | Branco Solar |
| 7 | FinalCtaSection | Branco Solar |

---

## 5. Sistema de Elevação

**Filosofia**: "Flat-By-Default" — superfícies nascem planas. Sombra é sempre resposta a estado (hover, foco), nunca decoração.

| Token | Valor | Uso |
|-------|-------|-----|
| `hover-card` | `0 4px 12px rgba(0,0,0,0.08)` | Hover de card (transição 200ms) |
| `focus-ring` | `0 0 0 2px var(--color-solar-primary)` | Todos elementos interativos |
| `header` | `backdrop-filter: blur(12px)` | Profundidade do cabeçalho fixo |

### Vocabulário de Sombras
- **hover-card**: Elevação sutil ao passar o mouse em cards. Suficiente para destacar, insuficiente para flutuar.
- **focus-ring**: Anel de foco visível em todos os elementos interativos. 2px sólido, sem blur — máximo de acessibilidade.
- **header**: Header fixo com desfoque de fundo. Transmite camadas sem sombra — o conteúdo rola "atrás" do header.

### Regra do Cartão Fantasma (No Ghost Card Rule)
Nunca combinar `border: 1px solid` com `box-shadow` de blur ≥16px no mesmo elemento. Borda OU sombra — nunca os dois como decoração. Cards em repouso têm apenas `border: 1px solid` sem sombra. *Verificado em PlanCard.tsx, MetricCard.tsx.*

---

## 6. Tokens de Design de Componentes

### Botões
```
Formato: rounded-lg (8px)
Tipografia: font-bold
Transição: 200ms ease-in-out (apenas cores, sem transform)
Foco: focus-visible:ring-2 focus-visible:ring-offset-2

Tamanhos:
  sm: px-4 py-2 text-sm
  md: px-6 py-2.5 text-base
  lg: px-6 py-2.5 text-base

Variantes:
  primary:   bg-solar-primary → hover:bg-solar-primary-dark
  secondary: bg-solar-secondary → hover:bg-solar-secondary-dark
  outline:   border-2 border-solar-primary → hover:bg-solar-primary-light
```

**Caráter**: Convidativos e confiáveis. Como interruptores físicos bem projetados — você sabe exatamente o que vai acontecer ao pressionar.

### Cards
```
Formato: rounded-lg (8px)
Borda: border border-solar-border
Fundo: bg-solar-bg
Padding: p-4
Hover: shadow-md (transição 200ms)

Destacado: ring-2 ring-solar-primary/30 scale-[1.02]
```

**Caráter**: Molduras de informação, não gavetas fechadas. Leves, abertos, com hierarquia interna clara.

### Inputs / Campos
```
Formato: rounded-lg (8px)
Borda: border border-solar-border
Padding: px-3 py-2.5
Foco: ring-2 ring-solar-primary (borda muda para primary)
Erro: border-solar-error-border bg-solar-error-surface
Disabled: opacidade 50%, cursor not-allowed
```

**Caráter**: Convidativos ao toque. A borda escurece no foco; o campo "acorda" quando você interage.

### Grupo de Rádio / Chips
```
Repouso: border-solar-border text-solar-text-muted
Selecionado: border-solar-primary bg-solar-primary-light text-solar-primary-dark
Hover: hover:border-solar-primary/50
```

---

## 7. Movimento e Animação

### Keyframes Definidos (`globals.css:62-69`)

```css
@keyframes fade-in {
  from { opacity: 0; transform: translateY(8px); }
  to   { opacity: 1; transform: translateY(0); }
}

@keyframes slide-in {
  from { transform: translateX(100%); }
  to   { transform: translateX(0); }
}
```

### Uso nos Componentes

| Animação | Duração | Easing | Onde |
|----------|---------|--------|------|
| `.animate-fade-in` | 400ms | ease-out | Resultados da calculadora |
| `.animate-slide-in` | 300ms | ease-out | Painel do menu mobile |
| Transição de cor dos botões | 200ms | ease-in-out | Todos botões |
| Rotação do chevron FAQ | 300ms | — | Toggle do accordion |
| Transição de altura FAQ | 300ms | — | Painel do accordion |
| Indicador de scroll | — | — | `motion-safe:animate-bounce` |
| Skeleton pulse | — | — | `motion-safe:animate-pulse` |

### Movimento Reduzido
Todas as animações respeitam `prefers-reduced-motion: reduce` via `globals.css:73-82`:
- Todas as durações → 0.01ms
- Todas as iterações → 1
- Scroll suave → auto
- `animate-fade-in` e `animate-slide-in` aplicados apenas dentro de `@media (prefers-reduced-motion: no-preference)`

---

## 8. Breakpoints Responsivos

| Breakpoint | Prefixo | Exemplos de Uso |
|------------|---------|-----------------|
| 640px | `sm:` | Grid 2 col, layout CTA horizontal |
| 768px | `md:` | Nav desktop visível, timeline horizontal, aumento padding seção |
| 1024px | `lg:` | Grid 3 col planos, tabela comparativa, footer 4 col |

### Padrões Mobile-First
- **Navegação**: Hambúrguer → nav desktop (md:)
- **Planos**: Pilha de cards → tabela comparativa (lg:)
- **Footer**: 1-col → 2-col (sm:) → 4-col (lg:)
- **Calculadora**: 1-col → 2-col (lg:)
- **ComoFunciona**: Timeline vertical → timeline horizontal (md:)

---

## 9. Implementação de Acessibilidade

| Requisito | Implementação | Localização |
|-----------|---------------|-------------|
| Link pular conteúdo | `sr-only focus:not-sr-only`, top-left, z-100 | `layout.tsx:45-50` |
| Anéis de foco | `focus-visible:ring-2 focus-visible:ring-offset-2` em todos elementos interativos | Button, inputs, links |
| Movimento reduzido | Sobrescrita global `prefers-reduced-motion: reduce` | `globals.css:73-82` |
| Landmarks ARIA | `<header>`, `<main>`, `<footer>`, `<nav>` | Componentes de layout |
| aria-expanded | Triggers do accordion FAQ | `FaqAccordion.tsx:27` |
| aria-controls | Trigger→painel do accordion FAQ | `FaqAccordion.tsx:28` |
| role="status" | Painéis de resultado da calculadora | `ConsumerCalculator.tsx:67` |
| aria-label | Toggle menu mobile, nav breadcrumb | `Header.tsx:56`, `Breadcrumb.tsx:10` |
| aria-hidden | Elementos decorativos (skeletons, scroll indicator) | Múltiplas localizações |
| text-wrap: balance | Todas as headlines | Global |
| text-wrap: pretty | Parágrafos longos | `HeroSection.tsx:24` |
| lang="pt-BR" | Elemento HTML | `layout.tsx:43` |
| Suavização de fonte | `-webkit-font-smoothing: antialiased` | `globals.css:47` |

---

## 10. O que Fazer e o que NÃO Fazer (Do's and Don'ts)

> Fonte: `DESIGN.md` — diretrizes originais do brief de design, verificadas contra a implementação.

### ✅ FAÇA:
- **Use** Inter em toda a interface — uma família, hierarquia por peso e tamanho
- **Use** Verde Terra (`#1E5631`) como acento primário em ≤15% de qualquer tela
- **Alterne** fundos de seção (Branco Solar → Luz do Sol → Névoa Teal) para ritmo visual
- **Use** `text-wrap: balance` em headlines e `text-wrap: pretty` em parágrafos longos
- **Mostre** números de economia em peso extrabold com Terra Escura
- **Inclua** anel de foco visível (2px) em todo elemento interativo
- **Use** fotografia real de casas e pessoas brasileiras — nada de stock photos genéricas
- **Escreva** alt text descritivo e em português natural

### ❌ NÃO FAÇA:
- **Não use** verde-limão, gradientes ecológicos, ou ícones de folha/gota — é o clichê genérico de energia
- **Não aplique** dark mode, neon, glassmorphism ou estética "startup tech fria"
- **Não use** azul marinho, brasões, bandeiras do Brasil ou estética de site governamental
- **Não combine** `border: 1px solid` com `box-shadow` de blur ≥16px no mesmo elemento (Regra do Cartão Fantasma)
- **Não use** `border-radius` acima de 16px em cards ou seções
- **Não coloque** eyebrow (texto pequeno tracked uppercase) acima de toda seção — um kicker deliberado é voz; em toda seção é gramática de IA
- **Não use** `border-left` ou `border-right` > 1px como faixa colorida decorativa em cards ou listas
- **Não esconda** números importantes em Sombra Suave (`#444444`) — números são sempre Terra Escura com peso bold
- **Não use** imagens genéricas de bancos de fotos internacionais — o Brasil real pede imagens reais
- **Não escreva** em inglês ou "portunhol" — a voz é em português brasileiro natural

---

## 11. Calculadora (Componente Assinatura)

> Fonte: `DESIGN.md` — especificação original do componente

**Caráter**: O coração funcional do site. Duas calculadoras lado a lado (Consumidor / Fornecedor) que transformam números abstratos em economia real.

- **Formato**: Abas horizontais para seleção de perfil, campos numéricos com labels claras, resultado em destaque
- **Exibição do resultado**: Número grande (peso Display 800, Verde Terra ou Terra Escura), label explicativa abaixo
- **Fundo**: Branco Solar com borda 1px Linha do Horizonte, cantos 8px
- **Input**: Campo numérico com input direto. Cor de foco: Verde Terra.

---

## 12. Avaliação de Fidelidade da Implementação

Comparando a especificação do DESIGN.md com o código CSS/componentes real:

| Especificação de Design | Status no Código | Observações |
|--------------------------|------------------|-------------|
| Cores como tokens CSS custom properties | ✅ Implementado | `--color-solar-*` em `@theme inline` |
| Inter como fonte única | ✅ Implementado | `next/font/google` no layout |
| Variantes de botão (primary/secondary/outline) | ✅ Implementado | `Button.tsx` com 3 variantes |
| Tamanhos de botão (sm/md/lg) | ✅ Implementado | Mas md e lg têm padding idêntico |
| Sombra de card apenas no hover | ✅ Implementado | `transition-shadow hover:shadow-md` |
| Anéis de foco em todos interativos | ✅ Implementado | Consistente `focus-visible:ring-2` |
| Header com backdrop-blur | ✅ Implementado | `bg-white/80 backdrop-blur-md` |
| Alternância de fundo de seção | ⚠️ Parcial | Luz do Sol usado ~40% vs meta 30% |
| Selo ANEEL visível | ✅ Implementado | Footer + ProofSection |
| Movimento reduzido | ✅ Implementado | Sobrescrita global em globals.css |
| `text-wrap: balance` | ✅ Implementado | Em todas as headlines |
| Modo escuro | ❌ Não implementado | Comentado como milestone futuro |
| Fotografia (casas brasileiras reais) | ❌ Não implementado | Nenhuma imagem no site ainda |
