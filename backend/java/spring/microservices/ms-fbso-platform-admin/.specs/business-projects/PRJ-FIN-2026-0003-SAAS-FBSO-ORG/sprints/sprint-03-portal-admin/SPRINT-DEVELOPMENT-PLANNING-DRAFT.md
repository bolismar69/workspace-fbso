# SPRINT-DEVELOPMENT-PLANNING-DRAFT.md — Planejamento de Desenvolvimento: Sprint 3

- **Sprint:** 3 de 7 — Portal Admin + Contas e Planos
- **Marco:** M2 (EP-01) + M3 (EP-02)
- **Datas:** 16/07/2026 → 31/08/2026
- **Duração efetiva:** ~15 dias úteis (12 originais + 3 para Frentes 0 e 3 de débitos técnicos)
- **Status:** 🟢 Em andamento — iniciada 16/07/2026. Frentes 0+1 concluídas (20/42 tarefas). Frentes 0 e 3 adicionadas após auditoria de débitos. 105 testes (77 unit + 28 IT), JaCoCo 87.1%
- **Documentos-mestre:** [TASKS.md](../../TASKS.md) v2.6 · [SPECS.md](../../SPECS.md) v1.8 · [TEST_PLAN.md](../../TEST_PLAN.md) v2.6 · [PRD.md](../../PRD.md) v1.9 · [ARCHITECTURE.md](../../ARCHITECTURE.md) v2.2
- **Débitos Técnicos:** [IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md](IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md) — 47 débitos, 12 pré-sprint, 7 durante-sprint, 9 postergados Sprint 4

---


## 1. Diagnóstico da Situação Atual

### 1.1 Progresso Acumulado (Sprints 1-2)

| Indicador | Valor |
|:---|:---|
| Tarefas concluídas | 16/80 (20%) |
| Sprints concluídas | 2/7 (Setup + Segurança) |
| Testes passando | 22/22 (Sprint 2) |
| Build | `mvn compile` SUCCESS |
| Cobertura de testes | Infraestrutura de segurança testada |

### 1.2 Código Existente (Pré-Sprint 3)

```
src/main/java/com/fbso/platform/admin/
├── config/           ✅ DataSourceConfig, SecurityConfig, TenantAwareDataSource
├── security/         ✅ JwtAuthenticationFilter, TenantContext, RbacAspect, AuditAspect,
│   ├── annotation/   ✅ RequiresPermission, Auditable
│   └── aspect/       ✅ RbacAspect, AuditAspect
├── common/           ✅ BaseEntity, Address
├── enums/            ✅ 8 enums (TenantStatus, Role, Recurrence, etc.)
├── exception/        ✅ BusinessException, GlobalExceptionHandler, PermissionDeniedException
├── repository/common/ ✅ BaseRepository (template JDBC)
├── dto/response/     ✅ ErrorResponse
├── utils/            ✅ JwtUtils
│
├── controller/       🔴 VAZIO — 0 controladores
├── service/          🔴 VAZIO — 0 services
├── entity/           🔴 VAZIO — 0 entidades
├── dto/request/      🔴 VAZIO — 0 DTOs de entrada
└── repository/       🔴 VAZIO — 0 repositories de domínio (apenas BaseRepository)
```

### 1.3 Diagnóstico

**O pipeline de segurança está pronto.** Toda requisição passa por:
- `JwtAuthenticationFilter` → valida JWT, extrai claims, seta `TenantContext`
- `RbacAspect` → verifica `@RequiresPermission` contra role do JWT
- `TenantAwareDataSource` → configura `app.current_tenant_id` na sessão PostgreSQL
- PostgreSQL RLS → força `WHERE tenant_id = current_setting(...)` em 5 tabelas

**O código de negócio é zero.** A Sprint 3 constrói a primeira camada visível do sistema. Todo o trabalho será greenfield sobre uma fundação sólida.

---

## 2. Escopo da Sprint 3

### 2.1 Visão Geral

| Dimensão | Quantidade |
|:---|:---:|
| Tarefas | 23 (T-016 a T-038) |
| Milestones | 2 (M2 + M3) |
| Features | 8 (F01-01 a F02-05) |
| Regras de negócio | 21 (RN01-01 a RN08-02) |
| Endpoints REST | 18 novos |
| Cenários de teste | 56 (24 unit + 20 integração + 6 E2E + 6 segurança) |
| Esforço estimado | ~35 dias-homem |

### 2.2 Features por Milestone

#### M2 — Portal Admin (EP-01) | T-016 a T-023

| Feature | Tarefas | O que entrega |
|:---|:---|:---|
| **F01-01** Dashboard Métricas | T-016, T-017, T-018, T-019 | 4 endpoints de dashboard com métricas agregadas |
| **F01-02** Lista de Contas | T-020 | Lista paginada com busca textual e filtros |
| **F01-03** Alertas | T-021 (Should) | Cards de onboarding >48h e assinatura suspensa |

#### M3 — Gestão de Clientes e Planos (EP-02) | T-024 a T-038

| Feature | Tarefas | O que entrega |
|:---|:---|:---|
| **F02-01** Criar Tenant | T-024, T-025, T-027, T-028 | CRUD tenants + envio de email de ativação |
| **F02-02** Status do Tenant | T-024, T-026, T-027 | Transições de status com máquina de estados |
| **F02-03** Planos | T-029, T-030, T-031 | CRUD planos com versionamento |
| **F02-04** Assinaturas | T-032, T-033, T-034 | CRUD assinaturas + change-plan atômico |
| **F02-05** Auditoria | T-035, T-036 | Consulta de auditoria com filtros |

---

## 3. Sequenciamento de Desenvolvimento

### 3.1 Grafo de Dependências

```
Semana 1 (M2 — Dashboard)
══════════════════════════
T-016 DashboardRepository ──┬── T-017 DashboardService ── T-019 DashboardController
                            │
T-018 DTOs ─────────────────┘
                            │
T-020 TenantRepository ─────┼── T-021 Alertas (Should, paralelo)

T-022 Testes Unitários M2 ──┤ (paralelo com desenvolvimento)
T-023 Testes Integração M2 ─┘ (após repositórios prontos)


Semana 2 (M3 — Domínio Core)
════════════════════════════
T-024 Entidade Tenant ────── T-025 TenantService ───── T-027 TenantController
                                                      │
                                          T-028 Email ┘ (paralelo, mock SMTP)

T-029 Entidade Plan ──────── T-030 PlanService ─────── T-031 PlanController

T-032 Entidade Subscription  T-033 SubscriptionService  T-034 SubscriptionController
    (depende de T-024, T-029)  (depende de T-032)       (depende de T-033)


Semana 3 (M3 — Finalização + Testes)
═══════════════════════════════════
T-035 AuditRepository ────── T-036 AuditController
    (depende de dados de T-025, T-030, T-033)

T-037 Testes Unitários M3 ──┤ (paralelo com desenvolvimento)
T-038 Testes Integração M3 ─┘ (após services prontos)
```

### 3.2 Ordem Recomendada de Implementação

| Dia | Bloco | Tarefas | Entregável |
|:---:|:---|:---|:---|
| **1** | Fundação M2 | T-016 (DashboardRepository), T-020 (TenantRepository), T-018 (DTOs) | Queries agregadas + listagem funcional |
| **2** | Lógica M2 | T-017 (DashboardService), T-019 (DashboardController) | 4 endpoints dashboard |
| **3** | Alertas + Testes | T-021 (Alertas), T-022 (início unitários M2) | Cards de alerta |
| **4** | Testes M2 | T-022 (conclusão), T-023 (integração) | M2 concluído — demo interno |
| **5** | Entidades M3 | T-024 (Tenant RowMapper — enums já existem), T-029 (Plan), T-032 (Subscription) | 3 entidades + DTOs |
| **6** | Services Core | T-025 (TenantService), T-030 (PlanService) | CRUD Tenant + Plan |
| **7** | Controllers Core | T-027 (TenantController), T-031 (PlanController) | 11 endpoints REST |
| **8** | Assinaturas | T-033 (SubscriptionService), T-034 (SubscriptionController) | CRUD Subscription + change-plan |
| **9** | Email + Status | T-026 (suspend/reactivate), T-028 (email) | Máquina de estados + envio email |
| **10** | Auditoria | T-035 (AuditRepository+Service), T-036 (AuditController) | GET /audit com filtros |
| **11** | Testes M3 | T-037 (unitários M3), T-038 (início integração) | Cobertura ≥ 80% |
| **12** | Finalização | T-038 (conclusão), smoke test completo, demo PO | Sprint 3 concluída |

> **Nota:** Dias 1-12 são dias úteis. V004 é opcional (índices já existem na V002). T-024 reduzido (enums `TenantStatus`, `SubscriptionStatus`, `Recurrence`, `Segment` já existem no código). Com 1 dev full-time ≈ 3 semanas. Com 2 devs ≈ 2 semanas.

---

## 4. Decisões Técnicas Pré-Desenvolvimento

### 4.1 Decisões que Precisam ser Tomadas ANTES de Codificar

| # | Decisão | Opções | Recomendação | Impacto |
|:---:|:---|:---|:---|:---|
| **D1** | RowMapper vs. ResultSetExtractor para queries de dashboard | RowMapper para queries simples; ResultSetExtractor para agregações complexas | `ResultSetExtractor` para T-016 (múltiplas métricas em 1 query) | Performance do dashboard |
| **D2** | Paginação: offset-based vs. keyset | Offset: `LIMIT ? OFFSET ?`. Keyset: `WHERE id > ? LIMIT ?` | **Offset-based** — simplicidade, 1000 tenants é volume baixo. Migrar para keyset se necessário | Complexidade do repository |
| **D3** | Versionamento de Planos: campo `version` incrementado vs. tabela `plan_version` separada | Campo: simples, 1 tabela. Tabela separada: histórico completo, mais complexo | **Campo `version` incrementado** — Fase 0 não exige histórico de versões. Se necessário, migrar na Fase 1 | Schema, complexidade |
| **D4** | Email: mock para dev vs. integração real desde o início | Mock (Mailhog/GreenMail): sem dependência externa. Real: testa integração | **Mock (Mailhog)** para dev + testes. Configuração real apenas em `staging` profile | T-028, testes |
| **D5** | Índices para performance do dashboard | V002 já cobre 5/6 índices. V004 adiciona só `idx_tenant_segment` + `idx_plan_status` | Criar **V004 enxuta** (2 índices). Revisado pelo Caveman Review v2.0 | Performance (meta ≤3s) |
| **D6** | AuditAspect: como capturar valores "antes" em updates | Reflection vs. Snapshot do entityManager vs. Query pré-update | **Query `SELECT *` antes do UPDATE** — mais simples e preciso que reflection | T-035, complexidade do AuditAspect |

### 4.2 Migration V004 — Índices de Performance (REVISADO)

**PRÉ-REQUISITO para T-016.** A maioria dos índices já existe na V002. Apenas 2 índices novos são necessários:

```sql
-- V004__dashboard_performance_indexes.sql
-- NOTA: idx_tenant_status, idx_tenant_created, idx_subscription_tenant_active,
--       idx_audit_log_timestamp, idx_audit_log_tenant já existem na V002.
--       Esta migration cria apenas os índices que faltam.

-- Novo: filtro por segmento no dashboard (não coberto pela V002)
CREATE INDEX IF NOT EXISTS idx_tenant_segment
    ON fbso_platform.tenant (segment) WHERE deleted_dt IS NULL;

-- Novo: busca de planos por status (único índice que faltava para plan)
CREATE INDEX IF NOT EXISTS idx_plan_status
    ON fbso_platform.plan (status) WHERE deleted_dt IS NULL;
```

---

## 5. Estratégia de Desenvolvimento por Camada

### 5.1 Padrão de Implementação (Ordem Strict)

Para **cada feature**, implementar nesta ordem:

```
1. Entity     → modelar a tabela como classe Java
2. DTOs       → request + response contracts
3. Repository → queries JDBC (começar com teste de integração!)
4. Service    → lógica de negócio + validações de RN
5. Controller → endpoints REST + @RequiresPermission
6. Testes     → unitário (service mockado) + integração (Testcontainers)
```

**Regra:** Nunca implementar o Controller antes do Service estar testado. Nunca implementar o Service antes do Repository estar funcional no Testcontainers.

### 5.2 Templates de Código

#### Repository (JDBC Template — estende BaseRepository)

```java
@Repository
public class DashboardRepository extends BaseRepository<Tenant> {

    public DashboardRepository(JdbcTemplate jdbc) {
        super(jdbc, "fbso_platform.tenant", new TenantRowMapper(), true);
    }

    public DashboardSummaryDTO getSummary() {
        String sql = """
            SELECT
                COUNT(*) FILTER (WHERE status = 'ACTIVE') AS active_count,
                COUNT(*) FILTER (WHERE status = 'PENDING_ONBOARDING') AS pending_count,
                COUNT(*) FILTER (WHERE status = 'SUSPENDED') AS suspended_count
            FROM fbso_platform.tenant
            WHERE deleted_dt IS NULL
            """;
        return jdbc.query(sql, new DashboardSummaryMapper());
    }
}
```

> ⚠️ **Padrão real do projeto:** Repositories de entidade DEVEM estender `BaseRepository<T>` (herda soft-delete, findById, findAll paginado com tenant filter, count). Queries customizadas (agregações, joins) são adicionadas como métodos extras. NÃO usar JdbcTemplate direto sem BaseRepository.

#### Service (com @Auditable e validações de RN)

```java
@Service
public class TenantService {
    private final TenantRepository tenantRepo;

    @Auditable(entityType = "TENANT", action = "CREATED")
    @Transactional
    public TenantResponse create(TenantCreateRequest req) {
        // RN04-02: validação de duplicidade
        if (tenantRepo.existsByNameCorporate(req.getNameCorporate())) {
            throw new DuplicateTenantException(req.getNameCorporate());
        }
        // RN04-01: status inicial PENDING_ONBOARDING
        Tenant tenant = Tenant.from(req);
        tenant.setStatus(TenantStatus.PENDING_ONBOARDING);
        tenantRepo.save(tenant);
        return TenantResponse.from(tenant);
    }
}
```

#### Controller (com @RequiresPermission)

```java
@RestController
@RequestMapping("/api/v1/tenants")
public class TenantController {
    private final TenantService tenantService;

    @PostMapping
    @RequiresPermission(resource = "TENANT", action = "create")
    public ResponseEntity<TenantResponse> create(@Valid @RequestBody TenantCreateRequest req) {
        return ResponseEntity.status(201).body(tenantService.create(req));
    }
}
```

---

## 6. Estratégia de Testes

### 6.1 Pirâmide para Sprint 3

| Nível | Ferramenta | Quantidade | Foco |
|:---|:---|:---:|:---|
| **Unitários** | JUnit 5 + Mockito | 24 | Services com repositories mockados. Cada RN testada positivo+negativo |
| **Integração** | Testcontainers + PostgreSQL 17 | 20 | Repositories com queries reais. Cenários de borda (409, 422, 403) |
| **Segurança** | MockMvc + JWT mock | 6 | RBAC: cada role × endpoint proibido → 403 |
| **E2E** | Playwright | 6 | Fluxos completos: dashboard, criar tenant, change-plan |

### 6.2 Testes Críticos (NÃO PODE FALHAR)

| ID | Cenário | Por que é crítico |
|:---|:---|:---|
| TC-F02-02-003 | ACTIVE→PENDING transição inválida → 422 | Se falhar, máquina de estados quebrada — tenants podem voltar a estado inconsistente |
| TC-F02-04-002 | Segunda assinatura ativa → 409 | Se falhar, tenant pode ter 2 planos simultâneos — billing incorreto |
| TC-F02-05-005 | UPDATE em audit_log → 403 | Se falhar, trilha de auditoria não é imutável — conformidade LGPD violada |
| TC-F01-01-005 | Soft-deleted tenants excluídos das métricas | Se falhar, dashboard mostra dados de tenants que não existem mais |
| TC-F02-04-009 | Concorrência: 2 assinaturas simultâneas não criam 2 ativas | Se falhar, race condition permite duplicação |

### 6.3 Testes de Performance (Dashboard)

```java
@Test
void dashboardSummary_shouldRespondUnder3Seconds_with1000Tenants() {
    // Seed: 1000 tenants com distribuição realista de status/planos
    // Warmup: 10 requisições
    // Medir: p95 de 100 requisições
    assertThat(p95Latency).isLessThan(Duration.ofSeconds(3));
}
```

---

## 7. Riscos e Mitigações Detalhadas

### 7.1 Matriz de Riscos

| Risco | Prob. | Impacto | Sinais de alerta | Mitigação | Plano B |
|:---|:---:|:---:|:---|:---|:---|
| **R1:** Dashboard lento (>3s) | Média | 🔴 Alto | Explain plan mostra full scan em T-016 | Índices via V004. Query de agregação otimizada com CTE. Teste de carga desde dia 1 | Materialized view refresh diário |
| **R2:** Sprint muito densa (23 tasks) | Alta | 🟡 Médio | Dia 5 e menos de 10 tasks concluídas | T-021 (alertas) já é Should. T-028 (email) pode ser Should se SMTP não disponível. Priorizar Must Have do M2 primeiro | Mover T-035/T-036 (auditoria) para Sprint 4 |
| **R3:** SMTP indisponível para T-028 | Média | 🟡 Médio | Mailhog não sobe no Docker Compose | Implementar `EmailService` com interface + mock para dev. Testes usam Mailhog. Produção requer configuração real | Enfileirar emails em tabela `email_queue` para envio futuro |
| **R4:** Race condition em change-plan (T-033) | Baixa | 🔴 Alto | Teste TC-F02-04-009 falha intermitente | `@Transactional` + `SELECT ... FOR UPDATE` na assinatura ativa. Teste de concorrência com `CountDownLatch` | Fila serializada por tenant (1 thread por tenant_id) |
| **R5:** Complexidade da máquina de estados (T-024) | Baixa | 🟡 Médio | Enum com >6 transições difícil de testar | Implementar como `Map<TenantStatus, Set<TenantStatus>>` com transições válidas. Teste parametrizado cobre todas as combinações | — |
| **R6:** Time reduzido (1 dev?) | Alta | 🔴 Alto | Velocidade abaixo de 2 tasks/dia | Paralelizar M2 e M3: Developer A faz M2 (dashboard), Developer B inicia M3 (entidades). Parear em T-033 (change-plan crítico) | Escopo reduzido: M2 completo + M3 apenas Tenant/Plan. Subscription e Audit vão para Sprint 4 |

### 7.2 Plano de Contingência por Semana

| Semana | Se atrasado... | Ação |
|:---|:---|:---|
| Fim da Semana 1 | <4 tasks M2 concluídas | Cortar T-021 (alertas). Focar T-016 a T-020 |
| Fim da Semana 2 | T-027 (TenantController) não funcional | Mover T-028 (email), T-032 a T-034 (assinaturas) para Sprint 4 |
| Fim da Semana 3 | <80% tasks concluídas | Mover T-035/T-036 (auditoria) para Sprint 4. Fechar sprint com M2 + CRUD Tenant/Plan |

---

## 8. Checkpoints e Marcos

### 8.1 Checkpoints Diários

| Checkpoint | Pergunta | Ação se "não" |
|:---|:---|:---|
| **Daily** | Build passa? (`mvn clean test`) | Corrigir antes de continuar |
| **Daily** | Testes novos passam? | Não acumular testes quebrados |
| **Daily** | Cobertura JaCoCo ≥ 80% nas classes novas? | Adicionar testes antes de nova feature |
| **A cada 3 tasks** | Review de código (self-review ou pair) | Bloquear PR até revisão OK |

### 8.2 Marcos Formais

| Marco | Data | Critério | Evidência |
|:---|:---|:---|:---|
| **M2 concluído** | ~Dia 4 | 4 endpoints dashboard + lista tenants funcionais. Testes M2 passando | Demo interno: dashboard carrega com 10+ tenants seed |
| **M3 Core** | ~Dia 9 | CRUD Tenant + Plan + Subscription funcionais. 11+ endpoints | Testes de integração passando para as 3 entidades |
| **Sprint 3 concluída** | 23/23 tasks. 56/56 testes. Cobertura ≥ 80%. Demo PO | Sprint review checklist preenchido |

---

## 9. Alocação de Esforço (Cenários)

### 9.1 Cenário A: 1 Desenvolvedor

| Semana | Foco | Tasks | Ritmo |
|:---|:---|:---|:---:|
| 1 | M2 — Dashboard | T-016 a T-023 (8 tasks) | 2 tasks/dia |
| 2 | M3 — Tenant + Plan | T-024 a T-031 (8 tasks) | 2 tasks/dia |
| 3 | M3 — Subscription + Audit + Testes | T-032 a T-038 (7 tasks) | 1.5 tasks/dia |

**Risco:** Sem buffer. Qualquer imprevisto empurra tasks para Sprint 4.

### 9.2 Cenário B: 2 Desenvolvedores (Recomendado)

| Dev | Semana 1 | Semana 2 | Semana 3 |
|:---|:---|:---|:---|
| **Dev A** | T-016, T-017, T-019 (Dashboard) | T-024, T-025, T-027 (Tenant CRUD) | T-028 (Email), T-035, T-036 (Auditoria) |
| **Dev B** | T-020 (TenantRepo), T-021 (Alertas), T-022, T-023 (Testes M2) | T-029, T-030, T-031 (Plans) | T-032, T-033, T-034 (Assinaturas) |
| **Pair** | — | T-033 (change-plan) | T-037, T-038 (Testes M3) |

**Vantagem:** 1 semana de buffer para imprevistos. Pair programming em T-033 (crítico).

---

## 10. Preparação do Ambiente de Desenvolvimento

### 10.1 Pré-requisitos (Verificar Antes de Começar)

```bash
# 1. Verificar Docker
docker ps && echo "Docker OK" || echo "Docker NÃO disponível"

# 2. Subir infraestrutura
docker compose up -d postgres keycloak mailhog

# 3. Verificar PostgreSQL
docker exec -it postgres psql -U fbso -d fbso_platform -c "SELECT COUNT(*) FROM fbso_platform.tenant;"

# 4. Executar migrations
./mvnw flyway:migrate -Dflyway.locations=filesystem:src/main/resources/db/migration

# 5. Criar migration V004 (índices de performance)
# Criar manualmente: src/main/resources/db/migration/V004__dashboard_performance_indexes.sql

# 6. Popular dados de teste (seed)
# Criar script: src/test/resources/db/seed/seed_sprint3.sql
# 50 tenants com distribuição: 35 ACTIVE, 10 PENDING_ONBOARDING, 5 SUSPENDED
# 3 planos: Básico (R$ 99/mês), Avançado (R$ 299/mês), Enterprise (R$ 999/mês)

# 7. Verificar build
./mvnw clean compile && echo "Build OK" || echo "Build FALHOU"
```

### 10.2 Ferramentas Recomendadas

| Ferramenta | Uso |
|:---|:---|
| **Mailhog** | Mock SMTP para T-028 (docker compose add `mailhog`) |
| **DBeaver / pgAdmin** | Verificar queries e índices durante desenvolvimento |
| **Postman / Bruno** | Testar endpoints manualmente antes de escrever testes |
| **JaCoCo Report** | `./mvnw jacoco:report` — verificar cobertura por pacote |

---

## 11. Critérios de Sucesso da Sprint

### 11.1 Must Have (Bloqueantes para considerar Sprint concluída)

- [ ] **M2:** Dashboard admin carrega em ≤3s (p95) com 1000 tenants seed
- [ ] **M2:** 4 endpoints de dashboard funcional (summary, evolution, accounts-by-status, accounts-by-plan)
- [ ] **M2:** Lista de tenants com paginação (25), busca textual (≥3 chars), filtros status/plano
- [ ] **M3:** CRUD Tenant: criar → PENDING_ONBOARDING, suspender exige motivo, reativar restaura permissões
- [ ] **M3:** CRUD Plan: criar/editar (versiona)/desativar (bloqueia se com assinantes, mínimo 1 ativo)
- [ ] **M3:** CRUD Subscription: 1 ativa por tenant, change-plan atômico sem gap
- [ ] **M3:** GET /audit funcional com filtros, paginação (25), registros imutáveis
- [ ] **Segurança:** Todos os endpoints anotados com `@RequiresPermission`
- [ ] **Testes:** 56/56 cenários passando (24 unit + 20 int + 6 E2E + 6 seg)
- [ ] **Cobertura:** JaCoCo ≥ 80% nas classes da sprint

### 11.2 Should Have (Desejável, não bloqueante)

- [ ] T-021: Alertas de onboarding >48h e assinatura suspensa
- [ ] T-028: Integração de email real (mock já cobre os testes)

### 11.3 Não Escopo (NÃO Implementar na Sprint 3)

- ❌ Qualquer endpoint de `/users`, `/permissions` (Sprint 4)
- ❌ Onboarding, dashboard do cliente (Sprint 5)
- ❌ CRUD `/business-units`, `/products` (Sprint 6)
- ❌ Frontend — escopo exclusivo do backend

---

## 12. Log de Decisões Durante o Desenvolvimento

> Decisões pré-sprint registradas em 16/07/2026 (Caveman Review v2.0). Novas decisões devem ser adicionadas durante a sprint.

| Data | Decisão | Contexto | Impacto |
|:---|:---|:---|:---|
| 16/07 | **D1:** ResultSetExtractor para T-016 | Múltiplas métricas em 1 query de agregação | Performance do dashboard |
| 16/07 | **D2:** Paginação offset-based | ~1000 tenants, volume baixo. Migrar para keyset se necessário | Complexidade do repository |
| 16/07 | **D3:** Campo `version` incrementado | Fase 0 não exige histórico. Migrar para tabela separada na Fase 1 | Schema |
| 16/07 | **D4:** Mailhog para dev/testes | Sem dependência externa. Configuração real apenas em staging | T-028, testes |
| 16/07 | **D5:** V004 enxuta (2 índices) | V002 já cobre 5/6 índices propostos | Migration reduzida |
| 16/07 | **D6:** Query SELECT antes do UPDATE | Mais simples e preciso que reflection para AuditAspect | T-035 |
| 16/07 | **D7:** Padrão BaseRepository | Todo repository de entidade estende BaseRepository. DashboardRepository é exceção (só agregação) | Templates de código |
| 16/07 | **D8:** SQL injection fix | TenantAwareDataSource migrado para PreparedStatement (defense-in-depth) | Segurança |
| 16/07 | **D9:** OpenAPI/Swagger | Adicionar springdoc-openapi para documentação dos 18 endpoints | DX |
| 17/07 | **D10 (T-039.DT-017):** V004 é OPCIONAL | `idx_tenant_segment` não existe na V002. Campo `segment` não é usado em WHERE de queries frequentes. Postergar para Sprint 4 se necessário | Migration V004 não será criada nesta sprint |
| 17/07 | **D11 (T-040.DT-019):** Day-by-day recalibrado | 12→15 dias úteis realista com Frentes 0+3. Frentes 0+1+2 concluídas em ~2.5 dias (antecipado) | Cronograma atualizado |
| 17/07 | **D12 (T-045.DT-046):** Testcontainers 1.21.4 | 1.20.6→1.21.4. CVE-2024-25710 (commons-compress transitiva) mitigada | Build compatível |
| 17/07 | **D13 (T-042.DT-025):** AccessDeniedException handler | Spring Security AccessDeniedException agora retorna 403 (antes caía no 500 genérico) | Segurança |

---

## Anexo A: Checklist de Início de Sprint

- [x] Ambiente Docker funcional (postgres + keycloak + mailhog)
- [x] Migration V004 (2 índices novos — V002 já cobre os demais) pronta para criar
- [ ] Script de seed com 50+ tenants pronto
- [x] `mvn clean compile` passa (33/33 testes BUILD SUCCESS)
- [x] 33 testes da Sprint 2 continuam passando
- [x] Decisões D1-D9 discutidas e registradas na Seção 12
- [ ] Dev(s) alocados e com ambiente pronto
- [ ] Mailhog acessível em `http://localhost:8025`
- [ ] **NOVO:** RowMapper para Tenant criado (dependência oculta de todos os repositories)
- [ ] **NOVO:** Padrão de teste de integração com Testcontainers validado (1 teste exemplo)
- [ ] **NOVO:** Dependência springdoc-openapi adicionada ao pom.xml
- [ ] Postman collection importada com endpoints base
- [ ] Repositório git com branch `feature/sprint-03-portal-admin` criada

---

🤖 *Planejamento gerado em 16/07/2026. Atualizado em 17/07/2026: Frentes 0+1+2 concluídas (35/42 tasks). 138 testes totais. Frente 3 em execução. Decisões D10-D13 registradas (V004 opcional, day-by-day recalibrado, Testcontainers 1.21.4, AccessDeniedException handler). Sprint 3 — a mais densa do projeto. O coração do Core Administrativo começa aqui.*
