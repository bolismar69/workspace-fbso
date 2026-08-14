# FLOWCHART: ROADMAP DE DOCUMENTOS WATERFALL

## Versão: 2.0 — Visualização Gráfica das 5 Fases, 22 Documentos, Dupla RTM e Gates Estruturais

> ⚠️ **VISUALIZAÇÃO LEGADA (histórica):** este diagrama reflete a versão original do roadmap (5 fases, 22 documentos). A estrutura vigente é a do roadmap master — **6 fases, 38 documentos** (`PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md`) e do `flowchart-WATERFALL.md` atualizado.

> **Documento de referência:** `PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md` v2.0
>
> Este documento complementa o roadmap textual com diagramas Mermaid que visualizam o fluxo de execução das **5 fases WATERFALL**, os **22 documentos**, o mecanismo de orquestração Generate→Gate→Fix, a dupla RTM (Negócio + Sistema) e a integração com WATERFALL-ESTIMATION.

---

## 1. Visão Macro — 5 Fases WATERFALL (22 Documentos)

```mermaid
flowchart TB
    START(["🚀 Início"]) --> F0["Fase 0: Bootstrap"]

    F0 --> FASE1

    subgraph FASE1["Fase 1: INICIAÇÃO E REQUISITOS DE NEGÓCIO"]
        direction LR
        D001["001: PROJECT-CHARTER"] --> D002["002: STAKEHOLDER-MAP"]
        D002 --> D005["005: BRD<br/>REQ-01, REQ-02..."]
        D005 --> D010["010: FRD<br/>FEAT-01, RN-01, UC-01"]
        D010 --> D015["015: RTM-FASE-1<br/>Rastr. Negócio"]
    end

    FASE1 --> FASE2

    subgraph FASE2["Fase 2: ESPECIFICAÇÃO DE SISTEMA E ARQUITETURA MACRO"]
        direction LR
        D020["020: SRS<br/>FR-01, NFRs"] --> D025["025: RTM-FASE-2<br/>Rastr. Sistema"]
        D025 --> D030["030: SAD"]
        D030 --> D035["035: HLD"]
    end

    FASE2 --> EST_GATE_UP{{🎯 GATE 1: UPSTREAM<br/>após 035-HLD}}
    EST_GATE_UP -->|"Opcional"| EST_UP["WATERFALL-ESTIMATION<br/>UPSTREAM/DISCOVERY<br/>ROM ±50% + GO/NO-GO"]
    EST_GATE_UP -->|"Pular"| FASE3
    EST_UP -->|"GO ✅"| FASE3
    EST_UP -->|"NO-GO ❌"| CANCEL["Projeto Cancelado"]

    FASE2 --> FASE3

    subgraph FASE3["Fase 3: ENGENHARIA DETALHADA E QUALIDADE"]
        direction LR
        D040["040: LLD"] --> D045["045: EST-PLAN"]
        D045 --> D050["050: EST-CASES"]
        D050 --> D055["055: RELATORIO-QUALIDADE"]
        D055 --> D060["060: EAP-WBS"]
    end

    FASE3 --> EST_GATE_DOWN{{🎯 GATE 2: DOWNSTREAM<br/>após 040-LLD + 060-EAP}}
    EST_GATE_DOWN -->|"Opcional"| EST_DOWN["WATERFALL-ESTIMATION<br/>DOWNSTREAM/REFINEMENT<br/>PERT ±15-25%<br/>→ Alimenta 065 + 070"]
    EST_GATE_DOWN -->|"Pular"| FASE4
    EST_DOWN --> FASE4

    FASE3 --> FASE4

    subgraph FASE4["Fase 4: PLANEJAMENTO E BASELINE"]
        direction LR
        D065["065: CRONOGRAMA-GANTT"] --> D070["070: ORCAMENTO"]
        D070 --> D075["075: PLANO-COMUNICACAO"]
        D075 --> D080["080: PLANO-RISCOS"]
        D080 --> D085["085: GESTAO-MUDANCAS 🆕"]
        D085 --> D090["090: STRATEGIC-IMPLEMENTATION-AND-DEPLOYMENT-PLAN"]
    end

    FASE4 --> FASE5

    subgraph FASE5["Fase 5: ENCERRAMENTO E OPERAÇÃO"]
        direction LR
        D095["095: MANUAIS-USUARIO"] --> D100["100: MANUAIS-OPERACIONAIS"]
        D100 --> D105["105: TERMO-ACEITE"]
        D105 --> D110["110: LICOES-APRENDIDAS"]
        D110 --> D115["115: TERMO-ENCERRAMENTO 🆕"]
    end

    FASE5 --> END(["✅ Roadmap Concluído"])

    style F0 fill:#6c5ce7,color:#fff
    style D001 fill:#0984e3,color:#fff
    style D002 fill:#0984e3,color:#fff
    style D005 fill:#0984e3,color:#fff
    style D010 fill:#0984e3,color:#fff
    style D015 fill:#0984e3,color:#fff
    style D020 fill:#6c5ce7,color:#fff
    style D025 fill:#6c5ce7,color:#fff
    style D030 fill:#6c5ce7,color:#fff
    style D035 fill:#6c5ce7,color:#fff
    style D040 fill:#00b894,color:#fff
    style D045 fill:#00b894,color:#fff
    style D050 fill:#00b894,color:#fff
    style D055 fill:#00b894,color:#fff
    style D060 fill:#00b894,color:#fff
    style D065 fill:#e17055,color:#fff
    style D070 fill:#e17055,color:#fff
    style D075 fill:#e17055,color:#fff
    style D080 fill:#e17055,color:#fff
    style D085 fill:#e17055,color:#fff
    style D090 fill:#e17055,color:#fff
    style D095 fill:#a29bfe,color:#fff
    style D100 fill:#a29bfe,color:#fff
    style D105 fill:#a29bfe,color:#fff
    style D110 fill:#a29bfe,color:#fff
    style D115 fill:#a29bfe,color:#fff
    style EST_GATE_UP fill:#fdcb6e,color:#333
    style EST_GATE_DOWN fill:#fdcb6e,color:#333
    style EST_UP fill:#fff3e0,stroke:#e65100
    style EST_DOWN fill:#e8f5e9,stroke:#2e7d32
    style CANCEL fill:#d63031,color:#fff
```

---

## 2. Fase 0 — Bootstrap Inteligente (Detalhado)

```mermaid
flowchart TD
    F0_START(["Fase 0: Bootstrap"]) --> F0_1

    subgraph PASSO_01["Passo 0.1 — Solicitar Inputs"]
        F0_1["Apresentar 7 inputs obrigatórios<br/>+ 4 opcionais<br/>+ diretiva HITL"] --> F0_1_DEC{Inputs<br/>fornecidos?}
        F0_1_DEC -->|Parcial| F0_1_ASK["Perguntar inputs faltantes<br/>de forma clara e objetiva"]
        F0_1_ASK --> F0_1_DEC
        F0_1_DEC -->|Todos ✅| F0_2
    end

    subgraph PASSO_02["Passo 0.2 — Validar Stack"]
        F0_2["Ler STACK-PADROES-CORPORATIVOS"] --> F0_2_VAL{Stack dentro<br/>do padrão?}
        F0_2_VAL -->|✅ Padrão| F0_3
        F0_2_VAL -->|⚠️ Fora| F0_2_JUST["Solicitar justificativa<br/>técnica ao usuário"]
        F0_2_JUST --> F0_3
    end

    subgraph PASSO_03["Passo 0.3-0.4 — Time e Skills"]
        F0_3["Solicitar PROJECT-TEAM-SKILLS-MAP<br/>Sugerir skills inferidos do contexto"] --> F0_4["Solicitar PROJECT-TEAM-CAPACITY<br/>Seniores, plenos, juniores, duração"]
        F0_4 --> F0_5
    end

    subgraph PASSO_04["Passo 0.5 — Exibir Caminhos"]
        F0_5["Exibir PROJECT_COMPLETE_PATH_NAME<br/>TECHNICAL_SOLUTION_NAMES<br/>ARCHITECTURE_GLOBAL · SECURITY_GLOBAL<br/>Stack · Time"] --> F0_5_CONF{Confirma?}
        F0_5_CONF -->|NÃO| F0_1
        F0_5_CONF -->|SIM| F0_6
    end

    subgraph PASSO_05["Passo 0.6-0.7 — Estrutura e Status"]
        F0_6["mkdir -p PROJECT_COMPLETE_PATH_NAME"] --> F0_7["Verificar existência e status<br/>COMPLIANCE dos 20 documentos"]
        F0_7 --> F0_7_DEC{Status?}
        F0_7_DEC -->|"Todos ❌"| F0_8A["Iniciar Fase 1, Doc #1"]
        F0_7_DEC -->|"Parcial"| F0_8B["Iniciar do primeiro doc<br/>sem COMPLIANCE"]
        F0_7_DEC -->|"Todos ✅"| F0_8C["Perguntar: revisar,<br/>novo ciclo ou encerrar"]
    end

    subgraph PASSO_06["Passo 0.8 — Resumo"]
        F0_8A --> F0_8["Exibir resumo e iniciar"]
        F0_8B --> F0_8
        F0_8C --> F0_8
    end

    style PASSO_01 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_02 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_03 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_04 fill:#dfe6e9,stroke:#fdcb6e
    style PASSO_05 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_06 fill:#dfe6e9,stroke:#6c5ce7
```

---

## 3. Fases 1-3 — Inicialização, Requisitos e Design

```mermaid
flowchart TD
    subgraph F1["Fase 1: INICIALIZAÇÃO"]
        direction TB
        D01_GEN["GENERATE #1<br/>PROJECT-CHARTER<br/>Skills: draft-project-charter, senior-pm"] --> D01_GATE["GATE #1<br/>Valida seções:<br/>Escopo, Cronograma Macro,<br/>Orçamento Macro, Riscos, RACI"]
        D01_GATE -->|FAIL| D01_FIX["FIX #1<br/>Corrige seções<br/>com VIOLATIONS"]
        D01_FIX -.->|loop| D01_GATE
        D01_GATE -->|PASS| D01_HUMAN["Validação Humana<br/>4 perguntas HITL"]
        D01_HUMAN -->|Aprovado| D01_OK["#1: COMPLIANCE ✅"]
        D01_HUMAN -->|Novos inputs| D01_GEN
    end

    D01_OK --> F2_START

    subgraph F2["Fase 2: REQUISITOS"]
        direction TB
        D02_GEN["GENERATE #2: BRD<br/>Skills: brd-creation, business-analyst"] --> D02_GATE["GATE #2"]
        D02_GATE -->|FAIL| D02_FIX["FIX #2"]
        D02_FIX -.->|loop| D02_GATE
        D02_GATE -->|PASS| D02_OK["#2: COMPLIANCE ✅"]

        D02_OK --> D03_GEN["GENERATE #3: SRS<br/>Skills: frs-creation, requirements-engineering"] --> D03_GATE["GATE #3"]
        D03_GATE -->|FAIL| D03_FIX["FIX #3"]
        D03_FIX -.->|loop| D03_GATE
        D03_GATE -->|PASS| D03_OK["#3: COMPLIANCE ✅"]

        D03_OK --> D04_GEN["GENERATE #4: RTM<br/>Skills: requirements-modeling, requirements-validation"] --> D04_GATE["GATE #4"]
        D04_GATE -->|FAIL| D04_FIX["FIX #4"]
        D04_FIX -.->|loop| D04_GATE
        D04_GATE -->|PASS| D04_OK["#4: COMPLIANCE ✅"]
    end

    F2_START --> D02_GEN

    D04_OK --> F3_START

    subgraph F3["Fase 3: DESIGN E ARQUITETURA"]
        direction TB
        D05_GEN["GENERATE #5: SAD<br/>Skills: software-architecture, architecture-designer"] --> D05_GATE["GATE #5"]
        D05_GATE -->|FAIL| D05_FIX["FIX #5"]
        D05_FIX -.->|loop| D05_GATE
        D05_GATE -->|PASS| D05_OK["#5: COMPLIANCE ✅"]

        D05_OK --> D06_GEN["GENERATE #6: HLD<br/>Skills: c4-container, system-design, ADRs"] --> D06_GATE["GATE #6"]
        D06_GATE -->|FAIL| D06_FIX["FIX #6"]
        D06_FIX -.->|loop| D06_GATE
        D06_GATE -->|PASS| D06_OK["#6: COMPLIANCE ✅"]

        D06_OK --> EST_CHECK_UP{{🎯 Após #6 HLD:<br/>Executar WATERFALL-ESTIMATION<br/>UPSTREAM/DISCOVERY?}}

        D06_OK --> D07_GEN["GENERATE #7: LLD<br/>Skills: c4-component, ddd-tactical-patterns"] --> D07_GATE["GATE #7"]
        D07_GATE -->|FAIL| D07_FIX["FIX #7"]
        D07_FIX -.->|loop| D07_GATE
        D07_GATE -->|PASS| D07_OK["#7: COMPLIANCE ✅"]
    end

    F3_START --> D05_GEN

    D07_OK --> F4_START

    style D01_GEN fill:#0984e3,color:#fff
    style D02_GEN fill:#0984e3,color:#fff
    style D03_GEN fill:#0984e3,color:#fff
    style D04_GEN fill:#0984e3,color:#fff
    style D05_GEN fill:#0984e3,color:#fff
    style D06_GEN fill:#0984e3,color:#fff
    style D07_GEN fill:#0984e3,color:#fff
    style D01_OK fill:#00b894,color:#fff
    style D02_OK fill:#00b894,color:#fff
    style D03_OK fill:#00b894,color:#fff
    style D04_OK fill:#00b894,color:#fff
    style D05_OK fill:#00b894,color:#fff
    style D06_OK fill:#00b894,color:#fff
    style D07_OK fill:#00b894,color:#fff
    style EST_CHECK_UP fill:#fdcb6e,color:#333
```

---

## 4. Fases 4-6 — Testes, Planejamento e Implantação

```mermaid
flowchart TD
    F4_START(["Após #7 LLD COMPLIANCE"])

    subgraph F4["Fase 4: TESTES E QUALIDADE"]
        direction TB
        D08_GEN["GENERATE #8: TEST-PLAN<br/>Skills: test-strategy-design, qa-test-planner"] --> D08_GATE["GATE #8"]
        D08_GATE -->|FAIL| D08_FIX["FIX #8"]
        D08_FIX -.->|loop| D08_GATE
        D08_GATE -->|PASS| D08_OK["#8: COMPLIANCE ✅"]

        D08_OK --> D09_GEN["GENERATE #9: TEST-CASES<br/>Skills: test-case-creation, acceptance-criteria"] --> D09_GATE["GATE #9"]
        D09_GATE -->|FAIL| D09_FIX["FIX #9"]
        D09_FIX -.->|loop| D09_GATE
        D09_GATE -->|PASS| D09_OK["#9: COMPLIANCE ✅"]

        D09_OK --> D10_GEN["GENERATE #10: Relatório Qualidade<br/>Skills: quality-documentation-manager, qa"] --> D10_GATE["GATE #10"]
        D10_GATE -->|FAIL| D10_FIX["FIX #10"]
        D10_FIX -.->|loop| D10_GATE
        D10_GATE -->|PASS| D10_OK["#10: COMPLIANCE ✅"]
    end

    F4_START --> D08_GEN

    D10_OK --> F5_START

    subgraph F5["Fase 5: PLANEJAMENTO"]
        direction TB
        D11_GEN["GENERATE #11: EAP/WBS<br/>Skills: decomposition-planning-roadmap"] --> D11_GATE["GATE #11"]
        D11_GATE -->|FAIL| D11_FIX["FIX #11"]
        D11_FIX -.->|loop| D11_GATE
        D11_GATE -->|PASS| D11_OK["#11: COMPLIANCE ✅"]

        D11_OK --> EST_CHECK_DOWN{{🎯 Após #7 LLD + #11 EAP:<br/>Executar WATERFALL-ESTIMATION<br/>DOWNSTREAM/REFINEMENT?}}

        D11_OK --> D12_GEN["GENERATE #12: Cronograma/Gantt<br/>+ CRONOGRAMA-CALCULADO se disponível<br/>Skills: roadmap-planning, project-estimation"] --> D12_GATE["GATE #12"]
        D12_GATE -->|FAIL| D12_FIX["FIX #12"]
        D12_FIX -.->|loop| D12_GATE
        D12_GATE -->|PASS| D12_OK["#12: COMPLIANCE ✅"]

        D12_OK --> D13_GEN["GENERATE #13: Orçamento<br/>+ ORCAMENTO-CALCULADO se disponível<br/>Skills: project-estimation"] --> D13_GATE["GATE #13"]
        D13_GATE -->|FAIL| D13_FIX["FIX #13"]
        D13_FIX -.->|loop| D13_GATE
        D13_GATE -->|PASS| D13_OK["#13: COMPLIANCE ✅"]

        D13_OK --> D14_GEN["GENERATE #14: Plano Comunicação<br/>Skills: stakeholder-analysis"] --> D14_GATE["GATE #14"]
        D14_GATE -->|PASS| D14_OK["#14: COMPLIANCE ✅"]

        D13_OK --> D15_GEN["GENERATE #15: Plano Riscos<br/>Skills: risk-manager"] --> D15_GATE["GATE #15"]
        D15_GATE -->|PASS| D15_OK["#15: COMPLIANCE ✅"]
    end

    F5_START --> D11_GEN

    D14_OK --> F6_START
    D15_OK --> F6_START

    subgraph F6["Fase 6: IMPLANTAÇÃO E ENCERRAMENTO"]
        direction TB
        D16_GEN["GENERATE #16: STRATEGIC-IMPLEMENTATION-AND-DEPLOYMENT-PLAN<br/>Skills: deployment-engineer, devops-rollout-plan"] --> D16_GATE["GATE #16"]
        D16_GATE -->|PASS| D16_OK["#16: COMPLIANCE ✅"]

        D16_OK --> D17_GEN["GENERATE #17: Manuais Usuário<br/>Skills: documentation-generation-doc-generate"] --> D17_GATE["GATE #17"]
        D17_GATE -->|PASS| D17_OK["#17: COMPLIANCE ✅"]

        D17_OK --> D18_GEN["GENERATE #18: Manuais Operacionais"] --> D18_GATE["GATE #18"]
        D18_GATE -->|PASS| D18_OK["#18: COMPLIANCE ✅"]

        D18_OK --> D19_GEN["GENERATE #19: Termo Aceite<br/>Skills: contract-and-proposal-writer"] --> D19_GATE["GATE #19"]
        D19_GATE -->|PASS| D19_OK["#19: COMPLIANCE ✅"]

        D19_OK --> D20_GEN["GENERATE #20: Lições Aprendidas<br/>Consolida todos os 19 docs anteriores"] --> D20_GATE["GATE #20"]
        D20_GATE -->|PASS| D20_OK["#20: COMPLIANCE ✅"]
    end

    F6_START --> D16_GEN

    D20_OK --> GIT["Git Workflow:<br/>commit → push → PR → merge"]

    style D08_GEN fill:#00b894,color:#fff
    style D09_GEN fill:#00b894,color:#fff
    style D10_GEN fill:#00b894,color:#fff
    style D11_GEN fill:#e17055,color:#fff
    style D12_GEN fill:#e17055,color:#fff
    style D13_GEN fill:#e17055,color:#fff
    style D14_GEN fill:#e17055,color:#fff
    style D15_GEN fill:#e17055,color:#fff
    style D16_GEN fill:#a29bfe,color:#fff
    style D17_GEN fill:#a29bfe,color:#fff
    style D18_GEN fill:#a29bfe,color:#fff
    style D19_GEN fill:#a29bfe,color:#fff
    style D20_GEN fill:#a29bfe,color:#fff
    style D08_OK fill:#00b894,color:#fff
    style D09_OK fill:#00b894,color:#fff
    style D10_OK fill:#00b894,color:#fff
    style D11_OK fill:#00b894,color:#fff
    style D12_OK fill:#00b894,color:#fff
    style D13_OK fill:#00b894,color:#fff
    style D14_OK fill:#00b894,color:#fff
    style D15_OK fill:#00b894,color:#fff
    style D16_OK fill:#00b894,color:#fff
    style D17_OK fill:#00b894,color:#fff
    style D18_OK fill:#00b894,color:#fff
    style D19_OK fill:#00b894,color:#fff
    style D20_OK fill:#00b894,color:#fff
    style EST_CHECK_DOWN fill:#fdcb6e,color:#333
    style GIT fill:#6c5ce7,color:#fff
```

---

## 5. Mecanismo de Orquestração — Loop Generate→Gate→Fix por Documento

```mermaid
flowchart TD
    ORCH(["Orquestrador: STEP 1<br/>Computar inputs para Doc N"]) --> HITL

    HITL["Checkpoint HITL:<br/>Deseja fornecer novas<br/>informações antes de gerar<br/>Documento N?"] --> GEN

    subgraph LOOP["Loop de Validação Soberana — Para CADA documento"]
        GEN["STEP 2: GENERATE<br/>Invocar project-documents-waterfall/<br/>PROMPT-GENERATE-{DOC-SLUG}.md<br/>Parâmetros: DOC_PATH, PROJECT_ID_NAME,<br/>UPSTREAM_DOCS, SKILLS, + domínio"] --> GATE

        GATE["STEP 3: GATE<br/>Invocar PROMPT-GATE-{DOC-SLUG}.md<br/>Ler DOC_PATH, aplicar CHECKLIST<br/>Status → Em revisão"] --> GATE_RESULT{Resultado<br/>da Auditoria?}

        GATE_RESULT -->|"FAIL<br/>VIOLATIONS[]"| FIX["STEP 4a: FIX CIRÚRGICO<br/>Invocar PROMPT-FIX-{DOC-SLUG}.md<br/>Editar APENAS seções em VIOLATIONS[]<br/>Manter status Em revisão"]
        FIX --> GATE

        GATE_RESULT -->|"PASS"| HUMAN

        HUMAN["STEP 4b: VALIDAÇÃO HUMANA<br/>Apresentar documento<br/>Status: aguardando aprovação"] --> P1["P1: Conteúdo aderente<br/>às necessidades?"]
        P1 --> P2["P2: Novos documentos<br/>de entrada?"]
        P2 --> P3["P3: Novos inputs,<br/>mudanças de escopo<br/>ou ajustes técnicos?"]
        P3 --> P4["P4 HITL: Novas informações<br/>antes de prosseguir?"]
        P4 --> HUMAN_DEC{Decisão<br/>do Humano?}

        HUMAN_DEC -->|"SIM / NÃO / NÃO / NÃO"| COMPLIANCE
        HUMAN_DEC -->|"Fornece novos inputs<br/>(P2, P3 ou P4)"| GEN
    end

    COMPLIANCE(["✅ STATUS: COMPLIANCE<br/>Documento congelado<br/>Próximo documento destravado"])

    style GEN fill:#0984e3,color:#fff
    style GATE fill:#fdcb6e,color:#333
    style FIX fill:#e17055,color:#fff
    style HUMAN fill:#6c5ce7,color:#fff
    style HITL fill:#6c5ce7,color:#fff
    style P1 fill:#dfe6e9,stroke:#6c5ce7,color:black
    style P2 fill:#dfe6e9,stroke:#6c5ce7,color:black
    style P3 fill:#dfe6e9,stroke:#6c5ce7,color:black
    style P4 fill:#dfe6e9,stroke:#6c5ce7,color:black
    style COMPLIANCE fill:#00b894,color:#fff
    style LOOP fill:#fff3e0,stroke:#f39c12
```

---

## 6. Matriz de UPSTREAM_DOCS — Fluxo de Dependências (v2.0)

```mermaid
flowchart LR
    D001["001: Charter"] --> D002["002: Stakeholder Map"]
    D001 --> D005["005: BRD"]
    D001 --> D010["010: FRD"]
    D001 --> D015["015: RTM-F1"]
    D001 --> D020["020: SRS"]
    D001 --> D030["030: SAD"]
    D001 --> D035["035: HLD"]
    D001 --> D040["040: LLD"]
    D001 --> D060["060: EAP/WBS"]
    D001 --> D065["065: Cronograma"]
    D001 --> D070["070: Orçamento"]
    D001 --> D075["075: Comunicação"]
    D001 --> D080["080: Riscos"]
    D001 --> D085["085: Gestão Mudanças"]
    D001 --> D090["090: Deploy"]
    D001 --> D095["095: Manuais User"]
    D001 --> D105["105: Termo Aceite"]
    D001 --> D115["115: Encerramento"]

    D002 --> D005
    D002 --> D075

    D005 --> D010
    D010 --> D015
    D005 --> D015
    D002 --> D015

    D015 --> D020
    D005 --> D020
    D010 --> D020

    D020 --> D025["025: RTM-F2"]
    D015 --> D025

    D025 --> D030
    D010 --> D030

    D030 --> D035
    D030 --> D040
    D030 --> D090
    D030 --> D100["100: Manuais Ops"]

    D035 --> D040
    D035 --> D090

    D040 --> D045["045: EST-PLAN"]
    D040 --> D090
    D040 --> D060

    D020 --> D045
    D045 --> D050["050: EST-CASES"]
    D010 --> D050
    D020 --> D050

    D050 --> D055["055: Qualidade"]
    D045 --> D055

    D050 --> D060
    D050 --> D065
    D045 --> D065

    D060 --> D065
    D060 --> D070
    D060 --> D085

    D065 --> D070
    D080 --> D085

    D090 --> D100
    D010 --> D095
    D020 --> D095

    D045 --> D105
    D055 --> D105

    D105 --> D110["110: Lições"]
    D110 --> D115

    style D001 fill:#6c5ce7,color:#fff
    style D002 fill:#6c5ce7,color:#fff
    style D005 fill:#0984e3,color:#fff
    style D010 fill:#0984e3,color:#fff
    style D015 fill:#0984e3,color:#fff
    style D020 fill:#6c5ce7,color:#fff
    style D025 fill:#6c5ce7,color:#fff
    style D030 fill:#6c5ce7,color:#fff
    style D035 fill:#6c5ce7,color:#fff
    style D040 fill:#00b894,color:#fff
    style D045 fill:#00b894,color:#fff
    style D050 fill:#00b894,color:#fff
    style D055 fill:#00b894,color:#fff
    style D060 fill:#00b894,color:#fff
    style D065 fill:#e17055,color:#fff
    style D070 fill:#e17055,color:#fff
    style D075 fill:#e17055,color:#fff
    style D080 fill:#e17055,color:#fff
    style D085 fill:#e17055,color:#fff
    style D090 fill:#e17055,color:#fff
    style D095 fill:#a29bfe,color:#fff
    style D100 fill:#a29bfe,color:#fff
    style D105 fill:#a29bfe,color:#fff
    style D110 fill:#a29bfe,color:#fff
    style D115 fill:#a29bfe,color:#fff
```

---

## 7. Efeitos Cascata — Impacto de Modificações (v2.0)

```mermaid
flowchart TD
    subgraph IMPACTO_ALTO["🔴 Impacto ALTO — Regenera 10+ docs"]
        CHARTER["001 Charter modificado"] -->|"Impacta TODOS<br/>os 21 docs downstream"| ALL["Regenerar 002 ao 115"]
        BRD["005 BRD modificado"] -->|"Impacta 13+ docs"| BRD_CASCADE["010, 015, 020, 025, 030, 035, 040,<br/>045, 050, 060, 065, 070, 090, 095"]
    end

    subgraph IMPACTO_MEDIO["🟡 Impacto MÉDIO — Regenera 4-8 docs"]
        FRD["010 FRD modificado"] -->|"Impacta 6 docs"| FRD_CASCADE["015, 020, 025, 030, 050, 095"]
        SAD["030 SAD modificado"] -->|"Impacta 5 docs"| SAD_CASCADE["035, 040, 045, 090, 100"]
        LLD["040 LLD modificado"] -->|"Impacta 6 docs<br/>+ WATERFALL-ESTIMATION"| LLD_CASCADE["045, 050, 060, 065,<br/>070, 090 + ⚡PERT"]
        HLD["035 HLD modificado"] -->|"Impacta + ⚡ROM"| HLD_CASCADE["040, 090 + ⚡WATERFALL-ESTIMATION<br/>UPSTREAM/DISCOVERY"]
    end

    subgraph IMPACTO_BAIXO["🟢 Impacto BAIXO — Regenera 1-3 docs"]
        EAP["060 EAP modificado"] -->|"Impacta 3 docs<br/>+ WATERFALL-ESTIMATION"| EAP_CASCADE["065, 070, 085 + ⚡PERT"]
        EST_PLAN["045 EST-PLAN"] -->|"Impacta 6 docs"| EST_CASCADE["050, 055, 060, 065, 070, 105"]
    end

    style IMPACTO_ALTO fill:#ffcccc,stroke:#d63031
    style IMPACTO_MEDIO fill:#fff3e0,stroke:#e65100
    style IMPACTO_BAIXO fill:#e8f5e9,stroke:#2e7d32
```

---

## 8. Integração com WATERFALL-ESTIMATION

```mermaid
flowchart LR
    subgraph WATERFALL["WATERFALL Docs"]
        direction TB
        W06["#6 HLD ✅"] --> GATE1{{Gate UPSTREAM}}
        W07["#7 LLD ✅"] --> W11["#11 EAP ✅"]
        W11 --> GATE2{{Gate DOWNSTREAM}}
    end

    subgraph ESTIMATION["WATERFALL-ESTIMATION"]
        direction TB
        UP["UPSTREAM/DISCOVERY<br/>F1: ROM ±50%<br/>F2: Scope Snapshot<br/>F3: Governance GO/NO-GO"]
        DOWN["DOWNSTREAM/REFINEMENT<br/>F4: PERT ±15-25%<br/>F5: Scope Snapshot<br/>F6: Cronograma Calculado<br/>F7: Orçamento Calculado"]
    end

    subgraph WATERFALL_OUT["WATERFALL Docs — Consumidores"]
        direction TB
        W12["#12 Cronograma/Gantt<br/>← CRONOGRAMA-CALCULADO"]
        W13["#13 Orçamento<br/>← ORCAMENTO-CALCULADO"]
    end

    GATE1 -->|"Opcional"| UP
    GATE1 -->|"Pular"| W07
    UP -->|"GO ✅"| W07
    UP -->|"NO-GO ❌"| CANCEL["Projeto Cancelado"]

    GATE2 -->|"Opcional"| DOWN
    GATE2 -->|"Pular"| W12
    DOWN -->|"UPSTREAM_DOCS adicional"| W12
    DOWN -->|"UPSTREAM_DOCS adicional"| W13

    style WATERFALL fill:#e3f2fd,stroke:#0984e3
    style ESTIMATION fill:#fff3e0,stroke:#e65100
    style WATERFALL_OUT fill:#e8f5e9,stroke:#2e7d32
    style CANCEL fill:#ffcccc,stroke:#d63031
```

---

## 9. Diagrama de Estados — Visão Unificada

```mermaid
stateDiagram-v2
    ["*"] --> Bootstrap

    state Bootstrap {
        ["*"] --> SolicitarInputs
        SolicitarInputs --> ValidarStack
        ValidarStack --> ColetarTime
        ColetarTime --> ConfirmarCaminhos
        ConfirmarCaminhos --> CriarEstrutura
        CriarEstrutura --> AuditarStatus
        AuditarStatus --> Resumo
    }

    Bootstrap --> Fase1_Inicializacao

    state Fase1_Inicializacao {
        ["*"] --> Doc1_Charter
        state Doc1_Charter {
            ["*"] --> Gen1
            Gen1 --> Gate1
            Gate1 --> Fix1: FAIL
            Fix1 --> Gate1
            Gate1 --> Human1: PASS
            Human1 --> Gen1: Novos inputs
            Human1 --> Done1: Aprovado
        }
        Doc1_Charter --> ["*"]: COMPLIANCE
    }

    Fase1_Inicializacao --> Fase2_Requisitos

    state Fase2_Requisitos {
        ["*"] --> Doc2_BRD
        Doc2_BRD --> Doc3_SRS: COMPLIANCE
        Doc3_SRS --> Doc4_RTM: COMPLIANCE
        Doc4_RTM --> ["*"]: COMPLIANCE
    }

    Fase2_Requisitos --> Fase3_Design

    state Fase3_Design {
        ["*"] --> Doc5_SAD
        Doc5_SAD --> Doc6_HLD: COMPLIANCE
        Doc6_HLD --> GateUpstream: COMPLIANCE
        GateUpstream --> Doc7_LLD: GO ou Pular
        GateUpstream --> Cancelado: NO-GO
        Doc7_LLD --> ["*"]: COMPLIANCE
    }

    Fase3_Design --> Fase4_Testes

    state Fase4_Testes {
        ["*"] --> Doc8_TestPlan
        Doc8_TestPlan --> Doc9_TestCases: COMPLIANCE
        Doc9_TestCases --> Doc10_Qualidade: COMPLIANCE
        Doc10_Qualidade --> ["*"]: COMPLIANCE
    }

    Fase4_Testes --> Fase5_Planejamento

    state Fase5_Planejamento {
        ["*"] --> Doc11_EAP
        Doc11_EAP --> GateDownstream: COMPLIANCE
        GateDownstream --> Doc12_Cronograma: Pular ou PERT
        Doc12_Cronograma --> Doc13_Orcamento: COMPLIANCE
        Doc13_Orcamento --> Doc14_Comunicacao: COMPLIANCE
        Doc13_Orcamento --> Doc15_Riscos: COMPLIANCE
        Doc14_Comunicacao --> ["*"]: COMPLIANCE
        Doc15_Riscos --> ["*"]: COMPLIANCE
    }

    Fase5_Planejamento --> Fase6_Implantacao

    state Fase6_Implantacao {
        ["*"] --> Doc16_Deploy
        Doc16_Deploy --> Doc17_User: COMPLIANCE
        Doc17_User --> Doc18_Ops: COMPLIANCE
        Doc18_Ops --> Doc19_Aceite: COMPLIANCE
        Doc19_Aceite --> Doc20_Licoes: COMPLIANCE
        Doc20_Licoes --> ["*"]: COMPLIANCE
    }

    Fase6_Implantacao --> GitWorkflow
    GitWorkflow --> ["*"]
```

---

## 10. Git Workflow de Finalização

```mermaid
flowchart LR
    ALL_COMPLIANCE(["22 docs COMPLIANCE ✅"]) --> F1["F.1: git add -A<br/>git commit"]

    F1 --> F2["F.2: git push origin<br/>feature/PROJECT_ID_NAME-waterfall-docs"]

    F2 --> F3["F.3: gh pr create<br/>--base main"]

    F3 --> F4["F.4: gh pr merge<br/>--merge --delete-branch"]

    F4 --> F5["F.5: git checkout main<br/>git branch -d WORK_BRANCH"]

    F5 --> DONE(["✅ Concluído"])

    style ALL_COMPLIANCE fill:#00b894,color:#fff
    style DONE fill:#00b894,color:#fff
```

---

## 11. Tabela de Símbolos e Convenções

| Símbolo/Cor | Significado |
|-------------|-------------|
| 🟣 Roxo (`#6c5ce7`) | Bootstrap / Validação Humana / Git Workflow / Fase 1 (docs 001-002) / Fase 2 (docs 020-035) |
| 🔵 Azul (`#0984e3`) | Fase 1: Requisitos de Negócio (docs 005-015) |
| 🟢 Verde (`#00b894`) | Fase 3: Engenharia Detalhada e Qualidade / Compliance |
| 🟠 Terracota (`#e17055`) | Fase 4: Planejamento e Baseline / Correções (FIX) |
| 🟣 Lilás (`#a29bfe`) | Fase 5: Encerramento e Operação |
| 🟡 Amarelo (`#fdcb6e`) | Gate / Decisões / Checkpoints |
| 🔴 Vermelho (`#d63031`) | Cancelado / Impacto ALTO |
| 🟤 Laranja escuro (`#e65100`) | WATERFALL-ESTIMATION UPSTREAM / Impacto MÉDIO |
| 🟢 Verde escuro (`#2e7d32`) | WATERFALL-ESTIMATION DOWNSTREAM / Impacto BAIXO |
| 🔲 Linha tracejada | Loop de retrabalho (GATE→FIX→GATE) |
| 🔲 Linha sólida | Fluxo sequencial normal |
| 🎯 | Gate de Estimativa |
| ⚡ | Dispara WATERFALL-ESTIMATION |
| 📥 | Inputs consumidos |
| 📄 | Documento WATERFALL |
| #N | Número do documento na sequência WATERFALL |

---

## 12. Mapeamento de Cores por Fase

| Fase | Cor | Docs |
|------|-----|------|
| **Fase 1: Iniciação e Requisitos de Negócio** | 🔵 Azul / 🟣 Roxo | 001, 002, 005, 010, 015 |
| **Fase 2: Especificação de Sistema e Arquitetura** | 🟣 Roxo | 020, 025, 030, 035 |
| **Fase 3: Engenharia Detalhada e Qualidade** | 🟢 Verde | 040, 045, 050, 055, 060 |
| **Fase 4: Planejamento e Baseline** | 🟠 Terracota | 065, 070, 075, 080, 085, 090 |
| **Fase 5: Encerramento e Operação** | 🟣 Lilás | 095, 100, 105, 110, 115 |

---

> **📁 Arquivos relacionados:**
> - `PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md` — Documento fonte (v2.0)
> - `../PROMPT-ROADMAP-GENERATE-WATERFALL-ESTIMATION.md` — Roadmap companion de estimativa
> - `../PROMPT-ROADMAP-GENERATE-SOURCING-FACTORY-BIDDING.md` — Roadmap de Sourcing (consome docs WATERFALL)
> - `PROMPT-GENERATE-*.md` — 22 prompts geradores
> - `PROMPT-GATE-*.md` — 22 prompts de auditoria
> - `PROMPT-FIX-*.md` — 22 prompts de correção
> - `PROMPT-GATE-*.md` — 20 prompts de auditoria
> - `PROMPT-FIX-*.md` — 20 prompts de correção
