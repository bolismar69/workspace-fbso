// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/credit/models.go
package credit

import (
	"time"

	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

// CreditCalculationRequest representa a requisição de cálculo de créditos (BR-08).
type CreditCalculationRequest struct {
	models.DocumentoFiscalEntrada
	CNPJFornecedor string `json:"cnpj_fornecedor"`
}

// CreditCalculationResponse representa o resultado do cálculo de créditos.
type CreditCalculationResponse struct {
	IDTransaction  string          `json:"id_transaction"`
	CNPJFornecedor string          `json:"cnpj_fornecedor"`
	CreditoCBS     decimal.Decimal `json:"credito_cbs"`
	CreditoIBS     decimal.Decimal `json:"credito_ibs"`
	CreditoTotal   decimal.Decimal `json:"credito_total"`
	MotivoBloqueio string          `json:"motivo_bloqueio,omitempty"`
	Itens          []CreditItem    `json:"itens"`
}

// CreditItem representa o crédito calculado por item.
type CreditItem struct {
	SKU          string          `json:"sku"`
	CBSDestacado decimal.Decimal `json:"cbs_destacado"`
	IBSDestacado decimal.Decimal `json:"ibs_destacado"`
	Creditavel   bool            `json:"creditavel"`
}

// CreditSummaryRequest para consulta de total de créditos no período.
type CreditSummaryRequest struct {
	Periodo string `json:"periodo"` // YYYY-MM
}

// CreditSummaryResponse para total de créditos disponíveis vs apropriados.
type CreditSummaryResponse struct {
	Periodo             string          `json:"periodo"`
	CreditoDisponivel   decimal.Decimal `json:"credito_disponivel"`
	CreditoApropriado   decimal.Decimal `json:"credito_apropriado"`
	TotalOperacoes      int             `json:"total_operacoes"`
	OperacoesBloqueadas int             `json:"operacoes_bloqueadas"`
	GeradoEm            time.Time       `json:"gerado_em"`
}
