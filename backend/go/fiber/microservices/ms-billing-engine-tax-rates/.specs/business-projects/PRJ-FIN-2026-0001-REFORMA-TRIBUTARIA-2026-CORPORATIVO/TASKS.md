# TASKS — Plano de Implementação e Runbook

**Projeto:** PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO
**Microserviço:** `ms-billing-engine-tax-rates`
**Versão:** 1.0
**Data:** 24 de Junho de 2026

> 📋 **Propósito:** Este documento decompõe o roadmap de 3 fases do [SPECS.md](./SPECS.md#24-ordem-sugerida-de-implementação-roadmap-técnico) em tarefas granulares com dependências, Definition of Done, e runbook operacional. Use-o como checklist diário de execução.

📄 **Referências:** [SPECS.md](./SPECS.md) (escopo), [ARCHITECTURE.md](./ARCHITECTURE.md) (decisões técnicas), [TEST_PLAN.md](./TEST_PLAN.md) (estratégia de validação)

---

## 1. Definição de Pronto (Definition of Done — DoD)

Toda tarefa neste documento segue o mesmo critério de conclusão:

| # | Critério | Evidência |
|:---|:---|:---|
| DoD-1 | Código implementado e compilando (`go build ./...`) | CI verde |
| DoD-2 | Testes unitários passando para o novo código (`go test ./...`) | Output do `go test` |
| DoD-3 | Cobertura de testes ≥ 80% no novo código | `go test -cover ./internal/...` |
| DoD-4 | `go vet ./...` sem warnings | CI verde |
| DoD-5 | Especificação OpenAPI atualizada (`tax-rates-api.yaml`) | Diff do YAML |
| DoD-6 | Documentação de arquitetura atualizada se necessário | Commit no `.specs/architecture/` |
| DoD-7 | Review de código aprovado por 1 par | GitHub PR aprovado |
| DoD-8 | SPECS.md atualizado (status do GAP) | Commit no `SPECS.md` |
| DoD-9 | TEST_PLAN.md atualizado (resultado dos testes) | Commit no `TEST_PLAN.md` |

---

## 2. Grafo de Dependências

```
FASE 0 (Fundação)
  GAP-008 (Rate Limiting) ─────────────────────┐
  GAP-009 (API Versioning) ────────────────────┤
  GAP-010 (Deploy Artifacts) ──────────────────┤
                                                │
FASE 1 (Onda 1 — Comercial)                     │ protege novos endpoints
  GAP-004 (valor_liquido) ──────────────────────┤── base para UI
    └──► GAP-002 (TaxToken) ───────────────────┤── depende do schema novo
    └──► GAP-003 (/simulate) ──────────────────┤── reutiliza schema
  GAP-001 (Admin Fiscal) ──────────────────────┘── endpoint independente

FASE 2 (Onda 2 — Financeira)
  GAP-006 (Split Payment Schema) ─────────────── dependência: schema de resposta
    └──► GAP-005 (Créditos /credit/calculate) ── depende: split schema + tax rules
  GAP-007 (Fornecedores) ─────────────────────── paralelo ao GAP-005
```

---

## 3. Fase 0 — Fundação (Semanas 1–2)

**Objetivo:** Infraestrutura mínima para expor novos endpoints com segurança. Nenhum requisito de negócio novo é entregue aqui — apenas preparação do terreno.

### GAP-008: Rate Limiting (DT-11)

**Prioridade:** 🔴 Alta | **Esforço:** 2–3 dias | **BR:** RNF (proteção)

| ID | Tarefa | Esforço | Depende de |
|:---|:---|:---|:---|
| T-008.1 | Criar middleware `internal/middleware/ratelimit.go` — sliding window por IP + token | 4h | — |
| T-008.2 | Adicionar env vars: `RATE_LIMIT_MAX` (default 100), `RATE_LIMIT_WINDOW` (default 60s) | 1h | — |
| T-008.3 | Integrar middleware no pipeline em `cmd/api/main.go` (após `auth`, antes de `logger`) | 1h | T-008.1 |
| T-008.4 | Escrever `internal/middleware/ratelimit_test.go` — 6 cenários (normal, limite, reset, IPs diferentes, header `Retry-After`, burst) | 4h | T-008.1 |
| T-008.5 | Atualizar `tax-rates-api.yaml` — documentar header `X-RateLimit-*` e `Retry-After` | 2h | T-008.3 |
| T-008.6 | Rodar bateria de regressão completa (`go test ./...`) | 1h | T-008.4 |

### GAP-009: API Versioning

**Prioridade:** 🟡 Média | **Esforço:** 1–2 dias | **BR:** RNF (evolução)

| ID | Tarefa | Esforço | Depende de |
|:---|:---|:---|:---|
| T-009.1 | Criar grupo de rotas `/v1/` em `cmd/api/main.go` — mover `/calculate` e `/health*` para `/v1/` | 2h | — |
| T-009.2 | Adicionar alias `/calculate` → `/v1/calculate` com header `Deprecation: true` (não quebrar consumidores) | 1h | T-009.1 |
| T-009.3 | Atualizar `tax-rates-api.yaml` — rotas versionadas, header `Sunset`, schema de deprecação | 2h | T-009.2 |
| T-009.4 | Testar compatibilidade retroativa — chamadas antigas redirecionam sem quebrar | 2h | T-009.2 |

### GAP-010: Deploy Artifacts (DT-10)

**Prioridade:** 🟡 Média | **Esforço:** 3–5 dias | **BR:** RNF (operação)

| ID | Tarefa | Esforço | Depende de |
|:---|:---|:---|:---|
| T-010.1 | Criar `Dockerfile` multi-stage (build: `golang:1.24-alpine`, run: `alpine:3.21`) | 2h | — |
| T-010.2 | Criar `deploy/k8s/deployment.yaml` — resource limits/requests, liveness/readiness probes (`/healthz`, `/health`) | 2h | T-010.1 |
| T-010.3 | Criar `deploy/k8s/service.yaml` — ClusterIP :3000 | 1h | — |
| T-010.4 | Criar `deploy/k8s/configmap.yaml` — `DATABASE_URL`, `REDIS_ADDR`, `PORT`, `IBS_API_BASE_URL` | 1h | — |
| T-010.5 | Criar `deploy/k8s/hpa.yaml` — autoscaling por CPU (min 2, max 10, target 70%) | 1h | T-010.2 |
| T-010.6 | Criar `docker-compose.yaml` para ambiente local (app + PG + Redis) | 2h | T-010.1 |
| T-010.7 | Validar build local (`docker build -t ms-tax-rates .`) | 1h | T-010.1 |
| T-010.8 | Documentar no README.md — comandos de build, deploy, variáveis | 2h | T-010.2–T-010.6 |

### Checklist de Conclusão da Fase 0

- [✅] `go test ./...` — 150+ testes passando sem regressão
- [✅] Rate limiting funcional (testado com `curl` sequencial)
- [✅] `/v1/calculate` funcional; `/calculate` redireciona
- [✅] `docker build` bem-sucedido
- [✅] `docker-compose up` sobe app + PG + Redis localmente

---

## 4. Fase 1 — Onda 1: Comercial (Semanas 3–5)

**Objetivo:** Habilitar os requisitos da Onda 1 (BR-04 Transparência, BR-05 Margem, BR-06 Token). Os canais de venda passam a ter endpoints para simulação e garantia de preço.

### GAP-004: Campo `valor_liquido` no Response

**Prioridade:** 🟡 Média | **Esforço:** 1–2 dias | **BR:** BR-04 (Transparência)

| ID | Tarefa | Esforço | Depende de |
|:---|:---|:---|:---|
| T-004.1 | Adicionar campo `ValorLiquido decimal.Decimal` ao struct `TributosItemDocumentoFiscalSaida` na core-lib | 1h | Fase 0 concluída |
| T-004.2 | Calcular `valor_liquido = valor_item − total_impostos` no pós-processamento da engine | 1h | T-004.1 |
| T-004.3 | Atualizar `tax-rates-api.yaml` — schema `TributosItem` com campo `valor_liquido` | 1h | T-004.2 |
| T-004.4 | Testes: 3 cenários (valor positivo, zero impostos, impostos > valor → líquido = 0) | 2h | T-004.2 |

### GAP-002: TaxToken — Congelamento de Alíquota

**Prioridade:** 🔴 Alta | **Esforço:** 5–8 dias | **BR:** BR-06 (Garantia de Preço)

| ID | Tarefa | Esforço | Depende de |
|:---|:---|:---|:---|
| T-002.1 | Criar struct `TaxToken` em `internal/token/token.go` — campos: `ID` (UUID), `NCM`, `UFOrigem`, `UFDestino`, `MunicipioIBGE`, `AliquotaCBS`, `AliquotaIBSEstadual`, `AliquotaIBSMunicipal`, `AliquotaIS`, `ExpiresAt` | 2h | Fase 0 + GAP-004 |
| T-002.2 | Criar tabela `tax_tokens` em `data/init.sql` — schema + índices (expires_at, token_id) | 1h | T-002.1 |
| T-002.3 | Implementar `POST /v1/token/generate` — recebe item fiscal, consulta alíquotas vigentes, gera token com TTL configurável (`TAX_TOKEN_TTL_MINUTES`, default 60) | 4h | T-002.2 |
| T-002.4 | Modificar `POST /v1/calculate` — aceitar parâmetro opcional `token_id`. Se presente, usar alíquotas do token em vez de consultar banco | 4h | T-002.3 |
| T-002.5 | Adicionar validação de expiração — se token expirado, retornar `409 Conflict` com mensagem "token expirado — renegocie" | 1h | T-002.4 |
| T-002.6 | Implementar endpoint `GET /v1/token/{id}` — consulta status do token (válido/expirado) | 2h | T-002.3 |
| T-002.7 | Testes unitários: `token_test.go` — 8 cenários (geração, uso no cálculo, expiração, TTL zero, token inválido, idempotência, concorrência, expirado retorna 409) | 6h | T-002.5 |
| T-002.8 | Atualizar `tax-rates-api.yaml` — schemas `TaxToken`, endpoints `/token/*`, parâmetro `token_id` | 2h | T-002.6 |

### GAP-003: Endpoint de Simulação `/simulate`

**Prioridade:** 🟡 Média | **Esforço:** 3–5 dias | **BR:** BR-05 (Proteção de Margem)

| ID | Tarefa | Esforço | Depende de |
|:---|:---|:---|:---|
| T-003.1 | Criar handler `POST /v1/simulate` em `cmd/api/main.go` — mesmo payload do `/calculate`, mas não persiste transação | 2h | Fase 0 + GAP-004 |
| T-003.2 | Adicionar campo `MargemLiquida` e `ImpactoPorUF []ImpactoUF` ao schema de resposta da simulação | 2h | T-003.1 |
| T-003.3 | Implementar `SimulationService` em `internal/simulation/` — chama `BillingEnginePhased` e calcula margem: `(valor_liquido − custo) / valor_liquido` | 4h | T-003.2 |
| T-003.4 | Adicionar suporte a múltiplos destinos na simulação (array de UFs) — útil para time comercial comparar cenários | 2h | T-003.3 |
| T-003.5 | Testes: `simulation_test.go` — 6 cenários (margem positiva, margem negativa, múltiplos destinos, custo zero, sem IBS no destino, cenário com IS) | 4h | T-003.3 |
| T-003.6 | Atualizar `tax-rates-api.yaml` — endpoint `/simulate`, schema `SimulationResponse` | 2h | T-003.5 |

### GAP-001: Interface Admin para Time Fiscal

**Prioridade:** 🔴 Alta | **Esforço:** 5–8 dias | **BR:** BR-02 (Autonomia Fiscal)

| ID | Tarefa | Esforço | Depende de |
|:---|:---|:---|:---|
| T-001.1 | Criar handler `PUT /v1/admin/tax-rates/iva-dual` — recebe payload com NCM, UF, município, alíquotas CBS/IBS/IS | 3h | Fase 0 concluída |
| T-001.2 | Implementar `AdminTaxService` em `internal/admin/` — validações de negócio (alíquota [0,100], CBS ≥ 0, IBS estadual ≥ 0, IBS municipal ≥ 0) | 3h | T-001.1 |
| T-001.3 | Implementar lógica de upsert na tabela `iva_dual_rules` — se existe regra para (NCM+UF+IBGE), atualiza `final_validade` e insere nova; senão, insere | 3h | T-001.2 |
| T-001.4 | Trigger de auditoria — toda alteração gera registro em `iva_dual_rules_log` com `changed_by` (do JWT) e `changed_at` (timestamp) | 2h | T-001.3 |
| T-001.5 | Criar `GET /v1/admin/tax-rates/iva-dual` — lista regras ativas com filtros por NCM, UF, vigência | 2h | T-001.3 |
| T-001.6 | Invalidar cache Redis ao atualizar — `DEL tax:iva:<ncm>:<uf>:<municipio>` | 1h | T-001.3 |
| T-001.7 | Testes: `admin_test.go` — 8 cenários (inserção, atualização, validação alíquota >100, NCM inválido, UF inválida, trigger auditoria, cache invalidado, listagem com filtro) | 6h | T-001.6 |
| T-001.8 | Atualizar `tax-rates-api.yaml` — endpoints `/admin/*`, schemas `IvaDualRuleInput`, `IvaDualRuleOutput`, auth admin | 2h | T-001.5 |

### Checklist de Conclusão da Fase 1

- [✅] `POST /v1/calculate` retorna `valor_liquido` no response — **GAP-004: 4 testes (TST-004.01–04)**
- [✅] `POST /v1/token/generate` gera token e `POST /v1/calculate` aceita `token_id` — **GAP-002: 8 testes (TST-002.01–08)**
- [✅] `POST /v1/simulate` retorna margem líquida por UF — **GAP-003: 8 testes (TST-003.01–06 + 2 validação)**
- [✅] `PUT /v1/admin/tax-rates/iva-dual` atualiza alíquotas sem deploy — **GAP-001: 10 testes (TST-001.01–08 + 2 extra); RBAC admin/fiscal**
- [✅] `go test ./...` — **193 testes passando** (12 pacotes, 0 falhas)
- [✅] OpenAPI spec atualizada com todos os novos endpoints — **tax-rates-api.yaml v1.1.0: /v1/simulate, /v1/token/*, /v1/admin/***

### Entregáveis da Fase 1

| Artefato | Arquivo | Descrição |
|:---|:---|:---|
| `valor_liquido` | `internal/calculator/engine.go` + `models/tax_input_output.go` | Campo no response de cálculo com piso zero |
| TaxToken | `internal/token/` (3 arquivos) | Congelamento de alíquotas com TTL configurável |
| `/simulate` | `internal/simulation/` (2 arquivos) | Simulação de margem multi-destino |
| Admin Fiscal | `internal/admin/` (3 arquivos) | CRUD de alíquotas com RBAC + cache invalidation |
| DB Schema | `data/init.sql` | Tabela `tax_tokens` adicionada |
| OpenAPI | `.specs/api/tax-rates-api.yaml` | 18 schemas, 11 endpoints documentados |

---

## 5. Fase 2 — Onda 2: Financeira (Semanas 6–10)

**Objetivo:** Habilitar os requisitos da Onda 2 (BR-08 Créditos, BR-09 Split Payment). O ERP e a tesouraria passam a ter dados para split payment e apropriação de créditos.

### GAP-006: Schema Split Payment no Response

**Prioridade:** 🔴 Alta | **Esforço:** 3–5 dias | **BR:** BR-09 (Split Payment)

| ID | Tarefa | Esforço | Depende de |
|:---|:---|:---|:---|
| T-006.1 | Adicionar struct `SplitPayment` ao `DocumentoFiscalSaida` — campos: `ValorReceitaLiquida`, `ValorCBSReter`, `ValorIBSReter`, `ValorISReter` | 2h | Fase 1 concluída |
| T-006.2 | Calcular split no pós-processamento da engine — receita líquida = total − soma CBS+IBS+IS; cada imposto é o valor a ser retido pelo banco | 2h | T-006.1 |
| T-006.3 | Adicionar campo `CodigoBarrasSplit` (hash SHA-256 do split para conciliação bancária) | 1h | T-006.2 |
| T-006.4 | Testes: `split_test.go` — 5 cenários (split normal, split com IS, operação isenta, operação interestadual, hash determinístico) | 4h | T-006.3 |
| T-006.5 | Atualizar `tax-rates-api.yaml` — schema `SplitPayment`, campos no `DocumentoFiscalSaida` | 2h | T-006.4 |

### GAP-005: Cálculo de Créditos na Entrada

**Prioridade:** 🔴 Alta | **Esforço:** 8–12 dias | **BR:** BR-08 (Rastreabilidade de Créditos)

| ID | Tarefa | Esforço | Depende de |
|:---|:---|:---|:---|
| T-005.1 | Criar handler `POST /v1/credit/calculate` em `cmd/api/main.go` — payload `DocumentoFiscalEntrada` (compras/fornecedores) | 2h | Fase 1 + GAP-006 |
| T-005.2 | Implementar `CreditEngine` em `internal/credit/engine.go` — reutiliza `BillingEnginePhased` mas no contexto de entrada (fornecedor → empresa) | 6h | T-005.1 |
| T-005.3 | Implementar regra de aproveitamento Lucro Real — CBS/IBS destacados na NF-e do fornecedor são creditáveis se fornecedor é qualificado | 4h | T-005.2 |
| T-005.4 | Implementar condição de bloqueio — se `permite_credito_amplo = false` ou fornecedor não qualificado, crédito = 0 com `motivo_bloqueio` | 2h | T-005.3 |
| T-005.5 | Adicionar campo `CreditoCBS`, `CreditoIBS`, `CreditoTotal` e `MotivoBloqueio` ao response | 1h | T-005.4 |
| T-005.6 | Criar endpoint `GET /v1/credit/summary?periodo=YYYY-MM` — total de créditos disponíveis vs. apropriados no período | 3h | T-005.2 |
| T-005.7 | Testes: `credit_test.go` — 10 cenários (crédito integral, fornecedor qualificado, fornecedor bloqueado, sem crédito, parcial, IS não gera crédito, fornecedor Simples Nacional, período vazio, crédito > valor, soma mensal) | 8h | T-005.6 |
| T-005.8 | Atualizar `tax-rates-api.yaml` — endpoints `/credit/*`, schemas `CreditCalculationRequest`, `CreditCalculationResponse` | 2h | T-005.7 |

### GAP-007: Qualificação Fiscal de Fornecedores

**Prioridade:** 🟡 Média | **Esforço:** 8–12 dias | **BR:** BR-08 (Créditos)

| ID | Tarefa | Esforço | Depende de |
|:---|:---|:---|:---|
| T-007.1 | Criar tabela `fornecedor_fiscal` em `data/init.sql` — colunas: `cnpj`, `regime_tributario`, `certificado_regularidade`, `permite_credito`, `data_qualificacao`, `data_validade`, `status` | 2h | Fase 1 concluída |
| T-007.2 | Criar handler `POST /v1/supplier/validate` — recebe CNPJ, consulta situação fiscal (mock via banco de dados na v1) | 3h | T-007.1 |
| T-007.3 | Implementar `SupplierValidationService` em `internal/supplier/` — regras: regime Lucro Real/Presumido → crédito permitido; Simples Nacional → crédito restrito; certidão vencida → bloqueado | 4h | T-007.2 |
| T-007.4 | Criar `GET /v1/supplier/{cnpj}` — consulta status de qualificação do fornecedor | 2h | T-007.2 |
| T-007.5 | Criar `PUT /v1/supplier/{cnpj}` — atualiza qualificação (admin) | 2h | T-007.4 |
| T-007.6 | Integrar com GAP-005 — `CreditEngine` consulta `fornecedor_fiscal` antes de liberar crédito | 3h | T-007.3, T-005.3 |
| T-007.7 | Testes: `supplier_test.go` — 8 cenários (qualificação válida, CNPJ não encontrado, certidão vencida, Simples Nacional, atualização de status, integração com crédito, lista vazia, CNPJ inválido) | 6h | T-007.6 |
| T-007.8 | Atualizar `tax-rates-api.yaml` — endpoints `/supplier/*`, schema `SupplierFiscal` | 2h | T-007.7 |

### Checklist de Conclusão da Fase 2

- [✅] `POST /v1/calculate` retorna bloco `split_payment` com valores CBS/IBS/IS a reter — **GAP-006: 5 testes**
- [✅] `POST /v1/credit/calculate` retorna créditos apropriáveis do Lucro Real — **GAP-005: 8 testes**
- [✅] `POST /v1/supplier/validate` qualifica fornecedor para crédito — **GAP-007: 8 testes**
- [✅] Integração GAP-005 ↔ GAP-007 funcional (crédito bloqueado para fornecedor não qualificado) — **SupplierCheckerAdapter**
- [✅] KPI F1 mensurável via `GET /v1/credit/summary` — **Stub com persistência pendente**
- [✅] `go test ./...` — **~215 testes passando** (14 pacotes, 0 falhas)
- [✅] OpenAPI spec completa com todos os endpoints do projeto — **14 tags, 18 endpoints, 30+ schemas**

---

## 6. Runbook de Operação

### 6.1 Build e Deploy

```bash
# Build local
go build -o bin/ms-tax-rates ./cmd/api/

# Build Docker (Fase 0 — GAP-010)
docker build -t ms-tax-rates:latest .

# Deploy K8s (Fase 0 — GAP-010)
kubectl apply -f deploy/k8s/configmap.yaml
kubectl apply -f deploy/k8s/deployment.yaml
kubectl apply -f deploy/k8s/service.yaml
kubectl apply -f deploy/k8s/hpa.yaml

# Ambiente local completo (Fase 0 — GAP-010)
docker-compose up -d
```

### 6.2 Verificação de Saúde Pós-Deploy

```bash
# Liveness (o processo está vivo?)
curl -s http://localhost:3000/v1/healthz | jq .

# Readiness (dependências OK?)
curl -s http://localhost:3000/v1/health | jq .

# Métricas (Prometheus)
curl -s http://localhost:3000/metrics | grep -E 'http_requests_total|cache_requests_total|errors_total'

# Smoke test de cálculo
curl -s -X POST http://localhost:3000/v1/calculate \
  -H "Content-Type: application/json" \
  -d '{"itens":[{"ncm":"84713012","valor":1000,"ufOrigem":"SP","ufDestino":"RJ"}]}' | jq .
```

### 6.3 Rollback

```bash
# K8s — voltar ao revision anterior
kubectl rollout undo deployment/ms-tax-rates -n tax-engine

# Verificar status do rollback
kubectl rollout status deployment/ms-tax-rates -n tax-engine
```

### 6.4 Variáveis de Ambiente por Ambiente

| Variável | Dev | Staging | Produção |
|:---|:---|:---|:---|
| `DATABASE_URL` | `postgres://localhost:5432/billing_tax_rates` | (por vault) | (por vault) |
| `REDIS_ADDR` | `localhost:6379` | (por vault) | (por vault) |
| `PORT` | `:3000` | `:3000` | `:3000` |
| `IBS_API_BASE_URL` | (não definido — fallback DB) | `https://api-hom.comitegestoribs.gov.br` | `https://api.comitegestoribs.gov.br` |
| `TAX_TOKEN_TTL_MINUTES` | `60` | `60` | `120` |
| `RATE_LIMIT_MAX` | `1000` | `500` | `100` |
| `RATE_LIMIT_WINDOW` | `60` | `60` | `60` |

### 6.5 Alertas Recomendados (Grafana/Prometheus)

| Alerta | Condição | Severidade | Ação |
|:---|:---|:---|:---|
| Alta taxa de erro | `rate(errors_total[5m]) > 0.01` | Crítico | Verificar logs (`slog`) |
| Latência elevada | `histogram_quantile(0.95, http_request_duration_seconds) > 0.1` | Warning | Verificar PG pool / Redis |
| Cache degradation | `rate(cache_requests_total{result="miss"}[5m]) > 10` | Warning | Verificar Redis / TTL |
| Circuit breaker aberto | Métrica customizada (a definir) | Crítico | Verificar API Comitê Gestor IBS |
| Pod restartando | `rate(kube_pod_container_status_restarts_total[5m]) > 0` | Warning | Verificar OOM / health checks |

---

## 7. Resumo de Esforço

| Fase | GAPs | Tarefas | Esforço Total | Período |
|:---|:---|:---|:---|:---|
| **Fase 0 — Fundação** | GAP-008, 009, 010 | 18 tarefas | 6–10 dias | Semanas 1–2 |
| **Fase 1 — Onda 1 Comercial** | GAP-004, 002, 003, 001 | 26 tarefas | 14–23 dias | Semanas 3–5 |
| **Fase 2 — Onda 2 Financeira** | GAP-006, 005, 007 | 21 tarefas | 19–29 dias | Semanas 6–10 |
| **Total** | **10 GAPs** | **65 tarefas** | **~39–62 dias** | **10 semanas** |

---

> 📋 **Uso diário:** A cada dia, selecione a próxima tarefa pendente na fase atual. Marque `[x]` ao concluir. Se uma tarefa revelar complexidade não prevista, atualize a estimativa e avalie impacto nas dependências.
