# User Stories: Histórico de Auditoria Administrativa

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** FEAT-EP-0002-0005 — Histórico de Auditoria Administrativa
- **Épico:** EP-0002 — Gestão de Clientes e Assinaturas
- **Prioridade:** Must Have
- **Data-Alvo:** 31/08/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

---

## Objetivo de Negócio

Registrar todas as ações administrativas realizadas pelo time interno, garantindo rastreabilidade e conformidade.

---

## User Stories

### US-022 — Registro Automático de Auditoria

**Como** Administrador FBSO, **quero** que toda ação de criação, alteração de status, mudança de plano e edição de dados de tenant seja automaticamente registrada em um histórico de auditoria.

**Critérios de Aceitação:**
- Registro inclui: tipo da ação, administrador responsável, data/hora, dados anteriores e novos. Ações de criação registram apenas dados novos. Ações de edição, alteração de status, mudança de plano e alteração de permissões registram dados anteriores e novos.
- Histórico acessível na tela de detalhes do Tenant
- Histórico não pode ser editado ou apagado

### US-023 — Filtro do Histórico de Auditoria

**Como** Administrador FBSO, **quero** filtrar o histórico de auditoria por período e por tipo de ação **para** localizar eventos específicos. (Papel 'Auditor' previsto para fase futura — nesta fase, Administrador FBSO consulta o histórico de auditoria.)

**Critérios de Aceitação:**
- Filtros disponíveis: período (data inicial e final), tipo de ação (criação, edição, alteração de status, mudança de plano, alteração de permissões — cobertura completa conforme RN-FEAT-EP-0002-0005-0001)
- Resultados ordenados do mais recente para o mais antigo

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN-FEAT-EP-0002-0005-0001** | Auditoria cobre 100% das ações administrativas (criação, edição, alteração de status, mudança de plano, alteração de permissões) |
| **RN-FEAT-EP-0002-0005-0002** | Registros de auditoria são imutáveis — não podem ser editados ou excluídos por nenhum usuário |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | 100% das ações administrativas registradas em auditoria | Amostra de ações vs. registros de auditoria |
| F2 | Registros imutáveis — tentativa de edição bloqueada | Teste de edição/remoção de registro |
| F3 | Filtro por período e tipo de ação funcional | Consulta filtrada com resultados corretos |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*
