# High-Level Design (HLD): PROJETO SHIELD — ms-shield-identity-auth
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Solução Técnica** | ms-shield-identity-auth |
| **Documentos Base** | 01-PROJECT-CHARTER, 10-SAD |
| **Stack** | Java 21 + Quarkus + GraalVM Native |
| **Data** | 03/08/2026 | **Versão** | 2.0 | **Metodologia** | WATERFALL |

---

## 1. System Context (C4 Level 1)

```mermaid
flowchart TB
    User["👤 Usuário\nBrowser"]

    subgraph DO["DigitalOcean"]
        subgraph Frontend["🖥️ Frontend (App Platform)"]
            SPA["SPA Estática\nPortal Escola / Reforma / SaaS\n(Zero secrets)"]
        end

        subgraph DOKS["DOKS Cluster"]
            Kong["Kong API Gateway"]
            subgraph Shield["🛡️ Plataforma Shield"]
                BFF["ms-shield-identity-auth\nQuarkus Native\nValidação de Sessão"]
            end
            subgraph Business["📦 Microserviços de Negócio"]
                MS["ms-escolas-core\nms-reforma-core\nms-saas-core"]
            end
        end

        KC["Keycloak\nMulti-Realm OIDC"]
        PG["PostgreSQL\nRLS Multi-Tenant"]
        Redis["Redis\nSession Store + Cache"]
    end

    CF["Cloudflare\nWAF + DNS + Proxy"]

    User -->|"1. https://escola-alfa.com"| CF
    CF -->|"2. Entrega estáticos"| SPA
    SPA -.->|"3. GET /api/v1/alunos"| CF
    CF -->|"4. Proxy API"| Kong
    Kong -->|"5. Validar sessão"| BFF
    BFF -->|"6. Host→Realm"| Redis
    BFF -.->|"7a. 302 login"| KC
    BFF -->|"7b. Injeta JWT"| Kong
    Kong -->|"8. Bearer JWT"| MS
    MS -->|"SET LOCAL tenant"| PG
```

---

## 2. Container Diagram (C4 Level 2)

```mermaid
flowchart TB
    User["👤 Browser"]

    subgraph Edge["🔒 Cloudflare Edge"]
        CF["Cloudflare\nWAF + DNS + Proxy API\nHeader: X-Tenant-Host"]
    end

    subgraph AppPlatform["🖥️ DigitalOcean App Platform / DOKS Nginx"]
        SPA["Frontend SPA\nHTML/JS/CSS estático"]
    end

    subgraph DOKS["DOKS Cluster"]
        subgraph KongNS["kong-system"]
            Kong["Kong API Gateway\nRouter: /api/* → validar sessão\nRouter: /auth/* → processar OIDC"]
        end

        subgraph ShieldNS["shield-system"]
            Web["Shield Web Layer\nRESTEasy Reactive + Vert.x"]
            SessionFilter["SessionFilter\nValida cookie → injeta JWT\nou 302 → Keycloak"]
        end

        subgraph BusinessNS["business-system"]
            MS["Microserviços de Negócio\nms-escolas-core, etc."]
        end
    end

    subgraph Dependencies["🔗 Serviços Externos (DOKS)"]
        KC["Keycloak\nMulti-Realm + Temas"]
        Redis["Redis\nSession Store\nHost→Realm Cache"]
        PG["PostgreSQL\nRLS Multi-Tenant"]
    end

    User -->|"1. https://escola-alfa.com"| CF
    CF -->|"2. Arquivos estáticos"| SPA
    SPA -.->|"3. GET /api/v1/alunos"| CF
    CF -->|"4. Proxy API"| Kong
    Kong -->|"5. Validar SHIELD_SESSION"| SessionFilter
    SessionFilter -->|"6. Host→Realm"| Redis
    SessionFilter -.->|"7a. 302 → Keycloak"| KC
    SessionFilter -->|"7b. Injeta Authorization: Bearer JWT"| Kong
    Kong -->|"8. JWT (tenant_id, roles)"| MS
    MS -->|"SET LOCAL tenant\nRLS"| PG
```

---

## 3. Technology Stack

```mermaid
flowchart LR
    subgraph Runtime["☕ Runtime"]
        Java21["Java 21"]
        Quarkus["Quarkus 3.x LTS"]
        GraalVM["GraalVM Native\nMandrel 24.1"]
    end

    subgraph Web["🌐 Web"]
        RESTEasy["RESTEasy Reactive"]
        Vertx["Vert.x"]
    end

    subgraph Clients["🔌 Reactive Clients"]
        OIDC["Quarkus OIDC\nExtension"]
        RedisClient["Quarkus Redis\nReactive"]
        PGClient["Quarkus Reactive\nPostgreSQL"]
    end

    subgraph Obs["📊 Observability"]
        OTel["OpenTelemetry\nJava Agent"]
        Micrometer["Micrometer\nPrometheus"]
    end

    subgraph Container["📦 Container"]
        UBIMicro["UBI-Micro 9\n~40MB + binary"]
        DOKS["DOKS\nKubernetes 1.30"]
    end

    Java21 --> Quarkus --> GraalVM
    Quarkus --> RESTEasy
    Quarkus --> OIDC
    Quarkus --> RedisClient
    Quarkus --> PGClient
    Quarkus --> Micrometer
    GraalVM --> UBIMicro --> DOKS
```

---

## 4. Integration Topology

```mermaid
flowchart LR
    User["👤 Browser"] -->|"1. HTTPS"| CF["Cloudflare"]
    CF -->|"2. Estáticos"| SPA["SPA Frontend\n(App Platform)"]
    SPA -.->|"3. GET /api/*"| CF
    CF -->|"4. Proxy API"| Kong["Kong Gateway"]
    Kong -->|"5. Validar sessão"| Shield["Shield BFF\n:8080"]
    Shield -->|"6. Host→Realm"| Redis["Redis\nSession Store"]
    Shield -.->|"7a. 302 OIDC"| KC["Keycloak\n:8443"]
    Shield -->|"7b. Injeta JWT\nAuthorization header"| Kong
    Kong -->|"8. Bearer JWT\n(tenant_id, roles)"| MS["Microserviços\nNegócio"]
    MS -->|"SET LOCAL\napp.current_tenant"| PG["PostgreSQL\nRLS"]
    Prometheus["Prometheus"] -.->|"/metrics"| Shield

    style CF fill:#f96,stroke:#333,color:#fff
    style Kong fill:#69f,stroke:#333,color:#fff
    style Shield fill:#6f9,stroke:#333
    style KC fill:#96f,stroke:#333,color:#fff
    style MS fill:#fc6,stroke:#333
```

---

## 5. Deployment Topology

```mermaid
flowchart TB
    subgraph DO["DigitalOcean NYC3"]
        subgraph AppPlatform["App Platform / DOKS Nginx"]
            SPA["Frontend SPA\nHTML/JS/CSS\n(Portal, Reforma, SaaS...)"]
        end

        subgraph DOKS["DOKS Cluster — 3 nodes (4vCPU/8GB)"]
            subgraph istio-system["istio-system"]
                IstioGW["Istio Ingress Gateway"]
            end
            subgraph kong-system["kong-system"]
                KongGW["Kong API Gateway"]
            end
            subgraph shield-system["shield-system"]
                Shield1["Shield Pod 1\nSessão + JWT Injector"]
                Shield2["Shield Pod 2\nSessão + JWT Injector"]
                ShieldN["Shield Pod N\nKEDA: 2→50\n~40MB RAM"]
            end
            subgraph business-system["business-system"]
                MS1["ms-escolas-core\nPod 1+"]
                MS2["ms-reforma-core\nPod 1+"]
            end
            subgraph iam-system["iam-system"]
                KC1["Keycloak Pod 1\nStatefulSet"]
                KC2["Keycloak Pod 2\nStatefulSet"]
            end
            subgraph monitoring["monitoring"]
                Prom["Prometheus"]
                Graf["Grafana"]
            end
        end

        subgraph Managed["Managed Services"]
            PGHA["PostgreSQL HA\n2 nodes, 4GB\nRLS Multi-Tenant"]
            RedisDO["Redis Managed\n2GB, TLS\nSession Store + Cache"]
        end
    end

    SPA -.->|"GET /api/*"| IstioGW
    IstioGW --> KongGW
    KongGW -->|"Validar sessão"| Shield1 & Shield2 & ShieldN
    Shield1 -->|"Injeta JWT"| KongGW
    KongGW -->|"Bearer JWT"| MS1 & MS2
    MS1 -->|"SET LOCAL tenant"| PGHA
    Shield1 --> RedisDO
    Shield1 -.-> Prom --> Graf
```

---

## 6. Data Flow — Interceptação de API pelo Shield (Kong Plugin)

O Shield atua como serviço de validação de sessão acoplado ao Kong. O frontend SPA **não chama o Shield diretamente** — ele faz chamadas de API normalmente, e o Kong+Shield interceptam a autenticação.

```mermaid
sequenceDiagram
    actor User as 👤 Usuário
    participant CF as Cloudflare
    participant SPA as 🖥️ SPA Frontend<br/>(App Platform)
    participant Kong as Kong API Gateway
    participant Shield as Shield BFF
    participant Redis as Redis
    participant KC as Keycloak
    participant MS as Microserviço<br/>Negócio
    participant PG as PostgreSQL

    User->>CF: https://escola-alfa.com
    CF->>SPA: Proxy → entrega HTML/JS/CSS
    SPA-->>User: SPA carregada

    User->>CF: GET /api/v1/alunos (SPA)
    CF->>Kong: Proxy API
    Kong->>Shield: Validar sessão (cookie SHIELD_SESSION)

    alt Sem sessao (primeira visita)
        Shield->>Redis: GET host:escola-alfa.com
        Redis-->>Shield: realm-escola-alfa
        Shield-->>Kong: 302 → Keycloak /realms/realm-escola-alfa/auth
        Kong-->>User: 302 → Login form (tema da escola)

        User->>KC: Credenciais
        KC-->>User: 302 → /auth/callback?code=xyz
        User->>Shield: /auth/callback?code=xyz
        Shield->>KC: POST /token (troca code por tokens)
        KC-->>Shield: {access_token, refresh_token, id_token}
        Shield->>Redis: Armazena JWT (session_id → token)
        Shield-->>User: 302 → escola-alfa.com
        Note over Shield,User: Set-Cookie: SHIELD_SESSION, HttpOnly, Secure, SameSite=Strict
    else Com sessao valida
        Shield->>Redis: GET session (recupera JWT)
        Redis-->>Shield: {access_token, claims}
        Shield->>Kong: Injeta Authorization: Bearer <JWT>
        Note over Shield,Kong: Claims: tenant_id, roles, user_id
        Kong->>MS: GET /api/v1/alunos + Authorization
        MS->>PG: SET LOCAL app.current_tenant
        PG-->>MS: Dados filtrados (RLS)
        MS-->>User: 200 {data: [...]}
    end
```

**Fluxo pós-login (chamadas subsequentes):**
1. SPA faz `GET /api/v1/alunos` → Cloudflare → Kong
2. Kong encaminha cookie para o Shield validar
3. Shield recupera JWT do Redis → injeta `Authorization: Bearer <JWT>` no header
4. Microserviço de negócio recebe o JWT com claims (tenant_id, roles, user_id)
5. Microserviço executa `SET LOCAL app.current_tenant` e consulta PostgreSQL com RLS
6. **O frontend nunca vê o JWT.** Toda a orquestração de segurança é transparente para a SPA.

**Por que o Shield fica separado do Frontend:**
- **Zero-Trust:** A SPA no App Platform é 100% estática, sem acesso a segredos do Keycloak
- **Escala independente:** KEDA escala apenas o Shield (Quarkus Native, ~40MB, cold start <100ms)
- **Isolamento:** Frontend cuida de UI/UX; Shield cuida de OIDC, cookies, injeção de JWT

---

## 7. NFR Allocation

```mermaid
flowchart TB
    subgraph NFRs["📋 Non-Functional Requirements (SRS)"]
        NFR01["NFR-01\nLatência p95 <15ms"]
        NFR02["NFR-02\nCache <5ms"]
        NFR03["NFR-03\nCold Start <100ms"]
        NFR04["NFR-04..06\nCookies HttpOnly/Secure/SameSite"]
        NFR07["NFR-07\nOWASP Top 10"]
        NFR08["NFR-08\nCross-Tenant Block"]
        NFR09["NFR-09\n99.9% Uptime"]
        NFR10["NFR-10\nEscala 2→50 pods"]
        NFR16["NFR-16\n<50MB RAM"]
    end

    subgraph Components["⚙️ Componentes"]
        GraalVM["GraalVM Native\nAOT Compilation"]
        RedisCache["Redis Cache\nVPC <1ms latency"]
        CookieMgr["CookieManager\nHttpOnly/Secure/SameSite"]
        KongGW["Kong + Istio\nRate Limit + mTLS"]
        RLS["PostgreSQL RLS\nSET LOCAL tenant"]
        DOKS["DOKS 3 nodes\nHA + Auto-healing"]
        KEDA["KEDA ScaledObject\nPrometheus trigger"]
    end

    NFR01 --> GraalVM
    NFR02 --> RedisCache
    NFR03 --> GraalVM
    NFR04 --> CookieMgr
    NFR07 --> KongGW
    NFR08 --> RLS
    NFR09 --> DOKS
    NFR10 --> KEDA
    NFR16 --> GraalVM
```

---

**[STATUS: SUCESSO]** — HLD v2.0 com diagramas Mermaid (C4 L1/L2, flowchart, sequenceDiagram, deployment topology, NFR allocation).
