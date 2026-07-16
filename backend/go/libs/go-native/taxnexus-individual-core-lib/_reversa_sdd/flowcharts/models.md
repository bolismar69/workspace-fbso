# Modelo de dados — módulo `models`

> Gerado pelo **Arqueólogo** (Reversa) em 2026-06-10
> Fonte: `models/tax_models.go` · Sem fluxo de controle (só definições de tipo)

## Relacionamento entre structs / tabelas

```mermaid
classDiagram
    class TaxDefinition {
        +int ID
        +string TaxCode
        +string Name
        +string Sphere
        +int RoundingPrecision
        +bool Active
    }
    class TaxRule {
        +int ID
        +int TaxDefinitionID
        +string Description
        +Decimal RangeMin
        +Decimal* RangeMax (NULL)
        +Decimal AliqPercent
        +Decimal DeductionVal
        +Time ValidFrom
        +Time* ValidTo (NULL)
    }
    class TaxCalculationLog {
        +UUID ID
        +Time CalculationDate
        +string TaxCode
        +Decimal InputBaseValue
        +Decimal CalculatedAmount
        +Decimal AppliedAliq
        +string TraceID
    }
    class TaxRequest {
        +string Type (monthly/annual)
        +Time RefDate
        +Decimal GrossIncome
        +Decimal PensionDeductionAmount
        +Decimal PensionDeductionPercentage
        +int DependentsCount
        +Decimal EducationExpenses
        +Decimal HealthExpenses
        +Decimal PGBLContribution
        +Decimal OtherDeductions
    }
    class UniversalTaxRequest {
        +string TaxCode
        +string CalculationType
        +Time ReferenceDate
        +Decimal GrossIncome
        +documentoFiscalRequest[] Inputs
    }
    class documentoFiscalRequest {
        +string Key
        +Decimal Value
        +string Unit
    }
    class TaxResponse {
        +string ReferenceDate
        +Decimal GrossIncome
        +Decimal TotalDeductionAmount
        +Decimal BaseValue
        +Decimal TaxAmount
        +Decimal EffectiveRate
        +bool IsRecommended
        +DeductionDetail[] DeductionDetails
        +string AppliedRule
        +map UsedConfigs
    }
    class DeductionDetail {
        +string Type
        +Decimal Amount
    }

    TaxDefinition "1" --> "N" TaxRule : tax_definition_id
    UniversalTaxRequest "1" --> "N" documentoFiscalRequest : Inputs
    TaxResponse "1" --> "N" DeductionDetail : DeductionDetails
```

> Legenda: `Decimal*` / `Time*` = ponteiro = coluna NULL-ável.
> `TaxRequest`/`UniversalTaxRequest` (entrada) e `TaxResponse` (saída) são consumidos pela camada de **cálculo ausente** deste recorte (🔴 LACUNA).
