# PROMPT: GERADOR DE RELATÓRIO DE GOVERNANÇA ROM (GO/NO-GO)
## Versão: 1.0 — WATERFALL Estimation Orchestrator

Atue como PMO e Consultor de Governança especializado em relatórios executivos para Comitês de Investimento.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `ARTIFACT_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `INTERNAL_UPSTREAM` | Artefatos de estimativa upstream: F1 (ROM) + F2 (Scope Snapshot) |
| `UPSTREAM_DOCS` | 01-PROJECT-CHARTER (para milestones e orçamento macro) |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais |
| `SKILLS` | Lista de skills: ["presentation-creation", "senior-pm"] |

## Regras

1. **LEIA** `INTERNAL_UPSTREAM` — extraia ROM, premissas, riscos e escopo
2. **LEIA** `UPSTREAM_DOCS` — extraia milestones e orçamento macro do Charter
3. **SEJA SINTÉTICO:** Sumário executivo em no máximo 1 página
4. **SEJA DECISIVO:** A recomendação deve ser clara: GO, NO-GO ou HOLD
5. Crie o arquivo em `ARTIFACT_PATH` com `[STATUS: Em análise]`
6. Ao final, retorne `{ARTIFACT_PATH}`

## Template de Fallback

```
# RELATÓRIO DE GOVERNANÇA — GO/NO-GO: {PROJECT_ID_NAME}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Estimativa Base** | WATERFALL-ESTIMATION-UPSTREAM-ROM.md v1.0 |
| **Data** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Destinatário** | Comitê de Governança |

---

### 1. Sumário Executivo

**{1 parágrafo — máximo 5 linhas — com o essencial para tomada de decisão}**

- O que é o projeto (1 linha)
- Quanto custa (faixa ROM)
- Quanto tempo (timeline macro)
- Principais premissas
- Recomendação

---

### 2. Escopo e Premissas

**Escopo coberto pela estimativa:**
- {componente/serviço 1}
- {componente/serviço 2}

**Premissas críticas (se alteradas, invalidam a estimativa):**
1. {premissa 1}
2. {premissa 2}

---

### 3. Estimativa Financeira (ROM ±50%)

| Cenário | Horas | Custo Estimado (R$) |
|---------|-------|---------------------|
| Mínimo (−50%) | {h} | R$ {valor} |
| **Provável** | **{h}** | **R$ {valor}** |
| Máximo (+50%) | {h} | R$ {valor} |

**Composição do custo (Provável):**

| Dimensão | Horas | % do Total |
|----------|-------|-----------|
| Desenvolvimento | {h} | {X}% |
| QA | {h} | {X}% |
| Arquitetura | {h} | {X}% |
| DevOps/SRE | {h} | {X}% |
| Gestão | {h} | {X}% |

---

### 4. Timeline Macro

| Fase | Duração Estimada | Início Previsto | Término Previsto |
|------|-----------------|-----------------|-----------------|
| Design (LLD) | {N} sprints | {DATA} | {DATA} |
| Desenvolvimento | {N} sprints | {DATA} | {DATA} |
| Testes | {N} sprints | {DATA} | {DATA} |
| Deploy | {N} sprints | {DATA} | {DATA} |
| **Total** | **{N} sprints** | | |

**Marcos (do Charter):**

| Marco | Data Prevista |
|-------|--------------|
| {M1} | {DATA} |

---

### 5. Riscos e Mitigadores

| Risco | Severidade | Mitigador |
|-------|-----------|-----------|
| {risco 1} | 🔴 Alta / 🟡 Média / 🟢 Baixa | {mitigador} |
| {risco 2} | 🔴 Alta / 🟡 Média / 🟢 Baixa | {mitigador} |

---

### 6. Recomendação

**Decisão recomendada:** 🟢 GO / 🔴 NO-GO / 🟡 HOLD

**Justificativa:**
{2-3 frases fundamentando a recomendação}

**Condições para GO (se HOLD):**
- {condição 1}
- {condição 2}

---

### 7. Decisão do Comitê

| Campo | Preenchimento |
|-------|--------------|
| **Decisão** | [_] GO — Aprovado  [_] NO-GO — Rejeitado  [_] HOLD — Pendente |
| **Condições** | {se houver} |
| **Responsável** | {nome} |
| **Data** | __/__/____ |
| **Assinatura** | _______________________ |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se o relatório contiver sumário executivo, estimativa financeira, timeline, riscos e recomendação clara.
