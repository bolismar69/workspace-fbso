# SPRINT-DEVELOPMENT-PLANNING-Frente-2.md — Plano de Desenvolvimento: Sprint 4 — Frente 2

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** `PRJ-FIN-2026-0003-SAAS-FBSO-ORG`
- **Sprint:** 4 de 7 — Frente 2 (Entidades RBAC + UserPermission F03-02)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17, JDBC Template (não JPA), Flyway, Maven
- **Branch:** `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-04-rbac`
- **Data:** 17/07/2026
- **Origem:** [SPRINT-CARD.md](./SPRINT-CARD.md) Frente 2 + [TASKS.md](../../TASKS.md) T-049 a T-050

---

## 1. Visão Geral

- **Sprint Goal (contexto):** "4 papéis aplicados com matriz de permissões RN10-01. Gestão de usuários com convite por e-mail. Vinculação Usuário × Unidade × Módulo."
- **Objetivo da Frente 2:** Validar seed data da matriz RN10-01 + criar entidade UserPermission para vínculo Usuário×BU×Role
- **Tasks a implementar:** 2 (T-049, T-050)
- **Ordem de execução:** T-049 → T-050 (T-050 depende indiretamente — User deve existir da Frente 1)
- **Estimativa total:** 2.5 dias-homem
- **Stack detectada:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + JDBC Template + Flyway 10.22.0 + Maven
- **Pré-requisitos:** Frente 0 ✅ (entities base, seed V004, FK V006) + Frente 1 ✅ (User completo)

### Status Inicial (Pós-Frente 0 + Frente 1)

| Artefato | Status | Observação |
|:---|:---:|:---|
| `entity/ResourceAction.java` | ✅ Existente | id, resourceName, action, toColumnMap() — tabela global |
| `entity/RoleResource.java` | ✅ Existente | id, role, resourceActionId, toColumnMap() — tabela global |
| `entity/BusinessUnit.java` | ✅ Existente | Entity mínima — referência estrutural |
| `entity/User.java` | ✅ Existente | +invitedDt (Frente 1) |
| V004 seed migration | ✅ Existente | 28 resource_actions + matriz 4 roles |
| U004 rollback | ✅ Existente | DELETE reverso do seed |
| V006 FK migration | ✅ Existente | FK user_permission → business_unit |
| U006 rollback | ✅ Existente | DROP CONSTRAINT reverso |
| `PermissionService.java` | ✅ Existente | Carrega matriz no startup, checkPermission(), getUserRoles() |
| `entity/UserPermission.java` | ❌ Ausente | A criar — T-050 |
| `repository/PermissionRepository.java` | ❌ Ausente | A criar — T-050 |
| Testes RBAC entity | ❌ Ausentes | A criar — T-049 |

---

## 2. Dependências entre Tasks

```
Frente 0 ✅ (ResourceAction, RoleResource, BusinessUnit, V004, V006, PermissionService)
Frente 1 ✅ (User completo com UserRepository)
    │
    ├── T-049 (Validar seed V004 + testes entity) — ~1d
    │   ⚠️ Entities já existem. Seed já existe. Foco: validação + testes
    │
    └── T-050 (UserPermission entity + repository) — ~1.5d
        Depende de: User (Frente 1) + BusinessUnit (Frente 0)
```

**Sobreposição com Frente 0:**

| Artefato | Criado na Frente 0 | Task original | O que falta |
|:---|:---|:---|:---|
| `ResourceAction.java` | T-102.DT-054 ✅ | T-049 | Nada — completo |
| `RoleResource.java` | T-102.DT-054 ✅ | T-049 | Nada — completo |
| V004 seed migration | T-105.DT-057 ✅ | T-049 | Validar em múltiplos ambientes |
| V006 FK migration | T-104.DT-056 ✅ | T-050 (pré-req) | Nada — FK já existe |

> **Conclusão:** T-049 está ~85% concluída pela Frente 0. Falta apenas validação e testes.

---

## 3. Plano por Task

### T-049 — Validar Seed V004 + Testes ResourceAction/RoleResource

- **Critério DONE:** Seed carrega corretamente em todos os ambientes. Consulta `findByRole` retorna recursos corretos
- **Estimativa:** 1d (reduzido — entities e seed já existem)
- **Status da base (Frente 0):**
  - `ResourceAction.java`: ✅ id, resourceName, action, toColumnMap()
  - `RoleResource.java`: ✅ id, role, resourceActionId, toColumnMap()
  - V004 seed: ✅ 28 resource_actions + matriz 4 roles (ADMIN_TENANT total, MANAGER_BU 6, OPERATOR_BU 2, AUDITOR 1)
  - `PermissionService.loadPermissionMatrix()`: ✅ Carrega via JOIN e popula ConcurrentHashMap
- **Abordagem:**
  1. Criar `ResourceActionTest.java` — validar toColumnMap()
  2. Criar `RoleResourceTest.java` — validar toColumnMap()
  3. Criar `RbacSeedValidationTest.java` — teste de integração que:
     - Verifica contagem: 28 resource_actions, role_resource: ADMIN_TENANT=28, MANAGER_BU=6, OPERATOR_BU=2, AUDITOR=1
     - Valida que PermissionService.loadPermissionMatrix() carrega sem erro
     - Verifica que ADMIN_TENANT tem permissão para DASHBOARD:view
     - Verifica que MANAGER_BU NÃO tem permissão para TENANT:create
     - Verifica que AUDITOR só tem AUDIT:view
  4. Validar consistência matriz SPRINT-CARD × seed V004 (auditoria de permissões)
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/test/.../unit/entity/ResourceActionTest.java` | 🆕 | Testes da entity ResourceAction |
  | `src/test/.../unit/entity/RoleResourceTest.java` | 🆕 | Testes da entity RoleResource |
  | `src/test/.../integration/RbacSeedValidationTest.java` | 🆕 | Validação da seed V004 contra matriz RN10-01 |
- **Dependências:** Frente 0 (entities + seed)
- **Riscos:**
  - Seed pode não carregar em CI se o schema não estiver criado (V001 depende de tabelas)
  - Colisão de ID se V004 rodar mais de uma vez — verificar idempotência
- **Skills aplicáveis:** `131-java-testing-unit-testing`, `313-frameworks-spring-db-migrations-flyway`, `311-frameworks-spring-jdbc`

### T-050 — UserPermission Entity + Repository

- **Critério DONE:** UserPermission com UNIQUE constraint. Query de permissões por usuário retorna BUs + role. Admin tem acesso implícito a todas as BUs
- **Estimativa:** 1.5d
- **Abordagem:**
  1. Criar `UserPermission.java` — entity com:
     - Campos: id, userId (FK→user), businessUnitId (FK→business_unit), role (String do enum Role)
     - UNIQUE(user_id, business_unit_id) — um usuário não pode ter 2 roles na mesma BU
     - `toColumnMap()` mapeando userId, businessUnitId, role
     - Não estende BaseEntity (sem tenant_id — a BU já tem tenant; sem soft delete próprio)
  2. Criar `PermissionRepository.java` (renome para não conflitar com PermissionService):
     - `findByUser(UUID userId)` — lista user_permissions do usuário
     - `assign(UUID userId, UUID businessUnitId, String role)` — INSERT
     - `revoke(UUID userId, UUID businessUnitId)` — DELETE físico (não soft delete)
     - `findByUserAndBu(UUID userId, UUID businessUnitId)` — verificar existência
  3. Criar `UserPermissionRowMapper.java`
  4. Atualizar `PermissionService.java`:
     - `assignRole(userId, buId, role)` — delega ao PermissionRepository + invalida cache
     - `revokeRole(userId, buId)` — delega ao PermissionRepository
     - `getUserRoles()` — já existe mas usa SQL inline; refatorar para usar PermissionRepository
  5. Criar testes:
     - `UserPermissionTest.java` — entity toColumnMap()
     - `PermissionRepositoryTest.java` — métodos do repository
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/.../entity/UserPermission.java` | 🆕 | Entidade de vínculo usuário×BU×role |
  | `src/main/.../repository/PermissionRepository.java` | 🆕 | Repository para user_permission |
  | `src/main/.../repository/rowmapper/UserPermissionRowMapper.java` | 🆕 | RowMapper para UserPermission |
  | `src/test/.../unit/entity/UserPermissionTest.java` | 🆕 | Testes da entity |
  | `src/test/.../unit/repository/PermissionRepositoryTest.java` | 🆕 | Testes do repository |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/.../service/PermissionService.java` | 🔄 | +assignRole(), +revokeRole(), refatorar getUserRoles() |
- **Dependências:** Frente 1 (User) + Frente 0 (BusinessUnit, V006 FK)
- **Riscos:**
  - UNIQUE constraint no banco — verificar se a tabela `user_permission` (V001) já tem a constraint
  - `PermissionService.getUserRoles()` usa SQL inline — manter compatibilidade com fallback JWT para ADMIN_TENANT
  - FK user_permission → user — garantir que UserRepository está funcional (Frente 1)
- **Skills aplicáveis:** `121-java-object-oriented-design`, `311-frameworks-spring-jdbc`, `304-frameworks-spring-boot-security`

---

## 4. Ordem de Execução

1. **T-049** — Criar 3 testes (ResourceActionTest, RoleResourceTest, RbacSeedValidationTest)
   - Checkpoint: `mvn test -Dtest="ResourceActionTest,RoleResourceTest,RbacSeedValidationTest"` — todos verdes
2. **T-050** — Criar UserPermission entity + PermissionRepository + testes
   - Atualizar PermissionService (assignRole, revokeRole)
   - Checkpoint: `mvn test -Dtest="UserPermissionTest,PermissionRepositoryTest"` — todos verdes
3. **Checkpoint final:** `mvn clean verify` — suite completa

---

## 5. Estratégia de Build e Verificação

- **Comando de build:** `mvn compile`
- **Comando de teste rápido:** `mvn test`
- **Comando de teste completo:** `mvn clean verify`

### Checkpoints

| Checkpoint | Após | Comando | Critério |
|:---|:---|:---|:---|
| CP1 | T-049 (testes entity + seed) | `mvn test -Dtest="ResourceActionTest,RoleResourceTest,RbacSeedValidationTest"` | Todos verdes |
| CP2 | T-050 (entity + repo + PermissionService) | `mvn test -Dtest="UserPermissionTest,PermissionRepositoryTest"` | Todos verdes |
| CP3 | Todas as tasks | `mvn clean verify` | BUILD SUCCESS |

---

## 6. Validação da Matriz RN10-01 (Checklist T-049)

| Papel | # resource_actions | Verificação |
|:---|:---:|:---|
| ADMIN_TENANT | 28 | Acesso total — todos os recursos |
| MANAGER_BU | 6 | BUSINESS_UNIT(view,create,edit) + PRODUCT_SERVICE(view,create,edit) |
| OPERATOR_BU | 2 | BUSINESS_UNIT(view) + PRODUCT_SERVICE(view) |
| AUDITOR | 1 | AUDIT(view) |
| **Total** | **37 registros em role_resource** | |

### Verificações Negativas (devem FALHAR)

| Papel | Recurso | Ação | Deve retornar |
|:---|:---|:---|:---|
| MANAGER_BU | TENANT | create | ❌ sem permissão |
| MANAGER_BU | AUDIT | view | ❌ sem permissão |
| OPERATOR_BU | BUSINESS_UNIT | edit | ❌ sem permissão |
| AUDITOR | DASHBOARD | view | ❌ sem permissão |

---

🤖 *Documento gerado em 17/07/2026 por Agente IA. Baseado no PROMPT-EXECUTE-SPRINT-TASKS.md Fase 1 e nos artefatos existentes das Frentes 0-1.*
