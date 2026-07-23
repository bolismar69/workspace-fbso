# SPRINT-TEST-PLANNING-Frente-1.md — Plano de Testes: Sprint 5 — Frente 1

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 5 de 7 — sprint-05-portal-cliente
- **Frente:** Frente 1 — Recomendados (🟡)
- **Stack:** Java 25 · Spring Boot 3.5.14 · PostgreSQL 17 · Caffeine 3.2.4 · JUnit 5 · Mockito · Testcontainers
- **Data:** 2026-07-23

---

## 1. Visão Geral

- **Tasks implementadas:** 10 (9 ativas + 1 NO-OP)
- **Cenários de teste mapeados:** 6 (1 unit + 3 integration + 2 segurança)
- **Meta de cobertura:** ≥ 80% para novo código (linhas)
- **Ferramentas:** JUnit 5 + Mockito (unit) · Testcontainers + PostgreSQL 17 (integration) · Spring Security Test (segurança)

---

## 2. Mapeamento Task → Cenários de Teste

| Task | Cenário(s) | Nível | Ferramenta | Status |
|:---|:---|:---|:---|:---:|
| T-139.DT-023 | CT-F1-001: findAllKeyset — primeira página (sem lastId) | Unit | JUnit 5 + H2 | ⬜ |
| T-139.DT-023 | CT-F1-002: findAllKeyset — segunda página (com lastId) | Unit | JUnit 5 + H2 | ⬜ |
| T-142.DT-107 | CT-F1-003: Migration V007 executa sem erro | Integration | Flyway + Testcontainers | ⬜ |
| T-142.DT-107 | CT-F1-004: is_matrix DEFAULT false para BUs existentes | Integration | Testcontainers | ⬜ |
| T-144.DT-110 | CT-F1-005: RateLimit — 5 falhas → bloqueio 15min | Integration | JUnit 5 + Caffeine | ⬜ |
| T-144.DT-110 | CT-F1-006: RateLimit — mensagem exibe tempo restante | Integration | JUnit 5 + Caffeine | ⬜ |
| T-146.DT-121 | CT-F1-007: AuthenticationException → 401 RFC 7807 | Unit | MockMvc | ⬜ |
| T-146.DT-121 | CT-F1-008: Token expirado → 401 (não 500) | Unit | MockMvc | ⬜ |
| T-148.DT-102 | CT-F1-009: JwtAuthenticationConverter extrai claims | Unit | Mockito | ⬜ |
| T-148.DT-102 | CT-F1-010: TenantContext populado via converter | Integration | Spring Security Test | ⬜ |

---

## 3. Estratégia por Nível de Teste

### 3.1 Testes Unitários

- **Ferramenta:** JUnit 5 + Mockito + MockMvc
- **Padrão:** AAA (Arrange-Act-Assert)
- **Localização:** `src/test/java/com/fbso/platform/admin/`
- **O que mockar:** `JdbcTemplate` (para BaseRepository), `HttpServletRequest/Response` (para RateLimitFilter), `JwtDecoder` (para JwtAuthenticationFilter)
- **O que NÃO mockar:** Entidades (`BusinessUnit`, `RateLimitEntry`), DTOs (`ErrorResponse`), lógica de cache Caffeine

### 3.2 Testes de Integração

- **Ferramenta:** Testcontainers + PostgreSQL 17 + Flyway
- **Localização:** `src/test/java/com/fbso/platform/admin/integration/`
- **O que usar real:** PostgreSQL (container), Flyway migrations (V001-V007), Caffeine cache
- **Dados de seed:** Migration V004 (RBAC seed) + dados inseridos via `@BeforeEach`

### 3.3 Testes de Segurança

- **Foco:** Rate limiting (bloqueio pós-5-falhas), 401 padronizado (RFC 7807), isolamento cross-tenant
- **Localização:** `src/test/java/com/fbso/platform/admin/security/`

---

## 4. Ordem de Execução dos Testes

1. **Testes unitários** (sem dependências externas — rodam primeiro)
   - CT-F1-001, CT-F1-002: BaseRepository keyset
   - CT-F1-007, CT-F1-008: GlobalExceptionHandler 401
   - CT-F1-009: JwtAuthenticationConverter extração de claims

2. **Testes de integração** (dependem de container/DB)
   - CT-F1-003, CT-F1-004: Migration V007 + is_matrix
   - CT-F1-005, CT-F1-006: RateLimitFilter com Caffeine

3. **Testes de segurança** (dependem de contexto autenticado)
   - CT-F1-010: TenantContext populado via converter + isolamento

---

## 5. Comandos de Execução

| Propósito | Comando |
|:---|:---|
| Unit tests (Frente 1) | `mvn test -Dtest="BaseRepositoryTest,GlobalExceptionHandlerTest,RateLimitFilterTest"` |
| Integration tests | `mvn verify -Dtest="*IntegrationTest"` |
| Full suite | `mvn test` |
| Coverage | `mvn jacoco:report` |
| Coverage check | `mvn jacoco:check` (meta: ≥80% linhas) |

---

## 6. Ações Manuais ou Externas

### Ação 1: Validar Migration V007 em Ambiente Dev

- **Cenário(s) relacionado(s):** CT-F1-003, CT-F1-004
- **Quem executa:** Dev BD (humano)
- **Pré-condições:** PostgreSQL 17 rodando (docker-compose). Flyway 12.11.0 configurado.
- **Ambiente:** dev

**Passo a passo:**

1. Execute `docker compose up postgres` na raiz do projeto para iniciar o PostgreSQL
2. Execute `mvn flyway:migrate` para aplicar todas as migrations (V001-V007)
3. Verifique a saída: `Successfully applied 7 migrations`

**Resultado esperado:**

- Migration V007 aplicada sem erro
- Coluna `is_matrix` existe na tabela `fbso_platform.business_unit`
- Valor DEFAULT é `false`

**Se falhar:**

- Verificar se a migration V007 não tem conflito de sintaxe: `docker compose exec postgres psql -U fbso -d fbso_platform -c "\d business_unit"`
- Se a migration já foi aplicada em outra branch, executar `mvn flyway:clean` (cuidado: apaga todos os dados)

**Evidência a coletar:**

- [ ] Output do `mvn flyway:info` mostrando V007 como "Success"
- [ ] Screenshot do DBeaver/pgAdmin mostrando a coluna `is_matrix`

### Ação 2: Validar Rate Limit com Teste Manual via curl

- **Cenário(s) relacionado(s):** CT-F1-005, CT-F1-006
- **Quem executa:** QA Engineer (humano)
- **Pré-condições:** Aplicação rodando (`mvn spring-boot:run`). Keycloak disponível (docker-compose).
- **Ambiente:** dev

**Passo a passo:**

1. Execute `curl -X POST http://localhost:8080/api/v1/auth/login -H "Content-Type: application/json" -d '{"username":"test@fbso.org","password":"wrong"}'` 5 vezes
2. Na 6ª tentativa, observe o código HTTP e o body da resposta

**Resultado esperado:**

- 5 primeiras respostas: 401 (Unauthorized)
- 6ª resposta: 429 (Too Many Requests) com body:
  ```json
  {
    "type": "https://api.fbso.org/errors/rate-limit",
    "title": "Muitas tentativas de login",
    "status": 429,
    "detail": "Conta bloqueada por 15 minutos. Tente novamente em X segundos."
  }
  ```
- Após 15 minutos, novas tentativas são permitidas

**Se falhar:**

- Verificar se o RateLimitFilter está registrado na cadeia: procure `RateLimitFilter` nos logs de startup
- Verificar se Caffeine está no classpath: `mvn dependency:tree | grep caffeine`

**Evidência a coletar:**

- [ ] Output do terminal mostrando as 5 tentativas 401 + 6ª tentativa 429
- [ ] Body JSON completo da resposta 429

---

## 7. Provenientes de Testes de Validação de Qualidade

> Nenhum erro encontrado durante a execução da Frente 1 que não fosse pré-existente.

| Task | Mensagem exata | Suspeita | Proposta solução |
|:-----|:---------------|:---------|:----------------:|
| — | `SubscriptionServiceTest > shouldCreateWithLockedPrice: IllegalStateException: TenantContext não inicializado` | Erro pré-existente (documentado no SPRINT-5-EXECUTION-REPORT-Frente-0.md). Teste não configura TenantContext antes de chamar SubscriptionService.create() | Corrigir na Frente 3: adicionar TenantContext.set(...) no @BeforeEach do teste |

---

## 8. Provenientes de Code Review

> Code review da Frente 1 será executada na Fase 7 do PROMPT-EXECUTE-SPRINT-TASKS.md.

| Task | Mensagem exata | Suspeita | Proposta solução | Skills |
|:-----|:---------------|:---------|:-----------------|-------:|
| — | (a preencher após execução da Fase 7) | — | — | — |

---

## Rodapé

🤖 *Documento gerado em 2026-07-23 como parte da Fase 3 do PROMPT-EXECUTE-SPRINT-TASKS.md. Stack detectada: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Caffeine 3.2.4 + JUnit 5 + Mockito + Testcontainers. Skills utilizados: 131-java-testing-unit-testing, 132-java-testing-integration-testing, 133-java-testing-acceptance-tests.*
