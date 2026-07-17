# SPRINT-DEVELOPMENT-PLANNING-Frente-1.md — Plano de Desenvolvimento: Sprint 3 — Frente 1

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 3 de 7 — Portal Admin + Contas e Planos
- **Frente:** 1 — M2: Portal Admin (EP-01)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + JDBC Template
- **Data:** 17/07/2026
- **Branch:** `feature/sprint-03-portal-admin`

---

## 1. Visão Geral

- **Frente Goal:** Dashboard administrativo exibe métricas operacionais do SaaS. Lista de tenants paginada com filtros e busca textual. Alertas de onboarding >48h e assinaturas suspensas.
- **Tasks a implementar:** 8 (T-016 a T-023)
- **Ordem de execução:** Sequencial com dependências (Repository → Service → DTO → Controller)
- **Features:** F01-01 (Dashboard), F01-02 (Lista de Contas), F01-03 (Alertas)
- **Build atual:** `./mvnw test` → BUILD SUCCESS, 36/36 testes passando

### Baseline (pós-Frente 0)

| Indicador | Valor |
|:---|:---|
| BaseRepository | ✅ save(), update(), findAll(), findById(), count(), softDelete() |
| AuditAspect | ✅ Funcional — captura contexto na thread principal |
| RbacAspect | ✅ TENANT, PLAN, SUBSCRIPTION, DASHBOARD na matriz |
| JaCoCo | ✅ 0.8.14 compatível com Java 25 |
| JavaMailSender | ✅ Disponível |
| Exceções de domínio | ✅ 5 classes criadas |

---

## 2. Dependências entre Tasks

```
T-018 (DTOs — sem dependências)
  │
  ├── T-016 (DashboardRepository — queries agregadas)
  │     │
  │     └── T-017 (DashboardService — lógica de métricas)
  │           │
  │           └── T-019 (DashboardController — 4 endpoints REST)
  │
  ├── T-020 (TenantRepository — findAll paginado + filtros + busca)
  │
  └── T-021 (AlertRepository/Service — onboarding >48h + assinatura suspensa)
        │
        └── T-019 (DashboardController — endpoint /alerts)

T-022 (Testes unitários — depende de T-016 a T-021)
T-023 (Testes integração — depende de T-022)
```

---

## 3. Plano por Task

### T-016 — DashboardRepository.java

- **Critério DONE:** Queries verificadas. Explain plan sem full scan
- **Estimativa:** 2d
- **Abordagem:** Criar `DashboardRepository` com queries agregadas sobre a tabela `tenant`:
  - `countActive()`: `SELECT COUNT(*) FROM tenant WHERE deleted_dt IS NULL`
  - `countByStatus(String status)`: COUNT com filtro status
  - `countBySegment(String segment)`: COUNT com filtro segment
  - `evolutionByPeriod(OffsetDateTime since)`: COUNT por período de criação
  - `accountsByStatus()`: `SELECT status, COUNT(*) GROUP BY status`
  - `accountsByPlan()`: JOIN com subscription + plan para agrupar por plano
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../repository/DashboardRepository.java` | 🆕 | Queries agregadas |
- **Dependências:** T-018 (DTOs para mapeamento de resultados)
- **Riscos:** Performance com 1000+ tenants — garantir índices V002 cobrem as queries
- **Skills:** `311-frameworks-spring-jdbc`, `postgres-pro`

### T-017 — DashboardService.java

- **Critério DONE:** Filtro recalcula métricas. Período padrão = mês atual
- **Estimativa:** 1d
- **Abordagem:** Criar `DashboardService` com:
  - `getSummary()`: agrega counts → `DashboardSummaryResponse`
  - `getEvolution(period)`: calcula evolução temporal. Períodos: 7d, 30d, 90d, `mes_atual`, `ano_atual`. Padrão = mês atual (RN01-02)
  - `getAccountsByStatus()`: agrupa por status
  - `getAccountsByPlan()`: agrupa por plano
  - `getAlerts()`: chama queries de T-021
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../service/DashboardService.java` | 🆕 | Lógica de métricas e filtros |
- **Dependências:** T-016 (DashboardRepository), T-018 (DTOs)
- **Riscos:** Cálculo correto de períodos relativos com OffsetDateTime
- **Skills:** `121-java-object-oriented-design`, `301-frameworks-spring-boot-core`

### T-018 — DTOs de Dashboard

- **Critério DONE:** JSON conforme contrato. Formatação R$. ISO 8601
- **Estimativa:** 0.5d
- **Abordagem:** Criar 4 records no pacote `dto/response/`:
  - `DashboardSummaryResponse`: totalAccounts, activeAccounts, pendingAccounts, suspendedAccounts, accountsByPlan (map), monthlyRevenue (BigDecimal)
  - `EvolutionResponse`: List<EvolutionDataPoint> com date (LocalDate) + count (int)
  - `AccountsByStatusResponse`: List<StatusCount> com status (String) + count (int)
  - `AccountsByPlanResponse`: List<PlanCount> com planName (String) + count (int)
  - `AlertResponse`: List<Alert> com type (WARNING/CRITICAL), message (String), entityId (UUID)
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../dto/response/DashboardSummaryResponse.java` | 🆕 | Record — indicadores SaaS |
  | `src/main/java/.../dto/response/EvolutionResponse.java` | 🆕 | Record — evolução temporal |
  | `src/main/java/.../dto/response/AccountsByStatusResponse.java` | 🆕 | Record — por status |
  | `src/main/java/.../dto/response/AccountsByPlanResponse.java` | 🆕 | Record — por plano |
  | `src/main/java/.../dto/response/AlertResponse.java` | 🆕 | Record — alertas |
- **Dependências:** Nenhuma
- **Riscos:** Nenhum — records simples com Jackson serialization
- **Skills:** `302-frameworks-spring-boot-rest`

### T-019 — DashboardController.java

- **Critério DONE:** p95 ≤ 3s com 1000 tenants
- **Estimativa:** 1d
- **Abordagem:** Criar `DashboardController` com 4 endpoints:
  - `GET /api/v1/dashboard/admin/summary` → `DashboardSummaryResponse`
  - `GET /api/v1/dashboard/admin/evolution?period=mes_atual` → `EvolutionResponse`
  - `GET /api/v1/dashboard/admin/accounts-by-status` → `AccountsByStatusResponse`
  - `GET /api/v1/dashboard/admin/accounts-by-plan` → `AccountsByPlanResponse`
  - `GET /api/v1/dashboard/admin/alerts` → `AlertResponse`
  - Todos anotados com `@RequiresPermission(DASHBOARD, view)`
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../controller/DashboardController.java` | 🆕 | Endpoints REST do dashboard |
- **Dependências:** T-017 (DashboardService), T-018 (DTOs)
- **Riscos:** Performance com 1000 tenants precisa de índices V002
- **Skills:** `302-frameworks-spring-boot-rest`, `304-frameworks-spring-boot-security`

### T-020 — TenantRepository.java

- **Critério DONE:** Paginação funcional. Soft delete respeitado
- **Estimativa:** 1.5d
- **Abordagem:** Criar `TenantRepository extends BaseRepository<Tenant>`:
  - `findAllPaginated(int page, int size, String status, String plan, String search)`:
    - Paginação: 25 itens padrão
    - Filtro status: `AND status = ?`
    - Filtro plano: JOIN subscription + plan
    - Busca textual: `AND (name_corporate ILIKE ? OR name_fantasy ILIKE ?)` com mínimo 3 chars
    - Ordenação: `created_dt DESC`
  - `countFiltered(...)`: COUNT com mesmos filtros para paginação
  - `findByNameCorporate(String)`: para validação de duplicidade
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../repository/TenantRepository.java` | 🆕 | Repository de Tenant |
- **Dependências:** T-015.4 (BaseRepository.save/update — já concluído)
- **Riscos:** Busca textual com ILIKE pode ser lenta sem índice GIN/trigram
- **Skills:** `311-frameworks-spring-jdbc`, `postgres-pro`

### T-021 — Alertas do Dashboard (F01-03)

- **Critério DONE:** Cards coloridos (WARNING/CRITICAL)
- **Estimativa:** 1.5d
- **Abordagem:** Adicionar queries ao `DashboardRepository`:
  - `onboardingStalled()`: tenants com `status = 'PENDING_ONBOARDING' AND created_dt < NOW() - INTERVAL '48 hours'` → type=WARNING
  - `suspendedSubscriptions()`: JOIN subscription com `status = 'SUSPENDED'` → type=CRITICAL
  - `getAlerts()` no `DashboardService`: agrega ambas as queries → `AlertResponse`
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../repository/DashboardRepository.java` | 🔄 | +2 queries de alerta |
  | `src/main/java/.../service/DashboardService.java` | 🔄 | +getAlerts() |
- **Dependências:** T-016, T-017
- **Riscos:** Nenhum — queries simples de leitura
- **Skills:** `311-frameworks-spring-jdbc`

### T-022 — Testes Unitários M2

- **Critério DONE:** ≥ 80%. Cada cenário de filtro testado positivo+negativo
- **Estimativa:** 1.5d
- **Abordagem:** Criar testes unitários com Mockito:
  - `DashboardServiceTest`: mock DashboardRepository, testar todos os períodos (7d, 30d, 90d, mes_atual, ano_atual), testar período inválido → mês atual, testar getAlerts()
  - `TenantRepositoryTest`: mock JdbcTemplate, testar paginação, filtro status, busca textual
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/test/java/.../unit/service/DashboardServiceTest.java` | 🆕 | Testes unitários do DashboardService |
  | `src/test/java/.../unit/repository/TenantRepositoryTest.java` | 🆕 | Testes unitários do TenantRepository |
- **Dependências:** T-016 a T-021
- **Skills:** `131-java-testing-unit-testing`

### T-023 — Testes de Integração M2

- **Critério DONE:** PostgreSQL real. Queries verificadas
- **Estimativa:** 1.5d
- **Abordagem:** Testes com Testcontainers PostgreSQL:
  - Popular 10+ tenants com status variados (ACTIVE, PENDING, SUSPENDED)
  - `DashboardRepositoryIT`: verificar queries agregadas retornam valores corretos
  - `TenantRepositoryIT`: verificar paginação, filtros, busca textual com dados reais
  - `DashboardControllerIT`: WebMvcTest com service real e DB container
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/test/java/.../integration/repository/DashboardRepositoryIT.java` | 🆕 | Integração DashboardRepository |
  | `src/test/java/.../integration/repository/TenantRepositoryIT.java` | 🆕 | Integração TenantRepository |
- **Dependências:** T-022
- **Skills:** `132-java-testing-integration-testing`

---

## 4. Ordem de Execução

1. **T-018** (DTOs) — 0.5d — Sem dependências, define contratos
2. **T-016** (DashboardRepository) — 2d — Queries agregadas
3. **T-020** (TenantRepository) — 1.5d — Paralelizável com T-016 (repositories diferentes)
4. **T-017** (DashboardService) — 1d — Depende de T-016 + T-018
5. **T-021** (Alertas) — 1.5d — Depende de T-016 + T-017
6. **T-019** (DashboardController) — 1d — Depende de T-017 + T-018
7. **T-022** (Testes unitários) — 1.5d — Depende de T-016 a T-021
8. **T-023** (Testes integração) — 1.5d — Depende de T-022

---

## 5. Estratégia de Build e Verificação

- **Comando de build:** `./mvnw compile`
- **Comando de teste:** `./mvnw test`
- **Checkpoints:**
  - Após T-018 → `./mvnw compile` (DTOs compilam)
  - Após T-016 + T-020 → `./mvnw compile` (repositories compilam)
  - Após T-017 + T-021 → `./mvnw compile` (services compilam)
  - Após T-019 → `./mvnw compile` (controller compila)
  - Após T-022 → `./mvnw test` (testes unitários passam)
  - Após T-023 → `./mvnw test` (testes integração passam)
  - Ao final → `./mvnw verify` (JaCoCo ≥ 80%)

---

🤖 *Gerado por Agente IA em 17/07/2026. Stack: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + JDBC Template. Skills: 311-frameworks-spring-jdbc, 302-frameworks-spring-boot-rest, 304-frameworks-spring-boot-security, 121-java-object-oriented-design, 131-java-testing-unit-testing, 132-java-testing-integration-testing, postgres-pro.*
