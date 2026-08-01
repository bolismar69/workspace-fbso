# FACTORY-COMPARISON — Matriz Comparativa de Fábricas (Discovery-Level)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F6 · **Mode:** Discovery-Level
- **Versão:** 1.3 · **Data:** 31/07/2026 · **Status:** ATUALIZADO COM 5 CRITÉRIOS (PIB INTEGRADO)

---

## Critérios de Avaliação (DTA-VALIDATION-STANDARDS §3.2 — Discovery)

| # | Critério | Peso |
|:---:|:---|:---:|
| 1 | Custo Total | 25% |
| 2 | Prazo de Entrega | 25% |
| 3 | Qualidade Técnica (QA+Arch) | 20% |
| 4 | PIB — Proximidade à Baseline Interna 🆕 | 15% |
| 5 | Consistência Prazo×Horas | 15% |

---

## Matriz Comparativa — Fábricas Aprovadas na F5 (antes do PIB)

| Fábrica | Total Horas | Prazo | Time | Valor | QA% | Arch% |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|
| **Stefanini** | 16.000 | 4 meses | 25 pessoas | R$ 1.6M | 11.9% | 10.5% |
| **CI&T** | 52.000 | 5 meses | 65 pessoas | R$ 5.2M | 9.1% | 7.7% |

> ⚠️ **Nota:** `valor_estimado` inferido (R$100/h) pois as fábricas não declararam. Com o schema v1.3, `valor_estimado` e `time_estimado_pessoas` são obrigatórios.

---

## Análise por Critério

### 1. Custo Total (Peso 25%)

| Fábrica | Horas | Valor Estimado | Nota |
|:---|:---:|:---|:---:|
| Stefanini | 16.000 | R$ 1.6M | 🟢 10 |
| CI&T | 52.000 | R$ 5.2M | 🟡 5 |

### 2. Prazo de Entrega (Peso 25%)

| Fábrica | Prazo | Time Necessário | Consistência | Nota |
|:---|:---:|:---:|:---|:---:|
| Stefanini | 4 meses | 25 pessoas | ✅ Viável | 🟢 10 |
| CI&T | 5 meses | 65 pessoas | ⚠️ Time 3× maior | 🟡 6 |

### 3. Qualidade Técnica QA+Arch (Peso 20%)

| Fábrica | QA% | Arch% | Completude | Nota |
|:---|:---:|:---:|:---|:---:|
| Stefanini | 11.9% | 10.5% | Completa | 🟢 8 |
| CI&T | 9.1% | 7.7% | Completa | 🟡 7 |

### 4. PIB — Proximidade à Baseline Interna (Peso 15%) 🆕

Baseline: ROM Upstream 6,080h (provavel).

| Fábrica | Horas | Desvio ROM | PIB Score | Nota |
|:---|:---:|:---:|:---:|:---:|
| Stefanini | 16.000h | +163% | 0.00 🔴 | 🔴 3 |
| CI&T | 52.000h | +755% | 0.00 🔴 | 🔴 1 |

> ⚠️ Ambas as fábricas têm PIB Score 0.00 (desvio >100% da baseline). Isso é esperado no modo discovery (ROM ±50%), mas penaliza fortemente no ranking.

### 5. Consistência Prazo×Horas (Peso 15%)

| Fábrica | Prazo Declarado | Prazo Calculado | Divergência | Nota |
|:---|:---:|:---:|:---:|:---:|
| Stefanini | 4 meses | 16.000/(25×160) = 4.0 | 0% ✅ | 🟢 10 |
| CI&T | 5 meses | 52.000/(65×160) = 5.0 | 0% ✅ | 🟢 10 |

---

## Ranking Final (5 critérios — v1.3)

| # | Fábrica | Custo (25%) | Prazo (25%) | Qualidade (20%) | PIB (15%) | Consist. (15%) | **Nota Ponderada** |
|:---:|:---|:---:|:---:|:---:|:---:|:---:|:---:|
| 🥇 | **Stefanini** | 10×0.25 | 10×0.25 | 8×0.20 | 3×0.15 | 10×0.15 | **8.05** |
| 🥈 | CI&T | 5×0.25 | 6×0.25 | 7×0.20 | 1×0.15 | 10×0.15 | **5.55** |

---

## Recomendação de Seleção

**🏆 Fábrica recomendada: STEFANINI**

**Justificativa:**
1. **Menor custo:** 16.000h (R$ 1.6M) — 3.25× menor que CI&T
2. **Menor prazo:** 4 meses — 1 mês mais rápida
3. **Time enxuto e consistente:** 25 pessoas (prazo calculado = declarado)
4. **Melhor QA/Arch:** 11.9% QA, 10.5% Arch — melhores índices entre as fábricas
5. **Nota final 45% superior** à segunda colocada (8.05 vs 5.55)

> ⚠️ **Ressalva PIB:** Ambas as fábricas têm PIB Score 0.00. Nenhuma se aproxima da baseline ROM de 6,080h. A Stefanini vence por ter a menor estimativa absoluta (16.000h), mas o descolamento de +163% da baseline indica que o discovery-level tem limitações de precisão. Recomenda-se executar a rodada **Full** (com PERT ±15-25%) para uma baseline mais precisa antes da contratação definitiva.

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

---

## Histórico de Alterações

| Versão | Data | Alteração |
|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial: matriz 4 critérios, Stefanini 9.20 vs CI&T 5.95 |
| 1.2 | 31/07/2026 | Adendo PIB v1.1 retroativo (sem integração ao ranking) |
| 1.3 | 31/07/2026 | **5 critérios integrados** (DTA-VALIDATION-STANDARDS §3.2): PIB 15% e Consistência 15% integrados ao cálculo ponderado. Ranking recalculado: Stefanini 8.05 vs CI&T 5.55. |
