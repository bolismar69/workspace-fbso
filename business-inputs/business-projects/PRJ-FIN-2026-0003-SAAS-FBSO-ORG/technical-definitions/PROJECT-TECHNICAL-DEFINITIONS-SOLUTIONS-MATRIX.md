# PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX — Matriz-Mestra de Soluções

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Programa:** FBSO Platform — Portal Administrativo SaaS
- **Versão:** 1.1
- **Data de Criação:** 26 de Julho de 2026
- **Última Atualização:** 27 de Julho de 2026 (alinhamento com docs de negócio v1.2)
- **Status:** ✅ COMPLIANCE — Validado pelo Time de Arquitetura
- **Baseline de Negócio:** [Project Charter v1.2](../01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [BRD v1.2](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [Épicos v1.2](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [Features FEAT-EP-](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

---

## 1. Tabela-Mestra: Solução × Stack × Owner × Status

| ID | Solução | Tipo | Stack Principal | Tech Lead | Time Alocado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| **S01** | `ms-fbso-platform-admin` | Backend API | Java 25 + Spring Boot 3.5.14 + GraalVM | Francisco Oliveira | Bolismar, Maria, Judith (temp) | ✅ Ativo (Sprint 5) |
| **S02** | `web-app-fbso-platform-portal` | Frontend Web | Next.js 15 + React 19 + Tailwind CSS | Francisco Oliveira (até 01/11) → Tom Santos | Bolismar (full-stack), Tom (a partir 01/11) | 🔮 Início 01/11/2026 |
| **S03** | PostgreSQL 17 | Banco de Dados | PostgreSQL 17 Alpine + RLS | Carlos Caldas | Carlos Caldas | ✅ Ativo |
| **S04** | Keycloak 26 | IAM | Keycloak 26.0 + OIDC + Realms | Gertrudes Paiva | Gertrudes Paiva | ✅ Ativo |
| **S05** | Docker Compose | Dev Environment | Docker Compose v3 | Davi Silva | Davi Silva | ✅ Ativo |
| **S06** | Flyway | Migrations | Flyway 10.x (Spring Boot) | Carlos Caldas | Carlos Caldas | ✅ Ativo |
| **S07** | MailHog | Email Dev | MailHog 1.0.1 | Davi Silva | Davi Silva | ✅ Ativo |
| **S08** | OTel Collector | Observabilidade | OpenTelemetry Collector | Davi Silva | Davi + Bolismar | 🔮 Sprint 0 |
| **S09** | Grafana | Dashboards | Grafana OSS | Davi Silva | Davi + Bolismar | 🔮 Sprint 1 |
| **S10** | RabbitMQ | Mensageria | RabbitMQ 3.13+ | A definir | A definir | 🔮 Futuro |
| **S11** | GitHub Actions | CI/CD | GitHub Actions | Davi Silva | Davi + Francisco | 🔮 Sprint 0 |
| **S12** | DOKS Secrets | Secrets Mgmt | DigitalOcean Kubernetes | Davi Silva | Davi Silva | 🔮 Sprint 1 |
| **S13** | Cloudflare+DO CDN | CDN/WAF | Cloudflare + DigitalOcean | Davi Silva | Davi + Francisco | 🔮 Sprint 4 |
| **S14** | Kong API Gateway | API Gateway | Kong + OIDC Plugin | Davi Silva | Davi + Francisco | 🔮 Sprint 0 |

---

## 2. Capacidade Alocada vs. Necessária

| Solução | Perfis Necessários | Capacidade Disponível | Carga Estimada | Cobertura |
|:---|:---|:---|:---|:---:|
| **S01** | Java Backend ×2, Full-Stack ×1 | Francisco ★★★, Bolismar ★★★, Maria ★☆☆, Judith ★★★ (temp) | 160h/semana (4 devs × 40h) | ✅ 85% (Maria reduz média) |
| **S02** | React/Next.js ×2 | Bolismar ★★★ (full-stack), Tom ★★★ (01/11) | 40h/semana (apenas Bolismar até 01/11) | 🔴 25% até 01/11, ✅ 100% após |
| **S03** | DBA ×1 | Carlos ★★★ | 40h/semana | ✅ 100% |
| **S04** | IAM Specialist ×1 | Gertrudes ★★★ | 40h/semana | ✅ 100% |
| **S05** | DevOps ×1 | Davi ★★★ | 10h/semana (manutenção) | ✅ 100% |
| **S06** | DBA ×1 | Carlos ★★★ | Compartilhado com S03 | ✅ 100% |
| **S07** | DevOps ×1 | Davi ★★★ | 5h/semana (manutenção) | ✅ 100% |
| **S08** | DevOps ×1, Dev ×1 | Davi ★★★, Bolismar ★★★ | 20h/semana setup inicial | ✅ 100% |
| **S09** | DevOps ×1 | Davi ★★★, Bolismar ★★☆ | 10h/semana | ✅ 100% |
| **S10** | A definir | A definir | 0h (futuro) | ⚪ N/A |
| **S11** | DevOps ×1, TL ×1 | Davi ★★★, Francisco ★★★ | 20h/semana setup inicial | ✅ 100% |
| **S12** | DevOps ×1 | Davi ★★★ | 10h/semana setup | ✅ 100% |
| **S13** | DevOps ×1 | Davi ★★★, Francisco ★★☆ | 20h/semana (Sprint 4) | ✅ 100% |
| **S14** | DevOps ×1, TL ×1 | Davi ★★★, Francisco ★★★ | 20h/semana setup inicial | ✅ 100% |

---

## 3. Indicadores por Solução

| Solução | Cobertura Skills | Risco Gargalo | Status Geral |
|:---|:---:|:---:|:---:|
| **S01** | 🟢 85% | 🟡 Maria (★☆☆) reduz velocidade | ✅ Saudável |
| **S02** | 🔴 25% | 🔴 Sem dev frontend dedicado até 01/11 | ⚠️ Crítico — Bolismar único |
| **S03** | 🟢 100% | 🟢 Sem gargalo | ✅ Saudável |
| **S04** | 🟢 100% | 🟡 SAML conhecimento único (Gertrudes) — mitigado: OIDC apenas | ✅ Saudável |
| **S05** | 🟢 100% | 🟢 Manutenção apenas | ✅ Saudável |
| **S06** | 🟢 100% | 🟢 Sem gargalo | ✅ Saudável |
| **S07** | 🟢 100% | 🟢 Manutenção apenas | ✅ Saudável |
| **S08** | 🟢 100% | 🟢 Setup simples | 🔮 Planejado |
| **S09** | 🟢 100% | 🟢 Setup simples | 🔮 Planejado |
| **S10** | ⚪ N/A | ⚪ Futuro | 🔮 Futuro |
| **S11** | 🟢 100% | 🟢 Setup simples | 🔮 Planejado |
| **S12** | 🟢 100% | 🟢 Setup simples | 🔮 Planejado |
| **S13** | 🟢 100% | 🟢 Setup Sprint 4 | 🔮 Planejado |
| **S14** | 🟢 100% | 🟢 Setup Sprint 0 | 🔮 Planejado |

---

## 4. Matriz RACI

| Atividade / Decisão | **R**esponsável | **A**utoridade | **C**onsultado | **I**nformado |
|:---|:---|:---|:---|:---|
| Arquitetura (decisões cross-solution) | Bruno (SA) | Francisco (TL) | Davi, Gertrudes | Time todo |
| Backend (S01) | Francisco (TL) | Francisco (TL) | Bolismar, Maria, Bruno | Time todo |
| Frontend (S02) | Bolismar (FS) → Tom (FE) | Francisco (TL) | Bruno, Gertrudes | Time todo |
| Banco de Dados (S03, S06) | Carlos (DB) | Francisco (TL) | Bruno | Backend team |
| IAM (S04) | Gertrudes (IAM) | Francisco (TL) | Bruno, Davi | Time todo |
| Infra/DevOps (S05, S08-S14) | Davi (DevOps) | Francisco (TL) | Bruno, Bolismar | Time todo |
| Qualidade / Testes | Felipe (QA) | Francisco (TL) | Bolismar, Mauro | Time todo |
| Homologação de Negócio | Mauro (BA) | Francisco (TL) | Felipe | Stakeholders |
| Deploy / Produção | Davi (DevOps) | Francisco (TL) | Bruno, Bolismar | Time todo |
| Segurança | Gertrudes (IAM) + Davi (DevOps) | Francisco (TL) | Bruno, Felipe | Time todo |

---

## 5. Repositórios e Pastas

| Solução | Caminho | Repositório |
|:---|:---|:---|
| **S01** | `backend/java/spring/microservices/ms-fbso-platform-admin/` | GitHub |
| **S02** | `frontend/javascript/react/web_apps/web_app-fbso-platform-portal/` | GitHub (a criar) |
| **S03-S10, S14** | Infraestrutura como código em `ms-fbso-platform-admin/` (docker-compose, keycloak) + DOKS | GitHub |
| **S11** | `.github/workflows/` | GitHub |
| **S12** | DOKS (não versionado em repositório) | DigitalOcean |
| **S13** | Cloudflare Dashboard + API | Cloudflare |

---

## 6. Referências

| Documento | Conteúdo |
|:---|:---|
| [TEAM-MAP](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md) | Skills por profissional |
| [TEAM-CAPACITY](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md) | Horas/semana + contatos |
| [TEAM-CAPACITY-EXCEPTIONS](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY-EXCEPTIONS.md) | Exceções de capacidade |
| [SOLUTIONS-CATALOG](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md) | Catálogo detalhado das 14 soluções |
| [STACK-MATRIX](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md) | Stacks por solução |
| [MILESTONES](./PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md) | Roadmap M1-M7 |

---

## Histórico de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 26/07/2026 | Criação inicial: tabela-mestra 14 soluções, capacidade vs demanda, indicadores, matriz RACI, repositórios. | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Consolidação de todas as fases anteriores. Resultado da Fase 9 do Roadmap de Definições Técnicas.*
