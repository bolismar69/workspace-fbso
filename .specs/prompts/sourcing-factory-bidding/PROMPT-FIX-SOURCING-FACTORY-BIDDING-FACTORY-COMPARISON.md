# PROMPT-FIX-SOURCING-FACTORY-BIDDING-FACTORY-COMPARISON

## Contexto

Este prompt implementa o **FIX do FACTORY-COMPARISON** para o processo de Sourcing & Factory Bidding (Fase 6).

**Propósito:** Consolida todas estimativas validadas em matriz comparativa com ranking e recomendação de seleção.

**Modo de operação:** Adapta-se ao `SOURCING_BIDDING_MODE` definido no Bootstrap (`discovery` ou `full`).

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{SOURCING_BIDDING_MODE}` | Modo: `discovery` ou `full` |
| `{SOURCING_BIDDING_PATH}` | Pasta sourcing-factory-bidding-{mode} |
| `{ESTIMATES_PATH}` | Pasta de estimativas recebidas |

## Fluxo de Execução

### Passo 0 — Validar Parâmetros e Modo
### Passo 1 — Carregar Artefatos Base (conforme modo)
### Passo 2 — Invocar Skills Especializadas
### Passo 3 — FIX o Artefato
### Passo 4 — Validação Pós-FIX

## Skills Utilizados

| 1 | `ads-budget` | Comparação orçamentária entre fábricas | 2 | `analyst-estimates` | Análise cross-fábrica |
| 3 | `trade-show-budget-planner` | ROI e modelo go/no-go por fábrica | 4 | `gap-analysis` | Análise de outliers e discrepâncias |
| 5 | `documentation-writer` | Relatório executivo de recomendação |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial — Fase 6 Sourcing & Factory Bidding | Time de Arquitetura |

🤖 *Sourcing & Factory Bidding — Fase 6 FIX*
