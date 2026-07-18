# SPRINT-DEVELOPMENT-PLANNING-Frente-4.md — Plano de Desenvolvimento: Sprint 4 — Frente 4

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** `PRJ-FIN-2026-0003-SAAS-FBSO-ORG`
- **Sprint:** 4 de 7 — Frente 4 (Integração RBAC + Segurança + Testes F03-04)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17, JDBC Template, Flyway 10.22.0, Caffeine Cache, REST Assured 5.5.7, Maven
- **Branch:** `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-04-rbac`
- **Data do plano:** 17/07/2026
- **Dependências:** Frentes 0, 1, 2 e 3 concluídas ✅
- **Tasks:** 4 (T-053, T-054, T-055, T-056)

---

## 1. Visão Geral

- **Sprint Goal (Frente 4):** Integração final do RBAC — RbacAspect 100% DB-backed (sem fallback JWT), resposta 403 amigável RFC 7807 em PT-BR, bateria completa de testes unitários e de segurança
- **Feature:** F03-04 — Acesso Condicional (403 Amigável)
- **RNs cobertas:** RN10-01 (matriz 4 papéis), RN12-01 (dupla camada UX + segurança), RN12-02 (mensagem amigável)
- **Tasks a implementar:** 4
- **Ordem de execução:** T-053 → T-054 ∥ T-055 → T-056
- **Estimativa total:** ~5.5 dias-homem
- **Stack detectada:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 (inferido de PRD.md §1 + pom.xml)

---

## 2. Dependências entre Tasks

```
Frentes 0-3 (concluídas)
  │
  ├── T-053 (1.5d) — Integrar RbacAspect com RoleResource do banco
  │     │              Remover fallback JWT. Cache em memória sem TTL
  │     │
  │     ├── T-054 (0.5d) — 403 padrão RFC 7807 PT-BR
  │     │     (paralelo com T-055)
  │     │
  │     ├── T-055 (1.5d) — Testes unitários M4
  │     │     (paralelo com T-054)
  │     │
  │     └── T-056 (2d) — Testes segurança RBAC parametrizados
  │           (sequencial após T-053, T-054, T-055)
  │
  └── Checkpoint final: mvn clean verify
```

**Paralelismo:** T-054 (0.5d) e T-055 (1.5d) podem ser executados em paralelo — T-054 é mudança pontual no handler de exceção, T-055 é criação de testes unitários. Ambos dependem apenas de T-053 (RbacAspect integrado).

**Cadeia crítica:** T-053 (1.5d) → T-056 (2d) = 3.5d

---

## 3. Plano por Task

### T-053 — Integrar RbacAspect com RoleResource do banco

- **Critério DONE:** RbacAspect consulta matriz do banco (tabelas `role_resource` + `resource_action`) via `PermissionService`. Cache em memória sem TTL. Sem fallback para JWT. Testes de integração validam que matriz carregada do banco é usada.
- **Estimativa:** 1.5d
- **Status atual (pós-Frente 0):**
  - ✅ `RbacAspect.java` (60 linhas) — já refatorado para injetar `PermissionService` e delegar `checkPermission(resource, action)` — sem Sets hardcoded, usa `Role` enum
  - ✅ `PermissionService.java` (303 linhas) — matriz carregada em memória via `@PostConstruct loadPermissionMatrix()`, `checkPermission()` implementado com suporte a ADMIN_TENANT implícito, `validateBusinessUnitAccess()` funcional
  - ✅ Migration V004 — seed data com matriz RN10-01 (4 roles × 8 resources) carregada
  - ⚠️ `RbacAspect` atual NÃO tem fallback JWT (já foi removido na Frente 0?) — verificar se há código residual de fallback
- **Abordagem:**
  1. Auditar `RbacAspect.java` — confirmar que NÃO há fallback para claims JWT (o `checkPermission` do `PermissionService` já resolve tudo via banco)
  2. Auditar `PermissionService.checkPermission()` — confirmar que a matriz em memória é a ÚNICA fonte de verdade
  3. Criar testes de integração `RbacAspectIntegrationTest.java` — provar que permissões vêm do banco, não do JWT:
     - Modificar seed data → permissão negada
     - Remover seed → ADMIN_TENANT ainda acessa tudo (implícito)
  4. Validar que `PermissionService.loadPermissionMatrix()` carrega em todos os profiles (dev, staging, CI)
  5. Documentar a estratégia de cache (sem TTL, recarregável via endpoint admin futuro)
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/test/.../integration/security/RbacAspectIntegrationTest.java` | 🆕 | Testes de integração provando RBAC DB-backed |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `security/aspect/RbacAspect.java` | 🔄 | Remover qualquer fallback JWT residual, logging aprimorado |
  | `service/PermissionService.java` | 🔄 | Documentar invariantes do cache, null-safety no loadMatrix |
- **Dependências:** Frentes 0-3 (RbacAspect refatorado, PermissionService completo, seed V004)
- **Riscos:**
  - Seed V004 não carregar em ambiente CI → adicionar `@Sql` nos testes ou verificar Flyway config
  - Cache sem TTL pode divergir se seed for alterado manualmente → documentar, criar endpoint de reload futuro
- **Skills aplicáveis:**
  - `304-frameworks-spring-boot-security` — integração RBAC com Spring Security
  - `132-java-testing-integration-testing` — testes de integração com @SpringBootTest
  - `124-java-secure-coding` — validação de que o fallback JWT foi removido

---

### T-054 — Garantir 403 padrão RFC 7807 PT-BR

- **Critério DONE:** Toda resposta 403 segue EXATAMENTE o formato: `{"title":"Acesso negado","detail":"Você não tem permissão para acessar esta área.","status":403}`. Mensagem em PT-BR. Sem stack trace. Sem detalhes técnicos. RFC 7807 com campo `type`.
- **Estimativa:** 0.5d
- **Status atual (pós-Frente 0):**
  - ✅ `GlobalExceptionHandler.java` (205 linhas) — já trata `PermissionDeniedException` → 403 com `ErrorResponse.of("https://api.fbso.org/errors/access-denied", "Acesso negado", 403, "Você não tem permissão para acessar esta área.")`
  - ✅ `PermissionDeniedException.java` (20 linhas) — existe com mensagem em PT-BR
  - ✅ `@ExceptionHandler(AccessDeniedException.class)` mapeado → 403 mesmo formato
  - ✅ `TenantIsolationException` → 403 com a mensagem correta
  - ⚠️ Verificar se o JSON gerado por `ErrorResponse.of()` contém o campo `type` (RFC 7807 obrigatório)
- **Abordagem:**
  1. Auditar `ErrorResponse.java` — confirmar que serializa `type`, `title`, `status`, `detail`
  2. Testar endpoint protegido sem permissão → verificar corpo da resposta 403
  3. Se faltar campo `type` na serialização → adicionar
  4. Garantir que `SecurityException` (401) NÃO vaza stack trace
  5. Escrever teste unitário `GlobalExceptionHandlerTest.java` para o handler 403
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/test/.../unit/exception/ErrorResponse403Test.java` | 🆕 | Teste de serialização JSON do ErrorResponse 403 |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `dto/response/ErrorResponse.java` | 🔄 | Garantir campo `type` na serialização (se ausente) |
  | `exception/GlobalExceptionHandler.java` | 🔄 | Verificar consistência entre todos os handlers 403 |
  | `exception/PermissionDeniedException.java` | 🔄 | Revisar mensagem padrão |
- **Dependências:** T-053 (RbacAspect integrado para testar o fluxo completo JWT → Aspect → 403)
- **Riscos:** Nenhum significativo — handlers já implementados, é refinamento
- **Skills aplicáveis:**
  - `303-frameworks-spring-boot-validation` — validação do formato RFC 7807
  - `302-frameworks-spring-boot-rest` — ResponseEntity e serialização JSON

---

### T-055 — Testes unitários M4: UserService, PermissionService

- **Critério DONE:** ≥ 80% cobertura em `UserService` e `PermissionService`. Cada RN testada com caso positivo + negativo:
  - RN09-03: Admin não pode desativar a si mesmo
  - RN10-01: Matriz 4 papéis — permissões corretas por role
  - RN11-01: Usuário requer ≥1 BU para acesso
  - RN11-02: Usuário requer ≥1 Módulo para acesso
- **Estimativa:** 1.5d
- **Status atual:**
  - ✅ `UserService` — implementado na Frente 1 (métodos: invite, deactivate, reactivate, list)
  - ✅ `PermissionService` — implementado nas Frentes 0+3 (métodos: checkPermission, assignRole, revokeRole, getUserPermissions, updateUserPermissions, validateBusinessUnitAccess, getUserBusinessUnits)
  - ❌ Testes unitários M4 inexistentes (criados como parte desta task)
- **Abordagem:**
  1. Criar `UserServiceTest.java` — MockMvc standalone ou Mockito puro:
     - Convidar usuário → sucesso
     - Email duplicado → DuplicateEmailException
     - Admin desativa a si mesmo → SelfDeactivationException (RN09-03)
     - Reativar usuário → sucesso
     - Listar com filtro por status → sucesso
  2. Criar `PermissionServiceTest.java` — Mockito puro:
     - ADMIN_TENANT → acesso implícito a qualquer resource+action (RN10-01)
     - MANAGER_BU → acesso a BUSINESS_UNIT, PRODUCT_SERVICE (edit)
     - OPERATOR_BU → acesso a BUSINESS_UNIT, PRODUCT_SERVICE (apenas view)
     - AUDITOR → acesso apenas a AUDIT (view)
     - Usuário sem BU vinculada → PermissionDeniedException (RN11-01)
     - Usuário sem módulo → PermissionDeniedException (RN11-02)
     - checkPermission com resource inexistente → negado
  3. Se cobertura < 80% → adicionar cenários de borda até atingir meta
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/test/.../unit/service/UserServiceTest.java` | 🆕 | Testes unitários UserService (RN09-03 + CRUD) |
  | `src/test/.../unit/service/PermissionServiceTest.java` | 🆕 | Testes unitários PermissionService (RN10-01, RN11-01, RN11-02) |
- **Arquivos a modificar:** Nenhum (apenas criação de testes)
- **Dependências:** T-053 (PermissionService com matriz DB-backed carregada)
- **Riscos:**
  - `PermissionService` depende de `TenantContext` (ThreadLocal) → mockar ou usar `@SpringBootTest`
  - Cobertura pode exigir mais cenários do que o estimado → priorizar RNs, deixar edge cases para Frente 5
- **Skills aplicáveis:**
  - `131-java-testing-unit-testing` — JUnit 5 + Mockito
  - `130-java-testing-strategies` — estratégia de cobertura de RNs

---

### T-056 — Testes segurança RBAC: cada papel × endpoint proibido → 403

- **Critério DONE:** Teste parametrizado com 20+ combinações papel×endpoint. Cada combinação valida que o endpoint proibido retorna 403 com o formato padrão. REST Assured + Testcontainers com PostgreSQL real. Matriz RN10-01 100% validada.
- **Estimativa:** 2d
- **Status atual:**
  - ✅ REST Assured 5.5.7 disponível (T-097)
  - ✅ Testcontainers configurado no projeto
  - ✅ Seed V004 carrega matriz RN10-01 completa
  - ✅ Migration V006 FK user_permission → business_unit
  - ✅ FORCE ROW LEVEL SECURITY nas tabelas
  - ✅ 27 endpoints REST com @RequiresPermission implementados
- **Abordagem:**
  1. Criar classe base `BaseRbacSecurityTest.java` com:
     - Container PostgreSQL + Flyway migrate
     - Método helper `givenUserWithRole(Role role, UUID businessUnitId)` que configura TenantContext + SecurityContext
     - REST Assured `RequestSpecification` com mock JWT
  2. Criar `RbacMatrixSecurityTest.java` com teste parametrizado:
     - Fonte de dados: `@CsvSource` ou `@MethodSource` com 20+ combinações
     - Cada linha: papel, método HTTP, endpoint, esperado 403
  3. Combinações a testar (mínimo 20):
     ```
     OPERATOR  × POST   /api/v1/tenants              → 403
     OPERATOR  × POST   /api/v1/plans                → 403
     OPERATOR  × POST   /api/v1/users                → 403
     OPERATOR  × PATCH  /api/v1/products             → 403
     OPERATOR  × POST   /api/v1/business-units       → 403
     AUDITOR   × POST   /api/v1/tenants              → 403
     AUDITOR   × PATCH  /api/v1/tenants              → 403
     AUDITOR   × POST   /api/v1/plans                → 403
     AUDITOR   × POST   /api/v1/users                → 403
     AUDITOR   × PATCH  /api/v1/products             → 403
     MANAGER   × POST   /api/v1/tenants              → 403
     MANAGER   × POST   /api/v1/plans                → 403
     MANAGER   × POST   /api/v1/users                → 403
     MANAGER   × GET    /api/v1/audit                → 403
     OPERATOR  × POST   /api/v1/subscriptions        → 403
     AUDITOR   × POST   /api/v1/subscriptions        → 403
     MANAGER   × POST   /api/v1/subscriptions        → 403
     OPERATOR  × PATCH  /api/v1/plans/{id}           → 403
     AUDITOR   × POST   /api/v1/business-units       → 403
     MANAGER   × POST   /api/v1/products/{id}/deactivate → 403
     ```
  4. Validar que cada resposta 403 contém:
     - `type`: "https://api.fbso.org/errors/access-denied"
     - `title`: "Acesso negado"
     - `status`: 403
     - `detail`: "Você não tem permissão para acessar esta área."
  5. Adicional: testar que ADMIN_TENANT acessa TODOS os endpoints (confirmar acesso total implícito)
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/test/.../security/BaseRbacSecurityTest.java` | 🆕 | Classe base com container PostgreSQL + helpers |
  | `src/test/.../security/RbacMatrixSecurityTest.java` | 🆕 | Teste parametrizado 20+ combinações |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/test/resources/application-test.yml` | 🔄 | Configurar REST Assured port, desabilitar RLS nos testes (mock context) |
- **Dependências:** T-053, T-054, T-055 (RbacAspect integrado, 403 handler refinado, testes unitários passando)
- **Riscos:**
  - Configuração do contexto de segurança nos testes de integração é complexa → criar `Test SecurityConfig` simplificado
  - REST Assured + Testcontainers podem ser lentos → usar `@TestInstance(Lifecycle.PER_CLASS)` para reutilizar container
  - Mock JWT + Keycloak → pode ser necessário um `JwtDecoder` mockado
- **Skills aplicáveis:**
  - `132-java-testing-integration-testing` — Testcontainers + REST Assured
  - `304-frameworks-spring-boot-security` — configuração de segurança nos testes
  - `133-java-testing-acceptance-tests` — testes parametrizados como contratos de segurança

---

## 4. Ordem de Execução

1. **T-053 — Integrar RbacAspect com RoleResource do banco (1.5d)**
   - Primeiro porque todas as outras tasks dependem do RbacAspect 100% DB-backed
   - Auditar código existente (Frente 0), remover fallback JWT residual
   - Criar `RbacAspectIntegrationTest.java`
   - Justificativa: Bloqueia T-054 e T-056

2. **T-054 — 403 padrão RFC 7807 PT-BR (0.5d) ∥ T-055 — Testes unitários M4 (1.5d)**
   - Paralelizáveis: T-054 mexe no handler (1 arquivo + teste), T-055 cria testes unitários (2 novas classes)
   - T-054 é quick win (0.5d) — pode ser concluído enquanto T-055 desenvolve cenários mais longos
   - Ambos dependem de T-053 concluído

3. **T-056 — Testes segurança RBAC parametrizados (2d)**
   - Sequencial após T-053, T-054, T-055
   - Depende do RbacAspect integrado, 403 handler validado e testes unitários verdes
   - É a tarefa mais complexa — requer contexto de teste integrado com PostgreSQL real

**Checkpoint intermediário (após T-053):** `mvn test` deve passar com testes de integração do RbacAspect
**Checkpoint final (após T-056):** `mvn clean verify` — todos os 27 cenários da suite verdes, matriz RN10-01 100% validada

---

## 5. Estratégia de Build e Verificação

- **Comando de build:** `./mvnw compile`
- **Comando de teste rápido (unit):** `./mvnw test -pl -Dtest="*Test" -DfailIfNoTests=false`
- **Comando de teste completo:** `./mvnw verify` (unit + integration + coverage)
- **Comando de cobertura:** `./mvnw jacoco:report` → target/site/jacoco/index.html
- **Meta de cobertura:** ≥ 80% (linhas) nos serviços M4 (`UserService`, `PermissionService`)

### Checkpoints:

| Checkpoint | Após | Comando | Esperado |
|:---|:---|:---|:---|
| CP-1 | T-053 | `./mvnw test` | RbacAspectIntegrationTest verde, sem fallback JWT |
| CP-2 | T-054 + T-055 | `./mvnw test` | 403 handler testado, UserService + PermissionService ≥ 80% cobertura |
| CP-3 | T-056 | `./mvnw verify` | Testes parametrizados 20+ combinações verde, matriz RN10-01 validada |

---

## 6. Ações Manuais ou Externas

> ⚠️ Nenhuma ação manual é necessária para esta frente. Todos os testes rodam com PostgreSQL via Testcontainers (automático). Não há dependência de serviços externos.

### Verificação Manual Opcional (Recomendada)

- **Cenário:** Validar visualmente o JSON 403
- **Quem executa:** Desenvolvedor
- **Pré-condições:** Aplicação rodando localmente com `./mvnw spring-boot:run`
- **Passo a passo:**
  1. Obter token JWT de um usuário OPERATOR
  2. Executar: `curl -X POST http://localhost:8080/api/v1/tenants -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{"name":"test"}'`
  3. Verificar resposta:
     ```json
     {"type":"https://api.fbso.org/errors/access-denied","title":"Acesso negado","status":403,"detail":"Você não tem permissão para acessar esta área."}
     ```

---

## 7. Notas de Sobreposição com Frentes Anteriores

> As Frentes 0-3 construíram a base que reduz o escopo da Frente 4:

| Task Frente 4 | Artefato pré-existente | Origem | O que falta |
|:---|:---|:---|:---|
| **T-053** (RbacAspect integração) | `RbacAspect.java` refatorado DB-backed, `PermissionService` com checkPermission | Frente 0 (T-100, T-115) | Remover fallback JWT residual, testes de integração |
| **T-054** (403 padrão) | `GlobalExceptionHandler` com handlers 403, `PermissionDeniedException` | Setup (T-015) + Frente 0 (T-123) | Validar serialização RFC 7807, teste de formato |
| **T-055** (Testes unitários) | `UserService`, `PermissionService` implementados | Frentes 0, 1, 3 | Criar classes de teste com cobertura ≥ 80% |
| **T-056** (Segurança parametrizada) | REST Assured configurado, seed V004, Testcontainers | Frente 0 (T-097, T-105) | Criar teste parametrizado com 20+ combinações |

---

## Rodapé

🤖 *Plano de desenvolvimento gerado em 17/07/2026 pelo Agente Claude como parte da execução do PROMPT-EXECUTE-SPRINT-TASKS. Stack detectada: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17. Skills aplicáveis: 304-spring-security, 132-integration-testing, 124-secure-coding, 303-validation, 302-rest, 131-unit-testing, 130-testing-strategies, 133-acceptance-tests.*
