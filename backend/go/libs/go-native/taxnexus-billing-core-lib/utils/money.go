// path: backend/go/libs/go-native/taxnexus-billing-core-lib/utils/money.go
package utils

import "github.com/shopspring/decimal"

// RoundTax aplica a regra de arredondamento fiscal padrão (2 casas decimais)
func RoundTax(value decimal.Decimal) decimal.Decimal {
	return RoundDecimal(value, 2)
}

// Percent aplica uma alíquota sobre uma base com precisão
func Percent(base decimal.Decimal, rate decimal.Decimal) decimal.Decimal {
	return RoundDecimal(base.Mul(rate), 2)
}
