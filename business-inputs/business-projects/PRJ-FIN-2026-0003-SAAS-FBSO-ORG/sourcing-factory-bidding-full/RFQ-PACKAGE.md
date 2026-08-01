# RFQ-PACKAGE — Pacote de Request for Quotation (Full Mode)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG — FBSO Platform (Portal Administrativo SaaS)
- **Modo:** Full — Projeto Completo (Features + User Stories + PERT Detail-Level)
- **Data:** 31/07/2026
- **Status:** RFQ ABERTO PARA COTAÇÃO

---

## 1. Resumo do Projeto

A **FBSO Platform** é uma plataforma SaaS multi-produto que opera no modelo de Suíte com módulos ativáveis por plano contratado. O Core administrativo — escopo desta RFQ — é a camada fundamental que gerencia contas de clientes (Tenants), planos comerciais, assinaturas, usuários e permissões de acesso (RBAC).

### Escopo

| Nível | Quantidade |
|:---|---:|
| Entregas (D1-D7) | 7 |
| Épicos | 4 |
| Features | 18 (16 Must Have + 2 Should Have) |
| User Stories | 62 |

### Baseline Técnica de Referência (Downstream Architecture Refinement)

A FBSO.ORG concluiu o **Downstream Architecture Refinement** com os seguintes artefatos disponíveis para consulta:

| Disciplina | Documento | Conteúdo |
|:---|:---|:---|
| PRD | `DETAIL-LEVEL-PRD.md` | Visão do produto, personas, jornadas, escopo D1-D7 |
| Arquitetura | `DETAIL-LEVEL-ARCHITECTURE-DEFINITION.md` | C4 L2/L3, 6 ADRs, package structure, matriz de integração |
| Segurança | `DETAIL-LEVEL-SECURITY-DEFINITION.md` | STRIDE 6 componentes, 21 controles OWASP ASVS, IAM specs, RBAC matrix |
| Dados | `DETAIL-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` | ERD 15 tabelas, RLS, particionamento, Flyway migrations |
| DevOps/SRE | `DETAIL-LEVEL-DEVOPS-SRE-DEFINITION.md` | GitHub Actions, Terraform, SLOs, Prometheus/Loki/Jaeger |
| Testes | `DETAIL-LEVEL-TEST-STRATEGY-DEFINITION.md` | Pirâmide de testes, matriz cobertura, quality gates |
| Infra/Cloud | `DETAIL-LEVEL-INFRA-CLOUD-DEFINITION.md` | Topologia, sizing DOKS, custos, DR (RPO 1h/RTO 4h) |

### Estimativa de Referência (PERT Bottom-Up)

A FBSO.ORG realizou estimativa bottom-up PERT independente como baseline de referência:

| Cenário | Horas | Homem-Mês |
|:---|---:|---:|
| Desenvolvimento (62 US) | 3,365h | 21 h-m |
| + QA (30%) + Arch (8%) + DevOps (7%) + Gestão (10%) | 5,697h | 36 h-m |
| **Com Contingência 20%** | **~7,300h** | **46 h-m** |
| **Duração estimada (7 FTE)** | **~5.4 – 6.5 meses** | |

> ⚠️ **Importante:** Esta estimativa é uma **referência interna** da FBSO.ORG para comparação. As fábricas devem produzir suas **próprias estimativas independentes** usando o schema padronizado.

---

## 2. Stack Tecnológica Obrigatória (Padrões Corporativos FBSO)

| Categoria | Tecnologia |
|:---|:---|
| **Cloud** | DigitalOcean (DOKS, PostgreSQL managed, Redis, Spaces) |
| **Edge/CDN/WAF** | Cloudflare |
| **Backend** | Java 25 LTS, Spring Boot 3.5.14, GraalVM Native Image |
| **Frontend** | TypeScript 5.x, Next.js 15, React 19 |
| **Banco de Dados** | PostgreSQL 17 (RLS multi-tenant) |
| **IAM** | Keycloak 26 (OIDC + realms por tenant) |
| **API Gateway** | Kong Gateway (Service-ID/Token-ID via Keycloak, header injection) |
| **Observabilidade** | Prometheus, Grafana Loki, Jaeger, OpenTelemetry, Grafana, Elastic Stack |
| **IaC** | Terraform, Ansible |
| **Orquestração** | Kubernetes (DOKS), Istio, Keda, Karpenter |
| **CI/CD** | GitHub Actions |
| **Cache** | Redis 7 |

---

## 3. Escopo Detalhado por Épico

### EP-0001 — Portal Admin (D1 · 7 US)

| Feature | US | Descrição |
|:---|---:|:---|
| Dashboard de Métricas | 3 | Indicadores principais, filtros por período, gráfico de evolução |
| Visão de Contas com Filtros | 2 | Lista paginada, busca textual |
| Alertas e Indicadores (Should) | 2 | Indicadores de atenção, destaque visual |

### EP-0002 — Clientes e Assinaturas (D2+D3 · 16 US)

| Feature | US | Descrição |
|:---|---:|:---|
| Cadastro e Ativação de Contas | 4 | CRUD tenant, email boas-vindas, reenvio ativação |
| Gestão de Status do Tenant | 3 | State machine (Pending→Active→Suspended→Inactive), bloqueio imediato |
| Configuração de Planos | 4 | CRUD planos, módulos inclusos, versionamento |
| Vinculação de Assinaturas | 3 | Vincular tenant×plano, upgrade/downgrade, suspensão |
| Histórico de Auditoria | 2 | Log automático, filtro por período/ação |

### EP-0003 — RBAC (D4 · 16 US)

| Feature | US | Descrição |
|:---|---:|:---|
| Cadastro e Convite de Usuários | 6 | Convite, CRUD, suspensão temporária, reativação |
| Papéis e Permissões | 4 | Admin/Gerente/Operador/Auditor, permissões predefinidas |
| Vinc. Usuário×Unidade×Módulo | 3 | Many-to-many: usuário ↔ BU ↔ módulo |
| Controle de Visibilidade | 3 | Menu dinâmico, botões condicionais, redirect acesso negado |

### EP-0004 — Portal Cliente (D5+D6+D7 · 23 US)

| Feature | US | Descrição |
|:---|---:|:---|
| Autenticação e Recuperação | 3 | Login OIDC, recuperação senha, bloqueio temporário |
| Onboarding Guiado | 5 | Wizard 4 passos, dados cadastrais, 1ª BU, resumo plano |
| Dashboard do Cliente (Should) | 3 | Resumo conta, notificações, upgrade self-service |
| App Switcher | 3 | Seletor de módulos, adaptação dinâmica de menu |
| Unidades de Negócio | 5 | Hierarquia matriz/filiais, CRUD, seletor de BU |
| Catálogo de Produtos | 4 | CRUD produtos/serviços, SKU, ativar/desativar |

---

## 4. Requisitos Técnicos

### 4.1 Multi-Tenancy
- Isolamento via PostgreSQL RLS (Row-Level Security) + discriminator column `tenant_id`
- Kong injeta header `X-Tenant-ID` — backend confia sem revalidar JWT
- Zero cross-tenant data leakage (testes automatizados obrigatórios)

### 4.2 Autenticação e Autorização
- Padrão corporativo Kong↔Keycloak Service-ID/Token-ID
- Microserviços NÃO revalidam JWT
- RBAC granular: 6 roles × 20+ permissões × escopo (BU + módulo)

### 4.3 Observabilidade
- Métricas: Prometheus + Grafana
- Logs: Grafana Loki (aplicação) + Elastic Stack (auditoria)
- Tracing: Jaeger via OpenTelemetry

### 4.4 Infraestrutura
- DigitalOcean DOKS (Kubernetes) com Istio service mesh
- Keda (pod autoscaling) + Karpenter (node autoscaling)
- IaC: Terraform + Ansible
- DR: RPO 1h, RTO 4h

---

## 5. Instruções para as Fábricas

1. **Preencher o schema CSV** (`ESTIMATION-SCHEMA.csv`) com estimativas por épico
2. **Estimativa independente** — não usar a baseline PERT como ponto de partida
3. **Incluir todos os papéis:** Dev, QA, Arquitetura, DevOps, Gestão
4. **Prazo:** Informar prazo de entrega em meses e time estimado
5. **Premissas:** Documentar premissas e exclusões
6. **Prazo de resposta:** 5 dias úteis

---

## 6. Critérios de Avaliação

| Critério | Peso |
|:---|---:|
| QA ≥ 25% do total de horas | Eliminatório |
| Arquitetura ≥ 5% do total | Eliminatório |
| Consistência Prazo×Horas | Eliminatório (divergência >50% → rejeitado) |
| Total de horas vs baseline PERT | Comparativo |
| Prazo de entrega | 30% |
| Custo total (horas × rate) | 40% |
| Qualidade da proposta técnica | 30% |

---

## 7. Documentos Anexos

Todos disponíveis em `downstream-architecture-refinement/`:

1. `DETAIL-LEVEL-PRD.md` — Visão do produto e escopo de negócio
2. `DETAIL-LEVEL-ARCHITECTURE-DEFINITION.md` — Arquitetura C4 + ADRs
3. `DETAIL-LEVEL-SECURITY-DEFINITION.md` — Threat model + OWASP + RBAC
4. `DETAIL-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Modelo de dados + RLS
5. `DETAIL-LEVEL-DEVOPS-SRE-DEFINITION.md` — Pipeline + SLOs
6. `DETAIL-LEVEL-TEST-STRATEGY-DEFINITION.md` — Estratégia de testes
7. `DETAIL-LEVEL-INFRA-CLOUD-DEFINITION.md` — Infra + sizing
8. `BOTTOM-UP-PERT-ESTIMATE.md` — Baseline de referência (NÃO usar como ponto de partida)
9. `SCOPE-SNAPSHOT.md` — Escopo completo (62 US)

---

🤖 *RFQ Package — Fase 1 do Sourcing & Factory Bidding (Full Mode). Baseline: Downstream Architecture Refinement v2.0.*
