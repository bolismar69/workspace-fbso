// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/admin/models.go
package admin

import (
	"time"

	"github.com/shopspring/decimal"
)

// IvaDualRuleInput representa o payload para criação/atualização de regra IVA Dual.
type IvaDualRuleInput struct {
	NCM                  string          `json:"ncm"`
	UFDestino            string          `json:"uf_destino"`
	MunicipioDestinoIBGE string          `json:"municipio_destino_ibge,omitempty"`
	AliquotaCBS          decimal.Decimal `json:"aliquota_cbs"`
	AliquotaIBSEstadual  decimal.Decimal `json:"aliquota_ibs_estadual"`
	AliquotaIBSMunicipal decimal.Decimal `json:"aliquota_ibs_municipal"`
	PercentualReducao    decimal.Decimal `json:"percentual_reducao"`
	IsImpostoSeletivo    bool            `json:"is_imposto_seletivo"`
	AliquotaIS           decimal.Decimal `json:"aliquota_is"`
	InicioValidade       time.Time       `json:"inicio_validade"`
	FinalValidade        *time.Time      `json:"final_validade,omitempty"`
}

// IvaDualRuleOutput representa a resposta com a regra IVA Dual.
type IvaDualRuleOutput struct {
	ID                   int64           `json:"id"`
	NCM                  string          `json:"ncm"`
	UFDestino            string          `json:"uf_destino"`
	MunicipioDestinoIBGE string          `json:"municipio_destino_ibge,omitempty"`
	AliquotaCBS          decimal.Decimal `json:"aliquota_cbs"`
	AliquotaIBSEstadual  decimal.Decimal `json:"aliquota_ibs_estadual"`
	AliquotaIBSMunicipal decimal.Decimal `json:"aliquota_ibs_municipal"`
	PercentualReducao    decimal.Decimal `json:"percentual_reducao"`
	IsImpostoSeletivo    bool            `json:"is_imposto_seletivo"`
	AliquotaIS           decimal.Decimal `json:"aliquota_is"`
	InicioValidade       time.Time       `json:"inicio_validade"`
	FinalValidade        *time.Time      `json:"final_validade,omitempty"`
}

// ListRulesFilter representa filtros para listagem de regras.
type ListRulesFilter struct {
	NCM           string `json:"ncm,omitempty"`
	UF            string `json:"uf,omitempty"`
	AtivasApenas  bool   `json:"ativas_apenas"`
}
