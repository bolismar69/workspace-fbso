# db — Tarefas de Implementação

> Sequência executável para reimplementar a fábrica de pool PostgreSQL.
> Gerado pelo **Redator** (Reversa) em 2026-06-10 · `doc_level = completo`
> Fonte primária: `db/postgres.go`. Ver `requirements.md` e `design.md` desta unit.
> 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

## Pré-requisitos

- [ ] Dependência `github.com/jackc/pgx/v5/pgxpool` (v5.9.1) disponível
- [ ] Variável de ambiente `DATABASE_URL` documentada para o chamador

## Tarefas

- [ ] **T-01** — Implementar `ConnectPostgres(connString string) (*pgxpool.Pool, error)`: `pgxpool.ParseConfig` + `pgxpool.NewWithConfig(context.Background(), config)`, propagando erros.
  - Origem no legado: `db/postgres.go:9`
  - Critério de pronto: string válida retorna pool utilizável; string inválida retorna erro e pool nil (sem panic)
  - Confiança: 🟢

## Tarefas de Teste

- [ ] **TT-01** — Happy path: connection string válida retorna pool não-nil
- [ ] **TT-02** — Erro: connection string malformada retorna erro descritivo

## Ordem Sugerida

1. **T-01** — unit autônoma; sem dependências internas.

## Lacunas Pendentes (🔴)

- Tuning do pool (max conns, lifetime, timeouts) — confirmar se vem do DSN ou de configuração externa na operação.
- Avaliar incluir `Ping`/health-check na criação se a inicialização precisar falhar rápido (decisão de operação).
