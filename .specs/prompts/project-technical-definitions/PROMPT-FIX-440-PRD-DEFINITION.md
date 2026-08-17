# PROMPT-FIX-440-PRD-DEFINITION

## Contexto

Este prompt é acionado quando o gate reprova `440-PRD-DEFINITION.md` (PRD de Negócio — Bloco 0, Fase 4). O agente corretor aplica correções cirúrgicas com base no relatório inline do gate. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

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
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), o PRD Definition atual, documentos de negócio e catálogo de soluções.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Requisito órfão (sem solução) | Mapear para solução existente ou justificar exclusão |
| P1 | Seção obrigatória ausente | Preencher seção completa |
| P2 | MVP Global vago | Detalhar escopo com critérios mensuráveis |
| P3 | Termo inconsistente | Alinhar com glossário do BRD |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções
100% P0 resolvidas, cobertura 100% do backlog, todas as seções obrigatórias completas.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `requirements-engineering` | Corrigir requisitos faltantes | Requirements |
| 3 | `prd-development` | Corrigir seções do PRD | Product |
| 4 | `documentation-writer` | Atualizar PRD Definition | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção do PRD Definition | Time de Arquitetura |
| 2.0 | 30/07/2026 | Atualização Bloco 0 (F4): adicionado contexto de fase (PRD de Negócio — Bloco 0, Fase 4) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
