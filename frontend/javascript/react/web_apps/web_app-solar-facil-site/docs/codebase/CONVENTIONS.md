# CONVENTIONS — Solar Fácil Site

> Padrões de código, naming, formatação, error handling e imports.
> Gerado por `acquire-codebase-knowledge` em 2026-07-08.
> Fontes: análise do código-fonte, `eslint.config.mjs`, `prettier` config.

---

## 1. Naming Conventions

### 1.1 Arquivos

| Tipo | Convenção | Exemplo |
|---|---|---|
| Componentes React | PascalCase | `HeroSection.tsx`, `Button.tsx` |
| Hooks | camelCase, prefixo `use` | `useCalculator.ts`, `usePlans.ts` |
| Serviços | camelCase, prefixo `service` | `servicePlans.ts`, `serviceFAQs.ts` |
| Funções utilitárias | camelCase | `calculator.ts`, `validation.ts` |
| Tipos | camelCase | `types.ts`, `concessionaria.ts` |
| Constantes | camelCase | `constants.ts` |
| Mocks JSON | camelCase, prefixo `mock` | `mockPlans.json` |
| Páginas | `page.tsx` (convenção Next.js) | `app/page.tsx`, `app/planos/page.tsx` |

### 1.2 Código

| Elemento | Convenção | Exemplo |
|---|---|---|
| Funções | camelCase | `calculateConsumerEconomy()` |
| Tipos/Interfaces | PascalCase | `ConsumerResult`, `LeadForm` |
| Constantes | UPPER_SNAKE_CASE | `CONSUMER_DISCOUNT_RATE`, `PROVIDER_RATE` |
| Enums (union types) | PascalCase | `PlanName`, `PersonaProfile` |
| Variáveis locais | camelCase | `monthlyBill`, `validationErrors` |
| Estados React | camelCase | `isSubmitting`, `hasCalculated` |

## 2. Estrutura de Componentes

### 2.1 Template Padrão

```typescript
// src/components/{domain}/{ComponentName}.tsx

import { ExternalLib } from 'external-lib';
import { InternalComponent } from '@/components/shared/InternalComponent';
import type { SomeType } from '@/lib/types';

interface ComponentNameProps {
  prop1: string;
  prop2?: number;
}

export function ComponentName({ prop1, prop2 }: ComponentNameProps) {
  // 1. Hooks
  // 2. Derived state
  // 3. Event handlers
  // 4. Render
  return (
    <div>{/* JSX */}</div>
  );
}
```

### 2.2 Client Components

```typescript
'use client';

import { useState } from 'react';
// ...
```

A diretiva `'use client'` SEMPRE na primeira linha.

## 3. Imports

### 3.1 Ordem de Importação

1. Bibliotecas externas (`react`, `next`, `lucide-react`)
2. Componentes internos (`@/components/...`)
3. Hooks (`@/hooks/...`)
4. Lib/Services (`@/lib/...`, `@/services/...`)
5. Tipos (`@/lib/types`, `@/types/...`)

### 3.2 Path Alias

```typescript
// ✅ Usar path alias
import { Button } from '@/components/shared/Button';
import { usePlans } from '@/hooks/usePlans';

// ❌ Evitar paths relativos longos
import { Button } from '../../components/shared/Button';
```

Configurado em `tsconfig.json`: `"@/*": ["./src/*"]`

### 3.3 Type Imports

```typescript
// ✅ Usar import type para tipos
import type { Plan, ConsumerResult } from '@/lib/types';

// Tipos de terceiros também
import type { NextConfig } from 'next';
```

## 4. Formatação

### 4.1 Prettier

Configuração padrão Prettier + `prettier-plugin-tailwindcss` para ordenação de classes.

### 4.2 Tailwind Classes

Classes CSS são ordenadas automaticamente pelo plugin Prettier. A ordem recomendada:
1. Layout (display, position)
2. Sizing (width, height)
3. Spacing (padding, margin)
4. Visual (colors, borders, shadows)
5. Typography (font, text)
6. Effects (opacity, transform)

## 5. Error Handling

### 5.1 Serviços (Data Layer)

```typescript
// Padrão: try-catch com fallback para array vazio
export async function fetchPlans(): Promise<Plan[]> {
  return new Promise((resolve) => {
    try {
      setTimeout(() => resolve(PLANS), 500);
    } catch (error) {
      console.error("Erro ao buscar dados de planos => ", error);
      resolve([]); // Fallback seguro
    }
  });
}
```

### 5.2 Hooks (State Layer)

```typescript
// Padrão: try-catch com mensagem de erro em pt-BR
try {
  // ... operação
} catch {
  setSubmitError('Não foi possível enviar. Verifique sua conexão ou fale conosco pelo WhatsApp.');
}
```

### 5.3 Validação (Business Logic)

```typescript
// Padrão: funções puras retornam string | null
export function validateEmail(email: string): string | null {
  if (!email || email.trim().length === 0) {
    return 'E-mail é obrigatório';
  }
  // ...
  return null; // Sem erro
}
```

## 6. Estados de UI

Cada componente interativo implementa estes estados:

| Estado | Implementação |
|---|---|
| **Default** | Renderização inicial |
| **Loading** | `<Skeleton />` ou flags `isSubmitting`/`isLoading` |
| **Empty** | Mensagem contextual (ex: "Nenhum plano disponível") |
| **Error** | Mensagem de erro em pt-BR, sem technical jargon |
| **Success** | Tela de confirmação (ex: `SuccessScreen`) |
| **Disabled** | Botões disabled durante submissão (`isSubmitting`) |

## 7. TypeScript Strict

- `strict: true` em `tsconfig.json`
- Todas as props de componentes têm interface explícita
- Todos os hooks retornam tipo explícito (implícito via inferência)
- Funções utilitárias têm parâmetros e retorno tipados

## 8. Comentários

- `//` para comentários inline
- `/** JSDoc */` para funções exportadas de `lib/`
- Referências a requisitos: `// Refs: BR-XXX-001`
- Placeholders marcados: `// Placeholder — atualizar com número real`

## 9. Variáveis de Ambiente

```bash
# .env.example
NEXT_PUBLIC_GA_ID=G-XXXXXXXXXX          # Google Analytics 4
NEXT_PUBLIC_FORM_ENDPOINT=https://...   # Endpoint do formulário (Formspree)
```

- Prefixo `NEXT_PUBLIC_` para variáveis expostas ao cliente
- Valores default seguros no código (fallback)
- Placeholders claramente identificados

## 10. Anti-Padrões Identificados

| Anti-Padrão | Status | Recomendação |
|---|---|---|
| Duplicação de dados (`PLANS` vs `mockPlans.json`) | ⚠️ Ativo | Escolher uma fonte da verdade |
| Placeholders em produção (`WHATSAPP_NUMBER`, `APP_STORE_URL`) | ⚠️ Ativo | Substituir por valores reais |
| Zero testes automatizados | ⚠️ Ativo | Adicionar Vitest + Testing Library |
| `console.error` em produção | ⚠️ Ativo | Substituir por serviço de logging |

---

Última atualização: 2026-07-08
