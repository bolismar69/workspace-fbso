# FEATURE - FEAT-EP-0001-0002: Visão de Contas com Filtros

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-0001-0002 — Visão de Contas com Filtros |
| **Épico** | [EP-0001 — Portal Administrativo Interno](../epics/EP-0001-portal-administrativo-interno.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0001](../epics/EP-0001-portal-administrativo-interno.md) | **Anterior:** [FEAT-EP-0001-0001 — Dashboard](../FEAT-EP-0001-0001-dashboard-metricas-operacionais.md) | **Próximo:** [FEAT-EP-0001-0003 — Alertas](../FEAT-EP-0001-0003-alertas-indicadores-atencao.md)

**Requisitos BRD Vinculados:** [BR-01](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Dashboard Administrativo

---

## Objetivo de Negócio
Permitir que o time administrativo localize rapidamente qualquer conta de cliente e visualize suas informações principais.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-FEAT-EP-0001-0002-0004 | Como **Administrador FBSO**, quero visualizar a lista completa de contas de clientes com informações resumidas (nome, plano, status, data de criação) para navegar pela base | • Tabela exibe: razão social, plano contratado, status, data de criação, data da última ação • Lista ordenada por data de criação (mais recentes primeiro) • Paginação a cada 25 registros |
| US-FEAT-EP-0001-0002-0005 | Como **Administrador FBSO**, quero buscar uma conta específica por nome ou razão social para localizar rapidamente um cliente | • Campo de busca textual no topo da lista • Busca filtra em tempo real (a partir de 3 caracteres digitados) • Resultados exibem correspondências parciais (ex: "Super" encontra "Supermercado Bom Preço") |

## Regras de Negócio

- **RN-FEAT-EP-0001-0002-0001:** Contas com status "Excluído" (soft delete) não aparecem na lista padrão
- **RN-FEAT-EP-0001-0002-0002:** A busca não diferencia maiúsculas de minúsculas

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-01** | Dashboard Administrativo | [EP-0001](../epics/EP-0001-portal-administrativo-interno.md) / J1: Acompanhamento diário da operação | **FEAT-EP-0001-0002** — Visão de Contas com Filtros |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0001](../epics/EP-0001-portal-administrativo-interno.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
