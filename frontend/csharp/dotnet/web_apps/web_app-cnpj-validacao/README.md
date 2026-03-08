# web_app-cnpj-validacao (Blazor Web App)

Página simples para entrada de CNPJ com máscara `99.999.999/9999-99` e validação do CNPJ (normalização + cálculo de dígitos verificadores).

## Rodar com Docker

Na pasta `dotnet_maui/web_app-cnpj-validacao`:

```bash
docker compose up --build
```

Abrir no browser:

- http://localhost:18080

Para parar:

```bash
docker compose down
```

## Testes E2E (via browser em container)

Executa um Playwright (Chromium) em container, acessando a aplicação pelo hostname do `docker compose`.

```bash
docker compose -f docker-compose.e2e.yml up --build --abort-on-container-exit --exit-code-from e2e
```

Limpeza:

```bash
docker compose -f docker-compose.e2e.yml down
```

## Contrato de UI (paridade com React/Angular)

A página mostra:

- `input`: o valor digitado
- `normalized`: somente dígitos (ou `null` quando vazio)
- `valid`: `true/false` (ou `null` quando vazio)
