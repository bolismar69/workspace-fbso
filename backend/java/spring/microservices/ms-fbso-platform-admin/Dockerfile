# =============================================================================
# Dockerfile — GraalVM Native Image (produção)
# =============================================================================
# Build: docker build -t fbso-platform-admin:latest .
# Run:   docker run -p 8080:8080 fbso-platform-admin:latest
# =============================================================================

FROM ghcr.io/graalvm/native-image-community:25 AS builder

WORKDIR /app
COPY pom.xml ./
COPY src ./src

# Build nativo
RUN ./mvnw -Pnative native:compile -DskipTests

# Imagem final mínima
FROM gcr.io/distroless/base-nossl-debian12:nonroot

WORKDIR /app
COPY --from=builder /app/target/ms-fbso-platform-admin /app/app

EXPOSE 8080

USER nonroot
ENTRYPOINT ["/app/app"]
