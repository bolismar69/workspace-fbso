# DOCS-SPRINT-CAVEMAN-REVIEW.md — Revisão da Documentação: Sprint 3

- **Sprint:** 3 de 7 — Portal Admin + Contas e Planos
- **Marco:** M2 (EP-01) + M3 (EP-02)
- **Docs-mestre de referência:** PRD v1.4, SPECS v1.5, TASKS v2.3, TEST_PLAN v2.3, ARCHITECTURE v1.3
- **Data:** 15/07/2026
- **Status da Sprint:** Não iniciada

---


## Resumo

| Arquivo | 🔴 bug | 🟡 risk | 🔵 nit | Total |
|---------|:---:|:---:|:---:|:---:|
| SPRINT-CARD.md | 2 | 3 | 2 | 7 |
| SPRINT-TEST-SUITE.md | 0 | 1 | 1 | 2 |
| SPRINT-REVIEW.md | 1 | 0 | 0 | 1 |
| **Total** | **3** | **4** | **3** | **10** |

---

## SPRINT-CARD.md

🔴 `L42 T-027`: "11 endpoints REST". TASKS v2.3 corrigiu para 7 endpoints (GET list, GET by id, POST, PATCH, POST /suspend, POST /reactivate, POST /resend-invite). Corrigir.

🔴 `L110 Métricas`: "RNs implementadas: 15". Contagem real: 20 RNs (F01-01:3 + F01-02:2 + F01-03:2 + F02-01:3 + F02-02:3 + F02-03:3 + F02-04:2 + F02-05:2 = 20). Com RN07-03 incluso = 21. Corrigir para 20.

🟡 `L91 Risco`: "T-021 e T-028 podem ser negociados como Should". T-021 já é Should (F01-03) conforme TASKS v2.3 (Must/Should M2 = 7/1). Apenas T-028 é Must negociável. Corrigir para "T-028 (email) pode ser negociado como Should; T-021 já é Should".

🟡 `L69 F02-04`: Lista RN07-01 e RN07-02. Omite RN07-03 (data de término opcional — `end_date` nullable). TASKS T-032 e SPECS §4.2 e §6.1 cobrem RN07-03. Adicionar.

🟡 `L48 T-033`: Critério DONE menciona "change-plan sem gap" mas não referencia RN07-02 explicitamente. Adicionar referência para rastreabilidade.

🔵 `L8`: Docs-mestre lista TASKS, SPECS, TEST_PLAN. Omite PRD.md e ARCHITECTURE.md (presentes nos cards das Sprints 1-2). Padronizar.

🔵 `L116 rodapé`: "Gerado a partir de TASKS.md v2.0". TASKS é v2.3. Atualizar.

---

## SPRINT-TEST-SUITE.md

🟡 `§Resumo L110-116`: Tabela de sumário mostra 48 cenários (Unit:18, Int:18, E2E:6, Seg:6). Soma real por feature: 7+4+5+7+9+7+9+8 = 56. Soma por nível real: Unit:24, Int:19, E2E:6, Seg:6 = 55. Divergência de ~8 cenários entre sumário e contagem individual. Recalcular sumário.

🔵 `L135 rodapé`: "Extraído de TEST_PLAN.md v2.0". TEST_PLAN é v2.3. Atualizar.

---

## SPRINT-REVIEW.md

🔴 `L90 Métricas`: "RNs implementadas: 15". Mesmo problema do SPRINT-CARD — contagem real é 20. Corrigir.

---

## Verificações OK

| Verificação | Status |
|:---|:---:|
| Tasks M2 (T-016 a T-023 = 8) conferem com TASKS v2.3 | ✅ |
| Tasks M3 (T-024 a T-038 = 15) conferem com TASKS v2.3 | ✅ |
| Total 23 tasks = 8+15 | ✅ |
| T-021 marcado Should — consistente com TASKS v2.3 | ✅ |
| 48 cenários cobrem 8 features (F01-01 a F02-05) | ✅ |
| Features listadas conferem com PRD v1.4 §4.5 | ✅ |
| RNs por feature conferem com SPECS v1.5 §3.3 | ✅ |
| Endpoints listados conferem com SPECS v1.5 §4.1 | ✅ |
| Sprint Goal alinhado com TASKS v2.3 §2 M2+M3 | ✅ |
| Dependências (Sprint 2 → Sprint 3 → Sprint 4) corretas | ✅ |

---

## Impacto das Correções nos Docs-Mestre

As seguintes alterações nos docs-mestre (já aplicadas) impactam a Sprint 3:

| Alteração no Doc-Mestre | Impacto na Sprint 3 |
|:---|:---|
| TASKS v2.3: T-027 "11→7 endpoints" | SPRINT-CARD L42 ainda diz 11 |
| TASKS v2.3: Must/Should M2 = 7/1 (T-021 = Should) | SPRINT-CARD L91 ainda sugere negociar T-021 |
| SPECS v1.5: entidades 11 (com AuditEntry) | OK — SPRINT-CARD já referencia F02-05 (Auditoria) |
| SPECS v1.5: 37 endpoints (com dashboard/client) | Não impacta Sprint 3 (dashboard/client é M5) |
| TEST_PLAN v2.3: 18 famílias RNs (não 16) | OK — já usa contagem por feature |
| PRD v1.4: RLS 5 tabelas (não 11) | Não impacta Sprint 3 (RLS é Sprint 2) |

---

## Plano de Correção

### Prioridade 1 — Bugs
| # | Arquivo | Linha | Ação |
|---|---------|-------|------|
| 1 | SPRINT-CARD.md | L42 | "11 endpoints" → "7 endpoints (GET list, GET by id, POST, PATCH, POST /suspend, POST /reactivate, POST /resend-invite)" |
| 2 | SPRINT-CARD.md | L110 | "RNs implementadas: 15" → "RNs implementadas: 20" |
| 3 | SPRINT-REVIEW.md | L90 | "RNs implementadas: 15" → "RNs implementadas: 20" |

### Prioridade 2 — Riscos
| # | Arquivo | Linha | Ação |
|---|---------|-------|------|
| 4 | SPRINT-CARD.md | L91 | Corrigir redação do risco (T-021 já é Should) |
| 5 | SPRINT-CARD.md | L69 | Adicionar RN07-03 em F02-04 |
| 6 | SPRINT-CARD.md | L48 | Adicionar "RN07-02" no critério DONE de T-033 |
| 7 | SPRINT-TEST-SUITE.md | §Resumo | Recalcular sumário (48→55/56) |

### Prioridade 3 — Nits
| # | Arquivo | Linha | Ação |
|---|---------|-------|------|
| 8 | SPRINT-CARD.md | L8 | Adicionar PRD.md e ARCHITECTURE.md nos docs-mestre |
| 9 | SPRINT-CARD.md | L116 | TASKS v2.0 → v2.3 |
| 10 | SPRINT-TEST-SUITE.md | L135 | TEST_PLAN v2.0 → v2.3 |

---

🤖 *Revisão gerada em 15/07/2026. 10 achados em 3 artefatos. Fonte: comparação com PRD v1.4, SPECS v1.5, TASKS v2.3, TEST_PLAN v2.3, ARCHITECTURE v1.3.*
