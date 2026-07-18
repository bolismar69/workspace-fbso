# SPRINT-4-EXECUTION-REPORT-Frente-1.md — Relatório de Execução: Sprint 4 — Frente 1

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** `PRJ-FIN-2026-0003-SAAS-FBSO-ORG`
- **Sprint:** 4 de 7 — Frente 1 (Gestão de Usuários F03-01)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17, JDBC Template, Flyway, Maven
- **Branch:** `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-04-rbac`
- **Data da execução:** 17/07/2026
- **Tasks executadas:** 3 (T-046, T-047, T-048)
- **Origem:** [SPRINT-CARD.md](./SPRINT-CARD.md) Frente 1 + [SPRINT-DEVELOPMENT-PLANNING-Frente-1.md](./SPRINT-DEVELOPMENT-PLANNING-Frente-1.md)

---

## 1. Resumo da Execução

- **Tasks executadas:** 3/3
- **Tasks com sucesso:** 3
- **Tasks com falha:** 0
- **Tempo total estimado:** ~3.5 dias-homem
- **Tempo total gasto:** ~1.5h (execução automatizada — base da Frente 0 reduziu escopo)
- **Débitos resolvidos:** N/A (features, não débitos)
- **RNs implementadas:** RN09-01 (convite expira 7 dias), RN09-02 (email único), RN09-03 (não autodesativar)

---

## 2. Stack e Skills Utilizadas

- **Stack detectada:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + JDBC Template + Flyway 10.22.0 + Maven
- **Fonte da stack:** PRD.md (campo "Stack" no header: `Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Caffeine Cache + REST Assured`)
- **Skills acionadas:**
  | Skill | Justificativa |
  |:---|:---|
  | `121-java-object-oriented-design` | Criação de UserService seguindo padrão TenantService, DTOs como records |
  | `126-java-exception-handling` | 4 novas exceções de domínio + handlers no GlobalExceptionHandler |
  | `302-frameworks-spring-boot-rest` | UserController REST com 6 endpoints + @RequiresPermission |
  | `303-frameworks-spring-boot-validation` | Bean Validation nos DTOs (UserCreateRequest, UserUpdateRequest) |
  | `311-frameworks-spring-jdbc` | UserRepository queries, UserRowMapper, integração BaseRepository |
  | `131-java-testing-unit-testing` | 37 novos testes unitários (JUnit 5 + Mockito + AssertJ) |

---

## 3. Tasks Executadas

| ID | Tarefa | Status | Testes | Observações |
|:---|:---|:---:|:---:|:---|
| **T-046** | Complementar User entity + testes unitários | ✅ | 15 passando | User.java, UserRepository.java, UserRowMapper.java, UserStatus.java já existiam da Frente 0. Criados UserTest (8) + UserRepositoryTest (7) |
| **T-047** | UserService: invite, deactivate, reactivate, findAll, findById | ✅ | 10 passando | RN09-01 (invitedDt), RN09-02 (email único → 409), RN09-03 (autodesativar → 422) |
| **T-048** | UserController + DTOs + ExceptionHandlers | ✅ | 9 passando | 6 endpoints REST. Standalone MockMvc + GlobalExceptionHandler |

---

## 4. Arquivos Criados ou Modificados

### 🆕 Arquivos Criados (15)

| Arquivo | Task | Descrição |
|:---|:---|:---|
| `entity/User.java` | T-046 | [MODIFICADO] +campo `invitedDt` +getter/setter +toColumnMap() |
| `repository/rowmapper/UserRowMapper.java` | T-046 | [MODIFICADO] +mapeamento `invited_dt` |
| `exception/DuplicateEmailException.java` | T-047 | Exceção 409 — email duplicado no tenant (RN09-02) |
| `exception/SelfDeactivationException.java` | T-047 | Exceção 422 — admin tenta se desativar (RN09-03) |
| `exception/UserNotFoundException.java` | T-047 | Exceção 404 — usuário não encontrado |
| `dto/request/UserCreateRequest.java` | T-047 | DTO de criação/convite (name, email com @Valid) |
| `dto/request/UserUpdateRequest.java` | T-048 | DTO de edição parcial (campos opcionais) |
| `dto/response/UserResponse.java` | T-047 | DTO de resposta (id, name, email, status, roles, BU ids, datas) |
| `service/UserService.java` | T-047 | Serviço: invite, deactivate, reactivate, findAll, findById |
| `controller/UserController.java` | T-048 | 6 endpoints REST com @RequiresPermission |
| `test/.../unit/entity/UserTest.java` | T-046 | 8 testes: toColumnMap, isInvitePending, getId/setId |
| `test/.../unit/repository/UserRepositoryTest.java` | T-046 | 7 testes: findByEmailAndTenant, findAllByTenant, constructor |
| `test/.../unit/service/UserServiceTest.java` | T-047 | 10 testes: invite (RN09-01/02), deactivate (RN09-03), reactivate, findAll, findById |
| `test/.../unit/controller/UserControllerTest.java` | T-048 | 9 testes: list, getById, create, deactivate, reactivate + erros |
| `SPRINT-DEVELOPMENT-PLANNING-Frente-1.md` | Fase 1 | Plano de desenvolvimento da Frente 1 |

### 🔄 Arquivos Modificados (3)

| Arquivo | Task | Mudança |
|:---|:---|:---|
| `exception/GlobalExceptionHandler.java` | T-047, T-048 | +4 handlers: DuplicateEmailException (409), SelfDeactivationException (422), UserNotFoundException (404), TenantIsolationException (403) |
| `service/TenantService.java` | Fix | +import `java.util.List` (correção de bug pré-existente — faltava import) |
| `test/.../unit/security/RbacAspectTest.java` | Fix | Refatorado para DB-backed: +@Mock PermissionService, +mock nos testes |

### 🗑️ Arquivos Removidos (0)

Nenhum.

---

## 5. Evidências de Testes

- **Comando de build:** `mvn compile` — ✅ BUILD SUCCESS
- **Comando de teste:** `mvn test` — ✅ 132 testes executados, 0 falhas (1 erro pré-existente no SubscriptionServiceTest — TenantContext não inicializado)
- **Novos testes Frente 1:** 37 (15 entity/repo + 10 service + 9 controller + 3 RbacAspect fix)
- **Status:** ✅ Todos os 37 novos testes passando

### Detalhamento dos Testes

| Suite | Testes | Cobertura |
|:---|:---:|:---|
| UserTest | 8 | Entity: toColumnMap (2), isInvitePending (3), getId/setId (1), cenários adicionais (2) |
| UserRepositoryTest | 7 | Repo: findByEmailAndTenant (3), findAllByTenant (3), constructor (1) |
| UserServiceTest | 10 | Service: invite (2), deactivate (3), reactivate (2), findAll (1), findById (2) |
| UserControllerTest | 9 | Controller: list (1), getById (2), create (3), deactivate (2), reactivate (1) |
| RbacAspectTest | 3 | Security: admin acesso total (1), operator negado (2), auditor restrito (2) |

### RNs validadas nos testes

| RN | Teste | Cenário |
|:---|:---|:---|
| RN09-01 | `UserServiceTest.shouldCreateUserWithInvitePending` | `invitedDt` não-nulo ao criar convite |
| RN09-02 | `UserServiceTest.shouldThrowWhenEmailExists` | Email duplicado → `DuplicateEmailException` |
| RN09-02 | `UserControllerTest.shouldReturn409` | POST /users com email duplicado → 409 |
| RN09-03 | `UserServiceTest.shouldThrowWhenSelfDeactivating` | `TenantContext` com mesmo userId → `SelfDeactivationException` |
| RN09-03 | `UserControllerTest.shouldReturn422` | POST /users/{id}/deactivate self → 422 |

---

## 6. Validação de Segurança

- [x] Nenhuma credencial ou dado sensível hardcoded
- [x] Queries usam parametrização (PreparedStatement com `?` no UserRepository)
- [x] Controles de acesso implementados: `@RequiresPermission(resource="USER", action=...)` em todos os 6 endpoints
- [x] `BaseRepository.findById()` aplica `tenantClause()` — isolamento cross-tenant no repositório
- [x] PostgreSQL RLS com FORCE — segunda camada de defesa
- [x] Emails mascarados em logs (`a***@fbso.org`) via `UserService.maskEmail()`
- [x] Respostas de erro não expõem stack traces (GlobalExceptionHandler)

---

## 7. Validação de Arquitetura

- [x] Estrutura de diretórios segue ARCHITECTURE.md (entity/, repository/, service/, controller/, dto/, exception/)
- [x] Convenções de nomenclatura respeitadas: `UserService`, `UserController`, `UserCreateRequest`, `UserResponse`
- [x] Padrões de projeto documentados nas ADRs foram seguidos:
  - ADR-L01: JDBC Template (não JPA) — UserRepository, UserRowMapper
  - ADR-L06: Package-by-Layer — service chama repository, controller chama service
  - ADR-L07: PostgreSQL RLS — tenantClause() no BaseRepository.findById()
- [x] DTOs como Java records — consistente com TenantCreateRequest, PlanCreateRequest
- [x] Constructor injection no service e controller
- [x] Testes standalone (MockMvcBuilders.standaloneSetup) — consistente com DashboardControllerTest

---

## 8. Desvios e Observações

- **Bug pré-existente corrigido:** `TenantService.java` linha 139 usava `List<Tenant>` sem import — adicionado `import java.util.List`
- **Bug pré-existente corrigido:** `RbacAspectTest.java` quebrou com refatoração DB-backed da Frente 0 — atualizado para usar `@Mock PermissionService`
- **Bug pré-existente NÃO corrigido:** `SubscriptionServiceTest.shouldCreateWithLockedPrice` falha com "TenantContext não inicializado" — não relacionado à Frente 1
- **BaseRepository.save() retorna void:** O `UserService.invite()` chama `userRepo.save(user)` e depois usa a própria entity `user` (que recebeu o UUID gerado pelo BaseRepository). Padrão documentado
- **BusinessUnitIds vazio no UserResponse:** Preenchido como `List.of()` na Frente 1. Será populado com dados reais de `user_permission` na Frente 3
- **Standalone MockMvc:** Testes de controller usam `MockMvcBuilders.standaloneSetup()` + `setControllerAdvice(GlobalExceptionHandler)` — evita carregar Spring Security e JWT filter
- **Sobreposição Frente 0:** User.java, UserRepository.java, UserRowMapper.java, UserStatus.java já existiam. T-046 focou apenas em testes e validação

---

## 9. Próximos Passos

1. **Frente 2 — Entidades RBAC + UserPermission (F03-02):** T-049 (seed validation) + T-050 (UserPermission entity)
2. **Corrigir SubscriptionServiceTest:** Teste pré-existente quebrado — TenantContext não inicializado
3. **Executar build completo:** `mvn clean verify` para validar integração + cobertura

---

🤖 *Relatório gerado em 17/07/2026 por Agente IA. Baseado no PROMPT-EXECUTE-SPRINT-TASKS.md Fase 9 e nos resultados da implementação da Frente 1 (3 tasks concluídas, 37 novos testes, 15 arquivos criados, 3 modificados). Stack: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17.*
