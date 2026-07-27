# User Stories: Gestão de Status do Tenant

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** FEAT-EP-0002-0002 — Gestão de Status do Tenant
- **Épico:** EP-0002 — Gestão de Clientes e Assinaturas
- **Prioridade:** Must Have
- **Data-Alvo:** 31/08/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

---

## Objetivo de Negócio

Controlar o ciclo de vida de cada conta de cliente, permitindo ativar, suspender e reativar conforme a situação operacional ou comercial.

---

## User Stories

### US-012 — Alteração de Status da Conta

**Como** Administrador FBSO, **quero** alterar o status de uma conta de cliente entre os estados: Pendente Onboarding, Ativo, Suspenso, Inativo **para** refletir a situação real da conta.

**Critérios de Aceitação:**
- Seletor de status na tela de detalhes do Tenant
- Transições permitidas seguem regra de negócio RN-FEAT-EP-0002-0002-0001
- Ao alterar para Suspenso, campo de motivo é obrigatório

### US-013 — Bloqueio de Acesso na Suspensão

**Como** Administrador FBSO, **quero** que ao suspender uma conta, todos os usuários daquele tenant tenham o acesso ao portal bloqueado imediatamente.

**Critérios de Aceitação:**
- Bloqueio efetivo em até 5 minutos após a suspensão
- Usuários logados recebem mensagem de sessão encerrada na próxima requisição ao servidor (refresh de página ou chamada de API)
- Status do tenant atualizado em tempo real para "Suspenso"

### US-014 — Histórico de Mudanças de Status

**Como** Administrador FBSO, **quero** visualizar o histórico de mudanças de status de cada conta (quando foi ativada, suspensa, reativada e por quem).

**Critérios de Aceitação:**
- Linha do tempo de status na tela de detalhes do Tenant
- Cada evento exibe: status anterior → novo status, responsável, data/hora, motivo (se aplicável)

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN-FEAT-EP-0002-0002-0001** | Transições permitidas: Pendente Onboarding→Ativo, Pendente Onboarding→Inativo (abandono de onboarding), Ativo→Suspenso, Suspenso→Ativo, Suspenso→Inativo (encerramento de conta suspensa), Ativo→Inativo, Inativo→Ativo |
| **RN-FEAT-EP-0002-0002-0002** | Suspensão exige motivo registrado (campo obrigatório) |
| **RN-FEAT-EP-0002-0002-0003** | Reativação de conta suspensa restaura as permissões anteriores dos usuários |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Todas as transições de status permitidas são executáveis e as proibidas são bloqueadas | Matriz de transições testada |
| F2 | Bloqueio de acesso em até 5 minutos após suspensão | Log de timestamp da suspensão vs. bloqueio |
| F3 | Linha do tempo de status completa e imutável | Print da timeline com múltiplos eventos |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*
