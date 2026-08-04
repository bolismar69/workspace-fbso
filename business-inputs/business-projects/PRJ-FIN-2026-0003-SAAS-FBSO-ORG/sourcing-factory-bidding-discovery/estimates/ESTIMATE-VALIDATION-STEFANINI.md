# ESTIMATE-VALIDATION-STEFANINI.md

| Campo | Detalhe |
|-------|---------|
| **Fábrica** | STEFANINI |
| **Arquivo** | ESTIMATION-SCHEMA-STEFANINI.csv |
| **Total Horas** | 4.679h |
| **Valor Total** | R$ 842.220 |
| **Time** | 6 pessoas |
| **Veredito** | ⚠️ APROVADA COM RESSALVA |

---

## Resultados DTA

| Regra | Resultado | Status |
|-------|-----------|:------:|
| QA Balanceado (Global) | 826h QA / 2.755h Dev = 30,0% | ✅ |
| QA por Épico | Todos épicos ≥ 29,9% | ✅ |
| Arquitetura/SRE | (136+136) / (2755+136+136+826) = 7,1% | ✅ |
| Outliers Cross-Fábrica | 4.679h dentro de [2.468h, 7.404h] | ✅ |
| Consistência Prazo×Horas | 2/4 épicos > 50% divergência | 🔴 |
| PIB vs Baseline | 71,4% acima (4.679h vs 2.730h baseline) | 🟡 |

## Detalhamento por Épico

| Épico | Horas | QA% | Prazo | Prazo Calc | Divergência |
|-------|:-----:|:---:|:-----:|:----------:|:-----------:|
| EP-0001 | 832h | 30,0% | 2 meses | 0,87 meses | 130,5% 🔴 |
| EP-0002 | 1.428h | 30,0% | 2 meses | 1,49 meses | 34,4% ⚠️ |
| EP-0003 | 1.172h | 30,0% | 1 mês | 1,22 meses | 18,1% ✅ |
| EP-0004 | 1.247h | 29,9% | 2 meses | 1,30 meses | 53,9% 🔴 |

## Não-Conformidades

| ID | Descrição | Severidade |
|----|-----------|:----------:|
| NC-STE-01 | Prazo EP-0001 declarado 2 meses vs calculado 0,87 meses (130,5%) | 🔴 |
| NC-STE-02 | Prazo EP-0004 declarado 2 meses vs calculado 1,30 meses (53,9%) | 🔴 |
| NC-STE-03 | Horas 71% acima da baseline (4.679h vs 2.730h) | 🟡 |

## Ações Recomendadas

1. **Ajustar prazo EP-0001:** 832h com 6 pessoas = 0,87 meses. Prazo declarado de 2 meses é 2,3× o calculado. Reduzir prazo ou justificar.
2. **Ajustar prazo EP-0004:** 1.247h com 6 pessoas = 1,30 meses. Prazo de 2 meses é 54% acima.
3. **EP-0003 está consistente:** 18,1% de divergência — dentro do limite. Serve como referência de boa prática.
