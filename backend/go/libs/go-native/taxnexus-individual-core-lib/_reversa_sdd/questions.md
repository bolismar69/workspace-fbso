# Perguntas para Validação Humana — taxnexus-individual-core-lib

> Modo de resposta: **file**.
> Status: ✅ **RESOLVIDO** (Todas as respostas incorporadas à spec em 2026-06-12)
> 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

---

## Bloco D — Detetive (domínio e regras de negócio)

### D1 ✅ — Tributos suportados 🟢
**Resposta:** IRPF, INSS

---

### D2 ✅ — Catálogo de `config_key` 🟢
**Resposta:** Chaves `pension_percentage`, `dependents_qty`, `pgbl_contribution`, `education_expenses`, `health_expenses` confirmadas via exemplos de API.

---

### D3 ✅ — Fórmula de cálculo (RN-03) 🟢
**Resposta:** `imposto = base × alíquota − parcela a deduzir` (Sim).

---

### D4 ✅ — Critério de recomendação (RN-06) 🟢
**Resposta:** Sempre recomendado o imposto de menor valor.

---

### D5 ✅ — Efeito de `TaxDefinition.Active` (RN-09) 🟢
**Resposta:** Reportado para o serviço consumidor.

---

### D6 ✅ — Invariante das faixas (RN-01 / ADR-0002) 🟢
**Resposta:** Garantido operacionalmente (Sim).

---

### D7 ✅ — Mensal vs. anual (RN-05) 🟢
**Resposta:** Aplicado pelo serviço consumidor.

---

### D8 ✅ — Autorização da camada consumidora (permissions.md) 🟢
**Resposta:** Controle existente nos serviços consumidores.

---

## Bloco A — Arquiteto (estrutura e recorte do monorepo)

### A1 ✅ (L1) — Camada de cálculo 🟢
**Resposta:** Código vive em monorepo original; responsabilidade externa à lib.

---

### A2 ✅ (L2) — Escrita do `tax_calculation_log` 🟢
**Resposta:** Estrutura preparada para evolução futura.

---

### A3 ✅ (L6) — DDL / migrations do schema `individual_tax_rates` 🟢
**Resposta:** DDL fornecido e incorporado ao `erd-complete.md`.

---

## Novas perguntas do Revisor 🔴

Nenhuma pergunta pendente no momento.
