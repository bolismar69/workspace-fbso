# Procedimentos Operacionais: Cálculo de Impostos Corporativos

**Código:** PROCEDURE-FIN-00001  
**Versão:** 1.1  
**Política vinculada:** [POLICE-FIN-00001](./POLICE-FIN-00001-CALCULO-DE-IMPOSTOS-CORPORATIVOS.md)  
**Regras de Negócio:** [RULES-CATALOG-FIN-00001](./RULES-CATALOG-FIN-00001.md)  
**Owner:** Tax Compliance Officer + Engineering Lead  
**Última Atualização:** 2026-06-21  
**Cadência de Revisão:** Semestral ou por evento (nova legislação)  
**Aprovadores:** CFO · Controller · Head of Engineering · Comitê Fiscal  

---

## 1. Propósito

Este documento define os procedimentos operacionais padrão (SOPs) necessários para executar, validar e manter o cálculo de tributos corporativos da empresa de **Telecomunicações**, em conformidade com a POLICE-FIN-00001 e o RULES-CATALOG-FIN-00001.

Cada SOP descreve **quem** executa, **quando** executa, **como** executa e **qual** o output esperado. Os procedimentos cobrem o ciclo completo: cálculo por operação, cálculo por período, reconciliação, gestão de exceções, atualização de alíquotas e contingência.

---

## 2. Escopo

### 2.1 Procedimentos Cobertos

| ID | Procedimento | Tipo | Gatilho | Engine |
|:---|:---|:---|:---|:---|
| SOP-001 | Cálculo de CBS "por fora" em serviços de telecom | Cálculo por Operação | Cada NF-e / operação de saída | Billing Engine — Reforma |
| SOP-002 | Consulta e cache de alíquotas IBS por destino | Infraestrutura | Cada operação com destino preenchido | Billing Engine — Reforma |
| SOP-003 | Classificação NCM e avaliação de incidência do IS | Pré-cálculo | Antes do cálculo da CBS | Billing Engine — Reforma |
| SOP-004 | Reconciliação mensal Shadow Run | Validação | Dia 5 de cada mês | Period Engine + Controller |
| SOP-005 | Atualização de alíquotas por nova legislação | Manutenção | Publicação de nova lei/decreto | Todos os engines |
| SOP-006 | Tratamento de exceções e creditamento | Fechamento | Apuração mensal/trimestral | Period Engine |
| SOP-007 | Cálculo de IRPJ/CSLL (Lucro Real e Presumido) | Cálculo por Período | Fechamento trimestral/anual | Period Engine — Lucratividade |
| SOP-008 | Cálculo de PIS/COFINS (Regime Não-Cumulativo) | Cálculo por Operação | Cada NF-e | Billing Engine — Legacy |
| SOP-009 | Cálculo de ICMS próprio, DIFAL e ICMS-ST | Cálculo por Operação | Cada NF-e | Billing Engine — Legacy |
| SOP-010 | Cálculo de ISS sobre serviços de telecom | Cálculo por Operação | Cada NF-e de serviço | Billing Engine — Legacy |
| SOP-011 | Cálculo de IPI sobre equipamentos | Cálculo por Operação | Cada NF-e com mercadoria | Billing Engine — Legacy |
| SOP-012 | Cálculo de CPP e FGTS (Encargos de Folha) | Cálculo por Período | Fechamento mensal | Period Engine — Folha |
| SOP-013 | Execução do pipeline de cálculo por fase | Orquestração | Cada requisição ao motor | Tax Pipeline |
| SOP-014 | Tratamento de contingência (circuit breaker IBS) | Contingência | Falha na API do Comitê Gestor | Billing Engine — Reforma |
| SOP-015 | Compensação cruzada de tributos (PER/DCOMP) | Fechamento | Quando há crédito excedente | Period Engine + Tax Officer |
| SOP-016 | Cálculo de FUST e FUNTTEL (Contribuições de Telecom) | Cálculo por Operação | Cada NF-e de serviço de telecom | Billing Engine — Telecom |
| SOP-017 | Cálculo de ICMS Desonerado | Cálculo por Operação | Cada NF-e com benefício fiscal de ICMS | Billing Engine — Legacy |

---

## 3. Matriz RACI Global

| Procedimento | R (Responsável) | A (Accountable) | C (Consultado) | I (Informado) |
|:---|:---|:---|:---|:---|
| SOP-001 — Cálculo CBS | `ms-billing-engine-tax-rates` | Engineering Lead | Tax Compliance Officer | Controller |
| SOP-002 — Consulta IBS | `ms-billing-engine-tax-rates` | Engineering Lead | Comitê Gestor IBS | Tax Compliance Officer |
| SOP-003 — Classificação IS | Pipeline de Dados + Engine | Engineering Lead | Consultoria Tributária | Tax Compliance Officer |
| SOP-004 — Reconciliação Shadow | Controller | Tax Compliance Officer | Engineering Lead | CFO, Comitê Fiscal |
| SOP-005 — Atualização Alíquotas | Engineering Lead | Tax Compliance Officer | Consultoria Tributária | Controller, CFO |
| SOP-006 — Exceções/Crédito | Engine + Controller | Tax Compliance Officer | Engineering Lead | CFO |
| SOP-007 — IRPJ/CSLL | Period Engine | Controller | Tax Compliance Officer | CFO |
| SOP-008 — PIS/COFINS | Billing Engine | Engineering Lead | Controller | Tax Compliance Officer |
| SOP-009 — ICMS/DIFAL/ST | Billing Engine | Engineering Lead | Controller | Tax Compliance Officer |
| SOP-010 — ISS | Billing Engine | Engineering Lead | Controller | Tax Compliance Officer |
| SOP-011 — IPI | Billing Engine | Engineering Lead | Controller | Tax Compliance Officer |
| SOP-012 — CPP/FGTS | Period Engine | Controller | RH Corporativo | Tax Compliance Officer |
| SOP-013 — Pipeline por Fase | Tax Pipeline | Engineering Lead | Tax Compliance Officer | Controller |
| SOP-014 — Contingência IBS | `ms-billing-engine-tax-rates` | Engineering Lead | Tax Compliance Officer | Comitê Fiscal |
| SOP-015 — PER/DCOMP | Controller | Tax Compliance Officer | Jurídico Corporativo | CFO |
| SOP-016 — FUST/FUNTTEL | `ms-billing-engine-tax-rates` | Engineering Lead | Tax Compliance Officer | Controller |
| SOP-017 — ICMS Desonerado | `ms-billing-engine-tax-rates` | Engineering Lead | Controller | Tax Compliance Officer |

---

# PROCEDIMENTOS DE CÁLCULO POR OPERAÇÃO

---

## SOP-001: Cálculo de CBS "por fora" em Serviços de Telecom

**Regras vinculadas:** BR-TAX-CALC-017, BR-TAX-CONS-008, BR-TAX-DEF-005  
**Engine:** Billing Engine — Reforma (`cbs_calculator.go`)

### Process Flow

```
┌─────────────────────────────────────────────────────────────────┐
│ SOP-001: Cálculo de CBS "por fora"                              │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  INÍCIO                                                         │
│    │                                                            │
│    ▼                                                            │
│  ┌──────────────────────────────┐                               │
│  │ 1. Receber TaxDocumentInput  │                               │
│  │    com valor_bruto,          │                               │
│  │    c_class_trib, fase ativa  │                               │
│  └──────────────┬───────────────┘                               │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────┐                               │
│  │ 2. Resolver Fase Tributária  │                               │
│  │    Ref: BR-TAX-INF-001 a 004 │                               │
│  └──────────────┬───────────────┘                               │
│                 │                                               │
│        ┌────────┴────────┐                                      │
│        ▼                 ▼                                      │
│  ┌──────────┐      ┌──────────────┐                             │
│  │ Shadow   │      │ CBS Plena /  │                             │
│  │ Run      │      │ Transição /  │                             │
│  │ (2026)   │      │ IVA Dual     │                             │
│  └────┬─────┘      └──────┬───────┘                             │
│       │                   │                                     │
│       ▼                   ▼                                     │
│  ┌──────────────┐   ┌──────────────────────┐                    │
│  │ Aliquota =   │   │ 3. Consultar tabela  │                    │
│  │ 0,001 (0,1%) │   │    cbs_rates por     │                    │
│  │              │   │    c_class_trib       │                    │
│  └──────┬───────┘   └──────────┬───────────┘                    │
│         │                      │                                │
│         └──────────┬───────────┘                                │
│                    ▼                                            │
│  ┌──────────────────────────────────────┐                       │
│  │ 4. Calcular:                         │                       │
│  │    CBS_Valor = Base × Aliquota_CBS   │                       │
│  │    (cálculo "por fora" — a CBS NÃO   │                       │
│  │     integra a própria base)           │                       │
│  └──────────────┬───────────────────────┘                       │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────┐                               │
│  │ 5. Se Shadow Run:            │                               │
│  │    → shadow_tax_result.cbs   │                               │
│  │    Se Produção:              │                               │
│  │    → TaxResponse.tributos[]  │                               │
│  │    → Compõe total_a_pagar    │                               │
│  └──────────────┬───────────────┘                               │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────┐                               │
│  │ 6. Registrar trilha de       │                               │
│  │    auditoria (log imutável)  │                               │
│  └──────────────┬───────────────┘                               │
│                 │                                               │
│                 ▼                                               │
│               FIM                                               │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Detailed Steps

#### Step 1: Receber Input da Operação
- **Who:** `CBSCalculator.Calculate()`
- **When:** Para cada item de cada NF-e de saída
- **How:** O motor recebe `TaxDocumentInput` contendo `valor_bruto`, `c_class_trib`, `data_operacao` e `itens[]`
- **Output:** Dados estruturados prontos para o cálculo

#### Step 2: Resolver Fase Tributária
- **Who:** `PhaseResolver.Resolve(data_operacao)`
- **When:** Antes de qualquer cálculo
- **How:** Conforme BR-TAX-INF-001 a 004, determina se a operação está em Shadow Run, CBS Plena, Transição Subnacional ou IVA Dual
- **Output:** `Phase` enum (`SHADOW_RUN`, `CBS_PLENA`, `TRANSICAO_SUBNACIONAL`, `IVA_DUAL`)

#### Step 3: Determinar Alíquota CBS
- **Who:** `CBSCalculator.GetRate(c_class_trib, phase)`
- **When:** Após resolver a fase
- **How:**
  - Se `phase == SHADOW_RUN`: retornar `0.001` (0,1% — alíquota de teste)
  - Se `phase >= CBS_PLENA`: consultar tabela SQL `billing_tax_rates.cbs_rates` filtrando por `c_class_trib` e `active = true`
  - Se alíquota setorial ainda não definida (Gap G1): retornar alíquota estimada de 0,12 (12%) com flag `rate_status = "ESTIMATED"`
- **Output:** `aliquota_cbs` (float64) e `rate_source` (string)

#### Step 4: Executar Cálculo "por fora"
- **Who:** `CBSCalculator.Calculate(base, aliquota)`
- **When:** Após obter a alíquota
- **How:**
  ```go
  cbsValor := base * aliquotaCBS
  // NOTA: A base NÃO é ajustada para incluir a CBS (cálculo "por fora")
  ```
- **Constraint:** BR-TAX-CONS-008 — A base usada para CBS não é acrescida da própria CBS
- **Output:** `cbs_valor` (float64)

#### Step 5: Roteamento do Resultado
- **Who:** `TaxResponseBuilder.Build()`
- **When:** Após cada cálculo
- **How:**
  - Se `phase == SHADOW_RUN`: armazenar em `shadow_tax_result.cbs` (não compor `total_a_pagar`)
  - Se `phase != SHADOW_RUN`: armazenar em `TaxResponse.tributos[]` e compor `total_a_pagar`
- **Output:** Resultado roteado para o campo correto

#### Step 6: Auditoria
- **Who:** `AuditLogger.Log()`
- **When:** Imediatamente após cada cálculo
- **How:** Registrar JSON imutável com timestamp, transaction_id, tax_type, base, rate, calculated_amount, rate_source, engine_version, git_commit
- **Output:** Registro de auditoria em log append-only

### Exceptions and Edge Cases

| Cenário | Ação |
|:---|:---|
| `c_class_trib` não encontrado na tabela `cbs_rates` | Log WARN + usar alíquota default do setor TELECOM (12%) + notificar Engineering Lead |
| Alíquota retornada = 0 (isento) | Registrar em auditoria com `rate_source = "EXEMPT"` + `cbs_valor = 0` |
| `data_operacao` em ano de transição (ex: 31/12/2026 vs 01/01/2027) | Usar data exata para resolução de fase; não arredondar para o ano fiscal |
| `valor_bruto` negativo (devolução/estorno) | Inverter sinal do CBS_Valor; registrar `tax_type = "CBS_DEVOLUCAO"` |

---

## SOP-002: Consulta e Cache de Alíquotas IBS por Destino

**Regras vinculadas:** BR-TAX-CALC-018, BR-TAX-CONS-009, BR-TAX-DEF-006, BR-TAX-ACT-001  
**Engine:** Billing Engine — Reforma (`ibs_client.go`, `ibs_cache.go`)

### Process Flow

```
┌─────────────────────────────────────────────────────────────────┐
│ SOP-002: Consulta e Cache de Alíquotas IBS                     │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  INÍCIO (para cada operação com destino.municipio_codigo_ibge)  │
│    │                                                            │
│    ▼                                                            │
│  ┌──────────────────────────────────┐                           │
│  │ 1. Montar chave de cache:        │                           │
│  │    ibs:rate:{ibge_code}:{date}   │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────┐                               │
│  │ 2. Verificar Redis Cache     │                               │
│  └──────────────┬───────────────┘                               │
│                 │                                               │
│        ┌────────┴────────┐                                      │
│        ▼                 ▼                                      │
│  ┌──────────┐      ┌──────────┐                                 │
│  │ Cache    │      │ Cache    │                                 │
│  │ HIT      │      │ MISS     │                                 │
│  └────┬─────┘      └────┬─────┘                                 │
│       │                 │                                       │
│       │                 ▼                                       │
│       │    ┌────────────────────────┐                           │
│       │    │ 3. Chamar API Comitê   │                           │
│       │    │    Gestor IBS:         │                           │
│       │    │    GET /api/v1/rates   │                           │
│       │    │    ?ibge_code={code}   │                           │
│       │    └───────────┬────────────┘                           │
│       │                │                                       │
│       │       ┌────────┴────────┐                               │
│       │       ▼                 ▼                               │
│       │  ┌──────────┐     ┌──────────────┐                      │
│       │  │ Sucesso   │     │ Falha        │                      │
│       │  │ (200 OK)  │     │ (4xx/5xx)    │                      │
│       │  └────┬─────┘     └──────┬───────┘                      │
│       │       │                  │                              │
│       │       │                  ▼                              │
│       │       │     ┌────────────────────────┐                  │
│       │       │     │ 4. Incrementar contador │                  │
│       │       │     │    de falhas.           │                  │
│       │       │     │    Se ≥ 3 falhas em 60s │                  │
│       │       │     │    → ABRIR Circuit      │                  │
│       │       │     │      Breaker            │                  │
│       │       │     │    → SOP-014            │                  │
│       │       │     └────────────────────────┘                  │
│       │       │                                                │
│       │       ▼                                                │
│       │  ┌────────────────────────┐                            │
│       │  │ 5. Split alíquota:     │                            │
│       │  │    total = estadual    │                            │
│       │  │    + municipal         │                            │
│       │  └───────────┬────────────┘                            │
│       │              │                                         │
│       │              ▼                                         │
│       │  ┌────────────────────────┐                            │
│       │  │ 6. Armazenar no Redis  │                            │
│       │  │    TTL = 24h           │                            │
│       │  └───────────┬────────────┘                            │
│       │              │                                         │
│       └──────────────┘                                         │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────┐                               │
│  │ 7. Retornar Aliquota_Total   │                               │
│  │    Para uso em BR-TAX-CALC-018                              │
│  └──────────────┬───────────────┘                               │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────┐                               │
│  │ 8. Registrar em trilha de    │                               │
│  │    auditoria (rate_source)   │                               │
│  └──────────────────────────────┘                               │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Detailed Steps

#### Step 1: Montar Chave de Cache
- **Who:** `IBSCache.Key(ibgeCode, date)`
- **When:** Para cada operação com `destino.municipio_codigo_ibge` preenchido
- **How:** `fmt.Sprintf("ibs:rate:%s:%s", ibgeCode, date.Format("2006-01-02"))`
- **Output:** Chave de cache string

#### Step 2: Verificar Cache
- **Who:** `IBSCache.Get(key)`
- **When:** Antes de chamar a API externa
- **How:** `redisClient.Get(ctx, key).Result()`
- **Output:** `IBSRate` (cache hit) ou `redis.Nil` (cache miss)

#### Step 3: Chamar API do Comitê Gestor IBS
- **Who:** `IBSClient.FetchRate(ibgeCode)`
- **When:** Cache miss
- **How:**
  ```go
  url := fmt.Sprintf("%s/api/v1/rates?ibge_code=%s", ibsBaseURL, ibgeCode)
  resp, err := httpClient.Get(url) // Timeout: 5s
  ```
- **Status:** ⚠️ Endpoint real ainda não publicado (Gap G2); usar mock em staging

#### Step 4: Tratamento de Falha
- **Who:** `IBSClient` + Circuit Breaker
- **When:** API retornar erro (4xx, 5xx, timeout)
- **How:** Incrementar contador de falhas. Se ≥ 3 falhas em janela de 60s → abrir circuit breaker (SOP-014)
- **Output:** Estado do circuit breaker + fallback

#### Step 5: Split da Alíquota
- **Who:** `IBSRate.Parse()`
- **When:** Após resposta bem-sucedida da API
- **How:** `aliquota_total = response.Estadual + response.Municipal`
- **Output:** `IBSRate{Estadual, Municipal, Total, FetchedAt}`

#### Step 6: Armazenar no Cache
- **Who:** `IBSCache.Set(key, rate)`
- **When:** Após obter a alíquota (cache hit ou fetch)
- **How:** `redisClient.Set(ctx, key, rateJSON, 24*time.Hour)`
- **Constraint:** BR-TAX-CONS-009 — TTL máximo = 24h
- **Output:** Confirmação de cache armazenado

#### Step 7: Retornar Alíquota
- **Who:** `IBSClient.GetRate(ibgeCode)`
- **When:** Como passo final do fluxo
- **How:** Retornar `IBSRate.Total` para uso em `BR-TAX-CALC-018`
- **Output:** `aliquota_total_ibs` (float64)

#### Step 8: Auditoria
- **Who:** `AuditLogger.Log()`
- **When:** Imediatamente após resolver a alíquota
- **How:** Registrar JSON com `rate_source` ("CACHE" ou "API_COMITE_GESTOR"), `ibge_code`, `aliquota_estadual`, `aliquota_municipal`, `aliquota_total`, `fetched_at`
- **Output:** Registro de auditoria

### Exceptions and Edge Cases

| Cenário | Ação |
|:---|:---|
| Cache com TTL expirado (24h+) | Reconsultar API automaticamente; não usar cache expirado |
| API retorna alíquota zero | Validar: consultar `is_exempt` flag. Se confirmado isento, aceitar zero. Senão, log ERROR + notificar |
| `ibge_code` inválido ou vazio | Log ERROR + retornar erro para o caller; não calcular IBS |
| API retorna alíquota > 50% (anômalo) | Log CRITICAL + não cachear + notificar Tax Compliance Officer imediatamente |
| Circuit breaker aberto | Usar última alíquota cacheada (mesmo que expirada) com flag `rate_source = "FALLBACK_CIRCUIT_OPEN"` |

---

## SOP-003: Classificação NCM e Avaliação de Incidência do IS

**Regras vinculadas:** BR-TAX-INF-005, BR-TAX-CONS-010  
**Engine:** Billing Engine — Reforma (`is_filter.go`)

### Process Flow

```
┌─────────────────────────────────────────────────────────────────┐
│ SOP-003: Classificação NCM e IS                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  INÍCIO (antes do cálculo da CBS para cada item)                │
│    │                                                            │
│    ▼                                                            │
│  ┌──────────────────────────────────┐                           │
│  │ 1. Extrair NCM e isento_is       │                           │
│  │    do TaxItem                    │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────┐                               │
│  │ 2. Verificar flag isento_is  │                               │
│  └──────────────┬───────────────┘                               │
│                 │                                               │
│        ┌────────┴────────┐                                      │
│        ▼                 ▼                                      │
│  ┌──────────┐      ┌──────────────┐                             │
│  │ isento_is│      │ isento_is    │                             │
│  │ = true   │      │ = false      │                             │
│  └────┬─────┘      └──────┬───────┘                             │
│       │                   │                                     │
│       ▼                   ▼                                     │
│  ┌──────────────┐   ┌──────────────────────┐                    │
│  │ IS = 0       │   │ 3. Consultar tabela  │                    │
│  │ (isento)     │   │    ncm_seletivo      │                    │
│  └──────┬───────┘   └──────────┬───────────┘                    │
│         │                      │                                │
│         │             ┌────────┴────────┐                       │
│         │             ▼                 ▼                       │
│         │        ┌──────────┐     ┌──────────┐                  │
│         │        │ NCM está │     │ NCM não  │                  │
│         │        │ na tabela│     │ está     │                  │
│         │        └────┬─────┘     └────┬─────┘                  │
│         │             │               │                         │
│         │             ▼               ▼                         │
│         │        ┌──────────┐   ┌──────────┐                    │
│         │        │ IS =     │   │ IS = 0   │                    │
│         │        │ Base ×   │   │ (não     │                    │
│         │        │ Aliq_IS  │   │ incide)  │                    │
│         │        └────┬─────┘   └────┬─────┘                    │
│         │             │               │                         │
│         └─────────────┴───────────────┘                         │
│                        │                                        │
│                        ▼                                        │
│  ┌──────────────────────────────────┐                           │
│  │ 4. Registrar em trilha de        │                           │
│  │    auditoria (mesmo se IS = 0)   │                           │
│  │    com flag is_exempt            │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────┐                               │
│  │ 5. Prosseguir para CBS e IBS │                               │
│  │    (ordem: IS → CBS → IBS)   │                               │
│  └──────────────────────────────┘                               │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Detailed Steps

#### Step 1: Extrair Dados do Item
- **Who:** `ISFilter.Extract(item)`
- **When:** Para cada item antes do pipeline de cálculo da Reforma
- **How:** Ler `item.NCM` e `item.IsentoIS` (bool)
- **Output:** Dados de classificação do item

#### Step 2: Verificar Override Manual
- **Who:** `ISFilter.CheckExemption(item)`
- **When:** Antes de consultar a tabela NCM
- **How:** Se `item.isento_is == true` → IS = 0 independente do NCM. Este é um override manual para itens comprovadamente isentos.
- **Output:** Decisão de isenção

#### Step 3: Consultar Tabela NCM Seletivo
- **Who:** `ISFilter.LookupNCM(ncm)`
- **When:** Se `isento_is == false`
- **How:** `SELECT aliquota_is FROM billing_tax_rates.ncm_seletivo WHERE ncm = $1 AND active = true`
- **Status:** ⚠️ Tabela oficial ainda não publicada (Gap G4); usar tabela preliminar com NCMs de bebidas, tabaco, açúcar, veículos e combustíveis
- **Output:** `aliquota_is` (float64) ou `sql.ErrNoRows` (não incide)

#### Step 4: Auditoria
- **Who:** `AuditLogger.Log()`
- **When:** Após determinação do IS (mesmo se zero)
- **How:** Registrar JSON com `ncm`, `is_exempt`, `is_rate`, `is_amount`, `lookup_source`
- **Output:** Registro de auditoria

#### Step 5: Prosseguir Pipeline
- **Who:** `TaxPipeline.Execute()`
- **When:** Após IS calculado
- **How:** Garantir ordem completa do pipeline conforme BR-TAX-CONS-010 (IS → IPI → CBS → ICMS → IBS → ISS → PIS/COFINS → FUST → FUNTTEL).
- **Output:** Próximo cálculo no pipeline

### Exceptions and Edge Cases

| Cenário | Ação |
|:---|:---|
| NCM não encontrado na tabela local | Log WARN + assumir não-incidência + agendar sincronização da tabela |
| Alíquota IS > 100% (anômalo) | Log CRITICAL + não aplicar + notificar Tax Compliance Officer |
| Item sem NCM (serviços puros) | IS = 0 (serviços não têm NCM e não são sujeitos ao IS) |
| Sincronização semanal da tabela falhou | Usar última tabela válida (cache em arquivo) + notificar Engenharia de Dados |

---

## SOP-008: Cálculo de PIS/COFINS (Regime Não-Cumulativo)

**Regras vinculadas:** BR-TAX-CALC-007, BR-TAX-CALC-009, BR-TAX-DEF-004, BR-TAX-CONS-002  
**Engine:** Billing Engine — Legacy (`pis_cofins_calculator.go`)

### Process Flow

```
┌─────────────────────────────────────────────────────────────────┐
│ SOP-008: Cálculo de PIS/COFINS Não-Cumulativo                  │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  INÍCIO (para cada item da NF-e)                                │
│    │                                                            │
│    ▼                                                            │
│  ┌──────────────────────────────────┐                           │
│  │ 1. Receber regime_tributario e   │                           │
│  │    cst_pis, cst_cofins do item   │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────┐                               │
│  │ 2. Validar regime TELECOM    │                               │
│  │    Deve ser NAO_CUMULATIVO   │                               │
│  │    Ref: BR-TAX-CONS-002      │                               │
│  └──────────────┬───────────────┘                               │
│                 │                                               │
│        ┌────────┴────────┐                                      │
│        ▼                 ▼                                      │
│  ┌──────────┐      ┌──────────────┐                             │
│  │ Regime   │      │ Regime       │                             │
│  │ OK       │      │ Inválido     │                             │
│  └────┬─────┘      └──────┬───────┘                             │
│       │                   │                                     │
│       │                   ▼                                     │
│       │           ┌──────────────┐                              │
│       │           │ Retornar     │                              │
│       │           │ ERRO + log   │                              │
│       │           └──────────────┘                              │
│       │                                                         │
│       ▼                                                         │
│  ┌──────────────────────────────────┐                           │
│  │ 3. Determinar CST e Alíquota     │                           │
│  │    CST_Tributado: 01 → 1,65%/7,6%                          │
│  │    CST_Credito:   50 → apurar crédito                       │
│  │    CST_Isento:    02,73 → 0%                                │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────┐                               │
│  │ 4. Calcular:                 │                               │
│  │    PIS = Base × 0,0165       │                               │
│  │    COFINS = Base × 0,076     │                               │
│  └──────────────┬───────────────┘                               │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────┐                               │
│  │ 5. Se CST = 50 (crédito):    │                               │
│  │    → acumular em creditos    │                               │
│  │    → abater débito no mês    │                               │
│  └──────────────┬───────────────┘                               │
│                 │                                               │
│                 ▼                                               │
│               FIM                                               │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Detailed Steps

#### Step 1: Receber Dados
- **Who:** `PISCOFINSCalculator.Calculate(input)`
- **When:** Cada item da NF-e
- **How:** Receber `regime_tributario`, `cst_pis`, `cst_cofins`, `valor` do item
- **Output:** Dados estruturados

#### Step 2: Validar Regime (TELECOM)
- **Who:** `RegimeValidator.Validate(regime, cnae)`
- **When:** Sempre, antes do cálculo
- **How:** Se CNAE é de telecomunicações, `regime_tributario` DEVE ser `NAO_CUMULATIVO` (BR-TAX-CONS-002)
- **Output:** Validação (ok/erro)

#### Step 3: Determinar Situação Tributária
- **Who:** `CSTResolver.ResolvePisCofins(cst)`
- **When:** Após validação do regime
- **How:**
  - `CST_PIS = "01"` → Alíquota PIS = 1,65% (tributado)
  - `CST_COFINS = "01"` → Alíquota COFINS = 7,6% (tributado)
  - `CST_PIS = "50"` → Crédito PIS (1,65%)
  - `CST_COFINS = "50"` → Crédito COFINS (7,6%)
  - `CST = "02", "73"` → Isento (0%)
- **Output:** Alíquotas PIS e COFINS aplicáveis

#### Step 4: Calcular
- **Who:** `PISCOFINSCalculator.Calculate()`
- **When:** Após determinar alíquotas
- **How:**
  ```go
  pis := base * 0.0165   // BR-TAX-CALC-007
  cofins := base * 0.076 // BR-TAX-CALC-009
  ```
- **Output:** `pis_valor`, `cofins_valor`

#### Step 5: Acumular Créditos
- **Who:** `CreditAccumulator.Add(tributo, valor)`
- **When:** Se CST = 50
- **How:** Acumular no extrato de créditos do período para abatimento na apuração mensal
- **Output:** Saldo de créditos atualizado

### Exceptions

| Cenário | Ação |
|:---|:---|
| CST não reconhecido | Log ERROR + não calcular + notificar Engineering Lead |
| Alíquota zero mas CST = 01 | Log WARN — inconsistência; usar alíquota do regime |
| Base negativa (devolução) | Inverter sinal + CST específico para devolução |

---

## SOP-009: Cálculo de ICMS Próprio, DIFAL e ICMS-ST

**Regras vinculadas:** BR-TAX-CALC-010 a 014, BR-TAX-CONS-004 a 006, DT-002  
**Engine:** Billing Engine — Legacy (`icms_calculator.go`, `icms_st_calculator.go`)

### Process Flow

```
┌─────────────────────────────────────────────────────────────────┐
│ SOP-009: Cálculo de ICMS (Próprio, DIFAL, ST)                  │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  INÍCIO (para cada item da NF-e)                                │
│    │                                                            │
│    ▼                                                            │
│  ┌──────────────────────────────────────┐                       │
│  │ 1. Compor Base de Cálculo:           │                       │
│  │    Valor + Frete + Seguro +          │                       │
│  │    Outras Despesas + IPI (se CF)     │                       │
│  │    Ref: BR-TAX-CALC-010              │                       │
│  └──────────────┬───────────────────────┘                       │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────────┐                       │
│  │ 2. Determinar Alíquota ICMS          │                       │
│  │    Intra-estadual: alíquota UF       │                       │
│  │    Interestadual: 7% ou 12%          │                       │
│  │    Ref: BR-TAX-CONS-004              │                       │
│  └──────────────┬───────────────────────┘                       │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────────┐                       │
│  │ 3. Calcular ICMS Próprio             │                       │
│  │    Ref: BR-TAX-CALC-011              │                       │
│  └──────────────┬───────────────────────┘                       │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────────┐                       │
│  │ 4. Verificar DIFAL (DT-002)          │                       │
│  │    UF_Origem != UF_Destino?          │                       │
│  │    Consumidor Final?                 │                       │
│  │    AliqIntDest > AliqInterestadual?  │                       │
│  └──────────────┬───────────────────────┘                       │
│                 │                                               │
│        ┌────────┴────────┐                                      │
│        ▼                 ▼                                      │
│  ┌──────────┐      ┌──────────┐                                 │
│  │ DIFAL    │      │ DIFAL    │                                 │
│  │ Aplicável│      │ = 0      │                                 │
│  └────┬─────┘      └────┬─────┘                                 │
│       │                 │                                       │
│       ▼                 │                                       │
│  ┌──────────────────┐   │                                       │
│  │ 5. Calcular DIFAL│   │                                       │
│  │ Ref:BR-TAX-CALC-012  │                                       │
│  └────┬─────────────┘   │                                       │
│       │                 │                                       │
│       └─────────┬───────┘                                       │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────────┐                       │
│  │ 6. Verificar ICMS-ST                  │                       │
│  │    CEST preenchido? MVA > 0?         │                       │
│  └──────────────┬───────────────────────┘                       │
│                 │                                               │
│        ┌────────┴────────┐                                      │
│        ▼                 ▼                                      │
│  ┌──────────┐      ┌──────────┐                                 │
│  │ ST       │      │ ST não   │                                 │
│  │ Aplicável│      │ se aplica│                                 │
│  └────┬─────┘      └──────────┘                                 │
│       │                                                         │
│       ▼                                                         │
│  ┌──────────────────────────────────────┐                       │
│  │ 7. Calcular ICMS-ST:                 │                       │
│  │    BaseST = Valor × (1 + MVA)        │                       │
│  │    Ref: BR-TAX-CALC-013              │                       │
│  │    ICMS_ST = (BaseST × AliqIntDest)  │                       │
│  │            - ICMS_Proprio             │                       │
│  │    Ref: BR-TAX-CALC-014              │                       │
│  └──────────────────────────────────────┘                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Detailed Steps

#### Step 1: Compor Base de Cálculo
- **Who:** `ICMSCalculator.CalcBase(item)`
- **When:** Primeiro passo do cálculo ICMS
- **How:** `base = item.Valor + item.Frete + item.Seguro + item.OutrasDespesas + (item.IPI if consumidorFinal)`
- **Output:** `base_icms` (float64)

#### Step 2: Determinar Alíquota
- **Who:** `ICMSResolver.GetAliq(ufOrigem, ufDestino)`
- **When:** Após compor a base
- **How:**
  - Operação intra-estadual: alíquota interna da UF
  - Operação interestadual:
    - 4% se `Origem_Mercadoria = 1` (importada) e `Conteudo_Importacao > 40%` (Resolução Senado 13/2012)
    - 7% se Sul/Sudeste (exceto ES) para outras regiões; ou para ES como destino
    - 12% nos demais casos
  - Ref: BR-TAX-CONS-004
- **Output:** `aliquota_icms` (float64)

#### Step 3: Calcular ICMS Próprio
- **Who:** `ICMSCalculator.Calculate(base, aliquota)`
- **When:** Após determinar alíquota
- **How:** `icmsProprio = base * aliquota` (BR-TAX-CALC-011)
- **Output:** `icms_proprio` (float64)

#### Step 4: Verificar DIFAL (DT-002)
- **Who:** `ICMSCalculator.ShouldApplyDIFAL(origem, destino, consumidorFinal)`
- **When:** Após ICMS próprio
- **How:** 3 condições: interestadual + consumidor final + alíquota destino > interestadual
- **Output:** Booleano + alíquotas para cálculo

#### Step 5: Calcular DIFAL
- **Who:** `ICMSCalculator.CalcularDIFAL(base, aliqDestino, aliqInterestadual)`
- **When:** Se DIFAL aplicável
- **How:** `difal = (base * aliqDestino) - (base * aliqInterestadual)` (BR-TAX-CALC-012)
- **Output:** `difal_valor` (float64)

#### Step 6-7: Verificar e Calcular ICMS-ST
- **Who:** `ICMSSTCalculator.Calculate(item, icmsProprio)`
- **When:** Se `item.CEST` preenchido e `item.MVA > 0`
- **How:**
  - `baseST = item.Valor * (1 + item.MVA/100)` (BR-TAX-CALC-013)
  - `icmsST = (baseST * aliqDestino) - icmsProprio` (BR-TAX-CALC-014)
- **Output:** `icms_st_valor` (float64)

### Exceptions

| Cenário | Ação |
|:---|:---|
| Alíquota interestadual não mapeada para o par UF origem-destino | Log ERROR + usar 12% como default conservador + notificar |
| CEST preenchido mas MVA = 0 | Log WARN + verificar se o CEST realmente está na ST (consultar CONFAZ) |
| FCP estadual > 2% | Log WARN + validar legislação do estado destino antes de aplicar |

---

## SOP-010: Cálculo de ISS sobre Serviços de Telecom

**Regras vinculadas:** BR-TAX-CALC-016, BR-TAX-CONS-007, BR-TAX-DEF-007  
**Engine:** Billing Engine — Legacy (`iss_calculator.go`)

### Detailed Steps

#### Step 1: Identificar Item como Serviço
- **Who:** `ISSCalculator.IsService(item)`
- **When:** Primeiro passo
- **How:** Verificar `item.ItemListaServico` preenchido (ex: "1.05" para telecom)
- **Output:** Booleano

#### Step 2: Validar Alíquota Municipal
- **Who:** `ISSValidator.ValidateAliquota(aliq)`
- **When:** Antes do cálculo
- **How:** `2% <= aliquota <= 5%` (BR-TAX-CONS-007)
- **Output:** Validação (ok/erro)

#### Step 3: Calcular ISS
- **Who:** `ISSCalculator.Calculate(precoServico, aliquota)`
- **When:** Após validação
- **How:** `iss = precoServico * aliquota` (BR-TAX-CALC-016)
- **Output:** `iss_valor`

#### Step 4: Verificar Retenção
- **Who:** `ISSCalculator.CheckRetencao(item)`
- **When:** Após cálculo
- **How:** Se `item.RetencaoISS == true` → o ISS foi retido pelo tomador; valor informado, não recolhido pelo prestador
- **Output:** Flag `iss_retido`

---

## SOP-011: Cálculo de IPI sobre Equipamentos

**Regras vinculadas:** BR-TAX-CALC-015, BR-TAX-CONS-003  
**Engine:** Billing Engine — Legacy (`ipi_calculator.go`)

### Detailed Steps

#### Step 1: Verificar Incidência
- **Who:** `IPICalculator.IsApplicable(item)`
- **When:** Primeiro passo
- **How:** Se o item é serviço puro de telecom → IPI não incide (BR-TAX-CONS-003). Se mercadoria/equipamento → prosseguir.
- **Output:** Booleano

#### Step 2: Consultar TIPI por NCM
- **Who:** `TIPILookup.GetRate(ncm)`
- **When:** Se IPI aplicável
- **How:** `SELECT aliquota_ipi FROM billing_tax_rates.tipi WHERE ncm = $1 AND active = true`
- **Output:** `aliquota_ipi`

#### Step 3: Calcular IPI
- **Who:** `IPICalculator.Calculate(valor, aliquota)`
- **When:** Após obter alíquota
- **How:** `ipi = valor * aliquota` (BR-TAX-CALC-015)
- **Output:** `ipi_valor`

---

## SOP-016: Cálculo de FUST e FUNTTEL (Contribuições de Telecom)

**Regras vinculadas:** BR-TAX-CALC-019, BR-TAX-CALC-020, BR-TAX-CONS-012, BR-TAX-INF-007, BR-TAX-DEF-010, BR-TAX-DEF-011  
**Engine:** Billing Engine — Telecom (`fust_calculator.go`, `funttel_calculator.go`)

### Process Flow

```
┌─────────────────────────────────────────────────────────────────┐
│ SOP-016: Cálculo de FUST e FUNTTEL                              │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  INÍCIO (para cada item de serviço da NF-e)                     │
│    │                                                            │
│    ▼                                                            │
│  ┌──────────────────────────────────┐                           │
│  │ 1. Identificar natureza do       │                           │
│  │    serviço (SCM, STFC, SVA)      │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│        ┌────────┴────────┐                                      │
│        ▼                 ▼                                      │
│  ┌──────────┐      ┌──────────────┐                             │
│  │ SCM/STFC │      │ SVA          │                             │
│  │ (incide) │      │ (não incide) │                             │
│  └────┬─────┘      └──────┬───────┘                             │
│       │                   │                                     │
│       │                   ▼                                     │
│       │            ┌──────────────┐                             │
│       │            │ FUST = 0     │                             │
│       │            │ FUNTTEL = 0  │                             │
│       │            └──────────────┘                             │
│       │                                                         │
│       ▼                                                         │
│  ┌──────────────────────────────────┐                           │
│  │ 2. Obter valores já calculados:  │                           │
│  │    ICMS, PIS, COFINS do item     │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 3. Calcular Base Líquida:        │                           │
│  │    Base = Valor_Serviço          │                           │
│  │         − ICMS                   │                           │
│  │         − PIS                    │                           │
│  │         − COFINS                 │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 4. Calcular FUST:                │                           │
│  │    FUST = Base × 0,01 (1%)       │                           │
│  │    Ref: BR-TAX-CALC-019          │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 5. Calcular FUNTTEL:             │                           │
│  │    FUNTTEL = Base × 0,005 (0,5%) │                           │
│  │    Ref: BR-TAX-CALC-020          │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 6. Registrar em trilha de        │                           │
│  │    auditoria                     │                           │
│  └──────────────────────────────────┘                           │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Detailed Steps

#### Step 1: Identificar Natureza do Serviço
- **Who:** `TelecomClassifier.Classify(item)`
- **When:** Para cada item classificado como serviço
- **How:** Verificar `item.NaturezaServico` ∈ {SCM, STFC, SVA}. Conforme BR-TAX-INF-007.
- **Output:** Classificação do serviço

#### Step 2: Obter Valores já Calculados
- **Who:** FUST/FUNTTEL Calculator
- **When:** Após ICMS, PIS e COFINS terem sido calculados (ordem do pipeline)
- **How:** Receber do contexto de cálculo: `icms_valor`, `pis_valor`, `cofins_valor`
- **Output:** Valores para compor a base líquida

#### Step 3: Calcular Base Líquida
- **Who:** `FUSTCalculator.CalcBase(valorServico, icms, pis, cofins)`
- **When:** Após obter os impostos
- **How:** `base = valorServico - icms - pis - cofins` (BR-TAX-CALC-019)
- **Output:** `base_fust_funttel` (float64)

#### Step 4: Calcular FUST
- **Who:** `FUSTCalculator.Calculate(base)`
- **When:** Após calcular a base
- **How:** `fust = base * 0.01` (BR-TAX-CALC-019)
- **Output:** `fust_valor`

#### Step 5: Calcular FUNTTEL
- **Who:** `FUNTTELCalculator.Calculate(base)`
- **When:** Após FUST (usa mesma base)
- **How:** `funttel = base * 0.005` (BR-TAX-CALC-020)
- **Output:** `funttel_valor`

### Exceptions and Edge Cases

| Cenário | Ação |
|:---|:---|
| SVA com FUST/FUNTTEL > 0 | Log ERROR — configuração incorreta do item; zerar FUST/FUNTTEL |
| Base líquida negativa (impostos > valor serviço) | FUST = 0; FUNTTEL = 0; log WARN |
| Item sem natureza de serviço definida | Assumir SVA (não incide) + log WARN + solicitar correção cadastral |
| Serviço com isenção judicial de FUST/FUNTTEL | Flag `isento_fust_funttel` no item → pular cálculo |

---

## SOP-017: Cálculo de ICMS Desonerado

**Regras vinculadas:** BR-TAX-CALC-021, BR-TAX-CALC-022, BR-TAX-CONS-013, BR-TAX-ACT-007  
**Engine:** Billing Engine — Legacy (`icms_calculator.go`)

### Process Flow

```
┌─────────────────────────────────────────────────────────────────┐
│ SOP-017: Cálculo de ICMS Desonerado                             │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  INÍCIO (quando item possui benefício fiscal de ICMS)           │
│    │                                                            │
│    ▼                                                            │
│  ┌──────────────────────────────────┐                           │
│  │ 1. Validar CST permite           │                           │
│  │    desoneração (≠ 00)            │                           │
│  │    Ref: BR-TAX-CONS-013          │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│        ┌────────┴────────┐                                      │
│        ▼                 ▼                                      │
│  ┌──────────┐      ┌──────────────┐                             │
│  │ CST OK   │      │ CST inválido │                             │
│  │ (20,30,  │      │ (ex: 00)     │                             │
│  │  40,41,  │      └──────┬───────┘                             │
│  │  50,70,90│             │                                     │
│  └────┬─────┘             ▼                                     │
│       │            ┌──────────────┐                             │
│       │            │ ERRO: CST não│                             │
│       │            │ permite      │                             │
│       │            │ desoneração  │                             │
│       │            └──────────────┘                             │
│       │                                                         │
│       ▼                                                         │
│  ┌──────────────────────────────────┐                           │
│  │ 2. Determinar motivo_desoneracao │                           │
│  │    (1-12, 90 conforme SEFAZ)     │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│        ┌────────┴────────┐                                      │
│        ▼                 ▼                                      │
│  ┌──────────┐      ┌──────────────┐                             │
│  │ Redução  │      │ Limitação de │                             │
│  │ de Base  │      │ Alíq. Efetiva│                             │
│  │ (%)      │      │ (alíquota    │                             │
│  │          │      │  alvo)       │                             │
│  └────┬─────┘      └──────┬───────┘                             │
│       │                   │                                     │
│       ▼                   ▼                                     │
│  ┌──────────────┐   ┌──────────────────┐                        │
│  │ 3a. Redução  │   │ 3b. Índice =     │                        │
│  │ Base Reduzida│   │ 1-(AliqAlvo/     │                        │
│  │ = Valor ×    │   │ AliqNominal)     │                        │
│  │ (1-PctRed%)  │   │ Base Reduzida =  │                        │
│  │ Ref: BR-TAX- │   │ Valor × (AliqAlvo│                        │
│  │ CALC-021     │   │ / AliqNominal)   │                        │
│  │              │   │ Ref: BR-TAX-     │                        │
│  │              │   │ CALC-022         │                        │
│  └──────┬───────┘   └────────┬─────────┘                        │
│         │                    │                                  │
│         └─────────┬──────────┘                                  │
│                   │                                             │
│                   ▼                                             │
│  ┌──────────────────────────────────┐                           │
│  │ 4. Calcular ICMS sobre base      │                           │
│  │    reduzida                      │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 5. Calcular vICMSDeson:          │                           │
│  │    = (Valor × AliqNominal)       │                           │
│  │      − ICMS_Efetivo              │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 6. Abater vICMSDeson do valor    │                           │
│  │    total da nota (BR-TAX-ACT-007)│                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 7. Registrar em trilha de        │                           │
│  │    auditoria (motDesICMS,        │                           │
│  │    vICMSDeson)                   │                           │
│  └──────────────────────────────────┘                           │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Detailed Steps

#### Step 1: Validar CST
- **Who:** `ICMSValidator.ValidateCSTDesoneracao(cst)`
- **When:** Antes de iniciar o cálculo de desoneração
- **How:** Se `cst ∈ {20, 30, 40, 41, 50, 70, 90}` → permitido. Caso contrário → erro. Ref: BR-TAX-CONS-013.
- **Output:** Validação

#### Step 2: Determinar Motivo
- **Who:** `ICMSDesoneracaoResolver.GetMotivo(item)`
- **When:** Após validação do CST
- **How:** Ler `item.motivo_desoneracao_icms` (1-12, 90). Se não informado, default = 9 (Outros).
- **Output:** `motDesICMS` (int)

#### Step 3a: Modo Redução de Base
- **Who:** `ICMSCalculator.CalcDesoneradoReducao(valor, aliquota, pctReducao)`
- **When:** Se `item.possui_desoneracao == true` e `item.percentual_reducao_base > 0`
- **How:** `baseReduzida = valor * (1 - pctReducao/100)`; `icms = baseReduzida * aliquota/100`
- **Output:** `icms_efetivo`, `vICMSDeson`

#### Step 3b: Modo Limitação de Alíquota
- **Who:** `ICMSCalculator.CalcDesoneradoLimite(valor, aliqNominal, aliqAlvo)`
- **When:** Se `item.aliquota_alvo > 0`
- **How:** `indice = 1 - (aliqAlvo/aliqNominal)`; `baseReduzida = valor * (aliqAlvo/aliqNominal)`; `icms = baseReduzida * aliqNominal/100`
- **Output:** `icms_efetivo`, `vICMSDeson`

#### Step 6: Abater do Total
- **Who:** `TaxResponseBuilder.ApplyDesoneracao(item, vICMSDeson)`
- **When:** Após calcular ICMS desonerado
- **How:** `valor_final_item = valor_item - vICMSDeson` (BR-TAX-ACT-007)
- **Output:** Valor final do item ajustado

### Exceptions and Edge Cases

| Cenário | Ação |
|:---|:---|
| CST 00 com flag desoneração | Log ERROR — CST 00 não permite desoneração; ignorar flag |
| motDesICMS não informado | Usar 9 (Outros) como default; log WARN |
| Redução = 100% (isenção total) | ICMS = 0; vICMSDeson = valor total do ICMS teórico |
| Simples Nacional (CRT 1) | Não aplicar desoneração "clássica" na NF-e; tratar na apuração PGDAS-D |
| vICMSDeson > ICMS teórico | Log CRITICAL — erro de configuração; limitar vICMSDeson = ICMS teórico |

---

# PROCEDIMENTOS DE CÁLCULO POR PERÍODO

---

## SOP-007: Cálculo de IRPJ/CSLL (Lucro Real e Presumido)

**Regras vinculadas:** BR-TAX-CALC-001 a 005, BR-TAX-CONS-001  
**Engine:** Period Engine — Lucratividade (`irpj_calculator.go`, `csll_calculator.go`)

### Process Flow

```
┌─────────────────────────────────────────────────────────────────┐
│ SOP-007: Cálculo de IRPJ/CSLL                                   │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  INÍCIO (fechamento trimestral/anual)                           │
│    │                                                            │
│    ▼                                                            │
│  ┌──────────────────────────────────┐                           │
│  │ 1. Receber TaxPeriodInput com    │                           │
│  │    LucroData e regime            │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│        ┌────────┴────────┐                                      │
│        ▼                 ▼                                      │
│  ┌──────────┐      ┌──────────────┐                             │
│  │ Lucro    │      │ Lucro        │                             │
│  │ Real     │      │ Presumido    │                             │
│  └────┬─────┘      └──────┬───────┘                             │
│       │                   │                                     │
│       ▼                   ▼                                     │
│  ┌──────────────────┐  ┌──────────────────┐                     │
│  │ 2a. Base =       │  │ 2b. Base =       │                     │
│  │ LucroContabil    │  │ ReceitaBruta ×   │                     │
│  │ + Adicoes LALUR  │  │ 8% (com/ind)     │                     │
│  │ - Exclusoes LALUR│  │ OU 32% (serv)    │                     │
│  │ - Comp Prejuizos │  │                  │                     │
│  │   (max 30%)      │  │                  │                     │
│  │ Ref: BR-TAX-     │  │ Ref: BR-TAX-     │                     │
│  │ CALC-001,        │  │ CALC-002         │                     │
│  │ CONS-001         │  │                  │                     │
│  └────────┬─────────┘  └────────┬─────────┘                     │
│           │                     │                               │
│           └──────────┬──────────┘                               │
│                      │                                          │
│                      ▼                                          │
│  ┌──────────────────────────────────┐                           │
│  │ 3. Calcular IRPJ:                │                           │
│  │    IRPJ = Base × 15%             │                           │
│  │    Ref: BR-TAX-CALC-003          │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 4. Calcular Adicional IRPJ:      │                           │
│  │    Se Base > 20.000 × n_meses    │                           │
│  │    Adicional = (Base - Limite)   │                           │
│  │               × 10%              │                           │
│  │    Ref: BR-TAX-CALC-004          │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 5. Calcular CSLL:                │                           │
│  │    CSLL = Base × 9%              │                           │
│  │    Ref: BR-TAX-CALC-005          │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 6. Abater retenções na fonte     │                           │
│  │    IRRF e CSLL retidos           │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 7. Emitir TaxPeriodResult com    │                           │
│  │    Grupo "LUCRO": IRPJ + CSLL    │                           │
│  └──────────────────────────────────┘                           │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Detailed Steps

#### Step 1: Receber Input do Período
- **Who:** `PeriodEngine.Calculate(input)`
- **When:** Fechamento trimestral (IRPJ/CSLL)
- **How:** Controller envia `TaxPeriodInput` com `LucroData` (lucro_contabil, adições, exclusões, prejuízo_acumulado, retenções) e `regime` (REAL/PRESUMIDO)
- **Output:** Dados do período estruturados

#### Step 2a: Base Lucro Real
- **Who:** `IRPJCalculator.CalcBaseLucroReal(data)`
- **When:** Se regime = REAL
- **How:**
  ```go
  lucroAjustado := data.LucroContabil + SUM(adicoes) - SUM(exclusoes)
  compMax := lucroAjustado * 0.30 // BR-TAX-CONS-001
  compEfetiva := math.Min(data.PrejuizoAcumulado, compMax)
  base := lucroAjustado - compEfetiva
  ```
- **Output:** `base_irpj_csll`

#### Step 2b: Base Lucro Presumido
- **Who:** `IRPJCalculator.CalcBasePresumido(receitaBruta, cnae)`
- **When:** Se regime = PRESUMIDO
- **How:** `base = receitaBruta * percentualPresuncao` (8% comércio/indústria, 32% serviços)
- **Output:** `base_irpj_csll`

#### Steps 3-5: Calcular IRPJ, Adicional e CSLL
- **Who:** `IRPJCalculator` e `CSLLCalculator`
- **When:** Após definir a base
- **How:** Fórmulas conforme BR-TAX-CALC-003, 004 e 005
- **Output:** Valores de IRPJ e CSLL

#### Step 6: Abater Retenções
- **Who:** `WithholdingDeductor.Apply(irpj, csll, retencoes)`
- **When:** Após calcular IRPJ e CSLL
- **How:** `irpj_a_pagar = irpj - retencoes.irrf`; `csll_a_pagar = csll - retencoes.csll`
- **Output:** Valores líquidos a recolher

### Exceptions

| Cenário | Ação |
|:---|:---|
| Prejuízo fiscal > 30% do lucro ajustado | Limitar compensação a 30%; saldo remanescente para próximos períodos |
| Adições LALUR sem código identificador | Log WARN + solicitar código ao Controller antes de processar |
| Retenções maiores que o imposto devido | IRPJ/CSLL = 0 + saldo de retenção a restituir (PER/DCOMP — SOP-015) |

---

## SOP-012: Cálculo de CPP e FGTS (Encargos de Folha)

**Regras vinculadas:** POLICE-FIN-00001 §4.1 (CPP, FGTS)  
**Engine:** Period Engine — Folha (`cpp_calculator.go`, `fgts_calculator.go`)

### Detailed Steps

#### Step 1: Receber Dados da Folha
- **Who:** `PeriodEngine.CalculateFolha(input)`
- **When:** Fechamento mensal
- **How:** Controller/RH envia `FolhaData` com `total_folha_bruta`, `base_fgts`, `base_inss_patronal`, `fap`, `rat`
- **Output:** Dados de folha

#### Step 2: Calcular CPP
- **Who:** `CPPCalculator.Calculate(folha)`
- **When:** Mensalmente
- **How:**
  - `cpp_base = folha.BaseINSSPatronal`
  - `cpp = cpp_base * 0.20` (20% — alíquota patronal básica)
  - `rat_ajustado = cpp_base * (folha.RAT/100) * folha.FAP`
  - `terceiros = cpp_base * 0.058` (5,8% — Sistema S)
  - `cpp_total = cpp + rat_ajustado + terceiros`
- **Output:** `cpp_total`

#### Step 3: Calcular FGTS
- **Who:** `FGTSCalculator.Calculate(folha)`
- **When:** Mensalmente
- **How:** `fgts = folha.BaseFGTS * 0.08` (8%)
- **Output:** `fgts_valor`

### Exceptions

| Cenário | Ação |
|:---|:---|
| FAP > 2.0 (anômalo) | Validar com RH — FAP varia entre 0.5 e 2.0 |
| RAT diferente do CNAE TELECOM | Log WARN + verificar enquadramento no CNAE correto |

---

# PROCEDIMENTOS DE ORQUESTRAÇÃO

---

## SOP-013: Execução do Pipeline de Cálculo por Fase

**Regras vinculadas:** BR-TAX-INF-001 a 004, BR-TAX-ACT-005, BR-TAX-ACT-006, DT-001, DT-005  
**Engine:** Tax Pipeline (`tax_pipeline.go`)

### Process Flow

```
┌─────────────────────────────────────────────────────────────────┐
│ SOP-013: Pipeline de Cálculo por Fase                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  INÍCIO (cada requisição POST /calculate)                       │
│    │                                                            │
│    ▼                                                            │
│  ┌──────────────────────────────────┐                           │
│  │ 1. Receber TaxDocumentInput      │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 2. Validar Input (CST, NCM,      │                           │
│  │    CFOP, regime, alíquotas)      │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 3. Resolver Fase Tributária      │                           │
│  │    Ref: BR-TAX-INF-001 a 004     │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 4. Selecionar Tributos (DT-001)  │                           │
│  │    Baseado na fase + natureza    │                           │
│  │    da operação                   │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 5. Executar Pipeline na Ordem:   │                           │
│  │    ┌─────────────────────────┐   │                           │
│  │    │ IS (SOP-003)            │   │ ← Pré-filtro: NCM restrito │
│  │    │   ↓                     │   │                           │
│  │    │ IPI (SOP-011)           │   │ ← "Por fora"; compõe base  │
│  │    │   ↓                     │   │   do ICMS (consum. final) │
│  │    │ CBS (SOP-001)           │   │ ← "Por fora"; não compõe  │
│  │    │   ↓                     │   │   base de outros          │
│  │    │ ICMS (SOP-009)          │   │ ← "Por dentro"; valor      │
│  │    │   ↓                     │   │   destacado excluído da   │
│  │    │ IBS (SOP-002 → CALC-018)│   │   base PIS/COFINS (STF)   │
│  │    │   ↓                     │   │                           │
│  │    │ ISS (SOP-010)           │   │ ← Serviços; não interage  │
│  │    │   ↓                     │   │   com ICMS                │
│  │    │ PIS/COFINS (SOP-008)    │   │ ← Base exclui ICMS        │
│  │    │   ↓                     │   │                           │
│  │    │ FUST (SOP-016)          │   │ ← Em cascata: base =      │
│  │    │   ↓                     │   │   bruto − ICMS − PIS −    │
│  │    │ FUNTTEL (SOP-016)       │   │   COFINS                  │
│  │    └─────────────────────────┘   │                           │
│  │    NOTA: Apenas tributos          │                           │
│  │    selecionados na DT-001 são     │                           │
│  │    executados                     │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 6. Roteamento Shadow vs Produção │                           │
│  │    Ref: BR-TAX-ACT-005           │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 7. Emitir TaxResponse            │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│               FIM                                               │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Detailed Steps

#### Step 1: Receber Input
- **Who:** Tax Pipeline entrypoint
- **When:** Cada requisição HTTP POST
- **How:** Deserializar `TaxDocumentInput` do JSON
- **Output:** Struct validada

#### Step 2: Validar Input
- **Who:** `InputValidator.Validate(input)`
- **When:** Antes de qualquer cálculo
- **How:**
  - Validar `regime_tributario` ∈ {SIMPLES, PRESUMIDO, REAL}
  - Validar `cst_icms`, `cst_pis`, `cst_cofins`, `cst_ipi`
  - Validar `cfop` (4 dígitos)
  - Validar `ncm` (8 dígitos)
  - Validar alíquotas nos ranges esperados
- **Output:** Input validado ou erros

#### Step 3: Resolver Fase
- **Who:** `PhaseResolver.Resolve(dataOperacao)`
- **When:** Após validação
- **How:** Mapear ano para fase conforme BR-TAX-INF-001 a 004
- **Output:** `Phase` enum

#### Step 4: Selecionar Tributos (DT-001)
- **Who:** `TaxSelector.Select(phase, natureza)`
- **When:** Após resolver fase
- **How:** Aplicar DT-001 — selecionar quais calculadoras executar baseado na fase e natureza da operação
- **Output:** Lista de calculadoras ativas

#### Step 5: Executar Pipeline
- **Who:** `TaxPipeline.Execute(calculadoras, input)`
- **When:** Após seleção
- **How:** Executar calculadoras na ordem fixa (IS → IPI → CBS → ICMS → IBS → ISS → PIS/COFINS → FUST → FUNTTEL). Calculadoras não selecionadas pela DT-001 são skip. A ordem é mandatória porque: IPI compõe base do ICMS para consumidor final; ICMS destacado é excluído da base de PIS/COFINS (Tese do Século — STF); FUST/FUNTTEL usam base líquida após ICMS+PIS+COFINS.
- **Output:** Resultados parciais de cada calculadora

#### Step 6: Roteamento Shadow/Produção
- **Who:** `TaxResponseBuilder.Build(results, phase)`
- **When:** Após todos os cálculos
- **How:**
  - Se Shadow Run: CBS/IBS → `shadow_tax_result`; Legacy → `total_a_pagar`
  - Se Produção Reforma: CBS/IBS → `total_a_pagar`; Legacy → zero
  - Se Hybrid: Ambos → `total_a_pagar` com fator de redução no Legacy
- **Output:** `TaxResponse` estruturada

### Exceptions

| Cenário | Ação |
|:---|:---|
| Calculadora retorna erro | Log ERROR + continuar com as demais + marcar item como `calculation_error` |
| Data da operação > 2033 mas ICMS ainda configurado | Log CRITICAL — possível erro de configuração |
| Fase não reconhecida para a data | Log ERROR + assumir fase mais recente configurada + notificar |

---

# PROCEDIMENTOS DE VALIDAÇÃO E MANUTENÇÃO

---

## SOP-004: Reconciliação Mensal Shadow Run

**Regras vinculadas:** BR-TAX-ACT-002  
**Owner:** Controller + Engineering Lead  
**Gatilho:** Dia 5 de cada mês  
**Engine:** N/A (processo manual/semi-automatizado)

### Process Flow

```
┌─────────────────────────────────────────────────────────────────┐
│ SOP-004: Reconciliação Mensal Shadow Run                        │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  DIA 5 DO MÊS (referente ao mês anterior)                       │
│    │                                                            │
│    ▼                                                            │
│  ┌──────────────────────────────────┐                           │
│  │ 1. Extrair totais mensais do     │                           │
│  │    Legacy: PIS, COFINS, ICMS, ISS│                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 2. Extrair totais mensais do     │                           │
│  │    Shadow: CBS, IBS              │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 3. Calcular Grupos:              │                           │
│  │    Leg_Federal = PIS + COFINS    │                           │
│  │    Shadow_Fed  = CBS             │                           │
│  │    Leg_Subnac  = ICMS + ISS      │                           │
│  │    Shadow_Sub  = IBS             │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 4. Calcular Variação:            │                           │
│  │    Var% = (Shadow-Leg)/Leg × 100 │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│        ┌────────┴────────┐                                      │
│        ▼                 ▼                                      │
│  ┌──────────┐      ┌──────────────┐                             │
│  │ Var% ≤   │      │ Var% > 10%   │                             │
│  │ 10%      │      │              │                             │
│  └────┬─────┘      └──────┬───────┘                             │
│       │                   │                                     │
│       ▼                   ▼                                     │
│  ┌──────────────┐   ┌──────────────────────┐                    │
│  │ Relatório OK │   │ 5. Investigar item   │                    │
│  │ Assinar +    │   │    a item            │                    │
│  │ Arquivar     │   │    Ref: BR-TAX-ACT-002                   │
│  └──────┬───────┘   └──────────┬───────────┘                    │
│         │                      │                                │
│         │                      ▼                                │
│         │           ┌──────────────────────┐                    │
│         │           │ Gerar relatório de   │                    │
│         │           │ divergência          │                    │
│         │           └──────────┬───────────┘                    │
│         │                      │                                │
│         │                      ▼                                │
│         │           ┌──────────────────────┐                    │
│         │           │ Agendar reunião      │                    │
│         │           │ Controller + Eng     │                    │
│         │           │ Lead (48h)           │                    │
│         │           └──────────┬───────────┘                    │
│         │                      │                                │
│         └──────────────────────┘                                │
│                        │                                        │
│                        ▼                                        │
│  ┌──────────────────────────────────┐                           │
│  │ 6. Controller assina relatório   │                           │
│  │    Tax Compliance Officer revisa │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 7. Arquivar em                    │                           │
│  │    compliance/shadow_run/YYYY/    │                           │
│  │    shadow_run_reconcile_MM.csv    │                           │
│  └──────────────────────────────────┘                           │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Detailed Steps

#### Step 1-2: Extrair Totais
- **Who:** Controller
- **When:** Dia 5 do mês
- **How:**
  ```sql
  -- Legacy
  SELECT SUM(pis_valor), SUM(cofins_valor), SUM(icms_valor), SUM(iss_valor)
  FROM billing_tax_rates.tax_results WHERE periodo = 'YYYY-MM' AND phase = 'PRODUCTION'

  -- Shadow
  SELECT SUM(cbs_valor), SUM(ibs_valor)
  FROM billing_tax_rates.shadow_tax_results WHERE periodo = 'YYYY-MM'
  ```
- **Output:** Totais mensais

#### Step 3-4: Calcular Variação
- **Who:** Controller
- **When:** Após extrair totais
- **How:** Planilha Excel ou script Python
  ```
  var_federal = (cbs_total - (pis_total + cofins_total)) / (pis_total + cofins_total) * 100
  var_subnacional = (ibs_total - (icms_total + iss_total)) / (icms_total + iss_total) * 100
  ```
- **Output:** Percentuais de variação

#### Step 5: Investigar (se necessário)
- **Who:** Controller + Engineering Lead
- **When:** Se variação > 10% em qualquer grupo
- **How:** Decompor divergência por UF, por NCM, por CST. Identificar outliers.
- **Output:** Relatório de divergência com causas identificadas

#### Step 6-7: Aprovar e Arquivar
- **Who:** Controller (assinatura) + Tax Compliance Officer (revisão)
- **When:** Após investigação concluída
- **How:** Assinar digitalmente o CSV; arquivar na pasta de compliance
- **Output:** Relatório assinado e arquivado

---

## SOP-005: Atualização de Alíquotas por Nova Legislação

**Regras vinculadas:** BR-TAX-ACT-004  
**Owner:** Tax Compliance Officer + Engineering Lead  
**Gatilho:** Publicação de nova lei, decreto ou ato normativo que altere alíquotas  
**SLA:** Deploy em produção em até 5 dias úteis

### Workflow (Timeline)

```
DIA 0 (Publicação)
├── Tax Compliance Officer identifica alteração e registra em tax_rate_change_log
├── Engineering Lead cria branch: tax-update/YYYY-MM-DD-<descricao>
│
DIA 0–1 (Implementação)
├── Atualizar tabelas SQL: billing_tax_rates.<tabela>
│   ├── icms_rules (alíquota interestadual, interna, FCP)
│   ├── cbs_rates (alíquota CBS por c_class_trib)
│   ├── iss_rates (alíquota municipal)
│   └── tipi (alíquota IPI por NCM)
├── Atualizar constantes em models/constants.go (se aplicável)
├── Atualizar testes unitários com novas alíquotas
│
DIA 1–2 (Code Review)
├── Abrir Pull Request
├── Dupla aprovação: Engineering Lead + Tax Compliance Officer
├── Merge na branch principal
│
DIA 2–3 (Staging)
├── Deploy em ambiente de staging
├── Executar suite de testes de regressão (hybrid_mode_test.go)
├── Executar smoke test com massa de dados de Shadow Run
├── Tax Compliance Officer valida resultados em staging
│
DIA 3–5 (Produção)
├── Deploy em produção (com rollback plan documentado)
├── Monitorar métricas por 24h (latency, error rate, precisão)
├── Comunicar Financeiro: nova alíquota, vigência, impacto estimado
├── Atualizar POLICE-FIN-00001 e RULES-CATALOG-FIN-00001 (se necessário)
│
DIA 5 (Fechamento)
└── Fechar tax_rate_change_log com status "DEPLOYED"
```

### Rollback Plan

Se a nova alíquota causar erro em produção:
1. **Rollback automático:** Se error rate > 1% em 5 minutos, reverter deploy automaticamente
2. **Rollback manual:** Engineering Lead pode reverter via `git revert` + deploy de emergência
3. **Comunicação:** Notificar Tax Compliance Officer e Controller em até 15 minutos

---

## SOP-006: Tratamento de Exceções e Creditamento

**Regras vinculadas:** BR-TAX-INF-006, BR-TAX-CONS-011  
**Owner:** Engine + Controller  
**Gatilho:** Apuração mensal/trimestral

### Process Flow

```
┌─────────────────────────────────────────────────────────────────┐
│ SOP-006: Tratamento de Exceções e Creditamento                  │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  INÍCIO (fechamento do período)                                  │
│    │                                                            │
│    ▼                                                            │
│  ┌──────────────────────────────────┐                           │
│  │ 1. Consolidar créditos do        │                           │
│  │    período por tributo           │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 2. Consolidar débitos do         │                           │
│  │    período por tributo           │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 3. Para cada tributo:            │                           │
│  │    Saldo = Créditos - Débitos    │                           │
│  │    Ref: BR-TAX-INF-006           │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│        ┌────────┼────────┐                                      │
│        ▼        ▼        ▼                                      │
│  ┌────────┐ ┌────────┐ ┌────────┐                               │
│  │ Saldo  │ │ Saldo  │ │ Saldo  │                               │
│  │ > 0    │ │ < 0    │ │ = 0    │                               │
│  └───┬────┘ └───┬────┘ └───┬────┘                               │
│      │          │          │                                    │
│      ▼          ▼          ▼                                    │
│  ┌────────┐ ┌────────┐ ┌────────┐                               │
│  │Saldo a │ │Valor a │ │Neutro  │                               │
│  │Recuper.│ │Recolher│ │        │                               │
│  └───┬────┘ └───┬────┘ └────────┘                               │
│      │          │                                               │
│      │          ▼                                               │
│      │   ┌──────────────────┐                                   │
│      │   │ 4. Emitir Guia   │                                   │
│      │   │ (DARF/GNRE/GPS)  │                                   │
│      │   └──────────────────┘                                   │
│      │                                                          │
│      ▼                                                          │
│  ┌──────────────────────────────────┐                           │
│  │ 5. Saldo credor: acumular para   │                           │
│  │    próximo período               │                           │
│  │    OU                             │                           │
│  │    Compensação cruzada (PER/      │                           │
│  │    DCOMP) — SOP-015              │                           │
│  └──────────────────────────────────┘                           │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Detailed Steps

#### Step 1-2: Consolidar
- **Who:** `PeriodEngine.ClosePeriod()`
- **When:** Fechamento do período
- **How:** Somar todos os créditos (entradas/compras) e débitos (saídas/vendas) do período por tributo
- **Output:** Extratos de créditos e débitos

#### Step 3: Calcular Saldo
- **Who:** `SaldoCalculator.Calculate(creditos, debitos)`
- **When:** Após consolidação
- **How:** `saldo = creditos - debitos` → classificar conforme BR-TAX-INF-006
- **Output:** `SaldoStatus` e valores

#### Step 4: Emitir Guias
- **Who:** Controller (manual) ou sistema de ERP
- **When:** Se `Saldo < 0`
- **How:** Gerar DARF (tributos federais), GNRE (ICMS interestadual), GPS (CPP/FGTS)
- **Output:** Guias de recolhimento

#### Step 5: Tratar Saldo Credor
- **Who:** `PeriodEngine` + Controller
- **When:** Se `Saldo > 0`
- **How:** Acumular em `saldos_a_recuperar[]` para o próximo período. Se crédito excedente e sem previsão de débito futuro → avaliar PER/DCOMP (SOP-015)
- **Output:** Saldo credor registrado

### Exceptions

| Cenário | Ação |
|:---|:---|
| Crédito acumulado por > 5 anos | Alerta — risco de prescrição (CTN art. 168) |
| Saldo credor > R$ 1.000.000 | Revisão obrigatória pelo Tax Compliance Officer |
| Divergência entre crédito contábil e fiscal | Conciliação Controller + Consultoria Tributária |

---

# PROCEDIMENTOS DE CONTINGÊNCIA

---

## SOP-014: Tratamento de Contingência — Circuit Breaker IBS

**Regras vinculadas:** BR-TAX-ACT-001  
**Owner:** `ms-billing-engine-tax-rates`  
**Gatilho:** 3 falhas na API IBS em janela de 60 segundos

### Estados do Circuit Breaker

```
     ┌──────────┐
     │  CLOSED  │ ← Estado normal
     └────┬─────┘
          │
          │ ≥ 3 falhas em 60s
          ▼
     ┌──────────┐
     │   OPEN   │ ← Todas as chamadas falham imediatamente (fast-fail)
     └────┬─────┘   Usa cache de fallback
          │
          │ Após 5 minutos
          ▼
     ┌──────────────┐
     │  HALF_OPEN   │ ← Permite 1 chamada de teste
     └──────┬───────┘
            │
     ┌──────┴──────┐
     ▼             ▼
  Sucesso       Falha
     │             │
     ▼             ▼
  CLOSED        OPEN
```

### Ações por Transição

| Transição | Ação |
|:---|:---|
| CLOSED → OPEN | Log CRITICAL + notificar Tax Compliance Officer + usar última taxa cacheada (TTL estendido) |
| OPEN → HALF_OPEN | Log INFO + permitir 1 chamada de teste |
| HALF_OPEN → CLOSED | Log INFO + notificar que API IBS normalizou |
| HALF_OPEN → OPEN | Log ERROR + continuar em fallback |
| OPEN (mantido > 1h) | Escalar para Nível 2 (Tax Compliance Officer) |

### Fallback Data

Durante o estado OPEN, usar:
- Última alíquota cacheada antes da falha (mesmo que TTL expirado)
- Flag `rate_source = "FALLBACK_CIRCUIT_OPEN"`
- Alerta visível no dashboard de monitoração

---

## SOP-015: Compensação Cruzada de Tributos (PER/DCOMP)

**Regras vinculadas:** BR-TAX-CONS-011  
**Owner:** Controller + Tax Compliance Officer  
**Gatilho:** Quando há crédito excedente de tributo federal e débito em outro tributo federal

### Condições de Elegibilidade

| De (Crédito Excedente) | Para (Débito) | Permitido? |
|:---|:---|:---|
| PIS | COFINS | ✅ Sim |
| COFINS | PIS | ✅ Sim |
| PIS/COFINS | IRPJ | ✅ Sim (com restrições) |
| PIS/COFINS | CSLL | ✅ Sim (com restrições) |
| PIS/COFINS | CPP (INSS Patronal) | ✅ Sim |
| IPI | IRPJ/CSLL | ❌ Não (naturezas distintas) |
| ICMS | Qualquer Federal | ❌ Não (esferas diferentes) |
| IBS | CBS | ⚠️ A definir pelo Comitê Gestor |

### Workflow

```
┌─────────────────────────────────────────────────────────────────┐
│ SOP-015: Compensação Cruzada (PER/DCOMP)                        │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  INÍCIO (Controller identifica crédito excedente)               │
│    │                                                            │
│    ▼                                                            │
│  ┌──────────────────────────────────┐                           │
│  │ 1. Controller elabora planilha   │                           │
│  │    de elegibilidade              │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 2. Tax Compliance Officer        │                           │
│  │    avalia e aprova/rejeita       │                           │
│  │    Ref: BR-TAX-CONS-011          │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│        ┌────────┴────────┐                                      │
│        ▼                 ▼                                      │
│  ┌──────────┐      ┌──────────┐                                 │
│  │ Aprovado │      │ Rejeitado│                                 │
│  └────┬─────┘      └────┬─────┘                                 │
│       │                 │                                       │
│       │                 ▼                                       │
│       │           ┌──────────────┐                              │
│       │           │ Documentar   │                              │
│       │           │ motivo +     │                              │
│       │           │ arquivar     │                              │
│       │           └──────────────┘                              │
│       │                                                         │
│       ▼                                                         │
│  ┌──────────────────────────────────┐                           │
│  │ 3. Controller transmite PER/     │                           │
│  │    DCOMP via sistema da RFB      │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 4. Acompanhar deferimento no     │                           │
│  │    e-CAC da Receita Federal      │                           │
│  └──────────────┬───────────────────┘                           │
│                 │                                               │
│                 ▼                                               │
│  ┌──────────────────────────────────┐                           │
│  │ 5. Atualizar saldos no           │                           │
│  │    Period Engine                 │                           │
│  └──────────────────────────────────┘                           │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## 4. Métricas de Performance dos Procedimentos

| Métrica | Alvo | Procedimento | Método de Medição |
|:---|:---|:---|:---|
| Tempo de cálculo CBS por item | p99 < 50ms | SOP-001 | APM tracing |
| Taxa de cache hit IBS | ≥ 95% | SOP-002 | Redis metrics |
| Precisão da tabela NCM Seletivo | 100% conforme lista oficial | SOP-003 | Hash SHA-256 semanal |
| Prazo de reconciliação Shadow Run | Até dia 5 do mês | SOP-004 | Calendário de compliance |
| SLA de atualização de alíquota | ≤ 5 dias úteis | SOP-005 | `tax_rate_change_log` |
| Precisão IRPJ/CSLL vs. contabilidade | ≥ 99,9% | SOP-007 | Conciliação trimestral |
| Cobertura de CST no cálculo PIS/COFINS | 100% dos CSTs previstos | SOP-008 | Suite de testes |
| Precisão DIFAL por UF | 100% conforme tab. IBPT/CONFAZ | SOP-009 | Testes paramétricos por UF |
| Latência do pipeline completo | p99 < 200ms | SOP-013 | APM end-to-end |
| Tempo de recuperação circuit breaker | < 5 min (HALF_OPEN) | SOP-014 | Dashboard |
| Volume de PER/DCOMP processados | 100% dos elegíveis | SOP-015 | Log de compensações |
| Precisão FUST/FUNTTEL vs. guia ANATEL | ≥ 99,9% | SOP-016 | Conciliação mensal |
| Cobertura de CSTs de desoneração | 100% dos CSTs 20,30,40,41,50,70,90 | SOP-017 | Suite de testes paramétricos |

---

## 5. Referências

| Documento | Código |
|:---|:---|
| Política de Cálculo de Impostos | [POLICE-FIN-00001](./POLICE-FIN-00001-CALCULO-DE-IMPOSTOS-CORPORATIVOS.md) |
| Catálogo de Regras de Negócio | [RULES-CATALOG-FIN-00001](./RULES-CATALOG-FIN-00001.md) |
| Escopo do Projeto | [README-ESCOPO.md](./README-ESCOPO.md) |
| Brainstorm de Arquitetura | [README-BRAINSTORM.md](./README-BRAINSTORM.md) |

---

## Controle de Versão

| Versão | Data | Autor | Alterações |
|:---|:---|:---|:---|
| 1.0 | 2026-06-21 | Comitê Fiscal + Engineering Lead | Versão inicial. 15 SOPs cobrindo cálculo por operação, cálculo por período, orquestração, validação, contingência e manutenção. Vinculado à POLICE-FIN-00001 v1.0 e RULES-CATALOG-FIN-00001 v1.0. |
| 1.1 | 2026-06-21 | Comitê Fiscal + Engineering Lead | Adicionados SOP-016 (FUST/FUNTTEL) e SOP-017 (ICMS Desonerado). SOP-009 Step 2 atualizado com alíquota 4% para importados. SOP-013 pipeline order corrigido (IS → IPI → CBS → ICMS → IBS → ISS → PIS/COFINS → FUST → FUNTTEL). Métricas e RACI atualizados. Vinculado à POLICE-FIN-00001 v1.1 e RULES-CATALOG-FIN-00001 v1.1. |

---

**Próxima Revisão Obrigatória:** 2026-12-21, alinhada à política-mãe e ao catálogo de regras.

---

_Documento classificado como **Confidencial — Uso Interno**. Vinculado à trilogia POLICE-FIN-00001 / RULES-CATALOG-FIN-00001 / PROCEDURE-FIN-00001. Distribuição controlada pelo Tax Compliance Officer._
