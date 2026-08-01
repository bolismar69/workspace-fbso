# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — SCOPE-SNAPSHOT (F11)
## Versão: 1.0 — Foto do Escopo Estimado (Snapshot Imutável)

Atue como um Configuration Manager e Analista de Requisitos especializado em baseline de escopo.

## OBJETIVO

Produzir uma FOTO IMUTÁVEL do escopo que foi incluído na estimativa PERT. Este documento NÃO planeja sprints, NÃO atribui tarefas, NÃO gera contratos técnicos. É um snapshot para auditoria futura: "o que estava no escopo quando estimamos em [data]?"

## INPUTS

1. **BOTTOM-UP-PERT-ESTIMATE.md** (F8) — lista de US estimadas
2. **05-USER-STORIES-{PROJECT_ID_NAME}.md** — RTM completa
3. **04-FEATURES-{PROJECT_ID_NAME}.md** — features e entregas
4. **03-EPICS-{PROJECT_ID_NAME}.md** — épicos

## REGRA CRÍTICA

⚠️ Este documento NÃO DISPARA planejamento de sprints ou entregas. O refinamento técnico das sprints e a geração de contratos são responsabilidade do roadmap `PROMPT-ROADMAP-GENERATE-PROJECT-TECHNICAL-DEFINITIONS.md`.

## ESTRUTURA DO DOCUMENTO

```markdown
# SCOPE-SNAPSHOT — Foto do Escopo Estimado

- **Data de Congelamento:** {data}
- **Estimativa Vinculada:** BOTTOM-UP-PERT-ESTIMATE.md (F8)
- **Hash do Escopo:** {checksum das US incluídas}
- **ATENÇÃO:** Este documento é um snapshot do escopo no momento da estimativa. O planejamento de sprints será feito no roadmap PROJECT-TECHNICAL-DEFINITIONS.

## 1. Escopo Incluído na Estimativa

| D# | EPIC-ID | FEATURE-ID | US-ID | Descrição | Complexidade | PERT (h) |
[62 linhas]

## 2. Sumário por Entrega
| Entrega | US | Horas PERT |
[7 entregas D1-D7 + Should Haves]

## 3. Sumário por Épico
| Épico | Features | US | Horas PERT |
[4 épicos]

## 4. Distribuição por Complexidade
| Complexidade | Qtd US | Horas PERT | % |
|:---|---:|---:|---:|
| Simples (1) | | | |
| Média (2) | | | |
| Complexa (3) | | | |

## 5. Itens Fora do Escopo (se houver)
[US ou features explicitamente excluídas da estimativa]

## 6. Registro de Imutabilidade
Este documento foi congelado em {data}. Qualquer alteração no escopo após esta data deve ser registrada como CHANGE REQUEST e requer reestimativa.
```

🤖 *Prompt gerador — Fase 11 do Downstream Architecture Refinement · Scope Snapshot*
