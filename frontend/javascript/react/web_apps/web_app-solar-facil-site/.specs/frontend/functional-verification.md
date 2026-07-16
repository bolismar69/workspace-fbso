# Functional Verification — Solar Fácil Site

> Verificação funcional de fluxos principais com base em análise de código.
> Gerado por `webapp-testing` em 2026-07-08.
> ⚠️ Verificação com Playwright indisponível (Chrome não instalado).

---

## 1. Fluxos de Usuário

### Fluxo 1: Consumidor — Calculadora → Plano → Contato

```
Home (/) → Hero → CTA "Quero Economizar"
  → ConsumerCalculator
    → Input: valor da conta de luz
    → Click "Calcular"
    → Resultado: economia em R$ + plano sugerido
  → PlansSection
    → Ver cards de planos
  → FinalCtaSection
    → CTA "Começar Agora"
  → /contato
    → JourneySummary (contexto da calculadora)
    → Preencher formulário
    → Submit → SuccessScreen
```

**Verificação de código:**
- [x] `calculateConsumerEconomy()` — função pura, testável
- [x] `suggestPlan()` — tabela de decisão DT-001 implementada
- [x] `useCalculator` — estados: input, result, error, hasCalculated
- [x] `trackCalculatorUse()` — evento GA4 conectado
- [x] `JourneySummary` — recebe `initialProfile` e `initialMessage`

### Fluxo 2: Fornecedor — Calculadora → Contato

```
Home (/) → Hero → CTA "Quero Compartilhar" (ou tab "Fornecedor")
  → ProviderCalculator
    → Input: kWh excedente mensal
    → Click "Calcular"
    → Resultado: ganho em R$
  → /contato
    → Preencher formulário
    → Submit → SuccessScreen
```

**Verificação de código:**
- [x] `calculateProviderGain()` — função pura, testável
- [x] Outliers tratados (<50 kWh, >10000 kWh)
- [x] `useCalculator('provider')` — compartilha hook com consumer

### Fluxo 3: Navegação — Home → Planos → Contato

```
Home (/) → Header → Link "Planos"
  → /planos
    → PlansComparisonTable
    → ProviderHighlight
    → FaqAccordion (expandir/fechar)
  → Header → Link "Contato"
  → /contato
    → ContactForm + DirectChannels
  → Footer → Links
```

**Verificação de código:**
- [x] 3 rotas configuradas no App Router
- [x] Header com links corretos
- [x] Footer com sitemap links
- [x] `useFaqAccordion` — estado de accordion

### Fluxo 4: Tratamento de Erros

```
Calculadora:
  → Input vazio → "Informe um valor para calcular"
  → Valor < 50 → outlier message
  → Valor > 5000 → outlier message

Formulário:
  → Nome vazio → "Nome é obrigatório"
  → Email inválido → "Formato de e-mail inválido"
  → Submit < 3s → fake success (anti-spam)
  → Rede falha → "Não foi possível enviar..."
```

**Verificação de código:**
- [x] `validateForm()` — 5 validações implementadas
- [x] Anti-spam: honeypot + timer 3s
- [x] Error handling com mensagens pt-BR
- [x] Fallback para WhatsApp em caso de falha

---

## 2. Cobertura de Estados por Componente

| Componente | Default | Loading | Empty | Error | Success | Disabled | Focus | Hover |
|---|---|---|---|---|---|---|---|---|
| **ConsumerCalculator** | ✅ | — | — | ✅ | ✅ | ✅ | ✅ | ✅ |
| **ProviderCalculator** | ✅ | — | — | ✅ | ✅ | ✅ | ✅ | ✅ |
| **PlansSection** | ✅ | ✅ | — | ✅ | — | — | — | ✅ |
| **FaqAccordion** | ✅ | — | — | — | — | — | ✅ | ✅ |
| **ContactForm** | ✅ | ✅ | — | ✅ | ✅ | ✅ | ✅ | ✅ |
| **MobileMenu** | ✅ | — | — | — | — | — | — | — |

---

## 3. Testes Manuais (a executar com navegador)

### Pré-condição
```bash
npm run dev  # http://localhost:3000
```

### Roteiro de Teste

```markdown
1. [ ] HOME — Abrir http://localhost:3000
   - Header visível com logo + links
   - Hero com headline + 2 CTAs
   - Scroll até Calculadora
   
2. [ ] CALCULADORA CONSUMIDOR
   - Input: deixar vazio, clicar Calcular → erro "Informe um valor"
   - Input: "500", clicar Calcular → economia R$ 60,00 + plano Special
   - Input: "30", clicar Calcular → mensagem outlier
   - Input: "6000", clicar Calcular → mensagem outlier

3. [ ] CALCULADORA FORNECEDOR
   - Clicar tab "Fornecedor"
   - Input: "300", clicar Calcular → ganho R$ 120,00
   - Input: "30", clicar Calcular → mensagem outlier

4. [ ] PLANOS — Navegar para /planos
   - 3 cards de plano visíveis
   - Special destacado como "Mais Popular"
   - FAQ: clicar nas perguntas → expandir respostas

5. [ ] CONTATO — Navegar para /contato
   - JourneySummary mostra contexto (se veio da calculadora)
   - Enviar vazio → erros de validação
   - Preencher nome (1 char) → erro "mínimo 2 caracteres"
   - Preencher email inválido → erro formato
   - Preencher corretamente → submit → SuccessScreen

6. [ ] MOBILE — Redimensionar para 375px
   - Header vira hamburger menu
   - Menu abre/fecha
   - Calculadora empilha verticalmente
   - Planos em 1 coluna
```

---

Última atualização: 2026-07-08
