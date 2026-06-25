// path: backend/go/libs/go-native/taxnexus-billing-core-lib/utils/calculator_utils.go
package utils

import (
	"github.com/shopspring/decimal"
)

// RoundDecimal arredonda um valor para a precisão desejada usando o padrão contábil (Half Away From Zero).
func RoundDecimal(val decimal.Decimal, precision int32) decimal.Decimal {
	return val.Round(precision)
}

// RoundDecimalWithUintPrecision arredonda um valor decimal quando a precisão é informada como uint.
func RoundDecimalWithUintPrecision(val decimal.Decimal, precision uint) decimal.Decimal {
	return RoundDecimal(val, int32(precision))
}

// CalcTaxPercentDecimal executa o cálculo de imposto percentual com precisão decimal.
func CalcTaxPercentDecimal(base decimal.Decimal, aliquota decimal.Decimal) decimal.Decimal {
	return CalcTax(base, aliquota)
}

// CalcTax realiza o cálculo de imposto: (Base * Aliquota) / 100
func CalcTax(base decimal.Decimal, aliquota decimal.Decimal) decimal.Decimal {
	// Multiplica base pela alíquota e divide por 100, arredondando para 2 casas
	valor := base.Mul(aliquota).Div(decimal.NewFromInt(100))
	return RoundDecimal(valor, 2)
}

// SumResults soma uma lista de valores monetários com precisão
func SumResults(values ...decimal.Decimal) decimal.Decimal {
	total := decimal.Zero
	for _, v := range values {
		total = total.Add(v)
	}
	return total
}

// CalcularAliquotaEfetivaSimples calcula a taxa que será usada para o crédito de ICMS.
// Formula: ((RBT12 * AliquotaNominal) - ParcelaDeduzir) / RBT12
// O resultado final é multiplicado pelo % de ICMS do anexo.
func CalcularAliquotaEfetivaSimples(rbt12, aliqNominal, parcelaDeduzir, percIcmsAnexo decimal.Decimal) decimal.Decimal {
	if rbt12.IsZero() {
		return decimal.Zero
	}

	// Alíquota Efetiva = (RBT12 * AliqNominal - ValorDeduzir) / RBT12
	v1 := rbt12.Mul(aliqNominal)
	v2 := v1.Sub(parcelaDeduzir)
	aliqEfetiva := v2.Div(rbt12)

	// O resultado deve ser multiplicado pelo percentual que o ICMS representa no anexo
	return RoundDecimal(aliqEfetiva.Mul(percIcmsAnexo), 4) // 4 casas para alíquotas é comum
}
