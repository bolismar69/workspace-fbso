---
title: "Arquitetura — web_app-tax-nexus-portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["architecture", "frontend", "react", "spa"]
---

# Visão Geral da Arquitetura — TaxNexus Portal

## Resumo Executivo

O **TaxNexus Portal (TaaS)** é uma Single Page Application (SPA) React 19 que fornece uma interface de simulação tributária para a Reforma Tributária brasileira de 2026. O portal consome uma API backend Go (`ms-billing-engine-tax-rates`) para cálculos fiscais e apresenta resultados comparativos (sistema legado vs. IVA Dual) com visualizações gráficas via Recharts.

**Modo de consumo:** SPA via navegador → API REST (hardcoded `localhost:8080`).  
**Padrão arquitetural:** Component-based SPA simples com hooks customizados.  
**Estado da aplicação:** Estado local via `useState` — sem estado global.  
**Estágio atual:** MVP funcional (Fase 0 concluída, Fase 1 em andamento).

---

## Stack Tecnológica

| Camada | Tecnologia | Versão | Propósito |
|---|---|---|---|
| Runtime | Node.js | 18+ (Docker) | Build e desenvolvimento |
| Framework UI | React | 19.2.4 | Biblioteca de componentes |
| Linguagem | TypeScript | 5.9.3 | Tipagem estática |
| Bundler | Vite | 8.0.0 | Build e HMR para desenvolvimento |
| Gráficos | Recharts | 3.8.0 | Visualização de dados (gráfico de barras) |
| Ícones | Lucide React | 0.577.0 | Ícones SVG |
| Linter | ESLint | 9.39.4 | Qualidade de código |
| Runtime prod | Nginx | stable-alpine | Servidor web de produção |
| Containerização | Docker | — | Build multi-stage e deploy |

---

## Estrutura de Diretórios

```
web_app-tax-nexus-portal/
├── public/                     # Assets estáticos
│   ├── favicon.svg
│   └── icons.svg
├── src/
│   ├── main.tsx                # Entry point — ReactDOM.createRoot
│   ├── App.tsx                 # Shell da aplicação (auth + layout)
│   ├── index.css               # Estilos globais (design tokens CSS)
│   ├── components/
│   │   └── TaxSimulator.tsx    # Componente principal de simulação
│   ├── hooks/
│   │   └── useTaxService.ts    # Hook de chamada à API de cálculo
│   ├── api/                    # [VAZIO] — Diretório planejado para clientes HTTP
│   ├── services/               # [VAZIO] — Diretório planejado para serviços
│   ├── store/                  # [VAZIO] — Diretório planejado para estado global
│   └── pages/                  # [VAZIO] — Diretório planejado para páginas (router)
├── Dockerfile                  # Build multi-stage (Node 18 → Nginx)
├── nginx.conf                  # Configuração do Nginx (SPA fallback)
├── vite.config.ts              # Configuração do Vite + plugin React
├── tsconfig.json               # Configuração base do TypeScript
├── tsconfig.app.json           # Config TS para código da aplicação
├── tsconfig.node.json          # Config TS para tooling (Vite)
├── eslint.config.js            # Configuração do ESLint (flat config)
├── index.html                  # HTML shell (entry point do Vite)
└── package.json                # Dependências e scripts
```

---

## Princípios de Design

| Princípio | Aplicação |
|---|---|
| **Simplicidade** | Sem router, sem estado global — apenas `useState` local. Adequado ao escopo MVP. |
| **Separação de concerns** | `hooks/` para lógica de API, `components/` para UI. Diretórios `api/`, `services/`, `store/` já criados como scaffolding. |
| **Type Safety** | TypeScript strict mode com interfaces explícitas para request/response da API. |
| **SPA Fallback** | Nginx configurado com `try_files $uri /index.html` para suporte a client-side routing futuro. |

---

## Cross-cutting Concerns

| Concern | Implementação atual |
|---|---|
| **Autenticação** | Pseudo-auth via CNPJ (validação local de 14 dígitos). Sem JWT/OAuth. |
| **Logging** | `console.error` em caso de falha na API. Sem sistema de logging estruturado. |
| **Error Handling** | Try/catch no hook `useTaxService`. Retorna `null` em caso de erro. |
| **Loading States** | `loading: boolean` exposto pelo hook. Botão desabilitado durante fetch. |
| **Observabilidade** | Não implementada (sem tracing, métricas ou monitoramento de erros). |
| **Cache** | Não implementado. Cada simulação faz uma nova chamada à API. |

---

## Diagrama de Fluxo de Dados

```
[Usuário] → [App.tsx] → [TaxSimulator.tsx] → [useTaxService.ts] → [API Backend :8080]
                │                │                      │
                │                │                      ├── Request: {cnpj, ncm, ibge, saldo_remanescente}
                │                │                      └── Response: {transaction_status, calculation, callback}
                │                │
                │                └── Renderiza: Form (UF/Cidade/NCM/Saldo) + Gráfico Recharts
                │
                └── Estado: authenticated (boolean) — controla exibição de Login vs. Simulador
```

🤖 *Documentação gerada por mineração reversa de especificações (spec-miner).*
