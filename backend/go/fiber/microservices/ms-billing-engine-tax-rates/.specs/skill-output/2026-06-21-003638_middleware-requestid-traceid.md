# Implementação — Middleware Request-ID + Trace-ID (W3C Trace Context)

**Data/Hora:** 2026-06-21T00:36:38-03:00

**Feature:** Middleware de rastreabilidade distribuída — Request-ID + Trace-ID via W3C traceparent

**Origem:** `.specs/product/feature-roadmap.md` — "Middleware requestid"

---

## Escopo

Middleware Fiber que implementa rastreamento distribuído seguindo o padrão W3C Trace Context:

1. **Request-ID**: ID único gerado pelo microserviço para cada requisição (32 hex chars)
2. **Trace-ID**: Segue o padrão W3C via header `traceparent` (32 hex chars)
3. **Span-ID**: Identifica o span deste serviço na trace distribuída (16 hex chars)
4. **Parent Span-ID**: Span ID upstream (extraído do traceparent ou gerado)

---

## Arquivos Criados

### 1. `internal/middleware/requestid.go` — Middleware completo

**Headers:**
| Header | Direção | Formato | Descrição |
|--------|---------|---------|-----------|
| `traceparent` | Request (in) | `00-{traceID}-{spanID}-01` | W3C Trace Context — recebido do upstream |
| `traceresponse` | Response (out) | `00-{traceID}-{newSpanID}-01` | W3C Trace Context — resposta para downstream |
| `X-Request-ID` | Response (out) | 32 hex chars | ID único do serviço para esta requisição |

**Contexto Fiber (`c.Locals()`):**

| Key | Tipo | Descrição |
|-----|------|-----------|
| `trace_id` (ContextKeyTraceID) | `string` | Trace ID (32 hex) — propagado entre serviços |
| `span_id` (ContextKeySpanID) | `string` | Span ID (16 hex) — representa este serviço |
| `request_id` (ContextKeyRequestID) | `string` | Request ID (32 hex) — único no serviço |
| `parent_span_id` (ContextKeyParentSpanID) | `string` | Span ID upstream (16 hex) |

**Funções exportadas:**

| Função | Retorno | Descrição |
|--------|---------|-----------|
| `ParseTraceParent(header)` | `(W3CTraceContext, bool)` | Parse do header traceparent W3C |
| `FormatTraceParent(traceID, spanID)` | `string` | Formata traceparent no padrão W3C |
| `GenerateTraceID()` | `string` | 32 hex chars (16 bytes aleatórios) |
| `GenerateSpanID()` | `string` | 16 hex chars (8 bytes aleatórios) |
| `GenerateRequestID()` | `string` | 32 hex chars (16 bytes aleatórios) |

**Algoritmo do middleware:**
1. Lê `traceparent` do header de entrada
2. Se válido → reutiliza Trace ID e Parent Span ID do upstream
3. Se inválido/ausente → gera novo Trace ID + Parent Span ID
4. Gera novo Span ID para este serviço (sempre)
5. Gera Request ID único (sempre)
6. Armazena no `c.Locals()` para uso por handlers/logging
7. Adiciona `X-Request-ID` e `traceresponse` nos headers de resposta

### 2. `internal/middleware/requestid_test.go` — 12 cenários

| # | Teste | Cenário |
|---|-------|---------|
| 1 | `TestParseTraceParent_Valid` | Parse de header W3C válido |
| 2 | `TestParseTraceParent_InvalidCases` | 9 casos inválidos (versão, tamanho, hex, flags) |
| 3 | `TestFormatTraceParent` | Formatação correta do header |
| 4 | `TestGenerateHexID` | Unicidade e tamanho (32 hex) |
| 5 | `TestGenerateTraceID` | Tamanho correto (32 hex) |
| 6 | `TestGenerateSpanID` | Tamanho correto (16 hex) |
| 7 | `TestGenerateRequestID` | Tamanho correto (32 hex) |
| 8 | `TestRoundTrip_ParseThenFormat` | Parse + Format idempotente |
| 9 | `TestMiddleware_WithoutTraceParent` | Sem traceparent → gera tudo novo |
| 10 | `TestMiddleware_WithValidTraceParent` | Com traceparent → reutiliza TraceID |
| 11 | `TestMiddleware_WithInvalidTraceParent_GeneratesNew` | traceparent inválido → fallback |
| 12 | `TestMiddleware_RequestID_Uniqueness` | 50 requisições → 50 IDs únicos |

### 3. `cmd/api/main.go` — Integração

**Middleware pipeline atualizado:**
```
Request → recover → requestid → logger → Handler
```

**Logging enriquecido:** Handlers incluem `trace_id` e `request_id` nos logs estruturados via `c.Locals()`.

---

## Exemplo de Funcionamento

### Cenário 1: Requisição sem traceparent (origem)

```
Request:  POST /calculate  (sem header traceparent)
Response: X-Request-ID: a1b2c3d4e5f6... (novo)
          traceresponse: 00-4bf92f3577b34da6a3ce929d0e0e4736-ff1234567890abcd-01
```

### Cenário 2: Requisição com traceparent upstream

```
Request:  POST /calculate
          traceparent: 00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01
Response: X-Request-ID: a1b2c3d4e5f6... (novo, do serviço)
          traceresponse: 00-4bf92f3577b34da6a3ce929d0e0e4736-<novoSpanID>-01
```

---

## Verificação

```
go test ./... -count=1    # OK — 76 testes passando
go build ./...            # OK — compilação limpa
go vet ./...              # OK — sem warnings
```

---

## Dependências

- **Stdlib apenas:** `crypto/rand`, `encoding/hex`, `fmt`, `strings` (sem dependências externas)
- **Framework:** Fiber v2 (`c.Locals()`, `c.Get()`, `c.Set()`, `c.Next()`)
