# db — Requisitos

> Fábrica de pool de conexão PostgreSQL via `pgxpool`.
> Gerado pelo **Redator** (Reversa) em 2026-06-10 · `doc_level = completo`
> Fontes: `db/postgres.go`, `surface.json`, `code-analysis.md`.
> 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

## Visão Geral

O pacote `db` é uma **fábrica de infraestrutura**: cria e devolve um pool de conexões PostgreSQL (`*pgxpool.Pool`) a partir de uma connection string. É a dependência de I/O que o `repository` usa para todas as leituras do schema `individual_tax_rates`.

## Responsabilidades

- Construir um `*pgxpool.Pool` a partir de uma connection string (`DATABASE_URL`). 🟢
- Propagar erros de parsing/conexão ao chamador. 🟢

## Regras de Negócio

- Não há regra de negócio — é infraestrutura pura. 🟢
- O pool é criado com `pgxpool.ParseConfig` + `pgxpool.NewWithConfig` usando `context.Background()` (sem deadline na criação). 🟢

## Requisitos Funcionais

| ID | Requisito | Prioridade | Critério de Aceite |
|----|-----------|-----------|-------------------|
| RF-01 | `ConnectPostgres(connString) (*pgxpool.Pool, error)` retorna um pool válido para uma string correta | Must | String válida retorna pool não-nil e erro nil |
| RF-02 | Connection string inválida resulta em erro propagado (não panic) | Must | String malformada retorna erro descritivo e pool nil |

## Requisitos Não Funcionais

| Tipo | Requisito inferido | Evidência no código | Confiança |
|------|--------------------|---------------------|-----------|
| Escalabilidade | Uso de pool de conexões (`pgxpool`) em vez de conexão única | `db/postgres.go:9` | 🟢 |
| Configurabilidade | Connection string externa via `DATABASE_URL` (12-factor) | `surface.json.env_vars` | 🟢 |
| Disponibilidade | Sem `Ping`/deadline na criação — validação efetiva ocorre no primeiro uso | `db/postgres.go:9` | 🟡 |

## Critérios de Aceitação

```gherkin
Dado uma connection string válida de PostgreSQL
Quando ConnectPostgres é chamado
Então retorna um *pgxpool.Pool utilizável e erro nil

Dado uma connection string malformada
Quando ConnectPostgres é chamado
Então retorna erro descritivo e pool nil, sem panic
```

## Prioridade (MoSCoW)

| Requisito | MoSCoW | Justificativa |
|-----------|--------|---------------|
| Criação do pool (RF-01) | Must | Dependência de toda leitura do `repository` |
| Propagação de erro (RF-02) | Must | Falha de conexão precisa ser tratável pelo chamador |
| Ping/health-check na criação | Could | Ausente hoje; melhoria de robustez (ver Lacunas) |

## Rastreabilidade de Código

| Arquivo | Função / Classe | Cobertura |
|---------|-----------------|-----------|
| `db/postgres.go` | `ConnectPostgres` | 🟢 |
