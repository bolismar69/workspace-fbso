# Spec: Roadmap de Definições Técnicas v5.0 — Redesign Estrutural

- **Data:** 2026-07-30
- **Status:** Design aprovado — aguardando implementação
- **Branch:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG-revisao-prompts-roadmap
- **Documento base:** `PROMPT-ROADMAP-GENERATE-PROJECT-TECHNICAL-DEFINITIONS.md` v4.0

---

## 1. Motivação e Objetivo

Reestruturar o roadmap de definições técnicas para:

1. Criar uma **ponte formal Negócio→TI** (Bloco 0) com ingestão de requisitos, DoR e PRD de Negócio
2. Expandir o Bloco B para cobrir **6 disciplinas técnicas** com artefatos independentes
3. Reorganizar o Bloco C para fluxo **Catálogo→Matriz→Stack→Specs→Milestones**
4. Repropositar o Bloco D como **Sprints com Discovery Técnico Contínuo**
5. Posicionar **EXECUTION-HISTORY** como fase standalone após o último bloco

**Abordagem escolhida:** Pipeline totalmente sequencial, sem paralelismo (Abordagem 1).

---

## 2. Arquitetura Macro dos Blocos (v5.0)

```
FASE 0: BOOTSTRAP
  │
  ▼
╔══════════════════════════════════════════════════════════════╗
║  BLOCO 0: Product Definition & Product Backlog & PRD  🆕   ║
║  F1: 🆕 Novas Features/Ideias (Lotes de Ingestão)          ║
║  F2: 📐 Refinamento de Negócio (PO/PM) → DoR de Negócio    ║
║  F3: 📋 PRODUCT BACKLOG LIST (Pronto para TI)              ║
║  F4: 📄 PRD-DEFINITION (Business PRD — congelado)          ║
║  ⛔ Barreira 0: GATE/FIX do Bloco                          ║
╚══════════════════════════════════════════════════════════════╝
  │
  ▼
╔══════════════════════════════════════════════════════════════╗
║  BLOCO A: People & Solutions (mantido)                      ║
║  F5: TEAM-SKILLS-MAP / F6: TEAM-CAPACITY                   ║
║  ⛔ Barreira A                                              ║
╚══════════════════════════════════════════════════════════════╝
  │
  ▼
╔══════════════════════════════════════════════════════════════╗
║  BLOCO B: Architecture & Security & Specialists  🔄        ║
║  F7: ARCHITECTURE / F8: SECURITY / F9: DATA-ARCH 🆕        ║
║  F10: DEVOPS-SRE 🆕 / F11: TEST-STRATEGY 🆕                ║
║  F12: INFRA-CLOUD 🆕                                        ║
║  👤 Tech Lead + Principal Architecture (revisores)          ║
║  ⛔ Barreira B: valida TODAS as disciplinas do escopo       ║
╚══════════════════════════════════════════════════════════════╝
  │
  ▼
╔══════════════════════════════════════════════════════════════╗
║  BLOCO C: Catálogo, Matriz, Stack, Specs & Milestones  🔄  ║
║  F13: CATALOG → F14: MATRIX → F15: STACK-MATRIX            ║
║  F16: SPECS-DEFINITION (consolidação técnica enxuta)        ║
║  F17: MILESTONES                                            ║
║  ⛔ Barreira C: com feedback loop → Bloco A se skills gaps  ║
╚══════════════════════════════════════════════════════════════╝
  │
  ▼
╔══════════════════════════════════════════════════════════════╗
║  BLOCO D: Sprints  🔄                                       ║
║  F18: SPRINT-BACKLOG-REFINED (technical-discovery/)         ║
║  F19: 🔍 Discovery Técnico Contínuo de User Stories         ║
║  ⛔ Barreira D                                              ║
╚══════════════════════════════════════════════════════════════╝
  │
  ▼
╔══════════════════════════════════════════════════════════════╗
║  EXECUTION-HISTORY 📊 (standalone, após último bloco)       ║
╚══════════════════════════════════════════════════════════════╝
```

**Paralelismo:** Nenhum. Pipeline totalmente sequencial — cada barreira libera exatamente um bloco seguinte.

---

## 3. Detalhamento dos Blocos

### 3.1 Bloco 0 — Product Definition & Product Backlog & PRD

**Propósito:** Ponte de integração Negócio→TI. Recebe requisitos, aplica DoR, gera backlog e PRD de Negócio.

| Fase | Artefato | Pipeline | Descrição |
|:---|:---|:---|:---|
| F1 | `INTAKE-LOG.md` 🆕 | G→G→F | Registro versionado dos lotes de ingestão (Waterfall: lote único; Ágil: múltiplas ondas) |
| F2 | `DOR-ASSESSMENT.md` 🆕 | G→G→F | Aplicação do DoR de Negócio pelo PO/PM. Itens que não passam voltam para refinamento |
| F3 | `PRODUCT-BACKLOG-LIST.md` 🆕 | G→G→F | Backlog consolidado "Pronto para TI", priorizado, com links para origem nos docs de negócio |
| F4 | `PRD-DEFINITION.md` (movido) | G→G→F | PRD de Negócio congelado. Visão do Produto, MVP Global, Glossário. Input para todos os blocos seguintes |

**Regras:**
- Ingestão versionada (v1, v2, ...)
- PRD congelado após Barreira 0 — alterações entram como novo lote F1
- Barreira 0 valida rastreabilidade: todo item do Backlog → origem nos docs de negócio e coberto pelo PRD

### 3.2 Bloco A — People & Solutions (mantido)

| Fase | Artefato | Pipeline |
|:---|:---|:---|
| F5 | `TEAM-SKILLS-MAP.md` | G→G→F |
| F6 | `TEAM-CAPACITY.md` | G→G→F |

Sem alterações estruturais. Mantém pipeline existente.

### 3.3 Bloco B — Architecture & Security & Specialists (expandido)

**Propósito:** 6 disciplinas técnicas com artefatos independentes. Tech Lead + Principal Architecture como revisores transversais.

| Fase | Disciplina | Artefato | Skills-Chave |
|:---|:---|:---|:---|
| F7 | Solution Architect | `ARCHITECTURE-DEFINITION.md` | `senior-architect`, `c4-architecture-c4-architecture`, `architecture-patterns` |
| F8 | Security Architect | `SECURITY-DEFINITION.md` | `senior-security`, `threat-modeling-expert`, `security-best-practices` |
| F9 🆕 | Data Architect | `DATA-ARCHITECTURE-DEFINITION.md` | `senior-data-engineer`, `data-modeling`, `database-architect`, `data-engineering-data-pipeline` |
| F10 🆕 | DevOps/SRE Architect | `DEVOPS-SRE-DEFINITION.md` | `senior-devops`, `cloud-devops`, `sre-engineer`, `kubernetes-specialist`, `observability-engineer` |
| F11 🆕 | Test Specialist | `TEST-STRATEGY-DEFINITION.md` | `senior-qa`, `qa-test-planner`, `test-strategy-design`, `tdd-guide`, `e2e-testing-patterns` |
| F12 🆕 | Infra/Cloud Specialist | `INFRA-CLOUD-DEFINITION.md` | `cloud-architect`, `aws-solution-architect`, `cloud-design-patterns`, `network-engineer` |

**Papéis transversais (sem artefato):**
- **Tech Lead:** presente em todas as reuniões do comitê, garante coesão cross-discipline
- **Principal Architecture:** acionado em projetos complexos, revisa decisões cross-discipline

**Escopo dos novos artefatos:**

- **F9 — DATA-ARCHITECTURE-DEFINITION:** Modelagem de dados (ERD, schemas), estratégia de armazenamento (SQL/NoSQL/cache/DW), pipelines ETL/ELT, integrações inter-banco, data governance, On-Premise vs Cloud
- **F10 — DEVOPS-SRE-DEFINITION:** Pipeline CI/CD, IaC, observabilidade, SLOs/SLIs, containers/orquestração, gestão de ambientes, runbooks
- **F11 — TEST-STRATEGY-DEFINITION:** Pirâmide de testes, automação, performance (carga/stress/soak), segurança (SAST/DAST), ambientes/dados de teste, quality gates
- **F12 — INFRA-CLOUD-DEFINITION:** Topologia, compute (VMs/K8s/serverless), networking, storage, DR (RPO/RTO), multi-region, auto-scaling

### 3.4 Bloco C — Catálogo, Matriz, Stack, Specs & Milestones

**Propósito:** Consolidar definições técnicas. CATALOG→MATRIX→STACK alimentam SPECS-DEFINITION, que deriva MILESTONES.

| Fase | Artefato | Pipeline | Descrição |
|:---|:---|:---|:---|
| F13 | `SOLUTIONS-CATALOG.md` | G→G→F | Catálogo de soluções técnicas |
| F14 | `SOLUTIONS-MATRIX.md` | G→G→F | Matriz solução×disciplina×owner |
| F15 | `SOLUTIONS-STACK-MATRIX.md` | G→G→F | Stack tecnológica por solução |
| F16 | `SPECS-DEFINITION.md` | G→G→F | Consolidação técnica enxuta — sumariza e referencia artefatos anteriores |
| F17 | `MILESTONES.md` | G→G→F | Roadmap com milestones, dependências, riscos |

**Fluxo interno:** F13→F14→F15→F16→F17 (sequencial). F13/F14/F15 convergem como inputs para F16.

**SPECS-DEFINITION como consolidação enxuta:** Cada seção tem ~1 parágrafo de sumário + link para o artefato completo (ARCHITECTURE-DEFINITION, DATA-ARCHITECTURE, SECURITY, etc.). Não repete conteúdo.

### 3.5 Bloco D — Sprints (Technical Discovery)

**Propósito:** Discovery técnico por sprint. Prepara US para execução com contratos concretos.

| Fase | Artefato | Pipeline | Descrição |
|:---|:---|:---|:---|
| F18 | `SPRINT-BACKLOG.md` (em `technical-discovery/`) | G→G→F | Backlog refinado com tarefas T-NNN → US-ID → Sprint-Alvo. Baseado no modelo existente, enriquecido com coluna CONTRACTS |
| F19 | `technical-discovery/sprint-NNN/` | G→G→F | Contratos técnicos por sprint. Iterativo. |

**Modelo SPRINT-BACKLOG (F18) — enriquecido do existente:**

| TASK-ID | TASK-DESCRIÇÃO | SPRINT-ALVO | US-ID | STATUS | CONTRACTS |
|---|---|---|---|---|---|
| T-000010 | Auditar endpoint GET /dashboard/admin/summary | Sprint 01 | US-FEAT-EP-0001-0001-0001 | TODO | [API](sprint-01/CONTRACTS-API.md) · [DATA](sprint-01/CONTRACTS-DATA.md) |

**Estrutura do technical-discovery/ (substitui sprints/):**

```
technical-discovery/
├── SPRINT-BACKLOG.md
├── sprint-00/
│   ├── CONTRACTS-API-sprint-00.md
│   ├── CONTRACTS-DATA-sprint-00.md
│   ├── CONTRACTS-SECURITY-sprint-00.md
│   ├── CONTRACTS-SRE-sprint-00.md
│   └── DEFINITION-INCREMENTS-sprint-00.md
├── sprint-01/ ...
└── sprint-NN/ ...
```

### 3.6 EXECUTION-HISTORY (standalone)

Dashboard consolidado com estado de todos os 20 artefatos. Pipeline: Generate → Revisão Humana (sem gate automatizado). Atualizado incrementalmente após cada fase.

---

## 4. Regras de Gating e Barreiras

### 4.1 Tabela de Barreiras

| Barreira | Posição | O que Valida | Comportamento Especial |
|:---|:---|:---|:---|
| ⛔ Barreira 0 | Após Bloco 0 | Rastreabilidade Backlog→Docs Negócio. DoR aplicado em 100%. PRD cobre todo backlog. | Itens sem DoR → voltam para F2. PRD incompleto → volta para F4 |
| ⛔ Barreira A | Após Bloco A | TEAM-SKILLS-MAP cobre todos os papéis. TEAM-CAPACITY preenchido. | — |
| ⛔ Barreira B | Após Bloco B | 6 disciplinas OK. N/A justificados. Consistência horizontal cross-artefacts. Tech Lead + Principal Architecture assinam. | Disciplina N/A sem justificativa = NÃO COMPLIANCE |
| ⛔ Barreira C | Após Bloco C | SPECS referencia todos artefatos. MILESTONES alinhado. | **Skills-Gap Detection:** Se skill necessária não coberta pelo Bloco A → `[SKILLS-GAP-DETECTED]` → propõe reabertura Bloco A |
| ⛔ Barreira D | Após Bloco D | 100% US da sprint com contratos. SPRINT-BACKLOG atualizado. | Iterativo — pergunta se continua próxima sprint ou encerra |

### 4.2 Mecanismo de Orquestração (mantido)

Todas as fases 1-19 seguem o loop trifásico:
1. **GENERATE** — Executa prompt gerador com inputs acumulados
2. **GATE** — Auditoria interna. Se erros → FIX → GATE. Se OK → Validação Humana
3. **VALIDAÇÃO HUMANA** — 3 perguntas obrigatórias. Aprova → COMPLIANCE. Novos inputs → volta ao GENERATE

### 4.3 Tratamento de Fases Especiais

- **F19 (Discovery Técnico):** Segue Generate→Gate→Fix, mas ao final de cada sprint o orquestrador pergunta se deseja continuar para próxima sprint ou encerrar
- **EXECUTION-HISTORY:** Generate → Revisão Humana direta (sem gate automatizado)

---

## 5. Estrutura Final de Diretórios (v5.0)

```
business-inputs/business-projects/{PROJECT_ID_NAME}/
└── technical-definitions/
    ├── PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md              (F1)  🆕
    ├── PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT.md          (F2)  🆕
    ├── PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md    (F3)  🆕
    ├── PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md          (F4)  🔄 movido do Bloco C
    ├── PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md         (F5)
    ├── PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md           (F6)
    ├── PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY-EXCEPTIONS.md
    ├── PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md (F7)
    ├── PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md     (F8)
    ├── PROJECT-TECHNICAL-DEFINITIONS-DATA-ARCHITECTURE-DEFINITION.md   (F9)  🆕
    ├── PROJECT-TECHNICAL-DEFINITIONS-DEVOPS-SRE-DEFINITION.md          (F10) 🆕
    ├── PROJECT-TECHNICAL-DEFINITIONS-TEST-STRATEGY-DEFINITION.md       (F11) 🆕
    ├── PROJECT-TECHNICAL-DEFINITIONS-INFRA-CLOUD-DEFINITION.md         (F12) 🆕
    ├── PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md       (F13)
    ├── PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md        (F14)
    ├── PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md  (F15)
    ├── PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md        (F16)
    ├── PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md              (F17)
    ├── technical-discovery/                                      🆕 substitui sprints/
    │   ├── SPRINT-BACKLOG.md                                    (F18) 🔄
    │   ├── sprint-00/
    │   │   ├── CONTRACTS-API-sprint-00.md
    │   │   ├── CONTRACTS-DATA-sprint-00.md
    │   │   ├── CONTRACTS-SECURITY-sprint-00.md
    │   │   ├── CONTRACTS-SRE-sprint-00.md
    │   │   └── DEFINITION-INCREMENTS-sprint-00.md
    │   └── sprint-01/ ...
    ├── PROJECT-TECHNICAL-DEFINITIONS-EXECUTION-HISTORY.md
    └── sprints/  → (substituído por technical-discovery/)
```

---

## 6. Estimativa de Prompts

| Bloco | Fases | Generate | Gate | Fix | Total | Notas |
|:---|:---|:---|:---|:---|:---|:---|
| Bootstrap | F0 | — | — | — | — | Orquestração |
| Bloco 0 | F1-F4 | 4 | 4 | 4 | **12** | 🆕 Todos novos |
| Bloco A | F5-F6 | 2 | 2 | 2 | **6** | Existente |
| Bloco B | F7-F12 | 6 | 6 | 6 | **18** | 2 existentes + 12 🆕 |
| Bloco C | F13-F17 | 5 | 5 | 5 | **15** | Reorganizado |
| Bloco D | F18-F19 | 2 | 2 | 2 | **6** | 🆕 Todos novos |
| History | standalone | 1 | — | — | **1** | Existente |
| **Total** | **20 fases** | **20** | **19** | **19** | **58** | |

**Novos prompts a criar:** 30 (10 Generate + 10 Gate + 10 Fix)
- Bloco 0: INTAKE-LOG, DOR-ASSESSMENT, PRODUCT-BACKLOG-LIST (3×3 = 9)
- Bloco B: DATA-ARCH, DEVOPS-SRE, TEST-STRATEGY, INFRA-CLOUD (4×3 = 12)
- Bloco D: SPRINT-BACKLOG-REFINED, TECHNICAL-DISCOVERY (2×3 = 6)
- PRD-DEFINITION (F4): 3 prompts movidos/atualizados do Bloco C

**Prompts existentes a atualizar:** 28 (reorganização, novos inputs, renomeação de referências)

---

## 7. Decisões de Design

| Decisão | Escolha | Justificativa |
|:---|:---|:---|
| Paralelismo | Nenhum — totalmente sequencial | Bloco B referencia Bloco A; Bloco C referencia todos anteriores; simplicidade operacional |
| PRD-DEFINITION | Negócio (Bloco 0), congelado | SPECS-DEFINITION no Bloco C é a consolidação técnica. PRD de Negócio não é reaberto |
| Bloco D pós-Bloco C | Sequencial | Discovery Técnico precisa de SPECS e MILESTONES definidos |
| SPECS-DEFINITION | Documento enxuto de consolidação | Referencia artefatos anteriores, não repete conteúdo |
| Barreira C feedback loop | Propõe reabertura Bloco A | Skills gaps detectados tardiamente podem ser corrigidos |
| sprints/ → technical-discovery/ | Substituição completa | Modelo SPRINT-BACKLOG.md preservado e enriquecido com coluna CONTRACTS |
| Tech Lead + Principal Arch | Revisores transversais, sem artefato | Participam de todas as fases do Bloco B, assinam validação cruzada |

---

## 8. Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Design inicial v5.0: 5 mudanças estruturais aprovadas | Time de Arquitetura / Claude |

---

🤖 *Design document gerado como parte do processo de brainstorming para o redesign do roadmap de definições técnicas v5.0.*
