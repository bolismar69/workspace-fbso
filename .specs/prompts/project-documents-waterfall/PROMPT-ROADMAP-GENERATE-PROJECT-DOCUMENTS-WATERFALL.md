# PROMPT: ROADMAP DE EXECUÇÃO MACRO E GUIA DE ORQUESTRAÇÃO DE DOCUMENTOS — METODOLOGIA WATERFALL
## Versão: 1.0 — Bootstrap Inteligente, Fluxo de Dados Explícito, Validação Soberana Humana (Human-in-the-Loop) e Git Workflow Automatizado

Atue como um Especialista em Gestão de Processos (BPM) e Arquiteto de Soluções Organizacionais, especializado em metodologia WATERFALL e Engenharia de Prompts.

Preciso que você execute um roadmap de criação, revisão, evolução e validação de 20 documentos base de um projeto, seguindo estritamente a metodologia WATERFALL em 5 fases sequenciais.

Objetivo Principal: Garantir que todos os documentos estejam criados, revisados e 100% alinhados conceitualmente entre si (rastreabilidade vertical de ponta a ponta), mitigando desvios de escopo (scope creep) e garantindo o sucesso do projeto.

Regra Crítica de Execução (Gating Rule): O processo é estritamente sequencial. Nenhum documento subsequente pode ser iniciado sem que o documento atual esteja marcado como `[STATUS: COMPLIANCE]` e aprovado explicitamente pelo humano.

---

## REGRAS DE OURO (7 REGRAS DE GATING — NÃO NEGOCIÁVEIS)

1. **NÃO ADIVINHAR INPUTS:** Nenhum prompt infere, busca ou descobre seus inputs. Todo parâmetro é passado explicitamente por este orquestrador.
2. **TODO GENERATE TEM GATE+FIX:** Cada um dos 20 documentos tem exatamente um trio GENERATE, GATE e FIX.
3. **FLUXO DE DADOS EXPLÍCITO:** `DOC_PATH` flui GENERATE→GATE→FIX. `VIOLATIONS[]` flui GATE→FIX. Sempre como parâmetros nomeados.
4. **FIX É CIRÚRGICO:** O prompt FIX edita apenas as seções com violações reportadas pelo GATE. Nunca regenera ou recria o documento.
5. **STATUS INICIAL: EM ANÁLISE:** Ao criar o documento, GENERATE escreve `[STATUS: Em análise]` no cabeçalho.
6. **EM REVISÃO DURANTE GATE/FIX:** GATE e FIX alteram o status para `[STATUS: Em revisão]`.
7. **SÓ AVANÇA COM COMPLIANCE:** O roadmap só avança para o próximo documento quando o atual estiver marcado `[STATUS: COMPLIANCE]` e o humano confirmar.

---

## VARIÁVEIS DE ENTRADA E BOOTSTRAP DO PROJETO (FASE 0)

### Tabela de Inputs Obrigatórios

| Variável | Descrição | Exemplo |
|---|---|---|
| `PROJECT_PATH` | Caminho base onde os projetos residem | `/home/bolismar/work/workspace-fbso/business-inputs/business-projects` |
| `PROJECT_ID` | Identificador único do projeto (ID corporativo) | `PRJ-FIN-2026-0003` |
| `PROJECT_NAME` | Nome curto do produto/projeto | `SAAS-FBSO-ORG` |
| `PROJECT_BRIEFING` | Briefing do projeto (texto inline ou caminho de arquivo) | `"Portal de autoatendimento..."` ou `/tmp/briefing.md` |
| `PROJECT_DOCUMENTS_INPUTS` | Lista de caminhos para documentos brutos de entrada | `[]` |
| `PROJECT_PROMPT_INPUTS` | Lista de caminhos para prompts auxiliares ou contextos adicionais | `[]` |
| `PROMPT_BRANCH` | Nome da branch Git. **Não pode** ser `main`, `master` ou `develop`. | `feature/PRJ-FIN-2026-0003-docs` |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_ID_NAME            = PROJECT_ID + "-" + PROJECT_NAME
PROJECT_COMPLETE_PATH_NAME = PROJECT_PATH + "/" + PROJECT_ID_NAME
```

### Workflow de Bootstrap (Execução Obrigatória)

#### Passo 0.1 — Solicitar Inputs ao Usuário

Se alguma das 7 variáveis não tiver sido fornecida, pergunte:

> "Para iniciar o Roadmap WATERFALL, preciso das seguintes informações:
> 1. **PROJECT_PATH** — Caminho base dos projetos
> 2. **PROJECT_ID** — ID do projeto
> 3. **PROJECT_NAME** — Nome do produto
> 4. **PROJECT_BRIEFING** — Briefing (texto ou caminho de arquivo)
> 5. **PROJECT_DOCUMENTS_INPUTS** — Documentos de entrada (deixe `[]` se não houver)
> 6. **PROJECT_PROMPT_INPUTS** — Prompts auxiliares (deixe `[]` se não houver)
> 7. **PROMPT_BRANCH** — Nome da branch Git. **Não pode** ser `main`, `master` ou `develop`."

**Validar PROMPT_BRANCH:** Se `main`, `master` ou `develop`, aplicar bloqueio:

> ⛔ **Branch Inválida.** O processo NÃO pode ser executado em branches protegidas. Informe um nome de branch de trabalho.

#### Passo 0.2 — Exibir Caminho Derivado e Solicitar Confirmação

> **📁 Caminho do Projeto:** `{PROJECT_COMPLETE_PATH_NAME}`
> **🏷️ Identificador:** `{PROJECT_ID_NAME}`
> **🌿 Branch Git:** `{PROMPT_BRANCH}`
>
> Confirma?
> - **SIM** → Prosseguir
> - **NÃO** → Corrigir inputs e repetir

#### Passo 0.3 — Criar Estrutura de Diretórios

```bash
mkdir -p {PROJECT_COMPLETE_PATH_NAME}
```

#### Passo 0.4 — Verificar Status dos Arquivos

Verificar existência e status de cada um dos 20 documentos. Para arquivos existentes, ler o cabeçalho e buscar por `[STATUS: COMPLIANCE]`.

| # | Arquivo | Status |
|---|---|---|
| 1 | `01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` | ✅/❌ |
| 2 | `02-BRD-{PROJECT_ID_NAME}.md` | ✅/❌ |
| ... | ... (todos os 20) | ... |
| 20 | `20-LICOES-APRENDIDAS-{PROJECT_ID_NAME}.md` | ✅/❌ |

**Decisão:**
- Todos ❌ → Iniciar da Fase 1, Documento 1 (PROJECT-CHARTER)
- Parcial → Iniciar do primeiro documento sem `[STATUS: COMPLIANCE]`
- Todos ✅ → Perguntar: revisar, novo ciclo ou encerrar

#### Passo 0.5 — Apresentar Resumo e Iniciar

> **📊 Resumo:** `{PROJECT_ID_NAME}` em `{PROJECT_COMPLETE_PATH_NAME}`
> **📝 Próximo:** Fase X, Documento Y — {NOME}
> Iniciando...

---

## FASES DO ROADMAP WATERFALL

### FASE 1 — INICIALIZAÇÃO

| # | Documento | Arquivo |
|---|---|---|
| 1 | PROJECT-CHARTER | `01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` |

### FASE 2 — PLANEJAMENTO E REQUISITOS

| # | Documento | Arquivo |
|---|---|---|
| 2 | BRD | `02-BRD-{PROJECT_ID_NAME}.md` |
| 3 | SRS | `03-SRS-{PROJECT_ID_NAME}.md` |
| 4 | RTM | `04-RTM-{PROJECT_ID_NAME}.md` |
| 5 | EAP/WBS | `05-EAP-WBS-{PROJECT_ID_NAME}.md` |
| 6 | Cronograma/Gantt | `06-CRONOGRAMA-GANTT-{PROJECT_ID_NAME}.md` |
| 7 | Orçamento | `07-ORCAMENTO-{PROJECT_ID_NAME}.md` |
| 8 | Plano de Comunicação | `08-PLANO-COMUNICACAO-{PROJECT_ID_NAME}.md` |
| 9 | Plano de Riscos | `09-PLANO-RISCOS-{PROJECT_ID_NAME}.md` |

### FASE 3 — DESIGN E ARQUITETURA

| # | Documento | Arquivo |
|---|---|---|
| 10 | SAD | `10-SAD-{PROJECT_ID_NAME}.md` |
| 11 | HLD | `11-HLD-{PROJECT_ID_NAME}.md` |
| 12 | LLD | `12-LLD-{PROJECT_ID_NAME}.md` |

### FASE 4 — TESTES E QUALIDADE

| # | Documento | Arquivo |
|---|---|---|
| 13 | TEST-PLAN | `13-TEST-PLAN-{PROJECT_ID_NAME}.md` |
| 14 | TEST-CASES | `14-TEST-CASES-{PROJECT_ID_NAME}.md` |
| 15 | Relatório de Qualidade | `15-RELATORIO-QUALIDADE-{PROJECT_ID_NAME}.md` |

### FASE 5 — IMPLANTAÇÃO E ENCERRAMENTO

| # | Documento | Arquivo |
|---|---|---|
| 16 | DEPLOYMENT-PLAN | `16-DEPLOYMENT-PLAN-{PROJECT_ID_NAME}.md` |
| 17 | Manuais de Usuário | `17-MANUAIS-USUARIO-{PROJECT_ID_NAME}.md` |
| 18 | Manuais Operacionais | `18-MANUAIS-OPERACIONAIS-{PROJECT_ID_NAME}.md` |
| 19 | Termo de Aceite | `19-TERMO-ACEITE-{PROJECT_ID_NAME}.md` |
| 20 | Lições Aprendidas | `20-LICOES-APRENDIDAS-{PROJECT_ID_NAME}.md` |

---

## LOOP DE EXECUÇÃO POR DOCUMENTO

Para cada documento `{DOC}` na ordem sequencial acima:

### STEP 1: Computar inputs para GENERATE

```
DOC_PATH    = {PROJECT_COMPLETE_PATH_NAME}/{NN}-{DOC-SLUG}-{PROJECT_ID_NAME}.md
UPSTREAM_DOCS = [lista de paths de documentos upstream já em COMPLIANCE]
BRIEFING    = PROJECT_BRIEFING
EXTRA_INPUTS = PROJECT_DOCUMENTS_INPUTS + PROJECT_PROMPT_INPUTS
SKILLS      = [lista de skills para este documento — vide tabela abaixo]
```

### STEP 2: Invocar GENERATE

Invocar `project-documents-waterfall/PROMPT-GENERATE-{DOC-SLUG}.md` passando **explicitamente**:
- `DOC_PATH`, `PROJECT_ID_NAME`, `BRIEFING`, `UPSTREAM_DOCS`, `EXTRA_INPUTS`, `SKILLS`

GENERATE deve:
- Criar o arquivo em `DOC_PATH`
- Escrever `[STATUS: Em análise]` como status inicial
- Retornar `{DOC_PATH}`

### STEP 3: Invocar GATE

Invocar `project-documents-waterfall/PROMPT-GATE-{DOC-SLUG}.md` passando **explicitamente**:
- `DOC_PATH` (recebido do GENERATE)
- `CHECKLIST` (definido no próprio prompt GATE)

GATE deve:
- Ler apenas o arquivo em `DOC_PATH`
- Atualizar status para `[STATUS: Em revisão]`
- Retornar `{PASS}` ou `{FAIL, VIOLATIONS: [{section, description, severity}]}`

### STEP 4a: Se GATE retornar FAIL → Invocar FIX

Invocar `project-documents-waterfall/PROMPT-FIX-{DOC-SLUG}.md` passando **explicitamente**:
- `DOC_PATH` (mesmo arquivo)
- `VIOLATIONS[]` (lista exata de não-conformidades do GATE)

FIX deve:
- Editar **apenas** as seções em `VIOLATIONS[]`
- Manter status como `[STATUS: Em revisão]`
- Retornar `{DOC_PATH}`
- → Voltar ao STEP 3 (re-executar GATE)

### STEP 4b: Se GATE retornar PASS → Validação Humana

Apresentar o documento e fazer 3 perguntas:

> **P1:** "O conteúdo deste documento está aderente às necessidades do projeto?"
> **P2:** "Existem novos documentos de entrada que devem ser incorporados a esta fase?"
> **P3:** "Há novas informações textuais, mudanças de escopo ou ajustes a serem considerados?"

- Se humano aprovar (SIM para P1, NÃO para P2/P3): `[STATUS: COMPLIANCE]`, documento congelado, próximo documento liberado
- Se humano fornecer novos inputs: Voltar ao STEP 2 (re-executar GENERATE com novo contexto)

---

## MATRIZ DE SKILLS POR DOCUMENTO

| # | Documento | Skills Primárias | Fallback |
|---|---|---|---|
| 1 | PROJECT-CHARTER | `draft-project-charter`, `senior-pm` | ✅ |
| 2 | BRD | `brd-creation`, `business-analyst`, `requirements-elicitation` | ✅ |
| 3 | SRS | `frs-creation`, `requirements-engineering` | ✅ |
| 4 | RTM | `requirements-modeling`, `requirements-validation` | ✅ |
| 5 | EAP/WBS | `decomposition-planning-roadmap`, `project-estimation` | ✅ |
| 6 | Cronograma/Gantt | `roadmap-planning`, `project-estimation` | ✅ |
| 7 | Orçamento | `project-estimation` | ✅ |
| 8 | Plano de Comunicação | `stakeholder-analysis`, `stakeholder-map` | ✅ |
| 9 | Plano de Riscos | `risk-manager`, `risk-management-specialist` | ✅ |
| 10 | SAD | `software-architecture`, `architecture-designer`, `architecture-patterns` | ✅ |
| 11 | HLD | `c4-container`, `system-design`, `architecture-decision-records` | ✅ |
| 12 | LLD | `c4-component`, `ddd-tactical-patterns`, `database-designer` | ✅ |
| 13 | TEST-PLAN | `test-strategy-design`, `qa-test-planner` | ✅ |
| 14 | TEST-CASES | `test-case-creation`, `acceptance-criteria` | ✅ |
| 15 | Relatório de Qualidade | `quality-documentation-manager`, `qa` | ✅ |
| 16 | DEPLOYMENT-PLAN | `deployment-engineer`, `devops-rollout-plan` | ✅ |
| 17 | Manuais de Usuário | `documentation-generation-doc-generate`, `docs-writer` | ✅ |
| 18 | Manuais Operacionais | `documentation-generation-doc-generate` | ✅ |
| 19 | Termo de Aceite | `contract-and-proposal-writer` | ✅ |
| 20 | Lições Aprendidas | — | ✅ |

---

## MAPEAMENTO DE UPSTREAM DOCS DE CADA DOCUMENTO

Cada GENERATE recebe em `UPSTREAM_DOCS` a lista de documentos anteriores que já estão em COMPLIANCE e que este documento deve referenciar:

| # | Documento | UPSTREAM_DOCS (paths relativos) |
|---|---|---|
| 1 | PROJECT-CHARTER | `[]` (documento raiz) |
| 2 | BRD | `[01-PROJECT-CHARTER]` |
| 3 | SRS | `[01-PROJECT-CHARTER, 02-BRD]` |
| 4 | RTM | `[01-PROJECT-CHARTER, 02-BRD, 03-SRS]` |
| 5 | EAP/WBS | `[01-PROJECT-CHARTER, 02-BRD]` |
| 6 | Cronograma/Gantt | `[01-PROJECT-CHARTER, 05-EAP-WBS]` |
| 7 | Orçamento | `[01-PROJECT-CHARTER, 05-EAP-WBS, 06-Cronograma]` |
| 8 | Plano de Comunicação | `[01-PROJECT-CHARTER]` |
| 9 | Plano de Riscos | `[01-PROJECT-CHARTER]` |
| 10 | SAD | `[01-PROJECT-CHARTER, 02-BRD, 03-SRS]` |
| 11 | HLD | `[01-PROJECT-CHARTER, 10-SAD]` |
| 12 | LLD | `[01-PROJECT-CHARTER, 10-SAD, 11-HLD]` |
| 13 | TEST-PLAN | `[01-PROJECT-CHARTER, 03-SRS, 10-SAD, 12-LLD]` |
| 14 | TEST-CASES | `[01-PROJECT-CHARTER, 03-SRS, 13-TEST-PLAN]` |
| 15 | Relatório de Qualidade | `[13-TEST-PLAN, 14-TEST-CASES]` |
| 16 | DEPLOYMENT-PLAN | `[01-PROJECT-CHARTER, 10-SAD, 11-HLD, 12-LLD]` |
| 17 | Manuais de Usuário | `[01-PROJECT-CHARTER, 03-SRS]` |
| 18 | Manuais Operacionais | `[01-PROJECT-CHARTER, 10-SAD, 16-DEPLOYMENT-PLAN]` |
| 19 | Termo de Aceite | `[01-PROJECT-CHARTER, 13-TEST-PLAN, 15-RELATORIO-QUALIDADE]` |
| 20 | Lições Aprendidas | `[todos os 19 documentos anteriores]` |

---

## EFEITOS CASCATA

Quando um documento já em COMPLIANCE é modificado:

| Se modificar... | Regenerar e revalidar... |
|---|---|
| 01-PROJECT-CHARTER | Documentos 2-20 (19 docs) |
| 02-BRD | Documentos 3-20 (18 docs downstream do BRD) |
| 10-SAD | Documentos 11-20 (10 docs downstream do SAD) |
| (regra geral) | Todos os documentos com numeração maior que dependem do modificado |

**Ação:** Alertar o humano, listar documentos afetados, perguntar: (A) regeneração completa, ou (B) marcar como "potencialmente desatualizados."

---

## FINALIZAÇÃO — GIT WORKFLOW

Quando os 20 documentos estiverem em COMPLIANCE e o humano confirmar a conclusão:

### Passo F.1 — Git Add e Commit

```bash
git add -A
git commit -m "docs: documentação WATERFALL completa — ${PROJECT_ID_NAME}

- 20 documentos WATERFALL gerados e validados
- Status: COMPLIANCE em todos os documentos
- Gerado pelo Waterfall Orchestrator v1.0
- Branch: ${PROMPT_BRANCH}

Co-Authored-By: Claude <noreply@anthropic.com>"
```

### Passo F.2 — Git Push

```bash
git push origin ${PROMPT_BRANCH}
```

Se falhar (branch remota existe): perguntar sobre `--force`.

### Passo F.3 — Criar e Mergear PR

```bash
gh pr create --base main --head ${PROMPT_BRANCH} --title "docs: documentação WATERFALL — ${PROJECT_ID_NAME}" --body "Documentação WATERFALL completa para ${PROJECT_ID_NAME}. 20 documentos validados."
gh pr merge --merge --delete-branch
```

### Passo F.4 — Cleanup Local

```bash
git checkout main
git branch -d ${PROMPT_BRANCH}
```

**Regra de segurança:** Nunca deletar branch local se o merge falhar.

---

## LOCALIZAÇÃO DOS PROMPTS

```
.specs/prompts/project-documents-waterfall/
├── PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md   ← ESTE ORQUESTRADOR
├── PROMPT-GENERATE-PROJECT-CHARTER.md
├── PROMPT-GATE-PROJECT-CHARTER.md
├── PROMPT-FIX-PROJECT-CHARTER.md
├── PROMPT-GENERATE-BRD.md
├── PROMPT-GATE-BRD.md
├── PROMPT-FIX-BRD.md
├── (total de 61 prompts)
└── PROMPT-FIX-LICOES-APRENDIDAS.md
```

---

🤖 *Roadmap gerado pelo Waterfall Orchestrator v1.0. Skills: draft-project-charter, senior-pm, brainstorming.*
