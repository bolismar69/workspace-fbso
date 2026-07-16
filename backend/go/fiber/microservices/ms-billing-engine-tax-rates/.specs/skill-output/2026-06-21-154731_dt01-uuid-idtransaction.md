# DT-01: Substituicao do IDTransaction placeholder por UUID real

**Data:** 2026-06-21 15:47  
**Feature:** DT-01 — Geracao de UUID para `IDTransaction`  
**Prioridade:** Baixa  
**Status:** Completo

## Contexto

O campo `IDTransaction` no modelo `DocumentoFiscalSaida` estava com valor placeholder `"0"` em dois pontos do codigo:

| Arquivo | Linha | Descricao |
|---------|-------|-----------|
| `internal/calculator/engine.go:49` | 49 | Response principal da engine (orquestrador central) |
| `internal/legacy/icms.go:26` | 26 | Response do calculador legado ICMS (descartado pelo adapter) |

## Implementacao

### 1. `internal/calculator/engine.go`

**Ponto central de impacto** — este e o unico local onde `IDTransaction` chega ao cliente, pois o `BillingEngineStruct.Process()` cria o `DocumentoFiscalSaida` final e mescla os resultados dos calculadores individuais.

Alteracoes:
- Adicionado import `github.com/google/uuid`
- Substituido `IDTransaction: "0"` por `IDTransaction: uuid.NewString()`

```go
response := models.DocumentoFiscalSaida{
    IDTransaction: uuid.NewString(),
    Itens:         make([]models.ItemDocumentoFiscalSaida, len(input.Itens)),
}
```

### 2. `internal/legacy/icms.go`

**Ponto secundario** — o `IDTransaction` do calculador legado e descartado pelo `icmsTaxAdapter` (que extrai apenas `response.Itens`). A correcao foi feita para consistencia e saneamento do codigo.

Alteracoes:
- Adicionado import `github.com/google/uuid`
- Substituido `IDTransaction: "0"` por `IDTransaction: uuid.NewString()`

```go
response := models.DocumentoFiscalSaida{
    IDTransaction: uuid.NewString(),
}
```

### 3. Dependencias

`github.com/google/uuid v1.6.0` promovido de dependencia indireta para direta no `go.mod` via `go mod tidy`.

## Validacao

- Build: `go build ./...` — OK
- Testes: `go test ./...` — OK (4 pacotes: calculator, legacy, middleware, reforma)

## Impacto no Response JSON

Antes:
```json
{
  "IDTransaction": "0",
  ...
}
```

Depois:
```json
{
  "IDTransaction": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  ...
}
```

Cada chamada ao `/calculate` gera um UUID v4 unico para rastreamento da transacao, visivel nos logs de auditoria (`slog.Info("Calculo finalizado", "IDTransaction", response.IDTransaction, ...)`).
