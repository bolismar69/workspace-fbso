# PROMPT: GERADOR DE RTM FASE 2 — RASTREABILIDADE DE SISTEMA
## Versão: 1.1 — WATERFALL Orchestrator v2.0

Atue como um Analista de Requisitos de Sistema e Auditor de Rastreabilidade. Sua missão é criar a **RTM-FASE-2** que cruza a RTM-FASE-1 (rastreabilidade de negócio) com a SRS (requisitos de sistema), provando que todo requisito de sistema possui lastro em um requisito de negócio através da cadeia RTM-FASE-1, incluindo a correspondência entre protótipos UX/UI (016) e os requisitos de sistema.

## Inputs

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `UPSTREAM_DOCS` | `[015-RTM-FASE-1, 020-SRS, 016-PROTOTIPOS-UX-UI]` |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais |
| `SKILLS` | `["requirements-modeling", "requirements-validation", "requirements-engineering"]` |

## Regras

1. **LEIA** 015-RTM-FASE-1 (cadeia Objetivo Charter → REQ → FEAT/RN/UC), 020-SRS (FR-NN, NFR-NN) e 016-PROTOTIPOS-UX-UI (PROTO-NN — validar que telas/fluxos dos protótipos têm correspondência nos requisitos de sistema)
2. Monte a matriz que prova: cada requisito de sistema da SRS (FR-NN, NFR-NN) → item de negócio da RTM-FASE-1 (FEAT/RN/UC) → requisito de negócio (REQ do BRD) → objetivo do Charter
3. **Garantia de Cobertura Total (Zero Lacunas):** Prove que cada Requisito Funcional (FR-NN) e cada Requisito Não-Funcional (NFR-NN) da SRS possui correspondência na cadeia de rastreabilidade da RTM-FASE-1
4. **Bloqueio de Gold-Plating (Zero Órfãos):** Garanta que nenhum FR/NFR foi inventado na SRS sem lastro em requisito de negócio através da RTM-FASE-1
5. Use o template abaixo. Status inicial: `[STATUS: Em análise]`

## Template de Fallback

```
# RTM Fase 2 — Matriz de Rastreabilidade de Sistema: {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documentos Base** | 015-RTM-FASE-1, 020-SRS, 016-PROTOTIPOS-UX-UI |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

## RTM Fase 2 — Rastreabilidade de Sistema

A **RTM-FASE-2** é o instrumento de governança que conecta a **Linha de Base de Escopo Funcional** (selada pela RTM-FASE-1) à **Especificação de Requisitos de Software (SRS)**. Ela valida que todos os requisitos de sistema (funcionais e não-funcionais) derivam de requisitos de negócio, garantindo que nenhum requisito técnico seja órfão.

### Conexão com o Pipeline

- **UPSTREAM:** Consome a cadeia de rastreabilidade de negócio da 015-RTM-FASE-1 e os requisitos de sistema da 020-SRS
- **DOWNSTREAM:** Alimenta 030-SAD (arquitetura rastreável), 035-HLD e 040-LLD

### Objetivos da RTM-FASE-2

- **Cobertura Total:** Prova que cada FR/NFR da SRS tem um item de negócio correspondente na RTM-FASE-1 (FEAT/RN/UC → REQ → Objetivo do Charter)
- **Zero Órfãos:** Garante que nenhum FR/NFR da SRS foi criado sem requisito de negócio explícito
- **Rastreabilidade de Mão Dupla:** Do negócio ao sistema (forward) e do sistema ao negócio (backward), permitindo análise de impacto imediata em qualquer mudança de requisito

---

## Matriz de Rastreabilidade — Sistema (RTM-FASE-1 × SRS)

| Requisito de Negócio (BRD) | Item RTM-F1 (FRD) | Requisito Funcional (SRS) | Requisito Não-Funcional (SRS) | Cobertura |
|:---|:---|:---|:---|:---|
| REQ-01 — {descrição} | FEAT-01 — {descrição} | FR-01 — {descrição} | NFR-PERF-01 — {descrição} | ✅ Completa |
| REQ-01 — {descrição} | RN-02 — {descrição} | FR-02 — {descrição} | NFR-SEC-01 — {descrição} | ✅ Completa |
| REQ-02 — {descrição} | UC-03 — {descrição} | FR-03 — {descrição} | — | ✅ Completa |

---

## Análise de Cobertura

### Requisitos de Sistema (SRS)

| FR/NFR | Item RTM-F1 Vinculado | REQ BRD | Status |
|--------|----------------------|---------|--------|
| FR-01 | FEAT-01 | REQ-01 | ✅ Coberto |
| FR-02 | RN-02 | REQ-01 | ✅ Coberto |
| NFR-PERF-01 | FEAT-01 | REQ-01 | ✅ Coberto |
| NFR-SEC-01 | RN-02 | REQ-01 | ✅ Coberto |
| FR-NN | ... | ... | ✅/⚠️/❌ |

### Órfãos (Gold-Plating) — Itens da SRS sem lastro na RTM-FASE-1

| Item SRS | Tipo | Justificativa | Ação |
|----------|------|---------------|------|
| {se vazio: "NENHUM — Todos os itens da SRS possuem lastro na RTM-FASE-1 ✅"} | | | |

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | {DATA ATUAL} | Baseline inicial de rastreabilidade de sistema | Time de Engenharia |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se a matriz estiver completa, todos os FR/NFR da SRS cobertos pela cadeia RTM-FASE-1 → BRD, e nenhum item órfão encontrado.
