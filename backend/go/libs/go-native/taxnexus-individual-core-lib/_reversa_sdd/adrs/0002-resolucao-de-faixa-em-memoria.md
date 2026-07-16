# ADR-0002 — Resolver a faixa aplicável em memória, não no SQL

> ADR **retroativo** inferido do código (sem histórico Git). · Detetive · 2026-06-10
> Confiança: 🟢 CONFIRMADO (decisão e refatoração explícitas no comentário do código)

## Status
Aceito (vigente). O comentário "agora utiliza a lógica de busca em memória após recuperar todas as faixas" (`tax_repository.go:25`) indica que houve uma **versão anterior que filtrava no SQL** e foi refatorada.

## Contexto
Para encontrar a faixa de imposto de uma base de cálculo, havia duas opções: (a) uma query `WHERE range_min <= v AND (range_max IS NULL OR range_max >= v)` por consulta, ou (b) buscar todas as faixas vigentes uma vez e percorrê-las em memória.

## Decisão
`GetApplicableRule` **reutiliza** `GetTaxRulesForPeriod` (que retorna todas as faixas vigentes, ordenadas por `range_min ASC`) e percorre o slice em memória, retornando o **primeiro match**. A seleção de faixa não toca o banco diretamente.
- Evidência: `tax_repository.go:26-41`.

## Consequências
- 🟢 **Cacheabilidade:** uma única chave Redis (`tax_rules_list:<taxCode>:<data>`) serve todas as consultas de faixa daquele imposto/data, independentemente da base de cálculo. Reduz drasticamente o hit no Postgres.
- 🟢 Cálculos progressivos (que precisam de **todas** as faixas, não só uma) reaproveitam a mesma chamada — coerente com o comentário "útil para cálculos progressivos como INSS".
- 🟡 **Pressuposto crítico:** faixas contíguas e não-sobrepostas + ordenação `range_min ASC` ⇒ "primeiro match" é correto. Dados com sobreposição quebram a invariante (ver lacuna D6).
- 🟡 Carrega todas as faixas para a memória; trivial dado o nº pequeno de escalões.
