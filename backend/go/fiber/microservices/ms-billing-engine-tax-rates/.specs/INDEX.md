# Mapa de Especificações — ms-billing-engine-tax-rates

Diretório centralizado de arquitetura de contexto e especificações do microsserviço de cálculo de tributos sobre faturamento (IPI, ICMS, PIS/COFINS).

**Confiança da documentação:** 99% (VERDE) — veja [governance/confidence-report.md](governance/confidence-report.md)
**Última atualização:** 2026-06-22 (C-001 Pipeline SOP-013 7-fases concluído — todas as 9 features do gap analysis implementadas: F-001 a F-007 + C-001 + C-002)

## Engenharia e Arquitetura
- **Visão Arquitetural e Regras de Solução:** Consulte [architecture/architecture.md](architecture/architecture.md)
- **Modelo de Dados (ERD Completo):** Consulte [architecture/erd.md](architecture/erd.md)
- **Diagrama de Fluxo e C4 Context:** Consulte [architecture/c4-context.md](architecture/c4-context.md)
- **Integrações, Dependências e Libs Locais:** Consulte [architecture/integrations.md](architecture/integrations.md)
- **Análise Técnica de Fluxo de Código (Handlers/Services):** Consulte [engineering/code-analysis.md](engineering/code-analysis.md)
- **Padrões de API, Erros e Observabilidade:** Consulte [engineering/api-guidelines.md](engineering/api-guidelines.md)

## Requisitos e Escopo de Produto
- **Especificação de Requisitos (RF-01 a RF-08 / RNF / MoSCoW):** Consulte [product/requirements.md](product/requirements.md)
- **Roadmap de Features e Dívidas Técnicas (lista canônica):** Consulte [product/feature-roadmap.md](product/feature-roadmap.md)
- **Contrato OpenAPI da API (Rotas, Schemas, Erros):** Consulte [api/tax-rates-api.yaml](api/tax-rates-api.yaml)

## Regras de Negócio e Domínio Fiscal
- **Glossário Geral de Domínio (IPI, ICMS, PIS, COFINS, CSTs):** Consulte [domain/domain.md](domain/domain.md)

## Governança e Qualidade das Specs
- **Inventário do Projeto e Cobertura de Testes:** Consulte [governance/inventory.md](governance/inventory.md)
- **Dívidas Técnicas (lista canônica):** Consulte [product/feature-roadmap.md](product/feature-roadmap.md#d%C3%ADvidas-t%C3%A9cnicas) (DT-01 a DT-11 com prioridades e localização)
- **Relatório de Confiança e Score da Documentação (98%):** Consulte [governance/confidence-report.md](governance/confidence-report.md)

## Lacunas e Perguntas
- **Histórico de Lacunas Resolvidas:** Consulte [questions/questions_01.md](questions/questions_01.md)

## Esquema de Banco e Dados
- **Schema DDL + Triggers (7 tabelas, PL/pgSQL):** `data/init.sql`

## Middleware e Infraestrutura
- **W3C Trace Context (Request-ID + Trace-ID):** `internal/middleware/requestid.go`
- **JWT Auth (Kong/Keycloak pass-through):** `internal/middleware/auth.go`
- **Métricas Prometheus (/metrics):** `internal/middleware/metrics.go`
- **ISS — Imposto sobre Serviços:** `internal/legacy/iss.go` — 7 testes
- **FUST/FUNTTEL — Contribuições de Telecom:** `internal/legacy/fust.go`, `funttel.go`, `telecom.go` — 10 testes
- **Motor 7-fases SOP-013 (BillingEnginePhased):** `internal/calculator/engine.go` — pipeline IS→IPI→CBS→ICMS→(IBS+ISS+PISCOFINS)→FUST→FUNTTEL (C-001)
- **CBS/IBS split (Reforma Tributária):** `internal/reforma/cbs_calculator.go`, `ibs_calculator.go` — calculadoras separadas
- **ICMS Desonerado (F-004):** `internal/legacy/icms_desoneracao.go` — 2 modos (Redução Base + Limitação Alíquota)
- **Phase Resolution System (F-005):** `internal/phase/phase.go`, `tax_selector.go` — SHADOW_RUN/CBS_PLENA/TRANSICAO/IVA_DUAL
- **IS Pré-Filtro (F-006):** `internal/legacy/is_filter.go` — Fase 0, tabela `ncm_seletivo`
- **IBS Circuit Breaker (F-007):** `internal/circuitbreaker/`, `internal/ibsclient/` — HTTP+Redis+Fallback DB

## Histórico de Implementações (skill-output)
- **15 registros de sessão** documentando cada feature implementada em [skill-output/](skill-output/)
- **Features implementadas:** [features/FEATURE-2026-06-21.md](features/FEATURE-2026-06-21.md) — gap analysis POLICE/PROCEDURE/RULES vs. código
- **Gap Analysis (conformidade spec→código):** [features/FEATURE-2026-06-21-GAP-ANALISYS.md](features/FEATURE-2026-06-21-GAP-ANALISYS.md) — 54/55 requisitos ✅
- **Schema DDL (10 tabelas + 3 índices):** `data/init.sql` — `icms_rules`, `federal_tax_rules`, `product_tax_exceptions`, `tax_equivalence`, `simples_nacional_rates`, `ipi_regras`, `iva_dual_rules`, `ncm_seletivo`, `cbs_rates`, `iss_rates`
