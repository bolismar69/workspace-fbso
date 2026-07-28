# User Story: US-FEAT-EP-0002-0005-0023 — filtrar o histórico de auditoria por período e por tipo de ação para l

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) ➔ Feature [FEAT-EP-0002-0005](../features/FEAT-EP-0002-0005-historico-de-auditoria-administrativa.md) ➔ User Story US-FEAT-EP-0002-0005-0023
- **Prioridade:** Must Have
- **Data-Alvo:** 31/08/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Auditor Interno,
- **quero** filtrar o histórico de auditoria por período e por tipo de ação para localizar eventos específicos,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Auditor Interno está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Filtros disponíveis: período (data inicial e final), tipo de ação (criação, suspensão, alteração de plano)**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Auditor Interno está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Resultados ordenados do mais recente para o mais antigo**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0002-0005 — Histórico de Auditoria Administrativa](../features/FEAT-EP-0002-0005-historico-de-auditoria-administrativa.md) | **Épico:** [EP-0002 — Gestao de Clientes e Assinaturas](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md)
