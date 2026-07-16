# API Guidelines e Padrões Globais — ms-billing-engine-tax-rates

Atualizado em 2026-06-21 15:47 (DT-01: UUID IDTransaction).

## Tratamento de Erros e Rastreabilidade

- **Middleware recover:** Todos os panics são capturados e convertidos em respostas HTTP adequadas.
- **Middleware logger:** Toda requisição é logada com método, path, status e latência.
- **Middleware requestid (W3C Trace Context):** Gera/extrai cabeçalhos `traceparent`, propaga Trace ID entre serviços, gera Request ID único por requisição. Headers de resposta: `traceresponse`, `X-Request-ID`.
- **Middleware auth:** Decodifica JWT (Base64, sem validação de assinatura — delegada ao Kong/Keycloak). Não bloqueante: tokens ausentes/inválidos passam com valores vazios. Injeta `X-User-Id`, `X-User-Name`, `X-User-Roles`.
- **Middleware metrics:** Coleta métricas Prometheus (contador de requisições, histograma de latência, cache hits/misses, erros). Expostas via `GET /metrics`.
- **Estrutura de erro:** Retorno consistente em JSON com mensagens descritivas (400 para payload inválido, 500 para erros internos).
- **Sanitização de erros:** Em erros 500, a mensagem exposta ao cliente é genérica (`Falha interna no processamento dos impostos`); detalhes internos ficam apenas nos logs (`slog.Error`).

**Fonte:** `cmd/api/main.go:72-78` (middleware pipeline), `cmd/api/main.go:130-153` (error handling)

## Validação de Entrada

- **Validação de payload:** `input.Validate()` é chamado no início do processamento do engine
- **Retorno de erros de validação:** O engine retorna resposta vazia + slice de erros se a validação falhar
- **HTTP 400:** Payload JSON inválido retorna `{"error": "Payload inválido", "details": "<err>"}`

**Fonte:** `internal/calculator/engine.go:44-46`, `cmd/api/main.go:130-136`

## Respostas da API

### Sucesso (200)
```json
{
  "IDTransaction": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "Itens": [
    {
      "SKU": "ITEM-001",
      "Total": "1000.00",
      "Tributos": [
        {
          "Tributo": "IPI",
          "CST": "50",
          "BaseCalculo": "1050.00",
          "Aliquota": "5.00",
          "Valor": "52.50",
          "MoreNumericDetails": [
            {"Key": "valor_item", "Value": "1000.00"},
            {"Key": "fator_rateio", "Value": "1.0"}
          ],
          "MoreTextDetails": [
            {"Key": "metodo_calculo", "Value": "ad_valorem"},
            {"Key": "fonte_regra", "Value": "repositorio"}
          ]
        },
        {
          "Tributo": "ICMS",
          "CST": "000",
          "BaseCalculo": "1052.50",
          "Aliquota": "18.00",
          "Valor": "189.45"
        },
        {
          "Tributo": "PIS",
          "CST": "01",
          "BaseCalculo": "863.05",
          "Aliquota": "1.65",
          "Valor": "14.24"
        },
        {
          "Tributo": "COFINS",
          "CST": "01",
          "BaseCalculo": "863.05",
          "Aliquota": "7.60",
          "Valor": "65.59"
        },
        {
          "Tributo": "CBS",
          "CST": "01",
          "BaseCalculo": "1000.00",
          "Aliquota": "3.52",
          "Valor": "35.20"
        },
        {
          "Tributo": "IBS",
          "CST": "01",
          "BaseCalculo": "1000.00",
          "Aliquota": "4.52",
          "Valor": "45.20"
        }
      ]
    }
  ],
  "TotalImpostos": "402.18",
  "TotalNota": "1000.00"
}
```

### Erro de Validação (400)
```json
{
  "error": "Payload inválido",
  "details": "invalid character '}' looking for beginning of object key string"
}
```

### Erro Interno (500)
```json
{
  "error": "Falha interna no processamento dos impostos"
}
```

## Métricas e Observabilidade

- **GET /metrics:** Endpoint Prometheus text exposition format com 4 métricas:
  - `http_requests_total{method,path,status}` — contador de requisições
  - `http_request_duration_seconds_count/sum/bucket{method,path}` — histograma de latência
  - `cache_requests_total{result="hit"|"miss"}` — contador de cache
  - `errors_total{type="validation"|"internal"}` — contador de erros por tipo
- **GET /healthz:** Liveness probe — retorna `200 {"status": "ok"}`
- **GET /health:** Readiness probe — verifica PostgreSQL + Redis, retorna `200` ou `503`
- **Structured logging:** `log/slog` JSON com trace_id e request_id em todas as mensagens

**Fonte:** `internal/middleware/metrics.go`, `cmd/api/main.go:81-118`
