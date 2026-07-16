# DOCS-SERVICE-CAVEMAN-REVIEW.md — Revisão dos Documentos do Serviço

- **Serviço:** `ms-fbso-platform-admin`
- **Docs revisados:** PRD.md v1.3, SPECS.md v1.4, TASKS.md v2.2, TEST_PLAN.md v2.2, ARCHITECTURE.md v1.2
- **Data:** 15/07/2026
- **Fontes:** Leitura completa 5 docs + DOCS-OTHERS-CAVEMAN-REVIEW.md + DOCS-REQS-REVIEW.md + DOCS-USER-STORY-REVIEW.md + auditoria cruzada multi-agente

---

## Resumo

| Doc | 🔴 bug | 🟡 risk | 🔵 nit | Total |
|:---|:---:|:---:|:---:|:---:|
| PRD.md | 4 | 2 | 3 | 9 |
| SPECS.md | 3 | 2 | 3 | 8 |
| TASKS.md | 3 | 3 | 2 | 8 |
| TEST_PLAN.md | 3 | 1 | 1 | 5 |
| ARCHITECTURE.md | 3 | 3 | 3 | 9 |
| Cross-doc | 4 | 1 | 1 | 6 |
| **Total** | **20** | **12** | **13** | **45** |

---

## PRD.md (v1.3)

🔴 `§5.2`: Claim "8 ADRs" mas lista 7 (ADR-01,02,04,05,06,07,08). ADR-03 ausente sem explicação. Corrigir count para 7 ou documentar ADR-03.

🔴 `§4.1`: Lista 10 entidades. Omite AuditEntry. SPECS §6.1 e ARCHITECTURE §2 listam 11. Adicionar AuditEntry à tabela.

🔴 `§3.3`: SPECS.md versão "v1.3" (real=v1.4), TASKS.md versão "v2.1" (real=v2.2). Corrigir.

🔴 `§6.1`: "Migration V003 em todas as 11 tabelas com tenant_id". T-015.1 implementou 5 tabelas. Corrigir para 5 com lista explícita.

🟡 `§6.7 BR-NFR02`: Descrição "TenantIsolationFilter + RbacInterceptor". Pipeline real: JWT Filter → TenantContext → RBAC → PostgreSQL RLS → BaseRepository. Atualizar.

🟡 `§4.6`: BR-B05 (Catálogo) "Should Have". DOCS-REQS-REVIEW recomenda Must Have. Verificar com PO.

🔵 `§4.5`: Tabela Cobertura de Features — coluna "Entrega" usa D1-D7. §4.3 usa M2-M7. Sem mapeamento Dx→Mx visível. Adicionar referência cruzada.

🔵 `§5.2`: ADR-01 descrito como "TenantIsolationFilter". Implementação real: PostgreSQL RLS + BaseRepository. Atualizar nome.

🔵 `§9.1`: Ordem leitura referencia ARCHITECTURE.md do business-inputs. Adicionar também ARCHITECTURE.md local do serviço.

---

## SPECS.md (v1.4)

🔴 `§1.2`: "10 entidades" mas §6.1 lista 11 (inclui AuditEntry). Corrigir para 11.

🔴 `§4.1`: Tabela lista 36 endpoints. Changelog v1.0 e TEST_PLAN §1.3 afirmam 37. Faltando `GET /api/v1/dashboard/client/summary`. Adicionar.

🔴 `§3.3`: "45/45 RNs mapeadas". Contagem real: 51 RNs individuais na matriz. Corrigir para 51.

🟡 `§5 BR-NFR02`: "PostgreSQL RLS (Camada 1) + BaseRepository (Camada 2) + Teste (Camada 3)". §1.2 pipeline lista 6 estágios (JWT→TenantContext→RBAC→RLS→BaseRepository→Auditoria). Clarificar: pipeline ≠ camadas de defesa.

🟡 `§11 Changelog v1.4`: "Corrigido glossário (§10): Migration V003 cobre 5 tabelas". Entrada autorreferencia a própria versão. Ajustar redação.

🔵 `§2.2`: Tabela cobre 8 de 18 features. Completar ou adicionar nota remetendo a §7.

🔵 `§10 Glossário - JWT`: Adicionar claims padrão (tenant_id, user_id, roles, business_unit_ids, modules).

🔵 `§10 Glossário - RBAC`: Referencia `RbacAspect`. ARCHITECTURE ADR-L07 removeu TenantIsolationAspect mas RbacAspect permanece. Verificar se nome está correto.

---

## TASKS.md (v2.2)

🔴 `Header`: "Origem: SPECS.md v1.3". SPECS é v1.4. Atualizar.

🔴 `§1`: "Should Have: 4 (F01-03, F04-03)". Tasks Should reais: T-021 + T-056 + T-057 = 3. §4 soma 4 Should (M2:2 + M5:2). M2 tem 1 Should (T-021), não 2. Corrigir para 3 Should total ou identificar 4ª tarefa.

🔴 `§4 M2`: "6 Must / 2 Should". Tasks M2: 7 Must + 1 Should (T-021). Corrigir para 7/1.

🟡 `§2 M3 - T-027`: "11 endpoints REST funcionais". SPECS §4.1 lista 7 endpoints tenant. Corrigir contagem.

🟡 `§2 M2 - T-022/T-023`: Marcadas Must mas cobrem F01-03 (Should). Testes de feature Should como Must — inconsistência de prioridade.

🟡 `§3`: Diagrama dependências mostra T-012 (TenantIsolationAspect) como nó ativo. T-012 substituído por T-015.1. Adicionar marcador [SUBSTITUÍDO] no diagrama.

🔵 `§2 M3`: Tasks referenciam papéis não-RBAC ("Gestor de Produto", "Líder Comercial", "Auditor Interno") em critérios DONE. Trocar por papéis RBAC reais.

🔵 `§1`: "Should Have: 4 (F01-03, F04-03)" lista features, não tasks. Clareza: "3 tarefas Should (T-021, T-056, T-057)".

---

## TEST_PLAN.md (v2.2)

🔴 `Header`: "Origem: SPECS.md v1.1". SPECS é v1.4 (3 versões atrás). Atualizar.

🔴 `§2`: "Total Geral de Cenarios: 154". Soma real: features(124) + cross-cutting(25) = 149. §9 adiciona 25 cenários infra não contabilizados. Total real: 149 ou 174 se incluir §9. Corrigir header.

🔴 `§1.3`: "100% das 16 RNs implementadas". PRD §6.6 tem 18 famílias, SPECS §3.3 tem ~51 RNs. Número "16" é artefato de versão anterior. Corrigir para contagem real.

🟡 `§2 Feature coverage`: Totais coluna não batem com contagem individual por feature na §3 (~127 vs 124). Auditar e reconciliar.

🔵 `§6.1`: Checklist regressão — colunas M2-M7 vazias. Esperado para sprints não iniciadas. OK.

---

## ARCHITECTURE.md (v1.2)

🔴 `L337`: SQL injection. `jdbcTemplate.execute("SET app.current_tenant_id = '" + tenantId + "'")`. Concatenando tenant_id no SQL. Usar `jdbcTemplate.update("SET app.current_tenant_id = ?", tenantId)`.

🔴 `L310-331`: Migration V003 SQL lista ALTER TABLE para 11 tabelas (tenant, user, business_unit, product_service, subscription, user_permission, audit_log, plan, plan_module, resource_action, role_resource). Incorreto: plan/plan_module/resource_action/role_resource não têm tenant_id. tenant com RLS quebraria Admin FBSO. Implementação real: 5 tabelas. Corrigir SQL exemplo.

🔴 `§2 security/aspect/`: Lista `TenantIsolationAspect.java`. ADR-L07 declara removido. TASKS T-012 marcado [SUBSTITUÍDO]. Remover do file listing ou adicionar nota.

🟡 `L164`: JWT payload exemplo com `modules: ["TRIBUTALI_ENGINE", "STOREKEEPER_PORTAL"]` — módulos Fase Futura. Adicionar `"FBSO_PLATFORM"` como módulo base Fase 0.

🟡 `§8`: ADRs numerados ADR-L01 a L07 (7). PRD §5.2 lista ADR-01 a 08 (7, com gap ADR-03). Unificar nomenclatura e contagem.

🟡 `§9 Changelog`: Ordem cronológica incorreta: 1.0 → 1.2 → 1.1. Reordenar: 1.0 → 1.1 → 1.2.

🔵 `§2`: Package `com.fbso.platform.admin/` vs serviço `ms-fbso-platform-admin`. Sem impacto — documentar convenção.

🔵 `§1.1`: Diagrama ASCII lista TenantIsolationAspect na camada security. Removido pelo ADR-L07. Atualizar diagrama.

🔵 `§2`: Package `utils/` (CnpjValidator, JwtUtils, DateUtils) não listado no PRD §5.3. Adicionar ou documentar como detalhe de arquitetura.

---

## Cross-Documento

🔴 **Entidade AuditEntry**: PRD §4.1 omite (10 entidades). SPECS §6.1 e ARCHITECTURE §2 incluem (11 entidades). Padronizar 11 em todos.

🔴 **Endpoint count 36 vs 37**: SPECS §4.1 lista 36. TASKS T-073 e TEST_PLAN §1.3 dizem 37. Faltando `GET /dashboard/client/summary` no SPECS. Adicionar.

🔴 **RLS 11 vs 5 tabelas**: PRD §6.1 e ARCHITECTURE §4.3: 11 tabelas. TASKS T-015.1, SPECS §10, TEST_PLAN §9.7: 5 tabelas. Decisão: meta=11, implementado=5. Documentos de design (PRD, ARCHITECTURE) devem declarar meta com nota de estado atual.

🔴 **Versões stale em referências cruzadas**: TASKS→SPECS v1.3 (real v1.4), TEST_PLAN→SPECS v1.1 (real v1.4), PRD→SPECS v1.3/TASKS v2.1 (real v1.4/v2.2). Atualizar todas.

🟡 **Contagem Must/Should**: TASKS total 76/4 OK no agregado. M2 diverge (6/2 tabela vs 7/1 tasks). M5: 2 Should na tabela, 2 Should nas tasks. Corrigir M2.

🔵 **App Switcher RNs**: RN12-02, RN16-02, RN16-03 marcadas "Frontend" no SPECS §3.3 mas descritas com impacto técnico no PRD §6.6. Alinhar responsabilidade.

---

## Impacto das Revisões Business-Inputs

Revisões dos docs de negócio geraram achados que precisam ser propagados aos docs de serviço:

### Alterações de escopo
1. **RBAC: 3 papéis MVP** (remover Auditor da Fase 0). Impacta PRD §4.6, SPECS §2.2, TASKS T-042/043/044, TEST_PLAN §3.10.
2. **BR-B05 (Catálogo): subir para Must Have**. Impacta PRD §4.6, TASKS T-065/066/067.
3. **F04-03 (Dashboard Cliente): manter Should Have** com fallback documentado pós-onboarding.

### Novos requisitos
4. **Backend enforcement US**: verificação de permissões em API (HTTP 403 + JSON). Impacta TASKS (nova task), TEST_PLAN (cenário segurança).
5. **Proteção admin lockout**: RN para impedir rebaixar/desativar último Admin. Impacta SPECS §3.1, TASKS T-044.
6. **~40 edge cases** (empty/error/loading states) ausentes nas US. Impacta TASKS critérios DONE, TEST_PLAN cenários.

### Correções de atores
7. **Papéis não-RBAC** ("Gestor de Produto", "Líder Comercial", "Auditor Interno"): trocar por papéis RBAC reais. Impacta TASKS tarefas M3.

### Terminologia
8. **"App Switcher" → "Seletor de Módulo"**. Impacta PRD §4.4/§4.5, SPECS §7, TASKS T-058/059.
9. **"Microserviços" → "Backend REST"** ou "Aplicação Modular". Impacta cabeçalhos ARCHITECTURE.md.

---

## Plano de Correção (Ordem Sugerida)

### Fase 1 — Correções críticas (bugs de segurança e dados)
| # | Doc | Achado | Ação |
|---|-----|--------|------|
| 1 | ARCHITECTURE | SQL injection L337 | Corrigir para PreparedStatement |
| 2 | ARCHITECTURE | RLS 11 tabelas L310-331 | Corrigir SQL exemplo para 5 tabelas |
| 3 | PRD | RLS 11 tabelas §6.1 | Corrigir para 5 com nota de meta futura |
| 4 | Cross-doc | RLS 11 vs 5 inconsistente | Alinhar PRD+ARCHITECTURE com realidade (5) |

### Fase 2 — Consistência de contagens
| # | Doc | Achado | Ação |
|---|-----|--------|------|
| 5 | SPECS | Faltando endpoint dashboard/client | Adicionar §4.1 |
| 6 | SPECS | RN count 45 vs 51 | Corrigir contagem |
| 7 | PRD | Entity count 10 vs 11 | Adicionar AuditEntry |
| 8 | TASKS | Must/Should M2 (6/2 vs 7/1) | Corrigir para 7/1 |
| 9 | TASKS | Should total 4 vs 3 | Corrigir para 3 |
| 10 | TEST_PLAN | Cenários 154 vs 149/174 | Recalcular e corrigir |
| 11 | TEST_PLAN | RNs 16 vs ~51 | Corrigir contagem |

### Fase 3 — Versões e referências
| # | Doc | Achado | Ação |
|---|-----|--------|------|
| 12 | PRD §3.3 | SPECS v1.3→v1.4, TASKS v2.1→v2.2 | Atualizar |
| 13 | TASKS header | SPECS v1.3→v1.4 | Atualizar |
| 14 | TEST_PLAN header | SPECS v1.1→v1.4 | Atualizar |
| 15 | ARCHITECTURE §9 | Ordem changelog 1.0,1.2,1.1 | Reordenar |
| 16 | PRD §5.2 | "8 ADRs" vs 7 listados | Corrigir count |

### Fase 4 — Design e terminologia
| # | Doc | Achado | Ação |
|---|-----|--------|------|
| 17 | ARCHITECTURE §2 | TenantIsolationAspect no file listing | Remover ou anotar [REMOVIDO] |
| 18 | ARCHITECTURE §1.1 | Diagrama com TenantIsolationAspect | Atualizar |
| 19 | ARCHITECTURE L164 | JWT modules Fase Futura | Adicionar FBSO_PLATFORM |
| 20 | ARCHITECTURE §8 | ADR-Lxx vs ADR-xx | Unificar nomenclatura |

---

🤖 *Revisão gerada em 15/07/2026. 45 achados: 20 bugs, 12 riscos, 13 nits. Fontes: leitura completa 5 docs + 3 revisões business-inputs + 2 agentes de auditoria paralelos.*
