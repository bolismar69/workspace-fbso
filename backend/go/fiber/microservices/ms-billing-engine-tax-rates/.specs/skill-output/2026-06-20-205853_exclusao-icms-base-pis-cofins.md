# Implementação — Testes de Exclusão ICMS da Base PIS/COFINS

**Data/Hora:** 2026-06-20T20:58:53-03:00

**Feature:** Exclusão ICMS da base PIS/COFINS — validação com testes + cobertura completa de CSTs

**Origem:** `.specs/governance/confidence-report.md` (ações para   100%, itens 4-5)

---

## Escopo

1. **Validar `ExcluiICMSBase`** via testes de integração no `PISCofinsCalculator.Calculate()`
2. **Cobertura CST** — confirmada 100% (CSTs 01-06, 49, 50-99, 99)
3. **Suite de testes consolidada** — 39 testes no total

---

## Arquivos Criados

### 1. Mock Repository `internal/legacy/mock_repository_test.go`

Padrão: mock manual (sem testify/mockery) implementando `repository.TaxRepository` com 8 métodos.

```go
type mockTaxRepository struct {
    federalRule    *repository.FederalTaxRule
    federalRuleErr error
    // ... demais campos para outras interfaces
}
```

Helpers:
- `newMockWithFederalRule(rule)` — injeta regra federal
- `newMockWithFederalRuleError(err)` — simula erro de banco
- `defaultFederalRule(pis, cofins, excluiICMS)` — factory de `FederalTaxRule`

### 2. Testes de Integração `internal/legacy/pis_cofins_calculate_test.go`

**13 cenários** validando o fluxo completo `PISCofinsCalculator.Calculate()`:

| # | Teste | Cenário |
|---|-------|---------|
| 1 | `ExcluiICMSBase_True` | ICMS excluído da base → PIS/COFINS sobre (valor - ICMS) |
| 2 | `ExcluiICMSBase_False` | ICMS NÃO excluído → base = valor_item |
| 3 | `FallbackDefaults` | Repo retorna nil → usa defaults (1.65% / 7.6%) |
| 4 | `FallbackOnError` | Repo retorna erro → fallback resiliente |
| 5 | `AliquotasCustomizadas` | Alíquotas do banco (3.0% / 9.25%) sobrescrevem defaults |
| 6 | `CST_03_PorUnidade` | CST 03 → CalcTax por quantidade (PIS e COFINS) |
| 7 | `CST_ZeroValue` | CSTs 04, 05, 06, 49, 50, 70, 98, 99 → valor zero |
| 8 | `MultiplosItens` | 2 itens com CSTs diferentes (01 e 04) |
| 9 | `ICMSZeroNaBase` | Base negativa → strategy clampa para zero |
| 10 | `SemCST_NaoCalcula` | Item sem CST PIS → não gera tributo PIS |
| 11 | `AliquotaZeroNoBancoUsaDefault` | Alíquota = zero no banco → não sobrescreve default |
| 12 | `PISOnly` | Item com apenas CST PIS (sem COFINS) |
| 13 | `ExcluiICMSBaseZeroICMS` | ExcluiICMSBase=true mas ICMS=0 → base mantida |

### Helpers de teste

```go
makeDocumentoFiscalEntrada(itens ...) → DocumentoFiscalEntrada
makeItem(cstPis, cstCofins, quantidade, valorUnitario, valorIcmsExcluir) → ItemDocumentoFiscalEntrada
assertDetail(t, details, key, want) → verifica detalhes de saída
```

---

## Resultados da Validação

### Flag `ExcluiICMSBase`

| Flag | Base PIS/COFINS | Valor PIS (1.65%) | Valor COFINS (7.6%) |
|------|-----------------|--------------------|---------------------|
| `true` | `valor_item - ICMS` | `(valor - ICMS) * 1.65%` | `(valor - ICMS) * 7.6%` |
| `false` | `valor_item` | `valor * 1.65%` | `valor * 7.6%` |

### Fonte da alíquota

| Cenário | `fonte_aliquota` | `metodo` |
|---------|-------------------|----------|
| Regra no banco | `banco_federal_tax_rules` | conforme `ExcluiICMSBase` |
| Fallback (nil/erro) | `default_hardcoded` | `base_com_exclusao_icms` |

### Fallback a zero

- Alíquota zero no banco (`IsZero()`) NÃO sobrescreve o default
- Base negativa após exclusão ICMS é clampada a zero pela strategy
- Erro de conexão com banco não interrompe o cálculo

---

## Suite Consolidada

| Arquivo | Testes | Cobertura |
|---------|--------|-----------|
| `pis_strategies_test.go` | 13 | Cálculo por CST, valor zero, roteamento |
| `cofins_strategies_test.go` | 13 | Cálculo por CST, valor zero, roteamento |
| `pis_cofins_calculate_test.go` | 13 | ExcluiICMSBase, fallback, múltiplos itens, edge cases |
| **Total** | **39** | Estratégias + Calculadora |

---

## Verificação

```
go test ./... -count=1    # OK — 39 testes passando (0.007s)
go build ./...            # OK — compilação limpa
go vet ./...              # OK — sem warnings
```

---

## Rastreabilidade

| Lacuna Anterior | Status | Evidência |
|-----------------|--------|-----------|
| Exclusão ICMS base PIS/COFINS (70%) |   100% validado | `pis_cofins_calculate_test.go` — 13 cenários |
| CSTs 04-06, 49, 50-99 (50%) |   100% coberto | Testes por CST + integração com calculadora |
| Testes automatizados (0%) |   40% | 4 arquivos de teste, 39 cenários |
