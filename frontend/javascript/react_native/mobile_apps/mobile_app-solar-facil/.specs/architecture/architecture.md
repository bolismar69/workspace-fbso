---
title: "Arquitetura — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["architecture", "mobile", "c4", "expo", "offline-first"]
---

# Arquitetura — Solar Fácil

## 1. Resumo Executivo

O Solar Fácil é um aplicativo mobile **offline-first** desenvolvido com **Expo SDK 53** e **React Native 0.79** (New Architecture). Utiliza **SQLite local** como fonte primária de dados, **React Context API** para gerenciamento de estado global, e **TanStack React Query** para cache e sincronização de dados. A interface é estilizada com **NativeWind 4** (Tailwind CSS) e a navegação é gerenciada via **Expo Router 5** (file-based routing).

## 2. Padrão Arquitetural

**Context + Hooks Architecture** com Provider Composition.

### 2.1. Princípios de Design

| Princípio | Implementação |
|---|---|
| **Component-Based** | Componentes React reutilizáveis com props tipadas (TypeScript strict) |
| **Separation of Concerns** | Screens → Components → Hooks → Services → Types |
| **Offline-First** | SQLite como fonte de verdade; sem dependência de rede |
| **Platform Agnostic** | Expo managed workflow; platform-specific code via `Platform.OS` |
| **Type Safety** | TypeScript strict mode + typed routes (Expo Router) |
| **Optimistic UI** | React Query mutations com cache invalidation |

### 2.2. Visão em Camadas

```
┌────────────────────────────────────────────────┐
│  APP LAYER (Expo Router)                       │
│  _layout.tsx → 8 rotas (tabs)                  │
├────────────────────────────────────────────────┤
│  PROVIDER LAYER                                │
│  ReactQuery → AppTheme → Auth → Database       │
├────────────────────────────────────────────────┤
│  SCREEN LAYER                                  │
│  Home, Planos, Login, Cadastro, FAQ, ...       │
├────────────────────────────────────────────────┤
│  COMPONENT LAYER                               │
│  Forms, Inputs, Cards, Lists, UI               │
├────────────────────────────────────────────────┤
│  HOOK LAYER                                    │
│  Queries (React Query), Mutations, Validation  │
├────────────────────────────────────────────────┤
│  SERVICE LAYER                                 │
│  SQLite (expo-sqlite) + Mock Services          │
├────────────────────────────────────────────────┤
│  INFRASTRUCTURE                                │
│  Types, Styles (NativeWind + Theme), Utils     │
└────────────────────────────────────────────────┘
```

## 3. Stack Tecnológica

| Categoria | Tecnologia | Versão |
|---|---|---|
| Framework | React Native | 0.79.3 |
| Plataforma | Expo (managed) | ~53.0.9 |
| Linguagem | TypeScript | ~5.8.3 |
| Navegação | Expo Router | ~5.0.6 |
| Estilização | NativeWind + Tailwind CSS | ^4.0.1 / ^3.4.17 |
| Estado Global | React Context API | built-in |
| Estado Assíncrono | TanStack React Query | ^5.83.0 |
| Formulários | react-hook-form + yup | ^7.56.4 / ^1.6.1 |
| Banco de Dados | expo-sqlite | ~15.2.13 |
| Animações | react-native-reanimated + moti | ~3.17.4 / ^0.30.0 |
| Build & Deploy | EAS Build + Expo Updates | >= 16.14.1 |
| Gráficos | victory-native | ^41.17.3 |

## 4. Decisões Arquiteturais (ADRs)

| ADR | Decisão | Status |
|---|---|---|
| [ADR-001](adrs/adr-001.md) | Expo SDK 53 + React Native 0.79 (New Architecture) | Aceito |
| [ADR-002](adrs/adr-002.md) | NativeWind 4 para estilização | Aceito |
| [ADR-003](adrs/adr-003.md) | Context API + React Query para estado | Aceito |
| [ADR-004](adrs/adr-004.md) | SQLite offline-first como fonte de dados | Aceito |

## 5. Cross-Cutting Concerns

| Concern | Implementação |
|---|---|
| **Autenticação** | AuthContext (CPF/CNPJ + senha, estado em memória) |
| **Tema** | AppThemeContext (claro/escuro, segue preferência do sistema) |
| **Persistência** | DatabaseContext (SQLite auto-inicialização) |
| **Validação** | yup schemas + react-hook-form |
| **Erro** | try/catch em serviços; [TODO] Error Boundary global |
| **Loading** | React Query `isLoading`/`isFetching`; [TODO] Skeleton components |
| **Empty State** | [TODO] — não padronizado |
| **Offline** | App é offline-first por design (SQLite local) |

## 6. Segurança

| Aspecto | Status |
|---|---|
| Senhas | ⚠️ Plain text no SQLite (DT-001) |
| Dados em repouso | ❌ Sem criptografia no SQLite |
| Comunicação | N/A (sem rede atualmente) |
| Código | ❌ Sem ofuscação |
| SSL Pinning | N/A (sem backend) |

## 7. Observações

- **README-ARQUITETURA.md** menciona Redux Toolkit incorretamente — o app não usa Redux.
- A arquitetura é adequada ao estágio atual do projeto (MVP offline), mas precisará evoluir quando um backend remoto for integrado.
- A duplicação SQLite/AsyncStorage indica uma migração em progresso — SQLite é o alvo final.
