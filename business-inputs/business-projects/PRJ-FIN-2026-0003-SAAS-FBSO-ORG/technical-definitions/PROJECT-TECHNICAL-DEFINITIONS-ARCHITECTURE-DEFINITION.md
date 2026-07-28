# PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION — Definição de Arquitetura do Projeto

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Programa:** FBSO Platform — Portal Administrativo SaaS
- **Versão:** 1.1
- **Data de Criação:** 26 de Julho de 2026
- **Última Atualização:** 27 de Julho de 2026 (alinhamento com docs de negócio v1.2)
- **Status:** ✅ COMPLIANCE — Validado pelo Time de Arquitetura
- **Baseline de Negócio:** [Project Charter v1.2](../01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [BRD v1.2](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [Épicos v1.2](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [Features FEAT-EP-](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)
- **Substitui:** `INTEGRATION-MAP.md` (absorvido neste documento)
- **Documentos Complementares:** [SOLUTIONS-CATALOG](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md) · [STACK-MATRIX](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md) · [PRD-DEFINITION](./PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md)

---

## 1. Objetivo

Este documento define a **arquitetura completa do projeto FBSO Platform** — como as 14 soluções técnicas se integram, comunicam e operam como um sistema coeso. Cobre os níveis C4 Level 1 (System Context) e Level 2 (Container), a matriz de integração entre todas as soluções, a topologia de deploy e os fluxos de sequência críticos.

---

## 2. C4 Level 1 — System Context

```mermaid
C4Context
    title FBSO Platform — System Context (C4 Level 1)

    Person(admin_fbso, "Administrador FBSO", "Time interno que gerencia o SaaS")
    Person(admin_tenant, "Administrador do Tenant", "Cliente que gerencia sua conta")
    Person(operador, "Operador", "Usuário operacional do cliente")

    System_Ext(cloudflare, "Cloudflare Edge", "CDN, WAF, SSL, DDoS Protection")
    System_Ext(email, "Email Service", "Envio de emails transacionais")

    System_Boundary(fbso_platform, "FBSO Platform — DigitalOcean") {
        System(portal_admin, "Portal Admin Interno", "Dashboard, tenants, planos, RBAC")
        System(portal_cliente, "Portal do Cliente", "Onboarding, catálogo, app switcher")
    }

    Rel(admin_fbso, portal_admin, "Acessa fbso.com", "HTTPS")
    Rel(admin_tenant, portal_cliente, "Acessa cliente.com", "HTTPS (white-label)")
    Rel(operador, portal_cliente, "Acessa", "HTTPS")

    Rel(portal_admin, cloudflare, "Passa por", "HTTPS/TLS")
    Rel(portal_cliente, cloudflare, "Passa por", "HTTPS/TLS")

    Rel(portal_admin, email, "Dispara emails", "SMTP")
    Rel(portal_cliente, email, "Dispara emails", "SMTP")
```

### 2.1 Atores Externos

| Ator | Tipo | Interação |
|:---|:---|:---|
| Administrador FBSO | Usuário interno | Acessa Portal Admin Interno |
| Administrador do Tenant | Usuário cliente | Acessa Portal do Cliente via domínio próprio (white-label) |
| Operador | Usuário cliente | Acessa Portal do Cliente com permissões restritas |
| Cloudflare Edge | Sistema externo | CDN, WAF, SSL termination, DDoS mitigation |
| Email Service | Sistema externo | Envio de emails transacionais (convite, recuperação de senha, notificações) |

---

## 3. C4 Level 2 — Container Diagram

```mermaid
C4Container
    title FBSO Platform — Container Diagram (C4 Level 2)

    Person(admin, "Administrador FBSO", "Time interno")
    Person(cliente, "Cliente (Admin Tenant)", "Auto-serviço")

    System_Ext(cloudflare, "Cloudflare Edge", "CDN + WAF + SSL")

    System_Boundary(fbso_platform, "FBSO Platform — DigitalOcean") {

        Container(kong, "Kong API Gateway", "Docker/Kong", "Gateway central: validação JWT, rate limiting, routing, header injection")
        Container(frontend, "Web App Portal", "Next.js 15 + React 19", "SPA com SSR — Portal Admin + Portal Cliente")
        Container(backend, "Backend API", "Java 25 + Spring Boot 3.5.14", "API REST monolítica modular. Lógica de negócio, multi-tenant, RBAC")
        Container(keycloak, "Keycloak IAM", "Keycloak 26.0", "Autenticação OIDC + emissão JWT. Realms por tenant.")
        Container(postgres, "PostgreSQL", "PostgreSQL 17 Alpine", "Banco relacional multi-tenant. Schemas: public, fbso_portal (RLS), keycloak")
        Container(otel, "OpenTelemetry Collector", "OTel Collector", "Coleta de traces e métricas")
        Container(grafana, "Grafana", "Grafana OSS", "Dashboards de monitoramento")

    }

    Rel(admin, cloudflare, "Acessa fbso.com", "HTTPS")
    Rel(cliente, cloudflare, "Acessa cliente.com", "HTTPS (white-label)")
    Rel(cloudflare, kong, "Encaminha", "HTTPS :443")

    Rel(kong, frontend, "Serve static + routes API", "HTTPS :3000 / :8080")
    Rel(kong, backend, "Routes API requests", "HTTP :8080")
    Rel(kong, keycloak, "Routes auth requests", "HTTP :8081")

    Rel(frontend, backend, "REST API (via Kong)", "JSON/HTTPS")
    Rel(frontend, keycloak, "OIDC Auth (via Kong)", "Authorization Code + PKCE")

    Rel(backend, postgres, "JDBC", "TCP :5432 — schema fbso_portal")
    Rel(backend, keycloak, "Validate JWT (JWKS)", "HTTPS")
    Rel(keycloak, postgres, "JDBC", "TCP :5432 — schema keycloak")

    Rel(backend, otel, "OTLP export", "gRPC :4317")
    Rel(kong, otel, "Metrics export", "Prometheus")
    Rel(otel, grafana, "Datasource", "HTTP :3001")
    Rel(grafana, postgres, "Business metrics", "TCP :5432")

    UpdateLayoutConfig($c4ShapeInRow="4", $c4BoundaryInRow="1")
```

### 3.1 Containers

| Container | Tecnologia | Porta | Responsabilidade |
|:---|:---|:---:|:---|
| **Kong API Gateway** | Kong (Docker/DO) | 443 | Validação JWT (JWKS), rate limiting, routing, header injection (`X-Tenant-ID`, `X-User-Permissions`) |
| **Web App Portal** | Next.js 15 + React 19 | 3000 | Frontend SPA/SSR. Serve assets estáticos + renderiza páginas. |
| **Backend API** | Java 25 + Spring Boot 3.5.14 | 8080 | API REST. Lógica de negócio. Multi-tenant via `SET app.current_tenant_id`. |
| **Keycloak IAM** | Keycloak 26.0 | 8081 | OIDC Provider. Realms por tenant. Emissão JWT com claims. |
| **PostgreSQL** | PostgreSQL 17 Alpine | 5432 | Persistência. Schemas: `public`, `fbso_portal` (RLS), `keycloak`. |
| **OTel Collector** | OpenTelemetry Collector | 4317/4318 | Coleta traces (gRPC) e métricas (Prometheus). |
| **Grafana** | Grafana OSS | 3001 | Dashboards: saúde, latência, métricas de negócio. |

---

## 4. Matriz de Integração

### 4.1 Matriz Completa (Origem → Destino)

| # | Origem | Destino | Protocolo | Porta | Autenticação | Dados Trafegados |
|:---|:---|:---|:---|:---:|:---|:---|
| I01 | Cloudflare | Kong | HTTPS/TLS | 443 | SSL (Cloudflare edge) | Requests HTTP |
| I02 | Kong | Frontend | HTTP | 3000 | — (interno) | Assets estáticos, páginas SSR |
| I03 | Kong | Backend | HTTP | 8080 | Header injection (X-Tenant-ID) | API requests, JSON |
| I04 | Kong | Keycloak | HTTP | 8081 | — (interno) | Auth requests |
| I05 | Frontend | Backend | REST/JSON | 80→Kong→8080 | Bearer JWT (via Kong) | CRUD tenants, plans, users, BUs, catalog |
| I06 | Frontend | Keycloak | OIDC | 443→Kong→8081 | Authorization Code + PKCE | Login, token refresh, logout |
| I07 | Backend | PostgreSQL | JDBC/PostgreSQL | 5432 | User/password (`fbso_app_user`) | Queries SQL no schema `fbso_portal` |
| I08 | Backend | Keycloak | HTTPS/JWKS | 8081 | — (valida assinatura JWT) | JWKS endpoint (public keys) |
| I09 | Keycloak | PostgreSQL | JDBC/PostgreSQL | 5432 | User/password (`fbso_keycloak_user`) | Tabelas Keycloak no schema `keycloak` |
| I10 | Backend | OTel Collector | OTLP/gRPC | 4317 | — (interno) | Traces, spans |
| I11 | Kong | OTel Collector | Prometheus | — | — (interno) | Métricas de API |
| I12 | OTel Collector | Grafana | HTTP | 3001 | — (interno) | Datasource traces/metrics |
| I13 | Grafana | PostgreSQL | JDBC | 5432 | Read-only user | Métricas de negócio (contas ativas, etc.) |
| I14 | Backend | MailHog | SMTP | 1025 | — (dev only) | Emails capturados (convite, reset senha) |

### 4.2 Padrões de Comunicação

| Padrão | Onde se Aplica | Justificativa |
|:---|:---|:---|
| **API Gateway** | Kong → Backend, Kong → Keycloak | Single entry point. Validação JWT centralizada. Rate limiting. |
| **REST/JSON** | Frontend → Backend (via Kong) | Padrão web. Contrato OpenAPI documentado. |
| **OIDC + PKCE** | Frontend → Keycloak | Authorization Code Flow com PKCE — mais seguro que Implicit Flow. |
| **JDBC Direto** | Backend → PostgreSQL, Keycloak → PostgreSQL | Conexão local no mesmo network. Pool de conexões gerenciado. |
| **JWKS Local** | Backend → Keycloak | Validação de assinatura JWT sem roundtrip para cada requisição. |
| **OTLP/gRPC** | Backend → OTel Collector | Protocolo padrão OpenTelemetry para tracing. |
| **SMTP (Dev)** | Backend → MailHog | Captura de emails em desenvolvimento. Produção: serviço externo. |

### 4.3 Matriz de Comunicação por Épico

| Épico | Frontend→Backend | Backend→DB | Keycloak | Fase 0 |
|:---|:---|:---|:---|:---|
| **EP-0001** (Dashboard Admin) | `GET /dashboard/admin` | `SELECT` agregado em tenant | JWT validation | ✅ |
| **EP-0002** (Clientes e Planos) | CRUD `/tenants`, `/plans`, `/subscriptions` | CRUD em tenant, plan, subscription | JWT validation + roles | ✅ |
| **EP-0002** (Auditoria) | `GET /audit` | `SELECT` em audit_log | JWT validation | ✅ |
| **EP-0003** (RBAC) | CRUD `/users`, `/permissions` | CRUD em user, user_permission, role_resource | JWT validation + roles | ✅ |
| **EP-0004** (Portal Cliente) | CRUD `/business-units`, `/products`, `/onboarding` | CRUD em business_unit, product_service | JWT + onboarding flow | ✅ |
| **EP-0004** (Dashboard Cliente) | `GET /dashboard/client` | `SELECT` agregado por BU | JWT validation | ✅ |

> **Origem:** Esta matriz foi migrada do INTEGRATION-MAP.md §6 (documento original absorvido e removido).

---

## 5. Topologia de Deploy

### 5.1 Ambiente de Desenvolvimento (Local)

```mermaid
flowchart TB
    subgraph docker["Docker Compose (fbso-network)"]
        PG["PostgreSQL :5432"]
        KC["Keycloak :8081"]
        MH["MailHog :1025/8025"]

        subgraph app["Aplicação (fora do compose)"]
            BE["Backend API :8080<br/>Spring Boot"]
            FE["Frontend :3000<br/>Next.js dev server"]
        end

        PG -->|"JDBC"| BE
        KC -->|"OIDC/JWT"| BE
        KC -->|"JDBC"| PG
        BE -->|"REST API"| FE
        BE -->|"SMTP"| MH
    end
```

| Container | Image | Porta | Usuário/Senha |
|:---|:---|:---:|:---|
| PostgreSQL | `postgres:17-alpine` | 5432 | `fbso_admin` / `fbso_admin` |
| Keycloak | `quay.io/keycloak/keycloak:26.0` | 8081 | `admin` / `admin` |
| MailHog | `mailhog/mailhog:v1.0.1` | 1025/8025 | — |

**Aplicação (fora do compose, via IDE/Maven):**
- Backend: `./mvnw spring-boot:run` → :8080
- Frontend: `npm run dev` → :3000

### 5.2 Ambiente de Produção (DigitalOcean)

```mermaid
flowchart TB
    subgraph cloudflare["Cloudflare Edge"]
        CF["CDN + WAF + SSL<br/>Custom Hostnames"]
    end

    subgraph doks["DigitalOcean Kubernetes (DOKS)"]
        subgraph gateway["API Gateway Layer"]
            KONG["Kong API Gateway<br/>:443 (externo)"]
        end

        subgraph apps["Application Layer"]
            FE["Frontend (Node.js)<br/>Next.js 15 :3000"]
            BE["Backend (Java 25)<br/>Spring Boot :8080"]
            KC["Keycloak IAM<br/>Keycloak 26 :8443"]
        end

        subgraph data["Data Layer"]
            PG["PostgreSQL 17<br/>Managed DO Database<br/>Schemas: public, fbso_portal, keycloak"]
        end

        subgraph observability["Observability"]
            OTEL["OTel Collector<br/>Traces + Metrics"]
            GRAF["Grafana<br/>Dashboards :3001"]
        end
    end

    CF -->|"HTTPS :443"| KONG
    KONG -->|"HTTP :3000"| FE
    KONG -->|"HTTP :8080"| BE
    KONG -->|"HTTP :8443"| KC
    BE -->|"JDBC :5432"| PG
    KC -->|"JDBC :5432"| PG
    BE -->|"OTLP gRPC :4317"| OTEL
    KONG -->|"Prometheus"| OTEL
    OTEL -->|"Datasource"| GRAF
    GRAF -->|"Read-only :5432"| PG
```

| Componente | Produção | Dev |
|:---|:---|:---|
| **Frontend** | Next.js build + Node.js no DOKS | `npm run dev` :3000 |
| **Backend** | GraalVM Native Image no DOKS | `./mvnw spring-boot:run` :8080 |
| **Kong** | Docker/DOKS :443 | Não usado em dev (opcional) |
| **Keycloak** | DOKS :8443 | Docker :8081 |
| **PostgreSQL** | DigitalOcean Managed Database | Docker :5432 |
| **OTel + Grafana** | DOKS sidecar | Docker (futuro) |
| **Rede** | DOKS internal network + Cloudflare | `fbso-network` bridge |

---

## 6. Estratégia de Comunicação

### 6.1 Síncrono (REST/HTTP)

| Fluxo | Protocolo | Timeout | Retry | Circuit Breaker |
|:---|:---|:---:|:---:|:---:|
| Frontend → Backend (via Kong) | REST/JSON | 30s | 3x (exponential backoff) | Sim (Kong) |
| Backend → Keycloak (JWKS) | HTTPS | 5s | 2x | Cache 15min |
| Backend → PostgreSQL | JDBC | 30s | Pool HikariCP | Connection pool |

### 6.2 Autenticação (OIDC + Kong → Backend)

```mermaid
sequenceDiagram
    actor User
    participant FE as Frontend (Next.js)
    participant KC as Keycloak
    participant Kong as Kong Gateway
    participant BE as Backend API

    Note over User,BE: === Fase 1: Autenticação OIDC ===

    User->>FE: Acessa portal
    FE->>KC: Redirect OIDC (Authorization Code + PKCE)
    KC-->>User: Tela de login (marca do tenant via Realm)
    User->>KC: Credenciais
    KC->>KC: Valida credenciais
    KC-->>FE: 302 /callback?code=xxx
    FE->>KC: POST /token (code + PKCE verifier)
    KC-->>FE: {access_token: JWT, id_token, refresh_token}

    Note over User,BE: === Fase 2: Requisição Autenticada via Kong ===

    User->>Kong: GET /api/tenants<br/>Authorization: Bearer <JWT>
    Kong->>KC: GET /.well-known/jwks (cache 15min)
    KC-->>Kong: JWKS keys
    Kong->>Kong: Valida assinatura JWT
    Kong->>Kong: Extrai claims: tenant_id, roles, permissions
    Kong->>Kong: Monta JWT interno com claims + acessos
    Kong->>BE: GET /api/tenants<br/>Authorization: Bearer <JWT interno><br/>Header: X-Tenant-ID, X-User-Permissions

    Note over BE: Backend "confia" no API Gateway.<br/>Sem double-check no Keycloak.<br/>Kong+Keycloak = camada de auth.

    BE->>BE: Abre JWT, extrai acessos+roles
    BE->>BE: Verifica o que os acessos+roles<br/>permitem obter dos serviços
    BE-->>Kong: 200 [tenants filtrados]
    Kong-->>User: 200 JSON
```

### 6.3 Assíncrono (Futuro — RabbitMQ)

| Evento | Publisher | Consumer | Quando |
|:---|:---|:---|:---|
| `tenant.activated` | Backend | Keycloak (Realm enable) | Tenant ativado |
| `tenant.suspended` | Backend | Keycloak (Realm disable) | Tenant suspenso |
| `subscription.changed` | Backend | Backend (módulos) | Upgrade/downgrade de plano |
| `user.invited` | Backend | Email Service | Convite de usuário |

> ⚠️ **Estado atual:** RabbitMQ (S10) é futuro. Eventos acima são implementados como chamadas síncronas diretas até que S10 seja ativado.

### 6.4 Tratamento de Erros entre Integrações

| Cenário | Comportamento Esperado |
|:---|:---|
| **Backend indisponível** | Frontend exibe tela de "Serviço temporariamente indisponível". Retry com exponential backoff (3 tentativas). |
| **PostgreSQL indisponível** | Backend retorna HTTP 503. Health check do K8s reinicia o pod se o banco não retornar em 30s. |
| **Keycloak indisponível** | Usuários já autenticados (JWT válido) continuam operando normalmente. Novos logins exibem "Autenticação temporariamente indisponível". |
| **JWT expirado** | Backend retorna HTTP 401. Frontend redireciona para tela de login (ou tenta refresh token silenciosamente). |
| **JWT inválido (assinatura)** | Backend retorna HTTP 401. Log de segurança gerado. Sem refresh — força novo login. |
| **Permissão negada (RBAC)** | Backend retorna HTTP 403. Frontend exibe tela de "Acesso Negado". |
| **Timeout de query** | Backend retorna HTTP 504. Alerta de monitoramento dispara. Query é cancelada no PostgreSQL. |

> **Origem:** Esta matriz de tratamento de erros foi migrada do INTEGRATION-MAP.md §7 (documento original absorvido e removido).

---

## 7. Diagramas de Sequência — Fluxos Críticos

### 7.1 Fluxo de Login OIDC + Primeira Requisição

```mermaid
sequenceDiagram
    actor User
    participant CF as Cloudflare
    participant Kong as Kong Gateway
    participant FE as Frontend (Next.js)
    participant KC as Keycloak
    participant BE as Backend API
    participant DB as PostgreSQL

    User->>CF: GET https://cliente.com
    CF->>Kong: Encaminha (SSL)
    Kong->>FE: Serve página (SSR)
    FE->>FE: header['host'] → domínio → tenant_id
    FE->>User: Página de login (marca tenant)

    User->>KC: POST /auth (credentials)
    KC->>DB: Valida credenciais (schema keycloak)
    KC-->>User: 302 /callback?code=xxx

    FE->>KC: POST /token (code + PKCE verifier)
    KC-->>FE: {access_token: JWT, refresh_token}

    User->>CF: GET /api/tenants (Bearer JWT)
    CF->>Kong: Encaminha
    Kong->>KC: GET /.well-known/jwks (cache 15min)
    KC-->>Kong: JWKS keys
    Kong->>Kong: Valida assinatura JWT
    Kong->>Kong: Extrai claims → headers
    Kong->>BE: GET /api/tenants (X-Tenant-ID, X-User-Permissions)
    BE->>DB: SET app.current_tenant_id = '<tenant_uuid>'
    BE->>DB: SELECT * FROM fbso_portal.tenants (RLS ativo)
    DB-->>BE: [tenant data]
    BE-->>Kong: 200 [tenants]
    Kong-->>CF: 200 [tenants]
    CF-->>User: JSON response
```

### 7.2 Fluxo de Onboarding de Novo Cliente

```mermaid
sequenceDiagram
    actor Admin as Admin FBSO
    participant FE as Portal Admin (S02)
    participant Kong as Kong Gateway
    participant BE as Backend API (S01)
    participant DB as PostgreSQL (S03)
    participant KC as Keycloak (S04)
    participant MH as MailHog (S07)

    Admin->>FE: Preenche formulário de novo tenant
    FE->>Kong: POST /api/tenants
    Kong->>BE: Forward + headers
    BE->>BE: Valida CNPJ (US-FEAT-EP-0002-0001-0010)
    BE->>DB: SET app.current_tenant_id = '00000000-...' (admin context)
    BE->>DB: INSERT INTO fbso_portal.tenants
    BE->>DB: INSERT INTO fbso_portal.business_units (primeira BU)
    DB-->>BE: tenant_id = <new_uuid>
    BE->>KC: POST /admin/realms (criar Realm)
    KC->>DB: Criar schema keycloak para o Realm
    KC-->>BE: Realm criado
    BE->>MH: Enviar email de boas-vindas
    MH-->>BE: Email capturado (dev)
    BE-->>Kong: 201 {tenant_id, credentials}
    Kong-->>FE: 201
    FE-->>Admin: Confirmação + credenciais geradas
```

### 7.3 Fluxo de Multi-Tenant Isolation (RLS)

```mermaid
sequenceDiagram
    actor UserA as Usuário Tenant A
    actor UserB as Usuário Tenant B
    participant Kong as Kong Gateway
    participant BE as Backend API
    participant DB as PostgreSQL

    par Requisição Tenant A
        UserA->>Kong: GET /api/business-units (JWT-A)
        Kong->>BE: Forward (X-Tenant-ID: A)
        BE->>DB: SET app.current_tenant_id = 'A'
        BE->>DB: SELECT * FROM fbso_portal.business_units
        Note over DB: RLS filtra: WHERE tenant_id = 'A'
        DB-->>BE: [BUs do Tenant A apenas]
        BE-->>UserA: 200 [3 BUs]
    and Requisição Tenant B
        UserB->>Kong: GET /api/business-units (JWT-B)
        Kong->>BE: Forward (X-Tenant-ID: B)
        BE->>DB: SET app.current_tenant_id = 'B'
        BE->>DB: SELECT * FROM fbso_portal.business_units
        Note over DB: RLS filtra: WHERE tenant_id = 'B'
        DB-->>BE: [BUs do Tenant B apenas]
        BE-->>UserB: 200 [5 BUs]
    end
```

### 7.4 Fluxo de Upgrade de Plano com Ativação de Módulos

```mermaid
sequenceDiagram
    actor Admin as Admin Tenant
    participant FE as Frontend (S02)
    participant Kong as Kong Gateway
    participant BE as Backend API (S01)
    participant DB as PostgreSQL (S03)
    participant KC as Keycloak (S04)
    actor Usuarios as Usuários do Tenant

    Note over Admin,Usuarios: Cenário: Tenant faz upgrade de plano (ex: Básico → Full Suite)

    Admin->>FE: Seleciona novo plano (Full Suite)
    FE->>Kong: POST /api/subscriptions/upgrade
    Kong->>BE: Forward + JWT interno
    BE->>BE: Valida: tenant ativo, upgrade permitido
    BE->>DB: SET app.current_tenant_id = '<tenant_uuid>'
    BE->>DB: UPDATE subscription SET plan_id = 'full_suite'
    BE->>DB: INSERT INTO audit_log (action='plan_upgraded')
    BE->>KC: Atualizar Realm: adicionar roles dos novos módulos

    Note over BE: "A se estudar forma segura e elegante<br/>e não intrusiva em sessões logadas<br/>para informar novos módulos"

    BE-->>FE: 200 {new_plan, activated_modules: [...]}
    FE-->>Admin: Confirmação + novos módulos visíveis

    Note over Usuarios,FE: Usuários logados precisam saber<br/>dos novos módulos sem logout forçado

    Usuarios->>FE: Próxima requisição (já logado)
    FE->>FE: Exibe notificação: "Novas funcionalidades<br/>disponíveis! Clique para atualizar."
    FE->>FE: Botão "Atualizar Acessos" disponível
    Usuarios->>FE: Clica "Atualizar Acessos"
    FE->>Kong: GET /api/me/permissions (force refresh)
    Kong->>BE: Forward
    BE->>BE: Re-consulta permissões atualizadas
    BE-->>FE: 200 {permissions: [..., new_modules]}
    FE->>FE: Re-renderiza menu com novos módulos
    FE-->>Usuarios: Menu atualizado + App Switcher com novos módulos
```

> ⚠️ **Estudo pendente:** Forma segura, elegante e não intrusiva de notificar usuários logados sobre novos módulos. Estratégia atual: notificação visual + botão "Atualizar Acessos" que força re-consulta de permissões e re-renderização da tela.

### 7.5 Fluxo de Downgrade de Plano com Inativação de Módulos

```mermaid
sequenceDiagram
    actor Admin as Admin Tenant
    participant FE as Frontend (S02)
    participant Kong as Kong Gateway
    participant BE as Backend API (S01)
    participant DB as PostgreSQL (S03)
    participant KC as Keycloak (S04)
    actor Usuario as Usuário do Tenant

    Note over Admin,Usuario: Cenário: Tenant faz downgrade (ex: Full Suite → Básico)

    Admin->>FE: Seleciona downgrade para plano Básico
    FE->>Kong: POST /api/subscriptions/downgrade
    Kong->>BE: Forward + JWT interno
    BE->>BE: Valida: downgrade permitido, módulos a remover
    BE->>DB: SET app.current_tenant_id = '<tenant_uuid>'
    BE->>DB: UPDATE subscription SET plan_id = 'basico'
    BE->>KC: Atualizar Realm: remover roles dos módulos inativos
    BE->>KC: Invalidar refresh tokens? (a estudar)
    BE->>DB: INSERT INTO audit_log (action='plan_downgraded')

    Note over BE: "A se estudar forma segura de<br/>bloquear acesso a menus e opções<br/>não mais permitidas"

    BE-->>FE: 200 {new_plan, deactivated_modules: [...]}
    FE-->>Admin: Confirmação

    Note over Usuario,FE: Usuário com tela renderizada antes do downgrade<br/>ainda vê menus/links dos módulos antigos

    Usuario->>FE: Clica em menu de módulo desativado
    FE->>Kong: GET /api/modulo-desativado/...
    Kong->>BE: Forward + JWT interno
    BE->>BE: Verifica permissões: módulo NÃO está mais ativo
    BE-->>FE: 403 {error: "Módulo não disponível no seu plano atual"}
    FE->>FE: Exibe mensagem de bloqueio amigável
    FE-->>Usuario: "Este módulo não está mais disponível<br/>no seu plano atual. Contate o administrador."

    Note over Usuario,FE: Próximo refresh de página ou "Atualizar Acessos"<br/>remove menus/links dos módulos inativos
```

> ⚠️ **Estudo pendente:** Forma segura de bloquear acesso a menus e opções não mais permitidas quando a tela foi renderizada antes do downgrade. Estratégia atual: backend retorna 403 com mensagem de bloqueio; frontend remove menus no próximo refresh ou botão "Atualizar Acessos".

### 7.6 Fluxo de Cancelamento de Plano com Logout Forçado

```mermaid
sequenceDiagram
    actor Admin as Admin Tenant
    participant FE as Frontend (S02)
    participant Kong as Kong Gateway
    participant BE as Backend API (S01)
    participant DB as PostgreSQL (S03)
    participant KC as Keycloak (S04)
    actor Usuarios as Todos os Usuários do Tenant

    Note over Admin,Usuarios: Cenário: Tenant cancela plano → todos perdem acesso

    Admin->>FE: Cancela assinatura
    FE->>Kong: POST /api/subscriptions/cancel
    Kong->>BE: Forward + JWT interno
    BE->>BE: Valida: tenant ativo, cancelamento permitido
    BE->>DB: SET app.current_tenant_id = '<tenant_uuid>'
    BE->>DB: UPDATE subscription SET status = 'cancelled'
    BE->>DB: UPDATE tenant SET status = 'suspended'
    BE->>DB: INSERT INTO audit_log (action='subscription_cancelled')

    Note over BE,KC: "A se estudar como derrubar<br/>usuários logados do tenant_id"

    BE->>KC: Desabilitar Realm do tenant
    KC->>KC: Revogar todas as sessões ativas do Realm
    KC->>KC: Invalidar todos os refresh tokens

    BE-->>FE: 200 {status: 'cancelled'}

    Note over Usuarios,FE: Todos os usuários do tenant são afetados

    Usuarios->>FE: Qualquer requisição (já logado)
    FE->>Kong: GET /api/... (JWT existente)
    Kong->>KC: Validar JWT
    KC-->>Kong: 401 — Realm desabilitado / token inválido
    Kong-->>FE: 401 Unauthorized
    FE->>FE: Limpa sessão local
    FE-->>Usuarios: Redireciona para tela de login<br/>"Sua conta foi desativada. Contate o suporte."
```

> ⚠️ **Estudo pendente:** Como "derrubar" usuários logados de um `tenant_id` específico. Estratégia atual: desabilitar Realm no Keycloak → revogar todas as sessões → invalidar refresh tokens. Próxima requisição de qualquer usuário recebe 401 e são redirecionados ao login com mensagem de conta desativada.

---

## 8. ADRs de Integração

### ADR-I01: Kong como API Gateway Central (Autenticação Delegada)

| Campo | Detalhe |
|:---|:---|
| **Decisão** | Toda requisição externa passa pelo Kong API Gateway. O Kong integra com o Keycloak, obtém a chave de acesso (token), valida o JWT e coloca as informações (tenant_id, roles, permissions) em um JWT interno no header da mensagem para o backend. |
| **Modelo de Confiança** | O backend recebe esse JWT interno e **confia** que o API Gateway realizou o trabalho de validação. NÃO está previsto que cada micro-serviço faça um "double-check" batendo novamente no Keycloak para validar o token. Kong+Keycloak formam a camada de autenticação/autorização. |
| **Fluxo** | Keycloak retorna se está autorizado → retorna o token válido → retorna os acessos+roles → Kong coloca essas informações no JWT → repassa ao backend → backend abre o JWT, captura os acessos+roles e verifica o que pode obter de seus serviços. |
| **Alternativas** | (1) NGINX reverse proxy, (2) Spring Cloud Gateway, (3) Istio Ingress, (4) Cada serviço validar token individualmente |
| **Justificativa** | Kong oferece plugin OIDC nativo, validação JWKS local (sem roundtrip), rate limiting declarativo. Delegar auth ao gateway simplifica o backend (sem lógica de validação JWT complexa). Spring Cloud Gateway exigiria lógica customizada em Java. Double-check no backend geraria latência desnecessária e carga dupla no Keycloak. |
| **Consequências** | Kong é ponto único de falha para autenticação — health checks + restart automático no DOKS. Latência adicional de ~5ms por requisição. Backend não tem lógica de validação de token independente (confia no Kong). |

### ADR-I02: JWKS Local Validation (não Introspection)

| Campo | Detalhe |
|:---|:---|
| **Decisão** | Kong valida JWT localmente via JWKS (chave pública), sem chamar o endpoint `/introspect` do Keycloak a cada requisição. |
| **Alternativas** | Token Introspection (RFC 7662) — Kong chama Keycloak a cada requisição. |
| **Justificativa** | JWKS é cacheado por 15 minutos. Latência de validação: <5ms vs ~50ms do introspection. Reduz carga no Keycloak. Revogação de token é tratada com short-lived access tokens (5min) + refresh tokens. |
| **Consequências** | Tokens revogados continuam válidos até expirar (max 5 min). Aceitável para o risco. |

### ADR-I03: Header Injection (não Token Forwarding)

| Campo | Detalhe |
|:---|:---|
| **Decisão** | Kong extrai claims do JWT e injeta como headers HTTP (`X-Tenant-ID`, `X-User-Permissions`, `X-User-Roles`). Backend NÃO recebe o JWT bruto. |
| **Alternativas** | Forward do JWT completo para o backend validar novamente. |
| **Justificativa** | Backend não precisa de lógica de validação JWT (simplifica código). Headers são confiáveis (rede interna DOKS). Spring Security configurado para confiar em headers pré-validados. |
| **Consequências** | Segurança depende da rede interna DOKS. Headers poderiam ser forjados se um container interno for comprometido. |

### ADR-I04: `SET app.current_tenant_id` por Transação

| Campo | Detalhe |
|:---|:---|
| **Decisão** | Backend executa `SET app.current_tenant_id = '<uuid>'` no início de cada transação JDBC. RLS no PostgreSQL usa `current_setting('app.current_tenant_id')` para filtrar. |
| **Alternativas** | (1) Adicionar `WHERE tenant_id = ?` em todas as queries manualmente, (2) Schema-per-tenant. |
| **Justificativa** | RLS garante que nem o owner da tabela escape do filtro (`FORCE ROW LEVEL SECURITY`). Zero risco de esquecer `WHERE tenant_id` em uma query. Schema-per-tenant não escala (milhares de schemas). |
| **Consequências** | `SET` deve ser a PRIMEIRA instrução de cada transação. Debug mais complexo (RLS oculta registros silenciosamente). |

### ADR-I05: Cloudflare + DigitalOcean (White-Label)

| Campo | Detalhe |
|:---|:---|
| **Decisão** | Cloudflare Edge (CDN, WAF, SSL, DDoS) → Kong → DigitalOcean (processamento). Custom Hostnames para domínios white-label. |
| **Alternativas** | (1) DigitalOcean CDN nativo, (2) AWS CloudFront + Route53 |
| **Justificativa** | Cloudflare Custom Hostnames permitem criar/remover domínios white-label programaticamente via API. WAF e DDoS na borda protegem cada domínio de cliente sem custo adicional por tenant. AWS seria alternativa viável mas DigitalOcean já é o provedor de infra. |
| **Consequências** | Dependência de dois provedores (Cloudflare + DigitalOcean). Latência adicional de ~10ms na borda. Custo varia com número de Custom Hostnames. |

---

## 9. Referências

| Documento | Relação |
|:---|:---|
| [SOLUTIONS-CATALOG](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md) | 14 soluções catalogadas |
| [STACK-MATRIX](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md) | Stacks por solução |
| [PRD-DEFINITION](./PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md) | Baseline de produto |
| [STACK-MATRIX](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md) | Stack tecnológica e ADRs de decisões |
| [PRD-DEFINITION](./PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md) | Baseline de requisitos de produto |
| `/architecture/` | ADRs globais, blueprints, data standards |
| [docker-compose.yml](../../../backend/java/spring/microservices/ms-fbso-platform-admin/docker-compose.yml) | Topologia dev |

---

## Histórico de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.1 | 26/07/2026 | Correções pós-validação humana: (1) seção 5.2 migrada para flowchart Mermaid, (2) seção 6.2 migrada para sequence diagram Mermaid, (3) ADR-I01 atualizado com modelo de confiança Kong+Keycloak (sem double-check no backend), (4) adicionados 3 diagramas de sequência: upgrade de plano com ativação de módulos (7.4), downgrade com inativação (7.5), cancelamento com logout forçado (7.6). 3 estudos pendentes documentados (notificação de novos módulos, bloqueio de acesso, derrubar sessões). | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Diagramas C4 em Mermaid. Resultado da Fase 5 do Roadmap de Definições Técnicas.*
