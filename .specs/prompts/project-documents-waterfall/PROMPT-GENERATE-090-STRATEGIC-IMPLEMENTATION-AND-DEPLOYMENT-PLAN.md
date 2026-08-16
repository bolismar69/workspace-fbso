# PROMPT: GERADOR DE PLANO DE DEPLOYMENT
## Versão: 1.0 — WATERFALL Orchestrator

Atue como DevOps Engineer Sênior especializado em estratégias de deployment.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado (formato: `090-STRATEGIC-IMPLEMENTATION-AND-DEPLOYMENT-PLAN-{PROJECT_ID_NAME}.md`) |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `ARCHITECTURE_GLOBAL` | Caminho para a pasta de arquitetura global — ADRs, blueprints, padrões corporativos |
| `SECURITY_GLOBAL` | Caminho para o GLOBAL-SECURITY.md — regras de ouro, checklist SDD, threat model global |
| `TECHNICAL_SOLUTIONS` | Lista de soluções técnicas do projeto (`TECHNICAL_SOLUTION_NAMES`) — nomes dos microsserviços, frontends, batches |
| `PROJECT-STACK` | Stack tecnológica validada contra baseline corporativa em `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |
| `UPSTREAM_DOCS` | Lista: `[001-PROJECT-CHARTER, 030-SAD, 035-HLD, 040-LLD]` |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais fornecidos pelo humano (`PROJECT_DOCUMENTS_INPUTS`) |
| `SKILLS` | Lista de skills: ["deployment-engineer", "devops-rollout-plan"] |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima. Os parâmetros listados na tabela de Inputs são a única fonte de dados — não leia outros arquivos além dos explicitamente fornecidos
2. **LEIA** os documentos em `UPSTREAM_DOCS` — todos os artefatos devem rastrear de volta a eles
3. Skills: tente usar as skills listadas em `SKILLS` via `Skill` tool. Se falharem, use o template de fallback abaixo
4. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
5. Ao final, retorne `{DOC_PATH}` confirmando a criação

## Template de Fallback

```
# PLANO DE DEPLOYMENT: {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documentos Base** | 001-PROJECT-CHARTER, 030-SAD, 035-HLD, 040-LLD |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

### 1. Deployment Strategy
[Blue-green, canary ou rolling — com justificativa]

### 2. Environment Inventory
| Ambiente | URL | Recursos | Owner |
|----------|-----|---------|-------|
| Dev | ... | ... | ... |
| Staging | ... | ... | ... |
| Prod | ... | ... | ... |

### 3. Pre-deployment Checklist
| Item | Responsável | Critério |
|------|------------|---------|
| ... | ... | ... |

### 4. Deployment Steps
| Step | Ação | Responsável | Rollback Plan |
|------|------|------------|--------------|
| 1 | ... | ... | ... |

### 5. Database Migration Plan
| Migração | Versão | Script | Rollback Script |
|----------|--------|--------|----------------|
| M001 | v1.0 | ... | ... |

### 6. Rollback Plan
[Procedimento completo de rollback com gatilhos e tempos]

### 7. Communication Plan (Deploy)
| Público | Quando | Canal | Mensagem |
|---------|-------|------|---------|
| ... | ... | ... | ... |

### 8. Validation & Smoke Tests
| Teste | Comando/Script | Critério de Sucesso |
|-------|--------------|-------------------|
| ... | ... | ... |

### 9. Go-Live Runbook
[Passo a passo do go-live com horários e responsáveis]
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se o documento estiver completo.
