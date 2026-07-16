# API Guidelines — Solar Fácil Site

> Padrões de consumo de API, tratamento de erros e observabilidade no frontend.
> Gerado por `api-designer` + `documentation-writer` em 2026-07-08.

---

## 1. Status Atual

⚠️ **A aplicação consome apenas 1 endpoint externo** (formulário de contato via Formspree). Todos os dados de negócio são estáticos (constantes TypeScript + JSON mocks).

## 2. Padrão de Chamadas HTTP

### 2.1 Fetch Nativo

```typescript
// Padrão: fetch nativo, sem axios/ky
const response = await fetch(ENDPOINT, {
  method: 'POST',
  body: formData,
});

if (!response.ok) {
  throw new Error(`HTTP ${response.status}`);
}
```

### 2.2 Tratamento de Erros

```typescript
// Padrão: try-catch com mensagem em pt-BR
try {
  // ... operação
} catch {
  setSubmitError(
    'Não foi possível enviar. Verifique sua conexão ou fale conosco pelo WhatsApp.',
  );
}
```

### 2.3 Fallback Seguro

```typescript
// Serviços retornam array vazio em caso de erro (nunca quebram a UI)
try {
  setTimeout(() => resolve(DATA), 500);
} catch (error) {
  console.error("Erro => ", error);
  resolve([]); // Fallback seguro
}
```

## 3. Camada de Serviços (Pronta para API Real)

```typescript
// Formato atual (mock):
export async function fetchPlans(): Promise<Plan[]> {
  return new Promise((resolve) => {
    setTimeout(() => resolve(PLANS), 500);
  });
}

// Formato futuro (API real):
export async function fetchPlans(): Promise<Plan[]> {
  const response = await fetch(`${API_BASE}/plans`);
  if (!response.ok) throw new Error(`HTTP ${response.status}`);
  return response.json();
}
```

## 4. Variáveis de Ambiente

| Variável | Default | Uso |
|---|---|---|
| `NEXT_PUBLIC_FORM_ENDPOINT` | `https://formspree.io/f/placeholder` | Endpoint do formulário |
| `NEXT_PUBLIC_GA_ID` | `G-XXXXXXXXXX` | Google Analytics 4 |

**Regra**: Prefixo `NEXT_PUBLIC_` apenas para variáveis expostas ao browser.

## 5. Headers de Segurança

```typescript
// next.config.ts — headers atuais
{ key: 'X-Frame-Options', value: 'DENY' },
{ key: 'X-Content-Type-Options', value: 'nosniff' },
{ key: 'Referrer-Policy', value: 'strict-origin-when-cross-origin' },
```

### Headers Recomendados (ausentes)

```typescript
{ key: 'Content-Security-Policy', value: "default-src 'self'; script-src 'self' 'unsafe-inline' https://www.googletagmanager.com; style-src 'self' 'unsafe-inline'" },
{ key: 'Strict-Transport-Security', value: 'max-age=31536000; includeSubDomains' },
{ key: 'Permissions-Policy', value: 'camera=(), microphone=(), geolocation=()' },
```

## 6. Observabilidade

### 6.1 Logging

- **Client-side**: `console.error` nos catch blocks
- **Recomendação**: Migrar para `pino` ou `winston` server-side, Sentry/LogRocket client-side

### 6.2 Analytics

- 4 eventos GA4 definidos: `cta_click`, `calculator_use`, `faq_open`, `lead_capture`
- 2 eventos conectados: `cta_click`, `calculator_use`
- 2 eventos pendentes: `faq_open`, `lead_capture`

## 7. TODO — Quando o Backend Existir

- [ ] Definir `API_BASE` como variável de ambiente
- [ ] Adicionar timeout e retry policy
- [ ] Adicionar circuit breaker para falhas repetidas
- [ ] Migrar `console.error` para logging service
- [ ] Conectar eventos `faq_open` e `lead_capture`

---

Última atualização: 2026-07-08
