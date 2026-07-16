# Design Tokens — Solar Fácil Site

> Tokens de design extraídos do código-fonte e do `.impeccable/design.json`.
> Gerado por `impeccable extract` + `documentation-writer` em 2026-07-08.

---

## 1. Paleta de Cores (OKLCH)

### Primary: Verde Terra

| Token | HEX | OKLCH | CSS Variable |
|---|---|---|---|
| `verde-terra-900` | `#0a1f10` | `oklch(18% 0.05 155)` | `--color-solar-primary-900` |
| `verde-terra-700` | `#153d23` | `oklch(28% 0.08 155)` | `--color-solar-primary-700` |
| `verde-terra-500` | `#1E5631` | `oklch(36% 0.08 155)` | `--color-solar-primary` |
| `verde-terra-300` | `#3a8a52` | `oklch(52% 0.10 155)` | `--color-solar-primary-300` |
| `verde-terra-100` | `#8fd99a` | `oklch(78% 0.08 155)` | `--color-solar-primary-100` |
| `verde-orvalho` | `#dcfce7` | `oklch(96% 0.03 155)` | `--color-solar-primary-bg` |

### Secondary: Teal

| Token | HEX | OKLCH | CSS Variable |
|---|---|---|---|
| `teal-profundo` | `#5a9a9a` | `oklch(68% 0.05 200)` | `--color-solar-secondary-700` |
| `teal-serena` | `#A5C9CA` | `oklch(78% 0.05 200)` | `--color-solar-secondary` |
| `nevoa-teal` | `#e6f4f4` | `oklch(95% 0.01 200)` | `--color-solar-secondary-bg` |

### Neutral

| Token | HEX | CSS Variable |
|---|---|---|
| `terra-escura` | `#111827` | `--color-solar-text-primary` |
| `sombra-suave` | `#444444` | `--color-solar-text-secondary` |
| `linha-do-horizonte` | `#cccccc` | `--color-solar-border` |
| `branco-solar` | `#ffffff` | `--color-solar-bg` |
| `luz-do-sol` | `#ffffbf` | `--color-solar-bg-alt` |

### Feedback

| Token | HEX | CSS Variable |
|---|---|---|
| `verde-confirma` | `#e0f7e9` | — |
| `amarelo-alerta` | `#fffde7` | — |
| `vermelho-atencao` | `#fbe9e7` | — |

---

## 2. Escala Tipográfica

```css
/* Display — Hero principal */
--font-display-size: clamp(2.25rem, 5vw, 3.75rem);
--font-display-weight: 800;
--font-display-line-height: 1.1;
--font-display-tracking: -0.02em;

/* Headline — Títulos de seção */
--font-headline-size: clamp(1.875rem, 3vw, 2.25rem);
--font-headline-weight: 700;
--font-headline-line-height: 1.2;
--font-headline-tracking: -0.01em;

/* Title — Cards, planos, formulários */
--font-title-size: 1.25rem;
--font-title-weight: 700;
--font-title-line-height: 1.3;

/* Body — Texto corrido */
--font-body-size: 1rem;
--font-body-weight: 400;
--font-body-line-height: 1.6;

/* Label — Formulários, navegação */
--font-label-size: 0.875rem;
--font-label-weight: 500;
--font-label-line-height: 1.4;
--font-label-tracking: normal;
```

---

## 3. Escala de Espaçamento

| Token | Valor | Uso |
|---|---|---|
| `section-y` | `clamp(4rem, 6vw, 6rem)` | Padding vertical de seções |
| `section-x` | `1rem` | Padding horizontal de seções |
| `card-padding` | `1rem` | Padding interno de cards |
| `button-sm` | `0.5rem 1rem` | Botão pequeno |
| `button-md` | `0.625rem 1.5rem` | Botão padrão |

---

## 4. Border Radius

| Token | Valor | Uso |
|---|---|---|
| `rounded-sm` | `8px` | Botões, inputs, cards |
| `rounded-full` | `9999px` | Badges, chips |

---

## 5. Sombras (Elevação)

| Token | Valor | Uso |
|---|---|---|
| `hover-card` | `0 4px 12px rgba(0,0,0,0.08)` | Hover de cards |
| `focus-ring` | `0 0 0 2px var(--color-solar-primary)` | Foco em interativos |

---

## 6. Implementação no Código

```css
/* src/app/globals.css */
@import "tailwindcss";

@theme inline {
  --color-solar-primary: #1E5631;
  --color-solar-secondary: #A5C9CA;
  --color-solar-bg: #ffffff;
  --color-solar-bg-alt: #ffffbf;
  --color-solar-text-primary: #111827;
  --color-solar-text-secondary: #444444;
  --color-solar-border: #cccccc;
}
```

---

Última atualização: 2026-07-08
