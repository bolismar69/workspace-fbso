---
title: "Domínio — Solar Fácil Site"
version: "2.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["domain", "glossary", "ubiquitous-language", "energia-solar"]
---

# Glossário de Domínio — Solar Fácil Site

## 1. Introdução

Este documento define a **linguagem ubíqua** (Ubiquitous Language) do domínio Solar Fácil — os termos que aparecem no código, na interface e nas conversas com stakeholders, com o mesmo significado em todos os contextos.

**Escopo do domínio**: Plataforma digital que conecta produtores de energia solar excedente com consumidores que desejam desconto na conta de luz, operando via cooperativas regulamentadas pela ANEEL. O site é a porta de entrada: educa, calcula economia, apresenta planos e converte visitantes em leads.

**Fontes**: Código-fonte (`src/lib/types.ts`, `src/lib/constants.ts`), documentação existente (`.specs/PRODUCT.md`, `.specs/ARCHITECTURE.md`).

---

## 2. Termos de Domínio

### 2.1 Atores (Personas)

| Termo | Definição | Sinônimos | Código (interface/type/enum) |
|---|---|---|---|
| **Consumidor** | Pessoa física ou pequena empresa que quer reduzir a conta de luz sem instalar painéis solares. Não é especialista em energia; busca economia sem complexidade. | Cliente, Usuário consumidor | `PersonaProfile = 'consumidor'` |
| **Fornecedor / Produtor** | Pessoa ou empresa que já tem painéis solares instalados e gera excedente de energia. Quer rentabilizar esse excedente de forma simples e legalizada. | Produtor, Gerador | `PersonaProfile = 'fornecedor'` |
| **Cooperativa** | Parceiro B2B que opera a distribuição de energia entre produtores e consumidores, dentro da regulação ANEEL. | Cooperativa de energia | `PersonaProfile = 'cooperativa'` |

### 2.2 Conceitos de Negócio Core

| Termo | Definição | Sinônimos | Código (interface/type/enum) |
|---|---|---|---|
| **Plano** | Assinatura mensal que define a capacidade de energia compartilhada (em kWh) e os benefícios incluídos. | Assinatura, Membership | `Plan` (interface), `PlanName` (type) |
| **Capacidade (kWh)** | Faixa de consumo mensal em quilowatt-hora que o plano cobre. Determina o plano sugerido pela calculadora. | Faixa de consumo | `Plan.capacityKwh: { min: number; max: number }` |
| **Economia** | Valor em reais que o consumidor deixa de pagar na conta de luz ao aderir à energia compartilhada. Calculado como 12% do valor da conta. | Desconto, Saving | `ConsumerResult.economy` |
| **Ganho** | Valor em reais que o fornecedor recebe por compartilhar seu excedente de energia. Calculado a R$ 0,40 por kWh. | Rendimento, Receita | `ProviderResult.gain` |
| **Taxa de desconto** | Percentual fixo de 12% aplicado sobre a conta de luz do consumidor para calcular a economia. | Discount rate | `CONSUMER_DISCOUNT_RATE = 0.12` |
| **Taxa do fornecedor** | Valor fixo de R$ 0,40 por kWh excedente compartilhado, pago ao fornecedor. | Provider rate | `PROVIDER_RATE = 0.4` |
| **Excedente de energia** | Energia solar gerada pelo fornecedor que excede seu próprio consumo e pode ser compartilhada. | Surplus, Excedente | Parâmetro `monthlySurplusKwh` em `calculateProviderGain()` |

### 2.3 Calculadora

| Termo | Definição | Sinônimos | Código (interface/type/enum) |
|---|---|---|---|
| **Calculadora do Consumidor** | Ferramenta que estima a economia mensal com base no valor da conta de luz. Retorna economia em R$ e plano sugerido. | Simulador de economia | `ConsumerCalculator`, `calculateConsumerEconomy()` |
| **Calculadora do Fornecedor** | Ferramenta que estima o ganho mensal com base no excedente de energia (kWh). Retorna ganho em R$. | Simulador de ganho | `ProviderCalculator`, `calculateProviderGain()` |
| **Outlier** | Valor de entrada (conta de luz ou kWh) fora da faixa típica atendida pelos planos. Retorna mensagem explicativa em vez de cálculo. | Valor atípico, Fora da faixa | `ConsumerResult.isOutlier`, `ProviderResult.isOutlier` |
| **Plano sugerido** | Plano recomendado automaticamente pela calculadora com base na faixa de consumo (tabela de decisão DT-001). | Recommended plan | `ConsumerResult.suggestedPlan` |

### 2.4 Fluxo de Contato/Lead

| Termo | Definição | Sinônimos | Código (interface/type/enum) |
|---|---|---|---|
| **Lead** | Visitante que preencheu o formulário de contato — potencial cliente. | Prospect, Contato | `LeadForm` (interface) |
| **Formulário de contato** | Formulário com nome, email, telefone, perfil e mensagem. Enviado via Formspree (placeholder). | Contact form | `useContactForm`, `ContactForm` |
| **Journey Summary** | Resumo do contexto do usuário antes de chegar ao formulário (ex: "Você veio da calculadora — seu plano sugerido é Special"). | Contexto da jornada | `JourneySummary` component |
| **Honeypot** | Campo hidden `website` no formulário para detectar bots. Bots preenchem; humanos não. | Anti-spam field | `HONEYPOT_FIELD = 'website'` |
| **Tempo mínimo de submissão** | 3 segundos entre carregar a página e submeter o formulário — anti-spam. | Anti-spam timer | `MIN_SUBMIT_TIME_MS = 3000` |

### 2.5 Métricas e Prova Social

| Termo | Definição | Sinônimos | Código (interface/type/enum) |
|---|---|---|---|
| **Métrica** | Par valor + label usado em cards de prova social e diferenciais. | Stat, KPI | `Metric` (interface) |
| **Diferencial** | Característica competitiva da Solar Fácil (ex: "Zero Capex", "Plataforma Legal ANEEL"). | Competitive advantage | `DIFFERENTIATORS: Metric[]` |

### 2.6 Analytics

| Termo | Definição | Sinônimos | Código (interface/type/enum) |
|---|---|---|---|
| **Evento GA4** | Evento disparado para o Google Analytics 4 via `window.gtag()`. | Tracking event | `AnalyticsEvent` (type) |
| **CTA Click** | Clique em botão de call-to-action (consumidor ou fornecedor, hero ou final). | Call-to-action | `cta_click` event |
| **Calculator Use** | Uso da calculadora (consumidor ou fornecedor) com valor de entrada e resultado. | Simulação | `calculator_use` event |
| **Lead Capture** | Formulário de contato enviado com sucesso. | Conversão | `lead_capture` event (definido, não conectado) |

### 2.7 Regulatório

| Termo | Definição | Sinônimos | Código (interface/type/enum) |
|---|---|---|---|
| **ANEEL** | Agência Nacional de Energia Elétrica — órgão regulador do setor elétrico brasileiro. | Agência reguladora | Mencionado em `constants.ts` e FAQ |
| **Resolução Normativa 687/2015** | Norma da ANEEL que regulamenta a geração distribuída e o compartilhamento de energia solar. | RN 687/2015 | Mencionado em FAQ `constants.ts:114` |

---

## 3. Relações entre Conceitos

```
┌──────────────┐     Calcula      ┌─────────────────┐
│  Consumidor  │─────────────────▶│  Economia (R$)   │
│  (Persona)   │                  │  12% da conta    │
└──────────────┘                  └────────┬────────┘
                                           │ Sugere
                                           ▼
                                  ┌─────────────────┐
                                  │  Plano           │
                                  │  Basic/Special/  │
                                  │  Premium         │
                                  └────────┬────────┘
                                           │ Contrata
                                           ▼
┌──────────────┐   Compartilha   ┌─────────────────┐
│  Fornecedor  │────────────────▶│  Excedente (kWh) │
│  (Persona)   │                  │  R$ 0,40/kWh    │
└──────────────┘                  └────────┬────────┘
                                           │ Via
                                           ▼
                                  ┌─────────────────┐
                                  │  Cooperativa     │
                                  │  (ANEEL)         │
                                  └─────────────────┘
```

### Fluxo Principal

1. **Consumidor** acessa o site → usa **Calculadora do Consumidor** → vê **Economia** estimada → recebe **Plano Sugerido** → preenche **Formulário de Contato** → torna-se **Lead**
2. **Fornecedor** acessa o site → usa **Calculadora do Fornecedor** → vê **Ganho** estimado → preenche **Formulário de Contato** → torna-se **Lead**
3. **Cooperativa** gerencia a distribuição da energia entre fornecedores e consumidores (fora do escopo do site)

---

## 4. Regras de Negócio Fundamentais

### 4.1 Cálculo de Economia (Consumidor)

| Ref | Regra | Implementação |
|---|---|---|
| BR-DER-001 | Economia = conta mensal × 12% | `calculateConsumerEconomy()` |
| BR-DER-005 | Conta < R$ 50 → outlier (abaixo da faixa) | `CONSUMER_OUTLIER_MIN = 50` |
| BR-DER-006 | Conta > R$ 5.000 → outlier (acima da faixa) | `CONSUMER_OUTLIER_MAX = 5000` |

### 4.2 Cálculo de Ganho (Fornecedor)

| Ref | Regra | Implementação |
|---|---|---|
| BR-DER-003 | Ganho = kWh excedente × R$ 0,40 | `calculateProviderGain()` |
| BR-DER-007 | < 50 kWh/mês → outlier | `PROVIDER_OUTLIER_MIN = 50` |
| BR-DER-008 | > 10.000 kWh/mês → outlier | `PROVIDER_OUTLIER_MAX = 10000` |

### 4.3 Sugestão de Plano (Tabela de Decisão DT-001)

| Faixa de Consumo (kWh) | Plano Sugerido |
|---|---|
| 100–200 | **Basic** (R$ 150/mês) |
| 200–350 | **Special** (R$ 250/mês) |
| 350–600 | **Premium** (R$ 400/mês) |
| Fora das faixas | `null` (outlier) |

### 4.4 Validação de Formulário

| Ref | Regra |
|---|---|
| BR-CON-001 | Nome: obrigatório, mínimo 2 caracteres |
| BR-CON-002 | Email: obrigatório, formato válido |
| BR-CON-003 | Telefone: opcional, 10-11 dígitos (DDD + número) |
| BR-CON-004 | Perfil: obrigatório (consumidor/fornecedor/cooperativa) |
| BR-CON-005 | Mensagem: opcional, máximo 1000 caracteres |
| BR-CON-006 | Anti-spam: tempo mínimo de 3s entre page load e submit |

---

## 5. Termos Fora de Escopo (Não aparecem no código)

| Termo | Motivo da ausência |
|---|---|
| **Fatura / Conta de luz** | Implícito no termo "monthlyBill" (conta mensal) |
| **Geração distribuída** | Termo técnico ANEEL — não usado na interface (usuário não precisa saber) |
| **Crédito de energia** | Mecanismo regulatório — abstraído pelo modelo Solar Fácil |
| **Usina solar / Fazenda solar** | Não aplicável — o modelo é compartilhamento, não geração centralizada |

---

Última atualização: 2026-07-08
