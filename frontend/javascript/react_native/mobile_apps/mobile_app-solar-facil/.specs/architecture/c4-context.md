---
title: "C4 — Contexto — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
diagram_type: "C4 — Nível 1 (System Context)"
---

# C4 — Nível 1: Contexto do Sistema

## Diagrama de Contexto

```mermaid
C4Context
    title System Context diagram for Solar Fácil

    Person(associado_pf, "Associado PF", "Pessoa física que busca economia na conta de energia via energia solar")
    Person(associado_pj, "Associado PJ", "Pessoa jurídica que busca reduzir custos operacionais com energia solar")
    Person(fornecedor, "Fornecedor", "Gerador de energia solar excedente que deseja vender créditos")

    System(solar_facil, "Solar Fácil App", "Aplicativo mobile para gestão de associações de energia solar: cadastro, simulação de economia, movimentações mensais e planos comerciais")

    System_Ext(concessionarias, "Concessionárias", "Distribuidoras de energia elétrica (dados de referência mockados)")
    System_Ext(expo_updates, "Expo Updates", "Serviço de atualização OTA do Expo")
    System_Ext(eas_build, "EAS Build", "Serviço de build e publicação nas lojas")

    Rel(associado_pf, solar_facil, "Usa para simular economia, gerenciar cadastro e visualizar movimentações")
    Rel(associado_pj, solar_facil, "Usa para reduzir custos operacionais com energia solar")
    Rel(fornecedor, solar_facil, "Usa para cadastrar oferta de energia excedente")

    Rel(solar_facil, concessionarias, "Consulta dados de concessionárias (mock)")
    Rel(solar_facil, expo_updates, "Recebe atualizações OTA", "HTTPS")
    Rel(solar_facil, eas_build, "Build e deploy gerenciados via EAS", "HTTPS")
```

## Descrição dos Elementos

### Pessoas (Actors)

| Ator | Descrição | Necessidades |
|---|---|---|
| **Associado PF** | Pessoa física, consumidora de energia | Cadastrar-se, simular economia com energia solar, visualizar movimentações mensais |
| **Associado PJ** | Pessoa jurídica, potencialmente consumidora e fornecedora | Gerenciar múltiplas unidades, analisar economia em escala |
| **Fornecedor** | Gerador de energia solar excedente | Cadastrar sistema de geração, disponibilizar créditos de energia |

### Sistemas Externos

| Sistema | Descrição | Tipo de Integração |
|---|---|---|
| **Concessionárias** | Dados de referência de distribuidoras de energia (mockados localmente) | JSON estático local |
| **Expo Updates** | Serviço de atualização over-the-air do Expo | HTTPS (expo.dev) |
| **EAS Build** | Serviço de build e publicação do Expo | HTTPS (expo.dev) |

## Fluxos de Dados Principais

### Fluxo 1: Cadastro de Associado
```
Associado → App → Formulário (react-hook-form + yup) → useMutationAssociadoInsertRecord → SQLite
```

### Fluxo 2: Login
```
Associado → App → CPF/CNPJ + Senha → useQueryAssociadosSearchByCpfCnpjSenha → SQLite → AuthContext
```

### Fluxo 3: Visualização de Movimentações
```
Associado → App → Tela de Movimentações → useQueryMovimentacoesSearchByAssociadoId → SQLite → Gráfico (Victory)
```
