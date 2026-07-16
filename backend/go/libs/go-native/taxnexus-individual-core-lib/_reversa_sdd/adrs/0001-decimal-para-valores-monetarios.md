# ADR-0001 — Usar `decimal.Decimal` para todos os valores monetários e alíquotas

> ADR **retroativo** inferido do código (sem histórico Git). · Detetive · 2026-06-10
> Confiança: 🟢 CONFIRMADO (a decisão é evidente no código; a justificativa é inferida 🟡)

## Status
Aceito (vigente no código).

## Contexto
O sistema lida com tributação de pessoa física — valores monetários, bases de cálculo e alíquotas. Erros de arredondamento de ponto flutuante (`float64`) são inaceitáveis em contexto fiscal/contábil, onde centavos importam e o resultado precisa ser auditável e reprodutível.

## Decisão
Todos os campos monetários e percentuais usam `github.com/shopspring/decimal.Decimal`, nunca `float32`/`float64`. As comparações de faixa usam os métodos `GreaterThanOrEqual`/`LessThanOrEqual` do tipo decimal.
- Evidência: `models/tax_models.go` (todos os campos de valor) e `repository/tax_repository.go:34-35`.

## Consequências
- 🟢 Precisão exata em somas e comparações de faixa.
- 🟡 Custo de desempenho e alocação maior que `float` — irrelevante no volume desta lib.
- 🟢 Serialização JSON do decimal é estável (string), compatível com cache Redis.
- A `RoundingPrecision` de cada imposto (`TaxDefinition`) governa o arredondamento final na camada de cálculo. 🟡
