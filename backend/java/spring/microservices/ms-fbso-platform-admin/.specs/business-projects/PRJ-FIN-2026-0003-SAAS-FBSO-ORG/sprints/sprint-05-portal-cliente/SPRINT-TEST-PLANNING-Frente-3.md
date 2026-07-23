# SPRINT-TEST-PLANNING-Frente-3.md — Plano de Testes: Sprint 5 — Frente 3

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 5 de 7 — sprint-05-portal-cliente
- **Frente:** Frente 3 — Features (🎯 Features)
- **Stack:** Java 25 · Spring Boot 3.5.14 · PostgreSQL 17 · Keycloak 26 · Flyway 12.11.0
- **Data do planejamento:** 2026-07-23

---

## 1. Visão Geral

- **Tasks implementadas:** 12 backend (10 features + 2 NO-OP)
- **Features cobertas:** F04-01 (Login/Auth) · F04-02 (Onboarding) · F04-03 (Dashboard Cliente) · F04-04 (App Switcher)
- **Meta de cobertura:** ≥ 80%
- **Ferramentas:** JUnit 5 + Mockito (unitários) · Testcontainers + PostgreSQL (integração)
- **Testes existentes:** 213 testes — base sólida para não regredir

---

## 2. Mapeamento Task → Cenários de Teste

| Task | Cenário(s) | Nível | Ferramenta |
|:---|:---|:---|:---|
| T-057 | Keycloak realm funcional, JWT claims corretos | Manual | Keycloak Admin Console + jwt.io |
| T-058 | Login redirect, forgot-password, reset-password | Unit | JUnit 5 + Mockito |
| T-059 | Rate limiting: 5 falhas→bloqueio, reset 15min, mensagem tempo | Unit | JUnit 5 + Caffeine |
| T-060 | Onboarding 4 passos, 6 edge cases, transição ACTIVE | Unit | JUnit 5 + Mockito |
| T-061 | 4 endpoints REST, validação CNPJ, @RequiresPermission | Unit | MockMvc |
| T-062 | BU Matriz: isMatrix=true, parentId=null, CNPJ validado | Unit | JUnit 5 + Mockito |
| T-063 | Cards dashboard: unidades, produtos, plano, notificações | Unit | JUnit 5 + Mockito |
| T-064 | GET /dashboard/client/summary, /notifications | Unit | MockMvc |
| T-065 | modules[] + business_unit_ids[] no TenantContext | Unit | JUnit 5 (NO-OP — já implementado) |
| T-066 | GET /auth/me: dados JWT, onboarding status | Unit | MockMvc |
| T-067 | OnboardingService, DashboardClientService, AuthService | Unit | JUnit 5 + Mockito — ≥ 80% |
| T-068 | Fluxo onboarding PENDING→ACTIVE, rate limiting, cross-tenant | Integration | Testcontainers + PostgreSQL |

---

## 3. Estratégia por Nível

### 3.1 Testes Unitários
- **Ferramenta:** JUnit 5 + Mockito + MockMvc
- **Padrão:** AAA (Arrange-Act-Assert)
- **O que mockar:** TenantRepository, BusinessUnitService, JdbcTemplate, Keycloak API calls
- **O que NÃO mockar:** Entidades, DTOs, TenantContext (ThreadLocal)

### 3.2 Testes de Integração
- **Ferramenta:** Testcontainers (PostgreSQL 17)
- **Dados de seed:** Tenant PENDING_ONBOARDING, BusinessUnit ativa, Subscription ativa
- **Cenários críticos:**
  - Fluxo onboarding completo: step-1 → step-2 → step-3 → complete → ACTIVE
  - Rate limiting com PostgreSQL real: 5 falhas → 429
  - Cross-tenant isolation: Tenant A não vê onboarding do Tenant B

---

## 4. Comandos de Execução

| Propósito | Comando |
|:---|:---|
| Unitários | `mvn test` |
| Integração | `mvn test -Dtest="*IntegrationTest"` |
| Coverage | `mvn jacoco:report` |

---

## 5. Resultado da Execução

| Métrica | Resultado |
|:---|---|
| `mvn compile` | ✅ BUILD SUCCESS (7.0s) |
| `mvn test` | ✅ 213 testes: 0 failures, 1 pre-existing error |

---

🤖 *Documento gerado em 2026-07-23 como parte da Fase 3 do PROMPT-EXECUTE-SPRINT-TASKS.md.*
