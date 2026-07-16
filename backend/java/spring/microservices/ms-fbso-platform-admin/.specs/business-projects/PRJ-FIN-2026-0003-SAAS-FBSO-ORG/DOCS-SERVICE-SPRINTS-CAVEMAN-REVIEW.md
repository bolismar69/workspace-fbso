# DOCS-SERVICE-SPRINTS-CAVEMAN-REVIEW.md — Revisão dos Artefatos de Sprint (1 e 2)

- **Serviço:** `ms-fbso-platform-admin`
- **Sprints revisadas:** Sprint 1 — Setup, Sprint 2 — Segurança
- **Docs-mestre de referência:** PRD v1.4, SPECS v1.5, TASKS v2.3, TEST_PLAN v2.3, ARCHITECTURE v1.3
- **Data:** 15/07/2026
- **Fontes:** Leitura completa de 11 artefatos de sprint + comparação com 5 docs-mestre atualizados

---

## Resumo

| Sprint | 🔴 bug | 🟡 risk | 🔵 nit | Total |
|--------|:---:|:---:|:---:|:---:|
| Sprint 1 — Setup | 0 | 0 | 3 | 3 |
| Sprint 2 — Segurança | 8 | 3 | 3 | 14 |
| Cross-sprint | 0 | 1 | 2 | 3 |
| **Total** | **8** | **4** | **8** | **20** |

---

## Sprint 1 — Setup

### SPRINT-CARD.md

🔵 `L89 rodapé`: "Gerado a partir de TASKS.md v2.0". TASKS.md é v2.3. Atualizar referência.

### SPRINT-REVIEW.md

🔵 `L52`: "Confirmar 14 pacotes conforme ARCHITECTURE.md §2". T-008 implementou 47 diretórios (14 pacotes top-level + sub-pacotes). Esclarecer contagem.

### SPRINT-TEST-SUITE.md

🔵 `L87 rodapé`: "Extraído de TEST_PLAN.md v2.1 §9". TEST_PLAN.md é v2.3. Atualizar referência.

---

## Sprint 2 — Segurança

### SPRINT-CARD.md

🔴 `L49 Entregável`: "Migration V003 — PostgreSQL RLS em **11 tabelas**". ARCHITECTURE v1.3 corrigiu para 5 tabelas (subscription, user, business_unit, product_service, audit_log). Corrigir para 5.

🔴 `L61 DoD checklist`: "PostgreSQL RLS habilitado em **11 tabelas** (V003)". Corrigir para 5.

🟡 `L94 Métricas`: "Tasks completadas \| 7/7". Sprint tem 8 tarefas (T-009 a T-015 + T-015.1). Corrigir para 8/8.

🔵 `L100 rodapé`: "Gerado a partir de TASKS.md v2.0". TASKS.md é v2.3. Atualizar referência.

### SPRINT-DEVELOPMENT-PLANNING.md

🔴 `L98 Critério DONE`: "RLS ativo em **11 tabelas**". Corrigir para 5.

🔴 `L101 Comentário SQL`: "-- ENABLE ROW LEVEL SECURITY para **11 tabelas**". Corrigir para 5.

🔴 `L115 Descrição`: "Migration V003: RLS + políticas em **11 tabelas**". Corrigir para 5.

🟡 `L107-108 Código exemplo`: SQL injection — concatenação `"SET app.current_tenant_id = '" + tenantId + "'"`. ARCHITECTURE v1.3 L337 corrigiu para `jdbcTemplate.update("SET LOCAL app.current_tenant_id = ?", tenantId)`. Atualizar exemplo para PreparedStatement.

### SPRINT-TEST-PLANNING.md

🔴 `L13`: "Tasks implementadas: **7** (T-009 a T-015)". Com T-015.1 são 8 tasks. Corrigir.

🟡 `L22-42`: Tabela de mapeamento cobre T-009 a T-015 (TC-S2-001 a TC-S2-021). Sem linha para T-015.1 (RLS). Adicionar mapeamento para TC-S2-022 a TC-S2-026.

### SPRINT-TEST-SUITE.md

🔴 `L40 TC-S2-022`: "RLS habilitado em **11 tabelas**". Corrigir para 5.

🔴 `L41 TC-S2-023`: "**11 políticas** tenant_isolation criadas". Corrigir para 5 (uma policy por tabela com RLS).

🔵 `L104 rodapé`: "Extraído de TEST_PLAN.md v2.0". TEST_PLAN.md é v2.3. Atualizar referência.

### SPRINT-2-EXECUTION-REPORT.md

🟡 `L8, L14`: "Tasks executadas: **7/7**". Sprint tem 8 tarefas totais. Explicação: T-015.1 tem relatório separado (SPRINT-2-EXECUTION-REPORT-T015.1.md), e T-012 foi substituído. Mas para consistência com SPRINT-CARD (que lista 8), ajustar para "7 implementadas + 1 substituída (T-012) = 8/8 planejadas".

### SPRINT-2-EXECUTION-REPORT-T015.1.md

✅ Sem achados. Documenta corretamente: 5 tabelas, 33/33 testes, DataSource proxy, SQL injection notado. Consistente com ARCHITECTURE v1.3 e SPECS v1.5.

### SPRINT-REVIEW.md

🔵 `L20`: Seção "TenantContext + PostgreSQL RLS (T-011, T-012, T-015.1)". T-012 foi substituído por T-015.1. Considerar renomear para "TenantContext + PostgreSQL RLS (T-011, T-015.1)".

---

## Cross-Sprint

🟡 **SQL injection em doc de desenvolvimento**: SPRINT-DEVELOPMENT-PLANNING.md L107-108 mostra padrão vulnerável de concatenação SQL. Documento de planejamento influencia código futuro — deve exibir padrão correto (PreparedStatement).

🔵 **Versões stale nos rodapés**: Todos os SPRINT-CARDs e SPRINT-TEST-SUITEs referenciam versões antigas dos docs-mestre (TASKS v2.0, TEST_PLAN v2.0/v2.1). Atualizar para v2.3.

🔵 **SPRINT-CARD Sprint 2**: Lista docs-mestre como TASKS.md + ARCHITECTURE.md, omitindo SPECS.md e TEST_PLAN.md. Sprint 1 lista SPECS + ARCHITECTURE. Padronizar.

---

## Plano de Correção

### Prioridade 1 — Bugs (RLS "11→5 tabelas")
| # | Arquivo | Ação |
|---|---------|------|
| 1 | SPRINT-CARD.md L49 | "11 tabelas" → "5 tabelas" |
| 2 | SPRINT-CARD.md L61 | "11 tabelas" → "5 tabelas" |
| 3 | SPRINT-DEVELOPMENT-PLANNING.md L98 | "11 tabelas" → "5 tabelas" |
| 4 | SPRINT-DEVELOPMENT-PLANNING.md L101 | "11 tabelas" → "5 tabelas" |
| 5 | SPRINT-DEVELOPMENT-PLANNING.md L115 | "11 tabelas" → "5 tabelas" |
| 6 | SPRINT-TEST-SUITE.md L40 | "11 tabelas" → "5 tabelas" |
| 7 | SPRINT-TEST-SUITE.md L41 | "11 políticas" → "5 políticas" |
| 8 | SPRINT-TEST-PLANNING.md L13 | "7 tasks" → "8 tasks" |

### Prioridade 2 — Riscos
| # | Arquivo | Ação |
|---|---------|------|
| 9 | SPRINT-CARD.md L94 | "7/7" → "8/8" |
| 10 | SPRINT-2-EXECUTION-REPORT.md L8,L14 | "7/7" → "7 implementadas + 1 substituída (T-012)" |
| 11 | SPRINT-DEVELOPMENT-PLANNING.md L107-108 | SQL injection → PreparedStatement |
| 12 | SPRINT-TEST-PLANNING.md L22-42 | Adicionar linha T-015.1 no mapeamento |

### Prioridade 3 — Nits
| # | Arquivo | Ação |
|---|---------|------|
| 13 | SPRINT-CARD Sprint 1 L89 | TASKS v2.0 → v2.3 |
| 14 | SPRINT-CARD Sprint 2 L100 | TASKS v2.0 → v2.3 |
| 15 | SPRINT-TEST-SUITE Sprint 1 L87 | TEST_PLAN v2.1 → v2.3 |
| 16 | SPRINT-TEST-SUITE Sprint 2 L104 | TEST_PLAN v2.0 → v2.3 |
| 17 | SPRINT-REVIEW Sprint 1 L52 | Esclarecer contagem pacotes |
| 18 | SPRINT-REVIEW Sprint 2 L20 | Remover T-012 da seção RLS |
| 19 | SPRINT-CARD Sprint 2 L8 | Adicionar SPECS.md + TEST_PLAN.md nos docs-mestre |

---

🤖 *Revisão gerada em 15/07/2026. 20 achados em 11 artefatos de sprint. Fonte: comparação com PRD v1.4, SPECS v1.5, TASKS v2.3, TEST_PLAN v2.3, ARCHITECTURE v1.3.*
