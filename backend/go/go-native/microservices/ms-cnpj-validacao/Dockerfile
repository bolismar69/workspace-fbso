# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS builder
WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/ms-cnpj-validacao ./

FROM alpine:3.19
WORKDIR /app

RUN adduser -D -g '' appuser
COPY --from=builder /out/ms-cnpj-validacao ./ms-cnpj-validacao

USER appuser
EXPOSE 8080

ENV PORT=8080
CMD ["./ms-cnpj-validacao"]
