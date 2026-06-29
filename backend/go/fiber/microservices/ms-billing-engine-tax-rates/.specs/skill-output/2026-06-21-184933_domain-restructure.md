# Reorganização da Estrutura por Domínio — 2026-06-21 18:49

## Contexto

O microserviço `ms-billing-engine-tax-rates` estava com a interface `TaxCalculator` definida no pacote `internal/calculator`, misturando contratos de domínio com lógica de engine. A separação em camadas seguindo **Domain-Driven Design** melhora a testabilidade e reduz o acoplamento.

## Mudanças Realizadas

### 1. Criação do pacote `internal/domain/`

**Arquivo novo:** `internal/domain/domain.go`

A interface `TaxCalculator` foi extraída do pacote `calculator` para o novo pacote `domain`, que representa a camada mais interna da arquitetura — sem dependências de outros pacotes internos, apenas da lib compartilhada `taxnexus-billing-core-lib`.

```go
package domain

type TaxCalculator interface {
    Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error)
}
```

### 2. Remoção de `internal/calculator/calculator.go`

O arquivo original que continha apenas a interface `TaxCalculator` foi removido, já que a interface agora reside em `domain/`.

### 3. Atualização de referências em `internal/calculator/`

| Arquivo | Alteração |
|---------|-----------|
| `engine.go` | Adicionado `import "ms-billing-engine-tax-rates/internal/domain"`; `TaxCalculator` → `domain.TaxCalculator` (5 ocorrências) |
| `legacy_adapter.go` | Adicionado `import "ms-billing-engine-tax-rates/internal/domain"`; `TaxCalculator` → `domain.TaxCalculator` (1 ocorrência) |
| `engine_test.go` | Adicionado `import "ms-billing-engine-tax-rates/internal/domain"`; `[]TaxCalculator{` → `[]domain.TaxCalculator{` (2 ocorrências) |

### 4. Atualização de `cmd/`

| Arquivo | Alteração |
|---------|-----------|
| `cmd/api/main.go` | Adicionado `import "...internal/domain"`; `calculator.TaxCalculator` → `domain.TaxCalculator` |
| `cmd/test_engine/main.go` | Adicionado `import "...internal/domain"`; `calculator.TaxCalculator` → `domain.TaxCalculator` |

### 5. Pacotes NÃO modificados

- **`internal/legacy/`** — Não referenciavam `calculator.TaxCalculator` diretamente. O `IPICalculator` implementa a interface via structural typing.
- **`internal/middleware/`** — Sem dependências da interface de cálculo.
- **`internal/reforma/`** — `ReformaCalculator` implementa `TaxCalculator` via structural typing (método `Calculate` com assinatura compatível), sem importar o pacote.

## Estrutura Final

```
internal/
├── domain/            # Interface central TaxCalculator (camada mais interna)
│   └── domain.go
├── calculator/        # BillingEngine + LegacyAdapter
│   ├── engine.go
│   ├── engine_test.go
│   └── legacy_adapter.go
├── legacy/            # Calculadoras legadas (ICMS, IPI, PIS/COFINS)
│   ├── icms.go, icms_calculate_test.go
│   ├── ipi.go, ipi_calculate_test.go
│   ├── pis_cofins.go, pis_cofins_calculate_test.go
│   ├── pis_strategies.go, pis_strategies_test.go
│   ├── cofins_strategies.go, cofins_strategies_test.go
│   └── mock_repository_test.go
├── middleware/        # HTTP middleware (auth, metrics, requestid)
│   ├── auth.go, auth_test.go
│   ├── metrics.go
│   └── requestid.go, requestid_test.go
└── reforma/           # Reforma Tributária (CBS/IBS/IS)
    ├── reforma.go
    └── reforma_test.go
```

## Verificação

- `go build ./...` — ✅ compilação limpa
- `go vet ./...` — ✅ sem warnings
- `go test ./internal/...` — ✅ todos os 5 pacotes passam
  - `calculator`: OK
  - `legacy`: OK
  - `middleware`: OK
  - `reforma`: OK

## Princípio de Design Aplicado

**Dependency Inversion**: A interface `TaxCalculator` está na camada mais interna (`domain/`). As implementações (`legacy/`, `reforma/`) e o orquestrador (`calculator/`) dependem da abstração, não de implementações concretas.

**Structural Typing (Go)**: `IPICalculator` e `ReformaCalculator` satisfazem `domain.TaxCalculator` sem precisar declarar `implements` ou importar o pacote `domain` — zero acoplamento desnecessário.
