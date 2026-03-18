# ms-geolocalidade

Microserviço para consulta de geolocalização por CEP e descoberta de localidades dentro de um raio, integrando:

- **AwesomeAPI** (geocodificação e busca por raio)
- **Base local DTB/IBGE** (enriquecimento oficial via tabelas já persistidas)

## Requisitos

- Java 21
- Maven Wrapper (`./mvnw`)

## Execução local (dev)

Por padrão o serviço sobe com H2 (memória).

```bash
./mvnw spring-boot:run
```

## Endpoint

`GET /api/v1/localidades/vizinhas-agil?cep=01001000&raio=5`

## Variáveis de ambiente

- `AWESOME_API_TOKEN` (usado na chamada `/json/{cep}` via query param)
- `AWESOME_API_KEY` (usado na chamada `/search` via header `x-api-key`)


## Build nativo (GraalVM)

```bash
./mvnw -DskipTests native:compile
```

## Docker (imagem nativa)

O `Dockerfile` assume build com **context na raiz do monorepo**.

```bash
docker build -f backend/java/spring/microservices/ms-geolocalidade/Dockerfile -t ms-geolocalidade:latest .
```
