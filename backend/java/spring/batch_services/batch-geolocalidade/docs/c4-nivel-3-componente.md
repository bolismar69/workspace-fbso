# C4 — Nível 3 (Componente) — batch-geolocalidade

## Objetivo

Descrever os **componentes** internos do contêiner "Spring Boot + Spring Batch" do `batch-geolocalidade`.

## Componentes principais

### Orquestração do Job

- `BatchConfig` (`br.com.fbso.geolocalidade.config`)
  - Define o Job `importacaoGeolocalidadeJob`.
  - Define 3 steps:
    - `stepMunicipio` (Municípios + hierarquia superior)
    - `stepDistrito`
    - `stepSubdistrito`

### Leitura de CSV (Readers)

- `FlatFileItemReader<MunicipioCsvDTO>` (`municipioReader`)
- `FlatFileItemReader<DistritoCsvDTO>` (`distritoReader`)
- `FlatFileItemReader<SubdistritoCsvDTO>` (`subdistritoReader`)

Características observadas:

- CSV delimitado por vírgula (`DelimitedLineTokenizer`) com `strict=false`.
- `linesToSkip(7)` para pular metadados do IBGE.
- `encoding("UTF-8")`.

### Processamento (Processors)

- `MunicipioProcessor`
  - Monta `Uf`, `RegiaoIntermediaria`, `RegiaoImediata` e `Municipio` a partir do DTO.
  - Mantém caches em memória durante o Job para evitar recriar instâncias repetidas.

- `DistritoProcessor`
  - Cria `Distrito` usando `MunicipioRepository.getReferenceById(...)`.

- `SubdistritoProcessor`
  - Cria `Subdistrito` usando `DistritoRepository.getReferenceById(...)`.

### Persistência (Writers + Repositories)

- Writers: `RepositoryItemWriter` com `save(...)`:
  - `MunicipioRepository`
  - `DistritoRepository`
  - `SubdistritoRepository`

- Repositórios adicionais:
  - `UfRepository` (presente no projeto; pode ser usado em evoluções/consultas)

### Runner opcional (carga local)

- `LoadTestRunner`
  - Componente condicional por `app.loadtest.enabled=true`.
  - Dispara o Job via `JobLauncher` e encerra o processo com exit code (0/1).

## Diagrama (Componentes)

```mermaid
flowchart TB
  %% C4-L3: Component diagram (representado em Mermaid)

  subgraph app["Container: Spring Boot + Spring Batch"]
    BatchConfig["BatchConfig\n(Job/Steps/Readers/Writers)"]

    subgraph readers["Readers (CSV)"]
      MunicipioReader["FlatFileItemReader<MunicipioCsvDTO>"]
      DistritoReader["FlatFileItemReader<DistritoCsvDTO>"]
      SubdistritoReader["FlatFileItemReader<SubdistritoCsvDTO>"]
    end

    subgraph processors["Processors"]
      MunicipioProcessor["MunicipioProcessor"]
      DistritoProcessor["DistritoProcessor"]
      SubdistritoProcessor["SubdistritoProcessor"]
    end

    subgraph writers["Writers (JPA)"]
      MunicipioWriter["RepositoryItemWriter<Municipio>"]
      DistritoWriter["RepositoryItemWriter<Distrito>"]
      SubdistritoWriter["RepositoryItemWriter<Subdistrito>"]
    end

    subgraph repos["Repositories"]
      MunicipioRepository["MunicipioRepository"]
      DistritoRepository["DistritoRepository"]
      SubdistritoRepository["SubdistritoRepository"]
      UfRepository["UfRepository"]
    end

    LoadTestRunner["LoadTestRunner\n(opcional)"]
  end

  fs[("File System / Volume\napp.import.path")]
  db[("PostgreSQL\n(spring_batch + localidade)")]
  scheduler["Operador/Scheduler"]

  %% Trigger
  scheduler --> LoadTestRunner
  scheduler -->|"start app"| app
  LoadTestRunner -->|"JobLauncher.run(job)"| BatchConfig

  %% Step 1 pipeline
  BatchConfig --> MunicipioReader
  MunicipioReader -->|"DTO"| MunicipioProcessor
  MunicipioProcessor -->|"Entity"| MunicipioWriter
  MunicipioWriter --> MunicipioRepository

  %% Step 2 pipeline
  BatchConfig --> DistritoReader
  DistritoReader --> DistritoProcessor
  DistritoProcessor --> DistritoWriter
  DistritoWriter --> DistritoRepository

  %% Step 3 pipeline
  BatchConfig --> SubdistritoReader
  SubdistritoReader --> SubdistritoProcessor
  SubdistritoProcessor --> SubdistritoWriter
  SubdistritoWriter --> SubdistritoRepository

  %% IO
  MunicipioReader --> fs
  DistritoReader --> fs
  SubdistritoReader --> fs

  MunicipioRepository --> db
  DistritoRepository --> db
  SubdistritoRepository --> db
  UfRepository --> db
```

