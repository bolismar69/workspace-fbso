# Dicionário de Dados — taxnexus-individual-core-lib

> Gerado pelo **Arqueólogo** (Reversa) em 2026-06-10 · `doc_level = completo`
> Fonte: `models/tax_models.go` + queries SQL em `repository/tax_repository.go`
> Confiança: 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

---

## Entidades persistidas

### `TaxDefinition` — tabela `individual_tax_rates.tax_definitions` 🟢

Definição de um imposto/tributo.

| Campo (Go) | Coluna (db) | Tipo Go | Obrigatório | Notas |
|------------|-------------|---------|-------------|-------|
| `ID` | `id` | `int` | sim | PK 🟡 |
| `TaxCode` | `tax_code` | `string` | sim | Chave de negócio (ex.: `IRPF`, `INSS`) 🟡; usada como filtro nas queries 🟢 |
| `Name` | `name` | `string` | sim | Nome legível do imposto |
| `Sphere` | `sphere` | `string` | sim | Esfera (federal/estadual/municipal) 🟡 |
| `RoundingPrecision` | `rounding_precision` | `int` | sim | Casas decimais para arredondamento do cálculo 🟡 |
| `Active` | `active` | `bool` | sim | Flag de ativação |

> 🟡 Apenas `id` e `tax_code` são confirmados pela query (JOIN em `GetTaxRulesForPeriod`); os demais vêm da struct.

---

### `TaxRule` — tabela `individual_tax_rates.tax_rules_history` 🟢

Faixa/escalão de imposto com vigência temporal (modelo historizado).

| Campo (Go) | Coluna (db) | Tipo Go | Obrigatório | Notas |
|------------|-------------|---------|-------------|-------|
| `ID` | `id` | `int` | sim | PK 🟡 |
| `TaxDefinitionID` | `tax_definition_id` | `int` | sim | FK → `tax_definitions.id` 🟢 (JOIN) |
| `Description` | `description` | `string` | sim | Descrição da faixa |
| `RangeMin` | `range_min` | `decimal.Decimal` | sim | Limite inferior do intervalo (inclusivo) 🟢 |
| `RangeMax` | `range_max` | `*decimal.Decimal` | **não (NULL)** | Limite superior inclusivo; `NULL` = faixa aberta (sem teto) 🟢 |
| `AliqPercent` | `aliq_percent` | `decimal.Decimal` | sim | Alíquota (%) da faixa |
| `DeductionVal` | `deduction_val` | `decimal.Decimal` | sim | Parcela a deduzir da faixa (tabela progressiva IRPF) 🟡 |
| `ValidFrom` | `valid_from` | `time.Time` | sim | Início de vigência (inclusivo) 🟢 |
| `ValidTo` | `valid_to` | `*time.Time` | **não (NULL)** | Fim de vigência; `NULL` = vigente indefinidamente 🟢 |

---

### `TaxCalculationLog` — tabela `tax_calculation_log` 🟡 🔴

Log de auditoria de um cálculo executado. **Inferida** — nenhuma query neste repositório a lê ou escreve.

| Campo (Go) | Coluna (db) | Tipo Go | Obrigatório | Notas |
|------------|-------------|---------|-------------|-------|
| `ID` | `id` | `uuid.UUID` | sim | PK (UUID gerado na aplicação) 🟡 |
| `CalculationDate` | `calculation_date` | `time.Time` | sim | Quando o cálculo ocorreu |
| `TaxCode` | `tax_code` | `string` | sim | Imposto calculado |
| `InputBaseValue` | `input_base_value` | `decimal.Decimal` | sim | Base de cálculo de entrada |
| `CalculatedAmount` | `calculated_amount` | `decimal.Decimal` | sim | Valor de imposto calculado |
| `AppliedAliq` | `applied_aliq` | `decimal.Decimal` | sim | Alíquota aplicada |
| `TraceID` | `trace_id` | `string` | sim | Correlação para observabilidade 🟡 |

> 🔴 **LACUNA:** existência da tabela e sua escrita dependem da camada de cálculo ausente.

---

## Configuração (não modelada como struct) 🟢

### tabela `individual_tax_rates.tax_configs`

Parâmetros chave-valor por imposto, com vigência. Lida em `GetConfig` e `GetTableConfigs`.

| Coluna | Tipo (inferido) | Obrigatório | Notas |
|--------|-----------------|-------------|-------|
| `tax_code` | text | sim | Imposto dono do parâmetro 🟢 |
| `config_key` | text | sim | Nome do parâmetro (ex.: limites de dedução, valor por dependente) 🟡 |
| `config_value` | numeric → `decimal.Decimal` | sim | Valor do parâmetro 🟢 |
| `valid_from` | date/timestamp | sim | Início de vigência 🟢 |
| `valid_to` | date/timestamp | **não (NULL)** | Fim de vigência; `NULL` = vigente 🟢 |

> Retorno de `GetTableConfigs`: `map[string]decimal.Decimal` (config_key → config_value). 🟢

---

## DTOs de transporte (não persistidos)

### `TaxRequest` 🟢 — payload do simulador (deduções IRPF)

| Campo (Go) | JSON | Tipo | Notas |
|------------|------|------|-------|
| `Type` | `calculation_type` | `string` | `"monthly"` ou `"annual"` 🟢 (comentário) |
| `RefDate` | `reference_date` | `time.Time` | Data de referência (define vigência aplicável) |
| `GrossIncome` | `gross_income` | `decimal.Decimal` | Renda bruta |
| `PensionDeductionAmount` | `pension_deduction_amount` | `decimal.Decimal` | Dedução previdenciária — valor fixo |
| `PensionDeductionPercentage` | `pension_deduction_percentage` | `decimal.Decimal` | Dedução previdenciária — percentual (ex.: 11.0) |
| `DependentsCount` | `dependents_count` | `int` | Nº de dependentes |
| `EducationExpenses` | `education_expenses` | `decimal.Decimal` | Gasto real com educação |
| `HealthExpenses` | `health_expenses` | `decimal.Decimal` | Gasto real com saúde |
| `PGBLContribution` | `pgbl_contribution` | `decimal.Decimal` | Contribuição PGBL |
| `OtherDeductions` | `other_deductions` | `decimal.Decimal` | Outras deduções |

### `UniversalTaxRequest` 🟢 — payload genérico

| Campo (Go) | JSON | Tipo | Notas |
|------------|------|------|-------|
| `TaxCode` | `tax_code` | `string` | Imposto a calcular |
| `CalculationType` | `calculation_type` | `string` | Tipo de cálculo |
| `ReferenceDate` | `reference_date` | `time.Time` | Data de referência |
| `GrossIncome` | `gross_income` | `decimal.Decimal` | Renda bruta |
| `Inputs` | `inputs` | `[]documentoFiscalRequest` | Lista de entradas genéricas |

### `documentoFiscalRequest` 🟢 (não exportada)

| Campo (Go) | JSON | Tipo |
|------------|------|------|
| `Key` | `key` | `string` |
| `Value` | `value` | `decimal.Decimal` |
| `Unit` | `unit` | `string` |

### `DeductionDetail` 🟢

| Campo (Go) | JSON | Tipo |
|------------|------|------|
| `Type` | `type` | `string` |
| `Amount` | `amount` | `decimal.Decimal` |

### `TaxResponse` 🟢 — resultado do cálculo

| Campo (Go) | JSON | Tipo | Notas |
|------------|------|------|-------|
| `ReferenceDate` | `reference_date` | `string` | |
| `GrossIncome` | `gross_income` | `decimal.Decimal` | |
| `TotalDeductionAmount` | `total_deduction_amount` | `decimal.Decimal` | |
| `BaseValue` | `base_value` | `decimal.Decimal` | Base após deduções |
| `TaxAmount` | `tax_amount` | `decimal.Decimal` | Imposto devido |
| `EffectiveRate` | `effective_rate` | `decimal.Decimal` | Alíquota efetiva |
| `IsRecommended` | `is_recommended` | `bool` | Marca o cenário recomendado 🟡 |
| `DeductionDetails` | `deduction_details` | `[]DeductionDetail` | Detalhamento de deduções |
| `AppliedRule` | `applied_rule_summary` | `string` | Resumo da regra aplicada |
| `UsedConfigs` | `used_configs` | `map[string]string` | Configs usadas no cálculo (rastreabilidade) 🟡 |

---

## Relacionamentos

```
tax_definitions (1) ──< (N) tax_rules_history   [via tax_definition_id]   🟢
tax_definitions.tax_code  ──  tax_configs.tax_code   [associação por código] 🟡
tax_definitions.tax_code  ──  tax_calculation_log.tax_code   [associação por código] 🟡🔴
```
