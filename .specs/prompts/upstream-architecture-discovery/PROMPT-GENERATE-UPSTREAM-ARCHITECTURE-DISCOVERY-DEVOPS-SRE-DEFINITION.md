# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION

## Contexto

Este prompt gera o artefato `UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION.md` — a **definição de DevOps e SRE do projeto** que especifica CI/CD, IaC, observabilidade, SLOs, runbooks e gestão de ambientes. Este documento é a referência operacional para toda a esteira de deploy e operação do projeto.

**Relação com ARCHITECTURE-DEFINITION:** Enquanto o ARCHITECTURE-DEFINITION (F7) define a topologia de containers e comunicação entre soluções, este documento detalha **como essas soluções são construídas, testadas, implantadas e operadas** em produção.

**Inputs upstream:** `UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md` (Fase 7) + `UPSTREAM-ARCHITECTURE-DISCOVERY-SECURITY-DEFINITION.md` (Fase 8) + `UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION.md` (Fase 9) + `{ARCHITECTURE_GLOBAL}` (blueprints, ADRs).

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

### Passo 1 — Carregar Documentos Base
Ler Architecture Definition (F7 — topologia de deploy, containers, redes), Security Definition (F8 — DevSecOps, SAST, SCA, secrets), Data Architecture (F9 — pipelines que precisam ser operados), ADRs globais, blueprints (Docker, K8s, CI/CD).

### Passo 2 — Invocar Skills Especializadas
Invocar skills de DevOps, SRE, observabilidade, containers, IaC e segurança operacional para projetar a estratégia de operações do projeto.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION.md` com:

1. **Pipeline CI/CD** — build, teste, deploy, rollback automatizado, promoção entre ambientes (Dev/Staging/Prod), approval gates
2. **Infrastructure as Code** — Terraform/CloudFormation/Pulumi, estrutura de repositórios, módulos compartilhados, state management, drift detection
3. **Observabilidade** — logging (formato estruturado, níveis, retenção), métricas (Micrometer/Prometheus), tracing distribuído (OpenTelemetry), alerting (PagerDuty/OpsGenie/Slack), dashboards (Grafana)
4. **SLOs/SLIs** — error budgets, latency targets (p50/p95/p99), availability targets (9s), burn rate alerts
5. **Containers e Orquestração** — Docker (Dockerfile multi-stage, distroless), Kubernetes (EKS/AKS/GKE), Helm charts, service mesh (Istio/Linkerd)
6. **Gestão de Ambientes** — Dev, Staging, Prod, feature branches (preview environments), ephemeral environments
7. **Runbooks** — procedimentos de incidente, escalação (severidade S1-S4), postmortems, disaster recovery runbooks
8. **Ferramentas** — CI/CD (GitHub Actions/GitLab CI/Jenkins), monitoring (Grafana/Datadog/New Relic), logging (ELK/Loki/Datadog Logs), containers (Docker/Helm)

### Passo 4 — Validação Pós-Geração
Verificar: pipeline CI/CD completo (build→deploy→rollback), IaC documentado, observabilidade cobre logs+métricas+tracing, SLOs definidos com burn rates, runbooks criados para S1-S2, consistência com ARCHITECTURE (topologia) e SECURITY (DevSecOps, SAST no pipeline).

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

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
