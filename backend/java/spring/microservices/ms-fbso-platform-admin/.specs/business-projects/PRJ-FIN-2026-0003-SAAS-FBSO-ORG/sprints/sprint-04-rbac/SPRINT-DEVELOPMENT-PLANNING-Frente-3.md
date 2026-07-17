# SPRINT-DEVELOPMENT-PLANNING-Frente-3.md — Plano de Desenvolvimento: Sprint 4 — Frente 3

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** `PRJ-FIN-2026-0003-SAAS-FBSO-ORG`
- **Sprint:** 4 de 7 — Frente 3 (PermissionService + API F03-02/F03-03)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17, JDBC Template, Flyway, Maven
- **Branch:** `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-04-rbac`
- **Data:** 17/07/2026
- **Origem:** [SPRINT-CARD.md](./SPRINT-CARD.md) Frente 3 + [TASKS.md](../../TASKS.md) T-051 a T-052

---

## 1. Visão Geral

- **Objetivo da Frente 3:** API REST para listar e atualizar permissões de usuários (GET/PUT) + completar PermissionService com operações de consulta e batch update
- **Tasks a implementar:** 2 (T-051, T-052)
- **Ordem de execução:** T-051 → T-052
- **Estimativa total:** 3 dias-homem
- **Pré-requisitos:** Frente 0 ✅ (PermissionService base) + Frente 1 ✅ (User) + Frente 2 ✅ (UserPermission, PermissionRepository)

### Status Inicial (Pós-Frentes 0-2)

| Artefato | Status | Observação |
|:---|:---:|:---|
| `PermissionService.java` | ✅ 233 linhas | loadPermissionMatrix, checkPermission, getUserRoles, validateBusinessUnitAccess, assignRole, revokeRole |
| `PermissionRepository.java` | ✅ | findByUser, assign (upsert), revoke, findRolesByUser |
| `UserPermission.java` | ✅ | entity com toColumnMap |
| `PermissionController.java` | ❌ Ausente | A criar — T-052 |
| `dto/response/PermissionResponse.java` | ❌ Ausente | A criar — T-051 |
| `dto/request/PermissionUpdateRequest.java` | ❌ Ausente | A criar — T-051 |
| Testes PermissionController | ❌ Ausentes | A criar — T-052 |

---

## 2. Dependências entre Tasks

```
Frente 0 ✅ (PermissionService base, ResourceAction, RoleResource, V004 seed)
Frente 1 ✅ (User completo)
Frente 2 ✅ (UserPermission entity, PermissionRepository)
    │
    ├── T-051 (completar PermissionService) — ~2d
    │   ⚠️ assignRole/revokeRole já existem. Foco: listUserPermissions (GET), updateUserPermissions (PUT batch)
    │
    └── T-052 (PermissionController + DTOs) — ~1d
        Depende de T-051
```

## 3. Plano por Task

### T-051 — Completar PermissionService (list + batch update)

- **Critério DONE:** Permissões atribuídas corretamente. Admin vê todas as BUs. Efeito imediato (RN11-03)
- **Estimativa:** 2d (reduzido — base já existe)
- **Status da base:**
  - `assignRole(userId, buId, role)`: ✅ com tenant validation + upsert
  - `revokeRole(userId, buId)`: ✅ com tenant validation
  - `checkPermission(resource, action)`: ✅ DB-backed via matriz carregada
  - `validateBusinessUnitAccess(buId)`: ✅ via JWT business_unit_ids claim
- **O que falta:**
  1. `getUserPermissions(UUID userId)` — retorna lista de UserPermission para o usuário, com validação de tenant
  2. `updateUserPermissions(UUID userId, List<PermissionAssignment> assignments)` — batch update:
     - Deleta todas as permissões atuais do usuário
     - Insere as novas (validadas uma a uma com tenant check)
     - Tudo em uma transação (@Transactional)
  3. `getUserBusinessUnits(UUID userId)` — retorna lista de BUs que o usuário tem acesso (admin: todas; demais: do user_permission)
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `dto/response/PermissionResponse.java` | 🆕 | DTO: userId, businessUnitId, role |
  | `dto/request/PermissionUpdateRequest.java` | 🆕 | DTO: lista de assignments [{businessUnitId, role}] |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `service/PermissionService.java` | 🔄 | +getUserPermissions(), +updateUserPermissions(), +getUserBusinessUnits() |
- **Dependências:** Frente 2 (PermissionRepository, UserPermission)
- **Riscos:**
  - Batch update: DELETE + INSERT em transação — garantir atomicidade com @Transactional
  - Admin vê todas as BUs: precisa de join ou fallback para lista completa (futuro BusinessUnitRepository)
- **Skills aplicáveis:** `121-java-object-oriented-design`, `311-frameworks-spring-jdbc`, `126-java-exception-handling`

### T-052 — PermissionController + DTOs

- **Critério DONE:** GET retorna permissões atuais. PUT atualiza vínculos. Auditoria registrada para cada alteração
- **Estimativa:** 1d
- **Abordagem:**
  1. Seguir padrão `UserController` / `TenantController`:
     - `@RestController` + `@RequestMapping("/api/v1/users/{userId}/permissions")`
     - Constructor injection (`PermissionService`)
     - `@RequiresPermission(resource="PERMISSION", action=...)`
  2. Endpoints:
     - `GET /api/v1/users/{userId}/permissions` — lista permissões do usuário
     - `PUT /api/v1/users/{userId}/permissions` — atualiza vínculos (batch)
  3. Tratamento de erros:
     - 404 UserNotFoundException se usuário não pertencer ao tenant
     - 403 PermissionDeniedException via @RequiresPermission
  4. Testes: PermissionControllerTest com standalone MockMvc + GlobalExceptionHandler
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `controller/PermissionController.java` | 🆕 | GET + PUT /api/v1/users/{userId}/permissions |
  | `test/.../unit/controller/PermissionControllerTest.java` | 🆕 | Testes do controller (MockMvc standalone) |
- **Dependências:** T-051 (PermissionService.getUserPermissions/updateUserPermissions)
- **Riscos:**
  - URL aninhada: `/users/{userId}/permissions` — userId precisa ser validado contra tenant
  - Batch update sem rollback parcial — @Transactional garante atomicidade
- **Skills aplicáveis:** `302-frameworks-spring-boot-rest`, `303-frameworks-spring-boot-validation`

---

## 4. Ordem de Execução

1. **T-051** — Adicionar getUserPermissions(), updateUserPermissions() ao PermissionService
   - Criar PermissionResponse + PermissionUpdateRequest DTOs
   - Checkpoint: `./mvnw test -Dtest="PermissionServiceTest"` (criar se não existir)
2. **T-052** — Criar PermissionController + PermissionControllerTest
   - Checkpoint: `./mvnw test -Dtest="PermissionControllerTest"` — todos verdes
3. **Checkpoint final:** `./mvnw test` — suite completa

---

## 5. Estratégia de Build

- **Build:** `./mvnw compile`
- **Teste:** `./mvnw test`
- **Verificação completa:** `./mvnw verify`

---

## 6. Estrutura dos DTOs

### PermissionResponse
```java
public record PermissionResponse(
    UUID userId,
    UUID businessUnitId,
    String role
) {
    public static PermissionResponse from(UserPermission up) { ... }
}
```

### PermissionUpdateRequest
```java
public record PermissionUpdateRequest(
    @NotNull UUID userId,
    @NotEmpty List<PermissionAssignment> permissions
) {}

public record PermissionAssignment(
    @NotNull UUID businessUnitId,
    @NotBlank String role
) {}
```

---

🤖 *Documento gerado em 17/07/2026. Baseado no PROMPT-EXECUTE-SPRINT-TASKS.md Fase 1. Frente 3 com alta sobreposição da Frente 0 — PermissionService já tem 233 linhas implementadas.*
