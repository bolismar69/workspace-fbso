# PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-SECURITY-DEFINITION (F3)

## Contexto

Acionado quando o gate reprova `DETAIL-LEVEL-SECURITY-DEFINITION.md`.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |
| `{SECURITY_GLOBAL}` | GLOBAL-SECURITY.md |

## Processamento de NCs

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Componente sem análise STRIDE | Adicionar threat analysis para o componente |
| P0 | Ameaça sem mitigação | Adicionar controle correspondente |
| P1 | OWASP ASVS incompleto | Completar matriz de controles |
| P1 | IAM spec incompleta | Detalhar realms, clients, claims, fluxos |
| P2 | Matriz RBAC com lacuna | Completar Role×Permission×Resource |
| P2 | Compliance não mapeado | Mapear requisito regulatório → controle |

## Skills Recomendados
- `gap-analysis`, `senior-security`, `security-best-practices`
- `security-threat-model`, `threat-modeling-expert`
- `gdpr-compliant`, `privacy-by-design`

🤖 *Fix — Fase 3 do Downstream Architecture Refinement*
