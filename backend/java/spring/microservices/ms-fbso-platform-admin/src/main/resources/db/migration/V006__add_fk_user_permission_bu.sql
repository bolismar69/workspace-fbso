-- =============================================================================
-- V006: FK user_permission.business_unit_id → business_unit.id
-- =============================================================================
-- Adiciona a constraint de integridade referencial que estava ausente desde V001.
-- Garante que user_permission não referencie business_units inexistentes.
--
-- Rollback: U006__drop_fk_user_permission_bu.sql
-- =============================================================================

ALTER TABLE fbso_platform.user_permission
    ADD CONSTRAINT fk_user_permission_business_unit
    FOREIGN KEY (business_unit_id) REFERENCES fbso_platform.business_unit(id);
