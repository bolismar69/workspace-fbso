# C4 — Nível 3: Componentes (container `taxnexus-core-lib`)

> Gerado pelo **Arquiteto** (Reversa) em 2026-06-10 · `doc_level = completo`
> Confiança: 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

Detalhamento interno da biblioteca: o pacote `repository` é o componente com lógica; `db` e `cache` são fábricas de infraestrutura injetadas; `models` fornece os tipos.

```mermaid
C4Component
    title Componentes — taxnexus-core-lib

    Container_Ext(api, "Serviço de Cálculo", "Go", "Consumidor in-process")

    Container_Boundary(lib, "taxnexus-core-lib") {
        Component(repo, "TaxRepository", "repository/tax_repository.go", "Resolve faixa em memória, lê configs, aplica vigência temporal, orquestra cache-aside")
        Component(models, "models", "models/tax_models.go", "TaxDefinition, TaxRule, TaxCalculationLog + DTOs (TaxRequest, TaxResponse, ...)")
        Component(dbf, "ConnectPostgres", "db/postgres.go", "Fábrica de *pgxpool.Pool a partir de DATABASE_URL")
        Component(cachef, "ConnectRedis", "cache/redis.go", "Fábrica de *redis.Client a partir de REDIS_ADDR")
    }

    ContainerDb_Ext(pg, "PostgreSQL", "individual_tax_rates")
    ContainerDb_Ext(redis, "Redis", "cache TTL 12h")

    Rel(api, repo, "NewTaxRepository(db, rdb) + chamadas Get*")
    Rel(api, dbf, "ConnectPostgres(DATABASE_URL)")
    Rel(api, cachef, "ConnectRedis(REDIS_ADDR)")
    Rel(repo, models, "Usa structs e (de)serializa")
    Rel(repo, pg, "Queries com filtro de vigência", "pgxpool")
    Rel(repo, redis, "GET/SET cache-aside", "go-redis")
    Rel(dbf, pg, "Cria pool")
    Rel(cachef, redis, "Cria client")

    UpdateRelStyle(api, repo, $offsetY="-30")
```

## Componentes do `repository`

| Operação | Cache | Lógica chave | Confiança |
|----------|-------|--------------|-----------|
| `GetApplicableRule` | indireto (via `GetTaxRulesForPeriod`) | Seleção da **primeira** faixa cujo `[range_min, range_max]` contém a base; `range_max NULL` = faixa aberta; erro se nenhuma casa | 🟢 |
| `GetConfig` | **não** | Lê um `config_value` único por `config_key`; retorna `(Zero, err)` em falha | 🟢 |
| `GetTableConfigs` | sim (`tax_configs:<code>:<YYYY-MM-DD>`) | Mapa `config_key → config_value` vigente | 🟢 |
| `GetTaxRulesForPeriod` | sim (`tax_rules_list:<code>:<YYYY-MM-DD>`) | Lista faixas vigentes (JOIN `tax_definitions`), ordenadas por `range_min ASC` | 🟢 |

## Padrões transversais

- **Cache-aside** (A3): GET → miss → Postgres → SET (TTL 12h); erro de Redis não aborta (degradação graciosa). 🟢
- **Vigência temporal** (A2): `valid_from <= refDate AND (valid_to IS NULL OR valid_to >= refDate)` em todas as queries. 🟢
- **Decimal exato** (RN-08): todo valor monetário/alíquota é `decimal.Decimal`. 🟢
- **DI por construtor**: `NewTaxRepository(pool, client)`; `db`/`cache` são fábricas puras. 🟢

> Faixas garantidamente contíguas e sem sobreposição (🟢 D6) — premissa que valida a corretude do "primeiro match vence" em `GetApplicableRule`.
