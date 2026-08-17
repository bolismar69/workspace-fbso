# PROMPT: ROADMAP DE IMPLEMENTAÇÃO DE AMBIENTE E FERRAMENTAS — TOOLING DA FASE 5 (IMPLEMENTATION-TOOLING)
## Versão: 1.0 — 4 Fases, 4 Trios (12 Prompts) + FLOWCHART, Numeração por Intervalos 6xx, Invocado pelo TECHLEAD (Bloco F) ou standalone — Independente da Metodologia (ágil ou waterfall)

Atue como um Arquiteto DevOps/Plataforma Sênior e Gestor de Execução, especializado em orquestração de infraestrutura como código e engenharia de prompts, atuando de forma independente da metodologia do projeto (ágil ou waterfall).

## Propósito e Posição no Fluxo

Este roadmap orquestra a **implementação de ambiente e ferramentas** da FASE 5 — atividades de setup que **não pertencem ao ciclo de código** (código é o Bloco E do TECHLEAD): geração de manifestos DevOps, implantação da stack de observabilidade, instalação de ferramentas middleware/ETL/orquestração e instalação de ferramentas de segurança.

Ele materializa em artefatos reais (Dockerfile, Helm charts, K8s YAML, Terraform, relatórios de instalação) as definições produzidas pelo TECHLEAD (`480/490/500/510/520/550`) e pelos setups WATERFALL (`041/043/044`).

```
FASE 0 (Bootstrap + Barreira de Entrada) → F1 (610 MANIFESTS-DEVOPS) → 🚧 Barreira 1
    → F2 (620 OBSERVABILITY-SETUP) → 🚧 Barreira 2
    → F3 (630 INSTALL-TOOL por ferramenta) → 🚧 Barreira 3
    → F4 (640 INSTALL-SECURITY-TOOL por ferramenta) → 🚩 CONCLUSÃO (relatório consolidado)
```

- **Entrada:** definições TECHLEAD Bloco B/C em `[STATUS: COMPLIANCE]` + setups WATERFALL `041/043/044` em `[STATUS: COMPLIANCE]` (M4 travado)
- **Saída:** manifestos e ferramentas implementadas/validadas por ambiente (DEV/QA/HMG/PROD conforme 087/096), com relatórios rastreáveis aos docs-base
- **Quem invoca:** `PROMPT-ROADMAP-GENERATE-PROJECT-TECHNICAL-DEFINITIONS.md` (Bloco F — tarefas de infra/ferramentas das janelas DEV/QA) ou `PROMPT-ROADMAP-GENERATE-WATERFALL-EXECUTION.md` (FASE 5, setup de ambiente) — também pode rodar standalone com os docs-base em mãos

## Regras de Ouro (7 REGRAS DE GATING — NÃO NEGOCIÁVEIS)

1. **NÃO ADIVINHAR INPUTS:** nenhum prompt infere, busca ou descobre inputs — todo parâmetro é passado explicitamente pelo orquestrador.
2. **TODO GENERATE TEM GATE+FIX:** cada um dos 4 artefatos tem exatamente um trio GENERATE, GATE e FIX.
3. **FLUXO DE DADOS EXPLÍCITO:** `ARTIFACT_PATH` flui GENERATE→GATE→FIX. `VIOLATIONS[]` flui GATE→FIX. Sempre como parâmetros nomeados.
4. **FIX É CIRÚRGICO:** o prompt FIX corrige apenas as seções/violações reportadas pelo GATE. Nunca regenera o artefato inteiro.
5. **ANCORAGEM DOCUMENTAL:** nenhuma ferramenta, versão ou configuração pode existir sem origem nos docs-base (500/520/490/480/550/044/043/041). Ferramenta não prevista nos docs-base exige aprovação humana + atualização dos docs-base ANTES da instalação.
6. **SÓ AVANÇA COM COMPLIANCE:** o roadmap só avança para o próximo artefato quando o atual estiver `[STATUS: COMPLIANCE]` e o humano confirmar.
7. **HITL OBRIGATÓRIO:** aplicar mudança em ambiente compartilhado (QA/HMG/PROD) exige aprovação humana explícita antes de executar — o agente propõe o plano, nunca aplica sozinho.

## VOCABULÁRIO WATERFALL (tabela de tradução — aplicar SOMENTE em contexto WATERFALL; em contexto ágil, usar a terminologia do próprio projeto)

| Termo ágil (PROIBIDO) | Equivalente WATERFALL (usar) |
|---|---|
| Epic | Pacote de trabalho da EAP (060-EAP-WBS) |
| Feature | Funcionalidade `FEAT-NN` (010-FRD) |
| User Story | Caso de Uso `UC-NN` (010-FRD) |
| Definition of Ready (DoR) | GATE de COMPLIANCE do documento de origem |
| Sprint | Ciclo de entrega `FILA-NN` (definido pelo 092) |
| Product Backlog | 088-PRODUCT-BACKLOG-LIST (operado pelo 092 na FASE 5) |
| Deploy contínuo | Automação de deploy alinhada ao 041/087 (GMUD no 090) |

---

## VARIÁVEIS DE ENTRADA E BOOTSTRAP (FASE 0)

### Tabela de Inputs

| Variável | Obrig. | Descrição | Exemplo |
|---|---|---|---|
| `PROJECT_PATH` | ✅ | Caminho base onde os projetos de negócio residem | `/home/bolismar/work/workspace-fbso/business-inputs/business-projects` |
| `PROJECT_ID_NAME` | ✅ | Identificador completo do projeto | `PRJ-TEC-2026-0004-PROJETO-SHIELD` |
| `TECH_DEFS_DIR` | ✅ | Pasta das definições técnicas TECHLEAD (480/490/500/510/520/550) | `{PROJECT_PATH}/{PROJECT_ID_NAME}/technical-definitions/` |
| `WATERFALL_DOCS_DIR` | ✅ | Pasta dos documentos WATERFALL (041/043/044/087/096) | `{PROJECT_PATH}/{PROJECT_ID_NAME}/` |
| `TECHNICAL_SOLUTIONS` | ✅ | Lista de soluções técnicas do projeto (para 610) | `["ms-fbso-platform-admin", "web-app-fbso-platform-portal"]` |
| `TARGET_TOOLS` | ✅ | Lista de ferramentas a implementar (F3/F4) | `["keycloak", "rabbitmq", "kafka", "kestra"]` |
| `TARGET_ENVIRONMENTS` | ❌ | Ambientes alvo das implementações (default: DEV e QA — HMG/PROD com GMUD em contexto WATERFALL) | `["DEV", "QA"]` |
| `ARCHITECTURE_GLOBAL` | ✅ | Caminho da pasta de arquitetura global (ADRs, blueprints) | `/home/bolismar/work/workspace-fbso/architecture/` |
| `SECURITY_GLOBAL` | ✅ | Caminho do GLOBAL-SECURITY.md | `/home/bolismar/work/workspace-fbso/.specs/security/GLOBAL-SECURITY.md` |
| `PROJECT_PROMPT_INPUTS` | ❌ | **(Diretiva comportamental)** Checkpoint HITL: porta sempre aberta para input humano | `{checkpoint HITL}` |

### Variáveis Derivadas

```
PROJECT_COMPLETE_PATH_NAME = PROJECT_PATH + "/" + PROJECT_ID_NAME
TOOLING_OUTPUT_DIR         = PROJECT_COMPLETE_PATH_NAME + "/implementation-tooling/"
  ├── manifests/            ← saída da F1 (610)
  ├── observability/        ← saída da F2 (620)
  ├── tools/{FERRAMENTA}/   ← saída da F3 (630, uma pasta por ferramenta)
  └── security-tools/{FERRAMENTA}/ ← saída da F4 (640)
```

### Barreira de Entrada (obrigatória — antes de qualquer fase)

1. Confirmar que TODOS os docs-base exigidos existem e estão `[STATUS: COMPLIANCE]`:
   - TECHLEAD: `500-DEVOPS-SRE-DEFINITION`, `520-INFRA-CLOUD-DEFINITION`, `480-SECURITY-DEFINITION`, `490-DATA-ARCHITECTURE-DEFINITION`, `550-SOLUTIONS-STACK-MATRIX` (+ `510` para thresholds de observabilidade/SLO).
   - WATERFALL: `041-DEVOPS-SETUP`, `043-SEC-SETUP`, `044-INFRA-SETUP` (e `087-PLANO-CI-CD-AMBIENTES` para ambientes).
   - Doc-base ausente ou sem COMPLIANCE → **PARE** e informe ao humano qual gate destravar primeiro.
2. Cruzar `TARGET_TOOLS` com os docs-base: toda ferramenta da lista deve estar prevista em `550` (versões), `490` (dados/ETL/streaming), `480` (IAM/secrets) ou `520` (infra). Ferramenta órfã → aprovação humana + atualização dos docs-base (Regra de Ouro 5).
3. Auditar `TOOLING_OUTPUT_DIR`: artefatos já existentes entram em modo atualização (GENERATE com `MODE=update`), nunca recriação cega.

---

## Arquitetura de Fases

```
        ┌────────────────────────────────────────────────────────────┐
        │  F1 — 610 MANIFESTS-DEVOPS (Dockerfile, Helm, K8s, IaC)    │
        └────────────────────────────┬───────────────────────────────┘
                                     🚧 Barreira 1 (manifests COMPLIANCE)
        ┌────────────────────────────────────────────────────────────┐
        │  F2 — 620 OBSERVABILITY-SETUP (stack + dashboards + alertas)│
        └────────────────────────────┬───────────────────────────────┘
                                     🚧 Barreira 2 (observabilidade COMPLIANCE)
        ┌────────────────────────────────────────────────────────────┐
        │  F3 — 630 INSTALL-TOOL-{FERRAMENTA} (uma execução por tool) │
        └────────────────────────────┬───────────────────────────────┘
                                     🚧 Barreira 3 (ferramentas COMPLIANCE)
        ┌────────────────────────────────────────────────────────────┐
        │  F4 — 640 INSTALL-SECURITY-TOOL-{FERRAMENTA} (uma por tool) │
        └────────────────────────────────────────────────────────────┘
                                     🚩 CONCLUSÃO (relatório consolidado)
```

## Loop de Execução por Artefato (idêntico ao master)

```
STEP 1  COMPUTAR inputs do artefato (docs-base do trio, caminhos de saída)
STEP 2  GENERATE  → cria o artefato com [STATUS: Em análise]
STEP 3  GATE      → valida contra os docs-base (checklist) → [STATUS: Em revisão]
STEP 4a FIX       → corrige APENAS as seções violadas (cirúrgico)
STEP 4b Validação humana (P1-P4: ancoragem, completude, segurança, vocabulário)
STEP 5  [STATUS: COMPLIANCE] + confirmação humana → próximo artefato
```

---

## F1 — 610 MANIFESTS-DEVOPS (Manifestos DevOps)

| Item | Detalhe |
|---|---|
| **Necessidade** | Gerar os manifestos reais das soluções: Dockerfile multi-stage, Helm charts, K8s YAML, módulos Terraform (IaC) |
| **Docs-base** | `500-DEVOPS-SRE-DEFINITION` (containers/Helm/IaC), `520-INFRA-CLOUD-DEFINITION` (topologia/compute/rede), `550-SOLUTIONS-STACK-MATRIX` (versões), `044-INFRA-SETUP` (provisionamento), `041-DEVOPS-SETUP` (IaC no pipeline) |
| **Saídas** | `manifests/{SOLUCAO}/Dockerfile`, `helm/`, `k8s/`, `terraform/` + relatório de rastreabilidade manifest → doc-base |
| **Trio** | `610-MANIFESTS-DEVOPS` (GENERATE/GATE/FIX) |

## F2 — 620 OBSERVABILITY-SETUP (Observabilidade)

| Item | Detalhe |
|---|---|
| **Necessidade** | Implantar a stack de observabilidade e os painéis de SLO: logging estruturado, métricas, tracing, alerting de burn rate |
| **Docs-base** | `500-DEVOPS-SRE-DEFINITION` (observabilidade/SLOs/SLIs/alerting/runbooks), `041-DEVOPS-SETUP` (observabilidade), `510-TEST-STRATEGY-DEFINITION` (thresholds) |
| **Saídas** | `observability/` — manifestos da stack (Prometheus/Grafana/Loki/Jaeger/OpenTelemetry), dashboards por SLO, regras de alerta, runbook de incidente S1-S4 |
| **Trio** | `620-OBSERVABILITY-SETUP` (GENERATE/GATE/FIX) |

## F3 — 630 INSTALL-TOOL-{FERRAMENTA} (Middleware/ETL/Orquestração)

Template parametrizado — **uma execução por ferramenta** de `TARGET_TOOLS`. Ferramentas previstas pelo catálogo:

| Ferramenta | Configurações típicas ancoradas nos docs-base |
|---|---|
| Keycloak | Realms, clients OIDC, flows, identity providers (`480` IAM) |
| RabbitMQ | Vhosts, queues, exchanges, políticas, TLS (`490` mensageria, `520` rede) |
| Apache Kafka | Cluster, topics, partitions, ACLs, schema registry (`490` streaming) |
| Kestra / Camunda | Instalação do orquestrador, workers, flows/processos iniciais (`490` orquestradores) |
| Airbyte / NiFi / dbt | Conexões ETL/ELT, pipelines iniciais, schedule (`490` ETL/ELT) |

| Item | Detalhe |
|---|---|
| **Docs-base** | `490-DATA-ARCHITECTURE-DEFINITION` (ETL/streaming/orquestradores), `480-SECURITY-DEFINITION` (IAM/secrets), `520-INFRA-CLOUD-DEFINITION` (compute/rede), `550-SOLUTIONS-STACK-MATRIX` (versões), `044-INFRA-SETUP` (topologia), `043-SEC-SETUP` (hardening) |
| **Saídas** | `tools/{FERRAMENTA}/` — manifestos de instalação (Helm/compose), configuração, hardening, smoke test e relatório de instalação com rastreabilidade |
| **Trio** | `630-INSTALL-TOOL` (GENERATE/GATE/FIX, parâmetro `{FERRAMENTA}`) |

## F4 — 640 INSTALL-SECURITY-TOOL-{FERRAMENTA} (Ferramentas de Segurança)

| Ferramenta | Configurações típicas ancoradas nos docs-base |
|---|---|
| Wazuh | Agentes, rulesets, alertas, integração SIEM (`043` controles, `520` rede) |
| HashiCorp Vault | Engines, policies, autenticação, integração com aplicações (`480` secrets) |
| WAF (Cloudflare/AWS WAF) | Regras, rate limiting, logging (`043`/`480`/`520` segurança de infra) |

| Item | Detalhe |
|---|---|
| **Docs-base** | `043-SEC-SETUP` (controles por camada/DevSecOps), `480-SECURITY-DEFINITION` (threat model/secrets/IAM), `520-INFRA-CLOUD-DEFINITION` (WAF/segurança de infra) |
| **Saídas** | `security-tools/{FERRAMENTA}/` — instalação, políticas, integração, teste de eficácia e relatório |
| **Trio** | `640-INSTALL-SECURITY-TOOL` (GENERATE/GATE/FIX, parâmetro `{FERRAMENTA}`) |

---

## Integração com os Demais Roadmaps

| Roadmap | Ponto de integração |
|---|---|
| `PROJECT-TECHNICAL-DEFINITIONS` (TECHLEAD, Bloco F) | Tarefas de infra/ferramentas das janelas DEV/QA invocam este roadmap; Bloco E referencia `610` (manifests) via `PROMPT-EXECUTE-CI-CD-PIPELINE` |
| `WATERFALL-EXECUTION` (FASE 5) | Setup de ambiente por ciclo de entrega (`FILA-NN`): tooling é pré-requisito de ambiente antes da execução do código |
| `PROJECT-DOCUMENTS-WATERFALL` (master) | Companion opcional da FASE 5 — consume `041/043/044` (F3) e alimenta `095-RELATORIO-QUALIDADE`/`100-MANUAIS-OPERACIONAIS` com evidências e runbooks |
| `sprint-tecnhnical-implementation/` | `PROMPT-EXECUTE-CI-CD-PIPELINE` e `PROMPT-EXECUTE-CVE-SCA-SCAN` consomem manifestos/ferramentas produzidos aqui |

## Git Workflow

1. Branch `feature/tooling-<slug>` a partir da principal (nunca main/master direto).
2. Commits convencionais: `feat(TOOLING-NN): <descrição>` por artefato aprovado.
3. PR obrigatório com revisão humana; merge somente após `[STATUS: COMPLIANCE]` + aprovação.
4. Arquivos de manifesto vivem no repositório de infra da organização (ou na pasta `devops/` do monorepo), conforme `500`/`520` — nunca em diretórios pessoais.

## Localização dos Prompts e Contagem

| Local | Arquivos |
|---|---|
| `.specs/prompts/` (raiz) | `PROMPT-ROADMAP-GENERATE-IMPLEMENTATION-TOOLING.md` |
| `.specs/prompts/implementation-tooling/` | 12 prompts (4 trios × GENERATE/GATE/FIX) + 1 FLOWCHART |
| **Total** | **13 arquivos (12 prompts + FLOWCHART)** |

| Trio | Arquivos |
|---|---|
| 610-MANIFESTS-DEVOPS | `PROMPT-GENERATE-610-MANIFESTS-DEVOPS.md`, `PROMPT-GATE-610-MANIFESTS-DEVOPS.md`, `PROMPT-FIX-610-MANIFESTS-DEVOPS.md` |
| 620-OBSERVABILITY-SETUP | `PROMPT-GENERATE-620-OBSERVABILITY-SETUP.md`, `PROMPT-GATE-620-OBSERVABILITY-SETUP.md`, `PROMPT-FIX-620-OBSERVABILITY-SETUP.md` |
| 630-INSTALL-TOOL | `PROMPT-GENERATE-630-INSTALL-TOOL.md`, `PROMPT-GATE-630-INSTALL-TOOL.md`, `PROMPT-FIX-630-INSTALL-TOOL.md` |
| 640-INSTALL-SECURITY-TOOL | `PROMPT-GENERATE-640-INSTALL-SECURITY-TOOL.md`, `PROMPT-GATE-640-INSTALL-SECURITY-TOOL.md`, `PROMPT-FIX-640-INSTALL-SECURITY-TOOL.md` |
| FLOWCHART | `FLOWCHART-PROMPT-ROADMAP-GENERATE-IMPLEMENTATION-TOOLING.md` |
