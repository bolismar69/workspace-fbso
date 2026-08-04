# PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-RESOURCE-ALLOCATION (F9)

## Contexto

Acionado quando o gate reprova `RESOURCE-ALLOCATION-PLAN.md`.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |

## Processamento de NCs

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Horas não derivadas do PERT | Substituir por valores da F8 |
| P0 | Fonte externa detectada | Remover referência a ROM/factory bids |
| P1 | Capacidade efetiva irrealista | Recalcular com cargas parciais corretas |
| P1 | Gargalo não identificado | Adicionar gargalo com recomendação |
| P2 | Papel necessário ausente | Adicionar papel/perfil necessário faltante |
| P2 | Duração inconsistente | Recalcular: horas ÷ capacidade |

## Skills Recomendados
- `gap-analysis`, `project-estimation`
- `senior-architect`

🤖 *Fix — Fase 9 do Downstream Architecture Refinement*
