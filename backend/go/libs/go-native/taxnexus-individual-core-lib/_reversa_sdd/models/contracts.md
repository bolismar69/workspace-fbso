# models — Contratos de Dados

> Contrato de fronteira: mapeamento campo a campo Go ↔ banco (`db:`) e Go ↔ wire JSON (`json:`).
> Esta é a base para a documentação dos serviços consumidores (A1).
> Gerado pelo **Redator** (Reversa) em 2026-06-10 · `doc_level = completo`
> Fontes: `models/tax_models.go`, `data-dictionary.md`, DDL (A3), exemplos de chamada (D2).
> 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

---

## 1. Entidades persistidas (contrato `db:` ↔ schema `individual_tax_rates`)

### `TaxDefinition` ↔ `tax_definitions` 🟢

| Campo Go | Coluna | Tipo db (DDL) | Nulável | Notas |
|----------|--------|---------------|---------|-------|
| `ID` | `id` | `serial4` (PK) | não | |
| `TaxCode` | `tax_code` | `varchar(20)` | não | **UNIQUE** (`tax_definitions_tax_code_key`); valores: `IRPF`, `INSS` (D1) |
| `Name` | `name` | `varchar(100)` | não | |
| `Sphere` | `sphere` | `bpchar(1)` | não | 1 caractere (ex.: esfera federal/estadual/municipal) 🟡 |
| `RoundingPrecision` | `rounding_precision` | `int4` (default 2) | sim | Casas decimais de arredondamento 🟡 |
| `Active` | `active` | `bool` (default true) | sim | Filtro **não** aplicado nesta lib — responsabilidade do consumidor (D5) |
| — | `created_at` | `timestamptz` (default now) | sim | Não mapeado na struct 🟡 |

### `TaxRule` ↔ `tax_rules_history` 🟢

| Campo Go | Coluna | Tipo db (DDL) | Nulável | Notas |
|----------|--------|---------------|---------|-------|
| `ID` | `id` | `serial4` (PK) | não | |
| `TaxDefinitionID` | `tax_definition_id` | `int4` | sim | **FK** → `tax_definitions(id)` |
| `Description` | `description` | `varchar(255)` | sim | |
| `RangeMin` | `range_min` | `numeric(18,4)` | não | Limite inferior inclusivo |
| `RangeMax` | `range_max` | `numeric(18,4)` | **sim** | `*decimal`; nil = faixa aberta. Constraint `check_ranges`: `range_max IS NULL OR range_max > range_min` |
| `AliqPercent` | `aliq_percent` | `numeric(18,4)` | não | Alíquota (%) |
| `DeductionVal` | `deduction_val` | `numeric(18,4)` (default 0) | sim | Parcela a deduzir (fórmula `base×alíquota−parcela` — D3) |
| `ValidFrom` | `valid_from` | `date` | não | Início de vigência (inclusivo) |
| `ValidTo` | `valid_to` | `date` | **sim** | `*time.Time`; nil = vigente indefinidamente |

> Índice `idx_tax_rules_validity (tax_definition_id, valid_from, valid_to)` suporta o filtro de vigência. 🟢

### `TaxCalculationLog` ↔ `tax_calculation_logs` 🟡 (evolução futura — A2)

⚠️ **Discrepância confirmada:** a struct mapeia `tax_calculation_log` (singular); o DDL define **`tax_calculation_logs`** (plural).

| Campo Go | Coluna (DDL) | Tipo db | Nulável | Notas |
|----------|--------------|---------|---------|-------|
| `ID` | `id` | `uuid` (default `uuid_generate_v4()`) | não | PK |
| `CalculationDate` | `calculation_date` | `timestamptz` (default now) | sim | |
| `TaxCode` | `tax_code` | `varchar(20)` | não | |
| `InputBaseValue` | `input_base_value` | `numeric(18,4)` | não | |
| `CalculatedAmount` | `calculated_amount` | `numeric(18,4)` | não | |
| `AppliedAliq` | `applied_aliq` | `numeric(18,4)` | sim | |
| *(ausente na struct)* | `request_metadata` | `jsonb` | sim | 🔴 Coluna sem campo Go correspondente — reconciliar na evolução A2 |
| `TraceID` | `trace_id` | `varchar(100)` | sim | Correlação de observabilidade |

> Índice `idx_calc_logs_date` em `calculation_date`. Nenhuma query desta lib lê/escreve esta tabela hoje. 🟢

---

## 2. Configuração (sem struct dedicada) ↔ `tax_configs` 🟢

Lida pelo `repository` como `map[config_key]config_value`.

| Coluna (DDL) | Tipo db | Nulável | Notas |
|--------------|---------|---------|-------|
| `id` | `serial4` (PK) | não | |
| `tax_code` | `varchar(20)` | não | UNIQUE composto: `(tax_code, config_key, valid_from)` |
| `config_key` | `varchar(50)` | não | Catálogo confirmado (D2): `pension_percentage`, `dependents_qty`, `pgbl_contribution`, `education_expenses`, `health_expenses` 🟢 |
| `config_value` | `numeric(18,4)` | não | |
| `description` | `text` | sim | |
| `valid_from` | `date` | não | |
| `valid_to` | `date` | sim | nil = vigente |

---

## 3. DTOs de transporte (contrato `json:` — wire com serviços consumidores)

### `TaxRequest` 🟢 — payload IRPF

| Campo Go | JSON | Tipo | Obrigatório |
|----------|------|------|-------------|
| `Type` | `calculation_type` | string (`monthly`/`annual`) | sim |
| `RefDate` | `reference_date` | RFC3339 | sim |
| `GrossIncome` | `gross_income` | decimal (string) | sim |
| `PensionDeductionAmount` | `pension_deduction_amount` | decimal | não |
| `PensionDeductionPercentage` | `pension_deduction_percentage` | decimal | não |
| `DependentsCount` | `dependents_count` | int | não |
| `EducationExpenses` | `education_expenses` | decimal | não |
| `HealthExpenses` | `health_expenses` | decimal | não |
| `PGBLContribution` | `pgbl_contribution` | decimal | não |
| `OtherDeductions` | `other_deductions` | decimal | não |

### `UniversalTaxRequest` 🟢 — payload genérico (formato observado em produção — D2)

| Campo Go | JSON | Tipo | Obrigatório |
|----------|------|------|-------------|
| `TaxCode` | `tax_code` | string (`IRPF`/`INSS`) | sim |
| `CalculationType` | `calculation_type` | string | sim |
| `ReferenceDate` | `reference_date` | RFC3339 | sim |
| `GrossIncome` | `gross_income` | decimal (string) | sim |
| `Inputs` | `inputs` | `[]documentoFiscalRequest` | não |

**`documentoFiscalRequest`** (item de `inputs`): `{ "key": string, "value": decimal, "unit": string }`. Unidades confirmadas: `percentage`, `count`, `amount` (D2).

Exemplo real (D2):
```json
{
  "tax_code": "IRPF",
  "reference_date": "2025-12-31T23:59:59Z",
  "calculation_type": "monthly",
  "gross_income": "36000.00",
  "inputs": [
    { "key": "pension_percentage", "value": "11.0",    "unit": "percentage" },
    { "key": "dependents_qty",     "value": "1",       "unit": "count" },
    { "key": "pgbl_contribution",  "value": "1200.11", "unit": "amount" },
    { "key": "education_expenses", "value": "3340.23", "unit": "amount" },
    { "key": "health_expenses",    "value": "1233.23", "unit": "amount" }
  ]
}
```
> 🟡 Os nomes de chave de `inputs` (`pension_percentage`, `dependents_qty`) diferem dos campos de `TaxRequest` (`pension_deduction_percentage`, `dependents_count`) — a normalização entre os dois formatos ocorre na camada consumidora (A1).

### `TaxResponse` 🟢 — resultado (produzido pela camada consumidora)

| Campo Go | JSON | Tipo | Notas |
|----------|------|------|-------|
| `ReferenceDate` | `reference_date` | string | |
| `GrossIncome` | `gross_income` | decimal | |
| `TotalDeductionAmount` | `total_deduction_amount` | decimal | Soma das deduções |
| `BaseValue` | `base_value` | decimal | Base após deduções |
| `TaxAmount` | `tax_amount` | decimal | Imposto devido |
| `EffectiveRate` | `effective_rate` | decimal | Alíquota efetiva |
| `IsRecommended` | `is_recommended` | bool | **true** no cenário de **menor imposto devido** (D4) |
| `DeductionDetails` | `deduction_details` | `[]DeductionDetail` | `{ type, amount }` |
| `AppliedRule` | `applied_rule_summary` | string | Resumo da regra aplicada |
| `UsedConfigs` | `used_configs` | `map[string]string` | Rastreabilidade das configs usadas |

---

## 4. Relacionamentos (confirmados pelo DDL — A3)

```
tax_definitions (1) ──< (N) tax_rules_history     [FK tax_definition_id → tax_definitions.id]   🟢
tax_definitions.tax_code  ──  tax_configs.tax_code            [associação por código]            🟢
tax_definitions.tax_code  ──  tax_calculation_logs.tax_code   [associação por código, futura A2] 🟡
```
