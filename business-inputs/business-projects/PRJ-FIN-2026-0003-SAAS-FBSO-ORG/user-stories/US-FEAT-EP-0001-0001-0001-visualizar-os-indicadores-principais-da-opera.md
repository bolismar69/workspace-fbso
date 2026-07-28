# User Story: US-FEAT-EP-0001-0001-0001 — visualizar os indicadores principais da operação em uma tela de dashbo

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0001](../epics/EP-0001-portal-administrativo-interno.md) ➔ Feature [FEAT-EP-0001-0001](../features/FEAT-EP-0001-0001-dashboard-de-metricas-operacionais.md) ➔ User Story US-FEAT-EP-0001-0001-0001
- **Prioridade:** Must Have
- **Data-Alvo:** 15/08/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador FBSO,
- **quero** visualizar os indicadores principais da operação em uma tela de dashboard para ter uma visão rápida da saúde do SaaS,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Dashboard carrega com os indicadores atualizados em até 3 segundos**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Indicadores exibidos: total de contas ativas, total de contas por status, total de contas por plano**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Cada indicador é clicável e leva à lista filtrada correspondente**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0001-0001 — Dashboard de Métricas Operacionais](../features/FEAT-EP-0001-0001-dashboard-de-metricas-operacionais.md) | **Épico:** [EP-0001 — Portal Administrativo Interno](../epics/EP-0001-portal-administrativo-interno.md)
