# Dependências — taxnexus-individual-core-lib

> Gerado pelo **Scout** (Reversa) em 2026-06-10
> Fonte: `go.mod` + `go.sum`

---

## Ambiente

| Item | Valor |
|------|-------|
| Gerenciador de pacotes | **Go Modules** (`go.mod` / `go.sum`) 🟢 |
| Versão do Go | **1.25.6** 🟢 |
| Módulo | `taxnexus-core-lib` 🟢 |

---

## Dependências diretas

| Pacote | Versão | Propósito | Onde é usado |
|--------|--------|-----------|--------------|
| `github.com/shopspring/decimal` | v1.4.0 | Aritmética decimal de precisão (valores monetários e alíquotas, evitando float) | `models`, `repository` 🟢 |
| `github.com/jackc/pgx/v5` | v5.9.1 | Driver/pool PostgreSQL (`pgxpool`) | `db`, `repository` 🟢 |
| `github.com/redis/go-redis/v9` | v9.18.0 | Cliente Redis para cache de regras e configs | `cache`, `repository` 🟢 |
| `github.com/google/uuid` | v1.6.0 | Identificadores UUID (campo `TaxCalculationLog.ID`) | `models` 🟢 |

---

## Dependências indiretas (// indirect)

| Pacote | Versão | Trazido por |
|--------|--------|-------------|
| `github.com/cespare/xxhash/v2` | v2.3.0 | go-redis |
| `github.com/dgryski/go-rendezvous` | v0.0.0-20200823014737 | go-redis (sharding) |
| `github.com/jackc/pgpassfile` | v1.0.0 | pgx |
| `github.com/jackc/pgservicefile` | v0.0.0-20240606120523 | pgx |
| `github.com/jackc/puddle/v2` | v2.2.2 | pgx (pool) |
| `go.uber.org/atomic` | v1.11.0 | transitiva |
| `golang.org/x/sync` | v0.17.0 | transitiva (pgx/puddle) |
| `golang.org/x/text` | v0.29.0 | transitiva |

---

## Dependências de teste presentes apenas no `go.sum`

> Não há testes próprios neste repositório. Estes aparecem no `go.sum` como dependências transitivas de teste das libs acima (não fazem parte do build da biblioteca).

- `github.com/stretchr/testify` (v1.11.1) — usado internamente por go-redis/pgx em seus testes
- `github.com/bsm/ginkgo/v2` (v2.12.0) + `github.com/bsm/gomega` (v1.27.10) — suíte de testes do go-redis
- `github.com/davecgh/go-spew`, `github.com/pmezard/go-difflib`, `github.com/stretchr/objx` — transitivas de testify
- `github.com/zeebo/xxh3`, `github.com/klauspost/cpuid/v2` — transitivas de xxhash
- `gopkg.in/yaml.v3`, `gopkg.in/check.v1` — transitivas

---

## Infraestrutura externa requerida

| Serviço | Evidência | Configuração |
|---------|-----------|--------------|
| **PostgreSQL** | `pgxpool`, queries no schema `individual_tax_rates` | `DATABASE_URL` (env) 🟢 |
| **Redis** | `go-redis`, cache com TTL 12h | `REDIS_ADDR` (env) 🟢 |

> 🟡 Nenhuma dependência aponta para HTTP server, gRPC, filas ou outras integrações — reforça que é uma **biblioteca de dados pura**, sem camada de API neste recorte.
