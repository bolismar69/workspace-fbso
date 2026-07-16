# Code/Spec Matrix — taxnexus-individual-core-lib

> Mapeia cada arquivo do legado à unit de spec que o cobre.
> Gerado pelo **Redator** (Reversa) em 2026-06-10 · `doc_level = completo`
> 🟢 coberto · 🟡 coberto parcialmente · n/a sem unit (candidato a análise adicional)

## Arquivos de código

| Arquivo do legado | Unit correspondente | Símbolos cobertos | Cobertura |
|-------------------|---------------------|-------------------|-----------|
| `repository/tax_repository.go` | `repository/` | `NewTaxRepository`, `GetApplicableRule`, `GetConfig`, `GetTableConfigs`, `GetTaxRulesForPeriod` | 🟢 |
| `models/tax_models.go` | `models/` | `TaxDefinition`, `TaxRule`, `TaxCalculationLog`, `TaxRequest`, `UniversalTaxRequest`, `documentoFiscalRequest`, `TaxResponse`, `DeductionDetail` | 🟢 |
| `db/postgres.go` | `db/` | `ConnectPostgres` | 🟢 |
| `cache/redis.go` | `cache/` | `ConnectRedis` | 🟢 |

## Arquivos de configuração / suporte

| Arquivo do legado | Unit correspondente | Cobertura | Observação |
|-------------------|---------------------|-----------|------------|
| `go.mod` | — | n/a | Dependências documentadas em `dependencies.md` |
| `go.sum` | — | n/a | Lockfile |
| `README.md` | — | n/a | Documentação textual; entry points em `inventory.md` |

## Artefatos referenciados mas ausentes no recorte

| Item | Status | Observação |
|------|--------|------------|
| `test_conn.go` | 🔴 ausente | Citado no `README.md`, não presente no recorte (ver `surface.json.gaps`) |
| Camada de cálculo (motor `TaxRequest`→`TaxResponse`) | 🔴 externa | Vive no monorepo de origem; **não** documentada aqui por decisão (A1) — será documentada nos serviços consumidores |
| Escrita de `tax_calculation_logs` | 🔴 externa/futura | Tabela existe no DDL (A3); persistência é **evolução futura** desta lib (A2) |
| DDL / migrations | 🟢 resolvido | Schema fornecido na resposta A3; refletido em `erd-complete.md` e `models/contracts.md` |

## Resumo de cobertura

- **Arquivos de código Go:** 4 de 4 mapeados a uma unit → **100%** 🟢
- **Total de arquivos do recorte:** 8 (4 Go + `go.mod`, `go.sum`, `README.md` + 1 ausente `test_conn.go`)
- **Cobertura de código-fonte executável:** 100% (toda função pública e struct mapeada a uma unit)
- **Fronteiras externas documentadas como tais:** camada de cálculo (A1), auditoria futura (A2), autorização/multi-tenant (D8)
