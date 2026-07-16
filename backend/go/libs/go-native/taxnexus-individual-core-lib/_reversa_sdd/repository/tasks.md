# Repository — Tarefas de Implementação

> Sequência executável para reimplementar a camada de acesso a dados a partir do legado.
> Gerado pelo **Redator** (Reversa) em 2026-06-10 · `doc_level = completo`
> Fonte primária: `repository/tax_repository.go`. Ver `requirements.md` e `design.md` desta unit.
> 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

## Pré-requisitos

- [ ] Unit `models` disponível (`TaxRule`, `TaxDefinition`) — ver `models/design.md`
- [ ] Unit `db` disponível (`ConnectPostgres` → `*pgxpool.Pool`)
- [ ] Unit `cache` disponível (`ConnectRedis` → `*redis.Client`)
- [ ] Schema `individual_tax_rates` provisionado conforme DDL (ver `erd-complete.md` / resposta A3): tabelas `tax_definitions`, `tax_rules_history`, `tax_configs`
- [ ] Dados de faixas/configs vigentes para `IRPF` e `INSS` carregados (`tax_code` confirmados — D1)

## Tarefas

- [ ] **T-01** — Definir a struct `TaxRepository` com os campos `db *pgxpool.Pool` e `rdb *redis.Client`, e o construtor `NewTaxRepository(db, rdb) *TaxRepository` (injeção de dependência via construtor — ADR-0005).
  - Origem no legado: `repository/tax_repository.go:17`
  - Critério de pronto: repositório instanciável recebendo pool e cliente externos; sem estado de negócio interno (stateless)
  - Confiança: 🟢

- [ ] **T-02** — Implementar `GetTaxRulesForPeriod(ctx, taxCode, refDate) ([]models.TaxRule, error)`: query em `tax_rules_history` com JOIN em `tax_definitions` por `tax_code`, filtro de vigência `valid_from <= refDate AND (valid_to IS NULL OR valid_to >= refDate)`, ordenada por `range_min ASC`.
  - Origem no legado: `repository/tax_repository.go` (query de listagem, ~:104-124)
  - Critério de pronto: retorna apenas faixas vigentes na data, ordenadas por `range_min`; mapeia `range_max`/`valid_to` nulos para ponteiros nil
  - Confiança: 🟢

- [ ] **T-03** — Adicionar cache-aside a `GetTaxRulesForPeriod`: chave `tax_rules_list:{taxCode}:{YYYY-MM-DD}`; em hit, `json.Unmarshal` e retorno; em miss, consulta banco, serializa em JSON e grava no Redis com **TTL de 12h** (ADR-0003).
  - Origem no legado: `repository/tax_repository.go:97,149`
  - Critério de pronto: segunda chamada idêntica dentro de 12h é servida do cache (sem ida ao banco); formato de data da chave = `2006-01-02`
  - Confiança: 🟢

- [ ] **T-04** — Implementar `GetApplicableRule(ctx, taxCode, baseValue, refDate) (*models.TaxRule, error)`: obtém as faixas via `GetTaxRulesForPeriod` e **resolve em memória** a primeira faixa cujo `[range_min, range_max]` contém `baseValue` (`range_max == nil` ⇒ sem teto). Se nenhuma casa, retorna erro `no applicable rule found for value <baseValue>` (RN-01, ADR-0002).
  - Origem no legado: `repository/tax_repository.go:25,33-42`
  - Critério de pronto: base dentro de faixa retorna a faixa; base na faixa aberta superior retorna o último escalão; base sem cobertura retorna erro explícito (não `nil`)
  - Confiança: 🟢

- [ ] **T-05** — Implementar `GetTableConfigs(ctx, taxCode, refDate) (map[string]decimal.Decimal, error)`: query em `tax_configs` com filtro de vigência, montando `map[config_key]config_value`; cache-aside com chave `tax_configs:{taxCode}:{YYYY-MM-DD}` e TTL 12h.
  - Origem no legado: `repository/tax_repository.go:61-97`
  - Critério de pronto: mapa contém todas as `config_key` vigentes (ex.: `pension_percentage`, `dependents_qty`, tetos); cache hit na segunda chamada
  - Confiança: 🟢

- [ ] **T-06** — Implementar `GetConfig(ctx, taxCode, key, refDate) (decimal.Decimal, error)`: leitura **direta no PostgreSQL (sem cache)** de uma única `config_key` vigente; erro se a chave não existir na vigência.
  - Origem no legado: `repository/tax_repository.go:44-59`
  - Critério de pronto: chave vigente retorna o valor; chave ausente retorna erro de "não encontrado"; confirmadamente não consulta Redis
  - Confiança: 🟢

- [ ] **T-07** — Implementar **degradação graciosa** do cache: erro de leitura/escrita/`unmarshal` no Redis não é propagado — registra-se o miss e segue-se para o PostgreSQL (ADR-0006).
  - Origem no legado: `repository/tax_repository.go` (tratamento de erro do Redis nos fluxos com cache)
  - Critério de pronto: com Redis indisponível, `GetTableConfigs` e `GetTaxRulesForPeriod` ainda retornam dados corretos do banco sem erro propagado
  - Confiança: 🟢

## Tarefas de Teste

- [ ] **TT-01** — Happy path `GetApplicableRule`: base dentro de uma faixa IRPF vigente retorna a faixa correta (ver Critérios de Aceitação de `requirements.md`)
- [ ] **TT-02** — Faixa aberta superior: base acima de todos os `range_max` definidos retorna o último escalão (`range_max == nil`)
- [ ] **TT-03** — Erro: base sem faixa correspondente retorna `no applicable rule found for value <X>`
- [ ] **TT-04** — Vigência temporal: mesma consulta em `refDate` distintas (ex.: 2025-12-31 vs 2026-03-31) retorna conjuntos de faixas diferentes
- [ ] **TT-05** — Cache: primeira chamada popula o Redis (miss → set); segunda chamada idêntica é hit e não consulta o banco
- [ ] **TT-06** — Degradação graciosa: com Redis fora, `GetTableConfigs`/`GetTaxRulesForPeriod` leem do PostgreSQL sem erro
- [ ] **TT-07** — `GetConfig` ignora o cache (vai sempre ao banco) — verificar ausência de leitura/escrita no Redis

## Tarefas de Migração de Dados (se aplicável)

- [ ] **TM-01** — Garantir que `tax_rules_history` respeite a constraint `check_ranges` (`range_max IS NULL OR range_max > range_min`) e o índice `idx_tax_rules_validity (tax_definition_id, valid_from, valid_to)` — ver DDL em `erd-complete.md`
- [ ] **TM-02** — Garantir unicidade de `tax_configs (tax_code, config_key, valid_from)` conforme constraint do DDL

## Ordem Sugerida

1. **T-01** (struct/construtor) primeiro — base de tudo.
2. **T-02 → T-03** (listagem + cache) antes de **T-04**, pois `GetApplicableRule` depende de `GetTaxRulesForPeriod`.
3. **T-05** e **T-06** (configs) podem ser feitas em paralelo a T-02/T-04.
4. **T-07** (degradação graciosa) integrada junto de T-03/T-05; não deixar como etapa final isolada — é parte do contrato de disponibilidade.

## Lacunas Pendentes (🔴)

- Nenhuma lacuna bloqueante para esta unit após D1–D8/A1–A3. A **fórmula de cálculo** (RN-03) e a distinção **mensal/anual** (RN-05/D7) são da **camada consumidora** (A1), fora do escopo deste pacote.
- A escrita de `tax_calculation_logs` é **evolução futura** (A2) — não implementar nesta versão.
