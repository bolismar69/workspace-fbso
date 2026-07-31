# DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION — Definição de Arquitetura de Dados (Discovery)
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F4 — Bloco B · **Disciplina:** Data Architect
- **Versão:** 1.0 · **Data:** 30/07/2026 · **Status:** CREATED

## 1. Modelo de Dados (Entidades Macro)
| Entidade | Descrição | Volume Estimado |
|:---|:---|:---:|
| TENANT | Conta do cliente | ~500 (ano 1) |
| USER | Usuários do sistema | ~5.000 (10 por tenant) |
| SUBSCRIPTION | Assinatura ativa | ~500 |
| BUSINESS_UNIT | Filiais do cliente | ~1.500 (3 por tenant) |
| AUDIT_LOG | Registro de auditoria | ~1M registros/ano |

## 2. Estratégia de Armazenamento
| Camada | Tecnologia | Propósito |
|:---|:---|:---|
| OLTP | PostgreSQL 17 | Dados transacionais |
| Cache | Redis | Sessões, rate limiting, métricas em tempo real |
| Logs/Auditoria | PostgreSQL + S3 (archive) | Retenção 5 anos |

## 3. Volumes e Crescimento
| Métrica | Ano 1 | Ano 2 |
|:---|:---:|:---:|
| Tenants | 500 | 2.000 |
| Tamanho do banco | ~50GB | ~200GB |
| Transações/dia | ~50K | ~200K |

## 4. Riscos de Dados
| Risco | Mitigação |
|:---|:---|
| Crescimento do banco impacta performance | Particionamento por tenant_id; read replicas |
| Volume de audit_log | Archive automático para S3 após 90 dias |
| Consistência em operações cross-tenant | Transações ACID; Row-Level Security |

## 5. Estimativa de Esforço
| Atividade | Esforço |
|:---|:---:|
| Schema design + migrações Flyway | 1-2 homem-mês |
| Configuração RLS + índices | 0.5-1 homem-mês |
| Pipeline auditoria → S3 | 0.5 homem-mês |
| **Total Dados** | **2-3.5 homem-mês** |

🤖 *Upstream Architecture Discovery — Fase 4*
