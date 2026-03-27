// path: backend/go/libs/go-native/taxnexus-core-lib/repository/tax_repository.go
package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"taxnexus-core-lib/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type TaxRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewTaxRepository(db *pgxpool.Pool, rdb *redis.Client) *TaxRepository {
	return &TaxRepository{db: db, rdb: rdb}
}

// GetApplicableRule agora utiliza a lógica de busca em memória após recuperar todas as faixas
func (r *TaxRepository) GetApplicableRule(ctx context.Context, taxCode string, baseValue decimal.Decimal, refDate time.Time) (*models.TaxRule, error) {
	// Reutilizamos a lógica de buscar todas as regras do período para evitar múltiplas queries
	rules, err := r.GetTaxRulesForPeriod(ctx, taxCode, refDate)
	if err != nil {
		return nil, err
	}

	for _, rule := range rules {
		if baseValue.GreaterThanOrEqual(rule.RangeMin) {
			if rule.RangeMax == nil || baseValue.LessThanOrEqual(*rule.RangeMax) {
				return &rule, nil
			}
		}
	}
	return nil, fmt.Errorf("no applicable rule found for value %s", baseValue.String())
}

// GetConfig busca um valor de configuração específico no banco/cache
func (r *TaxRepository) GetConfig(ctx context.Context, taxCode, key string, refDate time.Time) (decimal.Decimal, error) {
    var val decimal.Decimal
    query := `
        SELECT config_value 
        FROM individual_tax_rates.tax_configs 
        WHERE tax_code = $1 AND config_key = $2 
        AND valid_from <= $3 AND (valid_to IS NULL OR valid_to >= $3)
        LIMIT 1`
    
    err := r.db.QueryRow(ctx, query, taxCode, key, refDate).Scan(&val)
    if err != nil {
        return decimal.Zero, err
    }
    return val, nil
}

// GetTableConfigs busca todas as chaves de configuração de um imposto para uma data
func (r *TaxRepository) GetTableConfigs(ctx context.Context, taxCode string, refDate time.Time) (map[string]decimal.Decimal, error) {
	cacheKey := fmt.Sprintf("tax_configs:%s:%s", taxCode, refDate.Format("2006-01-02"))
	configs := make(map[string]decimal.Decimal)

	// 1. Tentar Cache
	if val, err := r.rdb.Get(ctx, cacheKey).Result(); err == nil && val != "" {
		if err := json.Unmarshal([]byte(val), &configs); err == nil {
			return configs, nil
		}
	}

	// 2. Query Postgres
	query := `
		SELECT config_key, config_value 
		FROM individual_tax_rates.tax_configs 
		WHERE tax_code = $1 
		AND valid_from <= $2 AND (valid_to IS NULL OR valid_to >= $2)`
	
	rows, err := r.db.Query(ctx, query, taxCode, refDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var key string
		var val decimal.Decimal
		if err := rows.Scan(&key, &val); err != nil {
			return nil, err
		}
		configs[key] = val
	}

	// 3. Salvar no Cache
	if len(configs) > 0 {
		cacheData, _ := json.Marshal(configs)
		r.rdb.Set(ctx, cacheKey, cacheData, 12*time.Hour)
	}

	return configs, nil
}

// GetTaxRulesForPeriod busca todas as fatias de imposto vigentes para um taxCode em uma data específica.
// Útil para cálculos progressivos como INSS.
func (r *TaxRepository) GetTaxRulesForPeriod(ctx context.Context, taxCode string, refDate time.Time) ([]models.TaxRule, error) {
	cacheKey := fmt.Sprintf("tax_rules_list:%s:%s", taxCode, refDate.Format("2006-01-02"))

	if val, err := r.rdb.Get(ctx, cacheKey).Result(); err == nil && val != "" {
		var rules []models.TaxRule
		if err := json.Unmarshal([]byte(val), &rules); err == nil {
			return rules, nil
		}
	}

	query := `
		SELECT 
			h.id, h.tax_definition_id, h.description, 
			h.range_min, h.range_max, h.aliq_percent, 
			h.deduction_val, h.valid_from, h.valid_to
		FROM individual_tax_rates.tax_rules_history h
		JOIN individual_tax_rates.tax_definitions d ON d.id = h.tax_definition_id
		WHERE d.tax_code = $1 
		  AND h.valid_from <= $2 
		  AND (h.valid_to IS NULL OR h.valid_to >= $2)
		ORDER BY h.range_min ASC`

	rows, err := r.db.Query(ctx, query, taxCode, refDate)
	if err != nil {
		return nil, fmt.Errorf("error querying tax rules: %w", err)
	}
	defer rows.Close()

	var rules []models.TaxRule
	for rows.Next() {
		var rule models.TaxRule
		err := rows.Scan(
			&rule.ID, &rule.TaxDefinitionID, &rule.Description,
			&rule.RangeMin, &rule.RangeMax, &rule.AliqPercent,
			&rule.DeductionVal, &rule.ValidFrom, &rule.ValidTo,
		)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}

	if len(rules) > 0 {
		data, _ := json.Marshal(rules)
		r.rdb.Set(ctx, cacheKey, data, 12*time.Hour)
	}

	return rules, nil
}
