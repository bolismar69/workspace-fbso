# SPRINT-DEVELOPMENT-PLANNING-Frente-3.md — Plano de Desenvolvimento: Sprint 5 — Frente 3

- **Solução:** `ms-fbso-platform-admin` (+ `web_app-fbso-platform-portal`)
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 5 de 7 — sprint-05-portal-cliente
- **Frente:** Frente 3 — Features do Portal do Cliente (🎯 Features)
- **Stack Backend:** Java 25 · Spring Boot 3.5.14 · PostgreSQL 17 · Keycloak 26 · Flyway 12.11.0 · Caffeine 3.2.4
- **Stack Frontend:** Next.js App Router · TypeScript · Tailwind CSS · next-auth · MSW
- **Branch:** `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-05-portal-cliente`
- **Data do planejamento:** 2026-07-23

---

## 1. Visão Geral

- **Objetivo da Frente 3:** Entregar as 4 features do Portal do Cliente (F04-01 a F04-04): autenticação Keycloak OIDC, onboarding guiado em 4 passos, dashboard do cliente com cards, e App Switcher com endpoint `/auth/me`.
- **Sprint Goal:** "Login via Keycloak OIDC funcional com recuperação de senha e rate limiting. Onboarding guiado em 4 passos — primeira BU vira Matriz, tenant transita para ACTIVE ao concluir. Dashboard do cliente com cards de unidades, produtos e plano. Endpoint /auth/me retorna dados do usuário logado."
- **Tasks a implementar:** 16 (12 backend + 4 frontend)
  - **Backend (Must):** 10 tasks (T-057..T-066) — 16.5d estimados
  - **Backend (Should):** 2 tasks (T-063, T-064) — 2.5d estimados
  - **Testes (Must):** 2 tasks (T-067, T-068) — 3.5d estimados
  - **Frontend (Must+Should):** 4 tasks (T-157..T-160) — 8d estimados (projeto separado)
- **Ordem de execução:** Mista — Backend e Frontend em paralelo. IAM Specialist é o caminho crítico. Frontend usa MSW mocks até APIs estarem prontas.
- **Estimativa total:** ~30.5d-h (≈8d wall-clock com time paralelo)
- **Features entregues:** F04-01 (Login/Auth) · F04-02 (Onboarding) · F04-03 (Dashboard Cliente) · F04-04 (App Switcher)

---

## 2. Dependências entre Tasks

```
CAMINHO CRÍTICO (IAM Specialist — 7d):
T-057 (Keycloak realm) ──→ T-058 (Auth endpoints) ──→ T-065 (JWT claims) ──→ T-066 (/auth/me)
                                                              │
                                                              └── T-059 (Rate limiting — integra Filter existente)


FLUXO ONBOARDING (Dev Full-Stack + Dev Backend — 5.5d):
T-060 (OnboardingService) ──→ T-061 (OnboardingController)
        │                           │
        └── T-062 (BU Matriz) ──────┘


FLUXO DASHBOARD (Should — 2.5d):
T-063 (DashboardClientService) ──→ T-064 (DashboardClientController)


TESTES (QA Engineer — 3.5d):
T-067 (Unitários) ──→ T-068 (Integração)
   └── depende de T-058..T-066 concluídas


FRONTEND (Dev Frontend — 8d, PARALELO ao backend):
T-157 (Bootstrap) ──→ T-158 (Auth UI) ──→ T-159 (Onboarding UI)
                    │
                    └── T-160 (Dashboard UI)
```

### Dependências das Frentes 0-1-2 (já concluídas ✅):

| Task Frente 3 | Depende de | Status | O que foi entregue |
|:---|:---|:---:|:---|
| T-057 | T-133 (docker-compose) | ✅ | Keycloak 26 + PG 17 + MailHog |
| T-058 | T-137+T-138 (OAuth2 Client) | ✅ | SecurityConfig com 2 filter chains |
| T-059 | T-144 (RateLimitFilter) | ✅ | Filter + Caffeine, 5 tentativas → 15min |
| T-060 | T-143+T-145 (máquinas de estado) | ✅ | ARCHITECTURE.md §8.1-8.2 |
| T-062 | T-142 (is_matrix + V007) | ✅ | Campo `is_matrix BOOLEAN` na tabela |
| T-065 | T-136+T-148 (JWT claims + Converter) | ✅ | FbsoJwtAuthenticationConverter |
| T-066 | T-065 + T-136 | ✅ | Claims disponíveis no TenantContext |
| T-146 | ExceptionHandler 401 | ✅ | AuthenticationException → 401 RFC 7807 |

---

## 3. Análise PonteTail (7 Rungs) por Task

### Resumo da Análise

| Task | R1-YAGNI | R2-JáExiste | R3-Stdlib | R4-DepExistente | R5-Padrão | R6-Simples | R7-Mínimo | Veredito |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---|
| T-057 | ✅ | ❌ | ✅ Keycloak Admin CLI | ✅ docker-compose | ✅ realm-config.json | ✅ 1 arquivo JSON | ✅ | Prosseguir |
| T-058 | ✅ | ❌ (parcial — OAuth2 Client existe) | ✅ Spring Security OAuth2 | ✅ spring-boot-starter-oauth2-client | ✅ SecurityConfig | ✅ 1 controller | ✅ | Prosseguir |
| T-059 | ✅ | ❌ (parcial — Filter existe) | ✅ Caffeine 3.2.4 | ✅ RateLimitFilter + Caffeine | ✅ Filter pattern | ✅ Integrar existente | ✅ | Prosseguir |
| T-060 | ✅ | ❌ | ✅ Spring @Transactional | ✅ TenantService, BusinessUnitService | ✅ Service pattern | ✅ 4 métodos | ✅ | Prosseguir |
| T-061 | ✅ | ❌ | ✅ Spring Web | ✅ spring-boot-starter-web | ✅ Controller pattern | ✅ 4 endpoints | ✅ | Prosseguir |
| T-062 | ✅ | ❌ (parcial — BU entity existe) | ✅ JDBC | ✅ BusinessUnitService.create() | ✅ Service pattern | ✅ Flag isMatrix=true | ✅ | Prosseguir |
| T-063 | ✅ | ❌ | ✅ JDBC | ✅ JdbcTemplate, TenantContext | ✅ Service pattern | ✅ Queries agregadas | ✅ | Prosseguir |
| T-064 | ✅ | ❌ | ✅ Spring Web | ✅ spring-boot-starter-web | ✅ Controller pattern | ✅ 2 endpoints | ✅ | Prosseguir |
| T-065 | ✅ | ❌ (parcial — Filter existe) | ✅ JWT Nimbus | ✅ JwtAuthenticationFilter, TenantContext | ✅ Security filter pattern | ✅ Atualizar claims | ✅ | Prosseguir |
| T-066 | ✅ | ❌ | ✅ Spring Web | ✅ TenantContext, JwtAuthenticationFilter | ✅ Controller pattern | ✅ 1 endpoint GET | ✅ | Prosseguir |
| T-067 | ✅ | ❌ | ✅ JUnit 5 + Mockito | ✅ Mockito, AssertJ | ✅ AAA pattern | ✅ Focado em services | ✅ | Prosseguir |
| T-068 | ✅ | ❌ | ✅ Testcontainers | ✅ PostgreSQL testcontainers | ✅ Given-When-Then | ✅ Fluxos críticos | ✅ | Prosseguir |
| T-157 | ✅ | ❌ | ✅ Next.js create-app | ✅ Node 20+, npx | ✅ App Router | ✅ Estrutura base | ✅ | Prosseguir |
| T-158 | ✅ | ❌ | ✅ next-auth | ✅ Next.js App Router | ✅ next-auth patterns | ✅ 3 páginas | ✅ | Prosseguir |
| T-159 | ✅ | ❌ | ✅ React hooks | ✅ Tailwind CSS, next-auth | ✅ Wizard pattern | ✅ 4 steps | ✅ | Prosseguir |
| T-160 | ✅ | ❌ | ✅ React + Tailwind | ✅ Next.js, Tailwind | ✅ Card grid pattern | ✅ 4 cards | ✅ | Prosseguir |

---

## 4. Plano por Task — Backend (Frente 3a)

---

### T-057 — Configurar Keycloak realm `fbso-platform`

- **Critério DONE:** Realm funcional. JWT com claims corretas (tenant_id, roles, business_unit_ids, modules). Config versionada em `realm-config.json`.
- **Estimativa:** 2d (16h)
- **Papel:** Especialista IAM
- **Prioridade:** Must — CAMINHO CRÍTICO #1
- **RACI:** Especialista IAM (R), Dev Backend (C), Arquiteto (C), Tech Lead (A)

**Abordagem:**
Criar o arquivo `keycloak/realm-config.json` com a configuração completa do realm `fbso-platform`. O docker-compose.yml já está configurado para importar realms via `--import-realm` e monta o arquivo como volume (`./keycloak/realm-config.json:/opt/keycloak/data/import/realm-config.json:ro`).

**Configuração necessária no realm:**
1. **Realm:** `fbso-platform` — enabled, registration disabled (apenas admin cria usuários)
2. **Client:** `fbso-platform-admin` — public (`confidential`), Authorization Code Flow, redirect URIs (localhost:3000, app.fbso.org), scopes (openid, profile, email)
3. **Roles:** `ADMIN_TENANT`, `MANAGER_BU`, `OPERATOR_BU`, `AUDITOR` (4 roles da matriz RN10-01)
4. **Client Scopes / Mappers:** Mapear atributos customizados do usuário para claims JWT:
   - `tenant_id` → claim `tenant_id` (String → UUID)
   - `business_unit_ids` → claim `business_unit_ids` (List<String>)
   - `modules` → claim `modules` (List<String>)
5. **Password Policy:** 8+ caracteres, letra + número (RN13-01)
6. **Token Settings:** Access Token Lifespan = 60 min (RN13-02 — sessão 60min)
7. **Required Actions:** `forgot-password` / `reset-password` configurado

**Arquivos a criar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `keycloak/realm-config.json` | 🆕 | Realm export com client, roles, mappers, password policy |

**Arquivos a modificar:** Nenhum (configuração standalone)

**Dependências:** T-133 (docker-compose ✅)

**Riscos:**
- **Alto:** O formato do realm-config.json do Keycloak 26 pode diferir de versões anteriores. Usar export via Admin Console/CLI como referência.
- **Médio:** Mapeamento de claims customizadas via user attributes requer configuração precisa. Testar com JWT decodificado (jwt.io) para validar.
- **Mitigação:** Após criar o realm-config.json, executar `docker compose up -d keycloak` e verificar: (a) Keycloak Admin Console acessível em localhost:8081, (b) OpenID Configuration disponível em `/realms/fbso-platform/.well-known/openid-configuration`.

**Skills aplicáveis:** `304-frameworks-spring-boot-security`, `124-java-secure-coding`

---

### T-058 — Endpoints de Autenticação: `POST /auth/login`, `/auth/forgot-password`, `/auth/reset-password`

- **Critério DONE:** Login funcional (delega Keycloak). Link reset expira 1h (RN13-03). Senha validada (8+ chars, letra+número — RN13-01). Sessão 60min inatividade (RN13-02).
- **Estimativa:** 2d (16h)
- **Papel:** Dev Backend
- **Prioridade:** Must — CAMINHO CRÍTICO #2
- **RACI:** Dev Backend (R), Especialista IAM (C), Arquiteto (C), Tech Lead (A)

**Abordagem:**
Criar `AuthController` e `AuthService` que orquestram a autenticação. O fluxo Authorization Code Flow é delegado ao Keycloak (frontend redireciona → Keycloak → callback). O backend expõe:

1. **`POST /api/v1/auth/login`** — Não implementa lógica própria. O login é via redirect OAuth2 (já configurado no SecurityConfig Frente 0). Este endpoint existe como entrypoint documentado na API, retornando 200 com redirect info ou delegando ao `oauth2LoginFilterChain`.

2. **`POST /api/v1/auth/forgot-password`** — Recebe `{ email }`. Chama Keycloak Admin REST API (`PUT /admin/realms/fbso-platform/users/{id}/execute-actions-email` com action `UPDATE_PASSWORD`). O link enviado pelo Keycloak expira em 1h (configurado no realm — RN13-03).

3. **`POST /api/v1/auth/reset-password`** — Recebe `{ token, newPassword }`. Valida complexidade client-side (RN13-01: 8+ chars, letra+número). Chama Keycloak Admin REST API para reset. O token é o `keycloak_session_state` ou token de action do Keycloak.

**Decisão de design:** Os endpoints `/auth/forgot-password` e `/auth/reset-password` atuam como **proxy/facade** para a Keycloak Admin REST API. Isso mantém o backend como ponto único de entrada e evita expor a Admin API do Keycloak diretamente ao frontend.

**Arquivos a criar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../controller/AuthController.java` | 🆕 | 3 endpoints: login, forgot-password, reset-password |
| `src/main/java/.../service/AuthService.java` | 🆕 | Lógica: delega login OAuth2, proxy Keycloak Admin API para recovery |
| `src/main/java/.../dto/request/LoginRequest.java` | 🆕 | DTO: email (opcional — login via redirect) |
| `src/main/java/.../dto/request/ForgotPasswordRequest.java` | 🆕 | DTO: email |
| `src/main/java/.../dto/request/ResetPasswordRequest.java` | 🆕 | DTO: token + newPassword |
| `src/main/java/.../dto/response/AuthResponse.java` | 🆕 | DTO: redirectUrl, token info |

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../config/SecurityConfig.java` | 🔄 | Garantir que `/auth/**` esteja no `oauth2LoginFilterChain` securityMatcher |

**Dependências:** T-057 (Keycloak realm funcional — claims mapeadas) · T-137+T-138 (OAuth2 Client ✅)

**Riscos:**
- **Alto:** Keycloak Admin REST API requer token de serviço (service account) com permissões `manage-users`, `manage-realm`. Configurar client `fbso-platform-admin` com Service Account Enabled + roles apropriadas.
- **Médio:** A comunicação backend→Keycloak é sensível a latência. Adicionar timeout e retry (Spring RestClient/RestTemplate).
- **Mitigação:** Testar fluxo completo com docker-compose: criar usuário no Keycloak → POST /auth/forgot-password → receber email no MailHog (localhost:8025) → POST /auth/reset-password → login funcional.

**Skills aplicáveis:** `302-frameworks-spring-boot-rest`, `304-frameworks-spring-boot-security`

---

### T-059 — Rate Limiting: integrar `RateLimitFilter` ao endpoint de login

- **Critério DONE:** 5 falhas → bloqueio 15min. Mensagem exibe tempo restante.
- **Estimativa:** 1.5d (12h)
- **Papel:** Dev Backend
- **Prioridade:** Must
- **RACI:** Dev Backend (R), Arquiteto (C), Tech Lead (A)

**Abordagem:**
O `RateLimitFilter` já existe (criado na Frente 1 — T-144.DT-110) e está registrado no `SecurityConfig.apiFilterChain()` via `addFilterBefore(rateLimitFilter(), JwtAuthenticationFilter.class)`. A integração consiste em:

1. **Confirmar que o Filter intercepta `POST /api/v1/auth/login`** — verificar o `shouldFilter()` method (herdado de `OncePerRequestFilter`)
2. **Extrair username do body** — Como o Filter lê o InputStream, usar `ContentCachingRequestWrapper` para permitir releitura
3. **Incrementar contador** no Caffeine Cache — chave: `username` (ou IP remoto como fallback)
4. **5ª falha → cache entry com TTL 15min** — `cache.put(username, attempts, 15, TimeUnit.MINUTES)`
5. **Resposta 429:** `Retry-After: 900` header + body RFC 7807 com detail "Tente novamente em X minutos"

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../security/RateLimitFilter.java` | 🔄 | Implementar `shouldFilter()`, extrair username, contador, resposta 429 |
| `src/main/java/.../config/SecurityConfig.java` | 🔄 | Verificar ordem dos filters (RateLimitFilter ANTES do JwtAuthenticationFilter) |

**Dependências:** T-144 (RateLimitFilter ✅) · T-058 (AuthController — endpoint login funcional)

**Riscos:**
- **Médio:** Consumir o `InputStream` do request pode quebrar o Controller (body já foi lido). Solução: `ContentCachingRequestWrapper` no Filter.
- **Baixo:** Caffeine local não escala em cluster. Documentado no SPRINT-CARD.md: trigger para Redis quando `INSTANCE_COUNT > 1`.

**Skills aplicáveis:** `304-frameworks-spring-boot-security`, `124-java-secure-coding`

---

### T-060 — `OnboardingService`: 4 passos obrigatórios

- **Critério DONE:** 4 passos. Ordem validada (RN14-01). Primeira BU = Matriz (RN14-02). Tenant → ACTIVE ao concluir (RN14-04). Estado salvo permite retomar.
- **Estimativa:** 2d (16h)
- **Papel:** Dev Full-Stack
- **Prioridade:** Must
- **RACI:** Dev Full-Stack (R), Dev Backend (C), Dev BD (C), Arquiteto (C), Tech Lead (A)

**Abordagem:**
Implementar serviço que gerencia o fluxo de onboarding conforme máquina de estados documentada no ARCHITECTURE.md §8.2. O estado do onboarding é armazenado no próprio Tenant (campo `onboarding_step` ou similar) para permitir retomada.

**Estados e transições (do ARCHITECTURE.md §8.2):**

| Estado | Próximo | Endpoint | Ação |
|:---|:---|:---|:---|
| NOT_STARTED | STEP1_DONE | `PATCH /onboarding/step-1` | Confirmar dados do tenant (razão social, segmento) |
| STEP1_DONE | STEP2_DONE | `POST /onboarding/step-2` | CNPJ + regime tributário → cria 1ª BU Matriz (T-062) |
| STEP2_DONE | STEP3_DONE | `PATCH /onboarding/step-3` | Configurações fiscais (opcional na Fase 0) |
| STEP3_DONE | COMPLETED | `POST /onboarding/complete` | Validar todos os passos → Tenant → ACTIVE |

**Edge cases implementados (6 cenários do DT-124):**
1. Tenant sai no meio do step 2 → retoma step 2 (último incompleto)
2. Step 2 falha (CNPJ inválido) → permanece em STEP1_DONE
3. Tentar step 3 sem step 2 → 422 "Conclua o passo anterior primeiro" (RN14-01)
4. Retomar após inatividade → sem expiração (Fase 0)
5. Admin resetar onboarding de tenant ativo → 422
6. Step 4 com dados inconsistentes → rollback `@Transactional`

**Métodos do OnboardingService:**

```java
@Service
@Transactional
public class OnboardingService {
    OnboardingStatusResponse getStatus(UUID tenantId);
    OnboardingStatusResponse completeStep1(UUID tenantId, Step1Request req);
    OnboardingStatusResponse completeStep2(UUID tenantId, Step2Request req); // Chama T-062
    OnboardingStatusResponse completeStep3(UUID tenantId, Step3Request req);
    OnboardingStatusResponse complete(UUID tenantId); // Valida + transita ACTIVE
}
```

**Arquivos a criar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../service/OnboardingService.java` | 🆕 | Lógica 4 passos, validação de ordem, transição Tenant→ACTIVE |
| `src/main/java/.../dto/request/OnboardingStep1Request.java` | 🆕 | DTO: razão social, segmento |
| `src/main/java/.../dto/request/OnboardingStep2Request.java` | 🆕 | DTO: cnpj, tax_regime |
| `src/main/java/.../dto/request/OnboardingStep3Request.java` | 🆕 | DTO: configs fiscais (placeholder Fase 0) |
| `src/main/java/.../dto/response/OnboardingStatusResponse.java` | 🆕 | DTO: currentStep, completedSteps, tenantStatus |

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../entity/Tenant.java` | 🔄 | Adicionar campo `onboardingStep` (VARCHAR: NOT_STARTED, STEP1_DONE, ...) |
| `src/main/resources/db/migration/V008__add_onboarding_step_to_tenant.sql` | 🆕 | Migration: ALTER TABLE tenant ADD COLUMN onboarding_step |

**Dependências:** T-143+T-145 (máquinas de estado ✅) · T-062 (criação BU Matriz)

**Riscos:**
- **Alto:** Consistência do estado de onboarding em caso de falha parcial. `@Transactional` cobre rollback no banco, mas chamadas externas (Keycloak, email) podem ficar inconsistentes.
- **Médio:** O campo `onboarding_step` no Tenant funciona para Fase 0 (single-tenant). Se o onboarding evoluir para multi-step complexo, migrar para tabela separada `onboarding_progress`.
- **Mitigação:** Usar `@Transactional` apenas nas operações de banco. Chamadas externas (Keycloak) são best-effort com log WARN em caso de falha.

**Skills aplicáveis:** `302-frameworks-spring-boot-rest`, `121-java-object-oriented-design`, `030-architecture-adr-general`

---

### T-061 — `OnboardingController`: 4 endpoints REST

- **Critério DONE:** 4 endpoints. Step-2 valida CNPJ. Complete só se todos passos OK (RN14-03).
- **Estimativa:** 1.5d (12h)
- **Papel:** Dev Backend
- **Prioridade:** Must
- **RACI:** Dev Backend (R), Dev Full-Stack (C), Tech Lead (A)

**Abordagem:**
Criar `OnboardingController` com 4 endpoints REST conforme SPECS.md §4.1. Cada endpoint delega ao `OnboardingService` e aplica `@RequiresPermission`. O controller é responsável por validação de entrada (Bean Validation) e conversão DTO↔Response.

**Endpoints (SPECS.md §4.1):**

| Método | Path | Descrição | Validações |
|:---|:---|:---|:---|
| `GET` | `/api/v1/onboarding/status` | Status atual do onboarding | Autenticado — tenant_id do JWT |
| `PATCH` | `/api/v1/onboarding/step-1` | Confirmar dados do tenant | `name_corporate` obrigatório, `segment` obrigatório |
| `POST` | `/api/v1/onboarding/step-2` | Cadastrar Matriz (CNPJ + regime) | `cnpj` obrigatório, formato válido. `tax_regime` obrigatório. |
| `POST` | `/api/v1/onboarding/complete` | Finalizar onboarding | Todos os passos concluídos. Tenant em PENDING_ONBOARDING. |

**Arquivos a criar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../controller/OnboardingController.java` | 🆕 | 4 endpoints REST com `@RequiresPermission` |

**Arquivos a modificar:** Nenhum (usa DTOs do T-060 e serviços existentes)

**Dependências:** T-060 (OnboardingService) · T-062 (BusinessUnitService para step-2)

**Riscos:**
- **Baixo:** Controller é fino — apenas validação + delegação. Complexidade está no Service.

**Skills aplicáveis:** `302-frameworks-spring-boot-rest`, `303-frameworks-spring-boot-validation`

---

### T-062 — Criação primeira BU como Matriz (step-2)

- **Critério DONE:** BU Matriz com is_matrix=true, parent_id=NULL. CNPJ validado.
- **Estimativa:** 1d (8h)
- **Papel:** Dev Backend + Dev BD
- **Prioridade:** Must
- **RACI:** Dev Backend (R), Dev BD (C), Tech Lead (A)

**Abordagem:**
Reutilizar `BusinessUnitService.create()` (já existente) com a flag `isMatrix=true` (campo adicionado na Frente 1 — T-142.DT-107/V007). O `OnboardingService.completeStep2()` chama `BusinessUnitService.create()` com os parâmetros do CNPJ e regime tributário informados no onboarding.

**Lógica:**
1. Validar CNPJ (formato + dígito verificador) — usar `CnpjValidator` existente
2. Verificar se já existe BU Matriz para este tenant (RN14-02: "primeira BU = Matriz")
3. Chamar `BusinessUnitService.create()` com:
   - `tenantId` = TenantContext.getTenantId()
   - `cnpj` = request.cnpj()
   - `taxRegime` = request.taxRegime()
   - `hierarchyType` = "MATRIZ"
   - `parentId` = null
   - `isMatrix` = true
4. Retornar `BusinessUnitResponse` com os dados da BU criada

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../service/BusinessUnitService.java` | 🔄 | Adicionar método `createMatrix(CreateMatrixRequest)` ou usar `create()` existente com flag |
| `src/main/java/.../service/OnboardingService.java` | 🔄 | Integrar chamada `BusinessUnitService.create()` no `completeStep2()` |

**Dependências:** T-142 (is_matrix ✅) · T-060 (OnboardingService)

**Riscos:**
- **Baixo:** `BusinessUnitService.create()` já existe e é funcional. Apenas adicionar flag `isMatrix=true`.
- **Mitigação:** Verificar constraint CNPJ único entre ativos do tenant (índice parcial `unique_cnpj_active`).

**Skills aplicáveis:** `311-frameworks-spring-jdbc`, `121-java-object-oriented-design`

---

### T-063 — `DashboardClientService`: cards do dashboard

- **Critério DONE:** Cards com dados resumidos (unidades ativas, produtos catálogo, plano contratado, notificações). Links para área.
- **Estimativa:** 1.5d (12h)
- **Papel:** Dev Full-Stack
- **Prioridade:** Should
- **RACI:** Dev Full-Stack (R), Dev Backend (C), Tech Lead (A)

**Abordagem:**
Criar serviço que agrega dados do tenant autenticado para o dashboard do cliente. O serviço consulta o banco filtrando por `tenant_id` do `TenantContext`. As queries são agregadas e otimizadas para baixa latência.

**Cards do dashboard:**
1. **Unidades Ativas:** `SELECT COUNT(*) FROM business_unit WHERE tenant_id = ? AND deleted_dt IS NULL AND status = 'ACTIVE'` → link para `/business-units`
2. **Produtos no Catálogo:** `SELECT COUNT(*) FROM product_service WHERE business_unit_id IN (SELECT id FROM business_unit WHERE tenant_id = ? AND deleted_dt IS NULL) AND deleted_dt IS NULL AND status = 'ACTIVE'` → link para `/products`
3. **Plano Contratado:** `SELECT p.name, s.status FROM subscription s JOIN plan p ON s.plan_id = p.id WHERE s.tenant_id = ? AND s.status = 'ACTIVE' AND s.deleted_dt IS NULL` → link para `/plan`
4. **Notificações:** Placeholder na Fase 0 (RN15-02: "dashboard genérico"). Retornar array vazio ou notificações mock. Link para área de notificações.

**DTO de resposta:**
```java
public record DashboardClientResponse(
    int activeUnits,           // Card 1
    String activeUnitsLink,
    int productCount,          // Card 2
    String productsLink,
    String planName,           // Card 3
    String planStatus,
    String planLink,
    List<Notification> notifications, // Card 4
    String notificationsLink
) {}
```

**Arquivos a criar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../service/DashboardClientService.java` | 🆕 | Agregação de dados para dashboard do cliente |
| `src/main/java/.../dto/response/DashboardClientResponse.java` | 🆕 | Record com 4 cards |
| `src/main/java/.../dto/response/NotificationResponse.java` | 🆕 | Record: title, message, link, severity |

**Dependências:** TenantContext (tenant_id do JWT) · BusinessUnitRepository · SubscriptionRepository

**Riscos:**
- **Baixo:** Queries simples de contagem. Performance adequada com índices existentes.
- **Mitigação:** Adicionar `@Cacheable` com TTL curto (1min) se houver preocupação de performance.

**Skills aplicáveis:** `311-frameworks-spring-jdbc`, `302-frameworks-spring-boot-rest`

---

### T-064 — `DashboardClientController`: endpoints do dashboard

- **Critério DONE:** Dados do cliente autenticado. Filtrado por tenant_id. `@RequiresPermission`.
- **Estimativa:** 1d (8h)
- **Papel:** Dev Backend
- **Prioridade:** Should
- **RACI:** Dev Backend (R), Dev Full-Stack (C), Tech Lead (A)

**Abordagem:**
Criar controller com 2 endpoints REST. Dashboard genérico na Fase 0 (RN15-01, RN15-02: adapta-se ao módulo, placeholder).

**Endpoints (SPECS.md §4.1):**

| Método | Path | Descrição |
|:---|:---|:---|
| `GET` | `/api/v1/dashboard/client/summary` | Cards do dashboard. Query param `?module_id` (RN15-01) |
| `GET` | `/api/v1/dashboard/client/notifications` | Notificações do cliente |

**Arquivos a criar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../controller/DashboardClientController.java` | 🆕 | 2 endpoints REST com `@RequiresPermission` |

**Dependências:** T-063 (DashboardClientService)

**Riscos:** Baixo — controller fino, lógica no Service.

**Skills aplicáveis:** `302-frameworks-spring-boot-rest`

---

### T-065 — Atualizar `JwtAuthenticationFilter`: claims `modules[]` + `business_unit_ids[]`

- **Critério DONE:** JWT retorna modules[]. App Switcher viável com 1 módulo.
- **Estimativa:** 1.5d (12h)
- **Papel:** Especialista IAM
- **Prioridade:** Must — CAMINHO CRÍTICO #3
- **RACI:** Especialista IAM (R), Dev Backend (C), Tech Lead (A)

**Abordagem:**
Atualizar `JwtAuthenticationFilter` (e/ou `FbsoJwtAuthenticationConverter`) para extrair as claims `modules[]` e `business_unit_ids[]` do JWT e populá-las no `TenantContext`. As claims já são mapeadas no Keycloak (T-057) como user attributes. O filter/converter precisa:

1. **Extrair `modules[]`** do JWT — `List<String> modules = jwt.getClaimAsStringList("modules")`
2. **Extrair `business_unit_ids[]`** do JWT — `List<String> buIds = jwt.getClaimAsStringList("business_unit_ids")`
3. **Popular `TenantContext`** — `TenantContext.set(tenantId, userId, roles, businessUnitIds, modules)`

**Placeholder "FBSO Platform" (RN16-01, RN16-02):** Se o claim `modules[]` estiver vazio ou ausente, popular com `List.of("FBSO Platform")` como módulo placeholder.

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../security/FbsoJwtAuthenticationConverter.java` | 🔄 | Extrair `modules[]` e `business_unit_ids[]` do JWT |
| `src/main/java/.../security/JwtAuthenticationFilter.java` | 🔄 | Popular TenantContext com os novos claims |
| `src/main/java/.../security/TenantContext.java` | 🔄 | Verificar se getters `getBusinessUnitIds()` e `getModules()` já existem |

**Dependências:** T-057 (Keycloak realm com claims mapeadas) · T-136 (JWT claims ✅) · T-148 (Converter ✅)

**Riscos:**
- **Baixo:** Claims já são extraídos pelo `FbsoJwtAuthenticationConverter`. Apenas adicionar 2 campos.
- **Mitigação:** Testar com JWT real do Keycloak (`docker compose up`) para validar claims.

**Skills aplicáveis:** `304-frameworks-spring-boot-security`

---

### T-066 — `GET /api/v1/auth/me`: dados do usuário logado

- **Critério DONE:** GET /auth/me funcional. Stateless (dados do token). 401 se sem token.
- **Estimativa:** 1d (8h)
- **Papel:** Dev Backend
- **Prioridade:** Must — CAMINHO CRÍTICO #4
- **RACI:** Dev Backend (R), Especialista IAM (C), Tech Lead (A)

**Abordagem:**
Criar endpoint que retorna os dados do usuário autenticado extraídos do JWT (stateless — sem consulta ao banco). O endpoint NÃO exige `@RequiresPermission` (autenticação já validada no filter). Implementar no `AuthController` ou em controller separado.

**Response (SPECS.md §4.1):**
```json
{
  "id": "uuid",
  "name": "João Silva",
  "email": "joao@exemplo.com",
  "role": "ADMIN_TENANT",
  "business_unit_ids": ["uuid1", "uuid2"],
  "modules": ["FBSO Platform"],
  "tenant_id": "uuid",
  "status_onboarding": "STEP1_DONE"
}
```

**Fonte dos dados:**
| Campo | Origem |
|:---|:---|
| `id` | JWT claim `sub` (Keycloak user ID) |
| `name` | JWT claim `name` ou `preferred_username` |
| `email` | JWT claim `email` |
| `role` | JWT claim `roles[0]` (papel principal) |
| `business_unit_ids` | JWT claim `business_unit_ids[]` (T-065) |
| `modules` | JWT claim `modules[]` (T-065) |
| `tenant_id` | JWT claim `tenant_id` |
| `status_onboarding` | Consulta ao banco: `SELECT onboarding_step FROM tenant WHERE id = ?` |

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../controller/AuthController.java` | 🔄 | Adicionar `GET /auth/me` (ou criar `ConfigController` separado) |

**Dependências:** T-065 (claims `modules[]` e `business_unit_ids[]` no TenantContext) · T-058 (AuthController existente)

**Riscos:**
- **Baixo:** Leitura de claims do JWT é operação trivial. Única consulta ao banco é `onboarding_step`.
- **Mitigação:** Cachear `onboarding_step` no próprio TenantContext (já disponível) para evitar query extra.

**Skills aplicáveis:** `302-frameworks-spring-boot-rest`, `304-frameworks-spring-boot-security`

---

### T-067 — Testes unitários M5

- **Critério DONE:** ≥ 80% cobertura. Rate limit testado (5 tentativas, bloqueio, reset 15min).
- **Estimativa:** 1.5d (12h)
- **Papel:** QA Engineer
- **Prioridade:** Must
- **RACI:** QA Engineer (R), Dev Backend (C), Tech Lead (A)

**Abordagem:**
Implementar testes unitários com JUnit 5 + Mockito para os novos services e controllers. Foco em:

1. **OnboardingServiceTest** — 6 edge cases (EC-1 a EC-6 do DT-124), transições de estado, rollback
2. **DashboardClientServiceTest** — queries mockadas, verificar agregação correta
3. **AuthServiceTest** — fluxo forgot-password, reset-password, delegação Keycloak
4. **RateLimitFilterTest** — (já parcialmente coberto na Frente 1) — completar com cenários de login
5. **AuthControllerTest** — MockMvc, 401 sem token, 200 com token válido

**Arquivos a criar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/test/.../unit/service/OnboardingServiceTest.java` | 🆕 | 6 edge cases + transições |
| `src/test/.../unit/service/DashboardClientServiceTest.java` | 🆕 | Queries mockadas |
| `src/test/.../unit/service/AuthServiceTest.java` | 🆕 | Fluxos auth |
| `src/test/.../unit/controller/AuthControllerTest.java` | 🆕 | MockMvc testes |
| `src/test/.../unit/controller/OnboardingControllerTest.java` | 🆕 | MockMvc testes |

**Dependências:** T-058..T-066 concluídas (código a testar)

**Skills aplicáveis:** `131-java-testing-unit-testing`, `321-frameworks-spring-boot-testing-unit-tests`

---

### T-068 — Testes de integração M5

- **Critério DONE:** Onboarding completo testado (PENDING→ACTIVE). Rate limiting com PostgreSQL real. Expiração link reset.
- **Estimativa:** 2d (16h)
- **Papel:** QA Engineer
- **Prioridade:** Must
- **RACI:** QA Engineer (R), Dev Backend (C), Tech Lead (A)

**Abordagem:**
Implementar testes de integração com Testcontainers (PostgreSQL 17 real). Cobrir fluxos completos:

1. **Fluxo onboarding completo:** Criar tenant → step-1 → step-2 (BU Matriz) → step-3 → complete → verificar Tenant.status = ACTIVE
2. **Rate limiting:** 5 tentativas de login com senha errada → 6ª bloqueada → aguardar 15min → desbloqueada
3. **Expiração link reset:** Simular token expirado → 401
4. **Isolamento cross-tenant:** Tenant A não vê onboarding do Tenant B
5. **Transições inválidas:** Tentar step-2 sem step-1 → 422

**Arquivos a criar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/test/.../integration/OnboardingIntegrationTest.java` | 🆕 | Fluxo completo com Testcontainers |
| `src/test/.../integration/AuthIntegrationTest.java` | 🆕 | Rate limit + reset password com PostgreSQL real |
| `src/test/.../integration/DashboardClientIntegrationTest.java` | 🆕 | Dashboard com dados reais |

**Dependências:** T-067 (testes unitários concluídos)

**Skills aplicáveis:** `132-java-testing-integration-testing`, `322-frameworks-spring-boot-testing-integration-tests`

---

## 5. Plano por Task — Frontend (Frente 3b)

> ⚠️ **Projeto separado:** O frontend está em `/frontend/javascript/react/web_apps/web_app-fbso-platform-portal/`. Stack: Next.js App Router + TypeScript + Tailwind CSS + next-auth (ADR-03). O desenvolvimento frontend ocorre em paralelo com o backend, usando MSW mock handlers para simular as APIs antes de estarem prontas.

---

### T-157 — Bootstrap do projeto frontend

- **Critério DONE:** `npm run dev` funcional. Rotas App Router corretas. MSW handlers ativos. ESLint/Prettier zero warnings.
- **Estimativa:** 2d (16h)
- **Papel:** Dev Frontend
- **Prioridade:** Must
- **Projeto:** `web_app-fbso-platform-portal` (separado do backend Java)

**Abordagem:**
Inicializar projeto Next.js com TypeScript e Tailwind CSS. Criar estrutura de diretórios com App Router (route groups para auth, onboarding, portal). Configurar MSW (Mock Service Worker) para interceptar chamadas à API durante desenvolvimento.

**Estrutura de diretórios:**
```
web_app-fbso-platform-portal/
├── app/
│   ├── (auth)/login/page.tsx
│   ├── (auth)/forgot-password/page.tsx
│   ├── (auth)/reset-password/page.tsx
│   ├── (onboarding)/steps/page.tsx
│   ├── (portal)/dashboard/page.tsx
│   └── layout.tsx
├── components/
│   ├── layout/ (Sidebar, Header, AppSwitcher)
│   ├── dashboard/ (MetricCard, NotificationList)
│   └── common/ (Skeleton, Toast, StatusBadge)
├── lib/
│   ├── auth.ts (next-auth config)
│   ├── api-client.ts (fetch wrapper com JWT)
│   └── permissions.ts (RBAC helpers)
├── mocks/
│   ├── handlers.ts (MSW handlers)
│   └── server.ts (MSW server setup)
└── tailwind.config.ts
```

**Dependências:** Nenhuma (projeto standalone)

**Skills aplicáveis:** `frontend-design`, `nextjs-best-practices`

---

### T-158 — Auth UI: Login + Reset Password

- **Critério DONE:** Login redirect Keycloak funcional. Form reset com validação. Sessão 60min.
- **Estimativa:** 2d (16h)
- **Papel:** Dev Frontend
- **Prioridade:** Must

**Abordagem:**
Implementar páginas de autenticação usando next-auth com Keycloak provider. Middleware protege rotas que exigem autenticação. Estados: loading skeleton, erro RFC 7807, sucesso.

**Páginas:**
1. `/login` — Botão "Entrar com FBSO" → redirect Keycloak OIDC
2. `/forgot-password` — Form: email → POST /auth/forgot-password. Toast: "Link enviado"
3. `/reset-password` — Form: nova senha + confirmação → POST /auth/reset-password. Validação RN13-01

**Dependências:** T-157 (bootstrap) · MSW mock para /auth/* endpoints

---

### T-159 — Onboarding wizard UI

- **Critério DONE:** 4 steps navegáveis. Barra progresso. Validação CNPJ. Mobile-first responsivo.
- **Estimativa:** 2.5d (20h)
- **Papel:** Dev Frontend
- **Prioridade:** Must

**Abordagem:**
Implementar wizard de 4 passos com barra de progresso (25%→50%→75%→100%). Cada passo valida antes de avançar. Estado salvo no backend permite retomar (GET /onboarding/status).

**Dependências:** T-157 (bootstrap) · T-158 (auth — proteção de rota) · MSW mock para /onboarding/* endpoints

---

### T-160 — Dashboard cliente: MetricCards

- **Critério DONE:** 4 cards responsivos. Skeleton loading. Empty state. Cards clicáveis.
- **Estimativa:** 1.5d (12h)
- **Papel:** Dev Frontend
- **Prioridade:** Should

**Abordagem:**
Grid responsivo com 4 MetricCards consumindo `GET /dashboard/client/summary`. Loading skeleton durante fetch, empty state se sem dados. Navegação via next/link.

**Dependências:** T-157 (bootstrap) · T-158 (auth) · MSW mock para /dashboard/client/*

---

## 6. Ordem de Execução

### Backend (Frente 3a)

| # | Task | Papel | Depende de | Paralelizável com |
|:---:|:---|:---|:---|:---|
| 1 | **T-057** — Keycloak realm | Especialista IAM | — | — (CAMINHO CRÍTICO) |
| 2 | **T-058** — Auth endpoints | Dev Backend | T-057 | T-060, T-063 |
| 3 | **T-060** — OnboardingService | Dev Full-Stack | T-057 (realm existe) | T-058, T-063 |
| 4 | **T-062** — BU Matriz | Dev Backend + Dev BD | T-060 | T-058, T-063 |
| 5 | **T-059** — Rate limiting | Dev Backend | T-058 (login endpoint) | T-061, T-064 |
| 6 | **T-061** — OnboardingController | Dev Backend | T-060, T-062 | T-059, T-064 |
| 7 | **T-063** — DashboardClientService | Dev Full-Stack | — (queries independentes) | T-058, T-060 |
| 8 | **T-064** — DashboardClientController | Dev Backend | T-063 | T-061, T-065 |
| 9 | **T-065** — JWT claims | Especialista IAM | T-057, T-058 | T-066 |
| 10 | **T-066** — GET /auth/me | Dev Backend | T-065 | — |
| 11 | **T-067** — Testes unitários | QA Engineer | T-058..T-066 | T-068 (início) |
| 12 | **T-068** — Testes integração | QA Engineer | T-067 | — |

### Frontend (Frente 3b) — Paralelo ao Backend

| # | Task | Papel | Depende de | Paralelizável com |
|:---:|:---|:---|:---|:---|
| 13 | **T-157** — Bootstrap | Dev Frontend | — | Backend T-057 |
| 14 | **T-158** — Auth UI | Dev Frontend | T-157 | Backend T-058..T-061 |
| 15 | **T-159** — Onboarding UI | Dev Frontend | T-157, T-158 | Backend T-062..T-065 |
| 16 | **T-160** — Dashboard UI (Should) | Dev Frontend | T-157, T-158 | Backend T-066 |

### Diagrama de Gantt (≈8d wall-clock com time paralelo):

```
DIA 1-2          DIA 3-4          DIA 5-6          DIA 7-8
├────────────────┼────────────────┼────────────────┼────────────────┤
│ T-057 (IAM)   │ T-058 (DevBack)│ T-065 (IAM)    │ T-066 (DevBack)│ ← CAMINHO CRÍTICO
│                │ T-060 (FullSt) │ T-062 (DevBack)│ T-061 (DevBack)│ ← ONBOARDING
│                │ T-063 (FullSt) │ T-064 (DevBack)│                │ ← DASHBOARD
│                │ T-059 (DevBack)│                │                │ ← RATE LIMIT
│                                                                │
│ T-157 (Front)  │ T-158 (Front)  │ T-159 (Front)  │ T-160 (Front) │ ← FRONTEND (paralelo)
│                                                                │
│                │                │                │ T-067 (QA)    │ ← TESTES
│                │                │                │ T-068 (QA)    │
└────────────────┴────────────────┴────────────────┴────────────────┘
```

---

## 7. Estratégia de Build e Verificação

### Comandos — Backend

| Propósito | Comando |
|:---|:---|
| Compilação | `mvn compile` |
| Testes unitários | `mvn test` |
| Testes específicos (onboarding) | `mvn test -Dtest="OnboardingServiceTest,OnboardingControllerTest"` |
| Testes específicos (auth) | `mvn test -Dtest="AuthServiceTest,AuthControllerTest"` |
| Testes integração | `mvn test -Dtest="*IntegrationTest"` |
| Keycloak realm validação | `docker compose up -d keycloak` → acessar localhost:8081/realms/fbso-platform |
| Verificação JWT claims | `curl localhost:8081/realms/fbso-platform/protocol/openid-connect/certs \| jq` |

### Comandos — Frontend

| Propósito | Comando |
|:---|:---|
| Dev server | `npm run dev` |
| Build | `npm run build` |
| Lint | `npm run lint` |
| Type check | `npx tsc --noEmit` |

### Checkpoints

| # | Momento | O que verificar |
|:---:|:---|:---|
| CP-1 | Após T-057 | Keycloak admin acessível. Realm `fbso-platform` existe. Claims mapeados via jwt.io |
| CP-2 | Após T-058 | `POST /auth/login` → redirect Keycloak. `POST /auth/forgot-password` → email no MailHog |
| CP-3 | Após T-059 | 5 tentativas falhas → 429. 6ª tentativa com tempo restante no body |
| CP-4 | Após T-060+T-061+T-062 | Onboarding completo: step-1→2→3→complete. Tenant.status = ACTIVE. BU com is_matrix=true |
| CP-5 | Após T-063+T-064 | `GET /dashboard/client/summary` retorna 4 cards com dados do tenant |
| CP-6 | Após T-065+T-066 | `GET /auth/me` retorna modules[], business_unit_ids[]. JWT contém claims |
| CP-7 | Após T-067+T-068 | `mvn test` completo. ≥ 80% cobertura. Onboarding integration test passa |
| CP-8 | Após T-157 | `npm run dev` funcional. MSW ativo. Rotas App Router navegáveis |
| CP-9 | Após T-158 | Login redirect Keycloak funcional no browser. Forms validados |
| CP-10 | Após T-159+T-160 | Onboarding wizard 4 steps. Dashboard 4 cards responsivos |

---

## 8. Observações

1. **IAM Specialist é o GARGALO:** T-057 (2d) + T-065 (1.5d) = 3.5d no caminho crítico. O SPRINT-CARD.md já alerta: "IAM Specialist 4h/dia pode ser insuficiente — T-057 + T-065 = 3.5d (28h = 7 dias wall-clock)". Recomenda-se iniciar T-057 IMEDIATAMENTE.

2. **Frontend é projeto separado:** T-157..T-160 estão em `/frontend/javascript/react/web_apps/web_app-fbso-platform-portal/` — Next.js + TypeScript. O desenvolvimento pode ocorrer em paralelo total com o backend usando MSW mocks.

3. **MSW mocks são críticos:** Sem os mocks, o frontend fica bloqueado até o backend entregar os endpoints. Prioridade: criar handlers MSW para `/auth/*`, `/onboarding/*`, `/dashboard/client/*` e `/api/v1/auth/me` logo no T-157.

4. **T-059 (Rate limiting) é integração, não construção:** O `RateLimitFilter` já existe. A task é integrar ao endpoint de login — estimativa de 1.5d é conservadora.

5. **Teste pré-existente quebrado:** `SubscriptionServiceTest.shouldCreateWithLockedPrice` — erro de TenantContext. Corrigir no `@BeforeEach` durante T-067.

6. **Onboarding Fase 0 — simplificado:** Step-3 (configurações fiscais) é placeholder. Sem expiração de onboarding. Sem envio de email durante onboarding. Estas capacidades são evolutivas (Sprint 6+).

7. **Placeholder "FBSO Platform" (RN16-01, RN16-02):** App Switcher exibe 1 módulo placeholder. Reais módulos virão com catálogo de produtos/serviços na Sprint 6.

---

🤖 *Documento gerado em 2026-07-23 como parte da Fase 1 do PROMPT-EXECUTE-SPRINT-TASKS.md. Stack detectada: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Keycloak 26 + Flyway 12.11.0 + Caffeine 3.2.4 (backend) + Next.js App Router + TypeScript + Tailwind CSS (frontend). Skills utilizados no planejamento: 302-frameworks-spring-boot-rest, 304-frameworks-spring-boot-security, 311-frameworks-spring-jdbc, 121-java-object-oriented-design, 030-architecture-adr-general, 131-java-testing-unit-testing, 132-java-testing-integration-testing.*
