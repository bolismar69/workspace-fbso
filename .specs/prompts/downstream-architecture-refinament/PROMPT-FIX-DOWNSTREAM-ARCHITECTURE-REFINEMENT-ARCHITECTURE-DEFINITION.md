# PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-ARCHITECTURE-DEFINITION (F2)

## Contexto

Acionado quando o gate reprova `DETAIL-LEVEL-ARCHITECTURE-DEFINITION.md`. Correções cirúrgicas apenas nos pontos de não-compliance.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho das soluções técnicas |

## Processamento de NCs

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Serviço ausente no C4 L2 | Adicionar container ao diagrama |
| P0 | ADR sem diagrama de sequência | Adicionar diagrama com fluxo |
| P1 | Componente C4 L3 incompleto | Completar decomposição interna |
| P1 | Integração sem protocolo | Documentar protocolo e autenticação |
| P2 | Padrão de código não documentado | Adicionar convenção |
| P2 | Estratégia multi-tenancy vaga | Detalhar mecanismo de isolamento |

## Skills Recomendados
- `gap-analysis`, `senior-architect`, `engineering-skills`
- `c4-container`, `c4-component`, `architecture-decision-records`

🤖 *Fix — Fase 2 do Downstream Architecture Refinement*
