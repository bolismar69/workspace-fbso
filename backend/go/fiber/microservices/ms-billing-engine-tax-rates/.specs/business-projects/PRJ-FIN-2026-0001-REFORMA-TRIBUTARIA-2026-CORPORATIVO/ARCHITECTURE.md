# Architecture Document — Adequação Corporativa à Reforma Tributária Nacional

**Código do Projeto:** PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO
**Versão:** 1.0
**Data:** 23 de Junho de 2026
**Microserviço-base:** `ms-billing-engine-tax-rates`

> ⚠️ **Aviso de Leitura:** Este documento define a arquitetura específica para atender ao projeto de Reforma Tributária. Ele **complementa**, e não substitui, os documentos de arquitetura do microserviço em `.specs/architecture/`, que são a fonte da verdade sobre a implementação. Consulte a Seção 7 para o mapa completo de referências cruzadas.

---

## 1. Propósito e Escopo Arquitetural

Este documento define as decisões arquiteturais necessárias para que o microserviço `ms-billing-engine-tax-rates` atenda integralmente aos 9 Requisitos de Negócio (BR-01 a BR-09) do programa de Adequação Corporativa à Reforma Tributária Nacional, cobrindo:

- **Onda 1 (Comercial):** Cálculo de tributos em tempo real nos canais de venda (simulação, checkout, CRM)
- **Onda 2 (Financeira):** Faturamento, split payment, conciliação contábil e apropriação de créditos no Lucro Real
- **Dimensão Estratégica:** Gestão do período híbrido 2029–2032 e transição completa até 2033

A arquitetura-base é o motor de cálculo multi-fase SOP-013 documentado em `.specs/architecture/architecture.md`. Este documento **estende** essa arquitetura com as especificidades do projeto de negócio.

📄 **Fonte da verdade técnica:** [.specs/architecture/architecture.md](../.specs/architecture/architecture.md)
📄 **Fonte dos requisitos de negócio:** [REQUIREMENTS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/REQUIREMENTS.md)

---

## 2. Mapeamento Requisitos de Negócio → Capacidades Técnicas

A tabela abaixo mostra como cada Business Requirement é atendido pela arquitetura atual e quais lacunas precisam ser endereçadas.

### 2.1 Bloco 1 — Fundação (BR-01 a BR-03)

| BR | Requisito de Negócio | Cobertura Atual | Gap / Ação Necessária |
|:---|:---|:---|:---|
| **BR-01** | Centralização da Inteligência de Regras | ✅ **Coberto** — O motor `BillingEnginePhased` é a fonte única de cálculo. O endpoint `POST /calculate` centraliza todas as regras. Tabelas `iva_dual_rules`, `federal_tax_rules`, `icms_rules` e `ipi_regras` no PostgreSQL garantem regra unificada. | Nenhum — arquitetura já suporta. Monitorar que novos canais (e-commerce, CRM, portal B2B) consumam o mesmo endpoint. |
| **BR-02** | Autonomia do Time Fiscal (No-Code/Dynamic Update) | ✅ **Coberto** — A tabela `iva_dual_rules` com `inicio_validade`/`final_validade` + trigger PL/pgSQL permite ao time fiscal atualizar alíquotas via SQL/interface sem deploy. Cache Redis com TTL 24h garante propagação automática. Tabela de auditoria `iva_dual_rules_log` registra todas as alterações. | 🟡 **Parcial** — Falta interface administrativa (UI/API) para o time fiscal. Atualmente a atualização é direta no banco. **Ação:** Criar endpoint `PUT /admin/tax-rates` com autenticação e log de auditoria. |
| **BR-03** | Qualificação Geográfica de Cadastro | 🔴 **Fora do escopo do microserviço** — A validação do código IBGE no cadastro é responsabilidade do CRM/sistemas de venda. O motor de cálculo **consome** o código IBGE como input (`municipioIBGE` em `ItemDocumentoFiscalEntrada`) e o utiliza para lookup de alíquotas IBS municipais. | **Ação:** Documentar contrato de API: o campo `municipioIBGE` é obrigatório para operações com IBS. Sistemas consumidores (CRM, e-commerce) devem validar antes de chamar `/calculate`. |

### 2.2 Bloco 2 — Onda 1: Comercial (BR-04 a BR-06)

| BR | Requisito de Negócio | Cobertura Atual | Gap / Ação Necessária |
|:---|:---|:---|:---|
| **BR-04** | Transparência e Cálculo "Por Fora" | ✅ **Coberto** — O pipeline SOP-013 posiciona CBS na Fase 2 (sequencial, "por fora") antes do ICMS. O `DocumentoFiscalSaida` retorna tributos segregados por item com `TributosItemDocumentoFiscalSaida` contendo CBS, IBS, IS, ICMS, PIS, COFINS individualmente. O campo `detalhes_base` permite decomposição visual. | 🟡 **Parcial** — O contrato de resposta atual não inclui campo explícito `preco_liquido` (preço base sem tributos). **Ação:** Adicionar `valor_liquido` ao schema de resposta para facilitar a UI de checkout "por fora". |
| **BR-05** | Proteção de Margem e Simulação Comercial | 🟡 **Parcialmente coberto** — O motor calcula tributos por destino, mas não expõe endpoint de simulação com projeção de margem. A Fase 4 paralela (IBS+ISS+PISCOFINS) já computa IBS por município. | **Ação:** Criar endpoint `POST /simulate` que retorna o impacto dos tributos por UF/município sem persistir transação, permitindo ao time comercial simular cenários antes do fechamento. |
| **BR-06** | Garantia de Preço Ofertado (Token de Validade) | 🔴 **Não implementado** — O motor não possui mecanismo de congelamento de alíquota por janela temporal. | **Ação:** Implementar `TaxToken` — estrutura que associa uma tupla `(ncm, ufDestino, municipioIBGE, aliquotasSnapshot, expiresAt)` com TTL configurável. Endpoint `POST /token/generate` e `POST /calculate` com parâmetro opcional `token_id` para usar alíquotas congeladas. |

### 2.3 Bloco 3 — Onda 2: Financeira (BR-07 a BR-09)

| BR | Requisito de Negócio | Cobertura Atual | Gap / Ação Necessária |
|:---|:---|:---|:---|
| **BR-07** | Unicidade Matemática Pedido ↔ Nota Fiscal | ✅ **Coberto** — A arquitetura garante consistência porque o mesmo motor (`BillingEnginePhased`) é usado tanto para simulação quanto para faturamento. O `IDTransaction` (UUID v4) gerado em cada cálculo fornece rastreabilidade. O Phase Resolution System garante que a mesma `DataOperacao` produza os mesmos resultados. | Nenhum — arquitetura já garante idempotência por design. |
| **BR-08** | Rastreabilidade de Créditos no Lucro Real | 🔴 **Não implementado** — O motor calcula tributos na saída (vendas), mas não possui lógica de crédito na entrada (compras/fornecedores). O campo `permite_credito_amplo` em `iva_dual_rules` existe mas não é consumido ativamente. | **Ação:** Criar endpoint `POST /credit/calculate` que, dado um documento fiscal de entrada (compra), calcula os créditos de CBS/IBS apropriáveis conforme regras do Lucro Real. Implementar validação de fornecedor (due diligence fiscal). |
| **BR-09** | Viabilização do Split Payment | 🔴 **Não implementado** — O motor não possui lógica de split payment. A arquitetura atual retorna o total de tributos, mas não discrimina a partição financeira esperada no momento da liquidação bancária. | **Ação:** Adicionar ao schema de resposta os campos `valor_receita_liquida`, `valor_cbs_reter`, `valor_ibs_reter`, `valor_is_reter` para que o sistema de tesouraria possa instruir a rede bancária sobre o split. |

📄 **Fonte dos requisitos:** [REQUIREMENTS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/REQUIREMENTS.md)
📄 **Fonte da cobertura técnica:** [.specs/product/requirements.md](../product/requirements.md)

---

## 3. Decisões Arquiteturais Específicas do Projeto

### 3.1 ADR-001: Motor de Cálculo Unificado como Fonte da Verdade

**Decisão:** O motor `BillingEnginePhased` (SOP-013, 7 fases) é a **única** fonte de cálculo de tributos para todo o ecossistema corporativo. Nenhum canal (CRM, e-commerce, ERP, portal B2B) pode implementar cálculo próprio de tributos.

**Racional:** Atende diretamente BR-01 (Centralização da Inteligência) e BR-07 (Unicidade Pedido ↔ Nota Fiscal). Elimina o risco de divergências de centavos entre canais.

**Mecanismo de enforcement:** O endpoint `POST /calculate` é o contrato único. O W3C Trace Context (`traceparent`/`traceresponse`) garante rastreabilidade ponta a ponta entre o canal de origem e o cálculo.

📄 **Fonte técnica:** [.specs/architecture/architecture.md — Seção 3](../.specs/architecture/architecture.md)

### 3.2 ADR-002: Phase Resolution System como Mecanismo de Transição Temporal

**Decisão:** O `PhaseResolver` (F-005) é o mecanismo canônico para gerenciar a transição entre regimes tributários (2026–2033). Toda requisição ao motor resolve a fase com base em `DataOperacao`:

```
DataOperacao ∈ 2026          → SHADOW_RUN         (CBS/IBS sombra, não compõe total)
DataOperacao ∈ {2027, 2028}  → CBS_PLENA          (CBS ativo, IBS sombra, PIS/COFINS extinto)
DataOperacao ∈ {2029..2032}  → TRANSICAO_SUBNACIONAL (redução progressiva ICMS/ISS)
DataOperacao ≥ 2033          → IVA_DUAL           (apenas CBS+IBS, legados extintos)
```

**Racional:** Atende aos requisitos de convivência de regimes durante o período híbrido (2029–2032) e à descontinuação progressiva de obrigações acessórias, conforme Project Charter Seções 3 e 4.

**Fases ativas por período (Matriz DT-001):**

| Período | CBS | IBS | PIS/COFINS | ICMS | ISS | IPI | IS |
|:---|:---|:---|:---|:---|:---|:---|:---|
| 2026 (Shadow) | 🟡 Sombra | 🟡 Sombra | ✅ Ativo | ✅ Ativo | ✅ Ativo | ✅ Ativo | 🟡 Sombra |
| 2027–28 (CBS Plena) | ✅ Ativo | 🟡 Sombra | ❌ Extinto | ✅ Ativo | ✅ Ativo | ✅ Ativo | ✅ Ativo |
| 2029–32 (Transição) | ✅ Ativo | ✅ Ativo | ❌ Extinto | 🔻 Reduzindo | 🔻 Reduzindo | ✅ Ativo | ✅ Ativo |
| 2033+ (IVA Dual) | ✅ Ativo | ✅ Ativo | ❌ Extinto | ❌ Extinto | ❌ Extinto | ✅ Ativo | ✅ Ativo |

📄 **Fonte técnica:** [.specs/domain/domain.md — Seção 13](../domain/domain.md)
📄 **Fonte negócio:** [PROJECT-CHARTER.md — Seção 4.2](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/PROJECT-CHARTER.md)

### 3.3 ADR-003: Cache Redis como Camada de Resiliência Regulatória

**Decisão:** O cache Redis (TTL 24h) com fallback para PostgreSQL é a camada primária de resiliência para consulta de alíquotas. O Circuit Breaker (`IBSRateFetcher` — F-007) adiciona proteção contra indisponibilidade da futura API do Comitê Gestor do IBS.

**Racional:** Atende ao risco "Indisponibilidade de Alíquotas do IBS por Município" identificado no Project Charter (Seção 7). A réplica local da matriz de alíquotas IBS no PostgreSQL (`iva_dual_rules`) garante operação contínua mesmo sem a API externa.

**Estratégia de fallback:**
```
1. Redis (TTL 24h, chave: tax:iva:<ncm>:<uf>:<municipio>)
2. PostgreSQL (tabela iva_dual_rules, com índice por NCM+UF+IBGE)  
3. API Comitê Gestor IBS (via CircuitBreakerIBSClient — planejada, bloqueada pelo Gap G2)
4. Alíquotas default hardcoded (último recurso, com slog.Warn)
```

📄 **Fonte técnica:** [.specs/architecture/integrations.md — IBS API](../.specs/architecture/integrations.md)

### 3.4 ADR-004: Contrato de API Versionado para Evolução Não-Quebrante

**Decisão:** O endpoint de cálculo será versionado (`/v1/calculate`) para permitir evolução do schema durante a transição tributária sem quebrar consumidores existentes. A versão `v2` está planejada para incorporar split payment e créditos do Lucro Real (BR-08, BR-09).

**Racional:** O calendário de transição (2026–2033) implica mudanças significativas nos schemas de entrada e saída. Versionamento permite que canais migrem gradualmente.

📄 **Fonte técnica:** [.specs/product/feature-roadmap.md — API Versioning](../product/feature-roadmap.md)

---

## 4. Arquitetura de Dados Específica do Projeto

### 4.1 Tabelas Relevantes para a Reforma Tributária

O schema `billing_tax_rates` contém 10 tabelas. As 3 tabelas diretamente vinculadas à Reforma Tributária são:

| Tabela | Propósito no Projeto | BR Atendido |
|:---|:---|:---|
| `iva_dual_rules` | Matriz de alíquotas CBS, IBS (estadual+municipal), IS por NCM + UF + município IBGE. Tabela-mestre da reforma. | BR-01, BR-03, BR-05 |
| `iva_dual_rules_log` | Auditoria de alterações nas alíquotas (triggers PL/pgSQL). Essencial para conformidade. | BR-02 |
| `ncm_seletivo` | Catálogo de NCMs sujeitos ao Imposto Seletivo (IS). Pré-filtro da Fase 0. | IS Compliance |
| `cbs_rates` | Alíquotas CBS por classe tributária (fase CBS Plena 2027+). | BR-04 |
| `iss_rates` | Alíquotas ISS por município/IBGE. Usado na Fase 4 e relevante para transição IBS. | BR-05 |

### 4.2 Modelo de Alíquotas IBS por Município

O campo `municipio_destino_ibge` em `iva_dual_rules` é o mecanismo central para o princípio do destino (BR-03). A query de lookup prioriza regra municipal sobre estadual:

```sql
SELECT * FROM iva_dual_rules
WHERE ncm = $1 AND uf_destino = $2
  AND (municipio_destino_ibge = $3 OR municipio_destino_ibge IS NULL)
  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
ORDER BY municipio_destino_ibge DESC NULLS LAST
LIMIT 1;
```

Isso garante que:
- Se existe regra específica para o município → usa a regra municipal
- Se não existe → usa a regra estadual (`municipio_destino_ibge IS NULL`)
- Se nenhuma regra → `slog.Warn` e skip do tributo

📄 **Fonte técnica:** [.specs/architecture/erd.md](../.specs/architecture/erd.md)

---

## 5. Diagrama de Contexto do Projeto (C4 — Nível 1)

```
┌─────────────────────────────────────────────────────────────────────────┐
│                     LIMITE DO SISTEMA CORPORATIVO                        │
│                                                                          │
│  ┌──────────┐   ┌──────────┐   ┌──────────┐   ┌──────────────────┐     │
│  │   CRM    │   │E-commerce│   │ Portal   │   │   ERP / SAP      │     │
│  │ (Vendas) │   │  (B2C)   │   │  (B2B)   │   │ (Faturamento)   │     │
│  └────┬─────┘   └────┬─────┘   └────┬─────┘   └────────┬─────────┘     │
│       │              │              │                   │               │
│       └──────────────┴──────────────┴───────────────────┘               │
│                          │ POST /calculate                              │
│                          ▼                                              │
│  ┌──────────────────────────────────────────────────────────────┐      │
│  │         ms-billing-engine-tax-rates (Go/Fiber :3000)          │      │
│  │                                                               │      │
│  │  ┌──────────────────────────────────────────────────┐        │      │
│  │  │        BillingEnginePhased (SOP-013, 7 fases)     │        │      │
│  │  │                                                  │        │      │
│  │  │  F0(IS)→F1(IPI)→F2(CBS)→F3(ICMS)                │        │      │
│  │  │  →F4(IBS+ISS+PISCOFINS)→F5(FUST)→F6(FUNTTEL)    │        │      │
│  │  └──────────────────────────────────────────────────┘        │      │
│  │                                                               │      │
│  │  ┌─────────────┐  ┌──────────────────┐  ┌──────────────┐    │      │
│  │  │PhaseResolver│  │  TaxSelector     │  │IBSRateFetcher│    │      │
│  │  │(F-005)      │  │  (Matriz DT-001) │  │(CircuitBreaker│    │      │
│  │  └─────────────┘  └──────────────────┘  └──────────────┘    │      │
│  └──────────────────────────┬───────────────────────────────────┘      │
│                             │                                          │
│              ┌──────────────┴──────────────┐                           │
│              ▼                             ▼                           │
│  ┌───────────────────┐        ┌──────────────────────┐                 │
│  │    PostgreSQL     │        │        Redis         │                 │
│  │ billing_tax_rates │        │  Cache de Regras     │                 │
│  │ (10 tabelas)      │        │  TTL 24h             │                 │
│  └───────────────────┘        └──────────────────────┘                 │
│                                                                          │
│  ┌──────────────────────────────────────────────────────────┐          │
│  │  Futuro: API Comitê Gestor IBS (Bloqueada — Gap G2)      │          │
│  │  GET /api/v1/rates?ibge_code={code}                       │          │
│  └──────────────────────────────────────────────────────────┘          │
└─────────────────────────────────────────────────────────────────────────┘
```

📄 **Fonte:** [.specs/architecture/c4-context.md](../.specs/architecture/c4-context.md)

---

## 6. Gaps e Roadmap Técnico do Projeto

### 6.1 Funcionalidades Existentes que Atendem ao Projeto

| Capacidade Técnica | Feature ID | BR Atendido | Status |
|:---|:---|:---|:---|
| Motor 7-fases SOP-013 | C-001 | BR-01, BR-04, BR-07 | ✅ Completo |
| Reforma Tributária CBS/IBS/IS | RF-08 | BR-04, BR-05 | ✅ Completo |
| Phase Resolution System | F-005 | Transição 2026–2033 | ✅ Completo |
| IS Pré-Filtro (NCM Seletivo) | F-006 | IS Compliance | ✅ Completo |
| IBS Circuit Breaker + Cache | F-007 | Resiliência regulatória | ✅ Completo |
| W3C Trace Context | RNF-08 | Rastreabilidade ponta a ponta | ✅ Completo |
| Autenticação JWT (Kong/Keycloak) | RNF-09 | Segurança | ✅ Completo |
| Métricas Prometheus | RNF-10 | Monitoramento (KPIs O2) | ✅ Completo |
| ICMS Desonerado | F-004 | Regimes especiais | ✅ Completo |
| Schema SQL Completo | C-002 | Persistência de regras | ✅ Completo |

### 6.2 Gaps Identificados (Funcionalidades Necessárias)

| Gap ID | Descrição | BR | Prioridade | Complexidade |
|:---|:---|:---|:---|:---|
| **GAP-001** | Interface administrativa para Time Fiscal atualizar alíquotas (BR-02) | BR-02 | 🔴 Alta | Média |
| **GAP-002** | TaxToken — congelamento de alíquota por janela temporal (BR-06) | BR-06 | 🔴 Alta | Média |
| **GAP-003** | Endpoint de simulação `/simulate` com projeção de margem (BR-05) | BR-05 | 🟡 Média | Baixa |
| **GAP-004** | Campo `valor_liquido` no schema de resposta (BR-04) | BR-04 | 🟡 Média | Baixa |
| **GAP-005** | Cálculo de créditos na entrada `/credit/calculate` (BR-08) | BR-08 | 🔴 Alta | Alta |
| **GAP-006** | Schema de split payment na resposta (BR-09) | BR-09 | 🔴 Alta | Média |
| **GAP-007** | Qualificação fiscal de fornecedores (BR-08) | BR-08 | 🟡 Média | Alta |
| **GAP-008** | Rate limiting no endpoint `/calculate` (DT-11) | RNF | 🔴 Alta | Baixa |
| **GAP-009** | API versioning `/v1/calculate` | RNF | 🟡 Média | Baixa |
| **GAP-010** | Deploy artifacts (Dockerfile, K8s manifests) (DT-10) | RNF | 🟡 Média | Média |

### 6.3 Dívidas Técnicas Relevantes ao Projeto

| DT ID | Descrição | Impacto no Projeto |
|:---|:---|:---|
| DT-03 | CSTs provisórios (`01`/`04`) para CBS/IBS/IS | Risco de não-conformidade quando RFB publicar tabela oficial |
| DT-04 | Créditos da Reforma (`permite_credito_amplo`) não implementados | Bloqueia BR-08 (rastreabilidade de créditos) |
| DT-11 | Sem rate limiting | Risco de indisponibilidade do cálculo em horário comercial |
| DT-10 | Sem artefatos de deploy | Bloqueia deploy em produção para homologação |

📄 **Fonte:** [.specs/product/feature-roadmap.md](../product/feature-roadmap.md)

---

## 7. Referências Cruzadas Completas

### 7.1 Documentos de Arquitetura do Microserviço (Fonte da Verdade)

| Documento | Conteúdo | Relevância para o Projeto |
|:---|:---|:---|
| [.specs/architecture/architecture.md](../.specs/architecture/architecture.md) | Visão arquitetural completa, regras de solução, pipeline SOP-013, injeção inter-fase | ⭐ Base técnica do motor de cálculo |
| [.specs/architecture/c4-context.md](../.specs/architecture/c4-context.md) | Diagrama de contexto C4, atores, fluxo de requisição | Contexto de integração |
| [.specs/architecture/erd.md](../.specs/architecture/erd.md) | Modelo de dados completo (10 tabelas, triggers, índices) | Estrutura de persistência das regras fiscais |
| [.specs/architecture/integrations.md](../.specs/architecture/integrations.md) | Dependências, variáveis de ambiente, bibliotecas, IBS API | Configuração de ambiente e integrações |

### 7.2 Documentos de Domínio e Produto

| Documento | Conteúdo |
|:---|:---|
| [.specs/domain/domain.md](../domain/domain.md) | Glossário fiscal, 13 regras de domínio, pipeline SOP-013 detalhado |
| [.specs/product/requirements.md](../product/requirements.md) | 11 RFs + 15 RNFs com evidência de código |
| [.specs/product/feature-roadmap.md](../product/feature-roadmap.md) | Features implementadas, planejadas e 17 dívidas técnicas |

### 7.3 Documentos de Negócio do Projeto

| Documento | Conteúdo |
|:---|:---|
| [PROJECT-CHARTER.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/PROJECT-CHARTER.md) | Termo de abertura, objetivos, escopo, riscos, stakeholders |
| [REQUIREMENTS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/REQUIREMENTS.md) | 9 BRs com matriz de rastreabilidade |
| [EPICS-01-COMMERCIAL-CHANNELS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/EPICS-01-COMMERCIAL-CHANNELS.md) | Onda 1: 3 Épicos comerciais |
| [EPICS-02-FINANCIAL-BILLING-ERP.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/EPICS-02-FINANCIAL-BILLING-ERP.md) | Onda 2: 3 Épicos financeiros |
| [MATRIZ-KPI.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/MATRIZ-KPI.md) | 8 KPIs em 3 dimensões |
| [PRD.md](./PRD.md) | Resumo de alto nível do produto (este projeto) |

---

## 8. Considerações de Infraestrutura e Deploy

### 8.1 Variáveis de Ambiente Específicas do Projeto

| Variável | Uso no Projeto | Status |
|:---|:---|:---|
| `DATABASE_URL` | Conexão PostgreSQL — schema `billing_tax_rates` com tabelas da reforma | ✅ Existente |
| `REDIS_ADDR` | Cache Redis — chaves `tax:iva:*` e `ibs:rate:*` | ✅ Existente |
| `PORT` | Porta HTTP (default `:3000`) | ✅ Existente |
| `IBS_API_BASE_URL` | API do Comitê Gestor IBS (não publicada — Gap G2) | ⚠️ Planejada |
| `TAX_TOKEN_TTL_MINUTES` | TTL do token de garantia comercial (GAP-002) | 🔴 Nova |
| `RATE_LIMIT_MAX` | Rate limiting para proteção do endpoint (GAP-008) | 🔴 Nova |

### 8.2 Requisitos de SLA para KPIs do Projeto

| KPI de Negócio | Métrica Técnica | SLA |
|:---|:---|:---|
| **O2:** Latência da API de Cálculo | `http_request_duration_seconds` (p95) | < 100ms |
| **C1:** Rejeição de Notas | `errors_total{type="validation"}` | < 0,1% |
| **F1:** Aproveitamento de Créditos | A ser definido com GAP-005 | ≥ 98% |
| **F3:** Divergência Split Payment | A ser definido com GAP-006 | R$ 0,00 |

📄 **Fonte KPIs:** [MATRIZ-KPI.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/MATRIZ-KPI.md)

---

> 📋 **Próximo passo:** Priorizar os GAPs da Seção 6.2 para implementação, começando pelos de prioridade Alta (GAP-001, GAP-002, GAP-005, GAP-006, GAP-008) que são bloqueantes para as Ondas 1 e 2 do projeto.
