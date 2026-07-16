# Implementação F-006 e F-007 + C-002 — IS Pré-Filtro + IBS Circuit Breaker + Schema SQL

**Data:** 2026-06-22  
**Skill:** golang-pro + spec-miner  
**Features:** [FEATURE-2026-06-21](../features/FEATURE-2026-06-21.md) — F-006, F-007, C-002

## F-006: IS como Pré-Filtro Independente

### Arquivos criados/modificados

| Arquivo | Ação | Descrição |
|---------|------|-----------|
| `internal/legacy/is_filter.go` | Criado | 175 linhas. `ISFilter` struct — pré-calculadora Fase 0. Consulta `ncm_seletivo`, flag `isento_is`, auditoria mesmo IS=0. |
| `internal/legacy/is_filter_test.go` | Criado | 8 testes unitários: bebida alcoólica incide, telecom não incide, isento flag doc/item, cigarros 100%, refrigerantes 25%, múltiplos itens, erro repo. |
| `internal/reforma/reforma.go` | Modificado | Removido bloco IS (34 linhas). `ReformaCalculator` agora calcula APENAS CBS + IBS. |
| `internal/reforma/reforma_test.go` | Modificado | `TestReformaCalculator_ImpostoSeletivo` → `TestReformaCalculator_CBS_IBS_SemIS` — verifica que IS NÃO é mais calculado. |
| `cmd/api/main.go` | Modificado | `ISFilter` adicionado como primeiro item dos `preCalculators` (Fase 0, antes do IPI). |

### Regras implementadas

- **BR-TAX-INF-005:** Consulta à tabela `ncm_seletivo` para verificar se NCM está sujeito ao IS
- **BR-TAX-CONS-010:** IS como pré-filtro obrigatório ANTES da CBS (executado na Fase 0)
- **SOP-003 Step 2:** Flag `isento_is` — se true, IS = 0 independente do NCM
- **SOP-003 Step 3:** Tabela `ncm_seletivo` com categorias (BEBIDAS_ALCOOLICAS, CIGARROS, REFRIGERANTES)
- **SOP-003 Step 4:** Registro de auditoria mesmo quando IS = 0 (flag `is_exempt`)
- **SOP-003 edge case:** NCM não listado → IS = 0, registrado com `is_exempt=true`

### Categorias de IS (ncm_seletivo)

| NCM | Categoria | Alíquota IS |
|-----|-----------|-------------|
| 22030000 | BEBIDAS_ALCOOLICAS | 50% |
| 24022000 | CIGARROS | 100% |
| 22021000 | REFRIGERANTES | 25% |

### Estrutura do ISFilter

```go
type ISFilter struct {
    repo repository.TaxRepository // consulta ncm_seletivo
}

func (f *ISFilter) Calculate(ctx context.Context, input models.DocumentoFiscalEntrada) ([]models.ItemDocumentoFiscalSaida, error) {
    // 1. Verifica flag isento_is (documento > item)
    // 2. Consulta ncm_seletivo via repo.GetNCMSeletivo()
    // 3. Se NCM na tabela: IS = Valor × Aliquota_IS_Categoria
    // 4. Se NCM não na tabela: IS = 0 (auditado)
    // 5. Se isento_is: IS = 0 (auditado)
}
```

### Modificações na shared lib (taxnexus-billing-core-lib)

| Arquivo | Ação |
|---------|------|
| `repository/entities.go` | Adicionado `NCMSeletivoRule` struct |
| `repository/contracts.go` | Adicionado `GetNCMSeletivo(ctx, ncm)` ao `TaxRepository` |
| `repository/postgres_repository.go` | Implementado `GetNCMSeletivo` — query na `ncm_seletivo` |
| `repository/cached_tax_repository.go` | Implementado `GetNCMSeletivo` com cache Redis `tax:ncm_seletivo:{ncm}` |

---

## F-007: IBS Circuit Breaker

### Arquivos criados

| Arquivo | Ação | Descrição |
|---------|------|-----------|
| `internal/circuitbreaker/circuit_breaker.go` | Criado | 180 linhas. State machine CLOSED→OPEN→HALF_OPEN com `sync.Mutex`. Config: 3 falhas/60s→OPEN, HALF_OPEN após 5min. |
| `internal/circuitbreaker/circuit_breaker_test.go` | Criado | 7 testes: sucesso fechado, abre após falhas, rejeita quando aberto, HALF_OPEN→CLOSED, HALF_OPEN→OPEN, window reset, sucesso reseta contador. |
| `internal/ibsclient/client.go` | Criado | 240 linhas. `IBSRateFetcher` interface + 4 implementações. |
| `internal/ibsclient/client_test.go` | Criado | 5 testes unitários (+ 2 skip integração). |
| `cmd/api/main.go` | Modificado | IBS Client instanciado com env var `IBS_API_BASE_URL`. |

### Arquitetura IBS Client

```
FallbackIBSClient (tenta API, fallback DB)
  ├── CircuitBreakerIBSClient (3 falhas/60s → OPEN)
  │     └── CachedIBSClient (Redis, TTL 24h)
  │           └── HTTPIBSClient (API Comitê Gestor)
  └── TaxRepository.GetIvaDualRule() (fallback banco)
```

### Implementações de IBSRateFetcher

| Implementação | Propósito | Fonte |
|---------------|-----------|-------|
| `HTTPIBSClient` | Chamada HTTP `GET /api/v1/rates?ibge_code={code}` | `API_COMITE_GESTOR` |
| `CachedIBSClient` | Cache Redis `ibs:rate:{ibge_code}:{date}` TTL 24h | `CACHE` |
| `CircuitBreakerIBSClient` | 3 falhas/60s→OPEN, cache expirado como fallback | `FALLBACK_CIRCUIT_OPEN` |
| `FallbackIBSClient` | API → DB fallback via `GetIvaDualRule()` | `FALLBACK_DB` |

### Regras implementadas

- **BR-TAX-ACT-001:** Circuit breaker com 3 estados — protege contra falhas em cascata
- **BR-TAX-CONS-009:** Cache Redis com chave `ibs:rate:{ibge_code}:{date}` e TTL = 24h
- **SOP-002 Step 3:** Chamada HTTP `GET /api/v1/rates?ibge_code={code}`
- **SOP-002 Step 8:** Flag `rate_source` nos detalhes: CACHE, API_COMITE_GESTOR, FALLBACK_CIRCUIT_OPEN, FALLBACK_DB
- **SOP-014:** Estados CLOSED→OPEN (≥3 falhas)→HALF_OPEN (5min)→CLOSED
- **Gap G2:** Enquanto API não publicada, fallback DB ativo via `GetIvaDualRule()`

### Estados do Circuit Breaker

```
CLOSED ──(3 falhas em 60s)──▶ OPEN ──(5 min timeout)──▶ HALF_OPEN
  ▲                                                         │
  └────────────────(sucesso)────────────────────────────────┘
  │                                                         │
  └────────────────(falha)──────────────────────────────────┘
                                     (volta a OPEN)
```

---

## C-002: Schema SQL

### Tabelas criadas em `data/init.sql`

| Tabela | Propósito | Dados de exemplo |
|--------|-----------|-----------------|
| `ncm_seletivo` | NCMs sujeitos ao Imposto Seletivo | 6 registros (cerveja, vinho, destilados, cigarros, xaropes, refrigerantes) |
| `cbs_rates` | Alíquotas CBS por classe tributária | Estrutura criada, dados pendentes (Gap G1) |
| `iss_rates` | Alíquotas ISS por município | 5 registros (SP, RJ, DF, BH, Curitiba — alíquota 5%) |

---

## Resumo de Testes

| Pacote | Testes | Status |
|--------|--------|--------|
| `internal/legacy` (ISFilter) | 8 | ✅ PASS |
| `internal/circuitbreaker` | 7 | ✅ PASS |
| `internal/ibsclient` | 5 + 2 skip | ✅ PASS |
| `internal/reforma` (refatorado) | 6 | ✅ PASS |
| Todos os pacotes existentes | ~140 | ✅ PASS |

**Total:** ~160 testes passando, `go vet` limpo, `go build` sem erros.
