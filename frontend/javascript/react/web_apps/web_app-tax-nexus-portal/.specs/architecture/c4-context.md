---
title: "C4 — Context — web_app-tax-nexus-portal"
level: "Context"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
---

# C4 — Nível 1: Contexto do Sistema

## Diagrama

```mermaid
C4Context
    title Diagrama de Contexto — TaxNexus Portal (TaaS)

    Person(contribuinte, "Contribuinte PJ", "Empresa brasileira que deseja simular o impacto da Reforma Tributária 2026")
    
    System(taxnexus_portal, "TaxNexus Portal", "SPA React que permite simular a transição tributária<br/>com gráficos comparativos (Legado vs. IVA Dual)")
    
    System_Ext(ms_tax_rates, "ms-billing-engine-tax-rates", "API Go/Fiber de cálculo tributário<br/>Provê o cálculo comparativo dos tributos")
    
    System_Ext(cadastro_unico, "Cadastro Único Corporativo", "Sistema de registro que recebe o callback<br/>da simulação para rastreabilidade")
    
    Rel(contribuinte, taxnexus_portal, "Acessa via navegador", "HTTPS :5173")
    Rel(taxnexus_portal, ms_tax_rates, "POST /v1/tax/calculate", "HTTP :8080")
    Rel(ms_tax_rates, cadastro_unico, "Registra simulação", "Callback")
    
    UpdateLayoutConfig($c4ShapeInRow="3", $c4BoundaryInRow="1")
```

## Elementos

| Nome | Tipo | Responsabilidade | Tecnologia |
|---|---|---|---|
| Contribuinte PJ | Person (Ator) | Usuário final — informa CNPJ, NCM, UF/Cidade e saldo para simular | Navegador web |
| TaxNexus Portal | System | Interface de simulação tributária com formulário e gráficos comparativos | React 19 + TypeScript + Vite + Recharts |
| ms-billing-engine-tax-rates | External System | API de cálculo tributário — computa PIS/COFINS/ICMS/ISS/IPI (legado) vs. CBS/IBS/IS (reforma) | Go 1.22 + Fiber |
| Cadastro Único Corporativo | External System | Recebe callback com ID de cadastro único para rastreabilidade da simulação | Desconhecido |

## Fluxos Principais

### Fluxo 1: Simulação Tributária Completa
1. Contribuinte acessa o portal via navegador e informa CNPJ (14 dígitos)
2. Portal valida CNPJ localmente e libera acesso ao simulador
3. Contribuinte seleciona UF, Cidade (IBGE), NCM e Saldo Remanescente
4. Portal envia requisição `POST /v1/tax/calculate` ao backend Go
5. Backend processa cálculo e retorna tributos legados + reforma + callback
6. Portal exibe resultados em cards comparativos e gráfico de barras (2026 vs. 2027)
7. Backend notifica Cadastro Único Corporativo com ID de rastreio

### Fluxo 2: Tratamento de Erro
1. Se a API backend estiver indisponível ou retornar erro
2. Portal registra erro no console e retorna `null`
3. UI permanece no estado atual sem crash (graceful degradation)
