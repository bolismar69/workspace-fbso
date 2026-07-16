# Implementação — Completar Estratégias PIS/COFINS

**Data/Hora:** 2026-06-20T00:00:00-03:00

**Feature:** Completar estratégias PIS/COFINS — CSTs 04, 05, 06, 49, 50-99

**Prioridade:** Média

**Origem:** `.specs/product/feature-roadmap.md` (linha 32) e `.specs/product/requirements.md` (RF-07)

---

## Arquivos Modificados

### 1. Core-Lib: Constantes CST `models/constants.go`

**Arquivo:** `backend/go/libs/go-native/taxnexus-billing-core-lib/models/constants.go:74-79`

Adicionados os novos valores ao tipo `CSTPISCOFINS`:

```go
CSTPISCOFINS04 CSTPISCOFINS = "04"  // Monofásico
CSTPISCOFINS05 CSTPISCOFINS = "05"  // Substituição Tributária
CSTPISCOFINS06 CSTPISCOFINS = "06"  // Alíquota Zero
CSTPISCOFINS49 CSTPISCOFINS = "49"  // Outras Operações de Saída
```

### 2. Core-Lib: Validação `models/tax_validation.go`

**Arquivo:** `backend/go/libs/go-native/taxnexus-billing-core-lib/models/tax_validation.go:282-287`

Novos CSTs adicionados ao mapa `validCSTPISCOFINS` para permitir validação de input.

### 3. Estratégias PIS `internal/legacy/pis_strategies.go`

**Arquivo:** `backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/legacy/pis_strategies.go`

Novas estratégias adicionadas:

| Strategy | CST | Descrição | Comportamento |
|----------|-----|-----------|---------------|
| `PIS04` | 04 | Monofásico — tributo concentrado no produtor/importador | Retorna `decimal.Zero` |
| `PIS05` | 05 | Substituição Tributária — tributo já recolhido por ST | Retorna `decimal.Zero` |
| `PIS06` | 06 | Alíquota Zero | Retorna `decimal.Zero` |
| `PIS49` | 49 | Outras Operações de Saída | Retorna `decimal.Zero` |
| `PIS50To99` | 50-99 | Operações de crédito, suspensão, outras | Retorna `decimal.Zero` (default) |

**Correção aplicada:** `PIS99` foi corrigido para retornar `decimal.Zero` (antes calculava `base * aliq`), alinhando com a semântica fiscal de "Outras Operações".

### 4. Estratégias COFINS `internal/legacy/cofins_strategies.go`

**Arquivo:** `backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/legacy/cofins_strategies.go`

Novas estratégias adicionadas:

| Strategy | CST | Descrição | Comportamento |
|----------|-----|-----------|---------------|
| `COFINS03` | 03 | Alíquota por Unidade (Quantidade) | `CalcTax(quantidade, aliquota)` |
| `COFINS04` | 04 | Monofásico — tributo concentrado no produtor/importador | Retorna `decimal.Zero` |
| `COFINS05` | 05 | Substituição Tributária — tributo já recolhido por ST | Retorna `decimal.Zero` |
| `COFINS06` | 06 | Alíquota Zero | Retorna `decimal.Zero` |
| `COFINS49` | 49 | Outras Operações de Saída | Retorna `decimal.Zero` |
| `COFINS50To99` | 50-99 | Operações de crédito, suspensão, outras | Retorna `decimal.Zero` (default) |

### 5. Fábricas de Estratégia `internal/legacy/pis_cofins.go`

**Arquivo:** `backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/legacy/pis_cofins.go:162-196`

Ambas as funções `GetPISStrategy` e `GetCOFINSStrategy` foram atualizadas com casos explícitos para cada novo CST. O fallback (`default`) agora retorna `PIS50To99` / `COFINS50To99` respectivamente, cobrindo automaticamente qualquer CST 50-99 não listado explicitamente.

### 6. Testes

**Arquivos criados:**
- `internal/legacy/pis_strategies_test.go` — 26 testes (cálculo, valor zero, roteamento de fábrica)
- `internal/legacy/cofins_strategies_test.go` — 26 testes (cálculo, valor zero, roteamento de fábrica)

Cobertura:
- Testes para cálculo ad valorem (CST 01/02) com e sem exclusão ICMS
- Testes para cálculo por unidade (CST 03)
- Testes de valor zero para CSTs 04, 05, 06, 49, 50-99, 99
- Testes de roteamento da fábrica para todos os CSTs implementados
- Testes de fallback para CSTs desconhecidos/vazios

---

## Regras de Negócio Implementadas

| CST | Fundamentação Fiscal | Efeito no Cálculo |
|-----|---------------------|-------------------|
| **04** | Lei 10.637/2002 Art. 2º — Tributação monofásica concentra PIS/COFINS em uma única etapa (produtor/importador) | Valor zero para etapas subsequentes |
| **05** | IN RFB 1.911/2019 Art. 38 — Substituição tributária: tributo já recolhido na fonte por substituto | Valor zero para o substituído |
| **06** | Lei 10.637/2002 Art. 1º c/c Lei 10.865/2004 — Alíquota zero para produtos específicos | Valor zero |
| **49** | IN RFB 2.121/2022 — Outras operações de saída (não configuram fato gerador próprio) | Valor zero |
| **50-99** | Arts. 31-40 da IN 1.911/2019 — Operações de crédito, suspensão, diferimento, regimes especiais | Valor zero (não geram débito na saída) |

---

## Verificação

```
go build ./...    # OK — compilação limpa
go vet ./...      # OK — sem warnings
go test ./... -v  # OK — 26 testes passando
```

---

## Rastreabilidade

| Requisito | Status |
|-----------|--------|
| RF-07 (PIS/COFINS por CST) |   Completo (100% dos CSTs) |
| Feature Roadmap — "Completar estratégias PIS/COFINS" |   Implementado |
| `validCSTPISCOFINS` validation map |   Atualizado com novos CSTs |
