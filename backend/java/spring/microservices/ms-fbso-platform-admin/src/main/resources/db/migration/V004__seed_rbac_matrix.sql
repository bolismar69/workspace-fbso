-- =============================================================================
-- V004: Seed Data — Matriz de Permissões RBAC (RN10-01)
-- =============================================================================
-- Popula resource_action com 8 recursos × 4 ações = 32 combinações
-- Popula role_resource com a matriz completa (4 papéis)
--
-- Matriz RN10-01 (SPRINT-CARD restritiva):
--   ADMIN_TENANT  → acesso total (todos os 32 resource_actions)
--   MANAGER_BU    → BUSINESS_UNIT(view,create,edit) + PRODUCT_SERVICE(view,create,edit)
--   OPERATOR_BU   → BUSINESS_UNIT(view) + PRODUCT_SERVICE(view)
--   AUDITOR       → AUDIT(view)
--
-- Rollback: U004__rollback_rbac_seed.sql
-- =============================================================================

-- ---------------------------------------------------------------------------
-- 1. RESOURCE_ACTION: 8 recursos × 4 ações
-- ---------------------------------------------------------------------------
INSERT INTO fbso_platform.resource_action (resource_name, action) VALUES
    -- DASHBOARD
    ('DASHBOARD', 'view'),
    -- TENANT
    ('TENANT', 'view'),
    ('TENANT', 'create'),
    ('TENANT', 'edit'),
    ('TENANT', 'suspend'),
    ('TENANT', 'reactivate'),
    -- PLAN
    ('PLAN', 'view'),
    ('PLAN', 'create'),
    ('PLAN', 'edit'),
    ('PLAN', 'deactivate'),
    -- SUBSCRIPTION
    ('SUBSCRIPTION', 'view'),
    ('SUBSCRIPTION', 'create'),
    ('SUBSCRIPTION', 'edit'),
    -- USER
    ('USER', 'view'),
    ('USER', 'create'),
    ('USER', 'edit'),
    ('USER', 'delete'),
    -- PERMISSION
    ('PERMISSION', 'view'),
    ('PERMISSION', 'edit'),
    -- BUSINESS_UNIT
    ('BUSINESS_UNIT', 'view'),
    ('BUSINESS_UNIT', 'create'),
    ('BUSINESS_UNIT', 'edit'),
    -- PRODUCT_SERVICE
    ('PRODUCT_SERVICE', 'view'),
    ('PRODUCT_SERVICE', 'create'),
    ('PRODUCT_SERVICE', 'edit'),
    -- AUDIT
    ('AUDIT', 'view');

-- ---------------------------------------------------------------------------
-- 2. ROLE_RESOURCE: Matriz RN10-01 (papel × resource_action)
-- ---------------------------------------------------------------------------

-- 2.1 ADMIN_TENANT — acesso total (todos os 28 resource_actions acima)
INSERT INTO fbso_platform.role_resource (role, resource_action_id)
SELECT 'ADMIN_TENANT', id FROM fbso_platform.resource_action;

-- 2.2 MANAGER_BU — BUSINESS_UNIT(view,create,edit) + PRODUCT_SERVICE(view,create,edit)
INSERT INTO fbso_platform.role_resource (role, resource_action_id)
SELECT 'MANAGER_BU', id FROM fbso_platform.resource_action
WHERE (resource_name = 'BUSINESS_UNIT'   AND action IN ('view', 'create', 'edit'))
   OR (resource_name = 'PRODUCT_SERVICE' AND action IN ('view', 'create', 'edit'));

-- 2.3 OPERATOR_BU — BUSINESS_UNIT(view) + PRODUCT_SERVICE(view)
INSERT INTO fbso_platform.role_resource (role, resource_action_id)
SELECT 'OPERATOR_BU', id FROM fbso_platform.resource_action
WHERE (resource_name = 'BUSINESS_UNIT'   AND action = 'view')
   OR (resource_name = 'PRODUCT_SERVICE' AND action = 'view');

-- 2.4 AUDITOR — AUDIT(view)
INSERT INTO fbso_platform.role_resource (role, resource_action_id)
SELECT 'AUDITOR', id FROM fbso_platform.resource_action
WHERE resource_name = 'AUDIT' AND action = 'view';
