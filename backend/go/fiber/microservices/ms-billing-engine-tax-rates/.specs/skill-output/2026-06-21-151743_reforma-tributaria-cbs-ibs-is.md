# Implementacao Reforma Tributaria: CBS/IBS/IS

**Data:** 2026-06-21 15:17:43
**Feature:** CBS/IBS/IS baseado no schema `iva_dual_rules`
**Prioridade:** Media
**Status:** Concluido

## Resumo

Implementacao do modulo de calculo da Reforma Tributaria (CBS, IBS, IS) utilizando a tabela `billing_tax_rates.iva_dual_rules` como fonte de aliquotas. O calculo segue o modelo de IVA Dual onde CBS (federal) e IBS (estadual + municipal) compartilham a mesma base de calculo e regras, enquanto o IS (Imposto Seletivo) e calculado separadamente apenas para produtos marcados como `is_imposto_seletivo = true`.

## Arquivos Modificados

### Core Library (`taxnexus-billing-core-lib`)

| Arquivo | Mudanca |
|---------|---------|
| `repository/entities.go:134-147` | Adicionada entidade `IvaDualRule` com campos: NCM, UFDestino, MunicipioDestinoIBGE, AliquotaCBS, AliquotaIBSEstadual, AliquotaIBSMunicipal, PercentualReducao, IsImpostoSeletivo, AliquotaIS |
| `repository/contracts.go:19` | Adicionado `GetIvaDualRule(ctx, ncm, ufDestino, municipioIBGE) (*IvaDualRule, error)` a interface `TaxRepository` |
| `repository/postgres_repository.go:359-401` | Implementacao PostgreSQL com query por NCM + UF destino + municipio IBGE (ORDER BY municipio DESC para priorizar regra especifica sobre padrao estadual) |
| `repository/cached_tax_repository.go:136-154` | Implementacao com cache Redis (TTL 24h, chave `tax:iva:<ncm>:<uf>:<municipio>`) |

### Microservice (`ms-billing-engine-tax-rates`)

| Arquivo | Mudanca |
|---------|---------|
| `internal/reforma/reforma.go` | **NOVO** - `ReformaCalculator` implementando `TaxCalculator`. Calcula CBS, IBS e IS por item baseado em `iva_dual_rules` |
| `internal/reforma/reforma_test.go` | **NOVO** - 7 testes cobrindo aliquotas normais, reducao 60%, isencao 100%, imposto seletivo, regra nao encontrada, multiplos itens, municipio especifico |
| `internal/legacy/mock_repository_test.go:63-65` | Adicionado stub `GetIvaDualRule` ao mock repository |
| `cmd/api/main.go:20,53,63` | Descomentada importacao `internal/reforma`, instanciado `reformaCalc`, adicionado ao engine como Fase 2 (paralela) |

## Design da Solucao

### Fluxo de Calculo

```
Para cada item do documento:
  1. Extrai NCM, UF destino, municipio IBGE
  2. Consulta GetIvaDualRule(ncm, ufDestino, municipioIBGE)
  3. Se regra nao encontrada -> ignora item (log WARN)
  4. Se regra encontrada:
     a. Calcula fatorReducao = 1 - (percentualReducao / 100)
     b. Se reducao < 100% (nao isento):
        - CBS = base * (aliquotaCBS * fatorReducao / 100)
        - IBS = base * ((aliquotaIBSEstadual + aliquotaIBSMunicipal) * fatorReducao / 100)
     c. Se isImpostoSeletivo && aliquotaIS > 0:
        - IS = base * (aliquotaIS / 100)
```

### Logica de Reducao

- `percentual_reducao = 0` → aliquotas integrais
- `percentual_reducao = 60` → aliquotas reduzidas a 40% do nominal (ex: 8.8% → 3.52%)
- `percentual_reducao >= 100` → isencao total (CBS e IBS nao sao gerados)

### Posicionamento no Engine

O `ReformaCalculator` e executado na **Fase 2** (paralela) do `BillingEngineStruct`, junto com ICMS e PIS/COFINS. Nao depende de valores pre-calculados da Fase 1 (IPI).

## Cobertura de Testes

```go
// 7 cenarios testados:
TestReformaCalculator_CBS_IBS_AliquotasNormais  // Aliquotas padrao: CBS 8.8%, IBS 11.3%
TestReformaCalculator_Reducao60                 // Reducao 60%: CBS 3.52%, IBS 4.52%
TestReformaCalculator_Isento100                 // Isencao total: sem CBS/IBS gerados
TestReformaCalculator_ImpostoSeletivo           // IS = 50% sobre cigarros (NCM 24022000)
TestReformaCalculator_RuleNotFound              // NCM sem regra: zero tributos
TestReformaCalculator_MultiplosItens            // 2 itens com SKUs diferentes
TestReformaCalculator_MunicipioEspecifico       // Aliquota municipal (13%) > padrao estadual
```

**Total de testes no projeto: 71** (64 existentes + 7 novos da Reforma)

## Tributos Gerados

| Tributo | CST | Descricao |
|---------|-----|-----------|
| `CBS` | `01` (normal) / `04` (com reducao) | Contribuicao sobre Bens e Servicos - federal |
| `IBS` | `01` (normal) / `04` (com reducao) | Imposto sobre Bens e Servicos - estadual + municipal |
| `IS` | `01` | Imposto Seletivo - apenas produtos marcados |

Cada tributo inclui `MoreNumericDetails` com aliquotas nominais/efetivas, fator de reducao e `MoreTextDetails` com UF destino, NCM e fonte (`iva_dual_rules`).

## Verificacao

- `go build ./...` → sucesso (microservice + core-lib)
- `go test ./internal/...` → 71/71 PASS
- Mock repository atualizado com `GetIvaDualRule`

## Dividas Tecnicas

- CST da Reforma Tributaria usa valores provisorios (`01` / `04`). A RFB ainda nao publicou a tabela oficial de CST para CBS/IBS.
- Integracao com creditos (cash forward / `permite_credito_amplo`) nao implementada.
