# PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION

## Contexto

Este prompt implementa o **Gate de Validação da Definição de Infraestrutura Cloud** para o artefato {UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION.md`. Verifica se a infraestrutura do projeto está completa, cobre topologia, compute, networking, storage, DR, dimensionamento e custos, e está alinhada com a arquitetura, segurança e DevOps.

**Princípio fundamental:** Toda solução deve ter sua infraestrutura de suporte definida. Nenhum recurso de infra pode ser provisionado sem estar documentado neste artefato, com dimensionamento justificado e custo estimado.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` — artefato auditado (F7)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Definição de Dados (F4)
5. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — Definição DevOps/SRE (F5)
6. `{ARCHITECTURE_GLOBAL}/` — ADRs, blueprints globais

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

**Arquivos gerados pelo GENERATE:** `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION.md`

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}` — base dos projetos de negócio
- `PROJECT_ID_NAME={PROJECT_ID_NAME}` — identificador do projeto
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
- `PROJECT-TEAM-SKILLS-MAP`, `PROJECT-TEAM-CAPACITY`, `PROJECT-STACK` — se fornecidos
Validar que o artefato auditado `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` existe.

### Passo 1 — Carregar Documentos Base
Confirmar leitura dos seguintes artefatos:
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` — artefato auditado (F7)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Definição de Dados (F4)
5. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — Definição DevOps/SRE (F5)
6. `{ARCHITECTURE_GLOBAL}/` — ADRs, blueprints globais

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de Infraestrutura
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Topologia definida | Diagrama de topologia para todos os ambientes (Dev/Staging/Prod) |
| 1.2 | Compute dimensionado | VMs, K8s, Serverless com sizing e justificativa |
| 1.3 | Networking completo | VPC, subnets, DNS, CDN, load balancers, NAT/VPN |
| 1.4 | Storage hierarquizado | Block, Object, File por performance/custo |

#### Dimensão 2: Resiliência e DR
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Disaster Recovery | RPO e RTO definidos por solução |
| 2.2 | Backup | Schedule, retenção, procedimento de restore |
| 2.3 | Multi-region/Failover | Estratégia ativo-ativo ou ativo-passivo documentada |

#### Dimensão 3: Dimensionamento e Custos
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Sizing inicial | Memória, CPU, storage por recurso |
| 3.2 | Auto-scaling | Políticas definidas (horizontal/vertical) |
| 3.3 | Estimativa de custos | Breakdown mensal por ambiente e serviço |

#### Dimensão 4: Consistência
| # | Verificação | Critério |
|---|---|---|
| 4.1 | Alinhamento com ARCHITECTURE | Containers e redes refletem a topologia da arquitetura |
| 4.2 | Alinhamento com SECURITY | WAF, security groups, NACLs, IAM de infra, encryption |
| 4.3 | Alinhamento com DEVOPS-SRE | Infra suporta pipeline CI/CD e observabilidade |

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

#### 📊 RELATÓRIO DE AUDITORIA DE INFRAESTRUTURA CLOUD: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-INFRA-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que a Architecture/Security/DevOps determinava:** [Descrever a referência]
  - **Impacto:** [O risco de infra sem DR, custo não estimado, segurança de rede inadequada]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir a definição de infraestrutura cloud, por favor, responda:
1. Quanto ao **[ID-CONFLITO-INFRA-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION.md)*

---

### ✅ CENÁRIO B: SE A INFRAESTRUTURA ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE INFRAESTRUTURA CLOUD: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION.md` gerado conforme Architecture Definition, Security Definition, Data Architecture e DevOps SRE.
- **AUDITORIA DA IA:** Topologia definida para todos os ambientes. Compute dimensionado e justificado. Networking completo. DR com RPO/RTO. Sizing com auto-scaling. Segurança de infra alinhada com SECURITY-DEFINITION. Estimativa de custos presente.
- **DIRETRIZ:** Peço que leia a definição de infraestrutura cloud para verificar se a topologia, dimensionamento e custos refletem as necessidades do projeto.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. A definição de infraestrutura cloud está em compliance e reflete corretamente a infraestrutura necessária para o projeto?
2. Deseja enviar mais documentos/arquivos para enriquecer a definição de infraestrutura?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Este documento não possui uma próxima fase de technical-definition (Bloco B completo). O fluxo segue para o Bloco C (F13-F17 — Solutions Catalog, Matrix, Stack Matrix). Se novos inputs → retrocede ao PROMPT-GENERATE).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `cloud-architect` | Validar arquitetura cloud | Cloud |
| 2 | `aws-solution-architect` | Validar topologia AWS | AWS |
| 3 | `kubernetes-architect` | Validar infra K8s | K8s |
| 4 | `network-engineer` | Validar redes e conectividade | Rede |
| 5 | `disaster-recovery` | Validar estratégia de DR | DR |
| 6 | `cost-optimization` | Validar estimativa de custos | Custos |
| 7 | `gap-analysis` | Identificar gaps de infra | Análise |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: gate de validação da definição de infraestrutura cloud | Time de Arquitetura |

---

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` | Artefato auditado (F7) |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` | Definição de Arquitetura (F2) |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` | Definição de Segurança (F3) |
| 4 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` | Definição de Dados (F4) |
| 5 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` | Definição DevOps/SRE (F5) |
| 6 | `{ARCHITECTURE_GLOBAL}/` | ADRs, blueprints globais |
| 7 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 8 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |
| 9 | `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (obter e validar com usuário) |
| 10 | `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time (obter e validar com usuário) |
| 11 | `{PROJECT-STACK}` | Stack tecnológica. Baseline: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
