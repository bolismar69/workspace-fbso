# TECHNICAL-REFERENCE.md — Referencia Tecnica para Agentes

- **Microservico:** `ms-fbso-platform-admin`
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Caffeine 3.2.4 + Keycloak 26
- **Versao:** 1.2
- **Data:** 23 de Julho de 2026
- **Situação implementação:** Em Execução — Sprint 5 Frentes 0-1-2-3a concluídas ✅ (36/40, 90%). Sprint 6 Frente 0 concluída ✅ (4/4 — BusinessUnit, ProductService, CnpjValidator alfanumérico IN RFB 2.119/2022). 26 endpoints REST. 4 features F04-01 a F04-04 entregues. 261 testes (0 failures).
- **Objetivo:** Guia rapido para agentes de desenvolvimento. Compacto, escaneavel, comandos exatos.

---

## 1. Setup Local

### 1.1 Pre-requisitos

| Ferramenta | Versao | Verificacao |
|:---|:---|:---|
| Java | 25 | `java -version` |
| Docker | 24+ | `docker --version` |
| Docker Compose | v2 | `docker compose version` |
| Maven | 3.9+ (via wrapper) | `./mvnw -v` |

### 1.2 Iniciar Infraestrutura

```bash
docker compose up -d
```

| Servico | Porta | Credenciais Dev |
|:---|:---:|:---|
| PostgreSQL 17 | 5432 | `fbso_admin` / `fbso_admin` (db: `fbso_platform`) |
| Keycloak 26 | 8081 | `admin` / `admin` (Admin Console) |
| MailHog SMTP | 1025 | — (Web UI: 8025) |

### 1.3 Credenciais de Teste (Keycloak)

| Usuario | Senha | Role | Tenant |
|:---|:---|:---|:---|
| `admin@fbso.org` | `admin` | ADMIN_TENANT | Qualquer |

Realm: `fbso-platform` | Client: `fbso-platform-admin`

### 1.4 Build e Run

```bash
./mvnw clean install          # Compila + roda unit tests
./mvnw spring-boot:run        # Roda em dev (porta 8080)
./mvnw verify                 # Compila + unit + integration tests (Docker necessario)
```

### 1.5 Profiles

| Profile | Ativacao | Configuracao |
|:---|:---|:---|
| `dev` | Default | DEBUG logging, Flyway clean habilitado |
| `staging` | `-Dspring.profiles.active=staging` | ENV vars, Flyway clean-desabilitado |
| `prod` | `-Dspring.profiles.active=prod` | Apenas `/actuator/health`, WARN logging |

---

## 2. Estrutura de Pacotes

```
com.fbso.platform.admin/
├── FbsoPlatformAdminApplication.java     # Entry point (@EnableAsync)
├── config/                                # SecurityConfig, CacheConfig, DataSourceConfig, TenantAwareDataSource
├── controller/                            # 11 controllers REST
├── dto/request/                           # 10 request DTOs (records Java)
├── dto/response/                          # 12 response DTOs (records Java)
├── entity/                                # 9 entidades (BaseEntity)
├── enums/                                 # 8 enums (TenantStatus, Role, etc.)
├── exception/                             # BusinessException + GlobalExceptionHandler (RFC 7807)
├── repository/                            # 9 repositories + common/BaseRepository
├── repository/rowmapper/                  # 6 RowMappers
├── security/                              # JwtAuthenticationFilter, TenantContext
├── security/annotation/                   # @RequiresPermission, @Auditable
├── security/aspect/                       # RbacAspect, AuditAspect
├── service/                               # 9 services + EmailServiceImpl
├── common/                                # BaseEntity, Address
└── utils/                                 # JwtUtils, CnpjValidator
```

**Detalhes:** Ver ARCHITECTURE.md Secao 2

---

## 3. Padroes de Codigo

### 3.1 Entidade

```java
public class Tenant extends BaseEntity {
    private String nameCorporate;
    private TenantStatus status;
    // ... campos de dominio
}
```

- Extende `BaseEntity` (created_dt, updated_dt, created_by, updated_by, deleted_dt, deleted_by)
- Mapeado por `RowMapper<T>` no repository (nao usa JPA)

### 3.2 Repository

```java
@Repository
public class TenantRepository extends BaseRepository<Tenant> {
    public TenantRepository(JdbcTemplate jdbc) {
        super(jdbc, "tenant", new TenantRowMapper(), true); // hasTenantColumn = true
    }
    // queries customizadas...
}
```

- `BaseRepository` injeta automaticamente `WHERE deleted_dt IS NULL` + `WHERE tenant_id = ?`
- `hasTenantColumn = false` para tabelas sem tenant_id (plan, resource_action, role_resource)
- `save(T)` e `update(T)` preenchem created_by/updated_by automaticamente

### 3.3 Service

```java
@Service
public class TenantService {
    @Auditable(entityType = "TENANT", action = "CREATED")
    public TenantResponse create(TenantCreateRequest req) { ... }

    @Auditable(entityType = "TENANT", action = "SUSPENDED")
    public TenantResponse suspend(String id, String reason) { ... }
}
```

- `@Auditable` em mutations (create, update, suspend, deactivate)
- `@Transactional` para operacoes atomicas (change-plan, etc.)
- Excecoes de dominio: `BusinessException` (422), `DuplicateCnpjException` (409), `TenantNotFoundException` (404)

### 3.4 Controller

```java
@RestController
@RequestMapping("/api/v1/tenants")
public class TenantController {
    @GetMapping
    @RequiresPermission(resource = "TENANT", action = "view")
    public ResponseEntity<Page<TenantResponse>> list(...) { ... }

    @PostMapping
    @RequiresPermission(resource = "TENANT", action = "create")
    public ResponseEntity<TenantResponse> create(@Valid @RequestBody TenantCreateRequest req) { ... }
}
```

- `@RequiresPermission(resource, action)` em TODOS os endpoints
- `@Valid` + Bean Validation nos request DTOs
- Retorna `ResponseEntity<T>` com status HTTP correto

### 3.5 DTO (Request)

```java
public record TenantCreateRequest(
    @NotBlank String nameCorporate,
    @NotBlank TenantSegment segment,
    String nameFantasy
) {}
```

- Records Java com anotacoes Bean Validation
- Nomes em camelCase no JSON

### 3.6 DTO (Response)

```java
public record TenantResponse(
    UUID id,
    String nameCorporate,
    TenantStatus status,
    OffsetDateTime createdDt
) {}
```

- ISO 8601 para datas (`spring.jackson.serialization.write-dates-as-timestamps: false`)
- Nulls omitidos (`spring.jackson.default-property-inclusion: non_null`)

---

## 4. Pipeline de Seguranca (por Request)

```
Request HTTP
  |
  +-- [Order 1] OAuth2 Login Chain (/auth/**, /login, /oauth2/**)
  |     +-- Keycloak Authorization Code Flow + sessao
  |
  +-- [Order 2] API Chain (todos os outros endpoints)
        |
        +-- 1. JwtAuthenticationFilter
        |     Valida RS256 via JWKS
        |     Extrai claims: tenant_id, user_id, roles, business_unit_ids, modules
        |     Seta TenantContext (ThreadLocal)
        |     Seta app.current_tenant_id (SET LOCAL via TenantAwareDataSource)
        |     -> 401 se invalido/expirado
        |
        +-- 2. RbacAspect
        |     Le @RequiresPermission(resource, action)
        |     Consulta PermissionService -> user_permission + role_resource (DB-backed)
        |     Valida contra matriz RN10-01
        |     -> 403 se negado
        |
        +-- 3. Controller -> Service
        |     Valida DTO (@Valid + Bean Validation)
        |     Executa logica de negocio
        |     -> 400 se validacao, 422 se regra de negocio violada
        |
        +-- 4. BaseRepository
        |     Injeta AND deleted_dt IS NULL
        |     Injeta AND tenant_id = ? (se hasTenantColumn)
        |     Preenche created_by/updated_by
        |
        +-- 5. PostgreSQL RLS
        |     FORCE ROW LEVEL SECURITY
        |     tenant_id = current_setting('app.current_tenant_id')::UUID
        |     Recusa queries sem filtro mesmo do table owner
        |
        +-- 6. AuditAspect (pos-execucao)
              Intercepta @Auditable
              Captura snapshot antes/depois
              Grava em audit_log (ASSINCRONO via ThreadPoolTaskExecutor)
```

**Detalhes:** Ver ARCHITECTURE.md Secao 4

---

## 5. Padrao de Erros RFC 7807

Todas as respostas de erro seguem [RFC 7807 (Problem Details)](https://www.rfc-editor.org/rfc/rfc7807):

```json
{
  "type": "https://api.fbso.org/errors/<error-code>",
  "title": "<mensagem amigavel em PT-BR>",
  "status": <codigo HTTP>,
  "detail": "<detalhe opcional>"
}
```

### 5.1 Codigo HTTP x Excecao

| HTTP | Excecao | Exemplo |
|:---:|:---|:---|
| **400** | `MethodArgumentNotValidException` | Bean Validation falhou |
| **401** | `JwtException`, `SecurityException` | Token invalido/expirado |
| **403** | `AccessDeniedException`, `TenantIsolationException`, `PermissionDeniedException` | Acesso negado (RBAC ou tenant isolation) |
| **404** | `TenantNotFoundException`, `UserNotFoundException` | Recurso nao encontrado |
| **409** | `DuplicateCnpjException`, `DuplicateEmailException` | Conflito (duplicado) |
| **422** | `BusinessException`, `SelfDeactivationException`, `InvalidStatusTransitionException`, `PlanHasActiveSubscribersException` | Regra de negocio violada |
| **500** | `Exception` (generico) | Erro interno (SEM stack trace na resposta) |

### 5.2 Erro de Validacao (400)

```json
{
  "type": "https://api.fbso.org/errors/validation-error",
  "title": "Erro de validacao",
  "status": 400,
  "errors": [
    {"field": "nameCorporate", "message": "Razao social e obrigatoria"},
    {"field": "price", "message": "Preco deve ser maior que zero"}
  ]
}
```

### 5.3 403 Padronizado (RBAC)

```json
{
  "type": "https://api.fbso.org/errors/access-denied",
  "title": "Acesso negado",
  "status": 403,
  "detail": "Voce nao tem permissao para acessar esta area."
}
```

**Regra:** 403 NUNCA expoe detalhes tecnicos, caminhos internos ou existencia do recurso (RN12-01, RN12-02).

---

## 6. Padroes de Teste

### 6.1 Organizacao

```
src/test/java/com/fbso/platform/admin/
├── unit/
│   ├── controller/          # MockMvc + JWT simulado
│   ├── service/             # Mockito mocks de repositories
│   ├── repository/          # Mocks de JdbcTemplate
│   ├── security/            # RbacAspectTest, AuditAspectTest
│   ├── entity/              # Testes de dominio
│   ├── exception/           # GlobalExceptionHandlerTest
│   └── config/              # TenantAwareDataSourceTest
├── integration/
│   ├── BaseIntegrationTest.java   # Testcontainers PostgreSQL 17
│   ├── repository/          # TenantServiceIT, DashboardRepositoryIT
│   └── security/            # RbacAspectIntegrationTest, RLSIsolationTest
└── security/
    ├── JwtAuthenticationFilterTest.java
    └── RbacMatrixValidationTest.java
```

### 6.2 Comandos

```bash
./mvnw test                              # Unit tests apenas
./mvnw verify                            # Unit + Integration (Docker necessario)
./mvnw test -pl . -Dtest="*Test"         # Todos os unit tests
./mvnw verify -Dit.test="*IT"            # Todos os integration tests
```

### 6.3 Convencoes

| Aspecto | Convencao |
|:---|:---|
| Naming | `*Test.java` (unit), `*IT.java` (integration) |
| Unit test path | `**/unit/**/*Test.java` |
| IT test path | `**/integration/**/*IT.java` |
| Security test path | `**/security/**/*Test.java` |
| Cobertura min | 85% linhas, 70% branches (JaCoCo) |
| Framework | JUnit 5 + Mockito (unit), Testcontainers (IT), REST Assured (API) |

### 6.4 Base de Integracao

```java
@SpringBootTest
@Testcontainers
@TestPropertySource(properties = {
    "spring.datasource.url=jdbc:tc:postgresql:17-alpine:///fbso_platform"
})
public abstract class BaseIntegrationTest {
    // Container PostgreSQL 17 estatico (compartilhado entre ITs)
    // TenantContext populado no @BeforeEach
}
```

---

## 7. Convencoes de Git

### 7.1 Branches

| Tipo | Prefixo | Exemplo |
|:---|:---|:---|
| Feature | `feature/` | `feature/T-057-keycloak-realm` |
| Bugfix | `bugfix/` | `bugfix/DT-070-audit-entity-id` |
| Hotfix | `hotfix/` | `hotfix/cve-2026-42198-pg-driver` |
| Chore | `chore/` | `chore/update-flyway-12.11.0` |

### 7.2 Commits (Conventional Commits)

```
<tipo>(<escopo>): <descricao curta>

[corpo opcional]
[footer opcional]
```

| Tipo | Uso |
|:---|:---|
| `feat` | Nova feature (endpoint, service, entity) |
| `fix` | Correcao de bug ou defeito tecnico |
| `test` | Adicao/correcao de testes |
| `chore` | Dependencias, configs, refactor sem mudanca funcional |
| `docs` | Documentacao |

Exemplos:
```
feat(tenant): implementar suspend e reativar endpoints
fix(audit): retornar null para entity_id invalido em parseEntityId
test(rbac): adicionar testes parametrizados para matriz RN10-01
chore(deps): bump postgresql driver 42.7.10 para 42.7.11
```

### 7.3 Pull Requests

- Titulo: resumo da mudanca (50 chars max)
- Corpo: o que mudou + por que + evidencia (testes passando)
- Referenciar task: `Closes T-057` ou `Fixes DT-070`

---

## 8. Variaveis de Ambiente

| Variavel | Default (dev) | Descricao |
|:---|:---|:---|
| `DATASOURCE_URL` | `jdbc:postgresql://localhost:5432/fbso_platform` | URL do PostgreSQL |
| `DATASOURCE_USERNAME` | `fbso_admin` | Usuario do PostgreSQL |
| `DATASOURCE_PASSWORD` | `fbso_admin` | Senha do PostgreSQL |
| `KEYCLOAK_JWKS_URI` | `http://localhost:8081/realms/fbso-platform/protocol/openid-connect/certs` | JWKS endpoint |
| `KEYCLOAK_ISSUER_URI` | `http://localhost:8081/realms/fbso-platform` | JWT issuer |
| `KEYCLOAK_CLIENT_ID` | `fbso-platform-admin` | Client ID no Keycloak |
| `KEYCLOAK_CLIENT_SECRET` | `changeme` | Client secret |
| `SMTP_HOST` | `localhost` | Servidor SMTP |
| `SMTP_PORT` | `1025` | Porta SMTP (MailHog) |
| `SERVER_PORT` | `8080` | Porta do servidor |
| `LOG_LEVEL` | `INFO` | Nivel de logging |

---

## 9. Endpoints Implementados (26/37)

### Dashboard Admin (5)
- `GET /api/v1/dashboard/admin/summary`
- `GET /api/v1/dashboard/admin/evolution?period=`
- `GET /api/v1/dashboard/admin/accounts-by-status`
- `GET /api/v1/dashboard/admin/accounts-by-plan`
- `GET /api/v1/dashboard/admin/alerts`

### Tenants (4)
- `GET /api/v1/tenants?page=&size=&search=`
- `GET /api/v1/tenants/{id}`
- `POST /api/v1/tenants`
- `GET /api/v1/tenants/me` (App Switcher — F04-04)

### Planos (4)
- `GET /api/v1/plans`
- `POST /api/v1/plans`
- `PATCH /api/v1/plans/{id}`
- `POST /api/v1/plans/{id}/deactivate`

### Subscriptions (2)
- `GET /api/v1/tenants/{tid}/subscriptions`
- `POST /api/v1/tenants/{tid}/subscriptions`

### Audit (1)
- `GET /api/v1/audit?start_date=&end_date=&action=&page=&size=`

### Users (2)
- `GET /api/v1/users`
- `POST /api/v1/users`

### Auth — Portal Cliente (4) 🆕 Sprint 5
- `POST /api/v1/auth/login` (proxy Keycloak)
- `POST /api/v1/auth/forgot-password`
- `POST /api/v1/auth/reset-password`
- `GET /api/v1/auth/me` (usuário logado + módulos)

### Onboarding (4) 🆕 Sprint 5
- `GET /api/v1/onboarding/status`
- `PATCH /api/v1/onboarding/step-1`
- `POST /api/v1/onboarding/step-2`
- `POST /api/v1/onboarding/complete`

### Dashboard Cliente (1) 🆕 Sprint 5
- `GET /api/v1/dashboard/client/summary?module_id=`

### Health (1)
- `GET /actuator/health`

**Lista completa de 37 endpoints:** Ver SPECS.md Secao 4.1

---

## 10. Referencias Rapidas

| Documento | Conteudo | Tamanho |
|:---|:---|:---:|
| **PRD.md** | Requisitos de negocio, user stories, RNs | 587 linhas |
| **ARCHITECTURE.md** | C4, design detalhado, ADRs, pipeline seguranca | 892 linhas |
| **SPECS.md** | Contrato API, 37 endpoints, 51 RNs, validacoes | 439 linhas |
| **TASKS.md** | Plano de tarefas, 176 tasks, criterios DONE | 558 linhas |
| **TEST_PLAN.md** | 176 cenarios, piramide de testes, seguranca | 418+ linhas |
| **SECURITY.md** | Threat model STRIDE, OWASP Top 10, pipeline DevSecOps | ~500 linhas |
| **TECHNICAL-REFERENCE.md** | Este documento (guia rapido) | ~300 linhas |

### Quanto consultar cada documento

| Necessidade | Documento |
|:---|:---|
| Como rodar o projeto | Este documento (Secao 1) |
| Padrao de codigo para nova feature | Este documento (Secao 3) |
| Regra de negocio especifica | SPECS.md Secao 3 |
| Detalhe de um endpoint (request/response) | SPECS.md Secao 4 |
| Arquitetura e decisoes de design | ARCHITECTURE.md |
| O que implementar em uma task especifica | TASKS.md |
| Cenarios de teste para uma feature | TEST_PLAN.md |
| Requisitos do produto | PRD.md |
| Plano de seguranca e threat model | SECURITY.md |

---

## 11. Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.2 | 23/07/2026 | **Sprint 6 Frente 0 concluída:** BusinessUnit.java + ProductService.java entities. CnpjValidator com algoritmo unificado (IN RFB 2.119/2022). validateBusinessUnitTenant() IDOR cross-tenant. 261 testes (0 failures). 22 débitos catalogados. Referências atualizadas (TASKS v3.9, SPECS v2.8, TEST_PLAN v3.4, ARCH v2.12, SECURITY v1.3). | Agente IA |
| 1.1 | 23/07/2026 | **Sprint 5 Frentes 1-3a concluídas:** Endpoints 18→26 (Auth + Onboarding + Dashboard Cliente + App Switcher). Stack confirmado: Flyway 12.11.0, PG driver 42.7.11, OAuth2 Client. Referências atualizadas (TASKS v3.7, SPECS v2.7, SECURITY v1.2). 227 testes. | Agente IA |
| 1.0 | 20/07/2026 | Criação inicial: setup local, padrões de código, pipeline de segurança, 18 endpoints, convenções git, variáveis de ambiente | Agente IA |
