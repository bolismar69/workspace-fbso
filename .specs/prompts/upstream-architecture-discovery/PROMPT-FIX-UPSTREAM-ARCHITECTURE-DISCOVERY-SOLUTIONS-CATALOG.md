# PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-SOLUTIONS-CATALOG

## Contexto

Este prompt é acionado quando o gate reprova `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md`. O agente corretor aplica correções cirúrgicas com base no relatório inline do gate. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

**Princípio fundamental:** O artefato Discovery-Level deve conter informações suficientes para embasar a análise de viabilidade e estimativa ROM 50%, sem detalhamento excessivo de implementação.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Definição de Dados (F4)
5. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — Definição DevOps/SRE (F5)
6. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md` — Estratégia de Testes (F6)
7. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` — Definição Infra/Cloud (F7)

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Documentos brutos adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Diretiva) Checkpoint HITL: sempre solicitar ao usuário se deseja fornecer informações adicionais ou novos direcionamentos via prompt |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
UPSTREAM_DISCOVERY_PATH       = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
```

**Arquivos gerados pelo GENERATE:** `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md`

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}` — base dos projetos de negócio
- `PROJECT_ID_NAME={PROJECT_ID_NAME}` — identificador do projeto
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
Validar que o artefato a ser corrigido `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` existe e que os artefatos upstream (F1-F7) estão acessíveis.

### Passo 1 — Carregar Relatório do Gate e Artefatos
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), o documento atual e os artefatos upstream do Discovery (F1–F7).

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Solução do Bloco B ausente do catálogo | Adicionar solução (nome, tipo, propósito) |
| P1 | Tipo da solução incorreto ou ausente | Corrigir tipo da solução |
| P1 | Propósito high-level ausente | Adicionar propósito macro |
| P2 | Solução duplicada | Mesclar entradas duplicadas |
| P3 | Detalhamento de stack/implementação | Remover detalhe excessivo (visão macro) |

### Passo 3 — Aplicar Correções Cirúrgicas
Aplicar as correções somente nas seções, tabelas ou linhas apontadas como NC, preservando o restante do documento.

### Passo 4 — Validar Correções
Revalidar cada NC contra o relatório do gate e garantir que todas foram resolvidas sem introduzir novas inconsistências.

## Skills Utilizados

| Ordem | Skill | Propósito |
|---|---|---|
| 1 | `requirements-validation` | Validar completude do catálogo |
| 2 | `gap-analysis` | Identificar soluções órfãs |
| 3 | `senior-architect` | Avaliar nível de detalhe do catálogo |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial — Upstream Architecture Discovery | Time de Arquitetura |

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` | Artefato corrigido (F8) |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` | PRD Discovery-Level (F1) — referência de soluções |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` | Definição de Arquitetura (F2) — referência |
| 4 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 5 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |

---

🤖 *Upstream Architecture Discovery — SOLUTIONS-CATALOG FIX*
