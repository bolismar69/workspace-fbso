package legacy

import (
	"testing"

	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

func TestCOFINS01_02_Calculate(t *testing.T) {
	s := &COFINS01_02{}

	tests := []struct {
		name     string
		item     models.ItemDocumentoFiscalEntrada
		info     COFINSInfo
		expected string
	}{
		{
			name: "CST 01 — tributação normal sem ICMS",
			item: models.ItemDocumentoFiscalEntrada{
				Quantidade:    decimal.NewFromInt(1),
				ValorUnitario: decimal.NewFromInt(1000),
			},
			info: COFINSInfo{
				CST:       "01",
				Aliquota:  decimal.NewFromFloat(7.6),
				ValorIcms: decimal.Zero,
			},
			expected: "76.00",
		},
		{
			name: "CST 01 — com exclusão do ICMS da base",
			item: models.ItemDocumentoFiscalEntrada{
				Quantidade:    decimal.NewFromInt(1),
				ValorUnitario: decimal.NewFromInt(1000),
			},
			info: COFINSInfo{
				CST:       "01",
				Aliquota:  decimal.NewFromFloat(7.6),
				ValorIcms: decimal.NewFromInt(120),
			},
			expected: "66.88",
		},
		{
			name: "CST 02 — alíquota diferenciada",
			item: models.ItemDocumentoFiscalEntrada{
				Quantidade:    decimal.NewFromInt(10),
				ValorUnitario: decimal.NewFromInt(100),
			},
			info: COFINSInfo{
				CST:       "02",
				Aliquota:  decimal.NewFromFloat(3.0),
				ValorIcms: decimal.NewFromInt(180),
			},
			expected: "24.60",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.Calculate(tt.item, tt.info)
			gotStr := got.StringFixed(2)
			if gotStr != tt.expected {
				t.Errorf("COFINS01_02.Calculate() = %s, want %s", gotStr, tt.expected)
			}
		})
	}
}

func TestCOFINS03_Calculate(t *testing.T) {
	s := &COFINS03{}

	tests := []struct {
		name     string
		item     models.ItemDocumentoFiscalEntrada
		info     COFINSInfo
		expected string
	}{
		{
			name: "CST 03 — alíquota por unidade (CalcTax divide por 100)",
			item: models.ItemDocumentoFiscalEntrada{
				Quantidade: decimal.NewFromInt(100),
			},
			info: COFINSInfo{
				CST:      "03",
				Aliquota: decimal.NewFromFloat(200),
			},
			expected: "200.00",
		},
		{
			name: "CST 03 — alíquota por unidade fracionada",
			item: models.ItemDocumentoFiscalEntrada{
				Quantidade: decimal.NewFromFloat(75.3),
			},
			info: COFINSInfo{
				CST:      "03",
				Aliquota: decimal.NewFromFloat(150),
			},
			expected: "112.95",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.Calculate(tt.item, tt.info)
			gotStr := got.StringFixed(2)
			if gotStr != tt.expected {
				t.Errorf("COFINS03.Calculate() = %s, want %s", gotStr, tt.expected)
			}
		})
	}
}

func TestCOFINSStrategiesZero(t *testing.T) {
	item := models.ItemDocumentoFiscalEntrada{
		Quantidade:    decimal.NewFromInt(1),
		ValorUnitario: decimal.NewFromInt(1000),
	}
	info := COFINSInfo{
		Aliquota:  decimal.NewFromFloat(7.6),
		ValorIcms: decimal.NewFromInt(120),
	}

	tests := []struct {
		name     string
		strategy COFINSStrategy
	}{
		{"COFINS04 — Monofásico", &COFINS04{}},
		{"COFINS05 — Substituição Tributária", &COFINS05{}},
		{"COFINS06 — Alíquota Zero", &COFINS06{}},
		{"COFINS49 — Outras Operações de Saída", &COFINS49{}},
		{"COFINS50To99 — Crédito/Suspensão/Outras", &COFINS50To99{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.strategy.Calculate(item, info)
			if !got.Equal(decimal.Zero) {
				t.Errorf("%s.Calculate() = %s, want 0", tt.name, got.String())
			}
		})
	}
}

func TestGetCOFINSStrategy(t *testing.T) {
	tests := []struct {
		cst      string
		strategy string
	}{
		{"01", "*legacy.COFINS01_02"},
		{"02", "*legacy.COFINS01_02"},
		{"03", "*legacy.COFINS03"},
		{"04", "*legacy.COFINS04"},
		{"05", "*legacy.COFINS05"},
		{"06", "*legacy.COFINS06"},
		{"49", "*legacy.COFINS49"},
		{"50", "*legacy.COFINS50To99"},
		{"70", "*legacy.COFINS50To99"},
		{"98", "*legacy.COFINS50To99"},
		{"", "*legacy.COFINS50To99"},
	}

	for _, tt := range tests {
		t.Run("CST "+tt.cst, func(t *testing.T) {
			got := GetCOFINSStrategy(tt.cst)
			gotType := getCOFINSTypeName(got)
			if gotType != tt.strategy {
				t.Errorf("GetCOFINSStrategy(%q) = %s, want %s", tt.cst, gotType, tt.strategy)
			}
		})
	}
}

func getCOFINSTypeName(i interface{}) string {
	switch i.(type) {
	case *COFINS01_02:
		return "*legacy.COFINS01_02"
	case *COFINS03:
		return "*legacy.COFINS03"
	case *COFINS04:
		return "*legacy.COFINS04"
	case *COFINS05:
		return "*legacy.COFINS05"
	case *COFINS06:
		return "*legacy.COFINS06"
	case *COFINS49:
		return "*legacy.COFINS49"
	case *COFINS50To99:
		return "*legacy.COFINS50To99"
	default:
		return "unknown"
	}
}
