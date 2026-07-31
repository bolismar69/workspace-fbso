# PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATION-SCHEMA

## Contexto

Este prompt implementa o **GENERATE do ESTIMATION-SCHEMA** para o processo de Sourcing & Factory Bidding (Fase 2).

**Propósito:** Gera o template CSV padronizado (DTA Estimation Schema) que as fábricas devem preencher e devolver.

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
### Passo 3 — GENERATE o Artefato
### Passo 4 — Validação Pós-GENERATE

## Skills Utilizados

| 1 | `estimate-builder` | Construção do template padronizado | 2 | `project-estimation` | Estrutura de colunas do schema |
| 3 | `afrexai-construction-estimator` | Metodologia de estimativa detalhada | 4 | `documentation-writer` | Instruções de preenchimento |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial — Fase 2 Sourcing & Factory Bidding | Time de Arquitetura |

🤖 *Sourcing & Factory Bidding — Fase 2 GENERATE*
