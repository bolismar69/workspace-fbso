# Sprints — ms-fbso-platform-admin

- **Projeto:** [PRJ-FIN-2026-0003-SAAS-FBSO-ORG](../)
- **Total de Sprints:** 7
- **Duração:** 14 semanas (24/07/2026 → 30/10/2026)
- **Documentos-mestre:** [TASKS.md](../TASKS.md) v2.4 | [SPECS.md](../SPECS.md) v1.6 | [TEST_PLAN.md](../TEST_PLAN.md) v2.4 | [PRD.md](../PRD.md) v1.5 | [ARCHITECTURE.md](../ARCHITECTURE.md) v1.4
- **🚫 Branch Obrigatória:** `feature/java-fbso-platform-admin` — [ver detalhes no PRD §8.4](../PRD.md#84-branch-de-desenvolvimento)

---

## Índice de Sprints

| Sprint | Datas | Marco | Tarefas | Artefatos |
|:---|:---|:---|:---:|:---|
| [**Sprint 1**](sprint-01-setup/) | 24/07 — 07/08 | Pre-M2 Setup | T-001 a T-008 (8) | [Card](sprint-01-setup/SPRINT-CARD.md) · [Testes](sprint-01-setup/SPRINT-TEST-SUITE.md) · [Review](sprint-01-setup/SPRINT-REVIEW.md) |
| [**Sprint 2**](sprint-02-seguranca/) | 07/08 — 15/08 | Pre-M2 Segurança | T-009 a T-015, T-015.1 (8) | [Card](sprint-02-seguranca/SPRINT-CARD.md) · [Testes](sprint-02-seguranca/SPRINT-TEST-SUITE.md) · [Review](sprint-02-seguranca/SPRINT-REVIEW.md) |
| [**Sprint 3**](sprint-03-portal-admin/) | 15/08 — 31/08 | M2 + M3 | T-016 a T-038 (23) | [Card](sprint-03-portal-admin/SPRINT-CARD.md) · [Testes](sprint-03-portal-admin/SPRINT-TEST-SUITE.md) · [Review](sprint-03-portal-admin/SPRINT-REVIEW.md) |
| [**Sprint 4**](sprint-04-rbac/) | 31/08 — 15/09 | M4 | T-039 a T-049 (11) | [Card](sprint-04-rbac/SPRINT-CARD.md) · [Testes](sprint-04-rbac/SPRINT-TEST-SUITE.md) · [Review](sprint-04-rbac/SPRINT-REVIEW.md) |
| [**Sprint 5**](sprint-05-portal-cliente/) | 15/09 — 30/09 | M5 | T-050 a T-061 (12) | [Card](sprint-05-portal-cliente/SPRINT-CARD.md) · [Testes](sprint-05-portal-cliente/SPRINT-TEST-SUITE.md) · [Review](sprint-05-portal-cliente/SPRINT-REVIEW.md) |
| [**Sprint 6**](sprint-06-bus-catalogo/) | 30/09 — 15/10 | M6 | T-062 a T-070 (9) | [Card](sprint-06-bus-catalogo/SPRINT-CARD.md) · [Testes](sprint-06-bus-catalogo/SPRINT-TEST-SUITE.md) · [Review](sprint-06-bus-catalogo/SPRINT-REVIEW.md) |
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
        │ Sprint 1│Sprint│ Sprint 3│ Sprint 4│ Sprint 5│ Sprint 6│ Sprint 7│
        │ Setup   │  2   │ M2 + M3 │   M4    │   M5    │   M6    │   M7    │
        │         │ Seg. │         │         │         │         │         │
        └─────────┴──────┴─────────┴─────────┴─────────┴─────────┴─────────┘
      24/07     07/08  15/08     31/08     15/09     30/09     15/10     30/10
```

---

## Matriz de Rastreabilidade

> **Atualizado:** 16/07/2026 · **Fonte:** [TASKS.md](../TASKS.md) v2.4 + [PRD.md](../PRD.md) v1.5
>
> Esta matriz é a **camada de integração** entre o planejamento ([TASKS.md](../TASKS.md)) e a execução (sprints/). Use-a como mapa único para navegar entre fases, sprints, tarefas, features e épicos.

| Fase | Sprint | Marco | Tasks | Qtd | Features | Épico | Status | Progresso | Atualização |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|:---:|
| **FASE 0** | Sprint 1 — Setup | Pre-M2 Setup | T-001 a T-008 | 8 | — | — | ✅ Concluída | 8/8 (100%) | 14/07/2026 |
| **FASE 0** | Sprint 2 — Segurança | Pre-M2 Segurança | T-009 a T-015, T-015.1 | 8 | — | — | ✅ Concluída | 8/8 (100%) | 14/07/2026 |
| **FASE 1** | Sprint 3 — Portal Admin | M2 (EP-01) | T-016 a T-023 | 8 | F01-01, F01-02, F01-03 | EP-01 | 🔄 Em andamento | 0/8 (0%) | 16/07/2026 |
| **FASE 2** | Sprint 3 — Portal Admin | M3 (EP-02) | T-024 a T-038 | 15 | F02-01 a F02-05 | EP-02 | 🔄 Em andamento | 0/15 (0%) | 16/07/2026 |
| **FASE 3** | Sprint 4 — RBAC | M4 (EP-03) | T-039 a T-049 | 11 | F03-01 a F03-04 | EP-03 | ⬜ Não iniciada | 0/11 (0%) | — |
| **FASE 4** | Sprint 5 — Portal Cliente | M5 (EP-04a) | T-050 a T-061 | 12 | F04-01 a F04-04 | EP-04 | ⬜ Não iniciada | 0/12 (0%) | — |
| **FASE 5** | Sprint 6 — BUs e Catálogo | M6 (EP-04b) | T-062 a T-070 | 9 | F04-05, F04-06 | EP-04 | ⬜ Não iniciada | 0/9 (0%) | — |
| **FASE 6** | Sprint 7 — Homologação | M7 (Homologação) | T-071 a T-079 | 9 | Todas (18) | Todos | ⬜ Não iniciada | 0/9 (0%) | — |
| **TOTAL** | **7 sprints** | **7 marcos** | **80 tasks** | **80** | **18 features** | **4 épicos** | **2 concluídos, 1 em andamento** | **16/80 (20%)** | **16/07/2026** |

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
| Sprint 3 — Portal Admin | 🔄 Em andamento | 16/07/2026 | — | — |
| Sprint 4 — RBAC | ⬜ Não iniciada | — | — | — |
| Sprint 5 — Portal Cliente | ⬜ Não iniciada | — | — | — |
| Sprint 6 — BUs e Catálogo | ⬜ Não iniciada | — | — | — |
| Sprint 7 — Homologação | ⬜ Não iniciada | — | — | — |

---

🤖 *Artefatos de sprint gerados a partir de TASKS.md v2.3, TEST_PLAN.md v2.3, SPECS.md v1.5 e PRD.md v1.4. Matriz de Rastreabilidade atualizada em 16/07/2026. Sprint 3 iniciada em 16/07/2026.*
