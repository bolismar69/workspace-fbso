---
title: "Convenções de Código — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["conventions", "coding-standards", "mobile", "react-native"]
---

# Convenções de Código — Solar Fácil

## 1. Nomenclatura

### 1.1. Arquivos

| Tipo | Convenção | Exemplos |
|---|---|---|
| Componentes React | PascalCase `.tsx` | `CardPlan.tsx`, `InputText.tsx`, `ThemedButton.tsx` |
| Telas (Screens) | PascalCase com sufixo `Screen.tsx` | `HomeScreen.tsx`, `AssociadoCadastroScreen.tsx` |
| Hooks | camelCase com prefixo `use` | `useFormValidation.ts`, `useQueryAssociadosSearchAll.ts` |
| Serviços | camelCase com prefixo `service` | `serviceAssociado.ts`, `serviceFAQs.ts` |
| Tipos | PascalCase com sufixo `Type.ts` | `AssociadoType.ts`, `PlanType.ts` |
| Contextos | PascalCase com sufixo `Context.tsx` | `AuthContext.tsx`, `AppThemeContext.tsx` |
| Temas | camelCase com sufixo `Theme.ts` | `lightTheme.ts`, `darkTheme.ts` |
| Mocks | camelCase com prefixo `mock` | `mockFAQs.json`, `mockPlans.json` |
| Validadores | camelCase com prefixo `validator` | `validatorCPF.ts`, `validatorCNPJ.ts` |

### 1.2. Funções & Variáveis

- **Funções**: camelCase — `initializeDatabase()`, `getDatabaseConnection()`
- **Componentes**: PascalCase — `FormBeneficiado`, `InputSelect`
- **Constantes**: UPPER_SNAKE_CASE para valores fixos, camelCase para objetos
- **Interfaces/Type aliases**: PascalCase — `AssociadoType`, `AppThemeStyles`
- **Enums/Union types**: valores em PascalCase ou string literals — `"Em cadastro" | "Ativo" | "Inativo"`

### 1.3. Diretórios

| Convenção | Exemplos |
|---|---|
| Domínio de negócio | `associado/`, `beneficiado/`, `fornecedor/`, `movimentacao/` |
| Função técnica | `components/`, `hooks/`, `services/`, `types/` |
| Subpastas por tipo | `components/forms/`, `components/inputs/`, `components/lists/` |
| Hooks por operação | `hooks/queries/`, `hooks/mutations/` |

## 2. Organização de Imports

### 2.1. Ordem

1. Bibliotecas externas (React, Expo, React Native)
2. Bibliotecas de terceiros (react-hook-form, yup, axios)
3. Contextos (`@/context/*`)
4. Hooks (`@/hooks/*`)
5. Componentes (`@/components/*`)
6. Serviços (`@/services/*`)
7. Tipos (`@/types/*`)
8. Utilitários (`@/utils/*`)
9. Estilos (`@/styles/*`)

### 2.2. Path Aliases

```typescript
// tsconfig.json paths
"@/*": ["src/*"]
```

## 3. Estilização

### 3.1. Método Principal

**NativeWind (Tailwind CSS)** para estilização utilitária, complementado por estilos definidos nos temas.

```tsx
// Classes Tailwind via NativeWind
<View className="flex-1 bg-background p-4">
  <Text className="text-primary font-bold text-xl">Título</Text>
</View>
```

### 3.2. Temas

Estilos complexos (cards, formulários, status) são definidos nos objetos `lightTheme` e `darkTheme` e acessados via `useAppTheme()`:

```tsx
const { theme } = useAppTheme();
<View style={theme.card}>
  <Text style={theme.cardTitle}>Título do Card</Text>
</View>
```

## 4. Formulários

### 4.1. Padrão

```tsx
// 1. Schema Yup
const schema = yup.object({
  nome: yup.string().required("Nome obrigatório"),
  email: yup.string().email("Email inválido").required("Email obrigatório"),
});

// 2. Hook useFormValidation
const { control, handleSubmit, errors } = useFormValidation(schema);

// 3. onSubmit
const onSubmit = (data) => { /* lógica */ };

// 4. Render com DynamicInput
<DynamicInput name="nome" control={control} errors={errors} />
```

## 5. Acesso a Dados

### 5.1. Padrão React Query (recomendado)

```tsx
// Query (leitura)
const { data, isLoading, error } = useQueryAssociadosSearchAll();

// Mutation (escrita)
const mutation = useMutationAssociadoInsertRecord();
mutation.mutate(associadoData);
```

### 5.2. Serviço Direto (legado)

```tsx
const { insertRecord, initialize } = await useAssociados();
await initialize();
await insertRecord(associado);
```

## 6. TypeScript

### 6.1. Configuração

- `strict: true`
- `jsx: "react-native"`
- Path aliases: `@/* → src/*`
- Typed routes: `experiments.typedRoutes: true`

### 6.2. Tipos Específicos de Mobile

```typescript
import { ImageStyle, TextStyle, ViewStyle } from "react-native";

// Estilos tipados
export interface AppThemeStyles {
  container: ViewStyle;
  title: TextStyle;
  imagePreview: ImageStyle;
}
```

## 7. Git

### 7.1. Branches

- `main` — branch principal
- `feature/*` — novas funcionalidades
- `fix/*` — correções

### 7.2. Commits (Conventional Commits)

```
feat: adiciona tela de cadastro de beneficiado
fix: corrige validação de CPF no formulário
chore: atualiza dependências do Expo
```

## 8. Plataforma

### 8.1. Código Específico de Plataforma

Usar `Platform.OS` ou extensões de arquivo (`.ios.ts`, `.android.ts`) quando necessário:

```typescript
import { Platform } from "react-native";

const isIOS = Platform.OS === "ios";
const isAndroid = Platform.OS === "android";
```

### 8.2. Safe Areas

Sempre usar `SafeAreaProvider` + `SafeAreaView` para conteúdo que fica próximo a bordas:

```tsx
import { SafeAreaProvider } from "react-native-safe-area-context";
```

## 9. Performance

- Preferir `FlatList` a `ScrollView` para listas longas
- Usar `React.memo()` para componentes puros
- Evitar renderizações desnecessárias em tabs (usar `lazy` quando disponível)
- Expo Image (`expo-image`) para imagens otimizadas em vez de `Image` do React Native

## 10. Segurança

- Senhas armazenadas no SQLite local (não criptografadas — [TODO] usar hash + salt)
- Autenticação local (sem backend remoto)
- [TODO]: SSL pinning para comunicações remotas
- [TODO]: armazenamento seguro de tokens no Keychain/Keystore

## 11. ESLint

Configuração: `eslint.config.js` com `eslint-config-expo` + TypeScript ESLint.

```bash
npx eslint . --ext .js,.jsx,.ts,.tsx
```
