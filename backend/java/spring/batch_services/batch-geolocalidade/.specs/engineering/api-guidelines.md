---
title: "API Guidelines — batch-geolocalidade"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["engineering", "api-guidelines", "batch"]
---

# API Guidelines — batch-geolocalidade

## Visão Geral

**batch-geolocalidade** é um serviço **Spring Batch headless** — ele NÃO expõe endpoints REST. A "API" do serviço é sua interface de linha de comando e o comportamento determinístico de seus Jobs.

Este documento define as guidelines para quando endpoints REST forem adicionados (versão futura).

## Interface Atual (CLI)

### Execução

```bash
java -jar target/spring-batch-geolocalidade-0.0.1-SNAPSHOT.jar \
  --app.loadtest.enabled=true
```

### Exit Codes

| Exit Code | Significado |
|---|---|
| `0` | Job concluído com sucesso (`BatchStatus.COMPLETED`) |
| `1` | Job falhou (`BatchStatus != COMPLETED`) |

### Logs de Execução

- `IBGE file [municipios] path=..., exists=true/false`
- `IBGE file [distritos] path=..., exists=true/false`
- `IBGE file [subdistritos] path=..., exists=true/false`
- `Load test finished. status=COMPLETED, durationMs=..., municipios=5570, distritos=10407, subdistritos=684`

## Guidelines para Endpoints REST Futuros

### Padrão de URL

```
/api/v1/jobs/{jobName}/launch
/api/v1/jobs/{jobName}/executions
/api/v1/jobs/{jobName}/executions/{executionId}
```

### Formato de Resposta

```json
{
  "data": { },
  "meta": {
    "timestamp": "2026-07-08T12:00:00Z",
    "version": "0.0.1"
  }
}
```

### Códigos de Erro

| HTTP Status | Caso |
|---|---|
| `202 Accepted` | Job disparado com sucesso |
| `400 Bad Request` | Parâmetros inválidos |
| `404 Not Found` | Job ou execução não encontrada |
| `409 Conflict` | Job já está em execução |
| `500 Internal Server Error` | Erro inesperado |

### Autenticação (Planejada)

- Spring Security com JWT Bearer Token
- RBAC: role `BATCH_OPERATOR` para disparar jobs
