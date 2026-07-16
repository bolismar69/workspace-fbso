# Domínio & Regras de Negócio — taxnexus-individual-core-lib

> Gerado pelo **Detetive** (Reversa) em 2026-06-10 · Atualizado pelo **Revisor** em 2026-06-12
> Fontes: `repository/tax_repository.go`, `models/tax_models.go`, artefatos do Arqueólogo e respostas humanas em `questions.md` (D1–D8).
> Confiança: 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

---

## 1. Contexto de domínio

O sistema pertence ao domínio de **tributação de pessoa física no Brasil** (IRPF, INSS e tributos correlatos). Esta biblioteca (`taxnexus-individual-core-lib`) é a **camada de acesso a parâmetros fiscais**: ela lê definições de impostos, faixas/escalões progressivos e parâmetros de configuração, com **vigência temporal** (a lei muda ao longo do tempo e cálculos retroativos precisam usar a regra da época).

🟢 O **motor de cálculo** que consome esses dados (`TaxRequest` → `TaxResponse`) **não está neste repositório**, residindo em serviços consumidores no monorepo original. As regras de cálculo (fórmula, soma de deduções, comparação de cenários) foram **confirmadas pelo usuário** (D3, D4, D7) e são aplicadas externamente a esta lib.

---

## 2. Glossário (ubiquitous language)

| Termo | Significado no domínio | Evidência | Confiança |
|-------|------------------------|-----------|-----------|
| **Imposto / Tributo** (`TaxDefinition`) | Tributo identificado por um código de negócio (`tax_code`), com esfera e precisão de arredondamento próprias | `models.TaxDefinition` | 🟢 |
| **`tax_code`** | Chave de negócio do imposto (valores suportados: `IRPF`, `INSS` 🟢 D1) | filtro em todas as queries | 🟢 |
| **Esfera (`sphere`)** | Nível de competência tributária: federal / estadual / municipal | campo `Sphere` | 🟢 |
| **Faixa / Escalão** (`TaxRule`) | Um intervalo `[range_min, range_max]` de base de cálculo com uma alíquota e uma parcela a deduzir | `models.TaxRule` | 🟢 |
| **Alíquota (`aliq_percent`)** | Percentual aplicado sobre a base dentro da faixa | campo `AliqPercent` | 🟢 |
| **Parcela a deduzir (`deduction_val`)** | Valor abatido do imposto na tabela progressiva (mecânica: `imposto = base × alíquota − parcela a deduzir` 🟢 D3) | campo `DeductionVal` | 🟢 |
| **Faixa aberta superior** | Último escalão sem teto (`range_max IS NULL`) | logic A1 no repository | 🟢 |
| **Vigência (`valid_from`/`valid_to`)** | Janela temporal em que uma regra/config está em vigor | filtro temporal nas queries | 🟢 |
| **Data de referência (`refDate`)** | Data que determina qual versão da regra/config aplicar | parâmetro no repositório | 🟢 |
| **Configuração (`tax_configs`)** | Parâmetro chave-valor (ex.: `dependents_qty`, `education_expenses` 🟢 D2) | `GetConfig`/`GetTableConfigs` | 🟢 |
| **Precisão de arredondamento (`rounding_precision`)** | Nº de casas decimais para arredondar o resultado (aplicado pelo consumidor) | campo `RoundingPrecision` | 🟢 |
| **Log de cálculo (`TaxCalculationLog`)** | Registro de auditoria (evolução futura — A2) | `models.TaxCalculationLog` | 🟢 |
| **Cenário recomendado (`IsRecommended`)** | Indica o cenário de **menor imposto devido** (🟢 D4) | campo `TaxResponse.IsRecommended` | 🟢 |

---

## 3. Regras de negócio extraídas

### RN-01 — Seleção de faixa progressiva por base de cálculo 🟢
**Regra:** dada uma base de cálculo e um imposto/data, a faixa aplicável é a **primeira** (na ordem `range_min` ASC) cujo intervalo contém a base.
- Intervalos são **fechados nos dois lados** (`>=` e `<=`).
- `range_max == NULL` ⇒ faixa sem teto (cobre qualquer valor ≥ `range_min`).
- Invariante: faixas são contíguas e sem sobreposição (garantia operacional 🟢 D6).
- **Fonte:** `GetApplicableRule`, `repository/tax_repository.go:33-40`.

### RN-02 — Vigência temporal (versionamento por data) 🟢
**Regra:** toda regra/config é válida apenas no intervalo `valid_from <= refDate AND (valid_to IS NULL OR valid_to >= refDate)`.
- Permite **cálculo retroativo** fiel à lei da época.
- Tabelas são append-only (histórico preservado).
- **Fonte:** cláusulas SQL em `repository/tax_repository.go`.

### RN-03 — Mecânica do imposto progressivo (alíquota − parcela a deduzir) 🟢
**Regra:** o cálculo segue o modelo brasileiro `imposto = base × aliq_percent − deduction_val`.
- Confirmado como padrão para IRPF e INSS. 🟢
- **Fonte:** Resposta humana D3.

### RN-04 — Deduções do IRPF (Simulador) 🟢
**Regra:** a base de cálculo do IRPF é a renda bruta menos deduções parametrizadas:
- Previdência (valor ou %), dependentes, educação (com teto), saúde (sem teto) e PGBL (com limite %).
- Os limites e valores vêm de `tax_configs` (chaves confirmadas em D2).
- **Fonte:** `TaxRequest` + Resposta humana D2.

### RN-05 — Cálculo mensal vs. anual 🟢
**Regra:** a distinção de período (`monthly`/`annual`) altera o conjunto de faixas e parâmetros aplicados.
- A lógica de seleção do tipo de cálculo reside no serviço consumidor. 🟢
- **Fonte:** Resposta humana D7.

### RN-06 — Comparação de cenários e recomendação 🟢
**Regra:** o motor calcula múltiplos cenários (ex.: desconto simplificado vs. deduções legais) e marca `IsRecommended = true` no de **menor imposto devido**. 🟢
- **Fonte:** Resposta humana D4.

### RN-07 — Rastreabilidade do cálculo 🟢
**Regra:** cada cálculo registra a regra aplicada e as configurações usadas para fins de explicabilidade ao usuário.
- **Fonte:** `TaxResponse.UsedConfigs` e `AppliedRule`.

### RN-08 — Precisão monetária exata 🟢
**Regra:** todo valor monetário e alíquota usa `decimal.Decimal` (nunca `float`), evitando erro de arredondamento.
- **Fonte:** `models` e DDL (A3: `numeric(18,4)`).

### RN-09 — Imposto pode estar inativo 🟢
**Regra:** `TaxDefinition.Active = false` sinaliza ao serviço consumidor que o imposto não deve ser processado.
- **Fonte:** Resposta humana D5.

---

## 4. Constantes e parâmetros de negócio embutidos

| Constante | Valor | Local | Natureza | Confiança |
|-----------|-------|-------|----------|-----------|
| TTL de cache | `12 * time.Hour` | `tax_repository.go` | Janela de estabilidade dos parâmetros | 🟢 |
| Formato de data cache | `2006-01-02` | `tax_repository.go` | Vigência diária | 🟢 |
| Schema do banco | `individual_tax_rates` | SQL queries | Isolamento de dados | 🟢 |

---

## 5. TODOs, FIXMEs e sinais de intenção não implementada

| Sinal | Evidência | Interpretação |
|-------|-----------|---------------|
| `TaxCalculationLog` definido mas não usado | `models:32` | 🟢 Confirmado como **evolução futura** (A2) para persistência de auditoria. |
| `documentoFiscalRequest` não exportada | `models:56,67` | 🟡 Smell: consumidores externos não conseguem instanciar `Inputs`. |
| Inconsistência de cache em `GetConfig` | `tax_repository.go:44` | 🟡 Única função de leitura que não utiliza a camada Redis. |

---

## 6. Lacunas resolvidas (Histórico) ✅

As lacunas D1–D8 e A1–A3 foram integralmente resolvidas via validação humana. Nenhuma lacuna de domínio permanece aberta para a Fase de Revisão.

---

## Referências cruzadas
- Visão Arquitetural: `architecture.md`
- Modelo de Dados: `erd-complete.md` (revisado com DDL real)
- ADRs retroativos: `adrs/`
