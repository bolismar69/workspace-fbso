# cache — Design Técnico

> Implementação da fábrica do cliente Redis, com base em `cache/redis.go`.
> Gerado pelo **Redator** (Reversa) em 2026-06-10 · Atualizado pelo **Revisor** em 2026-06-12
> 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

## Interface

| Símbolo | Assinatura | Retorno | Observação |
|---------|-----------|---------|------------|
| `ConnectRedis` | `(addr string)` | `*redis.Client` | Fábrica de cliente; sem retorno de erro (lazy) 🟢 |

## Fluxo Principal

1. Recebe `addr` (origem típica: env `REDIS_ADDR`). `cache/redis.go` 🟢
2. `redis.NewClient(&redis.Options{Addr: addr})` — apenas o endereço é configurado. `cache/redis.go:8` 🟢
3. Retorna `*redis.Client`. Nenhum `Ping` é feito; a conexão é estabelecida no primeiro comando. 🟢

## Fluxos Alternativos

- **Redis indisponível:** não há erro na criação (lazy); o erro aparece no primeiro `Get`/`Set` e é absorvido pela **degradação graciosa** do `repository` (ADR-0006). 🟢

## Dependências

- **`github.com/redis/go-redis/v9`** (v9.18.0) — cliente Redis. 🟢

## Decisões de Design Identificadas

| Decisão | Evidência no código | Confiança |
|---------|---------------------|-----------|
| Conexão lazy (sem `Ping` na criação) | `cache/redis.go:8` | 🟢 |
| Configuração mínima (apenas `Addr`) | `cache/redis.go:8` | 🟢 |
| Sem retorno de erro na fábrica | assinatura `ConnectRedis(addr) *redis.Client` | 🟢 |

## Estado Interno

Nenhum no pacote. O `*redis.Client` retornado gerencia seu próprio pool interno; o ciclo de vida (Close) é do chamador. 🟢

## Observabilidade

Nenhuma emitida pelo pacote. 🟢

## Riscos e Lacunas

- 🟡 **Segurança:** sem TLS/password no recorte — pressuposto de rede interna confiável. 🟢
- 🟡 **Sem `Ping`/validação:** design de degradação graciosa torna aceitável a falha tardia. 🟢
- 🟢 **Tuning do cliente:** confirmado como default do go-redis.
