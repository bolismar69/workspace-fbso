# SPRINT-BACKLOG — Índice de Tarefas Técnicas

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Data de Criação:** 27 de Julho de 2026
- **Versão:** 1.0
- **Status:** Em construção — Sprint 00 em andamento
- **Documento Vinculado:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — matriz de rastreabilidade US ↔ Feature ↔ Épico

---

## 1. Objetivo

Este documento é o **índice mestre de tarefas técnicas** do projeto. Cada entrada vincula uma tarefa (`T-NNN`) a uma User Story (`US-ID`) que, por sua vez, referencia Feature e Épico no [05-USER-STORIES](../../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md). A coluna `SPRINT-ALVO` indica em qual sprint a tarefa está alocada ou planejada.

**Fluxo:** Feature/US (05-USER-STORIES) → Refinamento Técnico → Tarefa (SPRINT-BACKLOG) → Sprint ativa

---

## 2. Status de Tarefa (Scrum/Kanban)

| Ícone | Código | Descrição |
|---|---|---|
| 📋 | `TODO` | Tarefa identificada, aguardando sprint ativa |
| 🏃 | `IN-PROGRESS` | Em execução pelo responsável |
| 👀 | `IN-REVIEW` | Code review ou revisão de artefato |
| 🧪 | `IN-TESTING` | Validação pelo QA |
| ✅ | `DONE` | Tarefa finalizada |
| 🚫 | `BLOCKED` | Impedimento externo — requer ação |

---

## 3. Backlog de Tarefas

| TASK-ID | TASK-DESCRIÇÃO | SPRINT-ALVO | US-ID | STATUS | DATA-INICIO | DATA-ENTREGA |
|---|---|---|---|---|---|---|
| T-000001 | Auditar backend existente vs SPECS-DEFINITION §6.1 — levantar gaps | Sprint 00 | — | `TODO` | | |
| T-000002 | Mover Docker Compose de `backend/` → `infra/docker/` | Sprint 00 | — | `TODO` | | |
| T-000003 | Mover migrations de `backend/` → `data_engineering/databases/db-postgresql/schema_fbso_platform/` | Sprint 00 | — | `TODO` | | |
| T-000004 | Extrair OpenAPI YAML dos controllers Spring existentes (SpringDoc) | Sprint 00 | — | `TODO` | | |
| T-000005 | Scaffold Frontend: Next.js 15 + React 19 + Tailwind CSS 4 | Sprint 00 | — | `TODO` | | |
| T-000006 | Configurar MSW Mock para endpoints EP-0001 e EP-0002 | Sprint 00 | — | `TODO` | | |
| T-000007 | Validar integração: `docker compose up` → backend healthy → frontend renderiza mock | Sprint 00 | — | `TODO` | | |
| T-000010 | Auditar/finalizar endpoint `GET /dashboard/admin/summary` + service | Sprint 01 | US-FEAT-EP-0001-0001-0001 | `TODO` | | |
| T-000011 | Auditar/finalizar endpoint `GET /dashboard/admin/evolution` | Sprint 01 | US-FEAT-EP-0001-0001-0003 | `TODO` | | |
| T-000012 | Auditar/finalizar endpoint `GET /dashboard/admin/accounts-by-status` + filtros | Sprint 01 | US-FEAT-EP-0001-0001-0001, US-FEAT-EP-0001-0001-0002 | `TODO` | | |
| T-000013 | Auditar/finalizar endpoint `GET /dashboard/admin/alerts` | Sprint 01 | US-FEAT-EP-0001-0003-0006 | `TODO` | | |
| T-000014 | Criar página Dashboard Admin no frontend (metrics cards + gráficos) | Sprint 01 | US-FEAT-EP-0001-0001-0001, US-FEAT-EP-0001-0001-0002, US-FEAT-EP-0001-0001-0003 | `TODO` | | |
| T-000015 | Criar tela de Lista de Contas com paginação, busca e filtros | Sprint 01 | US-FEAT-EP-0001-0002-0004 | `TODO` | | |
| T-000016 | Criar tela de Detalhe da Conta (dados cadastrais, status, assinatura) | Sprint 01 | US-FEAT-EP-0001-0002-0005 | `TODO` | | |
| T-000017 | Implementar indicadores visuais de atenção (badges, cores) no frontend | Sprint 01 | US-FEAT-EP-0001-0003-0007 | `TODO` | | |
| T-000018 | Testes de integração: dashboard admin + lista de contas | Sprint 01 | US-FEAT-EP-0001-0001-0001, US-FEAT-EP-0001-0002-0004 | `TODO` | | |
| T-000019 | Testes E2E (Playwright): fluxo dashboard → busca → detalhe conta | Sprint 01 | US-FEAT-EP-0001-0001-0001, US-FEAT-EP-0001-0002-0004, US-FEAT-EP-0001-0002-0005 | `TODO` | | |
| T-000020 | [Placeholder] — a preencher no planejamento do Sprint 02 | Sprint 02 | — | `TODO` | | |
| T-000030 | [Placeholder] — a preencher no planejamento do Sprint 03 | Sprint 03 | — | `TODO` | | |
| T-000040 | [Placeholder] — a preencher no planejamento do Sprint 04 | Sprint 04 | — | `TODO` | | |
| T-000050 | [Placeholder] — a preencher no planejamento do Sprint 05 | Sprint 05 | — | `TODO` | | |
| T-000060 | [Placeholder] — a preencher no planejamento do Sprint 06 | Sprint 06 | — | `TODO` | | |
| T-000070 | [Placeholder] — a preencher no planejamento do Sprint 07 | Sprint 07 | — | `TODO` | | |
| T-000080 | [Placeholder] — a preencher no planejamento do Sprint 08 | Sprint 08 | — | `TODO` | | |
| T-000090 | [Placeholder] — a preencher no planejamento do Sprint 09 | Sprint 09 | — | `TODO` | | |
| T-000100 | [Placeholder] — a preencher no planejamento do Sprint 10 | Sprint 10 | — | `TODO` | | |

---

## 4. Resumo por Sprint

| Sprint | Tarefas | US Vinculadas | Status |
|---|---|---|---|
| Sprint 00 | T-000001 a T-000007 (7) | — (infra) | 📋 Planejado |
| Sprint 01 | T-000010 a T-000019 (10) | US-FEAT-EP-0001-0001-0001, US-FEAT-EP-0001-0001-0002, US-FEAT-EP-0001-0001-0003, US-FEAT-EP-0001-0002-0004, US-FEAT-EP-0001-0002-0005, US-FEAT-EP-0001-0003-0006, US-FEAT-EP-0001-0003-0007 | 📋 Planejado |
| Sprint 02-10 | T-000020 a T-000100 (9) | A definir | 📋 Placeholder |

---

## 5. Referências

| Documento | Relação |
|---|---|
| [05-USER-STORIES](../../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | Matriz de rastreabilidade US ↔ Feature ↔ Épico |
| [questions.md](../questions.md) | Perguntas técnicas e entregáveis do Sprint 00 |
| [MILESTONES](../PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md) | Roadmap M1-M7 com datas-alvo |
| [SPECS-DEFINITION](../PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md) | Convenções técnicas que as tarefas devem seguir |

---

## Histórico de Alterações

| Versão | Data | Alteração | Autor |
|---|---|---|---|
| 1.0 | 27/07/2026 | Criação inicial: lista única com SPRINT-ALVO, 7 tarefas Sprint 00 + 10 tarefas Sprint 01 + placeholders Sprint 02-10 | Time Técnico |

---

🤖 *Documento gerado como parte da Fase 14 do Roadmap de Definições Técnicas. Ponte entre o 05-USER-STORIES (negócio) e a execução técnica.*
