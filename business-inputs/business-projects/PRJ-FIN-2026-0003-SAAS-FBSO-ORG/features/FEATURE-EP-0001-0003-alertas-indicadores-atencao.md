# FEATURE - EP-0001-0003: Alertas e Indicadores de Atenção

| Campo | Detalhe |
|-------|---------|
| **Feature** | EP-0001-0003 — Alertas e Indicadores de Atenção |
| **Épico** | [EP-0001 — Portal Administrativo Interno](../epics/EP-0001-portal-administrativo-interno.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0001](../epics/EP-0001-portal-administrativo-interno.md) | **Anterior:** [EP-0001-0002 — Visão de Contas](../FEATURE-EP-0001-0002-visao-contas-com-filtros.md)

**Requisitos BRD Vinculados:** [BR-01](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Dashboard Administrativo

---

## Objetivo de Negócio
Destacar proativamente situações que exigem ação do time administrativo, evitando que problemas passem despercebidos.

**Prioridade:** Should Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-006 | Como **Administrador FBSO**, quero ver indicadores de alerta no dashboard para contas que precisam de atenção (ex: onboarding incompleto há mais de 48h, assinatura suspensa) | • Alertas aparecem como cards coloridos (amarelo: atenção; vermelho: crítico) no topo do dashboard • Cada alerta exibe a quantidade de contas na situação e é clicável • Ao clicar, direciona para a lista filtrada das contas naquela situação |
| US-007 | Como **Administrador FBSO**, quero que o sistema destaque visualmente na lista de contas aquelas com status irregular para identificação rápida durante a navegação | • Contas suspensas exibem ícone ou cor de destaque na lista • Ao passar o cursor, tooltip explica o motivo da suspensão (se registrado) |

## Regras de Negócio

- **RN-FEAT-EP-0001-0003-0001:** Alerta de "onboarding incompleto" dispara após 48 horas da criação da conta
- **RN-FEAT-EP-0001-0003-0002:** Alertas são visíveis para todos os usuários do time interno; não há personalização por usuário nesta fase

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-01** | Dashboard Administrativo | [EP-0001](../epics/EP-0001-portal-administrativo-interno.md) / J1: Acompanhamento diário da operação | **EP-0001-0003** — Alertas e Indicadores de Atenção |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0001](../epics/EP-0001-portal-administrativo-interno.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
