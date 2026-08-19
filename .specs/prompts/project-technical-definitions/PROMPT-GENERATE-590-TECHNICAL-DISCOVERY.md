# PROMPT-GENERATE-590-TECHNICAL-DISCOVERY

## Contexto

Este prompt gera os **contratos técnicos por sprint** dentro do diretório `technical-discovery/590-ciclo-NNN/`. Este é um prompt **iterativo** — executa uma vez por sprint desejada, gerando 5 arquivos de contrato por sprint que detalham os acordos técnicos entre as disciplinas de API, Dados, Segurança e SRE, além da definição de incrementos retroativos.

**Objetivo:** Estabelecer contratos formais e rastreáveis entre as disciplinas técnicas para cada sprint, garantindo que todas as tarefas T-NNN tenham especificações técnicas detalhadas antes do início do desenvolvimento.

**Arquivos gerados por sprint:**
```
technical-discovery/590-ciclo-NNN/
├── CONTRACTS-API-ciclo-NNN.md       ← Endpoints, request/response, auth, rate limits
├── CONTRACTS-DATA-ciclo-NNN.md       ← Schemas, migrations, queries, índices
├── CONTRACTS-SECURITY-ciclo-NNN.md   ← Regras IAM, validações, threat model da sprint
├── CONTRACTS-SRE-ciclo-NNN.md        ← SLOs, dashboards, alertas, runbooks
└── DEFINITION-INCREMENTS-ciclo-NNN.md ← Atualizações retroativas nos docs base
```

**Inputs upstream (Bloco D):**
- **580-PACKAGE-BACKLOG-REFINED.md (F18):** Tarefas T-NNN da sprint atual, US vinculadas e contratos
- **SPECS-DEFINITION.md (F16):** Baseline de especificações técnicas cross-solution
- **Artefatos do Bloco B:** ARCHITECTURE-DEFINITION (F7), SECURITY-DEFINITION (F8), DATA-ARCHITECTURE-DEFINITION (F9), DEVOPS-SRE-DEFINITION (F10)

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
| `{SPRINT_NUMBER}` | Número da sprint alvo (ex: 01, 02, 03) |
| `{SPRINT_TASKS}` | Lista de TASK-ID da sprint conforme PACKAGE-BACKLOG |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Verificar se TODOS os parâmetros foram informados, especialmente `{SPRINT_NUMBER}` e `{SPRINT_TASKS}`.

### Passo 1 — Carregar Documentos Base
Ler PACKAGE-BACKLOG (F18), SPECS-DEFINITION (F16), ARCHITECTURE-DEFINITION (F7), SECURITY-DEFINITION (F8), DATA-ARCHITECTURE-DEFINITION (F9), DEVOPS-SRE-DEFINITION (F10) para fundamentar os contratos.

Filtrar do PACKAGE-BACKLOG apenas as tarefas T-NNN pertencentes à sprint corrente (`{SPRINT_NUMBER}`).

### Passo 2 — Invocar Skills Especializadas
Invocar skills de API design, modelagem de dados, segurança e SRE para gerar os contratos específicos de cada disciplina.

### Passo 3 — Gerar os Artefatos da Sprint

Criar o diretório `technical-discovery/sprint-{SPRINT_NUMBER}/` (se não existir) e gerar os 5 arquivos abaixo.

---

#### 3.1 CONTRACTS-API-sprint-{SPRINT_NUMBER}.md

Contrato de API para a sprint: define todos os endpoints, contratos request/response, autenticação, autorização e rate limits.

```markdown
# CONTRACTS-API — Sprint {SPRINT_NUMBER}

- **Projeto:** {PROJECT_ID_NAME}
- **Sprint:** {SPRINT_NUMBER}
- **Versão:** 1.0
- **Artefato Base:** [ARCHITECTURE-DEFINITION.md](../../470-ARCHITECTURE-DEFINITION.md)
- **Specs Base:** [SPECS-DEFINITION.md](../../560-SPECS-DEFINITION.md)

## Tarefas da Sprint

| TASK-ID | Descrição | US-ID |
|---------|-----------|-------|
| [tarefas da sprint com foco em API] |

## Contratos de API

### Endpoint [Método] /[path]
- **Descrição:** [descrição funcional]
- **Autenticação:** [Bearer JWT / API Key / OAuth2 / Public]
- **Escopo de Autorização:** [role/permission necessária]
- **Rate Limit:** [limite por intervalo]
- **Request:**
  - **Headers:** {[headers necessários]}
  - **Path Params:** {[parâmetros de path]}
  - **Query Params:** {[parâmetros de query]}
  - **Body:** {[schema do body]}
- **Response 200:**
  - **Body:** {[schema de resposta]}
- **Response 4xx/5xx:**
  - **Códigos:** {[possíveis erros]}

...

## Referências Cruzadas

| Endpoint | TASK-ID | US-ID | Security Contract |
|----------|---------|-------|-------------------|
| [endpoint] | [T-NNN] | [US-ID] | [link para CONTRACTS-SECURITY] |
```

---

#### 3.2 CONTRACTS-DATA-sprint-{SPRINT_NUMBER}.md

Contrato de dados: define schemas, migrations, queries e índices necessários para a sprint.

```markdown
# CONTRACTS-DATA — Sprint {SPRINT_NUMBER}

- **Projeto:** {PROJECT_ID_NAME}
- **Sprint:** {SPRINT_NUMBER}
- **Versão:** 1.0
- **Artefato Base:** [DATA-ARCHITECTURE-DEFINITION.md](../../490-DATA-ARCHITECTURE-DEFINITION.md)
- **Specs Base:** [SPECS-DEFINITION.md](../../560-SPECS-DEFINITION.md)

## Tarefas da Sprint

| TASK-ID | Descrição | US-ID |
|---------|-----------|-------|
| [tarefas da sprint com foco em dados] |

## Schema Changes

### [Nome da Tabela/Entidade]
- **Operação:** [CREATE / ALTER / DROP]
- **Colunas:** [nome, tipo, nullable, default, FK]
- **Índices:** [nome, colunas, unique?]
- **Migration:** [arquivo de migration]

## Queries Contratadas

| Query | Performance Esperada | Índice Necessário |
|-------|---------------------|-------------------|
| [descrição da query] | [tempo máximo] | [nome do índice] |

## Referências Cruzadas

| Tabela/Entidade | TASK-ID | US-ID | API Contract |
|-----------------|---------|-------|--------------|
| [entidade] | [T-NNN] | [US-ID] | [link para CONTRACTS-API] |
```

---

#### 3.3 CONTRACTS-SECURITY-sprint-{SPRINT_NUMBER}.md

Contrato de segurança: define regras IAM, validações de entrada, threat model específico da sprint e compliance.

```markdown
# CONTRACTS-SECURITY — Sprint {SPRINT_NUMBER}

- **Projeto:** {PROJECT_ID_NAME}
- **Sprint:** {SPRINT_NUMBER}
- **Versão:** 1.0
- **Artefato Base:** [SECURITY-DEFINITION.md](../../480-SECURITY-DEFINITION.md)
- **Specs Base:** [SPECS-DEFINITION.md](../../560-SPECS-DEFINITION.md)

## Tarefas da Sprint

| TASK-ID | Descrição | US-ID |
|---------|-----------|-------|
| [tarefas da sprint com foco em segurança] |

## Regras IAM

| Recurso | Ação | Role | Princípio |
|---------|------|------|-----------|
| [recurso] | [ação] | [role necessária] | [least privilege / separation of duties] |

## Validações de Entrada

| Campo | Regra | Código de Erro |
|-------|-------|----------------|
| [campo] | [regra de validação] | [HTTP status + código] |

## Threat Model da Sprint

| Ameaça | Mitigação | Severidade |
|--------|-----------|------------|
| [descrição] | [controle implementado] | [Alta/Média/Baixa] |

## Referências Cruzadas

| Controle | TASK-ID | US-ID | API/Data Contract |
|----------|---------|-------|-------------------|
| [controle] | [T-NNN] | [US-ID] | [link para CONTRACTS-API ou CONTRACTS-DATA] |
```

---

#### 3.4 CONTRACTS-SRE-sprint-{SPRINT_NUMBER}.md

Contrato de SRE: define SLOs, dashboards, alertas e runbooks para a sprint.

```markdown
# CONTRACTS-SRE — Sprint {SPRINT_NUMBER}

- **Projeto:** {PROJECT_ID_NAME}
- **Sprint:** {SPRINT_NUMBER}
- **Versão:** 1.0
- **Artefato Base:** [DEVOPS-SRE-DEFINITION.md](../../500-DEVOPS-SRE-DEFINITION.md)

## Tarefas da Sprint

| TASK-ID | Descrição | US-ID |
|---------|-----------|-------|
| [tarefas da sprint com foco em SRE/DevOps] |

## SLOs

| Serviço | Métrica | SLO | Janela |
|---------|---------|-----|--------|
| [serviço] | [latência/disponibilidade/throughput] | [valor] | [período] |

## Dashboards

| Dashboard | Métricas | Responsável |
|-----------|----------|-------------|
| [nome] | [métricas exibidas] | [time] |

## Alertas

| Alerta | Condição | Severidade | Ação |
|--------|----------|------------|------|
| [nome] | [condição de disparo] | [P1/P2/P3] | [runbook link] |

## Runbooks

| Incidente | Runbook | Responsável |
|-----------|---------|-------------|
| [tipo] | [link para runbook] | [time] |

## Referências Cruzadas

| Recurso | TASK-ID | US-ID | API/Data Contract |
|---------|---------|-------|-------------------|
| [serviço] | [T-NNN] | [US-ID] | [link para CONTRACTS-API ou CONTRACTS-DATA] |
```

---

#### 3.5 DEFINITION-INCREMENTS-sprint-{SPRINT_NUMBER}.md

Definição de incrementos: documenta as atualizações retroativas nos artefatos base necessárias após a conclusão da sprint.

```markdown
# DEFINITION-INCREMENTS — Sprint {SPRINT_NUMBER}

- **Projeto:** {PROJECT_ID_NAME}
- **Sprint:** {SPRINT_NUMBER}
- **Versão:** 1.0
- **Data:** [DATA ATUAL]

## Tarefas Concluídas

| TASK-ID | Descrição | US-ID | Contratos Impactados |
|---------|-----------|-------|----------------------|
| [tarefas concluídas na sprint] | | | |

## Atualizações Retroativas em Artefatos Base

### ARCHITECTURE-DEFINITION.md
- [ ] [Descrição da atualização necessária]
- **Responsável:** [papel]
- **Prazo:** [data]

### SECURITY-DEFINITION.md
- [ ] [Descrição da atualização necessária]
- **Responsável:** [papel]
- **Prazo:** [data]

### DATA-ARCHITECTURE-DEFINITION.md
- [ ] [Descrição da atualização necessária]
- **Responsável:** [papel]
- **Prazo:** [data]

### DEVOPS-SRE-DEFINITION.md
- [ ] [Descrição da atualização necessária]
- **Responsável:** [papel]
- **Prazo:** [data]

### 580-PACKAGE-BACKLOG-REFINED.md
- [ ] Atualizar status das tarefas concluídas para DONE
- [ ] Preencher DATA-INICIO e DATA-ENTREGA das tarefas
- [ ] Atualizar resumo por sprint

## Lições Aprendidas / Riscos Identificados

| Descoberta | Impacto | Ação |
|------------|---------|------|
| [aprendizado ou risco] | [impacto técnico] | [ação de mitigação] |
```

### Passo 4 — Validação Pós-Geração
Verificar para cada sprint:
- 5 arquivos de contrato criados no diretório correto
- Cada contrato referencia ao menos uma US-ID e um artefato base
- TASK-IDs consistentes com PACKAGE-BACKLOG
- Links markdown entre contratos são válidos

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados para esta fase. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `api-designer` | Definir contratos de API da sprint | API |
| 2 | `api-documentation` | Documentar endpoints, request/response, auth | API |
| 3 | `data-modeling` | Definir schemas, migrations e queries | Dados |
| 4 | `security-auditor` | Definir regras IAM e threat model por sprint | Segurança |
| 5 | `sre-engineer` | Definir SLOs, dashboards e alertas | SRE |
| 6 | `senior-architect` | Supervisão dos contratos cross-disciplina | Arquitetura |
| 7 | `documentation-writer` | Redigir contratos e incrementos | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt gerador iterativo de contratos técnicos por sprint | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência listados acima. Outros skills podem ser utilizados conforme aderência.*
