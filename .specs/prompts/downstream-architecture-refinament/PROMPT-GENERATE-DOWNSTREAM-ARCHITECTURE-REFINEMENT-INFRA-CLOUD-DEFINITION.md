# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — INFRA-CLOUD-DEFINITION (F7)
## Versão: 1.1 — Infra/Cloud Detail-Level (Sizing + Custos + Topologia) — Específico do Projeto

Atue como um Cloud Architect e Infrastructure Specialist especializado em DigitalOcean e Cloudflare.

> ⚠️ **Nota:** As tecnologias listadas neste prompt (DigitalOcean, Cloudflare, DOKS) são específicas do projeto atual. Para executar este prompt em outro projeto, ajustar conforme a stack de infraestrutura do projeto alvo.

## OBJETIVO

Produzir a definição de infraestrutura cloud em nível de implementação: sizing detalhado, cálculo de custos, topologia de rede, disaster recovery.

## INPUTS

1. **Arquitetura Detail-Level** (F2)
2. **DevOps/SRE Detail-Level** (F5)
3. **Segurança Detail-Level** (F3)
4. **Infra/Cloud Discovery-Level** (se existir)
5. **GLOBAL-SECURITY.md**

## ESTRUTURA DO DOCUMENTO

```markdown
# DETAIL-LEVEL-INFRA-CLOUD-DEFINITION — Infra/Cloud Detail-Level

## 1. Topologia de Rede
[Cloudflare → Kong → DOKS (Istio) → Backend/PostgreSQL/Redis/Keycloak]
[VPC, subnets, security groups, network policies]

## 2. Sizing Detalhado
| Recurso | Tier | vCPU | RAM | Storage | Justificativa |
| DOKS Node Pool | | | | | |
| PostgreSQL | | | | | |
| Redis | | | | | |
| Keycloak | | | | | |

## 3. Cálculo de Custos Mensais
| Provedor | Serviço | Custo/mês |
| DigitalOcean | DOKS | |
| DigitalOcean | PostgreSQL | |
| DigitalOcean | Redis | |
| Cloudflare | CDN + WAF | |
| **Total** | | |

## 4. Fluxo de Acesso
[Diagrama: usuário → Cloudflare DNS → Cloudflare WAF → Kong → Istio → Backend → PostgreSQL]

## 5. Disaster Recovery
- RPO: 1h (PostgreSQL WAL shipping)
- RTO: 4h (DOKS cluster recreation via Terraform)
- Backup: pg_dump diário + WAL contínuo, retenção 30 dias
- DR Drill: trimestral

## 6. Ambientes
| Ambiente | Provedor | Sizing | Custo |
| Dev | Docker Compose local | — | $0 |
| Staging | DOKS (1 node) | | |
| Produção | DOKS (3 nodes) | | |

## 7. Riscos de Infraestrutura

## 8. Padrões Corporativos FBSO

[Mapeie cada padrão abaixo para sizing, topologia e DR desta disciplina.]

As tecnologias e padrões abaixo são **definições corporativas da FBSO.ORG** e constituem a baseline obrigatória para este projeto. Qualquer tecnologia adicional utilizada além destas deve ser explicitamente documentada com sua justificativa técnica e aprovada pelo time de arquitetura.

### Cloud & Edge
| Padrão | Tecnologia | Aplicação no Projeto |
|:---|:---|:---|
| Cloud Provider | **DigitalOcean** | Provedor único — DOKS (Kubernetes), PostgreSQL managed, Redis managed, Spaces (S3) |
| Edge/CDN/WAF | **Cloudflare** | DNS, CDN, WAF, SSL termination, DDoS protection — toda entrada de tráfego externo passa pela Cloudflare |

### Autenticação & Autorização
| Padrão | Tecnologia | Aplicação no Projeto |
|:---|:---|:---|
| IAM | **Keycloak** | Autenticação OIDC (frontend) + emissão JWT. Realms por tenant. Provisioning automático |
| API Gateway | **Kong** | Gateway central: valida JWT (Service-ID/Token-ID via Keycloak), injeta headers, rate limiting, routing |
| Integração Kong↔Keycloak | **Service-ID/Token-ID** | Kong encaminha Service-ID/Token-ID ao Keycloak para validação. Keycloak retorna autenticação + roles + acessos do usuário sistêmico. Kong alimenta cabeçalho JWT com atributos recebidos. **Microserviços NÃO revalidam JWT** — premissa: toda comunicação com microserviços passa pelo Kong |

### Banco de Dados
| Padrão | Tecnologia | Aplicação no Projeto |
|:---|:---|:---|
| Banco Relacional | **PostgreSQL** | PostgreSQL 17. Schema `fbso_portal` com RLS multi-tenant. Schemas auxiliares: `public`, `keycloak` |

### SRE & Observabilidade
| Padrão | Tecnologia | Aplicação no Projeto |
|:---|:---|:---|
| Métricas | **Prometheus** | Coleta de métricas via exporters (Kong, PostgreSQL, Keycloak) + Micrometer (backend) |
| Dashboards | **Grafana** | Visualização unificada — datasources: Prometheus, Loki, PostgreSQL |
| Logs | **Grafana Loki** | Agregação de logs de aplicação e infraestrutura |
| Tracing | **Jaeger** | Distributed tracing — spans do OpenTelemetry exportados para Jaeger |
| Instrumentação | **OpenTelemetry** | Auto-instrumentação backend (Java agent) + traces manuais em pontos críticos |
| Logs de Auditoria | **Elastic Stack** | Complementar ao Loki para logs de auditoria com retenção longa e busca full-text |

### Infraestrutura como Código (IaC)
| Padrão | Tecnologia | Aplicação no Projeto |
|:---|:---|:---|
| Provisioning | **Terraform** | DOKS cluster, PostgreSQL, Redis, Spaces, networking, secrets |
| Configuração | **Ansible** | Provisioning de nós, configuração Kong, agentes de monitoramento |

### DevOps & Orquestração
| Padrão | Tecnologia | Aplicação no Projeto |
|:---|:---|:---|
| Orquestração | **Kubernetes (DOKS)** | Cluster gerenciado DigitalOcean — todos os workloads em containers |
| Service Mesh | **Istio** | mTLS entre serviços, controle de tráfego (canary, blue-green), observabilidade sidecar |
| Autoscaling (Pods) | **Keda** | Kubernetes Event-Driven Autoscaling — escala pods baseado em eventos (fila, métricas) |
| Autoscaling (Nodes) | **Karpenter** | Cluster Autoscaling — adiciona/remove nós automaticamente conforme demanda |

### CI/CD
| Padrão | Tecnologia | Aplicação no Projeto |
|:---|:---|:---|
| CI/CD | **GitHub Actions** | Build, SAST (Semgrep), Secret Scanning (Gitleaks), Docker build, deploy via kubectl/Helm |

### Tecnologias Adicionais
| Tecnologia | Justificativa | Aprovação |
|:---|:---|:---|
[Vazio por padrão — preencher apenas se tecnologias fora da baseline forem detectadas]

> ⚠️ **Regra de Compliance:** Tecnologias detectadas durante a análise que NÃO constam nesta lista de padrões corporativos devem ser explicitamente documentadas com justificativa técnica e aprovadas pelo time de arquitetura na seção "Tecnologias Adicionais".

🤖 *Documento gerado pelo Infra/Cloud Specialist — Fase 7 do Downstream Architecture Refinement · Skills utilizados: [lista de skills efetivamente acionados] · Padrões Corporativos FBSO.ORG*
```

### Skills Recomendados
- `engineering-skills`, `engineering-advanced-skills`
- `cloud-architect`, `aws-solution-architect` (padrões transferíveis)
- `hybrid-cloud-networking`, `kubernetes-architect`
- `disaster-recovery` (se disponível)

🤖 *Prompt gerador — Fase 7 do Downstream Architecture Refinement · Skills: `engineering-skills`, `engineering-advanced-skills`, `cloud-architect`, `aws-solution-architect`, `hybrid-cloud-networking`, `kubernetes-architect` · Padrões Corporativos FBSO.ORG*
