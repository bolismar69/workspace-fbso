# FLOWCHART: ROADMAP DE EXECUÇÃO MACRO E GUIA DE ORQUESTRAÇÃO DE DOCUMENTOS

## Versão: 5.0 — Visualização Gráfica do Pipeline de Documentação

> **Documento de referência:** `PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS.md` v5.0
>
> Este documento complementa o roadmap textual com diagramas Mermaid que visualizam o fluxo de execução, os mecanismos de orquestração e as regras de gating.

---

## 1. Visão Macro do Pipeline Completo

```mermaid
flowchart TB
    START([🚀 Início]) --> F0[Fase 0: Bootstrap Inteligente]
    F0 --> F1[Fase 1: Project Charter]
    F1 --> F2[Fase 2: BRD]
    F2 --> F3[Fase 3: Epics]
    F3 --> F4[Fase 4: Features]
    F4 --> F5[Fase 5: User Stories + RTM]
    F5 --> GIT[Pipeline Git: Commit → Push → PR → Merge]
    GIT --> END([🎉 Processo Finalizado])

    style F0 fill:#6c5ce7,color:#fff
    style F1 fill:#0984e3,color:#fff
    style F2 fill:#0984e3,color:#fff
    style F3 fill:#0984e3,color:#fff
    style F4 fill:#0984e3,color:#fff
    style F5 fill:#e17055,color:#fff
    style GIT fill:#00b894,color:#fff
```

---

## 2. Fase 0 — Bootstrap Inteligente (Detalhado)

```mermaid
flowchart TD
    F0_START([Fase 0: Bootstrap]) --> F0_1

    subgraph PASSO_01[Passo 0.1 — Coletar Inputs]
        F0_1[Solicitar 6 variáveis ao usuário] --> F0_1_VAL{PROMPT_BRANCH<br/>válida?}
        F0_1_VAL -->|Sim| F0_2
        F0_1_VAL -->|Não main/master/develop| F0_1_ERR[⛔ Exibir mensagem de bloqueio<br/>e solicitar novamente]
        F0_1_ERR --> F0_1
    end

    subgraph PASSO_02[Passo 0.2 — Confirmar Caminho]
        F0_2[Calcular PROJECT_COMPLETE_PATH_NAME<br/>e PROJECT_ID_NAME] --> F0_2_CONF{Usuário<br/>confirma?}
        F0_2_CONF -->|SIM| F0_3
        F0_2_CONF -->|NÃO| F0_1
    end

    subgraph PASSO_03[Passo 0.3 — Criar Estrutura]
        F0_3[mkdir -p PROJECT_COMPLETE_PATH_NAME/user-stories/] --> F0_4
    end

    subgraph PASSO_04[Passo 0.4 — Auditar Artefatos]
        F0_4[Verificar existência de 6 artefatos] --> F0_4_DEC{Status dos<br/>artefatos?}
        F0_4_DEC -->|Todos ❌| F0_5A[Projeto Novo → Iniciar Fase 1]
        F0_4_DEC -->|Alguns ✅| F0_5B[Projeto em Andamento →<br/>Perguntar fase de retomada]
        F0_4_DEC -->|Todos ✅| F0_5C[Projeto Completo →<br/>Perguntar se deseja revisão]
    end

    subgraph PASSO_05[Passo 0.5 — Resumo]
        F0_5A --> F0_5[Exibir resumo final<br/>e iniciar próxima fase]
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

## 3. Mecanismo de Orquestração Dinâmica (Loop Trifásico por Fase)

Este é o coração do roadmap. **TODAS** as fases 1-5 executam este mesmo loop.

```mermaid
flowchart TD
    ORCH([Orquestrador: Iniciar Fase N]) --> GEN

    subgraph LOOP[Loop de Validação Soberana — Fase N]
        GEN[1. GERAÇÃO<br/>Executar PROMPT-GENERATE-FASE-N.md] --> GATE

        GATE[2. AUDITORIA INTERNA<br/>Executar PROMPT-GATE-FASE-N.md] --> GATE_RESULT{Resultado<br/>da Auditoria?}

        GATE_RESULT -->|NÃO COMPLIANCE| FIX[2b. CORREÇÃO CIRÚRGICA<br/>Executar PROMPT-FIX-FASE-N.md<br/>apenas no artefato com defeito]
        FIX --> GATE

        GATE_RESULT -->|SEM ERROS| HUMAN_GATE

        HUMAN_GATE[3. PORTÃO DE VALIDAÇÃO HUMANA<br/>Status: PRÉ-COMPLIANCE INTERNO<br/>3 perguntas obrigatórias] --> HUMAN_DEC{Decisão<br/>do Humano?}

        HUMAN_DEC -->|Aprova sem novos inputs| COMPLIANCE
        HUMAN_DEC -->|Fornece novos documentos/inputs| GEN
    end

    COMPLIANCE([✅ STATUS: COMPLIANCE<br/>Arquivo congelado<br/>Próxima fase destravada])

    style GEN fill:#0984e3,color:#fff
    style GATE fill:#fdcb6e,color:#333
    style FIX fill:#e17055,color:#fff
    style HUMAN_GATE fill:#6c5ce7,color:#fff
    style COMPLIANCE fill:#00b894,color:#fff
    style LOOP fill:#fff3e0,stroke:#f39c12
```

### As 3 Perguntas do Portão de Validação Humana

| # | Pergunta | Propósito |
|---|----------|-----------|
| 1 | O documento está aderente às necessidades do negócio? | Validação de conteúdo |
| 2 | Você deseja anexar novos documentos de entrada? | Detecção de evolução incremental |
| 3 | Você deseja fornecer novos inputs textuais? | Captura de requisitos adicionais |

---

## 4. Fases 1-5 — Pipeline Sequencial com Dependências

```mermaid
flowchart LR
    subgraph F1[Fase 1: Project Charter]
        F1_GEN[GENERATE] --> F1_GATE[GATE] --> F1_FIX[FIX]
        F1_FIX -.->|loop| F1_GATE
        F1_GATE --> F1_OK[COMPLIANCE ✅]
    end

    subgraph F2[Fase 2: BRD]
        F2_GEN[GENERATE] --> F2_GATE[GATE] --> F2_FIX[FIX]
        F2_FIX -.->|loop| F2_GATE
        F2_GATE --> F2_OK[COMPLIANCE ✅]
    end

    subgraph F3[Fase 3: Epics]
        F3_GEN[GENERATE] --> F3_GATE[GATE] --> F3_FIX[FIX]
        F3_FIX -.->|loop| F3_GATE
        F3_GATE --> F3_OK[COMPLIANCE ✅]
    end

    subgraph F4[Fase 4: Features]
        F4_GEN[GENERATE] --> F4_GATE[GATE] --> F4_FIX[FIX]
        F4_FIX -.->|loop| F4_GATE
        F4_GATE --> F4_OK[COMPLIANCE ✅]
    end

    subgraph F5[Fase 5: User Stories + RTM]
        F5_GEN[GENERATE<br/>Modular] --> F5_GATE[GATE<br/>Valida arquivos<br/>+ links RTM] --> F5_FIX[FIX<br/>Isolado por<br/>arquivo]
        F5_FIX -.->|loop| F5_GATE
        F5_GATE --> F5_OK[COMPLIANCE ✅]
    end

    F1_OK -->|destrava| F2_GEN
    F2_OK -->|destrava| F3_GEN
    F3_OK -->|destrava| F4_GEN
    F4_OK -->|destrava| F5_GEN

    F1 --> INPUTS1[Inputs: docs brutos<br/>+ entrevistas]
    F2 --> INPUTS2[Inputs: 01-Charter<br/>+ stakeholders]
    F3 --> INPUTS3[Inputs: 02-BRD]
    F4 --> INPUTS4[Inputs: 03-Epics]
    F5 --> INPUTS5[Inputs: 04-Features<br/>+ 01-Charter original]

    style F1 fill:#0984e3,color:#fff
    style F2 fill:#0984e3,color:#fff
    style F3 fill:#0984e3,color:#fff
    style F4 fill:#0984e3,color:#fff
    style F5 fill:#e17055,color:#fff
```

### Artefatos Produzidos por Fase

| Fase | Arquivo Gerado | Conteúdo |
|------|----------------|----------|
| 1 | `01-PROJECT-CHARTER-{ID}.md` | 14 seções macro: escopo, objetivos, premissas, restrições, governança |
| 2 | `02-BRD-{ID}.md` | Requisitos de negócio detalhados, regras de atendimento |
| 3 | `03-EPICS-{ID}.md` | Grandes blocos de entrega de valor |
| 4 | `04-FEATURES-{ID}.md` | Funcionalidades tangíveis e implementáveis |
| 5 | `USER-STORIES-{PROJECT_ID_NAME}.md` + `user-stories/US-*.md` | Índice central + arquivos atômicos com Gherkin |

---

## 5. Matriz de Rastreabilidade (RTM) — Validação Cruzada Fase 5 vs Fase 1

```mermaid
flowchart TD
    RTM_START([Antes de aprovar Fase 5]) --> RTM_1

    subgraph RTM[Auditoria de Rastreabilidade Bidirecional]
        RTM_1["1. Mapeamento de Dependência:<br/>US → Feature → Epic → BRD → Charter"] --> RTM_2
        RTM_2["2. Identificação de Órfãos:<br/>Alguma US sem objetivo<br/>correspondente no Charter?"] --> RTM_3
        RTM_3["3. Verificação de Cobertura:<br/>100% dos Objetivos do Charter<br/>têm pelo menos 1 US?"] --> RTM_4
        RTM_4["4. Auto-Análise:<br/>Comparar user-stories/*.md + RTM.md<br/>contra 01-PROJECT-CHARTER"]
    end

    RTM_4 --> RTM_RESULT{Relatório de<br/>Conformidade?}

    RTM_RESULT -->|PASS| RTM_OK([✅ Compliance Confirmado])
    RTM_RESULT -->|FAIL| RTM_FIX[🔧 Acionar FIX nos<br/>arquivos com defeito]
    RTM_FIX --> RTM_1

    style RTM fill:#fff3e0,stroke:#e17055
    style RTM_OK fill:#00b894,color:#fff
    style RTM_FIX fill:#d63031,color:#fff
```

---

## 6. Pipeline Git — Finalização Automatizada

```mermaid
flowchart TD
    GIT_START([Usuário indica conclusão<br/>das 5 fases]) --> PRE_CHECK

    subgraph PRE[Pré-condições]
        PRE_CHECK[Validar: PROMPT_BRANCH ok?<br/>Repositório Git?<br/>Alterações pendentes?]
    end

    PRE_CHECK --> F1

    subgraph GIT_PIPE[Pipeline Git — 4 Passos]
        F1[Passo F.1: git add -A<br/>git commit -m ...] --> F1_CHECK{Commit<br/>OK?}
        F1_CHECK -->|Nada a commitar| F4
        F1_CHECK -->|Sucesso| F2

        F2[Passo F.2: git push] --> F2_CHECK{Push<br/>OK?}
        F2_CHECK -->|Branch remota existe| F2_ASK{Usuário<br/>autoriza<br/>--force?}
        F2_ASK -->|SIM| F2_FORCE[git push --force]
        F2_ASK -->|NÃO| ABORT([⛔ Pipeline abortado])
        F2_CHECK -->|Sucesso| F3
        F2_FORCE --> F3

        F3[Passo F.3: gh pr create<br/>gh pr merge --merge --delete-branch] --> F3_CHECK{Merge<br/>OK?}
        F3_CHECK -->|Sucesso| F4
        F3_CHECK -->|Conflitos| ABORT_CONF([⚠️ Conflitos!<br/>Branch NÃO deletada<br/>Resolução manual])

        F4[Passo F.4: git checkout main<br/>git branch -d PROMPT_BRANCH] --> RESUMO
    end

    RESUMO[📊 Exibir resumo final<br/>com status de cada etapa] --> END([🎉 Processo Finalizado])

    style PRE fill:#dfe6e9,stroke:#636e72
    style GIT_PIPE fill:#e8f8f5,stroke:#00b894
    style END fill:#00b894,color:#fff
    style ABORT fill:#d63031,color:#fff
    style ABORT_CONF fill:#e17055,color:#fff
```

---

## 7. Diagrama de Estados — Visão Unificada

```mermaid
stateDiagram-v2
    [*] --> Bootstrap: Início

    state Bootstrap {
        [*] --> ColetarInputs
        ColetarInputs --> ValidarBranch: PROMPT_BRANCH
        ValidarBranch --> ColetarInputs: main/master/develop
        ValidarBranch --> ConfirmarCaminho: branch válida
        ConfirmarCaminho --> ColetarInputs: NÃO confirma
        ConfirmarCaminho --> CriarEstrutura: SIM confirma
        CriarEstrutura --> AuditarArtefatos
        AuditarArtefatos --> ResumoInicial
    }

    Bootstrap --> Fase1_Charter

    state Fase1_Charter {
        [*] --> Charter_Gen
        Charter_Gen --> Charter_Gate
        Charter_Gate --> Charter_Fix: NÃO COMPLIANCE
        Charter_Fix --> Charter_Gate
        Charter_Gate --> Charter_Human: SEM ERROS
        Charter_Human --> Charter_Gen: Novos inputs
        Charter_Human --> Charter_Done: Aprovado
    }

    Fase1_Charter --> Fase2_BRD: COMPLIANCE

    state Fase2_BRD {
        [*] --> BRD_Gen
        BRD_Gen --> BRD_Gate
        BRD_Gate --> BRD_Fix: NÃO COMPLIANCE
        BRD_Fix --> BRD_Gate
        BRD_Gate --> BRD_Human: SEM ERROS
        BRD_Human --> BRD_Gen: Novos inputs
        BRD_Human --> BRD_Done: Aprovado
    }

    Fase2_BRD --> Fase3_Epics: COMPLIANCE

    state Fase3_Epics {
        [*] --> Epics_Gen
        Epics_Gen --> Epics_Gate
        Epics_Gate --> Epics_Fix: NÃO COMPLIANCE
        Epics_Fix --> Epics_Gate
        Epics_Gate --> Epics_Human: SEM ERROS
        Epics_Human --> Epics_Gen: Novos inputs
        Epics_Human --> Epics_Done: Aprovado
    }

    Fase3_Epics --> Fase4_Features: COMPLIANCE

    state Fase4_Features {
        [*] --> Features_Gen
        Features_Gen --> Features_Gate
        Features_Gate --> Features_Fix: NÃO COMPLIANCE
        Features_Fix --> Features_Gate
        Features_Gate --> Features_Human: SEM ERROS
        Features_Human --> Features_Gen: Novos inputs
        Features_Human --> Features_Done: Aprovado
    }

    Fase4_Features --> Fase5_UserStories: COMPLIANCE

    state Fase5_UserStories {
        [*] --> US_Gen
        US_Gen --> US_Gate
        US_Gate --> US_Fix: NÃO COMPLIANCE
        US_Fix --> US_Gate
        US_Gate --> RTM_Check: SEM ERROS
        RTM_Check --> US_Fix: FAIL na matriz
        RTM_Check --> US_Human: PASS na matriz
        US_Human --> US_Gen: Novos inputs
        US_Human --> US_Done: Aprovado
    }

    Fase5_UserStories --> GitWorkflow: Todas fases OK

    state GitWorkflow {
        [*] --> Commit
        Commit --> Push: Sucesso
        Commit --> Cleanup: Nada a commitar
        Push --> PR_Merge: Sucesso
        Push --> ForceDecision: Branch existe
        ForceDecision --> PR_Merge: --force autorizado
        ForceDecision --> [*]: Abortado
        PR_Merge --> Cleanup: Merge OK
        PR_Merge --> [*]: Conflitos
        Cleanup --> ResumoFinal
    }

    GitWorkflow --> [*]: 🎉 Finalizado
```

---

## 8. Tabela de Símbolos e Convenções

| Símbolo/Cor | Significado |
|-------------|-------------|
| 🟣 Roxo (`#6c5ce7`) | Bootstrap / Validação Humana |
| 🔵 Azul (`#0984e3`) | Fases de Geração (1-4) |
| 🟠 Laranja (`#e17055`) | Fase 5 (User Stories) / Correções |
| 🟢 Verde (`#00b894`) | Compliance / Pipeline Git |
| 🟡 Amarelo (`#fdcb6e`) | Auditoria Interna (Gate) |
| 🔴 Vermelho (`#d63031`) | Erro / Abort |
| 🔲 Linha tracejada | Loop de retrabalho |
| 🔲 Linha sólida | Fluxo sequencial normal |

---

## 9. Regras de Gating (Resumo Visual)

```mermaid
flowchart LR
    subgraph GATING[Regras Críticas de Bloqueio]
        G1[⛔ PROMPT_BRANCH<br/>não pode ser<br/>main/master/develop]
        G2[⛔ Nenhuma fase<br/>avança sem<br/>COMPLIANCE humano]
        G3[⛔ Branch local<br/>NUNCA deletada<br/>se merge falhar]
        G4[⛔ Fase 5 exige<br/>RTM 100%<br/>aprovada]
    end

    style G1 fill:#d63031,color:#fff
    style G2 fill:#d63031,color:#fff
    style G3 fill:#d63031,color:#fff
    style G4 fill:#d63031,color:#fff
```

---

> **📁 Arquivos relacionados:**
> - `PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS.md` — Documento fonte (v5.0)
> - `PROMPT-ORCHESTRATOR-GENERATE-ALL-ARTEFACTS.md` — Orquestrador de geração
> - `FLUXO-SPEC-DRIVEN-DEVELOPMENT-V1.md` — Fluxo spec-driven complementar
