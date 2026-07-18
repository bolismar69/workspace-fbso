# SPRINT-CARD: Sprint 5 — Portal do Cliente e Onboarding

- **Sprint:** 5 de 7
- **Marco:** M5 (EP-04a)
- **Datas:** 15/09/2026 → 30/09/2026
- **Duração:** 11 dias úteis
- **Responsável:** A definir
- **Documentos-mestre:** [TASKS.md](../../TASKS.md) · [SPECS.md](../../SPECS.md)

---

> 🚫 **BRANCH OBRIGATÓRIA:** Toda implementação deste sprint DEVE usar exclusivamente a branch `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-05-portal-cliente`. Antes de começar, execute:
> ```bash
> git checkout feature/sprint-05-portal-cliente
> git branch --show-current  # deve exibir: PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-05-portal-cliente
> ```
> 📖 Detalhes completos: [PRD.md §8.4](../../PRD.md#84-estratégia-de-branching--uma-branch-por-sprint)

## 🎯 Sprint Goal

**"Login via Keycloak OIDC funcional com recuperação de senha e rate limiting. Onboarding guiado em 4 passos — primeira BU vira Matriz, tenant transita para ACTIVE ao concluir. Dashboard do cliente com cards de unidades, produtos e plano. Endpoint /auth/me retorna dados do usuário logado."**

---

## 📋 Estrutura de Frentes

| Frente | Escopo | Tasks | Débitos | Estimativa | Quando |
|:---|:---|:---:|:---:|:---:|:---|
| **Frente 0** | 🔴 Bloqueantes (pré-sprint) | 6 | DT-095, DT-045, DT-068, DT-096, DT-099, DT-100 | ~6h (≈1d) | ✅ Concluída 17/07/2026 |
| **Frente 1** | 🟡 Recomendados | 10 | DT-023, DT-097, DT-098, DT-107, DT-108, DT-110, DT-124, DT-121, DT-106, DT-102 | ~17.5h (≈3d) | Durante a sprint |
| **Frente 2** | 🔵 Desejáveis (opcional) | 8 | DT-086, DT-089, DT-090, DT-092, DT-093, DT-101, DT-112, DT-113 | ~4.5h (≈0.5d) | Se houver capacidade |
| **Frente 3** | 🎯 Features Sprint 5 | 12 | — (T-057 a T-068) | ~18.5d-h | Corpo da sprint |

> 📖 Débitos detalhados em: [IDENTIFIED-TECHNICAL-DEBT-sprint-05-portal-cliente.md](./IDENTIFIED-TECHNICAL-DEBT-sprint-05-portal-cliente.md)

---

## 📋 Sprint Backlog

### Frente 0 — Bloqueantes (Pré-Sprint) ✅ CONCLUÍDA | 17/07/2026

| ID | Tarefa | DT | Est. | Status | Critério DONE |
|:---|:---|:---|:---:|:---:|:---|
| **T-133.DT-095** | Criar `docker-compose.yml` com PostgreSQL 17 + Keycloak 26 + MailHog. Exportar `realm-config.json` | DT-095 | 3h | ✅ | `docker compose up` funcional. Keycloak admin acessível em localhost:8081. MailHog em localhost:8025 |
| **T-134.DT-045** | Bump Flyway 10.22.0→12.11.0 no pom.xml. Atualizar `flyway-database-postgresql`. Rodar `mvn flyway:migrate` | DT-045 | 1h | ✅ | `mvn flyway:migrate` sucesso. V001-V006 reaplicadas sem erro |
| **T-135.DT-068** | Bump PostgreSQL driver 42.7.10→42.7.11 no pom.xml (CVE-2026-42198, CVSS 7.5) | DT-068 | 0.5h | ✅ | `mvn dependency:tree` confirma versão 42.7.11 |
| **T-136.DT-096** | Atualizar `JwtAuthenticationFilter`: extrair claims `modules[]` + `business_unit_ids[]` do JWT. Popular `TenantContext` | DT-096 | 0h | ✅ NO-OP | Já implementado — código existente já extrai claims (6 testes passando) |
| **T-137.DT-099** | Adicionar `spring-boot-starter-oauth2-client` ao pom.xml. Configurar `SecurityConfig` para Authorization Code Flow + Resource Server | DT-099 | 1.5h | ✅ | Dependência adicionada. SecurityConfig com 2 filter chains (@Order) |
| **T-138.DT-100** | Configurar `application.yml`: `spring.security.oauth2.client.registration.keycloak.*` + `provider.keycloak.*`. Issuer URI + client-id + client-secret | DT-100 | 1h | ✅ | Login redirect para Keycloak funcional. Token JWT validado |

### Frente 1 — Recomendados

| ID | Tarefa | DT | Prio. | Est. | Critério DONE |
|:---|:---|:---|:---:|:---:|:---|
| **T-139.DT-023** | Implementar `findAllKeyset()` no `BaseRepository` — paginação keyset em vez de offset | DT-023 | Must | 3h | Keyset funcional. Testes com >10k registros |
| **T-140.DT-097** | Corrigir contagem de cenários: 21→28 no SPRINT-TEST-SUITE.md e SPRINT-CARD.md | DT-097 | Must | 0.5h | Números consistentes entre documentos |
| **T-141.DT-098** | Conectar `TenantContext.businessUnitIds`/`modules` ao `JwtAuthenticationFilter` (junto com T-136) | DT-098 | Must | ✅ NO-OP | Getters chamados em T-065 e T-066 — já implementado |
| **T-142.DT-107** | Adicionar campo `is_matrix BOOLEAN NOT NULL DEFAULT false` à entidade `BusinessUnit` via migration V007 | DT-107 | Must | 1h | Migration aplicada. Entidade atualizada. T-062 usa flag |
| **T-143.DT-108** | Documentar máquina de estados de `TenantStatus` (transições válidas). Implementar validação no `OnboardingService` | DT-108 | Must | 1.5h | Diagrama de estados no ARCHITECTURE.md. Validação impede transições inválidas |
| **T-144.DT-110** | Implementar rate limiting via **Filter** (não @Aspect) + Caffeine. 5 tentativas → bloqueio 15min | DT-110 | Must | 3h (parte de T-059) | Filter funcional. Mensagem exibe tempo restante. Testes passam |
| **T-145.DT-124** | Criar diagrama de estados do onboarding antes de codificar `OnboardingService` | DT-124 | Must | 2h | Diagrama documentado no ARCHITECTURE.md. 4 estados + transições + edge cases |
| **T-146.DT-121** | Adicionar `@ExceptionHandler(AuthenticationException.class)` → 401 RFC 7807 no `GlobalExceptionHandler` | DT-121 | Must | 0.5h | 401 JSON padronizado. Token inválido/expirado não gera 500 |
| **T-147.DT-106** | Adicionar 5 cenários de teste ausentes: timeout sessão, complexidade senha, passo 3, segurança F04-03 (×2) | DT-106 | Must | 2h | 33 cenários no total (28+5). Suite atualizada |
| **T-148.DT-102** | Consolidar dupla decodificação JWT via `JwtAuthenticationConverter` (junto com DT-076) | DT-102 | Should | 2h | 1 decodificação por request. Testes de performance passam |

### Frente 2 — Desejáveis

| ID | Tarefa | DT | Prio. | Est. | Critério DONE |
|:---|:---|:---|:---:|:---:|:---|
| **T-149.DT-086** | Extrair helper `AuditFieldsRowMapper` — eliminar duplicação de 6 campos em 5 RowMappers | DT-086 | Could | 1.5h | Helper extraído. 5 RowMappers usam helper |
| **T-150.DT-089** | Injete `ObjectMapper` do Spring no `AuditAspect` (remover `new ObjectMapper()`) | DT-089 | Could | 0.5h | @Autowired ObjectMapper. Serialização consistente |
| **T-151.DT-090** | Substituir `OffsetDateTime.now()` → `OffsetDateTime.now(ZoneOffset.UTC)` no `BaseEntity` | DT-090 | Could | 0.5h | Timestamps UTC em todos os ambientes |
| **T-152.DT-092** | Bump springdoc-openapi 2.8.8→2.8.16 no pom.xml | DT-092 | Could | 0.5h | Swagger UI funcional. Sem breaking changes |
| **T-153.DT-093** | Externalizar CORS origins para `application.yml` (`cors.allowed-origins`) | DT-093 | Could | 0.5h | Origens configuráveis por ambiente |
| **T-154.DT-101** | Atualizar mitigação de riscos no SPRINT-CARD.md (referenciar DT-095 para docker-compose) | DT-101 | Could | 0.5h | Riscos atualizados |
| **T-155.DT-112** | Atualizar header do SPECS.md: "Próximo: Sprint 5 Portal do Cliente" | DT-112 | Could | 0.25h | Header correto |
| **T-156.DT-113** | Recalcular progresso no TASKS.md: 88 + 28 (S4 pendentes) + 24 (S5 débitos) + 12 (S5 features) = 152 total | DT-113 | Could | 0.25h | Progresso preciso |

### Frente 3 — Features da Sprint 5

| ID | Tarefa | Feature | Prio. | Est. | Critério DONE |
|:---|:---|:---|:---:|:---:|:---|
| **T-057** | Configurar Keycloak realm `fbso-platform` + client Authorization Code Flow. Mapear claims (tenant_id, roles, business_unit_ids, modules). Exportar `realm-config.json` | F04-01 | Must | 2d | Realm funcional. JWT com claims corretas. Config versionada |
| **T-058** | `POST /api/v1/auth/login` (delega Keycloak). Recuperação: `POST /auth/forgot-password` (expira 1h — RN13-03), `POST /auth/reset-password` (complexidade RN13-01) | F04-01 | Must | 2d | Login funcional. Link reset 1h. Senha 8+ chars, letra+número. Sessão 60min inatividade (RN13-02) |
| **T-059** | Rate limiting: 5 tentativas → bloqueio 15min (US-039, RN13-02). Implementar via Filter + Caffeine (decisão DT-110/T-144) | F04-01 | Must | 1.5d | 5 falhas → bloqueado. Mensagem exibe tempo restante |
| **T-060** | `OnboardingService`: 4 passos obrigatórios (US-040 a US-044). Ordem obrigatória (RN14-01). Primeira BU = Matriz (RN14-02). Tenant → ACTIVE ao concluir (RN14-04). Retomável. Diagrama de estados (DT-124/T-145) | F04-02 | Must | 2d | 4 passos. Ordem validada. Status salvo permite retomar |
| **T-061** | `OnboardingController`: `GET /onboarding/status`, `PATCH /step-1`, `POST /step-2`, `POST /complete`. `@RequiresPermission` | F04-02 | Must | 1.5d | 4 endpoints. Step-2 valida CNPJ. Complete só se todos passos OK (RN14-03) |
| **T-062** | Criação primeira BU como Matriz (step-2). Reutilizar `BusinessUnitService.create()` com flag `isMatrix=true` (DT-107/T-142). CNPJ validado | F04-02 | Must | 1d | BU Matriz com is_matrix=true, parent_id=NULL. CNPJ válido |
| **T-063** | `DashboardClientService`: cards — unidades ativas, produtos catálogo, plano contratado, notificações. Cards clicáveis | F04-03 | Should | 1.5d | Cards com dados resumidos. Link para área |
| **T-064** | `DashboardClientController`: `GET /dashboard/client/summary`, `/notifications`. `@RequiresPermission`. Dashboard genérico Fase 0 (RN15-01, RN15-02) | F04-03 | Should | 1d | Dados do cliente autenticado. Filtrado por tenant_id |
| **T-065** | Atualizar `JwtAuthenticationFilter`: claims `modules[]` + `business_unit_ids[]` no `TenantContext`. Placeholder "FBSO Platform" (RN16-01, RN16-02). Depende de T-136 (DT-096) | F04-04 | Must | 1.5d | JWT retorna modules[]. App Switcher viável com 1 módulo |
| **T-066** | `GET /api/v1/auth/me`: retorna dados do usuário logado (id, name, email, role, business_unit_ids, modules[], tenant_id, status onboarding). Sem `@RequiresPermission` | F04-04 | Must | 1d | GET /auth/me funcional. Stateless (dados do token). 401 se sem token |
| **T-067** | Testes unitários M5: `OnboardingService`, `DashboardClientService`, fluxo auth (login, recovery, rate-limit) | F04-01 a F04-04 | Must | 1.5d | ≥ 80%. Rate limit testado (5 tentativas, bloqueio, reset 15min) |
| **T-068** | Testes integração M5: fluxo onboarding completo (PENDING→ACTIVE). Rate limiting PostgreSQL real. Teste expiração link reset | F04-01, F04-02 | Must | 2d | Onboarding completo testado. Tenant → ACTIVE |

**Total:** 36 tarefas (6 Frente 0 + 10 Frente 1 + 8 Frente 2 + 12 Frente 3) · ~26.5 dias-homem (24 Must + 2 Should + 10 Could)

---

## 📦 Features Entregues

| Feature | Descrição | RNs Cobertas | Prio. |
|:---|:---|:---|:---:|
| **F04-01** | Login e Autenticação | RN13-01, RN13-02, RN13-03 | Must |
| **F04-02** | Onboarding Guiado | RN14-01, RN14-02, RN14-03, RN14-04 | Must |
| **F04-03** | Dashboard do Cliente | RN15-01, RN15-02 | Should |
| **F04-04** | App Switcher | RN16-01, RN16-02 | Must |

---

## ✅ Definition of Done (Sprint-Level)

- [ ] Login Keycloak funcional (Authorization Code Flow)
- [ ] Recuperação de senha com link (expira 1h)
- [ ] Rate limiting: 5 tentativas → bloqueio 15min
- [ ] Onboarding 4 passos: ordem validada, não pula etapas
- [ ] Primeira BU vira Matriz (parent_id=NULL)
- [ ] Tenant transita PENDING_ONBOARDING → ACTIVE ao concluir
- [ ] Dashboard cliente: cards com unidades, produtos, plano
- [ ] `GET /auth/me` retorna dados do token (stateless)
- [ ] JWT inclui claims `modules[]` e `business_unit_ids[]`

---

## ⚠️ Riscos e Bloqueadores

| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---:|:---|
| Keycloak não disponível para desenvolvimento local | Média | Crítico | **Frente 0 (T-133.DT-095):** Criar docker-compose.yml com Keycloak 26 + PostgreSQL 17 + MailHog. Realm exportado como `realm-config.json` |
| Rate limiting com Caffeine não funcionar em cluster (stateless) | Baixa | Médio | Fase 0: single instance. **Decisão DT-110/T-144:** Filter + Caffeine. Trigger para Redis: `INSTANCE_COUNT > 1` em produção |
| Onboarding com estado inconsistente (passo 2 falha, passo 1 salvo) | Média | Alto | `@Transactional` no `OnboardingService`. Rollback se qualquer passo falhar. **T-145.DT-124:** Diagrama de estados documentado |
| Flyway 2 majors atrás (10.22.0) — migrations V007+ podem ser incompatíveis | Alta | Médio | **Frente 0 (T-134.DT-045):** Bump Flyway → 12.11.0 antes de criar V007+ |
| CVE-2026-42198 ativa no PostgreSQL JDBC Driver (CVSS 7.5) | Alta | Alto | **Frente 0 (T-135.DT-068):** Bump driver → 42.7.11 |
| OAuth2 Client não configurado — login Authorization Code Flow não funciona | Alta | Crítico | **Frente 0 (T-137.DT-099 + T-138.DT-100):** Adicionar dependência + configurar application.yml |

---

## 🔗 Dependências

- **Pré-requisitos:** Sprint 2 (JwtAuthenticationFilter). Sprint 3 (TenantRepository). Sprint 4 (UserRepository). Keycloak provisionado.
- **Sucessor:** Sprint 6 (BUs e Catálogo) — depende de BusinessUnit criada durante onboarding.

---

## 📊 Métricas da Sprint

| Métrica | Meta |
|:---|:---:|
| Tasks Frente 0 (Bloqueantes) | 6/6 ✅ |
| Tasks Frente 1 (Recomendados) | 0/10 |
| Tasks Frente 2 (Desejáveis) | 0/8 |
| Tasks Frente 3 (Features) | 0/12 |
| **Total Tasks** | **7/36 (19%)** |
| Must Have | 24 |
| Should Have | 2 (F04-03) |
| Could Have | 10 (Frentes 1-2) |
| Débitos técnicos tratados | 24 |
| Cenários de teste | 33 (28 originais + 5 adicionados DT-106) |

---

🤖 *Gerado a partir de TASKS.md v2.0. Keycloak é dependência crítica — provisionar antes de começar.*
