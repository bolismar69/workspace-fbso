# SPRINT-TEST-SUITE: Sprint 6 — Unidades de Negócio e Catálogo

- **Sprint:** 6 de 7
- **Origem:** [TEST_PLAN.md](../../TEST_PLAN.md) §3.17 a §3.18 + §4.2
- **Features:** F04-05 a F04-06 (2 features)
- **Total de cenários:** 17

---

## 1. F04-05: Unidades de Negócio (9 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F04-05-001 | Criar BU CNPJ válido único → OK | Unit | §3.17 |
| TC-F04-05-002 | CNPJ duplicado → 409 | Unit | §3.17 |
| TC-F04-05-003 | Soft delete libera CNPJ para reúso | Unit | §3.17 |
| TC-F04-05-004 | Criar filha com parent_id válido → hierarquia | Unit | §3.17 |
| TC-F04-05-005 | POST /business-units → 201 | Integração | §3.17 |
| TC-F04-05-006 | CNPJ inválido → 400 | Integração | §3.17 |
| TC-F04-05-007 | POST deactivate → soft delete | Integração | §3.17 |
| TC-F04-05-008 | Tree de BUs → estrutura aninhada | Integração | §3.17 |
| TC-F04-05-009 | E2E: Criar Matriz + filiais, desativar, recriar com mesmo CNPJ | E2E | §3.17 |

## 2. F04-06: Catálogo de Produtos (7 cenários)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F04-06-001 | Criar produto vinculado BU → OK | Unit | §3.18 |
| TC-F04-06-002 | SKU duplicado mesma BU → erro | Unit | §3.18 |
| TC-F04-06-003 | Soft delete produto → não remove registro | Unit | §3.18 |
| TC-F04-06-004 | POST /products → 201 | Integração | §3.18 |
| TC-F04-06-005 | GET /products filtra por BU | Integração | §3.18 |
| TC-F04-06-006 | E2E: Criar, listar, desativar produto | E2E | §3.18 |
| TC-F04-06-007 | OPERATOR tenta criar produto → 403 | Segurança | §3.18 |

---

## 3. Isolamento Multi-Tenant (4 cenários — TEST_PLAN §4.2)

| ID | Descrição | Nível | Critério |
|:---|:---|:---|:---|
| TC-SEC-MT-001 | Query cross-tenant retorna vazio | Integração | TenantContext(tenant-A) → apenas dados tenant-A |
| TC-SEC-MT-002 | Acessar recurso outro tenant por ID → 404 | Segurança | GET /business-units/{id-tenant-B} como tenant-A → 404 |
| TC-SEC-MT-003 | JWT adulterado com tenant_id outro tenant | Segurança | Assinatura inválida → 401 |
| TC-SEC-MT-004 | Admin FBSO cross-tenant vê todos | Segurança | Admin FBSO sem tenant_id → visão global |

---

## 4. Teste de Isolamento Específico da Sprint (T-070)

| ID | Descrição | Nível | Critério |
|:---|:---|:---|:---|
| TC-S6-MT-001 | Tenant-A não vê BUs de Tenant-B | Integração | GET /business-units como tenant-A → 0 BUs de tenant-B |
| TC-S6-MT-002 | Tenant-A não vê produtos de Tenant-B | Integração | GET /products como tenant-A → 0 produtos de tenant-B |
| TC-S6-MT-003 | Tenant-A tenta POST product em BU de Tenant-B → 404 | Segurança | business_unit_id de tenant-B → 404 |

---

## 📊 Resumo

| Nível | Cenários |
|:---|:---:|
| Unit | 7 |
| Integração | 9 |
| E2E | 2 |
| Segurança | 5 |
| **Total** | **23** |

---

## 🔗 RNs Cobertas

| RN | Descrição | Feature |
|:---|:---|:---|
| RN17-01 | CNPJ único entre ativos, soft delete libera reúso | F04-05 |
| RN17-02 | BU desativada não pode ser "pai" | F04-05 |
| RN17-04 | Sem limite de níveis hierárquicos | F04-05 |
| RN17-05 | Seletor BU reflete permissões | F04-05 |
| RN18-01 | Catálogo segmentado por BU | F04-06 |
| RN18-02 | SKU opcional, único por BU ativo | F04-06 |
| RN18-03 | Indicador "Não mapeado" | F04-06 |
| RN18-04 | Soft delete em produtos | F04-06 |

---

🤖 *Extraído de TEST_PLAN.md v2.0. O isolamento multi-tenant (T-070) é o diferencial desta sprint.*
