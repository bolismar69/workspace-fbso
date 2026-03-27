# Comandos para iniciar o go.mod

### Inicializa o módulo
```bash
go mod init taxnexus-core-lib
```

### Adiciona as dependências necessárias
```bash
go get github.com/shopspring/decimal
go get github.com/jackc/pgx/v5
go get github.com/redis/go-redis/v9
go get github.com/google/uuid
```

### Limpa e organiza as dependências
```bash
go mod tidy
```

### Testar conexao com banco de dados

#### criando variaveis de ambiente
```bash
export DATABASE_URL="postgres://worker_user:worker_pass@localhost:5432/worker_db?sslmode=disable&search_path=individual_tax_rates"
export REDIS_ADDR="localhost:6379"
```

#### executar teste

```bash
# 1. Garante que as dependências estão atualizadas
go mod tidy

# 2. Executa o teste de conexão
go run test_conn.go

```

