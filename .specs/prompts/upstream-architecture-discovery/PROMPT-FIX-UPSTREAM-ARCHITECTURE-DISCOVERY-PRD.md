# PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-PRD

## Contexto

Este prompt é acionado quando o gate reprova `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md`. O agente corretor aplica correções cirúrgicas com base no relatório inline do gate. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

**Princípio fundamental:** O PRD Discovery-Level deve cobrir todos os Épicos definidos pelo Negócio com uma visão macro suficiente para embasar a análise de viabilidade técnica.

**Inputs upstream:** *(documentos de negócio)*
1. `{PROJECT_COMPLETE_PATH_NAME}/01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` — Project Charter
2. `{PROJECT_COMPLETE_PATH_NAME}/02-BRD-{PROJECT_ID_NAME}.md` — Business Requirements Document
3. `{PROJECT_COMPLETE_PATH_NAME}/03-EPICS-{PROJECT_ID_NAME}.md` — Índice de Épicos
4. `{PROJECT_COMPLETE_PATH_NAME}/epics/*.md` — Arquivos individuais de épicos

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de soluções técnicas |
| `{ARCHITECTURE_GLOBAL}` | Caminho da arquitetura global |
| `{SECURITY_GLOBAL}` | Caminho do GLOBAL-SECURITY.md |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Documentos brutos adicionais |
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
Validar que o artefato a ser corrigido `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` existe e que os documentos de negócio (Charter, BRD, Épicos) estão acessíveis.

### Passo 1 — Carregar Relatório do Gate e Artefatos
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), o documento atual e os documentos de negócio (Charter, BRD, Épicos incluindo `epics/*.md`).

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Épico sem cobertura no PRD | Adicionar visão/escopo macro do épico |
| P0 | Épico sem solução mapeada | Vincular cada épico a ≥1 solução técnica |
| P1 | Seção obrigatória ausente | Criar seção faltante (Visão, Matriz, MVP Macro, Restrições, Glossário) |
| P2 | MVP Macro indefinido | Definir escopo claro do MVP Macro |
| P3 | Inconsistência com Charter/BRD | Alinhar com os documentos de negócio |

### Passo 3 — Aplicar Correções Cirúrgicas
Aplicar as correções somente nas seções, tabelas ou linhas apontadas como NC, preservando o restante do documento.

### Passo 4 — Validar Correções
Revalidar cada NC contra o relatório do gate e garantir que todas foram resolvidas sem introduzir novas inconsistências.

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
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` | Artefato corrigido (F1) |
| 2 | `{PROJECT_COMPLETE_PATH_NAME}/01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` | Project Charter |
| 3 | `{PROJECT_COMPLETE_PATH_NAME}/02-BRD-{PROJECT_ID_NAME}.md` | Business Requirements Document |
| 4 | `{PROJECT_COMPLETE_PATH_NAME}/03-EPICS-{PROJECT_ID_NAME}.md` | Índice de Épicos |
| 5 | `{PROJECT_COMPLETE_PATH_NAME}/epics/*.md` | Épicos individuais |
| 6 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 7 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |

---

🤖 *Upstream Architecture Discovery — Fase 1 FIX*
