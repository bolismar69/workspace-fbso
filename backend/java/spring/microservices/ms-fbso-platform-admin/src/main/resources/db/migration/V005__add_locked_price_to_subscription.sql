-- Migration V005: Adiciona locked_price e locked_recurrence à tabela subscription
-- Débito: DT-009 — RN06-02: alteração de preço no plano não afeta assinaturas existentes
-- Data: 17/07/2026
-- Sprint: 3 (Frente 0 — Correções Pré-Sprint)

ALTER TABLE fbso_platform.subscription
    ADD COLUMN IF NOT EXISTS locked_price NUMERIC(10,2),
    ADD COLUMN IF NOT EXISTS locked_recurrence VARCHAR(20);

-- Nota: As colunas são NULLABLE inicialmente porque assinaturas existentes
-- (se houver) não terão esses valores. O SubscriptionService (T-033, Sprint 3)
-- é responsável por preencher locked_price e locked_recurrence no momento
-- da criação da assinatura, copiando do Plan corrente.
--
-- Rollback: ver U005__remove_locked_price_from_subscription.sql
