# FLOWCHART: ROADMAP DE SOURCING & FACTORY BIDDING

## Versão: 2.0 — Visualização Gráfica com 4 Modos (Agile Discovery, Agile Refinement, Waterfall Discovery, Waterfall Refinement)

> **Documento de referência:** `PROMPT-ROADMAP-GENERATE-SOURCING-FACTORY-BIDDING.md` v1.6
>
> Este documento complementa o roadmap textual com diagramas Mermaid que visualizam o fluxo de execução, os **quatro modos** de operação e o mecanismo de orquestração.

---

## 1. Visão Macro do Pipeline Completo

```mermaid
flowchart TB
    START([🚀 Início]) --> F0[Fase 0: Bootstrap]

    F0 --> MODE_DECISION{Modo de<br/>Estimativa?}

    MODE_DECISION -->|agile-discovery| MODE1[Mode 1: Agile Discovery]
    MODE_DECISION -->|agile-refinement| MODE2[Mode 2: Agile Refinement]
    MODE_DECISION -->|waterfall-discovery| MODE3[Mode 3: Waterfall Discovery]
    MODE_DECISION -->|waterfall-refinement| MODE4[Mode 4: Waterfall Refinement]

    MODE1 --> BLOCO_A
    MODE2 --> BLOCO_A
    MODE3 --> BLOCO_A
    MODE4 --> BLOCO_A

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

    MODE1 -.->|pasta| AGILE_DISCOVERY_PATH[📁 sourcing-factory-bidding-agile-discovery/]
    MODE2 -.->|pasta| AGILE_REFINEMENT_PATH[📁 sourcing-factory-bidding-agile-refinement/]
    MODE3 -.->|pasta| WATERFALL_DISCOVERY_PATH[📁 sourcing-factory-bidding-waterfall-discovery/]
    MODE4 -.->|pasta| WATERFALL_REFINEMENT_PATH[📁 sourcing-factory-bidding-waterfall-refinement/]

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
    style AGILE_DISCOVERY_PATH fill:#808080,stroke:#636e72
    style AGILE_REFINEMENT_PATH fill:#808080,stroke:#636e72
    style WATERFALL_DISCOVERY_PATH fill:#808080,stroke:#636e72
    style WATERFALL_REFINEMENT_PATH fill:#808080,stroke:#636e72
```

---

## 2. Fase 0 — Bootstrap Inteligente (Detalhado)

```mermaid
flowchart TD
    F0_START([Fase 0: Bootstrap]) --> F0_1

    subgraph PASSO_01[Passo 0.1 — Coletar Inputs]
        F0_1[Solicitar PROJECT_PATH, PROJECT_ID_NAME] --> F0_1_OPT{Solicitar<br/>SOURCING_BIDDING_MODE?}
        F0_1_OPT -->|Sim| F0_1_MODE[Coletar modo:<br/>agile-discovery, agile-refinement,<br/>waterfall-discovery ou<br/>waterfall-refinement]
        F0_1_OPT -->|Não| F0_2
        F0_1_MODE --> F0_2
    end

    subgraph PASSO_02[Passo 0.2 — Auditar Artefatos]
        F0_2["Verificar existência de:<br/>1️⃣ upstream-architecture-discovery/<br/>2️⃣ features/ + user-stories/<br/>3️⃣ 01-Charter + 06-HLD WATERFALL<br/>4️⃣ 07-LLD + 11-EAP/WBS WATERFALL<br/>5️⃣ waterfall-estimation/"] --> F0_2_DEC{"Quais<br/>artefatos<br/>existem?"}

        F0_2_DEC -->|"upstream-discovery ✅<br/>features ❌<br/>WATERFALL ❌"| F0_3A["Exibir opção:<br/>Mode 1 — Agile Discovery"]
        F0_2_DEC -->|"upstream-discovery ✅<br/>features ✅<br/>WATERFALL ❌"| F0_3B["Exibir opções:<br/>Mode 1 — Agile Discovery<br/>Mode 2 — Agile Refinement"]
        F0_2_DEC -->|"WATERFALL upstream ✅<br/>waterfall-estimation ✅"| F0_3C["Exibir opção:<br/>Mode 3 — Waterfall Discovery"]
        F0_2_DEC -->|"WATERFALL downstream ✅<br/>waterfall-estimation ✅"| F0_3D["Exibir opção:<br/>Mode 4 — Waterfall Refinement"]
        F0_2_DEC -->|"Múltiplos conjuntos ✅"| F0_3E["Exibir todas<br/>as opções disponíveis<br/>(até 4 modos)"]
    end

    subgraph PASSO_03[Passo 0.3 — Decisão Humana]
        F0_3A --> F0_3_HUMAN{Humano escolhe<br/>o modo}
        F0_3B --> F0_3_HUMAN
        F0_3C --> F0_3_HUMAN
        F0_3D --> F0_3_HUMAN
        F0_3E --> F0_3_HUMAN
    end

    subgraph PASSO_04[Passo 0.4 — Criar Estrutura]
        F0_3_HUMAN -->|agile-discovery| F0_4A[mkdir -p<br/>sourcing-factory-bidding-agile-discovery/estimates/]
        F0_3_HUMAN -->|agile-refinement| F0_4B[mkdir -p<br/>sourcing-factory-bidding-agile-refinement/estimates/]
        F0_3_HUMAN -->|waterfall-discovery| F0_4C[mkdir -p<br/>sourcing-factory-bidding-waterfall-discovery/estimates/]
        F0_3_HUMAN -->|waterfall-refinement| F0_4D[mkdir -p<br/>sourcing-factory-bidding-waterfall-refinement/estimates/]
    end

    subgraph PASSO_05[Passo 0.5 — Auditar Estimativas]
        F0_4A --> F0_5[Verificar estimativas já<br/>recebidas em estimates/]
        F0_4B --> F0_5
        F0_4C --> F0_5
        F0_4D --> F0_5
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
    style F0_3E fill:#ffeaa7,stroke:#fdcb6e,color:#333
    style F0_4A fill:#dfe6e9,stroke:#00b894,color:black
    style F0_4B fill:#dfe6e9,stroke:#00b894,color:black
    style F0_4C fill:#e0f7fa,stroke:#00b894,color:black
    style F0_4D fill:#e0f7fa,stroke:#00b894,color:black

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
    style P1 fill:#dfe6e9,stroke:#6c5ce7,color:black
    style P2 fill:#dfe6e9,stroke:#6c5ce7,color:black
    style P3 fill:#dfe6e9,stroke:#6c5ce7,color:black
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

    F1 --> I1["📥 Fontes por modo:<br/>Agile Discovery/Waterfall Discovery: PRD + Épicos<br/>Agile Refinement/Waterfall Refinement: Features + US + EAP"]
    F2 --> I2["📥 DTA Estimation Schema<br/>20 colunas — schema unificado"]
    F3 --> I3["📥 RFQ-PACKAGE.md<br/>+ ESTIMATION-SCHEMA.csv"]
    F4 --> I4["📥 estimates/<br/>nome-do-arquivo-csv-{fabrica}.md"]
    F5 --> I5["📥 DTA Validation Rules<br/>QA ≥ 25% · Arch ≥ 5% · Outliers<br/>+ PIB por modo (ROM/PERT)"]
    F6 --> I6["📥 Todas estimativas<br/>validadas (F5)"]

    style F1 fill:#0984e3,color:#fff
    style F2 fill:#0984e3,color:#fff
    style F3 fill:#0984e3,color:#fff
    style F4 fill:#0984e3,color:#fff
    style F5 fill:#0984e3,color:#fff
    style F6 fill:#0984e3,color:#fff
```

---

## 5. Modos de Operação — 4 Modos (Agile + Waterfall)

```mermaid
flowchart TD
    START([Bootstrap: Modo Selecionado]) --> MODE{SOURCING_BIDDING_MODE?}

    MODE -->|agile-discovery| AD_MODE[Mode 1: Agile Discovery]
    MODE -->|agile-refinement| AR_MODE[Mode 2: Agile Refinement]
    MODE -->|waterfall-discovery| WD_MODE[Mode 3: Waterfall Discovery]
    MODE -->|waterfall-refinement| WR_MODE[Mode 4: Waterfall Refinement]

    subgraph AD_MODE[Mode 1 — Agile Discovery]
        AD_PATH[📁 sourcing-factory-bidding-agile-discovery/]
        AD_F1[F1: RFQ com PRD + Bloco B + SPECS + ROM]
        AD_F2[F2: Schema macro — épicos × soluções]
        AD_F5[F5: Validação DTA adaptada — nível épico]
        AD_F6[F6: Comparação por épico + ROM baseline]
        AD_BASELINE[📊 Baseline PIB: upstream-architecture-discovery/ROM]
        AD_PATH --> AD_F1 --> AD_F2
        AD_F2 -.-> AD_F5 -.-> AD_F6
        AD_F5 -.-> AD_BASELINE
    end

    subgraph AR_MODE[Mode 2 — Agile Refinement]
        AR_PATH[📁 sourcing-factory-bidding-agile-refinement/]
        AR_F1[F1: RFQ com Features + US + RTM]
        AR_F2[F2: Schema detalhado — features × US]
        AR_F5[F5: Validação DTA completa — nível US]
        AR_F6[F6: Comparação detalhada US por US]
        AR_BASELINE[📊 Baseline PIB: downstream-architecture-refinement/PERT]
        AR_PATH --> AR_F1 --> AR_F2
        AR_F2 -.-> AR_F5 -.-> AR_F6
        AR_F5 -.-> AR_BASELINE
    end

    subgraph WD_MODE[Mode 3 — Waterfall Discovery]
        WD_PATH[📁 sourcing-factory-bidding-waterfall-discovery/]
        WD_F1[F1: RFQ com Charter + BRD + SAD + HLD]
        WD_F2[F2: Schema macro — épicos/componentes]
        WD_F5[F5: Validação DTA adaptada — nível componente]
        WD_F6[F6: Comparação por componente + ROM baseline]
        WD_BASELINE[📊 Baseline PIB: waterfall-estimation/UPSTREAM-ROM]
        WD_PATH --> WD_F1 --> WD_F2
        WD_F2 -.-> WD_F5 -.-> WD_F6
        WD_F5 -.-> WD_BASELINE
    end

    subgraph WR_MODE[Mode 4 — Waterfall Refinement]
        WR_PATH[📁 sourcing-factory-bidding-waterfall-refinement/]
        WR_F1[F1: RFQ com SRS + RTM + LLD + EAP/WBS]
        WR_F2[F2: Schema detalhado — pacotes EAP]
        WR_F5[F5: Validação DTA completa — nível pacote EAP]
        WR_F6[F6: Comparação detalhada por pacote EAP]
        WR_BASELINE[📊 Baseline PIB: waterfall-estimation/DOWNSTREAM-PERT]
        WR_PATH --> WR_F1 --> WR_F2
        WR_F2 -.-> WR_F5 -.-> WR_F6
        WR_F5 -.-> WR_BASELINE
    end

    AD_F6 --> SELECTION[🏆 Factory Selection]
    AR_F6 --> SELECTION
    WD_F6 --> SELECTION
    WR_F6 --> SELECTION

    style AD_MODE fill:#fff3e0,stroke:#e65100
    style AR_MODE fill:#f3e5f5,stroke:#7b1fa2
    style WD_MODE fill:#e0f7fa,stroke:#00838f
    style WR_MODE fill:#e8f5e9,stroke:#2e7d32
    style AD_PATH fill:#dfe6e9,stroke:#636e72,color:black
    style AR_PATH fill:#dfe6e9,stroke:#636e72,color:black
    style WD_PATH fill:#dfe6e9,stroke:#636e72,color:black
    style WR_PATH fill:#dfe6e9,stroke:#636e72,color:black
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

    subgraph VALIDACAO[Validação DTA + PIB — F5]
        EST_A --> VAL_A{"QA ≥ 25%?<br/>Arch ≥ 5%?<br/>Formato OK?<br/>PIB Score OK?"}
        EST_B --> VAL_B{"QA ≥ 25%?<br/>Arch ≥ 5%?<br/>Formato OK?<br/>PIB Score OK?"}
        EST_C --> VAL_C{"QA ≥ 25%?<br/>Arch ≥ 5%?<br/>Formato OK?<br/>PIB Score OK?"}
        VAL_A -->|✅| OK_A[APROVADA]
        VAL_A -->|❌| REJ_A[REJEITADA]
        VAL_B -->|✅| OK_B[APROVADA]
        VAL_B -->|❌| REJ_B[REJEITADA]
        VAL_C -->|✅| OK_C[APROVADA]
        VAL_C -->|❌| REJ_C[REJEITADA]
    end

    subgraph COMPARACAO[Comparação — F6: 5 Critérios]
        OK_A --> MATRIZ[Matriz Comparativa<br/>Custo 25-30% · Prazo 20-25%<br/>Qualidade QA+Arch 20%<br/>PIB 15% · Consistência 15%]
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
        AuditarArtefatos --> Detectar4Modos: Verifica Agile + WATERFALL
        Detectar4Modos --> EscolherModo: agile-discovery | agile-refinement | waterfall-discovery | waterfall-refinement
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

    subgraph WATERFALL_EST[WATERFALL-ESTIMATION]
        direction TB
        W1[Upstream ROM ±50%] --> W2[GO/NO-GO Governance]
        W3[Downstream PERT ±15-25%] --> W4[Cronograma + Orçamento]
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

    U3 -->|"Go-Ahead ✅<br/>dispara sourcing<br/>modo: agile-discovery"| S1
    W2 -->|"Go-Ahead ✅<br/>dispara sourcing<br/>modo: waterfall-discovery"| S1
    W4 -->|"PERT concluído ✅<br/>dispara sourcing<br/>modo: waterfall-refinement"| S1
    S4 -->|"Fábrica selecionada<br/>dispara definições detalhadas"| T1

    style UPSTREAM fill:#e3f2fd,stroke:#0984e3
    style WATERFALL_EST fill:#e0f7fa,stroke:#00838f
    style SOURCING fill:#fff3e0,stroke:#e65100
    style TECHDEF fill:#e8f5e9,stroke:#00b894
```

### Posicionamento na Cadeia de Valor

| Roadmap | Nível | Output | Consumido por |
|---------|-------|--------|---------------|
| **Upstream Architecture Discovery** | Estratégico / Viabilidade | PRD, 6 disciplinas, SPECS, ROM, GO/NO-GO | → Sourcing & Factory Bidding (modo: agile-discovery) |
| **WATERFALL-ESTIMATION** | Estratégico / Viabilidade | ROM ±50% (upstream), PERT ±15-25% (downstream), Cronograma, Orçamento, GO/NO-GO | → Sourcing & Factory Bidding (modos: waterfall-discovery, waterfall-refinement) |
| **Sourcing & Factory Bidding** ← | Tático / Sourcing | RFQ Package, Matriz Comparativa, Factory Selection — 4 modos | → Technical Definitions (detalhamento) |
| **Technical Definitions** | Tático / Implementação | 20 fases de definições detalhadas | → Times de Desenvolvimento |

---

## 9. Tabela de Símbolos e Convenções

| Símbolo/Cor | Significado |
|-------------|-------------|
| 🟣 Roxo (`#6c5ce7`) | Bootstrap / Validação Humana / Barreiras |
| 🔵 Azul (`#0984e3`) | Fases de Geração padrão (1-6) |
| 🟠 Laranja (`#e17055`) | Correções (FIX) |
| 🟢 Verde (`#00b894`) | Compliance / Factory Selection / Waterfall |
| 🟡 Amarelo (`#fdcb6e`) | Gate / Decisões |
| 🔴 Vermelho (`#d63031`) | Barreiras de bloqueio |
| ⬜ Cinza (`#808080`) | Pastas de trabalho |
| 🟤 Laranja escuro (`#e65100`) | Modo Agile Discovery |
| 🟣 Roxo escuro (`#7b1fa2`) | Modo Agile Refinement |
| 🔵 Ciano escuro (`#00838f`) | Modo Waterfall Discovery |
| 🟢 Verde escuro (`#2e7d32`) | Modo Waterfall Refinement |
| 🔲 Linha tracejada | Loop de retrabalho (GATE→FIX→GATE) |
| 🔲 Linha sólida | Fluxo sequencial normal |
| 📁 | Pasta de trabalho |
| 🏢 | Fábrica de software |
| 🏆 | Seleção / Vencedor |
| 📊 | Baseline PIB |
| 1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ | Passos da Fase 0 (Bootstrap) |

---

> **📁 Arquivos relacionados:**
> - `PROMPT-ROADMAP-GENERATE-SOURCING-FACTORY-BIDDING.md` — Documento fonte (v1.6)
> - `../PROMPT-ROADMAP-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY.md` — Roadmap upstream Agile (alimenta RFQ modo agile-discovery)
> - `../PROMPT-ROADMAP-GENERATE-WATERFALL-ESTIMATION.md` — Roadmap WATERFALL-ESTIMATION (alimenta RFQ modos waterfall-discovery/refinement)
> - `../PROMPT-ROADMAP-GENERATE-PROJECT-TECHNICAL-DEFINITIONS.md` — Roadmap técnico (consumido após seleção)
> - `../../standards/DTA-Engine-de-Bidding-e-Estimativas.md` — Engine de validação DTA
> - `../../standards/DTA-VALIDATION-STANDARDS.md` — Regras DTA, PIB e 5 critérios de comparação
> - `PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-*.md` — 8 prompts geradores (fases 1-7 + F5b)
> - `PROMPT-GATE-SOURCING-FACTORY-BIDDING-*.md` — 8 prompts de auditoria
> - `PROMPT-FIX-SOURCING-FACTORY-BIDDING-*.md` — 8 prompts de correção
