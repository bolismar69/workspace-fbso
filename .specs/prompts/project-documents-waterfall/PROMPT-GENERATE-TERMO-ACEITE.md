# PROMPT: GERADOR DE TERMO DE ACEITE (SIGN-OFF)
## Versão: 1.0 — WATERFALL Orchestrator

Atue como Gerente de Projetos especializado em processos formais de aceite e encerramento.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `PROJECT_CTX` | Contexto do projeto: stack (`PROJECT-STACK`), arquitetura global (`ARCHITECTURE_GLOBAL`), segurança global (`SECURITY_GLOBAL`) |
| `TECHNICAL_SOLUTIONS` | Lista de nomes das soluções técnicas do projeto (`TECHNICAL_SOLUTION_NAMES`) |
| `UPSTREAM_DOCS` | Lista de caminhos para documentos upstream já em COMPLIANCE |
| `TEAM_SKILLS` | Skills mapeados para o time de implementação (`PROJECT-TEAM-SKILLS-MAP`) |
| `TEAM_CAPACITY` | Capacidade do time: seniores, plenos, juniores, duração (`PROJECT-TEAM-CAPACITY`) |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais fornecidos pelo humano (`PROJECT_DOCUMENTS_INPUTS`) |
| `SKILLS` | Lista de skills: ["contract-and-proposal-writer"] |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima. O contexto completo do projeto está em `PROJECT_CTX`, `TECHNICAL_SOLUTIONS`, `TEAM_SKILLS` e `TEAM_CAPACITY`
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
| **Documentos Base** | 01-PROJECT-CHARTER, 13-TEST-PLAN, 15-RELATORIO-QUALIDADE |
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
