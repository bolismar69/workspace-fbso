# SPRINT-TEST-SUITE: Sprint 6 — Unidades de Negócio e Catálogo

- **Sprint:** 6 de 7
- **Origem:** [TEST_PLAN.md](../../TEST_PLAN.md) §3.17 a §3.18 + §4.2
- **Features:** F04-05 a F04-06 (2 features)
- **Status:** 🔄 Em Execução — Frente 0 concluída ✅ (4/4). 261 testes totais (0 failures, +34 CnpjValidator). Code review: 7 skills, 12 findings, HIGH-1 corrigido.
- **Total de cenários:** 17 (features) + 4 (multi-tenant) + 3 (isolamento sprint) + 45 (CnpjValidator F0) = 69 total

---


## 0. Frente 0 — Testes Concluídos ✅ | 23/07/2026

| ID | Descrição | Nível | Resultado |
|:---|:---|:---|:---:|
| TC-F0-CNPJ-001 | CNPJs numéricos válidos (5) | Unit | ✅ 5/5 |
| TC-F0-CNPJ-002 | CNPJs alfanuméricos válidos (11 — IN RFB 2.119/2022) | Unit | ✅ 11/11 |
| TC-F0-CNPJ-003 | CNPJs numéricos inválidos (6) | Unit | ✅ 6/6 |
| TC-F0-CNPJ-004 | CNPJs alfanuméricos inválidos (5) | Unit | ✅ 5/5 |
| TC-F0-CNPJ-005 | Entradas nulas/vazias (5) | Unit | ✅ 5/5 |
| TC-F0-CNPJ-006 | calculaDV — pesos e dígitos (6) | Unit | ✅ 6/6 |
| TC-F0-CNPJ-007 | strip — remoção de formatação (3) | Unit | ✅ 3/3 |
| TC-F0-CNPJ-008 | format — formatação com máscara (4) | Unit | ✅ 4/4 |

**CnpjValidatorTest:** 45/45 ✅. Algoritmo unificado `(int) char - (int) '0'` suporta numérico e alfanumérico (IN RFB 2.119/2022, Jul/2026).

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

## 4. Teste de Isolamento Específico da Sprint (T-077)

| ID | Descrição | Nível | Critério |
|:---|:---|:---|:---|
| TC-S6-MT-001 | Tenant-A não vê BUs de Tenant-B | Integração | GET /business-units como tenant-A → 0 BUs de tenant-B |
| TC-S6-MT-002 | Tenant-A não vê produtos de Tenant-B | Integração | GET /products como tenant-A → 0 produtos de tenant-B |
| TC-S6-MT-003 | Tenant-A tenta POST product em BU de Tenant-B → 404 | Segurança | business_unit_id de tenant-B → 404 |

---

## 📊 Resumo

| Nível | Frente 0 ✅ | Features ⬜ | Total |
|:---|:---:|:---:|:---:|
| Unit | 45 (CnpjValidator) | 7 | 52 |
| Integração | — | 9 | 9 |
| E2E | — | 2 | 2 |
| Segurança | — | 5 | 5 |
| **Total** | **45** | **23** | **68** |

**Testes totais no projeto:** 261 (0 failures, 1 pre-existing error DT-136)

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

🤖 *Extraído de TEST_PLAN.md v3.3. Frente 0 concluída em 23/07/2026: 45/45 CnpjValidator. O isolamento multi-tenant (T-077) é o diferencial desta sprint.*
