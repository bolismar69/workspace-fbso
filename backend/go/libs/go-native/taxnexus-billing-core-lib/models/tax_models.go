// path: backend/go/libs/go-native/taxnexus-billing-core-lib/models/tax_models.go
package models

import (
	"time"
	"github.com/shopspring/decimal"
	"github.com/google/uuid"
)

// TaxInput representa o fato gerador (NF-e, NFC-e, etc)
type TaxInput struct {
    ID             string    `json:"id"`
    DataOperacao   time.Time `json:"data_operacao"`
    
    // Geografia (Crucial para ICMS e o novo IBS)
    Origem         TaxLocation  `json:"origem"`
    Destino        TaxLocation  `json:"destino"`
    
    // Dados da Empresa
    Regime         string    `json:"regime_tributario"` // Simples, Presumido, Real
    CNAE           string    `json:"cnae"`

    // valor bruto
    // ValorBruto     decimal.Decimal `json:"valor_bruto"`
    ValorBruto     float64   `json:"valor_bruto"`
    
    // Itens da Nota
    Itens          []TaxItem `json:"itens"`
}

type TaxLocation struct {
    UF       string `json:"uf"`
    Municipio string `json:"municipio_codigo_ibge"` // Essencial para o IBS Municipal
}

type TaxItem struct {
	SKU      string  `json:"sku"`
	NCM      string  `json:"ncm"`
	Valor    float64 `json:"valor"`
	Quantity float64 `json:"quantity"`
}

// TaxResult é o padrão de saída para qualquer cálculo (Legacy ou Reforma)
type TaxResult struct {
	TaxName    string  `json:"tax_name"`
	BaseCalculo float64 `json:"base_calculo"`
	Aliquota    float64 `json:"aliquota"`
	Valor      float64 `json:"valor"`
}
