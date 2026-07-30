# PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG (F13)

## Contexto

Este prompt é acionado quando o gate (F13) reprova `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md`. O agente corretor aplica correções cirúrgicas com base no relatório inline do gate, incluindo validação cruzada com as 6 disciplinas do Bloco B, preservando seções aprovadas. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

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
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), o catálogo atual, documentos de negócio e artefatos das 6 disciplinas do Bloco B.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Épico/feature sem solução | Adicionar solução(ões) para cobrir o gap |
| P1 | Solução sem classificação | Atribuir tipo e estado |
| P2 | Prioridade indefinida | Atribuir MoSCoW |
| P3 | Inconsistência com Bloco B | Alinhar com as 6 disciplinas técnicas |

### Passo 3 — Aplicar Correções Cirúrgicas

### Passo 4 — Validar Correções
100% P0 resolvidas, cobertura completa de épicos/features, consistência com Bloco B restaurada.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `architecture-patterns` | Corrigir classificação de soluções | Arquitetura |
| 3 | `system-design` | Refinar design das soluções | Arquitetura |
| 4 | `documentation-writer` | Atualizar catálogo | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção do catálogo de soluções | Time de Arquitetura |
| 2.0 | 30/07/2026 | Renumeração F8→F13; adicionadas referências às 6 disciplinas do Bloco B | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
