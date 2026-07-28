# User Story: US-FEAT-EP-0003-0001-0059 — editar os dados básicos de um usuário (nome, e-mail) para manter o cadastro atualizado

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0003](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md) ➔ Feature [FEAT-EP-0003-0001](../features/FEAT-EP-0003-0001-cadastro-e-convite-de-usuarios.md) ➔ User Story US-FEAT-EP-0003-0001-0059
- **Prioridade:** Must Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.0 — Gap Analysis #11 (27/07/2026)
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador do Tenant,
- **quero** editar os dados básicos de um usuário (nome completo, e-mail) quando necessário,
- **para** manter o cadastro de usuários do meu tenant sempre atualizado sem depender do suporte da FBSO.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal — Editar nome e e-mail]
- **Dado que** o Administrador do Tenant está autenticado e visualiza os detalhes de um usuário,
- **Quando** aciona "Editar" e altera o nome completo e/ou e-mail do usuário, salvando as modificações,
- **Então** o sistema deve: **atualizar os dados imediatamente e, se o e-mail foi alterado, notificar o usuário no novo e-mail sobre a mudança**.

### Cenário 2: [Validação — E-mail duplicado]
- **Dado que** o Administrador do Tenant tenta alterar o e-mail de um usuário para um endereço já cadastrado no mesmo tenant,
- **Quando** submete o formulário,
- **Então** o sistema deve: **recusar a alteração e exibir a mensagem "Este e-mail já está em uso por outro usuário deste tenant"**.

### Cenário 3: [Restrição — Não alterar o próprio perfil administrativo]
- **Dado que** o Administrador do Tenant tenta alterar seu próprio e-mail ou rebaixar seu próprio papel,
- **Quando** acessa a tela de edição do próprio perfil,
- **Então** o sistema deve: **permitir a edição de nome e e-mail do próprio administrador, mas bloquear a alteração do próprio papel administrativo** (regra de segurança: o tenant nunca pode ficar sem admin).

### Cenário 4: [Auditoria]
- **Dado que** o Administrador do Tenant edita os dados de um usuário,
- **Quando** a alteração é salva,
- **Então** o sistema deve: **registrar no histórico de auditoria: administrador responsável, data/hora, campos alterados (antes → depois)**.

---

## 3. Regras de Negócio de Tela Relacionadas

- **RN-FEAT-EP-0003-0001-0059-01:** Apenas usuários com perfil `admin-do-tenant` podem editar dados de outros usuários
- **RN-FEAT-EP-0003-0001-0059-02:** O campo "e-mail" deve ser validado como formato de e-mail válido
- **RN-FEAT-EP-0003-0001-0059-03:** E-mail não pode ser duplicado dentro do mesmo tenant ativo
- **RN-FEAT-EP-0003-0001-0059-04:** O campo "nome completo" é obrigatório e deve ter no mínimo 3 caracteres

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0003-0001 — Cadastro e Convite de Usuários](../features/FEAT-EP-0003-0001-cadastro-e-convite-de-usuarios.md) | **Épico:** [EP-0003 — Gestao de Usuarios e Permissoes RBAC](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md)
