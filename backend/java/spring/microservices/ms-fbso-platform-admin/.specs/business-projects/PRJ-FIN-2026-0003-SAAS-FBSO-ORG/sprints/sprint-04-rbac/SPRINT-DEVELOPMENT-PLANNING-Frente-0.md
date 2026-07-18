# SPRINT-DEVELOPMENT-PLANNING-Frente-0.md — Plano de Desenvolvimento: Sprint 4 — Frente 0

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** `PRJ-FIN-2026-0003-SAAS-FBSO-ORG`
- **Sprint:** 4 de 7 — Frente 0 (Correções Pré-Sprint)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17, JDBC Template (não JPA), Flyway, Maven
- **Branch:** `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-04-rbac`
- **Data:** 17/07/2026
- **Origem:** [IDENTIFIED-TECHNICAL-DEBT-sprint-04-rbac.md](./IDENTIFIED-TECHNICAL-DEBT-sprint-04-rbac.md) — 20 débitos bloqueantes (DT-048 a DT-067)

---

## 1. Visão Geral

- **Sprint Goal (contexto):** "4 papéis (Admin Tenant, Gerente BU, Operador BU, Auditor) aplicados com matriz de permissões RN10-01. Gestão de usuários com convite por e-mail. Vinculação Usuário × Unidade × Módulo. Bloqueio de acesso direto com 403 amigável em PT-BR."
- **Objetivo da Frente 0:** Resolver 20 débitos técnicos bloqueantes ANTES de iniciar as 11 tasks de features RBAC (T-046 a T-056)
- **Tasks a implementar:** 20 (T-096.DT-048 a T-115.DT-067)
- **Ordem de execução:** Mista — blocos paralelizáveis com dependências sequenciais entre blocos
- **Estimativa total:** ~16-24h (2-3 dias com 1 dev)
- **Stack detectada:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + JDBC Template + Flyway 10.22.0 + Maven

### Status Inicial

| Métrica | Valor |
|:---|:---|
| Branch ativa | `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-04-rbac` ✅ |
| Entidades existentes | 4 (`Tenant`, `Plan`, `Subscription`, `AuditEntry`) |
| Migrations existentes | V001, V002, V003, V005 (+ U003, U005) |
| Dependências ausentes | Caffeine, REST Assured (apenas comentários no pom.xml) |
| T-113 (diagrama TASKS.md) | ✅ Já concluído durante a auditoria |

---

## 2. Dependências entre Tasks

```
BLOCO A — Dependências (paralelo, ~30min)
    T-096 (Caffeine + CacheConfig)
    T-097 (REST Assured)
    │
    ▼
BLOCO B — Entities (paralelo, ~2h)
    T-101 (User.java + UserRepository.java)
    T-102 (ResourceAction.java + RoleResource.java)
    T-103 (BusinessUnit.java)
    │
    ├── T-102 ──► T-105 (V004 seed RBAC — precisa das tabelas)
    ├── T-103 ──► T-104 (V006 FK — precisa da tabela business_unit)
    └── T-101 ──► T-098 (merge JWT×DB — precisa do UserRepository)
    │
    ▼
BLOCO C — Migrations + Enum (paralelo, ~2h)
    T-104 (V006 FK user_permission → business_unit + U006)
    T-105 (V004 seed RBAC: resource_action + role_resource + U004)
    T-106 (Role enum migration no RbacAspect)
    │
    ├── T-105 ──► T-100 (RbacAspect DB-backed — precisa do seed)
    ├── T-105 ──► T-099 (Cache strategy — precisa da matriz dimensionada)
    └── T-106 ──► T-100 (RbacAspect — usa Role enum)
    │
    ▼
BLOCO D — Refatoração RbacAspect (majoritariamente sequencial, ~9h)
    T-098 (Estratégia merge JWT×DB roles)
    │
    ├──► T-099 (Abandonar cache TTL 5min → findAll() indexado)
    │
    ├──► T-100 (Refatorar RbacAspect: inject PermissionService, remover Sets)
    │
    └──► T-115 (BU scope check via TenantContext.getBusinessUnitIds())
    │
    ▼
BLOCO E — Correções Pontuais (paralelo, ~7h)
    T-107 (TenantController.list() → TenantService)
    T-108 (JWT iss validation no SecurityConfig)
    T-109 (RLS FORCE nas 4 tabelas — migration V003 update)
    T-110 (Tenant isolation bypass — SubscriptionService)
    T-111 (Connection leak — TenantAwareDataSource)
    T-112 (Alinhar matriz RBAC — decisão SPRINT-CARD vs RbacAspect)
    T-114 (Corrigir header SPRINT-TEST-SUITE 19→27)
    │
    ▼
BLOCO F — Documentação (concluído ou trivial)
    T-113 ✅ JÁ CONCLUÍDO — diagrama TASKS.md §3 corrigido na auditoria
```

---

## 3. Plano por Task

### T-096.DT-048 — Adicionar Caffeine + CacheConfig

- **Critério DONE:** Caffeine no classpath. CacheConfig com `@EnableCaching`
- **Estimativa:** 1h
- **Abordagem:** Adicionar `spring-boot-starter-cache` + `caffeine:3.2.4` ao pom.xml. Criar `CacheConfig.java` em `config/` com `@Configuration` + `@EnableCaching` + `CacheManager` bean usando `CaffeineCacheManager`. TTL default 5min (para uso futuro; RBAC usará `findAll()` indexado, não cache).
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../config/CacheConfig.java` | 🆕 | Configuração de cache com Caffeine |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `pom.xml` | 🔄 | Adicionar `spring-boot-starter-cache` + `caffeine:3.2.4` |
- **Dependências:** Nenhuma
- **Riscos:** Baixo — dependência padrão Spring Boot
- **Skills aplicáveis:** `110-java-maven-best-practices`, `301-frameworks-spring-boot-core`

### T-097.DT-049 — Adicionar REST Assured

- **Critério DONE:** REST Assured disponível nos testes
- **Estimativa:** 30min
- **Abordagem:** Adicionar `rest-assured:5.5.7` + `rest-assured-spring-mock-mvc:5.5.7` (test scope) ao pom.xml. Usado por T-056 (testes parametrizados RBAC).
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `pom.xml` | 🔄 | Adicionar REST Assured dependencies (test scope) |
- **Dependências:** Nenhuma
- **Riscos:** Baixo — dependência padrão de teste
- **Skills aplicáveis:** `110-java-maven-best-practices`

### T-098.DT-050 — Definir Estratégia Merge JWT×DB para Roles

- **Critério DONE:** Roles do banco prevalecem sobre JWT. RN11-03 funcional
- **Estimativa:** 3h
- **Abordagem:** Decisão arquitetural: `RbacAspect` NÃO usa `TenantContext.getRoles()` (JWT). Em vez disso, consulta `user_permission` pelo `user_id` do `TenantContext`. Isso garante que alterações de permissão no banco têm efeito imediato (RN11-03), sem esperar refresh do token JWT. A consulta é: `SELECT role FROM user_permission WHERE user_id = ? AND deleted_dt IS NULL`.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `security/aspect/RbacAspect.java` | 🔄 | Substituir `TenantContext.getRoles()` por consulta ao `PermissionRepository` |
- **Dependências:** T-101 (UserRepository existente)
- **Riscos:** Performance — 1 query extra por request. Mitigação: query indexada por `user_id` (<1ms)
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `121-java-object-oriented-design`

### T-099.DT-051 — Abandonar Cache TTL 5min no RBAC

- **Critério DONE:** Sem stale permissions. RN11-03 "efeito imediato" respeitado
- **Estimativa:** 2h
- **Abordagem:** A matriz RN10-01 tem tamanho fixo (4 roles × 8 resources × 4 actions = 32 linhas na tabela `role_resource`). Carregar a matriz completa via `findAll()` com query indexada é <1ms — mais rápido que cache hit. Remover qualquer lógica de TTL do `RbacAspect`. O `PermissionService` carrega a matriz no startup ou sob demanda com `@PostConstruct`.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `security/aspect/RbacAspect.java` | 🔄 | Remover lógica de cache TTL |
- **Dependências:** T-098 (decisão de arquitetura), T-105 (seed data disponível)
- **Riscos:** Se a matriz crescer no futuro (>1000 linhas), reavaliar cache. Para 32 linhas, não é necessário
- **Skills aplicáveis:** `301-frameworks-spring-boot-core`, `121-java-object-oriented-design`

### T-100.DT-052 — Refatorar RbacAspect (DB-backed)

- **Critério DONE:** RbacAspect sem Sets estáticos. Permissões 100% DB-backed
- **Estimativa:** 4h
- **Abordagem:** Esta é a task central da Frente 0. Refatorar `RbacAspect.checkPermission()` para:
  1. Extrair `userId` do `TenantContext`
  2. Consultar `PermissionService.getUserRoles(userId)` → obtém roles do banco (`user_permission`)
  3. Consultar `PermissionService.getPermissionMatrix()` → obtém a matriz `role × resource × action` do banco (`role_resource` + `resource_action`)
  4. Verificar permissão contra a matriz (não contra Sets hardcoded)
  5. ADMIN_TENANT continua com acesso implícito total (não requer registros em `user_permission`)
  6. Remover TODOS os `private static final Set<String>` das linhas 48-60
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `service/PermissionService.java` | 🆕 | Serviço de consulta de permissões (roles + matriz) |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `security/aspect/RbacAspect.java` | 🔄 | Refatoração completa — inject PermissionService, remover Sets |
- **Dependências:** T-102 (ResourceAction + RoleResource entities), T-105 (V004 seed), T-106 (Role enum)
- **Riscos:** Refatoração de alto impacto — o aspecto é usado em todos os endpoints. Testar com `mvn test` após cada passo
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `121-java-object-oriented-design`, `126-java-exception-handling`

### T-101.DT-053 — Criar User.java + UserRepository.java

- **Critério DONE:** UserRepository funcional. Índice parcial email único por tenant ativo
- **Estimativa:** 2h
- **Abordagem:** Criar `User.java` entity estendendo `BaseEntity`. Campos: `id`, `tenantId`, `externalKeycloakId`, `email`, `name`, `status` (UserStatus enum — será adotado daqui). `toColumnMap()` mapeia todos os campos. `UserRepository` estende `BaseRepository<User>` com método `findByEmailAndTenant(String email, UUID tenantId)` usando o índice parcial `unique_email_active` (V002).
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `entity/User.java` | 🆕 | Entidade de usuário — 8 campos + auditoria |
  | `repository/UserRepository.java` | 🆕 | Repository com `findByEmailAndTenant()` |
- **Dependências:** Nenhuma (tabela `user` já existe desde V001)
- **Riscos:** Tabela usa aspas (`"user"` — palavra reservada SQL). O `BaseRepository` precisa lidar com tableName=`"user"` corretamente
- **Skills aplicáveis:** `311-frameworks-spring-jdbc`, `121-java-object-oriented-design`

### T-102.DT-054 — Criar ResourceAction.java + RoleResource.java

- **Critério DONE:** Entities compilam. Mapeamento colunas correto
- **Estimativa:** 1h
- **Abordagem:** Criar entities simples (não estendem `BaseEntity` — não têm `tenant_id`; a matriz RBAC é global). `ResourceAction`: `id`, `resourceName`, `action`. `RoleResource`: `id`, `role` (String do enum Role), `resourceActionId` (FK). Ambas com `toColumnMap()` e RowMapper. Sem soft delete (matriz é imutável por tenant).
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `entity/ResourceAction.java` | 🆕 | Entidade recurso+ação RBAC |
  | `entity/RoleResource.java` | 🆕 | Entidade papel×recurso (matriz RN10-01) |
- **Dependências:** Nenhuma (tabelas existem desde V001)
- **Riscos:** Baixo — entities simples sem lógica de negócio
- **Skills aplicáveis:** `311-frameworks-spring-jdbc`, `121-java-object-oriented-design`

### T-103.DT-055 — Criar BusinessUnit.java

- **Critério DONE:** Entity compila. Referenciável por `UserPermission`
- **Estimativa:** 1h
- **Abordagem:** Criar `BusinessUnit.java` estendendo `BaseEntity`. Campos: `id`, `tenantId`, `name`, `cnpj`, `hierarchyType`, `parentId`. `toColumnMap()` mapeia todos os campos. Entity mínima — sem repository dedicado nesta sprint (o CRUD de BU é Sprint 6). O propósito aqui é apenas ter a classe Java para referência no `UserPermission`.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `entity/BusinessUnit.java` | 🆕 | Entidade de unidade de negócio (mínima) |
- **Dependências:** Nenhuma (tabela existe desde V001)
- **Riscos:** Entity incompleta — sem RowMapper dedicado, sem repository. Apenas referência estrutural
- **Skills aplicáveis:** `311-frameworks-spring-jdbc`, `121-java-object-oriented-design`

### T-104.DT-056 — Criar Migration V006 (FK user_permission → business_unit)

- **Critério DONE:** FK criada. Rollback testado
- **Estimativa:** 30min
- **Abordagem:** Criar `V006__add_fk_user_permission_bu.sql` com `ALTER TABLE fbso_platform.user_permission ADD CONSTRAINT fk_up_bu FOREIGN KEY (business_unit_id) REFERENCES fbso_platform.business_unit(id)`. Criar `U006__drop_fk_user_permission_bu.sql` revertendo. A FK garante integridade referencial quando `UserPermission` for implementado na Sprint 4.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `db/migration/V006__add_fk_user_permission_bu.sql` | 🆕 | ALTER TABLE ADD CONSTRAINT |
  | `db/migration/U006__drop_fk_user_permission_bu.sql` | 🆕 | ALTER TABLE DROP CONSTRAINT |
- **Dependências:** T-103 (tabela `business_unit` deve existir — já existe desde V001)
- **Riscos:** Se houver dados órfãos em `user_permission` (business_unit_id que não existe), a FK falha. Verificar antes
- **Skills aplicáveis:** `313-frameworks-spring-db-migrations-flyway`

### T-105.DT-057 — Criar Migration V004 (Seed Matriz RN10-01)

- **Critério DONE:** Seed carrega. Matriz RN10-01 completa. Rollback funcional
- **Estimativa:** 1h
- **Abordagem:** Criar `V004__seed_rbac_matrix.sql` com INSERTs para:
  - 8 resources: DASHBOARD, TENANT, PLAN, SUBSCRIPTION, USER, PERMISSION, BUSINESS_UNIT, PRODUCT_SERVICE, AUDIT
  - 4 actions: view, create, edit, delete (32 resource_actions)
  - 4 roles × resource_actions conforme RN10-01 (SPRINT-CARD matrix restritiva)
  - ADMIN_TENANT: todos os 32 resource_actions
  - MANAGER_BU: BUSINESS_UNIT(view,create,edit), PRODUCT_SERVICE(view,create,edit)
  - OPERATOR_BU: BUSINESS_UNIT(view), PRODUCT_SERVICE(view)
  - AUDITOR: AUDIT(view)
  Criar `U004__rollback_rbac_seed.sql` com DELETE FROM role_resource + DELETE FROM resource_action.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `db/migration/V004__seed_rbac_matrix.sql` | 🆕 | INSERT resource_action + role_resource |
  | `db/migration/U004__rollback_rbac_seed.sql` | 🆕 | DELETE rollback |
- **Dependências:** T-102 (tabelas `resource_action` e `role_resource` existem desde V001)
- **Riscos:** A matriz restritiva (SPRINT-CARD) causa breaking change para MANAGER_BU (perde acesso a Dashboard/Tenants/Plans). Documentado no plano de cutover (T-112)
- **Skills aplicáveis:** `313-frameworks-spring-db-migrations-flyway`

### T-106.DT-058 — Migrar RbacAspect para Role Enum

- **Critério DONE:** Role enum referenciado. Type-safety na matriz
- **Estimativa:** 1h
- **Abordagem:** Substituir strings literais `"ADMIN_TENANT"`, `"MANAGER_BU"`, `"OPERATOR_BU"`, `"AUDITOR"` por `Role.ADMIN_TENANT`, `Role.MANAGER_BU`, etc. O enum `Role.java` já existe em `enums/`. Ajustar `RbacAspect` e `PermissionService` para usar `Role` em vez de `String`. A consulta ao banco retorna strings que são convertidas via `Role.valueOf()`.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `security/aspect/RbacAspect.java` | 🔄 | Substituir strings por Role enum |
  | `service/PermissionService.java` | 🔄 | Usar Role enum na consulta e mapeamento |
- **Dependências:** Nenhuma (enum já existe)
- **Riscos:** `Role.valueOf()` lança `IllegalArgumentException` se o banco tiver valor inválido. Adicionar tratamento
- **Skills aplicáveis:** `121-java-object-oriented-design`, `126-java-exception-handling`

### T-107.DT-059 — Refatorar TenantController.list()

- **Critério DONE:** RBAC aplicado no endpoint de listagem. TenantRepository não injetado no controller
- **Estimativa:** 1h
- **Abordagem:** Remover injeção de `TenantRepository` do `TenantController`. O método `list()` passa a chamar `tenantService.findAll(request)` (método a ser adicionado no `TenantService`). O `@RequiresPermission(resource="TENANT", action="view")` já existe no controller. Isso garante que o RBAC seja aplicado uniformemente via AOP.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `controller/TenantController.java` | 🔄 | Remover injeção direta de TenantRepository. Usar TenantService |
  | `service/TenantService.java` | 🔄 | Adicionar método `findAll(request)` delegando ao repository |
- **Dependências:** Nenhuma
- **Riscos:** Verificar que `TenantService` já existe e tem os métodos necessários
- **Skills aplicáveis:** `302-frameworks-spring-boot-rest`, `121-java-object-oriented-design`

### T-108.DT-060 — Adicionar JWT Issuer Validation

- **Critério DONE:** JWT de realms não autorizados → 401
- **Estimativa:** 1h
- **Abordagem:** Modificar o bean `jwtDecoder()` em `SecurityConfig.java` para adicionar validação de issuer. O `issuerUri` já está configurado em `application.yml` (`spring.security.oauth2.resourceserver.jwt.issuer-uri`). Usar `JwtValidators.createDefaultWithIssuer(issuerUri)` e `DelegatingJwtValidator`. Injetar `@Value("${spring.security.oauth2.resourceserver.jwt.issuer-uri}") String issuerUri`.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `config/SecurityConfig.java` | 🔄 | Adicionar JwtValidators.createDefaultWithIssuer() |
- **Dependências:** Nenhuma
- **Riscos:** Se `issuer-uri` não estiver configurado corretamente em staging/prod → 401 para todos. Validar configuração
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `124-java-secure-coding`

### T-109.DT-061 — Adicionar FORCE ROW LEVEL SECURITY

- **Critério DONE:** RLS aplicado inclusive para table owner. RLSIsolationTest passa sem FORCE manual
- **Estimativa:** 2h
- **Abordagem:** Modificar `V003__enable_rls.sql` para adicionar `ALTER TABLE fbso_platform.<table> FORCE ROW LEVEL SECURITY` nas 4 tabelas: `subscription`, `user`, `business_unit`, `audit_log`. Atualizar `U003__disable_rls.sql` correspondente. Atualizar `RLSIsolationTest.java` para remover `FORCE` manual do setup (linhas 61, 188). Remover `@Disabled` se possível.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `db/migration/V003__enable_rls.sql` | 🔄 | Adicionar FORCE ROW LEVEL SECURITY nas 4 tabelas |
  | `db/migration/U003__disable_rls.sql` | 🔄 | Adicionar NO FORCE ROW LEVEL SECURITY correspondente |
  | `test/.../integration/security/RLSIsolationTest.java` | 🔄 | Remover FORCE manual do setup |
- **Dependências:** Nenhuma
- **Riscos:** `FORCE ROW LEVEL SECURITY` afeta TODAS as queries do owner — testar exaustivamente. Pode quebrar migrations Flyway que rodam como owner
- **Skills aplicáveis:** `313-frameworks-spring-db-migrations-flyway`, `124-java-secure-coding`

### T-110.DT-062 — Corrigir Tenant Isolation Bypass

- **Critério DONE:** Tenant mismatch → 403. Queries com tenant_id do contexto
- **Estimativa:** 2h
- **Abordagem:** No `SubscriptionService.create()` e `changePlan()`, validar que `tenantId` da URL é igual a `TenantContext.getTenantId()`. Se diferente → `PermissionDeniedException`. Modificar `SubscriptionRepository.findActiveByTenantId()` para usar `tenantClause()` do `BaseRepository` (consistente com RLS). Adicionar log WARN em caso de mismatch.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `service/SubscriptionService.java` | 🔄 | Validar tenant_id URL vs JWT |
  | `repository/SubscriptionRepository.java` | 🔄 | Adicionar tenantClause() nas queries custom |
- **Dependências:** Nenhuma
- **Riscos:** Pode quebrar fluxos legítimos onde admin acessa dados de outro tenant (cross-tenant admin). Se necessário, criar role FBSO_ADMIN com bypass explícito
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `124-java-secure-coding`

### T-111.DT-063 — Corrigir Connection Leak no TenantAwareDataSource

- **Critério DONE:** Connection leak eliminado. Pool estável
- **Estimativa:** 1h
- **Abordagem:** No `TenantAwareDataSource.applyTenantContext()`, mover `UUID.fromString(tenantId)` para ANTES do bloco que usa a conexão. Se `IllegalArgumentException` for lançada, a conexão NÃO foi obtida do pool ainda → sem leak. Alternativa: manter no try mas adicionar `catch (RuntimeException e) { try { conn.close(); } catch (SQLException ignored) {} throw new TenantIsolationException(...); }`.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `config/TenantAwareDataSource.java` | 🔄 | Validar UUID antes de obter conexão OU fechar conexão no catch |
- **Dependências:** Nenhuma
- **Riscos:** Baixo — correção pontual
- **Skills aplicáveis:** `126-java-exception-handling`, `311-frameworks-spring-jdbc`

### T-112.DT-064 — Alinhar Matriz RBAC

- **Critério DONE:** Matriz unificada. MANAGER_BU impacto documentado
- **Estimativa:** 1h
- **Abordagem:** Decisão de design: adotar a **matriz restritiva do SPRINT-CARD** (RN10-01 como especificado). MANAGER_BU tem acesso apenas a BUSINESS_UNIT e PRODUCT_SERVICE. Documentar breaking change: MANAGER_BU perde acesso a Dashboard, Tenants, Plans, Subscriptions. Plano de cutover: seed V004 + deploy do RbacAspect refatorado no mesmo release. Atualizar SPRINT-CARD.md com nota de breaking change.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `SPRINT-CARD.md` | 🔄 | Adicionar nota de breaking change para MANAGER_BU |
- **Dependências:** Decisão do time (já tomada na auditoria)
- **Riscos:** MANAGER_BU existentes perdem acesso. Comunicar ao frontend e ao PO
- **Skills aplicáveis:** `121-java-object-oriented-design`

### T-113.DT-065 — Corrigir Diagrama TASKS.md §3 ✅ JÁ CONCLUÍDO

- **Status:** ✅ Concluído durante a auditoria (Passo 5). Diagrama FASE 3 corrigido de T-039..T-047 para T-046..T-056.
- **Estimativa:** 0min (já executado)

### T-114.DT-066 — Corrigir Header SPRINT-TEST-SUITE 19→27

- **Critério DONE:** Métricas corretas em todos os artefatos
- **Estimativa:** 15min
- **Abordagem:** Alterar `Total de cenários: 19` para `27` no header do `SPRINT-TEST-SUITE.md`. Atualizar `SPRINT-CARD.md` métricas: "Cenários de teste RBAC: 19" → "27".
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `SPRINT-TEST-SUITE.md` | 🔄 | Corrigir header L6: 19→27 |
  | `SPRINT-CARD.md` | 🔄 | Corrigir métricas L101: 19→27 |
- **Dependências:** Nenhuma
- **Riscos:** Nenhum
- **Skills aplicáveis:** N/A (edição de documentação)

### T-115.DT-067 — Adicionar BU Scope Check

- **Critério DONE:** MANAGER_BU acessa apenas BUs designadas. RN10-01 escopo correto
- **Estimativa:** 3h
- **Abordagem:** O `RbacAspect` (ou `PermissionService`) deve validar que o `business_unit_id` do request está na lista `TenantContext.getBusinessUnitIds()` do JWT. Para ADMIN_TENANT, bypass (acesso total). Para MANAGER_BU e OPERATOR_BU, validar escopo. Implementar como método no `PermissionService.validateBusinessUnitAccess(UUID businessUnitId)`. Chamado pelos services antes de operações em recursos com escopo BU.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `service/PermissionService.java` | 🔄 | Adicionar método `validateBusinessUnitAccess()` |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `service/ProductService.java` (a criar na Sprint 6) | 🔄 | Chamar validação de BU |
- **Dependências:** T-100 (PermissionService existente)
- **Riscos:** `TenantContext.getBusinessUnitIds()` nunca foi usado em produção — validar que o JWT contém a claim. Se não contiver, fallback para lista vazia = acesso negado para não-admin
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `126-java-exception-handling`

---

## 4. Ordem de Execução

1. **T-113** ✅ Já concluído — diagrama TASKS.md corrigido
2. **T-096 + T-097** (paralelo) — Dependências Maven (~30min)
3. **T-101 + T-102 + T-103** (paralelo) — Entities (~2h)
4. **T-104 + T-105 + T-106** (paralelo) — Migrations + Enum (~2h)
5. **T-098** — Estratégia merge JWT×DB (3h) — requisito para T-099, T-100
6. **T-099** — Abandonar cache TTL (2h) — requisito para T-100
7. **T-100** — Refatorar RbacAspect DB-backed (4h) — task central
8. **T-115** — BU scope check (3h) — depende de T-100
9. **T-107 + T-108 + T-109 + T-110 + T-111 + T-112 + T-114** (paralelo) — Correções pontuais (~7h)

**Checkpoints de build:**
- Após passo 3: `mvn compile` (entities compilam)
- Após passo 4: `mvn compile` (migrations criadas)
- Após passo 7: `mvn test` (RbacAspect refatorado não quebra testes existentes)
- Após passo 9: `mvn clean verify` (suite completa)

---

## 5. Estratégia de Build e Verificação

- **Comando de build:** `mvn compile`
- **Comando de teste rápido:** `mvn test`
- **Comando de teste completo:** `mvn clean verify`
- **Comando de cobertura:** `mvn jacoco:report`

### Checkpoints

| Checkpoint | Após | Comando | Critério |
|:---|:---|:---|:---|
| CP1 | Bloco B (entities) | `mvn compile` | BUILD SUCCESS |
| CP2 | Bloco C (migrations) | `mvn compile` | BUILD SUCCESS |
| CP3 | Bloco D (RbacAspect) | `mvn test` | Todos os testes passando |
| CP4 | Bloco E (correções) | `mvn clean verify` | BUILD SUCCESS + testes verdes |

### Estratégia de Rollback

Cada task que modifica código de produção deve ser comitável independentemente. Se uma task quebrar o build, reverter apenas aquela task e continuar com as demais. Tasks dos Blocos A, B, C, E são independentes entre si dentro de cada bloco.

---

🤖 *Documento gerado em 17/07/2026 por Agente IA. Baseado no PROMPT-EXECUTE-SPRINT-TASKS.md Fase 1 e nos 20 débitos bloqueantes do IDENTIFIED-TECHNICAL-DEBT-sprint-04-rbac.md.*
