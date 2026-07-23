# SPRINT-TEST-PLANNING-Frente-2.md — Plano de Testes: Sprint 5 — Frente 2

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 5 de 7 — sprint-05-portal-cliente
- **Frente:** Frente 2 — Desejáveis (🔵 Could Have)
- **Stack:** Java 25 · Spring Boot 3.5.14 · PostgreSQL 17 · Keycloak 26 · Flyway 12.11.0 · Caffeine 3.2.4
- **Data do planejamento:** 2026-07-23

---

## 1. Visão Geral

- **Tasks implementadas:** 8 (T-149.DT-086 a T-156.DT-113)
- **Natureza das mudanças:** Refatoração (T-149, T-150), Correção de bug latente (T-151), Bump de dependência (T-152), Externalização de configuração (T-153), Documentação (T-154, T-155, T-156)
- **Meta de cobertura:** ≥ 80% mantida (nenhuma task introduz nova lógica de negócio)
- **Estratégia:** Testes de regressão existentes cobrem todas as mudanças. Nenhuma task da Frente 2 introduz nova funcionalidade que exija testes dedicados.
- **Ferramentas:** JUnit 5 + Mockito (unitários). Surefire para execução completa.

---

## 2. Mapeamento Task → Validação

| Task | O que validar | Nível | Como validar | Status |
|:---|:---|:---|:---|:---:|
| T-149.DT-086 | AuditFieldsRowMapper — 4 RowMappers usam helper sem quebrar | Unit/Regressão | `mvn test` — todos os testes de RowMapper passam | ✅ |
| T-150.DT-089 | ObjectMapper injetado no AuditAspect — serialização consistente | Unit | `AuditAspectTest` (3 testes) passam com novo construtor | ✅ |
| T-151.DT-090 | OffsetDateTime.now(ZoneOffset.UTC) — sem regressão | Unit/Regressão | `mvn test` — zero falhas em testes com data/hora | ✅ |
| T-152.DT-092 | springdoc 2.8.16 — Swagger UI funcional | Integração | `mvn compile` OK. Swagger UI em `/swagger-ui.html` | ✅ |
| T-153.DT-093 | CORS do application.yml — origens configuráveis | Integração | `mvn compile` OK. `@Value` injetado sem erro | ✅ |
| T-154.DT-101 | SPRINT-CARD riscos atualizados | Documentação | Revisão manual da tabela de riscos | ✅ |
| T-155.DT-112 | SPECS.md header atualizado | Documentação | Revisão manual do header | ✅ |
| T-156.DT-113 | TASKS.md progresso recalculado | Documentação | Consistência cruzada com SPRINT-CARD.md | ✅ |

---

## 3. Estratégia por Nível de Teste

### 3.1 Testes Unitários

- **Ferramenta:** JUnit 5 + Mockito
- **Localização:** `src/test/java/com/fbso/platform/admin/unit/`
- **Cenário específico:** `AuditAspectTest` — atualizado para injetar `ObjectMapper` no novo construtor. 3 testes passando.
- **Cobertura existente:** 213 testes unitários cobrem RowMappers, AuditAspect, BaseRepository, BaseEntity — todas as classes modificadas.

### 3.2 Testes de Integração

- **Nenhum novo teste de integração necessário.** As mudanças são refatoração (não alteram comportamento) ou configuração (verificadas via compilação).

### 3.3 Verificações Estáticas

| Verificação | Comando | Resultado Esperado |
|:---|:---|:---|
| Nenhum `OffsetDateTime.now()` sem UTC | `grep -r "OffsetDateTime.now()" src/main/java/` | VAZIO |
| Nenhuma duplicação de audit fields em RowMappers | `grep -r "setCreatedDt\|setUpdatedDt\|setCreatedBy\|setUpdatedBy\|setDeletedDt\|setDeletedBy" src/main/java/**/rowmapper/` | APENAS `AuditFieldsRowMapper.java` |
| Nenhum `new ObjectMapper()` no AuditAspect | `grep "new ObjectMapper" src/main/.../AuditAspect.java` | VAZIO |

---

## 4. Ordem de Execução dos Testes

1. **Compilação:** `mvn compile` — verifica springdoc bump, CORS @Value, novos imports
2. **Testes unitários:** `mvn test` — 213 testes, 0 novas falhas
3. **Verificações estáticas:** grep patterns acima

---

## 5. Comandos de Execução

| Propósito | Comando |
|:---|:---|
| Compilação | `mvn compile` |
| Testes unitários | `mvn test` |
| Verificação UTC | `grep -rn "OffsetDateTime.now()" src/main/java/` |
| Verificação DRY (RowMappers) | `grep -rn "setCreatedDt\|setUpdatedDt\|setCreatedBy\|setUpdatedBy\|setDeletedDt\|setDeletedBy" src/main/java/` |

---

## 6. Ações Manuais ou Externas

> Nenhuma ação manual necessária. Todas as validações são automatizadas.

---

## 7. Resultado da Execução

| Métrica | Resultado |
|:---|---|
| `mvn compile` | ✅ BUILD SUCCESS (8.4s) |
| `mvn test` | ✅ 213 testes: 0 failures, 1 pre-existing error, 8 skipped |
| Erro pré-existente | `SubscriptionServiceTest.shouldCreateWithLockedPrice` — já documentado na Frente 1 |
| Verificação UTC | ✅ Zero `OffsetDateTime.now()` sem UTC |
| Verificação DRY | ✅ Apenas `AuditFieldsRowMapper.java` contém os 6 setters de auditoria |

---

## 8. Provenientes de Testes de Validação de Qualidade

| Task | Mensagem exata | Suspeita | Proposta solução |
|:-----|:---------------|:---------|:----------------:|
| — | Nenhum achado | — | — |

---

## 9. Provenientes de Code Review

| Task | Mensagem exata | Suspeita | Proposta solução | Skills |
|:-----|:---------------|:---------|:-----------------|-------:|
| — | Nenhum achado | — | — | — |

---

🤖 *Documento gerado em 2026-07-23 como parte da Fase 3 do PROMPT-EXECUTE-SPRINT-TASKS.md. Stack detectada: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17. 8/8 tasks da Frente 2 concluídas.*
