# Inventário do Projeto — taxnexus-individual-core-lib

> Gerado pelo **Scout** (Reversa) em 2026-06-10
> Escala de confiança: 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

---

## 1. Visão geral

| Item | Valor |
|------|-------|
| Nome do módulo Go | `taxnexus-core-lib` 🟢 (`go.mod`) |
| Pasta do repositório | `taxnexus-individual-core-lib` 🟢 |
| Tipo de artefato | **Biblioteca Go** (sem `package main` / sem entry point executável) 🟢 |
| Linguagem | Go 1.25.6 🟢 (`go.mod`) |
| Domínio de negócio | Cálculo de impostos sobre pessoa física (IRPF / INSS-like) no Brasil 🟡 |
| Total de arquivos `.go` | 4 🟢 |
| Caminho original (comentários) | `backend/go/libs/go-native/taxnexus-core-lib/` 🟡 (extraído de comentários `// path:` nos fontes — sugere ser um submódulo de um monorepo maior) |

> 🔴 **LACUNA:** Este repositório parece ser um recorte de um monorepo (`backend/go/libs/go-native/...`). A camada de serviço/cálculo que consome este `repository` **não está presente** aqui — só a camada de dados, modelos e conexões. Ver seção 7.

---

## 2. Estrutura de pastas

```
taxnexus-individual-core-lib/
├── go.mod                      # Definição do módulo e dependências
├── go.sum                      # Lockfile de dependências
├── README.md                  # Comandos de bootstrap do módulo + teste de conexão
├── cache/
│   └── redis.go               # Fábrica de cliente Redis
├── db/
│   └── postgres.go            # Fábrica de pool de conexão Postgres (pgxpool)
├── models/
│   └── tax_models.go          # Structs de domínio e DTOs de request/response
└── repository/
    └── tax_repository.go      # Acesso a dados: regras, configs e faixas de imposto
```

*(Excluídos da listagem: `.git`, `.reversa`, `_reversa_sdd`, `.claude`, `.agents` e artefatos de build.)*

---

## 3. Pacotes / Módulos identificados

| Pacote | Arquivo | Responsabilidade | Confiança |
|--------|---------|------------------|-----------|
| `models` | `models/tax_models.go` | Structs de domínio (`TaxDefinition`, `TaxRule`, `TaxCalculationLog`) e DTOs de entrada/saída (`TaxRequest`, `TaxResponse`, `UniversalTaxRequest`, `DeductionDetail`) | 🟢 |
| `repository` | `repository/tax_repository.go` | Camada de acesso a dados com cache: resolução de regra aplicável, leitura de configs e faixas progressivas | 🟢 |
| `db` | `db/postgres.go` | Conexão com PostgreSQL via `pgxpool` | 🟢 |
| `cache` | `cache/redis.go` | Conexão com Redis via `go-redis/v9` | 🟢 |

---

## 4. Pontos de entrada e configuração

| Tipo | Encontrado? | Detalhe |
|------|-------------|---------|
| Entry point executável (`main`) | ❌ Não | É uma biblioteca; nenhum `package main` 🟢 |
| Fábricas públicas (API da lib) | ✅ | `db.ConnectPostgres(connString)`, `cache.ConnectRedis(addr)`, `repository.NewTaxRepository(db, rdb)` 🟢 |
| Variáveis de ambiente | ✅ (via README) | `DATABASE_URL`, `REDIS_ADDR` 🟢 (`README.md`) |
| Arquivo `.env.example` | ❌ Não | 🔴 LACUNA — não há exemplo de env versionado |
| `config/` ou `settings` | ❌ Não | Configuração de imposto vive no **banco** (tabela `tax_configs`), não em arquivos 🟢 |
| CI/CD (`.github/workflows`, `Jenkinsfile`, `.gitlab-ci.yml`) | ❌ Não | Nenhum pipeline encontrado 🟢 |
| `Dockerfile` / `docker-compose.yml` | ❌ Não | Ausentes 🟢 |
| Script de teste citado no README | ⚠️ | `go run test_conn.go` referenciado, mas **`test_conn.go` não existe** no repositório 🔴 |

### Exemplo de `DATABASE_URL` (README)
```
postgres://worker_user:worker_pass@localhost:5432/worker_db?sslmode=disable&search_path=individual_tax_rates
```
> Indica schema dedicado **`individual_tax_rates`** no Postgres. 🟢

---

## 5. Banco de dados (análise superficial)

Não há arquivos DDL, migrations ou modelos ORM declarativos. O schema é inferido das **queries SQL embutidas** em `repository/tax_repository.go` e das structs em `models`. Análise detalhada caberá ao **Data Master**.

**Schema:** `individual_tax_rates` 🟢

| Tabela (referenciada em SQL) | Origem | Colunas observadas |
|------------------------------|--------|--------------------|
| `tax_configs` | `GetConfig`, `GetTableConfigs` | `tax_code`, `config_key`, `config_value`, `valid_from`, `valid_to` 🟢 |
| `tax_rules_history` | `GetTaxRulesForPeriod` | `id`, `tax_definition_id`, `description`, `range_min`, `range_max`, `aliq_percent`, `deduction_val`, `valid_from`, `valid_to` 🟢 |
| `tax_definitions` | `GetTaxRulesForPeriod` (JOIN) | `id`, `tax_code` (+ inferidas da struct: `name`, `sphere`, `rounding_precision`, `active`) 🟢/🟡 |

> 🟡 **INFERIDO:** A struct `TaxCalculationLog` (com `id uuid`, `trace_id`, `applied_aliq`) sugere uma tabela de log de cálculo (ex.: `tax_calculation_log`), mas **nenhuma query** a lê ou escreve neste repositório — a persistência do log deve ocorrer em outra camada ausente. 🔴 LACUNA

---

## 6. Testes

| Item | Resultado |
|------|-----------|
| Arquivos `*_test.go` | **0** 🟢 |
| Framework de teste declarado | Nenhum em uso direto; `go.sum` traz `stretchr/testify`, `bsm/ginkgo`+`gomega` como **dependências transitivas** (do `go-redis`), não testes próprios 🟡 |
| Cobertura estimada | **0%** — projeto sem suíte de testes 🟢 |
| Teste manual | `go run test_conn.go` citado no README, porém o arquivo está ausente 🔴 |

---

## 7. Lacunas e observações para os próximos agentes

- 🔴 **Recorte de monorepo:** comentários `// path: backend/go/libs/go-native/...` indicam que a lógica de cálculo (serviço que orquestra `TaxRequest` → `TaxResponse`, aplica deduções, escolhe entre cálculo `monthly`/`annual`) **não está neste repositório**. Os DTOs `TaxRequest`/`TaxResponse`/`UniversalTaxRequest` existem mas não há consumidor.
- 🔴 **`test_conn.go` ausente** apesar de citado no README.
- 🟡 **Struct não exportada `documentoFiscalRequest`** (minúscula) usada dentro de `UniversalTaxRequest` exportada — possível inconsistência ou design intencional para encapsular o item de input.
- 🟡 **Lógica de negócio no repositório:** `GetApplicableRule` faz filtragem de faixa em memória (não no SQL) após buscar todas as regras — decisão de performance/cache a ser documentada pelo Detective/Architect.
- 🟢 **Estratégia de cache:** Redis com TTL de 12h para `tax_configs` e `tax_rules_list`, chaveados por `taxCode` + data de referência.

---

## 8. Próximos passos sugeridos

1. **Decisão de organização das specs** (orquestrador Reversa).
2. **Archaeologist** — escavar pacote a pacote (`repository`, `models`, `db`, `cache`).
3. **Data Master** — reconstruir o schema `individual_tax_rates` a partir das queries.
4. **Detective** — extrair regras de negócio implícitas (faixas progressivas, vigência temporal, deduções IRPF).
