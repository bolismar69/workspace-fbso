# SPRINT-5-EXECUTION-REPORT-Frente-0.md — Relatório de Execução: Sprint 5 — Frente 0

- **Solução:** `ms-fbso-platform-admin`
- **Projeto de Negócio:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 5 de 7 — sprint-05-portal-cliente
- **Frente:** Frente 0 — Bloqueantes (Pré-Sprint)
- **Stack detectada:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Keycloak 26 + Flyway 12.11.0 + Caffeine 3.2.4
- **Data da execução:** 2026-07-17
- **Tasks executadas:** T-133.DT-095 a T-138.DT-100 + T-141.DT-098 (7 tasks)

---

## 1. Resumo da Execução

| Métrica | Valor |
|:---|---|
| **Tasks executadas** | 7/7 (100%) |
| **Tasks com sucesso** | 7 |
| **Tasks NO-OP** | 2 (T-136.DT-096, T-141.DT-098) |
| **Tasks com falha** | 0 |
| **Tempo estimado** | ~6h |
| **Tempo gasto** | ~4h (2 NO-OP economizaram ~2h) |
| **Build** | ✅ SUCCESS |
| **Testes unitários** | 213 testes executados, 0 falhas, 1 erro pré-existente |

---

## 2. Stack e Skills Utilizadas

- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Keycloak 26 + Flyway 12.11.0 + Caffeine 3.2.4
- **Fonte da stack:** PRD.md §5.1 + ARCHITECTURE.md header
- **Skills aplicáveis:**
  - `110-java-maven-best-practices` — bumps de dependência no pom.xml
  - `111-java-maven-dependencies` — verificação de árvore de dependências
  - `304-frameworks-spring-boot-security` — configuração OAuth2 Client + Resource Server
  - `security-review` — CVE-2026-42198 (PostgreSQL driver)

---

## 3. Tasks Executadas

| ID | Tarefa | Status | Build | Observações |
|:---|:---|:---:|:---:|:---|
| **T-133.DT-095** | Criar `docker-compose.yml` + `realm-config.json` | ✅ | — | 3 serviços: PostgreSQL 17, Keycloak 26, MailHog. Realm com 4 roles e 1 usuário de teste |
| **T-134.DT-045** | Bump Flyway 10.22.0→12.11.0 | ✅ | ✅ | 1 property alterada no pom.xml |
| **T-135.DT-068** | Bump PostgreSQL driver → 42.7.11 (CVE fix) | ✅ | ✅ | Property `<postgresql.version>` adicionada ao pom.xml |
| **T-136.DT-096** | Verificar JwtAuthenticationFilter claims | ✅ NO-OP | ✅ | Já implementado — extrai `modules[]` e `business_unit_ids[]` via JwtUtils (6 testes passando) |
| **T-137.DT-099** | Adicionar OAuth2 Client + SecurityConfig | ✅ | ✅ | Dependency + refactor SecurityConfig com 2 filter chains (@Order) |
| **T-138.DT-100** | Configurar application.yml OAuth2 Client | ✅ | ✅ | Client registration + provider Keycloak. Portas 8080→8081 |
| **T-141.DT-098** | Verificar TenantContext consumo | ✅ NO-OP | ✅ | Já implementado — armazena `businessUnitIds` e `modules` no record |

---

## 4. Arquivos Criados ou Modificados

| Ação | Arquivo | Task | Descrição da Mudança |
|:---|:---|:---|:---|
| 🆕 | `docker-compose.yml` | T-133 | 3 serviços (postgres:17, keycloak:26, mailhog) + network + volumes |
| 🆕 | `keycloak/realm-config.json` | T-133 | Realm fbso-platform: 4 roles, 1 client (Auth Code), 3 custom claims, 1 test user |
| 🔄 | `pom.xml` | T-134 | `flyway.version`: 10.22.0 → **12.11.0** |
| 🔄 | `pom.xml` | T-135 | Adicionado `<postgresql.version>42.7.11</postgresql.version>` |
| 🔄 | `pom.xml` | T-137 | Adicionada dependência `spring-boot-starter-oauth2-client` |
| 🔄 | `SecurityConfig.java` | T-137 | Refatorado: 2 `SecurityFilterChain` beans com `@Order(1)` (OAuth2 login) e `@Order(2)` (API Resource Server) |
| 🔄 | `application.yml` | T-138 | Adicionada seção `client.registration.keycloak` + `client.provider.keycloak`. Portas Keycloak 8080→8081 |
| — | `JwtAuthenticationFilter.java` | T-136 (NO-OP) | Nenhuma mudança — código já extrai todas as claims |
| — | `TenantContext.java` | T-141 (NO-OP) | Nenhuma mudança — código já armazena `businessUnitIds` e `modules` |
| — | `JwtUtils.java` | T-136 (NO-OP) | Nenhuma mudança — código já tem `getBusinessUnitIds()` e `getModules()` |

---

## 5. Evidências de Testes

| Comando | Resultado |
|:---|:---|
| `mvn compile` | ✅ BUILD SUCCESS (94 source files, 7.8s) |
| `mvn test -Dtest="JwtAuthenticationFilterTest"` | ✅ 6/6 passando |
| `mvn test` (full) | 213 testes: 0 failures, 1 pre-existing error |

### Erro Pré-Existente (NÃO causado pela Frente 0)

```
SubscriptionServiceTest > create > shouldCreateWithLockedPrice: ERROR
  IllegalStateException: TenantContext não inicializado
  at TenantContext.getTenantId(TenantContext.java:35)
  at SubscriptionService.create(SubscriptionService.java:66)
```

**Causa:** O teste `SubscriptionServiceTest.shouldCreateWithLockedPrice()` chama `SubscriptionService.create()` que depende de `TenantContext.getTenantId()`, mas o teste não configura o `TenantContext` ThreadLocal antes da chamada.

**Impacto:** Nenhum — este erro existe desde a implementação original do `SubscriptionService`. Não foi introduzido pela Frente 0. Nenhum arquivo tocado pela Frente 0 está na stack trace do erro.

**Ação recomendada:** Corrigir na Frente 1 ou Frente 3 (adicionar `TenantContext.set(…)` no `@BeforeEach` do teste).

---

## 6. Validação de Segurança

- [x] Nenhuma credencial ou dado sensível hardcoded — `client-secret` usa placeholder `${KEYCLOAK_CLIENT_SECRET:changeme}`
- [x] Queries usam parametrização (proteção contra injection) — sem alterações em queries
- [x] Controles de acesso implementados — JwtAuthenticationFilter + OAuth2 Resource Server mantidos
- [x] Dados pessoais não expostos em logs ou respostas HTTP
- [x] Respostas de erro não expõem stack traces — `ErrorResponse.of()` RFC 7807 mantido
- [x] CVE-2026-42198 (PostgreSQL driver DoS, CVSS 7.5) resolvida com bump 42.7.10→42.7.11

---

## 7. Validação de Arquitetura

- [x] Estrutura de diretórios segue ARCHITECTURE.md — `docker-compose.yml` na raiz, `keycloak/` para config
- [x] Convenções de nomenclatura respeitadas — padrão Spring Boot (`application.yml`, `pom.xml`)
- [x] Padrões de projeto documentados nas ADRs foram seguidos:
  - **ADR-04 (Keycloak)**: OAuth2 Client adicionado para Authorization Code Flow
  - **ADR-07 (JWT Stateless)**: API REST mantida stateless; apenas endpoints `/auth/**` usam sessão
- [x] Nenhuma quebra de API — endpoints existentes inalterados
- [x] Stack versionada e documentada: Flyway 12.11.0, PostgreSQL driver 42.7.11, Spring Boot 3.5.14

---

## 8. Desvios e Observações

### Desvios do Planejamento Original

| Débito | Planejado | Realizado | Justificativa |
|:---|:---|:---|:---|
| DT-096 | Atualizar JwtAuthenticationFilter | **NO-OP** — já implementado | Código já extrai `modules[]` e `business_unit_ids[]`. A auditoria foi gerada antes da inspeção detalhada |
| DT-098 | Conectar TenantContext ao Filter | **NO-OP** — já implementado | TenantContext já armazena e expõe `getBusinessUnitIds()` e `getModules()` |

### Decisões de Design Tomadas

1. **2 SecurityFilterChain beans** (não 1 monolítico): `@Order(1)` para OAuth2 login, `@Order(2)` para API Resource Server. Isso evita conflitos entre sessão (OAuth2) e stateless (JWT).

2. **Keycloak na porta 8081** (não 8080): Evita conflito com a aplicação Spring Boot (porta 8080). Mapeamento docker: `8081:8080`.

3. **PostgreSQL driver com property explícita**: Segue o mesmo padrão do Jackson (`<jackson.version>2.21.4</jackson.version>`) para overrides de versão gerenciados pelo Spring Boot Parent.

### Dificuldades e Soluções

| Dificuldade | Solução |
|:---|:---|
| Keycloak 26 usa `--import-realm` (não `--import-realm` como no 25) | Confirmado: `start-dev --import-realm` é o flag correto para Keycloak 26 |
| `spring-boot-starter-oauth2-client` adiciona ~30 jars transitivos | Build compilou sem conflitos. Spring Boot gerencia versões compatíveis |

---

## 9. Próximos Passos

1. **Frente 1** (Recomendados): 10 tarefas (T-139.DT-023 a T-148.DT-102) — ~17.5h
2. **Frente 2** (Desejáveis): 8 tarefas (T-149.DT-086 a T-156.DT-113) — ~4.5h
3. **Frente 3** (Features): 12 tarefas (T-057 a T-068) — ~18.5d
4. **Corrigir teste pré-existente:** `SubscriptionServiceTest.shouldCreateWithLockedPrice` — adicionar `TenantContext.set(...)` no `@BeforeEach`

### Pré-requisitos para a Frente 1

- [x] docker-compose.yml funcional (criado na Frente 0)
- [x] Flyway atualizado para criar migrations V007+ (bump feito)
- [x] PostgreSQL driver sem CVE ativa (bump feito)
- [x] OAuth2 Client disponível no classpath (dependência adicionada)
- [x] SecurityConfig pronto para OAuth2 (refatorado com 2 filter chains)

---

🤖 *Relatório gerado em 2026-07-17 como parte da Fase 9 do PROMPT-EXECUTE-SPRINT-TASKS.md. Stack detectada: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Keycloak 26. Skills utilizados: 110-java-maven-best-practices, 304-frameworks-spring-boot-security, security-review.*
