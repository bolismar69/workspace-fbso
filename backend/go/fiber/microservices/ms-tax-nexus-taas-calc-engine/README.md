Vamos elevar o **Calc Engine** para o nível de produção. Para garantir a performance que o **Fiber** promete, utilizaremos o driver `pgx`, que é o padrão ouro para PostgreSQL em Go, focado em baixa alocação de memória e alta velocidade.

### 📂 Estrutura do Projeto

Dentro da pasta `backend/go/fiber/microservices/ms-tax-nexus-taas-calc-engine`:
```bash
cd $(git rev-parse --show-toplevel)/backend/go/fiber/microservices/ms-tax-nexus-taas-calc-engine
```


### 🛠️ Comandos para Inicializar o Projeto

No seu terminal Ubuntu, dentro da pasta do projeto:

1. **Inicie o módulo Go:**
```bash
go mod init ms-tax-nexus-taas-calc-engine

```
> Isso criará o arquivo go.mod
```
module ms-tax-nexus-taas-calc-engine
go 1.25.6
```


2. **Instale as dependências:**
```bash
go get github.com/gofiber/fiber/v2
go get github.com/jackc/pgx/v5
go get github.com/jackc/puddle/v2
go get github.com/stretchr/testify
go get golang.org/x/sync
go get github.com/davecgh/go-spew
go get gopkg.in/yaml.v3
go get github.com/pmezard/go-difflib
go get github.com/redis/go-redis/v9

```

3. **O go mod tidy garante que as dependências necessárias estejam sincronizadas e remove as não utilizadas.**
```bash
go mod tidy
```


4. **Execute o serviço:**
```bash
go run .

```



---

### 🧪 Testando a Integração

Com o serviço rodando, abra outro terminal ou use o navegador para testar a consulta baseada no nosso **Seed Data**:

```bash
curl http://localhost:3000/v1/simulate/62011100/3550308

```

**O que deve acontecer:**
O Fiber vai receber a chamada, o Pool do `pgx` vai buscar na `VW_TAX_SIMULATOR_REFORMA` (que já resolve a lógica de transição 2026/2027) e retornar o JSON exato com o cálculo.

---

### 🛡️ Próximo Passo sugerido para você

Agora que o **Calc Engine** já está conversando com o **Postgres**, o que você acha de **atualizarmos o Gateway (Gin)**?

Ele precisa ser o orquestrador que recebe a requisição do Portal React, envia para este Engine e depois grava o log de auditoria na `TB_TAX_TRANSACTION` que criamos. **Deseja que eu prepare a integração do Gateway chamando este Engine?**