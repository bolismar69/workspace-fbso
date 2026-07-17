# IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin

- **Sprint alvo:** 3 de 7 — Portal Admin + Contas e Planos
- **Marco:** M2 (EP-01) + M3 (EP-02)
- **Data da análise:** 16/07/2026
- **Skills executadas:** engineering-skills, caveman-review, superpowers:brainstorming, ponytail-review, code-review, codebase-cleanup-tech-debt, codebase-cleanup-deps-audit
- **Stack:** Java 25 + Spring Boot 3.5.1 + PostgreSQL 17, JDBC Template (não JPA), Flyway
- **Total de achados:** 47 (18 🔴 críticos, 19 🟡 riscos, 10 🔵 nits)
- **Impeditivos para iniciar a sprint:** 12 SIM (devem ser corrigidos ANTES do primeiro commit de feature)

---

## Resumo Executivo

A auditoria multidisciplinar com 7 skills revelou 47 débitos técnicos — **12 são impeditivos para iniciar a Sprint 3**. Os 4 achados mais graves são: (1) **Spring Boot 3.5.1 com CVEs de authentication bypass (CVSS 8.2)** e Jackson com RCE (CVSS 8.1) — dependências precisam ser atualizadas; (2) **AuditAspect quebrado** — chama TenantContext depois que o filter já deu clear(), fazendo a auditoria nunca funcionar; (3) **BaseRepository sem save()/update()** — 7 repositories da Sprint 3 vão duplicar INSERT/UPDATE manualmente (~100 linhas de boilerplate por repository); (4) **JaCoCo 0.8.12 quebrado no Java 25** — a meta de cobertura ≥80% do DoD jamais foi verificada. A recomendação é **resolver os 12 impeditivos antes de iniciar a implementação das features** (estimativa: 6-8h de correções preparatórias).

---

## Matriz de Débitos Técnicos (Top 12 Impeditivos)

> **Legenda:** CR=caveman-review, PONY=ponytail-review, ARCH=brainstorming, ENG=engineering-skills, CODE=code-review, DEBT=codebase-cleanup-tech-debt, DEPS=codebase-cleanup-deps-audit

| ID | Arquivo/Artefato | Achado | Sev. | Skill | Compl. | Bloq.? | Efeito se não tratado |
|:---|:---|:---|:---:|:---:|:---:|:---:|:---|
| **DT-001** | `pom.xml` (spring-boot-starter-parent 3.5.1) | CVE-2026-22733 + CVE-2026-22731: Authentication bypass via Actuator (CVSS 8.2). CVE-2026-54512 + CVE-2026-54513: RCE via Jackson PTB bypass (CVSS 8.1) | 🔴 | DEPS | M | **SIM** | Aplicação em produção vulnerável a auth bypass e execução remota de código |
| **DT-002** | `AuditAspect.java:46-78` | @Async chama TenantContext.getTenantId() DEPOIS que JwtAuthenticationFilter já deu clear() no ThreadLocal. Auditoria **nunca funciona** — gera log WARN e grava UUID.randomUUID() | 🔴 | ARCH | M | **SIM** | T-035/T-036 (Auditoria) não podem ser concluídos. RN04-01, RN08-01, RN08-02 violadas. Falsa sensação de auditoria funcional |
| **DT-003** | `BaseRepository.java:28-140` | Não tem save() nem update(). 7 repositories da Sprint 3 vão duplicar INSERT/UPDATE manualmente (~100 linhas boilerplate cada) | 🔴 | ARCH | M | **SIM** | Atraso de 2-3 dias na Sprint 3. Risco de bugs em mapeamento de colunas (created_by, updated_by esquecidos) |
| **DT-004** | `pom.xml:173` (jacoco 0.8.12) | JaCoCo 0.8.12 rejeita class file v69 (Java 25). Cobertura NUNCA foi verificada. DoD exige ≥80% | 🔴 | DEBT | M | **SIM** | Meta de cobertura do DoD impossível de verificar. Código pode ter 0% de cobertura sem detecção |
| **DT-005** | `RbacAspect.java:48-54` | Matriz hardcoded sem entradas para TENANT, PLAN, SUBSCRIPTION, DASHBOARD. Endpoints da Sprint 3 retornarão 403 para todas as roles exceto ADMIN | 🔴 | ARCH | L | **SIM** | T-019, T-027, T-031, T-034, T-036 não funcionam com RBAC |
| **DT-006** | `TenantAwareDataSource.java:69-72` | SQLException no SET tenant_id é engolida com log.debug. Se falhar, conexão volta ao pool com tenant_id **residual de outra requisição** → vazamento cross-tenant | 🔴 | CODE | M | **SIM** | Violação gravíssima de isolamento multi-tenant. Dados de um tenant visíveis para outro |
| **DT-007** | `pom.xml` (spring-boot-starter-mail) | Dependência spring-boot-starter-mail AUSENTE. application.yml tem config mail mas sem dependência no classpath → JavaMailSender não existe | 🔴 | ARCH | L | **SIM** | T-028 (convite tenant por email) bloqueado. US-009 (link ativação) não pode ser implementada |
| **DT-008** | `AuditAspect.java:56,85-91` | extractEntityId() retorna args[0].toString(). Para TenantCreateRequest, retorna "TenantCreateRequest@1a2b3c" → UUID.fromString() quebra → entity_id vira UUID.randomUUID() | 🔴 | ENG | M | **SIM** | Registros de auditoria com entity_id aleatório. RN08-01 (100% rastreabilidade) violada |
| **DT-009** | `Plan.version` vs `Subscription` | Plan.version é incrementado mas NUNCA consultado por Subscription. RN06-02 ("alteração de preço não afeta assinaturas existentes") não tem implementação — subscription referencia plan_id sem snapshot de preço | 🔴 | ARCH | M | **SIM** | Alteração de preço de plano afeta assinaturas ativas. Violação de RN06-02. Risco contratual |
| **DT-010** | `pom.xml:228-231` (Surefire) | Padrão `**/security/**/*Test.java` NÃO captura JwtAuthenticationFilterTest. Testes de segurança NUNCA executam no build | 🔴 | ENG | L | **SIM** | Falso positivo — build passa mas testes de segurança não rodam |
| **DT-011** | `JwtAuthenticationFilter.java:126-132` | sendUnauthorized() escreve JSON manual com .formatted(). Se message contiver aspas → JSON inválido → 500. Duplica ErrorResponse sem necessidade | 🔴 | ENG | L | **SIM** | Respostas 401 com JSON malformado. Inconsistência com RFC 7807 do resto da API |
| **DT-012** | `BusinessException.java:9-13` (Javadoc) | Referencia 4 subclasses (DuplicateCnpjException, InvalidStatusTransitionException, PlanHasActiveSubscribersException, TenantNotFoundException) que **não existem** | 🟡 | ENG | M | **SIM** | Services da Sprint 3 forçados a usar BusinessException genérico. Sem catch específico |

---

## Matriz de Débitos Técnicos (Demais — 35 itens)

| ID | Arquivo/Artefato | Achado | Sev. | Skill | Compl. | Bloq.? | Efeito se não tratado |
|:---|:---|:---|:---:|:---:|:---:|:---:|:---|
| DT-013 | Diretórios controller/, service/, dto/request/ | Zero implementações. 18 endpoints precisam ser criados do zero. Não é débito — é pré-implementação | 🔴 | ENG | H | NÃO | Sprint 3 começa do zero. Risco de escopo |
| DT-014 | `application.yml:8,53` | Duas chaves spring: no mesmo arquivo (já corrigido) | 🟡 | CR | L | NÃO | — |
| DT-015 | `SPRINT-CARD.md:L117` | Endpoints: 17 (corrigido para 18) | 🟡 | CR | L | NÃO | — |
| DT-016 | `SPRINT-TEST-SUITE.md:L7,L139` | Cabeçalho/footer: 55 (corrigido para 56) | 🟡 | CR | L | NÃO | — |
| DT-017 | `PLANNING-DRAFT.md:L150` | Contradição V004 "opcional" vs "pré-requisito" | 🟡 | CR | M | NÃO | Dev pode pular V004 e dashboard ficar lento |
| DT-018 | `PLANNING-DRAFT.md:L280` | Tabela 6.1: Integração=19 (corrigido para 20) | 🟡 | CR | L | NÃO | — |
| DT-019 | `PLANNING-DRAFT.md` | 35.5 homem-dias em 12 dias úteis ≈ 3 tasks/dia. Irrealista com JDBC manual. Risco de cortar 3-5 tasks | 🟡 | ARCH | M | NÃO | Alta probabilidade de não completar a sprint no prazo |
| DT-020 | `V004` (planejada) | V004 enxuta (2 índices) insuficiente para queries dashboard com JOIN tenant+subscription | 🟡 | ARCH | L | NÃO | Dashboard pode degradar com >5k registros |
| DT-021 | `AuditAspect` | AfterReturning não captura valores "antes" em updates. Colunas previous_value/new_value (JSONB) nunca populadas | 🟡 | ARCH | M | NÃO | Não-conformidade com PRD §6.4 (auditoria sem diff antes/depois) |
| DT-022 | `V003__enable_rls.sql:16` | current_setting sem fallback: falha se variável nunca foi setada (ex: flyway migrate via CLI) | 🟡 | ENG | L | NÃO | Migrations executadas fora da aplicação falham |
| DT-023 | Offset-based pagination (D2) | Com >10k registros, páginas avançadas degradam. Keyset pagination é trivial e resolve | 🟡 | ARCH | L | NÃO | Performance degradada com escala. Mitigável na Sprint 7 |
| DT-024 | `Dockerfile.jvm:25` | USER 1001 pode não existir na imagem eclipse-temurin:25-jre-alpine. Container pode falhar ao iniciar | 🟡 | ENG | L | NÃO | Falha em produção JVM com "Unable to determine user name" |
| DT-025 | `GlobalExceptionHandler.java:75-86` | Handler captura java.lang.SecurityException (nunca lançada). AccessDeniedException do Spring Security NÃO é capturado → retorna 500 em vez de 403 | 🟡 | CODE | L | NÃO | Erro 500 para acesso negado. Difícil diagnosticar |
| DT-026 | `RLSIsolationTest.java` | 6 testes leem arquivos SQL e verificam strings. NÃO testam comportamento real do PostgreSQL RLS. Não usam Testcontainers | 🟡 | ENG | M | NÃO | Falsa segurança. RLS pode estar quebrado sem detecção |
| DT-027 | `RbacAspectTest.java:96-108` | Testes de "acesso permitido" sem verify(proceed()). Se aspecto parar de chamar proceed(), teste passa mesmo assim | 🟡 | ENG | L | NÃO | Falso positivo em testes de autorização |
| DT-028 | AuditAspect sem testes | Lógica complexa (@Async, extractEntityId, UUID.fromString) sem cobertura de teste | 🟡 | ENG | M | NÃO | Bug em auditoria só descoberto em produção |
| DT-029 | `BaseRepository.java:48-122` | Padrão if(hasTenantColumn) duplicado em 4 métodos (~96 linhas). Refatorar para método helper privado | 🟡 | PONY | L | NÃO | Correção de bug precisa ser replicada em 4 lugares |
| DT-030 | `JwtAuthenticationFilter` | JWT decodificado 2× por requisição: BearerTokenAuthenticationFilter + filtro customizado. Implementar Converter<Jwt, AbstractAuthToken> eliminaria o filtro customizado | 🟡 | PONY | M | NÃO | Dobro de validações JWKS por requisição. Custo em latência |
| DT-031 | `pom.xml` (checkstyle) | 711 violações, failOnViolation=false, maxAllowedViolations=300. Checkstyle decorativo | 🟡 | DEBT | L | NÃO | Zero barreira de qualidade estática |
| DT-032 | `pom.xml` (JaCoCo excludes) | Exclui entity/ (Tenant.isOperational()), config/ (TenantAwareDataSource, SecurityConfig). ~15% do código invisível para cobertura | 🟡 | ENG | L | NÃO | Métricas de cobertura infladas artificialmente |
| DT-033 | `application-dev.yml:14-18` | spring.jdbc: DEBUG loga queries completas com parâmetros. Se perfil dev rodar em staging → vazamento de dados | 🔵 | ENG | L | NÃO | Vazamento de dados em logs |
| DT-034 | `Address.java` (95 linhas) | Value Object sem consumidor. Criado para BusinessUnit (Sprint 6) | 🔵 | PONY | L | NÃO | Código morto até Sprint 6 |
| DT-035 | `Role.java` (19 linhas) | Enum não referenciado. RbacAspect usa strings literais | 🔵 | PONY | L | NÃO | Type-safety ghost |
| DT-036 | ProductType, TaxRegime, UserStatus, Recurrence, SubscriptionStatus | 5 enums sem referência em produção (~47 linhas). Criados antecipadamente | 🔵 | PONY | L | NÃO | Código morto até respectivas sprints |
| DT-037 | `BaseEntity.markAsDeleted()` | Método nunca chamado. Soft delete é feito via SQL no BaseRepository | 🔵 | PONY | L | NÃO | Confusão entre soft delete SQL vs in-memory |
| DT-038 | `PermissionDeniedException(String)` | Construtor com mensagem nunca usado. Apenas construtor default chamado | 🔵 | PONY | L | NÃO | — |
| DT-039 | `SecurityConfig.java:106-109` | CORS origins hardcoded. Sem suporte a staging com URL diferente | 🔵 | CODE | L | NÃO | Frontend de staging bloqueado por CORS |
| DT-040 | `BaseEntity.java:31-32` | OffsetDateTime.now() usa fuso do sistema, não UTC explícito. Timestamps inconsistentes entre ambientes | 🔵 | DEBT | L | NÃO | Inconsistência em multi-fuso |
| DT-041 | `README.md:1` | Apenas "TODO". Zero documentação de setup/build/run/test | 🔵 | DEBT | L | NÃO | Onboarding lento para novos devs |
| DT-042 | Sem docker-compose.yml | Setup local requer PostgreSQL + Keycloak manuais. ~30-60min por dev | 🔵 | DEBT | M | NÃO | Barreira de entrada alta |
| DT-043 | Sem scripts de seed | Zero dados para desenvolvimento. Dev precisa popular tabelas manualmente | 🔵 | DEBT | L | NÃO | Teste exploratório improdutivo |
| DT-044 | Sem logback-spring.xml | Logs apenas em stdout. Sem rotação, sem formato estruturado, sem arquivo | 🔵 | DEBT | M | NÃO | Diagnóstico de produção limitado |
| DT-045 | Flyway 10.22.0 → 12.11.0 | Major lag de 2 versões. Sem CVEs, mas acumula dívida | 🔵 | DEPS | H | NÃO | Migração mais complexa no futuro |
| DT-046 | Testcontainers 1.20.6 → 1.21.4 | Minor lag. Sem CVEs críticas. commons-compress transitiva com CVE-2024-25710 (DoS, CVSS 8.1) — apenas em testes | 🟡 | DEPS | L | NÃO | DoS teórico em pipeline de CI |
| DT-047 | `BaseRepository.java:50` | tableName concatenado sem sanitização. Confiável hoje, frágil para futuro | 🔵 | CODE | L | NÃO | SQL injection por erro humano em subclasse futura |

---

## Achados por Skill

### engineering-skills (21 achados)
DT-008, DT-010, DT-011, DT-012, DT-013, DT-022, DT-024, DT-026, DT-027, DT-028, DT-032, DT-033 — Cobre: zero controllers/services, AuditAspect quebrado, Surefire não executa testes de segurança, JSON manual com injection, JaCoCo exclui código com lógica de domínio, RLSIsolationTest falso, Docker USER 1001, RLS sem fallback, BusinessException incompleta.

### caveman-review (6 achados)
DT-014, DT-015, DT-016, DT-017, DT-018 — 5 off-by-one corrigidos imediatamente + 1 risco de contradição V004 "opcional" vs "pré-requisito". Documentação está 95% consistente após refresh.

### superpowers:brainstorming (14 achados)
DT-002, DT-003, DT-005, DT-007, DT-009, DT-019, DT-020, DT-021, DT-023 — Cobre: AuditAspect race condition (D11), BaseRepository sem save/update (D10), RbacAspect sem recursos Sprint 3 (D12), mail ausente (D16), Plan.version decorativo (D19), planejamento subestimado (D13), V004 insuficiente (D14), offset pagination (D22).

### ponytail-review (19 achados)
DT-029, DT-030, DT-034, DT-035, DT-036, DT-037, DT-038 — net: -393 lines possible, complexity score 4/10. Principais: BaseRepository duplicação (shrink), JWT dupla validação (native), Address+5 enums sem consumidor (yagni), markAsDeleted morto (delete), construtor PermissionDeniedException não usado (delete).

### code-review (10 achados)
DT-006, DT-025, DT-039, DT-047 — Cobre: TenantAwareDataSource engole SQLException (Critical C-01), AccessDeniedException não tratado (W-04), CORS fixo, tableName sem sanitização. CodeRabbit CLI indisponível — fallback manual.

### codebase-cleanup-tech-debt (17 achados)
DT-004, DT-031, DT-040, DT-041, DT-042, DT-043, DT-044 — Cobre: JaCoCo quebrado no Java 25 (C-01), Checkstyle decorativo, fuso não-UTC, README vazio, sem docker-compose, sem seed data, sem logback.

### codebase-cleanup-deps-audit (15 achados)
DT-001, DT-045, DT-046 — Cobre: 3 CVEs críticas em Spring Boot 3.5.1 + Jackson 2.19.1 (CVE-2026-22733/22731 auth bypass CVSS 8.2, CVE-2026-54512/54513 RCE CVSS 8.1), Flyway 2 majors atrás, Testcontainers minor lag, licenças 100% compatíveis com SaaS B2B.

---

## Plano de Ação Recomendado

### Correções Pré-Sprint (12 impeditivos — ANTES de iniciar features)

| Task ID | Débito | Ação | Estimativa |
|:---|:---|:---|:---:|
| **T-015.2.DT-001** | DT-001 | Atualizar spring-boot-starter-parent 3.5.1→3.5.14 + jackson 2.19.1→2.21.4 + logback + tomcat | 1-2h |
| **T-015.3.DT-002** | DT-002 | Refatorar AuditAspect: capturar tenantId/userId no JoinPoint ANTES do @Async, passar como parâmetro | 4-6h |
| **T-015.4.DT-003** | DT-003 | Adicionar save(T entity) e update(T entity) genéricos ao BaseRepository com preenchimento automático de created_by/updated_by | 3-4h |
| **T-015.5.DT-004** | DT-004 | Atualizar JaCoCo 0.8.12→0.8.14+ (ou versão com suporte a Java 25). Se não existir, substituir por OpenClover ou verificação manual | 2-3h |
| **T-015.6.DT-005** | DT-005 | Expandir RbacAspect com entradas para TENANT, PLAN, SUBSCRIPTION, DASHBOARD. Mapear ADMIN_TENANT→todos, MANAGER_BU→view | 1-2h |
| **T-015.7.DT-006** | DT-006 | Trocar log.debug→log.error em TenantAwareDataSource.applyTenantContext(). Considerar lançar TenantIsolationException | 30min |
| **T-015.8.DT-007** | DT-007 | Adicionar spring-boot-starter-mail no pom.xml | 10min |
| **T-015.9.DT-008** | DT-008 | Corrigir extractEntityId() para validar UUID.fromString() ou usar parâmetro anotado em vez de args[0] | 1-2h |
| **T-015.10.DT-009** | DT-009 | Adicionar locked_price + locked_recurrence em subscription (Migration V005). Popular no momento da criação da assinatura | 2-3h |
| **T-015.11.DT-010** | DT-010 | Corrigir padrão Surefire: `**/*Test.java` (remover filtro restritivo) | 5min |
| **T-015.12.DT-011** | DT-011 | Reescrever sendUnauthorized() usando ObjectMapper + ErrorResponse record | 30min |
| **T-015.13.DT-012** | DT-012 | Criar 4 subclasses de BusinessException (DuplicateCnpjException, InvalidStatusTransitionException, PlanHasActiveSubscribersException, TenantNotFoundException) | 30min |

**Estimativa total pré-sprint: 16-25h** (~2-3 dias com 1 dev)

### Correções Durante a Sprint (não-bloqueantes)

| Task ID | Débito | Ação | Estimativa | Sprint |
|:---|:---|:---|:---:|:---|
| **T-039.DT-017** | DT-017 | Decidir V004 "opcional" vs "pré-requisito" com o time | 5min discussão | 3 |
| **T-040.DT-019** | DT-019 | Recalibrar day-by-day: 12 dias → ~15 dias realista com 1 dev | 15min | 3 |
| **T-041.DT-021** | DT-021 | Implementar captura de valores "antes" no AuditAspect | 2-3h | 3 |
| **T-042.DT-025** | DT-025 | Adicionar @ExceptionHandler(AccessDeniedException.class) | 10min | 3 |
| **T-043.DT-026** | DT-026 | Refatorar RLSIsolationTest para usar Testcontainers + PostgreSQL real | 3-4h | 3 |
| **T-044.DT-029** | DT-029 | Extrair hasTenantColumn branching para método helper | 1h | 3 |
| **T-045.DT-046** | DT-046 | Atualizar Testcontainers 1.20.6→1.21.4 + commons-compress 1.24.0→1.28.0 | 15min | 3 |

### Correções Pós-Sprint (sprints futuras)

| Task ID | Débito | Ação | Estimativa | Sprint |
|:---|:---|:---|:---:|:---|
| **T-087.DT-023** | DT-023 | Migrar paginação offset→keyset no BaseRepository | 1d | 5 |
| **T-088.DT-030** | DT-030 | Consolidar dupla validação JWT via Converter customizado | 4h | 4 |
| **T-089.DT-031** | DT-031 | Reduzir maxAllowedViolations checkstyle progressivamente (300→100→0) | Contínuo | 4-7 |
| **T-090.DT-034** | DT-034 | Remover Address.java se não usado até Sprint 6 | 5min | 6 |
| **T-091.DT-035** | DT-035 | Migrar RbacAspect strings→Role enum | 1h | 4 |
| **T-092.DT-042** | DT-042 | Criar docker-compose.yml (postgres+keycloak+mailhog) | 1h | 4 |
| **T-093.DT-043** | DT-043 | Criar script seed SQL com 50+ tenants + 3 planos | 30min | 3 |
| **T-094.DT-044** | DT-044 | Criar logback-spring.xml com JSON appender + rotação | 1h | 7 |
| **T-095.DT-045** | DT-045 | Migrar Flyway 10→12 (planejar breaking changes) | 4h | 5 |

---

## Decisão do Time

> **Decisão registrada em 16/07/2026.**

| ID | Decisão | Sprint alvo | Justificativa |
|:---|:---|:---:|:---|
| DT-001 a DT-012 | **Tratar agora** — Frente 0 (Pré-Sprint) | **3** | Impeditivos: sem estas correções, features da Sprint 3 não podem ser implementadas com qualidade e segurança |
| DT-017, DT-019, DT-021, DT-025, DT-026, DT-029, DT-046 | **Tratar agora** — Frente 3 (Durante a Sprint) | **3** | Não-bloqueantes mas de alto impacto: corrigir durante o desenvolvimento |
| DT-023, DT-030, DT-031, DT-034, DT-035, DT-042, DT-043, DT-044, DT-045 | **Postergar** — Sprint 4+ | **4+** | Débitos de longo prazo. DT-042/DT-043 podem ser antecipados se houver buffer |
| DT-013, DT-014, DT-015, DT-016, DT-018 | **Já corrigidos** | — | Off-by-one e inconsistências de documentação resolvidos durante o Caveman Review refresh |
| DT-020, DT-022, DT-024, DT-027, DT-028, DT-032, DT-033, DT-036, DT-037, DT-038, DT-039, DT-040, DT-041, DT-047 | **Aceitar risco** | — | Baixo impacto ou complexidade trivial. Corrigir se sobrar tempo na Sprint 3 |

### Estrutura da Sprint 3 Revisada

| Frente | Escopo | Tasks | Estimativa |
|:---|:---|:---|:---:|
| **Frente 0** | Correções Pré-Sprint (12 impeditivos DT-001 a DT-012) | 12 correções | 16-25h (2-3 dias) |
| **Frente 1** | M2 — Portal Admin (T-016 a T-023) | 8 tasks originais | ~11 dias-homem |
| **Frente 2** | M3 — Clientes e Planos (T-024 a T-038) | 15 tasks originais | ~24 dias-homem |
| **Frente 3** | Correções Durante a Sprint (7 itens não-bloqueantes) | 7 correções | ~10h |
| **Total** | | **23 tasks + 19 correções** | ~38-47 dias-homem |

> ⚠️ **Impacto no prazo:** Com as correções pré-sprint (2-3 dias), o prazo original de 12 dias úteis sobe para ~15 dias. Recomendação: estender o marco M2 em 1 semana ou reduzir escopo (mover F01-03 Alertas para Sprint 4).

---

## Análise de Impacto nos Documentos

> **Esta seção deve ser preenchida APÓS a decisão do time.**

### Impacto nos Documentos-Mestre

| ID | Documento | Impacto | Ação necessária |
|:---|:---|:---|:---|
| DT-001 | ARCHITECTURE.md | Atualizar versões de dependências no stack | Atualizar §Stack |
| DT-003 | ARCHITECTURE.md | BaseRepository ganha save()/update() — atualizar ADR | Atualizar ADR-L01 |
| DT-004 | TASKS.md | Nova task: atualizar JaCoCo | Adicionar T-015.5.DT-004 |
| DT-005 | SPECS.md | Matriz RBAC expandida para recursos Sprint 3 | Atualizar RN10-01 |
| DT-009 | SPECS.md | Novos campos locked_price/locked_recurrence em subscription | Atualizar RN06-02, RN07-02 |
| DT-009 | TEST_PLAN.md | Novos cenários: change-plan preserva preço contratado | Adicionar TC-F02-04-010 |
| DT-002 | TASKS.md | Refatoração do AuditAspect impacta T-035 | Revisar estimativa T-035 |
| DT-001 | PRD.md | Stack atualizado: Spring Boot 3.5.14, Jackson 2.21.4 | Atualizar seção Stack |

### Impacto nos Artefatos da Sprint

| ID | Documento | Impacto | Ação necessária |
|:---|:---|:---|:---|
| DT-003 | SPRINT-CARD.md | Nova subtask: refatorar BaseRepository.save/update | Adicionar T-015.4.DT-003 |
| DT-005 | SPRINT-CARD.md | RbacAspect expandido: nova verificação no DONE de T-019/T-027/T-031/T-034 | Atualizar DONE criteria |
| DT-009 | SPRINT-CARD.md | Nova migration V005: locked_price em subscription | Adicionar em T-032 |
| DT-012 | SPRINT-CARD.md | 4 novas exceções: adicionar em T-024 (entidades Tenant) | Atualizar DONE de T-024 |
| DT-002 | SPRINT-TEST-SUITE.md | Novo cenário: auditoria funcional com @Async | Adicionar TC-F02-05-009 |
| DT-009 | SPRINT-TEST-SUITE.md | Novo cenário: change-plan preserva preço contratado | Adicionar TC-F02-04-010 |

---

🤖 *Análise gerada em 16/07/2026. 47 achados consolidados a partir de 7 skills (engineering-skills, caveman-review, superpowers:brainstorming, ponytail-review, code-review, codebase-cleanup-tech-debt, codebase-cleanup-deps-audit). Documento base para decisão do time.*
