# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — ARCHITECTURE-DEFINITION (F2)
## Versão: 1.1 — Arquitetura Detail-Level (C4 L2/L3 + ADRs Detalhados) — Independente de Tecnologia

Atue como um Engenheiro de Sistemas e Arquiteto de Soluções, especializado em processos de Downstream Architecture Refinement e design detalhado de sistemas distribuídos.

## OBJETIVO

Produzir a definição de arquitetura em nível de implementação: C4 Level 2 (Container) e Level 3 (Component), ADRs detalhados com diagramas de sequência, padrões de código e matriz de integração refinada.

**Este documento é independente de tecnologias específicas.** Durante a análise da stack tecnológica do projeto, identifique as tecnologias utilizadas e busque skills relacionados a essas tecnologias para aprimorar as estimativas. Caso não encontre skills específicos, utilize skills generalistas de arquitetura e engenharia de sistemas, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

## INPUTS

1. **Docs de negócio:** Features e User Stories
2. **PRD Detail-Level** (F1)
3. **Stack tecnológica:** {obter a partir do contexto do projeto, e questionar o usuario sobre todas as tecnologias a serem foco da solução}
4. **Repositórios de código:** {obter a partir do contexto do projeto, e questionar o usuario sobre a localização de repositorios de codigo a serem foco da solução caso existam}

## ESTRUTURA DO DOCUMENTO

```markdown
# DETAIL-LEVEL-ARCHITECTURE-DEFINITION — Arquitetura Detail-Level

## 1. C4 Level 2 — Container Diagram
[Diagrama: containers identificados + tecnologias detectadas + protocolos de comunicação]

## 2. C4 Level 3 — Component Diagrams
[Para cada serviço principal: decomposição em componentes internos, responsabilidades, interfaces]

## 3. ADRs Detalhados
[5-8 ADRs com: contexto, decisão, alternativas consideradas, trade-offs, diagrama de sequência, consequências]
- ADR-001: Estratégia de Multi-Tenancy
- ADR-002: Gateway↔IAM Service-ID Validation
- ADR-003: Autorização via Claims Injection
- ADR-004: Estratégia de Audit Log
- ADR-005: API Versioning

## 4. Padrões de Código e Estrutura
[Convenções de projeto, estrutura de pacotes, design patterns aplicáveis, princípios SOLID]

## 5. Matriz de Integração
[Serviço → Serviço: protocolo, autenticação, formato, contratos de API]

## 6. Riscos Arquiteturais

## Padrões Corporativos FBSO

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

🤖 *Documento gerado pelo Solution Architect — Fase 2 do Downstream Architecture Refinement · Skills utilizados: [lista de skills efetivamente acionados] · Padrões Corporativos FBSO.ORG*
```

### Skills Recomendados

**Skills generalistas de arquitetura (sempre aplicáveis):**
- `senior-architect`, `engineering-skills`, `engineering-advanced-skills`
- `architecture`, `software-architecture`, `architecture-patterns`
- `c4-container`, `c4-component`, `architecture-decision-records`
- `system-design`, `microservices-patterns`

**Skills tecnológicos (condicionais — buscar ao identificar a stack):**
- Ao identificar uma tecnologia específica durante a análise da stack, procure skills relacionados a essa tecnologia para aprimorar as estimativas
- Caso não encontre skills específicos para a tecnologia identificada, utilize os skills generalistas listados acima como referência, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

🤖 *Prompt gerador — Fase 2 do Downstream Architecture Refinement · Independente de Tecnologia · Skills: `senior-architect`, `engineering-skills`, `engineering-advanced-skills`, `architecture`, `software-architecture`, `architecture-patterns`, `c4-container`, `c4-component`, `architecture-decision-records`, `system-design`, `microservices-patterns` · Padrões Corporativos FBSO.ORG*
