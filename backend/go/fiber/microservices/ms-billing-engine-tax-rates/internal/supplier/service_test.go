// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/supplier/service_test.go
package supplier

import (
	"context"
	"testing"
)

// ─── TST-007.01: POST /validate com CNPJ válido → qualificado ─────────────

func TestSupplier_ValidateQualificado(t *testing.T) {
	store := NewMemorySupplierStore()
	svc := NewValidationService(store)

	sup, err := svc.Validate(context.Background(), ValidateRequest{
		CNPJ: "12345678000199", RegimeTributario: "LUCRO_REAL", CertificadoRegularidade: true,
	})
	if err != nil {
		t.Fatalf("erro: %v", err)
	}
	if sup.Status != "qualificado" {
		t.Errorf("Status = %q, esperado 'qualificado'", sup.Status)
	}
	if !sup.PermiteCredito {
		t.Error("PermiteCredito deve ser true para Lucro Real")
	}
}

// ─── TST-007.02: Certidão vencida → bloqueado ──────────────────────────────

func TestSupplier_CertidaoVencida(t *testing.T) {
	svc := NewValidationService(NewMemorySupplierStore())

	sup, err := svc.Validate(context.Background(), ValidateRequest{
		CNPJ: "12345678000199", RegimeTributario: "LUCRO_REAL", CertificadoRegularidade: false,
	})
	if err != nil {
		t.Fatalf("erro: %v", err)
	}
	if sup.Status != "bloqueado" {
		t.Errorf("Status = %q, esperado 'bloqueado'", sup.Status)
	}
	if sup.PermiteCredito {
		t.Error("PermiteCredito deve ser false com certidão vencida")
	}
}

// ─── TST-007.03: Simples Nacional → qualificado_restrito ───────────────────

func TestSupplier_SimplesNacional(t *testing.T) {
	svc := NewValidationService(NewMemorySupplierStore())

	sup, err := svc.Validate(context.Background(), ValidateRequest{
		CNPJ: "11111111000191", RegimeTributario: "SIMPLES_NACIONAL", CertificadoRegularidade: true,
	})
	if err != nil {
		t.Fatalf("erro: %v", err)
	}
	if sup.Status != "qualificado_restrito" {
		t.Errorf("Status = %q, esperado 'qualificado_restrito'", sup.Status)
	}
	if !sup.PermiteCredito {
		t.Error("PermiteCredito deve ser true para Simples (restrito)")
	}
}

// ─── TST-007.04: PUT atualiza dados de qualificação ────────────────────────

func TestSupplier_Update(t *testing.T) {
	store := NewMemorySupplierStore()
	svc := NewValidationService(store)

	// Cadastra
	svc.Validate(context.Background(), ValidateRequest{
		CNPJ: "12345678000199", RegimeTributario: "LUCRO_REAL", CertificadoRegularidade: true,
	})

	// Atualiza para certidão vencida
	sup, err := svc.Update(context.Background(), "12345678000199", ValidateRequest{
		CNPJ: "12345678000199", RegimeTributario: "LUCRO_REAL", CertificadoRegularidade: false,
	})
	if err != nil {
		t.Fatalf("erro ao atualizar: %v", err)
	}
	if sup.Status != "bloqueado" {
		t.Errorf("Status após update = %q, esperado 'bloqueado'", sup.Status)
	}
}

// ─── TST-007.05: GET /supplier/{cnpj} → dados completos ────────────────────

func TestSupplier_GetByCNPJ(t *testing.T) {
	store := NewMemorySupplierStore()
	svc := NewValidationService(store)

	svc.Validate(context.Background(), ValidateRequest{
		CNPJ: "12345678000199", RegimeTributario: "LUCRO_REAL", CertificadoRegularidade: true,
	})

	sup, err := svc.GetByCNPJ(context.Background(), "12345678000199")
	if err != nil {
		t.Fatalf("erro: %v", err)
	}
	if sup.CNPJ != "12345678000199" {
		t.Errorf("CNPJ = %q", sup.CNPJ)
	}
}

// ─── TST-007.06: CNPJ não cadastrado → erro ────────────────────────────────

func TestSupplier_NotFound(t *testing.T) {
	svc := NewValidationService(NewMemorySupplierStore())

	_, err := svc.GetByCNPJ(context.Background(), "00000000000000")
	if err == nil {
		t.Fatal("esperado erro para CNPJ não cadastrado")
	}
}

// ─── TST-007.07: Integração CreditEngine — fornecedor bloqueado zera crédito

func TestSupplier_CreditIntegration(t *testing.T) {
	store := NewMemorySupplierStore()
	svc := NewValidationService(store)

	// Cadastra fornecedor como bloqueado
	svc.Validate(context.Background(), ValidateRequest{
		CNPJ: "00000000000000", RegimeTributario: "SIMPLES_NACIONAL", CertificadoRegularidade: false,
	})

	// SupplierCheckerAdapter deve retornar bloqueado
	adapter := NewSupplierCheckerAdapter(svc)
	status, err := adapter.Check(context.Background(), "00000000000000")
	if err != nil {
		t.Fatalf("Check não deve errar: %v", err)
	}
	if status.PermiteCredito {
		t.Error("PermiteCredito deve ser false para fornecedor bloqueado")
	}
	if status.Status != "bloqueado" {
		t.Errorf("Status = %q, esperado 'bloqueado'", status.Status)
	}
}

// ─── TST-007.08: CNPJ inválido → erro de validação ─────────────────────────

func TestSupplier_CNPJInvalido(t *testing.T) {
	svc := NewValidationService(NewMemorySupplierStore())

	_, err := svc.Validate(context.Background(), ValidateRequest{
		CNPJ: "ABC", RegimeTributario: "LUCRO_REAL", CertificadoRegularidade: true,
	})
	if err == nil {
		t.Fatal("esperado erro para CNPJ inválido")
	}
}
