-- =============================================================================
-- U004: Rollback — Seed Data RBAC
-- =============================================================================
-- Reverte a migration V004__seed_rbac_matrix.sql removendo todos os registros
-- de role_resource e resource_action inseridos pelo seed.
-- =============================================================================

DELETE FROM fbso_platform.role_resource;
DELETE FROM fbso_platform.resource_action;
