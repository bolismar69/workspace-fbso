# syntax=docker/dockerfile:1.7
ARG APP_PORT=8080
ARG IMAGE_TO_BUILD=ghcr.io/graalvm/native-image-community:latest
ARG IMAGE_TO_RUNTIME=gcr.io/distroless/cc-debian12

########################
# Builder GraalVM Native
########################
FROM ${IMAGE_TO_BUILD} AS build

# caminho do serviço dentro do repo (útil em monorepo)
ARG SERVICE_PATH=.

WORKDIR /workspace

# com context na raiz, copia o repo inteiro
COPY . .

# entra na pasta do serviço para executar o mvnw/pom daquele módulo
WORKDIR /workspace/${SERVICE_PATH}

# Build Native (wrapper obrigatório)
RUN --mount=type=cache,target=/root/.m2 \
  --mount=type=cache,target=/root/.gradle \
  set -eux; \
  if [ -x ./mvnw ]; then \
    ./mvnw -q -DskipTests -Pnative package; \
  elif [ -x ./gradlew ]; then \
    ./gradlew build nativeCompile; \
  else \
    echo "ERRO: Padronize mvnw/gradlew no repo." && exit 1; \
  fi

# Normalize binário -> /workspace/app
RUN set -eux; \
  BIN=""; \
  if [ -d target ]; then BIN="$(find target -maxdepth 6 -type f -perm -111 | head -n 1 || true)"; fi; \
  if [ -z "$BIN" ] && [ -d build ]; then BIN="$(find build -maxdepth 8 -type f -perm -111 | head -n 1 || true)"; fi; \
  if [ -z "$BIN" ]; then \
    echo "ERRO: Binário nativo não encontrado." && exit 1; \
  fi; \
  cp "$BIN" /workspace/app; \
  chmod +x /workspace/app

########################
# Runtime deps (libs)
########################
FROM debian:12-slim AS deps

RUN set -eux; \
  apt-get update; \
  apt-get install -y --no-install-recommends zlib1g; \
  rm -rf /var/lib/apt/lists/*

########################
# Runtime Native
########################
FROM ${IMAGE_TO_RUNTIME} AS runtime

ARG APP_PORT
WORKDIR /app
COPY --from=deps /lib/x86_64-linux-gnu/libz.so.1 /lib/x86_64-linux-gnu/libz.so.1
COPY --from=build /workspace/app /app/app
USER 10001
EXPOSE ${APP_PORT}
ENTRYPOINT ["/app/app"]
