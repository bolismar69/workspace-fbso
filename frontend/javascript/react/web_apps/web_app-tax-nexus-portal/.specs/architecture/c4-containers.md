---
title: "C4 — Containers — web_app-tax-nexus-portal"
level: "Containers"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
---

# C4 — Nível 2: Containers

## Diagrama

```mermaid
C4Container
    title Diagrama de Containers — TaxNexus Portal

    Person(contribuinte, "Contribuinte PJ", "Usuário final")
    
    System_Boundary(taxnexus_system, "TaxNexus Portal") {
        Container(nginx, "Nginx Reverso", "stable-alpine", "Serve arquivos estáticos compilados<br/>Configurado com SPA fallback (try_files)")
        Container(spa, "SPA React", "React 19 + TypeScript + Vite", "Aplicação single-page com:<br/>- App.tsx (shell + auth)<br/>- TaxSimulator.tsx (formulário + gráfico)<br/>- useTaxService.ts (hook de API)")
    }
    
    System_Boundary(backend_system, "Backend de Cálculo Tributário") {
        Container(api, "API Go/Fiber", "Go 1.22 + Fiber", "Endpoint POST /v1/tax/calculate<br/>Processa cálculo comparativo de tributos")
        ContainerDb(db, "PostgreSQL", "PostgreSQL 16", "Armazena alíquotas, CSTs, CFOPs, NCMs<br/>e histórico de cálculos")
    }
    
    System_Ext(cadastro, "Cadastro Único", "Sistema corporativo de rastreabilidade")
    
    Rel(contribuinte, nginx, "HTTPS requisição", "TCP :5173")
    Rel(nginx, spa, "Serve assets", "Filesystem")
    Rel(spa, api, "POST /v1/tax/calculate", "HTTP :8080")
    Rel(api, db, "Leitura/Escrita", "TCP :5432")
    Rel(api, cadastro, "Callback registro", "HTTP")
    
    UpdateLayoutConfig($c4ShapeInRow="3", $c4BoundaryInRow="2")
```

## Elementos

| Nome | Tipo | Responsabilidade | Tecnologia |
|---|---|---|---|
| Nginx Reverso | Container (Web Server) | Servir bundle estático React com SPA fallback | Nginx stable-alpine |
| SPA React | Container (Frontend) | Interface de simulação tributária | React 19 + TS 5.9 + Vite 8 |
| API Go/Fiber | Container (Backend) | Cálculo de tributos comparativos | Go 1.22 + Fiber |
| PostgreSQL | Container (Database) | Persistência de alíquotas e histórico | PostgreSQL 16 |
| Cadastro Único | External System | Rastreabilidade corporativa | — |

## Fluxos Principais

### Fluxo: Carregamento Inicial
1. Navegador requisita `GET /` ao Nginx
2. Nginx serve `index.html` (SPA shell)
3. Navegador carrega bundle React (JS/CSS)
4. React hydrata e renderiza tela de login (CNPJ)

### Fluxo: Simulação
1. Usuário preenche formulário e clica "SIMULAR REFORMA TRIBUTÁRIA"
2. `useTaxService.calculateTax()` dispara `POST http://localhost:8080/v1/tax/calculate`
3. API Go processa no backend e consulta PostgreSQL para alíquotas
4. Resposta retorna ao frontend
5. `TaxSimulator` renderiza cards e gráfico Recharts

### Deploy
1. `docker build` executa multi-stage: Node 18 alpine (build) → Nginx stable-alpine (runtime)
2. Container expõe porta 5173
3. Nginx configurado com `try_files $uri /index.html` para client-side routing
