---
title: "C4 — Componentes — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
diagram_type: "C4 — Nível 3 (Component)"
---

# C4 — Nível 3: Componentes

## Diagrama de Componentes

```mermaid
C4Component
    title Component diagram for Solar Fácil App

    Container_Boundary(app, "Solar Fácil App") {
        Component(providers, "Provider Layer", "React Context", "ReactQuery → AppTheme → Auth → Database")
        Component(router, "Expo Router", "File-based Routing", "8 rotas em bottom tabs")
        Component(home, "HomeScreen", "React Component", "Tela inicial com visão geral")
        Component(planos, "PlanosScreen", "React Component", "Planos comerciais")
        Component(login, "LoginScreen", "React Component", "Autenticação local")
        Component(cadastro, "CadastroScreen", "React Component", "Cadastro de associados")
        Component(movimentacao, "MovimentacaoScreen", "React Component", "Movimentações mensais")
        Component(faq, "FAQScreen", "React Component", "Perguntas frequentes")

        Component(forms, "Form Components", "React Components", "FormBeneficiado, FormFornecedor, FormCadastroAssociado")
        Component(inputs, "Input Components", "React Components", "InputText, InputDate, InputSelect, InputSwitch, etc.")

        Component(queryHooks, "Query Hooks", "React Query", "useQueryAssociadosSearchAll, useQueryMovimentacoesSearchByAssociadoId, etc.")
        Component(mutationHooks, "Mutation Hooks", "React Query", "useMutationAssociadoInsertRecord, useMutationMovimentacoesUpdateRecord, etc.")
        Component(formHook, "useFormValidation", "react-hook-form + yup", "Validação de formulários")

        Component(sqliteService, "SQLite Service", "expo-sqlite", "initializeDatabase, getDatabaseConnection")
        Component(mockServices, "Mock Services", "JSON", "serviceConcessionarias, serviceFAQs, servicePlans")

        ComponentDb(sqliteDb, "SQLite DB", "solarfacil.db", "Tabelas: associados, movimentacoes")
    }

    Rel(router, providers, "Encapsula")
    Rel(router, home, "Rota /")
    Rel(router, planos, "Rota /planos")
    Rel(router, login, "Rota /login")
    Rel(router, cadastro, "Rota /cadastro")
    Rel(router, movimentacao, "Rota /movimentacao")
    Rel(router, faq, "Rota /faq")

    Rel(home, queryHooks, "Usa")
    Rel(planos, mockServices, "Usa")
    Rel(login, queryHooks, "Usa (auth)")
    Rel(cadastro, forms, "Usa")
    Rel(movimentacao, queryHooks, "Usa")
    Rel(faq, mockServices, "Usa")

    Rel(forms, inputs, "Usa")
    Rel(forms, formHook, "Usa (validação)")

    Rel(queryHooks, sqliteService, "Lê dados")
    Rel(mutationHooks, sqliteService, "Escreve dados")
    Rel(sqliteService, sqliteDb, "SQL queries")
```

## Árvore de Componentes Principal

```
App (Expo Router)
├── ReactQueryProvider
│   └── AppThemeProvider
│       └── AuthProvider
│           └── SafeAreaProvider
│               └── DatabaseProvider
│                   └── GestureHandlerRootView
│                       └── Tabs (8 rotas)
│                           ├── index (HomeScreen)
│                           ├── planos (PlanosScreen)
│                           ├── saibamais (SaibaMaisScreen)
│                           ├── faq (FAQScreen)
│                           ├── login (AssociadoLoginScreen)
│                           ├── cadastro (AssociadoCadastroScreen)
│                           ├── movimentacao (MovimentacaoComCardScreen)
│                           └── listatodos (AssociadoListaTodosScreen)
```

## Fluxos de Dados por Tela

### HomeScreen
```
HomeScreen → useQueryAssociadosSearchAll → SQLite → Exibe cards de associados
```

### Login
```
LoginScreen → Formulário (CPF/CNPJ + Senha)
  → useQueryAssociadosSearchByCpfCnpjSenha → SQLite
  → AuthContext.login() → Redireciona para Home
```

### Cadastro
```
CadastroScreen → FormCadastroAssociado
  → Inputs (InputText, InputSelect, InputDate, etc.)
  → useFormValidation (yup schema)
  → useMutationAssociadoInsertRecord → SQLite
```

### Movimentações
```
MovimentacaoScreen → useQueryMovimentacoesSearchByAssociadoId → SQLite
  → Victory Chart (gráfico de economia)
  → Cards de movimentação (valores, status)
```

### Planos
```
PlanosScreen → servicePlans (mock JSON) → CardPlan (exibe planos comerciais)
```
