# SPRINT-DEVELOPMENT-PLANNING.md — Plano de Desenvolvimento: Sprint 2

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** `PRJ-FIN-2026-0003-SAAS-FBSO-ORG`
- **Sprint:** 2 de 7 — Segurança Cross-Cutting
- **Stack:** Java 25 + Spring Boot 3.5.1 + PostgreSQL + Keycloak (JWT RS256)
- **Data:** 14 de Julho de 2026

---


## 1. Visão Geral

- **Sprint Goal:** JWT autentica via Keycloak RS256. TenantContext isola requisições por tenant_id. @RequiresPermission bloqueia acessos não autorizados (403). @Auditable grava auditoria de forma assíncrona. GlobalExceptionHandler retorna erros RFC 7807 em PT-BR.
- **Tasks a implementar:** 8 (T-009 a T-015 + T-015.1)
- **Ordem de execução:** Sequencial com paralelismo limitado — forte encadeamento de dependências
- **Stack detectada:** Java 25 + Spring Boot 3.5.1 + PostgreSQL 17 + Keycloak OIDC + Maven 3.9

---

## 2. Dependências entre Tasks

```
T-009 (SecurityConfig)
  │
  └── T-010 (JwtAuthenticationFilter)
        │
        ├── T-011 (TenantContext — complementar o da Sprint 1)
        │
        ├── T-015.1 (PostgreSQL RLS — Migration V003 + config no JwtFilter)
        │
        ├── T-013 (@RequiresPermission + RbacAspect)
        │     │
        │     └── T-012 (TenantIsolationAspect — redundante, substituído por T-015.1)
        │
        ├── T-014 (@Auditable + AuditAspect — relativamente independente)
        │
        └── T-015 (GlobalExceptionHandler — independente)
```

**Ordem recomendada:** T-009 → T-010 → T-011 → T-015.1 → T-013 → T-012 → T-014 → T-015

---

## 3. Plano por Task

### T-009 — SecurityConfig.java
- **Critério DONE:** Sem token → 401. Token válido → autenticado. Token inválido → 401. CORS permite origem do frontend
- **Estimativa:** 2d
- **Abordagem:** Classe `@Configuration` + `@EnableWebSecurity`. Configurar `SecurityFilterChain` com:
  - `oauth2ResourceServer` para JWT (Keycloak RS256 via JWKS URI)
  - `sessionCreationPolicy(STATELESS)` — API REST sem sessão
  - `cors()` configurado para aceitar origem do frontend
  - `csrf().disable()` — JWT substitui CSRF
  - Permitir `/actuator/health` sem autenticação
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../config/SecurityConfig.java` | 🆕 | Configuração Spring Security + JWT + CORS |
- **Dependências:** T-001 (pom.xml com spring-boot-starter-security, oauth2-resource-server)
- **Riscos:** JWKS URI do Keycloak indisponível → cache de JWKS. Versão Spring Security 3.5.1 com API estável
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `security-review`

### T-010 — JwtAuthenticationFilter.java
- **Critério DONE:** Filter em toda req. exceto `/actuator/health`. Claims extraídas. SecurityContext populado
- **Estimativa:** 2d
- **Abordagem:** `OncePerRequestFilter` que:
  - Extrai token do header `Authorization: Bearer <jwt>`
  - Valida assinatura RS256 com chave pública do Keycloak (JWKS)
  - Valida expiração (`exp`)
  - Extrai claims: `tenant_id`, `user_id`, `roles`, `business_unit_ids`, `modules`
  - Popula `TenantContext` e `SecurityContextHolder`
  - No `finally`: `TenantContext.clear()`
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../security/JwtAuthenticationFilter.java` | 🆕 | OncePerRequestFilter com validação JWT RS256 |
  | `src/main/java/.../utils/JwtUtils.java` | 🆕 | Utilitário para extração de claims do JWT |
- **Dependências:** T-009
- **Riscos:** JWKS endpoint inacessível → cache com TTL. Claims ausentes no token → 401 amigável
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `security-review`

### T-011 — TenantContext.java (complemento)
- **Critério DONE:** TenantContext.getTenantId() retorna tenant_id. Contexto limpo após requisição
- **Estimativa:** 0.5d
- **Abordagem:** A classe já existe da Sprint 1 com estrutura básica. Esta task é de **complemento**:
  - Verificar se já atende aos critérios DONE
  - Adicionar validação de limpeza no `clear()`
  - Garantir integração com `JwtAuthenticationFilter`
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../security/TenantContext.java` | 🔄 | Revisar e complementar — já criado na Sprint 1 |
- **Dependências:** T-010
- **Riscos:** Vazamento de ThreadLocal entre requisições → teste de concorrência
- **Skills aplicáveis:** `126-java-exception-handling`, `121-java-object-oriented-design`

### T-015.1 — PostgreSQL RLS (Migration V003 + JwtFilter config)
- **Critério DONE:** RLS ativo em 5 tabelas. Politicas criadas. JwtFilter seta `app.current_tenant_id`. INSERT cross-tenant → rejeitado
- **Estimativa:** 1.5d
- **Abordagem:** Defesa em profundidade — camada 1 de isolamento multi-tenant no nível do banco:
  1. Criar `V003__enable_rls.sql` na pasta `db/migration/` com:
     - `ALTER TABLE fbso_platform.<tabela> ENABLE ROW LEVEL SECURITY` para 5 tabelas
     - `CREATE POLICY tenant_isolation ON fbso_platform.<tabela> FOR ALL USING (tenant_id = current_setting('app.current_tenant_id')::UUID) WITH CHECK (tenant_id = current_setting('app.current_tenant_id')::UUID)` para cada tabela
  2. Criar `V003__enable_rls_undo.sql` para rollback
  3. No `JwtAuthenticationFilter`, após validar JWT e setar TenantContext:
     ```java
     jdbcTemplate.update("SET LOCAL app.current_tenant_id = ?", tenantId);
     ```
  4. Para Admin FBSO (sem tenant_id): `SET app.current_tenant_id = ''` ou omitir (RLS permite acesso global)
  5. Adicionar `finally { jdbcTemplate.execute("RESET app.current_tenant_id"); }` no filter
  6. Escrever teste de integração: `TC-INFRA-025` — tentar INSERT com tenant_id forjado
- **Arquivos a criar/modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/resources/db/migration/V003__enable_rls.sql` | 🆕 | RLS + políticas em 5 tabelas |
  | `src/main/resources/db/migration/U003__disable_rls.sql` | 🆕 | Rollback — desabilita RLS |
  | `src/main/java/.../security/JwtAuthenticationFilter.java` | 🔄 | Adicionar `SET app.current_tenant_id` + cleanup |
  | `src/test/java/.../integration/security/RLSIsolationTest.java` | 🆕 | Teste de violação RLS |
- **Dependências:** T-010 (JwtAuthenticationFilter), T-004 (Migration V001)
- **Riscos:** Admin FBSO (cross-tenant) precisa acessar todas as tabelas sem tenant_id → política especial ou bypass. Performance: testar com 1000 tenants que RLS não adiciona latência significativa
- **Skills aplicáveis:** `313-frameworks-spring-db-migrations-flyway`, `security-review`

### T-013 — @RequiresPermission + RbacAspect.java
- **Critério DONE:** Anotação bloqueia sem permissão. Role válida × resource × action. 403 com JSON amigável
- **Estimativa:** 2d
- **Abordagem:**
  - Criar anotação customizada `@RequiresPermission(resource, action)` com `@Retention(RUNTIME)` e `@Target(METHOD)`
  - Criar `RbacAspect` com `@Aspect` + `@Component`:
    - `@Around("@annotation(requiresPermission)")`
    - Lê roles do `TenantContext`
    - Consulta `RoleResource` no banco (ou cache) para verificar se role tem resource+action
    - Se negado → lança `PermissionDeniedException` → 403
    - Cache Caffeine com TTL 5min para matriz de permissões
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../security/annotation/RequiresPermission.java` | 🆕 | Anotação customizada (resource, action) |
  | `src/main/java/.../security/aspect/RbacAspect.java` | 🆕 | Aspecto AOP que verifica permissões |
  | `src/main/java/.../exception/PermissionDeniedException.java` | 🆕 | Exceção para 403 |
- **Dependências:** T-010 (TenantContext), T-042 (seed data RBAC — Sprint 4)
- **Riscos:** Matriz de permissões não carregada (seed data é Sprint 4) → mock/hardcoded para testes enquanto isso
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `121-java-object-oriented-design`

### T-012 — TenantIsolationAspect.java
- **Critério DONE:** Sem tenant_id → SecurityException. Query injetada. Tenants diferentes → dados diferentes
- **Estimativa:** 1.5d
- **Abordagem:**
  - Aspecto AOP que intercepta `@Repository` automaticamente
  - `@Around("@within(org.springframework.stereotype.Repository)")`
  - Injeta `tenant_id` no contexto da query antes da execução
  - Valida que `TenantContext.getTenantId()` não é nulo
  - `@Order(1)` — executa antes do método do repository
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../security/aspect/TenantIsolationAspect.java` | 🆕 | Aspecto AOP para isolamento multi-tenant |
- **Dependências:** T-011 (TenantContext)
- **Riscos:** Native queries sem tenant_id → auditoria de queries. Performance do aspecto → overhead mínimo (pointcut simples)
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `security-review`

### T-014 — @Auditable + AuditAspect.java
- **Critério DONE:** Registro gerado por operação. Valores capturados. Async não bloqueia thread principal
- **Estimativa:** 1.5d
- **Abordagem:**
  - Criar anotação `@Auditable(entityType, action)`
  - Criar `AuditAspect` com `@Aspect` + `@Component`:
    - `@AfterReturning` em métodos anotados
    - Captura valores antes/depois via reflection/snapshot
    - Grava em `audit_log` via `JdbcTemplate` de forma ASSÍNCRONA (`@Async`)
    - Configurar `ThreadPoolTaskExecutor` para async
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../security/annotation/Auditable.java` | 🆕 | Anotação customizada (entityType, action) |
  | `src/main/java/.../security/aspect/AuditAspect.java` | 🆕 | Aspecto AOP com @Async para auditoria |
- **Dependências:** T-004 (tabela audit_log já existe)
- **Riscos:** Perda de registros em crash (ADR-L03 — aceitável para Fase 0). Fila cheia → log warning
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `126-java-exception-handling`

### T-015 — GlobalExceptionHandler.java
- **Critério DONE:** type, title, status, detail. Sem stack trace. Mensagens PT-BR
- **Estimativa:** 1d
- **Abordagem:**
  - `@ControllerAdvice` com handlers para cada tipo de exceção:
    - `BusinessException` → 422 (Unprocessable Entity)
    - `PermissionDeniedException` → 403 (Forbidden)
    - `MethodArgumentNotValidException` → 400 (Bad Request)
    - `Exception` → 500 (Internal Server Error, genérico, sem stack trace)
  - Todas as respostas seguem RFC 7807: `type`, `title`, `status`, `detail`
  - Mensagens em PT-BR
  - `ErrorResponse` como DTO de resposta
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../exception/GlobalExceptionHandler.java` | 🆕 | @ControllerAdvice RFC 7807 |
  | `src/main/java/.../exception/BusinessException.java` | 🆕 | Exceção base para erros de negócio (422) |
  | `src/main/java/.../dto/response/ErrorResponse.java` | 🆕 | DTO RFC 7807 (type, title, status, detail) |
- **Dependências:** Nenhuma (independente)
- **Riscos:** Expor stack trace acidentalmente → testar todos os handlers
- **Skills aplicáveis:** `302-frameworks-spring-boot-rest`, `126-java-exception-handling`

---

## 4. Ordem de Execução

1. **T-009** — SecurityConfig (fundação de segurança)
2. **T-010** — JwtAuthenticationFilter (depende do SecurityConfig)
3. **T-011** — TenantContext (já existe, apenas complementar — depende do filter)
4. **T-013** — @RequiresPermission + RbacAspect (depende do TenantContext)
5. **T-012** — TenantIsolationAspect (depende do TenantContext)
6. **T-015** — GlobalExceptionHandler + exceções (independente — mas útil ter pronto para os testes)
7. **T-014** — @Auditable + AuditAspect (independente, @Async requer config adicional)

---

## 5. Estratégia de Build e Verificação

- **Comando de build:** `mvn compile`
- **Comando de teste:** `mvn test`
- **Comando de cobertura:** `mvn jacoco:check` (condicionado à compatibilidade JaCoCo × Java 25)
- **Checkpoints:**
  1. Após T-010: build + smoke test (requisição sem token → 401)
  2. Após T-013: build + teste RBAC (acesso negado → 403)
  3. Após T-015: build + teste RFC 7807 (erro formatado)
  4. Final: `mvn test` com todos os cenários da sprint

---


🤖 *Revisão Caveman em 15/07/2026 (DOCS-SERVICE-SPRINTS-CAVEMAN-REVIEW.md): RLS 11→5 tabelas, SQL injection corrigido (concatenação→PreparedStatement).*
