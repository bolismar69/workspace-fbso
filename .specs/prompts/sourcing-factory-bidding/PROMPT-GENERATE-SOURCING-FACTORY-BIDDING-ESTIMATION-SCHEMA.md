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
Confirmar `{PROJECT_PATH}`, `{PROJECT_ID_NAME}`, `{SOURCING_BIDDING_MODE}`.

### Passo 1 — Carregar Artefatos Base
- Discovery: `upstream-architecture-discovery/DISCOVERY-LEVEL-PRD.md` (épicos, escopo)
- Full: features + user stories do projeto

### Passo 2 — Invocar Skills Especializadas
- `estimate-builder` — Template CSV com 20 colunas do schema unificado
- `project-estimation` — Estrutura de colunas conforme DTA-VALIDATION-STANDARDS §2.5

### Passo 3 — GENERATE o Artefato

**Especificações do Artefato:**

1. **Formato:** CSV com separador ponto-e-vírgula (;)
2. **Schema (20 colunas unificadas):** `fabrica; id_epico; titulo; features_codigos; qtd_features; user_stories_codigos; qtd_user_stories; horas_dev; horas_qa; horas_arch; horas_devops; horas_gestao; total_horas; prazo_entrega_meses; time_estimado_pessoas; valor_estimado; complexidade; stack_aderencia; premissas; comentarios`
3. **Placeholder:** Cada linha de épico usa `{FABRICA}` como placeholder na coluna `fabrica`
4. **Dados pré-preenchidos:** `id_epico`, `titulo`, `features_codigos`, `qtd_features`, `user_stories_codigos`, `qtd_user_stories` vêm dos artefatos do projeto
5. **Colunas numéricas vazias:** `horas_dev` a `comentarios` deixadas em branco para preenchimento pelas fábricas
6. **Cabeçalho:** Linha 1 com nomes exatos das 20 colunas

### Passo 4 — Validação Pós-GENERATE
Verificar: 20 colunas presentes, separador ;, épicos corretos, placeholders {FABRICA} preenchidos.

## Skills Utilizados

| 1 | `estimate-builder` | Construção do template padronizado | 2 | `project-estimation` | Estrutura de colunas do schema |
| 3 | `afrexai-construction-estimator` | Metodologia de estimativa detalhada | 4 | `documentation-writer` | Instruções de preenchimento |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial — Fase 2 Sourcing & Factory Bidding | Time de Arquitetura |

🤖 *Sourcing & Factory Bidding — Fase 2 GENERATE*
