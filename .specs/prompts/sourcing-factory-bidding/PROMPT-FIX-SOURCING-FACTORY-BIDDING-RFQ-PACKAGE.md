# PROMPT-FIX-SOURCING-FACTORY-BIDDING-RFQ-PACKAGE

## Contexto

Este prompt implementa o **FIX do RFQ-PACKAGE** — Fase 1 do Sourcing & Factory Bidding. O FIX é acionado quando o GATE encontra não-conformidades no artefato.

**Postura do FIX:** Cirúrgico e contido. Corrige APENAS as não-conformidades apontadas pelo GATE. NUNCA reescreve o documento inteiro. NUNCA altera seções não relacionadas às NCs.

**Input:** Relatório de não-conformidades do GATE (IDs de conflito, localizações, sugestões).

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{SOURCING_BIDDING_MODE}` | Modo: `agile-discovery`, `agile-refinement`, `waterfall-discovery` ou `waterfall-refinement` |
| `{SOURCING_BIDDING_PATH}` | Pasta sourcing-factory-bidding-{mode} |

## Fluxo de Correção

### Passo 0 — Carregar Relatório do GATE
Ler o relatório de NCs emitido pelo GATE. Identificar IDs de conflito, localizações exatas e sugestões.

### Passo 1 — Priorizar NCs
| Prioridade | Tipo | Ação |
|---|---|---|
| P0 | Bloqueante — artefato não pode ser usado | Corrigir imediatamente |
| P1 | Importante — compromete qualidade | Corrigir |
| P2 | Menor — ajuste cosmético ou documental | Corrigir se simples |

### Passo 2 — Aplicar Correções Cirúrgicas
Para cada NC: localizar seção/linha exata → aplicar correção mínima → verificar que a correção resolve o problema sem introduzir novos.

### Passo 3 — Validar Correções
Conferir que todas as NCs foram endereçadas. Se alguma NC não puder ser resolvida, reportar como `[NÃO RESOLVIDA]` com justificativa.

**Regra de Ouro:** Nunca reescreva o documento do zero. Cada correção deve ser a menor mudança possível que resolve a NC.

## Skills Utilizados

| 1 | `documentation-writer` | Correção cirúrgica de documentos |
| 2 | `code-reviewer` | Revisão das correções aplicadas |

🤖 *Sourcing & Factory Bidding — Fase 1 FIX*
