# Diagrama de Contexto — ms-billing-engine-tax-rates

```mermaid
flowchart LR
    consumidor["Consumidor\nSistema de faturamento que\nprecisa calcular tributos"]
    ms_tax["MS Billing Engine Tax Rates\nMotor de cálculo de tributos sobre faturamento (Go/Fiber)\nPorta via env var PORT (default :3000)\nTributos: IPI, ICMS, PIS/COFINS, CBS/IBS/IS"]
    pg[("PostgreSQL\nSchema: billing_tax_rates\nRegras de ICMS, IPI, PIS, COFINS")]
    redis[("Redis\nCache de regras fiscais")]

    consumidor -->|"POST /calculate\nJSON/HTTPS\nJWT (Kong/Keycloak)"| ms_tax
    ms_tax -->|"pgx\nConsultas de regras fiscais\ne exceções por NCM"| pg
    ms_tax -->|"go-redis\nCache de regras com TTL"| redis

    subgraph "Limite do Sistema"
        ms_tax
    end

    subgraph "Persistência"
        pg
        redis
    end
```

## Atores e Sistemas

| Ator/Sistema | Tipo | Protocolo | Descrição |
|-------------|------|-----------|-----------|
| Consumidor (sistema de faturamento) | Sistema interno | HTTPS | Envia documentos fiscais para cálculo de tributos |
| MS Billing Engine Tax Rates | Sistema (este serviço) | — | Motor de cálculo fiscal Go/Fiber, execução bifásica |
| PostgreSQL | Banco de dados | pgx (TCP) | Persistência de regras fiscais (7 tabelas + log) |
| Redis | Cache | TCP | Cache de regras fiscais |

## Fluxo de Requisição

1. Sistema consumidor envia `POST /calculate` com payload `DocumentoFiscalEntrada` (itens com SKU, NCM, valores, UFs origem/destino)
2. Middleware pipeline: `recover` → `requestid (W3C Trace Context)` → `auth (JWT Kong/Keycloak)` → `logger` → `metrics`
3. Handler faz parse e validação do payload (`input.Validate()`)
4. Engine executa Fase 1 (sequencial): IPI para cada item
5. IPI calculado é injetado como `ITEM_IPI_VALOR` nos detalhes do item
6. Engine executa Fase 2 (paralela via goroutines): ICMS, PIS/COFINS e Reforma (CBS/IBS/IS) simultaneamente
7. ICMS utiliza valor do IPI como parte da base de cálculo; PIS/COFINS utiliza valor do ICMS para exclusão da base
8. Reforma Tributária calcula CBS, IBS (estadual+municipal) e IS via consulta à tabela `iva_dual_rules`
9. Resposta `DocumentoFiscalSaida` com tributos calculados e UUID de transação retorna ao consumidor
