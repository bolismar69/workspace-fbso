# Glossário e Regras de Domínio — ms-tax-individual-income

Gerado pelo agente **Detective** em 2026-06-08. Atualizado em 2026-06-20 com evidência de código.

## ⚖️ Regras de Negócio Principais

### 1. Seleção do Modelo de Cálculo

- **Regra:** O sistema executa as lógicas Completa e Simplificada em paralelo via goroutines com channel bufferizado (`chan CalculationResult, 2`). Após ambos retornarem, compara `TaxAmount` e marca `IsRecommended = true` no modelo de menor imposto.
- **Fonte:** `services/calculation_service.go:105-140`

### 2. Limites do Modelo Completo

- **Educação:** Limitado ao teto legal configurado multiplicado por (1 + número de dependentes). O teto-base vem de `education_limit_monthly` ou `education_limit_annual` conforme `calculation_type`.
- **Saúde:** Sem limite explícito de teto. Aceita o valor integral informado em `health_expenses`.
- **PGBL:** Limitado a `pgbl_limit_percentage`% da renda bruta total (default 12% quando a config não está disponível). Se a unidade for `percentage`, o valor é convertido: `pgbl_spent = gross_income × (pgbl_contribution / 100)`.
- **Pensão oficial:** Aceita `pension_amount` (valor fixo) ou `pension_percentage` (% da renda bruta).
- **Dependentes:** Dedução = `dependents_qty × dependent_deduction_monthly` (ou `annual`). Default R$ 189,59 por dependente/mês quando a config não está disponível.
- **Fonte:** `services/calculation_service.go:164-213`

### 3. Desconto Simplificado

- **Regra:** 20% da Renda Bruta, respeitando o teto configurado (`simplified_discount_monthly_limit` ou `simplified_discount_annual_limit`). Default mensal R$ 564,80 quando a config não está disponível.
- **Fonte:** `services/calculation_service.go:224-241`

### 4. Integração com INSS (Previdência Social)

- **Regra:** O modelo Completo consulta obrigatoriamente o microserviço externo de INSS via POST HTTP. O resultado (`inssRes.TaxAmount`) é adicionado como dedução do tipo `deduction_social_security`.
- **Resiliência:** Se a chamada ao INSS falhar (erro de rede, timeout 5s, ou status HTTP != 200), o sistema loga um warning (`slog.Warn`) e prossegue o cálculo do modelo Completo sem a dedução previdenciária. A requisição do contribuinte **não é abortada**.
- **Rastreabilidade:** O `X-Request-ID` é propagado no header da chamada HTTP ao INSS.
- **Timeout:** 5 segundos configurados no `http.Client`.
- **Fonte:** `services/calculation_service.go:150-161`, `services/inss_client.go:14-53`, `main.go:36-38`

### 5. Motor de Cálculo Progressivo (runTaxMath)

- **Base de cálculo:** `BaseValue = GrossIncome − TotalDeductions`. Se negativa, é zerada.
- **Tabela progressiva:** A regra aplicável é buscada via `GetApplicableRule` com base em `BaseValue` e `ReferenceDate`.
- **Fórmula:** `TaxAmount = BaseValue × (AliqPercent / 100) − DeductionVal`. Se negativo, é zerado.
- **Alíquota efetiva:** `EffectiveRate = (TaxAmount / GrossIncome) × 100`.
- **Fonte:** `services/calculation_service.go:243-314`

### 6. Regra de Não-Negatividade

- **BaseValue:** `if baseValue.IsNegative() { baseValue = decimal.Zero }` — a base de cálculo nunca pode ser negativa.
- **TaxAmount:** `if taxAmount.IsNegative() { taxAmount = decimal.Zero }` — o imposto devido nunca pode ser negativo.
- **Reforma 2026:** A redução adicional nunca pode exceder o `taxAmount` corrente (`reduction = min(reduction, taxAmount)`).
- **Fonte:** `services/calculation_service.go:245-247, 270-272, 288-291`

### 7. Configurações Dinâmicas (Defaults)

O sistema utiliza valores padrão (hardcoded) quando as configurações do banco não estão disponíveis:

| Config Key | Default | Descrição |
|---|---|---|
| `dependent_deduction_monthly` | 189.59 | Dedução mensal por dependente |
| `simplified_discount_monthly_limit` | 564.80 | Teto mensal do desconto simplificado |
| `pgbl_limit_percentage` | 12 | Limite de dedução PGBL (% da renda bruta) |

**Fonte:** `services/calculation_service.go:179, 208-209, 232-233`
