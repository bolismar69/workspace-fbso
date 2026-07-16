---
title: "Diretrizes de Performance — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
---

# Diretrizes de Performance — Solar Fácil

## 1. Renderização

### 1.1. Virtualização de Listas

**❌ Atual:**
```tsx
<ScrollView>
  {associados.map(a => <AssociadoItem key={a.id} associado={a} />)}
</ScrollView>
```

**✅ Recomendado:**
```tsx
import { FlashList } from "@shopify/flash-list";

<FlashList
  data={associados}
  renderItem={({ item }) => <AssociadoItem associado={item} />}
  keyExtractor={(item) => item.id.toString()}
  estimatedItemSize={80}
  getItemType={(item) => item.status} // Diferentes layouts por status
/>
```

### 1.2. Memoização de Componentes

```tsx
// Evitar re-renders desnecessários em listas
const AssociadoItem = React.memo(({ associado, onPress }: Props) => {
  // ...
}, (prev, next) => prev.associado.id === next.associado.id);
```

### 1.3. Evitar Cálculos em Render

```tsx
// ❌ Cálculo a cada render
const total = movimentacoes.reduce((acc, m) => acc + m.valorTotal, 0);

// ✅ Memorizar com useMemo
const total = useMemo(
  () => movimentacoes.reduce((acc, m) => acc + m.valorTotal, 0),
  [movimentacoes]
);
```

## 2. Animações (Reanimated)

### 2.1. Animações na UI Thread

```tsx
// ✅ Animações Reanimated rodam na UI thread (60fps)
import Animated, { useSharedValue, useAnimatedStyle, withSpring } from "react-native-reanimated";

const scale = useSharedValue(1);
const animatedStyle = useAnimatedStyle(() => ({
  transform: [{ scale: scale.value }],
}));

const onPress = () => {
  scale.value = withSpring(0.95, {}, () => {
    scale.value = withSpring(1);
  });
};
```

### 2.2. Evitar setState em Animações

```tsx
// ❌ Força re-render na JS thread
const [scale, setScale] = useState(1);

// ✅ Usa UI thread — sem re-render
const scale = useSharedValue(1);
```

## 3. Imagens

### 3.1. Otimização

```tsx
// ✅ Expo Image com cache otimizado
import { Image } from "expo-image";

<Image
  source={imageUrl}
  contentFit="cover"
  transition={200}
  cachePolicy="memory-disk"
/>
```

### 3.2. Tamanhos Apropriados

- Ícones: usar `@expo/vector-icons` (fontes vetoriais) em vez de PNGs
- Imagens de tela cheia: máximo 2× a largura da tela em pixels
- Splash: 200px (já configurado)

## 4. Bundle Size

### 4.1. Dependências a Monitorar

| Dependência | Peso | Alternativa |
|---|---|---|
| victory-native | ~2 MB | `react-native-svg-charts` (mais leve) ou canvas |
| moti | ~100 KB | Remover se não usado ativamente |

### 4.2. Importações

```tsx
// ❌ Importa tudo
import * as Icons from "@expo/vector-icons";

// ✅ Importa apenas o necessário
import { Ionicons, MaterialIcons } from "@expo/vector-icons";
```

## 5. Banco de Dados

### 5.1. Conexão Única

```typescript
// ✅ Singleton — uma conexão para toda a app
let dbInstance: SQLiteDatabase | null = null;

// ❌ Evitar abrir/fechar a cada operação
// O hook useAssociados atual faz isso — refatorar
```

### 5.2. Prepared Statements

```typescript
// ✅ Reutilizar statements
const statement = await db.prepareAsync(
  "SELECT * FROM associados WHERE id = ?"
);
// Reutilizar statement múltiplas vezes
await statement.executeAsync([id1]);
await statement.executeAsync([id2]);
await statement.finalizeAsync();
```

## 6. Startup

### 6.1. Inicialização Assíncrona

- Splash screen cobre inicialização do SQLite (já configurado)
- [TODO]: Adicionar lazy loading para telas não-visíveis inicialmente
- [TODO]: Carregar dados mock apenas quando necessário (não no bundle inicial)

### 6.2. Provider Tree

A árvore atual tem 5 providers aninhados — overhead aceitável, mas evitar adicionar mais.

## 7. Checklist de Performance

- [ ] Migrar ScrollView → FlashList para listas
- [ ] Adicionar React.memo em AssociadoItem
- [ ] Adicionar useMemo para dados derivados
- [ ] Verificar uso de Reanimated vs Animated padrão
- [ ] Analisar bundle size (`react-native-bundle-visualizer`)
- [ ] Remover moti se não utilizado
- [ ] Implementar lazy loading de tabs
- [ ] Otimizar imagens (máximo 2× resolução da tela)
- [ ] Unificar conexão SQLite (evitar abrir/fechar repetidamente)
