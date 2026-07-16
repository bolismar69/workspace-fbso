//go:build ignore
// +build ignore

// Script de diagnóstico de conexão PostgreSQL.
// Execute com: go run test_conn.go

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"taxnexus-individual-core-lib/db"
)

func main() {
	dsn := "postgres://worker_user:worker_pass@localhost:5432/worker_db?sslmode=disable&search_path=individual_tax_rates"

	fmt.Printf("Verificando conexão com: %s\n", dsn)

	pool, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatalf("Falha ao configurar o pool: %v", err)
	}
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco (Ping): %v", err)
	}

	fmt.Println("✅ Conexão estabelecida com sucesso no schema 'individual_tax_rates'!")
}
