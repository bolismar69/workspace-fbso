# SPECS.md — Especificação da Solução: ms-fbso-platform-admin

- **Solução:** `ms-fbso-platform-admin`
- **Tipo:** Backend
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Caffeine Cache
- **Projeto de Negócio:** [PRJ-FIN-2026-0003-SAAS-FBSO-ORG](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/)
- **Versão:** 2.4
- **Data:** 17 de Julho de 2026
- **Status:** Em Execução — Sprints 1-4 concluídas ✅ (Frentes 0-4). Sprint 5 Frente 0 concluída ✅. 18 endpoints REST. 213 testes. 21 RNs implementadas. RBAC DB-backed com matriz RN10-01 100% validada. Sprint 5 Frente 0: docker-compose dev (Keycloak 26 + PG 17 + MailHog), Flyway 10→12.11.0, PG driver→42.7.11, +OAuth2 Client. 95/160 tarefas (59%).
- **Origem:** [PRD.md](./PRD.md) + [ARCHITECTURE.md](./ARCHITECTURE.md)

---

## 1. Visão Geral da Solução

### 1.1 Propósito

O `ms-fbso-platform-admin` é o **backend do Core Administrativo da FBSO Platform**. Ele expõe a API REST consumida pelo frontend (`web_app-fbso-platform-portal`) e implementa toda a lógica de negócio para:

- **EP-01:** Dashboard administrativo com métricas operacionais do SaaS
- **EP-02:** Gestão de clientes (Tenants), planos comerciais e assinaturas
- **EP-03:** Governança de acessos — RBAC com 4 papéis, vinculação Usuário × Unidade × Módulo
- **EP-04:** Experiência do cliente — onboarding guiado, cadastro de Unidades de Negócio e Catálogo de Produtos

### 1.2 O Que Esta Solução Implementa

| Escopo | Entidades/Recursos |
|:---|:---|
| ✅ **Fase 0 (Core)** | 11 entidades: Tenant, Plan, PlanModule, Subscription, User, UserPermission, ResourceAction, RoleResource, BusinessUnit, ProductService, AuditEntry |
| ✅ **Fase 0 (Core)** | 11 recursos REST: `/tenants`, `/plans`, `/subscriptions`, `/users`, `/permissions`, `/business-units`, `/products`, `/dashboard/admin`, `/dashboard/client`, `/onboarding`, `/audit` |
| ✅ **Fase 0 (Core)** | Pipeline de segurança: JWT Filter → TenantContext → RBAC → PostgreSQL RLS → BaseRepository → Auditoria |
| ❌ **Fase Futura** | Tributali-Engine: Billable, ProductBillableMapping, SplitPayment |
| ❌ **Fase Futura** | Storekeeper: Order, Invoice, TransactionPayment, BillingInfo, BankAccount |

### 1.3 Relação com Entregas do Projeto

| Marco | Data | Entregas | O que esta solução entrega |
|:---|:---|:---|:---|
| M2 | 15/08/2026 | D1 — Portal Admin | 🔄 5 endpoints `GET /dashboard/admin/*` implementados (17/07/2026). 50 testes |
| M3 | 31/08/2026 | D2, D3 — Contas e Planos | ⬜ Pendente — CRUD `/tenants`, `/plans`, `/subscriptions`, `/audit` |
| M4 | 15/09/2026 | D4 — RBAC | CRUD `/users`, `/permissions` |
| M5 | 30/09/2026 | D5 — Portal Cliente (4 frentes, 36 tarefas) | `/onboarding`, `/dashboard/client` |
| M6 | 15/10/2026 | D6, D7 — BUs e Catálogo | CRUD `/business-units`, `/products` |
| M7 | 30/10/2026 | Homologação | Todos os 11 recursos |

---

## 2. Requisitos Funcionais

### 2.1 Mapeamento BR → Feature → User Stories

| BR | Descrição | Feature | User Stories |
|:---|:---|:---|:---|
| **BR-A01** | Dashboard Administrativo | F01-01, F01-02, F01-03 | US-001 a US-007 |
| **BR-A02** | Ativação e Gestão de Contas | F02-01, F02-02, F02-05 | US-008 a US-014, US-022, US-023 |
| **BR-A03** | Configuração de Planos Comerciais | F02-03 | US-015 a US-018 |
| **BR-A04** | Vinculação de Assinaturas | F02-04 | US-019 a US-021 |
| **BR-A05** | Gestão de Usuários e Permissões (RBAC) | F03-01 a F03-04 | US-024 a US-036 |
| **BR-B01** | Portal do Cliente com Autenticação | F04-01, F04-03 | US-037 a US-039, US-045, US-046 |
| **BR-B02** | Onboarding Guiado | F04-02 | US-040 a US-044 |
| **BR-B03** | App Switcher | F04-04 | US-047 a US-049 |
| **BR-B04** | Cadastro de Unidades de Negócio | F04-05 | US-050 a US-054 |
| **BR-B05** | Catálogo de Produtos/Serviços | F04-06 | US-055 a US-058 |

### 2.2 Resumo de Comportamento por Feature

| Feature | Comportamento Esperado |
|:---|:---|
| **F01-01** | Dashboard carrega em ≤3s. Indicadores: contas ativas, por status, por plano. Clicáveis — levam à lista filtrada |
| **F02-01** | Criar Tenant → status PENDING_ONBOARDING. Disparar e-mail com link único (expira em 7 dias). Reenvio disponível |
| **F02-02** | Transições de status: PENDING→ACTIVE, ACTIVE↔SUSPENDED, ACTIVE↔INACTIVE. Suspensão bloqueia acesso em ≤5min |
| **F02-03** | Planos com nome, preço, recorrências, módulos incluídos. Edição gera nova versão. Desativação não afeta assinantes |
| **F03-02** | 4 papéis com matriz de permissões (RN10-01). Admin vê tudo. Auditor só lê. Operador não edita |
| **F03-04** | Menu lateral e botões condicionais ao papel. Acesso direto a URL proibida → tela 403 amigável |
| **F04-02** | Onboarding em 4 passos. Primeira BU vira Matriz. Tenant só muda para ACTIVE após conclusão |
| **F04-05** | Unidades hierárquicas (Matriz/Filial). CNPJ único entre ativos. Soft delete libera CNPJ para reúso |

---

## 3. Regras de Negócio Implementadas

### 3.1 RNs com Especificação Formal

| RN | Descrição Formal | Feature | Casos de Borda |
|:---|:---|:---|:---|
| **RN01-01** | Métricas consideram apenas tenants com `deleted_dt IS NULL` | F01-01 | Tenant soft-deletado não aparece no dashboard |
| **RN04-01** | Criação de Tenant gera auditoria: admin, data/hora, dados iniciais | F02-01 | Criação sem admin autenticado → 401 |
| **RN05-01** | Transições de status permitidas: PENDING→ACTIVE, ACTIVE→SUSPENDED, SUSPENDED→ACTIVE, ACTIVE→INACTIVE, INACTIVE→ACTIVE | F02-02 | Tentar ACTIVE→PENDING → 422 |
| **RN05-02** | Suspensão exige motivo registrado | F02-02 | Motivo vazio → 400 |
| **RN06-01** | Plano com assinantes ativos não pode ser excluído | F02-03 | Tentar excluir → 422 |
| **RN06-02** | Alteração de preço não afeta assinaturas existentes (locked_price na subscription) | F02-03 | Upgrade mantém preço da data de contratação (atributo locked_price) (DT-009) |
| **RN06-03** | Deve existir pelo menos 1 plano ativo no sistema | F02-03 | Tentar desativar o último plano ativo → 422 "Não é possível desativar o último plano ativo". Validação: `SELECT COUNT(*) FROM plan WHERE status = 'ACTIVE' AND deleted_dt IS NULL` antes de desativar |
| **RN07-01** | Um tenant só pode ter 1 assinatura ativa por vez | F02-04 | Tentar criar segunda ativa → 409 |
| **RN07-02** | Upgrade/downgrade não pode deixar tenant sem assinatura ativa durante a transição | F02-04 | Transação atômica (`@Transactional`): finalizar assinatura anterior (`end_date = NOW()`) + criar nova (`start_date = NOW()`) no mesmo batch. Se qualquer operação falhar → rollback completo. Assinatura atual não pode estar SUSPENDED no momento do upgrade |
| **RN08-01** | Auditoria cobre 100% das ações administrativas com tenant_id e user_id corretos | F02-05 | Query sem registro de auditoria → falha de conformidade. Campo actor_id deve refletir user_id real (não UUID.randomUUID) (DT-002) |
| **RN08-02** | Registros de auditoria são imutáveis | F02-05 | Tentar UPDATE/ DELETE em audit_log → 403 |
| **RN09-03** | Admin do tenant não pode desativar a si mesmo | F03-01 | Tentar → 422 |
| **RN10-01** | Matriz de permissões: Admin total, Gerente edita sua BU, Operador só lê, Auditor só lê | F03-02 | Operador tentar PATCH /products → 403 |
| **RN12-01** | Dupla camada: ocultação UX + bloqueio de segurança no acesso direto | F03-04 | Acesso direto a URL sem permissão → 403 (não 404) |
| **RN14-01** | Onboarding obrigatório no primeiro acesso | F04-02 | Pular URL do passo → redireciona para passo correto |
| **RN14-02** | Primeira Unidade de Negócio cadastrada no onboarding é automaticamente definida como Matriz | F04-02 | Campo `parent_id = NULL` e `hierarchy_type = 'MATRIZ'` na primeira BU. Se o onboarding for retomado após interrupção, verificar se já existe BU before setting |
| **RN14-04** | Tenant só vira ACTIVE após conclusão do onboarding | F04-02 | API chamada antes de concluir → status ainda PENDING |
| **RN17-01** | CNPJ único entre Unidades ativas do mesmo tenant | F04-05 | Soft delete da BU antiga → CNPJ liberado para reúso |
| **RN17-02** | Unidade de Negócio desativada não pode ser definida como "pai" de novas filiais | F04-05 | Validação ao criar/editar BU: `if (parent != null && parent.status == INACTIVE) throw InvalidParentException("Unidade pai está desativada").` Aplicar também ao editar, se `parent_id` for alterado |
| **RN18-04** | Exclusão de produtos segue soft delete | F04-06 | Produto "excluído" não aparece em queries mas está no audit_log |

### 3.2 RNs NÃO Implementadas por Esta Solução

| RN | Justificativa |
|:---|:---|
| RNs da camada fiscal (NCM, IBS, CBS) | Fora do escopo — pertencem ao Tributali-Engine |
| RNs de pedidos e faturas (Quote, Order, Invoice) | Fora do escopo — pertencem ao Storekeeper |
| RNs de split de pagamento (SplitPayment) | Fora do escopo — pertencem ao Tributali-Engine |

---

### 3.3 Matriz de Cobertura de Todas as RNs

> **Origem:** [PRD.md §6.6](./PRD.md) — 18 famílias de RNs (51 regras individuais). Esta matriz mapeia CADA RN do PRD para a seção do SPECS que a cobre.

| RN (PRD §6.6) | Descrição Resumida | Coberta em |
|:---|:---|:---|
| **RN01-01** | Métricas excluem soft-deleted | §3.1 (formal) |
| **RN01-02** | Período padrão: mês atual | §4.2 (`GET /dashboard/admin/evolution`) |
| **RN01-03** | Zero exibido como "0" | §7 (F01-01) |
| **RN02-01** | Soft-deleted não aparece na lista | §6.1 (campos de auditoria) + ARCHITECTURE §5.1 (BaseRepository) |
| **RN02-02** | Busca case-insensitive | §7 (F01-02) |
| **RN03-01** | Alerta onboarding >48h | §7 (F01-03) |
| **RN03-02** | Alertas sem personalização | §7 (F01-03) |
| **RN04-01** | Criação Tenant → auditoria | §3.1 (formal) |
| **RN04-02** | Razão social obrigatória, única | §4.2 (`POST /tenants`) |
| **RN04-03** | Link ativação único, uso único | §7 (F02-01) |
| **RN05-01** | Transições de status permitidas | §3.1 (formal) |
| **RN05-02** | Suspensão exige motivo | §3.1 (formal) |
| **RN05-03** | Reativação restaura permissões | §7 (F02-02) |
| **RN06-01** | Plano com assinantes não pode ser excluído | §3.1 (formal) |
| **RN06-02** | Alteração preço não afeta existentes (locked_price) | §3.1 (formal) ← adicionado v1.7 |
| **RN06-03** | Mínimo 1 plano ativo | §3.1 (formal) ← adicionado v1.2 |
| **RN07-01** | 1 assinatura ativa por tenant | §3.1 (formal) |
| **RN07-02** | Transição atômica no upgrade | §3.1 (formal) ← adicionado v1.2 |
| **RN07-03** | Data término opcional (contínua) | §4.2 (`POST /subscriptions/{id}/change-plan`) + §6.1 (Subscription.end_date nullable) |
| **RN08-01** | Auditoria 100% ações admin com tenant_id e user_id corretos | §3.1 (formal) ← adicionado v1.7 |
| **RN08-02** | Auditoria imutável | §3.1 (formal) + §6.1 (AuditEntry immutable) |
| **RN09-01** | Convite expira em 7 dias | §7 (F03-01) |
| **RN09-02** | E-mail único por tenant | §4.2 (`POST /users`) |
| **RN09-03** | Admin não desativa a si mesmo | §3.1 (formal) |
| **RN10-01** | Matriz de permissões (4 papéis) | §3.1 (formal) + ARCHITECTURE §4.1 (@RequiresPermission) |
| **RN11-01** | Usuário requer ≥1 BU | §4.2 (`PUT /users/{uid}/permissions`) |
| **RN11-02** | Usuário requer ≥1 Módulo | §4.2 (`PUT /users/{uid}/permissions`) |
| **RN11-03** | Módulos = plano contratado | §7 (F03-03) |
| **RN12-01** | Dupla camada UX + segurança | §3.1 (formal) |
| **RN12-02** | Nome módulo ativo no topo | Frontend (não aplicável ao backend) |
| **RN13-01** | Senha mín. 8, letra+número | Delegado ao Keycloak (password policy) |
| **RN13-02** | Sessão expira 60min inatividade | Delegado ao Keycloak (accessTokenLifespan) |
| **RN13-03** | Link reset senha uso único | Delegado ao Keycloak (required actions) |
| **RN14-01** | Onboarding obrigatório | §3.1 (formal) |
| **RN14-02** | Primeira BU = Matriz | §3.1 (formal) ← adicionado v1.2 |
| **RN14-03** | Onboarding concluído = todos passos | §4.2 (`POST /onboarding/complete`) + §7 (F04-02) |
| **RN14-04** | Tenant → ACTIVE após onboarding | §3.1 (formal) |
| **RN15-01** | Dashboard adapta-se ao módulo | §7 (F04-03) |
| **RN15-02** | Fase 0: dashboard genérico | §7 (F04-03) |
| **RN16-01** | App Switcher: interseção plano×perms | Frontend + §2.2 (JWT `modules[]`) |
| **RN16-02** | Placeholder "FBSO Platform" | Frontend (não aplicável ao backend) |
| **RN16-03** | Troca módulo mantém contexto BU | Frontend (não aplicável ao backend) |
| **RN17-01** | CNPJ único entre ativos | §3.1 (formal) |
| **RN17-02** | BU desativada não pode ser pai | §3.1 (formal) ← adicionado v1.2 |
| **RN17-03** | Primeira BU = Matriz | §3.1 (RN14-02 — mesma lógica) |
| **RN17-04** | Sem limite níveis hierárquicos | §6.1 (parent_id auto-referenciado) |
| **RN17-05** | Seletor BU reflete permissões | §4.1 (`GET /business-units` RBAC) + ARCHITECTURE §3 (BU Filter) |
| **RN18-01** | Catálogo segmentado por BU | §6.1 (ProductService.business_unit_id FK) |
| **RN18-02** | SKU opcional, único por BU | §4.2 (`POST /products`) + §6.1 (índice parcial) |
| **RN18-03** | Indicador "Não mapeado" | §7 (F04-06) |
| **RN18-04** | Soft delete em produtos | §3.1 (formal) |

> **Legenda:** §3.1 = especificação formal com casos de borda | §4.2 = regra de validação por endpoint | §6.1 = modelo de dados/constraints | §7 = critério de aceitação | ARCHITECTURE = coberto pelo design técnico | Keycloak/Frontend = delegado a outro componente

**Cobertura total: 51/51 RNs mapeadas (100%).** 20 RNs em §3.1 (formato completo), 31 RNs cobertas via §4.2, §6.1, §7, ARCHITECTURE ou delegadas.

---

### 4.1 Endpoints REST

| Método | Path | Descrição | RBAC | Request Schema | Response Schema | Status HTTP | Erros HTTP |
|:---|:---|:---|:---|:---|:---|:---|:---|
| `GET` | `/api/v1/dashboard/admin/summary` | Indicadores do SaaS | Admin FBSO | — | `DashboardSummaryResponse` | 200 | 401, 403, 500 |
| `GET` | `/api/v1/dashboard/admin/evolution` | Evolução da base | Admin FBSO | `?period,start_date,end_date` | `DashboardSummaryResponse` | 200 | 400, 401, 403, 500 |
| `GET` | `/api/v1/tenants` | Listar tenants (paginado) | Admin FBSO | — | `TenantResponse[]` | 200 | 401, 403, 500 |
| `POST` | `/api/v1/tenants` | Criar tenant | Admin FBSO | `TenantCreateRequest` | `TenantResponse` | 201 | 400, 401, 403, 409, 422, 500 |
| `GET` | `/api/v1/tenants/{id}` | Detalhes do tenant | Admin FBSO | — | `TenantResponse` | 200 | 401, 403, 404, 500 |
| `PATCH` | `/api/v1/tenants/{id}` | Atualizar tenant | Admin FBSO | `TenantUpdateRequest` | `TenantResponse` | 200 | 400, 401, 403, 404, 422, 500 |
| `POST` | `/api/v1/tenants/{id}/suspend` | Suspender tenant | Admin FBSO | `SuspendRequest` | `TenantResponse` | 200 | 400, 401, 403, 404, 422, 500 |
| `POST` | `/api/v1/tenants/{id}/reactivate` | Reativar tenant | Admin FBSO | — | `TenantResponse` | 200 | 401, 403, 404, 422, 500 |
| `POST` | `/api/v1/tenants/{id}/resend-invite` | Reenviar convite | Admin FBSO | — | `TenantResponse` | 200 | 401, 403, 404, 422, 500 |
| `GET` | `/api/v1/plans` | Listar planos | Admin FBSO | — | `PlanResponse[]` | 200 | 401, 403, 500 |
| `POST` | `/api/v1/plans` | Criar plano | Admin FBSO | `PlanCreateRequest` | `PlanResponse` | 201 | 400, 401, 403, 422, 500 |
| `PATCH` | `/api/v1/plans/{id}` | Atualizar plano | Admin FBSO | `PlanUpdateRequest` | `PlanResponse` | 200 | 400, 401, 403, 404, 422, 500 |
| `POST` | `/api/v1/plans/{id}/deactivate` | Desativar plano | Admin FBSO | — | `PlanResponse` | 200 | 401, 403, 404, 422, 500 |
| `GET` | `/api/v1/tenants/{tid}/subscriptions` | Histórico de assinaturas | Admin FBSO | — | `SubscriptionResponse[]` | 200 | 401, 403, 404, 500 |
| `POST` | `/api/v1/tenants/{tid}/subscriptions` | Criar assinatura | Admin FBSO | `SubscriptionCreateRequest` | `SubscriptionResponse` | 201 | 400, 401, 403, 404, 409, 422, 500 |
| `POST` | `/api/v1/subscriptions/{id}/change-plan` | Upgrade/downgrade | Admin FBSO | `ChangePlanRequest` | `SubscriptionResponse` | 200 | 400, 401, 403, 404, 422, 500 |
| `POST` | `/api/v1/subscriptions/{id}/suspend` | Suspender assinatura | Admin FBSO | `SuspendSubscriptionRequest` | `SubscriptionResponse` | 200 | 400, 401, 403, 404, 422, 500 |
| `GET` | `/api/v1/users` | Listar usuários do tenant | Admin Tenant | — | `UserResponse[]` | 200 | 401, 403, 500 |
| `POST` | `/api/v1/users` | Convidar usuário | Admin Tenant | `UserInviteRequest` | `UserResponse` | 201 | 400, 401, 403, 409, 422, 500 |
| `PATCH` | `/api/v1/users/{id}` | Atualizar usuário | Admin Tenant | `UserUpdateRequest` | `UserResponse` | 200 | 400, 401, 403, 404, 422, 500 |
| `POST` | `/api/v1/users/{id}/deactivate` | Desativar usuário | Admin Tenant | — | `UserResponse` | 200 | 401, 403, 404, 422, 500 |
| `GET` | `/api/v1/users/{uid}/permissions` | Permissões do usuário | Admin Tenant | — | `PermissionResponse` | 200 | 401, 403, 404, 500 |
| `PUT` | `/api/v1/users/{uid}/permissions` | Atualizar vínculos | Admin Tenant | `PermissionUpdateRequest` | `PermissionResponse` | 200 | 400, 401, 403, 404, 422, 500 |
| `GET` | `/api/v1/business-units` | Listar unidades (hierárquico) | Admin/Manager/Operator | — | `BusinessUnitResponse[]` | 200 | 401, 403, 500 |
| `POST` | `/api/v1/business-units` | Criar unidade | Admin/Manager | `BusinessUnitCreateRequest` | `BusinessUnitResponse` | 201 | 400, 401, 403, 409, 422, 500 |
| `PATCH` | `/api/v1/business-units/{id}` | Atualizar unidade | Admin/Manager | `BusinessUnitUpdateRequest` | `BusinessUnitResponse` | 200 | 400, 401, 403, 404, 422, 500 |
| `POST` | `/api/v1/business-units/{id}/deactivate` | Desativar unidade | Admin Tenant | — | `BusinessUnitResponse` | 200 | 401, 403, 404, 422, 500 |
| `GET` | `/api/v1/products` | Listar produtos (filtrado BU) | Todos com acesso à BU | — | `ProductResponse[]` | 200 | 401, 403, 500 |
| `POST` | `/api/v1/products` | Criar produto | Admin/Manager | `ProductCreateRequest` | `ProductResponse` | 201 | 400, 401, 403, 422, 500 |
| `PATCH` | `/api/v1/products/{id}` | Atualizar produto | Admin/Manager | `ProductUpdateRequest` | `ProductResponse` | 200 | 400, 401, 403, 404, 422, 500 |
| `POST` | `/api/v1/products/{id}/deactivate` | Desativar produto | Admin/Manager | — | `ProductResponse` | 200 | 401, 403, 404, 422, 500 |
| `GET` | `/api/v1/onboarding/status` | Status do onboarding | Cliente autenticado | — | `OnboardingStatusResponse` | 200 | 401, 403, 500 |
| `PATCH` | `/api/v1/onboarding/step-1` | Confirmar dados | Cliente autenticado | `OnboardingStep1Request` | `OnboardingStatusResponse` | 200 | 400, 401, 403, 422, 500 |
| `POST` | `/api/v1/onboarding/step-2` | Cadastrar Matriz | Cliente autenticado | `OnboardingStep2Request` | `OnboardingStatusResponse` | 201 | 400, 401, 403, 422, 500 |
| `POST` | `/api/v1/onboarding/complete` | Finalizar onboarding | Cliente autenticado | — | `OnboardingStatusResponse` | 200 | 400, 401, 403, 422, 500 |
| `GET` | `/api/v1/dashboard/client/summary` | Dashboard do cliente | Cliente autenticado | `?module_id` | `DashboardClientResponse` | 200 | 401, 403, 500 |
| `GET` | `/api/v1/audit` | Consultar auditoria | Admin FBSO / Auditor | `?start_date,end_date,action,entity_type,page,size` | `AuditEntryResponse[]` | 200 | 400, 401, 403, 500 |

### 4.2 Regras de Validação por Endpoint

| Endpoint | Validações |
|:---|:---|
| `POST /tenants` | `name_corporate` obrigatório, único entre ativos. `segment` obrigatório |
| `POST /tenants/{id}/suspend` | `reason` obrigatório (não vazio). Status atual deve ser ACTIVE |
| `POST /plans` | `name` obrigatório. `price` > 0. Pelo menos 1 recorrência selecionada |
| `POST /users` | `email` obrigatório, único por tenant (ativos). `role` obrigatório, valor do enum Role. `name` obrigatório |
| `POST /business-units` | `cnpj` obrigatório, formato válido, único entre ativos do tenant. `tax_regime` obrigatório |
| `POST /products` | `name` obrigatório. `type` obrigatório (PRODUCT ou SERVICE). `sku` se informado, único por BU |
| `POST /onboarding/step-2` | `cnpj` obrigatório, válido. `tax_regime` obrigatório |
| `POST /onboarding/complete` | Todos os passos obrigatórios concluídos. Tenant em PENDING_ONBOARDING |
| `GET /dashboard/admin/evolution` | `period` opcional (7d, 30d, 90d, mes_atual, ano_atual). Se inválido → assume mês atual. `start_date` e `end_date` opcionais, mutuamente exclusivos com `period` |
| `POST /subscriptions/{id}/change-plan` | `plan_id` obrigatório, deve referir plano ativo. Assinatura atual não pode estar SUSPENDED. `effective_date` opcional (padrão: imediato). Transição não pode deixar tenant sem assinatura ativa |
| `GET /audit` | `start_date` e `end_date` opcionais (formato ISO 8601). `action` opcional (enum: CREATED, UPDATED, SUSPENDED, REACTIVATED, PLAN_CHANGED). `entity_type` opcional. Paginação obrigatória (padrão: 25 itens, máx: 100). `sort` opcional (timestamp DESC padrão) |
| `PATCH /business-units/{id}` | `cnpj` não pode ser alterado após cadastro. `corporate_name` opcional. `tax_regime` se alterado, deve ser valor do enum TaxRegime. `parent_id` se alterado, deve referir BU ativa do mesmo tenant |

---

## 5. Requisitos Não-Funcionais

| NFR | Descrição | Métrica | Como Verificar |
|:---|:---|:---|:---|
| **BR-NFR01** | Disponibilidade 99,5% (6h-23h) | Uptime ≥ 99,5% | Health check K8s + monitoramento |
| **BR-NFR02** | Isolamento total entre tenants e BUs | Zero incidentes de vazamento | PostgreSQL RLS (Camada 1) + BaseRepository (Camada 2) + Teste automatizado cross-tenant (Camada 3) |
| **BR-NFR03** | Auditoria em 100% das ações admin | Cobertura = 100% | Auditoria amostral mensal |
| **BR-NFR04** | Features autodescritivas com documentação inline (OpenAPI + README.md por pacote) | 100% dos endpoints com schemas no OpenAPI YAML; 1 README.md por módulo | Auditoria de documentação por sprint |
| **BR-NFR05** | Telas principais ≤ 3s | p95 latency ≤ 3s | Teste de carga no dashboard |
| **BR-NFR06** | Core suporta acoplamento de módulos | Novo módulo ativado em ≤ 1 sprint | Medição na Fase 1 |
| **BR-NFR07** | 100% dos erros de API seguem RFC 7807 (Problem Details) com mensagens em PT-BR, sem expor stack traces | Resposta de erro contém type, title, status, detail (RFC 7807); zero stack traces em respostas HTTP | Teste automatizado de cada endpoint para conformidade RFC 7807 |
| **BR-NFR08** | Interface em PT-BR | 100% das mensagens em português | Revisão de strings |
| **NFR-OWASP** | Zero vulnerabilidades críticas em scan SAST/DAST. Proteção contra OWASP Top 10: injection (SQL), broken auth (JWT), XSS (input validation), CSRF | 0 críticos/high em scan SAST; 0 falhas de injection em teste DAST | Relatório de scan SAST/DAST a cada sprint; teste de penetração a cada release |
| **NFR-LGPD** | Dados pessoais (nome, email, CPF/CNPJ) mascarados em logs. Soft delete preserva dados para auditoria (LGPD Art. 16) | Zero dados pessoais em texto claro em logs de produção; 100% das entidades com soft delete implementado | Inspeção de logs pós-release; revisão de schema (colunas `deleted_dt`) |

> **Mapeamento NFR → ADR:** ADR-L07 (PostgreSQL RLS) atende BR-NFR02 (isolamento). ADR-L02 (AOP) atende BR-NFR03 (auditoria). ADR-L03 (auditoria assíncrona) atende BR-NFR03. ADR-L04 (RFC 7807) atende BR-NFR07. ADR-L05 (índices parciais) atende NFR-LGPD (soft delete). **ADRs pendentes:** Criar ADR para BR-NFR01 (Disponibilidade 99,5% — health checks, readiness/liveness probes) e BR-NFR05 (Performance p95 ≤3s — HikariCP tuning, caching, índices).

---

## 6. Modelo de Dados

### 6.1 Entidades (Fase 0)

| Entidade | Tabela | Campos Essenciais | Constraints |
|:---|:---|:---|:---|
| **Tenant** | `tenant` | id, name_corporate, name_fantasy, segment, status | name_corporate UNIQUE (ativos) |
| **Plan** | `plan` | id, name, description, price, status, version | price > 0 |
| **PlanModule** | `plan_module` | id, plan_id (FK), module_name | (plan_id, module_name) UNIQUE |
| **Subscription** | `subscription` | id, tenant_id (FK), plan_id (FK), start_date, end_date, status | 1 ativa por tenant |
| **User** | `user` | id, tenant_id (FK), external_keycloak_id, email, name, status | email UNIQUE por tenant (ativos) |
| **UserPermission** | `user_permission` | id, user_id (FK), business_unit_id (FK), role | (user_id, business_unit_id) UNIQUE |
| **ResourceAction** | `resource_action` | id, resource_name, action | (resource_name, action) UNIQUE |
| **RoleResource** | `role_resource` | id, role, resource_action_id (FK) | (role, resource_action_id) UNIQUE |
| **BusinessUnit** | `business_unit` | id, tenant_id (FK), parent_id (FK), cnpj, corporate_name, tax_regime, address, status | cnpj UNIQUE por tenant (ativos) |
| **ProductService** | `product_service` | id, business_unit_id (FK), name, sku, type, description, status | sku UNIQUE por BU (ativos, não-nulo) |
| **AuditEntry** | `audit_log` | id, timestamp, tenant_id, action, entity_type, entity_id, actor_id, previous_value, new_value | Immutable — sem UPDATE ou DELETE |

### 6.2 Entidades Fora do Escopo (Referência Futura)

| Entidade | Fase | Tabela |
|:---|:---|:---|
| Billable | Tributali-Engine | `billable` |
| ProductBillableMapping | Tributali-Engine | `product_billable_mapping` |
| Order, OrderItem | Storekeeper | `order`, `order_item` |
| Invoice, InvoiceItem | Storekeeper | `invoice`, `invoice_item` |
| TransactionPayment | Storekeeper | `transaction_payment` |
| SplitPayment | Tributali-Engine | `split_payment` |
| BillingInfo | Storekeeper | `billing_info` |
| BankAccount | Storekeeper | `bank_account` |

### 6.3 Campos de Auditoria (TODAS as tabelas)

| Campo | Tipo | Descrição |
|:---|:---|:---|
| `created_dt` | TIMESTAMPTZ | Data/hora de criação (NOT NULL) |
| `updated_dt` | TIMESTAMPTZ | Data/hora da última atualização (NOT NULL) |
| `created_by` | UUID → user.id | Usuário criador |
| `updated_by` | UUID → user.id | Usuário da última atualização |
| `deleted_dt` | TIMESTAMPTZ | Soft delete (NULL = ativo) |
| `deleted_by` | UUID → user.id | Usuário que excluiu |

### 6.4 Tabela de Auditoria

```sql
CREATE TABLE audit_log (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    timestamp       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    tenant_id       UUID NOT NULL,
    action          VARCHAR(50) NOT NULL,
    entity_type     VARCHAR(50) NOT NULL,
    entity_id       UUID NOT NULL,
    actor_id        UUID NOT NULL,
    actor_name      VARCHAR(255),
    previous_value  JSONB,
    new_value       JSONB,
    reason          VARCHAR(500)
);
```

---

## 7. Critérios de Aceitação por Feature

| Feature | Critério DONE | Evidência | Nível DoD |
|:---|:---|:---|:---|
| **F01-01** | Dashboard carrega em ≤3s. Indicadores clicáveis levam à lista filtrada. Período padrão: mês atual | Print do dashboard | F |
| **F01-02** | Lista de contas exibe razão social, plano, status, data de criação. Paginação a cada 25 registros. Busca textual filtra em tempo real a partir de 3 caracteres | Print da lista com busca aplicada | F |
| **F01-03** | Alertas de onboarding incompleto (>48h) e assinatura suspensa aparecem como cards coloridos no topo do dashboard. Cards são clicáveis e levam à lista filtrada | Print dos cards de alerta | F |
| **F02-01** | Tenant criado → status PENDING. E-mail enviado. Link expira em 7 dias. Reenvio funcional | E-mail recebido + log | F |
| **F02-02** | Transições de status respeitam RN05-01. Suspensão bloqueia acesso em ≤5min. Timeline de status funcional | Teste de cada transição | F |
| **F02-03** | Plano criado disponível para assinatura. Edição gera nova versão. Desativação preserva assinantes | Lista de planos reflete versão atualizada após edição; assinantes existentes não são afetados por alteração de preço | F |
| **F02-04** | Assinatura vinculada a plano ativo. Upgrade/downgrade finaliza assinatura anterior e cria nova. Apenas 1 assinatura ativa por tenant. Suspensão bloqueia módulos em ≤5min. Change-plan preserva preço contratado (locked_price) (TC-F02-04-010, DT-009) | Histórico de assinaturas do tenant com timeline | F |
| **F02-05** | Auditoria registra 100% das ações admin (criação, edição, mudança de status, alteração de permissões). Registros imutáveis (sem UPDATE/DELETE). Filtros por período e tipo de ação funcionam. Auditoria @Async com tenant_id e user_id corretos (TC-F02-05-009, DT-002) | Log de auditoria com linha do tempo + tentativa de exclusão rejeitada (403) | F |
| **F03-01** | Usuário convidado recebe e-mail com link. E-mail único por tenant. Admin não pode desativar a si mesmo. Lista de usuários exibe nome, e-mail, papel, status, unidades vinculadas | E-mail de convite + teste de autodesativação rejeitada (422) | F |
| **F03-02** | Matriz RN10-01 aplicada. Admin vê tudo. Operador não edita. Auditor só lê | Teste de cada papel × endpoint | F |
| **F03-03** | Vinculação usuário × unidade × módulo configurável. Admin tem acesso implícito a todas as unidades. Usuário sem unidade ou módulo não acessa o portal. Efeito imediato na próxima ação | Teste de acesso com e sem vinculação | F |
| **F03-04** | Menu condicional ao papel. Acesso direto → 403 (amigável, sem detalhes técnicos) | Tentativa de URL proibida | F |
| **F04-01** | Login com e-mail/senha funcional. Recuperação de senha envia link (expira 1h). Bloqueio temporário após 5 tentativas incorretas (15min) | Fluxo completo de login + recuperação + bloqueio | F |
| **F04-02** | Onboarding 4 passos obrigatórios. Primeira BU = Matriz. Tenant → ACTIVE ao concluir. Barra de progresso visível. Não é possível pular etapas | Log de auditoria mostra transição PENDING_ONBOARDING→ACTIVE após conclusão do onboarding; TenantContext contém status ACTIVE | F |
| **F04-03** | Dashboard do cliente exibe cards: Unidades Ativas, Produtos no Catálogo, Plano Contratado. Notificações e lembretes visíveis com link para ação | Print do dashboard do cliente | F |
| **F04-04** | App Switcher exibe módulos disponíveis no plano e autorizados ao usuário. Troca de módulo atualiza menu lateral e conteúdo. Visível mesmo com 1 módulo (exibe nome do módulo ativo) | Print do App Switcher com módulos listados | F |
| **F04-05** | Unidades hierárquicas. CNPJ único entre ativos. Soft delete libera reúso | Cadastro pós-soft-delete | F |
| **F04-06** | Produto cadastrado vinculado à BU ativa. SKU único por BU. Indicador "Não mapeado" | Lista de produtos | F |

---

## 8. Dependências e Integrações

### 8.1 Dependências de Outras Soluções

| Dependência | Tipo | Contrato |
|:---|:---|:---|
| **web_app-fbso-platform-portal** (Frontend) | Consome esta API | [API-CONTRACTS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/API-CONTRACTS.md) |
| **Keycloak** (IdP) | Provedor de autenticação | JWT RS256 — realms, clients, roles |
| **PostgreSQL** | Banco de dados | JDBC — schema `fbso_platform` |
| **RabbitMQ** (Futuro) | Mensageria para billing/split | Fora do escopo da Fase 0 |

> **Integrações detalhadas:** Consulte [INTEGRATION-MAP.md](./INTEGRATION-MAP.md) para o diagrama completo de integrações, incluindo fluxo OIDC, sequência de chamadas e contratos de integração.
>
> **Fluxo OIDC (Keycloak):** O frontend obtém o token JWT via Authorization Code Flow (I-02) e o envia como `Authorization: Bearer <jwt>` para o backend (I-03). O backend valida a assinatura RS256 com a chave pública do Keycloak (JWKS endpoint), extrai as claims (`tenant_id`, `user_id`, `roles`, `business_unit_ids`, `modules`) e estabelece o `TenantContext` para isolamento multi-tenant.

### 8.2 Contratos Fornecidos

| Contrato | Localização | Consumidor |
|:---|:---|:---|
| `fbso-platform-api.yaml` (OpenAPI 3.0) | `.specs/api/` (a criar) | Frontend |

---

## 9. Restrições e Premissas Técnicas

### 9.1 Restrições

| Restrição | Origem | Impacto |
|:---|:---|:---|
| Time técnico reduzido | Project Charter C1 | Velocidade de entrega limitada |
| Orçamento limitado | Project Charter C2 | Sem contratações significativas |
| Prazo: 14 semanas (24/07→30/10) | Project Charter §7 | Escopo fixo, time fixo → qualidade negociável em Should Haves |
| Nenhum módulo-produto antes do Core | Project Charter C3 | Não implementar entidades de fases futuras |

### 9.2 Premissas

| Premissa | Se inválida |
|:---|:---|
| Keycloak disponível e configurado (realm `fbso-platform`) | Backend não autentica — bloqueante |
| PostgreSQL 17 com schema `fbso_platform` provisionado | Backend não persiste — bloqueante |
| Frontend consome API via JWT Bearer Token | Integração quebrada se frontend usar outro auth |
| E-mails enviados via SMTP (serviço externo) | Convites e ativações não chegam — impacto médio |

---

## 10. Glossário da Solução

| Termo de Negócio | Termo Técnico | Definição |
|:---|:---|:---|
| Tenant | `tenant` (tabela), `TenantContext` (ThreadLocal) | Conta corporativa do cliente |
| Unidade de Negócio | `business_unit` (tabela) | CNPJ/filial vinculada a um Tenant |
| Plano | `plan` (tabela) | Pacote comercial com módulos e preço |
| Assinatura | `subscription` (tabela) | Vínculo Tenant × Plano com vigência |
| RBAC | `@RequiresPermission` + `PermissionService` + `RbacAspect` (DB-backed via `resource_action` + `role_resource`) | Controle de acesso baseado em papéis — matriz RN10-01 carregada do banco |
| Soft Delete | `deleted_dt IS NULL` + índices parciais | Exclusão lógica — dados nunca removidos |
| Tenant Isolation | PostgreSQL RLS com FORCE + `BaseRepository` + Teste cross-tenant (3 camadas) | Isolamento de dados entre clientes — defesa em profundidade |
| RLS (Row-Level Security) | PostgreSQL — `FORCE ROW LEVEL SECURITY` + `CREATE POLICY tenant_isolation USING (tenant_id = current_setting('app.current_tenant_id')::UUID)` em 4 tabelas | Camada 1 do isolamento multi-tenant — FORCE garante que nem o table owner escapa |
| Migration V003 | Flyway — ativa RLS + FORCE + cria políticas em 4 tabelas (subscription, user, business_unit, audit_log) | Atualizada na Sprint 4 Frente 0 com FORCE ROW LEVEL SECURITY |
| Migration V004 | Flyway — seed data: `INSERT INTO resource_action` (32 ações) + `INSERT INTO role_resource` (matriz RN10-01) | Criada na Sprint 4 Frente 0 — matriz RBAC populada |
| Migration V006 | Flyway — FK `user_permission.business_unit_id → business_unit.id` | Criada na Sprint 4 Frente 0 — integridade referencial |
| App Switcher | Lógica no frontend; backend provê `modules[]` no JWT | Seletor de módulos do portal |
| Onboarding | `/onboarding/*` endpoints | Fluxo guiado de primeiro acesso |
| JWT | JSON Web Token — assinado pelo Keycloak (RS256) | Token de autenticação stateless |

---

## 11. Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 2.4 | 17/07/2026 | Sprint 5 Frente 0 concluida: ambiente dev dockerizado (Keycloak 26 + PG 17 + MailHog). Stack atualizado (Flyway 12.11.0, PG driver 42.7.11). OAuth2 Client configurado para Authorization Code Flow. | Agente IA |
| 2.3 | 17/07/2026 | Sprint 5 planejada: auditoria 9-skill identificou 42 débitos. 24 débitos + 12 features = 36 tarefas em 4 frentes. Referência ao [IDENTIFIED-TECHNICAL-DEBT-sprint-05-portal-cliente](./sprints/sprint-05-portal-cliente/IDENTIFIED-TECHNICAL-DEBT-sprint-05-portal-cliente.md). | Agente IA |
| 2.1 | 17/07/2026 | **Sprint 4 Frente 0 concluída:** Stack atualizado (Caffeine Cache). Glossário (§10): RBAC DB-backed, RLS com FORCE, novas migrations V004+V006. Status atualizado. Linha de status duplicada removida | Agente IA |
| 1.7 | 16/07/2026 | v1.7 — RN06-02 e RN08-01 reforçadas, 2 novos cenários de teste (TC-F02-04-010, TC-F02-05-009), referência a débitos técnicos (IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md) | Time Técnico |
| 1.6 | 16/07/2026 | Sprint 3 iniciada (16/07/2026). Status atualizado para "Em Desenvolvimento". | Time Técnico |
| 1.5 | 15/07/2026 | Revisão Caveman (DOCS-SERVICE-CAVEMAN-REVIEW.md): Corrigida contagem de entidades (10→11, §1.2). Adicionado endpoint `GET /dashboard/client/summary` (§4.1, 36→37). Corrigida contagem RN (45→51, §3.3). Corrigido header da matriz RN (~45→51). | Caveman/IA |
| 1.3 | 14/07/2026 | Adicionado PostgreSQL Row-Level Security (RLS) como camada 1 de defesa em profundidade (§5 BR-NFR02). Atualizado pipeline de segurança (§1.2) incluindo RLS. Glossário (§10) com novas entradas: RLS, Migration V003, Tenant Isolation atualizado para 3 camadas. | Agente Arquiteto/IA |
| 1.2 | 14/07/2026 | Correção pós-gate: 2 não-conformidades resolvidas do TECHNICAL_SPECS_FAIL_REPORT.md v1.0. 0 bloqueantes, 2 não-bloqueantes (NC-001, NC-002). Alterações: §3.1 renomeado para "RNs com Especificação Formal"; adicionadas 4 RNs de alto impacto (RN06-03, RN07-02, RN14-02, RN17-02) com casos de borda; nova §3.3 "Matriz de Cobertura de Todas as RNs" mapeando 45/45 RNs do PRD §6.6 (100% de cobertura). | Agente Corretor SPECS/IA |
| 1.1 | 14/07/2026 | Correção pós-gate: 12 não-conformidades resolvidas do TECHNICAL_SPECS_FAIL_REPORT.md v1.0. 4 bloqueantes (NC-003, NC-009, NC-010, NC-011), 8 não-bloqueantes (NC-001, NC-002, NC-004, NC-005, NC-006, NC-007, NC-008, NC-012). Alterações: schemas e erros HTTP em §4.1; critérios de aceitação para 18 features em §7 com coluna Nível DoD; NFRs OWASP e LGPD em §5; mapeamento BR→Feature corrigido em §2.1; validades expandidas em §4.2; AuditEntry em §6.1; métricas NFR refinadas; referências cruzadas adicionadas | Agente Corretor SPECS/IA |
| 1.0 | 13/07/2026 | Criação inicial: 11 seções. Cobertura completa dos 4 épicos, 18 features, 58 user stories, 10 BRs, 8 NFRs. Rastreabilidade BR→Feature→US. 37 endpoints REST especificados | Time Técnico |

---

## 12. Referências

| Documento | Localização |
|:---|:---|
| Débitos Técnicos Sprint 3 | [IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md](./IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md) |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: create-specification, spec-miner, domain-modeling, acceptance-criteria, documentation-writer. v1.7 em 16/07/2026: Sprint 3 — débitos técnicos incorporados.*
