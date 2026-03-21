# C4 — Nível 4 (Código) — batch-geolocalidade

## Objetivo

Detalhar a visão de **código** (classes e relacionamentos) do fluxo de importação DTB/IBGE.

O recorte é o pipeline típico de Spring Batch:

CSV → `FlatFileItemReader<DTO>` → `ItemProcessor<DTO, Entity>` → `RepositoryItemWriter<Entity>` → `JpaRepository.save(...)`.

## Fluxo de execução (alto nível)

- Job: `importacaoGeolocalidadeJob`
  1. `stepMunicipio`: importa municípios e hierarquia superior (UF, Regiões)
  2. `stepDistrito`: importa distritos (referencia município por `getReferenceById`)
  3. `stepSubdistrito`: importa subdistritos (referencia distrito por `getReferenceById`)

- Runner opcional:
  - `LoadTestRunner` dispara o Job quando `app.loadtest.enabled=true`.

## Diagrama (Classes principais)

```mermaid
classDiagram
  direction LR

  class BatchConfig {
    - String importPath
    - String municipiosFile
    - String distritosFile
    - String subdistritosFile
    + Job importacaoGeolocalidadeJob(Step, Step, Step)
    + Step stepMunicipio(ItemReader, ItemProcessor, ItemWriter)
    + Step stepDistrito(FlatFileItemReader, DistritoProcessor, RepositoryItemWriter)
    + Step stepSubdistrito(FlatFileItemReader, SubdistritoProcessor, RepositoryItemWriter)
    + FlatFileItemReader~MunicipioCsvDTO~ municipioReader()
    + FlatFileItemReader~DistritoCsvDTO~ distritoReader()
    + FlatFileItemReader~SubdistritoCsvDTO~ subdistritoReader()
    + RepositoryItemWriter~Municipio~ municipioWriter(MunicipioRepository)
    + RepositoryItemWriter~Distrito~ distritoWriter(DistritoRepository)
    + RepositoryItemWriter~Subdistrito~ subdistritoWriter(SubdistritoRepository)
  }

  class LoadTestRunner {
    - JobLauncher jobLauncher
    - Job job
    + run(ApplicationArguments args) void
  }

  class MunicipioProcessor {
    + process(MunicipioCsvDTO item) Municipio
  }

  class DistritoProcessor {
    - MunicipioRepository municipioRepository
    + process(DistritoCsvDTO item) Distrito
  }

  class SubdistritoProcessor {
    - DistritoRepository distritoRepository
    + process(SubdistritoCsvDTO item) Subdistrito
  }

  class MunicipioRepository {
    <<interface>>
    + save(Municipio) Municipio
    + getReferenceById(String id) Municipio
  }

  class DistritoRepository {
    <<interface>>
    + save(Distrito) Distrito
    + getReferenceById(String id) Distrito
  }

  class SubdistritoRepository {
    <<interface>>
    + save(Subdistrito) Subdistrito
  }

  class Uf
  class RegiaoIntermediaria
  class RegiaoImediata
  class Municipio
  class Distrito
  class Subdistrito

  class MunicipioCsvDTO
  class DistritoCsvDTO
  class SubdistritoCsvDTO

  %% Orquestração
  LoadTestRunner --> BatchConfig
  BatchConfig --> MunicipioProcessor
  BatchConfig --> DistritoProcessor
  BatchConfig --> SubdistritoProcessor

  %% Processors -> Repos (referências)
  DistritoProcessor --> MunicipioRepository
  SubdistritoProcessor --> DistritoRepository

  %% Entidades e hierarquia
  RegiaoIntermediaria --> Uf
  RegiaoImediata --> RegiaoIntermediaria
  Municipio --> RegiaoImediata
  Distrito --> Municipio
  Subdistrito --> Distrito

  %% DTOs
  MunicipioProcessor --> MunicipioCsvDTO
  DistritoProcessor --> DistritoCsvDTO
  SubdistritoProcessor --> SubdistritoCsvDTO
```

## Notas de implementação (observadas no código)

- Os readers usam `linesToSkip(7)` e `encoding("UTF-8")`.
- Os tokenizers estão com `strict=false` para tolerar colunas extras.
- `MunicipioProcessor` gera/associa `Uf`, `RegiaoIntermediaria`, `RegiaoImediata` e `Municipio` num único processamento.
- `DistritoProcessor` e `SubdistritoProcessor` usam `getReferenceById(...)` para manter integridade de FKs sem provocar inserts indevidos.

