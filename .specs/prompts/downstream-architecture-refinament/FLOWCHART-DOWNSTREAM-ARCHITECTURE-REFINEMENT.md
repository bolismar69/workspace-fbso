# Flowchart: Downstream Architecture Refinement

> Versão: 2.3 — Baseado no roadmap v2.3 (Skills Map + Team Capacity + Stack Validation + Checkpoint HITL)
> Este documento contém os diagramas de fluxo do processo de Downstream Architecture Refinement.

---

## 1. Diagrama Macro: Blocos, Fases e Barreiras

```mermaid
flowchart TB
    %% ═══════════════════════════════════════════
    %% ESTILOS GLOBAIS
    %% ═══════════════════════════════════════════
    classDef bootstrap fill:#1a1a2e,stroke:#16213e,color:#e0e0ff,stroke-width:2px
    classDef blocoA fill:#1a2a3a,stroke:#2a4a6a,color:#c8e0ff,stroke-width:2px
    classDef blocoB fill:#2a1a3a,stroke:#4a2a6a,color:#e0c8ff,stroke-width:2px
    classDef blocoC fill:#3a2a1a,stroke:#6a4a2a,color:#ffe0c8,stroke-width:2px
    classDef barreira fill:#5a1a1a,stroke:#8a2a2a,color:#ffc8c8,stroke-width:3px
    classDef gate fill:#1a3a3a,stroke:#2a6a6a,color:#c8ffff,stroke-width:3px
    classDef decisao fill:#3a3a1a,stroke:#6a6a2a,color:#ffffc8,stroke-width:2px
    classDef condicional fill:#1a2a2a,stroke:#2a4a4a,color:#a0c8c8,stroke-width:2px,stroke-dasharray: 5 5
    classDef downstream fill:#1a3a1a,stroke:#2a6a2a,color:#a0ffa0,stroke-width:2px

    %% ═══════════════════════════════════════════
    %% PRÉ-CONDIÇÕES E BOOTSTRAP
    %% ═══════════════════════════════════════════
    subgraph PRECOND[" 📋 Pré-condições de Entrada (Negócio) "]
        direction TB
        CH[Project Charter<br/>01-PROJECT-CHARTER.md]
        BRD[BRD<br/>02-BRD.md]
        EP[Épicos<br/>03-EPICS.md + epics/*.md]
        FEAT[Features<br/>04-FEATURES.md + features/*.md]
        US[User Stories<br/>05-USER-STORIES.md + user-stories/*.md]
        INPUTS["Inputs do Usuário<br/>PROJECT_PATH, PROJECT_ID_NAME<br/>TECHNICAL_SOLUTION_PATH, etc."]
        STACK[Stack Corporativa FBSO<br/>STACK-PADROES-CORPORATIVOS.md<br/>DigitalOcean, Cloudflare, Kong<br/>Keycloak, PostgreSQL, K8s/Istio]
        HITL_IN["🟢 Checkpoint HITL<br/>PROJECT_PROMPT_INPUTS<br/>Sempre perguntar ao usuário<br/>se deseja adicionar informações"]
    end

    subgraph PRECOND_OPT[" 📂 Pré-condição Opcional (Referência) "]
        direction TB
        UPSTREAM["Upstream Architecture Discovery<br/>upstream-architecture-discovery/<br/>🔹 Referência para design (Bloco A)<br/>🔹 Comparação na F12 (condicional)<br/>🚫 NUNCA usado na estimativa"]
    end

    %% ═══════════════════════════════════════════
    %% FASE 0: BOOTSTRAP
    %% ═══════════════════════════════════════════
    subgraph F0[" 🔧 FASE 0: Bootstrap Inteligente "]
        direction TB
        F0A["1. Solicitar inputs ao usuário"] --> F0B["2. Validar pré-condição de negócio<br/>(Charter, BRD, Épicos, Features, US)"]
        F0B --> F0C["3. Detectar upstream (informativo)<br/>Se existir: listar artefatos<br/>Se não: prossegue sem referência"]
        F0C --> F0D{"4. project-document-discovery<br/>Classificar projeto em 10 dimensões<br/>(4-signal algorithm)"}
        F0D --> F0E["5. discovery-process<br/>Framing → Synthesis → Experiments"]
        F0E --> F0F["6. Auditar User Stories<br/>Contar US por épico/feature<br/>Verificar distribuição e gaps"]
        F0F --> F0G["7. Exibir caminhos derivados<br/>e solicitar confirmação"]
        F0G --> F0H["8. Criar estrutura:<br/>mkdir -p {DOWNSTREAM_REFINEMENT_PATH}"]
        F0H --> F0I["9. Auditar artefatos existentes<br/>em downstream-architecture-refinement/"]
        F0I --> F0J["10. Apresentar resumo da situação<br/>e iniciar primeira fase pendente"]
        F0J --> F0K["11. Coletar Skills, Capacidade e Stack<br/>📋 PROJECT-TEAM-SKILLS-MAP<br/>👥 PROJECT-TEAM-CAPACITY<br/>🔧 PROJECT-STACK<br/>⏸️ Validar contra Stack Corporativa<br/>⏸️ Aguardar confirmação do usuário"]
    end

    %% ═══════════════════════════════════════════
    %% BLOCO A: ARCHITECTURE DEEP-DIVE (F1-F7)
    %% ═══════════════════════════════════════════
    subgraph BA[" 🏗️ BLOCO A: Architecture Deep-Dive (Detail-Level) "]
        direction LR
        F1["🏷️ F1<br/>DETAIL-LEVEL<br/>PRD<br/>⏸️ Alinhamento Negócio↔TI"]
        F2["🏗️ F2<br/>ARCHITECTURE<br/>DEFINITION<br/>⏸️ Validar Stack"]
        F3["🔒 F3<br/>SECURITY<br/>DEFINITION<br/>⏸️ Validar Stack"]
        F4["🗄️ F4<br/>DATA<br/>DEFINITION<br/>⏸️ Validar Stack"]
        F5["🚀 F5<br/>DEVOPS-SRE<br/>DEFINITION<br/>⏸️ Validar Stack"]
        F6["🧪 F6<br/>TEST-STRATEGY<br/>DEFINITION<br/>⏸️ Validar Stack"]
        F7["☁️ F7<br/>INFRA-CLOUD<br/>DEFINITION<br/>⏸️ Validar Stack"]
        F1 --> F2 --> F3 --> F4 --> F5 --> F6 --> F7
    end

    %% ═══════════════════════════════════════════
    %% BARREIRA A
    %% ═══════════════════════════════════════════
    BARRA{"⛔ BARREIRA A<br/>PRD + 6 disciplinas OK?<br/>Consistência horizontal?<br/>Seção FBSO presente em F2-F7?<br/>Detalhamento suficiente para<br/>time de TI validar viabilidade?"}

    %% ═══════════════════════════════════════════
    %% BLOCO B: BOTTOM-UP ESTIMATION (F8-F10)
    %% ═══════════════════════════════════════════
    subgraph BB[" ⭐ BLOCO B: Bottom-Up Estimation — INDEPENDENTE (PERT ±15-25%) "]
        direction LR
        F8["📐 F8<br/>BOTTOM-UP<br/>PERT-ESTIMATE<br/>⭐ Three-point estimation<br/>62 US individuais<br/>O/ML/P → PERT ±15-25%"]
        F9["👥 F9<br/>RESOURCE<br/>ALLOCATION<br/>Time Necessário<br/>Alocação por épico"]
        F10["⚠️ F10<br/>RISK-ADJUSTED<br/>ESTIMATE<br/>3 cenários<br/>Conservador/PERT/Pessimista"]
        F8 --> F9 --> F10
    end

    %% REGRA DE INDEPENDÊNCIA
    INDEP["🚫 REGRA CRÍTICA<br/>Estimativa 100% INDEPENDENTE<br/>NUNCA usa ROM upstream<br/>como baseline ou referência<br/>Calculada do zero, US por US"]

    %% ═══════════════════════════════════════════
    %% BARREIRA B
    %% ═══════════════════════════════════════════
    BARRB{"⛔ BARREIRA B<br/>Todas US estimadas individualmente?<br/>QA ≥ 25%? Arch ≥ 5%?<br/>Consistência Prazo×Horas?<br/>Outliers identificados?<br/>🚫 Evidência de contaminação<br/>pelo ROM upstream?"}

    %% ═══════════════════════════════════════════
    %% BLOCO C: SCOPE SNAPSHOT + CROSS-CHECK (F11-F12)
    %% ═══════════════════════════════════════════
    subgraph BC[" 📸 BLOCO C: Scope Snapshot + Cross-Check "]
        direction LR
        F11["📸 F11<br/>SCOPE<br/>SNAPSHOT<br/>Foto imutável do<br/>escopo estimado"]
        F12{"🔗 F12<br/>UPSTREAM<br/>COMPARISON<br/>Condicional:<br/>só se upstream existir"}
        F11 --> F12
    end

    %% ═══════════════════════════════════════════
    %% BARREIRA C
    %% ═══════════════════════════════════════════
    BARRC{"⛔ BARREIRA C<br/>Scope Snapshot cobre 100%<br/>das US estimadas?<br/>Se upstream existe:<br/>relatório comparativo gerado?<br/>Relatório NÃO altera PERT?"}

    %% ═══════════════════════════════════════════
    %% GATE DE ESTIMATIVA
    %% ═══════════════════════════════════════════
    GATE["📊 GATE DE ESTIMATIVA<br/>ESTIMATE-READY<br/><br/>Resumo Executivo:<br/>1. Escopo Estimado (Snapshot)<br/>2. Design Detalhado (6 disciplinas)<br/>3. Estimativa PERT ±15-25%<br/>4. Cross-Check com Upstream<br/>5. Parecer de Viabilidade Técnica"]
    
    DECISAO{"🏛️ Decisão do<br/>Comitê / Time de TI"}
    
    APPROVED["✅ ESTIMATE ACCEPTED<br/>━━━━━━━━━━━━━━<br/>Estimativa aprovada<br/>como baseline de<br/>prazo e orçamento"]
    
    PENDING["⚠️ PENDÊNCIAS<br/>━━━━━━━━━━━━━━<br/>Retorna ao bloco<br/>com gaps para ajuste"]
    
    DWN1["📋 Dispara Roadmap:<br/>TECHNICAL-DEFINITIONS<br/>(Refinamento técnico e<br/>planejamento de sprints)"]
    DWN2["🔄 Dispara Revisão:<br/>Re-estimar após<br/>2 sprints de execução<br/>(refinar para ±10%)"]

    %% ═══════════════════════════════════════════
    %% CONEXÕES PRINCIPAIS
    %% ═══════════════════════════════════════════
    INPUTS --> F0A
    CH --> F0B
    BRD --> F0B
    EP --> F0B
    FEAT --> F0B
    US --> F0B
    STACK --> F0K
    HITL_IN --> F0A
    UPSTREAM -.->|"referência (design)"| F2
    UPSTREAM -.->|"comparação (F12)"| F12

    F0K --> F1
    F7 --> BARRA
    BARRA -->|"✅ COMPLIANCE"| F8
    BARRA -->|"❌ NÃO COMPLIANCE"| F1

    INDEP -.->|"vigia"| F8

    F10 --> BARRB
    BARRB -->|"✅ COMPLIANCE"| F11
    BARRB -->|"❌ NÃO COMPLIANCE"| F8

    F12 --> BARRC
    BARRC -->|"✅ COMPLIANCE"| GATE
    BARRC -->|"❌ NÃO COMPLIANCE"| F11

    GATE --> DECISAO
    DECISAO -->|"Accepted ✅"| APPROVED
    DECISAO -->|"Pendências ⚠️"| PENDING
    PENDING --> F8
    APPROVED --> DWN1
    APPROVED --> DWN2

    %% ═══════════════════════════════════════════
    %% APLICAÇÃO DE ESTILOS
    %% ═══════════════════════════════════════════
    class F0,F0A,F0B,F0C,F0D,F0E,F0F,F0G,F0H,F0I,F0J,F0K bootstrap
    class BA,F1,F2,F3,F4,F5,F6,F7 blocoA
    class BB,F8,F9,F10 blocoB
    class BC,F11,F12 blocoC
    class BARRA,BARRB,BARRC barreira
    class GATE,APPROVED,PENDING gate
    class DECISAO decisao
    class F12 condicional
    class DWN1,DWN2 downstream
    class PRECOND,CH,BRD,EP,FEAT,US,INPUTS,STACK,HITL_IN bootstrap
    class PRECOND_OPT,UPSTREAM condicional
    class INDEP barreira
```

---

## 2. Ciclo HITL por Fase (Generate → Gate → Fix → Human Validation)

Este ciclo se repete em **todas as 12 fases** (F1 a F12). Cada fase gera um artefato, que passa pelo gate, e se necessário pelo fix, até atingir COMPLIANCE com validação humana.

```mermaid
flowchart TB
    %% ═══════════════════════════════════════════
    %% ESTILOS
    %% ═══════════════════════════════════════════
    classDef generate fill:#1a3a1a,stroke:#2a6a2a,color:#a0ffa0,stroke-width:2px
    classDef gate fill:#3a2a1a,stroke:#6a4a2a,color:#ffe0a0,stroke-width:2px
    classDef fix fill:#3a1a1a,stroke:#6a2a2a,color:#ffa0a0,stroke-width:2px
    classDef human fill:#1a2a3a,stroke:#2a4a6a,color:#a0c8ff,stroke-width:2px
    classDef compliance fill:#1a3a3a,stroke:#2a6a6a,color:#a0ffff,stroke-width:2px
    classDef nok fill:#5a1a1a,stroke:#8a2a2a,color:#ffc0c0,stroke-width:2px

    subgraph HITL[" 🔄 CICLO HITL POR FASE — Human-In-The-Loop "]
        direction TB
        
        %% PASSO 1: GENERATE
        subgraph GEN[" 1. GENERATE — Geração do Artefato "]
            G1["📥 Recebe inputs disponíveis<br/>(docs negócio + artefatos de fases anteriores<br/>+ upstream discovery se existir)"]
            G2["🤖 Executa PROMPT-GENERATE<br/>da fase correspondente"]
            G2a["⏸️ Para F2-F7: Validar Skills,<br/>Capacidade e Stack com usuário<br/>Verificar STACK-PADROES-CORPORATIVOS<br/>🟢 Checkpoint HITL: solicitar<br/>informações adicionais se necessário"]
            G3["📄 Artefato gerado:<br/>DETAIL-LEVEL-XXX.md<br/>(ou PERT-ESTIMATE / SCOPE-SNAPSHOT / etc.)"]
            G1 --> G2 --> G2a --> G3
        end

        %% PASSO 2: GATE (AUDITORIA INTERNA)
        subgraph GT[" 2. GATE — Auditoria Interna da IA "]
            GT1["🔍 Executa PROMPT-GATE<br/>da fase correspondente"]
            GT2{"📊 Resultado<br/>da Auditoria?"}
            GT_NOK["🚨 NÃO COMPLIANCE<br/>Lista conflitos com:<br/>ID, descrição, impacto, sugestão"]
            GT_OK["✅ PRÉ-COMPLIANCE INTERNO<br/>AGUARDANDO VALIDAÇÃO HUMANA"]
            GT1 --> GT2
            GT2 -->|"Encontrou erros"| GT_NOK
            GT2 -->|"100% OK"| GT_OK
        end

        %% PASSO 3: FIX (CORREÇÃO CIRÚRGICA)
        subgraph FX[" 3. FIX — Correção Cirúrgica "]
            FX1["🔧 Executa PROMPT-FIX<br/>Processa NCs por prioridade P0-P3"]
            FX2["✂️ Aplica correções cirúrgicas<br/>Nunca reescreve do zero"]
            FX3["🔄 Retorna ao GATE<br/>para revalidação"]
            FX1 --> FX2 --> FX3
        end

        %% PASSO 4: VALIDAÇÃO HUMANA
        subgraph HUM[" 4. VALIDAÇÃO HUMANA — 3 Perguntas Obrigatórias "]
            H_Q1["❓ 1. O documento está em compliance<br/>e alinhado com os documentos base<br/>(negócio + upstream discovery)?"]
            H_Q2["❓ 2. Deseja enviar mais documentos<br/>para enriquecer este artefato?"]
            H_Q3["❓ 3. Deseja enviar mais informações<br/>via input de texto?"]
            H_DEC{"🧑‍💼 Respostas<br/>do Humano?"}
            H_SIM["Sim, Não, Não"]
            H_NOVOS["Novos documentos<br/>ou informações fornecidos"]
            H_Q1 --> H_Q2 --> H_Q3 --> H_DEC
        end

        %% RESULTADO FINAL
        COMP["🎉 STATUS: COMPLIANCE<br/>Arquivo congelado<br/>Próxima fase destravada"]
        RETRO["🔙 RETROCESSO<br/>Retorna ao GENERATE<br/>com documento atual + novos insumos"]
    end

    %% CONEXÕES DO CICLO
    GEN --> GT
    GT_NOK --> FX
    FX3 --> GT1
    GT_OK --> H_Q1
    H_DEC -->|"Sim, Não, Não"| H_SIM
    H_DEC -->|"Novos inputs"| H_NOVOS
    H_SIM --> COMP
    H_NOVOS --> RETRO
    RETRO --> G1

    %% ESTILOS
    class G1,G2,G2a,G3 generate
    class GT1,GT2,GT_NOK,GT_OK gate
    class FX1,FX2,FX3 fix
    class H_Q1,H_Q2,H_Q3,H_DEC,H_SIM,H_NOVOS human
    class COMP compliance
    class RETRO nok
```

---

## 3. Barreiras de Bloco (Gating Rules)

Fluxo de decisão de cada barreira entre blocos:

```mermaid
flowchart LR
    %% ESTILOS
    classDef ok fill:#1a3a1a,stroke:#2a6a2a,color:#a0ffa0,stroke-width:2px
    classDef nok fill:#5a1a1a,stroke:#8a2a2a,color:#ffc0c0,stroke-width:2px
    classDef warn fill:#3a3a1a,stroke:#6a6a2a,color:#ffffa0,stroke-width:2px

    subgraph BARREIRAS[" ⛔ BARRAS DE BLOCO — Validação Acumulada "]
        direction TB
        
        BA{"⛔ Barreira A<br/>Após Bloco A (F7)"}
        BA_OK["✅ PRD + 6 disciplinas OK<br/>Consistência horizontal validada<br/>Seção FBSO presente em F2-F7<br/>Detalhamento suficiente para<br/>time de TI validar viabilidade"]
        BA_NOK["❌ Artefato muito similar ao<br/>Discovery-Level (sem refinamento)<br/>ou Seção FBSO ausente em F2-F7"]
        
        BB{"⛔ Barreira B<br/>Após Bloco B (F10)"}
        BB_OK["✅ Todas US com estimativa individual<br/>QA ≥ 25% · Arch ≥ 5%<br/>Consistência Prazo×Horas validada<br/>Outliers identificados<br/>🚫 Sem contaminação do ROM upstream"]
        BB_NOK["❌ US sem estimativa individual<br/>ou PERT sem IC 95%<br/>ou Evidência de contaminação<br/>pelo ROM upstream"]
        BB_WARN["⚠️ QA < 25% (Risco de Débito)<br/>ou Arch < 5% (Subinvestimento)<br/>ou Divergência Prazo×Horas >50%"]
        
        BC{"⛔ Barreira C<br/>Após Bloco C (F12)"}
        BC_OK["✅ Scope Snapshot cobre 100%<br/>das US estimadas<br/>Se upstream existe:<br/>relatório comparativo gerado<br/>Se upstream não existe:<br/>F12 pulada, barreira satisfeita<br/>Relatório NÃO altera PERT"]
        BC_NOK["❌ Scope Snapshot incompleto<br/>ou Relatório alterou estimativa PERT"]
        
        BA --> BA_OK
        BA --> BA_NOK
        BB --> BB_OK
        BB --> BB_NOK
        BB --> BB_WARN
        BC --> BC_OK
        BC --> BC_NOK
    end
    
    class BA_OK,BB_OK,BC_OK ok
    class BA_NOK,BB_NOK,BC_NOK nok
    class BB_WARN warn
```

---

## 4. Gate de Estimativa (ESTIMATE-READY)

Decisão final após a Barreira C — validação da estimativa como baseline de prazo e orçamento:

```mermaid
flowchart TB
    classDef gate fill:#1a3a3a,stroke:#2a6a6a,color:#c8ffff,stroke-width:3px
    classDef approved fill:#1a3a1a,stroke:#2a6a2a,color:#a0ffa0,stroke-width:2px
    classDef pending fill:#3a3a1a,stroke:#6a6a2a,color:#ffffa0,stroke-width:2px
    classDef roadmaps fill:#2a1a3a,stroke:#4a2a6a,color:#e0c8ff,stroke-width:2px
    classDef revisao fill:#1a2a3a,stroke:#2a4a6a,color:#a0c8ff,stroke-width:2px

    SNAPSHOT_IN["📸 SCOPE-SNAPSHOT.md<br/>Aprovado na Barreira C"] --> GATE_EXEC

    subgraph GATER[" 📊 GATE DE ESTIMATIVA — ESTIMATE-READY "]
        direction TB
        GATE_EXEC["📊 Resumo Executivo da<br/>Análise de Viabilidade e Estimativa"]
        GATE_SCOPE["1. Escopo Estimado<br/>Snapshot das US incluídas (F11)"]
        GATE_DESIGN["2. Design Detalhado<br/>6 disciplinas validadas pelo time de TI<br/>(Bloco A: F1-F7)"]
        GATE_PERT["3. Estimativa Bottom-Up PERT<br/>Faixa ±15-25% de confiança<br/>Resource Allocation (F9)<br/>Risk-Adjusted Estimate (F10)"]
        GATE_CROSS["4. Cross-Check com Upstream<br/>Relatório comparativo (F12)<br/>🔹 Se upstream existir<br/>🔹 Se não: 'sem ROM para comparação'"]
        GATE_PARECER["5. Parecer de Viabilidade<br/>Recomendação do time de TI<br/>sobre a viabilidade do projeto"]
    end

    GATE_EXEC --> GATE_SCOPE --> GATE_DESIGN --> GATE_PERT --> GATE_CROSS --> GATE_PARECER
    
    GATE_DEC{"🏛️ Decisão do<br/>Comitê /<br/>Time de TI"}
    
    PENDING_ST["⚠️ PENDÊNCIAS<br/>━━━━━━━━━━━━━━<br/>Retorna ao bloco<br/>com gaps para ajuste<br/>Após correção:<br/>revalidar Barreira B → C"]
    
    APPROVED_ST["✅ ESTIMATE ACCEPTED<br/>━━━━━━━━━━━━━━<br/>Estimativa aprovada<br/>como baseline de<br/>prazo e orçamento"]
    
    subgraph DWN[" 📋 Ações Disparadas "]
        direction TB
        DWN_TECH["TECHNICAL-DEFINITIONS<br/>Refinamento técnico e<br/>planejamento de sprints"]
        DWN_REV["🔄 Revisão da Estimativa<br/>Após 2 sprints de execução<br/>Refinar para ±10%"]
    end

    GATE_PARECER --> GATE_DEC
    GATE_DEC -->|"Accepted ✅"| APPROVED_ST
    GATE_DEC -->|"Pendências ⚠️"| PENDING_ST
    PENDING_ST --> GATE_SCOPE
    APPROVED_ST --> DWN_TECH
    APPROVED_ST --> DWN_REV
    
    class GATE_EXEC,GATE_SCOPE,GATE_DESIGN,GATE_PERT,GATE_CROSS,GATE_PARECER,GATE_DEC gate
    class APPROVED_ST approved
    class PENDING_ST pending
    class DWN_TECH roadmaps
    class DWN_REV revisao
```

---

## 5. Visão Consolidada: Entradas, Processo e Saídas (SIPOC)

```mermaid
flowchart LR
    classDef suppliers fill:#1a2a3a,stroke:#2a4a6a,color:#c8e0ff
    classDef inputs fill:#2a1a3a,stroke:#4a2a6a,color:#e0c8ff
    classDef process fill:#3a2a1a,stroke:#6a4a2a,color:#ffe0c8
    classDef outputs fill:#1a3a2a,stroke:#2a6a4a,color:#c8ffe0
    classDef customers fill:#2a3a1a,stroke:#4a6a2a,color:#e0ffc8

    subgraph SIPOC[" SIPOC: Downstream Architecture Refinement "]
        direction TB
        
        S["👥 SUPPLIERS<br/>━━━━━━━━━━<br/>Product Manager / PO<br/>Business Analyst<br/>Tech Lead / Arquiteto<br/>SQUAD de Desenvolvimento<br/>Security Architect<br/>Data Architect<br/>DevOps/SRE Architect<br/>QA/Test Specialist<br/>Infra/Cloud Specialist"]
        
        I["📥 INPUTS<br/>━━━━━━━━━━<br/>Project Charter + BRD + Épicos<br/>Features + User Stories (62 US)<br/>PROJECT_PATH / PROJECT_ID_NAME<br/>TECHNICAL_SOLUTION_PATH / NAMES<br/>ARCHITECTURE_GLOBAL<br/>SECURITY_GLOBAL<br/>🟢 PROJECT_PROMPT_INPUTS (Checkpoint HITL)<br/>📋 PROJECT-TEAM-SKILLS-MAP<br/>👥 PROJECT-TEAM-CAPACITY<br/>🔧 PROJECT-STACK<br/>📚 STACK-PADROES-CORPORATIVOS.md<br/>🔹 upstream-architecture-discovery/ (opcional)"]
        
        P["⚙️ PROCESS<br/>━━━━━━━━━━<br/>Fase 0: Bootstrap Inteligente<br/>Bloco A: 7 artefatos Detail-Level (F1-F7)<br/>  PRD + 6 disciplinas técnicas<br/>Bloco B: Estimativa PERT Independente (F8-F10)<br/>  ⭐ Three-point estimation por US<br/>  Resource Allocation + Risk Adjustment<br/>Bloco C: Scope Snapshot + Cross-Check (F11-F12)<br/>  📸 Foto imutável do escopo estimado<br/>  🔗 Comparação com ROM upstream (condicional)<br/>Ciclo HITL em cada fase (Generate→Gate→Fix)<br/>3 Barreiras de Bloco (A, B, C)<br/>Gate ESTIMATE-READY"]
        
        O["📤 OUTPUTS<br/>━━━━━━━━━━<br/>DETAIL-LEVEL-PRD.md (F1)<br/>DETAIL-LEVEL-ARCHITECTURE-DEFINITION.md (F2)<br/>DETAIL-LEVEL-SECURITY-DEFINITION.md (F3)<br/>DETAIL-LEVEL-DATA-ARCHITECTURE-DEFINITION.md (F4)<br/>DETAIL-LEVEL-DEVOPS-SRE-DEFINITION.md (F5)<br/>DETAIL-LEVEL-TEST-STRATEGY-DEFINITION.md (F6)<br/>DETAIL-LEVEL-INFRA-CLOUD-DEFINITION.md (F7)<br/>⭐ BOTTOM-UP-PERT-ESTIMATE.md (F8) ±15-25%<br/>RESOURCE-ALLOCATION-PLAN.md (F9)<br/>RISK-ADJUSTED-ESTIMATE.md (F10)<br/>📸 SCOPE-SNAPSHOT.md (F11)<br/>🔗 UPSTREAM-COMPARISON-REPORT.md (F12, condicional)"]
        
        C["🎯 CUSTOMERS<br/>━━━━━━━━━━<br/>Time de TI (validação de viabilidade)<br/>Comitê de Governança<br/>SQUAD de Desenvolvimento<br/>Tech Lead / Arquiteto<br/>Product Manager / PO<br/>Stakeholders do Projeto"]
    end
    
    S --> I --> P --> O --> C
    
    class S suppliers
    class I inputs
    class P process
    class O outputs
    class C customers
```

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 02/08/2026 | Criação inicial: 5 diagramas Mermaid cobrindo fluxo macro, ciclo HITL, barreiras, gate de estimativa e SIPOC. Baseado no roadmap v2.3 (Skills Map + Team Capacity + Stack Validation). Adaptado do FLOWCHART-UPSTREAM-ARCHITECTURE-DISCOVERY.md. | Time de Arquitetura |

---

🤖 *Diagramas gerados com Mermaid. Renderize em qualquer visualizador compatível (GitHub, Mermaid Live, VS Code).*
