# ESTIMATE-VALIDATION-INFOSYS.md

| Campo | Detalhe |
|-------|---------|
| **Fábrica** | INFOSYS |
| **Arquivo** | ESTIMATION-SCHEMA-INFOSYS.csv |
| **Total Horas** | 7.259h |
| **Valor Total** | R$ 2.177.700 |
| **Time** | 8 pessoas |
| **Veredito** | ⚠️ APROVADA COM RESSALVA |

---

## Resultados DTA

| Regra | Resultado | Status |
|-------|-----------|:------:|
| QA Balanceado (Global) | 1.452h QA / 3.630h Dev = 40,0% | ✅ |
| QA por Épico | Todos épicos = 40,0% | ✅ |
| Arquitetura/SRE | (544+181) / (3630+544+181+1452) = 12,5% | ✅ |
| Outliers Cross-Fábrica | 7.259h dentro de [2.468h, 7.404h] | ✅ |
| Consistência Prazo×Horas | 1/4 épicos > 50% divergência | 🔴 |
| PIB vs Baseline | 165,9% acima (7.259h vs 2.730h baseline) | 🔴 |

## Detalhamento por Épico

| Épico | Horas | QA% | Prazo | Prazo Calc | Divergência |
|-------|:-----:|:---:|:-----:|:----------:|:-----------:|
| EP-0001 | 1.400h | 40,0% | 2 meses | 1,09 meses | 83,1% 🔴 |
| EP-0002 | 2.059h | 40,0% | 1 mês | 1,61 meses | 37,9% ⚠️ |
| EP-0003 | 1.840h | 40,0% | 2 meses | 1,44 meses | 39,1% ⚠️ |
| EP-0004 | 1.960h | 40,0% | 2 meses | 1,53 meses | 30,6% ⚠️ |

## Não-Conformidades

| ID | Descrição | Severidade |
|----|-----------|:----------:|
| NC-INF-01 | Prazo EP-0001 declarado 2 meses vs calculado 1,09 meses (83,1%) | 🔴 |
| NC-INF-02 | Horas totais 166% acima da baseline (7.259h vs 2.730h) — mais que o dobro | 🔴 |
| NC-INF-03 | Maior custo total entre as 4 fábricas (R$ 2.177.700) | 🟡 |

## Ações Recomendadas

1. **Justificar dimensionamento:** 7.259h é 2,66× a baseline. Apresentar racional técnico detalhado — o que justifica 7.259h para 18 funcionalidades discovery-level?
2. **Revisar prazo EP-0001:** 2 meses declarados para 1.400h com 8 pessoas — prazo calculado é 1,1 meses.
3. **QA e Arch excelentes:** 40% QA e 12,5% Arch estão acima dos mínimos. Manter.
