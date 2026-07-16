---
title: "Design System — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["design", "design-system", "mobile", "tokens"]
---

# Design System — Solar Fácil

## 1. Visão Geral

O design system do Solar Fácil é construído sobre **NativeWind 4** (Tailwind CSS para React Native), complementado por um sistema de temas (claro/escuro) definido programaticamente em TypeScript.

## 2. Princípios de Design

| Princípio | Descrição |
|---|---|
| **Mobile First** | Design otimizado para telas mobile (375-430pt de largura) |
| **Energia & Natureza** | Paleta de verdes e amarelos — remetendo a sustentabilidade e energia solar |
| **Alto Contraste** | Texto legível com contraste adequado para uso em ambientes externos (luz solar) |
| **Touch-Friendly** | Elementos interativos com dimensões adequadas para toque |
| **Consistente** | Design tokens centralizados no `tailwind.config.js` e temas TypeScript |

## 3. Design Tokens

### 3.1. Cores (Tailwind)

| Token | Hex | Uso |
|---|---|---|
| `primary` | `#1E5631` | Verde escuro — botões, títulos, elementos de destaque primário |
| `secondary` | `#A4DE02` | Verde lima — destaques secundários, badges |
| `accent` | `#FFD700` | Dourado — elementos de acentuação |
| `background` | `#F6F6F6` | Fundo base da aplicação |
| `neutral` | `#FFFFFF` | Branco — cards, superfícies |

### 3.2. Cores (Tema TypeScript)

| Token | Light | Dark | Uso |
|---|---|---|---|
| `primary` | `#1E5631` | [TODO] | Cor primária |
| `secondary` | `#A5C9CA` | [TODO] | Cor secundária |
| `links` | `#1E90FF` | [TODO] | Links |
| `backgroundColor` | `#ffffbf` | [TODO] | Fundo principal |
| `textColor` | `#000` | [TODO] | Cor do texto |

### 3.3. Status Colors

| Status | Background | Uso |
|---|---|---|
| Pago | `#e0f7e9` | Movimentações pagas |
| Atrasado | `#fbe9e7` | Movimentações atrasadas |
| Pendente | `#fffde7` | Movimentações pendentes |

### 3.4. Tipografia

| Nível | Font Size | Weight | Uso |
|---|---|---|---|
| Título (header) | 24px | Bold | Títulos de tela (headerTitleStyle) |
| Título de Seção | 20px | Bold | Subtítulos |
| Corpo | 16px | Regular | Texto geral |
| Label | 14px | Medium (500) | Labels de formulário |
| Tab Label | 12px | Bold | Rótulos da tab bar |
| Card Title | 16px | Semi-bold (600) | Títulos de cards |

### 3.5. Spacing

| Token | Valor | Uso |
|---|---|---|
| Padding padrão | 16px | Container, cards, telas |
| Margin entre seções | 24px | marginTop em subtítulos |
| Margin entre itens | 8px | marginBottom em textos e botões |
| Tab margin | 0 | Tabs não têm margem extra |

### 3.6. Border Radius

| Elemento | Valor |
|---|---|
| Botões | 8px |
| Cards | 8-12px |
| Inputs | 8px |
| Container de ícone | 28px (círculo: width=height=56, radius=28) |

### 3.7. Sombras (Elevation)

| Elemento | Propriedades |
|---|---|
| Cards | `shadowColor: #000`, `shadowOffset: {0,2}`, `shadowOpacity: 0.1-0.3`, `shadowRadius: 4`, `elevation: 2` |
| Botões (baixa altura) | `shadowOpacity: 0.25`, `shadowRadius: 3.84`, `elevation: 5` |

## 4. Estados da UI

| Estado | Descrição | Implementação |
|---|---|---|
| Default | Estado normal do componente | — |
| Focused | Elemento com foco (tab, input) | `tabBarActiveTintColor: #43A047`, `inputBorder` com borda |
| Disabled | [TODO] | — |
| Loading | Carregamento de dados | React Query `isLoading`; [TODO] Skeleton components |
| Error | Erro de validação ou sistema | `inputError` (borda vermelha + fundo cinza), [TODO] Error Boundary |
| Empty | Lista vazia | [TODO] — não padronizado |

## 5. Componentes Principais

### 5.1. Botões

| Variante | Classe/Style | Uso |
|---|---|---|
| Primary | `theme.button` (`bg-[#1E5631]`, texto branco, bold) | Ações principais |
| Primary Low | `theme.buttonLow` (mesmo bg, altura 32px, largura 25%) | Ações secundárias |
| Secondary | `theme.secondaryButton` (`bg-[#A5C9CA]`) | Ações alternativas |
| Link | `theme.linkButton` (`bg-[#1E90FF]`) | Navegação |

### 5.2. Cards

| Variante | Componente | Uso |
|---|---|---|
| Standard | `theme.card` / `theme.cardContent` | Cards genéricos |
| Ícone Amarelo | `CardIconeAmarelo` | Cards com ícone e fundo amarelo |
| Ícone Padrão | `CardIconePadrao` | Cards com ícone padrão |
| Plano | `CardPlan` | Cards de plano comercial |

### 5.3. Inputs

| Tipo | Componente | Comportamento |
|---|---|---|
| Text | `InputText` | Entrada de texto padrão |
| Date | `InputDate` | DateTimePicker nativo |
| Select | `InputSelect` | Picker nativo com opções |
| Radio | `InputRadio` | Seleção única |
| Switch | `InputSwitch` | Toggle booleano |
| Password | `InputPasswordWithToggle` | Senha com toggle de visibilidade |
| Textarea | `InputTextarea` | Texto multilinha |
| Dynamic | `DynamicInput` | Renderização condicional baseada em `FieldDefinitionType` |

## 6. Layout Patterns

### 6.1. Safe Areas

```tsx
// Sempre encapsular conteúdo em SafeAreaView
<SafeAreaProvider>
  <SafeAreaView style={theme.safe}>
    {/* conteúdo */}
  </SafeAreaView>
</SafeAreaProvider>
```

### 6.2. Keyboard Avoidance

```tsx
// KeyboardSafeScreen componente customizado
<KeyboardSafeScreen>
  {/* formulário */}
</KeyboardSafeScreen>
```

### 6.3. Scrollable Content

- Para listas curtas: `ScrollView`
- Para listas longas: [TODO] migrar para `FlatList`/`FlashList`

## 7. Animações & Transições

Ver `design/animations.md` para catálogo completo.

| Tipo | Tecnologia |
|---|---|
| Animações de entrada | Moti (`fadeIn`, `slideIn`) |
| Micro-interações | Haptics (`expo-haptics`) |
| Gestos | Reanimated (`react-native-reanimated`) |
| Splash Screen | Expo Splash Screen (estático) |

## 8. Responsividade

| Breakpoint | Dispositivo | Largura (pt) |
|---|---|---|
| Small Phone | iPhone SE | 375 |
| Standard Phone | iPhone 14 | 390 |
| Large Phone | iPhone Pro Max | 430 |
| Tablet | iPad | 744+ |

**Status:** A aplicação não implementa breakpoints responsivos — o layout é único e se adapta via flexbox. [TODO]: Testar e ajustar para tablets.

## 9. Ícones

| Biblioteca | Uso |
|---|---|
| `@expo/vector-icons` (Ionicons) | Tabs (Home, Saiba Mais, FAQ, Login, Cadastro, Lista) |
| `@expo/vector-icons` (MaterialIcons) | Tabs (Solar, Plano) |
| `@expo/vector-icons` (SimpleLineIcons) | Tab Movimentações |
| `expo-symbols` | SF Symbols (iOS apenas) |

### Tab Icons

| Tab | Ícone | Cor Focada |
|---|---|---|
| Solar | `MaterialIcons: solar-power` | `#43A047` |
| Plano | `MaterialIcons: currency-exchange` | `#43A047` |
| Saiba | `Ionicons: information-circle-outline` | `#43A047` |
| FAQ | `Ionicons: help-circle-outline` | `#43A047` |
| Login | `Ionicons: flash-outline / flash-off-outline` | `#43A047` |
| User | `Ionicons: person-outline` | `#43A047` |
| Mov | `SimpleLineIcons: book-open` | `#43A047` |
| Lista | `Ionicons: list-outline` | `#43A047` |

## 10. Barra de Tabs

| Propriedade | Valor |
|---|---|
| Layout | Bottom tabs (8 abas) |
| Header | Visível, título centralizado, 24px bold |
| Header Background | `#ffffbf` (light) |
| Tab Bar Background | Igual ao header |
| Active Color | `#43A047` (verde) |
| Inactive Color | `#888` (cinza) |
| Active Background | `#ffffbf` |
| Inactive Background | `#fff` |
| Label Size | 12px bold |
