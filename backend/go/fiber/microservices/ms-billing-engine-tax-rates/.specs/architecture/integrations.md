# Integrações, Dependências e Bibliotecas Locais

## Protocolos de Integração Externa

### Persistência
- **PostgreSQL:** Via driver `pgx` v5, schema `billing_tax_rates`
- **Conexão:** Configurada via `DATABASE_URL` (obrigatória, Fatal sem)
- **Uso:** Consultas de regras de ICMS, IPI, PIS/COFINS, CBS/IBS/IS (Reforma), exceções por NCM, equivalência CSOSN, Simples Nacional
- **Fonte:** `cmd/api/main.go:35-42`

### Cache
- **Redis:** Via `go-redis` v9
- **Conexão:** Configurada via `REDIS_ADDR` (obrigatória, Fatal sem)
- **Uso:** Cache de regras fiscais (decorator pattern via `CachedTaxRepository`)
- **Fonte:** `cmd/api/main.go:43-47`

## Variáveis de Ambiente

| Variável | Obrigatória | Default | Descrição |
|----------|-------------|---------|-----------|
| `DATABASE_URL` | Sim | — | String de conexão PostgreSQL |
| `REDIS_ADDR` | Sim | — | Endereço do Redis |
| `PORT` | Não | `:3000` | Porta de escuta HTTP (aceita com ou sem prefixo `:`) |
| `IBS_API_BASE_URL` | Não | `https://api.comitegestoribs.gov.br` | URL base da API do Comitê Gestor IBS (Gap G2 — não publicada) |
| `RATE_LIMIT_MAX` | Não | `100` | Máximo de requisições por janela |
| `RATE_LIMIT_WINDOW` | Não | `60` | Janela de rate limiting em segundos |
| `TAX_TOKEN_TTL_MINUTES` | Não | — | TTL do token fiscal em minutos |
| `METRICS_REQUIRE_AUTH` | Não | `false` | Exige autenticação para `/v1/metrics` |

**Fonte:** `cmd/api/main.go`

## Tabela de Dependências do Projeto (go.mod)

| Dependência | Versão | Tipo | Uso |
|-------------|--------|------|-----|
| github.com/gofiber/fiber/v2 | v2.52.12 | Direta | Framework HTTP |
| github.com/gofiber/contrib/otelfiber/v2 | v2.2.3 | Direta | OpenTelemetry middleware para Fiber |
| github.com/shopspring/decimal | v1.3.1 | Direta | Matemática financeira de precisão |
| github.com/google/uuid | v1.6.0 | Direta | Geração de UUIDs para IDTransaction |
| taxnexus-billing-core-lib | v0.0.0 (replace) | Direta | Models, repository, db, cache |
| go.opentelemetry.io/otel | v1.44.0 | Direta | OpenTelemetry SDK |
| go.opentelemetry.io/otel/sdk | v1.44.0 | Direta | OpenTelemetry TracerProvider |
| go.opentelemetry.io/otel/trace | v1.44.0 | Direta | OpenTelemetry Tracing API |
| github.com/jackc/pgx/v5 | v5.x | Indireta | Driver PostgreSQL |
| github.com/redis/go-redis/v9 | v9.x | Indireta | Cliente Redis |
| github.com/go-playground/validator/v10 | v10.x | Indireta | Validação de structs |
| github.com/valyala/fasthttp | v1.x | Indireta | HTTP server (Fiber dependency) |

### Biblioteca Local (replace)

```go
replace taxnexus-billing-core-lib => ../../../libs/go-native/taxnexus-billing-core-lib
```

A biblioteca `taxnexus-billing-core-lib` fornece:
- `models` — `DocumentoFiscalEntrada`, `DocumentoFiscalSaida`, `ItemDocumentoFiscalEntrada`, `ItemDocumentoFiscalSaida`, `TributosItemDocumentoFiscalSaida`
- `db` — `ConnectPostgres(dsn)` para pool de conexões
- `cache` — `ConnectRedis(addr)` para cliente Redis
- `repository` — `NewPostgresTaxRepository(pool)`, `NewCachedTaxRepository(repo, rdb)`, interface `TaxRepository` (8 métodos incluindo `GetIvaDualRule`)

> **Nota de Ambiente:** O uso de `replace` exige a presença da biblioteca localmente no caminho relativo `../../../libs/go-native/taxnexus-billing-core-lib`.

## Comunicação com Outros Microserviços

### API Externa (planejada)

- **Comitê Gestor IBS:** `GET /api/v1/rates?ibge_code={code}` — consulta alíquotas IBS em tempo real (F-007)
- **Status:** ⚠️ API não publicada (Gap G2 — LC 214/2025 pendente). Fallback ativo: `GetIvaDualRule()` (banco de dados)
- **Circuit Breaker:** 3 falhas em 60s → OPEN, HALF_OPEN após 5min, cache expirado como fallback
- **Cache:** Redis com TTL 24h, chave `ibs:rate:{ibge_code}:{date}`
- **Fonte:** `internal/ibsclient/client.go:60-91`, `internal/circuitbreaker/circuit_breaker.go:35-48`

### Integração Interna

O endpoint `POST /v1/calculate` é o ponto de integração principal. A autenticação é delegada ao API Gateway (Kong/Keycloak) — o middleware de auth apenas decodifica o JWT e injeta metadados do usuário nos headers.

**Novos endpoints (PR #6):** `POST /v1/simulate`, `POST /v1/token/generate`, `POST /v1/credit/calculate`, `POST /v1/supplier/validate`, `GET /v1/supplier/:cnpj`, `GET /v1/admin/tax-rates/iva-dual`.

**Rate Limiting:** Middleware com headers `X-RateLimit-Limit`, `X-RateLimit-Remaining` e resposta `429 Too Many Requests` com `Retry-After`. Configurável via `RATE_LIMIT_MAX` e `RATE_LIMIT_WINDOW`.

**API Versioning:** Prefixo `/v1/` em todas as rotas. Rotas legacy (`/calculate`, `/healthz`, `/metrics`) com deprecation warning.

**Deploy:** Docker multi-stage (`Dockerfile`), `docker-compose.yaml` (app+PG+Redis), Kubernetes (`deploy/k8s/` — configmap, deployment, service, hpa).
