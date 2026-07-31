# ESTIMATE-VALIDATION — Validação DTA de Estimativas (Discovery-Level)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F5 · **Mode:** Discovery-Level
- **Versão:** 1.0 · **Data:** 31/07/2026 · **Status:** VALIDAÇÃO CONCLUÍDA

---

## Regras de Validação DTA (Discovery-Level — Épicos)

| Regra | Critério | Ação se não atender |
|:---|:---|:---|
| **QA Balanceado** | QA ≥ 20% por épico | ⚠️ Risco de Débito Técnico |
| **QA Global** | QA ≥ 25% do total de horas | ⚠️ Risco de Subinvestimento em Qualidade |
| **Arquitetura/SRE** | Arch ≥ 5% do total geral de horas | ⚠️ Risco de Subinvestimento Técnico |
| **Formato** | Colunas obrigatórias preenchidas conforme schema | ❌ REJEITADA |
| **Consistência Prazo×Horas** | `prazo_calculado = total_horas / (time_estimado × 160h)`. Divergência >50% entre prazo declarado e calculado → ⚠️ Alerta | 🔍 Revisão manual |
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

## Resultados por Fábrica

| # | Fábrica | Total Horas | Dev | QA | QA% Dev | Arch | Arch% Tot | Formato | Veredito |
|:---|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---|
| 1 | **Stefanini** | 16.000 | 12.800 | 1.520 | 11.9% ⚠️ | 1.680 | 10.5% ✅ | ✅ | 🟡 APROVADA c/ ressalva QA |
| 2 | **Capgemini** | 48.000 | 40.000 | 4.000 | 10.0% ⚠️ | 4.000 | 8.3% ✅ | ✅ | 🟡 APROVADA c/ ressalva QA |
| 3 | **CI&T** | 52.000 | 44.000 | 4.000 | 9.1% ⚠️ | 4.000 | 7.7% ✅ | ✅ | 🟡 APROVADA c/ ressalva QA |
| 4 | **TOTVS** | 64.000 | 52.800 | 5.520 | 10.5% ⚠️ | 5.680 | 8.9% ✅ | ✅ | 🟡 APROVADA c/ ressalva QA |
| 5 | **Deloitte** | 84.000 | 80.000 | 2.000 | 2.5% 🔴 | 2.000 | 2.4% 🔴 | ✅ | 🔴 REJEITADA — QA e Arch abaixo |
| 6 | **Infosys** | 97.400 | 90.000 | 2.700 | 3.0% 🔴 | 4.700 | 4.8% 🔴 | ✅ | 🔴 REJEITADA — QA abaixo |
| 7 | **TCS** | 104.000 | 92.800 | 5.520 | 5.9% 🔴 | 5.680 | 5.5% ✅ | ✅ | 🔴 REJEITADA — QA abaixo |
| 8 | **Overlabs** | 132.800 | 98.000 | 14.000 | 14.3% ⚠️ | 20.800 | 15.7% ✅ | ✅ | 🔴 REJEITADA — outlier extremo |

---

## Análise de Outliers

| Métrica | Valor |
|:---|:---|
| **Mediana cross-fábrica** | 74.000 h |
| **Limite inferior (-50%)** | 37.000 h |
| **Limite superior (+50%)** | 111.000 h |
| **Stefanini** | 16.000 h ⚠️ Abaixo do limite inferior (subestimado?) |
| **Overlabs** | 132.800 h 🔴 Acima do limite superior (superfaturado?) |

---

## Fábricas Aprovadas (4 de 8)

| # | Fábrica | Total Horas | Obs |
|:---|:---|:---:|:---|
| 1 | **Stefanini** | 16.000 h | ⚠️ Abaixo do ROM — possível subestimativa |
| 2 | **Capgemini** | 48.000 h | ✅ Dentro da faixa esperada |
| 3 | **CI&T** | 52.000 h | ✅ Dentro da faixa esperada |
| 4 | **TOTVS** | 64.000 h | ✅ Dentro da faixa esperada |

**4 fábricas aprovadas** seguem para F6 (Comparação).
**4 fábricas rejeitadas** (Deloitte, Infosys, TCS, Overlabs) — QA/Arch abaixo do mínimo ou outlier extremo.

🤖 *Sourcing & Factory Bidding — Fase 5*
