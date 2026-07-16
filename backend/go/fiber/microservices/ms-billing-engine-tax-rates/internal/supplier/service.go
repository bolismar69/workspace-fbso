// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/supplier/service.go
package supplier

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"ms-billing-engine-tax-rates/internal/credit"
)

// ValidationService gerencia a qualificação fiscal de fornecedores (GAP-007).
type ValidationService struct {
	store SupplierStore
}

func NewValidationService(store SupplierStore) *ValidationService {
	return &ValidationService{store: store}
}

// ValidateRequest é o payload para validação de fornecedor.
type ValidateRequest struct {
	CNPJ                  string `json:"cnpj"`
	RegimeTributario      string `json:"regime_tributario"`
	CertificadoRegularidade bool `json:"certificado_regularidade"`
}

// Validate cadastra ou atualiza a qualificação fiscal de um fornecedor.
func (s *ValidationService) Validate(ctx context.Context, req ValidateRequest) (*SupplierFiscal, error) {
	if len(req.CNPJ) != 14 {
		return nil, fmt.Errorf("CNPJ deve ter 14 dígitos, recebido %q (%d)", req.CNPJ, len(req.CNPJ))
	}
	for _, c := range req.CNPJ {
		if c < '0' || c > '9' {
			return nil, fmt.Errorf("CNPJ deve conter apenas dígitos")
		}
	}

	status := DetermineStatus(req.RegimeTributario, req.CertificadoRegularidade)
	permite := status == "qualificado" || status == "qualificado_restrito"

	sup := &SupplierFiscal{
		CNPJ:                  req.CNPJ,
		RegimeTributario:      req.RegimeTributario,
		CertificadoRegularidade: req.CertificadoRegularidade,
		PermiteCredito:        permite,
		DataQualificacao:      time.Now(),
		DataValidade:          time.Now().AddDate(1, 0, 0), // 1 ano
		Status:                status,
	}

	// Upsert: tenta salvar, se já existe atualiza
	if err := s.store.Save(ctx, sup); err != nil {
		if err := s.store.Update(ctx, req.CNPJ, sup); err != nil {
			return nil, fmt.Errorf("falha ao persistir fornecedor: %w", err)
		}
	}

	slog.Info("Fornecedor qualificado",
		"cnpj", sup.CNPJ,
		"regime", sup.RegimeTributario,
		"status", sup.Status,
	)

	return sup, nil
}

// GetByCNPJ consulta o status de qualificação de um fornecedor.
func (s *ValidationService) GetByCNPJ(ctx context.Context, cnpj string) (*SupplierFiscal, error) {
	return s.store.FindByCNPJ(ctx, cnpj)
}

// Update atualiza os dados de qualificação de um fornecedor existente.
func (s *ValidationService) Update(ctx context.Context, cnpj string, req ValidateRequest) (*SupplierFiscal, error) {
	status := DetermineStatus(req.RegimeTributario, req.CertificadoRegularidade)

	sup := &SupplierFiscal{
		CNPJ:                  req.CNPJ,
		RegimeTributario:      req.RegimeTributario,
		CertificadoRegularidade: req.CertificadoRegularidade,
		PermiteCredito:        status == "qualificado" || status == "qualificado_restrito",
		DataQualificacao:      time.Now(),
		DataValidade:          time.Now().AddDate(1, 0, 0),
		Status:                status,
	}

	if err := s.store.Update(ctx, cnpj, sup); err != nil {
		return nil, err
	}

	slog.Info("Fornecedor atualizado", "cnpj", cnpj, "status", sup.Status)
	return sup, nil
}

// ─── CreditEngine Adapter ──────────────────────────────────────────────────

// SupplierCheckerAdapter adapta o ValidationService para a interface
// credit.SupplierChecker usada pelo CreditEngine (GAP-005 ↔ GAP-007).
type SupplierCheckerAdapter struct {
	svc *ValidationService
}

func NewSupplierCheckerAdapter(svc *ValidationService) *SupplierCheckerAdapter {
	return &SupplierCheckerAdapter{svc: svc}
}

// Check implementa credit.SupplierChecker consultando o fornecedor cadastrado.
func (a *SupplierCheckerAdapter) Check(ctx context.Context, cnpj string) (*credit.SupplierStatus, error) {
	sup, err := a.svc.GetByCNPJ(ctx, cnpj)
	if err != nil {
		// Fornecedor não cadastrado → bloqueado por padrão
		slog.Warn("Fornecedor não cadastrado — crédito bloqueado", "cnpj", cnpj)
		return &credit.SupplierStatus{
			CNPJ:            cnpj,
			RegimeTributario: "NAO_CADASTRADO",
			CertidaoValida:  false,
			PermiteCredito:  false,
			Status:          "bloqueado",
		}, nil
	}

	// Verifica também se a certidão ainda está válida
	certidaoValida := sup.CertificadoRegularidade && time.Now().Before(sup.DataValidade)

	return &credit.SupplierStatus{
		CNPJ:            sup.CNPJ,
		RegimeTributario: sup.RegimeTributario,
		CertidaoValida:  certidaoValida,
		PermiteCredito:  sup.PermiteCredito && certidaoValida,
		Status:          sup.Status,
	}, nil
}
