# PROMPT: FIX CIRÚRGICO — SCOPE SNAPSHOT UPSTREAM/DISCOVERY
## Versão: 1.0 — WATERFALL Estimation Orchestrator

Atue como Especialista em Correção de Documentos de Escopo.

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
| SSU-01 | Componente estimado não listado no snapshot | P0 | Adicionar componente à seção 1 com referência ao documento fonte |
| SSU-02 | Item excluído sem motivo documentado | P1 | Adicionar motivo e fonte da decisão na seção 2 |
| SSU-03 | Rastreabilidade quebrada (célula vazia) | P1 | Preencher referência § em todos os documentos fonte aplicáveis |
| SSU-04 | Versão de documento fonte ausente | P1 | Adicionar versão e data do documento |
| SSU-05 | Premissa sem coluna de impacto | P2 | Adicionar "Impacto se Inválida" |
| SSU-06 | Metadados incompletos | P2 | Preencher campos do cabeçalho |
| SSU-07 | Seção ausente | P2 | Adicionar seção faltante |
