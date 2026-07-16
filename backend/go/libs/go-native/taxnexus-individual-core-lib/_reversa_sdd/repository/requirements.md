# Repository

> Camada de acesso a dados e resolução de regras fiscais com suporte a cache-aside e vigência temporal.

## Visão Geral
O módulo `repository` atua como o motor de busca de regras fiscais (faixas, alíquotas e deduções) e configurações do sistema. Ele abstrai a complexidade de busca por data de vigência e utiliza Redis para otimizar a performance de recuperação de tabelas fiscais.

## Responsabilidades
- Resolver a faixa de imposto aplicável a partir de um valor base e data de referência. 🟢
- Recuperar configurações específicas de tributos (ex: limites de dedução, valores por dependente). 🟢
- Gerenciar o ciclo de vida de cache (cache-aside) para tabelas de regras e configurações. 🟢
- Garantir que apenas regras vigentes na data de referência sejam consideradas. 🟢

## Regras de Negócio
- **RN-01: Resolução de Faixa Progressiva:** A faixa aplicável é a primeira (ordenada por `range_min` ASC) cujo intervalo fechado `[RangeMin, RangeMax]` contém o valor base; `RangeMax` nulo indica uma faixa aberta (sem teto). 🟢
- **RN-02: Vigência Temporal:** Um registro é considerado válido se `valid_from <= data_referencia` e (`valid_to` é nulo ou `valid_to >= data_referencia`). 🟢
- **RN-03: Cache-Aside de 12h:** Tabelas de regras e listas de configurações são cacheadas por 12 horas no Redis, usando chaves que incluem o código do tributo e a data de referência (YYYY-MM-DD). 🟢
- **RN-04: Transparência de Falha de Cache:** Caso o Redis esteja indisponível ou ocorra erro de desserialização, o repositório busca os dados diretamente no PostgreSQL sem interromper o serviço. 🟢
- **RN-05: Invariante de Faixas:** As faixas para um mesmo tributo e vigência devem ser contíguas e sem sobreposição. 🟢 (Regra confirmada pelo usuário — D6; garantia operacional do processo de manutenção das tabelas).

## Requisitos Funcionais

| ID | Requisito | Prioridade | Critério de Aceite |
|----|-----------|-----------|-------------------|
| RF-01 | Buscar regra aplicável por valor | Must | Retornar `aliq_percent` e `deduction_val` corretos para o valor informado na data. |
| RF-02 | Recuperar todas as regras de um período | Must | Retornar lista ordenada por `range_min` para suporte a cálculos progressivos (ex: INSS). |
| RF-03 | Recuperar mapa de configurações | Must | Retornar todas as chaves ativas (ex: `pension_percentage`, `dependents_qty`) para o tributo. |
| RF-04 | Recuperar configuração única | Should | Buscar valor diretamente no banco (bypass cache) para chaves isoladas. |

## Requisitos Não Funcionais

| Tipo | Requisito inferido | Evidência no código | Confiança |
|------|--------------------|---------------------|-----------|
| Performance | Cache Redis de 12h para tabelas pesadas | `repository/tax_repository.go:97, 149` | 🟢 |
| Escalabilidade | Uso de `pgxpool` para gerenciamento de conexões | `repository/tax_repository.go:17` | 🟢 |
| Disponibilidade | Degradação graciosa (Postgres como fallback do Redis) | `repository/tax_repository.go:66, 108` | 🟢 |
| Integridade | Precisão decimal exata via `shopspring/decimal`; colunas `numeric(18,4)` no banco | `models/tax_models.go` + DDL (A3) | 🟢 *(Revisor: corrigido — antes dizia "20 dígitos, 4 decimais" com evidência errada)* |

## Critérios de Aceitação

```gherkin
Dado que o valor base é 5000.00 e a data é 2025-01-01
Quando o repositório buscar a regra aplicável para IRPF
Então deve retornar a faixa que contém 5000.00 na vigência de jan/2025

Dado que o Redis está offline
Quando buscar configurações de tributo
Então o repositório deve ler diretamente do PostgreSQL e retornar o mapa de valores com sucesso

Dado que uma regra tem valid_to = '2025-12-31'
Quando buscar regras para a data '2026-01-01'
Então esta regra específica não deve ser incluída no resultado
```

## Prioridade (MoSCoW)

| Requisito | MoSCoW | Justificativa |
|-----------|--------|---------------|
| Resolução de faixa (`GetApplicableRule`) | Must | Função central para qualquer cálculo de imposto. |
| Filtro de vigência temporal | Must | Garante cálculos corretos retroativos e futuros. |
| Cache-Aside Redis | Should | Crítico para performance em escala, mas possui fallback. |
| Busca de configuração única | Could | Menos frequente que a busca de tabela completa. |

## Rastreabilidade de Código

| Arquivo | Função / Classe | Cobertura |
|---------|-----------------|-----------|
| `repository/tax_repository.go` | `TaxRepository` | 🟢 |
| `repository/tax_repository.go` | `GetApplicableRule` | 🟢 |
| `repository/tax_repository.go` | `GetTaxRulesForPeriod` | 🟢 |
| `repository/tax_repository.go` | `GetTableConfigs` | 🟢 |
