# ESTIMATE-VALIDATION-TOTVS.md

| Campo | Detalhe |
|-------|---------|
| **Fábrica** | TOTVS |
| **Arquivo** | ESTIMATION-SCHEMA-TOTVS.csv |
| **Total Horas** | 3.576h |
| **Valor Total** | R$ 715.200 |
| **Time** | 5 pessoas |
| **Veredito** | ⚠️ APROVADA COM RESSALVA |

---

## Resultados DTA

| Regra | Resultado | Status |
|-------|-----------|:------:|
| QA Balanceado (Global) | 541h QA / 2.170h Dev = 24,9% | ⚠️ Borderline |
| QA por Épico | Todos épicos ≥ 24,9% | ✅ |
| Arquitetura/SRE | (107+107) / (2170+107+107+541) = 7,3% | ✅ |
| Outliers Cross-Fábrica | 3.576h dentro de [2.468h, 7.404h] | ✅ |
| Consistência Prazo×Horas | 2/4 épicos > 50% divergência | 🔴 |
| PIB vs Baseline | 31,0% acima (3.576h vs 2.730h baseline) | 🏆 MELHOR |

## Detalhamento por Épico

| Épico | Horas | QA% | Prazo | Prazo Calc | Divergência |
|-------|:-----:|:---:|:-----:|:----------:|:-----------:|
| EP-0001 | 741h | 24,9% | 2 meses | 0,93 meses | 115,7% 🔴 |
| EP-0002 | 1.023h | 25,0% | 1 mês | 1,28 meses | 21,9% ✅ |
| EP-0003 | 873h | 24,9% | 1 mês | 1,09 meses | 8,2% ✅ |
| EP-0004 | 939h | 24,9% | 2 meses | 1,17 meses | 70,6% 🔴 |

## Não-Conformidades

| ID | Descrição | Severidade |
|----|-----------|:----------:|
| NC-TOT-01 | QA Global 24,9% — 0,1pp abaixo do limite de 25% | ⚠️ Borderline |
| NC-TOT-02 | Prazo EP-0001 declarado 2 meses vs calculado 0,93 meses (115,7%) | 🔴 |
| NC-TOT-03 | Prazo EP-0004 declarado 2 meses vs calculado 1,17 meses (70,6%) | 🔴 |

## Pontos Fortes

| ID | Descrição |
|----|-----------|
| PF-TOT-01 | 🏆 **Melhor PIB:** 31% de desvio da baseline — o menor entre as 4 fábricas |
| PF-TOT-02 | 🏆 **Menor custo total:** R$ 715.200 |
| PF-TOT-03 | **Melhor consistência Prazo×Horas:** EP-0002 (21,9%) e EP-0003 (8,2%) dentro do limite |
| PF-TOT-04 | **Melhor relação custo×hora:** ~R$ 200/h |

## Ações Recomendadas

1. **Ajustar QA:** Adicionar 2h de QA em qualquer épico para atingir 25% global. Impacto: ~R$ 400 no orçamento total.
2. **Corrigir prazo EP-0001 e EP-0004:** Alinhar com o prazo calculado ou aumentar o time declarado.
