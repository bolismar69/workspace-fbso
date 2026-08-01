# PROMPT: ROADMAP DE SOURCING & FACTORY BIDDING
## Versão: 1.3 — Schema Unificado + Time/Valor Obrigatórios + Fórmulas DTA

Atue como um Especialista em Gestão de Processos (BPM), PMO e Tech Lead, especializado em Sourcing de Fábricas de Software, Bidding e Estimativas.

Preciso que você crie um roadmap de execução detalhado e um guia de orquestração para o seguinte processo: **Empacotamento de RFQ, envio para fábricas de software, recebimento de estimativas padronizadas, validação automatizada, comparação e seleção da fábrica vencedora.**

Objetivo Principal: Produzir uma **Matriz Comparativa de Fábricas + Recomendação de Seleção** com qualidade suficiente para o PM/Comitê decidir qual fábrica contratar.

Regra Crítica de Execução (Gating Rule): O processo é estritamente sequencial em todas as fases. Nenhuma fase subsequente pode ser iniciada sem a aprovação formal, soberana e explícita do usuário humano na fase anterior.

**Base técnica:** 
- `.specs/standards/DTA-Engine-de-Bidding-e-Estimativas.md` — Schema original e regras de ouro
- `.specs/standards/DTA-VALIDATION-STANDARDS.md` — **Documento canônico:** TODAS as regras, fórmulas e padrões de validação e comparação. Inclui regra PIB (§2.6), matriz de decisão com 5 critérios (§3.2), e Internal Baseline por modo (§5). Consulta obrigatória antes de executar qualquer fase.

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

O roadmap é organizado em **8 fases** agrupadas em **3 blocos**:

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
  └─▶ BLOCO C: Validation, Comparison & Notification
        Fase 5 → [Fase 5b] → Fase 6 → Fase 7
        ⛔ Barreira C → Factory Selection
        
        [Fase 5b] = condicional: executa apenas se 0 fábricas aprovadas na F5
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
Template CSV padronizado (DTA Estimation Schema) que as fábricas devem preencher. Schema unificado para ambos os modos com colunas individuais de horas, prazo, time, e valor.

**Colunas (Discovery e Full — schema unificado):**
`fabrica; id_epico; titulo; features_codigos; qtd_features; user_stories_codigos; qtd_user_stories; horas_dev; horas_qa; horas_arch; horas_devops; horas_gestao; total_horas; prazo_entrega_meses; time_estimado_pessoas; valor_estimado; complexidade; stack_aderencia; premissas; comentarios`

> 💡 O discovery usa as mesmas colunas do full para padronização. No discovery, `features_codigos` e `user_stories_codigos` podem ser preenchidos com os épicos (o nível de detalhe disponível). `time_estimado_pessoas` e `valor_estimado` são **obrigatórios** — a FBSO.ORG não infere esses valores.

Pipeline: `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATION-SCHEMA.md` → Gate → Fix → COMPLIANCE

### Fase 3 — FACTORY-DISTRIBUTION.md 🆕
Registro de fábricas participantes, envio do RFQ e prazos. Pipeline: `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-FACTORY-DISTRIBUTION.md` → Gate → Fix → COMPLIANCE

### Fase 4 — ESTIMATE-RECEIPT.md 🆕
Guia para o time operacional salvar estimativas recebidas. Padrão: `nome-do-arquivo-csv-{nome-da-fabrica}.md` em `estimates/`. Pipeline: `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATE-RECEIPT.md` → Gate → Fix → COMPLIANCE

### Fase 5 — ESTIMATE-VALIDATION.md 🆕
Validação DTA de cada estimativa recebida: formato, QA balanceado (≥25%), Arquitetura (≥5%), consistência Prazo×Horas, outliers, e **PIB — Proximidade à Baseline Interna** (comparação com estimativa de referência da empresa). A baseline interna NÃO é enviada às fábricas — usada apenas na validação.

**Fontes da baseline por modo:**
- `discovery` → `upstream-architecture-discovery/DISCOVERY-LEVEL-ROM-ESTIMATE.md`
- `full` → `downstream-architecture-refinement/BOTTOM-UP-PERT-ESTIMATE.md`

Gera também arquivos individuais `ESTIMATE-VALIDATION-{FABRICA}.md` em `estimates/` com o racional detalhado de não-compliance. Pipeline: `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATE-VALIDATION.md` → Gate → Fix → COMPLIANCE

### Fase 5b — ESTIMATE-RETROSPECTIVE-PIB.md 🆕 (CONDICIONAL)

**Condição de execução:** Esta fase é executada **apenas quando 0 fábricas são aprovadas** na Fase 5. Se pelo menos 1 fábrica for aprovada, esta fase é pulada.

**Objetivo:** Análise retrospectiva aprofundada usando a métrica PIB para identificar observações adicionais de não-compliance que devem ser comunicadas às fábricas no realinhamento. Esta fase complementa a validação DTA com uma análise qualitativa que nenhuma regra automatizada captura sozinha.

**Dimensões analisadas:**

| Dimensão | O que verifica | Ação se detectado |
|:---|:---|:---|
| **PIB por Épico** | Qual fábrica mais se aproxima da baseline por épico? | Identificar se a melhor fábrica por épico ainda está muito distante |
| **Flat Estimates** | Mesmo valor para todos os épicos (CV entre épicos < 10%) | ⚠️ Alerta — fábrica não analisou escopo por épico |
| **QA/Arch como Overhead Fixo** | QA e Arch com valores absolutos iguais em todos os épicos | ⚠️ Alerta — devem ser percentuais do esforço |
| **Comentários Genéricos** | Textos idênticos entre fábricas ou genéricos na coluna `comentarios` | ⚠️ Alerta — solicitar racional detalhado |
| **Independência** | Valores idênticos entre fábricas diferentes em múltiplas colunas | 🔴 Crítico — possível violação de independência |

**Output:** `ESTIMATE-RETROSPECTIVE-PIB.md` com observações adicionais e recomendações para o realinhamento. As notificações da Fase 7 devem incorporar estas observações.

Pipeline: `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATE-RETROSPECTIVE-PIB.md` → Gate → Fix → COMPLIANCE

### Fase 6 — FACTORY-COMPARISON.md 🆕
Matriz comparativa cross-fábrica + ranking ponderado com **5 critérios**: Custo Total (25-30%), Prazo (20-25%), Qualidade Técnica QA+Arch (20%), **PIB — Proximidade à Baseline Interna (15%)** 🆕, Consistência Prazo×Horas (15%). Recomendação de seleção justificada. Pipeline: `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-FACTORY-COMPARISON.md` → Gate → Fix → COMPLIANCE

### Fase 7 — FACTORY-NOTIFICATION-{FAB}.md 🆕
Notificações formais individuais para cada fábrica em `notifications/`. Conteúdo:
- Se **aprovada:** carta de seleção (vencedora) ou segundo colocado
- Se **rejeitada:** feedback técnico detalhado com não-conformidades DTA + PIB + observações da retrospectiva (F5b, se aplicável)
- Orientações para realinhamento (6 itens: QA%, Arch%, diferenciação entre épicos, consistência prazo, comentários detalhados, reenvio)
- Confidencialidade: nome do arquivo NUNCA revela status da fábrica

Pipeline: `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-FACTORY-NOTIFICATION.md` → Gate → Fix → COMPLIANCE

---

## MECANISMO DE ORQUESTRAÇÃO DINÂMICA

Toda fase (1-7) segue o ecossistema trifásico Generate→Gate→Fix com validação humana. A Fase 5b é condicional — só executa se 0 fábricas aprovadas na F5:

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
| ⛔ Barreira C | Após Bloco C (F7) | Todas estimativas validadas (DTA + PIB). Se 0 aprovadas → F5b executada (retrospectiva). Matriz comparativa com 5 critérios (incluindo PIB 15%). Notificações geradas com observações da retrospectiva. Baseline interna consultada do modo correto (discovery→ROM, full→PERT). |

---

## ESTRUTURA DE DIRETÓRIOS GERADA

### Mode 1 — Discovery
```
business-inputs/business-projects/{PROJECT_ID_NAME}/
└── sourcing-factory-bidding-discovery/
    ├── RFQ-PACKAGE.md                         (F1)
    ├── ESTIMATION-SCHEMA.csv                  (F2)
    ├── FACTORY-DISTRIBUTION.md                (F3)
    ├── ESTIMATE-RECEIPT.md                    (F4)
    ├── ESTIMATE-VALIDATION.md                 (F5)
    ├── [ESTIMATE-RETROSPECTIVE-PIB.md]        (F5b — condicional)
    ├── FACTORY-COMPARISON.md                  (F6)
    ├── FACTORY-NOTIFICATION.md                (F7 — guia de notificações)
    ├── estimates/
    │   ├── ESTIMATION-SCHEMA-{FABRICA}.csv
    │   └── ESTIMATE-VALIDATION-{FABRICA}.md   (F5 — por fábrica)
    └── notifications/
        └── FACTORY-NOTIFICATION-{FABRICA}.md  (F7 — por fábrica)
```

### Mode 2 — Full
```
business-inputs/business-projects/{PROJECT_ID_NAME}/
└── sourcing-factory-bidding-full/
    ├── ... (mesma estrutura acima)
    ├── estimates/
    │   └── ...
    └── notifications/
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
├── PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATE-RETROSPECTIVE-PIB.md  🆕 F5b
├── PROMPT-GATE-SOURCING-FACTORY-BIDDING-ESTIMATE-RETROSPECTIVE-PIB.md      🆕 F5b
├── PROMPT-FIX-SOURCING-FACTORY-BIDDING-ESTIMATE-RETROSPECTIVE-PIB.md       🆕 F5b
├── PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-FACTORY-COMPARISON.md
├── PROMPT-GATE-SOURCING-FACTORY-BIDDING-FACTORY-COMPARISON.md
├── PROMPT-FIX-SOURCING-FACTORY-BIDDING-FACTORY-COMPARISON.md
└── (24 prompts: 8 GENERATE + 8 GATE + 8 FIX, incluindo F5b condicional)
```

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial: roadmap de Sourcing & Factory Bidding com 6 fases em 3 blocos, 2 modos (discovery/full), DTA Engine integration. | Time de Arquitetura |
| 1.1 | 31/07/2026 | **Regra PIB adicionada:** F5 inclui validação PIB (Proximidade à Baseline Interna) com fontes por modo. F6 atualizada para 5 critérios de comparação (PIB com peso 15%). Barreira C inclui verificação PIB. Skills `analyst-estimates` e `project-estimation` referenciados para PIB. Regra documentada em DTA-VALIDATION-STANDARDS.md §2.6. | Time de Arquitetura |
| 1.2 | 31/07/2026 | **Fase 5b — ESTIMATE-RETROSPECTIVE-PIB.md:** Nova fase condicional (executa apenas se 0 aprovadas na F5). Análise aprofundada: PIB por épico, detecção de flat estimates, QA/Arch como overhead fixo, comentários genéricos, independência. F7 atualizada para incorporar observações da retrospectiva nas notificações. Estrutura de diretórios inclui `notifications/`. Barreira C movida para após F7. 8 fases no total (F0-F7). | Time de Arquitetura |
| 1.3 | 31/07/2026 | **Schema Unificado:** Discovery e Full usam as mesmas colunas (20 colunas). `time_estimado_pessoas` e `valor_estimado` agora são **obrigatórios** — FBSO.ORG não infere. Colunas individuais de horas (dev/qa/arch/devops/gestao) padronizadas para ambos os modos. Modelos de exemplo salvos em `standards/`. Fórmulas DTA padronizadas: QA = qa/total, Arch = arch/total. | Time de Arquitetura |
| 1.4 | 31/07/2026 | **Auditoria de integridade:** 15 NCs corrigidas. RFQ-PACKAGE §4-5 atualizados para schema unificado de 20 colunas e 5 critérios. DTA-VALIDATION-STANDARDS §6 corrigido (contradição com §2.5/§3.2). GATEs com prefixos de ID de conflito ([SCHEMA-XX], [DIST-XX], etc.). GENERATEs F2-F7 expandidos com especificações detalhadas. FIXes F2-F7 expandidos com tabelas de priorização P0-P2. Contagem de prompts corrigida (21→24). Estruturas de diretório incluem FACTORY-NOTIFICATION.md, ESTIMATE-RETROSPECTIVE-PIB.md e ESTIMATE-VALIDATION-{FAB}.md. | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada. Skills de referência listados na seção Skills Utilizados.*
