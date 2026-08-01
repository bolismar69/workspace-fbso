# ESTIMATE-VALIDATION-INFOSYS — Validação DTA

- **Fábrica:** Infosys
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Modo:** Full — Projeto Completo
- **Data:** 31/07/2026
- **Baseline PERT:** 7,300h (c/ 20% contingência)

---

## 1. Estimativa Apresentada

| Épico | horas_dev | horas_qa | horas_arch | horas_devops | horas_gestao | Total |
|:---|---:|---:|---:|---:|---:|---:|
| EP-0001 — Portal Admin - Dashboard + Contas | 1,600h | 320h | 320h | 320h | 360h | **2,920h** |
| EP-0002 — Clientes e Assinaturas - Tenants + Planos + Audito | 1,600h | 320h | 320h | 320h | 360h | **2,920h** |
| EP-0003 — RBAC - Usuários + Papéis + Permissões + Visibilida | 1,600h | 320h | 320h | 320h | 360h | **2,920h** |
| EP-0004 — Portal do Cliente - Auth + Onboarding + BUs + Catá | 1,600h | 320h | 320h | 320h | 360h | **2,920h** |
| **Total** | **6,400h** | **1,280h** | **1,280h** | **1,280h** | **1,440h** | **11,680h** |


---

## 2. Métricas DTA

| Regra | Valor | Threshold | Status |
|:---|---:|---:|:---|
| QA Global | 11.0% | ≥25% | ❌ |
| Arquitetura | 11.0% | ≥5% | ✅ |
| Prazo Declarado | 3 meses | — | — |
| Time Estimado | 15 pessoas | — | — |
| Prazo Calculado | 4.9 meses | — | `11,680 / (15 × 160) = 4.9` |
| PIB (Proximidade Baseline) 🆕 | 0.40 (Nota 3) | ≥0.50 | ⚠️ |
| Divergência Prazo×Horas | 62% | ≤50% | ❌ |

---

## 3. Não-Conformidades

- **QA:** QA abaixo de 25% — tratado como overhead fixo em vez de proporcional ao esforço
- **Prazo:** Divergência Prazo×Horas >50% — prazo declarado inconsistente com horas e time

---

- **PIB:** PIB Score 0.40 (<0.50) — desvio significativo da baseline interna PERT (7,300h). Estimativa 60% acima da referência.

## 4. Comparação com Baseline PERT

| Fonte | Horas | Desvio |
|:---|---:|---:|
| PERT Downstream (F8) | 7,300h | — |
| **Infosys** | **11,680h** | **+60%** |

---

## 5. Veredito

**🔴 REJEITADA** — Motivo: QA+Prazo

**Ação:** Contatar Infosys para realinhamento. Solicitar:
1. Redistribuição de horas com QA ≥ 25% e Arch ≥ 5%
2. Racional detalhado de estimativa na coluna `comentarios`
3. Revisão do prazo declarado para consistência com horas e time

---

🤖 *Validação DTA — Fase 5 do Sourcing & Factory Bidding (Full Mode)*
