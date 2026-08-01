# RISK-ADJUSTED-ESTIMATE — Estimativa Ajustada a Risco

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Data:** 31/07/2026
- **Fase:** F10 — Downstream Architecture Refinement
- **Baseline:** [BOTTOM-UP-PERT-ESTIMATE (F8)](./BOTTOM-UP-PERT-ESTIMATE.md) — CONGELADA
- **Ajuste aplicado sobre:** PERT F8, NUNCA sobre ROM upstream

---

## 1. Matriz de Riscos

| Risco | Prob. | Impacto (h) | Valor Esperado | Ação |
|:---|:---:|:---:|:---:|:---|
| Frontend sem dev dedicado até 01/11 | 70% | +800h | **560h** | Contratar freelancer; priorizar backend |
| RBAC multi-tenant complexo (Keycloak+Kong) | 50% | +400h | **200h** | Spike técnico 40h; consultoria externa |
| Istio/Kong/Keda requer especialista | 40% | +300h | **120h** | Davi + consultoria; treinamento interno |
| Escopo creep (novas US descobertas) | 50% | +500h | **250h** | Gate de controle de mudança no M3 |
| Time reduzido limita paralelismo | 60% | +300h | **180h** | Priorização rigorosa; Should Haves postergados |
| Junior (Maria) em tarefas críticas | 30% | +100h | **30h** | Tasks complexidade 1; pair programming |

**Valor Esperado Total do Risco:** Σ(Prob × Impacto) = **1,340h**

---

## 2. Cenários Ajustados

| Cenário | Horas | h-m | Contingência | Composição |
|:---|---:|---:|:---|:---|
| **Conservador** | 6,552h | 41 h-m | 15% | PERT + riscos mitigados |
| **PERT (média)** | 6,077h | 38 h-m | — | Estimativa base (F8) |
| **Pessimista** | 7,121h | 45 h-m | 25% | PERT + todos os riscos materializados |

### Incluindo Valor Esperado do Risco

| Cenário | Horas | h-m |
|:---|---:|---:|
| PERT + Σ Risco Esperado | 7,417h | 46 h-m |
| PERT + Riscos Materializados (100%) | 8,477h | 53 h-m |

---

## 3. Análise de Sensibilidade

### Top 3 Riscos por Impacto

| # | Risco | Valor Esperado | % do Total |
|:---|:---|---:|---:|
| 1 | **Frontend sem dev dedicado** | 560h | 42% |
| 2 | **Escopo creep** | 250h | 19% |
| 3 | **RBAC complexo** | 200h | 15% |
| | **Demais (3 riscos)** | 330h | 24% |

> 🔑 **87% do risco total concentra-se em 3 riscos.** Mitigar o gargalo de frontend (contratar freelancer) e o escopo creep (gate M3) elimina mais da metade do valor esperado de risco.

---

## 4. Recomendações

1. **Mitigar Risco #1 (Frontend):** Alocar orçamento para freelancer frontend por 3 meses (~$15-25k) — reduz 560h de risco para ~100h
2. **Mitigar Risco #3 (RBAC):** Spike técnico de 40h antes do Sprint 04 — reduz 200h de risco para ~50h
3. **Mitigar Risco #2 (Escopo):** Gate de controle de mudança rígido no M3; qualquer nova US requer trade-off (remove outra US do escopo)
4. **Contingência recomendada:** 20% (média entre Conservador 15% e Pessimista 25%) = **~1,215h adicionais**

### Estimativa Recomendada para o Comitê

| Métrica | Valor |
|:---|---:|
| **Estimativa Base (PERT)** | **6,077h / 38 h-m** |
| Contingência Recomendada | 20% (+1,215h) |
| **Total Recomendado** | **~7,300h / 46 h-m** |
| **Duração (time 7 FTE efetivo)** | **~6.5 meses** |
| **Confiança** | ±15-20% |

---

🤖 *Ajuste aplicado exclusivamente sobre o PERT (F8) — Fase 10 do Downstream Architecture Refinement*
