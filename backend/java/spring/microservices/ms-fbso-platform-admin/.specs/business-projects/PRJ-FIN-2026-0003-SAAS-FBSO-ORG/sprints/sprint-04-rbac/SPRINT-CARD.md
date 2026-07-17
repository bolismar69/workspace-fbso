# SPRINT-CARD: Sprint 4 — Governança de Acessos (RBAC)

- **Sprint:** 4 de 7
- **Marco:** M4 (EP-03)
- **Datas:** 31/08/2026 → 15/09/2026
- **Duração:** 11 dias úteis
- **Responsável:** A definir
- **Documentos-mestre:** [TASKS.md](../../TASKS.md) · [SPECS.md](../../SPECS.md) · [PRD.md](../../PRD.md)
- **Execução:** [Frente 0 — Relatório](./SPRINT-4-EXECUTION-REPORT-Frente-0.md) · [Débitos Técnicos](./IDENTIFIED-TECHNICAL-DEBT-sprint-04-rbac.md)

---

> 🚫 **BRANCH OBRIGATÓRIA:** Toda implementação deste sprint DEVE usar exclusivamente a branch `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-04-rbac`. Antes de começar, execute:
> ```bash
> git checkout PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-04-rbac
> git branch --show-current  # deve exibir: PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-04-rbac
> ```
> 📖 Detalhes completos: [PRD.md §8.4](../../PRD.md#84-estratégia-de-branching--uma-branch-por-sprint)

---

## 🎯 Sprint Goal

**"4 papéis (Admin Tenant, Gerente BU, Operador BU, Auditor) aplicados com matriz de permissões RN10-01. Gestão de usuários com convite por e-mail. Vinculação Usuário × Unidade × Módulo. Bloqueio de acesso direto com 403 amigável em PT-BR."**

---

## 🗺️ Estrutura de Frentes — Visão Geral

> A Sprint 4 está organizada em **7 frentes sequenciais**. Cada frente agrupa tasks relacionadas e deve ser concluída antes de avançar para a próxima.

| # | Frente | Features | Tasks | Estimativa | Status | Progresso |
|:---:|:---|:---|:---:|:---:|:---:|:---:|
| **0** | Correções Pré-Sprint (Bloqueantes) | — | 20 | 16-24h | ✅ Concluída | 20/20 (100%) |
| **1** | Gestão de Usuários | F03-01 | 3 | 3.5d | ✅ Concluída | 3/3 (100%) |
| **2** | Entidades RBAC + UserPermission | F03-02 | 2 | 2.5d | ⬜ Pendente | 0/2 (0%) |
| **3** | PermissionService + API | F03-02, F03-03 | 2 | 3d | ⬜ Pendente | 0/2 (0%) |
| **4** | Integração RBAC + Segurança + Testes | F03-04 | 4 | 5.5d | ⬜ Pendente | 0/4 (0%) |
| **5** | Correções Recomendadas (Durante-Sprint) | — | 15 | 28-32h | ⬜ Pendente | 0/15 (0%) |
| **5b** | Infra + DevX (Backlog Sprint 3) | — | 2 | 1.5h | ⬜ Pendente | 0/2 (0%) |
| **TOTAL** | **7 frentes** | **4 features** | **48** | **~26d** | **Frentes 0-1 ✅** | **23/48 (48%)** |

---

## 📦 Frentes Detalhadas

### Frente 0 ✅ — Correções Pré-Sprint (Bloqueantes)

- **Objetivo:** Resolver 20 débitos técnicos impeditivos ANTES de iniciar as features RBAC
- **Status:** ✅ **CONCLUÍDA em 17/07/2026**
- **Evidência:** [SPRINT-4-EXECUTION-REPORT-Frente-0.md](./SPRINT-4-EXECUTION-REPORT-Frente-0.md)
- **Origem:** [IDENTIFIED-TECHNICAL-DEBT-sprint-04-rbac.md](./IDENTIFIED-TECHNICAL-DEBT-sprint-04-rbac.md) — DT-048 a DT-067

| ID | Tarefa | Débito | Est. | Status |
|:---|:---|:---:|:---:|:---:|
| **T-096.DT-048** | Adicionar Caffeine + CacheConfig | DT-048 | 1h | ✅ |
| **T-097.DT-049** | Adicionar REST Assured (test scope) | DT-049 | 30min | ✅ |
| **T-098.DT-050** | Estratégia merge JWT×DB: banco como fonte primária de roles | DT-050 | 3h | ✅ |
| **T-099.DT-051** | Abandonar cache TTL 5min: matriz carregada em memória | DT-051 | 2h | ✅ |
| **T-100.DT-052** | Refatorar RbacAspect: inject PermissionService, remover Sets | DT-052 | 4h | ✅ |
| **T-101.DT-053** | Criar User.java + UserRepository.java + UserRowMapper.java | DT-053 | 2h | ✅ |
| **T-102.DT-054** | Criar ResourceAction.java + RoleResource.java entities | DT-054 | 1h | ✅ |
| **T-103.DT-055** | Criar BusinessUnit.java entity (mínima) | DT-055 | 1h | ✅ |
| **T-104.DT-056** | Migration V006: FK user_permission → business_unit + U006 | DT-056 | 30min | ✅ |
| **T-105.DT-057** | Migration V004: seed resource_action + role_resource (RN10-01) + U004 | DT-057 | 1h | ✅ |
| **T-106.DT-058** | Migrar RbacAspect de strings para Role enum | DT-058, DT-035 | 1h | ✅ |
| **T-107.DT-059** | Refatorar TenantController.list() → TenantService | DT-059 | 1h | ✅ |
| **T-108.DT-060** | Adicionar JwtValidators.createDefaultWithIssuer() | DT-060 | 1h | ✅ |
| **T-109.DT-061** | FORCE ROW LEVEL SECURITY nas 4 tabelas (V003) | DT-061 | 2h | ✅ |
| **T-110.DT-062** | Validar tenant_id URL vs JWT no SubscriptionService | DT-062 | 2h | ✅ |
| **T-111.DT-063** | Corrigir connection leak no TenantAwareDataSource | DT-063 | 1h | ✅ |
| **T-112.DT-064** | Alinhar matriz RBAC: SPRINT-CARD restritiva + breaking change doc | DT-064 | 1h | ✅ |
| **T-113.DT-065** | Corrigir diagrama TASKS.md §3 | DT-065 | 30min | ✅ |
| **T-114.DT-066** | Corrigir header SPRINT-TEST-SUITE: 19→27 cenários | DT-066 | 15min | ✅ |
| **T-115.DT-067** | Adicionar BU scope check no PermissionService | DT-067 | 3h | ✅ |

**Artefatos criados na Frente 0 (13 novos arquivos):**
- `config/CacheConfig.java`, `entity/User.java`, `entity/ResourceAction.java`, `entity/RoleResource.java`, `entity/BusinessUnit.java`
- `repository/UserRepository.java`, `repository/rowmapper/UserRowMapper.java`
- `service/PermissionService.java`
- `db/migration/V004__seed_rbac_matrix.sql`, `U004__rollback_rbac_seed.sql`
- `db/migration/V006__add_fk_user_permission_bu.sql`, `U006__drop_fk_user_permission_bu.sql`

> ⚠️ **IMPORTANTE:** Várias tasks das Frentes 1-4 têm artefatos parcialmente criados na Frente 0. Ver notas de sobreposição abaixo.

---

### Frente 1 ✅ — Gestão de Usuários (F03-01) — CONCLUÍDA 17/07/2026

- **Objetivo:** CRUD completo de usuários com convite por e-mail, desativação e reativação
- **Dependências:** Frente 0 (UserRepository + User entity criados)
- **Plano:** [SPRINT-DEVELOPMENT-PLANNING-Frente-1.md](./SPRINT-DEVELOPMENT-PLANNING-Frente-1.md)
- **RNs:** RN09-01 (convite expira 7 dias), RN09-02 (email único por tenant), RN09-03 (não autodesativar)

| ID | Tarefa | Feature | Est. | Critério DONE |
|:---|:---|:---:|:---:|:---|
| **T-046** | Entidade User + UserRepository. `findByEmailAndTenant`. Email único por tenant ativo (RN09-02, índice parcial) | F03-01 | 1d | ⚠️ Base criada na Frente 0 (T-101). Complementar: UserStatus enum, índices, testes unitários |
| **T-047** | `UserService`: convite (email único RN09-02), desativar (não autodesativar RN09-03), reativar. Convite expira 7 dias (RN09-01) | F03-01 | 1.5d | Duplicado → 409. Autodesativar → 422 |
| **T-048** | `UserController`: CRUD `/api/v1/users` + `POST /{id}/deactivate`. `@RequiresPermission(USER, ...)` | F03-01 | 1d | Lista exibe nome, email, role, status, BUs vinculadas |

**Ordem interna:** T-046 → T-047 → T-048 (sequencial)
**Checkpoint:** `mvn test` após T-048 — todos os endpoints User operacionais

---

### Frente 2 ⬜ — Entidades RBAC + UserPermission (F03-02)

- **Objetivo:** Seed data da matriz RN10-01 + entidade de vínculo Usuário×BU×Role
- **Dependências:** Frente 0 (ResourceAction, RoleResource, BusinessUnit criados) + Frente 1 (User)

| ID | Tarefa | Feature | Est. | Critério DONE |
|:---|:---|:---:|:---:|:---|
| **T-049** | Entidades ResourceAction + RoleResource. Seed data com matriz RN10-01 (4 roles × 8 resources). Migration V004 | F03-02 | 1d | ⚠️ Entities criadas na Frente 0 (T-102). Seed criado (T-105). Complementar: validar seed carrega em todos os ambientes, testes |
| **T-050** | Entidade UserPermission (user_id, business_unit_id, role). UNIQUE(user_id, business_unit_id). Tabela ponte | F03-02, F03-03 | 1.5d | Constraint única. Admin tenant acesso implícito a todas BUs |

**Ordem interna:** T-049 → T-050 (T-050 referencia resources de T-049)
**Checkpoint:** `mvn test` — seed carrega, UserPermission funcional

---

### Frente 3 ⬜ — PermissionService + API (F03-02, F03-03)

- **Objetivo:** Serviço de atribuição/revogação de permissões + API REST
- **Dependências:** Frente 2 (UserPermission) + Frente 0 (PermissionService base)

| ID | Tarefa | Feature | Est. | Critério DONE |
|:---|:---|:---:|:---:|:---|
| **T-051** | `PermissionService`: atribuir/revogar, vincular BU, gerenciar módulos. Admin acesso implícito (RN11-01, RN11-02). Efeito imediato (RN11-03) | F03-02, F03-03 | 2d | ⚠️ Base criada na Frente 0 (T-100). Complementar: assignRole(), revokeRole(), validateBusinessUnitAccess() |
| **T-052** | `PermissionController`: `GET /users/{uid}/permissions`, `PUT /users/{uid}/permissions`. `@RequiresPermission(PERMISSION, ...)` | F03-02, F03-03 | 1d | GET retorna atuais. PUT atualiza vínculos. Auditoria registrada |

**Ordem interna:** T-051 → T-052 (sequencial)
**Checkpoint:** `mvn test` — atribuição/revogação funcional, efeito imediato validado

---

### Frente 4 ⬜ — Integração RBAC + Segurança + Testes (F03-04)

- **Objetivo:** Integração final do RBAC, resposta 403 amigável, bateria completa de testes
- **Dependências:** Todas as frentes acima (1, 2, 3)

| ID | Tarefa | Feature | Est. | Critério DONE |
|:---|:---|:---:|:---:|:---|
| **T-053** | Integrar `RbacAspect` com `RoleResource` do banco. Cache de matriz em memória (sem TTL — RN11-03) | F03-04 | 1.5d | ⚠️ RbacAspect refatorado na Frente 0 (T-100). Complementar: testes de integração, remover fallback JWT |
| **T-054** | Garantir 403 padrão: `{"title":"Acesso negado","detail":"Você não tem permissão para acessar esta área.","status":403}` (RN12-02) | F03-04 | 0.5d | Resposta sempre nesse formato. PT-BR. RFC 7807 |
| **T-055** | Testes unitários M4: `UserService`, `PermissionService`. RN09-03, RN10-01, RN11-01, RN11-02 | F03-01 a F03-03 | 1.5d | ≥ 80% cobertura. Cada RN positivo+negativo |
| **T-056** | Testes segurança RBAC: cada papel × endpoint proibido → 403. Teste parametrizado com REST Assured + Testcontainers | F03-02, F03-04 | 2d | Matriz RN10-01 validada. 20+ combinações papel×endpoint |

**Ordem interna:** T-053 → T-054 (paralelo com T-055) → T-056 (sequencial após todos)
**Checkpoint final:** `mvn clean verify` — todos os 27 cenários verdes, matriz RN10-01 100% validada

---

### Frente 5 ⬜ — Correções Recomendadas (Durante-Sprint)

- **Objetivo:** Resolver 15 débitos não-bloqueantes de segurança, qualidade e documentação
- **Dependências:** Nenhuma (paralelizável com Frentes 1-4)
- **Origem:** [IDENTIFIED-TECHNICAL-DEBT](./IDENTIFIED-TECHNICAL-DEBT-sprint-04-rbac.md) — DT-068 a DT-082

| ID | Tarefa | Débito | Est. | Critério DONE |
|:---|:---|:---:|:---:|:---|
| **T-116.DT-068** | Atualizar PostgreSQL driver 42.7.10→42.7.11 (CVE-2026-42198, CVSS 7.5) | DT-068 | 30min | Driver atualizado. Build passa |
| **T-117.DT-069** | Migrar Flyway 10.22.0→12.11.0 (10→11→12) | DT-069, DT-045 | 4h | Flyway 12.11.0. Migrations re-executadas |
| **T-118.DT-070** | Corrigir AuditAspect.parseEntityId(): null (não UUID.randomUUID) | DT-070 | 1h | Sem entity_id fantasma |
| **T-119.DT-071** | Criar @WebMvcTest para 4 controllers sem cobertura | DT-071 | 4h | 4 controllers >80% cobertura |
| **T-120.DT-072** | Corrigir RLSIsolationTest: remover @Disabled | DT-072 | 3h | RLS isolation tests executando |
| **T-121.DT-073** | Corrigir BaseRepositoryTest.save/update @Disabled | DT-073 | 2h | save/update cobertos |
| **T-122.DT-074** | Adicionar PERMISSION, ROLE no resolveTableName() do AuditAspect | DT-074 | 30min | Entidades RBAC auditadas |
| **T-123.DT-075** | Adicionar handler TenantIsolationException → 403 | DT-075 | 30min | TenantIsolationException → 403 |
| **T-124.DT-076** | Migrar para Converter<Jwt, AbstractAuthToken> customizado | DT-076, DT-030 | 4h | JWT decodificado 1× por request |
| **T-125.DT-077** | Reescrever RbacAspectTest com @SpringBootTest + @MockBean | DT-077 | 3h | Testes validam DB-backed RBAC |
| **T-126.DT-078** | Adicionar SELECT ... FOR UPDATE no PlanService.deactivate() | DT-078 | 1h | Race condition eliminada |
| **T-127.DT-079** | Try-catch DateTimeParseException → 400 no AuditService | DT-079 | 15min | Data inválida → 400 |
| **T-128.DT-080** | Documentar GET /plans/admin e GET /plans/{id} na SPECS.md §4.1 | DT-080 | 15min | API Contract First restaurado |
| **T-129.DT-081** | Corrigir SPRINT-CARD ref "V003" → "V004" | DT-081 | 5min | Referência correta |
| **T-130.DT-082** | Remover dependência falsa "Sprint 3 (UserRepository)" | DT-082 | 5min | Dependências reais |

**Ordem:** Paralelizável. Prioridade: T-116, T-118, T-122, T-123, T-127, T-128, T-129, T-130 primeiro (quick wins <1h). T-117, T-119, T-120, T-121, T-124, T-125, T-126 depois.

---

### Frente 5b ⬜ — Infra + DevX (Backlog Sprint 3)

- **Objetivo:** Ambiente de desenvolvimento local com 1 comando + dados realistas para testes
- **Dependências:** Nenhuma (independente de todas as frentes)

| ID | Tarefa | Débito | Est. | Critério DONE |
|:---|:---|:---:|:---:|:---|
| **T-131.DT-042** | Criar docker-compose.yml com PostgreSQL 17 + Keycloak 26 + MailHog | DT-042 | 1h | `docker compose up -d` → ambiente pronto |
| **T-132.DT-043** | Criar script seed-dev.sql com 50+ tenants, 3 planos, 4 roles | DT-043 | 30min | Dados realistas para dev |

---

## ⚠️ Notas de Sobreposição Frente 0 × Features

> A Frente 0 criou artefatos que originalmente fariam parte das tasks de feature. Isso reduz o escopo de várias tasks:

| Task Feature | Artefato já criado na Frente 0 | O que falta fazer |
|:---|:---|:---|
| **T-046** (User entity) | `User.java`, `UserRepository.java`, `UserRowMapper.java` (T-101) | Complementar com `UserStatus` enum, testes unitários, índices |
| **T-049** (ResourceAction + RoleResource) | `ResourceAction.java`, `RoleResource.java` (T-102) + seed V004 (T-105) | Validar seed em múltiplos ambientes, testes |
| **T-051** (PermissionService) | `PermissionService.java` com matriz carregada, `validateBusinessUnitAccess()` (T-100, T-115) | `assignRole()`, `revokeRole()`, testes |
| **T-053** (RbacAspect integração) | `RbacAspect` refatorado DB-backed, sem Sets hardcoded, usa `Role` enum (T-100, T-106) | Remover fallback JWT, testes de integração |

> **Regra:** Antes de iniciar cada task de feature, ler o código já existente da Frente 0 para evitar retrabalho.

---

## 📊 Matriz RN10-01 (Referência Rápida)

| Papel | Dashboard | Tenants | Plans | Users | Permissions | BUs | Products | Audit |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| **Admin Tenant** | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| **Gerente BU** | ❌ | ❌ | ❌ | ❌ | ❌ | ✅ edit | ✅ edit | ❌ |
| **Operador BU** | ❌ | ❌ | ❌ | ❌ | ❌ | ✅ view | ✅ view | ❌ |
| **Auditor** | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ | ✅ view |

> ⚠️ **BREAKING CHANGE — MANAGER_BU:** A matriz RN10-01 restritiva remove acesso de MANAGER_BU a Dashboard, Tenants, Plans e Subscriptions que eram temporariamente permitidos no RbacAspect hardcoded pré-Frente 0. MANAGER_BU passa a acessar apenas BUSINESS_UNIT e PRODUCT_SERVICE. **Comunicar ao frontend e ao PO antes do deploy.**

---

## 🔗 Ordem de Desenvolvimento

```
FASE 3 — M4 (EP-03): Sprint 4 — RBAC
│
├── FRENTE 0 ✅ (20 tasks) — CONCLUÍDA 17/07/2026
│   └── T-096.DT-048 a T-115.DT-067
│   └── Bloqueava o início de TODAS as features
│
├── FRENTE 1 ⬜ (3 tasks) — INÍCIO APÓS Frente 0
│   └── T-046 → T-047 → T-048
│   └── Depende de: User entity + UserRepository (Frente 0)
│
├── FRENTE 2 ⬜ (2 tasks) — INÍCIO APÓS Frente 1
│   └── T-049 → T-050
│   └── Depende de: ResourceAction/RoleResource (Frente 0) + User (Frente 1)
│
├── FRENTE 3 ⬜ (2 tasks) — INÍCIO APÓS Frente 2
│   └── T-051 → T-052
│   └── Depende de: UserPermission (Frente 2) + PermissionService base (Frente 0)
│
├── FRENTE 4 ⬜ (4 tasks) — INÍCIO APÓS Frentes 1, 2, 3
│   └── T-053 ∥ T-054 ∥ T-055 → T-056
│   └── Depende de: TODAS as frentes acima
│
├── FRENTE 5 ⬜ (15 tasks) — PARALELO com Frentes 1-4
│   └── Quick wins primeiro (<1h): T-116, T-118, T-122, T-123, T-127..T-130
│   └── Depois: T-117, T-119, T-120, T-121, T-124, T-125, T-126
│
└── FRENTE 5b ⬜ (2 tasks) — INDEPENDENTE (qualquer momento)
    └── T-131 ∥ T-132
```

### Cadeia Crítica (Caminho Mais Longo)

```
Frente 0 ✅ (concluída)
  → Frente 1: T-046 (1d) → T-047 (1.5d) → T-048 (1d)
  → Frente 2: T-049 (1d) → T-050 (1.5d)
  → Frente 3: T-051 (2d) → T-052 (1d)
  → Frente 4: T-053 (1.5d) → T-056 (2d)
TOTAL CRÍTICO: ~10.5d (dos 11 dias úteis disponíveis)
```

---

## ✅ Definition of Done (Sprint-Level)

### Frente 0 (Pré-Sprint)
- [x] Caffeine + REST Assured no classpath
- [x] 4 novas entities compilando (User, ResourceAction, RoleResource, BusinessUnit)
- [x] Migration V004 (seed RBAC) + V006 (FK) + rollbacks testados
- [x] RbacAspect DB-backed (sem Sets hardcoded, usa Role enum)
- [x] PermissionService com matriz carregada em memória
- [x] JWT issuer validation ativa
- [x] FORCE ROW LEVEL SECURITY nas 4 tabelas
- [x] Connection leak eliminado
- [x] Tenant isolation bypass corrigido
- [x] Matriz RN10-01 unificada (SPRINT-CARD = código = seed)

### Features (Frentes 1-4)
- [ ] CRUD Usuários funcional (convidar, desativar, reativar)
- [ ] Admin não pode desativar a si mesmo (RN09-03 → 422)
- [ ] Seed data da matriz RN10-01 carregada corretamente em todos os ambientes
- [ ] `@RequiresPermission` integrado com banco (não hardcoded)
- [ ] Teste parametrizado: 20+ combinações papel × endpoint proibido → 403
- [ ] Resposta 403 sempre no formato padrão (PT-BR, RFC 7807)
- [ ] Usuário sem BU vinculada → 403 (RN11-01)
- [ ] Usuário sem módulo → 403 (RN11-02)
- [ ] Cobertura ≥ 80% nos serviços M4

### Correções Recomendadas (Frente 5)
- [ ] PostgreSQL driver sem CVE conhecida
- [ ] 4 controllers com @WebMvcTest
- [ ] RLSIsolationTest executando sem @Disabled
- [ ] BaseRepository.save/update cobertos
- [ ] Entidades RBAC auditadas (AuditAspect)
- [ ] TenantIsolationException → 403 (não 500)
- [ ] API Contract First restaurado (SPECS.md)

---

## ⚠️ Riscos e Bloqueadores

| Risco | Prob. | Impacto | Mitigação | Status |
|:---|:---:|:---:|:---|:---:|
| ~~Cache de permissões servir dado stale~~ | — | — | ~~Abandonado na Frente 0 (T-099). Matriz em memória sem TTL~~ | ✅ Resolvido |
| ~~RbacAspect hardcoded vs DB~~ | — | — | ~~Refatorado na Frente 0 (T-100). PermissionService injetado~~ | ✅ Resolvido |
| ~~JWT sem validação de issuer~~ | — | — | ~~Corrigido na Frente 0 (T-108). JwtValidators.createDefaultWithIssuer()~~ | ✅ Resolvido |
| ~~RLS sem FORCE — isolamento ineficaz~~ | — | — | ~~Corrigido na Frente 0 (T-109). FORCE ROW LEVEL SECURITY~~ | ✅ Resolvido |
| MANAGER_BU perde acesso na migração | Média | Alto | Documentado. Comunicar frontend + PO. Deploy único (seed + RbacAspect) | ⚠️ Monitorar |
| Seed RBAC não carregar em todos os ambientes | Baixa | Alto | Migration V004 com seed. Testar em dev, staging e CI | ⬜ Pendente |
| Sobrecarga da cadeia crítica (10.5d em 11d) | Média | Alto | Frentes 5 e 5b são paralelizáveis. Quick wins não consomem dia completo | ⚠️ Monitorar |

---

## 📊 Métricas da Sprint

| Métrica | Original | Atualizada | Meta |
|:---|:---:|:---:|:---:|
| Tasks total | 11 | **48** | 48/48 |
| Tasks concluídas (Frente 0) | — | **20** ✅ | — |
| Tasks pendentes (Frentes 1-5b) | — | **28** | 28/28 |
| Features | 4 | 4 | F03-01 a F03-04 |
| Combinações papel×endpoint testadas | 20+ | 20+ | Matriz RN10-01 validada |
| Cenários de teste RBAC | 19→27 | **27** ✅ | Todos verdes |
| Resposta 403 padrão | 100% | 100% | RFC 7807 PT-BR |
| Cobertura de código | ≥ 80% | ≥ 80% | JaCoCo services M4 |
| Débitos resolvidos | — | **20** (Frente 0) | +15 (Frente 5) = 35 total |

---

## 📦 Features Entregues

| Feature | Descrição | RNs Cobertas | Frentes |
|:---|:---|:---|:---|
| **F03-01** | Gestão de Usuários | RN09-01, RN09-02, RN09-03 | Frente 1 |
| **F03-02** | Matriz de Permissões RBAC | RN10-01 | Frentes 2, 3 |
| **F03-03** | Vinculação Usuário × Unidade × Módulo | RN11-01, RN11-02, RN11-03 | Frentes 2, 3 |
| **F03-04** | Acesso Condicional (403 Amigável) | RN12-01, RN12-02 | Frente 4 |

---

## 🔗 Dependências

- **Pré-requisitos (já resolvidos):** Sprint 2 (RbacAspect, @RequiresPermission). Sprint 3 (UserRepository — criado na Frente 0).
- **Frente 0 ✅:** 20 correções bloqueantes — todas concluídas.
- **Sucessor:** Sprint 5 (Portal Cliente) — depende de UserRepository e PermissionRepository desta sprint.

---

🤖 *Documento reestruturado em 17/07/2026. A Frente 0 (20 tasks) está concluída. As Frentes 1-5b (28 tasks) compõem o backlog ativo da sprint. A matriz RN10-01 é o coração desta sprint — validar cada combinação.*
