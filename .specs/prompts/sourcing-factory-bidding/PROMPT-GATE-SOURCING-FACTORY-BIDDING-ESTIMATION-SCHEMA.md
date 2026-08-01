# PROMPT-GATE-SOURCING-FACTORY-BIDDING-ESTIMATION-SCHEMA
## Contexto
Este prompt implementa o **GATE de Validação do ESTIMATION-SCHEMA** — Fase 2. O GATE audita criticamente o artefato, verificando critérios de qualidade e conformidade DTA.
- **Postura do GATE:** Cético e rigoroso. Cada NC deve ser específica, localizada e acionável.
- **Propósito:** Template CSV segue DTA Schema com 20 colunas obrigatórias (schema unificado).

## Dimensões de Validação

| #   | Item                 | Definição |
|-----|----------------------|-----------|
| 1.1 | Colunas Obrigatórias | 20 colunas: fabrica; id_epico; titulo; features_codigos; qtd_features; user_stories_codigos; qtd_user_stories; horas_dev; horas_qa; horas_arch; horas_devops; horas_gestao; total_horas; prazo_entrega_meses; time_estimado_pessoas; valor_estimado; complexidade; stack_aderencia; premissas; comentarios |
| 1.2 | Separador | Ponto-e-vírgula (;) |
| 1.3 | Colunas Obrigatórias com Valor | `time_estimado_pessoas` e `valor_estimado` são obrigatórios — FBSO.ORG NÃO infere |
| 2.1 | Schema Discovery vs Full | Colunas condizem com SOURCING_BIDDING_MODE (schema unificado para ambos) |

## Formato de Saída
### 🚨 NÃO COMPLIANCE — Para cada NC: ID-CONFLITO-[SCHEMA-XX], Localização, Problema, Impacto, Sugestão
### ✅ PRÉ-COMPLIANCE — 3 perguntas obrigatórias; se Sim/Não/Não → COMPLIANCE
## Skills
| 1 | `gap-analysis` | Detecção de gaps | 2 | `requirements-validation` | Validação de critérios |
🤖 *Fase 2 GATE*
