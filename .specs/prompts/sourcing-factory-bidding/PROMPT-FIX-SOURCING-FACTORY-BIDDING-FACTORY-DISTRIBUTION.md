# PROMPT-FIX-SOURCING-FACTORY-BIDDING-FACTORY-DISTRIBUTION

## Contexto

Este prompt implementa o **FIX do FACTORY-DISTRIBUTION** — Fase 3. Acionado quando o GATE encontra NCs.

**Postura do FIX:** Cirúrgico e contido. Corrige APENAS as NCs apontadas pelo GATE. NUNCA reescreve o documento inteiro.

**Input:** Relatório de NCs do GATE (IDs [DIST-XX], localizações, sugestões).

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
| P0 | < 2 fábricas cadastradas | Cadastrar fábricas faltantes |
| P1 | Canal/E-mail/Telefone ausentes | Preencher com ⚠️ ou dados reais |
| P2 | Status inconsistente | Corrigir conforme situação real |

### Passo 2 — Aplicar Correções Cirúrgicas
Para cada NC: localizar linha exata → aplicar menor correção → verificar.

### Passo 3 — Validar Correções
Todas NCs endereçadas. Se não resolvida → `[NÃO RESOLVIDA]` com justificativa.

**Regra de Ouro:** Nunca reescreva o documento do zero. Cada correção é a menor mudança possível.

## Skills

| 1 | `documentation-writer` | Correção de registro |
| 2 | `code-reviewer` | Revisão das correções |

🤖 *Fase 3 FIX*
