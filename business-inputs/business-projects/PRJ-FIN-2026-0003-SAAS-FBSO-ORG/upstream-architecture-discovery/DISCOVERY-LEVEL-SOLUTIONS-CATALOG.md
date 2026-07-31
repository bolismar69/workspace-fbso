# DISCOVERY-LEVEL-SOLUTIONS-CATALOG — Catálogo de Soluções (Discovery)
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F8 — Bloco C · **Versão:** 1.0 · **Data:** 30/07/2026

## Catálogo Macro de Soluções

| ID | Solução | Tipo | Propósito | Complexidade |
|:---|:---|:---|:---|:---:|
| S01 | ms-fbso-platform-admin | Backend API | Lógica de negócio, CRUD, RBAC, métricas (Java/Spring) | Alta |
| S02 | web-app-fbso-platform-portal | Frontend Web | Portal Admin + Portal Cliente (Next.js/React) | Média |
| S03 | PostgreSQL 17 | Banco de Dados | Persistência multi-tenant (DigitalOcean Managed) | Média |
| S04 | Keycloak 26 | IAM | Autenticação OIDC, RBAC, realms | Média |
| S05 | Kong API Gateway | API Gateway | Roteamento, autenticação delegada, header injection | Média |
| S06 | Cloudflare | CDN/WAF | Proxy reverso, DDoS, cache, SSL | Baixa |
| S07 | Prometheus + Grafana | Observabilidade | Métricas e dashboards | Média |
| S08 | Grafana Loki | Logs | Agregação de logs | Baixa |
| S09 | Jaeger + OpenTelemetry | Tracing | Rastreamento distribuído | Média |
| S10 | Elastic Stack | Logs avançados | Análise e troubleshooting | Média |
| S11 | Terraform + Ansible | IaC | Provisionamento e configuração | Média |
| S12 | Istio | Service Mesh | Tráfego, mTLS, observabilidade na malha | Alta |
| S13 | Keda | Autoscaling | Escala de pods orientada a eventos | Média |
| S14 | Karpenter | Cluster AS | Escala de nós automaticamente | Média |
| S15 | GitHub Actions | CI/CD | Build, test, SAST, deploy | Média |

**Total:** 15 soluções · Complexidade: 2 Alta · 11 Média · 2 Baixa
🤖 *F8 — Upstream Architecture Discovery*
