# PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-SCOPE-SNAPSHOT (F11)

## Contexto

Acionado quando o gate reprova `SCOPE-SNAPSHOT.md`.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |

## Processamento de NCs

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | US ausente do snapshot | Adicionar US com ID, descrição, feature, épico, PERT |
| P0 | Snapshot contém planejamento de sprint | Remover sprints, tarefas ou contratos |
| P1 | Data de congelamento ausente | Registrar data explícita |
| P1 | Hash do escopo não calculado | Gerar checksum das US incluídas |
| P2 | Referência à F8 ausente | Adicionar link para BOTTOM-UP-PERT-ESTIMATE |
| P2 | Nota de imutabilidade faltando | Adicionar aviso explícito |

## Skills Recomendados
- `gap-analysis`, `configuration-manager` (se disponível)
- `documentation-writer`

🤖 *Fix — Fase 11 do Downstream Architecture Refinement*
