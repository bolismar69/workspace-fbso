# Legacy Mapping — módulo `models`

> Gerado pelo **Arqueólogo** (Reversa) em 2026-06-10

| Struct | Arquivo:linha | Persistida em | Confiança |
|--------|---------------|---------------|-----------|
| `TaxDefinition` | `models/tax_models.go:11` | `tax_definitions` | 🟢 |
| `TaxRule` | `models/tax_models.go:20` | `tax_rules_history` | 🟢 |
| `TaxCalculationLog` | `models/tax_models.go:32` | `tax_calculation_log` (inferida) | 🟡🔴 |
| `TaxRequest` | `models/tax_models.go:43` | — (DTO entrada) | 🟢 |
| `documentoFiscalRequest` | `models/tax_models.go:56` | — (DTO interno, não exportado) | 🟢 |
| `UniversalTaxRequest` | `models/tax_models.go:62` | — (DTO entrada) | 🟢 |
| `DeductionDetail` | `models/tax_models.go:71` | — (DTO saída) | 🟢 |
| `TaxResponse` | `models/tax_models.go:77` | — (DTO saída) | 🟢 |

Imports relevantes: `time`, `github.com/google/uuid`, `github.com/shopspring/decimal` (linhas 4–9).
