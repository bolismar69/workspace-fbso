// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/token/token.go
package token

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

// TaxToken representa um congelamento de alíquotas vigentes em um ponto no tempo,
// garantindo que um cálculo futuro use as mesmas alíquotas da simulação original
// (BR-06 — Garantia de Preço Ofertado).
//
// O token é gerado a partir de uma consulta às alíquotas atuais (CBS, IBS estadual,
// IBS municipal, IS) para uma tupla (NCM, UF destino, município IBGE) e é válido
// por um TTL configurável (TAX_TOKEN_TTL_MINUTES, default 60 min).
type TaxToken struct {
	ID                  string          `json:"id"`
	NCM                 string          `json:"ncm"`
	UFOrigem            string          `json:"uf_origem"`
	UFDestino           string          `json:"uf_destino"`
	MunicipioIBGE       string          `json:"municipio_ibge"`
	AliquotaCBS         decimal.Decimal `json:"aliquota_cbs"`
	AliquotaIBSEstadual decimal.Decimal `json:"aliquota_ibs_estadual"`
	AliquotaIBSMunicipal decimal.Decimal `json:"aliquota_ibs_municipal"`
	AliquotaIS          decimal.Decimal `json:"aliquota_is"`
	ExpiresAt           time.Time       `json:"expires_at"`
	CreatedAt           time.Time       `json:"created_at"`
}

// IsExpired verifica se o token expirou.
func (t *TaxToken) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}

// TokenStore define o contrato de persistência para tokens fiscais.
type TokenStore interface {
	// Save persiste um novo token.
	Save(ctx context.Context, token *TaxToken) error

	// FindByID recupera um token pelo seu ID.
	FindByID(ctx context.Context, id string) (*TaxToken, error)

	// DeleteExpired remove tokens expirados. Retorna o número de removidos.
	DeleteExpired(ctx context.Context) (int64, error)
}
