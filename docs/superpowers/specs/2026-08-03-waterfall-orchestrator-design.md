# Waterfall Orchestrator — Design Document

**Status:** Approved (revised)  
**Date:** 2026-08-03  
**Version:** 2.0  
**Topic:** WATERFALL methodology project management roadmap orchestrator skill  
**Revision:** Incorporates 7 gating rules for explicit data flow, status lifecycle, and surgical fixes

---

## 1. Purpose & Scope

### 1.1 What this is

A reusable **skill orchestrator** that walks a project through the 20 WATERFALL documents across 5 phases, following a strict `GENERATE → GATE → FIX → COMPLIANCE` state machine. The orchestrator is a meta-framework — it works as a template/methodology for **multiple future projects**, not a single-project script.

### 1.2 What this is NOT

- NOT a single-project quick-start tool
- NOT an agile/scrum backlog generator (existing `project-documents/` roadmap already covers that)
- NOT a replacement for the existing `PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS.md` (which follows an agile/epic/feature/user-story model)

### 1.3 Scope boundaries

| In scope | Out of scope |
|---|---|
| 20 WATERFALL documents × GENERATE/GATE/FIX | Implementation/code generation |
| 5 sequential phases with human gating | Agile artifacts (epics, features, user stories) |
| Hybrid skill delegation + fallback prompts | Real-time collaboration features |
| Cross-document RTM traceability validation | Integration with external PM tools (Jira, etc.) |
| Git workflow (branch→commit→PR→merge) | Multi-language localization |
| Project resume (detect existing docs, continue from first gap) | Automated scheduling/reminders |

---

## 2. Seven Mandatory Rules (Gating Contract)

These rules are non-negotiable and apply to every prompt and every document in the system.

### Rule 1 — No Guessing

> **Prompts must NEVER guess, infer, or "discover" their inputs.** Every input (document paths, parameters, upstream artifacts) must be passed EXPLICITLY by the orchestrator. A GENERATE prompt receives its inputs as named parameters from the orchestrator. A GATE prompt receives the exact file path of the document to validate. A FIX prompt receives the exact file path AND the exact list of violations. No prompt ever searches directories or reads files it wasn't explicitly told to read.

### Rule 2 — Every GENERATE Has a GATE and a FIX

> Every document in the WATERFALL flow has exactly one GENERATE prompt, one GATE prompt, and one FIX prompt. The only exception is if a document genuinely requires no validation (currently none of the 20 documents qualify for this exception — all documents are validated).

### Rule 3 — Explicit Data Flow Between Prompts

> GATE MUST explicitly receive the outputs from GENERATE (the generated document file path). GATE MUST explicitly pass its findings (file path + list of violations) to FIX. FIX MUST explicitly receive both the file path and the violation list from GATE. No prompt discovers its inputs by searching the filesystem.

### Rule 4 — FIX Is Surgical Only

> FIX only makes targeted corrections to the violations reported by GATE. FIX does NOT regenerate, recreate, or rewrite the document. FIX receives: (a) the full path to the existing document file, (b) the exact list of non-conformities from GATE. FIX edits only those sections. Sections that passed the gate are untouched.

### Rule 5 — New Documents Start as "Em análise"

> When GENERATE creates a document, the status marker `[STATUS: Em análise]` is written into the document header immediately upon creation. This is the initial state of every new document.

### Rule 6 — Under Review Documents Are "Em revisão"

> When a GATE or FIX prompt processes a document, the status changes to `[STATUS: Em revisão]`. The document remains "Em revisão" throughout the GENERATE→GATE→FIX loop until the human confirms compliance.

### Rule 7 — Only COMPLIANCE Unlocks the Next Document

> The orchestrator ONLY advances to the next document when the current document is marked `[STATUS: COMPLIANCE]` AND the human explicitly confirms. No automatic progression. No batch approval. The transition is strictly sequential.

---

## 3. Architecture

### 3.1 Data Flow Contract (Explicit, No Guessing)

```
┌─────────────────────────────────────────────────────────────────────┐
│                        ORCHESTRATOR                                   │
│                                                                       │
│   For each document, the orchestrator passes:                         │
│                                                                       │
│   ┌──────────────────┐                                                │
│   │ GENERATE         │                                                │
│   │ Inputs from      │────► Creates file at {DOC_PATH}               │
│   │ orchestrator:    │     Writes [STATUS: Em análise]               │
│   │ • BRIEFING       │     Returns: {DOC_PATH}                       │
│   │ • UPSTREAM_DOCS  │                                                │
│   │ • PROJECT_CTX    │                                                │
│   │ • SKILLS_LIST    │                                                │
│   └──────────────────┘                                                │
│            │                                                          │
│            ▼ {DOC_PATH} passed explicitly                             │
│   ┌──────────────────┐                                                │
│   │ GATE             │                                                │
│   │ Input:           │────► Reads {DOC_PATH}                         │
│   │ • DOC_PATH       │     Validates against checklist               │
│   │ • CHECKLIST      │     Updates status to [STATUS: Em revisão]    │
│   │                  │     Returns: {PASS | FAIL + VIOLATIONS[]}     │
│   └──────────────────┘                                                │
│            │                                                          │
│     ┌──────┴──────┐                                                   │
│     │ PASS        │ FAIL                                              │
│     ▼             ▼                                                   │
│  ┌────────┐   ┌──────────────────┐                                    │
│  │HUMAN    │   │ FIX              │                                   │
│  │GATE     │   │ Input:           │──► Edits only violated sections  │
│  │3 Qs     │   │ • DOC_PATH       │    Keeps status [Em revisão]     │
│  └────────┘   │ • VIOLATIONS[]   │    Returns: {DOC_PATH}           │
│     │         └──────────────────┘                                   │
│     │              │                                                  │
│     │              ▼ back to GATE (with same DOC_PATH)                │
│     │                                                                 │
│  ┌──┴──────────┐                                                      │
│  │ APPROVED?    │                                                     │
│  │ Y: COMPLIANCE│────► Next document unlocks                          │
│  │ N: back to   │                                                     │
│  │   GENERATE   │                                                     │
│  └──────────────┘                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### 3.2 Document Status Lifecycle

```
                               GENERATE creates file
                               ┌─────────────────────┐
                               │ [STATUS: Em análise] │ ◄── Initial state (Rule 5)
                               └──────────┬──────────┘
                                          │
                              GATE runs (Rule 3: receives DOC_PATH)
                                          │
                               ┌──────────▼──────────┐
                               │[STATUS: Em revisão]  │ ◄── Under review (Rule 6)
                               └──────────┬──────────┘
                                          │
                          ┌───────────────┴───────────────┐
                          │ GATE verdict                   │
                          ▼                               ▼
                  ┌───────────────┐               ┌───────────────┐
                  │ PASS          │               │ FAIL           │
                  │ → Human Gate  │               │ → FIX (Rule 4) │
                  └───────┬───────┘               │ surgical only  │
                          │                       └───────┬───────┘
                          │                               │
                          │                         back to GATE
                          │                         (still "Em revisão")
                          │
              ┌───────────┴───────────┐
              │ Human answers 3 Qs    │
              ▼                       ▼
      ┌───────────────┐       ┌───────────────┐
      │ APPROVED       │       │ NEEDS CHANGES  │
      │ → COMPLIANCE   │       │ → back to      │
      │   (Rule 7)     │       │   GENERATE     │
      └───────┬───────┘       │ (new inputs)   │
              │               └───────────────┘
              ▼
    ┌─────────────────────┐
    │[STATUS: COMPLIANCE]  │ ◄── Frozen. Next doc unlocks (Rule 7)
    └─────────────────────┘
```

### 3.3 Orchestrator Flow

```
┌──────────────────────────────────────────────────────────────────┐
│                    WATERFALL ORCHESTRATOR                          │
│                                                                    │
│  BOOTSTRAP FASE 0                                                 │
│  ├─ 0.1 Collect 7 variables (PROJECT_PATH, PROJECT_ID, ...)       │
│  ├─ 0.2 Validate PROMPT_BRANCH (block main/master/develop)        │
│  ├─ 0.3 Confirm derived paths                                     │
│  ├─ 0.4 Create directory structure                                │
│  ├─ 0.5 Scan existing documents (resume capability)               │
│  └─ 0.6 Display summary, identify first pending phase             │
│                                                                    │
│  PHASE LOOP (for each of 5 phases)                                │
│  │                                                                 │
│  ├─ FASE 1: INICIALIZACAO                                         │
│  │   └─ PROJECT-CHARTER                                           │
│  │                                                                 │
│  ├─ FASE 2: PLANEJAMENTO-REQUISITOS                               │
│  │   ├─ BRD → SRS → RTM → EAP/WBS → Cronograma/Gantt             │
│  │   ├─ Orçamento → Plano-Comunicacao → Plano-Riscos              │
│  │   └─ [optional: RTM cross-validation after last doc]           │
│  │                                                                 │
│  ├─ FASE 3: DESIGN-ARQUITETURA                                    │
│  │   └─ SAD → HLD → LLD                                           │
│  │                                                                 │
│  ├─ FASE 4: TESTES-QUALIDADE                                      │
│  │   └─ TEST-PLAN → TEST-CASES → RELATORIO-QUALIDADE              │
│  │                                                                 │
│  └─ FASE 5: IMPLANTACAO-ENCERRAMENTO                              │
│      └─ DEPLOYMENT-PLAN → MANUAIS-USUARIO → MANUAIS-OPERACIONAIS  │
│         → TERMO-ACEITE → LICOES-APRENDIDAS                        │
│                                                                    │
│  FINALIZATION: Git workflow (add→commit→push→PR→merge→cleanup)     │
└──────────────────────────────────────────────────────────────────┘
```

### 3.4 Hybrid Skill Strategy

Each GENERATE prompt follows the **Option C hybrid** approach:
- **Primary**: Delegates to the best-matching installed skill (via `Skill` tool)
- **Fallback**: The GENERATE prompt itself contains the full template and instructions to produce the document without any external skill

| Document | Primary Skill(s) | Has Fallback |
|---|---|---|
| PROJECT-CHARTER | `draft-project-charter`, `senior-pm` | ✅ |
| BRD | `brd-creation`, `business-analyst`, `requirements-elicitation` | ✅ |
| SRS | `frs-creation`, `requirements-engineering` | ✅ |
| RTM | `requirements-modeling`, `requirements-validation` | ✅ (fallback only) |
| EAP/WBS | `decomposition-planning-roadmap`, `project-estimation` | ✅ |
| Cronograma/Gantt | `roadmap-planning`, `project-estimation` | ✅ |
| Orçamento | `project-estimation` | ✅ (fallback only) |
| Plano de Comunicação | `stakeholder-analysis`, `stakeholder-map` | ✅ |
| Plano de Riscos | `risk-manager`, `risk-management-specialist` | ✅ |
| SAD | `software-architecture`, `architecture-designer`, `architecture-patterns` | ✅ |
| HLD | `c4-container`, `system-design`, `architecture-decision-records` | ✅ |
| LLD | `c4-component`, `ddd-tactical-patterns`, `database-designer` | ✅ |
| TEST-PLAN | `test-strategy-design`, `qa-test-planner` | ✅ |
| TEST-CASES | `test-case-creation`, `acceptance-criteria` | ✅ |
| RELATORIO-QUALIDADE | `quality-documentation-manager`, `qa` | ✅ (fallback only) |
| DEPLOYMENT-PLAN | `deployment-engineer`, `devops-rollout-plan` | ✅ |
| MANUAIS-USUARIO | `documentation-generation-doc-generate`, `docs-writer` | ✅ |
| MANUAIS-OPERACIONAIS | `documentation-generation-doc-generate` | ✅ |
| TERMO-ACEITE | `contract-and-proposal-writer` | ✅ (fallback only) |
| LICOES-APRENDIDAS | — | ✅ (fallback only) |

### 3.5 Architecture Document Vision Requirements

Each architecture document covers multiple views:

| Document | Required Views |
|---|---|
| **SAD** | Solution Architecture, Data Architecture, Security Architecture, DevOps/SRE Architecture, Infrastructure/Cloud Architecture, Testing Architecture |
| **HLD** | Macro components, data flows, integration topology, deployment topology, technology stack decisions |
| **LLD** | Class diagrams, API contracts, database schemas, sequence diagrams, state machines, component interfaces |

---

## 4. Directory Structure

### 4.1 Prompt files

```
.specs/prompts/project-documents-waterfall/
├── PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md   ← ORCHESTRATOR
├── PROMPT-GENERATE-PROJECT-CHARTER.md
├── PROMPT-GATE-PROJECT-CHARTER.md
├── PROMPT-FIX-PROJECT-CHARTER.md
├── PROMPT-GENERATE-BRD.md
├── PROMPT-GATE-BRD.md
├── PROMPT-FIX-BRD.md
├── PROMPT-GENERATE-SRS.md
├── PROMPT-GATE-SRS.md
├── PROMPT-FIX-SRS.md
├── PROMPT-GENERATE-RTM.md
├── PROMPT-GATE-RTM.md
├── PROMPT-FIX-RTM.md
├── PROMPT-GENERATE-EAP-WBS.md
├── PROMPT-GATE-EAP-WBS.md
├── PROMPT-FIX-EAP-WBS.md
├── PROMPT-GENERATE-CRONOGRAMA-GANTT.md
├── PROMPT-GATE-CRONOGRAMA-GANTT.md
├── PROMPT-FIX-CRONOGRAMA-GANTT.md
├── PROMPT-GENERATE-ORCAMENTO.md
├── PROMPT-GATE-ORCAMENTO.md
├── PROMPT-FIX-ORCAMENTO.md
├── PROMPT-GENERATE-PLANO-COMUNICACAO.md
├── PROMPT-GATE-PLANO-COMUNICACAO.md
├── PROMPT-FIX-PLANO-COMUNICACAO.md
├── PROMPT-GENERATE-PLANO-RISCOS.md
├── PROMPT-GATE-PLANO-RISCOS.md
├── PROMPT-FIX-PLANO-RISCOS.md
├── PROMPT-GENERATE-SAD.md
├── PROMPT-GATE-SAD.md
├── PROMPT-FIX-SAD.md
├── PROMPT-GENERATE-HLD.md
├── PROMPT-GATE-HLD.md
├── PROMPT-FIX-HLD.md
├── PROMPT-GENERATE-LLD.md
├── PROMPT-GATE-LLD.md
├── PROMPT-FIX-LLD.md
├── PROMPT-GENERATE-TEST-PLAN.md
├── PROMPT-GATE-TEST-PLAN.md
├── PROMPT-FIX-TEST-PLAN.md
├── PROMPT-GENERATE-TEST-CASES.md
├── PROMPT-GATE-TEST-CASES.md
├── PROMPT-FIX-TEST-CASES.md
├── PROMPT-GENERATE-RELATORIO-QUALIDADE.md
├── PROMPT-GATE-RELATORIO-QUALIDADE.md
├── PROMPT-FIX-RELATORIO-QUALIDADE.md
├── PROMPT-GENERATE-DEPLOYMENT-PLAN.md
├── PROMPT-GATE-DEPLOYMENT-PLAN.md
├── PROMPT-FIX-DEPLOYMENT-PLAN.md
├── PROMPT-GENERATE-MANUAIS-USUARIO.md
├── PROMPT-GATE-MANUAIS-USUARIO.md
├── PROMPT-FIX-MANUAIS-USUARIO.md
├── PROMPT-GENERATE-MANUAIS-OPERACIONAIS.md
├── PROMPT-GATE-MANUAIS-OPERACIONAIS.md
├── PROMPT-FIX-MANUAIS-OPERACIONAIS.md
├── PROMPT-GENERATE-TERMO-ACEITE.md
├── PROMPT-GATE-TERMO-ACEITE.md
├── PROMPT-FIX-TERMO-ACEITE.md
├── PROMPT-GENERATE-LICOES-APRENDIDAS.md
├── PROMPT-GATE-LICOES-APRENDIDAS.md
└── PROMPT-FIX-LICOES-APRENDIDAS.md
```

**Total:** 1 orchestrator + 20 generators + 20 gates + 20 fixers = **61 prompt files**

### 4.2 Generated project documents (at runtime)

```
{PROJECT_COMPLETE_PATH_NAME}/
├── 01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md
├── 02-BRD-{PROJECT_ID_NAME}.md
├── 03-SRS-{PROJECT_ID_NAME}.md
├── 04-RTM-{PROJECT_ID_NAME}.md
├── 05-EAP-WBS-{PROJECT_ID_NAME}.md
├── 06-CRONOGRAMA-GANTT-{PROJECT_ID_NAME}.md
├── 07-ORCAMENTO-{PROJECT_ID_NAME}.md
├── 08-PLANO-COMUNICACAO-{PROJECT_ID_NAME}.md
├── 09-PLANO-RISCOS-{PROJECT_ID_NAME}.md
├── 10-SAD-{PROJECT_ID_NAME}.md
├── 11-HLD-{PROJECT_ID_NAME}.md
├── 12-LLD-{PROJECT_ID_NAME}.md
├── 13-TEST-PLAN-{PROJECT_ID_NAME}.md
├── 14-TEST-CASES-{PROJECT_ID_NAME}.md
├── 15-RELATORIO-QUALIDADE-{PROJECT_ID_NAME}.md
├── 16-DEPLOYMENT-PLAN-{PROJECT_ID_NAME}.md
├── 17-MANUAIS-USUARIO-{PROJECT_ID_NAME}.md
├── 18-MANUAIS-OPERACIONAIS-{PROJECT_ID_NAME}.md
├── 19-TERMO-ACEITE-{PROJECT_ID_NAME}.md
└── 20-LICOES-APRENDIDAS-{PROJECT_ID_NAME}.md
```

---

## 5. Bootstrap (Phase 0) — Variables & Inputs

### 5.1 Required variables

| # | Variable | Description | Example |
|---|---|---|---|
| 1 | `PROJECT_PATH` | Base path for business projects | `/home/bolismar/work/workspace-fbso/business-inputs/business-projects` |
| 2 | `PROJECT_ID` | Corporate project identifier | `PRJ-FIN-2026-0003` |
| 3 | `PROJECT_NAME` | Short product/project name | `SAAS-FBSO-ORG` |
| 4 | `PROJECT_BRIEFING` | Free-text briefing or path to briefing file | `"Portal de autoatendimento B2B..."` or `/tmp/briefing.md` |
| 5 | `PROJECT_DOCUMENTS_INPUTS` | List of paths to raw input documents | `[]` |
| 6 | `PROJECT_PROMPT_INPUTS` | List of paths to auxiliary prompts/context | `[]` |
| 7 | `PROMPT_BRANCH` | Git branch name (blocked: main, master, develop) | `feature/PRJ-FIN-2026-0003-docs` |

### 5.2 Derived variables (computed automatically)

```
PROJECT_ID_NAME            = PROJECT_ID + "-" + PROJECT_NAME
PROJECT_COMPLETE_PATH_NAME = PROJECT_PATH + "/" + PROJECT_ID_NAME
```

### 5.3 Bootstrap steps

1. **Collect inputs** — ask for any missing variables. Never guess.
2. **Validate branch** — block if `main`, `master`, or `develop`
3. **Confirm paths** — display derived `PROJECT_COMPLETE_PATH_NAME` and `PROJECT_ID_NAME`, ask for explicit confirmation
4. **Create directory** — `mkdir -p {PROJECT_COMPLETE_PATH_NAME}`
5. **Scan existing documents** — check which of the 20 files exist; for each, read headers and check for `[STATUS: COMPLIANCE]` marker
6. **Resume decision**:
   - All absent → start from Fase 1 (PROJECT-CHARTER)
   - Partial → start from the first document without `[STATUS: COMPLIANCE]`
   - All compliant → ask: review specific phase, new evolution cycle, or exit

---

## 6. Orchestration Mechanism

### 6.1 Per-Document Execution Loop (Rules 1-7 Applied)

For each document `{DOC}` in sequence:

```
STEP 1: ORCHESTRATOR computes explicit inputs for GENERATE
        • DOC_PATH = {PROJECT_COMPLETE_PATH_NAME}/{NN}-{DOC}-{PROJECT_ID_NAME}.md
        • UPSTREAM_DOCS = list of file paths of upstream COMPLIANCE documents
        • BRIEFING = PROJECT_BRIEFING
        • EXTRA_INPUTS = PROJECT_DOCUMENTS_INPUTS + PROJECT_PROMPT_INPUTS
        • SKILLS = [{skill_list_for_this_doc}]

STEP 2: ORCHESTRATOR invokes PROMPT-GENERATE-{DOC}.md
        Passes all inputs from Step 1 EXPLICITLY as parameters.
        GENERATE MUST:
        - Never guess or search for inputs
        - Create file at DOC_PATH
        - Write [STATUS: Em análise] as the initial status (Rule 5)
        - Return {DOC_PATH} to orchestrator

STEP 3: ORCHESTRATOR invokes PROMPT-GATE-{DOC}.md
        Explicitly passes:
        - DOC_PATH (from GENERATE output)
        - CHECKLIST (from GATE prompt definition)
        GATE MUST:
        - Read ONLY the file at DOC_PATH
        - Update status to [STATUS: Em revisão] (Rule 6)
        - Return {PASS} or {FAIL, VIOLATIONS[]}

STEP 4a: IF GATE returns FAIL:
        ORCHESTRATOR invokes PROMPT-FIX-{DOC}.md
        Explicitly passes:
        - DOC_PATH (same file, from GATE)
        - VIOLATIONS[] (exact list of non-conformities from GATE)
        FIX MUST:
        - Edit ONLY the sections referenced in VIOLATIONS[] (Rule 4)
        - NOT regenerate or recreate the file
        - Keep status as [STATUS: Em revisão]
        - Return {DOC_PATH}
        → Go back to STEP 3 (re-run GATE)

STEP 4b: IF GATE returns PASS:
        ORCHESTRATOR presents the document to the human with 3 questions:
        P1: "O conteúdo está aderente às necessidades do projeto?"
        P2: "Existem novos documentos de entrada a incorporar?"
        P3: "Há novas informações textuais ou ajustes de escopo?"

        IF human approves (YES to P1, NO to P2/P3):
        → Status becomes [STATUS: COMPLIANCE] (Rule 7)
        → Document is frozen
        → Next document unlocks

        IF human provides new inputs (P2 or P3):
        → Go back to STEP 2 (re-run GENERATE with new context)
        → Document stays in [STATUS: Em análise]
```

### 6.2 Cascade Rules

When a document is modified after being marked COMPLIANCE, all downstream documents are flagged:

| If modifying... | Cascade impact (regenerate + revalidate)... |
|---|---|
| PROJECT-CHARTER | ALL 19 downstream documents |
| BRD | SRS → RTM → ... → LICOES-APRENDIDAS (16 docs) |
| SRS | RTM → ... → LICOES-APRENDIDAS (15 docs) |
| SAD | HLD → LLD → ... → LICOES-APRENDIDAS (10 docs) |
| (and so on, following the sequential dependency chain) |

**Cascade action**: The orchestrator alerts the human, lists affected documents, and asks: (A) proceed with full regeneration of downstreams, or (B) only update current document and mark downstreams as "potentially outdated."

### 6.3 RTM Cross-Validation (End of Phase 2)

After all Phase 2 documents reach COMPLIANCE:
- Every BRD requirement must trace to a PROJECT-CHARTER objective
- Every SRS specification must trace to a BRD requirement
- The RTM matrix must have zero orphans and 100% coverage
- Report: Pass/Fail with gap details

---

## 7. Prompt Templates (Explicit Data Flow)

### 7.1 GENERATE prompt template

```markdown
# PROMPT: GERADOR DE {DOCUMENT_NAME}

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado | `{PROJECT_COMPLETE_PATH_NAME}/01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` |
| `PROJECT_ID_NAME` | Identificador do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `BRIEFING` | Briefing do projeto (texto ou caminho de arquivo) | `"Portal de autoatendimento..."` |
| `UPSTREAM_DOCS` | Lista de caminhos para documentos upstream em COMPLIANCE | `[/path/to/01-PROJECT-CHARTER-....md]` |
| `EXTRA_INPUTS` | Documentos e prompts extras fornecidos pelo humano | `[]` |
| `SKILLS` | Lista de skills a serem utilizadas | `["draft-project-charter", "senior-pm"]` |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado acima
2. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
3. Skills: use as listadas em `SKILLS`. Se falharem, use o template de fallback abaixo
4. Ao final, retorne `{DOC_PATH}` confirmando a criação

## Template de Fallback

{Estrutura completa do documento com seções}
```

### 7.2 GATE prompt template

```markdown
# PROMPT: PORTÃO DE VALIDAÇÃO DE {DOCUMENT_NAME}

## Inputs (recebidos explicitamente — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo do arquivo a ser validado |
| `CHECKLIST` | Lista de itens de compliance a verificar |

## Regras

1. Leia **APENAS** o arquivo em `DOC_PATH`
2. Altere o status do documento para `[STATUS: Em revisão]`
3. Execute cada item do `CHECKLIST` contra o conteúdo do documento
4. Retorne `{PASS}` se todos os checks passarem
5. Retorne `{FAIL, VIOLATIONS: [{section, description, severity}]}` se houver falhas

## Checklist de Compliance

1. {CHECK_1}
2. {CHECK_2}
...
```

### 7.3 FIX prompt template

```markdown
# PROMPT: CORRETOR DE {DOCUMENT_NAME}

## Inputs (recebidos explicitamente do GATE — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo do arquivo a ser corrigido |
| `VIOLATIONS` | Lista de não-conformidades reportadas pelo GATE |

## Regras

1. Edite **APENAS** as seções listadas em `VIOLATIONS` (correção cirúrgica — Rule 4)
2. **NÃO** recrie, regenere ou reescreva o documento inteiro
3. **NÃO** altere seções que passaram no GATE
4. Mantenha o status como `[STATUS: Em revisão]`
5. Retorne `{DOC_PATH}` após as correções

## Não-Conformidades a Corrigir

(VIOLATIONS serão injetadas aqui pelo orquestrador a partir do output do GATE)
```

---

## 8. Implementation Plan

### 8.1 Build order

1. **Orchestrator** (`PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md`) — 1 file
2. **Phase 1** — PROJECT-CHARTER (GENERATE, GATE, FIX) — 3 files
3. **Phase 2** — BRD, SRS, RTM, EAP/WBS, Cronograma/Gantt, Orçamento, Plano-Comunicacao, Plano-Riscos — 24 files
4. **Phase 3** — SAD, HLD, LLD — 9 files
5. **Phase 4** — TEST-PLAN, TEST-CASES, RELATORIO-QUALIDADE — 9 files
6. **Phase 5** — DEPLOYMENT-PLAN, MANUAIS-USUARIO, MANUAIS-OPERACIONAIS, TERMO-ACEITE, LICOES-APRENDIDAS — 15 files

**Total**: 61 files

### 8.2 Each GATE checklist must be document-specific

Every GATE prompt defines its own unique checklist tailored to that document. Examples:

- **PROJECT-CHARTER GATE**: validates 14 sections present, RACI matrix complete, success criteria measurable, stakeholders named, scope boundaries explicit
- **SAD GATE**: validates 6 architecture views present, each view references upstream requirements, technology decisions have rationale, ADRs are traceable
- **RTM GATE**: validates zero orphan requirements, 100% BRD→SRS traceability, matrix format compliance, bidirectional linking

---

## 9. Error Handling

| Scenario | Behavior |
|---|---|
| Missing required variable | Ask user, do not proceed |
| Invalid PROMPT_BRANCH | Block, ask again |
| Gate finds non-compliance | Enter FIX loop (max 3 iterations, then escalate to human) |
| Human provides new inputs during human gate | Restart GENERATE with new context merged; document goes back to "Em análise" |
| Cascade triggered | Alert human, list affected docs, ask for decision |
| Git push fails (branch exists) | Ask about `--force` |
| Git merge conflict | Abort, preserve branch, report to human |
| Skill not found or fails | Fall back to embedded prompt template silently |

---

## 10. Design Decisions & Rationale

| Decision | Rationale |
|---|---|
| 61 separate prompt files | Matches existing project-documents and technical-solutions patterns; each file is independently editable and auditable |
| Option C (hybrid skills + fallbacks) | Maximizes reuse of 20+ installed skills while guaranteeing zero external dependency failures block the pipeline |
| Explicit data flow (Rule 1, 3) | Prevents prompts from guessing, searching directories, or discovering wrong artifacts — the orchestrator is the single source of truth for what files exist and what their paths are |
| Status lifecycle: Em análise → Em revisão → COMPLIANCE | Provides full audit trail; each status indicates exactly where in the GENERATE→GATE→FIX loop the document is |
| Surgical FIX only (Rule 4) | Prevents FIX from undoing human-approved content or introducing regressions in sections that already passed GATE |
| Sequential gating (Rule 7) | Core WATERFALL principle; enforced by human approval gate on every document |
| Architecture documents require 6+ views each | Ensures comprehensive coverage across solution, data, security, DevOps, infra, and testing dimensions |

---

## 11. Spec Self-Review

- **Placeholder scan**: No TBD or TODO markers. All 20 documents are named. All 7 rules are explicit.
- **Internal consistency**: The data flow contract (Section 3.1), status lifecycle (Section 3.2), and execution loop (Section 6.1) are mutually consistent. The prompt templates (Section 7) exactly match the execution loop steps. The 3-status lifecycle replaces the earlier 5-state model and is simpler and unambiguous.
- **Scope check**: Single spec for the orchestrator skill + its 61 prompts. They share the same architecture, state machine, and bootstrap context.
- **Ambiguity check**: Every transition has a precise condition. "Em análise" has exactly one trigger (GENERATE creates the file). "Em revisão" has exactly one trigger (GATE or FIX processes). "COMPLIANCE" has exactly one trigger (human approves). No prompt ever guesses its inputs — every parameter is named and passed explicitly.

---

🤖 *Design document v2.0 generated by Claude Code. Skills used: brainstorming, find-skills.*
