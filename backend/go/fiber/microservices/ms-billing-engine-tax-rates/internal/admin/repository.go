// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/admin/repository.go
package admin

import (
	"context"
	"fmt"

	"taxnexus-billing-core-lib/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

// AdminRepository define operações administrativas sobre alíquotas.
type AdminRepository interface {
	// UpsertIvaDualRule fecha a regra existente (se houver) e insere uma nova.
	// Retorna a nova regra inserida.
	UpsertIvaDualRule(ctx context.Context, input IvaDualRuleInput, changedBy string) (*IvaDualRuleOutput, error)

	// ListIvaDualRules lista regras com filtros opcionais.
	ListIvaDualRules(ctx context.Context, filter ListRulesFilter) ([]IvaDualRuleOutput, error)
}

var _ AdminRepository = (*PostgresAdminRepository)(nil)

// PostgresAdminRepository implementa AdminRepository usando PostgreSQL.
type PostgresAdminRepository struct {
	pool *pgxpool.Pool
}

// NewPostgresAdminRepository cria um novo repositório admin.
func NewPostgresAdminRepository(pool *pgxpool.Pool) *PostgresAdminRepository {
	return &PostgresAdminRepository{pool: pool}
}

func (r *PostgresAdminRepository) UpsertIvaDualRule(ctx context.Context, input IvaDualRuleInput, changedBy string) (*IvaDualRuleOutput, error) {
	// 1. Fecha regra existente (seta final_validade = hoje)
	_, err := r.pool.Exec(ctx,
		`UPDATE billing_tax_rates.iva_dual_rules
		 SET final_validade = CURRENT_DATE
		 WHERE ncm = $1 AND uf_destino = $2
		   AND COALESCE(municipio_destino_ibge, '') = $3
		   AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)`,
		input.NCM, input.UFDestino, input.MunicipioDestinoIBGE,
	)
	if err != nil {
		return nil, fmt.Errorf("falha ao fechar regra existente: %w", err)
	}

	// 2. Insere nova regra
	var id int64
	row := r.pool.QueryRow(ctx,
		`INSERT INTO billing_tax_rates.iva_dual_rules
		 (ncm, uf_destino, municipio_destino_ibge, aliquota_cbs,
		  aliquota_ibs_estadual, aliquota_ibs_municipal,
		  percentual_reducao, is_imposto_seletivo, aliquota_is,
		  inicio_validade, final_validade)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		 RETURNING id`,
		input.NCM, input.UFDestino, nullIfEmpty(input.MunicipioDestinoIBGE),
		input.AliquotaCBS, input.AliquotaIBSEstadual, input.AliquotaIBSMunicipal,
		input.PercentualReducao, input.IsImpostoSeletivo, input.AliquotaIS,
		input.InicioValidade, input.FinalValidade,
	)
	if err := row.Scan(&id); err != nil {
		return nil, fmt.Errorf("falha ao inserir nova regra: %w", err)
	}

	// 3. Registra auditoria manual (trigger captura a operação, aqui registramos changed_by)
	_, _ = r.pool.Exec(ctx,
		`INSERT INTO billing_tax_rates.iva_dual_rules_log
		 (operation_type, original_id, ncm, uf_destino, municipio_destino_ibge,
		  aliquota_cbs, aliquota_ibs_estadual, aliquota_ibs_municipal,
		  percentual_reducao, is_imposto_seletivo, aliquota_is,
		  inicio_validade, final_validade, changed_by)
		 VALUES ('I', $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
		id, input.NCM, input.UFDestino, nullIfEmpty(input.MunicipioDestinoIBGE),
		input.AliquotaCBS, input.AliquotaIBSEstadual, input.AliquotaIBSMunicipal,
		input.PercentualReducao, input.IsImpostoSeletivo, input.AliquotaIS,
		input.InicioValidade, input.FinalValidade, changedBy,
	)

	out := &IvaDualRuleOutput{
		ID:                   id,
		NCM:                  input.NCM,
		UFDestino:            input.UFDestino,
		MunicipioDestinoIBGE: input.MunicipioDestinoIBGE,
		AliquotaCBS:          input.AliquotaCBS,
		AliquotaIBSEstadual:  input.AliquotaIBSEstadual,
		AliquotaIBSMunicipal: input.AliquotaIBSMunicipal,
		PercentualReducao:    input.PercentualReducao,
		IsImpostoSeletivo:    input.IsImpostoSeletivo,
		AliquotaIS:           input.AliquotaIS,
		InicioValidade:       input.InicioValidade,
		FinalValidade:        input.FinalValidade,
	}
	return out, nil
}

func (r *PostgresAdminRepository) ListIvaDualRules(ctx context.Context, filter ListRulesFilter) ([]IvaDualRuleOutput, error) {
	query := `SELECT id, ncm, uf_destino, COALESCE(municipio_destino_ibge,''),
	          aliquota_cbs, aliquota_ibs_estadual, aliquota_ibs_municipal,
	          percentual_reducao, is_imposto_seletivo, aliquota_is,
	          inicio_validade, final_validade
	          FROM billing_tax_rates.iva_dual_rules WHERE 1=1`
	args := []interface{}{}
	argIdx := 1

	if filter.NCM != "" {
		query += fmt.Sprintf(" AND ncm = $%d", argIdx)
		args = append(args, filter.NCM)
		argIdx++
	}
	if filter.UF != "" {
		query += fmt.Sprintf(" AND uf_destino = $%d", argIdx)
		args = append(args, filter.UF)
		argIdx++
	}
	if filter.AtivasApenas {
		query += " AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)"
	}
	query += " ORDER BY inicio_validade DESC LIMIT 100"

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("falha ao listar regras: %w", err)
	}
	defer rows.Close()

	var results []IvaDualRuleOutput
	for rows.Next() {
		var o IvaDualRuleOutput
		if err := rows.Scan(&o.ID, &o.NCM, &o.UFDestino, &o.MunicipioDestinoIBGE,
			&o.AliquotaCBS, &o.AliquotaIBSEstadual, &o.AliquotaIBSMunicipal,
			&o.PercentualReducao, &o.IsImpostoSeletivo, &o.AliquotaIS,
			&o.InicioValidade, &o.FinalValidade); err != nil {
			return nil, fmt.Errorf("falha ao ler regra: %w", err)
		}
		results = append(results, o)
	}
	return results, rows.Err()
}

func nullIfEmpty(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}

// adminRepositoryAdapter adapta AdminRepository ao TaxRepository para DI compatível.
// Usado para que AdminTaxService possa acessar GetIvaDualRule sem duplicação.
type adminRepositoryAdapter struct {
	repo repository.TaxRepository
}

func (a *adminRepositoryAdapter) GetIvaDualRule(ctx context.Context, ncm, uf, ibge string) (*repository.IvaDualRule, error) {
	return a.repo.GetIvaDualRule(ctx, ncm, uf, ibge)
}
