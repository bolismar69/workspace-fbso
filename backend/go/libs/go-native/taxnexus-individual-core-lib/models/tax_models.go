// path: backend/go/libs/go-native/taxnexus-individual-core-lib/models/tax_models.go
package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TaxDefinition struct {
	ID                int    `db:"id" json:"id"`
	TaxCode           string `db:"tax_code" json:"tax_code"`
	Name              string `db:"name" json:"name"`
	Sphere            string `db:"sphere" json:"sphere"`
	RoundingPrecision int    `db:"rounding_precision" json:"rounding_precision"`
	Active            bool   `db:"active" json:"active"`
}

type TaxRule struct {
	ID              int              `db:"id" json:"id"`
	TaxDefinitionID int              `db:"tax_definition_id" json:"tax_definition_id"`
	Description     string           `db:"description" json:"description"`
	RangeMin        decimal.Decimal  `db:"range_min" json:"range_min"`
	RangeMax        *decimal.Decimal `db:"range_max" json:"range_max"` // Pointer para lidar com NULL
	AliqPercent     decimal.Decimal  `db:"aliq_percent" json:"aliq_percent"`
	DeductionVal    decimal.Decimal  `db:"deduction_val" json:"deduction_val"`
	ValidFrom       time.Time        `db:"valid_from" json:"valid_from"`
	ValidTo         *time.Time       `db:"valid_to" json:"valid_to"`
}

type TaxCalculationLog struct {
	ID               uuid.UUID       `db:"id" json:"id"`
	CalculationDate  time.Time       `db:"calculation_date" json:"calculation_date"`
	TaxCode          string          `db:"tax_code" json:"tax_code"`
	InputBaseValue   decimal.Decimal `db:"input_base_value" json:"input_base_value"`
	CalculatedAmount decimal.Decimal `db:"calculated_amount" json:"calculated_amount"`
	AppliedAliq      decimal.Decimal `db:"applied_aliq" json:"applied_aliq"`
	TraceID          string          `db:"trace_id" json:"trace_id"`
}

// TaxRequest representa o payload de entrada para o simulador
type TaxRequest struct {
	Type                       string          `json:"calculation_type"` // "monthly" ou "annual"
	RefDate                    time.Time       `json:"reference_date"`
	GrossIncome                decimal.Decimal `json:"gross_income"`
	PensionDeductionAmount     decimal.Decimal `json:"pension_deduction_amount"`     // Valor fixo
	PensionDeductionPercentage decimal.Decimal `json:"pension_deduction_percentage"` // Percentual (ex: 11.0)
	DependentsCount            int             `json:"dependents_count"`
	EducationExpenses          decimal.Decimal `json:"education_expenses"` // Gasto real com educação
	HealthExpenses             decimal.Decimal `json:"health_expenses"`    // Gasto real com saúde
	PGBLContribution           decimal.Decimal `json:"pgbl_contribution"`  // Contribuição PGBL
	OtherDeductions            decimal.Decimal `json:"other_deductions"`
}

type DocumentoFiscalRequest struct {
	Key   string          `json:"key"`
	Value decimal.Decimal `json:"value"`
	Unit  string          `json:"unit"`
}

type UniversalTaxRequest struct {
	TaxCode         string                   `json:"tax_code"`
	CalculationType string                   `json:"calculation_type"`
	ReferenceDate   time.Time                `json:"reference_date"`
	GrossIncome     decimal.Decimal          `json:"gross_income"`
	Inputs          []DocumentoFiscalRequest `json:"inputs"`
}

// DeductionDetail para o detalhamento na saída
type DeductionDetail struct {
	Type   string          `json:"type"`
	Amount decimal.Decimal `json:"amount"`
}

// TaxResponse representa o resultado do cálculo
type TaxResponse struct {
	ReferenceDate        string            `json:"reference_date"`
	GrossIncome          decimal.Decimal   `json:"gross_income"`
	TotalDeductionAmount decimal.Decimal   `json:"total_deduction_amount"`
	BaseValue            decimal.Decimal   `json:"base_value"`
	TaxAmount            decimal.Decimal   `json:"tax_amount"`
	EffectiveRate        decimal.Decimal   `json:"effective_rate"`
	IsRecommended        bool              `json:"is_recommended"`
	DeductionDetails     []DeductionDetail `json:"deduction_details"`
	AppliedRule          string            `json:"applied_rule_summary"`
	UsedConfigs          map[string]string `json:"used_configs"`
}
