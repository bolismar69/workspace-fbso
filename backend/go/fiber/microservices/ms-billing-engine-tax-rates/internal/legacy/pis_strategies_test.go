package legacy

import (
	"testing"

	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

func TestPIS01_02_Calculate(t *testing.T) {
	s := &PIS01_02{}

	tests := []struct {
		name     string
		item     models.ItemDocumentoFiscalEntrada
		info     PISInfo
		expected string
	}{
		{
			name: "CST 01 — tributação normal sem ICMS",
			item: models.ItemDocumentoFiscalEntrada{
				Quantidade:    decimal.NewFromInt(1),
				ValorUnitario: decimal.NewFromInt(1000),
			},
			info: PISInfo{
				CST:       "01",
				Aliquota:  decimal.NewFromFloat(1.65),
				ValorIcms: decimal.Zero,
			},
			expected: "16.50",
		},
		{
			name: "CST 01 — com exclusão do ICMS da base",
			item: models.ItemDocumentoFiscalEntrada{
				Quantidade:    decimal.NewFromInt(1),
				ValorUnitario: decimal.NewFromInt(1000),
			},
			info: PISInfo{
				CST:       "01",
				Aliquota:  decimal.NewFromFloat(1.65),
				ValorIcms: decimal.NewFromInt(120),
			},
			expected: "14.52",
		},
		{
			name: "CST 02 — alíquota diferenciada",
			item: models.ItemDocumentoFiscalEntrada{
				Quantidade:    decimal.NewFromInt(10),
				ValorUnitario: decimal.NewFromInt(100),
			},
			info: PISInfo{
				CST:       "02",
				Aliquota:  decimal.NewFromFloat(0.65),
				ValorIcms: decimal.NewFromInt(180),
			},
			expected: "5.33",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.Calculate(tt.item, tt.info)
			gotStr := got.StringFixed(2)
			if gotStr != tt.expected {
				t.Errorf("PIS01_02.Calculate() = %s, want %s", gotStr, tt.expected)
			}
		})
	}
}

func TestPIS03_Calculate(t *testing.T) {
	s := &PIS03{}

	tests := []struct {
		name     string
		item     models.ItemDocumentoFiscalEntrada
		info     PISInfo
		expected string
	}{
		{
			name: "CST 03 — alíquota por unidade (CalcTax divide por 100)",
			item: models.ItemDocumentoFiscalEntrada{
				Quantidade: decimal.NewFromInt(100),
			},
			info: PISInfo{
				CST:      "03",
				Aliquota: decimal.NewFromFloat(50),
			},
			expected: "50.00",
		},
		{
			name: "CST 03 — alíquota por unidade fracionada",
			item: models.ItemDocumentoFiscalEntrada{
				Quantidade: decimal.NewFromFloat(150.5),
			},
			info: PISInfo{
				CST:      "03",
				Aliquota: decimal.NewFromFloat(30),
			},
			expected: "45.15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.Calculate(tt.item, tt.info)
			gotStr := got.StringFixed(2)
			if gotStr != tt.expected {
				t.Errorf("PIS03.Calculate() = %s, want %s", gotStr, tt.expected)
			}
		})
	}
}

func TestPISStrategiesZero(t *testing.T) {
	item := models.ItemDocumentoFiscalEntrada{
		Quantidade:    decimal.NewFromInt(1),
		ValorUnitario: decimal.NewFromInt(1000),
	}
	info := PISInfo{
		Aliquota:  decimal.NewFromFloat(1.65),
		ValorIcms: decimal.NewFromInt(120),
	}

	tests := []struct {
		name     string
		strategy PISStrategy
	}{
		{"PIS04 — Monofásico", &PIS04{}},
		{"PIS05 — Substituição Tributária", &PIS05{}},
		{"PIS06 — Alíquota Zero", &PIS06{}},
		{"PIS49 — Outras Operações de Saída", &PIS49{}},
		{"PIS50To99 — Crédito/Suspensão/Outras", &PIS50To99{}},
		{"PIS99 — Outras Operações", &PIS99{}},
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

func TestGetPISStrategy(t *testing.T) {
	tests := []struct {
		cst      string
		strategy string
	}{
		{"01", "*legacy.PIS01_02"},
		{"02", "*legacy.PIS01_02"},
		{"03", "*legacy.PIS03"},
		{"04", "*legacy.PIS04"},
		{"05", "*legacy.PIS05"},
		{"06", "*legacy.PIS06"},
		{"49", "*legacy.PIS49"},
		{"99", "*legacy.PIS99"},
		{"50", "*legacy.PIS50To99"},
		{"70", "*legacy.PIS50To99"},
		{"98", "*legacy.PIS50To99"},
		{"", "*legacy.PIS50To99"},
	}

	for _, tt := range tests {
		t.Run("CST "+tt.cst, func(t *testing.T) {
			got := GetPISStrategy(tt.cst)
			gotType := getPISTypeName(got)
			if gotType != tt.strategy {
				t.Errorf("GetPISStrategy(%q) = %s, want %s", tt.cst, gotType, tt.strategy)
			}
		})
	}
}

func getPISTypeName(i interface{}) string {
	switch i.(type) {
	case *PIS01_02:
		return "*legacy.PIS01_02"
	case *PIS03:
		return "*legacy.PIS03"
	case *PIS04:
		return "*legacy.PIS04"
	case *PIS05:
		return "*legacy.PIS05"
	case *PIS06:
		return "*legacy.PIS06"
	case *PIS49:
		return "*legacy.PIS49"
	case *PIS99:
		return "*legacy.PIS99"
	case *PIS50To99:
		return "*legacy.PIS50To99"
	default:
		return "unknown"
	}
}
