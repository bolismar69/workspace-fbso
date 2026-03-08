# ms-cnpj-validacao (Python)

Micro-serviço HTTP para validar CNPJ, equivalente aos projetos Java/Go deste repositório.

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

Observações:

- `normalized` contém apenas dígitos.
- Se o input for `null`/ausente/vazio (após normalização), `normalized` é `null` e `valid` é `false`.

## Rodar localmente

Requisitos:

- Python 3.12+ (o código usa type hints com `str | None`)

Rodar:

```bash
python app.py
```

O serviço sobe em `0.0.0.0:8080` por padrão. Para trocar a porta:

```bash
PORT=9090 python app.py
```

Teste rápido:

```bash
curl 'http://localhost:8080/cnpj/validate?value=04.252.011/0001-10'
```

## Testes

```bash
python -m unittest -v
```

## Docker

O projeto inclui:

- `Dockerfile` multi-stage (cria um venv no stage builder e copia para a imagem final)
- `docker-compose.yml` expondo o serviço em `18080:8080`

### Subir com Docker Compose (recomendado)

```bash
docker compose up --build
```

Teste:

```bash
curl 'http://localhost:18080/cnpj/validate?value=04.252.011/0001-10'
```
