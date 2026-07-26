# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-MILESTONES

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md` — o **roadmap de milestones técnicos** do projeto alinhado com USER-STORIES, FEATURES, EPICS, BRD e PROJECT-CHARTER. Define quando cada solução será construída, em que ordem e com quais dependências.

**Inputs upstream:** `PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md` (Fase 7) + `PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md` (Fase 4) + Documentos de negócio (Charter, Features, User Stories) + Catálogo de Soluções (Fase 2) + Stack Matrix (Fase 3).

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

### Passo 1 — Carregar Documentos Base
Ler Project Charter (marcos M1-M7), Features, User Stories, Catálogo de Soluções, PRD Definition, Specs Definition.

### Passo 2 — Invocar Skills Especializadas
Invocar skills de planejamento, roadmap, gestão de projetos e análise de riscos.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md` com:
- Linha do tempo com milestones técnicos (derivados dos marcos M1-M7 do Charter)
- Para cada milestone: soluções afetadas, features/user stories cobertas
- Dependências entre soluções (ordem de construção)
- Sequenciamento: o que construir primeiro, o que depende do quê
- Critérios de aceitação por milestone
- Riscos técnicos por milestone
- Capacidade estimada vs. necessária (referenciando TEAM-CAPACITY)

### Passo 4 — Validação Pós-Geração
Verificar: milestones alinhados com Charter, dependências documentadas, critérios de aceitação definidos, riscos mapeados.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `roadmap-planning` | Planejar roadmap técnico | Roadmap |
| 2 | `roadmap-communicator` | Comunicar milestones claramente | Roadmap |
| 3 | `decomposition-planning-roadmap` | Decompor marcos em milestones técnicos | Roadmap |
| 4 | `agile-sprint-planning` | Planejar sprints alinhadas aos milestones | Agile |
| 5 | `scrum-framework` | Estrutura Scrum para entregas | Agile |
| 6 | `project-manager` | Gestão do plano de milestones | PM |
| 7 | `senior-pm` | Supervisão sênior de planejamento | PM |
| 8 | `prioritization-advisor` | Priorizar milestones e soluções | Priorização |
| 9 | `risk-manager` | Identificar e mitigar riscos por milestone | Risco |
| 10 | `gap-analysis` | Identificar gaps de cobertura | Análise |
| 11 | `stakeholder-alignment-checker` | Alinhar milestones com stakeholders | Stakeholder |
| 12 | `documentation-writer` | Redigir o Milestones | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador dos milestones técnicos | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
