# PROMPT-FIX-SOURCING-FACTORY-BIDDING-ESTIMATE-RETROSPECTIVE-PIB (F5b)

## Contexto

Acionado quando o gate reprova `ESTIMATE-RETROSPECTIVE-PIB.md`. Correções cirúrgicas apenas. Fase condicional.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |
| `{SOURCING_BIDDING_MODE}` | Modo: `agile-discovery`, `agile-refinement`, `waterfall-discovery` ou `waterfall-refinement` |

## Condicionalidade

⚠️ Se a F5 aprovou ≥ 1 fábrica, este FIX não executa (fase não aplicável).

## Processamento de NCs

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | PIB por épico não calculado | Calcular para cada épico com baseline correta do modo |
| P1 | Flat estimate não detectado | Calcular CV entre épicos; documentar fábricas com CV < 10% |
| P1 | Independência não verificada | Comparar valores cross-fábrica; documentar duplicatas |
| P2 | Comentários não analisados | Inspecionar coluna `comentarios`; classificar qualidade |
| P2 | Recomendações genéricas | Tornar acionáveis e específicas por fábrica |
| P3 | Fábricas afetadas não listadas | Para cada problema, listar quais fábricas são afetadas |

## Skills Recomendados
- `gap-analysis`, `analyst-estimates`

🤖 *Fix — Fase 5b do Sourcing & Factory Bidding. Condicional: 0 aprovadas.*
