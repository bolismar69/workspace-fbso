# SPRINT-TEST-SUITE: Sprint 5 — Portal do Cliente e Onboarding

- **Sprint:** 5 de 7
- **Origem:** [TEST_PLAN.md](../../TEST_PLAN.md) §3.13 a §3.16
- **Features:** F04-01 a F04-04 (4 features)
- **Total de cenários:** 33 (11 F04-01 + 11 F04-02 + 7 F04-03 + 4 F04-04)

---


## 1. F04-01: Login e Autenticação (11 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F04-01-001 | JWT valido estabelece TenantContext | Unit | §3.13 |
| TC-F04-01-002 | JWT expirado → 401 | Unit | §3.13 |
| TC-F04-01-003 | JWT assinatura inválida → 401 | Integração | §3.13 |
| TC-F04-01-004 | Bloqueio após 5 tentativas incorretas | Integração | §3.13 |
| TC-F04-01-005 | E2E: Fluxo login (sucesso + falha) | E2E | §3.13 |
| TC-F04-01-006 | E2E: Recuperação senha → link 1h | E2E | §3.13 |
| TC-F04-01-007 | JWT sem tenant_id → 401 | Segurança | §3.13 |
| TC-F04-01-008 | JWT com role inexistente → 401 | Segurança | §3.13 |
| TC-F04-01-009 | Brute force: rate limiting 429 | Segurança | §3.13 |
| **TC-F04-01-010** 🆕 | **Sessão 60min inatividade → 401 (RN13-02)** | **Integração** | §3.13 |
| **TC-F04-01-011** 🆕 | **Validação complexidade senha: 7 chars, 8 s/ número, 8 válida (RN13-01)** | **Unit** | §3.13 |

## 2. F04-02: Onboarding Guiado (11 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F04-02-001 | Passo 1 confirma dados | Unit | §3.14 |
| TC-F04-02-002 | Passo 2 cria BU Matriz (parent_id=null) | Unit | §3.14 |
| TC-F04-02-003 | Complete → tenant ACTIVE | Unit | §3.14 |
| TC-F04-02-004 | Pular passo → redireciona (RN14-01) | Unit | §3.14 |
| TC-F04-02-005 | GET /onboarding/status → progresso | Integração | §3.14 |
| TC-F04-02-006 | PATCH /step-1 → 200 | Integração | §3.14 |
| TC-F04-02-007 | POST /step-2 cria BU Matriz | Integração | §3.14 |
| TC-F04-02-008 | POST /complete antes dos passos → 422 | Integração | §3.14 |
| TC-F04-02-009 | E2E: Fluxo completo 4 passos | E2E | §3.14 |
| TC-F04-02-010 | Burlar onboarding via API → 403 | Segurança | §3.14 |
| **TC-F04-02-011** 🆕 | **Passo 3: sucesso, falha (falta CNPJ step 2), retomada pós-interrupção** | **Integração** | §3.14 |

## 3. F04-03: Dashboard do Cliente (7 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F04-03-001 | Cards: unidades, produtos, plano | Unit | §3.15 |
| TC-F04-03-002 | Notificações com link de ação | Unit | §3.15 |
| TC-F04-03-003 | GET /dashboard/client/summary → 200 | Integração | §3.15 |
| TC-F04-03-004 | Isolamento: dashboard não vaza dados outro tenant | Integração | §3.15 |
| TC-F04-03-005 | E2E: Cliente vê dashboard pós-onboarding | E2E | §3.15 |
| **TC-F04-03-006** 🆕 | **Segurança: cliente tenant A não acessa dashboard tenant B via troca de tenant_id** | **Segurança** | §3.15 |
| **TC-F04-03-007** 🆕 | **Segurança: GET /dashboard/client/summary sem token → 401** | **Segurança** | §3.15 |

## 4. F04-04: App Switcher (4 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F04-04-001 | /auth/me retorna dados + modules[] | Unit | §3.16 |
| TC-F04-04-002 | GET /auth/me → modules do plano | Integração | §3.16 |
| TC-F04-04-003 | E2E: App Switcher exibe módulos | E2E | §3.16 |
| TC-F04-04-004 | JWT sem modules[] → 403 | Segurança | §3.16 |

---

## 📊 Resumo

| Nível | Cenários |
|:---|:---:|
| Unit | 9 |
| Integração | 13 |
| E2E | 5 |
| Segurança | 6 |
| **Total** | **33** |

---

## 🔗 RNs Cobertas

| RN | Descrição | Feature |
|:---|:---|:---|
| RN13-01 | Senha 8+ chars, letra+número | F04-01 |
| RN13-02 | Sessão 60min inatividade / bloqueio 15min | F04-01 |
| RN13-03 | Link reset único, expira 1h | F04-01 |
| RN14-01 | Onboarding obrigatório, ordem fixa | F04-02 |
| RN14-02 | Primeira BU = Matriz | F04-02 |
| RN14-03 | Complete só se todos passos OK | F04-02 |
| RN14-04 | Tenant → ACTIVE após onboarding | F04-02 |
| RN15-01 | Dashboard adapta ao módulo | F04-03 |
| RN15-02 | Fase 0: dashboard genérico | F04-03 |
| RN16-01 | App Switcher: interseção plano×perms | F04-04 |
| RN16-02 | Placeholder "FBSO Platform" | F04-04 |

---

🤖 *Extraído de TEST_PLAN.md v2.0. O fluxo de onboarding completo é o cenário E2E mais importante desta sprint.*
