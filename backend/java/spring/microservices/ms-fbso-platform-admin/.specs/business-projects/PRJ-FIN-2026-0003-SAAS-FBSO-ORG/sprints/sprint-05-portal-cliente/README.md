# Sprint 5 — Portal do Cliente e Onboarding

- **Sprint:** 5 de 7
- **Marco:** M5 (EP-04a)
- **Branch:** `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-05-portal-cliente`
- **Status:** 🔄 Em andamento — Frente 0 concluída ✅ (7/36 tarefas, 19%)

---

## Estrutura de Frentes

| Frente | Escopo | Tasks | Quando |
|:---|:---|:---:|:---|
| **Frente 0** | 🔴 Bloqueantes (pré-sprint) | T-133.DT-095 a T-138.DT-100 (6) | ANTES da Frente 3 |
| **Frente 1** | 🟡 Recomendados | T-139.DT-023 a T-148.DT-102 (10) | Durante a sprint |
| **Frente 2** | 🔵 Desejáveis (opcional) | T-149.DT-086 a T-156.DT-113 (8) | Se houver capacidade |
| **Frente 3** | 🎯 Features da Sprint | T-057 a T-068 (12) | Corpo da sprint |

---

## Artefatos da Sprint

| Documento | Descrição |
|:---|:---|
| [SPRINT-CARD.md](./SPRINT-CARD.md) | Cartão da sprint com backlog detalhado |
| [SPRINT-TEST-SUITE.md](./SPRINT-TEST-SUITE.md) | Suite de testes (28 cenários) |
| [SPRINT-REVIEW.md](./SPRINT-REVIEW.md) | Template de review (a preencher) |
| [IDENTIFIED-TECHNICAL-DEBT-sprint-05-portal-cliente.md](./IDENTIFIED-TECHNICAL-DEBT-sprint-05-portal-cliente.md) | Auditoria técnica: 42 débitos identificados |

---

## Features (Frente 3)

| Feature | Descrição | Tasks | RNs |
|:---|:---|:---:|:---|
| **F04-01** | Login e Autenticação (Keycloak OIDC) | T-057, T-058, T-059 | RN13-01, RN13-02, RN13-03 |
| **F04-02** | Onboarding Guiado (4 passos) | T-060, T-061, T-062 | RN14-01, RN14-02, RN14-03, RN14-04 |
| **F04-03** | Dashboard do Cliente | T-063, T-064 | RN15-01, RN15-02 |
| **F04-04** | App Switcher | T-065, T-066 | RN16-01, RN16-02 |

---

## Débitos Técnicos

- **42 débitos identificados** pela auditoria multidisciplinar (9 skills)
- **24 serão tratados** na Sprint 5 (6 Frente 0 + 10 Frente 1 + 8 Frente 2)
- **16 postergados** para Sprints 6-7
- **28 débitos pendentes** de sprints anteriores (backlog)

📖 Detalhes completos em: [IDENTIFIED-TECHNICAL-DEBT-sprint-05-portal-cliente.md](./IDENTIFIED-TECHNICAL-DEBT-sprint-05-portal-cliente.md)

---

## Métricas

| Métrica | Meta |
|:---|:---:|
| Total de tarefas | 36 |
| Must Have | 24 |
| Should Have | 2 (F04-03) |
| Could Have | 10 |
| Débitos tratados | 24 |
| Cenários de teste | 33 |
| Dias-homem estimados | ~26.5d |

---

## Dependências

- **Pré-requisitos:** Sprint 2 (JwtAuthenticationFilter), Sprint 3 (TenantRepository), Sprint 4 (UserRepository, BusinessUnit entity)
- **Frente 0 é bloqueante:** docker-compose, Flyway, PostgreSQL driver, JWT claims, OAuth2 Client devem ser resolvidos antes da Frente 3
- **Sucessor:** Sprint 6 (BUs e Catálogo) — depende de BusinessUnit criada durante onboarding

---

## Quick Links

- [TASKS.md](../../TASKS.md) — Master task list (v3.2, 160 tasks)
- [PRD.md](../../PRD.md) — Product Requirements Document (v1.14)
- [SPECS.md](../../SPECS.md) — Technical Specifications (v2.3)
- [ARCHITECTURE.md](../../ARCHITECTURE.md) — Architecture Document (v2.7)
- [TEST_PLAN.md](../../TEST_PLAN.md) — Test Plan (v3.0)

---

🤖 *README gerado em 2026-07-17. Sprint 5 planejada com base na auditoria técnica IDENTIFIED-TECHNICAL-DEBT-sprint-05-portal-cliente.md.*
