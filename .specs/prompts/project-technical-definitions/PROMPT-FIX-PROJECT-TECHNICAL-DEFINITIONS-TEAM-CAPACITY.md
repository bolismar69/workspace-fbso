# PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY

## Contexto

Este prompt é acionado quando o `PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` emite `[NÃO COMPLIANCE]` para o artefato `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` (Bloco A — Fase 6). O agente corretor aplica correções cirúrgicas com base no relatório inline do gate. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

---

## Fluxo de Execução

### Passo 1 — Carregar Relatório do Gate e Artefatos
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` (artefato a corrigir) e `PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md` (referência de papéis).

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Papel do Discovery Team sem entrada no CAPACITY | Adicionar linha para o papel com capacidade a definir |
| P0 | Papel órfão no CAPACITY (sem correspondência no SKILLS-MAP) | Conforme decisão do humano: manter (papel adicional) ou remover |
| P1 | Capacidade semanal não definida | Preencher coluna de capacidade (horas/dia para cada dia da semana) |
| P2 | Nome no CAPACITY não confere com SKILLS-MAP | Alinhar nomes conforme definição do humano |
| P3 | Coluna ou seção ausente | Completar estrutura conforme modelo |

### Passo 3 — Aplicar Correções Cirúrgicas
Para cada NC: localizar seção/linha afetada, aplicar correção, preservar conteúdo não afetado.

### Passo 4 — Validar Correções
Verificar: 100% P0 resolvidas, todos os papéis do Discovery Team cobertos, sem órfãos não intencionais, capacidades definidas.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório do gate e priorizar | Análise |
| 2 | `team-composition-analysis` | Corrigir alocação de time | People |
| 3 | `project-manager` | Validar correções de capacidade | PM |
| 4 | `documentation-writer` | Atualizar TEAM-CAPACITY.md | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 28/07/2026 | Criação inicial: prompt de correção da capacidade do time (Fase 2) | Time de Arquitetura |
| 2.0 | 30/07/2026 | Atualização Bloco A (F5-F6): adicionado contexto de fase (Bloco A — Fase 6); alinhamento com novo roadmap | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
