# PROMPT: ROADMAP DE DEFINIÇÕES TÉCNICAS DO PROJETO
## Versão: 7.4 — Modos de Execução (agile/waterfall-discovery) + 6 Disciplinas Técnicas + Discovery Contínuo + Bloco E (Esteira de Construção + Pacote de Desenvolvimento — orquestrador 1000, 18 fases) + Bloco F (Janelas de Entrega + Tooling de Ambiente) + Ciclos de Entrega CICLO-NN + 590-ciclo-NNN

Atue como um Especialista em Gestão de Processos (BPM), Arquiteto de Soluções Ágeis e Tech Lead, especializado em definições técnicas de projetos e engenharia de prompts.

Preciso que você crie um roadmap de execução detalhado e um guia de orquestração para o seguinte processo: Criação, revisão, evolução e validação dos **documentos de definição técnica do projeto** — artefatos que preenchem o gap entre os documentos de negócio (`PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS.md`) e as especificações técnicas por solução (`PROMPT-ROADMAP-GENERATE-TECHNICAL_SOLUTIONS.md`, na pasta `technical-solutions/`).

Objetivo Principal: Garantir que todas as definições técnicas do projeto estejam criadas, revisadas e 100% alinhadas conceitualmente entre si e com os documentos de negócio (Charter, BRD, Epics, Features, User Stories), preparando o terreno para que cada time de solução técnica inicie seu trabalho com baseline consistente.

Regra Crítica de Execução (Gating Rule): O processo é estritamente sequencial em todas as fases. Nenhuma fase subsequente pode ser iniciada sem a aprovação formal, soberana e explícita do usuário humano na fase anterior.

---

## VARIÁVEIS DE ENTRADA E BOOTSTRAP (FASE 0)

### Tabela de Inputs

| Variável | Obrig. | Descrição | Exemplo |
|---|---|---|---|
| `PROJECT_PATH` | ✅ | Caminho base onde os projetos de negócio residem | `/home/bolismar/work/workspace-fbso/business-inputs/business-projects` |
| `PROJECT_ID_NAME` | ✅ | Identificador completo do projeto (ID + Nome) | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `TECHNICAL_SOLUTION_PATH` | ✅ | Caminho base onde as soluções técnicas residem | `/home/bolismar/work/workspace-fbso/backend/java/spring/microservices` |
| `TECHNICAL_SOLUTION_NAMES` | ✅ | Lista de nomes das soluções técnicas do projeto | `["ms-fbso-platform-admin", "web-app-fbso-platform-portal"]` |
| `ARCHITECTURE_GLOBAL` | ✅ | Caminho para a pasta de arquitetura global (ADRs, blueprints, padrões) | `/home/bolismar/work/workspace-fbso/architecture/` |
| `SECURITY_GLOBAL` | ✅ | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) | `/home/bolismar/work/workspace-fbso/.specs/security/GLOBAL-SECURITY.md` |
| `PROJECT_DOCUMENTS_INPUTS` | ❌ | Lista de caminhos para documentos brutos de entrada adicionais | `[]` |
| `PROJECT_PROMPT_INPUTS` | ❌ | Lista de caminhos para prompts auxiliares ou contextos adicionais | `[]` |
| `TECHNICAL_DEFINITIONS_MODE` | ❌ | Modo de execução: `agile-discovery` ou `waterfall-discovery`. Se não informado, o Bootstrap detecta e pergunta ao humano | `agile-discovery` |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
TECHNICAL_DEFINITIONS_PATH    = PROJECT_COMPLETE_PATH_NAME + "/technical-definitions"
```

---

## MODOS DE EXECUÇÃO

| Modo | Pipeline | Baseline | Quando |
|---|---|---|---|
| `agile-discovery` | Pipeline atual (Blocos 0–A–B–C–D, 20 fases) | Docs de negócio ágeis (`project-documents/`: Charter, BRD, Epics, Features, US) | Projetos ágeis — comportamento atual, retrocompatível |
| `waterfall-discovery` 🆕 | Bloco 0 reduzido (F1–F2 puladas, F3–F4 migradas) + Blocos A–D (migrados/validados) + **Bloco E** (Esteira de Construção) + **Bloco F** (Janelas de Entrega — consome o 096) + **595-TECHLEAD-RETURN-PACKAGE** | Docs WATERFALL F1–F4 em `[STATUS: COMPLIANCE]` (M4 travado: 088, 092, 010-FRD, 060-EAP-WBS, 062, 065, 070, 086, 087, 090) + `096-DEFINICAO-JANELAS-ENTREGA` | Projetos WATERFALL na FASE 5 (parceria com `PROMPT-ROADMAP-GENERATE-WATERFALL-EXECUTION.md` v2.3) |

**Detecção no Bootstrap (auditoria de artefatos):**

| Sinal auditado | Modo proposto |
|---|---|
| `project-documents/` ágil (Charter/BRD/Epics/Features/US) + `features/` + `user-stories/` | `agile-discovery` |
| Docs WATERFALL `088-PRODUCT-BACKLOG-LIST` + `092-BACKLOG-KANBAN` (com `CICLO-NN`) + `010-FRD` + `060-EAP-WBS` + 062/065/070/086/087/090 em `[STATUS: COMPLIANCE]` | `waterfall-discovery` |
| Nenhum dos dois | Perguntar ao humano; iniciar em `agile-discovery` |

A decisão final é sempre do humano (apresentar o modo proposto e confirmar).

---

## ARQUITETURA DE FASES

O roadmap é organizado em **20 fases** agrupadas em **7 blocos**:

```
FASE 0: BOOTSTRAP (sequencial)
  │
  ├─▶ BLOCO 0: Product Definition & Product Backlog & PRD (NOVO)
  │     Fase 1 → Fase 2 → Fase 3 → Fase 4
  │     ⛔ Barreira 0
  │
  ├─▶ BLOCO A: People & Solutions
  │     Fase 5 → Fase 6
  │     ⛔ Barreira A
  │
  ├─▶ BLOCO B: Architecture & Security & Specialists (EXPANDIDO)
  │     Fase 7 → Fase 8 → Fase 9 → Fase 10 → Fase 11 → Fase 12
  │     ⛔ Barreira B
  │
  ├─▶ BLOCO C: Catálogo, Matriz, Stack, Specs & Milestones (REORGANIZADO)
  │     Fase 13 → Fase 14 → Fase 15 → Fase 16 → Fase 17
  │     ⛔ Barreira C (com feedback loop → Bloco A)
  │
  ├─▶ BLOCO D: Ciclos/Sprints — Technical Discovery (REPROPOSITADO)
  │     Fase 18 → Fase 19 (iterativo)
  │     ⛔ Barreira D
  │
  └─▶ BLOCO E: Esteira de Construção por Ciclo (SOMENTE modo waterfall-discovery)
        Contexto base → Loop por CICLO-NN → 595-TECHLEAD-RETURN-PACKAGE → PM/PO
  └─▶ BLOCO F: Janelas de Entrega (SOMENTE modo waterfall-discovery)
        DEV (Bloco E) → QA → UAT → DEPLOY por CICLO-NN — consome o 096-DEFINICAO-JANELAS-ENTREGA

  └─▶ EXECUTION-HISTORY 📊 (standalone)
```

---

## FASES DO ROADMAP

### Fase 0 — Bootstrap Inteligente
(Mantido igual — solicitar inputs, criar estrutura, auditar artefatos)

Workflow:
1. Solicitar inputs ao usuário (se não fornecidos)
2. Exibir caminhos derivados e solicitar confirmação
3. Criar estrutura: `mkdir -p {TECHNICAL_DEFINITIONS_PATH}`
4. Criar template `465-TEAM-CAPACITY-EXCEPTIONS.md`
5. Auditar artefatos existentes — varredura completa nos documentos para identificar a situação atual do processo de criação da documentação
5a. **Detectar modo de execução** pela auditoria (tabela de detecção da seção MODOS DE EXECUÇÃO) e confirmar com o humano (`agile-discovery` ou `waterfall-discovery`)
6. Apresentar resumo da situação atual (incluindo o modo proposto) e iniciar a primeira fase pendente (ou ponto de retomada)

### Fase 1 — 410-INTAKE-LOG.md 🆕
Registro de lotes de ingestão de requisitos do Negócio. Pipeline: `PROMPT-GENERATE-410-INTAKE-LOG.md` → Gate → Fix → COMPLIANCE

### Fase 2 — 420-DOR-ASSESSMENT.md 🆕
Aplicação do DoR de Negócio (PO/PM). Pipeline: `PROMPT-GENERATE-420-DOR-ASSESSMENT.md` → Gate → Fix → COMPLIANCE

### Fase 3 — 430-PRODUCT-BACKLOG-LIST.md 🆕
Backlog consolidado "Pronto para TI". Pipeline: `PROMPT-GENERATE-430-PRODUCT-BACKLOG-LIST.md` → Gate → Fix → COMPLIANCE

### Fase 4 — 440-PRD-DEFINITION.md 🔄
PRD de Negócio (movido do Bloco C). Pipeline: `PROMPT-GENERATE-440-PRD-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 5 — 450-TEAM-SKILLS-MAP.md
Skills matrix do Discovery Team. Pipeline: `PROMPT-GENERATE-450-TEAM-SKILLS-MAP.md` → Gate → Fix → COMPLIANCE

### Fase 6 — 460-TEAM-CAPACITY.md
Capacidade de trabalho do time. Pipeline: `PROMPT-GENERATE-460-TEAM-CAPACITY.md` → Gate → Fix → COMPLIANCE

### Fase 7 — 470-ARCHITECTURE-DEFINITION.md
Solution Architect — integração C4, ADRs, topologia. Pipeline: `PROMPT-GENERATE-470-ARCHITECTURE-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 8 — 480-SECURITY-DEFINITION.md
Security Architect — threat model, IAM, compliance. Pipeline: `PROMPT-GENERATE-480-SECURITY-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 9 — 490-DATA-ARCHITECTURE-DEFINITION.md 🆕
Data Architect — modelagem, pipelines, storage strategy. Pipeline: `PROMPT-GENERATE-490-DATA-ARCHITECTURE-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 10 — 500-DEVOPS-SRE-DEFINITION.md 🆕
DevOps/SRE Architect — CI/CD, IaC, observabilidade, SLOs. Pipeline: `PROMPT-GENERATE-500-DEVOPS-SRE-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 11 — 510-TEST-STRATEGY-DEFINITION.md 🆕
Test Specialist — pirâmide de testes, automação, performance. Pipeline: `PROMPT-GENERATE-510-TEST-STRATEGY-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 12 — 520-INFRA-CLOUD-DEFINITION.md 🆕
Infra/Cloud Specialist — topologia, compute, networking, DR. Pipeline: `PROMPT-GENERATE-520-INFRA-CLOUD-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 13 — 530-SOLUTIONS-CATALOG.md
Catálogo de soluções técnicas. Pipeline: `PROMPT-GENERATE-530-SOLUTIONS-CATALOG.md` → Gate → Fix → COMPLIANCE

### Fase 14 — 540-SOLUTIONS-MATRIX.md
Matriz solução×disciplina×owner. Pipeline: `PROMPT-GENERATE-540-SOLUTIONS-MATRIX.md` → Gate → Fix → COMPLIANCE

### Fase 15 — 550-SOLUTIONS-STACK-MATRIX.md
Stacks tecnológicas por solução. Pipeline: `PROMPT-GENERATE-550-SOLUTIONS-STACK-MATRIX.md` → Gate → Fix → COMPLIANCE

### Fase 16 — 560-SPECS-DEFINITION.md
Consolidação técnica enxuta — sumariza e referencia artefatos anteriores. Pipeline: `PROMPT-GENERATE-560-SPECS-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 17 — 570-MILESTONES.md
Roadmap alinhado ao negócio. Pipeline: `PROMPT-GENERATE-570-MILESTONES.md` → Gate → Fix → COMPLIANCE

### Fase 18 — technical-discovery/580-PACKAGE-BACKLOG-REFINED.md 🆕
Backlog refinado com tarefas T-NNN → US-ID → Ciclo/Sprint-Alvo → CONTRACTS. Gerado em `technical-discovery/`. Pipeline: `PROMPT-GENERATE-580-PACKAGE-BACKLOG-REFINED.md` → Gate → Fix → COMPLIANCE

### Fase 19 — technical-discovery/590-ciclo-NNN/ 🆕
Discovery Técnico Contínuo — contratos API/Data/Security/SRE por ciclo/sprint + increments. Iterativo. Gera a estrutura `technical-discovery/590-ciclo-NNN/`. Pipeline: `PROMPT-GENERATE-590-TECHNICAL-DISCOVERY.md` → Gate → Fix → COMPLIANCE

### Execution History — 600-EXECUTION-HISTORY.md 📊
Dashboard de controle — estado de todos os documentos. Pipeline: Generate → Revisão humana (sem gate próprio).

---

## MODO WATERFALL-DISCOVERY (FASE 5 DO WATERFALL — PARCERIA PM/PO × TECHLEAD)

No modo `waterfall-discovery`, este roadmap atua como o **lado TECHLEAD da FASE 5**: recebe o pacote de demanda do PM/PO (via WATERFALL-EXECUTION v2.0), valida/refina as definições técnicas contra os docs WATERFALL, executa a esteira de construção (Bloco E) e devolve o `595-RETURN-PACKAGE-{CICLO-NN}.md`.

**Ownership:** "TECHLEAD propõe, PM/PO aplica" — este roadmap NUNCA edita 092/093/095/085/088; atualiza apenas os artefatos de `technical-definitions/`, `technical-discovery/`, o repositório da solução e o 600-EXECUTION-HISTORY.

### Mapeamento Fase a Fase (modo waterfall-discovery)

| Fase | Artefato | Comportamento | Inputs traduzidos |
|---|---|---|---|
| F1 | 410-INTAKE-LOG | **PULAR** | Ingestão já feita nas FASES 1–4; baseline = pacote de demanda do PM/PO |
| F2 | 420-DOR-ASSESSMENT | **PULAR** | DoR → GATE de COMPLIANCE do M4 |
| F3 | 430-PRODUCT-BACKLOG-LIST | **MIGRADO** | 088-PRODUCT-BACKLOG-LIST → 430 |
| F4 | 440-PRD-DEFINITION | **MIGRADO** | 010-FRD (FEAT-NN/UC-NN) → 440 |
| F5 | 450-TEAM-SKILLS-MAP | **MIGRADO** | 062-STAFFING-PLAN → 450 |
| F6 | 460-TEAM-CAPACITY | **MIGRADO** | 062 + 065/070 → 460 |
| F7 | 470-ARCHITECTURE-DEFINITION | **VALIDA/REFINA** | 030-SAD + 035-HLD → 470 |
| F8 | 480-SECURITY-DEFINITION | **VALIDA/REFINA** | 043-SEC-SETUP → 480 |
| F9 | 490-DATA-ARCHITECTURE-DEFINITION | **VALIDA/REFINA** | 042-DATA-SETUP → 490 |
| F10 | 500-DEVOPS-SRE-DEFINITION | **VALIDA/REFINA** | 041-DEVOPS-SETUP + 087 → 500 |
| F11 | 510-TEST-STRATEGY-DEFINITION | **VALIDA/REFINA** | 045-TEST-PLAN + 050-TEST-CASES → 510 |
| F12 | 520-INFRA-CLOUD-DEFINITION | **VALIDA/REFINA** | 044-INFRA-SETUP + 090 → 520 |
| F13 | 530-SOLUTIONS-CATALOG | **RODA** | 060-EAP-WBS + 010 |
| F14 | 540-SOLUTIONS-MATRIX | **RODA** | 062 + 086 |
| F15 | 550-SOLUTIONS-STACK-MATRIX | **RODA** | 087 + 044 + STACK-PADROES-CORPORATIVOS |
| F16 | 560-SPECS-DEFINITION | **RODA** | consolida Blocos 0–B |
| F17 | 570-MILESTONES | **MIGRADO** | 065-CRONOGRAMA-GANTT + 070-ORCAMENTO → 570 |
| F18 | 580-PACKAGE-BACKLOG-REFINED | **MIGRADO/RODA** | 092 (CICLO-NN/BL-NN) → 580 (tarefas T-NNN) |
| F19 | 590-TECHNICAL-DISCOVERY (iterativo) | **RODA** | por ciclo: pacote de demanda (CICLO-NN) → 5 contratos por 590-ciclo-NNN |
| 600 | 600-EXECUTION-HISTORY | **RODA** | registro interno do TECHLEAD |

**Regras de vocabulário:** US → UC-NN (010-FRD), sprint → CICLO-NN (092), DoR → GATE de COMPLIANCE, Epic → pacote EAP (060), Product Backlog → 088. **Regra de espelhamento:** `590-ciclo-NN ↔ CICLO-NN`.

### BLOCO E — Esteira de Construção por Ciclo (após Barreira D, somente no modo waterfall-discovery)

**E1 — Contexto base (uma vez por solução técnica):** reuso sem edição de `sprint-artefacts/PROMPT-ORCHESTRATOR-GENERATE-ALL-ARTEFACTS.md` com `IN_MACRO` = [010-FRD, 016-Protótipos, 020-SRS, 030-SAD, 035-HLD, 040-LLD, 045-TEST-PLAN, 050-TEST-CASES, 086, 087] → gera `PRD.md/SPECS.md`, `ARCH.md/LLD.md`, `TEST_PLAN.md`, `TASKS.md` no repositório da solução.

**E2 — Loop por ciclo (para cada CICLO-NN ativa):**

| STEP | Ação | Prompt (reuso sem edição) |
|---|---|---|
| 0 | SPRINT-CARD + SPRINT-TEST-SUITE | `sprint-artefacts/PROMPT-GENERATE-SPRINT-ARTEFACTS.md` |
| 1 | GATE dos artefatos do ciclo | `sprint-artefacts/PROMPT-GATE-SPRINT-ARTEFACTS.md` |
| 2 | FIX cirúrgico (máx. 3 loops) | `sprint-artefacts/PROMPT-FIX-SPRINT-ARTEFACTS.md` |
| 3 | **Executar o pacote de desenvolvimento (18 fases)** | `sprint-tecnhnical-implementation/PROMPT-EXECUTE-1000-PACKAGE-DEVELOPMENT-ORCHESTRATOR.md` — fases internas: 1010 débito pré-ciclo → 1020 pré-impl → 1030 plano → 1040 código → 1050 CVE/SCA + 1060 CI-CD (condicionais) → 1070–1100 testes/qualidade → 1110 falhas → 1120 code review → 1130 QA gate (HITL, 3 loops) → 1140–1170 pós/relatórios/artefatos → 1180 PR |
| 4 | (FALLBACK legado — somente se o orquestrador 1000 não se aplicar) Execução monolítica | `sprint-tecnhnical-implementation/PROMPT-EXECUTE-SPRINT-TASKS.md` + `PROMPT-QA-REVISOR-SECURITY.md` + `PROMPT-GENERATE-PULL-REQUEST.md` |
| 5 | **Empacotamento (em vez de governança direta)** | trio **595-TECHLEAD-RETURN-PACKAGE** (GENERATE → GATE → FIX → COMPLIANCE) → entrega `595-RETURN-PACKAGE-{CICLO-NN}.md` ao PM/PO |

**E3 — Regras:** vocabulário WATERFALL na invocação (nunca reescrever prompts reusados); branch `feature/sprint-NN-<slug>`; **padrão = orquestrador 1000** (18 fases numeradas na ordem de execução, artefato `PACKAGE-DEVELOPMENT-{FASE}.md` por fase; HITL interno = fase 1130, gate QA/segurança com 3 loops; débito centralizado na fase 1010); o legado monolítico `PROMPT-EXECUTE-SPRINT-TASKS.md` permanece como fallback; débito técnico `DT-XXX` entra no pacote 595 como CR Técnico; ownership (nunca editar 092/093/095/085/088). O Bloco E roda por `CICLO-NN` **dentro da janela DEV** — a passagem pelas janelas (QA/UAT/DEPLOY) é orquestrada pelo **Bloco F** (abaixo), conforme o 096-DEFINICAO-JANELAS-ENTREGA.

### BLOCO F — Janelas de Entrega (após Barreira D, somente no modo waterfall-discovery)

**F.1 — Visão das Janelas:** consome o **096-DEFINICAO-JANELAS-ENTREGA** (trio em `project-documents-waterfall/`) como upstream. As 4 janelas — DEV, QA, UAT, DEPLOY — regem a passagem de cada ciclo `CICLO-NN`. Regra: **janela ≠ ciclo** — a `CICLO-NN` vem do 092; a janela é o estágio de passagem da entrega.

**F.2 — Orquestração por ciclo CICLO-NN:**

`JAN-DEV (Bloco E steps 0–7) → gate DEV → JAN-QA (050 + QA-REVISOR-SECURITY + 095 GO/NO-GO) → JAN-UAT (DE-ACORDO/APROVAÇÃO por entrega — Key Users + PM/PO) → JAN-DEPLOY (090 + 087, GMUD em PROD) → pacote 595 ao PM/PO → PM/PO aplica no 092 → próxima CICLO-NN`

Tratativas de retorno (do 096): QA NO-GO → volta à DEV; UAT com divergência → CR via 085 + volta à DEV; DEPLOY bloqueado → `IMP-NN` (093) + 085.

**F.3 — Delegação de execução por frente (SEM prompts executores novos):**

| Janela | Frente responsável | Como executará (delegação futura) |
|---|---|---|
| DEV | TECHLEAD | Bloco E (prompts reusados atuais); roadmaps/skills/agentes de execução quando existirem |
| QA | TECHLEAD | QA-REVISOR-SECURITY + 050; skill/agente de QA futuro (funcional/carga/pentest); resultado proposto no 595 |
| UAT | PM/PO + usuários de negócio | Registro de DE-ACORDO/APROVAÇÃO por entrega; sem executor automatizado — validação humana |
| DEPLOY | TECHLEAD | 090 + 087; skill/agente de deploy futuro (GMUD em PROD) |

**F.4 — Rastreio por janela:** o 600-EXECUTION-HISTORY registra cada transição (`JAN-*-NN`, data, responsável, evidências); o pacote 595 (GENERATE) carrega a coluna "Janela" na seção 2, o **Aceite UAT (DE-ACORDO)** por item e as evidências agrupadas por janela na seção 4; o GATE-595 valida via `[595-08]`/`[595-09]`/`[595-10]`.

**F.5 — Regras:** vetos do 092 preservados (o 092 não define janelas — nenhuma edição); "TECHLEAD propõe, PM/PO aplica" por janela:

| Janela | TECHLEAD propõe | PM/PO aplica |
|---|---|---|
| DEV | conclusão do ciclo (evidências do Bloco E) | revisor humano aceita (086); ciente via 595 |
| QA | resultado da verificação | status no 092 + validação do 095 (GO/NO-GO) |
| UAT | — (não executa) | registro de DE-ACORDO/APROVAÇÃO por entrega |
| DEPLOY | execução (runbook 090) | validação do checklist + go/no-go GMUD |

HITL por transição; 095 GO é pré-requisito do UAT; o 105-TERMO-ACEITE permanece como aceite FINAL do projeto (FASE 6 — nunca gate por entrega); mudança de escopo pós-M4 via 085; vocabulário WATERFALL.

**F.6 — Setup de ambiente e ferramentas (tooling):** tarefas de infra/ferramentas das janelas DEV/QA são orquestradas pelo roadmap companion **`PROMPT-ROADMAP-GENERATE-IMPLEMENTATION-TOOLING.md`** (v1.0, metodologia-independente) — trios 610 (manifests), 620 (observabilidade), 630 (instalação de ferramentas middleware/ETL/orquestração) e 640 (ferramentas de segurança), invocado pelo Bloco F. Os manifests da 610 alimentam o `PROMPT-EXECUTE-CI-CD-PIPELINE` (step 3b do Bloco E). HITL por ambiente preservado; em contexto WATERFALL, HMG/PROD via GMUD (090).

---

## MECANISMO DE ORQUESTRAÇÃO DINÂMICA

Toda fase (1-19) deve rodar sob o ecossistema trifásico de prompts (Gerador, Auditor/Portão e Corretor), com controle final obrigatório do Humano. O fluxo segue estritamente a máquina de estados abaixo, idêntica à do roadmap de documentos de negócio (`PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS.md`):

1. **Geração / Evolução:** A IA recebe os inputs disponíveis e executa o prompt gerador da fase
2. **Auditoria Interna da IA:** O artefato é enviado para o gate. 
   - SE A IA ENCONTRAR ERROS: Emite o status `[NÃO COMPLIANCE]`, apresenta as falhas encontradas com sugestões de tratativa, faz as 3 perguntas obrigatórias do Portão de Validação Humana e aguarda o direcionamento do humano, que decide o que fazer com cada falha. Com base nas respostas, aciona o FIX de forma cirúrgica e retorna ao passo 2.
   - SE A IA NÃO ENCONTRAR ERROS (100% OK): Emite o status `[PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]`, apresenta as 3 perguntas obrigatórias e aguarda o direcionamento humano.
3. **Portão de Validação Humana (3 perguntas obrigatórias):**
   1. O documento está em compliance com a sua necessidade e perfeitamente alinhado com os documentos base?
   2. Deseja enviar mais documentos/arquivos para enriquecer este artefato?
   3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?
4. **Lógica de Decisão Baseada nas Respostas do Humano:**
   - CENÁRIO DE SUCESSO (Aprovação): Se o humano validar o documento e NÃO enviar novos arquivos ou inputs (Sim, Não, Não), a fase é dada por encerrada (`[STATUS: COMPLIANCE]`), o arquivo é congelado e a próxima fase é destravada.
   - CENÁRIO DE RETROCESSO (Evolução Incremental): Se o humano fornecer novos documentos ou novas informações, o orquestrador DEVE retroceder ao passo 1 (GENERATE), injetando o documento gerado até o momento + os novos insumos para uma atualização incremental.

**Relatório de Falha:** Os gates NÃO geram arquivos de FAIL_REPORT. As não-conformidades são apresentadas inline no próprio relatório do gate como parte do status `[NÃO COMPLIANCE]`.

### Esquema de Estados por Documento

```
CREATED → GATE ⟷ FIX → PRÉ-COMPLIANCE → COMPLIANCE
  │         │              │
  │         └─ NÃO-COMPLIANCE (falhas inline no relatório do gate)
  │
  └─ (estado inicial após geração)
```

---

## REGRAS DE BLOQUEIO (GATING RULES)

### Barreiras de Bloco

| Barreira | Posição | Validação | Regra Especial |
|---|---|---|---|
| ⛔ Barreira 0 | Após Bloco 0 (F4) | Rastreabilidade Backlog→Docs Negócio. DoR 100%. PRD cobre todo backlog. | Itens sem DoR → voltam F2. PRD incompleto → volta F4 |
| ⛔ Barreira A | Após Bloco A (F6) | TEAM-SKILLS-MAP cobre todos os papéis. TEAM-CAPACITY preenchido. | — |
| ⛔ Barreira B | Após Bloco B (F12) | 6 disciplinas OK. N/A justificados. Consistência horizontal. | Disciplina N/A sem justificativa = NÃO COMPLIANCE |
| ⛔ Barreira C | Após Bloco C (F17) | SPECS referencia todos artefatos. MILESTONES alinhado. | **Skills-Gap Detection:** Se skill necessária não coberta pelo Bloco A → `[SKILLS-GAP-DETECTED]` → propõe reabertura Bloco A |
| ⛔ Barreira D | Após Bloco D (F19) | 100% US do ciclo/sprint com contratos. PACKAGE-BACKLOG atualizado. | Iterativo — pergunta se continua no próximo ciclo/sprint ou encerra |

> **Modo waterfall-discovery:** a Barreira 0 valida os MIGRADOS (F3–F4 contra 088/010); as Barreiras A–D permanecem; após a Barreira D o fluxo entra no Bloco E (por ciclo, janela DEV) e no Bloco F (janelas QA/UAT/DEPLOY).

### Consistência Horizontal (Bloco B)

A Barreira B deve validar que os 6 artefatos do Bloco B são consistentes entre si:
- ARCHITECTURE ↔ SECURITY: controles de segurança implementam padrões arquiteturais
- ARCHITECTURE ↔ DATA: modelo de dados alinhado com topologia de containers
- ARCHITECTURE ↔ DEVOPS-SRE: pipeline de deploy suporta a topologia definida
- SECURITY ↔ INFRA-CLOUD: controles de rede e IAM consistentes
- TEST-STRATEGY ↔ ARCHITECTURE: pirâmide de testes cobre a topologia de integração
- TEST-STRATEGY ↔ SECURITY: testes de segurança (SAST/DAST) alinhados com threat model

### Skills-Gap Detection (Barreira C)

Se durante a validação do Bloco C for identificado que uma skill/disciplina não está coberta pelo TEAM-SKILLS-MAP ou TEAM-CAPACITY:

1. Emitir alerta `[SKILLS-GAP-DETECTED]` com:
   - Skill ausente
   - Impacto no projeto
   - Recomendação de ação
2. Propor reabertura do Bloco A para atualização dos documentos People
3. Aguardar decisão humana:
   - Prosseguir sem a skill (risco aceito)
   - Reabrir Bloco A (F5-F6) para atualização
   - Adicionar exceção documentada no TEAM-CAPACITY-EXCEPTIONS.md

---

## ESTRUTURA DE DIRETÓRIOS GERADA

```
business-inputs/business-projects/{PROJECT_ID_NAME}/
└── technical-definitions/
    ├── 410-INTAKE-LOG.md              (F1)  🆕
    ├── 420-DOR-ASSESSMENT.md          (F2)  🆕
    ├── 430-PRODUCT-BACKLOG-LIST.md    (F3)  🆕
    ├── 440-PRD-DEFINITION.md          (F4)  🔄
    ├── 450-TEAM-SKILLS-MAP.md         (F5)
    ├── 460-TEAM-CAPACITY.md           (F6)
    ├── 465-TEAM-CAPACITY-EXCEPTIONS.md
    ├── 470-ARCHITECTURE-DEFINITION.md (F7)
    ├── 480-SECURITY-DEFINITION.md     (F8)
    ├── 490-DATA-ARCHITECTURE-DEFINITION.md   (F9)  🆕
    ├── 500-DEVOPS-SRE-DEFINITION.md          (F10) 🆕
    ├── 510-TEST-STRATEGY-DEFINITION.md       (F11) 🆕
    ├── 520-INFRA-CLOUD-DEFINITION.md         (F12) 🆕
    ├── 530-SOLUTIONS-CATALOG.md       (F13)
    ├── 540-SOLUTIONS-MATRIX.md        (F14)
    ├── 550-SOLUTIONS-STACK-MATRIX.md  (F15)
    ├── 560-SPECS-DEFINITION.md        (F16)
    ├── 570-MILESTONES.md              (F17)
    ├── technical-discovery/
    │   ├── 580-PACKAGE-BACKLOG-REFINED.md                                    (F18)
    │   ├── 590-ciclo-00/
    │   │   ├── CONTRACTS-API-ciclo-00.md
    │   │   ├── CONTRACTS-DATA-ciclo-00.md
    │   │   ├── CONTRACTS-SECURITY-ciclo-00.md
    │   │   ├── CONTRACTS-SRE-ciclo-00.md
    │   │   └── DEFINITION-INCREMENTS-ciclo-00.md
    │   └── 590-ciclo-01/ ...
    ├── 595-RETURN-PACKAGE-{CICLO-NN}.md          (Bloco E — modo waterfall-discovery) 🆕
    └── 600-EXECUTION-HISTORY.md
```

---

## Skills Utilizados

> **📌 Nota sobre Skills:** A tabela abaixo lista os skills **recomendados** para o orquestrador. O agente tem autonomia para selecionar outros skills identificados como mais aderentes às necessidades específicas.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `superpowers:brainstorming` | Brainstorming inicial da arquitetura de definições | Orquestração |
| 2 | `superpowers:executing-plans` | Execução do plano de fases com gates | Orquestração |
| 3 | `superpowers:writing-plans` | Escrita e refino do plano de execução | Orquestração |
| 4 | `superpowers:verification-before-completion` | Verificação de completude antes de cada COMPLIANCE | Qualidade |
| 5 | `workflow-orchestration-patterns` | Padrões de orquestração de pipeline sequencial multi-bloco | Orquestração |
| 6 | `sequential-orchestration` | Orquestração de fases em sequência estrita com gates e barreiras | Orquestração |
| 7 | `gap-analysis` | Análise de gaps entre documentos de negócio e definições técnicas | Análise |
| 8 | `analyze-project` | Análise do projeto existente para bootstrap | Análise |
| 9 | `context-manager` | Gestão de contexto entre fases longas | Contexto |
| 10 | `documentation-writer` | Documentação do roadmap e execution history | Documentação |

> **🔄 Flexibilidade:** Se durante a execução o agente identificar que um skill diferente é mais adequado, substituí-lo e justificar no 600-EXECUTION-HISTORY.md.

---

## Localização dos Prompts das Fases

Os prompts de geração, gate e correção de cada fase estão na pasta `project-technical-definitions/`:

```
.specs/prompts/project-technical-definitions/
├── PROMPT-GENERATE-410-INTAKE-LOG.md            🆕
├── PROMPT-GATE-410-INTAKE-LOG.md                🆕
├── PROMPT-FIX-410-INTAKE-LOG.md                 🆕
├── PROMPT-GENERATE-420-DOR-ASSESSMENT.md         🆕
├── PROMPT-GATE-420-DOR-ASSESSMENT.md             🆕
├── PROMPT-FIX-420-DOR-ASSESSMENT.md              🆕
├── PROMPT-GENERATE-430-PRODUCT-BACKLOG-LIST.md   🆕
├── PROMPT-GATE-430-PRODUCT-BACKLOG-LIST.md       🆕
├── PROMPT-FIX-430-PRODUCT-BACKLOG-LIST.md        🆕
├── PROMPT-GENERATE-440-PRD-DEFINITION.md
├── PROMPT-GATE-440-PRD-DEFINITION.md
├── PROMPT-FIX-440-PRD-DEFINITION.md
├── PROMPT-GENERATE-450-TEAM-SKILLS-MAP.md
├── PROMPT-GATE-450-TEAM-SKILLS-MAP.md
├── PROMPT-FIX-450-TEAM-SKILLS-MAP.md
├── PROMPT-GENERATE-460-TEAM-CAPACITY.md
├── PROMPT-GATE-460-TEAM-CAPACITY.md
├── PROMPT-FIX-460-TEAM-CAPACITY.md
├── PROMPT-GENERATE-470-ARCHITECTURE-DEFINITION.md
├── PROMPT-GATE-470-ARCHITECTURE-DEFINITION.md
├── PROMPT-FIX-470-ARCHITECTURE-DEFINITION.md
├── PROMPT-GENERATE-480-SECURITY-DEFINITION.md
├── PROMPT-GATE-480-SECURITY-DEFINITION.md
├── PROMPT-FIX-480-SECURITY-DEFINITION.md
├── PROMPT-GENERATE-490-DATA-ARCHITECTURE-DEFINITION.md    🆕
├── PROMPT-GATE-490-DATA-ARCHITECTURE-DEFINITION.md        🆕
├── PROMPT-FIX-490-DATA-ARCHITECTURE-DEFINITION.md         🆕
├── PROMPT-GENERATE-500-DEVOPS-SRE-DEFINITION.md           🆕
├── PROMPT-GATE-500-DEVOPS-SRE-DEFINITION.md               🆕
├── PROMPT-FIX-500-DEVOPS-SRE-DEFINITION.md                🆕
├── PROMPT-GENERATE-510-TEST-STRATEGY-DEFINITION.md         🆕
├── PROMPT-GATE-510-TEST-STRATEGY-DEFINITION.md             🆕
├── PROMPT-FIX-510-TEST-STRATEGY-DEFINITION.md              🆕
├── PROMPT-GENERATE-520-INFRA-CLOUD-DEFINITION.md           🆕
├── PROMPT-GATE-520-INFRA-CLOUD-DEFINITION.md               🆕
├── PROMPT-FIX-520-INFRA-CLOUD-DEFINITION.md                🆕
├── PROMPT-GENERATE-530-SOLUTIONS-CATALOG.md
├── PROMPT-GATE-530-SOLUTIONS-CATALOG.md
├── PROMPT-FIX-530-SOLUTIONS-CATALOG.md
├── PROMPT-GENERATE-540-SOLUTIONS-MATRIX.md
├── PROMPT-GATE-540-SOLUTIONS-MATRIX.md
├── PROMPT-FIX-540-SOLUTIONS-MATRIX.md
├── PROMPT-GENERATE-550-SOLUTIONS-STACK-MATRIX.md
├── PROMPT-GATE-550-SOLUTIONS-STACK-MATRIX.md
├── PROMPT-FIX-550-SOLUTIONS-STACK-MATRIX.md
├── PROMPT-GENERATE-560-SPECS-DEFINITION.md
├── PROMPT-GATE-560-SPECS-DEFINITION.md
├── PROMPT-FIX-560-SPECS-DEFINITION.md
├── PROMPT-GENERATE-570-MILESTONES.md
├── PROMPT-GATE-570-MILESTONES.md
├── PROMPT-FIX-570-MILESTONES.md
├── PROMPT-GENERATE-580-PACKAGE-BACKLOG-REFINED.md           🆕
├── PROMPT-GATE-580-PACKAGE-BACKLOG-REFINED.md               🆕
├── PROMPT-FIX-580-PACKAGE-BACKLOG-REFINED.md                🆕
├── PROMPT-GENERATE-590-TECHNICAL-DISCOVERY.md              🆕
├── PROMPT-GATE-590-TECHNICAL-DISCOVERY.md                  🆕
├── PROMPT-FIX-590-TECHNICAL-DISCOVERY.md                   🆕
├── PROMPT-GENERATE-600-EXECUTION-HISTORY.md
├── PROMPT-GENERATE-595-TECHLEAD-RETURN-PACKAGE.md        🆕 (Bloco E)
├── PROMPT-GATE-595-TECHLEAD-RETURN-PACKAGE.md            🆕
├── PROMPT-FIX-595-TECHLEAD-RETURN-PACKAGE.md             🆕
└── ... (61 prompts no total: 21 GENERATE + 20 GATE + 20 FIX; + pacote de desenvolvimento em sprint-tecnhnical-implementation/ — orquestrador 1000 + 18 fases 1010-1180 — + 3 especialistas reusados EXECUTE-CI-CD-PIPELINE, EXECUTE-CVE-SCA-SCAN, EXECUTE-STRESS-PERFORMANCE-TEST e roadmap companion IMPLEMENTATION-TOOLING com 12 prompts em implementation-tooling/)
```

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: roadmap de 11 fases em 5 blocos para definições técnicas do projeto | Time de Arquitetura |
| 2.0 | 28/07/2026 | Refatoração HITL: gates unificados com 3 perguntas obrigatórias e veredito binário; Fase 4 (PRD) movida para Bloco B; novo prompt GENERATE-TEAM-CAPACITY; Fase 10 reduzida à pasta sprints/; FAIL_REPORT removido (relatório inline) | Time de Arquitetura |
| 3.0 | 28/07/2026 | TEAM-MAP renomeado para TEAM-SKILLS-MAP com Discovery Team (9 papéis); TEAM-CAPACITY promovido a Fase 2 com Gate/Fix; fases renumeradas (1→12); 31 prompts (11+10+10) | Time de Arquitetura |
| 4.0 | 28/07/2026 | Reestruturação TOGAF: Bloco A=People(2f), Bloco B=Architecture & Security(2f), Bloco C=Specs & Milestones(3f, PRD→SPECS→MILESTONES), Bloco D=Execution Portfolio(5f, CATALOG→MATRIX→STACK→SPRINTS→HISTORY); cadeia progressiva de inputs no Bloco D; inputs do PRD enriquecidos com Blocos A+B | Time de Arquitetura |
| 5.0 | 30/07/2026 | Redesign estrutural: Bloco 0 (ponte Negócio→TI com 4 fases), Bloco B expandido (6 disciplinas: +DATA, +DEVOPS-SRE, +TEST, +INFRA), Bloco C reorganizado (CATALOG→MATRIX→STACK→SPECS→MILESTONES), Bloco D=Sprints com Discovery Técnico, History standalone. 20 fases, pipeline sequencial, 58 prompts. | Time de Arquitetura |
| 6.0 | 16/08/2026 | Modos de execução (agile-discovery/waterfall-discovery) com detecção no Bootstrap; modo waterfall-discovery com mapeamento F1–F19 (pular/migrar/validar/rodar); Bloco E — Esteira de Construção por ciclo com reuso sem edição + trio 595-TECHLEAD-RETURN-PACKAGE; parceria PM/PO × TECHLEAD (WATERFALL-EXECUTION v2.0). 61 prompts. | Time de Arquitetura |
| 7.0 | 17/08/2026 | Bloco F — Janelas de Entrega (DEV→QA→UAT→DEPLOY por CICLO-NN): consome o 096-DEFINICAO-JANELAS-ENTREGA (WATERFALL), orquestra a passagem por janela com delegação por frente (sem prompts executores), rastreio no 600 e no pacote 595 (coluna Janela + Aceite UAT DE-ACORDO). Vetos do 092 preservados; 105 permanece aceite final. 61 prompts (sem executores novos). | Time de Arquitetura |
| 7.1 | 17/08/2026 | Bloco E ganha 3 especialistas opcionais reusados de sprint-tecnhnical-implementation/ (steps 3a/3b/4a: EXECUTE-CVE-SCA-SCAN, EXECUTE-CI-CD-PIPELINE, EXECUTE-STRESS-PERFORMANCE-TEST); Bloco F ganha F.6 — tooling de ambiente via roadmap companion IMPLEMENTATION-TOOLING v1.0 (trios 610/620/630/640). Sem prompts novos no folder TECHLEAD (61 preservados). | Time de Arquitetura |
| 7.2 | 17/08/2026 | Renomeação global FILA-NN → CICLO-NN (Ciclo de Entrega) em todo o ecossistema; regra de espelhamento `590-ciclo-NNN ↔ CICLO-NN`; alinhado ao Planejamento do Ciclo (WATERFALL-EXECUTION v2.3, Sub-fase 1.5). Trio 580 renomeado: `580-SPRINT-BACKLOG` → `580-PACKAGE-BACKLOG-REFINED` (prompt e artefato `technical-discovery/580-PACKAGE-BACKLOG-REFINED.md`), alinhado à terminologia de Pacote de Trabalho (060-EAP-WBS). | Time de Arquitetura |
| 7.3 | 17/08/2026 | Revisão de termos ágeis: vocabulário genérico FILA/CICLO/CICLOS/PACKAGE no roadmap e FLOWCHART — formas duplas `ciclo/sprint` onde o gerador de demanda pode vir de fluxos ágeis ou waterfall (Bloco D "Ciclos/Sprints", "Ciclo/Sprint-Alvo", Barreira D); pasta do discovery técnico renomeada `590-sprint-NNN` → `590-ciclo-NNN` (roadmap, FLOWCHART e prompts 580/590/595/600 + WATERFALL-EXECUTION). Nomes estruturais preservados (sprint-artefacts/, SPRINT-CARD, branch feature/sprint-NN). | Time de Arquitetura |
| 7.4 | 17/08/2026 | Bloco E passa a apontar o **pacote de desenvolvimento** — orquestrador `PROMPT-EXECUTE-1000-PACKAGE-DEVELOPMENT-ORCHESTRATOR` (18 fases numeradas na ordem de execução, artefato PACKAGE-DEVELOPMENT-{FASE}.md por fase, gate HITL 1130, débito centralizado 1010); legado monolítico mantido como fallback (step 4); empacotamento 595 inalterado. | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência listados na seção Skills Utilizados. Outros skills podem ser utilizados conforme aderência à necessidade específica.*
