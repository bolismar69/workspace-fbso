# Spec Impact Matrix — taxnexus-individual-core-lib

> Gerado pelo **Arquiteto** (Reversa) em 2026-06-10 · `doc_level = completo`
> Objetivo: rastrear qual componente impacta quais entidades, regras e integrações — base para análise de impacto de mudanças.
> Confiança: 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

---

## 1. Componente × Entidade de dados

| Componente | `tax_definitions` | `tax_rules_history` | `tax_configs` | `tax_calculation_log` | Redis |
|------------|:---:|:---:|:---:|:---:|:---:|
| `repository.GetApplicableRule` | leitura (JOIN) | leitura | — | — | leitura (indireta) |
| `repository.GetTaxRulesForPeriod` | leitura (JOIN) | leitura | — | — | leitura/escrita |
| `repository.GetConfig` | — | — | leitura | — | — |
| `repository.GetTableConfigs` | — | — | leitura | — | leitura/escrita |
| `db.ConnectPostgres` | infra | infra | infra | infra | — |
| `cache.ConnectRedis` | — | — | — | — | infra |
| Serviço de Cálculo (externo) | leitura (`active`) | — | — | **escrita** 🔴 | — |

---

## 2. Componente × Regra de negócio

| Componente | Regras impactadas | Confiança |
|------------|-------------------|-----------|
| `GetApplicableRule` | RN-01 (seleção de faixa), RN-08 (decimal) | 🟢 |
| `GetTaxRulesForPeriod` | RN-02 (vigência), RN-01 (ordenação), RN-08 | 🟢 |
| `GetConfig` / `GetTableConfigs` | RN-02 (vigência), RN-04 (parâmetros de dedução) | 🟢 / 🟡 |
| Serviço de Cálculo (externo) | RN-03 (fórmula), RN-04 (deduções), RN-05 (mensal/anual), RN-06 (recomendado), RN-09 (active), RN-07 (rastreabilidade) | 🟢 (Confirmado via D3/D4/D5/D7) |

---

## 3. Componente × Integração / Config

| Componente | PostgreSQL | Redis | `DATABASE_URL` | `REDIS_ADDR` |
|------------|:---:|:---:|:---:|:---:|
| `repository` | depende | depende (cache-aside) | indireto | indireto |
| `db.ConnectPostgres` | cria pool | — | consome | — |
| `cache.ConnectRedis` | — | cria client | — | consome |

---

## 4. Matriz de impacto de mudança

| Se mudar… | Impacta diretamente | Impacto secundário |
|-----------|---------------------|--------------------|
| Schema `individual_tax_rates` (colunas/tipos) | `models`, todas as queries do `repository` | (de)serialização de cache, serviço consumidor |
| Estrutura de `TaxRule` | `GetApplicableRule`, `GetTaxRulesForPeriod`, chave `tax_rules_list:*` | cache existente fica inválido |
| Catálogo de `config_key` (D2) | `GetConfig`/`GetTableConfigs`, serviço de cálculo | fórmula de deduções no consumidor |
| TTL de cache (12h) | `GetTableConfigs`, `GetTaxRulesForPeriod` | frescor x carga no Postgres |
| Regra de vigência (RN-02) | todas as 4 funções de leitura | corretude de cálculo retroativo |
| Invariante de faixas contíguas (D6) | corretude de `GetApplicableRule` | resultados de cálculo no consumidor |
| Assinatura de `NewTaxRepository` | serviço consumidor (wiring) | — |

---

## 5. Pontos de maior acoplamento (hotspots)

| Hotspot | Por quê | Risco |
|---------|---------|-------|
| `repository` ↔ `models` | repository (de)serializa e mapeia todas as structs | alto — mudança em models propaga para queries e cache |
| Chaves de cache ↔ formato de data `YYYY-MM-DD` | granularidade diária acoplada à montagem de chave | médio — mudar granularidade invalida cache e vigência |
| Lib ↔ serviço de cálculo | fronteira pura/cálculo; contrato via API de funções Go | médio — alterar assinaturas quebra o consumidor |
| Schema implícito (sem migrations, L6) | nenhuma fonte de verdade versionada do DDL | alto — drift silencioso entre código e banco |
