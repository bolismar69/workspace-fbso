# Modelo de Dados: Solar Fácil Site

> **Especificação de engenharia reversa** — todos os tipos, constantes e dados mock conforme observado no código-fonte.
> Gerado por `/spec-miner` em 2026-07-05. Revisado para pt-BR em 2026-07-06.

---

## 1. Tipos Core do Domínio

**Arquivo**: `src/lib/types.ts`

### Plan

Representa um plano de assinatura para compartilhamento de energia.

| Campo | Tipo | Descrição | Exemplo |
|-------|------|-----------|---------|
| `name` | `PlanName` (`'Basic' \| 'Special' \| 'Premium'`) | Identificador do plano | `'Special'` |
| `price` | `number` | Preço mensal em BRL | `250` |
| `capacity` | `string` | Faixa de capacidade legível | `'200 a 350 kWh'` |
| `capacityKwh` | `{ min: number; max: number }` | Faixa legível por máquina para matching | `{ min: 200, max: 350 }` |
| `features` | `string[]` | Lista de recursos (3–5 itens) | `['Monitoramento em tempo real', ...]` |
| `highlight` | `boolean` | Se é o plano em destaque/recomendado | `true` |

### Metric

Um ponto de dado estatístico para seções de prova social.

| Campo | Tipo | Descrição | Exemplo |
|-------|------|-----------|---------|
| `value` | `string` | Valor exibido (string permite formatação) | `'500+'` |
| `label` | `string` | Texto descritivo | `'usuários beta'` |
| `icon` | `string?` | Identificador de ícone opcional | `'Shield'` |

### Step

Um passo da timeline para a seção "Como Funciona".

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `icon` | `string` | Nome do ícone Lucide (mapeado no componente) |
| `title` | `string` | Título do passo |
| `description` | `string` | Descrição do passo |

### LeadForm

Modelo de dados do formulário de contato.

| Campo | Tipo | Obrigatório | Validação |
|-------|------|-------------|-----------|
| `name` | `string` | Sim | Mín 2 caracteres |
| `email` | `string` | Sim | Regex `/^[^\s@]+@[^\s@]+\.[^\s@]+$/` |
| `phone` | `string` | Não | Se informado: 10–11 dígitos |
| `profile` | `PersonaProfile \| ''` | Sim | Deve ser não-vazio |
| `message` | `string` | Não | Máx 1000 caracteres |

### FormErrors

Mapa de erros de validação indexado por nome do campo.

```typescript
interface FormErrors {
  name?: string;
  email?: string;
  phone?: string;
  profile?: string;
  message?: string;
}
```

### Resultados da Calculadora

#### ConsumerResult
| Campo | Tipo | Descrição |
|-------|------|-----------|
| `economy` | `number` | Economia mensal estimada em BRL |
| `suggestedPlan` | `PlanName \| null` | Plano correspondente (null se outlier) |
| `isOutlier` | `boolean` | Se o valor está fora da faixa válida |
| `message` | `string` | Mensagem de resultado legível |

#### ProviderResult
| Campo | Tipo | Descrição |
|-------|------|-----------|
| `gain` | `number` | Ganho mensal estimado em BRL |
| `rate` | `number` | Taxa aplicada (R$ 0,40/kWh) |
| `isOutlier` | `boolean` | Se o valor está fora da faixa válida |
| `message` | `string` | Mensagem de resultado legível |

### PersonaProfile

```typescript
type PersonaProfile = 'consumidor' | 'fornecedor' | 'cooperativa';
```

---

## 2. Constantes de Negócio

**Arquivo**: `src/lib/constants.ts`

### Dados dos Planos

| Plano | Preço (BRL) | Capacidade | Qtde Recursos | Destaque |
|-------|-------------|------------|---------------|----------|
| Basic | 150 | 100–200 kWh | 3 | Não |
| Special | 250 | 200–350 kWh | 5 | **Sim** |
| Premium | 400 | 350–600 kWh | 5 | Não |

### Taxas da Calculadora e Limiares de Outlier

| Constante | Valor | Descrição |
|-----------|-------|-----------|
| `PROVIDER_RATE` | 0.40 | R$ 0,40 por kWh (remuneração do fornecedor) |
| `CONSUMER_DISCOUNT_RATE` | 0.12 | 12% de desconto médio |
| `CONSUMER_OUTLIER_MIN` | 50 | Abaixo → outlier (R$ 50) |
| `CONSUMER_OUTLIER_MAX` | 5000 | Acima → outlier (R$ 5.000) |
| `PROVIDER_OUTLIER_MIN` | 50 | Abaixo → outlier (50 kWh) |
| `PROVIDER_OUTLIER_MAX` | 10000 | Acima → outlier (10.000 kWh) |

### Métricas de Prova Social

#### Métricas de Consumidor (`METRICS`)
| Valor | Texto |
|-------|-------|
| `'500+'` | `'usuários beta'` |
| `'4.8 ★'` | `'satisfação (NPS)'` |
| `'12%'` | `'desconto médio'` |

#### Métricas de Fornecedor (`PROVIDER_METRICS`)
| Valor | Texto |
|-------|-------|
| `'3'` | `'cooperativas ativas'` |
| `'1.000 kWh'` | `'compartilhados'` |
| `'R$ 0,40'` | `'por kWh excedente'` |

### Passos "Como Funciona"

| Passo | Ícone | Título | Descrição |
|-------|-------|--------|-----------|
| 1 | `Sun` | Produtor | Gera excedente de energia solar e compartilha via cooperativa |
| 2 | `RefreshCw` | Cooperativa | Gerencia a distribuição da energia entre os participantes |
| 3 | `Home` | Consumidor | Recebe energia limpa com desconto direto na conta de luz |

### Diferenciais

| Valor | Texto |
|-------|-------|
| `'R$ 0'` | Zero Capex — Sem investimento inicial |
| `'100%'` | Pareamento Automático |
| `'100%'` | Plataforma Legal (ANEEL) |
| `'API +'` | ANEEL — Escalável e transparente |

### Itens de FAQ (6 perguntas)

1. "O que é energia solar compartilhada?" → Explicação do modelo ANEEL
2. "Preciso instalar painéis solares?" → Não (consumidores não precisam)
3. "Como funciona o desconto?" → ~12% em média
4. "É legalizado?" → Sim (ANEEL RN 687/2015)
5. "Como recebo sendo fornecedor?" → R$ 0,40/kWh, pagamento mensal
6. "Tem fidelidade?" → Sem lock-in, cancele quando quiser

### URLs e Contato

| Constante | Valor | Status |
|-----------|-------|--------|
| `SITE_URL` | `'https://www.solarfacil.com.br'` | URL de produção |
| `CONTACT_EMAIL` | `'contato@solarfacil.com.br'` | Ativo |
| `WHATSAPP_NUMBER` | `'5511999999999'` | **Placeholder** |
| `APP_STORE_URL` | `'https://apps.apple.com/...'` | **Placeholder** |
| `GOOGLE_PLAY_URL` | `'https://play.google.com/...'` | **Placeholder** |
| `INSTAGRAM_URL` | `'https://instagram.com/solarfacil'` | Ativo |
| `LINKEDIN_URL` | `'https://linkedin.com/company/solarfacil'` | Ativo |

---

## 3. Tipos de Entidade de Domínio (do App Mobile)

**Diretório**: `src/types/`

Estes tipos espelham o modelo de domínio do app mobile e são usados por serviços que buscam de JSON mock.

### ConcessionariaType (`concessionaria.ts`)

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | `number` | Identificador único |
| `name` | `string` | Nome da concessionária (ex.: "Enel Distribuição São Paulo") |
| `status` | `'active' \| 'inactive'` | Status operacional |

### FAQType (`faq.ts`)

| Campo | Tipo |
|-------|------|
| `pergunta` | `string` |
| `resposta` | `string` |

### FAQCategoryType (`faq.ts`)

| Campo | Tipo |
|-------|------|
| `titulo` | `string` |
| `faqs` | `FAQType[]` |

### ConsumoMedioType (`consumo-medio.ts`)

| Campo | Tipo |
|-------|------|
| `id` | `number` |
| `name` | `string` |
| `initialValue` | `number` |
| `finalValue` | `number` |
| `unit` | `string` |
| `description` | `string` |
| `status` | `string` |

---

## 4. Estruturas de Dados Mock

### mockPlans.json (3 registros)

Dados estendidos de plano com campos não usados pelo site (reservados para compatibilidade com o app):
- `icon`, `powerRange`, `consumption`, `pricePerKwh`, `monthlyEstimate`, `energyCost`
- `commercialIndication`, `commercialAttraction`, `description`

### mockFAQs.json (3 categorias × ~3 FAQs cada)

Categorias: "FAQs Gerais", "Para Beneficiários", "Para Fornecedores"

### mockConcessionarias.json (50 registros)

Concessionárias de São Paulo ativas (Enel SP, CPFL Paulista, CPFL Piratininga, CPFL Santa Cruz, EDP SP, Energisa Sul Sudeste, Neoenergia Elektro). Todas as demais marcadas como `inactive`.

### mockConsumoMedio.json (6 faixas)

0–100, 101–200, 201–300, 301–400, 401–500, +500 kWh — todas com `status: "active"`.

---

## 5. Tipos de Eventos de Analytics

**Arquivo**: `src/lib/analytics.ts`

| Evento | Parâmetros | Gatilho |
|--------|-----------|---------|
| `cta_click` | `cta_type`, `location` | Clique em CTA do hero (consumidor/fornecedor) |
| `calculator_use` | `persona`, `input_value`, `result`, `plan_suggested?` | Resultado da calculadora computado |
| `faq_open` | `question_index` | Alternância do accordion FAQ |
| `lead_capture` | `persona`, `has_plan` | Formulário de contato enviado (ainda não conectado) |

---

## 6. Tabela de Decisão de Matching de Planos (DT-001)

**Arquivo**: `src/lib/calculator.ts:33-40`

| Faixa de Conta Mensal (kWh) | Plano Sugerido |
|------------------------------|----------------|
| 100 – 200 | Basic |
| 200 – 350 | Special |
| 350 – 600 | Premium |
| < 100 ou > 600 | `null` (nenhum plano corresponde) |

---

## 7. Transições de Estado

### Máquina de Estado da Calculadora

```
Inicial  →  [usuário digita valor]  →  Pronto
Pronto   →  [clica Calcular]       →  Resultado (válido) ou Erro (inválido)
Resultado → [clica Limpar]         →  Inicial
```

### Máquina de Estado do Formulário de Contato

```
Editando   →  [clica Enviar]     →  Validando
Validando  →  [temErros]         →  Editando (com erros nos campos)
Validando  →  [anti-spam falha]  →  FalsoSucesso (silencioso)
Validando  →  [POST sucesso]     →  Sucesso
Validando  →  [POST falha]       →  Editando (com submitError + fallback WhatsApp)
```

### Máquina de Estado do Accordion FAQ

```
TodosFechados → [clica pergunta N] →  NAberto
NAberto       → [clica pergunta M] →  MAberto (fecha N automaticamente)
NAberto       → [clica pergunta N] →  TodosFechados
Após 3+ FAQs únicos abertos → showContactHint = true
```
