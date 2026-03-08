# mobile_app-cnpj-validacao (React Native / Expo SDK 53)

Aplicativo mobile (Android/iOS) feito em React Native + TypeScript com:

- Entrada de CNPJ com máscara `99.999.999/9999-99`
- Normalização (somente dígitos)
- Validação (dígitos verificadores)

A UI exibe `input`, `normalized` e `valid` (mesmo contrato visual dos projetos `angular_js/web_app-cnpj-validacao` e `react/web_app-cnpj-validacao`).

## Rodar localmente (Expo)

Na pasta `react-native/mobile_app-cnpj-validacao`:

```bash
npm install
npm run start
```

Outros alvos:

```bash
npm run android
npm run ios
npm run web
```

Obs.: iOS nativo requer macOS (simulador). Para testar funcionalidade rapidamente, `npm run web` abre no browser.

## Testes funcionais via browser (Docker + Playwright)

Como não vamos publicar nas lojas agora, o fluxo de testes aqui usa **Expo Web export** (build estático) + **Playwright** em containers.

Executar E2E:

```bash
docker compose -f docker-compose.e2e.yml up --build --abort-on-container-exit --exit-code-from e2e
```

Limpar containers:

```bash
docker compose -f docker-compose.e2e.yml down
```

## Rodar o build web em container

```bash
docker compose up --build
```

Abrir no browser:

- http://localhost:18080

Parar:

```bash
docker compose down
```

## Não-publicação

Este projeto **não** inclui pipeline de publicação (Google Play / Apple App Store) neste momento.
