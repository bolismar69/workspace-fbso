# PROMPT-FIX-540-SOLUTIONS-MATRIX (F14)

## Contexto

Este prompt é acionado quando o gate (F14) reprova `540-SOLUTIONS-MATRIX.md`. O agente corretor aplica correções cirúrgicas com base no relatório inline do gate, incluindo validação cruzada com as 6 disciplinas do Bloco B. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

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
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), a matriz atual, Catálogo de Soluções (F13), artefatos do Bloco B, Stack Matrix, TEAM-CAPACITY.md.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Solução ausente da matriz | Adicionar linha completa para a solução |
| P0 | Stack inconsistente | Alinhar com STACK-MATRIX (F15) |
| P1 | Campo obrigatório vazio | Preencher com dado da fase correspondente |
| P2 | Indicador não calculado | Calcular cobertura de skills e risco |
| P3 | RACI incompleto | Completar matriz RACI |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `reference-builder` | Corrigir referências cruzadas | Mapeamento |
| 3 | `team-composition-analysis` | Corrigir alocação de time | People |
| 4 | `documentation-writer` | Atualizar matriz | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção da matriz de soluções | Time de Arquitetura |
| 2.0 | 30/07/2026 | Renumeração F9→F14; adicionadas referências às 6 disciplinas do Bloco B | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
