# 📌 Pull Request #6 — Reforma Tributária 2026 (Fases 0, 1, 2)

* **URL:** <https://github.com/bolismar69/workspace-fbso/pull/6>
* **Branch:** `feature/reforma-tributaria-2026-fases-0-1-2` → `main`
* **Data de criação:** 2026-06-25
* **Repositório:** `bolismar69/workspace-fbso` (monorepo)
* **Status:** 🟢 Aberta
* **Projeto:** [PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO](../business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/PRD.md)
* **Sumário de implementação:** [2026-06-25-073636-reforma-tributaria-fases-0-1-2.md](../skill-output/2026-06-25-073636-reforma-tributaria-fases-0-1-2.md)

---

## 📊 Estatísticas

| Métrica | Valor |
|---|---|
| Arquivos alterados | **176** |
| Linhas adicionadas | **+31,620** |
| Linhas removidas | **−119** |
| Pacotes Go | **14** |
| Endpoints REST | **18** |
| Schemas OpenAPI | **30+** |
| Tabelas SQL novas | **3** |
| Testes unitários | **~215** (100% passing) |
| GAPs implementados | **10** (Fases 0, 1, 2) |
| TASKS.md concluído | **18/18** ✅ |
| Tamanho imagem Docker | **46.4 MB** |

---

## 🛠️ Ações Realizadas

### 1. Análise pré-PR (2026-06-25)

| Aspecto | Estado |
|---|---|
| Branch atual | `main` ⚠️ (necessário criar feature branch) |
| Remote verificado | `git@github.com:bolismar69/workspace-fbso.git` ✅ |
| Arquivos no working tree | 61 (espalhados por múltiplos projetos do monorepo) |
| `ms-billing-engine-tax-rates/` | Diretório inteiramente novo (untracked) |
| `taxnexus-billing-core-lib/` | Mistura de M/D/?? (refactor: `tax_models.go` → múltiplos arquivos) |

### 2. Criação da Feature Branch

```bash
git checkout -b feature/reforma-tributaria-2026-fases-0-1-2
```

Branch criada a partir de `main` (commit `10c6935`).

### 3. Stage Seletivo (Apenas Escopo do Projeto)

```bash
git add backend/go/fiber/microservices/ms-billing-engine-tax-rates/  # 🆕 microserviço
git add backend/go/libs/go-native/taxnexus-billing-core-lib/          # 🔄 core-lib refactor
git add .specs/                                                        # 📋 specs + sumário
git add business-inputs/                                               # 📁 documentos de negócio
```

#### Arquivos incluídos (176)

| Grupo | Tipo | Descrição |
|---|---|---|
| `ms-billing-engine-tax-rates/` | 🆕 Novo | 14 pacotes Go + Dockerfile + K8s + docs |
| `taxnexus-billing-core-lib/` | 🔄 Refactor | 6 modificados, 2 deletados, 10 novos |
| `.specs/` (raiz) | 🆕 Novo | ci-cd, security, prompts, feature-roadmap |
| `business-inputs/` | 🆕 Novo | POLICE, RULES, PROCEDURE, EPICS, FEATURES, USER-STORYS |

#### Arquivos excluídos (39 — outros projetos)

| Grupo | Motivo |
|---|---|
| `ms-tax-individual-income/` | Projeto IRPF (diferente) |
| `ms-tax-nexus-taas-calc-engine/` | Projeto de cálculo diferente |
| `taxnexus-individual-core-lib/` | Lib de IRPF (diferente) |
| `batch-geolocalidade/` (Java) | Projeto Java, não relacionado |
| `ms-shoppingcart-engine/` | Contexto diferente |
| `ms-product-catalog-admin-simple/` | Contexto diferente |
| `README-FRAMEWORKS-MODELO.md` | Documentação geral do workspace |

### 4. Commit

```bash
git commit -m "feat(PRJ-FIN-2026-0001): Reforma Tributária 2026 — Fases 0, 1, 2\n\n..."
```

| Campo | Valor |
|---|---|
| Tipo | `feat` (Conventional Commits) |
| Escopo | `PRJ-FIN-2026-0001` |
| Hash | `83f6905` |
| Co-autoria | `Co-Authored-By: Claude <noreply@anthropic.com>` |

### 5. Push

```bash
git push -u origin feature/reforma-tributaria-2026-fases-0-1-2
```

Novo branch remoto criado com tracking configurado.

### 6. Criação da Pull Request

```bash
gh pr create \
  --base main \
  --head feature/reforma-tributaria-2026-fases-0-1-2 \
  --title "feat(PRJ-FIN-2026-0001): Reforma Tributária 2026 — Fases 0, 1, 2" \
  --body-file "backend/go/fiber/microservices/ms-billing-engine-tax-rates/.specs/skill-output/2026-06-25-073636-reforma-tributaria-fases-0-1-2.md"
```

- O corpo da PR contém o **sumário completo de implementação** (Markdown renderizado nativamente pelo GitHub).
- Ferramenta: GitHub CLI (`gh`) v2.x, autenticada como `bolismar69` com escopos `audit_log`, `project`, `read:org`, `repo`, `workflow`, `write:packages`.

### 7. Resultado

✅ **PR #6 criada**: <https://github.com/bolismar69/workspace-fbso/pull/6>

---

## 📦 Detalhamento por Fase

### FASE 0 — Fundação (GAPs 008, 009, 010)

| GAP | Feature | Artefatos |
|---|---|---|
| GAP-008 | Rate Limiting | `internal/middleware/ratelimit.go` (sliding window IP+JWT, 8 testes) |
| GAP-009 | API Versioning | `/v1/*` endpoints, headers `Deprecation`/`Sunset`/`Link` RFC 8594, 5 testes |
| GAP-010 | Deploy Artifacts | `Dockerfile` (multi-stage, 46.4MB), `deploy/k8s/*` (deployment, service, configmap, hpa), `docker-compose.yaml` |

### FASE 1 — Onda 1 Comercial (GAPs 001, 002, 003, 004)

| GAP | Feature | Artefatos |
|---|---|---|
| GAP-004 | valor_liquido | `internal/calculator/engine.go` — `valor_liquido = valorItem − impostos` (piso zero, BR-04, 4 testes) |
| GAP-002 | TaxToken | `internal/token/` — snapshot CBS/IBS/IS com TTL configurável (8 testes) |
| GAP-003 | /simulate | `internal/simulation/` — projeção multi-destino com `ImpactoUF` (8 testes) |
| GAP-001 | Admin Fiscal | `internal/admin/` — upsert `iva_dual_rules`, cache Redis, auditoria (10 testes) |

### FASE 2 — Onda 2 Financeira (GAPs 005, 006, 007)

| GAP | Feature | Artefatos |
|---|---|---|
| GAP-006 | Split Payment | `internal/calculator/engine.go` — SHA-256 barcode, receita líquida, CBS/IBS/IS a reter (5 testes) |
| GAP-005 | Créditos | `internal/credit/` — crédito CBS/IBS de NF-e entrada, `permite_credito_amplo` (8 testes) |
| GAP-007 | Fornecedores | `internal/supplier/` — validação fiscal CRUD, Lucro Real/Presumido/Simples (8 testes) |

---

## 🔒 Segurança

| Controle | Status |
|---|---|
| RBAC (`admin`/`fiscal`/`credit`) | ✅ Em todos endpoints sensíveis |
| Rate limiting anti-spoof | ✅ IP via `c.IP()` com trusted proxy |
| Métricas protegidas | ✅ `METRICS_REQUIRE_AUTH=true` |
| Fallback removido | ✅ Sem `"admin-api"` default |
| Menor privilégio | ✅ Públicos apenas `/v1/healthz`, `/v1/health` |
| `go vet ./...` | ✅ zero warnings |
| Credenciais hardcoded | ✅ Nenhuma |

### ⚠️ Vulnerabilidade pendente (MEDIUM)

**Spoofable-Field Auth Bypass — Fail-Open JWT Decoding** em `internal/middleware/auth.go`:

- `decodeJWTClaims()` não verifica assinatura do JWT (confia no Kong/Keycloak da borda)
- Middleware é **fail-open**: tokens inválidos/malformados são logados com warn e passam (`return c.Next()`)
- Risco: atacante com acesso ao ClusterIP do serviço pode forjar claims arbitrárias
- Correção recomendada: adicionar `github.com/golang-jwt/jwt/v5` com validação JWKS + fail-closed (401)

---

## 📄 Documentos Vinculados

| Documento | Localização |
|---|---|
| PRD | `business-inputs/business-projects/PRJ-FIN-2026-0001-*/PRD.md` |
| Architecture | `.specs/business-projects/PRJ-FIN-2026-0001-*/ARCHITECTURE.md` |
| Specs | `.specs/business-projects/PRJ-FIN-2026-0001-*/SPECS.md` |
| Tasks | `.specs/business-projects/PRJ-FIN-2026-0001-*/TASKS.md` (18/18 ✅) |
| Test Plan | `.specs/business-projects/PRJ-FIN-2026-0001-*/TEST_PLAN.md` |
| API Spec | `.specs/api/tax-rates-api.yaml` (v1.1.0, 30+ schemas) |
| ERD | `.specs/architecture/erd.md` |
| Security | `.specs/security/SECURITY.md` |
| Business Rules | `business-inputs/business-documents/business-rules/RULES-FIN-00001-*.md` |
| Business Policies | `business-inputs/business-documents/business-policies/POLICE-FIN-00001-*.md` |
| Business Procedures | `business-inputs/business-documents/business-procedures/PROCEDURE-FIN-00001-*.md` |

---

## 🧪 Evidências de Testes

```text
ok  	ms-billing-engine-tax-rates/cmd/api	        0.011s
ok  	ms-billing-engine-tax-rates/internal/admin	0.009s
ok  	ms-billing-engine-tax-rates/internal/calculator	0.063s
ok  	ms-billing-engine-tax-rates/internal/circuitbreaker	0.106s
ok  	ms-billing-engine-tax-rates/internal/credit	0.005s
ok  	ms-billing-engine-tax-rates/internal/ibsclient	0.144s
ok  	ms-billing-engine-tax-rates/internal/legacy	0.014s
ok  	ms-billing-engine-tax-rates/internal/middleware	0.022s
ok  	ms-billing-engine-tax-rates/internal/phase	        0.006s
ok  	ms-billing-engine-tax-rates/internal/reforma	0.004s
ok  	ms-billing-engine-tax-rates/internal/simulation	0.004s
ok  	ms-billing-engine-tax-rates/internal/supplier	0.004s
ok  	ms-billing-engine-tax-rates/internal/token	        0.004s

go vet ./... — clean (zero warnings)
go build ./... — success
Docker build — success (46.4MB)
```

---

## 📋 Dívidas Técnicas Remanescentes

| DT | Descrição | Impacto |
|---|---|---|
| DT-03 | CSTs provisórios da Reforma (aguardando tabela oficial RFB) | Não-conformidade futura |
| DT-04 | Créditos da Reforma (cash forward) | Bloqueia BR-08 avançado |
| DT-09 | API do Comitê Gestor IBS não publicada | Fallback DB ativo |

---

## 🔗 Links

- **Pull Request:** <https://github.com/bolismar69/workspace-fbso/pull/6>
- **Repositório:** <https://github.com/bolismar69/workspace-fbso>
- **Branch:** `feature/reforma-tributaria-2026-fases-0-1-2`
- **Commit:** `83f6905`
- **Projeto:** `PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO`

---

🤖 *Registro gerado em 2026-06-25. Histórico completo da criação da PR #6 para consulta humana e IA.*
