# PROJECT-TECHNICAL-DEFINITIONS-INFRA-CLOUD-DEFINITION — Definição de Infraestrutura Cloud

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Fase:** F12 — Bloco B · **Disciplina:** Infra/Cloud Specialist
- **Versão:** 1.0 · **Data:** 30/07/2026 · **Status:** CREATED

---

## 1. Topologia
- **Provedor:** AWS (us-east-1)
- **Ambientes:** Dev, Staging, Prod (contas AWS separadas)
- **VPC:** CIDR 10.0.0.0/16, subnets públicas (DMZ) + privadas (apps, dados)

## 2. Compute
| Serviço | Tecnologia | Dimensionamento |
|:---|:---|:---|
| API Backend | EKS (t3.medium × 3 nodes) | Auto-scaling 3-10 nodes |
| Frontend | CloudFront + S3 | Serverless |
| Background jobs | ECS Fargate | On-demand |

## 3. Networking
- **DNS:** Route53
- **CDN:** CloudFront (cache estático + API com TTL curto)
- **API Gateway:** Kong (self-hosted no EKS) ou AWS API Gateway
- **Load Balancer:** AWS ALB (interno para backend, público para API Gateway)

## 4. Storage
- **Block:** EBS gp3 para EKS persistent volumes
- **Object:** S3 (frontend static, data lake, backups)
- **File:** EFS para shared storage entre pods (quando necessário)

## 5. Disaster Recovery
- **RPO:** 1 hora (RDS automated backup + WAL archiving para S3)
- **RTO:** 4 horas (Multi-AZ RDS, EKS node groups pré-configurados, infra como código)
- **Estratégia:** Pilot Light — recursos mínimos na região secundária (us-east-2), escalar sob demanda

## 6. Custos Estimados (Produção)
| Serviço | Custo Mensal Est. |
|:---|:---|
| RDS PostgreSQL (db.t3.medium, Multi-AZ) | ~$150 |
| ElastiCache Redis (cache.t3.micro) | ~$30 |
| EKS (3× t3.medium) | ~$250 |
| CloudFront + S3 | ~$50 |
| Outros (MQ, Route53, WAF) | ~$70 |
| **Total** | **~$550/mês** |

🤖 *F12 — INFRA-CLOUD-DEFINITION · Roadmap v5.0*
