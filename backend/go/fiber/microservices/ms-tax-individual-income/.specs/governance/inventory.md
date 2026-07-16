# Inventário do Projeto — ms-tax-individual-income

Gerado pelo agente **Scout** em 2026-06-08. Atualizado em 2026-06-20.

## 📂 Estrutura Física do Código

```
.
├── main.go                         # Entry point: DI manual, setup Fiber, rotas (API + health)
├── go.mod                          # Módulo Go 1.25.6
├── go.sum                          # Checksums de dependências
├── test_conn.go                    # Script de diagnóstico PostgreSQL (build ignore)
├── .env                            # Variáveis de ambiente
├── handlers/
│   ├── tax_handler.go              # Handler HTTP: parse request, injetar trace ID
│   └── health_handler.go           # Health checks: liveness (/healthz), readiness (/health)
├── services/
│   ├── calculation_service.go      # Core: cálculo paralelo, deduções, Reforma 2026
│   └── inss_client.go              # Cliente HTTP para microserviço INSS (timeout 5s)
├── data/
│   └── init.sql                    # Schema + seed data (tax_definitions, rules, configs)
├── .specs/                         # Documentação de especificações
│   └── skill-output/               # Logs de implementação (cache Redis, health check)
├── .agents/                        # Agentes Reversa
├── .claude/                        # Configurações Claude
├── .reversa/                       # Artefatos Reversa
├── _reversa_sdd/                   # Specs SDD Reversa
├── CLAUDE.md                       # Instruções do agente Claude
├── AGENTS.md                       # Instruções para agentes AI
├── GEMINI.md                       # Instruções para Gemini
└── README-exemplos-de-uso.md       # Exemplos de curl para teste da API
```

## 🛠️ Tecnologias e Frameworks

- **Linguagem:** Go 1.25.6
- **Framework Web:** Fiber v2.52.12
- **Banco de Dados:** PostgreSQL via driver `pgx` v5.9.1 (schema: `individual_tax_rates`)
- **Cache:** Redis via `go-redis` v9.18.0 com nil-safety, logging e TTL configurável (`TAX_CACHE_TTL_HOURS`)
- **Matemática Financeira:** `shopspring/decimal` v1.4.0
- **Identificadores:** `google/uuid` v1.6.0
- **Lib Local:** `taxnexus-individual-core-lib` (replace → `../../../libs/go-native/taxnexus-individual-core-lib`)
- **Logging:** `log/slog` (stdlib) com handler JSON + nível Debug

## 📊 Cobertura de Testes

- **Testes unitários:** 0 arquivos `*_test.go`
- **Testes de integração:** 0
- **Cobertura:** 0%
- **Risco:** 🔴 Crítico para serviço de cálculo financeiro

## 🌐 Endpoints Expostos

| Método | Path | Handler | Descrição |
|--------|------|---------|-----------|
| GET | `/healthz` | HealthHandler.Liveness | Liveness probe (sempre 200) |
| GET | `/health` | HealthHandler.Readiness | Readiness probe (PostgreSQL + Redis) |
| GET | `/api/v1/health` | HealthHandler.Readiness | Readiness probe (grupo API) |
| POST | `/api/v1/calculate/irpf` | TaxHandler.CalculateIRPF | Cálculo de IRPF |

## 📁 Arquivos de Documentação (.specs/)

| Arquivo | Status | Descrição |
|---------|--------|-----------|
| INDEX.md | ✅ Completo | Mapa centralizador |
| architecture/architecture.md | ✅ Atualizado | Visão arquitetural (cache + health check) |
| architecture/c4-context.md | ✅ Completo | Diagrama de contexto |
| architecture/integrations.md | ✅ Atualizado | Integrações e dependências |
| architecture/erd.md | ✅ Novo | Modelo de dados completo |
| engineering/code-analysis.md | ✅ Completo | Análise de handlers/services |
| product/requirements.md | ✅ Atualizado | RF-01 a RF-06 + RNFs |
| product/feature-roadmap.md | ✅ Atualizado | Roadmap (cache e health check concluídos) |
| api/tax-api.yaml | ✅ Atualizado | Contrato OpenAPI 3.0.3 (inclui /health) |
| domain/domain.md | ✅ Atualizado | Regras de negócio + INSS |
| domain/reform-2026.md | ✅ Completo | Mecanismo de transição |
| governance/inventory.md | ✅ Atualizado | Este arquivo |
| governance/confidence-report.md | ✅ Atualizado | Score 🟢 92% |
| questions/questions_01.md | ✅ Completo | Lacunas resolvidas |
| skill-output/2026-06-20_174207.md | ✅ Log | Cache Redis |
| skill-output/2026-06-20_182911.md | ✅ Log | Health check endpoints |
