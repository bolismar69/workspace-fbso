-- =============================================================================
-- U009__rollback_tenant_id_product_service.sql
-- Rollback da V009 — remove RLS e coluna tenant_id de product_service
-- =============================================================================

-- 1. Remover política RLS
DROP POLICY IF EXISTS tenant_isolation ON fbso_platform.product_service;

-- 2. Desabilitar RLS
ALTER TABLE fbso_platform.product_service NO FORCE ROW LEVEL SECURITY;
ALTER TABLE fbso_platform.product_service DISABLE ROW LEVEL SECURITY;

-- 3. Remover índice
DROP INDEX IF EXISTS fbso_platform.idx_product_service_tenant_id;

-- 4. Remover coluna tenant_id
ALTER TABLE fbso_platform.product_service DROP COLUMN IF EXISTS tenant_id;
