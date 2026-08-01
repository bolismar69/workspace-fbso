# RESOURCE-ALLOCATION-PLAN — Plano de Alocação de Recursos

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Data:** 31/07/2026
- **Fase:** F9 — Downstream Architecture Refinement
- **Fonte Única:** [BOTTOM-UP-PERT-ESTIMATE (F8)](./BOTTOM-UP-PERT-ESTIMATE.md) — todas as horas derivadas exclusivamente do PERT

---

## 1. Time Disponível

| Nome | Papel | Disp. | Atuação Principal | Épicos |
|:---|:---|:---:|:---|:---|
| Francisco Bolismar | Tech Lead / Full-Stack | 100% | Backend + Frontend + Arquitetura | Todos |
| Bruno Gratto | SA / Frontend | 100% | Frontend + Design System | EP-0001, EP-0004 |
| Felipe Canedas | QA | 100% | Testes + Automação | Todos |
| Davi Silva | DevOps | 100% | IaC + CI/CD + Deploy + Kong | Infra (todos) |
| Carlos Caldas | DB | 100% | PostgreSQL + Flyway + Performance | Dados (todos) |
| Mauro | BA | 50% | Refinamento + Critérios de Aceite | Negócio (todos) |
| Francisco (TL) | Gestão | 50% | Governança + Compliance Gates | Transversal |
| Tom Santos | FE Dedicado | 0% até 01/11 | Frontend | EP-0004 (após 01/11) |
| Maria Madalena | Junior | 100% | Testes + Documentação + Tasks simples | Todos |

---

## 2. Capacidade Mensal

| Métrica | Valor |
|:---|---:|
| Pessoas (carga total) | 9 |
| Pessoas (carga efetiva, considerando parciais) | ~7 FTE |
| Horas/mês (bruto) | 9 × 160 = 1,440h |
| Horas/mês (efetivo) | ~7 × 160 = ~1,120h |

---

## 3. Projeção de Duração

| Cenário | Horas Totais | Cap. Efetiva | Duração |
|:---|---:|---:|:---|
| Conservador (15% contingência) | 6,552h | 1,120h/mês | **5.9 meses** |
| **PERT (média ponderada)** | **6,077h** | **1,120h/mês** | **5.4 meses** |
| Pessimista (25% contingência) | 7,121h | 1,120h/mês | **6.4 meses** |

> ⚠️ **Alerta de Prazo:** O Project Charter define 14 semanas (3.2 meses). Com capacidade efetiva de ~7 FTE, mesmo o cenário PERT demanda **5.4 meses** — uma diferença de **+2.2 meses**. Isso indica que o prazo do Charter é agressivo e requer ou reforço de time ou redução de escopo.

---

## 4. Alocação por Épico

| Épico | Horas PERT | % Total | Perfil Crítico | Duração Estimada |
|:---|---:|---:|:---|:---|
| EP-0001 Portal Admin | 562h | 9.2% | Full-stack, Frontend | ~2 semanas |
| EP-0002 Clientes e Assinaturas | 1,374h | 22.6% | Backend, DB | ~5 semanas |
| EP-0003 RBAC | 1,437h | 23.6% | Backend, Keycloak, Kong | ~5 semanas |
| EP-0004 Portal Cliente | 1,843h | 30.3% | **Frontend (gargalo)** | ~7 semanas |
| Infra + Gestão + Contingência | 861h | 14.2% | DevOps, TL | Distribuído |

---

## 5. Gargalos Identificados

| Gargalo | Impacto | Duração Adicional | Recomendação |
|:---|:---:|:---:|:---|
| **Frontend sem dev dedicado até 01/11** | 🔴 Crítico | +3-4 semanas no EP-0004 | Antecipar Tom Santos ou contratar freelancer para Ago-Set |
| **Keycloak/Kong especialista (RBAC)** | 🟡 Moderado | +1-2 semanas no EP-0003 | Spike técnico 40h + consultoria externa |
| **Maria Madalena (Junior) em tasks complexas** | 🟢 Baixo | +0.5 semana | Manter em tasks de complexidade 1; parear com sênior |
| **Time reduzido limita paralelismo** | 🟡 Moderado | +2-3 semanas | Priorização rigorosa do MVP; Should Haves postergados |

---

## 6. Recomendações

1. **Reforço de Frontend:** Contratar freelancer ou antecipar Tom Santos para cobrir EP-0001 (dashboard) e EP-0004 (portal cliente)
2. **Spike RBAC:** Dedicar 40h de Bolismar antes do Sprint 04 para validar integração Kong↔Keycloak Service-ID/Token-ID com claims injection
3. **Postergar Should Haves:** FEAT-EP-0001-0003 (Alertas) e FEAT-EP-0004-0003 (Dashboard Cliente) podem ir para fase 2
4. **Revisão no M3:** Refinar estimativa para ±10% com base nos dados reais dos primeiros 2 meses

---

🤖 *Plano derivado exclusivamente do PERT (F8) — Fase 9 do Downstream Architecture Refinement*
