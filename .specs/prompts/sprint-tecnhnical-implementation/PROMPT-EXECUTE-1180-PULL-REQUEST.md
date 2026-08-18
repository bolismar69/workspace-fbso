# PROMPT-EXECUTE-1180-PULL-REQUEST

## Contexto

Este prompt executa a **fase de Pull Request** do pacote de desenvolvimento — o fecho do ciclo: atualiza o repositório local/remoto e abre a PR com o `TASK-EXECUTED` anexado. **Delega a execução** aos prompts atuais `PROMPT-GENERATE-PULL-REQUEST` (abertura da PR) e `PROMPT-02-ATUALIZAR-REPOSITORIO-E-ABRIR-PULL-REQUEST` (protocolo git completo, mantido como executor do fluxo de repositório).

**Princípios fundamentais:**

1. **PR é o fecho do ciclo:** só acontece com todas as fases anteriores COMPLIANCE.
2. **Stage seletivo em monorepo:** somente arquivos do escopo do ciclo entram na PR (padrão `INCLUDE_PATHS`/`EXCLUDE_PATHS` do PROMPT-02).
3. **Evidência anexada:** o `TASK-EXECUTED-*.md` (Fase 1140) é o corpo da PR.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|:---|:---|:---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_NAME}` | Nome/código do projeto de negócio | `PRJ-TEC-2026-0004-PROJETO-SHIELD` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{CICLO_DIR}` | Pasta do ciclo | `.../sprints/sprint-01-setup/` |
| `{CICLO_NUMBER}` | Número do ciclo | `1` |
| `{CICLO_NAME}` | Nome curto do ciclo (kebab-case) | `sprint-01-setup` |
| `{REPO_ROOT}` | Raiz do repositório Git (monorepo) | `/home/user/work/workspace-fbso` |
| `{IS_MONOREPO}` | `true`/`false` | `true` |
| `{BASE_BRANCH}` | Branch destino da PR | `main` |
| `{INCLUDE_PATHS}` / `{EXCLUDE_PATHS}` | Globs de stage (monorepo) | `backend/...` / outros projetos |
| `{GITHUB_ORG_REPO}` | Org/repo no GitHub | `bolismar69/workspace-fbso` |

## Documentos de Referência

```
Ler obrigatoriamente:
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-TASK-EXECUTED-REPORT.md ← Caminho do TASK-EXECUTED (Fase 1140)
    └── {CICLO_DIR}/PACKAGE-DEVELOPMENT-EXECUTION-REPORT.md     ← Relatório do ciclo (Fase 1100)
```

> ⚠️ Se `PACKAGE-DEVELOPMENT-TASK-EXECUTED-REPORT.md` não existir → **PARE**: execute primeiro o `PROMPT-EXECUTE-1160-TASK-EXECUTED-REPORT`.

---

## Missão

Fechar o ciclo `{CICLO_NUMBER} — {CICLO_NAME}` com a abertura da Pull Request: stage seletivo, commit convencional, push e `gh pr create` com o TASK-EXECUTED anexado — registrando tudo no `PACKAGE-DEVELOPMENT-PULL-REQUEST.md`.

---

## Fluxo de Execução

1. **Auditoria do working tree** — protocolo do `PROMPT-02` (Passo 1): `git status --short`, `git branch --show-current`, `git remote -v`; nunca em `{BASE_BRANCH}` diretamente (criar `feature/...`).
2. **Stage seletivo** (se `{IS_MONOREPO}` = `true`): classificar arquivos por `{INCLUDE_PATHS}` / `{EXCLUDE_PATHS}`.
3. **Commit convencional** — padrão `feat(PRJ/TASK/ISSUE-XXXX): <mensagem>`.
4. **Push** com tracking upstream.
5. **Abrir PR** via `PROMPT-GENERATE-PULL-REQUEST` (executor mantido): `gh pr create` para `{BASE_BRANCH}` com `--body-file` = TASK-EXECUTED (caminho do artefato 1140).
6. **Registrar a PR** na pasta `{SOLUTION_PATH}/.specs/pull-requests/` (padrão `PR_N-{tipo}-{escopo}__{feature}.md` do executor) e no artefato da fase.

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-PULL-REQUEST.md`:

```markdown
# PACKAGE-DEVELOPMENT-PULL-REQUEST.md — Pull Request: Ciclo {N}
[Header: solução, projeto, ciclo, data]
## 1. PR Aberta
- URL: [url real da gh CLI]
- Branch: feature/... → {BASE_BRANCH}
- Corpo: TASK-EXECUTED-*.md anexado
## 2. Arquivos Incluídos
[Tabela: arquivo | tipo (🆕/🔄) | grupo]
## 3. Arquivos Excluídos (monorepo)
[Arquivos de outros projetos que ficaram de fora]
## 4. Registro
[.specs/pull-requests/PR_N-...md criado]
## Rodapé
[Indicação de geração por IA, data/hora]
```

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| `verification-before-completion` | automático | Conferir PR aberta e registro antes de concluir |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

> `PROMPT-GENERATE-PULL-REQUEST` e `PROMPT-02-ATUALIZAR-REPOSITORIO-E-ABRIR-PULL-REQUEST` são **prompts executores**, não skills.

---

## Regras de Ouro

1. PR somente com todas as fases anteriores COMPLIANCE.
2. Nunca commitar em `main`/`master` — sempre branch `feature/...`.
3. Stage seletivo obrigatório em monorepo (INCLUDE/EXCLUDE).
4. TASK-EXECUTED (1140) é o corpo obrigatório da PR.
5. Registro da PR salvo em `.specs/pull-requests/` (padrão do executor).
