---
name: Solar Fácil
description: Plataforma de energia solar compartilhada — conectamos produtores e consumidores via cooperativas ANEEL
colors:
  verde-terra: "#1E5631"
  verde-terra-profundo: "#153d23"
  verde-orvalho: "#dcfce7"
  teal-serena: "#A5C9CA"
  teal-profundo: "#7ab5b5"
  nevoa-teal: "#e6f4f4"
  branco-solar: "#ffffff"
  luz-do-sol: "#ffffbf"
  terra-escura: "#111827"
  sombra-suave: "#444444"
  linha-do-horizonte: "#cccccc"
  verde-confirma: "#e0f7e9"
  amarelo-alerta: "#fffde7"
  vermelho-atencao: "#fbe9e7"
typography:
  display:
    fontFamily: "Inter, ui-sans-serif, system-ui, sans-serif"
    fontSize: "clamp(2.25rem, 5vw, 3.75rem)"
    fontWeight: 800
    lineHeight: 1.1
    letterSpacing: "-0.02em"
  headline:
    fontFamily: "Inter, ui-sans-serif, system-ui, sans-serif"
    fontSize: "clamp(1.875rem, 3vw, 2.25rem)"
    fontWeight: 700
    lineHeight: 1.2
    letterSpacing: "-0.01em"
  title:
    fontFamily: "Inter, ui-sans-serif, system-ui, sans-serif"
    fontSize: "1.25rem"
    fontWeight: 700
    lineHeight: 1.3
  body:
    fontFamily: "Inter, ui-sans-serif, system-ui, sans-serif"
    fontSize: "1rem"
    fontWeight: 400
    lineHeight: 1.6
  label:
    fontFamily: "Inter, ui-sans-serif, system-ui, sans-serif"
    fontSize: "0.875rem"
    fontWeight: 500
    lineHeight: 1.4
    letterSpacing: "normal"
rounded:
  sm: "8px"
  full: "9999px"
spacing:
  section-y: "clamp(4rem, 6vw, 6rem)"
  section-x: "1rem"
  card-padding: "1rem"
  button-sm: "0.5rem 1rem"
  button-md: "0.625rem 1.5rem"
  button-lg: "0.625rem 1.5rem"
components:
  button-primary:
    backgroundColor: "{colors.verde-terra}"
    textColor: "{colors.branco-solar}"
    rounded: "{rounded.sm}"
    padding: "{spacing.button-lg}"
    typography: "{typography.label}"
  button-primary-hover:
    backgroundColor: "{colors.verde-terra-profundo}"
  button-secondary:
    backgroundColor: "{colors.teal-serena}"
    textColor: "{colors.terra-escura}"
    rounded: "{rounded.sm}"
    padding: "{spacing.button-lg}"
  button-secondary-hover:
    backgroundColor: "{colors.teal-profundo}"
  button-outline:
    backgroundColor: "transparent"
    textColor: "{colors.verde-terra}"
    rounded: "{rounded.sm}"
    padding: "{spacing.button-lg}"
  card-default:
    backgroundColor: "{colors.branco-solar}"
    rounded: "{rounded.sm}"
    padding: "{spacing.card-padding}"
  input-default:
    backgroundColor: "{colors.branco-solar}"
    textColor: "{colors.terra-escura}"
    rounded: "{rounded.sm}"
    padding: "{spacing.button-md}"
---

# Design System: Solar Fácil

## 1. Overview

**Creative North Star: "A Casa Conectada"**

A Solar Fácil é o encontro entre o conforto do lar brasileiro e a inteligência de uma rede de energia moderna. Como uma casa bem projetada, cada seção do site é um "cômodo" com sua própria atmosfera — a sala de estar convida à descoberta, o escritório oferece números precisos, a varanda acolhe o contato. A energia flui naturalmente entre esses espaços, assim como a luz solar percorre uma casa bem iluminada.

A estética é **acolhedora mas precisa** — a paleta verde-terra e teal transmite sustentabilidade sem clichês ecológicos; o amarelo-luz pontua momentos de destaque como a luz solar entrando por uma janela. Os componentes são **receptivos e vivos**: respondem ao toque com elevação sutil e transições imediatas, como interruptores bem projetados. A interface existe para conectar pessoas — não para impressionar com artificialidade.

Este sistema rejeita explicitamente: a frieza de startups tech (dark mode, neon, glassmorphism), a burocracia visual de sites governamentais (brasões, bandeiras, azul imperial), e os clichês genéricos de energia (folhas verdes, gradientes ecológicos, ícones de gota e lâmpada).

**Key Characteristics:**
- Paleta verde-terra ancorada em confiança, não em "ecologia de marketing"
- Tipografia única (Inter) com variação expressiva de pesos para hierarquia sem múltiplas famílias
- Componentes que respondem ao toque — hover com elevação, foco visível, transições rápidas
- Ritmo visual variado entre seções — cada "cômodo" tem sua própria atmosfera
- Transparência como linguagem visual: números grandes, claros, verificáveis

## 2. Colors

A paleta cresce do chão brasileiro: verdes profundos da terra, o teal sereno do horizonte, toques de luz solar amarela. Não há "verde ecológico" genérico nem azul corporativo.

### Primary
- **Verde Terra** (#1E5631): A cor da confiança. Usada em botões primários, headlines de seção, elementos de destaque. É o "chão" da marca — presente, estável, sem gritar. Convertido para OKLCH: `oklch(36% 0.08 155)`.
- **Verde Terra Profundo** (#153d23): Hover e estados ativos de elementos primários. Aprofunda sem mudar de família cromática.
- **Verde Orvalho** (#dcfce7): Fundo de elementos selecionados, badges de sucesso, áreas de destaque sutil. A versão mais clara do verde, como luz filtrada por folhas.

### Secondary
- **Teal Serena** (#A5C9CA): O contraponto arejado ao verde-terra. Usada em botões secundários, links decorativos, gráficos. Traz leveza sem perder seriedade. OKLCH: `oklch(78% 0.05 200)`.
- **Teal Profundo** (#7ab5b5): Hover de elementos secundários.
- **Névoa Teal** (#e6f4f4): Fundo alternativo frio para seções que pedem calma — FAQ, informações regulatórias.

### Neutral
- **Branco Solar** (#ffffff): Fundo principal. Branco puro, sem tintura — a luminosidade vem dos acentos, não do fundo.
- **Luz do Sol** (#ffffbf): Fundo de destaque para seções alternadas. Um amarelo bem claro e quente, como luz solar indireta. Usado com moderação — no máximo 30% das seções.
- **Terra Escura** (#111827): Cor de texto principal. Quase preto, com um leve tom azulado para suavizar a leitura.
- **Sombra Suave** (#444444): Texto secundário e muted. Contraste ≥4.5:1 sobre fundo branco (verificado).
- **Linha do Horizonte** (#cccccc): Bordas e divisores. Neutro, presente mas não dominante.

### Status
- **Verde Confirma** (#e0f7e9): Fundo de mensagens de sucesso, badges de status positivo.
- **Amarelo Alerta** (#fffde7): Fundo de avisos e alertas não-críticos.
- **Vermelho Atenção** (#fbe9e7): Fundo de erros e mensagens críticas.

### Named Rules
**The Uma Voz Rule.** O Verde Terra carrega ≤15% de qualquer tela. Sua raridade é o que lhe dá peso. Botões primários, um headline, um destaque — nunca mais que isso.

**The Luz Natural Rule.** Luz do Sol (#ffffbf) aparece em no máximo 30% das seções. Usado como fundo alternado para criar ritmo, não como cor dominante. Se toda seção tem fundo amarelo, nenhuma tem.

**The Confiança Visível Rule.** Todo número exibido na interface (economia, desconto, kWh) usa Terra Escura (#111827) com peso bold ou extrabold. Números não se escondem em Sombra Suave.

## 3. Typography

**Display Font:** Inter (ui-sans-serif, system-ui, sans-serif fallback)
**Body Font:** Inter (ui-sans-serif, system-ui, sans-serif fallback)
**Label/Mono Font:** Inter (mesma família; variação de peso supre a distinção)

**Character:** Uma única família bem calibrada. Inter é uma sans-serif geométrica-humanista que funciona em densidades extremas — do Thin (100) ao ExtraBold (800). A variação de peso constrói hierarquia onde outras marcas usariam múltiplas famílias. Limpa, legível, brasileira sem ser "tropical".

### Hierarchy
- **Display** (800, clamp(2.25rem, 5vw, 3.75rem), line-height 1.1): Hero principal. Máximo 2 linhas. Tracking levemente negativo (-0.02em). Reservado para a headline principal da página.
- **Headline** (700, clamp(1.875rem, 3vw, 2.25rem), line-height 1.2): Títulos de seção. Tracking -0.01em. Usar `text-wrap: balance` para linhas equilibradas.
- **Title** (700, 1.25rem, line-height 1.3): Títulos de cards, nomes de planos, headings de formulário.
- **Body** (400, 1rem, line-height 1.6): Texto corrido. Linha máxima de 65-75ch. `text-wrap: pretty` em parágrafos longos.
- **Label** (500, 0.875rem, line-height 1.4): Labels de formulário, navegação, chips, badges. Sem letter-spacing adicional.

### Named Rules
**The Peso Único Rule.** Inter é a única família do sistema. Sem fontes secundárias para "variedade". A hierarquia vem de peso (400 → 500 → 700 → 800), tamanho (escala modular de razão 1.25), e cor — não de troca de fonte.

**The Display Floor Rule.** Tracking do Display nunca abaixo de -0.04em. O atual -0.02em é o ponto ideal para Inter; mais apertado que isso e as letras se tocam.

## 4. Elevation

O sistema é **levemente elevado**: superfícies são planas em repouso, com elevação sutil em elementos interativos. O header usa backdrop-blur para transmitir profundidade sem sombra. Cards ganham sombra apenas no hover — a elevação é uma resposta, não um estado permanente. Seções alternam cores de fundo (Branco Solar / Luz do Sol / Névoa Teal) para criar ritmo vertical sem depender de sombras.

### Shadow Vocabulary
- **hover-card** (`box-shadow: 0 4px 12px rgba(0,0,0,0.08)`): Elevação sutil ao passar o mouse em cards. Suficiente para destacar, insuficiente para flutuar.
- **focus-ring** (`box-shadow: 0 0 0 2px var(--color-solar-primary)`): Anel de foco visível em todos os elementos interativos. 2px sólido, sem blur — máximo de acessibilidade.
- **header** (`backdrop-filter: blur(12px)`): Header fixo com desfoque de fundo. Transmite camadas sem sombra — o conteúdo rola "atrás" do header.

### Named Rules
**The Flat-By-Default Rule.** Superfícies nascem planas. Sombra é sempre uma resposta a estado (hover, foco), nunca decoração. Cards em repouso não têm sombra — apenas borda.

**The No Ghost Card Rule.** Nunca combinar `border: 1px solid` com `box-shadow` de blur ≥16px no mesmo elemento. Borda OU sombra — nunca os dois como decoração.

## 5. Components

### Buttons

**Character:** Convidativos e confiáveis. Como interruptores físicos bem projetados — você sabe exatamente o que vai acontecer ao pressionar.

- **Shape:** Cantos arredondados (8px). Sem bordas decorativas.
- **Primary:** Fundo Verde Terra (#1E5631), texto Branco Solar (#ffffff), font-weight 700. Padding: 10px 24px.
- **Secondary:** Fundo Teal Serena (#A5C9CA), texto Terra Escura (#111827), font-weight 700.
- **Outline:** Fundo transparente, borda 2px Verde Terra, texto Verde Terra. Hover: fundo Verde Orvalho.
- **Hover (Primary/Secondary):** Transição de 200ms ease-in-out na cor de fundo. Sem transform.
- **Focus:** Anel de foco 2px na cor do botão, offset 2px. Visível em todos os estados.
- **Sizes:** sm (8px 16px, text-sm), md (10px 24px, text-base), lg (10px 24px, text-base).

### Cards

**Character:** Molduras de informação, não gavetas fechadas. Leves, abertos, com hierarquia interna clara.

- **Shape:** Cantos arredondados (8px). Borda 1px Linha do Horizonte.
- **Background:** Branco Solar. Highlight: borda Verde Terra + anel sutil (2px, 30% opacidade).
- **Padding:** 16px internos. Stack vertical com gap de 16px entre preço, features, CTA.
- **Hover:** Sombra hover-card (0 4px 12px rgba(0,0,0,0.08)). Transição de 200ms.
- **Badge:** Pill (rounded-full), fundo Verde Terra, texto Branco Solar, centralizado na borda superior.

### Inputs / Fields

**Character:** Convidativos ao toque. A borda escurece no foco; o campo "acorda" quando você interage.

- **Style:** Borda 1px Linha do Horizonte, fundo Branco Solar, cantos 8px.
- **Focus:** Borda muda para Verde Terra, anel de foco 2px Verde Terra com offset 2px.
- **Error:** Borda vermelha, mensagem de erro em vermelho abaixo do campo.
- **Disabled:** Opacidade 50%, cursor not-allowed.
- **Radio/Chip group:** Labels com borda, selecionado = fundo Verde Orvalho + borda Verde Terra + texto Verde Terra Profundo.

### Navigation

**Character:** Presente mas não intrusivo. O header é uma moldura, não um outdoor.

- **Header:** Fixo no topo, altura 64px. Fundo Branco Solar com 80% opacidade + backdrop-blur 12px. Borda inferior 1px Linha do Horizonte.
- **Links:** Sombra Suave (#444444) em repouso, Terra Escura (#111827) no hover. Transição de 150ms. Font-weight 500.
- **CTA no header:** Botão primary, tamanho sm. "Baixar App" como ação principal.
- **Mobile:** Menu full-screen com overlay, links em stack vertical, botão de fechar no topo.

### Calculator (Signature Component)

**Character:** O coração funcional do site. Duas calculadoras lado a lado (Consumidor / Fornecedor) que transformam números abstratos em economia real.

- **Shape:** Abas horizontais para seleção de perfil, campos numéricos com labels claras, resultado em destaque.
- **Result display:** Número grande (Display weight 800, Verde Terra ou Terra Escura), label explicativa abaixo.
- **Background:** Branco Solar com borda 1px Linha do Horizonte, cantos 8px.
- **Input:** Campo numérico com slider + input direto. Range color: Verde Terra.

## 6. Do's and Don'ts

### Do:
- **Do** usar Inter em toda a interface — uma família, hierarquia por peso e tamanho
- **Do** usar Verde Terra (#1E5631) como acento primário em ≤15% de qualquer tela
- **Do** alternar fundos de seção (Branco Solar → Luz do Sol → Névoa Teal) para ritmo visual
- **Do** usar `text-wrap: balance` em headlines e `text-wrap: pretty` em parágrafos longos
- **Do** mostrar números de economia em peso extrabold com Terra Escura
- **Do** incluir anel de foco visível (2px) em todo elemento interativo
- **Do** usar fotografia real de casas e pessoas brasileiras — nada de stock photos genéricas
- **Do** escrever alt text descritivo e em português natural

### Don't:
- **Don't** usar verde-limão, gradientes ecológicos, ou ícones de folha/gota — é o clichê genérico de energia
- **Don't** aplicar dark mode, neon, glassmorphism ou estética "startup tech fria"
- **Don't** usar azul marinho, brasões, bandeiras do Brasil ou estética de site governamental
- **Don't** combinar `border: 1px solid` com `box-shadow` de blur ≥16px no mesmo elemento (The No Ghost Card Rule)
- **Don't** usar `border-radius` acima de 16px em cards ou seções
- **Don't** colocar eyebrow (texto pequeno tracked uppercase) acima de toda seção — um kicker deliberado é voz; em toda seção é gramática de IA
- **Don't** usar `border-left` ou `border-right` > 1px como faixa colorida decorativa em cards ou listas
- **Don't** esconder números importantes em Sombra Suave (#444444) — números são sempre Terra Escura com peso bold
- **Don't** usar imagens genéricas de bancos de fotos internacionais — o Brasil real pede imagens reais
- **Don't** escrever em inglês ou "portunhol" — a voz é em português brasileiro natural
