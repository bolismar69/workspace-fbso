# PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-UPSTREAM-COMPARISON (F12)

## Contexto

Acionado quando o gate reprova `UPSTREAM-COMPARISON-REPORT.md`. Fase condicional.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |

## Condicionalidade

⚠️ Se o gate retornou `[COMPLIANCE — FASE NÃO APLICÁVEL]` (upstream não existe), este FIX não executa.

## Processamento de NCs

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | PERT foi alterado no relatório | **Reverter imediatamente** — restaurar valores da F8 congelada |
| P1 | Tabela comparativa incompleta | Completar ROM vs PERT por épico |
| P1 | Desvio sem justificativa | Documentar causa (escopo, complexidade, metodologia) |
| P2 | Gráfico comparativo ausente | Adicionar visualização proporcional |
| P2 | Conclusão ambígua | Esclarecer: refina, substitui ou diverge? |

## Skills Recomendados
- `gap-analysis`, `documentation-writer`

🤖 *Fix — Fase 12 do Downstream Architecture Refinement · Cross-Check Condicional*
