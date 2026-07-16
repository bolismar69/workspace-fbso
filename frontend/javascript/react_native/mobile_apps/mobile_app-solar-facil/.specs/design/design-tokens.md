---
title: "Design Tokens — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
---

# Design Tokens — Solar Fácil

Tokens extraídos diretamente do código-fonte: `tailwind.config.js`, `lightTheme.ts`, `darkTheme.ts`, e `_layout.tsx`.

## Cores — Tailwind

Extraído de `tailwind.config.js`.

| Token | Valor Hex | RGB |
|---|---|---|
| `primary` | `#1E5631` | `rgb(30, 86, 49)` |
| `secondary` | `#A4DE02` | `rgb(164, 222, 2)` |
| `accent` | `#FFD700` | `rgb(255, 215, 0)` |
| `background` | `#F6F6F6` | `rgb(246, 246, 246)` |
| `neutral` | `#FFFFFF` | `rgb(255, 255, 255)` |

## Cores — Tema Light

Extraído de `src/styles/lightTheme.ts`.

| Token | Valor | Uso |
|---|---|---|
| `primary` | `#1E5631` | Cor primária |
| `secondary` | `#A5C9CA` | Cor secundária (≠ Tailwind secondary) |
| `links` | `#1E90FF` | Links clicáveis |
| `backgroundColor` | `#ffffbf` | Fundo principal |
| `textColor` | `#000` | Texto padrão |
| Tab Background | `#c3c3c3` | Fundo da tab bar |
| Tab Active | `#1E90FF` | Indicador de tab ativa |
| Tab Inactive | `#888` | Cor de tab inativa |

## Font Sizes

| Token | Valor | Uso |
|---|---|---|
| Title | 24px | `theme.title` |
| Subtitle | 20px | `theme.subtitle` |
| Body | 16px | `theme.text` |
| Label | 14px | `theme.label` |
| Small | 12px | Tab labels |

## Spacing

| Token | Valor |
|---|---|
| Container Padding | 16px |
| Section Margin Top | 24px |
| Item Margin Bottom | 8px |
| Button Padding Vertical | 10px |
| Button Padding Horizontal | 24px |

## Border Radius

| Token | Valor |
|---|---|
| Buttons | 8px |
| Cards | 8px / 12px |
| Inputs | 8px |
| Icon Container | 28px (circle) |
| Picker | 8px |

## Shadows / Elevation

| Token | Valor |
|---|---|
| Card Shadow Color | `#000` |
| Card Shadow Offset | `{width: 0, height: 2}` |
| Card Shadow Opacity | `0.1` (card) / `0.3` (cardContent) |
| Card Shadow Radius | `4` |
| Card Elevation (Android) | `2` |
| Button Shadow Opacity | `0.25` |
| Button Shadow Radius | `3.84` |
| Button Elevation (Android) | `5` |

## Tab Bar

| Token | Valor |
|---|---|
| Header Title Size | 24px |
| Header Title Weight | Bold |
| Header Title Align | Center |
| Tab Label Size | 12px |
| Tab Label Weight | Bold |
| Tab Bar Border Top | 0 (removida) |
| Active Tint Color | `#43A047` |
| Inactive Tint Color | `#888` |
| Active Background | `#ffffbf` |
| Inactive Background | `#fff` |

## Status Colors

| Status | Background | Texto (default) |
|---|---|---|
| Pago | `#e0f7e9` | `#000` |
| Atrasado | `#fbe9e7` | `#000` |
| Pendente | `#fffde7` | `#000` |
