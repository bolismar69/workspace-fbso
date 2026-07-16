# SPRINT-CARD: Sprint 3 — Portal Admin + Contas e Planos

- **Sprint:** 3 de 7
- **Status:** 🔄 Em andamento
- **Marco:** M2 (EP-01) + M3 (EP-02)
- **Datas:** 16/07/2026 → 31/08/2026 (início antecipado)
- **Duração:** 12 dias úteis
- **Responsável:** A definir
- **Documentos-mestre:** [TASKS.md](../../TASKS.md) v2.4 · [SPECS.md](../../SPECS.md) v1.6 · [TEST_PLAN.md](../../TEST_PLAN.md) v2.4 · [PRD.md](../../PRD.md) v1.5 · [ARCHITECTURE.md](../../ARCHITECTURE.md) v1.4

---

> 🚫 **BRANCH OBRIGATÓRIA:** Toda implementação deste sprint DEVE usar exclusivamente a branch `feature/sprint-03-portal-admin`. Antes de começar, execute:
> ```bash
> git checkout feature/sprint-03-portal-admin
> git branch --show-current  # deve exibir: feature/sprint-03-portal-admin
> ```
> 📖 Detalhes completos: [PRD.md §8.4](../../PRD.md#84-estratégia-de-branching--uma-branch-por-sprint)

## 🎯 Sprint Goal

**"Dashboard administrativo exibe métricas operacionais do SaaS em ≤3s. CRUD completo de Tenants, Planos e Assinaturas funcional. Auditoria registra 100% das ações administrativas e é consultável com filtros."**

> 🎯 **Primeiro sprint com entrega visível para o PO.**

---

## 📋 Sprint Backlog

### M2 — Portal Admin (EP-01)

| ID | Tarefa | Feature | Est. | Critério DONE |
|:---|:---|:---|:---:|:---|
| **T-016** | `DashboardRepository.java`: queries agregadas (contas ativas, por status, por plano, evolução temporal) | F01-01 | 2d | Queries verificadas. Explain plan sem full scan |
| **T-017** | `DashboardService.java`: lógica de métricas, filtro de período (7d, 30d, 90d, mês_atual, ano_atual). Padrão mês atual (RN01-02) | F01-01 | 1d | Filtro recalcula métricas. Período padrão = mês atual |
| **T-018** | DTOs: `DashboardSummaryResponse`, `EvolutionResponse`, `AccountsByStatusResponse`, `AccountsByPlanResponse` | F01-01 | 0.5d | JSON conforme contrato. Formatação R$. ISO 8601 |
| **T-019** | `DashboardController.java`: `GET /dashboard/admin/summary`, `/evolution`, `/accounts-by-status`, `/accounts-by-plan`. `@RequiresPermission(DASHBOARD, view)` | F01-01 | 1d | p95 ≤ 3s com 1000 tenants |
| **T-020** | `TenantRepository.java`: findAll paginado (25), filtros status/plano, busca textual (3+ chars, case-insensitive). Ordenação created_at DESC | F01-02 | 1.5d | Paginação funcional. Soft delete respeitado |
| **T-021** | Queries de alerta: onboarding >48h (RN03-01) + assinatura suspensa. Endpoint `GET /dashboard/admin/alerts` | F01-03 | 1.5d | Cards coloridos (WARNING/CRITICAL) |
| **T-022** | Testes unitários M2: `DashboardService`, `DashboardRepository` (mocks). Cobertura ≥ 80% | F01-01 a F01-03 | 1.5d | JUnit 5 + Mockito. Todos cenários de filtro |
| **T-023** | Testes integração M2: `DashboardRepository` com Testcontainers. Popular 10+ tenants | F01-01 a F01-03 | 1.5d | PostgreSQL real. Queries verificadas |

### M3 — Gestão de Clientes e Planos (EP-02)

| ID | Tarefa | Feature | Est. | Critério DONE |
|:---|:---|:---|:---:|:---|
| **T-024** | Entidades Tenant + DTOs. Mapa de transições de status (RN05-01) | F02-01, F02-02 | 1d | Transições válidas OK. Inválida → 422 |
| **T-025** | `TenantService.create()` status PENDING_ONBOARDING + `update()`. Validação razão social única (RN04-02) | F02-01 | 2d | Duplicada → 409. Auditoria @Auditable |
| **T-026** | `TenantService.suspend()` exige motivo (RN05-02). `reactivate()` restaura permissões (RN05-03) | F02-02 | 2d | Sem motivo → 400. Timeline status |
| **T-027** | `TenantController`: CRUD `/api/v1/tenants` + `/suspend`, `/reactivate`, `/resend-invite` | F02-01, F02-02 | 2d | 7 endpoints REST (GET list, GET by id, POST, PATCH, POST /suspend, POST /reactivate, POST /resend-invite). Bean Validation |
| **T-028** | Integração email: disparo na criação (US-009). Link único 7 dias (RN04-03). Reenvio `/resend-invite` | F02-01 | 2d | Email enviado. Link expira. Reenvio funcional |
| **T-029** | Entidade Plan + PlanModule + PlanRepository. Validação price > 0 (RN06-02) | F02-03 | 1d | Plano criado ACTIVE. PlanModule vinculado |
| **T-030** | `PlanService`: CRUD versionado. `deactivate()` preserva assinantes (RN06-01). Mínimo 1 plano ativo (RN06-03) | F02-03 | 2d | Edição gera nova versão. Último plano ativo → 422 |
| **T-031** | `PlanController`: CRUD `/api/v1/plans` + `POST /{id}/deactivate`. `@RequiresPermission` | F02-03 | 1d | CRUD completo. Desativado = "Descontinuado" |
| **T-032** | Entidade Subscription + Repository. 1 ativa por tenant (RN07-01) | F02-04 | 1d | Segunda ativa → 409 |
| **T-033** | `SubscriptionService`: criar, change-plan (RN07-02, RN07-03), suspender, reativar | F02-04 | 2d | Change-plan sem gap. Transação atômica |
| **T-034** | `SubscriptionController`: endpoints REST + `@RequiresPermission` | F02-04 | 1.5d | 4 endpoints. Validações OK |
| **T-035** | `AuditRepository` + `AuditService`: filtros período/ação/entidade. Paginação (25, max 100) (RN08-01, RN08-02) | F02-05 | 1.5d | Filtros funcionais. Imutável — UPDATE/DELETE → 403 |
| **T-036** | `AuditController`: `GET /api/v1/audit` com filtros. `@RequiresPermission(AUDIT, view)` | F02-05 | 1d | Admin vê tudo. Auditor vê tudo (leitura) |
| **T-037** | Testes unitários M3: todos os services. Cobrir RN05-01 a RN08-02 | F02-01 a F02-05 | 2d | ≥ 80%. Cada RN testada positivo+negativo |
| **T-038** | Testes integração M3: CRUD Tenant/Plan/Subscription/Audit com Testcontainers. Cenários de borda | F02-01 a F02-05 | 2d | PostgreSQL real. RN07-01, RN06-01, RN05-01, RN08-02 |

**Total:** 23 tarefas · ~35 dias-homem

---

## 📦 Features Entregues

| Feature | Descrição | RNs Cobertas |
|:---|:---|:---|
| **F01-01** | Dashboard Admin — Métricas | RN01-01, RN01-02, RN01-03 |
| **F01-02** | Lista de Contas | RN02-01, RN02-02 |
| **F01-03** | Alertas do Dashboard | RN03-01, RN03-02 |
| **F02-01** | Criar Tenant | RN04-01, RN04-02, RN04-03 |
| **F02-02** | Transições de Status | RN05-01, RN05-02, RN05-03 |
| **F02-03** | Configuração de Planos | RN06-01, RN06-02, RN06-03 |
| **F02-04** | Vinculação de Assinaturas | RN07-01, RN07-02, RN07-03 |
| **F02-05** | Auditoria | RN08-01, RN08-02 |

---

## ✅ Definition of Done (Sprint-Level)

- [ ] Dashboard admin carrega em ≤3s (p95) com 1000 tenants
- [ ] CRUD Tenant funcional: criar → PENDING_ONBOARDING, suspender → motivo obrigatório
- [ ] CRUD Plan funcional: edição versiona, desativação bloqueada se com assinantes
- [ ] CRUD Subscription funcional: 1 ativa por tenant, change-plan atômico
- [ ] GET /audit funcional com filtros de período, ação, entidade
- [ ] Todos os endpoints anotados com `@RequiresPermission`
- [ ] 55 cenários de teste (unit + integração) da suite extraída
- [ ] Cobertura JaCoCo ≥ 80%

---

## ⚠️ Riscos e Bloqueadores

| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---:|:---|
| Sprint mais densa do projeto (23 tarefas, 35d-homem) | Alta | Alto | Priorizar Must Have. T-021 já é Should (F01-03). T-028 (email) pode ser negociado como Should |
| Integração de email (T-028) depende de SMTP externo | Média | Médio | Mock SMTP com GreenMail ou Mailhog nos testes. Configuração real apenas em staging |
| Performance do dashboard com 1000 tenants | Média | Alto | Índices de desempenho (V002). Query de agregação otimizada. Teste de carga desde o início |

---

## 🔗 Dependências

- **Pré-requisitos:** ✅ Sprint 2 concluída (pipeline de segurança). ✅ T-020 (TenantRepository) é pré-requisito para M3.
- **Sucessor:** Sprint 4 (RBAC) — depende de UserRepository e PermissionRepository.

---

## 📊 Métricas da Sprint

| Métrica | Meta |
|:---|:---:|
| Tasks completadas | 23/23 |
| Endpoints REST | 17 novos |
| RNs implementadas | 20 |
| Cenários de teste | 55 |
| Cobertura JaCoCo | ≥ 80% |

---

🤖 *Gerado a partir de TASKS.md v2.4. Sprint 3 iniciada em 16/07/2026. Sprint mais densa do projeto — o coração do Core Administrativo.*
