# PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-DATA-ARCHITECTURE-DEFINITION (F4)

## Contexto

Acionado quando o gate reprova `DETAIL-LEVEL-DATA-ARCHITECTURE-DEFINITION.md`.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |

## Processamento de NCs

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Entidade ausente no modelo | Adicionar tabela com atributos, constraints, índices |
| P0 | Estratégia multi-tenancy não definida | Documentar mecanismo de isolamento |
| P1 | Query crítica sem índice | Adicionar índice e justificativa |
| P1 | Particionamento não definido | Definir estratégia para tabelas de alto volume |
| P2 | Migration strategy incompleta | Documentar versionamento, baseline, rollback |
| P2 | Projeção de volumes ausente | Adicionar estimativas de crescimento |

## Skills Recomendados
- `gap-analysis`, `senior-data-engineer`, `database-architect`
- `database-design`, `data-modeling`
- `database-migrations`, `database-migrations-sql-migrations`

🤖 *Fix — Fase 4 do Downstream Architecture Refinement*
