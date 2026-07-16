# ADR-0004 — Versionamento de regras por vigência temporal (temporal validity)

> ADR **retroativo** inferido do código (sem histórico Git). · Detetive · 2026-06-10
> Confiança: 🟢 CONFIRMADO (padrão explícito em todas as queries)

## Status
Aceito (vigente).

## Contexto
A legislação tributária muda ao longo do tempo, mas o sistema precisa calcular impostos para **datas passadas** (declarações retroativas, correções) usando a regra vigente **na época**. Apagar/sobrescrever regras antigas inviabilizaria isso.

## Decisão
Toda regra (`tax_rules_history`) e config (`tax_configs`) carrega `valid_from` e `valid_to` (NULL-ável). As consultas recebem uma `refDate` e filtram:
```sql
valid_from <= $refDate AND (valid_to IS NULL OR valid_to >= $refDate)
```
`valid_to IS NULL` significa vigente indefinidamente. A tabela de faixas chama-se `tax_rules_history`, sinalizando modelo **append-only/historizado**.
- Evidência: `GetConfig:50`, `GetTableConfigs:77`, `GetTaxRulesForPeriod:123-124`.

## Consequências
- 🟢 Cálculo retroativo correto: a `refDate` seleciona a versão da regra/config da época.
- 🟢 Auditabilidade: versões antigas preservadas como histórico.
- 🟡 Exige disciplina de manutenção: ao mudar a lei, encerra-se a versão antiga (`valid_to`) e insere-se a nova (`valid_from`). Esse processo de escrita **não está neste repositório** (lacuna — ver SM-01 e D6).
- 🟡 Risco de "buracos" ou sobreposição de vigência se a manutenção falhar — não há constraint visível que o impeça neste recorte. Validar com o Data Master.
