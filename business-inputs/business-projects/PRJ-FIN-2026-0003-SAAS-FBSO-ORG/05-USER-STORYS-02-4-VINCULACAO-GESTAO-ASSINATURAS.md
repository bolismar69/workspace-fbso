# User Stories: Vinculação e Gestão de Assinaturas

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** F02-04 — Vinculação e Gestão de Assinaturas
- **Épico:** EP-02 — Gestão de Clientes e Assinaturas
- **Prioridade:** Must Have
- **Data-Alvo:** 31/08/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES.md](../04-FEATURES.md)

---

## Objetivo de Negócio

Vincular clientes a planos com controle de vigência, permitindo upgrade/downgrade e acompanhamento do status da assinatura.

---

## User Stories

### US-019 — Vinculação de Cliente a Plano

**Como** Administrador FBSO, **quero** vincular um cliente a um plano comercial definindo data de início, vigência e status da assinatura.

**Critérios de Aceitação:**
- Seletor de plano (apenas planos ativos listados)
- Campos: data de início, data de término (ou "indeterminado"), status (Ativa, Suspensa, Cancelada)
- Ao ativar assinatura, módulos do plano são liberados para o tenant

### US-020 — Upgrade e Downgrade de Plano

**Como** Administrador FBSO, **quero** realizar upgrade ou downgrade de plano de um cliente, mantendo o histórico da assinatura anterior. (Na prática, funcionalidade será utilizada pelo Líder Comercial — perfil de stakeholder, não papel RBAC do sistema.)

**Critérios de Aceitação:**
- Ao trocar de plano, assinatura anterior é finalizada com data de término = data atual. Nova assinatura é criada com data de início = data atual (mesmo dia), garantindo que o tenant não fique sem assinatura ativa durante a transição (RN07-02).
- Histórico exibe todas as assinaturas do cliente em ordem cronológica

### US-021 — Suspensão de Assinatura com Bloqueio de Módulos

**Como** Administrador FBSO, **quero** suspender a assinatura de um cliente, o que deve bloquear o acesso dele aos módulos do plano.

**Critérios de Aceitação:**
- Suspensão da assinatura bloqueia acesso aos módulos em até 5 minutos
- Status da assinatura alterado para "Suspensa"
- Reativação da assinatura restaura acesso aos mesmos módulos

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN07-01** | Um tenant pode ter apenas uma assinatura ativa por vez |
| **RN07-02** | Upgrade/downgrade não pode deixar o tenant sem assinatura ativa durante a transição |
| **RN07-03** | Data de término de assinatura é opcional (planos sem data de término = vigência contínua) |
| **RN07-04** | Quando o Tenant está Suspenso, o acesso aos módulos é bloqueado independentemente do status da Assinatura. Quando a Assinatura está Suspensa e o Tenant está Ativo, o acesso aos módulos é bloqueado. Ambos devem estar ativos para liberação de acesso. |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Cliente vinculado a plano com módulos liberados | Login do cliente com módulos visíveis no App Switcher |
| F2 | Upgrade/downgrade com histórico preservado | Timeline de assinaturas com transições registradas |
| F3 | Suspensão bloqueia acesso em até 5 minutos | Log de timestamp da suspensão vs. bloqueio |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*
