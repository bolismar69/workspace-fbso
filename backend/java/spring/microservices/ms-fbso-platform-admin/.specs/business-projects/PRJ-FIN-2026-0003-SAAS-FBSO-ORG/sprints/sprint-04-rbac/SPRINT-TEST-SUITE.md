# SPRINT-TEST-SUITE: Sprint 4 — Governança de Acessos (RBAC)

- **Sprint:** 4 de 7
- **Origem:** [TEST_PLAN.md](../../TEST_PLAN.md) §3.9 a §3.12 + §4.1
- **Features:** F03-01 a F03-04 (4 features)
- **Total de cenários:** 27
- **Status atual:** Frente 0 ✅ (infra de testes pronta) · Frentes 1-4 ⬜ (cenários a implementar)
- **Pré-requisitos de teste concluídos na Frente 0:**
  - ✅ REST Assured 5.5.7 disponível (T-097)
  - ✅ FORCE ROW LEVEL SECURITY nas 4 tabelas (T-109) — RLSIsolationTest sem FORCE manual
  - ✅ RbacAspect DB-backed (T-100) — RbacAspectTest reescrevível com @MockBean
  - ✅ Migration V004 seed RBAC (T-105) — dados para testes parametrizados
  - ✅ Connection leak corrigido (T-111) — testes de integração estáveis

---

## 1. F03-01: Gestão de Usuários (7 cenários) — Frente 1

| ID | Descrição | Nível | Ref. TEST_PLAN | Status |
|:---|:---|:---|:---|:---:|
| TC-F03-01-001 | Convidar usuario email unico → OK | Unit | §3.9 | ⬜ |
| TC-F03-01-002 | Email duplicado → 409 | Unit | §3.9 | ⬜ |
| TC-F03-01-003 | Admin nao pode desativar a si mesmo | Unit | §3.9 | ⬜ |
| TC-F03-01-004 | POST /users → 201 | Integração | §3.9 | ⬜ |
| TC-F03-01-005 | Auto-desativacao → 422 | Integração | §3.9 | ⬜ |
| TC-F03-01-006 | E2E: Convidar → email → desativar | E2E | §3.9 | ⬜ |
| TC-F03-01-007 | OPERATOR tenta POST /users → 403 | Segurança | §3.9 | ⬜ |

## 2. F03-02: Matriz de Permissões RBAC (9 cenários) — Frentes 2, 3

| ID | Descrição | Nível | Ref. TEST_PLAN | Status |
|:---|:---|:---|:---|:---:|
| TC-F03-02-001 | ADMIN_TENANT permite qualquer ação | Unit | §3.10 | ⬜ |
| TC-F03-02-002 | AUDITOR permite apenas leitura | Unit | §3.10 | ⬜ |
| TC-F03-02-003 | **Teste parametrizado: cada papel × cada endpoint** | Segurança | §3.10 | ⬜ |
| TC-F03-02-004 | OPERATOR tenta PATCH /products → 403 | Segurança | §3.10 | ⬜ |
| TC-F03-02-005 | AUDITOR tenta POST /tenants → 403 | Segurança | §3.10 | ⬜ |
| TC-F03-02-006 | MANAGER pode criar BU | Integração | §3.10 | ⬜ |
| TC-F03-02-007 | ADMIN_TENANT pode criar BU | Integração | §3.10 | ⬜ |
| TC-F03-02-008 | E2E: Login cada papel, verifica menu | E2E | §3.10 | ⬜ |
| TC-F03-02-009 | Admin FBSO ve todos tenants (cross-tenant) | Segurança | §3.10 | ⬜ |

## 3. F03-03: Vinculação Usuário × Unidade × Módulo (6 cenários) — Frente 3

| ID | Descrição | Nível | Ref. TEST_PLAN | Status |
|:---|:---|:---|:---|:---:|
| TC-F03-03-001 | Atribuir permissao usuario x BU x role | Unit | §3.11 | ⬜ |
| TC-F03-03-002 | Admin acesso implícito todas BUs | Unit | §3.11 | ⬜ |
| TC-F03-03-003 | Usuario sem vinculacao → 403 | Integração | §3.11 | ⬜ |
| TC-F03-03-004 | Usuario BU-1 ve apenas produtos BU-1 | Integração | §3.11 | ⬜ |
| TC-F03-03-005 | Usuario BU-1 tenta acessar produto BU-2 por ID → 404 | Segurança | §3.11 | ⬜ |
| TC-F03-03-006 | Alterar permissao → efeito imediato (RN11-03) | Segurança | §3.11 | ⬜ |

## 4. F03-04: Acesso Condicional — 403 Amigável (5 cenários) — Frente 4

| ID | Descrição | Nível | Ref. TEST_PLAN | Status |
|:---|:---|:---|:---|:---:|
| TC-F03-04-001 | 403 sem detalhes técnicos | Unit | §3.12 | ⬜ |
| TC-F03-04-002 | 403 segue RFC 7807 (type, title, status, detail) | Integração | §3.12 | ⬜ |
| TC-F03-04-003 | E2E: URL proibida → tela 403 amigavel | E2E | §3.12 | ⬜ |
| TC-F03-04-004 | JWT adulterado (elevacao privilegio) → 401 | Segurança | §3.12 | ⬜ |
| TC-F03-04-005 | URL proibida → 403 (nao 404 — RN12-01) | Segurança | §3.12 | ⬜ |

---

## 5. Testes RBAC Parametrizados (§4.1 do TEST_PLAN)

> **Infra pronta (Frente 0):** REST Assured disponível, seed V004 carrega matriz completa.

```
Papel × Endpoint Proibido (20+ combinações):

OPERATOR × POST   /api/v1/tenants        → 403
OPERATOR × POST   /api/v1/plans          → 403
OPERATOR × POST   /api/v1/users          → 403
OPERATOR × PATCH  /api/v1/products       → 403
OPERATOR × POST   /api/v1/business-units → 403
AUDITOR  × POST   /api/v1/tenants        → 403
AUDITOR  × PATCH  /api/v1/tenants        → 403
AUDITOR  × POST   /api/v1/plans          → 403
AUDITOR  × POST   /api/v1/users          → 403
MANAGER  × POST   /api/v1/tenants        → 403
MANAGER  × POST   /api/v1/plans          → 403
MANAGER  × POST   /api/v1/users          → 403
MANAGER  × GET    /api/v1/audit          → 403
```

---

## 6. Testes de Infra (Frente 0 — Já Concluídos)

> Validados durante a execução da Frente 0. Não fazem parte dos 27 cenários acima.

| Item | Descrição | Status |
|:---|:---|:---:|
| RLS FORCE | FORCE ROW LEVEL SECURITY nas 4 tabelas (V003 atualizado) | ✅ |
| JWT issuer | JwtValidators.createDefaultWithIssuer() ativo | ✅ |
| Connection pool | TenantAwareDataSource sem leak (UUID validado antes da conexão) | ✅ |
| Tenant isolation | SubscriptionService valida tenant_id URL vs JWT | ✅ |
| Cache strategy | Matriz em memória (sem TTL) — conforme RN11-03 | ✅ |

---

## 📊 Resumo

| Nível | Cenários | Status |
|:---|:---:|:---:|
| Unit | 7 | ⬜ |
| Integração | 8 | ⬜ |
| E2E | 3 | ⬜ |
| Segurança | 9 | ⬜ |
| **Total** | **27** | **0/27 implementados** |

---

## 🔗 RNs Cobertas

| RN | Descrição | Cenários | Frente |
|:---|:---|:---:|:---:|
| RN09-01 | Convite expira 7 dias | F03-01 | 1 |
| RN09-02 | Email único por tenant | F03-01 | 1 |
| RN09-03 | Admin não desativa a si mesmo | F03-01 | 1 |
| RN10-01 | Matriz 4 papéis | F03-02 | 2, 3 |
| RN11-01 | Usuário requer ≥1 BU | F03-03 | 3 |
| RN11-02 | Usuário requer ≥1 Módulo | F03-03 | 3 |
| RN11-03 | Efeito imediato | F03-03 | 3 |
| RN12-01 | 403 (não 404) | F03-04 | 4 |
| RN12-02 | Mensagem amigável | F03-04 | 4 |

---

## 🔗 Migrations de Referência

> ⚠️ As migrations da Sprint 4 foram renumeradas para evitar conflitos:
> - **V003** — RLS policies + FORCE ROW LEVEL SECURITY (atualizado na Frente 0)
> - **V004** — Seed resource_action + role_resource (matriz RN10-01) — criado na Frente 0
> - **V006** — FK user_permission.business_unit_id → business_unit.id — criado na Frente 0
>
> A referência anterior a "V003 com seed data" no SPRINT-CARD original estava incorreta (V003 é RLS, não seed).

---

🤖 *Extraído de TEST_PLAN.md. Atualizado em 17/07/2026 pós-Frente 0. Infra de testes pronta (REST Assured, RLS FORCE, seed RBAC). O teste parametrizado da matriz RN10-01 é o centro desta suite.*
