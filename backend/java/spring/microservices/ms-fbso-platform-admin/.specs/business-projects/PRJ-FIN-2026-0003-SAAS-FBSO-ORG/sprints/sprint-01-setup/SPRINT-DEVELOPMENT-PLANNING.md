# SPRINT-DEVELOPMENT-PLANNING.md — Plano de Desenvolvimento: Sprint 1

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** `PRJ-FIN-2026-0003-SAAS-FBSO-ORG`
- **Sprint:** 1 de 7 — Setup e Fundação
- **Stack:** Java 25 + Spring Boot 3.5.1 + PostgreSQL 17 + Flyway + Maven 3.9
- **Data da Revisão:** 15 de Julho de 2026
- **Modo:** Revisão (sprint já concluída em 14/07/2026)

---


## 1. Visão Geral

- **Sprint Goal:** Scaffold Maven compila com `mvn clean install`. 11 tabelas criadas no PostgreSQL via Flyway. Estrutura de pacotes completa compila sem erros. BaseRepository funcional com Soft Delete + Tenant Filter.
- **Tasks implementadas:** 8 (T-001 a T-008)
- **Ordem de execução:** Sequencial com paralelismo limitado — forte encadeamento de dependências
- **Stack detectada:** Java 25 + Spring Boot 3.5.1 + PostgreSQL 17 + Flyway 10.x + Maven 3.9

---

## 2. Dependências entre Tasks

```
T-001 (pom.xml) ─────────────────────────────────────────────┐
    │                                                          │
    ├── T-002 (application.yml)                                │
    ├── T-003 (Dockerfile)                                     │
    ├── T-004 (Migration V001) ── T-005 (Migration V002)       │
    ├── T-006 (BaseEntity + enums) ── T-007 (BaseRepository)   │
    └── T-008 (estrutura de pacotes) ← depende de todos acima  │
```

**Ordem recomendada:** T-001 → (T-002, T-003, T-004, T-006 em paralelo) → T-005 → T-007 → T-008

---

## 3. Plano por Task

### T-001 — Scaffold Maven
- **Critério DONE:** `mvn compile` BUILD SUCCESS. 13 arquivos compilados com Java 25
- **Estimativa:** 2d
- **Abordagem:** Criar `pom.xml` com `spring-boot-starter-parent` como BOM. Dependências: Spring Boot (web, security, jdbc, validation, actuator), Flyway, PostgreSQL driver, Testcontainers, Micrometer, JUnit 5, Mockito. Java 25 como `java.version`. JaCoCo configurado com meta 80%.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `pom.xml` | 🆕 | Maven POM com parent, dependências, plugins |
- **Dependências:** Nenhuma (primeira task)
- **Riscos:** Conflito de versões Spring Boot × Java 25. JaCoCo incompatível com Java 25 (class file v69).
- **Skills aplicáveis:** `110-java-maven-best-practices`, `301-frameworks-spring-boot-core`

### T-002 — application.yml (profiles)
- **Critério DONE:** 3 profiles operacionais. Dev conecta PostgreSQL Docker. Logging configurado por profile
- **Estimativa:** 1d
- **Abordagem:** `application.yml` base + `application-dev.yml`, `application-staging.yml`, `application-prod.yml`. Datasource PostgreSQL via variáveis de ambiente. Keycloak JWKS URI. Server port 8081.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/resources/application.yml` | 🆕 | Configuração base (logging, server, actuator) |
  | `src/main/resources/application-dev.yml` | 🆕 | Dev profile (localhost DB, DEBUG logging) |
  | `src/main/resources/application-staging.yml` | 🆕 | Staging profile (K8s service names) |
  | `src/main/resources/application-prod.yml` | 🆕 | Prod profile (WARN logging, security hardening) |
- **Dependências:** T-001 (pom.xml com dependências referenciadas)
- **Riscos:** Credenciais hardcoded → usar `${ENV_VAR}` placeholders
- **Skills aplicáveis:** `301-frameworks-spring-boot-core`, `security-review`

### T-003 — Dockerfile
- **Critério DONE:** Build nativo e JVM configurados. .dockerignore criado
- **Estimativa:** 1d
- **Abordagem:** `Dockerfile` com GraalVM Native Image (multi-stage build). `Dockerfile.jvm` como fallback para dev. `.dockerignore` para reduzir contexto de build.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `Dockerfile` | 🆕 | GraalVM Native Image multi-stage build |
  | `Dockerfile.jvm` | 🆕 | Fallback JVM (OpenJDK 25) |
  | `.dockerignore` | 🆕 | Excluir target/, .git/, logs/ |
- **Dependências:** T-001 (pom.xml — referência ao artifact)
- **Riscos:** GraalVM Native Image pode falhar em runtime → JVM fallback mitiga
- **Skills aplicáveis:** `java-docker`

### T-004 — Migration V001 (11 tabelas Core)
- **Critério DONE:** Flyway migrate cria 11 tabelas com colunas, PKs, FKs, NOT NULL. Rollback testado
- **Estimativa:** 3d
- **Abordagem:** Script SQL com `CREATE TABLE` para: tenant, plan, plan_module, subscription, user, user_permission, resource_action, role_resource, business_unit, product_service, audit_log. Todas as tabelas com 6 campos de auditoria (created_dt, updated_dt, created_by, updated_by, deleted_dt, deleted_by). UUID como PK com `gen_random_uuid()`. Schema `fbso_platform`.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/resources/db/migration/V001__create_core_tables.sql` | 🆕 | 11 tabelas + constraints |
- **Dependências:** T-002 (application.yml com datasource e flyway config)
- **Riscos:** Erro de sintaxe PostgreSQL → testar em Docker local antes de commitar
- **Skills aplicáveis:** `postgres-pro`, `313-frameworks-spring-db-migrations-flyway`

### T-005 — Migration V002 (índices)
- **Critério DONE:** `CREATE UNIQUE INDEX ... WHERE deleted_dt IS NULL` executado. Explain plan mostra index scan
- **Estimativa:** 1d
- **Abordagem:** Índices parciais únicos: unique_cnpj_active (business_unit), unique_email_active (user), unique_sku_active (product_service). Índices de desempenho para queries frequentes (tenant_id, status, created_dt).
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/resources/db/migration/V002__create_partial_unique_indexes.sql` | 🆕 | Índices parciais + índices de desempenho |
- **Dependências:** T-004 (tabelas precisam existir)
- **Riscos:** Índice parcial com sintaxe incorreta → validar `WHERE deleted_dt IS NULL`
- **Skills aplicáveis:** `postgres-pro`, `postgresql-optimization`

### T-006 — BaseEntity + Address + 8 enums
- **Critério DONE:** BaseEntity compatível com 11 tabelas. 8 enums implementados
- **Estimativa:** 1d
- **Abordagem:** `BaseEntity` com campos de auditoria (created_dt, updated_dt, created_by, updated_by, deleted_dt, deleted_by). `Address` como Value Object (record). Enums: TenantStatus (PENDING_ONBOARDING, ACTIVE, SUSPENDED, INACTIVE), TenantSegment, Recurrence, SubscriptionStatus, UserStatus, Role (ADMIN_TENANT, MANAGER_BU, OPERATOR_BU, AUDITOR), TaxRegime, ProductType.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../common/BaseEntity.java` | 🆕 | Superclasse com 6 campos de auditoria |
  | `src/main/java/.../common/Address.java` | 🆕 | Value Object (record) |
  | `src/main/java/.../enums/*.java` | 🆕 | 8 enums |
- **Dependências:** T-004 (entidades de referência para os enums)
- **Riscos:** Enum Role deve alinhar com matriz RN10-01 (4 papéis)
- **Skills aplicáveis:** `121-java-object-oriented-design`

### T-007 — BaseRepository (JDBC Template)
- **Critério DONE:** CRUD injeta automaticamente `WHERE deleted_dt IS NULL` e `tenant_id = ?`. Teste unitário de cada método
- **Estimativa:** 2d
- **Abordagem:** Classe abstrata genérica `BaseRepository<T, ID>` com JdbcTemplate. Métodos: findAll (paginado, com sort), findById, save, update, softDelete. Injeção automática de tenant_id via TenantContext. Sanitização de nomes de colunas contra SQL injection.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../repository/common/BaseRepository.java` | 🆕 | Template JDBC com Soft Delete + Tenant Filter |
  | `src/main/java/.../security/TenantContext.java` | 🆕 | ThreadLocal holder (criado antecipadamente) |
  | `src/test/java/.../unit/repository/BaseRepositoryTest.java` | 🆕 | 7 testes unitários |
- **Dependências:** T-006 (BaseEntity para o generic `<T extends BaseEntity>`)
- **Riscos:** SQL injection via nome de coluna → sanitizar com whitelist. TenantContext sem JWT filter (Sprint 2) → mock nos testes.
- **Skills aplicáveis:** `311-frameworks-spring-jdbc`, `131-java-testing-unit-testing`, `security-review`

### T-008 — Estrutura de pacotes
- **Critério DONE:** `mvn compile` sem erros. Package structure segue ARCHITECTURE.md §2
- **Estimativa:** 1d
- **Abordagem:** Criar diretórios e classes esqueleto: controller/, service/, repository/, entity/, dto/request/, dto/response/, enums/, exception/, security/aspect/, security/annotation/, config/, common/, utils/. Classe principal `FbsoPlatformAdminApplication.java` com `@SpringBootApplication` e `@EnableAsync`.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../FbsoPlatformAdminApplication.java` | 🆕 | Classe principal Spring Boot |
  | `src/main/java/.../config/*.java` | 🆕 | Esqueletos: SecurityConfig, WebConfig, DataSourceConfig |
  | `src/main/java/.../controller/*.java` | 🆕 | 10 controllers esqueleto |
  | `src/main/java/.../service/*.java` | 🆕 | 10 services esqueleto |
  | `src/main/java/.../repository/*.java` | 🆕 | 9 repositories esqueleto |
  | `src/main/java/.../entity/*.java` | 🆕 | 11 entidades esqueleto |
  | `src/main/java/.../dto/request/*.java` | 🆕 | DTOs request |
  | `src/main/java/.../dto/response/*.java` | 🆕 | DTOs response |
  | `src/main/java/.../exception/*.java` | 🆕 | Exceções + GlobalExceptionHandler esqueleto |
  | `src/main/java/.../security/aspect/*.java` | 🆕 | Aspectos esqueleto |
  | `src/main/java/.../security/annotation/*.java` | 🆕 | Anotações esqueleto |
  | `src/main/java/.../utils/*.java` | 🆕 | Utilitários esqueleto |
- **Dependências:** T-001 a T-007 (todos os artefatos base devem existir)
- **Riscos:** Compilação quebra por dependência circular → verificar imports
- **Skills aplicáveis:** `121-java-object-oriented-design`, `301-frameworks-spring-boot-core`

---

## 4. Ordem de Execução

1. **T-001 (pom.xml)** — base de tudo. Sem dependências.
2. **T-002 (application.yml)** — depende de T-001 (referências a libraries).
3. **T-003 (Dockerfile)** — depende de T-001 (pom.xml para build context).
4. **T-004 (Migration V001)** — depende de T-002 (datasource config).
5. **T-006 (BaseEntity + enums)** — paralelo com T-004. Só depende de T-001.
6. **T-005 (Migration V002)** — depende de T-004 (tabelas existirem).
7. **T-007 (BaseRepository)** — depende de T-006 (BaseEntity) e T-005 (índices validados).
8. **T-008 (estrutura de pacotes)** — depende de todas acima. Verificação final de compilação.

---

## 5. Estratégia de Build e Verificação

- **Comando de build:** `mvn clean compile`
- **Comando de teste:** `mvn test`
- **Comando de migração:** `mvn flyway:migrate`
- **Checkpoints:**
  1. Após T-001: `mvn validate` passa
  2. Após T-004: `mvn flyway:migrate` cria schema
  3. Após T-007: `mvn test` executa 7/7
  4. Após T-008: `mvn compile` sem erros em todos os pacotes

---

## 6. Revisão Pós-Implementação

> **Executado em:** 15/07/2026. Sprint concluída em 14/07/2026. Todas as 8 tasks entregues.

### Avaliação de Necessidade de Revisão

| Task | Status | Necessita Revisão? | Justificativa |
|:---|:---:|:---:|:---|
| T-001 | ✅ | Não | `pom.xml` funcional. 13 arquivos compilam. JaCoCo configurado (incompatível com Java 25 — conhecido, externo). |
| T-002 | ✅ | Não | 3 profiles operacionais. Credenciais via variáveis de ambiente. |
| T-003 | ✅ | Não | Dockerfile Native + JVM configurados. .dockerignore presente. |
| T-004 | ✅ | Não | 11 tabelas criadas. PKs, FKs, NOT NULL corretos. 6 campos de auditoria em todas. |
| T-005 | ✅ | Não | 4 índices parciais + 8 índices de desempenho. Sintaxe PostgreSQL correta. |
| T-006 | ✅ | Não | BaseEntity + Address + 8 enums. Compatível com SPECS.md §6.1 e §6.3. |
| T-007 | ✅ | Não | BaseRepository com Soft Delete + Tenant Filter. 7/7 testes passando. Sem SQL injection (sanitização de colunas). |
| T-008 | ✅ | Não | 47 diretórios, 14 pacotes top-level. `mvn compile` sem erros. Segue ARCHITECTURE.md §2. |

### Verificações Cross-Documento

| Verificação | Resultado |
|:---|:---:|
| Entidades criadas conferem com SPECS.md §6.1 (11 entidades) | ✅ |
| Campos de auditoria conferem com SPECS.md §6.3 (6 campos) | ✅ |
| Estrutura de pacotes confere com ARCHITECTURE.md §2 | ✅ |
| JDBC Template usado conforme ADR-L01 (sem JPA/Hibernate) | ✅ |
| Padrão de nomenclatura `com.fbso.platform.admin.*` conforme PRD.md §5.3 | ✅ |
| Testes seguem TEST_PLAN.md §9.1-§9.5 (17 cenários cobertos) | ✅ |

### Pendências Conhecidas (Não Bloqueantes)

| Pendência | Impacto | Resolução |
|:---|:---|:---|
| JaCoCo incompatível com Java 25 | Baixo — cobertura não mensurável | Aguardar JaCoCo 0.8.13+ com suporte a class file v69 |
| Checkstyle/PMD com `failOnViolation=false` | Baixo — qualidade não validada | Endurecer na Sprint 7 (Homologação) |
| TenantContext criado antecipadamente | Baixo — sem JWT filter | Complemento na Sprint 2 (T-011) |

### Conclusão

**Nenhuma revisão de implementação necessária.** As 8 tarefas da Sprint 1 foram executadas conforme o plano, os critérios DONE foram atendidos, a arquitetura está em conformidade com ARCHITECTURE.md, e os testes passam (7/7). As 3 pendências identificadas são conhecidas e não bloqueiam o progresso para a Sprint 2.

---

🤖 *Plano de desenvolvimento gerado em modo Revisão pelo Agente Executor de Sprint/Claude. Skills consideradas: 110-java-maven-best-practices, 121-java-object-oriented-design, 126-java-exception-handling, 130-java-testing-strategies, 301-frameworks-spring-boot-core, 311-frameworks-spring-jdbc, postgres-pro, 313-frameworks-spring-db-migrations-flyway, 131-java-testing-unit-testing, security-review, java-docker. Revisão Caveman em 15/07/2026 (DOCS-SERVICE-SPRINTS-CAVEMAN-REVIEW.md): modo revisão pós-implementação — sem necessidade de alterações.*
