# PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION — Baseline de Especificações Técnicas

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Programa:** FBSO Platform — Portal Administrativo SaaS
- **Versão:** 1.1
- **Data de Criação:** 26 de Julho de 2026
- **Última Atualização:** 27 de Julho de 2026 (alinhamento com docs de negócio v1.2)
- **Status:** ✅ COMPLIANCE — Validado pelo Time de Arquitetura
- **Baseline de Negócio:** [Project Charter v1.2](../01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [BRD v1.2](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [Épicos v1.2](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [Features FEAT-EP-](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)
- **Documentos Complementares:** [ARCHITECTURE-DEFINITION](./PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md) · [SECURITY-DEFINITION](./PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md) · [STACK-MATRIX](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md)

---

## 1. Objetivo

Este documento define a **baseline de especificações técnicas** que TODAS as soluções do projeto FBSO Platform devem seguir. Ele garante consistência cross-solution em convenções, padrões de código, API, banco de dados, observabilidade e restrições técnicas. Cada time de solução parte desta baseline e a especializa em sua própria `SPEC.md`.

---

## 2. Convenções de Nomenclatura

### 2.1 Backend (Java / Spring Boot — S01)

| Elemento | Convenção | Exemplo |
|:---|:---|:---|
| **Pacote base** | `com.fbso.platform.admin` | — |
| **Camada Controller** | `{Entity}Controller` | `TenantController`, `PlanController` |
| **Camada Service** | `{Entity}Service` + `{Entity}ServiceImpl` | `TenantService`, `TenantServiceImpl` |
| **Camada Repository** | `{Entity}Repository` | `TenantRepository` |
| **Entidade** | `{Entity}` (substantivo singular) | `Tenant`, `Plan`, `User`, `BusinessUnit` |
| **DTO Request** | `{Entity}Request` ou `Create{Entity}Request` | `CreateTenantRequest`, `UpdatePlanRequest` |
| **DTO Response** | `{Entity}Response` | `TenantResponse`, `PlanListResponse` |
| **Enum** | `{Qualifier}{Type}` | `TenantStatus`, `PlanType`, `SubscriptionBillingCycle` |
| **Exception** | `{Context}Exception` | `TenantNotFoundException`, `DuplicateSubscriptionException` |
| **Método Controller** | RESTful: `getX()`, `createX()`, `updateX()`, `deleteX()` | `getTenants()`, `createPlan()`, `updateSubscription()` |
| **Testes** | `{Classe}Test` (unitário), `{Classe}IT` (integração) | `TenantServiceTest`, `TenantControllerIT` |

### 2.2 Frontend (Next.js / React — S02)

| Elemento | Convenção | Exemplo |
|:---|:---|:---|
| **Página** | `app/(area)/{rota}/page.tsx` | `app/(admin)/tenants/page.tsx` |
| **Componente** | PascalCase, kebab-case no arquivo | `TenantList.tsx`, `app-switcher.tsx` |
| **Hook** | `use{Nome}` | `useTenants`, `useAuth`, `usePermissions` |
| **API Client** | `lib/api/{recurso}.ts` | `lib/api/tenants.ts`, `lib/api/plans.ts` |
| **Tipo/Zod Schema** | `{Nome}Schema` (Zod) | `tenantSchema`, `createTenantSchema` |
| **Context** | `{Nome}Context` + `{Nome}Provider` | `AuthContext`, `TenantContext` |

### 2.3 Banco de Dados (PostgreSQL — S03)

| Elemento | Convenção | Exemplo |
|:---|:---|:---|
| **Tabela** | `snake_case` plural | `tenants`, `plans`, `subscriptions`, `business_units` |
| **Coluna PK** | `id UUID PRIMARY KEY DEFAULT gen_random_uuid()` | `id UUID` |
| **Coluna FK** | `{tabela_singular}_id` | `tenant_id`, `plan_id`, `business_unit_id` |
| **Auditoria** | `created_dt`, `updated_dt`, `created_by`, `updated_by` | — |
| **Soft Delete** | `deleted_dt TIMESTAMPTZ`, `deleted_by UUID` | — |
| **Índice** | `idx_{tabela}_{coluna}` | `idx_tenants_status`, `idx_users_email` |
| **Índice Único** | `uq_{tabela}_{coluna}` | `uq_tenants_cnpj` |
| **Índice Multi-Tenant** | `idx_{tabela}_tenant_{coluna}` | `idx_users_tenant_email` |
| **Policy RLS** | `{tabela}_isolation_policy` | `tenants_isolation_policy` |
| **Migration** | `V{NN}__{descricao_snake_case}.sql` | `V001__create_tenants_table.sql` |

### 2.4 API REST — Endpoints

| Elemento | Convenção | Exemplo |
|:---|:---|:---|
| **Base Path** | `/api/v1/{recurso}` | `/api/v1/tenants` |
| **Plural** | Substantivo plural | `/api/v1/tenants`, `/api/v1/plans` |
| **Recurso aninhado** | `/{pai_id}/{filho}` | `/api/v1/tenants/{id}/users` |
| **Ação não-CRUD** | Verbo no path (evitar) — usar sub-recurso | `/api/v1/tenants/{id}/activate` (POST) |
| **Query params** | `camelCase` | `?status=active&sortBy=createdDt&page=0&size=25` |

---

## 3. Padrões de API

### 3.1 URL Design

```
GET    /api/v1/tenants                 → Listar tenants (paginado)
POST   /api/v1/tenants                 → Criar tenant
GET    /api/v1/tenants/{id}            → Obter tenant por ID
PUT    /api/v1/tenants/{id}            → Atualizar tenant (completo)
PATCH  /api/v1/tenants/{id}/status     → Atualizar parcial (status)
DELETE /api/v1/tenants/{id}            → Soft-delete tenant
```

### 3.2 Paginação

```json
// Request
GET /api/v1/tenants?page=0&size=25&sort=createdDt,desc

// Response
{
  "content": [...],
  "page": 0,
  "size": 25,
  "totalElements": 142,
  "totalPages": 6,
  "first": true,
  "last": false
}
```

> **Regra:** Toda listagem é paginada. Máximo: `size=100`.

### 3.3 Formato de Erro

```json
{
  "timestamp": "2026-07-26T03:00:00Z",
  "status": 404,
  "error": "Not Found",
  "message": "Tenant não encontrado: id=abc-123",
  "path": "/api/v1/tenants/abc-123",
  "traceId": "a1b2c3d4e5f6"
}
```

> **Regra:** NUNCA expor stack trace em produção. `traceId` permite correlação com logs internos.

### 3.4 Versionamento

| Estratégia | Detalhe |
|:---|:---|
| **URL path** | `/api/v1/...` — versão no path |
| **Header (futuro)** | `Accept: application/vnd.fbso.v2+json` |
| **Regra** | Breaking change → nova versão. Additive change → mesma versão. |

### 3.5 Headers Padrão

| Header | Valor | Obrigatório |
|:---|:---|:---:|
| `Authorization` | `Bearer <JWT>` | ✅ (exceto health/public) |
| `Content-Type` | `application/json` | ✅ (POST/PUT/PATCH) |
| `Accept` | `application/json` | ✅ |
| `X-Request-Id` | UUID (gerado pelo Kong se ausente) | Recomendado |
| `X-Tenant-Id` | UUID (injetado pelo Kong) | ✅ (rotas tenant-scoped) |

### 3.6 Catálogo de Recursos API

Esta seção define o catálogo completo dos 11 recursos REST da FBSO Platform, migrado do API-CONTRACTS.md original (removido após consolidação). Os schemas detalhados de request/response estão no OpenAPI YAML canônico (`.specs/api/fbso-platform-api.yaml`).

| # | Recurso | Endpoint Base | Épico | Operações |
|:---|:---|:---|:---|:---|
| R-01 | **Tenants** | `/tenants` | EP-0002 | CRUD + ativar/suspender/reativar |
| R-02 | **Plans** | `/plans` | EP-0002 | CRUD + desativar |
| R-03 | **Subscriptions** | `/subscriptions` | EP-0002 | Criar, alterar, suspender, histórico |
| R-04 | **Users** | `/users` | EP-0003 | CRUD + convidar/desativar |
| R-05 | **Permissions** | `/permissions` | EP-0003 | Atribuir/revogar papéis e vínculos |
| R-06 | **Business Units** | `/business-units` | EP-0004 | CRUD + hierarquia + desativar |
| R-07 | **Products** | `/products` | EP-0004 | CRUD + ativar/desativar |
| R-08 | **Dashboard Admin** | `/dashboard/admin` | EP-0001 | Leitura de métricas operacionais |
| R-09 | **Dashboard Client** | `/dashboard/client` | EP-0004 | Leitura de métricas do cliente |
| R-10 | **Onboarding** | `/onboarding` | EP-0004 | Fluxo de primeiro acesso (4 passos) |
| R-11 | **Audit** | `/audit` | EP-0002 | Consulta ao histórico de auditoria |

#### R-01 — Tenants (`/tenants`)

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/tenants` | Listar todos os tenants (paginado, com filtros) | Admin FBSO |
| `GET` | `/tenants/{id}` | Obter detalhes de um tenant | Admin FBSO |
| `POST` | `/tenants` | Criar novo tenant (status: PENDING_ONBOARDING) | Admin FBSO |
| `PATCH` | `/tenants/{id}` | Atualizar dados cadastrais do tenant | Admin FBSO |
| `POST` | `/tenants/{id}/activate` | Ativar tenant manualmente | Admin FBSO |
| `POST` | `/tenants/{id}/suspend` | Suspender tenant (bloqueia acesso) | Admin FBSO |
| `POST` | `/tenants/{id}/reactivate` | Reativar tenant suspenso | Admin FBSO |

**Query Params:** `status` (ACTIVE, SUSPENDED, PENDING_ONBOARDING, INACTIVE), `planId`, `search` (razão social ou nome fantasia), `page`, `size`, `sort`

#### R-02 — Plans (`/plans`)

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/plans` | Listar todos os planos | Admin FBSO / Gestor Produto |
| `GET` | `/plans/{id}` | Obter detalhes de um plano | Admin FBSO / Gestor Produto |
| `POST` | `/plans` | Criar novo plano comercial | Gestor Produto |
| `PATCH` | `/plans/{id}` | Atualizar plano (gera nova versão) | Gestor Produto |
| `POST` | `/plans/{id}/deactivate` | Desativar plano | Gestor Produto |

#### R-03 — Subscriptions (`/subscriptions`)

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/tenants/{tenantId}/subscriptions` | Histórico de assinaturas do tenant | Admin FBSO / Admin Tenant |
| `GET` | `/tenants/{tenantId}/subscriptions/active` | Assinatura ativa do tenant | Admin FBSO / Admin Tenant |
| `POST` | `/tenants/{tenantId}/subscriptions` | Criar nova assinatura (vincula plano) | Admin FBSO |
| `POST` | `/subscriptions/{id}/change-plan` | Upgrade/downgrade de plano | Admin FBSO |
| `POST` | `/subscriptions/{id}/suspend` | Suspender assinatura | Admin FBSO |
| `POST` | `/subscriptions/{id}/reactivate` | Reativar assinatura | Admin FBSO |

#### R-04 — Users (`/users`)

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/users` | Listar usuários do tenant (contexto do JWT) | Admin Tenant |
| `GET` | `/users/{id}` | Detalhes do usuário | Admin Tenant / próprio usuário |
| `POST` | `/users` | Convidar novo usuário (dispara e-mail) | Admin Tenant |
| `PATCH` | `/users/{id}` | Atualizar dados do usuário | Admin Tenant / próprio usuário |
| `POST` | `/users/{id}/deactivate` | Desativar usuário (bloqueia acesso) | Admin Tenant |
| `POST` | `/users/{id}/reactivate` | Reativar usuário | Admin Tenant |

#### R-05 — Permissions (`/permissions`)

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/users/{userId}/permissions` | Permissões atuais do usuário | Admin Tenant |
| `PUT` | `/users/{userId}/permissions` | Substituir vínculos (BU + módulos) | Admin Tenant |
| `PATCH` | `/users/{userId}/permissions/role` | Alterar papel do usuário | Admin Tenant |
| `DELETE` | `/users/{userId}/permissions/{buId}` | Remover acesso a uma Unidade de Negócio | Admin Tenant |

#### R-06 — Business Units (`/business-units`)

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/business-units` | Listar unidades do tenant (hierárquico) | Admin Tenant / Manager BU |
| `GET` | `/business-units/{id}` | Detalhes da unidade | Admin Tenant / Manager BU / Operador BU |
| `POST` | `/business-units` | Cadastrar nova unidade | Admin Tenant / Manager BU |
| `PATCH` | `/business-units/{id}` | Atualizar dados da unidade | Admin Tenant / Manager BU |
| `POST` | `/business-units/{id}/deactivate` | Desativar unidade | Admin Tenant |

#### R-07 — Products (`/products`)

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/products` | Listar produtos (filtrado pela BU ativa) | Todos com acesso à BU |
| `GET` | `/products/{id}` | Detalhes do produto | Todos com acesso à BU |
| `POST` | `/products` | Cadastrar novo produto | Admin Tenant / Manager BU |
| `PATCH` | `/products/{id}` | Atualizar produto | Admin Tenant / Manager BU |
| `POST` | `/products/{id}/deactivate` | Desativar produto | Admin Tenant / Manager BU |
| `POST` | `/products/{id}/activate` | Reativar produto | Admin Tenant / Manager BU |

**Query Params:** `type` (PRODUCT, SERVICE), `status` (ACTIVE, INACTIVE), `search` (nome ou SKU)

#### R-08/R-09 — Dashboards (`/dashboard/admin`, `/dashboard/client`)

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/dashboard/admin/summary` | Indicadores principais do SaaS | Admin FBSO |
| `GET` | `/dashboard/admin/accounts-by-status` | Contas agrupadas por status | Admin FBSO |
| `GET` | `/dashboard/admin/accounts-by-plan` | Contas agrupadas por plano | Admin FBSO |
| `GET` | `/dashboard/admin/evolution` | Evolução da base ao longo do tempo | Admin FBSO |
| `GET` | `/dashboard/admin/alerts` | Alertas de atenção | Admin FBSO |
| `GET` | `/dashboard/client/summary` | Resumo da conta do cliente | Todos os papéis do tenant |

**Query Params:** `period` (7d, 30d, 90d, currentMonth, currentYear) ou customizado (`from` + `to`)

#### R-10 — Onboarding (`/onboarding`)

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/onboarding/status` | Status atual do onboarding | Cliente autenticado |
| `POST` | `/onboarding/step-1` | Confirmar/atualizar dados cadastrais | Cliente autenticado |
| `POST` | `/onboarding/step-2` | Cadastrar primeira Unidade de Negócio (Matriz) | Cliente autenticado |
| `GET` | `/onboarding/step-3` | Visualizar resumo do plano contratado | Cliente autenticado |
| `POST` | `/onboarding/complete` | Finalizar onboarding (tenant → ACTIVE) | Cliente autenticado |

#### R-11 — Audit (`/audit`)

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/audit` | Listar registros de auditoria | Admin FBSO / Auditor |
| `GET` | `/audit/{id}` | Detalhes de um registro de auditoria | Admin FBSO / Auditor |
| `GET` | `/tenants/{tenantId}/audit` | Auditoria de um tenant específico | Admin FBSO |

**Query Params:** `entityType` (TENANT, PLAN, SUBSCRIPTION, USER, PERMISSION, BUSINESS_UNIT, PRODUCT), `action` (CREATED, UPDATED, SUSPENDED, REACTIVATED, DEACTIVATED), `from`, `to`, `page`, `size`, `sort`

### 3.7 Matriz RBAC × Endpoints

| Recurso | Admin FBSO | Admin Tenant | Manager BU | Operator BU | Auditor |
|:---|:---:|:---:|:---:|:---:|:---:|
| **Tenants** (`/tenants`) | CRUD + ações | — | — | — | — |
| **Plans** (`/plans`) | CRUD | — | — | — | — |
| **Subscriptions** (`/subscriptions`) | CRUD | Ver (seu) | — | — | Ver |
| **Users** (`/users`) | — | CRUD | — | — | — |
| **Permissions** (`/permissions`) | — | CRUD | — | — | — |
| **Business Units** (`/business-units`) | — | CRUD | Ver + Editar (sua) | Ver (sua) | Ver |
| **Products** (`/products`) | — | CRUD | Criar, Editar, Ver | Ver | Ver |
| **Dashboard Admin** | Ver | — | — | — | — |
| **Dashboard Client** | — | Ver | Ver | Ver | Ver |
| **Onboarding** | — | Executar | Executar | Executar | — |
| **Audit** | Ver (todos) | — | — | — | Ver (tenant) |

> **Origem:** As seções 3.6 e 3.7 foram migradas do API-CONTRACTS.md original §3-5 (removido). Os schemas JSON detalhados foram removidos — o OpenAPI YAML em `.specs/api/fbso-platform-api.yaml` é a fonte canônica dos schemas de request/response. As convenções de nomenclatura desta seção prevalecem sobre o API-CONTRACTS.md original (ex: `camelCase` para query params, formato de erro flat).

---

## 4. Padrões de Banco de Dados

### 4.1 Modelo de Dados — Template de Tabela

```sql
CREATE TABLE fbso_portal.{tabela} (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id   UUID NOT NULL,  -- Multi-tenant identifier
    -- colunas de negócio aqui
    created_dt  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_dt  TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_by  UUID,
    updated_by  UUID,
    deleted_dt  TIMESTAMPTZ,    -- Soft delete
    deleted_by  UUID
);

-- RLS
ALTER TABLE fbso_portal.{tabela} ENABLE ROW LEVEL SECURITY;
ALTER TABLE fbso_portal.{tabela} FORCE ROW LEVEL SECURITY;

CREATE POLICY {tabela}_tenant_isolation ON fbso_portal.{tabela}
    USING (tenant_id = NULLIF(current_setting('app.current_tenant_id', true), '')::UUID)
    WITH CHECK (tenant_id = NULLIF(current_setting('app.current_tenant_id', true), '')::UUID);

-- Índices
CREATE INDEX idx_{tabela}_tenant_id ON fbso_portal.{tabela}(tenant_id);
CREATE INDEX idx_{tabela}_tenant_created ON fbso_portal.{tabela}(tenant_id, created_dt DESC);
```

### 4.2 Soft Delete + RLS

```sql
CREATE POLICY {tabela}_soft_delete ON fbso_portal.{tabela}
    USING (
        tenant_id = current_setting('app.current_tenant_id')::UUID
        AND (deleted_dt IS NULL OR pg_has_role(current_user, 'admin_role', 'member'))
    );
```

### 4.3 Índices Únicos com Soft Delete

```sql
-- Garantir unicidade apenas entre registros ativos
CREATE UNIQUE INDEX uq_{tabela}_{coluna}_active
    ON fbso_portal.{tabela}(tenant_id, {coluna})
    WHERE deleted_dt IS NULL;
```

### 4.4 Regras Obrigatórias

| Regra | Detalhe |
|:---|:---|
| **Schema sempre qualificado** | `SELECT * FROM fbso_portal.tenants` — nunca sem schema |
| **Nunca DELETE físico** | `UPDATE ... SET deleted_dt = now() WHERE ...` — nunca `DELETE FROM` |
| **SET tenant_id primeiro** | `SET app.current_tenant_id = '<uuid>'` antes de qualquer query |
| **Transaction boundary** | Toda operação de escrita em `@Transactional` |
| **Pool connection** | HikariCP — `spring.datasource.hikari.maximum-pool-size=20` |

---

## 5. Padrões de Logging e Observabilidade

### 5.1 Formato de Log

```json
{
  "timestamp": "2026-07-26T03:00:00.123Z",
  "level": "INFO",
  "logger": "com.fbso.platform.admin.service.TenantService",
  "message": "Tenant activated: id=a1b2c3",
  "traceId": "a1b2c3d4e5f6",
  "spanId": "1a2b3c4d5e6f",
  "tenantId": "a1b2c3d4-...",
  "userId": "u1u2u3u4-..."
}
```

### 5.2 O Que NUNCA Logar

- Senhas, tokens JWT completos, secrets
- CPF, RG, dados de cartão
- Corpo completo de requisições com dados sensíveis
- Stack traces em produção (usar `traceId` para correlação)

### 5.3 Níveis de Log

| Nível | Quando Usar |
|:---|:---|
| **ERROR** | Exceção não tratada, falha de integração, dados inconsistentes |
| **WARN** | Degradação, retry, recurso próximo do limite |
| **INFO** | Ações de negócio: tenant criado, plano alterado, usuário convidado |
| **DEBUG** | Detalhes de fluxo (dev/staging apenas) |
| **TRACE** | Query parameters, payloads (nunca em produção) |

### 5.4 Spans Manuais Obrigatórios

| Operação | Span Name |
|:---|:---|
| Validação de tenant | `tenant.validate` |
| Query com RLS | `db.query.rls.{tabela}` |
| Chamada externa (Keycloak) | `keycloak.{operacao}` |
| Operação de negócio crítica | `business.{acao}` (ex: `business.activate_tenant`) |

---

## 6. Padrões de Código

### 6.1 Java (S01)

| Regra | Detalhe |
|:---|:---|
| **Imutabilidade** | DTOs com `record` (Java 16+). Entidades com `@Value` ou getters sem setters. |
| **Injeção** | Injeção por construtor (não `@Autowired` em campos). |
| **Exceções** | Custom exceptions com `@ResponseStatus`. `@ControllerAdvice` global. |
| **Validação** | `@Valid` em todos os controllers. `@Validated` nos services. |
| **Logging** | `@Slf4j` (Lombok). Nunca `System.out.println`. |
| **Transactions** | `@Transactional(readOnly = true)` padrão. `readOnly = false` explícito em escritas. |
| **Build** | Maven wrapper (`./mvnw`). Proibido usar Maven do sistema. |

#### Estrutura de Pacotes (package-by-layer)

```
com.fbso.platform.admin/              ← Projeto: ms-fbso-platform-admin
│
├── config/
│   ├── SecurityConfig.java           ← Configuração Spring Security + JWT
│   ├── TenantContext.java            ← Holder do tenant_id da requisição atual
│   └── WebConfig.java                ← CORS, interceptors
│
├── security/
│   ├── JwtAuthenticationFilter.java  ← Filtro que valida JWT em toda requisição
│   ├── TenantIsolationFilter.java    ← Injeta tenant_id nas queries
│   └── RbacInterceptor.java          ← Verifica permissões por recurso/ação
│
├── tenant/
│   ├── TenantController.java         ← REST /tenants
│   ├── TenantService.java
│   └── TenantRepository.java
│
├── plan/
│   ├── PlanController.java           ← REST /plans
│   ├── PlanService.java
│   └── PlanRepository.java
│
├── subscription/
│   ├── SubscriptionController.java   ← REST /subscriptions
│   ├── SubscriptionService.java
│   └── SubscriptionRepository.java
│
├── user/
│   ├── UserController.java           ← REST /users
│   ├── UserService.java
│   ├── UserRepository.java
│   └── UserInviteService.java        ← Lógica de convite por e-mail
│
├── permission/
│   ├── PermissionController.java     ← REST /permissions
│   ├── PermissionService.java
│   └── PermissionRepository.java
│
├── businessunit/
│   ├── BusinessUnitController.java   ← REST /business-units
│   ├── BusinessUnitService.java
│   └── BusinessUnitRepository.java
│
├── product/
│   ├── ProductController.java        ← REST /products
│   ├── ProductService.java
│   └── ProductRepository.java
│
├── dashboard/
│   ├── DashboardController.java      ← REST /dashboard/admin, /dashboard/client
│   └── DashboardService.java
│
├── onboarding/
│   ├── OnboardingController.java     ← REST /onboarding
│   └── OnboardingService.java
│
├── audit/
│   ├── AuditController.java          ← REST /audit
│   ├── AuditService.java
│   └── AuditEntityListener.java      ← JPA Entity Listener para auto-auditoria
│
└── common/
    ├── BaseEntity.java               ← Superclasse com campos de auditoria
    └── SoftDeleteRepository.java     ← Repository base com filtro deleted_dt IS NULL
```

### 6.2 TypeScript/React (S02)

| Regra | Detalhe |
|:---|:---|
| **Tipagem** | `strict: true` no `tsconfig.json`. Sem `any` exceto com justificativa. |
| **Componentes** | Server Components por padrão. `'use client'` só quando necessário. |
| **Estado** | Zustand para estado global. `useState`/`useReducer` para local. |
| **Fetch** | SWR com tipos gerados do OpenAPI. Sem `fetch` cru nos componentes. |
| **Validação** | Zod schemas para formulários + tipos TypeScript via `z.infer`. |
| **Acessibilidade** | `jsx-a11y` rules no ESLint. `aria-label` em elementos sem texto visível. |

#### Estrutura de Rotas (Next.js App Router)

```
web_app-fbso-platform-portal/
│
├── app/
│   ├── (auth)/
│   │   ├── login/page.tsx            ← Tela de login (redireciona para Keycloak)
│   │   └── reset-password/page.tsx   ← Recuperação de senha
│   │
│   ├── (onboarding)/
│   │   └── onboarding/
│   │       ├── page.tsx              ← Wizard — Passo 1: Confirmar dados
│   │       ├── step-2/page.tsx       ← Passo 2: Cadastrar Matriz
│   │       ├── step-3/page.tsx       ← Passo 3: Resumo do Plano
│   │       └── step-4/page.tsx       ← Passo 4: Boas-vindas
│   │
│   ├── (admin)/                      ← Rotas do time FBSO.ORG
│   │   ├── dashboard/page.tsx        ← Dashboard administrativo
│   │   ├── tenants/page.tsx          ← Gestão de contas
│   │   ├── plans/page.tsx            ← Configuração de planos
│   │   └── audit/page.tsx            ← Histórico de auditoria
│   │
│   └── (portal)/                     ← Rotas do cliente
│       ├── dashboard/page.tsx        ← Dashboard do cliente
│       ├── business-units/page.tsx   ← Unidades de Negócio
│       ├── products/page.tsx         ← Catálogo de Produtos
│       ├── users/page.tsx            ← Gestão de Usuários (Admin Tenant)
│       └── profile/page.tsx          ← Perfil do usuário
│
├── components/
│   ├── layout/
│   │   ├── AppSwitcher.tsx           ← Seletor de módulos no topo
│   │   ├── Sidebar.tsx               ← Menu lateral dinâmico
│   │   └── BusinessUnitSelector.tsx  ← Seletor de Unidade de Negócio
│   ├── dashboard/
│   │   ├── MetricsCard.tsx           ← Card de indicador
│   │   └── TrendChart.tsx            ← Gráfico de evolução
│   └── common/
│       ├── DataTable.tsx             ← Tabela paginada reutilizável
│       └── StatusBadge.tsx           ← Badge de status colorido
│
├── lib/
│   ├── auth.ts                       ← Integração com Keycloak (next-auth)
│   ├── api-client.ts                 ← Cliente HTTP com injeção de JWT
│   └── permissions.ts                ← Hook usePermission(resource, action)
│
└── mocks/
    └── handlers/                     ← MSW handlers baseados no OpenAPI
```

> **Origem:** As estruturas de pacotes (§6.1) e rotas (§6.2) foram migradas do ARCHITECTURE.md original §5-6 (removido após consolidação).

### 6.3 SQL / Flyway (S06)

| Regra | Detalhe |
|:---|:---|
| **Migrations imutáveis** | Nunca alterar migration já aplicada. Criar nova. |
| **Rollback** | Criar migration reversa com prefixo `U{NN}`. |
| **Idempotência** | Usar `CREATE TABLE IF NOT EXISTS`, `DROP TABLE IF EXISTS`. |
| **Schema** | Sempre qualificar: `fbso_portal.{tabela}`. |

---

## 7. Restrições Técnicas Cross-Solution

| Restrição | Valor | Justificativa |
|:---|:---|:---|
| **Timeout HTTP (backend)** | 30s | Evitar conexões penduradas |
| **Timeout JDBC** | 30s | HikariCP `connection-timeout` |
| **Pool máximo PostgreSQL** | 20 conexões | DigitalOcean Managed DB tier básico |
| **Tamanho máximo payload** | 10 MB | Kong `client_max_body_size` |
| **Rate limit global** | 100 req/s por IP | Kong |
| **Rate limit por tenant** | 60 req/s | Kong |
| **Tamanho máximo página** | 100 registros | API paginação |
| **Tempo máximo JWT access token** | 5 minutos | Segurança: tokens de curta duração |
| **Tempo máximo refresh token** | 30 dias | UX: login transparente por 30 dias |
| **Cache JWKS** | 15 minutos | Kong |
| **Retenção logs (dev)** | 7 dias | Disco local |
| **Retenção logs (prod)** | 90 dias | CloudWatch / equivalente DO |

---

## 8. Referências a Blueprints

| Blueprint | Caminho | Uso |
|:---|:---|:---|
| Dockerfile Native | `ms-fbso-platform-admin/Dockerfile` | Template GraalVM Native Image |
| Dockerfile JVM | `ms-fbso-platform-admin/Dockerfile.jvm` | Fallback JVM |
| docker-compose | `ms-fbso-platform-admin/docker-compose.yml` | Ambiente dev |
| Keycloak Realm Config | `ms-fbso-platform-admin/keycloak/realm-config.json` | Template de Realm |
| Flyway Migration | `ms-fbso-platform-admin/src/main/resources/db/migration/` | Template SQL |
| application.yml | `ms-fbso-platform-admin/src/main/resources/application.yml` | Configuração Spring Boot |
| pom.xml | `ms-fbso-platform-admin/pom.xml` | Dependências e plugins |

---

## 9. Referências

| Documento | Relação |
|:---|:---|
| [ARCHITECTURE-DEFINITION](./PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md) | Diagramas C4, integrações |
| [SECURITY-DEFINITION](./PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md) | Políticas de segurança |
| [STACK-MATRIX](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md) | Versões de tecnologias |
| `/architecture/blueprints/` | Templates de código e configuração |
| `/architecture/data_standards/` | Padrões de modelagem de dados |

---

## Histórico de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 26/07/2026 | Criação inicial: convenções de nomenclatura (Java, React, PostgreSQL, API), padrões de API (URL, paginação, erros, versionamento), padrões de banco (RLS, Soft Delete, índices), padrões de código (Java, TypeScript, SQL), logging/observabilidade, restrições cross-solution, blueprints. | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Baseline de especificações que todas as soluções devem seguir. Resultado da Fase 7 do Roadmap de Definições Técnicas.*
