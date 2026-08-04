# Low-Level Design (LLD): PROJETO SHIELD — ms-shield-identity-auth
## [STATUS: Em revisão]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Solução Técnica** | ms-shield-identity-auth |
| **Documentos Base** | 01-PROJECT-CHARTER, 10-SAD, 11-HLD |
| **Stack** | Java 21 + Quarkus + GraalVM Native |
| **Data** | 03/08/2026 | **Versão** | 3.0 — Revisão de Integração | **Metodologia** | WATERFALL |

---

## 1. Component Diagram — Package Structure (Kong Filter Model)

O Shield atua como **serviço de validação de sessão acoplado ao Kong**. A SPA nunca chama o Shield diretamente — o Kong encaminha a validação de cookie para o Shield, que decide: injetar JWT (sessão válida) ou redirecionar para Keycloak (sem sessão).

```mermaid
classDiagram
    class KongRouter {
        +route(api_request) SessionCheck
    }
    class JWTInjector {
        +inject(jwt_claims) AuthorizationHeader
    }
    class SessionFilter {
        +validateCookie(cookie) SessionStatus
        +injectJWT(session) JWTClaims
        +redirectToLogin(host) Redirect
    }
    class TenantResolver {
        +resolveTenant(host) RealmMapping
        +invalidateCache(host) void
    }
    class OidcFlowService {
        +handleCallback(code, state) TokenSet
        +exchangeCode(realm, code, verifier) TokenSet
    }
    class SessionStore {
        +save(sessionId, jwt, ttl) void
        +get(sessionId) SessionData
        +delete(sessionId) void
    }
    class AuditService {
        +logEvent(event) void
    }
    class KeycloakClient {
        +authorize(realm, challenge) Redirect
        +token(realm, code, verifier) TokenResponse
        +logout(realm, refreshToken) void
    }
    class RedisSessionStore {
        +SET session_id jwt EX ttl
        +GET session_id
        +DEL session_id
    }
    class RedisHostCache {
        +GET host
        +SET host realm EX ttl
    }
    class SessionData {
        +String accessToken
        +String refreshToken
        +UUID tenantId
        +List~String~ roles
        +UUID userId
        +Instant expiresAt
    }
    class JWTClaims {
        +UUID tenantId
        +List~String~ roles
        +UUID userId
        +String email
    }

    KongRouter --> SessionFilter : "api + cookie"
    SessionFilter --> TenantResolver : "host lookup"
    SessionFilter --> SessionStore : "validate sessao"
    SessionFilter --> JWTInjector : "injeta claims"
    SessionStore --> RedisSessionStore
    TenantResolver --> RedisHostCache
    SessionFilter --> OidcFlowService : "callback"
    OidcFlowService --> KeycloakClient
    SessionStore --> AuditService
    SessionData --> JWTClaims
```

### Fluxo de Decisão do SessionFilter

```mermaid
flowchart TD
    A["Kong recebe GET /api/v1/alunos + cookie"] --> B{"Cookie SHIELD_SESSION presente?"}
    B -->|"Sim"| C["SessionStore.get(sessionId)"]
    C --> D{"JWT válido e não expirado?"}
    D -->|"Sim"| E["Extrai JWTClaims\ninjeta Authorization: Bearer <JWT>\nencaminha para microserviço"]
    D -->|"Expirado"| F["Tenta refresh token\nvia Keycloak"]
    F --> G{"Refresh OK?"}
    G -->|"Sim"| H["SessionStore.save(novo JWT)\ninjeta Authorization"]
    G -->|"Falhou"| I["302 → Keycloak login"]
    B -->|"Não"| J["TenantResolver.resolveTenant(host)"]
    J --> K["302 → Keycloak /realms/{realm}/auth?code_challenge=..."]
```

---

## 2. Endpoints — Uso Interno (Kong)

Estes endpoints são chamados **apenas pelo Kong API Gateway**, nunca diretamente pelo frontend:

| Endpoint | Quem Chama | Função |
|----------|-----------|--------|
| `POST /internal/session/validate` | Kong | Recebe cookie SHIELD_SESSION → retorna JWT claims (válido) ou 401 (expirado/inválido) |
| `GET /internal/tenant/resolve?host=` | Kong | Recebe host → retorna realm_id (para construir redirect Keycloak) |
| `GET /auth/callback?code=&state=` | Browser (via redirect Keycloak) | Recebe authorization code → troca por tokens → salva no Redis → seta cookie → redirect |
| `GET /health` | Kong / Prometheus | Health check |
| `GET /metrics` | Prometheus | Métricas operacionais |

**Nota:** Endpoints como `/auth/login`, `/auth/logout`, `/auth/refresh`, `/auth/me` **não existem mais como API pública**. O fluxo de login é disparado pela interceptação do Kong (passo 3 do SessionFilter). O logout é gerenciado pelo cookie — a SPA apenas remove o cookie local. O perfil do usuário está nas claims do JWT injetado — o microserviço de negócio as extrai diretamente.

---

## 3. Database Schema (inalterado — mantido para referência)

```mermaid
erDiagram
    user_sessions {
        uuid session_id PK "gen_random_uuid()"
        uuid user_id "NOT NULL"
        uuid tenant_id "NOT NULL"
        text access_hash "NOT NULL"
        timestamptz created_at "DEFAULT now()"
        timestamptz expires_at "NOT NULL"
        boolean revoked "DEFAULT false"
    }

    audit_events {
        uuid event_id PK "gen_random_uuid()"
        uuid correlation_id "NOT NULL"
        uuid tenant_id "NOT NULL"
        varchar event_type "LOGIN|LOGOUT|REFRESH|FAILED|SUSPENDED"
        timestamptz created_at "DEFAULT now()"
    }

    user_sessions ||--o{ audit_events : "triggers"
```

---

## 4. Sequence Diagrams

### 4.1 Login — SPA faz chamada API, Kong+Shield interceptam

```mermaid
sequenceDiagram
    actor User as 👤 Browser
    participant SPA as 🖥️ SPA Frontend
    participant CF as Cloudflare
    participant Kong as Kong Gateway
    participant Shield as SessionFilter
    participant Redis as Redis
    participant KC as Keycloak
    participant MS as Microserviço Negócio
    participant PG as PostgreSQL

    User->>CF: https://escola-alfa.com
    CF->>SPA: Entrega HTML/JS/CSS
    SPA-->>User: SPA carregada

    User->>CF: GET /api/v1/alunos (SPA)
    CF->>Kong: Proxy API + cookie SHIELD_SESSION (se existir)

    Kong->>Shield: POST /internal/session/validate (cookie)

    alt Sem cookie (primeira visita)
        Shield->>Redis: GET host:escola-alfa.com
        Redis-->>Shield: realm-escola-alfa
        Shield-->>Kong: 302 → Keycloak /realms/realm-escola-alfa/auth
        Kong-->>User: 302 → Login form (tema da escola)

        User->>KC: Credenciais
        KC-->>User: 302 → /auth/callback?code=xyz&state=abc
        User->>Shield: GET /auth/callback?code=xyz&state=abc
        Shield->>KC: POST /token (code + code_verifier)
        KC-->>Shield: {access_token, refresh_token, id_token}
        Shield->>Redis: SET session_id {jwt, claims, ttl}
        Shield-->>User: 302 → escola-alfa.com
        Note over Shield,User: Set-Cookie: SHIELD_SESSION=..., HttpOnly, Secure, SameSite=Strict
    else Com cookie valido
        Shield->>Redis: GET session_id
        Redis-->>Shield: {jwt, claims: tenant_id, roles, user_id}
        Shield-->>Kong: 200 {claims, jwt}
        Kong->>Kong: JWTInjector → Authorization: Bearer <JWT>
        Kong->>MS: GET /api/v1/alunos + Authorization header
        MS->>PG: SET LOCAL app.current_tenant = '{tenant_id}'
        PG-->>MS: Dados filtrados (RLS)
        MS-->>User: 200 {data: [...]}
    end
```

### 4.2 Cookie Expirado — Refresh ou Re-login

```mermaid
sequenceDiagram
    participant Kong as Kong Gateway
    participant Shield as SessionFilter
    participant Redis as Redis
    participant KC as Keycloak

    Kong->>Shield: POST /internal/session/validate (cookie)
    Shield->>Redis: GET session_id
    Redis-->>Shield: {jwt: expirado, refresh_token}
    Shield->>KC: POST /token (grant_type=refresh_token)
    alt Refresh OK
        KC-->>Shield: {new_access_token, new_refresh_token}
        Shield->>Redis: SET session_id {new_jwt, new_ttl}
        Shield-->>Kong: 200 {claims, jwt}
    else Refresh expirado
        KC-->>Shield: 400 invalid_grant
        Shield->>Redis: DEL session_id
        Shield-->>Kong: 302 → Keycloak /auth
    end
```

---

## 5. State Machine — Sessão do Usuário

```mermaid
stateDiagram-v2
    [*] --> Anonymous : Acessa SPA

    Anonymous --> SPA_Loaded : Cloudflare entrega HTML/JS/CSS
    SPA_Loaded --> Calling_API : SPA faz GET /api/v1/alunos
    Calling_API --> Redirecting : Kong+Shield: sem cookie → 302 Keycloak
    Calling_API --> Authorized : Kong+Shield: cookie válido → injeta JWT → MS recebe

    Redirecting --> Authenticating : Usuário vê login form
    Authenticating --> Callback : Keycloak redireciona com code
    Callback --> Session_Stored : Shield troca code → tokens → Redis
    Session_Stored --> SPA_Loaded : 302 → escola-alfa.com

    Authorized --> Calling_API : Próxima chamada API (cookie existe)

    Calling_API --> Refreshing : Cookie expirado → tenta refresh
    Refreshing --> Authorized : Refresh OK → novo JWT no Redis
    Refreshing --> Redirecting : Refresh falhou → 302 Keycloak

    Calling_API --> Session_Deleted : Logout / sessão revogada
    Session_Deleted --> Anonymous : Cookie removido
```

---

## 6. Error Handling Strategy

| Situação | Onde Ocorre | Resposta |
|----------|------------|---------|
| Cookie ausente | SessionFilter.validateCookie() | Kong retorna 302 → Keycloak login |
| Cookie expirado, refresh OK | SessionFilter + KeycloakClient | Novo JWT armazenado no Redis, Kong injeta Authorization |
| Cookie expirado, refresh falhou | SessionFilter + KeycloakClient | Kong retorna 302 → Keycloak login |
| JWT inválido (assinatura/tampering) | Kong JWTInjector | 401 — descartar cookie |
| Host não mapeado | TenantResolver.resolveTenant() | 401 — "Domínio não configurado" |
| Keycloak indisponível | KeycloakClient.token() | 503 — "Serviço de autenticação temporariamente indisponível" |
| Redis indisponível | SessionStore.get() | Degradado: validar JWT localmente (cache em memória, TTL curto). Alertar |
| Tenant suspenso | SessionFilter.validateCookie() | 403 — tenant_id bloqueado |

---

## 7. Component Interfaces (Key)

```mermaid
classDiagram
    class SessionFilter {
        <<interface>>
        +validate(cookie) SessionResult
    }

    class SessionStore {
        <<interface>>
        +save(sessionId, jwt, ttl) void
        +get(sessionId) SessionData
        +delete(sessionId) void
        +refresh(sessionId, refreshToken) SessionData
    }

    class TenantResolver {
        <<interface>>
        +resolve(host) RealmMapping
    }

    class JWTClaims {
        +UUID tenantId
        +List~String~ roles
        +UUID userId
        +String email
    }

    SessionFilter ..> SessionStore : uses
    SessionFilter ..> TenantResolver : uses
    SessionFilter ..> JWTClaims : produces
    SessionStore ..> SessionData : returns
```

### Kickback para o Kong

O SessionFilter retorna ao Kong uma destas 3 respostas:

| Resposta | Código | Ação do Kong |
|----------|--------|-------------|
| `{action: "INJECT", claims: {...}, jwt: "..."}` | 200 | JWTInjector insere `Authorization: Bearer <jwt>`, encaminha para MS |
| `{action: "REDIRECT", location: "https://keycloak/.../auth"}` | 302 | Kong retorna redirect HTTP para o browser |
| `{action: "REJECT", status: 403, reason: "tenant_suspended"}` | 403 | Kong retorna 403 para o browser |

---

**[STATUS: SUCESSO]** — LLD v3.0 alinhado com arquitetura Kong Filter. Shield não expõe endpoints REST públicos. SessionFilter decide: injetar JWT ou redirecionar.
