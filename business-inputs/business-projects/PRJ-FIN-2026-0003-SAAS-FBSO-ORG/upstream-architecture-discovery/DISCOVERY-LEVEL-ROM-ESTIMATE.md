# ROM-ESTIMATE — Estimativa ROM +-50% (Discovery)
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F11 — Bloco D · **Versão:** 1.0 · **Data:** 30/07/2026

## 1. Matriz de Esforço por Solução

| Solução | Complexidade | Esforço (h-m) | Faixa (+-50%) |
|:---|:---:|:---:|:---|
| S01 Backend API | Alta | 8-12 | 4-18 |
| S02 Frontend Portal | Média | 4-6 | 2-9 |
| S03 PostgreSQL + Redis | Média | 2-3 | 1-4.5 |
| S04 Keycloak | Média | 2-3 | 1-4.5 |
| S05 Kong API Gateway | Alta | 2-3 | 1-4.5 |
| S06 Cloudflare | Baixa | 0.5-1 | 0.25-1.5 |
| S07-S10 Observabilidade | Média | 3-5 | 1.5-7.5 |
| S11 IaC | Média | 2-3 | 1-4.5 |
| S12 Istio | Alta | 2-3 | 1-4.5 |
| S13-S14 Keda+Karpenter | Média | 1-2 | 0.5-3 |
| S15 CI/CD | Média | 1-2 | 0.5-3 |
| Gestão + Governança | — | 2-3 | 1-4.5 |
| **Total** | | **29.5-46** | **15-69** |

## 2. Premissas

- Time de 11 profissionais, alguns em carga parcial
- Stack corporativa já definida (sem POC de tecnologia)
- DigitalOcean como único provedor (sem complexidade multi-cloud)
- Frontend dedicado a partir de 01/11 (antes disso, full-stack cobre)
- Infraestrutura como código desde o início (sem dívida técnica de IaC)

## 3. Riscos e Mitigações

| Risco | Impacto na Estimativa | Mitigação |
|:---|:---:|:---|
| Istio requer especialista | +20% | Treinamento ou contratar consultoria |
| Time reduzido limita paralelismo | +30% | Priorização rigorosa do MVP |
| Frontend sem dev dedicado até 01/11 | +15% no portal | Backend primeiro; frontend após 01/11 |
| Elastic Stack + Loki = custo extra | +10% | Consolidar em uma stack após MVP |

## 4. ROM Consolidado

| Cenário | Homem-Mês | Duração (time 11 pessoas) | Custo Estimado |
|:---|:---:|:---:|:---|
| **Otimista** (-50%) | 15 | ~3 meses | — |
| **Provável** | 30-38 | ~6-8 meses | — |
| **Pessimista** (+50%) | 69 | ~12 meses | — |

## 5. Recomendação Técnica

**Parecer do Discovery Team:** Projeto **viável tecnicamente**. Stack corporativa madura. Riscos gerenciáveis. Recomendamos **Go-Ahead** com as seguintes condições:

1. Contratar ou treinar especialista Istio antes do deploy em produção
2. Iniciar desenvolvimento do backend imediatamente; frontend dedicado a partir de 01/11
3. Consolidar Elastic Stack + Loki em uma stack de logging após MVP
4. Revisar ROM após 3 meses de execução (refinar para +-20%)

🤖 *F11 — Upstream Architecture Discovery · Bloco D*
