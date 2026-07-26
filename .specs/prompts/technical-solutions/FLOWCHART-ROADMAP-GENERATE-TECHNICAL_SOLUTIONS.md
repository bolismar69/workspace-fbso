# FLOWCHART: ROADMAP DE EXECUÇÃO — SOLUÇÃO TÉCNICA

## Versão: 1.0 — Visualização Gráfica do Pipeline de Documentação Técnica

> **Documento de referência:** `PROMPT-ROADMAP-GENERATE-TECHNICAL_SOLUTIONS.md` v1.1
>
> Este documento complementa o roadmap textual com diagramas Mermaid que visualizam o fluxo de execução, os mecanismos de orquestração, as regras de cascata e a rastreabilidade vertical.

---

## 1. Visão Macro do Pipeline Completo

```mermaid
flowchart TB
    START([🚀 Início]) --> F0[Fase 0: Bootstrap Inteligente]
    F0 --> F1[Fase 1: TECHNICAL-SOLUTION-PRD.md]
    F1 --> F2[Fase 2: TECHNICAL-SOLUTION-ARCHITECTURE.md]
    F2 --> F3[Fase 3: TECHNICAL-SOLUTION-SECURITY.md]
    F3 --> F4[Fase 4: TECHNICAL-SOLUTION-SPECS.md]
    F4 --> F5[Fase 5: TECHNICAL-SOLUTION-TASKS.md]
    F5 --> F6[Fase 6: TECHNICAL-SOLUTION-TEST_PLAN.md]
    F6 --> END([✅ Solução Técnica Documentada])

    F0 -.->|inputs| GLOBAL[🏗️ ARCHITECTURE_GLOBAL<br/>🛡️ SECURITY_GLOBAL<br/>🛠️ STACK_DEFINITION]

    style F0 fill:#6c5ce7,color:#fff
    style F1 fill:#0984e3,color:#fff
    style F2 fill:#0984e3,color:#fff
    style F3 fill:#e17055,color:#fff
    style F4 fill:#0984e3,color:#fff
    style F5 fill:#0984e3,color:#fff
    style F6 fill:#00b894,color:#fff
    style GLOBAL fill:#dfe6e9,stroke:#636e72
```

---

## 2. Fase 0 — Bootstrap Inteligente (Detalhado)

```mermaid
flowchart TD
    F0_START([Fase 0: Bootstrap]) --> F0_1

    subgraph PASSO_01[Passo 0.1 — Coletar 7 Inputs Obrigatórios]
        F0_1[Solicitar PROJECT_PATH, PROJECT_ID_NAME,<br/>TECHNICAL_SOLUTION_PATH, TECHNICAL_SOLUTION_NAME,<br/>STACK, ARCHITECTURE_GLOBAL, SECURITY_GLOBAL] --> F0_1_STACK{Tipo de<br/>STACK?}
        F0_1_STACK -->|STACK_FILE_INPUT| F0_1_FILE[Ler arquivo de stack]
        F0_1_STACK -->|STACK_PROMPT_INPUT| F0_1_TEXT[Usar texto inline]
        F0_1_FILE --> F0_2
        F0_1_TEXT --> F0_2
    end

    subgraph PASSO_02[Passo 0.2 — Confirmar Caminhos]
        F0_2[Calcular variáveis derivadas:<br/>TECHNICAL_SOLUTION_COMPLETE_PATH_NAME<br/>PROJECT_COMPLETE_PATH_NAME<br/>SPECS_PATH<br/>STACK_DEFINITION] --> F0_2_DISP[Exibir todos os caminhos<br/>e stack para o usuário]
        F0_2_DISP --> F0_2_CONF{Usuário<br/>confirma?}
        F0_2_CONF -->|NÃO| F0_1
        F0_2_CONF -->|SIM| F0_3
    end

    subgraph PASSO_03[Passo 0.3 — Criar Estrutura]
        F0_3[mkdir -p SPECS_PATH<br/>Cria .specs/business-projects/ID_NAME/] --> F0_4
    end

    subgraph PASSO_04[Passo 0.4 — Auditar 6 Artefatos]
        F0_4[Verificar existência e compliance<br/>de PRD, ARCHITECTURE, SECURITY,<br/>SPECS, TASKS, TEST_PLAN] --> F0_4_DEC{Status dos<br/>artefatos?}
        F0_4_DEC -->|Todos ❌| F0_5A[🆕 Solução Nova<br/>→ Iniciar Fase 1]
        F0_4_DEC -->|Alguns ✅| F0_5B[📋 Solução em Andamento<br/>→ Iniciar da primeira fase pendente]
        F0_4_DEC -->|Todos ✅ Compliance| F0_5C[✅ Solução Completa<br/>→ Revisar / Evoluir / Encerrar]
    end

    subgraph PASSO_05[Passo 0.5 — Resumo]
        F0_5A --> F0_5[Exibir resumo final<br/>com stack, caminhos,<br/>próxima fase]
        F0_5B --> F0_5
        F0_5C --> F0_5
    end

    F0_5 --> ORCH

    style PASSO_01 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_02 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_03 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_04 fill:#dfe6e9,stroke:#6c5ce7
    style PASSO_05 fill:#dfe6e9,stroke:#6c5ce7
```

---

## 3. Mecanismo de Orquestração Dinâmica (Loop Trifásico)

Este é o coração do roadmap. **TODAS** as 6 fases executam este mesmo loop de validação.

```mermaid
flowchart TD
    ORCH([Orquestrador: Iniciar Fase N]) --> GEN

    subgraph LOOP[Loop de Validação Soberana — Fase N]
        GEN["1. GERAÇÃO / EVOLUÇÃO<br/>Executar PROMPT-GENERATE-{FASE}-TECHNICAL_SOLUTION.md<br/>Parâmetros: SOLUTION_PATH, PROJECT_PATH,<br/>STACK, SCOPE, + inputs da fase"] --> GATE

        GATE["2. AUDITORIA INTERNA DA IA<br/>Executar PROMPT-GATE-{FASE}-TECHNICAL_SOLUTION.md"] --> GATE_RESULT{Resultado<br/>da Auditoria?}

        GATE_RESULT -->|NÃO COMPLIANCE<br/>Erros encontrados| HUMAN_FB[Coletar feedback do humano<br/>para cada NC]
        HUMAN_FB --> FIX["2b. CORREÇÃO CIRÚRGICA<br/>Executar PROMPT-FIX-{FASE}-TECHNICAL_SOLUTION.md<br/>Apenas nas seções afetadas"]
        FIX --> GATE

        GATE_RESULT -->|SEM ERROS| HUMAN_GATE

        HUMAN_GATE[3. PORTÃO DE VALIDAÇÃO HUMANA<br/>Status: PRÉ-COMPLIANCE INTERNO<br/>— AGUARDANDO VALIDAÇÃO HUMANA] --> P1[P1: Aderente ao negócio<br/>e requisitos técnicos?]
        P1 --> P2[P2: Novos documentos<br/>de entrada?]
        P2 --> P3[P3: Novas informações,<br/>mudanças de escopo?]
        P3 --> HUMAN_DEC{Decisão<br/>do Humano?}

        HUMAN_DEC -->|Aprova sem novos inputs| COMPLIANCE
        HUMAN_DEC -->|Fornece novos docs/inputs| GEN
    end

    COMPLIANCE([✅ STATUS: COMPLIANCE<br/>Marcador gravado no cabeçalho<br/>Próxima fase destravada])

    style GEN fill:#0984e3,color:#fff
    style GATE fill:#fdcb6e,color:#333
    style HUMAN_FB fill:#ffeaa7,color:#333
    style FIX fill:#e17055,color:#fff
    style HUMAN_GATE fill:#6c5ce7,color:#fff
    style P1 fill:#dfe6e9,stroke:#6c5ce7
    style P2 fill:#dfe6e9,stroke:#6c5ce7
    style P3 fill:#dfe6e9,stroke:#6c5ce7
    style COMPLIANCE fill:#00b894,color:#fff
    style LOOP fill:#fff3e0,stroke:#f39c12
```

---

## 4. Fases 1-6 — Pipeline Sequencial com Dependências e Inputs

```mermaid
flowchart LR
    subgraph F1[Fase 1: TECHNICAL-SOLUTION-PRD.md]
        direction TB
        F1_GEN[GENERATE<br/>PRD] --> F1_GATE[GATE<br/>PRD] --> F1_FIX[FIX<br/>PRD]
        F1_FIX -.->|loop| F1_GATE
        F1_GATE --> F1_OK[COMPLIANCE ✅]
    end

    subgraph F2[Fase 2: TECHNICAL-SOLUTION-ARCHITECTURE.md]
        direction TB
        F2_GEN[GENERATE<br/>ARCH] --> F2_GATE[GATE<br/>ARCH] --> F2_FIX[FIX<br/>ARCH]
        F2_FIX -.->|loop| F2_GATE
        F2_GATE --> F2_OK[COMPLIANCE ✅]
    end

    subgraph F3[Fase 3: TECHNICAL-SOLUTION-SECURITY.md 🆕]
        direction TB
        F3_GEN[GENERATE<br/>SEC] --> F3_GATE[GATE<br/>SEC] --> F3_FIX[FIX<br/>SEC]
        F3_FIX -.->|loop| F3_GATE
        F3_GATE --> F3_OK[COMPLIANCE ✅]
    end

    subgraph F4[Fase 4: TECHNICAL-SOLUTION-SPECS.md]
        direction TB
        F4_GEN[GENERATE<br/>SPECS] --> F4_GATE[GATE<br/>SPECS] --> F4_FIX[FIX<br/>SPECS]
        F4_FIX -.->|loop| F4_GATE
        F4_GATE --> F4_OK[COMPLIANCE ✅]
    end

    subgraph F5[Fase 5: TECHNICAL-SOLUTION-TASKS.md]
        direction TB
        F5_GEN[GENERATE<br/>TASKS] --> F5_GATE[GATE<br/>TASKS] --> F5_FIX[FIX<br/>TASKS]
        F5_FIX -.->|loop| F5_GATE
        F5_GATE --> F5_OK[COMPLIANCE ✅]
    end

    subgraph F6[Fase 6: TECHNICAL-SOLUTION-TEST_PLAN.md]
        direction TB
        F6_GEN[GENERATE<br/>TEST] --> F6_GATE[GATE<br/>TEST] --> F6_FIX[FIX<br/>TEST]
        F6_FIX -.->|loop| F6_GATE
        F6_GATE --> F6_OK[COMPLIANCE ✅]
    end

    F1_OK -->|destrava| F2_GEN
    F2_OK -->|destrava| F3_GEN
    F3_OK -->|destrava| F4_GEN
    F4_OK -->|destrava| F5_GEN
    F5_OK -->|destrava| F6_GEN

    F1 --> I1[📥 Projeto de Negócio<br/>+ Docs adicionais]
    F2 --> I2[📥 TECHNICAL-SOLUTION-PRD.md +<br/>ARCHITECTURE_GLOBAL]
    F3 --> I3[📥 TECHNICAL-SOLUTION-PRD.md + TECHNICAL-SOLUTION-ARCHITECTURE.md<br/>+ SECURITY_GLOBAL]
    F4 --> I4[📥 TECHNICAL-SOLUTION-PRD.md + TECHNICAL-SOLUTION-ARCHITECTURE.md<br/>+ TECHNICAL-SOLUTION-SECURITY.md]
    F5 --> I5[📥 TECHNICAL-SOLUTION-PRD.md + TECHNICAL-SOLUTION-ARCHITECTURE.md<br/>+ TECHNICAL-SOLUTION-SPECS.md]
    F6 --> I6[📥 TECHNICAL-SOLUTION-PRD.md + TECHNICAL-SOLUTION-ARCHITECTURE.md<br/>+ TECHNICAL-SOLUTION-SPECS.md + TECHNICAL-SOLUTION-TASKS.md]

    style F1 fill:#0984e3,color:#fff
    style F2 fill:#0984e3,color:#fff
    style F3 fill:#e17055,color:#fff
    style F4 fill:#0984e3,color:#fff
    style F5 fill:#0984e3,color:#fff
    style F6 fill:#00b894,color:#fff
```

---

## 5. Efeitos Cascata (Cascade Rules)

Quando um artefato é modificado após COMPLIANCE, todos os artefatos downstream devem ser regenerados.

```mermaid
flowchart TD
    subgraph CASCADE[Regeneração em Cascata]
        direction TB

        C1[Se modificar TECHNICAL-SOLUTION-PRD.md] --> C1_DOWN[TECHNICAL-SOLUTION-ARCHITECTURE.md → TECHNICAL-SOLUTION-SECURITY.md →<br/>TECHNICAL-SOLUTION-SPECS.md → TECHNICAL-SOLUTION-TASKS.md → TECHNICAL-SOLUTION-TEST_PLAN.md]

        C2[Se modificar TECHNICAL-SOLUTION-ARCHITECTURE.md] --> C2_DOWN[TECHNICAL-SOLUTION-SECURITY.md → TECHNICAL-SOLUTION-SPECS.md →<br/>TECHNICAL-SOLUTION-TASKS.md → TECHNICAL-SOLUTION-TEST_PLAN.md]

        C3[Se modificar TECHNICAL-SOLUTION-SECURITY.md] --> C3_DOWN[TECHNICAL-SOLUTION-SPECS.md → TECHNICAL-SOLUTION-TASKS.md →<br/>TECHNICAL-SOLUTION-TEST_PLAN.md]

        C4[Se modificar TECHNICAL-SOLUTION-SPECS.md] --> C4_DOWN[TECHNICAL-SOLUTION-TASKS.md → TECHNICAL-SOLUTION-TEST_PLAN.md]

        C5[Se modificar TECHNICAL-SOLUTION-TASKS.md] --> C5_DOWN[TECHNICAL-SOLUTION-TEST_PLAN.md]
    end

    C1_DOWN --> ALERT
    C2_DOWN --> ALERT
    C3_DOWN --> ALERT
    C4_DOWN --> ALERT
    C5_DOWN --> ALERT

    ALERT[⚠️ Orquestrador alerta sobre<br/>efeito cascata e pergunta:] --> DECISION{Opção?}
    DECISION -->|A| REGEN[Regenerar TODOS<br/>os downstreams]
    DECISION -->|B| SIGNAL[Atualizar apenas o corrente<br/>+ sinalizar downstreams como<br/>'potencialmente desatualizados']

    style C1 fill:#d63031,color:#fff
    style C2 fill:#e17055,color:#fff
    style C3 fill:#fdcb6e,color:#333
    style C4 fill:#0984e3,color:#fff
    style C5 fill:#6c5ce7,color:#fff
    style ALERT fill:#ffeaa7,color:#333
```

### Matriz de Impacto Cascata

| Artefato Modificado | PRD | ARCH | SEC | SPECS | TASKS | TEST |
|---------------------|:---:|:---:|:---:|:---:|:---:|:---:|
| **TECHNICAL-SOLUTION-PRD.md** | — | 🔄 | 🔄 | 🔄 | 🔄 | 🔄 |
| **TECHNICAL-SOLUTION-ARCHITECTURE.md** | — | — | 🔄 | 🔄 | 🔄 | 🔄 |
| **TECHNICAL-SOLUTION-SECURITY.md** | — | — | — | 🔄 | 🔄 | 🔄 |
| **TECHNICAL-SOLUTION-SPECS.md** | — | — | — | — | 🔄 | 🔄 |
| **TECHNICAL-SOLUTION-TASKS.md** | — | — | — | — | — | 🔄 |
| **TECHNICAL-SOLUTION-TEST_PLAN.md** | — | — | — | — | — | — |

> 🔄 = Deve ser regenerado e revalidado

---

## 6. Matriz de Rastreabilidade — Validação Cruzada Fase 6 vs Fases 1-5

```mermaid
flowchart TD
    RTM_START([Antes de aprovar Fase 6: TECHNICAL-SOLUTION-TEST_PLAN.md]) --> RTM_1

    subgraph RTM[Auditoria de Rastreabilidade Vertical]
        RTM_1[1. MAPEAMENTO DE DEPENDÊNCIAS<br/>Cenário de Teste → Tarefa → Spec →<br/>Controle Sec → ADR → Req PRD →<br/>Objetivo de Negócio] --> RTM_2

        RTM_2[2. IDENTIFICAÇÃO DE ÓRFÃOS<br/>Algum cenário de teste sem<br/>requisito correspondente no PRD?] --> RTM_3

        RTM_3[3. VERIFICAÇÃO DE COBERTURA<br/>100% dos requisitos PRD + SPECS<br/>têm pelo menos 1 cenário de teste?] --> RTM_4

        RTM_4[4. VERIFICAÇÃO DE SEGURANÇA<br/>100% dos controles do TECHNICAL-SOLUTION-SECURITY.md<br/>têm cenários de teste no TEST_PLAN?]
    end

    RTM_4 --> RTM_RESULT{Relatório de<br/>Conformidade?}

    RTM_RESULT -->|✅ PASS| RTM_OK([✅ Compliance Confirmado<br/>Solução Técnica Completa])
    RTM_RESULT -->|❌ FAIL| RTM_FIX[🔧 Acionar FIX nos<br/>artefatos com defeito<br/>+ reportar gaps]
    RTM_FIX --> RTM_1

    style RTM fill:#fff3e0,stroke:#00b894
    style RTM_OK fill:#00b894,color:#fff
    style RTM_FIX fill:#d63031,color:#fff
```

---

## 7. Diagrama de Estados — Visão Unificada

```mermaid
stateDiagram-v2
    [*] --> Bootstrap: Início da Solução Técnica

    state Bootstrap {
        [*] --> ColetarInputs
        ColetarInputs --> ColetarStack: 7 inputs + STACK
        ColetarStack --> ValidarStack: Stack fornecida?
        ValidarStack --> ColetarInputs: NÃO (Stack é obrigatória)
        ValidarStack --> ConfirmarCaminhos: SIM
        ConfirmarCaminhos --> ColetarInputs: NÃO confirma
        ConfirmarCaminhos --> CriarEstrutura: SIM confirma
        CriarEstrutura --> AuditarArtefatos
        AuditarArtefatos --> ResumoInicial
    }

    Bootstrap --> Fase1_PRD

    state Fase1_PRD {
        [*] --> PRD_Gen
        PRD_Gen --> PRD_Gate
        PRD_Gate --> PRD_Fix: NÃO COMPLIANCE
        PRD_Fix --> PRD_Gate
        PRD_Gate --> PRD_Human: SEM ERROS
        PRD_Human --> PRD_Gen: Novos inputs
        PRD_Human --> PRD_Done: Aprovado → gravar [STATUS: COMPLIANCE]
    }

    Fase1_PRD --> Fase2_ARCH: COMPLIANCE

    state Fase2_ARCH {
        [*] --> ARCH_Gen
        ARCH_Gen --> ARCH_Gate
        ARCH_Gate --> ARCH_Fix: NÃO COMPLIANCE
        ARCH_Fix --> ARCH_Gate
        ARCH_Gate --> ARCH_Human: SEM ERROS
        ARCH_Human --> ARCH_Gen: Novos inputs
        ARCH_Human --> ARCH_Done: Aprovado → gravar [STATUS: COMPLIANCE]
    }

    Fase2_ARCH --> Fase3_SEC: COMPLIANCE

    state Fase3_SEC {
        [*] --> SEC_Gen
        SEC_Gen --> SEC_Gate
        SEC_Gate --> SEC_Fix: NÃO COMPLIANCE
        SEC_Fix --> SEC_Gate
        SEC_Gate --> SEC_Human: SEM ERROS
        SEC_Human --> SEC_Gen: Novos inputs
        SEC_Human --> SEC_Done: Aprovado → gravar [STATUS: COMPLIANCE]
    }

    Fase3_SEC --> Fase4_SPECS: COMPLIANCE

    state Fase4_SPECS {
        [*] --> SPECS_Gen
        SPECS_Gen --> SPECS_Gate
        SPECS_Gate --> SPECS_Fix: NÃO COMPLIANCE
        SPECS_Fix --> SPECS_Gate
        SPECS_Gate --> SPECS_Human: SEM ERROS
        SPECS_Human --> SPECS_Gen: Novos inputs
        SPECS_Human --> SPECS_Done: Aprovado → gravar [STATUS: COMPLIANCE]
    }

    Fase4_SPECS --> Fase5_TASKS: COMPLIANCE

    state Fase5_TASKS {
        [*] --> TASKS_Gen
        TASKS_Gen --> TASKS_Gate
        TASKS_Gate --> TASKS_Fix: NÃO COMPLIANCE
        TASKS_Fix --> TASKS_Gate
        TASKS_Gate --> TASKS_Human: SEM ERROS
        TASKS_Human --> TASKS_Gen: Novos inputs
        TASKS_Human --> TASKS_Done: Aprovado → gravar [STATUS: COMPLIANCE]
    }

    Fase5_TASKS --> Fase6_TEST: COMPLIANCE

    state Fase6_TEST {
        [*] --> TEST_Gen
        TEST_Gen --> TEST_Gate
        TEST_Gate --> TEST_Fix: NÃO COMPLIANCE
        TEST_Fix --> TEST_Gate
        TEST_Gate --> RTM_Check: SEM ERROS
        RTM_Check --> TEST_Fix: FAIL na matriz
        RTM_Check --> TEST_Human: PASS na matriz
        TEST_Human --> TEST_Gen: Novos inputs
        TEST_Human --> TEST_Done: Aprovado → gravar [STATUS: COMPLIANCE]
    }

    Fase6_TEST --> [*]: ✅ Solução Técnica Completa
```

---

## 8. Fluxo de Decisão do Bootstrap — Detalhamento Lógico

```mermaid
flowchart TD
    subgraph BOOTSTRAP_LOGIC[Lógica de Decisão do Passo 0.4]
        CHECK_ALL[Verificar 6 arquivos em SPECS_PATH]

        ALL_NO[❌ Todos ausentes<br/>→ SOLUÇÃO NOVA]
        SOME_YES[⚠️ Alguns existem<br/>→ SOLUÇÃO EM ANDAMENTO]
        ALL_YES[✅ Todos existem<br/>+ COMPLIANCE]

        ALL_NO --> NEW[Iniciar Fase 1: TECHNICAL-SOLUTION-PRD.md]

        SOME_YES --> FIND[Encontrar primeiro artefato<br/>ausente ou não-Compliance<br/>na ordem 1→6]
        FIND --> REPORT[Reportar status completo:<br/>'Iniciando da Fase N — Nome<br/>(primeiro artefato pendente)']

        ALL_YES --> ASK[Perguntar: Revisar /<br/>Novo ciclo evolutivo /<br/>Encerrar?]
        ASK -->|Revisar| FIND
        ASK -->|Novo ciclo| NEW
        ASK -->|Encerrar| DONE
    end

    REPORT --> ORCH

    style ALL_NO fill:#d63031,color:#fff
    style SOME_YES fill:#fdcb6e,color:#333
    style ALL_YES fill:#00b894,color:#fff
    style NEW fill:#0984e3,color:#fff
    style DONE fill:#00b894,color:#fff
```

---

## 9. Parâmetros por Fase — Resumo Visual

```mermaid
flowchart LR
    subgraph PARAMS[Parâmetros Injetados em Cada Fase]
        direction TB

        P1[Fase 1: PRD] --> P1_PARAMS["SOLUTION_PATH<br/>PROJECT_PATH<br/>PROJECT_NAME<br/>SOLUTION_NAME<br/>STACK<br/>SCOPE=full"]
        P2[Fase 2: ARCH] --> P2_PARAMS["SOLUTION_PATH<br/>PROJECT_PATH<br/>PROJECT_NAME<br/>SOLUTION_NAME<br/>STACK"]
        P3[Fase 3: SEC 🆕] --> P3_PARAMS["SOLUTION_PATH<br/>PROJECT_PATH<br/>PROJECT_NAME<br/>SOLUTION_NAME<br/>STACK<br/>SECURITY_GLOBAL"]
        P4[Fase 4: SPECS] --> P4_PARAMS["SOLUTION_PATH<br/>PROJECT_PATH<br/>PROJECT_NAME<br/>SOLUTION_NAME<br/>SOLUTION_TYPE=backend<br/>SCOPE=full"]
        P5[Fase 5: TASKS] --> P5_PARAMS["SOLUTION_PATH<br/>PROJECT_PATH<br/>PROJECT_NAME<br/>SOLUTION_NAME<br/>SCOPE=full"]
        P6[Fase 6: TEST] --> P6_PARAMS["SOLUTION_PATH<br/>PROJECT_PATH<br/>PROJECT_NAME<br/>SOLUTION_NAME<br/>STACK<br/>SCOPE=full"]
    end

    style P1 fill:#0984e3,color:#fff
    style P2 fill:#0984e3,color:#fff
    style P3 fill:#e17055,color:#fff
    style P4 fill:#0984e3,color:#fff
    style P5 fill:#0984e3,color:#fff
    style P6 fill:#00b894,color:#fff
```

---

## 10. Tabela de Símbolos e Convenções

| Símbolo/Cor | Significado |
|-------------|-------------|
| 🟣 Roxo (`#6c5ce7`) | Bootstrap / Validação Humana |
| 🔵 Azul (`#0984e3`) | Fases de Geração (1, 2, 4, 5) |
| 🟠 Laranja (`#e17055`) | Fase 3 (TECHNICAL-SOLUTION-SECURITY.md) / Correções |
| 🟢 Verde (`#00b894`) | Compliance / Fase 6 (TEST_PLAN) |
| 🟡 Amarelo (`#fdcb6e`) | Auditoria Interna (Gate) |
| 🔴 Vermelho (`#d63031`) | Erro / Falha / Bloqueio |
| 🔲 Linha tracejada | Loop de retrabalho |
| 🔲 Linha sólida | Fluxo sequencial normal |
| 🔄 | Regenerar em cascata |

---

## 11. Diferenças Chave: Roadmap de Negócio vs Roadmap Técnico

| Aspecto | Roadmap de Negócio | Roadmap Técnico |
|---------|-------------------|-----------------|
| **Fases** | 5 fases (Charter → BRD → Epics → Features → US) | 6 fases (PRD → ARCH → SEC → SPECS → TASKS → TEST) |
| **Fase de Segurança** | Integrada como NFR nos artefatos | Artefato dedicado (TECHNICAL-SOLUTION-SECURITY.md, Fase 3) |
| **Bootstrap** | 6 variáveis (com PROMPT_BRANCH) | 7+2 variáveis (com STACK e globais de arquitetura/segurança) |
| **Git Workflow** | Automatizado ao final (commit → push → PR → merge) | Não incluso (específico do projeto) |
| **Efeitos Cascata** | Revisão manual de downstreams | Regeneração automática com opção A/B |
| **RTM** | US → Feature → Epic → BRD → Charter | Teste → Tarefa → Spec → Sec → ADR → PRD → Negócio |
| **Nível** | Estratégico / Negócio | Tático / Implementação |
| **Output** | Pasta `user-stories/` + RTM central | 6 arquivos mestre em `SPECS_PATH` |

---

> **📁 Arquivos relacionados:**
> - `PROMPT-ROADMAP-GENERATE-TECHNICAL_SOLUTION.md` — Documento fonte (v1.0)
> - `PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS.md` — Roadmap de documentos de negócio (v5.0)
> - `FLOWCHART-ROADMAP-GENERATE-PROJECT-DOCUMENTS.md` — Visualização do roadmap de negócio
> - `PROMPT-ORCHESTRATOR-GENERATE-ALL-ARTEFACTS.md` — Orquestrador de geração
> - `FLUXO-SPEC-DRIVEN-DEVELOPMENT-V1.md` — Fluxo spec-driven complementar
