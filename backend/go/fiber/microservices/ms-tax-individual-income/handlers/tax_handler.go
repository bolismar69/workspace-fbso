// path: backend/go/fiber/microservices/ms-tax-individual-income/handlers/tax_handler.go
package handlers

import (
	"context"
	"time"
	"taxnexus-individual-core-lib/models"
	"ms-tax-individual-income/services"
	"github.com/gofiber/fiber/v2"
)

type TaxHandler struct {
	svc *services.CalculationService
}

func NewTaxHandler(svc *services.CalculationService) *TaxHandler {
	return &TaxHandler{svc: svc}
}

func (h *TaxHandler) CalculateIRPF(c *fiber.Ctx) error {
	// 1. Extrai o Request ID gerado pelo middleware do Fiber
	requestID := c.GetRespHeader(fiber.HeaderXRequestID)
	
	// 2. Injeta o ID no contexto Go padrão
	ctx := context.WithValue(c.UserContext(), "requestid", requestID)

	var req models.UniversalTaxRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "corpo da requisição inválido"})
	}

	// 3. Define data padrão se omitida
	if req.ReferenceDate.IsZero() {
		req.ReferenceDate = time.Now()
	}

	result, err := h.svc.Calculate(ctx, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}