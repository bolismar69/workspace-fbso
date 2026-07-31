# PROMPT-GATE-SOURCING-FACTORY-BIDDING-RFQ-PACKAGE

## Contexto

Este prompt implementa o **GATE do RFQ-PACKAGE** para o processo de Sourcing & Factory Bidding (Fase 1).

**Propósito:** Compila artefatos técnicos em um pacote RFQ padronizado para envio às fábricas.

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
### Passo 3 — GATE o Artefato
### Passo 4 — Validação Pós-GATE

## Skills Utilizados

| 1 | `project-estimation` | Estrutura de estimativas | 2 | `senior-architect` | Validação técnica do pacote |
| 3 | `documentation-writer` | Documentação do RFQ |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial — Fase 1 Sourcing & Factory Bidding | Time de Arquitetura |

🤖 *Sourcing & Factory Bidding — Fase 1 GATE*
