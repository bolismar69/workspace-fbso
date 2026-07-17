-- =============================================================================
-- U006: Rollback — FK user_permission.business_unit_id
-- =============================================================================
-- Reverte a migration V006 removendo a constraint de FK.
-- =============================================================================

ALTER TABLE fbso_platform.user_permission
    DROP CONSTRAINT IF EXISTS fk_user_permission_business_unit;
