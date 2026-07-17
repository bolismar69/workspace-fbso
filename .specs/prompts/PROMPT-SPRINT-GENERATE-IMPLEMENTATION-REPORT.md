# Contexto:
- Obter o historico de implementações, assim como o historico de logs e evidencias de testes, usando alguns skills para ajudar nessa atividade (skils: code-documenter, gap-analysis, spec-miner, code-reviewer, fullstack-guardian, caveman)

# Missão:
- Gerar um arquivo de documentação de saída na pasta correspondente da funcionalidade.
- O nome do arquivo deve seguir estritamente o padrão `{AAAA-MM-DD-HHMMSS}-TASK-EXECUTED-[nome-da-feature-em-kebab-case].md` (Substitua AAAA-MM-DD-HHMMSS pelo carimbo de data e hora exato do momento da criação).
- Gerar o arquivo na pasta `{humano-ira-informar-a-pasta}`
- O modelo para geração desse arquivo esta abaixo:

````markdown

# 📑 Relatório de Execução de Tarefa (TASK-EXECUTED)

* **Data e Hora da Conclusão:** [Preencher com AAAA-MM-DD HH:MM:SS]
**Skills utilizados:** golang-pro + 
* **Projeto/Incidente/Feature/Hot-Fix/POC/Incident/Other:** [ex:PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO]
* **Fase/Escopo/Feature/Issue/Sprint Concluído:**: [ex: Fundação]

---

## 🛠️ 1. Resumo do Desenvolvimento Realizado
[Descreva em um parágrafo técnico denso o que foi implementado, citando as principais classes, funções ou módulos criados/alterados na linguagem de programação (ex: Go)].

## 🗂️ 2. Arquivos Modificados ou Criados
Liste rigorosamente os caminhos dos arquivos afetados para auditoria do Git:
|Ação| Arquivo | Mudança |
|----|---------|---------|
| 🆕 | `internal/calculator/engine.go` | Refatorado — `CalculationPhase` + `ExecutionMode` + `BillingEnginePhased()` + `injectTributoValues()` |
| 🔄 | `internal/calculator/pipeline_test.go` | **Novo** — 22 cenários de teste |
| 🆕 | `internal/reforma/reforma.go` | Refatorado — `computeIvaDual()` compartilhado, `ReformaCalculator` legado |
| 🔄 | `internal/reforma/cbs_calculator.go` | **Novo** — `CBSCalculator` (Fase 2 sequencial) |
| 🆕 | `internal/reforma/ibs_calculator.go` | **Novo** — `IBSCalculator` (Fase 4 paralela) |
| 🆕 | `cmd/api/main.go` | Rewiring — 7 fases `BillingEnginePhased()`, CBS/IBS split, LegacyAdapter sem icmsSource |

## 🧪 3. Evidências e Resultados dos Testes (`TEST_PLAN.md`)
O agente de IA executou a suíte de testes com sucesso.
* **Comando Executado:** `[Insira o comando de terminal utilizado, ex: go test ./... -v]`
* **Total de Testes Rodados:** [Número de testes]
* **Status Final:** 🟩 100% PASSOU
* **Saída Sumarizada do Terminal:**
```text
[Cole aqui o trecho final do log do terminal que comprova que os testes passaram]
```

## 🔒 4. Validação de Segurança e Qualidade (`SECURITY.md`)
* [✅] Nenhuma credencial ou dado sensível foi deixada em formato hardcoded.
* [✅] Todos os novos inputs foram sanitizados via schemas/validadores.
* [✅] O código passou na verificação estática do Linter do projeto.

## Documentação Atualizada

- `FEATURE-2026-06-21.md` — C-001 marcado ✅
- `feature-roadmap.md` — C-001 concluído, DT-12 resolvido
- `architecture/architecture.md` — Seções 3, 7, 11, 16 reescritas

## Dividas Técnicas Resolvidas

- DT-12: Pipeline de cálculo completamente reordenado (SOP-013)

## Dividas Técnicas Que Surgiram

- C-001: Todas as dependências assimétricas implementadas (IPI→ICMS, ICMS→PIS/COFINS, ICMS+PIS+COFINS→FUST/FUNTTEL, IS→CBS)

---
🤖 *Documentação gerada de forma automatizada pelo agente de desenvolvimento de IA.*

````
