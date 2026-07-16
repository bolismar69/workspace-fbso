---
title: "Arquitetura de Código — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["architecture", "camadas", "padrões", "mobile"]
---

# Arquitetura de Código — Solar Fácil

## 1. Visão Geral

O Solar Fácil adota uma arquitetura **offline-first** com **SQLite local** como fonte primária de dados. A aplicação segue o padrão **Context + Hooks** do React, organizada em camadas bem definidas: Providers → Screens → Components → Hooks → Services.

## 2. Padrão Arquitetural

### 2.1. Camadas

```
┌──────────────────────────────────────┐
│            APP (Expo Router)          │  ← Entry point + File-based routing
├──────────────────────────────────────┤
│            PROVIDERS                  │  ← ReactQuery → AppTheme → Auth → Database
├──────────────────────────────────────┤
│            SCREENS                    │  ← Telas montadas a partir de componentes
├──────────────────────────────────────┤
│            COMPONENTS                 │  ← Componentes reutilizáveis (forms, inputs, cards)
├──────────────────────────────────────┤
│            HOOKS                      │  ← Lógica de negócio (queries, mutations, validação)
├──────────────────────────────────────┤
│            SERVICES                   │  ← Acesso a dados (SQLite, AsyncStorage, Mock)
├──────────────────────────────────────┤
│            TYPES / UTILS / STYLES     │  ← Infraestrutura compartilhada
└──────────────────────────────────────┘
```

### 2.2. Fluxo de Dados

```
User Action → Screen → Component → Hook (useQuery/useMutation) → Service → SQLite/AsyncStorage
                                                    ↓
                                              React Query Cache → Re-render Component
```

### 2.3. Injeção de Dependência via Context

A aplicação usa o padrão **Provider Composition**:

```tsx
<ReactQueryProvider>         // Cache e gerenciamento de estado assíncrono
  <AppThemeProvider>          // Tema claro/escuro
    <AuthProvider>            // Estado de autenticação
      <SafeAreaProvider>      // Safe areas
        <DatabaseProvider>    // Conexão SQLite (autoInitialize=true)
          {/* Rotas */}
        </DatabaseProvider>
      </SafeAreaProvider>
    </AuthProvider>
  </AppThemeProvider>
</ReactQueryProvider>
```

## 3. Estratégias de Dados

### 3.1. SQLite (Fonte Primária)

- Banco: `solarfacil.db`
- Tabelas: `associados` (30 colunas), `movimentacoes` (18 colunas)
- Acesso: `expo-sqlite` com `openDatabaseSync` e `withExclusiveTransactionAsync`
- Inicialização: automática via `DatabaseProvider` (autoInitialize=true)

### 3.2. AsyncStorage (Fallback/Legado)

- Usado como alternativa para persistência simples
- CRUDs completos em `services/storage/`

### 3.3. React Query (Cache Layer)

- Queries: leitura de dados com cache automático
- Mutations: escrita com invalidação de cache
- Provider via `ReactQueryProvider` (topo da árvore)

### 3.4. Mock Services

- Dados estáticos em `src/mocks/*.json`
- Serviços mock em `services/mock/` e `services/service*.ts`
- Usados para dados que não mudam (concessionárias, FAQs, planos)

## 4. Padrões de Código

### 4.1. Telas (Screens)

Cada tela é composta por:
- Um componente `*Screen.tsx` em `screens/{dominio}/`
- Consome hooks e contextos
- Renderiza componentes reutilizáveis

### 4.2. Componentes

- **Forms**: componentes de formulário completos com validação
- **Inputs**: inputs atômicos (Text, Date, Select, Radio, Switch, Textarea, Password)
- **Cards**: cards de apresentação (IconeAmarelo, IconePadrao, Plan)
- **UI**: botões temáticos, ícones, accordion

### 4.3. Hooks

- **Queries**: hooks `useQuery*` — leitura com cache React Query
- **Mutations**: hooks `useMutation*` — escrita com invalidação
- **Custom**: `useFormValidation` — ponte react-hook-form + yup

### 4.4. Serviços

- **Database**: acesso direto ao SQLite (raw SQL)
- **Storage**: abstração sobre AsyncStorage
- **Mock**: dados estáticos + simulação de latência

## 5. Navegação

- **Tipo**: File-based routing (Expo Router 5)
- **Estrutura**: 8 rotas em bottom tabs (`Tabs` component)
- **Auth Gate**: `AuthProtectedSlot` — renderização condicional baseada em `isLoggedIn`
- **Deep Links**: scheme `solar-facil://`

## 6. Temas

- **Light Theme**: fundo claro (#fffbf), texto escuro, verde primário (#1E5631)
- **Dark Theme**: fundo escuro, texto claro (definido em `styles/darkTheme.ts`)
- **Detecção**: `useColorScheme()` — segue preferência do sistema
- **Toggle**: disponível via `AppThemeContext.toggleTheme()`

## 7. Cross-Cutting Concerns

### 7.1. Autenticação

- Estado gerenciado via `AuthContext`
- Login: associa `userID`, `userName` e `AssociadoType` ao contexto
- Logout: limpa estado e redireciona
- Persistência: [TODO] — atualmente apenas em memória (perdido ao fechar app)

### 7.2. Tratamento de Erros

- Serviços retornam promessas com try/catch
- Erros de banco propagados via `throw`
- [TODO]: tratamento global de erros (Error Boundary)

### 7.3. Loading States

- [TODO]: indicadores de loading padronizados (Skeleton, Spinner)
- React Query provê `isLoading`, `isFetching`, `isError`

### 7.4. Offline Mode

- App é offline-first por design (dados locais no SQLite)
- Mock services funcionam sem rede
- [TODO]: sincronização com backend remoto (quando disponível)

## 8. Observações

- **Duplicação intencional**: existem implementações em `services/database/` (SQLite) e `services/storage/` (AsyncStorage) para o mesmo domínio — a versão SQLite é a mais recente.
- **Arquivos legacy**: existem arquivos `.txt` e `Copy.ts` no diretório `services/database/` que são backups/cópias de desenvolvimento — não usados em produção.
- **README-ARQUITETURA.md** menciona Redux Toolkit, mas o código atual não usa Redux — usa React Context + React Query.
