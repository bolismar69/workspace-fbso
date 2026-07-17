# DOCS-SPRINT-CAVEMAN-REVIEW.md — Revisão da Documentação: Sprint 3 (v2.0)

- **Sprint:** 3 de 7 — Portal Admin + Contas e Planos
- **Marco:** M2 (EP-01) + M3 (EP-02)
- **Docs-mestre de referência:** PRD v1.7, SPECS v1.6, TASKS v2.4, TEST_PLAN v2.4, ARCHITECTURE v2.0
- **Data:** 16/07/2026
- **Status da Sprint:** Em andamento (iniciada 16/07/2026)
- **Branch:** `feature/sprint-03-portal-admin`
- **Revisão anterior:** 15/07/2026 (10 achados — 6 corrigidos, 3 parcialmente, 1 pendente)

---

## Resumo Executivo

Esta é a **segunda iteração** do Caveman Review da Sprint 3, refeita integralmente com triangulação em 3 eixos:
1. **Docs-mestre × Docs-sprint** — consistência entre documentação canônica e artefatos da sprint
2. **Docs-sprint × Código** — aderência entre o que a documentação descreve e o que o código realmente contém
3. **Docs-sprint × Docs-sprint** — consistência cruzada entre os 5 artefatos da sprint

| Arquivo | 🔴 bug | 🟡 risk | 🔵 nit | Total |
|---------|:---:|:---:|:---:|:---:|
| SPRINT-CARD.md | 2 | 2 | 1 | 5 |
| SPRINT-TEST-SUITE.md | 1 | 1 | 0 | 2 |
| SPRINT-REVIEW.md | 1 | 1 | 1 | 3 |
| SPRINT-DEVELOPMENT-PLANNING-DRAFT.md | 2 | 4 | 2 | 8 |
| DOCS-SPRINT-CAVEMAN-REVIEW.md | — | — | 1 | 1 |
| **Código (achados cross-cutting)** | 2 | 3 | 2 | 7 |
| **Total** | **8** | **11** | **7** | **26** |

---

## PARTE 1: Docs-Mestre × Docs-Sprint

### 🔴 SPRINT-CARD.md

🔴 `L119 Métricas`: "RNs implementadas: 20". Soma real da tabela de features (L70-78): 3+2+2+3+3+3+3+2 = **21**. O RN07-03 foi adicionado em F02-04 mas a métrica não foi recalculada. Corrigir para 21.

🔴 `L120 Métricas`: "Cenários de teste: 55". TEST_PLAN v2.4 define 56 cenários (24U + 20I + 6E + 6S). SPRINT-TEST-SUITE também lista 56 por ID. Corrigir para 56.

🟡 `L9 Docs-mestre`: Lista ARCHITECTURE.md como **v1.4**. O doc real está em **v2.0** (refatoração completa em 16/07 com C4→Mermaid, 887 linhas). SPRINT-DEVELOPMENT-PLANNING-DRAFT referencia v2.0 corretamente. Corrigir para v2.0.

🟡 `L56 T-033 DONE`: Critério DONE diz "Change-plan sem gap. Transação atômica" mas não referencia RN07-02 explicitamente. A coluna Task referencia RN07-02 e RN07-03, mas DONE não. Adicionar "RN07-02" para rastreabilidade completa (achado #6 da review anterior, **ainda não corrigido**).

🔵 `L9 Docs-mestre`: PRD.md listado como v1.5. Doc real está em **v1.7** (branching strategy, §8.4). Corrigir para v1.7.

---

### 🔴 SPRINT-REVIEW.md

🔴 `L92 Métricas`: "20 RNs implementadas". Mesmo problema do SPRINT-CARD — contagem real é 21. Corrigir.

🟡 `L88-93 Métricas`: Tabela de métricas lista "17 endpoints" como target. Docs-mestre definem ~20 endpoints para Sprint 3 (4 em M2 + ~16 em M3). Verificar contagem exata: SPECS §4.1 lista GET dashboard/summary, GET dashboard/evolution, GET dashboard/alerts, GET tenants, POST tenants, PATCH tenants/{id}, POST tenants/{id}/suspend, POST tenants/{id}/reactivate, POST tenants/{id}/resend-invite, GET plans, POST plans, PATCH plans/{id}, POST plans/{id}/deactivate, GET tenants/{tid}/subscriptions, POST tenants/{tid}/subscriptions, POST subscriptions/{id}/change-plan, POST subscriptions/{id}/suspend, GET audit = **18 endpoints** (não 17, não 20). Corrigir target para 18.

🔵 `L97-101 Bloqueadores`: Seção de bloqueadores está vazia (template). Preencher com "Nenhum bloqueador identificado" ou remover seção.

---

### 🔴 SPRINT-TEST-SUITE.md

🔴 `L112-118 Sumário`: Tabela mostra 55 total (24U + 19I + 6E + 6S). Contagem real por ID:
- Unit: 24 ✓
- Integration: **20** (não 19) — F02-04 tem 4 testes de integração, não 3
- E2E: 6 ✓
- Security: 6 ✓ (embora F02-05 tenha 2 security, F01-03 tenha 0 — soma real = 6, não 7)
- **Total real: 56**, não 55

A divergência de Integration (19→20) já existia na review anterior. O sumário foi atualizado de 48→55 mas o subtotal de integração permaneceu errado. Corrigir: Integration=20, Total=56.

🟡 `L112-118 Sumário - Security`: A coluna "Seg" soma 6 mas a contagem por feature é: 1+0+0+1+1+1+1+2 = 7. O SPRINT-TEST-SUITE lista F01-03 sem teste de segurança (5 cenários: 2U+2I+1E). Verificar se Security=6 ou 7. **Recomendação:** manter 6 (F01-03 é feature Should, não exige teste de segurança dedicado) e explicitar que F01-03 não tem cenário Security.

---

## PARTE 2: Docs-Sprint × Código

### 🔴 SPRINT-DEVELOPMENT-PLANNING-DRAFT.md

🔴 `L7 Caminhos relativos`: Usa `../TASKS.md`, `../SPECS.md` etc. para referenciar docs-mestre. Ambos os arquivos estão em `sprints/sprint-03-portal-admin/`. O caminho `../` aponta para `sprints/`, não para a raiz do projeto. **Caminhos quebrados.** Corrigir para `../../TASKS.md` (padrão usado no SPRINT-CARD).

🔴 `L12 Status`: Diz "Não iniciada — planejamento prévio". Sprint 3 foi iniciada em 16/07/2026 conforme SPRINT-CARD e docs-mestre. Atualizar para "Em andamento — iniciada 16/07/2026".

🟡 `L68 Métricas`: "20 RNs" — mesmo problema de contagem (deveria ser 21).

🟡 `L70 Métricas`: "55 cenários de teste" — mesmo problema (deveria ser 56).

🟡 `L99-131 Sequenciamento`: Grafo de dependências está correto mas assume que T-024 (entidade/enums Tenant) é pré-requisito para T-025 e T-026. Na realidade, os enums `TenantStatus` e `Segment` **já existem no código** (`src/main/java/com/fbso/platform/admin/enums/`). T-024 pode ser parcialmente reduzido — só precisa criar RowMapper e validações.

🟡 `L156-180 Seção 5 - Migration V004`: Define 6 índices SQL como pré-requisito. V002 já criou `idx_tenant_status`, `idx_tenant_created`, `idx_subscription_tenant_active`. Verificar sobreposição antes de criar V004:
- `idx_tenant_status` — **já existe** (V002)
- `idx_tenant_created` — **já existe** (V002)
- `idx_plan_status` — novo, necessário
- `idx_subscription_tenant_active` — **já existe** (V002)
- `idx_audit_log_timestamp` — **já existe** (V002)
- `idx_audit_log_tenant` — **já existe** (V002)
Dos 6 índices propostos, **5 já existem**. Apenas `idx_plan_status` é novo. Revisar seção 5.

🔵 `L135-149 Plano dia-a-dia`: Assume 12 dias úteis com 1 desenvolvedor. Com base no código existente (infraestrutura pronta, enums existentes, BaseRepository funcional), o esforço real pode ser menor. Dia 1 (Migration V004) pode ser eliminado (índices já existem). Recalibrar para ~10 dias.

🔵 `L185-230 Seção 6 - Templates de código`: Templates usam `@Repository` e `JdbcTemplate` mas não herdam de `BaseRepository`. O padrão real do projeto é `extends BaseRepository<T>`. Templates deveriam refletir o padrão existente para evitar retrabalho.

---

### 🔴 Código (achados cross-cutting)

🔴 `TenantAwareDataSource.java:L57`: `SET app.current_tenant_id = '" + tenantId + "'"`. Concatenação de string para construir comando SQL — vulnerabilidade de SQL injection no `tenant_id`. Usar `SET app.current_tenant_id = ?` com `PreparedStatement`. **Mesmo sendo um valor interno do JWT, defense-in-depth exige parameterized statement.**

🔴 `application.yml:L8 e L53`: Duas chaves `spring:` no mesmo arquivo YAML. O segundo bloco (`spring.security.oauth2...`) sobrescreve o primeiro. Embora o YAML merge funcione, é uma irregularidade que pode causar bugs sutis se a ordem de merge mudar. Consolidar em um único bloco `spring:`.

🟡 `Cobertura de testes`: 5 classes de teste (todas unitárias/mock) — 0 testes de integração com Testcontainers apesar da dependência estar no pom.xml. TEST_PLAN v2.4 define 20 cenários de integração para Sprint 3. Nenhum teste de integração real existe ainda. Risco: os testes de integração serão escritos do zero durante a sprint, sem padrão estabelecido.

🟡 `RowMapper ausente`: Nenhum `RowMapper` implementado para as 5 entidades da Sprint 3 (Tenant, Plan, PlanModule, Subscription, AuditEntry). `BaseRepository` exige `RowMapper<T>` no construtor. Cada entidade precisará de seu RowMapper antes que qualquer repository funcione.

🟡 `OpenAPI/Swagger ausente`: Nenhuma dependência springdoc-openapi no pom.xml, nenhuma configuração Swagger. SPRINT-CARD menciona 17 endpoints REST mas não há mecanismo para documentá-los. Adicionar `springdoc-openapi-starter-webmvc-ui` nas dependências.

🔵 `pom.xml`: Testcontainers configurado mas sem uso. Se não for usado na Sprint 3, considerar mover para `test` scope com `@Disabled` até Sprint 4.

🔵 `Enums existentes`: `TenantStatus`, `SubscriptionStatus`, `Recurrence`, `Segment` já existem em `src/main/java/com/fbso/platform/admin/enums/`. Bom — reduz escopo de T-024 (não precisa criar enums, só validar se cobrem todos os estados necessários).

---

## PARTE 3: Docs-Sprint × Docs-Sprint (consistência cruzada)

| Métrica | SPRINT-CARD | SPRINT-REVIEW | TEST-SUITE | PLANNING-DRAFT | **Valor real** |
|:---|:---:|:---:|:---:|:---:|:---:|
| RNs | 20 ❌ | 20 ❌ | — | 20 ❌ | **21** |
| Testes | 55 ❌ | — | 55 ❌ | 55 ❌ | **56** |
| Endpoints | 17 ⚠️ | 17 ⚠️ | — | 17 ⚠️ | **18** |
| ARCH version | v1.4 ❌ | — | — | v2.0 ✓ | **v2.0** |
| PRD version | v1.5 ❌ | — | — | v1.5 ❌ | **v1.7** |
| Status sprint | Em andamento ✓ | Em andamento ✓ | — | Não iniciada ❌ | **Em andamento** |
| Docs-mestre paths | ../../ ✓ | — | — | ../ ❌ | **../../** |

---

## Verificações OK

| Verificação | Status |
|:---|:---:|
| Tasks M2 (T-016 a T-023 = 8) conferem com TASKS v2.4 | ✅ |
| Tasks M3 (T-024 a T-038 = 15) conferem com TASKS v2.4 | ✅ |
| Total 23 tasks = 8+15 | ✅ |
| T-021 marcado Should — consistente com TASKS v2.4 | ✅ |
| Features listadas (8) conferem com PRD v1.7 §4.5 | ✅ |
| Sprint Goal alinhado com TASKS v2.4 §2 M2+M3 | ✅ |
| Dependências Sprint 2→3→4 corretas | ✅ |
| Branch `feature/sprint-03-portal-admin` ativa e limpa | ✅ |
| Infraestrutura base pronta (JWT, RBAC, RLS, Auditoria) | ✅ |
| 11 tabelas criadas (V001) — todas as entidades da Sprint 3 têm tabelas | ✅ |
| Enums TenantStatus, SubscriptionStatus, Recurrence, Segment já existem | ✅ |
| BaseRepository funcional com soft-delete + tenant filter | ✅ |
| GlobalExceptionHandler (RFC 7807) implementado | ✅ |
| Migration V003 (RLS 5 tabelas) ativa | ✅ |
| T-033 DONE: "Change-plan sem gap" cobre RN07-02 implicitamente | ✅ |
| CAVEMAN anterior: 6/10 achados totalmente corrigidos | ✅ |

---

## Impacto das Correções nos Docs-Mestre

As seguintes alterações nos docs-mestre (já aplicadas) impactam a Sprint 3:

| Alteração no Doc-Mestre | Impacto na Sprint 3 |
|:---|:---|
| PRD v1.7: branching por sprint (§8.4) | Branch `feature/sprint-03-portal-admin` é a canônica |
| ARCHITECTURE v2.0: refatoração C4→Mermaid (887 linhas) | SPRINT-CARD referencia v1.4 (desatualizado) |
| TASKS v2.4: 16/80 tarefas concluídas | Sprint 3 é a atual, 23 tarefas pendentes |
| SPECS v1.6: RN07-03 incluso em F02-04 | Já refletido na tabela de features do SPRINT-CARD |
| TEST_PLAN v2.4: 56 cenários SS3.1-SS3.8 | SPRINT-TEST-SUITE lista 56 mas sumário diz 55 |

---

## Plano de Correção

### Prioridade 1 — Bugs (corrigir antes de iniciar implementação)
| # | Arquivo | Linha | Ação |
|---|---------|-------|------|
| 1 | SPRINT-CARD.md | L119 | "RNs implementadas: 20" → "RNs implementadas: 21" |
| 2 | SPRINT-CARD.md | L120 | "Cenários de teste: 55" → "Cenários de teste: 56" |
| 3 | SPRINT-REVIEW.md | L92 | "20 RNs" → "21 RNs" |
| 4 | SPRINT-TEST-SUITE.md | L112-118 | Integration: 19→20, Total: 55→56 |
| 5 | SPRINT-DEVELOPMENT-PLANNING-DRAFT.md | L7 | Corrigir paths relativos: `../` → `../../` |
| 6 | SPRINT-DEVELOPMENT-PLANNING-DRAFT.md | L12 | Status "Não iniciada" → "Em andamento — 16/07/2026" |
| 7 | TenantAwareDataSource.java | L57 | Concatenation → PreparedStatement (SQL injection) |
| 8 | application.yml | L8,L53 | Consolidar chaves `spring:` duplicadas |

### Prioridade 2 — Riscos (corrigir durante a sprint)
| # | Arquivo | Linha | Ação |
|---|---------|-------|------|
| 9 | SPRINT-CARD.md | L9 | ARCHITECTURE v1.4 → v2.0 |
| 10 | SPRINT-CARD.md | L9 | PRD v1.5 → v1.7 |
| 11 | SPRINT-CARD.md | L56 | Adicionar "RN07-02" no DONE de T-033 |
| 12 | SPRINT-REVIEW.md | L88-93 | Target endpoints: 17 → 18 (verificar contagem SPECS) |
| 13 | SPRINT-DEVELOPMENT-PLANNING-DRAFT.md | L68 | RNs: 20 → 21 |
| 14 | SPRINT-DEVELOPMENT-PLANNING-DRAFT.md | L70 | Testes: 55 → 56 |
| 15 | SPRINT-DEVELOPMENT-PLANNING-DRAFT.md | L156-180 | Seção 5: remover índices já existentes (5/6), manter só `idx_plan_status` |
| 16 | SPRINT-DEVELOPMENT-PLANNING-DRAFT.md | L185-230 | Templates: alinhar com padrão `BaseRepository` |
| 17 | pom.xml | — | Adicionar springdoc-openapi para documentação dos 18 endpoints |
| 18 | Código | — | Criar padrão de teste de integração com Testcontainers (1 exemplo) |
| 19 | Código | — | Criar RowMapper para Tenant (primeiro a ser implementado) |

### Prioridade 3 — Nits (corrigir quando oportuno)
| # | Arquivo | Linha | Ação |
|---|---------|-------|------|
| 20 | SPRINT-CARD.md | L9 | PRD v1.5 → v1.7 |
| 21 | SPRINT-REVIEW.md | L97-101 | Preencher ou remover seção de bloqueadores vazia |
| 22 | SPRINT-TEST-SUITE.md | L112-118 | Explicitar que F01-03 não tem Security (é Should) |
| 23 | SPRINT-DEVELOPMENT-PLANNING-DRAFT.md | L135-149 | Recalibrar day-by-day (Dia 1 Migration V004 desnecessário) |
| 24 | SPRINT-DEVELOPMENT-PLANNING-DRAFT.md | Seção 12 | Preencher Log de Decisões ou remover se vazio ao final |
| 25 | SPRINT-DEVELOPMENT-PLANNING-DRAFT.md | Anexo A | Marcar itens já concluídos no checklist |
| 26 | DOCS-SPRINT-CAVEMAN-REVIEW.md | Footer | Esta é a v2.0 da review; adicionar changelog |

---

## Status da Correção da Review Anterior (15/07/2026)

| # | Achado Original | Status | Nota |
|---|----------------|:---:|------|
| 1 | SPRINT-CARD L42: "11→7 endpoints" | ✅ Corrigido | |
| 2 | SPRINT-CARD L110: "15→20 RNs" | ⚠️ Parcial | Corrigido para 20, mas valor real é 21 |
| 3 | SPRINT-REVIEW L90: "15→20 RNs" | ⚠️ Parcial | Corrigido para 20, mas valor real é 21 |
| 4 | SPRINT-CARD L91: risco T-021 | ✅ Corrigido | |
| 5 | SPRINT-CARD L69: RN07-03 em F02-04 | ✅ Corrigido | |
| 6 | SPRINT-CARD L48: RN07-02 no DONE T-033 | ❌ Não corrigido | DONE ainda sem ref explícita |
| 7 | SPRINT-TEST-SUITE sumário: 48→55/56 | ⚠️ Parcial | Corrigido para 55, mas valor real é 56 |
| 8 | SPRINT-CARD L8: +PRD +ARCHITECTURE | ✅ Corrigido | |
| 9 | SPRINT-CARD L116: TASKS v2.0→v2.3 | ✅ Corrigido | Agora v2.4 |
| 10 | SPRINT-TEST-SUITE L135: TEST_PLAN v2.0→v2.3 | ✅ Corrigido | Agora v2.4 |

---

## Recomendações Estratégicas

1. **Antes de iniciar qualquer implementação:** corrigir os 8 bugs da Prioridade 1. São correções mecânicas que evitam arrastar números errados por toda a sprint.

2. **Primeiro código a escrever:** `TenantRowMapper` — é a dependência oculta de T-020 (TenantRepository), que por sua vez é pré-requisito para T-025 e T-026. Nenhum repository funciona sem RowMapper.

3. **Testes de integração:** escrever 1 teste de integração com Testcontainers como padrão antes de implementar os 20 cenários. O projeto tem a dependência mas nunca a usou — validar que funciona no ambiente local.

4. **Planejamento recalibrado:** com os índices V002 já existentes e enums prontos, o Dia 1 do plano pode ser eliminado. A sprint começa efetivamente no Dia 2 (T-016 DashboardRepository).

5. **SQL injection em TenantAwareDataSource:** embora o `tenant_id` venha do JWT (validado), o uso de concatenação de string viola o princípio de defense-in-depth. Corrigir com PreparedStatement — é uma mudança de ~5 linhas.

---

## Changelog

| Versão | Data | Mudanças |
|:---|:---|:---|
| v1.0 | 15/07/2026 | Review inicial: 10 achados em 3 artefatos (docs-mestre v1.4/v1.5/v2.3) |
| v2.0 | 16/07/2026 | Refeita integralmente: triangulação docs-mestre×docs-sprint×código. 26 achados em 5 artefatos + código. Docs-mestre v1.7/v1.6/v2.4/v2.4/v2.0. |

---

🤖 *Revisão gerada em 16/07/2026. 26 achados em 5 artefatos + código. Fonte: triangulação PRD v1.7, SPECS v1.6, TASKS v2.4, TEST_PLAN v2.4, ARCHITECTURE v2.0 + scan completo do código-fonte (43 arquivos analisados).*
