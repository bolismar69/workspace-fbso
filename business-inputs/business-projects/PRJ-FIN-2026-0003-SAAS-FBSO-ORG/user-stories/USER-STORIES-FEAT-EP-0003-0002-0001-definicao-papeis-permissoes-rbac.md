# User Stories: Definição de Papéis e Permissões (RBAC)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** FEAT-EP-0003-0002 — Definição de Papéis e Permissões (RBAC)
- **Épico:** EP-0003 — Governança de Acessos e Permissões
- **Prioridade:** Must Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

---

## Objetivo de Negócio

Estabelecer os papéis padrão da plataforma com conjuntos de permissões bem definidos, garantindo que cada usuário acesse apenas o que seu perfil permite.

---

## User Stories

### US-027 — Atribuição de Papel ao Usuário

**Como** Administrador do Tenant, **quero** atribuir um dos papéis padrão do MVP (Admin do Tenant, Gerente de Unidade, Operador) a cada usuário **para** definir seu nível de acesso na plataforma. Papel 'Auditor' documentado para fase futura — não disponível para atribuição nesta fase.

**Critérios de Aceitação:**
- Seletor de papel no cadastro e na edição do usuário
- Descrição de cada papel disponível como tooltip
- Alteração de papel registrada em auditoria

### US-028 — Conjunto Predefinido de Permissões por Papel

**Como** Gestor de Produto, **quero** que cada papel tenha um conjunto predefinido de permissões. MVP: 3 papéis — Admin do Tenant (acesso total), Gerente de Unidade (gerencia sua unidade), Operador (executa tarefas). Auditor (apenas visualiza — papel documentado para fase futura, não implementado no MVP).

**Critérios de Aceitação:**
- Permissões mapeadas conforme tabela de papéis (RN-FEAT-EP-0003-0002-0001)
- Permissões não são customizáveis por tenant nesta fase
- Papéis são os mesmos para todos os tenants

### US-029 — Acesso Total do Admin do Tenant

**Como** Administrador do Tenant, **quero** que ao atribuir o papel "Admin do Tenant" a um usuário, ele automaticamente tenha acesso a todas as Unidades de Negócio e todos os módulos do tenant.

**Critérios de Aceitação:**
- Admin do Tenant vê todas as unidades de negócio no seletor
- Admin do Tenant vê todos os módulos contratados no App Switcher
- Não é necessário configurar permissões individuais para Admin

### US-030 [FASE FUTURA — NÃO IMPLEMENTAR NESTA FASE] — Permissões do Auditor

**Status: Backlog — Fase Futura.** Este papel não será implementado no MVP. Documentado para referência de design. Schema e permissões previstos conforme tabela RN-FEAT-EP-0003-0002-0001.

**Como** Administrador do Tenant, **quero** que ao atribuir o papel "Auditor" a um usuário, ele possa visualizar todos os dados das unidades permitidas mas não possa criar, editar ou excluir nada.

**Critérios de Aceitação:**
- Botões de criação/edição/exclusão não visíveis para Auditor
- Menus de configuração não aparecem para Auditor
- Tentativa de acesso direto a funcionalidades de escrita é bloqueada

---

## Regras de Negócio

### RN-FEAT-EP-0003-0002-0001 — Tabela de Permissões por Papel

| Funcionalidade | Admin Tenant | Gerente BU | Operador BU | Auditor [Fase Futura] |
|---------------|-------------|-----------|------------|-----------------------|
| Dashboard | Ver | Ver | Ver | Ver |
| Unidades de Negócio | Criar, Editar, Ver | Ver (apenas sua) | Ver (apenas sua) | Ver |
| Catálogo de Produtos | Criar, Editar, Ver, Excluir | Criar, Editar, Ver | Ver | Ver |
| Usuários e Permissões | Criar, Editar, Ver, Excluir | — | — | — |
| Planos e Assinaturas | Ver (apenas seu plano) | — | — | Ver |
| Configurações Fiscais* | — | — | — | — |

> *Funcionalidades fiscais não fazem parte do escopo desta fase.

- **RN-FEAT-EP-0003-0002-0002:** Não é permitido rebaixar ou desativar o último usuário com papel Admin do Tenant. O sistema deve validar que sempre existe pelo menos um Admin ativo por tenant.
- **RN-FEAT-EP-0003-0002-0003:** O primeiro usuário Administrador do Tenant é criado pelo time FBSO.ORG no momento da ativação da conta (US-008/FEAT-EP-0002-0001), garantindo que o tenant sempre tenha pelo menos um Admin.

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Matriz de permissões (RN-FEAT-EP-0003-0002-0001) aplicada corretamente para todos os 3 papéis (MVP). Papel Auditor documentado para fase futura. | Checklist de papel × funcionalidade verificado |
| F2 | Admin do Tenant com acesso total implícito (sem configuração manual) | Login como Admin com todas as unidades e módulos visíveis |
| F3 | Auditor sem nenhum botão de ação visível e bloqueado em acesso direto | Teste de cada funcionalidade como Auditor |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*
