# SPRINT-TEST-SUITE: Sprint 2 — Segurança Cross-Cutting

- **Sprint:** 2 de 7
- **Origem:** [TEST_PLAN.md](../../TEST_PLAN.md) §4 — Pipeline de Segurança
- **Total de cenários:** 26

> ⚠️ Sprint crítica de segurança. Testes aqui são pré-requisito para qualquer endpoint de negócio.

---

## 1. JWT Authentication Filter

| ID | Descrição | Nível | Critério |
|:---|:---|:---|:---|
| TC-S2-001 | Requisição sem header Authorization → 401 | Unit | Filter retorna 401. Body RFC 7807 com title "Token de acesso não informado" |
| TC-S2-002 | JWT com assinatura RS256 inválida → 401 | Unit | Filter detecta `SignatureException`. 401 "Token inválido ou expirado" |
| TC-S2-003 | JWT expirado (`exp` no passado) → 401 | Unit | Filter detecta `ExpiredJwtException`. 401 |
| TC-S2-004 | JWT com `alg: none` → 401 (rejeitar) | Segurança | Filter rejeita token sem assinatura. 401 |
| TC-S2-005 | JWT válido → extrai claims (tenant_id, user_id, roles, business_unit_ids, modules) | Unit | SecurityContext populado. Claims acessíveis via TenantContext |
| TC-S2-006 | Endpoint `/actuator/health` NÃO passa pelo filter | Unit | Health check acessível sem token |

---

## 2. TenantContext e TenantIsolationAspect

| ID | Descrição | Nível | Critério |
|:---|:---|:---|:---|
| TC-S2-007 | `TenantContext.getTenantId()` retorna valor do JWT | Unit | ThreadLocal populado pelo filter |
| TC-S2-008 | `TenantContext.clear()` limpa após requisição | Unit | Após response, ThreadLocal vazio. Sem vazamento entre requisições |
| TC-S2-009 | `TenantIsolationAspect` injeta `WHERE tenant_id = ?` em query | Unit | SQL contém `tenant_id = ?`. Parâmetro = TenantContext.getTenantId() |
| TC-S2-010 | Repository sem tenant_id no contexto → SecurityException | Unit | TenantContext vazio → exceção lançada ANTES da query executar |
| TC-S2-011 | Dados de tenant-A não visíveis para tenant-B | Integração | Query com TenantContext(tenant-A) retorna apenas dados de tenant-A |

---

## 2.5 PostgreSQL Row-Level Security (RLS)

| ID | Descrição | Nível | Critério |
|:---|:---|:---|:---|
| TC-S2-022 | Migration V003: RLS habilitado em 5 tabelas | Integração | `rowsecurity=true` para todas as tabelas com tenant_id |
| TC-S2-023 | Politica `tenant_isolation` criada em cada tabela | Integração | 11 politicas. `USING (tenant_id = current_setting('app.current_tenant_id')::UUID)` |
| TC-S2-024 | `JwtAuthenticationFilter` seta `app.current_tenant_id` na sessão PostgreSQL | Unit | Após filter, `SHOW app.current_tenant_id` retorna UUID do tenant |
| TC-S2-025 | INSERT com tenant_id divergente → rejeitado pelo PostgreSQL | Integração | Tenant-A tenta INSERT com tenant_id='B' → erro POLICY violation |
| TC-S2-026 | SELECT sem WHERE tenant_id → filtrado automaticamente pelo RLS | Integração | 2 tenants com dados. SELECT * FROM business_unit → apenas BUs do tenant corrente |

---

## 3. RBAC — @RequiresPermission + RbacAspect

| ID | Descrição | Nível | Critério |
|:---|:---|:---|:---|
| TC-S2-012 | `@RequiresPermission(resource="PRODUCT", action="edit")` bloqueia OPERATOR (só lê) | Unit | RbacAspect lança PermissionDeniedException. GlobalExceptionHandler → 403 |
| TC-S2-013 | `@RequiresPermission` permite acesso se role tem resource+action | Unit | ADMIN_TENANT acessa qualquer recurso. RbacAspect permite execução |
| TC-S2-014 | Resposta 403 segue RFC 7807 e é amigável (sem stack trace) | Unit | `{"title":"Acesso negado","detail":"Você não tem permissão para executar esta operação.","status":403}` |

---

## 4. Auditoria — @Auditable + AuditAspect

| ID | Descrição | Nível | Critério |
|:---|:---|:---|:---|
| TC-S2-015 | `@Auditable(entityType="TENANT", action="CREATED")` grava registro em audit_log | Unit | AuditEntry criado. previous_value e new_value capturados |
| TC-S2-016 | Auditoria executa de forma ASSÍNCRONA (@Async) | Unit | Thread principal não bloqueada. Verificar que `@Async` está ativo |
| TC-S2-017 | Falha na auditoria NÃO interrompe operação principal | Unit | Exceção no AuditAspect capturada. Operação principal conclui com sucesso |

---

## 5. GlobalExceptionHandler (RFC 7807)

| ID | Descrição | Nível | Critério |
|:---|:---|:---|:---|
| TC-S2-018 | `BusinessException` → 422 com type, title, status, detail | Unit | RFC 7807 completo. Campos `fields` se houver erros de validação |
| TC-S2-019 | `PermissionDeniedException` → 403 com mensagem em PT-BR | Unit | `"Acesso negado"`. Sem detalhes internos |
| TC-S2-020 | `Exception` genérica → 500 com mensagem genérica, SEM stack trace | Unit | `"Erro interno do servidor"`. Stack trace NUNCA no response |
| TC-S2-021 | `MethodArgumentNotValidException` (Bean Validation) → 400 | Unit | Campos com erro listados em `fields[]`. Mensagens em PT-BR |

---

## 📊 Resumo

| Nível | Cenários |
|:---|:---:|
| Unit | 19 |
| Integração | 5 |
| Segurança | 2 |
| **Total** | **26** |

---

## 🔗 Pipeline de Segurança — Ordem de Execução

```
T-009 (SecurityConfig) → T-010 (JWT Filter) → T-011 (TenantContext)
                                              → T-013 (@RequiresPermission)
                         T-007 (BaseRepo)    → T-012 (TenantIsolation)
                         T-015 (GlobalExceptionHandler)
                         (independente)      → T-014 (@Auditable)
```

> Os testes devem validar o pipeline completo: JWT → TenantContext → RBAC → TenantIsolation → Auditoria.

---

🤖 *Extraído de TEST_PLAN.md v2.3. Revisão Caveman em 15/07/2026 (DOCS-SERVICE-SPRINTS-CAVEMAN-REVIEW.md): RLS 11→5 tabelas/políticas, TEST_PLAN ref v2.0→v2.3. Execute estes cenários antes de considerar a Sprint 2 concluída.*
