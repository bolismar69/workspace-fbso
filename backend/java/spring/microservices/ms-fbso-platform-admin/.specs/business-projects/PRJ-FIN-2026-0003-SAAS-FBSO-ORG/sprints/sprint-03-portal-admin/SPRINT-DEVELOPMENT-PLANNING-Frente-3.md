# SPRINT-DEVELOPMENT-PLANNING-Frente-3.md — Plano de Desenvolvimento: Sprint 3 — Frente 3

- **Solução:** `ms-fbso-platform-admin`
- **Sprint:** 3 de 7 — Portal Admin + Contas e Planos
- **Frente:** 3 — Correções Durante a Sprint (7 Débitos Técnicos Não-Bloqueantes)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + JDBC Template + Flyway + Testcontainers
- **Tasks:** T-039.DT-017 a T-045.DT-046 (7)
- **Estimativa total:** ~10h
- **Data:** 17/07/2026
- **Status:** 📋 Planejamento — aguardando início da implementação

---

## 1. Visão Geral

- **Objetivo:** Corrigir 7 débitos técnicos não-bloqueantes de alto impacto identificados na auditoria multidisciplinar.
- **Contexto:** Frentes 0+1+2 já concluídas (35/42 tasks). Esta é a última frente da Sprint 3.
- **Ordem:** Sequencial por complexidade crescente (mais simples primeiro, mais complexo por último).
- **Dependência externa:** Docker (para T-043).

---

## 2. Plano por Task

### T-045.DT-046 — Atualizar Testcontainers 1.20.6→1.21.4

- **Critério DONE:** CVE-2024-25710 mitigada. Build passa.
- **Estimativa:** 15min
- **Abordagem:**
  - Alterar `<testcontainers.version>` de 1.20.6 para 1.21.4 no pom.xml
  - `commons-compress` é dependência transitiva — versão 1.28.0+ já incluída no Testcontainers 1.21.4
  - Se o build falhar com conflito de versão transitiva, adicionar `<dependencyManagement>` para commons-compress 1.28.0
- **Arquivos a modificar:** `pom.xml`
- **Riscos:** Mudança de major version pode quebrar APIs do Testcontainers. Verificar compilação.
- **Skills:** 110-java-maven-best-practices

### T-042.DT-025 — Adicionar @ExceptionHandler(AccessDeniedException.class)

- **Critério DONE:** Acesso negado retorna 403 (não 500)
- **Estimativa:** 10min
- **Abordagem:**
  - Adicionar handler no `GlobalExceptionHandler` para `org.springframework.security.access.AccessDeniedException`
  - Handler retorna HTTP 403 com corpo RFC 7807 igual ao `PermissionDeniedException`
  - O handler existente para `java.lang.SecurityException` é mantido (cobre TenantIsolation)
- **Arquivos a modificar:** `exception/GlobalExceptionHandler.java`
- **Skills:** 126-java-exception-handling

### T-039.DT-017 — Decidir V004 "opcional" vs "pré-requisito"

- **Critério DONE:** Decisão registrada no Log de Decisões
- **Estimativa:** 15min
- **Abordagem:**
  - Confirmado: `idx_tenant_segment` NÃO existe na V002
  - O campo `segment` não é usado em WHERE de queries frequentes
  - Decisão: V004 é **OPCIONAL** — não bloqueia desenvolvimento. Postergar para Sprint 4 se necessário
  - Registrar decisão no SPRINT-DEVELOPMENT-PLANNING-DRAFT.md §12
- **Arquivos a modificar:** `SPRINT-DEVELOPMENT-PLANNING-DRAFT.md` (documentação)
- **Skills:** N/A (documentação)

### T-040.DT-019 — Recalibrar day-by-day: 12→~15 dias

- **Critério DONE:** Planejamento atualizado no SPRINT-DEVELOPMENT-PLANNING
- **Estimativa:** 15min
- **Abordagem:**
  - Atualizar estimativas no SPRINT-DEVELOPMENT-PLANNING-DRAFT.md
  - Documentar que Frentes 0+1+2 foram concluídas em ~2.5 dias (antecipado)
  - Ajustar cronograma: 15 dias úteis realista com 1 dev
- **Arquivos a modificar:** `SPRINT-DEVELOPMENT-PLANNING-DRAFT.md` (documentação)
- **Skills:** N/A (documentação)

### T-044.DT-029 — Extrair hasTenantColumn branching para helper

- **Critério DONE:** 4 métodos → 1 helper. Sem duplicação.
- **Estimativa:** 1h
- **Abordagem:**
  - `tenantClause()` e `buildParams()` já existem como helpers privados
  - Pendência: `softDelete()` (linha 100-102) ainda usa ternary inline:
    ```java
    Object[] params = hasTenantColumn
        ? new Object[]{now, deletedBy, id, TenantContext.getTenantId()}
        : new Object[]{now, deletedBy, id};
    ```
  - Refatorar para: `Object[] params = buildParams(now, deletedBy, id);`
  - `save()` e `update()` também têm branching inline para interpolar SQL dinâmico — isso é inevitável e não será refatorado
- **Arquivos a modificar:** `repository/common/BaseRepository.java`
- **Skills:** 141-java-refactoring-with-modern-features

### T-041.DT-021 — Implementar captura de valores "antes" no AuditAspect

- **Critério DONE:** Colunas previous_value/new_value populadas. PRD §6.4 atendido.
- **Estimativa:** 3h
- **Abordagem:**
  - Schema (V001), entidade (AuditEntry), DTO (AuditEntryResponse) e RowMapper (AuditEntryRowMapper) já suportam `previous_value` e `new_value` como JSONB/String
  - Apenas o `writeAuditLog()` não os popula — INSERE só 6 colunas
  - **Estratégia:**
    1. Mudar `@AfterReturning` para `@Around` para interceptar antes e depois
    2. Antes de `joinPoint.proceed()`: extrair entityId, consultar estado atual via `jdbc.queryForMap("SELECT * FROM ... WHERE id = ?")`
    3. Executar `joinPoint.proceed()` — manter retorno
    4. Após execução: consultar novo estado da entidade
    5. Serializar previous/new como JSON via `ObjectMapper`
    6. Passar para `writeAuditLog()` que insere as 8 colunas (6 + previous_value + new_value)
    7. `@Around` propaga exceções corretamente (não engole)
  - **Simplificação:** Se a entidade não existir antes (CREATE), previous_value = null. Se for delete/soft-delete, new_value = null.
- **Arquivos a modificar:** `security/aspect/AuditAspect.java`
- **Riscos:** @Around precisa de try-catch para propagar exceções. Consultas extras adicionam latência (~5-10ms por operação).
- **Skills:** 141-java-refactoring-with-modern-features, 311-frameworks-spring-jdbc

### T-043.DT-026 — Refatorar RLSIsolationTest para Testcontainers + PostgreSQL real

- **Critério DONE:** Testes de RLS com queries reais cross-tenant
- **Estimativa:** 4h
- **Abordagem:**
  - Manter testes estruturais existentes (validação de arquivos SQL) em métodos separados
  - Adicionar nova classe aninhada ou métodos que usam Testcontainers:
    1. Estender `BaseIntegrationTest`
    2. Criar 2 tenants (tenantA, tenantB) no seed
    3. Inserir dados em tabelas com RLS (subscription, business_unit, audit_log — "user" é palavra reservada, evitar)
    4. `SET app.current_tenant_id = '<tenantA_id>'` → SELECT subscription → apenas dados do tenantA
    5. `SET app.current_tenant_id = '<tenantB_id>'` → SELECT subscription → apenas dados do tenantB
    6. Sem `SET` → erro "unrecognized configuration parameter" (confirma que RLS exige o parâmetro)
    7. Tentar INSERT com tenant_id diferente do app.current_tenant_id → erro de política
  - Usar `DriverManagerDataSource` + Flyway manual (padrão já estabelecido)
- **Arquivos a modificar:** `integration/security/RLSIsolationTest.java`
- **Riscos:** "user" é palavra reservada SQL — usar aspas duplas. fbso_test é owner das tabelas e bypassa RLS por padrão — precisa usar `ALTER TABLE ... FORCE ROW LEVEL SECURITY` ou criar usuário sem privilégios de owner.
- **Skills:** 322-frameworks-spring-boot-testing-integration-tests

---

## 3. Ordem de Execução

| Ordem | Task | Tempo | Dependências |
|:---:|:---|:---:|:---|
| 1 | **T-045** — Testcontainers update | 15min | Nenhuma |
| 2 | **T-042** — AccessDeniedException handler | 10min | Nenhuma |
| 3 | **T-039 + T-040** — Documentação V004 + day-by-day | 30min | Nenhuma |
| 4 | **T-044** — BaseRepository softDelete refactor | 1h | Nenhuma |
| 5 | **T-041** — AuditAspect previous_value/new_value | 3h | Nenhuma |
| 6 | **T-043** — RLSIsolationTest com Testcontainers | 4h | Docker |

---

## 4. Estratégia de Build e Verificação

- **Compilação:** `./mvnw compile -Dcheckstyle.skip=true -q`
- **Testes unitários:** `./mvnw test -Dcheckstyle.skip=true`
- **Testes completos:** `./mvnw verify -Dcheckstyle.skip=true`

---

## Rodapé

🤖 *Plano gerado em 17/07/2026 pelo PROMPT-EXECUTE-SPRINT-TASKS (Fase 1). 7 tasks, ~10h estimadas. Stack: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17. Última frente da Sprint 3 — após conclusão, 42/42 tasks concluídas.*
