# ESTIMATE-VALIDATION — Validação DTA de Estimativas (Full Mode)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Modo:** Full — Projeto Completo (62 US · 18 Features · 4 Épicos)
- **Data:** 31/07/2026
- **Baseline de Referência:** PERT Downstream: 6,077h (Dev) / 7,300h (c/ contingência 20%)

---

## 1. Regras de Validação DTA

| Regra | Critério | Ação se não atender |
|:---|:---|:---|
| **QA Balanceado** | QA ≥ 20% por épico | ⚠️ Risco de Débito Técnico |
| **QA Global** | QA ≥ 25% do total de horas | ❌ REJEITADA |
| **Arquitetura/SRE** | Arch ≥ 5% do total geral de horas | ❌ REJEITADA |
| **Formato** | Colunas obrigatórias preenchidas conforme schema | ❌ REJEITADA |
| **Consistência Prazo×Horas** | `prazo_calculado = total_horas / (time × 160h)`. Divergência >50% → ❌ | ❌ REJEITADA |
| **Outliers** | Total de horas dentro de ±50% da mediana cross-fábrica | 🔍 Revisão manual |
| **PIB (Proximidade Baseline)** 🆕 | PIB Score ≥ 0.25. Baseline: PERT Downstream 7,300h | ⚠️ <0.50 alerta / 🔴 <0.25 rejeitada |

---

## 2. Resultados por Fábrica

| # | Fábrica | Total Horas | QA% Dev | Arch% Tot | Prazo (meses) | Time Est. | Consistência Prazo | PIB Score 🆕 | Veredito |
|:---|:---|:---:|:---:|:---:|:---:|:---:|:---|:---:|:---|
| 1 | **Stefanini** | 28,030h | 4.6% | 4.6% | 3 | 15p | 🔴 Divergente (calc: 11.7) | 0.00 🔴 | 🔴 **REJEITADA** — QA+Arch+Prazo+PIB |
| 2 | **Capgemini** | 11,680h | 11.0% | 11.0% | 3 | 15p | 🔴 Divergente (calc: 4.9, +62%) | 0.40 🔴 | 🔴 **REJEITADA** — QA+Prazo+PIB |
| 3 | **CI&T** | 51,680h | 2.5% | 2.5% | 6 | 45p | ⚠️ Divergente (calc: 7.2, +20%) | 0.00 🔴 | 🔴 **REJEITADA** — QA+Arch+PIB |
| 4 | **Deloitte** | 11,680h | 11.0% | 11.0% | 3 | 15p | 🔴 Divergente (calc: 4.9, +62%) | 0.40 🔴 | 🔴 **REJEITADA** — QA+Prazo+PIB |
| 5 | **Infosys** | 11,680h | 11.0% | 11.0% | 3 | 15p | 🔴 Divergente (calc: 4.9, +62%) | 0.40 🔴 | 🔴 **REJEITADA** — QA+Prazo+PIB |
| 6 | **Overlabs** | 27,680h | 4.6% | 4.6% | 4 | 45p | ✅ Consistente (calc: 3.8, -4%) | 0.00 🔴 | 🔴 **REJEITADA** — QA+Arch+PIB |
| 7 | **TCS** | 131,680h | 1.0% | 1.0% | 9 | 93p | ✅ Consistente (calc: 8.8, -2%) | 0.00 🔴 | 🔴 **REJEITADA** — QA+Arch+Outlier+PIB |
| 8 | **TOTVS** | 91,680h | 1.4% | 1.4% | 6 | 80p | ⚠️ Divergente (calc: 7.2, +19%) | 0.00 🔴 | 🔴 **REJEITADA** — QA+Arch+PIB |

**Mediana cross-fábrica:** 27,855h

---

## 3. Análise

### 3.1 Problema Sistêmico: QA e Arquitetura Subdimensionados

**NENHUMA fábrica atendeu aos critérios mínimos de QA (≥25%) e Arquitetura (≥5%).** Todas as 8 fábricas concentraram ~95% do esforço em desenvolvimento, com QA e Arquitetura simbólicos (~1-11%). Isso é um **problema sistêmico** que indica:

1. As fábricas interpretaram `horas_dev` como o esforço total e preencheram QA/Arch/DevOps/Gestão como overhead fixo (valores idênticos de 320/320/320/360 em várias fábricas)
2. O schema CSV pode não ter deixado claro que QA, Arch, DevOps e Gestão devem ser **proporcionais ao esforço de desenvolvimento**
3. Capgemini, Deloitte e Infosys submeteram valores **idênticos** (11,680h, mesmo QA/Arch, mesmo prazo, mesmo time) — possível coordenação ou uso do mesmo template

### 3.2 Comparação com PERT Baseline

| Fonte | Horas |
|:---|---:|
| **PERT Downstream (F8)** | 7,300h (c/ 20% contingência) |
| Capgemini / Deloitte / Infosys | 11,680h (+60%) |
| Overlabs | 27,680h (+279%) |
| Stefanini | 28,030h (+284%) |
| CI&T | 51,680h (+608%) |
| TOTVS | 91,680h (+1,156%) |
| TCS | 131,680h (+1,704%) |

### 3.3 Achados Críticos

- **3 fábricas (Capgemini, Deloitte, Infosys)** submeteram dados **idênticos** — possível violação de independência do processo
- **CI&T** foi a única com consistência Prazo×Horas aceitável (<20%) mas falhou em QA/Arch
- **TCS** estimou 131,680h com 93 pessoas — 18× acima do PERT baseline
- **Overlabs** foi a única com consistência Prazo×Horas dentro da margem (4%) mas QA/Arch insuficientes

---

## 4. Veredito Final

| Resultado | Quantidade |
|:---|---:|
| 🟢 Aprovada | **0** |
| 🔴 Rejeitada | **8** |

**8/8 fábricas rejeitadas.** Motivo primário: QA e Arquitetura abaixo dos thresholds DTA em todas as estimativas.

---

## 5. Recomendação

Diante da rejeição unânime, recomenda-se:

1. **Reabrir o RFQ** com instruções EXPLÍCITAS sobre a proporção mínima de QA (≥25%) e Arquitetura (≥5%), com exemplos numéricos
2. **Adicionar fórmula de validação** no schema CSV (`total_horas` = `horas_dev` + `horas_qa` + `horas_arch` + `horas_devops` + `horas_gestao`) para evitar que as fábricas tratem QA/Arch como overhead fixo
3. **Solicitar breakdown do racional** de estimativa na coluna `comentarios` — atualmente todas usam texto genérico "seguimos especificamente o material reportado"
4. **Investigar a coincidência** Capgemini/Deloitte/Infosys (valores idênticos)

---

🤖 *Validação DTA — Fase 5 do Sourcing & Factory Bidding (Full Mode). 0/8 aprovadas.*
