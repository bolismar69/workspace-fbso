# PROMPT: GERADOR DE REQUIREMENTS TRACEABILITY MATRIX (RTM) — FASE 1
## Versão: 1.0 — WATERFALL Orchestrator

Atue como um Analista de Rastreabilidade especializado em matrizes RTM.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado (formato: `015-RTM-FASE-1-{PROJECT_ID_NAME}.md`) |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `UPSTREAM_DOCS` | Lista: `[001-PROJECT-CHARTER, 002-STAKEHOLDER-MAP, 005-BRD, 010-FRD]` |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais fornecidos pelo humano (`PROJECT_DOCUMENTS_INPUTS`) |
| `SKILLS` | Lista de skills: `["requirements-modeling", "requirements-validation"]` |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima. Os parâmetros listados na tabela de Inputs são a única fonte de dados — não leia outros arquivos além dos explicitamente fornecidos
2. **LEIA** os documentos em `UPSTREAM_DOCS` (PROJECT-CHARTER, STAKEHOLDER-MAP, BRD e FRD) — a matriz deve rastrear cada elemento de volta aos documentos de origem
3. Skills: tente usar as skills listadas em `SKILLS` via `Skill` tool. Se falharem, use o template de fallback abaixo
4. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
5. Construa uma matriz **bidirecional** (forward FR→REQ→OBJ e backward OBJ→REQ→FR) com **zero órfãos** e **100% de cobertura** em ambas as direções
6. Ao final, retorne `{DOC_PATH}` confirmando a criação

## Template de Fallback (4 Seções)

```
# Requirements Traceability Matrix (RTM): {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento Base** | 001-PROJECT-CHARTER, 002-STAKEHOLDER-MAP, 005-BRD, 010-FRD |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---
## RTM — Requirements Traceability Matrix (FASE 1 — Rastreabilidade de Negócio)
A **RTM** foi dividida em duas fases. Esta é a **FASE 1 (015)**, gerada por padrão: rastreabilidade de negócio, vinculando objetivos do **001-PROJECT-CHARTER**, stakeholders do **002-STAKEHOLDER-MAP**, requisitos do **005-BRD** e funcionalidades do **010-FRD** (OBJ → REQ → FR). A **FASE 2 (025)** cobre a rastreabilidade técnica a partir do **020-SRS**, alimentando o desenho arquitetural (**030-SAD**).

### 1. Matriz de Rastreabilidade Bidirecional

Cada linha vincula um objetivo do Charter (OBJ-XX), um requisito de negócio do BRD (REQ-XX) e um requisito funcional do FRD (FR-XX). Nenhuma célula pode ficar vazia.

| OBJ (Charter) | REQ (BRD) | FR (FRD) | Status |
|---------------|-----------|----------|--------|
| OBJ-01 | REQ-01 | FR-01 | ✅ Vinculado |
| OBJ-01 | REQ-02 | FR-02 | ✅ Vinculado |
| OBJ-02 | REQ-03 | FR-03, FR-04 | ⚠️ Parcial |
| OBJ-03 | REQ-04 | — | ❌ Órfão |

Status: `✅ Vinculado` (todas as colunas preenchidas), `⚠️ Parcial` (vínculo incompleto, ex: REQ sem FR ou FR sem REQ), `❌ Órfão` (elemento sem vínculo em nenhuma das direções). Toda linha marcada ⚠️ ou ❌ deve ter o motivo documentado na Seção 4.

### 2. Cobertura Forward (FR → REQ → OBJ)

Todo FR-XX do FRD deve aparecer em pelo menos uma linha da matriz, vinculado ao seu REQ-XX (BRD) e, transitivamente, ao OBJ-XX (Charter).

| FR (FRD) | REQ (BRD) | OBJ (Charter) | Status |
|----------|-----------|---------------|--------|
| FR-01 | REQ-01 | OBJ-01 | ✅ Vinculado |

Regra: **100% dos FRs** cobertos. Qualquer FR sem REQ é órfão forward.

### 3. Cobertura Backward (OBJ → REQ → FR)

Todo OBJ-XX do Charter deve aparecer em pelo menos uma linha da matriz, vinculado ao seu REQ-XX (BRD) e ao FR-XX (FRD).

| OBJ (Charter) | REQ (BRD) | FR (FRD) | Status |
|---------------|-----------|----------|--------|
| OBJ-01 | REQ-01 | FR-01 | ✅ Vinculado |

Regra: **100% dos OBJs** cobertos. Qualquer REQ sem OBJ é órfão backward.

### 4. Resumo de Cobertura e Gaps

| Métrica | Valor |
|---------|-------|
| Total de OBJs (Charter) | N |
| Total de REQs (BRD) | N |
| Total de FRs (FRD) | N |
| Cobertura forward (FR → REQ → OBJ) | 100% |
| Cobertura backward (OBJ → REQ → FR) | 100% |
| Órfãos forward (FR sem REQ) | 0 |
| Órfãos backward (REQ sem OBJ) | 0 |
| Linhas ⚠️ Parcial | N — motivos: ... |
| Linhas ❌ Órfão | N — motivos: ... |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se a matriz estiver completa com as 4 seções, zero órfãos e 100% de cobertura bidirecional (forward e backward).
