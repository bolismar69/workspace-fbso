# User Story: US-FEAT-EP-0002-0004-0020 — realizar upgrade ou downgrade de plano de um cliente, mantendo o histó

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) ➔ Feature [FEAT-EP-0002-0004](../features/FEAT-EP-0002-0004-vinculacao-e-gestao-de-assinaturas.md) ➔ User Story US-FEAT-EP-0002-0004-0020
- **Prioridade:** Must Have
- **Data-Alvo:** 31/08/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Líder Comercial,
- **quero** realizar upgrade ou downgrade de plano de um cliente, mantendo o histórico da assinatura anterior,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Líder Comercial está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Ao trocar de plano, assinatura anterior é finalizada com data de término**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Líder Comercial está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Nova assinatura é criada com data de início igual ao dia seguinte ao término da anterior**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Líder Comercial está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Histórico exibe todas as assinaturas do cliente em ordem cronológica**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0002-0004 — Vinculação e Gestão de Assinaturas](../features/FEAT-EP-0002-0004-vinculacao-e-gestao-de-assinaturas.md) | **Épico:** [EP-0002 — Gestao de Clientes e Assinaturas](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md)
