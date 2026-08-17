# PROMPT-FIX-550-SOLUTIONS-STACK-MATRIX (F15)

## Contexto

Este prompt é acionado quando o gate (F15) reprova `550-SOLUTIONS-STACK-MATRIX.md`. O agente corretor aplica correções cirúrgicas com base no relatório inline do gate, considerando as 6 disciplinas do Bloco B. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{ARCHITECTURE_GLOBAL}` | Caminho da pasta de arquitetura global |

---

## Fluxo de Execução

### Passo 1 — Carregar Relatório do Gate
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), a matriz atual, catálogo de soluções (F13), matriz (F14) e ADRs.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Solução sem stack | Definir stack completa com versões |
| P1 | Stack incompleta | Preencher dimensões faltantes |
| P2 | Versão "latest" | Especificar versão exata |
| P3 | Inconsistência cross-solution | Alinhar versões entre soluções integradas |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `tech-stack-evaluator` | Corrigir escolhas de stack | Avaliação |
| 3 | `senior-architect` | Validar correções arquiteturais | Arquitetura |
| 4 | `documentation-writer` | Atualizar matriz | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção da matriz de stacks | Time de Arquitetura |
| 2.0 | 30/07/2026 | Renumeração F10→F15; adicionadas referências às 6 disciplinas do Bloco B | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
