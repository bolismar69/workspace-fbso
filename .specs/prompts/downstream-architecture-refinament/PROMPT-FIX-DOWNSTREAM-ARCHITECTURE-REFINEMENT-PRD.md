# PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-PRD (F1)

## Contexto

Acionado quando o gate reprova `DETAIL-LEVEL-PRD.md`. O agente corretor aplica **correções cirúrgicas** com base no relatório inline do gate. **Nunca reescreve o documento do zero.**

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos |
| `{PROJECT_ID_NAME}` | ID completo do projeto |

## Fluxo de Execução

### Passo 1 — Carregar Relatório do Gate
Ler o relatório de auditoria, documento atual e respostas do humano.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Referência ao upstream-architecture-discovery | Remover referência e buscar fonte nos docs de negócio |
| P1 | Entrega (D) sem feature mapeada | Completar matriz de escopo |
| P1 | Persona sem jornada | Adicionar cenários de uso |
| P2 | Glossário incompleto | Adicionar termos de domínio |
| P2 | Doc de projeto não referenciado | Adicionar referência com path |
| P3 | Doc de apoio existente não referenciado | Adicionar referência ao doc (DEFINITION_OF_DONE, GLOSSARY, MATRIZ-KPI, STAKEHOLDER-MAP) |
| P3 | Doc de apoio não existe | Prosseguir — não é obrigatório |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

## Skills Recomendados
- `gap-analysis`, `business-analyst`, `agile-ba-practices`
- `documentation-writer`

🤖 *Fix — Fase 1 do Downstream Architecture Refinement*
