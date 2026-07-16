# Fluxogramas — módulo `repository`

> Gerado pelo **Arqueólogo** (Reversa) em 2026-06-10
> Fonte: `repository/tax_repository.go`

---

## `GetApplicableRule` — resolução de faixa progressiva

```mermaid
flowchart TD
    A[GetApplicableRule ctx, taxCode, baseValue, refDate] --> B[GetTaxRulesForPeriod taxCode, refDate]
    B --> C{erro?}
    C -- sim --> Z[retorna nil, err]
    C -- não --> D[itera rules ordenadas por range_min ASC]
    D --> E{baseValue >= rule.RangeMin?}
    E -- não --> D
    E -- sim --> F{rule.RangeMax == nil<br/>OU baseValue <= rule.RangeMax?}
    F -- não --> D
    F -- sim --> G[retorna &rule, nil]
    D -. fim do loop sem match .-> H[retorna nil,<br/>erro 'no applicable rule found']
```

---

## `GetTaxRulesForPeriod` — cache-aside + query historizada

```mermaid
flowchart TD
    A[GetTaxRulesForPeriod taxCode, refDate] --> K[cacheKey = tax_rules_list:taxCode:YYYY-MM-DD]
    K --> R[Redis GET cacheKey]
    R --> H{hit e val != ''?}
    H -- sim --> U[json.Unmarshal]
    U --> UO{unmarshal OK?}
    UO -- sim --> RET1[retorna rules do cache]
    UO -- não --> Q
    H -- não --> Q[Query Postgres:<br/>JOIN tax_rules_history + tax_definitions<br/>WHERE tax_code e vigência<br/>ORDER BY range_min ASC]
    Q --> QE{erro?}
    QE -- sim --> ERR[retorna nil, erro envelopado %w]
    QE -- não --> SCAN[Scan linha a linha em TaxRule]
    SCAN --> NE{len rules > 0?}
    NE -- sim --> SET[Redis SET cacheKey, TTL 12h]
    SET --> RET2[retorna rules]
    NE -- não --> RET2
```

---

## `GetTableConfigs` — cache-aside de configs

```mermaid
flowchart TD
    A[GetTableConfigs taxCode, refDate] --> K[cacheKey = tax_configs:taxCode:YYYY-MM-DD]
    K --> R[Redis GET cacheKey]
    R --> H{hit e unmarshal OK?}
    H -- sim --> RET1[retorna map do cache]
    H -- não --> Q[Query Postgres tax_configs<br/>WHERE tax_code e vigência]
    Q --> QE{erro?}
    QE -- sim --> ERR[retorna nil, err]
    QE -- não --> SCAN[monta map config_key -> config_value]
    SCAN --> NE{len configs > 0?}
    NE -- sim --> SET[Redis SET cacheKey, TTL 12h]
    SET --> RET2[retorna map]
    NE -- não --> RET2
```

---

## `GetConfig` — leitura de valor único (sem cache)

```mermaid
flowchart TD
    A[GetConfig taxCode, key, refDate] --> Q[Query Postgres tax_configs<br/>WHERE tax_code, config_key e vigência<br/>LIMIT 1]
    Q --> S[QueryRow.Scan val]
    S --> E{erro?}
    E -- sim --> ERR[retorna decimal.Zero, err]
    E -- não --> RET[retorna val, nil]
```
