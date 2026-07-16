---
title: "Animações & Transições — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
---

# Animações & Transições — Solar Fácil

## 1. Stack de Animação

| Biblioteca | Versão | Propósito |
|---|---|---|
| react-native-reanimated | ~3.17.4 | Animações de alta performance na UI thread |
| moti | ^0.30.0 | Animações declarativas (wrapper sobre Reanimated) |
| expo-haptics | ~14.1.4 | Feedback tátil (vibração) |
| expo-splash-screen | ~0.30.8 | Animação de splash screen (estático) |

## 2. Animações de Tela

### Splash Screen
- **Tipo:** Estático
- **Config:** `app.json` → `expo-splash-screen` plugin
- **Imagem:** `solar-facil-splash-icon.png` (200px, resizeMode: contain)
- **Fundo:** `#ffffff`

### Transição entre Tabs
- **Tipo:** Padrão do React Navigation (sem animação customizada)
- **Comportamento:** Troca imediata de tela (sem crossfade ou slide)

## 3. Micro-interações

### Botões (Potencial)
```tsx
// Moti — animação de escala no press
import { MotiView } from "moti";

<MotiView
  from={{ scale: 1 }}
  animate={{ scale: 0.95 }}
  transition={{ type: "timing", duration: 100 }}
>
  <ThemedButton title="Pressione" onPress={handler} />
</MotiView>
```

### Feedback Tátil
```tsx
import * as Haptics from "expo-haptics";

// Vibração leve em ações importantes
Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Light);
```

### Accordion (FAQ)
- **Componente:** `FaqAccordion`
- **Tipo:** Expansão/colapso de conteúdo (provavelmente usando `LayoutAnimation` ou `Animated`)

## 4. Catálogo de Animações Recomendadas

| Nome | Tipo | Tecnologia | Uso Sugerido |
|---|---|---|---|
| Fade In | Entrada | Moti `fadeIn` | Aparição de cards e listas |
| Slide In | Entrada | Moti `slideIn` | Transição de telas |
| Scale on Press | Feedback | Moti + Haptics | Botões e itens tocáveis |
| Skeleton Loading | Loading | Reanimated | Placeholder durante carregamento |
| Pull to Refresh | Gesto | ScrollView/FlatList | Atualização de listas |

## 5. Observações

- **Moti está instalado** (`^0.30.0`) mas seu uso no código atual não foi identificado — potencialmente não utilizado ou usado minimamente.
- **Reanimated** está instalado e configurado (plugin no Babel) — principal motor de animações.
- **Nenhuma animação de transição entre telas** está configurada atualmente.
- [TODO]: Implementar animações de skeleton loading para melhor UX durante carregamento de dados.
- [TODO]: Adicionar `LayoutAnimation` para expansão de accordion FAQ.
- [TODO]: Implementar pull-to-refresh em listas (movimentações, associados).
