# PROMPT: ROADMAP DE UPSTREAM ARCHITECTURE DISCOVERY
## Versão: 1.0 — Discovery Upstream Engineering & Architecture + ROM 50%

Atue como um Especialista em Gestão de Processos (BPM), Arquiteto de Soluções Ágeis e Tech Lead, especializado em Discovery Técnico, Upstream Architecture e Análise de Viabilidade.

Preciso que você crie um roadmap de execução detalhado e um guia de orquestração para o seguinte processo: **Análise de viabilidade técnica e estimativa ROM 50%** na fase de Discovery/Upstream Engineering — quando o Negócio já definiu Project Charter, BRD e Épicos, e precisa que o time de TI avalie a viabilidade técnica, desenhe uma solução macro e produza uma estimativa de alto nível (ROM +-50%) para o comitê de governança decidir GO/NO-GO.

Objetivo Principal: Produzir um **Desenho de Solução Macro + Estimativa ROM 50%** com qualidade suficiente para o comitê de governança aprovar (ou não) o financiamento do projeto e o estabelecimento do SQUAD.

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
UPSTREAM_DISCOVERY_PATH       = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
```

### Pré-condição Obrigatória

O roadmap de documentos de negócio (`PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS.md`) deve ter sido executado **até a fase EPICS** (Fase 3). Os seguintes artefatos devem existir:

- `01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md`
- `02-BRD-{PROJECT_ID_NAME}.md`
- `03-EPICS-{PROJECT_ID_NAME}.md`
- `epics/*.md` (arquivos individuais de épicos)

---

## ARQUITETURA DE FASES

O roadmap é organizado em **11 fases** agrupadas em **5 blocos**:

```
FASE 0: BOOTSTRAP (sequencial)
  │
  ├─▶ BLOCO 0: Product Definition Discovery-Level
  │     Fase 1
  │     ⛔ Barreira 0
  │
  ├─▶ BLOCO B: Architecture & Security & Specialists (Discovery-Level)
  │     Fase 2 → Fase 3 → Fase 4 → Fase 5 → Fase 6 → Fase 7
  │     ⛔ Barreira B
  │
  ├─▶ BLOCO C: Catálogo, Matriz & Consolidação Discovery-Level
  │     Fase 8 → Fase 9 → Fase 10
  │     ⛔ Barreira C
  │
  ├─▶ BLOCO D: Estimativa & ROM
  │     Fase 11
  │     ⛔ Barreira D
  │
  └─▶ GATE DE GOVERNANÇA: GO / NO-GO
        ├── No-Go ❌ → Projeto Cancelado / Arquivado
        └── Go-Ahead ✅ → Dispara roadmaps downstream
```

---

## FASES DO ROADMAP

### Fase 0 — Bootstrap Inteligente

Workflow:
1. Solicitar inputs ao usuário (se não fornecidos)
2. Validar pré-condição: verificar existência de Charter, BRD, Épicos
3. Exibir caminhos derivados e solicitar confirmação
4. Criar estrutura: `mkdir -p {UPSTREAM_DISCOVERY_PATH}`
5. Auditar artefatos existentes no diretório `upstream-architecture-discovery/`
6. Apresentar resumo e iniciar a primeira fase pendente

### Fase 1 — DISCOVERY-LEVEL-PRD.md 🆕
PRD Discovery-Level — visão do produto baseada nos Épicos. Pipeline: `PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-PRD.md` → Gate → Fix → COMPLIANCE

### Fase 2 — DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md 🆕
Solution Architect — arquitetura macro, C4 Level 1, estratégia de integração high-level. Pipeline: `PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 3 — DISCOVERY-LEVEL-SECURITY-DEFINITION.md 🆕
Security Architect — threat model high-level, requisitos de compliance, estratégia de segurança macro. Pipeline: `PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SECURITY-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 4 — DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md 🆕
Data Architect — estratégia de dados macro, volumes estimados, tipo de armazenamento. Pipeline: `PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 5 — DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md 🆕
DevOps/SRE Architect — estratégia de deploy macro, observabilidade, ambientes. Pipeline: `PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 6 — DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md 🆕
Test Specialist — estratégia de testes macro, ambientes de teste, quality gates. Pipeline: `PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 7 — DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md 🆕
Infra/Cloud Specialist — topologia macro, provedor, estimativa de recursos. Pipeline: `PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 8 — DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md 🆕
Catálogo macro de soluções — nomes, tipos, propósito high-level. Pipeline: `PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SOLUTIONS-CATALOG.md` → Gate → Fix → COMPLIANCE

### Fase 9 — DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md 🆕
Matriz macro solução×disciplina×complexidade. Pipeline: `PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SOLUTIONS-MATRIX.md` → Gate → Fix → COMPLIANCE

### Fase 10 — DISCOVERY-LEVEL-SPECS.md 🆕
Consolidação técnica high-level — sumariza descobertas do Bloco B para embasar ROM. Pipeline: `PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SPECS.md` → Gate → Fix → COMPLIANCE

### Fase 11 — ROM-ESTIMATE.md 🆕
Consolidação da estimativa ROM +-50%. Matriz de esforço, premissas, riscos, faixa de valores. Pipeline: `PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-ROM-ESTIMATE.md` → Gate → Fix → COMPLIANCE

### Gate de Governança — GO / NO-GO 🚦

Após a Barreira D, o orquestrador apresenta o **Resumo Executivo para o Comitê de Governança**:

1. **Visão do Projeto:** Resumo do PRD Discovery-Level
2. **Desenho da Solução Macro:** Sumário do SPECS Discovery-Level
3. **Estimativa ROM 50%:** Faixa de valores, premissas, riscos
4. **Recomendação Técnica:** Parecer do time de arquitetura

O Comitê decide:
- **Go-Ahead ✅:** Verba aprovada, SQUAD estabelecida. Dispara:
  - `PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS.md` — continua de Features → User Stories
  - `PROMPT-ROADMAP-GENERATE-PROJECT-TECHNICAL-DEFINITIONS.md` — definições técnicas detalhadas
- **No-Go ❌:** Projeto cancelado ou arquivado para reavaliação futura

---

## MECANISMO DE ORQUESTRAÇÃO DINÂMICA

Toda fase (1-11) deve rodar sob o ecossistema trifásico de prompts (Gerador, Auditor/Portão e Corretor), com controle final obrigatório do Humano:

1. **Geração / Evolução:** A IA recebe os inputs disponíveis e executa o prompt gerador da fase
2. **Auditoria Interna da IA:** O artefato é enviado para o gate.
   - SE A IA ENCONTRAR ERROS: Emite o status `[NÃO COMPLIANCE]`, apresenta as falhas encontradas com sugestões de tratativa, faz as 3 perguntas obrigatórias do Portão de Validação Humana e aguarda o direcionamento do humano. Com base nas respostas, aciona o FIX de forma cirúrgica e retorna ao passo 2.
   - SE A IA NÃO ENCONTRAR ERROS (100% OK): Emite o status `[PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]`, apresenta as 3 perguntas obrigatórias e aguarda o direcionamento humano.
3. **Portão de Validação Humana (3 perguntas obrigatórias):**
   1. O documento está em compliance com a sua necessidade e perfeitamente alinhado com os documentos base?
   2. Deseja enviar mais documentos/arquivos para enriquecer este artefato?
   3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?
4. **Lógica de Decisão Baseada nas Respostas do Humano:**
   - CENÁRIO DE SUCESSO (Aprovação): Se o humano validar e NÃO enviar novos arquivos ou inputs (Sim, Não, Não), a fase é dada por encerrada (`[STATUS: COMPLIANCE]`), o arquivo é congelado e a próxima fase é destravada.
   - CENÁRIO DE RETROCESSO (Evolução Incremental): Se o humano fornecer novos documentos ou novas informações, o orquestrador DEVE retroceder ao passo 1 (GENERATE), injetando o documento gerado até o momento + os novos insumos para uma atualização incremental.

---

## REGRAS DE BLOQUEIO (GATING RULES)

### Barreiras de Bloco

| Barreira | Posição | Validação | Regra Especial |
|---|---|---|---|
| ⛔ Barreira 0 | Após Bloco 0 (F1) | PRD Discovery-Level cobre todos os Épicos. MVP Macro definido. | — |
| ⛔ Barreira B | Após Bloco B (F7) | 6 disciplinas OK. N/A justificados. Consistência horizontal entre os 6 artefatos Discovery-Level. | Disciplina N/A sem justificativa = NÃO COMPLIANCE |
| ⛔ Barreira C | Após Bloco C (F10) | SPECS referencia todos artefatos do Bloco B. Catálogo e Matriz consistentes. | — |
| ⛔ Barreira D | Após Bloco D (F11) | ROM presente. Premissas documentadas. Faixa de valores justificada. | ROM sem premissas = NÃO COMPLIANCE |

### Consistência Horizontal (Bloco B)

A Barreira B deve validar que os 6 artefatos Discovery-Level são consistentes entre si:
- ARCHITECTURE ↔ DATA: modelo de dados high-level alinhado com topologia proposta
- ARCHITECTURE ↔ DEVOPS-SRE: pipeline de deploy compatível com arquitetura
- ARCHITECTURE ↔ INFRA-CLOUD: topologia de infra suporta a arquitetura
- SECURITY ↔ ARCHITECTURE: controles de segurança alinhados com padrão arquitetural
- TEST-STRATEGY ↔ ARCHITECTURE: estratégia de testes cobre os componentes macro
- SECURITY ↔ INFRA-CLOUD: controles de rede e IAM consistentes

---

## ESTRUTURA DE DIRETÓRIOS GERADA

```
business-inputs/business-projects/{PROJECT_ID_NAME}/
└── upstream-architecture-discovery/
    ├── DISCOVERY-LEVEL-PRD.md                       (F1)  🆕
    ├── DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md    (F2)  🆕
    ├── DISCOVERY-LEVEL-SECURITY-DEFINITION.md        (F3)  🆕
    ├── DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md (F4) 🆕
    ├── DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md      (F5)  🆕
    ├── DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md    (F6)  🆕
    ├── DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md      (F7)  🆕
    ├── DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md           (F8)  🆕
    ├── DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md            (F9)  🆕
    ├── DISCOVERY-LEVEL-SPECS.md                       (F10) 🆕
    └── ROM-ESTIMATE.md                               (F11) 🆕
```

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `superpowers:brainstorming` | Brainstorming inicial da solução macro | Orquestração |
| 2 | `superpowers:executing-plans` | Execução do plano de fases com gates | Orquestração |
| 3 | `superpowers:writing-plans` | Escrita e refino do plano de execução | Orquestração |
| 4 | `senior-architect` | Desenho de solução high-level | Arquitetura |
| 5 | `cloud-architect` | Topologia de infra macro | Arquitetura |
| 6 | `senior-devops` | Estratégia de deploy e CI/CD macro | DevOps |
| 7 | `senior-security` | Threat model e compliance high-level | Segurança |
| 8 | `senior-data-engineer` | Estratégia de dados macro | Dados |
| 9 | `senior-qa` | Estratégia de testes macro | Qualidade |
| 10 | `gap-analysis` | Análise de gaps e riscos | Análise |
| 11 | `documentation-writer` | Documentação do roadmap | Documentação |

---

## Localização dos Prompts das Fases

Os prompts de geração, gate e correção de cada fase estão na pasta `upstream-architecture-discovery/`:

```
.specs/prompts/upstream-architecture-discovery/
├── PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-PRD.md                     🆕 F1
├── PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-PRD.md                         🆕 F1
├── PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-PRD.md                          🆕 F1
├── PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md  🆕 F2
├── PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md      🆕 F2
├── PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md       🆕 F2
├── PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SECURITY-DEFINITION.md      🆕 F3
├── PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SECURITY-DEFINITION.md          🆕 F3
├── PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-SECURITY-DEFINITION.md           🆕 F3
├── PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION.md 🆕 F4
├── PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION.md     🆕 F4
├── PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION.md      🆕 F4
├── PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION.md        🆕 F5
├── PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION.md            🆕 F5
├── PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION.md             🆕 F5
├── PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION.md     🆕 F6
├── PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION.md         🆕 F6
├── PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION.md          🆕 F6
├── PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION.md       🆕 F7
├── PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION.md           🆕 F7
├── PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION.md            🆕 F7
├── PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SOLUTIONS-CATALOG.md            🆕 F8
├── PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SOLUTIONS-CATALOG.md                🆕 F8
├── PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-SOLUTIONS-CATALOG.md                 🆕 F8
├── PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SOLUTIONS-MATRIX.md             🆕 F9
├── PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SOLUTIONS-MATRIX.md                 🆕 F9
├── PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-SOLUTIONS-MATRIX.md                  🆕 F9
├── PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SPECS.md                        🆕 F10
├── PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SPECS.md                            🆕 F10
├── PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-SPECS.md                             🆕 F10
├── PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-ROM-ESTIMATE.md                 🆕 F11
├── PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-ROM-ESTIMATE.md                     🆕 F11
├── PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-ROM-ESTIMATE.md                      🆕 F11
└── ... (33 prompts no total: 11 GENERATE + 11 GATE + 11 FIX)
```

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: roadmap de Upstream Architecture Discovery com 11 fases em 5 blocos + Gate de Governança GO/NO-GO. Baseado no DTA Framework V2. | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência listados na seção Skills Utilizados.*
