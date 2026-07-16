# db — Design Técnico

> Implementação da fábrica de pool PostgreSQL, com base em `db/postgres.go`.
> Gerado pelo **Redator** (Reversa) em 2026-06-10 · Atualizado pelo **Revisor** em 2026-06-12
> 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

## Interface

| Símbolo | Assinatura | Retorno | Observação |
|---------|-----------|---------|------------|
| `ConnectPostgres` | `(connString string)` | `(*pgxpool.Pool, error)` | Fábrica de pool; entry point de infraestrutura 🟢 |

## Fluxo Principal

1. Recebe `connString` (origem típica: env `DATABASE_URL`). `db/postgres.go` 🟢
2. `pgxpool.ParseConfig(connString)` → config; erro de parsing é retornado. `db/postgres.go:9` 🟢
3. `pgxpool.NewWithConfig(context.Background(), config)` cria o pool. 🟢
4. Retorna `(*pgxpool.Pool, error)` ao chamador. 🟢

## Fluxos Alternativos

- **String inválida:** `ParseConfig` retorna erro → propagado, pool nil. 🟢
- **Banco indisponível no momento da criação:** como não há `Ping`, a conexão é **lazy**; a falha aparece no primeiro uso pelo `repository`. 🟢

## Dependências

- **`github.com/jackc/pgx/v5/pgxpool`** (v5.9.1) — pool de conexões PostgreSQL. 🟢

## Decisões de Design Identificadas

| Decisão | Evidência no código | Confiança |
|---------|---------------------|-----------|
| Pool de conexões em vez de conexão única | `db/postgres.go:9` | 🟢 |
| `context.Background()` na criação (sem deadline) | `db/postgres.go:9` | 🟢 |
| Connection string injetada pelo chamador (não lê env internamente) | assinatura `ConnectPostgres(connString)` | 🟢 |

## Estado Interno

Nenhum no pacote. O `*pgxpool.Pool` retornado mantém o estado de conexões; seu ciclo de vida (Close) é responsabilidade do chamador. 🟢

## Observabilidade

Nenhuma emitida pelo pacote (sem logs/métricas). 🟢

## Riscos e Lacunas

- 🟡 **Sem `Ping`/health-check na criação:** falhas de conectividade só se manifestam no primeiro uso.
- 🟡 **Sem timeout configurável na criação:** uso de `context.Background()` significa que não há deadline para a montagem do pool.
- 🟢 **Tuning do pool:** confirmado como default do pgxpool (A3 não especificou parâmetros extras).
