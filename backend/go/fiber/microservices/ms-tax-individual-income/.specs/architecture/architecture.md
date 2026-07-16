# Visão Geral Arquitetural — ms-tax-individual-income

Gerado pelo agente **Architect** em 2026-06-08. Atualizado com evidência de código em 2026-06-20.

## 🧱 Regras de Arquitetura da Solução

### 1. Injeção de Dependência Manual (Sem Framework DI)

O sistema utiliza injeção de dependência manual no `main.go`, sem frameworks de DI. A ordem de inicialização é:

1. Conexão PostgreSQL (`db.ConnectPostgres`)
2. Conexão Redis (`cache.ConnectRedis` — retorna nil se `REDIS_ADDR` vazio)
3. Repository (`repository.NewTaxRepository`)
4. Service (`services.NewCalculationService`)
5. TaxHandler (`handlers.NewTaxHandler`)
6. HealthHandler (`handlers.NewHealthHandler`)

**Fonte:** `main.go:22-51`

### 2. Middleware Pipeline (Fiber)

O pipeline de middlewares é aplicado globalmente:

```
Request → requestid.New() → logger.New() → Handler
```

- `requestid`: Gera `X-Request-ID` automático para cada requisição
- `logger`: Loga método, path, status e latência de cada requisição

**Fonte:** `main.go:54-56`

### 3. Estrutura de Camadas

```
handlers/        ← Parse HTTP, validação de entrada, resposta HTTP, health checks
    ↓
services/        ← Lógica de negócio, orquestração, integração externa
    ↓
repository/      ← Acesso a dados + cache Redis (via taxnexus-individual-core-lib)
    ↓
[PostgreSQL]     ← Persistência de regras fiscais e configurações
[Redis]          ← Cache de regras fiscais com TTL configurável (12h default)
```

### 4. Concorrência com Goroutines

O cálculo paralelo (Completo + Simplificado) utiliza o padrão:

```go
resChan := make(chan CalculationResult, 2)
go func() { resChan <- calculateCompleta(...) }()
go func() { resChan <- calculateSimplificada(...) }()
// Coleta 2 resultados do channel
```

**Fonte:** `services/calculation_service.go:105-116`

### 5. Propagação de Trace ID

O `X-Request-ID` gerado pelo middleware `requestid` é:
1. Extraído no handler: `c.GetRespHeader(fiber.HeaderXRequestID)`
2. Injetado no `context.Context`: `context.WithValue(ctx, "requestid", requestID)`
3. Propagado para serviços externos: `httpReq.Header.Set("X-Request-ID", tid)`
4. Incluído em todos os logs: `s.logger.With("trace_id", tid)`

**Fonte:** `handlers/tax_handler.go:22-25`, `services/inss_client.go:34-36`, `services/calculation_service.go:38-43`

### 6. Logging Estruturado (slog)

O sistema utiliza `log/slog` (stdlib Go 1.21+) com:
- Handler JSON (`slog.NewJSONHandler`)
- Nível Debug (`slog.LevelDebug`)
- Contexto enriquecido com `trace_id`
- Saída para `os.Stdout`
- Logs de cache (hit/miss/error) no repository e no service

**Fonte:** `services/calculation_service.go:29-34`, `repository/tax_repository.go:44-52`

### 7. Configuração por Variáveis de Ambiente

| Variável | Uso | Default | Obrigatória |
|----------|-----|---------|-------------|
| `DATABASE_URL` | String de conexão PostgreSQL | — | Sim (Fatal sem) |
| `REDIS_ADDR` | Endereço do Redis (vazio = cache desabilitado) | — | Não |
| `INSS_SERVICE_URL` | URL base do microserviço INSS | `http://localhost:3001` | Não |
| `TAX_CACHE_TTL_HOURS` | TTL do cache Redis em horas | `12` | Não |

**Fonte:** `main.go:23,34,42-44`, `.env`, `repository/tax_repository.go:39-43`

### 8. Porta de Escuta

O servidor Fiber escuta em `:3000` (hardcoded).

**Fonte:** `main.go:67`

### 9. Health Checks (K8s/Docker)

Três endpoints para probes de orquestrador, implementados em `handlers/health_handler.go`:

| Endpoint | Tipo | Timeout | Descrição |
|----------|------|---------|-----------|
| `GET /healthz` | Liveness | — | Sempre 200 (processo vivo, sem I/O) |
| `GET /health` | Readiness | 2s por componente | Verifica PostgreSQL + Redis |
| `GET /api/v1/health` | Readiness | 2s por componente | Idêntico, dentro do grupo API |

Redis é opcional — se `REDIS_ADDR` não configurado, não é verificado e não afeta o status.

**Fonte:** `handlers/health_handler.go:1-117`, `main.go:51,58-64`

### 10. Cache Redis com Degradação Graciosa

O cache de regras fiscais é implementado no `TaxRepository` (core-lib) com:
- **Chave configs:** `tax_configs:{taxCode}:{YYYY-MM-DD}`
- **Chave rules:** `tax_rules_list:{taxCode}:{YYYY-MM-DD}`
- **TTL:** Configurável via `TAX_CACHE_TTL_HOURS` (default 12h)
- **Nil-safe:** Se `rdb == nil`, opera apenas com PostgreSQL
- **Logging:** Hit/miss/error em JSON via `slog`
- **Fallback:** Erro de Redis → PostgreSQL direto (não interrompe o cálculo)

**Fonte:** `repository/tax_repository.go:61-167`

## 🗄️ Modelo de Dados

O diagrama ERD completo foi extraído para [architecture/erd.md](erd.md).

## ⚠️ Dívidas Técnicas Identificadas

1. **Ausência de Testes Automatizados:** Não foram encontrados arquivos `*_test.go`, o que é crítico para um serviço de cálculo financeiro.
2. **Tratamento de Erros Externos:** O cliente INSS não implementa retry ou circuit breaker. Falhas são tratadas com degradação graciosa (warning + skip), mas sem política de retentativa.
3. **Porta hardcoded:** `:3000` está hardcoded no `main.go:67` — deveria ser configurável via variável de ambiente.
4. **Autenticação ausente:** O endpoint não possui middleware de autenticação/autorização. Presume-se operação em rede interna.
