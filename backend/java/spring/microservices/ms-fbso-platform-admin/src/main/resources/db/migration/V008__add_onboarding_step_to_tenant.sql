-- V008: Adiciona coluna onboarding_step à tabela tenant
-- Sprint 5 — Frente 3 (T-060: OnboardingService)
-- Estados: NOT_STARTED, STEP1_DONE, STEP2_DONE, STEP3_DONE, COMPLETED

ALTER TABLE fbso_platform.tenant
    ADD COLUMN IF NOT EXISTS onboarding_step VARCHAR(20) NOT NULL DEFAULT 'NOT_STARTED';

-- Índice para consultas por status de onboarding
CREATE INDEX IF NOT EXISTS idx_tenant_onboarding_step
    ON fbso_platform.tenant (onboarding_step)
    WHERE deleted_dt IS NULL;
