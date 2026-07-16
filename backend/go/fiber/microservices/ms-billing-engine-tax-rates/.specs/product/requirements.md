# Requirements: Cálculo de Tributos sobre Faturamento

> Fonte: `cmd/api/main.go:121-168`, `internal/calculator/engine.go:43-135`, `internal/legacy/`

## 1. Resumo executivo

O endpoint de cálculo de tributos processa um documento fiscal (`DocumentoFiscalEntrada`) contendo itens com NCM, CFOP, valores, e retorna os tributos calculados (IPI, ICMS, PIS, COFINS) para cada item. O motor executa os cálculos em duas fases: IPI sequencialmente, depois ICMS e PIS/COFINS em paralelo via goroutines.

## 2. Priority MoSCoW

| Item | MoSCoW | Justificativa | Evidência |
|------|--------|---------------|-----------|
| RF-01 | Must | Receber e validar payload do documento fiscal | `cmd/api/main.go:121-168` |
| RF-02 | Must | Calcular IPI (Ad Valorem e Ad Pauta) com rateio de despesas | `internal/legacy/ipi.go:1-217` |
| RF-03 | Must | Calcular ICMS com DIFAL (EC 87/2015) | `internal/legacy/icms.go:323-354` |
| RF-04 | Must | Executar motor bifásico (IPI → ICMS/PIS/COFINS) | `internal/calculator/engine.go:43-135` |
| RF-05 | Should | Calcular ICMS-ST com MVA | `internal/legacy/icms.go:129-174, 287-320` |
| RF-06 | Should | Calcular Simples Nacional via equivalência CSOSN | `internal/legacy/icms.go:72-127` |
| RF-07 | Must | Calcular PIS/COFINS por CST (01-06, 49, 50-99, 99) | `internal/legacy/pis_cofins.go:1-200`, `pis_strategies.go:1-98`, `pis_cofins_calculate_test.go` |
| RF-08 | Must | Implementar CBS/IBS/IS (Reforma Tributária) — ✅ Implementado | `internal/reforma/reforma.go` — `ReformaCalculator` com 7 testes, integrado ao motor na Fase 2 |
| RF-09 | Must | Calcular ISS sobre serviços de telecom (LC 116/2003) — ✅ Implementado | `internal/legacy/iss.go` — `ISSCalculator`, 7 testes |
| RF-10 | Must | Calcular FUST/FUNTTEL (contribuições setoriais de telecom) — ✅ Implementado | `internal/legacy/fust.go`, `funttel.go`, `telecom.go` — 10 testes |
| RF-11 | Should | Motor trifásico com Fase 3 pós-paralela — ✅ Implementado | `BillingEngineFull()` em `engine.go` |

## 3. Requisitos Funcionais

### RF-01 — Endpoint de Cálculo
**POST /calculate**

O sistema deve expor um endpoint HTTP que aceita um payload JSON `DocumentoFiscalEntrada`, valida os dados via `input.Validate()`, processa os cálculos de tributos e retorna um `DocumentoFiscalSaida`.

**Fonte:** `cmd/api/main.go:121-168`

**Erros:**
- `400` — JSON inválido ou validação falhou
- `500` — Erro interno no cálculo

### RF-02 — Cálculo de IPI

O cálculo do IPI deve seguir a seguinte prioridade:
1. Se o item tem overrides inline completos (`aliquota`, `cst`, `c_enq`), usar estes valores
2. Caso contrário, consultar `ipi_regras` no banco por NCM + CRT + tipo operação + perfil comprador + UF + zona especial
3. Fallback: usar detalhes do item como valores inline

**Rateio:** Frete, seguro, despesas e desconto são rateados proporcionalmente entre itens.

**Modalidades:**
- Ad Valorem: `valor_item × (aliquota / 100)`
- Ad Pauta: Valor fixo por unidade

**Fonte:** `internal/legacy/ipi.go:1-217`

### RF-03 — Cálculo de ICMS DIFAL

Para operações interestaduais com consumidor final não contribuinte, calcular o DIFAL:

```
DIFAL = Base × (Aliquota Interna Destino − Aliquota Interestadual)
```

**Condições:**
- UF origem ≠ UF destino
- Consumidor final (não contribuinte)
- Alíquota interna destino > alíquota interestadual
- Não Simples Nacional

**Fonte:** `internal/legacy/icms.go:323-354`

### RF-04 — Motor Bifásico

O motor deve executar os cálculos em duas fases:
1. **Fase 1 (Sequencial):** IPI calculado primeiro, resultado injetado como `ITEM_IPI_VALOR`
2. **Fase 2 (Paralela via goroutines):** ICMS e PIS/COFINS executados simultaneamente

**Fonte:** `internal/calculator/engine.go:43-135`

### RF-05 — Cálculo de ICMS-ST

Quando um produto possui exceção com CST "010" (tributado com ST), aplicar MVA:

```
Base ST = Valor do Item × (1 + MVA/100)
ICMS-ST = Base ST × Aliquota Interna Destino
```

**Fonte:** `internal/legacy/icms.go:129-174`

### RF-06 — Cálculo Simples Nacional

Para contribuintes do Simples Nacional:
1. Converter CSOSN para CST equivalente via `tax_equivalence`
2. Calcular alíquota efetiva: `((RBT12 × Aliquota Nominal) − Valor a Deduzir) / RBT12`
3. ICMS = Alíquota Efetiva × Percentual ICMS no Anexo

**Fonte:** `internal/legacy/icms.go:72-127`

### RF-07 — Cálculo PIS/COFINS por CST

Para cada item, determinar CST do regime e aplicar estratégia correspondente:
- **CST 01/02:** `(Valor × Aliquota) − Exclusão ICMS da Base`
- **CST 03:** Valor fixo por unidade (CalcTax)
- **CST 04:** Monofásico — valor zero (tributo já recolhido na produção/importação)
- **CST 05:** Substituição Tributária — valor zero (tributo já recolhido por substituto)
- **CST 06:** Alíquota Zero — valor zero
- **CST 49:** Outras Operações de Saída — valor zero
- **CST 50-99:** Operações de crédito, suspensão, regimes especiais — valor zero
- **CST 99:** Outras Operações — valor zero

**Fonte:** `internal/legacy/pis_cofins.go:1-200`, `internal/legacy/pis_strategies.go:1-98`, `internal/legacy/cofins_strategies.go:1-77`

### RF-08 — Reforma Tributária (CBS/IBS/IS)

Implementado em 2026-06-21 com `ReformaCalculator` usando `iva_dual_rules` como fonte de alíquotas.

**Regra de lookup:** Prioriza regra específica do município sobre a regra padrão estadual (`ORDER BY municipio_destino_ibge DESC`).

**Cálculo por item:**
1. `fatorReducao = 1 − (percentualReducao / 100)`
2. CBS = `valor_item × (aliquotaCBS × fatorReducao / 100)`
3. IBS = `valor_item × ((aliquotaIBSEstadual + aliquotaIBSMunicipal) × fatorReducao / 100)`
4. IS = `valor_item × (aliquotaIS / 100)` — somente se `isImpostoSeletivo && aliquotaIS > 0`

**Redução de alíquotas:**
- `0%` → alíquotas integrais
- `60%` → reduzidas a 40% do nominal
- `100%` → isenção total (sem CBS/IBS)

**CSTs provisórios:** `01` (normal) e `04` (com redução) — aguardando tabela oficial da RFB.

**Fonte:** `internal/reforma/reforma.go:1-168`, `repository/postgres_repository.go:359-401`

### RF-09 — Cálculo de ISS (Imposto sobre Serviços)

**Regras:** Alíquota municipal no intervalo [2%, 5%] (BR-TAX-CONS-007). Identificação de serviços via `ITEM_LISTA_SERVICO`. Cálculo: `ISS = Preço_Serviço × Alíquota`. Retenção na fonte via flag `ISS_RETIDO`.

**Fonte:** `internal/legacy/iss.go:1-140`

### RF-10 — Cálculo de FUST/FUNTTEL (Contribuições de Telecom)

**Regras:** Apenas SCM/STFC (não SVA). Base líquida = Valor − ICMS − PIS − COFINS (BR-TAX-CALC-019/020). FUST = Base × 1% (Lei 9.998/2000). FUNTTEL = Base × 0,5% (Lei 10.052/2000). Executados na Fase 3 (pós-paralela) por dependência dos tributos da Fase 2.

**Fonte:** `internal/legacy/fust.go:1-120`, `internal/legacy/funttel.go:1-100`, `internal/legacy/telecom.go:1-90`

## 4. Requisitos Não-Funcionais

| RNF | Descrição | Evidência |
|-----|-----------|-----------|
| RNF-01 | Concorrência na Fase 2 via goroutines com WaitGroup | `engine.go:83-114` |
| RNF-02 | Logging estruturado JSON via slog com trace_id/request_id | `main.go:29-30`, `main.go:138-143` |
| RNF-03 | Conexão com PostgreSQL via pool pgx | `main.go:38-42` |
| RNF-04 | Cache Redis com decorator pattern via CachedTaxRepository | `main.go:46-47` |
| RNF-05 | Middleware recover para captura de panics | `main.go:72` |
| RNF-06 | Middleware logger para todas as requisições | `main.go:75` |
| RNF-07 | Matemática financeira de precisão com shopspring/decimal | `go.mod:9` |
| RNF-08 | Middleware requestid (W3C Trace Context) | `internal/middleware/requestid.go`, `main.go:73` |
| RNF-09 | Middleware auth JWT (Kong/Keycloak pass-through) | `internal/middleware/auth.go`, `main.go:74` |
| RNF-10 | Métricas Prometheus (http, cache, errors) | `internal/middleware/metrics.go`, `main.go:77-78` |
| RNF-11 | Porta configurável via env var PORT | `main.go:169-175` |
| RNF-12 | Health checks: `/healthz` (liveness) + `/health` (readiness) | `main.go:81-116` |
| RNF-13 | Coleta não-bloqueante de erros em goroutines (Fase 2) | `engine.go:86-114` |
| RNF-14 | Motor trifásico com Fase 3 pós-paralela (FUST/FUNTTEL) | `engine.go:127-170` |
| RNF-15 | Injeção reversa de tributos calculados (Fase 2 → input) | `engine.go:127-148` |
