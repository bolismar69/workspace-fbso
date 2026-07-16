---
title: "Relatório de Performance Mobile — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
status: "code-only"
---

# Relatório de Performance Mobile — Solar Fácil

## 1. Bundle Size

[TODO]: Analisar com `react-native-bundle-visualizer` ou `expo-export`.

**Dependências de impacto:**

| Pacote | Tamanho Estimado | Impacto |
|---|---|---|
| react-native | ~8 MB (base) | Framework |
| expo | ~5 MB | Plataforma |
| victory-native | ~2 MB | Gráficos (maior dependência de UI) |
| react-native-reanimated | ~1 MB | Animações |
| moti | ~100 KB | Animações declarativas |
| nativewind | Compile-time (0 runtime) | — |

## 2. Render Performance

### Observações

| Aspecto | Status |
|---|---|
| Virtualização de listas | ❌ ScrollView em vez de FlatList/FlashList |
| React.memo | ❌ Não usado em componentes de lista |
| useMemo/useCallback | ⚠️ [RUNTIME] Verificar uso em hooks |
| Reanimated (UI Thread) | ✅ Instalado — uso verificado |
| Expo Image | ✅ Instalado — cache otimizado vs RN Image |

### Recomendações

- Migrar `AssociadoListaTodosScreen` para `FlashList` (Shopify)
- Migrar `MovimentacaoComCardScreen` para `FlashList`
- Adicionar `React.memo` em `AssociadoItem`
- Usar `getItemLayout` para altura fixa de itens de lista
- Configurar `windowSize` e `maxToRenderPerBatch` em FlatList

## 3. Startup Time

| Fase | Impacto |
|---|---|
| Splash Screen | ✅ Estático (rápido) |
| SQLite Init | ⚠️ `withExclusiveTransactionAsync` — [RUNTIME] Medir |
| Provider Tree | 5 providers aninhados — overhead inicial |
| React Query Init | ✅ Leve |

## 4. Consumo de Memória

[TODO]: Medir com Xcode Memory Debugger (iOS) / Android Profiler.

**Riscos identificados:**
- SQLite mantém conexão aberta (singleton `dbInstance`) — OK
- AsyncStorage mantido como legado junto com SQLite — consome memória duplicada
- Sem cleanup de listeners/timers [RUNTIME]

## 5. Otimizações Recomendadas

1. **Imediato:** Migrar ScrollView → FlashList em listas
2. **Curto prazo:** Adicionar React.memo em AssociadoItem
3. **Curto prazo:** Analisar bundle com `react-native-bundle-visualizer`
4. **Médio prazo:** Implementar lazy loading de tabs (Expo Router `lazy`)
5. **Médio prazo:** Remover AsyncStorage legado e consolidar em SQLite
