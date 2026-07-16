# Análise de Código — taxnexus-individual-core-lib

> Gerado pelo **Arqueólogo** (Reversa) em 2026-06-10 · `doc_level = completo`
> Escala de confiança: 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA
> Módulos analisados: `repository`, `models`, `db`, `cache`

---

## Sumário executivo

`taxnexus-core-lib` é uma **biblioteca Go de acesso a dados** para um motor de cálculo de impostos sobre pessoa física (IRPF/INSS no Brasil). Ela **não calcula** o imposto neste recorte — fornece a **camada de leitura** (regras, faixas e parâmetros de configuração) que um serviço de cálculo (ausente do repositório) consumiria.

Arquitetura em 4 pacotes:

```
┌─────────────┐     usa      ┌──────────────┐
│ repository  │─────────────▶│   models     │  (structs de domínio + DTOs)
│ (TaxRepo)   │              └──────────────┘
└──────┬──────┘
       │ injeta no construtor
   ┌───┴────────────┐
   ▼                ▼
┌──────┐        ┌────────┐
│  db  │        │ cache  │
│ (PG) │        │(Redis) │
└──────┘        └────────┘
```

| Módulo | Complexidade | Responsabilidade | Lógica de negócio? |
|--------|--------------|------------------|--------------------|
| `repository` | **média** | Acesso a dados com cache-aside; resolução de faixa de imposto | Sim (faixa progressiva + vigência temporal) 🟢 |
| `models` | baixa | Definições de structs/DTOs | Não (apenas dados) 🟢 |
| `db` | baixa | Fábrica de pool PostgreSQL | Não 🟢 |
| `cache` | baixa | Fábrica de cliente Redis | Não 🟢 |

---

## Módulo `repository` 🟢

**Arquivo:** `repository/tax_repository.go` (154 linhas)

Estrutura central:

```go
type TaxRepository struct {
    db  *pgxpool.Pool   // pool PostgreSQL
    rdb *redis.Client   // cliente Redis (cache)
}
```

Injeção de dependência explícita via construtor `NewTaxRepository(db, rdb)` — sem singletons globais. 🟢

### Funções

| Função | Assinatura | Retorno | Linha | Confiança |
|--------|-----------|---------|-------|-----------|
| `NewTaxRepository` | `(db *pgxpool.Pool, rdb *redis.Client)` | `*TaxRepository` | 21 | 🟢 |
| `GetApplicableRule` | `(ctx, taxCode string, baseValue decimal.Decimal, refDate time.Time)` | `(*models.TaxRule, error)` | 26 | 🟢 |
| `GetConfig` | `(ctx, taxCode, key string, refDate time.Time)` | `(decimal.Decimal, error)` | 44 | 🟢 |
| `GetTableConfigs` | `(ctx, taxCode string, refDate time.Time)` | `(map[string]decimal.Decimal, error)` | 61 | 🟢 |
| `GetTaxRulesForPeriod` | `(ctx, taxCode string, refDate time.Time)` | `([]models.TaxRule, error)` | 105 | 🟢 |

### Algoritmos e regras embutidas

#### A1 — Resolução de faixa aplicável (`GetApplicableRule`) 🟢
Decisão de design relevante: **a seleção de faixa NÃO é feita no SQL**. A função busca *todas* as regras vigentes via `GetTaxRulesForPeriod` (que é cacheada) e percorre em memória, retornando a **primeira** regra cujo intervalo contém o `baseValue`:

```
para cada rule em rules (ordenadas por range_min ASC):
    se baseValue >= rule.RangeMin:
        se rule.RangeMax == nil  OU  baseValue <= rule.RangeMax:
            retorna rule          ← primeiro match vence
retorna erro "no applicable rule found for value X"
```

- `RangeMax == nil` representa a **faixa aberta superior** (último escalão sem teto). 🟢
- Comparações usam `decimal.Decimal.GreaterThanOrEqual` / `LessThanOrEqual` (intervalos **fechados** em ambos os lados: `[min, max]`). 🟢
- 🟡 **INFERIDO:** a ordenação ASC por `range_min` (garantida pela query) é o que torna "primeiro match" correto — faixas contíguas não se sobrepõem. Se os dados tiverem sobreposição, o menor `range_min` ganha.
- 🟡 Reaproveitar `GetTaxRulesForPeriod` (em vez de uma query `WHERE range_min <= v AND range_max >= v`) é uma otimização para **cache**: uma única chave Redis serve todas as consultas de faixa daquele imposto/data.

#### A2 — Filtro de vigência temporal (todas as queries) 🟢
Padrão repetido de *temporal validity* (versionamento por data):

```sql
valid_from <= $refDate AND (valid_to IS NULL OR valid_to >= $refDate)
```

- `valid_to IS NULL` ⇒ regra/config vigente indefinidamente. 🟢
- Aplica-se a `tax_configs` (`GetConfig`, `GetTableConfigs`) e a `tax_rules_history` (`GetTaxRulesForPeriod`). 🟢
- A tabela chama-se `tax_rules_history` — 🟡 indica modelo **append-only / historização** das faixas de imposto.

#### A3 — Cache-aside (Redis) 🟢
`GetTableConfigs` e `GetTaxRulesForPeriod` implementam o padrão *cache-aside*:

```
1. monta cacheKey
2. GET no Redis → se hit e desserializa OK, retorna
3. (miss) consulta Postgres
4. se resultado não-vazio: SET no Redis com TTL = 12h
5. retorna
```

- **Chaves de cache:**
  - `tax_configs:<taxCode>:<YYYY-MM-DD>` 🟢
  - `tax_rules_list:<taxCode>:<YYYY-MM-DD>` 🟢
- **TTL:** `12 * time.Hour` (constante embutida, hardcoded — não configurável) 🟢 / 🟡 candidato a parâmetro
- **Serialização:** JSON (`encoding/json`). Erros de `json.Marshal` no `SET` são **ignorados** (`cacheData, _ := ...`). 🟢
- **Degradação graciosa:** erro de Redis no `GET` não aborta — cai para o Postgres. 🟢
- ⚠️ `GetConfig` (valor único) **não usa cache** — vai direto ao Postgres. 🟢 (possível inconsistência de estratégia vs. `GetTableConfigs`)

### Tratamento de erros
- `GetApplicableRule`: erro explícito quando nenhuma faixa casa (`fmt.Errorf("no applicable rule found for value %s", ...)`). 🟢
- `GetConfig`: propaga erro do `Scan`; retorna `decimal.Zero` junto. 🟢 (⚠️ chamador deve checar `err`, pois `Zero` é valor plausível)
- `GetTaxRulesForPeriod`: envelopa erro de query com `%w` (`error querying tax rules: %w`) — preserva cadeia de erro. 🟢
- Erros de cache (Redis) são deliberadamente engolidos para não derrubar o caminho feliz. 🟢

### Dependências do módulo
`models` (entidades), `pgx/v5/pgxpool`, `go-redis/v9`, `shopspring/decimal`, `encoding/json`, `fmt`, `time`, `context`. 🟢

> **Fluxograma:** ver `flowcharts/repository.md`.
> **Mapa de legado:** ver `repository/legacy-mapping.md`.

---

## Módulo `models` 🟢

**Arquivo:** `models/tax_models.go` (89 linhas) — somente definições de tipo, **sem lógica**.

8 structs (3 entidades persistidas + 5 DTOs de transporte). Dicionário completo em `data-dictionary.md`. Resumo:

| Struct | Papel | Persistida? |
|--------|-------|-------------|
| `TaxDefinition` | Definição de um imposto (código, nome, esfera, precisão) | 🟢 tabela `tax_definitions` |
| `TaxRule` | Faixa/escalão de imposto com vigência | 🟢 tabela `tax_rules_history` |
| `TaxCalculationLog` | Log de um cálculo executado (com `uuid` e `trace_id`) | 🟡 inferida (`tax_calculation_log`) — não lida/escrita aqui 🔴 |
| `TaxRequest` | Payload de entrada do simulador (mensal/anual, deduções IRPF) | DTO |
| `documentoFiscalRequest` | Item de input genérico (não exportado) | DTO interno |
| `UniversalTaxRequest` | Payload genérico de cálculo (`inputs` variáveis) | DTO |
| `DeductionDetail` | Detalhamento de uma dedução na saída | DTO |
| `TaxResponse` | Resultado do cálculo (base, imposto, alíquota efetiva, configs usadas) | DTO |

### Observações de domínio
- Valores monetários e alíquotas usam `shopspring/decimal.Decimal` (nunca `float`) — decisão correta para fiscal/contábil. 🟢
- Ponteiros (`*decimal.Decimal`, `*time.Time`) modelam colunas **NULL-áveis** (`range_max`, `valid_to`). 🟢
- `TaxRequest` revela o **escopo funcional do simulador** (camada ausente): tipo `monthly`/`annual`, deduções de previdência (valor fixo **e** percentual), dependentes, despesas de educação/saúde, PGBL, outras deduções. 🟡 → forte sinal das regras de IRPF brasileiro.
- `TaxResponse.IsRecommended` (bool) e `AppliedRule`/`UsedConfigs` sugerem que o motor compara cenários (ex.: dedução simplificada vs. completa) e marca o recomendado. 🟡 🔴 LACUNA (lógica não presente).
- `documentoFiscalRequest` é **não exportada** mas usada em campo de struct exportada (`UniversalTaxRequest.Inputs`) — JSON funciona, mas consumidores externos não conseguem instanciar o item diretamente. 🟡 possível smell.

> **Fluxograma:** N/A (sem fluxo de controle). Modelo de dados em `flowcharts/models.md` (diagrama de relacionamento de structs).

---

## Módulo `db` 🟢

**Arquivo:** `db/postgres.go` (15 linhas)

```go
func ConnectPostgres(connString string) (*pgxpool.Pool, error)
```

- Faz `pgxpool.ParseConfig(connString)` → `pgxpool.NewWithConfig(ctx, config)`. 🟢
- Usa `context.Background()` (sem timeout/cancel) na criação do pool. 🟡 (sem deadline de conexão)
- Não há configuração de pool (max conns, lifetime) além do default do pgx. 🟡
- `connString` vem de `DATABASE_URL` (ver README). 🟢

> **Fluxograma:** trivial — ver `flowcharts/db.md`.

---

## Módulo `cache` 🟢

**Arquivo:** `cache/redis.go` (11 linhas)

```go
func ConnectRedis(addr string) *redis.Client
```

- `redis.NewClient(&redis.Options{Addr: addr})` — apenas endereço. 🟢
- **Não retorna erro** e **não faz `Ping`** — a conexão é lazy; falha só aparece no primeiro comando. 🟡
- Sem auth (password), sem DB index, sem TLS. 🟡 (provavelmente OK para ambiente local/interno — `REDIS_ADDR`)

> **Fluxograma:** trivial — ver `flowcharts/cache.md`.

---

## Lacunas consolidadas (para Detective / Architect / Data Master)

| # | Lacuna | Impacto |
|---|--------|---------|
| L1 🔴 | Camada de **cálculo** (consumidora dos DTOs `TaxRequest`/`TaxResponse`) ausente — recorte de monorepo | Regras de cálculo IRPF (deduções, escolha mensal/anual, "recomendado") não documentáveis a partir do código |
| L2 🔴 | `TaxCalculationLog` não é lido nem escrito por nenhuma query | Persistência de auditoria está em outra camada |
| L3 🔴 | `test_conn.go` citado no README está ausente | Sem ponto de smoke-test reprodutível |
| L4 🟡 | TTL de cache (12h) hardcoded | Possível necessidade de config por ambiente |
| L5 🟡 | `GetConfig` não usa cache enquanto `GetTableConfigs` usa | Inconsistência de estratégia a confirmar com o time |
| L6 🟡 | Schema do banco só existe implícito nas queries (sem DDL/migrations) | Data Master precisa reconstruir o ERD por inferência |
