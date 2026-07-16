# TASKS — DT-03: Tabela Oficial CST para CBS/IBS

**Feature:** FEATURE-2026-06-30-DT03-CST-REFORMA-TRIBUTARIA
**Versão:** 1.0
**Data:** 01 de Julho de 2026
**Status:** ✅ Concluído

> 📋 **Propósito:** Este documento decompõe a feature em tarefas granulares com Definition of Done. Use-o como checklist de execução.

📄 **Referências:**
- [SPECS.md](./SPECS.md) — escopo completo
- [ARCHITECTURE.md](./ARCHITECTURE.md) — decisões de design (3 ADRs; registro canônico em [../../architecture/adrs/INDEX.md](../../architecture/adrs/INDEX.md))
- [TEST_PLAN.md](./TEST_PLAN.md) — cenários de teste
- [TASKS.md do projeto-base](../../business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/TASKS.md) — metodologia DoD

---

## 1. Definição de Pronto (Definition of Done)

| # | Critério | Evidência |
|:---|:---|:---|
| DoD-1 | Código compilando (`go build ./...`) | CI verde |
| DoD-2 | Testes passando para o novo código (`go test ./...`) | Output do `go test` |
| DoD-3 | `go vet ./...` sem warnings | CI verde |
| DoD-4 | Especificação OpenAPI atualizada se schema mudar | Diff do YAML |
| DoD-5 | DT-03 marcada como ✅ Resolvida no `feature-roadmap.md` | Commit |
| DoD-6 | **Documentação da feature concluída:** SPECS, TASKS, TEST_PLAN, ARCHITECTURE com status ✅ | Arquivos atualizados |
| DoD-7 | ADRs (010, 011, 012) com status ✅ Aceito no catálogo canônico | [adrs/INDEX.md](../../architecture/adrs/INDEX.md) |
| DoD-8 | Documentos do microserviço atualizados (domain, erd, changelog, api) | Commits |

---

## 2. Grafo de Dependências

```
T-01 (DDL + INSERTs)
  └─ T-02 (Repository interface + impl)
       ├─ T-03 (Refactor reforma.go)
       │    └─ T-05 (Atualizar testes)
       └─ T-04 (Cache Redis)
            └─ T-05 (Atualizar testes)
                 └─ T-06 (Sanity Check — Protocolo de Checagem Pós-Implementação)
                      ├─ Se falhas → gerar FEEDBACK_ERRORS.md → voltar a T-01
                      └─ Se OK → T-07 (Documentação da Feature)
                           └─ T-08 (Documentação da Solução — Relatório de Execução)
                                └─ T-09 (Auditoria compliance docs)
                                     ├─ Se divergências → voltar a T-07
                                     └─ Se OK → feature concluída ✅
```

---

## 3. Tarefas

### T-01: ✅ Criar tabela `cst_reforma` e popular dados

**Prioridade:** 🔴 Must
**Estimativa:** 2-3 horas
**Dependências:** Nenhuma

**Descrição:**
1. Adicionar `CREATE TABLE cst_reforma` ao `data/init.sql`
2. Criar índice `idx_cst_reforma_cst` para consultas por CST
3. Importar 164 registros do arquivo `classificacao_tributaria(1).csv` como INSERTs
4. Mapear colunas do CSV→colunas da tabela:
   - `Código da Situação Tributária` → `cst`
   - `Código da Classificação Tributária` → `cct`
   - `Descrição da Situação Tributária` → `descricao_cst`
   - `Descrição do Código da Classificação Tributária` → `descricao_cct`
   - Flags (`Exige Tributação`, `Redução BC CST`, etc.) → colunas booleanas
   - `Percentual Redução IBS`/`Percentual Redução CBS` → `percentual_reducao_ibs`/`_cbs`
   - `Url da Legislação` → `url_legislacao`

**Arquivos:**
- `data/init.sql`

**DoD:** `psql -f data/init.sql` cria a tabela com 164 registros sem erros

---

### T-02: ✅ Adicionar `GetCSTReforma()` ao Repository

**Prioridade:** 🔴 Must
**Estimativa:** 2-3 horas
**Dependências:** T-01

**Descrição:**
1. Adicionar struct `CSTReforma` aos models em `taxnexus-billing-core-lib/models/tax_models.go`
2. Adicionar struct `CSTFlags` para parâmetros de consulta
3. Adicionar método `GetCSTReforma(ctx, ncm, flags) (*CSTReforma, error)` à interface `TaxRepository`
4. Implementar `GetCSTReforma()` no `PostgresTaxRepository`:
   - Query: `SELECT * FROM cst_reforma WHERE cst = $1` (ou lógica mais complexa com flags)
   - Fallback: retornar `nil, nil` se nenhum match (motor usa `"000"` como default)
5. Implementar cache no `CachedTaxRepository` (decorator existente)

**Lógica de seleção de CST:**
```go
// Pseudocódigo
func (r *PostgresTaxRepository) GetCSTReforma(ctx, ncm string, flags CSTFlags) (*CSTReforma, error) {
    var cst string
    if flags.EfetivamenteIsento {
        cst = "800" // Sem tributação
    } else if flags.IsMonofasico {
        cst = "400" // Monofásica
    } else if flags.IsDiferimento {
        cst = "510" // Diferimento
    } else if flags.PercentualReducao > 0 {
        cst = "200" // Redução base de cálculo
    } else {
        cst = "000" // Tributação integral
    }
    return queryOne("SELECT * FROM cst_reforma WHERE cst = $1", cst)
}
```

**Arquivos:**
- `../../../libs/go-native/taxnexus-billing-core-lib/models/tax_models.go`
- `../../../libs/go-native/taxnexus-billing-core-lib/repository/tax_repository.go`

**DoD:** `GetCSTReforma()` retorna CST oficial para cenários de teste conhecidos

---

### T-03: ✅ Refatorar `reforma.go` — Remover Constantes Hardcoded

**Prioridade:** 🔴 Must
**Estimativa:** 3-4 horas
**Dependências:** T-02

**Descrição:**
1. Remover `const cstPadrao = "01"` e `const cstIsento = "04"`
2. Modificar `computeIvaDual()` para receber `repo repository.TaxRepository` como parâmetro e chamar `GetCSTReforma()`
3. Atualizar `CBSCalculator.Calculate()` para passar `repo` para `computeIvaDual()`
4. Atualizar `IBSCalculator.Calculate()` para passar `repo` para `computeIvaDual()`
5. Atualizar `ReformaCalculator.Calculate()` (legado) para consistência
6. Fallback: se `GetCSTReforma()` retornar `nil` ou erro, usar `"000"` (tributação integral)

**Antes:**
```go
result.CSTEfetivo = cstPadrao  // "01"
if efetivamenteIsento {
    result.CSTEfetivo = cstIsento  // "04"
}
```

**Depois:**
```go
cstFlags := CSTFlags{
    EfetivamenteIsento: efetivamenteIsento,
    PercentualReducao:  rule.PercentualReducao,
}
cstRule, err := repo.GetCSTReforma(ctx, ncm, cstFlags)
if err != nil {
    slog.Warn("Erro ao consultar CST, usando fallback", "err", err)
}
if cstRule != nil {
    result.CSTEfetivo = cstRule.CST
} else {
    result.CSTEfetivo = "000"
}
```

**Arquivos:**
- `internal/reforma/reforma.go`
- `internal/reforma/cbs_calculator.go`
- `internal/reforma/ibs_calculator.go`

**DoD:** `grep -r "cstPadrao\|cstIsento" internal/` não encontra resultados

---

### T-04: ⚠️ Cache Redis para `cst_reforma`

**Prioridade:** 🟡 Should
**Estimativa:** 1-2 horas
**Dependências:** T-02

**Descrição:**
1. Adicionar cache no `CachedTaxRepository.GetCSTReforma()`:
   - Chave: `cst:reforma:{cst}`
   - TTL: 24h (consistente com `iva_dual_rules`)
2. Invalidar cache no Admin Fiscal quando regra for atualizada

**Nota:** Se o padrão `CachedTaxRepository` já usar decorator genérico, pode ser automático. Verificar implementação atual.

**Arquivos:**
- `../../../libs/go-native/taxnexus-billing-core-lib/repository/tax_repository.go`
- `internal/admin/service.go` (invalidação)

**DoD:** Segunda chamada a `GetCSTReforma()` com mesmos parâmetros atinge cache Redis

---

### T-05: ✅ Atualizar Testes

**Prioridade:** 🔴 Must
**Estimativa:** 2-3 horas
**Dependências:** T-03

**Descrição:**
1. Atualizar `mock_repository_test.go` — adicionar mock de `GetCSTReforma()`
2. Atualizar `reforma_test.go` — cenários com CST oficial (3 dígitos)
3. Atualizar `pipeline_test.go` — garantir que pipeline tests esperam CST de 3 dígitos
4. Adicionar testes específicos:
   - **Cenário 1:** Tributação integral → CST `"000"`
   - **Cenário 2:** Isenção total → CST `"800"`
   - **Cenário 3:** Redução parcial → CST `"200"`
   - **Cenário 4:** Erro no repository → fallback `"000"`
   - **Cenário 5:** Repository retorna `nil` → fallback `"000"`
5. Rodar `go test ./...` e garantir 0 falhas (baseline: 211+ testes)

**Arquivos:**
- `internal/reforma/reforma_test.go`
- `internal/calculator/pipeline_test.go`
- `internal/legacy/mock_repository_test.go`

**DoD:** `go test ./...` — 211+ testes passando, 0 falhas

---

### T-06: ✅ Protocolo de Checagem Pós-Implementação (Sanity Check)

**Prioridade:** 🔴 Must
**Estimativa:** 1-2 horas
**Dependências:** T-05
**Ciclo:** Se falhas → gerar `FEEDBACK_ERRORS.md` → reportar ao desenvolvedor → voltar a T-01; se OK → ✅

**Descrição:**
Validar a qualidade do código implementado nas tarefas T-01 a T-05 antes de iniciar a fase de documentação. Esta tarefa é um **gate de qualidade técnica** — a feature não avança para documentação até que todas as verificações de código passem.

**Checklist de Verificação:**

### 6.1 Varredura de Limpeza
Executar busca e remoção nos arquivos alterados (T-01 a T-05) de:
- Trechos de código comentados (blocos `// TODO` temporários, código morto em `/* */`)
- Arquivos temporários de testes (`*_test.go.bak`, `*.tmp`, `testdata/` não versionado)
- Logs de debug (`fmt.Println`, `log.Print`, `slog.Debug` não estruturados, `console.log` equivalente)
- Comentários de rascunho como `// WIP`, `// TESTING`, `// DEBUG`

**Comando de referência:**
```bash
# Buscar logs de debug e código comentado suspeito nos arquivos do diff
git diff --name-only HEAD~1 | xargs grep -rn "fmt.Println\|log.Print\|// WIP\|// DEBUG\|// TEST" || echo "✅ Nenhum"
# Listar arquivos temporários/não rastreados
git status --porcelain | grep "^?" | grep -E "\.tmp|\.bak|\.old" || echo "✅ Nenhum"
```

### 6.2 Auditoria de Git Status & Arquitetura
1. Executar `git status` e listar os arquivos modificados
2. Validar se a localização e nomenclatura de cada arquivo cumpre estritamente o [ARCHITECTURE.md](./ARCHITECTURE.md):
   - `internal/reforma/` — calculadoras CBS/IBS
   - `internal/calculator/` — pipeline de fases
   - `internal/legacy/` — mock repository
   - `data/init.sql` — DDL e INSERTs
   - `libs/go-native/taxnexus-billing-core-lib/` — repository e models
3. Verificar se nenhum arquivo foi criado fora dos diretórios previstos na arquitetura

**Comando de referência:**
```bash
git status --porcelain
```

### 6.3 Travas do SECURITY.md
Realizar varredura no código gerado contra o [SECURITY.md](../../../../../.specs/security/SECURITY.md) para garantir:
- Nenhuma vulnerabilidade introduzida (OWASP Top 10)
- Nenhuma chave, token ou senha hardcoded (`grep -r "password\|secret\|token\|api_key\|private_key" --include="*.go"`)
- Nenhum dado sensível exposto em logs ou respostas de erro
- Validação de input em todas as funções públicas que recebem dados externos
- Princípio do menor privilégio respeitado nas novas queries

**Comando de referência:**
```bash
# Buscar segredos hardcoded
grep -rE "password|secret|token|api_key|private_key|access_key" --include="*.go" internal/ libs/ cmd/ data/ || echo "✅ Nenhum"
# Buscar vazamento de dados em logs
grep -rE "log\.(Print|Info|Warn|Error).*%v.*senha|log\.(Print|Info|Warn|Error).*%v.*cpf|log\.(Print|Info|Warn|Error).*%v.*token" --include="*.go" || echo "✅ Nenhum"
```

### 6.4 Travas do TEST_PLAN.md
Realizar varredura no código gerado contra o [TEST_PLAN.md](./TEST_PLAN.md) para garantir:
- Todos os cenários de teste (TST-03.01 a TST-03.27) estão implementados
- Testes de borda e erro (TST-03.24 a TST-03.27) cobertos
- Verificações documentais (TST-DOC.01 a TST-DOC.08) atendidas
- Baseline de 211+ testes mantida, meta de 238+ após feature atingida
- Nenhum critério de aceitação definido no TEST_PLAN.md foi ignorado

**Comando de referência:**
```bash
go test ./... -v 2>&1 | tail -20
```

### 6.5 Fechamento de Evidências
Confirmar se todas as tarefas funcionais e de testes desta feature estão com marcadores de conclusão:
- T-01: `[x]` ou `[✅]`
- T-02: `[x]` ou `[✅]`
- T-03: `[x]` ou `[✅]`
- T-04: `[x]` ou `[✅]`
- T-05: `[x]` ou `[✅]`

### 6.6 Procedimento em Caso de Falha

Se houver **qualquer** falha, vulnerabilidade, ou teste quebrado:

1. Gerar um relatório estruturado em Markdown no arquivo `.specs/features/FEATURE-2026-06-30-DT03-CST-REFORMA-TRIBUTARIA/FEEDBACK_ERRORS.md` seguindo o modelo abaixo
2. O relatório pode ser enriquecido com informações adicionais que ajudem o desenvolvedor a revisar a implementação
3. Marcar o relatório com a tag `[STATUS: FAILED]`
4. Voltar para a tarefa **T-01**, reportando que o relatório gerado deve ser usado como base para corrigir e re-executar a implementação

**Modelo de Saída (`FEEDBACK_ERRORS.md`):**

```markdown
## 🚨 Relatório de Falhas Encontradas (Loop {{LOOP_ATUAL}}/3)

### 📈 Resumo do Status
- **Total de Problemas:** [Quantidade]
- **Severidade Máxima:** [Crítico / Médio / Baixo]

### 🔍 Detalhamento dos Erros

#### Erro 1: [Nome Curto do Erro]
- **Severidade:** [Crítico | Médio | Baixo]
- **Arquivo e Linha:** `caminho/do/arquivo.ext` - Linha XX
- **Diretriz Violada:** [Citar a regra do SECURITY.md ou o caso de teste do TEST_PLAN.md]
- **O que está acontecendo:** [Explicação técnica e direta do comportamento incorreto do código atual]
- **Log de Erro / Evidência:**
```text
[Cole aqui o erro do compilador, log do teste ou trecho vulnerável]
```
- **Critério de Correção:** [Instrução explícita do que o desenvolvedor deve alterar para corrigir o problema]

**[STATUS: FAILED]**
```

### 6.7 Critério de Aprovação

Se **todas** as verificações (6.1 a 6.5) passarem sem falhas:
- Nenhum `FEEDBACK_ERRORS.md` gerado
- Sanity Check marcado como `[STATUS: ✅ APROVADO]`
- Prosseguir para T-07 (Documentação)

**Arquivos verificados:**
- Todos os arquivos alterados no diff (T-01 a T-05)
- [ARCHITECTURE.md](./ARCHITECTURE.md) (referência de estrutura)
- [SECURITY.md](../../../../../.specs/security/SECURITY.md) (checklist de segurança)
- [TEST_PLAN.md](./TEST_PLAN.md) (cenários de teste)

**DoD:** Nenhum `FEEDBACK_ERRORS.md` gerado OU relatório gerado e encaminhado para revisão; `git status` validado contra ARCHITECTURE.md; varredura de segurança limpa

---

### T-07: ✅ Atualizar Documentação (Entregável da Feature)

**Prioridade:** 🔴 Must
**Estimativa:** 1 hora
**Dependências:** T-06

**Descrição:**
1. Marcar DT-03 como ✅ Resolvida no `feature-roadmap.md`
2. Atualizar `domain/domain.md` — adicionar referência à tabela `cst_reforma`
3. Atualizar `erd.md` — adicionar tabela `cst_reforma` ao diagrama
4. Atualizar `tax-rates-api.yaml` — documentar que CST agora é 3 dígitos
5. Adicionar entrada no `CHANGELOG.md`
6. Atualizar este TASKS.md — marcar todas as tarefas como ✅ Concluídas
7. Atualizar SPECS.md — status ✅ Concluído
8. Atualizar TEST_PLAN.md — status ✅ Executado
9. Atualizar ADRs (010, 011, 012) — status 🔨 Proposto → ✅ Aceito
10. Atualizar `adrs/INDEX.md` — status dos ADRs

**Arquivos:**
- `.specs/product/feature-roadmap.md`
- `.specs/domain/domain.md`
- `.specs/architecture/erd.md`
- `.specs/api/tax-rates-api.yaml`
- `.specs/CHANGELOG.md`
- `.specs/architecture/adrs/adr-010.md`
- `.specs/architecture/adrs/adr-011.md`
- `.specs/architecture/adrs/adr-012.md`
- `.specs/architecture/adrs/INDEX.md`
- `.specs/features/FEATURE-2026-06-30-DT03-CST-REFORMA-TRIBUTARIA/SPECS.md`
- `.specs/features/FEATURE-2026-06-30-DT03-CST-REFORMA-TRIBUTARIA/TASKS.md`
- `.specs/features/FEATURE-2026-06-30-DT03-CST-REFORMA-TRIBUTARIA/TEST_PLAN.md`

**DoD:** DT-03 ✅ Resolvida, feature docs ✅ Concluído, ADRs ✅ Aceito, docs microserviço atualizados

---

### T-08: ✅ Documentação da Solução (Relatório de Execução)

**Prioridade:** 🔴 Must
**Estimativa:** 1-2 horas
**Dependências:** T-07

**Descrição:**
Gerar a documentação final da solução (relatório de implementação) usando o prompt padronizado do projeto como base. Esta tarefa produz o artefato de auditoria que comprova a execução completa da feature.

**Procedimento:**
1. Utilizar o prompt definido em `../../../../../.specs/prompts/PROMPT-GENERATE-IMPLEMENTATION-REPORT.md` como base para estruturar a documentação
2. O prompt instrui o uso de skills auxiliares (`code-documenter`, `gap-analysis`, `spec-miner`, `code-reviewer`, `fullstack-guardian`) para enriquecer o relatório
3. Gerar o arquivo de saída na pasta `.specs/skill-output/`
4. O nome do arquivo deve seguir estritamente o padrão: `{AAAA-MM-DD-HHMMSS}-TASK-EXECUTED-[nome-da-feature-em-kebab-case].md`

**Conteúdo mínimo do relatório:**
- Resumo do desenvolvimento realizado (parágrafo técnico com classes, funções, módulos criados/alterados)
- Lista de arquivos modificados ou criados (tabela com Ação, Arquivo, Mudança)
- Evidências e resultados dos testes (`go test ./... -v`, total de testes, status final)
- Validação de segurança e qualidade (checklist do SECURITY.md)
- Atualização de documentação realizada
- Dívidas técnicas resolvidas e novas dívidas técnicas identificadas

**Exemplos de saída na pasta `skill-output/`:**
- `2026-07-01-142800-dt03-cst-reforma-tributaria.md` (já existente, referência de formato)
- `2026-06-25-073636-reforma-tributaria-fases-0-1-2.md`
- `2026-06-22-011000-c001-pipeline-sop013.md`

**Arquivos referência:**
- `../../../../../.specs/prompts/PROMPT-GENERATE-IMPLEMENTATION-REPORT.md` — prompt base
- `.specs/skill-output/` — diretório de saída

**DoD:** Arquivo `{timestamp}-TASK-EXECUTED-dt03-cst-reforma-tributaria.md` gerado em `.specs/skill-output/` com todas as seções preenchidas

---

### T-09: ✅ Auditoria de Compliance Documental (Quality Gate)

**Prioridade:** 🔴 Must
**Estimativa:** 30-45 min
**Dependências:** T-08
**Ciclo:** Se divergências → voltar a T-07 para correção; se OK → ✅

**Descrição:**
Validar que todos os documentos da feature e da solução foram atualizados corretamente por T-07 e T-08. Esta tarefa é um **gate de qualidade** — a feature não é considerada concluída até que todas as verificações passem.

**Checklist de Verificação:**

| # | Verificação | Documento | Método |
|---|-------------|-----------|--------|
| V-01 | Status ✅ Concluído | `SPECS.md` | Ler header |
| V-02 | Status ✅ Executado | `TEST_PLAN.md` | Ler header |
| V-03 | Todas as tarefas ✅ | `TASKS.md` (este arquivo) | Verificar tabela resumo |
| V-04 | Status ✅ Aceito | `adrs/adr-010.md` | Ler header |
| V-05 | Status ✅ Aceito | `adrs/adr-011.md` | Ler header |
| V-06 | Status ✅ Aceito | `adrs/adr-012.md` | Ler header |
| V-07 | Status atualizado | `adrs/INDEX.md` | grep "010.*Aceito\|011.*Aceito\|012.*Aceito" |
| V-08 | DT-03 ✅ Resolvida | `product/feature-roadmap.md` | grep linha DT-03 |
| V-09 | Entrada da feature | `CHANGELOG.md` | grep "DT-03\|cst_reforma" |
| V-10 | Tabela `cst_reforma` referenciada | `architecture/erd.md` | grep "cst_reforma" |
| V-11 | CST 3 dígitos documentado | `api/tax-rates-api.yaml` | grep "ISS.*FUST.*FUNTTEL" (enum completo) |
| V-12 | Sem referências stale (`cstPadrao`/`cstIsento`) | Todo `.specs/` | `grep -r "cstPadrao\|cstIsento" .specs/` |
| V-13 | Links cross-ref funcionam | Feature docs → solução | Verificação manual de 5 links |
| V-14 | Datas consistentes com commit | Todos os docs | Verificar data/hora |

**Procedimento em caso de divergência:**

1. Registrar a divergência encontrada (ex: "V-10 falhou: `erd.md` não referencia `cst_reforma`")
2. Listar as divergências em um comentário na tabela de resumo abaixo
3. Atualizar T-07 com os arquivos pendentes
4. Re-executar T-07 para corrigir os docs apontados
5. Re-executar T-08 (Documentação da Solução) para refletir as correções
6. Re-executar T-09

**Arquivos verificados:**
- `.specs/features/FEATURE-2026-06-30-DT03-CST-REFORMA-TRIBUTARIA/SPECS.md`
- `.specs/features/FEATURE-2026-06-30-DT03-CST-REFORMA-TRIBUTARIA/TASKS.md`
- `.specs/features/FEATURE-2026-06-30-DT03-CST-REFORMA-TRIBUTARIA/TEST_PLAN.md`
- `.specs/architecture/adrs/adr-010.md`
- `.specs/architecture/adrs/adr-011.md`
- `.specs/architecture/adrs/adr-012.md`
- `.specs/architecture/adrs/INDEX.md`
- `.specs/product/feature-roadmap.md`
- `.specs/CHANGELOG.md`
- `.specs/architecture/erd.md`
- `.specs/api/tax-rates-api.yaml`
- `.specs/domain/domain.md`

**DoD:** 14/14 verificações passam OU divergências documentadas e T-07 re-executada

---

## 4. Resumo

| ID | Tarefa | Prioridade | Estimativa | Depende de |
|----|--------|-----------|------------|------------|
| T-01 | DDL + INSERTs `cst_reforma` | 🔴 Must | 2-3h | — |
| T-02 | `GetCSTReforma()` repository | 🔴 Must | 2-3h | T-01 |
| T-03 | Refactor `reforma.go` | 🔴 Must | 3-4h | T-02 |
| T-04 | Cache Redis | 🟡 Should | 1-2h | T-02 |
| T-05 | Atualizar testes | 🔴 Must | 2-3h | T-03 |
| T-06 | Protocolo de Checagem Pós-Implementação (Sanity Check) | 🔴 Must | 1-2h | T-05 |
| T-07 | Atualizar documentação (entregável) | 🔴 Must | 1h | T-06 |
| T-08 | Documentação da Solução (Relatório de Execução) | 🔴 Must | 1-2h | T-07 |
| T-09 | Auditoria compliance documental (quality gate) | 🔴 Must | 30-45min | T-08 |
| **Total** | | | **15-23h** | |

---

## 5. Risco e Mitigações

| Risco | Impacto | Mitigação |
|-------|---------|-----------|
| Tabela oficial ainda sofrer alterações (LC 214/2025 pendente de regulamentação) | Retrabalho | Tabela updatável via Admin Fiscal sem deploy |
| Consumidores quebrarem com CST 3 dígitos | Breaking change | Notificar consumidores; campo é informativo, não quebra lógica |
| Performance: +1 query por cálculo CBS/IBS | Latência adicional | Cache Redis (TTL 24h); segunda chamada é sub-ms |
| Complexidade da lógica de seleção de CST | CST incorreto | Fallback `"000"`; validação com time fiscal |
| Loop-back da Sanity Check (T-06 → T-01) causar retrabalho excessivo | Atraso na entrega | Relatório `FEEDBACK_ERRORS.md` permite revisão humana antes do loop; limitado a 3 ciclos |
