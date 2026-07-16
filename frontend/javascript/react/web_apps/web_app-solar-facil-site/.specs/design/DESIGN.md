# DESIGN.md — Solar Fácil Site

> Documentação do Design System. Gerado por `impeccable document` + análise de código.
> Fonte primária: `.impeccable/design.json` (2026-07-05), `.specs/DESIGN.md` original.
> Atualizado por `architecture-designer` em 2026-07-08.

---

## 1. North Star: "A Casa Conectada"

A Solar Fácil é o encontro entre o conforto do lar brasileiro e a inteligência de uma rede de energia moderna. Como uma casa bem projetada, cada seção do site é um "cômodo" com sua própria atmosfera — a sala de estar convida à descoberta, o escritório oferece números precisos, a varanda acolhe o contato.

**Características-chave:**
- Paleta verde-terra ancorada em confiança, não em "ecologia de marketing"
- Tipografia única (Inter) com variação expressiva de pesos para hierarquia
- Componentes que respondem ao toque — hover com elevação, foco visível, transições rápidas
- Ritmo visual variado entre seções — cada "cômodo" tem sua própria atmosfera
- Transparência como linguagem visual: números grandes, claros, verificáveis

---

## 2. Design Tokens

### 2.1 Cores

| Token | HEX | Papel | Uso |
|---|---|---|---|
| `verde-terra` | `#1E5631` | Primary | Botões primários, headlines destaque, badges |
| `verde-terra-profundo` | `#153d23` | Primary hover | Hover de botões, links ativos |
| `verde-orvalho` | `#dcfce7` | Primary bg | Fundo de chips selecionados, badges suaves |
| `teal-serena` | `#A5C9CA` | Secondary | Botões secundários, ícones, ilustrações |
| `teal-profundo` | `#7ab5b5` | Secondary hover | Hover de botões secondary |
| `nevoa-teal` | `#e6f4f4` | Secondary bg | Fundo alternado de seções |
| `branco-solar` | `#ffffff` | Neutral | Fundo principal, cards |
| `luz-do-sol` | `#ffffbf` | Neutral alt | Fundo alternado (máx 30% das seções) |
| `terra-escura` | `#111827` | Text primary | Texto principal, números de economia |
| `sombra-suave` | `#444444` | Text secondary | Texto secundário, labels |
| `linha-do-horizonte` | `#cccccc` | Border | Bordas de cards e inputs |

### 2.2 Tipografia

| Nível | Font Size | Weight | Line Height | Uso |
|---|---|---|---|---|
| **Display** | `clamp(2.25rem, 5vw, 3.75rem)` | 800 | 1.1 | Hero principal (máx 2 linhas) |
| **Headline** | `clamp(1.875rem, 3vw, 2.25rem)` | 700 | 1.2 | Títulos de seção |
| **Title** | `1.25rem` | 700 | 1.3 | Títulos de cards e planos |
| **Body** | `1rem` | 400 | 1.6 | Texto corrido (65-75ch max) |
| **Label** | `0.875rem` | 500 | 1.4 | Labels, navegação, chips |

### 2.3 Elevação

| Token | Valor | Uso |
|---|---|---|
| `hover-card` | `0 4px 12px rgba(0,0,0,0.08)` | Hover em cards |
| `focus-ring` | `0 0 0 2px var(--color-solar-primary)` | Anel de foco em interativos |

### 2.4 Movimento

| Token | Valor | Uso |
|---|---|---|
| `ease-standard` | `cubic-bezier(0.4, 0, 0.2, 1)` | Transições padrão |
| `ease-out-quart` | `cubic-bezier(0.25, 1, 0.5, 1)` | Entradas e revelações |
| `duration-fast` | `150ms` | Hover em links, ícones |
| `duration-standard` | `200ms` | Hover em botões, cards |
| `duration-slow` | `400ms` | Entradas de seção |

### 2.5 Breakpoints

| Nome | Largura |
|---|---|
| `sm` | 640px |
| `md` | 768px |
| `lg` | 1024px |
| `xl` | 1280px |

---

## 3. Componentes

### 3.1 Button (3 variantes)

| Variante | Background | Texto | Borda | Uso |
|---|---|---|---|---|
| **Primary** | `#1E5631` | Branco, bold | Nenhuma | Ação principal (máx 1 por seção) |
| **Secondary** | `#A5C9CA` | `#111827`, bold | Nenhuma | Ação secundária/alternativa |
| **Outline** | Transparente | `#1E5631`, bold | `2px solid #1E5631` | Ação terciária (cards não-destacados) |

Estados: default, hover (darken), focus-visible (ring 2px), disabled (opacity 0.5).

### 3.2 Plan Card

Card de plano com: badge "⭐ Mais Popular" (plano destacado), título, preço/mês, capacidade kWh, lista de features, CTA button. Plano destacado tem `border-color: #1E5631` + `transform: scale(1.02)`.

### 3.3 Text Input

Campo com label, placeholder, estados: default (border `#cccccc`), focus (`border-color: #1E5631` + ring), error (`border-color: #ef4444`). Label inclui `*` vermelho para required.

### 3.4 Metric Card

Valor grande (`2.25rem`, weight 800, `#1E5631`) + label pequena (`0.875rem`, weight 500, `#444444`). Layout vertical centralizado.

### 3.5 Header Navigation

Header fixo (`position: fixed`, `z-index: 50`, `height: 64px`), backdrop-blur, border-bottom. Logo SVG + links de navegação + CTA button.

---

## 4. Regras de Design (Design Rules)

1. **The Uma Voz Rule**: Verde Terra ≤15% de qualquer tela. Sua raridade é seu peso.
2. **The Luz Natural Rule**: Luz do Sol (#ffffbf) em ≤30% das seções. Ritmo, não dominância.
3. **The Confiança Visível Rule**: Números usam Terra Escura (#111827) com bold/extrabold.
4. **The Peso Único Rule**: Inter é a única família. Hierarquia = peso + tamanho + cor.
5. **The Flat-By-Default Rule**: Sombras são resposta a estado, nunca decoração.
6. **The No Ghost Card Rule**: Borda OU sombra — nunca os dois como decoração no mesmo elemento.

---

## 5. Do's e Don'ts

### ✅ Do
- Usar Inter em toda a interface
- Usar Verde Terra como acento primário em ≤15% da tela
- Alternar fundos de seção para ritmo visual
- `text-wrap: balance` em headlines, `text-wrap: pretty` em parágrafos
- Mostrar números de economia em extrabold com Terra Escura
- Anel de foco visível (2px) em todo interativo
- Alt text descritivo em português natural

### ❌ Don't
- Verde-limão, gradientes ecológicos, ícones de folha/gota
- Dark mode, neon, glassmorphism, estética "startup tech fria"
- Azul marinho, brasões, bandeiras do Brasil
- `border + box-shadow` no mesmo elemento como decoração
- `border-radius > 16px` em cards/seções
- Eyebrow acima de toda seção
- Números escondidos em Sombra Suave
- Inglês ou "portunhol"

---

Última atualização: 2026-07-08
