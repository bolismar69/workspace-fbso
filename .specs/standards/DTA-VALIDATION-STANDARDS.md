# DTA-VALIDATION-STANDARDS — Regras, Fórmulas e Padrões de Validação e Comparação

## Versão: 1.0 — Documento Canônico de Referência

> **Escopo:** Este documento define TODAS as regras de validação, fórmulas de comparação e padrões de nomenclatura usados nos processos de Sourcing & Factory Bidding. É a referência única consultada pelos prompts GENERATE, GATE e FIX.

---

## 1. Metodologia de Estimativa Base

**Skills de referência:** `project-estimation`, `estimate-builder`, `afrexai-construction-estimator`

### 1.1 Bottom-Up Estimation (afrexai-construction-estimator)

A estimativa é construída **de baixo para cima**: cada épico/feature/US é estimado individualmente, e o total é a soma das partes. Isso permite:

- Rastreabilidade: cada hora está vinculada a um item de escopo
- Validação granular: inconsistências são detectadas no nível do item, não no total
- Comparação justa: fábricas diferentes estimam o mesmo escopo, permitindo comparação linha a linha

### 1.2 Rough Order of Magnitude (ROM)

**Fórmula ROM (project-estimation):**
```
ROM = Estimativa_Provável × (1 ± 0.50)
Faixa: [ROM_min = 0.50 × Provável, ROM_max = 1.50 × Provável]
```

**Aplicação por modo:**
| Modo | Nível de Detalhe | Precisão ROM |
|:---|:---|:---|
| Discovery-Level | Épicos | ±50% |
| Full-Documentation | Features + User Stories | ±30% |

---

## 2. Regras de Validação DTA (F5)

**Skills de referência:** `estimate-builder-qmohd`, `analyst-estimates`

### 2.1 Regra do QA Balanceado (DoD Gate)

**Origem:** `estimate-builder-qmohd` — Quality validation gate

**Fórmula:**
```
QA_Ratio_Épico = horas_qa_épico / horas_desenvolvimento_épico
QA_Ratio_Global = Σ horas_qa / Σ horas_desenvolvimento
```

**Critérios:**

| Nível | Modo Discovery | Modo Full |
|:---|:---:|:---:|
| Por épico/feature | QA ≥ 20% do desenvolvimento | QA ≥ 20% por User Story |
| Global | QA ≥ 25% do total | QA ≥ 25% do total |

**Ação:** Se abaixo do mínimo → ⚠️ Risco de Débito Técnico. Se QA < 10% → ❌ REJEITADA.

### 2.2 Regra da Arquitetura/SRE (Technical Investment Gate)

**Fórmula:**
```
Arch_Ratio = Σ horas_arquitetura_sre / Σ (horas_desenvolvimento + horas_arquitetura_sre + horas_qa)
```

**Critério:** Arch ≥ 5% do total geral de horas (ambos os modos).

**Ação:** Se abaixo de 5% → ⚠️ Risco de Subinvestimento Técnico. Se Arch < 2% → ❌ REJEITADA.

### 2.3 Regra de Consistência Prazo × Horas (Timeline Feasibility Gate)

**Origem:** `project-estimation` — Resource-loaded scheduling

**Fórmula:**
```
Prazo_Calculado = Total_Horas / (Tamanho_Time_Estimado × 160h)
Tamanho_Time_Estimado = Total_Horas / (Prazo_Declarado × 160h)
```

**Validação:**
```
Divergência = |Prazo_Declarado - Prazo_Calculado| / Prazo_Calculado
```

| Divergência | Veredito |
|:---|:---|
| ≤ 30% | ✅ Consistente |
| 30% - 50% | ⚠️ Alerta — revisar |
| > 50% | 🔴 Inconsistente — REJEITADA |

### 2.4 Regra de Detecção de Outliers (Anomaly Detection Gate)

**Origem:** `analyst-estimates` — Cross-source variance analysis

**Fórmula:**
```
Mediana_Cross_Fábrica = mediana(total_horas_fábrica_1, ..., total_horas_fábrica_N)
Limite_Inferior = Mediana × 0.50
Limite_Superior = Mediana × 1.50
```

**Critério:** Total de horas da fábrica deve estar dentro de ±50% da mediana cross-fábrica.

**Ação:** Fora do intervalo → 🔍 Revisão manual. Se > 2× o limite superior → ❌ REJEITADA.

### 2.5 Regra de Formato (Schema Compliance Gate)

**Critério:** Todas as colunas obrigatórias preenchidas conforme `ESTIMATION-SCHEMA.csv`.

**Colunas obrigatórias (Schema Unificado — Discovery e Full):**
```
fabrica; id_epico; titulo; features_codigos; qtd_features; user_stories_codigos; qtd_user_stories; horas_dev; horas_qa; horas_arch; horas_devops; horas_gestao; total_horas; prazo_entrega_meses; time_estimado_pessoas; valor_estimado; complexidade; stack_aderencia; premissas; comentarios
```

> 💡 Schema unificado para ambos os modos. `time_estimado_pessoas` e `valor_estimado` são **obrigatórios** — a FBSO.ORG NÃO infere esses valores. No modo discovery, `features_codigos` e `user_stories_codigos` podem ser preenchidos com os épicos (nível de detalhe disponível).

**Ação:** Coluna ausente ou vazia → ❌ REJEITADA.

### 2.6 Regra de Proximidade à Baseline Interna — PIB (Internal Baseline Proximity Gate) 🆕

**Origem:** `analyst-estimates` — Cross-source variance analysis + `project-estimation` — Analogous estimation

**Objetivo:** Medir o quão próxima a estimativa da fábrica está da estimativa interna de referência da empresa (baseline). A baseline interna **NÃO é enviada às fábricas** no RFQ para evitar viés de ancoragem (anchoring bias).

**Fórmula:**
```
PIB_Score = 1 − (|Factory_Hours − Internal_Baseline| / Internal_Baseline)
Clampado em [0, 1] onde:
  1.0 = correspondência exata com a baseline
  0.0 = 2× ou mais de desvio da baseline
```

**Fontes da Baseline Interna por Modo:**

| Modo | Baseline | Arquivo | Precisão |
|:---|:---|:---|:---|
| `discovery` | ROM Upstream | `upstream-architecture-discovery/DISCOVERY-LEVEL-ROM-ESTIMATE.md` | ±50% |
| `full` | PERT Downstream | `downstream-architecture-refinement/BOTTOM-UP-PERT-ESTIMATE.md` | ±15-25% |

> ⚠️ A baseline usada é o valor **total do cenário recomendado** (com contingência). Para o modo Full, é o valor da linha "Total Recomendado" do `RISK-ADJUSTED-ESTIMATE.md` se existir, ou o valor "TOTAL com Contingência" do `BOTTOM-UP-PERT-ESTIMATE.md`.

**Escala de Pontuação:**

| Desvio da Baseline | PIB Score | Nota | Significado |
|:---|---:|:---:|:---|
| 0–15% | 0.85–1.00 | **9–10** | Excelente — muito próximo da baseline interna |
| 15–30% | 0.70–0.85 | **7–8** | Bom — dentro da margem esperada |
| 30–50% | 0.50–0.70 | **5–6** | Regular — desvio significativo; revisar |
| 50–100% | 0.00–0.50 | **2–4** | Ruim — muito distante da baseline |
| >100% | 0.00 | **1** | Inaceitável — mais que o dobro da baseline |

**Critério DTA:**
| PIB Score | Veredito |
|:---|:---|
| ≥ 0.50 | ✅ Aceitável |
| 0.25 – 0.50 | ⚠️ Alerta — justificar divergência |
| < 0.25 | 🔴 Rejeitada — descolamento extremo da baseline interna |

**Observação:** Fábricas com PIB Score < 0.25 NÃO são automaticamente rejeitadas se todos os outros critérios (QA, Arch, Prazo, Formato) forem atendidos — mas a divergência deve ser justificada detalhadamente na coluna `comentarios`. Isso permite que uma fábrica demonstre que sua estimativa mais alta se justifica por abordagem técnica superior, escopo adicional identificado ou riscos não considerados na baseline.

---

## 3. Critérios de Comparação (F6)

**Skills de referência:** `ads-budget`, `trade-show-budget-planner`, `analyst-estimates`

### 3.1 Matriz de Decisão Ponderada (ads-budget)

**Origem:** `ads-budget` — Budget allocation and comparison methodology

**Fórmula da Nota Ponderada:**
```
Nota_Final = Σ (Nota_Critério_i × Peso_Critério_i)
onde Σ Pesos = 100%
```

### 3.2 Pesos por Critério

| Critério | Peso (Discovery) | Peso (Full) | Skill Base |
|:---|:---:|:---:|:---|
| **Custo Total** | 25% | 30% | `ads-budget` |
| **Prazo de Entrega** | 25% | 20% | `project-estimation` |
| **Qualidade Técnica (QA+Arch)** | 20% | 20% | `estimate-builder-qmohd` |
| **Proximidade à Baseline Interna (PIB)** 🆕 | **15%** | **15%** | `analyst-estimates` |
| **Consistência Prazo×Horas** | 15% | 15% | `analyst-estimates` |

**Justificativa:** O PIB recebe 15% — suficiente para valorizar fábricas alinhadas com a referência interna sem dominar os critérios técnicos. Se o peso fosse >20%, fábricas poderiam "chutar" próximo da baseline sem qualidade real. Se fosse <10%, a baseline interna não teria efeito prático no ranking.

### 3.3 Escala de Notas (trade-show-budget-planner)

**Origem:** `trade-show-budget-planner` — ROI scoring and go/no-go models

| Nota | Significado |
|:---:|:---|
| 10 | Melhor entre os concorrentes neste critério |
| 8-9 | Acima da média |
| 6-7 | Na média |
| 4-5 | Abaixo da média |
| 1-3 | Significativamente pior |

### 3.4 Go/No-Go por Fábrica (trade-show-budget-planner)

**Fórmula de Viabilidade:**
```
Viabilidade = (Custo ≤ Budget_Máximo) AND (Prazo ≤ Prazo_Máximo) AND (Nota_Final ≥ 5.0)
```

Se `Viabilidade = FALSE` → Fábrica entra na lista de rejeitadas independente da nota.

---

## 4. Padrões de Nomenclatura e Estrutura

### 4.1 Arquivos de Estimativa (F4)

**Pasta:** `estimates/`
**Padrão:** `ESTIMATION-SCHEMA-{NOME-DA-FABRICA}.csv`
**Exemplo:** `ESTIMATION-SCHEMA-STEFANINI.csv`

### 4.2 Arquivos de Notificação (F7)

**Pasta:** `notifications/`
**Padrão:** `FACTORY-NOTIFICATION-{NOME-DA-FABRICA}.md`
**Exemplo:** `FACTORY-NOTIFICATION-STEFANINI.md`

> ⚠️ **Confidencialidade:** O status da fábrica (selecionada, rejeitada, segundo colocado) NUNCA aparece no nome do arquivo. Informação confidencial — apenas no conteúdo.

### 4.3 Estrutura de Pastas por Modo

```
# Mode 1 — Discovery-Level
{PROJECT_PATH}/sourcing-factory-bidding-discovery/
├── RFQ-PACKAGE.md
├── ESTIMATION-SCHEMA.csv
├── FACTORY-DISTRIBUTION.md
├── ESTIMATE-RECEIPT.md
├── ESTIMATE-VALIDATION.md
├── [ESTIMATE-RETROSPECTIVE-PIB.md]          (F5b — condicional)
├── FACTORY-COMPARISON.md
├── FACTORY-NOTIFICATION.md
├── estimates/
│   ├── ESTIMATION-SCHEMA-{FABRICA}.csv
│   └── ESTIMATE-VALIDATION-{FABRICA}.md
└── notifications/
    └── FACTORY-NOTIFICATION-{FABRICA}.md

# Mode 2 — Full-Documentation
{PROJECT_PATH}/sourcing-factory-bidding-full/
└── (mesma estrutura)
```

---

## 5. Internal Baseline — Estimativa Interna de Referência

**Skill base:** `project-estimation` — Analogous estimation + Bottom-up estimation

A baseline interna é a estimativa de referência calculada pelo time de TI da empresa e serve como **parâmetro de comparação** na validação (F5) e comparação (F6). **Não é enviada às fábricas** no RFQ para evitar viés de ancoragem (anchoring bias).

### 5.1 Fontes por Modo

| Modo | Baseline | Arquivo | Método | Precisão |
|:---|:---|:---|:---|:---|
| **discovery** | ROM Upstream | `upstream-architecture-discovery/DISCOVERY-LEVEL-ROM-ESTIMATE.md` | Top-down por solução (S01-S15) | ±50% |
| **full** | PERT Downstream | `downstream-architecture-refinement/BOTTOM-UP-PERT-ESTIMATE.md` | Bottom-up PERT por US (62 US) | ±15-25% |

### 5.2 Valor de Referência

O valor usado como baseline é o **total do cenário recomendado** (com contingência):

| Modo | Fonte do Valor | Exemplo (PRJ-FIN-2026-0003) |
|:---|:---|:---|
| **discovery** | Linha "Provável" do ROM | 6,080h / 38 h-m |
| **full** | `RISK-ADJUSTED-ESTIMATE.md` → "Total Recomendado", ou `BOTTOM-UP-PERT-ESTIMATE.md` → "TOTAL com Contingência" | 7,300h / 46 h-m |

### 5.3 Uso na Validação (Regra PIB — §2.6)

A baseline alimenta a **Regra de Proximidade à Baseline Interna (PIB)** que mede o quão próxima a estimativa da fábrica está da referência interna. Ver §2.6 para fórmula completa, escala de pontuação e critérios de aceitação.

---

## 6. Adaptações por Modo

| Aspecto | Discovery-Level | Full-Documentation |
|:---|:---|:---|
| **Unidade de estimativa** | Épico | Feature + User Story |
| **Schema columns** | 20 colunas (schema unificado — §2.5) | 20 colunas (schema unificado — §2.5) |
| **QA validation** | Por épico (QA ≥ 20% dev) | Por US (QA ≥ 20% dev) |
| **Outlier detection** | Total cross-fábrica | Por US cross-fábrica |
| **Precisão ROM** | ±50% | ±30% (PERT) |
| **Pesos comparação (§3.2)** | Custo 25% / Prazo 25% / Qualidade 20% / PIB 15% / Consistência 15% | Custo 30% / Prazo 20% / Qualidade 20% / PIB 15% / Consistência 15% |
| **Pasta** | `sourcing-factory-bidding-discovery/` | `sourcing-factory-bidding-full/` |

---

## 7. Referências

| Documento | Relação |
|:---|:---|
| `.specs/standards/DTA-Engine-de-Bidding-e-Estimativas.md` | Schema original e regras de ouro |
| `PROMPT-ROADMAP-GENERATE-SOURCING-FACTORY-BIDDING.md` | Roadmap de orquestração |
| `ESTIMATE-VALIDATION.md` (F5) | Aplicação das regras neste projeto |
| `FACTORY-COMPARISON.md` (F6) | Aplicação dos critérios neste projeto |

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial: consolidação de regras, fórmulas e padrões de validação e comparação. Baseado nos skills project-estimation, estimate-builder, estimate-builder-qmohd, analyst-estimates, ads-budget, trade-show-budget-planner, afrexai-construction-estimator. | Time de Arquitetura |
| 1.1 | 31/07/2026 | **Auditoria de integridade:** §6 corrigido (9→20 colunas, pesos 30%→25% alinhados com §3.2). §4.3 atualizado com ESTIMATE-RETROSPECTIVE-PIB.md e ESTIMATE-VALIDATION-{FAB}.md na estrutura de diretórios. | Time de Arquitetura |

---

🤖 *Documento canônico de referência para todos os processos de Sourcing & Factory Bidding. Consultar este documento antes de executar qualquer fase do roadmap.*
