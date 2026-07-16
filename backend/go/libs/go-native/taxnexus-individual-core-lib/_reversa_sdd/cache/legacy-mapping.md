# Legacy Mapping — módulo `cache`

> Gerado pelo **Arqueólogo** (Reversa) em 2026-06-10

| Símbolo | Tipo | Arquivo:linha | Notas |
|---------|------|---------------|-------|
| `ConnectRedis` | função | `cache/redis.go:7` | `redis.NewClient` com `Options{Addr: addr}` (l.8) |

Pacote: `cache`. Dependência: `github.com/redis/go-redis/v9` (l.4). Configuração via `REDIS_ADDR`. Conexão lazy (sem `Ping`).
