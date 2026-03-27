-- 4. Inserção do Imposto Base
INSERT INTO individual_tax_rates.tax_definitions (tax_code, name, sphere)
VALUES ('IRPF', 'Imposto de Renda Pessoa Física', 'F')
ON CONFLICT (tax_code) DO NOTHING;

-- 5. Inserção das Tabelas Progressivas (Exemplo Simplificado 2024-2026)
-- Nota: Para fins de teste, estamos usando a tabela vigente a partir de 02/2024.
INSERT INTO individual_tax_rates.tax_rules_history 
(tax_definition_id, description, range_min, range_max, aliq_percent, deduction_val, valid_from, valid_to)
SELECT id, 'Tabela Progressiva Mensal', 0.00, 2259.20, 0.0, 0.00, '2024-02-01', NULL FROM (SELECT id FROM individual_tax_rates.tax_definitions WHERE tax_code = 'IRPF');

INSERT INTO individual_tax_rates.tax_rules_history 
(tax_definition_id, description, range_min, range_max, aliq_percent, deduction_val, valid_from, valid_to)
SELECT id, 'Tabela Progressiva Mensal', 2259.21, 2826.65, 7.5, 169.44, '2024-02-01', NULL FROM (SELECT id FROM individual_tax_rates.tax_definitions WHERE tax_code = 'IRPF');

INSERT INTO individual_tax_rates.tax_rules_history 
(tax_definition_id, description, range_min, range_max, aliq_percent, deduction_val, valid_from, valid_to)
SELECT id, 'Tabela Progressiva Mensal', 2826.66, 3751.05, 15.0, 381.44, '2024-02-01', NULL FROM (SELECT id FROM individual_tax_rates.tax_definitions WHERE tax_code = 'IRPF');

INSERT INTO individual_tax_rates.tax_rules_history 
(tax_definition_id, description, range_min, range_max, aliq_percent, deduction_val, valid_from, valid_to)
SELECT id, 'Tabela Progressiva Mensal', 3751.06, 4664.68, 22.5, 662.77, '2024-02-01', NULL FROM (SELECT id FROM individual_tax_rates.tax_definitions WHERE tax_code = 'IRPF');

INSERT INTO individual_tax_rates.tax_rules_history 
(tax_definition_id, description, range_min, range_max, aliq_percent, deduction_val, valid_from, valid_to)
SELECT id, 'Tabela Progressiva Mensal', 4664.69, NULL, 27.5, 896.00, '2024-02-01', NULL FROM (SELECT id FROM individual_tax_rates.tax_definitions WHERE tax_code = 'IRPF');

-- TABELA DE CONFIGS

-- 1. Tabela para chaves de configuração (Ex: limites, valores fixos)
CREATE TABLE IF NOT EXISTS individual_tax_rates.tax_configs (
    id SERIAL PRIMARY KEY,
    tax_code VARCHAR(20) NOT NULL, -- IRPF, INSS, etc
    config_key VARCHAR(50) NOT NULL,
    config_value DECIMAL(18,4) NOT NULL,
    description TEXT,
    valid_from DATE NOT NULL,
    valid_to DATE,
    UNIQUE(tax_code, config_key, valid_from)
);

-- 2. Carga das Regras de Negócio (Valores vigentes para 2024-2026)
INSERT INTO individual_tax_rates.tax_configs 
(tax_code, config_key, config_value, description, valid_from)
VALUES 
-- IRPF: Dedução por Dependente
('IRPF', 'dependent_deduction_monthly', 189.59, 'Valor mensal por dependente', '2024-02-01'),
('IRPF', 'dependent_deduction_annual', 2275.08, 'Valor anual por dependente', '2024-02-01'),

-- IRPF: Limites de Educação
('IRPF', 'education_limit_annual', 3561.50, 'Limite anual de instrução por pessoa', '2024-02-01'),

-- IRPF: Desconto Simplificado
('IRPF', 'simplified_discount_monthly_limit', 564.80, 'Teto mensal do desconto simplificado (20%)', '2024-02-01'),
('IRPF', 'simplified_discount_annual_limit', 6777.60, 'Teto anual do desconto simplificado (20%)', '2024-02-01'),

-- IRPF: Previdência PGBL
('IRPF', 'pgbl_limit_percentage', 12.00, 'Limite de dedução PGBL sobre renda bruta', '2024-02-01')

-- IRPF: Limites de Saude
('IRPF', 'health_limit_annual', 999999999999.99, 'Limite de dedução em saude sobre renda bruta', '2024-02-01');
