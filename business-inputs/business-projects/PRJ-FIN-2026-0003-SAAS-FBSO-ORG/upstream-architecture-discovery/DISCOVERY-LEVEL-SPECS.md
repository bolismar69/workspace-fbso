# DISCOVERY-LEVEL-SPECS — Especificação Técnica Consolidada (Discovery)
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F10 — Bloco C · **Versão:** 1.0 · **Data:** 30/07/2026

## 1. Visão Técnica Consolidada

A FBSO Platform segue arquitetura de microserviços com API Gateway central (Kong) e Service Mesh (Istio). Autenticação delegada ao Kong↔Keycloak via Service-ID. Trust boundary: microserviços não revalidam JWT. Multi-tenant isolation via RLS no PostgreSQL.

## 2. Stack Tecnológica Corporativa

| Camada | Tecnologia |
|:---|:---|
| Cloud | DigitalOcean (DOKS, Managed PostgreSQL, Managed Redis, Spaces) |
| CDN/WAF | Cloudflare |
| API Gateway | Kong + OIDC Plugin |
| IAM | Keycloak 26 |
| Backend | Java 25 / Spring Boot 3.5 |
| Frontend | Next.js 15 / React 19 / Tailwind CSS |
| Banco | PostgreSQL 17 (Multi-Tenant RLS) |
| Cache | Redis |
| CI/CD | GitHub Actions |
| IaC | Terraform + Ansible |
| Orquestração | Kubernetes (DOKS) + Istio + Keda + Karpenter |
| Observabilidade | Prometheus + Grafana + Loki + Jaeger + OpenTelemetry + Elastic Stack |

## 3. Decisões Técnicas Transversais

| Decisão | Detalhe | Disciplina |
|:---|:---|:---|
| Multi-Tenant Isolation | Discriminator column + RLS (não schema-per-tenant) | DATA |
| Autenticação | Kong↔Keycloak Service-ID; microserviços não revalidam | SEC |
| API Gateway | Kong como único ponto de entrada público | ARCH |
| Service Mesh | Istio — mTLS, controle de tráfego, observabilidade | DEVOPS |
| Autoscaling | Keda (pods) + Karpenter (nós) | DEVOPS |
| Infra as Code | Terraform (provisionamento) + Ansible (configuração) | DEVOPS |

## 4. Estimativa ROM Consolidada (Bloco B)

| Disciplina | Esforço (homem-mês) |
|:---|:---:|
| Arquitetura (F2) | 7-10 |
| Segurança (F3) | 3-5 |
| Dados (F4) | 2-3.5 |
| DevOps/SRE (F5) | 4-8 |
| Testes (F6) | 2.5-4 |
| Infra/Cloud (F7) | 2.5-3.5 |
| **Total Bloco B** | **21.5-34** |

## 5. Artefatos Vinculados

| Fase | Artefato | Disciplina |
|:---|:---|:---|
| F1 | [DISCOVERY-LEVEL-PRD](./DISCOVERY-LEVEL-PRD.md) | Product Definition |
| F2 | [ARCHITECTURE](./DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md) | Solution Architect |
| F3 | [SECURITY](./DISCOVERY-LEVEL-SECURITY-DEFINITION.md) | Security Architect |
| F4 | [DATA](./DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md) | Data Architect |
| F5 | [DEVOPS-SRE](./DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md) | DevOps/SRE Architect |
| F6 | [TEST-STRATEGY](./DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md) | Test Specialist |
| F7 | [INFRA-CLOUD](./DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md) | Infra/Cloud Specialist |

🤖 *F10 — Upstream Architecture Discovery*
