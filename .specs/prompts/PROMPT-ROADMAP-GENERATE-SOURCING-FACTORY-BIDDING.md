# PROMPT: ROADMAP DE SOURCING & FACTORY BIDDING
## Versão: 1.0 — RFQ Package + DTA Estimation Schema + Factory Selection

Atue como um Especialista em Gestão de Processos (BPM), PMO e Tech Lead, especializado em Sourcing de Fábricas de Software, Bidding e Estimativas.

Preciso que você crie um roadmap de execução detalhado e um guia de orquestração para o seguinte processo: **Empacotamento de RFQ, envio para fábricas de software, recebimento de estimativas padronizadas, validação automatizada, comparação e seleção da fábrica vencedora.**

Objetivo Principal: Produzir uma **Matriz Comparativa de Fábricas + Recomendação de Seleção** com qualidade suficiente para o PM/Comitê decidir qual fábrica contratar.

Regra Crítica de Execução (Gating Rule): O processo é estritamente sequencial em todas as fases. Nenhuma fase subsequente pode ser iniciada sem a aprovação formal, soberana e explícita do usuário humano na fase anterior.

**Base técnica:** DTA Engine de Bidding e Estimativas (`/home/bolismar/work/workspace-fbso/.specs/DTA-Engine-de-Bidding-e-Estimativas.md`).

---

## VARIÁVEIS DE ENTRADA E BOOTSTRAP (FASE 0)

### Tabela de Inputs

| Variável | Obrig. | Descrição | Exemplo |
|---|---|---|---|
| `PROJECT_PATH` | ✅ | Caminho base onde os projetos de negócio residem | `/home/bolismar/work/workspace-fbso/business-inputs/business-projects` |
| `PROJECT_ID_NAME` | ✅ | Identificador completo do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `SOURCING_BIDDING_MODE` | ❌ | Modo de estimativa: `discovery` ou `full`. Se não informado, o Bootstrap detecta e pergunta. | `discovery` |
| `PROJECT_DOCUMENTS_INPUTS` | ❌ | Lista de caminhos para documentos brutos de entrada adicionais | `[]` |
| `PROJECT_PROMPT_INPUTS` | ❌ | Lista de caminhos para prompts auxiliares ou contextos adicionais | `[]` |

### Variáveis Derivadas

```
PROJECT_COMPLETE_PATH_NAME  = PROJECT_PATH + "/" + PROJECT_ID_NAME
SOURCING_BIDDING_PATH       = PROJECT_COMPLETE_PATH_NAME + "/sourcing-factory-bidding-" + SOURCING_BIDDING_MODE
ESTIMATES_PATH              = SOURCING_BIDDING_PATH + "/estimates"
```

### Modos de Operação

| Modo | Pasta | Nível | Artefatos de entrada |
|---|---|---|---|
| **Mode 1 — `discovery`** | `sourcing-factory-bidding-discovery/` | Alto nível (épicos) | Upstream Architecture Discovery (PRD + Bloco B + SPECS + ROM) |
| **Mode 2 — `full`** | `sourcing-factory-bidding-full/` | Detalhado (features + US) | Docs de negócio completos (Features + User Stories + RTM) |

---

## ARQUITETURA DE FASES

O roadmap é organizado em **6 fases** agrupadas em **3 blocos**:

```
FASE 0: BOOTSTRAP (detecta modo + pergunta ao humano)
  │
  ├─▶ BLOCO A: RFQ Package
  │     Fase 1 → Fase 2
  │     ⛔ Barreira A
  │
  ├─▶ BLOCO B: Distribution & Receipt
  │     Fase 3 → Fase 4
  │     ⛔ Barreira B
  │
  └─▶ BLOCO C: Validation & Comparison
        Fase 5 → Fase 6
        ⛔ Barreira C → Factory Selection
```

---

## FASES DO ROADMAP

### Fase 0 — Bootstrap Inteligente

Workflow:
1. Solicitar inputs ao usuário (se não fornecidos)
2. Auditar artefatos existentes:
   - Verificar existência de `upstream-architecture-discovery/`
   - Verificar existência de `features/` e `user-stories/`
3. Apresentar opções de modo:
   - Se `upstream-architecture-discovery/` existe mas `features/` não → **Mode 1 (discovery)**
   - Se `features/` e `user-stories/` existem → **Mode 1 (discovery)** e **Mode 2 (full)**
4. **Humano decide** qual modo (`SOURCING_BIDDING_MODE`)
5. Criar estrutura: `mkdir -p {ESTIMATES_PATH}`
6. Auditar estimativas já recebidas na pasta `estimates/`
7. Apresentar resumo e iniciar a primeira fase pendente

### Fase 1 — RFQ-PACKAGE.md 🆕
Pacote RFQ compilando artefatos técnicos conforme o modo. Pipeline: `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-RFQ-PACKAGE.md` → Gate → Fix → COMPLIANCE

### Fase 2 — ESTIMATION-SCHEMA.csv 🆕
Template CSV padronizado (DTA Estimation Schema) que as fábricas devem preencher. Colunas: id_epico, titulo, solucoes, horas_dev, horas_arch, horas_qa, **prazo_entrega_meses**, complexidade, comentarios. Pipeline: `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATION-SCHEMA.md` → Gate → Fix → COMPLIANCE

### Fase 3 — FACTORY-DISTRIBUTION.md 🆕
Registro de fábricas participantes, envio do RFQ e prazos. Pipeline: `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-FACTORY-DISTRIBUTION.md` → Gate → Fix → COMPLIANCE

### Fase 4 — ESTIMATE-RECEIPT.md 🆕
Guia para o time operacional salvar estimativas recebidas. Padrão: `nome-do-arquivo-csv-{nome-da-fabrica}.md` em `estimates/`. Pipeline: `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATE-RECEIPT.md` → Gate → Fix → COMPLIANCE

### Fase 5 — ESTIMATE-VALIDATION.md 🆕
Validação DTA de cada estimativa recebida (formato, QA balanceado, outliers). Pipeline: `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATE-VALIDATION.md` → Gate → Fix → COMPLIANCE

### Fase 6 — FACTORY-COMPARISON.md 🆕
Matriz comparativa cross-fábrica + ranking + recomendação de seleção. Pipeline: `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-FACTORY-COMPARISON.md` → Gate → Fix → COMPLIANCE

---

## MECANISMO DE ORQUESTRAÇÃO DINÂMICA

Toda fase (1-6) segue o ecossistema trifásico Generate→Gate→Fix com validação humana:

1. **Geração:** IA executa o prompt gerador
2. **Auditoria Interna:** Gate valida. Se erros → FIX → Gate. Se OK → Validação Humana
3. **Validação Humana (3 perguntas):** (1) Compliance? (2) Novos docs? (3) Novos inputs?
4. **Decisão:** Sim/Não/Não → COMPLIANCE. Novos inputs → retrocede ao GENERATE

---

## REGRAS DE BLOQUEIO (GATING RULES)

| Barreira | Posição | Validação |
|---|---|---|
| ⛔ Barreira A | Após Bloco A (F2) | RFQ completo. Schema CSV gerado conforme modo. |
| ⛔ Barreira B | Após Bloco B (F4) | Fábricas registradas. Estimativas recebidas salvas em `estimates/`. |
| ⛔ Barreira C | Após Bloco C (F6) | Todas estimativas validadas. Matriz comparativa preenchida. Recomendação justificada. |

---

## ESTRUTURA DE DIRETÓRIOS GERADA

### Mode 1 — Discovery
```
business-inputs/business-projects/{PROJECT_ID_NAME}/
└── sourcing-factory-bidding-discovery/
    ├── RFQ-PACKAGE.md                    (F1)
    ├── ESTIMATION-SCHEMA.csv             (F2)
    ├── FACTORY-DISTRIBUTION.md           (F3)
    ├── ESTIMATE-RECEIPT.md               (F4)
    ├── ESTIMATE-VALIDATION-FABRICA-*.md  (F5 — por fábrica)
    ├── FACTORY-COMPARISON.md             (F6)
    └── estimates/
        ├── nome-do-arquivo-csv-fabrica-A.md
        ├── nome-do-arquivo-csv-fabrica-B.md
        └── nome-do-arquivo-csv-fabrica-C.md
```

### Mode 2 — Full
```
business-inputs/business-projects/{PROJECT_ID_NAME}/
└── sourcing-factory-bidding-full/
    ├── ... (mesma estrutura acima)
    └── estimates/
        └── ...
```

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `superpowers:brainstorming` | Brainstorming da estratégia de sourcing | Orquestração |
| 2 | `superpowers:executing-plans` | Execução do plano de fases com gates | Orquestração |
| 3 | `project-estimation` | Estrutura geral de estimativas e metodologia | Estimativa |
| 4 | `estimate-builder` | Construção de templates de estimativa padronizados | Estimativa |
| 5 | `estimate-builder-qmohd` | Validação de estimativas recebidas | Estimativa |
| 6 | `analyst-estimates` | Análise comparativa de estimativas de fábricas | Análise |
| 7 | `ads-budget` | Comparação orçamentária e alocação entre fábricas | Financeiro |
| 8 | `trade-show-budget-planner` | Modelo de investimento ROI e go/no-go por fábrica | Financeiro |
| 9 | `afrexai-construction-estimator` | Metodologia de estimativa detalhada por linha | Estimativa |
| 10 | `senior-architect` | Validação técnica do pacote RFQ | Arquitetura |
| 11 | `gap-analysis` | Análise de outliers e discrepâncias entre fábricas | Análise |
| 12 | `documentation-writer` | Documentação do roadmap e relatórios | Documentação |

---

## Localização dos Prompts

```
.specs/prompts/sourcing-factory-bidding/
├── PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-RFQ-PACKAGE.md
├── PROMPT-GATE-SOURCING-FACTORY-BIDDING-RFQ-PACKAGE.md
├── PROMPT-FIX-SOURCING-FACTORY-BIDDING-RFQ-PACKAGE.md
├── PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATION-SCHEMA.md
├── PROMPT-GATE-SOURCING-FACTORY-BIDDING-ESTIMATION-SCHEMA.md
├── PROMPT-FIX-SOURCING-FACTORY-BIDDING-ESTIMATION-SCHEMA.md
├── PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-FACTORY-DISTRIBUTION.md
├── PROMPT-GATE-SOURCING-FACTORY-BIDDING-FACTORY-DISTRIBUTION.md
├── PROMPT-FIX-SOURCING-FACTORY-BIDDING-FACTORY-DISTRIBUTION.md
├── PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATE-RECEIPT.md
├── PROMPT-GATE-SOURCING-FACTORY-BIDDING-ESTIMATE-RECEIPT.md
├── PROMPT-FIX-SOURCING-FACTORY-BIDDING-ESTIMATE-RECEIPT.md
├── PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATE-VALIDATION.md
├── PROMPT-GATE-SOURCING-FACTORY-BIDDING-ESTIMATE-VALIDATION.md
├── PROMPT-FIX-SOURCING-FACTORY-BIDDING-ESTIMATE-VALIDATION.md
├── PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-FACTORY-COMPARISON.md
├── PROMPT-GATE-SOURCING-FACTORY-BIDDING-FACTORY-COMPARISON.md
├── PROMPT-FIX-SOURCING-FACTORY-BIDDING-FACTORY-COMPARISON.md
└── (18 prompts: 6 GENERATE + 6 GATE + 6 FIX)
```

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial: roadmap de Sourcing & Factory Bidding com 6 fases em 3 blocos, 2 modos (discovery/full), DTA Engine integration. | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada. Skills de referência listados na seção Skills Utilizados.*
