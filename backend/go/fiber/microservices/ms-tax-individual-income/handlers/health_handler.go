package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

const healthCheckTimeout = 2 * time.Second

type componentStatus struct {
	Status    string `json:"status"`
	LatencyMs int64  `json:"latency_ms,omitempty"`
	Error     string `json:"error,omitempty"`
}

type healthResponse struct {
	Status     string                    `json:"status"`
	Timestamp  string                    `json:"timestamp"`
	Components map[string]componentStatus `json:"components"`
}

// HealthHandler expõe endpoints de health check para orquestradores (K8s, Docker).
type HealthHandler struct {
	pgPool *pgxpool.Pool
	rdb    *redis.Client
}

// NewHealthHandler cria um handler de health check com as dependências a serem verificadas.
func NewHealthHandler(pgPool *pgxpool.Pool, rdb *redis.Client) *HealthHandler {
	return &HealthHandler{pgPool: pgPool, rdb: rdb}
}

// Liveness responde apenas se o processo está vivo (sem checagens externas).
// GET /healthz
func (h *HealthHandler) Liveness(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":    "alive",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// Readiness verifica a saúde de todos os componentes (PostgreSQL, Redis).
// GET /health | GET /api/v1/health
func (h *HealthHandler) Readiness(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), healthCheckTimeout)
	defer cancel()

	components := make(map[string]componentStatus)
	allHealthy := true

	// PostgreSQL
	pgStatus := h.checkPostgres(ctx)
	components["postgresql"] = pgStatus
	if pgStatus.Status != "healthy" {
		allHealthy = false
	}

	// Redis (apenas se configurado)
	if h.rdb != nil {
		redisStatus := h.checkRedis(ctx)
		components["redis"] = redisStatus
		if redisStatus.Status != "healthy" {
			allHealthy = false
		}
	}

	status := "healthy"
	httpStatus := fiber.StatusOK
	if !allHealthy {
		status = "unhealthy"
		httpStatus = fiber.StatusServiceUnavailable
	}

	return c.Status(httpStatus).JSON(healthResponse{
		Status:     status,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Components: components,
	})
}

func (h *HealthHandler) checkPostgres(ctx context.Context) componentStatus {
	start := time.Now()
	err := h.pgPool.Ping(ctx)
	latency := time.Since(start).Milliseconds()

	if err != nil {
		return componentStatus{
			Status:    "unhealthy",
			LatencyMs: latency,
			Error:     err.Error(),
		}
	}

	return componentStatus{
		Status:    "healthy",
		LatencyMs: latency,
	}
}

func (h *HealthHandler) checkRedis(ctx context.Context) componentStatus {
	start := time.Now()
	err := h.rdb.Ping(ctx).Err()
	latency := time.Since(start).Milliseconds()

	if err != nil {
		return componentStatus{
			Status:    "unhealthy",
			LatencyMs: latency,
			Error:     err.Error(),
		}
	}

	return componentStatus{
		Status:    "healthy",
		LatencyMs: latency,
	}
}
