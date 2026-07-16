---
title: "Gerenciamento de Estado — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
---

# Gerenciamento de Estado — Solar Fácil

## 1. Estratégia

O Solar Fácil divide o gerenciamento de estado em três camadas:

| Camada | Tecnologia | Escopo |
|---|---|---|
| Estado de UI Global | React Context API (4 contextos) | Tema, Autenticação, Database, React Query |
| Estado Assíncrono | TanStack React Query | Dados do SQLite (cache, loading, erro) |
| Estado Local | React useState/useEffect | Estado efêmero de componentes |

## 2. Contextos (Estado Global Síncrono)

### 2.1. AppThemeContext

```typescript
// Estado: tema claro/escuro
type AppThemeContextType = {
  theme: AppThemeStyles;   // Objeto com ~50 estilos
  toggleTheme: () => void; // Alterna entre temas
};

// Fonte de verdade: useColorScheme() (preferência do sistema)
// Persistência: Não (segue sistema a cada mount)
```

**Estrutura de estado:**
```
isDark: boolean (useState)
  → true → darkTheme (objeto completo)
  → false → lightTheme (objeto completo)
```

### 2.2. AuthContext

```typescript
// Estado: autenticação do usuário
type AuthContextType = {
  isLoggedIn: boolean;
  userID: number | null;
  userName: string | null;
  associado?: AssociadoType;
  login: (id, name, associado?) => void;
  logout: () => void;
  updatelogin: (associado) => void;
};

// Fonte de verdade: SQLite (validação CPF/CNPJ + senha)
// Persistência: NÃO — apenas em memória (perdido ao fechar)
```

**Estrutura de estado:**
```
isLoggedIn: false → login() → true → logout() → false
userID: null → login(id) → id → logout() → null
associado: undefined → login(assoc) → AssociadoType → logout() → undefined
```

### 2.3. DatabaseContext

```typescript
// Estado: conexão SQLite
type DatabaseContextType = {
  isDatabaseConnected: boolean;
  dbInstance: SQLiteDatabase | null;
  initializeDatabaseConnection: () => Promise<void>;
  finalizeDatabaseConnection: () => Promise<void>;
  getDatabaseConnection: () => SQLiteDatabase | null;
};

// Fonte de verdade: expo-sqlite
// Inicialização: automática (autoInitialize=true)
```

**Estrutura de estado:**
```
isDatabaseConnected: false
  → useEffect (autoInitialize=true)
    → initializeDatabaseConnection()
      → SQLite.openDatabaseSync("solarfacil.db")
      → CREATE TABLEs
    → setDbInstance(db)
    → setIsDatabaseConnected(true)
```

### 2.4. ReactQueryProvider

```typescript
// Configuração do QueryClient (cache global)
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 3,
      staleTime: 5 * 60 * 1000,
    },
  },
});
```

## 3. React Query (Estado Assíncrono)

### 3.1. Queries (Leitura)

```typescript
// Exemplo de Query Key structure:
["associados"]                    // Lista todos
["associados", id]                // Por ID
["associados", "cpf", cpf_cnpj]  // Por CPF/CNPJ + senha (login)
["movimentacoes"]                // Lista todas
["movimentacoes", associadoId]   // Por associado
```

### 3.2. Mutations (Escrita)

Cada mutation invalida queries relevantes no `onSuccess`:

```typescript
// Insert → invalida ["associados"]
// Update → invalida ["associados"] e ["associados", id]
// Delete → invalida ["associados"]
// Insert movimentacao → invalida ["movimentacoes", associadoId]
```

### 3.3. Cache & Stale Time

| Query | staleTime | gcTime |
|---|---|---|
| associados (lista) | 5 min | 30 min |
| associado (por ID) | 5 min | 30 min |
| movimentações | 2 min | 15 min |

## 4. Fluxo de Estado: Login → Home → Movimentações

```
[LoginScreen]
  useQueryAssociadosSearchByCpfCnpjSenha(cpf, senha)
    ↓ (dados do SQLite)
  AuthContext.login(userID, userName, associado)
    ↓ (isLoggedIn = true)

[AuthProtectedSlot]
  useAuth().isLoggedIn === true
    ↓ (tab muda para "Logout")

[MovimentacaoScreen]
  useQueryMovimentacoesSearchByAssociadoId(userID)
    ↓ (dados cacheados pelo React Query)
  Exibe cards + gráfico Victory
```

## 5. Boas Práticas Recomendadas

1. **Evitar prop drilling:** usar hooks customizados que consomem contextos
2. **Estado mínimo:** não duplicar estado — derivar do React Query cache
3. **Imutabilidade:** usar spread operator ou immer para updates complexos
4. **Memoização:** `useMemo` para dados derivados, `useCallback` para handlers passados a children
5. **Cleanup:** cancelar queries ao desmontar (`queryClient.cancelQueries` no useEffect cleanup)
