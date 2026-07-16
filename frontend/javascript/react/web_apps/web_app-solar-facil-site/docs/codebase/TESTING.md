# TESTING — Solar Fácil Site

> Frameworks de teste, organização, estratégia de mocking e cobertura.
> Gerado por `acquire-codebase-knowledge` em 2026-07-08.
> Fontes: `package.json`, análise do código-fonte.

---

## 1. Status Atual: ⚠️ ZERO TESTES

**Não há nenhum teste automatizado no projeto.**

- ❌ Nenhum framework de teste instalado (Jest, Vitest, Playwright)
- ❌ Nenhum arquivo `*.test.ts` ou `*.spec.ts`
- ❌ Nenhum diretório `__tests__/`
- ❌ Nenhum script de teste no `package.json`

## 2. Cobertura de Testes por Camada

| Camada | Cobertura | Risco |
|---|---|---|
| `lib/calculator.ts` | 0% | 🔴 Alto — lógica de cálculo sem verificação automatizada |
| `lib/validation.ts` | 0% | 🔴 Alto — validação de formulário sem testes |
| `lib/constants.ts` | 0% | 🟡 Médio — dados estáticos |
| `hooks/useCalculator.ts` | 0% | 🔴 Alto — hook principal sem testes |
| `hooks/useContactForm.ts` | 0% | 🔴 Alto — formulário sem testes |
| `components/` | 0% | 🟡 Médio — 25 componentes sem smoke tests |
| `services/` | 0% | 🟢 Baixo — apenas wrappers com delay |
| **Total** | **0%** | 🔴 |

## 3. Estratégia de Teste Recomendada

### 3.1 Testes Unitários (Vitest)

**Alvo**: Funções puras em `lib/`

```typescript
// ✅ Testáveis isoladamente (funções puras, sem side effects)
- calculator.ts: calculateConsumerEconomy, calculateProviderGain, suggestPlan, formatBRL
- validation.ts: validateName, validateEmail, validatePhone, validateProfile, validateMessage, validateForm
```

### 3.2 Testes de Hook (Vitest + React Testing Library)

**Alvo**: Custom hooks em `hooks/`

```typescript
// ✅ Testáveis com renderHook
- useCalculator: input → calculate → result/error
- useContactForm: fill fields → validate → submit
- useFaqAccordion: toggle → open/close state
```

### 3.3 Testes de Componente (Vitest + Testing Library)

**Alvo**: Componentes individuais

```typescript
// ✅ Testes de renderização e interação
- ConsumerCalculator: render → input value → click calcular → resultado
- ContactForm: render → fill fields → submit → success/error state
- PlanCard: render → verificar props → click CTA
```

### 3.4 Testes E2E (Playwright)

**Alvo**: Fluxos completos de usuário

```typescript
// ✅ Fluxos críticos
- Fluxo Consumidor: Home → Calcular economia → Ver plano → Contato
- Fluxo Fornecedor: Home → Calcular ganho → Ver planos → Contato
- Fluxo Navegação: Home → Planos → Contato → Home
- Fluxo Erro: Formulário vazio → submit → validação → erros
```

## 4. Frameworks Recomendados

| Framework | Propósito | Justificativa |
|---|---|---|
| **Vitest** | Testes unitários + hooks + componentes | Rápido, compatível com Vite/Next.js, mesma sintaxe do Jest |
| **@testing-library/react** | Renderização de componentes React | Padrão da indústria, foco em comportamento do usuário |
| **@testing-library/jest-dom** | Matchers de DOM | `toBeInTheDocument()`, `toHaveTextContent()` |
| **Playwright** | Testes E2E | Multi-browser, screenshots, traces, já temos Playwright MCP |

## 5. Scripts a Adicionar no package.json

```json
{
  "scripts": {
    "test": "vitest",
    "test:watch": "vitest --watch",
    "test:coverage": "vitest --coverage",
    "test:e2e": "playwright test"
  }
}
```

## 6. Dependências a Adicionar

```json
{
  "devDependencies": {
    "vitest": "^3",
    "@testing-library/react": "^16",
    "@testing-library/jest-dom": "^6",
    "@testing-library/user-event": "^14",
    "@playwright/test": "^1"
  }
}
```

## 7. Estrutura de Testes Recomendada

```
src/
├── __tests__/              ← Testes unitários (espelham src/)
│   ├── lib/
│   │   ├── calculator.test.ts
│   │   ├── validation.test.ts
│   │   └── constants.test.ts
│   └── hooks/
│       ├── useCalculator.test.ts
│       └── useContactForm.test.ts
│
├── components/
│   └── __tests__/           ← Testes de componente (junto ao componente)
│       ├── home/
│       │   └── ConsumerCalculator.test.tsx
│       └── contact/
│           └── ContactForm.test.tsx
│
e2e/                         ← Testes E2E (raiz do projeto)
├── flows/
│   ├── consumer-journey.spec.ts
│   ├── provider-journey.spec.ts
│   └── navigation.spec.ts
└── fixtures/
    └── test-data.ts
```

---

Última atualização: 2026-07-08
