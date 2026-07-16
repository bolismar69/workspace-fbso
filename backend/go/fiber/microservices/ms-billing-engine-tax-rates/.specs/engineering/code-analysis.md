# Análise Técnica Consolidada — ms-billing-engine-tax-rates

Gerado pelo agente **Spec Miner** em 2026-06-20. Atualizado em 2026-06-30 (PR #6 merge — Fases 0-1-2: Admin Fiscal, Créditos, TaxToken, Simulação, Fornecedores).

## Módulos e Componentes

### 1. Entry Points (`cmd/`)

#### `cmd/api/main.go` — Servidor HTTP

**Propósito:** Ponto de entrada do servidor HTTP da API de cálculo.

**Fluxo de inicialização:**
1. Logger JSON (`slog`) configurado com nível Debug
2. Conexão PostgreSQL (`db.ConnectPostgres`) — Fatal se falhar
3. Conexão Redis (`cache.ConnectRedis`) — Fatal se falhar
4. Repository com camada de cache (`CachedTaxRepository` envolvendo `PostgresTaxRepository`)
5. Instanciação das calculadoras: `IPICalculator`, `ICMSCalculator`, `PISCofinsCalculator`
6. Adaptação via `LegacyAdapter` para interface unificada `TaxCalculator`
7. Montagem do motor multi-fase: `BillingEnginePhased` com 7 fases SOP-013
8. Handlers e middlewares Fiber

**Fonte:** `cmd/api/main.go:1-70`

#### `cmd/test_engine/main.go` — CLI Test Runner

**Propósito:** Test harness manual para execução do motor de cálculo via linha de comando, output JSON.

**Fonte:** `cmd/test_engine/main.go:1-60`

### 2. Motor de Cálculo (`internal/calculator/`)

#### `engine.go` — Orquestrador Multi-Fase (SOP-013 / C-001)

**Estruturas chave:**
```go
type ExecutionMode int
const (
    Sequential ExecutionMode = iota  // calculadoras em ordem, erro propaga
    Parallel                          // goroutines, erro coletado e logado
)

type CalculationPhase struct {
    Name        string               // identificador para logging
    Mode        ExecutionMode        // Sequential ou Parallel
    Calculators []domain.TaxCalculator
}

type BillingEngineStruct struct {
    phases []CalculationPhase        // pipeline multi-fase genérico
}
```

**Funções chave:**
| Função | Descrição |
|--------|-----------|
| `BillingEnginePhased(phases ...)` | Construtor principal (C-001) — fases arbitrárias |
| `BillingEngineFull(pre, calcs, post)` | Compatibilidade — 3 fases internas |
| `BillingEngineOrdered(pre, calcs ...)` | Compatibilidade — 2 fases |
| `BillingEngine(calcs ...)` | Compatibilidade — 1 fase paralela |
| `Phase(name, mode, calcs ...)` | Helper para construção concisa de fases |
| `Process(ctx, input)` | Wrapper PRE_REFORMA → ProcessWithPhase |
| `ProcessWithPhase(ctx, input, filter)` | Itera fases → pós-processamento → consolida totais |
| `executeSequentialPhase(...)` | Executa calculadoras em ordem, injeta valores, propaga erros |
| `executeParallelPhase(...)` | Goroutines com WaitGroup, coleta erros via channel |
| `injectTributoValues(input, resItens)` | Injeta tributos calculados no input (chaves genéricas + específicas) |

**Pipeline SOP-013 (7 fases):**
```go
engine := BillingEnginePhased(
    Phase("IS", Sequential, isFilter),
    Phase("IPI", Sequential, ipiCalc),
    Phase("CBS", Sequential, cbsCalc),
    Phase("ICMS", Sequential, icmsAdapter),
    Phase("IBS+ISS+PISCOFINS", Parallel, ibsCalc, issCalc, pisCofinsAdapter),
    Phase("FUST", Sequential, fustCalc),
    Phase("FUNTTEL", Sequential, funttelCalc),
)
```

**Injeção inter-fase:** `injectTributoValues()` gera chaves padronizadas para cada tributo. ICMS recebe 3 chaves: `ICMS_VALOR` (genérica), `ITEM_ICMS_VALOR` (FUST/FUNTTEL) e `VALOR_EXCLUSAO_ICMS` (PIS/COFINS).

**Fonte:** `internal/calculator/engine.go:1-350`

#### `internal/domain/domain.go` — Interface Central do Domínio (DDD)

A interface `TaxCalculator` reside na camada mais interna da arquitetura (`domain/`), sem dependências de outros pacotes internos. As calculadoras em `legacy/` e `reforma/` a satisfazem via **structural typing** do Go — sem precisar importar o pacote `domain`.

```go
package domain

type TaxCalculator interface {
    Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error)
}
```

**Localização:** `internal/domain/domain.go:1-18`

**Fonte:** `internal/domain/domain.go:1-22`

#### `legacy_adapter.go` — Adaptadores

| Adapter | Entrada | Uso |
|---------|---------|-----|
| `icmsTaxAdapter` | Calculadora ICMS (legacy) | Converte `DocumentoFiscalSaida` → slice de `ItemDocumentoFiscalSaida` |
| `legacyTaxAdapter` | Calculadora PIS/COFINS + fonte ICMS opcional | Pré-calcula ICMS para exclusão da base, adapta saída de tributos |

**Função fábrica:** `LegacyAdapter(calc interface{}, icmsProvider ...interface{}) TaxCalculator`

**Fonte:** `internal/calculator/legacy_adapter.go:1-105`

### 3. Calculadoras Fiscais (`internal/legacy/`)

#### `ipi.go` — IPI Calculator

**Estrutura `IPICalculator`:**
```go
type IPICalculator struct {
    repo repository.TaxRepository
}
```

**Algoritmo:**
1. Rateio de frete, seguro, despesas e desconto entre itens
2. Para cada item, busca regra do banco (ou usa detalhes inline se overrides completos)
3. Cálculo Ad Valorem ou Ad Pauta

**Fonte:** `internal/legacy/ipi.go:1-124`

#### `icms.go` — ICMS Calculator

**Estrutura `ICMSCalculator`:**
```go
type ICMSCalculator struct {
    repo repository.TaxRepository
}
```

**Fluxo principal (reestruturado em 2026-06-20):**
1. Para cada item, obtém configuração efetiva via `getEffectiveTaxConfig()` (merge regra geral + exceção)
2. Se Simples Nacional → `calcularICMSSimples()` (equivalência CSOSN + alíquota efetiva)
3. Senão (Regime Normal):
   - Operação Interna → `calcularICMSOperacaoInterna()` (ICMS próprio + redução de base + FCP / ICMS-ST)
   - Operação Interestadual → `calcularICMSOperacaoInterestadual()` (ICMS interestadual + DIFAL + ICMS-ST)

**Métodos auxiliares:**
| Método | Responsabilidade |
|--------|-----------------|
| `calcularICMSSimples()` | CSOSN→CST, alíquota efetiva por anexo, fallback em erro |
| `calcularICMSOperacaoInterna()` | ICMS próprio com redução de base e FCP; ICMS-ST por CST 010 ou protocolo |
| `calcularICMSOperacaoInterestadual()` | ICMS interestadual (sempre); DIFAL se destino final; ICMS-ST se protocolo + MVA |
| `CalcularDIFAL()` | Diferencial de alíquotas usando `ICMSConfig` efetivo |
| `getEffectiveTaxConfig()` | Merge de `icms_rules` + `product_tax_exceptions` (override inteligente) |

**Fonte:** `internal/legacy/icms.go:1-400`

#### `pis_cofins.go` — PIS/COFINS Calculator

**Estrutura `PISCofinsCalculator`:**
```go
type PISCofinsCalculator struct {
    repo repository.TaxRepository
}
```

**Fluxo (migrado para banco em 2026-06-20):**
1. Para cada item, determina CST do regime e extrai `VALOR_EXCLUSAO_ICMS`
2. Consulta `GetFederalTaxRule(ctx, regime, cstPis, cstCofins)` no banco (`federal_tax_rules`)
3. Resolve alíquotas: banco > fallback para constantes (`defaultAliquotaPIS`=1.65%, `defaultAliquotaCOFINS`=7.6%)
4. Aplica `ExcluiICMSBase` da regra federal (flag controla "Tese do Século")
5. Delega cálculo para estratégia por CST (`GetPISStrategy` / `GetCOFINSStrategy`)
6. Loga fonte da alíquota (banco vs default) nos `MoreTextDetails` para auditoria

**Constantes de fallback:**
```go
const (
    defaultAliquotaPIS    = 1.65
    defaultAliquotaCOFINS = 7.6
)
```

**Fonte:** `internal/legacy/pis_cofins.go:1-180`

#### Strategy Implementations

- **PIS:** `PIS01_02`, `PIS03`, `PIS04`, `PIS05`, `PIS06`, `PIS49`, `PIS50To99`, `PIS99` — `internal/legacy/pis_strategies.go`
- **COFINS:** `COFINS01_02`, `COFINS03`, `COFINS04`, `COFINS05`, `COFINS06`, `COFINS49`, `COFINS50To99` — `internal/legacy/cofins_strategies.go`

**Cobertura de CSTs:** 100% (01, 02, 03, 04, 05, 06, 49, 50-99, 99). CSTs 04-06, 49, 50-99 retornam valor zero conforme legislação (monofásico, ST, alíquota zero, operações de crédito/suspensão).

**Testes:** 4 arquivos, 39 testes (`pis_strategies_test.go`, `cofins_strategies_test.go`, `pis_cofins_calculate_test.go`, `mock_repository_test.go`). A flag `ExcluiICMSBase` é validada com mock repository em 13 cenários (com/sem exclusão, fallback, múltiplos itens, edge cases).

## Padrões de Design Identificados

| Padrão | Local | Propósito |
|--------|-------|-----------|
| Strategy | `pis_strategies.go`, `cofins_strategies.go` | Comportamento de cálculo variável por CST |
| Adapter | `legacy_adapter.go` | Converte interfaces legacy para `TaxCalculator` |
| Pipeline (Chain of Responsibility) | `engine.go` | Pipeline multi-fase com modos Sequential/Parallel |
| Repository | `core-lib` | Acesso a dados com cache |
| Decorator | `core-lib` | `CachedTaxRepository` envolve `PostgresTaxRepository` |
| Factory | `pis_strategies.go:GetPISStrategy()` | Seleção de estratégia por código |
| Builder | `engine.go:BillingEnginePhased()` | Construção do motor com fases arbitrárias |
| Template Method | `reforma.go:computeIvaDual()` | Lógica compartilhada CBS/IBS (CBSCalculator e IBSCalculator) |
| Observer (Callback) | `pipeline_test.go` | Recording mocks verificam ordem de execução |

### 4. Reforma Tributária (`internal/reforma/`)

O pacote foi refatorado em 2026-06-22 (C-001) para suportar CBS e IBS como calculadoras independentes no pipeline SOP-013.

**Arquivos:**
| Arquivo | Conteúdo |
|---------|----------|
| `reforma.go` | Função interna `computeIvaDual()`, struct `ivaDualResult`, builders `buildCBSDetails`/`buildIBSDetails`, `ReformaCalculator` (legado CBS+IBS) |
| `cbs_calculator.go` | `CBSCalculator` — CBS apenas (Fase 2 sequencial) |
| `ibs_calculator.go` | `IBSCalculator` — IBS apenas (Fase 4 paralela) |

#### `cbs_calculator.go` — CBSCalculator

**Estrutura:**
```go
type CBSCalculator struct {
    repo repository.TaxRepository
}
```

Calcula apenas CBS (Contribuição sobre Bens e Serviços). Executa na Fase 2 do pipeline SOP-013 (sequencial, "por fora", antes do ICMS). Usa `computeIvaDual()` internamente e extrai apenas os campos CBS do resultado.

**Fonte:** `internal/reforma/cbs_calculator.go:1-64`

#### `ibs_calculator.go` — IBSCalculator

**Estrutura:**
```go
type IBSCalculator struct {
    repo repository.TaxRepository
}
```

Calcula apenas IBS (estadual + municipal). Executa na Fase 4 do pipeline SOP-013 (paralela com ISS e PIS/COFINS). Usa `computeIvaDual()` internamente e extrai apenas os campos IBS do resultado.

**Fonte:** `internal/reforma/ibs_calculator.go:1-65`

#### `reforma.go` — ReformaCalculator (legado) + lógica compartilhada

**Estrutura `computeIvaDual`:**
```go
func computeIvaDual(ctx, repo, item, ufDestino, municipioIBGE) *ivaDualResult
```

Função interna compartilhada por `CBSCalculator`, `IBSCalculator` e `ReformaCalculator`. Consulta `GetIvaDualRule` e retorna uma struct `ivaDualResult` com todos os valores computados (CBS, IBS estadual, IBS municipal). O cache Redis do `CachedTaxRepository` evita dupla consulta quando CBS e IBS são calculados em fases separadas.

**Algoritmo:**
1. Para cada item, extrai NCM, UF destino e município IBGE
2. Consulta `GetIvaDualRule()` — se `nil`, retorna `nil`
3. Calcula `fatorReducao = 1 − (percentualReducao / 100)`
4. Se `percentualReducao >= 100`: isenção total (EfetivamenteIsento = true)
5. CBS = base × (aliquotaCBS × fatorReducao / 100)
6. IBS = base × ((aliquotaIBSEstadual + aliquotaIBSMunicipal) × fatorReducao / 100)
7. IS = movido para `ISFilter` (Fase 0) — não faz parte do ReformaCalculator

**Fonte:** `internal/reforma/reforma.go:1-178`

**Testes:** `internal/reforma/reforma_test.go` — 7 cenários mantidos para o `ReformaCalculator` legado

### 5. ISS — Imposto sobre Serviços (`internal/legacy/iss.go`)

**Estrutura `ISSCalculator`:**
```go
type ISSCalculator struct{}
```

**Algoritmo:**
1. Para cada item, verifica `ITEM_LISTA_SERVICO` — se vazio, skip (mercadoria)
2. Obtém alíquota municipal (item > documento > default)
3. Valida range `[2%, 5%]` — `slog.Warn` se fora
4. Calcula: `ISS = Preço_Serviço × Alíquota / 100`
5. Verifica retenção na fonte (`ISS_RETIDO`)

**Constantes:** `issAliquotaMin = 2.0`, `issAliquotaMax = 5.0`, `issItemListaServicoTelecom = "1.05"`

**Fonte:** `internal/legacy/iss.go:1-140`

### 6. FUST e FUNTTEL — Contribuições de Telecom (`internal/legacy/fust.go`, `funttel.go`)

**Classificador compartilhado (`telecom.go`):**
```go
type TelecomClassifier struct{}
func (c *TelecomClassifier) MustCalculateFUST(item) bool // SCM/STFC → true, SVA → false
```

**FUSTCalculator:**
1. Classifica serviço (SCM/STFC vs SVA) via `TelecomClassifier`
2. Lê `ITEM_ICMS_VALOR`, `ITEM_PIS_VALOR`, `ITEM_COFINS_VALOR` dos detalhes
3. Calcula base líquida: `Valor_Serviço − ICMS − PIS − COFINS`
4. Se negativa → zero com `slog.Warn`
5. `FUST = Base × 0,01` (Lei 9.998/2000)

**FUNTTELCalculator:**
1. Mesma classificação e base do FUST
2. `FUNTTEL = Base × 0,005` (Lei 10.052/2000)

**Fonte:** `internal/legacy/fust.go:1-120`, `internal/legacy/funttel.go:1-100`, `internal/legacy/telecom.go:1-90`

**Testes:** `iss_test.go` (7), `fust_test.go` (6), `funttel_test.go` (4) — 17 cenários

### 7. Admin Fiscal (`internal/admin/`)

**Estrutura:**
- `models.go` — Modelos de regras fiscais (AdminFiscalRule)
- `repository.go` — Repository PostgreSQL para CRUD
- `service.go` — Service com validação de regras
- `service_test.go` — Testes do Admin Service

Implementado no GAP-004. Expõe `GET /v1/admin/tax-rates/iva-dual` para consulta e `POST /v1/admin/tax-rates/iva-dual` para upsert de alíquotas IVA Dual.

### 8. Créditos da Reforma Tributária (`internal/credit/`)

**Estrutura:**
- `engine.go` — Engine de cálculo de créditos (cash forward)
- `engine_test.go` — Testes do Credit Engine
- `models.go` — Modelos de crédito
- `supplier.go` — Fornecedor de créditos

Implementado no GAP-005. Expõe `POST /v1/credit/calculate`.

### 9. Simulação de Margem (`internal/simulation/`)

Projeção "what-if" para cenários fiscais (GAP-003). Expõe `POST /v1/simulate`.

### 10. Validação de Fornecedores (`internal/supplier/`)

**Estrutura:**
- `models.go` — Modelos de fornecedor
- `service.go` — Service de validação
- `service_test.go` — Testes do Supplier Service
- `store.go` — Store de fornecedores

Expõe `POST /v1/supplier/validate` e `GET /v1/supplier/:cnpj`.

### 11. TaxToken Snapshot (`internal/token/`)

Geração de token fiscal para snapshot de cálculos. Expõe `POST /v1/token/generate`.
