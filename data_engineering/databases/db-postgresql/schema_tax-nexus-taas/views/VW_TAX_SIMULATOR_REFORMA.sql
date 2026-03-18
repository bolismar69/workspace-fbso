CREATE OR REPLACE VIEW tax_nexus_taas.VW_TAX_SIMULATOR_REFORMA AS
WITH tax_base AS (
    SELECT 
        m.nome_municipio,
        m.codigo_ibge,
        e.sigla AS uf,
        c.codigo_ncm,
        SUM(CASE WHEN c.sigla_imposto = 'CBS' THEN c.aliquota ELSE 0 END) as aliq_cbs,
        SUM(CASE WHEN c.sigla_imposto = 'IBS_EST' THEN c.aliquota ELSE 0 END) as aliq_ibs_est,
        SUM(CASE WHEN c.sigla_imposto = 'IBS_MUN' THEN c.aliquota ELSE 0 END) as aliq_ibs_mun
    FROM tax_nexus_taas.TB_MUNICIPIO m
    JOIN tax_nexus_taas.TB_ESTADO e ON m.id_uf = e.id_uf
    JOIN tax_nexus_taas.TB_TAX_ENGINE_CACHE c ON m.codigo_ibge = c.codigo_ibge
    GROUP BY m.nome_municipio, m.codigo_ibge, e.sigla, c.codigo_ncm
)
SELECT 
    t.*,
    mt.ano_competencia,
    (t.aliq_cbs * mt.perc_transicao) as cbs_calculada,
    ((t.aliq_ibs_est + t.aliq_ibs_mun) * mt.perc_transicao) as ibs_calculado
FROM tax_base t
CROSS JOIN tax_nexus_taas.TB_MARCO_TEMPORAL_REGRA mt;
