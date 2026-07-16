# MS Billing Engine — Tax Rates

**Motor de cálculo de tributos sobre faturamento para o ecossistema TaxNexus.**

Microsserviço Go/Fiber que calcula tributos brasileiros (IPI, ICMS, PIS, COFINS, ISS, CBS, IBS, IS, FUST, FUNTTEL) sobre documentos fiscais, implementando o pipeline SOP-013 de 7 fases com arquitetura multi-fase genérica, incluindo regras da Reforma Tributária (EC 132/2023).

[![Go](https://img.shields.io/badge/Go-1.25.6-00ADD8?logo=go)](https://go.dev/)
[![Fiber](https://img.shields.io/badge/Fiber-v2.52.12-00ACC1?logo=go)](https://gofiber.io/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-pgx_v5-336791?logo=postgresql)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-go--redis_v9-DC382D?logo=redis)](https://redis.io/)
[![Tests](https://img.shields.io/badge/tests-150+-success)](./internal/)
[![Docs Confidence](https://img.shields.io/badge/docs_confidence-99%25-brightgreen)](./.specs/governance/confidence-report.md)

---

##   Visão Geral

O **MS Billing Engine Tax Rates** é o componente central de cálculo tributário da plataforma TaxNexus. Ele recebe um documento fiscal (`POST /calculate`) com itens de faturamento (NCM, CFOP, valores, UFs origem/destino) e retorna todos os tributos incidentes calculados conforme a legislação brasileira vigente.

### Tributos Suportados

| Tributo | Sigla | Esfera | Status |
|---------|-------|--------|--------|
| Imposto sobre Produtos Industrializados | IPI | Federal |   Implementado |
| Imposto sobre Circulação de Mercadorias e Serviços | ICMS | Estadual |   Implementado |
| ICMS Substituição Tributária | ICMS-ST | Estadual |   Implementado |
| Diferencial de Alíquotas (EC 87/2015) | DIFAL | Estadual |   Implementado |
| Programa de Integração Social | PIS | Federal |   Implementado |
| Contribuição para Financiamento da Seguridade Social | COFINS | Federal |   Implementado |
| Imposto Sobre Serviços (LC 116/2003) | ISS | Municipal |   Implementado |
| Fundo de Universalização das Telecomunicações | FUST | Federal |   Implementado |
| Fundo para Desenvolvimento Tecnológico das Telecom | FUNTTEL | Federal |   Implementado |
| Contribuição sobre Bens e Serviços (Reforma) | CBS | Federal |   Implementado |
| Imposto sobre Bens e Serviços (Reforma) | IBS | Subnacional |   Implementado |
| Imposto Seletivo (Reforma) | IS | Federal |   Implementado |

---

##   Arquitetura

### Diagrama de Contexto (C4 — Nível 1)

```mermaid
flowchart LR
    consumidor["Consumidor\nSistema de faturamento"]
    ms_tax["MS Billing Engine Tax Rates\nMotor de cálculo fiscal (Go/Fiber)\nPorta :3000 (configurável via PORT)"]
    pg[("PostgreSQL\nSchema: billing_tax_rates\nRegras fiscais")]
    redis[("Redis\nCache de regras fiscais")]

    consumidor -->|"POST /calculate\nJSON/HTTPS\nJWT (Kong/Keycloak)"| ms_tax
    ms_tax -->|"pgx v5\nConsultas de regras"| pg
    ms_tax -->|"go-redis v9\nCache com TTL"| redis
```

### Pipeline SOP-013 — 7 Fases

O motor executa os tributos em 7 fases porque há **dependências assimétricas** entre eles. Ex: o IPI compõe a base do ICMS, e o ICMS deve ser excluído da base do PIS/COFINS ("Tese do Século", STF).

```
F0 (Seq) → IS          ← pré-filtro (NCM seletivo)
  └─ F1 (Seq) → IPI         ← compõe base do ICMS
      └─ F2 (Seq) → CBS         ← "por fora", não compõe base de outros
          └─ F3 (Seq) → ICMS        ← antes do PIS/COFINS
              ├─ F4 (Par) → IBS         ← subnacional (estadual + municipal)
              ├─ F4 (Par) → ISS         ← municipal sobre serviços
              └─ F4 (Par) → PIS/COFINS  ← com exclusão do ICMS
                  └─ F5 (Seq) → FUST        ← depende ICMS+PIS+COFINS
                      └─ F6 (Seq) → FUNTTEL    ← mesma base do FUST
```

A arquitetura usa `CalculationPhase` genérica com modos `Sequential` e `Parallel`. Após cada fase, `injectTributoValues()` injeta os valores calculados nos detalhes do input para consumo pelas fases seguintes.

### Camadas (DDD)

```
internal/calculator/      ← Motor de orquestração multi-fase
    ↓
internal/domain/         ← Interface TaxCalculator (camada mais interna)
internal/legacy/         ← Calculadoras fiscais + Strategy Pattern (PIS/COFINS)
internal/reforma/        ← Reforma Tributária (CBS, IBS)
internal/phase/          ← Phase Resolver + TaxSelector (fases da Reforma)
internal/circuitbreaker/ ← Circuit Breaker para API IBS
internal/ibsclient/      ← IBS Client com cache Redis
    ↓
repository (core-lib)    ← Acesso a dados + cache (taxnexus-billing-core-lib)
    ↓
PostgreSQL + Redis       ← Persistência e cache
```

---

##   Tecnologias

| Componente | Tecnologia | Versão |
|-----------|-----------|--------|
| Linguagem | Go | 1.25.6 |
| Framework HTTP | Fiber (fasthttp-based) | v2.52.12 |
| Banco de Dados | PostgreSQL (pgx) | v5 |
| Cache | Redis (go-redis) | v9 |
| Matemática Financeira | shopspring/decimal | v1.3.1 |
| Validação | go-playground/validator | v10 |
| Identificadores | google/uuid | v1.6.0 |
| Observabilidade | OpenTelemetry + W3C Trace Context | v1.44.0 |
| Logging | log/slog (stdlib) | Go 1.21+ |
| Lib Local | taxnexus-billing-core-lib | replace → `../../../libs/go-native/taxnexus-billing-core-lib` |

---

##   Primeiros Passos

### Pré-requisitos

- Go 1.25+
- PostgreSQL 15+
- Redis 7+

> **Dica:** Se não tiver PostgreSQL e Redis instalados localmente, use o Docker Compose do projeto (veja abaixo).

### Ambiente de Testes com Docker Compose  

O projeto disponibiliza um manifesto Docker Compose que sobe PostgreSQL 16 + Redis 7 prontos para desenvolvimento local:

**Arquivo:** `utils/docker-compose/docker-compose-db-worker-postgres.yml`  
**Caminho absoluto:** `/home/bolismar/work/workspace-fbso/utils/docker-compose/docker-compose-db-worker-postgres.yml`

```bash
# Subir banco e cache (background)
docker compose -f /home/bolismar/work/workspace-fbso/utils/docker-compose/docker-compose-db-worker-postgres.yml up -d

# Verificar se os containers subiram
docker ps --filter "name=db-pg-worker" --filter "name=redis-tax-rules"

# Parar os containers quando terminar
docker compose -f /home/bolismar/work/workspace-fbso/utils/docker-compose/docker-compose-db-worker-postgres.yml down
```

**Serviços provisionados:**

| Serviço | Container | Porta | Credenciais |
|---------|-----------|-------|-------------|
| PostgreSQL 16 | `db-pg-worker` | `5432` | `worker_user` / `worker_pass` / `worker_db` |
| Redis 7 | `redis-tax-rules` | `6379` | — (sem senha) |

### Configuração

```bash
# Obrigatório — ajuste conforme seu ambiente
# Usando Docker Compose (valores do manifesto acima):
export DATABASE_URL="postgres://worker_user:worker_pass@localhost:5432/worker_db?sslmode=disable"
export REDIS_ADDR="localhost:6379"

# Ou, se tiver PostgreSQL/Redis próprios:
export DATABASE_URL="postgres://user:pass@localhost:5432/billing_tax_rates?sslmode=disable"
export REDIS_ADDR="localhost:6379"

# Opcional
export PORT=":3000"                                    # default: :3000
export IBS_API_BASE_URL="https://api.comitegestoribs.gov.br"  # API do Comitê Gestor IBS
```

### Instalação e Execução

```bash
# Instalar dependências
go mod tidy

# Executar o schema SQL (tabelas de regras fiscais)
psql "$DATABASE_URL" -f data/init.sql
# Se estiver usando Docker Compose:
# docker exec -i db-pg-worker psql -U worker_user -d worker_db < data/init.sql

# Compilar e rodar
go run cmd/api/main.go

# Testes
go test ./...
```

### Health Check

```bash
# Liveness (Kubernetes)
curl http://localhost:3000/healthz

# Readiness (verifica PostgreSQL + Redis)
curl http://localhost:3000/health

# Métricas Prometheus
curl http://localhost:3000/metrics
```

---

##   Deploy (Docker & Kubernetes)

O projeto inclui artefatos de deploy prontos para produção (GAP-010).

### Docker

```bash
# Build (contexto = raiz do workspace fbso/)
docker build \
  -f backend/go/fiber/microservices/ms-billing-engine-tax-rates/Dockerfile \
  -t ms-tax-rates:latest \
  .

# Executar (necessita PG + Redis — veja docker-compose abaixo)
docker run -p 3000:3000 \
  -e DATABASE_URL="postgres://user:pass@host:5432/db" \
  -e REDIS_ADDR="host:6379" \
  ms-tax-rates:latest
```

### Docker Compose (Ambiente Local Completo)

```bash
cd backend/go/fiber/microservices/ms-billing-engine-tax-rates

# Iniciar app + PostgreSQL 16 + Redis 7
docker compose up -d

# Verificar status
docker compose ps

# Logs da aplicação
docker compose logs -f app

# Smoke test
curl -X POST http://localhost:3000/v1/calculate \
  -H "Content-Type: application/json" \
  -d '{"itens":[{"sku":"TEST","ncm":"84713012","valor_unitario":"1000","quantidade":"1"}]}'

# Parar
docker compose down
```

### Kubernetes

```bash
# Aplicar manifests (ordem recomendada)
kubectl apply -f deploy/k8s/configmap.yaml
kubectl apply -f deploy/k8s/deployment.yaml
kubectl apply -f deploy/k8s/service.yaml
kubectl apply -f deploy/k8s/hpa.yaml

# Verificar rollout
kubectl rollout status deployment/ms-tax-rates -n tax-engine

# Rollback
kubectl rollout undo deployment/ms-tax-rates -n tax-engine
```

### Variáveis de Ambiente de Deploy

| Variável | Uso | Default |
|:---|:---|:---|
| `DATABASE_URL` | Conexão PostgreSQL | — (obrigatório) |
| `REDIS_ADDR` | Endereço Redis | — (obrigatório) |
| `PORT` | Porta HTTP | `:3000` |
| `IBS_API_BASE_URL` | API Comitê Gestor IBS | `https://api.comitegestoribs.gov.br` |
| `RATE_LIMIT_MAX` | Máximo req/janela | `100` |
| `RATE_LIMIT_WINDOW` | Janela em segundos | `60` |
| `TAX_TOKEN_TTL_MINUTES` | TTL token fiscal | — |

---

##   Endpoints

| Método | Path | Descrição |
|--------|------|-----------|
| `POST` | `/v1/calculate` | Cálculo de tributos sobre documento fiscal |
| `POST` | `/calculate` | ⚠️ Deprecado — use `/v1/calculate` |
| `GET` | `/v1/healthz` | Liveness probe (Kubernetes) |
| `GET` | `/v1/health` | Readiness probe (PostgreSQL + Redis) |
| `GET` | `/v1/metrics` | Métricas Prometheus (text exposition) |

### POST /calculate

**Request:**
```json
{
  "IDTransaction": "uuid-gerado-pelo-consumidor",
  "CFOP": "5101",
  "CRT": 3,
  "UFOrigem": "SP",
  "UFDestino": "RJ",
  "Itens": [
    {
      "SKU": "PROD-001",
      "NCM": "85176210",
      "CFOP": "5101",
      "Valor": "1000.00",
      "Quantidade": 1,
      "CST": "00",
      "UFOrigem": "SP",
      "UFDestino": "RJ"
    }
  ]
}
```

**Response `200`:**
```json
{
  "IDTransaction": "uuid-gerado-pelo-servico",
  "Itens": [
    {
      "SKU": "PROD-001",
      "Total": "1000.00",
      "Tributos": [
        {
          "Tributo": "IPI",
          "CST": "50",
          "BaseCalculo": "1050.00",
          "Aliquota": "5.00",
          "Valor": "52.50"
        }
      ]
    }
  ]
}
```

---

##  Estrutura do Projeto

```
.
├── cmd/
│   ├── api/main.go                  # Entry point — servidor Fiber
│   └── test_engine/main.go          # CLI test harness (output JSON)
├── internal/
│   ├── domain/domain.go             # Interface TaxCalculator (DDD)
│   ├── calculator/                  # Motor multi-fase SOP-013 (C-001)
│   │   ├── engine.go                # BillingEnginePhased — 7 fases
│   │   ├── pipeline_test.go         # 22 cenários de teste do pipeline
│   │   └── legacy_adapter.go        # Adapter: legacy → domain.TaxCalculator
│   ├── legacy/                      # Calculadoras fiscais
│   │   ├── icms.go                  # ICMS (normal, ST, DIFAL, Simples)
│   │   ├── icms_desoneracao.go      # ICMS Desonerado (F-004)
│   │   ├── ipi.go                   # IPI (Ad Valorem, Ad Pauta)
│   │   ├── pis_cofins.go            # PIS/COFINS
│   │   ├── pis_strategies.go        # Strategy Pattern PIS (15 estratégias)
│   │   ├── cofins_strategies.go     # Strategy Pattern COFINS
│   │   ├── iss.go                   # ISS (LC 116/2003)
│   │   ├── fust.go                  # FUST (Lei 9.998/2000)
│   │   ├── funttel.go               # FUNTTEL (Lei 10.052/2000)
│   │   ├── telecom.go               # Classificador SCM/STFC/SVA
│   │   └── is_filter.go             # IS pré-filtro (F-006)
│   ├── reforma/                     # Reforma Tributária (EC 132/2023)
│   │   ├── cbs_calculator.go        # CBS (Fase 2)
│   │   └── ibs_calculator.go        # IBS (Fase 4)
│   ├── phase/                       # Phase Resolver + TaxSelector (F-005)
│   ├── circuitbreaker/              # Circuit Breaker (F-007)
│   ├── ibsclient/                   # IBS Client HTTP + Cache
│   └── middleware/
│       ├── requestid.go             # W3C Trace Context
│       ├── auth.go                  # JWT pass-through (Kong/Keycloak)
│       └── metrics.go               # Prometheus metrics
├── data/init.sql                    # Schema DDL (10 tabelas)
├── docs/                            # Documentação de regras fiscais (13 arquivos)
└── .specs/                          # Especificações de engenharia (22 arquivos)
```

---

##   Documentação Técnica

> **IMPORTANTE — Entry point para agentes de IA e times técnicos.**

Toda a documentação técnica está nos diretórios `.specs/` e `docs/`. Comece pelo mapa centralizador:

###   Entry Point: [`.specs/INDEX.md`](./.specs/INDEX.md)

O `INDEX.md` é o **hub canônico de documentação** — mapeia cada aspecto do projeto para seu arquivo correspondente.

###   Especificações de Engenharia (`.specs/`)

| Categoria | Arquivo | Descrição |
|-----------|---------|-----------|
|   Arquitetura | [`architecture/architecture.md`](./.specs/architecture/architecture.md) | Visão arquitetural, patterns, middleware pipeline, camadas |
|   Diagrama C4 | [`architecture/c4-context.md`](./.specs/architecture/c4-context.md) | Diagrama de contexto (Mermaid) |
|   Modelo de Dados | [`architecture/erd.md`](./.specs/architecture/erd.md) | ERD completo (10 tabelas) |
|   Integrações | [`architecture/integrations.md`](./.specs/architecture/integrations.md) | Dependências, env vars, lib local, comunicação externa |
|   Domínio | [`domain/domain.md`](./.specs/domain/domain.md) | Regras de negócio, glossário fiscal |
|   Análise de Código | [`engineering/code-analysis.md`](./.specs/engineering/code-analysis.md) | Análise de fluxo handlers/services |
|   API Guidelines | [`engineering/api-guidelines.md`](./.specs/engineering/api-guidelines.md) | Padrões de API, erros, observabilidade |
|   Requisitos | [`product/requirements.md`](./.specs/product/requirements.md) | RF-01 a RF-10, RNF-01 a RNF-15 (MoSCoW) |
|   Roadmap | [`product/feature-roadmap.md`](./.specs/product/feature-roadmap.md) | Features implementadas, backlog, dívidas técnicas |
|   Contrato API | [`api/tax-rates-api.yaml`](./.specs/api/tax-rates-api.yaml) | OpenAPI 3.0.3 — rotas, schemas, erros |
|   Inventário | [`governance/inventory.md`](./.specs/governance/inventory.md) | Estrutura de código, cobertura de testes |
|   Confiança | [`governance/confidence-report.md`](./.specs/governance/confidence-report.md) | Score de confiança da documentação: **99%**   |
|   Histórico | [`skill-output/`](./.specs/skill-output/) | 15 registros de sessão de implementação |

###   Regras de Negócio (`.docs/`)

| Arquivo | Conteúdo |
|---------|----------|
| [`README-ESCOPO.md`](./docs/README-ESCOPO.md) | Definição de escopo do projeto |
| [`README-ICMS.md`](./docs/README-ICMS.md) | Regras detalhadas de ICMS |
| [`README-IPI.md`](./docs/README-IPI.md) | Regras detalhadas de IPI |
| [`README-PIS-COFINS.md`](./docs/README-PIS-COFINS.md) | Regras detalhadas de PIS/COFINS |
| [`README-SIMPLES-NACIONAL.md`](./docs/README-SIMPLES-NACIONAL.md) | Documentação Simples Nacional |
| [`README-TABELA-CST-CSON.md`](./docs/README-TABELA-CST-CSON.md) | Tabela de códigos CST/CSOSN |
| [`README-CONSTANTS.md`](./docs/README-CONSTANTS.md) | Constantes e valores de referência |

---

##   Padrões de Código

### Strategy Pattern
PIS e COFINS usam **Strategy Pattern com seleção por CST** — cada código de situação tributária mapeia para uma estratégia de cálculo específica, permitindo extensão sem modificar a calculadora principal.

### Adapter Pattern
`LegacyAdapter` converte calculadoras do módulo `legacy` (com interfaces específicas) para a interface unificada `domain.TaxCalculator`.

### Injeção de Dependência Manual
Sem frameworks DI. A ordem de inicialização no `main.go` é explícita: PostgreSQL → Redis → Repository → Repository com cache → Calculadoras → Engine.

### Middleware Pipeline
```
Request → recover → requestid (W3C) → auth (JWT) → logger → metrics → Handler
```

---

##   Testes

- **150+ cenários de teste** em 18 arquivos `*_test.go`
- Cobertura de todas as calculadoras, middleware, engine, pipeline, circuit breaker e phase resolver
- Mock de `TaxRepository` disponível em `mock_repository_test.go`
- Test harness manual: `cmd/test_engine/main.go`

```bash
# Executar todos os testes
go test ./...

# Com coverage
go test -cover ./...

# Testes específicos
go test ./internal/calculator/ -run TestPipeline
```

---

##   Observabilidade

- **Logging:** `log/slog` (stdlib) com handler JSON e nível Debug — cada log inclui `trace_id` e `request_id`
- **Métricas:** Prometheus via `GET /metrics` — `http_requests_total`, `http_request_duration_seconds`, `cache_requests_total`, `errors_total`
- **Tracing:** W3C Trace Context — headers `traceparent`/`traceresponse` + `X-Request-ID`
- **Health:** `/healthz` (liveness) + `/health` (readiness — verifica PG + Redis)

---

##   Dívidas Técnicas

Consulte [`product/feature-roadmap.md`](./.specs/product/feature-roadmap.md#d%C3%ADvidas-t%C3%A9cnicas) para a lista canônica (DT-01 a DT-11). Principais itens em aberto:

| ID | Descrição | Prioridade |
|----|-----------|------------|
| DT-03 | CSTs provisórios da Reforma (aguardando tabela oficial RFB) | Média |
| DT-04 | Créditos da Reforma (cash forward) | Média |
| DT-09 | API do Comitê Gestor IBS não publicada (Gap G2) | Média |

---

##   Licença

Proprietário — TaxNexus. Uso interno.
