# Visão Geral Arquitetural — ms-billing-engine-tax-rates

Gerado pelo agente **Spec Miner** em 2026-06-20. Atualizado em 2026-06-30 (PR #6 merge — Fases 0-1-2: Admin Fiscal, Créditos, TaxToken, Simulação, Fornecedores; Deploy Docker/K8s).

## Regras de Arquitetura da Solução

### 1. Injeção de Dependência Manual (Sem Framework DI)

O sistema utiliza injeção de dependência manual no `main.go`, sem frameworks de DI. A ordem de inicialização é:

1. Conexão PostgreSQL (`db.ConnectPostgres(DATABASE_URL)`)
2. Conexão Redis (`cache.ConnectRedis(REDIS_ADDR)`)
3. Repository PostgreSQL (`repository.NewPostgresTaxRepository(pgPool)`)
4. Repository com cache (`repository.NewCachedTaxRepository(taxRepo, rdb)`)
5. Calculadoras individuais: `IPICalculator`, `ICMSCalculator`, `PISCofinsCalculator`, `ReformaCalculator`
6. Motor de cálculo (`BillingEnginePhased`) com pipeline SOP-013 de 7 fases (Sequential/Parallel)

**Fonte:** `cmd/api/main.go:35-65`

### 2. Middleware Pipeline (Fiber)

O pipeline de middlewares é aplicado globalmente:

```
Request → recover.New() → requestid.NewRequestIDMiddleware() → auth.NewAuthMiddleware() → logger.New() → metrics.NewMetricsMiddleware() → Handler
```

- `recover`: Recupera de panics e retorna resposta adequada
- `requestid`: Implementa rastreamento distribuído W3C Trace Context — gera/extrai `traceparent`, propaga Trace ID entre serviços, gera Request ID único por requisição
- `auth`: Decodifica JWT (Base64, sem validação de assinatura — Kong/Keycloak faz isso upstream). Injeta `X-User-Id`, `X-User-Name`, `X-User-Roles` nos headers e contexto. Não bloqueante: tokens ausentes/inválidos passam com valores vazios.
- `logger`: Loga método, path, status e latência de cada requisição
- `metrics`: Coleta métricas Prometheus (contagem de requisições, histograma de latência, cache hits/misses, erros por tipo)

**Fonte:** `cmd/api/main.go:72-78`, `internal/middleware/requestid.go`, `internal/middleware/auth.go`, `internal/middleware/metrics.go`

### 2.1. W3C Trace Context (Trace-ID + Request-ID)

O middleware implementa o padrão W3C Trace Context para rastreamento distribuído entre microserviços:

**Headers de entrada:**
- `traceparent`: `00-{traceID}-{parentSpanID}-01` (W3C — recebido do upstream)

**Headers de resposta:**
- `traceresponse`: `00-{traceID}-{servicoSpanID}-01` (W3C — para downstream)
- `X-Request-ID`: 32 hex chars (ID único gerado pelo serviço)

**Contexto Fiber (`c.Locals()`):**
- `trace_id`: Trace ID propagado entre serviços
- `span_id`: Span ID do serviço atual
- `request_id`: Request ID único do serviço
- `parent_span_id`: Span ID upstream

**Geração:** `crypto/rand` — 16 bytes aleatórios para Trace/Request ID (32 hex), 8 bytes para Span ID (16 hex).

**Fonte:** `internal/middleware/requestid.go:1-157`, `cmd/api/main.go:138-143` (uso nos logs)

### 3. Motor de Cálculo Multi-Fase (SOP-013 / C-001)

O motor `BillingEngineStruct` foi refatorado de trifásico para uma arquitetura genérica multi-fase com `CalculationPhase`, suportando intercalação arbitrária de fases sequenciais e paralelas. O pipeline SOP-013 completo possui 7 fases:

```
Fase 0 (Sequencial): IS — pré-filtro (NCM seletivo, BR-TAX-CONS-010)
    └── IS_VALOR injetado nos detalhes do input
        └── Fase 1 (Sequencial): IPI — compõe base do ICMS para consumidor final
            └── IPI_VALOR injetado nos detalhes do input
                └── Fase 2 (Sequencial): CBS — "por fora", não compõe base de outros
                    └── CBS_VALOR injetado nos detalhes do input
                        └── Fase 3 (Sequencial): ICMS — antes do PIS/COFINS (Tese do Século)
                            ├── ITEM_ICMS_VALOR injetado (FUST/FUNTTEL)
                            └── VALOR_EXCLUSAO_ICMS injetado (PIS/COFINS)
                                └── Fase 4 (Paralela via goroutines): IBS + ISS + PIS/COFINS
                                    ├── IBS: subnacional (estadual + municipal)
                                    ├── ISS: municipal sobre serviços (LC 116/2003)
                                    └── PIS/COFINS: com exclusão do ICMS da base
                                        └── Fase 5 (Sequencial): FUST — depende de ICMS+PIS+COFINS
                                            └── FUST = (Valor − ICMS − PIS − COFINS) × 1%
                                                └── Fase 6 (Sequencial): FUNTTEL — mesma base FUST
                                                    └── FUNTTEL = Base × 0,5%
```

**Arquitetura de fases:**
- `CalculationPhase` contém `Name`, `Mode` (Sequential/Parallel) e `Calculators []TaxCalculator`
- `BillingEnginePhased(phases ...CalculationPhase)` — construtor principal (C-001)
- `BillingEngineFull(pre, calcs, post)` — compatibilidade retroativa (constrói 3 fases internamente)
- `BillingEngineOrdered(pre, calcs ...)` — compatibilidade retroativa (2 fases)
- `BillingEngine(calcs ...)` — compatibilidade retroativa (1 fase paralela)

**Injeção inter-fase automática (`injectTributoValues`):**
Após cada fase, os tributos calculados são injetados nos detalhes do `input` com chaves padronizadas:
- Chave genérica: `<TRIBUTO>_VALOR` (ex: `IPI_VALOR`, `ICMS_VALOR`)
- Chaves específicas: `ITEM_ICMS_VALOR` (FUST/FUNTTEL), `VALOR_EXCLUSAO_ICMS` (PIS/COFINS), `ITEM_PIS_VALOR`, `ITEM_COFINS_VALOR`, `ITEM_IPI_VALOR`

**Propagação de erros:**
- Fases sequenciais: erro propaga (interrompe pipeline) — calculadoras críticas
- Fases paralelas: erro coletado e logado (pipeline continua) — calculadoras independentes

**Fonte:** `internal/calculator/engine.go:1-330`

### 4. Strategy Pattern (PIS/COFINS por CST)

O cálculo de PIS e COFINS utiliza o padrão Strategy com seleção por código CST:

```go
type PISStrategy interface {
    Calculate(item models.ItemDocumentoFiscalEntrada, info PISInfo) decimal.Decimal
}

func GetPISStrategy(cst string) PISStrategy {
    switch cst {
    case "01", "02": return &PIS01_02{}
    case "03":       return &PIS03{}
    default:         return &PIS99{}
    }
}
```

**Fonte:** `internal/legacy/pis_strategies.go:1-35`, `internal/legacy/cofins_strategies.go:1-35`

**Atualização 2026-06-20:** As alíquotas agora são resolvidas via `GetFederalTaxRule()` (banco) com fallback para constantes do pacote. A flag `ExcluiICMSBase` do banco controla a exclusão do ICMS da base.

### 5. Adapter Pattern (Legacy Integration)

O `LegacyAdapter` converte calculadoras do módulo `legacy` (com interfaces específicas de ICMS e PIS/COFINS) para a interface unificada `TaxCalculator`:

```go
func LegacyAdapter(calc interface{}, icmsProvider ...interface{}) TaxCalculator
```

- Para calculadoras ICMS: cria `icmsTaxAdapter`
- Para calculadoras PIS/COFINS: cria `legacyTaxAdapter` (opcionalmente recebe fonte ICMS para exclusão da base)
- Para tipos não suportados: `panic()`

**Fonte:** `internal/calculator/legacy_adapter.go:80-101`

### 6. Strategy Pattern — TelecomClassifier (FUST/FUNTTEL)

O cálculo de FUST e FUNTTEL utiliza um classificador de domínio (`TelecomClassifier`) que encapsula a regra BR-TAX-INF-007 (SCM/STFC vs. SVA):

```go
type TelecomClassifier struct{}

func (c *TelecomClassifier) MustCalculateFUST(item models.ItemDocumentoFiscalEntrada) bool {
    natureza, _ := c.Classify(item)
    return natureza == "SCM" || natureza == "STFC" // SVA não incide
}
```

**Fonte:** `internal/legacy/telecom.go:1-90`

### 7. Padrão de Injeção Inter-Fase (Inter-Phase Injection)

O método `injectTributoValues()` da engine extrai os tributos calculados em cada fase e os injeta nos detalhes do `input` com chaves padronizadas. Isso permite que fases e calculadoras subsequentes leiam tributos já computados sem modificar a interface `domain.TaxCalculator`. O padrão é aplicado automaticamente após cada fase (sequencial ou paralela).

**Chaves injetadas por tributo:**
| Tributo | Chave(s) | Consumidores |
|---------|----------|-------------|
| IPI | `IPI_VALOR`, `ITEM_IPI_VALOR` | ICMS (base de cálculo) |
| CBS | `CBS_VALOR` | Auditoria |
| ICMS | `ICMS_VALOR`, `ITEM_ICMS_VALOR`, `VALOR_EXCLUSAO_ICMS` | PIS/COFINS (exclusão), FUST/FUNTTEL (base líquida) |
| PIS | `PIS_VALOR`, `ITEM_PIS_VALOR` | FUST/FUNTTEL |
| COFINS | `COFINS_VALOR`, `ITEM_COFINS_VALOR` | FUST/FUNTTEL |
| ISS | `ISS_VALOR` | Auditoria |
| IBS | `IBS_VALOR` | Auditoria |
| IS | `IS_VALOR` | Auditoria |

**Destaque — ICMS com chaves duplas:** O ICMS é injetado com ambas as chaves `ITEM_ICMS_VALOR` (para FUST e FUNTTEL) e `VALOR_EXCLUSAO_ICMS` (para PIS/COFINS — exclusão da base, Tese do Século), garantindo que ambos os consumidores encontrem o valor sem modificar suas interfaces.

**Fonte:** `internal/calculator/engine.go:280-330` (injectTributoValues)

### 8. Configuração por Variáveis de Ambiente

| Variável | Uso | Default | Obrigatória |
|----------|-----|---------|-------------|
| `DATABASE_URL` | String de conexão PostgreSQL | — | Sim (Fatal sem) |
| `REDIS_ADDR` | Endereço do Redis | — | Sim (Fatal sem) |

**Fonte:** `cmd/api/main.go:35-36`

### 9. Porta de Escuta Configurável

O servidor Fiber escuta na porta definida pela env var `PORT`, com fallback para `:3000`:

```go
port := os.Getenv("PORT")
if port == "" { port = ":3000" }
if port[0] != ':' { port = ":" + port }
```

Aceita valores com ou sem prefixo `:` (ex: `3000`, `:3000`, `:8080`).

**Fonte:** `cmd/api/main.go:169-175`

### 10. Logging Estruturado (slog)

O sistema utiliza `log/slog` (stdlib Go 1.21+) com:
- Handler JSON (`slog.NewJSONHandler`)
- Nível Debug (`slog.LevelDebug`)
- Saída para `os.Stdout`

**Fonte:** `cmd/api/main.go:29-30`

### 11. Estrutura de Camadas

```
internal/calculator/      ← Motor de orquestração multi-fase (BillingEnginePhased + CalculationPhase)
    ↓
internal/domain/         ← Interface TaxCalculator (DDD — camada mais interna)
internal/legacy/         ← Calculadoras fiscais (ICMS, IPI, PIS/COFINS, ISS, FUST, FUNTTEL, ISFilter) + TelecomClassifier
internal/reforma/        ← Reforma Tributária (CBSCalculator, IBSCalculator, ReformaCalculator legado)
internal/phase/          ← Phase Resolver + TaxSelector (fases da Reforma Tributária)
internal/circuitbreaker/ ← Circuit Breaker para API IBS (F-007)
internal/ibsclient/      ← IBS Client com cache Redis (F-007)
    ↓
repository (core-lib)    ← Acesso a dados + cache Redis (via taxnexus-billing-core-lib)
    ↓
[PostgreSQL]             ← Persistência de regras fiscais (15 tabelas — ver [data-dictionary.md](data-dictionary.md))
[Redis]                  ← Cache de regras fiscais + cache IBS (TTL 24h)
```

> 📋 **Dicionário de Dados:** A descrição funcional completa de cada tabela (propósito, padrões de lookup, calculadoras associadas, regras de negócio) está em [data-dictionary.md](data-dictionary.md).

### 12. Health Check Endpoints

O servidor expõe dois endpoints de health check:

**`GET /healthz`** — Liveness probe (Kubernetes):
- Retorna `200 {"status": "ok"}` se o processo está vivo.

**`GET /health`** — Readiness probe (Kubernetes):
- Verifica PostgreSQL (`pgPool.Ping()`) e Redis (`rdb.Ping()`)
- Retorna `200 {"status": "ok", "checks": {...}}` se todas as dependências saudáveis
- Retorna `503 {"status": "degraded", "checks": {...}}` se alguma dependência falhou

**Fonte:** `cmd/api/main.go:81-116`

### 13. Tratamento de Erros nas Goroutines (Fase 2)

Erros das calculadoras paralelas são coletados via buffered channel e logados como WARN:

```go
errChan := make(chan error, len(e.calculators))
// ... goroutines enviam erros para errChan ...
wg.Wait()
close(errChan)
for err := range errChan {
    slog.Warn("Erro em calculadora paralela (Fase 2) — cálculo parcial", "error", err)
}
```

Erros na Fase 2 são não-fatais: resultados parciais das calculadoras que funcionaram são retornados.

**Fonte:** `internal/calculator/engine.go:83-114`

### 14. Autenticação JWT (Delegada ao API Gateway)

O middleware de autenticação (`auth.go`) decodifica o payload JWT sem verificar assinatura (Kong/Keycloak no edge já fez isso). Extrai:
- `sub` → `X-User-Id`
- `name`/`preferred_username`/`email` → `X-User-Name`
- `roles`/`realm_access.roles` → `X-User-Roles`

As informações são injetadas em `c.Locals()` e nos headers da requisição para downstream. O middleware é **não bloqueante**: tokens ausentes ou inválidos não rejeitam a requisição (passam com valores vazios). Tenta 3 encodings Base64 (RawURL, RawStd, Std).

**Fonte:** `internal/middleware/auth.go:1-147`, `internal/middleware/auth_test.go` (9 cenários)

### 15. Métricas Prometheus (Stdlib Only)

Coletor de métricas thread-safe exposto via `GET /metrics` em formato Prometheus text exposition. Métricas:
- `http_requests_total` — counter por method/path/status
- `http_request_duration_seconds` — histogram com buckets [0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 5]
- `cache_requests_total` — counter com label result={hit,miss}
- `errors_total` — counter por error type (validation, internal)

Implementação sem dependências externas; apenas `sync.RWMutex` + texto formatado.

**Fonte:** `internal/middleware/metrics.go:1-147`

### 16. Reforma Tributária (CBS, IBS, IS)

O pacote `internal/reforma/` contém três calculadoras para a Reforma Tributária (EC 132/2023, LC 214/2025):

**`CBSCalculator` (`cbs_calculator.go`) — CBS apenas (Fase 2 do pipeline):**
Calcula a CBS (Contribuição sobre Bens e Serviços), tributo federal que substitui PIS e COFINS. No pipeline SOP-013, executa sequencialmente na Fase 2 ("por fora" — não compõe base de outros tributos), antes do ICMS.

**`IBSCalculator` (`ibs_calculator.go`) — IBS apenas (Fase 4 do pipeline):**
Calcula o IBS (Imposto sobre Bens e Serviços), tributo subnacional (estadual + municipal). No pipeline SOP-013, executa em paralelo na Fase 4 com ISS e PIS/COFINS.

**`ReformaCalculator` (`reforma.go`) — CBS + IBS combinados (legado):**
Mantido para compatibilidade com código existente. Calcula CBS e IBS em uma única passagem. Não recomendado para novos pipelines — use `CBSCalculator` e `IBSCalculator` separadamente.

**Lógica compartilhada (`computeIvaDual` em `reforma.go`):**
A função interna `computeIvaDual()` consulta a tabela `iva_dual_rules` via `GetIvaDualRule(ncm, ufDestino, municipioIBGE)` e retorna uma struct `ivaDualResult` com todos os valores computados. Tanto `CBSCalculator` quanto `IBSCalculator` chamam esta função — o `CachedTaxRepository` garante que a segunda consulta atinja o cache Redis.

**Fluxo de cálculo por item:**
1. Extrai NCM, UF destino e município IBGE (opcional) do documento
2. Consulta `GetIvaDualRule()` — se `nil`, skip com `slog.Warn`
3. Calcula `fatorReducao = 1 − (percentualReducao / 100)`
4. Se `percentualReducao >= 100`: isenção total (sem CBS/IBS)
5. CBS = base × (aliquotaCBS × fatorReducao / 100)
6. IBS = base × ((aliquotaIBSEstadual + aliquotaIBSMunicipal) × fatorReducao / 100)
7. IS = movido para `ISFilter` (Fase 0) — não faz parte do ReformaCalculator

**Cache:** Redis com chave `tax:iva:<ncm>:<uf>:<municipio>` (TTL 24h).

**Fonte:** `internal/reforma/reforma.go:1-178`, `internal/reforma/cbs_calculator.go:1-64`, `internal/reforma/ibs_calculator.go:1-65`

### Dívidas Técnicas Identificadas

1. ~~**ID de transação placeholder:** `IDTransaction: "0"` — substituído por geração real (`engine.go:49`, `icms.go:26` via `uuid.NewString()`). Resolvido em 2026-06-21.~~
2. **TODO CRT:** `TODO: VERIFICAR SE PRECISAR SER O CRT DO EMITENTE OU DESTINATARIO` (`icms.go:362`).
3. **CST provisório da Reforma:** CSTs `01`/`04` usados para CBS/IBS/IS aguardam tabela oficial da RFB (`internal/reforma/reforma.go`).
4. **Créditos da Reforma não implementados:** Cash forward (`permite_credito_amplo`) da Reforma Tributária não implementado.

### Dívidas Resolvidas (2026-06-21)

| Item | Resolução |
|------|-----------|
| ICMS Próprio não finalizado | `getEffectiveTaxConfig()` integrado; fluxo reestruturado com 3 métodos auxiliares |
| Taxas PIS/COFINS hardcoded | Migrado para `GetFederalTaxRule()` com fallback para constantes |
| Middleware requestid ausente | Implementado W3C Trace Context — Request-ID + Trace-ID em `internal/middleware/requestid.go` |
| Goroutine errors silenciosos | Channel de erro + `slog.Warn` na Fase 2 em `internal/calculator/engine.go:82-112` |
| Health checks ausentes | `/healthz` (liveness) e `/health` (readiness) em `cmd/api/main.go:81-116` |
| Sem autenticação/autorização | JWT middleware implementado em `internal/middleware/auth.go` — pass-through, não bloqueante |
| Porta hardcoded | Agora configurável via env var `PORT` com default `:3000` (`cmd/api/main.go:169-175`) |
| Módulo Reforma Tributária ausente | Implementado em `internal/reforma/` (CBS/IBS/IS), `GetIvaDualRule` no repository, wiring no engine |
| Calculadoras ISS, FUST, FUNTTEL ausentes | Implementado em 2026-06-21: `iss.go`, `fust.go`, `funttel.go` + `telecom.go` (classificador) — 17 testes |
| Motor bifásico insuficiente para dependências | Refatorado para trifásico com `BillingEngineFull()` — Fase 3 pós-paralela para FUST/FUNTTEL |
| Motor trifásico insuficiente para pipeline SOP-013 | Refatorado para multi-fase genérico com `BillingEnginePhased()` e `CalculationPhase` (7 fases SOP-013). C-001 concluído (2026-06-22). |
| Information disclosure em health/error responses | Sanitizado: erros movidos para `slog`, cliente vê mensagens genéricas |
