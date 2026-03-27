// path: backend/go/fiber/microservices/ms-tax-individual-social-security/handlers/inss_handler.go
package handlers

import (
	"context"
	"time"

	"ms-tax-individual-social-security/services"
	"taxnexus-core-lib/models"

	"github.com/gofiber/fiber/v2"
)

type INSSHandler struct {
	svc *services.INSSService
}

func NewINSSHandler(svc *services.INSSService) *INSSHandler {
	return &INSSHandler{svc: svc}
}

func (h *INSSHandler) Calculate(c *fiber.Ctx) error {
	// 1. Extrai o ID da requisição para rastreabilidade
	rid := c.GetRespHeader(fiber.HeaderXRequestID)
	ctx := context.WithValue(c.UserContext(), "requestid", rid)

	var req models.UniversalTaxRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "falha ao processar o corpo da requisição",
		})
	}

	// 2. Default para a data atual se não enviado
	if req.ReferenceDate.IsZero() {
		req.ReferenceDate = time.Now()
	}

	// 3. Força o TaxCode para garantir consistência
	req.TaxCode = "INSS"

	result, err := h.svc.Calculate(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
