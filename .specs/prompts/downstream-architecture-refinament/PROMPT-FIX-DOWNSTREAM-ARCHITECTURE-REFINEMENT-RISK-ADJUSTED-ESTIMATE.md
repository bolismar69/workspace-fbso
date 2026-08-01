# PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-RISK-ADJUSTED-ESTIMATE (F10)

## Contexto

Acionado quando o gate reprova `RISK-ADJUSTED-ESTIMATE.md`.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |

## Processamento de NCs

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Ajuste sobre ROM em vez de PERT | Recalcular usando baseline PERT (F8) |
| P0 | Risco crítico sem mitigação | Adicionar ação de mitigação |
| P1 | Probabilidade/Impacto não quantificados | Atribuir valores numéricos |
| P1 | Cenário (Conservador/PERT/Pessimista) ausente | Completar os 3 cenários |
| P2 | Análise de sensibilidade incompleta | Identificar top 3 riscos por impacto |
| P2 | Recomendação genérica | Tornar acionável e específica |

## Skills Recomendados
- `gap-analysis`, `project-estimation`
- `risk-manager` (se disponível)

🤖 *Fix — Fase 10 do Downstream Architecture Refinement*
