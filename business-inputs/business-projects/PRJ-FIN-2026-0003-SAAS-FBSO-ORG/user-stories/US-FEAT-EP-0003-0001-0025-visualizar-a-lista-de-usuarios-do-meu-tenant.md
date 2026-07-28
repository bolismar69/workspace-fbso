# User Story: US-FEAT-EP-0003-0001-0025 — visualizar a lista de usuários do meu tenant com seus respectivos papé

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0003](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md) ➔ Feature [FEAT-EP-0003-0001](../features/FEAT-EP-0003-0001-cadastro-e-convite-de-usuarios.md) ➔ User Story US-FEAT-EP-0003-0001-0025
- **Prioridade:** Must Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador do Tenant,
- **quero** visualizar a lista de usuários do meu tenant com seus respectivos papéis, unidades vinculadas e status (ativo, inativo, convite pendente),
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Administrador do Tenant está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Lista exibe: nome, e-mail, papel principal, unidades vinculadas, status**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Administrador do Tenant está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Filtro por status: Todos, Ativos, Pendentes (convite não aceito), Inativos**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Administrador do Tenant está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Indicador visual para convites ainda não aceitos**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0003-0001 — Cadastro e Convite de Usuários](../features/FEAT-EP-0003-0001-cadastro-e-convite-de-usuarios.md) | **Épico:** [EP-0003 — Gestao de Usuarios e Permissoes RBAC](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md)
