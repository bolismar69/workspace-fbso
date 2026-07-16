# TaxNexus Billing Core Lib

Biblioteca central de lógica tributária para o ecossistema TaxNexus.

## Funcionalidades
- **Precisão Decimal:** Cálculos usando `shopspring/decimal`.
- **Modelos Unificados:** Contratos de entrada e saída para transição da Reforma 2026.
- **Interfaces de Repositório:** Padronização de acesso a dados fiscais.

## Como usar
`go get taxnexus-billing-core-lib`


## ler varios arquivos e juntar em apenas um.

find . -type f -name "*.go" -exec cat {} + >> projeto_taxnexus-billing-core-lib.txt

find . -type f -name "*.go" -exec cat {} + >> projeto_ms-billing-engine-tax-rates.txt

find /caminho/do/diretorio -type f -exec cat {} + > arquivo_unico.txt

find /caminho/do/diretorio -type f | xargs cat > arquivo_unico.txt




find /home/bolismar/work/workspace-fbso/backend/go/fiber/microservices/ms-billing-engine-tax-rates -type f -name "*.go" -exec cat {} + >> projeto_ms-billing-engine-tax-rates.txt


find /home/bolismar/work/workspace-fbso/backend/go/libs/go-native/taxnexus-billing-core-lib -type f -name "*.go" -exec cat {} + >> projeto_taxnexus-billing-core-lib.txt

