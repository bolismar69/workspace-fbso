# syntax=docker/dockerfile:1.7
ARG APP_PORT=8080
ARG IMAGE_TO_BUILD=ghcr.io/graalvm/native-image-community:latest
ARG IMAGE_TO_RUNTIME=gcr.io/distroless/cc-debian12

FROM ${IMAGE_TO_BUILD} AS build

# caminho do serviço dentro do monorepo
ARG SERVICE_PATH=backend/java/spring/microservices/ms-geolocalidade

WORKDIR /workspace
COPY . .

WORKDIR /workspace/${SERVICE_PATH}

RUN --mount=type=cache,target=/root/.m2 \
  set -eux; \
  ./mvnw -q -DskipTests -Ddependency-check.skip=true native:compile; \
  BIN=""; \
  if [ -d target ]; then BIN="$(find target -maxdepth 6 -type f -perm -111 | head -n 1 || true)"; fi; \
  if [ -z "$BIN" ]; then echo "ERRO: Binário nativo não encontrado." && exit 1; fi; \
  cp "$BIN" /workspace/app; \
  chmod +x /workspace/app

FROM ${IMAGE_TO_RUNTIME} AS runtime
ARG APP_PORT
WORKDIR /app
COPY --from=build /workspace/app /app/app
USER 10001
EXPOSE ${APP_PORT}
ENTRYPOINT ["/app/app"]
