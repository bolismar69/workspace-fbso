# SPRINT-CARD: Sprint 2 — Segurança Cross-Cutting

- **Sprint:** 2 de 7
- **Marco:** Pre-M2 — Segurança
- **Datas:** 07/08/2026 → 15/08/2026
- **Duração:** 6 dias úteis
- **Responsável:** A definir
- **Documentos-mestre:** [TASKS.md](../../TASKS.md) · [SPECS.md](../../SPECS.md) · [TEST_PLAN.md](../../TEST_PLAN.md) · [ARCHITECTURE.md](../../ARCHITECTURE.md)

---

> ✅ Sprint concluída na branch `feature/java-fbso-platform-admin` (histórico). Estratégia atual: uma branch por sprint — veja [PRD.md §8.4](../../PRD.md#84-estratégia-de-branching--uma-branch-por-sprint).

## 🎯 Sprint Goal

**"JWT autentica via Keycloak RS256. TenantContext isola requisições por tenant_id. PostgreSQL RLS garante isolamento no nível do banco (defesa em profundidade). @RequiresPermission bloqueia acessos não autorizados (403). @Auditable grava auditoria de forma assíncrona. GlobalExceptionHandler retorna erros RFC 7807 em PT-BR."**

---

## ⚠️ Alerta Crítico

> Esta sprint é o **gargalo do projeto inteiro**. NENHUM endpoint de negócio (M2 a M6) pode ser implementado antes que este pipeline de segurança esteja funcional. A cadeia crítica do projeto passa por T-009 → T-010 → T-013.

---

## 📋 Sprint Backlog

| ID | Tarefa | Est. | Critério DONE |
|:---|:---|:---:|:---|
| **T-009** | `SecurityConfig.java`: Spring Security + JWT (Keycloak RS256 via JWKS). Desabilitar CSRF (API stateless), configurar CORS | 2d | Sem token → 401. Token válido → autenticado. Token inválido → 401. CORS permite origem do frontend |
| **T-010** | `JwtAuthenticationFilter.java` (OncePerRequestFilter): extrair JWT Bearer, validar assinatura RS256, validar exp, extrair claims (tenant_id, user_id, roles, business_unit_ids, modules), setar SecurityContext | 2d | Filter em toda req. exceto `/actuator/health`. Claims extraídas. SecurityContext populado |
| **T-011** | `TenantContext.java` (ThreadLocal): tenant_id, user_id, roles, business_unit_ids, modules. Métodos get/set/clear | 0.5d | TenantContext.getTenantId() retorna tenant_id. Contexto limpo após requisição |
| **T-012** | `TenantIsolationAspect.java`: interceptar @Repository via AOP, injetar `WHERE tenant_id = ?`. Ordem 1 | 1.5d | [SUBSTITUÍDO por T-015.1] Sem tenant_id → SecurityException. Query injetada. Tenants diferentes → dados diferentes |
| **T-015.1** 🆕 | Migration V003: PostgreSQL RLS em 5 tabelas + `TenantAwareDataSource` (proxy HikariCP) | 1.5d | ✅ `ALTER TABLE ... ENABLE ROW LEVEL SECURITY`. Politica `tenant_isolation` criada. DataSource proxy funcional. 33/33 testes passando |
| **T-013** | `@RequiresPermission(resource, action)` + `RbacAspect.java`: verificar role × resource × action conforme RN10-01 | 2d | Anotação bloqueia sem permissão. Role válida × resource × action. 403 com JSON amigável |
| **T-014** | `@Auditable(entityType, action)` + `AuditAspect.java`: capturar antes/depois, gravar em audit_log ASSÍNCRONO (@Async) | 1.5d | Registro gerado por operação. Valores capturados. Async não bloqueia thread principal |
| **T-015** | `GlobalExceptionHandler.java` (@ControllerAdvice) RFC 7807 + hierarchy: BusinessException(422), DuplicateCnpjException, InvalidStatusTransitionException, PlanHasActiveSubscribersException, TenantNotFoundException, PermissionDeniedException(403) | 1d | type, title, status, detail. Sem stack trace. Mensagens PT-BR |

**Total:** 8 tarefas · ~12 dias-homem

---

## 📦 Entregáveis da Sprint

1. **`SecurityConfig.java`** — Spring Security + JWT + CORS
2. **`JwtAuthenticationFilter.java`** — OncePerRequestFilter com validação RS256
3. **`TenantContext.java`** — ThreadLocal com claims do JWT
4. **`TenantIsolationAspect.java`** — AOP que injeta `WHERE tenant_id = ?`
5. **`@RequiresPermission`** + **`RbacAspect.java`** — anotação + aspecto RBAC
6. **`@Auditable`** + **`AuditAspect.java`** — anotação + aspecto de auditoria assíncrona
7. **`GlobalExceptionHandler.java`** — @ControllerAdvice RFC 7807
8. **Migration V003** — PostgreSQL RLS em 5 tabelas (subscription, user, business_unit, product_service, audit_log) + config `app.current_tenant_id` no JwtFilter

---

## ✅ Definition of Done (Sprint-Level)

- [ ] Requisição sem token JWT → 401 `{"title": "Token de acesso não informado"}`
- [ ] Token JWT válido → autenticado, TenantContext populado
- [ ] Token JWT expirado ou assinatura inválida → 401
- [ ] `TenantContext` limpo após cada requisição (sem vazamento entre threads)
- [ ] `TenantIsolationAspect` injeta `tenant_id` em toda query de repository
- [ ] PostgreSQL RLS habilitado em 5 tabelas (V003). Politica `tenant_isolation` ativa
- [ ] `JwtAuthenticationFilter` configura `app.current_tenant_id` na sessão PostgreSQL
- [ ] INSERT com tenant_id divergente do contexto → rejeitado pelo PostgreSQL
- [ ] `@RequiresPermission` bloqueia acesso não autorizado → 403 amigável
- [ ] `@Auditable` grava em `audit_log` de forma assíncrona
- [ ] `GlobalExceptionHandler` retorna RFC 7807 em todos os erros
- [ ] Nenhum stack trace em respostas HTTP

---

## ⚠️ Riscos e Bloqueadores

| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---:|:---|
| JWKS endpoint do Keycloak inacessível | Média | Crítico | Configurar cache de JWKS (TTL 5min). Fallback: chave pública em variável de ambiente |
| Vazamento de TenantContext entre requisições | Média | Crítico | `TenantContext.clear()` em `finally` no filter. Teste de concorrência com múltiplos tenants |
| `TenantIsolationAspect` não cobre queries nativas | Alta | Alto | Auditoria de todas as queries SQL na fase de teste. Proibir `@Query(nativeQuery=true)` sem `tenant_id` explícito |
| Lógica assíncrona do `AuditAspect` perde registros em crash | Baixa | Médio | Trade-off aceito (ADR-L03). Logar warning se fila cheia |

---

## 🔗 Dependências

- **Pré-requisitos:** Sprint 1 concluída (T-001, T-006, T-007, T-008). Keycloak disponível.
- **Sucessor:** Sprint 3 (Portal Admin) — depende de TODO o pipeline de segurança.
- **Features que dependem desta sprint:** TODAS (M2 a M6).

---

## 📊 Métricas da Sprint

| Métrica | Meta |
|:---|:---:|
| Tasks completadas | 8/8 |
| Tempo de resposta do filter JWT | ≤ 50ms (p95) |
| Testes de segurança | 3+ cenários por aspecto |

---

🤖 *Gerado a partir de TASKS.md v2.3. Revisão Caveman em 15/07/2026 (DOCS-SERVICE-SPRINTS-CAVEMAN-REVIEW.md): RLS 11→5 tabelas, tasks 7→8, docs-mestre expandido. Sprint crítica — o pipeline de segurança habilita todo o resto do projeto.*
