# ms-cnpj-validacao (C# / ASP.NET Core)

Micro-serviço HTTP para validar CNPJ, equivalente aos projetos Java/Go/Python deste repositório.

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

## Rodar com Docker Compose (recomendado)

Esse projeto foi pensado para rodar mesmo sem `dotnet` instalado localmente.

```bash
docker compose up --build
```

Ele sobe em `18080:8080`.

Teste:

```bash
curl 'http://localhost:18080/cnpj/validate?value=04.252.011/0001-10'
```

## Build e run com Docker (sem compose)

```bash
docker build -t ms-cnpj-validacao:csharp .
docker run --rm -p 18080:8080 ms-cnpj-validacao:csharp
```

## Testes (xUnit) via Docker

Os testes ficam no projeto `MsCnpjValidacao.Tests` e rodam usando um stage `test` no `Dockerfile`.

### Opção A: docker compose (recomendado)

```bash
docker compose -f docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from tests
```

Depois limpe recursos:

```bash
docker compose -f docker-compose.test.yml down
```

### Opção B: docker build (direto)

```bash
docker build --target test .
```

> Dica: como o stage `test` tem `ENTRYPOINT ["dotnet", "test", ...]`, o build não executa os testes; ele apenas constrói a imagem. Para executar, use o compose acima ou `docker run` na imagem gerada.
