# RFQ-PACKAGE — Request for Quotation (Discovery-Level)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Fase:** F1 — Bloco A (RFQ Package) · Mode 1: Discovery-Level
- **Versão:** 1.0 · **Data:** 31/07/2026 · **Status:** CREATED
- **Pasta:** `sourcing-factory-bidding-discovery/`

---

## 1. Carta Convite

Prezada Fábrica de Software,

A **FBSO.ORG** está em processo de seleção de parceiro de desenvolvimento para a construção da **FBSO Platform** — portal administrativo SaaS multi-tenant que servirá como plataforma Core para produtos futuros.

Solicitamos que V.Sa. analise o material técnico anexo e apresente estimativa de esforço conforme o **DTA Estimation Schema** (template CSV fornecido).

**Prazo para resposta:** 10 dias úteis a partir do recebimento.

---

## 2. Resumo Executivo do Projeto

A FBSO Platform é um portal administrativo SaaS multi-produto com RBAC. A plataforma permite:

- **Gestão de clientes e assinaturas** — cadastro, ativação, planos comerciais
- **Portal administrativo interno** — dashboard operacional, métricas, alertas
- **Governança de acessos (RBAC)** — papéis e permissões granulares por tenant
- **Portal de autoatendimento do cliente** — gestão autônoma de conta e usuários

**Métricas de negócio:** ~500 tenants (ano 1) · ~5.000 usuários · 4 épicos · múltiplos produtos como módulos ativáveis

---

## 3. Artefatos Técnicos Anexados

| # | Artefato | Conteúdo | Fase |
|:---|:---|:---|:---:|
| 1 | [DISCOVERY-LEVEL-PRD](../upstream-architecture-discovery/DISCOVERY-LEVEL-PRD.md) | Visão do produto, épicos, MVP macro, restrições | F1 |
| 2 | [DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION](../upstream-architecture-discovery/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md) | C4 System Context, containers, ADRs, matriz integração | F2 |
| 3 | [DISCOVERY-LEVEL-SECURITY-DEFINITION](../upstream-architecture-discovery/DISCOVERY-LEVEL-SECURITY-DEFINITION.md) | Threat model, compliance (LGPD), Kong trust boundary | F3 |
| 4 | [DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION](../upstream-architecture-discovery/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md) | Entidades macro, volumes, PostgreSQL/Redis | F4 |
| 5 | [DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION](../upstream-architecture-discovery/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md) | CI/CD, IaC, stack SRE, Istio/Keda/Karpenter | F5 |
| 6 | [DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION](../upstream-architecture-discovery/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md) | Pirâmide de testes, quality gates | F6 |
| 7 | [DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION](../upstream-architecture-discovery/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md) | DigitalOcean + Cloudflare, topologia, custos | F7 |
| 8 | [DISCOVERY-LEVEL-SOLUTIONS-CATALOG](../upstream-architecture-discovery/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md) | 15 soluções catalogadas | F8 |
| 9 | [DISCOVERY-LEVEL-SOLUTIONS-MATRIX](../upstream-architecture-discovery/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md) | Matriz 15 soluções × 6 disciplinas | F9 |
| 10 | [DISCOVERY-LEVEL-SPECS](../upstream-architecture-discovery/DISCOVERY-LEVEL-SPECS.md) | Especificação técnica consolidada + stack corporativa | F10 |

---

## 4. Instruções para Preenchimento do Schema

A fábrica deve preencher o arquivo **ESTIMATION-SCHEMA.csv** (fornecido na Fase 2) seguindo o formato DTA Estimation Schema:

| Coluna | Descrição | Obrigatório |
|:---|:---|:---:|
| `id_epico` | Identificador do épico (EP-0001 a EP-0004) | ✅ |
| `titulo_epico` | Título do épico conforme PRD | ✅ |
| `solucoes` | Soluções envolvidas (IDs S01-S15) | ✅ |
| `horas_desenvolvimento` | Horas estimadas de desenvolvimento | ✅ |
| `horas_arquitetura` | Horas estimadas de arquitetura/SRE | ✅ |
| `horas_qa` | Horas estimadas de QA/testes | ✅ |
| `prazo_entrega_meses` | Prazo de entrega em meses (ex: 6.5 para 6 meses e 2 semanas) | ✅ |
| `complexidade` | Alta / Média / Baixa | ✅ |
| `comentarios` | Observações técnicas relevantes | — |

**Regras de validação DTA que serão aplicadas:**
- QA ≥ 20% do total de horas de desenvolvimento por épico
- Arquitetura/SRE ≥ 5% do total geral
- Outliers detectados por comparação cross-fábrica

---

## 5. Critérios de Avaliação

| Critério | Peso | Descrição |
|:---|:---:|:---|
| **Custo total** | 35% | Homem-hora × taxa informada |
| **Prazo estimado** | 25% | Tempo total para conclusão |
| **Qualidade técnica** | 25% | Completude dos comentários, aderência ao schema |
| **QA/Arch balanceado** | 15% | Proporção adequada de QA e Arquitetura |

---

## 6. Prazo de Resposta

- **Data de envio:** 31/07/2026
- **Prazo limite:** 14/08/2026 (10 dias úteis)
- **Formato de devolução:** Arquivo CSV preenchido conforme ESTIMATION-SCHEMA.csv

---

## Histórico de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial: RFQ Package Discovery-Level — 10 artefatos anexados | PMO / Time de Arquitetura |

🤖 *Sourcing & Factory Bidding — Fase 1 · Mode 1 (Discovery-Level)*
