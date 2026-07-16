# Fluxograma — módulo `db`

> Gerado pelo **Arqueólogo** (Reversa) em 2026-06-10
> Fonte: `db/postgres.go`

## `ConnectPostgres`

```mermaid
flowchart TD
    A[ConnectPostgres connString] --> B[ctx = context.Background]
    B --> C[pgxpool.ParseConfig connString]
    C --> D{erro?}
    D -- sim --> E[retorna nil, err]
    D -- não --> F[pgxpool.NewWithConfig ctx, config]
    F --> G[retorna pool, err]
```

> Observações:
> - Sem timeout/deadline no contexto de criação. 🟡
> - Sem ajuste de parâmetros de pool (usa defaults do pgx). 🟡
> - `connString` provém de `DATABASE_URL`. 🟢
