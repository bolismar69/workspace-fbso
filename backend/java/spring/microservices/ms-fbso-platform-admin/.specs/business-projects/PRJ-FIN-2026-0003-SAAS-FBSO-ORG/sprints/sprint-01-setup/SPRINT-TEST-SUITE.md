# SPRINT-TEST-SUITE: Sprint 1 — Setup e Fundação

- **Sprint:** 1 de 7
- **Origem:** [TEST_PLAN.md](../../TEST_PLAN.md) §9 — Testes de Infraestrutura e Build
- **Total de cenários:** 17

> ⚠️ A Sprint 1 não implementa features de negócio — é fundação. Os testes aqui são estruturais: compilação, migrations, BaseRepository e validação de enums. Todos os cenários referenciam [TEST_PLAN.md §9](../../TEST_PLAN.md).

> ⚠️ Correção pós-gate: 3 NCs resolvidas do SPRINT_ARTEFACTS_FAIL_REPORT.md v1.0. NC-003: IDs renomeados de TC-S1-* para TC-INFRA-* (alinhados ao TEST_PLAN.md §9). NC-004: rastreabilidade restaurada com referências a §9. NC-005: contagem corrigida de 12→17. Data: 14/07/2026.

---

> 🚫 **Branch:** `feature/java-fbso-platform-admin` ([PRD §8.4](../../PRD.md#84-branch-de-desenvolvimento))

## 1. Testes de Compilação e Estrutura

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-INFRA-001 | `mvn clean compile` executa sem erros | Build | §9.1 |
| TC-INFRA-002 | `mvn clean install` executa sem erros | Build | §9.1 |
| TC-INFRA-003 | `mvn test` executa suite vazia (placeholder) | Build | §9.1 |
| TC-INFRA-004 | Estrutura de pacotes confere com ARCHITECTURE.md §2 | Estrutural | §9.1 |

---

## 2. Testes de Migrations Flyway

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-INFRA-005 | `mvn flyway:migrate` cria schema `fbso_platform` | Integração | §9.2 |
| TC-INFRA-006 | V001 cria 11 tabelas com colunas, PKs, FKs, NOT NULL | Integração | §9.2 |
| TC-INFRA-007 | V001 rollback desfaz migração (se configurado) | Integração | §9.2 |
| TC-INFRA-008 | V002 cria índices únicos parciais (`WHERE deleted_dt IS NULL`) | Integração | §9.2 |
| TC-INFRA-009 | Colunas de auditoria presentes em TODAS as tabelas | Integração | §9.2 |

---

## 3. Testes de BaseRepository

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-INFRA-010 | `findAll()` injeta `WHERE deleted_dt IS NULL` automaticamente | Unit | §9.3 |
| TC-INFRA-011 | `softDelete()` seta `deleted_dt = NOW()` e `deleted_by` | Unit | §9.3 |
| TC-INFRA-012 | `save()` preenche `created_by`/`updated_by` automaticamente | Unit | §9.3 |

---

## 4. Validação de Enums

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-INFRA-013 | Todos os 8 enums compilam com valores corretos | Unit | §9.4 |
| TC-INFRA-014 | Enum `TenantStatus` contém: PENDING_ONBOARDING, ACTIVE, SUSPENDED, INACTIVE | Unit | §9.4 |
| TC-INFRA-015 | Enum `Role` contém: ADMIN_TENANT, MANAGER_BU, OPERATOR_BU, AUDITOR | Unit | §9.4 |

---

## 5. Configuração de Profiles

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-INFRA-016 | Profile `dev` carrega `application-dev.yml` | Integração | §9.5 |
| TC-INFRA-017 | Profile `prod` carrega `application-prod.yml` | Integração | §9.5 |

---

## 📊 Resumo

| Nível | Cenários |
|:---|:---:|
| Build | 4 |
| Integração | 7 |
| Unit | 6 |
| **Total** | **17** |

---

## 🔗 Referência Rápida

| Documento | Seção Relevante |
|:---|:---|
| TEST_PLAN.md | §9 — Testes de Infraestrutura e Build |
| TASKS.md | §2 — Pre-M2 Setup (T-001 a T-008) |
| ARCHITECTURE.md | §2 — Estrutura de Pacotes, §5 — Design de Persistência |
| SPECS.md | §6.3 — Campos de Auditoria, §6.1 — Entidades |

---

🤖 *Extraído de TEST_PLAN.md v2.3. Revisão Caveman em 15/07/2026 (DOCS-SERVICE-SPRINTS-CAVEMAN-REVIEW.md): TEST_PLAN ref v2.1→v2.3 §9. Execute estes cenários antes de considerar a Sprint 1 concluída.*
