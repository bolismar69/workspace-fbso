# SPECS — Resolução DT-03: Tabela Oficial CST para CBS/IBS

**Feature:** FEATURE-2026-06-30-DT03-CST-REFORMA-TRIBUTARIA
**Microserviço:** `ms-billing-engine-tax-rates`
**Dívida Técnica:** [DT-03](../../product/feature-roadmap.md#d%C3%ADvidas-t%C3%A9cnicas) — CST da Reforma Tributária usa valores provisórios (`01`/`04`)
**Versão:** 1.0
**Data:** 01 de Julho de 2026
**Status:** ✅ Concluído

> 📋 **Propósito:** Este documento define o escopo, requisitos e decisões de design para substituir os valores provisórios de CST (`01`/`04`) pela tabela oficial de Classificação Tributária publicada pela RFB (LC 214/2025).

📄 **Referências:**
- Projeto-base: [PRJ-FIN-2026-0001](../../business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/SPECS.md) — Fases 0-1-2 concluídas
- Requisitos de negócio: [BR-04 (Transparência "Por Fora")](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/REQUIREMENTS.md)
- Dívida técnica: [DT-03](../../product/feature-roadmap.md)
- Código afetado: `internal/reforma/reforma.go:27-31`
- Dados oficiais: `CST_cClassTrib_2025-10-03_Public_verde.xlsx`, `classificacao_tributaria(1).csv`, `classificacao_tributaria(1).json`

---

## 1. Problema

### Situação Atual

O arquivo `internal/reforma/reforma.go` usa **2 constantes hardcoded** para CST:

```go
const (
    cstPadrao  = "01"   // tributação normal
    cstIsento  = "04"   // isenção/redução
)
```

A lógica de seleção é binária:
- Se `EfetivamenteIsento` → CST = `"04"`
- Se `PercentualReducao > 0` → CST = `"04"`
- Senão → CST = `"01"`

### Por que é um problema

1. **Não conformidade regulatória:** A LC 214/2025 define **18 CSTs** distintos para CBS/IBS, cobrindo cenários como monofásica, diferimento, crédito presumido, alíquotas uniformes setoriais — nenhum modelado pelo código atual.
2. **Impossibilidade de auditoria fiscal:** Com apenas `01`/`04`, não é possível distinguir uma operação monofásica de uma operação com redução de alíquota no documento fiscal.
3. **Bloqueio para features futuras:** DT-04 (créditos) e split payment (GAP-006) dependem de CST correto para calcular créditos apropriáveis.

---

## 2. O que Será Construído

### 2.1 Escopo

| Item | Descrição | Prioridade |
|------|-----------|------------|
| **S-01** | Criar tabela `cst_reforma` com 18 CSTs + 164 CCTs | 🔴 Must |
| **S-02** | Popular `data/init.sql` com dados oficiais da RFB | 🔴 Must |
| **S-03** | Adicionar `GetCSTReforma()` ao `TaxRepository` | 🔴 Must |
| **S-04** | Refatorar `reforma.go` para consultar CST do banco em vez de constantes | 🔴 Must |
| **S-05** | Mapear flags comportamentais (monofásica, diferimento, crédito) no modelo | 🟡 Should |
| **S-06** | Atualizar testes existentes com mock do novo repository method | 🔴 Must |

### 2.2 Fora do Escopo

- **Cálculo de créditos (DT-04):** A tabela `cst_reforma` habilita DT-04, mas não a implementa
- **Split payment (GAP-006):** Usa CST correto, mas a lógica de split é separada
- **Validação de CST no input:** O consumidor pode enviar CST no request? Escopo futuro
- **CCT (Classificação Tributária) como chave primária:** Nesta feature, usamos CST (3 dígitos). O CCT (6 dígitos) fica como coluna de referência para expansão futura

---

## 3. Dados de Entrada

### 3.1 Fonte Oficial

| Arquivo | Formato | Registros | Origem |
|---------|---------|-----------|--------|
| `CST_cClassTrib_2025-10-03_Public_verde.xlsx` | Excel | 164 CCTs | RFB — LC 214/2025 |
| `classificacao_tributaria(1).csv` | CSV (delim `;`) | 164 CCTs | Mesmo conteúdo |
| `classificacao_tributaria(1).json` | JSON | 164 CCTs | Mesmo conteúdo |

### 3.2 Mapeamento CST (3 dígitos oficiais)

| CST Oficial | Descrição | CCTs | Equivalente atual | Flags |
|-------------|-----------|------|-------------------|-------|
| `000` | Tributação integral | 5 | `01` (parcial) | — |
| `010` | Alíquotas uniformes | 2 | — | aliquota_uniforme |
| `011` | Alíquotas uniformes (var) | 5 | — | aliquota_uniforme |
| `200` | Redução base de cálculo | 3 | `04` (parcial) | reducao_bc |
| `220` | Redução de alíquota | 3 | `04` (parcial) | reducao_aliquota |
| `221` | Redução de alíquota (var) | 2 | `04` (parcial) | reducao_aliquota |
| `222` | Redução de alíquota (var) | 2 | `04` (parcial) | reducao_aliquota |
| `400` | Monofásica normal | 2 | — | monofasica |
| `410` | Monofásica com retenção | 2 | — | monofasica, retencao |
| `510` | Diferimento | 2 | — | diferimento |
| `515` | Diferimento (var) | 2 | — | diferimento |
| `550` | Diferimento ZFM | 2 | — | diferimento, zfm |
| `620` | Crédito presumido | 1 | — | credito_presumido |
| `800` | Sem tributação | 5 | `04` (parcial) | isento |
| `810` | Sem tributação (var) | 2 | — | isento, imunidade |
| `811` | Sem tributação (var) | 2 | — | isento, imunidade |
| `820` | Sem tributação (var) | 1 | — | isento |
| `830` | Sem tributação (var) | 1 | — | isento, nao_contribuinte |

---

## 4. Design Técnico

### 4.1 Nova Tabela: `cst_reforma`

```sql
CREATE TABLE IF NOT EXISTS cst_reforma (
    id              SERIAL PRIMARY KEY,
    cst             CHAR(3) NOT NULL,          -- Código Situação Tributária (000-830)
    cct             CHAR(6) NOT NULL UNIQUE,   -- Código Classificação Tributária
    descricao_cst   TEXT NOT NULL,             -- Descrição do CST
    descricao_cct   TEXT NOT NULL,             -- Descrição do CCT
    exige_tributacao BOOLEAN DEFAULT TRUE,
    reducao_bc      BOOLEAN DEFAULT FALSE,
    reducao_aliquota BOOLEAN DEFAULT FALSE,
    transferencia_credito BOOLEAN DEFAULT FALSE,
    diferimento     BOOLEAN DEFAULT FALSE,
    monofasica      BOOLEAN DEFAULT FALSE,
    credito_presumido BOOLEAN DEFAULT FALSE,
    percentual_reducao_ibs DECIMAL(5,2) DEFAULT 0,
    percentual_reducao_cbs DECIMAL(5,2) DEFAULT 0,
    tipo_aliquota   VARCHAR(50),
    url_legislacao  TEXT,
    created_at      TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_cst_reforma_cst ON cst_reforma(cst);
```

### 4.2 Repository Method

```go
// GetCSTReforma retorna o CST oficial para CBS/IBS com base no contexto fiscal.
// Parâmetros: NCM, UF destino, e flags da operação (percentual de redução, etc.)
// Fallback: se nenhuma regra específica, retorna "000" (tributação integral).
GetCSTReforma(ctx context.Context, ncm string, flags CSTFlags) (*CSTReforma, error)
```

### 4.3 Refactor em `reforma.go`

**Antes:**
```go
result.CSTEfetivo = cstPadrao   // "01"
if efetivamenteIsento {
    result.CSTEfetivo = cstIsento  // "04"
}
```

**Depois:**
```go
cstRule, err := repo.GetCSTReforma(ctx, ncm, CSTFlags{
    EfetivamenteIsento: efetivamenteIsento,
    PercentualReducao:  rule.PercentualReducao,
    UFDestino:          ufDestino,
})
if err != nil || cstRule == nil {
    result.CSTEfetivo = "000" // fallback: tributação integral
} else {
    result.CSTEfetivo = cstRule.CST
}
```

### 4.4 Diagrama de Integração

```
POST /v1/calculate
    │
    ▼
BillingEnginePhased (SOP-013)
    │
    ├─ F2 (CBS): CBSCalculator.Calculate()
    │   └─ computeIvaDual() → repo.GetIvaDualRule()
    │   └─ repo.GetCSTReforma()         ← NOVO
    │
    └─ F4 (IBS): IBSCalculator.Calculate()
        └─ computeIvaDual() → repo.GetIvaDualRule()
        └─ repo.GetCSTReforma()         ← NOVO
```

---

## 5. Impacto

### 5.1 Arquivos Modificados

| Arquivo | Tipo de Mudança |
|---------|-----------------|
| `data/init.sql` | +1 tabela `cst_reforma` + INSERTs |
| `internal/reforma/reforma.go` | Remover `const cstPadrao/cstIsento`, integrar `GetCSTReforma()` |
| `internal/reforma/cbs_calculator.go` | Passar `repo` para `computeIvaDual()` |
| `internal/reforma/ibs_calculator.go` | Passar `repo` para `computeIvaDual()` |
| `internal/reforma/reforma_test.go` | Atualizar mocks |
| `internal/calculator/pipeline_test.go` | Atualizar mocks |
| `.specs/product/feature-roadmap.md` | Marcar DT-03 como ✅ Resolvida |
| `.specs/domain/domain.md` | Adicionar referência à tabela `cst_reforma` |
| `.specs/architecture/erd.md` | Adicionar tabela `cst_reforma` ao diagrama |
| `.specs/api/tax-rates-api.yaml` | Documentar CST de 3 dígitos |
| `.specs/CHANGELOG.md` | Adicionar entrada da feature |
| `.specs/architecture/adrs/adr-010/011/012.md` | Atualizar status de Proposto → Aceito |
| `.specs/architecture/adrs/INDEX.md` | Atualizar status dos ADRs |
| `.specs/features/FEATURE-.../SPECS.md` | Atualizar status para ✅ Concluído |
| `.specs/features/FEATURE-.../TASKS.md` | Marcar todas as tarefas como concluídas |
| `.specs/features/FEATURE-.../TEST_PLAN.md` | Marcar como executado |

### 5.2 Não Quebra Compatibilidade

- A resposta da API **não muda de estrutura** — apenas o campo `CST` dentro de `TributosItemDocumentoFiscalSaida` passa a ter valores oficiais (3 dígitos) em vez de provisórios (2 dígitos).
- Consumidores da API devem ser notificados da mudança de `01`→`000` e `04`→`800` (ou o CST correspondente).

---

## 6. Referências Cruzadas

### 6.1 Documentos do Projeto-base

| Documento | Relevância |
|-----------|-----------|
| [PRJ-FIN-2026-0001 ARCHITECTURE.md](../../business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/ARCHITECTURE.md) | Seção 6.2 — GAPs originais. Seção 6.3 — DT-03 listada |
| [PRJ-FIN-2026-0001 SPECS.md](../../business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/SPECS.md) | BR-04 (Transparência), BR-06 (Token) |
| [PRJ-FIN-2026-0001 TASKS.md](../../business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/TASKS.md) | DT-03 não estava nas 65 tarefas originais |

### 6.2 Documentos do Microserviço

| Documento | Relevância |
|-----------|-----------|
| [feature-roadmap.md](../../product/feature-roadmap.md) | DT-03, DT-04 |
| [architecture.md](../../architecture/architecture.md) | Pipeline SOP-013, camadas DDD |
| [integrations.md](../../architecture/integrations.md) | Variáveis de ambiente, dependências |
| [erd.md](../../architecture/erd.md) | Modelo ER — nova tabela `cst_reforma` |

### 6.3 Documentos Corporativos

| Documento | Relevância |
|-----------|-----------|
| [REQUIREMENTS.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/REQUIREMENTS.md) | BR-04: Transparência "Por Fora" |
| [PROJECT-CHARTER.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/PROJECT-CHARTER.md) | Escopo do programa |

---

## 7. Critérios de Aceitação

- [ ] Tabela `cst_reforma` criada com 164 registros (18 CSTs × N CCTs)
- [ ] `GetCSTReforma()` implementado no `TaxRepository` (PostgreSQL + cache Redis)
- [ ] `reforma.go` sem constantes `cstPadrao`/`cstIsento`
- [ ] CBS e IBS retornam CST oficial de 3 dígitos no campo `CST`
- [ ] `go test ./internal/reforma/...` — 100% passes
- [ ] `go test ./internal/calculator/...` — 100% passes (pipeline tests)
- [ ] `go test ./...` — sem regressões (211+ testes)
- [ ] `go vet ./...` — sem warnings
- [ ] DT-03 marcada como ✅ Resolvida no `feature-roadmap.md`
- [ ] Consumidores da API notificados sobre mudança de formato do CST (2→3 dígitos)
- [ ] **📋 Documentação como Entregável:** Todos os documentos da feature atualizados para status ✅ Concluído (este SPECS.md, TASKS.md, TEST_PLAN.md, ARCHITECTURE.md)
- [ ] ADRs registrados no catálogo canônico com status atualizado de 🔨 Proposto → ✅ Aceito: [ADR-010](../../architecture/adrs/adr-010.md), [ADR-011](../../architecture/adrs/adr-011.md), [ADR-012](../../architecture/adrs/adr-012.md)
- [ ] [INDEX.md](../../architecture/adrs/INDEX.md) do catálogo de ADRs atualizado
- [ ] Documentos do microserviço atualizados: `feature-roadmap.md`, `domain/domain.md`, `erd.md`, `tax-rates-api.yaml`, `CHANGELOG.md`
- [ ] [ARCHITECTURE.md](../../architecture/architecture.md) do microserviço revisado e atualizado se necessário
