# ESTIMATE-VALIDATION — Validação DTA de Estimativas (Discovery-Level)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F5 · **Mode:** Discovery-Level
- **Versão:** 1.1 · **Data:** 31/07/2026 · **Status:** VALIDAÇÃO CONCLUÍDA (com prazo)

---

## Regras de Validação DTA (Discovery-Level — Épicos)

| Regra | Critério | Ação se não atender |
|:---|:---|:---|
| **QA Balanceado** | QA ≥ 20% por épico | ⚠️ Risco de Débito Técnico |
| **QA Global** | QA ≥ 25% do total de horas | ⚠️ Risco de Subinvestimento em Qualidade |
| **Arquitetura/SRE** | Arch ≥ 5% do total geral de horas | ⚠️ Risco de Subinvestimento Técnico |
| **Formato** | Colunas obrigatórias preenchidas conforme schema | ❌ REJEITADA |
| **Consistência Prazo×Horas** | `prazo_calculado = total_horas / (time_estimado × 160h)`. Divergência >50% → ⚠️ | 🔍 Revisão manual |
| **Outliers** | Total de horas dentro de ±50% da mediana cross-fábrica | 🔍 Revisão manual |

---

## ROM Interno (Baseline)

| Épico | Complexidade | ROM (h-m) | ROM (horas ≈160h/mês) |
|:---|:---:|:---:|:---:|
| EP-0001 Portal Admin | Média | 6-10 | 960-1.600 |
| EP-0002 Clientes e Assinaturas | Alta | 8-14 | 1.280-2.240 |
| EP-0003 RBAC | Alta | 7-12 | 1.120-1.920 |
| EP-0004 Portal Cliente | Alta | 8-14 | 1.280-2.240 |
| **Total ROM** | | **29-50 h-m** | **4.640-8.000 h** |

---

## Resultados por Fábrica (com Prazo)

| # | Fábrica | Total Horas | Prazo (meses) | Time Est.* | QA% Dev | Arch% Tot | Consistência Prazo | Veredito |
|:---|:---|:---:|:---:|:---:|:---:|:---:|:---|:---|
| 1 | **Stefanini** | 16.000 | 4 | 25p | 11.9% | 10.5% | ✅ (calc: 4.0) | 🟢 APROVADA |
| 2 | **Capgemini** | 48.000 | 4 | 75p | 10.0% | 8.3% | 🔴 Divergente (calc: 12) | 🔴 REJEITADA — prazo irreal |
| 3 | **CI&T** | 52.000 | 5 | 65p | 9.1% | 7.7% | ⚠️ Divergente (calc: 13) | 🟡 APROVADA c/ ressalva |
| 4 | **TOTVS** | 64.000 | 3-4 | 100-133p | 10.5% | 8.9% | 🔴 Divergente (calc: 16-20) | 🔴 REJEITADA — prazo irreal |
| 5 | Deloitte | 84.000 | — | — | 2.5% | 2.4% | — | 🔴 REJEITADA — QA/Arch |
| 6 | Infosys | 97.400 | — | — | 3.0% | 4.8% | — | 🔴 REJEITADA — QA/Arch |
| 7 | TCS | 104.000 | — | — | 5.9% | 5.5% | — | 🔴 REJEITADA — QA abaixo |
| 8 | Overlabs | 132.800 | — | — | 14.3% | 15.7% | — | 🔴 REJEITADA — outlier |

*\*Time estimado = total_horas / (prazo_meses × 160h)*

---

## Análise de Consistência Prazo×Horas

| Fábrica | Horas | Prazo | Time Necessário | Análise |
|:---|:---:|:---:|:---:|:---|
| **Stefanini** | 16.000 | 4 meses | 25 pessoas | ✅ Viável — time enxuto, prazo realista |
| **Capgemini** | 48.000 | 4 meses | 75 pessoas | 🔴 Impossível — 75p em 4 meses para mesmo escopo |
| **CI&T** | 52.000 | 5 meses | 65 pessoas | ⚠️ Duvidoso — 65p, 3× mais que Stefanini |
| **TOTVS** | 64.000 | 3-4 meses | 100-133 pessoas | 🔴 Impossível — 100+p para mesmo escopo |

---

## Fábricas Aprovadas (2 de 8)

| # | Fábrica | Total Horas | Prazo | Nota |
|:---|:---|:---:|:---:|:---|
| 1 | **Stefanini** | 16.000 h | 4 meses | ✅ Única com prazo consistente |
| 2 | **CI&T** | 52.000 h | 5 meses | ⚠️ Prazo divergente (52.000h em 5 meses = 65 pessoas) |

**6 fábricas rejeitadas** — QA/Arch abaixo, outlier extremo, ou prazo inconsistente com horas.

⚠️ **Nota:** Capgemini (48.000h em 4 meses) e TOTVS (64.000h em 3-4 meses) declararam prazos fisicamente impossíveis para o volume de horas — seria necessário um time de 75-133 pessoas, incompatível com o escopo de 4 épicos.

🤖 *Sourcing & Factory Bidding — Fase 5*
