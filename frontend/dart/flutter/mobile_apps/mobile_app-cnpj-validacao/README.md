# mobile_app-cnpj-validacao (Flutter)

Aplicativo mobile (Android/iOS) em Flutter/Dart com:

- Entrada de CNPJ com máscara `99.999.999/9999-99`
- Normalização (somente dígitos)
- Validação (dígitos verificadores)

A tela exibe `input`, `normalized` e `valid` (equivalente ao projeto React Native `react-native/mobile_app-cnpj-validacao`).

## Rodar localmente (sem container)

Na pasta `flutter/mobile_app-cnpj-validacao`:

```bash
flutter pub get
flutter run
```

Obs.: para iOS nativo você precisa de macOS (simulador). Para Android, funciona no Linux com Android Studio/emulador ou device físico.

## Rodar no browser (para testes funcionais)

```bash
flutter run -d chrome
```

## Containerização (web build + Nginx)

Como não vamos publicar em lojas agora, o fluxo em container gera um build **Flutter Web** e serve via **Nginx** (útil para testes via browser).

Subir e acessar:

```bash
docker compose up --build
```

- http://localhost:18080

Parar:

```bash
docker compose down
```

## E2E via browser em container (Playwright)

Executa Playwright (Chromium) em container, acessando o app pelo hostname do compose:

```bash
docker compose -f docker-compose.e2e.yml up --build --abort-on-container-exit --exit-code-from e2e
```

Limpeza:

```bash
docker compose -f docker-compose.e2e.yml down
```

## Não-publicação

Este projeto **não** inclui pipeline de publicação (Google Play / Apple App Store) neste momento.
