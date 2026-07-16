# Visão Arquitetural — taxnexus-individual-core-lib

> Gerado pelo **Arquiteto** (Reversa) em 2026-06-10 · Atualizado pelo **Revisor** em 2026-06-12
> Sintetiza: `inventory.md`, `dependencies.md`, `code-analysis.md`, `data-dictionary.md`, `domain.md`, `state-machines.md`, `permissions.md`, `adrs/` e as respostas humanas de `questions.md` (D1–D8, A1–A3).
> Confiança: 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

---

## 1. Resumo executivo

`taxnexus-individual-core-lib` (módulo Go `taxnexus-core-lib`) é uma **biblioteca de acesso a parâmetros fiscais** para tributação de pessoa física no Brasil (**IRPF** e **INSS** 🟢 D1). Ela **não calcula** o imposto: fornece a **camada de leitura versionada por vigência temporal** que os serviços de cálculo externos consomem.

Os serviços consumidores (expostos via API REST em `:3000` 🟢 D2) são onde residem: a fórmula de cálculo (confirmada como `base × alíquota − parcela a deduzir` 🟢 D3), a comparação de cenários para recomendação (🟢 D4), o controle de imposto ativo (🟢 D5), a semântica mensal/anual (🟢 D7) e a autorização/multi-tenant (🟢 D8). Esta lib é deliberadamente um **núcleo de dados puro** — essa fronteira de responsabilidade é a decisão arquitetural central do sistema.

| Atributo | Valor | Confiança |
|----------|-------|-----------|
| Linguagem | Go 1.25.6 | 🟢 |
| Tipo de artefato | Biblioteca (não executável) | 🟢 |
| Persistência | PostgreSQL (schema `individual_tax_rates`) via `pgxpool` 5.9.1 | 🟢 |
| Cache | Redis via `go-redis` 9.18.0 (cache-aside, TTL 12h) | 🟢 |
| Precisão numérica | `shopspring/decimal` 1.4.0 (nunca `float`) | 🟢 |
| IDs de auditoria | `google/uuid` 1.6.0 | 🟢 |
| Testes | Nenhum no recorte (`test_conn.go` ausente) | 🔴 |
| CI/CD / Docker | Ausente no recorte | 🔴 |

---

## 2. Estilo e princípios arquiteturais

- **Biblioteca de camada de dados (Repository pattern)** — uma única struct `TaxRepository` agrega o acesso a Postgres + Redis, exposta por API pública de pacote. 🟢
- **Injeção de dependência por construtor** — `NewTaxRepository(db, rdb)`; sem singletons globais (ver `adrs/0005`). 🟢
- **Versionamento por vigência temporal** — toda leitura é filtrada por `valid_from <= refDate AND (valid_to IS NULL OR valid_to >= refDate)`, permitindo cálculo retroativo (ver `adrs/0004`). 🟢
- **Cache-aside com degradação graciosa** — Redis é otimização; falha de cache cai para Postgres sem abortar (ver `adrs/0003`, `0006`). 🟢
- **Resolução de faixa em memória** — `GetApplicableRule` carrega todas as faixas vigentes (cacheáveis em uma chave) e seleciona em memória, em vez de filtrar no SQL (ver `adrs/0002`). 🟢
- **Núcleo puro / cálculo externalizado** — a lib lê parâmetros; cálculo, ativação (`active`), período (mensal/anual) e autorização ficam nos serviços consumidores (🟢 D5, D7, D8).

---

## 3. Decomposição em módulos

| Módulo | Responsabilidade | Lógica de negócio | Complexidade |
|--------|------------------|-------------------|--------------|
| `repository` | Acesso a dados com cache-aside; resolução de faixa progressiva; vigência temporal | **Sim** | média |
| `models` | Structs de domínio (entidades) + DTOs de transporte | Não | baixa |
| `db` | Fábrica de pool PostgreSQL (`ConnectPostgres`) | Não | baixa |
| `cache` | Fábrica de cliente Redis (`ConnectRedis`) | Não | baixa |

API pública do `repository` (contratos operacionais que o consumidor usa):

| Função | Papel |
|--------|-------|
| `GetApplicableRule(ctx, taxCode, baseValue, refDate)` | Resolve a faixa progressiva aplicável a uma base |
| `GetConfig(ctx, taxCode, key, refDate)` | Lê um parâmetro único de `tax_configs` (sem cache) |
| `GetTableConfigs(ctx, taxCode, refDate)` | Lê todos os parâmetros do imposto (cacheado) |
| `GetTaxRulesForPeriod(ctx, taxCode, refDate)` | Lista as faixas vigentes (cacheado) |

---

## 4. Catálogo de parâmetros de configuração (`tax_configs`) 🟢

Confirmado via payloads reais da API consumidora (D2). Chaves esperadas por `config_key`:

| `config_key` | Unidade | Significado |
|--------------|---------|-------------|
| `pension_percentage` | percentage | Percentual de dedução previdenciária (ex.: 11,0%) |
| `dependents_qty` | count | Quantidade de dependentes (multiplica valor por dependente) |
| `pgbl_contribution` | amount | Contribuição PGBL dedutível (limitada por % da renda) |
| `education_expenses` | amount | Despesas de educação (teto legal) |
| `health_expenses` | amount | Despesas de saúde (sem teto) |

---

## 5. Integrações externas

| Integração | Direção | Protocolo | Configuração | Confiança |
|------------|---------|-----------|--------------|-----------|
| **PostgreSQL** | consome | TCP / wire pgx | `DATABASE_URL` | 🟢 |
| **Redis** | consome | RESP (go-redis) | `REDIS_ADDR` | 🟢 |
| **Serviços de cálculo** | é consumida por | chamada de função Go (in-process) | — | 🟢 (D2 revela API REST `POST /api/v1/calculate/irpf` na frente do serviço) |

Esta biblioteca **não expõe HTTP/gRPC** diretamente; sua "API" é a superfície de funções Go do pacote `repository`. A API REST observada pertence aos serviços consumidores.

---

## 6. Dívidas técnicas e riscos

| # | Item | Severidade | Origem |
|---|------|------------|--------|
| T1 | Ausência total de testes; `test_conn.go` citado no README não existe | alta | L3 |
| T2 | TTL de cache (12h) hardcoded — não configurável por ambiente | média | L4 / `adrs/0003` |
| T3 | `GetConfig` não usa cache enquanto `GetTableConfigs` usa — estratégia inconsistente | média | L5 |
| T4 | `GetConfig` retorna `(decimal.Zero, err)`; chamador que ignore `err` lê zero plausível | média | code-analysis A3 |
| T5 | `documentoFiscalRequest` não exportada usada em campo exportado (`UniversalTaxRequest.Inputs`) | baixa | models |
| T6 | `ConnectPostgres` usa `context.Background()` sem timeout; pool sem tuning | baixa | db |
| T7 | `ConnectRedis` sem `Ping`/auth/TLS — falha só no primeiro comando | baixa | cache |
| T8 | Erros de `json.Marshal` no `SET` do cache são ignorados silenciosamente | baixa | repository A3 |
| T9 | Sem migrations versionadas no recorte (schema confirmado via DDL em A3) | média | L6 |
| T10 | `TaxCalculationLog` definido mas não escrito pela lib (responsabilidade externa confirmada) | baixa | A2 |

---

## 7. Lacunas remanescentes após validação humana

As lacunas de domínio D1–D8 e estruturais A1–A3 foram **resolvidas**. O que permanece em aberto refere-se à infraestrutura e testes do recorte:

| # | Lacuna | Estado |
|---|--------|--------|
| L3 🔴 | Estratégia de testes automatizados | `test_conn.go` ausente; sem suíte de testes no recorte |
| L4 🟡 | Parametrização de ambiente (TTL, timeouts) | Atualmente hardcoded ou default |

---

## Referências cruzadas
- C4: `c4-context.md`, `c4-containers.md`, `c4-components.md`
- Dados: `erd-complete.md`, `data-dictionary.md`
- Domínio/regras: `domain.md`, `state-machines.md`, `permissions.md`
- Decisões: `adrs/0001`–`adrs/0006`
- Rastreabilidade: `traceability/spec-impact-matrix.md`
