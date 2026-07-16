# User Stories: Cadastro e Convite de Usuários

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** F03-01 — Cadastro e Convite de Usuários
- **Épico:** EP-03 — Governança de Acessos e Permissões
- **Prioridade:** Must Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES.md](../04-FEATURES.md)

---

## Objetivo de Negócio

Permitir que o administrador do tenant cadastre e convide usuários para acessar a plataforma, definindo seus acessos de forma granular.

---

## User Stories

### US-024 — Convite de Novo Usuário

**Como** Administrador do Tenant, **quero** convidar um novo usuário para a plataforma informando nome, e-mail e perfil de acesso.

**Critérios de Aceitação:**
- Formulário com campos: nome completo, e-mail, perfil de acesso (seletor de papel: Admin, Gerente, Operador — obrigatório)
- Sistema valida se e-mail já está cadastrado no mesmo tenant
- Convite enviado por e-mail com link para definição de senha

### US-025 — Lista de Usuários do Tenant

**Como** Administrador do Tenant, **quero** visualizar a lista de usuários do meu tenant com seus respectivos papéis, unidades vinculadas e status (ativo, inativo, convite pendente).

**Critérios de Aceitação:**
- Lista exibe: nome, e-mail, papel principal, unidades vinculadas, status
- Filtro por status: Todos, Ativos, Pendentes (convite não aceito), Inativos
- Indicador visual para convites ainda não aceitos

### US-026 — Desativação de Usuário

**Como** Administrador do Tenant, **quero** desativar um usuário para bloquear imediatamente seu acesso à plataforma.

**Critérios de Aceitação:**
- Botão "Desativar" na lista de usuários
- Confirmação exigida antes da desativação
- Usuário desativado não consegue fazer login
- Reativação possível a qualquer momento pelo Administrador do Tenant. Usuário não pode reativar a si mesmo.

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN09-01** | Convite de usuário expira em 7 dias se não aceito |
| **RN09-02** | E-mail deve ser único por tenant (não pode haver dois usuários com mesmo e-mail no mesmo tenant) |
| **RN09-03** | Administrador do tenant não pode desativar a si mesmo |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Convite enviado e usuário consegue definir senha e acessar | Login bem-sucedido do novo usuário |
| F2 | Lista de usuários com filtro por status e indicador de pendentes | Print da lista com todos os status |
| F3 | Usuário desativado bloqueado; reativação funcional | Teste de login pós-desativação e pós-reativação |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*
