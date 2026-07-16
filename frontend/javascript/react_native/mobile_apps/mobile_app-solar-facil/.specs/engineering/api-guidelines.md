---
title: "Padrões de Consumo de Dados — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
---

# Padrões de Consumo de Dados — Solar Fácil

## 1. Visão Geral

O Solar Fácil opera em modo **offline-first** com SQLite local. Não há chamadas HTTP — todo acesso a dados é local. Este documento estabelece os padrões atuais e o plano de migração para quando um backend remoto estiver disponível.

## 2. Padrão Atual: Acesso a Dados Local

### 2.1. Queries (Leitura) — React Query

```typescript
// hooks/queries/useQueryAssociadosSearchAll.ts
import { useQuery } from "@tanstack/react-query";

export function useQueryAssociadosSearchAll() {
  return useQuery({
    queryKey: ["associados"],
    queryFn: async () => {
      const db = getDatabaseConnection();
      const result = await db.getAllAsync(
        "SELECT * FROM associados ORDER BY nome"
      );
      return result as AssociadoType[];
    },
  });
}
```

**Uso no componente:**
```tsx
const { data, isLoading, isError, error } = useQueryAssociadosSearchAll();
```

### 2.2. Mutations (Escrita) — React Query

```typescript
// hooks/mutations/useMutationAssociadoInsertRecord.ts
import { useMutation, useQueryClient } from "@tanstack/react-query";

export function useMutationAssociadoInsertRecord() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (associado: AssociadoType) => {
      const db = getDatabaseConnection();
      await db.runAsync(
        "INSERT INTO associados (...) VALUES (?, ?, ...)",
        [associado.nome, associado.email, ...]
      );
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["associados"] });
    },
  });
}
```

### 2.3. Serviços Mock (Dados de Referência)

```typescript
// services/serviceFAQs.ts
import mockFAQs from "@/mocks/mockFAQs.json";

export function getFAQs(): FAQType[] {
  return mockFAQs;
}
```

## 3. Padrão Futuro: API Backend

### 3.1. Estrutura de Serviço HTTP

```typescript
// services/api/client.ts
import axios from "axios";

const apiClient = axios.create({
  baseURL: process.env.EXPO_PUBLIC_API_URL,
  timeout: 10000,
  headers: { "Content-Type": "application/json" },
});

// Interceptor de auth
apiClient.interceptors.request.use((config) => {
  const token = getAuthToken(); // Keychain/Keystore
  if (token) config.headers.Authorization = `Bearer ${token}`;
  return config;
});
```

### 3.2. Offline-First com React Query

```typescript
// Configuração recomendada para React Query quando houver backend
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 5 * 60 * 1000, // 5 minutos (dados considerados frescos)
      gcTime: 30 * 60 * 1000, // 30 minutos (cache mantido)
      retry: 3,
      retryDelay: (attempt) => Math.min(1000 * 2 ** attempt, 30000),
      networkMode: "offlineFirst", // Prioriza cache offline
    },
    mutations: {
      networkMode: "offlineFirst",
    },
  },
});
```

### 3.3. Estratégia de Retry e Timeout

| Parâmetro | Valor | Justificativa |
|---|---|---|
| timeout | 10s | Mobile — rede instável |
| retry | 3 tentativas | Responsivo sem ser excessivo |
| backoff | Exponential (1s → 2s → 4s) | Evita sobrecarregar servidor |
| staleTime | 5 min | Dados de negócio mudam com baixa frequência |
| gcTime | 30 min | Permite uso offline prolongado |

## 4. Tratamento de Erros

### 4.1. Erros de Banco Local

```typescript
try {
  await db.runAsync(query, params);
} catch (error) {
  console.error("Database error:", error);
  // [TODO]: Mostrar erro ao usuário (Toast/SnackBar)
  throw error; // Propagar para React Query
}
```

### 4.2. Erros de Rede (Futuro)

```typescript
// Erros esperados:
// - Network Error → Offline (usar cache)
// - 401 → Renovar token ou redirecionar para login
// - 422 → Erro de validação (mostrar campos)
// - 500 → Erro interno (retry ou mensagem genérica)
```

## 5. Segurança

### 5.1. Armazenamento de Tokens

```typescript
// NUNCA armazenar tokens em AsyncStorage (plain text)
// Usar expo-secure-store:
import * as SecureStore from "expo-secure-store";

await SecureStore.setItemAsync("authToken", token);
const token = await SecureStore.getItemAsync("authToken");
```

### 5.2. SSL Pinning

[TODO]: Configurar SSL pinning para prevenir MITM attacks quando backend existir.

## 6. Estrutura de Hooks (Convenção)

```
hooks/
├── queries/
│   ├── index.ts                          # Barrels
│   ├── useQueryAssociadosSearchAll.ts     # GET /associados
│   ├── useQueryAssociadosSearchById.ts    # GET /associados/:id
│   └── useQueryMovimentacoesSearch*.ts
├── mutations/
│   ├── index.ts
│   ├── useMutationAssociadoInsertRecord.ts  # POST /associados
│   ├── useMutationAssociadoUpdateRecord.ts  # PUT /associados/:id
│   ├── useMutationAssociadoDeleteRecord.ts  # DELETE /associados/:id
│   └── useMutationMovimentacoes*.ts
└── useFormValidation.ts                    # Validação de formulários
```
