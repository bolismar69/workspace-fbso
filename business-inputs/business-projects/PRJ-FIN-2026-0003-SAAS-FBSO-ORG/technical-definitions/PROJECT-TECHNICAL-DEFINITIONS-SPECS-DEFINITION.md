# PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION — Baseline de Especificações Técnicas

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Programa:** FBSO Platform — Portal Administrativo SaaS
- **Versão:** 1.0
- **Data de Criação:** 26 de Julho de 2026
- **Última Atualização:** 26 de Julho de 2026
- **Status:** ✅ COMPLIANCE — Validado pelo Time de Arquitetura
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

### 6.2 TypeScript/React (S02)

| Regra | Detalhe |
|:---|:---|
| **Tipagem** | `strict: true` no `tsconfig.json`. Sem `any` exceto com justificativa. |
| **Componentes** | Server Components por padrão. `'use client'` só quando necessário. |
| **Estado** | Zustand para estado global. `useState`/`useReducer` para local. |
| **Fetch** | SWR com tipos gerados do OpenAPI. Sem `fetch` cru nos componentes. |
| **Validação** | Zod schemas para formulários + tipos TypeScript via `z.infer`. |
| **Acessibilidade** | `jsx-a11y` rules no ESLint. `aria-label` em elementos sem texto visível. |

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
