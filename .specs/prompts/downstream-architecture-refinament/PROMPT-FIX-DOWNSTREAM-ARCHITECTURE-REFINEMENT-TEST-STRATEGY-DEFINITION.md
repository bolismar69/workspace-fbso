# PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-TEST-STRATEGY-DEFINITION (F6)

## Contexto

Acionado quando o gate reprova `DETAIL-LEVEL-TEST-STRATEGY-DEFINITION.md`.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |

## Processamento de NCs

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | US sem cobertura de teste | Adicionar tipo(s) de teste para a US |
| P0 | Quality gate sem critério | Definir critério de aprovação |
| P1 | Camada de teste sem ferramenta | Definir framework/ferramenta |
| P1 | Teste multi-tenant ausente | Adicionar cenários de isolamento |
| P2 | Pirâmide desbalanceada | Ajustar proporção Unit>Integ>E2E |
| P2 | Caso de aceitação faltando | Criar caso baseado no cenário da US |

## Skills Recomendados
- `gap-analysis`, `senior-qa`, `testing-patterns`
- `test-master`, `testing-qa`, `test-case-creation`
- `e2e-testing`

🤖 *Fix — Fase 6 do Downstream Architecture Refinement*
