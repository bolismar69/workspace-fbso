# Design Document — Estratégia de Branching por Sprint

- **Projeto:** `ms-fbso-platform-admin` / `PRJ-FIN-2026-0003-SAAS-FBSO-ORG`
- **Data:** 16 de Julho de 2026
- **Status:** Aprovado — em implementação
- **Decisão:** Estratégia B — Branch por Sprint

---

## 1. Problema

O projeto adotou inicialmente uma **branch única** (`feature/java-fbso-platform-admin`) para todas as 7 sprints (14 semanas, 80 tarefas). Essa branch foi mergeada e deletada ao final da Sprint 2. Duas dores foram identificadas:

1. **Desalinhamento operacional:** commits de outras branches precisaram ser movidos manualmente para a branch correta, gerando retrabalho
2. **Risco de longevidade:** manter uma única branch viva por 14 semanas cria risco real de perda ou sobrescrita por descuido (ex.: `git reset --hard` errado, `push --force` acidental)

## 2. Alternativas Avaliadas

### A — Branch Única (descartada)
Uma branch para todas as sprints, merge único no final.
- **Descartada porque:** risco de perda proporcional ao tempo de vida; entrega não-incremental; conflitos com outras branches.

### B — Branch por Sprint (selecionada)
Uma branch por sprint (~2 semanas), merge no `main` ao final de cada sprint.
- **Selecionada porque:** reduz risco (branch vive 2 semanas); entrega incremental; isola falhas por sprint; resolve a dor de desalinhamento com nomes explícitos.

### C — Trunk-Based Development (descartada para agora)
Branches por tarefa (~1-2 dias), merge diário no `main`.
- **Descartada porque:** overkill para 1 desenvolvedor; exige feature flags e CI/CD maduro que o projeto ainda não tem.

## 3. Design da Estratégia B

### 3.1 Convenção de Nomenclatura

```
feature/sprint-<NN>-<slug-do-marco>
```

| Sprint | Branch | Marco | Duração |
|---|---|---|---|
| Sprint 3 | `feature/sprint-03-portal-admin` | M2+M3 — Portal Admin + Contas/Planos | ~2 sem |
| Sprint 4 | `feature/sprint-04-rbac` | M4 — RBAC | ~2 sem |
| Sprint 5 | `feature/sprint-05-portal-cliente` | M5 — Portal Cliente | ~2 sem |
| Sprint 6 | `feature/sprint-06-bus-catalogo` | M6 — BUs e Catálogo | ~2 sem |
| Sprint 7 | `feature/sprint-07-homologacao` | M7 — Homologação | ~2 sem |

### 3.2 Ciclo de Vida de Cada Branch

```
CRIAR ─── DESENVOLVER ─── PR + REVIEW ─── MERGE NO MAIN ─── DELETAR
  ↑                                                              │
  └────────────────── próxima sprint ────────────────────────────┘
```

**Criar:**
```bash
git checkout main && git pull
git checkout -b feature/sprint-NN-<slug>
```

**Desenvolver:** commits convencionais com prefixo `feat(sprint-NN):` ou `fix(sprint-NN):`

**PR + Review:**
```bash
gh pr create --base main --head feature/sprint-NN-<slug> \
  --title "Sprint NN: <marco>" \
  --body "Entrega da Sprint NN conforme TASKS.md. 🤖 Generated with Claude Code"
```

**Merge:** squash merge ou `--no-ff` no `main`

**Deletar:** branch local e remota removidas após merge

### 3.3 Regra de Ouro

> 🚫 Nenhum commit do projeto `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` pode ser feito diretamente em `main`. Todo desenvolvimento passa por `feature/sprint-NN-*`.

### 3.4 Hotfix em Sprint Anterior

Se uma sprint já mergeada precisar de correção:
```bash
# Encontra o merge commit da sprint
git log --oneline --merges main | grep "Sprint NN"
# Cria branch de hotfix a partir desse ponto
git checkout -b hotfix/sprint-NN-<descricao> <merge-commit-hash>
```

### 3.5 Integração com Artefatos do Projeto

- **PRD.md §8.4:** documento canônico da estratégia — tabela completa de branches, workflow, comandos git
- **SPRINT-CARD.md de cada sprint:** referência rápida com a branch correta para o executor
- **Demais artefatos** (SPRINT-REVIEW.md, SPRINT-TEST-SUITE.md, relatórios): atualizados sob demanda — são históricos, não bloqueiam

## 4. Plano de Migração

| Passo | Ação | Status |
|---|---|---|
| 1. Criar branch | `git checkout -b feature/sprint-03-portal-admin` a partir de `main` | ⬜ |
| 2. Atualizar PRD.md | Reescrever §8.4 com estratégia B | ⬜ |
| 3. Atualizar artefatos | Atualizar SPRINT-CARD.md de cada sprint (3 a 7) | ⬜ |
| 4. Commit inicial | `docs(strategy): adota branching por sprint` na nova branch | ⬜ |

## 5. Referências

- PRD.md v1.6 — Product Requirements Document
- TASKS.md v2.4 — 80 tarefas, 7 sprints
- [Git Flow](https://nvie.com/posts/a-successful-git-branching-model/) — modelo original de branching
- [GitHub Flow](https://docs.github.com/en/get-started/using-github/github-flow) — referência de feature branches com PRs
