# DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md
## Fase 7 — Bloco B: Architecture & Security & Specialists (Discovery-Level)

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Análise de Viabilidade) |
| **Data** | 02 de agosto de 2026 |
| **Autor** | Infra/Cloud Specialist |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em 02/08/2026 |

**Documentos Vinculados:**
- [`DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md`](DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md) — Definição de Arquitetura (F2)
- [`DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md`](DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md) — DevOps/SRE (F5)
- [`STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`](../../../.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md)

---

## 1. Provedor Cloud e Topologia

### 1.1 Provedor: DigitalOcean (Exclusivo)

| Critério | Decisão |
|----------|---------|
| **Provedor Primário** | DigitalOcean — padrão corporativo FBSO |
| **Região Primária** | São Paulo (saopaulo-1) — latência < 30ms para público brasileiro |
| **Região de DR** | Nova York (nyc3) — disaster recovery cross-region |
| **Edge/CDN/WAF** | Cloudflare — toda entrada de tráfego passa pela Cloudflare antes do DO |

### 1.2 Topologia de Rede

```
                          INTERNET
                             │
                    ┌────────┴────────┐
                    │   Cloudflare     │
                    │ DNS • CDN • WAF  │
                    │ DDoS • SSL Term  │
                    └────────┬────────┘
                             │ HTTPS (Full Strict)
                    ┌────────┴────────┐
                    │  DO Load         │
                    │  Balancer        │
                    │  (Regional)      │
                    └────────┬────────┘
                             │
              ┌──────────────┴──────────────┐
              │      DOKS Cluster            │
              │      (saopaulo-1)            │
              │                              │
              │  ┌────────────────────────┐  │
              │  │  Node Pool: Standard    │  │
              │  │  2-3× (8 vCPU, 16 GB)  │  │
              │  │                         │  │
              │  │  ┌───────────────────┐  │  │
              │  │  │ Kong (Ingress)     │  │  │
              │  │  │ Istio (Service Mes│  │  │
              │  │  │ Backend Pods       │  │  │
              │  │  │ Observability      │  │  │
              │  │  └───────────────────┘  │  │
              │  └────────────────────────┘  │
              └──────────────────────────────┘
                             │
              ┌──────────────┼──────────────┐
              │               │              │
     ┌────────┴──────┐ ┌─────┴──────┐ ┌─────┴──────────┐
     │ PostgreSQL 17  │ │  Redis     │ │ Spaces (S3)     │
     │ (DO Managed)   │ │ (Managed)  │ │ (Object Store)  │
     │ saopaulo-1     │ │ saopaulo-1│ │ saopaulo-1      │
     └───────────────┘ └───────────┘ └────────────────┘
```

### 1.3 Fluxo de Acesso Completo

```
Usuário (Browser)
  │
  ├─ 1. DNS resolve → Cloudflare (Anycast)
  ├─ 2. Cloudflare WAF inspeciona → aplica regras OWASP, DDoS, Bot
  ├─ 3. Cloudflare → DO Load Balancer (HTTPS, Full Strict SSL)
  ├─ 4. DO LB → Kong Gateway (NodePort)
  ├─ 5. Kong:
  │     ├─ Valida JWT via Keycloak (Service-ID/Token-ID)
  │     ├─ Aplica rate limiting (por tenant, IP, endpoint)
  │     ├─ Injeta headers: X-Tenant-Id, X-User-Id, X-Roles
  │     └─ Encaminha → Istio Ingress Gateway
  ├─ 6. Istio:
  │     ├─ mTLS entre Kong → Backend
  │     └─ AuthorizationPolicy: allow Kong → backend
  ├─ 7. Backend (ms-fbso-platform-admin):
  │     ├─ Configura tenant context na sessão PostgreSQL
  │     ├─ Processa request com RBAC
  │     └─ Responde
  └─ 8. Resposta volta pelo mesmo caminho (backend → Istio → Kong → LB → Cloudflare → usuário)
```

---

## 2. Recursos Cloud — Estimativa de Custos

### 2.1 DigitalOcean — Custos Mensais Estimados

| Recurso | Especificação | Dev | Staging | Prod | Custo Mensal (Prod) |
|---------|--------------|-----|---------|------|---------------------|
| **DOKS Nodes** | 3× (8 vCPU, 16 GB) | 1× (4/8) | 2× (8/16) | 3× (8/16) | $288 (3 × $96) |
| **PostgreSQL** | 4 vCPU, 8 GB, HA | 1 vCPU, 2 GB | 2 vCPU, 4 GB | 4 vCPU, 8 GB | $240 |
| **Redis** | 2 vCPU, 4 GB, HA | 1 vCPU, 1 GB | 1 vCPU, 2 GB | 2 vCPU, 4 GB | $120 |
| **Load Balancer** | Regional | 1× | 1× | 1× | $12 |
| **Spaces (S3)** | 250 GB + CDN | 50 GB | 100 GB | 250 GB | $12.50 |
| **Container Registry** | 50 GB | — | — | 50 GB | $0 (incluso) |
| **Bandwidth** | Estimado 500 GB/mês | 50 GB | 100 GB | 500 GB | $0 (incluso no DOKS) |
| **Snapshot Backups** | PostgreSQL + Spaces | — | — | Diário | $0 (incluso) |
| **Total Mensal Estimado (Prod)** | | | | | **~$673/mês** |

### 2.2 Cloudflare — Custos Mensais

| Plano | Features | Custo Mensal |
|-------|----------|-------------|
| **Pro** | WAF (OWASP + custom rules), Bot Management, DDoS avançado, SSL dedicado, Image Optimization | $20/mês |

### 2.3 Custo Total Estimado

| Ambiente | DigitalOcean | Cloudflare | Total |
|----------|-------------|-----------|-------|
| **Dev** | ~$80/mês | — | ~$80/mês |
| **Staging** | ~$200/mês | — | ~$200/mês |
| **Prod** | ~$673/mês | $20/mês | **~$693/mês** |
| **Total 3 ambientes** | | | **~$973/mês** |

> **Nota:** Estimativas Discovery-Level (±50%). Custos reais dependem de configuração final, uso de banda, e negociação com provedores.

---

## 3. Disaster Recovery e Continuidade

### 3.1 Estratégia de DR

| Cenário | RPO | RTO | Estratégia |
|---------|-----|-----|-----------|
| **Falha de node K8s** | 0 (zero data loss) | < 5 min | Karpenter recria node automaticamente; pods redistribuídos |
| **Falha de PostgreSQL primário** | < 1 min | < 5 min | DO Managed failover automático para replica standby |
| **Falha de Redis primário** | < 1 min | < 5 min | DO Managed failover automático para replica |
| **Falha de região (saopaulo-1)** | < 24h | < 4h | DR cross-region (nyc3) — PostgreSQL backup replicado + Terraform recreate |
| **Desastre total (perda de região)** | < 24h | < 8h | Restore de backup diário em nova região via Terraform |

### 3.2 Backup Strategy

| Recurso | Frequência | Retenção | Local |
|---------|-----------|----------|-------|
| **PostgreSQL** | Diário (automático DO) | 7 dias | DO Managed |
| **PostgreSQL PITR** | Contínuo (WAL) | 7 dias | DO Managed |
| **Spaces (documentos)** | Diário | 30 dias | Cross-region (nyc3) |
| **K8s Config (Helm releases)** | Por deploy | Indefinido | Git (GitHub) |
| **Terraform State** | Por apply | Indefinido | DO Spaces + versionamento |

---

## 4. Segurança de Infraestrutura

| Controle | Implementação |
|----------|---------------|
| **Network Isolation** | DO VPC — recursos gerenciados (DB, Redis) na mesma VPC do DOKS |
| **Firewall** | DO Cloud Firewall + Cloudflare WAF + Kong rate limiting — 3 camadas |
| **TLS in Transit** | TLS 1.3 em toda comunicação externa; mTLS Istio intra-cluster |
| **Encryption at Rest** | Ativado em PostgreSQL, Redis e Spaces (DO Managed) |
| **Secrets Management** | K8s Secrets + Vault (fase futura); `.env` nunca commitado |
| **SSH Access** | Desabilitado para nós DOKS (gerenciados); acesso administrativo via kubectl |
| **Audit Logging** | DO audit logs + Elastic Stack (audit trail de aplicação) |

---

## 5. Riscos e Estimativa de Esforço

### 5.1 Riscos de Infra/Cloud

| ID | Risco | Prob. | Impacto | Mitigação |
|----|-------|-------|---------|-----------|
| RI1 | DigitalOcean saopaulo-1 — indisponibilidade de região | Muito Baixa | 🔴 Crítico | DR cross-region (nyc3); Terraform idempotente permite recriar em outra região |
| RI2 | Custo de infra subestimado — estouro de orçamento | Baixa | 🟡 Médio | Monitoramento de custos DO; alertas de orçamento; Karpenter otimiza node usage |
| RI3 | DOKS limitações vs. AWS/Azure — necessidade futura de migração | Baixa | 🟡 Médio | Abstraction via Terraform + Helm facilita portabilidade; K8s é cloud-agnostic |

### 5.2 Estimativa de Esforço

| Atividade | Complexidade | Esforço (dias) | Responsável |
|-----------|-------------|----------------|-------------|
| Provisioning DOKS + DBs + Redis via Terraform | Complexa | 2 | Lucas Silva Neto |
| Configuração Cloudflare DNS + WAF + SSL | Moderada | 1 | Lucas Silva Neto |
| Rede e VPC | Leve | 0.5 | Lucas Silva Neto |
| Backup e DR strategy | Moderada | 0.5 | Lucas Silva Neto |
| Monitoramento de custos e alertas | Leve | 0.5 | Lucas Silva Neto |
| Documentação e estimativas | Leve | 0.5 | Lucas Silva Neto |
| **Total Infra/Cloud** | — | **~5 dias** | — |

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 02/08/2026 | Criação inicial: Infra/Cloud Definition. DO exclusivo, topologia rede, fluxo de acesso, estimativa de custos (~$973/mês total), DR (RPO/RTO), segurança de infra, 3 riscos, estimativa ~5 dias | Infra/Cloud Specialist |

---

🤖 *Upstream Architecture Discovery — Fase 7. Documento gerado pelo Infra/Cloud Specialist como parte do Bloco B.*
