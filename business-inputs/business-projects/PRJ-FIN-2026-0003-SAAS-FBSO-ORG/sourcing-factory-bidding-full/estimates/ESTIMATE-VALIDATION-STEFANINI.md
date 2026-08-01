# ESTIMATE-VALIDATION-STEFANINI — Validação DTA

- **Fábrica:** Stefanini
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Modo:** Full — Projeto Completo
- **Data:** 31/07/2026
- **Baseline PERT:** 7,300h (c/ 20% contingência)

---

## 1. Estimativa Apresentada

| Épico | horas_dev | horas_qa | horas_arch | horas_devops | horas_gestao | Total |
|:---|---:|---:|---:|---:|---:|---:|
| EP-0001 — Portal Admin - Dashboard + Contas | 7,000h | 320h | 320h | 320h | 360h | **8,320h** |
| EP-0002 — Clientes e Assinaturas - Tenants + Planos + Audito | 6,000h | 320h | 320h | 320h | 360h | **7,320h** |
| EP-0003 — RBAC - Usuários + Papéis + Permissões + Visibilida | 5,250h | 320h | 320h | 320h | 360h | **6,570h** |
| EP-0004 — Portal do Cliente - Auth + Onboarding + BUs + Catá | 4,500h | 320h | 320h | 320h | 360h | **5,820h** |
| **Total** | **22,750h** | **1,280h** | **1,280h** | **1,280h** | **1,440h** | **28,030h** |


---

## 2. Métricas DTA

| Regra | Valor | Threshold | Status |
|:---|---:|---:|:---|
| QA Global | 4.6% | ≥25% | ❌ |
| Arquitetura | 4.6% | ≥5% | ❌ |
| Prazo Declarado | 3 meses | — | — |
| Time Estimado | 15 pessoas | — | — |
| Prazo Calculado | 11.7 meses | — | `28,030 / (15 × 160) = 11.7` |
| PIB (Proximidade Baseline) 🆕 | 0.00 (Nota 1) | ≥0.50 | 🔴 |
| Divergência Prazo×Horas | 289% | ≤50% | ❌ |

---

## 3. Não-Conformidades

- **QA:** QA abaixo de 25% — tratado como overhead fixo em vez de proporcional ao esforço
- **Arch:** Arquitetura abaixo de 5% — subinvestimento técnico
- **Prazo:** Divergência Prazo×Horas >50% — prazo declarado inconsistente com horas e time

---

- **PIB:** PIB Score 0.00 (<0.25) — descolamento extremo da baseline interna PERT (7,300h). Estimativa 284% acima da referência.

## 4. Comparação com Baseline PERT

| Fonte | Horas | Desvio |
|:---|---:|---:|
| PERT Downstream (F8) | 7,300h | — |
| **Stefanini** | **28,030h** | **+284%** |

---

## 5. Veredito

**🔴 REJEITADA** — Motivo: QA+Arch+Prazo

**Ação:** Contatar Stefanini para realinhamento. Solicitar:
1. Redistribuição de horas com QA ≥ 25% e Arch ≥ 5%
2. Racional detalhado de estimativa na coluna `comentarios`
3. Revisão do prazo declarado para consistência com horas e time

---

🤖 *Validação DTA — Fase 5 do Sourcing & Factory Bidding (Full Mode)*
