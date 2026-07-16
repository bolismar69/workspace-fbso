-- =============================================================================
-- U003__disable_rls.sql
-- Rollback da V003 — remove políticas RLS e desabilita row-level security
-- =============================================================================

DROP POLICY IF EXISTS tenant_isolation ON fbso_platform.subscription;
ALTER TABLE fbso_platform.subscription DISABLE ROW LEVEL SECURITY;

DROP POLICY IF EXISTS tenant_isolation ON fbso_platform."user";
ALTER TABLE fbso_platform."user" DISABLE ROW LEVEL SECURITY;

DROP POLICY IF EXISTS tenant_isolation ON fbso_platform.business_unit;
ALTER TABLE fbso_platform.business_unit DISABLE ROW LEVEL SECURITY;

DROP POLICY IF EXISTS tenant_isolation ON fbso_platform.product_service;
ALTER TABLE fbso_platform.product_service DISABLE ROW LEVEL SECURITY;

DROP POLICY IF EXISTS tenant_isolation ON fbso_platform.audit_log;
ALTER TABLE fbso_platform.audit_log DISABLE ROW LEVEL SECURITY;
