// path: backend/go/libs/go-native/taxnexus-billing-core-lib/utils/money.go
package utils

import "github.com/shopspring/decimal"

// RoundTax aplica a regra de arredondamento fiscal padrão (2 casas decimais)
func RoundTax(value float64) float64 {
    d := decimal.NewFromFloat(value)
    res, _ := d.Round(2).Float64()
    return res
}

// Percent aplica uma alíquota sobre uma base com precisão
func Percent(base float64, rate float64) float64 {
    b := decimal.NewFromFloat(base)
    r := decimal.NewFromFloat(rate)
    res, _ := b.Mul(r).Round(2).Float64()
    return res
}
