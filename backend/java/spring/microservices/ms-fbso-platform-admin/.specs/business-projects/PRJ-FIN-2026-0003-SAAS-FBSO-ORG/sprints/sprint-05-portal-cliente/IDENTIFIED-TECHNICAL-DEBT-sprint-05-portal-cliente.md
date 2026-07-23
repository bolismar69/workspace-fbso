# IDENTIFIED-TECHNICAL-DEBT-sprint-05-portal-cliente

- **Sprint alvo:** 5 de 7 — sprint-05-portal-cliente
- **Data da análise:** 2026-07-17
- **Skills executadas:** code-reviewer, caveman-review, superpowers:brainstorming, ponytail-review, ponytail-debt, code-review (CodeRabbit), jscpd+dry-refactoring, tech-debt, security-review
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Keycloak (OAuth2 OIDC) + Caffeine Cache + Flyway 10.22.0 + Testcontainers 1.21.4 + REST Assured 5.5.7
- **Total de achados:** 42 (🔴 14 críticos/bloqueantes, 🟡 16 riscos, 🔵 12 nits)
- **Impeditivos para iniciar a sprint:** 6 SIM — 6 débitos devem ser resolvidos ANTES de iniciar a implementação

---

## Resumo Executivo

A Sprint 5 (Portal do Cliente e Onboarding) apresenta **6 débitos bloqueantes** que impedem o início imediato da implementação: (1) o `docker-compose.yml` com Keycloak + PostgreSQL + MailHog não existe, inviabilizando desenvolvimento local das 12 tarefas; (2) Flyway está 2 majors atrás (10.22.0 → 12.11.0) — a migração deve ocorrer antes de criar novas migrations V007+ para a sprint; (3) PostgreSQL driver tem CVE ativa (CVSS 7.5) que precisa de bump; (4) o `JwtAuthenticationFilter` atual não extrai claims `modules[]` e `business_unit_ids[]`, bloqueando T-065 e T-066; (5) a documentação da sprint contém inconsistências numéricas (21 vs 28 cenários) e está em estado pré-review; (6) `TenantContext` tem campos `businessUnitIds` e `modules` populados desde a Sprint 4 mas nunca consumidos — a Sprint 5 depende deles.

Há ainda **16 débitos de risco** (🟡) e **12 nits** (🔵) catalogados, além de **17 débitos pendentes de sprints anteriores** (backlog) que permanecem não resolvidos. Desses, 2 débitos do backlog (DT-023 — keyset pagination, DT-045 — Flyway migration) estão explicitamente designados para a Sprint 5 e devem ser incluídos no planejamento.

**Recomendação:** Resolver os 6 bloqueantes na Frente 0 (pré-sprint, ~2d) antes de iniciar as tarefas T-057 a T-068. Incluir DT-023 e DT-045 como tarefas da sprint. Os 11 nits DT-083 a DT-093 podem ser postergados para a Sprint 6 se a capacidade não permitir.

---

## Backlog de Débitos Técnicos (Sprints Anteriores)

Débitos técnicos identificados em sprints anteriores que **permanecem não resolvidos** e são candidatos a tratamento na sprint atual ou futuras.

| DT-XXX | Sprint Origem | Descrição | Severidade | Bloqueante? | Status | Resolução (do doc original ou revisada) |
|:---|:---|:---|:---:|:---:|:---|:---|
| DT-023 | Sprint 3 | Migrar paginação offset→keyset no BaseRepository — degradação >10k registros | 🟡 | NÃO | Pendente (target Sprint 5) | Implementar `findAllKeyset(keyset, pageSize)` usando `WHERE id > :lastId ORDER BY id LIMIT :size`. Estimar 3h |
| DT-045 | Sprint 3 | Migrar Flyway 10.22.0→12.11.0 (2 majors atrás) | 🟡 | NÃO → **SIM (bloqueia V007+)** | Pendente (target Sprint 5) | Bump `flyway.version` no pom.xml. Atualizar `flyway-database-postgresql`. Rodar `mvn flyway:migrate` para validar. Estimar 1h |
| DT-031 | Sprint 3 | Reduzir Checkstyle maxAllowedViolations 300→100→0 | 🔵 | NÃO | Pendente (Sprints 4-7) | Progressivo: Sprint 5 reduz para 200, Sprint 6 para 100, Sprint 7 para 0 |
| DT-034 | Sprint 3 | Remover Address.java se não usado até Sprint 6 | 🔵 | NÃO | Pendente (target Sprint 6) | ↗ ver DT-034 em IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md |
| DT-044 | Sprint 3 | Criar logback-spring.xml com appender JSON | 🔵 | NÃO | Pendente (target Sprint 7) | ↗ ver DT-044 em IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md |
| DT-068 | Sprint 4 | Atualizar PostgreSQL driver 42.7.10→42.7.11 (CVE-2026-42198, CVSS 7.5) | 🟡 | NÃO → **SIM (CVE ativa)** | Pendente (Sprint 4 recomendados) | Bump `postgresql.version` no pom.xml. CVE de negação de serviço. Estimar 0.5h |
| DT-069 | Sprint 4 | Migrar Flyway 10.22.0→12.11.0 (idêntico ao DT-045 — consolidar) | 🟡 | NÃO | Pendente (Sprint 4 recomendados) | Consolidado com DT-045 |
| DT-070 | Sprint 4 | Corrigir AuditAspect.parseEntityId() — retorna UUID aleatório para entrada inválida | 🟡 | NÃO | Pendente (Sprint 4 recomendados) | Adicionar validação UUID.fromString() com try-catch + log WARN. Estimar 1h |
| DT-071 | Sprint 4 | Adicionar @WebMvcTest para 4 controllers (PlanController, TenantController, SubscriptionController, AuditController) | 🟡 | NÃO | Pendente (Sprint 4 recomendados) | Criar testes WebMvcTest com mock de serviço. Estimar 4h |
| DT-072 | Sprint 4 | Corrigir RLSIsolationTest — remover @Disabled e testar isolamento real | 🟡 | NÃO | Pendente (Sprint 4 recomendados) | Refatorar para usar Testcontainers com SET LOCAL app.current_tenant_id. Estimar 2h |
| DT-073 | Sprint 4 | Corrigir BaseRepositoryTest.save/update — remover @Disabled | 🟡 | NÃO | Pendente (Sprint 4 recomendados) | Implementar testes unitários para save(T) e update(T). Estimar 1.5h |
| DT-074 | Sprint 4 | Adicionar PERMISSION, ROLE ao AuditAspect.resolveTableName() | 🟡 | NÃO | Pendente (Sprint 4 recomendados) | Adicionar cases no switch. Estimar 0.5h |
| DT-075 | Sprint 4 | Adicionar handler TenantIsolationException→403 no GlobalExceptionHandler | 🟡 | NÃO | Pendente (Sprint 4 recomendados) | Substituir handler genérico SecurityException por handler específico. Estimar 0.5h |
| DT-076 | Sprint 4 | Consolidar dupla decodificação JWT via Converter\<Jwt, AbstractAuthenticationToken\> | 🟡 | NÃO | Pendente (Sprint 4 recomendados) | Implementar JwtAuthenticationConverter customizado. Impacta T-065. Estimar 2h |
| DT-077 | Sprint 4 | Reescrever RbacAspectTest com @SpringBootTest (hoje instanciado sem Spring) | 🟡 | NÃO | Pendente (Sprint 4 recomendados) | Refatorar para teste de integração com contexto Spring real. Estimar 2h |
| DT-078 | Sprint 4 | Adicionar SELECT...FOR UPDATE em PlanService.deactivate() — race condition TOCTOU | 🟡 | NÃO | Pendente (Sprint 4 recomendados) | Adicionar lock pessimista na query de desativação. Estimar 1h |
| DT-083 | Sprint 4 | Dead code: Address + 5 enums (~211 linhas) — remover se não consumidos | 🔵 | NÃO | Pendente (target Sprint 5+) | ↗ ver DT-083 em IDENTIFIED-TECHNICAL-DEBT-sprint-04-rbac.md |
| DT-084 | Sprint 4 | EmailService interface YAGNI — 1 implementação, 88 linhas | 🔵 | NÃO | Pendente (target Sprint 5+) | ↗ ver DT-084 em IDENTIFIED-TECHNICAL-DEBT-sprint-04-rbac.md |
| DT-085 | Sprint 4 | TenantContext com campos BU/modules populados mas nunca consumidos | 🔵 | NÃO → **SIM (Sprint 5 depende)** | Pendente (target Sprint 5+) | **Revisado:** Sprint 5 (T-065) implementa consumo. Deixar de ser nit — agora é pré-requisito funcional |
| DT-086 | Sprint 4 | RowMappers duplicam 6 campos de auditoria (x4) — extrair helper | 🔵 | NÃO | Pendente (target Sprint 5+) | Extrair AuditFieldsRowMapper ou método static. Estimar 1.5h |
| DT-087 | Sprint 4 | Construção SQL duplicada (~70 linhas) entre repositories | 🔵 | NÃO | Pendente (target Sprint 5+) | Extrair SqlBuilder utility. Estimar 2h |
| DT-088 | Sprint 4 | Validação de plano ativo duplicada (12 linhas) — PlanService + SubscriptionService | 🔵 | NÃO | Pendente (target Sprint 5+) | Extrair método `validateActivePlan()` compartilhado. Estimar 0.5h |
| DT-089 | Sprint 4 | ObjectMapper não gerenciado pelo Spring no AuditAspect — `new ObjectMapper()` manual | 🔵 | NÃO | Pendente (target Sprint 5+) | Injete @Autowired ObjectMapper do contexto Spring. Estimar 0.5h |
| DT-090 | Sprint 4 | OffsetDateTime.now() sem ZoneOffset.UTC no BaseEntity — risco de fuso horário inconsistente | 🔵 | NÃO | Pendente (target Sprint 5+) | Substituir por `OffsetDateTime.now(ZoneOffset.UTC)`. Estimar 0.5h |
| DT-091 | Sprint 4 | Endpoint stub `POST /api/v1/tenants/{id}/resend-invite` — implementação vazia | 🔵 | NÃO | Pendente (target Sprint 5+) | Implementar reenvio de email ou remover endpoint. Estimar 1h |
| DT-092 | Sprint 4 | springdoc-openapi 2.8.8→2.8.16 (8 minors atrás) | 🔵 | NÃO | Pendente (target Sprint 5+) | Bump versão no pom.xml. Verificar breaking changes. Estimar 0.5h |
| DT-093 | Sprint 4 | CORS origins hardcoded em SecurityConfig — `allowedOrigins("http://localhost:3000")` | 🔵 | NÃO | Pendente (target Sprint 5+) | Externalizar para `application.yml` via `cors.allowed-origins`. Estimar 0.5h |
| DT-094 | Sprint 4 | README.md vazio ("TODO") | 🔵 | NÃO | Pendente (target Sprint 7) | Preencher com instruções de setup, build, deploy. Estimar 2h |

**Total em backlog:** 28 débitos pendentes de sprints anteriores (17 Sprint 4 recomendados + 9 Sprint 4+ postergados + 2 Sprint 3 backlog não migrados).

> **Nota sobre DT-085:** O campo `TenantContext.businessUnitIds` e `TenantContext.modules` foram populados na Sprint 4 mas nunca consumidos. A Sprint 5 (T-065, T-066) implementa o consumo desses campos no `JwtAuthenticationFilter` e `GET /auth/me`. Este débito deixa de ser "nit" e passa a ser **pré-requisito funcional** para a Sprint 5.

---

## Matriz de Débitos Técnicos

> **Esta matriz consolida TODOS os débitos — tanto os novos (descobertos nesta sprint) quanto os do backlog (sprints anteriores).**
>
> **Legenda das colunas:**
> - **ID:** DT-XXX (Débito Técnico, numeração sequencial e IMUTÁVEL)
> - **Sprint Origem:** Em qual sprint o débito foi identificado
> - **Severidade:** 🔴 Crítico (bloqueante) | 🟡 Risco (deve ser tratado) | 🔵 Nit (desejável)
> - **Skill:** Qual skill identificou (CR=caveman-review, PONY=ponytail-review, PDBT=ponytail-debt, ARCH=brainstorming, CREV=code-reviewer, CODE=code-review, JSCPD=jscpd+dry-refactoring, DEBT=tech-debt, SEC=security-review). Débitos de backlog: `BACKLOG`.
> - **Complexidade:** H (Alta, >4h) | M (Média, 1-4h) | L (Baixa, <1h)
> - **Bloqueante?:** SIM (impede o início/incremento da sprint) | NÃO (pode ser tratado depois)
> - **Efeito se não tratado:** O que acontece se este débito for ignorado

| ID | Sprint Origem | Arquivo/Artefato | Achado | Severidade | Skill | Complexidade | Bloqueante? | Efeito se não tratado |
|:---|:---|:---|:---|:---:|:---:|:---:|:---:|:---|
| DT-023 | Sprint 3 | `BaseRepository.java` | Paginação offset degrada com >10k registros — migrar para keyset | 🟡 | BACKLOG | M | NÃO | Dashboard lento com muitos tenants; timeout em produção |
| DT-045 | Sprint 3 | `pom.xml` | Flyway 10.22.0→12.11.0 (2 majors atrás). Bloqueia criação de V007+ | 🟡 | BACKLOG | L | SIM | Novas migrations não seguem padrão atualizado; risco de incompatibilidade futura |
| DT-068 | Sprint 4 | `pom.xml` | PostgreSQL driver 42.7.10→42.7.11 (CVE-2026-42198, CVSS 7.5 DoS) | 🟡 | BACKLOG | L | SIM | Vulnerabilidade de negação de serviço ativa em produção |
| DT-095 | Sprint 5 | `docker-compose.yml` (inexistente) | Sem docker-compose com Keycloak 26 + PostgreSQL 17 + MailHog — desenvolvimento local inviável | 🔴 | ARCH | M | SIM | 12 tarefas da sprint não podem ser desenvolvidas/testadas localmente |
| DT-096 | Sprint 5 | `JwtAuthenticationFilter.java` | Filter não extrai claims `modules[]` e `business_unit_ids[]` do JWT — T-065 e T-066 bloqueadas | 🔴 | CREV | M | SIM | App Switcher (F04-04) e GET /auth/me não funcionam sem claims |
| DT-097 | Sprint 5 | `SPRINT-TEST-SUITE.md:L3` | Cabeçalho declara 21 cenários, corpo lista 28 — inconsistência numérica (diferença de 7) | 🔴 | CR | L | SIM | Métricas de teste não confiáveis; planning incorreto |
| DT-098 | Sprint 5 | `TenantContext.java` | Campos `businessUnitIds` e `modules` populados mas getters nunca chamados — preenchimento sem consumo | 🔴 | PONY | L | SIM | T-065 e T-066 implementam consumo; sem isso, claims do JWT são descartadas |
| DT-099 | Sprint 5 | `pom.xml` | spring-boot-starter-oauth2-resource-server presente mas sem spring-boot-starter-oauth2-client — Authorization Code Flow requer ambos | 🔴 | ARCH | L | SIM | T-057 (Keycloak realm) e T-058 (login) não funcionam sem OAuth2 Client |
| DT-100 | Sprint 5 | `application.yml` | Configuração OAuth2 apenas Resource Server (JWKS). Falta Provider Details (issuer-uri completo + client-id/client-secret para Auth Code Flow) | 🔴 | SEC | M | SIM | Fluxo de login Authorization Code Flow não configurado |
| DT-101 | Sprint 5 | `SPRINT-CARD.md:L73-77` | Riscos de mitigação desatualizados — DT-042 (docker-compose) e DT-068 (driver CVE) ainda pendentes, mas mitigação lista "Keycloak container no Docker Compose" como se existisse | 🟡 | CR | L | NÃO | Planejamento de risco impreciso; falsa sensação de segurança |
| DT-102 | Sprint 5 | `JwtAuthenticationFilter.java:L42-68` | Dupla decodificação JWT por requisição (BearerTokenAuthenticationFilter + custom filter). DT-076 pendente desde Sprint 4 | 🟡 | DEBT | M | NÃO | Performance degradada; 2x validações de assinatura por request |
| DT-103 | Sprint 5 | `src/main/java/com/fbso/platform/admin/service/` (inexistente) | Nenhum service de Onboarding existe — T-060 precisa criar OnboardingService do zero. Sem reuso de padrões existentes | 🟡 | ARCH | H | NÃO | Curva de aprendizado; risco de inconsistência com padrões do projeto |
| DT-104 | Sprint 5 | `src/main/java/com/fbso/platform/admin/controller/` (inexistente) | Nenhum controller de Auth existe — T-058, T-061, T-064, T-066 são 4 novos controllers. Padrão de nomenclatura inconsistente (AuthController vs OnboardingController vs DashboardClientController) | 🟡 | CREV | M | NÃO | Inconsistência de API; 4 controllers seguirem 4 padrões diferentes |
| DT-105 | Sprint 5 | `DashboardController.java` (existente) | Dashboard admin e dashboard cliente compartilharão prefixo `/api/v1/dashboard/` — risco de colisão de rotas e confusão de responsabilidade | 🟡 | ARCH | M | NÃO | Rotas ambíguas; admin acessa endpoint de cliente sem querer |
| DT-106 | Sprint 5 | `SPRINT-TEST-SUITE.md` | 28 cenários definidos mas 0% executados. Suite não cobre: timeout de sessão (60min), validação complexidade senha (RN13-01), passo 3 do onboarding, segurança do dashboard cliente (F04-03) | 🟡 | DEBT | H | NÃO | Cobertura insuficiente; bugs escapam para produção |
| DT-107 | Sprint 5 | `BusinessUnit.java` (existente) | Entidade BusinessUnit existe mas sem campo `is_matrix` (boolean) — T-062 precisa distinguir BU Matriz de BU Filial | 🟡 | CREV | L | NÃO | Lógica de matriz/filial implementada via parent_id=NULL (frágil); sem flag explícita |
| DT-108 | Sprint 5 | `Tenant.java` (existente) | Enum TenantStatus tem PENDING_ONBOARDING e ACTIVE mas transições não documentadas via máquina de estados — T-060 implementa `PENDING_ONBOARDING→ACTIVE` | 🟡 | ARCH | M | NÃO | Transições de status inconsistentes; tenant pode pular onboarding |
| DT-109 | Sprint 5 | `pom.xml` | Dependência `spring-boot-starter-mail` existe desde DT-007 (Sprint 3) — verificar se MailHog/SMTP está funcional para T-058 (forgot-password) | 🟡 | SEC | L | NÃO | Fluxo de recuperação de senha não testável localmente |
| DT-110 | Sprint 5 | `SPRINT-CARD.md:L59` | Task T-059 define rate limiting "via @Aspect ou filter + Caffeine" — decisão de design não tomada. @Aspect adiciona complexidade AOP; filter é mais simples mas menos reutilizável | 🟡 | ARCH | M | NÃO | Implementação incorreta; retrabalho se escolha errada |
| DT-111 | Sprint 5 | `SPRINT-CARD.md:L73` | Risco "Rate limiting com Caffeine não funcionar em cluster (stateless)" — mitigação "Fase 1: migrar para Redis" não tem critério de trigger definido | 🟡 | ARCH | L | NÃO | Rate limiting ineficaz em multi-instância; sem critério claro para upgrade |
| DT-112 | Sprint 5 | `SPECS.md` header | Documento declara "Proximo: Sprint 4 Frentes 5-5b" mas Sprint 5 é a atual — referência desatualizada | 🔵 | CR | L | NÃO | Confusão para novos devs; documento não reflete sprint atual |
| DT-113 | Sprint 5 | `TASKS.md` header | Documento declara "88/136 tarefas (65%)" mas não reflete pendências da Sprint 4 (M4 11 tasks + recomendados 17 tasks = 28 pendentes) | 🔵 | CR | L | NÃO | Status report impreciso; stakeholders subestimam trabalho restante |
| DT-114 | Sprint 5 | `RateLimiter.java` (inexistente) | Checkstyle com 711 violações e `failOnViolation=false` (DT-031) — novo código da Sprint 5 deve seguir padrão zero-warning | 🔵 | CREV | H | NÃO | Dívida de estilo acumulada; difícil impor qualidade em código novo |
| DT-115 | Sprint 5 | `SecurityConfig.java` | SecurityConfig precisará de refactor significativo: adicionar OAuth2 Client (login), manter Resource Server (API), configurar CORS para frontend cliente | 🟡 | ARCH | M | NÃO | Configuração de segurança monolítica; difícil manutenção |
| DT-116 | Sprint 5 | `application.yml` | Propriedade `spring.security.oauth2.resourceserver.jwt.issuer-uri` precisa ser complementada com `spring.security.oauth2.client.registration.keycloak.*` e `spring.security.oauth2.client.provider.keycloak.*` | 🟡 | SEC | M | NÃO | Configuração incompleta; login Authorization Code Flow não funciona |
| DT-117 | Sprint 5 | `BaseEntity.java` | `OffsetDateTime.now()` sem ZoneOffset.UTC — todos os timestamps de auditoria podem divergir entre instâncias em timezones diferentes | 🔵 | PONY | L | NÃO | Timestamps inconsistentes entre ambientes (DEV vs PRD em timezones diferentes) |
| DT-118 | Sprint 5 | `src/main/resources/db/migration/` | Próximas migrations (V007+) para Sprint 5: tabelas de onboarding? Campos novos em tenant? Schema não planejado | 🟡 | DEBT | M | NÃO | Migrations ad-hoc; inconsistência de schema entre ambientes |
| DT-119 | Sprint 5 | `SPRINT-REVIEW.md` | Documento completamente vazio — sem checkmarks, sem bloqueios, sem resultados. Review não iniciada | 🔵 | CR | L | NÃO | Sem baseline de review; difícil medir progresso da sprint |
| DT-120 | Sprint 5 | Código existente | 5 RowMappers duplicam lógica de extração de campos de auditoria (6 campos × 5 mappers = 30 linhas duplicadas) — DT-086 pendente | 🔵 | JSCPD | L | NÃO | Código duplicado; mudança em campos de auditoria requer 5 edições |
| DT-121 | Sprint 5 | `GlobalExceptionHandler.java` | Handler de `SecurityException` captura `AccessDeniedException` mas não `AuthenticationException` — endpoints auth precisam de 401 claro (não 500) | 🟡 | CODE | L | NÃO | Erro 500 em vez de 401 para token inválido/expirado |
| DT-122 | Sprint 5 | `AuditAspect.java` | `new ObjectMapper()` manual no AuditAspect — não usa ObjectMapper do Spring (com módulos Jackson registrados). DT-089 pendente | 🔵 | PONY | L | NÃO | Serialização de auditoria pode divergir do resto da aplicação |
| DT-123 | Sprint 5 | `SPRINT-TEST-SUITE.md` | F04-03 (Dashboard Cliente) sem nenhum cenário de segurança. Isolamento entre tenants do dashboard cliente não validado | 🟡 | SEC | M | NÃO | Cliente de tenant A pode ver dados do tenant B |
| DT-124 | Sprint 5 | `OnboardingService` (a criar) | Máquina de estados do onboarding (4 passos, ordem obrigatória, retomável) sem especificação formal — risco de edge cases não tratados (ex: retomar após 30 dias) | 🟡 | DEBT | H | NÃO | Estados inconsistentes; tenant preso em PENDING_ONBOARDING |
| DT-125 | Sprint 5 | `pom.xml` | JaCoCo min.line.coverage=0.72 (72%) abaixo da meta de 80% do TEST_PLAN. Muitos testes @Disabled (DT-072, DT-073) puxam cobertura para baixo | 🔵 | DEBT | M | NÃO | Métrica de qualidade falsa; cobertura real abaixo do reportado |
| DT-126 | Sprint 5 | `application-dev.yml` | Perfil dev com `logging.level.com.fbso=DEBUG` — logs de SQL com dados sensíveis em desenvolvimento (DT-033 nunca resolvido) | 🔵 | SEC | L | NÃO | Vazamento de dados pessoais em logs de desenvolvimento |

> **Ordenação:** Débitos bloqueantes primeiro (🔴), depois por sprint origem (mais antigos primeiro — risco de apodrecimento), depois por severidade.

---

## Achados por Skill

### code-reviewer (8 achados)

**Regras violadas e code smells no código existente:**

1. **DT-096 — SOLID-SRP, SMELL-missing-claims:** `JwtAuthenticationFilter.attemptAuthentication()` extrai `tenant_id`, `roles`, `user_id` do JWT mas ignora `modules[]` e `business_unit_ids[]`. O filtro tem uma responsabilidade clara (extrair claims do token), mas está incompleto para a Sprint 5.
   - **Arquivo:** `security/JwtAuthenticationFilter.java:L42-L68`
   - **Severidade:** 🔴 Critical
   - **Ação:** Adicionar extração `modules[]` como `List<String>` e `business_unit_ids[]` como `List<UUID>` do claim `resource_access.fbso-platform`

2. **DT-104 — SMELL-naming-convention:** Quatro novos controllers para a Sprint 5 sem convenção de nomenclatura estabelecida: `AuthController` (autenticação), `OnboardingController` (onboarding), `DashboardClientController` (dashboard cliente), endpoint `GET /auth/me` (app switcher).
   - **Arquivo:** `controller/` (4 novos arquivos a criar)
   - **Severidade:** 🟡 Warning
   - **Ação:** Definir convenção: endpoints de cliente prefixo `/client/` ou controllers sufixo `ClientController`?

3. **DT-114 — SMELL-checkstyle-debt:** 711 violações Checkstyle acumuladas com `failOnViolation=false`. Novo código da Sprint 5 deve seguir padrão zero-warning para não agravar.
   - **Arquivo:** `pom.xml` (configuração checkstyle)
   - **Severidade:** 🔵 Info
   - **Ação:** Configurar `failOnViolation=true` apenas para novos arquivos da Sprint 5 via `suppressions.xml` seletivo

4. **DT-107 — SMELL-missing-field:** `BusinessUnit.java` não tem flag `is_matrix` (boolean). A lógica de "primeira BU = Matriz" depende de `parent_id=NULL`, o que é frágil — qualquer BU sem parent pode ser confundida com matriz.
   - **Arquivo:** `entity/BusinessUnit.java`
   - **Severidade:** 🟡 Warning
   - **Ação:** Adicionar campo `is_matrix BOOLEAN NOT NULL DEFAULT false` via migration V007

5. **SOLID-DIP genérico:** `RbacAspect` e `AuditAspect` dependem diretamente de `JdbcTemplate` em vez de uma abstração — baixa testabilidade. Já catalogado como DT-077.
   - **Status:** Já coberto por DT-077

6. **SMELL-long-method potencial:** `SecurityConfig.java` atual (filterChain, CORS, OAuth2) crescerá significativamente com OAuth2 Client + Resource Server. Risco de método `securityFilterChain` com >50 linhas.
   - **Arquivo:** `config/SecurityConfig.java`
   - **Severidade:** 🟡 Warning (DT-115)
   - **Ação:** Extrair `oauth2ClientConfig()` e `oauth2ResourceServerConfig()` como métodos separados

7. **SMELL-missing-tests:** `AuditAspect` tem 0% de cobertura de testes unitários. DT-028 aceito como risco na Sprint 3, mas com a Sprint 5 adicionando novos fluxos de auditoria (onboarding, login), o risco aumenta.
   - **Arquivo:** `security/aspect/AuditAspect.java`
   - **Severidade:** 🟡 Warning
   - **Ação:** Criar `AuditAspectTest` com mock de JdbcTemplate (1-2h)

8. **SMELL-null-handling:** `JwtUtils.extractClaim()` retorna null sem `Optional` — consumidores não são forçados a tratar ausência. DT-047 aceito como risco mas propaga-se para novas claims da Sprint 5.
   - **Arquivo:** `utils/JwtUtils.java`
   - **Severidade:** 🔵 Info
   - **Ação:** Refatorar para retornar `Optional<String>` com adaptação gradual

### caveman-review (9 achados)

**Inconsistências entre documentação e código real:**

1. **DT-097 — Contagem divergente:** `SPRINT-TEST-SUITE.md` cabeçalho declara **21 cenários**, quadro resumo lista **28** (9+10+5+4=28). Diferença de 7 cenários. O `SPRINT-CARD.md` declara **21 cenários** na tabela de métricas (§Métricas). Ambos divergem.
   - **Severidade:** 🔴 Bug — número errado propaga para planning e reports
   - **Correção:** Atualizar SPRINT-CARD.md métricas (21→28) e SPRINT-TEST-SUITE.md cabeçalho (21→28)

2. **DT-101 — Mitigação fantasma:** `SPRINT-CARD.md` (§Riscos e Bloqueadores) lista mitigação "Keycloak container no Docker Compose" como se `docker-compose.yml` existisse — mas DT-042 (docker-compose) está pendente desde Sprint 3.
   - **Severidade:** 🟡 Risk — falsa sensação de segurança
   - **Correção:** Atualizar mitigação para "Criar docker-compose.yml (DT-042/DT-095) antes de iniciar a sprint"

3. **DT-112 — Status desatualizado:** `SPECS.md` header declara "Proximo: Sprint 4 Frentes 5-5b" — mas Sprint 4 já concluiu e Sprint 5 é a atual.
   - **Severidade:** 🔵 Nit
   - **Correção:** Atualizar header para "Proximo: Sprint 5 Portal do Cliente (12 tarefas)"

4. **DT-113 — Progresso impreciso:** `TASKS.md` declara "88/136 tarefas (65%)" mas Sprint 4 tem 28 tarefas pendentes (M4 11 + recomendados 17). O número 88 inclui apenas concluídas; as pendentes da Sprint 4 não estão refletidas no status.
   - **Severidade:** 🔵 Nit
   - **Correção:** Recalcular: 88 concluídas + 28 pendentes Sprint 4 + 12 pendentes Sprint 5 = reporting mais preciso

5. **DT-119 — Review vazia:** `SPRINT-REVIEW.md` completamente sem preenchimento — 0 checkmarks, 0 bloqueios, 0 resultados. Documento serve como template mas não reflete estado real.
   - **Severidade:** 🔵 Nit
   - **Correção:** Preencher após revisão; ou marcar explicitamente como "Pré-sprint — revisão pendente"

6. **SPECS.md endpoint count:** Documento lista 37 endpoints REST mas Sprint 5 adiciona 4 novos (`/auth/*`, `/onboarding/*`, `/dashboard/client/*`). Total subirá para 41. SPECS.md deve ser atualizado.
   - **Severidade:** 🔵 Nit
   - **Correção:** Atualizar §4.1 com novos endpoints após implementação

7. **ARCHITECTURE.md diagrama C4:** Diagrama L3 (Component) referencia `OnboardingController` e `DashboardClientController` como existentes, mas eles ainda não foram implementados. Diagrama é aspiracional, não descritivo.
   - **Severidade:** 🔵 Nit
   - **Correção:** Adicionar nota "Planejado — Sprint 5" nos componentes ainda não implementados

8. **TEST_PLAN.md status:** Seção §3.13-3.16 (F04-01 a F04-04) lista 27 cenários como "Planejado" mas a suite da sprint (SPRINT-TEST-SUITE.md) lista 28 — divergência de 1 cenário entre documentos.
   - **Severidade:** 🔵 Nit
   - **Correção:** Alinhar contagem entre TEST_PLAN.md e SPRINT-TEST-SUITE.md

9. **PRD.md §4.3 milestones:** Datas dos milestones M4 (15/09) e M5 (30/09) são anteriores à data atual (17/07). Planejamento temporal desatualizado.
   - **Severidade:** 🔵 Nit
   - **Correção:** Revisar cronograma ou adicionar nota "datas originais — replanejamento em andamento"

### superpowers:brainstorming (6 achados)

**Decisões arquiteturais questionáveis e gaps de planejamento:**

1. **DT-110 — Rate limiting strategy não decidida:** `SPRINT-CARD.md:T-059` propõe "@Aspect ou filter + Caffeine" como alternativas em aberto. Isso é uma decisão de design que deveria estar tomada ANTES da sprint.
   - **Premissa original:** Rate limiting é simples, qualquer abordagem serve
   - **Realidade:** @Aspect adiciona complexidade AOP e requer Pointcut; filter é procedural mas mais testável
   - **Recomendação:** Usar `Filter` (não `@Aspect`) — mais simples, testável, e alinhado com `JwtAuthenticationFilter`

2. **DT-111 — Sem trigger para Redis:** Risco "Caffeine não funciona em cluster" documentado mas mitigação "Fase 1: migrar para Redis" não define critério objetivo (ex: "quando >1 instância em produção").
   - **Recomendação:** Definir trigger: "quando `INSTANCE_COUNT > 1` no ambiente de produção, migrar Caffeine→Redis em até 1 sprint"

3. **DT-105 — Colisão de rotas dashboard:** `DashboardController` (admin) usa prefixo `/api/v1/dashboard/admin/`. `DashboardClientController` (cliente) usará `/api/v1/dashboard/client/`. Rotas são distintas mas compartilham prefixo `/api/v1/dashboard/` — risco de middleware de segurança tratar ambas igualmente.
   - **Recomendação:** Separar completamente: `/api/v1/admin/dashboard/` vs `/api/v1/client/dashboard/`

4. **DT-124 — Onboarding sem especificação formal de estados:** A máquina de estados do onboarding (4 passos, ordem obrigatória, retomável) não tem diagrama de estados ou tabela de transições documentada. Isso é crítico porque:
   - Passo 2 falha mas passo 1 foi salvo? (consistência)
   - Usuário abandona por 30 dias e retorna? (timeout)
   - Tenant tem dados parciais? (rollback)
   - **Recomendação:** Criar diagrama de estados antes de implementar `OnboardingService`

5. **DT-118 — Schema de migrations não planejado:** Sprint 5 precisará de novas migrations (V007+). O que elas conterão?
   - Tabela `onboarding_step`? (tracking de progresso)
   - Campo `onboarding_status` em `tenant`?
   - Campo `is_matrix` em `business_unit`?
   - Novas tabelas para `rate_limit_attempt`?
   - **Recomendação:** Planejar schema antes de codificar

6. **DT-099 — OAuth2 Client vs Resource Server:** `pom.xml` tem `spring-boot-starter-oauth2-resource-server` (valida JWT) mas NÃO tem `spring-boot-starter-oauth2-client` (Authorization Code Flow). Para T-057 (Keycloak realm + login), ambos são necessários.
   - **Premissa original:** Backend só valida JWT, não inicia fluxo OAuth2
   - **Realidade Sprint 5:** Backend precisa iniciar Authorization Code Flow (login redirect para Keycloak)
   - **Recomendação:** Adicionar dependência `spring-boot-starter-oauth2-client`

### ponytail-review (5 achados)

**Complexidade e código que pode ser simplificado:**

1. **DT-098 — TenantContext com campos não consumidos:** `TenantContext` mantém `businessUnitIds` (List\<UUID\>) e `modules` (List\<String\>) via ThreadLocal mas nenhum código os lê. Os getters existem mas nunca são chamados fora de testes.
   - **Tag:** `yagni` — se Sprint 5 não consumir, remover
   - **Net lines:** -45 lines possible (remover setters/getters não usados + limpeza)

2. **DT-122 — ObjectMapper manual no AuditAspect:** `new ObjectMapper()` em vez de injetar o bean gerenciado pelo Spring (que já tem módulos Jackson registrados — JavaTimeModule, etc.).
   - **Tag:** `stdlib` — usar bean existente
   - **Net lines:** -2 lines (remover `new ObjectMapper()`, adicionar `@Autowired`)

3. **DT-117 — OffsetDateTime sem UTC:** `BaseEntity.java` usa `OffsetDateTime.now()` (fuso horário da JVM) em vez de `OffsetDateTime.now(ZoneOffset.UTC)`. Em ambientes com timezones diferentes, timestamps de auditoria divergem.
   - **Tag:** `bug-risk`
   - **Net lines:** 0 (correção de 1 linha)

4. **DT-086/DT-120 — Duplicação de campos de auditoria:** 5 RowMappers (`TenantRowMapper`, `PlanRowMapper`, `SubscriptionRowMapper`, `AuditEntryRowMapper`, `UserRowMapper`, `UserPermissionRowMapper`) repetem a mesma extração de 6 campos de auditoria.
   - **Tag:** `shrink`
   - **Net lines:** -30 lines possible (extrair helper)

5. **DT-084 — EmailService interface YAGNI:** Interface `EmailService` com 1 implementação concreta (`EmailServiceImpl`). A abstração não tem propósito — não há múltiplas implementações nem troca em runtime.
   - **Tag:** `yagni`
   - **Net lines:** -88 lines possible (remover interface)
   - **Cuidado:** Se testes mockam a interface, mantê-la

### ponytail-debt (1 achado)

**Comentários `ponytail:` no código-fonte:**

```
Grep: grep -rnE '(#|//|--) ?ponytail:' src/ --exclude-dir=target

Resultado: No ponytail: debt. Clean ledger.
```

**0 markers encontrados.** O projeto não utiliza marcadores `ponytail:` para sinalizar atalhos intencionais. Isso é positivo — significa que não há dívida técnica "assumida" não documentada. No entanto, os débitos catalogados nas seções anteriores indicam que atalhos foram tomados sem o marcador (ex: `@Disabled` em testes, `failOnViolation=false` no Checkstyle, `new ObjectMapper()` manual). 

**Recomendação:** Adotar o padrão `// ponytail: <ceiling> | trigger: <condição>` para novos atalhos intencionais na Sprint 5.

### code-review (5 achados)

**Bugs, vulnerabilidades e anti-padrões (revisão manual — CodeRabbit CLI não disponível):**

1. **DT-121 — AuthenticationException não tratada:** `GlobalExceptionHandler` tem handler para `AccessDeniedException` (403) mas não para `AuthenticationException` (401). Com OAuth2 Client na Sprint 5, erros de autenticação (token expirado, assinatura inválida) podem resultar em 500.
   - **Severidade:** 🟡 Warning
   - **Arquivo:** `exception/GlobalExceptionHandler.java`
   - **Ação:** Adicionar `@ExceptionHandler(AuthenticationException.class)` → 401 RFC 7807

2. **DT-062 — Isolamento tenant_id URL vs JWT:** `SubscriptionService` recebe `tenantId` da URL e do JWT mas não valida que são iguais. Corrigido na Sprint 4 (DT-062, Frente 0) — verificar se a correção persiste.
   - **Severidade:** ✅ Já resolvido (Sprint 4 Frente 0)
   - **Arquivo:** `service/SubscriptionService.java`
   - **Verificação:** Regressão — garantir que novos services (OnboardingService, DashboardClientService) também validem

3. **Rate limiting race condition:** Se implementado via Caffeine, `get(key, k -> new AtomicInteger(0))` tem race condition entre get e increment. Deve-se usar `asMap().computeIfAbsent()` atômico.
   - **Severidade:** 🟡 Warning (DT-110 relacionado)
   - **Arquivo:** `RateLimiter.java` (a criar — T-059)
   - **Ação:** Usar `AtomicInteger` + `synchronized` ou `ConcurrentHashMap.compute()`

4. **Onboarding inconsistência transacional:** Se passo 2 falha (erro ao criar BU) mas passo 1 já foi salvo, como fazer rollback? `@Transactional` no `OnboardingService` resolve para operações de banco, mas não para estado no Keycloak (ex: atualização de atributos do usuário).
   - **Severidade:** 🟡 Warning (DT-124 relacionado)
   - **Arquivo:** `OnboardingService.java` (a criar)
   - **Ação:** Implementar saga de compensação ou salvar estado do onboarding no banco (não no Keycloak)

5. **Senha em logs:** `application-dev.yml` com `logging.level.com.fbso=DEBUG` pode expor parâmetros de requisição em logs, incluindo senhas no fluxo de login. DT-033 aceito como risco, mas com a Sprint 5 adicionando endpoints de auth, o risco aumenta.
   - **Severidade:** 🔵 Info (DT-126)
   - **Arquivo:** `application-dev.yml`
   - **Ação:** Adicionar filter para mascarar campos `password`, `token`, `authorization` nos logs

### jscpd + dry-refactoring (4 achados)

**Duplicação de código (jscpd indisponível — análise manual de duplicação estrutural):**

1. **DT-120/DT-086 — Duplicação em RowMappers (6 campos × 5 mappers):**
   - **Arquivos:** `TenantRowMapper.java`, `PlanRowMapper.java`, `SubscriptionRowMapper.java`, `AuditEntryRowMapper.java`, `UserRowMapper.java`, `UserPermissionRowMapper.java`
   - **Bloco duplicado:** 6 campos de auditoria (`created_dt`, `updated_dt`, `created_by`, `updated_by`, `deleted_dt`, `deleted_by`) extraídos com `rs.getTimestamp()` + `rs.getString()`
   - **Estratégia:** Extract helper method — `AuditFieldsHelper.extract(ResultSet rs)` ou `BaseRowMapper` abstrato
   - **Estimativa:** -30 linhas, 1.5h

2. **DT-087 — Construção SQL duplicada:**
   - **Arquivos:** `TenantRepository.java`, `PlanRepository.java`, `UserRepository.java`, `SubscriptionRepository.java`
   - **Bloco duplicado:** Construção de cláusulas SQL com `StringBuilder` + concatenação condicional de filtros
   - **Estratégia:** Extract module — `SqlBuilder` utility com fluent API
   - **Estimativa:** -70 linhas, 2h

3. **DT-088 — Validação de plano ativo duplicada:**
   - **Arquivos:** `PlanService.java`, `SubscriptionService.java`
   - **Bloco duplicado:** Verificação `plan.getStatus() != ACTIVE` + lançamento de exceção (~12 linhas)
   - **Estratégia:** Extract function — `PlanService.validateActive(Plan plan)`
   - **Estimativa:** -12 linhas, 0.5h

4. **Duplicação em testes de integração:** `BaseIntegrationTest.java` configura PostgreSQL container — mas `RLSIsolationTest.java`, `RbacAspectIntegrationTest.java`, etc. cada um tem sua própria configuração de container. Potencial para extrair `SharedPostgresContainer`.
   - **Arquivos:** 6 classes de teste de integração
   - **Estratégia:** Extract base class — `SharedPostgresContainer` com singleton container
   - **Estimativa:** -40 linhas, 1h

**Resumo de duplicação:**
- % estimada de duplicação no projeto: ~8-12% (moderada)
- Top 5 arquivos mais duplicados: RowMappers (6), Repositories (4), Services (2), Tests (6)
- Estratégias sugeridas: 2 extract function, 1 extract module, 1 extract base class

### tech-debt (8 achados)

**Categorização e priorização de dívida técnica estrutural:**

| ID | Categoria | Localização | Métrica | Priority | Severidade | Impacto Vel. (h/mês) | ROI (h inv. / h econ.) |
|:---|:---|:---|:---|:---:|:---:|:---:|:---|
| DT-095 | Infrastructure | `docker-compose.yml` (inexistente) | 0% ambiente local funcional | 45 | 🔴 | 80h (12 devs × 6.5h setup manual) | 1h / 80h = 80x |
| DT-096 | Architecture | `JwtAuthenticationFilter.java` | 2 claims não extraídos | 42 | 🔴 | 20h (T-065, T-066 bloqueadas) | 2h / 20h = 10x |
| DT-045 | Dependency | `pom.xml` Flyway version | 2 majors de atraso | 36 | 🟡 | 2h/mês (bugs de migração) | 1h / 24h = 24x |
| DT-068 | Dependency | `pom.xml` PostgreSQL driver | CVE CVSS 7.5 | 40 | 🔴 | 0h (risco) | 0.5h / evita incidente = ∞ |
| DT-124 | Architecture | Onboarding state machine | 0% especificação formal | 32 | 🟡 | 16h (retrabalho em edge cases) | 3h / 16h = 5.3x |
| DT-102 | Code | `JwtAuthenticationFilter.java` | 2 decodificações/request | 28 | 🟡 | 0.5h/mês (latência) | 2h / 6h = 3x |
| DT-110 | Architecture | Rate limiting design | Decisão em aberto | 30 | 🟡 | 4h (retrabalho se escolha errada) | 0.5h / 4h = 8x |
| DT-114 | Code | Checkstyle 711 violações | 100% fora do padrão | 15 | 🔵 | 2h/mês (code review) | 20h / 24h = 1.2x |

**Plano de remediação faseado:**

**Fase 1 (Imediato — antes da sprint):** DT-095, DT-096, DT-045, DT-068, DT-099, DT-100 (6 itens, priority ≥36, ~6.5h)

**Fase 2 (Durante a sprint — Frente 0 ou junto com features):** DT-023, DT-102, DT-110, DT-124, DT-107, DT-108, DT-115, DT-116, DT-097, DT-098 (10 itens, priority 25-39, ~20h)

**Fase 3 (Sprints futuras):** DT-031, DT-083, DT-084, DT-086, DT-087, DT-088, DT-089, DT-090, DT-091, DT-092, DT-093, DT-094, DT-112, DT-113, DT-114, DT-117, DT-119, DT-120, DT-122, DT-125, DT-126 (21 itens, priority <25)

### security-review (9 achados)

**Segurança de código, dependências, licenças e supply chain:**

#### Dependency CVEs & Vulnerabilities

| ID | Componente | Versão Atual | CVE/Issue | CVSS | Severidade | Ação |
|:---|:---|:---|:---|:---:|:---:|:---|
| DT-068 | PostgreSQL JDBC Driver | 42.7.10 | CVE-2026-42198 (DoS) | 7.5 | 🔴 | Bump → 42.7.11 |
| DT-045 | Flyway | 10.22.0 | 2 majors atrás (12.11.0) | — | 🟡 | Bump → 12.11.0 |
| DT-092 | springdoc-openapi | 2.8.8 | 8 minors atrás (2.8.16) | — | 🔵 | Bump → 2.8.16 |

#### Injection Vulnerabilities

1. **DT-100 — Configuração OAuth2 incompleta:** `application.yml` configura Resource Server (JWT validation) mas NÃO OAuth2 Client (Authorization Code Flow). Sem `client-id`, `client-secret`, `authorization-grant-type`, o fluxo de login não funciona.
   - **Severidade:** 🔴 Critical
   - **Ação:** Adicionar propriedades `spring.security.oauth2.client.registration.keycloak.*`

2. **DT-109 — Email não testável:** `spring-boot-starter-mail` configurado para `localhost:1025` (MailHog) mas sem MailHog rodando (docker-compose pendente). Recuperação de senha (T-058) depende de SMTP.
   - **Severidade:** 🟡 High
   - **Ação:** Incluir MailHog no docker-compose (DT-095)

#### Secrets & Exposure

3. **DT-100 — Client secret no application.yml:** `client-secret` do Keycloak será armazenado em `application.yml` — isso é aceitável para dev, mas deve ser externalizado via variável de ambiente em staging/prod.
   - **Severidade:** 🟡 High
   - **Ação:** Documentar `KEYCLOAK_CLIENT_SECRET` como env var obrigatória em staging/prod

4. **DT-126 — Dados sensíveis em logs:** `application-dev.yml` com DEBUG logging pode expor tokens JWT e senhas em desenvolvimento. DT-033 nunca resolvido.
   - **Severidade:** 🔵 Medium
   - **Ação:** Adicionar `logging.level.org.springframework.security=INFO` e filter de masking

#### Auth & Access Control

5. **DT-096 — Claims de autorização ausentes:** JWT não carrega `modules[]` e `business_unit_ids[]` — App Switcher (F04-04) e escopo de BU (F04-02) não funcionam sem esses claims.
   - **Severidade:** 🔴 Critical
   - **Ação:** Configurar Keycloak realm para incluir claims nos tokens; extrair no `JwtAuthenticationFilter`

6. **DT-121 — AuthenticationException sem handler:** 401 não é tratado explicitamente — erros de autenticação viram 500.
   - **Severidade:** 🟡 High
   - **Ação:** Adicionar `@ExceptionHandler(AuthenticationException.class)` no `GlobalExceptionHandler`

7. **DT-123 — Sem testes de segurança para F04-03:** Dashboard do cliente não tem cenários de segurança na suite de testes.
   - **Severidade:** 🟡 High
   - **Ação:** Adicionar TC-F04-03-SEC-001: isolamento entre tenants; TC-F04-03-SEC-002: acesso sem onboarding concluído

#### Licenses & Supply Chain

8. **Dependências com licenças verificadas:** Todas as 21 dependências no `pom.xml` têm licenças compatíveis (Apache 2.0, MIT, EPL, BSD). Nenhuma dependência GPL viral ou AGPL encontrada.

9. **Keycloak versão:** A versão do Keycloak (26) é recente (março 2026). Sem CVEs críticas conhecidas. Container oficial `quay.io/keycloak/keycloak:26.0`.

**Resumo:** 21 dependências auditadas, 1 CVE crítica ativa (CVSS 7.5), 0 licenças problemáticas, 6 configurações de segurança pendentes.

---

## Recomendações Prioritárias

> **Esta seção lista débitos que DEVEM ou DEVERIAM ser tratados na sprint atual.**
> Inclui tanto débitos recém-descobertos (Agentes 1-9) quanto débitos do backlog (sprints anteriores) que se tornaram críticos para a sprint atual.

### 🔴 Bloqueantes (impeditivos — devem ser corrigidos ANTES de iniciar a sprint)

Débitos que **impedem** o início ou o avanço da Sprint 5. Sem correção, o desenvolvimento para ou produz código quebrado.

| ID (TASKS.md) | DT-XXX | Sprint Origem | Ação Corretiva | Estimativa | Responsável |
|:---|:---|:---|:---|:---:|:---|
| T-133.DT-095 | DT-095 | Sprint 5 | Criar `docker-compose.yml` com PostgreSQL 17 + Keycloak 26 + MailHog. Exportar `realm-config.json` do Keycloak | 3h | A definir |
| T-134.DT-045 | DT-045 | Sprint 3 | Bump Flyway 10.22.0→12.11.0 no pom.xml. Atualizar `flyway-database-postgresql`. Rodar `mvn flyway:migrate` | 1h | A definir |
| T-135.DT-068 | DT-068 | Sprint 4 | Bump PostgreSQL driver 42.7.10→42.7.11 no pom.xml | 0.5h | A definir |
| T-136.DT-096 | DT-096 | Sprint 5 | Atualizar `JwtAuthenticationFilter` para extrair claims `modules[]` e `business_unit_ids[]` do JWT. Popular `TenantContext` | 2h | A definir |
| T-137.DT-099 | DT-099 | Sprint 5 | Adicionar `spring-boot-starter-oauth2-client` ao pom.xml. Configurar `SecurityConfig` para Authorization Code Flow + Resource Server | 1.5h | A definir |
| T-138.DT-100 | DT-100 | Sprint 5 | Configurar `application.yml` com `spring.security.oauth2.client.registration.keycloak.*` e `provider.keycloak.*` | 1h | A definir |

**Total Frente 0 (Pré-Sprint):** 6 tarefas · ~9h (≈1.5 dia)

### 🟡 Recomendados (devem ser tratados — podem ser incluídos no backlog da sprint)

Débitos que **não bloqueiam** o início da sprint mas que, se ignorados, acumulam risco técnico significativo.

| ID (TASKS.md) | DT-XXX | Sprint Origem | Ação Corretiva | Estimativa | Sprint sugerida |
|:---|:---|:---|:---|:---:|:---|
| T-139.DT-023 | DT-023 | Sprint 3 | Implementar `findAllKeyset()` no `BaseRepository` | 3h | Sprint 5 |
| T-140.DT-097 | DT-097 | Sprint 5 | Corrigir contagem de cenários: 21→28 no SPRINT-TEST-SUITE.md e SPRINT-CARD.md | 0.5h | Sprint 5 |
| T-141.DT-098 | DT-098 | Sprint 5 | Conectar `TenantContext.businessUnitIds`/`modules` ao `JwtAuthenticationFilter` (junto com DT-096) | Incluso em T-136 | Sprint 5 |
| T-142.DT-107 | DT-107 | Sprint 5 | Adicionar campo `is_matrix` à entidade `BusinessUnit` via migration V007 | 1h | Sprint 5 |
| T-143.DT-108 | DT-108 | Sprint 5 | Documentar máquina de estados de TenantStatus (transições válidas) e implementar validação no `OnboardingService` | 1.5h | Sprint 5 |
| T-144.DT-110 | DT-110 | Sprint 5 | Decidir e implementar rate limiting via Filter (não @Aspect) + Caffeine | 3h (parte de T-059) | Sprint 5 |
| T-145.DT-124 | DT-124 | Sprint 5 | Criar diagrama de estados do onboarding antes de codificar `OnboardingService` | 2h | Sprint 5 |
| T-146.DT-121 | DT-121 | Sprint 5 | Adicionar `@ExceptionHandler(AuthenticationException.class)` → 401 RFC 7807 | 0.5h | Sprint 5 |
| T-147.DT-106 | DT-106 | Sprint 5 | Adicionar cenários de teste ausentes: timeout sessão, complexidade senha, passo 3, segurança F04-03 | 2h | Sprint 5 |
| T-148.DT-102 | DT-102 | Sprint 5 | Consolidar dupla decodificação JWT via `JwtAuthenticationConverter` (junto com DT-076) | 2h | Sprint 5 |

### 🔵 Desejáveis (nice-to-have — se houver capacidade)

| ID (TASKS.md) | DT-XXX | Sprint Origem | Ação Corretiva | Estimativa |
|:---|:---|:---|:---|:---:|
| T-149.DT-086 | DT-086 | Sprint 4 | Extrair helper `AuditFieldsRowMapper` para eliminar duplicação | 1.5h |
| T-150.DT-089 | DT-089 | Sprint 4 | Injete `ObjectMapper` do Spring no `AuditAspect` | 0.5h |
| T-151.DT-090 | DT-090 | Sprint 4 | Substituir `OffsetDateTime.now()` → `OffsetDateTime.now(ZoneOffset.UTC)` | 0.5h |
| T-152.DT-092 | DT-092 | Sprint 4 | Bump springdoc 2.8.8→2.8.16 | 0.5h |
| T-153.DT-093 | DT-093 | Sprint 4 | Externalizar CORS origins para `application.yml` | 0.5h |
| T-154.DT-101 | DT-101 | Sprint 5 | Atualizar mitigação de riscos no SPRINT-CARD.md (docker-compose pendente) | 0.5h |
| T-155.DT-112 | DT-112 | Sprint 5 | Atualizar header do SPECS.md para Sprint 5 | 0.25h |
| T-156.DT-113 | DT-113 | Sprint 5 | Recalcular progresso no TASKS.md (88+28+12 pendentes) | 0.25h |

---

## Decisão do Time

> **Decisão tomada em 2026-07-17.** O time definiu 3 frentes de correção para a Sprint 5, organizadas por prioridade e urgência.

### Estrutura de Frentes

| Frente | Escopo | Quantidade | Estimativa | Quando |
|:---|:---|:---:|:---:|:---|
| **Frente 0** | 🔴 Bloqueantes — impedem o início da sprint | 6 tarefas | ~9h (≈1.5 dia) | ANTES da sprint |
| **Frente 1** | 🟡 Recomendados — devem ser tratados junto com features | 10 tarefas | ~17.5h (≈3 dias) | Durante a sprint |
| **Frente 2** | 🔵 Desejáveis — se houver capacidade | 8 tarefas | ~4.5h (≈0.5 dia) | Durante a sprint (opcional) |
| **Frente 3** | 🎯 Tarefas da Sprint 5 (T-057 a T-068) | 12 tarefas | ~18.5 dias-homem | Corpo da sprint |

### Decisão por Débito

| ID | Decisão | Frente | Sprint alvo | Justificativa |
|:---|:---|:---:|:---:|:---|
| DT-095 | **Tratar agora** | Frente 0 | Sprint 5 | Sem docker-compose, 100% das tarefas bloqueadas — desenvolvimento local inviável |
| DT-045 | **Tratar agora** | Frente 0 | Sprint 5 | Flyway 2 majors atrás — migrations V007+ devem seguir versão atualizada |
| DT-068 | **Tratar agora** | Frente 0 | Sprint 5 | CVE-2026-42198 ativa (CVSS 7.5) — risco de negação de serviço |
| DT-096 | **Tratar agora** | Frente 0 | Sprint 5 | JWT sem claims `modules[]` e `business_unit_ids[]` — T-065 e T-066 bloqueadas |
| DT-099 | **Tratar agora** | Frente 0 | Sprint 5 | Sem OAuth2 Client, Authorization Code Flow não funciona — T-057 e T-058 bloqueadas |
| DT-100 | **Tratar agora** | Frente 0 | Sprint 5 | Configuração OAuth2 incompleta — login redirect não funciona |
| DT-023 | **Tratar agora** | Frente 1 | Sprint 5 | Keyset pagination designado para Sprint 5 desde Sprint 3 |
| DT-097 | **Tratar agora** | Frente 1 | Sprint 5 | Métricas incorretas propagam erro para reports e planning |
| DT-098 | **Tratar agora** | Frente 1 | Sprint 5 | TenantContext com campos não consumidos — Sprint 5 implementa consumo |
| DT-107 | **Tratar agora** | Frente 1 | Sprint 5 | Flag `is_matrix` necessária para T-062 (primeira BU = Matriz) |
| DT-108 | **Tratar agora** | Frente 1 | Sprint 5 | Máquina de estados documentada evita edge cases no OnboardingService |
| DT-110 | **Tratar agora** | Frente 1 | Sprint 5 | Decisão de design para T-059 — Filter + Caffeine em vez de @Aspect |
| DT-124 | **Tratar agora** | Frente 1 | Sprint 5 | Diagrama de estados antes de codificar evita retrabalho |
| DT-121 | **Tratar agora** | Frente 1 | Sprint 5 | AuthenticationException sem handler — 500 em vez de 401 |
| DT-106 | **Tratar agora** | Frente 1 | Sprint 5 | Cenários de teste ausentes — cobertura insuficiente |
| DT-102 | **Tratar agora** | Frente 1 | Sprint 5 | Dupla decodificação JWT — performance e risco de inconsistência |
| DT-086 | **Tratar agora** | Frente 2 | Sprint 5 | RowMapper duplicação — melhoria de qualidade |
| DT-089 | **Tratar agora** | Frente 2 | Sprint 5 | ObjectMapper não-Spring — consistência de serialização |
| DT-090 | **Tratar agora** | Frente 2 | Sprint 5 | OffsetDateTime sem UTC — timestamps inconsistentes |
| DT-092 | **Tratar agora** | Frente 2 | Sprint 5 | springdoc 8 minors atrás |
| DT-093 | **Tratar agora** | Frente 2 | Sprint 5 | CORS hardcoded — inflexível para múltiplos ambientes |
| DT-101 | **Tratar agora** | Frente 2 | Sprint 5 | Atualizar mitigação no SPRINT-CARD (documentação) |
| DT-112 | **Tratar agora** | Frente 2 | Sprint 5 | Atualizar header SPECS.md (documentação) |
| DT-113 | **Tratar agora** | Frente 2 | Sprint 5 | Recalcular progresso TASKS.md (documentação) |
| DT-031 | **Postergar** | — | Sprint 6-7 | Checkstyle: redução progressiva 300→200→100→0 |
| DT-034 | **Postergar** | — | Sprint 6 | Address.java — aguardar verificação de uso em M6 |
| DT-044 | **Postergar** | — | Sprint 7 | logback-spring.xml — infraestrutura, não bloqueia features |
| DT-083 | **Postergar** | — | Sprint 6 | Dead code Address + enums — depende de DT-034 |
| DT-084 | **Postergar** | — | Sprint 6 | EmailService YAGNI — pode quebrar testes |
| DT-087 | **Postergar** | — | Sprint 6 | SQL duplicado — refactor com risco de regressão |
| DT-088 | **Postergar** | — | Sprint 6 | Validação plano ativo duplicada — baixa prioridade |
| DT-091 | **Postergar** | — | Sprint 6 | Endpoint stub resend-invite — funcionalidade não prioritária |
| DT-094 | **Postergar** | — | Sprint 7 | README.md — documentação final |
| DT-114 | **Postergar** | — | Sprint 6-7 | Checkstyle — progressivo com DT-031 |
| DT-117 | **Postergar** | — | Sprint 6 | OffsetDateTime.now() — correção trivial mas ampla |
| DT-119 | **Postergar** | — | Sprint 5 | SPRINT-REVIEW — preencher durante a sprint |
| DT-120 | **Postergar** | — | Sprint 6 | RowMapper duplicação — refactor cosmético |
| DT-122 | **Postergar** | — | Sprint 6 | ObjectMapper manual — DT-089 pendente |
| DT-125 | **Postergar** | — | Sprint 6 | JaCoCo 72% — melhorar com novos testes |
| DT-126 | **Postergar** | — | Sprint 6 | Dados sensíveis em logs — masking |

---

## Débitos Técnicos Elegíveis para Sprints Futuras

> **Esta seção lista débitos que NÃO serão tratados na sprint atual, mas permanecem no radar para sprints futuras.**
> Inclui tanto débitos do backlog antigo quanto débitos recém-descobertos que o time decidiu postergar.

Débitos que o time decidiu **explicitamente postergar** — seja por baixa severidade, alta complexidade, ou dependência de outros fatores. Estes itens devem ser reavaliados no próximo ciclo de `PROMPT-GENERATE-IDENTIFIED-TECHNICAL-DEBT`.

| DT-XXX | Sprint Origem | Descrição | Severidade | Bloqueante? | Skill | Complexidade | Sprint Sugerida | Justificativa do Adiamento |
|:---|:---|:---|:---:|:---:|:---:|:---:|:---|:---|
| DT-031 | Sprint 3 | Reduzir Checkstyle de 711 violações para 0 | 🔵 | NÃO | BACKLOG | H | Sprint 5-7 | Progressivo: reduzir 300→200→100→0 ao longo de 3 sprints |
| DT-034 | Sprint 3 | Remover Address.java se não usado até Sprint 6 | 🔵 | NÃO | BACKLOG | L | Sprint 6 | Aguardar verificação de uso em M6 (BUs e Catálogo) |
| DT-044 | Sprint 3 | Criar logback-spring.xml | 🔵 | NÃO | BACKLOG | M | Sprint 7 | Infraestrutura — não bloqueia features |
| DT-083 | Sprint 4 | Dead code: Address + 5 enums (~211 linhas) | 🔵 | NÃO | BACKLOG | L | Sprint 6 | Depende de DT-034 |
| DT-084 | Sprint 4 | EmailService interface YAGNI | 🔵 | NÃO | BACKLOG | L | Sprint 6 | Pode impactar testes que mockam a interface |
| DT-087 | Sprint 4 | Construção SQL duplicada (~70 linhas) | 🔵 | NÃO | JSCPD | M | Sprint 6 | Refactor estrutural — risco de regressão |
| DT-088 | Sprint 4 | Validação de plano ativo duplicada (12 linhas) | 🔵 | NÃO | JSCPD | L | Sprint 6 | Baixa prioridade |
| DT-091 | Sprint 4 | Endpoint stub resend-invite | 🔵 | NÃO | BACKLOG | L | Sprint 6 | Funcionalidade não prioritária |
| DT-094 | Sprint 4 | README.md vazio | 🔵 | NÃO | BACKLOG | M | Sprint 7 | Documentação — última sprint |
| DT-114 | Sprint 5 | Checkstyle 711 violações | 🔵 | NÃO | CREV | H | Sprint 5-7 | Progressivo com DT-031 |
| DT-117 | Sprint 5 | OffsetDateTime.now() sem UTC | 🔵 | NÃO | PONY | L | Sprint 6 | Correção trivial mas afeta todos os timestamps |
| DT-119 | Sprint 5 | SPRINT-REVIEW.md vazio | 🔵 | NÃO | CR | L | Sprint 5 | A preencher durante a sprint |
| DT-120 | Sprint 5 | RowMapper duplicação de auditoria | 🔵 | NÃO | JSCPD | L | Sprint 6 | Refactor cosmético |
| DT-122 | Sprint 5 | ObjectMapper manual no AuditAspect | 🔵 | NÃO | PONY | L | Sprint 6 | DT-089 pendente desde Sprint 4 |
| DT-125 | Sprint 5 | JaCoCo min coverage 72% abaixo da meta 80% | 🔵 | NÃO | DEBT | M | Sprint 5-6 | Melhorar progressivamente com novos testes |
| DT-126 | Sprint 5 | Dados sensíveis em logs DEBUG | 🔵 | NÃO | SEC | L | Sprint 6 | Configuração de masking |

> **Regra de reavaliação:** Na próxima execução deste prompt (Sprint 6), os débitos desta seção devem ser relidos e reclassificados — um débito que era 🔵 pode se tornar 🔴 se o contexto mudou.

---

## Análise de Impacto nos Documentos

> **Esta seção deve ser preenchida APÓS a decisão do time,**
> para cada débito que foi aceito para correção na sprint atual.

### Impacto nos Documentos-Mestre

| ID | Documento | Impacto | Ação necessária |
|:---|:---|:---|:---|
| DT-045 | ARCHITECTURE.md | Stack atualizado (Flyway 12.11.0) | Atualizar §Stack e ADR-L01 |
| DT-068 | ARCHITECTURE.md | Stack atualizado (PostgreSQL driver 42.7.11) | Atualizar §Stack |
| DT-095 | ARCHITECTURE.md | Nova seção: ambiente de desenvolvimento local | Adicionar §Dev Environment com docker-compose |
| DT-096 | ARCHITECTURE.md | Claims do JWT expandidos | Atualizar §4 Pipeline de Segurança |
| DT-099 | ARCHITECTURE.md | OAuth2 Client adicionado | Atualizar ADR-04 (Keycloak) com fluxo Auth Code |
| DT-107 | SPECS.md | Nova entidade: campo `is_matrix` em BusinessUnit | Atualizar §6 Data Model |
| DT-108 | SPECS.md | Máquina de estados de TenantStatus documentada | Adicionar RN para transições de status |
| DT-124 | ARCHITECTURE.md | Diagrama de estados do onboarding | Adicionar §Onboarding State Machine |
| DT-023 | TASKS.md | Nova task T-139.DT-023 | Adicionar ao M5 ou M6 |
| DT-096 | TASKS.md | T-065 e T-066 dependem deste débito | Adicionar como pré-requisito |

### Impacto nos Artefatos da Sprint

| ID | Documento | Impacto | Ação necessária |
|:---|:---|:---|:---|
| DT-097 | SPRINT-TEST-SUITE.md | Corrigir contagem 21→28 | Atualizar cabeçalho |
| DT-097 | SPRINT-CARD.md | Corrigir métrica 21→28 | Atualizar tabela de métricas |
| DT-101 | SPRINT-CARD.md | Atualizar mitigação de riscos | Substituir "Keycloak container no Docker Compose" por referência a DT-095 |
| DT-106 | SPRINT-TEST-SUITE.md | Adicionar 5 novos cenários | Timeout sessão, complexidade senha, passo 3, segurança F04-03 |
| DT-110 | SPRINT-CARD.md | Especificar decisão de design para T-059 | "Filter + Caffeine" em vez de "@Aspect ou filter" |

---

---

## Resumo da Decisão

| Frente | Tasks | Débitos | Estimativa | Status |
|:---|:---:|:---:|:---:|:---:|
| Frente 0 (Pré-Sprint) | T-133.DT-095 a T-138.DT-100 | 6 | ~9h (≈1.5 dia) | 🔜 A iniciar |
| Frente 1 (Recomendados) | T-139.DT-023 a T-148.DT-102 | 10 | ~17.5h (≈3 dias) | 🔜 Planejado |
| Frente 2 (Desejáveis) | T-149.DT-086 a T-156.DT-113 | 8 | ~4.5h (≈0.5 dia) | 🔜 Opcional |
| Frente 3 (Sprint 5) | T-057 a T-068 | 12 (features) | ~18.5 dias-homem | 🔜 Corpo da sprint |
| **Total** | **36 tarefas** | **24 débitos + 12 features** | **~22.5 dias** | |
| **Postergados** | — | 16 débitos | — | Sprint 6-7 |

### Impacto nos Documentos (Pós-Decisão)

#### Documentos de Negócio
| ID | Documento | Impacto | Ação necessária |
|:---|:---|:---|:---|
| — | — | Sem impacto em docs de negócio | Nenhuma ação necessária |

#### Documentos-Mestre
| ID | Documento | Impacto | Ação necessária |
|:---|:---|:---|:---|
| DT-045, DT-068, DT-092 | ARCHITECTURE.md | Stack atualizado | Bump versão, atualizar §Stack |
| DT-095 | ARCHITECTURE.md | Nova seção Dev Environment | Adicionar §Dev Environment com docker-compose |
| DT-096, DT-099, DT-100 | ARCHITECTURE.md | Pipeline de segurança expandido | Atualizar §4 + ADR-04 |
| DT-107, DT-108, DT-124 | SPECS.md | Novas RNs e campos | Atualizar §3.1 + §6 |
| DT-023, DT-045, etc. | TASKS.md | 24 novas tasks (T-133 a T-156) | APPEND ao final do documento |
| DT-097 | PRD.md | Métricas atualizadas | Atualizar referências de cenários (21→28) |

#### Artefatos da Sprint
| ID | Documento | Impacto | Ação necessária |
|:---|:---|:---|:---|
| DT-097 | SPRINT-TEST-SUITE.md | Corrigir contagem 21→28 | Atualizar cabeçalho |
| DT-097 | SPRINT-CARD.md | Corrigir métrica 21→28 + adicionar Frentes 0/1/2 | Atualizar §Métricas + §Sprint Backlog |
| DT-101 | SPRINT-CARD.md | Atualizar mitigação de riscos | Referenciar DT-095 |
| DT-106 | SPRINT-TEST-SUITE.md | Adicionar 5 cenários | Timeout sessão, senha, passo 3, segurança F04-03 |

---

🤖 *Análise gerada em 2026-07-17. 42 achados consolidados a partir de 9 skills + backlog de 28 débitos de sprints anteriores. Decisão do time: 24 débitos tratados na Sprint 5 (6 Frente 0 + 10 Frente 1 + 8 Frente 2), 16 postergados para Sprints 6-7.*
