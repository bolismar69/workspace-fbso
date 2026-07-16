-- =============================================================================
-- V002__create_partial_unique_indexes.sql
-- Índices únicos parciais (WHERE deleted_dt IS NULL) + índices de desempenho
-- =============================================================================

-- ---------------------------------------------------------------------------
-- Índices Únicos Parciais — garantem unicidade apenas entre registros ativos
-- Permitem reúso de CNPJ/e-mail/SKU após soft delete
-- ---------------------------------------------------------------------------

-- business_unit: CNPJ único entre ativos do mesmo tenant
CREATE UNIQUE INDEX unique_cnpj_active
    ON fbso_platform.business_unit (tenant_id, cnpj)
    WHERE deleted_dt IS NULL;

-- user: e-mail único entre ativos do mesmo tenant
CREATE UNIQUE INDEX unique_email_active
    ON fbso_platform."user" (tenant_id, email)
    WHERE deleted_dt IS NULL;

-- product_service: SKU único entre ativos da mesma BU (apenas quando SKU é informado)
CREATE UNIQUE INDEX unique_sku_active
    ON fbso_platform.product_service (business_unit_id, sku)
    WHERE deleted_dt IS NULL AND sku IS NOT NULL;

-- tenant: razão social única entre ativos
CREATE UNIQUE INDEX unique_corporate_name_active
    ON fbso_platform.tenant (name_corporate)
    WHERE deleted_dt IS NULL;

-- ---------------------------------------------------------------------------
-- Índices de Desempenho — queries frequentes
-- ---------------------------------------------------------------------------

-- Dashboard: contagem por status
CREATE INDEX idx_tenant_status ON fbso_platform.tenant (status) WHERE deleted_dt IS NULL;

-- Lista de tenants: ordenação por criação
CREATE INDEX idx_tenant_created ON fbso_platform.tenant (created_dt DESC) WHERE deleted_dt IS NULL;

-- Subscriptions ativas por tenant (RN07-01 — validação de unicidade)
CREATE INDEX idx_subscription_tenant_active
    ON fbso_platform.subscription (tenant_id, status)
    WHERE deleted_dt IS NULL AND status = 'ACTIVE';

-- Produtos por business unit
CREATE INDEX idx_product_service_bu
    ON fbso_platform.product_service (business_unit_id)
    WHERE deleted_dt IS NULL;

-- Auditoria: consulta por período
CREATE INDEX idx_audit_log_timestamp ON fbso_platform.audit_log (timestamp DESC);

-- Auditoria: consulta por tenant
CREATE INDEX idx_audit_log_tenant ON fbso_platform.audit_log (tenant_id);

-- Auditoria: consulta por tipo de entidade
CREATE INDEX idx_audit_log_entity ON fbso_platform.audit_log (entity_type, entity_id);

-- Usuários por tenant
CREATE INDEX idx_user_tenant ON fbso_platform."user" (tenant_id) WHERE deleted_dt IS NULL;

-- Permissões por usuário
CREATE INDEX idx_user_permission_user ON fbso_platform.user_permission (user_id) WHERE deleted_dt IS NULL;

-- Permissões por business unit
CREATE INDEX idx_user_permission_bu ON fbso_platform.user_permission (business_unit_id) WHERE deleted_dt IS NULL;
