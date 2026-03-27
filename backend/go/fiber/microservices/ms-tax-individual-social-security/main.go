// path: backend/go/fiber/microservices/ms-tax-individual-social-security/main.go
package main

import (
	"context"
	"log"
	"os"
	"time"

	"ms-tax-individual-social-security/handlers"
	"ms-tax-individual-social-security/services"
	"taxnexus-core-lib/cache"
	"taxnexus-core-lib/db"
	"taxnexus-core-lib/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	// 1. Conexões de Infraestrutura
	pgPool, err := db.ConnectPostgres(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("erro ao conectar no postgres: ", err)
	}
	defer pgPool.Close()

	// Health check da conexão
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := pgPool.Ping(ctx); err != nil {
		log.Fatal("banco de dados inacessível: ", err)
	}

	rdb := cache.ConnectRedis(os.Getenv("REDIS_ADDR"))

	// 2. Injeção de Dependências
	repo := repository.NewTaxRepository(pgPool, rdb)
	svc := services.NewINSSService(repo)
	handler := handlers.NewINSSHandler(svc)

	// 3. Inicialização do Fiber
	app := fiber.New(fiber.Config{
		AppName: "TaxNexus - Individual Social Security (INSS)",
	})

	// Middlewares
	app.Use(requestid.New()) // Gera o X-Request-ID
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ID=${respHeader:X-Request-ID}\n",
	}))

	// Rotas
	api := app.Group("/api/v1")
	api.Post("/calculate/inss", handler.Calculate)

	// Início do Servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001" // Porta diferente do ms-tax-individual-income
	}

	log.Printf("Serviço INSS rodando na porta %s", port)
	log.Fatal(app.Listen(":" + port))
}