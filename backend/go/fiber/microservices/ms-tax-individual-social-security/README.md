Para iniciarmos o **`ms-tax-individual-social-security`**, seguiremos a lógica de cálculo progressivo do INSS, que difere do IRPF por ser aplicado sobre o salário bruto em faixas sucessivas até um teto máximo.

Abaixo, apresento a proposta de configuração de dados e a estrutura de diretórios para este novo microserviço, mantendo a consistência com o ecossistema **TaxNexus**.

---

### 1. Definição de Dados (SQL)

Diferente do IRPF, o INSS possui um **Teto Salarial**. Se o salário exceder esse teto, a contribuição é fixada no valor máximo.

#### Tabelas de Regras (`tax_rules_history`)
As faixas para 2026 devem ser inseridas vinculadas a um novo `tax_code` (ex: `INSS`).
* **Faixa 1**: Até R$ 1.412,00 (7,5%)
* **Faixa 2**: De R$ 1.412,01 a R$ 2.666,68 (9%)
* **Faixa 3**: De R$ 2.666,69 a R$ 4.000,03 (12%)
* **Faixa 4**: De R$ 4.000,04 até o Teto (14%)

#### Tabelas de Configuração (`tax_configs`)
Precisamos de chaves específicas para controlar o comportamento do cálculo:
* `inss_contribution_ceiling`: O valor máximo do salário sobre o qual incide o INSS (Ex: R$ 7.786,02).
* `calculation_method`: Definido como `progressive` (para indicar o cálculo por fatias).

---

### 2. Estrutura de Pastas do Microserviço

Seguiremos o padrão estabelecido no `ms-tax-individual-income`, garantindo que a separação de responsabilidades facilite a manutenção no seu ambiente Ubuntu.

```text
backend/go/fiber/microservices/ms-tax-individual-social-security/
├── cmd/
│   └── server/
│       └── main.go           # Ponto de entrada, setup do Fiber e DI
├── handlers/
│   └── inss_handler.go       # Recebe a requisição e injeta o Trace ID
├── services/
│   └── inss_service.go       # Lógica matemática (fatias progressivas e teto)
├── go.mod                    # Módulo: ms-tax-individual-social-security
├── go.sum
└── .env                      # DATABASE_URL e REDIS_ADDR
```

---

### 3. Diferenciais na Lógica do Serviço

Diferente do serviço de renda, o `inss_service.go` precisará implementar:
1.  **Cálculo por Fatias**: O algoritmo deve subtrair o limite da faixa anterior da atual para calcular o imposto devido em cada "degrau".
2.  **Verificação de Teto**: Antes de iniciar as faixas, validar se o `GrossIncome` supera o `inss_contribution_ceiling` recuperado do `TaxRepository`.
3.  **Logs com Trace ID**: Manteremos o uso do `slog` com o `requestid` vindo do Fiber para rastreabilidade total.

---

### Próximo Passo
Gostaria que eu gerasse o arquivo **`services/inss_service.go`** com a lógica específica de fatias progressivas (acumulador de faixas) para o INSS, já utilizando o seu `TaxRepository` para buscar os valores?



# Comandos para iniciar o go.mod

```bash
go mod init ms-tax-individual-social-security
```

```bash
go mod tidy
```

# Suba o serviço:
export DATABASE_URL="postgres://user:pass@localhost:5432/taxnexus"
export REDIS_ADDR="localhost:6379"
go run main.go

# Execute um teste com CURL:

```bash
curl -X POST http://localhost:3001/api/v1/calculate/inss \
-H "Content-Type: application/json" \
-d '{
  "gross_income": 5000.00,
  "reference_date": "2026-03-26T00:00:00Z"
}'
```





