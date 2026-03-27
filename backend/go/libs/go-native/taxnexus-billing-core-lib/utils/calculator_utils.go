package utils

import (
	"github.com/shopspring/decimal"
)

// CalcTax realiza o cálculo de imposto com precisão decimal
func CalcTax(base float64, rate float64) float64 {
	b := decimal.NewFromFloat(base)
	r := decimal.NewFromFloat(rate)
	
	// Calcula e arredonda para 2 casas decimais (Round Half Up)
	result, _ := b.Mul(r).Round(2).Float64()
	return result
}

// SumResults soma uma lista de valores monetários com precisão
func SumResults(values ...float64) float64 {
	total := decimal.Zero
	for _, v := range values {
		total = total.Add(decimal.NewFromFloat(v))
	}
	res, _ := total.Float64()
	return res
}
