---
title: "Testes — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["testing", "mobile", "jest", "react-native"]
---

# Testes — Solar Fácil

## 1. Status Atual

**⚠️ Nenhum framework de teste está configurado no projeto.**

O `package.json` não inclui Jest, React Native Testing Library, Detox, Maestro ou qualquer outro framework de teste como dependência (nem mesmo em `devDependencies`).

## 2. Frameworks Recomendados

### 2.1. Testes Unitários e de Componente

| Ferramenta | Propósito |
|---|---|
| Jest | Test runner padrão para React Native |
| @testing-library/react-native (RNTL) | Testes de componente com foco no usuário |
| @testing-library/jest-native | Matchers customizados (toBeVisible, toHaveTextContent) |

**Instalação sugerida:**
```bash
npx expo install jest @testing-library/react-native @testing-library/jest-native --dev
```

### 2.2. Testes End-to-End (E2E)

| Ferramenta | Propósito | Plataforma |
|---|---|---|
| Detox | E2E gray box testing (Wix) | iOS + Android |
| Maestro | E2E testing com YAML flows (simples) | iOS + Android |

**Instalação sugerida (Maestro — mais simples):**
```bash
brew install maestro
```

### 2.3. Testes de Acessibilidade

| Ferramenta | Propósito |
|---|---|
| eslint-plugin-react-native-a11y | Lint de accessibility props |

## 3. Organização de Testes (Planejada)

```
src/
├── __tests__/
│   ├── components/
│   │   ├── CardPlan.test.tsx
│   │   ├── InputText.test.tsx
│   │   └── ThemedButton.test.tsx
│   ├── hooks/
│   │   ├── useFormValidation.test.ts
│   │   └── queries/useQueryAssociadosSearchAll.test.ts
│   ├── screens/
│   │   ├── HomeScreen.test.tsx
│   │   └── LoginScreen.test.tsx
│   └── services/
│       └── database/initializeSQLiteDatabase.test.ts
└── e2e/
    └── flows/
        ├── login-flow.yaml
        ├── cadastro-flow.yaml
        └── movimentacao-flow.yaml
```

## 4. Estratégia de Mocking

### 4.1. Banco de Dados

Para testes unitários, mockar `expo-sqlite`:

```typescript
jest.mock("expo-sqlite", () => ({
  openDatabaseSync: jest.fn(() => ({
    withExclusiveTransactionAsync: jest.fn(),
    prepareAsync: jest.fn(),
    closeAsync: jest.fn(),
  })),
}));
```

### 4.2. Contextos

Envolver componentes em providers mockados ou reais:

```typescript
const wrapper = ({ children }) => (
  <AppThemeProvider>
    <AuthProvider>
      {children}
    </AuthProvider>
  </AppThemeProvider>
);

render(<Component />, { wrapper });
```

### 4.3. React Query

Para testes de hooks, usar `wrapper` com `QueryClientProvider`:

```typescript
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient({
  defaultOptions: { queries: { retry: false } },
});

const wrapper = ({ children }) => (
  <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
);
```

## 5. Cobertura Mínima Recomendada

| Área | Cobertura Alvo |
|---|---|
| Validadores (CPF, CNPJ, RG) | 100% |
| Hooks de query/mutation | 80% |
| Componentes de input | 70% |
| Telas (screens) | 50% |
| Serviços de banco | 60% |

## 6. CI/CD (Planejado)

[TODO]:
- Rodar `npx jest --coverage` em PRs
- Rodar E2E em builds de preview
- Integrar com GitHub Actions

## 7. Checklist de Implementação

- [ ] Instalar Jest + RNTL
- [ ] Configurar `jest.config.js`
- [ ] Criar mocks para `expo-sqlite` e `AsyncStorage`
- [ ] Escrever testes para validadores (CPF, CNPJ, RG)
- [ ] Escrever testes para hooks de query
- [ ] Escrever testes para componentes de input
- [ ] Configurar Maestro ou Detox para E2E
- [ ] Adicionar script `test` ao `package.json`
- [ ] Configurar GitHub Actions para CI
