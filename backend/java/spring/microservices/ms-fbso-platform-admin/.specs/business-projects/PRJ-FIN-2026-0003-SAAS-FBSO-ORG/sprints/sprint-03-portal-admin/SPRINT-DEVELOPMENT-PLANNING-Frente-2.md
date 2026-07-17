# SPRINT-DEVELOPMENT-PLANNING-Frente-2.md — Plano de Desenvolvimento: Sprint 3 — Frente 2 (M3)

- **Solução:** `ms-fbso-platform-admin`
- **Sprint:** 3 de 7 — Portal Admin + Contas e Planos
- **Frente:** 2 — M3: Gestão de Clientes e Planos (EP-02)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + JDBC Template + Flyway + Testcontainers
- **Tasks:** T-024 a T-038 (15)
- **Data:** 17/07/2026
- **Status:** 📋 Planejamento — aguardando início da implementação

---

## 1. Visão Geral

- **Sprint Goal (M3):** CRUD completo de Tenants, Planos e Assinaturas funcional. Auditoria registra 100% das ações administrativas e é consultável com filtros.
- **Tasks a implementar:** 15 (T-024 a T-038)
- **Features:** F02-01 a F02-05 (5 features)
- **Endpoints REST:** 13 novos (18 total na sprint)
- **RNs a implementar:** 18 (RN04-01 a RN08-02)
- **Ordem de execução:** Sequencial com paralelismo intra-feature
- **Dependência externa:** Nenhuma — todas as dependências já estão resolvidas pela Frente 0

### Pré-requisitos já atendidos pela Frente 0 ✅

| Pré-requisito | Status | Task |
|:---|:---:|:---|
| BaseRepository.save(T) e update(T) genéricos | ✅ | T-015.4.DT-003 |
| JavaMailSender injetável (spring-boot-starter-mail) | ✅ | T-015.8.DT-007 |
| Migration V005 (locked_price + locked_recurrence) | ✅ | T-015.10.DT-009 |
| 5 exceções de domínio (DuplicateCnpj, InvalidStatusTransition, etc.) | ✅ | T-015.13.DT-012 |
| RbacAspect cobre TENANT, PLAN, SUBSCRIPTION | ✅ | T-015.6.DT-005 |
| Tenant entity + RowMapper | ✅ | T-020 (Frente 1) |
| TenantRepository (findAllPaginated, countFiltered, findByNameCorporate) | ✅ | T-020 (Frente 1) |

---

## 2. Dependências entre Tasks

```
FASE 2 — M3 (EP-02)
═══════════════════════════════════════════

BLOCO A: Tenant (T-024 → T-025 → T-027)
         │
         ├── T-024 Tenant Entity + DTOs + Status Map
         │        │
         │        ├── T-025 TenantService (create + update)
         │        │        │
         │        │        └── T-027 TenantController (7 endpoints)
         │        │                 │
         │        │                 └── T-028 Email (depende de T-025.create)
         │        │
         │        └── T-026 TenantService (suspend + reactivate)
         │                 │
         │                 └── T-027 (compartilha controller com T-025)
         │
BLOCO B: Plan (independente do Bloco A — paralelizável)
         │
         ├── T-029 Plan Entity + PlanModule + PlanRepository
         │        │
         │        └── T-030 PlanService (CRUD versionado + deactivate)
         │                 │
         │                 └── T-031 PlanController (CRUD + deactivate)
         │
BLOCO C: Subscription (depende de T-024 Tenant + T-029 Plan)
         │
         ├── T-032 Subscription Entity + Repository
         │        │
         │        └── T-033 SubscriptionService (create, change-plan, suspend, reactivate)
         │                 │
         │                 └── T-034 SubscriptionController (4 endpoints)
         │
BLOCO D: Auditoria (depende de T-025, T-030, T-033 — precisa de dados)
         │
         ├── T-035 AuditRepository + AuditService (filtros, paginação)
         │        │
         │        └── T-036 AuditController (GET /audit)
         │
BLOCO E: Testes (após todos os blocos acima)
         │
         ├── T-037 Testes unitários M3 (todos os services)
         └── T-038 Testes integração M3 (CRUD completo com Testcontainers)
```

---

## 3. Plano por Task

### T-024 — Tenant Entity + DTOs + Status Map (RN05-01)

- **Critério DONE:** Transições válidas OK. Inválida → 422
- **Estimativa:** 1d
- **Abordagem:**
  - `Tenant.java` já existe (Frente 1). Adicionar DTOs request/response.
  - `TenantCreateRequest`: nameCorporate, nameFantasy, segment, cnpj (opcional)
  - `TenantUpdateRequest`: nameFantasy, segment (campos editáveis)
  - `TenantResponse`: id, nameCorporate, nameFantasy, segment, status, createdDt
  - Mapa de transições: `Map<TenantStatus, Set<TenantStatus>>` estático em TenantStatus ou TenantService
  - Transições: PENDING→ACTIVE, ACTIVE→SUSPENDED, SUSPENDED→ACTIVE, ACTIVE→INACTIVE
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `dto/request/TenantCreateRequest.java` | 🆕 | Record com Bean Validation |
  | `dto/request/TenantUpdateRequest.java` | 🆕 | Record com campos editáveis |
  | `dto/response/TenantResponse.java` | 🆕 | Record de resposta |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `entity/Tenant.java` | 🔄 | Adicionar campo cnpj, métodos de transição |
- **Dependências:** Nenhuma (Tenant entity + enums já existem)
- **Skills aplicáveis:** 121-java-object-oriented-design, 122-java-type-design, 302-frameworks-spring-boot-rest

### T-025 — TenantService.create() + update()

- **Critério DONE:** Duplicada → 409. Auditoria @Auditable
- **Estimativa:** 2d
- **Abordagem:**
  - `create(TenantCreateRequest)`: validar duplicidade via TenantRepository.findByNameCorporate(), status inicial PENDING_ONBOARDING, salvar via BaseRepository.save()
  - `update(UUID, TenantUpdateRequest)`: buscar tenant, aplicar campos, validar existência, salvar via BaseRepository.update()
  - Anotar com `@Auditable(entityType = "TENANT", action = "CREATED"/"UPDATED")`
  - Anotar com `@Transactional`
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `service/TenantService.java` | 🆕 | CRUD + validações + auditoria |
- **Dependências:** T-024 (DTOs), T-020 (TenantRepository)
- **Riscos:** Duplicidade de razão social precisa de índice parcial (V002 já criou)
- **Skills aplicáveis:** 301-frameworks-spring-boot-core, 126-java-exception-handling, 311-frameworks-spring-jdbc

### T-026 — TenantService.suspend() + reactivate()

- **Critério DONE:** Sem motivo → 400. Timeline status
- **Estimativa:** 2d
- **Abordagem:**
  - `suspend(UUID, String reason)`: validar transição ACTIVE→SUSPENDED, validar motivo não-vazio, atualizar status, registrar auditoria com reason
  - `reactivate(UUID)`: validar transição SUSPENDED→ACTIVE, atualizar status
  - Validação da máquina de estados via `isValidTransition(from, to)`
  - `@Auditable` em ambos os métodos
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `service/TenantService.java` | 🔄 | Adicionar métodos suspend/reactivate |
- **Dependências:** T-025 (TenantService base)
- **Riscos:** Transição inválida precisa de mensagem clara (qual transição foi tentada, quais são válidas)
- **Skills aplicáveis:** 126-java-exception-handling, 121-java-object-oriented-design

### T-027 — TenantController

- **Critério DONE:** 7 endpoints REST. Bean Validation
- **Estimativa:** 2d
- **Abordagem:**
  - `GET /api/v1/tenants` — paginado, filtros (herda de T-020 TenantRepository)
  - `GET /api/v1/tenants/{id}` — por ID
  - `POST /api/v1/tenants` — criar
  - `PATCH /api/v1/tenants/{id}` — editar
  - `POST /api/v1/tenants/{id}/suspend` — suspender (body: {reason})
  - `POST /api/v1/tenants/{id}/reactivate` — reativar
  - `POST /api/v1/tenants/{id}/resend-invite` — reenviar convite
  - Anotar com `@RequiresPermission(TENANT, create/edit/view/suspend/reactivate)`
  - `@Valid` nos DTOs de entrada
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `controller/TenantController.java` | 🆕 | 7 endpoints REST |
- **Dependências:** T-025, T-026 (TenantService completo)
- **Skills aplicáveis:** 302-frameworks-spring-boot-rest, 303-frameworks-spring-boot-validation

### T-028 — Integração Email

- **Critério DONE:** Email enviado. Link expira. Reenvio funcional
- **Estimativa:** 2d
- **Abordagem:**
  - `EmailService`: interface + implementação com JavaMailSender
  - Template de email: link único com token JWT de curta duração (7 dias)
  - Disparo automático após TenantService.create()
  - Reenvio via TenantController.resend-invite — gera novo token
  - Mock SMTP para testes (Mailhog/GreenMail)
  - Log de envio registrado em audit_log
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `service/EmailService.java` | 🆕 | Interface de envio de email |
  | `service/EmailServiceImpl.java` | 🆕 | Implementação com JavaMailSender |
- **Dependências:** T-025 (TenantService.create), spring-boot-starter-mail (T-015.8 ✅)
- **Riscos:** SMTP externo pode não estar disponível em dev. Usar Mailhog como fallback
- **Skills aplicáveis:** 301-frameworks-spring-boot-core

### T-029 — Plan Entity + PlanModule + PlanRepository

- **Critério DONE:** Plano criado ACTIVE. PlanModule vinculado
- **Estimativa:** 1d
- **Abordagem:**
  - `Plan` entity: id, name, description, price (>0), recurrence, status, version
  - `PlanModule` entity: id, planId, moduleName
  - `PlanRepository` extends BaseRepository\<Plan\>: findAll, findById, save, update
  - `PlanModuleRepository`: findByPlanId, save, deleteByPlanId
  - Validação price > 0 no construtor/setter
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `entity/Plan.java` | 🆕 | Entidade com campos de plano |
  | `entity/PlanModule.java` | 🆕 | Entidade de módulo vinculado |
  | `repository/PlanRepository.java` | 🆕 | Repository JDBC |
  | `repository/PlanModuleRepository.java` | 🆕 | Repository de módulos |
  | `repository/rowmapper/PlanRowMapper.java` | 🆕 | RowMapper JDBC |
- **Dependências:** Nenhuma (tabelas plan + plan_module já existem via V001)
- **Skills aplicáveis:** 121-java-object-oriented-design, 311-frameworks-spring-jdbc

### T-030 — PlanService CRUD Versionado + deactivate()

- **Critério DONE:** Edição gera nova versão. Último plano ativo → 422
- **Estimativa:** 2d
- **Abordagem:**
  - `create(PlanCreateRequest)`: criar plano ACTIVE, version=1, vincular PlanModules
  - `update(UUID, PlanUpdateRequest)`: se preço alterou → version++, novo registro. Se não → update normal
  - `deactivate(UUID)`: validar se não é o último plano ativo (RN06-03)
  - Validação: plano com assinantes ativos não pode ser desativado → 422 (RN06-01)
  - `@Auditable` em create/update/deactivate
  - `@Transactional`
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `service/PlanService.java` | 🆕 | CRUD + versionamento + deactivate |
  | `dto/request/PlanCreateRequest.java` | 🆕 | Record |
  | `dto/request/PlanUpdateRequest.java` | 🆕 | Record |
  | `dto/response/PlanResponse.java` | 🆕 | Record |
- **Dependências:** T-029 (PlanRepository)
- **Riscos:** Versionamento precisa de lógica cuidadosa para preservar assinaturas existentes
- **Skills aplicáveis:** 121-java-object-oriented-design, 126-java-exception-handling

### T-031 — PlanController

- **Critério DONE:** CRUD completo. Desativado = "Descontinuado"
- **Estimativa:** 1d
- **Abordagem:**
  - `GET /api/v1/plans` — lista planos ativos (para assinatura)
  - `GET /api/v1/plans/admin` — lista todos (admin)
  - `GET /api/v1/plans/{id}` — por ID
  - `POST /api/v1/plans` — criar
  - `PATCH /api/v1/plans/{id}` — editar
  - `POST /api/v1/plans/{id}/deactivate` — desativar
  - `@RequiresPermission(PLAN, create/edit/view/deactivate)`
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `controller/PlanController.java` | 🆕 | 6 endpoints REST |
- **Dependências:** T-030 (PlanService)
- **Skills aplicáveis:** 302-frameworks-spring-boot-rest, 303-frameworks-spring-boot-validation

### T-032 — Subscription Entity + Repository

- **Critério DONE:** Segunda ativa → 409
- **Estimativa:** 1d
- **Abordagem:**
  - `Subscription` entity: id, tenantId FK, planId FK, startDate, endDate, status, lockedPrice, lockedRecurrence
  - `SubscriptionRepository` extends BaseRepository\<Subscription\> (hasTenantColumn=true)
  - Query: `findActiveByTenantId(UUID)` — verificar RN07-01 (1 ativa por tenant)
  - Query: `findByTenantId(UUID)` — histórico completo
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `entity/Subscription.java` | 🆕 | Entidade com locked_price/locked_recurrence |
  | `repository/SubscriptionRepository.java` | 🆕 | Repository com findActiveByTenantId |
  | `repository/rowmapper/SubscriptionRowMapper.java` | 🆕 | RowMapper JDBC |
- **Dependências:** T-024 (Tenant entity), T-029 (Plan entity)
- **Riscos:** RLS ativo na tabela subscription — queries precisam ter app.current_tenant_id configurado
- **Skills aplicáveis:** 311-frameworks-spring-jdbc

### T-033 — SubscriptionService

- **Critério DONE:** RN07-02: Change-plan sem gap. Transação atômica
- **Estimativa:** 2d
- **Abordagem:**
  - `create(tenantId, planId)`: validar tenant existe, validar plano ativo, validar sem assinatura ativa (RN07-01), locked_price = plan.price, locked_recurrence = plan.recurrence
  - `changePlan(subscriptionId, newPlanId)`: finalizar atual (endDate=NOW), criar nova (startDate=NOW) — na mesma transação (RN07-02). `@Transactional`
  - `suspend(subscriptionId)`: status → SUSPENDED
  - `reactivate(subscriptionId)`: status → ACTIVE
  - `@Auditable` em todos os métodos
  - Race condition prevention: `SELECT ... FOR UPDATE` na assinatura ativa
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `service/SubscriptionService.java` | 🆕 | CRUD + change-plan atômico |
  | `dto/request/SubscriptionCreateRequest.java` | 🆕 | Record |
  | `dto/request/ChangePlanRequest.java` | 🆕 | Record |
  | `dto/response/SubscriptionResponse.java` | 🆕 | Record |
- **Dependências:** T-032 (SubscriptionRepository), T-025 (TenantService), T-030 (PlanService)
- **Riscos:** Race condition (2 assinaturas simultâneas) — mitigar com SELECT FOR UPDATE + @Transactional
- **Skills aplicáveis:** 125-java-concurrency, 126-java-exception-handling

### T-034 — SubscriptionController

- **Critério DONE:** 4 endpoints. Validações OK
- **Estimativa:** 1.5d
- **Abordagem:**
  - `POST /api/v1/tenants/{tid}/subscriptions` — criar
  - `GET /api/v1/tenants/{tid}/subscriptions` — listar histórico
  - `POST /api/v1/subscriptions/{id}/change-plan` — mudar plano
  - `POST /api/v1/subscriptions/{id}/suspend` — suspender
  - `@RequiresPermission(SUBSCRIPTION, create/view/edit)`
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `controller/SubscriptionController.java` | 🆕 | 4 endpoints REST |
- **Dependências:** T-033 (SubscriptionService)
- **Skills aplicáveis:** 302-frameworks-spring-boot-rest, 303-frameworks-spring-boot-validation

### T-035 — AuditRepository + AuditService

- **Critério DONE:** Filtros funcionais. Imutável — UPDATE/DELETE → 403
- **Estimativa:** 1.5d
- **Abordagem:**
  - `AuditEntry` entity: id, timestamp, tenantId, action, entityType, entityId, actorId, actorName, previousValue, newValue, reason
  - `AuditRepository` extends BaseRepository\<AuditEntry\> (hasTenantColumn=true)
  - Query: `findByFilters(startDate, endDate, action, entityType, page, size)` — ORDER BY timestamp DESC
  - `AuditService`: encapsular queries de filtro, paginação obrigatória (padrão 25, max 100)
  - Imutável: repository não expõe save/update/delete públicos (apenas find)
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `entity/AuditEntry.java` | 🆕 | Entidade imutável |
  | `repository/AuditRepository.java` | 🆕 | Repository somente-leitura |
  | `repository/rowmapper/AuditEntryRowMapper.java` | 🆕 | RowMapper |
  | `service/AuditService.java` | 🆕 | Serviço de consulta com filtros |
  | `dto/response/AuditEntryResponse.java` | 🆕 | Record de resposta |
- **Dependências:** Dados gerados por T-025, T-030, T-033 (auditoria já funciona)
- **Skills aplicáveis:** 311-frameworks-spring-jdbc, 126-java-exception-handling

### T-036 — AuditController

- **Critério DONE:** Admin vê tudo. Auditor vê tudo (leitura)
- **Estimativa:** 1d
- **Abordagem:**
  - `GET /api/v1/audit` — com query params: startDate, endDate, action, entityType, page, size
  - `@RequiresPermission(AUDIT, view)`
  - Paginação padrão 25, máximo 100
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `controller/AuditController.java` | 🆕 | 1 endpoint GET |
- **Dependências:** T-035 (AuditService)
- **Skills aplicáveis:** 302-frameworks-spring-boot-rest

### T-037 — Testes Unitários M3

- **Critério DONE:** ≥ 80%. Cada RN testada positivo+negativo
- **Estimativa:** 2d
- **Abordagem:**
  - `TenantServiceTest`: create (duplicidade → 409), update, suspend (sem motivo → 400), reactivate, transições inválidas → 422
  - `PlanServiceTest`: create, update (versiona), deactivate (com/sem assinantes), último plano → 422
  - `SubscriptionServiceTest`: create (1 ativa → 409), change-plan (atômico), suspend, reactivate
  - `AuditServiceTest`: filtros, paginação, imutabilidade
  - Mockito para repositories, AssertJ para assertions
  - Cobrir RN05-01 (todas as transições), RN07-01 (1 ativa), RN06-01 (desativar com assinantes)
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/test/.../unit/service/TenantServiceTest.java` | 🆕 | ~8 testes |
  | `src/test/.../unit/service/PlanServiceTest.java` | 🆕 | ~6 testes |
  | `src/test/.../unit/service/SubscriptionServiceTest.java` | 🆕 | ~7 testes |
  | `src/test/.../unit/service/AuditServiceTest.java` | 🆕 | ~5 testes |
- **Dependências:** T-025 a T-036 (todos os services implementados)
- **Skills aplicáveis:** 321-frameworks-spring-boot-testing-unit-tests, 131-java-testing-unit-testing

### T-038 — Testes Integração M3

- **Critério DONE:** PostgreSQL real. RN07-01, RN06-01, RN05-01, RN08-02
- **Estimativa:** 2d
- **Abordagem:**
  - Estender `BaseIntegrationTest` — Testcontainers PostgreSQL 17
  - Seed data: 2 tenants, 3 planos, 2 assinaturas
  - Cenários:
    - Criar tenant → 201 PENDING_ONBOARDING
    - Criar duplicado → 409
    - Transições de status: ciclo completo PENDING→ACTIVE→SUSPENDED→ACTIVE
    - Transição inválida → 422
    - Criar assinatura → 201
    - Segunda ativa → 409
    - Change-plan → assinatura anterior finalizada + nova criada
    - Desativar plano com assinantes → 422
    - UPDATE/DELETE em audit_log → 403
    - Concorrência: 2 threads tentando criar assinatura simultânea
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/test/.../integration/repository/TenantRepositoryIT.java` | 🆕 | ~5 testes |
  | `src/test/.../integration/repository/PlanRepositoryIT.java` | 🆕 | ~4 testes |
  | `src/test/.../integration/repository/SubscriptionRepositoryIT.java` | 🆕 | ~5 testes |
  | `src/test/.../integration/repository/AuditRepositoryIT.java` | 🆕 | ~4 testes |
- **Dependências:** T-037 (testes unitários) + T-024 a T-036 (tudo implementado)
- **Skills aplicáveis:** 322-frameworks-spring-boot-testing-integration-tests, 132-java-testing-integration-testing

---

## 4. Ordem de Execução

| Ordem | Task(s) | Bloco | Justificativa |
|:---:|:---|:---|:---|
| 1 | **T-024** | Tenant | DTOs + status map — pré-requisito para todo o bloco Tenant |
| 2 | **T-025** | Tenant | TenantService.create/update — depende de T-024 |
| 3 | **T-026** | Tenant | suspend/reactivate — expande T-025 |
| 4 | **T-027** | Tenant | TenantController — depende de T-025 + T-026 |
| 5 | **T-028** | Tenant | Email — depende de T-025, paralelizável com T-026/T-027 |
| 6 | **T-029** | Plan | Plan entity + repository — independente, pode iniciar junto com Tenant |
| 7 | **T-030** | Plan | PlanService — depende de T-029 |
| 8 | **T-031** | Plan | PlanController — depende de T-030 |
| 9 | **T-032** | Subscription | Entity + repository — depende de T-024 + T-029 |
| 10 | **T-033** | Subscription | SubscriptionService — depende de T-032 |
| 11 | **T-034** | Subscription | Controller — depende de T-033 |
| 12 | **T-035** | Auditoria | Repository + service — depende de dados existentes |
| 13 | **T-036** | Auditoria | Controller — depende de T-035 |
| 14 | **T-037** | Testes | Unitários — depende de todos os services |
| 15 | **T-038** | Testes | Integração — depende de T-037 |

> **Paralelismo possível:** Bloco B (Plan: T-029 a T-031) pode ser executado em paralelo com Bloco A (Tenant: T-024 a T-028), pois são independentes. Bloco C (Subscription) depende de A+B.

---

## 5. Estratégia de Build e Verificação

- **Comando de compilação:** `./mvnw compile -Dcheckstyle.skip=true -q`
- **Comando de teste unitário:** `./mvnw test -Dcheckstyle.skip=true`
- **Comando de teste integração (IT específico):** `./mvnw test -Dtest="NomeIT" -DfailIfNoTests=false -Dcheckstyle.skip=true`
- **Comando de verificação completa:** `./mvnw verify -Dcheckstyle.skip=true`
- **Checkpoints:**
  - Após cada bloco (A: Tenant, B: Plan, C: Subscription, D: Auditoria) → executar build + testes unitários
  - Após T-037 → verificar cobertura JaCoCo ≥ 80%
  - Após T-038 → executar `mvn verify` com todos os testes

---

## Rodapé

🤖 *Plano de desenvolvimento gerado em 17/07/2026 pelo PROMPT-EXECUTE-SPRINT-TASKS (Fase 1). Stack detectada: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + JDBC Template + Flyway 10.22.0 + Testcontainers 1.20.6. 15 tasks mapeadas em 5 blocos (A-E). Próximo passo: Fase 2 — Implementação sequencial.*
