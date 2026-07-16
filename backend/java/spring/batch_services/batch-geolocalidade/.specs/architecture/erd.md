---
title: "ERD — Diagrama Entidade-Relacionamento — batch-geolocalidade"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["architecture", "erd", "database", "postgresql"]
---

# Diagrama Entidade-Relacionamento — batch-geolocalidade

## Esquemas

O banco PostgreSQL utiliza **dois schemas**:

| Schema | Propósito | Tabelas |
|---|---|---|
| `spring_batch` | Metadata do Spring Batch | `BATCH_JOB_INSTANCE`, `BATCH_JOB_EXECUTION`, `BATCH_STEP_EXECUTION`, etc. |
| `localidade` | Tabelas de negócio (DTB/IBGE) | `uf`, `regiao_intermediaria`, `regiao_imediata`, `municipio`, `distrito`, `subdistrito` |

## ERD — Schema `localidade`

```mermaid
erDiagram
  UF {
    varchar id PK "Código IBGE 2 dígitos (ex: 35)"
    varchar sigla "Sigla oficial (ex: SP)"
    varchar nome "Nome do estado (ex: São Paulo)"
  }

  REGIAO_INTERMEDIARIA {
    varchar id PK "Código IBGE 4 dígitos (ex: 3501)"
    varchar nome "Nome da região intermediária"
    varchar uf_id FK "FK → uf.id"
  }

  REGIAO_IMEDIATA {
    varchar id PK "Código IBGE 6 dígitos (ex: 350001)"
    varchar nome "Nome da região imediata"
    varchar regiao_intermediaria_id FK "FK → regiao_intermediaria.id"
  }

  MUNICIPIO {
    varchar id PK "Código IBGE 7 dígitos (ex: 3550308)"
    varchar codigo "Código curto 5 dígitos"
    varchar nome "Nome do município"
    varchar regiao_imediata_id FK "FK → regiao_imediata.id"
  }

  DISTRITO {
    varchar id PK "Código IBGE 9 dígitos (ex: 355030805)"
    varchar codigo "Código curto 2 dígitos"
    varchar nome "Nome do distrito"
    varchar municipio_id FK "FK → municipio.id"
  }

  SUBDISTRITO {
    varchar id PK "Código IBGE 11 dígitos (ex: 35503080501)"
    varchar codigo "Código curto 2 dígitos"
    varchar nome "Nome do subdistrito"
    varchar distrito_id FK "FK → distrito.id"
  }

  UF ||--o{ REGIAO_INTERMEDIARIA : "1:N — possui"
  REGIAO_INTERMEDIARIA ||--o{ REGIAO_IMEDIATA : "1:N — possui"
  REGIAO_IMEDIATA ||--o{ MUNICIPIO : "1:N — possui"
  MUNICIPIO ||--o{ DISTRITO : "1:N — possui"
  DISTRITO ||--o{ SUBDISTRITO : "1:N — possui"
```

## Cardinalidade e Regras

| Relacionamento | Cardinalidade | FK Column | Nullable | On Delete |
|---|---|---|---|---|
| UF → RegiaoIntermediaria | 1:N | `uf_id` | NOT NULL | — |
| RegiaoIntermediaria → RegiaoImediata | 1:N | `regiao_intermediaria_id` | NOT NULL | — |
| RegiaoImediata → Municipio | 1:N | `regiao_imediata_id` | NOT NULL | — |
| Municipio → Distrito | 1:N | `municipio_id` | NOT NULL | — |
| Distrito → Subdistrito | 1:N | `distrito_id` | NOT NULL | — |

## Volumetria Esperada (Brasil, DTB 2024)

| Tabela | Registros (~) |
|---|---|
| `uf` | 27 |
| `regiao_intermediaria` | 133 |
| `regiao_imediata` | 510 |
| `municipio` | 5.570 |
| `distrito` | 10.407 |
| `subdistrito` | 684 |

## Índices

Índices são criados automaticamente pelo Hibernate/JPA via `@Id` nas chaves primárias. Índices adicionais em colunas FK não são criados automaticamente — recomenda-se adicionar manualmente em ambiente de produção:

```sql
CREATE INDEX IF NOT EXISTS idx_regiao_intermediaria_uf ON localidade.regiao_intermediaria(uf_id);
CREATE INDEX IF NOT EXISTS idx_regiao_imediata_inter ON localidade.regiao_imediata(regiao_intermediaria_id);
CREATE INDEX IF NOT EXISTS idx_municipio_imediata ON localidade.municipio(regiao_imediata_id);
CREATE INDEX IF NOT EXISTS idx_distrito_municipio ON localidade.distrito(municipio_id);
CREATE INDEX IF NOT EXISTS idx_subdistrito_distrito ON localidade.subdistrito(distrito_id);
```
