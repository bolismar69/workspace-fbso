# SPRINT-TEST-PLANNING.md — Plano de Testes: Sprint 2

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** `PRJ-FIN-2026-0003-SAAS-FBSO-ORG`
- **Sprint:** 2 de 7 — Segurança Cross-Cutting
- **Stack:** Java 25 + Spring Boot 3.5.1 + JUnit 5 + Mockito
- **Data:** 14 de Julho de 2026

---

## 1. Visão Geral

- **Tasks implementadas:** 8 (T-009 a T-015, T-015.1)
- **Cenários mapeados:** 26 (SPRINT-TEST-SUITE.md)
- **Meta de cobertura:** ≥ 80%
- **Ferramentas:** JUnit 5 + Mockito (unit), Spring Boot Test (integration)

## 2. Mapeamento Task → Cenários

| Task | Cenário(s) | Nível | Status |
|:---|:---|:---|:---:|
| T-009 | TC-S2-006 — Health check sem token | Unit | ⬜ |
| T-010 | TC-S2-001 — Sem header → 401 | Unit | ⬜ |
| T-010 | TC-S2-002 — JWT assinatura inválida → 401 | Unit | ⬜ |
| T-010 | TC-S2-003 — JWT expirado → 401 | Unit | ⬜ |
| T-010 | TC-S2-004 — JWT alg:none → 401 | Unit | ⬜ |
| T-010 | TC-S2-005 — JWT válido extrai claims | Unit | ⬜ |
| T-011 | TC-S2-007 — TenantContext.getTenantId() | Unit | ⬜ |
| T-011 | TC-S2-008 — TenantContext.clear() | Unit | ⬜ |
| T-012 | TC-S2-009 — TenantIsolation injeta tenant_id | Unit | ⬜ |
| T-012 | TC-S2-010 — Sem tenant_id → SecurityException | Unit | ⬜ |
| T-012 | TC-S2-011 — Dados tenant-A ≠ tenant-B | Integration | ⬜ |
| T-015.1 | TC-S2-022 — RLS habilitado em 5 tabelas | Integration | ⬜ |
| T-015.1 | TC-S2-023 — 5 políticas tenant_isolation | Integration | ⬜ |
| T-015.1 | TC-S2-024 — TenantAwareDataSource config | Unit | ⬜ |
| T-015.1 | TC-S2-025 — Rollback V003 remove RLS | Integration | ⬜ |
| T-015.1 | TC-S2-026 — Admin FBSO acesso cross-tenant | Unit | ⬜ |
| T-013 | TC-S2-012 — OPERATOR sem permissão → 403 | Unit | ⬜ |
| T-013 | TC-S2-013 — ADMIN acesso total | Unit | ⬜ |
| T-013 | TC-S2-014 — 403 RFC 7807 amigável | Unit | ⬜ |
| T-014 | TC-S2-015 — @Auditable grava audit_log | Unit | ⬜ |
| T-014 | TC-S2-016 — Auditoria assíncrona (@Async) | Unit | ⬜ |
| T-014 | TC-S2-017 — Falha auditoria não interrompe | Unit | ⬜ |
| T-015 | TC-S2-018 — BusinessException → 422 | Unit | ⬜ |
| T-015 | TC-S2-019 — PermissionDeniedException → 403 | Unit | ⬜ |
| T-015 | TC-S2-020 — Exception → 500 sem stack trace | Unit | ⬜ |
| T-015 | TC-S2-021 — MethodArgumentNotValid → 400 | Unit | ⬜ |

## 3. Estratégia por Nível

### 3.1 Testes Unitários
- **Ferramenta:** JUnit 5 + Mockito
- **Padrão:** AAA (Arrange-Act-Assert) com @Nested para agrupamento
- **Localização:** `src/test/java/.../unit/`
- **O que mockar:** JwtDecoder, HttpServletRequest, HttpServletResponse, JdbcTemplate
- **O que NÃO mockar:** TenantContext (usar valores reais), ErrorResponse (record)

### 3.2 Testes de Integração
- **Ferramenta:** Spring Boot Test + MockMvc
- **Localização:** `src/test/java/.../integration/security/`
- **Foco:** Pipeline JWT → TenantContext → RBAC (end-to-end)

## 4. Comandos

- **Unit:** `mvn test -Dtest="**/unit/**"`
- **Integration:** `mvn test -Dtest="**/integration/security/**"`
- **All:** `mvn test`
- **Coverage:** `mvn jacoco:check` (condicionado à compatibilidade JaCoCo × Java 25)

---


🤖 *Revisão Caveman em 15/07/2026 (DOCS-SERVICE-SPRINTS-CAVEMAN-REVIEW.md): tasks 7→8, adicionado mapeamento T-015.1 (TC-S2-022 a TC-S2-026), cenários 21→26.*
