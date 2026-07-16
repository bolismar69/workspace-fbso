# Feature: Migrar Alíquotas PIS/COFINS para Banco

> **Data:** 2026-06-20
> **Arquivo modificado:** `internal/legacy/pis_cofins.go`

## Resumo

Migração das alíquotas de PIS (1.65%) e COFINS (7.6%) de valores hardcoded para consulta à tabela `federal_tax_rules` via `repository.GetFederalTaxRule()`, com fallback para defaults quando a regra não existir no banco.

## Problema Anterior

```go
// No futuro, estas alíquotas virão do c.repo.GetTaxRate(...)
aliqPis := decimal.NewFromFloat(1.65)     // HARDCODED
aliqCofins := decimal.NewFromFloat(7.6)    // HARDCODED
```

As alíquotas estavam fixas no código, impossibilitando:
- Alteração de alíquotas sem deploy
- Diferentes alíquotas por regime tributário (Lucro Real vs Presumido)
- Diferentes alíquotas por CST
- Configuração de exclusão do ICMS da base por regra

## Solução Implementada

### Integração com `federal_tax_rules`

```go
federalRule, err := c.repo.GetFederalTaxRule(ctx, regime, cstPis, cstCofins)
```

A consulta usa:
- `regime`: CRT do emitente normalizado (LUCRO_REAL, LUCRO_PRESUMIDO, SIMPLES)
- `cstPis`: CST do PIS do item
- `cstCofins`: CST do COFINS do item

### Prioridade de alíquotas

```
Banco (federal_tax_rules) > Default hardcoded (1.65% / 7.6%)
```

1. Se `GetFederalTaxRule()` retornar regra com `AliquotaPIS` não-zero → usa do banco
2. Se `AliquotaPIS` for zero ou regra não encontrada → fallback para 1.65%
3. Mesma lógica para COFINS (fallback 7.6%)

### Exclusão do ICMS da base

Anteriormente, a exclusão do ICMS da base de PIS/COFINS ("Tese do Século") era sempre aplicada quando o `LegacyAdapter` injetava `VALOR_EXCLUSAO_ICMS`.

Agora, a flag `ExcluiICMSBase` da `federal_tax_rules` controla o comportamento:
- `ExcluiICMSBase = true` → ICMS é subtraído da base (comportamento padrão)
- `ExcluiICMSBase = false` → base = valor do item (sem exclusão)

### Observabilidade

- Log `slog.Warn` quando regra federal não encontrada (com regime, CSTs, erro)
- Log `slog.Debug` com alíquotas resolvidas e fonte (banco vs default)
- Campo `fonte_aliquota` nos `MoreTextDetails` do tributo para auditoria
- Campo `metodo` nos `MoreTextDetails` indica se ICMS foi excluído da base

### Constantes de fallback

```go
const (
    defaultAliquotaPIS    = 1.65
    defaultAliquotaCOFINS = 7.6
)
```

Centralizadas como constantes do pacote para facilitar manutenção futura.
