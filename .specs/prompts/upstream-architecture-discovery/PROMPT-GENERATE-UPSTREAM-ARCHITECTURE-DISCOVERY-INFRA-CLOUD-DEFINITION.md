# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION

## Contexto

Este prompt gera o artefato `UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION.md` — a **definição de infraestrutura cloud/on-premise do projeto** que especifica topologia, compute, networking, storage, disaster recovery, dimensionamento e custos. Este documento é a referência técnica para toda a infraestrutura que suporta as soluções do projeto.

**Relação com DEVOPS-SRE-DEFINITION:** Enquanto o DEVOPS-SRE-DEFINITION (F10) define o pipeline CI/CD e a operação das soluções, este documento detalha **o substrato de infraestrutura sobre o qual essas soluções e pipelines são executados** — VMs, clusters, redes, storage, DR.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Definição de Dados (F4)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — Definição DevOps/SRE (F5)

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
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2) — topologia e redes
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3) — firewall, WAF
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Definição de Dados (F4) — storage
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — Definição DevOps/SRE (F5) — clusters K8s

### Passo 2 — Invocar Skills Especializadas
Invocar skills de arquitetura cloud, infraestrutura, redes, storage, DR e custos para projetar a infraestrutura completa.

### Passo 2.5 — Apresentar Skills, Capacidade e Stack para Validação Humana

Avaliar e apresentar ao usuário para validação:

1. **PROJECT-TEAM-SKILLS-MAP:** Skills identificados como necessários para implementar a solução nesta disciplina.
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

2. **PROJECT-TEAM-CAPACITY:** Capacidade estimada do time nesta disciplina (ex: 2 seniores, 3 plenos).
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

3. **PROJECT-STACK:** Tecnologias identificadas para esta disciplina. Verificar conformidade com `STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`. Tecnologias fora do padrão corporativo DEVEM ser listadas com justificativa técnica e requerem aprovação explícita do usuário.
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

### Passo 3 — Gerar o Artefato
Gerar `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION.md` com:

1. **Topologia de Infraestrutura** — On-Premise e/ou Cloud (AWS/Azure/GCP), regiões, zonas de disponibilidade, ambientes (Dev/Staging/Prod), diagrama de topologia
2. **Compute** — VMs (EC2/Azure VM/GCE — famílias, sizing), Kubernetes (EKS/AKS/GKE — node pools, auto-scaling), Serverless (Lambda/Functions/Cloud Run), justificativa de escolha
3. **Networking** — VPC/VNet, subnets (públicas/privadas), DNS (Route53/Azure DNS), CDN (CloudFront/Cloudflare), API Gateway, load balancers (ALB/NLB/Application Gateway), NAT, VPN/Direct Connect
4. **Storage** — Block (EBS/Managed Disks), Object (S3/Blob Storage), File (EFS/Azure Files/NFS), hierarquia de storage por performance/custo
5. **Disaster Recovery** — RPO, RTO por solução, estratégia de backup (schedule, retenção), multi-region ativo-ativo ou ativo-passivo, failover automático
6. **Dimensionamento** — sizing inicial (compute, memória, storage), auto-scaling policies (horizontal/vertical), limites e quotas, plano de crescimento
7. **Segurança de Infra** — WAF, security groups, NACLs, IAM de infraestrutura (roles, policies), encryption at rest (KMS) e in transit (TLS), network policies
8. **Estimativa de Custos** — calculadora por provedor (AWS Pricing Calculator/Azure Pricing), custo mensal estimado por ambiente, breakdown por serviço, opções de reserva (RIs/Savings Plans)

---
## Layout do Documento (Modelo Estrutural)

> 📐 **Modelo de referência:** O documento `business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/upstream-architecture-discovery/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` ilustra a estrutura esperada. Use-o como referência de **formato e organização**, NÃO como fonte de dados — todo conteúdo deve ser gerado a partir dos artefatos do projeto corrente (`{PROJECT_ID_NAME}`).

### Estrutura Esperada do `DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md`

```markdown
# DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md
## Fase 7 — Bloco B: Architecture & Security & Specialists (Discovery-Level)

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Análise de Viabilidade) |
| **Data** | {DATA_ATUAL} |
| **Autor** | Infra/Cloud Specialist |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em {DATA} |

**Documentos Vinculados:**
- [`DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md`](DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md) — Definição de Arquitetura (F2)
- [`DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md`](DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md) — DevOps/SRE (F5)
- [`STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`](../../../.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md)

---

## 1. Provedor Cloud e Topologia
- **1.1 Provedor:** Tabela: Critério | Decisão (Provedor Primário, Região Primária, Região de DR, Edge/CDN/WAF)
- **1.2 Topologia de Rede:** Diagrama ASCII da topologia (Internet → Edge → Load Balancer → Cluster → DBs)
- **1.3 Fluxo de Acesso Completo:** Passos numerados (1 a N) com o caminho completo do request

## 2. Recursos Cloud — Estimativa de Custos
- **2.1 Custos Mensais Estimados (por provedor):** Tabela: Recurso | Especificação | Dev | Staging | Prod | Custo Mensal (Prod)
- **2.2 Edge/CDN/WAF — Custos:** Tabela: Plano | Features | Custo Mensal
- **2.3 Custo Total Estimado:** Tabela: Ambiente | Provedor Primário | Edge/WAF | Total

## 3. Disaster Recovery e Continuidade
- **3.1 Estratégia de DR:** Tabela: Cenário | RPO | RTO | Estratégia
- **3.2 Backup Strategy:** Tabela: Recurso | Frequência | Retenção | Local

## 4. Segurança de Infraestrutura
- Tabela: Controle | Implementação (Network Isolation, Firewall, TLS, Encryption at Rest, Secrets, SSH, Audit)

## 5. Riscos e Estimativa de Esforço
- **5.1 Riscos de Infra/Cloud:** Tabela: ID | Risco | Prob. | Impacto | Mitigação
- **5.2 Estimativa de Esforço:** Tabela: Atividade | Complexidade | Esforço (dias) | Responsável

---

## Registro de Alterações
| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | {DATA} | Criação inicial: Infra/Cloud Definition Discovery-Level | Infra/Cloud Specialist |
```

### Passo 4 — Validação Pós-Geração

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `cloud-architect` | Arquitetura cloud geral | Cloud |
| 2 | `aws-solution-architect` | Arquitetura AWS específica | AWS |
| 3 | `senior-devops` | Supervisão de infraestrutura | DevOps |
| 4 | `cloud-design-patterns` | Padrões de design cloud | Cloud |
| 5 | `kubernetes-architect` | Arquitetura de clusters K8s | K8s |
| 6 | `network-engineer` | Redes, VPC, DNS, CDN | Rede |
| 7 | `aws-well-architected-review` | Revisão Well-Architected | AWS |
| 8 | `disaster-recovery` | Estratégia de DR e backup | DR |
| 9 | `cost-optimization` | Otimização e estimativa de custos | Custos |
| 10 | `hybrid-cloud-networking` | Redes híbridas (on-prem to cloud) | Rede |
| 11 | `container-security-hardening` | Hardening de containers | Segurança |
| 12 | `mermaid-expert` | Diagramas de topologia | Diagramas |
| 13 | `documentation-writer` | Redigir o Infra Cloud Definition | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt gerador da definição de infraestrutura cloud | Time de Arquitetura |

---

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` | Definição de Arquitetura (F2) — topologia e redes |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` | Definição de Segurança (F3) — firewall, WAF |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` | Definição de Dados (F4) — storage |
| 4 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` | Definição DevOps/SRE (F5) — clusters K8s |
| 5 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 6 | `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (se fornecidos) |
| 7 | `{PROJECT-TEAM-CAPACITY}` | Capacidade do time (se fornecida) |
| 8 | `{PROJECT-STACK}` | Stack tecnológica (validar contra padrões corporativos) |
| 9 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
