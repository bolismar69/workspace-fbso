# PROMPT-GENERATE-TASKS-ARTEFACT

## Contexto

Este prompt orquestra skills especializadas em planejamento de tarefas e gestão de backlog para gerar ou revisar o artefato `TECHNICAL-SOLUTION-TASKS.md` na pasta de especificações de uma solução técnica.

O artefato gerado deve ser a **lista de tarefas acionável** para o time de desenvolvimento — derivada das user stories, features e épicos do projeto, organizada por marco de entrega e priorizada conforme MoSCoW.

---

## Parâmetros de Entrada

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{SCOPE}` | Escopo da geração | `full`, `delta`, `sprint` (apenas sprint atual) |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 5 parâmetros foram informados. Se algum estiver ausente, perguntar antes de prosseguir.

### Passo 1 — Verificar e Preparar a Estrutura

```
Verificar se existe: {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/
    │
    ├── NÃO existe → Criar a pasta (mkdir -p)
    │
    └── SIM, existe →
            │
            ├── Verificar se existem TECHNICAL-SOLUTION-PRD.md, TECHNICAL-SOLUTION-SPECS.md E TECHNICAL-SOLUTION-ARCHITECTURE.md
            │     ├── TODOS existem → Ler os 3 como ponto de partida
            │     ├── PARCIAL → Ler o que existir + docs do projeto
            │     └── NENHUM → Ler documentos do projeto:
            │           ├── 01-PROJECT-CHARTER-*.md (entregas D1-D7, marcos M1-M7)
            │           ├── 04-FEATURES.md (18 features, 58 user stories)
            │           ├── 05-USER-STORIES-*.md (critérios de aceitação)
            │           ├── TECHNICAL-PLAN.md (sequenciamento, fases)
            │           └── DEFINITION_OF_DONE.md (critérios de DONE)
            │
            └── Verificar se TECHNICAL-SOLUTION-TASKS.md já existe:
                  ├── SIM + SCOPE=full → Regenerar (incrementar versão)
                  ├── SIM + SCOPE=delta → Atualizar tarefas concluídas/novas
                  └── NÃO → Criar do zero
```

### Passo 2 — Invocar Skills Especializadas

| Ordem | Skill | Responsabilidade |
|---|---|---|
| 1ª | `breakdown-epic-pm` | Decompor épicos em features e features em tarefas |
| 2ª | `writing-plans` | Estruturar plano de tarefas com dependências e estimativas |
| 3ª | `acceptance-criteria` | Vincular cada tarefa aos critérios de aceitação |

### Passo 3 — Gerar TECHNICAL-SOLUTION-TASKS.md

Gerar em: `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-TASKS.md`

#### Estrutura Obrigatória

```markdown
# TECHNICAL-SOLUTION-TASKS.md — Plano de Tarefas: {SOLUTION_NAME}

**Projeto:** {PROJECT_NAME}
**Solução:** {SOLUTION_NAME}
**Versão do TASKS:** {X.0}
**Data:** {data atual}
**Status:** {status}
**Situação impplementação:** {a ser preenchido pelos processos de desenvolvimento}

## 1. Visão Geral
- Total de tarefas, organizadas por marco (M2→M7)
- Progresso atual (X/Y concluídas)

## 2. Tarefas por Marco de Entrega
### M2: {nome} (data)
| ID | Tarefa | Feature | US | Prioridade | Status | Responsável | Estimativa |
### M3...M7 (repetir)

## 3. Dependências entre Tarefas
- Diagrama ou lista de precedência

## 4. Registro de Alterações
```

### Passo 4 — Validação Pós-Geração (8 verificações)

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|---|---|---|
| 1ª | `breakdown-epic-pm` | Decompor épicos → features → tarefas |
| 2ª | `writing-plans` | Estruturar plano com dependências e estimativas |
| 3ª | `acceptance-criteria` | Vincular tarefas aos critérios de aceitação |

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.1 | 16/07/2026 | Removido `{BRANCH_NAME}` — informação de branch é documentada no TECHNICAL-SOLUTION-PRD.md e SPRINT-CARD.md, não no TECHNICAL-SOLUTION-TASKS.md. Header do template sem linha `Branch:`. Parâmetros reduzidos de 6 para 5. | Time Técnico |
| 1.0 | 13/07/2026 | Criação inicial | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices.*
