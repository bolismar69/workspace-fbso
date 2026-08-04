# FACTORY-COMPARISON.md — Matriz Comparativa + Recomendação
## Sourcing & Factory Bidding — Fase 6 — Bloco C

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FACTORY-COMPARISON-v1.0 |
| **Versão** | 1.0 — Discovery-Level |
| **Data** | 03 de agosto de 2026 |
| **Modo** | `discovery` |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em 03/08/2026 |

---

## 1. Matriz Comparativa Cross-Fábrica

### 1.1 Indicadores-Chave

| Indicador | CAPGEMINI | INFOSYS | STEFANINI | TOTVS |
|-----------|:---------:|:-------:|:---------:|:-----:|
| **Total Horas** | 5.193h | 7.259h | 4.679h | **3.576h** 🏆 |
| **Valor Total (R$)** | 2.077.200 | 2.177.700 | 842.220 | **715.200** 🏆 |
| **R$/hora** | ~400 | ~300 | ~180 | **~200** 🏆 |
| **Time Alocado** | 7 pess. | 8 pess. | 6 pess. | **5 pess.** 🏆 |
| **Prazo Calendar (soma)** | 8 meses | 7 meses | 7 meses | **6 meses** 🏆 |
| **QA% (Global)** | 34,9% | **40,0%** 🏆 | 30,0% | 24,9% |
| **Arch% (Global)** | 10,0% | **12,5%** 🏆 | 7,1% | 7,3% |
| **PIB Score** | 90,2% | 165,9% | 71,4% | **31,0%** 🏆 |
| **Consistência Prazo×Horas (média)** | 103,7% | **47,7%** 🏆 | 59,2% | 54,1% |

### 1.2 Comparação por Épico — Horas

| Épico | Baseline Mid | CAPGEMINI | INFOSYS | STEFANINI | TOTVS |
|-------|:-----------:|:---------:|:-------:|:---------:|:-----:|
| EP-0001 — Portal Admin | 490h | 1.044h | 1.400h | 832h | **741h** 🏆 |
| EP-0002 — Clientes e Assinaturas | 840h | 1.661h | 2.059h | 1.428h | **1.023h** 🏆 |
| EP-0003 — Governança e Permissões | 700h | 1.216h | 1.840h | 1.172h | **873h** 🏆 |
| EP-0004 — Experiência do Cliente | 700h | 1.272h | 1.960h | 1.247h | **939h** 🏆 |

### 1.3 Comparação por Épico — Valor (R$)

| Épico | CAPGEMINI | INFOSYS | STEFANINI | TOTVS |
|-------|----------:|--------:|----------:|------:|
| EP-0001 | 417.600 | 420.000 | 149.760 | **148.200** 🏆 |
| EP-0002 | 664.400 | 617.700 | 257.040 | **204.600** 🏆 |
| EP-0003 | 486.400 | 552.000 | 210.960 | **174.600** 🏆 |
| EP-0004 | 508.800 | 588.000 | 224.460 | **187.800** 🏆 |

---

## 2. Matriz de Decisão Ponderada (5 Critérios)

### 2.1 Metodologia de Pontuação

Cada critério recebe nota de 1 a 5 (5 = melhor). As notas são ponderadas pelos pesos definidos em `DTA-VALIDATION-STANDARDS.md` §3.2.

| Critério | Peso | Indicador | Melhor valor |
|----------|:----:|-----------|-------------|
| **Custo Total** | 27,5% | Valor total em R$ | Menor |
| **Prazo** | 22,5% | Prazo calendário (soma) | Menor |
| **Qualidade Técnica** | 20% | QA% + Arch% combinados | Maior |
| **PIB — Baseline** | 15% | Desvio da baseline interna | Menor |
| **Consistência Prazo×Horas** | 15% | Divergência média entre épicos | Menor |

### 2.2 Notas por Critério (1-5)

| Critério (peso) | CAPGEMINI | INFOSYS | STEFANINI | TOTVS |
|-----------------|:---------:|:-------:|:---------:|:-----:|
| Custo Total (27,5%) | 2,0 | 1,5 | 4,5 | **5,0** |
| Prazo (22,5%) | 2,0 | 3,5 | 3,5 | **5,0** |
| Qualidade Técnica (20%) | **5,0** | **5,0** | 4,0 | 3,5 |
| PIB — Baseline (15%) | 3,0 | 1,0 | 3,5 | **5,0** |
| Consistência Prazo×Horas (15%) | 1,0 | **5,0** | 3,0 | 4,0 |

### 2.3 Resultado Ponderado

| Fábrica | Custo (×0,275) | Prazo (×0,225) | Qualidade (×0,20) | PIB (×0,15) | Consist. (×0,15) | **Score Final** | **Ranking** |
|---------|:--------------:|:--------------:|:-----------------:|:-----------:|:----------------:|:---------------:|:-----------:|
| CAPGEMINI | 0,55 | 0,45 | 1,00 | 0,45 | 0,15 | **2,60** | 4º |
| INFOSYS | 0,41 | 0,79 | 1,00 | 0,15 | 0,75 | **3,10** | 3º |
| STEFANINI | 1,24 | 0,79 | 0,80 | 0,53 | 0,45 | **3,80** | 2º |
| **TOTVS** 🏆 | **1,38** | **1,13** | **0,70** | **0,75** | **0,60** | **4,55** | **1º** |

---

## 3. Análise por Critério

### 3.1 Custo Total (27,5%)

```
R$ 715k                                 R$ 842k              R$ 2.078k    R$ 2.178k
   ████████████████████                    ████████████████████  ████████████████████
   TOTVS                                  STEFANINI              CAPGEMINI    INFOSYS
```

TOTVS e STEFANINI têm custos na faixa de R$ 700-850k. CAPGEMINI e INFOSYS estão ~2,7× acima, na faixa de R$ 2,0-2,2M. A diferença entre o menor (TOTVS) e o maior (INFOSYS) é de **R$ 1.462.500** (3,0×).

### 3.2 Prazo (22,5%)

TOTVS declara 6 meses (menor prazo). CAPGEMINI declara 8 meses (maior). INFOSYS e STEFANINI empatam em 7 meses. Porém, todas as fábricas têm inconsistências Prazo×Horas — os prazos declarados estão sistematicamente superestimados em relação às horas.

### 3.3 Qualidade Técnica — QA + Arch (20%)

CAPGEMINI e INFOSYS lideram com QA 35-40% e Arch 10-12,5% — bem acima dos mínimos (25% e 5%). STEFANINI está na média (QA 30%, Arch 7,1%). TOTVS está no limite inferior (QA 24,9%, Arch 7,3%) mas ainda acima dos gates de rejeição (10% e 2%).

### 3.4 PIB — Proximidade à Baseline (15%)

TOTVS destaca-se com apenas **31%** de desvio da baseline interna — a métrica mais importante de aderência à realidade do projeto. STEFANINI (71%) e CAPGEMINI (90%) estão na faixa intermediária. INFOSYS (166%) está mais de 2,5× acima da baseline, sugerindo superdimensionamento significativo.

### 3.5 Consistência Prazo×Horas (15%)

INFOSYS lidera com a menor divergência média (47,7%). CAPGEMINI tem a pior consistência (103,7%), com 3 dos 4 épicos acima de 75% de divergência — os prazos declarados são 2-3× maiores que o calculado pelas horas.

---

## 4. Análise de Riscos por Fábrica

| Risco | CAPGEMINI | INFOSYS | STEFANINI | TOTVS |
|-------|:---------:|:-------:|:---------:|:-----:|
| **Superestimação de prazo** | 🔴 Alto | 🟡 Médio | 🟡 Médio | 🟡 Médio |
| **Custo acima do mercado** | 🔴 Alto | 🔴 Alto | 🟢 Baixo | 🟢 Baixo |
| **QA abaixo do mínimo** | 🟢 Baixo | 🟢 Baixo | 🟢 Baixo | 🟡 Médio |
| **Subinvestimento técnico** | 🟢 Baixo | 🟢 Baixo | 🟢 Baixo | 🟢 Baixo |
| **Desvio da baseline** | 🟡 Médio | 🔴 Alto | 🟡 Médio | 🟢 Baixo |
| **Time subdimensionado** | 🟢 Baixo | 🟢 Baixo | 🟢 Baixo | 🟡 Médio |

---

## 5. Recomendação de Seleção

### 🏆 Fábrica Recomendada: **TOTVS**

**Fundamentação:**

1. **Menor custo total:** R$ 715.200 — 3× menor que CAPGEMINI/INFOSYS, 15% menor que STEFANINI
2. **Melhor PIB (31%):** Estimativa mais próxima da baseline interna da FBSO.ORG — indica compreensão realista do escopo
3. **Menor prazo declarado:** 6 meses com time enxuto de 5 pessoas
4. **Boa consistência:** EP-0002 (21,9%) e EP-0003 (8,2%) dentro do limite de 30%
5. **Melhor relação custo×benefício:** R$ 200/hora com time multidisciplinar eficiente
6. **Risco gerenciável:** QA no limite (24,9%) — facilmente corrigível com +2h de QA. Prazo de EP-0001 e EP-0004 precisam de ajuste

### ⚠️ Ressalvas para Negociação com TOTVS

| # | Item | Ação |
|---|------|------|
| 1 | **QA Global** | Aumentar de 24,9% para ≥ 25% (+2h de QA) |
| 2 | **Prazo EP-0001** | 741h ÷ 5 pessoas = 0,93 meses. Prazo declarado de 2 meses é 2,1× acima. Revisar para 1-2 meses com justificativa |
| 3 | **Prazo EP-0004** | 939h ÷ 5 pessoas = 1,17 meses. Prazo declarado de 2 meses. Revisar para 1,5-2 meses com justificativa |
| 4 | **Time de 5 pessoas** | Validar se 5 pessoas (1 TL + 1 BE + 1 FE + 1 QA + 1 DevOps) é suficiente para entrega paralela de frontend e backend |

### 🥈 Segundo Colocado: **STEFANINI**

Caso a negociação com TOTVS não prospere, STEFANINI é a segunda opção:
- Custo competitivo (R$ 842.220)
- QA (30%) e Arch (7,1%) dentro dos limites
- PIB de 71,4% — aceitável para discovery
- Time de 6 pessoas (vs 5 da TOTVS) oferece mais folga

---

## 6. Resumo Executivo para o Comitê

| Critério | TOTVS (🏆) | STEFANINI (🥈) | CAPGEMINI | INFOSYS |
|----------|:----------:|:-------------:|:---------:|:-------:|
| Custo Total | R$ 715k | R$ 842k | R$ 2.078k | R$ 2.178k |
| Prazo | 6 meses | 7 meses | 8 meses | 7 meses |
| Qualidade QA+Arch | ⚠️ 24,9%+7,3% | ✅ 30%+7,1% | ✅ 35%+10% | ✅ 40%+12,5% |
| PIB (aderência) | 🏆 31% | ⚠️ 71% | ⚠️ 90% | 🔴 166% |
| Consistência | ⚠️ 54% | ⚠️ 59% | 🔴 104% | ⚠️ 48% |
| Score Ponderado | **4,55** 🏆 | 3,80 | 2,60 | 3,10 |

**Recomendação:** Contratar **TOTVS** como fábrica principal, com **STEFANINI** como backup. Negociar os 4 pontos de ressalva antes da assinatura do contrato.

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 03/08/2026 | Criação inicial: matriz comparativa 5 critérios, ranking ponderado, recomendação TOTVS com ressalvas, STEFANINI backup | PMO / Tech Lead |

---

🤖 *Sourcing & Factory Bidding — Fase 6. Matriz comparativa e recomendação de seleção.*
