# INTEGRATION-MAP.md — Mapa de Integrações do Sistema

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Programa:** FBSO Platform — Portal Administrativo SaaS (Fase 0 — Core)
- **Versão:** 1.0
- **Data:** 13 de Julho de 2026
- **Status:** Visão de Alto Nível — a ser detalhada pelos times técnicos
- **Origem:** [TECHNICAL-PLAN.md](./TECHNICAL-PLAN.md) e [ARCHITECTURE.md](./ARCHITECTURE.md)

---

## 1. Visão Geral das Integrações

```
┌────────────────────────────────────────────────────────────────────────┐
│                         FBSO PLATFORM                                  │
│                                                                        │
│                                                                        │
│  ┌──────────┐      HTTPS        ┌──────────┐      JDBC        ┌──────┐│
│  │ Frontend  │◄────────────────▶│ Backend   │───────────────▶│  DB  ││
│  │ React/    │   REST/JSON      │ Java 25/  │                │Postgre││
│  │ Next.js   │   + JWT Bearer   │ Spring    │                │ SQL  ││
│  └─────┬─────┘                  └─────┬─────┘                └──────┘│
│        │                              │                               │
│        │                              │                               │
│        │     OIDC/SAML                │  Validate JWT                 │
│        │     (Redirect)               │  (local)                      │
│        │                              │                               │
│        ▼                              ▼                               │
│  ┌──────────┐                                                         │
│  │ Keycloak  │                                                        │
│  │ (Docker)  │                                                        │
│  └──────────┘                                                         │
│                                                                        │
│  ┌ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┐  │
│  │                    FUTURO (Fase Tributali-Engine)                  │
│  │                                                                    │
│  │  ┌──────────┐      AMQP        ┌──────────┐      JDBC     ┌──────┐│
│  │  │ Backend   │◄──────────────▶│ RabbitMQ  │─────────────▶│  DB  ││
│  │  │ (Producer)│                 │ (Broker)  │              │Postgre││
│  │  └──────────┘                  └──────────┘              │ SQL  ││
│  │                                                          └──────┘│
│  └ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┘  │
└────────────────────────────────────────────────────────────────────────┘
```

---

## 2. Matriz de Integrações

| # | Origem | Destino | Protocolo | Direção | Autenticação | Fase |
|:---|:---|:---|:---|:---|:---|:---|
| **I-01** | Browser (Cliente) | Frontend (Next.js) | HTTPS | Bidirecional | — (página pública) + JWT (rotas autenticadas) | **Fase 0** |
| **I-02** | Browser → Frontend | Keycloak | HTTPS (Redirect OIDC/SAML) | Unidirecional (redirecionamento) | — | **Fase 0** |
| **I-03** | Keycloak → Browser | Frontend | HTTPS (Redirect com JWT) | Unidirecional (token) | Assinatura JWT (RS256) | **Fase 0** |
| **I-04** | Frontend (Next.js) | Backend (Spring Boot) | HTTPS + REST/JSON | Bidirecional | JWT Bearer Token no header `Authorization` | **Fase 0** |
| **I-05** | Backend (Spring Boot) | Keycloak | HTTPS (validação de token) | Unidirecional | Client credentials (backend → Keycloak admin API) | **Fase 0** |
| **I-06** | Backend (Spring Boot) | PostgreSQL | JDBC sobre TCP/TLS | Bidirecional | User/password + TLS | **Fase 0** |
| **I-07** | Backend (Producer) | RabbitMQ | AMQP sobre TCP/TLS | Unidirecional (publicação) | User/password + TLS | Futuro |
| **I-08** | RabbitMQ | Backend (Consumer) | AMQP sobre TCP/TLS | Unidirecional (consumo) | User/password + TLS | Futuro |

---

## 3. Detalhamento das Integrações (Fase 0)

### I-01/I-04 — Frontend ↔ Backend

```
[Frontend Next.js]                        [Backend Spring Boot]
        │                                          │
        │  GET /api/dashboard/admin                │
        │  Authorization: Bearer <JWT>             │
        │─────────────────────────────────────────▶│
        │                                          │ 1. Validar JWT (assinatura, exp)
        │                                          │ 2. Extrair tenant_id, roles
        │                                          │ 3. Aplicar TenantIsolationFilter
        │                                          │ 4. Executar query com tenant_id
        │                                          │ 5. Aplicar RBAC (role × resource)
        │  200 OK                                   │
        │  { "activeTenants": 42, ... }            │
        │◀─────────────────────────────────────────│
```

| Atributo | Valor |
|:---|:---|
| **Protocolo** | HTTPS + REST/JSON |
| **Autenticação** | JWT Bearer Token (`Authorization: Bearer eyJ...`) |
| **Formato** | JSON (request/response) |
| **Versionamento** | Prefixo `/api/v1/` (a confirmar) |
| **Timeout** | 30s (padrão), 60s (dashboard com agregações) |
| **Rate Limiting** | A definir (por tenant, por IP) |
| **CORS** | Origem: domínio do frontend. Métodos: GET, POST, PUT, PATCH, DELETE. Headers: Authorization, Content-Type |

### I-02/I-03 — Autenticação (Frontend ↔ Keycloak)

```
[Browser]                [Keycloak]                   [Backend]
    │                        │                            │
    │  GET /login            │                            │
    │───────────────────────▶│                            │
    │                        │                            │
    │  302 Redirect          │                            │
    │  → Keycloak login      │                            │
    │◀───────────────────────│                            │
    │                        │                            │
    │  POST credentials      │                            │
    │───────────────────────▶│                            │
    │                        │                            │
    │  302 Redirect          │                            │
    │  → /callback?code=...  │                            │
    │◀───────────────────────│                            │
    │                        │                            │
    │  GET /api/callback?code│                            │
    │────────────────────────────────────────────────────▶│
    │                        │                            │
    │                        │  POST /token (code)        │
    │                        │◀───────────────────────────│
    │                        │                            │
    │                        │  JWT Token                  │
    │                        │────────────────────────────▶│
    │                        │                            │
    │  Set-Cookie: JWT       │                            │
    │◀────────────────────────────────────────────────────│
```

| Atributo | Valor |
|:---|:---|
| **Protocolo** | OIDC (OpenID Connect) para SPA + SAML 2.0 (para clientes corporativos com IdP próprio) |
| **Fluxo** | Authorization Code Flow com PKCE (OIDC) |
| **Token** | JWT assinado com RS256. Expiração: 60 minutos (access token) + 24h (refresh token) |
| **Realm** | `fbso-platform` |
| **Clients** | `fbso-platform-frontend` (public), `fbso-platform-backend` (confidential) |
| **Roles** | `ADMIN_TENANT`, `MANAGER_BU`, `OPERATOR_BU`, `AUDITOR`. O perfil "Admin FBSO" (time interno) autentica via realm `master` do Keycloak com role `ROLE_ADMIN` e acessa endpoints administrativos cross-tenant |
| **Mapeamento de Claims** | `tenant_id` ← atributo do grupo do Keycloak, `business_unit_ids` ← grupos do usuário, `roles` ← roles do realm |
| **Logout** | RP-Initiated Logout (OIDC) — frontend redireciona para Keycloak `/logout` |

### I-05 — Validação de Token (Backend → Keycloak)

| Atributo | Valor |
|:---|:---|
| **Propósito** | Validar assinatura e expiração do JWT. Opcionalmente, consultar estado do usuário (ex: revogado) |
| **Frequência** | Toda requisição (validação local da assinatura JWT). Consulta ao Keycloak apenas em caso de dúvida (ex: refresh token) |
| **Endpoint Keycloak** | `GET /realms/fbso-platform/protocol/openid-connect/userinfo` |
| **Autenticação** | Client credentials (`fbso-platform-backend` + client secret) |

### I-06 — Persistência (Backend ↔ PostgreSQL)

```
[Backend Spring Boot]                           [PostgreSQL]
        │                                             │
        │  JDBC Connection (HikariCP Pool)            │
        │─────────────────────────────────────────────│
        │                                             │
        │  SELECT * FROM tenant                       │
        │  WHERE deleted_dt IS NULL                   │
        │    AND tenant_id = ?  ← TenantIsolation     │
        │─────────────────────────────────────────────▶│
        │                                             │
        │  ResultSet (apenas dados do tenant)         │
        │◀─────────────────────────────────────────────│
```

| Atributo | Valor |
|:---|:---|
| **Driver** | JDBC PostgreSQL (org.postgresql) |
| **Connection Pool** | HikariCP (padrão Spring Boot). Pool size: a definir (sugestão inicial: 10-20 conexões) |
| **Schema** | `fbso_platform` (schema dedicado dentro do banco) |
| **Migrations** | Flyway ou Liquibase (a definir). Versionamento sequencial (V001, V002...) |
| **Transações** | `@Transactional` do Spring. Isolation level padrão: READ_COMMITTED. Operações financeiras (assinaturas, planos) devem usar `@Transactional(isolation = REPEATABLE_READ)` ou `SELECT ... FOR UPDATE` para evitar phantom reads |
| **TLS** | TLS 1.3 para conexões em staging e produção. Dev local sem TLS |
| **Soft Delete** | Filtro automático `WHERE deleted_dt IS NULL` em todas as queries |

---

## 4. Fluxo de Dados — Cenários Principais

### 4.1 Criação de Tenant e Onboarding

```
[Admin FBSO]  →  [Frontend]  →  [Backend]  →  [DB]
     │               │              │            │
     │  POST/tenants  │              │            │
     │──────────────▶│              │            │
     │               │  POST/tenants│            │
     │               │─────────────▶│            │
     │               │              │ INSERT     │
     │               │              │───────────▶│
     │               │              │            │
     │               │              │ Tenant     │
     │               │              │ criado +   │
     │               │              │ email      │
     │               │◀─────────────│ enviado    │
     │  201 + link   │              │            │
     │◀──────────────│              │            │
     │               │              │            │
     │  [Cliente clica no link do e-mail]        │
     │               │              │            │
     │  GET /onboarding             │            │
     │──────────────▶│              │            │
     │  (redireciona │              │            │
     │   p/ Keycloak)│              │            │
     │               │              │            │
     │  [Autenticação OIDC - ver §3 I-02/I-03]   │
     │               │              │            │
     │  POST /onboarding/step-1    │            │
     │──────────────▶│─────────────▶│            │
     │               │              │ UPDATE     │
     │               │              │ tenant     │
     │               │              │───────────▶│
     │               │              │            │
     │  ...Passos 2-4 do onboarding...           │
     │               │              │            │
     │  Tenant status: ACTIVE       │            │
```

### 4.2 Verificação de Permissão (RBAC)

```
[Browser]  →  [Frontend]  →  [Backend Security]
    │              │                    │
    │  Clique em   │                    │
    │  "Editar     │                    │
    │   Produto"   │                    │
    │─────────────▶│                    │
    │              │  usePermission(    │
    │              │   'product',       │
    │              │   'edit')          │
    │              │                    │
    │              │  Verifica JWT:     │
    │              │  role = OPERATOR?  │
    │              │  → false           │
    │              │  Botão NÃO render  │
    │              │                    │
    │  (botão não  │                    │
    │   visível)   │                    │
    │◀─────────────│                    │
```

### 4.3 Isolamento Multi-Tenant em Query

```
[Backend]                               [PostgreSQL]
    │                                         │
    │  User: João (tenant_id = t-42)          │
    │  JWT: {tenant_id: "t-42",              │
    │        roles: [{bu: "bu-01",            │
    │                role: "OPERATOR_BU"}]}   │
    │                                         │
    │  GET /products?bu=bu-01                 │
    │                                         │
    │  1. TenantContext.set("t-42")           │
    │  2. RbacCheck: OPERATOR_BU pode         │
    │     ver products? → SIM                 │
    │  3. BuCheck: bu-01 está                 │
    │     nos roles[]? → SIM                  │
    │                                         │
    │  SELECT * FROM product_service          │
    │  WHERE business_unit_id = 'bu-01'       │
    │    AND tenant_id = 't-42'  ← AUTO      │
    │    AND deleted_dt IS NULL   ← AUTO     │
    │─────────────────────────────────────────▶
    │                                         │
    │  ResultSet (apenas produtos da BU-01)   │
    │◀─────────────────────────────────────────│
```

---

## 5. Dependências de Infraestrutura

| Componente | Tipo | Dev Local | Staging | Produção |
|:---|:---|:---|:---|:---|
| **PostgreSQL** | Banco de Dados | Docker (postgres:17, tag exata em produção) | K8s StatefulSet | K8s StatefulSet (HA) |
| **Keycloak** | IdP | Docker (docker-compose) | K8s Deployment (1 réplica) | K8s Deployment (2+ réplicas) |
| **Backend** | Aplicação | Docker (docker-compose) ou IDE | K8s Deployment | K8s Deployment (auto-scale) |
| **Frontend** | Aplicação | Node (`next dev`) ou Docker | K8s Deployment | K8s Deployment (auto-scale) |
| **RabbitMQ** | Mensageria | — (não provisionado) | — | K8s StatefulSet (futuro) |

### Docker Compose (Dev Local)

```yaml
# docker-compose.yml — Ambiente de Desenvolvimento Local
services:
  postgres:
    image: postgres:17
    environment:
      POSTGRES_DB: fbso_platform
      POSTGRES_USER: fbso_admin
      POSTGRES_PASSWORD: dev_password
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

  keycloak:
    image: quay.io/keycloak/keycloak:26
    environment:
      KC_DB: postgres
      KC_DB_URL: jdbc:postgresql://postgres:5432/keycloak
      KC_DB_USERNAME: keycloak
      KC_DB_PASSWORD: keycloak_password
      KEYCLOAK_ADMIN: ${KEYCLOAK_ADMIN_USER:-admin}
      KEYCLOAK_ADMIN_PASSWORD: ${KEYCLOAK_ADMIN_PASSWORD:-changeme}  # Sobrescrever via variável de ambiente em staging/produção
    ports:
      - "8443:8443"   # Keycloak na 8443 para evitar conflito com Spring Boot (8080). Ajustar KC_HOSTNAME_PORT conforme necessário
    depends_on:
      - postgres
    volumes:
      - ./keycloak/realm-config.json:/opt/keycloak/data/import/realm.json

  backend:
    build: ../../backend/java/spring/microservices/ms-fbso-platform-admin
    environment:
      SPRING_DATASOURCE_URL: jdbc:postgresql://postgres:5432/fbso_platform
      KEYCLOAK_AUTH_SERVER_URL: http://keycloak:8080
    ports:
      - "8081:8081"
    depends_on:
      - postgres
      - keycloak

  frontend:
    build: ../../frontend/javascript/react/web_apps/web_app-fbso-platform-portal
    environment:
      NEXT_PUBLIC_API_URL: http://localhost:8081
      NEXT_PUBLIC_KEYCLOAK_URL: http://localhost:8080
    ports:
      - "3000:3000"
    depends_on:
      - backend

volumes:
  pgdata:
```

---

## 6. Matriz de Comunicação por Épico

| Épico | Frontend→Backend | Backend→DB | Keycloak | Fase 0 |
|:---|:---|:---|:---|:---|
| **EP-01** (Dashboard Admin) | `GET /dashboard/admin` | `SELECT` agregado em tenant | JWT validation | ✅ |
| **EP-02** (Clientes e Planos) | CRUD `/tenants`, `/plans`, `/subscriptions` | CRUD em tenant, plan, subscription | JWT validation + roles | ✅ |
| **EP-02** (Auditoria) | `GET /audit` | `SELECT` em audit_log | JWT validation | ✅ |
| **EP-03** (RBAC) | CRUD `/users`, `/permissions` | CRUD em user, user_permission, role_resource | JWT validation + roles | ✅ |
| **EP-04** (Portal Cliente) | CRUD `/business-units`, `/products`, `/onboarding` | CRUD em business_unit, product_service | JWT + onboarding flow | ✅ |
| **EP-04** (Dashboard Cliente) | `GET /dashboard/client` | `SELECT` agregado por BU | JWT validation | ✅ |

---

## 7. Tratamento de Erros entre Integrações

| Cenário | Comportamento Esperado |
|:---|:---|
| **Backend indisponível** | Frontend exibe tela de "Serviço temporariamente indisponível". Retry com exponential backoff (3 tentativas) |
| **PostgreSQL indisponível** | Backend retorna HTTP 503. Health check do K8s reinicia o pod se o banco não retornar em 30s |
| **Keycloak indisponível** | Usuários já autenticados (JWT válido) continuam operando normalmente. Novos logins exibem "Autenticação temporariamente indisponível" |
| **JWT expirado** | Backend retorna HTTP 401. Frontend redireciona para tela de login (ou tenta refresh token silenciosamente) |
| **JWT inválido (assinatura)** | Backend retorna HTTP 401. Log de segurança gerado. Sem refresh — força novo login |
| **Permissão negada (RBAC)** | Backend retorna HTTP 403. Frontend exibe tela de "Acesso Negado" (US-036) |
| **Timeout de query** | Backend retorna HTTP 504. Alerta de monitoramento dispara. Query é cancelada no PostgreSQL |

---

## 8. Próximos Passos de Detalhamento

1. Refinar schema SQL a partir do ERD base (TECHNICAL-PLAN.md §2.4) — criar migrations Flyway/Liquibase com índices, constraints e partial unique indexes
2. Gerar OpenAPI YAML (`fbso-platform-api.yaml`) a partir dos contratos definidos em API-CONTRACTS.md
3. Detalhar contratos de erro (RFC 7807 Problem Details) para todos os endpoints
4. Configurar health checks para cada integração (readiness/liveness probes K8s)
5. Definir estratégia de cache (Redis?) para queries de dashboard
6. Especificar logging estruturado (JSON) com trace_id cross-service
7. Planejar integração futura com RabbitMQ (tópicos, filas, DLQ)

---

## 9. Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial: matriz de integrações (8), diagrama de comunicação, fluxos de dados (tenant onboarding, RBAC, Multi-Tenant query), dependências de infraestrutura, tratamento de erros | Time Técnico |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 033-architecture-diagrams, 030-architecture-adr-general.*
🔍 *Revisado pelo skill caveman-review em 15/07/2026. Ajustes aplicados: senha Keycloak hardcoded removida (variável de ambiente), porta Keycloak 8443 (evita conflito com Spring Boot), role Admin FBSO documentada, próximos passos alinhados com artefatos existentes, isolation level para operações financeiras, version pinning PostgreSQL.*
