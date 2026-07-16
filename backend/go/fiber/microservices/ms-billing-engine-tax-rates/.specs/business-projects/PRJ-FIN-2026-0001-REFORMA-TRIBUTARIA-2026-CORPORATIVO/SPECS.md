# SPECS — Especificações do Projeto

- **Projeto:** PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO
- **Programa:** Adequação Corporativa à Reforma Tributária Nacional
- **Microserviço-base:** `ms-billing-engine-tax-rates`
- **Versão do Documento:** 2.0
- **Data:** 30 de Junho de 2026
- **Status:** ✅ 10 GAPs implementados — PR #6 merged (Fases 0-1-2)

- 📋 **Propósito:** Este documento é o **índice-mestre de especificações** do projeto. Ele referencia e conecta três fontes de documentação: (1) os artefatos de produto deste projeto, (2) os documentos de definição de negócio, e (3) as especificações técnicas do microserviço em `.specs/`. Use-o como ponto de partida para navegar entre qualquer artefato do projeto.

---

## 1. Estrutura de Documentação do Projeto

```
PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/
│
├── 📄 SPECS.md          ← VOCÊ ESTÁ AQUI (índice-mestre)
├── 📄 PRD.md              Resumo de alto nível do produto (porta de entrada)
├── 📄 ARCHITECTURE.md     Decisões arquiteturais específicas do projeto
├── 📄 TASKS.md            Tarefas de implementação e runbook operacional
├── 📄 TEST_PLAN.md        Estratégia e cenários de teste por GAP
│
├── 🔗 business-inputs/    Documentos de definição de negócio (fonte primária)
│   └── PRJ-FIN-2026-0001/
│       ├── 01-PROJECT-CHARTER.md
│       ├── 02-BUSINESS-REQUIREMENTS.md
│       ├── 03-EPICS.md (1 arquivo unificado — 6 épicos)
│       ├── 04-FEATURES.md (1 arquivo unificado — 16 features)
│       ├── 05-USER-STORYS-*.md (16 arquivos)
│       ├── MATRIZ-KPI.md
│       ├── DEFINITION_OF_DONE.md
│       ├── GLOSSARY.md
│       ├── STAKEHOLDER-MAP.md
│       └── docs-suporte/ (16 arquivos técnico-tributários)
│
└── 🔗 .specs/             Especificações técnicas do microserviço (fonte da verdade)
    ├── architecture/ (4 arquivos)
    ├── domain/ (1 arquivo)
    ├── product/ (2 arquivos)
    ├── engineering/ (2 arquivos)
    ├── api/ (1 arquivo)
    ├── features/ (2 arquivos)
    ├── governance/ (2 arquivos)
    ├── questions/ (1 arquivo)
    └── skill-output/ (15 registros de implementação)
```

---

## 2. Escopo de Implementação no Microserviço

Esta seção define **exatamente o que será construído, modificado ou integrado** no microserviço `ms-billing-engine-tax-rates` para atender ao projeto PRJ-FIN-2026-0001. O detalhamento técnico de cada item está no [ARCHITECTURE.md](./ARCHITECTURE.md).

### 2.1 O Que JÁ EXISTE e Atende ao Projeto (Reaproveitamento)

Estas capacidades **não exigem nova implementação** — já estão prontas e testadas:

| # | Capacidade | Cobre BR(s) | Local no Código | Testes |
|:---|:---|:---|:---|:---|
| ✅ E-01 | Motor 7-fases SOP-013 (`BillingEnginePhased`) | BR-01, BR-04, BR-07 | `internal/calculator/engine.go` | 28 |
| ✅ E-02 | Cálculo CBS (federal, "por fora") | BR-04 | `internal/reforma/cbs_calculator.go` | 7 |
| ✅ E-03 | Cálculo IBS (estadual + municipal por destino) | BR-05 | `internal/reforma/ibs_calculator.go` | 7 |
| ✅ E-04 | Cálculo IS — Imposto Seletivo (pré-filtro NCM) | IS Compliance | `internal/legacy/is_filter.go` | 8 |
| ✅ E-05 | Phase Resolution System (transição 2026–2033) | Transição | `internal/phase/phase.go` | 14 |
| ✅ E-06 | TaxSelector — Matriz DT-001 (tributos ativos por fase) | Transição | `internal/phase/tax_selector.go` | 5 |
| ✅ E-07 | ICMS Desonerado (2 modos: Redução Base + Limitação Alíquota) | Regimes Especiais | `internal/legacy/icms_desoneracao.go` | 11 |
| ✅ E-08 | Tabela `iva_dual_rules` + trigger auditoria + cache Redis | BR-01, BR-02 | `data/init.sql`, `repository/` | — |
| ✅ E-09 | IBS Circuit Breaker + Cache Fallback | Resiliência | `internal/circuitbreaker/`, `internal/ibsclient/` | 12 |
| ✅ E-10 | W3C Trace Context (rastreabilidade ponta a ponta) | BR-07 | `internal/middleware/requestid.go` | 12 |
| ✅ E-11 | Métricas Prometheus (`/metrics`) | KPI O2 | `internal/middleware/metrics.go` | — |
| ✅ E-12 | Health Checks (`/healthz`, `/health`) | Operação | `cmd/api/main.go:81-116` | — |
| ✅ E-13 | Auth JWT (Kong/Keycloak pass-through) | Segurança | `internal/middleware/auth.go` | 9 |
| ✅ E-14 | PostgreSQL schema (10 tabelas, índices, unique constraints) | Persistência | `data/init.sql` | — |
| ✅ E-15 | Cálculo ICMS (Próprio, ST, DIFAL, Simples Nacional) | Legado ativo | `internal/legacy/icms.go` | 12 |
| ✅ E-16 | Cálculo PIS/COFINS (15 estratégias CST) | Legado ativo | `internal/legacy/pis_cofins.go` | 39 |
| ✅ E-17 | Cálculo IPI (Ad Valorem + Ad Pauta) | Legado ativo | `internal/legacy/ipi.go` | 7 |
| ✅ E-18 | Cálculo ISS + FUST + FUNTTEL | Serviços/Telecom | `internal/legacy/iss.go`, `fust.go`, `funttel.go` | 17 |

### 2.2 O Que SERÁ CONSTRUÍDO (Escopo de Implementação)

Estes são os **gaps que exigem desenvolvimento** no microserviço. Ordenados por prioridade:

| # | Gap | O que Implementar | BR | Prioridade | Esforço |
|:---|:---|:---|:---|:---|:---|
| 🔨 GAP-001 | **Interface Admin para Time Fiscal** | Endpoint `PUT/PATCH /admin/tax-rates` com autenticação, validação e log de auditoria. Permite que o time fiscal atualize alíquotas CBS/IBS/IS sem deploy | BR-02 | 🔴 Alta | 5–8 dias |
| 🔨 GAP-002 | **TaxToken — Congelamento de Alíquota** | Estrutura `TaxToken` (tupla NCM+UF+IBGE+aliquotas+TTL). Endpoints `POST /token/generate` e suporte a `token_id` no `POST /calculate` | BR-06 | 🔴 Alta | 5–8 dias |
| 🔨 GAP-003 | **Endpoint de Simulação de Margem** | `POST /simulate` — retorna impacto dos tributos por UF/município, margem líquida projetada, sem persistir transação | BR-05 | 🟡 Média | 3–5 dias |
| 🔨 GAP-004 | **Campo `valor_liquido` no Response** | Adicionar `valor_liquido` (preço base sem tributos) ao schema `DocumentoFiscalSaida` para UI de checkout "por fora" | BR-04 | 🟡 Média | 1–2 dias |
| 🔨 GAP-005 | **Cálculo de Créditos na Entrada** | `POST /credit/calculate` — calcula créditos CBS/IBS apropriáveis de documentos fiscais de compra (fornecedores) conforme regras do Lucro Real | BR-08 | 🔴 Alta | 8–12 dias |
| 🔨 GAP-006 | **Schema Split Payment no Response** | Adicionar campos `valor_receita_liquida`, `valor_cbs_reter`, `valor_ibs_reter`, `valor_is_reter` ao response para instrução bancária | BR-09 | 🔴 Alta | 3–5 dias |
| 🔨 GAP-007 | **Qualificação Fiscal de Fornecedores** | Tabela `fornecedor_fiscal` + endpoint `POST /supplier/validate` — due diligence fiscal para crédito (regime tributário, certidões, compliance) | BR-08 | 🟡 Média | 8–12 dias |
| 🔨 GAP-008 | **Rate Limiting** | Middleware Fiber com configuração por env var (`RATE_LIMIT_MAX`, `RATE_LIMIT_WINDOW`) — proteção contra abuso/DoS (DT-11) | RNF | 🔴 Alta | 2–3 dias |
| 🔨 GAP-009 | **API Versioning** | Migrar endpoint atual `/calculate` para `/v1/calculate` com redirect/alias para não quebrar consumidores existentes | RNF | 🟡 Média | 1–2 dias |
| 🔨 GAP-010 | **Deploy Artifacts** | Dockerfile multi-stage, K8s manifests (Deployment, Service, ConfigMap), resource limits/requests, probes (DT-10) | RNF | 🟡 Média | 3–5 dias |

### 2.3 O Que NÃO Será Implementado no Microserviço (Fora do Escopo)

Estas responsabilidades são de **outros sistemas** e estão documentadas aqui para clareza de fronteiras:

| # | Responsabilidade | Sistema Responsável | BR Relacionado |
|:---|:---|:---|:---|
| ❌ OUT-01 | Validação de código IBGE no cadastro do cliente | CRM / Sistema de Vendas | BR-03 |
| ❌ OUT-02 | Exibição visual da decomposição do IVA no checkout | Frontend e-commerce / Portal B2B | BR-04 |
| ❌ OUT-03 | Workflow de aprovação de fornecedores (due diligence) | ERP SAP / Sistema de Compras | BR-08 |
| ❌ OUT-04 | Conciliação bancária do split payment (CNAB) | Tesouraria / SAP FI | BR-09 |
| ❌ OUT-05 | Emissão de NF-e / NFS-e com layout IVA Dual | ERP SAP / Sistema de Faturamento | BR-07 |
| ❌ OUT-06 | Dashboard de KPIs para CFO | Ferramenta de BI / Grafana | MATRIZ-KPI |
| ❌ OUT-07 | Interface administrativa visual para time fiscal | Sistema de Backoffice (pode consumir GAP-001) | BR-02 |
| ❌ OUT-08 | Autenticação e autorização (JWT issuance) | Kong / Keycloak (API Gateway) | Segurança |
| ❌ OUT-09 | Coleta e armazenamento de logs centralizados | Plataforma de Observabilidade | Operação |
| ❌ OUT-10 | Agendamento de atualização de alíquotas | Sistema de Workflow / CronJobs | BR-02 |

### 2.4 Ordem Sugerida de Implementação (Roadmap Técnico)

```
FASE 0 — Fundação (Semanas 1–2)
├── GAP-008: Rate Limiting (proteção antes de expor novos endpoints)
├── GAP-009: API Versioning (/v1/)
└── GAP-010: Deploy Artifacts (Dockerfile + K8s)

FASE 1 — Onda 1: Comercial (Semanas 3–5)
├── GAP-004: Campo valor_liquido (pré-requisito para UI)
├── GAP-002: TaxToken (bloqueante para BR-06)
├── GAP-003: Endpoint /simulate (habilita simulação comercial)
└── GAP-001: Interface Admin (habilita time fiscal)

FASE 2 — Onda 2: Financeira (Semanas 6–10)
├── GAP-006: Schema Split Payment (bloqueante para BR-09)
├── GAP-005: Cálculo de Créditos (bloqueante para BR-08)
└── GAP-007: Qualificação de Fornecedores (complementa BR-08)
```

📄 **Detalhamento técnico de cada GAP:** [ARCHITECTURE.md — Seção 6.2](./ARCHITECTURE.md#62-gaps-identificados-funcionalidades-necessárias)
📄 **Mapeamento BR → GAP:** [ARCHITECTURE.md — Seção 2](./ARCHITECTURE.md#2-mapeamento-requisitos-de-negócio--capacidades-técnicas)

---

## 3. Artefatos do Projeto (Esta Pasta)

Documentos criados especificamente para este projeto de negócio:

| # | Documento | Descrição | Público-Alvo |
|:---|:---|:---|:---|
| 1 | [PRD.md](./PRD.md) | Product Requirements Document — resumo de alto nível do produto: visão geral, 9 BRs, hierarquia da documentação, cronograma, KPIs, riscos, stakeholders | Product Managers, Tech Leads, Stakeholders |
| 2 | [ARCHITECTURE.md](./ARCHITECTURE.md) | Decisões arquiteturais específicas do projeto: mapeamento BR→capacidades técnicas, 4 ADRs (registro canônico em [../../architecture/adrs/INDEX.md](../../architecture/adrs/INDEX.md)), gaps técnicos, diagrama C4, variáveis de ambiente, SLAs | Arquitetos, Tech Leads, DevOps |
| 3 | **[SPECS.md](./SPECS.md)** | **Este documento** — índice-mestre de todas as especificações, matriz de rastreabilidade completa, escopo de implementação, convenções do projeto | Todos os papéis |
| 4 | [TASKS.md](./TASKS.md) | Plano de implementação tático: 65 tarefas granulares em 3 fases, Definition of Done, grafo de dependências, runbook de deploy/operação, variáveis por ambiente, alertas | Tech Leads, Desenvolvedores, DevOps |
| 5 | [TEST_PLAN.md](./TEST_PLAN.md) | Estratégia de testes para os 10 GAPs: ~68 cenários em 4 níveis (L1–L4), mapeamento BDD→teste técnico, testes de performance (KPI O2 <100ms), suíte de regressão CI/CD | QA, Desenvolvedores, Tech Leads |

---

## 4. Documentos de Definição de Negócio (Fonte Primária)

Localização: `business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/`

### 4.1 Nível Estratégico

| # | Documento | Descrição | Status |
|:---|:---|:---|:---|
| 3.1.1 | [01-PROJECT-CHARTER.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/01-PROJECT-CHARTER.md) | Termo de Abertura do Programa v2.0 — Justificativa estratégica, 6 objetivos, escopo, macro-cronograma (2026–2033), stakeholders, riscos, critérios de sucesso | ✅ Aprovado |
| 3.1.2 | [PROJECT-CHARTER-001.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/PROJECT-CHARTER-001.md) | Versão 1.0 arquivada (histórico) | 📋 Arquivado |

### 4.2 Nível de Escopo — Requisitos de Negócio

| # | Documento | Descrição | Status |
|:---|:---|:---|:---|
| 3.2.1 | [02-BUSINESS-REQUIREMENTS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/02-BUSINESS-REQUIREMENTS.md) | 9 Requisitos de Negócio (BR-01 a BR-09) em 3 blocos: Fundação, Onda 1 (Comercial), Onda 2 (Financeira). Inclui matriz de rastreabilidade e critérios de homologação (UAT) | ✅ Aprovado |

### 4.3 Nível Macro-Escopo — Épicos

| # | Documento | Onda | Épicos | Status |
|:---|:---|:---|:---|:---|
| 3.3.1 | [03-EPICS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/03-EPICS.md) | Onda 1 | Épico 01 (CRM), Épico 02 (Integração), Épico 03 (Precificação) | ✅ Pronto |
| 3.3.2 | [03-EPICS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/03-EPICS.md) | Onda 2 | Épico 01 (Faturamento), Épico 02 (Split), Épico 03 (Créditos) | ✅ Pronto |

### 4.4 Nível Funcional — Features

| # | Documento | Onda | Features | Status |
|:---|:---|:---|:---|:---|
| 3.4.1 | [04-FEATURES.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/04-FEATURES.md) | Onda 1 | 7 Features: Validação Cadastral, Trava Comercial, Simulador Omnicanal, Resiliência, Decomposição IVA, Painel B2B, Token Fiscal | ✅ Pronto |
| 3.4.2 | [04-FEATURES.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/04-FEATURES.md) | Onda 2 | 9 Features: Validação Pré-Emissão, Conversão ISS→IBS, Benefícios Fiscais, Split, Ajuste Split, Painel Auditoria, Bloqueio Créditos, Segregação Contábil, Reserva Incentivos | ✅ Pronto |

### 4.5 Nível de Execução — User Stories (18 arquivos)

**Onda 1 — Canais Comerciais (7 US):**

| # | Documento | Feature |
|:---|:---|:---|
| 3.5.1 | [USER-STORYS-01-01-1-VALIDACAO-CADASTRAL-GEOGRAFICA-TEMPO-REAL.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-01-01-1-VALIDACAO-CADASTRAL-GEOGRAFICA-TEMPO-REAL.md) | 01.01 — Validação Cadastral |
| 3.5.2 | [USER-STORYS-01-01-2-TRAVA-COMERCIAL-NO-CRM-POR-FALTA-DE-HIGIENIZACAO-CADASTRAL.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-01-01-2-TRAVA-COMERCIAL-NO-CRM-POR-FALTA-DE-HIGIENIZACAO-CADASTRAL.md) | 01.02 — Trava Comercial |
| 3.5.3 | [USER-STORYS-01-02-1-SIMULADOR-UNIFICADO-OMNICANAL.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-01-02-1-SIMULADOR-UNIFICADO-OMNICANAL.md) | 02.01 — Simulador Omnicanal |
| 3.5.4 | [USER-STORYS-01-02-2-RESILIENCIA-VENDAS-CONTINGENCIA-LOCAL.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-01-02-2-RESILIENCIA-VENDAS-CONTINGENCIA-LOCAL.md) | 02.02 — Resiliência |
| 3.5.5 | [USER-STORYS-01-03-1-INTERFACE-VISUAL-CHECKOUT-DECOMPOSICAO-IVA.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-01-03-1-INTERFACE-VISUAL-CHECKOUT-DECOMPOSICAO-IVA.md) | 03.01 — Decomposição IVA |
| 3.5.6 | [USER-STORYS-01-03-2-PAINEL-ATRATIVIDADE-B2B-CALCULADORA-CREDITO.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-01-03-2-PAINEL-ATRATIVIDADE-B2B-CALCULADORA-CREDITO.md) | 03.02 — Painel B2B |
| 3.5.7 | [USER-STORYS-01-03-3-CHAVE-DE-GARANTIA-TOKEN-DE-VALIDADE-FISCAL.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-01-03-3-CHAVE-DE-GARANTIA-TOKEN-DE-VALIDADE-FISCAL.md) | 03.03 — Token Fiscal |

**Onda 2 — Finanças, Faturamento e ERP (9 US):**

| # | Documento | Feature |
|:---|:---|:---|
| 3.5.8 | [USER-STORYS-02-01-1-VALIDACAO-FATURAMENTO-PRE-EMISSAO-TRAVA-CONTABIL.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-02-01-1-VALIDACAO-FATURAMENTO-PRE-EMISSAO-TRAVA-CONTABIL.md) | 01.01 — Validação Pré-Emissão |
| 3.5.9 | [USER-STORYS-02-01-2-MOTOR-CONVERSAO-ISS-IBS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-02-01-2-MOTOR-CONVERSAO-ISS-IBS.md) | 01.02 — Conversão ISS→IBS |
| 3.5.10 | [USER-STORYS-02-01-3-AUTOMACAO-BENEFICIOS-REGIMES-ESPECIAIS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-02-01-3-AUTOMACAO-BENEFICIOS-REGIMES-ESPECIAIS.md) | 01.03 — Benefícios |
| 3.5.11 | [USER-STORYS-02-02-1-LIQUIDACAO-CONCILIACAO-SEGREGADA-SPLIT.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-02-02-1-LIQUIDACAO-CONCILIACAO-SEGREGADA-SPLIT.md) | 02.01 — Conciliação Split |
| 3.5.12 | [USER-STORYS-02-02-2-AJUSTE-SPLIT-OPERACOES-INCENTIVADAS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-02-02-2-AJUSTE-SPLIT-OPERACOES-INCENTIVADAS.md) | 02.02 — Ajuste Split |
| 3.5.13 | [USER-STORYS-02-02-3-PAINEL-AUDITORIA-RECONCILIACAO-SPLIT.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-02-02-3-PAINEL-AUDITORIA-RECONCILIACAO-SPLIT.md) | 02.03 — Auditoria Split |
| 3.5.14 | [USER-STORYS-02-03-1-AUDITORIA-FISCAL-ENTRADA-BLOQUEIO-CREDITOS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-02-03-1-AUDITORIA-FISCAL-ENTRADA-BLOQUEIO-CREDITOS.md) | 03.01 — Bloqueio Créditos |
| 3.5.15 | [USER-STORYS-02-03-2-SEGREGACAO-CONTABIL-CUSTOS-ESTOQUE-ATIVOS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-02-03-2-SEGREGACAO-CONTABIL-CUSTOS-ESTOQUE-ATIVOS.md) | 03.02 — Segregação Contábil |
| 3.5.16 | [USER-STORYS-02-03-3-ESCRITURACAO-RESERVA-INCENTIVOS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/05-USER-STORYS-02-03-3-ESCRITURACAO-RESERVA-INCENTIVOS.md) | 03.03 — Reserva Incentivos |

- 💡 As User Stories acima são as 16 que foram detalhadas com critérios de aceite BDD. O total de 41 US do projeto inclui variações e cenários adicionais documentados nos arquivos de Features. Consulte o [README.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/README.md) do projeto para a lista completa.

### 4.6 Governança e Métricas

| # | Documento | Descrição | Status |
|:---|:---|:---|:---|
| 3.6.1 | [MATRIZ-KPI.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/MATRIZ-KPI.md) | 8 KPIs em 3 dimensões (Financeira F1–F3, Compliance C1–C2, Operacional O1–O2) com fórmulas, metas e template de dashboard para CFO | ✅ Pronto |

### 4.7 Base de Conhecimento Técnico-Tributário (docs-suporte/)

| # | Documento | Conteúdo |
|:---|:---|:---|
| 3.7.1 | [README-ESCOPO.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/docs-suporte/README-ESCOPO.md) | Catálogo de impostos PJ, estrutura de microserviços, matriz de convivência de regimes |
| 3.7.2 | [README-BRAINSTORM.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/docs-suporte/README-BRAINSTORM.md) | Modelagem técnica Go para motores de cálculo fiscal |
| 3.7.3 | [README-ICMS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/docs-suporte/README-ICMS.md) | Regras detalhadas de ICMS, alíquotas interestaduais, DIFAL |
| 3.7.4 | [README-ICMS-EXTENSAO-REGRAS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/docs-suporte/README-ICMS-EXTENSAO-REGRAS.md) | Extensões de regras ICMS, ST, FCP |
| 3.7.5 | [README-PIS-COFINS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/docs-suporte/README-PIS-COFINS.md) | Regimes cumulativo e não-cumulativo |
| 3.7.6 | [README-IPI.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/docs-suporte/README-IPI.md) | TIPI, NCM, fato gerador, alíquotas |
| 3.7.7 | [README-CONSTANTS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/docs-suporte/README-CONSTANTS.md) | Constantes fiscais, tabelas de alíquotas, códigos IBGE |
| 3.7.8 | [README-SIMPLES-NACIONAL.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/docs-suporte/README-SIMPLES-NACIONAL.md) | Anexos I-V, RBT12, regras de transição |
| 3.7.9 | [README-TABELA-CST-CSON.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/docs-suporte/README-TABELA-CST-CSON.md) | Tabelas CST, CSOSN, CFOP |
| 3.7.10 | docs-suporte/*.pdf (11 arquivos) | Material do curso RTC CFC-RFB, Nota Técnica NT 2025.002, cartilha da reforma |

---

## 5. Especificações Técnicas do Microserviço (`.specs/`)

Localização: `.specs/` (raiz do projeto `ms-billing-engine-tax-rates`)

### 5.1 Índice Geral (.specs/INDEX.md)

| # | Documento | Descrição |
|:---|:---|:---|
| 4.1.1 | [INDEX.md](../INDEX.md) | Mapa geral das especificações do microserviço com links para todas as subpastas |

### 5.2 Arquitetura

| # | Documento | Descrição | Confiança |
|:---|:---|:---|:---|
| 4.2.1 | [architecture/architecture.md](../architecture/architecture.md) | Visão arquitetural completa: pipeline SOP-013 7-fases, injeção de dependência manual, middleware pipeline, strategy/adapter patterns, injeção inter-fase, health checks, métricas, autenticação, dívidas técnicas | ✅ 100% |
| 4.2.2 | [architecture/c4-context.md](../architecture/c4-context.md) | Diagrama C4 Nível 1 — atores, sistemas, fluxo de requisição, protocolos de integração | ✅ 100% |
| 4.2.3 | [architecture/erd.md](../architecture/erd.md) | Modelo Entidade-Relacionamento completo (10 tabelas, triggers PL/pgSQL, índices, padrões de vigência temporal) | ✅ 100% |
| 4.2.4 | [architecture/integrations.md](../architecture/integrations.md) | Dependências (go.mod), variáveis de ambiente, bibliotecas locais, integração com API do Comitê Gestor IBS, Circuit Breaker | ✅ 100% |

### 5.3 Domínio Fiscal

| # | Documento | Descrição | Confiança |
|:---|:---|:---|:---|
| 4.3.1 | [domain/domain.md](../domain/domain.md) | Glossário completo de domínio fiscal: 13 regras de negócio principais, pipeline SOP-013 fase a fase, Phase Resolution System, ICMS Desonerado, tributos implementados, glossário com 21 termos | ✅ 100% |

### 5.4 Produto e Requisitos

| # | Documento | Descrição | Confiança |
|:---|:---|:---|:---|
| 4.4.1 | [product/requirements.md](../product/requirements.md) | 11 Requisitos Funcionais (RF-01 a RF-11) + 15 Requisitos Não-Funcionais (RNF-01 a RNF-15) com evidência de código e priorização MoSCoW | ✅ 100% |
| 4.4.2 | [product/feature-roadmap.md](../product/feature-roadmap.md) | Roadmap completo: 31 features implementadas, 6 features planejadas, 17 dívidas técnicas com prioridade e localização no código | ✅ 100% |

### 5.5 Engenharia e Padrões

| # | Documento | Descrição | Confiança |
|:---|:---|:---|:---|
| 4.5.1 | [engineering/code-analysis.md](../engineering/code-analysis.md) | Análise técnica consolidada: entry points, módulos, fluxo de inicialização, estrutura de componentes | ✅ 100% |
| 4.5.2 | [engineering/api-guidelines.md](../engineering/api-guidelines.md) | Padrões de API: tratamento de erros, rastreabilidade (W3C Trace Context), validação de entrada, autenticação JWT, sanitização de erros, métricas Prometheus | ✅ 100% |

### 5.6 API e Contratos

| # | Documento | Descrição | Confiança |
|:---|:---|:---|:---|
| 4.6.1 | [api/tax-rates-api.yaml](../api/tax-rates-api.yaml) | Contrato OpenAPI 3.0 — schemas, endpoints, erros, auth JWT, enum de 11 tributos | ✅ 100% |

### 5.7 Features e Gap Analysis

| # | Documento | Descrição |
|:---|:---|:---|
| 4.7.1 | [features/FEATURE-2026-06-21.md](../features/FEATURE-2026-06-21.md) | Registro das features implementadas em 2026-06-21 |
| 4.7.2 | [features/FEATURE-2026-06-21-GAP-ANALISYS.md](../features/FEATURE-2026-06-21-GAP-ANALISYS.md) | Gap Analysis — conformidade spec→código: 54/55 requisitos ✅, 3 gaps de documentação, 1 divergência |

### 5.8 Governança e Qualidade

| # | Documento | Descrição |
|:---|:---|:---|
| 4.8.1 | [governance/inventory.md](../governance/inventory.md) | Inventário do projeto: estrutura física do código, arquivos fonte, cobertura de testes |
| 4.8.2 | [governance/confidence-report.md](../governance/confidence-report.md) | Relatório de confiança das especificações: score VERDE (99%), cobertura de artefatos |

### 5.9 Questões e Lacunas

| # | Documento | Descrição |
|:---|:---|:---|
| 4.9.1 | [questions/questions_01.md](../questions/questions_01.md) | Histórico de lacunas resolvidas e perguntas respondidas |

### 5.10 Histórico de Implementação (skill-output/)

15 registros documentando cada sessão de implementação:

| # | Arquivo | Feature Implementada |
|:---|:---|:---|
| 4.10.1 | `skill-output/2026-06-20-192700_icms-proprio-st.md` | ICMS Próprio e ST |
| 4.10.2 | `skill-output/2026-06-20-192730_pis-cofins-db-rates.md` | PIS/COFINS via banco |
| 4.10.3 | `skill-output/2026-06-20-204235-pis-cofins-cst-completion.md` | Cobertura total de CSTs PIS/COFINS |
| 4.10.4 | `skill-output/2026-06-20-205853_exclusao-icms-base-pis-cofins.md` | Exclusão ICMS da base PIS/COFINS |
| 4.10.5 | `skill-output/2026-06-20-214728_testes-icms-ipi-engine.md` | Testes ICMS, IPI e Engine |
| 4.10.6 | `skill-output/2026-06-21-003638_middleware-requestid-traceid.md` | Middleware W3C Trace Context |
| 4.10.7 | `skill-output/2026-06-21-005150_health-checks_goroutine-errors.md` | Health Checks + Goroutine Errors |
| 4.10.8 | `skill-output/2026-06-21-013456_auth-metrics-port.md` | Auth JWT + Métricas + Porta Configurável |
| 4.10.9 | `skill-output/2026-06-21-151743_reforma-tributaria-cbs-ibs-is.md` | Reforma Tributária CBS/IBS/IS |
| 4.10.10 | `skill-output/2026-06-21-154731_dt01-uuid-idtransaction.md` | UUID IDTransaction (DT-01) |
| 4.10.11 | `skill-output/2026-06-21-184933_domain-restructure.md` | Reorganização DDD (domain/) |
| 4.10.12 | `skill-output/2026-06-21-230924_iss-fust-funttel.md` | ISS + FUST + FUNTTEL |
| 4.10.13 | `skill-output/2026-06-21-235900_icms-desonerado-phase-resolver.md` | ICMS Desonerado + Phase Resolver |
| 4.10.14 | `skill-output/2026-06-22-004300_is-prefilter-ibs-circuitbreaker.md` | IS Pré-Filtro + IBS Circuit Breaker |
| 4.10.15 | `skill-output/2026-06-22-011000_c001-pipeline-sop013.md` | Pipeline SOP-013 7-fases (C-001) |

---

## 6. Matriz de Rastreabilidade Completa

### 6.1 Requisitos de Negócio → Requisitos Funcionais → Features Técnicas → Código

| BR (Negócio) | RF (Técnico) | Feature ID | Arquivo Principal |
|:---|:---|:---|:---|
| BR-01 — Centralização | RF-01 (Endpoint /calculate) | C-001 | `cmd/api/main.go:121-168`, `internal/calculator/engine.go` |
| BR-02 — Autonomia Fiscal | — (Tabela `iva_dual_rules` + triggers) | C-002 | `data/init.sql`, `repository/postgres_repository.go:359-401` |
| BR-03 — Qualificação Geográfica | — (Campo `municipioIBGE` no input) | C-001 | `internal/reforma/ibs_calculator.go:1-65` |
| BR-04 — Transparência "Por Fora" | RF-08 (CBS/IBS/IS) | C-001, F-006 | `internal/reforma/cbs_calculator.go`, `is_filter.go` |
| BR-05 — Proteção de Margem | RF-08 (CBS/IBS/IS) + Phase Resolver | F-005 | `internal/phase/phase.go`, `internal/phase/tax_selector.go` |
| BR-06 — Token de Validade | ✅ GAP-002 | `internal/token/` | Implementado |
| BR-07 — Unicidade Pedido↔NF | RF-04 (Motor), RNF-08 (Trace Context) | C-001 | `internal/calculator/engine.go`, `internal/middleware/requestid.go` |
| BR-08 — Créditos Lucro Real | ✅ GAP-005 | `internal/credit/` | Implementado |
| BR-09 — Split Payment | ⚠️ GAP-006 | `internal/calculator/engine.go` | Parcial |

### 6.2 Épicos de Negócio → Capacidades Técnicas

| Épico (Negócio) | Capacidade Técnica | Feature ID | Status |
|:---|:---|:---|:---|
| Épico 01.01 (CRM/Saneamento) | Consumo do campo `municipioIBGE` no input | C-001 | ✅ |
| Épico 01.02 (Integração Comercial) | `POST /calculate` como endpoint único | RF-01 | ✅ |
| Épico 01.03 (Precificação Dinâmica) | Motor 7-fases + Phase Resolver | C-001, F-005 | ✅ |
| Épico 02.01 (Faturamento SAP) | Unicidade via UUID + Trace Context | RF-04, RNF-08 | ✅ |
| Épico 02.02 (Split Payment) | ⚠️ GAP-006 — Schema de split na resposta | `internal/calculator/engine.go` | Parcial |
| Épico 02.03 (Créditos) | ✅ GAP-005 — `/credit/calculate` | `internal/credit/` | ✅ |

### 6.3 KPIs de Negócio → Métricas Técnicas

| KPI | Métrica Prometheus / Técnica | SLA |
|:---|:---|:---|
| O2 — Latência API Cálculo | `http_request_duration_seconds` (p95) | < 100ms |
| C1 — Rejeição de Notas | `errors_total{type="validation"}` | < 0,1% |
| F1 — Aproveitamento Créditos | 🔴 A definir com GAP-005 | ≥ 98% |
| F3 — Divergência Split | 🔴 A definir com GAP-006 | R$ 0,00 |

---

## 7. Estatísticas Consolidadas do Projeto

| Métrica | Valor |
|:---|:---|
| **Documentos de Negócio** | 25 (1 Charter + 1 Requirements + 2 Epics + 2 Features + 16 US + 1 KPIs + 1 README + 1 histórico) |
| **Documentos Técnicos (.specs/)** | 22 (4 architecture + 1 domain + 2 product + 2 engineering + 1 api + 2 features + 2 governance + 1 questions + 1 INDEX + 15 skill-output) |
| **Artefatos do Projeto** | 5 (PRD + ARCHITECTURE + SPECS + TASKS + TEST_PLAN) |
| **Requisitos de Negócio (BR)** | 9 |
| **Requisitos Funcionais (RF)** | 11 |
| **Requisitos Não-Funcionais (RNF)** | 15 |
| **Épicos** | 6 |
| **Features de Negócio** | 16 |
| **Features Técnicas** | 9 (F-001 a F-007 + C-001 + C-002) |
| **User Stories** | 41 |
| **Regras de Negócio (RN)** | ~90 |
| **Critérios de Aceite BDD** | ~55 cenários |
| **KPIs** | 8 |
| **Dívidas Técnicas** | 17 (7 resolvidas, 10 ativas) |
| **Gaps do Projeto** | 10 (GAP-001 a GAP-010) |
| **Testes Automatizados** | 211+ passando (25 arquivos de teste) |
| **Tabelas no Schema** | 10 |
| **Tributos Implementados** | 10 |
| **Confiança da Documentação** | VERDE (99%) |

---

## 8. Convenções e Padrões do Projeto

### 8.1 Nomenclatura de Arquivos

| Tipo | Padrão | Exemplo |
|:---|:---|:---|
| Artefatos do Projeto | `NOME.md` (MAIÚSCULO) | `PRD.md`, `ARCHITECTURE.md`, `SPECS.md` |
| Documentos de Negócio | `TIPO-DESCRIÇÃO.md` | `01-PROJECT-CHARTER.md`, `03-EPICS.md` |
| User Stories | `05-USER-STORYS-{onda}-{epico}-{feature}-{descricao}.md` | `05-USER-STORYS-01-01-1-VALIDACAO-CADASTRAL.md` |
| Especificações Técnicas | `kebab-case.md` ou `kebab-case.yaml` | `architecture.md`, `tax-rates-api.yaml` |
| Registros de Implementação | `YYYY-MM-DD-HHMMSS_descricao.md` | `2026-06-21-151743_reforma-tributaria.md` |

### 8.2 Rastreabilidade

- Todo documento de negócio referencia seu antecessor hierárquico via marcador `[INDEX]`
- User Stories seguem formato: Descrição (Como/Quero/Para) + Regras de Negócio (RN01...) + Critérios de Aceite BDD (Cenário: Dado/Quando/Então)
- Documentos técnicos em `.specs/` referenciam o código-fonte com `file_path:line_number`
- O PRD referencia documentos de negócio; o ARCHITECTURE.md referencia documentos técnicos; o SPECS.md referencia ambos

### 8.3 Status de Documentos

| Status | Significado |
|:---|:---|
| ✅ Aprovado / Pronto / Completo | Documento validado e liberado para uso |
| 🟡 Em Definição / Parcial | Conteúdo existe mas requer refinamento |
| 🔴 Não Implementado | Gap identificado, ação necessária |
| 📋 Arquivado | Versão histórica, não usar como referência ativa |

### 8.4 Níveis de Confiança (.specs/)

| Score | Significado |
|:---|:---|
| 🟢 VERDE (95–100%) | Especificação verificada contra código-fonte |
| 🟡 AMARELO (70–94%) | Especificação parcialmente verificada |
| 🔴 VERMELHO (<70%) | Especificação não verificada ou desatualizada |

---

## 9. Guia de Navegação Rápida

### "Preciso entender o que o projeto faz..."
→ Comece pelo [PRD.md](./PRD.md) (Seções 1–3)

### "Preciso saber COMO implementar..."
→ Vá para [ARCHITECTURE.md](./ARCHITECTURE.md) (Seção 2: Mapeamento BR→Técnico, Seção 6: Gaps)

### "Preciso das tarefas detalhadas e do runbook de deploy..."
→ Vá para [TASKS.md](./TASKS.md) — 65 tarefas (✅ todas concluídas), DoD, dependências, runbook, alertas

### "Preciso saber o que testar e como..."
→ Vá para [TEST_PLAN.md](./TEST_PLAN.md) — ~68 cenários, níveis L1–L4, performance, BDD mapeado

### "Preciso dos detalhes de negócio (regras, BDD)..."
→ Consulte os [Documentos de Negócio](#4-documentos-de-definição-de-negócio-fonte-primária) na Seção 4

### "Preciso entender o código/arquitetura do microserviço..."
→ Consulte as [Especificações Técnicas](#5-especificações-técnicas-do-microsserviço-specs) na Seção 5

### "Preciso saber o que já foi implementado e o que falta..."
→ Veja a [Matriz de Rastreabilidade](#6-matriz-de-rastreabilidade-completa) na Seção 6 e os [Gaps](./ARCHITECTURE.md#62-gaps-identificados-funcionalidades-necessárias) no ARCHITECTURE.md

### "Preciso do histórico de implementações..."
→ Consulte os [15 registros em skill-output/](#510-histórico-de-implementação-skill-output) na Seção 5.10

### "Preciso saber quantos testes existem e onde..."
→ [governance/inventory.md](../governance/inventory.md) + [governance/confidence-report.md](../governance/confidence-report.md)

### "Preciso ver o contrato da API..."
→ [api/tax-rates-api.yaml](../api/tax-rates-api.yaml)

### "Preciso entender o modelo de dados..."
→ [architecture/erd.md](../architecture/erd.md)

---

- 📋 **Manutenção deste documento:** Atualize o SPECS.md sempre que um novo artefato for criado, um gap for resolvido, ou uma nova feature for implementada. Este é o documento de referência para onboarding de novos membros do time e para auditoria de conformidade do projeto.
