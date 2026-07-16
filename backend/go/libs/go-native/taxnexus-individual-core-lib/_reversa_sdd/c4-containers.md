# C4 — Nível 2: Containers

> Gerado pelo **Arquiteto** (Reversa) em 2026-06-10 · `doc_level = completo`
> Confiança: 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

A biblioteca não é um container executável independente — ela é **linkada in-process** ao serviço de cálculo. Os containers de runtime relevantes são o serviço consumidor, os dois data stores e a própria lib como dependência compilada.

```mermaid
C4Container
    title Containers — taxnexus-core-lib em runtime

    Person(contribuinte, "Contribuinte", "Solicita simulação")

    System_Boundary(b, "Plataforma TaxNexus") {
        Container(api, "Serviço de Cálculo", "Go (fora do recorte)", "Expõe POST /api/v1/calculate/irpf; aplica fórmula e regras de período/ativação/autorização")
        Container(lib, "taxnexus-core-lib", "Biblioteca Go 1.25.6", "TaxRepository: GetApplicableRule, GetConfig, GetTableConfigs, GetTaxRulesForPeriod")
        ContainerDb(pg, "PostgreSQL", "schema individual_tax_rates", "tax_definitions, tax_rules_history, tax_configs (+ tax_calculation_log inferida)")
        ContainerDb(redis, "Redis", "go-redis 9.18", "Chaves tax_configs:* e tax_rules_list:* com TTL 12h")
    }

    Rel(contribuinte, api, "POST /api/v1/calculate/irpf", "HTTPS/JSON")
    Rel(api, lib, "Invoca funções do repositório", "in-process (Go)")
    Rel(lib, pg, "SELECT com filtro de vigência", "pgxpool / SQL")
    Rel(lib, redis, "GET/SET cache-aside", "RESP")
    Rel(api, pg, "Persiste log de cálculo (inferido)", "SQL")

    UpdateRelStyle(api, lib, $offsetX="-40")
```

## Containers

| Container | Tecnologia | Responsabilidade | Confiança |
|-----------|------------|------------------|-----------|
| Serviço de Cálculo | Go (fora do recorte) | API REST, fórmula, cenários, período, autorização | 🟡 (existência 🟢) |
| **taxnexus-core-lib** | Go 1.25.6 (biblioteca) | Leitura versionada de parâmetros + cache | 🟢 |
| PostgreSQL | RDBMS, schema `individual_tax_rates` | Fonte de verdade dos parâmetros fiscais | 🟢 |
| Redis | Cache em memória | Aceleração de leitura (TTL 12h) | 🟢 |

## Configuração de runtime

| Variável | Consumida por | Uso | Confiança |
|----------|---------------|-----|-----------|
| `DATABASE_URL` | `db.ConnectPostgres` | String de conexão pgx | 🟢 |
| `REDIS_ADDR` | `cache.ConnectRedis` | Endereço do Redis | 🟢 |

> A escrita em `tax_calculation_log` (seta tracejada) é **inferida** — nenhuma query desta lib a executa (L2). Pertence ao serviço de cálculo.
