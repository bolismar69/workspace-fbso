# PROMPT-FIX-410-INTAKE-LOG

## Contexto

Este prompt é acionado quando o `PROMPT-GATE-410-INTAKE-LOG.md` emite `[NÃO COMPLIANCE]` para o artefato `410-INTAKE-LOG.md`.

O agente corretor atua como **cirurgião de ingestão** — aplica correções pontuais no registro de lotes com base no relatório do gate, preservando todas as seções aprovadas. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas que foram apontadas como Não Compliance.**

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |

---

## Fluxo de Execução

### Passo 1 — Carregar Artefatos e Relatório do Gate
Ler `410-INTAKE-LOG.md` (artefato a corrigir), documentos de negócio (Charter, BRD, Epics, Features, User Stories) como referência, e o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano).

### Passo 2 — Processar Não-Conformidades por Prioridade

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Lote sem versionamento | Atribuir número de versão sequencial e data |
| P1 | Lote sem rastreabilidade | Adicionar referências aos documentos de negócio |
| P2 | Cobertura incompleta | Incluir documentos de negócio não cobertos em lote(s) existente(s) ou criar novo lote |
| P3 | Metadados ausentes | Preencher origem, responsável, tipo e status |
| P4 | Matriz de cobertura ausente | Gerar matriz de cobertura |

### Passo 3 — Aplicar Correções Cirúrgicas
Para cada NC: localizar seção/lote afetado, aplicar correção, preservar conteúdo não afetado.

### Passo 4 — Validar Correções
Verificar: 100% P0 resolvidas, todos os lotes versionados, rastreabilidade completa, cobertura total dos documentos de negócio, metadados preenchidos.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório do gate e priorizar NCs | Análise |
| 2 | `requirements-engineering` | Corrigir estruturação dos lotes | Requisitos |
| 3 | `business-analyst` | Validar escopo corrigido dos lotes | Análise |
| 4 | `documentation-writer` | Atualizar INTAKE-LOG.md com correções | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt de correção do registro de ingestão (F1 — Bloco 0) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
