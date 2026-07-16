CREATE SCHEMA IF NOT EXISTS billing_tax_rates;

-- 1. Tabela de Alíquotas de ICMS (Matriz Origem/Destino)
CREATE TABLE billing_tax_rates.icms_rules (
    id SERIAL PRIMARY KEY,
    uf_origem CHAR(2) NOT NULL,
    uf_destino CHAR(2) NOT NULL,
    aliquota_interna DECIMAL(5,2) NOT NULL, -- Ex: 18.00
    aliquota_interestadual DECIMAL(5,2) NOT NULL, -- Ex: 7.00 ou 12.00
    cst_padrao CHAR(2) DEFAULT '00',
    reducao_base DECIMAL(5,2) DEFAULT 0.00,
    active BOOLEAN DEFAULT true,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 2. Tabela de Alíquotas PIS/COFINS por Regime
CREATE TABLE billing_tax_rates.federal_tax_rules (
    id SERIAL PRIMARY KEY,
    regime_tributario VARCHAR(50) NOT NULL, -- 'LUCRO_REAL', 'LUCRO_PRESUMIDO'
    cst_pis CHAR(2) NOT NULL,
    cst_cofins CHAR(2) NOT NULL,
    aliquota_pis DECIMAL(5,2) NOT NULL,     -- Ex: 1.65
    aliquota_cofins DECIMAL(5,2) NOT NULL,  -- Ex: 7.60
    exclui_icms_base BOOLEAN DEFAULT true,  -- Tese do Século
    active BOOLEAN DEFAULT true,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 3. Tabela de Exceções por NCM (Produtos com Alíquota Zero, Monofásicos ou ST)
CREATE TABLE billing_tax_rates.product_tax_exceptions (
    id SERIAL PRIMARY KEY,
    ncm VARCHAR(10) NOT NULL,
    uf_destino CHAR(2), -- Pode ser específica para um estado
    cst_pis CHAR(2),
    cst_cofins CHAR(2),
    cst_icms CHAR(2),
    aliquota_pis_unitario DECIMAL(10,4),    -- Para CST 03 (Unidade de Medida)
    aliquota_cofins_unitario DECIMAL(10,4), -- Para CST 03
    mva_st DECIMAL(5,2),                    -- Margem de Valor Agregado para ICMS ST
    active BOOLEAN DEFAULT true,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Exemplo de Carga de Dados (População Inicial)
-- Configurando ICMS SP -> RJ (Interestadual 12%)
INSERT INTO billing_tax_rates.icms_rules (uf_origem, uf_destino, aliquota_interna, aliquota_interestadual)
VALUES ('SP', 'RJ', 18.00, 12.00);

-- Configurando PIS/COFINS para Lucro Real com Exclusão de ICMS
INSERT INTO billing_tax_rates.federal_tax_rules (regime_tributario, cst_pis, cst_cofins, aliquota_pis, aliquota_cofins)
VALUES ('LUCRO_REAL', '01', '01', 1.65, 7.60);

-- 1. Tabela de Equivalência (Conforme seu README-TABELA-CST-CSON.md)
CREATE TABLE billing_tax_rates.tax_equivalence (
    id SERIAL PRIMARY KEY,
    csosn CHAR(3) NOT NULL,
    cst_equivalente CHAR(2) NOT NULL,
    permite_credito BOOLEAN DEFAULT false,  
    descricao TEXT
);


-- 2. Atualização da tabela de exceções para aceitar CSOSN
ALTER TABLE billing_tax_rates.product_tax_exceptions 
ADD COLUMN csosn CHAR(3);

-- Tabela para armazenar os anexos e faixas do Simples Nacional
CREATE TABLE billing_tax_rates.simples_nacional_rates (
    id SERIAL PRIMARY KEY,
    anexo VARCHAR(10) NOT NULL, -- 'ANEXO_I' (Comércio), 'ANEXO_II' (Indústria)
    faixa INTEGER NOT NULL,     -- 1 a 6
    receita_min DECIMAL(15,2),
    receita_max DECIMAL(15,2),
    aliquota_nominal DECIMAL(5,4) NOT NULL, -- Alíquota da faixa (ex: 0.0400 para 4%)
    valor_deduzir DECIMAL(15,2) NOT NULL,   -- Parcela a deduzir da faixa
    percentual_icMS DECIMAL(5,4) NOT NULL,  -- % do ICMS dentro do Simples (ex: 0.3350 para 33,5%)
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index para busca rápida por valor de faturamento
CREATE INDEX idx_rbt12_range ON billing_tax_rates.simples_nacional_rates (receita_min, receita_max);

-- Exemplo de População (Baseado no Anexo I - Comércio)
INSERT INTO billing_tax_rates.simples_nacional_rates 
(anexo, faixa, faturamento_min, faturamento_max, aliquota_nominal, valor_deduzir, percentual_icms)
VALUES 
('ANEXO_I', 1, 0.00, 180000.00, 0.0400, 0.00, 0.3350),
('ANEXO_I', 2, 180000.01, 360000.00, 0.0730, 5940.00, 0.3350),
('ANEXO_I', 3, 360000.01, 720000.00, 0.0950, 13860.00, 0.3350);

-- EVOLUCAO MODELO DE DADOS
ALTER TABLE billing_tax_rates.icms_rules 
ADD COLUMN percentual_fcp DECIMAL(5,2) DEFAULT 0.00,
ADD COLUMN mva_padrao DECIMAL(5,2) DEFAULT 0.00,
ADD COLUMN possui_protocolo_st BOOLEAN DEFAULT FALSE;

ALTER TABLE billing_tax_rates.product_tax_exceptions 
ADD COLUMN possui_protocolo_st BOOLEAN DEFAULT FALSE,
ADD COLUMN aliquota_interna_destino DECIMAL(5,2), -- Ex: Alíquota diferenciada para o item na UF destino
ADD COLUMN aliquota_interestadual DECIMAL(5,2); -- Ex: Itens importados (4%) vs Nacionais (7/12%)

ALTER TABLE billing_tax_rates.product_tax_exceptions 
ADD COLUMN IF NOT EXISTS possui_protocolo_st BOOLEAN DEFAULT FALSE,
ADD COLUMN IF NOT EXISTS aliquota_interna_destino DECIMAL(5,2),
ADD COLUMN IF NOT EXISTS aliquota_interestadual DECIMAL(5,2),
ADD COLUMN IF NOT EXISTS percentual_fcp DECIMAL(5,2) DEFAULT 0.00,
ADD COLUMN IF NOT EXISTS reducao_base DECIMAL(5,2) DEFAULT 0.00;

COMMENT ON COLUMN billing_tax_rates.product_tax_exceptions.possui_protocolo_st IS 'Indica se há acordo/protocolo de ST para este NCM entre as UFs';
COMMENT ON COLUMN billing_tax_rates.product_tax_exceptions.aliquota_interna_destino IS 'Override da alíquota interna da UF destino para este produto específico';


-- ATUALIZACAO 
-- Exemplo de ajuste para operações saindo de SP para RJ
UPDATE billing_tax_rates.icms_rules
SET 
    aliquota_interestadual = 12.0,
    aliquota_interna = 18.0, -- Alíquota base do RJ
    percentual_fcp = 2.0,    -- Adicional do RJ
    cst_padrao = '00',
    possui_protocolo_st = false, -- Default para regra geral sem NCM específico
    reducao_base = 0.0,
    mva_padrao = 0.0
WHERE uf_origem = 'SP' AND uf_destino = 'RJ';

-- Exemplo para produtos com ST (Ex: NCM de Pneus citado no brainstorm)
-- Note que aqui usamos a tabela de exceções que você criou para maior granularidade
UPDATE billing_tax_rates.product_tax_exceptions
SET 
    mva_st = 40.0,
    cst_icms = '010',
    possui_protocolo_st = true
WHERE ncm = '40111000' AND uf_destino = 'RJ';

-- REFATORANDO AS TABELAS - NOVAS COLUNAS - DROPANDO COLUNAS QUE SERAO SUBSTITUIDAS
ALTER TABLE billing_tax_rates.icms_rules ADD COLUMN IF NOT EXISTS inicio_validade date NOT NULL DEFAULT CURRENT_DATE;
ALTER TABLE billing_tax_rates.icms_rules ADD COLUMN IF NOT EXISTS final_validade date NULL;
ALTER TABLE billing_tax_rates.icms_rules DROP COLUMN IF EXISTS active;
ALTER TABLE billing_tax_rates.icms_rules DROP COLUMN IF EXISTS updated_at;
ALTER TABLE billing_tax_rates.icms_rules ADD COLUMN IF NOT EXISTS criado_em date NOT NULL DEFAULT CURRENT_DATE;
ALTER TABLE billing_tax_rates.icms_rules ADD COLUMN IF NOT EXISTS atualizado_em date NOT NULL DEFAULT CURRENT_DATE;

ALTER TABLE billing_tax_rates.federal_tax_rules ADD COLUMN IF NOT EXISTS inicio_validade date NOT NULL DEFAULT CURRENT_DATE;
ALTER TABLE billing_tax_rates.federal_tax_rules ADD COLUMN IF NOT EXISTS final_validade date NULL;
ALTER TABLE billing_tax_rates.federal_tax_rules DROP COLUMN IF EXISTS active;
ALTER TABLE billing_tax_rates.federal_tax_rules DROP COLUMN IF EXISTS updated_at;
ALTER TABLE billing_tax_rates.federal_tax_rules ADD COLUMN IF NOT EXISTS criado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE billing_tax_rates.federal_tax_rules ADD COLUMN IF NOT EXISTS atualizado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE billing_tax_rates.product_tax_exceptions ADD COLUMN IF NOT EXISTS inicio_validade date NOT NULL DEFAULT CURRENT_DATE;
ALTER TABLE billing_tax_rates.product_tax_exceptions ADD COLUMN IF NOT EXISTS final_validade date NULL;
ALTER TABLE billing_tax_rates.product_tax_exceptions DROP COLUMN IF EXISTS active;
ALTER TABLE billing_tax_rates.product_tax_exceptions DROP COLUMN IF EXISTS updated_at;
ALTER TABLE billing_tax_rates.product_tax_exceptions ADD COLUMN IF NOT EXISTS criado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE billing_tax_rates.product_tax_exceptions ADD COLUMN IF NOT EXISTS atualizado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE billing_tax_rates.simples_nacional_rates ADD COLUMN IF NOT EXISTS inicio_validade date NOT NULL DEFAULT CURRENT_DATE;
ALTER TABLE billing_tax_rates.simples_nacional_rates ADD COLUMN IF NOT EXISTS final_validade date NULL;
ALTER TABLE billing_tax_rates.simples_nacional_rates DROP COLUMN IF EXISTS active;
ALTER TABLE billing_tax_rates.simples_nacional_rates DROP COLUMN IF EXISTS updated_at;
ALTER TABLE billing_tax_rates.simples_nacional_rates ADD COLUMN IF NOT EXISTS criado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE billing_tax_rates.simples_nacional_rates ADD COLUMN IF NOT EXISTS atualizado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE billing_tax_rates.tax_equivalence ADD COLUMN IF NOT EXISTS inicio_validade date NOT NULL DEFAULT CURRENT_DATE;
ALTER TABLE billing_tax_rates.tax_equivalence ADD COLUMN IF NOT EXISTS final_validade date NULL;
ALTER TABLE billing_tax_rates.tax_equivalence DROP COLUMN IF EXISTS active;
ALTER TABLE billing_tax_rates.tax_equivalence DROP COLUMN IF EXISTS updated_at;
ALTER TABLE billing_tax_rates.tax_equivalence ADD COLUMN IF NOT EXISTS criado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE billing_tax_rates.tax_equivalence ADD COLUMN IF NOT EXISTS atualizado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- CRIANDO FUNCAO PARA CONTROLAR DATA DE FIM DE VALIDADE DE REGRAS
CREATE OR REPLACE FUNCTION public.fechar_fim_validade_generica()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
BEGIN
    -- Executa o update na tabela que disparou a trigger (TG_TABLE_NAME)
    EXECUTE format('UPDATE %I SET fim_validade = ( $1.inicio_validade - 1 )
                    WHERE codigo_cst = $1.codigo_cst 
                      AND uf_origem = $1.uf_origem 
                      AND uf_destino = $1.uf_destino
                      AND fim_validade IS NULL', TG_TABLE_NAME)
    USING NEW;

    RETURN NEW;
END;
$function$
;

-- CRIAR TRIGGERS PARA USAREM A FUNÇÃO DE FECHAMENTO DE VALIDADE
create trigger billing_tax_rates_federal_tax_rules_fim_validade before
insert
    on
    billing_tax_rates.federal_tax_rules for each row execute function fechar_fim_validade_generica();

create trigger billing_tax_rates_icms_rules_fim_validade before
insert
    on
    billing_tax_rates.icms_rules for each row execute function fechar_fim_validade_generica();

create trigger billing_tax_rates_product_tax_exceptions_fim_validade before
insert
    on
    billing_tax_rates.product_tax_exceptions for each row execute function fechar_fim_validade_generica();

create trigger billing_tax_rates_simples_nacional_rates_fim_validade before
insert
    on
    billing_tax_rates.simples_nacional_rates for each row execute function fechar_fim_validade_generica();

create trigger billing_tax_rates_tax_equivalence_fim_validade before
insert
    on
    billing_tax_rates.tax_equivalence for each row execute function fechar_fim_validade_generica();

-- FUNCAO PARA CONTROLAR A DATA DE ATUALIZACAO (updated_at -> atualizado_em)
CREATE OR REPLACE FUNCTION public.atualizar_data_atualizacao_generica()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
BEGIN
    NEW.atualizado_em = CURRENT_TIMESTAMP,
    NEW.criado_em = COALESCE(OLD.criado_em, COALESCE(OLD.atualizado_em, CURRENT_TIMESTAMP));
    RETURN NEW;
END;
$function$
;

-- CRIAR TRIGGERS PARA USAREM A FUNÇÃO DE ATUALIZAÇÃO DE DATA
create trigger billing_tax_rates_federal_tax_rules_atualizado_em before
insert
    on
    billing_tax_rates.federal_tax_rules for each row execute function public.atualizar_data_atualizacao_generica();

create trigger billing_tax_rates_icms_rules_atualizado_em before
insert
    on
    billing_tax_rates.icms_rules for each row execute function public.atualizar_data_atualizacao_generica();

create trigger billing_tax_rates_product_tax_exceptions_atualizado_em before
insert
    on
    billing_tax_rates.product_tax_exceptions for each row execute function public.atualizar_data_atualizacao_generica();

create trigger billing_tax_rates_simples_nacional_rates_atualizado_em before
insert
    on
    billing_tax_rates.simples_nacional_rates for each row execute function public.atualizar_data_atualizacao_generica();

create trigger billing_tax_rates_tax_equivalence_atualizado_em before
insert
    on
    billing_tax_rates.tax_equivalence for each row execute function public.atualizar_data_atualizacao_generica();

--- COLUNAS PARA CONFIGURAR DESONERAÇÃO
-- Adicionando na regra geral
ALTER TABLE billing_tax_rates.icms_rules 
ADD COLUMN motivo_desoneracao_icms int4,
ADD COLUMN possui_desoneracao bool DEFAULT false;

ALTER TABLE billing_tax_rates.icms_rules
ADD CONSTRAINT billing_tax_rates_icms_rules_check_cst_desoneracao
CHECK (
    (cst_icms IN ('020', '030', '040', '041', '050', '070', '090')) 
    OR (motivo_desoneracao_icms IS NULL AND possui_desoneracao = false)
);

-- Adicionando na exceção por NCM
ALTER TABLE billing_tax_rates.product_tax_exceptions 
ADD COLUMN motivo_desoneracao_icms int4,
ADD COLUMN possui_desoneracao bool DEFAULT false;

ALTER TABLE billing_tax_rates.product_tax_exceptions
ADD CONSTRAINT billing_tax_rates_product_tax_exceptions_check_cst_desoneracao
CHECK (
    (cst_icms IN ('020', '030', '040', '041', '050', '070', '090')) 
    OR (motivo_desoneracao_icms IS NULL AND possui_desoneracao = false)
);

---

ALTER TABLE billing_tax_rates.product_tax_exceptions 
ADD COLUMN IF NOT EXISTS motivo_desoneracao_icms varchar(15) NULL,
ADD COLUMN IF NOT EXISTS possui_desoneracao bool DEFAULT false;
COMMENT ON COLUMN billing_tax_rates.product_tax_exceptions.motivo_desoneracao_icms IS 'Refere-se aos códigos da SEFAZ (ex: 3-Uso na agropecuária, 9-Outros, 7-SUFRAMA). Sem isso, sua calculadora não saberá qual código gerar no XML da nota';

ALTER TABLE billing_tax_rates.icms_rules 
ADD COLUMN IF NOT EXISTS motivo_desoneracao_icms varchar(15) NULL,
ADD COLUMN IF NOT EXISTS possui_desoneracao bool DEFAULT false;
COMMENT ON COLUMN billing_tax_rates.icms_rules.motivo_desoneracao_icms IS 'Refere-se aos códigos da SEFAZ (ex: 3-Uso na agropecuária, 9-Outros, 7-SUFRAMA). Sem isso, sua calculadora não saberá qual código gerar no XML da nota';

ALTER TABLE billing_tax_rates.product_tax_exceptions ADD COLUMN IF NOT EXISTS regime_tributario_destino varchar(15) NULL;
COMMENT ON COLUMN billing_tax_rates.product_tax_exceptions.regime_tributario_destino IS 'Para o ICMS-ST, um detalhe crítico é que o MVA pode mudar dependendo do Regime Tributário do destinatário em alguns estados';

ALTER TABLE billing_tax_rates.simples_nacional_rates ADD COLUMN IF NOT EXISTS aliquota_repasse_credito numeric(5, 4) NULL;
COMMENT ON COLUMN billing_tax_rates.simples_nacional_rates.aliquota_repasse_credito IS 'Empresas do Simples que vendem para o Regime Normal podem permitir que o comprador aproveite crédito de ICMS. Esse valor é fixo por faixa e anexo, e deve aparecer na nota';

ALTER TABLE billing_tax_rates.tax_equivalence ADD COLUMN IF NOT EXISTS tipo_operacao_fiscal varchar(15) NULL;
COMMENT ON COLUMN billing_tax_rates.tax_equivalence.tipo_operacao IS 'A equivalência de código pode mudar se a nota for de entrada ou saída.';

---

ALTER TABLE billing_tax_rates.icms_rules 
ADD COLUMN IF NOT EXISTS motivo_desoneracao_icms int4,
ADD COLUMN IF NOT EXISTS possui_desoneracao bool DEFAULT false;
-- Adicionando na exceção por NCM
ALTER TABLE billing_tax_rates.product_tax_exceptions 
ADD COLUMN IF NOT EXISTS motivo_desoneracao_icms int4,
ADD COLUMN IF NOT EXISTS possui_desoneracao bool DEFAULT false;

---

-- Refinamento para permitir regras específicas de origem/destino para o mesmo NCM, considerando o regime tributário do destinatário
ALTER TABLE billing_tax_rates.product_tax_exceptions 
ADD COLUMN IF NOT EXISTS uf_origem bpchar(2) NOT NULL DEFAULT '**'; -- '**' pode representar 'Todas' se desejar

-- Criar um índice único para garantir que não haja sobreposição de regras para o mesmo cenário
CREATE UNIQUE INDEX idx_ncm_origem_destino_validade 
ON billing_tax_rates.product_tax_exceptions (ncm, uf_origem, uf_destino, inicio_validade)
WHERE final_validade IS NULL;

--

-- PRIMEIRO RASCUNHO DA TABELA DE REFORMA TRIBUTARIA
CREATE TABLE billing_tax_rates.reforma_tributaria_rules (
    id serial4 PRIMARY KEY,
    ncm varchar(10) NOT NULL,
    tipo_aliquota varchar(20) DEFAULT 'padrao', -- padrao, reduzida_60, reduzida_100 (isento), seletivo
    aliquota_cbs numeric(5, 2), -- Federal
    aliquota_ibs_estadual numeric(5, 2),
    aliquota_ibs_municipal numeric(5, 2),
    aliquota_is numeric(5, 2) DEFAULT 0.00, -- Imposto Seletivo ("Imposto do Pecado")
    inicio_validade date NOT NULL,
    final_validade date,
    -- Campos para o Cash Forward (crédito imediato) se necessário
    permite_credito_amplo bool DEFAULT true
);


--

-- product_tax_exceptions: Adicione uma coluna codigo_beneficio_fiscal (cBenef no XML). 
-- Muitos estados exigem esse código específico vinculado ao motivo da desoneração para validar a nota.
ALTER TABLE billing_tax_rates.product_tax_exceptions 
ADD COLUMN IF NOT EXISTS desoneracao_codigo_beneficio_fiscal varchar(20) NULL;
COMMENT ON COLUMN billing_tax_rates.product_tax_exceptions.desoneracao_codigo_beneficio_fiscal IS 'Código específico exigido por alguns estados para validar a desoneração na nota fiscal. Geralmente está vinculado ao motivo da desoneração e é essencial para garantir que a nota seja aceita sem erros de validação. Ex: Para desoneração por uso na agropecuária, o código pode ser "3-AGROPECUARIA". Sem esse código, mesmo que o motivo da desoneração esteja correto, a nota pode ser rejeitada.';

---

--- Tabela dua DUAL que atenda ao CBS e IBS com as mesmas regras de validade e regime tributário
-- PREMISSA:
-- 1) TODO NCM DEVE ESTAR CONFIGURADO NESSA TABELA PARA SE GERAR O CALCULO DO IMPOSTO
-- 2) AS COLUNAS NCM, UF_DESTINO SERÃO OBRIGATORIAS
-- 3) AS COLUNAS aliquota_cbs, aliquota_ibs_estadual, aliquota_ibs_municipal SERÃO OBRIGATORIAS
-- 4) TODO UF_DESTINO TERÁ UM LINHA, SEM INDICAÇÃO DE MUNICIO_DESTINO (CODIGO IBGE) PARA INDICAR A QUE NESSA LINHA ESTÃO CONFIGURADAS "aliquota_cbs PADRÃO", "aliquota_ibs_estadual PADRÃO", "aliquota_ibs_municipal PADRÃO para os municipios que seguiram a definição estadual".
-- 5) QUANDO UM MUNICIPIO TIVER UMA REGRA DIFERENTE EXISTIRÁ UMA LINHA ESPECIFICA PARA O MUNICIPIO, ONDE DEVERÁ SER INDICADAS TODAS AS ALIQUOTAS, inclusive percentual_reducao, is_imposto_seletivo, aliquota_is

CREATE TABLE billing_tax_rates.iva_dual_rules (
    id serial4 PRIMARY KEY,
    ncm varchar(10) NOT NULL,
    uf_destino bpchar(2) NOT NULL, -- O imposto é 100% no destino
    municipio_destino_ibge bpchar(7) NULL, -- Código IBGE do município de destino
    aliquota_cbs numeric(5, 2) NOT NULL, -- Federal (unificada)
    aliquota_ibs_estadual numeric(5, 2) NOT NULL,
    aliquota_ibs_municipal numeric(5, 2) NOT NULL,
    percentual_reducao numeric(5, 2) DEFAULT 0.00, -- Para os casos de 60% ou 100% de redução
    is_imposto_seletivo bool DEFAULT false, -- Se incide o "Imposto do Pecado"
    aliquota_is numeric(5, 2) DEFAULT 0.00,
    inicio_validade date NOT NULL,
    final_validade date,
);

CREATE UNIQUE INDEX idx_iva_rules_lookup 
ON billing_tax_rates.iva_dual_rules (ncm, uf_destino, COALESCE(municipio_destino_ibge, '0000000'));


--- IDEIA ABAIXO DE CONSULTA:
SELECT * FROM billing_tax_rates.iva_dual_rules
WHERE ncm = '12345678' 
  AND uf_destino = 'SP'
  AND (municipio_destino_ibge = '3550308' OR municipio_destino_ibge IS NULL)
ORDER BY municipio_destino_ibge DESC -- Garante que a regra específica venha antes da nula
LIMIT 1;

-- Tabela de LOG auditoria
CREATE TABLE billing_tax_rates.iva_dual_rules_log (
    log_id serial4 PRIMARY KEY,
    operation_type char(1) NOT NULL, -- 'I' (Insert), 'U' (Update), 'D' (Delete)
    changed_by varchar(100),
    changed_at timestamp DEFAULT CURRENT_TIMESTAMP,
    
    -- Réplica das colunas originais
    original_id int4,
    ncm varchar(10),
    uf_destino bpchar(2),
    municipio_destino_ibge bpchar(7),
    aliquota_cbs numeric(5, 2),
    aliquota_ibs_estadual numeric(5, 2),
    aliquota_ibs_municipal numeric(5, 2),
    percentual_reducao numeric(5, 2),
    is_imposto_seletivo bool,
    aliquota_is numeric(5, 2),
    inicio_validade date,
    final_validade date
);

-- trigger de auditoria para a tabela iva_dual_rules
CREATE OR REPLACE FUNCTION billing_tax_rates.fn_log_iva_dual_rules()
RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'DELETE') THEN
        INSERT INTO billing_tax_rates.iva_dual_rules_log (operation_type, original_id, ncm, uf_destino, municipio_destino_ibge, changed_by)
        VALUES ('D', OLD.id, OLD.ncm, OLD.uf_destino, OLD.municipio_destino_ibge, current_user);
        RETURN OLD;
    ELSIF (TG_OP = 'UPDATE') THEN
        INSERT INTO billing_tax_rates.iva_dual_rules_log (operation_type, original_id, ncm, uf_destino, municipio_destino_ibge, aliquota_cbs, aliquota_ibs_estadual, aliquota_ibs_municipal, percentual_reducao, is_imposto_seletivo, aliquota_is, inicio_validade, final_validade, changed_by)
        VALUES ('U', OLD.id, OLD.ncm, OLD.uf_destino, OLD.municipio_destino_ibge, OLD.aliquota_cbs, OLD.aliquota_ibs_estadual, OLD.aliquota_ibs_municipal, OLD.percentual_reducao, OLD.is_imposto_seletivo, OLD.aliquota_is, OLD.inicio_validade, OLD.final_validade, current_user);
        RETURN NEW;
    ELSIF (TG_OP = 'INSERT') THEN
        INSERT INTO billing_tax_rates.iva_dual_rules_log (operation_type, original_id, ncm, uf_destino, municipio_destino_ibge, aliquota_cbs, changed_by)
        VALUES ('I', NEW.id, NEW.ncm, NEW.uf_destino, NEW.municipio_destino_ibge, NEW.aliquota_cbs, current_user);
        RETURN NEW;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_audit_iva_dual_rules
AFTER INSERT OR UPDATE OR DELETE ON billing_tax_rates.iva_dual_rules
FOR EACH ROW EXECUTE FUNCTION billing_tax_rates.fn_log_iva_dual_rules();

---



-- ///////// RASCUNHO PARA REFAZER A TABELA DE REGRAS DE ICMS //////////////////
-- ///////// E ASSIM TERMOS SOMENTE UMA TABELA QUE SERVIRÁ PARA A REGRAS E PARA AS EXCEÇÕES A REGRA

-- CREATE TABLE billing_tax_rates.product_tax_exceptions (
-- 	id serial4 NOT NULL,

-- 	uf_origem bpchar(2) DEFAULT NOT NULL,
-- 	uf_destino bpchar(2) DEFAULT uf_origem NOT NULL,
-- 	ncm varchar(10) '*'::bpchar NOT NULL,
-- 	natureza_operacao bpchar(2) '*'::bpchar NOT NULL,
-- 	perfil_comprador bpchar(15) '*'::bpchar NOT NULL,

-- 	aliquota_interestadual numeric(5, 2) NULL,
-- 	aliquota_uf_destino numeric(5, 2) NOT NULL,
-- 	cst_pis bpchar(3) NULL,
-- 	cst_cofins bpchar(3) NULL,
-- 	cst_icms bpchar(3) NULL,
-- 	reducao_base numeric(5, 2) DEFAULT 0.00 NULL,
-- 	percentual_fcp numeric(5, 2) DEFAULT 0.00 NULL,
-- 	mva_st numeric(5, 2) NULL,
-- 	possui_protocolo_st bool DEFAULT false NULL,
	
-- 	motivo_desoneracao_icms varchar(15) NULL,
-- 	regime_tributario_destino varchar(15) NULL,
-- 	possui_desoneracao bool DEFAULT false NULL,
-- 	desoneracao_codigo_beneficio_fiscal varchar(20) NULL,
	
-- 	aliquota_pis_unitario numeric(10, 4) NULL,
-- 	aliquota_cofins_unitario numeric(10, 4) NULL,

-- 	// aqui abaixo estão os codigos a serem usados para preencher os documentos oficiais
-- 	csN bpchar(3) NULL,
-- 	csosn bpchar(3) NULL,
-- 	cest varchar(7) NULL,

-- 	// campos de auditoria
-- 	inicio_validade date DEFAULT CURRENT_DATE NOT NULL,
-- 	final_validade date NULL,
-- 	criado_em timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
-- 	atualizado_em timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,



-- FUTURA TABELA PARA REGRAS DE IPI
-- billing_tax_rates.ipi_regras definition
-- Drop table
-- DROP TABLE billing_tax_rates.ipi_regras;
CREATE TABLE billing_tax_rates.ipi_regras (
	id serial4 NOT NULL,
	ncm varchar(8) DEFAULT '*'::character varying NOT NULL,
	ex_ipi varchar(3) DEFAULT '*'::character varying NOT NULL,
	crt_emitente varchar(15) DEFAULT '*'::character varying NOT NULL,
	tipo_operacao_fiscal varchar(15) DEFAULT '*'::character varying NOT NULL,
	perfil_comprador varchar(15) DEFAULT '*'::character varying NOT NULL,
	uf_destino bpchar(2) DEFAULT '*'::bpchar NOT NULL,
	zona_especial varchar(15) DEFAULT '*'::character varying NOT NULL,
	aliquota_ipi numeric(15, 4) DEFAULT 0 NULL,
	valor_pauta_ipi numeric(15, 4) DEFAULT 0 NULL,
	cst_ipi bpchar(2) NOT NULL,
	c_enq bpchar(3) DEFAULT '999'::bpchar NOT NULL,
	possui_desoneracao bool DEFAULT false NULL,
	motivo_desoneracao varchar(15) NULL,
	desoneracao_codigo_beneficio_fiscal varchar(15) NULL,
	motivo_desoneracao_ipi varchar(15) NULL,
    inicio_validade date NOT NULL DEFAULT CURRENT_DATE,
    final_validade date NULL,
	criado_em timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	atualizado_em timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT ipi_regras_pkey PRIMARY KEY (id),
	CONSTRAINT ipi_regras_ukey UNIQUE (ncm, ex_ipi, crt_emitente, tipo_operacao_fiscal, perfil_comprador, zona_especial, uf_destino, inicio_validade),
	CONSTRAINT check_datas CHECK (final_validade IS NULL OR inicio_validade <= final_validade)
);

CREATE INDEX idx_ipi_busca_regra ON billing_tax_rates.ipi_regras 
(ncm, ex_ipi, crt_emitente, tipo_operacao_fiscal,perfil_comprador, zona_especial, uf_destino, inicio_validade);

-- ============================================================================
-- TABELA ncm_seletivo — NCMs sujeitos ao Imposto Seletivo (IS)
-- Ref: PROCEDURE-FIN-00001 SOP-003, RULES-CATALOG-FIN-00001 BR-TAX-INF-005
-- ============================================================================
CREATE TABLE billing_tax_rates.ncm_seletivo (
    id serial4 PRIMARY KEY,
    ncm varchar(10) NOT NULL,
    categoria varchar(100) NOT NULL,         -- ex: 'BEBIDAS_ALCOOLICAS', 'CIGARROS', 'REFRIGERANTES'
    aliquota_is numeric(5, 2) NOT NULL,      -- Aliquota do IS para esta categoria
    descricao text,
    inicio_validade date NOT NULL DEFAULT CURRENT_DATE,
    final_validade date NULL,
    criado_em timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    atualizado_em timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_ncm_seletivo_lookup
ON billing_tax_rates.ncm_seletivo (ncm, inicio_validade)
WHERE final_validade IS NULL;

COMMENT ON TABLE billing_tax_rates.ncm_seletivo IS 'NCMs sujeitos ao Imposto Seletivo (IS) — "Imposto do Pecado". Gerido pelo Ministerio da Fazenda.';

-- Dados de exemplo para ncm_seletivo (baseados nas categorias previstas na EC 132/2023)
INSERT INTO billing_tax_rates.ncm_seletivo (ncm, categoria, aliquota_is, descricao) VALUES
    ('22030000', 'BEBIDAS_ALCOOLICAS', 50.00, 'Cervejas de malte — categoria bebidas alcoólicas'),
    ('22041010', 'BEBIDAS_ALCOOLICAS', 50.00, 'Vinhos espumantes — categoria bebidas alcoólicas'),
    ('22082000', 'BEBIDAS_ALCOOLICAS', 75.00, 'Aguardentes e destilados — categoria bebidas alcoólicas'),
    ('24022000', 'CIGARROS', 100.00, 'Cigarros com filtro — categoria tabaco'),
    ('21069010', 'REFRIGERANTES', 25.00, 'Xaropes e concentrados para refrigerantes'),
    ('22021000', 'REFRIGERANTES', 25.00, 'Águas minerais adicionadas de açúcar/edulcorantes — refrigerantes');

-- ============================================================================
-- TABELA cbs_rates — Alíquotas CBS por classe tributária
-- Ref: PROCEDURE-FIN-00001 SOP-001
-- ============================================================================
CREATE TABLE billing_tax_rates.cbs_rates (
    id serial4 PRIMARY KEY,
    c_class_trib varchar(50) NOT NULL,       -- Classe tributaria CBS (ex: 'TELECOM', 'GERAL', 'SAUDE')
    aliquota_cbs numeric(5, 2) NOT NULL,     -- Aliquota padrao da classe
    descricao text,
    inicio_validade date NOT NULL DEFAULT CURRENT_DATE,
    final_validade date NULL,
    criado_em timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    atualizado_em timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_cbs_rates_lookup
ON billing_tax_rates.cbs_rates (c_class_trib, inicio_validade)
WHERE final_validade IS NULL;

COMMENT ON TABLE billing_tax_rates.cbs_rates IS 'Aliquotas CBS (Contribuicao sobre Bens e Servicos) por classe tributaria. Ref: LC 214/2025.';

-- ============================================================================
-- TABELA iss_rates — Alíquotas ISS por município
-- Ref: PROCEDURE-FIN-00001 SOP-010
-- ============================================================================
CREATE TABLE billing_tax_rates.iss_rates (
    id serial4 PRIMARY KEY,
    codigo_ibge varchar(7) NOT NULL,         -- Codigo IBGE do municipio (7 digitos)
    municipio_nome varchar(200) NOT NULL,
    uf bpchar(2) NOT NULL,
    aliquota_iss numeric(5, 2) NOT NULL,     -- Aliquota padrao do ISS no municipio
    item_lista_servico varchar(5),            -- Item da Lista de Servicos da LC 116/2003 (ex: '1.05' para telecom)
    descricao text,
    inicio_validade date NOT NULL DEFAULT CURRENT_DATE,
    final_validade date NULL,
    criado_em timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    atualizado_em timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_iss_rates_lookup
ON billing_tax_rates.iss_rates (codigo_ibge, COALESCE(item_lista_servico, ''), inicio_validade)
WHERE final_validade IS NULL;

COMMENT ON TABLE billing_tax_rates.iss_rates IS 'Aliquotas ISS por municipio. Range constitucional: [2%, 5%] conforme LC 116/2003 art. 8o-A.';

-- Dados de exemplo para iss_rates
INSERT INTO billing_tax_rates.iss_rates (codigo_ibge, municipio_nome, uf, aliquota_iss, descricao) VALUES
    ('3550308', 'Sao Paulo', 'SP', 5.00, 'Aliquota maxima padrao — conforme Lei Municipal 13.701/2003'),
    ('3304557', 'Rio de Janeiro', 'RJ', 5.00, 'Aliquota maxima padrao — conforme Lei Municipal 3.720/2004'),
    ('5300108', 'Brasilia', 'DF', 5.00, 'Aliquota maxima padrao — conforme Lei Complementar Distrital'),
    ('3106200', 'Belo Horizonte', 'MG', 5.00, 'Aliquota maxima padrao'),
    ('4106902', 'Curitiba', 'PR', 5.00, 'Aliquota maxima padrao');

-- - QUERY DE BUSCA - Lógica de "Match" (Busca da Regra)
SELECT * FROM billing_tax_rates.ipi_regras
WHERE (ncm = $1 OR ncm = '*')
  AND (ex_ipi = $2 OR ex_ipi = '*')
  AND (crt_emitente = $3 OR crt_emitente = '*')
  AND (tipo_operacao_fiscal = $4 OR tipo_operacao_fiscal = '*')
  AND (perfil_comprador = $5 OR perfil_comprador = '*')
  AND (uf_destino = $6 OR uf_destino = '*')
  AND (zona_especial = $7 OR zona_especial = '*')
Filtro de Validade
  AND ($8 >= inicio_validade)
  AND ($8 <= final_validade OR final_validade IS NULL)
ORDER BY 
  (ncm != '*') DESC, 
  (ex_ipi != '*') DESC, 
  (crt_emitente != '*') DESC, 
  (tipo_operacao_fiscal != '*') DESC,
  (perfil_comprador != '*') DESC,
  (uf_destino != '*') DESC,
  (zona_especial != '*') DESC,
  inicio_validade DESC -- Pega a regra mais recente em caso de empate
LIMIT 1;

-- ============================================================================
-- 11. Tabela de Tokens Fiscais (GAP-002 — BR-06 Garantia de Preço Ofertado)
-- ============================================================================
-- Armazena tokens de congelamento de alíquotas gerados pelo endpoint
-- POST /v1/token/generate. Cada token snapshot as alíquotas CBS, IBS
-- (estadual + municipal) e IS para uma tupla (NCM, UF destino, IBGE)
-- e é válido por TAX_TOKEN_TTL_MINUTES (default 60 min).
CREATE TABLE billing_tax_rates.tax_tokens (
    id UUID PRIMARY KEY,
    ncm VARCHAR(10) NOT NULL,
    uf_origem CHAR(2) NOT NULL,
    uf_destino CHAR(2) NOT NULL,
    municipio_ibge CHAR(7),
    aliquota_cbs NUMERIC(5,2) NOT NULL,
    aliquota_ibs_estadual NUMERIC(5,2) NOT NULL,
    aliquota_ibs_municipal NUMERIC(5,2) NOT NULL,
    aliquota_is NUMERIC(5,2) DEFAULT 0.00,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Índice para limpeza de tokens expirados (background job)
CREATE INDEX idx_tax_tokens_expires_at ON billing_tax_rates.tax_tokens (expires_at);

-- Índice para lookup rápido por ID (GET /v1/token/{id})
CREATE INDEX idx_tax_tokens_id ON billing_tax_rates.tax_tokens (id);

-- ============================================================================
-- 12. Tabela de Fornecedores Fiscais (GAP-007 — BR-08 Qualificação)
-- ============================================================================
CREATE TABLE billing_tax_rates.fornecedor_fiscal (
    cnpj CHAR(14) PRIMARY KEY,
    regime_tributario VARCHAR(50) NOT NULL,
    certificado_regularidade BOOLEAN DEFAULT false,
    permite_credito BOOLEAN DEFAULT false,
    data_qualificacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_validade TIMESTAMP NOT NULL,
    status VARCHAR(30) NOT NULL
);

CREATE INDEX idx_fornecedor_status ON billing_tax_rates.fornecedor_fiscal (status);

-- ============================================================================
-- Tabela: cst_reforma
-- Propósito: Tabela oficial de CST (Código de Situação Tributária) para CBS/IBS
-- Fonte: RFB — LC 214/2025 (CST_cClassTrib_2025-10-03_Public_verde.xlsx)
-- Feature: FEATURE-2026-06-30-DT03-CST-REFORMA-TRIBUTARIA
-- Registros: 164 CCTs, 18 CSTs
-- ============================================================================

CREATE TABLE IF NOT EXISTS billing_tax_rates.cst_reforma (
    id                    SERIAL PRIMARY KEY,
    cst                   CHAR(3) NOT NULL,
    cct                   CHAR(6) NOT NULL UNIQUE,
    descricao_cst         TEXT NOT NULL,
    descricao_cct         TEXT NOT NULL,
    exige_tributacao      BOOLEAN DEFAULT TRUE,
    reducao_bc            BOOLEAN DEFAULT FALSE,
    reducao_aliquota      BOOLEAN DEFAULT FALSE,
    transferencia_credito BOOLEAN DEFAULT FALSE,
    diferimento           BOOLEAN DEFAULT FALSE,
    monofasica            BOOLEAN DEFAULT FALSE,
    credito_presumido     BOOLEAN DEFAULT FALSE,
    ajuste_competencia    BOOLEAN DEFAULT FALSE,
    percentual_reducao_ibs DECIMAL(5,2) DEFAULT 0,
    percentual_reducao_cbs DECIMAL(5,2) DEFAULT 0,
    tipo_aliquota         VARCHAR(50),
    url_legislacao        TEXT,
    simples_nacional      VARCHAR(100),
    created_at            TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_cst_reforma_cst ON billing_tax_rates.cst_reforma(cst);
CREATE INDEX IF NOT EXISTS idx_cst_reforma_cct ON billing_tax_rates.cst_reforma(cct);

-- ============================================================================
-- Dados: 164 CCTs (Classificações Tributárias) oficiais
-- ============================================================================

INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('000', '000001', 'Tributação integral', 'Situações tributadas integralmente pelo IBS e CBS.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art4', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('000', '000002', 'Tributação integral', 'Exploração de via, observado o art. 11 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art11', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('000', '000003', 'Tributação integral', 'Regime automotivo - projetos incentivados, observado o art. 311 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art311', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('000', '000004', 'Tributação integral', 'Regime automotivo - projetos incentivados, observado o art. 312 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art312', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('000', '000005', 'Tributação integral', 'Operação com EAC destinado à mistura com gasolina A, mas com saída do biocombustível com destinação diversa, observado o art. 179 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art179', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('010', '010001', 'Tributação com alíquotas uniformes', 'Operações do FGTS não realizadas pela Caixa Econômica Federal, observado o art. 212 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '5 - Uniforme Setorial', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art212', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('010', '010002', 'Tributação com alíquotas uniformes', 'Operações do serviço financeiro', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '5 - Uniforme Setorial', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art233', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('011', '011001', 'Tributação com alíquotas uniformes reduzidas', 'Planos de assistência funerária, observado o art. 236 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '4 - Uniforme Nacional', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art236', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('011', '011002', 'Tributação com alíquotas uniformes reduzidas', 'Planos de assistência à saúde, observado o art. 237 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '4 - Uniforme Nacional', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art237', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('011', '011003', 'Tributação com alíquotas uniformes reduzidas', 'Intermediação de planos de assistência à saúde, observado o art. 240 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '4 - Uniforme Nacional', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art240', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('011', '011004', 'Tributação com alíquotas uniformes reduzidas', 'Concursos e prognósticos, observado o art. 246 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '4 - Uniforme Nacional', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art246', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('011', '011005', 'Tributação com alíquotas uniformes reduzidas', 'Planos de assistência à saúde de animais domésticos, observado o art. 243 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 30, 30, '4 - Uniforme Nacional', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art243', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200001', 'Alíquota reduzida', 'Serviços de transporte de bens até as zonas de processamento de exportação e bens exportados a partir das zonas de processamento de exportação, observado o art. 103 da Lei Complementar n 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art103', '3 - Receita Bruta Exportação Direta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200002', 'Alíquota reduzida', 'Fornecimento ou importação de tratores, máquinas e implementos agrícolas, destinados a produtor rural não contribuinte, e de veículos de transporte de carga destinados a transportador autônomo de carga pessoa física não contribuinte, observado o art. 110 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art110', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200003', 'Alíquota reduzida', 'Vendas de produtos destinados à alimentação humana relacionados no Anexo I da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NCM/SH, que compõem a Cesta Básica Nacional de Alimentos, criada nos termos do art. 8º da Emenda Constitucional nº 132, de 20 de dezembro de 2023, observado o art. 125 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art125', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200004', 'Alíquota reduzida', 'Fornecimento de dispositivos médicos com a especificação das respectivas classificações da NCM/SH previstas no Anexo XII da Lei Complementar nº 214, de 2025, observado o art. 144 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art144', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200005', 'Alíquota reduzida', 'Fornecimento de dispositivos médicos com a especificação das respectivas classificações da NCM/SH previstas no Anexo IV da Lei Complementar nº 214, de 2025, quando adquiridos por órgãos da administração pública direta, autarquias, fundações públicas e entidades de saúde imunes, observado o art. 144 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art144', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200006', 'Alíquota reduzida', 'Situação de emergência de saúde pública reconhecida pelo Poder Legislativo federal, estadual, distrital ou municipal competente, ato conjunto do Ministro da Fazenda e do Comitê Gestor do IBS poderá ser editado, a qualquer momento, para incluir dispositivos não listados no Anexo XII da Lei Complementar nº 214, de 2025, limitada a vigência do benefício ao período e à localidade da emergência de saúde pública, observado o art. 144 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art144', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200007', 'Alíquota reduzida', 'Fornecimento dos dispositivos de acessibilidade próprios para pessoas com deficiência relacionados no Anexo XIII da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NCM/SH, observado o art. 145 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art145', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200008', 'Alíquota reduzida', 'Fornecimento dos dispositivos de acessibilidade próprios para pessoas com deficiência relacionados no Anexo V da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NCM/SH, quando adquiridos por órgãos da administração pública direta, autarquias, fundações públicas e entidades imunes, observado o art. 145 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art145', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200009', 'Alíquota reduzida', 'Fornecimento dos medicamentos registrados na Anvisa, observado o art. 146 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art146', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200010', 'Alíquota reduzida', 'Fornecimento dos medicamentos registrados na Anvisa, quando adquiridos por órgãos da administração pública direta, autarquias, fundações públicas e entidades imunes, observado o art. 146 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art146', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200011', 'Alíquota reduzida', 'Fornecimento das composições para nutrição enteral e parenteral, composições especiais e fórmulas nutricionais destinadas às pessoas com erros inatos do metabolismo relacionadas no Anexo VI da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NCM/SH, quando adquiridas por órgãos da administração pública direta, autarquias e fundações públicas, observado o art. 146 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art146', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200012', 'Alíquota reduzida', 'Situação de emergência de saúde pública reconhecida pelo Poder Legislativo federal, estadual, distrital ou municipal competente, ato conjunto do Ministro da Fazenda e do Comitê Gestor do IBS poderá ser editado, a qualquer momento, limitada a vigência do benefício ao período e à localidade da emergência de saúde pública, observado o art. 146 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art146', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200013', 'Alíquota reduzida', 'Fornecimento de tampões higiênicos, absorventes higiênicos internos ou externos, descartáveis ou reutilizáveis, calcinhas absorventes e coletores menstruais, observado o art. 147 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art147', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200014', 'Alíquota reduzida', 'Fornecimento dos produtos hortícolas, frutas e ovos, relacionados no Anexo XV da Lei Complementar nº 214 , de 2025, com a especificação das respectivas classificações da NCM/SH e desde que não cozidos, observado o art. 148 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art148', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200015', 'Alíquota reduzida', 'Venda de automóveis de passageiros de fabricação nacional de, no mínimo, 4 (quatro) portas, inclusive a de acesso ao bagageiro, quando adquiridos por motoristas profissionais que exerçam, comprovadamente, em automóvel de sua propriedade, atividade de condutor autônomo de passageiros, na condição de titular de autorização, permissão ou concessão do poder público, e que destinem o automóvel à utilização na categoria de aluguel (táxi), ou por pessoas com deficiência física, visual, auditiva, deficiência mental severa ou profunda, transtorno do espectro autista, com prejuízos na comunicação social e em padrões restritos ou repetitivos de comportamento de nível moderado ou grave, nos termos da legislação relativa à matéria, observado o disposto no art. 149 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art149', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200016', 'Alíquota reduzida', 'Prestação de serviços de pesquisa e desenvolvimento por Instituição Científica, Tecnológica e de Inovação (ICT) sem fins lucrativos para a administração pública direta, autarquias e fundações públicas ou para o contribuinte sujeito ao regime regular do IBS e da CBS, observado o disposto no art. 156  da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art156', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200017', 'Alíquota reduzida', 'Operações relacionadas ao FGTS, considerando aquelas necessárias à aplicação da Lei nº 8.036, de 1990, realizadas pelo Conselho Curador ou Secretaria Executiva do FGTS, observado o art. 212 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art212', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200018', 'Alíquota reduzida', 'Operações de resseguro e retrocessão ficam sujeitas à incidência à alíquota zero, inclusive quando os prêmios de resseguro e retrocessão forem cedidos ao exterior, observado o art. 223 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art223', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200019', 'Alíquota reduzida', 'Importador dos serviços financeiros que seja contribuinte e tenha direito de apropriação de créditos na aquisição do mesmo serviço financeiro no País, observado o art. 231 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art231', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200020', 'Alíquota reduzida', 'Operação praticada por sociedades cooperativas optantes por regime específico do IBS e CBS, quando o associado destinar bem ou serviço à cooperativa de que participa, e a cooperativa fornecer bem ou serviço ao associado sujeito ao regime regular do IBS e da CBS, observado o art. 271 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art271', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200021', 'Alíquota reduzida', 'Serviços de transporte público coletivo de passageiros ferroviário e hidroviário urbanos, semiurbanos e metropolitanos, observado o art. 285 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art285', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200022', 'Alíquota reduzida', 'Operação originada fora da Zona Franca de Manaus que destine bem material industrializado de origem nacional a contribuinte estabelecido na Zona Franca de Manaus que seja habilitado nos termos do art. 442 da Lei Complementar nº 214, de 2025, e sujeito ao regime regular do IBS e da CBS ou optante pelo regime do Simples Nacional de que trata o art. 12 da Lei Complementar nº 123, de 2006, observado o art. 445 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art445', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200023', 'Alíquota reduzida', 'Operação realizada por indústria incentivada que destine bem material intermediário para outra indústria incentivada na Zona Franca de Manaus, desde que a entrega ou disponibilização dos bens ocorra dentro da referida área, observado o art. 448 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art448', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200024', 'Alíquota reduzida', 'Operação originada fora das Áreas de Livre Comércio que destine bem material industrializado de origem nacional a contribuinte estabelecido nas Áreas de Livre Comércio que seja habilitado nos termos do art. 456 da Lei Complementar nº 214, de 2025, e sujeito ao regime regular do IBS e da CBS ou optante pelo regime do Simples Nacional de que trata o art. 12 da Lei Complementar nº 123, de 2006, observado o art. 463 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art463', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200025', 'Alíquota reduzida', 'Fornecimento dos serviços de educação relacionados ao Programa Universidade para Todos (Prouni), instituído pela Lei nº 11.096, de 13 de janeiro de 2005, observado o art. 308 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art308', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200026', 'Alíquota reduzida', 'Locação de imóveis localizados nas zonas reabilitadas, pelo prazo de 5 (cinco) anos, contado da data de expedição do habite-se, e relacionados a projetos de reabilitação urbana de zonas históricas e de áreas críticas de recuperação e reconversão urbanística dos Municípios ou do Distrito Federal, a serem delimitadas por lei municipal ou distrital, observado o art. 158 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 80, 80, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art158', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200027', 'Alíquota reduzida', 'Operações de locação, cessão onerosa e arrendamento de bens imóveis, observado o art. 261 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 70, 70, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art261', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200028', 'Alíquota reduzida', 'Fornecimento dos serviços de educação relacionados no Anexo II da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da Nomenclatura Brasileira de Serviços, Intangíveis e Outras Operações que Produzam Variações no Patrimônio (NBS), observado o art. 129 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art129', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200029', 'Alíquota reduzida', 'Fornecimento dos serviços de saúde humana relacionados no Anexo III da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NBS, observado o art. 130 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art130', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200030', 'Alíquota reduzida', 'Venda dos dispositivos médicos relacionados no Anexo IV da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NCM/SH, observado o art. 131 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art131', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200031', 'Alíquota reduzida', 'Fornecimento dos dispositivos de acessibilidade próprios para pessoas com deficiência relacionados no Anexo V da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NCM/SH, observado o art. 132 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art132', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200032', 'Alíquota reduzida', 'Fornecimento dos medicamentos registrados na Anvisa ou produzidos por farmácias de manipulação, ressalvados os medicamentos sujeitos à alíquota zero de que trata o art. 146 da Lei Complementar nº 214, de 2025, observado o art. 133 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art133', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200033', 'Alíquota reduzida', 'Fornecimento das composições para nutrição enteral e parenteral, composições especiais e fórmulas nutricionais destinadas às pessoas com erros inatos do metabolismo relacionadas no Anexo VI da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NCM/SH, observado o art. 133 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art133', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200034', 'Alíquota reduzida', 'Fornecimento dos alimentos destinados ao consumo humano relacionados no Anexo VII da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NCM/SH, observado o art. 135 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art135', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200035', 'Alíquota reduzida', 'Fornecimento dos produtos de higiene pessoal e limpeza relacionados no Anexo VIII da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NCM/SH, observado o art. 136 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art136', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200036', 'Alíquota reduzida', 'Fornecimento de produtos agropecuários, aquícolas, pesqueiros, florestais e extrativistas vegetais in natura, observado o art. 137 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art137', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200037', 'Alíquota reduzida', 'Fornecimento de serviços ambientais de conservação ou recuperação da vegetação nativa, mesmo que fornecidos sob a forma de manejo sustentável de sistemas agrícolas, agroflorestais e agrossilvopastoris, em conformidade com as definições e requisitos da legislação específica, observado o art. 137 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art137', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200038', 'Alíquota reduzida', 'Fornecimento dos insumos agropecuários e aquícolas relacionados no Anexo IX da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NCM/SH e da NBS, observado o art. 138 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art138', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200039', 'Alíquota reduzida', 'Fornecimento dos bens e serviços listados no Anexo X da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NCM/SH e NBS, nos casos relacionados com produções nacionais artísticas, culturais, de eventos, jornalísticas e audiovisuais, observado o art. 139 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art139', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200040', 'Alíquota reduzida', 'Fornec dos seguintes serv de comunic instit à admin púb direta, autarq e fund púb: serviços direcionados ao planej, criação, programação e manutenção de páginas eletrônicas da admin pública, ao monitor e gestão de suas redes sociais e à otimização de páginas e canais digitais para mecanismos de buscas e produção de mensagens, infográficos, painéis interativos e conteúdo institucional, serviços de relações com a imprensa, que reúnem estrat org para promover e reforçar a comunicação dos órgãos e das entidades contratantes com seus públicos de interesse, por meio da interação com prof da imprensa, e serviços de relações públicas, que compreendem o esforço de comunic planej, coeso e contínuo que tem por obj estab adequada percepção da atuação e dos obj instituc, a partir do estímulo à compreensão mútua e da manut de padrões de relac e fluxos de inf entre os órgãos e as entidades contrat e seus públicos de interesse, no País e no exterior, obs o art. 140 da Lei Compl nº 214, de 2025', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art140', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200041', 'Alíquota reduzida', 'Operações relacionadas às seguintes atividades desportivas: fornecimento de serviço de educação desportiva, classificado no código 1.2205.12.00 da NBS, e gestão e exploração do desporto por associações e clubes esportivos filiados ao órgão estadual ou federal responsável pela coordenação dos desportos, inclusive por meio de venda de ingressos para eventos desportivos, fornecimento oneroso ou não de bens e serviços, inclusive ingressos, por meio de programas de sócio-torcedor, cessão dos direitos desportivos dos atletas e transferência de atletas para outra entidade desportiva ou seu retorno à atividade em outra entidade desportiva, observado o art. 141 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art141', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200042', 'Alíquota reduzida', 'Operações relacionadas às seguintes atividades desportivas: gestão e exploração do desporto por associações e clubes esportivos filiados ao órgão estadual ou federal responsável pela coordenação dos desportos, observado o art. 141 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art141', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200043', 'Alíquota reduzida', 'Fornecimento à administração pública direta, autarquias e fundações púbicas dos serviços e dos bens relativos à soberania e à segurança nacional, à segurança da informação e à segurança cibernética relacionados no Anexo XI da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NBS e da NCM/SH, observado o art. 142 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art142', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200044', 'Alíquota reduzida', 'Operações e prestações de serviços de segurança da informação e segurança cibernética desenvolvidos por sociedade que tenha sócio brasileiro com o mínimo de 20% (vinte por cento) do seu capital social, relacionados no Anexo XI da Lei Complementar nº 214, de 2025, com a especificação das respectivas classificações da NBS e da NCM/SH, observado o art. 142 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art142', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200045', 'Alíquota reduzida', 'Operações relacionadas a projetos de reabilitação urbana de zonas históricas e de áreas críticas de recuperação e reconversão urbanística dos Municípios ou do Distrito Federal, a serem delimitadas por lei municipal ou distrital, observado o art. 158 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art158', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200046', 'Alíquota reduzida', 'Operações com bens imóveis, observado o art. 261 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 50, 50, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art261', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200047', 'Alíquota reduzida', 'Bares e Restaurantes, observado o art. 275 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 40, 40, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art275', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200048', 'Alíquota reduzida', 'Hotelaria, Parques de Diversão e Parques Temáticos, observado o art. 281 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 40, 40, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art281', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200049', 'Alíquota reduzida', 'Transporte coletivo de passageiros rodoviário, ferroviário e hidroviário intermunicipais e interestaduais, observado o art. 286 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 40, 40, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art286', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200050', 'Alíquota reduzida', 'Serviços de transporte aéreo regional coletivo de passageiros ou de carga, observado o art. 287 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 40, 40, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art287', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200051', 'Alíquota reduzida', 'Agências de Turismo, observado o art. 289 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 40, 40, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art289', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200052', 'Alíquota reduzida', 'Prestação de serviços das seguintes profissões intelectuais de natureza científica, literária ou artística, submetidas à fiscalização por conselho profissional: administradores, advogados, arquitetos e urbanistas, assistentes sociais, bibliotecários, biólogos, contabilistas, economistas, economistas domésticos, profissionais de educação física, engenheiros e agrônomos, estatísticos, médicos veterinários e zootecnistas, museólogos, químicos, profissionais de relações públicas, técnicos industriais e técnicos agrícolas, observado o art. 127 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 30, 30, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art127', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200053', 'Alíquota reduzida', 'Fornecimento de medicamentos registrados na Anvisa, quando  classificados como soros ou vacinas, observado o art. 146 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art146', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('200', '200054', 'Alíquota reduzida', 'Fornecimento de bem material pela cooperativa de produção agropecuária a associado não sujeito ao regime regular do IBS e da CBS com anulação de créditos referentes ao bem fornecido, observado o art. 271 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, 100, 100, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art271', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('220', '220001', 'Alíquota fixa', 'Incorporação imobiliária submetida ao regime especial de tributação, observado o art. 485 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '1 - Fixa', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art485', '-');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('220', '220002', 'Alíquota fixa', 'Incorporação imobiliária submetida ao regime especial de tributação, observado o art. 485 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '1 - Fixa', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art485', '-');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('220', '220003', 'Alíquota fixa', 'Alienação de imóvel decorrente de parcelamento do solo, observado o art. 486 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '1 - Fixa', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art486', '-');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('221', '221001', 'Alíquota fixa proporcional', 'Locação, cessão onerosa ou arrendamento de bem imóvel com alíquota sobre a receita bruta, observado o art. 487 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '1 - Fixa', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art484', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('221', '221002', 'Alíquota fixa proporcional', 'Incorporação imobiliária submetida ao regime especial de tributação, observado o art. 485 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '1 - Fixa', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art485', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('221', '221003', 'Alíquota fixa proporcional', 'Incorporação imobiliária submetida ao regime especial de tributação, observado o art. 485 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '1 - Fixa', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art485', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('221', '221004', 'Alíquota fixa proporcional', 'Alienação de imóvel decorrente de parcelamento do solo, observado o art. 486 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '1 - Fixa', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art486', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('222', '222001', 'Redução de Base de Cálculo', 'Transporte internacional de passageiros, caso os trechos de ida e volta sejam vendidos em conjunto, a base de cálculo será a metade do valor cobrado, observado o Art. 12 § 8º da Lei Complementar nº 214, de 2025.', TRUE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art12', '5 - Receita Bruta Mercado Interno/Exportação');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('400', '400001', 'Isenção', 'Fornecimento de serviços de transporte público coletivo de passageiros rodoviário e metroviário de caráter urbano, semiurbano e metropolitano, sob regime de autorização, permissão ou concessão pública, observado o art. 157 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art157', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('400', '400002', 'Isenção', 'Fornecimento de serviços de transporte público coletivo de passageiros rodoviário e metroviário de caráter urbano, semiurbano e metropolitano, sob regime de autorização, permissão ou concessão pública, com medição por quilômetro rodado, observado o art. 157 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art157', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410001', 'Imunidade e não incidência', 'Fornecimento de bonificações quando constem do respectivo documento fiscal e que não dependam de evento posterior, observado o art. 5º da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art5', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410002', 'Imunidade e não incidência', 'Transferências entre estabelecimentos pertencentes ao mesmo contribuinte, observado o art. 6º da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art6', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410003', 'Imunidade e não incidência', 'Doações que não tenham por objeto bens ou serviços que tenham permitido a apropriação de créditos pelo doador, observado o art. 6º da Lei Complementar nº 214, de 2025.
', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art6', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410004', 'Imunidade e não incidência', 'Exportações de bens e serviços, observado o art. 8º da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art8', '3 - Receita Bruta Exportação Direta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410005', 'Imunidade e não incidência', 'Fornecimentos realizados pela União, pelos Estados, pelo Distrito Federal e pelos Municípios, observado o art. 9º da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art9', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410006', 'Imunidade e não incidência', 'Fornecimentos realizados por entidades religiosas e templos de qualquer culto, inclusive suas organizações assistenciais e beneficentes, observado o art. 9º da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art9', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410007', 'Imunidade e não incidência', 'Fornecimentos realizados por partidos políticos, inclusive suas fundações, entidades sindicais dos trabalhadores e instituições de educação e de assistência social, sem fins lucrativos, observado o art. 9º da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art9', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410008', 'Imunidade e não incidência', 'Fornecimentos de livros, jornais, periódicos e do papel destinado a sua impressão, observado o art. 9º da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art9', '2 - Receita Bruta Interna sem Cálculo IBS/CBS');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410009', 'Imunidade e não incidência', 'Fornecimentos de fonogramas e videofonogramas musicais produzidos no Brasil contendo obras musicais ou literomusicais de autores brasileiros e/ou obras em geral interpretadas por artistas brasileiros, bem como os suportes materiais ou arquivos digitais que os contenham, salvo na etapa de replicação industrial de mídias ópticas de leitura a laser, observado o art. 9º da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art9', '2 - Receita Bruta Interna sem Cálculo IBS/CBS');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410010', 'Imunidade e não incidência', 'Fornecimentos de serviço de comunicação nas modalidades de radiodifusão sonora e de sons e imagens de recepção livre e gratuita, observado o art. 9º da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art9', '2 - Receita Bruta Interna sem Cálculo IBS/CBS');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410011', 'Imunidade e não incidência', 'Fornecimentos de ouro, quando definido em lei como ativo financeiro ou instrumento cambial, observado o art. 9º da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art9', '2 - Receita Bruta Interna sem Cálculo IBS/CBS');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410012', 'Imunidade e não incidência', 'Fornecimento de condomínio edilício não optante pelo regime regular, observado o art. 26 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art26', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410013', 'Imunidade e não incidência', 'Exportações de combustíveis, observado o art. 98 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art98', '3 - Receita Bruta Exportação Direta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410014', 'Imunidade e não incidência', 'Fornecimento de produtor rural não contribuinte, observado o art. 164 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art164', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410015', 'Imunidade e não incidência', 'Fornecimento por transportador autônomo não contribuinte, observado o art. 169 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art169', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410016', 'Imunidade e não incidência', 'Fornecimento ou aquisição de resíduos sólidos, observado o art. 170 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art170', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410017', 'Imunidade e não incidência', 'Aquisição de bem móvel com crédito presumido sob condição de revenda realizada, observado o art. 171 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art171', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410018', 'Imunidade e não incidência', 'Operações relacionadas aos fundos garantidores e executores de políticas públicas, inclusive de habitação, previstos em lei, assim entendidas os serviços prestados ao fundo pelo seu agente operador e por entidade encarregada da sua administração, observado o art. 213 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art213', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410019', 'Imunidade e não incidência', 'Exclusão da gorjeta na base de cálculo no fornecimento de alimentação, observado o art. 274 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art274', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410020', 'Imunidade e não incidência', 'Exclusão do valor de intermediação na base de cálculo no fornecimento de alimentação, observado o art. 274 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art274', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410021', 'Imunidade e não incidência', 'Contribuição de que trata o art. 149-A da Constituição Federal, observado o art. 12 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art12', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410022', 'Imunidade e não incidência', 'Consolidação da propriedade pelo credor de bens móveis ou imóveis que tenham sido objeto de garantia, observado o art. 200 da Lei Complementar nº 214, de 2025.
', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art200', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410023', 'Imunidade e não incidência', 'Alienação de bens móveis ou imóveis que tenham sido objeto de garantia constituída em favor de credor em que o prestador da garantia não seja contribuinte, observado o art. 200 da Lei Complementar nº 214, de 2025.
', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art200', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410024', 'Imunidade e não incidência', 'Consolidação da propriedade pelo grupo de consórcio de bem que tenha sido objeto de garantia, observado o art. 204 da Lei Complementar nº 214, de 2025.
', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art204', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410025', 'Imunidade e não incidência', 'Alienação de bem que tenha sido objeto de garantia constituída em favor do grupo de consórcio em que o prestador da garantia não seja contribuinte, observado o art. 204 da Lei Complementar nº 214, de 2025.
', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art204', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410026', 'Imunidade e não incidência', 'Doações sem contraprestação em benefício do doador, com anulação de crédito apropriados pelo doador referente ao fornecimento doado, observado o art. 6º da Lei Complementar nº 214, de 2025.
', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art6', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410027', 'Imunidade e não incidência', 'Fornecimento de bens e serviços, desde que vinculados direta e exclusivamente à exportação de bens materiais ou associados à entrega no exterior de bens materiais, observado o art. 6º da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art80', '3 - Receita Bruta Exportação Direta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410028', 'Imunidade e não incidência', 'Operações com bens imóveis realizadas por pessoas físicas não consideradas contribuintes do regime regular do IBS e da CBS, observado o art. 251 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art251', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410029', 'Imunidade e não incidência', 'Operações não sujeitas à incidência de IBS e de CBS, alcançadas apenas por obrigação acessória do ICMS, observado o art. 4º da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art4', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410030', 'Imunidade e não incidência', 'Estorno de crédito apropriado de bens adquiridos e venham a perecer, deteriorar-se ou ser objeto de roubo, furto ou extravio, observado o art. 47 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art47', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410031', 'Imunidade e não incidência', 'Fornecimento em período anterior ao início de vigência de incidências de CBS e IBS, observado o art. 544 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art544', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410032', 'Imunidade e não incidência', 'Tributos incidentes na operação que não integram a base de cálculo do IBS e da CBS, observado o art. 12 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art12', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410033', 'Imunidade e não incidência', 'Operações com bens imóveis, inclusive operações com direitos reais sobre bens imóveis, realizadas por Fundos de Investimento Imobiliário (FII) e Fundos de Investimento nas Cadeias Produtivas do Agronegócio (Fiagro), observado o art. 26 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art26', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410034', 'Imunidade e não incidência', 'Fundos de investimento cujo patrimônio seja constituído exclusivamente por aplicações em participações societárias, certificados, direitos, títulos, valores mobiliários e demais ativos financeiros permitidos pela Comissão de Valores Mobiliários, observado o art. 26 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art26', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410035', 'Imunidade e não incidência', 'Fornecimento realizado por nanoempreendedor, observado o art. 26 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art26', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410036', 'Imunidade e não incidência', 'Descontos incondicionais, observado o art. 12 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art12', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410037', 'Imunidade e não incidência', 'Importação de bens materiais sem incidência de IBS e CBS, observado o art. 66 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art66', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('410', '410999', 'Imunidade e não incidência', 'Operações não onerosas sem previsão de tributação, não especificadas anteriormente, observado o art. 4º da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art4', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('510', '510001', 'Diferimento', 'Operações, sujeitas a diferimento, com energia elétrica ou com direitos a ela relacionados, relativas à importação, geração, comercialização, distribuição e transmissão, observado o art. 28 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, TRUE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art28', '2 - Receita Bruta Interna sem Cálculo IBS/CBS');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('515', '515001', 'Diferimento com redução de alíquota', 'Operações, sujeitas a diferimento, com insumos agropecuários e aquícolas, observado o art. 138 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, 60, 60, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art138', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550001', 'Suspensão', 'Exportações de bens materiais, observado o art. 82 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art82', '4 - Receita Bruta Exportação Indireta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550002', 'Suspensão', 'Regime de Trânsito, observado o art. 84 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art84', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550003', 'Suspensão', 'Regimes de Depósito, observado o art. 85 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art85', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550004', 'Suspensão', 'Regimes de Depósito, observado o art. 87 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art87', '4 - Receita Bruta Exportação Indireta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550005', 'Suspensão', 'Regimes de Depósito, observado o art. 87 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art87', '4 - Receita Bruta Exportação Indireta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550006', 'Suspensão', 'Regimes de Permanência Temporária, observado o art. 88 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art88', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550007', 'Suspensão', 'Regimes de Aperfeiçoamento, observado o art. 90 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art90', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550008', 'Suspensão', 'Importação de bens para o Regime de Repetro-Temporário, de que tratam o inciso I do art. 93 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art93', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550009', 'Suspensão', 'GNL-Temporário, de que trata o inciso II do art. 93 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art93', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550010', 'Suspensão', 'Repetro-Permanente, de que trata o inciso III do art. 93 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art93', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550011', 'Suspensão', 'Repetro-Industrialização, de que trata o inciso IV do art. 93 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art93', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550012', 'Suspensão', 'Repetro-Nacional, de que trata o inciso V do art. 93 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art93', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550013', 'Suspensão', 'Repetro-Entreposto, de que trata o inciso VI do art. 93 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art93', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550014', 'Suspensão', 'Zona de Processamento de Exportação, observado os arts. 99, 100 e 102 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art99', '4 - Receita Bruta Exportação Indireta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550015', 'Suspensão', 'Regime Tributário para Incentivo à Modernização e à Ampliação da Estrutura Portuária - Reporto, observado o art. 105 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art105', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550016', 'Suspensão', 'Regime Especial de Incentivos para o Desenvolvimento da Infraestrutura - Reidi, observado o art. 106 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art106', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550017', 'Suspensão', 'Regime Tributário para Incentivo à Atividade Econômica Naval – Renaval, observado o art. 107 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art107', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550018', 'Suspensão', 'Desoneração da aquisição de bens de capital, observado o art. 109 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art109', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550019', 'Suspensão', 'Importação de bem material por indústria incentivada para utilização na Zona Franca de Manaus, observado o art. 443 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art443', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550020', 'Suspensão', 'Áreas de livre comércio, observado o art. 461 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art461', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550021', 'Suspensão', 'Fornecimento de produtos agropecuários in natura para contribuinte do regime regular que promova industrialização destinada a exportação, observado o art. 82 da Lei Complementar nº 214, de 2025.
', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art82', '4 - Receita Bruta Exportação Indireta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550022', 'Suspensão', 'Regime Especial de Incentivos para a Produção de Hidrogênio de Baixa Emissão de Carbono (Rehidro),  observado o art. 106 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art106', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550023', 'Suspensão', 'Operações com hidrocarbonetos líquidos derivados de petróleo não combustíveis ou de gás natural, inclusive nafta, observado o art. 172 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art172', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550024', 'Suspensão', 'Importações e nas aquisições no mercado interno de máquinas, equipamentos e veículos destinados a utilização nas atividades de que trata o inciso IIIdo art. 107 efetuadas para incorporação a seu ativo imobilizado, observado o art. 107 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art107', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('550', '550025', 'Suspensão', 'Importações e nas aquisições no mercado interno de matérias-primas, produtos intermediários, partes, peças e componentes para utilização na construção, conservação, modernização e reparo de embarcações pré-registradas ou registradas no REB, observado o art. 107 da Lei Complementar nº 214, de 2025.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art107', '1 - Receita Bruta Interna');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('620', '620001', 'Tributação Monofásica', 'Tributação monofásica sobre combustíveis, observados os art. 172 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, TRUE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art172', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('620', '620002', 'Tributação Monofásica', 'Tributação monofásica com responsabilidade pela retenção sobre combustíveis, observado o art. 178 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, TRUE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art178', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('620', '620003', 'Tributação Monofásica', 'Tributação monofásica com responsabilidade de retenção de tributos por terceiros, observado o art. 178 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, TRUE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art178', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('620', '620004', 'Tributação Monofásica', 'Tributação monofásica sobre mistura de EAC com gasolina A em percentual superior ou inferior ao obrigatório, observado o art. 179 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, TRUE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art179', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('620', '620005', 'Tributação Monofásica', 'Tributação monofásica sobre mistura de EAC com gasolina A em percentual superior ou inferior ao obrigatório, observado o art. 179 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, TRUE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art179', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('620', '620006', 'Tributação Monofásica', 'Tributação monofásica sobre combustíveis cobrada anteriormente, observador o art. 180 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, TRUE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art180', '2 - Receita Bruta Interna sem Cálculo IBS/CBS');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('620', '620007', 'Tributação Monofásica', 'Perecimento, deteriorização, roubo, furto ou extravio no regime monofásico sem estorno de crédito, observado o art. 47 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, TRUE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art47', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('800', '800001', 'Transferência de crédito', 'Fusão, cisão ou incorporação, observado o art. 55 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art55', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('800', '800002', 'Transferência de crédito', 'Transferência de crédito do associado, inclusive as cooperativas singulares, para cooperativa de que participa das operações antecedentes às operações em que fornece bens e serviços e os créditos presumidos, observado o art. 272 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art272', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('810', '810001', 'Ajuste de IBS na ZFM', 'Crédito presumido sobre o valor apurado nos fornecimentos a partir da Zona Franca de Manaus, observado o art. 450 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, TRUE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art450', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('811', '811001', 'Ajustes', 'Anulação de crédito proporcional ao valor das operações imunes e isentas, observado o art. 51 da Lei Complementar nº 214, de 2025.
', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, TRUE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art51', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('811', '811002', 'Ajustes', 'Débitos de notas fiscais não processadas na apuração, observado o art. 45 da Lei Complementar nº 214, de 2025.
', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, TRUE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art45', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('811', '811003', 'Ajustes', 'Débitos apurados após o desenquadramento do regime Simples Nacional, observado o art. 41 da Lei Complementar nº 214, de 2025.
', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, TRUE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art41', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('820', '820001', 'Tributação em documento específico', 'Documento com informações de fornecimento de serviços de planos de assistência à saúde elencados no art. 234 da Lei Complementar nº 214, de 2025, mas com tributação realizada por outro meio', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art234', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('820', '820002', 'Tributação em documento específico', 'Documento com informações de fornecimento de serviços de planos de assinstência funerária, mas com tributação realizada por outro meio, observado o art. 236 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art236', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('820', '820003', 'Tributação em documento específico', 'Documento com informações de fornecimento de serviços de planos de assinstência à saúde de animais domésticos, mas com tributação realizada por outro meio, observado o art. 243 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art243', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('820', '820004', 'Tributação em documento específico', 'Documento com informações de prestação de serviços de consursos de prognósticos, mas com tributação realizada por outro meio, observado o art. 248 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art248', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('820', '820005', 'Tributação em documento específico', 'Documento com informações de alienação de bens imóveis, mas com tributação realizada por outro meio,, observado o art. 254 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art254', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('820', '820006', 'Tributação em documento específico', 'Documento com informações de fornecimento de serviços de exploração de via, mas com tributação realizada por outro meio, observado o art. 11 da Lei Complementar nº 214, de 2025.
', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art11', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('820', '820007', 'Tributação em documento específico', 'Documento com informações de fornecimento de serviços financeiros, mas com tributação realizada por outro meio, observado o art. 181 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art181', '9 - Fornecimento Incompatível com SN');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('820', '820008', 'Tributação em documento específico', 'Documento com informações de fornecimento de serviço continuado, mas com tributação realizada em fatura anterior, observado o art. 10 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art10', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('820', '820009', 'Tributação em documento específico', 'Cobrança relativa a fornecimentos declarados em outro documento, observado o art. 60 da Lei Complementar nº 214, de 2025.', FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '3 - Sem aliquota', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art60', '0 - Não Receita Bruta');
INSERT INTO billing_tax_rates.cst_reforma (cst, cct, descricao_cst, descricao_cct, exige_tributacao, reducao_bc, reducao_aliquota, transferencia_credito, diferimento, monofasica, credito_presumido, ajuste_competencia, percentual_reducao_ibs, percentual_reducao_cbs, tipo_aliquota, url_legislacao, simples_nacional) VALUES ('830', '830001', 'Exclusão da Base de Cálculo', 'Documento com  exclusão da base de cálculo da CBS e do IBS refrente à energia elétrica fornecida pela distribuidora à unidade consumidora, conforme  Art 28, parágrafos 3° e 4°.', TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, 0, 0, '2 - Padrão', 'https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp214.htm#art28', '0 - Não Receita Bruta');
