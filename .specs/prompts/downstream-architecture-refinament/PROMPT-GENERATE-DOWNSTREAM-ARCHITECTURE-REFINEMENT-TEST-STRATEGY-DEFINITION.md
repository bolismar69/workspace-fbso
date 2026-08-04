# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — TEST-STRATEGY-DEFINITION (F6)
## Versão: 1.1 — Estratégia de Testes Detail-Level (Matriz de Cobertura + Casos de Teste) — Independente de Tecnologia

Atue como um QA Lead e Test Specialist especializado em estratégias de teste para aplicações multi-tenant e sistemas distribuídos.

## OBJETIVO

Produzir a estratégia de testes em nível de implementação: matriz de cobertura por US, casos de teste de aceitação, estratégia de automação, quality gates.

**Este documento é independente de tecnologias específicas de teste.** Durante a análise da stack do projeto, identifique as ferramentas e frameworks de teste utilizados e busque skills relacionados. Caso não encontre skills específicos, utilize skills generalistas de qualidade e teste, utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

## INPUTS

1. **Arquitetura Detail-Level** (F2)
2. **Segurança Detail-Level** (F3)
3. **Todas as User Stories** (cenários de aceitação)

## ESTRUTURA DO DOCUMENTO

```markdown
# DETAIL-LEVEL-TEST-STRATEGY-DEFINITION — Estratégia de Testes Detail-Level

## 1. Pirâmide de Testes Refinada
- Testes Unitários: framework e ferramentas (meta ≥80%)
- Testes de Integração: framework e ferramentas (meta ≥60%)
- Testes E2E: ferramentas (fluxos críticos)
- Testes de Performance: ferramentas (carga, stress)
- Testes de Segurança: SAST, DAST, Secret Scanning

## 2. Matriz de Cobertura por US
| US-ID | Unitário | Integração | E2E | Performance | Segurança | Responsável |
[Uma linha por US]

## 3. Casos de Teste de Aceitação
[Baseados nos cenários de aceitação das US — 3-5 casos por feature]

## 4. Estratégia de Automação
- CI: testes unitários + integração a cada commit/PR
- Nightly: E2E + performance smoke tests
- Release: E2E completo + performance completo + security scan

## 5. Quality Gates
| Gate | Critério | Bloqueia? |
| PR | Unit ≥80%, Integ ≥60%, SAST limpo | Sim |
| Staging | E2E críticos passam, Perf smoke OK | Sim |
| Release | E2E 100%, Security 0 críticas | Sim |

## 6. Testes de Isolamento Multi-Tenant
[Cenários de verificação de isolamento entre tenants]

## 7. Riscos de Qualidade

## 8. Padrões Corporativos FBSO

[Mapeie cada padrão abaixo para a matriz de testes e quality gates desta disciplina.]

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

🤖 *Documento gerado pelo QA Lead / Test Specialist — Fase 6 do Downstream Architecture Refinement · Skills utilizados: [lista de skills efetivamente acionados] · Padrões Corporativos FBSO.ORG*
```

### Skills Recomendados

**Skills generalistas de teste e qualidade (sempre aplicáveis):**
- `engineering-skills`, `engineering-advanced-skills`
- `senior-qa`, `testing-patterns`, `test-strategy-design`
- `test-master`, `testing-qa`, `test-driven-development`
- `e2e-testing`, `e2e-testing-patterns`
- `test-automator`, `test-case-creation`
- `qa`, `qa-test-planner`, `acceptance-criteria`
- `performance-testing-review-multi-agent-review`

**Skills tecnológicos de teste (condicionais — buscar ao identificar a stack):**
- Ao identificar um framework ou ferramenta de teste específica durante a análise da stack, busque skills relacionados a essa tecnologia para aprimorar as especificações de teste
- Caso não encontre skills específicos para a ferramenta identificada, utilize os skills generalistas listados acima como referência, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

🤖 *Prompt gerador — Fase 6 do Downstream Architecture Refinement · Independente de Tecnologia · Skills: `engineering-skills`, `engineering-advanced-skills`, `senior-qa`, `testing-patterns`, `test-strategy-design`, `test-master`, `testing-qa`, `test-driven-development`, `e2e-testing`, `e2e-testing-patterns`, `test-automator`, `test-case-creation`, `qa`, `qa-test-planner`, `acceptance-criteria` · Padrões Corporativos FBSO.ORG*
