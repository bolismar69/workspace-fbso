# PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATE-RECEIPT

## Contexto

Este prompt implementa o **GENERATE do ESTIMATE-RECEIPT** para o processo de Sourcing & Factory Bidding (Fase 4).

**Propósito:** Guia para o time operacional salvar estimativas recebidas no padrão: `ESTIMATION-SCHEMA-{NOME-DA-FABRICA}.csv` em `estimates/`.

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
Confirmar `{PROJECT_PATH}`, `{PROJECT_ID_NAME}`, `{SOURCING_BIDDING_MODE}`.

### Passo 1 — Carregar Artefatos Base
- FACTORY-DISTRIBUTION.md (F3) — lista de fábricas participantes

### Passo 2 — Invocar Skills Especializadas
- `analyst-estimates` — Análise de estimativas recebidas
- `documentation-writer` — Guia de recebimento

### Passo 3 — GENERATE o Artefato

**Especificações do Artefato:**

1. **Padrão de nomenclatura:** `ESTIMATION-SCHEMA-{NOME-DA-FABRICA}.csv` (extensão `.csv`)
2. **Pasta destino:** `estimates/`
3. **Checklist de recebimento** com colunas: `# | Fábrica | Arquivo | Data Recebimento | Total Horas | Status`
4. **Uma linha por fábrica** da F3, com nome de arquivo previsto
5. **Resumo** com total de fábricas contactadas, recebidas e mediana de horas

### Passo 4 — Validação Pós-GENERATE
Verificar: checklist reflete F3, padrão de arquivo `.csv`, pasta `estimates/` referenciada.

## Skills Utilizados

| 1 | `analyst-estimates` | Análise de estimativas recebidas | 2 | `documentation-writer` | Guia de recebimento |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial — Fase 4 Sourcing & Factory Bidding | Time de Arquitetura |

🤖 *Sourcing & Factory Bidding — Fase 4 GENERATE*
