# PROMPT: GERADOR DE TERMO DE ACEITE (SIGN-OFF)
## Versão: 1.0 — WATERFALL Orchestrator

Atue como Gerente de Projetos especializado em processos formais de aceite e encerramento.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado (formato: `105-TERMO-ACEITE-{PROJECT_ID_NAME}.md`) |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `UPSTREAM_DOCS` | Lista: `[001-PROJECT-CHARTER, 045-TEST-PLAN, 055-RELATORIO-QUALIDADE]` |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais fornecidos pelo humano (`PROJECT_DOCUMENTS_INPUTS`) |
| `SKILLS` | Lista de skills: ["contract-and-proposal-writer"] |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima. Os parâmetros listados na tabela de Inputs são a única fonte de dados — não leia outros arquivos além dos explicitamente fornecidos
2. **LEIA** os documentos em `UPSTREAM_DOCS` — todos os artefatos devem rastrear de volta a eles
3. Skills: tente usar as skills listadas em `SKILLS` via `Skill` tool. Se falharem, use o template de fallback abaixo
4. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
5. Ao final, retorne `{DOC_PATH}` confirmando a criação

## Template de Fallback

```
# TERMO DE ACEITE (SIGN-OFF): {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documentos Base** | 001-PROJECT-CHARTER, 045-TEST-PLAN, 055-RELATORIO-QUALIDADE |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

### 1. Acceptance Criteria Checklist
| # | Critério (Charter Seção 4 + Test Plan) | Status | Evidência |
|---|---------------------------------------|--------|----------|
| 1 | ... | ✅/❌ | ... |

### 2. Deliverable Acceptance Status
| Entrega (Charter Seção 4) | Aceito? | Data | Observação |
|--------------------------|---------|------|-----------|
| D1 | Sim/Não | DD/MM/AAAA | ... |

### 3. Quality Gate Results
| Gate (Relatório Qualidade) | Status | Comentário |
|---------------------------|--------|-----------|
| Unit Test | GO/NO-GO | ... |

### 4. Punch List (Pendências)
| # | Item | Severity | Responsável | Prazo |
|---|------|---------|-------------|-------|
| 1 | ... | High/Medium/Low | ... | DD/MM/AAAA |

### 5. Formal Acceptance Statement
[Declaração formal de aceite do produto/serviço]

### 6. Signatures
| Nome | Papel | Data | Assinatura |
|------|------|------|-----------|
| ... | Sponsor | DD/MM/AAAA | ______________ |
| ... | Product Owner | DD/MM/AAAA | ______________ |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se o documento estiver completo.
