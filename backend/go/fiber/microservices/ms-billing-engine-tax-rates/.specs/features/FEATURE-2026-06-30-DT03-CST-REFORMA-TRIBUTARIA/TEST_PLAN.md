# TEST PLAN — DT-03: Tabela Oficial CST para CBS/IBS

**Feature:** FEATURE-2026-06-30-DT03-CST-REFORMA-TRIBUTARIA
**Versão:** 1.0
**Data:** 01 de Julho de 2026
**Status:** ✅ Executado

> 📋 **Propósito:** Este documento define a estratégia de testes para a resolução da DT-03. Ele complementa o [TEST_PLAN.md do projeto-base](../../business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/TEST_PLAN.md) (v2.0, todos os cenários executados). Baseline atual: **211+ testes em 25 arquivos**.

📄 **Referências:**
- [SPECS.md](./SPECS.md) — escopo e requisitos
- [TASKS.md](./TASKS.md) — tarefas de implementação
- [ARCHITECTURE.md](./ARCHITECTURE.md) — decisões de design
- [TEST_PLAN.md do projeto-base](../../business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/TEST_PLAN.md) — baseline e níveis de teste

---

## 1. Estratégia de Testes

| Nível | Escopo | Ferramenta |
|:---|:---|:---|
| **L1 — Unitário** | `GetCSTReforma()` repository, lógica de seleção de CST | `go test` + mocks |
| **L2 — Integração** | `CBSCalculator` e `IBSCalculator` com repository real | `go test` + PostgreSQL real (testcontainers ou local) |
| **L3 — Contrato** | API response: campo `CST` com 3 dígitos no response | `go test` + `httptest` |
| **L4 — Regressão** | Pipeline SOP-013 completo | `go test ./...` |

---

## 2. Linha de Base (Pré-Feature)

| Métrica | Valor |
|---------|-------|
| Testes totais | 211+ |
| Arquivos de teste | 25 |
| Cobertura `internal/reforma/` | 7 testes (`reforma_test.go`) |
| Cobertura pipeline | 22 cenários (`pipeline_test.go`) |

> ⚠️ **Regra:** Nenhuma tarefa é concluída se a linha de base regredir. `go test ./...` deve manter 100% de passes.

---

## 3. Cenários de Teste

### 3.1 L1 — Unitário: Repository `GetCSTReforma()`

**Arquivo:** `internal/legacy/mock_repository_test.go` (atualizar mock)

| ID | Cenário | Entrada | Esperado |
|:---|:---|:---|:---|
| TST-03.01 | Tributação integral — sem redução, sem isenção | `flags{EfetivamenteIsento: false, PercentualReducao: 0}` | `CSTReforma{CST: "000"}` |
| TST-03.02 | Isenção total (redução ≥ 100%) | `flags{EfetivamenteIsento: true}` | `CSTReforma{CST: "800"}` |
| TST-03.03 | Redução parcial de base de cálculo | `flags{PercentualReducao: 60}` | `CSTReforma{CST: "200"}` |
| TST-03.04 | Operação monofásica (NCM de combustível) | `flags{IsMonofasico: true}` | `CSTReforma{CST: "400"}` |
| TST-03.05 | Operação com diferimento | `flags{IsDiferimento: true}` | `CSTReforma{CST: "510"}` |
| TST-03.06 | Erro no repository (DB offline) | `repo retorna erro` | `nil, error` — motor usa fallback `"000"` |
| TST-03.07 | Nenhum CST encontrado (match vazio) | `repo retorna nil, nil` | `nil, nil` — motor usa fallback `"000"` |
| TST-03.08 | CST com alíquotas uniformes (setorial) | NCM financeiro | `CSTReforma{CST: "010"}` |

### 3.2 L1 — Unitário: Refactor `reforma.go`

**Arquivo:** `internal/reforma/reforma_test.go` (atualizar cenários existentes)

| ID | Cenário | Entrada | Esperado |
|:---|:---|:---|:---|
| TST-03.09 | CBS com CST `"000"` (tributação integral) | NCM normal, sem redução | `CBS calculado, CST=000` |
| TST-03.10 | CBS com CST `"800"` (isento) | NCM isento, redução 100% | `CBS=0, CST=800` |
| TST-03.11 | IBS com CST `"200"` (redução BC) | NCM cesta básica, redução 60% | `IBS reduzido, CST=200` |
| TST-03.12 | `computeIvaDual()` com erro no `GetCSTReforma()` | Repo mock retorna erro | `CSTEfetivo="000"` (fallback), CBS/IBS calculados normalmente |
| TST-03.13 | `computeIvaDual()` com `GetCSTReforma()` retornando `nil` | Repo mock retorna `nil, nil` | `CSTEfetivo="000"` (fallback) |
| TST-03.14 | Constantes `cstPadrao`/`cstIsento` removidas | `grep` no source | Nenhum resultado |

### 3.3 L2 — Integração: CBS/IBS com Repository Real

**Arquivo:** `internal/reforma/reforma_test.go`

| ID | Cenário | Nível | Esperado |
|:---|:---|:---|:---|
| TST-03.15 | `CBSCalculator` com `GetCSTReforma()` real | L2 | CST oficial no response |
| TST-03.16 | `IBSCalculator` com `GetCSTReforma()` real | L2 | CST oficial no response |
| TST-03.17 | Cache Redis: segunda chamada usa cache | L2 | Latência < 1ms na segunda chamada |

### 3.4 L3 — Contrato: API Response

**Arquivo:** `cmd/api/main_test.go` (ou teste de integração HTTP)

| ID | Cenário | Esperado |
|:---|:---|:---|
| TST-03.18 | `POST /v1/calculate` — resposta contém `CST` com 3 dígitos | `Tributos[].CST` é string de 3 caracteres (`"000"`, `"800"`, etc.) |
| TST-03.19 | `POST /v1/calculate` — resposta NÃO contém CST de 2 dígitos | Nenhum `Tributos[].CST` é `"01"` ou `"04"` |
| TST-03.20 | `POST /v1/simulate` — consistente com `/calculate` | CST de 3 dígitos na simulação |

### 3.5 L4 — Regressão: Pipeline SOP-013

**Arquivo:** `internal/calculator/pipeline_test.go`

| ID | Cenário | Esperado |
|:---|:---|:---|
| TST-03.21 | Pipeline 7 fases com CST oficial | 22 cenários existentes passando, CST de 3 dígitos nos resultados |
| TST-03.22 | Ordem das fases mantida | IS(F0)→IPI(F1)→CBS(F2)→ICMS(F3)→(IBS+ISS+PISCOFINS)(F4)→FUST(F5)→FUNTTEL(F6) |
| TST-03.23 | Injeção inter-fase preservada | `injectTributoValues()` funciona com novos CSTs |

---

## 4. Cenários de Borda e Erro

| ID | Cenário | Entrada | Comportamento Esperado |
|:---|:---|:---|:---|
| TST-03.24 | CST não mapeado (flag combination desconhecida) | Flags que não batem nenhum CST | Fallback `"000"` + `slog.Warn` |
| TST-03.25 | Tabela `cst_reforma` vazia (DB sem dados) | `init.sql` não executado | Fallback `"000"` para todos os itens |
| TST-03.26 | Concorrência: CBS e IBS consultam CST simultaneamente | Fase 2 (CBS) + Fase 4 (IBS) | Cache Redis evita dupla consulta ao DB |
| TST-03.27 | Admin Fiscal atualiza `cst_reforma` durante cálculo | `PUT /admin/tax-rates/cst` durante `POST /calculate` | Consistência: cálculo usa snapshot do momento da consulta |

### 4.1 Verificação Documental (Pós-Implementação)

| ID | Cenário | Esperado |
|:---|:---|:---|
| TST-DOC.01 | SPECS.md — status ✅ Concluído | Campo `Status` no header |
| TST-DOC.02 | TASKS.md — todas as tarefas marcadas ✅ | Tabela de resumo com 6 ✅ |
| TST-DOC.03 | TEST_PLAN.md — status ✅ Executado | Campo `Status` no header |
| TST-DOC.04 | ADR-010, ADR-011, ADR-012 — status ✅ Aceito | Campo `Status` em cada ADR |
| TST-DOC.05 | `feature-roadmap.md` — DT-03 ✅ Resolvida | Linha DT-03 na tabela |
| TST-DOC.06 | `adrs/INDEX.md` — status atualizado | ADR-010/011/012 como ✅ Aceito |
| TST-DOC.07 | `CHANGELOG.md` — entrada da feature | Entrada DT-03 |
| TST-DOC.08 | `erd.md` — tabela `cst_reforma` | Nova tabela no diagrama |

---

## 5. Matriz de Rastreabilidade

| Requisito (SPECS.md) | Cenários de Teste |
|:---|:---|
| S-01 (Tabela `cst_reforma`) | TST-03.01 a TST-03.08 |
| S-02 (Popular init.sql) | TST-03.25 (tabela vazia) |
| S-03 (`GetCSTReforma()`) | TST-03.01 a TST-03.08, TST-03.17 |
| S-04 (Refatorar reforma.go) | TST-03.09 a TST-03.14 |
| S-05 (Flags comportamentais) | TST-03.04, TST-03.05 |
| S-06 (Atualizar testes) | TST-03.09 a TST-03.23 |

---

## 6. Suíte de Regressão

```bash
# Executar antes de cada commit
go test ./internal/reforma/...        # CBS, IBS, computeIvaDual
go test ./internal/calculator/...     # Pipeline 7 fases
go test ./internal/legacy/...         # Mock repository
go test ./...                         # Full suite (211+ testes)

# Verificar remoção de constantes
grep -r "cstPadrao\|cstIsento" internal/  # Deve retornar vazio
```

---

## 7. Resumo

| Métrica | Valor |
|---------|-------|
| Cenários novos | **27** (TST-03.01 a TST-03.27) |
| Cenários de verificação documental | **8** (TST-DOC.01 a TST-DOC.08) |
| Cenários atualizados | ~14 (reforma_test.go, pipeline_test.go) |
| Níveis de teste | L1 (14), L2 (3), L3 (3), L4 (3), Borda (4), Doc (8) |
| Baseline esperada pós-feature | **238+ testes** (211 + 27 novos) |
