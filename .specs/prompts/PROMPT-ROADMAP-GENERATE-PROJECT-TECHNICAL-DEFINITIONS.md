# PROMPT: ROADMAP DE DEFINIÇÕES TÉCNICAS DO PROJETO
## Versão: 1.0 — Bootstrap Inteligente + Validação Soberana Humana (Human-in-the-Loop)

Atue como um Especialista em Gestão de Processos (BPM), Arquiteto de Soluções Ágeis e Tech Lead, especializado em definições técnicas de projetos e engenharia de prompts.

Preciso que você crie um roadmap de execução detalhado e um guia de orquestração para o seguinte processo: Criação, revisão, evolução e validação dos **documentos de definição técnica do projeto** — artefatos que preenchem o gap entre os documentos de negócio (`PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS.md`) e as especificações técnicas por solução (`PROMPT-ROADMAP-GENERATE-TECHNICAL_SOLUTIONS.md`, na pasta `technical-solutions/`).

Objetivo Principal: Garantir que todas as definições técnicas do projeto estejam criadas, revisadas e 100% alinhadas conceitualmente entre si e com os documentos de negócio (Charter, BRD, Epics, Features, User Stories), preparando o terreno para que cada time de solução técnica inicie seu trabalho com baseline consistente.

Regra Crítica de Execução (Gating Rule): O processo é estritamente sequencial dentro de cada bloco. Nenhuma fase subsequente pode ser iniciada sem a aprovação formal, soberana e explícita do usuário humano na fase anterior. Blocos independentes podem ser executados em paralelo.

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

O roadmap é organizado em **11 fases** agrupadas em **5 blocos**:

```
FASE 0: BOOTSTRAP (sequencial)
  │
  ├─▶ BLOCO A: People & Solutions
  │     Fase 1 → Fase 2 → Fase 3
  │
  ├─▶ (barreira — Bloco A concluído)
  │     └─▶ Fase 4: PRD Definition 🆕
  │
  ├─▶ BLOCO B: Architecture & Security
  │     Fase 5 → Fase 6
  │
  └─▶ (barreira de sincronização)
        │
        ├─▶ BLOCO C: Specs & Milestones
        │     Fase 7 → Fase 8
        │
        └─▶ BLOCO D: Matriz, Sprints, Histórico
              Fase 9 → Fase 10 → Fase 11
```

---

## FASES DO ROADMAP

### Fase 0 — Bootstrap Inteligente

Workflow:
1. Solicitar inputs ao usuário (se não fornecidos)
2. Exibir caminhos derivados e solicitar confirmação
3. Criar estrutura: `mkdir -p {TECHNICAL_DEFINITIONS_PATH}`
4. Migrar TECHNICAL-TEAM-MAP.md (se existir) para `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md`
5. Criar template `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY-EXCEPTIONS.md`
6. Auditar artefatos existentes e determinar ponto de partida
7. Apresentar resumo e iniciar primeira fase pendente

### Fase 1 — PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md
Skills matrix do time. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md` → Gate → Fix → COMPLIANCE

### Fase 2 — PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md
Catálogo de soluções técnicas. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md` → Gate → Fix → COMPLIANCE

### Fase 3 — PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md
Stacks por solução. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md` → Gate → Fix → COMPLIANCE

### Fase 4 — PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md 🆕
Baseline de PRD do projeto. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 5 — PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md
Como as soluções se integram. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 6 — PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md
Regras de segurança do projeto. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 7 — PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md
Baseline de especificações. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 8 — PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md
Roadmap alinhado ao negócio. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md` → Gate → Fix → COMPLIANCE

### Fase 9 — PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md
Matriz solução×stack×owner. Pipeline: `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md` → Gate → Fix → COMPLIANCE

### Fase 10 — Criação da Estrutura de Sprints
Ação Bash para criar pastas de sprints em cada solução. Sem pipeline Generate→Gate→Fix — é uma ação direta com confirmação humana.

### Fase 11 — PROJECT-TECHNICAL-DEFINITIONS-EXECUTION-HISTORY.md
Dashboard de controle — estado de todos os documentos. Pipeline: Generate → Revisão humana (sem gate próprio — atualizado após cada fase).

---

## MECANISMO DE ORQUESTRAÇÃO DINÂMICA

Toda fase (1-9) deve rodar sob o ecossistema trifásico de prompts (Gerador, Auditor/Portão e Corretor), com controle final obrigatório do Humano:

1. **Geração / Evolução:** A IA executa o prompt gerador da fase
2. **Auditoria Interna:** O artefato é enviado para o gate. Se encontrar erros → `[NÃO COMPLIANCE]` → aciona o FIX → retorna ao gate. Se não encontrar erros → avança para validação humana
3. **Portão de Validação Humana:** Status `[PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]` com 3 perguntas obrigatórias
4. **Lógica de Decisão:** Aprovação = COMPLIANCE e próxima fase. Novos inputs = retrocesso ao gerador (evolução incremental)

### Esquema de Estados por Documento

```
CREATED → GATE ⟷ FIX → PRE-COMPLIANCE → COMPLIANCE
  │         │              │
  │         └─ NÃO-COMPLIANCE (com FAIL_REPORT)
  │
  └─ (estado inicial após geração)
```

---

## ESTRUTURA DE DIRETÓRIOS GERADA

```
business-inputs/business-projects/{PROJECT_ID_NAME}/
└── technical-definitions/
    ├── PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md
    ├── PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY-EXCEPTIONS.md
    ├── PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md              (Fase 1)
    ├── PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md     (Fase 2)
    ├── PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md (Fase 3)
    ├── PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md        (Fase 4)
    ├── PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md (Fase 5)
    ├── PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md   (Fase 6)
    ├── PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md      (Fase 7)
    ├── PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md            (Fase 8)
    ├── PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md      (Fase 9)
    ├── PROJECT-TECHNICAL-DEFINITIONS-EXECUTION-HISTORY.md     (Fase 11)
    └── sprints/
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
| 5 | `workflow-orchestration-patterns` | Padrões de orquestração de workflows multi-fase | Orquestração |
| 6 | `dispatching-parallel-agents` | Disparo de agentes em paralelo para Blocos A e B | Orquestração |
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
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md
├── PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md
├── PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md
├── PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md
├── ... (32 prompts no total)
```

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: roadmap de 11 fases em 5 blocos para definições técnicas do projeto | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência listados na seção Skills Utilizados. Outros skills podem ser utilizados conforme aderência à necessidade específica.*
