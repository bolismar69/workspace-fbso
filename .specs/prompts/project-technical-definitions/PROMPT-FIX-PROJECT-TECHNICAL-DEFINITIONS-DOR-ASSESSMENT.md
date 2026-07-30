# PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT

## Contexto

Este prompt é acionado quando o `PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT.md` emite `[NÃO COMPLIANCE]` para o artefato `PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT.md`.

O agente corretor atua como **cirurgião de prontidão** — aplica correções pontuais na avaliação DoR com base no relatório do gate, preservando todas as seções aprovadas. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas que foram apontadas como Não Compliance.**

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
Ler `PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT.md` (artefato a corrigir), `PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md` (referência de itens) e o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano).

### Passo 2 — Processar Não-Conformidades por Prioridade

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Item do INTAKE-LOG não avaliado | Avaliar item contra checklist DoR e registrar resultado |
| P0 | Item devolvido sem justificativa | Adicionar justificativa com critério(s) não atendido(s) |
| P1 | Item aprovado sem atender todos os critérios | Reavaliar e atualizar status para devolvido ou complementar critérios |
| P2 | Matriz de pendências ausente ou incompleta | Preencher ação necessária e responsável para cada item devolvido |
| P3 | Assinatura PO/PM ausente | Adicionar campo de sign-off formal |

### Passo 3 — Aplicar Correções Cirúrgicas
Para cada NC: localizar seção/item afetado, aplicar correção, preservar conteúdo não afetado.

### Passo 4 — Validar Correções
Verificar: 100% P0 resolvidas, todos os itens avaliados, rejeições justificadas, itens aprovados em conformidade com o DoR, sign-off presente.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `requirements-validation` | Corrigir avaliação de critérios DoR | Requisitos |
| 2 | `gap-analysis` | Analisar pendências e justificativas | Análise |
| 3 | `acceptance-criteria` | Ajustar verificação de testabilidade | Qualidade |
| 4 | `documentation-writer` | Atualizar DOR-ASSESSMENT.md com correções | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt de correção da avaliação DoR (F2 — Bloco 0) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
