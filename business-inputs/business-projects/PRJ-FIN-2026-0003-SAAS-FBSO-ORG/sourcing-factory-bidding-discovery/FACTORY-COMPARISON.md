# FACTORY-COMPARISON — Matriz Comparativa de Fábricas (Discovery-Level)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F6 · **Mode:** Discovery-Level
- **Versão:** 1.0 · **Data:** 31/07/2026 · **Status:** COMPARAÇÃO CONCLUÍDA

---

## Matriz Comparativa — Fábricas Aprovadas

| Fábrica | Total Horas | Prazo Declarado | Custo Est.* | QA% Dev | Arch% Tot | Variação vs ROM |
|:---|:---:|:---:|:---:|:---:|:---:|:---|
| **Stefanini** | 16.000 | ⚠️ Não informado | R$ 1.6M | 11.9% | 10.5% | +108% acima ROM |
| **Capgemini** | 48.000 | ⚠️ Não informado | R$ 4.8M | 10.0% | 8.3% | +525% acima ROM |
| **CI&T** | 52.000 | ⚠️ Não informado | R$ 5.2M | 9.1% | 7.7% | +550% acima ROM |
| **TOTVS** | 64.000 | ⚠️ Não informado | R$ 6.4M | 10.5% | 8.9% | +700% acima ROM |

> ⚠️ **Prazo não informado:** As estimativas foram recebidas antes da adição da coluna `prazo_entrega_meses` ao schema. Solicitar às fábricas que preencham a coluna de prazo para refinar a comparação.

*\* Custo estimado a R$100/h (referência de mercado)*

---

## Análise por Critério

### 1. Custo Total (Peso 35%)

| Fábrica | Horas | Nota |
|:---|:---:|:---:|
| Stefanini | 16.000 | 🟢 10 |
| Capgemini | 48.000 | 🟡 6 |
| CI&T | 52.000 | 🟡 5 |
| TOTVS | 64.000 | 🟠 4 |

### 2. Prazo Estimado (Peso 25%)

| Fábrica | h-m | Time 11p | Nota |
|:---|:---:|:---|:---:|
| Stefanini | 100 | ~9 meses | 🟢 9 |
| Capgemini | 300 | ~27 meses | 🟠 4 |
| CI&T | 325 | ~30 meses | 🟠 3 |
| TOTVS | 400 | ~36 meses | 🔴 2 |

### 3. Qualidade Técnica (Peso 25%)

| Fábrica | QA% | Arch% | Completude | Nota |
|:---|:---:|:---:|:---|:---:|
| Stefanini | 11.9% | 10.5% | Completa | 🟢 8 |
| Capgemini | 10.0% | 8.3% | Completa | 🟡 7 |
| CI&T | 9.1% | 7.7% | Completa | 🟡 7 |
| TOTVS | 10.5% | 8.9% | Completa | 🟡 7 |

### 4. QA/Arch Balanceado (Peso 15%)

| Fábrica | QA% Dev | Arch% Tot | Nota |
|:---|:---:|:---:|:---:|
| Stefanini | 11.9% | 10.5% | 🟢 8 |
| TOTVS | 10.5% | 8.9% | 🟡 7 |
| Capgemini | 10.0% | 8.3% | 🟡 7 |
| CI&T | 9.1% | 7.7% | 🟡 6 |

---

## Ranking Final

| # | Fábrica | Custo (35%) | Prazo (25%) | Qualidade (25%) | QA/Arch (15%) | **Nota Ponderada** | Recomendação |
|:---:|:---|:---:|:---:|:---:|:---:|:---:|:---|
| 🥇 | **Stefanini** | 10×0.35 | 9×0.25 | 8×0.25 | 8×0.15 | **8.95** | ✅ RECOMENDADA |
| 🥈 | Capgemini | 6×0.35 | 4×0.25 | 7×0.25 | 7×0.15 | **5.90** | Alternativa |
| 🥉 | CI&T | 5×0.35 | 3×0.25 | 7×0.25 | 6×0.15 | **5.15** | Alternativa |
| 4 | TOTVS | 4×0.35 | 2×0.25 | 7×0.25 | 7×0.15 | **4.70** | Terceira opção |

---

## Recomendação de Seleção

**🏆 Fábrica recomendada: STEFANINI**

**Justificativa:**
1. **Menor custo:** 16.000 horas (R$ 1.6M estimado) — 3× menor que a segunda colocada
2. **Menor prazo:** ~9 meses com time de 11 pessoas — alinhado com expectativa do negócio (6-8 meses)
3. **Melhor balanceamento QA/Arch:** 11.9% QA e 10.5% Arch — os melhores índices entre as 4 aprovadas
4. **Alinhamento com ROM:** 100 h-m está acima do ROM interno (29-50 h-m), mas é a estimativa mais próxima

**⚠️ Atenção:** O valor de Stefanini (16.000h / 100 h-m) está abaixo da mediana cross-fábrica (74.000h). Recomenda-se:
- Validar se o escopo compreendido pela Stefanini cobre todos os 4 épicos + 15 soluções
- Solicitar detalhamento da estimativa por solução (S01-S15) para confirmar cobertura
- Incluir cláusula contratual de variação máxima sobre a estimativa

---

## Fábricas Rejeitadas (F5)

| Fábrica | Total | Motivo |
|:---|:---:|:---|
| Deloitte | 84.000 | QA 2.5% + Arch 2.4% — abaixo do mínimo |
| Infosys | 97.400 | QA 3.0% + Arch 4.8% — abaixo do mínimo |
| TCS | 104.000 | QA 5.9% — abaixo do mínimo |
| Overlabs | 132.800 | Outlier extremo (8.3× Stefanini) |

🤖 *Sourcing & Factory Bidding — Fase 6*
