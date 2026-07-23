# SPRINT-TEST-PLANNING-Frente-1.md — Plano de Testes: Sprint 6 — Frente 1

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 6 de 7 — sprint-06-bus-catalogo
- **Frente:** Frente 1 — Recomendados
- **Stack:** Java 25 · Spring Boot 3.5.14 · PostgreSQL 17 · Flyway 12.11.0 · JUnit 5 + Mockito
- **Data:** 23 de Julho de 2026

---

## 1. Visão Geral

- **Tasks implementadas:** 3 (T-165, T-168, T-169) + 2 já concluídas (T-166, T-167)
- **Cenários de teste mapeados:** 12
- **Meta de cobertura:** ≥ 80% (padrão JaCoCo)
- **Ferramentas:** JUnit 5 + Mockito (unit), Testcontainers + PostgreSQL 17 (integration)

---

## 2. Mapeamento Task → Cenários de Teste

| Task | ID | Descrição | Nível | Ferramenta | Status |
|:---|:---|:---|:---|:---|:---:|
| **T-165** | TC-F1-165-001 | V009 migration aplica sem erro — tenant_id populado | Integração | Flyway + PostgreSQL | ⬜ |
| **T-165** | TC-F1-165-002 | U009 rollback reverte — coluna tenant_id removida | Integração | Flyway + PostgreSQL | ⬜ |
| **T-165** | TC-F1-165-003 | ProductService entity — tenantId no toColumnMap() | Unit | JUnit 5 | ⬜ |
| **T-168** | TC-F1-168-001 | findTree() retorna árvore com 3 níveis — ordem correta | Unit | Mockito + JdbcTemplate | ⬜ |
| **T-168** | TC-F1-168-002 | findTree() com tenant sem BUs → lista vazia | Unit | Mockito + JdbcTemplate | ⬜ |
| **T-168** | TC-F1-168-003 | findChildren() retorna filhas diretas de uma BU | Unit | Mockito + JdbcTemplate | ⬜ |
| **T-168** | TC-F1-168-004 | RowMapper mapeia todos os 15 campos + is_matrix | Unit | JUnit 5 | ⬜ |
| **T-168** | TC-F1-168-005 | AuditorsFields mapeados via AuditFieldsRowMapper | Unit | JUnit 5 | ⬜ |
| **T-169** | TC-F1-169-001 | extractKey() usa X-Forwarded-For de proxy confiável | Unit | Mockito + MockHttpServletRequest | ⬜ |
| **T-169** | TC-F1-169-002 | extractKey() ignora X-Forwarded-For de IP não confiável | Unit | Mockito + MockHttpServletRequest | ⬜ |
| **T-169** | TC-F1-169-003 | extractKey() usa remoteAddr quando sem X-Forwarded-For | Unit | Mockito + MockHttpServletRequest | ⬜ |
| **T-169** | TC-F1-169-004 | Construtor aceita null trustedProxyIps → lista vazia | Unit | JUnit 5 | ⬜ |

---

## 3. Estratégia por Nível de Teste

### 3.1 Testes Unitários

- **Ferramenta:** JUnit 5 + Mockito
- **Padrão:** AAA (Arrange-Act-Assert)
- **Localização:** `src/test/java/.../unit/`
- **O que mockar:** JdbcTemplate (para BusinessUnitRepository), HttpServletRequest (para RateLimitFilter)
- **O que NÃO mockar:** BusinessUnitRowMapper, BusinessUnit entity, ProductService entity

### 3.2 Testes de Integração

- **Ferramenta:** Flyway migration test (via `mvn flyway:migrate` + `mvn flyway:undo`)
- **Localização:** migrations são testadas estruturalmente — `mvn flyway:migrate` sem erro confirma V009
- **Dados de seed:** PostgreSQL com business_unit existente para testar populamento de tenant_id

---

## 4. Ordem de Execução dos Testes

1. **Testes unitários** (sem dependências externas — rodam primeiro)
   - BusinessUnitRowMapperTest (TC-F1-168-004, TC-F1-168-005)
   - RateLimitFilterTest (TC-F1-169-001 a TC-F1-169-004)
   - BusinessUnitRepositoryTest (TC-F1-168-001 a TC-F1-168-003)
   - ProductServiceTest (TC-F1-165-003)

2. **Testes de integração** (dependem de container/DB)
   - V009 Migration (TC-F1-165-001, TC-F1-165-002) — validado via `mvn flyway:migrate`

---

## 5. Comandos de Execução

- Unit: `./mvnw test`
- Integration: `./mvnw verify` (requer Docker para Testcontainers)
- Coverage: `./mvnw jacoco:check`
- Lint/Quality: `./mvnw checkstyle:check`

---

## 6. Ações Manuais ou Externas

> Nenhuma ação manual necessária. Todos os testes são automatizados.

---

## 7. Provenientes de Testes de Validação de Qualidade

> A preencher após Fase 5, se necessário.

| Task | Mensagem exata | Suspeita | Proposta solução |
|:-----|:---------------|:---------|:----------------:|
| — | — | — | — |

---

## 8. Provenientes de Code Review

> A preencher após Fase 7 (Code Review), se necessário.

| Task | Mensagem exata | Suspeita | Proposta solução | Skills |
|:-----|:---------------|:---------|:-----------------|-------:|
| — | — | — | — | — |

---

## Rodapé

🤖 *Documento gerado em 23/07/2026 conforme PROMPT-EXECUTE-SPRINT-TASKS.md Fase 3. Stack: JUnit 5 + Mockito + Testcontainers. 12 cenários mapeados para 3 tasks.*
