// path: backend/go/libs/go-native/taxnexus-billing-core-lib/repository/tax_repository.go
package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bolismar/taxnexus-billing-core-lib/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

// type TaxRepository struct {
// 	db  *pgxpool.Pool
// 	rdb *redis.Client
// }

// func NewTaxRepository(db *pgxpool.Pool, rdb *redis.Client) *TaxRepository {
// 	return &TaxRepository{db: db, rdb: rdb}
// }

type TaxRepository interface {
	GetRateByNCM(ctx context.Context, ncm string, uf string) (float64, error)
	GetIBSRate(ctx context.Context, municipioIBGE string) (float64, error)
}

func (r *TaxRepository) GetRateByNCM(ctx context.Context, ncm string, uf string) (float64, error) {
	var rate float64
	err := r.db.QueryRow(ctx, "SELECT rate FROM tax_rates WHERE ncm = $1 AND uf = $2", ncm, uf).Scan(&rate)
	if err != nil {
		return 0, fmt.Errorf("failed to get tax rate: %w", err)
	}
	return rate, nil
}

func (r *TaxRepository) GetIBSRate(ctx context.Context, municipioIBGE string) (float64, error) {
	var rate float64
	err := r.db.QueryRow(ctx, "SELECT ibs_rate FROM ibs_rates WHERE municipio_ibge = $1", municipioIBGE).Scan(&rate)
	if err != nil {
		return 0, fmt.Errorf("failed to get IBS rate: %w", err)
	}
	return rate, nil
}
