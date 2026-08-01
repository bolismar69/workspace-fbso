# PROMPT-FIX-SOURCING-FACTORY-BIDDING-FACTORY-COMPARISON

## Contexto

Este prompt implementa o **FIX do FACTORY-COMPARISON** — Fase 6. Acionado quando o GATE encontra NCs.

**Postura do FIX:** Cirúrgico e contido. Corrige APENAS as NCs apontadas pelo GATE. NUNCA reescreve o documento inteiro.

**Input:** Relatório de NCs do GATE (IDs [COMP-XX], localizações, sugestões).

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos |
| `{PROJECT_ID_NAME}` | Identificador do projeto |
| `{SOURCING_BIDDING_MODE}` | Modo: `discovery` ou `full` |

## Fluxo de Correção

### Passo 0 — Carregar Relatório do GATE
IDs de conflito, localizações e sugestões.

### Passo 1 — Priorizar NCs

| Prioridade | Tipo | Ação |
|---|---|---|
| P0 | Pesos não somam 100% ou critérios errados (4 em vez de 5) | Corrigir pesos conforme DTA §3.2 |
| P0 | PIB não integrado ao ranking | Integrar PIB como 5º critério (15%) |
| P1 | Cálculo de nota ponderada errado | Recalcular |
| P1 | Fábrica rejeitada na F5 incluída no ranking | Remover do ranking |
| P2 | Justificativa com < 3 razões | Expandir justificativa |

### Passo 2 — Aplicar Correções Cirúrgicas
Para cada NC: localizar seção exata → aplicar menor correção → verificar.

### Passo 3 — Validar Correções
Todas NCs endereçadas. Se não resolvida → `[NÃO RESOLVIDA]` com justificativa.

**Regra de Ouro:** Nunca reescreva o documento do zero. Cada correção é a menor mudança possível.

## Skills

| 1 | `ads-budget` | Correção de pesos e valores |
| 2 | `analyst-estimates` | Recálculo de ranking |
| 3 | `documentation-writer` | Correção de justificativas |

🤖 *Fase 6 FIX*
