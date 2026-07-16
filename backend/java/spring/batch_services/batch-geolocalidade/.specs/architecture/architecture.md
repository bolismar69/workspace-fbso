---
title: "Arquitetura — batch-geolocalidade"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["architecture", "spring-batch", "c4", "java"]
---

# Arquitetura — batch-geolocalidade

## Resumo Executivo

**batch-geolocalidade** é um serviço Spring Batch projetado para executar como um **job efêmero** (Kubernetes Job / CronJob) que importa dados da Divisão Territorial Brasileira (DTB) do IBGE a partir de arquivos CSV e os persiste em PostgreSQL. Após a execução, o processo se encerra com exit code determinístico.

O serviço é consumido tipicamente pelo microserviço `ms-geolocalidade`, que consulta as tabelas populadas por este batch para enriquecer respostas de API com dados oficiais do IBGE.

## Padrão Arquitetural

**Batch Processing (Chunk-oriented)** — O Spring Boot atua como container do Spring Batch, que orquestra a leitura, processamento e escrita em chunks transacionais de 100 registros. O padrão segue a arquitetura clássica do Spring Batch:

```
Reader → Processor → Writer
  ↓          ↓          ↓
 CSV    Transformação   JPA
```

**Justificativa:** Spring Batch foi escolhido por fornecer gerenciamento transacional, restartabilidade, métricas e separação clara de responsabilidades (Reader/Processor/Writer), essenciais para processamento confiável de arquivos CSV com milhares de registros.

## Stack Tecnológica

| Componente | Tecnologia | Versão |
|---|---|---|
| Linguagem | Java | 21 |
| Framework | Spring Boot | 3.5.12 |
| Batch | Spring Batch | 5.x (transitivo do Boot) |
| Persistência | Spring Data JPA + Hibernate | 6.x |
| Banco Primário | PostgreSQL | 16 |
| Banco de Testes | H2 | (in-memory) |
| Pool de Conexões | HikariCP | (transitivo do Boot) |
| Build | Maven Wrapper | 3.x |
| Encoding | UTF-8 | — |

## Estrutura de Diretórios

```
batch-geolocalidade/
├── .specs/                              ← Documentação técnica (esta pasta)
├── src/
│   ├── main/
│   │   ├── java/br/com/fbso/geolocalidade/
│   │   │   ├── SpringBatchGeolocalidadeApplication.java  ← Entry point
│   │   │   ├── config/
│   │   │   │   └── BatchConfig.java        ← Job, Steps, Readers, Writers
│   │   │   ├── dto/
│   │   │   │   ├── MunicipioCsvDTO.java     ← Record: 9 campos CSV
│   │   │   │   ├── DistritoCsvDTO.java      ← Record: 12 campos CSV
│   │   │   │   └── SubdistritoCsvDTO.java   ← Record: 15 campos CSV
│   │   │   ├── entity/
│   │   │   │   ├── Uf.java                  ← JPA: uf (PK string 2)
│   │   │   │   ├── RegiaoIntermediaria.java ← JPA: regiao_intermediaria (PK string 4)
│   │   │   │   ├── RegiaoImediata.java      ← JPA: regiao_imediata (PK string 6)
│   │   │   │   ├── Municipio.java           ← JPA: municipio (PK string 7)
│   │   │   │   ├── Distrito.java            ← JPA: distrito (PK string 9)
│   │   │   │   └── Subdistrito.java         ← JPA: subdistrito (PK string 11)
│   │   │   ├── load/
│   │   │   │   └── LoadTestRunner.java      ← Runner condicional (exit code)
│   │   │   ├── processor/
│   │   │   │   ├── MunicipioProcessor.java  ← CSV→JPA + cache UF/Região
│   │   │   │   ├── DistritoProcessor.java   ← CSV→JPA com getReferenceById
│   │   │   │   └── SubdistritoProcessor.java← CSV→JPA com getReferenceById
│   │   │   └── repository/
│   │   │       ├── UfRepository.java
│   │   │       ├── MunicipioRepository.java
│   │   │       ├── DistritoRepository.java
│   │   │       └── SubdistritoRepository.java
│   │   └── resources/
│   │       ├── application.yaml             ← Configuração principal
│   │       └── db/
│   │           └── init-postgres.sql         ← Criação idempotente de schemas
│   └── test/
│       ├── java/.../SpringBatchGeolocalidadeApplicationTests.java
│       └── resources/
│           └── application.yaml             ← Config H2 para testes
├── pom.xml                                  ← Maven (Spring Boot 3.5.12 parent)
└── README.md                                ← Documentação de uso
```

## Princípios de Design

1. **Separation of Concerns**: Reader, Processor e Writer são componentes independentes, facilitando teste e manutenção.
2. **Natural Keys**: IDs das entidades são os próprios códigos IBGE (sem surrogate keys), simplificando integração com sistemas externos.
3. **Chunk-oriented Processing**: Processamento em lotes de 100 registros por transação, balanceando performance e atomicidade.
4. **Conditional Runner**: `LoadTestRunner` é ativado apenas com flag `--app.loadtest.enabled=true`, permitindo que o mesmo artefato seja usado em diferentes cenários.
5. **Idempotent Schema Creation**: `init-postgres.sql` usa `CREATE SCHEMA IF NOT EXISTS`, seguro para múltiplas execuções.
6. **Dual Schema Pattern**: Separação entre schema de framework (`spring_batch`) e schema de negócio (`localidade`).

## Cross-cutting Concerns

| Concern | Implementação |
|---|---|
| **Logging** | SLF4J via Spring Boot. Batch: INFO, Hibernate SQL: ERROR |
| **Transações** | Gerenciadas pelo Spring Batch via `PlatformTransactionManager` |
| **Métricas** | Spring Batch built-in (readCount, writeCount, skipCount). LoadTestRunner loga counts e duration |
| **Configuração** | Externalizada via `application.yaml` + variáveis de ambiente |
| **Autenticação** | N/A — serviço batch headless |
| **Observabilidade** | Logs estruturados (SLF4J). Spring Actuator planejado para futuro |
