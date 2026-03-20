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

ou

```bash
cd backend/java/spring/microservices/ms-geolocalidade && ./mvnw -q spring-boot:run -Dspring-boot.run.arguments=--server.port=8081
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


## ADENDO DE INFORMACOES DE SITES PESQUISADOS COMO BASE DE CONHECIMENTO

https://www.ibge.gov.br/geociencias/organizacao-do-territorio/divisao-regional/23701-divisao-territorial-brasileira.html

https://www.ibge.gov.br/estatisticas/sociais/populacao/38734-cadastro-nacional-de-enderecos-para-fins-estatisticos.html?edicao=38891&t=resultados

https://www.ibge.gov.br/geociencias/organizacao-do-territorio/estrutura-territorial.html



