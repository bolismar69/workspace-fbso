# web_app-cnpj-validacao (React + TypeScript)

Página simples para entrada de CNPJ com máscara `99.999.999/9999-99` e validação do dígito verificador, equivalente ao app em [angular_js/web_app-cnpj-validacao](angular_js/web_app-cnpj-validacao).

## Funcionalidades

- Campo de entrada com máscara `99.999.999/9999-99`
- Normalização (somente dígitos)
- Validação de CNPJ (mesmo algoritmo usado nos microsserviços)

## Rodar com Docker (recomendado)

```bash
docker compose up --build
```

Abra no browser:

- http://localhost:18080

## Teste via “browser” em container (Playwright)

Esse fluxo sobe o container da aplicação e roda um teste e2e (Chromium headless) validando:

- A página abre
- O campo aceita um CNPJ válido
- O resultado mostra `valid: true` e `normalized: 04252011000110`

Rodar:

```bash
docker compose -f docker-compose.e2e.yml up --build --abort-on-container-exit --exit-code-from e2e
```

Limpar:

```bash
docker compose -f docker-compose.e2e.yml down
```
