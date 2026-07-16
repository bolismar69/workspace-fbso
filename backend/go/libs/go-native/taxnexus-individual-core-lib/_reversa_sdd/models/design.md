# models — Design Técnico

> Estrutura dos tipos de domínio e DTOs, com base em `models/tax_models.go` e no DDL real (A3).
> Gerado pelo **Redator** (Reversa) em 2026-06-10 · Atualizado pelo **Revisor** em 2026-06-12
> 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

## Interface

Pacote sem comportamento — expõe apenas tipos. Os contratos de campo detalhados (Go ↔ db ↔ JSON) estão em `contracts.md`.

| Símbolo | Natureza | Tabela / Uso | Observação |
|---------|----------|--------------|------------|
| `TaxDefinition` | entidade | `individual_tax_rates.tax_definitions` | Definição de um imposto (`tax_code` único) 🟢 |
| `TaxRule` | entidade | `individual_tax_rates.tax_rules_history` | Faixa/escalão historizado por vigência 🟢 |
| `TaxCalculationLog` | entidade | `individual_tax_rates.tax_calculation_logs` | Log de auditoria — **evolução futura** (A2) 🟢 |
| `TaxRequest` | DTO | payload IRPF | Deduções específicas do IRPF 🟢 |
| `UniversalTaxRequest` | DTO | payload genérico | `tax_code` + `inputs[]` chave-valor-unidade 🟢 |
| `documentoFiscalRequest` | DTO (não exportado) | item de `Inputs` | `key`, `value`, `unit` 🟢 |
| `TaxResponse` | DTO | resultado do cálculo | Produzido pela camada consumidora 🟢 |
| `DeductionDetail` | DTO | item de `DeductionDetails` | `type`, `amount` 🟢 |

## Fluxo Principal

O pacote não tem fluxo de execução. O ciclo de vida típico dos tipos:

1. `repository` faz `Scan` de linhas de `tax_definitions`/`tax_rules_history` para `TaxDefinition`/`TaxRule` (tags `db:`). `repository/tax_repository.go` 🟢
2. O serviço consumidor desserializa o request HTTP (`TaxRequest` ou `UniversalTaxRequest`) a partir do JSON (tags `json:`). 🟢 *(consumidor externo — A1)*
3. O serviço consumidor calcula e serializa `TaxResponse` de volta ao cliente. 🟢 *(consumidor externo — A1)*

## Fluxos Alternativos

- **Coluna NULL:** `range_max`/`valid_to` ausentes → ponteiro nil; o `repository` interpreta nil como "sem teto" / "vigente". 🟢
- **Input genérico:** `UniversalTaxRequest.Inputs` carrega chaves arbitrárias com `unit` (`percentage`/`count`/`amount`), permitindo que o mesmo DTO sirva IRPF e INSS. 🟢

## Dependências

- **`shopspring/decimal`** — tipo de todos os valores monetários e percentuais. 🟢
- **`google/uuid`** — `TaxCalculationLog.ID` (UUID). 🟢
- **`time`** — campos de vigência e data de referência. 🟢

## Decisões de Design Identificadas

| Decisão | Evidência no código | Confiança |
|---------|---------------------|-----------|
| `decimal.Decimal` para valores fiscais (ADR-0001) | `models/tax_models.go` | 🟢 |
| Ponteiros para colunas NULL-áveis | `models/tax_models.go:25` | 🟢 |
| `documentoFiscalRequest` não exportada usada em campo exportado (`Inputs`) | `models/tax_models.go:56,67` | 🟢 |
| `UniversalTaxRequest` como generalização de `TaxRequest` | `models/tax_models.go` | 🟢 |

## Estado Interno

Nenhum. São tipos de dados imutáveis após construção; sem cache, sem singletons.

## Observabilidade

Nenhuma emitida pelo pacote. `TaxResponse.UsedConfigs` e `AppliedRule` carregam **rastreabilidade do cálculo** (quais configs/regra foram aplicadas), consumida pela camada de cálculo. `TaxCalculationLog.TraceID` destina-se a correlação de auditoria na evolução futura (A2). 🟢

## Riscos e Lacunas

- 🟢 **Resolvido:** A discrepância struct ↔ tabela em `TaxCalculationLog` foi mapeada via DDL. A tabela real é `tax_calculation_logs` (plural) e possui `request_metadata jsonb`. A implementação da evolução futura deve reconciliar a struct com essas colunas.
- 🟢 **Confirmado:** A semântica de `sphere` (federal/estadual/municipal) e `rounding_precision` é puramente informativa para esta lib, sendo aplicada pelo serviço de cálculo.
- 🟡 **`documentoFiscalRequest` não exportada:** consumidores externos não conseguem instanciar `Inputs` diretamente fora deste pacote.
