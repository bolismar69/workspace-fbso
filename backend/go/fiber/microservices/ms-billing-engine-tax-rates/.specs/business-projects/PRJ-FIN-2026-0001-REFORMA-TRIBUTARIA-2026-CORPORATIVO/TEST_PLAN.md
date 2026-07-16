# TEST PLAN — Plano de Testes

**Projeto:** PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO
**Microserviço:** `ms-billing-engine-tax-rates`
**Versão:** 2.0
**Data:** 30 de Junho de 2026
**Status:** ✅ **Todos os cenários de teste executados** — PR #6 merged

> 📋 **Propósito:** Este documento foi o plano de testes para os 10 GAPs. Todos os cenários foram implementados e executados. Baseline atual: **211+ testes em 25 arquivos** (vs. 150+ em 12 arquivos no plano original). As seções abaixo refletem o plano original; os testes reais estão em `internal/*/`.

📄 **Referências:** [SPECS.md](./SPECS.md) (escopo dos GAPs), [TASKS.md](./TASKS.md) (tarefas de implementação), [ARCHITECTURE.md](./ARCHITECTURE.md) (decisões técnicas)

---

## 1. Estratégia de Testes por Nível

Cada GAP é testado em até 4 níveis, conforme sua natureza:

| Nível | Escopo | Ferramenta | Quando Aplicar |
|:---|:---|:---|:---|
| **L1 — Unitário** | Funções, métodos, estratégias isoladas | `go test` + mocks | Todos os GAPs |
| **L2 — Integração** | Múltiplos componentes reais (engine + repository + cache) | `go test` com PostgreSQL e Redis reais (testcontainers ou docker-compose) | GAPs com nova lógica de negócio |
| **L3 — Contrato** | API endpoints (request/response, status codes, headers) | `go test` + `net/http/httptest` | GAPs com novos endpoints |
| **L4 — Performance** | Latência, throughput, consumo de recursos | `go test -bench` + `wrk`/`k6` | KPI O2 (<100ms p95) |

---

## 2. Linha de Base — Testes Existentes

Antes de adicionar novos testes, a linha de base deve permanecer verde:

| Suite | Arquivos | Cenários | Comando |
|:---|:---|:---|:---|
| Pipeline SOP-013 | `engine_test.go`, `pipeline_test.go` | 28 | `go test ./internal/calculator/...` |
| CBS/IBS/IS | `reforma/` (calculadoras + testes) | 7 | `go test ./internal/reforma/...` |
| IS Filter | `is_filter_test.go` | 8 | `go test ./internal/legacy/... -run ISFilter` |
| Phase Resolver | `phase_test.go` | 14 | `go test ./internal/phase/...` |
| ICMS (Próprio, ST, DIFAL, SN) | `icms_test.go` | 12 | `go test ./internal/legacy/... -run ICMS` |
| PIS/COFINS | `pis_cofins_calculate_test.go` | 39 | `go test ./internal/legacy/... -run PISCOFINS` |
| IPI | `ipi_test.go` | 7 | `go test ./internal/legacy/... -run IPI` |
| ISS + FUST + FUNTTEL | `iss_test.go`, `fust_test.go`, `funttel_test.go` | 17 | `go test ./internal/legacy/... -run 'ISS|FUST|FUNTTEL'` |
| ICMS Desonerado | `icms_desoneracao_test.go` | 11 | `go test ./internal/legacy/... -run Desoneracao` |
| Circuit Breaker + IBS Client | `circuitbreaker/`, `ibsclient/` | 12 | `go test ./internal/circuitbreaker/... ./internal/ibsclient/...` |
| Middleware (requestid, auth) | `requestid_test.go`, `auth_test.go` | 21 | `go test ./internal/middleware/...` |
| **Total Baseline** | **25 arquivos** | **211+** | `go test ./...` |

> ⚠️ **Atualização pós-implementação (2026-06-30):** Baseline executada com sucesso — 211+ testes passando em 25 arquivos. Nenhuma regressão.

---

## 3. Fase 0 — Fundação

### GAP-008: Rate Limiting

**Níveis:** L1 (Unitário) + L3 (Contrato)

| ID | Cenário de Teste | Nível | Tipo | BDD de Origem |
|:---|:---|:---|:---|:---|
| TST-008.01 | Requisições dentro do limite passam normalmente (até 100 req/min) | L3 | Positivo | — |
| TST-008.02 | Requisição acima do limite retorna `429 Too Many Requests` | L3 | Negativo | — |
| TST-008.03 | Header `Retry-After` presente na resposta 429 com segundos restantes | L3 | Contrato | — |
| TST-008.04 | Após janela de tempo (`RATE_LIMIT_WINDOW`), contador reseta e requisições voltam a passar | L1 | Temporal | — |
| TST-008.05 | IPs diferentes têm contadores independentes (cliente A bloqueado não afeta cliente B) | L1 | Isolamento | — |
| TST-008.06 | Rate limit com burst — permite rajada curta acima da média mas abaixo do burst máximo | L1 | Edge | — |

### GAP-009: API Versioning

**Níveis:** L3 (Contrato)

| ID | Cenário de Teste | Nível | Tipo | BDD de Origem |
|:---|:---|:---|:---|:---|
| TST-009.01 | `POST /v1/calculate` retorna 200 com payload válido | L3 | Positivo | — |
| TST-009.02 | `POST /calculate` (sem versão) redireciona para `/v1/calculate` com header `Deprecation: true` | L3 | Compatibilidade | — |
| TST-009.03 | `GET /v1/healthz` e `GET /v1/health` retornam 200 | L3 | Positivo | — |
| TST-009.04 | Endpoint inexistente `/v2/calculate` retorna 404 (não quebra consumidores que tentarem v2 antes do tempo) | L3 | Negativo | — |

### GAP-010: Deploy Artifacts

**Níveis:** N/A (testes manuais de infraestrutura)

| ID | Cenário de Teste | Nível | Tipo |
|:---|:---|:---|:---|
| TST-010.01 | `docker build -t ms-tax-rates .` completa sem erro | — | Build |
| TST-010.02 | Container sobe e responde em `localhost:3000/v1/healthz` em <5s | — | Saúde |
| TST-010.03 | `docker-compose up` sobe app + PostgreSQL + Redis — `/v1/health` reporta todos os checks verdes | — | Integração |
| TST-010.04 | K8s liveness probe (`/v1/healthz`) mantém pod Running por 5min sem restart | — | Estabilidade |
| TST-010.05 | K8s readiness probe (`/v1/health`) falha quando PG está offline → pod removido do Service | — | Resiliência |

---

## 4. Fase 1 — Onda 1 Comercial

### GAP-004: Campo `valor_liquido` no Response

**Níveis:** L1 (Unitário) + L3 (Contrato)

| ID | Cenário de Teste | Nível | Tipo | BDD de Origem |
|:---|:---|:---|:---|:---|
| TST-004.01 | Item com tributos > 0 — `valor_liquido = valor_item − total_impostos` (ex: 1000 − 250 = 750) | L1 | Positivo | BR-04 — Transparência |
| TST-004.02 | Item isento (todos os tributos = 0) — `valor_liquido = valor_item` (1000 − 0 = 1000) | L1 | Positivo | BR-04 |
| TST-004.03 | Item com impostos > valor (cenário anômalo) — `valor_liquido = 0` (nunca negativo) | L1 | Edge | — |
| TST-004.04 | Response JSON contém campo `valor_liquido` no schema `TributosItem` | L3 | Contrato | BR-04 |

### GAP-002: TaxToken — Congelamento de Alíquota

**Níveis:** L1 (Unitário) + L2 (Integração) + L3 (Contrato)

| ID | Cenário de Teste | Nível | Tipo | BDD de Origem |
|:---|:---|:---|:---|:---|
| TST-002.01 | `POST /v1/token/generate` com NCM+UF+IBGE válidos retorna token com alíquotas corretas e TTL futuro | L3 | Positivo | BR-06 — Garantia de Preço |
| TST-002.02 | `POST /v1/calculate` com `token_id` válido usa alíquotas do token (não consulta banco) e retorna mesmo valor do token | L2 | Integração | BR-06 |
| TST-002.03 | `POST /v1/calculate` com `token_id` expirado retorna `409 Conflict` com mensagem "token expirado" | L3 | Negativo | BR-06 |
| TST-002.04 | Token com TTL = 0 (não configurado) — endpoint rejeita com `400 Bad Request` | L1 | Validação | — |
| TST-002.05 | `GET /v1/token/{id}` retorna token válido com `status: "valido"` e `expires_at` | L3 | Positivo | BR-06 |
| TST-002.06 | `GET /v1/token/{id}` para token expirado retorna `status: "expirado"` | L3 | Temporal | BR-06 |
| TST-002.07 | Dois cálculos com mesmo token produzem resultados idênticos (idempotência) | L1 | Idempotência | BR-06, BR-07 |
| TST-002.08 | Concorrência — 10 requisições simultâneas com mesmo token não corrompem estado | L1 | Concorrência | — |

### GAP-003: Endpoint de Simulação `/simulate`

**Níveis:** L1 (Unitário) + L2 (Integração) + L3 (Contrato)

| ID | Cenário de Teste | Nível | Tipo | BDD de Origem |
|:---|:---|:---|:---|:---|
| TST-003.01 | Simulação com 1 destino — retorna margem líquida calculada: `(valor_liquido − custo) / valor_liquido` | L2 | Positivo | BR-05 — Proteção de Margem |
| TST-003.02 | Simulação com 3 destinos (SP, RJ, MG) — retorna array com impacto por UF e margem de cada | L2 | Positivo | BR-05 |
| TST-003.03 | Simulação com margem negativa (impostos > margem bruta) — retorna `margem_liquida` negativo e flag `alerta_margem: true` | L1 | Edge | BR-05 |
| TST-003.04 | Simulação com custo zero — `margem_liquida = valor_liquido / valor_liquido = 1` (100%) | L1 | Edge | — |
| TST-003.05 | Destino sem alíquota IBS cadastrada — retorna `impacto_por_uf` com `ibs: 0` e `warning: "aliquota_ibs_nao_encontrada"` | L2 | Edge | BR-05 |
| TST-003.06 | Simulação com produto sujeito a IS — IS incluso no impacto e na margem | L1 | Positivo | IS Compliance |

### GAP-001: Interface Admin para Time Fiscal

**Níveis:** L1 (Unitário) + L2 (Integração) + L3 (Contrato)

| ID | Cenário de Teste | Nível | Tipo | BDD de Origem |
|:---|:---|:---|:---|:---|
| TST-001.01 | `PUT /v1/admin/tax-rates/iva-dual` com payload válido insere nova regra e retorna 201 | L3 | Positivo | BR-02 — Autonomia Fiscal |
| TST-001.02 | `PUT` com mesmo NCM+UF+IBGE atualiza regra existente — `final_validade` da antiga é fechado, nova é inserida | L2 | Integração | BR-02 |
| TST-001.03 | `PUT` com alíquota CBS > 100% retorna `400 Bad Request` com campo inválido | L3 | Validação | BR-02 |
| TST-001.04 | `PUT` com NCM inválido (não numérico) retorna `400 Bad Request` | L3 | Validação | — |
| TST-001.05 | `GET /v1/admin/tax-rates/iva-dual?ncm=84713012` retorna lista de regras ativas para aquele NCM | L3 | Positivo | BR-02 |
| TST-001.06 | `GET` com filtro `?uf=SP` retorna apenas regras para São Paulo | L3 | Filtro | BR-02 |
| TST-001.07 | Trigger de auditoria — após `PUT`, `iva_dual_rules_log` contém registro com `operation_type='U'`, `changed_by` do JWT, snapshot completo | L2 | Auditoria | BR-02 |
| TST-001.08 | Cache Redis invalidado após `PUT` — próxima consulta `GetIvaDualRule` vai ao banco (cache miss) e reaquece | L2 | Cache | BR-02 |

---

## 5. Fase 2 — Onda 2 Financeira

### GAP-006: Schema Split Payment no Response

**Níveis:** L1 (Unitário) + L3 (Contrato)

| ID | Cenário de Teste | Nível | Tipo | BDD de Origem |
|:---|:---|:---|:---|:---|
| TST-006.01 | Cálculo normal — `valor_receita_liquida = total − (cbs + ibs + is)`, cada imposto com valor exato a reter | L1 | Positivo | BR-09 — Split Payment |
| TST-006.02 | Operação isenta (todos os tributos = 0) — `valor_receita_liquida = total`, todos os `valor_*_reter = 0` | L1 | Positivo | BR-09 |
| TST-006.03 | Operação interestadual — IBS estadual e municipal corretamente segregados no split | L1 | Positivo | BR-09 |
| TST-006.04 | `CodigoBarrasSplit` (SHA-256) é determinístico — mesmo input produz mesmo hash | L1 | Idempotência | BR-09 |
| TST-006.05 | Response JSON contém bloco `split_payment` com todos os campos obrigatórios | L3 | Contrato | BR-09 |

### GAP-005: Cálculo de Créditos na Entrada

**Níveis:** L1 (Unitário) + L2 (Integração) + L3 (Contrato)

| ID | Cenário de Teste | Nível | Tipo | BDD de Origem |
|:---|:---|:---|:---|:---|
| TST-005.01 | NF-e de fornecedor qualificado (Lucro Real, certidão válida) — crédito CBS e IBS integrais | L2 | Positivo | BR-08 — Rastreabilidade de Créditos |
| TST-005.02 | NF-e de fornecedor não qualificado (certidão vencida) — `credito_total = 0`, `motivo_bloqueio = "fornecedor_nao_qualificado"` | L2 | Negativo | BR-08 |
| TST-005.03 | Fornecedor Simples Nacional — crédito parcial (apenas percentual permitido do repasse) | L1 | Regra | BR-08 |
| TST-005.04 | Produto com `permite_credito_amplo = false` — crédito bloqueado mesmo com fornecedor qualificado | L1 | Regra | BR-08 |
| TST-005.05 | Imposto Seletivo (IS) na NF-e de entrada — IS NÃO gera crédito (não é recuperável) | L1 | Regra | IS Compliance |
| TST-005.06 | NF-e com valor zero — `credito_total = 0`, sem erro | L1 | Edge | — |
| TST-005.07 | `GET /v1/credit/summary?periodo=2026-06` retorna total de créditos disponíveis vs. apropriados no mês | L3 | Positivo | BR-08, KPI F1 |
| TST-005.08 | Período sem dados — `GET /v1/credit/summary?periodo=2025-01` retorna totais zerados (não 404) | L3 | Edge | — |
| TST-005.09 | NF-e com crédito > valor do item (cenário anômalo) — crédito limitado ao valor do item | L1 | Edge | — |
| TST-005.10 | Múltiplos itens na mesma NF-e — crédito calculado por item e totalizado corretamente | L1 | Positivo | BR-08 |

### GAP-007: Qualificação Fiscal de Fornecedores

**Níveis:** L1 (Unitário) + L2 (Integração) + L3 (Contrato)

| ID | Cenário de Teste | Nível | Tipo | BDD de Origem |
|:---|:---|:---|:---|:---|
| TST-007.01 | `POST /v1/supplier/validate` com CNPJ válido e dados fiscais ok — retorna `status: "qualificado"` | L3 | Positivo | BR-08 |
| TST-007.02 | `POST /v1/supplier/validate` com certidão vencida — retorna `status: "bloqueado"`, `motivo: "certidao_vencida"` | L2 | Negativo | BR-08 |
| TST-007.03 | `POST /v1/supplier/validate` com regime Simples Nacional — retorna `status: "qualificado_restrito"`, `permite_credito: false` (crédito limitado ao repasse) | L2 | Regra | BR-08 |
| TST-007.04 | `PUT /v1/supplier/{cnpj}` atualiza dados de qualificação — nova data de validade, novo status | L3 | Positivo | BR-08 |
| TST-007.05 | `GET /v1/supplier/{cnpj}` para fornecedor qualificado — retorna dados completos | L3 | Positivo | BR-08 |
| TST-007.06 | `GET /v1/supplier/{cnpj}` para CNPJ não cadastrado — retorna `404` | L3 | Negativo | — |
| TST-007.07 | Integração GAP-005 ↔ GAP-007 — `POST /v1/credit/calculate` para fornecedor bloqueado automaticamente zera crédito | L2 | Integração | BR-08 |
| TST-007.08 | CNPJ inválido (formato errado) — `POST /v1/supplier/validate` retorna `400 Bad Request` | L3 | Validação | — |

---

## 6. Testes de Performance — KPI O2 (<100ms p95)

### 6.1 Benchmark de Cálculo (L4)

```go
// Arquivo: internal/calculator/benchmark_test.go
func BenchmarkCalculateFullPipeline(b *testing.B) {
    input := &models.DocumentoFiscalEntrada{
        Itens: []models.ItemDocumentoFiscalEntrada{
            {NCM: "84713012", Valor: decimal.NewFromFloat(1000), UFOrigem: "SP", UFDestino: "RJ", MunicipioIBGE: "3304557"},
        },
    }
    engine := BillingEnginePhased(SOP013Phases()...)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        engine.Calculate(context.Background(), input)
    }
}
```

| Métrica | Alvo | Comando |
|:---|:---|:---|
| Tempo médio por cálculo (1 item) | < 50ms | `go test -bench=CalculateFullPipeline -benchtime=100x ./internal/calculator/` |
| Tempo médio por cálculo (100 itens) | < 100ms | Benchmark com lote de 100 itens |
| Alocações por cálculo | < 1MB | `go test -bench=CalculateFullPipeline -benchmem ./internal/calculator/` |

### 6.2 Load Test (L4)

Após deploy em staging, executar com `wrk` ou `k6`:

```bash
# wrk — 100 conexões por 30 segundos
wrk -t4 -c100 -d30s -s post_calculate.lua http://localhost:3000/v1/calculate

# k6 — script de carga progressiva (ramp-up)
k6 run deploy/loadtest/calculate-load.js
```

| Métrica | Alvo |
|:---|:---|
| p95 latência | < 100ms |
| p99 latência | < 200ms |
| Throughput | ≥ 1000 req/s (single instance) |
| Taxa de erro | < 0.01% |
| Cache hit rate (Redis) | ≥ 90% |

### 6.3 Cenários de Carga

| Cenário | Descrição | Esperado |
|:---|:---|:---|
| **Carga normal** | 500 req/s constantes por 5min | p95 < 100ms, zero erros |
| **Pico de vendas** | 2000 req/s por 2min (simula Black Friday) | p95 < 200ms, throttling via rate limit, sem OOM |
| **Recuperação** | Após pico, volta a 500 req/s | p95 retorna a < 100ms em < 60s |
| **Cache frio** | Redis vazio, 100 req/s | p95 < 150ms (banco direto), cache aquece em < 30s |

---

## 7. Suíte de Regressão (CI/CD)

Executada a cada commit e obrigatória antes de merge:

```bash
#!/bin/bash
# Arquivo: scripts/regression.sh
set -e

echo "=== L1: Testes Unitários ==="
go test -v -count=1 -timeout 120s ./...

echo "=== L1: Cobertura de Código ==="
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out | grep total

echo "=== L3: Testes de Contrato (OpenAPI) ==="
# Validar que tax-rates-api.yaml é válido
# (ferramenta: spectral lint ou openapi-generator validate)

echo "=== Sanidade: go vet ==="
go vet ./...

echo "=== Regressão completa concluída ==="
```

| Gatilho | Suite | Timeout Máximo |
|:---|:---|:---|
| Push para branch de feature | L1 (unitários) | 2 min |
| PR aberto/atualizado | L1 + L3 (contrato) | 5 min |
| Merge para main | L1 + L2 + L3 (completa) | 10 min |
| Deploy para staging | L4 (performance smoke) | 15 min |

---

## 8. Matriz BDD → Teste Técnico

Mapeamento dos critérios de aceite BDD das User Stories para os testes técnicos definidos neste plano:

### Onda 1 — Canais Comerciais

| User Story | Critério BDD (resumo) | Teste Técnico |
|:---|:---|:---|
| US 01-03-1 — Decomposição IVA | Dado checkout com item de SP para RJ, Quando exibo preço, Então vejo preço base + CBS + IBS segregados | TST-004.01, TST-004.04 |
| US 01-03-3 — Token Fiscal | Dado proposta comercial fechada, Quando gero token, Então cálculo em até 60min usa mesmas alíquotas | TST-002.01 a TST-002.08 |
| US 01-02-1 — Simulador Omnicanal | Dado vendedor no CRM, Quando simula venda para 3 estados, Então vê margem líquida por UF | TST-003.01 a TST-003.06 |
| US 01-01-2 — Trava Comercial CRM | Dado cliente sem IBGE, Quando CRM tenta simular, Então endpoint retorna validação (responsabilidade do CRM — OUT-01) | Fora do escopo do microserviço |

### Onda 2 — Finanças, Faturamento e ERP

| User Story | Critério BDD (resumo) | Teste Técnico |
|:---|:---|:---|
| US 02-02-1 — Conciliação Split | Dado pagamento de fatura, Quando banco processa, Então valor líquido + CBS + IBS batem com NF-e | TST-006.01 a TST-006.05 |
| US 02-03-1 — Bloqueio Créditos | Dado NF-e de entrada, Quando fornecedor não é qualificado, Então crédito = 0 com motivo | TST-005.01 a TST-005.10, TST-007.07 |
| US 02-02-2 — Ajuste Split Incentivado | Dado operação com benefício fiscal, Quando split é calculado, Então retenção usa alíquota reduzida | TST-006.02, TST-006.03 |
| US 02-03-3 — Reserva Incentivos | Dado ganho por regime especial, Quando contabilizo, Então valor vai para reserva de subvenção (SAP, não microserviço) | Fora do escopo (OUT-05) |

---

## 9. Resumo de Cobertura de Testes

| Fase | GAPs | Cenários de Teste | L1 | L2 | L3 | L4 |
|:---|:---|:---|:---|:---|:---|:---|
| **Fase 0** | GAP-008, 009, 010 | 15 cenários | ✅ 2 | — | ✅ 8 | — |
| **Fase 1** | GAP-004, 002, 003, 001 | 26 cenários | ✅ 9 | ✅ 9 | ✅ 11 | — |
| **Fase 2** | GAP-006, 005, 007 | 23 cenários | ✅ 8 | ✅ 9 | ✅ 8 | — |
| **Performance** | KPI O2 | 4 cenários de carga | — | — | — | ✅ 4 |
| **Regressão** | Todos | 150+ baseline + 64 novos | CI | PR | Main | Staging |
| **Total** | **10 GAPs + KPI** | **~68 novos cenários** | — | — | — | — |

---

> 📋 **Execução:** Cada tarefa no [TASKS.md](./TASKS.md) referencia seu cenário de teste correspondente (coluna `ID` nas tabelas acima). Ao implementar uma tarefa, execute o cenário de teste associado. Ao concluir um GAP, verifique que todos os cenários daquele GAP passam.
