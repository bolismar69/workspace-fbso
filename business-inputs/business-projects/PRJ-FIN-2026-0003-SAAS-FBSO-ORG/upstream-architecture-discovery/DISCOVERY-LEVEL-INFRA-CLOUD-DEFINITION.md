# DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION — Definição de Infraestrutura Cloud (Discovery)
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F7 — Bloco B · **Disciplina:** Infra/Cloud Specialist
- **Versão:** 1.1 · **Data:** 30/07/2026 · **Status:** CREATED

## 1. Provedor e Topologia
| Decisão | Escolha | Justificativa |
|:---|:---|:---|
| Cloud Provider | **DigitalOcean** (padrão corporativo) | Kubernetes gerenciado (DOKS), banco gerenciado, custo previsível |
| CDN/WAF | **Cloudflare** (padrão corporativo) | Proxy reverso, DDoS protection, cache estático, SSL termination |
| API Gateway | **Kong** (padrão corporativo) | Self-hosted no Kubernetes, OIDC Plugin |
| IAM | **Keycloak** (padrão corporativo) | Autenticação OIDC, autorização RBAC |
| Compute | DOKS (DigitalOcean Kubernetes) | Orquestração de containers, auto-scaling |
| Banco | DigitalOcean Managed PostgreSQL | Alta disponibilidade gerenciada |

## 2. Recursos Estimados (Produção)
| Recurso | Especificação | Custo Mensal Est. |
|:---|:---|:---:|
| DOKS Kubernetes | 3× nodes (4 vCPU, 8GB) | ~$144 |
| Managed PostgreSQL | 2 vCPU, 8GB, 100GB SSD | ~$60 |
| Managed Redis | 1 vCPU, 2GB | ~$15 |
| Cloudflare | Pro plan (WAF + CDN) | ~$20 |
| Load Balancer | DO LB | ~$12 |
| **Total** | | **~$251/mês** |

## 3. Fluxo de Acesso
```
Usuário → Cloudflare (WAF/CDN) → Kong API Gateway → Backend API → PostgreSQL
                                                      ↘ Keycloak (auth)
```

## 4. Disaster Recovery
| Métrica | Alvo |
|:---|:---:|
| RPO | 1 hora (backup automatizado PostgreSQL) |
| RTO | 4 horas (DOKS node pools pré-configurados, infra como código) |
| Estratégia | Pilot Light — recursos mínimos em região secundária |

## 5. Riscos de Infraestrutura
| Risco | Mitigação |
|:---|:---|
| DigitalOcean tem menos regiões que AWS | Usar região mais próxima (NYC); Cloudflare mitiga latência global |
| Complexidade Kubernetes para time reduzido | DOKS abstrai camada de controle; começar com poucos nós |
| Custo Cloudflare pode crescer com tráfego | Plano Pro fixo; revisar em 6 meses |

## 6. Estimativa de Esforço
| Atividade | Esforço |
|:---|:---:|
| Setup DigitalOcean (DOKS, VPC, networking) | 1 homem-mês |
| PostgreSQL + Redis provisionamento | 0.5 homem-mês |
| Cloudflare configuração (DNS, WAF, CDN) | 0.5 homem-mês |
| Kong deploy no DOKS + OIDC Plugin | 1 homem-mês |
| **Total Infra/Cloud** | **3 homem-mês** |

🤖 *Upstream Architecture Discovery — Fase 7*
