# User Story: US-FEAT-EP-0002-0002-0014 — visualizar o histórico de mudanças de status de cada conta (quando foi

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) ➔ Feature [FEAT-EP-0002-0002](../features/FEAT-EP-0002-0002-gestao-de-status-do-tenant.md) ➔ User Story US-FEAT-EP-0002-0002-0014
- **Prioridade:** Must Have
- **Data-Alvo:** 31/08/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador FBSO,
- **quero** visualizar o histórico de mudanças de status de cada conta (quando foi ativada, suspensa, reativada e por quem),
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Linha do tempo de status na tela de detalhes do Tenant**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Cada evento exibe: status anterior → novo status, responsável, data/hora, motivo (se aplicável)**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0002-0002 — Gestão de Status do Tenant](../features/FEAT-EP-0002-0002-gestao-de-status-do-tenant.md) | **Épico:** [EP-0002 — Gestao de Clientes e Assinaturas](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md)
