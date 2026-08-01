# DETAIL-LEVEL-INFRA-CLOUD-DEFINITION — Infra/Cloud Detail-Level

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Data:** 31/07/2026
- **Fase:** F7 — Downstream Architecture Refinement
- **Padrões FBSO:** DigitalOcean, Cloudflare

---

## 1. Topologia de Rede

```
                          Internet
                             │
                     ┌───────▼────────┐
                     │  Cloudflare     │
                     │  DNS + WAF +    │
                     │  CDN + DDoS     │
                     └───────┬────────┘
                             │ HTTPS :443
                     ┌───────▼────────┐
                     │  Kong Gateway   │  ← DigitalOcean Droplet / DOKS pod
                     │  (Trust Bound)  │
                     └───┬─────────────┘
                         │ HTTP interno
              ┌──────────┼──────────────────┐
              │          │                   │
     ┌────────▼──┐ ┌─────▼──────┐  ┌────────▼──────┐
     │ Backend   │ │ Frontend   │  │ Keycloak IAM  │
     │ API (S01) │ │ Portal(S02)│  │ (S04)          │
     └─────┬─────┘ └───────────┘  └───────┬────────┘
           │                              │
     ┌─────▼──────┐              ┌────────▼──────┐
     │ PostgreSQL │              │ PostgreSQL     │
     │ fbso_portal│              │ keycloak schema│
     └────────────┘              └───────────────┘
```

### Fluxo de Acesso

1. Usuário acessa `fbso.com` (admin) ou `cliente.com` (white-label) → Cloudflare DNS resolve
2. Cloudflare aplica WAF, DDoS protection, SSL termination
3. Cloudflare encaminha para Kong Gateway (HTTPS :443)
4. Kong valida JWT (Service-ID/Token-ID via Keycloak) + injeta headers
5. Kong roteia para Frontend (:3000) ou Backend (:8080) ou Keycloak (:8081)
6. Backend acessa PostgreSQL (:5432) e Redis (:6379) via rede interna DOKS
7. Toda comunicação interna é HTTP (rede privada VPC) — TLS termina no Kong

---

## 2. Sizing Detalhado

### 2.1 Produção (DOKS — DigitalOcean Kubernetes)

| Recurso | Tier | vCPU | RAM | Storage | Qtd | Custo/mês |
|:---|:---|:---:|:---:|:---:|:---:|:---|
| DOKS Node Pool | s-4vcpu-8gb | 4 | 8 GB | 160 GB SSD | 3-10 (auto-scale) | $144-480 |
| PostgreSQL Managed | db-s-4vcpu-8gb | 4 | 8 GB | 115 GB SSD | 1 (+ standby) | $240 |
| Redis Managed | db-s-2vcpu-4gb | 2 | 4 GB | 38 GB SSD | 1 | $100 |
| Load Balancer | — | — | — | — | 1 | $12 |
| Spaces (S3) | — | — | — | 250 GB | 1 | $5 |
| Container Registry | — | — | — | 50 GB | 1 | $0 (incluso) |
| **Subtotal DigitalOcean** | | | | | | **~$501-837** |

### 2.2 Cloudflare

| Plano | Custo/mês |
|:---|:---|
| Cloudflare Pro (WAF + CDN + DDoS) | $20 |
| **Subtotal Cloudflare** | **$20** |

### 2.3 Custo Total Mensal Estimado

| Ambiente | Provedor | Custo |
|:---|:---|---:|
| Produção | DigitalOcean | ~$500-840 |
| Produção | Cloudflare | $20 |
| Staging | DigitalOcean (1 node + db-s-2vcpu) | ~$120 |
| **Total Mensal** | | **~$640-980** |

---

## 3. Ambientes

| Ambiente | Infra | Sizing | Uso | Custo/mês |
|:---|:---|:---|:---|:---|
| **Dev** | Docker Compose local | — | Desenvolvimento individual | $0 |
| **CI** | GitHub Actions runners | — | Testes automatizados a cada PR | $0 (incluído) |
| **Staging** | DOKS (1 node) + PostgreSQL (db-s-2vcpu) | Reduzido (1/3 prod) | Homologação, E2E, performance smoke | ~$120 |
| **Produção** | DOKS (3-10 nodes) + PostgreSQL HA + Redis | Completo | Operação real | ~$500-840 |

---

## 4. Disaster Recovery

| Parâmetro | Valor |
|:---|:---|
| **RPO** | 1 hora (PostgreSQL WAL shipping contínuo para standby) |
| **RTO** | 4 horas (DOKS cluster recreation via Terraform + restore PostgreSQL from standby) |
| **Backup PostgreSQL** | Full diário (pg_dump) + WAL contínuo. Retenção: 30 dias |
| **Backup Redis** | Snapshot diário (RDB). Retenção: 7 dias |
| **Backup Kong Config** | `deck dump` versionado no Git. Sempre atualizado |
| **DR Drill** | Trimestral — simular failover e medir RTO real |
| **Failover PostgreSQL** | DigitalOcean promove standby automaticamente em caso de falha do primary |

---

## 5. Riscos de Infraestrutura

| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---|:---|
| DOKS node failure durante pico | Baixa | Alto | Multi-node pool + Karpenter auto-scale + multi-AZ (se disponível) |
| PostgreSQL primary failure | Baixa | Crítico | Standby com WAL shipping; DigitalOcean managed HA |
| Cloudflare outage | Muito Baixa | Alto | Status page monitoring; fallback DNS secundário |
| Custo acima do orçamento | Média | Médio | Karpenter scale-down noturno; Keda evita over-provisioning; alerta de billing |
| DR não testado → RTO real > 4h | Média | Alto | DR drill trimestral obrigatório; documentar RTO real vs estimado |

---

🤖 *Documento gerado pelo Infra/Cloud Specialist — Fase 7 do Downstream Architecture Refinement. Padrões FBSO: DigitalOcean, Cloudflare.*
