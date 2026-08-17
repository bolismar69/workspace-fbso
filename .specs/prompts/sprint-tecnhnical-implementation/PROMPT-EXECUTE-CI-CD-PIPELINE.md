# PROMPT-EXECUTE-CI-CD-PIPELINE

## Contexto

Este prompt **implementa os pipelines de CI/CD concretos** de uma solução técnica, materializando os planos definidos nos documentos-base do projeto (041/087/090 no padrão WATERFALL, ou equivalentes do TECHLEAD em outros contextos) e nas definições técnicas do TECHLEAD — independente da metodologia adotada. Ele não inventa estratégia — ele traduz os componentes `DED-NN` (041-DEVOPS-SETUP), `CICD-NN` (087-PLANO-CI-CD-AMBIENTES) e as quality gates da 510-TEST-STRATEGY em arquivos de pipeline reais (GitHub Actions / GitLab CI), prontos para uso.

**Princípios fundamentais:**

1. **Os documentos-base são a fonte da verdade:** nenhuma etapa de pipeline pode existir aqui sem estar ancorada em `041`, `087`, `090`, `500`, `510` ou `086`. Pipeline inventado = violação.
2. **Stack-agnóstico:** o prompt adapta-se a qualquer linguagem/framework (o build/teste/lint seguem os comandos já padronizados no `PROMPT-EXECUTE-SPRINT-TASKS`).
3. **Gate obrigatório:** ao final, o pipeline implementado passa pelo `PROMPT-QA-REVISOR-SECURITY` e pela validação humana antes de ser aceito.
4. **Localidade:** os arquivos de pipeline são criados dentro do repositório da solução (`{SOLUTION_PATH}/.github/workflows/` ou `.gitlab-ci.yml`), nunca em diretórios globais.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|:---|:---|:---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_NAME}` | Nome/código do projeto de negócio | `PRJ-TEC-2026-0004-PROJETO-SHIELD` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{CI_PROVIDER}` | Provedor de CI/CD (deve constar no 500-DEVOPS-SRE-DEFINITION) | `github-actions` ou `gitlab-ci` |
| `{PROJECT_DOCS_DIR}` | Pasta dos documentos WATERFALL do projeto (041/043/086/087/090) | `/home/user/work/business-inputs/business-projects/PRJ-TEC-2026-0004-PROJETO-SHIELD/` |
| `{TECH_DEFS_DIR}` | Pasta das definições técnicas TECHLEAD (480/500/510/520/550) | `/home/user/work/business-inputs/business-projects/PRJ-TEC-2026-0004-PROJETO-SHIELD/technical-definitions/` |

---

## Documentos de Referência (obrigatórios — fonte da verdade)

```
Ler obrigatoriamente antes de gerar qualquer pipeline:

    ├── {PROJECT_DOCS_DIR}/041-DEVOPS-SETUP-*.md                  ← Pipelines DED-NN (build/test/deploy/rollback), IaC, observabilidade
    ├── {PROJECT_DOCS_DIR}/087-PLANO-CI-CD-AMBIENTES-*.md         ← Estratégia de branches, pipelines CICD-NN, ambientes DEV/QA/HMG/PROD, GMUD
    ├── {PROJECT_DOCS_DIR}/090-STRATEGIC-IMPLEMENTATION-AND-DEPLOYMENT-PLAN-*.md ← Estratégia de deploy (blue-green/canary/rolling) e rollback
    ├── {PROJECT_DOCS_DIR}/086-PADROES-CODIGO-DOD-*.md            ← DOD-NN (critérios de aceite) e checklists de revisão de código
    ├── {TECH_DEFS_DIR}/500-DEVOPS-SRE-DEFINITION.md              ← Ferramentas de CI/CD e IaC do projeto
    ├── {TECH_DEFS_DIR}/510-TEST-STRATEGY-DEFINITION.md           ← Quality gates (cobertura ≥80%, SAST/SCA/DAST) e thresholds
    └── {SOLUTION_PATH}/README.md                                 ← Comandos de build/teste/lint da solução
```

> ⚠️ Se algum documento-base não existir no caminho informado → **PARE** e solicite o caminho correto ao humano. NUNCA inferir pipelines sem ancoragem documental.

---

## Missão

Implementar, no repositório da solução `{SOLUTION_NAME}`, os pipelines de CI/CD especificados nos documentos-base, usando `{CI_PROVIDER}`, com os estágios `build → lint → test → scan → deploy` e approval gates por ambiente.

---

## Fluxo de Execução

### Fase 0 — Pré-implementação

0. **Validar parâmetros e branch:**
   - Confirmar que `{CI_PROVIDER}` consta na seção de ferramentas do `500-DEVOPS-SRE-DEFINITION` (se divergir → perguntar ao humano).
   - Trabalhar em branch `feature/ci-cd-<slug>` criada a partir da branch principal (nunca main/master diretamente).
1. **Carregar documentos-base** — ler os 7 documentos da seção acima e extrair:
   - De `041`: lista de pipelines `DED-NN` e suas etapas + componentes de rollback.
   - De `087`: matriz de ambientes (DEV/QA/HMG/PROD), proteções de branch, critérios de merge (PR + code review + CI verde) e pipeline `CICD-NN` por ambiente.
   - De `090`: estratégia de deploy em produção (blue-green/canary/rolling) e procedimento de rollback.
   - De `510`: quality gates objetivos (ex.: cobertura unitária ≥80%, SAST sem achados críticos, SCA sem CVEs críticos/alto).
   - De `086`: itens de DoD que o pipeline deve automatizar (lint, testes, revisão).
   - De `500`: ferramentas concretas (ex.: GitHub Actions + Semgrep + Dependabot).

### Fase 1 — Desenho do Mapa de Pipelines

2. **Gerar o mapa de implementação** (arquivo `{SOLUTION_PATH}/.specs/pull-requests/`... não — dentro do relatório final): uma tabela `CICD-NN → arquivo(s) de pipeline → estágios → ambiente → gate de aprovação`, rastreando cada pipeline de volta ao seu ID no 087/041.

### Fase 2 — Implementação dos Arquivos de Pipeline

3. **Para cada pipeline do mapa, gerar o arquivo conforme o provedor:**

   - **GitHub Actions:** criar `{SOLUTION_PATH}/.github/workflows/ci.yml` (build → lint → test → scan) e `deploy-<ambiente>.yml` por ambiente com `environment:` protegido e aprovação manual (GMUD), além do job de rollback do 090.
   - **GitLab CI:** criar `.gitlab-ci.yml` com `stages: [build, lint, test, scan, deploy]`, `rules` por branch (padrão do 087) e `when: manual` nos gates de deploy.
   - Estágios obrigatórios por pipeline (se constarem no 087/041):
     | Estágio | Conteúdo | Origem |
     |:---|:---|:---|
     | build | Compilação usando o comando do README.md | 041-DEVOPS-SETUP |
     | lint | Linters da stack (checkstyle/pmd, eslint, ruff...) | 086-PADROES-CODIGO |
     | test | Testes unitários/integração + cobertura (gate ≥80%) | 510-TEST-STRATEGY |
     | scan | SAST/SCA/secret-scanning — invocar `PROMPT-EXECUTE-CVE-SCA-SCAN` para configurar as ferramentas | 480/510 |
     | deploy | Deploy por ambiente com approval gates e rollback | 087/090 |

### Fase 3 — Validação e Relatório

4. **Validar sintaxe** dos arquivos gerados (ex.: `actionlint` para GitHub Actions; checar formato YAML).
5. **Gerar o relatório de implementação** em `{SOLUTION_PATH}/.specs/skill-output/{AAAA-MM-DD-HHMMSS}-TASK-EXECUTED-ci-cd-pipeline-[slug].md` seguindo o modelo do `PROMPT-GENERATE-IMPLEMENTATION-REPORT`, com: mapa CICD-NN → arquivo, estágios implementados, gates, rastreabilidade com 041/087/090/510.
6. **Gate:** submeter o pipeline implementado ao `PROMPT-QA-REVISOR-SECURITY` (referência: 510 + 480). `[STATUS: APPROVED]` → prosseguir para abertura de PR via `PROMPT-GENERATE-PULL-REQUEST`. `[STATUS: FAILED]` → corrigir e revalidar (máx. 3 loops).
7. **Validação humana:** apresentar o mapa de pipelines e aguardar aprovação antes do merge.

---

## Regras de Ouro

1. NUNCA criar etapa de pipeline que não esteja ancorada em 041/087/090/510/086.
2. NUNCA publicar deploy automático em PROD sem approval gate (GMUD do 087/090 em contexto WATERFALL; approval gates equivalentes em contexto ágil).
3. SEMPRE incluir job de rollback (estratégia do 090) em pipelines de deploy.
4. SEMPRE anexar o relatório TASK-EXECUTED na PR.
5. Toda correção pós-gate é cirúrgica — alterar apenas o que o gate reprovou.
