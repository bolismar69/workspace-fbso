# FACTORY-COMPARISON — Matriz Comparativa de Fábricas (Discovery-Level)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F6 · **Mode:** Discovery-Level
- **Versão:** 1.1 · **Data:** 31/07/2026 · **Status:** COMPARAÇÃO CONCLUÍDA (com prazo)

---

## Matriz Comparativa — Fábricas Aprovadas

| Fábrica | Total Horas | Prazo | Time Est. | Custo Est.* | QA% | Arch% | Consistência |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---|
| **Stefanini** | 16.000 | 4 meses | 25p | R$ 1.6M | 11.9% | 10.5% | ✅ Perfeita |
| **CI&T** | 52.000 | 5 meses | 65p | R$ 5.2M | 9.1% | 7.7% | ⚠️ Divergente |

---

## Análise por Critério

### 1. Custo Total (Peso 30%)

| Fábrica | Horas | Custo Est. | Nota |
|:---|:---:|:---|:---:|
| Stefanini | 16.000 | R$ 1.6M | 🟢 10 |
| CI&T | 52.000 | R$ 5.2M | 🟡 5 |

### 2. Prazo de Entrega (Peso 30%)

| Fábrica | Prazo | Time Necessário | Consistência | Nota |
|:---|:---:|:---:|:---|:---:|
| Stefanini | 4 meses | 25 pessoas | ✅ Viável | 🟢 10 |
| CI&T | 5 meses | 65 pessoas | ⚠️ 3× mais pessoas | 🟡 6 |

### 3. Qualidade Técnica (Peso 25%)

| Fábrica | QA% | Arch% | Completude | Nota |
|:---|:---:|:---:|:---|:---:|
| Stefanini | 11.9% | 10.5% | Completa | 🟢 8 |
| CI&T | 9.1% | 7.7% | Completa | 🟡 7 |

### 4. QA/Arch Balanceado (Peso 15%)

| Fábrica | QA% Dev | Arch% Tot | Nota |
|:---|:---:|:---:|:---:|
| Stefanini | 11.9% | 10.5% | 🟢 8 |
| CI&T | 9.1% | 7.7% | 🟡 6 |

---

## Ranking Final

| # | Fábrica | Custo (30%) | Prazo (30%) | Qualidade (25%) | QA/Arch (15%) | **Nota Ponderada** |
|:---:|:---|:---:|:---:|:---:|:---:|:---:|
| 🥇 | **Stefanini** | 10×0.30 | 10×0.30 | 8×0.25 | 8×0.15 | **9.20** |
| 🥈 | CI&T | 5×0.30 | 6×0.30 | 7×0.25 | 6×0.15 | **5.95** |

---

## Recomendação de Seleção

**🏆 Fábrica recomendada: STEFANINI**

**Justificativa:**
1. **Menor custo:** 16.000h (R$ 1.6M) — 3.25× menor que CI&T
2. **Menor prazo:** 4 meses — 1 mês mais rápida
3. **Prazo consistente:** 16.000h ÷ 4 meses = 25 pessoas (time realista para 4 épicos)
4. **Melhor QA/Arch:** 11.9% QA, 10.5% Arch — melhores índices
5. **Nota final 54% superior** à segunda colocada (9.20 vs 5.95)

---

## Fábricas Rejeitadas (F5)

| Fábrica | Motivo da Rejeição |
|:---|:---|
| Capgemini | Prazo irreal: 48.000h em 4 meses (75 pessoas) |
| TOTVS | Prazo impossível: 64.000h em 3-4 meses (100-133 pessoas) |
| Deloitte | QA 2.5% + Arch 2.4% — abaixo do mínimo |
| Infosys | QA 3.0% + Arch 4.8% — abaixo do mínimo |
| TCS | QA 5.9% — abaixo do mínimo |
| Overlabs | Outlier extremo: 132.800h (8.3× Stefanini) |

🤖 *Sourcing & Factory Bidding — Fase 6*
