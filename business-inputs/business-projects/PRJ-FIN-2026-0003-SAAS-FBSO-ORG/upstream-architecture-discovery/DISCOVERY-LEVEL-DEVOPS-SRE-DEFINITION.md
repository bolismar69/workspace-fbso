# DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION — Definição DevOps/SRE (Discovery)
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F5 — Bloco B · **Disciplina:** DevOps/SRE Architect
- **Versão:** 1.1 · **Data:** 30/07/2026 · **Status:** CREATED

## 1. Estratégia de Deploy
| Ambiente | Infra | Deploy |
|:---|:---|:---|
| Dev | Docker Compose (local) | Automático (push main) |
| Staging | DigitalOcean DOKS | Manual (tag) |
| Produção | DigitalOcean DOKS | Manual (release) |

## 2. CI/CD (Macro)
| Etapa | Ferramenta | Tempo Est. |
|:---|:---|:---:|
| Build + Test | GitHub Actions (padrão corporativo) | ~5 min |
| SAST | Semgrep | ~2 min |
| Container Build | Docker | ~3 min |
| Deploy K8s | kubectl + Helm | ~2 min |

## 3. Observabilidade (Stack Corporativa)

| Pilar | Ferramenta | Propósito |
|:---|:---|:---|
| **Métricas** | **Prometheus + Grafana** (padrão corporativo) | Coleta, armazenamento e dashboards de métricas |
| **Logs** | **Grafana Loki** (padrão corporativo) | Agregação e consulta de logs estruturados |
| **Logs (complementar)** | **Elastic Stack** (ELK) | Análise avançada de logs, troubleshooting |
| **Tracing** | **Jaeger + OpenTelemetry** (padrão corporativo) | Rastreamento distribuído entre microserviços |
| **Instrumentação** | OpenTelemetry SDK (Java/Spring, Node.js) | Coleta automática de traces, métricas e logs |

## 4. Infra as Code (IaC)
| Ferramenta | Propósito |
|:---|:---|
| **Terraform** (padrão corporativo) | Provisionamento de infraestrutura DigitalOcean (DOKS, PostgreSQL, Redis, Load Balancers) |
| **Ansible** (padrão corporativo) | Configuração de servidores, hardening, instalação de agentes |

## 5. Orquestração e Escalabilidade
| Ferramenta | Propósito |
|:---|:---|
| **Kubernetes (DOKS)** | Orquestração de containers |
| **Istio** (Service Mesh) | Gerenciamento de tráfego, mTLS entre serviços, observabilidade na malha |
| **Keda** (Event-Driven Autoscaling) | Escala automática de pods baseada em eventos (filas, métricas, triggers) |
| **Karpenter** (Cluster Autoscaling) | Adiciona/remove nós automaticamente conforme demanda de recursos |

## 6. SLOs Alvo
| Serviço | Disponibilidade | Latência p99 |
|:---|:---:|:---:|
| Backend API | 99.9% | < 500ms |
| Frontend Portal | 99.5% | LCP < 2.5s |
| Banco de Dados | 99.95% | < 100ms |

## 7. Estimativa de Esforço
| Atividade | Esforço |
|:---|:---:|
| Pipeline CI/CD (GitHub Actions) | 1-2 homem-mês |
| IaC (Terraform + Ansible) | 1-2 homem-mês |
| Observabilidade (Prometheus, Loki, Jaeger, Grafana) | 1-2 homem-mês |
| K8s + Istio + Keda + Karpenter | 1-2 homem-mês |
| **Total DevOps/SRE** | **4-8 homem-mês** |

🤖 *Upstream Architecture Discovery — Fase 5*
