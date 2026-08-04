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

    subgraph Shield["🛡️ Plataforma Shield"]
        BFF["ms-shield-identity-auth\nQuarkus Native : 8080"]
    end

    KC["Keycloak\nMulti-Realm OIDC"]
    PG["PostgreSQL\nRLS Multi-Tenant"]
    Redis["Redis\nHost→Realm Cache"]
    Produtos["Produtos FBSO\nGestão Escolar, Portal, Comunidades"]

    User -->|"HTTPS + Cookie"| BFF
    BFF -->|"OIDC/PKCE"| KC
    BFF -->|"Reactive SQL"| PG
    BFF -->|"Cache GET/SET"| Redis
    Produtos -->|"Consome /auth/*"| BFF
```

---

## 2. Container Diagram (C4 Level 2)

```mermaid
flowchart TB
    subgraph Edge["🔒 Edge"]
        CF["Cloudflare\nWAF + Header Injection"]
    end

    subgraph Gateway["🚪 Gateway"]
        Kong["Kong API Gateway\nRoute: /auth/* → Shield\nRoute: /api/* → Produtos\nJWT Validation + Rate Limit"]
    end

    subgraph Mesh["🔐 Mesh"]
        Istio["Istio Service Mesh\nmTLS + Traffic Control"]
    end

    subgraph BFF["⚙️ ms-shield-identity-auth"]
        Web["Web Layer\nRESTEasy Reactive + Vert.x"]
        Svc["Service Layer\nCDI Beans"]
        Int["Integration Layer\nReactive Clients"]
    end

    subgraph Dependencies["🔗 Dependencies"]
        Keycloak["Keycloak\nMulti-Realm"]
        RedisCache["Redis\nCache"]
        PostgresDB["PostgreSQL\nRLS"]
        PrometheusMetrics["Prometheus\n/metrics scrape"]
    end

    CF -->|"TLS 1.3\nX-Tenant-Host"| Kong
    Kong -->|"HTTP/2 mTLS"| Istio
    Istio -->|"/auth/*"| Web
    Web --> Svc
    Svc --> Int
    Int -->|"OIDC"| Keycloak
    Int -->|"Redis Protocol"| RedisCache
    Int -->|"Reactive SQL"| PostgresDB
    Web -->|"/metrics"| PrometheusMetrics
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
    CF["Cloudflare"] -->|"HTTPS\nX-Tenant-Host"| Kong["Kong API Gateway"]
    Kong -->|"HTTP/2 mTLS\nJWT Bearer"| BFF["Shield BFF\n:8080"]
    BFF -->|"OIDC Client Secret\nTLS 1.3"| KC["Keycloak\n:8443"]
    BFF -->|"Redis TLS\nPassword"| Redis["Redis\n:6379"]
    BFF -->|"SQL TLS\nSET LOCAL tenant"| PG["PostgreSQL\n:5432"]
    Prometheus["Prometheus"] -.->|"Pull /metrics"| BFF

    style CF fill:#f96,stroke:#333,color:#fff
    style Kong fill:#69f,stroke:#333,color:#fff
    style BFF fill:#6f9,stroke:#333
    style KC fill:#96f,stroke:#333,color:#fff
```

---

## 5. Deployment Topology

```mermaid
flowchart TB
    subgraph DO["DigitalOcean NYC3"]
        subgraph DOKS["DOKS — 3 nodes"]
            subgraph istio-system["istio-system"]
                IstioGW["Istio Ingress Gateway"]
            end
            subgraph shield-system["shield-system"]
                direction TB
                BFF1["Shield BFF Pod 1\nQuarkus Native\n~40MB RAM"]
                BFF2["Shield BFF Pod 2\nQuarkus Native\n~40MB RAM"]
                BFFN["Shield BFF Pod N\nKEDA: 2→50\n~40MB RAM"]
                KCPod1["Keycloak Pod 1\nStatefulSet"]
                KCPod2["Keycloak Pod 2\nStatefulSet"]
            end
            subgraph monitoring["monitoring"]
                Prom["Prometheus"]
                Graf["Grafana"]
                Loki["Grafana Loki"]
                Jaeger["Jaeger"]
            end
        end
        subgraph Managed["Managed Services"]
            PGHA["PostgreSQL HA\n2 nodes, 4GB"]
            RedisDO["Redis Managed\n2GB, TLS"]
        end
    end

    IstioGW --> BFF1 & BFF2 & BFFN
    BFF1 --> KCPod1 & KCPod2
    BFF1 --> PGHA
    BFF1 --> RedisDO
    BFF1 -.-> Prom
    Prom --> Graf
    BFF1 -.-> Loki
    BFF1 -.-> Jaeger
```

---

## 6. Data Flow — Login Sequence

```mermaid
sequenceDiagram
    actor User as 👤 Browser
    participant CF as Cloudflare
    participant Kong as Kong Gateway
    participant BFF as Shield BFF
    participant Redis as Redis
    participant KC as Keycloak

    User->>CF: GET /app
    CF->>Kong: HTTPS + X-Tenant-Host: escola-alfa.com
    Kong->>BFF: GET /auth/login?redirect_uri=/app

    BFF->>Redis: GET host:escola-alfa.com
    Redis-->>BFF: realm-escola-alfa

    BFF->>KC: Redirect /realms/realm-escola-alfa/protocol/openid-connect/auth?code_challenge=...
    KC-->>User: 302 → Login Form

    User->>KC: POST credentials
    KC-->>User: 302 → /auth/callback?code=xyz&state=...

    User->>BFF: GET /auth/callback?code=xyz&state=...
    BFF->>KC: POST /token (code + code_verifier)
    KC-->>BFF: {access_token, refresh_token, id_token}

    BFF-->>User: 302 /app + Set-Cookie
    Note over BFF,User: Cookie flags: HttpOnly, Secure, SameSite=Strict
```

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
