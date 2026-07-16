# ERD Completo — taxnexus-individual-core-lib

> Gerado pelo **Arquiteto** (Reversa) em 2026-06-10 · Atualizado pelo **Revisor** em 2026-06-12
> Fonte: `data-dictionary.md`, queries em `repository/tax_repository.go` e DDL real (A3).
> Confiança: 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

Schema PostgreSQL: **`individual_tax_rates`**.

```mermaid
erDiagram
    tax_definitions ||--o{ tax_rules_history : "tem faixas (tax_definition_id)"
    tax_definitions ||..o{ tax_configs : "parametriza (por tax_code)"
    tax_definitions ||..o{ tax_calculation_logs : "auditoria (por tax_code)"

    tax_definitions {
        int id PK
        string tax_code "UQ — IRPF, INSS"
        string name
        string sphere "bpchar(1)"
        int rounding_precision
        bool active "padrão true"
    }

    tax_rules_history {
        int id PK
        int tax_definition_id FK
        string description
        decimal range_min "numeric(18,4)"
        decimal range_max "NULL = faixa aberta"
        decimal aliq_percent
        decimal deduction_val "parcela a deduzir"
        date valid_from "inclusivo"
        date valid_to "NULL = vigente"
    }

    tax_configs {
        int id PK
        string tax_code "FK lógico"
        string config_key "UQ (code, key, from)"
        decimal config_value
        text description
        date valid_from
        date valid_to "NULL = vigente"
    }

    tax_calculation_logs {
        uuid id PK "default uuid_generate_v4()"
        timestamp calculation_date "default now()"
        string tax_code
        decimal input_base_value
        decimal calculated_amount
        decimal applied_aliq
        jsonb request_metadata
        string trace_id
    }
```

## Entidades e cardinalidades

| Relacionamento | Cardinalidade | Chave | Confiança |
|----------------|---------------|-------|-----------|
| `tax_definitions` → `tax_rules_history` | 1 : N | `tax_rules_history.tax_definition_id` → `tax_definitions.id` | 🟢 |
| `tax_definitions` → `tax_configs` | 1 : N | associação lógica por `tax_code` | 🟢 |
| `tax_definitions` → `tax_calculation_logs` | 1 : N | associação lógica por `tax_code` | 🟢 |

## Invariantes de dados

| Invariante | Estado | Confiança |
|------------|--------|-----------|
| Faixas (`tax_rules_history`) contíguas e sem sobreposição por imposto/vigência | Garantido (Operacional) | 🟢 D6 |
| Último escalão com `range_max IS NULL` (faixa aberta superior) | Confirmado | 🟢 |
| Versionamento temporal: nunca apaga, encerra via `valid_to` (append-only) | Confirmado | 🟢 |
| `tax_code` único em `tax_definitions` | Confirmado (UQ no DDL) | 🟢 |

## Observações

- **Schema verificado:** O DDL confirmou a estrutura de schemas, tabelas, sequences e índices.
- **Tipagem exata:** Todos os campos monetários/percentuais são `numeric(18,4)`.
- **Auditoria:** A tabela `tax_calculation_logs` (plural) inclui `request_metadata jsonb`, confirmando sua vocação para evolução futura de auditoria detalhada.
