# User Stories — Consumo da `taxnexus-individual-core-lib`

> Histórias do ponto de vista do **serviço consumidor** (motor de cálculo / API fiscal) que integra esta biblioteca.
> Gerado pelo **Redator** (Reversa) em 2026-06-10 · `doc_level = completo`
> Esta lib é a base para a documentação dos serviços consumidores (A1).
> 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

> **Persona principal:** *Serviço de Cálculo Fiscal* — backend que recebe requisições de simulação (ex.: `POST /api/v1/calculate/irpf`), busca parâmetros nesta lib e executa a fórmula. A lib **não calcula**; entrega os parâmetros vigentes.

---

## US-01 — Resolver a faixa progressiva aplicável 🟢

**Como** serviço de cálculo fiscal
**Quero** obter a faixa de imposto (`aliq_percent`, `deduction_val`) aplicável a uma base e data
**Para** aplicar a fórmula `imposto = base × alíquota − parcela a deduzir` (D3)

**Critérios de aceitação:**
```gherkin
Dado um tax_code "IRPF", uma base de cálculo e uma reference_date
Quando chamo GetApplicableRule(ctx, "IRPF", baseValue, refDate)
Então recebo a TaxRule cujo [range_min, range_max] contém a base na vigência da data
E quando a base excede todos os tetos, recebo a faixa aberta superior
E quando nenhuma faixa cobre a base, recebo erro "no applicable rule found for value <X>"
```
Origem: `repository/tax_repository.go` (`GetApplicableRule`).

---

## US-02 — Calcular impostos retroativos com a regra da época 🟢

**Como** serviço de cálculo fiscal
**Quero** que toda leitura respeite a `reference_date`
**Para** simular corretamente períodos passados ou futuros (a lei muda no tempo — RN-02)

**Critérios de aceitação:**
```gherkin
Dado faixas/configs com vigências distintas para IRPF
Quando consulto com reference_date = 2025-12-31 e depois 2026-03-31
Então cada consulta retorna o conjunto vigente na respectiva data
```
Origem: filtro de vigência em todas as funções do `repository`. Confirma D7 (mensal/anual e mudanças de tabela aplicados pelo consumidor com base na data).

---

## US-03 — Obter os parâmetros de dedução do imposto 🟢

**Como** serviço de cálculo fiscal
**Quero** ler o mapa de configurações vigentes (`tax_configs`)
**Para** aplicar limites/valores de dedução (dependentes, PGBL, educação, saúde, previdência)

**Critérios de aceitação:**
```gherkin
Dado um tax_code "IRPF" e uma reference_date
Quando chamo GetTableConfigs(ctx, "IRPF", refDate)
Então recebo um map[config_key]decimal com as chaves vigentes
E o mapa pode conter pension_percentage, dependents_qty, pgbl_contribution, education_expenses, health_expenses (D2)
E uma segunda chamada idêntica dentro de 12h é servida do cache
```
Origem: `repository/tax_repository.go` (`GetTableConfigs`, `GetConfig`).

---

## US-04 — Suportar cálculo progressivo (ex.: INSS) 🟢

**Como** serviço de cálculo fiscal
**Quero** listar todas as faixas vigentes de um imposto em uma data
**Para** somar a contribuição faixa a faixa (cálculo progressivo)

**Critérios de aceitação:**
```gherkin
Dado um tax_code "INSS" e uma reference_date
Quando chamo GetTaxRulesForPeriod(ctx, "INSS", refDate)
Então recebo todas as faixas vigentes ordenadas por range_min ASC
```
Origem: `repository/tax_repository.go` (`GetTaxRulesForPeriod`; comentário "útil para cálculos progressivos como INSS").

---

## US-05 — Resiliência ao cache indisponível 🟢

**Como** serviço de cálculo fiscal
**Quero** que falhas do Redis não derrubem o cálculo
**Para** manter disponibilidade mesmo com o cache fora (ADR-0006)

**Critérios de aceitação:**
```gherkin
Dado que o Redis está indisponível
Quando consulto regras ou configurações
Então a lib lê do PostgreSQL e retorna os dados sem propagar erro de cache
```
Origem: degradação graciosa no `repository`; fábrica lazy em `cache/redis.go`.

---

## US-06 — Escolher o cenário recomendado (na camada consumidora) 🟡🔴

**Como** serviço de cálculo fiscal
**Quero** comparar cenários (ex.: dedução simplificada vs. completa) e marcar o de menor imposto
**Para** retornar `is_recommended = true` no resultado de menor imposto devido (D4)

**Critérios de aceitação:**
```gherkin
Dado dois cenários calculados para a mesma renda
Quando comparo o tax_amount de cada um
Então marco is_recommended = true no de menor imposto devido
```
> 🔴 **Fora do escopo desta lib** — a comparação e a fórmula vivem na camada consumidora (A1). A lib fornece apenas os parâmetros (faixas e configs). Story incluída para orientar a documentação do serviço consumidor.

---

## Notas de fronteira (A1)

- Filtro por `TaxDefinition.Active` (D5), distinção mensal/anual (D7), soma de deduções e escolha do recomendado (D3/D4) são responsabilidade do **serviço consumidor**.
- Autorização e isolamento multi-tenant (D8) existem nos serviços consumidores, não nesta lib.
- Persistência de auditoria em `tax_calculation_logs` é **evolução futura** desta lib (A2).
