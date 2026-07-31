# FLOWCHART: ROADMAP DE SOURCING & FACTORY BIDDING

## Versão: 1.0 — Visualização Gráfica do Pipeline de Sourcing, Bidding e Seleção de Fábricas

> **Documento de referência:** `PROMPT-ROADMAP-GENERATE-SOURCING-FACTORY-BIDDING.md` v1.0
>
> Este documento complementa o roadmap textual com diagramas Mermaid que visualizam o fluxo de execução, os dois modos de operação e o mecanismo de orquestração.

---

## 1. Visão Macro do Pipeline Completo

```mermaid
flowchart TB
    START([🚀 Início]) --> F0[Fase 0: Bootstrap]

    F0 --> MODE_DECISION{Modo de<br/>Estimativa?}
    MODE_DECISION -->|discovery| MODE1[Mode 1: Discovery-Level]
    MODE_DECISION -->|full| MODE2[Mode 2: Full-Documentation]

    MODE1 --> BLOCO_A
    MODE2 --> BLOCO_A

    subgraph BLOCO_A[Bloco A: RFQ Package]
        direction LR
        F1[F1: RFQ-PACKAGE] --> F2[F2: ESTIMATION-SCHEMA]
    end

    BLOCO_A --> BARREIRA_A{{⛔ Barreira A}}

    BARREIRA_A --> BLOCO_B

    subgraph BLOCO_B[Bloco B: Distribution & Receipt]
        direction LR
        F3[F3: FACTORY-DISTRIBUTION] --> F4[F4: ESTIMATE-RECEIPT]
    end

    BLOCO_B --> BARREIRA_B{{⛔ Barreira B}}

    BARREIRA_B --> BLOCO_C

    subgraph BLOCO_C[Bloco C: Validation & Comparison]
        direction LR
        F5[F5: ESTIMATE-VALIDATION] --> F6[F6: FACTORY-COMPARISON]
    end

    BLOCO_C --> BARREIRA_C{{⛔ Barreira C}}

    BARREIRA_C --> SELECTION[🏆 Factory Selection]

    SELECTION --> END([✅ Fábrica Selecionada])

    MODE1 -.->|pasta| DISCOVERY_PATH[📁 sourcing-factory-bidding-discovery/]
    MODE2 -.->|pasta| FULL_PATH[📁 sourcing-factory-bidding-full/]

    style F0 fill:#6c5ce7,color:#fff
    style F1 fill:#0984e3,color:#fff
    style F2 fill:#0984e3,color:#fff
    style F3 fill:#0984e3,color:#fff
    style F4 fill:#0984e3,color:#fff
    style F5 fill:#0984e3,color:#fff
    style F6 fill:#0984e3,color:#fff
    style SELECTION fill:#00b894,color:#fff
    style BARREIRA_A fill:#d63031,color:#fff
    style BARREIRA_B fill:#d63031,color:#fff
    style BARREIRA_C fill:#d63031,color:#fff
    style MODE_DECISION fill:#fdcb6e,color:#333
    style DISCOVERY_PATH fill:#dfe6e9,stroke:#636e72
    style FULL_PATH fill:#dfe6e9,stroke:#636e72
```

---

## 2. Fase 0 — Bootstrap Inteligente (Detalhado)

```mermaid
flowchart TD
    F0_START([Fase 0: Bootstrap]) --> F0_1

    subgraph PASSO_01[Passo 0.1 — Coletar Inputs]
        F0_1[Solicitar PROJECT_PATH, PROJECT_ID_NAME] --> F0_1_OPT{Solicitar<br/>SOURCING_BIDDING_MODE?}
        F0_1_OPT -->|Sim| F0_1_MODE[Coletar modo: discovery ou full]
        F0_1_OPT -->|Não| F0_2
        F0_1_MODE --> F0_2
    end

    subgraph PASSO_02[Passo 0.2 — Auditar Artefatos]
        F0_2[Verificar existência de:<br/>upstream-architecture-discovery/<br/>features/ + user-stories/] --> F0_2_DEC{Quais artefatos<br/>existem?}

        F0_2_DEC -->|"upstream-discovery/ ✅<br/>features/ ❌"| F0_3A[Exibir opção:<br/>Mode 1 - Discovery]
        F0_2_DEC -->|"upstream-discovery/ ✅<br/>features/ ✅"| F0_3B[Exibir opções:<br/>Mode 1 - Discovery<br/>Mode 2 - Full]
    end

    subgraph PASSO_03[Passo 0.3 — Decisão Humana]
        F0_3A --> F0_3_HUMAN{Humano escolhe<br/>o modo}
        F0_3B --> F0_3_HUMAN
    end

    subgraph PASSO_04[Passo 0.4 — Criar Estrutura]
        F0_3_HUMAN -->|discovery| F0_4A[mkdir -p<br/>sourcing-factory-bidding-discovery/estimates/]
        F0_3_HUMAN -->|full| F0_4B[mkdir -p<br/>sourcing-factory-bidding-full/estimates/]
    end

    subgraph PASSO_05[Passo 0.5 — Auditar Estimativas]
        F0_4A --> F0_5[Verificar estimativas já<br/>recebidas em estimates/]
        F0_4B --> F0_5
    end

    subgraph PASSO_06[Passo 0.6 — Resumo]
        F0_5 --> F0_6[Exibir resumo:<br/>modo, pasta, estimativas existentes,<br/>próxima fase a executar]
    end

    F0_6 --> ORCH

    style PASSO_01 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_02 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_03 fill:#dfe6e9,stroke:#fdcb6e
    style PASSO_04 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_05 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_06 fill:#dfe6e9,stroke:#6c5ce7
```

---

## 3. Mecanismo de Orquestração Dinâmica (Loop Trifásico)

Todas as fases 1-6 executam este mesmo loop de validação:

```mermaid
flowchart TD
    ORCH([Orquestrador: Iniciar Fase N]) --> GEN

    subgraph LOOP[Loop de Validação Soberana — Fases 1 a 6]
        GEN["1. GERAÇÃO / EVOLUÇÃO<br/>Executar PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-{FASE}.md<br/>Parâmetros: PROJECT_PATH, PROJECT_ID_NAME,<br/>SOURCING_BIDDING_MODE, SOURCING_BIDDING_PATH, ESTIMATES_PATH"] --> GATE

        GATE["2. AUDITORIA INTERNA DA IA<br/>Executar PROMPT-GATE-SOURCING-FACTORY-BIDDING-{FASE}.md"] --> GATE_RESULT{Resultado<br/>da Auditoria?}

        GATE_RESULT -->|NÃO COMPLIANCE<br/>Erros encontrados| FIX["2b. CORREÇÃO CIRÚRGICA<br/>Executar PROMPT-FIX-SOURCING-FACTORY-BIDDING-{FASE}.md<br/>Apenas nas seções afetadas"]
        FIX --> GATE

        GATE_RESULT -->|SEM ERROS| HUMAN_GATE

        HUMAN_GATE[3. PORTÃO DE VALIDAÇÃO HUMANA<br/>Status: PRÉ-COMPLIANCE INTERNO<br/>— AGUARDANDO VALIDAÇÃO HUMANA] --> P1[P1: Documento aderente<br/>às necessidades?]
        P1 --> P2[P2: Deseja anexar novos<br/>documentos de entrada?]
        P2 --> P3[P3: Deseja fornecer novos<br/>inputs textuais ou mudanças<br/>de escopo?]
        P3 --> HUMAN_DEC{Decisão<br/>do Humano?}

        HUMAN_DEC -->|Aprova sem novos inputs| COMPLIANCE
        HUMAN_DEC -->|Fornece novos docs/inputs| GEN
    end

    COMPLIANCE([✅ STATUS: COMPLIANCE<br/>Arquivo congelado<br/>Próxima fase destravada])

    style GEN fill:#0984e3,color:#fff
    style GATE fill:#fdcb6e,color:#333
    style FIX fill:#e17055,color:#fff
    style HUMAN_GATE fill:#6c5ce7,color:#fff
    style P1 fill:#dfe6e9,stroke:#6c5ce7
    style P2 fill:#dfe6e9,stroke:#6c5ce7
    style P3 fill:#dfe6e9,stroke:#6c5ce7
    style COMPLIANCE fill:#00b894,color:#fff
    style LOOP fill:#fff3e0,stroke:#f39c12
```

---

## 4. Pipeline Sequencial — Fases 1-6 com Dependências

```mermaid
flowchart LR
    subgraph F1[Fase 1: RFQ-PACKAGE]
        direction TB
        F1_GEN[GENERATE<br/>Compila artefatos] --> F1_GATE[GATE<br/>Valida pacote] --> F1_FIX[FIX<br/>Corrige gaps]
        F1_FIX -.->|loop| F1_GATE
        F1_GATE --> F1_OK[COMPLIANCE ✅]
    end

    subgraph F2[Fase 2: ESTIMATION-SCHEMA]
        direction TB
        F2_GEN[GENERATE<br/>Template CSV] --> F2_GATE[GATE<br/>Valida schema] --> F2_FIX[FIX<br/>Corrige colunas]
        F2_FIX -.->|loop| F2_GATE
        F2_GATE --> F2_OK[COMPLIANCE ✅]
    end

    subgraph F3[Fase 3: FACTORY-DISTRIBUTION]
        direction TB
        F3_GEN[GENERATE<br/>Registra fábricas] --> F3_GATE[GATE<br/>Valida registros] --> F3_FIX[FIX<br/>Corrige dados]
        F3_FIX -.->|loop| F3_GATE
        F3_GATE --> F3_OK[COMPLIANCE ✅]
    end

    subgraph F4[Fase 4: ESTIMATE-RECEIPT]
        direction TB
        F4_GEN[GENERATE<br/>Guia recebimento] --> F4_GATE[GATE<br/>Valida estimativas<br/>em estimates/] --> F4_FIX[FIX<br/>Corrige formato]
        F4_FIX -.->|loop| F4_GATE
        F4_GATE --> F4_OK[COMPLIANCE ✅]
    end

    subgraph F5[Fase 5: ESTIMATE-VALIDATION]
        direction TB
        F5_GEN[GENERATE<br/>Valida regras DTA] --> F5_GATE[GATE<br/>Verifica QA, outliers] --> F5_FIX[FIX<br/>Corrige validação]
        F5_FIX -.->|loop| F5_GATE
        F5_GATE --> F5_OK[COMPLIANCE ✅]
    end

    subgraph F6[Fase 6: FACTORY-COMPARISON]
        direction TB
        F6_GEN[GENERATE<br/>Matriz comparativa] --> F6_GATE[GATE<br/>Valida ranking] --> F6_FIX[FIX<br/>Corrige análise]
        F6_FIX -.->|loop| F6_GATE
        F6_GATE --> F6_OK[COMPLIANCE ✅]
    end

    F1_OK -->|destrava| F2_GEN
    F2_OK -->|destrava| F3_GEN
    F3_OK -->|destrava| F4_GEN
    F4_OK -->|destrava| F5_GEN
    F5_OK -->|destrava| F6_GEN

    F1 --> I1[📥 upstream-architecture-discovery/<br/>PRD + Bloco B + SPECS + ROM]
    F2 --> I2[📥 DTA Estimation Schema<br/>+ BACKLOG-LIST.csv]
    F3 --> I3[📥 RFQ-PACKAGE.md<br/>+ ESTIMATION-SCHEMA.csv]
    F4 --> I4[📥 estimates/<br/>nome-do-arquivo-csv-{fabrica}.md]
    F5 --> I5[📥 DTA Validation Rules<br/>QA ≥ 20% · Arch ≥ 5% · Outliers]
    F6 --> I6[📥 Todas estimativas<br/>validadas (F5)]

    style F1 fill:#0984e3,color:#fff
    style F2 fill:#0984e3,color:#fff
    style F3 fill:#0984e3,color:#fff
    style F4 fill:#0984e3,color:#fff
    style F5 fill:#0984e3,color:#fff
    style F6 fill:#0984e3,color:#fff
```

---

## 5. Modos de Operação — Discovery vs Full

```mermaid
flowchart TD
    START([Bootstrap: Modo Selecionado]) --> MODE{SOURCING_BIDDING_MODE?}

    MODE -->|discovery| D_MODE[Mode 1: Discovery-Level]
    MODE -->|full| F_MODE[Mode 2: Full-Documentation]

    subgraph D_MODE[Mode 1 — Discovery-Level]
        D_PATH[📁 sourcing-factory-bidding-discovery/]
        D_F1[F1: RFQ com PRD + Bloco B + SPECS + ROM]
        D_F2[F2: Schema macro — épicos × soluções]
        D_F5[F5: Validação DTA adaptada — nível épico]
        D_F6[F6: Comparação por épico + ROM baseline]
        D_PATH --> D_F1 --> D_F2
        D_F2 -.-> D_F5 -.-> D_F6
    end

    subgraph F_MODE[Mode 2 — Full-Documentation]
        F_PATH[📁 sourcing-factory-bidding-full/]
        F_F1[F1: RFQ com Features + US + RTM]
        F_F2[F2: Schema detalhado — features × US]
        F_F5[F5: Validação DTA completa — nível US]
        F_F6[F6: Comparação detalhada US por US]
        F_PATH --> F_F1 --> F_F2
        F_F2 -.-> F_F5 -.-> F_F6
    end

    D_F6 --> SELECTION[🏆 Factory Selection]
    F_F6 --> SELECTION

    style D_MODE fill:#fff3e0,stroke:#e65100
    style F_MODE fill:#f3e5f5,stroke:#7b1fa2
    style D_PATH fill:#dfe6e9,stroke:#636e72
    style F_PATH fill:#dfe6e9,stroke:#636e72
    style SELECTION fill:#00b894,color:#fff
```

---

## 6. Fluxo de Estimativa — Da Fábrica à Seleção

```mermaid
flowchart LR
    subgraph ENVIO[Envio para Fábricas]
        RFQ[📦 RFQ Package<br/>F1 + F2] -->|email| FAB_A[🏢 Fábrica A]
        RFQ -->|email| FAB_B[🏢 Fábrica B]
        RFQ -->|email| FAB_C[🏢 Fábrica C]
    end

    subgraph RETORNO[Retorno das Fábricas]
        FAB_A -->|CSV preenchido| EST_A[nome-do-arquivo-csv-fabrica-A.md]
        FAB_B -->|CSV preenchido| EST_B[nome-do-arquivo-csv-fabrica-B.md]
        FAB_C -->|CSV preenchido| EST_C[nome-do-arquivo-csv-fabrica-C.md]
    end

    subgraph VALIDACAO[Validação DTA — F5]
        EST_A --> VAL_A{QA ≥ 20%?<br/>Arch ≥ 5%?<br/>Formato OK?}
        EST_B --> VAL_B{QA ≥ 20%?<br/>Arch ≥ 5%?<br/>Formato OK?}
        EST_C --> VAL_C{QA ≥ 20%?<br/>Arch ≥ 5%?<br/>Formato OK?}
        VAL_A -->|✅| OK_A[APROVADA]
        VAL_A -->|❌| REJ_A[REJEITADA]
        VAL_B -->|✅| OK_B[APROVADA]
        VAL_B -->|❌| REJ_B[REJEITADA]
        VAL_C -->|✅| OK_C[APROVADA]
        VAL_C -->|❌| REJ_C[REJEITADA]
    end

    subgraph COMPARACAO[Comparação — F6]
        OK_A --> MATRIZ[Matriz Comparativa<br/>Custo × Prazo × Qualidade]
        OK_B --> MATRIZ
        OK_C --> MATRIZ
        MATRIZ --> RANKING[🏆 Ranking + Recomendação]
    end

    RANKING --> FIM[✅ Factory Selection]

    style ENVIO fill:#e3f2fd,stroke:#0984e3
    style RETORNO fill:#fff3e0,stroke:#e65100
    style VALIDACAO fill:#e8f5e9,stroke:#00b894
    style COMPARACAO fill:#f3e5f5,stroke:#6c5ce7
    style FIM fill:#00b894,color:#fff
```

---

## 7. Diagrama de Estados — Visão Unificada

```mermaid
stateDiagram-v2
    [*] --> Bootstrap: Início do Sourcing

    state Bootstrap {
        [*] --> ColetarInputs
        ColetarInputs --> AuditarArtefatos
        AuditarArtefatos --> EscolherModo
        EscolherModo --> CriarEstrutura
        CriarEstrutura --> AuditarEstimativas
        AuditarEstimativas --> ResumoInicial
    }

    Bootstrap --> F1_RFQ

    state F1_RFQ {
        [*] --> F1_Gen
        F1_Gen --> F1_Gate
        F1_Gate --> F1_Fix: NÃO COMPLIANCE
        F1_Fix --> F1_Gate
        F1_Gate --> F1_Human: SEM ERROS
        F1_Human --> F1_Gen: Novos inputs
        F1_Human --> F1_Done: Aprovado
    }

    F1_RFQ --> F2_Schema: COMPLIANCE

    state F2_Schema {
        [*] --> F2_Gen
        F2_Gen --> F2_Gate
        F2_Gate --> F2_Fix: NÃO COMPLIANCE
        F2_Fix --> F2_Gate
        F2_Gate --> F2_Human: SEM ERROS
        F2_Human --> F2_Gen: Novos inputs
        F2_Human --> F2_Done: Aprovado
    }

    F2_Schema --> BarreiraA: COMPLIANCE

    state BarreiraA {
        [*] --> CheckA
        CheckA --> F3_Dist: Bloco A 100% OK
    }

    BarreiraA --> F3_Dist

    state F3_Dist {
        [*] --> F3_Gen
        F3_Gen --> F3_Gate
        F3_Gate --> F3_Fix: NÃO COMPLIANCE
        F3_Fix --> F3_Gate
        F3_Gate --> F3_Human: SEM ERROS
        F3_Human --> F3_Gen: Novos inputs
        F3_Human --> F3_Done: Aprovado
    }

    F3_Dist --> F4_Receipt: COMPLIANCE

    state F4_Receipt {
        [*] --> F4_Gen
        F4_Gen --> F4_Gate
        F4_Gate --> F4_Fix: NÃO COMPLIANCE
        F4_Fix --> F4_Gate
        F4_Gate --> F4_Human: SEM ERROS
        F4_Human --> F4_Gen: Novos inputs
        F4_Human --> F4_Done: Aprovado
    }

    F4_Receipt --> BarreiraB: COMPLIANCE

    state BarreiraB {
        [*] --> CheckB
        CheckB --> F5_Valid: Bloco B 100% OK
    }

    BarreiraB --> F5_Valid

    state F5_Valid {
        [*] --> F5_Gen
        F5_Gen --> F5_Gate
        F5_Gate --> F5_Fix: NÃO COMPLIANCE
        F5_Fix --> F5_Gate
        F5_Gate --> F5_Human: SEM ERROS
        F5_Human --> F5_Gen: Novos inputs
        F5_Human --> F5_Done: Aprovado
    }

    F5_Valid --> F6_Compare: COMPLIANCE

    state F6_Compare {
        [*] --> F6_Gen
        F6_Gen --> F6_Gate
        F6_Gate --> F6_Fix: NÃO COMPLIANCE
        F6_Fix --> F6_Gate
        F6_Gate --> F6_Human: SEM ERROS
        F6_Human --> F6_Gen: Novos inputs
        F6_Human --> F6_Done: Aprovado
    }

    F6_Compare --> BarreiraC: COMPLIANCE

    state BarreiraC {
        [*] --> CheckC
        CheckC --> Selection: Bloco C 100% OK
    }

    BarreiraC --> Selection

    state Selection {
        [*] --> Choose
        Choose --> [*]: 🏆 Fábrica Selecionada
    }

    Selection --> [*]
```

---

## 8. Integração com os Demais Roadmaps

```mermaid
flowchart LR
    subgraph UPSTREAM[Upstream Architecture Discovery]
        direction TB
        U1[PRD + Bloco B] --> U2[SPECS + ROM]
        U2 --> U3[GO/NO-GO]
    end

    subgraph SOURCING[Sourcing & Factory Bidding — Este Documento]
        direction TB
        S1[RFQ Package] --> S2[Factory Distribution]
        S2 --> S3[Validation + Comparison]
        S3 --> S4[Factory Selection]
    end

    subgraph TECHDEF[Technical Definitions]
        direction TB
        T1[Bloco 0] --> T2[Bloco A → B]
        T2 --> T3[Bloco C → D]
    end

    U3 -->|"Go-Ahead ✅<br/>dispara sourcing"| S1
    S4 -->|"Fábrica selecionada<br/>dispara definições detalhadas"| T1

    style UPSTREAM fill:#e3f2fd,stroke:#0984e3
    style SOURCING fill:#fff3e0,stroke:#e65100
    style TECHDEF fill:#e8f5e9,stroke:#00b894
```

### Posicionamento na Cadeia de Valor

| Roadmap | Nível | Output | Consumido por |
|---------|-------|--------|---------------|
| **Upstream Architecture Discovery** | Estratégico / Viabilidade | PRD, 6 disciplinas, SPECS, ROM, GO/NO-GO | → Sourcing & Factory Bidding (RFQ) |
| **Sourcing & Factory Bidding** ← | Tático / Sourcing | RFQ Package, Matriz Comparativa, Factory Selection | → Technical Definitions (detalhamento) |
| **Technical Definitions** | Tático / Implementação | 20 fases de definições detalhadas | → Times de Desenvolvimento |

---

## 9. Tabela de Símbolos e Convenções

| Símbolo/Cor | Significado |
|-------------|-------------|
| 🟣 Roxo (`#6c5ce7`) | Bootstrap / Validação Humana / Barreiras |
| 🔵 Azul (`#0984e3`) | Fases de Geração padrão (1-6) |
| 🟠 Laranja (`#e17055`) | Correções (FIX) |
| 🟢 Verde (`#00b894`) | Compliance / Factory Selection |
| 🟡 Amarelo (`#fdcb6e`) | Gate / Decisões |
| 🔴 Vermelho (`#d63031`) | Barreiras de bloqueio |
| 🔲 Linha tracejada | Loop de retrabalho (GATE→FIX→GATE) |
| 🔲 Linha sólida | Fluxo sequencial normal |
| 📁 | Pasta de trabalho |
| 🏢 | Fábrica de software |
| 🏆 | Seleção / Vencedor |

---

> **📁 Arquivos relacionados:**
> - `PROMPT-ROADMAP-GENERATE-SOURCING-FACTORY-BIDDING.md` — Documento fonte (v1.0)
> - `../PROMPT-ROADMAP-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY.md` — Roadmap upstream (alimenta o RFQ)
> - `../PROMPT-ROADMAP-GENERATE-PROJECT-TECHNICAL-DEFINITIONS.md` — Roadmap técnico (consumido após seleção)
> - `../../DTA-Engine-de-Bidding-e-Estimativas.md` — Engine de validação DTA
> - `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-*.md` — 6 prompts geradores (fases 1-6)
> - `PROMPT-GATE-SOURCING-FACTORY-BIDDING-*.md` — 6 prompts de auditoria (fases 1-6)
> - `PROMPT-FIX-SOURCING-FACTORY-BIDDING-*.md` — 6 prompts de correção (fases 1-6)
