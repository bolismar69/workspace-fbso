// path: backend/go/libs/go-native/taxnexus-billing-core-lib/db/postgres.go
package db // <--- Adicione isto

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

func ConnectPostgres(connString string) (*pgxpool.Pool, error) {
	slog.Info("Conectando ao banco de dados fiscal", "db", "PostgreSQL", "connString", connString)
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	return pgxpool.NewWithConfig(ctx, config)
}
