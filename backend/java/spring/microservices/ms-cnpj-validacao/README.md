# ms-cnpj-validacao (Java 21 + GraalVM native)

Micro-serviço HTTP para validar CNPJ.

## Endpoints

- `GET /cnpj/validate?value=04.252.011/0001-10`
- `POST /cnpj/validate` com body:

```json
{ "cnpj": "04.252.011/0001-10" }
```

Resposta:

```json
{ "input": "...", "normalized": "...", "valid": true }
```

## Rodar em modo dev

```bash
mvn quarkus:dev
```

## Build JVM

```bash
mvn package
java -jar target/quarkus-app/quarkus-run.jar
```

## Docker (imagem nativa)

O projeto inclui:

- `Dockerfile` (multi-stage) que compila binário nativo no build e roda em `distroless`.
- `docker-compose.yml` que sobe o serviço já mapeado em `18080:8080`.

### Subir com Docker Compose (recomendado)

```bash
docker compose up --build
```

Teste:

```bash
curl 'http://localhost:18080/cnpj/validate?value=04.252.011/0001-10'
```

### Build e run com Docker (sem compose)

```bash
docker build -t ms-cnpj-validacao:native .
docker run --rm -p 18080:8080 ms-cnpj-validacao:native
```

## Build nativo (GraalVM)

### Opção A: GraalVM instalado localmente

- Use GraalVM compatível com Java 21 e com o componente `native-image` instalado.

```bash
mvn -Pnative package
./target/ms-cnpj-validacao-0.1.0-SNAPSHOT-runner
```

### Opção B: build nativo via container (sem instalar GraalVM)

```bash
mvn -Pnative package -Dquarkus.native.container-build=true
```

> Dica: essa opção depende de Docker/Podman disponível na máquina.

O binário gerado costuma ficar em:

```bash
./target/ms-cnpj-validacao-0.1.0-SNAPSHOT-native-image-source-jar/ms-cnpj-validacao-0.1.0-SNAPSHOT-runner
```
