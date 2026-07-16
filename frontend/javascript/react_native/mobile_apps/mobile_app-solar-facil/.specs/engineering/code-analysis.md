---
title: "Análise de Código — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
---

# Análise de Código — Solar Fácil

## 1. Fluxo de Renderização

```
App.tsx → RootLayout (_layout.tsx)
  ├── ReactQueryProvider
  │   └── QueryClient (cache global)
  ├── AppThemeProvider
  │   └── useColorScheme() → lightTheme | darkTheme
  ├── AuthProvider
  │   └── useState: isLoggedIn, userID, userName, associado
  ├── SafeAreaProvider
  ├── DatabaseProvider (autoInitialize=true)
  │   └── useEffect → initializeDatabase() → SQLite.openDatabaseSync("solarfacil.db")
  │       ├── CREATE TABLE IF NOT EXISTS associados
  │       └── CREATE TABLE IF NOT EXISTS movimentacoes
  └── GestureHandlerRootView
      └── AuthProtectedSlot
          └── Tabs (8 rotas)
```

## 2. Fluxo de Dados: Cadastro de Associado

```
1. Usuário preenche formulário (CadastroScreen → FormCadastroAssociado)
2. react-hook-form gerencia estado dos campos via DynamicInput
3. yup schema valida em tempo real (onChange/onBlur)
4. onSubmit → useMutationAssociadoInsertRecord.mutate(data)
5. MutationFn → getDatabaseConnection() → db.runAsync(INSERT INTO associados ...)
6. onSuccess → queryClient.invalidateQueries(["associados"])
7. HomeScreen re-renderiza com novo associado na lista
```

## 3. Fluxo de Dados: Login

```
1. Usuário digita CPF/CNPJ + senha (LoginScreen)
2. useQueryAssociadosSearchByCpfCnpjSenha(cpf, senha)
3. QueryFn → db.getFirstAsync("SELECT * FROM associados WHERE cpf_cnpj=? AND senha=?")
4. Se encontrado → AuthContext.login(id, nome, associado)
    → isLoggedIn = true
    → Tab "Login" muda para "Logout"
5. Se não encontrado → [TODO] Mensagem de erro
```

## 4. Dependência entre Contextos

```
ReactQueryProvider (topo — sem dependências)
  └── AppThemeProvider (sem dependências)
      └── AuthProvider (sem dependências)
          └── SafeAreaProvider (sem dependências)
              └── DatabaseProvider (sem dependências)
                  └── AuthProtectedSlot (consome: useAuth, useColorScheme)
                      └── Tabs (consome: useAuth via AuthProtectedSlot)
```

**Ordem de inicialização:** Top-down. Cada provider renderiza seus children, então a ordem importa para disponibilidade de hooks dos providers pais.

## 5. Hooks Customizados

| Hook | Dependências | Retorno |
|---|---|---|
| `useFormValidation(schema)` | react-hook-form, yup | `{control, handleSubmit, errors}` |
| `useQueryAssociadosSearchAll()` | React Query, SQLite | `{data, isLoading, isError, error}` |
| `useMutationAssociadoInsertRecord()` | React Query, SQLite | `mutate, mutateAsync, isLoading, isError` |
| `useAuth()` | AuthContext | `{isLoggedIn, login, logout, userID, userName, associado}` |
| `useAppTheme()` | AppThemeContext | `{theme, toggleTheme}` |
| `useDatabase()` | DatabaseContext | `{isDatabaseConnected, dbInstance, ...}` |

## 6. Ciclo de Vida de Componentes

### DatabaseProvider
```
Mount → useEffect (autoInitialize=true)
  → initializeDatabaseConnection()
    → initializeDatabase() [do initializeSQLiteDatabase.ts]
      → SQLite.openDatabaseSync("solarfacil.db")
      → withExclusiveTransactionAsync (CREATE TABLE IF NOT EXISTS)
    → setDbInstance(db)
    → setIsDatabaseConnected(true)
Unmount → finalizeDatabaseConnection()
  → dbInstance.closeAsync()
```

### AuthProvider
```
Mount → useState: isLoggedIn=false, userID=null, userName=null
login(id, name, associado) → setState
logout() → setState (clear)
Render → AuthContext.Provider value={{isLoggedIn, login, logout, ...}}
```

## 7. Pontos de Atenção

- **Auth em memória:** ao fechar o app, `isLoggedIn` é perdido. Usuário precisa fazer login novamente.
- **Database singleton:** `dbInstance` é global no módulo `initializeSQLiteDatabase.ts` — thread-safe para acesso síncrono.
- **Sem cleanup de queries:** React Query gerencia cache, mas queries ativas não são canceladas ao desmontar componentes. [TODO]: Usar `AbortController` ou `queryClient.cancelQueries`.
- **Re-renders em cascata:** `AuthProtectedSlot` re-renderiza todas as tabs quando `isLoggedIn` muda — pode causar flicker.
