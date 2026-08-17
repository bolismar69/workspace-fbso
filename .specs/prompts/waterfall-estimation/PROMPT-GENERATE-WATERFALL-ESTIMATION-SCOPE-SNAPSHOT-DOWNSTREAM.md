# PROMPT: GERADOR DE SCOPE SNAPSHOT DOWNSTREAM/REFINEMENT
## Versão: 1.0 — WATERFALL Estimation Orchestrator

Atue como Analista de Negócios e Gerente de Projetos especializado em congelamento de escopo detalhado e rastreabilidade.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `ARTIFACT_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `UPSTREAM_DOCS` | Documentos WATERFALL upstream: 03-SRS, 04-RTM, 07-LLD, 11-EAP/WBS |
| `INTERNAL_UPSTREAM` | Artefato de estimativa downstream: WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md (Fase 4) |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais |
| `SKILLS` | Lista de skills: ["gap-analysis", "business-analyst"] |

## Regras

1. **LEIA** `INTERNAL_UPSTREAM` — extraia todos os pacotes EAP estimados
2. **LEIA** `UPSTREAM_DOCS` — cruze cada pacote com SRS (§X), RTM (§Y), LLD (§Z), EAP (§W)
3. **LISTE EXPLICITAMENTE** exclusões
4. **INCLUA** a declaração de independência da estimativa
5. Crie o arquivo em `ARTIFACT_PATH` com `[STATUS: Em análise]`
6. Ao final, retorne `{ARTIFACT_PATH}`

## Template de Fallback

```
# SCOPE SNAPSHOT — DOWNSTREAM/REFINEMENT: {PROJECT_ID_NAME}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Estimativa Vinculada** | WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md v1.0 |
| **Data de Congelamento** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Modo** | DOWNSTREAM/REFINEMENT |

---

### 1. Pacotes EAP Estimados

| ID EAP | Pacote de Trabalho | Fase WATERFALL | Fonte (EAP §) | Estimado |
|--------|-------------------|---------------|--------------|----------|
| 1.1 | {nome} | {fase} | §{X} | ✅ |
| 1.2 | {nome} | {fase} | §{X} | ✅ |

**Total de pacotes estimados:** {N}

---

### 2. Exclusões Explícitas (NÃO Estimado)

| Item | Motivo | Fonte da Decisão |
|------|--------|-----------------|
| {item} | {motivo} | {doc / decisão} |

---

### 3. Matriz de Rastreabilidade (Pacote × Documento Fonte)

| ID EAP | 03-SRS | 04-RTM | 07-LLD | 11-EAP/WBS |
|--------|--------|--------|--------|-----------|
| 1.1 | §{X} | §{Y} | §{Z} | §{W} |
| 1.2 | §{X} | §{Y} | §{Z} | §{W} |

---

### 4. Versões dos Documentos Fonte

| Documento | Versão | Data | Status |
|-----------|--------|------|--------|
| 03-SRS | v{X} | {DATA} | COMPLIANCE |
| 04-RTM | v{X} | {DATA} | COMPLIANCE |
| 07-LLD | v{X} | {DATA} | COMPLIANCE |
| 11-EAP/WBS | v{X} | {DATA} | COMPLIANCE |

---

### 5. Independência da Estimativa

> Esta estimativa PERT foi calculada exclusivamente a partir dos documentos WATERFALL listados acima. NENHUM valor do ROM upstream (se existente) foi usado como baseline ou referência. O escopo congelado neste snapshot reflete exclusivamente os pacotes EAP/WBS do projeto.

---

### 6. Premissas de Escopo

| Premissa | Impacto se Inválida |
|----------|---------------------|
| {premissa} | {impacto na estimativa} |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se o snapshot contiver todas as 6 seções, todos os pacotes EAP listados e rastreabilidade completa.
