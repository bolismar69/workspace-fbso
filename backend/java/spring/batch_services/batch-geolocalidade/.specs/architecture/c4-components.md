---
title: "C4 — Componentes — batch-geolocalidade"
level: "Components"
---

# C4 — Nível 3: Componentes

## Diagrama

```mermaid
C4Component
  title Diagrama de Componentes — batch-geolocalidade

  Container_Boundary(batch_app, "Spring Batch Application") {

    Component(load_runner, "LoadTestRunner", "ApplicationRunner", "Valida CSVs, dispara Job, loga métricas e encerra processo com exit code")

    Component(job_config, "BatchConfig", "@Configuration", "Define Job, 3 Steps, Readers e Writers")

    Component(mun_proc, "MunicipioProcessor", "ItemProcessor", "Transforma CSV em entidades; cache de UF/Regiões (HashMap)")

    Component(dist_proc, "DistritoProcessor", "ItemProcessor", "Transforma CSV em Distrito; getReferenceById para FK")

    Component(sub_proc, "SubdistritoProcessor", "ItemProcessor", "Transforma CSV em Subdistrito; getReferenceById para FK")

    Component(mun_repo, "MunicipioRepository", "JpaRepository", "Persiste e consulta Municípios")
    Component(dist_repo, "DistritoRepository", "JpaRepository", "Persiste e consulta Distritos")
    Component(sub_repo, "SubdistritoRepository", "JpaRepository", "Persiste e consulta Subdistritos")
    Component(uf_repo, "UfRepository", "JpaRepository", "Persiste e consulta UFs")
  }

  ContainerDb(postgres, "PostgreSQL", "PostgreSQL 16", "Schemas: spring_batch + localidade")

  System_Ext(fs, "File System", "CSVs IBGE")

  Rel(load_runner, job_config, "Dispara importacaoGeolocalidadeJob", "JobLauncher.run()")
  Rel(job_config, mun_proc, "Usa no stepMunicipio", "ItemProcessor")
  Rel(job_config, dist_proc, "Usa no stepDistrito", "ItemProcessor")
  Rel(job_config, sub_proc, "Usa no stepSubdistrito", "ItemProcessor")
  Rel(job_config, fs, "Lê CSVs via FlatFileItemReader", "File I/O")
  Rel(mun_proc, mun_repo, "Cache de UF/Regiões cria entidades", "JPA cascade")
  Rel(dist_proc, dist_repo, "getReferenceById", "JPA proxy")
  Rel(sub_proc, sub_repo, "getReferenceById", "JPA proxy")

  Rel(job_config, postgres, "Persiste via RepositoryItemWriter", "JDBC")
  Rel(mun_repo, postgres, "save()", "JDBC")
  Rel(dist_repo, postgres, "save()", "JDBC")
  Rel(sub_repo, postgres, "save()", "JDBC")
  Rel(uf_repo, postgres, "save() (cascade)", "JDBC")
```

## Elementos

| Nome | Tipo | Responsabilidade | Tecnologia |
|---|---|---|---|
| **LoadTestRunner** | ApplicationRunner | Validar existência dos CSVs, disparar Job, logar métricas, encerrar processo com exit code 0/1 | Spring Boot `ApplicationRunner` |
| **BatchConfig** | @Configuration | Definir Job (`importacaoGeolocalidadeJob`), 3 Steps, 3 Readers (FlatFileItemReader), 3 Writers (RepositoryItemWriter) | Spring Batch Java DSL |
| **MunicipioProcessor** | ItemProcessor<DTO, Entity> | Transformar CSV em entidades UF→RegInter→RegImed→Municipio com cache HashMap | Java, JPA CascadeType.MERGE |
| **DistritoProcessor** | ItemProcessor<DTO, Entity> | Transformar CSV em Distrito referenciando Municipio via `getReferenceById` | Java, JPA proxy |
| **SubdistritoProcessor** | ItemProcessor<DTO, Entity> | Transformar CSV em Subdistrito referenciando Distrito via `getReferenceById` | Java, JPA proxy |
| **Repositories** | JpaRepository | Persistência e consulta de entidades | Spring Data JPA |

## Fluxos Principais

### Fluxo: stepMunicipio (chunk de 100)

1. `FlatFileItemReader<MunicipioCsvDTO>` lê 100 linhas de `DTB_Municipios.csv` (pula 7 linhas de metadados)
2. Para cada linha, `MunicipioProcessor.process(item)`:
   - `ufCache.computeIfAbsent(id, ...)` → cria `Uf` com sigla do mapa estático `UF_SIGLAS`
   - `interCache.computeIfAbsent(id, ...)` → cria `RegiaoIntermediaria`
   - `imedCache.computeIfAbsent(id, ...)` → cria `RegiaoImediata`
   - Retorna `new Municipio(id, codigo, nome, imediata)`
3. `RepositoryItemWriter<Municipio>` persiste 100 entidades via `municipioRepository.save()`
4. Transação commitada; caches mantidos na memória até fim do step

### Fluxo: stepDistrito (chunk de 100)

1. `FlatFileItemReader<DistritoCsvDTO>` lê 100 linhas de `DTB_Distritos.csv`
2. Para cada linha, `DistritoProcessor.process(item)`:
   - `municipioRepository.getReferenceById(item.municipioId())` → proxy JPA (sem SELECT)
   - Retorna `new Distrito(id, codigo, nome, municipioRef)`
3. `RepositoryItemWriter<Distrito>` persiste 100 entidades

### Fluxo: stepSubdistrito (chunk de 100)

1. Idêntico ao stepDistrito, mas lendo `DTB_Subdistritos.csv` e referenciando `Distrito`
