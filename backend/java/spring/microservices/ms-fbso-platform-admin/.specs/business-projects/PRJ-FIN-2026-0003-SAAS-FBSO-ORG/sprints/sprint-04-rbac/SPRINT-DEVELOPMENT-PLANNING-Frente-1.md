# SPRINT-DEVELOPMENT-PLANNING-Frente-1.md — Plano de Desenvolvimento: Sprint 4 — Frente 1

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** `PRJ-FIN-2026-0003-SAAS-FBSO-ORG`
- **Sprint:** 4 de 7 — Frente 1 (Gestão de Usuários F03-01)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17, JDBC Template (não JPA), Flyway, Maven
- **Branch:** `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-04-rbac`
- **Data:** 17/07/2026
- **Origem:** [SPRINT-CARD.md](./SPRINT-CARD.md) Frente 1 + [TASKS.md](../../TASKS.md) T-046 a T-048

---

## 1. Visão Geral

- **Sprint Goal (contexto):** "4 papéis (Admin Tenant, Gerente BU, Operador BU, Auditor) aplicados com matriz de permissões RN10-01. Gestão de usuários com convite por e-mail. Vinculação Usuário × Unidade × Módulo. Bloqueio de acesso direto com 403 amigável em PT-BR."
- **Objetivo da Frente 1:** CRUD completo de usuários com convite por e-mail, desativação segura e reativação
- **Tasks a implementar:** 3 (T-046, T-047, T-048)
- **Ordem de execução:** Sequencial (T-046 → T-047 → T-048)
- **Estimativa total:** 3.5 dias-homem
- **Stack detectada:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + JDBC Template + Flyway 10.22.0 + Maven
- **Pré-requisito:** Frente 0 ✅ (concluída 17/07/2026)

### Status Inicial (Pós-Frente 0)

| Artefato | Status | Observação |
|:---|:---:|:---|
| `entity/User.java` | ✅ Existente | 6 campos + auditoria, `toColumnMap()`, `isInvitePending()` |
| `enums/UserStatus.java` | ✅ Existente | ACTIVE, INACTIVE, INVITE_PENDING |
| `repository/UserRepository.java` | ✅ Existente | `findByEmailAndTenant()`, `findAllByTenant()` |
| `repository/rowmapper/UserRowMapper.java` | ✅ Existente | Mapeamento completo (dados + auditoria) |
| `service/UserService.java` | ❌ Ausente | A criar — T-047 |
| `controller/UserController.java` | ❌ Ausente | A criar — T-048 |
| `dto/request/UserCreateRequest.java` | ❌ Ausente | A criar — T-047/T-048 |
| `dto/response/UserResponse.java` | ❌ Ausente | A criar — T-047/T-048 |
| Testes unitários User | ❌ Ausentes | A criar — T-046/T-047 |
| Testes integração User | ❌ Ausentes | A criar — T-048 |

---

## 2. Dependências entre Tasks

```
Frente 0 ✅ (User entity, UserRepository, UserRowMapper, UserStatus)
    │
    ▼
T-046 ⚠️ (Complementar User + testes unitários) — ~1d
    │   ⚠️ Base já existe. Escopo: UserStatus enum já usado,
    │       UserRepository já funcional. Foco: testes unitários
    │
    ▼
T-047 (UserService) — ~1.5d
    │   Depende de T-046 (UserRepository finalizado)
    │
    ▼
T-048 (UserController + DTOs) — ~1d
        Depende de T-047 (UserService)
```

**Sobreposição com Frente 0:**

| Artefato | Criado na Frente 0 | Task original | O que falta |
|:---|:---|:---|:---|
| `User.java` | T-101.DT-053 ✅ | T-046 | Nada — completo. Campos: id, tenantId, externalKeycloakId, email, name, status + auditoria |
| `UserStatus.java` | Frente 0 ✅ | T-046 | Nada — ACTIVE, INACTIVE, INVITE_PENDING |
| `UserRepository.java` | T-101.DT-053 ✅ | T-046 | Nada — `findByEmailAndTenant()`, `findAllByTenant()` |
| `UserRowMapper.java` | T-101.DT-053 ✅ | T-046 | Nada — mapeamento completo |

> **Conclusão:** A task T-046 está ~70% concluída pela Frente 0. O esforço restante foca em testes unitários e validação de que os artefatos existentes atendem aos critérios DONE.

---

## 3. Plano por Task

### T-046 — Complementar User Entity + Testes Unitários

- **Critério DONE:** UserRepository funcional. Email único por tenant ativo. Soft delete respeitado. Índice parcial unique_email_active
- **Estimativa:** 1d (reduzido de 1d original — base já existe)
- **Status da base (Frente 0):**
  - `User.java`: ✅ Completo — 6 campos + auditoria, `toColumnMap()`, getters/setters, `isInvitePending()`
  - `UserStatus.java`: ✅ ACTIVE, INACTIVE, INVITE_PENDING
  - `UserRepository.java`: ✅ `findByEmailAndTenant()`, `findAllByTenant()`, `BaseRepository<User>` com tableName `"user"`
  - `UserRowMapper.java`: ✅ Mapeamento completo de todas as colunas
- **Abordagem:** O foco desta task agora é **validação e testes**:
  1. Verificar que `UserRepository` funciona com soft delete (herdado de `BaseRepository`)
  2. Criar `UserRepositoryTest.java` — testes unitários para `findByEmailAndTenant()` e `findAllByTenant()`
  3. Validar que o índice parcial `unique_email_active` (V002) está documentado e funcional
  4. Criar `UserTest.java` — teste unitário da entity (`toColumnMap()`, `isInvitePending()`)
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/test/.../entity/UserTest.java` | 🆕 | Testes unitários da entity User |
  | `src/test/.../repository/UserRepositoryTest.java` | 🆕 | Testes unitários do UserRepository |
- **Arquivos a modificar:** Nenhum (código de produção já existe)
- **Dependências:** Frente 0 concluída
- **Riscos:** Baixo — código base já existe. Risco: tabela `"user"` com aspas (palavra reservada SQL) — verificar se `BaseRepository` trata corretamente
- **Skills aplicáveis:** `131-java-testing-unit-testing`, `311-frameworks-spring-jdbc`

### T-047 — UserService

- **Critério DONE:** Convite com email válido. Email duplicado no mesmo tenant → 409. Autodesativação → 422. Reativação restaura acesso. Convite expira em 7 dias (RN09-01)
- **Estimativa:** 1.5d
- **Abordagem:**
  1. Criar `UserService.java` seguindo o padrão de `TenantService`:
     - Injeção via constructor (`UserRepository`)
     - Métodos: `invite(CreateUserRequest)`, `deactivate(UUID userId)`, `reactivate(UUID userId)`, `findAll()`, `findById()`
  2. `invite()`:
     - Validar email único por tenant ativo (RN09-02) → `findByEmailAndTenant()` + 409 se duplicado
     - Criar User com status `INVITE_PENDING`
     - Setar `invited_dt = now()` (ou usar `created_dt`)
     - Integrar com `EmailService.sendInvite()` (enviar email de convite)
     - Convite expira em 7 dias (RN09-01) — campo `invited_dt` para cálculo
  3. `deactivate()`:
     - Buscar usuário por ID
     - Validar RN09-03: se `userId == TenantContext.getUserId()` → 422 "Um administrador não pode desativar a si mesmo"
     - `softDelete()` via `BaseRepository`
  4. `reactivate()`:
     - Buscar usuário (incluindo soft-deleted)
     - Setar `deleted_dt = null`, `status = ACTIVE`
     - `update()` via `BaseRepository`
  5. `findAll()` / `findById()`:
     - Delegar ao `UserRepository`
     - Mapear `User → UserResponse` (via método privado ou mapper)
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/.../service/UserService.java` | 🆕 | Serviço de gestão de usuários |
  | `src/main/.../dto/request/UserCreateRequest.java` | 🆕 | DTO para criação/convite de usuário |
  | `src/main/.../dto/response/UserResponse.java` | 🆕 | DTO de resposta com dados do usuário |
  | `src/test/.../service/UserServiceTest.java` | 🆕 | Testes unitários do UserService |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/.../entity/User.java` | 🔄 | Adicionar campo `invitedDt` para RN09-01 (expiração do convite) |
  | `src/main/.../repository/rowmapper/UserRowMapper.java` | 🔄 | Adicionar mapeamento de `invited_dt` |
- **Dependências:** T-046 (UserRepository validado)
- **Riscos:**
  - `EmailService.sendInvite()` é stub — verificar se existe implementação real. Se não, criar envio mock para não bloquear
  - Autodesativação: validação cross-field (userId vs contexto) requer `TenantContext`
  - Convite expirado: precisa de lógica de verificação de data (7 dias a partir de `invited_dt`)
- **Skills aplicáveis:** `302-frameworks-spring-boot-rest`, `126-java-exception-handling`, `311-frameworks-spring-jdbc`

### T-048 — UserController + DTOs

- **Critério DONE:** CRUD usuários funcional. Deactivate bloqueia login. Lista exibe nome, email, role, status, BUs vinculadas. Filtro por status
- **Estimativa:** 1d
- **Abordagem:**
  1. Criar `UserController.java` seguindo o padrão de `TenantController`:
     - `@RestController` + `@RequestMapping("/api/v1/users")`
     - Injeção via constructor (`UserService`)
     - `@RequiresPermission(resource="USER", action=...)` em cada endpoint
  2. Endpoints:
     - `GET /api/v1/users` — lista paginada com filtro por status. `@RequiresPermission(resource="USER", action="view")`
     - `GET /api/v1/users/{id}` — buscar por ID. `@RequiresPermission(resource="USER", action="view")`
     - `POST /api/v1/users` — criar/convite. `@RequiresPermission(resource="USER", action="create")`
     - `PATCH /api/v1/users/{id}` — editar (nome, email). `@RequiresPermission(resource="USER", action="edit")`
     - `POST /api/v1/users/{id}/deactivate` — desativar. `@RequiresPermission(resource="USER", action="delete")`
     - `POST /api/v1/users/{id}/reactivate` — reativar. `@RequiresPermission(resource="USER", action="edit")`
  3. DTOs:
     - `UserCreateRequest`: email (NotEmpty, Email), name (NotBlank)
     - `UserUpdateRequest`: email, name (opcionais)
     - `UserResponse`: id, email, name, status, role, businessUnitIds, invitedDt, createdAt
  4. Tratamento de erros:
     - 409 DuplicateEmailException → `{"title":"Email já cadastrado","detail":"Este email já está em uso no tenant.","status":409}`
     - 422 SelfDeactivationException → `{"title":"Operação não permitida","detail":"Um administrador não pode desativar a si mesmo.","status":422}`
     - 404 UserNotFoundException → padrão RFC 7807
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/.../controller/UserController.java` | 🆕 | Endpoints REST de usuários |
  | `src/main/.../dto/request/UserCreateRequest.java` | 🆕 | DTO de criação/convite |
  | `src/main/.../dto/request/UserUpdateRequest.java` | 🆕 | DTO de edição |
  | `src/main/.../dto/response/UserResponse.java` | 🆕 | DTO de resposta |
  | `src/test/.../controller/UserControllerTest.java` | 🆕 | Testes unitários do controller (MockMvc) |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/.../exception/GlobalExceptionHandler.java` | 🔄 | Adicionar handlers para DuplicateEmailException, SelfDeactivationException, UserNotFoundException |
- **Dependências:** T-047 (UserService)
- **Riscos:**
  - `UserResponse` com lista de BUs vinculadas — precisa de join com `user_permission` table. Se complexo, deixar BUs como array vazio na Frente 1 e preencher na Frente 3
  - Filtro por status: `UserRepository.findAllByTenant()` não tem filtro — precisa adicionar método com filtro
- **Skills aplicáveis:** `302-frameworks-spring-boot-rest`, `126-java-exception-handling`, `303-frameworks-spring-boot-validation`

---

## 4. Ordem de Execução

1. **T-046** — Criar `UserTest.java` + `UserRepositoryTest.java` (validação da base existente)
   - Checkpoint: `mvn test -pl . -Dtest="UserTest,UserRepositoryTest"` — todos verdes
2. **T-047** — Criar `UserService.java` + `UserCreateRequest.java` + `UserResponse.java`
   - Adicionar `invitedDt` ao `User.java` e `UserRowMapper.java`
   - Checkpoint: `mvn test -pl . -Dtest="UserServiceTest"` — todos verdes
3. **T-048** — Criar `UserController.java` + `UserUpdateRequest.java`
   - Adicionar exceções ao `GlobalExceptionHandler`
   - Checkpoint: `mvn test -pl . -Dtest="UserControllerTest"` — todos verdes
4. **Checkpoint final:** `mvn clean verify` — todos os 142 testes existentes + novos passando

---

## 5. Estratégia de Build e Verificação

- **Comando de build:** `mvn compile`
- **Comando de teste rápido:** `mvn test`
- **Comando de teste completo:** `mvn clean verify`
- **Comando de cobertura:** `mvn jacoco:report`

### Checkpoints

| Checkpoint | Após | Comando | Critério |
|:---|:---|:---|:---|
| CP1 | T-046 (testes entity + repo) | `mvn test -Dtest="UserTest,UserRepositoryTest"` | Todos verdes |
| CP2 | T-047 (UserService + DTOs) | `mvn test -Dtest="UserServiceTest"` | Todos verdes |
| CP3 | T-048 (UserController + exc handlers) | `mvn test -Dtest="UserControllerTest"` | Todos verdes |
| CP4 | Todas as tasks | `mvn clean verify` | BUILD SUCCESS + todos os testes passando |

### Estratégia de Rollback

Cada task é independente dentro da Frente 1. Se T-047 falhar, T-046 permanece válida. Se T-048 falhar, T-046 e T-047 permanecem válidas. Reverter apenas a task com falha.

---

## 6. RNs Cobertas pela Frente 1

| RN | Descrição | Onde implementar | Validação |
|:---|:---|:---|:---|
| **RN09-01** | Convite expira em 7 dias | `UserService.invite()` — seta `invited_dt`. Validação no login (futuro) | Teste: criar convite, verificar `invited_dt` |
| **RN09-02** | Email único por tenant ativo | `UserService.invite()` — `findByEmailAndTenant()` → 409 | Teste: mesmo email 2× → 409 |
| **RN09-03** | Admin não desativa a si mesmo | `UserService.deactivate()` — compara `userId` com `TenantContext.getUserId()` | Teste: admin tenta se desativar → 422 |

---

## 7. Novas Exceções de Domínio

| Exceção | HTTP | Gatilho | Mensagem PT-BR |
|:---|:---:|:---|:---|
| `DuplicateEmailException` | 409 | Email já cadastrado no tenant | "Este email já está em uso no tenant." |
| `SelfDeactivationException` | 422 | Admin tenta desativar a si mesmo | "Um administrador não pode desativar a si mesmo." |
| `UserNotFoundException` | 404 | Usuário não encontrado ou soft-deleted | "Usuário não encontrado." |
| `InviteExpiredException` | 422 | Convite com mais de 7 dias | "O convite expirou. Solicite um novo convite ao administrador." |

---

## 8. Padrões de Código (Referência)

### Service (seguir `TenantService` como template)

```java
@Service
public class UserService {
    private static final Logger log = LoggerFactory.getLogger(UserService.class);
    private final UserRepository userRepo;

    public UserService(UserRepository userRepo) {
        this.userRepo = userRepo;
    }
    // métodos públicos com @Transactional, logging, validações RN
}
```

### Controller (seguir `TenantController` como template)

```java
@RestController
@RequestMapping("/api/v1/users")
@Tag(name = "Users", description = "Gestão de usuários — Admin Tenant")
public class UserController {
    private final UserService userService;

    public UserController(UserService userService) {
        this.userService = userService;
    }
    // @RequiresPermission(resource = "USER", action = "...")
}
```

### DTO (seguir `TenantCreateRequest` como template)

```java
public record UserCreateRequest(
    @NotEmpty @Email String email,
    @NotBlank String name
) {}
```

---

🤖 *Documento gerado em 17/07/2026 por Agente IA. Baseado no PROMPT-EXECUTE-SPRINT-TASKS.md Fase 1 e nos artefatos existentes da Frente 0 (User entity, UserRepository, UserRowMapper, UserStatus, PermissionService).*
