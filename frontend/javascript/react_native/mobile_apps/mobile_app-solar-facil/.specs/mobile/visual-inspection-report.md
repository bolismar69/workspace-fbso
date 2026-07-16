---
title: "Inspeção Visual Mobile — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
status: "code-only"
note: "⚠️ Inspeção visual não realizada — simulador/device indisponível. Análise baseada apenas em código-fonte."
---

# Inspeção Visual Mobile — Solar Fácil

## ⚠️ Status: Não Executada

A inspeção visual requer o app rodando em simuladores/dispositivos para:
1. Capturar screenshots em múltiplas resoluções
2. Detectar problemas de layout visual
3. Verificar safe areas, notch, keyboard avoidance
4. Validar consistência entre iOS e Android

**Dispositivos planejados para inspeção:**

| Dispositivo | Resolução | Densidade | Status |
|---|---|---|---|
| iPhone SE (3rd gen) | 375×667 | @2x | ❌ Não testado |
| iPhone 14 | 390×844 | @3x | ❌ Não testado |
| iPhone 15 Pro Max | 430×932 | @3x | ❌ Não testado |
| iPad (10th gen) | 744×1133 | @2x | ❌ Não testado |
| Pixel 6a | 393×851 | 2.75x | ❌ Não testado |
| Pixel 7 Pro | 412×915 | 3.5x | ❌ Não testado |
| Samsung Galaxy Tab | 800×1280 | 2x | ❌ Não testado |

## Previsão de Problemas (Baseada em Análise de Código)

### Safe Areas & Notch

| Risco | Localização | Severidade |
|---|---|---|
| SafeAreaProvider configurado corretamente | `_layout.tsx` ✅ | — |
| Conteúdo em `cadastro.tsx` com `headerShown: false` — pode invadir notch/Dynamic Island | `_layout.tsx:227` | ⚠️ Média |
| `KeyboardSafeScreen` deve lidar com home indicator (iPhone) | `components/inputs/KeyboardSafeScreen.tsx` | ⚠️ [RUNTIME] |

### Keyboard Avoidance

| Risco | Localização |
|---|---|
| `KeyboardSafeScreen` implementado (KeyboardAvoidingView + ScrollView) | ✅ |
| Comportamento iOS vs Android diferente (padding vs height) | ⚠️ [RUNTIME] Verificar |

### Layout & Overflow

| Risco |
|---|
| Formulários com ~30 campos — possibilidade de overflow e scroll inadequado |
| 8 tabs na barra inferior — labels podem truncar em telas pequenas (iPhone SE: 375pt / 8 = ~47pt por tab) |
| `buttonLow` com `width: "25%"` — pode ser muito estreito em telas pequenas |

### Touch Targets

| Risco | Medida |
|---|---|
| `buttonLow` height: 32px | ❌ Abaixo do mínimo de 44pt (iOS) / 48dp (Android) |
| Tab icons: tamanho padrão do sistema | ⚠️ [RUNTIME] Verificar se ≥ 44pt |
| Inputs com padding 10-12px | ⚠️ Pode resultar em touch target < 44pt |

### Consistência Cross-Platform

| Elemento | iOS | Android | Risco |
|---|---|---|---|
| DateTimePicker | Spinner | Calendar dialog | ⚠️ UX diferente |
| Picker | Wheel | Dropdown | ⚠️ UX diferente |
| SF Symbols | ✅ `expo-symbols` | ❌ Não disponível | ⚠️ Fallback necessário |
| Shadow | `shadowColor/Offset/Opacity/Radius` | `elevation` | ✅ Implementado |
| StatusBar | `style: "auto"` | `edgeToEdgeEnabled: true` | ⚠️ [RUNTIME] Verificar |

### Estados da Aplicação

| Estado | Status |
|---|---|
| Splash Screen | ✅ Configurado |
| Loading | ⚠️ [RUNTIME] Sem Skeleton — verificar spinner/ActivityIndicator |
| Empty | ❌ Não implementado |
| Error | ❌ Não implementado na UI |
| Offline | N/A (app é offline-first) |

## Checklist para Inspeção Futura

- [ ] Executar `npx expo start --ios` e inspecionar no iPhone SE simulator
- [ ] Capturar screenshot de cada tela em 4 dispositivos
- [ ] Verificar safe areas em iPhone com notch/Dynamic Island
- [ ] Verificar keyboard avoidance em formulários
- [ ] Medir touch targets com Accessibility Inspector (iOS) / Accessibility Scanner (Android)
- [ ] Verificar contraste de cores em diferentes modos (claro/escuro)
- [ ] Testar font scaling 200% em ambas plataformas
- [ ] Verificar comportamento landscape vs. portrait
