# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-MILESTONES (F17)

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md` — o **roadmap de milestones técnicos** do projeto alinhado com USER-STORIES, FEATURES, EPICS, BRD e PROJECT-CHARTER. Define quando cada solução será construída, em que ordem e com quais dependências.

**Inputs upstream (Bloco C — F17):** Este artefato é o último do Bloco C e consolida os insumos de todos os blocos anteriores:
- **F16 (SPECS-DEFINITION):** `PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md` — consolidação técnica enxuta (principal insumo)
- **Bloco 0 (Product Def & Backlog & PRD):** `PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md` (F4) + `PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md` (F3)
- **Bloco A (People & Solutions):** `TEAM-SKILLS-MAP.md`, `TEAM-CAPACITY.md`
- **Bloco B (6 Disciplinas Técnicas):** `ARCHITECTURE-DEFINITION.md`, `DATA-ARCHITECTURE-DEFINITION.md`, `SECURITY-DEFINITION.md`, `DEVOPS-SRE-DEFINITION.md`, `TEST-STRATEGY-DEFINITION.md`, `INFRA-CLOUD-DEFINITION.md`
- **Bloco C:** `SOLUTIONS-CATALOG.md` (F13), `SOLUTIONS-MATRIX.md` (F14), `SOLUTIONS-STACK-MATRIX.md` (F15)
- **Documentos de Negócio:** Charter (marcos M1-M7), Features, User Stories

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

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
