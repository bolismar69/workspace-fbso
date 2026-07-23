# API-CONTRACTS.md — Contratos de API da FBSO Platform

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Programa:** FBSO Platform — Portal Administrativo SaaS (Fase 0 — Core)
- **Versão:** 1.0
- **Data:** 13 de Julho de 2026
- **Status:** Esboço Inicial — a ser detalhado com OpenAPI YAML completo
- **Origem:** [TECHNICAL-PLAN.md](./TECHNICAL-PLAN.md) e [ARCHITECTURE.md](./ARCHITECTURE.md)

---

## 1. Objetivo

Este documento define os contratos de API entre o **Frontend** (React/Next.js) e o **Backend** (Java 25/Spring Boot) da FBSO Platform. Ele é a "fonte da verdade" que ambos os times consomem durante o desenvolvimento paralelo (Cenário C do TECHNICAL-PLAN.md).

> ⚠️ **Status:** Esboço de alto nível. O OpenAPI YAML detalhado (`fbso-platform-api.yaml`) será gerado a partir deste documento na Fase 0 — Fundação (pré-Kickoff).

---

## 2. Convenções da API

### 2.1 Padrões Gerais

| Convenção | Valor |
|:---|:---|
| **Base URL** | `/api/v1` (confirmado — consistente com INTEGRATION-MAP.md) |
| **Formato** | JSON (request/response) |
| **Encoding** | UTF-8 |
| **Autenticação** | JWT Bearer Token (`Authorization: Bearer <token>`) |
| **Content-Type** | `application/json` |
| **Versionamento** | Prefixo de URL (`/api/v1`, `/api/v2`...) |

### 2.2 Padrões de Nomenclatura

| Recurso | Convenção | Exemplo |
|:---|:---|:---|
| **Coleções** | Plural, lowercase, hífen | `/tenants`, `/business-units` |
| **Recurso individual** | `/{recurso}/{id}` | `/tenants/t-12345` |
| **Sub-recursos** | `/{recurso}/{id}/{sub-recurso}` | `/tenants/t-12345/users` |
| **Ações customizadas** | Verbo no path | `/tenants/t-12345/suspend` |
| **Query params** | snake_case | `?start_date=2026-01-01&status=ACTIVE` |

### 2.3 Códigos de Status HTTP

| Código | Quando usar |
|:---|:---|
| **200 OK** | Requisição bem-sucedida (GET, PATCH) |
| **201 Created** | Recurso criado com sucesso (POST) |
| **204 No Content** | Operação bem-sucedida sem corpo de resposta (DELETE, ações) |
| **400 Bad Request** | Dados de entrada inválidos (validação de negócio) |
| **401 Unauthorized** | Token JWT ausente, expirado ou inválido |
| **403 Forbidden** | Token válido mas sem permissão (RBAC) para o recurso/ação |
| **404 Not Found** | Recurso não encontrado (ou fora do tenant do usuário) |
| **409 Conflict** | Conflito de estado (ex: CNPJ duplicado, transição de status inválida) |
| **422 Unprocessable Entity** | Dados sintaticamente corretos mas semanticamente inválidos |
| **500 Internal Server Error** | Erro inesperado do servidor |

### 2.4 Estrutura de Resposta de Erro

```json
{
  "error": {
    "type": "https://api.fbso.org/errors/validation-error",
    "title": "Dados de entrada inválidos",
    "status": 400,
    "detail": "O campo 'cnpj' contém um valor duplicado para este tenant.",
    "instance": "/api/v1/business-units",
    "timestamp": "2026-07-13T18:00:00Z",
    "fields": [
      {
        "name": "cnpj",
        "reason": "CNPJ já cadastrado para unidade ativa neste tenant"
      }
    ]
  }
}
```

> **Padrão:** RFC 7807 — Problem Details for HTTP APIs

---

## 3. Visão Geral dos Recursos

| # | Recurso | Endpoint Base | Épico | Operações |
|:---|:---|:---|:---|:---|
| R-01 | **Tenants** | `/tenants` | EP-02 | CRUD + ativar/suspender/reativar |
| R-02 | **Plans** | `/plans` | EP-02 | CRUD + desativar |
| R-03 | **Subscriptions** | `/subscriptions` | EP-02 | Criar, alterar, suspender, histórico |
| R-04 | **Users** | `/users` | EP-03 | CRUD + convidar/desativar |
| R-05 | **Permissions** | `/permissions` | EP-03 | Atribuir/revogar papéis e vínculos |
| R-06 | **Business Units** | `/business-units` | EP-04 | CRUD + hierarquia + desativar |
| R-07 | **Products** | `/products` | EP-04 | CRUD + ativar/desativar |
| R-08 | **Dashboard Admin** | `/dashboard/admin` | EP-01 | Leitura de métricas operacionais |
| R-09 | **Dashboard Client** | `/dashboard/client` | EP-04 | Leitura de métricas do cliente |
| R-10 | **Onboarding** | `/onboarding` | EP-04 | Fluxo de primeiro acesso |
| R-11 | **Audit** | `/audit` | EP-02 | Consulta ao histórico de auditoria |

---

## 4. Contratos por Recurso (Esboço)

### R-01 — Tenants (`/tenants`)

**Descrição:** Gestão de contas de clientes (Tenants) pelo time FBSO.ORG.

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/tenants` | Listar todos os tenants (com paginação e filtros) | Admin FBSO |
| `GET` | `/tenants/{id}` | Obter detalhes de um tenant | Admin FBSO |
| `POST` | `/tenants` | Criar novo tenant (status: PENDING_ONBOARDING) | Admin FBSO |
| `PATCH` | `/tenants/{id}` | Atualizar dados cadastrais do tenant | Admin FBSO |
| `POST` | `/tenants/{id}/activate` | Ativar tenant manualmente | Admin FBSO |
| `POST` | `/tenants/{id}/suspend` | Suspender tenant (bloqueia acesso) | Admin FBSO |
| `POST` | `/tenants/{id}/reactivate` | Reativar tenant suspenso | Admin FBSO |
| `POST` | `/tenants/{id}/resend-invite` | Reenviar e-mail de ativação (idioma padronizado: inglês para endpoints) | Admin FBSO |

<details>
<summary>Schemas (esboço)</summary>

**Tenant (response):**
```json
{
  "id": "t-12345",
  "name_corporate": "Supermercado Bom Preço Ltda",
  "name_fantasy": "Bom Preço",
  "segment": "RETAIL",
  "status": "ACTIVE",
  "created_dt": "2026-07-13T10:00:00Z",
  "subscription": {
    "plan_name": "Core",
    "status": "ACTIVE",
    "start_date": "2026-07-13"
  }
}
```

**TenantCreate (request):**
```json
{
  "name_corporate": "Supermercado Bom Preço Ltda",
  "name_fantasy": "Bom Preço",
  "segment": "RETAIL"
}
```

**Query Params (GET /tenants):**
- `status` — filtrar por status (ACTIVE, SUSPENDED, PENDING_ONBOARDING, INACTIVE)
- `plan_id` — filtrar por plano
- `search` — busca textual por razão social ou nome fantasia
- `page` — número da página (default: 0)
- `size` — registros por página (default: 25)
- `sort` — ordenação (ex: `created_dt,desc`)

</details>

---

### R-02 — Plans (`/plans`)

**Descrição:** Configuração de planos comerciais do SaaS.

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/plans` | Listar todos os planos | Admin FBSO / Gestor Produto |
| `GET` | `/plans/{id}` | Obter detalhes de um plano | Admin FBSO / Gestor Produto |
| `POST` | `/plans` | Criar novo plano comercial | Gestor Produto |
| `PATCH` | `/plans/{id}` | Atualizar plano (gera nova versão) | Gestor Produto |
| `POST` | `/plans/{id}/deactivate` | Desativar plano (não disponível para novas assinaturas) | Gestor Produto |

<details>
<summary>Schemas (esboço)</summary>

**Plan (response):**
```json
{
  "id": "p-001",
  "name": "Plano Core",
  "description": "Módulos essenciais para operação",
  "price": 499.90,
  "recurrences": ["MONTHLY", "YEARLY"],
  "modules": [
    {"name": "FBSO Platform", "included": true},
    {"name": "Tributali-Engine", "included": false}
  ],
  "status": "ACTIVE",
  "version": 1
}
```

</details>

---

### R-03 — Subscriptions (`/subscriptions`)

**Descrição:** Vinculação de clientes a planos e gestão do ciclo de vida da assinatura.

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/tenants/{tenantId}/subscriptions` | Histórico de assinaturas do tenant | Admin FBSO / Admin Tenant |
| `GET` | `/tenants/{tenantId}/subscriptions/active` | Assinatura ativa do tenant | Admin FBSO / Admin Tenant |
| `POST` | `/tenants/{tenantId}/subscriptions` | Criar nova assinatura (vincula plano) | Admin FBSO |
| `POST` | `/subscriptions/{id}/change-plan` | Upgrade/downgrade de plano | Admin FBSO |
| `POST` | `/subscriptions/{id}/suspend` | Suspender assinatura (bloqueia módulos) | Admin FBSO |
| `POST` | `/subscriptions/{id}/reactivate` | Reativar assinatura | Admin FBSO |

<details>
<summary>Schemas (esboço)</summary>

**SubscriptionCreate (request):**
```json
{
  "plan_id": "p-001",
  "start_date": "2026-07-13",
  "end_date": null
}
```

**ChangePlan (request):**
```json
{
  "new_plan_id": "p-002",
  "effective_date": "2026-08-01"
}
```

</details>

---

### R-04 — Users (`/users`)

**Descrição:** Cadastro e gestão de usuários do ecossistema (por tenant). A autenticação é delegada ao Keycloak; este recurso gerencia o perfil e vínculos do usuário na plataforma.

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/users` | Listar usuários do tenant (contexto do JWT) | Admin Tenant |
| `GET` | `/users/{id}` | Detalhes do usuário | Admin Tenant / próprio usuário |
| `POST` | `/users` | Convidar novo usuário (dispara e-mail) | Admin Tenant |
| `PATCH` | `/users/{id}` | Atualizar dados do usuário | Admin Tenant / próprio usuário |
| `POST` | `/users/{id}/deactivate` | Desativar usuário (bloqueia acesso) | Admin Tenant |
| `POST` | `/users/{id}/reactivate` | Reativar usuário | Admin Tenant |

<details>
<summary>Schemas (esboço)</summary>

**UserCreate (request):**
```json
{
  "name": "João Silva",
  "email": "joao@bompreco.com.br",
  "role": "OPERATOR_BU",
  "business_unit_ids": ["bu-001"],
  "module_names": ["FBSO_PLATFORM"]
}
```

**User (response):**
```json
{
  "id": "u-987",
  "name": "João Silva",
  "email": "joao@bompreco.com.br",
  "status": "INVITE_PENDING",
  "role": "OPERATOR_BU",
  "business_units": [
    {"id": "bu-001", "corporate_name": "Bom Preço Matriz"}
  ],
  "modules": ["FBSO_PLATFORM"],
  "external_keycloak_id": null
}
```

</details>

---

### R-05 — Permissions (`/permissions`)

**Descrição:** Gestão de permissões e vínculos dos usuários (RBAC).

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/users/{userId}/permissions` | Permissões atuais do usuário | Admin Tenant |
| `PUT` | `/users/{userId}/permissions` | Substituir vínculos do usuário (BU + módulos) | Admin Tenant |
| `PATCH` | `/users/{userId}/permissions/role` | Alterar papel do usuário (atualização parcial) | Admin Tenant |
| `DELETE` | `/users/{userId}/permissions/{buId}` | Remover acesso a uma Unidade de Negócio | Admin Tenant |

<details>
<summary>Schemas (esboço)</summary>

**PermissionsUpdate (request):**
```json
{
  "business_unit_ids": ["bu-001", "bu-002"],
  "module_names": ["FBSO_PLATFORM", "STOREKEEPER_PORTAL"]
}
```

**RoleUpdate (request):**
```json
{
  "role": "MANAGER_BU"
}
```

</details>

---

### R-06 — Business Units (`/business-units`)

**Descrição:** Cadastro e gestão de Unidades de Negócio (CNPJs/filiais) vinculadas a um Tenant.

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/business-units` | Listar unidades do tenant (hierárquico) | Admin Tenant / Manager BU (apenas suas) |
| `GET` | `/business-units/{id}` | Detalhes da unidade | Admin Tenant / Manager BU / Operador BU |
| `POST` | `/business-units` | Cadastrar nova unidade | Admin Tenant / Manager BU |
| `PATCH` | `/business-units/{id}` | Atualizar dados da unidade | Admin Tenant / Manager BU |
| `POST` | `/business-units/{id}/deactivate` | Desativar unidade | Admin Tenant |

<details>
<summary>Schemas (esboço)</summary>

**BusinessUnit (response):**
```json
{
  "id": "bu-001",
  "parent_id": null,
  "cnpj": "12.345.678/0001-90",
  "corporate_name": "Supermercado Bom Preço Ltda",
  "tax_regime": "LUCRO_REAL",
  "address": {
    "street": "Av. Principal, 1000",
    "city": "São Paulo",
    "state": "SP",
    "zip": "01000-000"
  },
  "status": "ACTIVE",
  "hierarchy_level": 0,
  "children_count": 3
}
```

**BusinessUnitCreate (request):**
```json
{
  "cnpj": "12.345.678/0002-71",
  "corporate_name": "Bom Preço Filial Campinas",
  "tax_regime": "LUCRO_REAL",
  "parent_id": "bu-001",
  "address": {
    "street": "Rua do Comércio, 500",
    "city": "Campinas",
    "state": "SP",
    "zip": "13000-000"
  }
}
```

</details>

---

### R-07 — Products (`/products`)

**Descrição:** Catálogo de Produtos e Serviços por Unidade de Negócio.

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/products` | Listar produtos (filtrado pela BU ativa no seletor) | Admin Tenant / Manager BU / Operador BU / Auditor |
| `GET` | `/products/{id}` | Detalhes do produto | Todos com acesso à BU |
| `POST` | `/products` | Cadastrar novo produto | Admin Tenant / Manager BU |
| `PATCH` | `/products/{id}` | Atualizar produto | Admin Tenant / Manager BU |
| `POST` | `/products/{id}/deactivate` | Desativar produto | Admin Tenant / Manager BU |
| `POST` | `/products/{id}/activate` | Reativar produto | Admin Tenant / Manager BU |

<details>
<summary>Schemas (esboço)</summary>

**Product (response):**
```json
{
  "id": "prd-001",
  "business_unit_id": "bu-001",
  "name": "Arroz Tipo 1 — Pacote 5kg",
  "sku": "ARZ-5KG-001",
  "type": "PRODUCT",
  "description": "Arroz branco tipo 1, pacote de 5kg",
  "status": "ACTIVE",
  "fiscal_mapping": "NOT_MAPPED"
}
```

**Query Params (GET /products):**
- `type` — filtrar por tipo (PRODUCT, SERVICE)
- `status` — filtrar por status (ACTIVE, INACTIVE)
- `search` — busca textual por nome ou SKU
- `page`, `size`, `sort`

</details>

---

### R-08 — Dashboard Admin (`/dashboard/admin`)

**Descrição:** Métricas operacionais do SaaS para o time interno FBSO.ORG.

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/dashboard/admin/summary` | Indicadores principais do SaaS | Admin FBSO |
| `GET` | `/dashboard/admin/accounts-by-status` | Contas agrupadas por status | Admin FBSO |
| `GET` | `/dashboard/admin/accounts-by-plan` | Contas agrupadas por plano | Admin FBSO |
| `GET` | `/dashboard/admin/evolution` | Evolução da base ao longo do tempo | Admin FBSO |
| `GET` | `/dashboard/admin/alerts` | Alertas de atenção (onboarding pendente, suspensos) | Admin FBSO |

<details>
<summary>Schemas (esboço)</summary>

**AdminSummary (response):**
```json
{
  "total_accounts": 142,
  "active_accounts": 130,
  "suspended_accounts": 5,
  "pending_onboarding": 7,
  "new_accounts_this_month": 12,
  "period": {
    "from": "2026-07-01",
    "to": "2026-07-13"
  }
}
```

**Query Params:**
- `period` — predefinido (7d, 30d, 90d, current_month, current_year) ou customizado (`from` + `to`)

</details>

---

### R-09 — Dashboard Client (`/dashboard/client`)

**Descrição:** Dashboard do cliente logado, com informações da sua conta.

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/dashboard/client/summary` | Resumo da conta do cliente | Todos os papéis do tenant |
| `GET` | `/dashboard/client/notifications` | Notificações e lembretes | Todos os papéis do tenant |
| `POST` | `/dashboard/client/notifications/{id}/dismiss` | Dispensar notificação | Todos os papéis do tenant |

<details>
<summary>Schemas (esboço)</summary>

**ClientSummary (response):**
```json
{
  "active_business_units": 4,
  "total_products": 230,
  "current_plan": {
    "name": "Plano Core",
    "modules": ["FBSO Platform"]
  },
  "pending_invites": 2
}
```

</details>

---

### R-10 — Onboarding (`/onboarding`)

**Descrição:** Fluxo guiado de primeiro acesso do cliente.

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/onboarding/status` | Status atual do onboarding (qual passo, o que falta) | Cliente autenticado |
| `POST` | `/onboarding/step-1` | Confirmar/atualizar dados cadastrais | Cliente autenticado |
| `POST` | `/onboarding/step-2` | Cadastrar primeira Unidade de Negócio (Matriz) | Cliente autenticado |
| `GET` | `/onboarding/step-3` | Visualizar resumo do plano contratado | Cliente autenticado |
| `POST` | `/onboarding/complete` | Finalizar onboarding (tenant → ACTIVE) | Cliente autenticado |

<details>
<summary>Schemas (esboço)</summary>

**OnboardingStatus (response):**
```json
{
  "current_step": 2,
  "total_steps": 4,
  "steps": [
    {"step": 1, "label": "Confirmar dados", "completed": true},
    {"step": 2, "label": "Cadastrar empresa", "completed": false},
    {"step": 3, "label": "Seu plano", "completed": false},
    {"step": 4, "label": "Boas-vindas", "completed": false}
  ],
  "tenant_status": "PENDING_ONBOARDING"
}
```

</details>

---

### R-11 — Audit (`/audit`)

**Descrição:** Consulta ao histórico de auditoria das ações administrativas.

| Método | Path | Descrição | RBAC |
|:---|:---|:---|:---|
| `GET` | `/audit` | Listar registros de auditoria (com filtros) | Admin FBSO / Auditor |
| `GET` | `/audit/{id}` | Detalhes de um registro de auditoria | Admin FBSO / Auditor |
| `GET` | `/tenants/{tenantId}/audit` | Auditoria de um tenant específico | Admin FBSO |

<details>
<summary>Schemas (esboço)</summary>

**AuditEntry (response):**
```json
{
  "id": "aud-5001",
  "timestamp": "2026-07-13T15:30:00Z",
  "action": "TENANT_SUSPENDED",
  "entity_type": "TENANT",
  "entity_id": "t-12345",
  "actor": {
    "id": "u-admin-01",
    "name": "Admin FBSO"
  },
  "changes": {
    "field": "status",
    "previous_value": "ACTIVE",
    "new_value": "SUSPENDED"
  },
  "reason": "Inadimplência — fatura junho/2026 vencida"
}
```

**Query Params (GET /audit):**
- `entity_type` — TENANT, PLAN, SUBSCRIPTION, USER, PERMISSION, BUSINESS_UNIT, PRODUCT
- `action` — CREATED, UPDATED, SUSPENDED, REACTIVATED, DEACTIVATED, DELETED
- `from` / `to` — filtro por período
- `tenant_id` — filtrar por tenant (admin FBSO apenas)
- `page`, `size`, `sort`

</details>

---

## 5. Mapeamento RBAC × Endpoints

Matriz de permissões por papel (conforme RN10-01 do [FEATURES.md](./04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)):

| Recurso | Admin FBSO | Admin Tenant | Manager BU | Operator BU | Auditor |
|:---|:---:|:---:|:---:|:---:|:---:|
| **Tenants** (`/tenants`) | CRUD + ações | — | — | — | — |
| **Plans** (`/plans`) | CRUD | — | — | — | — |
| **Subscriptions** (`/subscriptions`) | CRUD | Ver (seu) | — | — | Ver |
| **Users** (`/users`) | — | CRUD | — | — | — |
| **Permissions** (`/permissions`) | — | CRUD | — | — | — |
| **Business Units** (`/business-units`) | — | CRUD | Ver + Editar (sua) | Ver (sua) | Ver |
| **Products** (`/products`) | — | CRUD | Criar, Editar, Ver | Ver | Ver |
| **Dashboard Admin** (`/dashboard/admin`) | Ver | — | — | — | — |
| **Dashboard Client** (`/dashboard/client`) | — | Ver | Ver | Ver | Ver |
| **Onboarding** (`/onboarding`) | — | Executar | Executar | Executar | — |
| **Audit** (`/audit`) | Ver (todos) | — | — | — | Ver (tenant) |

---

## 6. Versionamento e Evolução da API

### 6.1 Política de Versionamento

- **Versão corrente:** `v1` (prefixo `/api/v1/`)
- **Nova versão major:** Quando há breaking changes (ex: remoção de campo, mudança de tipo)
- **Adições non-breaking:** Novos endpoints, campos opcionais em response — NÃO exigem nova versão

### 6.2 Processo de Alteração de Contrato

```
1. Necessidade de mudança identificada (PO + Tech Leads)
2. API-CONTRACTS.md atualizado com a proposta (Seção "Alterações Propostas")
3. Revisão conjunta Backend + Frontend
4. OpenAPI YAML atualizado no backend
5. Comunicação formal: "API v1 — endpoint X alterado. Efetivo em: data"
6. Frontend atualiza mock (MSW) e tipos TypeScript
```

---

## 7. Próximos Passos

1. **[Pré-Kickoff]** Aprovar este esboço de API-CONTRACTS.md com os times
2. **[Pré-Kickoff]** Gerar `fbso-platform-api.yaml` (OpenAPI 3.0+) a partir destes contratos
3. **[Pré-Kickoff]** Criar coleção Postman/Bruno para testes manuais
4. **[Fase 0]** Gerar tipos TypeScript (frontend) e interfaces Java (backend) via codegen do OpenAPI
5. **[Fase 1+]** Manter este documento sincronizado com o OpenAPI YAML — o YAML é a fonte canônica; este documento é a referência de alto nível

---

## 8. Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Esboço inicial: 11 recursos, convenções, schemas de exemplo, matriz RBAC × endpoints, política de versionamento | Time Técnico |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 701-technologies-openapi, agile-ba-practices.*
🔍 *Revisado pelo skill caveman-review em 15/07/2026. Ajustes aplicados: PUT→PATCH para role update, onboarding verbs padronizados (POST), versionamento /api/v1 confirmado, idioma de ações documentado.*
