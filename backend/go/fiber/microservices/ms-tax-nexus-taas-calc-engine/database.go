package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	// DSN configurada para o schema tax_nexus_taas
	dsn := "postgres://bolismar:admin_tax@localhost:5432/tax_nexus_db?search_path=tax_nexus_taas"
	
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao configurar pool: %v\n", err)
		os.Exit(1)
	}

	// Configurações de performance do Pool
	config.MaxConns = 20
	config.MinConns = 5

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao conectar ao banco: %v\n", err)
		os.Exit(1)
	}

	DB = pool
	fmt.Println("🐘 Conectado ao PostgreSQL no schema tax_nexus_taas")
}
