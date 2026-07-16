# Inventário do Projeto — ms-billing-engine-tax-rates

Gerado pelo agente **Spec Miner** em 2026-06-20. Atualizado em 2026-06-22 (C-001 pipeline SOP-013 7-fases + CBS/IBS split + todas as 9 features concluídas).

## Estrutura Física do Código

```
.
├── cmd/
│   ├── api/main.go                  # Entry point: servidor HTTP Fiber (porta via PORT env, default :3000)
│   └── test_engine/main.go          # CLI test harness manual (output JSON)
├── internal/
│   ├── domain/
│   │   └── domain.go                # Interface TaxCalculator (camada mais interna — DDD)
│   ├── calculator/
│   │   ├── engine.go                # Motor multi-fase SOP-013 7-fases (C-001)
│   │   ├── engine_test.go           # Testes do motor (6 cenários legados)
│   │   ├── pipeline_test.go         # Testes do pipeline SOP-013 (22 cenários)
│   │   └── legacy_adapter.go        # Adaptadores ICMS e PIS/COFINS → domain.TaxCalculator
│   ├── middleware/
│   │   ├── requestid.go             # Middleware Request-ID + Trace-ID (W3C Trace Context)
│   │   ├── requestid_test.go        # Testes do middleware requestid (12 cenários)
│   │   ├── auth.go                  # Middleware JWT Auth (Kong/Keycloak pass-through)
│   │   ├── auth_test.go             # Testes do middleware auth (9 cenários)
│   │   └── metrics.go               # Middleware + handler de métricas Prometheus
│   ├── phase/
│   │   ├── phase.go                 # PhaseResolver — SHADOW_RUN/CBS_PLENA/TRANSICAO/IVA_DUAL
│   │   ├── tax_selector.go          # TaxSelector com matriz DT-001
│   │   └── phase_test.go            # Testes do PhaseResolver (5 cenários)
│   ├── reforma/
│   │   ├── reforma.go                # Lógica compartilhada + ReformaCalculator (legado)
│   │   ├── cbs_calculator.go         # CBSCalculator — CBS apenas (Fase 2 sequencial)
│   │   ├── ibs_calculator.go         # IBSCalculator — IBS apenas (Fase 4 paralela)
│   │   └── reforma_test.go           # Testes da Reforma (7 cenários legados)
│   ├── circuitbreaker/
│   │   ├── circuit_breaker.go        # Circuit Breaker (CLOSED→OPEN→HALF_OPEN)
│   │   └── circuit_breaker_test.go   # Testes Circuit Breaker (7 cenários)
│   └── ibsclient/
│       ├── client.go                 # IBS Client (HTTP + Cache Redis + Fallback DB)
│       └── client_test.go            # Testes IBS Client (5 cenários)
└── internal/
    └── legacy/
│       ├── icms.go                  # Calculadora ICMS (normal, ST, DIFAL, Simples)
│       ├── icms_calculate_test.go   # Testes ICMS (12 cenários)
│       ├── icms_desoneracao.go      # ICMS Desonerado (F-004 — Redução Base + Limitação Alíquota)
│       ├── icms_desoneracao_test.go # Testes ICMS Desonerado (14 cenários)
│       ├── ipi.go                   # Calculadora IPI (Ad Valorem, Ad Pauta, rateio)
│       ├── ipi_calculate_test.go    # Testes IPI (7 cenários)
│       ├── is_filter.go             # ISFilter — pré-filtro IS (F-006, Fase 0)
│       ├── is_filter_test.go        # Testes ISFilter (8 cenários)
│       ├── iss.go                   # Calculadora ISS (serviços, LC 116/2003)
│       ├── iss_test.go              # Testes ISS (7 cenários)
│       ├── telecom.go               # Classificador SCM/STFC/SVA (FUST/FUNTTEL)
│       ├── fust.go                  # Calculadora FUST (1%, Lei 9.998/2000)
│       ├── fust_test.go             # Testes FUST (6 cenários)
│       ├── funttel.go               # Calculadora FUNTTEL (0,5%, Lei 10.052/2000)
│       ├── funttel_test.go          # Testes FUNTTEL (4 cenários)
│       ├── pis_cofins.go            # Calculadora PIS/COFINS
│       ├── pis_strategies.go        # Strategy por CST PIS (01-06, 49, 50-99, 99)
│       ├── pis_strategies_test.go    # Testes unitários PIS (13 testes)
│       ├── cofins_strategies.go     # Strategy por CST COFINS (01-06, 49, 50-99)
│       ├── cofins_strategies_test.go # Testes unitários COFINS (13 testes)
│       ├── pis_cofins_calculate_test.go # Testes integração PISCofinsCalculator (13 cenários)
│       └── mock_repository_test.go   # Mock de TaxRepository para testes
├── data/
│   └── init.sql                     # Schema DDL + triggers (7 tabelas)
├── docs/
│   ├── README-BRAINSTORM.md         # Brainstorming de regras fiscais
│   ├── README-CONSTANTS.md          # Constantes e valores de referência
│   ├── README-ESCOPO.md             # Definição de escopo do projeto
│   ├── README-ESCOPO_ADENDO.md      # Adendo ao escopo
│   ├── README-ICMS.md               # Documentação de regras ICMS
│   ├── README-ICMS-EXTENSAO-REGRAS.md # Extensão de regras ICMS
│   ├── README-ICSM-TAXA-DESONERACAO.md # Taxa de desoneração ICMS
│   ├── README-IPI.md                # Documentação de regras IPI
│   ├── README-PIS-COFINS.md         # Documentação de regras PIS/COFINS
│   ├── README-PIS-COFINS-ADENDO.md  # Adendo PIS/COFINS
│   ├── README-PIS-COFINS-DESONERACAO.md # Desoneração PIS/COFINS
│   ├── README-SIMPLES-NACIONAL.md   # Documentação Simples Nacional
│   └── README-TABELA-CST-CSON.md    # Tabela de códigos CST/CSOSN
├── .remember/
│   ├── .gitignore
│   └── logs/
│       ├── hook-errors.log
│       └── memory-2026-06-20.log
├── .specs/                          # Documentação de especificações (este diretório)
├── go.mod                           # Módulo Go 1.25.6
├── go.sum                           # Checksums de dependências
└── README.md                        # README principal do projeto
```

## Tecnologias e Frameworks

- **Linguagem:** Go 1.25.6
- **Framework Web:** Fiber v2.52.12 (fasthttp-based)
- **Banco de Dados:** PostgreSQL via driver `pgx` v5 (schema: `billing_tax_rates`)
- **Cache:** Redis via `go-redis` v9
- **Matemática Financeira:** `shopspring/decimal` v1.3.1
- **Validação:** `go-playground/validator/v10`
- **Identificadores:** `google/uuid`
- **Lib Local:** `taxnexus-billing-core-lib` (replace → `../../../libs/go-native/taxnexus-billing-core-lib`)
- **Observabilidade:** OpenTelemetry (`go.opentelemetry.io/otel`) + W3C Trace Context
- **Logging:** `log/slog` (stdlib) com handler JSON + nível Debug

## Cobertura de Testes

- **Testes unitários:** 18 arquivos `*_test.go`, 150+ testes
- **Testes de integração:** `pis_cofins_calculate_test.go` (13 cenários com mock repository)
- **Cobertura:** Alta — todas as calculadoras, middleware, reforma, pipeline e circuit breaker cobertos: PIS/COFINS 39, IPI 7, ICMS 12, Engine 6, Pipeline SOP-013 22, Middleware requestid 12, Middleware auth 9, Reforma 7, ISS 7, FUST 6, FUNTTEL 4, ICMS Desonerado 14, PhaseResolver 5, ISFilter 8, CircuitBreaker 7, IBSClient 5
- **Mock repository:** `mock_repository_test.go` implementa `repository.TaxRepository`
- **Test harness manual:** `cmd/test_engine/main.go`
- **Risco:**   Baixo para motor de cálculo fiscal (cobertura de testes adequada — cada fase do pipeline SOP-013 coberta)

## Endpoints Expostos

| Método | Path | Handler | Descrição |
|--------|------|---------|-----------|
| POST | `/calculate` | inline handler | Cálculo de tributos sobre documento fiscal |
| GET | `/healthz` | inline handler | Liveness probe (Kubernetes) |
| GET | `/health` | inline handler | Readiness probe (PostgreSQL + Redis) |
| GET | `/metrics` | `metrics.Handler()` | Métricas Prometheus (text exposition format) |

## Arquivos de Documentação (.specs/)

| Arquivo | Status | Descrição |
|---------|--------|-----------|
| INDEX.md |   Completo | Mapa centralizador |
| architecture/architecture.md |   Completo | Visão arquitetural |
| architecture/c4-context.md |   Completo | Diagrama de contexto |
| architecture/integrations.md |   Completo | Integrações e dependências |
| architecture/erd.md |   Completo | Modelo de dados (7 tabelas) |
| engineering/code-analysis.md |   Completo | Análise de módulos internos |
| engineering/api-guidelines.md |   Completo | Padrões de API e erros |
| product/requirements.md |   Completo | Requisitos funcionais |
| product/feature-roadmap.md |   Completo | Roadmap de features |
| api/tax-rates-api.yaml |   Completo | Contrato OpenAPI 3.0.3 |
| domain/domain.md |   Completo | Regras de negócio e glossário |
| governance/inventory.md |   Completo | Este arquivo |
| governance/confidence-report.md |   Completo | Score   95% |
| questions/questions_01.md |   Completo | Lacunas e dúvidas |
