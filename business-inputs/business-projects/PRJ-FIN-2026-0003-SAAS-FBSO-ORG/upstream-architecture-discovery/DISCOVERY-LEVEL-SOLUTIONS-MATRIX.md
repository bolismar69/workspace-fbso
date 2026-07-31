# DISCOVERY-LEVEL-SOLUTIONS-MATRIX — Matriz Solução×Disciplina (Discovery)
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F9 — Bloco C · **Versão:** 1.0 · **Data:** 30/07/2026

## Matriz Solução × Disciplina × Complexidade

| Solução | ARCH | SEC | DATA | DEVOPS | TEST | INFRA | Complexidade |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| S01 Backend API | Alta | Alta | Alta | Média | Alta | Média | **Alta** |
| S02 Frontend Portal | Média | Média | — | Média | Alta | Baixa | **Média** |
| S03 PostgreSQL | Média | Alta | Alta | Baixa | Baixa | Média | **Média** |
| S04 Keycloak | Média | Alta | — | Média | Média | Média | **Média** |
| S05 Kong | Alta | Alta | — | Alta | Média | Alta | **Alta** |
| S06 Cloudflare | Baixa | Média | — | Baixa | — | Média | **Baixa** |
| S07-S10 Observabilidade | Baixa | Baixa | — | Alta | — | Média | **Média** |
| S11 IaC (Terraform+Ansible) | Baixa | Baixa | — | Alta | — | Alta | **Média** |
| S12 Istio | Alta | Alta | — | Alta | — | Alta | **Alta** |
| S13 Keda | — | — | — | Alta | — | Média | **Média** |
| S14 Karpenter | — | — | — | Alta | — | Alta | **Média** |
| S15 GitHub Actions | — | Baixa | — | Alta | Média | — | **Média** |

## Gaps e Riscos Identificados

| Gap | Disciplina | Risco |
|:---|:---|:---|
| Istio requer conhecimento especializado | DEVOPS | Curva de aprendizado alta para time reduzido |
| Elastic Stack + Loki = redundância parcial | DEVOPS | Avaliar se ambas são necessárias ou consolidar |
| Sem ferramenta de API Management além do Kong | ARCH | OK para MVP; revisar no crescimento |

🤖 *F9 — Upstream Architecture Discovery*
