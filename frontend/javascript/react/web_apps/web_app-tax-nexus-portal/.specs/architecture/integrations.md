---
title: "Integrações — web_app-tax-nexus-portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["integrations", "api", "dependencies"]
---

# Integrações Externas — TaxNexus Portal

## 1. API de Cálculo Tributário

| Atributo | Valor |
|---|---|
| **Serviço** | `ms-billing-engine-tax-rates` |
| **Endpoint** | `POST http://localhost:8080/v1/tax/calculate` |
| **Propósito** | Calcular comparativo de tributos (legado vs. reforma) |
| **Protocolo** | HTTP/1.1 REST |
| **Content-Type** | `application/json` |
| **Timeout** | Não configurado (usa default do `fetch`) |
| **Retry** | Não implementado |
| **Circuit Breaker** | Não implementado |
| **Autenticação** | Nenhuma (MVP) |

### Request Schema

```typescript
interface TaxRequest {
  cnpj: string;              // 14 dígitos
  ncm: string;               // 8 dígitos (ex: "62011100")
  ibge: string;              // 7 dígitos (ex: "3550308" = São Paulo)
  saldo_remanescente: number; // Crédito acumulado em R$
}
```

### Response Schema

```typescript
interface TaxResponse {
  transaction_status: string;
  calculation: {
    municipio: string;
    uf: string;
    ncm: string;
    ano: number;
    // Tributos Legado
    pis: number;
    cofins: number;
    icms: number;
    iss: number;
    ipi: number;
    // Tributos Reforma
    cbs_calculada: number;
    ibs_calculado: number;
    imposto_seletivo: number;
    ipva_novo: number;
    itcmd_novo: number;
    total?: number;
  };
  callback: {
    id_cadastro_unico: string;
    status_integracao: string;
  };
}
```

### Códigos de Erro Esperados

| HTTP Status | Significado | Tratamento no Frontend |
|---|---|---|
| 200 | Sucesso | Dados processados e exibidos |
| 400 | Dados inválidos | `console.error` do `errorData`, retorna `null` |
| 500 | Erro interno do backend | `console.error`, retorna `null` |
| Network Error | API indisponível | `console.error("Erro na chamada Fetch:", error)`, retorna `null` |

### Localização no Código

- Cliente HTTP: `src/hooks/useTaxService.ts:46` — `fetch('http://localhost:8080/v1/tax/calculate', ...)`
- URL hardcoded — sem variável de ambiente ou configuração externa.

---

## 2. Dependências de Pacotes (npm)

### Produção

| Pacote | Versão | Propósito | Licença |
|---|---|---|---|
| `react` | ^19.2.4 | Biblioteca UI | MIT |
| `react-dom` | ^19.2.4 | Renderização DOM | MIT |
| `recharts` | ^3.8.0 | Gráficos (BarChart, ResponsiveContainer) | MIT |
| `lucide-react` | ^0.577.0 | Ícones SVG (não usado atualmente no código) | ISC |

### Desenvolvimento

| Pacote | Versão | Propósito |
|---|---|---|
| `typescript` | ~5.9.3 | Type checker |
| `vite` | ^8.0.0 | Bundler + dev server |
| `@vitejs/plugin-react` | ^6.0.0 | Suporte React no Vite (Oxc) |
| `eslint` | ^9.39.4 | Linter |
| `typescript-eslint` | ^8.56.1 | Linting TypeScript |
| `eslint-plugin-react-hooks` | ^7.0.1 | Regras de hooks React |
| `eslint-plugin-react-refresh` | ^0.5.2 | HMR com eslint |
| `@types/react` | ^19.2.14 | Tipos React |
| `@types/react-dom` | ^19.2.3 | Tipos ReactDOM |
| `@types/node` | ^24.12.0 | Tipos Node.js |

---

## 3. Integrações de Infraestrutura

### Docker

| Atributo | Valor |
|---|---|
| **Build** | Multi-stage: `node:18-alpine` (builder) → `nginx:stable-alpine` (runtime) |
| **Porta exposta** | 5173 |
| **Config Nginx** | `nginx.conf` — SPA fallback (`try_files $uri /index.html`) |
| **Build cmd** | `npm ci` → `npm run build` (tsc + vite build) |

### Ambiente de Desenvolvimento

| Atributo | Valor |
|---|---|
| **Dev server** | Vite dev server na porta 5173 |
| **HMR** | Hot Module Replacement via Vite |
| **Backend** | Esperado em `localhost:8080` |

---

## 4. Dívidas Técnicas de Integração

| ID | Descrição | Severidade |
|---|---|---|
| INT-001 | URL do backend hardcoded (`localhost:8080`) — sem suporte a ambientes (staging, produção) | Alta |
| INT-002 | Sem timeout configurado no `fetch` — pode travar indefinidamente | Média |
| INT-003 | Sem retry ou circuit breaker — falha de rede não é tratada com resiliência | Média |
| INT-004 | Sem autenticação na chamada à API — backend exposto sem proteção | Alta |
| INT-005 | `lucide-react` listado como dependência mas não utilizado no código | Baixa |

🤖 *Documentação gerada por mineração reversa de especificações (spec-miner).*
