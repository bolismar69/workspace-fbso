# User Story: US-FEAT-EP-0001-0001-0002 — filtrar as métricas do dashboard por período (últimos 7, 30, 90 dias, 

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0001](../epics/EP-0001-portal-administrativo-interno.md) ➔ Feature [FEAT-EP-0001-0001](../features/FEAT-EP-0001-0001-dashboard-de-metricas-operacionais.md) ➔ User Story US-FEAT-EP-0001-0001-0002
- **Prioridade:** Must Have
- **Data-Alvo:** 15/08/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Líder Comercial,
- **quero** filtrar as métricas do dashboard por período (últimos 7, 30, 90 dias, mês atual, ano atual) para analisar tendências de crescimento,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Líder Comercial está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Filtro de período disponível no topo do dashboard**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Líder Comercial está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Ao alterar o período, todos os indicadores são recalculados**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Líder Comercial está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Gráfico de evolução da base reflete o período selecionado**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0001-0001 — Dashboard de Métricas Operacionais](../features/FEAT-EP-0001-0001-dashboard-de-metricas-operacionais.md) | **Épico:** [EP-0001 — Portal Administrativo Interno](../epics/EP-0001-portal-administrativo-interno.md)
