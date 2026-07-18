# IDENTIFIED-TECHNICAL-DEBT-sprint-04-rbac

- **Sprint alvo:** 4 de 7 — Governança de Acessos (RBAC)
- **Marco:** M4 (EP-03)
- **Data da análise:** 17/07/2026
- **Skills executadas:** code-reviewer, caveman-review, superpowers:brainstorming, ponytail-review, ponytail-debt, code-review (CodeRabbit), jscpd+dry-refactoring, tech-debt, security-review
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17, JDBC Template (não JPA), Flyway
- **Total de achados:** 47 (20 🔴 críticos, 15 🟡 riscos, 12 🔵 nits) + 9 backlog Sprint 3
- **Impeditivos para iniciar a sprint:** 20 SIM (devem ser corrigidos no Frente 0 antes do primeiro commit de feature)

---

## Resumo Executivo

A auditoria multidisciplinar com 9 skills revelou **47 novos débitos técnicos**, somados a **9 débitos não resolvidos da Sprint 3** (backlog). **20 são impeditivos** para iniciar a Sprint 4 (RBAC). Os 5 achados mais graves:

1. **RLS sem `FORCE ROW LEVEL SECURITY`** (DT-061) — PostgreSQL não aplica RLS ao table owner. Flyway cria tabelas como owner → app user = owner → **RLS 100% ineficaz**. O `RLSIsolationTest.java` já adiciona `FORCE` durante o setup (linhas 61, 188) provando que a falha é conhecida mas nunca corrigida na migration.

2. **JWT decoder sem validação de `iss`** (DT-062) — aceita tokens de QUALQUER realm Keycloak. Um atacante com seu próprio realm pode forjar tokens aceitos por esta aplicação.

3. **Duas fontes de verdade para roles** (DT-050) — `RbacAspect` lê roles do JWT (`TenantContext.getRoles()`), mas Sprint 4 gerencia permissões no banco (`user_permission`). Sem resolução de conflito, alterações no banco não têm efeito até novo login (até 60min). RN11-03 ("efeito imediato") impossível.

4. **Dependências bloqueantes ausentes** (DT-058, DT-059) — Caffeine e REST Assured aparecem como comentários no `pom.xml` sem os blocos `<dependency>`. T-053 (cache) e T-056 (testes) não implementáveis.

5. **Cache TTL 5min viola RN11-03** (DT-051) — Se admin revoga permissão às 10:00, usuário ainda acessa até 10:05. Violação contratual da spec. Recomendação: abandonar cache; carregar matriz de 32 linhas com query indexada (<1ms).

**Estimativa Frente 0:** 16-24h (2-3 dias) para resolver os 20 impeditivos antes de iniciar as 11 tasks RBAC.

---

## Backlog de Débitos Técnicos (Sprints Anteriores)

> Débitos identificados na Sprint 3 que **permanecem não resolvidos** e são candidatos a tratamento na Sprint 4.

| DT-XXX | Sprint Origem | Descrição | Severidade | Bloqueante? | Status | Resolução |
|:---|:---|:---|:---:|:---:|:---|:---|
| DT-023 | Sprint 3 | Migrar paginação offset→keyset no BaseRepository | 🟡 | NÃO | Pendente (Sprint 5) | ↗ ver T-087.DT-023 em TASKS.md |
| DT-030 | Sprint 3 | Consolidar dupla validação JWT via `Converter<Jwt, AbstractAuthToken>` customizado | 🟡 | NÃO | Pendente (Sprint 4) | Eliminar `JwtAuthenticationFilter` redundante; usar `Converter` no `SecurityConfig` |
| DT-031 | Sprint 3 | Reduzir `maxAllowedViolations` Checkstyle: 300→100→0 | 🔵 | NÃO | Pendente (Sprint 4-7) | Redução progressiva a cada sprint |
| DT-034 | Sprint 3 | Remover `Address.java` se não usado até Sprint 6 | 🔵 | NÃO | Pendente (Sprint 6) | 94 linhas de código morto |
| **DT-035** | Sprint 3 | Migrar `RbacAspect` de strings literais para `Role` enum | 🟡 | **SIM** | Pendente (Sprint 4) | `Role.java` existe mas `RbacAspect` usa strings; type-safety zero |
| DT-042 | Sprint 3 | Criar `docker-compose.yml` com PostgreSQL + Keycloak + MailHog | 🟡 | NÃO | Pendente (Sprint 4) | Setup local ~30min/dev sem ele |
| DT-043 | Sprint 3 | Criar script seed SQL com 50+ tenants + 3 planos | 🔵 | NÃO | Pendente (Sprint 3→4) | Dados realistas para desenvolvimento |
| DT-044 | Sprint 3 | Criar `logback-spring.xml` com JSON appender + rotação | 🔵 | NÃO | Pendente (Sprint 7) | Logs stdout-only em produção |
| DT-045 | Sprint 3 | Migrar Flyway 10.22.0→12.11.0 (breaking changes analisados) | 🟡 | NÃO | Pendente (Sprint 5) | 2 majors atrás; migração 10→11→12 |

**Total em backlog:** 9 débitos pendentes da Sprint 3. **2 com Sprint 4 como alvo** (DT-030, DT-035). **DT-035 é bloqueante** para T-053.

---

## Matriz de Débitos Técnicos

> **Legenda:** CREV=code-reviewer, CR=caveman-review, ARCH=brainstorming, PONY=ponytail-review, PDBT=ponytail-debt, CODE=code-review, JSCPD=jscpd+dry-refactoring, DEBT=tech-debt, SEC=security-review, BACKLOG=sprints anteriores

### 🔴 Bloqueantes (20 — Frente 0 antes de iniciar features RBAC)

| ID | Sprint Origem | Arquivo/Artefato | Achado | Sev. | Skill | Compl. | Efeito se não tratado |
|:---|:---|:---|:---|:---:|:---:|:---:|:---|
| **DT-048** | Sprint 4 | `pom.xml:119-120` | **Caffeine ausente** — apenas comentário `<!-- Cache (Caffeine) — Sprint 4 -->`. T-053 requer cache de matriz RBAC | 🔴 | CREV, ARCH, DEBT | L | T-053 não implementável. Sem cache, N+1 queries por request |
| **DT-049** | Sprint 4 | `pom.xml:157` | **REST Assured ausente** — apenas comentário `<!-- REST Assured — Sprint 4 -->`. T-056 requer 20+ combinações papel×endpoint | 🔴 | CREV, ARCH | L | T-056 (testes parametrizados RBAC) não implementável |
| **DT-050** | Sprint 4 | `JwtAuthenticationFilter:92` + `RbacAspect:68` | **Conflito arquitetural**: roles vêm do JWT (`TenantContext.getRoles()`), Sprint 4 gerencia no banco (`user_permission`). Sem merge/override, alterações no banco não têm efeito até refresh do token (60min) | 🔴 | ARCH, CODE | H | RN11-03 impossível. Permissões stale por até 60min |
| **DT-051** | Sprint 4 | `SPRINT-CARD.md` T-053 + SPECS.md RN11-03 | **Cache TTL 5min viola RN11-03** "efeito imediato". Se admin revoga permissão, usuário ainda acessa por 5min. Violação contratual | 🔴 | ARCH | M | RN11-03 violada. Recomendação: abandonar cache; query indexada <1ms |
| **DT-052** | Sprint 4 | `RbacAspect.java:48-103` | **Matriz hardcoded** — sem injeção de repository/cache. Comentário `// ponytail: ceiling = Sprint 4` desde Sprint 2. Zero progresso | 🔴 | CREV, CODE, DEBT, PDBT | H | T-053 inviável sem refatoração completa do aspecto |
| **DT-053** | Sprint 4 | `src/main/java/.../entity/` | **`User.java` + `UserRepository.java` inexistentes**. Tabela `user` criada no V001 com RLS (V003), mas zero código Java | 🔴 | CREV, ARCH | M | T-046, T-047, T-048 bloqueados |
| **DT-054** | Sprint 4 | `src/main/java/.../entity/` | **`ResourceAction.java` + `RoleResource.java` inexistentes**. Tabelas no V001, vazias, sem entity | 🔴 | CREV, ARCH | M | T-049 e T-053 bloqueados |
| **DT-055** | Sprint 4 | `src/main/java/.../entity/` | **`BusinessUnit.java` inexistente**. Tabela `business_unit` no V001 referenciada por `user_permission.business_unit_id` | 🔴 | CREV | M | T-050 (`UserPermission`) não pode referenciar BU sem entity |
| **DT-056** | Sprint 4 | `V001__create_core_tables.sql:106` | **FK ausente** `user_permission.business_unit_id → business_unit(id)`. Órfãos quando BU é soft-deletada | 🔴 | CREV | L | Integridade comprometida. Precisa migration V006 |
| **DT-057** | Sprint 4 | `src/main/resources/db/migration/` | **V004 ausente** — seed data matriz RN10-01 (`resource_action` + `role_resource`). Tabelas vazias desde V001 | 🔴 | CREV, ARCH | M | T-049 sem seed. `RbacAspect` carrega matriz vazia → 403 universal |
| **DT-058** | Sprint 4 | `RbacAspect.java:82-91` | **Strings literais** ("MANAGER_BU", "OPERATOR_BU") em vez de `Role` enum. DT-035 (Sprint 3) nunca resolvido | 🔴 | CREV, PONY, CODE | L | Type-safety zero. Divergência banco↔código quebra RBAC silenciosamente |
| **DT-059** | Sprint 4 | `TenantController.java:44-45` | **Controller bypassa service** — `list()` chama `tenantRepo.findAllPaginated()` direto. `@RequiresPermission` NÃO aplicado | 🔴 | CREV | L | RBAC bypass no endpoint de listagem de tenants |
| **DT-060** | Sprint 4 | `SecurityConfig.java:94-98` | **JWT decoder sem validação de `iss`** — `NimbusJwtDecoder.withJwkSetUri()` só valida assinatura. Token de QUALQUER realm Keycloak é aceito | 🔴 | CODE, SEC | M | Bypass de autoridade JWT. Atacante com realm próprio forja tokens |
| **DT-061** | Sprint 4 | `V003__enable_rls.sql` | **RLS sem `FORCE ROW LEVEL SECURITY`** — PostgreSQL NÃO aplica RLS ao table owner. Flyway cria tabelas como owner → app user = owner → RLS 100% ineficaz. `RLSIsolationTest.java` prova a falha (adiciona FORCE no setup) | 🔴 | SEC | M | Isolamento multi-tenant quebrado. Dados cross-tenant visíveis |
| **DT-062** | Sprint 4 | `SubscriptionService.java:40-68` + `SubscriptionRepository.java:28-34` | **Tenant isolation bypass**: `tenant_id` da URL vs JWT não validado. `findActiveByTenantId()` faz query custom sem `tenantClause()` do `BaseRepository` | 🔴 | CODE, SEC | H | Violação isolamento multi-tenant. Tenant X lê/escreve dados de tenant Y |
| **DT-063** | Sprint 4 | `TenantAwareDataSource.java:62-63` | **Connection leak**: `UUID.fromString()` dentro do try, catch só pega `SQLException`. `IllegalArgumentException` → conexão vaza do pool | 🔴 | CODE, DEBT | M | Pool exhaustion. Multi-tenant quebra para todos |
| **DT-064** | Sprint 4 | `SPRINT-CARD.md:61-63` vs `RbacAspect.java:48-60` | **Matriz RBAC diverge**: SPRINT-CARD diz MANAGER_BU ❌ Dashboard/Tenants/Plans; RbacAspect permite VIEW de TENANT, PLAN, SUBSCRIPTION, DASHBOARD. MANAGER_BU perde acesso na migração → breaking change sem plano de cutover | 🔴 | CR, ARCH | M | Regressão 403 para MANAGER_BU. Sem rollback |
| **DT-065** | Sprint 4 | `TASKS.md:270-330` + `TASKS.md:278` | **Diagrama de dependências referencia tasks erradas**: T-039..T-047 (inexistentes) para M4. `T-046` rotulado como "RbacAspect" mas é "User entity" | 🔴 | CR | L | Desenvolvedor segue tasks fantasmas. Ordem de implementação errada |
| **DT-066** | Sprint 4 | `SPRINT-TEST-SUITE.md:6` vs `:83-90` | **Header 19 cenários, corpo 27** (7+9+6+5=27). SPRINT-CARD também diz 19. Métricas de teste quebradas | 🔴 | CR | L | Sprint reporting errado. Aceitação baseada em métricas falsas |
| **DT-067** | Sprint 4 | `RbacAspect.java` + `TenantContext.java:60-63` | **Sem verificação de escopo por BU** — `RbacAspect` nunca lê `TenantContext.getBusinessUnitIds()`. MANAGER_BU pode editar produtos de BU alheia. RLS filtra por tenant_id, não por business_unit_id | 🔴 | ARCH | H | MANAGER_BU edita qualquer BU. Violação RN10-01 |

### 🟡 Riscos (15 — tratar durante a sprint)

| ID | Sprint Origem | Arquivo/Artefato | Achado | Sev. | Skill | Compl. | Efeito se não tratado |
|:---|:---|:---|:---|:---:|:---:|:---:|:---|
| **DT-068** | Sprint 4 | `pom.xml` (postgresql 42.7.10) | CVE-2026-42198 (CVSS 7.5): SCRAM-SHA-256 DoS. Fix: 42.7.11 | 🟡 | SEC | L | DoS via servidor PostgreSQL malicioso |
| **DT-069** | Sprint 4 | `pom.xml:25` (Flyway 10.22.0→12.11.0) | 2 majors atrás. DT-045 postergado. Migration path: 10→11→12 | 🟡 | SEC, DEBT | H | Migração mais complexa a cada sprint |
| **DT-070** | Sprint 4 | `AuditAspect.java:200-209` | `parseEntityId()` retorna `UUID.randomUUID()` para entrada inválida → entity_id fantasma no audit_log. CREATE operations: entity ID não capturado do retorno | 🟡 | CODE, CREV | L | RN08-01 (100% rastreabilidade) violada |
| **DT-071** | Sprint 4 | `controller/{Audit,Tenant,Subscription,Plan}Controller.java` | 4 controllers com **0% JaCoCo**. Sem `@WebMvcTest`. Só `DashboardController` tem teste | 🟡 | DEBT, CREV | M | RBAC via HTTP nunca testado. Regressão invisível |
| **DT-072** | Sprint 4 | `RLSIsolationTest.java:129` | **RLS real @Disabled** ("RLS FORCE + SingleConnectionDataSource requer refatoração"). Isolamento cross-tenant nunca testado com PostgreSQL real | 🟡 | CREV | H | Falsa segurança. RLS pode estar quebrado sem detecção |
| **DT-073** | Sprint 4 | `BaseRepositoryTest.java:216-251` | `save()` e `update()` **@Disabled** ("Mockito varargs incompatibility"). Métodos centrais do BaseRepository sem cobertura unitária | 🟡 | CREV | M | Regressão em RBAC repositories sem detecção |
| **DT-074** | Sprint 4 | `AuditAspect.java:142-183` | `resolveTableName()` switch não reconhece PERMISSION, ROLE → retorna null. Novas entidades RBAC sem auditoria | 🟡 | CREV | M | Auditoria cega para operações RBAC |
| **DT-075** | Sprint 4 | `GlobalExceptionHandler.java:120-131` | Handler `java.lang.SecurityException` depreciado e inútil (JDK 17+). `TenantIsolationException` → 500 genérico em vez de 403 | 🟡 | CODE | L | Violação de segurança retorna 500 |
| **DT-076** | Sprint 4 | `JwtAuthenticationFilter:80` + `SecurityConfig:58` | JWT decodificado 2×: `BearerTokenAuthenticationFilter` + `JwtAuthenticationFilter`. DT-030 postergado | 🟡 | CODE, PONY | M | Dobro de latência. Custo em JWKS fetch |
| **DT-077** | Sprint 4 | `RbacAspectTest.java:27-143` | Testes instanciam `new RbacAspect()` sem Spring. T-053 injeta repository → todos quebram | 🟡 | CREV | H | Falsa segurança. Testes mockados não validam DB-backed RBAC |
| **DT-078** | Sprint 4 | `SubscriptionService.java:94-100` | **TOCTOU race condition**: `countActive()<=1` e `setStatus("DISCONTINUED")` sem lock → 2 requests zeram planos ativos | 🟡 | CODE | M | Ficar sem plano ativo |
| **DT-079** | Sprint 4 | `AuditService.java:38-41` | `OffsetDateTime.parse()` sem try-catch → `DateTimeParseException` = 500 em vez de 400 | 🟡 | CODE | L | Erro confuso para o cliente |
| **DT-080** | Sprint 4 | `SPECS.md:194-200` vs `PlanController.java` | 2 endpoints Plan não documentados (`GET /admin`, `GET /{id}`). Quebra ADR-06 API Contract First | 🟡 | CR | L | API specs desatualizadas |
| **DT-081** | Sprint 4 | `SPRINT-CARD.md:85` | Referência errada a "Migration V003 com seed data" — V003 é RLS. Seed RBAC deveria ser V006 | 🟡 | CR | L | Colisão de migration. Deploy quebrado |
| **DT-082** | Sprint 4 | `SPRINT-CARD.md:92` | Dependência lista "Sprint 3 (UserRepository)" mas `UserRepository.java` NÃO existe — é deliverable da Sprint 4 | 🟡 | CR | L | Falsa dependência. Confunde planejamento |

### 🔵 Nits (12 — sprints futuras)

| ID | Sprint Origem | Arquivo/Artefato | Achado | Sev. | Skill | Compl. |
|:---|:---|:---|:---|:---:|:---:|:---:|
| **DT-083** | Sprint 4 | `Address.java` + 5 enums (`ProductType`, `TaxRegime`, `UserStatus`, `SubscriptionStatus`, `Recurrence`) | ~211 linhas de código morto. `UserStatus` deveria ser usado pela Sprint 4 | 🔵 | PONY | L |
| **DT-084** | Sprint 4 | `EmailService.java` + `EmailServiceImpl.java` | Interface YAGNI com única implementação. `sendInvite()` nunca chamado. 88 linhas | 🔵 | PONY | L |
| **DT-085** | Sprint 4 | `TenantContext.java:60-71` + `JwtUtils.java:68-85` | `getBusinessUnitIds()` + `getModules()` populados mas **nunca consumidos** por código de produção | 🔵 | PONY | L |
| **DT-086** | Sprint 4 | `repository/rowmapper/*.java` | 4 RowMappers duplicam 6 linhas de campos de auditoria. 24 linhas duplicadas | 🔵 | PONY, JSCPD | L |
| **DT-087** | Sprint 4 | `TenantRepository.java:42-109` + `AuditRepository.java:39-99` | SQL construction duplicado: `findAllPaginated()` e `countFiltered()` replicam JOINs, WHERE, params (~70 linhas) | 🔵 | JSCPD, DEBT | M |
| **DT-088** | Sprint 4 | `SubscriptionService.java:52-58,87-93` | Validação de plano ativo duplicada em `create()` e `changePlan()`. 12 linhas | 🔵 | JSCPD | L |
| **DT-089** | Sprint 4 | `AuditAspect.java:40` | `static final ObjectMapper` ignora o `ObjectMapper` configurado pelo Spring (sem JavaTimeModule) | 🔵 | PONY, CODE | L |
| **DT-090** | Sprint 4 | `common/BaseEntity.java:32-35` | `OffsetDateTime.now()` usa fuso do sistema, não UTC | 🔵 | DEBT | L |
| **DT-091** | Sprint 4 | `TenantController.java:129-135` | `POST /{id}/resend-invite` é stub — retorna 202 sem enviar email | 🔵 | CODE | L |
| **DT-092** | Sprint 4 | `springdoc-openapi 2.8.8→2.8.16` | 8 minors atrás | 🔵 | SEC | L |
| **DT-093** | Sprint 4 | `SecurityConfig.java:106-107` | CORS origins hardcoded. Sem suporte a env var | 🔵 | CODE | L |
| **DT-094** | Sprint 4 | `README.md:1` | Ainda "TODO". Zero documentação | 🔵 | DEBT | M |

---

## Achados por Skill

### code-reviewer (CREV — 33 achados → 8 únicos após dedup)
Foco: SOLID violations (DT-048 DIP, DT-049 OCP, DT-050 SRP, DT-051 ISP), code smells (DT-054 hardcoded matrix, DT-055 magic strings, DT-060 disabled tests), architecture (DT-063 controller bypassing service = RBAC bypass, DT-064 missing RBAC integration tests), security (DT-069 string-based roles, DT-070 AuditAspect won't recognize RBAC entities), tests (DT-074 missing RBAC integration, DT-075 RLS tests disabled).

### caveman-review (CR — 20 achados → 8 únicos após dedup)
Foco: divergência matriz RBAC SPRINT-CARD×código (DT-064), diagrama TASKS.md com tasks erradas (DT-065), header SPRINT-TEST-SUITE 19 vs 27 (DT-066), 2 endpoints Plan não documentados (DT-080), migration V003 vs V006 confusão (DT-081), UserRepository citado como pré-requisito inexistente (DT-082).

### superpowers:brainstorming (ARCH — 14 achados → 3 únicos após dedup)
Foco: conflito arquitetural JWT×banco roles (DT-050), cache TTL viola RN11-03 (DT-051), sem verificação escopo BU (DT-067). **Recomendação-chave:** abandonar cache TTL; carregar matriz 32 linhas com query indexada (<1ms).

### ponytail-review (PONY — 17 achados → 4 únicos após dedup)
Foco: 211 linhas código morto (Address + 5 enums), EmailService YAGNI 88 linhas, TenantContext BU/modules nunca consumidos. **net: -296 lines possible.** Zero bloqueantes.

### ponytail-debt (PDBT — 2 achados)
**2 markers, 0 with no trigger.** Ambos em `RbacAspect.java:33,46`: "matriz hardcoded até Sprint 4" e "ceiling = Sprint 4 carrega do banco". Ambos com trigger definido → baixo risco de apodrecimento.

### code-review (CODE — 22 achados → 6 únicos após dedup)
Foco: 6 críticos — JWT sem validação `iss` (DT-060), tenant isolation bypass subscription (DT-062), connection leak TenantAwareDataSource (DT-063), AuditAspect entity_id fantasma (DT-070), DateTimeParseException → 500 (DT-079).

### jscpd + dry-refactoring (JSCPD — 18 achados → 3 únicos após dedup)
**1.64% duplicação** (jscpd) + 8 clones estruturais manuais = **~2.6% total**. Projeto com baixa duplicação. Principais: SQL filter-building duplicado (AuditRepository + TenantRepository, ~70 linhas), RowMapper audit fields (24 linhas), validação plano ativo duplicada (12 linhas).

### tech-debt (DEBT — 47 achados → 2 únicos após dedup)
Priorização quantitativa: 2 críticos (score≥40), 13 warnings (25-39), 32 baixo (<25). Média: 22.3. Dependências comentadas (Caffeine, REST Assured) bloqueiam Sprint 4. 4 controllers com 0% cobertura.

### security-review (SEC — 12 achados → 4 únicos após dedup)
Foco: RLS sem FORCE (DT-061 — isolamento quebrado), AuditRepository sem tenant_id filter (cross-tenant leak), Caffeine/REST Assured ausentes, SMTP sem TLS, JWKS HTTP default. Licenças 100% compatíveis (Apache 2.0/BSD/MIT). Sem GPL/AGPL.

---

## Recomendações Prioritárias

### 🔴 Bloqueantes (20 — Frente 0, ANTES de iniciar features)

| ID (TASKS.md) | DT-XXX | Sprint Origem | Ação Corretiva | Estimativa |
|:---|:---|:---|:---|:---:|
| T-096.DT-048 | DT-048 | Sprint 4 | Adicionar `spring-boot-starter-cache` + `caffeine:3.2.4` ao pom.xml. Criar `CacheConfig.java` | 1h |
| T-097.DT-049 | DT-049 | Sprint 4 | Adicionar `rest-assured:5.5.7` + `rest-assured-spring-mock-mvc:5.5.7` (test scope) | 30min |
| T-098.DT-050 | DT-050 | Sprint 4 | Definir estratégia merge JWT×banco: `RbacAspect` consulta `user_permission` como fonte primária OU JWT inclui permissões do banco | 3h |
| T-099.DT-051 | DT-051 | Sprint 4 | Abandonar cache TTL 5min. Carregar matriz via `findAll()` com query indexada (<1ms). OU embed version hash no JWT | 2h |
| T-100.DT-052 | DT-052 | Sprint 4 | Refatorar `RbacAspect`: injetar `PermissionService`, remover Sets hardcoded, carregar do banco | 4h |
| T-101.DT-053 | DT-053 | Sprint 4 | Criar `User.java` entity + `UserRepository.java` (estende `BaseRepository<User>`). Índice parcial email único | 2h |
| T-102.DT-054 | DT-054 | Sprint 4 | Criar `ResourceAction.java` + `RoleResource.java` entities | 1h |
| T-103.DT-055 | DT-055 | Sprint 4 | Criar `BusinessUnit.java` entity (estende `BaseEntity`) | 1h |
| T-104.DT-056 | DT-056 | Sprint 4 | Criar migration V006: `ALTER TABLE user_permission ADD CONSTRAINT fk_up_bu FOREIGN KEY (business_unit_id) REFERENCES business_unit(id)` | 30min |
| T-105.DT-057 | DT-057 | Sprint 4 | Criar migration V004: `INSERT INTO resource_action` (8×4) + `INSERT INTO role_resource` (matriz RN10-01) + **U004 rollback** | 1h |
| T-106.DT-058 | DT-058 | Sprint 4 | Migrar `RbacAspect` de strings para `Role` enum. DT-035 do backlog | 1h |
| T-107.DT-059 | DT-059 | Sprint 4 | Refatorar `TenantController.list()` para usar `TenantService` (não `TenantRepository` direto). Aplicar `@RequiresPermission` | 1h |
| T-108.DT-060 | DT-060 | Sprint 4 | Adicionar `JwtValidators.createDefaultWithIssuer(issuerUri)` no `NimbusJwtDecoder` | 1h |
| T-109.DT-061 | DT-061 | Sprint 4 | Adicionar `ALTER TABLE ... FORCE ROW LEVEL SECURITY` nas 4 tabelas RLS do V003. OU criar role não-owner para app | 2h |
| T-110.DT-062 | DT-062 | Sprint 4 | Validar `tenant_id` da URL contra `TenantContext.getTenantId()`. Adicionar `tenantClause()` no `SubscriptionRepository.findActiveByTenantId()` | 2h |
| T-111.DT-063 | DT-063 | Sprint 4 | Corrigir `TenantAwareDataSource`: wrap `UUID.fromString()` em try-catch próprio; fechar conexão antes de relançar | 1h |
| T-112.DT-064 | DT-064 | Sprint 4 | Alinhar matriz RBAC: decidir SPRINT-CARD vs RbacAspect. Documentar breaking change e plano de cutover | 1h |
| T-113.DT-065 | DT-065 | Sprint 4 | Corrigir diagrama TASKS.md §3: T-039..T-047 → T-046..T-056 | 30min |
| T-114.DT-066 | DT-066 | Sprint 4 | Corrigir header SPRINT-TEST-SUITE.md: 19→27. Atualizar SPRINT-CARD.md | 15min |
| T-115.DT-067 | DT-067 | Sprint 4 | Adicionar verificação de escopo BU no `RbacAspect` ou `PermissionService`. `@RequiresBusinessUnitAccess` ou check manual nos services | 3h |

### 🟡 Recomendados (15 — incluir no backlog da sprint)

| ID (TASKS.md) | DT-XXX | Sprint Origem | Ação Corretiva | Estimativa |
|:---|:---|:---|:---|:---:|
| T-116.DT-068 | DT-068 | Sprint 4 | Atualizar PostgreSQL driver 42.7.10→42.7.11 (DoS CVE) | 30min |
| T-117.DT-069 | DT-069 | Sprint 4 | Migrar Flyway 10.22.0→12.11.0 (10→11→12, testar migrations) | 4h |
| T-118.DT-070 | DT-070 | Sprint 4 | Corrigir `AuditAspect.parseEntityId()`: retornar null (não random UUID). Capturar entity ID do retorno em CREATE | 1h |
| T-119.DT-071 | DT-071 | Sprint 4 | Criar `@WebMvcTest` para 4 controllers (MockMvc + JWT simulado) | 4h |
| T-120.DT-072 | DT-072 | Sprint 4 | Corrigir `RLSIsolationTest`: remover `@Disabled`, mockar SingleConnectionDataSource | 3h |
| T-121.DT-073 | DT-073 | Sprint 4 | Corrigir `BaseRepositoryTest.save/update`: remover `@Disabled`, usar argumentCaptor | 2h |
| T-122.DT-074 | DT-074 | Sprint 4 | Adicionar PERMISSION, ROLE no `resolveTableName()` do `AuditAspect` | 30min |
| T-123.DT-075 | DT-075 | Sprint 4 | Adicionar handler `TenantIsolationException` (403). Remover handler `SecurityException` | 30min |
| T-124.DT-076 | DT-076 | Sprint 4 | Migrar para `Converter<Jwt, AbstractAuthenticationToken>` customizado (DT-030) | 4h |
| T-125.DT-077 | DT-077 | Sprint 4 | Reescrever `RbacAspectTest` com `@SpringBootTest` + `@MockBean PermissionService` | 3h |
| T-126.DT-078 | DT-078 | Sprint 4 | Adicionar `SELECT ... FOR UPDATE` no `PlanService.deactivate()` | 1h |
| T-127.DT-079 | DT-079 | Sprint 4 | Adicionar try-catch `DateTimeParseException` → 400 no `AuditService` | 15min |
| T-128.DT-080 | DT-080 | Sprint 4 | Documentar `GET /plans/admin` e `GET /plans/{id}` na SPECS.md §4.1 | 15min |
| T-129.DT-081 | DT-081 | Sprint 4 | Corrigir SPRINT-CARD.md:85 "V003" → "V006 (ou próxima disponível)" | 5min |
| T-130.DT-082 | DT-082 | Sprint 4 | Remover "Sprint 3 (UserRepository)" das dependências no SPRINT-CARD.md | 5min |

### 🔵 Desejáveis (12 — se houver capacidade)

| ID (TASKS.md) | DT-XXX | Sprint Origem | Ação Corretiva | Estimativa |
|:---|:---|:---|:---|:---:|
| T-131.DT-083 | DT-083 | Sprint 4 | Remover Address + 5 enums sem referência (~211 linhas). OU adotar UserStatus na Sprint 4 | 15min |
| T-132.DT-084 | DT-084 | Sprint 4 | Remover `EmailService` interface (YAGNI). Usar `EmailServiceImpl` direto | 5min |
| T-133.DT-085 | DT-085 | Sprint 4 | Remover `getBusinessUnitIds()`/`getModules()` se continuarem sem consumidor | 5min |
| T-134.DT-086 | DT-086 | Sprint 4 | Extrair `BaseRowMapper.setAuditFields(entity, rs)` (24 linhas duplicadas) | 15min |
| T-135.DT-087 | DT-087 | Sprint 4 | Extrair `buildFilterQuery()` helper nos repositories | 30min |
| T-136.DT-088 | DT-088 | Sprint 4 | Extrair `getActivePlan(UUID)` helper no `SubscriptionService` | 15min |
| T-137.DT-089 | DT-089 | Sprint 4 | Substituir `static ObjectMapper` por injeção Spring no `AuditAspect` | 15min |
| T-138.DT-090 | DT-090 | Sprint 4 | `OffsetDateTime.now(ZoneOffset.UTC)` no `BaseEntity` | 10min |
| T-139.DT-091 | DT-091 | Sprint 4 | Remover endpoint stub `POST /{id}/resend-invite` | 5min |
| T-140.DT-092 | DT-092 | Sprint 4 | Atualizar springdoc-openapi 2.8.8→2.8.16 | 15min |
| T-141.DT-093 | DT-093 | Sprint 4 | Externalizar CORS origins para `application.yml` | 15min |
| T-142.DT-094 | DT-094 | Sprint 4 | Escrever `README.md` (quickstart, build, test, deploy) | 2h |

---

## Débitos Técnicos Elegíveis para Sprints Futuras

> Débitos postergados por decisão do time em 17/07/2026. Reavaliar no `PROMPT-GENERATE-IDENTIFIED-TECHNICAL-DEBT` da Sprint 5.

| DT-XXX | Sprint Origem | Descrição | Severidade | Skill | Compl. | Sprint Sugerida | Justificativa |
|:---|:---|:---|:---:|:---:|:---:|:---|:---|
| DT-023 | Sprint 3 | Migrar paginação offset→keyset | 🟡 | BACKLOG | L | Sprint 5 | Performance — não bloqueia features |
| DT-031 | Sprint 3 | Reduzir Checkstyle maxAllowedViolations 300→0 | 🔵 | BACKLOG | L | Sprint 5-7 | Qualidade contínua |
| DT-034 | Sprint 3 | Remover Address.java | 🔵 | BACKLOG | L | Sprint 6 | Código morto — sem pressa |
| DT-044 | Sprint 3 | Criar logback-spring.xml | 🔵 | BACKLOG | M | Sprint 7 | Infra — prioridade baixa |
| DT-045 | Sprint 3 | Migrar Flyway 10→12 | 🟡 | BACKLOG | H | Sprint 5 | Complexo — requer janela dedicada |
| DT-083 | Sprint 4 | Código morto: Address + 5 enums (~211 linhas) | 🔵 | PONY | L | Sprint 5 | Cosmético |
| DT-084 | Sprint 4 | EmailService interface YAGNI (88 linhas) | 🔵 | PONY | L | Sprint 5 | Cosmético |
| DT-085 | Sprint 4 | TenantContext BU/modules não consumidos | 🔵 | PONY | L | Sprint 5 | Cosmético |
| DT-086 | Sprint 4 | RowMappers duplicam 6 linhas auditoria (×4) | 🔵 | PONY, JSCPD | L | Sprint 5 | Baixo impacto |
| DT-087 | Sprint 4 | SQL construction duplicado (~70 linhas) | 🔵 | JSCPD, DEBT | M | Sprint 5 | Refactor — requer testes |
| DT-088 | Sprint 4 | Validação plano ativo duplicada (12 linhas) | 🔵 | JSCPD | L | Sprint 5 | Trivial |
| DT-089 | Sprint 4 | ObjectMapper não-Spring no AuditAspect | 🔵 | PONY, CODE | L | Sprint 5 | Cosmético |
| DT-090 | Sprint 4 | OffsetDateTime sem UTC no BaseEntity | 🔵 | DEBT | L | Sprint 5 | Cosmético |
| DT-091 | Sprint 4 | Endpoint stub resend-invite | 🔵 | CODE | L | Sprint 5 | Substituído pelo UserController |
| DT-092 | Sprint 4 | springdoc 2.8.8→2.8.16 (8 minors atrás) | 🔵 | SEC | L | Sprint 5 | Não crítico |
| DT-093 | Sprint 4 | CORS origins hardcoded | 🔵 | CODE | L | Sprint 5 | Funciona para os ambientes atuais |
| DT-094 | Sprint 4 | README.md vazio | 🔵 | DEBT | M | Sprint 7 | Documentação — final da fase |

---

## Decisão do Time

> **Decisão do time em 17/07/2026.**

| ID | Decisão | Sprint alvo | Justificativa |
|:---|:---|:---:|:---|
| **DT-048 a DT-067** (20) | ✅ **Tratar agora** (Frente 0) | Sprint 4 | Bloqueantes — impedem início das features RBAC |
| **DT-068 a DT-082** (15) | ✅ **Tratar agora** (durante sprint) | Sprint 4 | Riscos que degradam qualidade e segurança |
| DT-030 (backlog) | ✅ **Tratar agora** | Sprint 4 | Incluído via T-124.DT-076 (JWT consolidation) |
| DT-035 (backlog) | ✅ **Tratar agora** | Sprint 4 | Incluído via T-106.DT-058 (Role enum) |
| DT-042 (backlog) | ✅ **Tratar agora** | Sprint 4 | Docker-compose pendente — incluído nas Recomendadas |
| DT-043 (backlog) | ✅ **Tratar agora** | Sprint 4 | Seed scripts pendentes — incluído nas Recomendadas |
| **DT-083 a DT-094** (12) | ⬜ **Postergar** | Sprint 5+ | Nits — baixo impacto, revisar na Sprint 5 |
| DT-023, DT-031, DT-034, DT-044, DT-045 (backlog) | ⬜ **Postergar** | Sprint 5+ | Backlog — manter no radar, reavaliar na Sprint 5 |

---

## Análise de Impacto nos Documentos

### Impacto nos Documentos-Mestre

| ID | Documento | Impacto | Ação necessária |
|:---|:---|:---|:---|
| DT-048, DT-049, DT-068 | PRD.md | Novas dependências (Caffeine, REST Assured) + version bumps | Atualizar stack |
| DT-050, DT-051, DT-052, DT-058, DT-067 | ARCHITECTURE.md | Refatoração completa pipeline RBAC: JWT→TenantContext→DB | Novo ADR ou revisão ADR-L06 |
| DT-061 | ARCHITECTURE.md | ADR-L07 (RLS) incorreto — falta FORCE | Corrigir ADR-L07 |
| DT-053..DT-055 | SPECS.md | Novas entities (User, ResourceAction, RoleResource, BusinessUnit) em §5 | Atualizar modelo de dados |
| DT-057, DT-056 | ARCHITECTURE.md | Novas migrations V004 (seed RBAC) + V006 (FK) | Registrar |
| DT-064 | SPECS.md, SPRINT-CARD.md | Matriz RN10-01 unificada | Alinhar definição |
| DT-065 | TASKS.md | Diagrama dependências corrigido | Regenerar §3 |
| DT-066 | SPRINT-TEST-SUITE.md, SPRINT-CARD.md | Header 19→27 | Corrigir métricas |
| DT-080 | SPECS.md | 2 endpoints Plan não documentados | Adicionar §4.1 |

---

## Verificação Pós-Consolidação

- [ ] BUILD: `mvn clean verify` passa com Caffeine + REST Assured + Spring Boot 3.5.15
- [ ] 4 novas entities compilam (User, ResourceAction, RoleResource, BusinessUnit)
- [ ] Migration V004 (seed RBAC) + U004 (rollback) executam sem erro
- [ ] Migration V006 (FK user_permission) executa sem erro
- [ ] `RbacAspect` injeta `PermissionService` — sem Sets hardcoded
- [ ] `RbacAspect` usa `Role` enum, não strings
- [ ] `PermissionService.assignRole()` invalida estado imediatamente (RN11-03)
- [ ] `NimbusJwtDecoder` valida `iss` claim
- [ ] V003 com `FORCE ROW LEVEL SECURITY` nas 4 tabelas
- [ ] `SubscriptionService` valida `tenant_id` URL contra JWT
- [ ] `TenantController.list()` passa pelo service (com `@RequiresPermission`)
- [ ] `AuditAspect.resolveTableName()` reconhece PERMISSION, ROLE
- [ ] `AuditAspect.parseEntityId()` retorna null (não UUID aleatório)
- [ ] Matriz RBAC unificada (SPRINT-CARD = código = seed)
- [ ] TASKS.md diagrama dependências corrigido
- [ ] SPRINT-TEST-SUITE header corrigido (27)
- [ ] `TenantAwareDataSource` connection leak corrigido
- [ ] 47 débitos catalogados + 9 backlog = 56 total

---

🤖 *Análise gerada em 2026-07-17. 47 achados consolidados a partir de 9 skills paralelas (code-reviewer, caveman-review, superpowers:brainstorming, ponytail-review, ponytail-debt, code-review, jscpd+dry-refactoring, tech-debt, security-review) + 9 débitos backlog Sprint 3. Prompt v3.1.*
