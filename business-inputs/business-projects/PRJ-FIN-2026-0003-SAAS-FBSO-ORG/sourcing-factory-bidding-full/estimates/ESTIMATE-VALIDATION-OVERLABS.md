# ESTIMATE-VALIDATION-OVERLABS — Validação DTA

- **Fábrica:** Overlabs
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Modo:** Full — Projeto Completo
- **Data:** 31/07/2026
- **Baseline PERT:** 7,300h (c/ 20% contingência)

---

## 1. Estimativa Apresentada

| Épico | horas_dev | horas_qa | horas_arch | horas_devops | horas_gestao | Total |
|:---|---:|---:|---:|---:|---:|---:|
| EP-0001 — Portal Admin - Dashboard + Contas | 5,600h | 320h | 320h | 320h | 360h | **6,920h** |
| EP-0002 — Clientes e Assinaturas - Tenants + Planos + Audito | 5,600h | 320h | 320h | 320h | 360h | **6,920h** |
| EP-0003 — RBAC - Usuários + Papéis + Permissões + Visibilida | 5,600h | 320h | 320h | 320h | 360h | **6,920h** |
| EP-0004 — Portal do Cliente - Auth + Onboarding + BUs + Catá | 5,600h | 320h | 320h | 320h | 360h | **6,920h** |
| **Total** | **22,400h** | **1,280h** | **1,280h** | **1,280h** | **1,440h** | **27,680h** |


---

## 2. Métricas DTA

| Regra | Valor | Threshold | Status |
|:---|---:|---:|:---|
| QA Global | 4.6% | ≥25% | ❌ |
| Arquitetura | 4.6% | ≥5% | ❌ |
| Prazo Declarado | 4 meses | — | — |
| Time Estimado | 45 pessoas | — | — |
| Prazo Calculado | 3.8 meses | — | `27,680 / (45 × 160) = 3.8` |
| PIB (Proximidade Baseline) 🆕 | 0.00 (Nota 1) | ≥0.50 | 🔴 |
| Divergência Prazo×Horas | 4% | ≤50% | ✅ |

---

## 3. Não-Conformidades

- **QA:** QA abaixo de 25% — tratado como overhead fixo em vez de proporcional ao esforço
- **Arch:** Arquitetura abaixo de 5% — subinvestimento técnico

---

- **PIB:** PIB Score 0.00 (<0.25) — descolamento extremo da baseline interna PERT (7,300h). Estimativa 279% acima da referência.

## 4. Comparação com Baseline PERT

| Fonte | Horas | Desvio |
|:---|---:|---:|
| PERT Downstream (F8) | 7,300h | — |
| **Overlabs** | **27,680h** | **+279%** |

---

## 5. Veredito

**🔴 REJEITADA** — Motivo: QA+Arch

**Ação:** Contatar Overlabs para realinhamento. Solicitar:
1. Redistribuição de horas com QA ≥ 25% e Arch ≥ 5%
2. Racional detalhado de estimativa na coluna `comentarios`
3. Revisão do prazo declarado para consistência com horas e time

---

🤖 *Validação DTA — Fase 5 do Sourcing & Factory Bidding (Full Mode)*
