# User Story: US-FEAT-EP-0003-0001-0060 — suspender temporariamente um usuário definindo período de ausência

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0003](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md) ➔ Feature [FEAT-EP-0003-0001](../features/FEAT-EP-0003-0001-cadastro-e-convite-de-usuarios.md) ➔ User Story US-FEAT-EP-0003-0001-0060
- **Prioridade:** Should Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.0 — Gap Analysis #12 (27/07/2026)
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador do Tenant,
- **quero** suspender temporariamente um usuário definindo data de início e data de retorno prevista, com um motivo (férias, licença, afastamento),
- **para** gerenciar ausências programadas sem precisar desativar e recriar o usuário depois.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal — Suspensão temporária com prazo]
- **Dado que** o Administrador do Tenant está autenticado e acessa os detalhes de um usuário ativo,
- **Quando** seleciona "Suspender Temporariamente", define a data de início, a data prevista de retorno, seleciona o motivo (Férias, Licença, Afastamento, Outros) e confirma,
- **Então** o sistema deve: **bloquear o acesso do usuário a partir da data de início informada, exibir o status "Suspenso Temporariamente" com o motivo e a data de retorno na lista de usuários, e não exigir nenhuma ação no dia do retorno para reativar**.

### Cenário 2: [Validação — Datas inválidas]
- **Dado que** o Administrador do Tenant preenche o formulário de suspensão temporária,
- **Quando** informa uma data de retorno anterior à data de início ou uma data de início no passado,
- **Então** o sistema deve: **recusar e exibir mensagens de validação específicas para cada campo inválido**.

### Cenário 3: [Usuário suspenso tenta login]
- **Dado que** um usuário com status "Suspenso Temporariamente" tenta fazer login durante o período de suspensão,
- **Quando** informa credenciais válidas,
- **Então** o sistema deve: **recusar o acesso e exibir mensagem informando que a conta está suspensa temporariamente, com a data prevista de retorno, orientando contato com o administrador se necessário**.

### Cenário 4: [Reativação automática na data de retorno]
- **Dado que** a data prevista de retorno de um usuário suspenso temporariamente é atingida,
- **Quando** o sistema processa a virada de data,
- **Então** o sistema deve: **alterar automaticamente o status do usuário para "Ativo", restaurar o acesso normalmente e registrar a reativação automática no histórico de auditoria**.

### Cenário 5: [Auditoria]
- **Dado que** o Administrador do Tenant suspende temporariamente um usuário,
- **Quando** a operação é concluída,
- **Então** o sistema deve: **registrar no histórico de auditoria: administrador responsável, data/hora, usuário afetado, período (início e retorno), motivo e tipo da ação (SUSPENDER_TEMPORARIAMENTE)**.

---

## 3. Regras de Negócio de Tela Relacionadas

- **RN-FEAT-EP-0003-0001-0060-01:** Apenas usuários com perfil `admin-do-tenant` podem suspender temporariamente outros usuários
- **RN-FEAT-EP-0003-0001-0060-02:** Um `admin-do-tenant` não pode suspender a si mesmo temporariamente
- **RN-FEAT-EP-0003-0001-0060-03:** O campo "motivo" é obrigatório, com opções predefinidas: Férias, Licença, Afastamento, Outros (com campo texto livre)
- **RN-FEAT-EP-0003-0001-0060-04:** A data de retorno é obrigatória e deve ser posterior à data de início em pelo menos 1 dia
- **RN-FEAT-EP-0003-0001-0060-05:** A reativação na data de retorno é automática — o administrador não precisa executar nenhuma ação manual
- **RN-FEAT-EP-0003-0001-0060-06:** Um administrador pode reativar manualmente um usuário suspenso temporariamente antes da data prevista de retorno (ver US-0061)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0003-0001 — Cadastro e Convite de Usuários](../features/FEAT-EP-0003-0001-cadastro-e-convite-de-usuarios.md) | **Épico:** [EP-0003 — Gestao de Usuarios e Permissoes RBAC](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md)
