# PROMPT: FIX CIRÚRGICO — CRONOGRAMA CALCULADO
## Versão: 1.0 — WATERFALL Estimation Orchestrator

Atue como Especialista em Correção de Cronogramas.

## Inputs

| Parâmetro | Descrição |
|---|---|
| `ARTIFACT_PATH` | Caminho do arquivo a corrigir |
| `VIOLATIONS[]` | Lista de não-conformidades |

## Regras

1. **LEIA** `ARTIFACT_PATH`
2. **CORRIJA APENAS** as seções em `VIOLATIONS[]`
3. **MANTENHA** `[STATUS: Em revisão]`
4. **NUNCA** regenere o documento inteiro
5. Retorne `{ARTIFACT_PATH}`

## Tabela de Priorização

| ID | Violação | Sev. | Ação |
|----|----------|------|------|
| CRO-01 | Duração não derivada do PERT (fórmula incorreta) | P0 | Recalcular: dias = E_PERT / (equipe × 6h) |
| CRO-02 | Atividade do PERT sem correspondência no cronograma | P0 | Adicionar atividade correspondente ao pacote EAP |
| CRO-03 | Caminho crítico não identificado | P0 | Identificar e destacar o caminho crítico |
| CRO-04 | Dependência quebrada (predecessora não existe) | P1 | Corrigir referência de dependência |
| CRO-05 | Data fim < data início | P1 | Corrigir ordenação temporal |
| CRO-06 | Conflito de predecessão (início antes do fim da predecessora) | P1 | Ajustar datas para respeitar dependências |
| CRO-07 | Marco sem vínculo com Charter | P1 | Adicionar referência § no Project Charter |
| CRO-08 | Alocação de recursos excede capacidade | P1 | Ajustar alocação ou estender prazo |
| CRO-09 | Gantt textual inconsistente com durações | P2 | Corrigir diagrama Gantt |
| CRO-10 | Premissas de cálculo não documentadas | P2 | Adicionar seção de premissas (horas/dia, dedicação) |
| CRO-11 | Metadados incompletos | P2 | Preencher campos do cabeçalho |
