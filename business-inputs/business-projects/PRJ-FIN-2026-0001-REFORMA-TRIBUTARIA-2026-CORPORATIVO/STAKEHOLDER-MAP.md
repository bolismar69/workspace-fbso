# Stakeholder Map

> **Programa:** Adequação Corporativa à Reforma Tributária Nacional
> **Código:** PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO
> **Versão:** 2.0
> **Atualizado:** 2026-07-08

---

## 1. Identificação das Partes Interessadas

### 1.1 Patrocinadores Executivos (Sponsors)

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Sponsor — Diretor Financeiro (CFO) | `<nome>` | Liberação de recursos orçamentários, validação de diretrizes de rentabilidade, aprovação de políticas de crédito e split payment, governança geral do programa | `<email>` |
| Sponsor — Diretor de Operações (COO) | `<nome>` | Continuidade operacional durante a transição, processos de faturamento e supply chain, políticas comerciais interestaduais | `<email>` |

### 1.2 Governança Fiscal e Jurídica

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Comitê Fiscal e Jurídico de Negócios | `<nome>` | Interpretação das regras do Lucro Real, validação de alíquotas aplicadas, acompanhamento do Comitê Gestor do IBS, conformidade regulatória (CBS, IBS, IS), classificação fiscal de produtos (NCM/NBS/CClassTrib) | `<email>` |
| Especialista em Direito Tributário | `<nome>` | Pareceres jurídicos sobre incidência tributária, contencioso fiscal, interpretação de legislação complementar | `<email>` |
| Controladoria Corporativa | `<nome>` | Regras de escrituração contábil, segregação de créditos CBS/IBS, reserva de incentivos, apuração do Lucro Real | `<email>` |

### 1.3 Lideranças de Negócio

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Diretor Comercial / Canais e Mercado | `<nome>` | Estratégias de precificação, políticas de repasse ou absorção de tributos, gestão da carteira de clientes, regras de simulação comercial | `<email>` |
| Gerentes Comerciais Regionais | `<nome>` | Impacto dos preços nas vendas interestaduais, relacionamento com clientes locais, feedback de mercado sobre aceitação de preços | `<email>` |
| Diretor de Suprimentos e Compras (CPO) | `<nome>` | Revisão da política de compras, qualificação fiscal de fornecedores, negociação de contratos para maximização de créditos CBS/IBS, due diligence fiscal pré-contratação | `<email>` |
| Gerência de Tesouraria | `<nome>` | Modelagem do impacto do split payment no capital de giro, projeções de fluxo de caixa, negociação de linhas de crédito-ponte, conciliação bancária (CNAB) | `<email>` |

### 1.4 Governança do Programa

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Gerência Geral do Programa (PMO Corporativo) | `<nome>` | Coordenação das ondas estratégicas, gestão de riscos corporativos, desdobramento de diretrizes de negócios para equipes técnicas, priorização de escopo | `<email>` |
| Product Owner (PO) Corporativo | `<nome>` | Priorização de features e user stories, validação de critérios de aceite, definição de pronto (DoD), trade-offs de escopo vs. prazo | `<email>` |
| Product Manager (PM) Corporativo | `<nome>` | Visão de produto, roadmap de releases, alinhamento estratégico das features com objetivos de negócio | `<email>` |

### 1.5 Execução Técnica

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Tech Lead / Arquiteto de Sistemas | `<nome>` | Decisões de arquitetura, padrões de código, escolha de bibliotecas e frameworks, design de APIs e contratos entre microserviços, aprovação de PRs | `<email>` |
| Tech Lead — ms-billing-engine | `<nome>` | Design do motor de cálculo de impostos (CBS/IBS/IS), pipeline de fases, otimização de performance (SLO <100ms) | `<email>` |
| Tech Lead — ms-tax-rates | `<nome>` | Design do serviço de alíquotas, matriz de alíquotas nacionais, atualização dinâmica (no-code), APIs de consulta | `<email>` |
| SecOps / Segurança da Informação | `<nome>` | Exceções de segurança, validação de SECURITY.md, revisão de vulnerabilidades, autenticação e autorização entre microserviços | `<email>` |
| DevOps / Infraestrutura | `<nome>` | Estratégia de deploy, ambientes (dev/staging/prod), CI/CD, monitoramento e observabilidade | `<email>` |
| QA Lead | `<nome>` | Estratégia de testes, validação do TEST_PLAN.md, aprovação de cobertura de testes, testes de integração entre serviços | `<email>` |

---

## 2. Matriz RACI por Fase do Programa

**Legenda:** R = Responsible (executa) | A = Accountable (aprova/responde) | C = Consulted (consultado) | I = Informed (informado)

### 2.1 Fase de Preparação Corporativa (Mês 1-2: Diagnóstico e Precificação)

| Atividade | CFO | COO | Comitê Fiscal | Controladoria | Dir. Comercial | CPO | Tesouraria | PMO | PO | Tech Lead |
|-----------|-----|-----|---------------|---------------|----------------|-----|------------|-----|----|-----------|
| Homologação da matriz de alíquotas nacionais | I | I | **A** | R | C | — | — | I | C | C |
| Definição de políticas de preços base | C | C | C | C | **A** | — | — | I | R | C |
| Mapeamento de produtos sujeitos ao IS | I | I | **A** | R | C | C | — | I | C | — |
| Qualificação inicial de fornecedores | — | C | C | C | — | **A** | — | I | R | — |
| Projeções de fluxo de caixa (split payment) | C | C | — | C | — | — | **A** | I | — | — |

### 2.2 Onda 1 — Ativação Comercial (Mês 3-4: Canais de Vendas)

| Atividade | CFO | COO | Comitê Fiscal | Controladoria | Dir. Comercial | CPO | Tesouraria | PMO | PO | Tech Lead |
|-----------|-----|-----|---------------|---------------|----------------|-----|------------|-----|----|-----------|
| Simulação de preços e tributos nos canais | I | I | C | C | **A** | — | — | I | R | R |
| Exibição transparente de CBS/IBS/IS | — | — | C | — | **A** | — | — | I | R | R |
| Validação cadastral geográfica (IBGE) | — | C | — | — | **A** | — | — | I | R | R |
| Trava comercial por inconformidade cadastral | — | — | C | — | **A** | — | — | I | R | R |
| Token de garantia comercial | C | — | C | — | **A** | — | — | I | R | R |

### 2.3 Onda 2 — Ativação Financeira (Mês 5-6: Faturamento e ERP)

| Atividade | CFO | COO | Comitê Fiscal | Controladoria | Dir. Comercial | CPO | Tesouraria | PMO | PO | Tech Lead |
|-----------|-----|-----|---------------|---------------|----------------|-----|------------|-----|----|-----------|
| Emissão de faturamento IVA Dual | C | C | C | **A** | I | — | C | I | R | R |
| Apropriação de créditos CBS/IBS (Lucro Real) | C | — | C | **A** | — | C | — | I | R | R |
| Ativação do split payment | C | C | C | C | — | — | **A** | I | R | R |
| Conversão ISS → IBS para serviços | — | — | C | **A** | C | — | — | I | R | R |
| Conciliação bancária (CNAB) | — | — | — | C | — | — | **A** | I | R | R |
| Segregação contábil de créditos | C | — | C | **A** | — | C | — | I | R | R |

### 2.4 Período Híbrido (2029–2032: Convivência de Modelos)

| Atividade | CFO | COO | Comitê Fiscal | Controladoria | Dir. Comercial | CPO | Tesouraria | PMO | PO | Tech Lead |
|-----------|-----|-----|---------------|---------------|----------------|-----|------------|-----|----|-----------|
| Operação simultânea de regimes (legado + IVA Dual) | C | C | **A** | R | C | C | C | I | R | R |
| Monitoramento do Comitê Gestor do IBS | I | I | **A** | R | C | C | C | I | C | C |
| Roadmap de descontinuação de obrigações acessórias | — | — | **A** | R | — | — | — | I | C | R |

---

## 3. Canais de Comunicação e Frequência

| Fórum | Participantes | Frequência | Objetivo | Artefato de saída |
|-------|---------------|------------|----------|-------------------|
| **Comitê Executivo do Programa** | CFO, COO, PMO, Comitê Fiscal | Mensal | Aprovar direcionamento estratégico, liberar recursos, revisar KPIs | Dashboard Executivo (MATRIZ-KPI.md) |
| **Reunião de Governança Fiscal** | Comitê Fiscal, Controladoria, PO | Quinzenal | Validar alíquotas, interpretar mudanças regulatórias, revisar classificação de produtos | Ata + atualização de GLOSSARY.md |
| **Sprint Review (Onda 1 e Onda 2)** | PMO, PO, Tech Leads, QA Lead | Quinzenal (por sprint) | Demonstrar features concluídas, validar critérios de aceite, revisar impedimentos | TASKS.md atualizado + Demo |
| **War Room Comercial** | Dir. Comercial, Gerentes Regionais, PO | Semanal (durante Onda 1) | Feedback de campo sobre precificação, ajustes de simulação, tratamento de exceções | Log de decisões comerciais |
| **War Room Financeira** | Tesouraria, Controladoria, CPO, PO | Semanal (durante Onda 2) | Acompanhar split payment, conciliar divergências, revisar créditos | Log de decisões financeiras |
| **Daily Técnica** | Tech Leads, DevOps, QA Lead | Diário | Sincronizar progresso técnico, destravar impedimentos, revisar métricas de qualidade | Board atualizado |

---

## 4. Caminho de Escalação (Escalation Path)

```
[IMPEDIMENTO TÉCNICO]
    │
    ├─ 1. Resolver no time (Tech Lead + dev)
    │
    ├─ 2. Se impacta escopo/prazo:
    │      Escalar para PO → PMO
    │
    └─ 3. Se impacta orçamento ou diretriz estratégica:
           Escalar para PMO → Sponsors (CFO/COO)

[IMPEDIMENTO DE NEGÓCIO]
    │
    ├─ 1. Resolver com Product Owner (PO)
    │
    ├─ 2. Se requer decisão fiscal:
    │      Escalar para Comitê Fiscal → Controladoria
    │
    ├─ 3. Se impacta cliente/fornecedor:
    │      Escalar para Dir. Comercial (vendas) ou CPO (compras)
    │
    └─ 4. Se impacta resultado financeiro:
           Escalar para CFO

[IMPEDIMENTO REGULATÓRIO]
    │
    ├─ 1. Comitê Fiscal analisa impacto
    │
    ├─ 2. Comitê Fiscal emite parecer
    │
    └─ 3. Se impacto crítico:
           PMO convoca Comitê Executivo extraordinário
```

---

## 5. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 2.0 | 2026-07-08 | Expansão completa: RACI por fase, canais de comunicação, escalation path, roles técnicos | Time Técnico |
| 1.0 | 2026-06-22 | Versão inicial (papéis de negócio) | PMO |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: stakeholder-analysis, agile-ba-practices.*
