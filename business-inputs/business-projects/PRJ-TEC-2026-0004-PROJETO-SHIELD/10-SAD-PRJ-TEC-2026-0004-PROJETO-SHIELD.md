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

**Estilo Arquitetural:** Microserviço de borda (Backend-For-Frontend) com arquitetura reativa e compilação nativa GraalVM. O Shield atua como **serviço de validação de sessão acoplado ao Kong API Gateway** — não é chamado diretamente pelo frontend.

```mermaid
flowchart TB
    User["👤 Usuário\n(navegador)"]

    subgraph Edge["🔒 Edge Layer"]
        CF["Cloudflare\nWAF + DNS + SSL + Proxy"]
    end

    subgraph Frontend["🖥️ Frontend (App Platform / DOKS Nginx)"]
        SPA["SPA Estática\nHTML/JS/CSS\n(Zero secrets)"]
    end

    subgraph Gateway["🚪 API Gateway + Auth Filter"]
        Kong["Kong API Gateway\nRate Limiting + Routing"]
        Shield["ms-shield-identity-auth\nValidação de Sessão\nInjeção de JWT"]
    end

    subgraph Mesh["🔐 Service Mesh"]
        Istio["Istio\nmTLS + Traffic Control"]
    end

    subgraph Core["⚙️ Backend Services"]
        MS["Microserviços de Negócio\nms-escolas-core\n(Portal, Reforma, SaaS...)"]
        KC["Keycloak\nMulti-Realm\nOIDC Provider"]
    end

    subgraph Data["💾 Data Layer"]
        PG["PostgreSQL\nRLS Multi-Tenant"]
        Redis["Redis\nSession Store + Host→Realm Cache"]
    end

    subgraph Obs["📊 Observability"]
        Prom["Prometheus"]
        Graf["Grafana"]
    end

    User -->|"1. https://escola-alfa.com"| CF
    CF -->|"2. Proxy estático"| SPA
    SPA -.->|"3. GET /api/v1/alunos"| CF
    CF -->|"4. Proxy de API"| Kong
    Kong -->|"5. Validar sessão"| Shield
    Shield -->|"6. Host→Realm"| Redis
    Shield -.->|"7a. Sem sessão: 302 → Keycloak"| KC
    Shield -->|"7b. Com sessão: injeta JWT"| Kong
    Kong -->|"8. Authorization: Bearer <JWT>"| Istio
    Istio --> MS
    MS -->|"SET LOCAL tenant"| PG
    MS -.-> Prom --> Graf
    
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

O `ms-shield-identity-auth` expõe endpoints REST mas **não é chamado diretamente pelo frontend**. O Kong API Gateway é o entry point — ele encaminha a validação de sessão para o Shield, que decide se libera a requisição (injetando JWT) ou redireciona para o Keycloak.

```mermaid
flowchart TB
    subgraph Kong["Kong API Gateway"]
        direction LR
        KongRouter["Router\n/api/* → Shield valida\n/auth/* → Shield processa"]
        KongJWT["JWT Injector\n(plugin Shield)"]
    end

    subgraph Shield["ms-shield-identity-auth"]
        subgraph API["REST Endpoints"]
            Callback["/auth/callback\nTroca code por tokens"]
            Session["/auth/session\nValida cookie. Injeta JWT"]
        end

        subgraph Services["Service Layer"]
            TenantRes["TenantResolver\nHost→Realm"]
            OidcSvc["OidcFlowService\nAuth Code + PKCE"]
            SessionSvc["SessionService\nCookie + JWT storage"]
        end

        subgraph Clients["Integration Layer"]
            KCClient["KeycloakClient"]
            RedisClient["RedisClient\nSession Store"]
        end
    end

    KongRouter --> Session
    Session --> SessionSvc
    SessionSvc --> RedisClient
    SessionSvc -->|"Com sessão: injeta JWT"| KongJWT
    Session -->|"Sem sessão: 302"| TenantRes
    TenantRes --> RedisClient
    Callback --> OidcSvc
    OidcSvc --> KCClient

    KongJWT -->|"Authorization: Bearer <JWT>\nclaims: tenant_id, roles, user_id"| MS["Microserviços de Negócio"]
```

### Passo a Passo — Interceptação de API pelo Shield

**1. Acesso e Entrega do Frontend:**
- O cliente digita `https://escola-alfa.com`
- A Cloudflare recebe e repassa para a DigitalOcean (App Platform ou DOKS/Nginx), que entrega os arquivos estáticos (HTML/JS/CSS) da SPA para o navegador
- O frontend é **100% estático e sem segredos** — não conhece client_secret do Keycloak

**2. Tentativa de Chamada de API:**
- O JavaScript carrega no navegador e faz uma requisição de dados (ex: `GET /api/v1/alunos`)
- A chamada passa pela Cloudflare (proxy de API) e chega ao Kong API Gateway no DOKS

**3. Interceptação pelo Shield (onde ele atua):**
- O Kong repassa a checagem de sessão para o `ms-shield-identity-auth`
- **Se o usuário NÃO estiver logado:** O Shield consulta Redis (host → realm), descobre o Realm no Keycloak baseado na URL (`escola-alfa.com`) e devolve `302` para a tela de login do Keycloak
- **Se o usuário ESTIVER logado:** O navegador envia o cookie `HttpOnly`. O Shield decodifica a sessão, recupera o JWT do Redis e **injeta o JWT no header `Authorization: Bearer <JWT>`** da requisição interna, liberando a chamada para o microserviço de negócio

**4. Autenticação no Keycloak:**
- O cliente digita usuário e senha na tela do Keycloak (tema visual da escola)
- O Keycloak valida as credenciais e devolve um *Authorization Code* para a URL de callback (`/auth/callback`)

**5. Troca do Código por Tokens (Back-Channel):**
- A rota `/auth/callback` é processada pelo Shield
- O Shield faz a chamada interna (back-channel) para o Keycloak, trocando o código pelo **Access Token (JWT)** e **Refresh Token**
- Em vez de devolver o JWT para o frontend, o Shield **armazena os tokens no Redis** e responde ao navegador gravando o cookie seguro: `Set-Cookie: SHIELD_SESSION=...; HttpOnly; Secure; SameSite=Strict`

**6. Retorno ao Cliente:**
- A resposta atravessa DOKS → Cloudflare → Navegador
- O cliente está com a SPA carregada na URL `escola-alfa.com`, com sessão gravada via cookie seguro e **pronto para usar a aplicação**
- Nas próximas chamadas de API, o fluxo repete a partir do passo 2 — mas agora o cookie existe, então o Shield vai direto para a injeção do JWT (passo 3, caminho "com sessão")

### System Integration Architecture — Shield como Filtro de Sessão do Kong

O `ms-shield-identity-auth` **não é chamado diretamente pelo frontend**. Ele atua como um serviço de validação de sessão acoplado ao Kong API Gateway. O frontend (SPA estática) faz chamadas de API normalmente — o Kong + Shield interceptam e gerenciam a autenticação de forma transparente.

**Por que esta separação:**
- **Frontend 100% estático:** A SPA no App Platform/Nginx não conhece segredos do Keycloak (client_secret). Ela apenas faz chamadas de API.
- **Isolamento de responsabilidade:** O Shield cuida de OIDC, cookies HttpOnly, gestão de sessão e injeção de JWT. O frontend cuida de UI/UX.
- **Escala independente:** Se o volume de logins disparar, o KEDA escala apenas as instâncias do Shield (Quarkus Native, ~40MB RAM, cold start <100ms).

```mermaid
sequenceDiagram
    actor User as 👤 Usuário
    participant CF as Cloudflare
    participant SPA as 🖥️ Frontend SPA<br/>(App Platform / Nginx)
    participant Kong as Kong API Gateway
    participant Shield as Shield BFF
    participant Redis as Redis
    participant KC as Keycloak
    participant MS as Microserviço Negócio<br/>(ms-escolas-core)
    participant PG as PostgreSQL (RLS)

    User->>CF: https://escola-alfa.com
    CF->>SPA: Proxy → entrega arquivos estáticos (HTML/JS/CSS)
    SPA-->>User: 🖥️ SPA carregada no navegador

    User->>CF: GET /api/v1/alunos (chamada de API da SPA)
    CF->>Kong: Proxy de API → encaminha requisição
    Kong->>Shield: Valida sessão (cookie SHIELD_SESSION)

    alt Sem sessão
        Shield->>Redis: GET host:escola-alfa.com
        Redis-->>Shield: realm-escola-alfa
        Shield-->>Kong: 302 → Keycloak /realms/realm-escola-alfa/auth
        Kong-->>User: 302 → Login form (tema da escola)

        User->>KC: Usuário + senha
        KC-->>User: 302 → /auth/callback?code=xyz
        User->>Shield: GET /auth/callback?code=xyz
        Shield->>KC: POST /token (code → troca por tokens)
        KC-->>Shield: {access_token, refresh_token, id_token}
        Shield->>Redis: Armazena JWT (key: session_id)
        Shield-->>User: 302 → escola-alfa.com
        Note over Shield,User: Set-Cookie: SHIELD_SESSION=..., HttpOnly, Secure, SameSite=Strict
    else Com sessão válida
        Shield->>Redis: GET session (recupera JWT)
        Redis-->>Shield: {access_token, claims}
        Shield->>Kong: Injeta Authorization: Bearer <JWT> (tenant_id, roles, user_id)
        Kong->>MS: GET /api/v1/alunos + Authorization header
        MS->>PG: SET LOCAL app.current_tenant = '<tenant_id>'
        PG-->>MS: Dados filtrados por tenant (RLS)
        MS-->>User: 200 {data: [...]}
    end
```

**Fluxo pós-autenticação (sessão válida):**
1. SPA faz chamada de API → Cloudflare → Kong
2. Kong encaminha para o Shield validar cookie `SHIELD_SESSION`
3. Shield recupera JWT do Redis, injeta `Authorization: Bearer <JWT>` no header
4. Kong encaminha requisição com JWT para o microserviço de negócio
5. Microserviço executa `SET LOCAL app.current_tenant` e consulta PostgreSQL com RLS
6. Resposta retorna ao cliente — **transparente, sem intervenção do frontend**

**Regra de Ouro:** Microserviços de negócio **NUNCA** chamam Keycloak ou Shield diretamente. Eles apenas recebem o header `Authorization` injetado pelo Kong e confiam nas claims (tenant_id, roles, user_id).

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
        CF["Cloudflare Edge\nDNS + WAF + Proxy"]

        subgraph AppPlatform["App Platform / DOKS Nginx"]
            SPA["Frontend SPA\nHTML/JS/CSS estático\n(Zero secrets, Zero backend)"]
        end

        subgraph DOKS["DOKS Cluster — 3 nodes (4vCPU/8GB)"]
            IstioGW["Istio Ingress Gateway"]
            KongGW["Kong API Gateway"]
            subgraph ShieldNS["shield-system"]
                Shield["Shield BFF Pods\n2-50 (KEDA)\nValida sessão + Injeta JWT"]
                KC["Keycloak Pods\n2 (StatefulSet)"]
            end
            subgraph BusinessNS["business-system"]
                MS["Microserviços Negócio\nms-escolas-core, etc."]
            end
            Argo["Argo CD"]
            KEDA["KEDA Scaler"]
        end

        subgraph Managed["Managed Services"]
            PG["PostgreSQL HA\n2 nodes, 4GB\nRLS Multi-Tenant"]
            RedisDO["Redis Managed\n2GB, TLS\nSession Store + Cache"]
        end
    end

    CF -->|"1. https://escola-alfa.com"| SPA
    SPA -.->|"3. GET /api/v1/alunos"| CF
    CF -->|"4. Proxy API"| IstioGW
    IstioGW --> KongGW
    KongGW -->|"5. Validar sessão"| Shield
    Shield -->|"6. Cache"| RedisDO
    Shield -.->|"7a. 302 → Keycloak"| KC
    Shield -->|"7b. Injeta JWT"| KongGW
    KongGW -->|"8. Bearer JWT"| MS
    MS -->|"9. SET LOCAL tenant"| PG
    KEDA -.-> Shield
    Argo -.-> Shield
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
