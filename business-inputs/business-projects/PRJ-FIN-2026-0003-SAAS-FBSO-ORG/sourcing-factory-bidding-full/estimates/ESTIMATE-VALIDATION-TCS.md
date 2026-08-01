# ESTIMATE-VALIDATION-TCS — Validação DTA

- **Fábrica:** TCS
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Modo:** Full — Projeto Completo
- **Data:** 31/07/2026
- **Baseline PERT:** 7,300h (c/ 20% contingência)

---

## 1. Estimativa Apresentada

| Épico | horas_dev | horas_qa | horas_arch | horas_devops | horas_gestao | Total |
|:---|---:|---:|---:|---:|---:|---:|
| EP-0001 — Portal Admin - Dashboard + Contas | 31,600h | 320h | 320h | 320h | 360h | **32,920h** |
| EP-0002 — Clientes e Assinaturas - Tenants + Planos + Audito | 31,600h | 320h | 320h | 320h | 360h | **32,920h** |
| EP-0003 — RBAC - Usuários + Papéis + Permissões + Visibilida | 31,600h | 320h | 320h | 320h | 360h | **32,920h** |
| EP-0004 — Portal do Cliente - Auth + Onboarding + BUs + Catá | 31,600h | 320h | 320h | 320h | 360h | **32,920h** |
| **Total** | **126,400h** | **1,280h** | **1,280h** | **1,280h** | **1,440h** | **131,680h** |


---

## 2. Métricas DTA

| Regra | Valor | Threshold | Status |
|:---|---:|---:|:---|
| QA Global | 1.0% | ≥25% | ❌ |
| Arquitetura | 1.0% | ≥5% | ❌ |
| Prazo Declarado | 9 meses | — | — |
| Time Estimado | 93 pessoas | — | — |
| Prazo Calculado | 8.8 meses | — | `131,680 / (93 × 160) = 8.8` |
| PIB (Proximidade Baseline) 🆕 | 0.00 (Nota 1) | ≥0.50 | 🔴 |
| Divergência Prazo×Horas | 2% | ≤50% | ✅ |

---

## 3. Não-Conformidades

- **QA:** QA abaixo de 25% — tratado como overhead fixo em vez de proporcional ao esforço
- **Arch:** Arquitetura abaixo de 5% — subinvestimento técnico
- **Outlier:** Total de horas fora de ±50% da mediana cross-fábrica (27,855h)

---

- **PIB:** PIB Score 0.00 (<0.25) — descolamento extremo da baseline interna PERT (7,300h). Estimativa 1704% acima da referência.

## 4. Comparação com Baseline PERT

| Fonte | Horas | Desvio |
|:---|---:|---:|
| PERT Downstream (F8) | 7,300h | — |
| **TCS** | **131,680h** | **+1704%** |

---

## 5. Veredito

**🔴 REJEITADA** — Motivo: QA+Arch+Outlier

**Ação:** Contatar TCS para realinhamento. Solicitar:
1. Redistribuição de horas com QA ≥ 25% e Arch ≥ 5%
2. Racional detalhado de estimativa na coluna `comentarios`
3. Revisão do prazo declarado para consistência com horas e time

---

🤖 *Validação DTA — Fase 5 do Sourcing & Factory Bidding (Full Mode)*
