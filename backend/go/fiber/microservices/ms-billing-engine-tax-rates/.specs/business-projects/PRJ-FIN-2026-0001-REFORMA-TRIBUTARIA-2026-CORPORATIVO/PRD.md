# Product Requirements Document (PRD)
## Programa de Adequação Corporativa à Reforma Tributária Nacional

**Código:** PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO
**Versão do PRD:** 2.0
**Data:** 30 de Junho de 2026
**Status:** ✅ Fases 0-1-2 implementadas (PR #6 merged)
**Tipo:** Resumo de Alto Nível — porta de entrada para a documentação de negócio completa

> ⚠️ **Aviso de Leitura:** Este documento é um sumário executivo. Todas as especificações detalhadas, regras de negócio, critérios de aceite BDD e user stories completas residem nos documentos-fonte referenciados na Seção 4.

---

## 1. Visão Geral do Produto

A **Reforma Tributária Brasileira** introduz o modelo de **IVA Dual** — composto pela **CBS** (Contribuição sobre Bens e Serviços, federal) e pelo **IBS** (Imposto sobre Bens e Serviços, estadual/municipal) — além do **Imposto Seletivo (IS)** sobre produtos com externalidades negativas. Esta transição constitucional, com período híbrido de convivência entre regimes antigo e novo de **2029 a 2032**, representa a maior transformação operacional da história da companhia.

A empresa opera sob o regime de **Lucro Real**, com sede em **Santana de Parnaíba (SP)** e presença comercial em todos os estados brasileiros. O programa abrange três dimensões críticas:

| Dimensão | Escopo |
|:---|:---|
| **Comercial (Onda 1)** | Experiência do cliente, simulação de preços "por fora", transparência fiscal nos canais de venda |
| **Financeira (Onda 2)** | Faturamento, split payment, escrituração contábil, apropriação de créditos no Lucro Real |
| **Estratégica** | Gestão do período híbrido, monitoramento regulatório, descontinuação de obrigações acessórias legadas |

📄 **Documento-fonte:** [01-PROJECT-CHARTER.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/01-PROJECT-CHARTER.md)

---

## 2. Objetivos Estratégicos de Negócio

Os 6 objetivos que direcionam todas as iniciativas do programa:

1. **Unificação da Inteligência de Preços** — consistência absoluta entre simulação comercial e liquidação financeira
2. **Preservação de Margens** — mecanismos de precificação que absorvam ou repassem estrategicamente as novas alíquotas regionais de IBS
3. **Otimização de Fluxo de Caixa e Créditos** — captura máxima de créditos não cumulativos do Lucro Real e preparação para o split payment
4. **Sustentabilidade das Operações Nacionais** — continuidade de vendas e faturamento em todas as UFs, sem interrupções por inconformidades fiscais
5. **Conformidade com o Imposto Seletivo** — classificação correta de produtos sujeitos ao IS e repasse adequado ao preço
6. **Integridade Ponta a Ponta** — unicidade das regras de preço e tributação em toda a jornada da operação nacional

📄 **Documento-fonte:** [01-PROJECT-CHARTER.md — Seção 2](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/01-PROJECT-CHARTER.md)

---

## 3. Requisitos de Negócio (Resumo)

Os 9 Requisitos de Negócio (BR-01 a BR-09) estão organizados em 3 blocos:

### Bloco 1 — Fundação: Dados, Alíquotas e Governança

| ID | Requisito | Essência |
|:---|:---|:---|
| BR-01 | Centralização da Inteligência de Regras | Fonte única da verdade para cálculo de impostos em todos os canais |
| BR-02 | Autonomia do Time Fiscal (No-Code) | Atualização dinâmica de alíquotas sem intervenção de desenvolvimento |
| BR-03 | Qualificação Geográfica de Cadastro | Validação obrigatória do código IBGE do destino no início da transação |

### Bloco 2 — Onda 1: Jornada Comercial

| ID | Requisito | Essência |
|:---|:---|:---|
| BR-04 | Transparência e Cálculo "Por Fora" | Exibição do preço base + IVA Dual segregado para o cliente |
| BR-05 | Proteção de Margem e Simulação | Visibilidade do preço líquido e margem real antes do fechamento de contratos |
| BR-06 | Garantia de Preço Ofertado (Token) | Congelamento do valor tributário simulado por janela temporal definida |

### Bloco 3 — Onda 2: Operação Financeira e Faturamento

| ID | Requisito | Essência |
|:---|:---|:---|
| BR-07 | Unicidade Pedido ↔ Nota Fiscal | 100% de consistência matemática entre simulação e faturamento |
| BR-08 | Rastreabilidade de Créditos no Lucro Real | Mapeamento do potencial de crédito de CBS/IBS por operação |
| BR-09 | Viabilização do Split Payment | Discriminação exata entre receita líquida da empresa e parcelas CBS/IBS retidas |

📄 **Documento-fonte:** [02-BUSINESS-REQUIREMENTS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/02-BUSINESS-REQUIREMENTS.md)

---

## 4. Hierarquia da Documentação e Rastreabilidade

```
[Estratégico]  01-PROJECT-CHARTER.md
       │
       ▼
[Escopo]       02-BUSINESS-REQUIREMENTS.md  (BR-01 a BR-09)
       │
       ▼
[Macro]  03-EPICS.md (6 Épicos unificados)
       │
       ├──────────────────────┐
       ▼                      ▼
      Onda 1                  Onda 2
   (Épicos 01.01-01.03)   (Épicos 02.01-02.03)
       │                      │
       ▼                      ▼
[Func]  04-FEATURES.md (16 Features unificadas)
       │
       ├──────────────────────┐
       ▼                      ▼
   Onda 1 (7 Features)    Onda 2 (9 Features)
       │                      │
       ▼                      ▼
[Exec]  21 User Stories        20 User Stories
      (05-USER-STORYS-01-*)  (05-USER-STORYS-02-*)
```

### 4.1 Documentos da Onda 1 — Canais Comerciais

| Nível | Documento | Conteúdo |
|:---|:---|:---|
| Épicos | [03-EPICS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/03-EPICS.md) | Épicos 01.01 (CRM), 01.02 (Integração), 01.03 (Precificação) |
| Features | [04-FEATURES.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/04-FEATURES.md) | 7 Features detalhadas da Onda 1 |
| User Stories | 21 USs em `05-USER-STORYS-01-*.md` (7 arquivos) | Critérios de aceite BDD completos |

### 4.2 Documentos da Onda 2 — Finanças, Faturamento e ERP

| Nível | Documento | Conteúdo |
|:---|:---|:---|
| Épicos | [03-EPICS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/03-EPICS.md) | Épicos 02.01 (Faturamento), 02.02 (Split), 02.03 (Créditos) |
| Features | [04-FEATURES.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/04-FEATURES.md) | 9 Features detalhadas da Onda 2 |
| User Stories | 20 USs em `05-USER-STORYS-02-*.md` (9 arquivos) | Critérios de aceite BDD completos |

### 4.3 Governança e Métricas

| Documento | Conteúdo |
|:---|:---|
| [MATRIZ-KPI.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/MATRIZ-KPI.md) | 8 KPIs em 3 dimensões (Financeira, Compliance, Operacional) |
| [README.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/README.md) | Índice completo com diagrama de relacionamentos |

### 4.4 Base de Conhecimento Técnico-Tributário

Documentos em `docs-suporte/` fornecem a fundamentação técnica:
- Catálogo de impostos PJ (ICMS, ISS, PIS/COFINS, IPI)
- Regras de DIFAL, ST, FCP
- Tabelas CST, CSOSN, CFOP
- Simples Nacional (Anexos I-V)
- Constantes fiscais e códigos IBGE
- Material do curso RTC CFC-RFB (11 PDFs)

---

## 5. Cronograma Macro

| Período | Fase | Foco |
|:---|:---|:---|
| **Mês 1–2** | Diagnóstico e Precificação | Homologação da matriz de alíquotas nacionais, políticas de preço base |
| **Mês 3–4** | Onda 1 — Ativação Comercial | Simulação e exibição transparente dos novos tributos nos canais de venda |
| **Mês 5–6** | Onda 2 — Ativação Financeira | Faturamento IVA Dual, apropriação de créditos, split payment |
| **2026** | Shadow Run | CBS (0,9%) e IBS (0,1%) — validação de cálculos sem impacto material |
| **2027** | Extinção PIS/COFINS | Início pleno da CBS, transição da carga federal |
| **2029–2032** | Período Híbrido | Convivência ICMS/ISS + CBS/IBS, dupla apuração |
| **2033** | Full IVA Dual | Operação estabilizada, descontinuação de processos legados |

📄 **Documento-fonte:** [01-PROJECT-CHARTER.md — Seção 4](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/01-PROJECT-CHARTER.md)

---

## 6. KPIs e Critérios de Sucesso

### KPIs Financeiros
| KPI | Meta |
|:---|:---|
| **F1:** Aproveitamento de Créditos do IVA Dual | ≥ 98% das entradas elegíveis |
| **F2:** Margem Líquida por UF | Oscilação ≤ ±0,5% vs. orçamento pré-reforma |
| **F3:** Retenção Indevida no Split Payment | R$ 0,00 de divergência |

### KPIs de Compliance
| KPI | Meta |
|:---|:---|
| **C1:** Taxa de Rejeição de NF-e/NFS-e | < 0,1% (apenas falhas técnicas) |
| **C2:** Perda de Subvenção de Investimento | 100% de conformidade na escrituração |

### KPIs Operacionais
| KPI | Meta |
|:---|:---|
| **O1:** Higienização Cadastral (Código IBGE) | 100% dos clientes ativos |
| **O2:** Latência da API de Cálculo Fiscal | < 100ms por requisição |

### Critérios Globais de Sucesso
- Zero interrupção comercial em todas as fases de transição
- Alinhamento contábil integral entre valores projetados e arrecadação real
- Ausência de notificações ou autuações de órgãos reguladores
- Captura mínima de 95% dos créditos não cumulativos disponíveis
- Satisfação do cliente > 80% quanto à clareza da carga tributária
- Processos de dupla apuração operacionais até dezembro/2028

📄 **Documento-fonte:** [MATRIZ-KPI.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/MATRIZ-KPI.md) e [01-PROJECT-CHARTER.md — Seção 8](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/01-PROJECT-CHARTER.md)

---

## 7. Principais Riscos de Negócio

| Risco | Impacto | Mitigação |
|:---|:---|:---|
| Divergência Oferta ↔ Liquidação por flutuação de alíquotas | Alto | Token de garantia de cotação com janela temporal |
| Perda de margem por fornecedores não qualificados (sem crédito CBS/IBS) | Crítico | Due diligence fiscal pré-contratação e programa de qualificação |
| Compressão do capital de giro pelo split payment | Crítico | Projeções de fluxo de caixa, linhas de crédito-ponte |
| Complexidade do período híbrido (2029–2032) | Crítico | Planejamento antecipado da dupla apuração desde 2026 |
| Indisponibilidade de alíquotas IBS por município | Alto | Réplica local da matriz com atualização periódica |

📄 **Documento-fonte:** [01-PROJECT-CHARTER.md — Seção 7](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/01-PROJECT-CHARTER.md)

---

## 8. Estatísticas do Projeto

| Métrica | Valor |
|:---|:---|
| Requisitos de Negócio (BR) | 9 |
| Épicos | 6 (3 Onda 1 + 3 Onda 2) |
| Features | 16 (7 Onda 1 + 9 Onda 2) |
| User Stories | 41 (21 Onda 1 + 20 Onda 2) |
| Regras de Negócio (RN) | ~90 |
| Critérios de Aceite BDD | ~55 cenários |
| KPIs | 8 (3 Financeiros + 2 Compliance + 2 Operacionais) |

---

## 9. Stakeholders Principais

| Papel | Responsabilidade |
|:---|:---|
| Patrocinadores (CFO, COO) | Validação de diretrizes de rentabilidade, liberação orçamentária |
| Comitê Fiscal e Jurídico | Regras do Lucro Real, alíquotas, acompanhamento do Comitê Gestor do IBS |
| Lideranças Comerciais | Impacto nos preços, relação com clientes, estratégia de repasse/absorção |
| CPO / Supply Chain | Política de compras, qualificação fiscal de fornecedores |
| Tesouraria e Controladoria | Modelagem do split payment, fluxo de caixa, liquidez |
| PMO Corporativo | Coordenação das ondas, riscos, desdobramento para execução técnica |

---

## 10. Referências Cruzadas

| Este PRD referencia | Documento-fonte |
|:---|:---|
| Seção 1 — Visão Geral | [01-PROJECT-CHARTER.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/01-PROJECT-CHARTER.md) |
| Seção 2 — Objetivos | [01-PROJECT-CHARTER.md — Seção 2](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/01-PROJECT-CHARTER.md) |
| Seção 3 — Requisitos | [02-BUSINESS-REQUIREMENTS.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/02-BUSINESS-REQUIREMENTS.md) |
| Seção 4 — Hierarquia | [README.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/README.md) (índice completo) |
| Seção 5 — Cronograma | [01-PROJECT-CHARTER.md — Seção 4](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/01-PROJECT-CHARTER.md) |
| Seção 6 — KPIs | [MATRIZ-KPI.md](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/MATRIZ-KPI.md) |
| Seção 7 — Riscos | [01-PROJECT-CHARTER.md — Seção 7](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/01-PROJECT-CHARTER.md) |
| Seção 9 — Stakeholders | [01-PROJECT-CHARTER.md — Seção 5](../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/01-PROJECT-CHARTER.md) |

---

> 📋 **Próximo passo:** Este PRD serve como documento de entrada para o desdobramento técnico. Consulte os documentos-fonte listados acima para as especificações completas de cada área.
