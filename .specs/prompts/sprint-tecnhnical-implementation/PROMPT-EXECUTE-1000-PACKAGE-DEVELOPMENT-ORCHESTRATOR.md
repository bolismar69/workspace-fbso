# PROMPT-EXECUTE-1000-PACKAGE-DEVELOPMENT-ORCHESTRATOR

## Contexto

Este prompt **orquestra o ciclo completo de desenvolvimento de um Pacote (PACKAGE)** — a sequência de 18 fases específicas (`1010` a `1180`) que decompõem a execução de um ciclo de entrega. Ele não executa os passos das fases: **delega cada fase ao prompt específico correspondente** e controla a sequência, os gates e o loop de retorno.

> **Relação com o legado:** o `PROMPT-EXECUTE-SPRINT-TASKS.md` permanece intacto como executor histórico monolítico. Este orquestrador + os 18 prompts específicos são a via estruturada, fase a fase, com artefato próprio por fase (`PACKAGE-DEVELOPMENT-{FASE}.md` em `{CICLO_DIR}`).

**Princípios fundamentais:**

1. **Sequência obrigatória:** cada fase só inicia quando a anterior gerou seu artefato em `[STATUS: COMPLIANCE]` e o humano validou.
2. **Delegação total:** o orquestrador NUNCA executa conteúdo de fase — invoca o prompt específico e aguarda o resultado.
3. **Loop de retorno controlado:** o 1120 (Code Review) pode retornar ao 1070 (Test Planning) no máximo **2 ciclos** (regra herdada da Fase 7 original).
4. **Falha persistente = impedimento:** após 3 tentativas na mesma falha, o 1110 registra `IMPEDIMENT-SPRINT-{N}.md` e pausa para decisão humana.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|:---|:---|:---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_NAME}` | Nome/código do projeto de negócio | `PRJ-TEC-2026-0004-PROJETO-SHIELD` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{CICLO_DIR}` | Pasta do ciclo (sprint) — onde os artefatos do ciclo vivem | `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/sprints/sprint-01-setup/` |
| `{CICLO_NUMBER}` | Número do ciclo (1 a N) | `1` |
| `{CICLO_NAME}` | Nome curto do ciclo (kebab-case) | `sprint-01-setup` |
| `{STACK}` | Stack tecnológica principal (opcional — auto-detectada no 1020) | `Java 25 + Spring Boot + PostgreSQL` |
| `{TASK_IDS}` | IDs das tarefas a executar (opcional — vazio = todas pendentes) | `T-001,T-002` |

## Documentos de Referência

```
Ler obrigatoriamente (contexto do orquestrador):
    ├── {CICLO_DIR}/SPRINT-CARD.md        ← Backlog, Branch, critérios DONE, métricas
    ├── {CICLO_DIR}/SPRINT-TEST-SUITE.md  ← Cenários de teste aplicáveis ao ciclo
    ├── {CICLO_DIR}/SPRINT-REVIEW.md      ← O que demonstrar na review
    └── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/
        PRD.md, SPECS.md, TASKS.md, TEST_PLAN.md, ARCHITECTURE.md   ← docs-mestre (fonte da verdade)
```

---

## Missão

Conduzir o ciclo `{CICLO_NUMBER} — {CICLO_NAME}` pelas 18 fases do pacote de desenvolvimento, invocando o prompt específico de cada fase, validando cada artefato `PACKAGE-DEVELOPMENT-*.md` e encerrando com o relatório de execução e a atualização dos artefatos.

---

## Fluxo de Execução

### Fase de Orquestração 0 — Preparação

0. **Validar parâmetros** — os 6 obrigatórios informados; `{STACK}` e `{TASK_IDS}` opcionais.
1. **Ler SPRINT-CARD.md** e confirmar a branch do ciclo (veto: main/master — erro crítico herdado do passo 0 original).
1a. **(OPCIONAL — pré-ciclo) Invocar `PROMPT-EXECUTE-1010-TECHNICAL-DEBT-AUDIT`** (modo `pre-ciclo`) quando o ciclo exigir auditoria prévia de débito — decisão humana sobre tratar agora vs. depois.
2. **Apresentar o plano de execução** ao humano: sequência das fases (com as condicionais 3a/3b/4a/6a), tasks do ciclo e artefatos esperados. Aguardar confirmação.

### Sequência de Delegação (18 fases)

| Ordem | Prompt específico | Artefato gerado em `{CICLO_DIR}/` | Gate |
|:---:|:---|:---|:---:|
| 0 | `PROMPT-EXECUTE-1010-TECHNICAL-DEBT-AUDIT` (opcional — pré-ciclo) | `PACKAGE-DEVELOPMENT-TECHNICAL-DEBT-AUDIT.md` | Decisão humana |
| 1 | `PROMPT-EXECUTE-1020-PRE-IMPLEMENTATION` | `PACKAGE-DEVELOPMENT-PRE-IMPLEMENTATION.md` | Humano + status |
| 2 | `PROMPT-EXECUTE-1030-DEVELOPMENT-PLANNING` | `PACKAGE-DEVELOPMENT-PLANNING.md` | Humano + status |
| 3 | `PROMPT-EXECUTE-1040-IMPLEMENTATION` | `PACKAGE-DEVELOPMENT-IMPLEMENTATION.md` | Build/testes verdes |
| 3a | `PROMPT-EXECUTE-1050-CVE-SCA-SCAN` (condicional — se 480/510 exigirem) | `PACKAGE-DEVELOPMENT-CVE-SCA-SCAN.md` | Veredito verde / DT-XXX |
| 3b | `PROMPT-EXECUTE-1060-CI-CD-PIPELINE` (condicional — se 500/087 exigirem) | `PACKAGE-DEVELOPMENT-CI-CD-PIPELINE.md` | Gate 1130 + humano |
| 4 | `PROMPT-EXECUTE-1070-TEST-PLANNING` | `PACKAGE-DEVELOPMENT-TEST-PLANNING.md` | Humano + status |
| 4a | `PROMPT-EXECUTE-1080-STRESS-PERFORMANCE-TEST` (condicional — se 510/050 tiverem cenários) | `PACKAGE-DEVELOPMENT-STRESS-PERFORMANCE-TEST.md` | PASS / FAIL + decisão |
| 5 | `PROMPT-EXECUTE-1090-TEST-IMPLEMENTATION` | `PACKAGE-DEVELOPMENT-TEST-IMPLEMENTATION.md` | Testes verdes, cobertura ≥80% |
| 6 | `PROMPT-EXECUTE-1100-QUALITY-VALIDATION` | `PACKAGE-DEVELOPMENT-QUALITY-VALIDATION.md` | Zero warnings/violations |
| 6a | `PROMPT-EXECUTE-1110-FAILURE-HANDLING` (condicional — só se houve falhas) | `PACKAGE-DEVELOPMENT-FAILURE-HANDLING.md` | Decisão humana |
| 7 | `PROMPT-EXECUTE-1120-CODE-REVIEW` | `PACKAGE-DEVELOPMENT-CODE-REVIEW.md` | Humano; loop máx. 2 |
| 7a | `PROMPT-EXECUTE-1130-QUALITY-ASSURANCE-REVIEW` (HITL — máx. 3 loops) | `PACKAGE-DEVELOPMENT-QUALITY-ASSURANCE-REVIEW.md` | APPROVED |
| 8 | `PROMPT-EXECUTE-1140-POST-IMPLEMENTATION` | `PACKAGE-DEVELOPMENT-POST-IMPLEMENTATION.md` | Checklist concluído |
| 9 | `PROMPT-EXECUTE-1150-EXECUTION-REPORT` | `PACKAGE-DEVELOPMENT-EXECUTION-REPORT.md` | Humano + status |
| 9a | `PROMPT-EXECUTE-1160-TASK-EXECUTED-REPORT` | `PACKAGE-DEVELOPMENT-TASK-EXECUTED-REPORT.md` | Arquivo gerado |
| 10 | `PROMPT-EXECUTE-1170-ARTIFACTS-UPDATE` | `PACKAGE-DEVELOPMENT-ARTIFACTS-UPDATE.md` | Consistência cruzada |
| 10a | `PROMPT-EXECUTE-1180-PULL-REQUEST` (fecho do ciclo) | `PACKAGE-DEVELOPMENT-PULL-REQUEST.md` | PR aberta |

### Regras do Loop de Retorno (herdadas da Fase 7 original)

- Se o **1120** gerar ajustes de código → retornar ao **1070** e reexecutar 1070 → 1090 → 1100 → 1120.
- **Máximo de 2 ciclos completos** (1070→1090→1100→1120). Após o 2º ciclo, achados Critical/High não resolvidos são registrados no relatório (§8 do 1150) e o fluxo prossegue; Medium/Low viram débito técnico documentado.
- Falha que persiste por 3 tentativas em qualquer fase → **1110** registra o impedimento e o orquestrador **PÁRA** para decisão humana.
- **Gate 1130:** após o 1120, o `1130-QUALITY-ASSURANCE-REVIEW` (HITL, 3 loops) é obrigatório — FAILED devolve a correção ao 1040/1120 e re-executa o gate.
- **Débitos centralizados:** todo `DT-XXX` gerado em 1110/1120/1050/1080 é incorporado pelo `1010-TECHNICAL-DEBT-AUDIT` (modo `catalogo`).

### Encerramento

Ao final do 1180 (fecho com PR), apresentar o resumo consolidado: tasks concluídas, artefatos gerados (18 arquivos `PACKAGE-DEVELOPMENT-*.md` + TASK-EXECUTED + PR), impedimentos e próximo passo (pacote 595 / Bloco F conforme o roadmap TECHLEAD).

---

## Skills

| Skill | Modo | Uso no orquestrador |
|:---|:---|:---|
| `verification-before-completion` | automático | Verificar que o artefato da fase está COMPLIANCE antes de avançar |
| `caveman` | full | Comunicação interativa com o humano (nunca em artefatos permanentes) |

> As skills específicas de cada fase (stack, testes, auditorias) são acionadas pelos prompts 1010–1180 — o orquestrador não aciona skills de stack.

---

## Regras de Ouro

1. NUNCA executar conteúdo de fase — somente delegar ao prompt específico correspondente.
2. Fase N só inicia com o artefato da fase N-1 em `[STATUS: COMPLIANCE]` + validação humana.
3. Branch `main`/`master` = erro crítico — abortar (herdado do passo 0).
4. Loop de retorno limitado a 2 ciclos; falhas persistentes viram impedimento (1110) com decisão humana.
5. Todos os artefatos ficam em `{CICLO_DIR}/` — nunca em diretórios globais.
6. O orquestrador não substitui nem altera o `PROMPT-EXECUTE-SPRINT-TASKS.md` (legado intacto).
