# ESTIMATE-VALIDATION — Validação DTA + PIB (Discovery Mode)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Modo:** Discovery — Alto Nível (Épicos)
- **Data:** 31/07/2026 (atualizado com regra PIB v1.1)
- **Baseline PIB:** ROM Upstream 6,080h (provavel) — `upstream-architecture-discovery/DISCOVERY-LEVEL-ROM-ESTIMATE.md`

---

## 1. Regras de Validação DTA + PIB (v1.1)

| Regra | Critério | Ação se não atender |
|:---|:---|:---|
| **QA Balanceado** | QA ≥ 20% por épico | ⚠️ Risco de Débito Técnico |
| **QA Global** | QA ≥ 25% do total de horas | ❌ REJEITADA |
| **Arquitetura/SRE** | Arch ≥ 5% do total geral de horas | ❌ REJEITADA |
| **Formato** | Colunas obrigatórias preenchidas conforme schema | ❌ REJEITADA |
| **Consistência Prazo×Horas** | `prazo_calculado = total_horas / (time × 160h)`. Divergência >50% → ❌ | ❌ REJEITADA |
| **Outliers** | Total de horas dentro de ±50% da mediana cross-fábrica (78,500h discovery) | 🔍 Revisão manual |
| **PIB (Proximidade Baseline)** 🆕 | PIB Score ≥ 0.25. Baseline: ROM Upstream 6,080h | ⚠️ <0.50 alerta / 🔴 <0.25 rejeitada |

---

## 2. Resultados por Fábrica

| # | Fábrica | Total Horas | QA% | Arch% | Prazo | PIB Score 🆕 | Veredito Original | Veredito c/ PIB |
|:---|:---|:---:|:---:|:---:|:---:|:---:|:---|:---|
| 1 | **Stefanini** | 16,000h | 9.5% | 10.5% | 4 meses | 0.00 🔴 | 🟢 APROVADA | 🔴 REJEITADA (PIB) |
| 2 | **Capgemini** | 48,000h | 8.3% | 8.3% | 4 meses | 0.00 🔴 | 🔴 REJEITADA | 🔴 REJEITADA |
| 3 | **CI&T** | 52,000h | 7.7% | 7.7% | 5 meses | 0.00 🔴 | 🟡 APROVADA c/ ressalva | 🔴 REJEITADA (PIB) |
| 4 | **TOTVS** | 64,000h | 8.6% | 8.9% | 3-4 meses | 0.00 🔴 | 🔴 REJEITADA | 🔴 REJEITADA |
| 5 | **Deloitte** | 84,000h | 2.4% | 2.4% | — | 0.00 🔴 | 🔴 REJEITADA | 🔴 REJEITADA |
| 6 | **Infosys** | 97,400h | 2.8% | 4.8% | — | 0.00 🔴 | 🔴 REJEITADA | 🔴 REJEITADA |
| 7 | **TCS** | 104,000h | 5.3% | 5.5% | — | 0.00 🔴 | 🔴 REJEITADA | 🔴 REJEITADA |
| 8 | **Overlabs** | 132,800h | 10.5% | 15.7% | — | 0.00 🔴 | 🔴 REJEITADA | 🔴 REJEITADA |

> ⚠️ **Nota:** O PIB Score de todas as fábricas é 0.00 porque mesmo a menor estimativa (Stefanini, 16,000h) está 163% acima da baseline ROM de 6,080h. No modo discovery, com ROM ±50%, espera-se que estimativas fiquem entre 3,040h e 9,120h.

---

## 3. PIB por Épico (Discovery Mode)

Baseline ROM por épico (extraída do `DISCOVERY-LEVEL-ROM-ESTIMATE.md` §2):

| Épico | ROM Baseline | % do Total |
|:---|---:|---:|
| EP-0001 Portal Admin | ~1,280h | 21% |
| EP-0002 Clientes e Assinaturas | ~1,760h | 29% |
| EP-0003 RBAC e Permissões | ~1,520h | 26% |
| EP-0004 Portal do Cliente | ~1,760h | 29% |

### PIB Score por Épico e Fábrica

| Fábrica | EP-0001 | EP-0002 | EP-0003 | EP-0004 | PIB Total |
|:---|---:|---:|---:|---:|:---:|
| Stefanini | 4,000h (0.00) | 4,000h (0.00) | 4,000h (0.00) | 4,000h (0.00) | 0.00 |
| Capgemini | 12,000h (0.00) | 12,000h (0.00) | 12,000h (0.00) | 12,000h (0.00) | 0.00 |
| CI&T | 13,000h (0.00) | 13,000h (0.00) | 13,000h (0.00) | 13,000h (0.00) | 0.00 |
| Deloitte | 21,000h (0.00) | 21,000h (0.00) | 21,000h (0.00) | 21,000h (0.00) | 0.00 |
| Infosys | 34,400h (0.00) | 21,000h (0.00) | 21,000h (0.00) | 21,000h (0.00) | 0.00 |
| Overlabs | 33,200h (0.00) | 33,200h (0.00) | 33,200h (0.00) | 33,200h (0.00) | 0.00 |
| TCS | 26,000h (0.00) | 26,000h (0.00) | 26,000h (0.00) | 26,000h (0.00) | 0.00 |
| TOTVS | 16,000h (0.00) | 16,000h (0.00) | 16,000h (0.00) | 16,000h (0.00) | 0.00 |

> 🔴 **PIB por épico zerado em todas as fábricas.** A menor estimativa por épico (Stefanini, 4,000h) está 127-212% acima da baseline ROM (1,280-1,760h por épico). Além disso, **todas as 8 fábricas** apresentam flat estimates — mesmo valor para os 4 épicos, indicando que nenhuma analisou o escopo individualmente. Apenas Infosys diferenciou EP-0001 (34,400h) dos demais (21,000h).

---

## 4. Análise PIB Total (Discovery Mode)

| Fábrica | Horas | Desvio ROM | PIB Score | Nota |
|:---|---:|---:|:---:|:---:|
| Stefanini | 16,000h | +163% | 0.00 | 1 |
| Capgemini | 48,000h | +689% | 0.00 | 1 |
| CI&T | 52,000h | +755% | 0.00 | 1 |
| TOTVS | 64,000h | +953% | 0.00 | 1 |
| Deloitte | 84,000h | +1,282% | 0.00 | 1 |
| Infosys | 97,400h | +1,502% | 0.00 | 1 |
| TCS | 104,000h | +1,611% | 0.00 | 1 |
| Overlabs | 132,800h | +2,084% | 0.00 | 1 |

**Conclusão PIB:** Nenhuma fábrica atinge PIB ≥ 0.25 no modo discovery. Todas superestimaram significativamente em relação ao ROM interno. Isso é esperado no modo discovery (ROM ±50% é uma estimativa inicial de baixa precisão), mas o descolamento de 163-2,084% indica que as fábricas não se basearam no escopo macro (épicos) com a mesma interpretação do time de arquitetura.

---

## 5. Comparação com Rodada Full

| Métrica | Discovery (ROM) | Full (PERT) |
|:---|---:|:---|
| Baseline Interna | 6,080h | 7,300h |
| Fábricas aprovadas (original) | 2/8 | 0/8 |
| Fábricas aprovadas (c/ PIB) | 0/8 | 0/8 |
| Menor estimativa | Stefanini 16,000h | Capgemini/Deloitte/Infosys 11,680h |
| Mediana cross-fábrica | 78,500h | 27,855h |

---

🤖 *Validação DTA + PIB — Atualizada com regra PIB v1.1. Baseline: ROM Upstream 6,080h.*
