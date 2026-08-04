# Software Architecture Document (SAD): PROJETO SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Solução Técnica** | ms-shield-identity-auth |
| **Documentos Base** | 01-PROJECT-CHARTER, 02-BRD, 03-SRS |
| **Arquitetura Global** | `/home/bolismar/work/workspace-fbso/architecture/` |
| **Segurança Global** | `/home/bolismar/work/workspace-fbso/.specs/security/GLOBAL-SECURITY.md` |
| **Stack** | Java 21 + Quarkus + GraalVM Native + Keycloak + Kong + Istio + DOKS + PostgreSQL + Redis |
| **Data** | 03/08/2026 | **Versão** | 2.0 | **Metodologia** | WATERFALL |

---

## 1. Architectural Overview

**Estilo Arquitetural:** Microserviço de borda (Backend-For-Frontend) com arquitetura reativa e compilação nativa GraalVM.

```mermaid
flowchart TB
    subgraph Edge["🔒 Edge Layer"]
        CF["Cloudflare\nWAF + DNS + SSL"]
    end

    subgraph Gateway["🚪 API Gateway Layer"]
        Kong["Kong API Gateway\nRate Limiting + JWT Validation"]
    end

    subgraph Mesh["🔐 Service Mesh"]
        Istio["Istio\nmTLS + Traffic Control"]
    end

    subgraph Core["⚙️ Core Services"]
        BFF["ms-shield-identity-auth\nQuarkus Native\nPort 8080"]
        KC["Keycloak\nMulti-Realm\nOIDC Provider"]
    end

    subgraph Data["💾 Data Layer"]
        PG["PostgreSQL\nRLS Multi-Tenant"]
        Redis["Redis\nHost→Realm Cache"]
    end

    subgraph Obs["📊 Observability"]
        Prom["Prometheus"]
        Graf["Grafana"]
        Loki["Grafana Loki"]
        Jaeger["Jaeger"]
    end

    CF --> Kong
    Kong --> Istio
    Istio --> BFF
    BFF --> KC
    BFF --> Redis
    BFF --> PG
    BFF --> Prom
    Prom --> Graf
    BFF --> Loki
    BFF --> Jaeger
```

### ADR Registry

```mermaid
requirementDiagram
    requirement ADR_001 {
        id: "ADR-001"
        text: "GraalVM Native — binário nativo para latência <15ms e ~40MB RAM"
    }
    requirement ADR_002 {
        id: "ADR_002"
        text: "BFF único ponto Keycloak — microsserviços nunca chamam Keycloak"
    }
    requirement ADR_003 {
        id: "ADR_003"
        text: "Istio mTLS + Kong — comunicação interna criptografada"
    }
    requirement ADR_004 {
        id: "ADR_004"
        text: "Redis Host→Realm — cache com TTL 1h e invalidação sob demanda"
    }
    requirement ADR_005 {
        id: "ADR_005"
        text: "RLS Multi-Tenant — SET LOCAL app.current_tenant"
    }
    requirement ADR_006 {
        id: "ADR_006"
        text: "GitOps Argo CD — deploy declarativo via PR"
    }
    requirement ADR_007 {
        id: "ADR_007"
        text: "Cookies HttpOnly — tokens nunca expostos ao JS"
    }

    element NFR_Performance {
        type: "NFR-01,03,16"
    }
    element NFR_Security {
        type: "NFR-04 a 08"
    }
    element NFR_Availability {
        type: "NFR-09,11"
    }

    ADR_001 - satisfies -> NFR_Performance
    ADR_007 - satisfies -> NFR_Security
    ADR_006 - satisfies -> NFR_Availability
```

---

## 2. Solution Architecture

```mermaid
flowchart TB
    subgraph API["REST Layer"]
        Login["GET /auth/login"]
        Callback["GET /auth/callback"]
        Logout["POST /auth/logout"]
        Refresh["POST /auth/refresh"]
        Me["GET /auth/me"]
        Health["GET /health"]
        Metrics["GET /metrics"]
    end

    subgraph Services["Service Layer — CDI Beans"]
        TenantRes["TenantResolverService\nHost→Realm lookup"]
        OidcSvc["OidcFlowService\nAuth Code + PKCE"]
        SessionSvc["SessionService\nCookie management"]
        TokenSvc["TokenService\nToken exchange/refresh"]
        AuditSvc["AuditService\nEvent persistence"]
    end

    subgraph Clients["Integration Layer"]
        KCClient["KeycloakClient\nOIDC HTTP"]
        RedisClient["RedisClient\nReactive Cache"]
        PGClient["PostgresClient\nReactive SQL + RLS"]
    end

    Login --> TenantRes
    Callback --> OidcSvc
    Logout --> SessionSvc
    Refresh --> TokenSvc
    Me --> SessionSvc

    TenantRes --> RedisClient
    OidcSvc --> KCClient
    SessionSvc --> PGClient
    TokenSvc --> KCClient
    AuditSvc --> PGClient
```

```mermaid
flowchart LR
    Browser["🌐 Browser"] -->|"HTTPS + X-Tenant-Host"| CF["Cloudflare"]
    CF -->|"TLS 1.3"| Kong["Kong Gateway"]
    Kong -->|"HTTP/2 mTLS"| Istio["Istio Mesh"]
    Istio -->|"/auth/login"| Auth["AuthController"]
    Auth -->|"resolve(host)"| Redis["Redis"]
    Auth -->|"redirect OIDC"| KC["Keycloak"]
    KC -->|"callback?code=xyz"| Auth
    Auth -->|"exchange code"| KC
    Auth -->|"set cookies"| Browser
```

### System Integration Architecture — Shield × Sistemas Existentes

O Shield é a camada de identidade centralizada. Todos os sistemas existentes do ecossistema FBSO (Portal Escola, Portal Reforma, SaaS Corporativo, Comunidades de Ensino) delegam autenticação a ele. O contrato de integração é:

1. Cada sistema expõe suas páginas no domínio `*.fbso.org`
2. Ao detectar usuário sem sessão, o sistema redireciona para `shield.fbso.org/auth/login?redirect_uri=<url-de-retorno>`
3. O Shield gerencia todo o fluxo OIDC com Keycloak e retorna o usuário autenticado ao `redirect_uri`
4. Após retorno, o sistema chama `shield.fbso.org/auth/me` para obter o perfil
5. Cookies de sessão são compartilhados via domínio `.fbso.org` — login único para todos os sistemas

```mermaid
sequenceDiagram
    actor User as 👤 Usuário
    participant Sys as 🖥️ Portal Escola<br/>(sistema existente)
    participant CF as Cloudflare
    participant Kong as Kong Gateway
    participant BFF as Shield BFF
    participant Redis as Redis
    participant KC as Keycloak

    User->>Sys: Acessa escola-alfa.portal.fbso.org
    Sys->>Sys: Verifica cookie de sessão Shield
    Note over Sys: Sem sessão → redireciona
    Sys-->>User: 302 → shield.fbso.org/auth/login?redirect_uri=escola-alfa.portal.fbso.org

    User->>CF: GET shield.fbso.org/auth/login?redirect_uri=...
    CF->>Kong: X-Tenant-Host: escola-alfa.portal.fbso.org
    Kong->>BFF: GET /auth/login?redirect_uri=escola-alfa.portal.fbso.org

    BFF->>Redis: GET host:escola-alfa.portal.fbso.org
    Redis-->>BFF: realm-escola-alfa

    BFF-->>User: 302 → Keycloak /realms/realm-escola-alfa/auth?code_challenge=...
    User->>KC: Login form (credenciais)
    KC-->>User: 302 → shield.fbso.org/auth/callback?code=xyz&state=...

    User->>BFF: GET /auth/callback?code=xyz&state=...
    BFF->>KC: POST /token (code + code_verifier)
    KC-->>BFF: {access_token, refresh_token, id_token}

    BFF-->>User: 302 → escola-alfa.portal.fbso.org
    Note over BFF,User: Cookie: HttpOnly, Secure, SameSite=Strict, Domain=.fbso.org

    User->>Sys: GET escola-alfa.portal.fbso.org (com cookie)
    Sys->>BFF: GET shield.fbso.org/auth/me (com cookie)
    BFF-->>Sys: {user_id, email, roles, tenant_id}
    Sys-->>User: 🖥️ Aplicação renderizada

    Note over Sys,Kong: Chamadas API subsequentes passam pelo Kong<br/>Kong valida cookie → injeta Authorization: Bearer JWT<br/>com claims: tenant_id, roles, user_id
```

**Regra de Ouro da Integração:** Microserviços de negócio **NUNCA** chamam Keycloak diretamente. Toda validação de sessão é feita pelo Kong (que injeta o JWT) ou pelo Shield BFF (`/auth/me`).

---

## 3. Data Architecture

```mermaid
erDiagram
    Tenant {
        uuid tenant_id PK
        string domain "escola-alfa.com"
        string realm_id "realm-escola-alfa"
        string status "active|inactive|suspended"
    }

    UserSession {
        uuid session_id PK
        uuid user_id
        uuid tenant_id FK
        string access_hash
        string refresh_hash
        timestamp created_at
        timestamp expires_at
        boolean revoked
    }

    AuditEvent {
        uuid event_id PK
        uuid correlation_id
        uuid tenant_id
        uuid user_id
        string event_type "LOGIN|LOGOUT|REFRESH|FAILED"
        string ip_address
        string user_agent
        timestamp created_at
    }

    RealmMap {
        string host PK
        string realm
        uuid tenant_id
        timestamp cached_at
        int ttl
    }

    JWKSCache {
        string realm_id PK
        text jwks_json
        timestamp cached_at
        int ttl
    }

    Tenant ||--o{ UserSession : "has sessions"
    Tenant ||--o{ AuditEvent : "generates events"
    Tenant ||--|| RealmMap : "maps to"
    Tenant ||--|| JWKSCache : "has keys for"
```

---

## 4. Security Architecture

```mermaid
flowchart TB
    subgraph Edge["🛡️ Edge Security"]
        CF["Cloudflare\nWAF + DDoS + SSL"]
    end

    subgraph Gateway["🔑 Gateway Security"]
        Kong["Kong\nRate Limiting\nJWT Validation\nService-ID/Token-ID"]
    end

    subgraph Mesh["🔐 Zero-Trust Mesh"]
        Istio["Istio mTLS\nPod-to-Pod Encryption\nTraffic Control"]
    end

    subgraph App["🔒 Application Security"]
        BFF["Shield BFF\nInput Validation\nCookie HttpOnly/Secure/SameSite\nPKCE State Check"]
    end

    subgraph Identity["🪪 Identity Provider"]
        KC["Keycloak\nMulti-Realm Isolation\nOIDC + PKCE\nRBAC Roles"]
    end

    subgraph Data["🗄️ Data Security"]
        PG["PostgreSQL\nRLS per Tenant\nAES-256 Encryption\nZero PII in Logs"]
    end

    CF -->|"Header: X-Tenant-Host"| Kong
    Kong -->|"Bearer JWT + tenant_id"| Istio
    Istio -->|"mTLS"| BFF
    BFF -->|"OIDC Client Secret"| KC
    BFF -->|"SET LOCAL tenant"| PG
```

### Threat Model (STRIDE)

```mermaid
flowchart LR
    subgraph Threats["STRIDE Threat Categories"]
        S["Spoofing\nUsuário finge ser outro tenant"]
        T["Tampering\nToken JWT modificado"]
        R["Repudiation\nUsuário nega ação"]
        I["Info Disclosure\nCross-Tenant leak"]
        D["DoS\nBrute-force /auth/*"]
        E["Elevation\nAcesso a roles não autorizados"]
    end

    subgraph Mitigations["Mitigações"]
        M1["Cookie SameSite=Strict\nX-Tenant-Host via Cloudflare"]
        M2["JWT Signature validation\nKong verifica com Keycloak"]
        M3["AuditEvent imutável\ncorrelation_id + tenant_id"]
        M4["RLS: SET LOCAL tenant\nQuery retorna 0 linhas"]
        M5["Kong Rate Limiting\n100 req/min por IP"]
        M6["RBAC no Kong\nClaims do JWT + roles"]
    end

    S --> M1
    T --> M2
    R --> M3
    I --> M4
    D --> M5
    E --> M6
```

---

## 5. DevOps / SRE Architecture

```mermaid
flowchart LR
    subgraph CI["🔄 CI — GitHub Actions"]
        PR["Pull Request"]
        SAST["SAST\nSemgrep"]
        Secrets["Secret Scan\nGitleaks"]
        Build["GraalVM\nNative Build"]
        Test["Unit + IT\nTests"]
        Docker["Docker Build\nUBI-Micro"]
        Push["Push GHCR"]
    end

    subgraph CD["🚀 CD — GitOps"]
        GitOps["GitOps Repo\nUpdate Image Tag"]
        Argo["Argo CD\nSync DOKS"]
        Deploy["RollingUpdate\nZero Downtime"]
    end

    subgraph Obs["📊 Observability"]
        Metrics["Prometheus\nMicrometer"]
        Dash["Grafana\nDashboards"]
        Logs["Loki\nStructured JSON"]
        Trace["Jaeger\nDistributed Tracing"]
    end

    PR --> SAST --> Secrets --> Build --> Test --> Docker --> Push
    Push --> GitOps --> Argo --> Deploy
    Deploy --> Metrics --> Dash
    Deploy --> Logs
    Deploy --> Trace
```

### CI/CD Pipeline (Git Graph)

```mermaid
gitGraph
   commit id: "feat: initial scaffold"
   commit id: "feat: auth endpoints"
   branch feature/oauth2-pkce
   checkout feature/oauth2-pkce
   commit id: "feat: PKCE flow"
   commit id: "fix: state validation"
   checkout main
   merge feature/oauth2-pkce
   commit id: "feat: Redis cache"
   branch feature/rls-policies
   checkout feature/rls-policies
   commit id: "feat: RLS per tenant"
   commit id: "test: cross-tenant"
   checkout main
   merge feature/rls-policies
   commit id: "release: v1.0.0"
```

---

## 6. Infrastructure / Cloud Architecture

### Deployment Topology

```mermaid
flowchart TB
    subgraph DO["DigitalOcean — NYC3"]
        CF["Cloudflare Edge"]

        subgraph DOKS["DOKS Cluster — 3 nodes (4vCPU/8GB)"]
            IstioGW["Istio Ingress Gateway"]
            KongGW["Kong API Gateway"]
            subgraph ShieldNS["shield-system"]
                BFF["Shield BFF Pods\n2-50 (KEDA)"]
                KC["Keycloak Pods\n2 (StatefulSet)"]
            end
            Argo["Argo CD"]
            KEDA["KEDA Scaler"]
        end

        subgraph Managed["Managed Services"]
            PG["PostgreSQL HA\n2 nodes, 4GB"]
            RedisDO["Redis Managed\n2GB, TLS"]
        end

        subgraph Monitoring["monitoring"]
            Prom["Prometheus"]
            Graf["Grafana"]
        end
    end

    CF --> IstioGW
    IstioGW --> KongGW
    KongGW --> BFF
    BFF --> KC
    BFF --> PG
    BFF --> RedisDO
    BFF -.-> Prom --> Graf
    KEDA -.-> BFF
    Argo -.-> BFF
```

---

## 7. Testing Architecture

```mermaid
flowchart TB
    subgraph Pyramid["🧪 Testing Pyramid"]
        E2E["E2E — Cypress/Playwright\nFluxo de login completo\nCross-tenant validation"]
        Integration["Integration — RestAssured + k6\nAPI Contracts, Carga 200+ req/s\nRLS Cross-Tenant"]
        Unit["Unit — JUnit 5 + Mockito\nServices isolados\n>80% coverage"]
    end

    subgraph Security["🔒 Security Testing"]
        ZAP["OWASP ZAP\nAutomated Scan"]
        Manual["Manual Penetration\nCookie inspection\nCross-Tenant Data Leak"]
    end

    subgraph Gates["🚦 Quality Gates"]
        Gate1["Gate 1: Build\nCompile + Unit Tests"]
        Gate2["Gate 2: Integration\nAPI Tests + RLS"]
        Gate3["Gate 3: Security\nOWASP + Manual"]
        Gate4["Gate 4: Performance\nk6 Load + KEDA Scale"]
        Gate5["Gate 5: Go-Live\nAll gates pass"]
    end

    Unit --> Gate1
    Integration --> Gate2
    E2E --> Gate3
    Security --> Gate3
    ZAP --> Gate3
    Manual --> Gate3
    Gate2 --> Gate4
    Gate3 --> Gate5
    Gate4 --> Gate5
```

---

## 8. Cross-Cutting Concerns

| Concern | Implementação |
|---------|--------------|
| **Logging** | JSON estruturado: `{correlation_id, tenant_id, event, timestamp}` — Zero PII |
| **Error Handling** | Exceções mapeadas para HTTP status codes; stack traces nunca expostos |
| **Caching** | Redis: Host→Realm (TTL 1h), JWKS (TTL 15min). Invalidação sob demanda |
| **Rate Limiting** | Kong: 100 req/min/IP em /auth/*; 1000 req/min para endpoints internos |

---

## 9. Traceability SAD → SRS → BRD → Charter

| ADR (SAD) | NFR (SRS) | REQ (BRD) | OBJ (Charter) |
|-----------|-----------|-----------|---------------|
| ADR-001 GraalVM | NFR-01, NFR-03, NFR-16 | REQ-05, REQ-09 | C3, C4 |
| ADR-002 BFF único | NFR-04 a NFR-08 | REQ-02, REQ-03, REQ-04 | C1, C2 |
| ADR-003 Istio+Kong | NFR-07, NFR-09 | REQ-02, REQ-06 | C1, C4, C7 |
| ADR-004 Redis Cache | NFR-02 | REQ-01, REQ-08 | C1, C6 |
| ADR-005 RLS | NFR-08 | REQ-02 | C1 |
| ADR-006 GitOps | NFR-09, NFR-11 | REQ-09 | C4, C7 |
| ADR-007 Cookies | NFR-04 a NFR-06 | REQ-03 | C2 |

---

**[STATUS: SUCESSO]** — SAD v2.0 com diagramas Mermaid (flowchart, requirementDiagram, erDiagram, architecture-beta, gitGraph).
