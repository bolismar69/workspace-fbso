# PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-BOTTOM-UP-PERT-ESTIMATE (F8)

## Contexto

Acionado quando o gate reprova `BOTTOM-UP-PERT-ESTIMATE.md`. **Este é o prompt FIX mais crítico do roadmap.**

⚠️ **Regra especial:** Se o conflito for **ID-EST-01 (Contaminação pelo ROM upstream)**, o FIX cirúrgico NÃO é permitido. A estimativa deve ser **refeita do zero** pelo GENERATE (Fase 8). Este prompt FIX aplica-se apenas aos demais tipos de NC.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |

## Processamento de NCs

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| **P0** | **Contaminação pelo ROM** | **🚫 NÃO CORRIGIR. Retornar ao GENERATE F8 para refazer do zero.** |
| P0 | US sem estimativa | Adicionar O, ML, P, PERT, σ para a US faltante |
| P0 | QA < 25% | Revisar % de QA por épico ou adicionar buffer |
| P1 | Arch < 5% | Revisar % de arquitetura global |
| P1 | IC 95% não calculado | Calcular para todos os níveis |
| P2 | Rollup inconsistente | Corrigir somas Feature→Épico→Projeto |
| P2 | Outlier sem justificativa | Documentar razão do desvio |

## Skills Recomendados
- `gap-analysis`, `project-estimation`
- Referências: `bottom-up-estimation.md`, `three-point-estimation-pert.md`

🤖 *Fix — Fase 8 do Downstream Architecture Refinement · Regra especial para contaminação*
