# ESTIMATE-VALIDATION-CAPGEMINI.md

| Campo | Detalhe |
|-------|---------|
| **Fábrica** | CAPGEMINI |
| **Arquivo** | ESTIMATION-SCHEMA-CAPGEMINI.csv |
| **Total Horas** | 5.193h |
| **Valor Total** | R$ 2.077.200 |
| **Time** | 7 pessoas |
| **Veredito** | ⚠️ APROVADA COM RESSALVA |

---

## Resultados DTA

| Regra | Resultado | Status |
|-------|-----------|:------:|
| QA Balanceado (Global) | 956h QA / 2.735h Dev = 34,9% | ✅ |
| QA por Épico | Todos épicos ≥ 34,9% | ✅ |
| Arquitetura/SRE | (273+135) / (2735+273+135+956) = 10,0% | ✅ |
| Outliers Cross-Fábrica | 5.193h dentro de [2.468h, 7.404h] | ✅ |
| Consistência Prazo×Horas | 3/4 épicos > 50% divergência | 🔴 |
| PIB vs Baseline | 90,2% acima (5.193h vs 2.730h baseline) | 🔴 |

## Detalhamento por Épico

| Épico | Horas | QA% | Prazo | Prazo Calc | Divergência |
|-------|:-----:|:---:|:-----:|:----------:|:-----------:|
| EP-0001 | 1.044h | 34,9% | 3 meses | 0,93 meses | 222% 🔴 |
| EP-0002 | 1.661h | 35,0% | 1 mês | 1,48 meses | 32,6% ⚠️ |
| EP-0003 | 1.216h | 35,0% | 2 meses | 1,09 meses | 84,0% 🔴 |
| EP-0004 | 1.272h | 34,9% | 2 meses | 1,14 meses | 76,0% 🔴 |

## Não-Conformidades

| ID | Descrição | Severidade |
|----|-----------|:----------:|
| NC-CAP-01 | Prazo EP-0001 declarado 3 meses vs calculado 0,93 meses (222%) | 🔴 |
| NC-CAP-02 | Prazo EP-0003 declarado 2 meses vs calculado 1,09 meses (84%) | 🔴 |
| NC-CAP-03 | Prazo EP-0004 declarado 2 meses vs calculado 1,14 meses (76%) | 🔴 |
| NC-CAP-04 | Horas totais 90% acima da baseline interna (5.193h vs 2.730h) | 🟡 |

## Ações Recomendadas

1. **Revisar prazos:** Para 7 pessoas, 5.193h exigem ~4,6 meses de calendário. Prazos declarados somam 8 meses. Ajustar `prazo_entrega_meses` ou justificar tempo ocioso entre épicos.
2. **Justificar superestimação:** 90% acima da baseline — apresentar racional detalhado do dimensionamento por épico.
3. **Manter QA e Arch:** Percentuais de QA (35%) e Arquitetura (10%) estão excelentes.
