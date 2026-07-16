# ADR-0003 — Cache-aside com Redis e TTL de 12 horas

> ADR **retroativo** inferido do código (sem histórico Git). · Detetive · 2026-06-10
> Confiança: 🟢 CONFIRMADO (padrão e TTL explícitos no código; valor exato do TTL 🟡 quanto à motivação)

## Status
Aceito (vigente).

## Contexto
Parâmetros fiscais (faixas e configs) mudam raramente — tipicamente quando a legislação muda, com vigência granular por **dia** (`refDate`). Consultá-los do Postgres a cada cálculo é desperdício; eles são praticamente imutáveis dentro de uma data de referência.

## Decisão
`GetTableConfigs` e `GetTaxRulesForPeriod` implementam **cache-aside** sobre Redis:
1. `GET` na chave (`tax_configs:<taxCode>:<data>` / `tax_rules_list:<taxCode>:<data>`).
2. Hit + desserialização OK ⇒ retorna.
3. Miss ⇒ consulta Postgres.
4. Resultado não-vazio ⇒ `SET` com TTL `12 * time.Hour`.

Serialização via `encoding/json`. A chave inclui a `refDate` (YYYY-MM-DD), então cada data tem seu próprio cache.
- Evidência: `tax_repository.go:61-101` e `103-153`.

## Consequências
- 🟢 Reduz carga no Postgres para parâmetros estáveis.
- 🟢 Chave por data evita servir parâmetros de vigência errada.
- 🟡 **TTL de 12h hardcoded**, não configurável por ambiente (lacuna L4). Como a chave é por dia e os dados raramente mudam intra-dia, o impacto é baixo, mas uma mudança de lei publicada às 9h só se reflete totalmente após a expiração/rollover do dia.
- 🟢 **Degradação graciosa** (ver ADR-0006): falha de Redis não derruba o caminho feliz.
- ⚠️ **Inconsistência deliberada/acidental:** `GetConfig` (valor único) **não** usa cache, enquanto `GetTableConfigs` usa (lacuna L5 — a confirmar com o time).
