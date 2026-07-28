# User Story: US-FEAT-EP-0003-0002-0029 — que ao atribuir o papel "Admin do Tenant" a um usuário, ele automatica

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0003](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md) ➔ Feature [FEAT-EP-0003-0002](../features/FEAT-EP-0003-0002-definicao-de-papeis-e-permissoes-rbac.md) ➔ User Story US-FEAT-EP-0003-0002-0029
- **Prioridade:** Must Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador do Tenant,
- **quero** que ao atribuir o papel "Admin do Tenant" a um usuário, ele automaticamente tenha acesso a todas as Unidades de Negócio e todos os módulos do tenant,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Administrador do Tenant está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Admin do Tenant vê todas as unidades de negócio no seletor**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Administrador do Tenant está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Admin do Tenant vê todos os módulos contratados no App Switcher**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Administrador do Tenant está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Não é necessário configurar permissões individuais para Admin**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0003-0002 — Definição de Papéis e Permissões (RBAC)](../features/FEAT-EP-0003-0002-definicao-de-papeis-e-permissoes-rbac.md) | **Épico:** [EP-0003 — Gestao de Usuarios e Permissoes RBAC](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md)
