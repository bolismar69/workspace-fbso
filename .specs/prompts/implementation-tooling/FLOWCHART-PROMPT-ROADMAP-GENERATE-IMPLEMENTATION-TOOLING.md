# FLOWCHART: ROADMAP DE IMPLEMENTAÇÃO DE AMBIENTE E FERRAMENTAS

## Versão: 1.0 — 4 Fases, 4 Trios (610/620/630/640), Barreiras e Gates HITL

> **Documento de referência:** `PROMPT-ROADMAP-GENERATE-IMPLEMENTATION-TOOLING.md` v1.0
>
> Este documento complementa o roadmap textual com diagramas Mermaid que visualizam o fluxo de execução, a arquitetura de fases com barreiras, o loop GENERATE→GATE→FIX por artefato e a integração com os roadmaps TECHLEAD/WATERFALL.

---

## 1. Visão Macro do Roadmap

```mermaid
flowchart TB
    START(["🚀 Início"]) --> F0["FASE 0: Bootstrap e Auditoria"]
    F0 --> BE{{"⛔ Barreira de Entrada<br/>docs-base em COMPLIANCE?"}}
    BE -->|NÃO| STOP["⛔ PARE: destrave o gate do doc-base"]
    BE -->|SIM| F1B

    subgraph F1B["F1 — Manifestos DevOps (610)"]
        F1["610-MANIFESTS-DEVOPS<br/>Dockerfile, Helm, K8s, Terraform"] --> F1G["GATE-610"]
        F1G -.->|violações| F1X["FIX-610 (cirúrgico)"]
        F1X -.-> F1G
    end

    F1B --> B1{{"🚧 Barreira 1<br/>manifests COMPLIANCE"}}
    B1 --> F2B

    subgraph F2B["F2 — Observabilidade (620)"]
        F2["620-OBSERVABILITY-SETUP<br/>stack, dashboards SLO, alertas"] --> F2G["GATE-620"]
        F2G -.->|violações| F2X["FIX-620 (cirúrgico)"]
        F2X -.-> F2G
    end

    F2B --> B2{{"🚧 Barreira 2<br/>observabilidade COMPLIANCE"}}
    B2 --> F3B

    subgraph F3B["F3 — Ferramentas middleware/ETL (630)"]
        F3["630-INSTALL-TOOL-FERRAMENTA<br/>uma execução por ferramenta"] --> F3G["GATE-630"]
        F3G -.->|violações| F3X["FIX-630 (cirúrgico)"]
        F3X -.-> F3G
        F3G -->|"PASS + próxima ferramenta"| F3
    end

    F3B --> B3{{"🚧 Barreira 3<br/>ferramentas COMPLIANCE"}}
    B3 --> F4B

    subgraph F4B["F4 — Ferramentas de segurança (640)"]
        F4["640-INSTALL-SECURITY-TOOL-FERRAMENTA<br/>uma execução por ferramenta"] --> F4G["GATE-640"]
        F4G -.->|violações| F4X["FIX-640 (cirúrgico)"]
        F4X -.-> F4G
    end

    F4B --> FIM["🚩 CONCLUSÃO<br/>relatório consolidado + 095/100"]
```

---

## 2. Loop GENERATE → GATE → FIX por Artefato

```mermaid
flowchart LR
    G["GENERATE<br/>[STATUS: Em análise]"] --> GA["GATE<br/>checklist vs docs-base"]
    GA -->|"PASS"| HR["Validação humana<br/>(ancoragem, completude,<br/>segurança, vocabulário)"]
    HR -->|aprova| C["[STATUS: COMPLIANCE]<br/>próximo artefato"]
    HR -->|reprova| GA
    GA -->|"FAIL<br/>VIOLATIONS[]"| X["FIX (cirúrgico)<br/>apenas seções violadas"]
    X --> GA
```

---

## 3. Integração com os Roadmaps Existentes

```mermaid
flowchart TB
    PMPO["WATERFALL-EXECUTION (FASE 5)<br/>janelas DEV/QA do 096"] --> TL["TECHLEAD v7.x — Bloco F<br/>tarefas de infra/ferramentas"]
    TL --> IT["IMPLEMENTATION-TOOLING v1.0<br/>610 / 620 / 630 / 640"]
    IT -->|manifests e ferramentas| BL["Bloco E do TECHLEAD<br/>EXECUTE-SPRINT-TASKS,<br/>EXECUTE-CI-CD-PIPELINE,<br/>EXECUTE-CVE-SCA-SCAN"]
    BL -->|evidências| PMPO
    IT -.->|runbooks e evidências| DOCS["095-RELATORIO-QUALIDADE<br/>100-MANUAIS-OPERACIONAIS"]
```

> **Legenda:** `{{⛔}}` = barreira estrutural com decisão; `-.->` = fluxo condicional/de correção; em contexto WATERFALL, HMG/PROD via GMUD (090); em contexto ágil, approval gates do pipeline — sempre com aprovação humana por ambiente.
