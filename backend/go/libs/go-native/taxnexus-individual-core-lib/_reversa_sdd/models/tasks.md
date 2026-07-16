# models — Tarefas de Implementação

> Sequência executável para reimplementar os tipos de domínio e DTOs a partir do legado.
> Gerado pelo **Redator** (Reversa) em 2026-06-10 · `doc_level = completo`
> Fonte primária: `models/tax_models.go`. Ver `requirements.md`, `design.md` e `contracts.md` desta unit.
> 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

## Pré-requisitos

- [ ] Dependências disponíveis: `shopspring/decimal`, `google/uuid`, pacote `time`
- [ ] DDL do schema `individual_tax_rates` como referência de tipos (ver `contracts.md` / resposta A3)

## Tarefas

- [ ] **T-01** — Definir `TaxDefinition` com tags `db:` para `tax_definitions` (id, tax_code, name, sphere, rounding_precision, active).
  - Origem no legado: `models/tax_models.go`
  - Critério de pronto: `Scan` de uma linha de `tax_definitions` preenche todos os campos; `tax_code` aceita `IRPF`/`INSS`
  - Confiança: 🟢

- [ ] **T-02** — Definir `TaxRule` com `RangeMax *decimal.Decimal` e `ValidTo *time.Time` (nuláveis por ponteiro) e tags `db:` para `tax_rules_history`.
  - Origem no legado: `models/tax_models.go:25`
  - Critério de pronto: linha com `range_max`/`valid_to` NULL produz nil; demais campos em `decimal.Decimal`
  - Confiança: 🟢

- [ ] **T-03** — Definir os DTOs `TaxRequest` (IRPF) e `UniversalTaxRequest` + `documentoFiscalRequest` (genérico) com tags `json:` conforme `contracts.md`.
  - Origem no legado: `models/tax_models.go:43,56,67`
  - Critério de pronto: desserializa os exemplos reais de D2 (inputs `pension_percentage`, `dependents_qty`, `pgbl_contribution`, `education_expenses`, `health_expenses`) sem perda de precisão
  - Confiança: 🟢

- [ ] **T-04** — Definir `TaxResponse` e `DeductionDetail` com tags `json:` (incluindo `is_recommended`, `deduction_details`, `applied_rule_summary`, `used_configs`).
  - Origem no legado: `models/tax_models.go`
  - Critério de pronto: serializa um resultado completo de cálculo; `is_recommended` reflete o cenário de menor imposto (D4)
  - Confiança: 🟢

- [ ] **T-05** — Definir `TaxCalculationLog` (evolução futura — A2). Ao reconciliar com o DDL, usar tabela **`tax_calculation_logs`** (plural) e prever a coluna `request_metadata jsonb`.
  - Origem no legado: `models/tax_models.go:32` + DDL (A3)
  - Critério de pronto: struct alinhada ao DDL; documentada como não-utilizada nesta versão
  - Confiança: 🟡

## Tarefas de Teste

- [ ] **TT-01** — Round-trip de `TaxRule`: linha NULL → nil → serialização preserva ausência de teto/vigência
- [ ] **TT-02** — Desserialização de `UniversalTaxRequest` a partir dos payloads reais de D2 (2 exemplos: gross 36000 e 360000)
- [ ] **TT-03** — Precisão: valor `"360000.00"` mapeado para `decimal.Decimal` sem erro de float
- [ ] **TT-04** — Serialização de `TaxResponse` com `deduction_details` e `used_configs` preenchidos

## Tarefas de Migração de Dados (se aplicável)

- [ ] **TM-01** — (Evolução A2) Alinhar struct `TaxCalculationLog` à tabela `tax_calculation_logs`, adicionando o campo para `request_metadata jsonb` e corrigindo o nome plural da tabela

## Ordem Sugerida

1. **T-01** e **T-02** (entidades persistidas) primeiro — consumidas pelo `repository`.
2. **T-03** e **T-04** (DTOs) — independentes, podem ser paralelas.
3. **T-05** por último (evolução futura, não bloqueante).

## Lacunas Pendentes (🔴)

- Coluna `request_metadata jsonb` sem campo Go correspondente — decidir o tipo Go na implementação da evolução A2.
- Semântica exata de `sphere` (`bpchar(1)`) e `rounding_precision` — não aplicadas por esta lib; documentadas como inferidas.
