# Requirements: Cálculo de IRPF

> Identificador: `001-calculate-irpf`
> Fonte: handlers/tax_handler.go:20-43, services/calculation_service.go:86-314, services/inss_client.go:26-53

## 1. Resumo executivo

O endpoint de Cálculo de IRPF processa dados financeiros de uma pessoa física para determinar o imposto de renda devido. Ele realiza cálculos simultâneos — via goroutines — nos modelos **Completo** e **Simplificado**, comparando os resultados e marcando automaticamente a opção mais vantajosa (`IsRecommended`) para o contribuinte. Suporta a transição da Reforma Tributária 2026 com regras de isenção total e redução proporcional.

## 2. Priority MoSCoW

| Item | MoSCoW | Justificativa | Evidência |
|------|--------|---------------|-----------|
| RF-01 | Must | Core do motor de cálculo: receber requisição, parsear payload, injetar trace ID | handlers/tax_handler.go:20-43 |
| RF-02 | Must | Execução paralela Completo/Simplificado com recomendação automática | services/calculation_service.go:107-140 |
| RF-03 | Must | Integração com INSS para dedução previdenciária no modelo completo | services/calculation_service.go:150-161, services/inss_client.go:26-53 |
| RF-04 | Must | Motor de cálculo progressivo com aplicação de deduções (dependentes, educação, saúde, PGBL, pensão) | services/calculation_service.go:145-213 |
| RF-05 | Should | Regras de transição da Reforma Tributária 2026 (isenção total + redução proporcional) | services/calculation_service.go:275-296 |
| RF-06 | Should | Resiliência externa: degradação graciosa na falha do INSS | services/calculation_service.go:153-154 |

## 3. Requisitos Funcionais

### RF-01 — Endpoint de Cálculo
**POST /api/v1/calculate/irpf**

O sistema deve expor um endpoint HTTP que aceita um payload JSON com dados financeiros do contribuinte (`models.UniversalTaxRequest`), injeta um `X-Request-ID` via middleware `requestid` do Fiber, e retorna um `map[string]TaxResponse` contendo os resultados de ambos os modelos de cálculo.

**Fonte:** `main.go:52-53`, `handlers/tax_handler.go:20-43`

**Campos do Request (`UniversalTaxRequest`):**
| Campo | Tipo | Obrigatório | Descrição |
|-------|------|-------------|-----------|
| `tax_code` | string | Sim | Código do tributo (ex: "IRPF") |
| `reference_date` | datetime | Não (default: now) | Data de referência para vigência das regras |
| `calculation_type` | string | Sim | "monthly" ou "annual" |
| `gross_income` | decimal | Sim | Renda bruta total |
| `inputs` | array de `{key, value, unit}` | Sim | Deduções e parâmetros de entrada |

**Campos do Response (`TaxResponse`):**
| Campo | Tipo | Descrição |
|-------|------|-----------|
| `reference_date` | string (date) | Data de referência usada no cálculo |
| `gross_income` | decimal(2) | Renda bruta |
| `total_deduction_amount` | decimal(2) | Total de deduções aplicadas |
| `base_value` | decimal(2) | Base de cálculo (renda − deduções) |
| `tax_amount` | decimal(2) | Imposto devido |
| `effective_rate` | decimal(2) | Alíquota efetiva (%) |
| `deduction_details` | array de `{type, amount}` | Detalhamento de cada dedução aplicada |
| `applied_rule` | string | Descrição da faixa/alíquota aplicada |
| `used_configs` | map[string]string | Configurações utilizadas no cálculo |
| `is_recommended` | bool | Indica o modelo mais vantajoso |

### RF-02 — Cálculo Paralelo e Recomendação
O sistema deve executar os modelos Completo e Simplificado em paralelo usando goroutines com channel `chan CalculationResult` de buffer 2. Após ambos retornarem, deve comparar `TaxAmount` e marcar `IsRecommended = true` no modelo de menor imposto.

**Fonte:** `services/calculation_service.go:105-140`

### RF-03 — Integração com INSS
O sistema deve consultar o microserviço externo de INSS via `POST {INSS_SERVICE_URL}/api/v1/calculate/inss`, repassando o `X-Request-ID` para rastreabilidade distribuída. O timeout da chamada HTTP é de 5 segundos. Em caso de falha, o sistema loga um warning (nível `Warn`) e prossegue o cálculo sem a dedução previdenciária (degradação graciosa).

**Fonte:** `services/inss_client.go:14-53`, `services/calculation_service.go:150-161`, `main.go:36-38`

### RF-04 — Motor de Cálculo Progressivo
O sistema deve aplicar as seguintes deduções no modelo Completo, nesta ordem:
1. Dedução de INSS (via integração externa)
2. Pensão oficial (`pension_amount` ou `pension_percentage` × `gross_income`)
3. Dependentes (`dependents_qty` × `dependent_deduction_monthly` ou `annual`)
4. Educação (`education_expenses`, limitado a `education_limit × (1 + dependents_qty)`)
5. Saúde (`health_expenses`, sem teto explícito)
6. PGBL (`pgbl_contribution`, limitado a `pgbl_limit_percentage`% da renda bruta; default 12%)

O modelo Simplificado aplica 20% de desconto sobre a renda bruta, respeitando o teto `simplified_discount_annual_limit` ou `simplified_discount_monthly_limit`.

A base de cálculo (`BaseValue`) nunca pode ser negativa.

**Fonte:** `services/calculation_service.go:145-241`, `data/init.sql:31-61`

### RF-05 — Reforma Tributária 2026
Quando `tax_code == "IRPF"` e `reference_date >= 2026-01-01`, o sistema deve aplicar regras de transição:

- **Isenção total:** Se `BaseValue <= transition_2026_floor`, o imposto é zerado.
- **Redução adicional:** Se `BaseValue` entre `floor` e `ceiling`, aplica `Reduction = fA − (fB × BaseValue)`, limitada ao valor do imposto (nunca negativo).

Os fatores `transition_2026_floor`, `transition_2026_ceiling`, `transition_2026_factor_a` e `transition_2026_factor_b` são obtidos dinamicamente via `GetTableConfigs` da biblioteca `taxnexus-individual-core-lib`.

**Fonte:** `services/calculation_service.go:274-296`

### RF-06 — Resiliência de Integração Externa
Quando a chamada ao serviço de INSS falha (erro de rede, timeout de 5s, ou status != 200), o sistema deve logar a falha com nível `Warn` e prosseguir o cálculo do modelo Completo sem a dedução previdenciária, sem abortar a requisição do contribuinte.

**Fonte:** `services/calculation_service.go:153-154`, `services/inss_client.go:20-23`

## 4. Requisitos Não-Funcionais

| ID | Descrição | Evidência |
|----|-----------|-----------|
| RNF-01 | Rastreabilidade: toda requisição recebe `X-Request-ID` (middleware) e o propaga para serviços externos | main.go:49, services/inss_client.go:34-36 |
| RNF-02 | Logging estruturado: `slog` com nível Debug, logs em JSON, trace_id incluso em todas as entradas | services/calculation_service.go:29-34, 38-43 |
| RNF-03 | Concorrência: cálculos Completo e Simplificado executam em goroutines com channel bufferizado | services/calculation_service.go:105-116 |
| RNF-04 | Timeout HTTP externo: 5 segundos para chamadas ao INSS | services/inss_client.go:22 |
| RNF-05 | Porta padrão: serviço escuta em `:3000` | main.go:55 |
| RNF-06 | Sem autenticação: o endpoint opera sem middleware de auth (rede interna presumida) | main.go:46-53 |

## 5. Critérios de Aceitação

### AC-001: Cálculo paralelo com recomendação
**Given** um request válido com `tax_code="IRPF"` e dados de renda
**When** o endpoint `POST /api/v1/calculate/irpf` é chamado
**Then** a resposta contém as chaves `"completa"` e `"simplificada"`, e exatamente uma delas tem `is_recommended: true`

### AC-002: Degradação graciosa do INSS
**Given** o serviço INSS está indisponível
**When** o cálculo é solicitado
**Then** o modelo Completo retorna resultado sem a dedução `deduction_social_security` e um warning é logado

### AC-003: Isenção Reforma 2026
**Given** `reference_date >= 2026-01-01` e `BaseValue <= transition_2026_floor`
**When** o cálculo é executado
**Then** `tax_amount = 0` e `deduction_details` contém `reforma_2026_isencao_total`

### AC-004: Limite PGBL
**Given** contribuição PGBL > 12% da renda bruta
**When** o modelo Completo processa deduções
**Then** a dedução PGBL é limitada a 12% da renda bruta

### AC-005: BaseValue não-negativa
**Given** deduções totais > renda bruta
**When** a base de cálculo é computada
**Then** `base_value = 0` (nunca negativo)
