// path: backend/go/libs/go-native/taxnexus-billing-core-lib/repository/postgres_repository.go
package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type postgresTaxRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTaxRepository(db *pgxpool.Pool) TaxRepository {
	return &postgresTaxRepository{db: db}
}

func parseDataOperacaoIPI(dataOperacao string) (time.Time, error) {
	if dataOperacao == "" {
		return time.Now(), nil
	}

	layouts := []string{time.RFC3339, "2006-01-02"}
	for _, layout := range layouts {
		parsed, err := time.Parse(layout, dataOperacao)
		if err == nil {
			return parsed, nil
		}
	}

	return time.Time{}, fmt.Errorf("data_operacao invalida para IPI: %s", dataOperacao)
}

func (r *postgresTaxRepository) GetFederalTaxRule(ctx context.Context, regimeTributario, cstPIS, cstCOFINS string) (*FederalTaxRule, error) {
	query := `
		SELECT id,
		       regime_tributario,
		       cst_pis,
		       cst_cofins,
		       aliquota_pis,
		       aliquota_cofins,
		       COALESCE(exclui_icms_base, true),
		       inicio_validade,
		       final_validade,
		       criado_em,
		       atualizado_em
		FROM billing_tax_rates.federal_tax_rules
		WHERE regime_tributario = $1
		  AND cst_pis = $2
		  AND cst_cofins = $3
		  AND inicio_validade <= CURRENT_DATE
		  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
		ORDER BY inicio_validade DESC, id DESC
		LIMIT 1`

	var rule FederalTaxRule
	err := r.db.QueryRow(ctx, query, regimeTributario, cstPIS, cstCOFINS).Scan(
		&rule.Id,
		&rule.RegimeTributario,
		&rule.CSTPIS,
		&rule.CSTCOFINS,
		&rule.AliquotaPIS,
		&rule.AliquotaCOFINS,
		&rule.ExcluiICMSBase,
		&rule.InicioValidade,
		&rule.FinalValidade,
		&rule.CriadoEm,
		&rule.AtualizadoEm,
	)
	if err != nil {
		return nil, err
	}

	return &rule, nil
}

func (r *postgresTaxRepository) GetICMSRule(ctx context.Context, orig, dest string) (*ICMSRule, error) {
	query := `
		SELECT id,
		       uf_origem,
		       uf_destino,
		       aliquota_interna,
		       aliquota_interestadual,
		       COALESCE(cst_padrao, '00'),
		       COALESCE(reducao_base, 0.00),
		       COALESCE(percentual_fcp, 0.00),
		       COALESCE(mva_padrao, 0.00),
		       COALESCE(possui_protocolo_st, false),
		       inicio_validade,
		       final_validade,
		       criado_em,
		       atualizado_em
		FROM billing_tax_rates.icms_rules
		WHERE uf_origem = $1
		  AND uf_destino = $2
		  AND inicio_validade <= CURRENT_DATE
		  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
		ORDER BY inicio_validade DESC, id DESC
		LIMIT 1`
	var rule ICMSRule
	err := r.db.QueryRow(ctx, query, orig, dest).Scan(
		&rule.Id,
		&rule.UFOrigem,
		&rule.UFDestino,
		&rule.AliquotaInterna,
		&rule.AliquotaInterestadual,
		&rule.CSTPadrao,
		&rule.ReducaoBase,
		&rule.PercentualFCP,
		&rule.MVAPadrao,
		&rule.PossuiProtocoloST,
		&rule.InicioValidade,
		&rule.FinalValidade,
		&rule.CriadoEm,
		&rule.AtualizadoEm,
	)
	if err != nil {
		return nil, err
	}

	return &rule, nil
}

func (r *postgresTaxRepository) GetEquivalence(ctx context.Context, CSOSN string, tipoOperacao string) (*TaxEquivalence, error) {
	query := `
		SELECT id,
		       csosn,
		       cst_equivalente,
		       COALESCE(permite_credito, false),
		       COALESCE(descricao, ''),
		       inicio_validade,
		       final_validade,
		       criado_em,
		       atualizado_em
		FROM billing_tax_rates.tax_equivalence
		WHERE csosn = $1
		  AND (tipo_operacao_fiscal = $2 OR tipo_operacao_fiscal IS NULL)
		  AND inicio_validade <= CURRENT_DATE
		  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
		ORDER BY inicio_validade DESC, id DESC
		LIMIT 1`
	var eq TaxEquivalence
	err := r.db.QueryRow(ctx, query, CSOSN, tipoOperacao).Scan(
		&eq.Id,
		&eq.CSOSN,
		&eq.CSTEquivalente,
		&eq.PermiteCredito,
		&eq.Descricao,
		&eq.InicioValidade,
		&eq.FinalValidade,
		&eq.CriadoEm,
		&eq.AtualizadoEm,
	)
	return &eq, err
}

// Implementando os métodos que faltavam para satisfazer a interface
func (r *postgresTaxRepository) GetRateByNCM(ctx context.Context, ncm string, uf string) (float64, error) {
	var rate float64
	err := r.db.QueryRow(ctx, "SELECT rate FROM billing_tax_rates.tax_rates WHERE ncm = $1 AND uf = $2", ncm, uf).Scan(&rate)
	return rate, err
}

func (r *postgresTaxRepository) GetIBSRate(ctx context.Context, municipioIBGE string) (float64, error) {
	var rate float64
	err := r.db.QueryRow(ctx, "SELECT rate FROM billing_tax_rates.ibs_rates WHERE ibge_code = $1", municipioIBGE).Scan(&rate)
	return rate, err
}

func (r *postgresTaxRepository) GetSimplesFaixa(ctx context.Context, anexo string, rbt12 decimal.Decimal) (*SimplesFaixa, error) {
	query := `SELECT 
		               id,
		               anexo,
		               faixa,
		               COALESCE(faturamento_min, 0.00),
		               COALESCE(faturamento_max, 999999999999999.99),
		               aliquota_nominal,
		               valor_deduzir,
		               percentual_icms,
		               inicio_validade,
		               final_validade,
		               criado_em,
		               atualizado_em
					FROM billing_tax_rates.simples_nacional_rates
					WHERE anexo = $1
					  AND $2 >= COALESCE(faturamento_min, 0.00)
					  AND $2 <= COALESCE(faturamento_max, 999999999999999.99)
					  AND inicio_validade <= CURRENT_DATE
					  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
					ORDER BY inicio_validade DESC, faixa ASC, id DESC
					LIMIT 1;`
	var f SimplesFaixa
	err := r.db.QueryRow(ctx, query, anexo, rbt12).Scan(
		&f.Id,
		&f.Anexo,
		&f.Faixa,
		&f.FaturamentoMin,
		&f.FaturamentoMax,
		&f.AliqNominal,
		&f.ValorDeduzir,
		&f.PercIcmsAnexo,
		&f.InicioValidade,
		&f.FinalValidade,
		&f.CriadoEm,
		&f.AtualizadoEm,
	)
	return &f, err
}

func (r *postgresTaxRepository) GetProductException(ctx context.Context, ncmFull, ncmGroup, ufDestino string, regimeTributarioDestino string) (*ProductException, error) {
	query := `
        SELECT
            id,
            ncm,
            COALESCE(uf_destino, ''),
            COALESCE(cst_pis, ''),
            COALESCE(cst_cofins, ''),
            COALESCE(cst_icms, ''),
            COALESCE(aliquota_pis_unitario, 0.00),
            COALESCE(aliquota_cofins_unitario, 0.00),
            COALESCE(mva_st, 0.00),
            COALESCE(csosn, ''),
            COALESCE(cest, ''),
            COALESCE(possui_protocolo_st, false),
            COALESCE(aliquota_interna_destino, 0.00),
            COALESCE(aliquota_interestadual, 0.00),
            COALESCE(percentual_fcp, 0.00),
            COALESCE(reducao_base, 0.00),
            inicio_validade,
            final_validade,
            criado_em,
            atualizado_em
        FROM billing_tax_rates.product_tax_exceptions
        WHERE (ncm = $1 OR ncm = $2)
          AND (uf_destino = $3 OR uf_destino IS NULL)
          AND (regime_tributario_destino = $4 OR regime_tributario_destino IS NULL)
          AND inicio_validade <= CURRENT_DATE
		  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
        ORDER BY
            CASE
                WHEN ncm = $1 AND uf_destino = $3 AND regime_tributario_destino = $4 THEN 1
                WHEN ncm = $1 AND uf_destino = $3 AND regime_tributario_destino IS NULL THEN 2
                WHEN ncm = $1 AND uf_destino IS NULL AND regime_tributario_destino = $4 THEN 3
                WHEN ncm = $1 AND uf_destino IS NULL AND regime_tributario_destino IS NULL THEN 4
                WHEN ncm = $2 AND uf_destino = $3 AND regime_tributario_destino = $4 THEN 5
                WHEN ncm = $2 AND uf_destino = $3 AND regime_tributario_destino IS NULL THEN 6
                WHEN ncm = $2 AND uf_destino IS NULL AND regime_tributario_destino = $4 THEN 7
                ELSE 8
            END ASC,
            inicio_validade DESC,
            id DESC
        LIMIT 1;`

	var e ProductException
	err := r.db.QueryRow(ctx, query, ncmFull, ncmGroup, ufDestino, regimeTributarioDestino).Scan(
		&e.Id,
		&e.NCM,
		&e.UFDestino,
		&e.CSTPIS,
		&e.CSTCOFINS,
		&e.CSTICMS,
		&e.AliquotaPISUnitario,
		&e.AliquotaCOFINSUnitario,
		&e.MVAST,
		&e.CSOSN,
		&e.CEST,
		&e.PossuiProtocoloST,
		&e.AliquotaInternaDestino,
		&e.AliquotaInterestadual,
		&e.PercentualFCP,
		&e.ReducaoBase,
		&e.InicioValidade,
		&e.FinalValidade,
		&e.CriadoEm,
		&e.AtualizadoEm,
	)

	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *postgresTaxRepository) GetIPIRegra(ctx context.Context, NCM string, ExIPI string, CrtEmitente string, TipoOperacaoFiscal string, PerfilComprador string, UFDestino string, ZonaEspecial bool, DataOperacao string) (*IPIRegra, error) {
	dataOperacao, err := parseDataOperacaoIPI(DataOperacao)
	if err != nil {
		return nil, err
	}

	zonaEspecial := fmt.Sprintf("%t", ZonaEspecial)

	query := `
		SELECT * FROM billing_tax_rates.ipi_regras
		WHERE (ncm = $1 OR ncm = '*')
		  AND (ex_ipi = $2 OR ex_ipi = '*')
		  AND (crt_emitente = $3 OR crt_emitente = '*')
		  AND (tipo_operacao_fiscal = $4 OR tipo_operacao_fiscal = '*')
		  AND (perfil_comprador = $5 OR perfil_comprador = '*')
		  AND (uf_destino = $6 OR uf_destino = '*')
		  AND (zona_especial = $7 OR zona_especial = '*')
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
		  inicio_validade DESC
		LIMIT 1;`

	var regra IPIRegra
	err = r.db.QueryRow(ctx, query,
		NCM,
		ExIPI,
		CrtEmitente,
		TipoOperacaoFiscal,
		PerfilComprador,
		UFDestino,
		zonaEspecial,
		dataOperacao,
	).Scan(
		&regra.ID,
		&regra.NCM,
		&regra.ExIPI,
		&regra.CrtEmitente,
		&regra.TipoOperacaoFiscal,
		&regra.PerfilComprador,
		&regra.UFDestino,
		&regra.ZonaEspecial,
		&regra.AliquotaIPI,
		&regra.ValorPautaIPI,
		&regra.CSTIPI,
		&regra.CEnq,
		&regra.PossuiDesoneracao,
		&regra.MotivoDesoneracao,
		&regra.InicioValidade,
		&regra.FinalValidade,
		&regra.CriadoEm,
		&regra.AtualizadoEm,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &regra, nil
}

// GetNCMSeletivo consulta a tabela ncm_seletivo para verificar se um NCM
// esta sujeito ao Imposto Seletivo (IS).
// Retorna nil quando o NCM nao esta na tabela (nao incide IS).
// Ref: PROCEDURE-FIN-00001 SOP-003, RULES-CATALOG-FIN-00001 BR-TAX-INF-005
func (r *postgresTaxRepository) GetNCMSeletivo(ctx context.Context, ncm string) (*NCMSeletivoRule, error) {
	query := `
		SELECT ncm,
		       categoria,
		       aliquota_is,
		       descricao,
		       inicio_validade,
		       final_validade
		FROM billing_tax_rates.ncm_seletivo
		WHERE ncm = $1
		  AND inicio_validade <= CURRENT_DATE
		  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
		ORDER BY inicio_validade DESC
		LIMIT 1`

	var rule NCMSeletivoRule
	err := r.db.QueryRow(ctx, query, ncm).Scan(
		&rule.NCM,
		&rule.Categoria,
		&rule.AliquotaIS,
		&rule.Descricao,
		&rule.InicioValidade,
		&rule.FinalValidade,
	)
	if err == sql.ErrNoRows {
		return nil, nil // NCM nao esta na tabela — IS nao incide
	}
	if err != nil {
		return nil, err
	}

	return &rule, nil
}

func (r *postgresTaxRepository) GetIvaDualRule(ctx context.Context, ncm, ufDestino, municipioIBGE string) (*IvaDualRule, error) {
	query := `
		SELECT ncm,
		       uf_destino,
		       municipio_destino_ibge,
		       aliquota_cbs,
		       aliquota_ibs_estadual,
		       aliquota_ibs_municipal,
		       percentual_reducao,
		       COALESCE(is_imposto_seletivo, false),
		       COALESCE(aliquota_is, 0.00),
		       inicio_validade,
		       final_validade
		FROM billing_tax_rates.iva_dual_rules
		WHERE ncm = $1
		  AND uf_destino = $2
		  AND (municipio_destino_ibge = $3 OR ($3 = '' AND municipio_destino_ibge IS NULL))
		  AND inicio_validade <= CURRENT_DATE
		  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
		ORDER BY municipio_destino_ibge DESC, inicio_validade DESC
		LIMIT 1`

	var rule IvaDualRule
	err := r.db.QueryRow(ctx, query, ncm, ufDestino, municipioIBGE).Scan(
		&rule.NCM,
		&rule.UFDestino,
		&rule.MunicipioDestinoIBGE,
		&rule.AliquotaCBS,
		&rule.AliquotaIBSEstadual,
		&rule.AliquotaIBSMunicipal,
		&rule.PercentualReducao,
		&rule.IsImpostoSeletivo,
		&rule.AliquotaIS,
		&rule.InicioValidade,
		&rule.FinalValidade,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &rule, nil
}

// GetCSTReforma seleciona o CST oficial para CBS/IBS com base nos flags de contexto da operacao.
// A logica de prioridade segue a LC 214/2025:
//   - EfetivamenteIsento -> CST "800" (Sem tributacao)
//   - Monofasica -> CST "400" (Monofasica normal)
//   - Diferimento -> CST "510" (Diferimento)
//   - PercentualReducao > 0 -> CST "200" (Reducao de base de calculo)
//   - Default -> CST "000" (Tributacao integral)
func (r *postgresTaxRepository) GetCSTReforma(ctx context.Context, flags CSTFlags) (*CSTReforma, error) {
	var cst string
	switch {
	case flags.EfetivamenteIsento:
		cst = "800"
	case flags.IsMonofasico:
		cst = "400"
	case flags.IsDiferimento:
		cst = "510"
	case flags.PercentualReducao.GreaterThan(decimal.Zero):
		cst = "200"
	default:
		cst = "000"
	}

	query := `
		SELECT id, cst, cct, descricao_cst, descricao_cct,
		       exige_tributacao, reducao_bc, reducao_aliquota,
		       transferencia_credito, diferimento, monofasica,
		       credito_presumido, ajuste_competencia,
		       percentual_reducao_ibs, percentual_reducao_cbs,
		       tipo_aliquota, url_legislacao, simples_nacional
		FROM billing_tax_rates.cst_reforma
		WHERE cst = $1
		LIMIT 1`

	var rule CSTReforma
	err := r.db.QueryRow(ctx, query, cst).Scan(
		&rule.ID, &rule.CST, &rule.CCT,
		&rule.DescricaoCST, &rule.DescricaoCCT,
		&rule.ExigeTributacao, &rule.ReducaoBC, &rule.ReducaoAliquota,
		&rule.TransferenciaCredito, &rule.Diferimento,
		&rule.Monofasica, &rule.CreditoPresumido, &rule.AjusteCompetencia,
		&rule.PercentualReducaoIBS, &rule.PercentualReducaoCBS,
		&rule.TipoAliquota, &rule.UrlLegislacao, &rule.SimplesNacional,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &rule, nil
}
