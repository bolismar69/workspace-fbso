# PROMPT-GATE-SOURCING-FACTORY-BIDDING-ESTIMATION-SCHEMA
## Contexto
Este prompt implementa o **GATE de Validação do ESTIMATION-SCHEMA** — Fase 2. O GATE audita criticamente o artefato, verificando critérios de qualidade e conformidade DTA.
**Postura do GATE:** Cético e rigoroso. Cada NC deve ser específica, localizada e acionável.
**Propósito:** Template CSV segue DTA Schema com colunas obrigatórias.
## Dimensões de Validação
| 1.1 | Colunas Obrigatórias | id_epico, titulo, solucoes, horas_dev, horas_arch, horas_qa, prazo_entrega_meses, complexidade, comentarios |
| 1.2 | Separador | Ponto-e-vírgula (;) |
| 2.1 | Schema Discovery vs Full | Colunas condizem com SOURCING_BIDDING_MODE |
## Formato de Saída
### 🚨 NÃO COMPLIANCE — Para cada NC: ID-CONFLITO, Localização, Problema, Impacto, Sugestão
### ✅ PRÉ-COMPLIANCE — 3 perguntas obrigatórias; se Sim/Não/Não → COMPLIANCE
## Skills
| 1 | `gap-analysis` | Detecção de gaps | 2 | `requirements-validation` | Validação de critérios |
🤖 *Fase 2 GATE*
