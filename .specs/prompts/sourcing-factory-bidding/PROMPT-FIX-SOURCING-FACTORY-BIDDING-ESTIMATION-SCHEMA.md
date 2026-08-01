# PROMPT-FIX-SOURCING-FACTORY-BIDDING-ESTIMATION-SCHEMA

## Contexto

Este prompt implementa o **FIX do ESTIMATION-SCHEMA** — Fase 2. Acionado quando o GATE encontra NCs.

**Postura do FIX:** Cirúrgico e contido. Corrige APENAS as NCs apontadas pelo GATE. NUNCA reescreve o documento inteiro.

**Input:** Relatório de NCs do GATE (IDs [SCHEMA-XX], localizações, sugestões).

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
| P0 | Schema fora do padrão (20 colunas ausentes) | Corrigir imediatamente |
| P1 | Separador errado, placeholder não preenchido | Corrigir |
| P2 | Colunas de escopo com dados errados | Corrigir se simples |

### Passo 2 — Aplicar Correções Cirúrgicas
Para cada NC: localizar linha exata → aplicar menor correção → verificar.

### Passo 3 — Validar Correções
Todas NCs endereçadas. Se não resolvida → `[NÃO RESOLVIDA]` com justificativa.

**Regra de Ouro:** Nunca reescreva o CSV do zero. Cada correção é a menor mudança possível.

## Skills

| 1 | `estimate-builder` | Correção estrutural do schema |
| 2 | `documentation-writer` | Correção de instruções |

🤖 *Fase 2 FIX*
