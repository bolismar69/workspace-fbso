# PROMPT-GENERATE-SPRINT-ARTEFACTS

## Contexto

Este prompt orquestra a geração dos **3 artefatos de uma sprint** a partir dos documentos-mestre do projeto (`TASKS.md`, `TEST_PLAN.md`, `SPECS.md`, `ARCHITECTURE.md`).

Os artefatos gerados são:

| Artefato | Propósito | Público |
|:---|:---|:---|
| `SPRINT-CARD.md` | Goal, backlog, DONE criteria, riscos e métricas da sprint | Time técnico |
| `SPRINT-TEST-SUITE.md` | Extrato do TEST_PLAN.md com cenários de teste aplicáveis à sprint | Time técnico / QA |
| `SPRINT-REVIEW.md` | Checklist de demonstração para o Product Owner na review | PO / Stakeholders |

**Princípio fundamental:** Os documentos-mestre (`TASKS.md`, `TEST_PLAN.md`, `SPECS.md`) são a **fonte da verdade**. Os artefatos de sprint são **derivados** — eles extraem, filtram e reorganizam informações já existentes, nunca criam conteúdo novo que não esteja ancorado nos documentos-mestre.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|:---|:---|:---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica (microsserviço) | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_NAME}` | Nome/código do projeto de negócio | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{SPRINT_NUMBER}` | Número da sprint (1 a N) | `1` |
| `{SPRINT_NAME}` | Nome curto da sprint (kebab-case) | `sprint-01-setup` |
| `{STACK}` | Stack tecnológica principal | `Java 25 + Spring Boot + PostgreSQL` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Antes de qualquer ação, verificar se TODOS os 6 parâmetros foram informados. Se algum estiver ausente, perguntar ao humano antes de prosseguir.

### Passo 1 — Carregar Documentos-Mestre

```
SPRINTS_INDEX = {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/sprints/README.md
SPRINT_DIR   = {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/sprints/{SPRINT_NAME}/

Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TASKS.md      ← Fonte da verdade para tarefas
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TEST_PLAN.md  ← Fonte da verdade para testes
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/SPECS.md      ← Fonte da verdade para specs
    └── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/ARCHITECTURE.md ← Estrutura e padrões

Se TASKS.md não existir → ERRO: "TASKS.md não encontrado. Execute o gerador de TASKS primeiro."
Se TEST_PLAN.md não existir → ERRO: "TEST_PLAN.md não encontrado. Execute o gerador de TEST_PLAN primeiro."
Se SPECS.md não existir → ERRO: "SPECS.md não encontrado. Execute o gerador de SPECS primeiro."
```

### Passo 2 — Identificar as Tarefas da Sprint

Extrair do `TASKS.md` §2 (ou seção equivalente) as tarefas que pertencem a esta sprint. Para cada tarefa:

- ID, descrição, feature associada, user stories, prioridade (Must/Should/Could), estimativa em dias, critério DONE
- Verificar o diagrama de dependências (§3 do TASKS.md) para identificar pré-requisitos e sucessores

### Passo 3 — Identificar os Cenários de Teste da Sprint

Extrair do `TEST_PLAN.md` os cenários de teste cujas features são implementadas nesta sprint:

- Cruzar as features da sprint (extraídas do TASKS.md) com as seções do TEST_PLAN.md (§3.1 a §3.18)
- Incluir também cenários transversais aplicáveis (ex: pipeline de segurança, isolamento multi-tenant)
- Preservar os IDs originais do TEST_PLAN.md para rastreabilidade bidirecional

### Passo 4 — Identificar Critérios de Aceitação e RNs

Extrair do `SPECS.md`:

- **§3.1** — Regras de Negócio formais que se aplicam às features desta sprint
- **§4.1** — Endpoints REST que serão implementados nesta sprint
- **§7** — Critérios de aceitação das features desta sprint

### Passo 5 — Gerar SPRINT-CARD.md

```
{SPRINT_DIR}/SPRINT-CARD.md
```

#### Estrutura Obrigatória

```markdown
# SPRINT-CARD: Sprint {N} — {Nome Descritivo}

[Header: sprint N de M, marco, datas, duração, responsável, docs-mestre com links]

## 🎯 Sprint Goal
[1-2 frases: o que esta sprint entrega de valor. Formato: "Verbo no imperativo. Métrica de sucesso."]

## 📋 Sprint Backlog
[Tabela: ID | Tarefa | Feature | Prio. | Est. | Critério DONE]
[Cada tarefa extraída do TASKS.md, com estimativa e critério DONE preservados]

## 📦 Features Entregues
[Tabela: Feature | Descrição | RNs Cobertas | Prio.]

## ✅ Definition of Done (Sprint-Level)
[Checklist de verificações que definem a sprint como concluída]

## ⚠️ Riscos e Bloqueadores
[Tabela: Risco | Prob. | Impacto | Mitigação]

## 🔗 Dependências
[Pré-requisitos (sprints anteriores). Sucessor (próxima sprint). Features que dependem desta sprint.]

## 📊 Métricas da Sprint
[Tabela: Métrica | Meta]

## Rodapé
[Indicação de geração por IA, referência ao TASKS.md]
```

### Passo 6 — Gerar SPRINT-TEST-SUITE.md

```
{SPRINT_DIR}/SPRINT-TEST-SUITE.md
```

#### Estrutura Obrigatória

```markdown
# SPRINT-TEST-SUITE: Sprint {N} — {Nome Descritivo}

[Header: sprint N de M, origem (TEST_PLAN.md), features cobertas, total de cenários]

## 1. {Feature ID}: {Nome da Feature} (N cenários)
[Tabela: ID | Descrição | Nível (Unit/Integração/E2E/Segurança) | Ref. TEST_PLAN (seção)]

## 2. {Feature ID}: ... (repetir para cada feature da sprint)

## 📊 Resumo
[Tabela: Nível | Cenários]

## 🔗 RNs Cobertas
[Tabela: RN | Descrição | Feature]
```

> **Regra de ouro:** Os IDs dos cenários devem ser EXATAMENTE os mesmos do TEST_PLAN.md. Se o TEST_PLAN.md referencia `TC-F02-01-001`, o SPRINT-TEST-SUITE.md deve usar o mesmo ID. Isso garante rastreabilidade bidirecional.

### Passo 7 — Gerar SPRINT-REVIEW.md

```
{SPRINT_DIR}/SPRINT-REVIEW.md
```

#### Estrutura Obrigatória

```markdown
# SPRINT-REVIEW: Sprint {N} — {Nome Descritivo}

[Header: sprint N de M, data da review, participantes, features]

## 🎯 O Que Demonstrar

### 1. {Feature / Funcionalidade}
- [ ] Item demonstrável (com checkboxes)
- [ ] Incluir scripts de demonstração (🎬 Script:) para features com interação visível

### 2. ...

## 📋 Pontos de Verificação (PO)
[Tabela: Verificação | Status (⬜)]

## 🚧 Bloqueios Identificados
[Tabela: Bloqueio | Ação | Responsável — preencher durante a review]

## ➡️ Próximo Passo
[Próxima sprint com datas e goal resumido]

## Rodapé
[Indicação de geração por IA]
```

> **Regra para scripts de demo:** Para features que têm interface visível (a partir da Sprint 3), incluir blocos `🎬 Script:` com falas prontas para o apresentador — o que dizer, o que clicar, qual resultado esperar. Isso elimina improvisação na hora da review com o PO.

### Passo 8 — Atualizar o Índice de Sprints

Atualizar o arquivo `{SPRINTS_INDEX}` (`sprints/README.md`):

- Se não existir, criá-lo com a estrutura completa (índice, timeline, progresso)
- Se existir, verificar se a sprint atual está listada. Se não estiver, adicioná-la à tabela de índice.

### Passo 9 — Validar os Artefatos Gerados

| # | Verificação | Critério |
|:---|:---|:---|
| 1 | SPRINT-CARD.md existe | Arquivo criado no path correto |
| 2 | Sprint backlog completo | Todas as tarefas do TASKS.md para esta sprint estão listadas |
| 3 | Estimativas preservadas | Cada tarefa no backlog tem estimativa em dias |
| 4 | Critérios DONE preservados | Cada tarefa no backlog tem critério DONE do TASKS.md |
| 5 | SPRINT-TEST-SUITE.md existe | Arquivo criado no path correto |
| 6 | IDs de cenário preservados | IDs batem exatamente com TEST_PLAN.md |
| 7 | Features cobertas | Todas as features da sprint têm cenários de teste |
| 8 | SPRINT-REVIEW.md existe | Arquivo criado no path correto |
| 9 | Scripts de demo inclusos | Features visíveis têm 🎬 Script: |
| 10 | Índice atualizado | README.md lista a sprint |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|:---|:---|:---|
| 1ª | `breakdown-plan` | Decompor o marco em tarefas gerenciáveis e extrair do TASKS.md |
| 2ª | `test-strategy-design` | Extrair e organizar cenários de teste por feature/sprint |
| 3ª | `acceptance-criteria` | Extrair critérios de aceitação do SPECS.md para o SPRINT-REVIEW.md |
| 4ª | `documentation-writer` | Qualidade e clareza dos 3 artefatos |
| 5ª | `ponytail` | Revisor final — os artefatos estão simples e diretos? Há conteúdo redundante? |

---

## Observações

1. **Documentos-mestre são a fonte da verdade.** Os artefatos de sprint NUNCA criam requisitos, tarefas ou cenários novos. Se algo está nos artefatos de sprint, DEVE estar nos documentos-mestre.

2. **Rastreabilidade bidirecional é mandatória.** IDs de tarefas (T-XXX), cenários de teste (TC-XXX-XXX) e RNs (RNXX-XX) devem ser preservados exatamente como nos documentos-mestre. Isso permite navegar do artefato de sprint → documento-mestre e vice-versa.

3. **Scripts de demo são investimento, não desperdício.** O tempo gasto escrevendo `🎬 Script:` no SPRINT-REVIEW.md se paga na hora da review com o PO — zero improvisação, zero "deixa eu ver como faz isso".

4. **O Sprint Goal é a bússola.** Se uma tarefa não contribui para o goal, ela não deveria estar nesta sprint. O goal deve ser específico e mensurável — "Implementar CRUD de tenants" é melhor que "Trabalhar no módulo de tenants".

5. **Sprint 1 e 2 são especiais.** A Sprint 1 (Setup) e Sprint 2 (Segurança) são de fundação — não têm features visíveis para o PO. Seus SPRINT-REVIEW.md são focados em validação técnica, não em demonstração de produto.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 14/07/2026 | Criação inicial: geração de 3 artefatos de sprint (SPRINT-CARD, SPRINT-TEST-SUITE, SPRINT-REVIEW) a partir de TASKS.md, TEST_PLAN.md e SPECS.md | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, breakdown-plan, test-strategy-design, acceptance-criteria, documentation-writer, ponytail.*
