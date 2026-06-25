# Implementação — Health Check Endpoints + Coleta de Erros em Goroutines

**Data/Hora:** 2026-06-21T00:51:50-03:00

**Features:**
- Health check endpoints (`/healthz`, `/health`)
- Coleta de erros em goroutines (Fase 2 da engine)

**Origem:** `.specs/product/feature-roadmap.md` — "Health check endpoints" + "Coleta de erros em goroutines"

---

## 1. Health Check Endpoints

### `/healthz` — Liveness Probe

**Método:** GET
**Resposta 200:** `{"status": "ok"}`
**Propósito:** Kubernetes liveness probe — indica que o processo está vivo e responde.

### `/health` — Readiness Probe

**Método:** GET
**Resposta 200:** `{"status": "ok", "checks": {"postgres": "healthy", "redis": "healthy"}}`
**Resposta 503:** `{"status": "degraded", "checks": {"postgres": "healthy", "redis": "unhealthy: connection refused"}}`
**Propósito:** Kubernetes readiness probe — verifica dependências (PostgreSQL ping, Redis ping). Só marca como pronto quando todas as dependências estão saudáveis.

**Fonte:** `cmd/api/main.go:77-110`

---

## 2. Coleta de Erros em Goroutines (Fase 2)

### Problema

Anteriormente, erros nas goroutines da Fase 2 eram silenciosamente ignorados (`engine.go:91-93`):

```go
// ANTES (bug):
go func(c TaxCalculator) {
    defer wg.Done()
    _, err := c.Calculate(ctx, input)
    if err != nil {
        return  // erro perdido!
    }
    // ...
}(calc)
```

### Solução

Channel de erro com buffer + log estruturado pós `wg.Wait()`:

```go
// DEPOIS (corrigido):
errChan := make(chan error, len(e.calculators))

go func(c TaxCalculator) {
    defer wg.Done()
    resItens, err := c.Calculate(ctx, input)
    if err != nil {
        errChan <- err  // erro coletado!
        return
    }
    // ...
}(calc)

wg.Wait()
close(errChan)

for err := range errChan {
    slog.Warn("Erro em calculadora paralela (Fase 2) — cálculo parcial",
        "error", err,
    )
}
```

**Comportamento:**
- Erros na Fase 2 são **não-fatais** para o cálculo global
- Resultados parciais são retornados (calculadoras que funcionaram contribuem)
- Cada erro é logado via `slog.Warn` com nível WARN para visibilidade em produção
- Buffer do channel dimensionado pelo número de calculadoras (`len(e.calculators)`) evita deadlock

**Fonte:** `internal/calculator/engine.go:82-112`

---

## Arquivos Modificados

| Arquivo | Mudança |
|---------|---------|
| `internal/calculator/engine.go` | Import `log/slog`, errChan com buffer, log pós WaitGroup |
| `internal/calculator/engine_test.go` | `TestEngine_CalcParaleloErro_SilenciosamenteIgnorado` → `TestEngine_CalcParaleloErro_ColetaELoga` + `TestEngine_CalcParaleloMultiplosErros_ColetaTodos` |
| `cmd/api/main.go` | Rotas `/healthz` e `/health` com checks de PostgreSQL e Redis |

---

## Testes

| # | Teste | Cenário |
|---|-------|---------|
| 1 | `TestEngine_CalcParaleloErro_ColetaELoga` | 1 calculadora falha, 1 sucesso → resultado parcial + log WARN |
| 2 | `TestEngine_CalcParaleloMultiplosErros_ColetaTodos` | 2 calculadoras falham, 1 sucesso → apenas sucesso contribui |

---

## Verificação

```
go test ./... -count=1    # OK — 78 testes passando
go build ./...            # OK — compilação limpa
go vet ./...              # OK — sem warnings
```
