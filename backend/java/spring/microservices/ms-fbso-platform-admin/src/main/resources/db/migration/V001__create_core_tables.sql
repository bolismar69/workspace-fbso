-- =============================================================================
-- V001__create_core_tables.sql
-- Cria o schema fbso_platform e as 11 tabelas Core da Fase 0
-- =============================================================================

CREATE SCHEMA IF NOT EXISTS fbso_platform;

-- ---------------------------------------------------------------------------
-- 1. tenant — Conta corporativa do cliente
-- ---------------------------------------------------------------------------
CREATE TABLE fbso_platform.tenant (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name_corporate  VARCHAR(255) NOT NULL,
    name_fantasy    VARCHAR(255),
    segment         VARCHAR(50)  NOT NULL,
    status          VARCHAR(30)  NOT NULL DEFAULT 'PENDING_ONBOARDING',
    -- Auditoria
    created_dt      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_dt      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    created_by      UUID,
    updated_by      UUID,
    deleted_dt      TIMESTAMPTZ,
    deleted_by      UUID
);

-- ---------------------------------------------------------------------------
-- 2. plan — Pacote comercial com módulos e preço
-- ---------------------------------------------------------------------------
CREATE TABLE fbso_platform.plan (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(255) NOT NULL,
    description     TEXT,
    price           NUMERIC(12,2) NOT NULL CHECK (price > 0),
    recurrence      VARCHAR(20)  NOT NULL DEFAULT 'MONTHLY',
    status          VARCHAR(30)  NOT NULL DEFAULT 'ACTIVE',
    version         INT          NOT NULL DEFAULT 1,
    -- Auditoria
    created_dt      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_dt      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    created_by      UUID,
    updated_by      UUID,
    deleted_dt      TIMESTAMPTZ,
    deleted_by      UUID
);

-- ---------------------------------------------------------------------------
-- 3. plan_module — Módulos incluídos em um plano
-- ---------------------------------------------------------------------------
CREATE TABLE fbso_platform.plan_module (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    plan_id         UUID         NOT NULL REFERENCES fbso_platform.plan(id),
    module_name     VARCHAR(100) NOT NULL,
    -- Auditoria
    created_dt      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_dt      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    created_by      UUID,
    updated_by      UUID,
    deleted_dt      TIMESTAMPTZ,
    deleted_by      UUID,
    UNIQUE (plan_id, module_name)
);

-- ---------------------------------------------------------------------------
-- 4. subscription — Vínculo Tenant × Plano com vigência
-- ---------------------------------------------------------------------------
CREATE TABLE fbso_platform.subscription (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id       UUID         NOT NULL REFERENCES fbso_platform.tenant(id),
    plan_id         UUID         NOT NULL REFERENCES fbso_platform.plan(id),
    start_date      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    end_date        TIMESTAMPTZ,
    status          VARCHAR(30)  NOT NULL DEFAULT 'ACTIVE',
    -- Auditoria
    created_dt      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_dt      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    created_by      UUID,
    updated_by      UUID,
    deleted_dt      TIMESTAMPTZ,
    deleted_by      UUID
);

-- ---------------------------------------------------------------------------
-- 5. user — Usuário do sistema
-- ---------------------------------------------------------------------------
CREATE TABLE fbso_platform."user" (
    id                      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id               UUID         NOT NULL REFERENCES fbso_platform.tenant(id),
    external_keycloak_id    UUID,
    email                   VARCHAR(320) NOT NULL,
    name                    VARCHAR(255) NOT NULL,
    status                  VARCHAR(30)  NOT NULL DEFAULT 'INVITE_PENDING',
    -- Auditoria
    created_dt              TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_dt              TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    created_by              UUID,
    updated_by              UUID,
    deleted_dt              TIMESTAMPTZ,
    deleted_by              UUID
);

-- ---------------------------------------------------------------------------
-- 6. user_permission — Vínculo Usuário × Unidade de Negócio × Papel
-- ---------------------------------------------------------------------------
CREATE TABLE fbso_platform.user_permission (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id             UUID NOT NULL REFERENCES fbso_platform."user"(id),
    business_unit_id    UUID NOT NULL,
    role                VARCHAR(30) NOT NULL,
    -- Auditoria
    created_dt          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_dt          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by          UUID,
    updated_by          UUID,
    deleted_dt          TIMESTAMPTZ,
    deleted_by          UUID,
    UNIQUE (user_id, business_unit_id)
);

-- ---------------------------------------------------------------------------
-- 7. resource_action — Recurso + Ação para RBAC
-- ---------------------------------------------------------------------------
CREATE TABLE fbso_platform.resource_action (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    resource_name   VARCHAR(100) NOT NULL,
    action          VARCHAR(50)  NOT NULL,
    -- Auditoria
    created_dt      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_dt      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    created_by      UUID,
    updated_by      UUID,
    deleted_dt      TIMESTAMPTZ,
    deleted_by      UUID,
    UNIQUE (resource_name, action)
);

-- ---------------------------------------------------------------------------
-- 8. role_resource — Papel × ResourceAction (matriz RN10-01)
-- ---------------------------------------------------------------------------
CREATE TABLE fbso_platform.role_resource (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role                VARCHAR(30) NOT NULL,
    resource_action_id  UUID NOT NULL REFERENCES fbso_platform.resource_action(id),
    -- Auditoria
    created_dt          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_dt          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by          UUID,
    updated_by          UUID,
    deleted_dt          TIMESTAMPTZ,
    deleted_by          UUID,
    UNIQUE (role, resource_action_id)
);

-- ---------------------------------------------------------------------------
-- 9. business_unit — CNPJ/filial vinculada a um Tenant
-- ---------------------------------------------------------------------------
CREATE TABLE fbso_platform.business_unit (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id       UUID         NOT NULL REFERENCES fbso_platform.tenant(id),
    parent_id       UUID         REFERENCES fbso_platform.business_unit(id),
    cnpj            VARCHAR(18)  NOT NULL,
    corporate_name  VARCHAR(255) NOT NULL,
    tax_regime      VARCHAR(30)  NOT NULL,
    street          VARCHAR(255),
    number          VARCHAR(20),
    complement      VARCHAR(100),
    neighborhood    VARCHAR(100),
    city            VARCHAR(100),
    state           VARCHAR(2),
    zip_code        VARCHAR(9),
    status          VARCHAR(30)  NOT NULL DEFAULT 'ACTIVE',
    -- Auditoria
    created_dt      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_dt      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    created_by      UUID,
    updated_by      UUID,
    deleted_dt      TIMESTAMPTZ,
    deleted_by      UUID
);

-- ---------------------------------------------------------------------------
-- 10. product_service — Item do catálogo vinculado a uma BU
-- ---------------------------------------------------------------------------
CREATE TABLE fbso_platform.product_service (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_unit_id    UUID         NOT NULL REFERENCES fbso_platform.business_unit(id),
    name                VARCHAR(255) NOT NULL,
    sku                 VARCHAR(50),
    type                VARCHAR(20)  NOT NULL DEFAULT 'SERVICE',
    description         TEXT,
    status              VARCHAR(30)  NOT NULL DEFAULT 'ACTIVE',
    -- Auditoria
    created_dt          TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_dt          TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    created_by          UUID,
    updated_by          UUID,
    deleted_dt          TIMESTAMPTZ,
    deleted_by          UUID
);

-- ---------------------------------------------------------------------------
-- 11. audit_log — Registro imutável de auditoria
-- ---------------------------------------------------------------------------
CREATE TABLE fbso_platform.audit_log (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    timestamp       TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    tenant_id       UUID         NOT NULL,
    action          VARCHAR(50)  NOT NULL,
    entity_type     VARCHAR(50)  NOT NULL,
    entity_id       UUID         NOT NULL,
    actor_id        UUID         NOT NULL,
    actor_name      VARCHAR(255),
    previous_value  JSONB,
    new_value       JSONB,
    reason          VARCHAR(500),
    -- auditoria mínima (apenas criação — registros são imutáveis)
    created_dt      TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
