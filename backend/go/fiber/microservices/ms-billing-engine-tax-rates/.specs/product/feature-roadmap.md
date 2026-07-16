# Feature Roadmap — ms-billing-engine-tax-rates

Atualizado em 2026-06-30 — PR #6 merge (Fases 0-1-2 Reforma Tributária: Admin Fiscal, Créditos, TaxToken, Simulação, Fornecedores). Pipeline SOP-013 7-fases estável. Rate Limiting e API Versioning implementados (Fase 0). Docker/K8s deploy implementados (GAP-010). 211+ testes passando. 10 tabelas SQL + 3 índices.

## Features Implementadas

| Feature | Status | Descrição |
|---------|--------|-----------|
| Motor 7-fases SOP-013 (C-001) |   Completo | Pipeline F0–F6: IS→IPI→CBS→ICMS→(IBS+ISS+PISCOFINS)→FUST→FUNTTEL. Arquitetura multi-fase genérica com `CalculationPhase` (Sequential/Parallel). ICMS sequencial antes do PIS/COFINS (Tese do Século), CBS "por fora" sequencial antes do ICMS. ReformaCalculator dividido em CBSCalculator + IBSCalculator. Injeção inter-fase automática com chaves duplas para ICMS (ITEM_ICMS_VALOR + VALOR_EXCLUSAO_ICMS). 22 testes de pipeline. |
| Cálculo IPI |   Completo | Ad Valorem, Ad Pauta, rateio de despesas acessórias, 3-tier priority |
| Cálculo DIFAL |   Completo | EC 87/2015 — diferencial de alíquotas interestadual |
| Cálculo Simples Nacional |   Completo | Equivalência CSOSN→CST, alíquota efetiva por anexo/RBT12 |
| Cálculo ICMS Próprio |   Completo | `getEffectiveTaxConfig()` integrado, redução de base, FCP |
| Cálculo ICMS-ST regras gerais |   Completo | ST interno e interestadual com MVA, sem exceção de produto |
| PIS/COFINS via banco |   Completo | Alíquotas do `federal_tax_rules` com fallback para defaults (1.65%/7.6%) |
| Cobertura total de CSTs PIS/COFINS |   Completo | CSTs 01-06, 49, 50-99, 99 — 100% de cobertura (15 estratégias) |
| Exclusão ICMS da base PIS/COFINS |   Completo | Flag `ExcluiICMSBase` validada com 13 cenários de teste |
| Testes automatizados |   Completo | 25 arquivos de teste, 211+ cenários (PIS/COFINS 39, IPI 7, ICMS 12, Engine 6, Pipeline SOP-013 22, Middleware requestid 12, Middleware auth 9, Reforma 7, ISS 7, FUST 6, FUNTTEL 4, ICMS Desonerado 14, PhaseResolver 5, ISFilter 8, CircuitBreaker 7, IBSClient 5, Admin 4, Credit 4, Simulation 3, Supplier 4, Token 4) |
| Middleware requestid/traceid |   Completo | W3C Trace Context — Request-ID + Trace-ID com 12 testes |
| Health check endpoints |   Completo | `/healthz` (liveness) e `/health` (readiness com PG + Redis) |
| Coleta de erros em goroutines |   Completo | Channel de erro + `slog.Warn` na Fase 2 (não bloqueante) |
| Auth JWT (Kong/Keycloak) |   Completo | Decode JWT → `X-User-Id`, `X-User-Name`, `X-User-Roles` com 9 testes |
| Métricas Prometheus |   Completo | `GET /metrics` — 4 métricas (stdlib only, sem dependências externas) |
| Porta configurável |   Completo | `PORT` env var com default `:3000` |
| LegacyAdapter |   Completo | Conversão de calculadoras legacy para interface `TaxCalculator` unificada |
| Cache Redis |   Completo | Cache de regras via decorator pattern (`CachedTaxRepository`) |
| Documentação de regras |   Completo | Consolidada em `.specs/domain/domain.md` + 22 arquivos `.specs/` |
| Reforma Tributária (CBS/IBS/IS) |   Completo (2026-06-21 15:17) | CBS, IBS (estadual+municipal) e IS baseado em `iva_dual_rules` com 7 testes. Redução de alíquotas, isenção, imposto seletivo. `ReformaCalculator` implementa `TaxCalculator` diretamente — sem adapter. |
| Reorganização domain/ |   Completo (2026-06-21 18:49) | Extração da interface `TaxCalculator` para `internal/domain/` seguindo DDD. Pacote `domain` como camada mais interna (zero dependências internas). 5 arquivos atualizados, `go build`, `go vet` e `go test` passando. |
| Calculadora ISS |   Completo (2026-06-21 23:09) | `ISSCalculator` — alíquota [2%,5%], item 1.05 LC 116/2003, retenção fonte. 7 testes. |
| Calculadora FUST |   Completo (2026-06-21 23:09) | `FUSTCalculator` — base líquida (Valor−ICMS−PIS−COFINS), 1% (Lei 9.998/2000). 6 testes. |
| Calculadora FUNTTEL |   Completo (2026-06-21 23:09) | `FUNTTELCalculator` — mesma base FUST, 0,5% (Lei 10.052/2000). 4 testes. |
| TelecomClassifier |   Completo (2026-06-21 23:09) | Classificador SCM/STFC/SVA compartilhado entre FUST e FUNTTEL. `internal/legacy/telecom.go`. |
| Motor trifásico |   Completo (2026-06-21 23:09) | `BillingEngineFull()` — Fase 3 pós-paralela + injeção reversa ICMS/PIS/COFINS. `engine.go` adaptado. |
| Security: information disclosure |   Completo (2026-06-21 23:09) | Health check e bad request sanitizados — erros movidos para `slog`. |
| ICMS Desonerado (F-004) |   Completo (2026-06-21 23:59) | 2 modos: Redução de Base (BR-TAX-CALC-021) e Limitação de Alíquota (BR-TAX-CALC-022). CST validation (≠ 00). motDesICMS (1-12, 90). vICMSDeson abatido. FCP integrado. Simples Nacional excluído. 11 testes. Arquivos: `icms_desoneracao.go`, `icms_desoneracao_test.go`. |
| Phase Resolution System (F-005) |   Completo (2026-06-21 23:59) | `PhaseResolver` mapeia DataOperacao→Phase (SHADOW_RUN/2026, CBS_PLENA/2027-28, TRANSICAO_SUBNACIONAL/2029-32, IVA_DUAL/2033+). `TaxSelector` com matriz DT-001. `ProcessWithPhase()` no engine com shadow tax (não compõe total), redução subnacional, extinção IVA Dual. Integrado em `/calculate`. 14 testes. Arquivos: `phase/phase.go`, `phase/tax_selector.go`, `phase/phase_test.go`, `engine.go` (modificado). |
| IS Pré-Filtro (F-006) |   Completo (2026-06-22 00:43) | `ISFilter` como Fase 0 (pré-calculadora antes de IPI). Consulta tabela `ncm_seletivo` independente. Flag `isento_is` como override manual. IS NÃO sofre redução (ao contrário de CBS/IBS). `ReformaCalculator` refatorado: apenas CBS + IBS. 8 testes. Arquivos: `is_filter.go`, `is_filter_test.go`, `reforma.go` (modificado), `main.go` (modificado). |
| IBS Circuit Breaker (F-007) |   Completo (2026-06-22 00:43) | `IBSRateFetcher` interface com 4 implementações: `HTTPIBSClient` (API Comitê Gestor), `CachedIBSClient` (Redis TTL 24h), `CircuitBreakerIBSClient` (3 falhas/60s→OPEN, HALF_OPEN 5min), `FallbackIBSClient` (API→DB fallback). Integrado em `main.go` com env var `IBS_API_BASE_URL`. 7 testes circuit breaker + 5 testes IBS client. Arquivos: `circuitbreaker/`, `ibsclient/`. |
| Schema SQL (C-002) |   Completo (2026-06-22 00:43) | Tabelas criadas: `ncm_seletivo` (NCMs sujeitos ao IS), `cbs_rates` (alíquotas CBS por classe tributária), `iss_rates` (alíquotas ISS por município/IBGE). Dados de exemplo para categorias IS (bebidas, cigarros, refrigerantes) e 5 capitais para ISS. |
| Pipeline SOP-013 (C-001) |   Completo (2026-06-22 01:10) | Reordenação do pipeline de 3 para 7 fases conforme SOP-013. `CalculationPhase` genérico com modos Sequential/Parallel. CBS "por fora" (F2) antes do ICMS (F3). ICMS sequencial antes do PIS/COFINS (Tese do Século). IBS paralelo com ISS e PIS/COFINS (F4). ReformaCalculator dividido em CBSCalculator e IBSCalculator. `injectTributoValues()` com chaves duplas ICMS. 22 testes de pipeline (ordenação, concorrência, injeção, phase-aware). |

## Features Planejadas

| Feature | Prioridade | Descrição |
|---------|-----------|-----------|
| Rate Limiting | ~~Alta~~ ✅ Concluída (Fase 0) | Proteger endpoint `/calculate` contra abuso/DoS. Middleware Fiber com `RATE_LIMIT_MAX` e `RATE_LIMIT_WINDOW`. |
| API Versioning (`/v1/calculate`) | ~~Média~~ ✅ Concluída (Fase 0) | Prefixo `/v1/` adotado. Rotas legacy em `/calculate`, `/healthz`, `/metrics` com deprecation warning. |
| Deploy (Docker & Kubernetes) | ~~Média~~ ✅ Concluída (GAP-010) | Dockerfile multi-stage, docker-compose.yaml (app+PG+Redis), deploy/k8s/ (configmap, deployment, service, hpa). |
| Admin Fiscal (CRUD) | ~~Alta~~ ✅ Concluída (GAP-004) | `internal/admin/` — CRUD de regras fiscais (models, repository, service). `GET/POST /v1/admin/tax-rates/iva-dual`. |
| Créditos Reforma Tributária | ~~Alta~~ ✅ Concluída (GAP-005) | `internal/credit/` — engine de créditos, supplier. `POST /v1/credit/calculate`. |
| Simulação de Margem | ~~Média~~ ✅ Concluída (GAP-003) | `internal/simulation/` — projeção "what-if". `POST /v1/simulate`. |
| TaxToken Snapshot | ~~Média~~ ✅ Concluída | `internal/token/` — geração de token fiscal. `POST /v1/token/generate`. |
| Validação de Fornecedores | ~~Média~~ ✅ Concluída | `internal/supplier/` — models, service, store. `POST /v1/supplier/validate`, `GET /v1/supplier/:cnpj`. |
| Atualização da especificação OpenAPI | ~~Média~~ ✅ Concluída | `tax-rates-api.yaml` reescrito em 2026-06-21 (Spec Miner): v1.0.0 com schemas alinhados, 4 endpoints, auth JWT documentado, 11 tributos no enum |
| Payload Size Limits |   Média | Configurar `fiber.Config{BodyLimit: N}` para prevenir oversized payloads. Atualmente sem limite máximo de corpo de requisição. |
| Monitoring Dashboard Templates |   Baixa | Templates Grafana para as 4 métricas Prometheus expostas (`http_requests_total`, `http_request_duration_seconds`, `cache_requests_total`, `errors_total`) + alertas para taxa de erro > 1% e latência p95 > 500ms. |
| Performance Benchmark Suite |   Baixa | Suite de benchmark (`go test -bench`) para cálculos fiscais com SLAs documentados: latência esperada por item, throughput máximo (req/s), consumo de memória por cálculo. |
| ~~Reordenar Pipeline (C-001)~~ | ~~Alta~~ ✅ Concluída | Motor 7-fases SOP-013 implementado. `BillingEnginePhased()` com `CalculationPhase` genérico. CBS(F2)→ICMS(F3)→IBS+ISS+PISCOFINS(F4). `injectTributoValues()` automático entre fases. |
| ~~IS como Pré-Filtro (F-006)~~ | ~~Média~~ ✅ Concluída | `ISFilter` implementado como Fase 0. `ncm_seletivo` populado. `ReformaCalculator` refatorado (CBS+IBS apenas). |
| ~~Expandir Schema SQL (C-002)~~ | ~~Média~~ ✅ Concluída | Tabelas `ncm_seletivo`, `cbs_rates`, `iss_rates` criadas em `data/init.sql` com dados de exemplo. |
| ~~IBS Circuit Breaker (F-007)~~ | ~~Baixa~~ ✅ Concluída | `IBSRateFetcher` com HTTP+Redis cache+CircuitBreaker+FallbackDB. ⚠️ API real bloqueada por Gap G2 — fallback DB ativo. |

## Dívidas Técnicas

| ID | Descrição | Prioridade | Local |
|----|-----------|-----------|-------|
| DT-01 | `IDTransaction: "0"` placeholder — substituir por geração real (UUID) | ~~Baixa~~ ✅ Resolvida | `engine.go:49`, `icms.go:26` (`github.com/google/uuid`) |
| DT-02 | `TODO: VERIFICAR SE PRECISAR SER O CRT DO EMITENTE OU DESTINATARIO` — resolvido silenciosamente: código agora usa `CRTEmitente` (ver `icms.go:31,369`). Marcador TODO removido durante reestruturação do fluxo ICMS. | ~~Baixa~~ ✅ Resolvida | `icms.go:31,369` (`models.NormalizeCRTEmitente(input.CRTEmitente)`) |
| DT-03 | CST da Reforma Tributária usa valores provisórios (`01`/`04`) — aguardando tabela oficial da RFB | ~~Média~~ ✅ Resolvida (2026-07-01) | Tabela `cst_reforma` com 164 CCTs oficiais, `GetCSTReforma()` |
| DT-04 | Integração com créditos (cash forward / `permite_credito_amplo`) não implementada na Reforma |   Média | `internal/reforma/reforma.go` |
| DT-05 | CI/CD pipeline não documentado/configurado para este microserviço Go |   Baixa | — |
| DT-06 | Estrutura `legacy/` ainda contém lógica de negócio misturada com acesso a dados — separar em handlers/services/repository |   Baixa | `internal/legacy/` |
| DT-07 | Especificação OpenAPI (`tax-rates-api.yaml`) estava desatualizada — reescrita em 2026-06-21 para v1.0.0: schemas alinhados com modelos reais, 4 endpoints documentados, enum de tributos completo, auth documentado | ~~Média~~ ✅ Resolvida | `.specs/api/tax-rates-api.yaml` |
| DT-08 | Nome do pacote `internal/legacy/` é enganoso — contém as calculadoras principais (ICMS, IPI, PIS/COFINS), não código legado. Renomear para `internal/taxes/` ou `internal/calculators/` alinharia com a semântica real. |   Baixa | `internal/legacy/` (7 arquivos .go) |
| DT-09 | `middleware.InitTracing()` (OpenTelemetry W3C Trace Context) chamado em `main.go:33` não está documentado nos diagramas de arquitetura (`architecture.md` e `c4-context.md`). O middleware pipeline documentado omite a inicialização do propagador OTEL. |   Baixa | `cmd/api/main.go:33`, `internal/middleware/requestid.go:13` |
| DT-10 | Ausência de artefatos de deploy para produção: sem Dockerfile, sem manifests Kubernetes (Deployment/Service/ConfigMap), sem configuração de resource limits/requests. | ~~Média~~ ✅ Resolvida (GAP-010) | `Dockerfile`, `docker-compose.yaml`, `deploy/k8s/` (configmap, deployment, service, hpa) |
| DT-11 | Sem rate limiting no endpoint `/calculate` — risco de exaustão de recursos (PG pool, Redis connections, CPU) sob carga excessiva ou ataque DoS. | ~~Alta~~ ✅ Resolvida (Fase 0) | Env vars `RATE_LIMIT_MAX`, `RATE_LIMIT_WINDOW` |
| DT-12 | Pipeline de cálculo completamente reordenado (C-001). Motor 7-fases SOP-013: IS(F0)→IPI(F1)→CBS(F2)→ICMS(F3)→(IBS+ISS+PISCOFINS)(F4)→FUST(F5)→FUNTTEL(F6). CBS "por fora" antes do ICMS. ICMS sequencial antes do PIS/COFINS (Tese do Século — VALOR_EXCLUSAO_ICMS injetado automaticamente). FUST/FUNTTEL em cascata pós-paralela. | ~~Alta~~ ✅ Resolvida | `internal/calculator/engine.go`, `cmd/api/main.go` |
| DT-15 | `ReformaCalculator` acumulava CBS, IBS e IS em um único calculator. IS extraído como pré-filtro independente (F-006) — `ISFilter` em Fase 0. | ~~Média~~ ✅ Resolvida | `internal/reforma/reforma.go`, `internal/legacy/is_filter.go` |
| DT-16 | Phase Resolution System (F-005) implementado mas sem integração de pipeline automática: o `TaxSelector` seleciona calculadoras ativas, mas o wiring em `main.go` ainda é estático. Idealmente as calculadoras da Fase 2 deveriam ser montadas dinamicamente via `TaxSelector.Filter()`. |   Baixa | `cmd/api/main.go:74-86` |
| DT-17 | Testes de integração do PhaseResolver com engine existem via unit tests mas não ha testes de integração end-to-end com o endpoint `/calculate` passando datas de operação em diferentes fases. |   Baixa | `internal/phase/phase_test.go`, `internal/calculator/engine_test.go` |
