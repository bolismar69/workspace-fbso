# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SPRINT-BACKLOG-REFINED

## Contexto

Este prompt gera o artefato `technical-discovery/SPRINT-BACKLOG.md` — o **backlog refinado de tarefas técnicas** do projeto. Diferente do `PRODUCT-BACKLOG-LIST` (F3), que lista itens de backlog de negócio priorizados, este documento decompõe cada User Story em tarefas técnicas T-NNN, associa a sprints-alvo e vincula a contratos técnicos.

**Objetivo:** Servir como índice mestre entre User Stories de negócio e tarefas técnicas, habilitando o rastreamento bidirecional US → Tarefas → Contratos.

**Estrutura do artefato:**
- `technical-discovery/SPRINT-BACKLOG.md` — backlog T-NNN enriquecido com sprints, status, datas e contratos

**Inputs upstream (Bloco C → Bloco D):**
- **PRODUCT-BACKLOG-LIST.md (F3):** Backlog priorizado de negócio para derivar tarefas técnicas
- **SPECS-DEFINITION.md (F16):** Baseline de especificações técnicas que as tarefas devem respeitar
- **MILESTONES.md (F17):** Roadmap de milestones com sprints-alvo para alinhamento temporal
- **PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md:** Milestones e dependências entre entregas
- **Documentos de negócio:** Charter, Features, User Stories para contexto adicional

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

### Passo 0 — Validação de Parâmetros
Verificar se TODOS os parâmetros foram informados.

### Passo 1 — Carregar Documentos Base
Ler PRODUCT-BACKLOG-LIST (F3), SPECS-DEFINITION (F16), MILESTONES (F17), Project Charter, Features, User Stories, Catálogo de Soluções.

### Passo 2 — Invocar Skills Especializadas
Invocar skills de planejamento ágil, backlog, sprint e gestão de projetos para decompor User Stories em tarefas técnicas T-NNN.

### Passo 3 — Gerar o Artefato

Criar o diretório `technical-discovery/` (se não existir) e gerar `{PROJECT_PATH}/technical-discovery/SPRINT-BACKLOG.md` com a estrutura abaixo.

#### 3.1 Cabeçalho do Documento
```markdown
# SPRINT-BACKLOG — Backlog Refinado de Tarefas Técnicas

- **Projeto:** {PROJECT_ID_NAME}
- **Versão:** 1.0
- **Data de Criação:** [DATA ATUAL]
- **Última Atualização:** [DATA ATUAL]
- **Total de Tarefas:** [N]
- **Total de Sprints:** [N]
- **Documento Base:** [PRODUCT-BACKLOG-LIST.md](../../business/project-documents/PRODUCT-BACKLOG-LIST.md)
```

#### 3.2 Status de Tarefa (Scrum/Kanban)
Definir a legenda de status que será usada no backlog:

| Status | Significado | Cor |
|--------|-------------|-----|
| **TODO** | Tarefa pendente, ainda não iniciada | 🔵 |
| **IN-PROGRESS** | Tarefa em desenvolvimento ativo | 🟡 |
| **IN-REVIEW** | Tarefa em revisão de código | 🟠 |
| **IN-TESTING** | Tarefa em teste (unitário, integração, QA) | 🔴 |
| **DONE** | Tarefa concluída e aprovada | 🟢 |
| **BLOCKED** | Tarefa bloqueada por dependência externa | ⚪ |

#### 3.3 Backlog de Tarefas (Tabela Enriquecida T-NNN)

Gerar a tabela completa de tarefas técnicas derivadas das User Stories, seguindo o modelo:

```markdown
| TASK-ID | TASK-DESCRIÇÃO | SPRINT-ALVO | US-ID | STATUS | DATA-INICIO | DATA-ENTREGA | CONTRACTS |
|---------|----------------|-------------|-------|--------|-------------|--------------|-----------|
| T-000010 | Auditar endpoint GET /dashboard/admin/summary | Sprint 01 | US-FEAT-EP-0001-0001-0001 | TODO | | | [API](sprint-01/CONTRACTS-API.md) · [DATA](sprint-01/CONTRACTS-DATA.md) · [SEC](sprint-01/CONTRACTS-SECURITY.md) · [SRE](sprint-01/CONTRACTS-SRE.md) |
```

**Regras de preenchimento:**
- **TASK-ID:** Formato `T-NNNNNN` sequencial (6 dígitos, zero-padded)
- **TASK-DESCRIÇÃO:** Ação + componente técnico (verbo + objeto). Ex: "Implementar endpoint POST /api/v1/users", "Criar migration de schema usuário", "Configurar alerta de latência no dashboard"
- **SPRINT-ALVO:** Sprint planejada conforme MILESTONES.md. Ex: "Sprint 01", "Sprint 02"
- **US-ID:** User Story de origem que a tarefa atende. Deve existir nos documentos de negócio
- **STATUS:** Conforme legenda da seção 3.2. Iniciar todas como TODO
- **DATA-INICIO:** (Opcional) Preenchido quando a tarefa é iniciada
- **DATA-ENTREGA:** (Opcional) Preenchido quando a tarefa é concluída
- **CONTRACTS:** Links para contratos técnicos da sprint. Para sprints futuras sem contratos gerados, manter apenas os links para os placeholders

#### 3.4 Critérios de Decomposição Tarefa × US

Cada User Story deve ser decomposta em tarefas seguindo:

1. **Tarefas de API:** Endpoints, handlers, validações de request/response
2. **Tarefas de Dados:** Migrations, schemas, queries, índices
3. **Tarefas de Segurança:** IAM, autenticação, autorização, audit logging
4. **Tarefas de SRE/DevOps:** Pipeline, observabilidade, deploy, infra
5. **Tarefas de Testes:** Testes unitários, integração, contrato, e2e
6. **Tarefas de Documentação:** ADRs, OpenAPI, README

#### 3.5 Resumo por Sprint

Após a tabela principal, adicionar um resumo consolidado por sprint:

```markdown
## Resumo por Sprint

### Sprint 01
| Métrica | Valor |
|---------|-------|
| Total de Tarefas | [N] |
| User Stories Vinculadas | [US-FEAT-EP-0001-0001-0001, US-FEAT-EP-0001-0002-0001] |
| Status | 5 TODO · 2 IN-PROGRESS · 0 DONE |
| Contratos | [API](sprint-01/CONTRACTS-API.md) · [DATA](sprint-01/CONTRACTS-DATA.md) · [SEC](sprint-01/CONTRACTS-SECURITY.md) · [SRE](sprint-01/CONTRACTS-SRE.md) |
| Início Previsto | [DATA] |
| Término Previsto | [DATA] |

### Sprint 02
...
```

#### 3.6 Referências

```markdown
## Referências

| Documento | Relação |
|:---|:---|
| [PRODUCT-BACKLOG-LIST.md](../../business/project-documents/PRODUCT-BACKLOG-LIST.md) | Backlog priorizado de negócio (F3) |
| [SPECS-DEFINITION.md](../PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md) | Baseline de especificações técnicas (F16) |
| [MILESTONES.md](../PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md) | Roadmap de milestones com sprints-alvo (F17) |
| [Project Charter](../../business/project-documents/PROJECT-CHARTER.md) | Documento de abertura do projeto |
| [User Stories](../../business/project-documents/) | User Stories do projeto |
```

### Passo 4 — Validação Pós-Geração
Verificar:
- TASK-ID sequencial sem duplicatas
- Toda tarefa vinculada a uma US existente
- SPRINT-ALVO coerente com os milestones do MILESTONES.md
- Coluna CONTRACTS preenchida (links válidos ou placeholders)
- Resumo por sprint reflete a tabela corretamente

---

## Modelo do Arquivo

O arquivo `technical-discovery/SPRINT-BACKLOG.md` deve seguir a estrutura consolidada abaixo:

```markdown
# SPRINT-BACKLOG — Backlog Refinado de Tarefas Técnicas

- **Projeto:** {PROJECT_ID_NAME}
- **Versão:** 1.0
- **Data de Criação:** [DATA ATUAL]
- **Última Atualização:** [DATA ATUAL]
- **Total de Tarefas:** [N]
- **Total de Sprints:** [N]

---

## 1. Status de Tarefa (Scrum/Kanban)

| Status | Significado | Cor |
|--------|-------------|-----|
| TODO | Tarefa pendente, ainda não iniciada | 🔵 |
| IN-PROGRESS | Tarefa em desenvolvimento ativo | 🟡 |
| IN-REVIEW | Tarefa em revisão de código | 🟠 |
| IN-TESTING | Tarefa em teste (unitário, integração, QA) | 🔴 |
| DONE | Tarefa concluída e aprovada | 🟢 |
| BLOCKED | Tarefa bloqueada por dependência externa | ⚪ |

---

## 2. Backlog de Tarefas

| TASK-ID | TASK-DESCRIÇÃO | SPRINT-ALVO | US-ID | STATUS | DATA-INICIO | DATA-ENTREGA | CONTRACTS |
|---------|----------------|-------------|-------|--------|-------------|--------------|-----------|
| [preencher conforme decomposição] |

---

## 3. Resumo por Sprint

[preencher para cada sprint com tarefas no backlog]

---

## 4. Referências

| Documento | Relação |
|:---|:---|
| [PRODUCT-BACKLOG-LIST.md](../../business/project-documents/PRODUCT-BACKLOG-LIST.md) | Backlog priorizado de negócio (F3) |
| [SPECS-DEFINITION.md](../PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md) | Baseline de especificações técnicas (F16) |
| [MILESTONES.md](../PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md) | Roadmap de milestones com sprints-alvo (F17) |

---

## Histórico de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | [DATA ATUAL] | Criação inicial: backlog refinado com tarefas T-NNN | Time de Arquitetura |
```

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados para esta fase. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `scrum-master` | Planejamento de sprints e decomposição de US em tarefas | Agile |
| 2 | `agile-sprint-planning` | Planejamento de sprints e alocação de tarefas | Agile |
| 3 | `backlog-management` | Gestão e refinamento do backlog técnico | Agile |
| 4 | `project-manager` | Coordenação entre tarefas, US e sprints | PM |
| 5 | `technical-change-tracker` | Rastreamento de mudanças técnicas no backlog | Engenharia |
| 6 | `documentation-writer` | Redigir o SPRINT-BACKLOG.md consolidado | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt gerador do backlog refinado de tarefas técnicas T-NNN | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência listados acima. Outros skills podem ser utilizados conforme aderência.*
