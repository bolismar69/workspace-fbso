# path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/Dockerfile
#
# Dockerfile multi-stage para ms-billing-engine-tax-rates (GAP-010).
#
# Build:
#   docker build \
#     -f backend/go/fiber/microservices/ms-billing-engine-tax-rates/Dockerfile \
#     -t ms-tax-rates:latest \
#     .
#
# O contexto de build DEVE ser a raiz do workspace (fbso/) para que o replace
# directive do go.mod (taxnexus-billing-core-lib => ../../../libs/go-native/...)
# seja resolvido. Como o replace usa caminho relativo ao go.mod do microserviço
# (backend/go/fiber/...), os caminhos no Dockerfile são mapeados da raiz.

# ─── Stage 1: Build ────────────────────────────────────────────────────────

FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /build

# Copia go.mod/go.sum do microserviço e da lib para cache de dependências
COPY backend/go/fiber/microservices/ms-billing-engine-tax-rates/go.mod \
     backend/go/fiber/microservices/ms-billing-engine-tax-rates/go.sum \
     ./backend/go/fiber/microservices/ms-billing-engine-tax-rates/

COPY backend/go/libs/go-native/taxnexus-billing-core-lib/go.mod \
     backend/go/libs/go-native/taxnexus-billing-core-lib/go.sum \
     ./backend/go/libs/go-native/taxnexus-billing-core-lib/

# Download de dependências (cache Docker se go.mod/go.sum não mudaram)
WORKDIR /build/backend/go/fiber/microservices/ms-billing-engine-tax-rates
RUN go mod download

# Copia o código fonte completo
WORKDIR /build
COPY backend/go/fiber/microservices/ms-billing-engine-tax-rates/ \
     ./backend/go/fiber/microservices/ms-billing-engine-tax-rates/
COPY backend/go/libs/ \
     ./backend/go/libs/

# Build
WORKDIR /build/backend/go/fiber/microservices/ms-billing-engine-tax-rates
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/ms-tax-rates ./cmd/api/

# ─── Stage 2: Runtime ──────────────────────────────────────────────────────

FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata && \
    addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=builder /app/bin/ms-tax-rates /app/bin/ms-tax-rates

USER app

EXPOSE 3000

# Health check usando liveness probe
HEALTHCHECK --interval=15s --timeout=3s --start-period=5s --retries=3 \
    CMD wget -qO- http://localhost:3000/v1/healthz || exit 1

ENTRYPOINT ["/app/bin/ms-tax-rates"]
