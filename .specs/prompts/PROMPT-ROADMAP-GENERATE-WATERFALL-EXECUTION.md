# PROMPT: ROADMAP DE EXECUÇÃO E CONSTRUÇÃO — FASE 5 DO WATERFALL (WATERFALL-EXECUTION)
## Versão: 1.0 — WATERFALL Orchestrator v3.0 (6 Fases, 38 Documentos)

Atue como um Arquiteto de Soluções Organizacionais e Gestor de Execução, especializado em metodologia WATERFALL, engenharia de prompts e orquestração de esteiras de desenvolvimento.

## Propósito e Posição no Fluxo

Este roadmap orquestra a **FASE 5 — EXECUÇÃO E CONSTRUÇÃO** do fluxo WATERFALL. É a fase em que o time de desenvolvimento **recebe toda a documentação produzida nas FASES 1–4** e executa o projeto em ciclos de entrega.

```
FASE 0 (Bootstrap) → FASE 1 (Negócio) → FASE 2 (Especificação) → FASE 3 (Engenharia) → FASE 4 (Baseline)
    → 🚩 M4: PROJECT BASELINE LOCKED
    → FASE 5 (EXECUÇÃO E CONSTRUÇÃO — ESTE ROADMAP)
    → 🚩 M5: GO-LIVE & HANDOVER
    → FASE 6 (ENCERRAMENTO E OPERAÇÃO — docs 105/110/115)
```

- **Entrada:** M4 travado (baseline de escopo, cronograma e orçamento selados)
- **Saída:** M5 (GO-LIVE & HANDOVER) — com 095-RELATORIO-QUALIDADE alimentado com evidências
- **Posição:** orquestrador da FASE 5 do roadmap master `PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md`

## Regras de Ouro de Execução (NÃO NEGOCIÁVEIS)

1. **M4 TRAVADO:** execução só inicia com todos os documentos F1–F4 em `[STATUS: COMPLIANCE]`. Mudança de escopo pós-M4 passa OBRIGATORIAMENTE pelo 085-PLANO-GESTAO-MUDANCAS
2. **NÃO ADIVINHAR INPUTS:** nenhum prompt infere, busca ou descobre inputs — todo parâmetro é passado explicitamente
3. **TODO GENERATE TEM GATE+FIX:** vale para os triads 092/093 e para os artefatos de ciclo (GATEs delegados dos prompts reusados)
4. **FIX É CIRÚRGICO:** correção apenas nas seções com violações reportadas
5. **HITL OBRIGATÓRIO POR CICLO:** a revisão humana do loop de execução (code review) é obrigatória — nenhum ciclo avança sem aprovação explícita do humano
6. **VOCABULÁRIO WATERFALL:** vetar termos ágeis (Epic/User Story/DoR/Sprint) nos artefatos WATERFALL — usar a tabela de tradução (abaixo)
7. **BASELINE PRESERVADA:** 092 opera o 088, mas o 088 permanece o registro congelado em M4 — mudanças são registradas com rastreabilidade, nunca sobrescritas silenciosamente

## VOCABULÁRIO WATERFALL (tabela de tradução obrigatória)

| Termo ágil (PROIBIDO) | Equivalente WATERFALL (usar) |
|---|---|
| Epic | Pacote de trabalho da EAP (060-EAP-WBS) |
| Feature | Funcionalidade `FEAT-NN` (010-FRD) |
| User Story | Caso de Uso `UC-NN` (010-FRD) |
| Definition of Ready (DoR) | GATE de COMPLIANCE do documento de origem |
| Sprint | Ciclo de entrega `FILA-NN` (definido pelo 092) |
| Product Backlog | 088-PRODUCT-BACKLOG-LIST (operado pelo 092 na FASE 5) |

---

## Inputs da Fase 5 — "Time recebe tudo pronto" (UPSTREAM F1–F4)

O time de desenvolvimento recebe, no dia 1 da FASE 5:

| Categoria | Documentos |
|---|---|
| Negócio (F1–F2) | 001-Charter, 002-Stakeholder-Map, 003-Personas-Jornadas, 004-AS-IS/TO-BE, 005-BRD, 010-FRD, 015-RTM-FASE-1 |
| Funcional (F1–F2) | 010-FRD, 016-Protótipos UX/UI, 020-SRS, 025-RTM-FASE-2 |
| Design Técnico (F3) | 030-SAD, 035-HLD, 040-LLD, 041-DEVOPS-SETUP, 042-DATA-SETUP, 043-SEC-SETUP, 044-INFRA-SETUP, 045-EST-PLAN, 050-EST-CASES, 060-EAP-WBS |
| Planejamento e Regras (F4) | 062-STAFFING-PLAN, 065-Cronograma, 070-Orçamento, 075-Comunicação, 080-Riscos, 085-Mudanças, 086-Padrões/DoD, 087-CI-CD, 088-Backlog, 090-STRATEGIC-IMPLEMENTATION-AND-DEPLOYMENT-PLAN |

**Regra:** o orquestrador só inicia a FASE 5 com M4 travado e 095-RELATORIO-QUALIDADE com a estrutura de métricas criada na FASE 3 (as evidências são alimentadas ao longo da FASE 5).

---

## Sub-fase 1 — Gestão Diária & Operacional

Triads WATERFALL próprios (em `project-documents-waterfall/`):

| Doc | Prompt | Responsabilidades |
|---|---|---|
| 092-BACKLOG-KANBAN | `PROMPT-GENERATE/GATE/FIX-092-BACKLOG-KANBAN.md` | Revisa/expande o 088 via **Change-Request de Negócio** e **Change-Request Técnico** (formalizadas pelo 085); atualiza **status** (A Fazer → Em Execução → Em Revisão → Concluído/Impedido); define **Filas/Ciclos** (`FILA-NN`) de implementação |
| 093-GESTAO-TIMES | `PROMPT-GENERATE/GATE/FIX-093-GESTAO-TIMES.md` | Capacidade, impedimentos (`IMP-NN`) e alocação contra o 062-STAFFING-PLAN |

O par 092 ↔ 093 roda em loop contínuo durante toda a FASE 5 (como no flowchart: `092 → 093 → 092`).

## Sub-fase 2 — Janelas de Entrega (Ciclos/Sprints)

> 🚧 **TBD — FORA DE ESCOPO nesta revisão.** A estrutura das janelas (DEV → QA → UAT → DEPLOY) permanece como no `flowchart-WATERFALL.md`, mas NENHUMA solução de prompt/orquestração foi definida. A esteira da sub-fase 3 executa por **ciclo de entrega** (`FILA-NN` do 092), sem depender da definição das janelas.

## Sub-fase 3 — Esteira de Construção por Solução Técnica

Implementa o Diagrama 2 do flowchart ("Esteira Mão-na-Massa — IA + Dev Execution Flow"), **reusando prompts existentes sem edição**:

### 3.1 Contexto Base (uma vez por solução técnica)

- Executar `sprint-artefacts/PROMPT-ORCHESTRATOR-GENERATE-ALL-ARTEFACTS.md` com `IN_MACRO` composto dos docs WATERFALL: 010-FRD, 016-Prodótipos, 020-SRS, 030-SAD, 035-HLD, 040-LLD, 045-EST-PLAN, 050-EST-CASES, 086-Padrões/DoD e 087-CI-CD
- Gera no repositório da solução: `PRD.md/SPECS.md`, `ARCH.md/LLD.md`, `TEST_PLAN.md`, `TASKS.md`

### 3.2 Loop por Ciclo de Entrega (para cada FILA-NN ativa do 092)

| STEP | Ação | Prompt |
|---|---|---|
| 0 | Carregar SPRINT-CARD do ciclo | `sprint-artefacts/PROMPT-GENERATE-SPRINT-ARTEFACTS.md` (gera SPRINT-CARD.md + SPRINT-TEST-SUITE.md a partir de TASKS/TEST_PLAN/SPECS/ARCHITECTURE) |
| 1 | GATE dos artefatos do ciclo | `PROMPT-GATE-SPRINT-ARTEFACTS.md` → `{PASS}` ou `{FAIL, VIOLATIONS[]}` |
| 2 | FIX cirúrgico (máx. 3 loops) | `PROMPT-FIX-SPRINT-ARTEFACTS.md` → re-GATE |
| 3 | Executar tarefas | `sprint-tecnhnical-implementation/PROMPT-EXECUTE-SPRINT-TASKS.md` (consome SPRINT-CARD.md/SPRINT-TEST-SUITE.md; gera código + testes + relatório) |
| 4 | Revisão humana obrigatória (HITL) | Loop `PROMPT_ENG → CODE_GEN → HUMAN_REVIEW` do Diagrama 2 — `sprint-tecnhnical-implementation/PROMPT-QA-REVISOR-SECURITY.md` documenta falhas acionáveis (máx. 3 tentativas) |
| 5 | Fechamento do ciclo | `sprint-tecnhnical-implementation/PROMPT-GENERATE-IDENTIFIED-TECHNICAL-DEBT.md` → `SPRINT-REVIEW.md`; `PROMPT-GENERATE-IMPLEMENTATION-REPORT.md` → `TASK-EXECUTED` |
| 6 | Git e PR | `sprint-tecnhnical-implementation/PROMPT-GENERATE-PULL-REQUEST.md` ou `PROMPT-02-ATUALIZAR-REPOSITORIO-E-ABRIR-PULL-REQUEST.md` |
| 7 | Saída de governança | Atualizar 095-RELATORIO-QUALIDADE (evidências), 092 (status BL-NN/FILA-NN), 093 (impedimentos); desvios acionam 085 |

### 3.3 Regras da esteira

- **Vocabulário:** prompts ágeis reusados são agnósticos de stack; o orquestrador traduz o vocabulário (US → UC-NN do 010; sprint → FILA-NN do 092) na invocação — nunca reescrever os prompts
- **Git:** branch por ciclo `feature/sprint-NN-<slug>` (padrão já embutido nos prompts reusados); merge via PR com code review (086)
- **Débito técnico:** `IDENTIFIED-TECHNICAL-DEBT.md` realimenta `TASKS.md` (novas tarefas entram como CR Técnico no 092)

## Sub-fase 4 — Documentação de Suporte e Evidências

| Doc | Quando | Notas |
|---|---|---|
| 095-RELATORIO-QUALIDADE | estrutura criada na F3; **alimentado com evidências ao longo da FASE 5** | executado a cada ciclo (STEP 7) |
| 097-MANUAIS-USUARIO | durante a execução | upstream 003/010/016 — telas e jornadas reais |
| 100-MANUAIS-OPERACIONAIS | durante a execução | upstream 041/044/087/090 — runbooks reais |

---

## Integração com WATERFALL-ESTIMATION

Se a EAP mudar via 085 (mudança de escopo aprovada), o orquestrador alerta o humano e oferece **reexecutar o WATERFALL-ESTIMATION modo DOWNSTREAM/REFINEMENT** — nova PERT atualiza 065/070 e, consequentemente, o 088 (via 092).

## Git Workflow da Execução

- Branch por ciclo: `feature/sprint-NN-<slug>`; PR obrigatório com code review (086) e CI verde (087)
- O commit de encerramento (git add/commit/push/PR do roadmap master) migra para o final da FASE 6

## Localização dos Prompts

```
.specs/prompts/
├── PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md      ← roadmap master (6 fases)
├── PROMPT-ROADMAP-GENERATE-WATERFALL-EXECUTION.md              ← ESTE ROADMAP (FASE 5)
├── project-documents-waterfall/                                ← triads WATERFALL (092, 093 e demais docs)
├── sprint-artefacts/                                           ← reuso: contexto base + artefatos de ciclo
├── sprint-tecnhnical-implementation/                           ← reuso: execução, QA-revisor, PR, débito técnico
└── technical-solutions/                                        ← reuso: tarefas por solução técnica
```

---

🤖 *Roadmap gerado pelo Waterfall Orchestrator v3.0. Evolução aprovada em 14/08/2026 — sub-fase 2 (Janelas de Entrega) permanece TBD/fora de escopo.*
