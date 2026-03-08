# ms-cnpj-validacao (Go)

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

Observações de compatibilidade:

- `normalized` contém apenas dígitos.
- Se o input for `null`/ausente/vazio (após normalização), `normalized` vira `null` e `valid` é `false`.

## Rodar localmente

Requisitos:

- Go 1.22+

Rodar:

```bash
go run .
```

O serviço sobe em `0.0.0.0:8080` por padrão. Para trocar a porta:

```bash
PORT=9090 go run .
```

Teste rápido:

```bash
curl 'http://localhost:8080/cnpj/validate?value=04.252.011/0001-10'
```

## Testes

```bash
go test ./...
```

## Docker

O projeto inclui:

- `Dockerfile` multi-stage (build do binário no stage `builder` e runtime leve em Alpine)
- `docker-compose.yml` expondo o serviço em `18080:8080`

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
docker build -t ms-cnpj-validacao:go .
docker run --rm -p 18080:8080 -e PORT=8080 ms-cnpj-validacao:go
```
