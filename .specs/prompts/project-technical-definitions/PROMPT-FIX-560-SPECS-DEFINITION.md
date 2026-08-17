# PROMPT-FIX-560-SPECS-DEFINITION (F16)

## Contexto

Este prompt é acionado quando o gate (F16) reprova `560-SPECS-DEFINITION.md`. O agente corretor aplica correções cirúrgicas com base no relatório inline do gate, focando em restaurar referências quebradas, eliminar duplicação de conteúdo e garantir sumários concisos. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

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
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), o documento atual e TODOS os artefatos dos Blocos 0, A, B e C.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Artefato não referenciado | Adicionar seção com sumário + `→ ver [ARTEFATO]` |
| P0 | Conteúdo duplicado | Substituir por sumário conciso + referência |
| P1 | Link markdown inválido | Corrigir caminho do artefato |
| P2 | Sumário muito extenso | Reduzir para ~1 parágrafo |
| P3 | Seção ausente | Preencher sumário referenciando artefato correto |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções
100% artefatos referenciados, links válidos, sem duplicação.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `reference-builder` | Corrigir referências cruzadas | Mapeamento |
| 3 | `senior-architect` | Validar correções técnicas | Arquitetura |
| 4 | `documentation-writer` | Atualizar documento | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção da baseline de especificações | Time de Arquitetura |
| 2.0 | 30/07/2026 | Reformulação: correção focada em referências, links e eliminação de duplicação | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
