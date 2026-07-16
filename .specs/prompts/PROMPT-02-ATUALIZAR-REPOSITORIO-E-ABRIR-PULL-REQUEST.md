# Contexto:
  - Este prompt é **genérico** — adapta-se a qualquer demanda: projeto, feature, issue, incidente, PoC, hotfix, etc.
  - Antes de executar, o humano deve fornecer os parâmetros da seção `# ⚙️ Parâmetros de Entrada`.
  - ⚠️ **Pré-requisito:** O Prompt #1 (`PROMPT-01`) já deve ter sido executado e o arquivo TASK-EXECUTED deve estar disponível em `{SKILL_OUTPUT_DIR}/`.
  - O repositório pode ser **monorepo** (múltiplos projetos independentes) ou **single-repo**. O protocolo se adapta a ambos.

---

# ⚙️ Parâmetros de Entrada (preencher antes de executar)

> **Instrução:** No momento de invocar este prompt, o humano deve informar os valores abaixo. Parâmetros marcados com `[*]` são obrigatórios.

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{DEMAND_ID}` `[*]` | Identificador único da demanda | `PRJ-FIN-2026-0001`, `ISSUE-422`, `INC-2026-0091`, `POC-ML-001` |
| `{DEMAND_TYPE}` `[*]` | Tipo da demanda | `projeto`, `feature`, `issue`, `incidente`, `hotfix`, `poc` |
| `{DEMAND_TITLE}` `[*]` | Título descritivo curto | `Reforma Tributária 2026 Corporativo` |
| `{DEMAND_SLUG}` `[*]` | Slug em kebab-case para nome de branch e arquivos | `reforma-tributaria-2026-fases-0-1-2` |
| `{SOLUTION_ROOT}` `[*]` | Caminho absoluto da raiz da solução sistêmica | `/home/user/work/ms-billing-engine-tax-rates` |
| `{REPO_ROOT}` `[*]` | Caminho absoluto da raiz do repositório Git (pode ser igual ou ancestral de `{SOLUTION_ROOT}`) | `/home/user/work/workspace-fbso` |
| `{IS_MONOREPO}` `[*]` | `true` se o repositório contém múltiplos projetos independentes, `false` caso contrário | `true` |
| `{BASE_BRANCH}` `[*]` | Branch destino da PR | `main`, `master`, `develop` |
| `{GIT_REMOTE_NAME}` | Nome do remote Git | `origin` |
| `{GITHUB_ORG_REPO}` `[*]` | Organização/repositório no GitHub | `bolismar69/workspace-fbso` |
| `{SPECS_DIR}` | Caminho relativo da pasta de especificações (a partir de `{SOLUTION_ROOT}`) | `.specs` |
| `{SKILL_OUTPUT_DIR}` | Pasta onde o TASK-EXECUTED foi salvo (mesmo valor do Prompt #1) | `{SPECS_DIR}/skill-output` |
| `{PULL_REQUESTS_DIR}` | Pasta onde o registro histórico da PR será salvo | `{SPECS_DIR}/pull-requests` |
| `{INDEX_FILE}` | Caminho para o INDEX.md das specs (se não existir, será criado) | `{SPECS_DIR}/INDEX.md` |
| `{INCLUDE_PATHS}` `[*]` | Lista de paths/globs a incluir no stage (um por linha). Use `*` para diretórios inteiros | Ver abaixo |
| `{EXCLUDE_PATHS}` | Lista de paths/globs a **NUNCA** incluir — outros projetos do monorepo, arquivos pessoais, etc. | Ver abaixo |
| `{COMMIT_SCOPE}` `[*]` | Escopo do Conventional Commit | `PRJ-FIN-2026-0001` |

### `{INCLUDE_PATHS}` — Exemplo (monorepo)

```
backend/go/fiber/microservices/ms-billing-engine-tax-rates/
backend/go/libs/go-native/taxnexus-billing-core-lib/
.specs/
business-inputs/business-projects/PRJ-FIN-2026-0001-*/
```

### `{INCLUDE_PATHS}` — Exemplo (single-repo)

```
.
```

### `{EXCLUDE_PATHS}` — Exemplo (itens de outros projetos no monorepo)

```
backend/go/fiber/microservices/ms-tax-individual-income/
backend/go/fiber/microservices/ms-tax-nexus-taas-calc-engine/
backend/go/libs/go-native/taxnexus-individual-core-lib/
backend/java/
*.md
```
> **Nota:** `*.md` exclui arquivos como `AGENTS.md`, `CLAUDE.md`, `GEMINI.md` que são de configuração pessoal e não do projeto. Ajuste conforme necessário.

---

# Missão:
> Atualizar o repositório local e remoto com as implementações concluídas para a demanda `{DEMAND_ID}` e abrir uma Pull Request no GitHub anexando o documento de execução (TASK-EXECUTED) como corpo da descrição.

---

# 🔁 Protocolo de Atualização de Repositório e Abertura de PR (Passo a Passo)

## Passo 1 — Auditoria do Working Tree
```bash
cd {REPO_ROOT}
git status --short
git branch --show-current
git remote -v
```

1. Liste TODOS os arquivos alterados/novos/deletados.
2. **⚠️ Se estiver em `{BASE_BRANCH}`**: prossiga para criar uma feature branch (Passo 3).
3. **⚠️ Se já estiver em feature branch**: avalie se os arquivos atuais pertencem ao mesmo escopo de `{DEMAND_ID}` ou se é necessário um novo branch.

## Passo 2 — Classificação dos Arquivos (crítico se `{IS_MONOREPO}` = `true`)
Separe os arquivos do `git status` em duas listas:

### ✅ Arquivos da Demanda `{DEMAND_ID}` (entrarão na PR)
- Correspondem a algum dos padrões em `{INCLUDE_PATHS}`.
- Estão no escopo definido por `{ARCHITECTURE_FILE}` (se existir).

### ❌ Arquivos de OUTRAS demandas/projetos (ficarão de fora)
- Correspondem a algum dos padrões em `{EXCLUDE_PATHS}`.
- Pertencem a outros projetos do monorepo.

> **Registre ambas as listas** — os paths excluídos serão documentados no arquivo de registro da PR para rastreabilidade.

## Passo 3 — Criação da Feature Branch
1. Nome do branch:
   ```
   feature/{DEMAND_SLUG}
   ```
   > Ajuste o prefixo conforme o tipo: `feature/` para projeto/feature, `fix/` para issue/hotfix, `poc/` para prova de conceito.

2. Crie o branch:
   ```bash
   git checkout -b feature/{DEMAND_SLUG}
   ```

## Passo 4 — Stage Seletivo
**⚠️ Se `{IS_MONOREPO}` = `true`:** Use `git add` com **paths específicos**. **NUNCA** use `git add .` ou `git add -A`!

**⚠️ Se `{IS_MONOREPO}` = `false`:** Pode usar `git add .` se o repositório inteiro pertence à mesma demanda. Mas ainda assim prefira paths específicos para auditoria.

```bash
# Para cada path em {INCLUDE_PATHS}:
git add [path-1]
git add [path-2]
git add [path-3]
# ...
```

Validação pós-stage:
```bash
# O que entrou (deve corresponder a {INCLUDE_PATHS}):
git diff --cached --stat

# O que ficou de fora (deve corresponder a {EXCLUDE_PATHS}):
git status --short | grep -v "^M \|^D \|^A \|^R "
```

## Passo 5 — Commit
Siga **Conventional Commits**:

```
{type}({COMMIT_SCOPE}): {DEMAND_TITLE} — [resumo em uma linha]

[corpo com bullet points dos principais itens implementados]

Co-Authored-By: Claude <noreply@anthropic.com>
```

Onde `{type}` é mapeado de `{DEMAND_TYPE}`:

| `{DEMAND_TYPE}` | `{type}` (Conventional Commit) |
|---|---|
| `projeto`, `feature` | `feat` |
| `issue`, `hotfix` | `fix` |
| `incidente` | `fix` |
| `poc` | `feat` |

## Passo 6 — Push
```bash
git push -u {GIT_REMOTE_NAME} feature/{DEMAND_SLUG}
```

Confirme que o remote reportou o novo branch e sugeriu o link para criar a PR.

## Passo 7 — Criação da Pull Request
1. Verifique autenticação:
   ```bash
   gh auth status
   ```

2. Localize o arquivo TASK-EXECUTED mais recente para esta demanda:
   ```bash
   ls -t {SKILL_OUTPUT_DIR}/*.md | head -1
   ```
   > Alternativa: se souber o nome exato, use-o diretamente:
   > `{SKILL_OUTPUT_DIR}/AAAA-MM-DD-HHMMSS_{DEMAND_SLUG}.md`

3. Crie a PR usando o sumário como corpo da descrição:
   ```bash
   gh pr create \
     --base {BASE_BRANCH} \
     --head feature/{DEMAND_SLUG} \
     --title "{type}({COMMIT_SCOPE}): {DEMAND_TITLE}" \
     --body-file "{SKILL_OUTPUT_DIR}/[nome-do-arquivo-task-executed].md"
   ```

   > **Por que `--body-file`?** O GitHub renderiza Markdown nativamente no corpo da PR. Isso preserva tabelas, checkboxes, blocos de código e emojis do documento de sumário sem poluir o commit message.

4. Anote o número e URL da PR retornados (ex: `https://github.com/{GITHUB_ORG_REPO}/pull/6`).

## Passo 8 — Registro Histórico da PR (para consulta humana e IA)
1. Crie o diretório se não existir:
   ```bash
   mkdir -p {PULL_REQUESTS_DIR}
   ```

2. Crie o arquivo de registro seguindo o padrão:
   ```
   PR_[NUMERO]__{DEMAND_SLUG}.md
   ```
   Exemplo: `PR_6__reforma-tributaria-2026-fases-0-1-2.md`

3. Conteúdo mínimo do registro:

```markdown
# 📌 Pull Request #[N] — {DEMAND_TITLE}

* **URL:** https://github.com/{GITHUB_ORG_REPO}/pull/[N]
* **Branch:** `feature/{DEMAND_SLUG}` → `{BASE_BRANCH}`
* **Data de criação:** AAAA-MM-DD
* **Repositório:** `{GITHUB_ORG_REPO}` ({monorepo|single-repo})
* **Status:** 🟢 Aberta
* **Demanda:** {DEMAND_ID} — {DEMAND_TITLE} ({DEMAND_TYPE})
* **Sumário de implementação:** [link relativo para o TASK-EXECUTED.md]

---

## 📊 Estatísticas
| Métrica | Valor |
|---|---|
| Arquivos alterados | [N] |
| Linhas adicionadas | +[N] |
| Linhas removidas | −[N] |

---

## 🛠️ Ações Realizadas

### 1. Análise pré-PR (AAAA-MM-DD)
| Aspecto | Estado |
|---|---|
| Branch anterior | `{BASE_BRANCH}` |
| Remote verificado | `git@github.com:{GITHUB_ORG_REPO}.git` ✅ |
| IS_MONOREPO | {true|false} |
| Arquivos no working tree | [N] total |

### 2. Criação da Feature Branch
```bash
git checkout -b feature/{DEMAND_SLUG}
```

### 3. Stage Seletivo
#### ✅ Arquivos incluídos ([N])
| Path | Tipo | Descrição |
|---|---|---|

#### ❌ Arquivos excluídos ([N])
| Path | Motivo |
|---|---|

### 4. Commit
| Campo | Valor |
|---|---|
| Tipo | {type} |
| Escopo | {COMMIT_SCOPE} |
| Hash | [hash] |
| Co-autoria | `Co-Authored-By: Claude <noreply@anthropic.com>` |

### 5. Push
```bash
git push -u {GIT_REMOTE_NAME} feature/{DEMAND_SLUG}
```

### 6. Criação da PR
```bash
gh pr create --base {BASE_BRANCH} --head feature/{DEMAND_SLUG} --title "..." --body-file "..."
```

✅ **PR #[N] criada**: https://github.com/{GITHUB_ORG_REPO}/pull/[N]

---

## 📦 Detalhamento por Item
[copiado do TASK-EXECUTED, organizado por fase/GAP/issue/patch]

## 🔒 Segurança
[checklist do SECURITY.md]

### ⚠️ Vulnerabilidades pendentes
[se houver — documentar para correção futura]

## 📄 Documentos Vinculados
| Documento | Localização |
|---|---|

## 🧪 Evidências de Testes
[output dos testes]

---

## 🔗 Links

- **Pull Request:** https://github.com/{GITHUB_ORG_REPO}/pull/[N]
- **Branch:** `feature/{DEMAND_SLUG}`
- **Commit:** `[hash]`

---
🤖 *Registro gerado em AAAA-MM-DD. Histórico completo da criação da PR para consulta humana e IA.*
```

## Passo 9 — Atualização do Índice de Specs
1. Se `{INDEX_FILE}` existir, adicione na seção `## Histórico de Pull Requests` (crie a seção se necessário):
   ```markdown
   - **PR #[N] — {DEMAND_TITLE}:** Consulte [pull-requests/PR_[N]__{DEMAND_SLUG}.md](pull-requests/PR_[N]__{DEMAND_SLUG}.md) — [N] arquivos, +[N] inserções
   ```

2. Se `{INDEX_FILE}` não existir, crie-o com a estrutura básica:
   ```markdown
   # Mapa de Especificações — [nome da solução]

   ## Histórico de Pull Requests
   - **PR #[N] — {DEMAND_TITLE}:** Consulte [pull-requests/PR_[N]__{DEMAND_SLUG}.md](pull-requests/PR_[N]__{DEMAND_SLUG}.md)
   ```

## Passo 10 — Resumo Final
Emita uma mensagem de confirmação com:
- ✅ Link da PR no GitHub.
- 📊 Estatísticas (arquivos, linhas, itens, testes).
- 📄 Localização do registro de PR criado (`{PULL_REQUESTS_DIR}/`).
- 📄 Localização do índice atualizado (`{INDEX_FILE}`).
- ⚠️ Lista de arquivos que ficaram de fora (se `{IS_MONOREPO}` = `true`).
- 🔒 Vulnerabilidades pendentes (se houver).

---

# 🚫 Anti-Padrões (Coisas a NUNCA fazer)

| ❌ Anti-Padrão | ✅ Correto |
|---|---|
| `git add .` ou `git add -A` em monorepo | `git add [paths específicos de {INCLUDE_PATHS}]` |
| Commitar direto em `{BASE_BRANCH}` | Criar feature branch primeiro |
| Mensagem de commit `"fix"` ou `"updates"` | Conventional Commits com `{COMMIT_SCOPE}` |
| PR sem corpo ou com "—" | `--body-file` com o TASK-EXECUTED completo |
| Misturar alterações de múltiplas demandas na mesma PR | Uma PR por `{DEMAND_ID}` |
| Não documentar o que ficou de fora | Listar exclusões no arquivo de registro da PR |
| Pular o registro histórico | Criar `PR_[N]__{DEMAND_SLUG}.md` em `{PULL_REQUESTS_DIR}/` |
| Esquecer de atualizar o índice | Nova entrada em `{INDEX_FILE}` |

---

# 📋 Exemplo de Invocação

> **Humano:** "Executar o Prompt #2 com os parâmetros:"
>
> - `{DEMAND_ID}` = `PRJ-FIN-2026-0001`
> - `{DEMAND_TYPE}` = `projeto`
> - `{DEMAND_TITLE}` = `Reforma Tributária 2026 — Fases 0, 1, 2`
> - `{DEMAND_SLUG}` = `reforma-tributaria-2026-fases-0-1-2`
> - `{SOLUTION_ROOT}` = `/home/user/work/ms-billing-engine-tax-rates`
> - `{REPO_ROOT}` = `/home/user/work/workspace-fbso`
> - `{IS_MONOREPO}` = `true`
> - `{BASE_BRANCH}` = `main`
> - `{GIT_REMOTE_NAME}` = `origin`
> - `{GITHUB_ORG_REPO}` = `bolismar69/workspace-fbso`
> - `{SPECS_DIR}` = `.specs`
> - `{SKILL_OUTPUT_DIR}` = `.specs/skill-output`
> - `{PULL_REQUESTS_DIR}` = `.specs/pull-requests`
> - `{INDEX_FILE}` = `.specs/INDEX.md`
> - `{COMMIT_SCOPE}` = `PRJ-FIN-2026-0001`
> - `{INCLUDE_PATHS}`:
>   ```
>   backend/go/fiber/microservices/ms-billing-engine-tax-rates/
>   backend/go/libs/go-native/taxnexus-billing-core-lib/
>   .specs/
>   business-inputs/business-projects/PRJ-FIN-2026-0001-*/
>   ```
> - `{EXCLUDE_PATHS}`:
>   ```
>   backend/go/fiber/microservices/ms-tax-individual-income/
>   backend/go/fiber/microservices/ms-tax-nexus-taas-calc-engine/
>   backend/go/libs/go-native/taxnexus-individual-core-lib/
>   backend/java/
>   README-FRAMEWORKS-MODELO.md
>   ```

---

# Skills Recomendadas
`git-commit`, `code-documenter`, `fullstack-guardian`

---

# 🔗 Dependência
> ⚠️ Este protocolo assume que o **Prompt #1** (`PROMPT-01-PROCESSAR-TASKS-E-GERAR-DOCUMENTO-DE-EXECUCAO.md`) já foi executado e o arquivo TASK-EXECUTED está disponível em `{SKILL_OUTPUT_DIR}/`.
