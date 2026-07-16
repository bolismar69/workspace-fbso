# Legacy Mapping — módulo `repository`

> Gerado pelo **Arqueólogo** (Reversa) em 2026-06-10

| Símbolo | Tipo | Arquivo:linha | Notas |
|---------|------|---------------|-------|
| `TaxRepository` | struct | `repository/tax_repository.go:16` | Campos `db *pgxpool.Pool`, `rdb *redis.Client` |
| `NewTaxRepository` | construtor | `repository/tax_repository.go:21` | Injeção de dependência |
| `GetApplicableRule` | método | `repository/tax_repository.go:26` | Resolução de faixa em memória (linha 33–39) |
| `GetConfig` | método | `repository/tax_repository.go:44` | Sem cache; `LIMIT 1` |
| `GetTableConfigs` | método | `repository/tax_repository.go:61` | Cache-aside; chave `tax_configs:...` (linha 62) |
| `GetTaxRulesForPeriod` | método | `repository/tax_repository.go:105` | Cache-aside; JOIN history+definitions (linha 115–125) |

**SQL embutido:**
- `tax_configs` (single) — linhas 46–51
- `tax_configs` (todas) — linhas 73–77
- `tax_rules_history` JOIN `tax_definitions` — linhas 115–125

**Chaves de cache Redis:** `tax_configs:<taxCode>:<date>` (l.62), `tax_rules_list:<taxCode>:<date>` (l.106) — TTL `12*time.Hour` (l.97, l.149).
