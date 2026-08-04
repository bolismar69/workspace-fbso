# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — DATA-ARCHITECTURE-DEFINITION (F4)
## Versão: 1.1 — Arquitetura de Dados Detail-Level (ERD + Query Patterns + Migrations) — Independente de Tecnologia

Atue como um Data Architect e Especialista em Modelagem de Dados, com expertise em sistemas multi-tenant e estratégias de isolamento de dados.

## OBJETIVO

Produzir a definição de arquitetura de dados em nível de implementação: modelo de dados completo, estratégia de isolamento multi-tenant, query patterns otimizadas, estratégia de versionamento de schema.

**Este documento é independente de tecnologias específicas de banco de dados.** Durante a análise da stack do projeto, identifique o banco de dados utilizado e busque skills relacionados a esse banco para aprimorar as especificações. Caso não encontre skills específicos, utilize skills generalistas de arquitetura de dados, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

## INPUTS

1. **Arquitetura Detail-Level** (F2)
2. **Segurança Detail-Level** (F3) — políticas de isolamento
3. **Features relevantes:** dashboards, tenants, RBAC, auditoria

## ESTRUTURA DO DOCUMENTO

```markdown
# DETAIL-LEVEL-DATA-ARCHITECTURE-DEFINITION — Arquitetura de Dados Detail-Level

## 1. Modelo de Dados Completo
[Todas as entidades: atributos, tipos, constraints, índices, relacionamentos]
Esquemas e domínios de dados

## 2. Estratégia de Multi-Tenancy
- Abordagem de isolamento (discriminator column, RLS, schema-per-tenant, database-per-tenant)
- Políticas de isolamento por tabela
- Verificação de isolamento

## 3. Estratégia de Particionamento e Arquivamento
- Tabelas de alto volume: estratégia de particionamento
- Política de retenção e arquivamento

## 4. Query Patterns Otimizadas
[Índices, views materializadas, estratégias de cache para queries críticas]

## 5. Estratégia de Versionamento de Schema
- Ferramenta de migration
- Versionamento e baseline
- Estratégia de rollback

## 6. Volumes e Crescimento
[Projeções: year 1-3, tenants, users, transactions, storage]

## 7. Riscos de Dados

## 8. Padrões Corporativos FBSO

[Mapeie cada padrão abaixo para as decisões de dados desta disciplina (modelo, RLS, migrações).]

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

🤖 *Documento gerado pelo Data Architect — Fase 4 do Downstream Architecture Refinement · Skills utilizados: [lista de skills efetivamente acionados] · Padrões Corporativos FBSO.ORG*
```

### Skills Recomendados

**Skills generalistas de dados (sempre aplicáveis):**
- `engineering-skills`, `engineering-advanced-skills`
- `senior-data-engineer`, `database-architect`, `database-design`
- `data-engineer`, `data-modeling`, `database`
- `database-migrations`, `database-migrations-sql-migrations`

**Skills tecnológicos de banco de dados (condicionais — buscar ao identificar a stack):**
- Ao identificar um banco de dados específico durante a análise da stack, busque skills relacionados a esse banco para aprimorar as especificações e estimativas
- Caso não encontre skills específicos para o banco identificado, utilize os skills generalistas listados acima como referência, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

🤖 *Prompt gerador — Fase 4 do Downstream Architecture Refinement · Independente de Tecnologia · Skills: `engineering-skills`, `engineering-advanced-skills`, `senior-data-engineer`, `database-architect`, `database-design`, `data-engineer`, `data-modeling`, `database`, `database-migrations`, `database-migrations-sql-migrations` · Padrões Corporativos FBSO.ORG*
