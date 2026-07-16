# FEATURE-2026-06-21 — Novas Features Identificadas via Gap Analysis

**Origem:** Cross-reference entre POLICE-FIN-00001 v1.1, PROCEDURE-FIN-00001 v1.1, RULES-CATALOG-FIN-00001 v1.1 × código-fonte `ms-billing-engine-tax-rates`

**Data da análise:** 2026-06-21  
**Data da revisão:** 2026-06-30  
**Método:** Spec Miner + Gen-Specs-as-Issues  
**Código analisado (snapshot inicial):** 22 arquivos Go (cmd/, internal/), 7 tabelas SQL (data/init.sql)  
**Estado atual (pós-PR #6):** 40+ arquivos Go, 14 tabelas SQL, 10 GAPs (incluindo Fase 0: rate limiting, versioning, deploy)

---

## Resumo Executivo

A comparação entre os 3 documentos organizacionais (Política, Procedimentos e Regras de Negócio) e o código-fonte real do `ms-billing-engine-tax-rates` revelou **7 features críticas não implementadas** e **2 correções estruturais** no pipeline de cálculo. As ausências mais graves são: **ISS** (tributo municipal obrigatório para telecom), **FUST/FUNTTEL** (contribuições setoriais de telecom), e o sistema de **Resolução de Fase Tributária** que controla quais tributos são calculados em cada período (2026–2033).

---

## Features Identificadas

### F-001: Calculadora de ISS (Imposto sobre Serviços)

**Prioridade:** 🔴 Crítica  
**Estimativa de esforço:** Médio (3–5 dias)  
**Referências organizacionais:**
- POLICE-FIN-00001 §4.1 #9, §4.3 (ISS)
- PROCEDURE-FIN-00001 SOP-010
- RULES-CATALOG-FIN-00001 BR-TAX-CALC-016, BR-TAX-CONS-007, BR-TAX-DEF-007

**Situação atual:** ❌ Não implementado. O código não possui calculadora de ISS.

**O que implementar:**

| # | Requisito | Regra vinculada |
|---|-----------|----------------|
| 1 | Criar `ISSCalculator` em `internal/legacy/iss.go` implementando a interface `domain.TaxCalculator` | BR-TAX-CALC-016 |
| 2 | Identificar item como serviço via `ItemListaServico` (ex: "1.05" para telecom — LC 116/2003) | BR-TAX-DEF-007 |
| 3 | Validar alíquota municipal no intervalo `[2%, 5%]` | BR-TAX-CONS-007 |
| 4 | Calcular `ISS = Preço_Serviço × Alíquota_Municipal` | BR-TAX-CALC-016 |
| 5 | Tratar retenção de ISS na fonte (`RetencaoISS == true`) | SOP-010 Step 4 |
| 6 | Integrar ao motor bifásico (Fase 2 — paralela) | SOP-013 pipeline order |
| 7 | Implementar testes unitários: `TestISS_Valor_AliquotaMunicipal`, `TestISS_AliquotaForaDoIntervalo_RetornaErro`, `TestISS_Telecom_ItemLista105`, `TestISS_RetencaoFonte` | — |

**Estrutura esperada:**
```go
// internal/legacy/iss.go
type ISSCalculator struct {
    repo repository.TaxRepository
}

func (c *ISSCalculator) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
    // 1. Para cada item, verificar se é serviço (ItemListaServico preenchido)
    // 2. Se não for serviço → skip (ISS = 0)
    // 3. Validar aliquota ∈ [2%, 5%] — constraint BR-TAX-CONS-007
    // 4. ISS = Preço_Serviço × Alíquota_Municipal — BR-TAX-CALC-016
    // 5. Se RetencaoISS → flag iss_retido = true no MoreTextDetails
}
```

---

### F-002: Calculadora de FUST (Fundo de Universalização dos Serviços de Telecomunicações)

**Prioridade:** 🔴 Crítica  
**Estimativa de esforço:** Médio (2–3 dias)  
**Referências organizacionais:**
- POLICE-FIN-00001 §4.1 #10, §4.3 (FUST/FUNTTEL)
- PROCEDURE-FIN-00001 SOP-016
- RULES-CATALOG-FIN-00001 BR-TAX-CALC-019, BR-TAX-DEF-010, BR-TAX-CONS-012, BR-TAX-INF-007

**Situação atual:** ❌ Não implementado.

**O que implementar:**

| # | Requisito | Regra vinculada |
|---|-----------|----------------|
| 1 | Criar `FUSTCalculator` em `internal/legacy/fust.go` | BR-TAX-CALC-019 |
| 2 | Classificar natureza do serviço: SCM/STFC (incide) vs. SVA (não incide) | BR-TAX-INF-007 |
| 3 | Obter valores de ICMS, PIS, COFINS já calculados do contexto do item | SOP-016 Step 2 |
| 4 | Calcular base líquida: `Base = Valor_Serviço − ICMS − PIS − COFINS` | BR-TAX-CALC-019 |
| 5 | Se base líquida negativa → FUST = 0 + log WARN | SOP-016 edge case |
| 6 | Calcular: `FUST = Base × 0,01` (1%) — Lei 9.998/2000 | BR-TAX-CALC-019 |
| 7 | Integrar ao motor APÓS PIS/COFINS (depende dos valores deles) | SOP-013 pipeline order |
| 8 | Testes: `TestFUST_SCM_Incide`, `TestFUST_SVA_NaoIncide`, `TestFUST_BaseLiquida`, `TestFUST_BaseNegativa` | — |

**Constraint crítica:** FUST depende dos valores de ICMS, PIS e COFINS já calculados — portanto, deve ser executado DEPOIS destes tributos no pipeline (não pode ser paralelo).

---

### F-003: Calculadora de FUNTTEL (Fundo para o Desenvolvimento Tecnológico das Telecomunicações)

**Prioridade:** 🔴 Crítica  
**Estimativa de esforço:** Pequeno (1–2 dias)  
**Referências organizacionais:**
- POLICE-FIN-00001 §4.1 #11, §4.3 (FUST/FUNTTEL)
- PROCEDURE-FIN-00001 SOP-016
- RULES-CATALOG-FIN-00001 BR-TAX-CALC-020, BR-TAX-DEF-011

**Situação atual:** ❌ Não implementado.

**O que implementar:**

| # | Requisito | Regra vinculada |
|---|-----------|----------------|
| 1 | Criar `FUNTTELCalculator` — mesma base que FUST | BR-TAX-CALC-020 |
| 2 | Calcular: `FUNTTEL = Base_FUST × 0,005` (0,5%) — Lei 10.052/2000 | BR-TAX-CALC-020 |
| 3 | Integrar no pipeline imediatamente APÓS FUST | SOP-013 pipeline order |
| 4 | Testes: `TestFUNTTEL_SCM_Incide`, `TestFUNTTEL_MesmaBaseFUST` | — |

**Nota de design:** FUST e FUNTTEL compartilham a mesma base de cálculo e a mesma lógica de classificação SCM/STFC vs. SVA. Recomenda-se compartilhar a classificação de natureza do serviço entre ambos, possivelmente via um `TelecomClassifier` comum.

---

### F-004: ICMS Desonerado — Redução de Base e Limitação de Alíquota

**Prioridade:** 🟠 Alta  
**Estimativa de esforço:** Grande (4–6 dias)  
**Referências organizacionais:**
- POLICE-FIN-00001 §4.3 (ICMS Desonerado)
- PROCEDURE-FIN-00001 SOP-017
- RULES-CATALOG-FIN-00001 BR-TAX-CALC-021, BR-TAX-CALC-022, BR-TAX-CONS-013, BR-TAX-ACT-007

**Situação atual:** 🟡 Parcial. O banco de dados tem colunas (`motivo_desoneracao_icms`, `possui_desoneracao`, `percentual_reducao_base`) mas a lógica de cálculo de desoneração NÃO está implementada na calculadora ICMS.

**O que implementar:**

| # | Requisito | Regra vinculada |
|---|-----------|----------------|
| 1 | Validar CST permite desoneração: `CST ∈ {20, 30, 40, 41, 50, 70, 90}`. CST 00 NÃO permite. | BR-TAX-CONS-013 |
| 2 | Determinar `motDesICMS` (códigos 1–12, 90 conforme SEFAZ). Default = 9 (Outros). | SOP-017 Step 2 |
| 3 | **Modo Redução de Base:** `Base_Reduzida = Valor × (1 − PctRedução/100)`, `ICMS = Base_Reduzida × Alíquota/100`, `vICMSDeson = (Valor × Alíquota) − ICMS` | BR-TAX-CALC-021 |
| 4 | **Modo Limitação de Alíquota:** `Índice = 1 − (AliqAlvo/AliqNominal)`, `Base_Reduzida = Valor × (AliqAlvo/AliqNominal)`, `ICMS = Base_Reduzida × Alíquota/100` | BR-TAX-CALC-022 |
| 5 | Abater `vICMSDeson` do valor total da nota fiscal | BR-TAX-ACT-007 |
| 6 | Simples Nacional (CRT=1): NÃO aplicar desoneração "clássica" na NF-e | SOP-017 edge case |
| 7 | Testes: `TestICMS_Desonerado_ReducaoBase`, `TestICMS_Desonerado_LimitacaoAliquota`, `TestDesoneracao_CSTInvalido_RetornaErro`, `TestDesoneracao_AbateValorTotal`, `TestDesoneracao_SimplesNacional_NaoAplica` | — |

---

### F-005: Sistema de Resolução de Fase Tributária (Phase Resolver)

**Prioridade:** 🟠 Alta  
**Estimativa de esforço:** Grande (5–7 dias)  
**Referências organizacionais:**
- POLICE-FIN-00001 §5.1, §5.2, §5.3
- PROCEDURE-FIN-00001 SOP-013
- RULES-CATALOG-FIN-00001 BR-TAX-INF-001 a 004, DT-001, DT-005, BR-TAX-ACT-005, BR-TAX-ACT-006

**Situação atual:** ❌ Não implementado. O código atual calcula todos os tributos incondicionalmente, sem distinção entre Shadow Run (2026), CBS Plena (2027), Transição Subnacional (2029–2032) ou IVA Dual (2033+).

**O que implementar:**

| # | Requisito | Regra vinculada |
|---|-----------|----------------|
| 1 | Criar `PhaseResolver` que mapeia `DataOperacao` → `Phase` enum: `SHADOW_RUN` (2026), `CBS_PLENA` (2027), `TRANSICAO_SUBNACIONAL` (2029–2032), `IVA_DUAL` (2033+) | BR-TAX-INF-001 a 004 |
| 2 | Criar `TaxSelector` que seleciona calculadoras ativas por fase + natureza da operação (DT-001) | DT-001 |
| 3 | Implementar `shadow_tax_result` separado de `total_a_pagar`: na Fase Shadow Run, CBS e IBS não compõem o total a pagar | BR-TAX-ACT-005 |
| 4 | Na Fase IVA Dual (2033+), PIS, COFINS, ICMS e ISS são extintos (valor zero) | BR-TAX-ACT-006 |
| 5 | Na Fase Transição Subnacional, aplicar fator de redução no ICMS e ISS | POLICE-FIN-00001 §5.2 |
| 6 | Integrar ao `main.go` e ao motor `BillingEngineStruct.Process()` | — |
| 7 | Testes: `TestPhaseResolver_2026_ShadowRun`, `TestPhaseResolver_2027_CBSPlena`, `TestPhaseResolver_2030_Transicao`, `TestPhaseResolver_2033_IVADual`, `TestShadowRun_CBS_NaoCompoeTotal`, `TestIVADual_TributosLegacy_Extintos` | — |

**Modelo de dados esperado:**
```go
// internal/phase/phase.go
type Phase string
const (
    PhaseShadowRun          Phase = "SHADOW_RUN"
    PhaseCBSPlena           Phase = "CBS_PLENA"
    PhaseTransicaoSubnacional Phase = "TRANSICAO_SUBNACIONAL"
    PhaseIVADual            Phase = "IVA_DUAL"
)

type PhaseResolver struct{}

func (r *PhaseResolver) Resolve(dataOperacao time.Time) Phase {
    year := dataOperacao.Year()
    switch {
    case year == 2026: return PhaseShadowRun
    case year == 2027: return PhaseCBSPlena
    case year >= 2029 && year <= 2032: return PhaseTransicaoSubnacional
    case year >= 2033: return PhaseIVADual
    default: return PhaseShadowRun // safe default
    }
}

type TaxSelector struct{}

func (s *TaxSelector) Select(phase Phase, natureza string) []domain.TaxCalculator {
    // Aplica DT-001 — retorna apenas calculadoras ativas para esta fase
}
```

---

### F-006: IS como Pré-Filtro Independente (antes da CBS)

**Prioridade:** 🟡 Média  
**Estimativa de esforço:** Médio (3–4 dias)  
**Referências organizacionais:**
- POLICE-FIN-00001 §4.3 (IS)
- PROCEDURE-FIN-00001 SOP-003
- RULES-CATALOG-FIN-00001 BR-TAX-INF-005, BR-TAX-CONS-010

**Situação atual:** 🟡 Parcial. O IS é calculado dentro do `ReformaCalculator` (`reforma.go`) como parte do IVA Dual, mas as regras de negócio exigem que o IS seja verificado ANTES do cálculo da CBS (é um pré-filtro), usando a tabela separada `ncm_seletivo`. O IS não deve sofrer redução (ao contrário de CBS/IBS).

**O que implementar:**

| # | Requisito | Regra vinculada |
|---|-----------|----------------|
| 1 | Criar `ISFilter` como pré-calculadora (Fase 0 — antes de todos) | BR-TAX-CONS-010 |
| 2 | Verificar flag `isento_is` — se true, IS = 0 independente do NCM | SOP-003 Step 2 |
| 3 | Consultar tabela `ncm_seletivo` para verificar se NCM está sujeito ao IS | BR-TAX-INF-005 |
| 4 | Se NCM na tabela: `IS = Valor × Aliquota_IS_Categoria` | BR-TAX-INF-005 |
| 5 | Se NCM não na tabela: IS = 0 (não incide) | SOP-003 edge case |
| 6 | Registrar em auditoria mesmo se IS = 0 (com flag `is_exempt`) | SOP-003 Step 4 |
| 7 | Criar tabela `billing_tax_rates.ncm_seletivo` no schema SQL | SOP-003 Step 3 |
| 8 | Refatorar `ReformaCalculator` para remover lógica de IS (fica apenas CBS + IBS) | — |
| 9 | Testes: `TestIS_BebidaAlcoolica_Incide`, `TestIS_TelecomPuro_NaoIncide`, `TestIS_IsentoFlag_Override`, `TestIS_Antes_CBS_Pipeline` | — |

---

### F-007: Circuit Breaker para API do Comitê Gestor IBS

**Prioridade:** 🟡 Média  
**Estimativa de esforço:** Médio (3–4 dias)  
**Referências organizacionais:**
- POLICE-FIN-00001 §4.3 (IBS), §11.1
- PROCEDURE-FIN-00001 SOP-002, SOP-014
- RULES-CATALOG-FIN-00001 BR-TAX-ACT-001, BR-TAX-CONS-009

**Situação atual:** 🟡 Parcial. O código atual consulta alíquotas IBS via banco de dados (`GetIvaDualRule()`), enquanto a política exige consulta à API do Comitê Gestor IBS em tempo real, com cache Redis (TTL 24h) e circuit breaker.

**O que implementar:**

| # | Requisito | Regra vinculada |
|---|-----------|----------------|
| 1 | Criar `IBSClient` com chamada HTTP `GET /api/v1/rates?ibge_code={code}` | SOP-002 Step 3 |
| 2 | Implementar cache Redis com chave `ibs:rate:{ibge_code}:{date}` e TTL = 24h | BR-TAX-CONS-009 |
| 3 | Implementar circuit breaker: 3 falhas em 60s → OPEN → usar cache expirado + notificar | BR-TAX-ACT-001 |
| 4 | Estados: CLOSED → OPEN (≥3 falhas) → HALF_OPEN (após 5min) → CLOSED (sucesso) | SOP-014 |
| 5 | Flag `rate_source` nos detalhes: "CACHE", "API_COMITE_GESTOR", "FALLBACK_CIRCUIT_OPEN" | SOP-002 Step 8 |
| 6 | Alerta visível no dashboard quando circuit breaker aberto | SOP-014 |
| 7 | Testes: `TestIBS_Cache_TTL_24h`, `TestIBS_CircuitBreaker_AbreApos3Falhas`, `TestIBS_CircuitBreaker_UsaCacheFallback`, `TestIBS_AliquotaAnomala_Maior50pct` | — |

**Nota de design:** Enquanto a API do Comitê Gestor não é publicada (Gap G2), o fallback deve continuar usando `GetIvaDualRule()` (banco de dados). O `IBSClient` deve ser projetado com uma interface que permita trocar a fonte de dados (mock → API → DB fallback) sem alterar a calculadora.

---

## Correções Estruturais

### C-001: Reordenar Pipeline de Cálculo

**Prioridade:** 🔴 Crítica  
**Referência:** SOP-013, BR-TAX-CONS-010

**Situação atual:** O pipeline atual é:
```
Fase 1 (Sequencial): IPI
Fase 2 (Paralela):  ICMS + PIS/COFINS + Reforma (CBS/IBS/IS)
```

**Pipeline requerido pela política:**
```
Fase 0 (Sequencial): IS (pré-filtro — verifica NCM seletivo antes de tudo)
Fase 1 (Sequencial): IPI (compõe base do ICMS para consumidor final)
Fase 2 (Sequencial): CBS ("por fora", não compõe base de outros)
Fase 3 (Sequencial): ICMS (seu valor é excluído da base PIS/COFINS — Tese do Século)
Fase 4 (Paralela):  IBS + ISS + PIS/COFINS (independentes entre si)
Fase 5 (Sequencial): FUST (depende de ICMS + PIS + COFINS)
Fase 6 (Sequencial): FUNTTEL (depende da base do FUST)
```

**Impacto:** Esta reordenação é mandatória porque as dependências entre tributos são assimétricas:
- IPI → ICMS: IPI compõe a base do ICMS para consumidor final
- ICMS → PIS/COFINS: ICMS destacado é excluído da base de PIS/COFINS (STF, "Tese do Século")
- ICMS + PIS + COFINS → FUST/FUNTTEL: base líquida após impostos principais
- IS → CBS: IS é pré-filtro obrigatório (BR-TAX-CONS-010)

**O que modificar:**
1. `internal/calculator/engine.go`: Refatorar `BillingEngineStruct` para suportar múltiplas fases sequenciais
2. `cmd/api/main.go`: Atualizar wiring das calculadoras na nova ordem
3. Atualizar `legacy_adapter.go`: Garantir que dependências entre fases sejam injetadas corretamente

### C-002: Expandir Modelo de Dados (Schema SQL)

**Prioridade:** 🟠 Alta  
**Referência:** SOP-001, SOP-003, SOP-002

**Tabelas a criar em `data/init.sql`:**

| Tabela | Propósito | SOP |
|--------|-----------|-----|
| `cbs_rates` | Alíquotas CBS por classe tributária (`c_class_trib`) | SOP-001 |
| `ncm_seletivo` | NCMs sujeitos ao Imposto Seletivo com alíquotas por categoria | SOP-003 |
| `iss_rates` | Alíquotas ISS por município (código IBGE) | SOP-010 |

---

## Matriz de Rastreabilidade

| Feature | Regras de Negócio | SOPs | Prioridade | Esforço | Dependências | Status |
|---------|-------------------|------|-----------|---------|-------------| - |
| F-001 ISS | BR-TAX-CALC-016, BR-TAX-CONS-007, BR-TAX-DEF-007 | SOP-010 | 🔴 Crítica | Médio | — | ✅ |
| F-002 FUST | BR-TAX-CALC-019, BR-TAX-DEF-010, BR-TAX-CONS-012, BR-TAX-INF-007 | SOP-016 | 🔴 Crítica | Médio | C-001 (ordem pipeline) | ✅ |
| F-003 FUNTTEL | BR-TAX-CALC-020, BR-TAX-DEF-011 | SOP-016 | 🔴 Crítica | Pequeno | F-002 (mesma base FUST) | ✅ |
| F-004 ICMS Desonerado | BR-TAX-CALC-021/022, BR-TAX-CONS-013, BR-TAX-ACT-007 | SOP-017 | 🟠 Alta | Grande | — | ✅ |
| F-005 Phase Resolver | BR-TAX-INF-001 a 004, DT-001, DT-005, BR-TAX-ACT-005/006 | SOP-013 | 🟠 Alta | Grande | — | ✅ |
| F-006 IS Pré-Filtro | BR-TAX-INF-005, BR-TAX-CONS-010 | SOP-003 | 🟡 Média | Médio | C-002 (ncm_seletivo) | ✅ |
| F-007 IBS Circuit Breaker | BR-TAX-ACT-001, BR-TAX-CONS-009 | SOP-002, SOP-014 | 🟡 Média | Médio | Gap G2 (API externa) | ✅ |
| C-001 Pipeline Order | BR-TAX-CONS-010 | SOP-013 | 🔴 Crítica | Grande | Bloqueia F-002, F-003 | ✅ |
| C-002 Schema SQL | — | SOP-001, SOP-003, SOP-010 | 🟠 Alta | Pequeno | Bloqueia F-006 | ✅ |

---

## Ordem de Implementação Recomendada

```
Fase 1 (Fundação):
  ├── C-001: Reordenar pipeline de cálculo
  └── C-002: Criar tabelas SQL (cbs_rates, ncm_seletivo, iss_rates)

Fase 2 (Tributos Críticos):
  ├── F-001: ISS Calculator
  ├── F-002: FUST Calculator
  └── F-003: FUNTTEL Calculator

Fase 3 (Reforma e Governança):
  ├── F-005: Phase Resolution System
  ├── F-006: IS Pre-filter
  └── F-004: ICMS Desonerado

Fase 4 (Resiliência):
  └── F-007: IBS Circuit Breaker (dependente de API externa)
```

---

## Gaps Bloqueantes (Dependências Externas)

| ID | Gap | Features Impactadas | Previsão |
|----|-----|---------------------|----------|
| G1 | Alíquota CBS setorial TELECOM não definida | F-005 (CBS Plena 2027) | Ministério da Fazenda |
| G2 | API Comitê Gestor IBS não publicada | F-007 | LC 214/2025 pendente |
| G4 | Lista oficial NCMs sujeitos ao IS não publicada | F-006 | Ministério da Fazenda |

---

## Referências

| Documento | Código |
|-----------|--------|
| Política de Cálculo de Impostos | [POLICE-FIN-00001](../../../../../business-inputs/business-documents/business-policies/POLICE-FIN-00001-CALCULO-DE-IMPOSTOS-CORPORATIVOS.md) §4.1, §4.3, §5.1–5.3 |
| Procedimentos Operacionais | [PROCEDURE-FIN-00001](../../../../../business-inputs/business-documents/business-procedures/PROCEDURE-FIN-00001-CALCULO-DE-IMPOSTOS-CORPORATIVOS.md) SOP-001 a SOP-017 |
| Catálogo de Regras de Negócio | [RULES-CATALOG-FIN-00001](../../../../../business-inputs/business-documents/business-rules/RULES-FIN-00001-CALCULO-DE-IMPOSTOS-CORPORATIVOS.md) 60 regras (DEF-001 a ACT-007) |
| Especificação OpenAPI | [tax-rates-api.yaml](../api/tax-rates-api.yaml) v1.0.0 |
| Modelo de Dados (ERD) | [erd.md](../architecture/erd.md) |
| Documentação de Domínio | [domain.md](../domain/domain.md) |
