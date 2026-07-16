---
title: "API Guidelines — web_app-tax-nexus-portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["api", "guidelines", "how-to"]
---

# Padrões de Consumo de API — TaxNexus Portal

## 1. Formato de Requisição

### Headers obrigatórios

```
Content-Type: application/json
Accept: application/json
```

### Método HTTP

Apenas `POST` é utilizado na comunicação com o backend.

### Encoding

UTF-8 — `Content-Type: application/json`.

---

## 2. Estrutura de Payload

### Request Envelope

```typescript
// Todas as requisições seguem este padrão
const response = await fetch(url, {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  },
  body: JSON.stringify(payload)
});
```

### Response Envelope

```typescript
// Resposta de sucesso
{
  transaction_status: "completed",
  calculation: { /* dados do cálculo */ },
  callback: {
    id_cadastro_unico: string,
    status_integracao: "registered" | "pending" | "failed"
  }
}

// Resposta de erro (inferido)
{
  error: string,
  details?: object
}
```

---

## 3. Tratamento de Erros

### Padrão Atual

```typescript
// src/hooks/useTaxService.ts
try {
  const response = await fetch(url, options);
  if (!response.ok) {
    const errorData = await response.json();
    console.error("Erro do servidor:", errorData);
    return null;
  }
  return await response.json();
} catch (error) {
  console.error("Erro na chamada Fetch:", error);
  return null;
}
```

### Problemas Identificados

| ID | Problema | Recomendação |
|---|---|---|
| ERR-001 | `errorData` assumido como JSON — pode falhar se resposta não for JSON | Verificar `Content-Type` antes de `response.json()` |
| ERR-002 | Erro de rede e erro HTTP tratados de forma diferente mas ambos retornam `null` | Diferenciar erros para permitir retry vs. user feedback |
| ERR-003 | Sem `AbortController` — chamada não é cancelada ao desmontar componente | Adicionar `AbortController` com cleanup no `useEffect` |
| ERR-004 | `console.error` como única estratégia de logging | Adotar serviço de logging estruturado (ex: Sentry, Datadog) |

### Padrão Recomendado

```typescript
const calculateTax = async (payload: TaxRequest, signal?: AbortSignal): Promise<TaxResponse> => {
  const response = await fetch(url, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json', 'Accept': 'application/json' },
    body: JSON.stringify(payload),
    signal,
  });

  if (!response.ok) {
    const contentType = response.headers.get('content-type');
    const errorBody = contentType?.includes('application/json')
      ? await response.json()
      : await response.text();
    throw new ApiError(response.status, errorBody);
  }

  return response.json();
};
```

---

## 4. Versionamento

### Atual

- **URL:** `http://localhost:8080/v1/tax/calculate`
- **Versão:** v1 (prefixo de path)
- **Sem negociação de versão** (Accept header não usado para versionamento)

### Recomendação

- Manter versionamento por path (`/v1/`, `/v2/`)
- Backend deve retornar header `X-API-Version: 1.0.0`

---

## 5. Segurança

### Estado Atual (MVP)

- **Sem autenticação** na chamada à API
- **Sem HTTPS** (localhost)
- **Sem CORS configurado** no frontend

### Recomendações

| ID | Recomendação | Prioridade |
|---|---|---|
| SEC-001 | Adicionar JWT Bearer token nas requisições | P1 |
| SEC-002 | Usar HTTPS em produção | P1 |
| SEC-003 | Configurar CORS no backend para origem específica | P2 |
| SEC-004 | Adicionar rate limiting no backend | P2 |
| SEC-005 | Não expor `id_cadastro_unico` em logs client-side | P3 |

---

## 6. Observabilidade

### Estado Atual

- Apenas `console.error` para erros
- Sem métricas de latência/taxa de sucesso
- Sem tracing distribuído

### Recomendações

- Instrumentar chamadas fetch com medição de latência
- Reportar erros para serviço de monitoramento (Sentry, Datadog RUM)
- Adicionar header `X-Request-ID` para correlacionar logs frontend/backend

---

## 7. Configuração de Ambiente

### Problema

A URL do backend está hardcoded:

```typescript
// src/hooks/useTaxService.ts:46
const response = await fetch('http://localhost:8080/v1/tax/calculate', ...);
```

### Recomendação

Usar variáveis de ambiente Vite:

```typescript
// vite.config.ts
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/v1': 'http://localhost:8080'
    }
  }
});

// useTaxService.ts
const API_BASE = import.meta.env.VITE_API_BASE || '';
const response = await fetch(`${API_BASE}/v1/tax/calculate`, ...);
```

🤖 *Documentação gerada por mineração reversa de especificações (spec-miner).*
