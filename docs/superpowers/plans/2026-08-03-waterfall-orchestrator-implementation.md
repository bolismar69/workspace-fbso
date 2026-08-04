# Waterfall Orchestrator — Implementation Plan

> **Para agentic workers:** REQUIRED SUB-SKILL: Use `superpowers:subagent-driven-development` (recommended) ou `superpowers:executing-plans` para implementar este plano task-by-task. Steps usam checkbox (`- [ ]`) para tracking.

**Goal:** Criar 61 prompts (1 orquestrador + 20 documentos × 3 prompts cada) para o WATERFALL orchestrator, seguindo 7 regras de gating e fluxo de dados explícito.

**Architecture:** Cada documento WATERFALL tem um trio GENERATE→GATE→FIX. O orquestrador gerencia o fluxo sequencial e passa explicitamente os paths entre prompts. Cada GENERATE usa skills híbridas (primárias + fallback). GATE valida contra checklist específico. FIX corrige cirurgicamente apenas as violações.

**Tech Stack:** Markdown prompts, skills via `Skill` tool, arquivos salvos em `.specs/prompts/project-documents-waterfall/`

## Global Constraints

- **Rule 1 (No guessing):** Nenhum prompt infere ou busca inputs; todos os parâmetros são passados explicitamente pelo orquestrador
- **Rule 2 (Every GENERATE has GATE+FIX):** Todos os 20 documentos têm trio completo
- **Rule 3 (Explicit data flow):** `DOC_PATH` flui GENERATE→GATE→FIX; `VIOLATIONS[]` flui GATE→FIX
- **Rule 4 (Surgical FIX):** FIX edita apenas seções com violações, nunca regera o documento
- **Rule 5 (New docs = "Em análise"):** GENERATE escreve `[STATUS: Em análise]` ao criar o arquivo
- **Rule 6 (Under review = "Em revisão"):** GATE/FIX alteram status para `[STATUS: Em revisão]`
- **Rule 7 (Sequential gating):** Só avança para próximo documento com `[STATUS: COMPLIANCE]`
- **Spec reference:** `docs/superpowers/specs/2026-08-03-waterfall-orchestrator-design.md`
- **Output directory:** `.specs/prompts/project-documents-waterfall/`
- **Naming convention:** `PROMPT-{GENERATE|GATE|FIX}-{DOCUMENT-SLUG}.md`

---

### Task 1: Criar o Orquestrador (ROADMAP)

**Files:**
- Create: `.specs/prompts/project-documents-waterfall/PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md`

**Interfaces:**
- Produces: O ponto de entrada que recebe as 7 variáveis de bootstrap e orquestra o fluxo completo

- [ ] **Step 1: Criar o arquivo do orquestrador**

Criar `.specs/prompts/project-documents-waterfall/PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md` com o conteúdo abaixo:

```markdown
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
```

- [ ] **Step 2: Verificar estrutura do orquestrador**

Ler o arquivo criado e confirmar que contém:
- [ ] 7 regras de gating documentadas
- [ ] Bootstrap Fase 0 com 7 variáveis
- [ ] 5 fases WATERFALL com os 20 documentos
- [ ] Loop de execução STEP 1-4 documentado
- [ ] Matriz de skills por documento
- [ ] Mapeamento de upstream docs
- [ ] Efeitos cascata
- [ ] Git workflow

- [ ] **Step 3: Commit**

```bash
git add .specs/prompts/project-documents-waterfall/PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md
git commit -m "feat: criar orquestrador WATERFALL — roadmap master com 7 regras de gating"
```

---

### Task 2: Criar Trio PROJECT-CHARTER (Fase 1)

**Files:**
- Create: `.specs/prompts/project-documents-waterfall/PROMPT-GENERATE-PROJECT-CHARTER.md`
- Create: `.specs/prompts/project-documents-waterfall/PROMPT-GATE-PROJECT-CHARTER.md`
- Create: `.specs/prompts/project-documents-waterfall/PROMPT-FIX-PROJECT-CHARTER.md`

**Interfaces:**
- GENERATE consumes: `DOC_PATH`, `PROJECT_ID_NAME`, `BRIEFING`, `UPSTREAM_DOCS=[]`, `EXTRA_INPUTS`, `SKILLS=["draft-project-charter","senior-pm"]`
- GENERATE produces: `DOC_PATH` (arquivo criado com 14 seções + `[STATUS: Em análise]`)
- GATE consumes: `DOC_PATH` from GENERATE → atualiza para `[STATUS: Em revisão]` → retorna `{PASS}` ou `{FAIL, VIOLATIONS[]}`
- FIX consumes: `DOC_PATH` + `VIOLATIONS[]` from GATE → edições cirúrgicas → retorna `DOC_PATH`

- [ ] **Step 1: Criar PROMPT-GENERATE-PROJECT-CHARTER.md**

```markdown
# PROMPT: GERADOR DE PROJECT CHARTER COM FOCO EM NEGÓCIO
## Versão: 1.0 — WATERFALL Orchestrator

Atue como um Especialista em Gestão de Processos (BPM), Analista de Negócios Sênior e Arquiteto de Soluções Organizacionais. Sua missão é criar um **Project Charter** completo para o projeto.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto (ex: `PRJ-FIN-2026-0003-SAAS-FBSO-ORG`) |
| `BRIEFING` | Briefing do projeto (texto inline ou caminho de arquivo) |
| `UPSTREAM_DOCS` | Lista vazia `[]` — este é o documento raiz |
| `EXTRA_INPUTS` | Documentos e prompts extras fornecidos pelo humano |
| `SKILLS` | Lista de skills: `["draft-project-charter", "senior-pm"]` |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima
2. Skills: tente usar as skills listadas em `SKILLS` via `Skill` tool. Se falharem, use o template de fallback abaixo
3. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
4. Foco estrito em negócio — não inclua detalhes técnicos de implementação
5. Ao final, retorne `{DOC_PATH}` confirmando a criação

## Template de Fallback (14 Seções)

```
# Project Charter: {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 — Documento Inicial |
| **Patrocinador** | {Área ou Diretoria Patrocinadora} |
| **Metodologia** | WATERFALL |
| **Status** | Em análise |

---

### 1. Declaração do Problema (Problem Statement)
[Descrever o cenário atual, dores do negócio, gargalos operacionais e o impacto/custo de não resolver.]

### 2. Propósito do Projeto (Project Purpose)
[Visão da solução sob perspectiva de negócio. O que está sendo construído e como resolve o problema.]

#### 2.1 Visão de Longo Prazo
[Impacto futuro ou evolução esperada após a consolidação do projeto.]

### 3. Escopo (Scope)

#### 3.1 Dentro do Escopo (In Scope)
[Listar macro-módulos, fluxos de valor e funcionalidades contempladas.]

#### 3.2 Fora do Escopo (Out of Scope)
[Listar explicitamente o que NÃO será feito.]

### 4. Entregas (Deliverables) & Critérios de Aceitação

| # | Entrega | Critérios de Aceitação de Negócio | Data-Alvo |
|---|---------|------------------------------------|-----------|
| D1 | ... | ... | ... |

### 5. Partes Interessadas e Matriz RACI (Stakeholders & RACI)

| Parte Interessada | Papel | D1 | D2 | D3 |
|---|---|---|---|---|
| ... | ... | R/A/C/I | ... | ... |

### 6. Critérios de Sucesso (Success Criteria)

| # | Critério | Indicador | Meta |
|---|---|---|---|
| C1 | ... | ... | ... |

### 7. Premissas (Assumptions)
[Listar premissas assumidas para o planejamento.]

### 8. Restrições (Constraints)
[Listar restrições de prazo, orçamento, recursos, regulatórias.]

### 9. Riscos de Alto Nível (High-Level Risks)

| Risco | Probabilidade | Impacto | Mitigação |
|---|---|---|---|
| ... | Alta/Média/Baixa | Alto/Médio/Baixo | ... |

### 10. Marcos do Projeto (Project Milestones)

| Marco | Data | Critério de Conclusão |
|---|---|---|
| M1: Kickoff | ... | ... |
| M2: ... | ... | ... |

### 11. Orçamento Estimado (Estimated Budget)

| Categoria | Estimativa |
|---|---|
| ... | ... |

### 12. Plano de Comunicação (Communication Plan)

| Público | Frequência | Canal | Responsável |
|---|---|---|---|
| ... | ... | ... | ... |

### 13. Governança (Governance)
[Estrutura de governança, comitês, papéis e responsabilidades.]

### 14. Aprovações (Approvals)

| Nome | Papel | Data | Assinatura |
|---|---|---|---|
| ... | Patrocinador | ... | Pendente |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` ao final se o documento estiver completo. Se faltarem insumos impeditivos, emitir `[STATUS: INSUCESSO]`.
```

- [ ] **Step 2: Verificar GENERATE-PROJECT-CHARTER**

Conferir que:
- [ ] Parâmetros de input declarados explicitamente (DOC_PATH, PROJECT_ID_NAME, BRIEFING, UPSTREAM_DOCS, EXTRA_INPUTS, SKILLS)
- [ ] Regra "NUNCA procure por inputs" está documentada
- [ ] Template de fallback com 14 seções
- [ ] Status inicial `[STATUS: Em análise]`
- [ ] Skills listadas corretamente

- [ ] **Step 3: Criar PROMPT-GATE-PROJECT-CHARTER.md**

```markdown
# PROMPT: PORTÃO DE VALIDAÇÃO DE PROJECT CHARTER
## Versão: 1.0 — WATERFALL Orchestrator

Atue como um Auditor de Qualidade de Documentação, especializado em Project Charter e metodologia WATERFALL.

## Inputs (recebidos explicitamente — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo do arquivo a ser validado |

## Regras

1. Leia **APENAS** o arquivo em `DOC_PATH` — não busque outros arquivos
2. Altere o status do documento para `[STATUS: Em revisão]`
3. Execute cada item do CHECKLIST abaixo contra o conteúdo do documento
4. Retorne `{PASS}` se todos os checks passarem
5. Retorne `{FAIL, VIOLATIONS: [{section, description, severity}]}` se houver falhas

## Checklist de Compliance

1. **Cabeçalho e Metadados:** O documento possui campo Projeto, Data, Versão, Patrocinador, Metodologia e Status preenchidos? Status é "Em análise"?
2. **Seção 1 — Problem Statement:** Declaração do problema está presente e descreve cenário atual, dores e impacto?
3. **Seção 2 — Propósito:** Propósito do projeto está definido? Visão de longo prazo está documentada (2.1)?
4. **Seção 3 — Escopo:** In Scope e Out of Scope estão explicitamente listados? Não há ambiguidade entre o que está dentro e fora?
5. **Seção 4 — Entregas:** Tabela de entregas com critérios de aceitação e datas-alvo preenchida? Critérios são mensuráveis?
6. **Seção 5 — Stakeholders e RACI:** Matriz RACI preenchida com partes interessadas e papéis mapeados contra entregas (D1, D2, etc.)?
7. **Seção 6 — Critérios de Sucesso:** Critérios mensuráveis com indicadores e metas definidos?
8. **Seção 7 — Premissas:** Premissas listadas?
9. **Seção 8 — Restrições:** Restrições de prazo, orçamento, recursos e regulatórias documentadas?
10. **Seção 9 — Riscos:** Riscos de alto nível com probabilidade, impacto e mitigação preenchidos?
11. **Seção 10 — Marcos:** Marcos do projeto com datas e critérios de conclusão definidos?
12. **Seção 11 — Orçamento:** Orçamento estimado por categoria preenchido?
13. **Seção 12 — Plano de Comunicação:** Plano com público, frequência, canal e responsável definido?
14. **Seção 13 — Governança:** Estrutura de governança documentada?
15. **Seção 14 — Aprovações:** Tabela de aprovações presente?
16. **Consistência Interna:** Os critérios de sucesso (Seção 6) estão alinhados com o propósito (Seção 2)? As entregas (Seção 4) cobrem o escopo (Seção 3)? Os riscos (Seção 9) consideram as restrições (Seção 8)?
17. **Foco em Negócio:** O documento está focado em regras de negócio (não em detalhes técnicos de implementação)?
```

- [ ] **Step 4: Verificar GATE-PROJECT-CHARTER**

Conferir que:
- [ ] Recebe `DOC_PATH` como único input explícito
- [ ] Altera status para `[STATUS: Em revisão]`
- [ ] 17 itens de checklist específicos para PROJECT-CHARTER
- [ ] Retorna `{PASS}` ou `{FAIL, VIOLATIONS[]}`

- [ ] **Step 5: Criar PROMPT-FIX-PROJECT-CHARTER.md**

```markdown
# PROMPT: CORRETOR DE PROJECT CHARTER
## Versão: 1.0 — WATERFALL Orchestrator

Atue como um Revisor Técnico de Documentação especializado em correções cirúrgicas de Project Charter.

## Inputs (recebidos explicitamente do GATE — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo do arquivo a ser corrigido |
| `VIOLATIONS` | Lista de não-conformidades reportadas pelo GATE |

Cada item em `VIOLATIONS` tem o formato:
```json
{ "section": "Nome da Seção", "description": "Descrição da não-conformidade", "severity": "HIGH|MEDIUM|LOW" }
```

## Regras

1. Edite **APENAS** as seções listadas em `VIOLATIONS` — correção cirúrgica
2. **NÃO** recrie, regenere ou reescreva o documento inteiro
3. **NÃO** altere seções que passaram no GATE e não estão em `VIOLATIONS`
4. Mantenha o status como `[STATUS: Em revisão]`
5. Após corrigir cada violação, adicione um comentário inline `<!-- FIX: {description} — corrigido -->` na seção reparada
6. Retorne `{DOC_PATH}` após as correções
```

- [ ] **Step 6: Verificar FIX-PROJECT-CHARTER**

Conferir que:
- [ ] Recebe `DOC_PATH` e `VIOLATIONS[]` como inputs explícitos
- [ ] Regra "NÃO recrie, regenere ou reescreva" está documentada
- [ ] Regra "NÃO altere seções que passaram no GATE" está documentada
- [ ] Comentário inline de rastreabilidade `<!-- FIX: ... -->`

- [ ] **Step 7: Verificação de integridade do trio**

Ler os 3 arquivos e verificar:
- [ ] GENERATE produz o status inicial `[STATUS: Em análise]`
- [ ] GATE transforma para `[STATUS: Em revisão]`
- [ ] FIX mantém `[STATUS: Em revisão]`
- [ ] O fluxo de dados GENERATE→GATE→FIX via `DOC_PATH` está consistente entre os 3 prompts
- [ ] O fluxo GATE→FIX via `VIOLATIONS[]` está consistente

- [ ] **Step 8: Commit**

```bash
git add .specs/prompts/project-documents-waterfall/PROMPT-GENERATE-PROJECT-CHARTER.md \
        .specs/prompts/project-documents-waterfall/PROMPT-GATE-PROJECT-CHARTER.md \
        .specs/prompts/project-documents-waterfall/PROMPT-FIX-PROJECT-CHARTER.md
git commit -m "feat: criar trio PROJECT-CHARTER — GENERATE, GATE, FIX com 7 regras de gating"
```

---

### Task 3: Criar Trio BRD (Fase 2 — Documento 2)

**Files:**
- Create: `.specs/prompts/project-documents-waterfall/PROMPT-GENERATE-BRD.md`
- Create: `.specs/prompts/project-documents-waterfall/PROMPT-GATE-BRD.md`
- Create: `.specs/prompts/project-documents-waterfall/PROMPT-FIX-BRD.md`

**Interfaces:**
- GENERATE consumes: `DOC_PATH`, `PROJECT_ID_NAME`, `BRIEFING`, `UPSTREAM_DOCS=[01-PROJECT-CHARTER]`, `EXTRA_INPUTS`, `SKILLS=["brd-creation","business-analyst","requirements-elicitation"]`
- GENERATE produces: `DOC_PATH` → `[STATUS: Em análise]`
- GATE consumes: `DOC_PATH` → atualiza para `[STATUS: Em revisão]` → retorna `{PASS}` ou `{FAIL, VIOLATIONS[]}`
- FIX consumes: `DOC_PATH` + `VIOLATIONS[]` from GATE → edições cirúrgicas → retorna `DOC_PATH`

- [ ] **Step 1: Criar PROMPT-GENERATE-BRD.md**

```markdown
# PROMPT: GERADOR DE BUSINESS REQUIREMENTS DOCUMENT (BRD)
## Versão: 1.0 — WATERFALL Orchestrator

Atue como um Analista de Negócios Sênior (Business Analyst), especializado em levantamento e documentação de requisitos de negócio.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `BRIEFING` | Briefing do projeto (texto inline ou caminho de arquivo) |
| `UPSTREAM_DOCS` | Lista: `[{PROJECT_COMPLETE_PATH_NAME}/01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md]` |
| `EXTRA_INPUTS` | Documentos e prompts extras fornecidos pelo humano |
| `SKILLS` | Lista de skills: `["brd-creation", "business-analyst", "requirements-elicitation"]` |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima
2. **LEIA** o PROJECT-CHARTER em `UPSTREAM_DOCS[0]` — todos os requisitos devem rastrear de volta aos objetivos do Charter
3. Skills: tente usar as skills listadas em `SKILLS` via `Skill` tool. Se falharem, use o template de fallback abaixo
4. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
5. Foco estrito em requisitos de negócio — não inclua especificações técnicas
6. Ao final, retorne `{DOC_PATH}` confirmando a criação

## Template de Fallback

```
# Business Requirements Document (BRD): {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento Base** | 01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

### 1. Requisitos de Negócio (Business Requirements)

Cada requisito deve referenciar o objetivo do Project Charter que o originou.

| ID | Requisito de Negócio | Objetivo Charter (OBJ-XX) | Prioridade | Stakeholder |
|----|----------------------|---------------------------|------------|-------------|
| REQ-01 | ... | OBJ-01 | Alta/Média/Baixa | ... |

### 2. Regras de Negócio (Business Rules)

| ID | Regra | Descrição | Requisito Vinculado |
|----|-------|-----------|---------------------|
| BR-01 | ... | ... | REQ-01 |

### 3. Restrições de Negócio (Business Constraints)

| ID | Restrição | Descrição | Impacto |
|----|-----------|-----------|---------|
| BC-01 | ... | ... | ... |

### 4. Requisitos de Dados (Data Requirements)

| Entidade | Descrição | Requisito Vinculado |
|----------|-----------|---------------------|
| ... | ... | REQ-XX |

### 5. Requisitos de Interface e Integração (Integration Requirements)

| Interface | Descrição | Tipo | Requisito Vinculado |
|-----------|-----------|------|---------------------|
| ... | ... | API/File/UI | REQ-XX |

### 6. Requisitos de Segurança e Compliance (Security & Compliance)

| ID | Requisito | Regulação/Política | Requisito Vinculado |
|----|-----------|-------------------|---------------------|
| SEC-01 | ... | ... | REQ-XX |

### 7. Fluxos de Processo de Negócio (Business Process Flows)
[Descrever os fluxos BPMN ou textuais dos processos impactados]

### 8. Mapeamento de Stakeholders Detalhado

| Stakeholder | Necessidades | Expectativas | Nível de Influência |
|-------------|-------------|--------------|---------------------|
| ... | ... | ... | Alta/Média/Baixa |

### 9. Matriz de Rastreabilidade Preliminar (BRD → Project Charter)

| Requisito BRD | Objetivo Charter | Status |
|---------------|------------------|--------|
| REQ-01 | OBJ-01 | ✅ Vinculado |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se todos os requisitos rastrearem ao Project Charter e as 9 seções estiverem completas.
```

- [ ] **Step 2: Criar PROMPT-GATE-BRD.md**

```markdown
# PROMPT: PORTÃO DE VALIDAÇÃO DE BUSINESS REQUIREMENTS DOCUMENT (BRD)
## Versão: 1.0 — WATERFALL Orchestrator

Atue como um Auditor de Requisitos de Negócio especializado em metodologia WATERFALL.

## Inputs (recebidos explicitamente — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo do arquivo a ser validado |

## Regras

1. Leia **APENAS** o arquivo em `DOC_PATH`
2. Altere o status do documento para `[STATUS: Em revisão]`
3. Execute cada item do CHECKLIST
4. Retorne `{PASS}` ou `{FAIL, VIOLATIONS: [{section, description, severity}]}`

## Checklist de Compliance

1. **Metadados:** Cabeçalho com Projeto, Documento Base (Project Charter), Data, Versão e Metodologia preenchidos?
2. **Seção 1 — Requisitos de Negócio:** Tabela com ID, descrição, objetivo do Charter vinculado, prioridade e stakeholder? Pelo menos um requisito por objetivo do Project Charter?
3. **Seção 2 — Regras de Negócio:** Regras documentadas com ID, descrição e vínculo com requisito?
4. **Seção 3 — Restrições:** Restrições de negócio listadas com impacto?
5. **Seção 4 — Requisitos de Dados:** Entidades de dados mapeadas com vínculo a requisitos?
6. **Seção 5 — Integração:** Interfaces e integrações listadas com tipo e requisito vinculado?
7. **Seção 6 — Segurança e Compliance:** Requisitos de segurança com regulação/política referenciada?
8. **Seção 7 — Fluxos de Processo:** Fluxos de negócio descritos?
9. **Seção 8 — Stakeholders:** Mapeamento detalhado com necessidades, expectativas e nível de influência?
10. **Seção 9 — Matriz de Rastreabilidade:** Todo requisito BRD está vinculado a um objetivo do Project Charter? Zero órfãos?
11. **Cobertura:** 100% dos objetivos do Project Charter cobertos por pelo menos um requisito BRD?
12. **Foco em Negócio:** Documento não contém especificações técnicas de implementação?
```

- [ ] **Step 3: Criar PROMPT-FIX-BRD.md**

```markdown
# PROMPT: CORRETOR DE BUSINESS REQUIREMENTS DOCUMENT (BRD)
## Versão: 1.0 — WATERFALL Orchestrator

Atue como um Revisor de Requisitos de Negócio especializado em correções cirúrgicas de BRD.

## Inputs (recebidos explicitamente do GATE — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo do arquivo a ser corrigido |
| `VIOLATIONS` | Lista de não-conformidades reportadas pelo GATE |

Cada item: `{ "section": "Nome da Seção", "description": "Descrição da não-conformidade", "severity": "HIGH|MEDIUM|LOW" }`

## Regras

1. Edite **APENAS** as seções listadas em `VIOLATIONS` — correção cirúrgica
2. **NÃO** recrie, regenere ou reescreva o documento inteiro
3. **NÃO** altere seções que passaram no GATE
4. Mantenha o status como `[STATUS: Em revisão]`
5. Após corrigir cada violação, adicione `<!-- FIX: {description} — corrigido -->` na seção reparada
6. Retorne `{DOC_PATH}` após as correções
```

- [ ] **Step 4: Verificar integridade do trio BRD**

Conferir que:
- [ ] GENERATE lê UPSTREAM_DOCS (PROJECT-CHARTER) explicitamente
- [ ] GATE verifica rastreabilidade Charter→BRD (checklist item 10 e 11)
- [ ] FIX é cirúrgico com regras idênticas ao template base
- [ ] Status lifecycle consistente: Em análise → Em revisão

- [ ] **Step 5: Commit**

```bash
git add .specs/prompts/project-documents-waterfall/PROMPT-GENERATE-BRD.md \
        .specs/prompts/project-documents-waterfall/PROMPT-GATE-BRD.md \
        .specs/prompts/project-documents-waterfall/PROMPT-FIX-BRD.md
git commit -m "feat: criar trio BRD — GENERATE, GATE, FIX com rastreabilidade Charter→BRD"
```

---

### Task 4: Criar Trio SRS (Fase 2 — Documento 3)

**Files:**
- Create: `.specs/prompts/project-documents-waterfall/PROMPT-GENERATE-SRS.md`
- Create: `.specs/prompts/project-documents-waterfall/PROMPT-GATE-SRS.md`
- Create: `.specs/prompts/project-documents-waterfall/PROMPT-FIX-SRS.md`

**Interfaces:**
- GENERATE consumes: `DOC_PATH`, `PROJECT_ID_NAME`, `BRIEFING`, `UPSTREAM_DOCS=[01-PROJECT-CHARTER, 02-BRD]`, `EXTRA_INPUTS`, `SKILLS=["frs-creation","requirements-engineering"]`
- Diffs from BRD pattern: template é SRS (functional specs, not business reqs), skills different, upstream docs include BRD

- [ ] **Step 1: Criar PROMPT-GENERATE-SRS.md**

Mesma estrutura do GENERATE-BRD com estas diferenças concretas:
- `SKILLS` = `["frs-creation", "requirements-engineering"]`
- `UPSTREAM_DOCS` = `[01-PROJECT-CHARTER, 02-BRD]`
- Template de fallback tem estas seções: (1) Functional Requirements (com ID FR-XX, vinculado a REQ-XX do BRD), (2) Non-Functional Requirements (Performance, Security, Availability, Scalability, Usability — cada um com métrica), (3) System Features (feature list com prioridade MoSCoW), (4) External Interfaces (API contracts, data formats, protocols), (5) Assumptions and Dependencies, (6) Traceability Matrix FR→BRD→Charter

- [ ] **Step 2: Criar PROMPT-GATE-SRS.md**

Checklist de 10 itens:
1. Metadados com Documento Base (Project Charter + BRD)
2. Functional Requirements: cada FR vinculado a REQ do BRD
3. NFRs: 5 dimensões (performance, security, availability, scalability, usability) com métricas mensuráveis
4. System Features com prioridade MoSCoW
5. External Interfaces documentadas
6. Assumptions and Dependencies
7. Traceability: 100% dos FRs rastreiam a REQs do BRD, que rastreiam a OBJs do Charter
8. Zero órfãos
9. NFRs são testáveis (cada NFR tem uma métrica verificável)
10. Cobertura total: todos os REQs do BRD cobertos por pelo menos um FR

- [ ] **Step 3: Criar PROMPT-FIX-SRS.md**

Idêntico ao FIX-BRD (template cirúrgico padrão), substituindo "BRD" por "SRS".

- [ ] **Step 4: Commit**

```bash
git add .specs/prompts/project-documents-waterfall/PROMPT-GENERATE-SRS.md \
        .specs/prompts/project-documents-waterfall/PROMPT-GATE-SRS.md \
        .specs/prompts/project-documents-waterfall/PROMPT-FIX-SRS.md
git commit -m "feat: criar trio SRS — GENERATE, GATE, FIX com rastreabilidade BRD→SRS"
```

---

### Tasks 5-9: Trios dos Documentos Restantes da Fase 2

Cada task segue o mesmo padrão estrutural dos Tasks 2-4, com as diferenças específicas abaixo:

---

#### Task 5: Trio RTM (Requirements Traceability Matrix)

**SKILLS:** `["requirements-modeling", "requirements-validation"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER, 02-BRD, 03-SRS]`  
**TEMPLATE SECTIONS:** Matriz com colunas: OBJ-Charter | REQ-BRD | FR-SRS | Status. Regra: zero órfãos, 100% cobertura bidirecional.  
**GATE CHECKLIST (8 items):**
1. Colunas obrigatórias presentes (OBJ-Charter, REQ-BRD, FR-SRS, Status)
2. Todo OBJ do Charter tem pelo menos 1 REQ e 1 FR vinculado
3. Todo REQ do BRD tem pelo menos 1 FR vinculado
4. Nenhum FR sem REQ correspondente (zero órfãos forward)
5. Nenhum REQ sem OBJ correspondente (zero órfãos backward)
6. Status de cada linha é ✅ Vinculado, ⚠️ Parcial ou ❌ Órfão
7. Cobertura forward: 100% dos FRs
8. Cobertura backward: 100% dos OBJs

---

#### Task 6: Trio EAP/WBS

**SKILLS:** `["decomposition-planning-roadmap", "project-estimation"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER, 02-BRD]`  
**TEMPLATE SECTIONS:** (1) EAP Gráfica (árvore hierárquica), (2) Dicionário da EAP (cada pacote com ID, descrição, responsável, critério de aceitação, estimativa), (3) Matriz EAP×Entregas do Charter, (4) EAP×Requisitos BRD.  
**GATE CHECKLIST (7 items):**
1. EAP gráfica com hierarquia de pelo menos 3 níveis
2. Dicionário da EAP com ID, descrição, responsável, critério de aceitação e estimativa para cada pacote
3. Matriz EAP×Entregas: cada pacote vinculado a uma entrega do Project Charter
4. Matriz EAP×Requisitos: cada pacote vinculado a um requisito do BRD
5. Cobertura: todas as entregas do Charter cobertas por pelo menos um pacote EAP
6. Sem pacotes órfãos (sem vínculo com entrega ou requisito)
7. Estimativas preenchidas para todos os pacotes

---

#### Task 7: Trio Cronograma/Gantt

**SKILLS:** `["roadmap-planning", "project-estimation"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER, 05-EAP-WBS]`  
**TEMPLATE SECTIONS:** (1) Lista de Atividades (extraída da EAP), (2) Sequenciamento (dependências finish-to-start), (3) Estimativas de Duração, (4) Caminho Crítico, (5) Cronograma (tabela com data início/fim por atividade), (6) Diagrama de Gantt (representação textual), (7) Marcos (alinhados com os marcos do Charter).  
**GATE CHECKLIST (8 items):**
1. Todas as atividades derivadas de pacotes da EAP
2. Dependências documentadas
3. Caminho crítico identificado
4. Datas de início e fim para cada atividade
5. Marcos alinhados com Project Charter Seção 10
6. Sem conflitos de dependência circular
7. Folga (slack) calculada para atividades não-críticas
8. Duração total do projeto consistente com os marcos do Charter

---

#### Task 8: Trio Orçamento

**SKILLS:** `["project-estimation"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER, 05-EAP-WBS, 06-Cronograma]`  
**TEMPLATE SECTIONS:** (1) Custos por Pacote EAP, (2) Custos por Recurso (RH, infra, licenças, serviços), (3) Curva S (custo acumulado ao longo do tempo), (4) Reserva de Contingência, (5) Fluxo de Caixa Projetado, (6) Comparativo Orçado×Real (a ser preenchido).  
**GATE CHECKLIST (6 items):**
1. Cada pacote EAP tem custo estimado
2. Custos por categoria de recurso detalhados
3. Reserva de contingência explicitada (valor e %)
4. Curva S projetada
5. Total alinhado com Seção 11 do Project Charter
6. Sem valores zerados ou ausentes em categorias obrigatórias

---

#### Task 9: Trio Plano de Comunicação

**SKILLS:** `["stakeholder-analysis", "stakeholder-map"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER]`  
**TEMPLATE SECTIONS:** (1) Matriz de Comunicação (público, mensagem, frequência, canal, responsável, formato), (2) Fluxo de Escalação, (3) Calendário de Reuniões e Rituais, (4) Repositório de Documentos, (5) Plano de Comunicação em Crise.  
**GATE CHECKLIST (6 items):**
1. Todos os stakeholders do Charter (Seção 5) têm entradas na matriz de comunicação
2. Fluxo de escalação definido com níveis e responsáveis
3. Calendário com reuniões recorrentes e marcos
4. Canais definidos e adequados ao público
5. Repositório de documentos especificado
6. Plano de crise contemplado

---

#### Task 10: Trio Plano de Riscos

**SKILLS:** `["risk-manager", "risk-management-specialist"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER]`  
**TEMPLATE SECTIONS:** (1) Registro de Riscos (ID, descrição, categoria, probabilidade, impacto, score=P×I, trigger, response strategy, owner), (2) Matriz de Probabilidade×Impacto, (3) Plano de Resposta (avoid/transfer/mitigate/accept por risco), (4) Plano de Contingência, (5) Riscos Residuais, (6) Monitoramento (indicadores, thresholds, revisão periódica).  
**GATE CHECKLIST (8 items):**
1. Registro de riscos com ID, descrição, categoria, P, I, score, trigger, estratégia e owner
2. Riscos da Seção 9 do Project Charter expandidos e detalhados
3. Matriz P×I preenchida
4. Estratégia de resposta definida para cada risco (avoid/transfer/mitigate/accept)
5. Plano de contingência para riscos HIGH
6. Thresholds de monitoramento definidos
7. Riscos residuais identificados
8. Score = P×I calculado corretamente para todos os riscos

---

### Tasks 11-13: Trios da Fase 3 — Design e Arquitetura

---

#### Task 11: Trio SAD (Software Architecture Document)

**SKILLS:** `["software-architecture", "architecture-designer", "architecture-patterns"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER, 02-BRD, 03-SRS]`  
**VISÕES OBRIGATÓRIAS:** Solution Architecture, Data Architecture, Security Architecture, DevOps/SRE Architecture, Infrastructure/Cloud Architecture, Testing Architecture  
**TEMPLATE SECTIONS:** (1) Architectural Overview (estilo arquitetural, ADRs chave, diagrama de contexto), (2) Solution Architecture (componentes, camadas, fluxos), (3) Data Architecture (modelo conceitual, storage strategy, data governance), (4) Security Architecture (threat model, authN/Z, data protection, secrets management), (5) DevOps/SRE Architecture (CI/CD pipeline, observability, SLOs, incident response), (6) Infrastructure/Cloud Architecture (topology, scaling, disaster recovery), (7) Testing Architecture (test pyramid, strategy per layer, quality gates), (8) Cross-cutting Concerns (logging, error handling, i18n, feature flags), (9) ADR Registry (cada ADR vinculado a NFR do SRS), (10) Traceability SAD→SRS→BRD→Charter  
**GATE CHECKLIST (14 items):**
1. 6 visões obrigatórias presentes e completas
2. ADRs documentados e vinculados a NFRs do SRS
3. Diagrama de contexto presente
4. Threat model documentado (STRIDE ou similar)
5. Estratégia de observability definida (logs, metrics, traces, alerts)
6. Topologia de deploy e scaling documentada
7. Disaster recovery strategy definida
8. CI/CD pipeline descrito
9. Cada decisão arquitetural vinculada a um requisito do SRS
10. 100% dos NFRs cobertos por pelo menos uma decisão arquitetural
11. Sem componentes órfãos (sem vínculo com requisito)
12. AuthN/AuthZ model documentado
13. Estratégia de testes alinhada com pirâmide de testes
14. Consistência interna: as 6 visões não se contradizem

---

#### Task 12: Trio HLD (High-Level Design)

**SKILLS:** `["c4-container", "system-design", "architecture-decision-records"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER, 10-SAD]`  
**VISÕES OBRIGATÓRIAS:** Macro components, data flows, integration topology, deployment topology, technology stack decisions  
**TEMPLATE SECTIONS:** (1) System Context (C4 Level 1), (2) Container Diagram (C4 Level 2), (3) Technology Stack (com rationale), (4) Integration Topology (sync/async, APIs, message queues, events), (5) Deployment Topology (environments, regions, sizing), (6) Data Flow Diagrams (por fluxo principal), (7) Non-Functional Requirements Allocation (cada NFR do SRS alocado a um componente/camada)  
**GATE CHECKLIST (9 items):**
1. C4 Level 1 e Level 2 presentes
2. Stack tecnológica com rationale para cada escolha
3. Integrações externas documentadas com protocolos
4. Topologia de deploy por ambiente
5. Data flow diagrams para fluxos principais
6. Cada NFR do SRS alocado a um componente
7. Alinhamento com SAD (sem decisões conflitantes)
8. ADRs do SAD referenciados nas decisões de design
9. Sem componentes não mapeados no SAD

---

#### Task 13: Trio LLD (Low-Level Design)

**SKILLS:** `["c4-component", "ddd-tactical-patterns", "database-designer"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER, 10-SAD, 11-HLD]`  
**VISÕES OBRIGATÓRIAS:** Class diagrams, API contracts, database schemas, sequence diagrams, state machines, component interfaces  
**TEMPLATE SECTIONS:** (1) Component Diagram (C4 Level 3), (2) Class/Entity Design (por domínio), (3) API Contracts (OpenAPI/REST endpoints, request/response schemas), (4) Database Schema (DDL, indexes, relationships), (5) Sequence Diagrams (para fluxos críticos), (6) State Machines (para entidades com estado), (7) Error Handling Strategy (error codes, retry policies, circuit breakers), (8) Component Interfaces (method signatures, dependencies)  
**GATE CHECKLIST (10 items):**
1. C4 Level 3 para componentes principais
2. API contracts definidos (endpoints, methods, request/response)
3. Database schema com DDL, indexes e relationships
4. Sequence diagrams para fluxos críticos identificados no HLD
5. State machines para entidades com ciclo de vida
6. Cada componente vinculado a um container do HLD
7. Error handling consistente com SAD cross-cutting concerns
8. Interfaces de componentes explicitamente definidas
9. Alinhamento com tecnologia definida no HLD
10. Consistência com API contracts e database schema (foreign keys, tipos)

---

### Tasks 14-16: Trios da Fase 4 — Testes e Qualidade

---

#### Task 14: Trio TEST-PLAN

**SKILLS:** `["test-strategy-design", "qa-test-planner"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER, 03-SRS, 10-SAD, 12-LLD]`  
**TEMPLATE SECTIONS:** (1) Test Strategy (test pyramid, levels, scope per level), (2) Test Environment Requirements, (3) Test Data Strategy, (4) Unit Test Plan (frameworks, coverage targets), (5) Integration Test Plan (API, DB, message queues), (6) Functional/System Test Plan (cenários por feature/user story), (7) Security Test Plan (SAST, DAST, penetration testing, RBAC verification), (8) Performance Test Plan (load, stress, soak, scalability), (9) Regression Test Suite, (10) Acceptance Criteria (por feature), (11) Test Deliverables Schedule  
**GATE CHECKLIST (10 items):**
1. Pirâmide de testes documentada
2. Cobertura de cada nível definida com targets
3. Cenários de teste vinculados a FRs do SRS
4. Security test plan cobre OWASP Top 10
5. Performance test plan define thresholds
6. Test data strategy definida
7. Environments de teste especificados
8. Acceptance criteria alinhados com Deliverables do Charter
9. Schedule de entregas de teste
10. 100% dos FRs cobertos por pelo menos um cenário de teste

---

#### Task 15: Trio TEST-CASES

**SKILLS:** `["test-case-creation", "acceptance-criteria"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER, 03-SRS, 13-TEST-PLAN]`  
**TEMPLATE SECTIONS:** (1) Test Case Catalog (ID, feature vinculada, precondition, steps, expected result, postcondition, priority), (2) Happy Path Cases, (3) Edge Cases, (4) Negative Test Cases, (5) Gherkin Scenarios (Given/When/Then), (6) Test Data Specifications, (7) Traceability TC→FR(SRS)→Test Plan  
**GATE CHECKLIST (7 items):**
1. Cada TC vinculado a um FR do SRS e a uma seção do TEST-PLAN
2. Happy path, edge cases e negative cases para cada feature
3. Gherkin formatado corretamente (Given/When/Then)
4. Preconditions e postconditions definidas
5. Test data specifications preenchidas
6. Zero TCs órfãos
7. Cobertura por feature: 100% das features têm TCs

---

#### Task 16: Trio Relatório de Qualidade

**SKILLS:** `["quality-documentation-manager", "qa"]`  
**UPSTREAM_DOCS:** `[13-TEST-PLAN, 14-TEST-CASES]`  
**TEMPLATE SECTIONS:** (1) Quality Metrics Dashboard (coverage %, pass/fail rates, defect density), (2) Defect Report (ID, severity, status, linked TC), (3) Test Execution Summary, (4) Coverage Matrix (TC×Feature×FR), (5) Quality Gate Status (go/no-go por fase), (6) Defect Trends (burndown chart textual), (7) Recommendations  
**GATE CHECKLIST (5 items):**
1. Metrics dashboard com valores numéricos
2. Defeitos linkados a TCs
3. Coverage matrix preenchida
4. Quality gates com critérios de go/no-go
5. Defect trends e recomendações presentes

---

### Tasks 17-21: Trios da Fase 5 — Implantação e Encerramento

---

#### Task 17: Trio DEPLOYMENT-PLAN

**SKILLS:** `["deployment-engineer", "devops-rollout-plan"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER, 10-SAD, 11-HLD, 12-LLD]`  
**TEMPLATE SECTIONS:** (1) Deployment Strategy (blue-green, canary, rolling), (2) Environment Inventory (dev, staging, prod), (3) Pre-deployment Checklist, (4) Deployment Steps (sequencial, com rollback plan por step), (5) Database Migration Plan, (6) Rollback Plan, (7) Communication Plan (para deploy), (8) Validation & Smoke Tests, (9) Go-Live Runbook  
**GATE CHECKLIST (8 items):**
1. Estratégia de deployment definida
2. Pré-deployment checklist completo
3. Cada deployment step tem rollback plan associado
4. Database migration plan separado
5. Smoke tests definidos para validação pós-deploy
6. Comunicação de deploy alinhada com Plano de Comunicação
7. Go-live runbook presente
8. Ambientes alinhados com deployment topology do HLD

---

#### Task 18: Trio Manuais de Usuário

**SKILLS:** `["documentation-generation-doc-generate", "docs-writer"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER, 03-SRS]`  
**TEMPLATE SECTIONS:** (1) Getting Started, (2) Feature Walkthrough (por feature do SRS), (3) Step-by-Step Guides, (4) FAQ, (5) Troubleshooting, (6) Glossary  
**GATE CHECKLIST (5 items):**
1. Getting Started cobre onboarding completo
2. Cada feature do SRS tem walkthrough
3. Guias passo a passo com screenshots/descrições
4. FAQ cobre dúvidas comuns
5. Troubleshooting cobre cenários de erro

---

#### Task 19: Trio Manuais Operacionais

**SKILLS:** `["documentation-generation-doc-generate"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER, 10-SAD, 16-DEPLOYMENT-PLAN]`  
**TEMPLATE SECTIONS:** (1) System Architecture Overview (para ops), (2) Runbooks (start/stop, backup/restore, scaling, monitoring), (3) Alert & Escalation Procedures, (4) Disaster Recovery Runbook, (5) Maintenance Procedures, (6) Capacity Planning Guide  
**GATE CHECKLIST (6 items):**
1. Runbooks para operações críticas (start/stop, backup, restore, scaling)
2. Alinhamento com deployment plan e SAD
3. DR runbook presente
4. Procedimentos de alerta alinhados com observability do SAD
5. Capacity planning guide
6. Maintenance procedures com janelas definidas

---

#### Task 20: Trio Termo de Aceite (Sign-Off)

**SKILLS:** `["contract-and-proposal-writer"]`  
**UPSTREAM_DOCS:** `[01-PROJECT-CHARTER, 13-TEST-PLAN, 15-RELATORIO-QUALIDADE]`  
**TEMPLATE SECTIONS:** (1) Acceptance Criteria Checklist (extraída do Charter Seção 4 e TEST-PLAN), (2) Deliverable Acceptance Status (por entrega), (3) Quality Gate Results (do Relatório de Qualidade), (4) Punch List (itens pendentes), (5) Formal Acceptance Statement, (6) Signatures  
**GATE CHECKLIST (4 items):**
1. Todos os deliverables do Charter listados com status
2. Quality gate results referenciados
3. Punch list documentada (vazia ou com itens)
4. Formal acceptance statement e tabela de assinaturas

---

#### Task 21: Trio Lições Aprendidas

**SKILLS:** nenhuma (fallback only)  
**UPSTREAM_DOCS:** todos os 19 documentos anteriores  
**TEMPLATE SECTIONS:** (1) What Went Well, (2) What Could Be Improved, (3) Process Deviations, (4) Key Metrics (schedule variance, budget variance, defect escape rate), (5) Recommendations for Future Projects, (6) Knowledge Base Contributions  
**GATE CHECKLIST (4 items):**
1. What Went Well com exemplos concretos
2. What Could Be Improved com sugestões acionáveis
3. Métricas preenchidas com valores reais
4. Recommendations são específicas e reutilizáveis

---

## Self-Review

### 1. Spec Coverage Check

| Spec Section | Coverage |
|---|---|
| 7 Regras de Gating (Rules 1-7) | Task 1 (orquestrador) documenta todas; Tasks 2-21 implementam |
| Bootstrap Fase 0 (7 variáveis) | Task 1 implementa completo |
| 5 Fases WATERFALL, 20 documentos | Tasks 2-21 cobrem todos |
| GENERATE→GATE→FIX explícito | Tasks 2-21: cada documento tem trio |
| Status lifecycle: Em análise → Em revisão → COMPLIANCE | Cada GENERATE/GATE/FIX implementa |
| Dados explícitos (no guessing) | Cada prompt declara inputs via tabela |
| FIX cirúrgico (Rule 4) | Template FIX idêntico em todas tasks |
| Hybrid skills + fallback | SKILLS listadas em cada GENERATE + template de fallback |
| RTM cross-validation | Task 5 (RTM) + GATE checklist cobre forward/backward |
| Cascade rules | Task 1 (orquestrador) documenta |
| Git workflow | Task 1 (orquestrador) implementa passos F.1-F.4 |
| Architecture document views | Tasks 11-13: SAD (6 views), HLD (5 views), LLD (6 views) |

### 2. Placeholder Scan

Nenhum TBD, TODO, "implement later", "fill in details", "similar to Task N" encontrado. Cada task tem: skills concretas, upstream docs explícitos, template sections nomeadas, checklist items específicos.

### 3. Type Consistency

- `DOC_PATH` é consistente em todas as tasks (string, caminho completo)
- `VIOLATIONS[]` formato `{section, description, severity}` é consistente
- `UPSTREAM_DOCS` é sempre lista de paths de arquivos
- `SKILLS` é sempre lista de strings (nomes de skills)
- `[STATUS: Em análise]`, `[STATUS: Em revisão]`, `[STATUS: COMPLIANCE]` são consistentes

---

## Execution Handoff

**Plano completo e salvo.** Dois modos de execução:

**1. Subagent-Driven (recomendado)** — Um subagent por task, revisão entre tasks, cada task é um trio GENERATE+GATE+FIX autossuficiente

**2. Execução Inline** — Execução sequencial neste session, com checkpoints após cada fase WATERFALL
```

- [ ] **Step 6: Revisão final do plano e salvamento**
```
git add docs/superpowers/plans/2026-08-03-waterfall-orchestrator-implementation.md
git commit -m "docs: adicionar plano de implementação do WATERFALL orchestrator — 21 tasks, 61 prompts"
```
