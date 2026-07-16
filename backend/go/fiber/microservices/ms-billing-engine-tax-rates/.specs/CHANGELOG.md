# Changelog — ms-billing-engine-tax-rates

Histórico de mudanças significativas no microsserviço de cálculo de tributos.

---

## 2026-07-01 — DT-03: Tabela Oficial CST para CBS/IBS (LC 214/2025)

- Tabela `cst_reforma` criada com 164 CCTs oficiais da RFB (18 CSTs)
- `GetCSTReforma()` adicionado ao `TaxRepository` (PostgreSQL + cache Redis)
- `reforma.go` refatorado: constantes `cstPadrao`/`cstIsento` removidas
- CST oficial de 3 dígitos no response (substitui valores provisórios `01`/`04`)
- 3 ADRs registrados (ADR-010, ADR-011, ADR-012)
- Feature: `FEATURE-2026-06-30-DT03-CST-REFORMA-TRIBUTARIA`

---

## 2026-06-30 — PR #6: Reforma Tributária — Fases 0, 1, 2

**Fase 0 (Fundação):**
- Rate Limiting: middleware com `RATE_LIMIT_MAX` e `RATE_LIMIT_WINDOW`
- API Versioning: prefixo `/v1/` com rotas legacy deprecadas
- Deploy Docker/K8s: `Dockerfile`, `docker-compose.yaml`, `deploy/k8s/` (configmap, deployment, service, hpa)

**Fase 1 — Onda 1 (Comercial):**
- Admin Fiscal (GAP-004): CRUD de regras IVA Dual — `internal/admin/`
- TaxToken (snapshot fiscal): `internal/token/` — `/v1/token/generate`
- Simulação de margem (GAP-003): `internal/simulation/` — `/v1/simulate`

**Fase 2 — Onda 2 (Fornecedores e Créditos):**
- Validação de fornecedores: `internal/supplier/`
- Créditos Reforma Tributária (GAP-005): `internal/credit/` — `/v1/credit/calculate`

**Métricas:** 25 arquivos de teste, 211+ cenários, 10 tabelas SQL

---

## 2026-06-22 — C-001 Pipeline SOP-013 7-fases

- Pipeline reordenado de 3 para 7 fases (IS→IPI→CBS→ICMS→IBS+ISS+PISCOFINS→FUST→FUNTTEL)
- `BillingEnginePhased` com arquitetura multi-fase genérica (`CalculationPhase`, `ExecutionMode`)
- CBS/IBS split: `CBSCalculator` (Fase 2) + `IBSCalculator` (Fase 4) separados
- ICMS Desonerado (F-004), Phase Resolver (F-005), IS Pré-Filtro (F-006), IBS Circuit Breaker (F-007)
- Schema expandido (C-002): `ncm_seletivo`, `cbs_rates`, `iss_rates`
- 150+ testes

---

## 2026-06-21 — Tributos, Middleware e Reforma

- ISS (LC 116/2003), FUST (Lei 9.998/2000), FUNTTEL (Lei 10.052/2000) — F-001 a F-003
- Middleware: W3C Trace Context, JWT Auth, Métricas Prometheus
- Reforma Tributária: CBS, IBS, IS — `ReformaCalculator` integrado
- Domain restructure: interface `TaxCalculator` extraída para `internal/domain/`
- 109 testes, 12 arquivos de teste

---

## 2026-06-20 — Fundação

- ICMS (normal, ST, DIFAL, Simples Nacional) — `internal/legacy/icms.go`
- IPI (Ad Valorem, Ad Pauta, rateio) — `internal/legacy/ipi.go`
- PIS/COFINS com Strategy Pattern (15 estratégias, 100% CSTs)
- Exclusão ICMS da base PIS/COFINS ("Tese do Século")
- Motor bifásico inicial (`BillingEngineOrdered`)
- Schema SQL inicial (7 tabelas)
- 64 testes
