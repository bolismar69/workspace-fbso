# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION

## Contexto

Este prompt gera o artefato `UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION.md` — a **definição de DevOps e SRE do projeto** que especifica CI/CD, IaC, observabilidade, SLOs, runbooks e gestão de ambientes. Este documento é a referência operacional para toda a esteira de deploy e operação do projeto.

**Relação com ARCHITECTURE-DEFINITION:** Enquanto o ARCHITECTURE-DEFINITION (F7) define a topologia de containers e comunicação entre soluções, este documento detalha **como essas soluções são construídas, testadas, implantadas e operadas** em produção.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura de Dados (F4)
4. `{ARCHITECTURE_GLOBAL}/` — blueprints, ADRs

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Diretiva) Checkpoint HITL: sempre solicitar ao usuário se deseja fornecer informações adicionais ou novos direcionamentos via prompt |
| `{PROJECT-TEAM-SKILLS-MAP}` | Skills necessários para o time de implementação (obter e validar com usuário) |
| `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time — seniores, plenos, juniores, duração (obter e validar com usuário) |
| `{PROJECT-STACK}` | Stack tecnológica da solução. Baseline corporativa: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`. Tecnologias fora do padrão exigem justificativa |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
UPSTREAM_DISCOVERY_PATH       = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
```

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}` — base dos projetos de negócio
- `PROJECT_ID_NAME={PROJECT_ID_NAME}` — identificador do projeto
- `TECHNICAL_SOLUTION_PATH={TECHNICAL_SOLUTION_PATH}` — base das soluções técnicas
- `TECHNICAL_SOLUTION_NAMES={TECHNICAL_SOLUTION_NAMES}` — soluções do projeto
- `ARCHITECTURE_GLOBAL={ARCHITECTURE_GLOBAL}` — ADRs e blueprints globais
- `SECURITY_GLOBAL={SECURITY_GLOBAL}` — documento de segurança global
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
- `PROJECT-TEAM-SKILLS-MAP` — skills do time (se fornecidos)
- `PROJECT-TEAM-CAPACITY` — capacidade do time (se fornecida)
- `PROJECT-STACK` — stack tecnológica; validar contra STACK-PADROES-CORPORATIVOS-FBSO-ORG.md
Validar que `{UPSTREAM_DISCOVERY_PATH}` existe e contém os artefatos upstream.

### Passo 1 — Carregar Documentos Base
Confirmar leitura dos seguintes artefatos upstream:
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2) — topologia de deploy
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3) — DevSecOps
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Definição de Dados (F4) — pipelines
4. `{ARCHITECTURE_GLOBAL}/` — ADRs globais, blueprints (Docker, K8s, CI/CD)

### Passo 2 — Invocar Skills Especializadas
Invocar skills de DevOps, SRE, observabilidade, containers, IaC e segurança operacional para projetar a estratégia de operações do projeto.

### Passo 2.5 — Apresentar Skills, Capacidade e Stack para Validação Humana

Avaliar e apresentar ao usuário para validação:

1. **PROJECT-TEAM-SKILLS-MAP:** Skills identificados como necessários para implementar a solução nesta disciplina.
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

2. **PROJECT-TEAM-CAPACITY:** Capacidade estimada do time nesta disciplina (ex: 2 seniores, 3 plenos).
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

3. **PROJECT-STACK:** Tecnologias identificadas para esta disciplina. Verificar conformidade com `STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`. Tecnologias fora do padrão corporativo DEVEM ser listadas com justificativa técnica e requerem aprovação explícita do usuário.
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

### Passo 3 — Gerar o Artefato
Gerar `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION.md` com:

1. **Pipeline CI/CD** — build, teste, deploy, rollback automatizado, promoção entre ambientes (Dev/Staging/Prod), approval gates
2. **Infrastructure as Code** — Terraform/CloudFormation/Pulumi, estrutura de repositórios, módulos compartilhados, state management, drift detection
3. **Observabilidade** — logging (formato estruturado, níveis, retenção), métricas (Micrometer/Prometheus), tracing distribuído (OpenTelemetry), alerting (PagerDuty/OpsGenie/Slack), dashboards (Grafana)
4. **SLOs/SLIs** — error budgets, latency targets (p50/p95/p99), availability targets (9s), burn rate alerts
5. **Containers e Orquestração** — Docker (Dockerfile multi-stage, distroless), Kubernetes (EKS/AKS/GKE), Helm charts, service mesh (Istio/Linkerd)
6. **Gestão de Ambientes** — Dev, Staging, Prod, feature branches (preview environments), ephemeral environments
7. **Runbooks** — procedimentos de incidente, escalação (severidade S1-S4), postmortems, disaster recovery runbooks
8. **Ferramentas** — CI/CD (GitHub Actions/GitLab CI/Jenkins), monitoring (Grafana/Datadog/New Relic), logging (ELK/Loki/Datadog Logs), containers (Docker/Helm)

---
## Layout do Documento (Modelo Estrutural)

> 📐 **Modelo de referência:** O documento `business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/upstream-architecture-discovery/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` ilustra a estrutura esperada. Use-o como referência de **formato e organização**, NÃO como fonte de dados — todo conteúdo deve ser gerado a partir dos artefatos do projeto corrente (`{PROJECT_ID_NAME}`).

### Estrutura Esperada do `DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md`

```markdown
# DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md
## Fase 5 — Bloco B: Architecture & Security & Specialists (Discovery-Level)

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Análise de Viabilidade) |
| **Data** | {DATA_ATUAL} |
| **Autor** | DevOps/SRE Architect |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em {DATA} |

**Documentos Vinculados:**
- [`DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md`](DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md) — Definição de Arquitetura (F2)
- [`DISCOVERY-LEVEL-SECURITY-DEFINITION.md`](DISCOVERY-LEVEL-SECURITY-DEFINITION.md) — Definição de Segurança (F3)
- [`STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`](../../../.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md)

---

## 1. Estratégia DevOps/SRE — Visão Macro
- **1.1 Pipeline CI/CD:** Diagrama ASCII do fluxo (Commit → Build & Test → Docker Build & Push → Deploy to K8s → Verify Health)
- **1.2 Workflow (GitHub Actions ou equivalente):** Tabela: Stage | Ação | Ferramenta | Gate
- **1.3 Ambientes:** Tabela: Ambiente | Propósito | Deploy | Recursos

## 2. Infraestrutura como Código (IaC)
- **2.1 Terraform (ou ferramenta equivalente) — Provisioning:** Tabela: Recurso | Provider | Configuração
- **2.2 Ansible (ou equivalente) — Configuração:** Tabela: Playbook | Alvo | Configura

## 3. Kubernetes — Orquestração e Service Mesh
- **3.1 Topologia de Deploy K8s:** Estrutura YAML de namespaces
- **3.2 Service Mesh (Istio ou equivalente):** Tabela: Feature | Configuração
- **3.3 Autoscaling:** Tabela: Componente | Ferramenta | Trigger

## 4. Observabilidade — Stack Completa
- **4.1 Stack Overview:** Diagrama ASCII do fluxo de telemetria (Aplicação → Coleta → Dashboards)
- **4.2 Métricas:** Tabela: Métrica | Tipo | Descrição
- **4.3 Alertas Críticos:** Tabela: Alerta | Condição | Severidade | Canal
- **4.4 Dashboards:** Tabela: Dashboard | Público | Métricas

## 5. SLOs — Service Level Objectives
- Tabela: SLO | Target | Measurement Window

## 6. Riscos e Estimativa de Esforço
- **6.1 Riscos DevOps/SRE:** Tabela: ID | Risco | Prob. | Impacto | Mitigação
- **6.2 Estimativa de Esforço:** Tabela: Atividade | Complexidade | Esforço (dias) | Responsável

---

## Registro de Alterações
| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | {DATA} | Criação inicial: DevOps/SRE Definition Discovery-Level | DevOps/SRE Architect |
```

### Passo 4 — Validação Pós-Geração

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `senior-devops` | Supervisão sênior de DevOps | DevOps |
| 2 | `cloud-devops` | Estratégia DevOps multi-cloud | Cloud |
| 3 | `sre-engineer` | SLOs, SLIs, error budgets | SRE |
| 4 | `kubernetes-specialist` | Configuração e operação K8s | K8s |
| 5 | `docker-expert` | Containerização e Dockerfiles | Docker |
| 6 | `terraform-specialist` | IaC com Terraform/OpenTofu | IaC |
| 7 | `observability-engineer` | Observabilidade (logs+métricas+tracing) | Observabilidade |
| 8 | `slo-implementation` | Implementação de SLOs e burn rates | SRE |
| 9 | `monitoring-expert` | Monitoramento e alerting | Monitoramento |
| 10 | `cicd-automation-workflow-automate` | Pipeline CI/CD automatizado | CI/CD |
| 11 | `deployment-pipeline-design` | Design de pipeline de deploy | Deploy |
| 12 | `incident-response-incident-response` | Runbooks e resposta a incidentes | Incidente |
| 13 | `helm-chart-scaffolding` | Helm charts para deploy K8s | K8s |
| 14 | `mermaid-expert` | Diagramas de pipeline e fluxo | Diagramas |
| 15 | `documentation-writer` | Redigir o DevOps SRE Definition | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt gerador da definição de DevOps e SRE | Time de Arquitetura |

---

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` | Definição de Arquitetura (F2) — topologia de deploy |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` | Definição de Segurança (F3) — DevSecOps |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` | Definição de Dados (F4) — pipelines |
| 4 | `{ARCHITECTURE_GLOBAL}/` | ADRs globais, blueprints |
| 5 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 6 | `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (se fornecidos) |
| 7 | `{PROJECT-TEAM-CAPACITY}` | Capacidade do time (se fornecida) |
| 8 | `{PROJECT-STACK}` | Stack tecnológica (validar contra padrões corporativos) |
| 9 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
