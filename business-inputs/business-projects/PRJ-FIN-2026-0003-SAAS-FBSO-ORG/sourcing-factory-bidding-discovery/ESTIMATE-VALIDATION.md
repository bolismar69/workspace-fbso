# ESTIMATE-VALIDATION.md — Validação DTA + PIB
## Sourcing & Factory Bidding — Fase 5 — Bloco C

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | ESTIMATE-VALIDATION-v1.0 |
| **Versão** | 1.0 — Discovery-Level |
| **Data** | 03 de agosto de 2026 |
| **Modo** | `discovery` |
| **Baseline PIB** | `DISCOVERY-LEVEL-ROM-ESTIMATE.md` — {3.180-4.508}h |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em 03/08/2026 |

---

## 1. Regras de Validação DTA + PIB

| Regra | Fórmula (Discovery) | Limite |
|-------|---------------------|--------|
| **QA Balanceado (Global)** | `Σhoras_qa / Σhoras_dev` | ≥ 25% ✅ / 10-25% ⚠️ / < 10% ❌ |
| **QA por Épico** | `horas_qa / horas_dev` por épico | ≥ 20% ✅ / < 20% ⚠️ |
| **Arquitetura/SRE** | `(Σarch + Σdevops) / (Σdev + Σarch + Σdevops + Σqa)` | ≥ 5% ✅ / 2-5% ⚠️ / < 2% ❌ |
| **Consistência Prazo×Horas** | `|prazo - prazo_calc| / prazo_calc` onde `prazo_calc = total_horas / (time × 160)` | ≤ 30% ✅ / 30-50% ⚠️ / > 50% 🔴 |
| **Outliers Cross-Fábrica** | `total_horas < mediana×0.5` ou `> mediana×1.5` | Dentro do limite ✅ / Fora ⚠️ |
| **PIB — Proximidade à Baseline** | `|total_fabrica - baseline_midpoint| / baseline_midpoint` | Menor = melhor |

**Baseline ROM (midpoints):**
| Épico | Baseline Midpoint |
|-------|:-----------------:|
| EP-0001 — Portal Admin Interno | 490h |
| EP-0002 — Clientes e Assinaturas | 840h |
| EP-0003 — Governança e Permissões | 700h |
| EP-0004 — Experiência do Cliente | 700h |
| **Total Baseline** | **2.730h** |

---

## 2. Resultados por Fábrica

### 2.1 Tabela Consolidada

| Fábrica | Total Horas | QA% (Global) | Arch% | PIB Score | Veredito |
|---------|:----------:|:------------:|:-----:|:---------:|----------|
| **CAPGEMINI** | 5.193h | 34,9% ✅ | 10,0% ✅ | 90,2% | ⚠️ APROVADA COM RESSALVA |
| **INFOSYS** | 7.259h | 40,0% ✅ | 12,5% ✅ | 165,9% | ⚠️ APROVADA COM RESSALVA |
| **STEFANINI** | 4.679h | 30,0% ✅ | 7,1% ✅ | 71,4% | ⚠️ APROVADA COM RESSALVA |
| **TOTVS** | 3.576h | 24,9% ⚠️ | 7,3% ✅ | 31,0% | ⚠️ APROVADA COM RESSALVA |

### 2.2 QA por Épico

| Fábrica | EP-0001 | EP-0002 | EP-0003 | EP-0004 | Global |
|---------|:-------:|:-------:|:-------:|:-------:|:------:|
| CAPGEMINI | 34,9% ✅ | 35,0% ✅ | 35,0% ✅ | 34,9% ✅ | 34,9% ✅ |
| INFOSYS | 40,0% ✅ | 40,0% ✅ | 40,0% ✅ | 40,0% ✅ | 40,0% ✅ |
| STEFANINI | 30,0% ✅ | 30,0% ✅ | 30,0% ✅ | 29,9% ✅ | 30,0% ✅ |
| TOTVS | 24,9% ✅ | 25,0% ✅ | 24,9% ✅ | 24,9% ✅ | 24,9% ⚠️ |

### 2.3 Consistência Prazo×Horas por Épico

| Fábrica | Time | EP-0001 | EP-0002 | EP-0003 | EP-0004 |
|---------|:----:|:-------:|:-------:|:-------:|:-------:|
| CAPGEMINI | 7 | 222% 🔴 | 32,6% ⚠️ | 84,0% 🔴 | 76,0% 🔴 |
| INFOSYS | 8 | 83,1% 🔴 | 37,9% ⚠️ | 39,1% ⚠️ | 30,6% ⚠️ |
| STEFANINI | 6 | 130,5% 🔴 | 34,4% ⚠️ | 18,1% ✅ | 53,9% 🔴 |
| TOTVS | 5 | 115,7% 🔴 | 21,9% ✅ | 8,2% ✅ | 70,6% 🔴 |

### 2.4 Valor por Épico (R$)

| Fábrica | EP-0001 | EP-0002 | EP-0003 | EP-0004 | **Total** |
|---------|--------:|--------:|--------:|--------:|----------:|
| CAPGEMINI | 417.600 | 664.400 | 486.400 | 508.800 | **2.077.200** |
| INFOSYS | 420.000 | 617.700 | 552.000 | 588.000 | **2.177.700** |
| STEFANINI | 149.760 | 257.040 | 210.960 | 224.460 | **842.220** |
| TOTVS | 148.200 | 204.600 | 174.600 | 187.800 | **715.200** |

---

## 3. PIB — Proximidade à Baseline Interna

### 3.1 PIB por Épico (desvio do midpoint)

| Fábrica | EP-0001 (490h) | EP-0002 (840h) | EP-0003 (700h) | EP-0004 (700h) | Média PIB |
|---------|:-------------:|:-------------:|:-------------:|:-------------:|:---------:|
| CAPGEMINI | 113,1% | 97,7% | 73,7% | 81,7% | 91,6% |
| INFOSYS | 185,7% | 145,1% | 162,9% | 180,0% | 168,4% |
| STEFANINI | 69,8% | 70,0% | 67,4% | 78,1% | 71,3% |
| TOTVS | **51,2%** | **21,8%** | **24,7%** | **34,1%** | **33,0%** 🏆 |

### 3.2 PIB Total

| Fábrica | Total Horas | Baseline Midpoint | Desvio | PIB Score | Nota |
|---------|:----------:|:-----------------:|:------:|:---------:|:----:|
| CAPGEMINI | 5.193h | 2.730h | +2.463h | 90,2% | 3 |
| INFOSYS | 7.259h | 2.730h | +4.529h | 165,9% | 1 |
| STEFANINI | 4.679h | 2.730h | +1.949h | 71,4% | 4 |
| **TOTVS** | **3.576h** | **2.730h** | **+846h** | **31,0%** 🏆 | **5** |

> 🏆 **Melhor PIB:** TOTVS — apenas 31% acima da baseline interna, consistente em todos os épicos.
> ⚠️ **Pior PIB:** INFOSYS — 166% acima da baseline, com todos os épicos acima de 145% de desvio.

---

## 4. Detecção de Outliers

| Métrica | Valor |
|---------|-------|
| **Mediana Cross-Fábrica** | 4.936h |
| **Limite Inferior (×0,5)** | 2.468h |
| **Limite Superior (×1,5)** | 7.404h |

| Fábrica | Total Horas | Status Outlier |
|---------|:----------:|:--------------:|
| CAPGEMINI | 5.193h | ✅ Dentro |
| INFOSYS | 7.259h | ✅ Dentro |
| STEFANINI | 4.679h | ✅ Dentro |
| TOTVS | 3.576h | ✅ Dentro |

> Nenhuma fábrica detectada como outlier cross-fábrica. Todas dentro da faixa [2.468h, 7.404h].

---

## 5. Resumo de Não-Conformidades

| # | ID | Regra | Fábricas Afetadas | Severidade |
|---|-----|------|-------------------|:----------:|
| NC1 | QA Global < 25% | QA Balanceado §2.1 | TOTVS (24,9%) | ⚠️ Borderline |
| NC2 | Prazo inconsistente > 50% | Consistência Prazo×Horas §2.3 | CAPGEMINI (3 épicos), INFOSYS (1), STEFANINI (2), TOTVS (2) | 🔴 Crítico |
| NC3 | PIB > 100% | Proximidade à Baseline §2.6 | INFOSYS (165,9%) | 🟡 Moderado |

---

## 6. Veredito Final

| Fábrica | Veredito | Justificativa |
|---------|----------|---------------|
| **CAPGEMINI** | ⚠️ APROVADA COM RESSALVA | QA e Arch passam. Prazo severamente inconsistente (3/4 épicos >50%). Horas 90% acima da baseline. Precisa revisar prazos por épico. |
| **INFOSYS** | ⚠️ APROVADA COM RESSALVA | QA (40%) e Arch (12,5%) excelentes. Horas 166% acima da baseline — mais que o dobro. Prazo inconsistente em 1 épico. Precisa justificar dimensionamento. |
| **STEFANINI** | ⚠️ APROVADA COM RESSALVA | QA (30%) e Arch (7,1%) passam. Horas 71% acima da baseline. Prazo inconsistente em 2 épicos. |
| **TOTVS** | ⚠️ APROVADA COM RESSALVA | 🏆 Melhor PIB (31%). QA no limite (24,9% — precisa de +1h de QA para atingir 25%). Prazo inconsistente em 2 épicos. Melhor relação custo×aderência. |

> **Nota:** Nenhuma fábrica foi REJEITADA nesta rodada. Todas passam nos gates críticos (QA ≥ 10%, Arch ≥ 2%). As ressalvas concentram-se em inconsistência de prazo e superestimação vs. baseline interna. A Fase 5b (Retrospectiva PIB) **não é necessária** (há fábricas aprovadas).

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 03/08/2026 | Criação inicial: validação DTA+PIB das 4 fábricas. Regras aplicadas: QA, Arch, Prazo×Horas, Outliers, PIB. Arquivos individuais gerados. | PMO / Tech Lead |

---

🤖 *Sourcing & Factory Bidding — Fase 5. Validação DTA completa das estimativas recebidas.*
