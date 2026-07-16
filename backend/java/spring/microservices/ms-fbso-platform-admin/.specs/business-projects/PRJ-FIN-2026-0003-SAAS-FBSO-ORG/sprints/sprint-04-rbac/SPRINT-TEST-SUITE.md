# SPRINT-TEST-SUITE: Sprint 4 — Governança de Acessos (RBAC)

- **Sprint:** 4 de 7
- **Origem:** [TEST_PLAN.md](../../TEST_PLAN.md) §3.9 a §3.12 + §4.1
- **Features:** F03-01 a F03-04 (4 features)
- **Total de cenários:** 19

---

## 1. F03-01: Gestão de Usuários (7 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F03-01-001 | Convidar usuario email unico → OK | Unit | §3.9 |
| TC-F03-01-002 | Email duplicado → 409 | Unit | §3.9 |
| TC-F03-01-003 | Admin nao pode desativar a si mesmo | Unit | §3.9 |
| TC-F03-01-004 | POST /users → 201 | Integração | §3.9 |
| TC-F03-01-005 | Auto-desativacao → 422 | Integração | §3.9 |
| TC-F03-01-006 | E2E: Convidar → email → desativar | E2E | §3.9 |
| TC-F03-01-007 | OPERATOR tenta POST /users → 403 | Segurança | §3.9 |

## 2. F03-02: Matriz de Permissões RBAC (9 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F03-02-001 | ADMIN_TENANT permite qualquer ação | Unit | §3.10 |
| TC-F03-02-002 | AUDITOR permite apenas leitura | Unit | §3.10 |
| TC-F03-02-003 | **Teste parametrizado: cada papel × cada endpoint** | Segurança | §3.10 |
| TC-F03-02-004 | OPERATOR tenta PATCH /products → 403 | Segurança | §3.10 |
| TC-F03-02-005 | AUDITOR tenta POST /tenants → 403 | Segurança | §3.10 |
| TC-F03-02-006 | MANAGER pode criar BU | Integração | §3.10 |
| TC-F03-02-007 | ADMIN_TENANT pode criar BU | Integração | §3.10 |
| TC-F03-02-008 | E2E: Login cada papel, verifica menu | E2E | §3.10 |
| TC-F03-02-009 | Admin FBSO ve todos tenants (cross-tenant) | Segurança | §3.10 |

## 3. F03-03: Vinculação Usuário × Unidade × Módulo (6 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F03-03-001 | Atribuir permissao usuario x BU x role | Unit | §3.11 |
| TC-F03-03-002 | Admin acesso implícito todas BUs | Unit | §3.11 |
| TC-F03-03-003 | Usuario sem vinculacao → 403 | Integração | §3.11 |
| TC-F03-03-004 | Usuario BU-1 ve apenas produtos BU-1 | Integração | §3.11 |
| TC-F03-03-005 | Usuario BU-1 tenta acessar produto BU-2 por ID → 404 | Segurança | §3.11 |
| TC-F03-03-006 | Alterar permissao → efeito imediato | Segurança | §3.11 |

## 4. F03-04: Acesso Condicional — 403 Amigável (5 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F03-04-001 | 403 sem detalhes técnicos | Unit | §3.12 |
| TC-F03-04-002 | 403 segue RFC 7807 (type, title, status, detail) | Integração | §3.12 |
| TC-F03-04-003 | E2E: URL proibida → tela 403 amigavel | E2E | §3.12 |
| TC-F03-04-004 | JWT adulterado (elevacao privilegio) → 401 | Segurança | §3.12 |
| TC-F03-04-005 | URL proibida → 403 (nao 404 — RN12-01) | Segurança | §3.12 |

---

## 5. Testes RBAC Parametrizados (§4.1 do TEST_PLAN)

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

## 📊 Resumo

| Nível | Cenários |
|:---|:---:|
| Unit | 7 |
| Integração | 8 |
| E2E | 3 |
| Segurança | 9 |
| **Total** | **27** |

---

## 🔗 RNs Cobertas

| RN | Descrição | Cenários |
|:---|:---|:---:|
| RN09-01 | Convite expira 7 dias | F03-01 |
| RN09-02 | Email único por tenant | F03-01 |
| RN09-03 | Admin não desativa a si mesmo | F03-01 |
| RN10-01 | Matriz 4 papéis | F03-02 |
| RN11-01 | Usuário requer ≥1 BU | F03-03 |
| RN11-02 | Usuário requer ≥1 Módulo | F03-03 |
| RN11-03 | Efeito imediato | F03-03 |
| RN12-01 | 403 (não 404) | F03-04 |
| RN12-02 | Mensagem amigável | F03-04 |

---

🤖 *Extraído de TEST_PLAN.md v2.0. O teste parametrizado da matriz RN10-01 é o centro desta suite.*
