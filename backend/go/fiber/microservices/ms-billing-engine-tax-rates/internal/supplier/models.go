// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/supplier/models.go
package supplier

import (
	"context"
	"time"
)

// SupplierFiscal representa a qualificação fiscal de um fornecedor (BR-08).
type SupplierFiscal struct {
	CNPJ                  string    `json:"cnpj"`
	RegimeTributario      string    `json:"regime_tributario"`
	CertificadoRegularidade bool    `json:"certificado_regularidade"`
	PermiteCredito        bool      `json:"permite_credito"`
	DataQualificacao      time.Time `json:"data_qualificacao"`
	DataValidade          time.Time `json:"data_validade"`
	Status                string    `json:"status"` // "qualificado", "qualificado_restrito", "bloqueado"
}

// SupplierStore define o contrato de persistência para fornecedores.
type SupplierStore interface {
	Save(ctx context.Context, s *SupplierFiscal) error
	FindByCNPJ(ctx context.Context, cnpj string) (*SupplierFiscal, error)
	Update(ctx context.Context, cnpj string, s *SupplierFiscal) error
}

// DetermineStatus calcula o status com base no regime e certidão.
func DetermineStatus(regime string, certidaoValida bool) string {
	if !certidaoValida {
		return "bloqueado"
	}
	switch regime {
	case "LUCRO_REAL", "LUCRO_PRESUMIDO":
		return "qualificado"
	case "SIMPLES_NACIONAL":
		return "qualificado_restrito"
	default:
		return "bloqueado"
	}
}
