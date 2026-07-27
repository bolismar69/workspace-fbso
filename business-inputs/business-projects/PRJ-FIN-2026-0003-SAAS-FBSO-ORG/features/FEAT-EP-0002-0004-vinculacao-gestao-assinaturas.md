# FEATURE - FEAT-EP-0002-0004: Vinculação e Gestão de Assinaturas

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-0002-0004 — Vinculação e Gestão de Assinaturas |
| **Épico** | [EP-0002 — Gestão de Clientes e Assinaturas](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) | **Anterior:** [FEAT-EP-0002-0003 — Planos Comerciais](../FEAT-EP-0002-0003-configuracao-planos-comerciais.md) | **Próximo:** [FEAT-EP-0002-0005 — Histórico de Auditoria](../FEAT-EP-0002-0005-historico-auditoria-administrativa.md)

**Requisitos BRD Vinculados:** [BR-04](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Vinculação de Assinaturas

---

## Objetivo de Negócio
Vincular clientes a planos com controle de vigência, permitindo upgrade/downgrade e acompanhamento do status da assinatura.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-019 | Como **Administrador FBSO**, quero vincular um cliente a um plano comercial definindo data de início, vigência e status da assinatura | • Seletor de plano (apenas planos ativos listados) • Campos: data de início, data de término (ou "indeterminado"), status (Ativa, Suspensa, Cancelada) • Ao ativar assinatura, módulos do plano são liberados para o tenant |
| US-020 | Como **Líder Comercial**, quero realizar upgrade ou downgrade de plano de um cliente, mantendo o histórico da assinatura anterior | • Ao trocar de plano, assinatura anterior é finalizada com data de término • Nova assinatura é criada com data de início igual ao dia seguinte ao término da anterior • Histórico exibe todas as assinaturas do cliente em ordem cronológica |
| US-021 | Como **Administrador FBSO**, quero suspender a assinatura de um cliente, o que deve bloquear o acesso dele aos módulos do plano | • Suspensão da assinatura bloqueia acesso aos módulos em até 5 minutos • Status da assinatura alterado para "Suspensa" • Reativação da assinatura restaura acesso aos mesmos módulos |

## Regras de Negócio

- **RN-FEAT-EP-0002-0004-0001:** Um tenant pode ter apenas uma assinatura ativa por vez
- **RN-FEAT-EP-0002-0004-0002:** Upgrade/downgrade não pode deixar o tenant sem assinatura ativa durante a transição
- **RN-FEAT-EP-0002-0004-0003:** Data de término de assinatura é opcional (planos sem data de término = vigência contínua)

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-04** | Vinculação de Assinaturas | [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) / J1: Ativação de novo cliente · J3: Upgrade de plano | **FEAT-EP-0002-0004** — Vinculação e Gestão de Assinaturas |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
