# cache — Tarefas de Implementação

> Sequência executável para reimplementar a fábrica do cliente Redis.
> Gerado pelo **Redator** (Reversa) em 2026-06-10 · `doc_level = completo`
> Fonte primária: `cache/redis.go`. Ver `requirements.md` e `design.md` desta unit.
> 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

## Pré-requisitos

- [ ] Dependência `github.com/redis/go-redis/v9` (v9.18.0) disponível
- [ ] Variável de ambiente `REDIS_ADDR` documentada para o chamador

## Tarefas

- [ ] **T-01** — Implementar `ConnectRedis(addr string) *redis.Client`: `redis.NewClient(&redis.Options{Addr: addr})`, conexão lazy, sem `Ping`.
  - Origem no legado: `cache/redis.go:8`
  - Critério de pronto: retorna cliente não-nil; não bloqueia/erra na criação mesmo com Redis fora
  - Confiança: 🟢

## Tarefas de Teste

- [ ] **TT-01** — `ConnectRedis` retorna cliente não-nil para um endereço qualquer
- [ ] **TT-02** — Integração: com Redis ativo, um `Set`/`Get` simples funciona via o cliente retornado

## Ordem Sugerida

1. **T-01** — unit autônoma; sem dependências internas.

## Lacunas Pendentes (🔴)

- Requisitos de segurança (TLS/password) e seleção de DB index — confirmar com operação se devem ser adicionados.
- Tuning do cliente (pool size, timeouts) — hoje usa defaults do go-redis.
