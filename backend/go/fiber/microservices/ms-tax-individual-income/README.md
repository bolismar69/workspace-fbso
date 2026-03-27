
# Comandos para iniciar o go.mod

```bash
go mod init ms-tax-individual-income
```

```bash
go mod tidy
```

# Configuraçao ambiente de execucao e trabalho

### variaveis de ambiente a criar
```bash

# # Formato: postgres://USUARIO:SENHA@HOST:PORTA/BANCO?opcoes
export DATABASE_URL="postgres://worker_user:worker_pass@localhost:5432/worker_db?sslmode=disable&search_path=individual_tax_rates"

export REDIS_ADDR="localhost:6379"

```

### executar serviço
```bash
# Atualiza dependencias
go mod tidy

# Inicia o serviço
go run main.go

# Atualiza dependencias e Inicia o serviço
go mod tidy && go run main.go

```

### Testando chamada via curl
```bash

curl -X POST http://localhost:3000/api/v1/calculate/irpf \
     -H "Content-Type: application/json" \
     -d '{"salary": "5000.00"}'

```

```bash

# Calculo mensal
curl -X POST http://localhost:3000/api/v1/calculate/irpf \
     -H "Content-Type: application/json" \
     -d '{
          "calculation_type": "monthly",
          "gross_income": "5000.00",
          "pension_deduction": "500.00",
          "dependents_count": 0,
          "other_deductions": "0.00"
     }'

```

```bash
curl -X POST http://localhost:3000/api/v1/calculate/irpf \
     -H "Content-Type: application/json" \
     -d '{
          "calculation_type": "monthly",
          "reference_date": "2026-03-25T10:00:00Z",
          "gross_income": "5000.00",
          "pension_deduction": "500.00",
          "dependents_count": 0,
          "other_deductions": "0.00"
     }'

```bash

# Calculo anual
curl -X POST http://localhost:3000/api/v1/calculate/irpf \
     -H "Content-Type: application/json" \
     -d '{
          "calculation_type": "annual",
          "reference_date": "2025-12-31",
          "total_annual_income": "82000.00",
          "total_deductions": "15000.00",
          "discount_type": "simplified" 
     }'

```

```bash
curl -X POST http://localhost:3000/api/v1/calculate/irpf \
     -H "Content-Type: application/json" \
     -d '{
          "calculation_type": "annual",
          "reference_date": "2025-12-31T23:59:59Z",
          "gross_income": "80000.00",
          "pension_deduction": "8000.00",
          "dependents_count": 0,
          "other_deductions": "2000.00"
     }'
```

```bash
curl -X POST http://localhost:3000/api/v1/calculate/irpf      -H "Content-Type: application/json"      -d '{
          "tax_code": "IRPF",
          "reference_date": "2025-12-31T23:59:59Z",
          "calculation_type": "monthly",
          "gross_income": "36000.00",
          "inputs": [
            { "key": "pension_percentage", "value": "11.0", "unit": "percentage" },
            { "key": "dependents_qty", "value": "1", "unit": "count" },
            { "key": "pgbl_contribution", "value": "1200.11", "unit": "amount" },
            { "key": "education_expenses", "value": "3340.23", "unit": "amount" },
            { "key": "health_expenses", "value": "1233.23", "unit": "amount" }
          ]
     }' | jq .

curl -X POST http://localhost:3000/api/v1/calculate/irpf      -H "Content-Type: application/json"      -d '{
          "tax_code": "IRPF",
          "reference_date": "2026-03-31T23:59:59Z",
          "calculation_type": "monthly",
          "gross_income": "36000.00",
          "inputs": [
            { "key": "pension_percentage", "value": "11.0", "unit": "percentage" },
            { "key": "dependents_qty", "value": "1", "unit": "count" },
            { "key": "pgbl_contribution", "value": "1200.11", "unit": "amount" },
            { "key": "education_expenses", "value": "3340.23", "unit": "amount" },
            { "key": "health_expenses", "value": "1233.23", "unit": "amount" }
          ]
     }' | jq .

```

```bash
curl -X POST http://localhost:3000/api/v1/calculate/irpf      -H "Content-Type: application/json"      -d '{
          "tax_code": "IRPF",
          "reference_date": "2025-12-31T23:59:59Z",
          "calculation_type": "monthly",
          "gross_income": "360000.00",
          "inputs": [
            { "key": "pension_percentage", "value": "11.0", "unit": "percentage" },
            { "key": "dependents_qty", "value": "1", "unit": "count" },
            { "key": "pgbl_contribution", "value": "14333.78", "unit": "amount" },
            { "key": "education_expenses", "value": "8340.23", "unit": "amount" },
            { "key": "health_expenses", "value": "13450.23", "unit": "amount" }
          ]
     }' | jq .

curl -X POST http://localhost:3000/api/v1/calculate/irpf      -H "Content-Type: application/json"      -d '{
          "tax_code": "IRPF",
          "reference_date": "2026-03-31T23:59:59Z",
          "calculation_type": "monthly",
          "gross_income": "360000.00",
          "inputs": [
            { "key": "pension_percentage", "value": "11.0", "unit": "percentage" },
            { "key": "dependents_qty", "value": "1", "unit": "count" },
            { "key": "pgbl_contribution", "value": "14333.78", "unit": "amount" },
            { "key": "education_expenses", "value": "8340.23", "unit": "amount" },
            { "key": "health_expenses", "value": "13450.23", "unit": "amount" }
          ]
     }' | jq .

```

# Estrutura do microserviço

```
go/fiber/microservices/ms-tax-individual-income/
├── go.mod
├── main.go
├── handlers/
│   └── tax_handler.go
└── services/
    └── calculation_service.go
```

