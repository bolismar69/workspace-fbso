package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"taxnexus-individual-core-lib/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type TaxRepository struct {
	db          *pgxpool.Pool
	rdb         *redis.Client
	logger      *slog.Logger
	cacheTTL    time.Duration
	cacheEnabled bool
}

func NewTaxRepository(db *pgxpool.Pool, rdb *redis.Client) *TaxRepository {
	ttl := 12 * time.Hour
	if v := os.Getenv("TAX_CACHE_TTL_HOURS"); v != "" {
		if h, err := strconv.Atoi(v); err == nil && h > 0 {
			ttl = time.Duration(h) * time.Hour
		}
	}

	return &TaxRepository{
		db:           db,
		rdb:          rdb,
		cacheTTL:     ttl,
		cacheEnabled: rdb != nil,
		logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})),
	}
}

func (r *TaxRepository) logCacheHit(key string) {
	r.logger.Debug("cache hit", "key", key, "ttl", r.cacheTTL.String())
}

func (r *TaxRepository) logCacheMiss(key string, reason string) {
	r.logger.Debug("cache miss", "key", key, "reason", reason)
}

func (r *TaxRepository) logCacheSet(key string, err error) {
	if err != nil {
		r.logger.Warn("cache set failed", "key", key, "error", err)
	} else {
		r.logger.Debug("cache set", "key", key, "ttl", r.cacheTTL.String())
	}
}

// GetApplicableRule busca a regra aplicável para um baseValue, utilizando cache Redis.
func (r *TaxRepository) GetApplicableRule(ctx context.Context, taxCode string, baseValue decimal.Decimal, refDate time.Time) (*models.TaxRule, error) {
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

// GetConfig busca uma configuração pontual no banco.
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

// GetTableConfigs busca todas as chaves de configuração via cache Redis ou PostgreSQL.
// Cache key: tax_configs:{taxCode}:{YYYY-MM-DD}
func (r *TaxRepository) GetTableConfigs(ctx context.Context, taxCode string, refDate time.Time) (map[string]decimal.Decimal, error) {
	cacheKey := fmt.Sprintf("tax_configs:%s:%s", taxCode, refDate.Format("2006-01-02"))
	configs := make(map[string]decimal.Decimal)

	// 1. Tentar Cache (somente se Redis estiver habilitado)
	if r.cacheEnabled && r.rdb != nil {
		val, err := r.rdb.Get(ctx, cacheKey).Result()
		if err == nil && val != "" {
			if err := json.Unmarshal([]byte(val), &configs); err == nil {
				r.logCacheHit(cacheKey)
				return configs, nil
			}
			r.logCacheMiss(cacheKey, "unmarshal_failed")
		} else if err != redis.Nil {
			r.logCacheMiss(cacheKey, fmt.Sprintf("redis_error: %v", err))
		} else {
			r.logCacheMiss(cacheKey, "not_cached")
		}
	}

	// 2. Query PostgreSQL
	query := `
		SELECT config_key, config_value 
		FROM individual_tax_rates.tax_configs 
		WHERE tax_code = $1 
		AND valid_from <= $2 AND (valid_to IS NULL OR valid_to >= $2)`

	rows, err := r.db.Query(ctx, query, taxCode, refDate)
	if err != nil {
		return nil, fmt.Errorf("error querying tax configs: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var key string
		var val decimal.Decimal
		if err := rows.Scan(&key, &val); err != nil {
			return nil, fmt.Errorf("error scanning tax config: %w", err)
		}
		configs[key] = val
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating tax configs: %w", err)
	}

	// 3. Salvar no Cache (best-effort)
	if len(configs) > 0 && r.cacheEnabled && r.rdb != nil {
		cacheData, err := json.Marshal(configs)
		if err == nil {
			setErr := r.rdb.Set(ctx, cacheKey, cacheData, r.cacheTTL).Err()
			r.logCacheSet(cacheKey, setErr)
		}
	}

	return configs, nil
}

// GetTaxRulesForPeriod busca todas as faixas de imposto vigentes via cache Redis ou PostgreSQL.
// Cache key: tax_rules_list:{taxCode}:{YYYY-MM-DD}
func (r *TaxRepository) GetTaxRulesForPeriod(ctx context.Context, taxCode string, refDate time.Time) ([]models.TaxRule, error) {
	cacheKey := fmt.Sprintf("tax_rules_list:%s:%s", taxCode, refDate.Format("2006-01-02"))

	// 1. Tentar Cache
	if r.cacheEnabled && r.rdb != nil {
		val, err := r.rdb.Get(ctx, cacheKey).Result()
		if err == nil && val != "" {
			var rules []models.TaxRule
			if err := json.Unmarshal([]byte(val), &rules); err == nil {
				r.logCacheHit(cacheKey)
				return rules, nil
			}
			r.logCacheMiss(cacheKey, "unmarshal_failed")
		} else if err != redis.Nil {
			r.logCacheMiss(cacheKey, fmt.Sprintf("redis_error: %v", err))
		} else {
			r.logCacheMiss(cacheKey, "not_cached")
		}
	}

	// 2. Query PostgreSQL
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
			return nil, fmt.Errorf("error scanning tax rule: %w", err)
		}
		rules = append(rules, rule)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating tax rules: %w", err)
	}

	// 3. Salvar no Cache (best-effort)
	if len(rules) > 0 && r.cacheEnabled && r.rdb != nil {
		data, err := json.Marshal(rules)
		if err == nil {
			setErr := r.rdb.Set(ctx, cacheKey, data, r.cacheTTL).Err()
			r.logCacheSet(cacheKey, setErr)
		}
	}

	return rules, nil
}
