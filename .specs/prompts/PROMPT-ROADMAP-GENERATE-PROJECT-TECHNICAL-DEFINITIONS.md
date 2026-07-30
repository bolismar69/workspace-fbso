# PROMPT: ROADMAP DE DEFINIÇÕES TÉCNICAS DO PROJETO
## Versão: 5.0 — Ponte Negócio→TI + 6 Disciplinas Técnicas + Discovery Contínuo

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

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
TECHNICAL_DEFINITIONS_PATH    = PROJECT_COMPLETE_PATH_NAME + "/technical-definitions"
```

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
  ├─▶ BLOCO D: Sprints — Technical Discovery (REPROPOSITADO)
  │     Fase 18 → Fase 19 (iterativo)
  │     ⛔ Barreira D
  │
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
4. Criar template `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY-EXCEPTIONS.md`
5. Auditar artefatos existentes — varredura completa nos documentos para identificar a situação atual do processo de criação da documentação
6. Apresentar resumo da situação atual e iniciar a primeira fase pendente (ou ponto de retomada)

### Fase 1 — PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md 🆕
Registro de lotes de ingestão de requisitos do Negócio. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md` → Gate → Fix → COMPLIANCE

### Fase 2 — PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT.md 🆕
Aplicação do DoR de Negócio (PO/PM). Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT.md` → Gate → Fix → COMPLIANCE

### Fase 3 — PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md 🆕
Backlog consolidado "Pronto para TI". Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md` → Gate → Fix → COMPLIANCE

### Fase 4 — PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md 🔄
PRD de Negócio (movido do Bloco C). Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 5 — PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md
Skills matrix do Discovery Team. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md` → Gate → Fix → COMPLIANCE

### Fase 6 — PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md
Capacidade de trabalho do time. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` → Gate → Fix → COMPLIANCE

### Fase 7 — PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md
Solution Architect — integração C4, ADRs, topologia. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 8 — PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md
Security Architect — threat model, IAM, compliance. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 9 — PROJECT-TECHNICAL-DEFINITIONS-DATA-ARCHITECTURE-DEFINITION.md 🆕
Data Architect — modelagem, pipelines, storage strategy. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-DATA-ARCHITECTURE-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 10 — PROJECT-TECHNICAL-DEFINITIONS-DEVOPS-SRE-DEFINITION.md 🆕
DevOps/SRE Architect — CI/CD, IaC, observabilidade, SLOs. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-DEVOPS-SRE-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 11 — PROJECT-TECHNICAL-DEFINITIONS-TEST-STRATEGY-DEFINITION.md 🆕
Test Specialist — pirâmide de testes, automação, performance. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-TEST-STRATEGY-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 12 — PROJECT-TECHNICAL-DEFINITIONS-INFRA-CLOUD-DEFINITION.md 🆕
Infra/Cloud Specialist — topologia, compute, networking, DR. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-INFRA-CLOUD-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 13 — PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md
Catálogo de soluções técnicas. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md` → Gate → Fix → COMPLIANCE

### Fase 14 — PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md
Matriz solução×disciplina×owner. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md` → Gate → Fix → COMPLIANCE

### Fase 15 — PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md
Stacks tecnológicas por solução. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md` → Gate → Fix → COMPLIANCE

### Fase 16 — PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md
Consolidação técnica enxuta — sumariza e referencia artefatos anteriores. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 17 — PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md
Roadmap alinhado ao negócio. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md` → Gate → Fix → COMPLIANCE

### Fase 18 — PROJECT-TECHNICAL-DEFINITIONS-SPRINT-BACKLOG-REFINED.md 🆕
Backlog refinado com tarefas T-NNN → US-ID → Sprint-Alvo → CONTRACTS. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SPRINT-BACKLOG-REFINED.md` → Gate → Fix → COMPLIANCE

### Fase 19 — PROJECT-TECHNICAL-DEFINITIONS-TECHNICAL-DISCOVERY.md 🆕
Discovery Técnico Contínuo — contratos API/Data/Security/SRE por sprint + increments. Iterativo. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-TECHNICAL-DISCOVERY.md` → Gate → Fix → COMPLIANCE

### Execution History — PROJECT-TECHNICAL-DEFINITIONS-EXECUTION-HISTORY.md 📊
Dashboard de controle — estado de todos os documentos. Pipeline: Generate → Revisão humana (sem gate próprio).

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
| ⛔ Barreira D | Após Bloco D (F19) | 100% US da sprint com contratos. SPRINT-BACKLOG atualizado. | Iterativo — pergunta se continua próxima sprint ou encerra |

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
    ├── PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md              (F1)  🆕
    ├── PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT.md          (F2)  🆕
    ├── PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md    (F3)  🆕
    ├── PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md          (F4)  🔄
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
    ├── technical-discovery/
    │   ├── SPRINT-BACKLOG.md                                    (F18)
    │   ├── sprint-00/
    │   │   ├── CONTRACTS-API-sprint-00.md
    │   │   ├── CONTRACTS-DATA-sprint-00.md
    │   │   ├── CONTRACTS-SECURITY-sprint-00.md
    │   │   ├── CONTRACTS-SRE-sprint-00.md
    │   │   └── DEFINITION-INCREMENTS-sprint-00.md
    │   └── sprint-01/ ...
    └── PROJECT-TECHNICAL-DEFINITIONS-EXECUTION-HISTORY.md
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

> **🔄 Flexibilidade:** Se durante a execução o agente identificar que um skill diferente é mais adequado, substituí-lo e justificar no PROJECT-TECHNICAL-DEFINITIONS-EXECUTION-HISTORY.md.

---

## Localização dos Prompts das Fases

Os prompts de geração, gate e correção de cada fase estão na pasta `project-technical-definitions/`:

```
.specs/prompts/project-technical-definitions/
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md            🆕
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md                🆕
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md                 🆕
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT.md         🆕
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT.md             🆕
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT.md              🆕
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md   🆕
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md       🆕
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md        🆕
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-DATA-ARCHITECTURE-DEFINITION.md    🆕
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-DATA-ARCHITECTURE-DEFINITION.md        🆕
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-DATA-ARCHITECTURE-DEFINITION.md         🆕
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-DEVOPS-SRE-DEFINITION.md           🆕
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-DEVOPS-SRE-DEFINITION.md               🆕
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-DEVOPS-SRE-DEFINITION.md                🆕
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-TEST-STRATEGY-DEFINITION.md         🆕
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-TEST-STRATEGY-DEFINITION.md             🆕
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-TEST-STRATEGY-DEFINITION.md              🆕
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-INFRA-CLOUD-DEFINITION.md           🆕
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-INFRA-CLOUD-DEFINITION.md               🆕
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-INFRA-CLOUD-DEFINITION.md                🆕
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SPRINT-BACKLOG-REFINED.md           🆕
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SPRINT-BACKLOG-REFINED.md               🆕
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SPRINT-BACKLOG-REFINED.md                🆕
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-TECHNICAL-DISCOVERY.md              🆕
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-TECHNICAL-DISCOVERY.md                  🆕
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-TECHNICAL-DISCOVERY.md                   🆕
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-EXECUTION-HISTORY.md
└── ... (58 prompts no total: 20 GENERATE + 19 GATE + 19 FIX)
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

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência listados na seção Skills Utilizados. Outros skills podem ser utilizados conforme aderência à necessidade específica.*
