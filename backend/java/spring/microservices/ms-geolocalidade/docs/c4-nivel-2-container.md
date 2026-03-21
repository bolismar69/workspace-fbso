# C4 — Nível 2 (Contêiner) — ms-geolocalidade

## Objetivo

Detalhar os **contêineres** (executáveis/armazenamentos) que compõem o `ms-geolocalidade` e suas integrações.

## Contêineres

- **API Spring Boot (Java 21)**
  - Expõe endpoints REST `/api/v1/localidades/**`.
  - Orquestra chamadas à AwesomeAPI e leituras no PostgreSQL.
  - Responsável por validação de entrada e padronização de resposta (`PageResponseDTO`).

- **Cache in-process (Caffeine)**
  - Cacheia geocodificação de CEP (`AwesomeCepService.obterCoordenadas`) via `@Cacheable`.

- **PostgreSQL (schema `localidade`)**
  - Armazena tabelas DTB/IBGE (UF, regiões, municípios, distritos, subdistritos).
  - O `ms-geolocalidade` tipicamente faz **somente leitura** dessas tabelas.

- **AwesomeAPI (externo)**
  - Geocodificação por CEP: `GET /json/{cep}`
  - Busca por raio: `GET /search?lat=...&lng=...&d=...`

## Diagrama (Contêiner)

```mermaid
flowchart LR
  %% C4-L2: Container diagram (representado em Mermaid)

  consumer["Consumidor\n(outros serviços / front-ends)"]

  subgraph sys["Software System: ms-geolocalidade"]
    api["Container: API Spring Boot\nJava 21 + Spring Web\nPorta 8080"]
    cache["Container: Caffeine Cache\n(in-memory, no processo)"]
  end

  db[("Container: PostgreSQL\nDTB/IBGE (schema localidade)")]
  awesome["External Container: AwesomeAPI\n(API CEP)"]

  consumer -->|"HTTP/JSON"| api

  api -->|"usa"| cache
  api -->|"JPA (read)"| db
  api -->|"HTTP (RestClient)"| awesome

  %% Observação: carga do DTB/IBGE é externa ao ms-geolocalidade
  batch["External System: batch-geolocalidade\n(carga DTB/IBGE)"]
  batch -->|"write/update"| db
```

## Tecnologias observadas no projeto

- Spring Boot + Spring Web + Validation (via parent)
- Spring Data JPA
- PostgreSQL driver
- Cache: `spring-boot-starter-cache` + Caffeine
- HTTP client: `RestClient` com `JdkClientHttpRequestFactory`

