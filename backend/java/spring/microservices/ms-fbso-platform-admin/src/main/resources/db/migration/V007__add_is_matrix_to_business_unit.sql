-- V007: Adiciona flag is_matrix à tabela business_unit
-- Sprint 5 — Portal do Cliente (DT-107/T-142)
-- Feature F04-02: Primeira BU do onboarding é a Matriz (RN14-02)
--
-- DEFAULT false garante que BUs existentes (se houver) não sejam
-- incorretamente marcadas como Matriz.
-- parent_id IS NULL deixa de ser o único critério para identificar
-- a Matriz — a flag is_matrix é semanticamente explícita.

ALTER TABLE fbso_platform.business_unit
    ADD COLUMN IF NOT EXISTS is_matrix BOOLEAN NOT NULL DEFAULT false;

-- Índice para queries que filtram pela flag (ex: "liste todas as matrizes")
CREATE INDEX IF NOT EXISTS idx_business_unit_is_matrix
    ON fbso_platform.business_unit (tenant_id, is_matrix)
    WHERE deleted_dt IS NULL;
