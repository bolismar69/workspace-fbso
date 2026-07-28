# User Story: US-FEAT-EP-0003-0001-0061 — reativar manualmente um usuário antes do fim do período de suspensão temporária

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0003](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md) ➔ Feature [FEAT-EP-0003-0001](../features/FEAT-EP-0003-0001-cadastro-e-convite-de-usuarios.md) ➔ User Story US-FEAT-EP-0003-0001-0061
- **Prioridade:** Should Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.0 — Gap Analysis #13 (27/07/2026)
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador do Tenant,
- **quero** reativar manualmente um usuário que está suspenso temporariamente antes da data prevista de retorno (ex: funcionário retornou antes do previsto),
- **para** restaurar o acesso do usuário imediatamente sem precisar esperar a data de reativação automática.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal — Reativar antes da data prevista]
- **Dado que** o Administrador do Tenant está autenticado e visualiza um usuário com status "Suspenso Temporariamente" cuja data de retorno ainda não foi atingida,
- **Quando** aciona "Reativar Agora" e confirma a operação,
- **Então** o sistema deve: **alterar o status do usuário para "Ativo" imediatamente, restaurar o acesso normalmente, cancelar a reativação automática que estava agendada e registrar a reativação manual no histórico de auditoria**.

### Cenário 2: [Usuário reativado tenta login]
- **Dado que** um usuário foi reativado manualmente antes do fim do prazo de suspensão,
- **Quando** tenta fazer login com suas credenciais,
- **Então** o sistema deve: **permitir o acesso normalmente, sem qualquer indicação residual de suspensão, mantendo todas as permissões e vinculações originais**.

### Cenário 3: [Auditoria — Reativação manual vs. automática]
- **Dado que** o Administrador do Tenant reativa manualmente um usuário,
- **Quando** a operação é concluída,
- **Então** o sistema deve: **registrar no histórico de auditoria: administrador responsável, data/hora, usuário afetado, indicando que a reativação foi "MANUAL" (antes do prazo) — diferenciando-a da reativação automática por data**.

---

## 3. Regras de Negócio de Tela Relacionadas

- **RN-FEAT-EP-0003-0001-0061-01:** Apenas usuários com perfil `admin-do-tenant` podem reativar manualmente usuários suspensos temporariamente
- **RN-FEAT-EP-0003-0001-0061-02:** A reativação manual cancela a reativação automática agendada — não há dupla reativação
- **RN-FEAT-EP-0003-0001-0061-03:** O histórico de auditoria deve diferenciar claramente: `REATIVAR_MANUAL` (ação do admin) vs `REATIVAR_AUTOMATICO` (gatilho por data)
- **RN-FEAT-EP-0003-0001-0061-04:** Usuário reativado manualmente herda exatamente as mesmas permissões, papéis e vinculações que possuía antes da suspensão temporária

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0003-0001 — Cadastro e Convite de Usuários](../features/FEAT-EP-0003-0001-cadastro-e-convite-de-usuarios.md) | **Épico:** [EP-0003 — Gestao de Usuarios e Permissoes RBAC](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md)
