# FLOWCHART: ROADMAP DE WATERFALL-ESTIMATION

## Versão: 1.0 — Visualização Gráfica dos 2 Modos (UPSTREAM/DISCOVERY ROM ±50% + DOWNSTREAM/REFINEMENT PERT ±15-25%)

> **Documento de referência:** `PROMPT-ROADMAP-GENERATE-WATERFALL-ESTIMATION.md` v1.0
>
> Este documento complementa o roadmap textual com diagramas Mermaid que visualizam o fluxo de execução, os **dois modos** de estimativa, o mecanismo de orquestração e a integração com os documentos WATERFALL.

---

## 1. Visão Macro do Pipeline Completo

```mermaid
flowchart TB
    START(["🚀 Início"]) --> F0["Fase 0: Bootstrap"]

    F0 --> MODE_DECISION{Modo de<br/>Estimativa?}

    MODE_DECISION -->|upstream-discovery| UPSTREAM_MODE["UPSTREAM/DISCOVERY<br/>ROM ±50%<br/>Fases 1 → 2 → 3"]
    MODE_DECISION -->|downstream-refinement| DOWNSTREAM_MODE["DOWNSTREAM/REFINEMENT<br/>PERT ±15-25%<br/>Fases 4 → 5 → 6 → 7"]

    UPSTREAM_MODE --> BLOCO_UPSTREAM
    DOWNSTREAM_MODE --> BLOCO_DOWNSTREAM

    subgraph BLOCO_UPSTREAM["UPSTREAM/DISCOVERY — ROM ±50%"]
        direction LR
        F1["F1: UPSTREAM-ROM"] --> F2["F2: SCOPE-SNAPSHOT"]
        F2 --> F3["F3: GOVERNANCE-ROM-REPORT"]
    end

    subgraph BLOCO_DOWNSTREAM["DOWNSTREAM/REFINEMENT — PERT ±15-25%"]
        direction LR
        F4["F4: DOWNSTREAM-PERT"] --> F5["F5: SCOPE-SNAPSHOT"]
        F5 --> F6["F6: CRONOGRAMA-CALCULADO"]
        F6 --> F7["F7: ORCAMENTO-CALCULADO"]
    end

    BLOCO_UPSTREAM --> BARREIRA_UPSTREAM{{"⛔ Barreira UPSTREAM"}}
    BARREIRA_UPSTREAM --> GONOGO{GO / NO-GO?}

    GONOGO -->|GO ✅| CONTINUA_WATERFALL["Continua WATERFALL<br/>Doc #07 LLD"]
    GONOGO -->|NO-GO ❌| CANCELADO["Projeto Cancelado<br/>ou Arquivado"]

    BLOCO_DOWNSTREAM --> BARREIRA_DOWNSTREAM{{"⛔ Barreira DOWNSTREAM"}}
    BARREIRA_DOWNSTREAM --> ALIMENTA_DOCS["Alimenta Docs WATERFALL"]

    ALIMENTA_DOCS --> DOC12["📄 Doc #12: Cronograma/Gantt"]
    ALIMENTA_DOCS --> DOC13["📄 Doc #13: Orçamento"]

    UPSTREAM_MODE -.->|pasta| EST_PATH["📁 waterfall-estimation/"]
    DOWNSTREAM_MODE -.->|pasta| EST_PATH

    style F0 fill:#6c5ce7,color:#fff
    style F1 fill:#0984e3,color:#fff
    style F2 fill:#0984e3,color:#fff
    style F3 fill:#0984e3,color:#fff
    style F4 fill:#00b894,color:#fff
    style F5 fill:#00b894,color:#fff
    style F6 fill:#00b894,color:#fff
    style F7 fill:#00b894,color:#fff
    style BARREIRA_UPSTREAM fill:#d63031,color:#fff
    style BARREIRA_DOWNSTREAM fill:#d63031,color:#fff
    style MODE_DECISION fill:#fdcb6e,color:#333
    style GONOGO fill:#fdcb6e,color:#333
    style CANCELADO fill:#d63031,color:#fff
    style CONTINUA_WATERFALL fill:#00b894,color:#fff
    style EST_PATH fill:#808080,stroke:#636e72
```

---

## 2. Fase 0 — Bootstrap Inteligente (Detalhado)

```mermaid
flowchart TD
    F0_START(["Fase 0: Bootstrap"]) --> F0_1

    subgraph PASSO_01["Passo 0.1 — Coletar Inputs"]
        F0_1["Solicitar PROJECT_PATH, PROJECT_ID_NAME"] --> F0_1_OPT{Solicitar<br/>WATERFALL_ESTIMATION_MODE?}
        F0_1_OPT -->|Sim| F0_1_MODE["Coletar modo:<br/>upstream-discovery ou<br/>downstream-refinement"]
        F0_1_OPT -->|Não| F0_2
        F0_1_MODE --> F0_2
    end

    subgraph PASSO_02["Passo 0.2 — Auditar Docs WATERFALL"]
        F0_2["Verificar existência e status<br/>COMPLIANCE de:<br/>1️⃣ 01-Charter, 02-BRD, 05-SAD, 06-HLD<br/>2️⃣ 03-SRS, 04-RTM, 07-LLD, 11-EAP/WBS<br/>3️⃣ waterfall-estimation/"] --> F0_2_DEC{Quais<br/>pré-condições<br/>atendidas?}

        F0_2_DEC -->|"01,02,05,06 ✅<br/>(todas COMPLIANCE)"| F0_3A["Exibir opção:<br/>Modo UPSTREAM/DISCOVERY<br/>ROM ±50%"]
        F0_2_DEC -->|"03,04,07,11 ✅<br/>(todas COMPLIANCE)"| F0_3B["Exibir opção:<br/>Modo DOWNSTREAM/REFINEMENT<br/>PERT ±15-25%"]
        F0_2_DEC -->|"Ambos os conjuntos ✅"| F0_3C["Exibir opções:<br/>1️⃣ UPSTREAM/DISCOVERY<br/>2️⃣ DOWNSTREAM/REFINEMENT<br/>3️⃣ Ambos em sequência"]
        F0_2_DEC -->|"Nenhum ❌"| F0_3D["Informar docs faltantes<br/>e encerrar"]
    end

    subgraph PASSO_03["Passo 0.3 — Decisão Humana"]
        F0_3A --> F0_3_HUMAN{Humano escolhe<br/>o modo}
        F0_3B --> F0_3_HUMAN
        F0_3C --> F0_3_HUMAN
    end

    subgraph PASSO_04["Passo 0.4 — Coletar Info Complementar"]
        F0_3_HUMAN --> F0_4["Coletar PROJECT-STACK,<br/>PROJECT-TEAM-SKILLS-MAP,<br/>PROJECT-TEAM-CAPACITY<br/>Validar contra baseline corporativa"]
    end

    subgraph PASSO_05["Passo 0.5 — Criar Estrutura"]
        F0_4 --> F0_5["mkdir -p<br/>waterfall-estimation/"]
    end

    subgraph PASSO_06["Passo 0.6 — Auditar Artefatos Existentes"]
        F0_5 --> F0_6["Verificar artefatos de<br/>estimativa já gerados em<br/>waterfall-estimation/"]
    end

    subgraph PASSO_07["Passo 0.7 — Resumo"]
        F0_6 --> F0_7["Exibir resumo:<br/>modo, pré-condições, artefatos existentes,<br/>próxima fase a executar"]
    end

    F0_7 --> ORCH

    style PASSO_01 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_02 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_03 fill:#dfe6e9,stroke:#fdcb6e
    style PASSO_04 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_05 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_06 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_07 fill:#dfe6e9,stroke:#6c5ce7
    style F0_3C fill:#ffeaa7,stroke:#fdcb6e,color:#333
    style F0_3D fill:#ffcccc,stroke:#d63031,color:#333
```

---

## 3. Modo UPSTREAM/DISCOVERY — ROM ±50% (Fases 1-3)

```mermaid
flowchart TD
    GATE_HLD["06-HLD COMPLIANCE ✅"] --> F1_START

    subgraph F1_ROM["Fase 1: UPSTREAM-ROM — ROM ±50%"]
        direction TB
        F1_GEN["GENERATE<br/>Bottom-Up por componente HLD<br/>5 dimensões DTA<br/>ROM = Provável × 1±0.50"] --> F1_GATE["GATE<br/>Valida QA≥25%, Arch≥5%<br/>Verifica fórmula ROM<br/>Sanidade numérica"]
        F1_GATE -->|FAIL| F1_FIX["FIX<br/>Corrige dimensões<br/>Recalcula DTA"]
        F1_FIX -.->|loop| F1_GATE
        F1_GATE -->|PASS| F1_OK["COMPLIANCE ✅"]
    end

    F1_START --> F1_GEN

    subgraph F2_SCOPE["Fase 2: SCOPE-SNAPSHOT-UPSTREAM"]
        direction TB
        F2_GEN["GENERATE<br/>Congela escopo estimado<br/>Lista componentes + exclusões<br/>Matriz rastreabilidade × docs fonte"] --> F2_GATE["GATE<br/>Verifica integralidade<br/>Todos componentes do HLD listados<br/>Exclusões com justificativa"]
        F2_GATE -->|FAIL| F2_FIX["FIX<br/>Adiciona componentes faltantes<br/>Completa rastreabilidade"]
        F2_FIX -.->|loop| F2_GATE
        F2_GATE -->|PASS| F2_OK["COMPLIANCE ✅"]
    end

    F1_OK --> F2_GEN

    subgraph F3_GOV["Fase 3: GOVERNANCE-ROM-REPORT"]
        direction TB
        F3_GEN["GENERATE<br/>Sumário executivo ≤1 pág<br/>Estimativa financeira ROM<br/>Timeline macro<br/>Recomendação GO/NO-GO/HOLD"] --> F3_GATE["GATE<br/>Valida consistência com ROM<br/>Recomendação explícita<br/>Campos para decisão do comitê"]
        F3_GATE -->|FAIL| F3_FIX["FIX<br/>Corrige valores<br/>Clarifica recomendação"]
        F3_FIX -.->|loop| F3_GATE
        F3_GATE -->|PASS| F3_OK["COMPLIANCE ✅"]
    end

    F2_OK --> F3_GEN

    F3_OK --> BARREIRA_UP{{⛔ Barreira UPSTREAM}}

    BARREIRA_UP --> COMITE{Comitê de<br/>Governança}

    COMITE -->|GO ✅| GO_OUT["Prossegue WATERFALL<br/>Doc #07 LLD"]
    COMITE -->|NO-GO ❌| NOGO_OUT["Projeto Cancelado"]
    COMITE -->|HOLD ⏸️| HOLD_OUT["Aguardar info adicional"]

    F1_ROM --> INPUTS1["📥 Inputs:<br/>01-Charter + 02-BRD<br/>05-SAD + 06-HLD<br/>PROJECT-STACK<br/>PROJECT-TEAM-CAPACITY"]
    F2_SCOPE --> INPUTS2["📥 Inputs:<br/>F1 UPSTREAM-ROM<br/>Docs WATERFALL fonte"]
    F3_GOV --> INPUTS3["📥 Inputs:<br/>F1 ROM + F2 Snapshot<br/>01-Project Charter"]

    style F1_GEN fill:#0984e3,color:#fff
    style F1_GATE fill:#fdcb6e,color:#333
    style F1_FIX fill:#e17055,color:#fff
    style F2_GEN fill:#0984e3,color:#fff
    style F3_GEN fill:#0984e3,color:#fff
    style F1_OK fill:#00b894,color:#fff
    style F2_OK fill:#00b894,color:#fff
    style F3_OK fill:#00b894,color:#fff
    style BARREIRA_UP fill:#d63031,color:#fff
    style COMITE fill:#fdcb6e,color:#333
    style GO_OUT fill:#00b894,color:#fff
    style NOGO_OUT fill:#d63031,color:#fff
    style HOLD_OUT fill:#fdcb6e,color:#333
```

---

## 4. Modo DOWNSTREAM/REFINEMENT — PERT ±15-25% (Fases 4-7)

```mermaid
flowchart TD
    GATE_LLD_EAP["07-LLD + 11-EAP/WBS COMPLIANCE ✅"] --> F4_START

    subgraph F4_PERT["Fase 4: DOWNSTREAM-PERT"]
        direction TB
        F4_GEN["GENERATE<br/>Three-Point por pacote EAP<br/>E = O+4M+P/6 · σ = P−O/6<br/>σ_consolidado = √Σσ²<br/>Estimativa independente do ROM"] --> F4_GATE["GATE<br/>Valida O≤M≤P em todos<br/>QA≥25%, Arch≥5%<br/>Precisão ±15-25%<br/>Declaração independência"]
        F4_GATE -->|FAIL| F4_FIX["FIX<br/>Corrige Three-Point<br/>Recalcula DTA<br/>Refina pacotes com σ alto"]
        F4_FIX -.->|loop| F4_GATE
        F4_GATE -->|PASS| F4_OK["COMPLIANCE ✅"]
    end

    F4_START --> F4_GEN

    subgraph F5_SCOPE["Fase 5: SCOPE-SNAPSHOT-DOWNSTREAM"]
        direction TB
        F5_GEN["GENERATE<br/>Congela escopo detalhado<br/>Pacotes EAP + exclusões<br/>Rastreabilidade SRS×RTM×LLD×EAP<br/>Declaração independência"] --> F5_GATE["GATE<br/>Todos pacotes EAP listados<br/>Rastreabilidade completa<br/>Independência declarada"]
        F5_GATE -->|FAIL| F5_FIX["FIX<br/>Adiciona pacotes faltantes<br/>Completa rastreabilidade"]
        F5_FIX -.->|loop| F5_GATE
        F5_GATE -->|PASS| F5_OK["COMPLIANCE ✅"]
    end

    F4_OK --> F5_GEN

    subgraph F6_CRONO["Fase 6: CRONOGRAMA-CALCULADO"]
        direction TB
        F6_GEN["GENERATE<br/>dias = E_PERT / equipe×6h<br/>Caminho crítico<br/>Diagrama Gantt<br/>Marcos × Charter"] --> F6_GATE["GATE<br/>Durações derivadas do PERT<br/>Caminho crítico identificado<br/>Sem conflitos dependência<br/>Alocação ≤ capacidade"]
        F6_GATE -->|FAIL| F6_FIX["FIX<br/>Corrige durações<br/>Ajusta dependências"]
        F6_FIX -.->|loop| F6_GATE
        F6_GATE -->|PASS| F6_OK["COMPLIANCE ✅"]
    end

    F5_OK --> F6_GEN

    subgraph F7_ORC["Fase 7: ORCAMENTO-CALCULADO"]
        direction TB
        F7_GEN["GENERATE<br/>Custo RH = Horas_PERT × Taxa<br/>Curva S acumulada<br/>Contingência = fσ<br/>Fluxo de caixa"] --> F7_GATE["GATE<br/>Custos derivados do PERT<br/>Contingência baseada em σ<br/>Contingência ≤ 50% direto<br/>Somas consistentes"]
        F7_GATE -->|FAIL| F7_FIX["FIX<br/>Corrige custos<br/>Recalcula contingência"]
        F7_FIX -.->|loop| F7_GATE
        F7_GATE -->|PASS| F7_OK["COMPLIANCE ✅"]
    end

    F6_OK --> F7_GEN

    F7_OK --> BARREIRA_DOWN{{⛔ Barreira DOWNSTREAM}}

    BARREIRA_DOWN --> INTEGRACAO["Integração com WATERFALL"]

    INTEGRACAO --> DOC12_OUT["📄 Alimenta Doc #12<br/>Cronograma/Gantt<br/>via UPSTREAM_DOCS"]
    INTEGRACAO --> DOC13_OUT["📄 Alimenta Doc #13<br/>Orçamento<br/>via UPSTREAM_DOCS"]

    F4_PERT --> INPUTS4["📥 Inputs:<br/>03-SRS + 04-RTM<br/>07-LLD + 11-EAP/WBS<br/>PROJECT-STACK<br/>TEAM-SKILLS + CAPACITY"]
    F5_SCOPE --> INPUTS5["📥 Inputs:<br/>F4 DOWNSTREAM-PERT<br/>Docs WATERFALL fonte"]
    F6_CRONO --> INPUTS6["📥 Inputs:<br/>F4 PERT + F5 Snapshot<br/>01-Charter milestones<br/>TEAM-CAPACITY"]
    F7_ORC --> INPUTS7["📥 Inputs:<br/>F4 PERT + F6 Cronograma<br/>PROJECT-STACK<br/>TEAM-CAPACITY"]

    style F4_GEN fill:#00b894,color:#fff
    style F4_GATE fill:#fdcb6e,color:#333
    style F4_FIX fill:#e17055,color:#fff
    style F5_GEN fill:#00b894,color:#fff
    style F6_GEN fill:#00b894,color:#fff
    style F7_GEN fill:#00b894,color:#fff
    style F4_OK fill:#00b894,color:#fff
    style F5_OK fill:#00b894,color:#fff
    style F6_OK fill:#00b894,color:#fff
    style F7_OK fill:#00b894,color:#fff
    style BARREIRA_DOWN fill:#d63031,color:#fff
    style DOC12_OUT fill:#0984e3,color:#fff
    style DOC13_OUT fill:#0984e3,color:#fff
```

---

## 5. Mecanismo de Orquestração Dinâmica (Loop Trifásico)

Todas as fases (1-7) executam este mesmo loop de validação:

```mermaid
flowchart TD
    ORCH(["Orquestrador: Iniciar Fase N"]) --> HITL_CHECK

    HITL_CHECK["Checkpoint HITL:<br/>Deseja fornecer novas<br/>informações antes de gerar?"] --> GEN

    subgraph LOOP["Loop de Validação Soberana — Fases 1 a 7"]
        GEN["1. GERAÇÃO<br/>Executar PROMPT-GENERATE-WATERFALL-ESTIMATION-{ARTIFACT}.md<br/>Parâmetros: ARTIFACT_PATH, PROJECT_ID_NAME,<br/>UPSTREAM_DOCS, INTERNAL_UPSTREAM, SKILLS"] --> GATE

        GATE["2. AUDITORIA INTERNA DA IA<br/>Executar PROMPT-GATE-WATERFALL-ESTIMATION-{ARTIFACT}.md<br/>Status → Em revisão"] --> GATE_RESULT{Resultado<br/>da Auditoria?}

        GATE_RESULT -->|NÃO COMPLIANCE<br/>Erros encontrados| FIX["2b. CORREÇÃO CIRÚRGICA<br/>Executar PROMPT-FIX-WATERFALL-ESTIMATION-{ARTIFACT}.md<br/>Apenas nas seções com VIOLATIONS[]"]
        FIX --> GATE

        GATE_RESULT -->|SEM ERROS| HUMAN_GATE

        HUMAN_GATE["3. PORTÃO DE VALIDAÇÃO HUMANA<br/>Status: PRÉ-COMPLIANCE INTERNO<br/>— AGUARDANDO VALIDAÇÃO HUMANA"] --> P1["P1: Artefato aderente<br/>às necessidades?"]
        P1 --> P2["P2: Novos documentos<br/>de entrada?"]
        P2 --> P3["P3: Novos inputs ou<br/>mudanças de escopo?"]
        P3 --> P4["P4 HITL: Novas informações<br/>antes de prosseguir?"]
        P4 --> HUMAN_DEC{Decisão<br/>do Humano?}

        HUMAN_DEC -->|Aprova<br/>SIM/NÃO/NÃO/NÃO| COMPLIANCE
        HUMAN_DEC -->|Fornece novos inputs| GEN
    end

    COMPLIANCE(["✅ STATUS: COMPLIANCE<br/>Artefato congelado<br/>Próximo artefato destravado"])

    style GEN fill:#0984e3,color:#fff
    style GATE fill:#fdcb6e,color:#333
    style FIX fill:#e17055,color:#fff
    style HUMAN_GATE fill:#6c5ce7,color:#fff
    style HITL_CHECK fill:#6c5ce7,color:#fff
    style P1 fill:#dfe6e9,stroke:#6c5ce7,color:black
    style P2 fill:#dfe6e9,stroke:#6c5ce7,color:black
    style P3 fill:#dfe6e9,stroke:#6c5ce7,color:black
    style P4 fill:#dfe6e9,stroke:#6c5ce7,color:black
    style COMPLIANCE fill:#00b894,color:#fff
    style LOOP fill:#fff3e0,stroke:#f39c12
```

---

## 6. Integração com o Roadmap WATERFALL

```mermaid
flowchart LR
    subgraph WATERFALL_DOCS["WATERFALL Docs — Fase 3: Design e Arquitetura"]
        direction TB
        W5["05-SAD ✅"] --> W6["06-HLD ✅"]
    end

    subgraph EST_UPSTREAM["WATERFALL-ESTIMATION — UPSTREAM/DISCOVERY"]
        direction TB
        EU1["F1: ROM ±50%"] --> EU2["F2: Scope Snapshot"]
        EU2 --> EU3["F3: Governance GO/NO-GO"]
    end

    subgraph WATERFALL_DOCS2["WATERFALL Docs — Fase 3 + Fase 5"]
        direction TB
        W7["07-LLD ✅"] --> W11["11-EAP/WBS ✅"]
    end

    subgraph EST_DOWNSTREAM["WATERFALL-ESTIMATION — DOWNSTREAM/REFINEMENT"]
        direction TB
        ED4["F4: PERT ±15-25%"] --> ED5["F5: Scope Snapshot"]
        ED5 --> ED6["F6: Cronograma Calculado"]
        ED6 --> ED7["F7: Orçamento Calculado"]
    end

    subgraph WATERFALL_DOCS3["WATERFALL Docs — Fase 5: Planejamento"]
        direction TB
        W12["12-Cronograma/Gantt"]
        W13["13-Orçamento"]
    end

    W6 -->|"Gatilho UPSTREAM<br/>após HLD COMPLIANCE"| EU1
    EU3 -->|"GO ✅"| W7
    EU3 -->|"NO-GO ❌"| CANCEL["Projeto Cancelado"]

    W11 -->|"Gatilho DOWNSTREAM<br/>após LLD + EAP COMPLIANCE"| ED4
    ED6 -->|"UPSTREAM_DOC adicional"| W12
    ED7 -->|"UPSTREAM_DOC adicional"| W13

    style WATERFALL_DOCS fill:#e3f2fd,stroke:#0984e3
    style EST_UPSTREAM fill:#fff3e0,stroke:#e65100
    style WATERFALL_DOCS2 fill:#e3f2fd,stroke:#0984e3
    style EST_DOWNSTREAM fill:#e8f5e9,stroke:#2e7d32
    style WATERFALL_DOCS3 fill:#e3f2fd,stroke:#0984e3
    style CANCEL fill:#ffcccc,stroke:#d63031
```

---

## 7. Diagrama de Estados — Visão Unificada

```mermaid
stateDiagram-v2
    [*] --> Bootstrap: Início da Estimativa WATERFALL

    state Bootstrap {
        [*] --> ColetarInputs
        ColetarInputs --> AuditarDocsWATERFALL
        AuditarDocsWATERFALL --> DetectarModo: upstream-discovery | downstream-refinement | ambos
        DetectarModo --> EscolherModo
        EscolherModo --> ColetarStackETime
        ColetarStackETime --> CriarEstrutura
        CriarEstrutura --> ResumoInicial
    }

    Bootstrap --> ModoUpstream: upstream-discovery
    Bootstrap --> ModoDownstream: downstream-refinement

    state ModoUpstream {
        [*] --> F1_ROM
        state F1_ROM {
            [*] --> F1Gen
            F1Gen --> F1Gate
            F1Gate --> F1Fix: FAIL
            F1Fix --> F1Gate
            F1Gate --> F1Human: PASS
            F1Human --> F1Gen: Novos inputs
            F1Human --> F1Done: Aprovado
        }
        F1_ROM --> F2_SnapUp: COMPLIANCE
        state F2_SnapUp {
            [*] --> F2Gen
            F2Gen --> F2Gate
            F2Gate --> F2Fix: FAIL
            F2Fix --> F2Gate
            F2Gate --> F2Human: PASS
            F2Human --> F2Gen: Novos inputs
            F2Human --> F2Done: Aprovado
        }
        F2_SnapUp --> F3_Gov: COMPLIANCE
        state F3_Gov {
            [*] --> F3Gen
            F3Gen --> F3Gate
            F3Gate --> F3Fix: FAIL
            F3Fix --> F3Gate
            F3Gate --> F3Human: PASS
            F3Human --> F3Gen: Novos inputs
            F3Human --> F3Done: Aprovado
        }
        F3_Gov --> BarreiraUpstream: COMPLIANCE
    }

    state BarreiraUpstream {
        [*] --> Governanca
        Governanca --> GoAhead: GO ✅
        Governanca --> Cancelado: NO-GO ❌
        Governanca --> Aguardando: HOLD ⏸️
    }

    state ModoDownstream {
        [*] --> F4_PERT
        state F4_PERT {
            [*] --> F4Gen
            F4Gen --> F4Gate
            F4Gate --> F4Fix: FAIL
            F4Fix --> F4Gate
            F4Gate --> F4Human: PASS
            F4Human --> F4Gen: Novos inputs
            F4Human --> F4Done: Aprovado
        }
        F4_PERT --> F5_SnapDown: COMPLIANCE
        state F5_SnapDown {
            [*] --> F5Gen
            F5Gen --> F5Gate
            F5Gate --> F5Fix: FAIL
            F5Fix --> F5Gate
            F5Gate --> F5Human: PASS
            F5Human --> F5Gen: Novos inputs
            F5Human --> F5Done: Aprovado
        }
        F5_SnapDown --> F6_Crono: COMPLIANCE
        state F6_Crono {
            [*] --> F6Gen
            F6Gen --> F6Gate
            F6Gate --> F6Fix: FAIL
            F6Fix --> F6Gate
            F6Gate --> F6Human: PASS
            F6Human --> F6Gen: Novos inputs
            F6Human --> F6Done: Aprovado
        }
        F6_Crono --> F7_Orc: COMPLIANCE
        state F7_Orc {
            [*] --> F7Gen
            F7Gen --> F7Gate
            F7Gate --> F7Fix: FAIL
            F7Fix --> F7Gate
            F7Gate --> F7Human: PASS
            F7Human --> F7Gen: Novos inputs
            F7Human --> F7Done: Aprovado
        }
        F7_Orc --> BarreiraDownstream: COMPLIANCE
    }

    state BarreiraDownstream {
        [*] --> IntegrarWATERFALL
        IntegrarWATERFALL --> [*]: Cronograma → Doc #12 | Orçamento → Doc #13
    }

    GoAhead --> [*]
    Cancelado --> [*]
    Aguardando --> Governanca: Info recebida
```

---

## 8. Efeitos Cascata — Dependências entre Artefatos

```mermaid
flowchart TD
    subgraph UPSTREAM_CASCADE["Cascata UPSTREAM/DISCOVERY"]
        F1["F1: UPSTREAM-ROM"] -->|"modificação em F1<br/>impacta"| F2_C["F2: Scope Snapshot"]
        F2_C -->|"modificação em F2<br/>impacta"| F3_C["F3: Governance Report"]
    end

    subgraph DOWNSTREAM_CASCADE["Cascata DOWNSTREAM/REFINEMENT"]
        F4["F4: DOWNSTREAM-PERT"] -->|"modificação em F4<br/>impacta"| F5_C["F5: Scope Snapshot"]
        F5_C -->|"modificação em F5<br/>impacta"| F6_C["F6: Cronograma"]
        F4 -->|"modificação em F4<br/>impacta"| F6_C
        F6_C -->|"modificação em F6<br/>impacta"| F7_C["F7: Orçamento"]
        F5_C -->|"modificação em F5<br/>impacta"| F7_C
    end

    subgraph CROSS_CASCADE["Cascata Cruzada WATERFALL → Estimativa"]
        HLD["06-HLD modificado"] -->|"dispara reexecução"| F1
        LLD["07-LLD modificado"] -->|"dispara reexecução"| F4
        EAP["11-EAP/WBS modificado"] -->|"dispara reexecução"| F4
    end

    style UPSTREAM_CASCADE fill:#fff3e0,stroke:#e65100
    style DOWNSTREAM_CASCADE fill:#e8f5e9,stroke:#2e7d32
    style CROSS_CASCADE fill:#ffcccc,stroke:#d63031
```

---

## 9. Estrutura de Diretórios Gerada

```mermaid
flowchart LR
    subgraph PROJECT["PROJECT_ID_NAME/"]
        DOCS["📄 01-20 docs WATERFALL ..."]
        EST_DIR["📁 waterfall-estimation/"]
    end

    EST_DIR --> F1_FILE["WATERFALL-ESTIMATION-UPSTREAM-ROM.md"]
    EST_DIR --> F2_FILE["ESTIMATION-SCOPE-SNAPSHOT-UPSTREAM.md"]
    EST_DIR --> F3_FILE["GOVERNANCE-ROM-REPORT.md"]
    EST_DIR --> F4_FILE["WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md"]
    EST_DIR --> F5_FILE["ESTIMATION-SCOPE-SNAPSHOT-DOWNSTREAM.md"]
    EST_DIR --> F6_FILE["CRONOGRAMA-CALCULADO.md"]
    EST_DIR --> F7_FILE["ORCAMENTO-CALCULADO.md"]

    style PROJECT fill:#e3f2fd,stroke:#0984e3
    style EST_DIR fill:#808080,stroke:#636e72,color:#fff
    style F1_FILE fill:#fff3e0,stroke:#e65100,color:#333
    style F2_FILE fill:#fff3e0,stroke:#e65100,color:#333
    style F3_FILE fill:#fff3e0,stroke:#e65100,color:#333
    style F4_FILE fill:#e8f5e9,stroke:#2e7d32,color:#333
    style F5_FILE fill:#e8f5e9,stroke:#2e7d32,color:#333
    style F6_FILE fill:#e8f5e9,stroke:#2e7d32,color:#333
    style F7_FILE fill:#e8f5e9,stroke:#2e7d32,color:#333
```

---

## 10. Tabela de Símbolos e Convenções

| Símbolo/Cor | Significado |
|-------------|-------------|
| 🟣 Roxo (`#6c5ce7`) | Bootstrap / Validação Humana / HITL |
| 🔵 Azul (`#0984e3`) | Modo UPSTREAM/DISCOVERY (ROM) / Docs WATERFALL |
| 🟢 Verde (`#00b894`) | Modo DOWNSTREAM/REFINEMENT (PERT) / Compliance |
| 🟠 Laranja (`#e17055`) | Correções (FIX) |
| 🟡 Amarelo (`#fdcb6e`) | Gate / Decisões |
| 🔴 Vermelho (`#d63031`) | Barreiras / NO-GO / Cancelado |
| ⬜ Cinza (`#808080`) | Pastas de trabalho |
| 🟤 Laranja escuro (`#e65100`) | Artefatos UPSTREAM/DISCOVERY |
| 🟢 Verde escuro (`#2e7d32`) | Artefatos DOWNSTREAM/REFINEMENT |
| 🔲 Linha tracejada | Loop de retrabalho ou associação |
| 🔲 Linha sólida | Fluxo sequencial normal |
| 📁 | Pasta de trabalho |
| 📄 | Documento |
| 📥 | Inputs consumidos |
| ⛔ | Barreira de bloqueio |
| 1️⃣ 2️⃣ 3️⃣ | Passos do Bootstrap |
| ✅ | COMPLIANCE / Aprovado |
| ❌ | Cancelado / Rejeitado |
| ⏸️ | HOLD / Aguardando |

---

## 11. Métricas e Fórmulas Chave

| Métrica | Fórmula | Aplicação |
|---------|---------|-----------|
| **ROM** | `ROM = Provável × (1 ± 0.50)` | UPSTREAM F1 — por componente HLD |
| **PERT (E)** | `E = (O + 4M + P) / 6` | DOWNSTREAM F4 — por pacote EAP |
| **Desvio Padrão (σ)** | `σ = (P − O) / 6` | DOWNSTREAM F4 — incerteza por pacote |
| **σ Consolidado** | `σ_total = √(Σ σ²)` | DOWNSTREAM F4 — incerteza global |
| **Precisão PERT** | `Precisão = σ_total / E_total` | DOWNSTREAM F4 — esperado ≤ 25% |
| **DTA — QA Ratio** | `QA_Ratio = Σ horas_qa / Σ horas_dev` | UPSTREAM F1 + DOWNSTREAM F4 — ≥ 25% |
| **DTA — Arch Ratio** | `Arch_Ratio = Σ horas_arch / Σ total_horas` | UPSTREAM F1 + DOWNSTREAM F4 — ≥ 5% |
| **Duração (dias)** | `dias = E_PERT / (equipe × 6h)` | DOWNSTREAM F6 — conversão horas→cronograma |
| **Contingência** | `Contingência = f(σ, nível_confiança)` | DOWNSTREAM F7 — reserva financeira |

---

> **📁 Arquivos relacionados:**
> - `PROMPT-ROADMAP-GENERATE-WATERFALL-ESTIMATION.md` — Documento fonte (v1.0)
> - `../PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md` — Roadmap WATERFALL Docs (integração Docs #12 e #13)
> - `../PROMPT-ROADMAP-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY.md` — Roadmap Upstream Agile (referência metodológica)
> - `../PROMPT-ROADMAP-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT.md` — Roadmap Downstream Agile (referência metodológica)
> - `../PROMPT-ROADMAP-GENERATE-SOURCING-FACTORY-BIDDING.md` — Roadmap Sourcing (consome baseline interna WATERFALL-ESTIMATION)
> - `../../standards/DTA-Engine-de-Bidding-e-Estimativas.md` — Engine de validação DTA
> - `../../standards/DTA-VALIDATION-STANDARDS.md` — Regras DTA, ROM, PERT e PIB
> - `PROMPT-GENERATE-WATERFALL-ESTIMATION-*.md` — 7 prompts geradores (F1-F7)
> - `PROMPT-GATE-WATERFALL-ESTIMATION-*.md` — 7 prompts de auditoria
> - `PROMPT-FIX-WATERFALL-ESTIMATION-*.md` — 7 prompts de correção
