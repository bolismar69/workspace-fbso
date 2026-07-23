# Sprints — ms-fbso-platform-admin

- **Projeto:** [PRJ-FIN-2026-0003-SAAS-FBSO-ORG](../)
- **Total de Sprints:** 7
- **Duração:** 14 semanas (24/07/2026 → 30/10/2026)
- **Documentos-mestre:** [TASKS.md](../TASKS.md) v3.8 | [SPECS.md](../SPECS.md) v2.7 | [TEST_PLAN.md](../TEST_PLAN.md) v3.3 | [PRD.md](../PRD.md) v1.18 | [ARCHITECTURE.md](../ARCHITECTURE.md) v2.11 | [SECURITY.md](../SECURITY.md) v1.2 | [TECHNICAL-REFERENCE.md](../TECHNICAL-REFERENCE.md) v1.1
- **🚫 Estratégia de Branching:** Uma branch por sprint — [ver detalhes no PRD §8.4](../PRD.md#84-estratégia-de-branching--uma-branch-por-sprint)

---

## Índice de Sprints

| Sprint | Datas | Marco | Tarefas | Artefatos |
|:---|:---|:---|:---:|:---|
| [**Sprint 1**](sprint-01-setup/) | 24/07 — 07/08 | Pre-M2 Setup | T-001 a T-008 (8) | [Card](sprint-01-setup/SPRINT-CARD.md) · [Testes](sprint-01-setup/SPRINT-TEST-SUITE.md) · [Review](sprint-01-setup/SPRINT-REVIEW.md) |
| [**Sprint 2**](sprint-02-seguranca/) | 07/08 — 15/08 | Pre-M2 Segurança | T-009 a T-015, T-015.1 (8) | [Card](sprint-02-seguranca/SPRINT-CARD.md) · [Testes](sprint-02-seguranca/SPRINT-TEST-SUITE.md) · [Review](sprint-02-seguranca/SPRINT-REVIEW.md) |
| [**Sprint 3 ✅**](sprint-03-portal-admin/) | 16/07 — 17/07 | M2 + M3 | 42/42 ✅ | [Card](sprint-03-portal-admin/SPRINT-CARD.md) · [Testes](sprint-03-portal-admin/SPRINT-TEST-SUITE.md) · [Review](sprint-03-portal-admin/SPRINT-REVIEW.md) · [Unified Report](sprint-03-portal-admin/SPRINT-3-EXECUTION-REPORT-UNIFIED.md) · [Tech Debt](sprint-03-portal-admin/IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md) |
| [**Sprint 4 🔄**](sprint-04-rbac/) | 31/08 — 15/09 | M4 | 48 tasks (Frentes 0-4 ✅, 31/48) | [Card](sprint-04-rbac/SPRINT-CARD.md) · [Testes](sprint-04-rbac/SPRINT-TEST-SUITE.md) · [Review](sprint-04-rbac/SPRINT-REVIEW.md) · [Exec Reports F0-F4](sprint-04-rbac/) · [Tech Debt](sprint-04-rbac/IDENTIFIED-TECHNICAL-DEBT-sprint-04-rbac.md) |
| [**Sprint 5 🔄**](sprint-05-portal-cliente/) | 15/09 — 30/09 | M5 | 40 tasks (Frentes 0-1-2-3a ✅, 34/40 85%) | [Card](sprint-05-portal-cliente/SPRINT-CARD.md) · [Testes](sprint-05-portal-cliente/SPRINT-TEST-SUITE.md) · [Review](sprint-05-portal-cliente/SPRINT-REVIEW.md) · [Dev Plan F1](sprint-05-portal-cliente/SPRINT-DEVELOPMENT-PLANNING-Frente-1.md) · [Dev Plan F2](sprint-05-portal-cliente/SPRINT-DEVELOPMENT-PLANNING-Frente-2.md) · [Dev Plan F3](sprint-05-portal-cliente/SPRINT-DEVELOPMENT-PLANNING-Frente-3.md) · [Exec Reports F0-F3](sprint-05-portal-cliente/) · [Code Reviews F1-F3](sprint-05-portal-cliente/) · [Tech Debt](sprint-05-portal-cliente/IDENTIFIED-TECHNICAL-DEBT-sprint-05-portal-cliente.md) |
| [**Sprint 6 🔄**](sprint-06-bus-catalogo/) | 30/09 — 15/10 | M6 | 18 tasks (F0 4/4 ✅ + F1 0/5 + M6 0/9) | [Card](sprint-06-bus-catalogo/SPRINT-CARD.md) · [Testes](sprint-06-bus-catalogo/SPRINT-TEST-SUITE.md) · [Review](sprint-06-bus-catalogo/SPRINT-REVIEW.md) · [Exec Report F0](sprint-06-bus-catalogo/SPRINT-6-EXECUTION-REPORT-Frente-0.md) · [Dev Plan F0](sprint-06-bus-catalogo/SPRINT-DEVELOPMENT-PLANNING-Frente-0.md) · [Tech Debt](sprint-06-bus-catalogo/IDENTIFIED-TECHNICAL-DEBT-sprint-06-bus-catalogo.md) |
| [**Sprint 7**](sprint-07-homologacao/) | 15/10 — 30/10 | M7 | T-071 a T-079 (9) | [Card](sprint-07-homologacao/SPRINT-CARD.md) · [Testes](sprint-07-homologacao/SPRINT-TEST-SUITE.md) · [Review](sprint-07-homologacao/SPRINT-REVIEW.md) |

---

## Estrutura de Cada Sprint

```
sprint-NN-nome/
├── SPRINT-CARD.md         ← 1 página: goal, tasks, DONE criteria, datas, responsável
├── SPRINT-TEST-SUITE.md   ← Extrato do TEST_PLAN.md com cenários desta sprint
└── SPRINT-REVIEW.md       ← Checklist de demonstração para o Product Owner
```

## Como Usar Este Diretório

1. **Início da sprint:** Leia o `SPRINT-CARD.md` para alinhar o time sobre o goal.
2. **Durante a sprint:** Use o `SPRINT-TEST-SUITE.md` como checklist de qualidade.
3. **Fim da sprint:** Execute o `SPRINT-REVIEW.md` na demonstração para o PO.
4. **Após a review:** Atualize o status dos cenários no documento-mestre `TEST_PLAN.md`.

## Relação com Documentos-Mestre

```
Documentos-Mestre (fonte da verdade)
    │
    ├── TASKS.md      ──extrai──→  SPRINT-CARD.md (tasks, estimativas, critérios DONE)
    ├── TEST_PLAN.md  ──extrai──→  SPRINT-TEST-SUITE.md (cenários de teste da sprint)
    └── SPECS.md      ──extrai──→  SPRINT-REVIEW.md (critérios de aceitação demonstráveis)
```

> **Regra:** Os documentos-mestre (`TASKS.md`, `TEST_PLAN.md`, `SPECS.md`) são a **fonte da verdade**. Os artefatos de sprint são **derivados** deles. Se houver conflito, os documentos-mestre prevalecem. Atualize ambos quando houver mudanças.

---

## Timeline Visual

```
SEMANA  1  2  3  4  5  6  7  8  9 10 11 12 13 14
        ├─────────┬──────┬─────────┬─────────┬─────────┬─────────┬─────────┤
        │ Sprint 1│Sprint│ Sprint 3│ Sprint 4│Sprint 5 │ Sprint 6│ Sprint 7│
        │ Setup   │  2   │ M2 + M3 │   M4    │  M5 🔄  │  M6 🔄  │   M7    │
        │         │ Seg. │         │         │ F0 ✅   │         │         │
        └─────────┴──────┴─────────┴─────────┴─────────┴─────────┴─────────┘
      24/07     07/08  15/08     31/08     15/09     30/09     15/10     30/10
```

---

## Matriz de Rastreabilidade

> **Atualizado:** 23/07/2026 · **Fonte:** [TASKS.md](../TASKS.md) v3.8 + [PRD.md](../PRD.md) v1.18
>
> Sprint 5 em andamento: Frentes 0-1-2-3a ✅ (36/40, 90%). Sprint 6 🔄: Frente 0 concluída ✅ (4/4). 261 testes (0 failures). CnpjValidator com algoritmo CNPJ alfanumérico (IN RFB 2.119/2022).
>
> Esta matriz é a **camada de integração** entre o planejamento ([TASKS.md](../TASKS.md)) e a execução (sprints/). Use-a como mapa único para navegar entre fases, sprints, tarefas, features e épicos.

| Fase | Sprint | Marco | Tasks | Qtd | Features | Épico | Status | Progresso | Atualização |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|:---:|
| **FASE 0** | Sprint 1 — Setup | Pre-M2 Setup | T-001 a T-008 | 8 | — | — | ✅ Concluída | 8/8 (100%) | 14/07/2026 |
| **FASE 0** | Sprint 2 — Segurança | Pre-M2 Segurança | T-009 a T-015, T-015.1 | 8 | — | — | ✅ Concluída | 8/8 (100%) | 14/07/2026 |
| **FASE 0** | Sprint 3 — Frente 0 | Pre-M2 Correções | T-015.2.DT-001 a T-015.13.DT-012 | 12 | — | — | ✅ Concluída | 12/12 (100%) | 17/07/2026 |
| **FASE 1** | Sprint 3 — Frente 1 | M2 (EP-01) | T-016 a T-023 | 8 | F01-01, F01-02, F01-03 | EP-01 | ✅ Concluída | 8/8 (100%) | 17/07/2026 |
| **FASE 2** | Sprint 3 — Frente 2 | M3 (EP-02) | T-024 a T-038 | 15 | F02-01 a F02-05 | EP-02 | ✅ Concluída | 15/15 (100%) | 17/07/2026 |
| **FASE 2** | Sprint 3 — Frente 3 | Durante-Sprint | T-039.DT-017 a T-045.DT-046 | 7 | — | — | ✅ Concluída | 7/7 (100%) | 17/07/2026 |
| **FASE 3** | Sprint 4 — RBAC | M4 (EP-03) | T-046 a T-056 + T-096 a T-132 | 48 | F03-01 a F03-04 | EP-03 | 🔄 Em andamento | 31/48 (65%) | 17/07/2026 |
| **FASE 4** | Sprint 5 — Frente 0 | Pre-M5 Bloqueantes | T-133.DT-095 a T-138.DT-100 | 6 | — | — | ✅ Concluída | 6/6 (100%) | 17/07/2026 |
| **FASE 4** | Sprint 5 — Frente 1 | Pre-M5 Recomendados | T-139.DT-023 a T-148.DT-102 | 10 | — | — | ✅ Concluída | 10/10 (100%) | 23/07/2026 |
| **FASE 4** | Sprint 5 — Frente 2 | Pre-M5 Desejáveis | T-149.DT-086 a T-156.DT-113 | 8 | — | — | ✅ Concluída | 8/8 (100%) | 23/07/2026 |
| **FASE 4** | Sprint 5 — Frente 3a | M5 Features Backend | T-057 a T-068 | 12 | F04-01 a F04-04 | EP-04 | ✅ Concluída | 12/12 (100%) | 23/07/2026 |
| **FASE 4** | Sprint 5 — Portal Cliente | M5 (EP-04a) | T-157 a T-160 | 4 | F04-01 a F04-04 | EP-04 | ⬜ Não iniciada | 0/4 (0%) | — |
| **FASE 5** | Sprint 6 — Frente 0 | Pre-M6 Bloqueantes | T-161.DT-126 a T-164.DT-129 | 4 | — | — | ✅ Concluída | 4/4 (100%) | 23/07/2026 |
| **FASE 5** | Sprint 6 — Frente 1 | Pre-M6 Recomendados | T-165.DT-130 a T-169.DT-137 | 5 | — | — | ✅ Concluída | 5/5 (100%) | 23/07/2026 |
| **FASE 5** | Sprint 6 — BUs e Catálogo | M6 (EP-04b) | T-069 a T-077 | 9 | F04-05, F04-06 | EP-04 | ⬜ Não iniciada | 0/9 (0%) | — |
| **FASE 6** | Sprint 7 — Homologação | M7 (Homologação) | T-071 a T-079 | 9 | Todas (18) | Todos | ⬜ Não iniciada | 0/9 (0%) | — |
| **TOTAL** | **7 sprints** | **7 marcos** | **176 tasks** | **176** | **18 features** | **4 épicos** | **3 sprints concluídos + Sprint 4/5/6 em andamento** | **129/176 (73%)** | **23/07/2026** |

### Legenda

| Ícone | Significado |
|:---:|:---|
| ✅ | Concluída — Sprint entregue e aprovada pelo PO |
| 🔄 | Em andamento — Sprint em execução |
| ⬜ | Não iniciada — Sprint pendente |

---

## Progresso

| Sprint | Status | Início Real | Fim Real | Review PO |
|:---|:---:|:---|:---|:---:|
| Sprint 1 — Setup | ✅ Concluída | 14/07/2026 | 14/07/2026 | ✅ Aprovada |
| Sprint 2 — Segurança | ✅ Concluída | 14/07/2026 | 14/07/2026 | ✅ Aprovada |
| Sprint 3 — Portal Admin | ✅ Concluída | 16/07/2026 | 17/07/2026 | ✅ Aprovada |
| Sprint 4 — RBAC | 🔄 Em andamento | 17/07/2026 | — | Frentes 0-4 concluídas (31/48 tasks, 65%) |
| Sprint 5 — Portal Cliente | 🔄 Em andamento | 23/07/2026 | — | Frentes 0-1-2-3a concluídas (36/40 tasks, 90%) |
| Sprint 6 — BUs e Catálogo | 🔄 Em andamento | 23/07/2026 | — | Frentes 0+1 concluídas (9/18 tasks, 50%) |
| Sprint 7 — Homologação | ⬜ Não iniciada | — | — | — |

---

🤖 *Artefatos de sprint gerados a partir de TASKS.md v3.9, TEST_PLAN.md v3.4, SPECS.md v2.8 e PRD.md v1.19. Sprint 3 concluída em 17/07/2026: 42/42 (100%). Sprint 4/5 em andamento: Frentes 0-4 (31/48) + Frentes 0-1-2-3a (36/40). Sprint 6 Frentes 0+1 concluídas (9/18 = 50%). 3 sprints concluídos + 3 em andamento. 129/176 (73%). 288 testes (0 failures).*
