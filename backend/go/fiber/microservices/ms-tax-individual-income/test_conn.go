package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"taxnexus-core-lib/db" // Sua lib local
  // "taxnexus-core-lib/cache"
)

func main() {
	// 1. Defina sua DSN (Ajuste usuário e senha se necessário)
	// Se estiver usando o usuário padrão do Postgres sem senha no localhost:
	// dsn := "postgres://postgres:root@localhost:5432/worker_db?sslmode=disable"
	dsn := "postgres://worker_user:worker_pass@localhost:5432/worker_db?sslmode=disable&search_path=individual_tax_rates"

	fmt.Printf("Verificando conexão com: %s\n", dsn)

	// 2. Tenta conectar usando a Lib que criamos
	pool, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatalf("Falha ao configurar o pool: %v", err)
	}
	defer pool.Close()

	// 3. Contexto com Timeout para o Ping
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 4. Executa o Ping Real
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco (Ping): %v", err)
	}

	fmt.Println("✅ Conexão estabelecida com sucesso no schema 'individual_tax_rates'!")
}
