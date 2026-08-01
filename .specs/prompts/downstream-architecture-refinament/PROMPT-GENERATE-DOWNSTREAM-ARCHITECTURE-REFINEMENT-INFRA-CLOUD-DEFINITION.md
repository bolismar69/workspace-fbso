# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — INFRA-CLOUD-DEFINITION (F7)
## Versão: 1.0 — Infra/Cloud Detail-Level (Sizing + Custos + Topologia)

Atue como um Cloud Architect e Infrastructure Specialist especializado em DigitalOcean e Cloudflare.

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
```

### Skills Recomendados
- `cloud-architect`, `aws-solution-architect` (padrões transferíveis)
- `hybrid-cloud-networking`, `kubernetes-architect`
- `disaster-recovery` (se disponível)

🤖 *Prompt gerador — Fase 7 do Downstream Architecture Refinement*
