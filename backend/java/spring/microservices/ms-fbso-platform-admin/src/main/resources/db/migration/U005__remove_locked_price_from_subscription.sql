-- Rollback U005: Remove locked_price e locked_recurrence da tabela subscription
-- Reverte V005__add_locked_price_to_subscription.sql
-- Data: 17/07/2026

ALTER TABLE fbso_platform.subscription
    DROP COLUMN IF EXISTS locked_price,
    DROP COLUMN IF EXISTS locked_recurrence;
