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

## 4. Instruções para Preenchimento do Schema (Schema Unificado v1.3)

A fábrica deve preencher o arquivo **ESTIMATION-SCHEMA.csv** (fornecido na Fase 2) seguindo o **DTA Estimation Schema — 20 colunas padronizadas** (conforme `.specs/standards/DTA-VALIDATION-STANDARDS.md` §2.5):

### Bloco A: Identificação

| # | Coluna | Descrição | Obrigatório |
|:---:|:---|:---|:---:|
| 1 | `fabrica` | Nome da fábrica | ✅ |
| 2 | `id_epico` | Identificador do épico (EP-0001 a EP-0004) | ✅ |
| 3 | `titulo` | Título do épico conforme PRD | ✅ |

### Bloco B: Escopo

| # | Coluna | Descrição | Obrigatório |
|:---:|:---|:---|:---:|
| 4 | `features_codigos` | Códigos das features (ex: FEAT-EP-0001-0001) | ✅ |
| 5 | `qtd_features` | Quantidade de features no épico | ✅ |
| 6 | `user_stories_codigos` | Códigos das User Stories (ex: US-0001 a US-0007) | ✅ |
| 7 | `qtd_user_stories` | Quantidade de User Stories no épico | ✅ |

> 💡 No modo Discovery, `features_codigos` e `user_stories_codigos` podem ser preenchidos com os épicos (nível de detalhe disponível).

### Bloco C: Horas (por disciplina)

| # | Coluna | Descrição | Obrigatório |
|:---:|:---|:---|:---:|
| 8 | `horas_dev` | Horas estimadas de desenvolvimento | ✅ |
| 9 | `horas_qa` | Horas estimadas de QA/testes | ✅ |
| 10 | `horas_arch` | Horas estimadas de arquitetura | ✅ |
| 11 | `horas_devops` | Horas estimadas de DevOps/SRE | ✅ |
| 12 | `horas_gestao` | Horas estimadas de gestão | ✅ |
| 13 | `total_horas` | Soma das horas (dev+qa+arch+devops+gestao) | ✅ |

### Bloco D: Prazo, Time e Valor (OBRIGATÓRIOS)

| # | Coluna | Descrição | Obrigatório |
|:---:|:---|:---|:---:|
| 14 | `prazo_entrega_meses` | Prazo de entrega em meses (ex: 4.0) | ✅ |
| 15 | `time_estimado_pessoas` | Tamanho do time estimado (nº de pessoas) | ✅ |
| 16 | `valor_estimado` | Valor total estimado (R$) | ✅ |

> ⚠️ **Importante:** `time_estimado_pessoas` e `valor_estimado` são **obrigatórios**. A FBSO.ORG **NÃO infere** esses valores — devem ser declarados pela fábrica.

### Bloco E: Metadados

| # | Coluna | Descrição | Obrigatório |
|:---:|:---|:---|:---:|
| 17 | `complexidade` | Alta / Média / Baixa | ✅ |
| 18 | `stack_aderencia` | Aderência à stack proposta (Alta / Média / Baixa) | ✅ |
| 19 | `premissas` | Premissas consideradas na estimativa | ✅ |
| 20 | `comentarios` | Observações técnicas, metodologia, riscos | ✅ |

**Regras de validação DTA que serão aplicadas (F5):**
- **QA Balanceado:** QA ≥ 20% do desenvolvimento por épico; QA ≥ 25% do total global
- **Arquitetura:** Arch ≥ 5% do total geral de horas
- **Consistência Prazo×Horas:** `prazo_calculado = total_horas / (time × 160h)`. Divergência > 50% entre prazo declarado e calculado → rejeitada
- **Outliers:** Total de horas dentro de ±50% da mediana cross-fábrica
- **PIB (Proximidade à Baseline Interna):** Comparação com estimativa de referência da FBSO.ORG. PIB Score < 0.25 → rejeitada
- **Formato:** Todas as 20 colunas obrigatórias preenchidas; separador ponto-e-vírgula (;)

---

## 5. Critérios de Avaliação (5 critérios — DTA-VALIDATION-STANDARDS §3.2)

| # | Critério | Peso | Descrição |
|:---:|:---|:---:|:---|
| 1 | **Custo Total** | 25% | `valor_estimado` declarado pela fábrica |
| 2 | **Prazo de Entrega** | 25% | `prazo_entrega_meses` declarado pela fábrica |
| 3 | **Qualidade Técnica (QA+Arch)** | 20% | Percentual de QA e Arquitetura sobre o total de horas |
| 4 | **PIB — Proximidade à Baseline Interna** 🆕 | 15% | Comparação com estimativa de referência da FBSO.ORG (baseline interna não divulgada) |
| 5 | **Consistência Prazo×Horas** | 15% | Coerência entre `total_horas`, `prazo_entrega_meses` e `time_estimado_pessoas` |

> 💡 O critério PIB compara a estimativa da fábrica com uma **baseline interna da FBSO.ORG** (não divulgada no RFQ para evitar viés de ancoragem). Fábricas que se aproximam da baseline recebem pontuação mais alta neste critério. |

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
