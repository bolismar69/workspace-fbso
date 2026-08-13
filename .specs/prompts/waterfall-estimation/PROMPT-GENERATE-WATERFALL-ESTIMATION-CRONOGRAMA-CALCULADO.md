# PROMPT: GERADOR DE CRONOGRAMA CALCULADO (DERIVADO DO PERT)
## Versão: 1.0 — WATERFALL Estimation Orchestrator

Atue como Planejador de Projetos especializado em cronogramas derivados de estimativas PERT e caminho crítico.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `ARTIFACT_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `INTERNAL_UPSTREAM` | Artefatos de estimativa: F4 (PERT) + F5 (Scope Snapshot) |
| `UPSTREAM_DOCS` | 01-PROJECT-CHARTER (para milestones) |
| `PROJECT-TEAM-CAPACITY` | Capacidade do time: N seniores, N plenos, N juniores |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais |
| `SKILLS` | Lista de skills: ["roadmap-planning", "project-estimation", "senior-pm"] |

## Regras

1. **LEIA** `INTERNAL_UPSTREAM` — extraia:
   - Do PERT (F4): durações E por pacote, caminho crítico, σ
   - Do Scope Snapshot (F5): pacotes EAP e fases WATERFALL
2. **LEIA** `UPSTREAM_DOCS` — extraia milestones do Charter
3. **CALCULE** durações: `dias = E / (tamanho_equipe × horas_dia)` onde `horas_dia = 6` (produtivas)
4. **RESPEITE** dependências do LLD e sequenciamento da EAP/WBS
5. Crie o arquivo em `ARTIFACT_PATH` com `[STATUS: Em análise]`
6. Ao final, retorne `{ARTIFACT_PATH}`

## Template de Fallback

```
# CRONOGRAMA CALCULADO — DERIVADO DO PERT: {PROJECT_ID_NAME}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Estimativa Base** | WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md v1.0 |
| **Capacidade do Time** | {N} seniores, {N} plenos, {N} juniores |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Modo** | DOWNSTREAM/REFINEMENT |

---

### 1. Lista de Atividades com Durações

| ID | Atividade | Pacote EAP | Duração PERT (h) | Equipe | Duração (dias) |
|----|----------|-----------|-----------------|--------|---------------|
| A1 | {nome} | 1.1 | {E} h | {N} pessoas | {E / (N×6)} d |
| A2 | {nome} | 1.2 | {E} h | {N} pessoas | {d} |

**Premissas de cálculo:**
- 6 horas produtivas/dia/pessoa
- Equipe dedicada (não compartilhada com outros projetos)
- {outras premissas}

---

### 2. Sequenciamento e Dependências

| Atividade | Depende de | Tipo | Folga (dias) |
|-----------|-----------|------|-------------|
| A2 | A1 | Finish-to-Start | 0 |
| A3 | A1, A2 | Finish-to-Start | {d} |

---

### 3. Caminho Crítico

**Duração total do projeto:** {N} dias úteis / {N} meses

**Sequência crítica:**
```
A1 → A2 → A5 → A8 → A10
```

| Atividade Crítica | Duração (dias) | Início | Término |
|------------------|---------------|--------|---------|
| A1 | {d} | {DATA} | {DATA} |
| A2 | {d} | {DATA} | {DATA} |

---

### 4. Cronograma (Datas)

| Atividade | Data Início | Data Fim | Folga Total (dias) |
|-----------|------------|---------|-------------------|
| A1 | DD/MM/AAAA | DD/MM/AAAA | 0 |
| A2 | DD/MM/AAAA | DD/MM/AAAA | {d} |

---

### 5. Diagrama de Gantt (Textual)

```
ATIVIDADE  | M1 | M2 | M3 | M4 | M5 | M6 | ...
A1 (1.1)   | ██ | ██ |    |    |    |    |
A2 (1.2)   |    | ██ | ██ |    |    |    |
A3 (2.1)   |    |    | ██ | ██ | ██ |    |
...
CRÍTICO    | ██ | ██ | ██ | ██ | ██ | ██ |
```

---

### 6. Marcos (Milestones)

| Marco | Data | Vinculado a Marco do Charter |
|-------|------|---------------------------|
| M1: Kickoff | DD/MM/AAAA | M1: Kickoff (§{X} do Charter) |
| M2: LLD Concluído | DD/MM/AAAA | — |
| M3: Dev Concluído | DD/MM/AAAA | M3: Code Freeze (§{X}) |
| M4: Testes Concluídos | DD/MM/AAAA | — |
| M5: Deploy Produção | DD/MM/AAAA | M5: Go-Live (§{X}) |

---

### 7. Alocação de Recursos por Período

| Período | Seniores | Plenos | Juniores | Total |
|---------|---------|--------|---------|-------|
| Mês 1 | {N} | {N} | {N} | {N} |
| Mês 2 | {N} | {N} | {N} | {N} |
| Mês 3 | {N} | {N} | {N} | {N} |

---

### 8. Compatibilidade com WATERFALL Doc #12

> **Instrução para o orquestrador WATERFALL:** Este artefato é consumido como `UPSTREAM_DOC` adicional pelo `PROMPT-GENERATE-CRONOGRAMA-GANTT.md`. As seções 1-7 acima fornecem os dados estruturados para o template do Documento #12 (Cronograma/Gantt) da sequência WATERFALL.
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se o cronograma contiver todas as 8 seções, durações calculadas a partir do PERT e caminho crítico identificado.
