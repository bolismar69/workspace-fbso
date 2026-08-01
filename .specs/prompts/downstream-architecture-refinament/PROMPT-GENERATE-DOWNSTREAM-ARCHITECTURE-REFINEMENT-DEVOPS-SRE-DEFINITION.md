# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — DEVOPS-SRE-DEFINITION (F5)
## Versão: 1.1 — DevOps/SRE Detail-Level (Pipeline Specs + IaC + SLOs) — Específico do Projeto

Atue como um DevOps/SRE Architect especializado em infraestrutura cloud native com Kubernetes, Istio e GitOps.

> ⚠️ **Nota:** As tecnologias listadas neste prompt (DigitalOcean, Kong, Istio, GitHub Actions, DOKS) são específicas do projeto atual. Para executar este prompt em outro projeto, ajustar conforme a stack tecnológica do projeto alvo.

## OBJETIVO

Produzir a definição de DevOps/SRE em nível de implementação: pipeline specs, IaC templates, observabilidade stack, SLOs com SLIs, estratégia de deploy.

## INPUTS

1. **Arquitetura Detail-Level** (F2)
2. **Segurança Detail-Level** (F3)
3. **DevOps/SRE Discovery-Level** (se existir)
4. **INFRA-CLOUD Discovery-Level** (se existir)

## ESTRUTURA DO DOCUMENTO

```markdown
# DETAIL-LEVEL-DEVOPS-SRE-DEFINITION — DevOps/SRE Detail-Level

## 1. Pipeline Specs (GitHub Actions)
[Workflows por ambiente: dev / staging / prod]
- pr-checks.yml (SAST + Secret Scanning)
- build-deploy-dev.yml
- deploy-staging.yml
- deploy-prod.yml (com approval gate)

## 2. IaC (Terraform + Ansible)
[Terraform: DOKS cluster, PostgreSQL, Redis, networking]
[Ansible: node provisioning, Kong config, monitoring agents]

## 3. Observabilidade Stack
- Prometheus: alert rules (latência, erro, saturação)
- Grafana: dashboards (Backend, Frontend, DB, Kong)
- Loki: log queries por tenant, error patterns
- Jaeger: sampling strategy, trace propagation (OpenTelemetry)
- Elastic Stack: complementar ao Loki (logs de auditoria)

## 4. SLOs com SLIs
| Serviço | SLO | SLI | Janela |
| Backend API | 99.9% | latência p99 < 500ms | 30d |
| Frontend | 99.5% | LCP < 2.5s | 30d |
| PostgreSQL | 99.95% | latência p99 < 100ms | 30d |

## 5. Estratégia de Deploy
- Blue-Green para serviços stateless
- Canary para mudanças críticas (RBAC, autenticação)
- Rollback automático se SLO violado

## 6. Runbooks
[Procedimentos: tenant isolation breach, DB failover, Kong misconfiguration]

## 7. Riscos DevOps
```

### Skills Recomendados
- `senior-devops`, `kubernetes-specialist`, `terraform-engineer`
- `observability-engineer`, `monitoring-expert`
- `github-actions-advanced`, `sre-engineer`

🤖 *Prompt gerador — Fase 5 do Downstream Architecture Refinement*
