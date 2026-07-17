# SPRINT-REVIEW: Sprint 4 — Governança de Acessos (RBAC)

- **Sprint:** 4 de 7
- **Marco:** M4 (EP-03)
- **Data da Review:** 15/09/2026
- **Participantes:** Time Técnico, Tech Lead, **Product Owner** 🎯
- **Features:** 4 (F03-01 a F03-04)
- **Status atual:** Frente 0 ✅ concluída (17/07/2026) · Frentes 1-5b ⬜ pendentes
- **Docs relacionados:** [SPRINT-CARD](./SPRINT-CARD.md) · [Exec Report Frente 0](./SPRINT-4-EXECUTION-REPORT-Frente-0.md) · [Tech Debt](./IDENTIFIED-TECHNICAL-DEBT-sprint-04-rbac.md)

---

## 📋 Checklists de Review por Frente

### Frente 0 ✅ — Correções Pré-Sprint (CONCLUÍDA 17/07/2026)

> **Não requer demonstração ao PO.** Estas foram correções técnicas internas. Evidências no [Relatório de Execução](./SPRINT-4-EXECUTION-REPORT-Frente-0.md).

- [x] Caffeine + REST Assured adicionados ao pom.xml
- [x] 4 novas entities compilando (User, ResourceAction, RoleResource, BusinessUnit)
- [x] Migration V004 (seed RBAC) + V006 (FK) + rollbacks
- [x] RbacAspect refatorado: DB-backed, sem Sets hardcoded, usa Role enum
- [x] PermissionService com matriz carregada em memória + BU scope check
- [x] JWT issuer validation ativa no SecurityConfig
- [x] FORCE ROW LEVEL SECURITY nas 4 tabelas (subscription, user, business_unit, audit_log)
- [x] Tenant isolation bypass corrigido no SubscriptionService
- [x] Connection leak eliminado no TenantAwareDataSource
- [x] TenantController.list() passa pelo TenantService (RBAC aplicado)
- [x] Diagrama TASKS.md corrigido (tasks reais)
- [x] Métricas SPRINT-TEST-SUITE corrigidas (19→27)
- [x] Breaking change MANAGER_BU documentado

**Build:** `mvn compile` esperado SUCCESS · **Arquivos:** 13 criados, 12 modificados

---

### Frente 1 — Gestão de Usuários (F03-01)

- [ ] **Convidar:** Admin Tenant convida usuário com email → status INVITE_PENDING
- [ ] **Validação:** Tentar mesmo email no mesmo tenant → 409
- [ ] **Autodesativação:** Admin tentar desativar a si mesmo → 422 "Um administrador não pode desativar a si mesmo"
- [ ] **Reativação:** Usuário desativado pode ser reativado
- [ ] **Lista:** Usuários exibidos com nome, email, papel, status, BUs vinculadas
- [ ] **Filtro:** Lista filtrável por status (ACTIVE, INACTIVE, INVITE_PENDING)

---

### Frente 2 — Entidades RBAC (F03-02)

- [ ] **Seed carregado:** Migration V004 executada em dev, staging e CI
- [ ] **Matriz completa:** 4 roles × 8 resources × 4 actions conforme RN10-01
- [ ] **UserPermission:** Tabela ponte com UNIQUE(user_id, business_unit_id)
- [ ] **Admin implícito:** Admin Tenant acessa todas BUs sem registros em user_permission

---

### Frente 3 — PermissionService + API (F03-02, F03-03)

- [ ] **Atribuir:** Vincular usuário a BU-A com papel Gerente
- [ ] **Revogar:** Remover permissão → efeito imediato na próxima requisição (RN11-03)
- [ ] **API:** GET /users/{uid}/permissions retorna vínculos atuais
- [ ] **API:** PUT /users/{uid}/permissions atualiza vínculos com auditoria
- [ ] **Isolamento:** Usuário BU-A vê apenas produtos da BU-A
- [ ] **Tentativa de acesso:** Usuário BU-A tenta acessar produto BU-B por ID direto → 404

---

### Frente 4 — Integração RBAC + Segurança (F03-04)

#### Demonstração por Papel

- [ ] **Login como Admin Tenant:** Acessa tudo — dashboard, tenants, planos, usuários, permissões, BUs, produtos
- [ ] **Login como Gerente BU:** Vê BUs e produtos da sua unidade. Pode criar/editar. Não vê tenants, planos, auditoria
- [ ] **Login como Operador BU:** Apenas leitura de BUs e produtos. Tentar editar → 403
- [ ] **Login como Auditor:** Apenas leitura de auditoria. Tentar criar qualquer coisa → 403

> 🎬 **Script:** "Vou fazer login como cada um dos 4 papéis. Reparem que o menu lateral muda. O Operador tenta criar um produto... 403."

#### 403 Amigável

- [ ] **Acesso direto:** Operador digita URL `/admin/plans/create` → 403
- [ ] **Formato:** `{"title":"Acesso negado","detail":"Você não tem permissão para acessar esta área.","status":403}`
- [ ] **Sem vazamento:** 403 (nunca 404 — não revela existência do recurso, RN12-01)
- [ ] **Sem stack trace:** Nenhum detalhe técnico na resposta

#### Testes

- [ ] **Unitários:** UserService, PermissionService ≥ 80% cobertura
- [ ] **Parametrizado:** 20+ combinações papel × endpoint proibido → 403
- [ ] **Segurança:** Matriz RN10-01 100% validada com REST Assured + Testcontainers

---

## 📋 Pontos de Verificação (PO)

| # | Verificação | Frente | Status |
|:---|:---|:---:|:---:|
| 1 | Correções técnicas pré-sprint concluídas | Frente 0 | ✅ |
| 2 | Convidar usuário → email enviado | Frente 1 | ⬜ |
| 3 | Email duplicado → bloqueado (409) | Frente 1 | ⬜ |
| 4 | Admin não desativa a si mesmo (422) | Frente 1 | ⬜ |
| 5 | Seed RBAC carregado em todos os ambientes | Frente 2 | ⬜ |
| 6 | Admin Tenant acessa tudo | Frente 4 | ⬜ |
| 7 | Gerente BU edita apenas sua BU | Frente 4 | ⬜ |
| 8 | Operador BU apenas lê | Frente 4 | ⬜ |
| 9 | Auditor apenas lê auditoria | Frente 4 | ⬜ |
| 10 | Usuário sem BU → não acessa (403) | Frente 3 | ⬜ |
| 11 | Usuário sem módulo → não acessa (403) | Frente 3 | ⬜ |
| 12 | 403 amigável em PT-BR (RFC 7807) | Frente 4 | ⬜ |
| 13 | Acesso direto URL proibida → 403 (não 404) | Frente 4 | ⬜ |
| 14 | Testes automatizados RBAC passando (27 cenários) | Frente 4 | ⬜ |

---

## 🚧 Bloqueios Identificados

| Bloqueio | Ação | Responsável |
|:---|:---|:---|
| (preencher na review) | | |

---

## 📊 Progresso Acumulado

| Frente | Tasks | Concluídas | Pendentes |
|:---|:---:|:---:|:---:|
| Frente 0 — Correções Pré-Sprint | 20 | 20 ✅ | 0 |
| Frente 1 — Gestão de Usuários | 3 | 0 | 3 |
| Frente 2 — Entidades RBAC | 2 | 0 | 2 |
| Frente 3 — PermissionService + API | 2 | 0 | 2 |
| Frente 4 — Integração + Testes | 4 | 0 | 4 |
| Frente 5 — Correções Recomendadas | 15 | 0 | 15 |
| Frente 5b — Infra + DevX | 2 | 0 | 2 |
| **TOTAL** | **48** | **20 (42%)** | **28** |

---

## ➡️ Próximo Passo

**Sprint 5 — Portal do Cliente** (15/09 → 30/09): Login Keycloak, Onboarding guiado 4 passos, Dashboard do cliente, App Switcher.

---

🤖 *Checklist de review da Sprint 4 reestruturado em 17/07/2026. Frente 0 concluída com 20/20 tasks. O teste parametrizado da matriz RN10-01 deve estar 100% verde antes da review com PO.*
