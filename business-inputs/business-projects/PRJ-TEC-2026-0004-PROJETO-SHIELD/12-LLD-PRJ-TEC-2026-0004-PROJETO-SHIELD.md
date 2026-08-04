# Low-Level Design (LLD): PROJETO SHIELD — ms-shield-identity-auth
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Solução Técnica** | ms-shield-identity-auth |
| **Documentos Base** | 01-PROJECT-CHARTER, 10-SAD, 11-HLD |
| **Stack** | Java 21 + Quarkus + GraalVM Native |
| **Data** | 03/08/2026 | **Versão** | 2.0 | **Metodologia** | WATERFALL |

---

## 1. Component Diagram (C4 Level 3) — Package Structure

```mermaid
classDiagram
    namespace web {
        class AuthController {
            +login(redirectUri) Redirect
            +callback(code, state) Redirect
            +logout() Redirect
            +refresh() Response
            +me() UserProfile
        }
        class HealthController {
            +health() HealthStatus
        }
        class MetricsController {
            +metrics() PrometheusData
        }
    }

    namespace service {
        class TenantResolverService {
            +resolveTenant(host) TenantMapping
            +invalidateCache(host) void
        }
        class OidcFlowService {
            +buildAuthUrl(realm, redirect, challenge) String
            +exchangeCode(realm, code, verifier) TokenSet
            +logout(realm, refreshToken) void
        }
        class SessionService {
            +createSession(profile, tokens) void
            +getProfile(sessionId) UserProfile
            +destroySession(sessionId) void
            +refreshSession(refreshToken) void
        }
        class TokenService {
            +exchangeTokens(realm, code, verifier) TokenSet
            +refreshTokens(realm, refreshToken) TokenSet
            +validateSession(accessToken) boolean
        }
        class AuditService {
            +logEvent(event) void
            +queryEvents(tenantId, filter) List~AuditEvent~
        }
    }

    namespace client {
        class KeycloakClient {
            +authorize(realm, params) Redirect
            +token(realm, body) TokenResponse
            +logout(realm, refreshToken) void
            +userinfo(realm, accessToken) UserInfo
        }
        class RedisClient {
            +get(key) String
            +set(key, value, ttl) void
            +del(key) void
        }
        class PostgresClient {
            +query(sql, tenantId) List~Row~
            +execute(sql, tenantId) void
            +withTenant(tenantId) Session
        }
    }

    namespace model {
        class TenantMapping {
            +String host
            +String realm
            +UUID tenantId
            +LocalDateTime cachedAt
        }
        class TokenSet {
            +String accessToken
            +String refreshToken
            +String idToken
            +int expiresIn
        }
        class UserProfile {
            +UUID userId
            +String email
            +List~String~ roles
            +UUID tenantId
        }
        class AuditEvent {
            +UUID eventId
            +UUID correlationId
            +UUID tenantId
            +String eventType
            +Instant timestamp
        }
    }

    AuthController --> TenantResolverService
    AuthController --> OidcFlowService
    AuthController --> SessionService
    TenantResolverService --> RedisClient
    OidcFlowService --> KeycloakClient
    SessionService --> PostgresClient
    SessionService --> TokenService
    TokenService --> KeycloakClient
    AuditService --> PostgresClient
    TenantResolverService --> TenantMapping
    OidcFlowService --> TokenSet
    SessionService --> UserProfile
    AuditService --> AuditEvent
```

---

## 2. API Contracts

```mermaid
flowchart LR
    subgraph Endpoints["🌐 REST Endpoints"]
        Login["GET /auth/login\n?redirect_uri=...\n→ 302 Keycloak + PKCE"]
        Callback["GET /auth/callback\n?code=...&state=...\n→ 302 App + Set-Cookie"]
        Logout["POST /auth/logout\nCookie: session\n→ 200 + Clear-Cookie"]
        Refresh["POST /auth/refresh\nCookie: refresh_token\n→ 200 + Set-Cookie"]
        Me["GET /auth/me\nCookie: session\n→ 200 UserProfile JSON"]
        Health["GET /health\n→ 200 {status, checks}"]
        Metrics["GET /metrics\n→ Prometheus text"]
    end

    style Login fill:#6f9,stroke:#333
    style Callback fill:#6f9,stroke:#333
    style Logout fill:#fc6,stroke:#333
    style Refresh fill:#fc6,stroke:#333
    style Me fill:#69f,stroke:#333
    style Health fill:#ccc,stroke:#333
    style Metrics fill:#ccc,stroke:#333
```

---

## 3. Database Schema

```mermaid
erDiagram
    user_sessions {
        uuid session_id PK "gen_random_uuid()"
        uuid user_id "NOT NULL"
        uuid tenant_id "NOT NULL — FK"
        text access_hash "NOT NULL"
        text refresh_hash "NOT NULL"
        timestamptz created_at "DEFAULT now()"
        timestamptz expires_at "NOT NULL"
        boolean revoked "DEFAULT false"
    }

    audit_events {
        uuid event_id PK "gen_random_uuid()"
        uuid correlation_id "NOT NULL"
        uuid tenant_id "NOT NULL"
        uuid user_id "nullable"
        varchar event_type "LOGIN|LOGOUT|REFRESH|FAILED|SUSPENDED"
        inet ip_address ""
        text user_agent ""
        timestamptz created_at "DEFAULT now()"
    }

    user_sessions ||--o{ audit_events : "triggers"
```

### RLS Policies

```sql
ALTER TABLE shield.user_sessions ENABLE ROW LEVEL SECURITY;
ALTER TABLE shield.audit_events   ENABLE ROW LEVEL SECURITY;

CREATE POLICY tenant_isolation_sessions ON shield.user_sessions
    FOR ALL USING (
        tenant_id = current_setting('app.current_tenant')::UUID
    );

CREATE POLICY tenant_isolation_audit ON shield.audit_events
    FOR ALL USING (
        tenant_id = current_setting('app.current_tenant')::UUID
    );
```

---

## 4. Sequence Diagrams

### Login Flow (Authorization Code + PKCE)

```mermaid
sequenceDiagram
    actor U as 👤 Browser
    participant CF as Cloudflare
    participant Kong as Kong Gateway
    participant BFF as AuthController
    participant TR as TenantResolver
    participant OIDC as OidcFlowService
    participant KC as Keycloak
    participant Redis as Redis
    participant SS as SessionService
    participant PG as PostgreSQL

    U->>CF: GET /app
    CF->>Kong: X-Tenant-Host: escola-alfa.com
    Kong->>BFF: GET /auth/login?redirect_uri=/app

    BFF->>TR: resolveTenant("escola-alfa.com")
    TR->>Redis: GET host:escola-alfa.com
    Redis-->>TR: realm-escola-alfa
    TR-->>BFF: TenantMapping(realm=realm-escola-alfa)

    BFF->>OIDC: buildAuthUrl(realm, redirectUri, challenge)
    OIDC-->>BFF: https://keycloak/realms/realm-escola-alfa/protocol/openid-connect/auth?code_challenge=...
    BFF-->>U: 302 → Keycloak

    U->>KC: Login form (credentials)
    KC-->>U: 302 → /auth/callback?code=xyz&state=abc

    U->>BFF: GET /auth/callback?code=xyz&state=abc
    BFF->>OIDC: exchangeCode(realm, code, verifier)
    OIDC->>KC: POST /token (code + code_verifier)
    KC-->>OIDC: {access_token, refresh_token, id_token}
    OIDC-->>BFF: TokenSet

    BFF->>SS: createSession(profile, tokens)
    SS->>PG: INSERT user_sessions (SET LOCAL tenant)
    SS-->>BFF: sessionId

    BFF-->>U: 302 /app + Set-Cookie
    Note over BFF,U: Cookie flags: HttpOnly, Secure, SameSite=Strict
```

### Cross-Tenant Attack Blocked

```mermaid
sequenceDiagram
    actor Attacker as 👤 Escola A (token)
    participant Kong as Kong Gateway
    participant BFF as AuthController
    participant PG as PostgreSQL (RLS)

    Attacker->>Kong: GET /api/v1/dados?escola=B
    Note over Attacker,Kong: Cookie: session (tenant=A)

    Kong->>Kong: Validate JWT → tenant=A, roles=[...]
    Kong->>BFF: Authorization: Bearer <JWT>

    BFF->>PG: SET LOCAL app.current_tenant = 'A'
    BFF->>PG: SELECT * FROM dados WHERE escola = 'B'
    Note over PG: RLS Policy: tenant_id = 'A'
    Note over PG: Row: tenant_id='B' → FILTERED OUT
    PG-->>BFF: 0 rows

    BFF-->>Attacker: 200 OK {data: []}
    Note over BFF,Attacker: Dados da Escola B NÃO vazaram
```

---

## 5. State Machine — User Session

```mermaid
stateDiagram-v2
    [*] --> Anonymous : Acessa /app

    Anonymous --> Redirecting : GET /auth/login
    Redirecting --> Authenticating : 302 → Keycloak
    Authenticating --> Active : /auth/callback OK
    Authenticating --> Failed : Invalid credentials
    Failed --> Redirecting : Retry

    Active --> Refreshing : POST /auth/refresh
    Refreshing --> Active : New tokens OK
    Refreshing --> Expired : Refresh token expired

    Active --> LoggedOut : POST /auth/logout
    LoggedOut --> [*] : Cookies cleared

    Active --> Blocked : Tenant suspended
    Blocked --> [*] : 403 Access Denied

    Expired --> Redirecting : Redirect to login
```

---

## 6. Error Handling Strategy

```mermaid
flowchart TB
    subgraph Exceptions["🚨 Exception Hierarchy"]
        TKE["TokenExchangeException\n502 — Keycloak unavailable"]
        TNE["TenantNotFoundException\n401 — Domain not mapped"]
        ISE["InvalidStateException\n401 — PKCE state mismatch"]
        SEE["SessionExpiredException\n401 — Token expired"]
        STE["SuspendedTenantException\n403 — Tenant blocked"]
        RE["RedisException\nWARN — Cache miss, fallback KC"]
        PE["PostgresException\n503 — DB unavailable, retry 3x"]
        GE["GenericException\n500 — Log + correlation_id"]
    end

    subgraph Mapper["🔄 GlobalExceptionMapper"]
        JAXRS["JAX-RS ExceptionMapper"]
    end

    subgraph Response["📤 Standard Error Response"]
        JSON["{error, message, correlation_id}"]
    end

    TKE --> JAXRS
    TNE --> JAXRS
    ISE --> JAXRS
    SEE --> JAXRS
    STE --> JAXRS
    RE --> JAXRS
    PE --> JAXRS
    GE --> JAXRS
    JAXRS --> JSON
```

---

## 7. Component Interfaces (Key)

```mermaid
classDiagram
    class TenantResolver {
        <<interface>>
        +resolveTenant(String host) Uni~TenantMapping~
        +invalidateCache(String host) Uni~Void~
    }

    class OidcFlow {
        <<interface>>
        +buildAuthorizationUrl(String realm, String redirectUri, String challenge) Uni~String~
        +exchangeCodeForTokens(String realm, String code, String verifier) Uni~TokenSet~
        +logout(String realm, String refreshToken) Uni~Void~
        +refreshTokens(String realm, String refreshToken) Uni~TokenSet~
    }

    class SessionManager {
        <<interface>>
        +createSession(UserProfile profile, TokenSet tokens) Uni~Void~
        +getProfile(String sessionId) Uni~UserProfile~
        +destroySession(String sessionId) Uni~Void~
        +refreshSession(String refreshToken) Uni~Void~
    }

    TenantResolver ..> TenantMapping : returns
    OidcFlow ..> TokenSet : returns
    SessionManager ..> UserProfile : returns
```

---

**[STATUS: SUCESSO]** — LLD v2.0 com diagramas Mermaid (classDiagram, erDiagram, sequenceDiagram ×2, stateDiagram-v2, flowchart).
