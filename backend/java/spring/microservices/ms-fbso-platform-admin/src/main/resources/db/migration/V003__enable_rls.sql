-- =============================================================================
-- V003__enable_rls.sql
-- PostgreSQL Row-Level Security — defesa em profundidade para isolamento multi-tenant
--
-- ADR-L07: Camada 1 de 3 do isolamento entre tenants.
-- Toda tabela com coluna tenant_id recebe política que força o filtro.
-- A variável de sessão app.current_tenant_id é configurada pelo JwtAuthenticationFilter.
-- =============================================================================

-- ---------------------------------------------------------------------------
-- 1. subscription — Vínculo Tenant × Plano
-- ---------------------------------------------------------------------------
ALTER TABLE fbso_platform.subscription ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON fbso_platform.subscription
    FOR ALL
    USING (tenant_id = current_setting('app.current_tenant_id')::UUID)
    WITH CHECK (tenant_id = current_setting('app.current_tenant_id')::UUID);

-- ---------------------------------------------------------------------------
-- 2. user — Usuário do sistema
-- ---------------------------------------------------------------------------
ALTER TABLE fbso_platform."user" ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON fbso_platform."user"
    FOR ALL
    USING (tenant_id = current_setting('app.current_tenant_id')::UUID)
    WITH CHECK (tenant_id = current_setting('app.current_tenant_id')::UUID);

-- ---------------------------------------------------------------------------
-- 3. business_unit — CNPJ/filial vinculada a um Tenant
-- ---------------------------------------------------------------------------
ALTER TABLE fbso_platform.business_unit ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON fbso_platform.business_unit
    FOR ALL
    USING (tenant_id = current_setting('app.current_tenant_id')::UUID)
    WITH CHECK (tenant_id = current_setting('app.current_tenant_id')::UUID);

-- ---------------------------------------------------------------------------
-- 4. product_service — Catálogo de Produtos/Serviços por BU
-- ---------------------------------------------------------------------------
ALTER TABLE fbso_platform.product_service ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON fbso_platform.product_service
    FOR ALL
    USING (tenant_id = current_setting('app.current_tenant_id')::UUID)
    WITH CHECK (tenant_id = current_setting('app.current_tenant_id')::UUID);

-- ---------------------------------------------------------------------------
-- 5. audit_log — Trilha de auditoria (imutável, segmentada por tenant)
-- ---------------------------------------------------------------------------
ALTER TABLE fbso_platform.audit_log ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON fbso_platform.audit_log
    FOR ALL
    USING (tenant_id = current_setting('app.current_tenant_id')::UUID)
    WITH CHECK (tenant_id = current_setting('app.current_tenant_id')::UUID);

-- =============================================================================
-- NOTA: As tabelas abaixo NÃO recebem RLS porque são globais/compartilhadas:
--   - tenant       → root (Admin FBSO vê todos; isolamento via BaseRepository)
--   - plan         → catálogo global de planos
--   - plan_module  → vinculado a plan (global)
--   - user_permission → herda isolamento via JOIN com user
--   - resource_action  → matriz RBAC global
--   - role_resource    → matriz RBAC global
-- =============================================================================
