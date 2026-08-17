# PROMPT-FIX-430-PRODUCT-BACKLOG-LIST

## Contexto

Este prompt é acionado quando o `PROMPT-GATE-430-PRODUCT-BACKLOG-LIST.md` emite `[NÃO COMPLIANCE]` para o artefato `430-PRODUCT-BACKLOG-LIST.md`.

O agente corretor atua como **cirurgião de backlog** — aplica correções pontuais na lista de backlog com base no relatório do gate, preservando todas as seções aprovadas. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas que foram apontadas como Não Compliance.**

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
Ler `430-PRODUCT-BACKLOG-LIST.md` (artefato a corrigir), `420-DOR-ASSESSMENT.md` (referência de itens aprovados), `410-INTAKE-LOG.md` (referência de lotes) e o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano).

### Passo 2 — Processar Não-Conformidades por Prioridade

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Item aprovado no DoR ausente do backlog | Adicionar item com todos os campos obrigatórios |
| P0 | Item devolvido no DoR presente no backlog | Remover item do backlog |
| P1 | Item sem rastreabilidade de origem | Adicionar referência ao documento de negócio de origem |
| P1 | Link markdown inválido | Corrigir sintaxe ou caminho do link |
| P2 | Item sem prioridade MoSCoW | Atribuir prioridade com base nas informações disponíveis |
| P3 | Resumo por prioridade ausente | Gerar tabela de contagem e percentual |
| P3 | Resumo por lote ausente | Agrupar itens por lote de ingestão |

### Passo 3 — Aplicar Correções Cirúrgicas
Para cada NC: localizar seção/item afetado, aplicar correção, preservar conteúdo não afetado.

### Passo 4 — Validar Correções
Verificar: 100% P0 resolvidas, todos os itens aprovados no backlog, rastreabilidade completa com links válidos, priorização aplicada, resumos gerados.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `backlog-management` | Corrigir estrutura e composição do backlog | Gestão |
| 2 | `requirements-prioritization` | Ajustar priorização MoSCoW | Requisitos |
| 3 | `gap-analysis` | Identificar itens ausentes ou mal referenciados | Análise |
| 4 | `documentation-writer` | Atualizar PRODUCT-BACKLOG-LIST.md com correções | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt de correção do backlog de produto (F3 — Bloco 0) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
