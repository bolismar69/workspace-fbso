# SPRINT-CARD: Sprint 5 — Portal do Cliente e Onboarding

- **Sprint:** 5 de 7
- **Marco:** M5 (EP-04a)
- **Datas:** 15/09/2026 → 30/09/2026
- **Duração:** 11 dias úteis
- **Responsável:** A definir (time técnico v1.5 — 10 papéis)
- **Documentos-mestre:** [TASKS.md](../../TASKS.md) v3.5 · [SPECS.md](../../SPECS.md) v2.5 · [TECHNICAL-TEAM-MAP.md](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/TECHNICAL-TEAM-MAP.md) v1.5
- **Revisão:** 22/07/2026 — Reavaliação com time técnico v1.5 (10 papéis, 4 tasks frontend adicionadas)

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

| Frente | Escopo | Tasks | Débitos | Estimativa | Papéis-Chave | Quando |
|:---|:---|:---:|:---:|:---:|:---|:---|
| **Frente 0** | 🔴 Bloqueantes (pré-sprint) | 6 | DT-095, DT-045, DT-068, DT-096, DT-099, DT-100 | ~6h (≈1d) | Dev Backend (Agente IA) | ✅ Concluída 17/07/2026 |
| **Frente 1** | 🟡 Recomendados | 10 | DT-023, DT-097, DT-098, DT-107, DT-108, DT-110, DT-124, DT-121, DT-106, DT-102 | ~17.5h (≈3d serial, ≈1.5d paralelo) | Arquiteto, Dev BD, IAM, QA, Dev Backend, Tech Lead | Durante a sprint |
| **Frente 2** | 🔵 Desejáveis (opcional) | 8 | DT-086, DT-089, DT-090, DT-092, DT-093, DT-101, DT-112, DT-113 | ~4.5h (≈0.5d) | Dev Backend, DevOps, Dev BD, Tech Lead | ✅ Concluída 23/07/2026 |
| **Frente 3** | 🎯 Features Sprint 5 | 16 | — (T-057 a T-068 + T-157 a T-160) | ~34d-h (≈8d paralelo) | IAM, Dev Backend, Dev Full-Stack, Dev BD, Dev Frontend, QA | 🔄 12/12 backend ✅ · 0/4 frontend ⬜ |

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

### Frente 1 — Recomendados ✅ CONCLUÍDA | 23/07/2026

| ID | Tarefa | DT | Prio. | Est. | Status | Papel | Critério DONE |
|:---|:---|:---|:---:|:---:|:---:|:---|:---|
| **T-139.DT-023** | Implementar `findAllKeyset()` no `BaseRepository` — paginação keyset em vez de offset | DT-023 | Must | 3h | ✅ | Dev BD | Keyset funcional. Testes com >10k registros |
| **T-140.DT-097** | Corrigir contagem de cenários: 21→28 no SPRINT-TEST-SUITE.md e SPRINT-CARD.md | DT-097 | Must | 0.5h | ✅ | Tech Lead | Números consistentes entre documentos |
| **T-141.DT-098** | Conectar `TenantContext.businessUnitIds`/`modules` ao `JwtAuthenticationFilter` (junto com T-136) | DT-098 | Must | — | ✅ NO-OP | — | Getters chamados em T-065 e T-066 — já implementado |
| **T-142.DT-107** | Adicionar campo `is_matrix BOOLEAN NOT NULL DEFAULT false` à entidade `BusinessUnit` via migration V007 | DT-107 | Must | 1h | ✅ | Dev BD | Migration aplicada. Entidade atualizada. T-062 usa flag |
| **T-143.DT-108** | Documentar máquina de estados de `TenantStatus` (transições válidas). Implementar validação no `OnboardingService` | DT-108 | Must | 1.5h | ✅ | Arquiteto | Diagrama de estados no ARCHITECTURE.md §8.1 |
| **T-144.DT-110** | Implementar rate limiting via **Filter** (não @Aspect) + Caffeine. 5 tentativas → bloqueio 15min | DT-110 | Must | 3h | ✅ | Dev Backend | Filter funcional. Mensagem exibe tempo restante. Build OK |
| **T-145.DT-124** | Criar diagrama de estados do onboarding antes de codificar `OnboardingService` | DT-124 | Must | 2h | ✅ | Arquiteto | Diagrama documentado no ARCHITECTURE.md §8.2. 6 edge cases |
| **T-146.DT-121** | Adicionar `@ExceptionHandler(AuthenticationException.class)` → 401 RFC 7807 no `GlobalExceptionHandler` | DT-121 | Must | 0.5h | ✅ | Dev Backend | 401 JSON padronizado. Token inválido/expirado não gera 500 |
| **T-147.DT-106** | Adicionar 5 cenários de teste ausentes: timeout sessão, complexidade senha, passo 3, segurança F04-03 (×2) | DT-106 | Must | 2h | ✅ | QA Engineer | 33 cenários no total (28+5). Suite atualizada |
| **T-148.DT-102** | Consolidar dupla decodificação JWT via `JwtAuthenticationConverter` (junto com DT-076) | DT-102 | Should | 2h | ✅ | Especialista IAM | 1 decodificação por request. Converter registrado no SecurityConfig |

### Frente 2 — Desejáveis ✅ CONCLUÍDA | 23/07/2026

| ID | Tarefa | DT | Prio. | Est. | Status | Papel | Critério DONE |
|:---|:---|:---|:---:|:---:|:---:|:---|:---|
| **T-149.DT-086** | Extrair helper `AuditFieldsRowMapper` — eliminar duplicação de 6 campos em 5 RowMappers | DT-086 | Could | 1.5h | ✅ | Dev Backend | Helper extraído. 4 RowMappers usam helper |
| **T-150.DT-089** | Injete `ObjectMapper` do Spring no `AuditAspect` (remover `new ObjectMapper()`) | DT-089 | Could | 0.5h | ✅ | Dev Backend | Injeção por construtor. Serialização consistente |
| **T-151.DT-090** | Substituir `OffsetDateTime.now()` → `OffsetDateTime.now(ZoneOffset.UTC)` — 13 ocorrências em 9 arquivos | DT-090 | Could | 0.5h | ✅ | Dev BD | Timestamps UTC em todos os ambientes |
| **T-152.DT-092** | Bump springdoc-openapi 2.8.8→2.8.16 no pom.xml | DT-092 | Could | 0.5h | ✅ | DevOps | Swagger UI funcional. Sem breaking changes |
| **T-153.DT-093** | Externalizar CORS origins para `application.yml` (`app.cors.allowed-origins`) | DT-093 | Could | 0.5h | ✅ | DevOps | Origens configuráveis por ambiente |
| **T-154.DT-101** | Atualizar mitigação de riscos no SPRINT-CARD.md (referenciar DT-095 para docker-compose) | DT-101 | Could | 0.5h | ✅ | Tech Lead | Riscos atualizados com ✅ |
| **T-155.DT-112** | Atualizar header do SPECS.md: "Próximo: Sprint 5 Frente 3 — Features" | DT-112 | Could | 0.25h | ✅ | Tech Lead | Header correto |
| **T-156.DT-113** | Recalcular progresso no TASKS.md: 105/167 (63%) + Frentes 0-1-2 ✅ | DT-113 | Could | 0.25h | ✅ | Tech Lead | Progresso preciso |

### Frente 3 — Features da Sprint 5

> ⚡ **Paralelização com time v1.5:** Backend (T-057..T-068) e Frontend (T-157..T-160) executam simultaneamente. IAM Specialist é o caminho crítico (T-057 → T-058 → T-065 → T-066). Frontend consome APIs via MSW mock até backend estar pronto. QA prepara cenários durante desenvolvimento.

#### Frente 3a — Backend (Admin API) ✅ CONCLUÍDA | 23/07/2026

| ID | Tarefa | Feature | Prio. | Est. | Status | Papel | Critério DONE |
|:---|:---|:---|:---:|:---:|:---|:---|
| **T-057** ✅ | Configurar Keycloak realm `fbso-platform` + client Authorization Code Flow. Mapear claims (tenant_id, roles, business_unit_ids, modules). Exportar `realm-config.json` | F04-01 | Must | 2d | Especialista IAM | Realm funcional. JWT com claims corretas. Config versionada |
| **T-058** ✅ | `POST /api/v1/auth/login` (delega Keycloak). Recuperação: `POST /auth/forgot-password` (expira 1h — RN13-03), `POST /auth/reset-password` (complexidade RN13-01) | F04-01 | Must | 2d | Dev Backend | Login funcional. Link reset 1h. Senha 8+ chars, letra+número. Sessão 60min inatividade (RN13-02) |
| **T-059** ✅ | Rate limiting: 5 tentativas → bloqueio 15min (US-039, RN13-02). Implementar via Filter + Caffeine (decisão DT-110/T-144) | F04-01 | Must | 1.5d | Dev Backend | 5 falhas → bloqueado. Mensagem exibe tempo restante |
| **T-060** ✅ | `OnboardingService`: 4 passos obrigatórios (US-040 a US-044). Ordem obrigatória (RN14-01). Primeira BU = Matriz (RN14-02). Tenant → ACTIVE ao concluir (RN14-04). Retomável. Diagrama de estados (DT-124/T-145) | F04-02 | Must | 2d | Dev Full-Stack | 4 passos. Ordem validada. Status salvo permite retomar |
| **T-061** ✅ | `OnboardingController`: `GET /onboarding/status`, `PATCH /step-1`, `POST /step-2`, `POST /complete`. `@RequiresPermission` | F04-02 | Must | 1.5d | Dev Backend | 4 endpoints. Step-2 valida CNPJ. Complete só se todos passos OK (RN14-03) |
| **T-062** ✅ | Criação primeira BU como Matriz (step-2). Reutilizar `BusinessUnitService.create()` com flag `isMatrix=true` (DT-107/T-142). CNPJ validado | F04-02 | Must | 1d | Dev Backend + Dev BD | BU Matriz com is_matrix=true, parent_id=NULL. CNPJ válido |
| **T-063** ✅ | `DashboardClientService`: cards — unidades ativas, produtos catálogo, plano contratado, notificações. Cards clicáveis | F04-03 | Should | 1.5d | Dev Full-Stack | Cards com dados resumidos. Link para área |
| **T-064** ✅ | `DashboardClientController`: `GET /dashboard/client/summary`, `/notifications`. `@RequiresPermission`. Dashboard genérico Fase 0 (RN15-01, RN15-02) | F04-03 | Should | 1d | Dev Backend | Dados do cliente autenticado. Filtrado por tenant_id |
| **T-065** ✅ | Atualizar `JwtAuthenticationFilter`: claims `modules[]` + `business_unit_ids[]` no `TenantContext`. Placeholder "FBSO Platform" (RN16-01, RN16-02). Depende de T-136 (DT-096) | F04-04 | Must | 1.5d | Especialista IAM | JWT retorna modules[]. App Switcher viável com 1 módulo |
| **T-066** ✅ | `GET /api/v1/auth/me`: retorna dados do usuário logado (id, name, email, role, business_unit_ids, modules[], tenant_id, status onboarding). Sem `@RequiresPermission` | F04-04 | Must | 1d | Dev Backend | GET /auth/me funcional. Stateless (dados do token). 401 se sem token |
| **T-067** ✅ | Testes unitários M5: `OnboardingService`, `DashboardClientService`, fluxo auth (login, recovery, rate-limit) | F04-01 a F04-04 | Must | 1.5d | QA Engineer | ≥ 80%. Rate limit testado (5 tentativas, bloqueio, reset 15min) |
| **T-068** ✅ | Testes integração M5: fluxo onboarding completo (PENDING→ACTIVE). Rate limiting PostgreSQL real. Teste expiração link reset | F04-01, F04-02 | Must | 2d | QA Engineer | Onboarding completo testado. Tenant → ACTIVE |

#### Frente 3b — Frontend (web_app-fbso-platform-portal) 🆕

> Stack: Next.js App Router + TypeScript + Tailwind CSS + next-auth (ADR-03). Projeto em `/frontend/javascript/react/web_apps/web_app-fbso-platform-portal/`.

| ID | Tarefa | Feature | Prio. | Est. | Papel | Critério DONE |
|:---|:---|:---|:---:|:---:|:---|:---|
| **T-157** | Bootstrap `web_app-fbso-platform-portal`: Next.js App Router + TypeScript + Tailwind CSS. Estrutura: `app/(auth)/`, `app/(onboarding)/`, `app/(portal)/`, `components/`, `lib/`. Configurar MSW mock handlers para dev paralelo. ESLint + Prettier | F04-01, F04-03, F04-04 | Must | 2d | Dev Frontend | `npm run dev` funcional. Rotas App Router corretas. MSW handlers ativos. ESLint/Prettier zero warnings |
| **T-158** | Auth UI: Login via redirect Keycloak (next-auth). Callback handler + middleware proteção rotas. Forgot/reset password forms (RN13-01). Estados: loading skeleton, erro RFC 7807, sucesso | F04-01 | Must | 2d | Dev Frontend | Login funcional com redirect Keycloak. Form reset com validação client+server. Sessão 60min (RN13-02) |
| **T-159** | Onboarding wizard UI: 4 steps sequenciais com barra progresso. Step-2: CNPJ + regime tributário. Validação por step — não pula etapas (RN14-01). Retomável. Mobile-first. Toast notificações | F04-02 | Must | 2.5d | Dev Frontend | 4 steps navegáveis. Progresso 25→50→75→100%. Validação CNPJ. Retoma de onde parou. Responsivo |
| **T-160** | Dashboard cliente: 4 MetricCards (Unidades, Produtos, Plano, Notificações). API via `lib/api-client.ts`. Loading skeleton. Empty state. Cards clicáveis → navegação. Grid responsivo (4→2→1 col) | F04-03 | Should | 1.5d | Dev Frontend | 4 cards com dados da API. Skeleton durante fetch. Responsivo. Notificações com badge + link |

**Total:** 40 tarefas (6 Frente 0 + 10 Frente 1 + 8 Frente 2 + 16 Frente 3) · ~34 dias-homem · **36/40 (90%) concluído**

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

- [x] Login Keycloak funcional (Authorization Code Flow)
- [x] Recuperação de senha com link (expira 1h)
- [x] Rate limiting: 5 tentativas → bloqueio 15min
- [x] Onboarding 4 passos: ordem validada, não pula etapas
- [x] Primeira BU vira Matriz (parent_id=NULL)
- [x] Tenant transita PENDING_ONBOARDING → ACTIVE ao concluir
- [x] Dashboard cliente: cards com unidades, produtos, plano
- [x] `GET /auth/me` retorna dados do token (stateless)
- [x] JWT inclui claims `modules[]` e `business_unit_ids[]`
- [ ] 🆕 **Login UI:** redirect Keycloak via next-auth, middleware proteção rotas
- [ ] 🆕 **Onboarding wizard UI:** 4 steps sequenciais, barra progresso, mobile-first
- [ ] 🆕 **Dashboard cliente UI:** 4 cards responsivos, loading skeleton, empty state
- [ ] 🆕 **MSW mock handlers:** desenvolvimento frontend paralelo ao backend

---

## ⚠️ Riscos e Bloqueadores

| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---:|:---|
| Keycloak não disponível para desenvolvimento local | Média | Crítico | ✅ **Frente 0 (T-133.DT-095):** docker-compose.yml criado com Keycloak 26 + PostgreSQL 17 + MailHog. Realm exportado como `realm-config.json` |
| Rate limiting com Caffeine não funcionar em cluster (stateless) | Baixa | Médio | ✅ **Frente 1 (T-144.DT-110):** RateLimitFilter implementado. Caffeine Cache com 5 tentativas → bloqueio 15min. Trigger para Redis: `INSTANCE_COUNT > 1` em produção |
| Onboarding com estado inconsistente (passo 2 falha, passo 1 salvo) | Média | Alto | ✅ **Frente 1 (T-143.DT-108 + T-145.DT-124):** Máquinas de estado documentadas no ARCHITECTURE.md §8.1-8.2. `@Transactional` no `OnboardingService` (T-060). 6 edge cases documentados |
| Flyway 2 majors atrás (10.22.0) — migrations V007+ podem ser incompatíveis | — | — | ✅ **Frente 0 (T-134.DT-045):** Flyway 12.11.0. Migrations V001-V007 reaplicadas sem erro. **Risco eliminado** |
| CVE-2026-42198 ativa no PostgreSQL JDBC Driver (CVSS 7.5) | — | — | ✅ **Frente 0 (T-135.DT-068):** Driver 42.7.11. **Risco eliminado** |
| OAuth2 Client não configurado — login Authorization Code Flow não funciona | — | — | ✅ **Frente 0 (T-137.DT-099 + T-138.DT-100):** OAuth2 Client configurado. 2 SecurityFilterChain beans. **Risco eliminado** |
| JWT dupla decodificação — CPU desperdiçada por request | Baixa | Baixo | ✅ **Frente 1 (T-148.DT-102):** `FbsoJwtAuthenticationConverter` registrado. 1 decodificação por request |
| 🆕 **Time técnico não preenchido** — 10 papéis com `<TODO>` no TECHNICAL-TEAM-MAP.md | **Alta** | **Crítico** | Preencher nomes antes do início da Frente 3 (15/09). Risco R2 do Project Charter permanece com severidade Crítica |
| 🆕 **Frontend sem design system** — `web_app-fbso-platform-portal` não tem componentes base | Média | Médio | Tailwind CSS provê design tokens. MSW mock permite dev paralelo. Componentes começam simples (evoluir na Sprint 6) |
| 🆕 **IAM Specialist 4h/dia pode ser insuficiente** — T-057 + T-065 = 3.5d (28h = 7 dias wall-clock) | Média | Alto | Iniciar T-057 durante a Frente 1 (antes da janela oficial da sprint). Considerar bump para 8h/dia na Sprint 5 |

---

## 🔗 Dependências

- **Pré-requisitos:** Sprint 2 (JwtAuthenticationFilter). Sprint 3 (TenantRepository). Sprint 4 (UserRepository). Keycloak provisionado.
- **Sucessor:** Sprint 6 (BUs e Catálogo) — depende de BusinessUnit criada durante onboarding.

---

## 📊 Métricas da Sprint

| Métrica | Meta |
|:---|:---:|
| Tasks Frente 0 (Bloqueantes) | 6/6 ✅ |
| Tasks Frente 1 (Recomendados) | 10/10 ✅ |
| Tasks Frente 2 (Desejáveis) | 8/8 ✅ |
| Tasks Frente 3a — Backend | 12/12 ✅ |
| Tasks Frente 3b — Frontend 🆕 | 0/4 |
| **Total Tasks** | **36/40 (90%)** |
| Must Have | 27 → 21 concluídos |
| Should Have | 3 (F04-03: T-063, T-064 ✅, T-160 pendente) |
| Could Have | 2 (Frente 2 — já executada ✅) |
| Débitos técnicos tratados | 24 |
| Cenários de teste | 33 (28 originais + 5 DT-106) |
| Time alocado | 10 papéis × 68h/dia (748h total sprint) |

---

🤖 *Gerado a partir de TASKS.md v3.5. Revisado em 22/07/2026 com time técnico v1.5 (10 papéis). Keycloak é dependência crítica — IAM Specialist deve iniciar T-057 antes da janela oficial da sprint. Frontend (T-157..T-160) em paralelo com backend via MSW mocks.*
