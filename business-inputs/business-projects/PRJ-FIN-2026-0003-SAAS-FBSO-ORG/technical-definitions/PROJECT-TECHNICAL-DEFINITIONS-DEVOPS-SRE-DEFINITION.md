# PROJECT-TECHNICAL-DEFINITIONS-DEVOPS-SRE-DEFINITION — Definição DevOps/SRE

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Fase:** F10 — Bloco B · **Disciplina:** DevOps/SRE Architect
- **Versão:** 1.0 · **Data:** 30/07/2026 · **Status:** CREATED

---

## 1. Pipeline CI/CD
- **Ferramenta:** GitHub Actions
- **Workflows:** build→test→SAST→container→deploy (Dev/Staging/Prod)
- **Estratégia de branch:** feature/* → PR → main → deploy automático Dev; tag semver → Staging; release → Prod

## 2. Infrastructure as Code
- **Ferramenta:** Terraform (AWS) + Docker Compose (local)
- **Módulos:** RDS, ElastiCache, EKS, S3, IAM, Route53, CloudFront
- **State:** S3 backend com DynamoDB lock

## 3. Observabilidade
- **Logging:** SLF4J/Logback → CloudWatch Logs, formato JSON estruturado
- **Metrics:** Micrometer → Prometheus → Grafana
- **Tracing:** OpenTelemetry → AWS X-Ray
- **Alerting:** CloudWatch Alarms → PagerDuty

## 4. SLOs/SLIs
| Serviço | SLO | SLI |
|:---|:---|:---|
| API Backend | 99.9% availability | Latência p99 < 500ms |
| Portal Frontend | 99.5% availability | LCP < 2.5s |
| Banco de Dados | 99.95% availability | Replication lag < 1s |

## 5. Containers & Orquestração
- **Container:** Docker, imagens multi-stage (Java 21 + Node 22)
- **Orquestração:** AWS EKS (Kubernetes) para produção; Docker Compose para dev local
- **Registry:** AWS ECR

## 6. Ambientes
| Ambiente | Propósito | Deploy |
|:---|:---|:---|
| Dev | Desenvolvimento contínuo | Auto (merge main) |
| Staging | Validação pré-produção | Manual (tag) |
| Prod | Produção | Manual (release) |

## 7. Runbooks
- **Incidente crítico (P1):** PagerDuty → Tech Lead + DevOps on-call → diagnóstico → rollback ou hotfix
- **Postmortem:** Documento no Confluence em até 48h após resolução

🤖 *F10 — DEVOPS-SRE-DEFINITION · Roadmap v5.0*
