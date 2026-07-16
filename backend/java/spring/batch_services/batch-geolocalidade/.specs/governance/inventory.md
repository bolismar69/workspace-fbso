---
title: "Inventário do Projeto — batch-geolocalidade"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["governance", "inventory", "reference"]
---

# Inventário do Projeto — batch-geolocalidade

## Módulos e Cobertura

| Módulo | Arquivos Java | Cobertura de Testes |
|---|---|---|
| `entity` | 6 (Uf, RegiaoIntermediaria, RegiaoImediata, Municipio, Distrito, Subdistrito) | — |
| `dto` | 3 (MunicipioCsvDTO, DistritoCsvDTO, SubdistritoCsvDTO) | — |
| `config` | 1 (BatchConfig) | — |
| `processor` | 3 (MunicipioProcessor, DistritoProcessor, SubdistritoProcessor) | — |
| `repository` | 4 (UfRepository, MunicipioRepository, DistritoRepository, SubdistritoRepository) | — |
| `load` | 1 (LoadTestRunner) | — |
| **Total** | **18 arquivos Java** | 1 teste (`SpringBatchGeolocalidadeApplicationTests`) |

## Estrutura de Diretórios e Arquivos

```
batch-geolocalidade/
├── .specs/                              ← 18 arquivos (criados 2026-07-08)
├── src/main/java/br/com/fbso/geolocalidade/
│   ├── SpringBatchGeolocalidadeApplication.java
│   ├── config/BatchConfig.java
│   ├── dto/
│   │   ├── MunicipioCsvDTO.java
│   │   ├── DistritoCsvDTO.java
│   │   └── SubdistritoCsvDTO.java
│   ├── entity/
│   │   ├── Uf.java
│   │   ├── RegiaoIntermediaria.java
│   │   ├── RegiaoImediata.java
│   │   ├── Municipio.java
│   │   ├── Distrito.java
│   │   └── Subdistrito.java
│   ├── load/LoadTestRunner.java
│   ├── processor/
│   │   ├── MunicipioProcessor.java
│   │   ├── DistritoProcessor.java
│   │   └── SubdistritoProcessor.java
│   └── repository/
│       ├── UfRepository.java
│       ├── MunicipioRepository.java
│       ├── DistritoRepository.java
│       └── SubdistritoRepository.java
├── src/main/resources/
│   ├── application.yaml
│   └── db/init-postgres.sql
├── src/test/java/.../
│   └── SpringBatchGeolocalidadeApplicationTests.java
├── src/test/resources/
│   └── application.yaml
├── pom.xml
└── README.md
```

## Dependências Externas

| Dependência | Versão | CVE Status |
|---|---|---|
| Spring Boot | 3.5.12 | — |
| PostgreSQL Driver | (transitivo) | — |
| H2 | (transitivo) | — |
