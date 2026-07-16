---
title: "Suporte Offline — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
---

# Suporte Offline — Solar Fácil

## 1. Estratégia Atual

O Solar Fácil é **offline-first por design**. O app opera completamente sem conexão de rede:

- **Fonte de verdade:** SQLite local (`solarfacil.db`)
- **Dados de referência:** JSON mockado em bundle
- **Sem backend remoto:** 100% das funcionalidades disponíveis offline

## 2. Comportamento Atual

| Cenário | Comportamento |
|---|---|
| App abre sem internet | ✅ Funciona normalmente |
| CRUD de associados | ✅ Opera no SQLite local |
| Movimentações | ✅ Opera no SQLite local |
| Planos/FAQ/Saiba Mais | ✅ Dados mockados em bundle |
| Login | ✅ Verificação local (CPF/CNPJ + senha contra SQLite) |
| Sincronização | N/A — sem backend |

## 3. Cache Strategy

| Dados | Estratégia | Persistent? |
|---|---|---|
| Associados | SQLite | ✅ Sim (até delete explícito) |
| Movimentações | SQLite | ✅ Sim |
| Planos | Mock JSON (bundle) | ✅ Sim (estático) |
| FAQs | Mock JSON (bundle) | ✅ Sim (estático) |
| Concessionárias | Mock JSON (bundle) | ✅ Sim (estático) |
| Auth State | Memória (Context) | ❌ Perdido ao fechar app |

## 4. Sincronização Futura

### 4.1. Estratégia Planejada

Quando um backend remoto for integrado:

```
[App] ──write──► [SQLite (local cache)]
                      │
                      ▼
              [Sync Queue] ──HTTP──► [Backend API]
                      │
                      ▼
              [React Query Cache]
```

### 4.2. Políticas de Conflito

| Cenário | Estratégia |
|---|---|
| Conflito de update | Last Write Wins (timestamp) |
| Dados offline pendentes | Fila FIFO — processa ao reconectar |
| Dados stale vs. fresh | React Query `staleTime` + `gcTime` |

## 5. Recomendações

- **Persistir Auth State:** salvar sessão no AsyncStorage ou SQLite para manter login entre sessões
- **NetworkListener:** adicionar `@react-native-community/netinfo` para detectar conectividade quando backend existir
- **Sync Indicator:** indicador visual de dados pendentes de sincronização (quando backend existir)
