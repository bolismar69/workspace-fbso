---
title: "Auditoria de Acessibilidade Mobile — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
status: "code-only"
note: "⚠️ Auditoria realizada apenas via análise estática de código. Testes com VoiceOver/TalkBack e ferramentas automatizadas requerem app em execução no simulador/device."
---

# Auditoria de Acessibilidade Mobile — Solar Fácil

## ⚠️ Status: Análise Estática (Code-Only)

Testes não realizados por indisponibilidade de simulador/device:
- ❌ VoiceOver (iOS) — teste manual de leitor de tela
- ❌ TalkBack (Android) — teste manual de leitor de tela
- ❌ Accessibility Inspector (Xcode) — auditoria automatizada iOS
- ❌ Accessibility Scanner (Google) — auditoria automatizada Android
- ❌ `eslint-plugin-react-native-a11y` — lint não configurado

## Sumário de Conformidade

| Diretriz | Status | Issues |
|---|---|---|
| WCAG 1.1.1 — Text Alternatives | ❌ | Imagens sem accessibilityLabel |
| WCAG 1.4.3 — Contrast (Minimum) | ⚠️ | [RUNTIME] Verificar |
| WCAG 2.5.5 — Target Size | ❌ | Botão 32px < 44pt mínimo |
| WCAG 4.1.2 — Name, Role, Value | ❌ | Componentes sem accessibility props |
| HIG A11y — iOS | ❌ | Sem suporte a VoiceOver, Dynamic Type |
| Material A11y — Android | ❌ | Sem suporte a TalkBack, Font Size |

## Violações Detalhadas

### 1. Labels e Roles de Acessibilidade

**Status:** ❌ Não implementado

Todos os componentes interativos não possuem propriedades de acessibilidade:

```tsx
// ❌ Atual
<ThemedButton title="Cadastrar" onPress={handleCadastro} />

// ✅ Correto
<ThemedButton
  title="Cadastrar"
  onPress={handleCadastro}
  accessibilityLabel="Cadastrar novo associado"
  accessibilityRole="button"
  accessibilityHint="Toque para abrir o formulário de cadastro"
/>
```

| Componente | accessibilityLabel | accessibilityRole | accessibilityHint |
|---|---|---|---|
| ThemedButton | ❌ | ❌ | ❌ |
| CardIconeAmarelo | ❌ | ❌ | ❌ |
| CardIconePadrao | ❌ | ❌ | ❌ |
| CardPlan | ❌ | ❌ | ❌ |
| FaqAccordion | ❌ | ❌ | ❌ |
| InputText | ❌ | ❌ | ❌ |
| InputDate | ❌ | ❌ | ❌ |
| InputSelect | ❌ | ❌ | ❌ |
| AssociadoItem | ❌ | ❌ | ❌ |

### 2. Touch Targets

| Elemento | Tamanho | Mínimo iOS | Mínimo Android | Resultado |
|---|---|---|---|---|
| `buttonLow` | 32px altura | 44pt | 48dp | ❌ Reprovado |
| Input text | ~44px (12 padding × 2 + 16 font ≈ 44) | 44pt | 48dp | ⚠️ No limite |
| Tab icons | [RUNTIME] | 44pt | 48dp | ⚠️ Verificar |

### 3. Contraste de Cores

| Combinação | Contraste Estimado | Mínimo AA | Resultado |
|---|---|---|---|
| `#1E5631` (primary) sobre `#FFFFFF` | ~8.1:1 | 4.5:1 | ✅ Passa |
| `#888` (inactive tab) sobre `#FFFFFF` | ~3.5:1 | 3:1 (large) | ⚠️ Limite |
| `#1E90FF` (links) sobre `#FFFFFF` | ~3.9:1 | 3:1 (large) | ⚠️ Limite |
| `#A4DE02` (secondary) sobre `#FFFFFF` | ~1.6:1 | 4.5:1 | ❌ Reprovado |
| `#A5C9CA` (theme secondary) sobre `#FFFFFF` | ~1.9:1 | 4.5:1 | ❌ Reprovado |

**⚠️ `secondary` (#A4DE02 e #A5C9CA) não passam no contraste mínimo sobre fundo branco.**

### 4. Font Scaling (Dynamic Type / Font Size)

| Aspecto | Status |
|---|---|
| `allowFontScaling` configurado | ❌ Não configurado em nenhum Text |
| Texto testado em 200% scale | ❌ [RUNTIME] |
| Layout adapta-se a fontes maiores | ⚠️ [RUNTIME] |

### 5. Leitores de Tela

| Aspecto | Status |
|---|---|
| Ordem de foco lógica | ⚠️ [RUNTIME] |
| Imagens decorativas `accessible={false}` | ❌ Não implementado |
| Agrupamentos (`accessibilityRole`) | ❌ Não implementado |
| Anúncios de mudança de tela | ❌ Não implementado |
| Estados customizados (`accessibilityState`) | ❌ Não implementado |

### 6. Redução de Movimento

| Aspecto | Status |
|---|---|
| `prefers-reduced-motion` verificado | ❌ Não implementado |
| Animações Reanimated/Moti desligam com a preferência | ❌ Não implementado |

### 7. Navegação Alternativa

| Aspecto | Status |
|---|---|
| Switch Control (iOS) / Switch Access (Android) | ❌ [RUNTIME] Não testado |
| Navegação por teclado externo | ❌ Não suportada |

## Mapa de Remediação

### Fase 1 — Imediato (P1)

```tsx
// 1. Adicionar accessibilityLabel + accessibilityRole em TODOS componentes interativos
// 2. Aumentar touch targets para ≥ 44pt
// 3. Substituir cores secondary com contraste insuficiente
```

### Fase 2 — Curto Prazo (P2)

```tsx
// 4. Configurar eslint-plugin-react-native-a11y
// 5. Implementar font scaling (allowFontScaling + testar)
// 6. Marcar imagens decorativas com accessible={false}
```

### Fase 3 — Médio Prazo (P3)

```tsx
// 7. Testar com VoiceOver e TalkBack em todas as telas
// 8. Implementar accessibilityState para componentes com estado
// 9. Implementar suporte a prefers-reduced-motion
// 10. Adicionar accessibilityHint para ações não-óbvias
```

## Guia Rápido: Accessibility Props Essenciais

```tsx
// Botão
<TouchableOpacity
  accessibilityLabel="Nome descritivo da ação"
  accessibilityRole="button"
  accessibilityHint="O que acontece ao tocar"
  accessibilityState={{ disabled: false }}
/>

// Input
<TextInput
  accessibilityLabel="Nome do campo"
  accessibilityHint="Formato esperado ou instrução"
/>

// Imagem decorativa
<Image
  accessible={false}
  accessibilityElementsHidden={true}
/>

// Imagem informativa
<Image
  accessibilityLabel="Descrição da imagem"
/>

// Lista
<FlatList
  accessibilityLabel="Lista de associados"
  accessibilityHint="Deslize para navegar pelos itens"
/>

// Switch/Toggle
<Switch
  accessibilityLabel="Nome da configuração"
  accessibilityRole="switch"
  accessibilityState={{ checked: isEnabled }}
/>

// Accordion
<TouchableOpacity
  accessibilityLabel={isOpen ? "Fechar" : "Abrir" + " pergunta: " + title}
  accessibilityRole="button"
  accessibilityState={{ expanded: isOpen }}
/>
```
