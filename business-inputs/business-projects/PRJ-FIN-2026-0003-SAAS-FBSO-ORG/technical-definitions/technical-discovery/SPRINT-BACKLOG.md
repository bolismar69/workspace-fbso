# SPRINT-BACKLOG — Índice de Tarefas Técnicas

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Fase:** F18 — Bloco D (Sprints — Technical Discovery)
- **Versão:** 2.0 · **Data:** 30/07/2026 · **Status:** CREATED
- **Documento Vinculado:** [05-USER-STORIES](../../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

---

## 1. Objetivo

Índice mestre de tarefas técnicas. Cada entrada vincula T-NNN → US-ID → SPRINT-ALVO → CONTRACTS.

## 2. Status de Tarefa

| Ícone | Código | Descrição |
|---|---|---|
| 📋 | `TODO` | Aguardando sprint ativa |
| 🏃 | `IN-PROGRESS` | Em execução |
| 👀 | `IN-REVIEW` | Code review |
| 🧪 | `IN-TESTING` | Validação QA |
| ✅ | `DONE` | Finalizada |
| 🚫 | `BLOCKED` | Impedimento |

## 3. Backlog de Tarefas

| TASK-ID | TASK-DESCRIÇÃO | SPRINT-ALVO | US-ID | RESPONSÁVEL | STATUS | CONTRACTS |
|---|---|---|---|---|---|---|---|
| T-000001 | Auditar backend vs SPECS-DEFINITION §6.1 | Sprint 00 | — | A definir | `TODO` | — |
| T-000002 | Mover Docker Compose → `infra/docker/` | Sprint 00 | — | Davi Silva (DevOps) | `TODO` | — |
| T-000003 | Mover migrations → `data_engineering/` | Sprint 00 | — | Carlos Caldas (DB) | `TODO` | [DATA](sprint-00/CONTRACTS-DATA-sprint-00.md) |
| T-000010 | Endpoint `GET /dashboard/admin/summary` | Sprint 01 | US-FEAT-EP-0001-0001-0001 | Bolismar Oliveira (BE) | `TODO` | [API](sprint-01/CONTRACTS-API-sprint-01.md) |
| T-000011 | Endpoint `GET /dashboard/admin/evolution` | Sprint 01 | US-FEAT-EP-0001-0001-0003 | Bolismar Oliveira (BE) | `TODO` | [API](sprint-01/CONTRACTS-API-sprint-01.md) |
| T-000014 | Criar página Dashboard Admin (métricas+gráficos) | Sprint 01 | US-FEAT-EP-0001-0001-0001 | Bruno Gratto (FE) | `TODO` | [API](sprint-01/CONTRACTS-API-sprint-01.md) |
| T-000018 | Testes integração: dashboard + lista contas | Sprint 01 | US-FEAT-EP-0001-0001-0001 | Felipe Canedas (QA) | `TODO` | [SRE](sprint-01/CONTRACTS-SRE-sprint-01.md) |
| T-000020 | [Placeholder] Sprint 02 | Sprint 02 | — | A definir | `TODO` | — |
| T-000030 | [Placeholder] Sprint 03 | Sprint 03 | — | A definir | `TODO` | — |

## 4. Resumo por Sprint

| Sprint | Tarefas | US Vinculadas | Status |
|---|---|---|---|
| Sprint 00 | T-000001 a T-000003 (3) | — (infra) | 📋 Planejado |
| Sprint 01 | T-000010 a T-000019 (10) | US-EP-0001-* | 📋 Planejado |
| Sprint 02-10 | T-000020 a T-000100 (9) | A definir | 📋 Placeholder |

## 5. Referências

| Documento | Relação | Fase |
|---|---|---|
| [TEAM-SKILLS-MAP](../PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md) | Skills do time executor | F5 |
| [ARCHITECTURE-DEFINITION](../PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md) | C4, ADRs, integração | F7 |
| [SECURITY-DEFINITION](../PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md) | Threat model, IAM, compliance | F8 |
| [DATA-ARCHITECTURE](../PROJECT-TECHNICAL-DEFINITIONS-DATA-ARCHITECTURE-DEFINITION.md) | ERD, schemas, pipelines | F9 |
| [DEVOPS-SRE](../PROJECT-TECHNICAL-DEFINITIONS-DEVOPS-SRE-DEFINITION.md) | CI/CD, IaC, observabilidade, SLOs | F10 |
| [TEST-STRATEGY](../PROJECT-TECHNICAL-DEFINITIONS-TEST-STRATEGY-DEFINITION.md) | Pirâmide de testes, quality gates | F11 |
| [INFRA-CLOUD](../PROJECT-TECHNICAL-DEFINITIONS-INFRA-CLOUD-DEFINITION.md) | Topologia, networking, DR | F12 |
| [SPECS-DEFINITION](../PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md) | Convenções técnicas cross-solution | F16 |
| [MILESTONES](../PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md) | Roadmap M1-M7 | F17 |

🤖 *F18 — SPRINT-BACKLOG · technical-discovery/ · Roadmap v5.0*
