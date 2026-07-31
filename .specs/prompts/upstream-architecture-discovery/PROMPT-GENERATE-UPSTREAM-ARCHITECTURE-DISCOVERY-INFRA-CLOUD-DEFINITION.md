# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION

## Contexto

Este prompt gera o artefato `UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION.md` — a **definição de infraestrutura cloud/on-premise do projeto** que especifica topologia, compute, networking, storage, disaster recovery, dimensionamento e custos. Este documento é a referência técnica para toda a infraestrutura que suporta as soluções do projeto.

**Relação com DEVOPS-SRE-DEFINITION:** Enquanto o DEVOPS-SRE-DEFINITION (F10) define o pipeline CI/CD e a operação das soluções, este documento detalha **o substrato de infraestrutura sobre o qual essas soluções e pipelines são executados** — VMs, clusters, redes, storage, DR.

**Inputs upstream:** `UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md` (Fase 7) + `UPSTREAM-ARCHITECTURE-DISCOVERY-SECURITY-DEFINITION.md` (Fase 8) + `UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION.md` (Fase 9) + `UPSTREAM-ARCHITECTURE-DISCOVERY-DEVOPS-SRE-DEFINITION.md` (Fase 10).

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
Ler Architecture Definition (F7 — topologia de containers, redes, ambientes), Security Definition (F8 — firewall, IAM de infra, WAF, criptografia), Data Architecture (F9 — requisitos de storage para bancos e pipelines), DevOps SRE (F10 — clusters K8s, IaC, observabilidade), ADRs globais.

### Passo 2 — Invocar Skills Especializadas
Invocar skills de arquitetura cloud, infraestrutura, redes, storage, DR e custos para projetar a infraestrutura completa.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-INFRA-CLOUD-DEFINITION.md` com:

1. **Topologia de Infraestrutura** — On-Premise e/ou Cloud (AWS/Azure/GCP), regiões, zonas de disponibilidade, ambientes (Dev/Staging/Prod), diagrama de topologia
2. **Compute** — VMs (EC2/Azure VM/GCE — famílias, sizing), Kubernetes (EKS/AKS/GKE — node pools, auto-scaling), Serverless (Lambda/Functions/Cloud Run), justificativa de escolha
3. **Networking** — VPC/VNet, subnets (públicas/privadas), DNS (Route53/Azure DNS), CDN (CloudFront/Cloudflare), API Gateway, load balancers (ALB/NLB/Application Gateway), NAT, VPN/Direct Connect
4. **Storage** — Block (EBS/Managed Disks), Object (S3/Blob Storage), File (EFS/Azure Files/NFS), hierarquia de storage por performance/custo
5. **Disaster Recovery** — RPO, RTO por solução, estratégia de backup (schedule, retenção), multi-region ativo-ativo ou ativo-passivo, failover automático
6. **Dimensionamento** — sizing inicial (compute, memória, storage), auto-scaling policies (horizontal/vertical), limites e quotas, plano de crescimento
7. **Segurança de Infra** — WAF, security groups, NACLs, IAM de infraestrutura (roles, policies), encryption at rest (KMS) e in transit (TLS), network policies
8. **Estimativa de Custos** — calculadora por provedor (AWS Pricing Calculator/Azure Pricing), custo mensal estimado por ambiente, breakdown por serviço, opções de reserva (RIs/Savings Plans)

### Passo 4 — Validação Pós-Geração
Verificar: topologia definida para todos os ambientes, compute dimensionado e justificado, networking completo (VPC, subnets, DNS, CDN), DR com RPO/RTO por solução, sizing com auto-scaling, segurança de infra alinhada com SECURITY-DEFINITION, estimativa de custos presente.

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

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
