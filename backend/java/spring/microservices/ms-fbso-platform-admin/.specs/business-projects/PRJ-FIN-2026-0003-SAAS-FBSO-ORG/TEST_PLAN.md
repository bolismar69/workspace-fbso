# TEST_PLAN.md — Plano de Testes: ms-fbso-platform-admin

- **Solucao:** `ms-fbso-platform-admin`
- **Stack:** Java 25 + Spring Boot + PostgreSQL
- **Projeto de Negocio:** [PRJ-FIN-2026-0003-SAAS-FBSO-ORG](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/)
- **Versao:** 2.7
- **Data:** 17 de Julho de 2026
- **Status:** Em Execucao — M2 (EP-01) 100% testado. 105 testes totais: 77 unitários (Surefire) + 28 integração PostgreSQL real (Failsafe). 0 falhas, 6 skipped. JaCoCo: Instructions 87.1%, Lines 85.8%, Branches 64.6%. T-023 DashboardRepositoryIT concluído (23 cenários). maven-failsafe-plugin configurado. F01-01 a F01-03 cobertas.
- **Origem:** [SPECS.md](./SPECS.md) v1.4 + [ARCHITECTURE.md](./ARCHITECTURE.md) + [PRD.md](./PRD.md)

---

## 1. Estrategia de Testes

### 1.1 Piramide de Testes

```
                    /\
                   /  \
                  / E2E\   ~5%   — Playwright (frontend + backend integrados)
                 /      \         Fluxos completos: login, onboarding, CRUD multi-papel
                /--------\
               / Seguranca\ ~5%  — OWASP ZAP + JWT manipulation + RBAC adversarial
              /            \     Testes de penetracao automatizados
             /--------------\
            /                \
           /   Integracao     \ ~20% — Testcontainers (PostgreSQL 17 real)
          /                    \      Repositorios + aspectos AOP (TenantIsolation, RBAC, Auditoria)
         /----------------------\
        /                        \
       /       Unitarios          \ ~70% — JUnit 5 + Mockito
      /                            \      Services, validators, exceptions, utils
     /------------------------------\
```

**Distribuicao alvo:** ~70% Unitarios, ~20% Integracao, ~5% E2E, ~5% Seguranca

### 1.2 Ferramentas por Nivel

| Nivel | Ferramenta | Framework | Responsabilidade |
|:---|:---|:---|:---|
| **Unitario** | JUnit 5 + Mockito | Spring Boot Test | Services com repositorios mockados, validadores, utilitarios, excecoes |
| **Integracao** | Testcontainers | spring-boot-starter-test | Repositorios (PostgreSQL real), Aspectos AOP, Pipeline de seguranca, APIs |
| **E2E** | Playwright + REST Assured | Testcontainers + Keycloak containers | Fluxos completos: frontend consumindo API, multi-papel, onboarding |
| **Seguranca** | OWASP ZAP + JWT manipulation | Testcontainers | Injection (SQL, XSS), Broken Access Control, CSRF, data leakage |
| **Performance** | JMeter + Micrometer | spring-boot-starter-actuator | Carga no dashboard, listas paginadas, endpoints de tenants |
| **Cobertura** | JaCoCo | maven-surefire-plugin | Meta minima: 80% (linhas), 70% (branchs) |
| **Qualidade** | Checkstyle + PMD | maven-checkstyle-plugin | Zero warnings no codigo de producao |
| **Documentacao** | OpenAPI diff | swagger-request-validator | Contrato OpenAPI nunca quebrado |

### 1.3 Metas de Cobertura

| Nivel | Aspecto | Meta | Verificacao |
|:---|:---|:---:|:---|
| Unitario | Cobertura de linhas (Services) | >= 85% | JaCoCo |
| Unitario | Cobertura de branchs (Services) | >= 75% | JaCoCo |
| Integracao | Cobertura de endpoints REST | 100% dos 37 endpoints | Teste automatizado por controller |
| Integracao | Cobertura de RNs | 100% das 18 familias de RNs (PRD §6.6) | Checklist de aceitacao |
| Seguranca | RBAC cada papel x endpoint proibido | 100% da matriz RN10-01 | Teste parametrizado |
| Seguranca | Isolamento Multi-Tenant | 100% dos repositorios testados | Teste integracao |
| E2E | Fluxos criticos | 5 fluxos completos | Playwright |
| Performance | p95 <= 3s (dashboard) | 3 endpoints criticos | JMeter |

### 1.4 Organizacao dos Testes no Projeto

```
src/test/java/com/fbso/platform/admin/
│
├── unit/
│   ├── service/           ← Testes unitarios de services (mocks)
│   ├── validator/         ← Testes de validacao (CNPJ, status transition, etc.)
│   ├── exception/         ← Testes de excecoes e RFC 7807
│   └── utils/             ← Testes de utilitarios (CnpjValidator, JwtUtils, DateUtils)
│
├── integration/
│   ├── repository/        ← Testes com Testcontainers (queries reais)
│   ├── security/          ← Testes de JWT Filter, TenantIsolation, RBAC, Auditoria
│   ├── controller/        ← Testes de API (WebMvcTest + Testcontainers)
│   └── aspect/            ← Testes dos aspectos AOP
│
├── e2e/
│   └── flow/              ← Testes E2E com Playwright + Keycloak container
│
└── security/
    ├── owasp/             ← Testes OWASP (SQL Injection, XSS, CSRF)
    ├── rbac/              ← Testes de matriz de permissoes (parametrizado)
    └── multitenant/       ← Testes de isolamento entre tenants
```

---

## 2. Mapa de Cobertura: Feature x Cenarios

| Feature | Descricao | Unit | Int | E2E | Seg | Total |
|:---|:---|:---:|:---:|:---:|:---:|:---:|
| **F01-01** | Dashboard Admin — metricas | 3 | 2 | 1 | 1 | 7 |
| **F01-02** | Lista de Contas | 2 | 2 | — | — | 4 |
| **F01-03** | Alertas do Dashboard | 2 | 2 | 1 | — | 5 |
| **F02-01** | Criar Tenant | 3 | 2 | 1 | 1 | 7 |
| **F02-02** | Transicoes de Status Tenant | 4 | 3 | 1 | 1 | 9 |
| **F02-03** | Configuracao de Planos | 3 | 2 | 1 | 1 | 7 |
| **F02-04** | Vinculacao de Assinaturas | 5 | 3 | 1 | 1 | 10 |
| **F02-05** | Auditoria | 3 | 4 | — | 2 | 9 |
| **F03-01** | Gestao de Usuarios | 3 | 2 | 1 | 1 | 7 |
| **F03-02** | Matriz de Permissoes RBAC | 2 | 2 | 1 | 4 | 9 |
| **F03-03** | Vinculacao Usuario x Unidade x Modulo | 2 | 2 | — | 2 | 6 |
| **F03-04** | Acesso Condicional (403) | 1 | 1 | 1 | 2 | 5 |
| **F04-01** | Login e Autenticacao | 2 | 2 | 2 | 3 | 9 |
| **F04-02** | Onboarding Guiado | 4 | 3 | 1 | 1 | 9 |
| **F04-03** | Dashboard do Cliente | 2 | 2 | 1 | — | 5 |
| **F04-04** | App Switcher | 1 | 1 | 1 | 1 | 4 |
| **F04-05** | Unidades de Negocio | 4 | 3 | 1 | 1 | 9 |
| **F04-06** | Catalogo de Produtos/Servicos | 3 | 2 | 1 | 1 | 7 |
| **Total** | **18 features** | **49** | **38** | **16** | **23** | **126** |

### Cobertura Adicional (Transversal)

| Area | Unit | Int | Seg | Carga | Total |
|:---|:---:|:---:|:---:|:---:|:---:|
| Pipeline de seguranca (JWT -> TenantContext -> RBAC -> Auditoria) | — | 1 | 3 | — | 4 |
| Conformidade RFC 7807 (erros) | 2 | 1 | — | — | 3 |
| Isolamento Multi-Tenant | — | 2 | 5 | — | 7 |
| Soft Delete | 2 | 2 | — | — | 4 |
| Concorrencia / Race Conditions | — | 3 | — | — | 3 |
| Carga / Performance | — | — | — | 3 | 3 |
| Regressao | — | — | — | — | 1 |

**Total Geral de Cenarios: 151 (features + cross-cutting) + 25 (infraestrutura §9) = 176**

---

## 3. Cenarios de Teste por Feature

### 3.1 F01-01: Dashboard Administrativo — Metricas

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F01-01-001 | Dashboard summary carrega indicadores corretos | Unit | Repositorio mockado com 5 tenants (3 ACTIVE, 1 PENDING, 1 SUSPENDED) | 1. Chamar DashboardService.getSummary() | Retorna total=5, ativos=3, pendentes=1, suspensos=1. Metricas consideram apenas tenants com deleted_dt IS NULL | Planejado |
| TC-F01-01-002 | Dashboard evolution com periodo "30d" retorna dados filtrados | Unit | Repositorio mockado com dados de 60 dias | 1. Chamar DashboardService.getEvolution("30d") | Retorna apenas registros dos ultimos 30 dias | Planejado |
| TC-F01-01-003 | Dashboard evolution com periodo invalido assume mes atual | Unit | Repositorio mockado | 1. Chamar DashboardService.getEvolution("invalido") | Retorna dados do mes corrente (fallback seguro) | Planejado |
| TC-F01-01-004 | GET /dashboard/admin/summary retorna 200 com JSON valido | Integracao | Testcontainers com dados de seed, JWT valido | 1. Autenticar como Admin FBSO<br>2. GET /api/v1/dashboard/admin/summary | Status 200. Response contem totalTenants, byStatus, byPlan. Resposta em <=3s | Planejado |
| TC-F01-01-005 | GET /dashboard/admin/summary exclui tenants soft-deletados | Integracao | Testcontainers com 1 tenant deletado (deleted_dt preenchido) | 1. Autenticar como Admin FBSO<br>2. GET /api/v1/dashboard/admin/summary | Tenant soft-deletado NAO aparece nas metricas | Planejado |
| TC-F01-01-006 | Fluxo E2E: login Admin FBSO visualiza dashboard com dados reais | E2E | Keycloak + PostgreSQL com dados seed de 20 tenants | 1. Login como admin_fbso@fbso.org<br>2. Navegar para /admin/dashboard<br>3. Verificar metricas | Dashboard carrega em <=3s. Cards exibem contagem correta. Periodo padrao: mes atual | Planejado |
| TC-F01-01-007 | Dashboard sem autenticacao retorna 401 | Seguranca | Nenhum token JWT | 1. GET /api/v1/dashboard/admin/summary sem header Authorization | Status 401. Mensagem: "Token de acesso nao informado" | Planejado |

### 3.2 F01-02: Lista de Contas

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F01-02-001 | Lista paginada de tenants retorna page correta | Unit | Repositorio mockado com 30 tenants | 1. Chamar TenantService.findAll(page=0, size=25) | Retorna 25 itens na pagina 0, totalPages=2, totalElements=30 | Planejado |
| TC-F01-02-002 | Busca textual por "Mercado" filtra resultados | Unit | Repositorio mockado com 5 tenants (3 contem "Mercado") | 1. Chamar TenantService.findAll(search="Mercado") | Retorna apenas 3 tenants com "Mercado" no nome | Planejado |
| TC-F01-02-003 | GET /tenants com paginacao e filtros retorna dados corretos | Integracao | Testcontainers com 10 tenants, Admin FBSO auth | 1. GET /api/v1/tenants?page=0&size=5&status=ACTIVE | Status 200. 5 itens. Campos: razao social, plano, status, data criacao | Planejado |
| TC-F01-02-004 | GET /tenants com busca de 2 caracteres retorna vazio (min 3) | Integracao | Testcontainers, busca textual | 1. GET /api/v1/tenants?search=ab | Status 200. Lista vazia. Busca so ativa com >=3 caracteres | Planejado |

### 3.3 F01-03: Alertas do Dashboard

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F01-03-001 | Alertas incluem tenants com onboarding incompleto >48h | Unit | Repositorio mockado com 2 tenants PENDING_ONBOARDING ha 72h | 1. Chamar DashboardService.getAlerts() | Retorna alerta de onboarding incompleto para cada tenant com >48h | Planejado |
| TC-F01-03-002 | Alertas incluem tenants com assinatura suspensa | Unit | Repositorio mockado com 1 tenant com subscription SUSPENDED | 1. Chamar DashboardService.getAlerts() | Retorna alerta de assinatura suspensa | Planejado |
| TC-F01-03-003 | GET /dashboard/admin/alerts retorna cards de alerta | Integracao | Testcontainers com dados que disparam alertas | 1. GET /api/v1/dashboard/admin/alerts | Status 200. Lista de alertas com tipo, mensagem, tenant_id, link | Planejado |
| TC-F01-03-004 | Sem alertas, endpoint retorna lista vazia | Integracao | Testcontainers sem condicoes de alerta | 1. GET /api/v1/dashboard/admin/alerts | Status 200. Lista vazia | Planejado |
| TC-F01-03-005 | E2E: Admin ve cards de alerta coloridos no dashboard | E2E | Dados seed com onboarding pendente >48h e assinatura suspensa | 1. Login como Admin FBSO<br>2. Visualizar dashboard<br>3. Verificar cards de alerta | Cards aparecem no topo com cores distintas (amarelo/vermelho). Clicaveis — levam a lista filtrada | Planejado |

### 3.4 F02-01: Criar Tenant

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F02-01-001 | Criar tenant define status PENDING_ONBOARDING | Unit | Repositorio mockado, name_corporate unico | 1. Chamar TenantService.create(req) com req valida | Status = PENDING_ONBOARDING. Auditoria registrada (RN04-01) | Planejado |
| TC-F02-01-002 | Criar tenant com razao social duplicada retorna erro | Unit | Repositorio mockado indica que nome ja existe | 1. Chamar TenantService.create(req) com nome duplicado | Lanca BusinessException. Mensagem: "Razao social ja cadastrada" | Planejado |
| TC-F02-01-003 | Criar tenant sem name_corporate retorna erro de validacao | Unit | Mock de validator | 1. Chamar validacao de TenantCreateRequest sem name_corporate | Erro de Bean Validation: "Razao social e obrigatoria" | Planejado |
| TC-F02-01-004 | POST /tenants cria registro e retorna 201 | Integracao | Testcontainers, Admin FBSO auth | 1. POST /api/v1/tenants com JSON valido | Status 201. Location header presente. TenantResponse com status PENDING_ONBOARDING | Planejado |
| TC-F02-01-005 | POST /tenants com name_corporate duplicado retorna 409 | Integracao | Testcontainers com tenant existente | 1. POST /api/v1/tenants com mesmo name_corporate | Status 409 Conflict. Detail: "Razao social ja cadastrada" | Planejado |
| TC-F02-01-006 | E2E: Admin cria tenant, ve email de convite, tenant aparece no dashboard | E2E | Keycloak + PostgreSQL + SMTP mock | 1. Login como Admin FBSO<br>2. Preencher formulario de criacao de tenant<br>3. Submeter<br>4. Verificar email recebido<br>5. Verificar tenant na lista | Tenant criado com status PENDING_ONBOARDING. Email enviado com link unico. Aparece no dashboard | Planejado |
| TC-F02-01-007 | Criar tenant sem autenticacao retorna 401 | Seguranca | Nenhum token | 1. POST /api/v1/tenants sem Authorization | Status 401 | Planejado |

### 3.5 F02-02: Transicoes de Status Tenant

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F02-02-001 | Transicao PENDING->ACTIVE permitida | Unit | Tenant status PENDING | 1. Chamar TenantService.activate(id) | Status alterado para ACTIVE. Auditoria registrada | Planejado |
| TC-F02-02-002 | Transicao ACTIVE->SUSPENDED permitida com motivo | Unit | Tenant status ACTIVE | 1. Chamar TenantService.suspend(id, "Inadimplencia") | Status = SUSPENDED. Motivo registrado. Auditoria: action=SUSPENDED | Planejado |
| TC-F02-02-003 | Tentativa ACTIVE->PENDING retorna 422 | Unit | Tenant status ACTIVE | 1. Chamar TenantService.suspend(id, null) | Lanca InvalidStatusTransitionException | Planejado |
| TC-F02-02-004 | Transicao ACTIVE->INACTIVE permitida | Unit | Tenant status ACTIVE | 1. Chamar TenantService.deactivate(id) | Status = INACTIVE | Planejado |
| TC-F02-02-005 | POST /tenants/{id}/suspend sem motivo retorna 400 | Integracao | Testcontainers, Admin FBSO auth | 1. POST /api/v1/tenants/{id}/suspend com JSON sem campo reason | Status 400. Detail: "Motivo da suspensao e obrigatorio" | Planejado |
| TC-F02-02-006 | POST /tenants/{id}/suspend com motivo valido retorna 200 | Integracao | Testcontainers, Admin FBSO auth, tenant ACTIVE | 1. POST /api/v1/tenants/{id}/suspend com reason="Inadimplencia" | Status 200. TenantResponse.status = SUSPENDED. Auditoria com motivo registrado | Planejado |
| TC-F02-02-007 | POST /tenants/{id}/suspend bloqueia acesso em <=5min | Integracao | Testcontainers, tenant ACTIVE | 1. POST /api/v1/tenants/{id}/suspend<br>2. Tentar GET como usuario do tenant | Acesso bloqueado (403). TenantIsolation impede operacoes | Planejado |
| TC-F02-02-008 | E2E: Transicao completa ACTIVE->SUSPENDED->ACTIVE | E2E | Admin FBSO logado, tenant ACTIVE | 1. Suspender tenant com motivo<br>2. Verificar tenant aparece como SUSPENDED<br>3. Reativar tenant<br>4. Verificar tenant volta a ACTIVE | Timeline de status exibe historico completo. Transicoes respeitam RN05-01 | Planejado |
| TC-F02-02-009 | Operador tenta suspender tenant — 403 | Seguranca | JWT de OPERATOR | 1. POST /api/v1/tenants/{id}/suspend como OPERATOR | Status 403. Apenas Admin FBSO pode suspender | Planejado |

### 3.6 F02-03: Configuracao de Planos

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F02-03-001 | Criar plano com nome, price>0, recorrencia valida | Unit | Repositorio mockado | 1. Chamar PlanService.create(req) com req valida | Plano criado com status ACTIVE, version=1 | Planejado |
| TC-F02-03-002 | Criar plano com price=0 retorna erro | Unit | Mock de validator | 1. Chamar validacao com price=0 | Erro: "Preco deve ser maior que zero" | Planejado |
| TC-F02-03-003 | Editar plano gera nova versao | Unit | Plano existente version=1 | 1. Chamar PlanService.update(id, req) | Plano atualizado, version=2. Preco antigo preservado em assinaturas ativas (RN06-02) | Planejado |
| TC-F02-03-004 | POST /plans retorna 201 | Integracao | Testcontainers, Admin FBSO auth | 1. POST /api/v1/plans com JSON valido | Status 201. PlanResponse com id, name, price, status | Planejado |
| TC-F02-03-005 | POST /plans/{id}/deactivate com assinantes ativos retorna 422 | Integracao | Testcontainers, plano com 2 assinaturas ACTIVE | 1. POST /api/v1/plans/{id}/deactivate | Status 422. Detail: "Plano possui assinantes ativos e nao pode ser desativado" (RN06-01) | Planejado |
| TC-F02-03-006 | E2E: Admin cria plano, assina tenant, altera preco — assinatura mantem preco original | E2E | Admin FBSO logado, tenant ativo | 1. Criar plano (price=100)<br>2. Assinar tenant<br>3. Alterar preco do plano para 200<br>4. Verificar assinatura do tenant | Assinatura mantem preco original (100). Plano exibe nova versao com preco 200 | Planejado |
| TC-F02-03-007 | Operador tenta criar plano — 403 | Seguranca | JWT de OPERATOR | 1. POST /api/v1/plans como OPERATOR | Status 403 | Planejado |

### 3.7 F02-04: Vinculacao de Assinaturas

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F02-04-001 | Criar assinatura para tenant sem assinatura ativa permite criacao | Unit | Tenant sem subscription ACTIVE | 1. Chamar SubscriptionService.create(tid, req) | Assinatura criada com status ACTIVE, start_date = now() | Planejado |
| TC-F02-04-002 | Criar segunda assinatura ativa para mesmo tenant retorna erro | Unit | Tenant com 1 subscription ACTIVE | 1. Chamar SubscriptionService.create(tid, req) | Lanca BusinessException: "Tenant ja possui uma assinatura ativa" (RN07-01) | Planejado |
| TC-F02-04-003 | Upgrade/downgrade: assinatura anterior finalizada, nova criada | Unit | Tenant com 1 subscription ACTIVE | 1. Chamar SubscriptionService.changePlan(subId, newPlanId) | Assinatura anterior end_date = now(). Nova assinatura ACTIVE com novo plan_id | Planejado |
| TC-F02-04-004 | Suspender assinatura bloqueia modulos | Unit | Subscription ACTIVE | 1. Chamar SubscriptionService.suspend(subId, motivo) | Status = SUSPENDED. Registro em audit_log | Planejado |
| TC-F02-04-005 | POST /tenants/{tid}/subscriptions retorna 201 | Integracao | Testcontainers, Admin FBSO auth, tenant sem assinatura | 1. POST /api/v1/tenants/{tid}/subscriptions com JSON valido | Status 201. SubscriptionResponse com status ACTIVE | Planejado |
| TC-F02-04-006 | POST /tenants/{tid}/subscriptions com tenant ja ativo retorna 409 | Integracao | Testcontainers, tenant ja com 1 subscription ACTIVE | 1. POST /api/v1/tenants/{tid}/subscriptions | Status 409 Conflict. Detail: "Tenant ja possui uma assinatura ativa" | Planejado |
| TC-F02-04-007 | POST /subscriptions/{id}/change-plan valido retorna 200 | Integracao | Testcontainers, subscription ACTIVE, plano ativo | 1. POST /api/v1/subscriptions/{id}/change-plan com plan_id valido | Status 200. Assinatura anterior encerrada. Nova subscription ACTIVE | Planejado |
| TC-F02-04-008 | E2E: Tenant assina plano, faz upgrade, historico mostra timeline | E2E | Admin FBSO logado, tenant ativo, multiplos planos | 1. Assinar tenant no plano Basico<br>2. Fazer upgrade para plano Avancado<br>3. Verificar historico de assinaturas | Historico mostra ambas assinaturas com timeline. Apenas 1 ativa por vez | Planejado |
| TC-F02-04-009 | Concorrencia: duas requisicoes simultaneas de assinatura para mesmo tenant | Integracao | Testcontainers, race condition | 1. Disparar 2 POST /tenants/{tid}/subscriptions simultaneos | Uma retorna 201, a outra retorna 409. Nenhuma condicao de corrida permite 2 assinaturas ativas | Planejado |
| TC-F02-04-010 | Change-plan: locked_price preserva preço da assinatura original | Unit | Subscription ACTIVE com locked_price = 100, plano alterado para price = 200 | 1. Chamar SubscriptionService.changePlan(subId, newPlanId) onde locked_price=100 e novo plano custa 200 | Nova subscription criada com locked_price = 100 (preço original). Preço do novo plano (200) não altera locked_price (DT-009) | Planejado |

### 3.8 F02-05: Auditoria

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F02-05-001 | Auditoria registra criacao de tenant | Unit | AuditAspect mockado | 1. Chamar TenantService.create(req) | AuditEntry gerado com action=CREATED, entity_type=TENANT, previous_value=null, new_value preenchido | Planejado |
| TC-F02-05-002 | Auditoria registra alteracao de status (SUSPENDED) | Unit | AuditAspect mockado | 1. Chamar TenantService.suspend(id, motivo) | AuditEntry com action=SUSPENDED, previous_value="ACTIVE", new_value="SUSPENDED" | Planejado |
| TC-F02-05-003 | Consulta de auditoria com filtro por periodo retorna registros corretos | Unit | AuditRepository mockado com 10 registros | 1. Chamar AuditService.findAll(startDate, endDate) | Retorna apenas registros no periodo. Paginacao e sorting funcionam | Planejado |
| TC-F02-05-004 | GET /audit com filtros retorna dados paginados | Integracao | Testcontainers com 50 registros de auditoria | 1. GET /api/v1/audit?start_date=2026-08-01&end_date=2026-08-31&page=0&size=25 | Status 200. 25 itens. Formato AuditEntryResponse | Planejado |
| TC-F02-05-005 | Tentativa de UPDATE direto em audit_log retorna erro | Integracao | Testcontainers | 1. Executar UPDATE audit_log SET ... WHERE id = ? | Erro PostgreSQL (trigger ou constraint impede). Registros imutaveis (RN08-02) | Planejado |
| TC-F02-05-006 | GET /audit sem paginacao usa defaults (size=25, sort=timestamp DESC) | Integracao | Testcontainers com auditoria | 1. GET /api/v1/audit | Status 200. Page com size=25, ordenado por timestamp DESC | Planejado |
| TC-F02-05-007 | Auditoria captura acao de usuario sem permissao (403) | Seguranca | AuditAspect registra tentativas | 1. Tentar acesso proibido como OPERATOR<br>2. Consultar audit_log | Tentativa de acesso negado registrada em audit_log (RN08-01) | Planejado |
| TC-F02-05-008 | Tentativa de DELETE em audit_log via API retorna 403 | Seguranca | Admin FBSO auth | 1. Qualquer verbo de escrita em /audit (POST, PUT, PATCH, DELETE) | Status 403. "Registros de auditoria sao imutaveis" | Planejado |
| TC-F02-05-009 | Auditoria @Async: tenant_id e user_id corretos no registro (não UUID.randomUUID) | Integracao | AuditAspect com @Async configurado, TenantContext com tenant_id e user_id reais | 1. Executar ação auditável (ex: criar tenant)<br>2. Consultar audit_log gerado | AuditEntry.tenant_id = tenant_id real do TenantContext (não UUID.randomUUID). AuditEntry.actor_id = user_id real do JWT (DT-002) | Planejado |

### 3.9 F03-01: Gestao de Usuarios

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F03-01-001 | Convidar usuario com email unico no tenant cria registro | Unit | Repositorio mockado, email unico | 1. Chamar UserService.invite(req) com email valido | User criado, status=INVITE_PENDING. Auditoria registrada | Planejado |
| TC-F03-01-002 | Convidar usuario com email duplicado no tenant retorna erro | Unit | Repositorio mockado com email ja existente | 1. Chamar UserService.invite(req) com mesmo email | Lanca BusinessException: "Email ja cadastrado neste tenant" | Planejado |
| TC-F03-01-003 | Admin nao pode desativar a si mesmo | Unit | Contexto com admin_id = user_id | 1. Chamar UserService.deactivate(adminId) | Lanca BusinessException: "Um administrador nao pode desativar a si mesmo" (RN09-03) | Planejado |
| TC-F03-01-004 | POST /users retorna 201 | Integracao | Testcontainers, Admin Tenant auth | 1. POST /api/v1/users com JSON valido | Status 201. UserResponse com status INVITE_PENDING | Planejado |
| TC-F03-01-005 | POST /users/{id}/deactivate com auto-desativacao retorna 422 | Integracao | Testcontainers, Admin Tenant tentando desativar a si mesmo | 1. POST /api/v1/users/{adminId}/deactivate | Status 422. Detail: "Um administrador nao pode desativar a si mesmo" | Planejado |
| TC-F03-01-006 | E2E: Admin convida usuario, usuario recebe email, admin desativa usuario | E2E | Admin Tenant logado, SMTP mock | 1. Convidar usuario com email valido<br>2. Verificar email de convite<br>3. Desativar usuario<br>4. Verificar status | Usuario aparece como INVITE_PENDING, depois INACTIVE. Email de convite enviado | Planejado |
| TC-F03-01-007 | Convidar usuario sem permissao (OPERATOR) — 403 | Seguranca | JWT de OPERATOR | 1. POST /api/v1/users como OPERATOR | Status 403. Apenas Admin Tenant pode convidar | Planejado |

### 3.10 F03-02: Matriz de Permissoes RBAC

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F03-02-001 | Matriz RN10-01: role ADMIN_TENANT permite qualquer acao | Unit | RoleResource seed data | 1. Verificar ADMIN_TENANT contra todos resources/actions | ADMIN_TENANT tem permissao para todos os resources/actions | Planejado |
| TC-F03-02-002 | Matriz RN10-01: role AUDITOR permite apenas leitura | Unit | RoleResource seed data | 1. Verificar AUDITOR contra resources de escrita (create, edit, delete) | AUDITOR NAO tem permissao para nenhuma acao de escrita | Planejado |
| TC-F03-02-003 | Teste parametrizado: cada papel x cada endpoint | Seguranca | Testcontainers + JWT para cada role | 1. Para cada (papel, endpoint, metodo) na matriz: chamar endpoint | Papeis autorizados recebem 200/201. Papeis nao autorizados recebem 403. 100% da matriz coberta | Planejado |
| TC-F03-02-004 | OPERATOR tentando PATCH /products retorna 403 | Seguranca | JWT de OPERATOR | 1. PATCH /api/v1/products/{id} como OPERATOR | Status 403. Detail: "Voce nao tem permissao para executar esta operacao." | Planejado |
| TC-F03-02-005 | AUDITOR tentando POST /tenants retorna 403 | Seguranca | JWT de AUDITOR | 1. POST /api/v1/tenants como AUDITOR | Status 403 | Planejado |
| TC-F03-02-006 | MANAGER pode criar BusinessUnit (BU do seu tenant) | Integracao | Testcontainers, JWT de MANAGER | 1. POST /api/v1/business-units como MANAGER | Status 201 (MANAGER pode criar BU) | Planejado |
| TC-F03-02-007 | ADMIN_TENANT pode criar BusinessUnit | Integracao | Testcontainers, JWT de ADMIN_TENANT | 1. POST /api/v1/business-units como ADMIN_TENANT | Status 201 | Planejado |
| TC-F03-02-008 | E2E: Login com cada papel verifica menu e permissoes | E2E | Keycloak com 4 usuarios (Admin FBSO, Admin Tenant, Manager, Operator, Auditor) | 1. Login como Admin — ver menu completo<br>2. Login como Manager — ver BUs e produtos<br>3. Login como Operator — apenas leitura<br>4. Login como Auditor — apenas leitura | Menu condicional ao papel. Botoes de acao aparecem apenas para papeis autorizados | Planejado |
| TC-F03-02-009 | Admin FBSO pode ver todos os tenants (cross-tenant) | Seguranca | Testcontainers, Admin FBSO auth | 1. GET /api/v1/tenants | Admin FBSO ve TODOS os tenants (diferente de Admin Tenant que ve apenas o seu) | Planejado |

### 3.11 F03-03: Vinculacao Usuario x Unidade x Modulo

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F03-03-001 | Atribuir permissao: usuario vinculado a BU e modulo | Unit | Repositorio mockado | 1. Chamar PermissionService.assign(userId, buId, role) | UserPermission criado com (userId, buId, role). Auditoria registrada | Planejado |
| TC-F03-03-002 | Admin tem acesso implicito a todas as unidades | Unit | Role = ADMIN_TENANT | 1. Chamar PermissionService.hasAccess(adminId, qualquerBU) | Retorna true — admin ve todas as BUs | Planejado |
| TC-F03-03-003 | Usuario sem vinculacao nao acessa portal | Integracao | Testcontainers, JWT de usuario sem UserPermission | 1. GET /api/v1/products | Status 403. Usuario sem permissao de acesso | Planejado |
| TC-F03-03-004 | Usuario com vinculacao a BU especifica ve apenas produtos daquela BU | Integracao | Testcontainers, 2 BUs com produtos, usuario vinculado a BU-1 | 1. GET /api/v1/products | Retorna apenas produtos da BU-1. Produtos da BU-2 invisiveis | Planejado |
| TC-F03-03-005 | Protecao adversarial: usuario BU-1 tenta acessar produto BU-2 por ID direto | Seguranca | Testcontainers, JWT usuario BU-1 | 1. GET /api/v1/products/{productId-BU2} | Status 404 (nao 403 — tenant nao sabe que o registro existe) | Planejado |
| TC-F03-03-006 | Efeito imediato: alterar permissao reflete na proxima requisicao | Seguranca | Testcontainers | 1. Remover permissao de usuario<br>2. Imediatamente tentar GET /api/v1/products | Status 403. Sem cache de permissoes — efeito imediato | Planejado |

### 3.12 F03-04: Acesso Condicional (403 Amigavel)

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F03-04-001 | Acesso direto a URL proibida retorna 403 amigavel (sem detalhes tecnicos) | Unit | RbacAspect mock negado | 1. Chamar endpoint com @RequiresPermission sem permissao | GlobalExceptionHandler retorna 403 com titulo "Acesso negado". SEM stack trace, SEM detalhes internos (RN12-01) | Planejado |
| TC-F03-04-002 | Resposta 403 segue RFC 7807 | Integracao | Testcontainers, OPERATOR tenta POST /plans | 1. POST /api/v1/plans como OPERATOR | Status 403. Body: {type, title: "Acesso negado", status:403, detail}. Sem stack trace. Mensagem em PT-BR | Planejado |
| TC-F03-04-003 | E2E: Usuario tenta acessar URL proibida via navegador — ve tela 403 amigavel | E2E | Login como OPERATOR | 1. Navegar para /admin/plans/create | Tela 403 exibida: "Voce nao tem permissao para acessar esta area." Botao "Voltar ao Dashboard" presente | Planejado |
| TC-F03-04-004 | JWT com role modificado (tentativa de elevacao de privilegio) | Seguranca | JWT com role adulterado de OPERATOR para ADMIN_TENANT | 1. Forjar JWT com role=ADMIN_TENANT, assinatura invalida<br>2. Chamar endpoint protegido | Status 401 (assinatura invalida). Se assinatura valida mas claim divergente do banco: 403 | Planejado |
| TC-F03-04-005 | Acesso direto a URL proibida retorna 403 (nao 404) | Seguranca | JWT de OPERATOR | 1. GET /api/v1/plans/create (endpoint proibido) | Status 403 (RN12-01). Nunca 404 — nao revela existencia do recurso | Planejado |

### 3.13 F04-01: Login e Autenticacao

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F04-01-001 | JWT valido com claims corretas estabelece TenantContext | Unit | JwtAuthenticationFilter com JWT valido | 1. Extrair claims do JWT<br>2. Setar TenantContext | TenantContext contem tenant_id, user_id, roles, business_unit_ids, modules | Planejado |
| TC-F04-01-002 | JWT expirado retorna 401 | Unit | JWT com exp no passado | 1. Validar JWT no filtro | Filtro rejeita: 401 Unauthorized | Planejado |
| TC-F04-01-003 | JWT com assinatura invalida retorna 401 | Integracao | JWT adulterado | 1. GET /api/v1/tenants com JWT de assinatura invalida | Status 401. "Token invalido ou expirado" | Planejado |
| TC-F04-01-004 | Bloqueio temporario apos 5 tentativas incorretas (Keycloak) | Integracao | Keycloak container, usuario valido | 1. Tentar login 5x com senha errada<br>2. Tentar login com senha correta | Apos 5 tentativas: conta bloqueada por 15min. Sexta tentativa (correta) falha por bloqueio | Planejado |
| TC-F04-01-005 | E2E: Fluxo completo de login (sucesso e senha incorreta) | E2E | Keycloak + frontend | 1. Acessar pagina de login<br>2. Inserir email/senha corretos<br>3. Login realizado<br>4. Logout<br>5. Tentar login com senha incorreta | Login sucesso redireciona para dashboard. Login falha exibe mensagem "Email ou senha incorretos" | Planejado |
| TC-F04-01-006 | E2E: Recuperacao de senha envia link que expira em 1h | E2E | Keycloak + SMTP mock | 1. Clicar "Esqueci minha senha"<br>2. Inserir email<br>3. Verificar email recebido<br>4. Clicar link de recuperacao<br>5. Redefinir senha | Email enviado com link. Link funciona dentro de 1h. Nova senha permite login | Planejado |
| TC-F04-01-007 | JWT sem tenant_id claim — 401 | Seguranca | JWT valido mas sem claim tenant_id | 1. GET /api/v1/tenants com JWT sem tenant_id | Status 401. "Token nao contem identificacao do tenant" | Planejado |
| TC-F04-01-008 | JWT com role inexistente — 401 | Seguranca | JWT com role "SUPER_ADMIN" (invalido) | 1. GET /api/v1/tenants com role invalida | Status 401. Role nao reconhecida | Planejado |
| TC-F04-01-009 | Brute force: muitas requisicoes sem JWT — rate limiting | Seguranca | Nenhum token | 1. Enviar 100 requisicoes sem JWT em 10s | Apos N tentativas: rate limiting ativado (429 Too Many Requests) | Planejado |

### 3.14 F04-02: Onboarding Guiado

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F04-02-001 | Onboarding: passo 1 confirma dados do tenant | Unit | Tenant PENDING_ONBOARDING, OnboardingService | 1. Chamar OnboardingService.step1(tenantId, req) | Passo 1 concluido. Status do onboarding avanca. Tenant continua PENDING | Planejado |
| TC-F04-02-002 | Onboarding: passo 2 cria primeira BU como Matriz | Unit | Tenant PENDING_ONBOARDING, passo 1 concluido | 1. Chamar OnboardingService.step2(tenantId, buReq) | BU criada com parent_id = null (Matriz). Tenant continua PENDING | Planejado |
| TC-F04-02-003 | Onboarding: complete muda tenant para ACTIVE | Unit | Todos os passos concluidos | 1. Chamar OnboardingService.complete(tenantId) | Tenant.status = ACTIVE. Auditoria registra transicao PENDING_ONBOARDING->ACTIVE | Planejado |
| TC-F04-02-004 | Onboarding: tentar pular passo redireciona para passo correto | Unit | Tenant PENDING_ONBOARDING, passo 1 NAO concluido | 1. Chamar OnboardingService.step3(tenantId, req) | Lanca BusinessException: "Complete o passo 1 antes de continuar" (RN14-01) | Planejado |
| TC-F04-02-005 | GET /onboarding/status retorna progresso | Integracao | Testcontainers, cliente auth, onboarding em andamento | 1. GET /api/v1/onboarding/status | Status 200. OnboardingStatusResponse com stepsCompleted[], currentStep, progress (25%, 50%, 75%, 100%) | Planejado |
| TC-F04-02-006 | PATCH /onboarding/step-1 com dados validos retorna 200 | Integracao | Testcontainers, cliente auth, PENDING_ONBOARDING | 1. PATCH /api/v1/onboarding/step-1 com JSON valido | Status 200. StepCompleted. Progresso atualizado | Planejado |
| TC-F04-02-007 | POST /onboarding/step-2 cria BU Matriz | Integracao | Testcontainers, passo 1 concluido | 1. POST /api/v1/onboarding/step-2 com CNPJ valido | Status 201. BU criada com parent_id = null. TaxRegime valido | Planejado |
| TC-F04-02-008 | POST /onboarding/complete antes dos passos retorna 422 | Integracao | Testcontainers, cliente sem completar passos | 1. POST /api/v1/onboarding/complete | Status 422. Detail: "Complete todos os passos do onboarding antes de finalizar" | Planejado |
| TC-F04-02-009 | E2E: Fluxo completo de onboarding (4 passos) | E2E | Cliente logado, tenant PENDING_ONBOARDING | 1. Acessar /onboarding<br>2. Passo 1: confirmar dados<br>3. Passo 2: cadastrar Matriz (CNPJ)<br>4. Passo 3: configurar (se houver)<br>5. Passo 4: finalizar<br>6. Verificar dashboard do cliente | Barra de progresso visivel. Tenant transita para ACTIVE. Dashboard do cliente disponivel. Auditoria registra transicao | Planejado |
| TC-F04-02-010 | Tentativa de burlar onboarding via API — chamar endpoints antes de concluir | Seguranca | Cliente autenticado, PENDING_ONBOARDING | 1. Chamar GET /api/v1/dashboard/client/summary antes de concluir onboarding | Status 403. Detail: "Complete o onboarding antes de acessar o portal" (RN14-04) | Planejado |

### 3.15 F04-03: Dashboard do Cliente

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F04-03-001 | Dashboard cliente retorna cards: unidades, produtos, plano | Unit | Repositorio mockado com 2 BUs, 5 produtos, 1 assinatura | 1. Chamar DashboardService.getClientSummary() | Retorna: totalUnidades=2, totalProdutos=5, planName="Basico", subscriptionStatus=ACTIVE | Planejado |
| TC-F04-03-002 | Dashboard cliente exibe notificacoes e lembretes | Unit | Repositorio mockado com 2 notificacoes | 1. Chamar DashboardService.getNotifications() | Retorna lista de notificacoes com link para acao | Planejado |
| TC-F04-03-003 | GET /dashboard/client/summary retorna dados do cliente logado | Integracao | Testcontainers, cliente ACTIVE com dados | 1. GET /api/v1/dashboard/client/summary | Status 200. Campos: unidadesAtivas, produtosCatalogo, planoContratado, notificacoes | Planejado |
| TC-F04-03-004 | GET /dashboard/client/summary de outro tenant retorna vazio (isolamento) | Integracao | Testcontainers, cliente Tenant-A acessando | 1. Garantir que dashboard so retorna dados do tenant do JWT | NENHUM dado de Tenant-B no response (BR-NFR02) | Planejado |
| TC-F04-03-005 | E2E: Cliente ve dashboard apos onboarding completo | E2E | Cliente ACTIVE, dados seed | 1. Login como cliente<br>2. Visualizar dashboard | Cards: Unidades Ativas, Produtos no Catalogo, Plano Contratado. Notificacoes visiveis com link | Planejado |

### 3.16 F04-04: App Switcher

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F04-04-001 | Endpoint /auth/me retorna usuarios logado + modulos disponiveis | Unit | JWT com modules[] claim | 1. Chamar AuthService.me(userId) | Retorna: user, roles, businessUnits, modules[] | Planejado |
| TC-F04-04-002 | GET /auth/me retorna modulos do plano + autorizados | Integracao | Testcontainers, usuario com 2 modulos habilitados | 1. GET /api/v1/auth/me | Status 200. Modules listados. Usuario sem modulo nao acessa portal | Planejado |
| TC-F04-04-003 | E2E: App Switcher exibe modulos disponiveis e permite troca | E2E | Usuario com 2 modulos (Dashboard, Produtos) | 1. Login como Admin Tenant<br>2. Ver App Switcher no header<br>3. Trocar de modulo | Switcher exibe modulos disponiveis. Troca atualiza menu lateral e conteudo. Visivel mesmo com 1 modulo | Planejado |
| TC-F04-04-004 | JWT sem modules[] claim tem acesso negado a modulos | Seguranca | JWT sem claim modules | 1. Tentar acessar recurso de modulo nao listado | Status 403. Usuario sem modulo nao acessa o portal (F03-03) | Planejado |

### 3.17 F04-05: Unidades de Negocio

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F04-05-001 | Criar BU com CNPJ valido e unico no tenant | Unit | Repositorio mockado, CNPJ unico | 1. Chamar BusinessUnitService.create(req) | BU criada com status ACTIVE. parent_id = null (Matriz) | Planejado |
| TC-F04-05-002 | Criar BU com CNPJ duplicado no tenant retorna erro | Unit | Repositorio mockado com CNPJ ja existente | 1. Chamar BusinessUnitService.create(req) com CNPJ duplicado | Lanca DuplicateCnpjException. Mensagem: "CNPJ ja cadastrado para este tenant" (RN17-01) | Planejado |
| TC-F04-05-003 | Soft delete de BU libera CNPJ para reuso | Unit | BU deletada (deleted_dt preenchido) | 1. Chamar BusinessUnitService.create(req) com CNPJ da BU deletada | BU criada com sucesso. CNPJ reutilizado (RN17-01) | Planejado |
| TC-F04-05-004 | Criar BU filha com parent_id valido cria hierarquia | Unit | BU Matriz existente | 1. Chamar BusinessUnitService.create(req) com parent_id da Matriz | BU criada com parent_id = id da Matriz. Hierarquia correta | Planejado |
| TC-F04-05-005 | POST /business-units retorna 201 | Integracao | Testcontainers, Admin Tenant auth, CNPJ valido | 1. POST /api/v1/business-units com JSON valido | Status 201. BusinessUnitResponse com status ACTIVE | Planejado |
| TC-F04-05-006 | POST /business-units com CNPJ invalido retorna 400 | Integracao | Testcontainers, Admin Tenant auth | 1. POST /api/v1/business-units com CNPJ "00.000.000/0000-00" | Status 400. Detail: "CNPJ invalido" | Planejado |
| TC-F04-05-007 | POST /business-units/{id}/deactivate faz soft delete | Integracao | Testcontainers, BU ACTIVE | 1. POST /api/v1/business-units/{id}/deactivate | Status 200. deleted_dt preenchido. Indice unico parcial permite reuso do CNPJ | Planejado |
| TC-F04-05-008 | Hierarquia: tree de BUs retorna estrutura aninhada | Integracao | Testcontainers com Matriz + 3 filiais | 1. GET /api/v1/business-units | Arvore hierarquica: Matriz (parent_id=null) com children[] contendo filiais | Planejado |
| TC-F04-05-009 | E2E: Admin cria BU Matriz + filiais, desativa uma, recria com mesmo CNPJ | E2E | Admin Tenant logado | 1. Criar BU Matriz (CNPJ-A)<br>2. Criar BU Filial (CNPJ-B)<br>3. Desativar Filial<br>4. Criar nova BU com CNPJ-B | CNPJ-B reutilizado apos soft delete. Hierarquia exibida corretamente | Planejado |

### 3.18 F04-06: Catalogo de Produtos/Servicos

| ID | Descricao | Nivel | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-F04-06-001 | Criar produto vinculado a BU ativa | Unit | Repositorio mockado, BU ACTIVE | 1. Chamar ProductService.create(req) | Produto criado com status ACTIVE, vinculado a BU | Planejado |
| TC-F04-06-002 | Criar produto com SKU duplicado na mesma BU retorna erro | Unit | Repositorio mockado com SKU existente na BU | 1. Chamar ProductService.create(req) com mesmo SKU | Lanca BusinessException: "SKU ja cadastrado para esta unidade" | Planejado |
| TC-F04-06-003 | Soft delete de produto nao remove registro | Unit | Product ACTIVE | 1. Chamar ProductService.deactivate(id) | deleted_dt preenchido. Produto nao aparece em queries SELECT mas esta no banco (RN18-04) | Planejado |
| TC-F04-06-004 | POST /products retorna 201 | Integracao | Testcontainers, Admin/Manager auth | 1. POST /api/v1/products com JSON valido (name, type, business_unit_id) | Status 201. ProductResponse com status ACTIVE | Planejado |
| TC-F04-06-005 | GET /products filtra por BU | Integracao | Testcontainers com 2 BUs e produtos | 1. GET /api/v1/products?business_unit_id={buId} | Retorna apenas produtos da BU especificada | Planejado |
| TC-F04-06-006 | E2E: Admin cria produto na BU, lista catalogo, desativa produto | E2E | Admin/Manager logado, BU ativa | 1. Criar produto com SKU unico<br>2. Listar catalogo<br>3. Desativar produto<br>4. Verificar lista | Produto aparece no catalogo. Apos desativacao: nao aparece em consultas, mas esta no audit_log. Indicador "Nao mapeado" visivel se sem tributacao | Planejado |
| TC-F04-06-007 | OPERATOR tenta criar produto — 403 | Seguranca | JWT de OPERATOR | 1. POST /api/v1/products como OPERATOR | Status 403. Apenas Admin/Manager podem criar | Planejado |

---

## 4. Testes de Seguranca

### 4.1 RBAC — Matriz RN10-01: Cada Papel x Endpoint Proibido

Teste parametrizado automatizado que percorre a matriz completa:

```java
@ParameterizedTest
@CsvSource({
    "OPERATOR, POST, /api/v1/tenants",
    "OPERATOR, POST, /api/v1/plans",
    "OPERATOR, POST, /api/v1/users",
    "OPERATOR, PATCH, /api/v1/products",
    "OPERATOR, POST, /api/v1/business-units",
    "AUDITOR, POST, /api/v1/tenants",
    "AUDITOR, PATCH, /api/v1/tenants",
    "AUDITOR, POST, /api/v1/plans",
    "AUDITOR, POST, /api/v1/users",
    "MANAGER, POST, /api/v1/tenants",
    "MANAGER, POST, /api/v1/plans",
    // ... matriz completa (20+ combinacoes)
})
void shouldReturn403ForForbiddenAction(Role role, String method, String path) {
    // Configurar JWT com role especifico
    // Chamar endpoint
    // Verificar 403
}
```

| ID | Descricao | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---:|
| TC-SEC-RBAC-001 | Teste parametrizado: 20+ combinacoes papel x endpoint proibido | Testcontainers + JWT para cada role | 1. Para cada (role, methodo, path) na matriz RN10-01: configurar JWT, chamar endpoint | Status 403. Resposta RFC 7807. Mensagem "Acesso negado" | Planejado |
| TC-SEC-RBAC-002 | Admin FBSO pode acessar todos os endpoints | JWT Admin FBSO | 1. Para cada endpoint REST: chamar com JWT Admin FBSO | Status 200/201 (sucesso) — Admin FBSO tem permissao total | Planejado |
| TC-SEC-RBAC-003 | Auditor tentando DELETE em qualquer recurso retorna 403 | JWT AUDITOR | 1. Para cada DELETE endpoint: chamar como AUDITOR | 403 em todos | Planejado |
| TC-SEC-RBAC-004 | OPERATOR tentando POST/PATCH/PATCH/DELETE em recurso admin retorna 403 | JWT OPERATOR | 1. Para cada endpoint de escrita: chamar como OPERATOR | 403 em todos. OPERATOR so le | Planejado |

### 4.2 Multi-Tenant — Isolamento entre Tenants

| ID | Descricao | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---:|
| TC-SEC-MT-001 | Query cross-tenant retorna vazio (nao 403) | Tenant-A e Tenant-B com dados isolados | 1. Configurar TenantContext = tenant-A<br>2. SELECT * FROM business_unit | Retorna apenas BUs de tenant-A. Dados de tenant-B nao visiveis | Planejado |
| TC-SEC-MT-002 | Tentar acessar recurso de outro tenant por ID direto retorna 404 | Tenant-A, recurso de tenant-B | 1. GET /api/v1/tenants/{tenant-B-id} como usuario tenant-A | Status 404 (nao 403 — tenant nao sabe que o outro existe) | Planejado |
| TC-SEC-MT-003 | Injetar tenant_id de outro tenant no JWT adulterado | JWT valido de tenant-A com tenant_id de tenant-B | 1. Forjar JWT: manter assinatura valida, alterar tenant_id claim | Se assinatura invalida: 401. Se assinatura valida mas claim divergente: comportamento por design — idealmente 403 | Planejado |
| TC-SEC-MT-004 | Admin FBSO (cross-tenant) pode ver todos os tenants | JWT Admin FBSO (sem tenant_id especifico) | 1. GET /api/v1/tenants como Admin FBSO | Retorna TODOS os tenants. Admin FBSO e o unico papel com visao global | Planejado |
| TC-SEC-MT-005 | PostgreSQL RLS: INSERT com tenant_id divergente do app.current_tenant_id é REJEITADO | PostgreSQL RLS ativo, TenantContext com tenant-A | 1. SET app.current_tenant_id = 'tenant-A-uuid'<br>2. INSERT INTO tenant (...) VALUES (tenant_id='tenant-B-uuid', ...) | PostgreSQL rejeita com POLICY violation. RLS impede INSERT cross-tenant mesmo se query forjar tenant_id | Implementado (estrutural) |
| TC-SEC-MT-006 | PostgreSQL RLS: SELECT sem WHERE tenant_id é automaticamente filtrado (não retorna dados de outros tenants) | PostgreSQL RLS ativo, dados de tenant-A e tenant-B | 1. SET app.current_tenant_id = 'tenant-A-uuid'<br>2. SELECT * FROM business_unit (sem WHERE) | Retorna apenas BUs de tenant-A. RLS aplica filtro implícito. Zero vazamento | Implementado (estrutural) |
| TC-SEC-MT-007 | PostgreSQL RLS: tentativa de UPDATE em registro de outro tenant falha | PostgreSQL RLS ativo | 1. SET app.current_tenant_id = 'tenant-A-uuid'<br>2. UPDATE business_unit SET corporate_name='HACK' WHERE tenant_id='tenant-B-uuid' | PostgreSQL: 0 rows affected. Nenhum registro de tenant-B é alterado | Implementado (estrutural) |

### 4.3 OWASP Top 10 — Cenarios Especificos

| ID | Categoria OWASP | Descricao | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---|:---:|
| TC-SEC-OWASP-001 | **A1 - SQL Injection** | Tentativa de SQL injection no campo de busca textual | Endpoint GET /tenants?search= | 1. GET /api/v1/tenants?search=' OR 1=1 --<br>2. GET /api/v1/tenants?search=' UNION SELECT * FROM pg_catalog.pg_tables --<br>3. GET /api/v1/tenants?search='; DROP TABLE tenant; -- | Parametro tratado como string literal (JDBC PreparedStatement). Nenhum SQL injection bem-sucedido. Query retorna dados filtrados ou vazios, sem erros de banco | Planejado |
| TC-SEC-OWASP-002 | **A1 - SQL Injection** | SQL injection em campo de ID (UUID) | Endpoint GET /tenants/{id} | 1. GET /api/v1/tenants/' OR '1'='1<br>2. GET /api/v1/tenants/'; SELECT pg_sleep(10); -- | Status 400 (UUID invalido). Sem time-based injection | Planejado |
| TC-SEC-OWASP-003 | **A2 - Broken Authentication** | JWT com assinatura invalida (tentativa de forjar token) | JWT adulterado | 1. Gerar JWT com payload modificado, assinatura aleatoria<br>2. GET /api/v1/dashboard/admin/summary | Status 401. JwtAuthenticationFilter detecta assinatura invalida | Planejado |
| TC-SEC-OWASP-004 | **A2 - Broken Authentication** | JWT com "alg": "none" (tentativa de bypass) | JWT com alg=none | 1. Criar JWT com header {alg: "none"}, payload valido, sem assinatura<br>2. GET /api/v1/tenants | Status 401. Filtro rejeita tokens sem assinatura | Planejado |
| TC-SEC-OWASP-005 | **A2 - Broken Authentication** | Replay attack: reutilizar JWT apos logout | JWT valido, usuario faz logout | 1. Login, obter JWT<br>2. Logout (Keycloak invalida sessao)<br>3. Reutilizar JWT antigo | Status 401. JWT deve ser verificado contra blacklist ou expiracao curta | Planejado |
| TC-SEC-OWASP-006 | **A2 - Broken Authentication** | JWT com expiracao muito longa (7 dias+ para admin) | JWT com exp = now + 30 dias | 1. Validar JWT no filtro | Filtro pode aceitar (se dentro do exp), mas politica recomendada: exp <= 24h para admin | Planejado |
| TC-SEC-OWASP-007 | **A3 - XSS** | Tentativa de XSS stored em campos de texto | POST /business-units com nome contendo HTML/JS | 1. POST /api/v1/business-units com corporate_name="<script>alert('XSS')</script>"<br>2. GET /api/v1/business-units | Nome armazenado como texto puro. HTML/JS escapado na resposta JSON. Nao executavel no frontend | Planejado |
| TC-SEC-OWASP-008 | **A3 - XSS** | Tentativa de XSS reflected em query params | GET /tenants com search contendo script | 1. GET /api/v1/tenants?search=<script>alert('XSS')</script> | Parametro tratado como dado, nao executado. Resposta JSON escapa caracteres especiais | Planejado |
| TC-SEC-OWASP-009 | **A5 - Broken Access Control** | Usuario de baixo privilegio acessa endpoint de alta permisao via method smuggling | JWT OPERATOR, varios endpoints | 1. OPERATOR tenta GET com header malicioso (ex: X-HTTP-Method-Override: POST)<br>2. OPERATOR tenta OPTIONS, TRACE, etc. em endpoints restritos | Spring Security bloqueia method override. Acesso negado | Planejado |
| TC-SEC-OWASP-010 | **A5 - Broken Access Control** | Inclusao de usuario sem permissao em modulo restrito (IDOR) | JWT OPERATOR | 1. PUT /api/v1/users/{uid}/permissions com role=ADMIN_TENANT | Status 403. OPERATOR nao pode alterar permissoes | Planejado |
| TC-SEC-OWASP-011 | **A6 - Security Misconfiguration** | Respostas de erro expoem stack traces | Nenhum | 1. POST /api/v1/tenants com JSON invalido (tipo errado)<br>2. GET /api/v1/tenants/{id-invalido} | Erro segue RFC 7807. SEM stack trace. SEM rastros de implementacao. Mensagens em PT-BR (BR-NFR07, BR-NFR08) | Planejado |
| TC-SEC-OWASP-012 | **A6 - Security Misconfiguration** | Headers de seguranca ausentes | Nenhum | 1. GET qualquer endpoint<br>2. Verificar response headers | Headers: X-Content-Type-Options: nosniff, X-Frame-Options: DENY, Strict-Transport-Security, Cache-Control | Planejado |
| TC-SEC-OWASP-013 | **A8 - CSRF** | CORS permite origens nao autorizadas | Nenhum | 1. OPTIONS /api/v1/tenants com Origin: https://malicious-site.com | Origin nao autorizada deve ser bloqueada. Apenas origem do frontend permitida | Planejado |
| TC-SEC-OWASP-014 | **A8 - CSRF** | Requisicao POST sem token CSRF (se aplicavel) | Nenhum | 1. POST /api/v1/tenants com Cookie de sessao mas sem CSRF token | Se CSRF habilitado: 403. Se stateless JWT: nao aplicavel (verificar configuracao) | Planejado |
| TC-SEC-OWASP-015 | **A9 - Using Components with Known Vuln** | Dependencias com CVEs conhecidas | pom.xml | 1. Rodar mvn dependency-check:check | Zero vulnerabilidades criticas ou high. Medium aceitas com justificativa | Planejado |
| TC-SEC-OWASP-016 | **A10 - Insufficient Logging & Monitoring** | Acoes suspeitas nao registradas | Nenhum | 1. Tentativas de acesso negado (403)<br>2. Tentativas de SQL injection<br>3. Tentativas de autenticacao falha | Todas registradas em audit_log. Action = "ACCESS_DENIED", entity_type = "SECURITY_EVENT" | Planejado |

### 4.4 Data Leakage — Exposicao de Dados Sensiveis

| ID | Descricao | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---:|
| TC-SEC-DL-001 | Logs nao contem dados pessoais em texto claro | Logger configurado com filtro | 1. Executar operacoes com dados pessoais (nome, email, CNPJ)<br>2. Verificar logs | Dados pessoais mascarados: "joao.silva@email.com" -> "joa***@email.com". CNPJ: "**.***.***/0001-**" (NFR-LGPD) | Planejado |
| TC-SEC-DL-002 | Resposta HTTP nao expoe dados sensiveis de outros tenants | Testcontainers multi-tenant | 1. GET /api/v1/users como usuario tenant-A | Response nao contem emails, nomes ou dados de usuarios de tenant-B | Planejado |
| TC-SEC-DL-003 | Erro 500 nao expoe detalhes de implementacao | Exception nao tratada | 1. Provocar erro interno (ex: constraint violation)<br>2. Verificar response | Erro segue RFC 7807. Sem stack trace, sem query SQL, sem nomes de tabelas. Apenas "Erro interno do servidor" | Planejado |
| TC-SEC-DL-004 | Respostas de lista paginada nao expoem total geral se configurado | Configuracao de seguranca | 1. GET /api/v1/tenants com paginacao | Verificar se totalElements nao vaza informacao de volume (relevante para alguns contextos) | Planejado |

---

## 5. Testes de Performance (NFRs)

### 5.1 Cenario de Carga: Dashboard Admin

| ID | Descricao | Pre-condicao | Passos | Metricas | Criterio de Aceite | Status |
|:---|:---|:---|:---|:---|:---:|:---:|
| TC-PERF-001 | Dashboard summary com 1000 tenants (concorrencia: 50 usuarios) | 1000 tenants no banco, 50 threads simultaneas | 1. JMeter: 50 usuarios, ramp-up 10s<br>2. GET /api/v1/dashboard/admin/summary<br>3. Coletar p95 e p99 | p95 latency <= 3s, p99 <= 5s, 0 erros | BR-NFR05 | Planejado |
| TC-PERF-002 | Dashboard evolution com 12 meses de dados (1000 tenants) | 12 meses de dados historicos | 1. JMeter: 20 usuarios simultaneos<br>2. GET /api/v1/dashboard/admin/evolution?period=ano_atual | p95 latency <= 3s | BR-NFR05 | Planejado |
| TC-PERF-003 | Lista paginada de tenants com 1000 registros | 1000 tenants no banco | 1. JMeter: 30 usuarios simultaneos<br>2. GET /api/v1/tenants?page=0&size=25 | p95 latency <= 2s. Paginacao responsiva | BR-NFR05 | Planejado |
| TC-PERF-004 | Criacao simultanea de 50 tenants (concorrencia) | Banco vazio | 1. JMeter: 50 threads, cada uma POST /tenants<br>2. Verificar consistencia | 50 tenants criados sem duplicatas. Sem deadlocks. Nome unico respeitado | — | Planejado |
| TC-PERF-005 | Upsert de permissao com 1000 usuarios vinculados | 1000 usuarios, 10 BUs | 1. JMeter: 10 threads, PUT /users/{uid}/permissions | p95 <= 2s. Vinculacao consistente | — | Planejado |

### 5.2 Cenario de Concorrencia / Race Condition

| ID | Descricao | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---:|
| TC-RACE-001 | Duas assinaturas simultaneas para mesmo tenant (race condition) | Tenant sem assinatura | 1. Thread-1: POST subscription<br>2. Thread-2: POST subscription (mesmo instante) | Uma retorna 201, outra retorna 409. Banco consistente — 1 subscription ACTIVE (RN07-01) | Planejado |
| TC-RACE-002 | Duas atualizacoes simultaneas de preco do plano | Plano existente | 1. Thread-1: PATCH plan/price=100<br>2. Thread-2: PATCH plan/price=200 (mesmo instante) | Ambas atualizacoes aplicadas em sequencia (versao incrementada 2x). Sem perda de dado | Planejado |
| TC-RACE-003 | Suspensao e reativacao simultanea de tenant | Tenant ACTIVE | 1. Thread-1: POST suspend<br>2. Thread-2: POST reactivate (mesmo instante) | Estado final consistente. Auditoria registra ambas operacoes na ordem correta | Planejado |

---

## 6. Testes de Regressao

### 6.1 Checklist de Features Ja Homologadas

> A ser preenchido a cada marco (M2 a M7). Ao final do M7, todas as 18 features devem estar marcadas.

| Feature | Marco | Testada em | Status M2 | Status M3 | Status M4 | Status M5 | Status M6 | Status M7 |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| F01-01 — Dashboard Admin | M2 | | ⬜ | — | — | — | — | ⬜ |
| F01-02 — Lista de Contas | M2 | | ⬜ | — | — | — | — | ⬜ |
| F01-03 — Alertas Dashboard | M2 | | ⬜ | — | — | — | — | ⬜ |
| F02-01 — Criar Tenant | M3 | | — | ⬜ | — | — | — | ⬜ |
| F02-02 — Transicoes Status | M3 | | — | ⬜ | — | — | — | ⬜ |
| F02-03 — Planos | M3 | | — | ⬜ | — | — | — | ⬜ |
| F02-04 — Assinaturas | M3 | | — | ⬜ | — | — | — | ⬜ |
| F02-05 — Auditoria | M3 | | — | ⬜ | — | — | — | ⬜ |
| F03-01 — Gestao Usuarios | M4 | | — | — | ⬜ | — | — | ⬜ |
| F03-02 — RBAC | M4 | | — | — | ⬜ | — | — | ⬜ |
| F03-03 — Vinculacao UxM | M4 | | — | — | ⬜ | — | — | ⬜ |
| F03-04 — Acesso Condicional | M4 | | — | — | ⬜ | — | — | ⬜ |
| F04-01 — Login | M5 | | — | — | — | ⬜ | — | ⬜ |
| F04-02 — Onboarding | M5 | | — | — | — | ⬜ | — | ⬜ |
| F04-03 — Dashboard Cliente | M5 | | — | — | — | ⬜ | — | ⬜ |
| F04-04 — App Switcher | M5 | | — | — | — | ⬜ | — | ⬜ |
| F04-05 — Business Units | M6 | | — | — | — | — | ⬜ | ⬜ |
| F04-06 — Produtos | M6 | | — | — | — | — | ⬜ | ⬜ |

### 6.2 Script de Regressao Automatico (M7)

```bash
#!/bin/bash
# Regressao completa — executa antes de cada release

echo "=== REGRESSAO: ms-fbso-platform-admin ==="

# 1. Testes unitarios
./mvnw test -Dtest="**/unit/**" && echo "OK: Unitarios" || exit 1

# 2. Testes de integracao
./mvnw test -Dtest="**/integration/**" -Dspring.profiles.active=test && echo "OK: Integracao" || exit 1

# 3. Testes de seguranca
./mvnw test -Dtest="**/security/**" && echo "OK: Seguranca" || exit 1

# 4. Testes E2E (requer Docker Compose)
docker compose -f infra/docker/docker-compose.test.yml up -d
./mvnw test -Dtest="**/e2e/**" && echo "OK: E2E" || exit 1

# 5. Cobertura
./mvnw jacoco:check && echo "OK: Cobertura >= 80%" || exit 1

# 6. Qualidade
./mvnw checkstyle:check pmd:check && echo "OK: Qualidade" || exit 1

echo "=== REGRESSAO COMPLETA ==="
```

### 6.3 Smoke Tests Pos-Deploy

| ID | Verificacao | Comando / Acao | Resultado Esperado | Status |
|:---|:---|:---|:---|:---:|
| SMOKE-01 | Health check | `curl /actuator/health` | Status 200. Body: {"status":"UP"} | Planejado |
| SMOKE-02 | Autenticacao | `curl /api/v1/tenants -H "Authorization: Bearer <jwt>"` | Status 200. Lista de tenants retornada | Planejado |
| SMOKE-03 | Erro 401 sem token | `curl /api/v1/tenants` | Status 401 | Planejado |
| SMOKE-04 | Erro 403 sem permissao | `curl /api/v1/plans -H "Authorization: Bearer <jwt-operator>"` | Status 403 | Planejado |
| SMOKE-05 | RFC 7807 | `curl /api/v1/tenants/99999` | Status 404. Body no formato RFC 7807 | Planejado |
| SMOKE-06 | Banco conectado | `curl /actuator/health` | "db" no health check mostra "UP" | Planejado |

---

## 7. Pipeline de Seguranca — Teste Integrado Fim-a-Fim

| ID | Descricao | Pre-condicao | Passos | Resultado Esperado | Status |
|:---|:---|:---|:---|:---|:---:|
| TC-PIPE-001 | Pipeline completa: JWT -> TenantContext -> RBAC -> TenantIsolation -> Auditoria | Testcontainers + Keycloak, Admin FBSO | 1. Enviar requisicao com JWT valido<br>2. Verificar fluxo em cada estagio | JWT Filter extrai claims. TenantContext setado. RbacAspect permite. TenantIsolation filtra. AuditAspect registra | Planejado |
| TC-PIPE-002 | Pipeline: requisicao sem JWT para no filtro | Nenhum token | 1. Enviar requisicao sem Authorization | JWTFilter retorna 401. RbacAspect e TenantIsolation NAO executam (curto-circuito seguro) | Planejado |
| TC-PIPE-003 | Pipeline: JWT invalido para no filtro | JWT com assinatura invalida | 1. Enviar requisicao com JWT invalido | 401. Nenhum aspecto apos o filtro executa | Planejado |
| TC-PIPE-004 | Pipeline: sem tenant_id no JWT | JWT valido sem claim tenant_id | 1. Enviar requisicao | 401. TenantIsolationAspect nunca executado sem tenant_id | Planejado |
| TC-PIPE-005 | Conformidade RFC 7807 em todos os endpoints de erro | Testcontainers | 1. Para cada cenario de erro (400, 401, 403, 404, 409, 422, 500): verificar response | Resposta contem type, title, status, detail. Sem stack trace. Mensagens em PT-BR | Planejado |

---

## 8. Ferramentas e Configuracao

### 8.1 Dependencias de Teste (pom.xml)

```xml
<!-- Testes Unitarios -->
<dependency>
    <groupId>org.junit.jupiter</groupId>
    <artifactId>junit-jupiter</artifactId>
    <scope>test</scope>
</dependency>
<dependency>
    <groupId>org.mockito</groupId>
    <artifactId>mockito-core</artifactId>
    <scope>test</scope>
</dependency>
<dependency>
    <groupId>org.mockito</groupId>
    <artifactId>mockito-junit-jupiter</artifactId>
    <scope>test</scope>
</dependency>

<!-- Testes de Integracao -->
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-test</artifactId>
    <scope>test</scope>
</dependency>
<dependency>
    <groupId>org.testcontainers</groupId>
    <artifactId>testcontainers</artifactId>
    <scope>test</scope>
</dependency>
<dependency>
    <groupId>org.testcontainers</groupId>
    <artifactId>postgresql</artifactId>
    <scope>test</scope>
</dependency>
<dependency>
    <groupId>org.testcontainers</groupId>
    <artifactId>junit-jupiter</artifactId>
    <scope>test</scope>
</dependency>

<!-- Testes E2E -->
<dependency>
    <groupId>com.microsoft.playwright</groupId>
    <artifactId>playwright</artifactId>
    <scope>test</scope>
</dependency>

<!-- Testes de Seguranca -->
<dependency>
    <groupId>io.rest-assured</groupId>
    <artifactId>rest-assured</artifactId>
    <scope>test</scope>
</dependency>
```

### 8.2 Plugins de Qualidade

```xml
<plugin>
    <groupId>org.jacoco</groupId>
    <artifactId>jacoco-maven-plugin</artifactId>
    <configuration>
        <excludes>
            <exclude>**/config/**</exclude>
            <exclude>**/FbsoPlatformAdminApplication.class</exclude>
        </excludes>
    </configuration>
    <executions>
        <execution>
            <goals><goal>prepare-agent</goal></goals>
        </execution>
        <execution>
            <id>check</id>
            <goals><goal>check</goal></goals>
            <configuration>
                <rules>
                    <rule>
                        <element>BUNDLE</element>
                        <limits>
                            <limit><counter>LINE</counter><value>COVEREDRATIO</value><minimum>0.80</minimum></limit>
                            <limit><counter>BRANCH</counter><value>COVEREDRATIO</value><minimum>0.70</minimum></limit>
                        </limits>
                    </rule>
                </rules>
            </configuration>
        </execution>
    </executions>
</plugin>
<plugin>
    <groupId>org.apache.maven.plugins</groupId>
    <artifactId>maven-checkstyle-plugin</artifactId>
    <configuration>
        <failOnViolation>true</failOnViolation>
    </configuration>
</plugin>
```

---

## 9. Testes de Infraestrutura e Build (Sprints 1-2)

> **Adicionado v2.1 — 14/07/2026.** Esta secao cobre testes estruturais executados nas Sprints 1 (Setup) e 2 (Seguranca) que nao se enquadram no modelo feature-oriented das secoes §3-§7.

### 9.1 Build e Compilacao

| ID | Descricao | Nivel | Criterio | Sprint |
|:---|:---|:---|:---|:---:|
| TC-INFRA-001 | `mvn clean compile` executa sem erros | Build | Zero erros de compilacao. Todos os pacotes e imports resolvidos | 1 |
| TC-INFRA-002 | `mvn clean install` executa sem erros | Build | BUILD SUCCESS. JAR gerado em `target/` | 1 |
| TC-INFRA-003 | `mvn test` executa suite vazia (placeholder inicial) | Build | Zero falhas. Surefire configurado | 1 |
| TC-INFRA-004 | Estrutura de pacotes confere com ARCHITECTURE.md §2 | Estrutural | 14 pacotes criados. Classes esqueleto presentes em todos os pacotes | 1 |

### 9.2 Migrations Flyway

| ID | Descricao | Nivel | Criterio | Sprint |
|:---|:---|:---|:---|:---:|
| TC-INFRA-005 | `mvn flyway:migrate` cria schema `fbso_platform` | Integracao | Schema existe no PostgreSQL. `\dn fbso_platform` | 1 |
| TC-INFRA-006 | V001 cria 11 tabelas com colunas, PKs, FKs, NOT NULL | Integracao | `\dt fbso_platform.*` lista 11 tabelas. Constraints verificadas via `information_schema` | 1 |
| TC-INFRA-007 | V001 rollback desfaz migracao | Integracao | `mvn flyway:undo` reverte schema. DROP manual como fallback | 1 |
| TC-INFRA-008 | V002 cria indices unicos parciais (`WHERE deleted_dt IS NULL`) | Integracao | Indices: unique_cnpj_active, unique_email_active, unique_sku_active | 1 |
| TC-INFRA-009 | Colunas de auditoria presentes em TODAS as tabelas | Integracao | Toda tabela tem: created_dt, updated_dt, created_by, updated_by, deleted_dt, deleted_by | 1 |

### 9.3 BaseRepository

| ID | Descricao | Nivel | Criterio | Sprint |
|:---|:---|:---|:---|:---:|
| TC-INFRA-010 | `findAll()` injeta `WHERE deleted_dt IS NULL` automaticamente | Unit | Query gerada contem `deleted_dt IS NULL`. TenantContext mockado | 1 |
| TC-INFRA-011 | `softDelete()` seta `deleted_dt = NOW()` e `deleted_by` | Unit | Registro nao aparece em `findAll()` subsequente. Continua no banco | 1 |
| TC-INFRA-012 | `save()` preenche `created_by`/`updated_by` automaticamente | Unit | Campos de auditoria preenchidos a partir do TenantContext | 1 |

### 9.4 Validacao de Enums e Entidades Base

| ID | Descricao | Nivel | Criterio | Sprint |
|:---|:---|:---|:---|:---:|
| TC-INFRA-013 | Todos os 8 enums compilam com valores corretos | Unit | TenantStatus, TenantSegment, Recurrence, SubscriptionStatus, UserStatus, Role, TaxRegime, ProductType | 1 |
| TC-INFRA-014 | Enum `TenantStatus` contem: PENDING_ONBOARDING, ACTIVE, SUSPENDED, INACTIVE | Unit | 4 valores. Ordem preservada | 1 |
| TC-INFRA-015 | Enum `Role` contem: ADMIN_TENANT, MANAGER_BU, OPERATOR_BU, AUDITOR | Unit | 4 valores. Compativel com matriz RN10-01 | 1 |

### 9.5 Configuracao de Profiles

| ID | Descricao | Nivel | Criterio | Sprint |
|:---|:---|:---|:---|:---:|
| TC-INFRA-016 | Profile `dev` carrega `application-dev.yml` | Integracao | Datasource aponta para `localhost:5432`. Logging level DEBUG | 1 |
| TC-INFRA-017 | Profile `prod` carrega `application-prod.yml` | Integracao | Datasource via `DATASOURCE_URL`. Logging level WARN | 1 |

### 9.6 Pipeline de Seguranca — Testes Estruturais

| ID | Descricao | Nivel | Criterio | Sprint |
|:---|:---|:---|:---|:---:|
| TC-INFRA-018 | `JwtAuthenticationFilter` registrado no Spring Security filter chain | Unit | Filter aparece na cadeia. Ordem correta (antes de qualquer filtro de autorizacao) | 2 |
| TC-INFRA-019 | `TenantContext.clear()` invocado no finally do filter | Unit | ThreadLocal limpo apos response. Verificar com teste de concorrencia | 2 |
| TC-INFRA-020 | `@RequiresPermission` anotacao compilada com resource e action | Unit | Anotacao acessivel via reflection. RetentionPolicy.RUNTIME | 2 |
| TC-INFRA-021 | `@Auditable` anotacao compilada com entityType e action | Unit | Anotacao acessivel via reflection. @Async configurado no aspecto | 2 |

### 9.7 PostgreSQL Row-Level Security

| ID | Descricao | Nivel | Criterio | Sprint |
|:---|:---|:---|:---|:---:|
| TC-INFRA-022 | Migration V003: RLS habilitado em 5 tabelas | Integracao | `SELECT tablename FROM pg_tables WHERE schemaname='fbso_platform' AND rowsecurity=true` retorna 5 tabelas | 2 |
| TC-INFRA-023 | Politica tenant_isolation criada em cada tabela | Integracao | `SELECT policyname FROM pg_policies WHERE schemaname='fbso_platform' AND policyname='tenant_isolation'` retorna 5 politicas | 2 |
| TC-INFRA-024 | `TenantAwareDataSource` configura `app.current_tenant_id` em cada `getConnection()` | Unit | DataSource proxy intercepta getConnection() e executa SET com tenant_id do TenantContext | 2 |
| TC-INFRA-025 | Rollback V003 remove politicas RLS | Integracao | `mvn flyway:undo` (V003) remove politicas e desabilita RLS. `rowsecurity=false` em todas as tabelas | 2 |
| TC-INFRA-026 | Admin FBSO (sem tenant_id) → RESET app.current_tenant_id (acesso global) | Unit | TenantContext sem tenant_id → DataSource proxy executa RESET, permitindo visao cross-tenant | 2 |

### Resumo da Secao

| Nivel | Cenarios |
|:---|:---:|
| Build | 4 |
| Integracao | 9 |
| Unit | 10 |
| Estrutural | 2 |
| **Total** | **25** |

> **Uso pelas Sprints:** Sprint 1 cobre TC-INFRA-001 a TC-INFRA-017. Sprint 2 adiciona TC-INFRA-018 a TC-INFRA-026.

---

## 10. Criterios de Aceite (Definition of Done — Testes)

Alem da [DoD do projeto](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/DEFINITION_OF_DONE.md), este plano adiciona:

- [ ] Cobertura de testes >= 80% (JaCoCo LINE), >= 70% (BRANCH)
- [ ] Zero warnings do Checkstyle/PMD no codigo de producao
- [ ] 100% dos 37 endpoints testados (integracao)
- [ ] 100% das 18 familias de RNs (PRD §6.6) cobertas por teste automatizado
- [ ] 100% da matriz RBAC (RN10-01) testada (cada papel x endpoint proibido)
- [ ] Isolamento Multi-Tenant verificado (query cross-tenant retorna vazio)
- [ ] Pipeline de seguranca testada (JWT -> TenantContext -> RBAC -> TenantIsolation -> Auditoria)
- [ ] Respostas de erro conforme RFC 7807 em todos os endpoints
- [ ] Zero dados pessoais em logs (NFR-LGPD)
- [ ] Carga do dashboard testada com 1000+ tenants (p95 <= 3s)
- [ ] Testes OWASP executados: SQL Injection, XSS, Broken Auth, CSRF
- [ ] Smoke test pos-deploy executado

---

## 11. Registro de Alteracoes

| Versao | Data | Alteracao | Autor |
|:---|:---|:---|:---|
| 2.5 | 16/07/2026 | v2.5 — 2 novos cenários (DT-002, DT-009), referência a débitos técnicos (IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md), total 126+25=176 | Time Técnico |
| 2.4 | 16/07/2026 | Sprint 3 iniciada (16/07/2026). Status atualizado para "Em Execucao". Cenários §3.1-§3.8 em validação. | Time Técnico |
| 2.3 | 15/07/2026 | Revisão Caveman (DOCS-SERVICE-CAVEMAN-REVIEW.md): Corrigido total de cenários (154→149+25=174, §2). Corrigida contagem RN (16→18 famílias, §1.3 e §10). Atualizada referência SPECS (v1.1→v1.4). | Caveman/IA |
| 2.1 | 14/07/2026 | Correcao pos-gate de artefatos de sprint (SPRINT_ARTEFACTS_FAIL_REPORT.md v1.0 — Sprint 1). NC-003/NC-004: adicionada §9 "Testes de Infraestrutura e Build" com 21 cenarios estruturais (TC-INFRA-001 a TC-INFRA-021) cobrindo Sprints 1-2. Secoes renumeradas: §9→§10, §10→§11 | Agente Corretor Sprint/IA |
| 2.0 | 14/07/2026 | Regeneracao completa baseada em SPECS.md v1.1. Correcao da piramide de testes (~70% Unit, ~20% Int, ~5% E2E, ~5% Seg). Adicao de 146 cenarios (124 por feature + 22 transversais). Cobertura das 18 features (incluindo F01-03, F04-03, F04-04). 16 cenarios E2E com 5 fluxos completos. 23 cenarios de seguranca (RBAC, Multi-Tenant, OWASP, Data Leak). 5 cenarios de carga. 3 cenarios de concorrencia. Mapa de cobertura Feature x Cenarios. Checklist de regressao por marco. Coluna Status (Planejado/Implementado/Aprovado/Falhou) em todos os cenarios | Agente Gerador TEST_PLAN |
| 1.0 | 13/07/2026 | Criacao inicial: 59 cenarios funcionais + 4 seguranca Multi-Tenant + 5 performance + matriz RBAC + checklist de regressao por marco | Time Tecnico |

---

## 12. Referencias

| Documento | Localizacao |
|:---|:---|
| Debitos Tecnicos Sprint 3 | [IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md](./IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md) |

---

🤖 *Documentacao gerada de forma automatizada pelo Agente: Gerador de TEST_PLAN. Foram utilizados os skills: test-strategy-design, qa-test-planner, acceptance-criteria, security-reviewer. v2.5 em 16/07/2026: Sprint 3 — debitos tecnicos incorporados.*
