---
title: "Code Analysis — batch-geolocalidade"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["engineering", "code-analysis", "batch"]
---

# Análise de Fluxo de Código — batch-geolocalidade

## Fluxo Principal de Execução

### 1. Bootstrap da Aplicação

```
SpringApplication.run(SpringBatchGeolocalidadeApplication.class, args)
  ├── Carrega application.yaml
  ├── Executa init-postgres.sql (CREATE SCHEMA IF NOT EXISTS)
  ├── Hibernate DDL: cria/atualiza tabelas no schema localidade
  ├── Spring Batch: cria tabelas BATCH_* no schema spring_batch
  └── ApplicationRunners são executados
```

### 2. LoadTestRunner (se --app.loadtest.enabled=true)

```
LoadTestRunner.run(args)
  ├── logImportFile("municipios")  → valida CSV, loga exists=true/false
  ├── logImportFile("distritos")   → idem
  ├── logImportFile("subdistritos")→ idem
  ├── jobLauncher.run(job, params) → dispara Job
  │   ├── Step 1: stepMunicipio
  │   ├── Step 2: stepDistrito
  │   └── Step 3: stepSubdistrito
  ├── Loga métricas: status, duration, counts
  └── System.exit(0 ou 1)
```

### 3. Detalhe do stepMunicipio (Chunk Processing)

```
Para cada chunk de 100 linhas:
  FlatFileItemReader.read()
    ├── Abre FileSystemResource (UTF-8)
    ├── Pula 7 linhas (metadados IBGE)
    └── DelimitedLineTokenizer parseia 9 campos → MunicipioCsvDTO

  MunicipioProcessor.process(dto)
    ├── ufCache.computeIfAbsent(id) → Uf(id, sigla, nome)
    ├── interCache.computeIfAbsent(id) → RegiaoIntermediaria(id, nome, uf)
    ├── imedCache.computeIfAbsent(id) → RegiaoImediata(id, nome, inter)
    └── return new Municipio(id, codigo, nome, imediata)

  RepositoryItemWriter.write(chunk)
    └── municipioRepository.saveAll(chunk)
        ├── CascadeType.MERGE: persiste UF/RegInter/RegImed se necessário
        └── Hibernate flush + commit
```

### 4. Detalhe do stepDistrito

```
Para cada chunk de 100 linhas:
  FlatFileItemReader.read()
    └── DelimitedLineTokenizer parseia 12 campos → DistritoCsvDTO

  DistritoProcessor.process(dto)
    ├── municipioRepository.getReferenceById(municipioId)
    │   └── Retorna proxy JPA (sem SELECT!)
    └── return new Distrito(id, codigo, nome, municipioRef)

  RepositoryItemWriter.write(chunk)
    └── distritoRepository.saveAll(chunk)
```

## Cache Strategy

| Processor | Cache | Chave | Tamanho Máximo |
|---|---|---|---|
| MunicipioProcessor | `ufCache` (HashMap) | Código UF (2 dígitos) | 27 |
| MunicipioProcessor | `interCache` (HashMap) | Código Região Intermediária (4 dígitos) | 133 |
| MunicipioProcessor | `imedCache` (HashMap) | Código Região Imediata (6 dígitos) | 510 |
| DistritoProcessor | Nenhum | — | — |
| SubdistritoProcessor | Nenhum | — | — |

**Nota:** Os caches do `MunicipioProcessor` vivem apenas durante a execução de um Job. Eles são recriados a cada nova instância do processor (a cada Job).

## Tratamento de Erros

| Cenário | Comportamento |
|---|---|
| CSV não encontrado | `FlatFileItemReader` lança `ItemStreamException`: "Input resource must exist (reader is in 'strict' mode)" |
| Linha com colunas extras | Tokenizer com `strict=false` → tolera e ignora colunas extras |
| FK inválida (Distrito/Subdistrito) | `getReferenceById()` lança `EntityNotFoundException` se a referência não existir |
| Encoding incorreto | Caracteres acentuados aparecem quebrados (ex: `RondÃ´nia`) — resolvido com UTF-8 |
| Falha de transação | Spring Batch faz rollback do chunk e registra falha no `BATCH_STEP_EXECUTION` |
