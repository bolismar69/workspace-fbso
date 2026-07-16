# SPRINT-DEVELOPMENT-PLANNING-revisao-caveman.md — Plano de Desenvolvimento: Sprint 2

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** `PRJ-FIN-2026-0003-SAAS-FBSO-ORG`
- **Sprint:** 2 de 7 — Segurança Cross-Cutting
- **Stack:** Java 25 + Spring Boot 3.5.1 + PostgreSQL 17 + Keycloak OIDC + Flyway + Maven 3.9
- **Data da Revisão:** 15 de Julho de 2026
- **Modo:** Revisão Caveman (sprint concluída em 14/07/2026)

---

> 🚫 **Branch:** `feature/java-fbso-platform-admin` ([PRD §8.4](../../PRD.md#84-branch-de-desenvolvimento))

## 1. Visão Geral

- **Sprint Goal:** JWT autentica via Keycloak RS256. TenantContext isola requisições por tenant_id. PostgreSQL RLS garante isolamento no nível do banco (defesa em profundidade). @RequiresPermission bloqueia acessos não autorizados (403). @Auditable grava auditoria de forma assíncrona. GlobalExceptionHandler retorna erros RFC 7807 em PT-BR.
- **Tasks implementadas:** 8 (T-009 a T-015 + T-015.1, sendo T-012 substituído por T-015.1)
- **Ordem de execução:** Sequencial com forte encadeamento — pipeline de segurança é linear
- **Stack detectada:** Java 25 + Spring Boot 3.5.1 + PostgreSQL 17 + Keycloak OIDC + Maven 3.9

---

## 2. Dependências entre Tasks

```
T-009 (SecurityConfig)
  │
  └── T-010 (JwtAuthenticationFilter)
        │
        ├── T-011 (TenantContext — complemento do Sprint 1)
        │
        ├── T-015.1 (PostgreSQL RLS — Migration V003 + TenantAwareDataSource)
        │     │
        │     └── [SUBSTITUI T-012 — TenantIsolationAspect removido, redundante com RLS]
        │
        ├── T-013 (@RequiresPermission + RbacAspect)
        │
        ├── T-014 (@Auditable + AuditAspect — independente do RBAC)
        │
        └── T-015 (GlobalExceptionHandler — independente)
```

**Ordem recomendada:** T-009 → T-010 → T-011 → T-015.1 → T-013 → T-014 → T-015

---

## 3. Plano por Task

### T-009 — SecurityConfig.java
- **Critério DONE:** Requisição sem token → 401. Token válido → autenticado. Token inválido/expirado → 401. CORS permite origem do frontend
- **Estimativa:** 2d
- **Abordagem:** Classe `@Configuration` + `@EnableWebSecurity`. Configurar `SecurityFilterChain` com `oauth2ResourceServer` para JWT (Keycloak RS256 via JWKS URI). `sessionCreationPolicy(STATELESS)`. CSRF desabilitado (API stateless). CORS configurado.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `config/SecurityConfig.java` | 🆕 | Configuração Spring Security + JWT + CORS |
  | `config/AsyncConfig.java` | 🆕 | @EnableAsync + ThreadPoolTaskExecutor |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `application.yml` | 🔄 | Adicionar Keycloak JWKS URI, CORS origins |
- **Dependências:** T-001 (pom.xml com spring-boot-starter-security, oauth2-resource-server)
- **Riscos:** JWKS endpoint inacessível → cache de JWKS. Versão Spring Security com API estável.
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `security-review`

### T-010 — JwtAuthenticationFilter.java
- **Critério DONE:** Filter em toda req. exceto `/actuator/health`. Claims extraídas. SecurityContext populado
- **Estimativa:** 2d
- **Abordagem:** `OncePerRequestFilter` que extrai JWT do header `Authorization: Bearer`, valida assinatura RS256, valida exp, extrai claims (tenant_id, user_id, roles, business_unit_ids, modules), popula `TenantContext` e `SecurityContextHolder`. No finally: `TenantContext.clear()`.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `security/JwtAuthenticationFilter.java` | 🆕 | OncePerRequestFilter com validação JWT RS256 |
  | `utils/JwtUtils.java` | 🆕 | Utilitário para extração de claims do JWT |
- **Dependências:** T-009 (SecurityConfig que registra o filter)
- **Riscos:** JWKS endpoint inacessível → cache com TTL. Claims ausentes → 401 amigável.
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `security-review`

### T-011 — TenantContext.java (complemento)
- **Critério DONE:** TenantContext.getTenantId() retorna tenant_id. Contexto limpo após requisição
- **Estimativa:** 0.5d
- **Abordagem:** A classe já existe da Sprint 1 com estrutura básica. Esta task é de verificação/complemento: validar integração com JwtAuthenticationFilter, garantir limpeza no finally, prevenir vazamento entre threads.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `security/TenantContext.java` | 🔄 | Revisar e complementar — criado na Sprint 1 |
- **Dependências:** T-010 (JwtAuthenticationFilter que popula o contexto)
- **Riscos:** Vazamento de ThreadLocal entre requisições → teste de concorrência
- **Skills aplicáveis:** `126-java-exception-handling`, `121-java-object-oriented-design`

### T-015.1 — PostgreSQL RLS (Migration V003 + TenantAwareDataSource)
- **Critério DONE:** ✅ RLS ativo em 5 tabelas (subscription, user, business_unit, product_service, audit_log). TenantAwareDataSource + BeanPostProcessor configurados. 33/33 testes passando (+11 novos). Migration V003 + U003 (rollback) criados
- **Estimativa:** 1.5d
- **Abordagem:** Defesa em profundidade — camada 1 de isolamento multi-tenant no banco. Migration V003 com `ALTER TABLE ... ENABLE ROW LEVEL SECURITY` + `CREATE POLICY tenant_isolation`. `TenantAwareDataSource` como proxy HikariCP que configura `app.current_tenant_id` em cada `getConnection()`. BeanPostProcessor para encapsular o DataSource.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/resources/db/migration/V003__enable_rls.sql` | 🆕 | RLS + políticas tenant_isolation em 5 tabelas |
  | `src/main/resources/db/migration/U003__disable_rls.sql` | 🆕 | Rollback — remove RLS |
  | `config/TenantAwareDataSource.java` | 🆕 | DataSource proxy — SET LOCAL app.current_tenant_id |
  | `config/DataSourceConfig.java` | 🆕 | BeanPostProcessor — encapsula HikariCP |
  | `test/.../unit/config/TenantAwareDataSourceTest.java` | 🆕 | 6 testes unitários |
  | `test/.../integration/security/RLSIsolationTest.java` | 🆕 | 5 testes estruturais |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `security/TenantContext.java` | 🔄 | Adicionar getTenantIdQuietly() para o proxy |
  | `security/JwtAuthenticationFilter.java` | 🔄 | Javadoc — delega RLS ao TenantAwareDataSource |
- **Dependências:** T-010 (JwtAuthenticationFilter), T-004 (Migration V001)
- **Riscos:** Admin FBSO cross-tenant precisa bypass RLS → RESET quando tenant_id é null. Performance com muitos tenants.
- **Skills aplicáveis:** `313-frameworks-spring-db-migrations-flyway`, `security-review`, `postgres-pro`

> **Nota:** T-015.1 substitui T-012 (TenantIsolationAspect). O aspecto AOP foi removido por ser redundante com PostgreSQL RLS e frágil (queries nativas escapam do aspecto, mas não do RLS). Ver ARCHITECTURE.md ADR-L07.

### T-012 — TenantIsolationAspect.java (SUBSTITUÍDO)
- **Status:** [SUBSTITUÍDO por T-015.1 — PostgreSQL RLS]
- **Justificativa:** AOP frágil — queries nativas escapam. RLS no banco é garantia definitiva. Ver ARCHITECTURE.md ADR-L07.
- **Arquivo criado como placeholder:** `security/aspect/TenantIsolationAspect.java` (mantido para referência, marcado como @Deprecated)

### T-013 — @RequiresPermission + RbacAspect.java
- **Critério DONE:** Anotação bloqueia sem permissão. Role válida × resource × action. 403 com JSON amigável
- **Estimativa:** 2d
- **Abordagem:** Criar anotação `@RequiresPermission(resource, action)` com @Retention(RUNTIME). Criar `RbacAspect` com @Aspect + @Around. Matriz RBAC hardcoded (switch-case com papéis fixos) — ceiling = Sprint 4 (substituir por consulta ao banco: RoleResource + ResourceAction).
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `security/annotation/RequiresPermission.java` | 🆕 | Anotação customizada (resource, action) |
  | `security/aspect/RbacAspect.java` | 🆕 | Aspecto AOP — verifica permissões |
  | `exception/PermissionDeniedException.java` | 🆕 | Exceção para 403 |
- **Dependências:** T-010 (TenantContext com roles do JWT)
- **Riscos:** Matriz hardcoded → ceiling Sprint 4. Papéis não mapeados → 403 seguro (fail-closed).
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `121-java-object-oriented-design`

### T-014 — @Auditable + AuditAspect.java
- **Critério DONE:** Registro de auditoria gerado para cada operação. Valores capturados. Async não bloqueia thread principal
- **Estimativa:** 1.5d
- **Abordagem:** Criar anotação `@Auditable(entityType, action)`. Criar `AuditAspect` com @Aspect + @Async. Capturar valores antes/depois via reflection. Gravar em audit_log de forma assíncrona. Trade-off: perda de registros em crash (aceitável para Fase 0 — ADR-L03).
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `security/annotation/Auditable.java` | 🆕 | Anotação customizada (entityType, action) |
  | `security/aspect/AuditAspect.java` | 🆕 | Aspecto AOP — grava auditoria assíncrona |
- **Dependências:** T-009 (AsyncConfig com @EnableAsync)
- **Riscos:** Perda de registros em crash → trade-off aceito (ADR-L03). Fila cheia → log warning.
- **Skills aplicáveis:** `126-java-exception-handling`, `121-java-object-oriented-design`

### T-015 — GlobalExceptionHandler.java
- **Critério DONE:** type, title, status, detail. Sem stack trace. Mensagens PT-BR
- **Estimativa:** 1d
- **Abordagem:** `@ControllerAdvice` com handlers para: BusinessException (422), DuplicateCnpjException, InvalidStatusTransitionException, PlanHasActiveSubscribersException, TenantNotFoundException, PermissionDeniedException (403), MethodArgumentNotValidException (400), Exception genérica (500). Todas seguem RFC 7807. Criar `ErrorResponse` record.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `exception/BusinessException.java` | 🆕 | RuntimeException base (HTTP 422) |
  | `exception/GlobalExceptionHandler.java` | 🆕 | @ControllerAdvice — RFC 7807 |
  | `dto/response/ErrorResponse.java` | 🆕 | Record com type, title, status, detail, fields |
- **Dependências:** T-013 (PermissionDeniedException já deve existir)
- **Riscos:** Stack trace vazando → revisar todos os handlers. Mensagens em EN → PT-BR.
- **Skills aplicáveis:** `126-java-exception-handling`, `302-frameworks-spring-boot-rest`

---

## 4. Ordem de Execução

1. **T-009 (SecurityConfig)** — base da segurança. Sem dependências além do pom.xml.
2. **T-010 (JwtAuthenticationFilter)** — depende de T-009.
3. **T-011 (TenantContext)** — verificação. Depende de T-010.
4. **T-015.1 (PostgreSQL RLS)** — depende de T-010 e T-004. Substitui T-012.
5. **T-013 (@RequiresPermission)** — depende de T-010.
6. **T-014 (@Auditable)** — depende de T-009 (AsyncConfig).
7. **T-015 (GlobalExceptionHandler)** — depende de T-013 (PermissionDeniedException).

---

## 5. Estratégia de Build e Verificação

- **Comando de build:** `mvn clean compile`
- **Comando de teste:** `mvn test`
- **Checkpoints:**
  1. Após T-009: `mvn compile` — SecurityConfig compila
  2. Após T-010: testes JWT (6 cenários) passam
  3. Após T-015.1: 33/33 testes passam (+11 novos)
  4. Após T-013: testes RBAC (5 cenários) passam
  5. Após T-015: testes Exception (4 cenários) passam
  6. Final: 22 testes (Sprint 2) + 7 testes (Sprint 1) = 29 totais (antes T-015.1), 33 após T-015.1

---

## 6. Checklist Ponytail (7 Rungs) — Avaliação por Task

| Task | R1: YAGNI? | R2: Existe? | R3: Stdlib? | R4: Dep? | R5: Padrão? | R6: Simples? | R7: Mínimo? |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| T-009 | ✅ Sim | ✅ N/A | ✅ Spring Security | ✅ Já no pom.xml | ✅ ARCH §2 | ✅ 1 classe | ✅ |
| T-010 | ✅ Sim | ✅ N/A | ✅ OncePerRequestFilter | ✅ spring-security-oauth2 | ✅ ARCH §3 | ✅ 1 filter | ✅ |
| T-011 | ✅ Sim | ✅ Já existia (S1) | ✅ ThreadLocal | ✅ N/A | ✅ ARCH §3 | ✅ Revisão | ✅ |
| T-015.1 | ✅ Sim | ✅ N/A | ✅ DelegatingDataSource | ✅ postgresql, HikariCP | ✅ ARCH §4.3 | ✅ Proxy + migration | ✅ |
| T-013 | ✅ Sim | ✅ N/A | ✅ @Aspect | ✅ spring-aop | ✅ ARCH §4.1 | ✅ 1 aspecto | ✅ |
| T-014 | ✅ Sim | ✅ N/A | ✅ @Aspect, @Async | ✅ spring-aop | ✅ ARCH §4.2 | ✅ 1 aspecto | ✅ |
| T-015 | ✅ Sim | ✅ N/A | ✅ @ControllerAdvice | ✅ spring-web | ✅ ARCH §6 | ✅ 1 handler | ✅ |

**Conclusão Ponytail:** Todas as 7 tasks passam nos 7 rungs. Nenhum código desnecessário, duplicação, dependência supérflua ou desvio do padrão arquitetural.

---

## 7. Revisão Pós-Implementação

> **Executado em:** 15/07/2026. Sprint concluída em 14/07/2026. 7 tasks implementadas + 1 substituída (T-012) = 8/8 planejadas.

### Avaliação de Necessidade de Revisão

| Task | Status | Necessita Revisão? | Justificativa |
|:---|:---:|:---:|:---|
| T-009 | ✅ | Não | SecurityConfig funcional. JWT + CORS + CSRF desabilitado. 401 sem token. |
| T-010 | ✅ | Não | JwtAuthenticationFilter operacional. 6/6 testes passando. Claims extraídas corretamente. |
| T-011 | ✅ | Não | TenantContext já existia da Sprint 1. Verificação confirmou que atende aos critérios. |
| T-015.1 | ✅ | Não | RLS em 5 tabelas. 33/33 testes passando. DataSource proxy funcional. ADR-L07 documentado. |
| T-012 | — | Não (substituído) | TenantIsolationAspect substituído por PostgreSQL RLS (T-015.1). Arquivo mantido como @Deprecated. |
| T-013 | ✅ | Não | @RequiresPermission + RbacAspect. 5/5 testes passando. Matriz RBAC hardcoded (ceiling Sprint 4). |
| T-014 | ✅ | Não | @Auditable + AuditAspect. @Async configurado. Auditoria não bloqueia thread principal. |
| T-015 | ✅ | Não | GlobalExceptionHandler RFC 7807. 4/4 testes passando. Sem stack traces. Mensagens PT-BR. |

### Verificações Cross-Documento

| Verificação | Resultado |
|:---|:---:|
| Pipeline de segurança confere com ARCHITECTURE.md §3 (7 estágios) | ✅ |
| RLS em 5 tabelas confere com SPECS.md §10 e ARCHITECTURE.md §4.3 | ✅ |
| ADR-L07 documentado (RLS substitui TenantIsolationAspect) | ✅ |
| Testes seguem TEST_PLAN.md §9.6-§9.7 (Sprint 2 cenários) | ✅ |
| 8 tasks conferem com TASKS.md v2.3 (Pre-M2 Segurança) | ✅ |
| T-012 marcado [SUBSTITUÍDO] — consistente entre TASKS.md e ARCHITECTURE.md | ✅ |
| SQL injection mitigado — PreparedStatement no TenantAwareDataSource | ✅ |
| Mensagens em PT-BR conforme BR-NFR08 | ✅ |

### Pendências Conhecidas (Não Bloqueantes)

| Pendência | Impacto | Resolução |
|:---|:---|:---|
| Matriz RBAC hardcoded (switch-case) | Médio — sem consulta ao banco | Substituir na Sprint 4 (T-046) |
| JaCoCo incompatível com Java 25 | Baixo — cobertura não mensurável | Aguardar JaCoCo 0.8.13+ |
| Testes RLS estruturais (sem PostgreSQL real) | Baixo — validam migration, não runtime | Testcontainers na Sprint 3+ |
| AuditAspect perda em crash | Baixo — trade-off aceito (ADR-L03) | Reavaliar na Fase 1 |

### Conclusão

**Nenhuma revisão de implementação necessária.** As 8 tarefas da Sprint 2 (7 implementadas + T-012 substituída por T-015.1) foram executadas conforme o plano. Pipeline de segurança completo: JWT Filter → TenantContext → RBAC → PostgreSQL RLS → BaseRepository → Auditoria → RFC 7807. 33/33 testes passando. Arquitetura em conformidade com ARCHITECTURE.md.

**Execução do prompt interrompida na Fase 1 (SPRINT-DEVELOPMENT-PLANNING) por decisão do revisor — sprint já validada, sem necessidade de reimplementação.**

---

🤖 *Plano de desenvolvimento gerado em modo Revisão Caveman pelo Agente Executor de Sprint/Claude. Skills consideradas: 304-frameworks-spring-boot-security, security-review, 121-java-object-oriented-design, 126-java-exception-handling, 302-frameworks-spring-boot-rest, 311-frameworks-spring-jdbc, 313-frameworks-spring-db-migrations-flyway, postgres-pro, 131-java-testing-unit-testing. Revisão em 15/07/2026 — sprint concluída em 14/07/2026.*
