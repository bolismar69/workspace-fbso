# SPRINT-DEVELOPMENT-PLANNING.md — Plano de Desenvolvimento: Sprint 3 — Frente 0

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 3 de 7 — Portal Admin + Contas e Planos
- **Frente:** 0 — Correções Pré-Sprint (12 Débitos Técnicos Impeditivos)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + JDBC Template
- **Data:** 17/07/2026
- **Branch:** `feature/sprint-03-portal-admin`

---

## 1. Visão Geral

- **Frente Goal:** Resolver bugs e vulnerabilidades que IMPEDEM o início da implementação das features da Sprint 3
- **Tasks a implementar:** 12 (T-015.2.DT-001 a T-015.13.DT-012)
- **Ordem de execução:** Sequencial (não paralelizável por ser 1 dev)
- **Stack:** Java 25 (GraalVM 25.0.3) + Spring Boot 3.5.14 + PostgreSQL 17 + JDBC Template (ADR-L01)
- **Build atual:** `./mvnw test` → BUILD SUCCESS, 33/33 testes passando, 1 skipped

### Baseline (pré-Frente 0)

| Indicador | Valor |
|:---|:---|
| Spring Boot | 3.5.14 ✅ (já atualizado, DT-001 parcial) |
| Jackson | 2.21.4 ✅ (já atualizado, DT-001 parcial) |
| JaCoCo | 0.8.12 ❌ (incompatível com Java 25 class file v69) |
| Surefire | Já inclui `**/security/**/*Test.java` ✅ (DT-010 parcial) |
| spring-boot-starter-mail | ❌ Ausente |
| AuditAspect | ❌ @Async chama TenantContext depois do clear() |
| BaseRepository | ❌ Sem save()/update() genéricos |
| RbacAspect | ❌ Sem entradas TENANT/PLAN/SUBSCRIPTION/DASHBOARD |
| Exceções de domínio | ❌ 4 subclasses referenciadas no Javadoc não existem |

---

## 2. Dependências entre Tasks

```
T-015.2.DT-001 (pom.xml: verificar versões)
  │
  ├── T-015.8.DT-007 (pom.xml: spring-boot-starter-mail)
  │
  ├── T-015.5.DT-004 (pom.xml: JaCoCo compatível Java 25)
  │
  └── T-015.11.DT-010 (pom.xml: Surefire — verificar já corrigido)

T-015.4.DT-003 (BaseRepository.save/update)
  │
  └── T-015.3.DT-002 (AuditAspect — usa save/update do repo?)

T-015.3.DT-002 (AuditAspect: capturar tenantId/userId ANTES do @Async)
  │
  ├── T-015.9.DT-008 (AuditAspect.extractEntityId: corrigir extração)
  │
  └── T-015.7.DT-006 (TenantAwareDataSource: log.error + exceção)

T-015.6.DT-005 (RbacAspect: expandir matriz)

T-015.10.DT-009 (Migration V005: locked_price + locked_recurrence)

T-015.12.DT-011 (JwtAuthenticationFilter: ObjectMapper + ErrorResponse)

T-015.13.DT-012 (4 exceções de domínio)
```

> **Nota:** As dependências são majoritariamente de edição de arquivos (não de lógica). Tasks que tocam o mesmo arquivo são agrupadas.

---

## 3. Plano por Task

### T-015.2.DT-001 — Atualizar Dependências (Spring Boot + Jackson + logback + tomcat)

- **Critério DONE:** Build passa. 33/33 testes. CVEs de auth bypass (8.2) e RCE (8.1) eliminadas
- **Estimativa:** 2h
- **Abordagem:** Verificar versões atuais no pom.xml. Spring Boot 3.5.14 e Jackson 2.21.4 já estão configurados. Verificar se logback e tomcat (transitivas do parent) estão nas versões corrigidas. Spring Boot 3.5.14 já traz tomcat 10.1.x (sem CVEs) e logback 1.5.x.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `pom.xml` | 🔄 | Verificar e confirmar versões spring-boot-starter-parent 3.5.14, jackson 2.21.4 |
- **Dependências:** Nenhuma
- **Riscos:** Se versões transitivas de logback/tomcat ainda tiverem CVEs, override manual necessário
- **Skills aplicáveis:** `110-java-maven-best-practices`, `111-java-maven-dependencies`

### T-015.3.DT-002 — Refatorar AuditAspect: @Async Thread Safety

- **Critério DONE:** Auditoria funcional: registro com tenant_id e user_id corretos. TesteAuditAspect criado
- **Estimativa:** 6h
- **Abordagem:** O problema: `@Async` executa em thread separada. Quando a thread do pool executa `audit()`, o `JwtAuthenticationFilter.finally` já chamou `TenantContext.clear()`. Solução: capturar `tenantId` e `userId` no JoinPoint (thread principal) e passá-los como parâmetros para o método assíncrono.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/test/java/.../unit/security/AuditAspectTest.java` | 🆕 | Teste unitário do AuditAspect com mock JdbcTemplate |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../security/aspect/AuditAspect.java` | 🔄 | Capturar tenantId/userId ANTES do @Async |
- **Dependências:** T-015.4.DT-003 (BaseRepository.save/update — contexto de como entidades são persistidas)
- **Riscos:** Mudança de assinatura do método @Async — verificar se Spring proxy não quebra
- **Skills aplicáveis:** `126-java-exception-handling`, `301-frameworks-spring-boot-core`, `131-java-testing-unit-testing`

### T-015.4.DT-003 — Adicionar save(T) e update(T) ao BaseRepository

- **Critério DONE:** INSERT/UPDATE com created_by/updated_by automáticos. Testes unitários para save/update
- **Estimativa:** 4h
- **Abordagem:** Gerar INSERT dinâmico com reflection ou aceitar Map<String, Object> de colunas. A abordagem mais idiomática com JDBC: cada entidade fornece um método `toColumnMap()` que retorna as colunas. O BaseRepository.save() chama `entity.toColumnMap()`, adiciona `created_by`/`created_dt`, e gera o INSERT. O update() similar com `updated_by`/`updated_dt`.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../repository/common/BaseRepository.java` | 🔄 | Adicionar save(T) e update(T) genéricos |
  | `src/main/java/.../common/BaseEntity.java` | 🔄 | Adicionar método abstrato toColumnMap() |
  | `src/test/java/.../unit/repository/BaseRepositoryTest.java` | 🔄 | Adicionar testes para save/update |
- **Dependências:** Nenhuma (infraestrutura já existe)
- **Riscos:** Complexidade da abordagem de reflection vs Map — escolher a mais simples e idiomática
- **Skills aplicáveis:** `121-java-object-oriented-design`, `311-frameworks-spring-jdbc`, `131-java-testing-unit-testing`

### T-015.5.DT-004 — Atualizar JaCoCo para Compatibilidade com Java 25

- **Critério DONE:** Relatório JaCoCo gerado sem erro "class file major version 69". Meta ≥80% verificável
- **Estimativa:** 3h
- **Abordagem:** JaCoCo 0.8.12 não suporta Java 25 (class file v69). Verificar qual versão do JaCoCo suporta Java 25. Se 0.8.14+ não suportar, avaliar 0.8.15-SNAPSHOT ou substituir por OpenClover. Alternativa: usar JaCoCo 0.8.13 com flag `-Xmax-classfile-name` ou aguardar release oficial.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `pom.xml` | 🔄 | Atualizar `${jacoco.version}` para versão compatível com Java 25 |
- **Dependências:** Nenhuma
- **Riscos:** ALTO — JaCoCo pode não ter release oficial com suporte a Java 25. Plano B: desabilitar check do JaCoCo temporariamente e verificar cobertura via IntelliJ ou OpenClover
- **Skills aplicáveis:** `110-java-maven-best-practices`, `112-java-maven-plugins`

### T-015.6.DT-005 — Expandir RbacAspect: TENANT, PLAN, SUBSCRIPTION, DASHBOARD

- **Critério DONE:** Endpoints da Sprint 3 não retornam 403 para ADMIN_TENANT. Matriz atualizada
- **Estimativa:** 2h
- **Abordagem:** Adicionar `TENANT`, `PLAN`, `SUBSCRIPTION`, `DASHBOARD` aos conjuntos `MANAGER_RESOURCES` e `OPERATOR_RESOURCES` com ações apropriadas. MANAGER_BU: view em todos, create/edit em TENANT/PLAN/SUBSCRIPTION. OPERATOR_BU: apenas view. AUDITOR: apenas AUDIT (já existe).
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../security/aspect/RbacAspect.java` | 🔄 | Expandir matriz com 4 novos resources |
  | `src/test/java/.../unit/security/RbacAspectTest.java` | 🔄 | Adicionar cenários para novos resources |
- **Dependências:** T-015.13.DT-012 (exceções de domínio — usadas pelo aspecto)
- **Riscos:** Matriz hardcoded cresce — risco de manutenção até Sprint 4
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `131-java-testing-unit-testing`

### T-015.7.DT-006 — TenantAwareDataSource: log.error + TenantIsolationException

- **Critério DONE:** Se SET falhar, conexão NÃO retorna ao pool. Erro logado como ERROR
- **Estimativa:** 30min
- **Abordagem:** Trocar `log.debug` → `log.error` no catch. Após logar, lançar `TenantIsolationException` (runtime) para impedir que a conexão com tenant_id residual volte ao pool.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../exception/TenantIsolationException.java` | 🆕 | Exceção específica para falha de isolamento |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../config/TenantAwareDataSource.java` | 🔄 | log.debug → log.error + throw |
  | `src/test/java/.../unit/config/TenantAwareDataSourceTest.java` | 🔄 | Atualizar testes |
- **Dependências:** Nenhuma
- **Riscos:** Baixo — mudança pontual de 3 linhas
- **Skills aplicáveis:** `126-java-exception-handling`

### T-015.8.DT-007 — Adicionar spring-boot-starter-mail

- **Critério DONE:** JavaMailSender disponível para injeção. T-028 desbloqueada
- **Estimativa:** 10min
- **Abordagem:** Descomentar/adicionar a dependência `spring-boot-starter-mail` no pom.xml. A configuração de email no `application.yml` já existe.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `pom.xml` | 🔄 | Adicionar spring-boot-starter-mail |
- **Dependências:** Nenhuma
- **Riscos:** Nenhum
- **Skills aplicáveis:** `110-java-maven-best-practices`

### T-015.9.DT-008 — Corrigir extractEntityId() no AuditAspect

- **Critério DONE:** entity_id nos registros de auditoria corresponde à entidade real
- **Estimativa:** 2h
- **Abordagem:** O método atual usa `args[0].toString()` que retorna `Classe@hash` para objetos. Solução: usar a anotação `@Auditable` para especificar qual parâmetro contém o ID, ou tentar `UUID.fromString()` com fallback para reflection (procurar método `getId()` no argumento).
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../security/aspect/AuditAspect.java` | 🔄 | Refatorar extractEntityId() |
  | `src/main/java/.../security/annotation/Auditable.java` | 🔄 | Adicionar campo `idParamIndex` (opcional) |
- **Dependências:** T-015.3.DT-002 (AuditAspect refatorado)
- **Riscos:** Médio — abordagem de reflection pode ser frágil. Preferir anotação explícita
- **Skills aplicáveis:** `121-java-object-oriented-design`, `301-frameworks-spring-boot-core`

### T-015.10.DT-009 — Migration V005: locked_price + locked_recurrence

- **Critério DONE:** RN06-02 atendida. Alteração de preço no plano não afeta assinaturas existentes
- **Estimativa:** 3h
- **Abordagem:** Criar migration V005 que adiciona colunas `locked_price NUMERIC(10,2)` e `locked_recurrence VARCHAR(20)` na tabela `subscription`. Criar migration de rollback U005. Quando uma subscription é criada, copiar o `price` e `recurrence` atuais do plano para essas colunas.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/resources/db/migration/V005__add_locked_price_to_subscription.sql` | 🆕 | ALTER TABLE subscription ADD locked_price, locked_recurrence |
  | `src/main/resources/db/migration/U005__remove_locked_price_from_subscription.sql` | 🆕 | Rollback |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../entity/Tenant.java` | 🔄 | Se Subscription entity existir, adicionar campos |
- **Dependências:** Nenhuma (migration independente)
- **Riscos:** Se a tabela subscription ainda não existir (não foi criada nas migrations V001-V004), esta migration falhará. Verificar V001.
- **Skills aplicáveis:** `313-frameworks-spring-db-migrations-flyway`, `postgres-pro`

### T-015.11.DT-010 — Surefire: Verificar Padrão de Inclusão

- **Critério DONE:** JwtAuthenticationFilterTest executado no build
- **Estimativa:** 5min (verificação)
- **Abordagem:** O pom.xml já inclui `**/security/**/*Test.java` (linha 231). Verificar se `JwtAuthenticationFilterTest` está no caminho correto (`src/test/java/.../security/JwtAuthenticationFilterTest.java`). Se sim, esta task é NO-OP.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `pom.xml` | 🔄 | Confirmar que padrão captura todos os testes de segurança |
- **Dependências:** Nenhuma
- **Riscos:** Nenhum — parece já corrigido
- **Skills aplicáveis:** `110-java-maven-best-practices`

### T-015.12.DT-011 — Reescrever sendUnauthorized() com ObjectMapper + ErrorResponse

- **Critério DONE:** JSON 401 consistente com RFC 7807. Sem injection via mensagem
- **Estimativa:** 30min
- **Abordagem:** Substituir `response.getWriter().write("""...""".formatted(message))` por `new ObjectMapper().writeValue(response.getWriter(), new ErrorResponse(...))`. Isso garante escape JSON correto e consistência com o resto da API.
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../security/JwtAuthenticationFilter.java` | 🔄 | Reescrever sendUnauthorized() |
- **Dependências:** Nenhuma
- **Riscos:** Nenhum — ObjectMapper já disponível via spring-boot-starter-web
- **Skills aplicáveis:** `302-frameworks-spring-boot-rest`, `124-java-secure-coding`

### T-015.13.DT-012 — Criar 4 Subclasses de BusinessException

- **Critério DONE:** Services podem lançar exceções específicas. Catch semântico possível
- **Estimativa:** 30min
- **Abordagem:** Criar 4 classes no pacote `exception/`: `DuplicateCnpjException`, `InvalidStatusTransitionException`, `PlanHasActiveSubscribersException`, `TenantNotFoundException`. Todas estendem `BusinessException` com construtor que recebe mensagem e define errorCode.
- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../exception/DuplicateCnpjException.java` | 🆕 | Construtor com mensagem |
  | `src/main/java/.../exception/InvalidStatusTransitionException.java` | 🆕 | Construtor com mensagem |
  | `src/main/java/.../exception/PlanHasActiveSubscribersException.java` | 🆕 | Construtor com mensagem |
  | `src/main/java/.../exception/TenantNotFoundException.java` | 🆕 | Construtor com mensagem |
- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/.../exception/GlobalExceptionHandler.java` | 🔄 | Adicionar handlers específicos para cada nova exceção |
- **Dependências:** Nenhuma
- **Riscos:** Nenhum — classes simples
- **Skills aplicáveis:** `126-java-exception-handling`

---

## 4. Ordem de Execução

1. **T-015.11.DT-010** (5min) — Verificar Surefire (provavelmente NO-OP, já corrigido)
2. **T-015.8.DT-007** (10min) — Adicionar spring-boot-starter-mail (desbloqueia T-028)
3. **T-015.5.DT-004** (3h) — Atualizar JaCoCo (risco: compatibilidade Java 25)
4. **T-015.2.DT-001** (2h) — Verificar versões finais (Spring Boot/Jackson/logback/tomcat)
5. **T-015.13.DT-012** (30min) — Criar 4 exceções de domínio
6. **T-015.7.DT-006** (30min) — TenantAwareDataSource: log.error + throw
7. **T-015.6.DT-005** (2h) — Expandir RbacAspect
8. **T-015.4.DT-003** (4h) — BaseRepository.save() e update()
9. **T-015.3.DT-002** (6h) — Refatorar AuditAspect @Async
10. **T-015.9.DT-008** (2h) — Corrigir extractEntityId()
11. **T-015.12.DT-011** (30min) — Reescrever sendUnauthorized()
12. **T-015.10.DT-009** (3h) — Migration V005

**Ordem justificada:** Tasks de baixo risco e infraestrutura primeiro (pom.xml, exceções); aspectos cross-cutting depois (RbacAspect, AuditAspect); migration por último (depende do schema existente).

---

## 5. Estratégia de Build e Verificação

- **Comando de build:** `./mvnw compile`
- **Comando de teste:** `./mvnw test`
- **Comando de coverage:** `./mvnw verify` (quando JaCoCo estiver funcional)
- **Checkpoints:**
  - Após cada task de pom.xml → `./mvnw compile`
  - Após cada task de código → `./mvnw test`
  - Ao final da Frente 0 → `./mvnw verify` (com meta ≥80%)

---

🤖 *Gerado por Agente IA em 17/07/2026. Stack detectada: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + JDBC Template. Skills: 110-java-maven-best-practices, 121-java-object-oriented-design, 126-java-exception-handling, 130-java-testing-strategies, 301-frameworks-spring-boot-core, 304-frameworks-spring-boot-security, 311-frameworks-spring-jdbc, postgres-pro.*
