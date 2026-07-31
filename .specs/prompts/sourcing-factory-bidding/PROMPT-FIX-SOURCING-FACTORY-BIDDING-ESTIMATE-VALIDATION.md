# PROMPT-FIX-SOURCING-FACTORY-BIDDING-ESTIMATE-VALIDATION

## Contexto

Este prompt implementa o **FIX do ESTIMATE-VALIDATION** para o processo de Sourcing & Factory Bidding (Fase 5).

**Propósito:** Valida cada estimativa recebida contra regras DTA: formato, QA balanceado, outliers vs ROM interno.

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

| 1 | `analyst-estimates` | Análise e validação de estimativas | 2 | `estimate-builder` | Verificação estrutural |
| 3 | `estimate-builder-qmohd` | Validação de qualidade | 4 | `afrexai-construction-estimator` | Metodologia de verificação |
| 5 | `gap-analysis` | Detecção de outliers |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial — Fase 5 Sourcing & Factory Bidding | Time de Arquitetura |

🤖 *Sourcing & Factory Bidding — Fase 5 FIX*
