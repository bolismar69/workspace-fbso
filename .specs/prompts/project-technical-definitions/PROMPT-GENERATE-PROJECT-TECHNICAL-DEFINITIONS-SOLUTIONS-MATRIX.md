# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md` — a **matriz-mestra** que consolida todas as definições anteriores em uma tabela única: solução → responsável → repositório → stack → perfis → status.

**Inputs upstream (Bloco C — F14):** Este artefato consolida insumos de todas as fases anteriores na cadeia progressiva do Bloco C:
- **Inputs Globais do Roadmap** (todas as 8 variáveis)
- **Bloco 0 (Product Def & Backlog & PRD):** `INTAKE-LOG.md`, `DOR-ASSESSMENT.md`, `PRODUCT-BACKLOG-LIST.md`, `PRD-DEFINITION.md`
- **Bloco A (People & Solutions):** `TEAM-SKILLS-MAP.md` + `TEAM-CAPACITY.md` + `TEAM-CAPACITY-EXCEPTIONS.md`
- **Bloco B (6 Disciplinas Técnicas):** `ARCHITECTURE-DEFINITION.md` + `SECURITY-DEFINITION.md` + `DATA-ARCHITECTURE-DEFINITION.md` + `DEVOPS-SRE-DEFINITION.md` + `TEST-STRATEGY-DEFINITION.md` + `INFRA-CLOUD-DEFINITION.md`
- **F13 (Bloco C):** `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md` — catálogo de soluções (primeiro artefato do Bloco C)

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
Ler TODOS os artefatos das fases 1-8 + TEAM-CAPACITY.md + TEAM-CAPACITY-EXCEPTIONS.md.

### Passo 2 — Invocar Skills Especializadas
Invocar skills de mapeamento, portfolio e documentação.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md` com:
- **Tabela-mestra:** Solução | Tipo | Repositório/Pasta | Stack | Tech Lead | Time Alocado | Perfis Necessários | Status
- Referência ao TEAM-CAPACITY.md para nomes e contatos
- Referência ao TEAM-CAPACITY-EXCEPTIONS.md para exceções de capacidade
- Capacidade alocada vs. necessária por solução
- Indicadores: cobertura de skills (%), risco de gargalo, status geral
- Matriz RACI simplificada (Responsável, Autoridade, Consultado, Informado)

### Passo 4 — Validação Pós-Geração
Verificar: todas as soluções na matriz, responsáveis definidos, stacks referenciadas, capacidade calculada.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `reference-builder` | Construir matriz de referência cruzada | Mapeamento |
| 2 | `team-composition-analysis` | Analisar alocação de time por solução | People |
| 3 | `project-manager` | Validar alocação de recursos | PM |
| 4 | `track-management` | Estruturar tracking de soluções | Portfolio |
| 5 | `stakeholder-map` | Mapear stakeholders por solução | Stakeholder |
| 6 | `documentation-writer` | Redigir a matriz consolidada | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador da matriz de soluções | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
