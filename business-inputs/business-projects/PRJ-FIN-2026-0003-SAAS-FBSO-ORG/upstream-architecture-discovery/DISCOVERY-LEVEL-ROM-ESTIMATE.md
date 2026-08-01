# ROM-ESTIMATE — Estimativa ROM +-50% (Discovery)
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F11 — Bloco D · **Versão:** 1.1 · **Data:** 31/07/2026

## 1. Matriz de Esforço por Solução

| Solução | Complexidade | Esforço (h-m) | Horas Aprox. | Faixa (+-50%) |
|:---|:---:|:---:|:---:|:---|
| S01 Backend API | Alta | 8-12 | ~1.280-1.920 | 4-18 h-m (640-2.880h) |
| S02 Frontend Portal | Média | 4-6 | ~640-960 | 2-9 h-m (320-1.440h) |
| S03 PostgreSQL + Redis | Média | 2-3 | ~320-480 | 1-4.5 h-m (160-720h) |
| S04 Keycloak | Média | 2-3 | ~320-480 | 1-4.5 h-m (160-720h) |
| S05 Kong API Gateway | Alta | 2-3 | ~320-480 | 1-4.5 h-m (160-720h) |
| S06 Cloudflare | Baixa | 0.5-1 | ~80-160 | 0.25-1.5 h-m (40-240h) |
| S07-S10 Observabilidade | Média | 3-5 | ~480-800 | 1.5-7.5 h-m (240-1.200h) |
| S11 IaC | Média | 2-3 | ~320-480 | 1-4.5 h-m (160-720h) |
| S12 Istio | Alta | 2-3 | ~320-480 | 1-4.5 h-m (160-720h) |
| S13-S14 Keda+Karpenter | Média | 1-2 | ~160-320 | 0.5-3 h-m (80-480h) |
| S15 CI/CD | Média | 1-2 | ~160-320 | 0.5-3 h-m (80-480h) |
| Gestão + Governança | — | 2-3 | ~320-480 | 1-4.5 h-m (160-720h) |
| **Total** | | **29.5-46** | **~4.720-7.360** | **15-69 h-m (2.400-11.040h)** |

> 💡 **Horas Aprox.** = h-m × 160h. Valores arredondados para referência rápida.

## 2. Visão por Épico

Distribuição do esforço total estimado entre os 4 épicos do projeto, com mapeamento para as soluções envolvidas.

| Épico | Soluções | Complexidade | Esforço (h-m) | Horas Aprox. | % do Total |
|:---|:---|:---:|:---:|:---:|:---:|
| EP-0001 Portal Admin | S01, S03 | Média | 6-10 | ~960-1.600 | ~21% |
| EP-0002 Clientes e Assinaturas | S01, S03, S04, S07 | Alta | 8-14 | ~1.280-2.240 | ~29% |
| EP-0003 RBAC e Permissões | S01, S04 | Alta | 7-12 | ~1.120-1.920 | ~26% |
| EP-0004 Portal do Cliente | S01, S02, S03, S04 | Alta | 8-14 | ~1.280-2.240 | ~29% |
| **Total** | | | **29-50** | **~4.640-8.000** | |

> 💡 Os épicos não são independentes — compartilham soluções (ex: S01 Backend é usado em todos). O rateio entre épicos é aproximado e baseado na proporção de features e User Stories de cada um.

## 3. Premissas

- Time de 11 profissionais, alguns em carga parcial
- Stack corporativa já definida (sem POC de tecnologia)
- DigitalOcean como único provedor (sem complexidade multi-cloud)
- Frontend dedicado a partir de 01/11 (antes disso, full-stack cobre)
- Infraestrutura como código desde o início (sem dívida técnica de IaC)

## 4. ROM Consolidado

| Cenário | Homem-Mês | Horas Aprox. | Duração (time 11 pessoas) |
|:---|:---:|:---:|:---|
| **Otimista** (-50%) | 15 | ~2.400 | ~3 meses |
| **Provável** | 30-38 | **~4.800-6.080** | ~6-8 meses |
| **Pessimista** (+50%) | 69 | ~11.040 | ~12 meses |

> 💡 O valor de **~6.080h / 38 h-m** (cenário provável) é a **baseline de referência** usada como PIB (Proximidade à Baseline Interna) no modo discovery do Sourcing & Factory Bidding. Comparar com PERT Downstream (~7.300h / 46 h-m) no modo full.

## 5. Riscos e Mitigações

| Risco | Impacto na Estimativa | Mitigação |
|:---|:---:|:---|
| Istio requer especialista | +20% | Treinamento ou contratar consultoria |
| Time reduzido limita paralelismo | +30% | Priorização rigorosa do MVP |
| Frontend sem dev dedicado até 01/11 | +15% no portal | Backend primeiro; frontend após 01/11 |
| Elastic Stack + Loki = custo extra | +10% | Consolidar em uma stack após MVP |

## 6. Recomendação Técnica

**Parecer do Discovery Team:** Projeto **viável tecnicamente**. Stack corporativa madura. Riscos gerenciáveis. Recomendamos **Go-Ahead** com as seguintes condições:

1. Contratar ou treinar especialista Istio antes do deploy em produção
2. Iniciar desenvolvimento do backend imediatamente; frontend dedicado a partir de 01/11
3. Consolidar Elastic Stack + Loki em uma stack de logging após MVP
4. Revisar ROM após 3 meses de execução (refinar para +-20%)

🤖 *F11 — Upstream Architecture Discovery · Bloco D · v1.1 com visão por épico e horas aproximadas*
