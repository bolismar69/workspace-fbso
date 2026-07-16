// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/legacy/cofins_strategies.go
package legacy

import (
	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/utils"

	"github.com/shopspring/decimal"
)

type COFINSStrategy interface {
	Calculate(item models.ItemDocumentoFiscalEntrada, info COFINSInfo) decimal.Decimal
}

type COFINSInfo struct {
	CST       string
	Aliquota  decimal.Decimal
	ValorIcms decimal.Decimal
}

// COFINS01_02 implementa CST 01 (Básica) e 02 (Diferenciada) - Ad Valorem
type COFINS01_02 struct{}

func (s *COFINS01_02) Calculate(item models.ItemDocumentoFiscalEntrada, info COFINSInfo) decimal.Decimal {
	vItem := item.ValorUnitario.Mul(item.Quantidade)
	vIcms := info.ValorIcms
	aliq := info.Aliquota.Div(decimal.NewFromInt(100))

	base := vItem.Sub(vIcms)
	if base.IsNegative() {
		base = decimal.Zero
	}

	return base.Mul(aliq).Round(2)
}

// COFINS03 implementa CST 03 (Alíquota por Unidade)
type COFINS03 struct{}

func (s *COFINS03) Calculate(item models.ItemDocumentoFiscalEntrada, info COFINSInfo) decimal.Decimal {
	return utils.CalcTax(item.Quantidade, info.Aliquota).Round(2)
}

// COFINS04 implementa CST 04 (Monofásico) — tributo concentrado no produtor/importador
type COFINS04 struct{}

func (s *COFINS04) Calculate(item models.ItemDocumentoFiscalEntrada, info COFINSInfo) decimal.Decimal {
	return decimal.Zero
}

// COFINS05 implementa CST 05 (Substituição Tributária) — tributo já recolhido por substituição
type COFINS05 struct{}

func (s *COFINS05) Calculate(item models.ItemDocumentoFiscalEntrada, info COFINSInfo) decimal.Decimal {
	return decimal.Zero
}

// COFINS06 implementa CST 06 (Alíquota Zero)
type COFINS06 struct{}

func (s *COFINS06) Calculate(item models.ItemDocumentoFiscalEntrada, info COFINSInfo) decimal.Decimal {
	return decimal.Zero
}

// COFINS49 implementa CST 49 (Outras Operações de Saída)
type COFINS49 struct{}

func (s *COFINS49) Calculate(item models.ItemDocumentoFiscalEntrada, info COFINSInfo) decimal.Decimal {
	return decimal.Zero
}

// COFINS50To99 implementa CSTs 50-99 (operações de crédito, suspensão, outras)
type COFINS50To99 struct{}

func (s *COFINS50To99) Calculate(item models.ItemDocumentoFiscalEntrada, info COFINSInfo) decimal.Decimal {
	return decimal.Zero
}
