# PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION

## Contexto

Este prompt implementa o **Gate de Validação da Definição de DevOps e SRE** para o artefato `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION.md`. Verifica se a estratégia de operações do projeto está completa, cobre CI/CD, IaC, observabilidade, SLOs e runbooks, e está alinhada com a arquitetura e segurança.

**Princípio fundamental:** Toda solução deve ter seu pipeline CI/CD, observabilidade e SLOs definidos. Nenhum deploy pode ser manual sem justificativa. Nenhuma solução pode ficar sem monitoramento.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — artefato auditado (F5)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Definição de Dados (F4)

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
| `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (obter e validar com usuário) |
| `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time (obter e validar com usuário) |
| `{PROJECT-STACK}` | Stack tecnológica. Baseline: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
UPSTREAM_DISCOVERY_PATH       = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
```

---

**Arquivos gerados pelo GENERATE:** `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION.md`

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}`, `PROJECT_ID_NAME={PROJECT_ID_NAME}`
- `PROJECT_DOCUMENTS_INPUTS`, `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
- `PROJECT-TEAM-SKILLS-MAP`, `PROJECT-TEAM-CAPACITY`, `PROJECT-STACK` — se fornecidos
Validar que o artefato auditado `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` existe.

### Passo 1 — Carregar Documentos Base
Confirmar leitura dos artefatos:
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — artefato auditado (F5)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Dados (F4)

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Pipeline CI/CD
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Pipeline documentado | Build, teste, deploy, rollback definidos |
| 1.2 | Ambientes mapeados | Dev, Staging, Prod com promoção entre eles |
| 1.3 | Approval gates | Deploy em Prod requer aprovação documentada |

#### Dimensão 2: Infraestrutura e Observabilidade
| # | Verificação | Critério |
|---|---|---|
| 2.1 | IaC definido | Terraform/CloudFormation/Pulumi com estrutura de repositórios |
| 2.2 | Logging | Formato estruturado, níveis, retenção |
| 2.3 | Métricas | Micrometer/Prometheus, dashboards (Grafana) |
| 2.4 | Tracing | OpenTelemetry implementado |
| 2.5 | Alerting | Canais (PagerDuty/Slack) e thresholds definidos |

#### Dimensão 3: SLOs e Runbooks
| # | Verificação | Critério |
|---|---|---|
| 3.1 | SLOs/SLIs definidos | Error budgets, latency targets, availability targets |
| 3.2 | Runbooks de incidente | Procedimentos S1-S4 com escalação |
| 3.3 | Postmortem | Processo de blameless postmortem definido |

#### Dimensão 4: Consistência
| # | Verificação | Critério |
|---|---|---|
| 4.1 | Alinhamento com ARCHITECTURE | Topologia de deploy reflete containers da arquitetura |
| 4.2 | Alinhamento com SECURITY | DevSecOps integrado (SAST, SCA, secrets no pipeline) |

#### Dimensão 5: Alinhamento com Stack Corporativa e Time

| # | Verificação | Critério |
|---|---|---|
| 5.1 | Skills mapeados | Skills necessários para esta disciplina estão documentados no artefato |
| 5.2 | Capacidade estimada | Capacidade do time está dimensionada proporcionalmente à complexidade |
| 5.3 | Stack corporativa | Tecnologias propostas constam no `STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |
| 5.4 | Tecnologias adicionais | Tecnologias fora do padrão corporativo têm justificativa técnica documentada e aprovada |

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE DEVOPS E SRE: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-DEVOPS-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que a Architecture/Security determinava:** [Descrever a referência]
  - **Impacto:** [O risco operacional — deploy inseguro, sem observabilidade, sem SLO]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir a definição de DevOps e SRE, por favor, responda:
1. Quanto ao **[ID-CONFLITO-DEVOPS-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION.md)*

---

### ✅ CENÁRIO B: SE A DEFINIÇÃO DE DEVOPS/SRE ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE DEVOPS E SRE: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION.md` gerado conforme Architecture Definition, Security Definition e Data Architecture.
- **AUDITORIA DA IA:** Pipeline CI/CD completo. IaC documentado. Observabilidade cobre logs+métricas+tracing. SLOs definidos com burn rates. Runbooks criados. Consistente com Architecture (topologia) e Security (DevSecOps).
- **DIRETRIZ:** Peço que leia a definição de DevOps e SRE para verificar se o pipeline, observabilidade e SLOs atendem às expectativas operacionais.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. A definição de DevOps e SRE está em compliance e reflete corretamente a estratégia de operações do projeto?
2. Deseja enviar mais documentos/arquivos para enriquecer a definição de DevOps e SRE?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se "Sim, Não, Não" → [STATUS: COMPLIANCE] e Fase 11 (TEST-STRATEGY-DEFINITION). Se novos inputs → retrocede ao PROMPT-GENERATE).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `senior-devops` | Validar estratégia DevOps | DevOps |
| 2 | `sre-engineer` | Validar SLOs e runbooks | SRE |
| 3 | `observability-engineer` | Validar observabilidade | Observabilidade |
| 4 | `kubernetes-specialist` | Validar topologia K8s | K8s |
| 5 | `terraform-specialist` | Validar IaC | IaC |
| 6 | `gap-analysis` | Identificar gaps operacionais | Análise |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: gate de validação da definição de DevOps e SRE | Time de Arquitetura |

---

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` | Artefato auditado (F5) |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` | Definição de Arquitetura (F2) |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` | Definição de Segurança (F3) |
| 4 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` | Definição de Dados (F4) |
| 5 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 6 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |
| 7 | `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (obter e validar com usuário) |
| 8 | `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time (obter e validar com usuário) |
| 9 | `{PROJECT-STACK}` | Stack tecnológica. Baseline: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
