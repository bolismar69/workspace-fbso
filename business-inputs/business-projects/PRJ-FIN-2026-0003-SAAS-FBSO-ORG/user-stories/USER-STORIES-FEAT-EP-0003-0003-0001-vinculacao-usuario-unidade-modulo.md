# User Stories: Vinculação Usuário × Unidade de Negócio × Módulo

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** F03-03 — Vinculação Usuário × Unidade × Módulo
- **Épico:** EP-03 — Governança de Acessos e Permissões
- **Prioridade:** Must Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES.md](../04-FEATURES.md)

---

## Objetivo de Negócio

Permitir controle granular de acesso, definindo exatamente quais unidades de negócio e quais módulos cada usuário pode acessar.

---

## User Stories

### US-031 — Definição de Unidades de Negócio Vinculadas

**Como** Administrador do Tenant, **quero** definir quais Unidades de Negócio um usuário pode acessar (uma, várias ou todas) **para** restringir seu escopo de atuação.

**Critérios de Aceitação:**
- No cadastro/edição do usuário, lista de Unidades de Negócio com checkbox
- Permite selecionar "Todas" ou unidades específicas
- Para Admin do Tenant, "Todas" é fixo e não pode ser alterado

### US-032 — Definição de Módulos Vinculados

**Como** Administrador do Tenant, **quero** definir quais módulos/produtos um usuário pode acessar (ex: apenas Storekeeper, apenas Tributali-Engine, ou ambos) **para** restringir sua visão da plataforma.

**Critérios de Aceitação:**
- No cadastro/edição do usuário, lista de módulos contratados pelo tenant
- Cada módulo com checkbox (marcado = acesso permitido)
- Usuário sem acesso a um módulo não o vê no App Switcher

### US-033 — Alteração de Vinculações com Efeito Imediato

**Como** Administrador do Tenant, **quero** alterar as vinculações de um usuário a qualquer momento (adicionar/remover unidade, adicionar/remover módulo) com efeito imediato.

**Critérios de Aceitação:**
- Alterações salvas têm efeito na próxima requisição ao servidor (refresh de página ou chamada de API). A sessão ativa é atualizada para refletir as novas permissões imediatamente.
- Se usuário estiver logado e tiver acesso a uma unidade removida, a sessão é ajustada
- Registro de auditoria gerado para cada alteração

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN11-01** | Um usuário deve ter pelo menos uma Unidade de Negócio vinculada para acessar o portal (exceto Admin do Tenant, que tem acesso implícito a todas) |
| **RN11-02** | Um usuário deve ter pelo menos um módulo vinculado para acessar o portal |
| **RN11-03** | A lista de módulos disponíveis para vinculação é determinada pelos módulos incluídos no plano contratado pelo tenant |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Usuário vê apenas unidades e módulos vinculados | Login com usuário restrito e verificação de visibilidade |
| F2 | Remoção de unidade tem efeito imediato na sessão ativa | Usuário logado perde acesso à unidade removida na próxima ação |
| F3 | Lista de módulos disponíveis para vinculação reflete o plano contratado | Comparação plano × lista de módulos exibida |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*
