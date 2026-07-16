---
title: "Dicionário de Dados — batch-geolocalidade"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["architecture", "data-dictionary", "postgresql"]
---

# Dicionário de Dados — batch-geolocalidade

## Schema `localidade`

### Tabela: `uf`

**Propósito:** Armazenar as 27 Unidades Federativas do Brasil.

| Coluna | Tipo | Nullable | Default | Descrição | Regra |
|---|---|---|---|---|---|
| `id` | `VARCHAR(2)` | NOT NULL | — | Código IBGE de 2 dígitos da UF | PK |
| `sigla` | `VARCHAR(5)` | NOT NULL | — | Sigla oficial (ex: SP, RJ, MG) | Derivado via `MunicipioProcessor.UF_SIGLAS` |
| `nome` | `VARCHAR(50)` | NOT NULL | — | Nome completo do estado | Vem do CSV do IBGE |

### Tabela: `regiao_intermediaria`

**Propósito:** Armazenar as Regiões Intermediárias definidas pelo IBGE (~133 registros).

| Coluna | Tipo | Nullable | Default | Descrição | Regra |
|---|---|---|---|---|---|
| `id` | `VARCHAR(4)` | NOT NULL | — | Código IBGE de 4 dígitos | PK |
| `nome` | `VARCHAR(100)` | NOT NULL | — | Nome da região intermediária | Vem do CSV |
| `uf_id` | `VARCHAR(2)` | NOT NULL | — | FK para `uf.id` | FK → `uf(id)`, CascadeType.MERGE |

### Tabela: `regiao_imediata`

**Propósito:** Armazenar as Regiões Imediatas definidas pelo IBGE (~510 registros).

| Coluna | Tipo | Nullable | Default | Descrição | Regra |
|---|---|---|---|---|---|
| `id` | `VARCHAR(6)` | NOT NULL | — | Código IBGE de 6 dígitos | PK |
| `nome` | `VARCHAR(100)` | NOT NULL | — | Nome da região imediata | Vem do CSV |
| `regiao_intermediaria_id` | `VARCHAR(4)` | NOT NULL | — | FK para `regiao_intermediaria.id` | FK, CascadeType.MERGE |

### Tabela: `municipio`

**Propósito:** Armazenar os 5.570 municípios brasileiros.

| Coluna | Tipo | Nullable | Default | Descrição | Regra |
|---|---|---|---|---|---|
| `id` | `VARCHAR(7)` | NOT NULL | — | Código IBGE completo de 7 dígitos | PK |
| `codigo` | `VARCHAR(5)` | NOT NULL | — | Código curto (sem DV) | Vem do CSV |
| `nome` | `VARCHAR(100)` | NOT NULL | — | Nome do município | Vem do CSV |
| `regiao_imediata_id` | `VARCHAR(6)` | NOT NULL | — | FK para `regiao_imediata.id` | FK, CascadeType.MERGE |

**Legacy Getters:**
- `getCodigoIbge7()` → retorna `id`
- `getNomeMunicipio()` → retorna `nome`
- `getUfSigla()` → navega `regiaoImediata → regiaoIntermediaria → uf.sigla`

### Tabela: `distrito`

**Propósito:** Armazenar os ~10.407 distritos brasileiros.

| Coluna | Tipo | Nullable | Default | Descrição | Regra |
|---|---|---|---|---|---|
| `id` | `VARCHAR(9)` | NOT NULL | — | Código IBGE completo de 9 dígitos | PK |
| `codigo` | `VARCHAR(2)` | NOT NULL | — | Código curto do distrito | Vem do CSV |
| `nome` | `VARCHAR(100)` | NOT NULL | — | Nome do distrito | Vem do CSV |
| `municipio_id` | `VARCHAR(7)` | NOT NULL | — | FK para `municipio.id` | FK, CascadeType.MERGE |

### Tabela: `subdistrito`

**Propósito:** Armazenar os ~684 subdistritos brasileiros.

| Coluna | Tipo | Nullable | Default | Descrição | Regra |
|---|---|---|---|---|---|
| `id` | `VARCHAR(11)` | NOT NULL | — | Código IBGE completo de 11 dígitos | PK |
| `codigo` | `VARCHAR(2)` | NOT NULL | — | Código curto do subdistrito | Vem do CSV |
| `nome` | `VARCHAR(100)` | NOT NULL | — | Nome do subdistrito | Vem do CSV |
| `distrito_id` | `VARCHAR(9)` | NOT NULL | — | FK para `distrito.id` | FK, CascadeType.MERGE |

**Legacy Getters:**
- `getCodigoSubdistrito11()` → retorna `id`
- `getNomeSubdistrito()` → retorna `nome`

## Schema `spring_batch`

Tabelas gerenciadas automaticamente pelo Spring Batch (`spring.batch.jdbc.initialize-schema=always`).

| Tabela | Propósito |
|---|---|
| `BATCH_JOB_INSTANCE` | Registro de cada execução de Job |
| `BATCH_JOB_EXECUTION` | Dados de execução (status, timestamps) |
| `BATCH_JOB_EXECUTION_PARAMS` | Parâmetros usados na execução |
| `BATCH_STEP_EXECUTION` | Dados de cada Step (readCount, writeCount, status) |
| `BATCH_JOB_EXECUTION_CONTEXT` | Contexto serializado da execução |
| `BATCH_STEP_EXECUTION_CONTEXT` | Contexto serializado do step |
| `BATCH_STEP_EXECUTION_SEQ` | Sequências para IDs |
| `BATCH_JOB_EXECUTION_SEQ` | Sequências para IDs |
| `BATCH_JOB_SEQ` | Sequências para IDs |

## Tipos de Dados JPA → PostgreSQL

| Java / JPA | PostgreSQL |
|---|---|
| `String` com `@Column(length=N)` | `VARCHAR(N)` |
| Os IDs são `String` em Java mas armazenados como `VARCHAR` no banco | — |

## Regras de Negócio Implementadas via Constraints/Triggers

| Regra | Implementação |
|---|---|
| IDs únicos por tabela | `@Id` → PK constraint |
| FKs obrigatórias (não nulas) | `@ManyToOne(optional=false)` + `@JoinColumn(nullable=false)` |
| Nomes e siglas obrigatórios | `@Column(nullable=false)` |
| Cascade MERGE (não PERSIST) | `cascade = CascadeType.MERGE` — seguro para batch |
