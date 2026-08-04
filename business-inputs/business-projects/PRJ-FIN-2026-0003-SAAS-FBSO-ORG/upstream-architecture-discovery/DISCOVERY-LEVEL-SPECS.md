# DISCOVERY-LEVEL-SPECS.md
## Fase 10 — Bloco C: Catálogo, Matriz & Consolidação Discovery-Level

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | DISCOVERY-LEVEL-SPECS-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Consolidação Técnica) |
| **Data** | 02 de agosto de 2026 |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em 02/08/2026 |

**Documentos Referenciados (Bloco B):**
- F1: [`DISCOVERY-LEVEL-PRD.md`](DISCOVERY-LEVEL-PRD.md)
- F2: [`DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md`](DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md)
- F3: [`DISCOVERY-LEVEL-SECURITY-DEFINITION.md`](DISCOVERY-LEVEL-SECURITY-DEFINITION.md)
- F4: [`DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md`](DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md)
- F5: [`DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md`](DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md)
- F6: [`DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md`](DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md)
- F7: [`DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md`](DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md)
- F8: [`DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md`](DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md)
- F9: [`DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md`](DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md)

---

## 1. Sumário Técnico Executivo

A **FBSO Platform** é a fundação SaaS multi-produto da FBSO.ORG. A fase Core (escopo deste Discovery) entrega o **Portal Administrativo** — gestão de clientes, planos, permissões e autoatendimento — como um **monólito modular** (backend Java/Spring Boot + frontend React/Next.js) sobre infraestrutura **DigitalOcean DOKS** com **Kong↔Keycloak** como trust boundary, **PostgreSQL RLS** multi-tenant e stack completa de observabilidade.

O projeto é organizado em **4 épicos** (EP-0001 a EP-0004), implementados sequencialmente em **16 semanas**, com **12 soluções técnicas** (2 de aplicação, 6 de infraestrutura, 3 de automação, 1 de storage). O time de Discovery conta com **7 profissionais seniores** dedicados por 1 sprint (Sprint 0).

---

## 2. Consolidação por Disciplina

### 2.1 Arquitetura (F2)

| Decisão | Detalhe |
|---------|---------|
| **Abordagem** | Monólito Modular com 9 módulos internos e boundaries de domínio explícitas |
| **Integração** | Kong Gateway como trust boundary; REST/HTTPS para comunicação síncrona; RabbitMQ planejado para fase futura |
| **ADRs** | 4 decisões documentadas: Kong↔Keycloak Trust Boundary, RLS Multi-Tenant, Soft Delete, REST (sem WebSockets) |
| **Principais Riscos** | GraalVM Native Image com Spring Boot; acoplamento excessivo entre módulos internos |

### 2.2 Segurança (F3)

| Decisão | Detalhe |
|---------|---------|
| **Modelo** | Defense-in-depth: 5 camadas (Cloudflare → Kong → Istio → App → PostgreSQL) |
| **Threat Model** | 10 ameaças STRIDE priorizadas (3 críticas: cross-tenant, cross-BU, privilege escalation) |
| **Compliance** | LGPD (4 artigos mapeados); ASVS Level 2 como alvo; PCI-DSS e SOC2 futuros |
| **Trust Boundary** | Kong exclusivo — backend nunca exposto diretamente; NetworkPolicy K8s enforce |

### 2.3 Dados (F4)

| Decisão | Detalhe |
|---------|---------|
| **Banco** | PostgreSQL 17 (DO Managed) — schema `fbso_portal` com RLS |
| **Isolamento** | RLS multi-tenant via `tenant_id` + `bu_id`; Soft Delete em todas as tabelas |
| **Cache** | Redis para sessão, permissões, rate limiting, plan modules |
| **Volumes** | Ano 1: ~500 tenants, 5 GB; Ano 3: ~2.500 tenants, 50 GB |

### 2.4 DevOps/SRE (F5)

| Decisão | Detalhe |
|---------|---------|
| **CI/CD** | GitHub Actions (10 stages); Deploy via Helm + ArgoCD |
| **IaC** | Terraform para provisioning (DOKS, DBs, Redis, Spaces); Ansible para configuração |
| **Observabilidade** | Prometheus + Loki + Jaeger + Grafana + Elastic Stack; 5 dashboards; 5 alertas críticos |
| **SLOs** | 99.9% availability; P95 < 500ms; Error rate < 1%; MTTR < 30 min |

### 2.5 Testes (F6)

| Decisão | Detalhe |
|---------|---------|
| **Pirâmide** | 60% unitários, 30% integração, 10% E2E |
| **Cobertura** | ≥ 80% backend; 6 fluxos E2E críticos; teste de isolamento cross-tenant automatizado |
| **Quality Gates** | SAST + Secret Scan + Unit Tests + Integration Tests — block on failure |

### 2.6 Infra/Cloud (F7)

| Decisão | Detalhe |
|---------|---------|
| **Provedor** | DigitalOcean exclusivo (saopaulo-1); Cloudflare no edge |
| **Orquestração** | DOKS + Istio (mTLS STRICT) + Keda + Karpenter |
| **Custos** | Prod: ~$693/mês (DO + Cloudflare); 3 ambientes: ~$973/mês |
| **DR** | RPO < 24h, RTO < 4h (cross-region nyc3); PostgreSQL failover automático |

---

## 3. Matriz de Consistência Cross-Disciplina

| Par | Verificação | Status |
|-----|------------|--------|
| ARCH ↔ SEC | Kong trust boundary + 5 camadas de defesa alinhadas | ✅ |
| ARCH ↔ DATA | 9 módulos internos ↔ 9 entidades core mapeadas | ✅ |
| ARCH ↔ DEVOPS | Monólito modular ↔ pipeline único com build unificado | ✅ |
| ARCH ↔ TEST | E2E cobre fluxos cross-module (onboarding, RBAC, tenant lifecycle) | ✅ |
| ARCH ↔ INFRA | DOKS suporta topologia: Kong → Istio → Backend → PostgreSQL/Redis | ✅ |
| SEC ↔ DATA | RLS enforce isolamento; audit trail imutável para compliance | ✅ |
| SEC ↔ INFRA | NetworkPolicy + mTLS + Cloudflare WAF — 3 camadas de rede | ✅ |
| DEVOPS ↔ INFRA | Terraform provisiona; Helm deploy; ArgoCD sync | ✅ |
| DEVOPS ↔ TEST | Quality gates no pipeline CI; SAST + Secret Scan + Unit Tests | ✅ |
| TEST ↔ SEC | Teste de isolamento cross-tenant automatizado; pentest pré-lançamento | ✅ |

---

## 4. Premissas para ROM

1. **Time:** 7 profissionais seniores, 100% dedicados, Sprint 0 de 1 semana para Discovery; implementação em squads dedicados
2. **Stack:** Corporativa FBSO (Java/Spring Boot, React/Next.js, PostgreSQL, DOKS, Kong↔Keycloak)
3. **Escopo:** 4 épicos, 13 funcionalidades MVP (conforme PRD F1)
4. **Infra:** DigitalOcean exclusivo; 3 ambientes (Dev, Staging, Prod)
5. **Riscos Mitigados:** GraalVM Native Image validado em POC nas primeiras 48h; Kong↔Keycloak configurado por IAM Specialist
6. **Exclusões:** Faturamento real, gateways de pagamento, módulos Tributali-Engine e Storekeeper Portal — fora do escopo

---

## 5. Riscos Consolidados (Top 5)

| # | Risco | Disciplina | Severidade |
|---|-------|-----------|------------|
| 1 | Cross-Tenant Data Access — falha no RLS ou tenant context | Segurança | 🔴 Crítica |
| 2 | Kong Bypass — atacante acessa backend diretamente | Segurança | 🟠 Alta |
| 3 | GraalVM Native Image — problemas de reflection/proxies | Arquitetura | 🟠 Alta |
| 4 | Istio + Keda + Karpenter — complexidade excessiva | DevOps | 🟡 Média |
| 5 | Custo de infra subestimado | Infra | 🟡 Média |

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 02/08/2026 | Criação inicial: Consolidação técnica das 6 disciplinas (F2-F7), matriz de consistência (10 pares), premissas ROM, riscos top 5 | Tech Lead |

---

🤖 *Upstream Architecture Discovery — Fase 10. Documento de consolidação para embasar ROM.*
