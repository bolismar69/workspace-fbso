# PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-PRD

## Contexto

Este prompt implementa o **GATE do PRD Discovery-Level** para o artefato `DISCOVERY-LEVEL-PRD.md` (Fase 1 — Bloco 0).

**Princípio fundamental:** O PRD Discovery-Level deve cobrir todos os Épicos definidos pelo Negócio com uma visão macro suficiente para embasar a análise de viabilidade técnica.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — artefato auditado (F1)
2. `{PROJECT_COMPLETE_PATH_NAME}/01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` — Project Charter
3. `{PROJECT_COMPLETE_PATH_NAME}/02-BRD-{PROJECT_ID_NAME}.md` — Business Requirements Document
4. `{PROJECT_COMPLETE_PATH_NAME}/03-EPICS-{PROJECT_ID_NAME}.md` — Índice de Épicos
5. `{PROJECT_COMPLETE_PATH_NAME}/epics/*.md` — Arquivos individuais de épicos

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de soluções técnicas |
| `{ARCHITECTURE_GLOBAL}` | Caminho da arquitetura global |
| `{SECURITY_GLOBAL}` | Caminho do GLOBAL-SECURITY.md |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Diretiva) Checkpoint HITL: sempre solicitar ao usuário se deseja fornecer informações adicionais ou novos direcionamentos via prompt |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
UPSTREAM_DISCOVERY_PATH       = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
```

**Arquivos gerados pelo GENERATE:** `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md`

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}` — base dos projetos de negócio
- `PROJECT_ID_NAME={PROJECT_ID_NAME}` — identificador do projeto
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
Validar que o artefato auditado `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` existe.

### Passo 1 — Carregar Documentos Base
Confirmar leitura dos seguintes artefatos:
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — artefato auditado (F1)
2. `{PROJECT_COMPLETE_PATH_NAME}/01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` — Project Charter
3. `{PROJECT_COMPLETE_PATH_NAME}/02-BRD-{PROJECT_ID_NAME}.md` — Business Requirements Document
4. `{PROJECT_COMPLETE_PATH_NAME}/03-EPICS-{PROJECT_ID_NAME}.md` — Índice de Épicos
5. `{PROJECT_COMPLETE_PATH_NAME}/epics/*.md` — Arquivos individuais de épicos

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de Épicos
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Cobertura total | 100% dos Épicos referenciados no PRD |
| 1.2 | Soluções mapeadas | Cada Épico vinculado a ≥1 solução |

#### Dimensão 2: Completude
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Seções presentes | Visão, Matriz, MVP Macro, Restrições, Glossário |
| 2.2 | MVP definido | Escopo do MVP Macro claro |

### Passo 3 — Emitir Veredito

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE — lista conflitos com ID, descrição, impacto, sugestão
### ✅ CENÁRIO B: PRÉ-COMPLIANCE — 3 perguntas obrigatórias de validação humana

*(Instrução de Orquestração: Se "Sim, Não, Não" → COMPLIANCE e Bloco B (F2). Se novos inputs → retrocede ao GENERATE).*

## Skills Utilizados

| Ordem | Skill | Propósito |
|---|---|---|
| 1 | `requirements-validation` | Validar cobertura de épicos |
| 2 | `gap-analysis` | Identificar épicos órfãos |
| 3 | `product-manager` | Validar MVP Macro |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial | Time de Arquitetura |

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` | Artefato auditado (F1) |
| 2 | `{PROJECT_COMPLETE_PATH_NAME}/01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` | Project Charter |
| 3 | `{PROJECT_COMPLETE_PATH_NAME}/02-BRD-{PROJECT_ID_NAME}.md` | Business Requirements Document |
| 4 | `{PROJECT_COMPLETE_PATH_NAME}/03-EPICS-{PROJECT_ID_NAME}.md` | Índice de Épicos |
| 5 | `{PROJECT_COMPLETE_PATH_NAME}/epics/*.md` | Arquivos individuais de épicos |
| 6 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 7 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |

---

🤖 *Upstream Architecture Discovery — Fase 1 GATE*
