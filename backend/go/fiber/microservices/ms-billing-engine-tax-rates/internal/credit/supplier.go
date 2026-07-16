// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/credit/supplier.go
package credit

import "context"

// SupplierStatus representa o status de qualificação fiscal de um fornecedor.
type SupplierStatus struct {
	CNPJ             string
	RegimeTributario  string // "LUCRO_REAL", "LUCRO_PRESUMIDO", "SIMPLES_NACIONAL"
	CertidaoValida   bool
	PermiteCredito   bool
	Status           string // "qualificado", "qualificado_restrito", "bloqueado"
}

// SupplierChecker verifica a qualificação fiscal de fornecedores (BR-08).
// A implementação real virá com GAP-007 (Qualificação Fiscal de Fornecedores).
type SupplierChecker interface {
	Check(ctx context.Context, cnpj string) (*SupplierStatus, error)
}

// AlwaysQualifiedSupplierChecker é um checker padrão que considera todo
// fornecedor como qualificado (Lucro Real, certidão válida).
// Será substituído pela implementação real no GAP-007.
type AlwaysQualifiedSupplierChecker struct{}

func (c *AlwaysQualifiedSupplierChecker) Check(ctx context.Context, cnpj string) (*SupplierStatus, error) {
	return &SupplierStatus{
		CNPJ:            cnpj,
		RegimeTributario: "LUCRO_REAL",
		CertidaoValida:  true,
		PermiteCredito:  true,
		Status:          "qualificado",
	}, nil
}

// BlockedSupplierChecker bloqueia crédito para testes (fornecedor não qualificado).
type BlockedSupplierChecker struct{}

func (c *BlockedSupplierChecker) Check(ctx context.Context, cnpj string) (*SupplierStatus, error) {
	return &SupplierStatus{
		CNPJ:            cnpj,
		RegimeTributario: "SIMPLES_NACIONAL",
		CertidaoValida:  false,
		PermiteCredito:  false,
		Status:          "bloqueado",
	}, nil
}
