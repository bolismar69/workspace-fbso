package db // <--- Adicione isto

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres(connString string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	return pgxpool.NewWithConfig(ctx, config)
}
