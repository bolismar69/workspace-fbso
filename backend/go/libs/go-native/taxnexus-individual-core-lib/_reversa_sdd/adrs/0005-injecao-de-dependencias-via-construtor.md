# ADR-0005 — Injeção de dependências via construtor (sem singletons globais)

> ADR **retroativo** inferido do código (sem histórico Git). · Detetive · 2026-06-10
> Confiança: 🟢 CONFIRMADO

## Status
Aceito (vigente).

## Contexto
A biblioteca precisa de um pool PostgreSQL e um cliente Redis. Havia a opção de variáveis/singletons globais (comum em código Go pequeno) ou injeção explícita.

## Decisão
`TaxRepository` recebe `*pgxpool.Pool` e `*redis.Client` no construtor `NewTaxRepository(db, rdb)` e os guarda como campos privados. Não há estado global; a fábrica de conexões fica nos pacotes `db` e `cache`, separada do repositório.
- Evidência: `tax_repository.go:16-23`; `db/postgres.go`; `cache/redis.go`.

## Consequências
- 🟢 Testabilidade: dá para injetar pools/clients de teste (ou fakes que satisfaçam as interfaces concretas — ver ressalva abaixo).
- 🟢 Aplicação consumidora controla o ciclo de vida das conexões.
- 🟡 As dependências são **tipos concretos** (`*pgxpool.Pool`, `*redis.Client`), não interfaces — dificulta mock puro em testes unitários sem um banco/redis real. Possível melhoria futura (definir interfaces). 🟡
- 🟢 Sem acoplamento a configuração global; `connString`/`addr` vêm de fora (12-factor).
