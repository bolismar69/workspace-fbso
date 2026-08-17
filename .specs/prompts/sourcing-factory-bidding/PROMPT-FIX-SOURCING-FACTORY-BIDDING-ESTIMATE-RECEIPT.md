# PROMPT-FIX-SOURCING-FACTORY-BIDDING-ESTIMATE-RECEIPT

## Contexto

Este prompt implementa o **FIX do ESTIMATE-RECEIPT** — Fase 4. Acionado quando o GATE encontra NCs.

**Postura do FIX:** Cirúrgico e contido. Corrige APENAS as NCs apontadas pelo GATE. NUNCA reescreve o documento inteiro.

**Input:** Relatório de NCs do GATE (IDs [RECEIPT-XX], localizações, sugestões).

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
| P0 | Padrão de arquivo errado (extensão .md em vez de .csv) | Corrigir imediatamente |
| P1 | Checklist não reflete arquivos reais em estimates/ | Atualizar tabela |
| P2 | Total de horas ou data ausentes | Preencher dados |

### Passo 2 — Aplicar Correções Cirúrgicas
Para cada NC: localizar seção exata → aplicar menor correção → verificar.

### Passo 3 — Validar Correções
Todas NCs endereçadas. Se não resolvida → `[NÃO RESOLVIDA]` com justificativa.

**Regra de Ouro:** Nunca reescreva o documento do zero. Cada correção é a menor mudança possível.

## Skills

| 1 | `documentation-writer` | Correção do guia |
| 2 | `code-reviewer` | Revisão das correções |

🤖 *Fase 4 FIX*
