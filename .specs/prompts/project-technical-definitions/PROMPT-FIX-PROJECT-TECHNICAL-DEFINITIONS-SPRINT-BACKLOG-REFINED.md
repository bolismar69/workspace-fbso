# PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SPRINT-BACKLOG-REFINED

## Contexto

Este prompt é acionado quando o gate reprova `technical-discovery/SPRINT-BACKLOG.md`. O agente corretor aplica correções cirúrgicas com base no relatório inline do gate. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |

---

## Fluxo de Execução

### Passo 1 — Carregar Relatório do Gate e Artefatos
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), o documento atual, PRODUCT-BACKLOG-LIST, User Stories, MILESTONES.

### Passo 2 — Processar NCs por Prioridade

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | US do backlog sem tarefa | Criar tarefa(s) T-NNN para cobrir a US |
| P0 | Tarefa sem US vinculada | Remover tarefa ou vincular à US correta |
| P0 | TASK-ID duplicado | Renumerar tarefa com ID sequencial único |
| P1 | SPRINT-ALVO incoerente com MILESTONES | Ajustar sprint-alvo conforme roadmap |
| P1 | Link CONTRACTS inválido | Corrigir caminho do link markdown |
| P2 | Descrição vaga da tarefa | Padronizar: verbo + objeto técnico |
| P2 | Status inválido | Substituir por valor do conjunto permitido |
| P3 | Resumo por sprint inconsistente | Recalcular totais e métricas conforme tabela |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar NCs | Análise |
| 2 | `backlog-management` | Corrigir backlog e tarefas | Agile |
| 3 | `scrum-master` | Ajustar decomposição de US em tarefas | Agile |
| 4 | `documentation-writer` | Atualizar documento | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt de correção do backlog refinado de tarefas técnicas | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
