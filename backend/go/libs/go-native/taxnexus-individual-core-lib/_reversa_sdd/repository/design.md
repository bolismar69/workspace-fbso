# Repository, Design Técnico

> Detalhamento da implementação do módulo de acesso a dados e resolução de regras fiscais.

## Interface

| Símbolo | Assinatura | Retorno | Observação |
|---------|-----------|---------|------------|
| `NewTaxRepository` | `(db *pgxpool.Pool, rdb *redis.Client)` | `*TaxRepository` | Construtor que recebe pools de conexão. 🟢 |
| `GetApplicableRule` | `(ctx, taxCode: string, baseValue: decimal, refDate: time)` | `(*models.TaxRule, error)` | Resolve a faixa progressiva em memória após buscar do banco/cache. 🟢 |
| `GetConfig` | `(ctx, taxCode: string, key: string, refDate: time)` | `(decimal, error)` | Busca configuração pontual direto no PostgreSQL (sem cache). 🟢 |
| `GetTableConfigs` | `(ctx, taxCode: string, refDate: time)` | `(map[string]decimal, error)` | Recupera todas as configs do tributo com suporte a cache-aside. 🟢 |
| `GetTaxRulesForPeriod` | `(ctx, taxCode: string, refDate: time)` | `([]models.TaxRule, error)` | Recupera lista ordenada de fatias de imposto com suporte a cache-aside. 🟢 |

## Fluxo Principal (Resolução de Faixa)
1. **Recuperação de Dados:** Chama `GetTaxRulesForPeriod` para obter todas as faixas vigentes para o `taxCode` na `refDate`. `repository/tax_repository.go:35` 🟢
2. **Iteração em Memória:** Percorre a lista de regras (que já vem ordenada por `range_min` ASC do banco). `repository/tax_repository.go:40` 🟢
3. **Avaliação de Intervalo:** 🟢
    - Verifica se `baseValue >= rule.RangeMin`.
    - Verifica se `rule.RangeMax == nil` (faixa aberta) OU `baseValue <= rule.RangeMax`.
4. **Retorno:** Retorna o ponteiro para a primeira `TaxRule` que satisfaça as condições, ou erro caso nenhuma faixa cubra o valor. `repository/tax_repository.go:42` 🟢

## Fluxo de Cache-Aside (Tabelas e Configurações)
1. **Chave de Cache:** Gera chave no formato `{prefixo}:{taxCode}:{YYYY-MM-DD}`. `repository/tax_repository.go:68, 110` 🟢
2. **Hit:** Tenta recuperar valor do Redis. Se existir e for válido, faz o `unmarshal` JSON e retorna. `repository/tax_repository.go:72, 112` 🟢
3. **Miss:** 🟢
    - Executa Query SQL no PostgreSQL filtrando por `valid_from` e `valid_to`.
    - Realiza o scan para o modelo Go.
    - Se houver resultados, serializa para JSON e salva no Redis com **TTL de 12 horas**. `repository/tax_repository.go:97, 143`

## Dependências
- **`models`**: Utiliza as structs `TaxRule` e `TaxDefinition` para mapeamento de dados. 🟢
- **`pgxpool`**: Gerenciamento de pool de conexões com PostgreSQL. 🟢
- **`go-redis`**: Cliente para interação com cache Redis. 🟢
- **`shopspring/decimal`**: Precisão numérica para valores fiscais. 🟢

## Decisões de Design Identificadas

| Decisão | Evidência no código | Confiança |
|---------|---------------------|-----------|
| Resolução de faixa em memória | `repository/tax_repository.go:40` | 🟢 |
| Cache-aside com fallback silencioso | `repository/tax_repository.go:73` | 🟢 |
| TTL fixo de 12 horas | `repository/tax_repository.go:97` | 🟢 |
| Uso de JSON para armazenamento no Redis | `repository/tax_repository.go:96` | 🟢 |
| Filtro de vigência temporal nas queries SQL | `repository/tax_repository.go:56, 85, 122` | 🟢 |

## Estado Interno
O `TaxRepository` é um componente stateless em relação aos dados de negócio, mantendo apenas as referências para os pools de conexão (`db` e `rdb`). `repository/tax_repository.go:17` 🟢

## Observabilidade
- **Tratamento de Erros:** Erros de banco e cache são propagados com contexto via `fmt.Errorf`. 🟢
- **Logs:** Não foram identificados logs explícitos no código legado (presume-se que a camada consumidora realize o log). 🟡

## Riscos e Lacunas
- 🟢 **Confirmado:** A responsabilidade de persistir o `tax_calculation_log` é uma evolução futura do módulo (A2), não estando presente neste recorte.
- 🟢 **Confirmado:** A diferença entre cálculos mensais e anuais é tratada pelo serviço consumidor (A1), não pelo repositório.
