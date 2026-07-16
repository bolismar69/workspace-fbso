# Inventário do Projeto — ms-tax-individual-income

Gerado pelo agente **Scout** em 2026-06-08.

## 📂 Estrutura de Pastas

```
.
├── handlers/
│   └── tax_handler.go
├── services/
│   ├── calculation_service.go
│   └── inss_client.go
├── data/
│   └── init.sql
├── main.go
├── test_conn.go
├── go.mod
├── go.sum
├── README.md
├── CLAUDE.md
└── GEMINI.md
```

## 🛠️ Tecnologias e Frameworks

| Categoria | Tecnologia | Versão |
|-----------|------------|--------|
| Linguagem | Go | 1.25.6 |
| Framework Web | Fiber | v2.52.12 |
| Banco de Dados | PostgreSQL | v5.9.1 (pgx) |
| Cache | Redis | v9.18.0 |
| Biblioteca | shopspring/decimal | v1.4.0 |
| Biblioteca | taxnexus-individual-core-lib | local/replace |

## 🚀 Pontos de Entrada e Configuração

- **Entry Point:** `main.go` (Inicia servidor Fiber na porta 3000)
- **Configurações:** `.env` (DATABASE_URL, REDIS_ADDR, INSS_SERVICE_URL)
- **Docker:** Não detectado.
- **CI/CD:** Não detectado.

## 🗄️ Banco de Dados

- **SQL Inicial:** `data/init.sql`
- **Tabelas Identificadas:**
    - `individual_tax_rates.tax_definitions`
    - `individual_tax_rates.tax_rules_history`
    - `individual_tax_rates.tax_configs`

## 🧪 Cobertura de Testes

- **Arquivos de teste:** 0 encontrados (*_test.go).
- **Framework:** Go testing (nativo).
- **Status:** Sem testes automatizados no repositório.
