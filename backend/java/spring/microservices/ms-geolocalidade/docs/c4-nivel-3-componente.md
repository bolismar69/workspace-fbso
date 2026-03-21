# C4 — Nível 3 (Componente) — ms-geolocalidade

## Objetivo

Descrever os **componentes** internos do contêiner "API Spring Boot" do `ms-geolocalidade`, com foco em responsabilidades e dependências.

## Componentes principais (por pacote)

### Controllers (`com.fbso.geolocalidade.controller`)

- `LocalidadeController`
  - Endpoint: `GET /api/v1/localidades/vizinhas-agil`
  - Caso de uso: vizinhança por CEP (AwesomeAPI + enriquecimento DTB/IBGE)

- `CepController`
  - Endpoint: `GET /api/v1/localidades/cep/{cep}`
  - Caso de uso: geocodificação por CEP (AwesomeAPI)

- `UfController`
  - Endpoints: `GET /api/v1/localidades/uf`, `GET /api/v1/localidades/uf/{sigla}/municipios`, ...
  - Caso de uso: navegação UF → municípios (DTB/IBGE)

- `MunicipioController`
  - Endpoints: `GET /api/v1/localidades/municipios`, `.../id/{id}`, `.../nome/{nome}`
  - Caso de uso: busca/paginação em municípios (DTB/IBGE)

### Services (`com.fbso.geolocalidade.service`)

- `AwesomeCepService`
  - Integração HTTP com AwesomeAPI (`/json/{cep}`, `/search`)
  - Cache: `@Cacheable` em `obterCoordenadas(cep)`

- `LocalidadeService`
  - Orquestra: geocodificação do CEP, busca por raio, enriquecimento via repositórios

- `UfService`, `MunicipioService`
  - Consultas paginadas e montagem de DTOs a partir das entidades DTB/IBGE

### Repositories (`com.fbso.geolocalidade.repository`)

- `MunicipioRepository`, `UfRepository`
  - Leitura com `@EntityGraph` para carregar hierarquia UF/Regiões

- `DistritoRepository`, `SubdistritoRepository`
  - Busca de primeiro distrito/subdistrito (payload resumido)
  - Consulta por prefixo do código IBGE (`SubdistritoRepository.findNomeByCodigo`)

### Config e infraestrutura (`com.fbso.geolocalidade.config`)

- `ClientConfig`
  - Bean de `RestClient` com `HttpClient` (timeouts)

- `AwesomeApiProperties`
  - Config: base URL + token/key

- `CacheConfig` / `CacheProperties`
  - Nome do cache e TTL/tamanho (Caffeine)

- `RequestIdFilter`
  - Propaga `requestId` em MDC para logging

## Diagrama (Componentes)

```mermaid
flowchart TB
  %% C4-L3: Component diagram (representado em Mermaid)

  subgraph api["Container: API Spring Boot (ms-geolocalidade)"]
    subgraph controllers["Controllers"]
      LocalidadeController["LocalidadeController\nGET /localidades/vizinhas-agil"]
      CepController["CepController\nGET /localidades/cep/{cep}"]
      UfController["UfController\nGET /localidades/uf/**"]
      MunicipioController["MunicipioController\nGET /localidades/municipios/**"]
    end

    subgraph services["Services"]
      LocalidadeService["LocalidadeService\nOrquestra vizinhança por CEP"]
      AwesomeCepService["AwesomeCepService\nIntegra AwesomeAPI + cache"]
      UfService["UfService\nConsulta UFs"]
      MunicipioService["MunicipioService\nConsulta municípios"]
    end

    subgraph repos["Repositories (JPA)"]
      UfRepository["UfRepository"]
      MunicipioRepository["MunicipioRepository"]
      DistritoRepository["DistritoRepository"]
      SubdistritoRepository["SubdistritoRepository"]
    end

    subgraph infra["Infra/Config"]
      RestClient["RestClient (Spring)"]
      Cache["Caffeine Cache"]
      RequestIdFilter["RequestIdFilter"]
    end
  end

  db[("PostgreSQL\nDTB/IBGE (schema localidade)")]
  awesome["AwesomeAPI\n/json/{cep}, /search"]

  %% Controller -> Service
  LocalidadeController --> LocalidadeService
  CepController --> AwesomeCepService
  UfController --> UfService
  MunicipioController --> MunicipioService

  %% Service -> Service
  LocalidadeService --> AwesomeCepService

  %% Service -> Repos
  LocalidadeService --> MunicipioRepository
  LocalidadeService --> SubdistritoRepository

  UfService --> UfRepository
  MunicipioService --> MunicipioRepository
  MunicipioService --> DistritoRepository
  MunicipioService --> SubdistritoRepository

  %% Infra
  AwesomeCepService --> RestClient
  AwesomeCepService --> Cache
  RestClient --> awesome

  %% DB access
  UfRepository --> db
  MunicipioRepository --> db
  DistritoRepository --> db
  SubdistritoRepository --> db
```

## Observações

- O caso de uso "vizinhas por CEP" combina dados **dinâmicos** (AwesomeAPI) com dados **oficiais** (DTB/IBGE).
- O cache é aplicado apenas na geocodificação (`/json/{cep}`), não na busca por raio.

