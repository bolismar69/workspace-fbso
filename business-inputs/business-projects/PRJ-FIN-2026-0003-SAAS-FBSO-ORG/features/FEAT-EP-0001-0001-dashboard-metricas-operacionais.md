# FEATURE - FEAT-EP-0001-0001: Dashboard de Métricas Operacionais

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-0001-0001 — Dashboard de Métricas Operacionais |
| **Épico** | [EP-0001 — Portal Administrativo Interno](../epics/EP-0001-portal-administrativo-interno.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0001](../epics/EP-0001-portal-administrativo-interno.md) | **Próximo:** [FEAT-EP-0001-0002 — Visão de Contas com Filtros](../FEAT-EP-0001-0002-visao-contas-com-filtros.md)

**Requisitos BRD Vinculados:** [BR-01](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Dashboard Administrativo

---

## Objetivo de Negócio
Prover visibilidade em tempo real da operação do SaaS para o time interno da FBSO.ORG, permitindo acompanhar a saúde da base de clientes e tomar decisões operacionais com agilidade.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-001 | Como **Administrador FBSO**, quero visualizar os indicadores principais da operação em uma tela de dashboard para ter uma visão rápida da saúde do SaaS | • Dashboard carrega com os indicadores atualizados em até 3 segundos • Indicadores exibidos: total de contas ativas, total de contas por status, total de contas por plano • Cada indicador é clicável e leva à lista filtrada correspondente |
| US-002 | Como **Líder Comercial**, quero filtrar as métricas do dashboard por período (últimos 7, 30, 90 dias, mês atual, ano atual) para analisar tendências de crescimento | • Filtro de período disponível no topo do dashboard • Ao alterar o período, todos os indicadores são recalculados • Gráfico de evolução da base reflete o período selecionado |
| US-003 | Como **Diretoria**, quero visualizar um gráfico de evolução da base de clientes ao longo do tempo para acompanhar o crescimento do SaaS | • Gráfico de linhas ou barras exibe a quantidade de novas contas por mês • Gráfico permite alternar entre visão de contas totais e novas contas • Período do gráfico segue o filtro aplicado no dashboard |

## Regras de Negócio

- **RN-FEAT-EP-0001-0001-0001:** Métricas consideram apenas tenants com status diferente de "Excluído" (soft delete)
- **RN-FEAT-EP-0001-0001-0002:** Período padrão do dashboard ao carregar: mês atual
- **RN-FEAT-EP-0001-0001-0003:** Indicadores que exibem "zero" devem ser apresentados com o número 0, nunca em branco

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-01** | Dashboard Administrativo | [EP-0001](../epics/EP-0001-portal-administrativo-interno.md) / J1: Acompanhamento diário da operação · J2: Análise de crescimento por plano | **FEAT-EP-0001-0001** — Dashboard de Métricas Operacionais |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0001](../epics/EP-0001-portal-administrativo-interno.md) | **Próximo:** [FEAT-EP-0001-0002 — Visão de Contas com Filtros](../FEAT-EP-0001-0002-visao-contas-com-filtros.md)

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Estrutura modular v2.0.*

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
