# PROMPT-EXECUTE-1010-PRE-IMPLEMENTATION

## Contexto

Este prompt executa a **Fase de Pré-implementação** do pacote de desenvolvimento (extraída do `PROMPT-EXECUTE-SPRINT-TASKS.md`, Fase 0 — passos 0 a 7). Prepara o terreno do ciclo: valida a branch, carrega artefatos e documentos, identifica as tasks, valida a ordem de execução, descobre a stack e aciona as skills corretas, aplicando o checklist ponytail antes de qualquer código.

**Princípios fundamentais:**

1. **Stack-agnóstico:** a stack é descoberta dos documentos do projeto (nunca presumida).
2. **Nada de código nesta fase:** apenas preparação e decisões registradas no artefato de saída.
3. **Segurança primeiro:** branch `main`/`master` é erro crítico e aborta a execução.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|:---|:---|:---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_NAME}` | Nome/código do projeto de negócio | `PRJ-TEC-2026-0004-PROJETO-SHIELD` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{CICLO_DIR}` | Pasta do ciclo (onde os artefatos vivem) | `.../sprints/sprint-01-setup/` |
| `{CICLO_NUMBER}` | Número do ciclo | `1` |
| `{CICLO_NAME}` | Nome curto do ciclo (kebab-case) | `sprint-01-setup` |
| `{STACK}` | Stack tecnológica principal (opcional — se omitido, auto-detectada) | `Java 25 + Spring Boot + PostgreSQL` |
| `{TASK_IDS}` | IDs das tarefas a executar (opcional — vazio = todas pendentes) | `T-001,T-002` |

## Documentos de Referência

### Artefatos do Ciclo (obrigatórios — fonte direta da execução)

```
Ler obrigatoriamente antes de iniciar:
    ├── {CICLO_DIR}/SPRINT-CARD.md        ← Backlog, Branch, estimativas, critérios DONE, riscos
    ├── {CICLO_DIR}/SPRINT-TEST-SUITE.md  ← Cenários de teste aplicáveis a este ciclo
    └── {CICLO_DIR}/SPRINT-REVIEW.md      ← O que demonstrar na review
```

### Documentos-Mestre (obrigatórios — baseline de verdade)

```
SPECS_DIR = {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/
    ├── PRD.md, SPECS.md, TASKS.md, TEST_PLAN.md, ARCHITECTURE.md
```

### Documentos Transversais (obrigatórios)

```
    ├── {SOLUTION_PATH}/.specs/security/SECURITY.md   (se existir)
    └── {SOLUTION_PATH}/README.md                     (build, execução, testes)
```

---

## Missão

Preparar a execução do ciclo `{CICLO_NUMBER} — {CICLO_NAME}`: validar branch, carregar artefatos/documentos, selecionar tasks, validar ordem, descobrir stack e acionar skills — registrando tudo no artefato `PACKAGE-DEVELOPMENT-PRE-IMPLEMENTATION.md`.

---

## Fluxo de Execução

0. **Validar branch de desenvolvimento** — lida do campo `**Branch:**` no header do `SPRINT-CARD.md`:
   - Branch inexistente no card → ERRO CRÍTICO. PARE: o artefato foi gerado sem a branch — reexecutar `PROMPT-GENERATE-SPRINT-ARTEFACTS` com `BRANCH_NAME`.
   - Branch = `main`/`master` → ERRO CRÍTICO. PARE: viola políticas de GitOps/GitSecOps.
   - `git checkout {BRANCH_NAME}` (criar `git checkout -b` se não existir); confirmar com `git branch --show-current`.
1. **Carregar artefatos do ciclo** — SPRINT-CARD.md, SPRINT-TEST-SUITE.md, SPRINT-REVIEW.md.
2. **Carregar documentos-mestre** — PRD.md, SPECS.md, TASKS.md, TEST_PLAN.md, ARCHITECTURE.md.
3. **Carregar documentos transversais** — SECURITY.md, README.md.
4. **Identificar tasks a executar:** `{TASK_IDS}` informado → só as listadas; vazio → todas as pendentes do SPRINT-CARD.
5. **Validar ordem de execução** — dependências do SPRINT-CARD.md e TASKS.md §3 (pré-requisitos primeiro).
6. **Descobrir a stack e acionar as skills corretas:**
   - **6.1** Identificar a stack (ordem): `{STACK}` → PRD.md → SPECS.md → ARCHITECTURE.md → TASKS.md → README.md (ex.: `mvn`, `go`, `npm`, `cargo`).
   - **6.2** Decompor em componentes: linguagens (Java, Go, Python, TypeScript, Rust, Kotlin, C#, PHP, Ruby, Dart), frameworks (Spring Boot, Quarkus, Micronaut, Fiber, Gin, Express, FastAPI, Next.js, Django, Rails, Laravel, Flutter, React, React Native, Angular, Vue), bancos (PostgreSQL, MySQL, MongoDB, Redis, Elasticsearch), infraestrutura (Docker, K8s, RabbitMQ, Kafka).
   - **6.3** Selecionar skills por componente via `001-skills-inventory` (tabela de referência do `PROMPT-EXECUTE-SPRINT-TASKS.md` passo 6.3 — NÃO exaustiva).
   - **6.4** Registrar a decisão (skills acionadas e por quê) — entra no artefato de saída.
   - ⚠️ Dúvida sobre skills → PERGUNTAR ao humano antes de prosseguir.
7. **Aplicar o checklist ponytail (7 rungs) ANTES de codificar** em cada task:
   | Rung | Pergunta |
   |:---:|:---|
   | 1 | Isso precisa existir? (YAGNI) |
   | 2 | Já existe no codebase? |
   | 3 | A stdlib da linguagem cobre? |
   | 4 | Dependência já declarada resolve? |
   | 5 | O padrão do projeto (ARCHITECTURE.md) define como fazer? |
   | 6 | Dá pra ser uma unidade simples? |
   | 7 | Só então: código mínimo que funciona |

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-PRE-IMPLEMENTATION.md`:

```markdown
# PACKAGE-DEVELOPMENT-PRE-IMPLEMENTATION.md — Pré-implementação: Ciclo {N}
[Header: solução, projeto, ciclo, data]
## 1. Branch Validada
[Branch do SPRINT-CARD + confirmação git branch --show-current]
## 2. Artefatos e Documentos Carregados
[Tabela: artefato | carregado? | observações]
## 3. Tasks Selecionadas
[Tabela: ID | Tarefa | Prio | Dependências | Status]
## 4. Ordem de Execução Validada
[Lista ordenada com justificativa]
## 5. Stack Detectada
[Stack + fonte da detecção ({STACK} | PRD | SPECS | README...)]
## 6. Skills Acionadas
[Tabela: componente | skill | justificativa]
## 7. Checklist Ponytail (7 rungs)
[Registro por task: rungs aplicados e decisões]
## Rodapé
[Indicação de geração por IA, data/hora]
```

---

## Skills

### Dinâmico — seleção da stack (passo 6)

| Skill | Uso |
|:---|:---|
| `001-skills-inventory` | Listar skills disponíveis (ponto de partida da seleção) |

**Protocolo:** 1. invocar `001-skills-inventory` → 2. cruzar componentes da stack com as skills → 3. priorizar skills específicas sobre genéricas → 4. ambiguidade = PERGUNTAR ao humano → 5. ausência = conhecimento geral + documentar → 6. registrar a seleção no artefato (§6).

**Tabela de referência stack→skills** (herdada do original passo 6.3 — NÃO exaustiva):

| Componente | Skills |
|:---|:---|
| Java | `110-java-maven-best-practices`, `121-java-object-oriented-design`, `126-java-exception-handling`, `130-java-testing-strategies`, `141-java-refactoring-with-modern-features` |
| Go | `golang-pro`, `golang-performance`, `golang-testing`, `golang-error-handling` |
| TypeScript | `typescript-pro`, `javascript-typescript-jest` |
| Python | `python-pro`, `pytest-coverage` |
| Rust | `rust-engineer` |
| Dart/Flutter | `flutter-expert` |
| Spring Boot / Quarkus / Micronaut | `301/302/304/311-frameworks-spring-boot-*` · `401/402-frameworks-quarkus-*` · `501/502-frameworks-micronaut-*` |
| React / React Native / Angular / Vue | `react-best-practices`, `react-patterns` · `react-native-expert` · `angular` · `vue-expert` |
| Bancos | `postgres-pro`, `postgresql-optimization`, `mongodb-query-optimizer`, `mongodb-schema-design` |
| Testes (qualquer stack) | `131-java-testing-unit-testing`, `golang-testing`, `javascript-typescript-jest`, `pytest-coverage` |
| Segurança (qualquer stack) | `security-review`, `124-java-secure-coding`, `gdpr-compliant` |

### Transversais

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| `ponytail` | full | Checklist de 7 rungs antes de codificar |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

---

## Regras de Ouro

1. Branch `main`/`master` = aborto imediato (erro crítico).
2. NUNCA presumir stack — sempre do parâmetro ou dos documentos.
3. NENHUMA linha de código nesta fase — apenas preparação e decisões.
4. Dúvida sobre skills → pergunta ao humano (nunca decidir sozinho).
5. Artefato de saída SEMPRE em `{CICLO_DIR}/`, com `[STATUS: Em análise]` inicial.
