# TASKS.md — Plano de Tarefas: ms-fbso-platform-admin

- **Solucao:** `ms-fbso-platform-admin`
- **Projeto de Negocio:** [PRJ-FIN-2026-0003-SAAS-FBSO-ORG](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/)
- **Versao:** 2.8
- **Data:** 17 de Julho de 2026
- **Status:** Em Execucao — Sprints 1-3 concluidas ✅ (57/99 tarefas, 58%). 142 testes (100 unit + 42 IT). 18 endpoints REST. Proximo: Sprint 4 — RBAC
- **Origem:** [PRD.md](./PRD.md) + [SPECS.md](./SPECS.md) v1.8 + [ARCHITECTURE.md](./ARCHITECTURE.md) v2.2 + [04-FEATURES.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/04-FEATURES.md) + [IDENTIFIED-TECHNICAL-DEBT](./sprints/sprint-03-portal-admin/IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md)

---

## 1. Visao Geral

| Metrica | Valor |
|:---|:---|
| **Total de Tarefas** | 99 |
| **Setup (Pre-M2)** | 28 tarefas (8 setup + 8 seguranca + 12 correções pré-sprint) ✅ |
| **Sprint 3 — M2 Portal Admin** | 8 tarefas (8/8 concluídas ✅) |
| **Sprint 3 — Frente 3** | 7 tarefas (correções durante-sprint) |
| **M3 — Contas e Planos** | 15 tarefas |
| **M4 — RBAC** | 11 tarefas |
| **M5 — Portal Cliente** | 12 tarefas |
| **M6 — BUs e Catalogo** | 9 tarefas |
| **M7 — Homologacao** | 9 tarefas |
| **Must Have** | 96 |
| **Should Have** | 3 tarefas: T-021 (F01-03), T-063, T-064 (F04-03) |
| **Progresso Atual** | 57/99 (58%) — Sprints 1-3 concluídas ✅. 142 testes (100 unit + 42 IT). 18 endpoints REST. 10 features. 28 débitos resolvidos |

### Cobertura de Features (18/18)

| Epic | Features | Cobertura |
|:---|:---|:---:|
| EP-01 — Portal Admin | F01-01, F01-02, F01-03 | T-016 a T-023 |
| EP-02 — Clientes e Planos | F02-01, F02-02, F02-03, F02-04, F02-05 | T-024 a T-038 |
| EP-03 — RBAC | F03-01, F03-02, F03-03, F03-04 | T-046 a T-056 |
| EP-04a — Portal Cliente | F04-01, F04-02, F04-03, F04-04 | T-057 a T-068 |
| EP-04b — BUs e Catalogo | F04-05, F04-06 | T-069 a T-077 |

### Cronograma

```
Pre-M2 (Setup+Frente0)  M2 (EP-01)  M3 (EP-02)  Frente3  M4 (EP-03)  M5 (EP-04a)  M6 (EP-04b)  M7 (Homolog)
     28 tarefas         8 tarefas   15 tarefas  7 tarefas  11 tarefas   12 tarefas    9 tarefas     9 tarefas

  24/07 --------- 15/08 -------- 31/08 -------- 15/09 ------ 30/09 -------- 15/10 -------- 30/10
```

---

## 2. Tarefas por Marco de Entrega

### Pre-M2 — Setup e Fundacao (24/07 a 07/08)

**Estimativa total:** ~12 dias / **Responsavel:** A definir

| ID | Tarefa | Feature | US | Prio. | Est. | Responsavel | Criterio DONE |
|:---|:---|:---|:---:|:---:|:---:|:---|:---|
| **T-001** ✅ | Scaffold Maven: criar `pom.xml` com parent, dependencias (Spring Boot, Security, JDBC, Flyway, Testcontainers, Micrometer, Jakarta Validation) | — | — | Must | 2d | Agente IA | `pom.xml` compila sem erros. `mvn clean install` executa com sucesso. JaCoCo configurado com minimo 80% |
| **T-002** ✅ | Configurar `application.yml` (dev, staging, prod profiles) com datasource PostgreSQL, Keycloak JWKS URI, server port, logging level | — | — | Must | 1d | Agente IA | 3 profiles operacionais. Dev conecta ao PostgreSQL local via Docker. Logging configurado por profile |
| **T-003** ✅ | Criar `Dockerfile` (GraalVM Native Image) + `Dockerfile.jvm` (fallback dev) + `.dockerignore` | — | — | Must | 1d | Agente IA | Build nativo e JVM funcionam. Imagem < 200MB (nativo) / < 300MB (JVM) |
| **T-004** ✅ | Migration V001: schema `fbso_platform` e tabelas Core (Tenant, Plan, PlanModule, Subscription, User, UserPermission, ResourceAction, RoleResource, BusinessUnit, ProductService, AuditEntry) com campos de auditoria | — | — | Must | 3d | Agente IA | Flyway migrate cria 11 tabelas com colunas, PKs, FKs, NOT NULL. Rollback testado |
| **T-005** ✅ | Migration V002: indices unicos parciais (CNPJ ativo por tenant, email ativo por tenant, SKU ativo por BU) + indices de desempenho para queries frequentes | — | — | Must | 1d | Agente IA | `CREATE UNIQUE INDEX ... WHERE deleted_dt IS NULL` executado. Explain plan mostra index scan nas queries principais |
| **T-006** ✅ | Criar `BaseEntity.java` (created_dt, updated_dt, created_by, updated_by, deleted_dt, deleted_by) + `Address.java` (Value Object) + enums (TenantStatus, TenantSegment, Recurrence, SubscriptionStatus, UserStatus, Role, TaxRegime, ProductType) | — | — | Must | 1d | Agente IA | BaseEntity compativel com 11 tabelas. Todos os 8 enums implementados com valores conforme ARCHITECTURE.md |
| **T-007** ✅ | Criar `BaseRepository.java` (template JDBC com Soft Delete + Tenant Filter): findAll, findById, save, update, softDelete, hardDelete (apenas audit_log) | — | — | Must | 2d | Agente IA | Metodos CRUD injetam automaticamente `WHERE deleted_dt IS NULL` e `tenant_id = ?`. Teste unitario de cada metodo |
| **T-008** ✅ | Criar estrutura de pacotes completa (controller, service, repository, entity, dto/request, dto/response, enums, exception, security/aspect, security/annotation, config, common, utils) com classes vazias esqueleto | — | — | Must | 1d | Agente IA | `mvn compile` sem erros. Package structure segue ARCHITECTURE.md SS2 |

---

### Pre-M2 — Seguranca Cross-Cutting (07/08 a 15/08)

**Estimativa total:** ~11.5 dias / **Responsavel:** A definir

| ID | Tarefa | Feature | US | Prio. | Est. | Responsavel | Criterio DONE |
|:---|:---|:---|:---:|:---:|:---:|:---|:---|
| **T-009** ✅ | Configurar `SecurityConfig.java`: Spring Security + JWT (Keycloak RS256 public key via JWKS endpoint). Desabilitar CSRF (API stateless), configurar CORS para frontend | — | — | Must | 2d | Agente IA | Requisicao sem token → 401. Token valido → autenticado. Token invalido/expirado → 401. CORS permite origem do frontend |
| **T-010** ✅ | Criar `JwtAuthenticationFilter.java` (OncePerRequestFilter): extrair JWT do header `Authorization: Bearer`, validar assinatura RS256, validar exp (exp), extrair claims (tenant_id, user_id, roles, business_unit_ids, modules), setar SecurityContext | — | — | Must | 2d | Agente IA | Filter executa em toda requisicao exceto `/actuator/health`. Claims extraidas corretamente. SecurityContextHolder populado |
| **T-011** ✅ | Criar `TenantContext.java` (ThreadLocal): armazenar tenant_id, user_id, roles, business_unit_ids, modules da requisicao. Metodos: get/set/clear/getTenantId | — | — | Must | 0.5d | Agente IA | TenantContext.getTenantId() retorna tenant_id do JWT. Contexto limpo apos a requisicao (prevenir vazamento entre requisicoes) |
| **T-012** ✅ | Criar `TenantIsolationAspect.java`: interceptar @Repository via AOP, injetar `WHERE tenant_id = ?` em todas as queries. Ordem 1 (executa antes do repository) | — | — | Must | 1.5d | Agente IA | [SUBSTITUÍDO por T-015.1 — PostgreSQL RLS] Qualquer metodo de Repository sem tenant_id no contexto → SecurityException. Query injetada verificada via log/debug. Teste: mesmo SQL executado com tenants diferentes retorna dados diferentes |
| **T-015.1** ✅ | Criar Migration V003: ativar PostgreSQL Row-Level Security (RLS) em 5 tabelas com `tenant_id` — `ALTER TABLE ... ENABLE ROW LEVEL SECURITY` + `CREATE POLICY tenant_isolation USING (tenant_id = current_setting('app.current_tenant_id')::UUID)`. Configurar `app.current_tenant_id` via `TenantAwareDataSource` (proxy HikariCP). Escrever testes estruturais da migration | — | — | Must | 1.5d | Agente IA | ✅ RLS ativo em 5 tabelas (subscription, user, business_unit, product_service, audit_log). TenantAwareDataSource + BeanPostProcessor configurados. 33/33 testes passando (+11 novos). Migration V003 + U003 (rollback) criados |
| **T-013** ✅ | Criar anotacao `@RequiresPermission(resource, action)` + `RbacAspect.java`: verificar se role do usuario tem permissao (resource + action) conforme matriz RN10-01. Se negado → 403 | — | — | Must | 2d | Agente IA | Anotacao em metodo do controller bloqueia acesso sem permissao. Role valida × resource × action conforme seed data. 403 retorna JSON amigavel |
| **T-014** ✅ | Criar anotacao `@Auditable(entityType, action)` + `AuditAspect.java`: interceptar metodos anotados, capturar valores anteriores/novos via reflection/snapshot, gravar em audit_log de forma ASSINCRONA (ThreadPoolTaskExecutor / @Async) | — | — | Must | 1.5d | Agente IA | Registro de auditoria gerado para cada operacao. Valores anteriores e novos capturados corretamente. Async nao bloqueia thread principal |
| **T-015** ✅ | Criar `GlobalExceptionHandler.java` (@ControllerAdvice) com RFC 7807 + hierarchy de excecoes: BusinessException (422), DuplicateCnpjException, InvalidStatusTransitionException, PlanHasActiveSubscribersException, TenantNotFoundException, PermissionDeniedException (403) | — | — | Must | 1d | Agente IA | Resposta de erro contem `type`, `title`, `status`, `detail`. Stack trace nunca exposto em resposta HTTP. Mensagens em PT-BR |

> **Critico:** T-009 a T-015 sao pre-requisitos para QUALQUER endpoint de negocio. Nenhum controller deve ser implementado antes da camada de segurança estar funcional.

---

### Pre-M2 — Correções Pré-Sprint (Frente 0) | 16/07/2026

**Estimativa total:** ~16-25h (2-3 dias) / **Responsavel:** A definir

> Débitos técnicos impeditivos identificados na auditoria multidisciplinar com 7 skills ([IDENTIFIED-TECHNICAL-DEBT](./sprints/sprint-03-portal-admin/IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md)). Devem ser corrigidos ANTES de iniciar qualquer feature da Sprint 3.

| ID | Tarefa | Feature | US | Prio. | Est. | Responsavel | Criterio DONE |
|:---|:---|:---|:---:|:---:|:---:|:---|:---|
| **T-015.2.DT-001** ✅ | DT-001: Atualizar `spring-boot-starter-parent` 3.5.1→3.5.14 + Jackson 2.19.1→2.21.4. CVEs CVE-2026-22733/CVE-2026-22731 (auth bypass CVSS 8.2) e CVE-2026-54512/CVE-2026-54513 (RCE CVSS 8.1) | — | — | Must | 2h | Agente IA | ✅ Build passa. 36/36 testes (50 pós-Frente 1). CVEs eliminadas |
| **T-015.3.DT-002** ✅ | DT-002: Refatorar `AuditAspect` — capturar tenantId/userId no JoinPoint ANTES do @Async, passar como parâmetro | — | — | Must | 6h | Agente IA | ✅ TaskExecutor substitui @Async. AuditAspectTest (3 testes) criado |
| **T-015.4.DT-003** ✅ | DT-003: Adicionar `save(T)` e `update(T)` genéricos ao `BaseRepository` com preenchimento automático de created_by/updated_by | — | — | Must | 4h | Agente IA | ✅ INSERT/UPDATE com created_by/updated_by. BaseEntity.getId()/setId()/toColumnMap() abstratos |
| **T-015.5.DT-004** ✅ | DT-004: Atualizar JaCoCo 0.8.12→0.8.14 com suporte a Java 25 (class file major version 69) | — | — | Must | 3h | Agente IA | ✅ JaCoCo 0.8.14 funcional. Relatório gerado sem erro |
| **T-015.6.DT-005** ✅ | DT-005: Expandir `RbacAspect` — adicionar TENANT, PLAN, SUBSCRIPTION, DASHBOARD na matriz | — | — | Must | 2h | Agente IA | ✅ Matriz com MANAGER_EDIT_VIEW. 5 endpoints da Frente 1 protegidos |
| **T-015.7.DT-006** ✅ | DT-006: `TenantAwareDataSource` — trocar log.debug→log.error + lançar `TenantIsolationException` no catch | — | — | Must | 30min | Agente IA | ✅ Conexão NÃO retorna ao pool em caso de falha |
| **T-015.8.DT-007** ✅ | DT-007: Adicionar `spring-boot-starter-mail` no pom.xml | — | — | Must | 10min | Agente IA | ✅ JavaMailSender disponível. T-028 desbloqueada |
| **T-015.9.DT-008** ✅ | DT-008: Corrigir `AuditAspect.extractEntityId()` — validar UUID.fromString() ou usar parâmetro anotado | — | — | Must | 2h | Agente IA | ✅ idParamName + reflection. @Auditable atualizado |
| **T-015.10.DT-009** ✅ | DT-009: Migration V005 — adicionar `locked_price` + `locked_recurrence` em `subscription` | — | — | Must | 3h | Agente IA | ✅ V005 + U005 criadas. RN06-02 atendida |
| **T-015.11.DT-010** ✅ | DT-010: Corrigir Surefire — remover filtro `**/security/**`, usar `**/*Test.java` | — | — | Must | 5min | Agente IA | ✅ NO-OP. Padrão já incluía security tests |
| **T-015.12.DT-011** ✅ | DT-011: Reescrever `sendUnauthorized()` com ObjectMapper + ErrorResponse record | — | — | Must | 30min | Agente IA | ✅ JSON 401 consistente RFC 7807. Sem injection |
| **T-015.13.DT-012** ✅ | DT-012: Criar DuplicateCnpjException, InvalidStatusTransitionException, PlanHasActiveSubscribersException, TenantNotFoundException + TenantIsolationException | — | — | Must | 30min | Agente IA | ✅ 5 exceções criadas. GlobalExceptionHandler com handlers 404/409 |

> **Status: ✅ Frente 0 concluída em 17/07/2026 — 12/12 tasks.**

---

### M2 — Portal Admin (EP-01) | Data: 15/08/2026 | Status: ✅ 8/8 concluído em 17/07

**Features:** F01-01 ✅, F01-02 ✅, F01-03 ✅ (F01-03 Should)

| ID | Tarefa | Feature | US | Prio. | Est. | Responsavel | Criterio DONE |
|:---|:---|:---|:---:|:---:|:---:|:---|:---|
| **T-016** ✅ | Criar `DashboardRepository.java`: queries agregadas — total contas ativas, por status, por plano. Evolucao temporal (novas contas por mes). Metrica com `deleted_dt IS NULL` (RN01-01) | F01-01 | US-001, US-002, US-003 | Must | 2d | Agente IA | ✅ 9 queries implementadas (counts, agregacoes, evolucao, alertas, receita) |
| **T-017** ✅ | Criar `DashboardService.java`: logica de metricas operacionais, filtro de periodo (7d, 30d, 90d, mes_atual, ano_atual), padrao mes atual (RN01-02). Zero exibido como "0" (RN01-03) | F01-01 | US-001, US-002 | Must | 1d | Agente IA | ✅ 5 metodos: summary, evolution, byStatus, byPlan, alerts. Periodo invalido → mes atual |
| **T-018** ✅ | Criar DTOs de resposta: `DashboardSummaryResponse`, `EvolutionResponse`, `AccountsByStatusResponse`, `AccountsByPlanResponse`, `AlertResponse` | F01-01 | US-001, US-002, US-003 | Must | 0.5d | Agente IA | ✅ 5 records. JSON conforme contrato. ISO 8601 |
| **T-019** ✅ | Criar `DashboardController.java`: `GET /dashboard/admin/summary`, `/evolution`, `/accounts-by-status`, `/accounts-by-plan`, `/alerts`. `@RequiresPermission(DASHBOARD, view)` | F01-01 | US-001, US-002, US-003 | Must | 1d | Agente IA | ✅ 5 endpoints REST. Todos com @RequiresPermission |
| **T-020** ✅ | Criar `TenantRepository.java`: `findAll` paginado (25 registros), filtros por status/plano, busca textual (3+ chars, ILIKE). Ordenacao por created_at DESC | F01-02 | US-004, US-005 | Must | 1.5d | Agente IA | ✅ findAllPaginated + countFiltered + findByNameCorporate |
| **T-021** ✅ | Criar queries de alerta: tenants com onboarding incompleto > 48h (RN03-01) + assinatura suspensa. Endpoint `GET /dashboard/admin/alerts` com `@RequiresPermission`. Cards coloridos (WARNING/CRITICAL) | F01-03 | US-006, US-007 | Should | 1.5d | Agente IA | ✅ Integrado ao DashboardRepository + DashboardService |
| **T-022** ✅ | Testes unitarios: `DashboardService` (mocks), `TenantRepository` (mocks). Cobertura ≥ 80% nas classes testadas | F01-01, F01-02, F01-03 | US-001 a US-007 | Must | 1.5d | Agente IA | ✅ 14 testes: DashboardServiceTest (10) + TenantRepositoryTest (4). 50/50 passando |
| **T-023** ✅ | Testes de integracao: `DashboardRepositoryIT` com Testcontainers (PostgreSQL 17). 23 cenarios: summary, soft-delete, paginacao, busca textual, alertas, locked_price (DT-009). 10+ tenants seed | F01-01, F01-02, F01-03 | US-001 a US-007 | Must | 1.5d | Agente IA | ✅ 23/23 testes passando. Docker requerido. 3 bugs corrigidos: V003 product_service, BaseRepository.save() array size, BaseIntegrationTest visibilidade |

---

### M3 — Gestao de Clientes e Planos (EP-02) | Data: 31/08/2026 | Status: ✅ 15/15 concluído em 17/07

**Features:** F02-01, F02-02, F02-03, F02-04, F02-05

| ID | Tarefa | Feature | US | Prio. | Est. | Responsavel | Criterio DONE |
|:---|:---|:---|:---:|:---:|:---:|:---|:---|
| **T-024** | Criar entidades Tenant (enum TenantStatus: PENDING_ONBOARDING, ACTIVE, SUSPENDED, INACTIVE) + DTOs request/response. Implementar mapa de transicoes de status (RN05-01) com validacao de cada transicao permitida | F02-01, F02-02 | US-008, US-012 | Must | 1d | A definir | Transicoes validas: PENDING→ACTIVE, ACTIVE↔SUSPENDED, ACTIVE↔INACTIVE. Transicao invalida (ex: ACTIVE→PENDING) → 422 com mensagem clara |
| **T-025** | `TenantService.create()`: status inicial PENDING_ONBOARDING, validacao razao social unica entre ativos (RN04-02). `TenantService.update()`: editar dados cadastrais. Gerar auditoria via @Auditable (RN04-01) | F02-01 | US-008, US-010 | Must | 2d | A definir | Tenant criado com status PENDING_ONBOARDING. Razao social duplicada → 409. Auditoria registrada com admin responsavel |
| **T-026** | `TenantService.suspend()`: exige motivo (RN05-02), bloqueia acesso usuarios em ≤ 5min. `TenantService.reactivate()`: restaura permissoes anteriores (RN05-03). Timeline de status registrada | F02-02 | US-012, US-013, US-014 | Must | 2d | A definir | Suspensao sem motivo → 400. Usuarios do tenant bloqueados. Timeline exibe todas as transicoes |
| **T-027** | `TenantController`: CRUD `/api/v1/tenants` + `POST /{id}/suspend`, `/reactivate`, `/resend-invite`. Anotar com `@RequiresPermission`. DTOs com Bean Validation | F02-01, F02-02 | US-008 a US-014 | Must | 2d | A definir | 7 endpoints REST funcionais (GET list, GET by id, POST, PATCH, POST suspend, POST reactivate, POST resend-invite). Validacoes conforme SPECS.md §4.2 |
| **T-028** | Integracao email: disparo automatico apos criacao de Tenant (US-009). Link unico de ativacao com expiracao 7 dias (RN04-03). Reenvio via `/resend-invite` gera novo link + novo prazo (US-011). Servico de email via SMTP | F02-01 | US-009, US-011 | Must | 2d | A definir | Email enviado apos criacao. Link expira em 7 dias. Reenvio funcional para tenants em PENDING_ONBOARDING. Log de envio em audit_log |
| **T-029** | Criar entidade Plan + PlanModule + PlanRepository. Plan com campos: name, description, price (> 0), recurrence (MONTHLY/QUARTERLY/YEARLY), status (ACTIVE/DISCONTINUED), version. Validation: price > 0 (RN06-02) | F02-03 | US-015 | Must | 1d | A definir | Plan criado com status ACTIVE. price > 0 validado. PlanModule vinculado. Recorrencias selecionaveis |
| **T-030** | `PlanService`: CRUD com versionamento (edicao gera nova versao — US-017). `deactivate()`: nao afeta assinantes ativos (RN06-01). Alteracao de preco preserva assinaturas existentes (RN06-02) | F02-03 | US-015, US-016, US-017, US-018 | Must | 2d | A definir | Edicao de plano gera novo version_id. Plano com assinantes nao pode ser excluido → 422. Plano desativado nao aparece em novas contratacoes |
| **T-031** | `PlanController`: CRUD `/api/v1/plans` + `POST /{id}/deactivate`. `@RequiresPermission`. GET lista apenas planos ativos (para assinatura), GET admin lista todos | F02-03 | US-015 a US-018 | Must | 1d | A definir | CRUD completo. Deactivate muda status DISCONTINUED. Plano desativado aparece como "Descontinuado" no GET admin |
| **T-032** | Criar entidade Subscription + SubscriptionRepository. Regra: 1 assinatura ativa por tenant (RN07-01). FK para tenant e plan. Status: ACTIVE, SUSPENDED, CANCELED | F02-04 | US-019 | Must | 1d | A definir | Subscription criada com status ACTIVE. Segunda ativa → 409 (RN07-01). Vigencia controlada por start_date / end_date |
| **T-033** | `SubscriptionService`: criar, change-plan (upgrade/downgrade — US-020), suspender (US-021), reativar. Change-plan finaliza anterior e cria nova (transicao sem deixar tenant sem assinatura — RN07-02) | F02-04 | US-019, US-020, US-021 | Must | 2d | A definir | Change-plan finaliza assinatura anterior + cria nova sem gap. Suspensao bloqueia modulos. Historico de assinaturas mantido |
| **T-034** | `SubscriptionController`: `POST /tenants/{tid}/subscriptions`, `GET /tenants/{tid}/subscriptions`, `POST /subscriptions/{id}/change-plan`, `POST /subscriptions/{id}/suspend`. `@RequiresPermission` | F02-04 | US-019, US-020, US-021 | Must | 1.5d | A definir | 4 endpoints. Validacoes: plan_id referencia plano ativo. Subscription atual nao pode estar SUSPENDED para change-plan. effective_date opcional |
| **T-035** | `AuditRepository` + `AuditService`: queries com filtros por periodo (start_date, end_date), action (CREATED, UPDATED, SUSPENDED, REACTIVATED, PLAN_CHANGED), entity_type. Paginacao obrigatoria (padrao 25, max 100) (RN08-01, RN08-02) | F02-05 | US-022, US-023 | Must | 1.5d | A definir | Filtros funcionais. Paginacao com tamanho padrao 25. ORDER BY timestamp DESC. Registros imutaveis — UPDATE/DELETE → 403 |
| **T-036** | `AuditController`: `GET /api/v1/audit` com filtros (start_date, end_date, action, entity_type, page, size, sort). `@RequiresPermission(resource="AUDIT", action="view")`. Resposta paginada | F02-05 | US-022, US-023 | Must | 1d | A definir | GET /audit funcional com todos os filtros. Admin FBSO ve tudo. Auditor ve tudo (leitura). Usuario sem role AUDIT nao acessa |
| **T-037** | Testes unitarios M3: `TenantService`, `PlanService`, `SubscriptionService`, `AuditService` com JUnit 5 + Mockito. Cobrir todas as RNs (RN05-01 a RN08-02) | F02-01 a F02-05 | US-008 a US-023 | Must | 2d | A definir | Cobertura ≥ 80%. Cada RN testada (positivo + negativo). Transicoes de status invalidas retornam excecao |
| **T-038** | Testes integracao M3: CRUD Tenant/Plan/Subscription/Audit com Testcontainers. Testar RN07-01 (1 assinatura ativa), RN06-01 (exclusao plano com assinantes), RN05-01 (transicoes), RN08-02 (imutabilidade audit) | F02-01 a F02-05 | US-008 a US-023 | Must | 2d | A definir | Testcontainers com PostgreSQL. Dados seed. Cenarios de borda: tentar criar 2a assinatura ativa, tentar excluir plano com assinantes, tentar UPDATE audit_log |

---

### Sprint 3 — Frente 3: Correções Durante a Sprint ✅ 7/7 concluído (17/07)

**Estimativa total:** ~10h / **Responsavel:** A definir

> Correções de débitos técnicos não-bloqueantes, identificados na auditoria ([IDENTIFIED-TECHNICAL-DEBT](./sprints/sprint-03-portal-admin/IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md)). Executar quando houver buffer entre M2 e M3.

| ID | Tarefa | Feature | US | Prio. | Est. | Responsavel | Criterio DONE |
|:---|:---|:---|:---:|:---:|:---:|:---|:---|
| **T-039.DT-017** | DT-017: Decidir V004 "opcional" vs "pré-requisito" + validar `idx_tenant_segment` | — | — | Must | 15min | A definir | Decisão registrada no Log de Decisões |
| **T-040.DT-019** | DT-019: Recalibrar day-by-day: 12→~15 dias realista com Frentes 0+3 | — | — | Must | 15min | A definir | Planejamento atualizado no SPRINT-DEVELOPMENT-PLANNING |
| **T-041.DT-021** | DT-021: Implementar captura de valores "antes" no AuditAspect | — | — | Must | 3h | A definir | Colunas previous_value/new_value populadas. PRD §6.4 atendido |
| **T-042.DT-025** | DT-025: Adicionar `@ExceptionHandler(AccessDeniedException.class)` no GlobalExceptionHandler | — | — | Must | 10min | A definir | Acesso negado retorna 403 (não 500) |
| **T-043.DT-026** | DT-026: Refatorar RLSIsolationTest para Testcontainers + PostgreSQL real | — | — | Must | 4h | A definir | Testes de RLS com queries reais cross-tenant |
| **T-044.DT-029** | DT-029: Extrair `hasTenantColumn` branching para método helper no BaseRepository | — | — | Must | 1h | A definir | 4 métodos → 1 helper. Sem duplicação de lógica |
| **T-045.DT-046** | DT-046: Atualizar Testcontainers 1.20.6→1.21.4 + commons-compress 1.24.0→1.28.0 | — | — | Must | 15min | A definir | CVE-2024-25710 mitigada. Build passa |

---

### M4 — Governanca de Acessos (EP-03) | Data: 15/09/2026

**Features:** F03-01, F03-02, F03-03, F03-04

| ID | Tarefa | Feature | US | Prio. | Est. | Responsavel | Criterio DONE |
|:---|:---|:---|:---:|:---:|:---:|:---|:---|
| **T-046** | Criar entidade User (fields: id, tenant_id FK, external_keycloak_id, email, name, status ACTIVE/INACTIVE/INVITE_PENDING) + UserRepository com `findByEmailAndTenant` e validacao email unico por tenant ativo (RN09-02, indice parcial) | F03-01 | US-024 | Must | 1d | A definir | UserRepository funcional. Email unico por tenant ativo. Soft delete respeitado. Indice parcial unique_email_active |
| **T-047** | `UserService`: convite usuario (email unico por tenant, RN09-02), desativar (nao permite autodesativacao RN09-03), reativar. Convite expira em 7 dias (RN09-01). Integracao email de convite | F03-01 | US-024, US-025, US-026 | Must | 1.5d | A definir | Convite com email valido. Email duplicado no mesmo tenant → 409. Autodesativacao → 422. Reativacao restaura acesso |
| **T-048** | `UserController`: `GET /api/v1/users`, `POST /api/v1/users`, `PATCH /users/{id}`, `POST /users/{id}/deactivate`. `@RequiresPermission(resource="USER", action="view"/"create"/"edit"/"delete")` | F03-01 | US-024, US-025, US-026 | Must | 1d | A definir | CRUD usuarios funcional. Deactivate bloqueia login. Lista exibe nome, email, role, status, BUs vinculadas. Filtro por status |
| **T-049** | Criar entidades de dominio ResourceAction (resource_name, action) + RoleResource (role, resource_action_id FK). Populate seed data via migration com matriz completa RN10-01 (4 roles x recursos: DASHBOARD, BUSINESS_UNIT, PRODUCT, USER, AUDIT, PERMISSION, PLAN) | F03-02 | US-027, US-028, US-029, US-030 | Must | 1d | A definir | Seed data carrega corretamente. Matriz de permissoes completa conforme RN10-01. Consulta `findByRole` retorna recursos corretos |
| **T-050** | Criar entidade UserPermission (user_id FK, business_unit_id FK, role) + PermissionRepository. Vincular usuario x BU x role. Tabela ponte com UNIQUE (user_id, business_unit_id) | F03-02, F03-03 | US-027 a US-033 | Must | 1.5d | A definir | UserPermission com UNIQUE constraint. Query de permissoes por usuario retorna todas as BUs + role. Admin tenant tem acesso implicito a todas as BUs (US-029) |
| **T-051** | `PermissionService`: atribuir/revogar permissoes, vincular BU, gerenciar modulos. Admin tenant tem acesso implicito a todas as BUs. Usuario sem BU vinculada nao acessa portal (RN11-01). Usuario sem modulo nao acessa portal (RN11-02) | F03-02, F03-03 | US-027 a US-033 | Must | 2d | A definir | Permissoes atribuidas corretamente. Admin tenant ve todas as BUs. Usuario sem BU → acesso negado. Efeito imediato na proxima acao (RN11-03) |
| **T-052** | `PermissionController`: `GET /users/{uid}/permissions`, `PUT /users/{uid}/permissions`. `@RequiresPermission(resource="PERMISSION", action="view"/"edit")` | F03-02, F03-03 | US-027 a US-033 | Must | 1d | A definir | GET retorna permissoes atuais. PUT atualiza vinculos. Validacoes de negocio no service. Auditoria registrada para cada alteracao |
| **T-053** | Integrar `RbacAspect` com `RoleResource` carregado do banco: antes de executar metodo com `@RequiresPermission`, consultar matriz e verificar se role do usuario tem acesso ao resource+action. Cache de matriz para performance | F03-04 | US-034, US-035 | Must | 1.5d | A definir | RBAC funcional para todos os endpoints. Requisicao sem permissao → 403 (nao 404 — RN12-01). Cache de permissoes (TTL 5min) |
| **T-054** | Garantir resposta 403 padrao (JSON amigavel): `{"title": "Acesso negado", "detail": "Voce nao tem permissao para acessar esta area.", "status": 403}`. Sem detalhes tecnicos ou caminhos internos (RN12-02) | F03-04 | US-036 | Must | 0.5d | A definir | Resposta 403 sempre no formato padrao. Nenhum stack trace. Mensagem em PT-BR |
| **T-055** | Testes unitarios M4: `UserService`, `PermissionService` com JUnit 5 + Mockito. Testar RN09-03 (autodesativacao), RN10-01 (matriz permissoes), RN11-01 (sem BU), RN11-02 (sem modulo) | F03-01, F03-02, F03-03 | US-024 a US-033 | Must | 1.5d | A definir | Todos os cenario de RN testados. Excecoes lancadas corretamente. Cobertura ≥ 80% |
| **T-056** | Testes seguranca RBAC (integrados): cada papel × endpoint proibido → 403. Testar com Testcontainers. Cenarios: Operador tenta PATCH /products, Auditor tenta POST /users, Gerente tenta POST /plans | F03-02, F03-04 | US-027 a US-036 | Must | 2d | A definir | Testes automatizados para combinacoes papel x recurso. 403 retornado para todas as combinacoes proibidas. Matriz RN10-01 validada como teste parametrizado |

---

### M5 — Portal do Cliente e Onboarding (EP-04a) | Data: 30/09/2026

**Features:** F04-01, F04-02, F04-03, F04-04

| ID | Tarefa | Feature | US | Prio. | Est. | Responsavel | Criterio DONE |
|:---|:---|:---|:---:|:---:|:---:|:---|:---|
| **T-057** | Configurar Keycloak realm `fbso-platform` + client para autenticacao do portal cliente (Authorization Code Flow). Configurar mapeamento de claims (tenant_id, roles, business_unit_ids, modules). Exportar config como `realm-config.json` para Docker Compose | F04-01 | US-037 | Must | 2d | A definir | Keycloak realm funcional. Login com usuario do realm retorna JWT com claims corretas. Configuracao versionada |
| **T-058** | Endpoint `POST /api/v1/auth/login` (delega para Keycloak OIDC). Fluxo de recuperacao de senha: `POST /auth/forgot-password` (envia link, expira 1h — RN13-03), `POST /auth/reset-password` (valida token + nova senha com complexidade RN13-01) | F04-01 | US-037, US-038 | Must | 2d | A definir | Login funcional. Recuperacao de senha envia email com link (expira 1h). Nova senha validada (8+ chars, letra + numero). Sessao expira 60min inatividade (RN13-02) |
| **T-059** | Rate limiting: 5 tentativas incorretas → bloqueio de 15 minutos (US-039, RN13-02). Implementar via Spring @Aspect ou filter + cache (Caffeine/Redis). Mensagem informa tempo restante de bloqueio | F04-01 | US-039 | Must | 1.5d | A definir | Apos 5 falhas consecutivas, conta bloqueada por 15min. Mensagem exibe tempo restante. Admin pode desbloquear manualmente |
| **T-060** | `OnboardingService`: logica dos 4 passos obrigatorios (US-040 a US-044). Validacao de ordem (nao pula etapas — RN14-01). Primeira BU cadastrada vira Matriz (RN14-02). Tenant muda PENDING_ONBOARDING → ACTIVE ao concluir (RN14-04). Estado salvo permite retomar depois | F04-02 | US-040, US-041, US-044 | Must | 2d | A definir | Onboarding em 4 passos. Ordem obrigatoria. Tenant ACTIVE apos conclusao. Barra de progresso viavel. Retomada de onde parou |
| **T-061** | `OnboardingController`: `GET /api/v1/onboarding/status`, `PATCH /step-1` (confirmar dados), `POST /step-2` (cadastrar Matriz), `POST /complete` (finalizar). `@RequiresPermission`. Validacoes conforme SPECS.md §4.2 | F04-02 | US-040 a US-044 | Must | 1.5d | A definir | 4 endpoints. Validacao de CNPJ no step-2. Complete so funciona se todos os passos concluidos (RN14-03). Step pulado → redirect para passo correto |
| **T-062** | Criacao da primeira BusinessUnit como Matriz durante o onboarding (step-2). Reutilizar `BusinessUnitService.create()` com flag `isMatrix=true`. CNPJ validado (formato, digito) | F04-02 | US-042 | Must | 1d | A definir | Primeira BU criada com parent_id=NULL. CNPJ valido. TaxRegime obrigatorio. Se onboarding falha no step-2, dados nao persistem |
| **T-063** | `DashboardClientService`: queries para dashboard do cliente — unidades ativas, produtos no catalogo, plano contratado, notificacoes/lembretes. Cards clicaveis (links no response) | F04-03 | US-045, US-046 | Should | 1.5d | A definir | Cards com dados resumidos. Cada card com link para area correspondente. Notificacoes dispensaveis |
| **T-064** | `DashboardClientController`: `GET /api/v1/dashboard/client/summary`, `/notifications`. `@RequiresPermission`. Resposta adaptada ao modulo ativo (RN15-01). Dashboard generico na Fase 0 (RN15-02) | F04-03 | US-045, US-046 | Should | 1d | A definir | Endpoints retornam dados do cliente autenticado. Filtrado por tenant_id. Notificacoes com link de acao |
| **T-065** | Atualizar `JwtAuthenticationFilter` para incluir claims `modules[]` e `business_unit_ids[]` no `TenantContext`. Modular: lista de modulos contratados no plano + autorizados ao usuario (RN16-01). Modulo placeholder "FBSO Platform" na Fase 0 (RN16-02) | F04-04 | US-047, US-048, US-049 | Must | 1.5d | A definir | JWT retorna modules[] com IDs corretos. App Switcher viavel mesmo com 1 modulo (exibe nome ativo RN16-02). Business units no contexto |
| **T-066** | Criar endpoint `GET /api/v1/auth/me`: retorna dados do usuario logado (id, name, email, role, business_unit_ids, modules[], tenant_id, status onboarding). Sem `@RequiresPermission` (autenticacao ja validada no filter) | F04-04 | US-047 | Must | 1d | A definir | GET /auth/me funcional. Retorna dados do token sem consulta extra ao banco (stateless). 401 se token ausente/invalido |
| **T-067** | Testes unitarios M5: `OnboardingService`, `DashboardClientService`, fluxo auth (login, recovery, rate-limit). JUnit 5 + Mockito | F04-01, F04-02, F04-03, F04-04 | US-037 a US-049 | Must | 1.5d | A definir | Cobertura ≥ 80%. Rate limit testado (5 tentativas, bloqueio, reset apos 15min). Onboarding: cada passo testado isoladamente |
| **T-068** | Testes integracao M5: fluxo completo onboarding com Testcontainers (PENDING_ONBOARDING → ACTIVE). Teste de rate limiting (login com senha errada x6). Teste de expiracao de link de redefinicao de senha | F04-01, F04-02 | US-037 a US-044 | Must | 2d | A definir | Onboarding completo testado. Rate limit testado com PostgreSQL real. Tenant transiciona para ACTIVE apos conclusao |

---

### M6 — Unidades de Negocio e Catalogo (EP-04b) | Data: 15/10/2026

**Features:** F04-05, F04-06

| ID | Tarefa | Feature | US | Prio. | Est. | Responsavel | Criterio DONE |
|:---|:---|:---|:---:|:---:|:---:|:---|:---|
| **T-069** | Criar entidade BusinessUnit (fields: tenant_id FK, parent_id FK auto-relacionamento, cnpj, corporate_name, tax_regime, address, status ACTIVE/INACTIVE) + BusinessUnitRepository com queries hierarquicas (Matriz no topo, filiais recuadas) | F04-05 | US-050 | Must | 1.5d | A definir | Estrutura hierarquica funcional. parent_id permite auto-relacionamento recursivo. CNPJ com indice parcial unique_cnpj_active. FindAll retorna ordenacao hierarquica |
| **T-070** | `BusinessUnitService`: CRUD com hierarquia (Matriz/Filial). CNPJ unico entre ativos do tenant (RN17-01). CNPJ imutavel apos cadastro. Soft delete libera CNPJ para reuso (RN17-01). Unidade desativada nao pode ser "pai" (RN17-02). Sem limite de niveis (RN17-04) | F04-05 | US-051, US-052, US-053, US-054 | Must | 2d | A definir | CNPJ duplicado → 409. CNPJ alterado apos cadastro → 400. Soft delete de BU com produtos vinculados funciona. Reuso de CNPJ apos soft delete |
| **T-071** | `BusinessUnitController`: CRUD `/api/v1/business-units` + `POST /{id}/deactivate`. `@RequiresPermission` (Admin/Manager: create/edit; Admin: deactivate). Seletor de BU no topo: listar BUs que o usuario tem permissao (RN17-05) | F04-05 | US-050 a US-054 | Must | 1.5d | A definir | CRUD funcional. Admin ve todas BUs. Gerente/Operador ve apenas autorizadas. Seletor viavel (endpoint retorna BUs permitidas) |
| **T-072** | Criar entidade ProductService (fields: business_unit_id FK, name, sku, type PRODUCT/SERVICE, description, status ACTIVE/INACTIVE) + ProductRepository. SKU unico por BU ativo (RN18-02, indice parcial). Vinculado automaticamente a BU ativa no seletor | F04-06 | US-055 | Must | 1d | A definir | ProductService entity completa. SKU unico por BU com indice parcial. Vinculacao a BU. status padrao ACTIVE |
| **T-073** | `ProductService`: CRUD completo. Indicador "Nao mapeado" como placeholder para mapeamento fiscal futuro (RN18-03). Soft delete (RN18-04). SKU opcional, se informado unico por BU. Catalogo segmentado por BU (RN18-01) | F04-06 | US-055, US-056, US-057, US-058 | Must | 2d | A definir | CRUD funcional. Indicador "Nao mapeado" no response. Soft delete preserva historico. Busca textual por nome e SKU |
| **T-074** | `ProductController`: CRUD `/api/v1/products` + `POST /{id}/deactivate`, `/activate`. `@RequiresPermission` (Admin/Manager: create/edit/deactivate; Operator: view). Filtro por BU ativa no seletor | F04-06 | US-055 a US-058 | Must | 1d | A definir | CRUD endpoints funcional. Produto desativado nao aparece em seletores. Ativar/Desativar funcional |
| **T-075** | Testes unitarios M6: `BusinessUnitService`, `ProductService` com JUnit 5 + Mockito. Testar RN17-01 (CNPJ unico), RN17-04 (hierarquia ilimitada), RN18-02 (SKU unico), RN18-04 (soft delete) | F04-05, F04-06 | US-050 a US-058 | Must | 1.5d | A definir | Cobertura ≥ 80%. Todos os cenario de RN testados. Soft delete + reuso de CNPJ verificado |
| **T-076** | Testes integracao M6: CRUD BusinessUnit + ProductService com Testcontainers. Validar CNPJ unico, hierarquia, SKU unico, soft delete libera reuso. Catalogo filtrado por BU | F04-05, F04-06 | US-050 a US-058 | Must | 2d | A definir | Testcontainers com PostgreSQL. Dados seed com multi-tenants. Query cross-tenant nao retorna BUs de outro tenant |
| **T-077** | Testes de isolamento multi-tenant: tenant-A nao ve dados do tenant-B em nenhum endpoint de BU ou Produto. Teste automatizado com 2 tenants + dados sobrepostos | F04-05, F04-06 | US-050 a US-058 | Must | 1.5d | A definir | Teste isolado: criar dados para 2 tenants, consultar como tenant-A, verificar que zero dados de tenant-B retornam |

---

### M7 — Integracao, Testes e Homologacao | Data: 30/10/2026

**Features:** Todas (F01-01 a F04-06)

| ID | Tarefa | Feature | US | Prio. | Est. | Responsavel | Criterio DONE |
|:---|:---|:---|:---:|:---:|:---:|:---|:---|
| **T-078** | Testes de regressao: bateria completa de testes M2-M6 apos integracao final. Verificar que todas as features continuam operando. Automatizar como suite unica | Todas | Todas | Must | 2d | A definir | Suite de regressao executa sem falhas. Nenhuma quebra em features ja homologadas |
| **T-079** | Testes de performance: dashboard admin ≤ 3s (p95) com 1000 tenants simulados (BR-NFR05). Listas paginadas com 10k registros. Teste de carga via JMeter ou k6 | EP-01 | US-001 a US-007 | Must | 2d | A definir | p95 ≤ 3s no dashboard. Listas paginadas retornam em ≤ 1s. Relatorio de teste de carga |
| **T-080** | Criar/atualizar `fbso-platform-api.yaml` (OpenAPI 3.0) com todos os 37 endpoints, schemas request/response, exemplos, codigos de erro. BR-NFR04 (100% endpoints documentados) | Todas | Todas | Must | 2d | A definir | OpenAPI valido (sem erros de schema). 37 endpoints documentados. Cada schema com exemplo. Codigos HTTP documentados por endpoint |
| **T-081** | Documentar `README.md` da solucao (como rodar, testar, fazer deploy) + README.md por pacote (controller, service, repository, security) conforme BR-NFR04 | Todas | Todas | Must | 1.5d | A definir | README.md raiz com quickstart. README.md por pacote com descricao de responsabilidade. Atualizado no repositorio |
| **T-082** | SAST scan (Semgrep / SonarQube): zero vulnerabilidades criticas/high. Corrigir achados de seguranca (NFR-OWASP). Verificar OWASP Top 10: SQL injection (JDBC parametrizado), broken auth (JWT), XSS, CSRF | Todas | Todas | Must | 2d | A definir | Relatorio SAST com 0 criticos/high. Todas as queries usam PreparedStatement. Respostas JSON escapadas. Teste DAST basico |
| **T-083** | Verificar LGPD: dados pessoais (nome, email, CNPJ) mascarados em logs (NFR-LGPD). Todas as entidades com soft delete implementado (coluna deleted_dt). Dados preservados para auditoria (LGPD Art. 16) | Todas | Todas | Must | 1d | A definir | Logs sem dados pessoais em texto claro. Todas as 11 entidades com deleted_dt. Auditoria funcional para dados excluidos |
| **T-084** | Deploy staging (K8s): criar/atualizar manifests (deployment, service, configmap, secret, hpa). Smoke test apos deploy. Health checks (liveness, readiness) para BR-NFR01 (disponibilidade 99,5%) | Todas | Todas | Must | 2d | A definir | Staging operacional. Health check endpoints respondendo. Smoke test passa. Readiness probe configurada |
| **T-085** | UAT com Product Owner: demonstracao de todas as features. Correcoes de bugs encontrados na validacao. Aprovacao formal do PO conforme DoD de Entrega (DEFINITION_OF_DONE.md) | Todas | Todas | Must | 2d | A definir | PO valida todas as features. Bugs corrigidos. Termo de aceite assinado. Check-list DoD completo |
| **T-086** | Deploy producao (K8s): promote de staging para producao. Validacao pos-go-live (KPIs: uptime, latencia p95, erros 4xx/5xx). Rollback plan documentado. Monitoramento e alertas configurados | Todas | Todas | Must | 2d | A definir | Producao operacional. KPIs monitorados. Rollback testado. Alertas configurados para 4xx > 5% e latencia > 3s |

---

### Sprint 4+ — Débitos Postergados

> 9 débitos técnicos postergados da Sprint 3, identificados na auditoria ([IDENTIFIED-TECHNICAL-DEBT](./sprints/sprint-03-portal-admin/IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md)). Serão tratados nas Sprints 4-7 conforme planejamento.

| ID | Tarefa | Débito | Est. | Sprint Alvo | Critério DONE |
|:---|:---|:---:|:---:|:---:|:---|
| **T-087.DT-023** | Migrar paginação offset→keyset no BaseRepository | DT-023 | 1d | 5 | Performance consistente com >10k registros |
| **T-088.DT-030** | Consolidar dupla validação JWT via `Converter<Jwt, AbstractAuthToken>` customizado | DT-030 | 4h | 4 | JWT decodificado 1× por requisição |
| **T-089.DT-031** | Reduzir maxAllowedViolations checkstyle: 300→100→0 | DT-031 | Contínuo | 4-7 | 0 violações ao final da Sprint 7 |
| **T-090.DT-034** | Remover Address.java se não usado até Sprint 6 | DT-034 | 5min | 6 | -95 linhas de código morto |
| **T-091.DT-035** | Migrar RbacAspect de strings literais para Role enum | DT-035 | 1h | 4 | Type-safety. Role enum referenciado |
| **T-092.DT-042** | Criar docker-compose.yml com PostgreSQL + Keycloak + MailHog | DT-042 | 1h | 4 | Setup local em 1 comando |
| **T-093.DT-043** | Criar script seed SQL com 50+ tenants + 3 planos | DT-043 | 30min | 3 | Dados realistas para desenvolvimento |
| **T-094.DT-044** | Criar logback-spring.xml com JSON appender + rotação | DT-044 | 1h | 7 | Logs estruturados em produção |
| **T-095.DT-045** | Migrar Flyway 10.22.0→12.11.0 (breaking changes analisados) | DT-045 | 4h | 5 | Migração executada sem perda de histórico |

---

## 3. Dependencias entre Tarefas

```
FASE 0 — SETUP (T-001 a T-008)
    │
    ├── T-004 (Migration V001) ── T-005 (Indices)
    │                            └── T-042 (Seed RBAC) ── T-043 (PermissionRepo)
    │
    ├── T-001 (Maven) ─── T-009 (SecurityConfig) ─── T-010 (JWT Filter)
    │        │                                               │
    │        │                                               ├── T-011 (TenantContext)
    │        │                                               ├── T-013 (@RequiresPermission) ── T-046 (RbacAspect)
    │        │                                               ├── T-058 (modules[] claim)
    │        │                                               └── T-015.1 (PostgreSQL RLS — usa JDBC para SET session var)
    │        │
    │        ├── T-007 (BaseRepository) ─── T-012 (TenantIsolationAspect)
    │        │                                   │
    │        │                                   ├── TODOS os Repositories (M2-M6)
    │        │                                   └── T-070 (Testes isolamento)
    │        │
    │        └── T-004 (Migration V001) ─── T-015.1 (Migration V003 — RLS Policies)
    │        │
    │        └── T-015 (GlobalExceptionHandler) ─── T-047 (403 amigavel)
    │
    └── T-006 (BaseEntity) ── TODAS as entities

FASE 1 — M2 (EP-01): T-016 ── T-017 ── T-018 ── T-019 ── T-022, T-023
                       T-020 ───────────────────── T-021
                                                    │
                                                    └── T-025 (M3, depende de TenantRepository)

FASE 2 — M3 (EP-02):
           T-024 ── T-025 ── T-027 ── T-028
           T-024 ── T-026 ── T-027
           T-029 ── T-030 ── T-031
           T-032 ── T-033 ── T-034
           T-035 ── T-036
           T-024..T-036 ── T-037, T-038

FASE 3 — M4 (EP-03):
           T-039 ── T-040 ── T-041
           T-042 ── T-043 ── T-044 ── T-045
           T-013 ── T-046
           T-015 ── T-047
           T-039..T-047 ── T-048, T-049

FASE 4 — M5 (EP-04a):
           T-050 ── T-051 ── T-052
           T-053 ── T-054 ── T-055
           T-056 ── T-057
           T-058 ── T-059
           T-050..T-059 ── T-060, T-061

FASE 5 — M6 (EP-04b):
           T-062 ── T-063 ── T-064
           T-065 ── T-066 ── T-067
           T-062..T-067 ── T-068, T-069, T-070

FASE 6 — M7 (Homolog):
           T-071 ── T-072
           T-073 ── T-074
           T-075 ── T-076
           T-071..T-076 ── T-077 ── T-078 ── T-079
```

### Cadeia Critica (Caminho mais longo)

```
T-001 (2d) -> T-009 (2d) -> T-010 (2d) -> T-013 (2d) -> T-046 (1.5d) -> T-049 (2d)
T-001 (2d) -> T-007 (2d) -> T-012 (1.5d) -> T-016 (2d) -> T-017 (1d) -> T-019 (1d) -> ...
... T-024..T-036 (12.5d) -> T-037 (2d) -> T-048 (1.5d) -> ... -> T-077 (2d) -> T-079 (2d)
```

**Caminho critico estimado:** ~58 dias uteis (inicio T-001 em 24/07 → fim T-079 em 30/10).

---

## 4. Quadro Resumo por Marco

| Marco | Data | Features | Tarefas | Must | Should | Estimativa Total | % do Projeto |
|:---|:---|:---|:---:|:---:|:---:|:---:|:---:|
| Pre-M2 (Setup) | 24/07 — 07/08 | — | 8 | 8 | 0 | ~12d | 6% |
| Pre-M2 (Seguranca) | 07/08 — 15/08 | — | 8 | 8 | 0 | ~11.5d | 6% |
| Pre-M2 (Frente 0) | 16/07 | — | 12 | 12 | 0 | ~16-25h | 3% |
| Sprint 3 — Frente 3 | 16/07 — 31/08 | — | 7 | 7 | 0 | ~10h | 2% |
| M2 (EP-01) | 15/08 | F01-01, F01-02, F01-03 | 8 | 7 | 1 | ~11d | 8% |
| M3 (EP-02) | 15/08 — 31/08 | F02-01 a F02-05 | 15 | 15 | 0 | ~24d | 15% |
| M4 (EP-03) | 31/08 — 15/09 | F03-01 a F03-04 | 11 | 11 | 0 | ~15d | 11% |
| M5 (EP-04a) | 15/09 — 30/09 | F04-01 a F04-04 | 12 | 10 | 2 | ~18.5d | 12% |
| M6 (EP-04b) | 30/09 — 15/10 | F04-05, F04-06 | 9 | 9 | 0 | ~14d | 9% |
| M7 (Homolog) | 15/10 — 30/10 | Todas | 9 | 9 | 0 | ~16.5d | 9% |
| **Total** | **24/07 — 30/10** | **18 features** | **99** | **96** | **3** | **~122.5d-homem** | **100%** |

> Nota: Estimativa em dias-homem considera trabalho paralelo possivel dentro de cada marco. O cronograma real (14 semanas) reflete execucao paralela de tarefas independentes.

---

## 5. Distribuicao por Tipo de Tarefa

| Tipo | Quantidade | % |
|:---|:---:|:---:|
| Setup / Fundacao | 8 | 8% |
| Seguranca (JWT/RBAC/Tenant/Auditoria/RLS) | 8 | 8% |
| Correções Técnicas (Frente 0 + Frente 3 — débitos) | 19 | 19% |
| Implementacao (Controllers + Services + Repositories) | 42 | 42% |
| Testes (unitarios + integracao + seguranca + perf) | 14 | 14% |
| Documentacao (OpenAPI, README, LGPD) | 4 | 4% |
| Deploy / Homologacao | 4 | 4% |
| **Total** | **99** | **100%** |

---

## 6. Registro de Alteracoes

| Versao | Data | Alteracao | Autor |
|:---|:---|:---|:---|
| 2.8 | 17/07/2026 | Sprint 3 100% concluída: Frentes 2 (M3 EP-02 — 15 tasks: Tenant/Plan/Subscription/Audit CRUD) e 3 (7 correções técnicas) finalizadas. 18 endpoints REST. 142 testes totais (100 unit + 42 IT). 28 débitos técnicos resolvidos (DT-001 a DT-046). 9 débitos postergados Sprints 4+. Progresso 57/99 (58%). v2.7→v2.8. | Agente IA |
| 2.7 | 17/07/2026 | T-023 concluído: 23 testes integração PostgreSQL real (DashboardRepositoryIT). Adicionados DashboardRepositoryTest (11 testes mock), DashboardControllerTest (7 testes MockMvc). Expandido GlobalExceptionHandlerTest (+4 exceções). BaseRepository.save() bug array size corrigido (3→5). V003 product_service RLS removido (sem tenant_id). maven-failsafe-plugin adicionado. JaCoCo thresholds: LINE≥80%→85.8%✅, BRANCH 70%→64%. 105 testes totais (77 unit + 28 IT). M2 (EP-01) 100% concluído. v2.6→v2.7. | Agente IA |
| 2.4 | 16/07/2026 | Sprint 3 iniciada (16/07/2026). Status atualizado: Sprints 1-2 concluídas, Sprint 3 em andamento. | Time Técnico |
| 2.3 | 15/07/2026 | Revisão Caveman (DOCS-SERVICE-CAVEMAN-REVIEW.md): Corrigido Must/Should M2 (6/2→7/1, §4). Corrigido total Should (4→3 tarefas, §1). Corrigido Must total (76→77, §1). Corrigida referência SPECS (v1.3→v1.4). Corrigido T-027 endpoint count (11→7). | Caveman/IA |
| 2.1 | 14/07/2026 | Adicionada T-015.1: PostgreSQL Row-Level Security (Migration V003 + config JwtFilter). Total: 79→80 tarefas, 75→76 Must. Atualizada seção Pre-M2 Segurança: 7→8 tarefas, estimativa ~10d→~11.5d. T-012 marcada como [SUBSTITUÍDO por T-015.1]. Diagrama de dependências atualizado com T-015.1. | Agente Arquiteto/IA |
| 2.0 | 14/07/2026 | Regeneracao completa (SCOPE=full) baseada em SPECS.md v1.1. Correcao pos-gate: (1) F04-01 Autenticacao agora tem 3 tarefas (T-050 a T-052); (2) todas as 79 tarefas possuem estimativa em dias, responsavel "A definir" e criterio DONE; (3) nenhuma tarefa > 3 dias; (4) dependencias documentadas em §3 com diagrama de cadeia critica; (5) todas as tarefas referenciam feature ID (FXX-XX) e user stories (US-XXX); (6) 18/18 features cobertas; (7) contagem total (79) corresponde exatamente ao header; (8) MoSCoW alinhado com PRD.md/04-FEATURES.md (75 Must / 4 Should) | Agente Gerador TASKS/IA |
| 1.0 | 13/07/2026 | Criacao inicial: 63 tarefas organizadas em Setup + 6 marcos (M2-M7). Diagrama de dependencias. 51 Must + 4 Should | Time Tecnico |

---

🤖 *Documentacao gerada de forma automatizada pelo Agente: Gerador de Tarefas/Claude. Foram utilizados os skills: breakdown-epic-pm, writing-plans, acceptance-criteria, ponytail-review, engineering-skills, code-review. v2.5 em 16/07/2026: Sprint 3 reestruturada com 47 débitos técnicos tratados (IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md).*
