# C4 — Nível 1: Contexto

> Gerado pelo **Arquiteto** (Reversa) em 2026-06-10 · `doc_level = completo`
> Confiança: 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

O sistema central deste recorte é a biblioteca `taxnexus-core-lib`. Ela é consumida in-process por um **serviço de cálculo de imposto** (fora do recorte), que por sua vez expõe a API REST observada nos payloads de validação (`POST /api/v1/calculate/irpf` 🟢 D2).

```mermaid
C4Context
    title Contexto — taxnexus-individual-core-lib

    Person(contribuinte, "Contribuinte / Cliente", "Pessoa física que solicita simulação de IRPF/INSS")
    Person(admin, "Administrador fiscal", "Mantém faixas, definições e parâmetros (tax_definitions, tax_rules_history, tax_configs)")

    System_Boundary(b, "Plataforma TaxNexus") {
        System(svc, "Serviço de Cálculo de Imposto", "Aplica fórmula, compara cenários, controla active e período mensal/anual, autoriza requisições. Expõe POST /api/v1/calculate/irpf. [FORA DO RECORTE]")
        System(lib, "taxnexus-core-lib", "Biblioteca Go: leitura versionada de parâmetros fiscais (faixas, configs) com cache. [ESTE REPOSITÓRIO]")
    }

    SystemDb_Ext(pg, "PostgreSQL", "Schema individual_tax_rates: tax_definitions, tax_rules_history, tax_configs")
    System_Ext(redis, "Redis", "Cache-aside de configs e faixas (TTL 12h)")

    Rel(contribuinte, svc, "Solicita cálculo de imposto", "HTTPS/JSON")
    Rel(admin, pg, "Cadastra/versiona parâmetros fiscais", "SQL")
    Rel(svc, lib, "Lê faixas e parâmetros por refDate", "chamada de função Go")
    Rel(lib, pg, "Consulta parâmetros vigentes", "pgx / SQL")
    Rel(lib, redis, "Cache-aside (GET/SET TTL 12h)", "RESP")

    UpdateRelStyle(svc, lib, $offsetY="-10")
```

## Atores e sistemas

| Elemento | Tipo | Papel | Confiança |
|----------|------|-------|-----------|
| Contribuinte / Cliente | Persona | Solicita simulação via serviço consumidor | 🟡 |
| Administrador fiscal | Persona | Mantém os dados de `individual_tax_rates` | 🟡 |
| Serviço de Cálculo de Imposto | Sistema (externo ao recorte) | Cálculo, comparação de cenários, `active`, período, autorização | 🟢 (existência confirmada por D2/D5/D7/D8) |
| **taxnexus-core-lib** | Sistema (este repo) | Leitura versionada de parâmetros fiscais com cache | 🟢 |
| PostgreSQL | Sistema externo | Persistência dos parâmetros | 🟢 |
| Redis | Sistema externo | Cache | 🟢 |

## Fronteira de responsabilidade (decisão central)

O que está **dentro** da lib: resolver faixa por base/data, ler configs, vigência temporal, cache.
O que está **fora** (no serviço consumidor): fórmula `base × alíquota − deduction_val` (🟢 D3), escolha do cenário recomendado de menor imposto (🟢 D4), bloqueio por `active` (🟢 D5), diferença mensal/anual (🟢 D7), autorização e multi-tenant (🟢 D8).
