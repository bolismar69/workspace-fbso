# PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-DEVOPS-SRE-DEFINITION (F5)

## Contexto

Acionado quando o gate reprova `DETAIL-LEVEL-DEVOPS-SRE-DEFINITION.md`.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |

## Processamento de NCs

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Pipeline de produção ausente | Adicionar workflow de deploy prod |
| P0 | SLO sem SLI | Definir métrica mensurável e janela |
| P1 | IaC faltando recurso | Adicionar template para o recurso |
| P1 | Runbook ausente | Documentar procedimento de recuperação |
| P2 | Observabilidade incompleta | Completar stack (métricas, logs, traces) |
| P2 | Estratégia de deploy vaga | Detalhar blue-green/canary |

## Skills Recomendados
- `gap-analysis`, `senior-devops`, `sre-engineer`
- `observability-engineer`, `monitoring-expert`

🤖 *Fix — Fase 5 do Downstream Architecture Refinement*
