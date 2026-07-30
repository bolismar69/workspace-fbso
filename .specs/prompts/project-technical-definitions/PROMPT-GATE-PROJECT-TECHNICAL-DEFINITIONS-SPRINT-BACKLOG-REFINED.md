# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SPRINT-BACKLOG-REFINED

## Contexto

Este prompt implementa o **Gate de Validação do Backlog Refinado de Tarefas Técnicas** para o artefato `technical-discovery/SPRINT-BACKLOG.md`. Verifica se 100% das User Stories do backlog de negócio estão cobertas por tarefas T-NNN, se os links na coluna CONTRACTS são válidos e se os status são consistentes.

**Princípio fundamental:** Toda US do backlog de negócio deve ter ao menos uma tarefa técnica correspondente. Nenhuma tarefa pode ficar sem US vinculada ou sem sprint-alvo definida.

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

**Arquivos gerados pelo GENERATE:** `technical-discovery/SPRINT-BACKLOG.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `technical-discovery/SPRINT-BACKLOG.md`, PRODUCT-BACKLOG-LIST (F3), User Stories, MILESTONES (F17).

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de User Stories
| # | Verificação | Critério |
|---|---|---|
| 1.1 | US do backlog cobertas | Toda US do PRODUCT-BACKLOG-LIST tem ao menos uma tarefa T-NNN |
| 1.2 | US-ID existe | Toda US-ID na tabela existe nos documentos de User Stories |
| 1.3 | Tarefas sem US | Nenhuma tarefa deve ficar sem US vinculada |

#### Dimensão 2: Integridade das Tarefas
| # | Verificação | Critério |
|---|---|---|
| 2.1 | TASK-ID sequencial | IDs T-NNNNNN sequenciais sem duplicatas ou saltos |
| 2.2 | TASK-DESCRIÇÃO informativa | Descrição segue padrão verbo + objeto técnico |
| 2.3 | SPRINT-ALVO definido | Toda tarefa tem sprint-alvo coerente com MILESTONES |
| 2.4 | STATUS válido | Status pertence ao conjunto {TODO, IN-PROGRESS, IN-REVIEW, IN-TESTING, DONE, BLOCKED} |

#### Dimensão 3: Coluna CONTRACTS
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Links preenchidos | Toda tarefa tem ao menos um link na coluna CONTRACTS |
| 3.2 | Links válidos | Links markdown apontam para arquivos .md no caminho esperado |
| 3.3 | Placeholders aceitáveis | Sprints futuras podem ter links para arquivos ainda não criados |

#### Dimensão 4: Resumo por Sprint
| # | Verificação | Critério |
|---|---|---|
| 4.1 | Consistência numérica | Totais do resumo batem com a contagem da tabela |
| 4.2 | Sprint-alvo alinhada | Sprints do resumo refletem as sprints-alvo do MILESTONES |
| 4.3 | Datas informadas | Datas de início e término previsto preenchidas |

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE SPRINT-BACKLOG: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-SB-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o PRODUCT-BACKLOG-LIST/US/MILESTONES determinavam:** [Descrever a referência]
  - **Impacto:** [O risco de lacuna de rastreabilidade entre negócio e tarefas]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir o backlog refinado, por favor, responda:
1. Quanto ao **[ID-CONFLITO-SB-01]**, qual é a tarefa ou US correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SPRINT-BACKLOG-REFINED.md)*

---

### ✅ CENÁRIO B: SE O BACKLOG REFINADO ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE SPRINT-BACKLOG: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `technical-discovery/SPRINT-BACKLOG.md` gerado conforme PRODUCT-BACKLOG-LIST e MILESTONES.
- **AUDITORIA DA IA:** Cobertura de US validada. Tarefas T-NNN sequenciais e sem duplicatas. Todas as US do backlog de negócio cobertas. Coluna CONTRACTS preenchida. Resumo por sprint consistente com a tabela.
- **DIRETRIZ:** Peço que leia o backlog refinado para verificar se a decomposição em tarefas está adequada para o planejamento das sprints.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. O backlog refinado está em compliance e a decomposição das US em tarefas T-NNN está adequada?
2. Deseja enviar mais documentos/arquivos para enriquecer o backlog?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se "Sim, Não, Não" → [STATUS: COMPLIANCE] e Fase 19 (DISCOVERY TÉCNICO). Se novos inputs → retrocede ao PROMPT-GENERATE).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `scrum-master` | Validar decomposição de US em tarefas | Agile |
| 2 | `backlog-management` | Validar integridade do backlog técnico | Agile |
| 3 | `gap-analysis` | Identificar US sem tarefas ou tarefas sem US | Análise |
| 4 | `project-manager` | Validar alinhamento com milestones | PM |
| 5 | `kanban-method` | Validar fluxo de status das tarefas | Agile |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: gate de validação do backlog refinado de tarefas técnicas | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
