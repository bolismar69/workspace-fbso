// Este é o ponto de entrada do microserviço de cálculo de imposto de renda para pessoas físicas.
// path: backend/go/fiber/microservices/ms-tax-individual-income/main.go

package main

import (
	"context"
	"log"
	"os"
	"ms-tax-individual-income/handlers"
	"ms-tax-individual-income/services"
	"taxnexus-individual-core-lib/db"
	"taxnexus-individual-core-lib/cache"
	"taxnexus-individual-core-lib/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	// 1. Inicializa Conexões (Usando a Lib)
	pgPool, err := db.ConnectPostgres(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}

	if err := pgPool.Ping(context.Background()); err != nil {
    log.Fatal("Database is unreachable: ", err)
	}
	log.Println("Database connection verified.")
	
	rdb := cache.ConnectRedis(os.Getenv("REDIS_ADDR"))

	// Pega a URL do microserviço de INSS das variáveis de ambiente
	inssServiceURL := os.Getenv("INSS_SERVICE_URL")
	if inssServiceURL == "" {
		inssServiceURL = "http://localhost:3001" // Default local
	}

	// 2. Injeção de Dependência
	repo := repository.NewTaxRepository(pgPool, rdb)
	svc := services.NewCalculationService(repo, inssServiceURL) // <-- Alterado para INSS URL
	handler := handlers.NewTaxHandler(svc)

	// 3. Setup Fiber
	app := fiber.New()
	app.Use(logger.New())
	app.Use(requestid.New()) // Ativa geração de IDs únicos

	// Rotas
	api := app.Group("/api/v1")
	api.Post("/calculate/irpf", handler.CalculateIRPF)

	log.Fatal(app.Listen(":3000"))
}
