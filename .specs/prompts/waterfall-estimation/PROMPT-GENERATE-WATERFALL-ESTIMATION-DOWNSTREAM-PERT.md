# PROMPT: GERADOR DE ESTIMATIVA PERT DOWNSTREAM/REFINEMENT (±15-25%)
## Versão: 1.0 — WATERFALL Estimation Orchestrator

Atue como Tech Lead e Arquiteto de Soluções especializado em estimativas PERT (Program Evaluation and Review Technique) baseadas em design detalhado e pacotes de trabalho.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `ARTIFACT_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `UPSTREAM_DOCS` | Documentos WATERFALL upstream: 03-SRS, 04-RTM, 07-LLD, 11-EAP/WBS |
| `WATERFALL_DOCS_PATH` | Caminho base dos documentos WATERFALL |
| `PROJECT-STACK` | Stack tecnológica validada contra baseline corporativa |
| `PROJECT-TEAM-SKILLS-MAP` | Skills disponíveis no time |
| `PROJECT-TEAM-CAPACITY` | Capacidade do time: seniores, plenos, juniores, duração prevista |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais |
| `SKILLS` | Lista de skills: ["project-estimation", "afrexai-construction-estimator", "senior-architect"] |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima
2. **LEIA** `UPSTREAM_DOCS` — extraia:
   - Da EAP/WBS (Doc #11): pacotes de trabalho (nível 3+), seus IDs, descrições e dependências
   - Do LLD (Doc #07): APIs, tabelas, diagramas de classe/sequência, complexidade técnica
   - Do SRS (Doc #03): requisitos funcionais e não-funcionais vinculados a cada pacote
   - Da RTM (Doc #04): rastreabilidade requisitos ↔ componentes
3. **ESTIMATIVA INDEPENDENTE:** Calcule do zero, sem usar ROM upstream como baseline ou referência
4. Skills: tente usar `SKILLS` via `Skill` tool. Se falharem, use o template de fallback
5. Crie o arquivo em `ARTIFACT_PATH` com `[STATUS: Em análise]`
6. Ao final, retorne `{ARTIFACT_PATH}`

## Metodologia PERT (Three-Point Estimation)

### Por pacote EAP/WBS:
- `O = Otimista` (melhor cenário — tudo corre bem)
- `M = Mais Provável` (cenário esperado — com imprevistos normais)
- `P = Pessimista` (pior cenário — com problemas significativos)
- **PERT Ponderado:** `E = (O + 4M + P) / 6`
- **Desvio Padrão:** `σ = (P − O) / 6`

### Consolidação:
- Soma dos E por fase WATERFALL (Design, Desenvolvimento, Testes, Deploy)
- σ consolidado = √(Σ σ²) (raiz quadrada da soma dos quadrados)
- Faixa de confiança: `[E − 1σ, E + 1σ]` = 68.3% de confiança

### Dimensões DTA obrigatórias por pacote:
- Desenvolvimento, QA, Arquitetura, DevOps/SRE, Gestão
- **Validação DTA interna:** QA ≥ 25% do Dev, Arch ≥ 5% do total

### Declaração de Independência:
Incluir seção explícita declarando que esta estimativa NÃO usou o ROM upstream como baseline.

## Template de Fallback

```
# ESTIMATIVA PERT DOWNSTREAM/REFINEMENT (±15-25%): {PROJECT_ID_NAME}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documentos Base** | 03-SRS, 04-RTM, 07-LLD, 11-EAP/WBS |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Modo** | DOWNSTREAM/REFINEMENT — PERT ±15-25% |
| **Metodologia** | Three-Point Estimation por Pacote EAP/WBS × Dimensões DTA |

---

### 1. Escopo Estimado

**Pacotes EAP/WBS considerados:**

| ID EAP | Pacote de Trabalho | Fase WATERFALL | Fonte (EAP) |
|--------|-------------------|---------------|-------------|
| 1.1 | {nome} | {fase} | EAP §{X} |
| 1.2 | {nome} | {fase} | EAP §{X} |

**Exclusões explícitas:**
- {item não estimado}
- {item não estimado}

---

### 2. Matriz PERT por Pacote EAP × Dimensões

#### 2.1 Desenvolvimento (horas)

| ID EAP | O | M | P | E = (O+4M+P)/6 | σ = (P−O)/6 |
|--------|---|---|---|-----------------|-------------|
| 1.1 | {h} | {h} | {h} | {E} | {σ} |
| 1.2 | {h} | {h} | {h} | {E} | {σ} |
| **Total Dev** | | | | **{ΣE}** | **√(Σσ²)** |

#### 2.2 QA (horas)

| ID EAP | O | M | P | E | σ |
|--------|---|---|---|---|---|
| 1.1 | {h} | {h} | {h} | {E} | {σ} |
| **Total QA** | | | | **{ΣE}** | **√(Σσ²)** |

#### 2.3 Arquitetura (horas)

| ID EAP | O | M | P | E | σ |
|--------|---|---|---|---|---|
| 1.1 | {h} | {h} | {h} | {E} | {σ} |
| **Total Arch** | | | | **{ΣE}** | **√(Σσ²)** |

#### 2.4 DevOps/SRE (horas)

| ID EAP | O | M | P | E | σ |
|--------|---|---|---|---|---|
| 1.1 | {h} | {h} | {h} | {E} | {σ} |
| **Total DevOps** | | | | **{ΣE}** | **√(Σσ²)** |

#### 2.5 Gestão (horas)

| ID EAP | O | M | P | E | σ |
|--------|---|---|---|---|---|
| 1.1 | {h} | {h} | {h} | {E} | {σ} |
| **Total Gestão** | | | | **{ΣE}** | **√(Σσ²)** |

---

### 3. PERT Consolidado por Fase WATERFALL

| Fase WATERFALL | Total E (h) | σ Consolidado | Faixa 68.3% [E−σ, E+σ] |
|---------------|-------------|---------------|--------------------------|
| Design (LLD) | {E} | {σ} | [{min}h, {max}h] |
| Desenvolvimento | {E} | {σ} | [{min}h, {max}h] |
| Testes | {E} | {σ} | [{min}h, {max}h] |
| Deploy | {E} | {σ} | [{min}h, {max}h] |
| **TOTAL** | **{ΣE}** | **√(Σσ²)** | **[{min}h, {max}h]** |

---

### 4. Caminho Crítico

**Duração total (caminho crítico):** {N} dias úteis

| ID | Atividade Crítica | Duração (dias) | Dependências |
|----|------------------|---------------|-------------|
| A1 | {atividade} | {d} | — |
| A2 | {atividade} | {d} | A1 |

---

### 5. Desvio Padrão e Faixa de Confiança

| Nível de Confiança | Multiplicador | Faixa |
|-------------------|--------------|-------|
| 68.3% (1σ) | ×1.00 | [{min}h, {max}h] |
| 95.4% (2σ) | ×2.00 | [{min}h, {max}h] |
| 99.7% (3σ) | ×3.00 | [{min}h, {max}h] |

**Precisão da estimativa:** ±{X}% (calculado como σ_total / E_total)

---

### 6. Validação DTA Interna

| Métrica | Valor | Limite | Status |
|---------|-------|--------|--------|
| QA / Dev | {X}% | ≥ 25% | ✅ / ⚠️ / 🔴 |
| Arch / Total | {X}% | ≥ 5% | ✅ / ⚠️ / 🔴 |

> ⚠️ Ações corretivas necessárias se alguma métrica estiver abaixo do limite.

---

### 7. Premissas por Pacote

| ID EAP | Premissa | Impacto se inválida |
|--------|----------|---------------------|
| 1.1 | {premissa} | {impacto} |
| 1.2 | {premissa} | {impacto} |

---

### 8. Independência da Estimativa

> **Declaração de Independência:** Esta estimativa PERT foi calculada exclusivamente a partir dos documentos WATERFALL (03-SRS, 04-RTM, 07-LLD, 11-EAP/WBS) usando Three-Point Estimation. NENHUM valor do ROM upstream (se existente) foi usado como baseline, ponto de partida ou referência para os cálculos. A estimativa é completamente independente e foi construída "do zero".

**Documentos utilizados (versões):**
- 03-SRS: v{X} — {DATA}
- 04-RTM: v{X} — {DATA}
- 07-LLD: v{X} — {DATA}
- 11-EAP/WBS: v{X} — {DATA}
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se o documento contiver todas as 8 seções, Three-Point para cada pacote EAP, validação DTA e declaração de independência.
