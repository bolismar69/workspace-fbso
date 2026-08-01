# DETAIL-LEVEL-DEVOPS-SRE-DEFINITION — DevOps/SRE Detail-Level

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Data:** 31/07/2026
- **Fase:** F5 — Downstream Architecture Refinement
- **Padrões FBSO:** Prometheus, Grafana Loki, Jaeger, OpenTelemetry, Grafana, Elastic Stack, Terraform, Ansible, Kubernetes/DOKS, Istio, Keda, Karpenter, GitHub Actions

---

## 1. Pipeline Specs (GitHub Actions)

### 1.1 Workflows

| Workflow | Trigger | Ações |
|:---|:---|:---|
| `pr-checks.yml` | Pull Request | SAST (Semgrep), Secret Scanning (Gitleaks), Build (Maven), Unit Tests (JUnit), Lint |
| `build-deploy-dev.yml` | Push to `develop` | Docker build (GraalVM Native + JVM fallback), Push to registry, Deploy to Dev DOKS |
| `deploy-staging.yml` | PR merged to `staging` | Deploy to Staging DOKS, Integration tests (Testcontainers), Smoke tests (k6) |
| `deploy-prod.yml` | Release tag `v*` | Manual approval gate → Deploy to Prod DOKS (blue-green), Post-deploy health check |
| `nightly.yml` | Cron (daily 02:00) | Full test suite, E2E (Playwright), Performance (k6), Security scan (OWASP ZAP) |

### 1.2 Quality Gates por Ambiente

| Gate | Dev | Staging | Prod |
|:---|:---:|:---:|:---:|
| Unit tests pass | ✅ ≥80% | ✅ | ✅ |
| SAST limpo (Semgrep) | ✅ 0 críticas | ✅ | ✅ |
| Secret scan limpo (Gitleaks) | ✅ | ✅ | ✅ |
| Integration tests pass | — | ✅ ≥60% | ✅ |
| E2E críticos passam | — | ✅ | ✅ |
| Performance smoke (k6) | — | ✅ p99<500ms | ✅ |
| Security scan (OWASP ZAP) | — | ✅ 0 críticas | ✅ |
| Manual approval | — | — | ✅ Required |

---

## 2. IaC (Terraform + Ansible)

### 2.1 Terraform — Recursos

```hcl
# digitalocean-kubernetes/main.tf
resource "digitalocean_kubernetes_cluster" "fbso_prod" {
  name    = "fbso-prod"
  region  = "nyc3"
  version = "1.30"
  node_pool {
    name       = "default-pool"
    size       = "s-4vcpu-8gb"
    auto_scale = true
    min_nodes  = 3
    max_nodes  = 10
  }
}

resource "digitalocean_database_cluster" "postgres" {
  name       = "fbso-postgres-prod"
  engine     = "pg"
  version    = "17"
  size       = "db-s-4vcpu-8gb"
  region     = "nyc3"
  node_count = 2
}

resource "digitalocean_redis_cluster" "redis" {
  name       = "fbso-redis-prod"
  engine     = "redis"
  version    = "7"
  size       = "db-s-2vcpu-4gb"
  region     = "nyc3"
  node_count = 1
}
```

### 2.2 Ansible — Provisioning

| Playbook | Função |
|:---|:---|
| `kong-setup.yml` | Instala e configura Kong Gateway, plugins OIDC, rate limiting, logging |
| `monitoring-agents.yml` | Instala Prometheus exporters (node, PostgreSQL, Kong), OTel Collector |
| `istio-setup.yml` | Configura Istio sidecar injection, mTLS, VirtualServices, DestinationRules |
| `keda-setup.yml` | Instala Keda via Helm, configura ScaledObjects para backend |

---

## 3. Observabilidade Stack

### 3.1 Stack Completa

| Camada | Ferramenta | O que coleta | Destino |
|:---|:---|:---|:---|
| Métricas | **Prometheus** | Kong, PostgreSQL, Keycloak, Backend (Micrometer), DOKS nodes | Grafana |
| Logs | **Grafana Loki** | Logs de aplicação Spring Boot, Kong access logs | Grafana |
| Tracing | **Jaeger** | Spans via OpenTelemetry Collector | Jaeger UI |
| Instrumentação | **OpenTelemetry** | Java agent (auto), spans manuais em pontos críticos | OTel Collector |
| Dashboards | **Grafana** | Unificado: Prometheus + Loki + PostgreSQL datasources | — |
| Auditoria | **Elastic Stack** | Logs de auditoria (audit_log) — busca full-text, retenção longa | Kibana |

### 3.2 Alert Rules (Prometheus)

| Alerta | Condição | Severidade | Canal |
|:---|:---|:---|:---|
| Backend API Down | `up{job="backend"} == 0` > 2min | Crítico | PagerDuty |
| Backend p99 > 1s | `histogram_quantile(0.99, http_request_duration_seconds) > 1` por 5min | Warning | Slack |
| DB Connection Pool > 80% | `hikaricp_connections_active / hikaricp_connections_max > 0.8` por 5min | Warning | Slack |
| Kong 5xx Rate > 1% | `rate(kong_http_status{code="5xx"}[5m]) > 0.01` | Warning | Slack |
| Tenant RLS Failure | `rate(rls_policy_violations_total[5m]) > 0` | Crítico | PagerDuty |
| Redis Memory > 80% | `redis_memory_used_bytes / redis_memory_max_bytes > 0.8` | Warning | Slack |

---

## 4. SLOs com SLIs

| Serviço | SLO | SLI | Janela |
|:---|:---|:---|:---|
| Backend API | 99.9% disponibilidade | `rate(http_requests_total{status!~"5.."}[30d]) / rate(http_requests_total[30d])` | 30 dias |
| Backend API | p99 < 500ms | `histogram_quantile(0.99, http_request_duration_seconds)` | 30 dias |
| Frontend Portal | 99.5% disponibilidade | Cloudflare analytics: success rate | 30 dias |
| Frontend Portal | LCP < 2.5s | Web Vitals (p75) | 30 dias |
| PostgreSQL | 99.95% disponibilidade | DigitalOcean managed: uptime | 30 dias |
| PostgreSQL | p99 < 100ms | `pg_stat_statements` metrics | 30 dias |
| Keycloak | 99.9% disponibilidade | Health endpoint | 30 dias |
| Kong Gateway | 99.9% disponibilidade | Health endpoint + request success rate | 30 dias |

---

## 5. Estratégia de Deploy

| Estratégia | Quando Usar | Rollback |
|:---|:---|:---|
| **Blue-Green** (padrão) | Deploys regulares de backend e frontend | Switch instantâneo via Istio VirtualService |
| **Canary** | Mudanças críticas: RBAC, autenticação, multi-tenancy | Reduzir tráfego canary → 0% se anomalia detectada |
| **Rollback Automático** | Se SLO violado em até 5min pós-deploy | Prometheus alert → webhook → `helm rollback` |

### Fluxo de Deploy (Blue-Green)

```
1. Build + Test → Docker image tag v{version}
2. Deploy Green environment (novo ReplicaSet)
3. Health check Green (health endpoint + smoke tests)
4. Istio VirtualService: switch 100% tráfego Blue → Green
5. Monitor 5 minutos (Prometheus alerts)
6. Se OK: remove Blue. Se SLO violado: Istio switch back → Blue
```

---

## 6. Runbooks

| Cenário | Procedimento |
|:---|:---|
| **Tenant Isolation Breach** | 1. Verificar RLS policies no PostgreSQL. 2. Auditar logs de acesso cross-tenant. 3. Bloquear tenant ofensor via Kong rate limit. 4. Corrigir bug e aplicar migration de reparo |
| **DB Failover** | 1. DigitalOcean promove standby automaticamente. 2. Verificar connection string no backend → aponta para novo primary. 3. Reiniciar backend pods para refresh de conexões |
| **Kong Misconfiguration** | 1. Rollback Kong config via `deck sync` com última config versionada. 2. Validar que plugins OIDC e rate limiting voltaram. 3. Testar endpoint de health check |
| **Keycloak Outage** | 1. Backend usa cache de JWKS → tolera outage curta. 2. Kong health check detecta → redireciona para fallback. 3. Restaurar Keycloak via backup do PostgreSQL |

---

## 7. Riscos DevOps

| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---|:---|
| Falha de deploy em produção | Média | Alto | Blue-green + rollback automático + Istio traffic splitting |
| Observabilidade cega (métricas/logs/traces ausentes) | Baixa | Alto | Health check dos exporters; alerta se Prometheus sem scrape > 5min |
| IaC drift (config manual vs Terraform) | Média | Médio | Terraform state versionado; `terraform plan` semanal; drift detection |
| Keda/Karpenter escala insuficiente em pico | Baixa | Médio | Definir min/max replicas e nodes; teste de carga trimestral |

---

🤖 *Documento gerado pelo DevOps/SRE Architect — Fase 5 do Downstream Architecture Refinement. Padrões FBSO: Prometheus, Loki, Jaeger, OTel, Grafana, Elastic Stack, Terraform, Ansible, K8s, Istio, Keda, Karpenter, GitHub Actions.*
