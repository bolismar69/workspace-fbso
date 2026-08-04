# DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md
## Fase 5 — Bloco B: Architecture & Security & Specialists (Discovery-Level)

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Análise de Viabilidade) |
| **Data** | 02 de agosto de 2026 |
| **Autor** | DevOps/SRE Architect |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em 02/08/2026 |

**Documentos Vinculados:**
- [`DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md`](DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md) — Definição de Arquitetura (F2)
- [`DISCOVERY-LEVEL-SECURITY-DEFINITION.md`](DISCOVERY-LEVEL-SECURITY-DEFINITION.md) — Definição de Segurança (F3)
- [`STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`](../../../.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md)

---

## 1. Estratégia DevOps/SRE — Visão Macro

### 1.1 Pipeline CI/CD

```
┌──────────┐    ┌──────────────┐    ┌──────────────┐    ┌──────────────┐    ┌──────────┐
│  Commit   │───▶│  Build &     │───▶│  Docker       │───▶│  Deploy       │───▶│  Verify   │
│  (Push)   │    │  Test        │    │  Build & Push │    │  to K8s       │    │  Health   │
└──────────┘    └──────────────┘    └──────────────┘    └──────────────┘    └──────────┘
                     │                    │                    │                  │
                     ▼                    ▼                    ▼                  ▼
               ┌──────────┐       ┌──────────┐        ┌──────────┐       ┌──────────┐
               │ Semgrep   │       │ Docker    │        │ Helm      │       │ Prometheus│
               │ Gitleaks  │       │ Registry  │        │ kubectl   │       │ Health    │
               │ JUnit     │       │ (GHCR)    │        │ ArgoCD    │       │ Check     │
               │ npm test  │       │           │        │           │       │           │
               └──────────┘       └──────────┘        └──────────┘       └──────────┘
```

### 1.2 GitHub Actions Workflow

| Stage | Ação | Ferramenta | Gate |
|-------|------|-----------|------|
| **1. Checkout** | Clonar repositório | actions/checkout | — |
| **2. SAST** | Análise estática de segurança | Semgrep | ❌Block on CRITICAL |
| **3. Secrets Scan** | Detectar secrets hardcoded | Gitleaks | ❌Block on ANY finding |
| **4. Backend Build** | Compilar + testes unitários | Maven + JUnit 5 | ❌Block on test failure |
| **5. Frontend Build** | Compilar + testes | npm + Jest | ❌Block on test failure |
| **6. Docker Build** | Build imagem GraalVM Native | Docker +GraalVM | — |
| **7. Docker Push** | Push para GHCR | docker push | — |
| **8. Deploy Dev** | Deploy automático em dev | Helm + kubectl | ✅Auto |
| **9. Smoke Tests** | Testes de sanidade pós-deploy | Playwright (API) | ❌Block on failure |
| **10. Deploy Staging** | Deploy manual em staging | Helm + kubectl | ⏸️Manual approval |
| **11. Deploy Prod** | Deploy manual em produção | Helm + kubectl + ArgoCD | ⏸️Manual approval |

### 1.3 Ambientes

| Ambiente | Propósito | Deploy | Recursos |
|----------|-----------|--------|----------|
| **Dev** | Desenvolvimento contínuo | Automático (push na branch) | 1 node DOKS reduzido |
| **Staging** | Validação pré-produção | Manual (após smoke tests passarem) | 2 nodes, equivalente a Prod |
| **Prod** | Produção | Manual (com approval gate) | 2-3 nodes, HA |

---

## 2. Infraestrutura como Código (IaC)

### 2.1 Terraform — Provisioning

| Recurso | Provider | Configuração |
|---------|----------|-------------|
| **DOKS Cluster** | `digitalocean_kubernetes_cluster` | 2-3 nodes, 8 vCPU/16 GB, auto-upgrade |
| **PostgreSQL** | `digitalocean_database_cluster` | pg-17, 4 vCPU/8 GB, HA, backup diário |
| **Redis** | `digitalocean_database_cluster` | redis-7, 2 vCPU/4 GB, HA |
| **Spaces** | `digitalocean_spaces_bucket` | 250 GB, CDN enabled |
| **Load Balancer** | `digitalocean_loadbalancer` | Regional, SSL termination |
| **Container Registry** | `digitalocean_container_registry` | Imagens Docker privadas |
| **Cloudflare DNS** | `cloudflare_record` | DNS records apontando para DO LB |

### 2.2 Ansible — Configuração

| Playbook | Alvo | Configura |
|----------|------|-----------|
| **kong-setup** | Kong Gateway node | Plugins, rotas, rate limiting, CORS |
| **monitoring-agents** | Todos os nodes | Prometheus node exporter, Loki agent, OTel collector |
| **hardening** | Todos os nodes | CIS benchmarks, firewalls, SSH hardening |

---

## 3. Kubernetes — Orquestração e Service Mesh

### 3.1 Topologia de Deploy K8s

```yaml
# Estrutura de namespaces
namespaces:
  - kong-gateway       # Kong Gateway (Ingress)
  - istio-system       # Istio control plane
  - backend            # ms-fbso-platform-admin
  - observability      # Prometheus, Grafana, Loki, Jaeger
  - elastic            # Elastic Stack (audit logs)
  - cert-manager       # TLS certificate management
```

### 3.2 Istio — Service Mesh

| Feature | Configuração |
|---------|-------------|
| **mTLS** | STRICT mode em produção (todos os serviços) |
| **AuthorizationPolicy** | Allow backend→postgres, backend→redis; deny all else |
| **Traffic Control** | Canary (10% → 50% → 100%), Blue-Green, Circuit Breaking |
| **Observabilidade** | Sidecar (Envoy) injetado automaticamente — métricas, traces, logs |

### 3.3 Autoscaling

| Componente | Ferramenta | Trigger |
|------------|-----------|---------|
| **Pods** | Keda | CPU > 70%, Memory > 80%, HTTP requests pendentes |
| **Nodes** | Karpenter | Pods pendentes por falta de recursos; otimização de custo |

---

## 4. Observabilidade — Stack Completa

### 4.1 Stack Overview

```
┌────────────────────────────────────────────────────────────┐
│                   APLICAÇÃO (Spring Boot)                   │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐  │
│  │Micrometer│  │  OTel    │  │  Logback │  │ Audit    │  │
│  │(Métricas)│  │  Agent   │  │  Appender│  │ Logger   │  │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘  └────┬─────┘  │
└───────┼──────────────┼─────────────┼─────────────┼────────┘
        │              │             │             │
        ▼              ▼             ▼             ▼
   ┌─────────┐   ┌─────────┐   ┌─────────┐   ┌──────────┐
   │Prometheus│  │ Jaeger  │   │  Loki   │   │ Elastic  │
   │(Scrape) │   │(gRPC)   │   │(HTTP)   │   │ Stack    │
   └────┬────┘   └────┬────┘   └────┬────┘   └────┬─────┘
        │             │             │             │
        └─────────────┴─────────────┴─────────────┘
                      │
                      ▼
               ┌───────────┐
               │  Grafana   │
               │ Dashboards │
               └───────────┘
```

### 4.2 Métricas (Prometheus + Micrometer)

| Métrica | Tipo | Descrição |
|---------|------|-----------|
| `http_server_requests_seconds` | Histogram | Latência de endpoints REST |
| `tenant_active_count` | Gauge | Tenants ativos |
| `subscription_active_count` | Gauge | Assinaturas ativas |
| `audit_actions_total` | Counter | Ações de auditoria registradas |
| `db_connection_pool_active` | Gauge | Conexões ativas no pool |
| `jvm_memory_used_bytes` | Gauge | Memória JVM utilizada |

### 4.3 Alertas Críticos

| Alerta | Condição | Severidade | Canal |
|--------|----------|-----------|-------|
| **API Error Rate > 5%** | `rate(http_server_requests_seconds_count{status=~"5.."}[5m]) > 0.05` | 🔴 Crítico | PagerDuty |
| **Latência P99 > 2s** | `histogram_quantile(0.99, http_server_requests_seconds) > 2` | 🟠 Warning | Slack #alerts |
| **DB Connection Pool > 80%** | `db_connection_pool_active / db_connection_pool_max > 0.8` | 🟠 Warning | Slack #alerts |
| **Tenant Suspenso com acesso** | Custom check: `suspended_tenants_with_access > 0` | 🔴 Crítico | PagerDuty |
| **Certificado TLS expirando** | `cert_expiry_days < 30` | 🟡 Info | Slack #alerts |

### 4.4 Dashboards Grafana

| Dashboard | Público | Métricas |
|-----------|---------|----------|
| **Platform Overview** | Diretoria | Tenants ativos, novas contas/dia, assinaturas por plano |
| **API Performance** | Tech Lead | Latência P50/P95/P99, error rate, throughput |
| **Infrastructure Health** | DevOps | CPU, memória, disco, conexões DB, Redis hit rate |
| **Security & Audit** | IAM Specialist | Alertas de segurança, ações de auditoria anômalas |
| **Business Metrics** | Comercial | Onboarding completion rate, plan distribution, churn signals |

---

## 5. SLOs — Service Level Objectives

| SLO | Target | Measurement Window |
|-----|--------|-------------------|
| **API Availability** | 99.9% | 30 dias |
| **API Latency P95** | < 500ms | 30 dias |
| **API Latency P99** | < 2s | 30 dias |
| **Error Rate** | < 1% | 30 dias |
| **Deploy Frequency** | ≥ 1/semana | 30 dias |
| **Deploy Lead Time** | < 1 hora (do merge ao deploy em dev) | — |
| **Mean Time to Recovery (MTTR)** | < 30 min | — |

---

## 6. Riscos e Estimativa de Esforço

### 6.1 Riscos DevOps/SRE

| ID | Risco | Prob. | Impacto | Mitigação |
|----|-------|-------|---------|-----------|
| RDV1 | Istio + Keda + Karpenter — complexidade excessiva para fase inicial | Média | 🟡 Médio | Iniciar com Istio mínimo (apenas mTLS); Keda/Karpenter progressivamente |
| RDV2 | GraalVM Native Image — build lento e complexo | Média | 🟡 Médio | Cache de build layers; fallback JVM HotSpot no CI se build > 10 min |
| RDV3 | DigitalOcean — recursos limitados vs. AWS/Azure | Baixa | 🟡 Médio | DO atende ao porte previsto; migração cloud é possível mas custosa |

### 6.2 Estimativa de Esforço

| Atividade | Complexidade | Esforço (dias) | Responsável |
|-----------|-------------|----------------|-------------|
| Pipeline CI/CD (GitHub Actions) | Moderada | 1.5 | Lucas Silva Neto |
| Terraform (DOKS, DBs, Spaces) | Complexa | 2 | Lucas Silva Neto |
| Ansible (Kong, monitoring agents, hardening) | Moderada | 1 | Lucas Silva Neto |
| K8s namespaces + Helm charts | Moderada | 1.5 | Lucas Silva Neto |
| Istio setup (mTLS + AuthorizationPolicy) | Complexa | 1.5 | Lucas Silva Neto |
| Observabilidade (Prometheus + Loki + Jaeger + Grafana) | Moderada | 2 | Lucas Silva Neto |
| Elastic Stack (audit logs) | Moderada | 1 | Lucas Silva Neto |
| Alertas e dashboards | Leve | 1 | Lucas Silva Neto |
| SLOs e documentação | Leve | 0.5 | Lucas Silva Neto |
| **Total DevOps/SRE** | — | **~12 dias** | — |

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 02/08/2026 | Criação inicial: DevOps/SRE Definition. Pipeline CI/CD (10 stages), IaC (Terraform+Ansible), K8s+Istio+Keda+Karpenter, observabilidade completa, 5 dashboards, 7 SLOs, estimativa ~12 dias | DevOps/SRE Architect |

---

🤖 *Upstream Architecture Discovery — Fase 5. Documento gerado pelo DevOps/SRE Architect como parte do Bloco B.*
