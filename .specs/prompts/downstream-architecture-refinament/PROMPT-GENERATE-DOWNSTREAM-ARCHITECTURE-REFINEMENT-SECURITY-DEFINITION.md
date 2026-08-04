# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — SECURITY-DEFINITION (F3)
## Versão: 1.1 — Segurança Detail-Level (STRIDE + OWASP ASVS + IAM Specs) — Independente de Tecnologia

Atue como um Security Architect especializado em segurança de aplicações multi-tenant e análise de ameaças.

## OBJETIVO

Produzir a definição de segurança em nível de implementação: threat model STRIDE por componente, matriz de controles OWASP ASVS L1+L2, especificação de IAM, matriz de autorização granular e conformidade regulatória.

**Este documento é independente de tecnologias específicas de segurança.** Durante a análise da stack de segurança do projeto, identifique as tecnologias utilizadas (IAM, API Gateway, WAF) e busque skills relacionados. Caso não encontre skills específicos, utilize skills generalistas de segurança, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

## INPUTS

1. **Arquitetura Detail-Level** (F2)
2. **GLOBAL-SECURITY.md** — padrões corporativos de segurança
3. **Features de segurança:** EP-0003 (RBAC) e aspectos de autenticação cross-épico (Keycloak, OIDC, MFA)

## ESTRUTURA DO DOCUMENTO

```markdown
# DETAIL-LEVEL-SECURITY-DEFINITION — Segurança Detail-Level

## 1. Threat Model (STRIDE por Componente)
| Componente | Spoofing | Tampering | Repudiation | Info Disclosure | DoS | Elevation | Mitigação |

## 2. Matriz de Controles (OWASP ASVS L1+L2)
| ASVS ID | Categoria | Controle | Como Implementar |

## 3. Especificação IAM
- Realms/Tenants, Clients, Protocol Mappers
- Claims: roles, permissions, tenant_id, business_unit_ids
- Fluxos de autenticação e autorização
- Política de Senhas e MFA

## 4. Matriz de Autorização Granular
[Role × Permission × Resource — tabela completa de permissões]

## 5. Data Protection
[Criptografia em repouso e trânsito, isolamento de dados, masking]

## 6. Compliance Regulatória
[Mapeamento: requisito regulatório → controle implementado]

## 7. Riscos de Segurança

## 8. Padrões Corporativos FBSO

[Mapeie cada padrão abaixo para os controles de segurança desta disciplina (IAM, criptografia, compliance).]

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

🤖 *Documento gerado pelo Security Architect — Fase 3 do Downstream Architecture Refinement · Skills utilizados: [lista de skills efetivamente acionados] · Padrões Corporativos FBSO.ORG*
```

### Skills Recomendados

**Skills generalistas de segurança (sempre aplicáveis):**
- `engineering-skills`, `engineering-advanced-skills`
- `senior-security`, `security-best-practices`, `security-review`
- `security-reviewer`, `security-audit`, `security-auditor`
- `security-scanning-security-sast`, `security-threat-model`
- `threat-modeling-expert`, `threat-model-analyst`
- `privacy-by-design`, `gdpr-compliant`
- `secrets-management`, `secret-scanning`

**Skills tecnológicos de segurança (condicionais — buscar ao identificar a stack):**
- Ao identificar uma tecnologia específica de segurança (IAM, API Gateway, WAF) durante a análise da stack, procure skills relacionados a essa tecnologia para aprimorar as especificações
- Caso não encontre skills específicos para a tecnologia identificada, utilize os skills generalistas listados acima como referência, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

🤖 *Prompt gerador — Fase 3 do Downstream Architecture Refinement · Independente de Tecnologia · Skills: `engineering-skills`, `engineering-advanced-skills`, `senior-security`, `security-best-practices`, `security-review`, `security-audit`, `security-threat-model`, `threat-modeling-expert`, `privacy-by-design`, `gdpr-compliant`, `secrets-management`, `secret-scanning` · Padrões Corporativos FBSO.ORG*
