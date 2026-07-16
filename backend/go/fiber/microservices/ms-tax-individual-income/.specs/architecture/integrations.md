# Integrações, Dependências e Bibliotecas Locais

## 🔌 Protocolos de Integração Externa

### API Externa Consumida: MS INSS
- **Endpoint:** `POST {INSS_SERVICE_URL}/api/v1/calculate/inss`
- **Timeout:** 5 segundos
- **Rastreabilidade:** Propaga header `X-Request-ID`
- **Resiliência:** Degradação graciosa — falha gera warning log, cálculo prossegue sem dedução previdenciária
- **Fonte:** `services/inss_client.go:14-53`, `services/calculation_service.go:150-161`

### Persistência
- **PostgreSQL:** Via driver `pgx` v5.9.1, schema `individual_tax_rates`
- **Conexão:** Configurada via `DATABASE_URL` (obrigatória, Fatal sem)
- **Fonte:** `main.go:23-30`, `go.mod:20`, `.env`

### Cache
- **Redis:** Via `go-redis` v9.18.0 com nil-safety, logging `slog` e TTL configurável
- **Conexão:** Configurada via `REDIS_ADDR` (vazio = cache desabilitado, serviço opera normalmente)
- **TTL:** Configurável via `TAX_CACHE_TTL_HOURS` (default 12h)
- **Chaves de cache:**
  - Configurações: `tax_configs:{taxCode}:{YYYY-MM-DD}`
  - Regras: `tax_rules_list:{taxCode}:{YYYY-MM-DD}`
- **Observabilidade:** Logs JSON com hit/miss/error em nível Debug/Warn
- **Fonte:** `repository/tax_repository.go:61-167`, `cache/redis.go:1-27`, `main.go:34-38`

### Health Checks
- **Endpoints:** `GET /healthz` (liveness), `GET /health` (readiness), `GET /api/v1/health`
- **Verificação:** PostgreSQL (`pgxpool.Ping`) + Redis (`redis.Ping`, se configurado)
- **Timeout:** 2 segundos por componente
- **Fonte:** `handlers/health_handler.go:1-117`, `main.go:51,58-64`

## 🌐 Variáveis de Ambiente

| Variável | Obrigatória | Default | Descrição |
|----------|-------------|---------|-----------|
| `DATABASE_URL` | Sim | — | String de conexão PostgreSQL |
| `REDIS_ADDR` | Não | — | Endereço do Redis (vazio = sem cache) |
| `INSS_SERVICE_URL` | Não | `http://localhost:3001` | URL base do microserviço INSS |
| `TAX_CACHE_TTL_HOURS` | Não | `12` | TTL do cache Redis em horas |

**Fonte:** `main.go:23,34,42-44`, `.env`, `repository/tax_repository.go:39-43`

## 📦 Tabela de Dependências do Projeto (go.mod)

| Dependência | Versão | Tipo | Uso |
|-------------|--------|------|-----|
| github.com/gofiber/fiber/v2 | v2.52.12 | Direta | Framework HTTP |
| github.com/shopspring/decimal | v1.4.0 | Direta | Matemática financeira de precisão |
| taxnexus-individual-core-lib | v0.0.0 (replace) | Direta | Models, repository, db, cache |
| github.com/jackc/pgx/v5 | v5.9.1 | Indireta | Driver PostgreSQL |
| github.com/redis/go-redis/v9 | v9.18.0 | Indireta | Cliente Redis |
| github.com/google/uuid | v1.6.0 | Indireta | Geração de UUIDs |
| github.com/valyala/fasthttp | v1.51.0 | Indireta | HTTP server (Fiber dependency) |

### Biblioteca Local (replace)

```go
replace taxnexus-individual-core-lib => ../../../libs/go-native/taxnexus-individual-core-lib
```

A biblioteca `taxnexus-individual-core-lib` fornece:
- `models` — `UniversalTaxRequest`, `TaxResponse`, `DeductionDetail`, `DocumentoFiscalRequest`
- `db` — `ConnectPostgres(dsn)` para pool de conexões
- `cache` — `ConnectRedis(addr)` com nil-safety (retorna nil se addr vazio)
- `repository` — `NewTaxRepository(pool, rdb)`, `GetTableConfigs()`, `GetApplicableRule()`, `GetTaxRulesForPeriod()` (todos com cache Redis + logging)

> **Nota de Ambiente:** O uso de `replace` exige a presença da biblioteca localmente no caminho relativo `../../../libs/go-native/taxnexus-individual-core-lib`. Em CI/CD, este caminho deve existir ou ser substituído por um módulo versionado.

**Fonte:** `go.mod:5,10`, `main.go:12-14`, `services/calculation_service.go:11-12`
