// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/cmd/api/versioning_test.go
package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

// setupVersioningApp cria uma app mínima com as rotas versionadas para teste.
func setupVersioningApp() *fiber.App {
	app := fiber.New()

	v1 := app.Group("/v1")

	healthzHandler := func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"status": "ok"})
	}
	v1.Get("/healthz", healthzHandler)
	app.Get("/healthz", withDeprecation(healthzHandler))

	healthHandler := func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"status": "ok", "checks": fiber.Map{}})
	}
	v1.Get("/health", healthHandler)
	app.Get("/health", withDeprecation(healthHandler))

	calculateHandler := func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"result": "ok"})
	}
	v1.Post("/calculate", calculateHandler)
	app.Post("/calculate", withDeprecation(calculateHandler))

	return app
}

// ─── TST-009.01: /v1/calculate retorna 200 com payload válido ──────────────

func TestVersioning_V1Calculate(t *testing.T) {
	app := setupVersioningApp()

	req := httptest.NewRequest("POST", "/v1/calculate", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("erro ao testar /v1/calculate: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("esperado 200, obtido %d", resp.StatusCode)
	}
}

// ─── TST-009.02: /calculate (legado) retorna Deprecation e Link headers ─────

func TestVersioning_LegacyCalculateDeprecation(t *testing.T) {
	app := setupVersioningApp()

	req := httptest.NewRequest("POST", "/calculate", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("erro ao testar /calculate: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("esperado 200, obtido %d", resp.StatusCode)
	}

	deprecation := resp.Header.Get("Deprecation")
	if deprecation != "true" {
		t.Errorf("esperado Deprecation: true, obtido %q", deprecation)
	}

	sunset := resp.Header.Get("Sunset")
	if sunset == "" {
		t.Error("esperado header Sunset não vazio")
	}

	link := resp.Header.Get("Link")
	if link == "" {
		t.Error("esperado header Link não vazio")
	}
	t.Logf("Legacy /calculate headers: Deprecation=%q, Sunset=%q, Link=%q",
		deprecation, sunset, link)
}

// ─── TST-009.03: /v1/healthz e /v1/health retornam 200 ──────────────────────

func TestVersioning_V1HealthEndpoints(t *testing.T) {
	app := setupVersioningApp()

	tests := []struct {
		method string
		path   string
	}{
		{"GET", "/v1/healthz"},
		{"GET", "/v1/health"},
		{"GET", "/healthz"},
		{"GET", "/health"},
	}

	for _, tt := range tests {
		t.Run(tt.method+" "+tt.path, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("erro: %v", err)
			}

			if resp.StatusCode != 200 {
				t.Errorf("esperado 200, obtido %d", resp.StatusCode)
			}

			// Health legacy paths must have Deprecation header
			if tt.path == "/healthz" || tt.path == "/health" {
				dep := resp.Header.Get("Deprecation")
				if dep != "true" {
					t.Errorf("%s: esperado Deprecation: true, obtido %q", tt.path, dep)
				}
			}
		})
	}
}

// ─── TST-009.04: Endpoint inexistente /v2/calculate retorna 404 ────────────

func TestVersioning_UnknownVersion(t *testing.T) {
	app := setupVersioningApp()

	req := httptest.NewRequest("POST", "/v2/calculate", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("erro: %v", err)
	}

	if resp.StatusCode != 404 {
		t.Errorf("esperado 404 para /v2/calculate, obtido %d", resp.StatusCode)
	}

	// /v1/calculate ainda deve funcionar
	req2 := httptest.NewRequest("POST", "/v1/calculate", nil)
	resp2, _ := app.Test(req2)
	if resp2.StatusCode != 200 {
		t.Errorf("/v1/calculate deveria funcionar, obtido %d", resp2.StatusCode)
	}
}

// ─── Teste adicional: Legado healthz tem Deprecation ────────────────────────

func TestVersioning_LegacyHealthzDeprecation(t *testing.T) {
	app := setupVersioningApp()

	req := httptest.NewRequest("GET", "/healthz", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("erro: %v", err)
	}

	deprecation := resp.Header.Get("Deprecation")
	if deprecation != "true" {
		t.Errorf("esperado Deprecation: true em /healthz, obtido %q", deprecation)
	}

	body, _ := io.ReadAll(resp.Body)
	t.Logf("Legacy /healthz body: %s", string(body))
}
