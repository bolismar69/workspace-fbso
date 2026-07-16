---
title: "C4 — Containers — batch-geolocalidade"
level: "Containers"
---

# C4 — Nível 2: Containers

## Diagrama

```mermaid
C4Container
  title Diagrama de Containers — batch-geolocalidade

  Container(jvm, "JVM (Spring Boot)", "Java 21, Spring Boot 3.5.12", "Orquestra a execução do Spring Batch e gerencia o ciclo de vida da aplicação")

  Container(batch_engine, "Spring Batch Engine", "Spring Batch 5.x", "Gerencia Jobs, Steps, transações, restartabilidade e métricas de execução")

  Container(jpa_layer, "JPA / Hibernate Layer", "Spring Data JPA + Hibernate 6.x", "Mapeia entidades para tabelas no schema localidade e gerencia persistência")

  ContainerDb(postgres, "PostgreSQL Database", "PostgreSQL 16", "Armazena tabelas em dois schemas: spring_batch (metadata) e localidade (negócio)")

  System_Ext(filesystem, "File System", "Volume montado (local ou K8s)", "Diretório com arquivos CSV do IBGE")

  Rel(jvm, batch_engine, "Delega execução do Job", "Spring Context")
  Rel(batch_engine, filesystem, "Lê CSVs via FlatFileItemReader", "File I/O (UTF-8)")
  Rel(batch_engine, jpa_layer, "Persiste via RepositoryItemWriter", "Spring Data JPA")
  Rel(jpa_layer, postgres, "JDBC (HikariCP pool)", "JDBC:5432")
```

## Elementos

| Nome | Tipo | Responsabilidade | Tecnologia |
|---|---|---|---|
| JVM (Spring Boot) | Container | Application context, injeção de dependências, configuração, ciclo de vida | Java 21, Spring Boot 3.5.12 |
| Spring Batch Engine | Container | Orquestração de Jobs/Steps, gerenciamento transacional, chunk processing, métricas | Spring Batch 5.x |
| JPA / Hibernate Layer | Container | Mapeamento objeto-relacional, persistência de entidades, proxies lazy | Spring Data JPA, Hibernate 6.x |
| PostgreSQL Database | Container (Database) | Armazenamento durável em dois schemas | PostgreSQL 16 |
| File System | External System | Fonte de dados — arquivos CSV do IBGE | Volume (local ou K8s PV) |

## Fluxos Principais

### Fluxo: Chunk Processing (exemplo: Step Municípios)

1. `FlatFileItemReader` abre `DTB_Municipios.csv` do File System
2. Lê 100 linhas (1 chunk), pula 7 linhas de metadados
3. Para cada linha, `MunicipioProcessor` transforma `MunicipioCsvDTO` → `Municipio` (com cache de UF/Regiões)
4. `RepositoryItemWriter` persiste as 100 entidades via `MunicipioRepository.save()`
5. Spring Batch commita a transação
6. Repete até EOF

### Fluxo: Dual Schema

1. `init-postgres.sql` cria schemas `spring_batch` e `localidade` (idempotente)
2. Conexão JDBC usa `currentSchema=spring_batch` → tabelas `BATCH_*` criadas aqui
3. Hibernate usa `default_schema=localidade` → tabelas de negócio (`uf`, `municipio`, etc.) criadas aqui
