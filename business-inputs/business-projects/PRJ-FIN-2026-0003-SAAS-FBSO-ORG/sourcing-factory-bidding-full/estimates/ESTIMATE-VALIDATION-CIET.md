# ESTIMATE-VALIDATION-CIET — Validação DTA

- **Fábrica:** CI&T
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Modo:** Full — Projeto Completo
- **Data:** 31/07/2026
- **Baseline PERT:** 7,300h (c/ 20% contingência)

---

## 1. Estimativa Apresentada

| Épico | horas_dev | horas_qa | horas_arch | horas_devops | horas_gestao | Total |
|:---|---:|---:|---:|---:|---:|---:|
| EP-0001 — Portal Admin - Dashboard + Contas | 11,600h | 320h | 320h | 320h | 360h | **12,920h** |
| EP-0002 — Clientes e Assinaturas - Tenants + Planos + Audito | 11,600h | 320h | 320h | 320h | 360h | **12,920h** |
| EP-0003 — RBAC - Usuários + Papéis + Permissões + Visibilida | 11,600h | 320h | 320h | 320h | 360h | **12,920h** |
| EP-0004 — Portal do Cliente - Auth + Onboarding + BUs + Catá | 11,600h | 320h | 320h | 320h | 360h | **12,920h** |
| **Total** | **46,400h** | **1,280h** | **1,280h** | **1,280h** | **1,440h** | **51,680h** |


---

## 2. Métricas DTA

| Regra | Valor | Threshold | Status |
|:---|---:|---:|:---|
| QA Global | 2.5% | ≥25% | ❌ |
| Arquitetura | 2.5% | ≥5% | ❌ |
| Prazo Declarado | 6 meses | — | — |
| Time Estimado | 45 pessoas | — | — |
| Prazo Calculado | 7.2 meses | — | `51,680 / (45 × 160) = 7.2` |
| PIB (Proximidade Baseline) 🆕 | 0.00 (Nota 1) | ≥0.50 | 🔴 |
| Divergência Prazo×Horas | 20% | ≤50% | ✅ |

---

## 3. Não-Conformidades

- **QA:** QA abaixo de 25% — tratado como overhead fixo em vez de proporcional ao esforço
- **Arch:** Arquitetura abaixo de 5% — subinvestimento técnico

---

- **PIB:** PIB Score 0.00 (<0.25) — descolamento extremo da baseline interna PERT (7,300h). Estimativa 608% acima da referência.

## 4. Comparação com Baseline PERT

| Fonte | Horas | Desvio |
|:---|---:|---:|
| PERT Downstream (F8) | 7,300h | — |
| **CI&T** | **51,680h** | **+608%** |

---

## 5. Veredito

**🔴 REJEITADA** — Motivo: QA+Arch

**Ação:** Contatar CI&T para realinhamento. Solicitar:
1. Redistribuição de horas com QA ≥ 25% e Arch ≥ 5%
2. Racional detalhado de estimativa na coluna `comentarios`
3. Revisão do prazo declarado para consistência com horas e time

---

🤖 *Validação DTA — Fase 5 do Sourcing & Factory Bidding (Full Mode)*
