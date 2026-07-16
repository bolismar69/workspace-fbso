# Implementação — Auth JWT + Métricas Prometheus + Porta Configurável

**Data/Hora:** 2026-06-21T01:34:56-03:00

**Features:**
- Middleware de autenticação JWT (Kong/Keycloak)
- Métricas Prometheus (latência, cache, erros)
- Porta configurável via env var `PORT`

**Origem:** `.specs/product/feature-roadmap.md`

---

## 1. Porta Configurável

**Lógica:** Lê `PORT` do ambiente, default `:3000`. Aceita com ou sem prefixo `:`.

```go
port := os.Getenv("PORT")
if port == "" {
    port = ":3000"
}
if port[0] != ':' {
    port = ":" + port
}
app.Listen(port)
```

**Fonte:** `cmd/api/main.go:169-178`

---

## 2. Middleware de Autenticação JWT

### Arquitetura Kong → Microserviço

```
Client → Kong API Gateway → Keycloak (valida JWT)
                              ↓
                    Kong repassa JWT validado
                              ↓
              Microserviço decodifica Base64
              (sem verificar assinatura)
                              ↓
              Extrai: sub, name, roles
              Injeta: X-User-Id, X-User-Name, X-User-Roles
```

### Funcionamento

1. Lê `Authorization: Bearer <jwt>` do header
2. Extrai o payload (parte central entre `.`) do JWT
3. Decodifica Base64 (tenta RawURL, RawStd, Std)
4. Parseia JSON das claims
5. Injeta no `c.Locals()` e `c.Request().Header`:
   - `sub` → `X-User-Id`
   - `name` / `preferred_username` / `email` → `X-User-Name`
   - `roles` / `realm_access.roles` → `X-User-Roles` (separado por vírgula)

### Comportamento não-bloqueante

- Token ausente → `c.Next()` (continua sem auth)
- Token inválido → `slog.Warn` + `c.Next()` (continua sem auth)
- Kong/Keycloak são responsáveis por rejeitar tokens inválidos na borda

### Helpers

```go
GetUserID(c)    → string  // X-User-Id do contexto
GetUserName(c)  → string  // X-User-Name do contexto
GetUserRoles(c) → string  // X-User-Roles (csv)
```

**Fonte:** `internal/middleware/auth.go:1-145`

### Testes — 9 cenários

| # | Teste | Cenário |
|---|-------|---------|
| 1 | `TestDecodeJWTClaims_Valid` | Token com sub, name, preferred_username, roles |
| 2 | `TestDecodeJWTClaims_RealmRoles` | realm_access.roles (Keycloak) |
| 3 | `TestDecodeJWTClaims_PaddedBase64` | Base64 com padding `==` |
| 4 | `TestDecodeJWTClaims_Invalid` | 6 casos: vazio, sem dots, json inválido, base64 inválido |
| 5 | `TestAuthMiddleware_WithValidToken` | Middleware extrai claims corretamente |
| 6 | `TestAuthMiddleware_WithoutToken` | Continua sem token |
| 7 | `TestAuthMiddleware_WithInvalidToken` | Continua com token inválido |
| 8 | `TestAuthMiddleware_WithRealmRoles` | realm_access.roles → X-User-Roles |
| 9 | `TestAuthMiddleware_SetsHeaders` | Headers X-User-* definidos |

---

## 3. Métricas Prometheus

### Implementação (stdlib apenas, sem dependências externas)

**Coletor:** `MetricsCollector` thread-safe com `sync.RWMutex`.

**Métricas expostas em `/metrics`:**
- `http_requests_total{method,path,status}` — counter de requisições
- `http_request_duration_seconds_count{method,path}` — histogram count
- `http_request_duration_seconds_sum{method,path}` — histogram sum
- `cache_requests_total{result="hit"|"miss"}` — cache hit/miss
- `errors_total{type="validation"|"internal"}` — erros por tipo

**Middleware automático:**
```go
app.Use(middleware.NewMetricsMiddleware(metrics))
```
Registra automaticamente método, path, status e duração de todas as requisições.

**Cache tracking (via handler):**
```go
metrics.RecordCacheHit()   // quando Redis retorna cache
metrics.RecordCacheMiss()  // quando Redis não tem cache
```

**Fonte:** `internal/middleware/metrics.go:1-160`

---

## Pipeline Atualizado

```
Request → recover → requestid → auth → logger → metrics → Handler
```

**Endpoints:**
- `GET /metrics` → Prometheus text exposition format
- `GET /healthz` → Liveness
- `GET /health` → Readiness
- `POST /calculate` → Cálculo de tributos

---

## Verificação

```
go test ./... -count=1    # OK — 90+ testes passando
go build ./...            # OK — compilação limpa
go vet ./...              # OK — sem warnings
```
