# cache — Requisitos

> Fábrica de cliente Redis (conexão lazy) via `go-redis`.
> Gerado pelo **Redator** (Reversa) em 2026-06-10 · `doc_level = completo`
> Fontes: `cache/redis.go`, `surface.json`, `code-analysis.md`.
> 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

## Visão Geral

O pacote `cache` é uma **fábrica de infraestrutura**: cria e devolve um cliente Redis (`*redis.Client`) a partir de um endereço. É a dependência usada pelo `repository` para a estratégia de cache-aside (TTL 12h) das tabelas de regras e configurações.

## Responsabilidades

- Construir um `*redis.Client` a partir de um endereço (`REDIS_ADDR`). 🟢

## Regras de Negócio

- Não há regra de negócio — é infraestrutura pura. 🟢
- O cliente é criado **somente com `Addr`** — sem `Ping`, sem password, sem DB index, sem TLS. Conexão **lazy** (estabelecida no primeiro comando). 🟢

## Requisitos Funcionais

| ID | Requisito | Prioridade | Critério de Aceite |
|----|-----------|-----------|-------------------|
| RF-01 | `ConnectRedis(addr) *redis.Client` retorna um cliente configurado com o endereço informado | Must | Cliente não-nil retornado para qualquer `addr`; conexão lazy |
| RF-02 | A indisponibilidade do Redis não quebra a inicialização (conexão lazy) | Must | `ConnectRedis` retorna cliente mesmo com Redis fora; falha só aparece no uso, e é tratada via degradação graciosa no `repository` |

## Requisitos Não Funcionais

| Tipo | Requisito inferido | Evidência no código | Confiança |
|------|--------------------|---------------------|-----------|
| Disponibilidade | Conexão lazy + degradação graciosa no consumidor (`repository`) | `cache/redis.go:8` + `repository/tax_repository.go` | 🟢 |
| Configurabilidade | Endereço externo via `REDIS_ADDR` (12-factor) | `surface.json.env_vars` | 🟢 |
| Segurança | Sem TLS/password no recorte — possível pressuposto de rede interna | `cache/redis.go:8` | 🟡 |

## Critérios de Aceitação

```gherkin
Dado um endereço Redis qualquer
Quando ConnectRedis é chamado
Então retorna um *redis.Client não-nil (conexão lazy, sem Ping)

Dado que o Redis está indisponível
Quando ConnectRedis é chamado
Então ainda retorna um cliente; a falha só se manifesta no primeiro comando e é absorvida pela degradação graciosa do repository
```

## Prioridade (MoSCoW)

| Requisito | MoSCoW | Justificativa |
|-----------|--------|---------------|
| Criação do cliente (RF-01) | Must | Dependência do cache-aside do `repository` |
| Conexão lazy / não bloquear init (RF-02) | Must | Sustenta a degradação graciosa |
| TLS / password / DB index | Could | Ausentes; melhorias de segurança/config (ver Lacunas) |

## Rastreabilidade de Código

| Arquivo | Função / Classe | Cobertura |
|---------|-----------------|-----------|
| `cache/redis.go` | `ConnectRedis` | 🟢 |
