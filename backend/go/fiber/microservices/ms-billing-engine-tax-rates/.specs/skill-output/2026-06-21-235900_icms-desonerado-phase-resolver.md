# Implementação F-004 e F-005 — ICMS Desonerado + Phase Resolver

**Data:** 2026-06-21  
**Skill:** golang-pro + spec-miner  
**Features:** [FEATURE-2026-06-21](../features/FEATURE-2026-06-21.md) — F-004, F-005

## F-004: ICMS Desonerado — Redução de Base e Limitação de Alíquota

### Arquivos criados/modificados

| Arquivo | Ação | Descrição |
|---------|------|-----------|
| `internal/legacy/icms_desoneracao.go` | Criado | 310 linhas. `ICMSDesoneracao` struct com 2 modos (BR-TAX-CALC-021/022), validação CST (BR-TAX-CONS-013), 13 motivos SEFAZ. |
| `internal/legacy/icms_desoneracao_test.go` | Criado | 11 testes unitários cobrindo ambos os modos, validação CST, modo default, abate de valor, Simples Nacional. |
| `internal/legacy/icms.go` | Modificado | Integração da desoneração em `calcularICMSOperacaoInterna()` — verifica flag `ITEM_DESONERACAO_APLICAR`, valida CST, aplica modo apropriado antes do fallback ICMS próprio. |

### Regras implementadas

- **BR-TAX-CALC-021 (Redução de Base):** `Base_Reduzida = Valor × (1 − PctRedução/100)`, `ICMS = Base_Reduzida × Alíquota/100`
- **BR-TAX-CALC-022 (Limitação de Alíquota):** `Base_Reduzida = Valor × (AliqAlvo/AliqNominal)`, `ICMS = Base_Reduzida × Alíquota/100`
- **BR-TAX-CONS-013 (Validação CST):** CSTs {20, 30, 40, 41, 50, 70, 90} permitem desoneração. CST 00 não permite.
- **BR-TAX-ACT-007 (Abate vICMSDeson):** `vICMSDeson = (Valor × Alíquota) − ICMS` registrado nos detalhes.
- **SOP-017 edge case (Simples Nacional):** Desoneração "clássica" não aplicada em CRT=1.
- **FCP integrado** sobre base reduzida quando `PercentualFCP > 0`.

### Estrutura dos motivos SEFAZ

```go
const (
    MotDesoneracaoTaxi                  = 1  // Taxi
    MotDesoneracaoDeficienteFisico      = 2  // Deficiente fisico
    MotDesoneracaoProdutorAgropecuario  = 3  // Produtor Agropecuario
    MotDesoneracaoFrotistaLocadora      = 4  // Frotista/Locadora
    MotDesoneracaoDiplomaticoConsular   = 5  // Diplomatico/Consular
    MotDesoneracaoUtilitariosAmazonia   = 6  // Utilitarios Amazonia
    MotDesoneracaoSUFRAMA               = 7  // SUFRAMA
    MotDesoneracaoVendaOrgaosPublicos   = 8  // Venda a Orgaos Publicos
    MotDesoneracaoOutros                = 9  // Outros (default)
    MotDesoneracaoCATGuiado             = 10 // CAT Guiado
    MotDesoneracaoCATNaoGuiado          = 11 // CAT Nao Guiado
    MotDesoneracaoOrgaoFomento          = 12 // Orgao de Fomento
    MotDesoneracaoOlimpiadas            = 90 // Olimpiadas (historico)
)
```

## F-005: Phase Resolution System — Fases da Reforma Tributária

### Arquivos criados/modificados

| Arquivo | Ação | Descrição |
|---------|------|-----------|
| `internal/phase/phase.go` | Criado | 198 linhas. `Phase` enum (4 fases), `PhaseResolver`, `PhaseInfo`, `GetReductionFactor()`. |
| `internal/phase/tax_selector.go` | Criado | 160 linhas. `TaxSelector` com `Filter()`, `CalculatorFilter`, `ShouldIncludeInTotal()`. |
| `internal/phase/phase_test.go` | Criado | 14 testes: 4 fases resolvidas, flags shadow/IVADual, fator redução, inclusão total. |
| `internal/calculator/engine.go` | Modificado | `ProcessWithPhase()` com shadow tax exclusion, subnational reduction, IVA Dual extinction. Process() delega para ProcessWithPhase com default filter. |
| `cmd/api/main.go` | Modificado | Wiring de `PhaseResolver` + `TaxSelector`. Endpoint `/calculate` agora chama `ProcessWithPhase()` com filter baseado em `DataOperacao`. |

### Pipeline Phase-Aware

```
POST /calculate
  ↓
TaxSelector.Filter(input.DataOperacao)
  ↓
engine.ProcessWithPhase(ctx, input, filter)
  ├── Fase 1 (Sequencial): IPI (sempre ativo)
  ├── Fase 2 (Paralela): ICMS + PIS/COFINS + Reforma + ISS
  ├── Pós-processamento:
  │   ├── IVA_DUAL? → zera PIS/COFINS/ICMS/ISS (BR-TAX-ACT-006)
  │   └── Transicao? → aplica fator redução ICMS/ISS
  ├── Fase 3 (Sequencial): FUST → FUNTTEL
  └── Consolidação:
      ├── ShadowCBS/ShadowIBS → excluídos de total_impostos
      └── total_impostos ← apenas tributos ativos na fase
```

### Matriz DT-001 (TaxSelector)

| Fase | CBS | IBS | PIS/COFINS | ICMS | ISS | IPI | IS |
|------|-----|-----|-----------|------|-----|-----|----|
| SHADOW_RUN | Shadow | Shadow | Ativo | Ativo | Ativo | Ativo | Ativo |
| CBS_PLENA | Ativo | Shadow | Extinto | Ativo | Ativo | Ativo | Ativo |
| TRANSICAO | Ativo | Ativo | Extinto | Reduzido* | Reduzido* | Ativo | Ativo |
| IVA_DUAL | Ativo | Ativo | Extinto | Extinto | Extinto | Ativo | Ativo |

*Fator redução: 2029=25%, 2030=50%, 2031=75%, 2032=100%

## Resultado dos testes

```
$ go test ./...
ok  ms-billing-engine-tax-rates/internal/calculator  0.005s
ok  ms-billing-engine-tax-rates/internal/legacy       0.009s
ok  ms-billing-engine-tax-rates/internal/middleware    0.010s
ok  ms-billing-engine-tax-rates/internal/phase        0.004s
ok  ms-billing-engine-tax-rates/internal/reforma      0.004s

Total: 136 testes passando (25 novos: 11 desoneração + 14 phase resolver)
```

## Dívidas técnicas remanescentes

- DT-16: Wiring em `main.go` ainda é estático — ideal seria montagem dinâmica via `TaxSelector.Filter()`
- DT-17: Sem testes de integração end-to-end com diferentes datas de operação no endpoint `/calculate`
