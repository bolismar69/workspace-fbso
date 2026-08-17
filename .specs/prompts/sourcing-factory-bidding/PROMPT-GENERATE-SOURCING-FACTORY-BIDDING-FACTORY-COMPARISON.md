# PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-FACTORY-COMPARISON

## Contexto

Este prompt implementa o **GENERATE do FACTORY-COMPARISON** para o processo de Sourcing & Factory Bidding (Fase 6).

**Propósito:** Consolida estimativas validadas em matriz comparativa conforme `.specs/standards/DTA-VALIDATION-STANDARDS.md` §3: Matriz de Decisão Ponderada, Pesos por Critério, Escala de Notas, Go/No-Go por Fábrica.

**Modo de operação:** Adapta-se ao `SOURCING_BIDDING_MODE` definido no Bootstrap (`agile-discovery`, `agile-refinement`, `waterfall-discovery` ou `waterfall-refinement`).

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{SOURCING_BIDDING_MODE}` | Modo: `agile-discovery`, `agile-refinement`, `waterfall-discovery` ou `waterfall-refinement` |
| `{SOURCING_BIDDING_PATH}` | Pasta sourcing-factory-bidding-{mode} |
| `{ESTIMATES_PATH}` | Pasta de estimativas recebidas |

## Fluxo de Execução

### Passo 0 — Validar Parâmetros e Modo
Confirmar `{PROJECT_PATH}`, `{PROJECT_ID_NAME}`, `{SOURCING_BIDDING_MODE}`.

### Passo 1 — Carregar Artefatos Base
- ESTIMATE-VALIDATION.md (F5) — vereditos por fábrica
- ESTIMATE-RETROSPECTIVE-PIB.md (F5b) — se disponível
- DTA-VALIDATION-STANDARDS.md §3 — matriz de decisão

### Passo 2 — Invocar Skills Especializadas
- `ads-budget` — Comparação orçamentária
- `analyst-estimates` — Análise cross-fábrica
- `trade-show-budget-planner` — ROI e go/no-go

### Passo 3 — GENERATE o Artefato

**Especificações do Artefato:**

1. **Critérios (5):** Tabela com pesos conforme DTA §3.2: Custo (25%), Prazo (25%), Qualidade (20%), PIB (15%), Consistência (15%)
2. **Matriz por critério:** Nota 1-10 para cada fábrica aprovada, com justificativa
3. **Ranking Final:** Tabela `# | Fábrica | Custo (25%) | Prazo (25%) | Qualidade (20%) | PIB (15%) | Consist. (15%) | Nota Ponderada`
4. **Recomendação:** Fábrica vencedora com ≥ 3 razões de seleção
5. **Fábricas rejeitadas:** Lista com motivo específico da F5
6. **Go/No-Go:** `Viabilidade = (Custo ≤ Budget_Máximo) AND (Prazo ≤ Prazo_Máximo) AND (Nota_Final ≥ 5.0)`

### Passo 4 — Validação Pós-GENERATE
Verificar: apenas aprovadas F5, pesos somam 100%, cálculo sem erro, recomendação justificada.

## Skills Utilizados

| 1 | `ads-budget` | Comparação orçamentária entre fábricas | 2 | `analyst-estimates` | Análise cross-fábrica |
| 3 | `trade-show-budget-planner` | ROI e modelo go/no-go por fábrica | 4 | `gap-analysis` | Análise de outliers e discrepâncias |
| 5 | `documentation-writer` | Relatório executivo de recomendação |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial — Fase 6 Sourcing & Factory Bidding | Time de Arquitetura |

🤖 *Sourcing & Factory Bidding — Fase 6 GENERATE*
