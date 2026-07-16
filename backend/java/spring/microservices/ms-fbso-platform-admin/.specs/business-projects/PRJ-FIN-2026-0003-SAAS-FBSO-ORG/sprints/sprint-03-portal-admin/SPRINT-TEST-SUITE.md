# SPRINT-TEST-SUITE: Sprint 3 — Portal Admin + Contas e Planos

- **Sprint:** 3 de 7
- **Status:** 🔄 Em andamento (iniciada em 16/07/2026)
- **Origem:** [TEST_PLAN.md](../../TEST_PLAN.md) v2.4 §3.1 a §3.8
- **Features:** F01-01 a F02-05 (8 features)
- **Total de cenários:** 55 (extraídos do TEST_PLAN.md) — a executar durante a sprint

---

> 🚫 **Branch:** `feature/java-fbso-platform-admin` ([PRD §8.4](../../PRD.md#84-branch-de-desenvolvimento))

## 1. F01-01: Dashboard Admin — Métricas (7 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F01-01-001 | Dashboard summary carrega indicadores corretos | Unit | §3.1 |
| TC-F01-01-002 | Dashboard evolution com periodo "30d" | Unit | §3.1 |
| TC-F01-01-003 | Periodo invalido assume mes atual | Unit | §3.1 |
| TC-F01-01-004 | GET /dashboard/admin/summary → 200 | Integração | §3.1 |
| TC-F01-01-005 | Soft-deleted tenants excluídos das métricas | Integração | §3.1 |
| TC-F01-01-006 | E2E: Admin visualiza dashboard | E2E | §3.1 |
| TC-F01-01-007 | Sem autenticação → 401 | Segurança | §3.1 |

## 2. F01-02: Lista de Contas (4 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F01-02-001 | Lista paginada page 0, size 25 | Unit | §3.2 |
| TC-F01-02-002 | Busca textual "Mercado" filtra | Unit | §3.2 |
| TC-F01-02-003 | GET /tenants paginado com filtros | Integração | §3.2 |
| TC-F01-02-004 | Busca <3 chars retorna vazio | Integração | §3.2 |

## 3. F01-03: Alertas do Dashboard (5 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F01-03-001 | Onboarding >48h gera alerta | Unit | §3.3 |
| TC-F01-03-002 | Assinatura suspensa gera alerta | Unit | §3.3 |
| TC-F01-03-003 | GET /dashboard/admin/alerts → cards | Integração | §3.3 |
| TC-F01-03-004 | Sem alertas → lista vazia | Integração | §3.3 |
| TC-F01-03-005 | E2E: Cards coloridos no dashboard | E2E | §3.3 |

## 4. F02-01: Criar Tenant (7 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F02-01-001 | Criar tenant → PENDING_ONBOARDING | Unit | §3.4 |
| TC-F02-01-002 | Razão social duplicada → erro | Unit | §3.4 |
| TC-F02-01-003 | Sem name_corporate → erro validação | Unit | §3.4 |
| TC-F02-01-004 | POST /tenants → 201 | Integração | §3.4 |
| TC-F02-01-005 | POST /tenants duplicado → 409 | Integração | §3.4 |
| TC-F02-01-006 | E2E: Admin cria tenant, email enviado | E2E | §3.4 |
| TC-F02-01-007 | Sem autenticação → 401 | Segurança | §3.4 |

## 5. F02-02: Transições de Status (9 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F02-02-001 | PENDING→ACTIVE permitida | Unit | §3.5 |
| TC-F02-02-002 | ACTIVE→SUSPENDED com motivo | Unit | §3.5 |
| TC-F02-02-003 | ACTIVE→PENDING → 422 | Unit | §3.5 |
| TC-F02-02-004 | ACTIVE→INACTIVE permitida | Unit | §3.5 |
| TC-F02-02-005 | POST suspend sem motivo → 400 | Integração | §3.5 |
| TC-F02-02-006 | POST suspend com motivo → 200 | Integração | §3.5 |
| TC-F02-02-007 | Suspensão bloqueia acesso ≤5min | Integração | §3.5 |
| TC-F02-02-008 | E2E: Ciclo ACTIVE→SUSPENDED→ACTIVE | E2E | §3.5 |
| TC-F02-02-009 | OPERATOR tenta suspender → 403 | Segurança | §3.5 |

## 6. F02-03: Planos (7 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F02-03-001 | Criar plano com price>0 | Unit | §3.6 |
| TC-F02-03-002 | price=0 → erro | Unit | §3.6 |
| TC-F02-03-003 | Edição gera nova versão | Unit | §3.6 |
| TC-F02-03-004 | POST /plans → 201 | Integração | §3.6 |
| TC-F02-03-005 | Deactivate com assinantes → 422 | Integração | §3.6 |
| TC-F02-03-006 | E2E: Alterar preço preserva assinaturas | E2E | §3.6 |
| TC-F02-03-007 | OPERATOR tenta criar plano → 403 | Segurança | §3.6 |

## 7. F02-04: Assinaturas (9 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F02-04-001 | Criar assinatura sem ativa → OK | Unit | §3.7 |
| TC-F02-04-002 | Segunda ativa → 409 | Unit | §3.7 |
| TC-F02-04-003 | Change-plan: anterior finalizada, nova criada | Unit | §3.7 |
| TC-F02-04-004 | Suspender assinatura bloqueia módulos | Unit | §3.7 |
| TC-F02-04-005 | POST subscriptions → 201 | Integração | §3.7 |
| TC-F02-04-006 | POST com ativa existente → 409 | Integração | §3.7 |
| TC-F02-04-007 | Change-plan válido → 200 | Integração | §3.7 |
| TC-F02-04-008 | E2E: Assinar, upgrade, histórico | E2E | §3.7 |
| TC-F02-04-009 | Concorrência: 2 assinaturas simultâneas | Integração | §3.7 |

## 8. F02-05: Auditoria (8 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F02-05-001 | Criação de tenant gera auditoria | Unit | §3.8 |
| TC-F02-05-002 | Alteração status gera auditoria | Unit | §3.8 |
| TC-F02-05-003 | Filtro por período funciona | Unit | §3.8 |
| TC-F02-05-004 | GET /audit com filtros → 200 | Integração | §3.8 |
| TC-F02-05-005 | UPDATE em audit_log → erro | Integração | §3.8 |
| TC-F02-05-006 | GET /audit sem paginação usa defaults | Integração | §3.8 |
| TC-F02-05-007 | Tentativa acesso negado → registrada | Segurança | §3.8 |
| TC-F02-05-008 | DELETE em audit_log → 403 | Segurança | §3.8 |

---

## 📊 Resumo

| Nível | Cenários |
|:---|:---:|
| Unit | 24 |
| Integração | 19 |
| E2E | 6 |
| Segurança | 6 |
| **Total** | **55** |

---

## 🔗 Features e RNs Cobertas

| Feature | RNs | Cenários |
|:---|:---|:---:|
| F01-01 | RN01-01, RN01-02, RN01-03 | 7 |
| F01-02 | RN02-01, RN02-02 | 4 |
| F01-03 | RN03-01, RN03-02 | 5 |
| F02-01 | RN04-01, RN04-02, RN04-03 | 7 |
| F02-02 | RN05-01, RN05-02, RN05-03 | 9 |
| F02-03 | RN06-01, RN06-02, RN06-03 | 7 |
| F02-04 | RN07-01, RN07-02, RN07-03 | 9 |
| F02-05 | RN08-01, RN08-02 | 8 |

---

🤖 *Extraído de TEST_PLAN.md v2.4. Sprint 3 iniciada em 16/07/2026. Execute estes 55 cenários antes de considerar a Sprint 3 concluída.*
