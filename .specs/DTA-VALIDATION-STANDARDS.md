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

**Colunas obrigatórias (Discovery-Level):**
```
id_epico; titulo_epico; solucoes; horas_desenvolvimento; horas_arquitetura; horas_qa; prazo_entrega_meses; complexidade; comentarios
```

**Ação:** Coluna ausente ou vazia → ❌ REJEITADA.

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
| **Custo Total** | 30% | 35% | `ads-budget` |
| **Prazo de Entrega** | 30% | 25% | `project-estimation` |
| **Qualidade Técnica** | 25% | 25% | `estimate-builder-qmohd` |
| **QA/Arch Balanceado** | 15% | 15% | `analyst-estimates` |

**Justificativa (Discovery-Level):** No modo Discovery, prazo e custo têm peso igual (30% cada) porque o time-to-market é tão crítico quanto o orçamento na fase de viabilidade.

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
├── FACTORY-COMPARISON.md
├── FACTORY-NOTIFICATION.md
├── estimates/
│   └── ESTIMATION-SCHEMA-{FABRICA}.csv
└── notifications/
    └── FACTORY-NOTIFICATION-{FABRICA}.md

# Mode 2 — Full-Documentation
{PROJECT_PATH}/sourcing-factory-bidding-full/
└── (mesma estrutura)
```

---

## 5. ROM Baseline — Estimativa Interna de Referência

**Skill base:** `project-estimation` — Analogous estimation

O ROM interno é calculado pelo time de arquitetura durante o Upstream Discovery e serve como **baseline de comparação** — não é enviado às fábricas para evitar ancoragem (anchoring bias).

**Fórmula:**
```
ROM_Interno = Σ (Complexidade_Épico × Fator_Esforço)
onde Fator_Esforço: Alta = 8-14 h-m, Média = 6-10 h-m, Baixa = 2-5 h-m
```

---

## 6. Adaptações por Modo

| Aspecto | Discovery-Level | Full-Documentation |
|:---|:---|:---|
| **Unidade de estimativa** | Épico | Feature + User Story |
| **Schema columns** | 9 colunas (macro) | 9 colunas (detalhado) |
| **QA validation** | Por épico (QA ≥ 20%) | Por US (QA ≥ 20%) |
| **Outlier detection** | Total cross-fábrica | Por US cross-fábrica |
| **Precisão ROM** | ±50% | ±30% |
| **Pesos comparação** | Custo 30% / Prazo 30% | Custo 35% / Prazo 25% |
| **Pasta** | `sourcing-factory-bidding-discovery/` | `sourcing-factory-bidding-full/` |

---

## 7. Referências

| Documento | Relação |
|:---|:---|
| `DTA-Engine-de-Bidding-e-Estimativas.md` | Schema original e regras de ouro |
| `PROMPT-ROADMAP-GENERATE-SOURCING-FACTORY-BIDDING.md` | Roadmap de orquestração |
| `ESTIMATE-VALIDATION.md` (F5) | Aplicação das regras neste projeto |
| `FACTORY-COMPARISON.md` (F6) | Aplicação dos critérios neste projeto |

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial: consolidação de regras, fórmulas e padrões de validação e comparação. Baseado nos skills project-estimation, estimate-builder, estimate-builder-qmohd, analyst-estimates, ads-budget, trade-show-budget-planner, afrexai-construction-estimator. | Time de Arquitetura |

---

🤖 *Documento canônico de referência para todos os processos de Sourcing & Factory Bidding. Consultar este documento antes de executar qualquer fase do roadmap.*
