# User Story: US-FEAT-EP-0003-0002-0028 — que cada papel tenha um conjunto predefinido de permissões: Admin do T

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0003](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md) ➔ Feature [FEAT-EP-0003-0002](../features/FEAT-EP-0003-0002-definicao-de-papeis-e-permissoes-rbac.md) ➔ User Story US-FEAT-EP-0003-0002-0028
- **Prioridade:** Must Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Gestor de Produto,
- **quero** que cada papel tenha um conjunto predefinido de permissões: Admin do Tenant (acesso total), Gerente de Unidade (gerencia sua unidade), Operador (executa tarefas), Auditor (apenas visualiza),
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Gestor de Produto está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Permissões mapeadas conforme tabela de papéis (RN-FEAT-EP-0003-0002-0001)**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Gestor de Produto está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Permissões não são customizáveis por tenant nesta fase**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Gestor de Produto está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Papéis são os mesmos para todos os tenants**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0003-0002 — Definição de Papéis e Permissões (RBAC)](../features/FEAT-EP-0003-0002-definicao-de-papeis-e-permissoes-rbac.md) | **Épico:** [EP-0003 — Gestao de Usuarios e Permissoes RBAC](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md)
