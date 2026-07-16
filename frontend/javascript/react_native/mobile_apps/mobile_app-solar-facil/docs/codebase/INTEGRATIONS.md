---
title: "Integrações — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["integrations", "dependencies", "mobile", "services"]
---

# Integrações — Solar Fácil

## 1. Visão Geral

O Solar Fácil é uma aplicação **offline-first** — não depende de backend remoto em sua configuração atual. Todas as integrações são locais (SQLite, AsyncStorage) ou baseadas em dados mock.

## 2. Integrações Locais

### 2.1. SQLite (expo-sqlite)

| Propriedade | Valor |
|---|---|
| Propósito | Banco de dados local primário |
| Biblioteca | `expo-sqlite` ~15.2.13 |
| Arquivo de banco | `solarfacil.db` |
| Inicialização | `services/database/initializeSQLiteDatabase.ts` |
| Tabelas | `associados` (30 colunas), `movimentacoes` (18 colunas) |
| Acesso | `DatabaseContext` + hooks React Query |
| Transações | `withExclusiveTransactionAsync` |

**Schema — associados:**
```
id INTEGER PRIMARY KEY AUTOINCREMENT
dataCadastro, dataAtualizacao TEXT
senha, status, tipoAssociado TEXT
tipoPessoa, cpf_cnpj (UNIQUE), nome, email, telefone TEXT
cep, endereco, numero, bairro, cidade, estado, complemento TEXT
aceitaTermos, observacoes TEXT
dataNascimento, nomeSocial TEXT
dataAbertura, razaoSocial, nomeFantasia TEXT
nomeConcessionaria, consumoMedio, planoDesejado TEXT
potenciaInstalada, disponibilidade, tipoConexao TEXT
```

**Schema — movimentacoes:**
```
id INTEGER PRIMARY KEY AUTOINCREMENT
dataCadastro, dataAtualizacao TEXT
associadoId INTEGER (FK → associados.id)
mes, ano INTEGER
valorTotal, energiaRecebidaKwh, valorEnergiaRecebida, tarifaUnitariaKwh REAL
valorCobrado, valorEconomizado, percentualEconomizado REAL
dataVencimento, dataPagamento, statusPagamento TEXT
observacoes TEXT
```

### 2.2. AsyncStorage

| Propriedade | Valor |
|---|---|
| Propósito | Armazenamento key-value (legado/fallback) |
| Biblioteca | `@react-native-async-storage/async-storage` 2.1.2 |
| Uso | CRUD de associados, beneficiados, fornecedores, movimentações |
| Localização | `services/storage/` |

## 3. Dados Mock (Estáticos)

### 3.1. Fontes de Dados

| Arquivo | Dados | Uso |
|---|---|---|
| `mocks/mockConcessionarias.json` | Lista de concessionárias de energia | Select de concessionária |
| `mocks/mockConsumoMedio.json` | Faixas de consumo médio (kWh) | Cálculo de economia |
| `mocks/mockFAQs.json` | Perguntas frequentes | Tela de FAQ |
| `mocks/mockPlans.json` | Planos comerciais | Tela de planos |

### 3.2. Serviços Mock

| Arquivo | Propósito |
|---|---|
| `services/serviceConcessionarias.ts` | Retorna lista de concessionárias |
| `services/serviceConsumoMedio.ts` | Retorna dados de consumo médio |
| `services/serviceFAQs.ts` | Retorna FAQs |
| `services/servicePlans.ts` | Retorna planos comerciais |
| `services/mock/serviceBeneficiadoMock.ts` | Simula API de beneficiado |
| `services/mock/serviceFornecedorMock.ts` | Simula API de fornecedor |

## 4. Serviços Externos (Potenciais)

### 4.1. Expo Updates (OTA)

| Propriedade | Valor |
|---|---|
| Propósito | Atualizações over-the-air |
| URL | `https://u.expo.dev/eaf32393-0c03-4e40-9005-fd5955794e5a` |
| Runtime Version | `appVersion` (policy) |
| Configuração | `app.json` → `updates` |

### 4.2. EAS Build & Submit

| Propriedade | Valor |
|---|---|
| Propósito | Build e publicação nas lojas |
| Project ID | `eaf32393-0c03-4e40-9005-fd5955794e5a` |
| Owner | `bolismar69` |
| Perfis | development, preview, production |

### 4.3. Expo Linking (Deep Links)

| Propriedade | Valor |
|---|---|
| Propósito | Deep linking e universal links |
| Scheme | `solar-facil` |
| Biblioteca | `expo-linking` ~7.1.5 |

## 5. Pacotes Nativos (Expo Plugins)

| Plugin | Propósito |
|---|---|
| `expo-router` | File-based routing |
| `expo-splash-screen` | Splash screen (imagem: 200px, fundo branco) |
| `expo-sqlite` | Banco de dados SQLite |

## 6. APIs Externas

**Status atual: Nenhuma API remota configurada.**

A aplicação opera completamente offline com:
- Dados transacionais no SQLite local
- Dados de referência em arquivos mock JSON
- Axios incluído como dependência para uso futuro

### 6.1. Plano de Integração Futura

Quando um backend for implementado:
- URL base: `[ASK USER]`
- Autenticação: JWT Bearer token
- Sincronização: estratégia offline-first com fila de pendências
- Cache: React Query com `staleTime` configurável
- [TODO]: Definir contrato OpenAPI dos endpoints

## 7. Monitoramento & Analytics

**Status: Não configurado.**

[TODO]:
- Crash reporting: Sentry ou Crashlytics
- Analytics: Firebase Analytics ou Amplitude
- Push notifications: FCM + APNs via Expo Notifications

## 8. Diagrama de Dependências

```
┌──────────────────────────────┐
│   Solar Fácil App            │
├──────────────────────────────┤
│  SQLite (expo-sqlite)        │ ← Dados transacionais
│  AsyncStorage                │ ← Persistência legada
│  Mock JSON files             │ ← Dados de referência
│  Expo Updates (OTA)          │ ← Atualizações
│  Expo Linking (Deep Links)   │ ← Navegação externa
│  EAS Build                   │ ← Build & Deploy
└──────────────────────────────┘
```
