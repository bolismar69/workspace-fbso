# DETAIL-LEVEL-TEST-STRATEGY-DEFINITION — Estratégia de Testes Detail-Level

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Data:** 31/07/2026
- **Fase:** F6 — Downstream Architecture Refinement
- **Referências:** [Arquitetura (F2)](./DETAIL-LEVEL-ARCHITECTURE-DEFINITION.md), [Segurança (F3)](./DETAIL-LEVEL-SECURITY-DEFINITION.md)

---

## 1. Pirâmide de Testes

```
        ╱  E2E  ╲          Playwright — fluxos críticos (10% das US)
       ╱──────────╲
      ╱ Integração ╲       Spring Boot Test + Testcontainers (≥60%)
     ╱──────────────╲
    ╱   Unitários    ╲     JUnit 5 + Mockito (backend) / Jest (frontend) — ≥80%
   ╱──────────────────╲
```

| Camada | Meta Cobertura | Ferramenta Backend | Ferramenta Frontend |
|:---|:---:|:---|:---|
| Unitários | ≥80% | JUnit 5 + Mockito | Jest + Testing Library |
| Integração | ≥60% | Spring Boot Test + Testcontainers (PostgreSQL, Keycloak, Redis) | MSW (Mock Service Worker) |
| E2E | Fluxos críticos | Playwright | Playwright |
| Performance | Smoke test | k6 | Lighthouse (LCP, CLS) |
| Segurança | Scan | Semgrep (SAST), Gitleaks (Secrets), OWASP ZAP (DAST) | npm audit, CSP header check |

---

## 2. Matriz de Cobertura por Épico

| Épico | US | Unit | Integração | E2E | Perf | Security |
|:---|---:|:---:|:---:|:---:|:---:|:---:|
| EP-0001 Portal Admin | 7 | ✅ Todas | ✅ Dashboard queries, filtros | ✅ Fluxo dashboard completo | ✅ k6 smoke dashboard | ✅ Semgrep |
| EP-0002 Clientes e Assinaturas | 16 | ✅ Todas | ✅ CRUD tenants, status machine, planos, assinaturas, audit trigger | ✅ Ativação tenant + suspensão | — | ✅ Semgrep |
| EP-0003 RBAC | 16 | ✅ Todas | ✅ PermissionEvaluator, RLS, Keycloak roles | ✅ Convite → atribuir papel → validar acesso | — | ✅ Semgrep + OWASP ZAP |
| EP-0004 Portal Cliente | 23 | ✅ Todas | ✅ Auth OIDC, onboarding wizard, BU hierarchy, product CRUD, app switcher | ✅ Onboarding completo + upgrade self-service | ✅ k6 LCP/Load | ✅ Semgrep + OWASP ZAP |

---

## 3. Casos de Teste de Aceitação (por Feature Crítica)

### FEAT-EP-0002-0002 — Gestão de Status do Tenant

| Cenário | Dado | Quando | Então |
|:---|:---|:---|:---|
| Suspender tenant ativo | Tenant status=ACTIVE | Admin FBSO clica "Suspender" | Status muda para SUSPENDED; todos os usuários do tenant bloqueados |
| Reativar tenant suspenso | Tenant status=SUSPENDED | Admin FBSO clica "Reativar" | Status volta para ACTIVE; usuários desbloqueados |
| Bloqueio imediato de usuários | Tenant com 5 usuários ativos | Admin suspende o tenant | Em até 60 segundos, todos os 5 usuários não conseguem autenticar |

### FEAT-EP-0003-0002 — Papéis e Permissões (RBAC)

| Cenário | Dado | Quando | Então |
|:---|:---|:---|:---|
| Admin acessa tudo | Usuário com ROLE_TENANT_ADMIN | Acessa qualquer endpoint do tenant | 200 OK |
| Operador sem acesso a gestão de usuários | Usuário com ROLE_TENANT_OPERATOR | Tenta POST /api/v1/users | 403 Forbidden |
| Auditor vê tudo, não edita nada | Usuário com ROLE_TENANT_AUDITOR | Tenta PUT em qualquer recurso | 403 Forbidden; GET funciona |

### FEAT-EP-0004-0002 — Onboarding Guiado

| Cenário | Dado | Quando | Então |
|:---|:---|:---|:---|
| Onboarding completo | Tenant recém-criado, status=PENDING_ONBOARDING | Admin tenant completa os 4 passos | Status muda para ACTIVE; tela de boas-vindas exibida |
| Abandono e retomada | Tenant no passo 2/4 | Fecha o navegador e volta depois | Retoma do passo 2; dados já preenchidos preservados |

---

## 4. Estratégia de Automação

| Momento | O que Roda | Bloqueia Merge? |
|:---|:---|:---:|
| **PR (push)** | Unit tests (JUnit + Jest), SAST (Semgrep), Secret scan (Gitleaks) | ✅ Sim |
| **Merge → Staging** | Integration tests (Testcontainers), Smoke E2E (Playwright subset), k6 smoke | ✅ Sim |
| **Nightly (02:00)** | Full E2E (Playwright), Performance (k6 full), OWASP ZAP | ⚠️ Alerta |
| **Release Tag** | Todos os testes acima + manual approval + post-deploy health check | ✅ Sim |

---

## 5. Quality Gates

| Gate | Critério | Bloqueia? |
|:---|:---|:---:|
| **PR Gate** | Unit tests ≥80% coverage. SAST 0 críticas. Gitleaks 0 findings | ✅ Sim |
| **Staging Gate** | Integration tests pass ≥60%. E2E críticos passam. k6 smoke p99<500ms | ✅ Sim |
| **Release Gate** | E2E 100%. OWASP ZAP 0 críticas. Multi-tenant isolation tests pass. Manual QA approval | ✅ Sim |

---

## 6. Testes de Isolamento Multi-Tenant

| Cenário | Verificação |
|:---|:---|
| Tenant A não vê dados do Tenant B | Autenticar como Tenant A, chamar GET /api/v1/tenants/{tenant_b_id} → 404 |
| RLS ativo em todas as tabelas | Query direta SQL sem SET app.current_tenant_id → zero resultados |
| Header injection bypass | Request sem X-Tenant-ID → 403 Forbidden |
| Admin FBSO vê todos os tenants | Autenticar como FBSO Admin, GET /api/v1/tenants → retorna todos (admin bypass) |

---

🤖 *Documento gerado pelo Test Specialist — Fase 6 do Downstream Architecture Refinement.*
