# Diagrama de Contexto — ms-tax-individual-income

```mermaid
flowchart LR
    contribuinte["Contribuinte\nPessoa física que deseja calcular seu imposto"]
    ms_tax["MS Tax Individual Income\nMicroserviço de cálculo de IRPF (Go/Fiber)\nPorta :3000"]
    ms_inss["MS INSS\nMicroserviço externo para cálculo de previdência social\nTimeout: 5s"]
    pg[("PostgreSQL\nSchema: individual_tax_rates\nRegras fiscais e configs")]
    redis[("Redis\nCache de regras\n(planejado)")]

    contribuinte -->|"POST /api/v1/calculate/irpf\nJSON/HTTPS\nSem autenticação"| ms_tax
    ms_tax -->|"POST /api/v1/calculate/inss\nJSON/HTTPS\nPropaga X-Request-ID"| ms_inss
    ms_tax -->|"pgx\nGetTableConfigs\nGetApplicableRule"| pg
    ms_tax -.->|"go-redis\nCache de tabelas progressivas\n(não implementado)"| redis

    subgraph "Limite do Sistema"
        ms_tax
    end

    subgraph "Sistemas Externos"
        ms_inss
    end

    subgraph "Persistência"
        pg
        redis
    end
```

## Atores e Sistemas

| Ator/Sistema | Tipo | Protocolo | Descrição |
|-------------|------|-----------|-----------|
| Contribuinte | Pessoa (usuário final) | HTTPS | Solicita cálculo de IRPF via API |
| MS Tax Individual Income | Sistema (este serviço) | — | Motor de cálculo fiscal Go/Fiber |
| MS INSS | Sistema externo | HTTPS | Microserviço de previdência social |
| PostgreSQL | Banco de dados | pgx (TCP) | Persistência de regras fiscais e configs |
| Redis | Cache | TCP | Cache de tabelas progressivas (planejado) |

## Fluxo de Requisição

1. Contribuinte envia `POST /api/v1/calculate/irpf` com dados financeiros
2. Middleware `requestid` gera `X-Request-ID`
3. Handler faz parse do body (`UniversalTaxRequest`), define `reference_date` default se ausente
4. Service carrega configurações fiscais do PostgreSQL (`GetTableConfigs`)
5. Service dispara goroutines para cálculos Completo e Simplificado em paralelo
6. Modelo Completo consulta INSS externo (com timeout 5s, propaga `X-Request-ID`)
7. Ambos os modelos executam `runTaxMath` com tabela progressiva e regras da Reforma 2026
8. Sistema compara resultados e marca `IsRecommended` no modelo mais barato
9. Resposta JSON com ambos os resultados retorna ao contribuinte
