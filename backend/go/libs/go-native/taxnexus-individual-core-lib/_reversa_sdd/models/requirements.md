# models — Requisitos

> Definições de structs de domínio (entidades persistidas) e DTOs de request/response do simulador de impostos.
> Gerado pelo **Redator** (Reversa) em 2026-06-10 · Atualizado pelo **Revisor** em 2026-06-12
> Fontes: `models/tax_models.go`, `data-dictionary.md`, DDL (resposta A3), respostas D1–D8.
> Confiança: 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

## Visão Geral

O pacote `models` é a **camada de tipos** da biblioteca: define as entidades persistidas (`TaxDefinition`, `TaxRule`, `TaxCalculationLog`) e os DTOs de transporte (`TaxRequest`, `UniversalTaxRequest`, `TaxResponse` e auxiliares). Não contém lógica — apenas structs com tags `db:`/`json:` que estabelecem o **contrato de dados** entre o banco, esta lib e os serviços consumidores.

## Responsabilidades

- Modelar as entidades persistidas do schema `individual_tax_rates`. 🟢
- Definir os DTOs de entrada/saída que os serviços consumidores usam ao integrar com o simulador. 🟢
- Garantir **precisão monetária** usando `shopspring/decimal` (nunca `float`) em todo valor financeiro/percentual. 🟢
- Modelar colunas NULL-áveis via ponteiros (`*decimal.Decimal`, `*time.Time`). 🟢

## Regras de Negócio

- **Precisão fiscal:** todo valor monetário e alíquota usa `decimal.Decimal`; o banco usa `numeric(18,4)` (confirmado no DDL — A3). 🟢
- **Nuláveis por ponteiro:** `range_max` e `valid_to` são `*decimal.Decimal`/`*time.Time` — `nil` significa "sem teto" / "vigente indefinidamente". 🟢
- **`tax_code` suportados:** `IRPF` e `INSS` (D1). 🟢
- **Escopo do `TaxRequest`:** modela o simulador IRPF — tipo `monthly`/`annual`, deduções previdenciárias (valor fixo **ou** percentual), dependentes, educação, saúde, PGBL, outras. 🟢 *(a soma/fórmula é da camada consumidora — A1)*
- **`UniversalTaxRequest`:** payload genérico por `tax_code` com lista de `inputs` chave-valor-unidade; as chaves confirmadas em produção (D2) são `pension_percentage`, `dependents_qty`, `pgbl_contribution`, `education_expenses`, `health_expenses`. 🟢
- **`IsRecommended`:** marca, entre cenários alternativos, o de **menor imposto devido** (D4); a decisão é tomada pela camada consumidora. 🟢

## Requisitos Funcionais

| ID | Requisito | Prioridade | Critério de Aceite |
|----|-----------|-----------|-------------------|
| RF-01 | `TaxDefinition` mapeia `tax_definitions` (id, tax_code, name, sphere, rounding_precision, active) | Must | Scan de uma linha de `tax_definitions` preenche todos os campos sem perda |
| RF-02 | `TaxRule` mapeia `tax_rules_history` com `range_max`/`valid_to` opcionais | Must | Linha com `range_max`/`valid_to` NULL produz ponteiros nil; linha preenchida produz valores |
| RF-03 | `TaxRequest` desserializa o payload IRPF (JSON) preservando precisão decimal | Must | JSON com `gross_income`, deduções e `calculation_type` mapeia para os campos corretos |
| RF-04 | `UniversalTaxRequest` desserializa payload genérico com `inputs[]` (key/value/unit) | Must | Os 5 `inputs` confirmados (D2) são lidos com `value` in `decimal` e `unit` correta |
| RF-05 | `TaxResponse` serializa o resultado do cálculo (base, imposto, alíquota efetiva, recomendação, detalhamento) | Must | Resposta inclui `is_recommended`, `deduction_details`, `applied_rule_summary`, `used_configs` |
| RF-06 | `TaxCalculationLog` modela o log de auditoria (evolução futura — A2) | Must | Struct existe e alinha com a tabela `tax_calculation_logs` (plural) |

## Requisitos Não Funcionais

| Tipo | Requisito inferido | Evidência no código | Confiança |
|------|--------------------|---------------------|-----------|
| Integridade | Precisão decimal exata (sem ponto flutuante) em valores fiscais | `models/tax_models.go` (todos os campos monetários) | 🟢 |
| Compatibilidade | Tags `json:` definem o contrato de wire com serviços consumidores | `models/tax_models.go` | 🟢 |
| Compatibilidade | Tags `db:` alinham as structs ao schema `individual_tax_rates` | `models/tax_models.go` + DDL (A3) | 🟢 |

> Inferido a partir do código e do DDL fornecido (A3).

## Critérios de Aceitação

```gherkin
Dado uma linha de tax_rules_history com range_max e valid_to NULL
Quando a linha é mapeada para TaxRule
Então RangeMax e ValidTo são ponteiros nil (faixa aberta / vigente indefinidamente)

Dado um payload UniversalTaxRequest com inputs [pension_percentage, dependents_qty, pgbl_contribution, education_expenses, health_expenses]
Quando o JSON é desserializado
Então cada input tem Key, Value (decimal) e Unit corretos (percentage / count / amount)

Dado um valor monetário como "36000.00" no JSON
Quando mapeado para um campo decimal.Decimal
Então o valor é preservado com precisão exata, sem erro de ponto flutuante
```

## Prioridade (MoSCoW)

| Requisito | MoSCoW | Justificativa |
|-----------|--------|---------------|
| `TaxRule` / `TaxDefinition` (entidades persistidas) | Must | Consumidas diretamente pelo `repository` |
| `decimal.Decimal` em valores fiscais | Must | Regra de negócio sem alternativa (precisão fiscal) |
| `TaxResponse` / `TaxRequest` / `UniversalTaxRequest` (DTOs) | Must | Contrato de fronteira com serviços consumidores |
| `DeductionDetail` / `documentoFiscalRequest` (auxiliares) | Should | Compõem os DTOs principais |
| `TaxCalculationLog` | Could | Evolução futura (A2); struct definida conforme DDL |

## Rastreabilidade de Código

| Arquivo | Função / Classe | Cobertura |
|---------|-----------------|-----------|
| `models/tax_models.go` | `TaxDefinition` | 🟢 |
| `models/tax_models.go` | `TaxRule` | 🟢 |
| `models/tax_models.go` | `TaxCalculationLog` | 🟢 (struct vs. tabela `tax_calculation_logs` conforme DDL) |
| `models/tax_models.go` | `TaxRequest` | 🟢 |
| `models/tax_models.go` | `UniversalTaxRequest` / `documentoFiscalRequest` | 🟢 |
| `models/tax_models.go` | `TaxResponse` / `DeductionDetail` | 🟢 |
