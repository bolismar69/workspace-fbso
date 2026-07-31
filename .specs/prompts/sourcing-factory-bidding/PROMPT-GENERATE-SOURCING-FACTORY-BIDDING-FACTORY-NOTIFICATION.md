# PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-FACTORY-NOTIFICATION

## Contexto

Este prompt implementa o **GENERATE do FACTORY-NOTIFICATION** — Fase 7 do Sourcing & Factory Bidding.

**Propósito:** Gerar notificações formais para todas as fábricas participantes informando o resultado do processo de seleção: carta de seleção para a vencedora, feedback técnico para as rejeitadas.

**Inputs upstream:** FACTORY-COMPARISON (F6) + ESTIMATE-VALIDATION (F5).

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{SOURCING_BIDDING_MODE}` | Modo: `discovery` ou `full` |
| `{SOURCING_BIDDING_PATH}` | Pasta sourcing-factory-bidding-{mode} |

## Fluxo de Execução

### Passo 0 — Validar Parâmetros
### Passo 1 — Carregar FACTORY-COMPARISON e ESTIMATE-VALIDATION
### Passo 2 — Invocar Skills
### Passo 3 — Gerar notificações: (a) Carta de Seleção para vencedora, (b) Carta de feedback para cada rejeitada com motivo específico, (c) Carta de segundo colocado quando aplicável
### Passo 4 — Validação Pós-Geração

## Skills Utilizados

| 1 | `documentation-writer` | Redação das cartas de notificação |
| 2 | `business-analyst` | Linguagem adequada ao relacionamento com fornecedores |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial — Fase 7 Sourcing & Factory Bidding | Time de Arquitetura |

🤖 *Sourcing & Factory Bidding — Fase 7 GENERATE*
