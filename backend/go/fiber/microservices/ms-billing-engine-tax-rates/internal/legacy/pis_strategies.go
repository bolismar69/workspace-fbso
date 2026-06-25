// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/legacy/pis_strategies.go
package legacy

import (
	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/utils"

	"github.com/shopspring/decimal"
)

// PISStrategy define o contrato para cada tipo de CST de PIS
type PISStrategy interface {
	Calculate(item models.ItemDocumentoFiscalEntrada, info PISInfo) decimal.Decimal
}

// PISInfo carrega dados extras necessários para o cálculo
type PISInfo struct {
	CST       string
	Regime    string
	Aliquota  decimal.Decimal
	ValorIcms decimal.Decimal // Valor do ICMS Próprio para exclusão da base
}

// PIS01_02 implementa CST 01 (Básica) e 02 (Diferenciada) - Ad Valorem
type PIS01_02 struct{}

func (s *PIS01_02) Calculate(item models.ItemDocumentoFiscalEntrada, info PISInfo) decimal.Decimal {
	vItem := item.ValorUnitario.Mul(item.Quantidade)
	vIcms := info.ValorIcms
	aliq := info.Aliquota.Div(decimal.NewFromInt(100))

	// Base = Valor Item - ICMS
	base := vItem.Sub(vIcms)
	if base.IsNegative() {
		base = decimal.Zero
	}

	// Valor Final = (Base * Alíquota) arredondado para 2 casas
	return base.Mul(aliq).Round(2)
}

// PIS03 implementa CST 03 - Alíquota por Unidade (Quantidade)
// Nota: CST 03 geralmente não sofre impacto da exclusão do ICMS da base,
// pois a base é quantidade, não valor. Mantemos por integridade.
type PIS03 struct{}

func (s *PIS03) Calculate(item models.ItemDocumentoFiscalEntrada, info PISInfo) decimal.Decimal {
	// CalcTax retorna decimal.Decimal
	return utils.CalcTax(item.Quantidade, info.Aliquota).Round(2)
}

// PIS04 implementa CST 04 (Monofásico) — tributo concentrado no produtor/importador
type PIS04 struct{}

func (s *PIS04) Calculate(item models.ItemDocumentoFiscalEntrada, info PISInfo) decimal.Decimal {
	return decimal.Zero
}

// PIS05 implementa CST 05 (Substituição Tributária) — tributo já recolhido por substituição
type PIS05 struct{}

func (s *PIS05) Calculate(item models.ItemDocumentoFiscalEntrada, info PISInfo) decimal.Decimal {
	return decimal.Zero
}

// PIS06 implementa CST 06 (Alíquota Zero)
type PIS06 struct{}

func (s *PIS06) Calculate(item models.ItemDocumentoFiscalEntrada, info PISInfo) decimal.Decimal {
	return decimal.Zero
}

// PIS49 implementa CST 49 (Outras Operações de Saída)
type PIS49 struct{}

func (s *PIS49) Calculate(item models.ItemDocumentoFiscalEntrada, info PISInfo) decimal.Decimal {
	return decimal.Zero
}

// PIS50To99 implementa CSTs 50-99 (operações de crédito, suspensão, outras)
type PIS50To99 struct{}

func (s *PIS50To99) Calculate(item models.ItemDocumentoFiscalEntrada, info PISInfo) decimal.Decimal {
	return decimal.Zero
}

// PIS99 implementa CST 99 (Outras Operações)
type PIS99 struct{}

func (s *PIS99) Calculate(item models.ItemDocumentoFiscalEntrada, info PISInfo) decimal.Decimal {
	return decimal.Zero
}
