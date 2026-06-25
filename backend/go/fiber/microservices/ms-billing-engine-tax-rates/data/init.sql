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
