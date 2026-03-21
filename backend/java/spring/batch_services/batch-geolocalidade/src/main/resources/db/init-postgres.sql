-- Criação idempotente dos schemas usados pela aplicação
-- - spring_batch: tabelas internas do Spring Batch (BATCH_*)
-- - localidade: tabelas de negócio (JPA)
CREATE SCHEMA IF NOT EXISTS spring_batch;
CREATE SCHEMA IF NOT EXISTS localidade;
