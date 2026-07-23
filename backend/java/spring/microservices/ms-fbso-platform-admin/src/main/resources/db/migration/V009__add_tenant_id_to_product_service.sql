-- =============================================================================
-- V009__add_tenant_id_to_product_service.sql
-- PostgreSQL Row-Level Security — estende RLS para product_service
--
-- ADR-L07: Camada 1 de 3 do isolamento entre tenants.
-- A tabela product_service originalmente não possuía tenant_id próprio —
-- o isolamento era feito indiretamente via JOIN com business_unit.tenant_id.
-- Esta migration desnormaliza tenant_id para permitir RLS direto.
--
-- DT-130 (Sprint 6 Frente 1): Fecha gap de segurança — product_service
-- agora tem RLS com FORCE, alinhado com as outras 4 tabelas (V003).
-- =============================================================================

-- ---------------------------------------------------------------------------
-- 1. Adicionar coluna tenant_id (nullable inicialmente para permitir UPDATE)
-- ---------------------------------------------------------------------------
ALTER TABLE fbso_platform.product_service ADD COLUMN IF NOT EXISTS tenant_id UUID;

-- ---------------------------------------------------------------------------
-- 2. Preencher tenant_id via JOIN com business_unit
--    NOTA: Se houver product_service órfão (business_unit_id sem BU
--    correspondente), o tenant_id ficará NULL. A constraint NOT NULL
--    abaixo rejeitará e a migração falhará — isso é intencional para
--    garantir integridade dos dados.
-- ---------------------------------------------------------------------------
UPDATE fbso_platform.product_service ps
SET tenant_id = bu.tenant_id
FROM fbso_platform.business_unit bu
WHERE ps.business_unit_id = bu.id
  AND bu.deleted_dt IS NULL;

-- ---------------------------------------------------------------------------
-- 3. Tornar NOT NULL — se houver NULLs, a migração falha aqui
-- ---------------------------------------------------------------------------
ALTER TABLE fbso_platform.product_service ALTER COLUMN tenant_id SET NOT NULL;

-- ---------------------------------------------------------------------------
-- 4. Índice para performance de queries filtradas por tenant_id
-- ---------------------------------------------------------------------------
CREATE INDEX IF NOT EXISTS idx_product_service_tenant_id
    ON fbso_platform.product_service(tenant_id);

-- ---------------------------------------------------------------------------
-- 5. ENABLE + FORCE RLS (padrão V003 para todas as tabelas multi-tenant)
-- ---------------------------------------------------------------------------
ALTER TABLE fbso_platform.product_service ENABLE ROW LEVEL SECURITY;
ALTER TABLE fbso_platform.product_service FORCE ROW LEVEL SECURITY;

-- ---------------------------------------------------------------------------
-- 6. Política de isolamento (idêntica às outras 4 tabelas em V003)
-- ---------------------------------------------------------------------------
CREATE POLICY tenant_isolation ON fbso_platform.product_service
    FOR ALL
    USING (tenant_id = current_setting('app.current_tenant_id')::UUID)
    WITH CHECK (tenant_id = current_setting('app.current_tenant_id')::UUID);

-- =============================================================================
-- RESULTADO: product_service agora faz parte da camada 1 de defesa em
-- profundidade (RLS). Total de tabelas com RLS: 5 (subscription, user,
-- business_unit, product_service, audit_log).
-- =============================================================================
