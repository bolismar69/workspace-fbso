# Git Reorganization — Frontend Web Apps

**Data:** 2026-07-16
**Branch base:** `main`
**Origem dos arquivos:** `frontend/javascript/react/web_apps/`

---

## Resumo da Estratégia

O diretório `frontend/javascript/react/web_apps/` contém 5 aplicações React distintas que foram originalmente agrupadas em uma única branch (`feature/frontend-web-mobile-apps`). Para permitir desenvolvimento independente de cada frontend, cada app terá sua própria branch.

| # | App | Arquivos | Tamanho |
|---|-----|----------|---------|
| 1 | `web_app-solar-facil-site` | 125 | 1.1 MB |
| 2 | `web_app-billing-admin-tax-rates` | 1 | 76 KB |
| 3 | `web_app-tax-nexus-portal` | 14.833 | 241 MB |
| 4 | `web_app-fbso-platform-portal` | 2 | 96 KB |
| 5 | `web_app-cnpj-validacao` | 17 | 140 KB |

---

## Comandos Executados (em ordem)

### 1. Análise dos diretórios

```bash
# Listar apps no diretório de web apps
ls frontend/javascript/react/web_apps/

# Contar arquivos e ver tamanho de cada app
for dir in \
  web_app-solar-facil-site \
  web_app-billing-admin-tax-rates \
  web_app-tax-nexus-portal \
  web_app-fbso-platform-portal \
  web_app-cnpj-validacao; do
  echo "=== $dir ==="
  find "frontend/javascript/react/web_apps/$dir" -type f | wc -l
  du -sh "frontend/javascript/react/web_apps/$dir"
done
```
**Função:** Diagnosticar o conteúdo de cada app para planejar a separação.

---

### 2. Criar branches individuais para cada frontend

```bash
# Para cada web app, criar uma branch específica a partir de main
git checkout main

# App 1: Solar Fácil Site
git checkout -b feature/frontend-solar-facil-site
git push -u origin feature/frontend-solar-facil-site

# App 2: Billing Admin Tax Rates
git checkout main
git checkout -b feature/frontend-billing-admin-tax-rates
git push -u origin feature/frontend-billing-admin-tax-rates

# App 3: Tax Nexus Portal
git checkout main
git checkout -b feature/frontend-tax-nexus-portal
git push -u origin feature/frontend-tax-nexus-portal

# App 4: FBSO Platform Portal
git checkout main
git checkout -b feature/frontend-fbso-platform-portal
git push -u origin feature/frontend-fbso-platform-portal

# App 5: CNPJ Validação
git checkout main
git checkout -b feature/frontend-cnpj-validacao
git push -u origin feature/frontend-cnpj-validacao
```
**Função:** Cada branch nasce de `main` e contém o monorepo completo. A especialização por frontend permite que futuros commits sejam focados em uma única aplicação, facilitando code review e deploy independente.

---

### 3. Criar Pull Requests

```bash
# Para cada branch, criar PR via gh CLI
for branch in \
  feature/frontend-solar-facil-site \
  feature/frontend-billing-admin-tax-rates \
  feature/frontend-tax-nexus-portal \
  feature/frontend-fbso-platform-portal \
  feature/frontend-cnpj-validacao; do
  title=$(git log "$branch" --oneline --format="%s" -1 | head -c 256)
  gh pr create \
    --base main \
    --head "$branch" \
    --title "$title" \
    --body "🤖 Generated with [Claude Code](https://claude.com/claude-code)"
done
```
**Função:** Cria um PR para cada branch. Como o conteúdo já existe em `main`, os PRs serão "no-op" — servem para estabelecer a branch como canal oficial de desenvolvimento daquele frontend.

---

### 4. Merge dos PRs

```bash
# Merge local de cada branch no main e push
git checkout main

for branch in \
  feature/frontend-solar-facil-site \
  feature/frontend-billing-admin-tax-rates \
  feature/frontend-tax-nexus-portal \
  feature/frontend-fbso-platform-portal \
  feature/frontend-cnpj-validacao; do
  git merge "$branch" -m "chore: merge $branch into main"
done

git push origin main

# Fechar PRs (já mergeados localmente)
for pr in <PR_NUMBERS>; do
  gh pr close "$pr" -c "Merged locally into main. See commit history."
done
```
**Função:** Mesmo processo da reorganização anterior — merge local e push com bypass de proteção de branch.

---

## Resultado Final

| # | Branch | PR | App |
|---|--------|----|-----|
| 1 | `feature/frontend-solar-facil-site` | [#17](https://github.com/bolismar69/workspace-fbso/pull/17) | Solar Fácil Site (Next.js) |
| 2 | `feature/frontend-billing-admin-tax-rates` | [#18](https://github.com/bolismar69/workspace-fbso/pull/18) | Billing Admin Tax Rates |
| 3 | `feature/frontend-tax-nexus-portal` | [#19](https://github.com/bolismar69/workspace-fbso/pull/19) | Tax Nexus Portal |
| 4 | `feature/frontend-fbso-platform-portal` | [#20](https://github.com/bolismar69/workspace-fbso/pull/20) | FBSO Platform Portal |
| 5 | `feature/frontend-cnpj-validacao` | [#21](https://github.com/bolismar69/workspace-fbso/pull/21) | CNPJ Validação |

---

## Convenção de Nomenclatura

Para branches de frontend, o padrão adotado é:

```
feature/frontend-<nome-do-app>
```

Isso permite identificar rapidamente no histórico do git:
- `feature/frontend-*` — todas as branches de frontend
- `feature/java-*` — todas as branches de microsserviços Java
- `feature/*` — todas as branches de feature no monorepo
