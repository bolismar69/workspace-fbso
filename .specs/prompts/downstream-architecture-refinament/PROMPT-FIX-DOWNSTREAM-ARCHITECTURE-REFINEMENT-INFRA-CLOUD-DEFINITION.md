# PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-INFRA-CLOUD-DEFINITION (F7)

## Contexto

Acionado quando o gate reprova `DETAIL-LEVEL-INFRA-CLOUD-DEFINITION.md`.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |

## Processamento de NCs

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Recurso sem sizing | Dimensionar com justificativa |
| P0 | RPO/RTO não definidos | Definir objetivos de recuperação |
| P1 | Custo não calculado | Detalhar por provedor e serviço |
| P1 | Backup strategy ausente | Documentar ferramenta, frequência, retenção |
| P2 | Ambiente não definido | Adicionar dev/staging/prod |
| P2 | Topologia de rede sem diagrama | Adicionar fluxo de acesso |

## Skills Recomendados
- `gap-analysis`, `cloud-architect`, `senior-devops`
- `disaster-recovery` (se disponível), `hybrid-cloud-networking`

🤖 *Fix — Fase 7 do Downstream Architecture Refinement*
