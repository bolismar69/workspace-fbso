// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/token/service.go
package token

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"taxnexus-billing-core-lib/repository"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// EnvTaxTokenTTLMinutes é a variável de ambiente para o TTL do token em minutos.
	EnvTaxTokenTTLMinutes = "TAX_TOKEN_TTL_MINUTES"
	defaultTTLMinutes     = 60
)

// TokenGenerateRequest representa a requisição para gerar um token fiscal.
type TokenGenerateRequest struct {
	NCM           string `json:"ncm"`
	UFOrigem      string `json:"uf_origem"`
	UFDestino     string `json:"uf_destino"`
	MunicipioIBGE string `json:"municipio_ibge"`
}

// TokenStatusResponse representa a resposta de consulta de status do token.
type TokenStatusResponse struct {
	ID            string          `json:"id"`
	Status        string          `json:"status"` // "valido" ou "expirado"
	NCM           string          `json:"ncm"`
	UFDestino     string          `json:"uf_destino"`
	MunicipioIBGE string          `json:"municipio_ibge"`
	AliquotaCBS   decimal.Decimal `json:"aliquota_cbs"`
	AliquotaIBS   decimal.Decimal `json:"aliquota_ibs"`
	AliquotaIS    decimal.Decimal `json:"aliquota_is"`
	ExpiresAt     time.Time       `json:"expires_at"`
	CreatedAt     time.Time       `json:"created_at"`
}

// TokenService gerencia o ciclo de vida dos tokens fiscais (BR-06).
type TokenService struct {
	store      TokenStore
	repo       repository.TaxRepository
	ttlMinutes int
}

// NewTokenService cria um TokenService com TTL via env var TAX_TOKEN_TTL_MINUTES.
func NewTokenService(store TokenStore, repo repository.TaxRepository) *TokenService {
	ttl := defaultTTLMinutes
	if v := os.Getenv(EnvTaxTokenTTLMinutes); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			ttl = n
		} else {
			slog.Warn("TAX_TOKEN_TTL_MINUTES inválido, usando default",
				"valor", v, "default", defaultTTLMinutes)
		}
	}

	slog.Info("TokenService inicializado",
		"ttl_minutes", ttl,
	)

	// Background cleanup a cada 5 minutos
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			if deleted, err := store.DeleteExpired(context.Background()); err != nil {
				slog.Warn("Falha ao limpar tokens expirados", "error", err)
			} else if deleted > 0 {
				slog.Debug("Tokens expirados removidos", "count", deleted)
			}
		}
	}()

	return &TokenService{
		store:      store,
		repo:       repo,
		ttlMinutes: ttl,
	}
}

// Generate consulta as alíquotas vigentes e gera um TaxToken.
//
// Fluxo:
//  1. Consulta iva_dual_rules para (NCM, UF destino, município IBGE)
//  2. Se não encontrar regra, retorna erro
//  3. Gera UUID, calcula ExpiresAt (now + TTL)
//  4. Persiste no store
//  5. Retorna o token
func (s *TokenService) Generate(ctx context.Context, req TokenGenerateRequest) (*TaxToken, error) {
	// 1. Consulta alíquotas vigentes
	rule, err := s.repo.GetIvaDualRule(ctx, req.NCM, req.UFDestino, req.MunicipioIBGE)
	if err != nil {
		slog.Warn("Falha ao consultar alíquotas para geração de token",
			"ncm", req.NCM,
			"uf_destino", req.UFDestino,
			"municipio_ibge", req.MunicipioIBGE,
			"error", err,
		)
		return nil, fmt.Errorf("não foi possível consultar alíquotas para NCM %s / UF %s: %w",
			req.NCM, req.UFDestino, err)
	}
	if rule == nil {
		return nil, fmt.Errorf("nenhuma regra de alíquota encontrada para NCM %s / UF %s / IBGE %s",
			req.NCM, req.UFDestino, req.MunicipioIBGE)
	}

	// 2. Cria o token com as alíquotas congeladas
	now := time.Now()
	token := &TaxToken{
		ID:                   uuid.NewString(),
		NCM:                  req.NCM,
		UFOrigem:             req.UFOrigem,
		UFDestino:            req.UFDestino,
		MunicipioIBGE:        req.MunicipioIBGE,
		AliquotaCBS:          rule.AliquotaCBS,
		AliquotaIBSEstadual:  rule.AliquotaIBSEstadual,
		AliquotaIBSMunicipal: rule.AliquotaIBSMunicipal,
		AliquotaIS:           rule.AliquotaIS,
		ExpiresAt:            now.Add(time.Duration(s.ttlMinutes) * time.Minute),
		CreatedAt:            now,
	}

	// 3. Persiste
	if err := s.store.Save(ctx, token); err != nil {
		return nil, fmt.Errorf("falha ao persistir token: %w", err)
	}

	slog.Info("Token fiscal gerado",
		"token_id", token.ID,
		"ncm", token.NCM,
		"uf_destino", token.UFDestino,
		"municipio_ibge", token.MunicipioIBGE,
		"expires_at", token.ExpiresAt,
		"ttl_minutes", s.ttlMinutes,
	)

	return token, nil
}

// Validate verifica se um token é válido (existe e não expirou).
// Retorna o token se válido, ou erro com código apropriado.
func (s *TokenService) Validate(ctx context.Context, tokenID string) (*TaxToken, error) {
	token, err := s.store.FindByID(ctx, tokenID)
	if err != nil {
		return nil, fmt.Errorf("token não encontrado: %w", err)
	}

	if token.IsExpired() {
		slog.Warn("Token expirado",
			"token_id", tokenID,
			"expires_at", token.ExpiresAt,
		)
		return nil, ErrTokenExpired{ID: tokenID, ExpiresAt: token.ExpiresAt}
	}

	return token, nil
}

// Status retorna o status do token (válido ou expirado).
// Diferente de Validate, NÃO retorna erro para token expirado —
// apenas retorna status="expirado".
func (s *TokenService) Status(ctx context.Context, tokenID string) (*TokenStatusResponse, error) {
	token, err := s.store.FindByID(ctx, tokenID)
	if err != nil {
		return nil, fmt.Errorf("token não encontrado: %w", err)
	}

	status := "valido"
	if token.IsExpired() {
		status = "expirado"
	}

	return &TokenStatusResponse{
		ID:            token.ID,
		Status:        status,
		NCM:           token.NCM,
		UFDestino:     token.UFDestino,
		MunicipioIBGE: token.MunicipioIBGE,
		AliquotaCBS:   token.AliquotaCBS,
		AliquotaIBS:   token.AliquotaIBSEstadual.Add(token.AliquotaIBSMunicipal),
		AliquotaIS:    token.AliquotaIS,
		ExpiresAt:     token.ExpiresAt,
		CreatedAt:     token.CreatedAt,
	}, nil
}

// ErrTokenExpired representa um token que expirou.
type ErrTokenExpired struct {
	ID        string
	ExpiresAt time.Time
}

func (e ErrTokenExpired) Error() string {
	return fmt.Sprintf("token %s expirado em %s", e.ID, e.ExpiresAt.Format(time.RFC3339))
}
