# User Story: US-FEAT-EP-0003-0001-0026 — desativar ou reativar um usuário para controlar seu acesso à plataforma

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0003](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md) ➔ Feature [FEAT-EP-0003-0001](../features/FEAT-EP-0003-0001-cadastro-e-convite-de-usuarios.md) ➔ User Story US-FEAT-EP-0003-0001-0026
- **Prioridade:** Must Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.1 — Adicionado cenário de reativação (Gap Analysis #15, 27/07/2026)
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador do Tenant,
- **quero** desativar um usuário para bloquear imediatamente seu acesso e também poder reativá-lo quando necessário,
- **para** controlar o ciclo de vida dos acessos dos usuários do meu tenant de forma autônoma.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal — Desativar]
- **Dado que** o Administrador do Tenant está autenticado no portal e visualiza a lista de usuários ativos,
- **Quando** aciona a ação "Desativar" em um usuário e confirma a operação,
- **Então** o sistema deve: **bloquear imediatamente o login do usuário desativado e exibi-lo na lista com o status "Inativo"**.

### Cenário 2: [Confirmação obrigatória antes da desativação]
- **Dado que** o Administrador do Tenant aciona "Desativar" em um usuário,
- **Quando** o diálogo de confirmação é exibido,
- **Então** o sistema deve: **exigir confirmação explícita antes de efetivar a desativação, exibindo o nome e e-mail do usuário que será desativado**.

### Cenário 3: [Fluxo Principal — Reativar]
- **Dado que** o Administrador do Tenant está autenticado no portal e visualiza a lista de usuários inativos,
- **Quando** aciona a ação "Reativar" em um usuário desativado e confirma a operação,
- **Então** o sistema deve: **restaurar o acesso do usuário imediatamente, mantendo suas permissões, papéis e vinculações originais, e exibi-lo com o status "Ativo"**.

### Cenário 4: [Segurança — Usuário desativado tenta login]
- **Dado que** um usuário com status "Inativo" tenta fazer login no portal,
- **Quando** informa credenciais válidas,
- **Então** o sistema deve: **recusar o acesso e exibir mensagem informando que a conta está desativada, orientando contato com o administrador do tenant**.

### Cenário 5: [Auditoria]
- **Dado que** o Administrador do Tenant desativa ou reativa um usuário,
- **Quando** a operação é concluída,
- **Então** o sistema deve: **registrar a ação no histórico de auditoria com: administrador responsável, data/hora, usuário afetado e tipo da ação (DESATIVAR / REATIVAR)**.

---

## 3. Regras de Negócio de Tela Relacionadas

- **RN-FEAT-EP-0003-0001-0026-01:** Apenas usuários com perfil `admin-do-tenant` podem desativar ou reativar outros usuários
- **RN-FEAT-EP-0003-0001-0026-02:** Um `admin-do-tenant` não pode desativar a si mesmo
- **RN-FEAT-EP-0003-0001-0026-03:** Usuário reativado herda exatamente as mesmas permissões, papéis e vinculações que possuía antes da desativação
- **RN-FEAT-EP-0003-0001-0026-04:** A reativação é possível a qualquer momento e não tem limite de tempo — um usuário pode ser reativado dias, semanas ou meses após a desativação

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0003-0001 — Cadastro e Convite de Usuários](../features/FEAT-EP-0003-0001-cadastro-e-convite-de-usuarios.md) | **Épico:** [EP-0003 — Gestao de Usuarios e Permissoes RBAC](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md)
