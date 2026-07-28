# User Story: US-FEAT-EP-0003-0002-0027 — atribuir um dos papéis padrão (Admin, Gerente, Operador, Auditor) a ca

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0003](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md) ➔ Feature [FEAT-EP-0003-0002](../features/FEAT-EP-0003-0002-definicao-de-papeis-e-permissoes-rbac.md) ➔ User Story US-FEAT-EP-0003-0002-0027
- **Prioridade:** Must Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador do Tenant,
- **quero** atribuir um dos papéis padrão (Admin, Gerente, Operador, Auditor) a cada usuário para definir seu nível de acesso na plataforma,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Administrador do Tenant está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Seletor de papel no cadastro e na edição do usuário**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Administrador do Tenant está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Descrição de cada papel disponível como tooltip**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Administrador do Tenant está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Alteração de papel registrada em auditoria**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0003-0002 — Definição de Papéis e Permissões (RBAC)](../features/FEAT-EP-0003-0002-definicao-de-papeis-e-permissoes-rbac.md) | **Épico:** [EP-0003 — Gestao de Usuarios e Permissoes RBAC](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md)
