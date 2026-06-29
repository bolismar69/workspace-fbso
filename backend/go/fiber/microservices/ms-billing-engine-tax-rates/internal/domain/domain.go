// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/domain/domain.go
// Package domain define as interfaces e tipos centrais do dominio fiscal.
// Como camada mais interna da arquitetura, nao possui dependencias de outros
// pacotes internos — apenas da lib compartilhada taxnexus-billing-core-lib.
package domain

import (
	"context"

	"taxnexus-billing-core-lib/models"
)

// TaxCalculator e a interface central do motor de calculo tributario.
// Cada implementacao (ICMS, IPI, PIS/COFINS, Reforma, etc.) deve satisfaze-la.
// O BillingEngine orquestra as calculadoras registradas e consolida os resultados.
//
// A interface retorna uma lista de itens fiscais ja calculados — a engine
// encarrega-se de mergear os tributos de cada item no response final.
type TaxCalculator interface {
	Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error)
}
