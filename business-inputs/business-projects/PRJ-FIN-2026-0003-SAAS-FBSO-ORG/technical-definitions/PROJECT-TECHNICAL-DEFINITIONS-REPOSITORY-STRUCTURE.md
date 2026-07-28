# PROJECT-TECHNICAL-DEFINITIONS-REPOSITORY-STRUCTURE — Estrutura de Diretórios do Workspace

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Programa:** FBSO Platform — Portal Administrativo SaaS
- **Versão:** 1.0
- **Data de Criação:** 27 de Julho de 2026
- **Última Atualização:** 27 de Julho de 2026
- **Status:** ✅ COMPLIANCE — Validado pelo Time de Arquitetura
- **Baseline de Negócio:** [Project Charter v1.2](../01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [BRD v1.2](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [Épicos v1.2](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [Features FEAT-EP-](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)
- **Documentos Complementares:** [SOLUTIONS-CATALOG](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md) · [STACK-MATRIX](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md) · [ARCHITECTURE-DEFINITION](./PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md)
- **Origem:** Planejamento técnico original (TECHNICAL-PLAN.md §4.3, removido após consolidação)

---

## 1. Objetivo

Este documento define a **estrutura de diretórios esperada** para o workspace `workspace-fbso/`, abrangendo os diretórios de documentação de negócio, código-fonte (backend + frontend), especificações técnicas por solução (`.specs/`) e infraestrutura como código (Docker + Kubernetes). Ele serve como referência de localização para novos desenvolvedores e como baseline para scripts de scaffold e CI/CD.

---

## 2. Convenções de Localização

| Tipo de Artefato | Localização | Exemplo |
|:---|:---|:---|
| **Documentos de Negócio** | `business-inputs/business-projects/PRJ-FIN-YYYY-NNNN-.../` | Charter, BRD, Épicos, Features, User Stories |
| **Definições Técnicas** | `business-inputs/.../technical-definitions/` | Stack, Arquitetura, Segurança, Specs, PRD |
| **Código Backend** | `backend/java/spring/microservices/{nome}/` | `ms-fbso-platform-admin` |
| **Código Frontend** | `frontend/javascript/react/web_apps/{nome}/` | `web_app-fbso-platform-portal` |
| **Infraestrutura** | `infra/` | Docker Compose, K8s manifests |
| **Especificações por Solução** | `{projeto}/.specs/` | API (OpenAPI), arquitetura, domínio, engenharia, governança |
| **Contrato de API (fonte canônica)** | `backend/.../ms-fbso-platform-admin/.specs/api/` | `fbso-platform-api.yaml` |
| **Contrato de API (cópia consumo)** | `frontend/.../web_app-fbso-platform-portal/.specs/api/` | `fbso-platform-api.yaml` (cópia sincronizada) |

---

## 3. Árvore de Diretórios (Alvo)

```
workspace-fbso/
│
├── business-inputs/business-projects/
│   └── PRJ-FIN-2026-0003-SAAS-FBSO-ORG/
│       ├── README.md                                    ← Índice geral do projeto
│       ├── 01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md
│       ├── 02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md
│       ├── 03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md
│       ├── 04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md
│       ├── user-stories/
│       │   └── 05-USER-STORIES-*.md                     ← 18 arquivos (1 por feature)
│       ├── technical-definitions/
│       │   ├── PROJECT-TECHNICAL-DEFINITIONS-EXECUTION-HISTORY.md
│       │   ├── PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md
│       │   ├── PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY-EXCEPTIONS.md
│       │   ├── PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md
│       │   ├── PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md
│       │   ├── PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md
│       │   ├── PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md
│       │   ├── PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md
│       │   ├── PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md
│       │   ├── PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md
│       │   ├── PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md
│       │   ├── PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md
│       │   ├── PROJECT-TECHNICAL-DEFINITIONS-REPOSITORY-STRUCTURE.md  ← este documento
│       │   └── sprints/
│       │       ├── sprint-00/    ← Setup & Fundação
│       │       ├── sprint-01/    ← EP-0001 (Dashboard + Contas)
│       │       ├── sprint-02/    ← EP-0001 + EP-0002 início
│       │       ├── sprint-03/    ← EP-0002 (Clientes e Assinaturas)
│       │       ├── sprint-04/    ← EP-0002 final + EP-0003 início
│       │       ├── sprint-05/    ← EP-0003 (RBAC)
│       │       ├── sprint-06/    ← EP-0003 final + EP-04a início
│       │       ├── sprint-07/    ← EP-04a (Portal Cliente)
│       │       ├── sprint-08/    ← EP-04a final + EP-04b início
│       │       ├── sprint-09/    ← EP-04b (BUs + Catálogo)
│       │       └── sprint-10/    ← Homologação final (M7)
│       ├── 05-MATRIZ-RASTREABILIDADE-RTM.md
│       ├── DEFINITION_OF_DONE.md
│       ├── GLOSSARY.md
│       ├── MATRIZ-KPI.md
│       ├── STAKEHOLDER-MAP.md
│       ├── TECHNICAL-TEAM-MAP.md
│       └── PROJECT-TEAM-MAP.md
│
├── backend/java/spring/microservices/
│   └── ms-fbso-platform-admin/                          ← Backend: S01
│       ├── pom.xml
│       ├── Dockerfile                                    ← GraalVM Native Image
│       ├── Dockerfile.jvm                                ← Fallback JVM
│       ├── mvnw / mvnw.cmd                               ← Maven Wrapper
│       ├── src/main/java/com/fbso/platform/admin/
│       │   ├── config/           ← SecurityConfig, TenantContext, WebConfig
│       │   ├── security/         ← JWT Filter, Tenant Isolation, RBAC Interceptor
│       │   ├── tenant/           ← REST /tenants
│       │   ├── plan/             ← REST /plans
│       │   ├── subscription/     ← REST /subscriptions
│       │   ├── user/             ← REST /users
│       │   ├── permission/       ← REST /permissions
│       │   ├── businessunit/     ← REST /business-units
│       │   ├── product/          ← REST /products
│       │   ├── dashboard/        ← REST /dashboard/admin, /dashboard/client
│       │   ├── onboarding/       ← REST /onboarding
│       │   ├── audit/            ← REST /audit + Entity Listener
│       │   └── common/           ← BaseEntity, SoftDeleteRepository
│       ├── src/main/resources/
│       │   ├── application.yml
│       │   └── db/migration/     ← Flyway migrations (V001__..., U001__...)
│       ├── src/test/java/        ← Testes unitários e de integração
│       └── .specs/
│           ├── api/
│           │   └── fbso-platform-api.yaml               ← OpenAPI (fonte canônica)
│           ├── architecture/
│           │   └── ARCHITECTURE-C4.md                    ← C4 L3 + Deployment
│           ├── domain/
│           ├── engineering/
│           ├── governance/
│           │   └── ENVIRONMENTS.md
│           └── business-projects/
│               └── PRJ-FIN-2026-0003-SAAS-FBSO-ORG/
│                   ├── SPECS.md
│                   ├── TASKS.md
│                   ├── TEST_PLAN.md
│                   └── ARCHITECTURE.md
│
├── frontend/javascript/react/web_apps/
│   └── web_app-fbso-platform-portal/                    ← Frontend: S02
│       ├── package.json
│       ├── next.config.js
│       ├── tailwind.config.js
│       ├── tsconfig.json
│       ├── src/
│       │   ├── app/
│       │   │   ├── (auth)/          ← Login, reset password
│       │   │   ├── (onboarding)/    ← Wizard 4 passos
│       │   │   ├── (admin)/         ← Rotas time FBSO.ORG
│       │   │   └── (portal)/        ← Rotas cliente
│       │   ├── components/
│       │   │   ├── layout/          ← AppSwitcher, Sidebar, BUSelector
│       │   │   ├── dashboard/       ← MetricsCard, TrendChart
│       │   │   └── common/          ← DataTable, StatusBadge
│       │   ├── lib/
│       │   │   ├── auth.ts          ← Keycloak integration (next-auth)
│       │   │   ├── api-client.ts    ← HTTP client with JWT injection
│       │   │   └── permissions.ts   ← usePermission(resource, action)
│       │   └── mocks/
│       │       └── handlers/        ← MSW handlers (OpenAPI-based)
│       └── .specs/
│           ├── api/
│           │   └── fbso-platform-api.yaml               ← Cópia do contrato
│           ├── architecture/
│           ├── design/
│           ├── engineering/
│           ├── governance/
│           │   └── ENVIRONMENTS.md
│           └── business-projects/
│               └── PRJ-FIN-2026-0003-SAAS-FBSO-ORG/
│                   ├── SPECS.md
│                   ├── TASKS.md
│                   ├── TEST_PLAN.md
│                   └── ARCHITECTURE.md
│
└── infra/
    ├── docker/
    │   ├── docker-compose.yml                           ← Dev local (DB + Keycloak + MailHog)
    │   └── keycloak/
    │       └── realm-config.json                        ← Realm FBSO Platform
    └── k8s/
        ├── namespace.yaml
        ├── backend-deployment.yaml
        ├── frontend-deployment.yaml
        ├── keycloak-deployment.yaml
        └── postgres-statefulset.yaml
```

---

## 4. Estrutura `.specs/` por Solução

Cada solução técnica (S01 Backend, S02 Frontend) mantém seu próprio diretório `.specs/` com artefatos técnicos versionados junto ao código. Esta estrutura garante que as especificações evoluam em sincronia com a implementação.

| Subdiretório | Conteúdo | Soluções |
|:---|:---|:---|
| **`api/`** | OpenAPI YAML — contrato de API | S01 (fonte canônica), S02 (cópia de consumo) |
| **`architecture/`** | Diagramas C4 L3, deployment, decisões de design internas | S01, S02 |
| **`domain/`** | Modelos de domínio, diagramas de sequência, DDD | S01 |
| **`design/`** | Design system, tokens, wireframes | S02 |
| **`engineering/`** | Guias de setup, convenções, troubleshooting | S01, S02 |
| **`governance/`** | ENVIRONMENTS.md, políticas de code review, quality gates | S01, S02 |
| **`business-projects/PRJ-FIN-YYYY-NNNN-.../`** | SPECS.md, TASKS.md, TEST_PLAN.md, ARCHITECTURE.md por projeto | S01, S02 |

### 4.1 Fluxo do Contrato de API

```
[Backend] .specs/api/fbso-platform-api.yaml  ← FONTE CANÔNICA
        │
        │  Codegen (OpenAPI Generator)
        │
        ├──► [Backend] Interfaces Java geradas
        │
        └──► [Frontend] Cópia sincronizada → Tipos TypeScript gerados
             └── .specs/api/fbso-platform-api.yaml  ← CÓPIA DE CONSUMO
```

---

## 5. Localização de Artefatos por Tipo

| Artefato | Projeto de Negócio | Backend (S01) | Frontend (S02) | Infra |
|:---|:---:|:---:|:---:|:---:|
| **Project Charter** | ✅ | — | — | — |
| **BRD** | ✅ | — | — | — |
| **Épicos / Features / US** | ✅ | — | — | — |
| **RTM** | ✅ | — | — | — |
| **Definições Técnicas** | ✅ | 🔗 ref | 🔗 ref | 🔗 ref |
| **OpenAPI YAML** | — | ✅ fonte | 🔗 cópia | — |
| **SPECS.md** | — | ✅ `.specs/` | ✅ `.specs/` | — |
| **TASKS.md** | — | ✅ `.specs/` | ✅ `.specs/` | — |
| **TEST_PLAN.md** | — | ✅ `.specs/` | ✅ `.specs/` | — |
| **ARCHITECTURE.md** | — | ✅ `.specs/` | ✅ `.specs/` | — |
| **Dockerfile** | — | ✅ | ✅ | — |
| **docker-compose.yml** | — | — | — | ✅ |
| **K8s Manifests** | — | — | — | ✅ |

---

## 6. Referências

| Documento | Relação |
|:---|:---|
| [STACK-MATRIX](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md) | Stack tecnológica de cada solução — decisões originais de planejamento |
| [SOLUTIONS-CATALOG](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md) | Catálogo das 14 soluções referenciadas na árvore |
| [STACK-MATRIX](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md) | Stack tecnológica de cada solução |
| [ARCHITECTURE-DEFINITION](./PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md) | Topologia de deploy (dev + produção) que esta árvore suporta |
| [SPECS-DEFINITION](./PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md) | Baseline de convenções que os arquivos `.specs/` devem seguir |

---

## 7. Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 27/07/2026 | Criação inicial: árvore de diretórios alvo migrada do TECHNICAL-PLAN.md §4.3, expandida com convenções de localização, estrutura `.specs/` e matriz de artefatos | Time Técnico |

---

🤖 *Documentação migrada do TECHNICAL-PLAN.md §4.3 e expandida com explicações da estrutura `.specs/` por solução.*
🔍 *Revisado pelo time de arquitetura em 27/07/2026.*
