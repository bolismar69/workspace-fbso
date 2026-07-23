# PRD — Product Requirements Document (Backend: ms-fbso-platform-admin)

- **Microserviço:** `ms-fbso-platform-admin`
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Caffeine Cache + REST Assured
- **Projeto de Negócio:** [PRJ-FIN-2026-0003-SAAS-FBSO-ORG](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/)
- **Versão:** 1.16
- **Data:** 21 de Julho de 2026
- **Situação implementação:** Em Execucao
- **Status:** [STATUS: COMPLIANCE] — Validado via GATE-PRD-SCOPE em 21/07/2026. 5 dimensões validadas (4 RESSALVAS, 1 APROVADO). 6 NCs corrigidas.

---

## 1. Propósito

Este documento serve como **guia de entrada** para qualquer desenvolvedor, tech lead ou arquiteto que ingressar no time do backend `ms-fbso-platform-admin`. Ele resume o que é o projeto, onde estão os documentos oficiais de negócio e quais são as entregas esperadas deste microserviço.

> **Regra de ouro:** Os documentos de negócio no diretório `business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/` são a **fonte da verdade**. Em caso de conflito entre este PRD e os documentos de negócio, **os documentos de negócio prevalecem**.

---

## 2. O Que é a FBSO Platform

A **FBSO Platform** é o futuro SaaS multi-produto da FBSO.ORG. Nesta **Fase 0** (este projeto), estamos construindo o **Core Administrativo** — a fundação sobre a qual os módulos-produto (Tributali-Engine, Storekeeper Portal) serão acoplados futuramente.

**Este microserviço (`ms-fbso-platform-admin`) é o backend do Core.** Ele expõe a API REST que o frontend (`web_app-fbso-platform-portal`) consome.

### O Que Este Microserviço NÃO É

- ❌ **Tributali-Engine** — cálculos fiscais, IBS/CBS, Split Payment — isso virá em fase futura
- ❌ **Storekeeper Portal** — PDV, estoque, frente de caixa — fase futura
- ❌ **Gateway de pagamento** — processamento financeiro real — fase futura
- ❌ **Integrações Externas** — ERPs (Totvs, SAP, Omie), sistemas bancários e de arrecadação — fase futura
- ❌ **Comercialização de Produtos Finais** — venda direta de produtos/serviços ao consumidor final — fora do escopo do Core

---

## 3. Documentos Oficiais do Projeto

### 3.1 Índice de Documentos de Negócio

| Documento | Caminho Relativo | O que contém |
|:---|:---|:---|
| **Project Charter** | [`01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | Escopo, entregas D1-D7, marcos M1-M7, riscos, stakeholders |
| **Business Requirements** | [`02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | 10 BRs funcionais, 8 NFRs |
| **Épicos** | [`03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | EP-01 a EP-04 com objetivos e jornadas |
| **Features** | [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | 18 features, 58 user stories, 18 regras de negócio |
| **User Stories** | [`05-USER-STORIES-*.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/user-stories/) | 18 arquivos com critérios de aceitação detalhados |
| **Matriz RTM** | [`05-MATRIZ-RASTREABILIDADE-RTM.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/05-MATRIZ-RASTREABILIDADE-RTM.md) | Rastreabilidade D1-D7 → EP-01 a EP-04 → 18 features → 18 US |

### 3.2 Documentos Técnicos Transversais

| Documento | Caminho Relativo | O que contém |
|:---|:---|:---|
| **Technical Plan** | [`TECHNICAL-PLAN.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/TECHNICAL-PLAN.md) | Stack, ERD, cenários de organização, sequenciamento |
| **Architecture** | [`ARCHITECTURE.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/ARCHITECTURE.md) | Diagramas C4, 7 ADRs, fluxo de auth, estrutura de pacotes |
| **Integration Map** | [`INTEGRATION-MAP.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/INTEGRATION-MAP.md) | 8 integrações, fluxos de dados, Docker Compose |
| **API Contracts** | [`API-CONTRACTS.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/API-CONTRACTS.md) | 11 recursos REST, schemas JSON, matriz RBAC × endpoints |
| **Definition of Done** | [`DEFINITION_OF_DONE.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/DEFINITION_OF_DONE.md) | Critérios de DONE para US, Feature e Entrega |
| **Technical Team Map** | [`TECHNICAL-TEAM-MAP.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/TECHNICAL-TEAM-MAP.md) | Estrutura de times e competências |

### 3.3 Artefatos Locais (deste microserviço)

| Artefato | Localização | Status |
|:---|:---|:---|
| **PRD.md** | `.specs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/PRD.md` | ✅ Este documento |
| **SPECS.md** | `.specs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/SPECS.md` | ✅ Criado (v1.4) |
| **TASKS.md** | `.specs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/TASKS.md` | ✅ Criado (v2.2) |
| **TEST_PLAN.md** | `.specs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/TEST_PLAN.md` | ✅ Criado (v2.2) |
| **ARCHITECTURE.md** | `.specs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/ARCHITECTURE.md` | ✅ Criado (v1.2) |
| **API YAML** | `.specs/api/fbso-platform-api.yaml` | ⬜ A criar |

---

## 4. Escopo do Backend

### 4.1 Entidades que Este Microserviço Implementa

| Entidade | Tabela | Épico | API Resource |
|:---|:---|:---|:---|
| Tenant | `tenant` | EP-02 | `/tenants` |
| Plan | `plan` | EP-02 | `/plans` |
| PlanModule | `plan_module` | EP-02 | (sub-recurso de `/plans`) |
| Subscription | `subscription` | EP-02 | `/subscriptions` |
| User | `user` | EP-03 | `/users` |
| UserPermission | `user_permission` | EP-03 | `/permissions` |
| ResourceAction | `resource_action` | EP-03 | (tabela de domínio) |
| RoleResource | `role_resource` | EP-03 | (tabela de domínio) |
| BusinessUnit | `business_unit` | EP-04 | `/business-units` |
| ProductService | `product_service` | EP-04 | `/products` |
| AuditEntry | `audit_log` | EP-02 | `/audit` |

### 4.2 Entidades Fora do Escopo (NÃO IMPLEMENTAR)

| Entidade | Quando será feita |
|:---|:---|
| Billable, ProductBillableMapping | Fase Tributali-Engine |
| Order, OrderItem | Fase Storekeeper / Tributali-Engine |
| Invoice, InvoiceItem | Fase Storekeeper / Tributali-Engine |
| TransactionPayment, SplitPayment | Fase Tributali-Engine |
| BillingInfo, BankAccount | Fase Storekeeper |

### 4.3 Endpoints por Marco de Entrega

| Marco | Data | Endpoints | Épicos |
|:---|:---|:---|:---|
| **M2** | 15/08/2026 | `GET /dashboard/admin/*` | EP-01 |
| **M3** | 31/08/2026 | CRUD `/tenants`, `/plans`, `/subscriptions`, `/audit` | EP-02 |
| **M4** | 15/09/2026 | CRUD `/users`, `/permissions` | EP-03 |
| **M5** | 30/09/2026 | `/onboarding`, `/dashboard/client` | EP-04 |
| **M6** | 15/10/2026 | CRUD `/business-units`, `/products` | EP-04 |
| **M7** | 30/10/2026 | Homologação completa (todos os endpoints) | Todos |

### 4.4 Matriz de Rastreabilidade (BR → Feature → US → Entrega → API)

> **Origem:** [BRD §6.1](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [04-FEATURES](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [03-EPICS](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

| BR | Feature | User Stories | Entrega | Pacote Java | API Resource |
|:---|:---|:---|:---|:---|:---|
| BR-A01 | F01-01 Dashboard Métricas | US-001, US-002, US-003 | D1 | `dashboard` | `GET /dashboard/admin/*` |
| BR-A01 | F01-02 Visão de Contas | US-004, US-005 | D1 | `dashboard` | `GET /dashboard/admin/*` |
| BR-A01 | F01-03 Alertas e Indicadores | US-006, US-007 | D1 | `dashboard` | `GET /dashboard/admin/*` |
| BR-A02 | F02-01 Cadastro Contas | US-008 a US-011 | D2 | `tenant` | CRUD `/tenants` |
| BR-A02 | F02-02 Status do Tenant | US-012 a US-014 | D2 | `tenant` | CRUD `/tenants` |
| BR-A02 | F02-05 Histórico Auditoria | US-022, US-023 | D2 | `audit` | `GET /audit` |
| BR-A03 | F02-03 Configuração Planos | US-015 a US-018 | D3 | `plan` | CRUD `/plans` |
| BR-A04 | F02-04 Gestão Assinaturas | US-019 a US-021 | D3 | `subscription` | CRUD `/subscriptions` |
| BR-A05 | F03-01 Convite Usuários | US-024 a US-026 | D4 | `user` | CRUD `/users` |
| BR-A05 | F03-02 Papéis e Permissões | US-027 a US-030 | D4 | `permission` | CRUD `/permissions` |
| BR-A05 | F03-03 Vínculo Usuário×BU×Módulo | US-031 a US-033 | D4 | `permission` | CRUD `/permissions` |
| BR-A05 | F03-04 Visibilidade Menus | US-034 a US-036 | D4 | `permission` | (frontend + interceptor) |
| BR-B01 | F04-01 Autenticação | US-037 a US-039 | D5 | `security`, `config` | (Keycloak) |
| BR-B02 | F04-02 Onboarding Guiado | US-040 a US-044 | D5 | `onboarding` | `/onboarding` |
| BR-B01 | F04-03 Dashboard Cliente | US-045, US-046 | D5 | `dashboard` | `GET /dashboard/client` |
| BR-B03 | F04-04 App Switcher | US-047 a US-049 | D5 | `config` | (frontend + `/tenants/me`) |
| BR-B04 | F04-05 Unidades de Negócio | US-050 a US-054 | D6 | `businessunit` | CRUD `/business-units` |
| BR-B05 | F04-06 Catálogo Produtos | US-055 a US-058 | D7 | `product` | CRUD `/products` |

### 4.5 Cobertura de Features e User Stories

| Feature ID | Nome | Épico | US | Entrega | Pacote Java |
|:---|:---|:---|:---:|:---|:---|
| F01-01 | Dashboard de Métricas Operacionais | EP-01 | 3 | D1 | `dashboard` |
| F01-02 | Visão de Contas com Filtros | EP-01 | 2 | D1 | `dashboard` |
| F01-03 | Alertas e Indicadores de Atenção | EP-01 | 2 | D1 | `dashboard` |
| F02-01 | Cadastro e Ativação de Contas | EP-02 | 4 | D2 | `tenant` |
| F02-02 | Gestão de Status do Tenant | EP-02 | 3 | D2 | `tenant` |
| F02-03 | Configuração de Planos Comerciais | EP-02 | 4 | D3 | `plan` |
| F02-04 | Vinculação e Gestão de Assinaturas | EP-02 | 3 | D3 | `subscription` |
| F02-05 | Histórico de Auditoria Administrativa | EP-02 | 2 | D2 | `audit` |
| F03-01 | Cadastro e Convite de Usuários | EP-03 | 3 | D4 | `user` |
| F03-02 | Definição de Papéis e Permissões (RBAC) | EP-03 | 4 | D4 | `permission` |
| F03-03 | Vinculação Usuário × Unidade × Módulo | EP-03 | 3 | D4 | `permission` |
| F03-04 | Controle de Visibilidade de Menus e Ações | EP-03 | 3 | D4 | `permission`, `security` |
| F04-01 | Autenticação e Recuperação de Senha | EP-04 | 3 | D5 | `security`, `config` |
| F04-02 | Onboarding Guiado de Primeiro Acesso | EP-04 | 5 | D5 | `onboarding` |
| F04-03 | Dashboard do Cliente | EP-04 | 2 | D5 | `dashboard` |
| F04-04 | App Switcher (Seletor de Módulos) | EP-04 | 3 | D5 | `config` |
| F04-05 | Gestão de Unidades de Negócio | EP-04 | 5 | D6 | `businessunit` |
| F04-06 | Catálogo de Produtos e Serviços | EP-04 | 4 | D7 | `product` |
| **Total** | **18 features** | **4 épicos** | **58** | **7 entregas** | **12 pacotes** |

### 4.6 Cobertura de Business Requirements (BRs)

| BR | Descrição | Épico | Features | Prioridade |
|:---|:---|:---|:---|:---|
| BR-A01 | Dashboard Administrativo | EP-01 | F01-01, F01-02, F01-03 | Must Have |
| BR-A02 | Ativação e Gestão de Contas | EP-02 | F02-01, F02-02, F02-05 | Must Have |
| BR-A03 | Configuração de Planos Comerciais | EP-02 | F02-03 | Must Have |
| BR-A04 | Vinculação de Assinaturas | EP-02 | F02-04 | Must Have |
| BR-A05 | Gestão de Usuários e Permissões | EP-03 | F03-01 a F03-04 | Must Have |
| BR-B01 | Portal do Cliente com Autenticação | EP-04 | F04-01, F04-03 (¹) | Must Have |

> ¹ **F04-03 (Dashboard do Cliente):** Classificada como **Should Have** no nível de feature (conforme RTM — item bônus, sem entrega obrigatória no Charter). Incluída em BR-B01 para cobertura completa da jornada do cliente, mas sua implementação pode ser postergada sem bloquear D5.
| BR-B02 | Onboarding Guiado | EP-04 | F04-02 | Must Have |
| BR-B03 | App Switcher | EP-04 | F04-04 | Must Have |
| BR-B04 | Unidades de Negócio | EP-04 | F04-05 | Must Have |
| BR-B05 | Catálogo de Produtos/Serviços | EP-04 | F04-06 | Must Have |

---

### 4.7 Matriz de Dependências entre Features

| Feature | Depende de | Motivo |
|:---|:---|:---|
| F02-01 (Cadastro Contas) | F01-01 (Dashboard) | Métricas do dashboard alimentadas pelos tenants cadastrados |
| F02-03 (Planos) | F02-01 (Cadastro Contas) | Planos configurados antes da ativação de tenants |
| F02-04 (Assinaturas) | F02-01 (Cadastro Contas), F02-03 (Planos) | Assinatura vincula tenant a plano — ambos devem existir |
| F03-01 (Convite Usuários) | F02-01 (Cadastro Contas) | Usuários pertencem a tenants ativos |
| F03-02 (Permissões RBAC) | F03-01 (Convite Usuários) | Permissões atribuídas a usuários já cadastrados |
| F04-01 (Autenticação) | F02-01 (Cadastro Contas) | Login exige tenant ativo |
| F04-02 (Onboarding) | F04-01 (Autenticação) | Primeiro acesso ocorre após autenticação |
| F04-04 (App Switcher) | F02-03 (Planos), F03-02 (Permissões) | Módulos exibidos = interseção plano × permissões |
| F04-05 (Unidades Negócio) | F02-01 (Cadastro Contas) | BUs pertencem a tenant; primeira BU = Matriz no onboarding |
| F04-06 (Catálogo) | F04-05 (Unidades Negócio) | Produtos vinculados a uma BU |

### 4.8 Premissas

> **Origem:** [Project Charter §9.1](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) e [BRD §7](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

| ID | Premissa | Impacto no Microserviço |
|:---|:---|:---|
| A1 | Time técnico com competência em Java + Spring Boot | Autonomia para decisões de implementação dentro dos ADRs |
| A2 | Contratos de interface congelados no M2 (15/08/2026) | APIs definidas até M2 não sofrerão breaking changes sem renegociação |
| A3 | Stakeholders disponíveis para validação a cada sprint | Feedback de negócio disponível no prazo máximo de 3 dias úteis |
| A4 | Early adopters para validação do Portal do Cliente | Ambiente de homologação com tenants reais para testes de aceitação |
| A5 | Cronograma da Reforma Tributária estável até 10/2026 | Fase Tributali-Engine permanece como escopo futuro; sem antecipação |
| A6 | Modelo de negócios (planos com módulos ativáveis) estável | Estrutura Plan/Module/Subscription não sofrerá refatoração estrutural |

> ⚠️ **Regra de governança:** Se qualquer premissa A1-A6 for invalidada, o PRD.md deve ser revalidado e os impactos em ARCHITECTURE.md, SPECS.md, TASKS.md e TEST_PLAN.md avaliados (efeito cascata).

---

## 5. Stack e Decisões Técnicas

### 5.1 Stack

| Camada | Tecnologia |
|:---|:---|
| **Linguagem** | Java 25 |
| **Framework** | Spring Boot (versão compatível com Java 25) |
| **Build** | Maven (preferencial) |
| **Persistência** | Spring Data JDBC / JDBC Template + Flyway (migrations) |
| **Segurança** | Spring Security + JWT (Keycloak) + JWT Issuer Validation |
| **Cache** | Caffeine Cache (spring-boot-starter-cache) |
| **Container** | Docker + GraalVM Native Image (preferencial para produção) |
| **Testes** | JUnit 5 + Mockito + Testcontainers + REST Assured |

### 5.2 ADRs que Impactam Este Microserviço

| ADR | Decisão | O que significa para o backend |
|:---|:---|:---|
| **ADR-01** | Shared Database + tenant_id | Toda query DEVE filtrar por `tenant_id`. Implementar via `TenantIsolationFilter` |
| **ADR-02** | Java 25 + Spring Boot | Stack definida. Usar versão mais recente estável do Spring Boot compatível |
| **ADR-04** | Keycloak como IdP | Backend NÃO gerencia senhas. Apenas valida JWT e extrai claims |
| **ADR-05** | Soft Delete universal | `deleted_dt IS NULL` em toda query. Índices únicos parciais |
| **ADR-06** | API Contract First | OpenAPI YAML é a verdade. Código gerado ou validado contra ele |
| **ADR-07** | JWT Stateless | Sem sessão no servidor. `TenantContext` por request (ThreadLocal) |
| **ADR-08** | PostgreSQL Row-Level Security + FORCE | Camada de defesa em profundidade. `FORCE ROW LEVEL SECURITY` aplicado a 4 tabelas (subscription, user, business_unit, audit_log). Política: `USING (tenant_id = current_setting('app.current_tenant_id')::UUID)`. FORCE garante que nem o table owner (app user) escapa do RLS |

### 5.3 Estrutura de Pacotes Esperada

```
com.fbso.platform.admin/
├── config/           ← SecurityConfig, CacheConfig, TenantContext, TenantAwareDataSource
├── security/         ← JwtAuthenticationFilter, aspect/RbacAspect, aspect/AuditAspect
├── entity/           ← User, ResourceAction, RoleResource, BusinessUnit, Tenant, Plan, etc.
├── enums/            ← Role, UserStatus, SubscriptionStatus, etc.
├── controller/       ← TenantController, PlanController, UserController, etc.
├── service/          ← TenantService, PermissionService, SubscriptionService, etc.
├── repository/       ← UserRepository, TenantRepository, etc.
│   └── rowmapper/    ← UserRowMapper, TenantRowMapper, etc.
├── dto/              ← Request/Response DTOs
├── exception/        ← GlobalExceptionHandler, domain exceptions
└── common/           ← BaseEntity, BaseRepository, BaseRowMapper
```

---

## 6. Requisitos Técnicos Críticos

### 6.1 Isolamento Multi-Tenant

**Regra absoluta:** NENHUMA query SQL pode ser executada sem filtro `tenant_id`.

**Estratégia de defesa em profundidade (3 camadas):**

| Camada | Mecanismo | Tipo | Garantia |
|:---|:---|:---|:---|
| **1. PostgreSQL RLS** | `CREATE POLICY tenant_isolation USING (tenant_id = current_setting('app.current_tenant_id')::UUID)` em TODAS as tabelas com `tenant_id` | **Preventiva (banco)** | Bloqueia query que esquecer `WHERE tenant_id = ?` — nível de banco, impossível de burlar via aplicação |
| **2. BaseRepository** | Template JDBC com `AND tenant_id = ?` em todas as queries (ADR-L01) | **Preventiva (aplicação)** | Convenção explícita — todo repository que estende `BaseRepository` herda o filtro |
| **3. Teste de Isolamento** | Teste automatizado: tenant-A tenta acessar dados de tenant-B | **Detectiva** | Falha o build se isolamento for violado |

**Implementação esperada:**
- `TenantContext` — holder ThreadLocal que armazena o `tenant_id` da requisição atual
- `JwtAuthenticationFilter` — extrai `tenant_id` do JWT e seta no `TenantContext` E configura `app.current_tenant_id` na sessão PostgreSQL
- **Migration V003** — `ALTER TABLE ... ENABLE ROW LEVEL SECURITY` + `FORCE ROW LEVEL SECURITY` + `CREATE POLICY tenant_isolation` nas tabelas multi-tenant (subscription, user, business_unit, audit_log). FORCE garante que nem o table owner escapa do RLS
- **Migration V004** — Seed data da matriz RN10-01: `INSERT INTO resource_action` (8 resources × 4 actions) + `INSERT INTO role_resource` (4 roles). Criada na Sprint 4 Frente 0
- **Migration V006** — `ALTER TABLE user_permission ADD CONSTRAINT fk_up_bu FOREIGN KEY (business_unit_id) REFERENCES business_unit(id)`. Criada na Sprint 4 Frente 0
- **BaseRepository** — template JDBC com `AND tenant_id = ?` automático
- **Teste obrigatório:** Tentar acessar dados de outro tenant deve retornar vazio (não 403 — o tenant nem sabe que o outro existe). Tentar INSERT com `tenant_id` diferente do contexto → rejeitado pelo PostgreSQL RLS

### 6.2 RBAC (Role-Based Access Control)

**Implementação (DB-backed — Sprint 4 Frente 0):**
- `RbacAspect` (AOP) + `PermissionService` — verifica permissões consultando o banco (não o JWT)
- `PermissionService` carrega a matriz RN10-01 das tabelas `resource_action` + `role_resource` (seed V004)
- Roles do usuário vêm da tabela `user_permission` (user_id, business_unit_id, role) — não do JWT
- Sem cache TTL — alterações de permissão têm efeito imediato (RN11-03)
- `@RequiresPermission(resource, action)` nos controllers — o aspecto intercepta e valida
- **Teste obrigatório:** 20+ combinações papel × endpoint proibido → 403 (via REST Assured + Testcontainers)

### 6.3 Soft Delete

**Implementação esperada:**
- Superclasse `BaseEntity` com campos: `created_dt`, `updated_dt`, `created_by`, `updated_by`, `deleted_dt`, `deleted_by`
- Repository base com filtro automático `WHERE deleted_dt IS NULL`
- Endpoint de DELETE não remove fisicamente — seta `deleted_dt = NOW()` e `deleted_by = current_user`
- Índices únicos parciais (PostgreSQL `WHERE deleted_dt IS NULL`)

### 6.4 Trilha de Auditoria

**Implementação esperada:**
- `AuditEntityListener` (JPA EntityListener) — toda operação INSERT/UPDATE/DELETE gera registro em `audit_log`
- Registro contém: entidade, ID, ação, usuário, timestamp, valores anteriores e novos
- Endpoint `GET /audit` para consulta (com filtros por entidade, ação, período, tenant)

### 6.5 Validação de Entrada

**Implementação esperada:**
- Bean Validation (`jakarta.validation`) em todos os DTOs de entrada
- Mensagens de erro em português (Brasil)
- Respostas de erro no formato RFC 7807 (Problem Details)

### 6.6 Regras de Negócio Aplicáveis

> **Origem:** [04-FEATURES](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — 18 regras de negócio (RN01-01 a RN18-04)

As regras de negócio abaixo são a fonte da verdade para validações, transições de estado e restrições. **Toda RN deve ser implementada e testada.**

#### Domínio: Dashboard (EP-01)

| RN | Descrição | Feature | Impacto Técnico |
|:---|:---|:---|:---|
| RN01-01 | Métricas excluem tenants com soft delete | F01-01 | `WHERE deleted_dt IS NULL` em todas as queries do dashboard |
| RN01-02 | Período padrão do dashboard: mês atual | F01-01 | Query com range de data padrão no service |
| RN01-03 | Indicadores zero exibem "0", nunca em branco | F01-01 | Tratamento de null/empty no DTO |
| RN02-01 | Contas soft-deletadas não aparecem na lista | F01-02 | Filtro `deleted_dt IS NULL` na listagem |
| RN02-02 | Busca case-insensitive | F01-02 | `LOWER(razão_social) LIKE LOWER(:termo)` |
| RN03-01 | Alerta onboarding incompleto após 48h | F01-03 | Job ou query: `status = 'PENDING' AND created_dt < NOW() - 48h` |
| RN03-02 | Alertas sem personalização por usuário | F01-03 | Mesma query para todos os admins |

#### Domínio: Tenant (EP-02)

| RN | Descrição | Feature | Impacto Técnico |
|:---|:---|:---|:---|
| RN04-01 | Criação de Tenant gera auditoria | F02-01 | `AuditEntityListener` — INSERT na tabela `audit_log` |
| RN04-02 | Razão social obrigatória; validação de duplicidade | F02-01 | Bean Validation `@NotBlank` + query `SELECT 1 FROM tenant WHERE razao_social = :rs AND deleted_dt IS NULL` |
| RN04-03 | Link de ativação único e de uso único por Tenant | F02-01 | UUID na tabela `tenant.activation_token`; token invalidado após uso |
| RN05-01 | Transições de status permitidas (máquina de estados) | F02-02 | Enum de transições válidas; `IllegalStateTransitionException` para transições inválidas |
| RN05-02 | Suspensão exige motivo | F02-02 | Campo `motivo_suspensao` obrigatório no DTO; `@NotBlank` |
| RN05-03 | Reativação restaura permissões anteriores | F02-02 | Manter `user_permission` com soft delete; reativar em vez de recriar |

#### Domínio: Planos e Assinaturas (EP-02)

| RN | Descrição | Feature | Impacto Técnico |
|:---|:---|:---|:---|
| RN06-01 | Plano com clientes ativos não pode ser excluído | F02-03 | Apenas soft delete (desativação); constraint de integridade |
| RN06-02 | Alteração de preço não afeta assinaturas existentes | F02-03 | Versão do plano (plan_version); assinatura aponta para versão específica |
| RN06-03 | Deve existir pelo menos 1 plano ativo | F02-03 | Validação antes de desativar: `SELECT COUNT(*) FROM plan WHERE status = 'ACTIVE' AND deleted_dt IS NULL` |
| RN07-01 | Apenas 1 assinatura ativa por tenant | F02-04 | Constraint única parcial: `tenant_id WHERE status = 'ACTIVE' AND deleted_dt IS NULL` |
| RN07-02 | Upgrade/downgrade sem janela sem assinatura | F02-04 | Transação atômica: finalizar anterior + criar nova no mesmo batch |
| RN07-03 | Data de término opcional (vigência contínua) | F02-04 | Campo `end_date` nullable |
| RN08-01 | Auditoria cobre 100% das ações administrativas | F02-05 | `AuditEntityListener` intercepta create/update/delete; §6.4 |
| RN08-02 | Registros de auditoria são imutáveis (append-only) | F02-05 | Tabela `audit_entry` sem operações de UPDATE/DELETE; apenas INSERT e SELECT |

#### Domínio: Usuários e RBAC (EP-03)

| RN | Descrição | Feature | Impacto Técnico |
|:---|:---|:---|:---|
| RN09-01 | Convite expira em 7 dias | F03-01 | `invitation_expires_at = NOW() + 7 days`; filter: `WHERE expires_at > NOW()` |
| RN09-02 | E-mail único por tenant | F03-01 | Constraint única: `(tenant_id, email) WHERE deleted_dt IS NULL` |
| RN09-03 | Admin não pode desativar a si mesmo | F03-01 | Validação: `if (targetUserId == currentUserId) throw SelfDeactivationException` |
| RN10-01 | Tabela de permissões por papel (Admin/Gerente/Operador/Auditor) | F03-02 | Matriz RBAC no `RbacInterceptor`; ver §6.2 |
| RN11-01 | Usuário requer ≥1 Unidade de Negócio | F03-03 | Validação no DTO: `@Size(min=1) List<UUID> businessUnitIds` |
| RN11-02 | Usuário requer ≥1 Módulo vinculado | F03-03 | Validação no DTO: `@Size(min=1) List<UUID> moduleIds` |
| RN11-03 | Módulos disponíveis = módulos do plano contratado | F03-03 | Query: módulos do plano → filtrar lista exibida |
| RN12-01 | Ocultação de menu (UX) + bloqueio por permissão (segurança) | F03-04 | Dupla camada: frontend condicional + interceptor backend |
| RN12-02 | Nome do módulo ativo visível no topo | F03-04 | Contexto da sessão; incluído no header de resposta |

#### Domínio: Portal do Cliente (EP-04)

| RN | Descrição | Feature | Impacto Técnico |
|:---|:---|:---|:---|
| RN13-01 | Senha: mín. 8 caracteres, letra + número | F04-01 | Delegado ao Keycloak (password policy) |
| RN13-02 | Sessão expira após 60 min de inatividade | F04-01 | Configuração Keycloak: `accessTokenLifespan` |
| RN13-03 | Link de redefinição de senha de uso único | F04-01 | Delegado ao Keycloak (required actions) |
| RN14-01 | Onboarding obrigatório no primeiro acesso | F04-02 | Flag `onboarding_completed` no Tenant; redirect se false |
| RN14-02 | Primeira BU = automaticamente Matriz | F04-02 | Campo `hierarchy_type = 'MATRIZ'` na primeira BU |
| RN14-03 | Onboarding concluído = todos os passos obrigatórios finalizados | F04-02 | Validação server-side: checklist de passos × status |
| RN14-04 | Tenant muda para "Ativo" após onboarding | F04-02 | Transição automática: `PENDING_ONBOARDING → ACTIVE` |
| RN15-01 | Dashboard adapta-se ao módulo ativo no App Switcher | F04-03 | Endpoint recebe `module_id` como parâmetro |
| RN15-02 | Fase 0: dashboard genérico expansível | F04-03 | Interface `DashboardProvider` com implementação default |
| RN16-01 | App Switcher: interseção plano × permissões | F04-04 | Query: `SELECT module FROM plan_modules WHERE plan_id = :plan INTERSECT SELECT module FROM user_modules WHERE user_id = :user` |
| RN16-02 | Fase 0: módulo placeholder "FBSO Platform" | F04-04 | Módulo default criado na migration inicial |
| RN16-03 | Troca de módulo mantém contexto da BU | F04-04 | `TenantContext` preserva `current_bu_id` durante switch |

#### Domínio: Unidades de Negócio e Catálogo (EP-04)

| RN | Descrição | Feature | Impacto Técnico |
|:---|:---|:---|:---|
| RN17-01 | CNPJ único entre BUs ativas do mesmo tenant | F04-05 | Constraint única parcial: `(tenant_id, cnpj) WHERE deleted_dt IS NULL` |
| RN17-02 | Unidade desativada não pode ser "pai" | F04-05 | Validação: `if (parent.status == INACTIVE) throw InvalidParentException` |
| RN17-03 | Primeira BU = Matriz (durante onboarding) | F04-05 | Mesma lógica de RN14-02 |
| RN17-04 | Sem limite de níveis hierárquicos | F04-05 | Estrutura de árvore (parent_id auto-referenciado) |
| RN17-05 | Seletor de BU reflete permissões do usuário | F04-05 | Query: `SELECT bu FROM user_bu_permission WHERE user_id = :userId` (Admin = todas) |
| RN18-01 | Catálogo segmentado por Unidade de Negócio | F04-06 | Toda query de produto filtra por `business_unit_id` |
| RN18-02 | SKU opcional mas único por BU | F04-06 | Constraint única parcial: `(business_unit_id, sku) WHERE sku IS NOT NULL AND deleted_dt IS NULL` |
| RN18-03 | Indicador "Não mapeado" para todos os itens (placeholder fiscal) | F04-06 | Campo `fiscal_mapping_status` default = `NOT_MAPPED` |
| RN18-04 | Soft delete em produtos | F04-06 | `deleted_dt IS NULL` — mesma política de RN17-01 |

### 6.7 Cobertura de Requisitos Não-Funcionais (NFRs)

> **Origem:** [BRD §6.2](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

| NFR | Categoria | Requisito | Responsabilidade | Como o Backend Atende |
|:---|:---|:---|:---|:---|
| BR-NFR01 | Disponibilidade | 99,5% em horário comercial (6h-23h) | **Infra/DevOps** | Health checks (`/actuator/health`), readiness probes K8s, graceful shutdown |
| BR-NFR02 | Segurança | Isolamento total entre tenants e BUs | **Backend** | §6.1 (TenantIsolationFilter) + §6.2 (RbacInterceptor) — 100% das queries filtradas |
| BR-NFR03 | Auditabilidade | 100% das ações administrativas registradas | **Backend** | §6.4 (AuditEntityListener) — INSERT/UPDATE/DELETE → `audit_log` |
| BR-NFR04 | Usabilidade | Intuitivo, máx. 2h capacitação | **Frontend** | Backend provê APIs RESTful padronizadas; mensagens de erro em PT-BR (§6.5) |
| BR-NFR05 | Performance | Telas principais ≤ 3s | **Frontend + Backend** | Backend: query optimization, índices, paginação (25 registros), cache onde aplicável |
| BR-NFR06 | Escalabilidade | Acoplamento de módulos sem reestruturação | **Backend** | ADR-01 (Shared Database + tenant_id), estrutura de planos com módulos ativáveis |
| BR-NFR07 | Acessibilidade | Contraste, navegação teclado, textos alternativos | **Frontend** | Backend: sem impacto direto (APIs REST retornam dados; frontend renderiza) |
| BR-NFR08 | Idioma | PT-BR (preparado para multi-idioma) | **Backend + Frontend** | Backend: mensagens de erro e validação em PT-BR via `messages_pt_BR.properties`; estrutura `i18n` preparada |

---

### 7.1 Sequência Recomendada

```
1. Setup do projeto
   ├── Scaffold Maven (pom.xml, dependências)
   ├── Configurar Dockerfile (Native Image + JVM fallback)
   ├── Configurar application.yml (dev, staging, prod profiles)
   └── Setup Flyway/Liquibase + migrations iniciais

2. Segurança (ANTES de qualquer endpoint de negócio)
   ├── Configurar Spring Security + JWT validation
   ├── Implementar TenantContext + TenantIsolationFilter
   ├── Implementar PostgreSQL Row-Level Security (Migration V003)
   ├── Implementar RbacInterceptor
   └── Implementar AuditEntityListener

3. Desenvolvimento por épico (ordem dos marcos)
   ├── M2: EP-01 — Dashboard Admin (endpoints de leitura)
   ├── M3: EP-02 — Tenants, Plans, Subscriptions, Audit
   ├── M4: EP-03 — Users, Permissions (RBAC completo)
   ├── M5: EP-04a — Onboarding, Dashboard Client
   └── M6: EP-04b — Business Units, Products

4. Integração e Testes
   ├── Testes unitários (JUnit 5 + Mockito)
   ├── Testes de integração (Testcontainers + PostgreSQL real)
   ├── Testes de segurança (RBAC, Multi-Tenant, OWASP)
   └── Testes E2E (junto com frontend)

5. Deploy
   ├── Docker build (Native Image)
   ├── Push para registry
   └── Deploy K8s (staging → produção)
```

### 7.2 Definição de Pronto (DoD Local)

Além da [DoD do projeto](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/DEFINITION_OF_DONE.md) (12 critérios de User Story, 5 de Feature, 6 de Entrega), este microserviço adiciona:

- [ ] Cobertura de testes ≥ 80% (JaCoCo)
- [ ] Nenhum warning do Checkstyle
- [ ] OpenAPI YAML atualizado (se endpoint novo ou alterado)
- [ ] Migration script versionado e testado (rollback validado)
- [ ] Teste de isolamento Multi-Tenant executado (acesso cross-tenant bloqueado)
- [ ] Teste de RBAC executado (cada papel testado contra endpoints proibidos)
- [ ] Dockerfile Native Image build funcional

> **Critérios de Aceite por User Story:** Os critérios de aceite detalhados de cada uma das 58 user stories estão documentados em:
> - [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — coluna "Critérios de Aceitação" nas tabelas de user stories
> - [`05-USER-STORIES-*.md`](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/user-stories/) — 18 arquivos com especificações detalhadas
> - A matriz de rastreabilidade (§4.4) vincula cada US à sua feature, entrega e API resource
> - O DoD do projeto (§3 — DoD de USER STORY, §4 — DoD de FEATURE, §5 — DoD de ENTREGA) aplica-se cumulativamente

---

## 8. Ambiente de Desenvolvimento

### 8.1 Pré-requisitos

- Java 25 SDK
- Maven 3.9+
- Docker + Docker Compose
- PostgreSQL 17 (via Docker)
- Keycloak 26 (via Docker)

### 8.2 Início Rápido

```bash
# 1. Subir infraestrutura (DB + Keycloak)
cd infra/docker
docker compose up -d postgres keycloak

# 2. Executar migrations
cd backend/java/spring/microservices/ms-fbso-platform-admin
./mvnw flyway:migrate

# 3. Iniciar aplicação em modo dev
./mvnw spring-boot:run -Dspring-boot.run.profiles=dev

# 4. Verificar saúde
curl http://localhost:8081/actuator/health

# 5. Testar endpoint (com JWT de exemplo)
curl http://localhost:8081/api/v1/dashboard/admin/summary \
  -H "Authorization: Bearer <jwt-token>"
```

### 8.3 Profiles

| Profile | Descrição | Banco | Keycloak |
|:---|:---|:---|:---|
| `dev` | Desenvolvimento local | localhost:5432 | localhost:8080 |
| `staging` | Homologação (K8s) | postgres.staging.svc | keycloak.staging.svc |
| `prod` | Produção (K8s) | postgres.prod.svc | keycloak.prod.svc |

### 8.4 Estratégia de Branching — Uma Branch por Sprint

> **🚫 Regra de ouro:** Toda implementação deste projeto (`PRJ-FIN-2026-0003-SAAS-FBSO-ORG`) **DEVE** passar por uma `feature/sprint-NN-*`. **NUNCA** commitar diretamente em `main`.

**Decisão:** Cada sprint tem sua própria branch de vida curta (~2 semanas). Ao final da sprint, a branch é mergeada no `main` via PR e deletada. Esta estratégia substitui o modelo anterior de branch única (descontinuado em 16/07/2026).

**Justificativa:**
- **Segurança:** branches de 2 semanas têm risco muito menor de perda/sobrescrita que uma branch de 14 semanas
- **Isolamento:** problemas em uma sprint não afetam as demais
- **Entrega incremental:** `main` sempre contém o último estado validado e funcional
- **Clareza operacional:** `feature/sprint-03-portal-admin` é inequívoco — elimina confusão entre branches

#### Mapeamento Sprint → Branch

| Sprint | Branch | Marco | Status |
|:---|:---|:---|:---|
| Sprint 1–2 | `feature/java-fbso-platform-admin` | Setup + Segurança | ✅ Mergeada e deletada |
| Sprint 3 | `feature/sprint-03-portal-admin` | M2+M3 — Portal Admin + Contas/Planos | ✅ Mergeada e deletada |
| Sprint 4 | `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-04-rbac` | M4 — RBAC | 🔄 Ativa (Frente 0 ✅) |
| Sprint 5 | `feature/sprint-05-portal-cliente` | M5 — Portal Cliente | ⬜ Planejada — 4 frentes definidas |
| Sprint 6 | `feature/sprint-06-bus-catalogo` | M6 — BUs e Catálogo | ⬜ Pendente |
| Sprint 7 | `feature/sprint-07-homologacao` | M7 — Homologação | ⬜ Pendente |

#### Ciclo de Vida de Cada Sprint

```
CRIAR ─── DESENVOLVER ─── PR + REVIEW ─── MERGE NO MAIN ─── DELETAR
  ↑                                                              │
  └────────────────── próxima sprint ────────────────────────────┘
```

**Fluxo de trabalho:**

```bash
# === INÍCIO DA SPRINT ===
# 1. A partir de main, crie a branch da sprint
git checkout main && git pull
git checkout -b feature/sprint-NN-<slug>

# === DURANTE A SPRINT ===
# 2. Commits convencionais com prefixo da sprint
git add .
git commit -m "feat(sprint-03): implementa endpoint X (T-0XX)"
git push origin feature/sprint-NN-<slug>

# === FINAL DA SPRINT ===
# 3. Abra PR contra main
gh pr create --base main --head feature/sprint-NN-<slug> \
  --title "Sprint NN: <marco>" \
  --body "Entrega da Sprint NN conforme TASKS.md.

🤖 Generated with [Claude Code](https://claude.com/claude-code)"

# 4. Após merge, delete a branch
git branch -d feature/sprint-NN-<slug>
git push origin --delete feature/sprint-NN-<slug>
```

> ⚠️ **Antes de começar qualquer tarefa:** verifique com `git branch --show-current` se você está na branch correta da sprint. Commits em branches erradas serão rejeitados no code review.

#### Hotfix em Sprint Anterior

Se uma sprint já mergeada precisar de correção, crie uma branch de hotfix a partir do merge commit:

```bash
# Encontre o merge commit da sprint
git log --oneline --merges main | grep "Sprint"
# Crie branch de hotfix a partir desse ponto
git checkout -b hotfix/sprint-NN-<descricao> <merge-commit-hash>
```

---

> 📖 **Documento canônico completo:** [`docs/superpowers/specs/2026-07-16-sprint-branching-strategy-design.md`](../../../../../../../docs/superpowers/specs/2026-07-16-sprint-branching-strategy-design.md)

---

## 9. Referência Rápida para o Time Técnico

### 9.1 O que ler primeiro (ordem sugerida)

| Ordem | Documento | Tempo estimado |
|:---|:---|:---:|
| 1º | Este PRD.md | 10 min |
| 2º | [Project Charter](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Seções 1, 2, 3, 4, 7 | 15 min |
| 3º | [API-CONTRACTS.md](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/API-CONTRACTS.md) | 20 min |
| 4º | [ARCHITECTURE.md](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/ARCHITECTURE.md) | 20 min |
| 5º | [INTEGRATION-MAP.md](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/INTEGRATION-MAP.md) | 15 min |
| 6º | [04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Foco nas features do épico que for implementar | 30 min |
| 7º | [05-USER-STORIES-*.md](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/) — User stories do épico atual | 20 min |

**Total para onboard completo:** ~2 horas

---

## 10. Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.16 | 21/07/2026 | **GATE-PRD-SCOPE COMPLIANCE:** Validação em 5 dimensões (4 RESSALVAS, 1 APROVADO). 6 NCs corrigidas: BR-B05 Must Have (§4.6), out-of-scope expandido (§2), F04-03 classificação clarificada (§4.6), matriz de dependências (§4.7), premissas A1-A6 (§4.8), RN08-01/RN08-02 adicionadas (§6.6). Status: COMPLIANCE. | Agente GATE-PRD-SCOPE/IA |
| 1.15 | 17/07/2026 | Sprint 5 Frente 0 concluida: docker-compose (Keycloak+PG+MailHog), Flyway 10→12.11.0, PG driver→42.7.11 (CVE fix), OAuth2 Client adicionado, SecurityConfig refatorado com 2 filter chains. 2 NO-OP (DT-096/DT-098 ja implementados). Build ✅, 213 testes. | Agente IA |
| 1.14 | 17/07/2026 | Sprint 5 planejada: auditoria técnica com 9 skills identificou 42 débitos ([IDENTIFIED-TECHNICAL-DEBT](./sprints/sprint-05-portal-cliente/IDENTIFIED-TECHNICAL-DEBT-sprint-05-portal-cliente.md)). Time decidiu tratar 24 débitos na Sprint 5 em 3 frentes (6🔴 Frente 0 pré-sprint + 10🟡 Frente 1 + 8🔵 Frente 2) + 12 tarefas de features (Frente 3). 16 débitos postergados Sprints 6-7. Stack atualizado: Flyway 12.11.0, PostgreSQL driver 42.7.11, spring-boot-starter-oauth2-client adicionado. | Agente IA |
| 1.12 | 17/07/2026 | **Sprint 4 Frente 0 concluída:** Stack atualizado (Caffeine Cache, REST Assured). §5.3 Estrutura de pacotes reflete código real (entity/, aspect/, enums/). §6.1 Migrations: V004 (seed RBAC) + V006 (FK). §6.2 RBAC reescrito: DB-backed via PermissionService + ResourceAction/RoleResource (não JWT). ADR-08 com FORCE ROW LEVEL SECURITY. §8.4 Branch status: Sprint 3 mergeada, Sprint 4 ativa. [Detalhes](sprints/sprint-04-rbac/SPRINT-4-EXECUTION-REPORT-Frente-0.md) | Agente IA |
| 1.8 | 16/07/2026 | Atualização de stack (Spring Boot 3.5.14, Jackson 2.21.4 — CVE-2026-22733/CVE-2026-22731 auth bypass CVSS 8.2). Adicionada referência a débitos técnicos da Sprint 3 ([IDENTIFIED-TECHNICAL-DEBT](sprints/sprint-03-portal-admin/IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md) — auditoria com 7 skills). | Time Técnico |
| 1.5 | 16/07/2026 | Sprint 3 iniciada (16/07/2026). Status atualizado para "Em Execução". | Time Técnico |
| 1.7 | 16/07/2026 | Estratégia de branching alterada: modelo de branch única substituído por uma branch por sprint (§8.4). Adicionada tabela de mapeamento Sprint→Branch para Sprints 3–7. Documento canônico: `docs/superpowers/specs/2026-07-16-sprint-branching-strategy-design.md`. | Time Técnico |
| 1.6 | 16/07/2026 | Adicionada seção 8.4 — Branch de Desenvolvimento (`feature/java-fbso-platform-admin`) como obrigatória para todas as tarefas do projeto. (Substituída pela v1.7.) | Time Técnico |
| 1.4 | 15/07/2026 | Revisão Caveman (DOCS-SERVICE-CAVEMAN-REVIEW.md): Adicionado AuditEntry à tabela de entidades (§4.1, 10→11). Corrigida contagem ADR (8→7, §5.2). Corrigido RLS "11 tabelas"→"5 tabelas" (§6.1). Atualizadas versões dos artefatos em §3.3. | Caveman/IA |
| 1.2 | 14/07/2026 | Adicionado PostgreSQL Row-Level Security (RLS) como camada 1 de defesa em profundidade no isolamento multi-tenant (§6.1). Novo ADR-08 (§5.2). Atualizada sequência de segurança (§7.1) com migration V003 para RLS. Estratégia de 3 camadas: RLS (banco) + BaseRepository (aplicação) + Teste de Isolamento (detecção). | Agente Arquiteto/IA |
| 1.1 | 14/07/2026 | Correção pós-gate: 7 não-conformidades resolvidas do PRD_SCOPE_FAIL_REPORT.md v1.0. Adicionadas seções: §4.4 Matriz de Rastreabilidade, §4.5 Cobertura de Features/US, §4.6 Cobertura de BRs, §6.6 Regras de Negócio (18 RNs), §6.7 Cobertura de NFRs (8 NFRs). Atualizada §7.2 com referência explícita a critérios de aceite por US. NCs resolvidas: NC-001 a NC-007. | Agente Corretor PRD/IA |
| 1.0 | 13/07/2026 | Criação inicial: guia de entrada para time técnico, referência aos documentos oficiais, escopo do backend, ADRs, estrutura de pacotes, sequência de desenvolvimento | Time Técnico |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 030-architecture-adr-general, agile-ba-practices. v1.5 em 16/07/2026: Sprint 3 iniciada.*
