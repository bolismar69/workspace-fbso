# PROMPT-FIX-SOURCING-FACTORY-BIDDING-ESTIMATE-VALIDATION

## Contexto

Este prompt implementa o **FIX do ESTIMATE-VALIDATION** — Fase 5. Acionado quando o GATE encontra NCs.

**Postura do FIX:** Cirúrgico e contido. Corrige APENAS as NCs apontadas pelo GATE. NUNCA reescreve o documento inteiro.

**Input:** Relatório de NCs do GATE (IDs [VALID-XX], localizações, sugestões).

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos |
| `{PROJECT_ID_NAME}` | Identificador do projeto |
| `{SOURCING_BIDDING_MODE}` | Modo: `agile-discovery`, `agile-refinement`, `waterfall-discovery` ou `waterfall-refinement` |

## Fluxo de Correção

### Passo 0 — Carregar Relatório do GATE
IDs de conflito, localizações e sugestões.

### Passo 1 — Priorizar NCs

| Prioridade | Tipo | Ação |
|---|---|---|
| P0 | QA/Arch/Prazo calculado errado | Recalcular com fórmula correta do DTA §2 |
| P0 | PIB não calculado | Calcular PIB Score com baseline do modo correto |
| P1 | Fábrica sem veredito | Emitir veredito claro (aprovada/rejeitada/ressalva) |
| P1 | Arquivo individual por fábrica ausente | Gerar ESTIMATE-VALIDATION-{FAB}.md |
| P2 | Tabela PIB por épico incompleta | Completar com baseline ROM/PERT por épico |

### Passo 2 — Aplicar Correções Cirúrgicas
Para cada NC: localizar seção exata → aplicar menor correção → verificar.

### Passo 3 — Validar Correções
Todas NCs endereçadas. Se não resolvida → `[NÃO RESOLVIDA]` com justificativa.

**Regra de Ouro:** Nunca reescreva o documento do zero. NUNCA altere vereditos da F5 sem recalcular com regras corretas.

## Skills

| 1 | `analyst-estimates` | Recálculo de métricas |
| 2 | `estimate-builder-qmohd` | Validação de QA |
| 3 | `gap-analysis` | Detecção de outliers |

🤖 *Fase 5 FIX*
